package mathoperations

import "testing"

func TestSum(t *testing.T) {
	input1 := 10.0
	input2 := 20.0

	expected := 30.0

	got := Sum(input1, input2)

	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
