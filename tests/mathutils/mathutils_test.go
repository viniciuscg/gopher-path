package mathutils

import (
	"testing"
)

// Basic test for Add
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

// Table-driven tests for Add
func TestAddTableDriven(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 1, 2},
		{2, -2, 0},
		{0, 0, 0},
		{-3, -3, -6},
	}
	for _, tt := range tests {
		t.Run("Add", func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}

// Test Divide with normal input
func TestDivide(t *testing.T) {
	result := Divide(10, 2)
	expected := 5
	if result != expected {
		t.Errorf("Divide(10, 2) = %d; want %d", result, expected)
	}
}

// Test Divide with division by zero — expects panic
func TestDivideByZero(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Divide should panic when dividing by zero")
		}
	}()
	_ = Divide(10, 0)
}

// Skipping a test conditionally (e.g. integration, network, etc.)
func TestSkipExample(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test in short mode")
	}
	// long-running test logic here
}

// Benchmark test
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Add(100, 200)
	}
}
