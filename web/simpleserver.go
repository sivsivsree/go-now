package web

import (
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sivsivsree/go-now/web/modals"
	"os"

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
	r.HandleFunc("/echo", handlers.EchoHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST", "GET")

	log.Fatal(http.ListenAndServe(os.Getenv("PORT"), r))

}

func simpleMw(next http.Handler) http.Handler {
	allowedApi := map[string]bool{
		"/login": true,
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		log.Println(r.RemoteAddr, r.RequestURI, r.Method)
		if allowedApi[r.RequestURI] == true {
			next.ServeHTTP(w, r)
			return
		}

		token := r.Header.Get("Authorization")
		if token == "" {
			token = r.FormValue("token")
		}

		if token == "" {
			err := modals.ResErr{}
			err.ErrorMessage("No token provided")
			w.WriteHeader(401)
			_ = json.NewEncoder(w).Encode(err)
			return
		}

		parsedToken, err := handlers.VerifyJWT(token)

		if err != nil {
			log.Println(err)
			resErr := modals.ResErr{}
			resErr.ErrorMessage(err.Error())
			w.WriteHeader(401)
			_ = json.NewEncoder(w).Encode(resErr)
			return
		}

		log.Println("token", parsedToken.Claims)
		context.Set(r, "token", parsedToken.Claims)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
