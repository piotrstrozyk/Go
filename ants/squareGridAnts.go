// package main

// import (
// 	"math/rand"
// 	"time"

// 	rl "github.com/gen2brain/raylib-go/raylib"
// )

// const (
// 	MaxX         = 100
// 	MaxY         = 50
// 	NumAnts      = 20
// 	NumLeaves    = 100
// 	RestPeriod   = 5
// 	CellSize     = 10
// 	ScreenWidth  = MaxX * CellSize
// 	ScreenHeight = MaxY * CellSize
// )

// type Position struct {
// 	X, Y int
// }

// type Ant struct {
// 	Pos      Position
// 	Carrying bool
// 	Resting  int
// }

// type Leaf struct {
// 	Pos Position
// }

// type Board struct {
// 	Cells  [][]rune
// 	Ants   []*Ant
// 	Leaves []*Leaf
// }

// func NewBoard(maxX, maxY int) *Board {
// 	board := &Board{
// 		Cells: make([][]rune, maxY),
// 	}

// 	for i := range board.Cells {
// 		board.Cells[i] = make([]rune, maxX)
// 		for j := range board.Cells[i] {
// 			board.Cells[i][j] = '.'
// 		}
// 	}

// 	return board
// }

// func (b *Board) AddAnt(a *Ant) {
// 	b.Ants = append(b.Ants, a)
// 	b.Cells[a.Pos.Y][a.Pos.X] = 'A'
// }

// func (b *Board) AddLeaf(l *Leaf) {
// 	b.Leaves = append(b.Leaves, l)
// 	b.Cells[l.Pos.Y][l.Pos.X] = 'L'
// }

// func (b *Board) MoveAnt(a *Ant, newPos Position) {
// 	b.Cells[a.Pos.Y][a.Pos.X] = '.'
// 	a.Pos = newPos
// 	b.Cells[a.Pos.Y][a.Pos.X] = 'A'
// }

// func (b *Board) SimulateStep() {
// 	for _, ant := range b.Ants {
// 		if ant.Resting > 0 {
// 			ant.Resting--
// 			continue
// 		}

// 		newPos := ant.Pos
// 		switch rand.Intn(4) {
// 		case 0:
// 			newPos.X = (newPos.X + 1) % MaxX
// 		case 1:
// 			newPos.X = (newPos.X - 1 + MaxX) % MaxX
// 		case 2:
// 			newPos.Y = (newPos.Y + 1) % MaxY
// 		case 3:
// 			newPos.Y = (newPos.Y - 1 + MaxY) % MaxY
// 		}

// 		if b.Cells[newPos.Y][newPos.X] == 'L' {
// 			if ant.Carrying {
// 				ant.Carrying = false
// 				b.AddLeaf(&Leaf{Pos: ant.Pos})
// 				ant.Resting = RestPeriod
// 			} else {
// 				ant.Carrying = true
// 			}
// 		}

// 		b.MoveAnt(ant, newPos)
// 	}
// }
// func drawBoard(board *Board) {
// 	for y, row := range board.Cells {
// 		for x, cell := range row {
// 			color := rl.Gray
// 			switch cell {
// 			case 'A':
// 				color = rl.Black
// 				for _, ant := range board.Ants {
// 					if ant.Pos.X == x && ant.Pos.Y == y {
// 						if ant.Carrying {
// 							color = rl.Red
// 						}
// 						break
// 					}
// 				}
// 			case 'L':
// 				color = rl.Green
// 			}
// 			rl.DrawRectangle(int32(x*CellSize), int32(y*CellSize), CellSize, CellSize, color)
// 		}
// 	}
// }

// func main() {
// 	rand.Seed(time.Now().UnixNano())

// 	board := NewBoard(MaxX, MaxY)

// 	for i := 0; i < NumAnts; i++ {
// 		ant := &Ant{Pos: Position{X: rand.Intn(MaxX), Y: rand.Intn(MaxY)}}
// 		board.AddAnt(ant)
// 	}

// 	for i := 0; i < NumLeaves; i++ {
// 		leaf := &Leaf{Pos: Position{X: rand.Intn(MaxX), Y: rand.Intn(MaxY)}}
// 		board.AddLeaf(leaf)
// 	}

// 	rl.InitWindow(ScreenWidth, ScreenHeight, "Ants Simulation")
// 	defer rl.CloseWindow()

// 	rl.SetTargetFPS(10)

// 	for !rl.WindowShouldClose() {
// 		board.SimulateStep()

// 		rl.BeginDrawing()
// 		rl.ClearBackground(rl.RayWhite)
// 		drawBoard(board)
// 		rl.EndDrawing()
// 	}
// }
