package exptree

import (
	"errors"
	"laba5/internal/stack"
	"math"
	"strconv"
	"strings"
)

var (
	ErrTokCount  = errors.New("Incorrect number of tokens")
	ErrBadTokens = errors.New("Incorrect tokens")
	ErrEmptyTree = errors.New("Empty tree")
	ErrBadExp    = errors.New("Bad expression")
)

type Tree struct {
	Parent *Node
}

const operands = "+-*/^"

func Build(tokens []string) (*Tree, error) {
	if len(tokens)%2 == 0 {
		return nil, ErrTokCount
	}

	stack := stack.New()
	for _, elem := range tokens {
		if strings.Contains(operands, elem) {
			right, ok := stack.Pop()
			if !ok {
				return nil, ErrBadTokens
			}

			left, ok := stack.Pop()
			if !ok {
				return nil, ErrBadTokens
			}

			op := &Node{
				Operation: elem,
				Right:     right.(*Node),
				Left:      left.(*Node),
			}

			stack.Push(op)
			continue
		}

		val, err := strconv.Atoi(elem)
		if err != nil {
			return nil, ErrBadTokens
		}
		node := &Node{
			Value: val,
		}
		stack.Push(node)
	}

	parent, ok := stack.Pop()
	if !ok || !stack.IsEmpty() {
		return nil, ErrBadTokens
	}

	return &Tree{
		Parent: parent.(*Node),
	}, nil
}

func Evaluate(root *Node) (int, error) {
	if root == nil {
		return 0, ErrEmptyTree
	}

	if root.Operation != "" {
		left, err := Evaluate(root.Left)
		if err != nil {
			return 0, err
		}
		right, err := Evaluate(root.Right)
		if err != nil {
			return 0, err
		}

		switch root.Operation {
		case "+":
			return left + right, nil
		case "-":
			return left - right, nil
		case "*":
			return left * right, nil
		case "/":
			if right == 0 {
				return 0, errors.New("division by zero")
			}
			return left / right, nil
		case "^":
			if right < 0 {
				return 0, errors.New("negative exponent not supported for integers")
			}
			return int(math.Pow(float64(left), float64(right))), nil
		}
	}

	return root.Value, nil
}

func PostfixString(root *Node) string {
	if root == nil {
		return ""
	}

	if root.Operation == "" {
		return strconv.Itoa(root.Value)
	}

	left := PostfixString(root.Left)
	right := PostfixString(root.Right)
	return left + " " + right + " " + root.Operation
}

func PrefixString(root *Node) string {
	if root == nil {
		return ""
	}

	if root.Operation == "" {
		return strconv.Itoa(root.Value)
	}

	left := PrefixString(root.Left)
	right := PrefixString(root.Right)
	return root.Operation + " " + left + " " + right
}

func Height(root *Node) int {
	if root == nil {
		return 0
	}
	if root.Operation == "" {
		return 1
	}
	leftH := Height(root.Left)
	rightH := Height(root.Right)
	if leftH > rightH {
		return 1 + leftH
	}
	return 1 + rightH
}

func CountOperations(root *Node) int {
	if root == nil || root.Operation == "" {
		return 0
	}
	count := 1
	return count + CountOperations(root.Left) + CountOperations(root.Right)
}
