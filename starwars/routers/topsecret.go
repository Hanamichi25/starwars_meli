package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/starwars/modelos"
	"github.com/starwars/servicios"
)

func Topsecret(w http.ResponseWriter, r *http.Request) {
	var satelites modelos.SatelitesInput
	fmt.Println("Body", r.Body)
	err := json.NewDecoder(r.Body).Decode(&satelites)
	fmt.Println("** Satelites ", satelites)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), http.StatusBadGateway)
		return
	}
	outPut, err := servicios.ProcessInput(satelites)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if outPut.X == 0 && outPut.Y == 0 {
		http.Error(w, "No existe suficiente informacion", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(outPut)
}
