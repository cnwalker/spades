package player

import "fmt"

import "game/card"
import "game/hand"

type Player interface {
	Play(GameState) card.Card
}
