package main

import (
	"fmt"
	"sort"
	"testing"
)

func Test_NewPokerSet(t *testing.T) {
	p := NewPokerSet()
	fmt.Println(p)
	pa := p.Arr()
	sort.Sort(pa)
	fmt.Println(pa)
}

func Test_PokerSet_Cut(t *testing.T) {
	ps := NewPokerSet()
	for i := 0; i < 20; i++ {
		fmt.Println(ps.Cut())
	}
	fmt.Println(ps)
}
