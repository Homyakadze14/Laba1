package brackets

import (
	"errors"
	"fmt"
	"laba3/internal/stack"
	"strings"
)

const opn = "({[<"
const cls = ")}]>"

type Checker struct {
	s stack.Stack
}

func NewChecker() *Checker {
	return &Checker{
		s: stack.New(),
	}
}

func (b *Checker) Check(line string) (int, bool) {
	b.s.Clear()

	for i, elem := range line {
		ch := string(elem)
		if strings.Contains(opn, ch) {
			b.s.Push(elem)
		} else if idx := strings.Index(cls, ch); idx != -1 {
			prev, ok := b.s.Pop()
			if !ok {
				return i, false
			}

			if prev != rune(opn[idx]) {
				return i, false
			}
		}
	}

	return -1, b.s.IsEmpty()
}

func (b *Checker) CheckAndCount(line string) (int, bool) {
	b.s.Clear()

	count := 0

	for _, elem := range line {
		ch := string(elem)
		if strings.Contains(opn, ch) {
			b.s.Push(elem)
		} else if idx := strings.Index(cls, ch); idx != -1 {
			prev, ok := b.s.Pop()
			if !ok {
				return 0, false
			}

			if prev != rune(opn[idx]) {
				return 0, false
			} else {
				count++
			}
		}
	}

	return count, b.s.IsEmpty()
}

func (b *Checker) CheckOnly(line string) bool {
	b.s.Clear()
	for _, elem := range line {
		switch elem {
		case '(':
			b.s.Push(elem)
		case ')':
			if _, ok := b.s.Pop(); !ok {
				return false
			}
		}
	}
	return b.s.IsEmpty()
}

func (b *Checker) IgnoreCheck(line string, ignore string) (int, bool) {
	b.s.Clear()

	for i, elem := range line {
		ch := string(elem)
		if strings.Contains(ignore, ch) {
			continue
		}

		if strings.Contains(opn, ch) {
			b.s.Push(elem)
		} else if idx := strings.Index(cls, ch); idx != -1 {
			prev, ok := b.s.Pop()
			if !ok {
				return i, false
			}

			if prev != rune(opn[idx]) {
				return i, false
			}
		}
	}

	return -1, b.s.IsEmpty()
}

func (b *Checker) CountBrackets(line string) int {
	b.s.Clear()

	count := 0
	for _, elem := range line {
		ch := string(elem)

		if strings.Contains(opn, ch) {
			count++
		} else if strings.Contains(cls, ch) {
			count--
		}
	}

	return count
}

func (b *Checker) CheckAndCountDepth(line string) (int, bool) {
	b.s.Clear()

	maxDepth := 0
	depth := 0

	for _, elem := range line {
		ch := string(elem)
		if strings.Contains(opn, ch) {
			b.s.Push(elem)
			depth++
		} else if idx := strings.Index(cls, ch); idx != -1 {
			prev, ok := b.s.Pop()
			if !ok {
				return 0, false
			}

			if prev != rune(opn[idx]) {
				return 0, false
			} else {
				maxDepth = max(maxDepth, depth)
				depth--
			}
		}
	}

	return maxDepth, b.s.IsEmpty()
}

func (b *Checker) CheckWithError(line string) (bool, error) {
	b.s.Clear()

	for _, elem := range line {
		ch := string(elem)
		if strings.Contains(opn, ch) {
			b.s.Push(elem)
		} else if idx := strings.Index(cls, ch); idx != -1 {
			prev, ok := b.s.Pop()
			if !ok {
				return false, errors.New("Лишняя закрывающая скобка")
			}

			if prev != rune(opn[idx]) {
				return false, errors.New("Несовпадение открывающей и закрывающей скобки")
			}
		}
	}

	if !b.s.IsEmpty() {
		return false, errors.New("Нехватает закрывающей скобки")
	}

	return true, nil
}

func (b *Checker) CheckWithPrintStack(line string, print func(args ...any)) (int, bool) {
	b.s.Clear()

	for i, elem := range line {
		ch := string(elem)
		if strings.Contains(opn, ch) {
			b.s.Push(elem)
			print(fmt.Sprintf("Stack after open brack: %s", b.s.String()))
		} else if idx := strings.Index(cls, ch); idx != -1 {
			prev, ok := b.s.Pop()
			print(fmt.Sprintf("Stack after pop: %s", b.s.String()))
			if !ok {
				return i, false
			}

			if prev != rune(opn[idx]) {
				return i, false
			}
		}
	}

	print(fmt.Sprintf("Stack after checking line: %s", b.s.String()))

	return -1, b.s.IsEmpty()
}

func (b *Checker) MultiplyCheck(lines []string) []bool {
	res := make([]bool, len(lines))

	for i, line := range lines {
		_, res[i] = b.Check(line)
	}

	return res
}
