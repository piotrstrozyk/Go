package main
import ("fmt")

func pesel() {
	liczba := 12345678971
	cyfry := make([]int, 0)

	for liczba > 0 {
		cyfra := liczba % 10
		cyfry = append([]int{cyfra}, cyfry...)
		liczba /= 10
	}
	var sum = 0
	waga := [10]int{1,3,7,9,1,3,7,9,1,3,0}
	for i, v := range cyfry{
		fmt.Println(v, waga[i])
		if v*waga[i] >= 10{
			sum += v*waga[i] % 10
		} else{
			sum += v*waga[i]
		}
		
	}
	fmt.Println(10 - sum % 10)
	
}

func main() {
	fmt.Println("Ok")
	pesel()
}