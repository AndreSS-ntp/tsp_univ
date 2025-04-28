package view

import (
	"fmt"
	"github.com/AndreSS-ntp/tsp_univ/tree/main/TSP/internal/model"
)

type ConsoleView struct {
	MapWidth  int
	MapHeight int
}

func NewConsoleView(width, height int) ConsoleView {
	return ConsoleView{
		MapWidth:  width,
		MapHeight: height,
	}
}

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

func (cv ConsoleView) DisplayMap(cities []model.City, route []int) {
	// Создаем карту
	grid := make([][]rune, cv.MapHeight)
	for i := range grid {
		grid[i] = make([]rune, cv.MapWidth)
		for j := range grid[i] {
			grid[i][j] = ' '
		}
	}

	// Нормализуем координаты городов к размерам карты
	normalized := make([]struct{ X, Y int }, len(cities))
	maxX, maxY := 0.0, 0.0
	for _, city := range cities {
		if city.X > maxX {
			maxX = city.X
		}
		if city.Y > maxY {
			maxY = city.Y
		}
	}

	for i, city := range cities {
		normalized[i].X = int((city.X / maxX) * float64(cv.MapWidth-1))
		normalized[i].Y = int((city.Y / maxY) * float64(cv.MapHeight-1))
	}

	// Рисуем маршруты
	for i := 0; i < len(route)-1; i++ {
		from := route[i]
		to := route[i+1]
		cv.drawLine(grid, normalized[from].X, normalized[from].Y, normalized[to].X, normalized[to].Y)
	}

	// Рисуем города
	for i, pos := range normalized {
		label := rune('A' + i)
		if i >= 26 {
			label = rune('0' + i - 26)
		}
		grid[pos.Y][pos.X] = label
	}

	// Выводим карту
	fmt.Println("\nКарта маршрута:")
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (cv ConsoleView) drawLine(grid [][]rune, x0, y0, x1, y1 int) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	sx, sy := 1, 1
	if x0 >= x1 {
		sx = -1
	}
	if y0 >= y1 {
		sy = -1
	}
	err := dx - dy

	for {
		if grid[y0][x0] == ' ' {
			if dx > dy {
				grid[y0][x0] = '-'
			} else {
				grid[y0][x0] = '|'
			}
		}

		if x0 == x1 && y0 == y1 {
			break
		}

		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x0 += sx
		}
		if e2 < dx {
			err += dx
			y0 += sy
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
