package main

import "fmt"

func colatz(num int) {
	i := 0
	for i = 0; num != 1; i++ {
		if num%2 == 0 {
			num /= 2
		} else {
			num = num*3 + 1
		}
	} 
	fmt.Println(i)
}

func main() {
	var executions []int = make ([]int, 1)

	for i := 1; i <= 1000; i++ {
		executions = append(executions, colatz(i))
	}
	sum := 0

	for i := 0; i <= len(executions) - 1; i++ {
		sum += executions[i]
	}
	mean := sum / len(executions)

	fmt.Println("Sum: ", sum)
}