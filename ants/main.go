package main

import (
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	MaxX       = 80
	MaxY       = 40
	NumAnts    = 50
	NumLeaves  = 200
	RestPeriod = 10
	HexRadius  = 20
	HexHeight  = HexRadius * 2
	Scale      = 0.7
)

var (
	HexWidth     = math.Sqrt(3) * HexRadius
	ScreenWidth  = int32(float32(MaxX) * float32(HexWidth) * Scale)
	ScreenHeight = int32(float32(MaxY) * float32(HexHeight) * Scale)
	OffsetX      = (int32(rl.GetScreenWidth()) - ScreenWidth) / 2
	OffsetY      = (int32(rl.GetScreenHeight()) - ScreenHeight) / 2
)

type Position struct {
	X, Y int
}

type MyRand struct {
	seed int64
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
	Ants   []*Ant
	Leaves []*Leaf
}

func NewMyRand() *MyRand {
	return &MyRand{seed: time.Now().UnixNano()}
}

func (r *MyRand) Intn(n int) int {
	r.seed = (r.seed*9301 + 49297) % 233280
	return int(r.seed % int64(n))
}

func NewBoard() *Board {
	board := &Board{}
	return board
}

func (b *Board) AddAnt(a *Ant) {
	b.Ants = append(b.Ants, a)
}

func (b *Board) AddLeaf(l *Leaf) {
	b.Leaves = append(b.Leaves, l)
}

func (b *Board) MoveAnt(a *Ant, newPos Position) {
	a.Pos = newPos
}

func (b *Board) FindEmptyAdjacent(pos Position) (Position, bool) {
	directions := []Position{
		{1, 0}, {1, -1}, {0, -1},
		{-1, 0}, {-1, 1}, {0, 1},
	}

	for _, dir := range directions {
		adjacentPos := Position{X: (pos.X + dir.X + MaxX) % MaxX, Y: (pos.Y + dir.Y + MaxY) % MaxY}
		if !b.IsOccupied(adjacentPos) {
			return adjacentPos, true
		}
	}

	return Position{}, false
}

func (b *Board) IsOccupied(pos Position) bool {
	for _, ant := range b.Ants {
		if ant.Pos == pos {
			return true
		}
	}
	for _, leaf := range b.Leaves {
		if leaf.Pos == pos {
			return true
		}
	}
	return false
}

func (b *Board) SimulateStep() {

	myRand := NewMyRand()
	directions := []Position{
		{1, 0}, {1, -1}, {0, -1},
		{-1, 0}, {-1, 1}, {0, 1},
	}

	for _, ant := range b.Ants {
		dir := directions[myRand.Intn(len(directions))]
		newPos := Position{X: (ant.Pos.X + dir.X + MaxX) % MaxX, Y: (ant.Pos.Y + dir.Y + MaxY) % MaxY}

		if ant.Resting > 0 {
			// Ant is resting, just move without interacting with leaves
			ant.Resting--
			b.MoveAnt(ant, newPos)
			continue
		}

		if leaf := b.LeafAt(newPos); leaf != nil {
			if ant.Carrying {
				// Drop the leaf and start resting
				adjacentPos, found := b.FindEmptyAdjacent(newPos)
				if found {
					b.AddLeaf(&Leaf{Pos: adjacentPos})
				}
				ant.Carrying = false
				ant.Resting = RestPeriod
			} else {
				// Pick up the leaf
				ant.Carrying = true
				b.RemoveLeaf(leaf)
			}
		}

		b.MoveAnt(ant, newPos)
	}
}

func (b *Board) LeafAt(pos Position) *Leaf {
	for _, leaf := range b.Leaves {
		if leaf.Pos == pos {
			return leaf
		}
	}
	return nil
}

func (b *Board) RemoveLeaf(leaf *Leaf) {
	for i, l := range b.Leaves {
		if l == leaf {
			b.Leaves = append(b.Leaves[:i], b.Leaves[i+1:]...)
			break
		}
	}
}

func hexToPixel(pos Position) (float32, float32) {
	x := float32(HexRadius) * 3.0 / 2.0 * float32(pos.X)
	y := float32(HexHeight) * (float32(pos.Y) + 0.5*float32(pos.X%2))
	return x*Scale + float32(OffsetX), y*Scale + float32(OffsetY) // Zastosowanie skali i przesunięcia
}

func drawHexagon(x, y float32, radius float32, color rl.Color) {
	points := make([]rl.Vector2, 6)
	for i := 0; i < 6; i++ {
		angle := float32(i) * (math.Pi / 3)
		points[i] = rl.NewVector2(x+radius*float32(math.Cos(float64(angle)))*Scale, y+radius*float32(math.Sin(float64(angle)))*Scale) // Zastosowanie skali
	}
	rl.DrawPoly(rl.NewVector2(x, y), 6, radius*Scale, 0, color) // Zastosowanie skali
}

func drawBoard(board *Board) {
	for _, ant := range board.Ants {
		x, y := hexToPixel(ant.Pos)
		color := rl.Black
		if ant.Carrying {
			color = rl.Red
		}
		drawHexagon(x, y, HexRadius, color)
	}

	for _, leaf := range board.Leaves {
		x, y := hexToPixel(leaf.Pos)
		drawHexagon(x, y, HexRadius, rl.Green)
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

	rl.InitWindow(int32(float32(MaxX)*float32(HexWidth)*Scale), int32(float32(MaxY)*float32(HexHeight*3/4)*Scale), "Ants Simulation on Hexagonal Grid")
	defer rl.CloseWindow()

	OffsetX = (int32(rl.GetScreenWidth()) - ScreenWidth) / 2
	OffsetY = (int32(rl.GetScreenHeight()) - ScreenHeight) / 2
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
