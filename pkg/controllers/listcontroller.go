package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/maskedemann/go-todo/pkg/models"
	"github.com/maskedemann/go-todo/pkg/utils"
)

var NewTask models.Task

// func multiplexer(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		GetAll(w, r)
// 	case http.MethodPost:
// 		CreateTask(w, r)
// 	case http.MethodPut:
// 		UpdateTask(w, r)
// 	// case "GET":
// 	// 	GetTaskById(w,r)
// 	case http.MethodDelete:
// 		DeleteTask(w, r)
// 	}
// }

func GetAll(w http.ResponseWriter, r *http.Request) {

	newTasks := models.GetAll()
	res, _ := json.Marshal(newTasks)
	w.Header().Set("Content-Type", "app/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTaskById(w http.ResponseWriter, r *http.Request) {
	println("-------------------------")

	vars := mux.Vars(r)
	taskId := vars["taskId"]
	fmt.Printf("taskId: %v\n", taskId)
	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	taskDetails, _ := models.GetTaskById(ID)
	res, _ := json.Marshal(taskDetails)
	w.Header().Set("Content-Type", "app/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {

	CreateTask := &models.Task{}
	utils.ParseBody(r, CreateTask)
	l := CreateTask.CreateTask()
	res, _ := json.Marshal(l)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateTask(w http.ResponseWriter, r *http.Request) {

	UpdateList := &models.Task{}
	utils.ParseBody(r, UpdateList)
	vars := mux.Vars(r)
	taskId := vars["taskId"]

	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	listDetails, db := models.GetTaskById(ID)
	if UpdateList.Name != "" {
		listDetails.Name = UpdateList.Name
	}
	if UpdateList.CreationDate != nil {
		listDetails.CreationDate = UpdateList.CreationDate
	}
	if UpdateList.Deadline != nil {
		listDetails.Deadline = UpdateList.Deadline
	}
	if UpdateList.Status != nil {
		listDetails.Status = UpdateList.Status
	}
	db.Save(&listDetails)
	res, _ := json.Marshal(listDetails)
	w.Header().Set("Content-Type", "app/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["taskId"]

	ID, err := strconv.ParseInt(taskId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	task := models.DeleteTask(ID)
	res, _ := json.Marshal(task)
	w.Header().Set("Content-Type", "app/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
