package main

import "fmt"

func collatz(num int) {
	if num%2 == 0 {
		return num / 2
	} else {
		return num*3 + 1
	}
}

func main() {
	var n = 1

	fmt.Println("First", 10_000, "liczb ciągu Collatza:")
	for i := 0; i < 10_000; i++ {
		fmt.Println(n, " ")
		n = collatz(n)
	}
}