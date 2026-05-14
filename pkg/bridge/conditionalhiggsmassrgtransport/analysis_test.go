package conditionalhiggsmassrgtransport

import (
	"math"
	"strings"
	"testing"
)

func TestBoundaryAndBetaSystem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Boundary.BoundaryInherited || a.Boundary.GStarSquaredSealed != 1 || a.Boundary.AbsoluteFinalClaim {
		t.Fatalf("bad boundary: %s", FormatBoundary(a.Boundary))
	}
	wantLambda := 1197.0 / 4624.0
	if math.Abs(a.Boundary.LambdaUV-wantLambda) > 1e-15 {
		t.Fatalf("wrong lambda boundary %.15g", a.Boundary.LambdaUV)
	}
	if !a.Beta.OneLoopOnly || !a.Beta.UsesStandardContinuumQFT || a.Beta.DerivedAsFiniteCoreTheorem {
		t.Fatalf("bad beta system: %s", FormatBeta(a.Beta))
	}
}

func TestThresholdAndTopLanes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if len(a.Thresholds) != 3 || len(a.TopLanes) != 2 {
		t.Fatalf("unexpected lane counts thresholds=%d top=%d", len(a.Thresholds), len(a.TopLanes))
	}
	if !a.Thresholds[1].ConditionalOnPeVSeal || !a.Thresholds[2].ConditionalOnPeVSeal || a.Thresholds[1].FiniteThresholdOrigin || a.Thresholds[2].MatchingCorrectionsDerived {
		t.Fatalf("bad PeV threshold lanes: %s | %s", FormatThreshold(a.Thresholds[1]), FormatThreshold(a.Thresholds[2]))
	}
	if a.TopLanes[1].Name != "r_plus_top_yukawa_boundary_seal" || a.TopLanes[1].FiniteCoreDerived || a.TopLanes[1].DiagnosticOnly || a.TopLanes[1].YtUV < 1.28 || a.TopLanes[1].YtUV > 1.29 {
		t.Fatalf("bad r_plus top lane: %s", FormatTopLane(a.TopLanes[1]))
	}
}

func TestTransportResults(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if len(a.Results) != 6 {
		t.Fatalf("expected six transport results, got %d", len(a.Results))
	}
	pureFailures := 0
	computedPeV := 0
	for _, r := range a.Results {
		if strings.Contains(r.ThresholdName, "pure_SM") && !r.Computed {
			pureFailures++
		}
		if strings.Contains(r.ThresholdName, "PeV") && r.Computed && r.LambdaPositive && r.HiggsMassGeV > 0 {
			computedPeV++
		}
	}
	if pureFailures != 2 || computedPeV != 4 {
		t.Fatalf("bad transport classification pureFailures=%d computedPeV=%d results=%+v", pureFailures, computedPeV, a.Results)
	}
}

func TestConditionalMassDiagnostic(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Prediction.RPlusTopLaneComputed || !a.Prediction.PureSMRunInvalid || a.Prediction.ObservedHiggsMassInserted || !a.Prediction.MeasuredComparisonOnly {
		t.Fatalf("bad prediction audit: %s", FormatPrediction(a.Prediction))
	}
	if a.Prediction.PrimaryConditionalMassGeV < 320 || a.Prediction.PrimaryConditionalMassGeV > 340 {
		t.Fatalf("unexpected primary conditional mass: %s", FormatPrediction(a.Prediction))
	}
	if a.Prediction.GaugeOnlyDiagnosticMassGeV < 150 || a.Prediction.GaugeOnlyDiagnosticMassGeV > 165 {
		t.Fatalf("unexpected gauge-only mass: %s", FormatPrediction(a.Prediction))
	}
	if a.Prediction.RPlusTopLaneNearObserved {
		t.Fatalf("r_plus lane should be a tension diagnostic, not near-observed: %s", FormatPrediction(a.Prediction))
	}
}

func TestFirewallsAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoObservedHiggsMassUsedForDerivation || !a.Firewalls.NoObservedTopMassUsedForDerivation || !a.Firewalls.NoTwoLoopTermsInserted || !a.Firewalls.NoThresholdMatchingInserted || !a.Firewalls.NoPoleMassMatchingInserted || !a.Firewalls.PeVThresholdsRemainSealed || !a.Firewalls.TopYukawaOriginRemainsSealed || !a.Firewalls.PureSMPathologyRecorded || a.Firewalls.FinalColliderPredictionClaimed || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failed: %s", FormatFirewalls(a.Firewalls))
	}
	res := ConditionalHiggsMassFromQuarticRGTransportAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
