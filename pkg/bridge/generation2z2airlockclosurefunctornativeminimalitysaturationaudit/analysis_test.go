package generation2z2airlockclosurefunctornativeminimalitysaturationaudit

import (
	"strings"
	"testing"
)

func TestGate928ClosureAxiomsForceLevelsUnderAxioms(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited != Gate927ShortStatus {
		t.Fatalf("bad inherited status: %s", a.Inherited)
	}
	if !a.Basepoint.ForcedByAxiom || !a.Basepoint.EmptyHasNoActivation || !a.Basepoint.MatchesReducedForm || a.Basepoint.NativeAxiomCertified || a.Basepoint.ClosureTarget != "F_0" {
		t.Fatalf("bad basepoint: %s", FormatBasepoint(a.Basepoint))
	}
	if !a.Monotonicity.Monotone || !a.Monotonicity.MatchingThreeLevels || a.Monotonicity.NativeActionTheorem {
		t.Fatalf("bad monotonicity: %s", FormatMonotonicity(a.Monotonicity))
	}
	if !a.Minimality.ForcedByMinimality || a.Minimality.SkipsSaturation || !a.Minimality.FirstNonbaseClosure || a.Minimality.NativeAxiomCertified || a.Minimality.ClosureTarget != "F_1" {
		t.Fatalf("bad minimality: %s", FormatMinimality(a.Minimality))
	}
	if !a.Saturation.ForcedBySaturation || !a.Saturation.TopBoundaryDegree || !a.Saturation.CannotRemainAtF1 || a.Saturation.NativeAxiomCertified || a.Saturation.ClosureTarget != "F_2" {
		t.Fatalf("bad saturation: %s", FormatSaturation(a.Saturation))
	}
}

func TestGate928Z2FixedBaseUniquenessAndMeasure(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Z2.PhaseFlipCommutes || !a.Z2.RanksInvariant || !a.Z2.RepresentativeFree || a.Z2.NativeZ2Theorem {
		t.Fatalf("bad z2: %s", FormatZ2(a.Z2))
	}
	if !a.FixedBase.ForcesCumulative || !a.FixedBase.RejectsAssociated || a.FixedBase.NativeRuleCertified || a.FixedBase.TopTarget != "[F_2/F_0]_{Z2}" {
		t.Fatalf("bad fixed base: %s", FormatFixedBase(a.FixedBase))
	}
	if !a.Uniqueness.UniqueUnderAxioms || !a.Uniqueness.Basepoint || !a.Uniqueness.Monotone || !a.Uniqueness.Minimal || !a.Uniqueness.Saturated || !a.Uniqueness.Z2Invariant || a.Uniqueness.NativeAxiomSource {
		t.Fatalf("bad uniqueness: %s", FormatUniqueness(a.Uniqueness))
	}
	if !a.Measure.UniqueClosureSuppliesTheta || !a.Measure.TargetsFixedByClosure || !a.Measure.AlphaReconstructed || a.Measure.NativeAlphaTheorem || a.Measure.ThetaOneRank != RankF1OverF0 || a.Measure.ThetaTwoRank != RankF2OverF0 {
		t.Fatalf("bad measure: %s", FormatMeasure(a.Measure))
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("bad firewalls: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate928Theorem(t *testing.T) {
	res := Generation2Z2AirlockClosureFunctorNativeMinimalitySaturationAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range append(append(Statuses(), Supports()...), Failures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
	for _, want := range []string{FinalTruth, Classification, ShortStatus, BoundarySubsetChain, ExteriorSourceChain, AirlockFlagChain, ClosureFunctor, ClosureZero, ClosureOne, ClosureTwo, ThetaViaClosure, ThetaOne, ThetaTwo, AlphaViaClosure, NextGate} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}
