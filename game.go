package main

import (
	"fmt"
	"strings"
)

type Round int8

const (
	PreFlop Round = 0
	Flop    Round = 1
	Turn    Round = 2
	River   Round = 3
)

var rounds = []Round{PreFlop, Flop, Turn, River}

type PlayerStatus map[string]*Status
type PlayerMap map[string]*Player

func (ps PlayerStatus) String() string {
	var text string
	for _, s := range ps {
		text += fmt.Sprintf("%s\n", s)
	}
	return strings.TrimRight(text, "\n")
}

type Game struct {
	PokerSet     PokerSet
	PlayerStatus PlayerStatus
	PlayerMap    PlayerMap
	Pot          Pokers
	currentTurn  int
}

func NewGame(maxPlayers int8) Game {
	return Game{
		PlayerStatus: map[string]*Status{},
		PlayerMap:    map[string]*Player{},
		PokerSet:     NewPokerSet(),
	}
}

func (g Game) AddPlayer(player Player) {
	g.PlayerStatus[player.Name] = &Status{}
	g.PlayerMap[player.Name] = &player
}

func (g Game) Players() int {
	return len(g.PlayerStatus)
}

func (g Game) String() string {
	var text string
	pText := ""
	for p, s := range g.PlayerStatus {
		pText += fmt.Sprintf("%s\n%s\n", g.PlayerMap[p], s)
	}

	text += fmt.Sprintf("\n================ Players ================\n%s\n", strings.TrimRight(pText, "\n"))
	text += fmt.Sprintf("\n================== Pot ==================\n%s\n", g.Pot)
	text += fmt.Sprintf("\n=============== Poker Set ===============\n%s\n", g.PokerSet)
	return text
}

func (g *Game) Next() *Round {
	g.currentTurn += 1
	if g.currentTurn > len(rounds)-1 {
		return nil
	}
	r := rounds[g.currentTurn]
	return &r
}

func (g Game) Round() *Round {
	if g.currentTurn > len(rounds)-1 {
		return nil
	}
	r := rounds[g.currentTurn]
	return &r
}

func (g *Game) DealPot() {
	g.Pot.Add(*g.PokerSet.Cut())
}

func (g *Game) DealPlayer() {
	for i := 0; i < 2; i++ {
		for player := range g.PlayerStatus {
			g.PlayerStatus[player].AddCard(*g.PokerSet.Cut())
		}
	}
}

func (g *Game) Deal() {

	if g.currentTurn > len(rounds)-1 {
		return
	}
	switch rounds[g.currentTurn] {
	case PreFlop:
		g.DealPlayer()
		g.Next()
	case Flop:
		g.DealPot()
		g.DealPot()
		g.DealPot()
		g.Next()
	case Turn, River:
		g.DealPot()
		g.Next()
	}
}

func (g Game) FaDiPai() {
	for _, _ = range g.PlayerStatus {

	}
}

func (g Game) Bet() bool {
	for _, _ = range g.PlayerStatus {

	}
	return false
}

func (g Game) XiaoPai() {
	_ = g.PokerSet.Cut()
}

func (g Game) Compete() {

}

func (g Game) Clear() {

}

func (g *Game) Start() {

	defer func() {
		g.Clear()
	}()

	g.Deal()
	if over := g.Bet(); over {
		return
	}
	g.XiaoPai()
	g.Deal()
	if over := g.Bet(); over {
		return
	}
	g.XiaoPai()
	g.Deal()
	if over := g.Bet(); over {
		return
	}
	g.XiaoPai()
	g.Deal()
	if over := g.Bet(); over {
		return
	}
	g.Compete()
}
