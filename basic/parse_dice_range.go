package basic

import (
	"fmt"
	"regexp"
	"strconv"
)

var matcher = regexp.MustCompile(`^([1-9][0-9]*)(?:-([1-9][0-9]*))?$`)

func parseDiceRange(rng string) (uint, uint, error) {
	if tokens := matcher.FindStringSubmatch(rng); len(tokens) > 0 {
		min, err := strconv.ParseUint(tokens[1], 10, 0)
		if err == nil {
			umin := uint(min)
			if len(tokens[2]) == 0 {
				return umin, umin, nil
			}
			max, err2 := strconv.ParseUint(tokens[2], 10, 0)
			if err2 == nil {
				return umin, uint(max), nil
			}
		}
	}

	return 0, 0, fmt.Errorf(`Malformed dice range "%s"`, rng)
}
