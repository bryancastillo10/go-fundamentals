package iteration

import "testing"

func TestRepeat(t *testing.T){
	repeated := Repeat("b")
	expected := "bbbbb"

	if repeated != expected{
		t.Errorf("Expected %q but got %q",expected,repeated)
	}
}

func BenchMarkRepeat(b *testing.B){
	for i :=0; i < b.N; i++{
		Repeat("b")
	}
}