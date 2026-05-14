package clifford

import "testing"

func TestCliffordDimension(t *testing.T) {
	alg, err := New(Signature{Positive: 1, Negative: 7})
	if err != nil {
		t.Fatal(err)
	}
	if alg.AlgebraDimension() != 256 {
		t.Fatalf("got %d want 256", alg.AlgebraDimension())
	}
	if alg.AnticommutatorCoefficient(0, 0) != 2 {
		t.Fatal("expected positive time-like generator")
	}
	if alg.AnticommutatorCoefficient(7, 7) != -2 {
		t.Fatal("expected negative space-like generator")
	}
}
