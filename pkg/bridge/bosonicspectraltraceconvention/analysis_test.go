package bosonicspectraltraceconvention

import (
	"math"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if a.Inputs.HighestInheritedGate != inheritedHighestGate || a.Inputs.AddsEmpiricalAlpha {
		t.Fatalf("bad inputs: %s", FormatInputs(a.Inputs))
	}
}

func TestTraceAxiom(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Trace.BosonicTraceUsesFullHilbertSpace || a.Trace.FermionicHalfAppliesToBosons {
		t.Fatalf("bad trace axiom: %s", FormatTrace(a.Trace))
	}
}

func TestMirrorAndSeparation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !nearlyEqual(a.Mirror.TotalBosonicIndex, 2.0, 1e-12) || !a.Mirror.SameSign || !a.Mirror.Positive {
		t.Fatalf("bad mirror: %s", FormatMirror(a.Mirror))
	}
	if !a.Separation.HalfFactorConfinedToFermions || !a.Separation.QuotientLaneRejected || !a.Separation.ApplyingHalfToBosonsBreaksEightPi {
		t.Fatalf("bad separation: %s", FormatSeparation(a.Separation))
	}
}

func TestCouplingLanes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !nearlyEqual(a.Single.AlphaInverse, 4.0*math.Pi, 1e-12) || !nearlyEqual(a.Single.GStarSquared, 1.0, 1e-12) {
		t.Fatalf("bad single lane: %s", FormatLane(a.Single))
	}
	if !nearlyEqual(a.Doubled.AlphaInverse, 8.0*math.Pi, 1e-12) || !nearlyEqual(a.Doubled.GStarSquared, 0.5, 1e-12) || !a.Doubled.NativeBosonicTrace {
		t.Fatalf("bad doubled lane: %s", FormatLane(a.Doubled))
	}
	if !nearlyEqual(a.Quotient.AlphaInverse, 2.0*math.Pi, 1e-12) || a.Quotient.NativeBosonicTrace {
		t.Fatalf("bad quotient lane: %s", FormatLane(a.Quotient))
	}
}

func TestPromotionAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Promotion.Gate330TraceConventionSuppliesFactorTwo || !a.Promotion.EightPiPromotedWithinBosonicAction || a.Promotion.AlphaUnconditional {
		t.Fatalf("bad promotion: %s", FormatPromotion(a.Promotion))
	}
	if !a.Audit.NoEmpiricalAlphaInserted || !a.Audit.NoObservedMassFitted || !a.Audit.NoPoleMassClaimed || !a.Audit.NoFinalColliderClaimed || !a.Audit.RepresentationIndexStillFirewalled || !a.Audit.TopologicalActionMapStillFirewalled {
		t.Fatalf("bad audit: %s", FormatAudit(a.Audit))
	}
}

func TestStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	statuses := Statuses(a)
	required := []string{StatusBosonicFullTraceNative, StatusEightPiBranchPromotedByTraceConvention, StatusFailedAlphaUnconditionalNotDerived, StatusFailedColliderMassNotClaimed}
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
	res := BosonicSpectralActionTraceConventionFullDoubledSpaceAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
