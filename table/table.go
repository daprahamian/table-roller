package table

import (
	"table-roller/dice"
)

// Result todo
type Result struct {
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

// Table todo
type Table interface {
	RollN(dice.Dice) (*Result, uint, error)
	Roll() (*Result, uint, error)
}
