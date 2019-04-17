package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/sivsivsree/go-now/web/database"
	M "github.com/sivsivsree/go-now/web/modals"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	err := M.ResErr{}
	err.ErrorMessage("API not found")
	_ = json.NewEncoder(w).Encode(err)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {

	user, _ := database.CreateUsers(fmt.Sprintf("%v", context.Get(r, "user")))
	_ = json.NewEncoder(w).Encode(user)
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {

	if todo, err := database.CreateTodo(fmt.Sprintf("%v", context.Get(r, "user"))); err == nil {
		_ = json.NewEncoder(w).Encode(todo)
	} else {
		err := M.ResErr{}
		err.ErrorMessage("User Creation failed")
		_ = json.NewEncoder(w).Encode(err)
	}

}

func ListTodo(w http.ResponseWriter, r *http.Request) {

	//token := fmt.Sprintf("%v", context.Get(r, "token"))
	todos := database.GetAll()
	_ = json.NewEncoder(w).Encode(todos)

}

func EchoHandler(w http.ResponseWriter, r *http.Request) {

}

func FindHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	if users, err := database.FindTodoByUserID(id); err == nil {
		_ = json.NewEncoder(w).Encode(users)
	} else {
		err := M.ResErr{}
		err.ErrorMessage("The record not found")
		_ = json.NewEncoder(w).Encode(err)
	}

}

func ListUsers(writer http.ResponseWriter, request *http.Request) {
	users := database.GetAllUsers()
	_ = json.NewEncoder(writer).Encode(users)
}
