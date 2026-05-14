package matter

import "testing"

func TestFockContactBridge(t *testing.T) {
	b, err := BuildDefaultFockContactBridge()
	if err != nil {
		t.Fatal(err)
	}
	if b.FockStateCount != 16 {
		t.Fatalf("expected 16 Fock states, got %d", b.FockStateCount)
	}
	if !b.ModeToActiveScalarMatch {
		t.Fatalf("expected Fock modes to match active scalar directions")
	}
	if !b.SpatialToProtectedMatch {
		t.Fatalf("expected 3 spatial modes to resonate with 3 protected contact directions")
	}
	if b.YukawaOperatorConstructed {
		t.Fatalf("Yukawa operator should remain open in this gate")
	}
}
