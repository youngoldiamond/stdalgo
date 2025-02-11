package hash

import "math/rand/v2"

type Table interface {
	Insert(key int)
	Search(key int) bool
	Delete(key int)
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
	const gr = 0.618033989
	return func(key int) int {
		return int(float64(len) * (float64(key)*gr - float64(int(float64(key)*gr))))
	}
}

// Создаёт функцию из универсального множества хэш функций с ключами до 1^31
func UniFunc(len int) func(int) int {
	if len < 1 {
		panic("Length is less than 1 in UniFunc()")
	}
	const maxInt32 = (1 << 31) - 1
	a := rand.Int64N(maxInt32-1) + 1
	b := a
	for a == b {
		b = rand.Int64N(maxInt32)
	}
	return func(key int) int { return int((a*int64(key)+b)%(maxInt32+1)) % len }
}
