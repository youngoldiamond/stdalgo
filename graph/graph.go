package graph

import "github.com/youngoldiamond/stdalgo/queue"

type Node struct {
	key      int
	adjacent []*Node
}

func Bfs(start *Node) map[int]int {
	res := make(map[int]int)

	q, err := queue.New[*Node](5)
	if err != nil {
		panic(err)
	}
	q.Push(start)
	for !q.Empty() {
		cur := q.Pop()
		for _, adj := range cur.adjacent {
			if res[adj.key] == 0 {
				res[adj.key] = res[cur.key] + 1
				q.Push(adj)
			}
		}
	}

	return res
}
