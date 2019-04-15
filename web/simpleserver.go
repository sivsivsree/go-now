package web

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Todo is..
type Todo struct {
	Name      string
	Completed bool
	Due       time.Time
}

// Todos is..
type Todos []Todo

func Server() {

	http.HandleFunc("/", rootHandle)

	log.Fatal(http.ListenAndServe(":9876", nil))

}

func rootHandle(w http.ResponseWriter, r *http.Request) {
	todos := Todos{
		Todo{Name: "Write presentation", Due: time.Now()},
		Todo{Name: "Host meetup"},
	}

	json.NewEncoder(w).Encode(todos)

}
