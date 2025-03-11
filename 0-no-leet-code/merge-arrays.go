package __no_leet_code

// смержить массивы в один оптимальным образом (без лишних выделений памяти)

func mergeSlices(arrs ...[]int) []int {
	length := 0
	for _, arr := range arrs {
		length += len(arr)
	}
	result := make([]int, length) // создаем слайс длинной length
	temp := result[:]             // вспомогательный слайс - копия ссылки на result, указывает на тот же массив в памяти, но при этом его границы можно изменять.
	for _, arr := range arrs {
		copy(temp, arr)
		temp = temp[len(arr):] // по сути - перемещаем указатель на начало места, в котрое будем копировать на следующей итерации
	}
	return result
}
