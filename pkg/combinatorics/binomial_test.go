package combinatorics

import "testing"

func TestGradeDimensionsR8(t *testing.T) {
	dims, err := GradeDimensions(8)
	if err != nil {
		t.Fatal(err)
	}
	want := []int{1, 8, 28, 56, 70, 56, 28, 8, 1}
	if len(dims) != len(want) {
		t.Fatalf("len got %d want %d", len(dims), len(want))
	}
	for i := range want {
		if dims[i] != want[i] {
			t.Fatalf("grade %d got %d want %d", i, dims[i], want[i])
		}
	}
	if Sum(dims) != 256 {
		t.Fatalf("sum got %d want 256", Sum(dims))
	}
}

func TestSubsets(t *testing.T) {
	subs, err := Subsets(4, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(subs) != 6 {
		t.Fatalf("got %d subsets want 6", len(subs))
	}
}
