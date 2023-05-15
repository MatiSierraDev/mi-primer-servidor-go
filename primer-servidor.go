package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matiraw/mi-primer-servidor-go/db"
	"github.com/matiraw/mi-primer-servidor-go/models"
	"github.com/matiraw/mi-primer-servidor-go/routes"
)

func main() {

	db.DBConnection()

	//creando el modelo
	db.DB.AutoMigrate(&models.User{}, &models.Task{})

	//instacia para crear una nueva ruta
	r := mux.NewRouter()

	//definicion y manejo de rutas
	//users
	r.HandleFunc("/air", routes.HomeHandler)
	r.HandleFunc("/api/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/api/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/api/users/{id}", routes.GetUserhandler).Methods("GET")
	r.HandleFunc("/api/users/{id}", routes.DeleteUserById).Methods("DELETE")
	r.HandleFunc("/api/users/{id}", routes.UpdateUserHandler).Methods("PUT")

	//routes task for users
	r.HandleFunc("/api/tasks", routes.PostTaskHandler).Methods("POST")
	r.HandleFunc("/api/tasks", routes.GetAllTasksHandler).Methods("GET")
	r.HandleFunc("/api/tasks/{id}", routes.GetTaskhandler).Methods("GET")
	r.HandleFunc("/api/tasks/{id}", routes.DeleteTaskHandler).Methods("DELETE")
	r.HandleFunc("/api/tasks/{id}", routes.UpdateTaskHandler).Methods("PUT")
	//levantando el servidor
	fmt.Println("Servidor en puerto: 3000")
	http.ListenAndServe(":3000", r)
}
