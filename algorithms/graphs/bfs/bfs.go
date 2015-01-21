package bfs

import (
	"github.com/arnauddri/algorithms/data-structures/graph"
)

func bfs(g *graph.Graph, start graph.VertexId) map[graph.VertexId]bool {
	queue := []graph.VertexId{start}
	visited := make(map[graph.VertexId]bool)
	var next []graph.VertexId

	for len(queue) > 0 {
		next = []graph.VertexId{}
		for _, vertex := range queue {
			visited[vertex] = true
			neighbours := g.GetNeighbours(vertex).VerticesIter()

			for neighbour := range neighbours {

				_, ok := visited[neighbour]
				if !ok {
					next = append(next, neighbour)
				}
			}
		}
		queue = next
	}
	return visited
}