package thresholdorigindichotomy

import "testing"

func TestBuildDefaultThresholdOriginDichotomy(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Dichotomy.PreviousGate178NoGo || !a.PreviousGate178.Requirements.NoCandidateHasAllPieces {
		t.Fatalf("expected Gate 178 no-go input: %s", FormatDichotomy(a.Dichotomy))
	}
	if len(a.Branches) != 4 {
		t.Fatalf("expected four origin branches, got %d", len(a.Branches))
	}
}

func TestContinuumBridgeRequiredNotDerived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	c := a.Continuum
	if len(c.CandidateExistingAnchors) < 4 {
		t.Fatalf("expected existing finite anchors: %s", FormatContinuum(c))
	}
	if !c.OrientedFourCycleRequired || !c.PrincipalBundleRequired || !c.ChernWeilNormalizationRequired || !c.ContinuumTraceConventionRequired || !c.LocalFieldMapRequired || !c.PhysicalMassUnitRequired || !c.ActivationPredicateRequired || !c.DecouplingLawRequired || !c.GaugeRepresentationRowsRequired {
		t.Fatalf("continuum bridge requirement inventory incomplete: %s", FormatContinuum(c))
	}
	if c.AllRequiredObjectsPresent || c.BridgeDerived || c.CanPromoteGate177Repair {
		t.Fatalf("continuum bridge should remain underived: %s", FormatContinuum(c))
	}
}

func TestNewSectorBranchOpenNotDerived(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	n := a.NewSector
	if !n.KnownFiniteInventoryExhausted || len(n.RequiredFeatures) < 6 {
		t.Fatalf("expected known finite inventory exhausted and feature requirements: %s", FormatNewSector(n))
	}
	if n.DerivedNewSectors != 0 || n.RepresentationCompleteHeavyMultiplets != 0 || n.CanonicalMassSpectrumCount != 0 || n.CanGenerateNonUniversalDeltaB || n.CanPromoteGate177Repair {
		t.Fatalf("new-sector branch should be open, not derived: %s", FormatNewSector(n))
	}
}

func TestRejectedOriginsAreFirewallSafe(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Rejected) < 3 || !allRejectedOriginsSafe(a.Rejected) {
		t.Fatalf("rejected shortcuts should preserve firewall: %s", FormatRejected(a.Rejected))
	}
	if !a.Dichotomy.ObservedFitRejectedAsOrigin || !a.Dichotomy.SchemeOnlyRejectedAsThresholdOrigin {
		t.Fatalf("fit and scheme routes should be rejected as threshold origins: %s", FormatDichotomy(a.Dichotomy))
	}
}

func TestDichotomyCompleteNoNullityReduction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Dichotomy.DichotomyCompleteAtCurrentStage || a.Dichotomy.ThresholdOriginDerived || a.Dichotomy.Gate177RepairPromoted {
		t.Fatalf("expected complete open dichotomy without promoted repair: %s", FormatDichotomy(a.Dichotomy))
	}
	f := a.Firewall
	if f.UsesObservedInputForDerivation || f.NonUniversalDeltaBDerived || f.ThresholdOperatorDerived || f.ThresholdCorrectedBetaDerived || f.PhysicalConstantsDerived || f.BoundaryScaleDerived {
		t.Fatalf("firewall violation: %s", FormatFirewall(f))
	}
	if f.StrictNullityBefore != 3 || f.StrictNullityAfter != 3 || f.ConditionalNullityBefore != 2 || f.ConditionalNullityAfter != 2 {
		t.Fatalf("nullity should remain unchanged: %s", FormatFirewall(f))
	}
}
