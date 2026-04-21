package task2

func FindScndMax(data []int) int {
	if len(data) < 2 {
		return -1
	}

	max := data[0]
	second_max := data[0]

	for _, v := range data {
		if v > max {
			second_max = max
			max = v
		} else if v > second_max && v != max {
			second_max = v
		}
	}

	return second_max
}
