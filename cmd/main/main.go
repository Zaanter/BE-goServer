package main

import (
	"fmt"
	"log"
	"net/http"

	"context"

	"github.com/Zaanter/be-goserver/pkg/routes"
	"github.com/gorilla/mux"

	firebase "firebase.google.com/go"

	"google.golang.org/api/option"
)

func main() {
	fmt.Println("Server started on port 8080")
	// Creo un router
	http.NewServeMux()
	r := mux.NewRouter()

	opt := option.WithCredentialsFile("../../pkg/config/zcoding.json")
	_, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Printf("error initializing app: %v", err)
	}

	// Inicializamos las rutas
	routes.InitializeBaseRoutes(r)
	routes.InitializeUsersRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
