package quicksort

import (
	"laba1/internal/common"
	"slices"
	"testing"
	"time"
)

func TestSort(t *testing.T) {
	arr := []int{1, -23, 123, 32, 4}
	needArr := []int{-23, 1, 4, 32, 123}

	Sort(arr, 0, len(arr)-1)

	if !slices.Equal(arr, needArr) {
		t.Errorf("expected %v, got %v", needArr, arr)
	}
}

func TestSortTime(t *testing.T) {
	sizes := []int{100, 1000, 5000, 10000}
	times := make([]time.Duration, len(sizes))
	for i, size := range sizes {
		data := common.GenerateData(size)
		time := common.MeasureTime(func() {
			Sort(data, 0, len(data)-1)
		})
		times[i] = time
	}
	res := common.GenerateTable(times, sizes)
	t.Logf("QuickSort:\n%s", res)
}

func TestSortSize(t *testing.T) {
	sizes := []int{100, 1000, 5000, 10000}
	usages := make([]uint64, len(sizes))
	for i, size := range sizes {
		data := common.GenerateData(size)
		usage := common.MeasureMemoryUsage(func() {
			Sort(data, 0, len(data)-1)
		})
		usages[i] = usage
	}
	res := common.GenerateMemTable(usages, sizes)
	t.Logf("QuickSort:\n%s", res)
}
