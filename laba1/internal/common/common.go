package common

import (
	"fmt"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

func MeasureTime(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}

func MeasureMemoryUsage(fn func()) uint64 {
	runtime.GC()

	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	before := memStats.Alloc

	fn()

	runtime.ReadMemStats(&memStats)
	after := memStats.Alloc
	used := after - before

	return used
}

func GenerateData(size int) []int {
	data := make([]int, size)
	for i := range size {
		data[i] = rand.Int()
	}
	return data
}

func GenerateTable(times []time.Duration, sizes []int) string {
	var result strings.Builder
	result.WriteString("| Size | Time |\n|------|-----------|\n")
	for i, size := range sizes {
		result.WriteString(fmt.Sprintf("| %d | %v |\n", size, times[i]))
	}
	return result.String()
}

func GenerateMemTable(times []uint64, sizes []int) string {
	var result strings.Builder
	result.WriteString("| Size | Mem (byte) |\n|------|-----------|\n")
	for i, size := range sizes {
		result.WriteString(fmt.Sprintf("| %d | %v |\n", size, times[i]))
	}
	return result.String()
}
