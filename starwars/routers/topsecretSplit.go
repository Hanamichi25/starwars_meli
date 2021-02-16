package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/starwars/comun"
	"github.com/starwars/modelos"
	"github.com/starwars/servicios"
)

var listNombresSatelites = []string{comun.SateliteKenobi, comun.SateliteSato, comun.SateliteSkywalker}

func TopsecretSplit(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["satellite_name"]
	index, exist := validarNombreSatelite(name)
	if exist == false {
		http.Error(w, "El satelite no existe en esta universo ", http.StatusBadGateway)
		return
	} else {
		var input modelos.InputSplit
		err := json.NewDecoder(r.Body).Decode(&input)
		fmt.Println("*** input", err)
		if input.Distance > 0 {
			var r1 float32
			var r2 float32
			var r3 float32
			switch index {
			case 0:
				r1 = input.Distance
			case 1:
				r2 = input.Distance
			case 2:
				r3 = input.Distance
			}
			x, y := servicios.GetLocation(r1, r2, r3)
			if x == 0 && y == 0 {
				w.WriteHeader(http.StatusNotFound)
				return
			}
			position := &modelos.Position{
				X: x,
				Y: y,
			}
			message := servicios.GetMessage(input.Message)
			output := &modelos.OutPut{
				Position: *position,
				Message:  message,
			}
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(output)
		} else {
			w.Header().Set("Content-type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			http.Error(w, "El forma no coincide con el requerdio", http.StatusBadRequest)
		}
	}
}

func validarNombreSatelite(nombre string) (int, bool) {
	return contains(listNombresSatelites, nombre)
}

func contains(s []string, str string) (int, bool) {
	for i, v := range s {
		if v == str {
			return i, true
		}
	}
	return 0, false
}
