package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/starwars/routers"
)

/*Manejadores se asigna el puerto el handler, y se sube el servidor web*/
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/", routers.Index).Methods("GET")
	router.HandleFunc("/topsecret", routers.Topsecret).Methods("POST")
	router.HandleFunc("/topsecret_split/{satellite_name:[a-z]+}", routers.TopsecretSplit).Methods("POST", "GET")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "80"
	}
	fmt.Println("Servidor iniciado")
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
