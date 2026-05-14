package linear

import "testing"

func TestIdentityMultiplication(t *testing.T) {
	a := NewMatrix(2, 2)
	a.Set(0, 0, 2)
	a.Set(0, 1, 3)
	a.Set(1, 0, 5)
	a.Set(1, 1, 7)

	got, err := Identity(2).Mul(a)
	if err != nil {
		t.Fatal(err)
	}
	if !got.AlmostEqual(a, 1e-12) {
		t.Fatal("I*A should equal A")
	}
}
