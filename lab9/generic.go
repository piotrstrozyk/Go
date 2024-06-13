package main

import (
	"fmt"
	"math"
)

type Numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

type Point[T Numeric] struct {
	X T
	Y T
}

func (p Point[T]) Distance() T {
	return T(math.Sqrt(float64(p.X)*float64(p.X) + float64(p.Y)*float64(p.Y)))
}

func main() {
	p1 := Point[int]{X: 3, Y: 4}
	fmt.Printf("Distance from (0,0) to (%d,%d): %v\n", p1.X, p1.Y, p1.Distance())

	p2 := Point[float64]{X: 3.0, Y: 4.0}
	fmt.Printf("Distance from (0,0) to (%.1f,%.1f): %v\n", p2.X, p2.Y, p2.Distance())
}
