package servicios

import (
	"fmt"
	"math"

	"github.com/starwars/modelos"
)

var P1 modelos.Satelite
var P2 modelos.Satelite
var P3 modelos.Satelite

func init() {
	fmt.Println("Inicializando informacion de los satelites")
	P1 = modelos.Satelite{
		Punto: modelos.Punto{
			CoordenadaX: -500,
			CoordenadaY: -200,
		},
		Nombre: "kenobi",
	}
	P2 = modelos.Satelite{
		Punto: modelos.Punto{
			CoordenadaX: 500,
			CoordenadaY: 100,
		},
		Nombre: "sato",
	}
	P3 = modelos.Satelite{
		Punto: modelos.Punto{
			CoordenadaX: 100,
			CoordenadaY: -100,
		},
		Nombre: "skywalker",
	}
}

func GetLocation(distances ...float32) (x1, y1 float32) {
	r1 := float64(distances[0])
	r2 := float64(distances[1])
	r3 := float64(distances[2])
	fmt.Println("r1 ", r1)
	fmt.Println("r2 ", r2)
	fmt.Println("r3 ", r3)
	// Según la documentación de wikipedia
	// https://stackoverflow.com/questions/9747227/2d-trilateration
	// https://en.wikipedia.org/wiki/True-range_multilateration
	// Primero se calcula ex
	// Se calcula la diferencia entre dos puntos
	// ex = (P2 - P1) / ‖P2 - P1‖
	ex := dividirPunto(diferenciaEntreDosPuntos(P2.Punto, P1.Punto), calcularDividendo(P2.Punto, P1.Punto))
	// Se calcula i ex(P3 - P1)
	i := multiplicarDosPuntos(ex, diferenciaEntreDosPuntos(P3.Punto, P1.Punto))
	// Se calcula el ey (P3 - P1 - i · ex) / ‖P3 - P1 - i · ex‖
	// primero multiplicto i * ex
	iex := multiplicarUnPuntoPorUnValor(ex, i)
	// P3 - P1
	delta := diferenciaEntreDosPuntos(P3.Punto, P1.Punto)
	// P3 - P1 - i · ex
	divey := diferenciaEntreDosPuntos(*delta, *iex)
	// ‖P3 - P1 - i · ex‖
	divisorey := obtenerNorma(divey)
	ey := dividirPunto(divey, divisorey)
	// Se calcula d
	d := obtenerNorma(diferenciaEntreDosPuntos(P2.Punto, P1.Punto))
	// se calcula J
	// ey(P3 - P1)
	j := multiplicarDosPuntos(ey, diferenciaEntreDosPuntos(P3.Punto, P1.Punto))
	// Se calcula x (r12 - r22 + d2) / 2d
	// donde r1 es la distancia entre el emisor y el satelite p1
	// donde r2 es la distancia entre el emisor y el satelite p2
	x := ((r1 * r1) - (r2 * r2) + (d * d)) / (2 * d)
	// Se calcula y  (r12 - r32 + i2 + j2) / 2j - ix / j
	y := ((r1*r1)-(r3*r3)+(i*i)+(j*j))/(2*j) - (i/j)*x
	// donde r3 es la distancia entre el emisor y el satelite p3
	// 8. p1,2 = P1 + x*ex + y*ey
	// x*ex
	xex := multiplicarUnPuntoPorUnValor(ex, x)
	// y*ey
	yey := multiplicarUnPuntoPorUnValor(ey, y)
	// x*ex + y*ey
	xexAddYey := sumarDosPuntos(xex, yey)
	// P1 + x*ex + y*ey
	puntoFinal := sumarDosPuntos(&P1.Punto, xexAddYey)
	// Se cacula z
	fmt.Println("ex ", ex)
	fmt.Println("i ", i)
	fmt.Println("iex ", iex)
	fmt.Println("delta ", delta)
	fmt.Println("divey ", divey)
	fmt.Println("divisorey ", divisorey)
	fmt.Println("ey ", ey)
	fmt.Println("d ", d)
	fmt.Println("j ", j)
	fmt.Println("x ", x)
	fmt.Println("y1 ", y)
	fmt.Println("xex ", xex)
	fmt.Println("yey ", yey)
	fmt.Println("puntoFinal ", puntoFinal)
	return float32(puntoFinal.CoordenadaX), float32(puntoFinal.CoordenadaY)
}

func diferenciaEntreDosPuntos(p2, p1 modelos.Punto) *modelos.Punto {
	return &modelos.Punto{
		CoordenadaX: p2.CoordenadaX - p1.CoordenadaX,
		CoordenadaY: p2.CoordenadaY - p1.CoordenadaY,
	}
}

func dividirPunto(p1 *modelos.Punto, valor float64) *modelos.Punto {
	return &modelos.Punto{
		CoordenadaX: p1.CoordenadaX / valor,
		CoordenadaY: p1.CoordenadaY / valor,
	}
}

func calcularDividendo(p2, p1 modelos.Punto) float64 {
	return obtenerNorma(diferenciaEntreDosPuntos(p2, p1))
}

func obtenerNorma(punto *modelos.Punto) float64 {
	return math.Sqrt((punto.CoordenadaX * punto.CoordenadaX) + (punto.CoordenadaY * punto.CoordenadaY))
}

//multiplicarDosPuntos PRODUCTO ESCALAR
func multiplicarDosPuntos(punto1, punto2 *modelos.Punto) float64 {
	return ((punto1.CoordenadaX * punto2.CoordenadaX) + (punto1.CoordenadaY * punto2.CoordenadaY))
}

func multiplicarUnPuntoPorUnValor(punto *modelos.Punto, valor float64) *modelos.Punto {
	return &modelos.Punto{
		CoordenadaX: punto.CoordenadaX * valor,
		CoordenadaY: punto.CoordenadaY * valor,
	}
}

func sumarDosPuntos(p1, p2 *modelos.Punto) *modelos.Punto {
	return &modelos.Punto{
		CoordenadaX: p1.CoordenadaX + p2.CoordenadaX,
		CoordenadaY: p1.CoordenadaY + p2.CoordenadaY,
	}
}
