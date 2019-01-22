package deck

import "game/card"

type Deck struct {
	cards []card.Card
}

func New(initialCards []card.Card) *Deck {
	return &Deck{
		cards: initialCards,
	}
}

// Utilities
func (deck *Deck) Add(card *card.Card) {
	deck.cards = append(deck.cards, card)
}

func (deck *Deck) Clear() {
	deck.cards = nil
}

func (deck *Deck) Pop() (*card.Card, error) {
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

func (deck *Deck) Shuffle() {
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	randomizer.Shuffle(len(deck.cards), func(i, j int) {
		deck.cards[i], deck.cards[j] = deck.cards[j], deck.cards[i]
	})
}

func (deck *Deck) Distribute(numGroups int, numCards int) ([]hand.Hand, error) {
	// NOTE: Don't allow more cards than can fill the groups?
	var curGroup int
	hands := make([]hand.Hand, numGroups)

	for cardNum := 0; cardNum < numCards; cardNum++ {
		poppedCard, err := deck.Pop()
		if err != nil {
			return nil, err
		}
		cardGroups[curGroup] = append(cardGroups[curGroup], poppedCard)

		curGroup %= numGroups
	}

	return cardGroups, nil
}

func (deck *Deck) DistributeAll(numGroups int) ([]hand.Hand, error) {

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

func (deck *Deck) Size() int {
	return len(deck.cards)
}
