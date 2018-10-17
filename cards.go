package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Suit int

type Kind int

// Suits
const (
	Hearts Suit = iota
	Spades
	Clubs
	Diamonds
)

// Kinds
const (
	Two         Kind = 2
	Three       Kind = 3
	Four        Kind = 4
	Five        Kind = 5
	Six         Kind = 6
	Seven       Kind = 7
	Eight       Kind = 8
	Nine        Kind = 9
	Ten         Kind = 10
	Jack        Kind = 11
	Queen       Kind = 12
	King        Kind = 13
	Ace         Kind = 14
	LittleJoker Kind = 15
	BigJoker    Kind = 16
)

var suitNames = map[Suit]string{
	Hearts:   "hearts",
	Spades:   "spades",
	Clubs:    "clubs",
	Diamonds: "diamonds",
}

var kindNames = map[Kind]string{
	Two:         "two",
	Three:       "three",
	Four:        "four",
	Five:        "five",
	Six:         "six",
	Seven:       "seven",
	Eight:       "eight",
	Nine:        "nine",
	Ten:         "ten",
	Jack:        "jack",
	Queen:       "queen",
	King:        "king",
	Ace:         "ace",
	LittleJoker: "little joker",
	BigJoker:    "big joker",
}

type Card struct {
	suit Suit
	kind Kind
}

// NOTE: Do we really need a deck or just functions on cards
// IDEA: OOOH we can just make a list of cards a Decks

type Deck struct {
	cards []*Card
}

func NewCard(kind Kind, suit Suit) *Card {
	return &Card{
		suit: suit,
		kind: kind,
	}
}

func (card *Card) String() string {
	return fmt.Sprintf("<%s of %s>", kindNames[card.kind], suitNames[card.suit])
}

func NewDeck(initialCards []*Card) *Deck {
	return &Deck{
		cards: initialCards,
	}
}

// Utilities
func (deck *Deck) Add(card *Card) {
	deck.cards = append(deck.cards, card)
}

func (deck *Deck) Clear() {
	deck.cards = nil
}

func (deck *Deck) Pop() (*Card, error) {
	if len(deck.cards) == 0 {
		return nil, fmt.Errorf("No cards remaining to pop")
	}
	poppedCard := deck.cards[len(deck.cards)-1]
	deck.cards = deck.cards[:len(deck.cards)-1]
	return poppedCard, nil
}

func (deck *Deck) NumCards() int {
	return len(deck.cards)
}

// TODO: shuffle in-place
func (deck *Deck) Shuffle() {
	if len(deck.cards) == 0 {
		return
	}

	newCards := []*Card{}
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	for _, index := range randomizer.Perm(len(deck.cards)) {
		newCards = append(newCards, deck.cards[index])
	}
	deck.cards = newCards
}

func (deck *Deck) Distribute(numGroups int, numCards int) ([][]*Card, error) {
	// NOTE: Don't allow more cards than can fill the groups?
	var curGroup int
	cardGroups := make([][]*Card, numGroups)

	for cardNum := 0; cardNum < numCards; cardNum++ {
		poppedCard, err := deck.Pop()
		if err != nil {
			return nil, err
		}
		cardGroups[curGroup] = append(cardGroups[curGroup], poppedCard)

		if curGroup > numGroups {
			curGroup %= numGroups
		}
	}

	return cardGroups, nil
}

// Generation functions
func (deck *Deck) GenerateStandard() {
	deck.Clear()
	for suit := 0; suit <= 3; suit++ {
		for kind := 2; kind <= 14; kind++ {
			deck.cards = append(deck.cards, NewCard(Kind(kind), Suit(suit)))
		}
	}
}

func main() {
	testCard := NewCard(Ace, Diamonds)
	fmt.Println(testCard)
}
