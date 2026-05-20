package generation2gaugescalarboundarystresssealaudit

import (
	"math"
	"testing"
)

func TestGate613BoundaryStressCompression(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	xi := findCandidate(a.CompressionCandidates, "xi_mean")
	if math.Abs(xi.Xi-0.0503471644870914) > 1e-12 {
		t.Fatalf("unexpected xi_mean %.15g", xi.Xi)
	}
	if math.Abs(a.SignedStressVector.SPlus-0.00129244481881585) > 1e-12 {
		t.Fatalf("unexpected SPlus %.15g", a.SignedStressVector.SPlus)
	}
	if !a.AntiAlignment.AntiAligned || a.AntiAlignment.RelativeAntiAlignment > 0.03 {
		t.Fatalf("expected anti-alignment, got %+v", a.AntiAlignment)
	}
}

func TestGate613EtaComparisonAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	eta := etaFor(a.EtaComparisons, "xi_mean")
	if !(eta.EtaOverTwoXi > 0.93 && eta.EtaOverTwoXi < 0.96) {
		t.Fatalf("unexpected eta/(2xi_mean): %.15g", eta.EtaOverTwoXi)
	}
	if a.NativeStatus.ProvidesNativeXiBoundary || a.NativeStatus.ProvidesNativeGaugeScalarEquation || a.NativeStatus.ClaimsHiggsPrediction {
		t.Fatalf("native theorem/firewall violation: %+v", a.NativeStatus)
	}
	if a.Firewalls.ClaimsGaugeUnification || a.Firewalls.ClaimsScalarStability || a.Firewalls.ClaimsThresholdExistence {
		t.Fatalf("firewall violation: %+v", a.Firewalls)
	}
}
