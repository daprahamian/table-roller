package basic

import (
	"encoding/json"
	"fmt"
	"strconv"
	"table-roller/dice"
	"table-roller/table"
)

// Table A basic table that rolls a single dice and returns the result
type Table struct {
	dice       dice.Dice
	diceString string
	name       string
	results    []possibleResult
}

type possibleResult struct {
	min, max uint
	result   table.Result
}

type tableJSON struct {
	Name    string                  `json:"name"`
	Dice    string                  `json:"dice"`
	Results map[string]table.Result `json:"results"`
}

// FromJSON Returns a basic table based on input json
func FromJSON(s string) (*Table, error) {
	j := tableJSON{}
	if err := json.Unmarshal([]byte(s), &j); err != nil {
		return nil, err
	}

	dice, err := dice.FromString(j.Dice)
	if err != nil {
		return nil, err
	}

	results := make([]possibleResult, len(j.Results))

	for k, jResult := range j.Results {
		min, max, err := parseDiceRange(k)
		if err != nil {
			return nil, err
		}

		result := possibleResult{
			min:    min,
			max:    max,
			result: jResult,
		}

		results = append(results, result)
	}

	ret := Table{
		name:       j.Name,
		diceString: j.Dice,
		dice:       dice,
		results:    results,
	}

	return &ret, nil
}

// Name returns the table's name
func (t *Table) Name() string {
	return t.name
}

// Roll Rolls the table's default dice against the table
func (t *Table) Roll() (*table.Result, uint, error) {
	return t.RollD(t.dice)
}

// RollD Rolls a provided Dice against the table
func (t *Table) RollD(d dice.Dice) (*table.Result, uint, error) {
	value := d.Roll()

	for _, p := range t.results {
		if p.min <= value && value <= p.max {
			tags := make([]string, len(p.result.Tags))
			copy(tags, p.result.Tags)
			ret := table.Result{
				Description: p.result.Description,
				Tags:        tags,
			}
			return &ret, value, nil
		}
	}

	return nil, value, fmt.Errorf("Rolled %d, but found no valid results", value)
}

// ToJSON returns a json representation of a basic table
func (t *Table) ToJSON() ([]byte, error) {
	tj := tableJSON{
		Name:    t.name,
		Dice:    t.diceString,
		Results: map[string]table.Result{},
	}

	for _, r := range t.results {
		min := int(r.min)
		max := int(r.max)
		var key string
		if min == max {
			key = strconv.Itoa(min)
		} else {
			key = strconv.Itoa(min) + "-" + strconv.Itoa(max)
		}
		tj.Results[key] = r.result
	}

	return json.Marshal(tj)
}
