package generation2hodgeprojectorplanemetricratioaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate644Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ResidualTensorCertified || !a.Inherited.SameSectorHodgeDiagonal || !a.Inherited.OffSectorCarrierRejected || a.Inherited.NativeTraceIdentityFound || a.Inherited.SplitG2Certified || a.Inherited.BoundaryStressAssignment || a.Inherited.SevenOver72Theorem || a.Inherited.ScalarFlavorTransport || a.Inherited.PhysicalAngle || a.Inherited.PhysicalMetric || !a.Inherited.Gate643FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if math.Abs(a.Inherited.RPlusPlusFrobSquared-3.0/7.0) > 1e-8 || math.Abs(a.Inherited.RMinusMinusFrobSquared-4.0/7.0) > 1e-8 || math.Abs(a.Inherited.TwiceRPlusMinusFrobSq) > 1e-8 {
		t.Fatalf("bad inherited residual block profile: %+v", a.Inherited)
	}
	if !a.Definitions.ProjectorPlaneMetricsCertified || a.Definitions.BHatTargetResidual > strictTolerance || a.Definitions.ProjectorPlaneTargetResidual > strictTolerance {
		t.Fatalf("bad definitions: %+v", a.Definitions)
	}
	if math.Abs(a.Definitions.GHatPlusWeight-1/math.Sqrt(31)) > strictTolerance || math.Abs(a.Definitions.GHatMinusWeight+3/math.Sqrt(31)) > strictTolerance {
		t.Fatalf("bad G_hat weights: %+v", a.Definitions)
	}
	if !a.MetricRatio.AllRoutesRatioCertified || len(a.MetricRatio.Routes) != 3 || a.MetricRatio.MaxProjectorPlaneResidual > ratioTolerance || a.MetricRatio.MaxReconstructedResidual > ratioTolerance || a.MetricRatio.MaxOffDiagonalNorm > ratioTolerance || a.MetricRatio.MaxRatioDrift > ratioTolerance {
		t.Fatalf("bad metric ratio audit: %+v", a.MetricRatio)
	}
	seen := map[string]bool{}
	for _, r := range a.MetricRatio.Routes {
		seen[r.Name] = true
		if !r.Ratio1ToMinus3Certified || math.Abs(r.PlusBlockMean-1/math.Sqrt(31)) > ratioTolerance || math.Abs(r.MinusBlockMean+3/math.Sqrt(31)) > ratioTolerance || math.Abs(r.ObservedMinusToPlusRatio+3) > ratioTolerance || r.PlusMinusFrobNorm > ratioTolerance {
			t.Fatalf("bad route %s: %+v", r.Name, r)
		}
	}
	for _, want := range []string{"omega_1_alt", "omega_2_alt", "omega_B_alt"} {
		if !seen[want] {
			t.Fatalf("missing route %s", want)
		}
	}
	if !a.AngleFromPlane.AngleDerivedFromPlane || a.AngleFromPlane.NativeTraceIdentityFound || math.Abs(a.AngleFromPlane.ComputedCosine-13/math.Sqrt(217)) > strictTolerance || math.Abs(a.AngleFromPlane.ComputedSinSquared-48.0/217.0) > strictTolerance {
		t.Fatalf("bad angle: %+v", a.AngleFromPlane)
	}
	if a.MinusThree.CertifiedNativeSource || !strings.Contains(a.MinusThree.Verdict, StatusNoMinusThreeSource) {
		t.Fatalf("bad minus-three audit: %+v", a.MinusThree)
	}
	if !a.Interpretation.Gate643Inherited || !a.Interpretation.GHatReconstructed || !a.Interpretation.RatioCertified || !a.Interpretation.AngleFromPlane || a.Interpretation.MinusThreeSourceFound || a.Interpretation.NativeTraceIdentityFound {
		t.Fatalf("bad interpretation: %+v", a.Interpretation)
	}
	if a.Firewalls.ClaimsNativeTraceIdentity || a.Firewalls.ClaimsSplitG2 || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72Theorem || a.Firewalls.ClaimsScalarFlavor || a.Firewalls.ClaimsPhysicalAngle || a.Firewalls.ClaimsPhysicalMetric || a.Firewalls.ClaimsFlavor || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2HodgeProjectorPlaneMetricRatioAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate643ResidualInherited, StatusGHatReconstructed, StatusProjectorPlaneMetricsDefined, StatusRouteMetricRatiosComputed, StatusHodgeDiagonalRatio, StatusProjectorPlaneAngle, StatusMinusThreeSourceCandidate, StatusNoMinusThreeSource, StatusNoNativeTraceIdentity, StatusNoCertifiedSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72Theorem, StatusNoScalarFlavorTransport, StatusNoPhysicalAngle, StatusNoPhysicalMetric, StatusNoHiggsFlavorGauge, StatusGate644Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
