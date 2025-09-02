package handlers

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	root, _ := os.Getwd()
	path := filepath.Join(root, "templates", "index.html")
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(w, "Could not parse template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]string{
		"Title":   "Go Web App",
		"Message": "Welcome to my Go web application!",
	}
	tmpl.Execute(w, data)
}