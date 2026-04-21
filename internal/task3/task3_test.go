package task3

import (
	"laba1/internal/common"
	"math/rand/v2"
	"testing"
	"time"
)

func TestBinarySearch(t *testing.T) {
	data := []int{100, 1000, 5000, 10000}
	res := BinarySearch(data, 3)
	if res != -1 {
		t.Errorf("Expected -1 got %v", res)
	}

	res = BinarySearch(data, 5000)
	if res != 2 {
		t.Errorf("Expected 2 got %v", res)
	}
}

func TestBinarySearchTime(t *testing.T) {
	sizes := []int{100, 1000, 5000, 10000}
	times := make([]time.Duration, len(sizes))
	for i, size := range sizes {
		data := common.GenerateData(size)
		time := common.MeasureTime(func() {
			BinarySearch(data, rand.Int())
		})
		times[i] = time
	}
	res := common.GenerateTable(times, sizes)
	t.Logf("Task3:\n%s", res)
}
