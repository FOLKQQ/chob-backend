// main.go

package main

import (
	authcontrollers "backend/controllers/authController"
	companycontroller "backend/controllers/companyController"
	rolecontrollers "backend/controllers/roleController"
	teamcontroller "backend/controllers/teamController"
	"backend/database"
	middlewarejwt "backend/middleware"
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

const Role = ""

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
		r.Use(middlewarejwt.ValidateToken)
		r.Use(middlewarejwt.Rolesv)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			authcontrollers.ListAdmin(w, r, db)
		})
		r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
			authcontrollers.AddAdmin(w, r, db)
		})
		r.Post("/update", func(w http.ResponseWriter, r *http.Request) {
			authcontrollers.UpdateAdmin(w, r, db)
		})
	})

	r.Route("/roles", func(r chi.Router) {
		r.Use(middlewarejwt.ValidateToken)
		r.Use(middlewarejwt.Rolesv)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			rolecontrollers.Listroles(w, r, db)
		})
		r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
			rolecontrollers.Addroles(w, r, db)
		})
		r.Post("/update", func(w http.ResponseWriter, r *http.Request) {
			rolecontrollers.UpdateRoles(w, r, db)
		})
		r.Post("/delete", func(w http.ResponseWriter, r *http.Request) {
			rolecontrollers.DeleteRoles(w, r, db)
		})
	})

	r.Route("/teams", func(r chi.Router) {
		r.Use(middlewarejwt.ValidateToken)
		r.Use(middlewarejwt.Rolesv)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			teamcontroller.Listteams(w, r, db)
		})
		r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
			teamcontroller.Addteams(w, r, db)
		})
		r.Post("/update", func(w http.ResponseWriter, r *http.Request) {
			teamcontroller.Updateteams(w, r, db)
		})
		r.Post("/delete", func(w http.ResponseWriter, r *http.Request) {
			teamcontroller.Deleteteams(w, r, db)
		})
	})

	r.Route("/companys", func(r chi.Router) {
		r.Use(middlewarejwt.ValidateToken)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			companycontroller.ListCompany(w, r, db)
		})
		r.Post("/add", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World!"))
		})
		r.Post("/update", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World!"))
		})
		r.Post("/delete", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World!"))
		})
	})

	http.ListenAndServe(":8000", r)
}
