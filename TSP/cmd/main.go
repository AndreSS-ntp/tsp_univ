package main

import (
	"bufio"
	"fmt"
	"github.com/AndreSS-ntp/tsp_univ/tree/main/TSP/internal/controller"
	"github.com/AndreSS-ntp/tsp_univ/tree/main/TSP/internal/view"
	"os"
	"strconv"
	"strings"
)

func main() {
	view := view.NewConsoleView(60, 20)
	tsp := controller.NewTSPController(view)

	reader := bufio.NewReader(os.Stdin)

	for {
		view.DisplayMenu()
		input, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println("Ошибка: введите число")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Введите количество городов (2-26): ")
			countInput, _ := reader.ReadString('\n')
			count, err := strconv.Atoi(strings.TrimSpace(countInput))
			if err != nil || count <= 1 || count > 26 {
				fmt.Println("Ошибка: введите число от 2 до 26")
				continue
			}
			tsp.GenerateRandomCities(count)
			route, routeIndices, dist := tsp.SolveTSP()
			view.DisplayRoute(route, dist)
			view.DisplayMap(tsp.Cities, routeIndices)

		case 2:
			tsp.InputCities()
			if len(tsp.Cities) > 0 {
				route, routeIndices, dist := tsp.SolveTSP()
				view.DisplayRoute(route, dist)
				view.DisplayMap(tsp.Cities, routeIndices)
			}

		case 3:
			fmt.Println("Выход...")
			os.Exit(0)

		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}
