package spectralactioncouplingnormalization

import (
	"math"
	"testing"
)

func TestInputs(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Inputs.HighestInheritedGate != inheritedHighestGate || a.Inputs.AddsEmpiricalFit || a.Inputs.UsesObservedHiggsFit {
		t.Fatalf("bad inputs: %s", FormatInputs(a.Inputs))
	}
}

func TestSpectralActionLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !nearlyEqual(a.Spectral.N4ForEightPi, 2.0/7.0, 1e-12) {
		t.Fatalf("expected N4=2/7 for 8π: %s", FormatSpectral(a.Spectral))
	}
	expectedRequiredTrace := 4.0 * math.Pi * math.Pi / 7.0
	if !nearlyEqual(a.Spectral.RequiredTraceRepIndexFor8Pi, expectedRequiredTrace, 1e-12) || a.Spectral.StandardTraceIndexKnown || a.Spectral.EightPiDerivedByThisLane {
		t.Fatalf("bad spectral ledger: %s", FormatSpectral(a.Spectral))
	}
}

func TestTopologicalAndDimensionLanes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Topological.MatchesEightPi || !nearlyEqual(a.Topological.AlphaInverse, 8.0*math.Pi, 1e-12) || !nearlyEqual(a.Topological.GStarSquared, 0.5, 1e-12) {
		t.Fatalf("bad topological lane: %s", FormatTopological(a.Topological))
	}
	if !a.Dimension.EqualsTopologicalLane || !a.Dimension.UsesOnlyDerivedCounts || a.Dimension.ProvedAsSpectralTheorem {
		t.Fatalf("bad dimension lane: %s", FormatDimension(a.Dimension))
	}
}

func TestHiggsProxy(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	expectedLambda := (1197.0 / 4624.0) * 0.5
	expectedMass := electroweakVEVGeV * math.Sqrt(2.0*expectedLambda)
	if !nearlyEqual(a.Higgs.LambdaH, expectedLambda, 1e-12) || !nearlyEqual(a.Higgs.PredictedMassGeV, expectedMass, 1e-9) {
		t.Fatalf("bad Higgs proxy: %s", FormatHiggs(a.Higgs))
	}
	if a.Higgs.PredictedMassGeV < 125.0 || a.Higgs.PredictedMassGeV > 125.6 || a.Higgs.PoleMassClaimed {
		t.Fatalf("proxy out of expected firewalled range: %s", FormatHiggs(a.Higgs))
	}
}

func TestFirewallsAndStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Audit.NoAlphaGUTFitInserted || !a.Audit.NoTraceIndexInvented || !a.Audit.NoSpectralActionProofOverclaimed || !a.Audit.NoObservedHiggsMassFitInserted || !a.Audit.NoPoleMassClaimed || !a.Audit.NoTwoLoopClaimed || !a.Audit.NoFinalColliderMassClaimed || a.Audit.FiniteCorePolluted {
		t.Fatalf("bad audit: %s", FormatAudit(a.Audit))
	}
	statuses := Statuses(a)
	required := []string{StatusTopologicalActionCouplingWitness, StatusGStarHalfBoundaryComputed, StatusHiggsProxyFromTopologicalCoupling, StatusFailedAlphaGUTDerivationNotClosed, StatusFailedDimensionFormulaNotSpectralProof, StatusFailedFinalColliderMassNotClaimed}
	for _, req := range required {
		found := false
		for _, s := range statuses {
			if s == req {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("missing status %s in %v", req, statuses)
		}
	}
}

func TestTheoremPasses(t *testing.T) {
	res := SpectralActionCouplingNormalizationAlphaGUTAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
