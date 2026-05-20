package generation2compactsplitresidualtensorblockstructureaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate643Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.HodgePolaritySkeleton || a.Inherited.NativeTraceIdentityCertified || a.Inherited.SplitG2Certified || a.Inherited.BoundaryStressAssignment || a.Inherited.SevenOver72Theorem || a.Inherited.ScalarFlavorTransport || a.Inherited.PhysicalAngle || a.Inherited.PhysicalMetric || !a.Inherited.Gate642FirewallPreserved {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if math.Abs(a.Inherited.CosTheta-(float64(alignmentRoot)/math.Sqrt(float64(angleDenominator)))) > strictTolerance {
		t.Fatalf("bad inherited cosine: %+v", a.Inherited)
	}
	if math.Abs(a.Inherited.SinTheta-(4*math.Sqrt(3)/math.Sqrt(float64(angleDenominator)))) > strictTolerance {
		t.Fatalf("bad inherited sine: %+v", a.Inherited)
	}
	if !a.ResidualTensor.ResidualTensorsCertified || len(a.ResidualTensor.Routes) != 3 || a.ResidualTensor.MaxOrthogonalityToBHat > angleTolerance || a.ResidualTensor.MaxResidualUnitNormDrift > angleTolerance || a.ResidualTensor.MaxCosineDrift > 1e-6 || a.ResidualTensor.MaxRhoDrift > 1e-6 {
		t.Fatalf("bad residual tensor audit: %+v", a.ResidualTensor)
	}
	seen := map[string]bool{}
	for _, r := range a.ResidualTensor.Routes {
		seen[r.Name] = true
		if math.Abs(r.BlockNormSum-1) > 1e-8 || r.PlusPlusRank == 0 || r.MinusMinusRank == 0 || !r.TypedBlockProfile {
			t.Fatalf("bad block profile for %s: %+v", r.Name, r)
		}
	}
	for _, want := range []string{"omega_1_alt", "omega_2_alt", "omega_B_alt"} {
		if !seen[want] {
			t.Fatalf("missing residual route %s", want)
		}
	}
	if a.BlockSummary.RouteCount != 3 || !a.BlockSummary.HasTypedBlockStructure || a.BlockSummary.NativeTraceIdentityFound || !a.BlockSummary.SameRankProfileAllRoutes {
		t.Fatalf("bad block summary: %+v", a.BlockSummary)
	}
	if !a.Interpretation.AnglePairInherited || !a.Interpretation.ResidualTensorDefined || !a.Interpretation.BlocksComputed || !a.Interpretation.TypedBlockStructure || a.Interpretation.NativeTraceIdentityFound {
		t.Fatalf("bad interpretation: %+v", a.Interpretation)
	}
	if a.Firewalls.ClaimsNativeTraceIdentity || a.Firewalls.ClaimsSplitG2 || a.Firewalls.ClaimsBoundaryStress || a.Firewalls.ClaimsSevenOver72Theorem || a.Firewalls.ClaimsScalarFlavor || a.Firewalls.ClaimsPhysicalAngle || a.Firewalls.ClaimsPhysicalMetric || a.Firewalls.ClaimsFlavor || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsCKMPMNS || a.Firewalls.ClaimsGaugeUnification {
		t.Fatalf("firewall breach: %+v", a.Firewalls)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2CompactSplitResidualTensorBlockStructureAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{StatusGate642AngleInherited, StatusResidualTensorDefined, StatusHodgePolarityBlocksComputed, StatusRouteBlockProfilesComputed, StatusResidualBlockStructure, StatusNoNativeTraceIdentity, StatusNoCertifiedSplitG2, StatusNoBoundaryStress, StatusNoSevenOver72Theorem, StatusNoScalarFlavorTransport, StatusNoPhysicalAngle, StatusNoPhysicalMetric, StatusNoHiggsFlavorGauge, StatusGate643Boundary} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
