package main

import (
	"testing"
)

// Calculate returns x + 2.
func Calculate(x int) (result int) {
	result = x + 2
	return result
}
func benchmarkCalculate(input int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Calculate(input)
	}
}

// go test -v -run=BenchmarkExportTyreSickerExcelWithGoRoutine -bench=. -benchmem -count 5 eco-tries-project-back-end/controllers/report/sticker_report

func BenchmarkCalculate100(b *testing.B)         { benchmarkCalculate(100, b) }
func BenchmarkCalculateNegative100(b *testing.B) { benchmarkCalculate(-100, b) }
func BenchmarkCalculateNegative1(b *testing.B)   { benchmarkCalculate(-1, b) }
