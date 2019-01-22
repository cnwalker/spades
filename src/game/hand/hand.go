package hand

import "game/card"

type Hand struct {
	cards []card.Card
}

func New(hand []card.Card) *Hand {
	return &Hand{
		cards: hand,
	}
}

func (h *Hand) GetCards() []card.Card {
	return h.cards
}

func (h *Hand) PlayIndex(i int) (card.Card, error) {
	if i < 0 || i >= len(hand) {
		return card.Card{}, fmt.Errorf("invalid hand index %d/%d", i, len(hand))
	}
	removedCard := h.cards[i]
	p.Hand = append(h.cards[:i], h.cards[i+1:])
	return removedCard, nil
}

func (h *Hand) PlayCard(c card.card) error {
	for d, i := range h.cards {
		if d == c {
			p.Hand = append(h.cards[:i], h.cards[i+1:])
			return nil
		}
	}
	return card.Card{}, fmt.Errorf("card %v not found", c)
}
