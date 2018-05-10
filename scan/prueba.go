package main

import (
	"fmt"
)

func main() {
	var pagoPorHora float64
	fmt.Println("Cuanto te pagan por hora negruri?")
	fmt.Scan(&pagoPorHora)

	var HorasPorDia float64
	fmt.Println("Cuantas horas trabajaste por dia negruri?")
	fmt.Scan(&HorasPorDia)

	var DiasTrabajados float64
	fmt.Println("Cuantos dias de la semana trabajaste negruri?")
	fmt.Scan(&DiasTrabajados)

	pagoTotal := pagoPorHora * HorasPorDia * DiasTrabajados

	fmt.Println("Entonces te van a pagar ", pagoTotal)

}
