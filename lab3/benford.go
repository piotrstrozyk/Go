// package main
// import ("fmt"
// 		"flag")


// func collatz( N int64 ) {
// 	for i := range(N) {
// 		fmt.Printf("%d ", i)
// 		for i > 1 {
// 			if i % 2 == 0 {
// 				i /= 2
// 			} else{
// 				i = i*3+1
// 			}
// 			fmt.Printf("%d", i)

// 		}
// 		fmt.Println("1")
// 	}
// }

// func main() {
// 	var argN int64 = 10
// 	flag.Int64Var(&argN, "N", argN, "Wartość liczby max dla której liczymy Collatza")
// 	flag.Parse()
// 	collatz(argN)
// }
