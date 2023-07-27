package main

import (
	"fmt"
	"net/http"
)
func homeHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html; charset=utf-8")
  fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
  fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in contact, go to  <a href=\"https://www.linkedin.com/in/davinamedina/\">Davina's linkedin</a>.</p>")
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
  switch r.URL.Path {
  case "/":
    homeHandler(w, r)
  case "/contact":
    contactHandler(w, r)
  default:
    http.Error(w,"Page Not Found 🧐", http.StatusNotFound)
  }
}

func main() {
  http.HandleFunc("/",pathHandler)
  fmt.Println("Starting the server on :3000...")
  http.ListenAndServe(":3000", nil)
}