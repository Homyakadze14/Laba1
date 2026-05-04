package task4

import (
	"laba1/internal/common"
	"testing"
	"time"
)

func TestMultTable(t *testing.T) {
	arr := MultTable(3)
	for _, row := range arr {
		for _, val := range row {
			print(val, " ")
		}
		println()
	}
}

func TestMultTableTime(t *testing.T) {
	sizes := []int{100, 1000, 5000, 10000}
	times := make([]time.Duration, len(sizes))
	for i, size := range sizes {
		time := common.MeasureTime(func() {
			MultTable(size)
		})
		times[i] = time
	}
	res := common.GenerateTable(times, sizes)
	t.Logf("Task4:\n%s", res)
}
