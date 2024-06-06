package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MaxX       = 80
	MaxY       = 30
	NumAnts    = 50
	NumLeaves  = 200
	RestPeriod = 10
	HexRadius  = 20
	HexHeight  = HexRadius * 2
)

var (
	HexWidth     = math.Sqrt(3) * HexRadius
	ScreenWidth  = int32(MaxX) * int32(HexWidth)
	ScreenHeight = int32(MaxY) * int32(HexHeight*3/4)
)

type Position struct {
	X, Y int
}

type Ant struct {
	Pos      Position
	Carrying bool
	Resting  int
}

type Leaf struct {
	Pos Position
}

type Board struct {
	Cells  map[Position]rune
	Ants   []*Ant
	Leaves []*Leaf
}

func NewBoard() *Board {
	board := &Board{
		Cells: make(map[Position]rune),
	}

	return board
}

func (b *Board) AddAnt(a *Ant) {
	b.Ants = append(b.Ants, a)
	b.Cells[a.Pos] = 'A'
}

func (b *Board) AddLeaf(l *Leaf) {
	b.Leaves = append(b.Leaves, l)
	b.Cells[l.Pos] = 'L'
}

func (b *Board) MoveAnt(a *Ant, newPos Position) {
	delete(b.Cells, a.Pos)
	a.Pos = newPos
	b.Cells[a.Pos] = 'A'
}

func (b *Board) FindEmptyAdjacent(pos Position) (Position, bool) {
	directions := []Position{
		{1, 0}, {1, -1}, {0, -1},
		{-1, 0}, {-1, 1}, {0, 1},
	}

	for _, dir := range directions {
		adjacentPos := Position{X: (pos.X + dir.X + MaxX) % MaxX, Y: (pos.Y + dir.Y + MaxY) % MaxY}
		if b.Cells[adjacentPos] == 0 {
			return adjacentPos, true
		}
	}

	return Position{}, false
}

func (b *Board) SimulateStep() {
	directions := []Position{
		{1, 0}, {1, -1}, {0, -1},
		{-1, 0}, {-1, 1}, {0, 1},
	}

	for _, ant := range b.Ants {
		if ant.Resting > 0 {
			dir := directions[rand.Intn(len(directions))]
			newPos := Position{X: (ant.Pos.X + dir.X + MaxX) % MaxX, Y: (ant.Pos.Y + dir.Y + MaxY) % MaxY}

			if b.Cells[newPos] == 'L' {
				adjacentPos, found := b.FindEmptyAdjacent(newPos)
				if found {
					b.AddLeaf(&Leaf{Pos: adjacentPos})
				}
			}

			b.MoveAnt(ant, newPos)
			ant.Resting--
			continue
		}

		dir := directions[rand.Intn(len(directions))]
		newPos := Position{X: (ant.Pos.X + dir.X + MaxX) % MaxX, Y: (ant.Pos.Y + dir.Y + MaxY) % MaxY}

		if b.Cells[newPos] == 'L' {
			if ant.Carrying {
				// Odłóż jeden liść na aktualne pole i drugi na sąsiednie puste pole
				adjacentPos, found := b.FindEmptyAdjacent(newPos)
				if found {
					b.AddLeaf(&Leaf{Pos: adjacentPos})
				}
				adjacentPos2, found2 := b.FindEmptyAdjacent(newPos)
				if found2 {
					b.AddLeaf(&Leaf{Pos: adjacentPos2})
				}
				ant.Carrying = false
				ant.Resting = RestPeriod
			} else {
				ant.Carrying = true
			}
		}

		b.MoveAnt(ant, newPos)
	}
}

func hexToPixel(pos Position) (float32, float32) {
	x := float32(HexWidth) * (float32(pos.X) + 0.5*float32(pos.Y%2))
	y := float32(HexHeight) * (3.0 / 4.0 * float32(pos.Y))
	return x, y
}

func drawHexagon(x, y float32, radius float32, color rl.Color) {
	points := make([]rl.Vector2, 6)
	for i := 0; i < 6; i++ {
		angle := float32(i) * (math.Pi / 3)
		points[i] = rl.NewVector2(x+radius*float32(math.Cos(float64(angle))), y+radius*float32(math.Sin(float64(angle))))
	}
	rl.DrawPoly(rl.NewVector2(x, y), 6, radius, 0, color)
	// Draw outline
	for i := 0; i < 6; i++ {
		rl.DrawLineV(points[i], points[(i+1)%6], rl.Black)
	}
}
//obrócić hexagony

func drawBoard(board *Board) {
	for pos, cell := range board.Cells {
		x, y := hexToPixel(pos)
		color := rl.NewColor(0, 121, 241, 255)
		switch cell {
		case 'A':
			for _, ant := range board.Ants {
				if ant.Pos == pos {
					if ant.Carrying {
						color = rl.Red
					} else {
						color = rl.Black
					}
					break
				}
			}
		case 'L':
			color = rl.Green
		}
		fmt.Printf("Drawing hex at %d,%d -> %f,%f with color %v\n", pos.X, pos.Y, x, y, color)
		drawHexagon(x, y, HexRadius, color)
	}
}

func drawGrid() {
	for x := 0; x < MaxX; x++ {
		for y := 0; y < MaxY; y++ {
			px, py := hexToPixel(Position{X: x, Y: y})
			drawHexagon(px, py, HexRadius, rl.NewColor(0, 121, 241, 255))
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	board := NewBoard()

	for i := 0; i < NumAnts; i++ {
		ant := &Ant{Pos: Position{X: rand.Intn(MaxX), Y: rand.Intn(MaxY)}}
		board.AddAnt(ant)
	}

	for i := 0; i < NumLeaves; i++ {
		leaf := &Leaf{Pos: Position{X: rand.Intn(MaxX), Y: rand.Intn(MaxY)}}
		board.AddLeaf(leaf)
	}

	rl.InitWindow(ScreenWidth, ScreenHeight, "Ants Simulation on Hexagonal Grid")
	defer rl.CloseWindow()

	rl.SetTargetFPS(10)

	for !rl.WindowShouldClose() {
		board.SimulateStep()

		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(0, 121, 241, 255)) // Zmieniamy kolor tła na niebieski
		drawGrid()
		drawBoard(board)
		rl.EndDrawing()
	}
}
//mrówki bez planszy? pozbyć się pustych pól dla optymalizacji

