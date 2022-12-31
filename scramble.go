package main

import (
	"math/rand"
	// "strings"
	"time"
)

func returnLetter() string {
	rand.Seed(time.Now().UnixNano())
	const letters = "UDLRFB"
	var letter byte = letters[rand.Intn(6)]
	return string(letter)
}

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

func getScramble() string {
	rand.Seed(time.Now().UnixNano())
	var scramble string
	var moveArr [25]string
	i := 1
	scramble = returnLetter() + returnIndicator()
	for i < 25 {
		moveArr[i] = returnLetter()
		if moveArr[i] != moveArr[i-1] {
			scramble = scramble + " " + moveArr[i] + returnIndicator()
			i++
		}

	}
	return scramble
}
