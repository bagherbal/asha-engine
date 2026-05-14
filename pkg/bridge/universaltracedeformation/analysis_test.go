package universaltracedeformation

import (
	"math"
	"testing"
)

func TestUniversalBetaBoundaryOffsetEquivalence(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	eq := a.Equivalence
	if !eq.UniversalBetaShiftEquivalentToBoundaryOffset || !eq.RelativeRunningUnaffectedByUniversalRow || !eq.BoundaryOffsetActsOnlyAsCommonIntercept || !eq.SignConventionChecked || eq.PhysicalPredictionClaim {
		t.Fatalf("bad equivalence audit: %s", FormatEquivalence(eq))
	}
}

func TestRequiredOffsetsComeOnlyFromConditionalGate201Shapes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.RequiredOffsets) != 2 {
		t.Fatalf("expected two Gate-201 conditional shape offsets, got %d: %s", len(a.RequiredOffsets), FormatRequiredOffsets(a.RequiredOffsets))
	}
	for _, r := range a.RequiredOffsets {
		recomputed := r.UniversalBetaDelta * r.LeverArmLog / (8 * math.Pi * math.Pi)
		if math.Abs(recomputed-r.RequiredDeltaU) > 1e-12 {
			t.Fatalf("delta formula mismatch: %s recomputed=%g", FormatRequiredOffset(r), recomputed)
		}
		if !r.FromGate201 || !r.ConditionalOnly || r.FiniteDerived || r.RequiredDeltaU <= 0 || r.DefectAdjustedU <= 1 {
			t.Fatalf("offset overclaimed or invalid: %s", FormatRequiredOffset(r))
		}
	}
}

func TestFiniteTraceCandidatesDoNotCanonicalizeOffset(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	ft := a.FiniteTrace
	if ft.BGapValue <= 0 || ft.ContactZetaValues != 5 || ft.ContactActionCandidates != 10 || ft.CandidatesAudited != len(a.TraceCandidates) {
		t.Fatalf("bad finite trace audit: %s", FormatFiniteTrace(ft))
	}
	if ft.CanonicalBoundaryOffsetCandidates != 0 || ft.CanonicalPerfectAbsorptions != 0 || ft.UniversalVolumeDefectCanonicalized || ft.BGapCanonicalOffsetDerived || ft.ContactZetaCanonicalOffsetDerived {
		t.Fatalf("finite trace offset was overclaimed: %s", FormatFiniteTrace(ft))
	}
	if len(a.AbsorptionTests) == 0 {
		t.Fatalf("expected absorption tests")
	}
	for _, x := range a.AbsorptionTests {
		if x.ConditionalBridgeAllowed || (x.PerfectlyAbsorbs && x.CandidateCanonical) {
			t.Fatalf("unexpected canonical absorption: %s", FormatAbsorptionTest(x))
		}
	}
}

func TestFirewallSealed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Firewall
	if f.DefectAdjustedBoundaryDerived || f.BGapUsedAsPhysicalMass || f.BGapUsedAsBetaRow || f.ContactZetaUsedAsBetaRow || f.ArbitraryCoefficientInserted || f.PhysicalUnificationClaimed || f.ThresholdCorrectedPhysicalFitClaimed || f.AbsoluteMassPredicted || f.FiniteMatchingCorrectionsDerived {
		t.Fatalf("firewall leak: %s", FormatFirewall(f))
	}
	if f.StrictNullityBefore != f.StrictNullityAfter || f.PhysicalPredictionNullityBefore != f.PhysicalPredictionNullityAfter {
		t.Fatalf("nullity changed unexpectedly: %s", FormatFirewall(f))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := UniversalTraceDeformationTopologicalBoundaryOffsetAuditTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
