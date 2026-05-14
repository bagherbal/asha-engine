package cliffordpullback

import "testing"

func TestBuildDefaultCliffordPullback(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.CliffordAction.CliffordMultiplicationAvailable || a.CliffordAction.RealSpinorDimension != 32 || a.CliffordAction.ComplexSpinorDimension != 16 {
		t.Fatalf("bad Clifford action audit: %s", FormatCliffordAction(a.CliffordAction))
	}
	if a.TauEtaPullback.CliffordActionApplicable || a.TauEtaPullback.EndomorphismConstructed || !a.TauEtaPullback.HypotheticalOperatorRejected {
		t.Fatalf("tau_eta should remain outside Clifford domain: %s", FormatTauEtaPullback(a.TauEtaPullback))
	}
	if a.Functor.PullbackFunctorDerived || a.Summary.EndomorphismConstructed {
		t.Fatalf("pullback should not be derived: %s :: %s", FormatFunctor(a.Functor), FormatSummary(a.Summary))
	}
}

func TestSpatialAndTrialityRemainConditional(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Spatial.TauMagnitudeSelectorCapacity || a.Spatial.WeakPlaneDerived || a.Spatial.S3DegeneracyBroken {
		t.Fatalf("spatial sieve leaked a selection: %s", FormatSpatial(a.Spatial))
	}
	if a.Spatial.UniqueAxisConditionallySeen != "a†_3" || a.Spatial.ComplementPlaneConditionally != "U={a†_1,a†_2}" {
		t.Fatalf("unexpected conditional weak-plane roadmap: %s", FormatSpatial(a.Spatial))
	}
	if !a.Triality.DistinctEigenvalueCapacity || a.Triality.GenerationTextureDerived || a.Triality.DiagonalGenerationOperator {
		t.Fatalf("triality sieve should be capacity only: %s", FormatTriality(a.Triality))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := CliffordActionPullbackTauEtaEndomorphismAuditTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
