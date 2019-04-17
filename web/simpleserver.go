package web

import (
	"encoding/json"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Todo is..
type Todo struct {
	gorm.Model
	Name      string
	Token     string
	Completed bool
	Due       time.Time
}

var db = getDB()

func getDB() *gorm.DB {

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}

	db.LogMode(true)
	// Migrate the schema
	db.AutoMigrate(&Todo{})

	return db

}

func createTodo(todo Todo) *gorm.DB {
	// Create
	db.Create(&todo)
	return db
}

func getAll() []Todo {
	var todos []Todo
	db.Find(&todos)

	return todos;
}

/*func setUpDB() {

	db := getDB()
	// Read
	var product Product
	db.First(&product, 1)                   // find product with id 1
	db.First(&product, "code = ?", "L1212") // find product with code l1212

	// Update - update product's price to 2000
	db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	db.Delete(&product)

}*/

func handler(w http.ResponseWriter, r *http.Request) {

	// token := fmt.Sprintf("%v", context.Get(r, "token"))
	todos := getAll()
	json.NewEncoder(w).Encode(todos)

}

func createHandler(w http.ResponseWriter, r *http.Request) {
	todo := Todo{Name: "Siv", Completed: false, Due: time.Now(), Token: strconv.FormatInt(time.Now().UnixNano(), 10)}
	createTodo(todo)
	json.NewEncoder(w).Encode(todo)
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

func Server() {

	// db.DropTable(&Todo{})

	r := mux.NewRouter()
	r.Use(simpleMw)
	r.HandleFunc("/list", handler)
	r.HandleFunc("/create", createHandler)

	log.Fatal(http.ListenAndServe(":9876", r))

	defer db.Close()

}
