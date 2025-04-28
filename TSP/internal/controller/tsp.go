package controller

import (
	"bufio"
	"fmt"
	"github.com/AndreSS-ntp/tsp_univ/tree/main/TSP/internal/model"
	"github.com/AndreSS-ntp/tsp_univ/tree/main/TSP/internal/view"
	"math"
	"os"
	"strconv"
	"strings"
)

type TSPController struct {
	View   view.ConsoleView
	Cities []model.City
}

func (tc *TSPController) GenerateRandomCities(count int) {
	tc.Cities = make([]model.City, count)
	for i := 0; i < count; i++ {
		tc.Cities[i] = model.NewCity()
	}
	tc.View.DisplayCities(tc.Cities)
}

func (tc *TSPController) InputCities() {
	reader := bufio.NewReader(os.Stdin)
	tc.Cities = []model.City{}

	fmt.Println("Введите координаты городов (x y). Пустая строка - завершение ввода.")
	for {
		fmt.Print("Город ", len(tc.Cities)+1, ": ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			break
		}

		coords := strings.Fields(input)
		if len(coords) != 2 {
			fmt.Println("Ошибка: нужно ввести две координаты через пробел")
			continue
		}

		x, err1 := strconv.ParseFloat(coords[0], 64)
		y, err2 := strconv.ParseFloat(coords[1], 64)
		if err1 != nil || err2 != nil {
			fmt.Println("Ошибка: координаты должны быть числами")
			continue
		}

		tc.Cities = append(tc.Cities, model.City{X: x, Y: y})
	}

	tc.View.DisplayCities(tc.Cities)
}

// Алгоритм ближайшего соседа
func (tc *TSPController) SolveTSP() ([]model.City, float64) {
	if len(tc.Cities) == 0 {
		return nil, 0
	}

	unvisited := make(map[int]bool)
	for i := range tc.Cities {
		unvisited[i] = true
	}

	route := make([]model.City, 0, len(tc.Cities))
	current := 0
	delete(unvisited, current)
	route = append(route, tc.Cities[current])

	totalDistance := 0.0

	for len(unvisited) > 0 {
		nearest := -1
		minDist := math.MaxFloat64

		for cityIdx := range unvisited {
			dist := tc.Cities[current].DistanceTo(tc.Cities[cityIdx])
			if dist < minDist {
				minDist = dist
				nearest = cityIdx
			}
		}

		totalDistance += minDist
		current = nearest
		delete(unvisited, current)
		route = append(route, tc.Cities[current])
	}

	// возврат в начальный город
	totalDistance += tc.Cities[current].DistanceTo(tc.Cities[0])
	route = append(route, tc.Cities[0])

	return route, totalDistance
}
