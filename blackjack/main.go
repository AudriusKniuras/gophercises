package main

import (
	"fmt"
	"strings"

	"github.com/AudriusKniuras/gophercises/cards"
)

type Hand []cards.Card

type State int8

const (
	StatePlayerTurn State = iota
	StateDealerTurn
	StateHandOver
)

type GameState struct {
	Deck   []cards.Card
	State  State
	Player Hand
	Dealer Hand
}

func (h Hand) String() string {
	strs := make([]string, len(h))
	// i is index; i,v := range h, here v would be card
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", *HIDDEN*"
}

func (gs *GameState) CurrentPlayer() *Hand {
	switch gs.State {
	case StatePlayerTurn:
		return &gs.Player
	case StateDealerTurn:
		return &gs.Dealer
	default:
		panic("Wrong game state")
	}
}

// not very efficient. Might be better to use pointers to...
// alter the game state instead of copying it
func clone(gs GameState) GameState {
	// not doing ret := gs because GameState struct might change
	// and we might be altering gamestate incorrectly
	// not necessarily in this game, but in general
	ret := GameState{
		Deck:   make([]cards.Card, len(gs.Deck)),
		State:  gs.State,
		Player: make(Hand, len(gs.Player)),
		Dealer: make(Hand, len(gs.Dealer)),
	}
	copy(ret.Deck, gs.Deck)
	copy(ret.Player, gs.Player)
	copy(ret.Dealer, gs.Dealer)
	return ret
}

func Shuffle(gs GameState) GameState {
	ret := clone(gs)
	ret.Deck = cards.New(cards.Deck(1), cards.Shuffle)
	return ret
}

func Deal(gs GameState) GameState {
	ret := clone(gs)
	// 5 - capacity of the initial slice
	ret.Player = make(Hand, 0, 5)
	ret.Dealer = make(Hand, 0, 5)
	var card cards.Card
	// 1 card player, 1 card dealer until they both have 2 cards
	for i := 0; i < 2; i++ {
		card, ret.Deck = ret.Deck[0], ret.Deck[1:]
		ret.Player = append(ret.Player, card)
		card, ret.Deck = ret.Deck[0], ret.Deck[1:]
		ret.Dealer = append(ret.Dealer, card)

	}
	ret.State = StatePlayerTurn
	return ret
}

func (h Hand) calculateHand() int {
	sum := 0
	for _, card := range h {
		if card.Rank != 1 && card.Rank < 11 {
			sum += int(card.Rank)
		} else if card.Rank > 10 {
			sum += 10
		} else if card.Rank == 1 {
			if sum+11 > 21 {
				sum += 1
			} else {
				sum += 11
			}
		}
	}
	return sum
}

func Hit(gs GameState) GameState {
	ret := clone(gs)
	hand := ret.CurrentPlayer()
	var card cards.Card
	card, ret.Deck = gs.Deck[0], ret.Deck[1:]
	*hand = append(*hand, card)
	if hand.calculateHand() > 21 {
		return Stand(ret)
	}
	return ret
}

func Stand(gs GameState) GameState {
	ret := clone(gs)
	ret.State++
	return ret
}

func EndHand(gs GameState) GameState {
	ret := clone(gs)

	pScore, dScore := ret.Player.calculateHand(), ret.Dealer.calculateHand()
	fmt.Println("Player hand:", ret.Player, ", Score:", pScore)
	fmt.Println("Dealer hand:", ret.Dealer, ", Score:", dScore)

	switch {
	case pScore > 21:
		fmt.Println("You Lost!")
	case dScore > 21 || pScore > dScore:
		fmt.Println("You Won!")
	case dScore > pScore:
		fmt.Println("You Lost!")
	case dScore == pScore:
		fmt.Println("Draw!")
	default:
		fmt.Println("Uncought end hand")
	}
	fmt.Println()
	gs.Dealer = nil
	gs.Player = nil
	return ret
}

func main() {
	var gs GameState
	gs = Shuffle(gs)
	gs = Deal(gs)

	var input string
	for gs.State == StatePlayerTurn {
		fmt.Println("Player:", gs.Player.String(), ", Score:", gs.Player.calculateHand())
		fmt.Println("Dealer:", gs.Dealer.DealerString())
		fmt.Println("Your move: (h)it, (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			gs = Hit(gs)
		case "s":
			gs = Stand(gs)
		default:
			fmt.Println("Invalid option", input)
		}
	}

	for gs.State == StateDealerTurn {
		if gs.Dealer.calculateHand() < 18 {
			gs = Hit(gs)
		} else {
			gs = Stand(gs)
		}
	}

	gs = EndHand(gs)
}
