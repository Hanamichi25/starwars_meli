package servicios

import (
	"errors"
	"fmt"

	"github.com/starwars/comun"
	"github.com/starwars/modelos"
)

func ProcessInput(satelites modelos.SatelitesInput) (*modelos.OutPut, error) {
	fmt.Println("He recibido una peticion para estos satelites ", satelites)
	var r1 float32 = 0
	var r2 float32 = 0
	var r3 float32 = 0
	if len(satelites.SatelitesInput) == 0 {
		return nil, errors.New("Es necesario como minimo un satelite")
	}
	var listMessage [][]string
	for _, value := range satelites.SatelitesInput {
		fmt.Println("####", value)
		if value.Name == comun.SateliteKenobi {
			r1 = value.Distance
		} else if value.Name == comun.SateliteSkywalker {
			r2 = value.Distance
		} else if value.Name == comun.SateliteSato {
			r3 = value.Distance
		} else if value.Name == "" {
			return nil, errors.New("El nombre del satelite es necesario")
		} else {
			return nil, errors.New("El satelite " + value.Name + " no se encuentra en esta galaxia ")
		}
		listMessage = append(listMessage, value.Message)
	}
	x, y := GetLocation(r1, r2, r3)
	position := &modelos.Position{
		X: x,
		Y: y,
	}
	message := GetMessage(listMessage[:]...)
	outPut := &modelos.OutPut{
		Position: *position,
		Message:  message,
	}

	fmt.Println("r1 ", r1)
	fmt.Println("r2 ", r2)
	fmt.Println("r3 ", r3)
	return outPut, nil
}
