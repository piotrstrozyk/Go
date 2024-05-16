package main

import "strings"

func ToASCII(s string) string {
	// Mapa zamiany polskich znaków na ASCII
	toASCII := map[rune]rune{
		'ą': 'a', 'ć': 'c', 'ę': 'e', 'ł': 'l', 'ń': 'n',
		'ó': 'o', 'ś': 's', 'ź': 'z', 'ż': 'z',
	}

	s = strings.ToLower(s)

	for i, r := range s {
		if replacement, exists := toASCII[r]; exists {
			s = s[:i] + string(replacement) + s[i+1:]
		}
	}
	return s
}

func CreateNick(firstName, lastName string) string {

	firstName = ToASCII(firstName)
	lastName = ToASCII(lastName)

	firstName = firstName[:3]

	lastName = lastName[:3]

	nick := firstName + lastName

	return nick
}

func NickToBytes(nick string) []byte {
	return []byte(nick)
}

func bytesToInts(bytes []byte) []int {
	ints := make([]int, len(bytes))
	for i, b := range bytes {
		ints[i] = int(b)
	}
	return ints
}
