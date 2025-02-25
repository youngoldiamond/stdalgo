package sort

// Соритировка вставками
func Insertion(data []int) {
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1
		for (j >= 0) && (data[j] > key) {
			//ans[j+1] = ans[j]
			j--
		}
		copy(data[j+2:i+1], data[j+1:i])
		data[j+1] = key
	}
}

// Сортировка слиянием вариант без рекурсии
func Merge(data []int) {
	tmp := make([]int, len(data))
	for length := 2; ; length *= 2 {
		copy(tmp, data)
		for i := 0; i < len(data); i += length {
			l := i
			r := i + length/2
			if r >= len(data) {
				break
			}
			for j := i; (j < i+length) && (j < len(data)); j++ {
				if tmp[l] <= tmp[r] {
					data[j] = tmp[l]
					l++
					if l >= i+length/2 {
						break
					}
				} else {
					data[j] = tmp[r]
					r++
					if r >= i+length {
						copy(data[j+1:i+length], tmp[l:i+length/2])
						break
					} else if r >= len(data) {
						copy(data[j+1:], tmp[l:i+length/2])
						break
					}
				}
			}
		}
		if length >= len(data) {
			break
		}
	}
}

// Сортировка слиянием вариант с рекурсией
func MergeRec(data []int) {
	if len(data) == 1 {
		return
	}
	MergeRec(data[:len(data)/2])
	MergeRec(data[len(data)/2:])
	tmp := make([]int, len(data))
	copy(tmp, data)
	l := 0
	r := len(data) / 2
	for i := range data {
		if tmp[l] <= tmp[r] {
			data[i] = tmp[l]
			l++
			if l == len(data)/2 {
				break
			}
		} else {
			data[i] = tmp[r]
			r++
			if r == len(data) {
				copy(data[i+1:], tmp[l:len(data)/2])
				break
			}
		}
	}
}

// Быстрая сортировка с рекурсией
func Quick(data []int) {
	if len(data) <= 1 {
		return
	}
	key := data[len(data)-1]
	midle := -1
	for i := range data {
		if data[i] <= key {
			midle++
			if i != midle {
				data[midle] ^= data[i]
				data[i] ^= data[midle]
				data[midle] ^= data[i]
			}
		}
	}
	Quick(data[:midle])
	Quick(data[midle+1:])
}
