package graph

import "github.com/youngoldiamond/stdalgo/queue"

type adjacencyList [][]int

func newAdjacencyList(adjacents [][]int) *adjacencyList {
	gr := new(adjacencyList)
	*gr = adjacencyList(adjacents)
	return gr
}

func (g *adjacencyList) Bfs(start int) ([]int, error) {
	distance := make([]int, len([][]int(*g)))
	for i := range distance {
		distance[i] = -1
	}

	q, err := queue.New[int](5)
	if err != nil {
		return nil, err
	}
	q.Push(start)

	distance[start] = 0

	for !q.Empty() {
		cur := q.Pop()
		for _, adj := range (*g)[cur] {
			if distance[adj] != -1 {
				continue
			}
			distance[adj] = distance[cur] + 1
			q.Push(adj)
		}
	}

	return distance, nil
}

func (g *adjacencyList) Dfs(start int) {

}
