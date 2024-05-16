package main

import "fmt"

func main() {

	// Tworzę nick
	firstName := "Piotr"
	lastName := "Stróżyk"

	nick := CreateNick(firstName, lastName)
	fmt.Println("Wygenerowany nick:", nick)

	// Zamieniam nick na bajty
	bajty := NickToBytes(nick)
	fmt.Println(bajty)

	// bajty na inty...
	array := bytesToInts(bajty)

	// ...i szukam moich liczb
	findMyNumbers(array)
}
