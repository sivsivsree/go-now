package web

import (
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/sivsivsree/go-now/practices"
	"os"

	"log"
	"net/http"
)

func Server() {

	r := mux.NewRouter()
	r.Use(simpleMw)
	r.NotFoundHandler = http.HandlerFunc(practices.ErrorHandler)

	r.HandleFunc("/", practices.CreateUser)
	r.HandleFunc("/create", practices.CreateTodoHandler)
	r.HandleFunc("/list", practices.ListTodo)
	r.HandleFunc("/users", practices.ListUsers)
	r.HandleFunc("/find/{id:[0-9]+}", practices.FindHandler)
	r.HandleFunc("/echo", practices.EchoHandler).Methods("POST")
	r.HandleFunc("/login", practices.LoginHandler).Methods("POST", "GET")

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
			err := practices.ResErr{}
			err.ErrorMessage("No token provided")
			w.WriteHeader(401)
			_ = json.NewEncoder(w).Encode(err)
			return
		}

		parsedToken, err := practices.VerifyJWT(token)

		if err != nil {
			log.Println(err)
			resErr := practices.ResErr{}
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
