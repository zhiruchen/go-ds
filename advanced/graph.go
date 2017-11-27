package advanced

import (
	"fmt"
	"log"
)

type AdjListNode struct {
	dest int
	next *AdjListNode
}

type AdjList struct {
	head *AdjListNode
}

type Graph struct {
	v     int
	array []*AdjList
}

func newAdjListNode(dest int) *AdjListNode {
	return &AdjListNode{
		dest: dest,
	}
}

func newGraph(v int) *Graph {
	g := &Graph{
		v:     v,
		array: make([]*AdjList, 0),
	}

	for i := 0; i < v; i++ {
		g.array = append(g.array, &AdjList{})
	}

	return g
}

func (g *Graph) addEdge(v int, e int) {
	if v < 0 || v >= len(g.array) {
		log.Printf("add (%d, %d) error: v is invalid\n", v, e)
		return
	}
	node := &AdjListNode{
		dest: e,
		next: nil,
	}

	head := g.array[v].head
	if head == nil {
		g.array[v].head = node
		return
	}

	for head.next != nil {
		fmt.Printf("head != nil: %d\n", head.dest)
		head = head.next
	}

	fmt.Printf("head(%v) = next: %d\n", head, e)
	head.next = node
}

func (g *Graph) printGraph() {
	for i := 0; i < g.v; i++ {
		head := g.array[i].head

		fmt.Printf("\n adj list vertex: %d\n head", i)
		for head != nil {
			fmt.Printf("-> %d", head.dest)
			head = head.next
		}

		fmt.Println()
	}
}

func (g *Graph) DFS(v int) {
	visited := make(map[int]bool)

	g.dfs(v, visited)
}

func (g *Graph) dfs(v int, visited map[int]bool) {
	visited[v] = true
	fmt.Printf("visiting: %d\n", v)

	head := g.array[v].head
	for head != nil {
		if !visited[head.dest] {
			g.dfs(head.dest, visited)
		}

		head = head.next
	}
}

func (g *Graph) BFS(v int) {
	visited := make(map[int]bool)

	queue := make([]int, 0)
	queue = append(queue, v)
	visited[v] = true

	for len(queue) > 0 {
		v1 := queue[0]
		fmt.Printf("queue: %v\n", queue)
		queue = queue[1:]

		head := g.array[v1].head
		for head != nil {
			if !visited[head.dest] {
				fmt.Printf("visiting: %d\n", head.dest)
				visited[head.dest] = true
				queue = append(queue, head.dest)
			}

			head = head.next
		}
	}
}
