package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Necesito recibir un router como pointer y asignarle las rutas

func InitializeBaseRoutes(r *mux.Router) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Estoy en la route base"))
	})
}
