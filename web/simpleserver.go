package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// Todo is..
type Todo struct {
	Name      string
	Token     string
	Completed bool
	Due       time.Time
}

// Todos is..
type Todos []Todo

func Server() {

	r := mux.NewRouter()
	r.HandleFunc("/test", handler)
	r.Use(simpleMw)

	log.Fatal(http.ListenAndServe(":9876", r))

}

func handler(w http.ResponseWriter, r *http.Request) {

	token := fmt.Sprintf("%v", context.Get(r, "token"))

	todos := Todos{
		Todo{Name: "Write presentation", Due: time.Now(), Token: token},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)

}

func simpleMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RemoteAddr, r.RequestURI, r.Method)
		context.Set(r, "token", "23412342")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
