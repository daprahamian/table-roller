package basic

import (
	"reflect"
	"table-roller/dice"
	"table-roller/table"
	"testing"
)

const testTable = `{ "dice": "d6", "name": "foobar", "results": { "1-2": { "description": "res1" }, "3": {"description": "res2"}, "4-6": { "description": "res3", "tags": ["a"] }}}`

func TestFromJSON(t *testing.T) {
	tests := []struct {
		json string
		err  bool
	}{
		{`{`, true},
		{`{"dice": "c20"}`, true},
		{testTable, false},
	}

	for _, test := range tests {
		_, err := FromJSON(test.json)
		if err == nil && test.err {
			t.Errorf("Expected FromJSON(`%s`) to error, but it did not", test.json)
		}
		if err != nil && !test.err {
			t.Errorf("Expected FromJSON(`%s`) to not error, but it did did with '%v'", test.json, err)
		}
	}
}

func TestRoll(t *testing.T) {
	tb, _ := FromJSON(testTable)

	tests := []struct {
		val    uint
		result table.Result
	}{
		{1, table.Result{Description: "res1", Tags: []string{}}},
		{2, table.Result{Description: "res1", Tags: []string{}}},
		{3, table.Result{Description: "res2", Tags: []string{}}},
		{4, table.Result{Description: "res3", Tags: []string{"a"}}},
		{5, table.Result{Description: "res3", Tags: []string{"a"}}},
		{6, table.Result{Description: "res3", Tags: []string{"a"}}},
	}

	for _, test := range tests {
		dice := dice.DC(test.val)
		res, _, err := tb.RollD(dice)
		if err != nil {
			t.Errorf(`For val=%d, expected table to not error, but it errored with "%v"`, test.val, err)
		}
		if !reflect.DeepEqual(*res, test.result) {
			t.Errorf(`For val=%d, expected (%+v), got (%+v)`, test.val, test.result, *res)
		}
	}
}
