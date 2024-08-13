package main

import (
	"log"
	"net/http"
)


type Todo struct{
    ID          int
    Content     string 
}

var todos = []Todo{
    {ID:1 ,Content: "Hello world"},
    {ID:2, Content: "Say something"},
    {ID:3, Content: "Do something"},
}





func main() {
    h := NewHandler()
    mux := http.NewServeMux()
    mux.HandleFunc("/", h.Index)
    mux.HandleFunc("/add-todo", h.AddTodo)
    log.Println("Let's go to 8080 ")
    http.ListenAndServe(":8080", mux)
}
