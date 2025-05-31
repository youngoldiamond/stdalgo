package graph

type Graph interface {
	Bfs(start int) ([]int, error)
	Dfs(start int)
}

func New(data any) *Graph {
	var gr Graph
	switch data := data.(type) {
	default:
		return nil
	case [][]int:
		gr = newAdjacencyList(data)
		return &gr
	}
}
