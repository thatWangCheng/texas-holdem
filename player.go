package main

import "fmt"

type Player struct {
	Name  string
	Chips []Chip
}

func NewPlayer(name string) Player {
	return Player{
		Name:  name,
		Chips: []Chip{},
	}
}

func (p Player) EqualTo(pp Player) bool {
	return p.Name == pp.Name
}

func (p *Player) FillChips(chips []Chip) {
	p.Chips = append(p.Chips, chips...)
}

func (p Player) String() string {
	var text string
	text += fmt.Sprintf("\033[1;31;32m%s\033[0m\t\033[1;31;33m$%d\033[0m", p.Name, p.Funds())
	return text
}

func (p Player) Funds() int64 {
	var res int64
	for _, v := range p.Chips {
		res += v.Value
	}
	return res
}
