package inversebsectordeformation

import "testing"

func TestInverseFamilyRejectsSingleScaleUniqueness(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.InverseFamily
	if !f.CanEvaluateIfBoundaryScaleSealed || !f.MismatchTriangleClosedByConstruction || !f.UOneBoundaryEnforcedByConstruction || f.PhysicalPredictionClaim {
		t.Fatalf("bad inverse family: %s", FormatInverseFamily(f))
	}
	if !f.SingleThresholdScaleOnly || !f.BoundaryScaleStillFree || !f.UnderdeterminedByOneContinuousParameter {
		t.Fatalf("expected single-scale no-go: %s", FormatInverseFamily(f))
	}
}

func TestBenchmarkFormulaClosesUOne(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	p := a.BenchmarkPoint
	if !p.ValidOrderedScales || p.MaxAbsResidual > 1e-8 || p.TriangleArea != 0 {
		t.Fatalf("benchmark did not close by construction: %s", FormatPoint(p))
	}
}

func TestKnownRawRepresentationRowsNoGo(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	r := a.Representation
	if r.KnownRationalRowsAudited == 0 || r.RawExactKnownRepresentationFound || len(r.RawNoGoMatches) != r.KnownRationalRowsAudited || r.PhysicalRepresentationClaimed {
		t.Fatalf("expected raw representation no-go: %s", FormatRepresentation(r))
	}
}

func TestUniversalCompletionIsConditionalOnly(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	r := a.Representation
	if !r.ConditionalUniversalShapeMatchFound || len(r.UniversalCompletionMatches) == 0 {
		t.Fatalf("expected conditional universal-completion shape matches: %s", FormatRepresentation(r))
	}
	if r.UniversalCompletionFiniteDerived || r.IntegerOrRationalTotalDeltaDerived || r.PhysicalRepresentationClaimed {
		t.Fatalf("universal completion overclaimed: %s", FormatRepresentation(r))
	}
	for _, m := range r.UniversalCompletionMatches {
		if !m.ConditionalAlive || !m.PositiveOrderedScales || !m.NonnegativeUniversalRow || m.FiniteDerived || m.MaxAbsResidual > 1e-7 {
			t.Fatalf("bad universal match: %s", FormatUniversalMatch(m))
		}
	}
}

func TestInternalSpectralFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	i := a.Internal
	if !i.DimensionlessSpectralAnchorsKnown || i.BGapValue <= 0 || i.BPositiveModeCount == 0 || i.ContactPartialModeCount != 7 || i.ScalarActiveModeCount != 4 {
		t.Fatalf("bad finite spectral inventory: %s", FormatInternal(i))
	}
	if i.BGapHasRepresentationRow || i.ContactModesHaveRepresentationRows || i.ThresholdActivationRuleDerived || i.FiniteToContinuumMatchingDerived || i.StructuralBGapMatchFound || i.StructuralContactMatchFound || i.CountResonancePromoted {
		t.Fatalf("internal spectral firewall leaked: %s", FormatInternal(i))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := InverseBSectorDeformationThresholdPredictionAuditTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
