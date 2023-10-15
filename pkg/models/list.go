package models

import (
	"time"

	"github.com/maskedemann/go-todo/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Task struct {
	gorm.Model
	Name         string     `json:"name"`
	CreationDate *time.Time `json:"creationDate"`
	Deadline     *time.Time `json:"deadline"`
	Status       *bool      `json:"status"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Task{})
}

func GetAll() []Task {
	var Tasks []Task
	db.Find(&Tasks)
	return Tasks
}

func GetTaskById(Id int64) (*Task, *gorm.DB) {
	var getTask Task
	db := db.Where("ID=?", Id).Find(&getTask)
	return &getTask, db

}

func (b *Task) CreateTask() *Task {
	db.Create(&b)
	return b
}

func DeleteTask(Id int64) Task {
	var task Task
	db.Where("ID=?", Id).Unscoped().Delete(task)
	return task
}
