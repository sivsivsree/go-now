package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Todos []Todo `gorm:"foreignkey:UserRefer"`
}

// Todo is..
type Todo struct {
	gorm.Model
	Name      string
	Token     string
	Completed bool
	Due       time.Time
	UserRefer uint
}

var db = getDB()

func getDB() *gorm.DB {

	db, err := gorm.Open("sqlite3", "database.db")
	if err != nil {
		panic("failed to connect database")
	}

	db.LogMode(true)
	// Migrate the schema
	db.AutoMigrate(&User{}, &Todo{})

	return db

}

func createTodo(todo *Todo) *gorm.DB {
	var user User
	db.Where("Name = ?", "siv").First(&user)
	todo.UserRefer = user.ID
	fmt.Println(user)
	fmt.Println(todo)
	db.Create(&todo)
	return db
}

func getAll() []Todo {
	var todos []Todo
	db.Find(&todos)
	return todos
}

func getAllUsers() []User {
	var users []User
	db.Find(&users)
	fmt.Println(users)
	return users
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
	_ = json.NewEncoder(w).Encode(todos)

}

func createHandler(w http.ResponseWriter, r *http.Request) {
	todo := Todo{Name: "Siv", Completed: false, Due: time.Now(), Token: strconv.FormatInt(time.Now().UnixNano(), 10)}

	createTodo(&todo)
	_ = json.NewEncoder(w).Encode(todo)
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

	r := mux.NewRouter()
	r.Use(simpleMw)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		db.Create(&User{Name: "siv"})
		var user User
		db.Where(&User{Name: "siv"}).First(&user)
		_ = json.NewEncoder(w).Encode(user)
	})
	r.HandleFunc("/list", handler)
	r.HandleFunc("/users", listUsers)
	r.HandleFunc("/create", createHandler)
	r.HandleFunc("/find/{id:[0-9]+}", findHandler)

	log.Fatal(http.ListenAndServe(":9876", r))

	defer db.Close()

}

func findHandler(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	var users User
	//var todo Todo

	// db.Model(&users).Related(&todo, "Todos")
	db.Preload("Todos").First(&users)
	fmt.Println(id)

	_ = json.NewEncoder(w).Encode(users)

}

func listUsers(writer http.ResponseWriter, request *http.Request) {
	users := getAllUsers()
	_ = json.NewEncoder(writer).Encode(users)
}
