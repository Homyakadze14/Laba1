package task1

import (
	"laba1/internal/common"
	"math/rand/v2"
	"testing"
	"time"
)

func TestExists(t *testing.T) {
	data := []int{100, 1000, 5000, 10000}
	ex := Exists(data, -1)
	if ex {
		t.Errorf("Want false got %v", ex)
	}

	ex = Exists(data, 5000)
	if !ex {
		t.Errorf("Want true got %v", ex)
	}
}

func TestExistsTime(t *testing.T) {
	sizes := []int{100, 1000, 5000, 10000}
	times := make([]time.Duration, len(sizes))
	for i, size := range sizes {
		data := common.GenerateData(size)
		time := common.MeasureTime(func() {
			Exists(data, rand.Int())
		})
		times[i] = time
	}
	res := common.GenerateTable(times, sizes)
	t.Logf("Task1:\n%s", res)
}
