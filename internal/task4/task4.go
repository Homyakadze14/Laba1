package task4

func MultTable(n int) [][]int {
	table := make([][]int, n)
	for i := range n {
		table[i] = make([]int, n)
		for j := range n {
			if i-1 < 0 {
				table[i][j] = j + 1
			} else if j-1 < 0 {
				table[i][j] = i + 1
			} else {
				table[i][j] = table[0][j] * table[i][0]
			}
		}
	}
	return table
}
