package tree

import "fmt"

type Node struct {
	key                 int
	left, right, parent *Node
}

func NewNode(key int) *Node {
	return &Node{key, nil, nil, nil}
}

func (n *Node) Key() int {
	return n.key
}

func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) Right() *Node {
	return n.right
}

func (n *Node) Parent() *Node {
	return n.parent
}

// Центрированный обход (выводит элементы по порядку)
func (n *Node) InorderTreeWalk() {
	if n != nil {
		(n.left).InorderTreeWalk()
		fmt.Println(n.key)
		(n.right).InorderTreeWalk()
	}
}

// Прямой обход
func (n *Node) StraightTreeWalk() {
	if n != nil {
		fmt.Println(n.key)
		(n.left).StraightTreeWalk()
		(n.right).StraightTreeWalk()
	}
}

// Обратный обход
func (n *Node) ReverseTreeWalk() {
	if n != nil {
		(n.left).ReverseTreeWalk()
		(n.right).ReverseTreeWalk()
		fmt.Println(n.key)
	}
}

// Поиск элемента по ключу
func (n *Node) IterativeTreeSearch(val int) *Node {
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
func (n *Node) TreeMinimum() *Node {
	cur := n
	for cur.left != nil {
		cur = cur.left
	}
	return cur
}

// Самый большой элемент
func (n *Node) TreeMaximum() *Node {
	cur := n
	for cur.right != nil {
		cur = cur.right
	}
	return cur
}

// Следующий элемент
func (n *Node) TreeSuccessor() *Node {
	if n.right != nil {
		return n.right.TreeMinimum()
	}
	for (n.parent != nil) && (n == n.parent.right) {
		n = n.parent
	}
	return n.parent
}

// Вставка элемента
func (n *Node) Insert(x *Node) *Node {

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
func (n *Node) deleteIfNotTwoKid() {
	var kid *Node
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
func (n *Node) Delete() {
	if (n.left == nil) || (n.right == nil) {
		n.deleteIfNotTwoKid()
	} else {
		nextValue := n.TreeSuccessor()
		n.key = nextValue.key
		nextValue.deleteIfNotTwoKid()
	}
}

// Создание дерева из массива
func New(values []int) *Node {
	var Tree *Node
	for _, val := range values {
		Tree = Tree.Insert(NewNode(val))
	}
	return Tree
}

// Проверка идентичности
func IsEqual(first *Node, second *Node) bool {
	if first != nil && second != nil {
		return (first.key == second.key) && (IsEqual(first.left, second.left) && (IsEqual(first.right, second.right)))
	} else {
		return first == second
	}
}
