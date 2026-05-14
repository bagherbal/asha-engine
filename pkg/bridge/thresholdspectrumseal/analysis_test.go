package thresholdspectrumseal

import "testing"

func TestGate213BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Summary.Status != StatusConditionalSealPreflight {
		t.Fatalf("unexpected status: %s", FormatSummary(a.Summary))
	}
	if !a.Summary.ThresholdSpectrumSealIntroduced || !a.Summary.SelectedBestPair {
		t.Fatalf("seal/subject not established: %s", FormatSummary(a.Summary))
	}
	if a.Subject.Row1Rep != "(1,3,Y=1)" || a.Subject.Row2Rep != "(8,2,Y=1/2)" {
		t.Fatalf("unexpected sealed subject: %s", FormatSubject(a.Subject))
	}
}

func TestMatchingCorrectionObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Matching.ThresholdMatchingCoefficientsDerived || a.Matching.CanonicalSubtractionSchemeDerived || a.Matching.MSbarOrDimRegImported {
		t.Fatalf("matching firewall leak: %s", FormatMatching(a.Matching))
	}
	if a.Matching.Status != MatchingCorrectionsFailed {
		t.Fatalf("expected matching failed route: %s", FormatMatching(a.Matching))
	}
}

func TestExactTwoLoopHeavyMatrix(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	m := a.TwoLoop.HeavyInducedTwoLoopMatrix
	checks := map[[2]int]string{
		{0, 0}: "144/25",
		{0, 1}: "108/5",
		{0, 2}: "144/5",
		{1, 0}: "36/5",
		{1, 1}: "108",
		{1, 2}: "48",
		{2, 0}: "18/5",
		{2, 1}: "18",
		{2, 2}: "192",
	}
	for ij, want := range checks {
		if got := m.M[ij[0]][ij[1]].String(); got != want {
			t.Fatalf("heavy two-loop matrix[%d,%d]=%s, want %s; matrix=%s", ij[0], ij[1], got, want, m.String())
		}
	}
	if a.TwoLoop.ImportedAsFiniteCore || !a.TwoLoop.UsesStandardQFTFormula {
		t.Fatalf("two-loop provenance invalid: %s", FormatTwoLoop(a.TwoLoop))
	}
}

func TestTwoLoopWarningAndFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Stability.Status != TwoLoopWarning || !a.Stability.OneLoopScalesValidOnlyAtOneLoop || !a.Stability.RequiresMatchingCorrections {
		t.Fatalf("unexpected stability status: %s", FormatStability(a.Stability))
	}
	if a.Stability.MaxTwoLoopToOneLoopRatio < 1.0 {
		t.Fatalf("expected non-small two-loop warning ratio, got %s", FormatStability(a.Stability))
	}
	f := a.Firewall
	if f.UniquePhysicalSpectrumClaimed || f.MatchingCorrectionsDerived || f.TwoLoopCoefficientsFiniteDerived || f.TwoLoopScalesClaimedAsPrediction || f.PhysicalPredictionClaimed || f.ProtonLifetimeComputed {
		t.Fatalf("firewall leak: %s", FormatFirewall(f))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := ThresholdSpectrumSealMatchingCorrectionTwoLoopPreflightTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
