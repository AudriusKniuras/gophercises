package cards

import (
	"fmt"
	"testing"
)

func ExampleCard() {
	fmt.Println(Card{Rank: Ace, Suit: Heart})
	fmt.Println(Card{Rank: Two, Suit: Club})
	fmt.Println(Card{Suit: Joker})
	// Output:
	// Ace of Hearts
	// Two of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := New()
	if len(cards) != 52 {
		t.Error("Wrong number of cards")
	}
}

func TestDefaultSort(t *testing.T) {
	cards := New(DefaultSort)
	if len(cards) != 52 {
		t.Error("Wrong number of cards")
	}
}

// exactly the same thing as DefaultSort, but verifies that additional Sort funcion is implemented correctly
// Sort function allows package user to create their own sort functions
func TestSort(t *testing.T) {
	cards := New(Sort(Less))
	if len(cards) != 52 {
		t.Error("Wrong number of cards")
	}
}

func TestJokers(t *testing.T) {
	cards := New(Jokers(3))
	count := 0
	for _, c := range cards {
		if c.Suit == Joker {
			count++
		}
	}
	if count != 3 {
		t.Error("Expected 3 jokers, got:", count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	cards := New(Filter(filter))

	for _, c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("Found rank Two or Three")
		}
	}

}

func TestDeck(t *testing.T) {
	cards := New(Deck(3))
	if len(cards) != 52*3 {
		t.Errorf("Expected %d cards, got: %d", 52*3, len(cards))
	}
}
