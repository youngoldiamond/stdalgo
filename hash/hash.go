package hash

import "fmt"

type Node struct {
	key        int
	next, prev *Node
}

func NewNode(key int) *Node {
	return &Node{key, nil, nil}
}

func (n *Node) Key() int {
	return n.key
}

type Table struct {
	data []*Node
	h    func(int) int
}

func NewTable(len int, h func(int) int) *Table {
	return &Table{make([]*Node, len), h}
}

func ModFunc(len int) func(int) int {
	return func(key int) int { return key % len }
}

func (t *Table) Insert(n *Node) {
	next := t.data[t.h(n.key)]
	t.data[t.h(n.key)] = n
	n.next = next
	if next != nil {
		next.prev = n
	}
}

func (t *Table) Search(key int) *Node {
	cur := t.data[t.h(key)]
	for cur != nil {
		if cur.key == key {
			return cur
		}
		cur = cur.next
	}
	return cur
}

func (t *Table) Delete(n *Node) {
	if n.prev != nil {
		n.prev.next = n.next
		if n.next != nil {
			n.next.prev = n.prev
		}
	} else {
		t.data[t.h(n.key)] = n.next
	}
}

func (t *Table) Print() {
	for i, str := range t.data {
		fmt.Printf("%d: ", i)
		cur := str
		for cur != nil {
			fmt.Printf("%d, ", cur.key)
			cur = cur.next
		}
		fmt.Println()
	}
}
