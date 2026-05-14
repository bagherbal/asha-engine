package betamatching

import (
	"math"
	"testing"
)

func TestThresholdRepresentationCompletionBetaMatchingAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if !a.ScalarSectorRowConstructed {
		t.Fatalf("expected scalar sector baseline row")
	}
	if math.Abs(a.ScalarSectorDeltaB.B1-0.1) > 1e-10 || math.Abs(a.ScalarSectorDeltaB.B2-(1.0/6.0)) > 1e-10 || math.Abs(a.ScalarSectorDeltaB.B3) > 1e-10 {
		t.Fatalf("unexpected scalar sector Δb: %+v", a.ScalarSectorDeltaB)
	}
	if a.ScalarSectorIsThresholdCorrection {
		t.Fatalf("scalar sector baseline row must not be a heavy threshold correction")
	}
	if a.IncompleteOpenRows != 8 {
		t.Fatalf("expected 8 incomplete open rows, got %d", a.IncompleteOpenRows)
	}
	if a.BGapRepresentationCompleted || a.ContactOverlapRepresentationCompleted || a.AllOpenModesRepresentationComplete {
		t.Fatalf("open modes must remain representation-incomplete")
	}
	if a.BetaCorrectionRowsAllowed != 0 || a.ThresholdCorrectedBetaDerived || a.FullFiniteBetaMatchingTensorDerived {
		t.Fatalf("threshold-corrected beta tensor must remain sealed")
	}
	if a.ResidualNullityAfter != 3 || a.ResidualNullityAfter != a.ResidualNullityBefore {
		t.Fatalf("unexpected residual nullity before=%d after=%d", a.ResidualNullityBefore, a.ResidualNullityAfter)
	}
	if a.HiddenObservedInputUsed || a.PhysicalWeakAngleDerived || a.FineStructureDerived || a.PhysicalMassesDerived {
		t.Fatalf("physical predictions must remain sealed")
	}
}

func TestAmbiguityWitnessesAreCompatibleButNonCanonical(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault: %v", err)
	}
	if len(a.AmbiguityWitnesses) < 3 {
		t.Fatalf("expected ambiguity witnesses")
	}
	zero, nonzero := false, false
	for _, w := range a.AmbiguityWitnesses {
		if !w.CompatibleWithFiniteData {
			t.Fatalf("witness %q should be finite-compatible", w.Name)
		}
		if w.DeltaB.Zero(1e-12) {
			zero = true
		} else {
			nonzero = true
		}
	}
	if !zero || !nonzero {
		t.Fatalf("expected both zero and nonzero compatible Δb witnesses")
	}
}
