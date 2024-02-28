package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/plaja-app/back-end/config"
	c "github.com/plaja-app/back-end/controllers"
	m "github.com/plaja-app/back-end/middleware"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/api/v1/course-categories", c.Controller.GetCourseCategory)
	r.Get("/api/v1/courses", c.Controller.GetCourse)

	r.Post("/api/v1/users/signup", c.Controller.SignUp)
	r.Post("/api/v1/users/login", c.Controller.Login)

	r.Route("/api/v1/users", func(r chi.Router) {
		r.Use(m.Middleware.RequireAuth)
		r.Get("/getme", c.Controller.GetMe)
	})

	return r
}
