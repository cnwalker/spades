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
	hands := make([]hand.Hand, numGroups)

	for cardNum := 0; cardNum < numCards; cardNum++ {
		poppedCard, err := deck.Pop()
		if err != nil {
			return nil, err
		}
		curGroup := cardNum % numGroups
		cardGroups[curGroup] = append(cardGroups[curGroup], poppedCard)
	}

	return cardGroups, nil
}

// Distributes all cards in the deck evenly. If they cannot be distributed
// evenly, deal as many cards as possible to make it even.
func (deck *Deck) DistributeAll(numGroups int) ([]hand.Hand, error) {
	// Assumption: want to distribute cards evenly, even if that means that not
	// all cards are dealt.
	size := deck.Size()
	roundedSize := (size / numGroups) * numGroups
	return deck.Distribute(numGroups, roundedSize)
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
