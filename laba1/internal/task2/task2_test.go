package task2

import (
	"laba1/internal/common"
	"testing"
	"time"
)

func TestFindScndMax(t *testing.T) {
	data := []int{100, 1000, 5000, 10000}
	mx := FindScndMax(data)

	if mx != 5000 {
		t.Errorf("Want 5000 got %v", mx)
	}
}

func TestFindScndMaxTime(t *testing.T) {
	sizes := []int{100, 1000, 5000, 10000}
	times := make([]time.Duration, len(sizes))
	for i, size := range sizes {
		data := common.GenerateData(size)
		time := common.MeasureTime(func() {
			FindScndMax(data)
		})
		times[i] = time
	}
	res := common.GenerateTable(times, sizes)
	t.Logf("Task2:\n%s", res)
}
