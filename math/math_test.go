package math_test

import (
	"testing"

	"github.com/artsupawat/buggy-math-go/math"
)

func TestAdd(t *testing.T) {
	// Arrange & Act
	got := math.Add(1, 5)

	// Assert
	want := 6
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestSubtract(t *testing.T) {
	// Arrange & Act
	got := math.Subtract(5, 3)

	// Assert
	want := 2
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
