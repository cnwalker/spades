package round

import "game/deck"
import "game/hand"
import "game/player"

type handedPlayer struct {
	hand        hand.Hand
	player      player.Player
	tricksTaken int
}

type Round struct {
	hands          []hand.Hand
	startingPlayer int
}

// TODO: abstract out spades logic here? (`DistributeAll`)
func New(d deck.Deck, players []player.Player, startingPlayer int) *Round {
	deck.GenerateStandard()
	deck.Shuffle()
	hands := deck.DistributeAll(len(players))
}
