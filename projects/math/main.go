package main

import (
	"fmt"
	"math"
)

var k, p, v float64 = 1296, 6, 6

// Функция для вычисления массы
func M() float64 {
	return p * v
}

// Функция для вычисления циклической частоты
func W() float64 {
	m := M()
	return math.Sqrt(k / m)
}

// Функция для вычисления периода колебаний
func T() float64 {
	w := W()
	return 6 / w
}

func main() {
	// Пример использования функции T()
	period := T()
	fmt.Println("Период колебаний (t):", period) //1
}
