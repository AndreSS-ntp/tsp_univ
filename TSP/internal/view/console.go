package view

import (
	"fmt"
	"github.com/AndreSS-ntp/tsp_univ/tree/main/TSP/internal/model"
)

type ConsoleView struct{}

func (cv ConsoleView) DisplayCities(cities []model.City) {
	fmt.Println("Города:")
	for i, city := range cities {
		fmt.Printf("%d: (%.2f, %.2f)\n", i+1, city.X, city.Y)
	}
	fmt.Println()
}

func (cv ConsoleView) DisplayRoute(route []model.City, distance float64) {
	fmt.Println("Оптимальный маршрут:")
	for i, city := range route {
		fmt.Printf("%d: (%.2f, %.2f)\n", i+1, city.X, city.Y)
	}
	fmt.Printf("Общее расстояние: %.2f\n", distance)
}

func (cv ConsoleView) DisplayMenu() {
	fmt.Println("1. Сгенерировать случайные города")
	fmt.Println("2. Ввести города вручную")
	fmt.Println("3. Выход")
	fmt.Print("Выберите вариант: ")
}
