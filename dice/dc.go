package dice

// C A dice that rolls a constant value
type C struct {
	val uint
}

// DC Returns a Dice that rolls a constant value every time
func DC(val uint) Dice {
	return &C{val: val}
}

// Roll roll
func (c *C) Roll() uint {
	return c.val
}
