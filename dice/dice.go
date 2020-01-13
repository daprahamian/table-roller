package dice

import (
	"fmt"
	"regexp"
	"strconv"
)

// Dice Interface for a set of dice
type Dice interface {
	Roll() uint
}

var dXRegexp = regexp.MustCompile("^d([1-9][0-9]*)$")
var dCRegexp = regexp.MustCompile("^([1-9][0-9]*)$")

// FromString returns a Dice based on the input string (ex: "d6")
func FromString(diceCombo string) (Dice, error) {
	if val := intFromPattern(dXRegexp, diceCombo); val > 0 {
		return DX(val), nil
	}
	if val := intFromPattern(dCRegexp, diceCombo); val > 0 {
		return DC(val), nil
	}
	return nil, fmt.Errorf("Invalid dice format %s", diceCombo)
}

// Roll rolls a dice based on the input string (ex: "d6") and returns the result
func Roll(diceCombo string) (uint, error) {
	roller, err := FromString(diceCombo)
	if err != nil {
		return 0, err
	}
	return roller.Roll(), nil
}

func intFromPattern(pattern *regexp.Regexp, input string) uint {
	if val := pattern.FindStringSubmatch(input); len(val) > 0 {
		num, err := strconv.ParseUint(val[1], 10, 0)
		if num > 0 && err == nil {
			return uint(num)
		}
	}
	return 0
}
