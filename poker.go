package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Suit struct {
	Color  string
	Symbol string
}

var colorToCode = map[string]string{
	"red":   "31m",
	"black": "37m",
}

func (s Suit) String() string {
	return fmt.Sprintf("\033[1;31;%s%s\033[0m", colorToCode[s.Color], s.Symbol)
}

var (
	Heart   = Suit{Color: "red", Symbol: "♥️️"}
	Club    = Suit{Color: "black", Symbol: "♣️"}
	Spade   = Suit{Color: "black", Symbol: "♠️"}
	Diamond = Suit{Color: "red", Symbol: "♦️️"}
)

type Poker struct {
	Name  string
	Value int8
	Suit  Suit
}

func (p *Poker) String() string {
	if p == nil {
		return ""
	}
	coloredName := fmt.Sprintf("\033[1;31;%s%s\033[0m", colorToCode[p.Suit.Color], p.Name)
	return fmt.Sprintf("%s%s", p.Suit, coloredName)
}

func (p *Poker) SameColor(pp Poker) bool {
	return p.Suit.Symbol == pp.Suit.Symbol
}

func (p *Poker) SameValue(pp Poker) bool {
	return p.Value == pp.Value
}

var pokerValues = []int8{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
var pokerSuites = []Suit{Heart, Club, Spade, Diamond}
var valueNames = map[int8]string{
	2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "10", 11: "J", 12: "Q", 13: "K", 14: "A",
}

type PokerSet map[int]*Poker

var pokerSet = PokerSet{}

func NewPoker() Poker {
	rand.Seed(time.Now().UnixNano())
	idx1 := rand.Intn(len(pokerValues))
	idx2 := rand.Intn(len(pokerSuites))
	return Poker{
		Value: pokerValues[idx1],
		Name:  valueNames[pokerValues[idx1]],
		Suit:  pokerSuites[idx2],
	}
}

func PokerSetCapacity() int {
	return len(pokerValues) * len(pokerSuites)
}

func (ps PokerSet) Size() int {
	return len(ps)
}

func (ps PokerSet) Empty() bool {
	for _, p := range ps {
		if p != nil {
			return false
		}
	}
	return true
}

func (ps PokerSet) String() string {
	text := ""
	cnt := 0
	for _, poker := range ps {
		cnt += 1
		if cnt%13 == 1 {
			text += fmt.Sprintf("%s", poker)
		} else {
			text += fmt.Sprintf("\t%s", poker)
		}
		if cnt%13 == 0 {
			text += "\n"
			strings.TrimLeft(text, "\t")
		}
	}
	return strings.TrimRight(text, "\n")
}

func (ps PokerSet) Exist(v Poker) bool {
	for _, p := range ps {
		if p != nil && *p == v {
			return true
		}
	}
	return false
}

func (ps PokerSet) Arr() (arr Pokers) {
	for _, v := range ps {
		if v != nil {
			arr = append(arr, *v)
		}
	}
	return arr
}

type Pokers []Poker

func (ps *Pokers) Add(p Poker) {
	*ps = append(*ps, p)
}

func (ps Pokers) String() string {
	text := ""
	for _, p := range ps {
		text += fmt.Sprintf("%s ", p.String())
	}
	return strings.TrimRight(text, " ")
}

func (ps Pokers) Len() int {
	return len(ps)
}

func (ps Pokers) Less(i, j int) bool {
	if ps[i].Value != ps[j].Value {
		return ps[i].Value < ps[j].Value
	}
	return ps[i].Suit.Symbol > ps[j].Suit.Symbol
}

func (ps Pokers) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func NewPokerSet() PokerSet {
	for i := 0; i < len(pokerValues); i++ {
		for j := 0; j < len(pokerSuites); j++ {
			p := NewPoker()
			for pokerSet.Exist(p) {
				p = NewPoker()
			}
			pokerSet[j+i*len(pokerSuites)] = &p
		}
	}
	return pokerSet
}

func (ps PokerSet) Cut() *Poker {
	// 牌没生成完整
	if ps.Size() != PokerSetCapacity() || ps.Empty() {
		return nil
	}

	rand.Seed(time.Now().UnixNano())
	idx := rand.Intn(len(pokerValues) * len(pokerSuites))

	for ps[idx] == nil {
		rand.Seed(time.Now().UnixNano())
		idx = rand.Intn(len(pokerValues) * len(pokerSuites))
	}

	p := ps[idx]
	ps[idx] = nil
	return p
}
