package web

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/sivsivsree/go-now/web/handlers"
	"log"
	"net/http"
)

func Server() {

	r := mux.NewRouter()
	r.Use(simpleMw)
	r.NotFoundHandler = http.HandlerFunc(handlers.ErrorHandler)

	r.HandleFunc("/", handlers.CreateUser)
	r.HandleFunc("/create", handlers.CreateTodoHandler)
	r.HandleFunc("/list", handlers.ListTodo)
	r.HandleFunc("/users", handlers.ListUsers)
	r.HandleFunc("/find/{id:[0-9]+}", handlers.FindHandler)
	r.HandleFunc("/echo", handlers.EchoHandler)

	log.Fatal(http.ListenAndServe(":9876", r))

}

func simpleMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RemoteAddr, r.RequestURI, r.Method)
		context.Set(r, "user", "akhil")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
