package main

import (
    "fmt"
    "math/rand"
    "time"
)

func randomWalk(steps int) int {
    position := 0
    for i := 0; i < steps; i++ {
        if rand.Intn(2) == 0 {
            position++
        } else {
            position--
        }
    }
    return position
}

func main() {
    rand.Seed(time.Now().UnixNano())
    steps := 1000
    trials := 1000
    sum := 0
    for i := 0; i < trials; i++ {
        sum += abs(randomWalk(steps))
    }
    fmt.Printf("Average distance after %d steps: %f\n", steps, float64(sum)/float64(trials))
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}