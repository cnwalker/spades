package card

import "testing"

import "github.com/stretchr/testify/assert"

func TestString(t *testing.T) {
	c := New(Ace, Spades)
	assert.Equal(t, "<ace of spades>", c.String())
}
