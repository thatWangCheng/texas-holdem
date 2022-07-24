package main

import "fmt"

type Status struct {
	Cards Pokers
}

func (s *Status) AddCard(p Poker) {
	s.Cards.Add(p)
}

func (s *Status) String() string {
	var text string
	text = fmt.Sprintf("cards: %s", s.Cards)
	return text
}
