package generation2k7minustrialityintertwinerconstructionr4stabilizationaudit

import (
	"strings"
	"testing"
)

func TestGate955BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.DualSealRequired || a.NativeR3 || a.OfficialLedgerUpdate || a.GenerationCarrierCertified || a.FlavorOrientationCertified || a.IndividualYukawaCertified || a.PhysicalAssignmentCertified {
		t.Fatalf("overclaimed status: %#v", a)
	}
	if a.K7MinusDimension != 3 || a.K7PlusDimension != 4 {
		t.Fatalf("bad K7 polarity dimensions")
	}
	if !a.Action.OrderThree || !a.Action.Nontrivial || !nearly(a.Action.Trace, 0) || !nearly(a.Action.Determinant, 1) || !a.Action.MetricPreserving || a.Action.OrbitSpan != 3 {
		t.Fatalf("abstract C3 action sanity failed: %#v", a.Action)
	}
	if a.Action.CanonicalTrialityInput || a.Action.NativeRestriction {
		t.Fatalf("native triality restriction overclaimed: %#v", a.Action)
	}
	if a.Intertwiner.Certified || !a.Intertwiner.ArbitraryBasisChoiceRequired || !a.Intertwiner.UsesR3RowsAsGenerationLabels {
		t.Fatalf("intertwiner obstruction flags wrong: %#v", a.Intertwiner)
	}
	if a.Intertwiner.UsesFlavorBacksolve || a.Intertwiner.UsesObservedYukawaOrMassData || a.Intertwiner.UsesCKMPMNSInput {
		t.Fatalf("forbidden empirical input used: %#v", a.Intertwiner)
	}
	if len(a.Items) != 8 {
		t.Fatalf("expected 8 items, got %d", len(a.Items))
	}
	joined := strings.Join(appendAll(a.Supports, a.Failures, ItemSupports(a.Items), ItemFailures(a.Items)), "\n")
	for _, want := range append(RequiredSupports(), RequiredFailures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}

func TestGate955Theorem(t *testing.T) {
	res := Generation2K7MinusTrialityIntertwinerConstructionR4StabilizationAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{Verdict, Classification, ShortStatus, NextGate, "K7MinusTrialityTracebodyIntertwiner", "FAILED_ROUTE_NO_K7_MINUS_TRIALITY_TO_R3_TRACEBODY_INTERTWINER"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
