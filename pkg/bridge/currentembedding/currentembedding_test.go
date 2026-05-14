package currentembedding

import "testing"

func TestCurrentFieldEmbedding(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault() error: %v", err)
	}
	if a.SectorFieldCount != 4 {
		t.Fatalf("sectors=%d, want 4", a.SectorFieldCount)
	}
	if a.GeneratorFieldCount != 16 {
		t.Fatalf("generator fields=%d, want 16", a.GeneratorFieldCount)
	}
	if !a.FieldSlotsDefined {
		t.Fatalf("field slots should be typed")
	}
	if a.CurrentToContactEmbeddingDerived {
		t.Fatalf("current-to-contact embedding should remain open")
	}
	if a.HessianComputable || a.CurrentHessianDerived || a.PropagatorRuleDerived {
		t.Fatalf("Hessian/propagator must remain open until embedding/source are derived")
	}
	if a.HiddenObservedInputUsed {
		t.Fatalf("must not use observed inputs")
	}
}
