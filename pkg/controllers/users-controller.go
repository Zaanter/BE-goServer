package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Zaanter/be-goserver/pkg/errors"

	"github.com/Zaanter/be-goserver/pkg/entities"
	"github.com/Zaanter/be-goserver/pkg/repository"
	"github.com/Zaanter/be-goserver/pkg/services"
	"github.com/gorilla/mux"
)

var (
	repo    repository.UserRepository = repository.NewUserRepository()
	service services.UserService      = services.NewUserService()
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	// utils.ParseBody(r, newUser)
	var newUser entities.IUser
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "failed decoding body"})
		return
	}

	validateError := service.Validate(&newUser)

	if validateError != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: validateError.Error()})
		return
	}

	result, serviceErr := service.Create(&newUser)
	if serviceErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "failed creating new user"})
		return
	}

	res, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("Error applying json marshal to new user. err: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "failed applying json marshal to new user"})
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := repo.FindAll()
	if err != nil {
		log.Fatalf("Error getting all users %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "failed getting all users"`))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid := vars["uid"]
	user, err := repo.GetUser(uid)

	if err != nil {
		log.Fatalf("Error getting user %v", err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "Error getting user"`))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(user)
	w.Write(res)
}

// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Update User")
// 	w.Write([]byte("Update User"))

// }

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid := vars["uid"]

	err := service.Delete(uid)
	if err != nil {
		log.Fatalf("Error deleting user with uid: %v - err: %v", uid, err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "error deleting user"`))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entities.IResponse{Message: "user deleted successfully", Timestamp: time.Now()})

}
