package main

import (
	"fmt"
	"html/template"
	"path/filepath"

	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("processing template: %v", err)
		http.Error(w, "There was an error processing the template.", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  tplPath := filepath.Join("templates", "home.gohtml")
  executeTemplate(w, tplPath)
}


func contactHandler(w http.ResponseWriter, r *http.Request) {
  tplPath := filepath.Join("templates", "home.gohtml")
  executeTemplate(w, tplPath)
  fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in contact, go to  <a href=\"https://www.linkedin.com/in/davinamedina/\">Davina's linkedin</a>.</p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
  tplPath := filepath.Join("templates", "home.gohtml")
  executeTemplate(w, tplPath)
  
  fmt.Fprint(w, `<h1>FAQ Page</h1>
<ul>
  <li>
    <b>Is there a free version?</b>
    Yes! We offer a free trial for 30 days on any paid plans.
  </li>
  <li>
    <b>What are your support hours?</b>
    We have support staff answering emails 24/7, though response
    times may be a bit slower on weekends.
  </li>
  <li>
    <b>How do I contact support?</b>
    Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
  </li>
</ul>
`)
}

func main() {
  r := chi.NewRouter()
  r.Get("/", homeHandler)
  r.Get("/contact", contactHandler)
  r.Get("/faq", faqHandler)

  r.NotFound(func(w http.ResponseWriter, r *http.Request) {
    http.Error(w, "Page not found", http.StatusNotFound)
  })
  fmt.Println("Starting the server on :3000...")
  http.ListenAndServe(":3000", r)
}