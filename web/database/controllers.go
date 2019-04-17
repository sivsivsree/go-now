package database

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	M "github.com/sivsivsree/go-now/web/modals"
	"log"
	"strconv"
	"time"
)

var db = GetDB()

// defer db.Close()

func GetDB() *gorm.DB {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(sql12.freemysqlhosting.net:3306)/%s?charset=utf8&parseTime=True",
		"sql12288441",
		"3mHKHz3C8g",
		"sql12288441",
	)

	log.Println(dsn)
	//db, err := gorm.Open("sqlite3", "database.db")
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	// Migrate the schema
	db.AutoMigrate(&M.User{}, &M.Todo{})

	return db

}

func CreateUsers(name string) (M.User, error) {

	db.Create(&M.User{Name: name})
	var user M.User
	db.Where(&M.User{Name: name}).First(&user)

	return user, nil
}

func CreateTodo(name string) (M.Todo, error) {
	todo := M.Todo{Name: "avooo", Completed: false, Due: time.Now(), Token: strconv.FormatInt(time.Now().UnixNano(), 10)}
	var user M.User
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

func GetAll() []M.Todo {

	var todos []M.Todo
	db.Find(&todos)
	return todos
}

func GetAllUsers() []M.User {

	var users []M.User
	db.Find(&users)
	fmt.Println(users)
	return users
}

func FindTodoByUserID(id string) (M.User, error) {
	var users M.User
	//var todo Todo

	// db.Model(&users).Related(&todo, "Todos")
	if noUsers := db.Preload("Todos").Where("id = ?", id).First(&users).RecordNotFound(); !noUsers {

		return users, nil

	} else {
		return M.User{}, errors.New("No Records found")
	}
}
