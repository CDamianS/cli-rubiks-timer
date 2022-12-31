package main

import (
	"fmt"
	// "github.com/charmbracelet/bubbles/list"
	// tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func main() {
	fmt.Println(getScramble())
}

const divisor = 2

var (
	columnStyle  = lipgloss.NewStyle().Padding(1, 2)
	focusedStyle = lipgloss.NewStyle().Padding(1, 2)
	helpStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

type Time struct {
	number   int
	time     float32
	scramble string
	note     string
}

func (t Time) Number() int {
	return t.number
}

func (t Time) Time() float32 {
	return t.time
}

func (t Time) Note() string {
	return t.note
}
