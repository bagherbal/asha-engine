package doubledbosonictraceindex

import (
	"math"
	"testing"
)

func TestInputs(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Inputs.HighestInheritedGate != inheritedHighestGate || !nearlyEqual(a.Inputs.RequiredMultiplier, 2.0, 1e-12) || a.Inputs.AddsEmpiricalFit {
		t.Fatalf("bad inputs: %s", FormatInputs(a.Inputs))
	}
}

func TestMirrorCarrier(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !nearlyEqual(a.Mirror.FullDoubledTraceIndex, 2.0, 1e-12) || !a.Mirror.CurvaturesHaveSameF2Sign || !a.Mirror.ComplexConjugationNeutral {
		t.Fatalf("bad mirror carrier: %s", FormatMirror(a.Mirror))
	}
}

func TestTraceLanes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !nearlyEqual(a.BaseLane.AlphaInverse, 4.0*math.Pi, 1e-12) || !nearlyEqual(a.BaseLane.GStarSquared, 1.0, 1e-12) {
		t.Fatalf("bad base lane: %s", FormatLane(a.BaseLane))
	}
	if !a.Doubled.MatchesEightPi || !nearlyEqual(a.Doubled.AlphaInverse, 8.0*math.Pi, 1e-12) || !nearlyEqual(a.Doubled.GStarSquared, 0.5, 1e-12) || !a.Doubled.MatchesHiggs {
		t.Fatalf("bad doubled lane: %s", FormatLane(a.Doubled))
	}
	if !nearlyEqual(a.Quotient.TraceMultiplier, 0.5, 1e-12) || a.Quotient.MatchesEightPi || a.Quotient.MatchesHiggs {
		t.Fatalf("bad quotient lane: %s", FormatLane(a.Quotient))
	}
}

func TestPromotionAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Promotion.MultiplierMatches || !a.Promotion.ConditionalPromotion || a.Promotion.UnconditionalDerivation || a.Promotion.FullBosonicTraceNativeHere || !a.Promotion.QuotientConventionRejected {
		t.Fatalf("bad promotion audit: %s", FormatPromotion(a.Promotion))
	}
	if !a.Audit.NoEmpiricalAlphaInserted || !a.Audit.NoObservedHiggsFitInserted || !a.Audit.NoPoleMassClaimed || !a.Audit.NoFinalColliderMassClaimed || !a.Audit.QuotientLaneKeptVisible || !a.Audit.TraceConventionStillConditional || a.Audit.FiniteCorePolluted {
		t.Fatalf("bad firewall audit: %s", FormatAudit(a.Audit))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	statuses := Statuses(a)
	required := []string{StatusDoubledBosonicTraceFactorFormalized, StatusEightPiBranchConditionallyPromoted, StatusTensionBosonicTraceConventionRequired, StatusFailedTraceConventionNotNative, StatusFailedUnconditionalAlphaNotDerived, StatusFailedColliderMassNotClaimed}
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
	res := DoubledBosonicTraceIndexJMirrorGaugeCapacityAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
