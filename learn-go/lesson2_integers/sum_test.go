package integers

import "testing"

func TestSum(t *testing.T) {
	sum := Sum(2,2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected '%d' but got '%d'",expected,sum)
	}
}