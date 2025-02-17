package hash

import (
	"errors"
	"fmt"
)

type Node struct {
	key        int
	next, prev *Node
}

// Создание нового элемента таблицы
func NewNode(key int) *Node {
	return &Node{key, nil, nil}
}

func (n *Node) Key() int {
	return n.key
}

// Хэш-таблица с закрытой адресацией
type ClosedTable struct {
	data []*Node
	h    func(int) int
}

// Создание новой таблицы с закрытой адресацией
func NewClosedTable(length int, h func(int) int) (*ClosedTable, error) {
	if h == nil {
		return nil, errors.New("Nill hash func in NewClosedTable()")
	} else if length < 1 {
		return nil, errors.New("Length is less than 1 in NewClosedTable()")
	}
	return &ClosedTable{make([]*Node, length), h}, nil
}

// Вставка элемента
func (t *ClosedTable) InsertNode(n *Node) {
	if n == nil {
		return
	}
	next := t.data[t.h(n.key)]
	t.data[t.h(n.key)] = n
	n.next = next
	if next != nil {
		next.prev = n
	}
}

// Вставка ключа
func (t *ClosedTable) Insert(key int) {
	t.InsertNode(NewNode(key))
}

// Поиск элкмента по ключу
func (t *ClosedTable) SearchNode(key int) *Node {
	cur := t.data[t.h(key)]
	for cur != nil {
		if cur.key == key {
			return cur
		}
		cur = cur.next
	}
	return cur
}

// Содержит ли таблица ключ
func (t *ClosedTable) Search(key int) bool {
	if t.SearchNode(key) != nil {
		return true
	} else {
		return false
	}
}

// Удаление элемента
func (t *ClosedTable) DeleteNode(n *Node) {
	if n == nil {
		return
	}
	if n.prev != nil {
		n.prev.next = n.next
		if n.next != nil {
			n.next.prev = n.prev
		}
	} else {
		t.data[t.h(n.key)] = n.next
	}
}

// Удаление по ключу
func (t *ClosedTable) Delete(key int) {
	t.DeleteNode(t.SearchNode(key))
}

// Вывод таблицы по хэшам
func (t *ClosedTable) Print() {
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
