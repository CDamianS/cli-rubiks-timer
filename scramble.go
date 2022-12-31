package main

import (
	"math/rand"
	"time"
)

// Returns random letter for a move
func returnLetter() string {
	rand.Seed(time.Now().UnixNano())
	const letters = "UDLRFB"
	var letter byte = letters[rand.Intn(6)]
	return string(letter)
}

// Returns an indicator for prime and double moves
func returnIndicator() string {
	rand.Seed(time.Now().UnixNano())
	ind := rand.Intn(3)
	switch ind {
	default:
		return ""
	case 1:
		return "2"
	case 2:
		return "'"
	}
}

// Merges letter and indicator in a string
func getScramble() string {
	var scramble string
	var moveArr [25]string
	i := 1
	for i < rand.Intn(5)+20 {
		moveArr[i] = returnLetter()
		// Check if the move is the same as the previous one
		if moveArr[i] != moveArr[i-1] {
			scramble = scramble + " " + moveArr[i] + returnIndicator()
			i++
		}

	}
	return scramble
}
