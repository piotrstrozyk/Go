package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/eiannone/keyboard"
)

const (
	width  = 20
	height = 20
	tree   = '🌲'
	fire   = '🔥'
	empty  = ' '
	axe    = '🪓'
)

type Forest struct {
	grid [][]rune
	axeX int
	axeY int
	sync.Mutex
}

func newForest() *Forest {
	rand.Seed(time.Now().UnixNano())
	forest := &Forest{
		grid: make([][]rune, height),
		axeX: width / 2,
		axeY: height / 2,
	}
	for i := range forest.grid {
		forest.grid[i] = make([]rune, width)
		for j := range forest.grid[i] {
			if rand.Float64() < 0.6 {
				forest.grid[i][j] = tree
			} else {
				forest.grid[i][j] = empty
			}
		}
	}
	forest.grid[forest.axeY][forest.axeX] = axe
	return forest
}

func (f *Forest) print() {
	f.Lock()
	defer f.Unlock()
	fmt.Print("\033[H\033[2J") // Clear terminal
	fmt.Print("    ")
	for i := 0; i < width; i++ {
		fmt.Printf("%2d ", i)
	}
	fmt.Println()
	for y, row := range f.grid {
		fmt.Printf("%2d ", y)
		for _, cell := range row {
			fmt.Printf("%c  ", cell)
		}
		fmt.Println()
	}
}

func (f *Forest) spreadFire(x, y int) {
	if x < 0 || x >= width || y < 0 || y >= height {
		return
	}
	f.Lock()
	if f.grid[y][x] != tree {
		f.Unlock()
		return
	}
	f.grid[y][x] = fire
	f.Unlock()
	f.print()

	// Spread to neighboring cells
	directions := []struct{ dx, dy int }{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
	for _, d := range directions {
		nx, ny := x+d.dx, y+d.dy
		f.Lock()
		if nx >= 0 && nx < width && ny >= 0 && ny < height && f.grid[ny][nx] == tree {
			f.grid[ny][nx] = fire
		}
		f.Unlock()
	}
}

func (f *Forest) ignite() {
	// Find a random tree to ignite
	var x, y int
	for {
		x, y = rand.Intn(width), rand.Intn(height)
		f.Lock()
		if f.grid[y][x] == tree {
			f.Unlock()
			break
		}
		f.Unlock()
	}
	f.spreadFire(x, y)
}

func (f *Forest) moveAxe(dx, dy int) {
	f.Lock()
	newX, newY := f.axeX+dx, f.axeY+dy
	if newX >= 0 && newX < width && newY >= 0 && newY < height {
		if f.grid[newY][newX] == empty || f.grid[newY][newX] == tree {
			f.grid[f.axeY][f.axeX] = empty
			f.axeX, f.axeY = newX, newY
			f.grid[f.axeY][f.axeX] = axe
		}
	}
	f.Unlock()
	f.print()
	f.spreadFireAfterCut()
}

func (f *Forest) userInteraction() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer keyboard.Close()

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch key {
		case keyboard.KeyEsc:
			fmt.Println("Exiting...")
			return
		case keyboard.KeyArrowUp, 'w':
			f.moveAxe(0, -1)
		case keyboard.KeyArrowDown, 's':
			f.moveAxe(0, 1)
		case keyboard.KeyArrowLeft, 'a':
			f.moveAxe(-1, 0)
		case keyboard.KeyArrowRight, 'd':
			f.moveAxe(1, 0)
		case keyboard.KeyEnter:
			f.Lock()
			f.Unlock()
			f.print()
			f.spreadFireAfterCut()
			f.printRemainingPercentage()
			os.Exit(0)

		}
	}
}

func (f *Forest) spreadFireAfterCut() {
	f.Lock()
	newFire := make([][2]int, 0)
	for y, row := range f.grid {
		for x, cell := range row {
			if cell == fire {
				directions := []struct{ dx, dy int }{
					{-1, -1}, {0, -1}, {1, -1},
					{-1, 0}, {1, 0},
					{-1, 1}, {0, 1}, {1, 1},
				}
				for _, d := range directions {
					nx, ny := x+d.dx, y+d.dy
					if nx >= 0 && nx < width && ny >= 0 && ny < height && f.grid[ny][nx] == tree {
						newFire = append(newFire, [2]int{nx, ny})
					}
				}
			}
		}
	}
	for _, pos := range newFire {
		f.grid[pos[1]][pos[0]] = fire
	}
	f.Unlock()
	f.print()
}

func (f *Forest) printRemainingPercentage() {
	f.Lock()
	defer f.Unlock()

	total, burned := 0, 0
	for _, row := range f.grid {
		for _, cell := range row {
			if cell == tree || cell == fire {
				total++
			}
			if cell == fire {
				burned++
			}
		}
	}

	fmt.Printf("\nPercentage of forest burned: %.2f%%\n", float64(burned)/float64(total)*100)
}

func main() {
	forest := newForest()
	forest.print()
	forest.ignite()
	forest.userInteraction()
}

//Na przykład, metoda spreadFire() sprawdza, czy dana komórka jest drzewem,
//a następnie zmienia jej stan na ogień. Jeśli dwa wątki wywołałyby tę metodę jednocześnie na
//tej samej komórce, mogłoby dojść do sytuacji, w której jeden wątek sprawdziłby, czy komórka jest drzewem,
//a następnie drugi wątek zmieniłby jej stan na ogień zanim pierwszy wątek zdążyłby to zrobić.
//W rezultacie, pierwszy wątek mógłby próbować zmienić stan komórki, która już jest ogniem,
//co mogłoby prowadzić do nieprawidłowego zachowania programu.

//go get github.com/eiannone/keyboard
