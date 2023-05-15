package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matiraw/mi-primer-servidor-go/db"
	"github.com/matiraw/mi-primer-servidor-go/models"
	"github.com/matiraw/mi-primer-servidor-go/utils"
)

func PostTaskHandler(w http.ResponseWriter, r *http.Request) {

	var task models.Task

	json.NewDecoder(r.Body).Decode(&task)

	createdTask := db.DB.Create(&task)
	err := createdTask.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&task)

}

func GetAllTasksHandler(w http.ResponseWriter, r *http.Request) {

	var tasks []models.Task

	db.DB.Find(&tasks)

	json.NewEncoder(w).Encode(&tasks)

}

func GetTaskhandler(w http.ResponseWriter, r *http.Request) {

	taskId := mux.Vars(r)["id"]
	i := utils.Options{IdTask: &taskId}
	if !utils.CheckId(i) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No existe el usuario con el ID proporcionado")
		return
	}

	var task models.Task
	db.DB.First(&task, taskId)

	json.NewEncoder(w).Encode(&task)

}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	taskId := mux.Vars(r)["id"]

	i := utils.Options{IdTask: &taskId}
	if !utils.CheckId(i) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No existe la tarea con el ID proporcionado")
		return
	}

	var task models.Task

	db.DB.Delete(&task, taskId)
	json.NewEncoder(w).Encode("Sucessfull")

}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {

	taskId := mux.Vars(r)["id"]

	i := utils.Options{IdTask: &taskId}
	if !utils.CheckId(i) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No existe la tarea con el ID proporcionado")
		return
	}

	var task models.Task
	db.DB.First(&task, taskId)
	json.NewDecoder(r.Body).Decode(&task)
	db.DB.Save(&task)

	json.NewEncoder(w).Encode("Tarea actulizada correctamente")
}
