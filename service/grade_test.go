package service_test

import (
	"UnitTest/service"
	"fmt"
	"testing"
)

// Unit test for CheckGrade function
func TestCheckGrade(t *testing.T) {
	testCases := []struct {
		name     string
		score    int
		expected string
	}{
		{"A+ grade", 100, "A+"},
		{"A grade", 90, "A"},
		{"B grade", 80, "B"},
		{"C grade", 70, "C"},
		{"D grade", 60, "D"},
		{"F grade", 50, "F"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := service.CheckGrade(tc.score)
			if actual != tc.expected {
				t.Errorf("expected: %s, actual: %s", tc.expected, actual)
			}
		})
	}
}

// Benchmark test for CheckGrade function
func BenchmarkCheckGrade(b *testing.B) {
	for i := 0; i < b.N; i++ {
		service.CheckGrade(100)
	}
}

func ExampleCheckGrade() {
	service.CheckGrade(100)
	fmt.Println(service.CheckGrade(100))
	// Output:
	// A+
}
