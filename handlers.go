package main

import (
	"html/template"
	"net/http"

)

type Handler struct {
	templ *template.Template

}

func NewHandler() *Handler {
    return &Handler{
        templ: template.Must(template.ParseGlob("templates/*.html")),
    }
}

func (h *Handler) Render(w http.ResponseWriter, name string, data interface{}) error {
    return h.templ.ExecuteTemplate(w, name, data)
} 


func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	h.Render(w, "index", nil)
}


func (h *Handler) AddTodo(w http.ResponseWriter, r *http.Request)  {
    id := len(todos)      
    content := r.FormValue("content")
    task := Todo{
        ID: id,
        Content: content,
    }
    todos = append(todos, task)
    w.WriteHeader(http.StatusOK)
    h.Render(w, "tasks", todos)
}






