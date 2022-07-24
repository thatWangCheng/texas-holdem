package main

import (
	"fmt"
	"testing"
)

func Test_NewGame(t *testing.T) {
	defaultChips := []Chip{{Value: 100}}
	g := NewGame(5)
	p1 := NewPlayer("Player 1")
	p2 := NewPlayer("Player 2")
	p1.FillChips(defaultChips)
	p2.FillChips(defaultChips)
	g.AddPlayer(p1)
	g.AddPlayer(p2)
	g.Start()
	fmt.Println(g)
}
