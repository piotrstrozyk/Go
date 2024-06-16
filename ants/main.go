package main

import (
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Constants for the simulation
const (
	MaxX       = 80            // Maximum X coordinate
	MaxY       = 40            // Maximum Y coordinate
	NumAnts    = 50            // Number of ants
	NumLeaves  = 200           // Number of leaves
	RestPeriod = 10            // Rest period for ants
	HexRadius  = 20            // Radius of hexagon
	HexHeight  = HexRadius * 2 // Height of hexagon
	Scale      = 0.7           // Scale for drawing
)

// Variables for the simulation
var (
	HexWidth     = math.Sqrt(3) * HexRadius
	ScreenWidth  = int32(float32(MaxX) * float32(HexWidth) * Scale)
	ScreenHeight = int32(float32(MaxY) * float32(HexHeight) * Scale)
	OffsetX      = (int32(rl.GetScreenWidth()) - ScreenWidth) / 2
	OffsetY      = (int32(rl.GetScreenHeight()) - ScreenHeight) / 2
)

// Position struct represents a position on the board
type Position struct {
	X, Y int
}

// MyRand struct represents a random number generator
type MyRand struct {
	seed int64
}

// Ant struct represents an ant
type Ant struct {
	Pos      Position // Position of the ant
	Carrying bool     // Whether the ant is carrying a leaf
	Resting  int      // How long the ant is resting
}

// Leaf struct represents a leaf
type Leaf struct {
	Pos Position // Position of the leaf
}

// Board struct represents the board of the simulation
type Board struct {
	Ants   []*Ant  // List of ants
	Leaves []*Leaf // List of leaves
}

// NewMyRand creates a new random number generator to confirm a possible diagonal bias
func NewMyRand() *MyRand {
	return &MyRand{seed: time.Now().UnixNano()}
}

// Intn generates a random integer
func (r *MyRand) Intn(n int) int {
	r.seed = (r.seed*9301 + 49297) % 233280
	return int(r.seed % int64(n))
}

// NewBoard creates a new board
func NewBoard() *Board {
	board := &Board{}
	return board
}

// AddAnt adds an ant to the board
func (b *Board) AddAnt(a *Ant) {
	b.Ants = append(b.Ants, a)
}

// AddLeaf adds a leaf to the board
func (b *Board) AddLeaf(l *Leaf) {
	b.Leaves = append(b.Leaves, l)
}

// MoveAnt moves an ant to a new position
func (b *Board) MoveAnt(a *Ant, newPos Position) {
	a.Pos = newPos
}

// FindEmptyAdjacent finds an empty adjacent position
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

// IsOccupied checks if a position is occupied
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

// SimulateStep simulates a step of the simulation
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

// LeafAt finds a leaf at a position
func (b *Board) LeafAt(pos Position) *Leaf {
	for _, leaf := range b.Leaves {
		if leaf.Pos == pos {
			return leaf
		}
	}
	return nil
}

// RemoveLeaf removes a leaf from the board
func (b *Board) RemoveLeaf(leaf *Leaf) {
	for i, l := range b.Leaves {
		if l == leaf {
			b.Leaves = append(b.Leaves[:i], b.Leaves[i+1:]...)
			break
		}
	}
}

// hexToPixel converts a hexagonal coordinate to a pixel coordinate
func hexToPixel(pos Position) (float32, float32) {
	x := float32(HexRadius) * 3.0 / 2.0 * float32(pos.X)
	y := float32(HexHeight) * (float32(pos.Y) + 0.5*float32(pos.X%2))
	return x*Scale + float32(OffsetX), y*Scale + float32(OffsetY) // Apply scale and offset
}

// drawHexagon draws a hexagon
func drawHexagon(x, y float32, radius float32, color rl.Color) {
	points := make([]rl.Vector2, 6)
	for i := 0; i < 6; i++ {
		angle := float32(i) * (math.Pi / 3)
		points[i] = rl.NewVector2(x+radius*float32(math.Cos(float64(angle)))*Scale, y+radius*float32(math.Sin(float64(angle)))*Scale) // Apply scale
	}
	rl.DrawPoly(rl.NewVector2(x, y), 6, radius*Scale, 0, color) // Apply scale
}

// drawBoard draws the board
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

// drawGrid draws the grid
func drawGrid() {
	for x := 0; x < MaxX; x++ {
		for y := 0; y < MaxY; y++ {
			px, py := hexToPixel(Position{X: x, Y: y})
			drawHexagon(px, py, HexRadius, rl.NewColor(0, 121, 241, 255))
		}
	}
}

// main function of the program
func main() {
	rand.Seed(time.Now().UnixNano())

	board := NewBoard()

	// Add ants to the board
	for i := 0; i < NumAnts; i++ {
		ant := &Ant{Pos: Position{X: rand.Intn(MaxX), Y: rand.Intn(MaxY)}}
		board.AddAnt(ant)
	}

	// Add leaves to the board
	for i := 0; i < NumLeaves; i++ {
		leaf := &Leaf{Pos: Position{X: rand.Intn(MaxX), Y: rand.Intn(MaxY)}}
		board.AddLeaf(leaf)
	}

	// Initialize the window
	rl.InitWindow(int32(float32(MaxX)*float32(HexWidth)*Scale), int32(float32(MaxY)*float32(HexHeight*3/4)*Scale), "Ants Simulation on Hexagonal Grid")
	defer rl.CloseWindow()

	OffsetX = (int32(rl.GetScreenWidth()) - ScreenWidth) / 2
	OffsetY = (int32(rl.GetScreenHeight()) - ScreenHeight) / 2
	rl.SetTargetFPS(10)

	// Main loop of the simulation
	for !rl.WindowShouldClose() {
		board.SimulateStep()

		rl.BeginDrawing()
		rl.ClearBackground(rl.NewColor(0, 121, 241, 255)) // Change the background color to blue
		drawGrid()
		drawBoard(board)
		rl.EndDrawing()
	}
}
