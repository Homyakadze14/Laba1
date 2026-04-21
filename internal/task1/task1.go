package task1

func Exists(data []int, target int) bool {
	for _, v := range data {
		if v == target {
			return true
		}
	}
	return false
}
