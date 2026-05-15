package generation2structuralzero

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Collapse.UniqueMinimal || !a.Collapse.ForcesMiddleZero {
		t.Fatalf("expected unique primitive structural zero: %+v", a.Collapse)
	}
	if !a.Axiom.GeometricallyForced || !a.Axiom.Generation2BareZero {
		t.Fatalf("expected geometrically forced K_gen axiom: %+v", a.Axiom)
	}
	if a.Firewall.NativeFlavorDimAfter != NativeFlavorDim || !a.Firewall.PhysicalMassRequiresBridgeData {
		t.Fatalf("firewall changed incorrectly: %+v", a.Firewall)
	}
}

func TestEnumerationScaleFamily(t *testing.T) {
	e := enumerate(4)
	if len(e.PassingFamilies) != 4 {
		t.Fatalf("expected q=1..4 passing scale family, got %d: %s", len(e.PassingFamilies), FormatCandidates(e.PassingFamilies))
	}
	if len(e.PrimitivePassing) != 1 || !e.PrimitivePassing[0].CanonicalMinimal {
		t.Fatalf("expected exactly one primitive canonical triplet, got: %s", FormatCandidates(e.PrimitivePassing))
	}
	if !e.RejectedDegenerateZero {
		t.Fatalf("degenerate tracial spectrum must be rejected by three-level boundary")
	}
}

func TestTheoremPasses(t *testing.T) {
	res := Generation2StructuralZeroIntersectionSieveTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}

func TestRenderAuditContainsKeyStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusGen2StructuralZeroProved, StatusKGenGeometricallyForcedAxiom, StatusNoYukawaPrediction, "K_gen = diag(-1, 0, 1)"} {
		if !stringsContains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}

func stringsContains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
