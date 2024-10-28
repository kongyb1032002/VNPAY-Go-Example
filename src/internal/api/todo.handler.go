package api

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var todos []Todo

type TodoHandler interface {
	POST(rw http.ResponseWriter, r *http.Request)
	PUT(rw http.ResponseWriter, r *http.Request, id int)
	GET(rw http.ResponseWriter, r *http.Request, id int)
	DELETE(rw http.ResponseWriter, r *http.Request, id int)
}

type todoHandler struct {
}

func NewTodoHandler() TodoHandler {
	return &todoHandler{}
}

func (h *todoHandler) POST(rw http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)
	todos = append(todos, todo)
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(todo)
}
func (h *todoHandler) PUT(rw http.ResponseWriter, r *http.Request, id int) {
	for index, todo := range todos {
		if todo.ID == id {
			json.NewDecoder(r.Body).Decode(&todo)
			todos[index].ID = todo.ID
			todos[index].Task = todo.Task
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte(`{"message": "Success to update todo"}`))
		}
	}
}
func (h *todoHandler) GET(rw http.ResponseWriter, r *http.Request, id int) {
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(todos)
}
func (h *todoHandler) DELETE(rw http.ResponseWriter, r *http.Request, id int) {
	for index, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte(`{"message": "Success to delete todo"}`))
		}
	}
}
