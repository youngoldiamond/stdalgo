package tree

import "fmt"

type Numeric interface {
	int | int8 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Node[T Numeric] struct {
	key                 T
	left, right, parent *Node[T]
}

func NewNode[T Numeric](key T) *Node[T] {
	return &Node[T]{key, nil, nil, nil}
}

func (n *Node[T]) Key() T {
	return n.key
}

func (n *Node[T]) Left() *Node[T] {
	return n.left
}

func (n *Node[T]) Right() *Node[T] {
	return n.right
}

func (n *Node[T]) Parent() *Node[T] {
	return n.parent
}

// Центрированный обход (выводит элементы по порядку)
func (n *Node[T]) InorderTreeWalk() {
	if n != nil {
		(n.left).InorderTreeWalk()
		fmt.Println(n.key)
		(n.right).InorderTreeWalk()
	}
}

// Прямой обход
func (n *Node[T]) StraightTreeWalk() {
	if n != nil {
		fmt.Println(n.key)
		(n.left).StraightTreeWalk()
		(n.right).StraightTreeWalk()
	}
}

// Обратный обход
func (n *Node[T]) ReverseTreeWalk() {
	if n != nil {
		(n.left).ReverseTreeWalk()
		(n.right).ReverseTreeWalk()
		fmt.Println(n.key)
	}
}

// Поиск элемента по ключу
func (n *Node[T]) IterativeTreeSearch(val T) *Node[T] {
	cur := n
	for (cur != nil) && (val != cur.key) {
		if val < cur.key {
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return cur
}

// Самый маленький элемент
func (n *Node[T]) TreeMinimum() *Node[T] {
	cur := n
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

// Самый большой элемент
func (n *Node[T]) TreeMaximum() *Node[T] {
	cur := n
	for cur.right != nil {
		cur = cur.right
	}
	return cur
}

// Следующий элемент
func (n *Node[T]) TreeSuccessor() *Node[T] {
	if n.right != nil {
		return n.right.TreeMinimum()
	}
	for (n.parent != nil) && (n == n.parent.right) {
		n = n.parent
	}
	return n.parent
}

// Вставка элемента
func (n *Node[T]) Insert(x *Node[T]) *Node[T] {

	//Вариант с рекурсией
	/*if n == nil {
		return n
	}
	if x.key < n.key {
		n.left = n.left.Insert(x)
		n.left.parent = n
	} else {
		n.right = n.right.Insert(x)
		n.right.parent = n
	}
	return n*/

	// Вариант без рекурсии
	if n == nil {
		return x
	}
	cur := n
	for {
		if x.key < cur.key {
			if cur.left == nil {
				cur.left = x
				x.parent = cur
				break
			} else {
				cur = cur.left
			}
		} else {
			if cur.right == nil {
				cur.right = x
				x.parent = cur
				break
			} else {
				cur = cur.right
			}
		}
	}
	return n
}

// Вспомогательная функция удаляет вершину с одним ребёнком или без детей
func (n *Node[T]) deleteIfNotTwoKid() {
	var kid *Node[T]
	if n.left == nil {
		kid = n.right
	} else if n.right == nil {
		kid = n.left
	} else {
		panic("Node have 2 kids")
	}
	if n.parent.left == n {
		n.parent.left = kid
	} else {
		n.parent.right = kid
	}
	if kid != nil {
		kid.parent = n.parent
	}
}

// Удаление элемента
func (n *Node[T]) Delete() {
	if (n.left == nil) || (n.right == nil) {
		n.deleteIfNotTwoKid()
	} else {
		nextValue := n.TreeSuccessor()
		n.key = nextValue.key
		nextValue.deleteIfNotTwoKid()
	}
}

// Создание дерева из массива
func New[T Numeric](values []T) *Node[T] {
	var Tree *Node[T]
	for _, val := range values {
		Tree = Tree.Insert(NewNode(val))
	}
	return Tree
}

// Проверка идентичности
func IsEqual[T Numeric](first *Node[T], second *Node[T]) bool {
	if first != nil && second != nil {
		return (first.key == second.key) && (IsEqual(first.left, second.left) && (IsEqual(first.right, second.right)))
	} else {
		return first == second
	}
}
