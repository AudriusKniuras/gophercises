package main

import (
	"fmt"
	"strings"

	"github.com/AudriusKniuras/gophercises/cards"
)

type Hand []cards.Card

type Player interface {
	stand()
	hit() Hand
}

func (h Hand) String() string {
	strs := make([]string, len(h))
	// i is index; i,v := range h, here v would be card
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func main() {
	deck := cards.New(cards.Deck(1), cards.Shuffle)
	var card cards.Card
	for i := 0; i < 10; i++ {
		_, deck = deck[0], deck[1:]
		fmt.Println(card)
	}
	var h Hand = deck[0:3]
	fmt.Println(h) // easy to read, uses custom String() method
	// fmt.Println(deck[0:3]) // hard to read, not separated
}
