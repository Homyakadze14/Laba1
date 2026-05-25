package exptree

type Node struct {
	Value     int
	Operation string
	Left      *Node
	Right     *Node
}
