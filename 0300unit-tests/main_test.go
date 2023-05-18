package main

import (
	"fmt"
	"testing"
)

func TestIntMin(t *testing.T) {
	min := IntMin(0, 2)
	if min != 0 {
		t.Errorf("InitMin(0,2) = %d; wanted 0", min)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b   int
		wanted int
	}{
		{1, 2, 1},
		{2, 1, 1},
		{2, 2, 2},
		{0, 2, 0},
		{0, -2, -2},
	}
	for _, data := range tests {
		name := fmt.Sprintf("InitMin(%d,%d)", data.a, data.b)
		t.Run(name, func(t *testing.T) {
			min := IntMin(data.a, data.b)
			if min != data.wanted {
				t.Errorf("InitMin(%d,%d) = %d; wanted %d",
					data.a, data.b, min, data.wanted)
			}
		})
	}
}

func BenchmarkIntMin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntMin(0, 0)
	}
}
