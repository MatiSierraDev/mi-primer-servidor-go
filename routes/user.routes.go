package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matiraw/mi-primer-servidor-go/db"
	"github.com/matiraw/mi-primer-servidor-go/models"
	"github.com/matiraw/mi-primer-servidor-go/utils"
)

func PostUserHandler(w http.ResponseWriter, r *http.Request) {

	var user models.User

	json.NewDecoder(r.Body).Decode(&user)

	createdUser := db.DB.Create(&user)
	err := createdUser.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&user)

}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {

	var user []models.User

	db.DB.Find(&user)

	json.NewEncoder(w).Encode(&user)

}

func GetUserhandler(w http.ResponseWriter, r *http.Request) {

	// var user models.User

	userId := mux.Vars(r)["id"]
	i := utils.Options{IdUser: &userId}
	if !utils.CheckId(i) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No existe el usuario con el ID proporcionado")
		return
	}

	var user models.User
	db.DB.First(&user, userId)

	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)

	json.NewEncoder(w).Encode(&user)

}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {

	userId := mux.Vars(r)["id"]

	if !utils.CheckUserId(userId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No existe el usuario con el ID proporcionado")
		return
	}

	var user models.User
	db.DB.Delete(&user, userId)
	// w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("El usuario ah sido borrado correctamente.")

}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	userId := mux.Vars(r)["id"]

	if !utils.CheckUserId(userId) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No se econtro el usuario proporcionado")
		return
	}

	var user models.User

	db.DB.First(&user, userId)
	json.NewDecoder(r.Body).Decode(&user)
	db.DB.Save(&user)
	json.NewEncoder(w).Encode(&user)

}
