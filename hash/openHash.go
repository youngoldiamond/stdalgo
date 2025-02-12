package hash

import "fmt"

// Хэш-таблица с открытой адресацией (только неотрицательные ключи)
type OpenTable struct {
	data []int // -1 - свободная ячейка, -2 - удаленый ключ
	step int
	h    func(int) int
}

// Создание новой таблицы с открытой адресацией
func NewOpenTable(length int, h func(int) int) *OpenTable {
	if h == nil {
		panic("Nill hash func in NewOpenTable()")
	} else if length < 1 {
		panic("Length is less than 1 in NewOpenTable()")
	}
	step := 3
	for length%step == 0 {
		step++
	}
	table := &OpenTable{make([]int, length), step, h}
	for i := 0; i < len(table.data); i++ {
		table.data[i] = -1
	}
	return table
}

// Находит позицию от хэш-функции и шага
func (t *OpenTable) hashFunc(key, iter int) int {
	if (key < 0) || (iter < 0) {
		panic("Invalid input in hashFunc()")
	}
	pos := (t.h(key) + iter*t.step) % len(t.data)
	return pos
}

// Вставка ключа
func (t *OpenTable) Insert(key int) {
	if key < 0 {
		panic("Key is less than 0 in *OpenTable.Insert()")
	}
	for i := 0; i < len(t.data); i++ {
		pos := t.hashFunc(key, i)
		if (t.data[pos] == -1) || (t.data[pos] == -2) {
			t.data[pos] = key
			return
		}
	}
	panic(fmt.Sprintf("Hash-table owerflow, try to add %d", key))
}

// Поиск позиции по ключу
func (t *OpenTable) SearchPos(key int) int {
	if key < 0 {
		panic("Key is less than 0 in *OpenTable.SearchPos()")
	}
	for i := 0; i < len(t.data); i++ {
		pos := t.hashFunc(key, i)
		if t.data[pos] == key {
			return pos
		} else if t.data[pos] == -1 {
			return -1
		}
	}
	return -1
}

// Содержит ли таблица ключ
func (t *OpenTable) Search(key int) bool {
	if key < 0 {
		panic("Key is less than 0 in *OpenTable.Search()")
	}
	pos := t.SearchPos(key)
	if pos != -1 {
		return true
	} else {
		return false
	}
}

// Удаление по ключу
func (t *OpenTable) Delete(key int) {
	if key < 0 {
		panic("Key is less than 0 in *OpenTable.Delete()")
	}
	pos := t.SearchPos(key)
	if pos != -1 {
		t.data[pos] = -2
	}
}

// Вывод таблицы
func (t *OpenTable) Print() {
	for i, val := range t.data {
		if (val != -1) && (val != -2) {
			fmt.Printf("%d: %d, ", i, val)
		}
	}
	fmt.Println()
}
