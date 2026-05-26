package main

import (
	"fmt"
	"laba5/internal/exptree"
)

func main() {
	exp := []string{"3", "4", "+", "5", "*", "2", "^"}
	tree, err := exptree.Build(exp)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}

	res, err := exptree.Evaluate(tree.Parent)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	fmt.Printf("Result: %d\n", res)

	psfx := exptree.PostfixString(tree.Parent)
	fmt.Printf("Result: %s\n", psfx)

	prfx := exptree.PrefixString(tree.Parent)
	fmt.Printf("Result: %s\n", prfx)

	res = exptree.Height(tree.Parent)
	fmt.Printf("Result: %d\n", res)

	res = exptree.CountOperations(tree.Parent)
	fmt.Printf("Result: %d\n", res)
}
