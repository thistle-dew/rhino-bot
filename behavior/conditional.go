package behavior

import (
	"math/rand"
)

type Conditional interface {
	Behavior
}

// GenerateBool generates a random boolean with a `chance` percent likelihood to return true.
func GenerateBool(chance int) bool {
	rando := rand.Intn(100)
	return rando < chance
}
