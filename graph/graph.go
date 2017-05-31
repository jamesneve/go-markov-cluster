package graph

import (
	"github.com/gonum/matrix/mat64"
	"github.com/jamesneve/go-mcl/mcl"
)

type Graph struct {
	Nodes []*Node
	Edges []*Edge
}

func NewGraph() Graph {
	return Graph{}
}

func (g *Graph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(n1 *Node, n2 *Node, w float64) {
	e := NewEdge(w, n1, n2)

	n1.Edges = append(n1.Edges, &e)
	n2.Edges = append(n2.Edges, &e)
	g.Edges = append(g.Edges, &e)
}

func (g *Graph) clearSeenBefore() {
	for _, n := range g.Nodes {
		n.SeenBefore = false
	}

	for _, e := range g.Edges {
		e.Processed = false
	}
}

func (g *Graph) GetClusters(p, i, l int) {
	m := g.MakeAdjacencyMatrix()
	markovClusters := mcl.NewMCL(p, i, l)
	markovClusters.GenerateClusters(m)
}

func (g *Graph) MakeAdjacencyMatrix() *mat64.Dense {
	g.clearSeenBefore()

	m := mat64.NewDense(len(g.Nodes), len(g.Nodes), nil)
	if len(g.Nodes) < 2 {
		return m
	}

	g.Nodes[0].SeenBefore = true
	queue := []*Node{g.Nodes[0]}
	for len(queue) != 0 {
		n := queue[0]
		queue = queue[1:]

		for _, e := range n.Edges {
			p1 := g.nodePosition(n)
			o := e.OtherNode(n)
			if !e.Processed {
				p2 := g.nodePosition(o)
				m.Set(p1, p2, e.Weight)
				m.Set(p2, p1, e.Weight)
				e.Processed = true
			}

			if !o.SeenBefore {
				queue = append(queue, o)
				o.SeenBefore = true
			}
		}
	}

	g.clearSeenBefore()
	return m
}

func (g *Graph) nodePosition(n *Node) int {
	for i, n2 := range g.Nodes {
		if n == n2 {
			return i
		}
	}
	return 0
}