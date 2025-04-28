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
	view := view.ConsoleView{}
	tsp := controller.TSPController{View: view}

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
			fmt.Print("Введите количество городов: ")
			countInput, _ := reader.ReadString('\n')
			count, err := strconv.Atoi(strings.TrimSpace(countInput))
			if err != nil || count <= 0 {
				fmt.Println("Ошибка: введите положительное число")
				continue
			}
			tsp.GenerateRandomCities(count)
			route, dist := tsp.SolveTSP()
			view.DisplayRoute(route, dist)

		case 2:
			tsp.InputCities()
			if len(tsp.Cities) > 0 {
				route, dist := tsp.SolveTSP()
				view.DisplayRoute(route, dist)
			}

		case 3:
			fmt.Println("Выход...")
			os.Exit(0)

		default:
			fmt.Println("Неверный выбор. Попробуйте снова.")
		}
	}
}
