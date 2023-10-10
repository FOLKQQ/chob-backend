// main.go

package main

import (
	admincontroller "backend/controllers/adminController"
	authcontrollers "backend/controllers/authController"
	billingcontroller "backend/controllers/billingController"
	chatcontroller "backend/controllers/chatController"
	companycontroller "backend/controllers/companyController"
	rolecontrollers "backend/controllers/roleController"
	servicecontroller "backend/controllers/serviceController"
	tagController "backend/controllers/tagController"
	taskController "backend/controllers/taskController"
	taxcontroller "backend/controllers/taxController"
	teamcontroller "backend/controllers/teamController"
	"backend/database"
	//middlewarejwt "backend/middleware"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(503)
		w.Write([]byte("bad"))
	}
}

func checkrole(w http.ResponseWriter, r *http.Request) {

	dotenv := goDotEnvVariable("SecretKey")
	tokenString := r.Header.Get("Authorization")[7:]
	// สร้างตัวแปรเพื่อเก็บผลลัพธ์ที่ได้จากการตรวจสอบ token
	claims := jwt.MapClaims{}
	// สร้างตัวแปรเพื่อเก็บผลลัพธ์ที่ได้จากการตรวจสอบ token
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(dotenv), nil
	})
	// ตรวจสอบว่า token มีข้อผิดพลาดหรือไม่
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	Role := token.Claims.(jwt.MapClaims)["role_id"].(float64)
	// return Role
	fmt.Println(Role)

}

func main() {
	// Initialize the database connection
	db, err := database.InitDB()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.AllowContentEncoding("deflate", "gzip"))
	r.Use(middleware.AllowContentType("application/json", "text/xml"))
	r.Use(middleware.CleanPath)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Post("/loginadmins", func(w http.ResponseWriter, r *http.Request) {
		authcontrollers.LoginAdmins(w, r, db)
	})

	r.Route("/admins", func(r chi.Router) {
		//r.Use(middlewarejwt.ValidateToken)
		//r.Use(middlewarejwt.Rolesv)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {

			admincontroller.ListAdmin(w, r, db)
		})
		r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
			admincontroller.AddAdmin(w, r, db)
		})
		r.Delete("/update", func(w http.ResponseWriter, r *http.Request) {
			admincontroller.UpdateAdmin(w, r, db)
		})
	})
	r.Route("/admins/dashboard", func(r chi.Router) {
		//r.Use(middlewarejwt.ValidateToken)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			admincontroller.DashboardAdmin(w, r, db)
		})
		r.Get("/statuswork", func(w http.ResponseWriter, r *http.Request) {
			admincontroller.StatusWork(w, r, db)
		})
	})

	r.Route("/roles", func(r chi.Router) {
		//r.Use(middlewarejwt.ValidateToken)
		//r.Use(middlewarejwt.Rolesv)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			rolecontrollers.Listroles(w, r, db)
		})
		r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
			rolecontrollers.Addroles(w, r, db)
		})
		r.Put("/update", func(w http.ResponseWriter, r *http.Request) {
			rolecontrollers.UpdateRoles(w, r, db)
		})
		r.Delete("/delete", func(w http.ResponseWriter, r *http.Request) {
			rolecontrollers.DeleteRoles(w, r, db)
		})
	})

	r.Route("/teams", func(r chi.Router) {
		//r.Use(middlewarejwt.ValidateToken)
		//r.Use(middlewarejwt.Rolesv)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			teamcontroller.Listteams(w, r, db)
		})
		r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
			teamcontroller.Addteams(w, r, db)
		})
		r.Put("/update", func(w http.ResponseWriter, r *http.Request) {
			teamcontroller.Updateteams(w, r, db)
		})
		r.Delete("/delete", func(w http.ResponseWriter, r *http.Request) {
			teamcontroller.Deleteteams(w, r, db)
		})
	})

	r.Route("/companys", func(r chi.Router) {
		//r.Use(middlewarejwt.ValidateToken)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			companycontroller.ListCompany(w, r, db)
		})
		r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
			companycontroller.AddCompany(w, r, db)
		})
		r.Put("/update", func(w http.ResponseWriter, r *http.Request) {
			companycontroller.UpdateCompany(w, r, db)
		})
		r.Delete("/delete", func(w http.ResponseWriter, r *http.Request) {
			companycontroller.DeleteCompany(w, r, db)
		})
	})

	r.Route("/services", func(r chi.Router) {
		//r.Use(middlewarejwt.ValidateToken)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			servicecontroller.Listservice(w, r, db)
		})

		r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
			servicecontroller.Addservice(w, r, db)
		})
		r.Put("/update", func(w http.ResponseWriter, r *http.Request) {
			servicecontroller.Updateservice(w, r, db)
		})
		r.Delete("/delete", func(w http.ResponseWriter, r *http.Request) {
			servicecontroller.Deleteservice(w, r, db)
		})
	})

	r.Route("/servicetypes", func(r chi.Router) {
		//r.Use(middlewarejwt.ValidateToken)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			servicecontroller.Listservicetype(w, r, db)
		})
		r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
			servicecontroller.Addservicetype(w, r, db)
		})
		r.Put("/update", func(w http.ResponseWriter, r *http.Request) {
			servicecontroller.Updateservicetype(w, r, db)
		})
		r.Delete("/delete", func(w http.ResponseWriter, r *http.Request) {
			servicecontroller.Deleteservicetype(w, r, db)
		})
	})

	r.Route("/task", func(r chi.Router) {
		//r.Use(middlewarejwt.ValidateToken)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			taskController.ListTask(w, r, db)
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			taskController.GetTask(w, r, db)
		})
		r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
			taskController.CreateTask(w, r, db)
		})
		r.Put("/update", func(w http.ResponseWriter, r *http.Request) {
			taskController.UpdateTask(w, r, db)
		})
		r.Delete("/delete", func(w http.ResponseWriter, r *http.Request) {
			taskController.DeleteTask(w, r, db)
		})
	})

	r.Route("/subtask", func(r chi.Router) {
		//r.Use(middlewarejwt.ValidateToken)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			taskController.ListSubtask(w, r, db)
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			taskController.GetSubtask(w, r, db)
		})

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			taskController.CreateSubtask(w, r, db)
		})
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			taskController.UpdateSubtask(w, r, db)
		})
		r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
			taskController.DeleteSubtask(w, r, db)
		})
	})

	r.Route("/taskdue", func(r chi.Router) {
		//r.Use(middlewarejwt.ValidateToken)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			taskController.ListTaskdue(w, r, db)
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			taskController.CreateTaskdue(w, r, db)
		})
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			taskController.UpdateTaskdue(w, r, db)
		})
		r.Delete("/", func(w http.ResponseWriter, r *http.Request) {
			taskController.DeleteTaskdue(w, r, db)
		})
	})

	r.Route("/tag", func(r chi.Router) {
		//r.Use(middlewarejwt.ValidateToken)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			tagController.ListTag(w, r, db)
		})
		r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
			tagController.CreateTag(w, r, db)
		})
		r.Put("/update", func(w http.ResponseWriter, r *http.Request) {
			tagController.UpdateTag(w, r, db)
		})
		r.Delete("/delete", func(w http.ResponseWriter, r *http.Request) {
			tagController.DeleteTag(w, r, db)
		})
	})

	r.Route("/chat", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			// List all chats
			chatcontroller.ListChat_Team(w, r, db)
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			// Create a new chat
			chatcontroller.CreateChat_Team(w, r, db)
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			// Get a specific chat by ID
			chatcontroller.ChatByID(w, r, db)
		})
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			// Update a specific chat by ID
			chatcontroller.UpdateChat_Team(w, r, db)
		})
		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			// Delete a specific chat by ID
			chatcontroller.DeleteChat_Team(w, r, db)
		})
	})

	r.Route("/chattask", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			// List all chat tasks
			chatcontroller.ListChat_Task(w, r, db)
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			// Create a new chat task
			chatcontroller.CreateChat_Task(w, r, db)
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			// Get a specific chat task by ID
			chatcontroller.Chat_TaskByID(w, r, db)
		})
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			// Update a specific chat task by ID
			chatcontroller.UpdateChat_Task(w, r, db)

		})
		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			// Delete a specific chat task by ID
			chatcontroller.DeleteChat_Task(w, r, db)
		})
	})

	r.Route("/tax_30", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			taxcontroller.ListTax30(w, r, db)
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			taxcontroller.CreateTax30(w, r, db)
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			taxcontroller.GetTax30ById(w, r, db)
		})
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			taxcontroller.UpdateTax30(w, r, db)
		})
		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			taxcontroller.DeleteTax30(w, r, db)
		})
	})

	r.Route("/tax_from", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			taxcontroller.ListTax(w, r, db)
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			taxcontroller.CreateTax(w, r, db)
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			taxcontroller.GetTaxById(w, r, db)
		})
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			taxcontroller.UpdateTax(w, r, db)
		})
		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			taxcontroller.DeleteTax(w, r, db)
		})
	})

	r.Route("/billing", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			// List all billing records
			billingcontroller.ListBilling(w, r, db)
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			// Create a new billing record
			billingcontroller.CreateBilling(w, r, db)
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			// Get a specific billing record by ID
			billingcontroller.GetBillingById(w, r, db)
		})
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			// Update a specific billing record by ID
			billingcontroller.UpdateBilling(w, r, db)
		})
		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			// Delete a specific billing record by ID
			billingcontroller.DeleteBilling(w, r, db)
		})
	})

	r.Route("/billing_tax", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			// List all billing_tax records
			billingcontroller.ListBilling_tax(w, r, db)
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			// Create a new billing_tax record
			billingcontroller.CreateBilling_tax(w, r, db)
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			// Get a specific billing_tax record by ID
			billingcontroller.GetBilling_taxById(w, r, db)
		})
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			// Update a specific billing_tax record by ID
			billingcontroller.UpdateBilling_tax(w, r, db)
		})
		r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
			// Delete a specific billing_tax record by ID
			billingcontroller.DeleteBilling_tax(w, r, db)
		})
	})

	http.ListenAndServe(":8000", r)
}
