package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Home!")
}

type PageVariables struct {
	PageTitle string
	PageTodos []Todo
}

type Todo struct {
	Title string
	Content string
}

var todos []Todo

func getTodos(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables {
		PageTitle: "Get Todos",
		PageTodos: todos,
	}
	t, err := template.ParseFiles("./../../todos.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template parsing error: ", err)
	}

	err = t.Execute(w, pageVariables)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template parsing error: ", err)
	}

	todo := Todo {
		Title: r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	todos = append(todos, todo);
	log.Print(todos)
	http.Redirect(w, r, "/todos/", http.StatusSeeOther)
}

func main() {
	port := ":3000"
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", getTodos)
	http.HandleFunc("/add-todo/", addTodo)

	fmt.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
