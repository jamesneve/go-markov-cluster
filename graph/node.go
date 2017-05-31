package graph

type Node struct {
	Label string
	Edges []*Edge
	SeenBefore bool
}

func NewNode(l string) Node {
	return Node{Label: l, SeenBefore: false}
}