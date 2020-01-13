package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/daprahamian/table-roller/basic"
	"github.com/daprahamian/table-roller/dice"
	"github.com/daprahamian/table-roller/table"
)

func main() {
	filename, roll := getArgs()
	filejson, err := ioutil.ReadFile(filename)
	check(err)

	tb, err := basic.FromJSON(string(filejson))
	check(err)

	result, rollResult, err := rollOnTable(tb, roll)

	fmt.Printf("You rolled on table %s\n", tb.Name())
	fmt.Printf("  You rolled a %d\n", rollResult)
	fmt.Printf("  Result: \"%s\"\n", result.Description)
}

const usage = `
Usage: 
  roll <path-to-file> [roll value]
`

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getArgs() (string, uint) {
	args := os.Args[1:]
	if len(args) == 0 || args[0] == "" {
		fmt.Println(usage)
		os.Exit(1)
	}

	filename := args[0]
	var roll uint = 0
	if len(args) >= 2 {
		tmp, err := strconv.ParseUint(args[1], 10, 0)
		check(err)
		if tmp <= 0 {
			fmt.Printf("Cannot roll with value <=0, but got %d\n", tmp)
			fmt.Println(usage)
			os.Exit(1)
		}
		roll = uint(tmp)
	}

	return filename, roll
}

func rollOnTable(tb *basic.Table, rl uint) (*table.Result, uint, error) {
	if rl > 0 {
		return tb.RollD(dice.DC(rl))
	}
	return tb.Roll()
}
