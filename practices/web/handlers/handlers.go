package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/sivsivsree/go-now/practices"
	"net/http"
	"time"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	err := practices.ResErr{}
	err.ErrorMessage("API not found")
	_ = json.NewEncoder(w).Encode(err)
}
func CreateUser(w http.ResponseWriter, r *http.Request) {

	user := context.Get(r, "token").(jwt.MapClaims)
	if userCreated, err := practices.CreateUsers(fmt.Sprintf("%v", user["UserName"])); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(userCreated)
	} else {
		err := practices.ResErr{}
		err.ErrorMessage("User Creation failed")
		_ = json.NewEncoder(w).Encode(err)
	}

}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	user := context.Get(r, "token").(jwt.MapClaims)
	if todo, err := practices.CreateTodo(fmt.Sprintf("%v", user["UserName"])); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(todo)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		errRes := practices.ResErr{}
		errRes.ErrorMessage("User Creation failed")
		_ = json.NewEncoder(w).Encode(err)
	}

}

func ListTodo(w http.ResponseWriter, r *http.Request) {

	//token := fmt.Sprintf("%v", context.Get(r, "token"))
	todos := practices.GetAll()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(todos)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	user := r.FormValue("user")
	if user == "" {
		err := practices.ResErr{}
		err.ErrorMessage("No user..")
		_ = json.NewEncoder(w).Encode(err)
		return
	}

	tk := practices.Token{
		UserId:   1,
		UserName: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 60).Unix(),
			Issuer:    "Siv",
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := practices.CreateJWT(&tk)
	res := map[string]string{
		"succes": "true",
		"token":  token,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}

func EchoHandler(w http.ResponseWriter, r *http.Request) {

	user := context.Get(r, "token").(jwt.MapClaims)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

func FindHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	if users, err := practices.FindTodoByUserID(id); err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_ = json.NewEncoder(w).Encode(users)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err := practices.ResErr{}
		err.ErrorMessage("The record not found")
		_ = json.NewEncoder(w).Encode(err)
	}

}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	users := practices.GetAllUsers()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(users)
}
