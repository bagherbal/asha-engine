package topologicalcouplingnormalization

import (
	"math"
	"testing"
)

func TestInputs(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Inputs.HighestInheritedGate != inheritedHighestGate || a.Inputs.AddsEmpiricalFit {
		t.Fatalf("bad inputs: %s", FormatInputs(a.Inputs))
	}
}

func TestNormalizationLanes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !nearlyEqual(a.PiLane.AlphaInverse, 8.0*math.Pi, 1e-12) || !nearlyEqual(a.PiLane.GStarSquared, 0.5, 1e-12) || !a.PiLane.MatchesNearHiggs || a.PiLane.PromotedTheorem {
		t.Fatalf("bad pi lane: %s", FormatLane(a.PiLane))
	}
	if !nearlyEqual(a.TwoPiLane.AlphaInverse, 4.0*math.Pi, 1e-12) || !nearlyEqual(a.TwoPiLane.GStarSquared, 1.0, 1e-12) || a.TwoPiLane.MatchesNearHiggs || a.TwoPiLane.PromotedTheorem {
		t.Fatalf("bad two-pi lane: %s", FormatLane(a.TwoPiLane))
	}
}

func TestChernWeilFactorTwo(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !nearlyEqual(a.ChernWeil.FactorOfTwo, 2.0, 1e-12) || !a.ChernWeil.PiLaneRequiresExtraHalf || !a.ChernWeil.DoubledSpaceCouldSupplyHalf || a.ChernWeil.DoubledSpaceHalfDerivedHere || !a.ChernWeil.RepresentationTraceRequired || a.ChernWeil.DerivedAsSpectralActionProof {
		t.Fatalf("bad Chern-Weil audit: %s", FormatChernWeil(a.ChernWeil))
	}
}

func TestDimensionWitness(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !nearlyEqual(a.Dimension.Value, 8.0*math.Pi, 1e-12) || !a.Dimension.EqualsPiLane || !a.Dimension.UsesOnlyDerivedIntegers || !a.Dimension.RequiresPiNormalization || a.Dimension.PromotedToActionTheorem {
		t.Fatalf("bad dimension witness: %s", FormatDimension(a.Dimension))
	}
}

func TestFirewallsAndStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Audit.NoEmpiricalAlphaInserted || !a.Audit.NoObservedHiggsFitInserted || !a.Audit.NoFactorTwoInvented || !a.Audit.NoTraceIndexInvented || !a.Audit.NoPoleMassClaimed || !a.Audit.NoFinalColliderMassClaimed || !a.Audit.EightPiKeptAsWitness || a.Audit.FiniteCorePolluted {
		t.Fatalf("bad firewall audit: %s", FormatAudit(a.Audit))
	}
	statuses := Statuses(a)
	required := []string{StatusPiDenominatorWitnessComputed, StatusTwoPiDenominatorLaneComputed, StatusFactorTwoObligationIdentified, StatusFailedEightPiNotPromoted, StatusFailedFactorTwoNotDerived, StatusFailedColliderMassNotClaimed}
	for _, req := range required {
		found := false
		for _, got := range statuses {
			if got == req {
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
	res := TopologicalActionChernWeilCouplingNormalizationFactorAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
