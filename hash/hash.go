package hash

import "fmt"

type Node struct {
	key        int
	next, prev *Node
}

// Создание нового элемента таблицы
func NewNode(key int) *Node {
	return &Node{key, nil, nil}
}

func (n *Node) Key() int {
	if n == nil {
		panic("Nill pointer in *Node.Key()")
	}
	return n.key
}

type Table struct {
	data []*Node
	h    func(int) int
}

// Создание новой таблицы с закрытой адресацией
func NewTable(len int, h func(int) int) *Table {
	if h == nil {
		panic("Nill hash func in NewTable()")
	}
	return &Table{make([]*Node, len), h}
}

// Создаёт простую хэш функцию с методом деления
func ModFunc(len int) func(int) int {
	if len < 1 {
		panic("Length is less than 1 in ModFunc()")
	}
	return func(key int) int { return key % len }
}

// Создаёт хэш функцию с методом умножения
func MultFunc(len int) func(int) int {
	if len < 1 {
		panic("Length is less than 1 in MultFunc()")
	}
	const gr = 1.0 / 1.6180339887
	return func(key int) int {
		return int(float64(len) * (float64(key)*gr - float64(int(float64(key)*gr))))
	}
}

// Вставка элемента
func (t *Table) Insert(n *Node) {
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

// Поиск по ключу
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

// Удаление элемента
func (t *Table) Delete(n *Node) {
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

// Вывод таблицы по хэшам
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
