package server

import (
	"net/http"
	"text/template"
)

func (s *Server) makeHandler() {
	mux := s.Mux

	mux.HandleFunc("/", s.homeHandler)
	mux.HandleFunc("/about", s.aboutHandler)
}

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	t, err := template.ParseFiles("static/template/index.tmpl")
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    t.Execute(w, nil)
}

var templatesTest = template.Must(
	template.ParseFiles(
		"static/template/base.tmpl", 
		"static/template/about.tmpl",
	))


func (s *Server) aboutHandler(w http.ResponseWriter, r *http.Request) {
	err := templatesTest.ExecuteTemplate(w, "base", nil)
	// t, err := template.ParseFiles("static/template/about.tmpl")
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}