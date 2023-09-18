// main.go

package main

import (
	"backend/controllers/authController"
	companycontroller "backend/controllers/companyController"
	rolecontrollers "backend/controllers/roleController"
	teamcontroller "backend/controllers/teamController"
	"backend/database"
	"backend/middleware"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(503)
		w.Write([]byte("bad"))
	}
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
