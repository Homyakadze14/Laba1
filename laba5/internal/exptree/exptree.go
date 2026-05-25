package exptree

import (
	"errors"
	"laba5/internal/stack"
	"strconv"
	"strings"
)

const defcap = 10

var (
	ErrTokCount  = errors.New("Incorrect number of tokens")
	ErrBadTokens = errors.New("Incorrect tokens")
	ErrEmptyTree = errors.New("Empty tree")
	ErrBadExp    = errors.New("Bad expression")
)

type Tree struct {
	Parent *Node
}

const operands = "+-*/"

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
		}
	}

	return root.Value, nil
}

func ConvertToPostfix(exp string) ([]string, error) {
	pr := "*/"
	out := make([]string, 0, defcap)
	stack := stack.New()

	for _, elem := range exp {
		el := string(elem)
		if el == ")" {
			op, ok := stack.Pop()
			for ok {
				if op == "(" {
					break
				}
				out = append(out, op.(string))
				op, ok = stack.Pop()
			}
			if !ok && op != "(" {
				return nil, ErrBadExp
			}
		} else if strings.Contains(operands, el) {
			if !stack.IsEmpty() {
				sel, _ := stack.Pop()
				op := sel.(string)
				if (strings.Contains(pr, op) && strings.Contains(pr, el)) || (!strings.Contains(pr, op) && strings.Contains(pr, el)) {
					out = append(out, el)
					stack.Push(op)
				} else {
					out = append(out, op)
					stack.Push(el)
				}
			} else {
				out = append(out, el)
			}
		} else {
			out = append(out, el)
		}
	}

	if !stack.IsEmpty() {
		return nil, ErrBadExp
	}

	return out, nil
}
