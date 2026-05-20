package generation2scalarflavordeficitclosuretriangleaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate659Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ScalarTransportSpineInherited || !a.Inherited.ScalarBoundarySpineActive || !a.Inherited.LowScaleMatchingActive || !a.Inherited.BoundaryStressTransportActive || !a.Inherited.KappaLambdaDefined || !a.Inherited.NoNativeProxyRuntimeTheorem || !a.Inherited.NoNativeRGThresholdTheorem || !a.Inherited.NoNativeBoundaryStressTheorem || !a.Inherited.NoHiggsMassOrStabilityClaim || !a.Inherited.FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.Flavor.EnvironmentalSeal || a.Flavor.NativeFlavorTheorem || math.Abs(a.Flavor.KappaE-kappaE) > 1e-15 || !strings.Contains(a.Flavor.OrientationBalanceExpression, "J_CKM") {
		t.Fatalf("bad flavor seal: %+v", a.Flavor)
	}
	if math.Abs(a.Closure.KSum-0.04982659728765166) > 5e-15 || math.Abs(a.Closure.DeltaClosure-0.0001256552099683575) > 5e-15 || !a.Closure.ClosesOnScalarWound || !a.Closure.RawClosureResidualSmall || a.Closure.RelativeToAbsLambda >= 0.003 {
		t.Fatalf("bad closure: %+v", a.Closure)
	}
	if math.Abs(a.Boundary.BoundarySplit-0.0012924448188162962) > 5e-15 || math.Abs(a.Boundary.ObservedWeight-0.09722288188941064) > 5e-13 || !a.Boundary.SevenOver72Closest || !a.Boundary.UsedTypedSetOnly || len(a.Boundary.Candidates) != 4 {
		t.Fatalf("bad boundary weight: %+v", a.Boundary)
	}
	if math.Abs(a.Interpolation.WeightedTarget-0.04982659643506822) > 5e-15 || math.Abs(a.Interpolation.WeightedResidual-8.525834413464217e-10) > 5e-18 || a.Interpolation.ImprovementFactor < 100000 || a.Interpolation.ResidualRelativeToKSum >= 2e-8 || !a.Interpolation.BridgeLayerOnly {
		t.Fatalf("bad interpolation: %+v", a.Interpolation)
	}
	if len(a.Sources.Objects) != 5 || a.Sources.SevenOver72InFanoLane || !a.Sources.SevenOver72InTransportLane || a.Sources.FanoBoundaryMapConstructed || a.Sources.RandomConstantsSearched || !a.Sources.TypedCandidatesOnly {
		t.Fatalf("bad source audit: %+v", a.Sources)
	}
	if a.Firewalls.ClaimsNativeFlavorTheorem || a.Firewalls.ClaimsNativeScalarTheorem || a.Firewalls.ClaimsNativeSevenOver72Theorem || a.Firewalls.ClaimsBoundaryStressDerivation || a.Firewalls.ClaimsHiggsPrediction || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.ClaimsCKMPMNSDerivation || a.Firewalls.ClaimsPhysicalSpacetime || a.Firewalls.ClaimsNativeClosureTheorem || a.Firewalls.Verdict != StatusGate659Boundary {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2ScalarFlavorDeficitClosureTriangleAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate658ScalarTransportSpineInherited, StatusFlavorKappaESealInherited, StatusKappaSumComputed, StatusKappaSumClosesOnAbsLambda, StatusBoundarySplitRatioComputed, StatusTypedWeightCandidatesAudited, StatusSevenOverSeventyTwoInterpolation, StatusSourceTypeAuditComputed, StatusKappaClosureSupport, StatusResidualTracksBoundarySplit, StatusSevenOver72ReappearsActiveTransport, StatusBoundaryWeightedClosureSupport, StatusNoNativeKappaClosureTheorem, StatusNoNativeSevenOver72SourceTheorem, StatusNoNativeScalarFlavorBoundaryTheorem, StatusNoNativeFlavorTheorem, StatusNoNativeScalarTheorem, StatusNoBoundaryStressDerivation, StatusNoHiggsGaugeOrCKMClaim, StatusGate659Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
