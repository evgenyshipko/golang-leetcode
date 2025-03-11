package __no_leet_code

import "math/rand"

// сгенерировать случайную матрицу m*n и чтобы числа в ней не повторялись

func generateMatrix(n, m int) [][]int {
	used := map[int]bool{}

	rows := make([][]int, n)

	for range n {
		row := make([]int, m)
		for range m {
			var random int
			for {
				random := rand.Int()
				if !used[random] {
					used[random] = true
					break
				}
			}
			row = append(row, random)
		}
		rows = append(rows, row)
	}
	return rows
}
