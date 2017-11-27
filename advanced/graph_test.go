package advanced

import "testing"

func TestGraph_BFS(t *testing.T) {
	graph := newGraph(6)

	graph.addEdge(0, 2)
	graph.addEdge(0, 3)
	graph.addEdge(0, 1)
	graph.addEdge(1, 2)
	graph.addEdge(1, 3)
	graph.addEdge(2, 1)
	graph.addEdge(2, 3)
	graph.addEdge(3, 2)
	graph.addEdge(3, 1)
	graph.addEdge(3, 4)
	graph.addEdge(3, 5)

	graph.BFS(0)
	graph.DFS(0)
}

func TestGraph_DFS(t *testing.T) {

}
