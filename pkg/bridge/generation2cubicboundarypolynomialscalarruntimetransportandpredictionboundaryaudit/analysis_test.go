package generation2cubicboundarypolynomialscalarruntimetransportandpredictionboundaryaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate734CubicRuntimeBridge(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate733.Inherited || !a.Gate733.CurrentBestClosure || !a.Gate733.NoNativeGeneratingFunction || !a.Gate733.NoNativeMomentExpansion {
		t.Fatalf("bad Gate733 inheritance: %+v", a.Gate733)
	}
	expectedW3 := math.Abs(lambdaLambda12) + a.Gate733.FWall3
	if math.Abs(a.BoundarySub.W3-expectedW3) > 1e-18 || math.Abs(a.BoundarySub.KappaLambdaApprox-(a.BoundarySub.W3-a.Gate733.KappaE)) > 1e-18 || math.Abs(a.BoundarySub.KappaLambdaApprox+a.BoundarySub.DroppedPolynomialResidual-kappaLambda) > runtimeTolerance {
		t.Fatalf("bad boundary substitution: %+v", a.BoundarySub)
	}
	expectedRuntime := lambdaProxyMZ * (1 + a.Gate733.L*(1-a.BoundarySub.W3+a.Gate733.KappaE))
	if !a.Runtime.UsesCubicBoundaryWound || math.Abs(a.Runtime.RuntimeApprox-expectedRuntime) > 1e-18 {
		t.Fatalf("bad runtime bridge: %+v", a.Runtime)
	}
	if !a.SourceType.SourceTypingRecorded || !strings.Contains(a.SourceType.BoundaryPolynomial, "M1_wall") || !strings.Contains(a.SourceType.RadialHopfLoop, "P_rad") {
		t.Fatalf("bad source typing: %+v", a.SourceType)
	}
}

func TestGate734ResidualPredictionBoundaryAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	expectedResidual := a.Gate733.LambdaProxy * a.Gate733.L * a.Gate733.EPoly3
	if math.Abs(a.Propagation.RuntimeResidual-expectedResidual) > 1e-18 || !a.Propagation.MatchesPropagation || !a.Propagation.NearlyEliminated {
		t.Fatalf("bad residual propagation: %+v", a.Propagation)
	}
	if !a.Prediction.KappaLambdaDefinedFromRuntime || a.Prediction.CubicRuntimeIndependentPrediction || !a.Prediction.ConsistencyClosure {
		t.Fatalf("bad prediction boundary: %+v", a.Prediction)
	}
	if !a.Seals.DependsOnN || !a.Seals.DependsOnPRad || !a.Seals.DependsOnRhoPlus || !a.Seals.DependsOnRho72 || !a.Seals.DependsOnPK7 || !a.Seals.DependsOnKappaE || !a.Seals.DependsOnLambdaProxy || !a.Seals.DependsOnL || a.Seals.PremisesNativelyDerived {
		t.Fatalf("bad seal dependence: %+v", a.Seals)
	}
	if a.Firewall.ClaimsHiggsPoleMassPrediction || a.Firewall.ClaimsNativeScalarRuntimeTheorem || a.Firewall.ClaimsNativeScalarPotentialTheorem || a.Firewall.ClaimsYukawaEigenvalueTheorem || a.Firewall.ClaimsFlavorHierarchyTheorem || a.Firewall.ClaimsCKMPMNSTheorem {
		t.Fatalf("forecast firewall failed: %+v", a.Firewall)
	}

	res := Generation2CubicBoundaryPolynomialScalarRuntimeTransportAndPredictionBoundaryAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected construction failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
