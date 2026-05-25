package stack

import (
	"fmt"
	"strings"
)

const defcap = 10

type Stack struct {
	arr []interface{}
}

func New() Stack {
	return Stack{
		arr: make([]interface{}, 0, defcap),
	}
}

func (s *Stack) IsEmpty() bool {
	return len(s.arr) == 0
}

func (s *Stack) Push(elem interface{}) {
	s.arr = append(s.arr, elem)
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	elem := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return elem, true
}

func (s *Stack) Peek() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	return s.arr[len(s.arr)-1], true
}

func (s *Stack) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for i, elem := range s.arr {
		fmt.Fprintf(&sb, "%v", elem)
		if i < len(s.arr)-1 {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

func (s *Stack) Clear() {
	s.arr = make([]interface{}, 0, cap(s.arr))
}
