package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/Davina-dev/First-steps-with-Go/lenslocked/controllers"
	template "github.com/Davina-dev/First-steps-with-Go/lenslocked/templates"
	"github.com/Davina-dev/First-steps-with-Go/lenslocked/views"
)



func main() {
  r := chi.NewRouter() 
 
  r.Get("/", controllers.StaticHandler(
	views.Must(views.ParseFS(template.FS, "home.gohtml" ))))
  r.Get("/contact", controllers.StaticHandler(
	views.Must(views.ParseFS(template.FS, "contact.gohtml"))))
  r.Get("/faq", controllers.StaticHandler(
	views.Must(views.ParseFS(template.FS, "faq.gohtml"))))

  r.NotFound(func(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "Page not found", http.StatusNotFound)
  })
  fmt.Println("Starting the server on :3000...")
  http.ListenAndServe(":3000", r)
}