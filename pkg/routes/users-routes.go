package routes

import (
	"github.com/Zaanter/be-goserver/pkg/controllers"
	"github.com/gorilla/mux"
)

// Defino las rutas que van a manejar a los users

func InitializeUsersRoutes(r *mux.Router) {
	r.HandleFunc("/user", controllers.CreateUser).Methods("POST")
	r.HandleFunc("/user", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/user/{uid}", controllers.GetUser).Methods("GET")
	// r.HandleFunc("/user/{uid}", controllers.UpdateUser).Methods("POST")
	r.HandleFunc("/user/{uid}", controllers.DeleteUser).Methods("DELETE")
}
