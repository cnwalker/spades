package card

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

func New(kind Kind, suit Suit) Card {
	return Card{
		suit: suit,
		kind: kind,
	}
}

func (card Card) String() string {
	return fmt.Sprintf("<%s of %s>", kindNames[card.kind], suitNames[card.suit])
}
