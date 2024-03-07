package main
import "fmt"

func collatz(num int64) (int){
	fmt.Println("Pierwsze", 30, "liczb ciągu Collatza:")
	if num%2 == 0 {
		return num / 2
	} else {
		return num*3 + 1
	}
}

func main() {
	var n = 61

	fmt.Println("Pierwsze", 30, "liczb ciągu Collatza:")
	for i := 0; i < 30; i++ {
		fmt.Println(n, " ")
		n = collatz(n)
	}
}

