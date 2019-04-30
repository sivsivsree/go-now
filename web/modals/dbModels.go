package modals

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name  string `gorm:"unique"`
	Todos []Todo `gorm:"foreignkey:UserRefer"`
}

// Todo is ..
type Todo struct {
	gorm.Model
	Name      string
	Token     string
	Completed bool
	Due       time.Time
	UserRefer uint
}
