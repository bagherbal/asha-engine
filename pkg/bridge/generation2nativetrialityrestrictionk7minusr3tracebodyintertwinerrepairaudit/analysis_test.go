package generation2nativetrialityrestrictionk7minusr3tracebodyintertwinerrepairaudit

import (
	"strings"
	"testing"
)

func TestGate956BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.R3DualSealRequired || a.NativeR3 || a.OfficialLedgerUpdate || a.GenerationCarrierCertified || a.FlavorOrientationCertified || a.IndividualYukawaCertified || a.PhysicalAssignmentCertified {
		t.Fatalf("overclaimed status: %#v", a)
	}
	if a.K7MinusDimension != 3 || a.K7PlusDimension != 4 || a.R3TraceRows != TraceRows {
		t.Fatalf("bad inherited carrier data")
	}
	if a.Transport.NativeTrialityOperatorConstructed || a.Transport.TransportToLambda4Certified || a.Transport.PreservesK7ContactCarrier || a.Transport.PreservesK7Minus || a.Transport.AbstractC3FromGate955Realized {
		t.Fatalf("native triality route overclaimed: %#v", a.Transport)
	}
	if !a.Transport.LeakageToK7PlusUnknown {
		t.Fatalf("K7+ leakage should remain unknown without native operator")
	}
	if a.Intertwiner.Certified || !a.Intertwiner.RequiresNativeK7MinusAction || !a.Intertwiner.UsesArbitraryBasisFit || !a.Intertwiner.UsesR3RowsAsGenerationLabels || !a.Intertwiner.PreservesR3DualSeal {
		t.Fatalf("intertwiner obstruction flags wrong: %#v", a.Intertwiner)
	}
	if a.Intertwiner.UsesFlavorBacksolve || a.Intertwiner.UsesObservedMassesOrYukawas || a.Intertwiner.UsesCKMPMNSInput {
		t.Fatalf("forbidden empirical/flavor input used: %#v", a.Intertwiner)
	}
	if len(a.Items) != 7 {
		t.Fatalf("expected 7 items, got %d", len(a.Items))
	}
	joined := strings.Join(appendAll(a.Supports, a.Failures, ItemSupports(a.Items), ItemFailures(a.Items)), "\n")
	for _, want := range append(RequiredSupports(), RequiredFailures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}

func TestGate956Theorem(t *testing.T) {
	res := Generation2NativeTrialityRestrictionK7MinusR3TracebodyIntertwinerRepairAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{Verdict, Classification, ShortStatus, NextGate, "FAILED_ROUTE_NO_TRIALITY_TO_LAMBDA4_TRANSPORT_MAP", "FAILED_ROUTE_NO_NATIVE_K7_MINUS_TRIALITY_R3_TRACEBODY_INTERTWINER"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
