package main

import (
	"net/http"
	"text/template"
)

type Todo struct {
    Text string
    ID int 
}

var Tasks = []Todo{
    {Text: "Do some maths", ID: 1,},
    {Text: "Do a little bit of coding", ID:2},
    {Text: "Sleep for 10 hours on days off", ID: 3,},
    {Text: "something else, do whatever the heck comes to your mind", ID: 4,},

}



func main() {
    mux := http.NewServeMux()
    server := &http.Server{
        Addr: ":8080",
        Handler: mux,
    }

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    mux.HandleFunc("/", HandleIndex)
    server.ListenAndServe()
}




func HandleIndex(w http.ResponseWriter, r *http.Request){
    templ := template.Must(template.ParseFiles("templates/index.templ"))
    templ.Execute(w, map[string]interface{}{
        "Todos": Tasks,
    })
}





func HandleAdd(w http.ResponseWriter, r *http.Request) {
   if r.Method == http.MethodPost {
        r.ParseForm()
        todo := r.FormValue("todo")
        task := Todo{
            Text: todo,
            ID: len(Tasks), 
        }
        Tasks = append(Tasks, task)
   } 
}




func HandleDelte(w http.ResponseWriter, r *http.Request) {
    
}




