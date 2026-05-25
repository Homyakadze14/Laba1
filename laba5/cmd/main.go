package main

import (
	"fmt"
	"laba5/internal/exptree"
)

func main() {
	exp := []string{"3", "4", "+", "5", "*"}
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

	exp, err = exptree.ConvertToPostfix("(2-3)*4+5*6")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	tree, err = exptree.Build(exp)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}

	res, err = exptree.Evaluate(tree.Parent)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return
	}
	fmt.Printf("Result: %d\n", res)
}
