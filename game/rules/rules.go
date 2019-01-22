package rules

type Rules interface {
	MaxCard
}

type SpadesRules struct{}

func (s SpadesRules) ArgMaxCard(t trick.Trick) int {

}
