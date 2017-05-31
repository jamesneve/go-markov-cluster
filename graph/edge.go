package graph

type Edge struct {
	Weight float64
	Node1 *Node
	Node2 *Node
	Processed bool
}

func NewEdge(w float64, n1 *Node, n2 *Node) Edge {
	return Edge{
		Weight: w,
		Node1: n1,
		Node2: n2,
		Processed: false,
	}
}

func (e *Edge) OtherNode(n *Node) *Node {
	if n.Label == e.Node1.Label {
		return e.Node2
	} else {
		return e.Node1
	}
}