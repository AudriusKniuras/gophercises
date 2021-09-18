package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"

// 	"github.com/AudriusKniuras/gophercises/cards"
// )

// type Hand []cards.Card

// func (h Hand) String() string {
// 	strs := make([]string, len(h))
// 	// i is index; i,v := range h, here v would be card
// 	for i := range h {
// 		strs[i] = h[i].String()
// 	}
// 	return strings.Join(strs, ", ")
// }

// func (h Hand) DealerString() string {
// 	return h[0].String() + ", *HIDDEN*"
// }

// // Hit action - give a new card to a player
// func (h Hand) hit(d []cards.Card) (Hand, []cards.Card) {
// 	var card cards.Card
// 	var deck []cards.Card
// 	card, deck = d[0], d[1:]
// 	h = append(h, card)
// 	return h, deck
// }

// // Deal a new set of cards to a player
// // n - how many cards to deal
// // basically same as hit
// func (h Hand) deal(n int, d []cards.Card) (Hand, []cards.Card) {
// 	var deck []cards.Card
// 	h = d[0:n]
// 	deck = d[n:]

// 	return h, deck
// }

// func (h Hand) calculateHand() int {
// 	sum := 0
// 	for _, card := range h {
// 		if card.Rank != 1 && card.Rank < 11 {
// 			sum += int(card.Rank)
// 		} else if card.Rank > 10 {
// 			sum += 10
// 		} else if card.Rank == 1 {
// 			if sum+11 > 21 {
// 				sum += 1
// 			} else {
// 				sum += 11
// 			}
// 		}
// 	}
// 	return sum
// }

// func (h Hand) dealerPlay(d []cards.Card) (Hand, []cards.Card) {
// 	for h.calculateHand() < 18 {
// 		h, d = h.hit(d)
// 	}
// 	return h, d
// }

// func main() {
// 	deck := cards.New(cards.Deck(1), cards.Shuffle)
// 	var p1_hand Hand
// 	var dealer_hand Hand

// 	p1_hand, deck = p1_hand.deal(2, deck)
// 	fmt.Println("P1 hand is:", p1_hand)
// 	fmt.Println("P1 hand score is:", p1_hand.calculateHand())

// 	dealer_hand, deck = dealer_hand.deal(2, deck)
// 	fmt.Println("Dealer hand is:", dealer_hand.DealerString())

// 	reader := bufio.NewReader(os.Stdin)
// 	playing := true
// 	for playing {
// 		fmt.Println("Hit or Stand?")
// 		answer, _ := reader.ReadString('\n')
// 		answer = strings.ToLower(strings.TrimRight(answer, "\r\n"))

// 		if answer == "hit" {
// 			p1_hand, deck = p1_hand.hit(deck)
// 			fmt.Println("Your hand is:", p1_hand)
// 			p1_hand_score := p1_hand.calculateHand()
// 			fmt.Println("Your score is:", p1_hand_score)
// 			if p1_hand_score > 21 {
// 				fmt.Println("You lost!")
// 				playing = false
// 			} else if p1_hand_score == 21 {
// 				fmt.Println("Blackjack! You won!")
// 				playing = false
// 			}
// 		} else if answer == "stand" {
// 			dealer_hand, deck = dealer_hand.dealerPlay(deck)
// 			fmt.Println("Dealer hand is:", dealer_hand.String())
// 			fmt.Println("Dealer hand score is:", dealer_hand.calculateHand())
// 			if dealer_hand.calculateHand() > 21 || dealer_hand.calculateHand() < p1_hand.calculateHand() {
// 				fmt.Println("You won!")
// 			} else if dealer_hand.calculateHand() == p1_hand.calculateHand() {
// 				fmt.Println("Equal")
// 			} else if dealer_hand.calculateHand() > p1_hand.calculateHand() {
// 				fmt.Println("You lost!")
// 			} else {
// 				fmt.Println("Uncought condition")
// 			}
// 			playing = false
// 		} else {
// 			fmt.Println("your answer:", answer)
// 			fmt.Println("Unrecognized input. Hit or Stand?")
// 		}
// 	}

// }
