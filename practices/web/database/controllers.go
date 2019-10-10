package database

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/sivsivsree/go-now/practices"
	"log"
	"os"
	"strconv"
	"time"
)

var db = GetDB()

// defer db.Close()

func GetDB() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// This is public sql database anyone can take a peek on it. if wanted.
	dsn := fmt.Sprintf(
		"%s:%s@tcp(sql12.freemysqlhosting.net:3306)/%s?charset=utf8&parseTime=True",
		os.Getenv("USERNAME"),
		os.Getenv("PASSWORD"),
		os.Getenv("DATABASE"),
	)

	log.Println(dsn)
	//db, err := gorm.Open("sqlite3", "database.db")
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	// Migrate the schema
	db.AutoMigrate(&practices.User{}, &practices.Todo{})

	return db

}

func CreateUsers(name string) (practices.User, error) {

	db.Create(&practices.User{Name: name})
	var user practices.User
	db.Where(&practices.User{Name: name}).First(&user)

	return user, nil
}

func CreateTodo(name string) (practices.Todo, error) {
	todo := practices.Todo{Name: "avooo", Completed: false, Due: time.Now(), Token: strconv.FormatInt(time.Now().UnixNano(), 10)}
	var user practices.User
	if noUsers := db.Where("Name = ?", name).First(&user).RecordNotFound(); !noUsers {
		todo.UserRefer = user.ID
		fmt.Println(user)
		fmt.Println(todo)
		db.Create(&todo)
		return todo, nil
	} else {
		return todo, errors.New("Todo Creation Failed")
	}

}

func GetAll() []practices.Todo {

	var todos []practices.Todo
	db.Find(&todos)
	return todos
}

func GetAllUsers() []practices.User {

	var users []practices.User
	db.Find(&users)
	fmt.Println(users)
	return users
}

func FindTodoByUserID(id string) (practices.User, error) {
	var users practices.User
	//var to do To do

	// db.Model(&users).Related(&to do, "Todos")
	if noUsers := db.Preload("Todos").Where("id = ?", id).First(&users).RecordNotFound(); !noUsers {

		return users, nil

	} else {
		return practices.User{}, errors.New("No Records found")
	}
}
