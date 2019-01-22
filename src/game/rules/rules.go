package rules

// How we move from the abstraction of "trick-taking game" to Spades, etc.
type Rules interface {
	WinningCard(trick.Trick) int
	// TODO: Scoring/points abstraction
}

type SpadesRules struct{}

// Get winning trick
func (s SpadesRules) WinningCard(t trick.Trick) int {
	// TODO
	panic("Not implemented.")
}
