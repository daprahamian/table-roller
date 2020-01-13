package dice

import (
	"math/rand"
	"time"
)

// X A dice that rolls between 1 and max
type X struct {
	max  uint
	rand *rand.Rand
}

// DX returns a dice that rolls between 1 and max
func DX(max uint) Dice {
	return &X{
		max:  max,
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// Roll Rolls a dice values
func (x *X) Roll() uint {
	return uint(x.rand.Intn(int(x.max))) + 1
}
