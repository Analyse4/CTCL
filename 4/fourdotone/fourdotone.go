package fourdotone

import "github.com/Analyse4/mew/adt/graph"

func GenerateExGraph() *graph.Graph {
	g := graph.New(7)
	g.AddEdge(0, 2)
	g.AddEdge(0, 1)
	g.AddEdge(0, 5)
	g.AddEdge(2, 1)
	g.AddEdge(2, 3)
	g.AddEdge(2, 4)
	g.AddEdge(3, 5)
	g.AddEdge(3, 4)
	return g
}

// O(V+E)
func RouteBetweenNodeDFS(o, e int, g *graph.Graph, mark []bool) bool {
	if mark[o] {
		return false
	}
	mark[o] = true
	if o == e {
		return true
	}
	for _, n := range g.Adj(o) {
		if RouteBetweenNodeDFS(n.Value, e, g, mark) {
			return true
		}
	}
	return false
}

// O(V+E)
func RouteBetweenNodeBFS(o, e int, g *graph.Graph) bool {
	q := make([]*graph.Node, 0)
	mark := make([]bool, g.V())
	q = append(q, g.Nodes[o])
	for len(q) != 0 {
		n := q[0]
		q = q[1:]
		if mark[n.Value] {
			continue
		}
		mark[n.Value] = true
		if n.Value == e {
			return true
		}
		for _, elem := range g.Adj(n.Value) {
			q = append(q, elem)
		}
	}
	return false
}
