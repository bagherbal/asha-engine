package universalbetasource

import "testing"

func TestGate202Inheritance(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.PreviousGate202.Gate202Inherited || !a.PreviousGate202.BGapContactTraceOffsetFailedRoute || !a.PreviousGate202.NoPhysicalPredictionClaim || len(a.PreviousGate202.Requirements) != 2 {
		t.Fatalf("bad Gate 202 inheritance: %s", FormatGate202(a.PreviousGate202))
	}
}

func TestCompleteMultipletRowsAreExactButDoNotMatchRequiredUniversalRows(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.MultipletBasis) < 6 || a.MultipletAudit.BasisRowsAudited != len(a.MultipletBasis) || a.MultipletAudit.GUTCompleteRows != len(a.MultipletBasis) || a.MultipletAudit.ExactOneLoopRows != len(a.MultipletBasis) {
		t.Fatalf("bad multiplet basis audit: %s", FormatCompleteMultipletAudit(a.MultipletAudit))
	}
	if a.MultipletAudit.ExactIntegerMultipletMatches != 0 || a.MultipletAudit.CompleteMultipletSourceFound || a.MultipletAudit.ConditionalPredictions != 0 {
		t.Fatalf("unexpected complete multiplet source: %s :: %s", FormatCompleteMultipletAudit(a.MultipletAudit), FormatMultipletFits(a.MultipletFits, 8))
	}
}

func TestFiniteInventoryDoesNotPromoteContactOrFockToHeavyThreshold(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.FiniteInventory
	if f.ContactPartialOverlapModes != 7 || f.ContactRowsHaveChargeSemantics || f.ContactRowsHaveGaugeRepresentation || f.ContactRowsHaveDynkinIndex || f.ContactRowsHaveBetaPermission {
		t.Fatalf("contact inventory overclaimed: %s", FormatFiniteInventory(f))
	}
	if f.FockStates != 16 || !f.FockKinematicSO10SixteenAvailable || !f.FockRepTraceBoundarySeedClosed || f.FockHeavyDuplicateDerived || f.FockThresholdMassDerived || f.FockCompleteMultipletBetaActivated || f.FiniteCompleteMultipletFound {
		t.Fatalf("fock inventory overclaimed: %s", FormatFiniteInventory(f))
	}
}

func TestRegulatorTraceRouteIsBlocked(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	r := a.RegulatorAudit
	if r.CandidatesAudited != len(a.RegulatorCandidates) || r.CanonicalTraces == 0 {
		t.Fatalf("bad regulator inventory: %s", FormatRegulatorAudit(r))
	}
	if r.UniversalAnomalyCandidates != 0 || r.ConformalAnomalyDerived || r.GhostBRSTCancellationComplete || r.SpectralTripleComplete || r.GaugeMeasureMapDerived || r.BetaRowPermission || r.ExactRequiredMatches != 0 || r.ConditionalPredictions != 0 || r.RegulatorTraceSourceFound {
		t.Fatalf("regulator route overclaimed: %s :: %s", FormatRegulatorAudit(r), FormatRegulatorCandidates(a.RegulatorCandidates, 5))
	}
}

func TestFirewallSealed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	f := a.Firewall
	if f.UniversalBetaSourceDerived || f.CompleteHeavyMultipletDerived || f.RegulatorTraceAnomalyDerived || f.ContactModesPromotedToBetaRows || f.FockGenerationPromotedToNewThreshold || f.ArbitraryIntegerMultiplicityInserted || f.ArbitraryRegulatorCoefficientInserted || f.PhysicalUnificationClaimed || f.ThresholdCorrectedPhysicalFitClaimed || f.AbsoluteMassPredicted || f.FiniteMatchingCorrectionsDerived {
		t.Fatalf("firewall leak: %s", FormatFirewall(f))
	}
	if f.StrictNullityBefore != f.StrictNullityAfter || f.PhysicalPredictionNullityBefore != f.PhysicalPredictionNullityAfter {
		t.Fatalf("nullity changed unexpectedly: %s", FormatFirewall(f))
	}
}

func TestTheoremChecksPass(t *testing.T) {
	res := UniversalBetaSourceClassificationAuditTheorem().Verify()
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check %q failed: %s", c.Name, c.Detail)
		}
	}
}
