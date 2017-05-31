# Markov Clustering

Go implementation of Markov Clustering. It isn't currently particularly well optimised or anything (although it
mostly uses mat64 methods for matrix operations, so I imagine it does OK). But there wasn't another one for
golang, so it's better than nothing.

I've tested it on graphs of a few thousand nodes, where it completes in a couple of seconds.

## Usage

If you've already got a graph in adjacency matrix form, then you can just throw it directly into the
MCL method

```go
m := NewMCL(power, inflation, maxLoops)
m.GenerateClusters(adjacencyMatrix)
```

If you just have a bunch of data, I've included a small graph implementation.

```go
g := graph.NewGraph()
n1 := graph.NewNode("1")
n2 := graph.NewNode("2")
n3 := graph.NewNode("3")
n4 := graph.NewNode("4")

g.AddNode(&n1)
g.AddNode(&n2)
g.AddNode(&n3)
g.AddNode(&n4)

g.AddEdge(&n1, &n2, 2)
g.AddEdge(&n2, &n3, 2)
g.AddEdge(&n1, &n3, 2)

c, err := g.GetClusters(2, 2, 100)
```

The implementation won't output clusters smaller than 2 elements, and won't repeat clusters.