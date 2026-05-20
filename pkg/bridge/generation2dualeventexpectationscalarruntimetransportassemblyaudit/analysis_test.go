package generation2dualeventexpectationscalarruntimetransportassemblyaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate728DualEventAssembly(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate700.Inherited || math.Abs(a.Gate700.P_K7-7.0/72.0) > 1e-18 || math.Abs(a.Gate700.DBase-a.Gate700.ExpectedHistoryResponse-a.Gate700.EWall) > 1e-18 {
		t.Fatalf("bad Gate700 inheritance: %+v", a.Gate700)
	}
	if !a.Gate727.Inherited || math.Abs(a.Gate727.L-1/(8*math.Pi)) > 1e-18 || math.Abs(a.Gate727.L-a.Gate727.RadialHopfExpectation) > 1e-18 || !a.Gate727.ConditionallyExact {
		t.Fatalf("bad Gate727 inheritance: %+v", a.Gate727)
	}
	if math.Abs(a.BoundarySub.W72-(math.Abs(lambdaLambda12)+a.Gate700.P_K7*a.Gate700.SSplit)) > 1e-18 || math.Abs(a.BoundarySub.KappaLambdaApprox+a.BoundarySub.EWall-kappaLambda) > 1e-12 || !a.BoundarySub.BoundaryMinusFlavorReading {
		t.Fatalf("bad boundary substitution: %+v", a.BoundarySub)
	}
	if !a.RadialSub.UsesRadialHopfExpectation || math.Abs(a.RadialSub.FactorApprox-(1-a.BoundarySub.W72+kappaE)) > 1e-18 || math.Abs(a.RadialSub.RuntimePredApprox-lambdaProxyMZ*(1+a.Gate727.L*(1-a.BoundarySub.W72+kappaE))) > 1e-18 {
		t.Fatalf("bad radial substitution: %+v", a.RadialSub)
	}
	if !a.Assembly.DualEventExpectationForm || math.Abs(a.Assembly.AssembledRuntimeApprox-a.RadialSub.RuntimePredApprox) > 1e-18 {
		t.Fatalf("bad dual assembly: %+v", a.Assembly)
	}
}

func TestGate728ResidualNoncircularityAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	expectedDelta := lambdaProxyMZ * a.Gate727.L * a.Gate700.EWall
	if math.Abs(a.Propagation.DeltaLambdaPred-expectedDelta) > 1e-18 || !a.Propagation.MatchesPropagationFormula || !a.Propagation.RuntimeResidualIsWallResidual {
		t.Fatalf("bad propagation: %+v", a.Propagation)
	}
	if !a.NonCircular.KappaLambdaDefinedFromRuntime || a.NonCircular.AssembledIndependentPrediction || !a.NonCircular.BridgeConsistencyClosure {
		t.Fatalf("bad noncircularity audit: %+v", a.NonCircular)
	}
	if !a.Seals.DependsOnN || !a.Seals.DependsOnPRad || !a.Seals.DependsOnRhoPlus || !a.Seals.DependsOnRho72 || !a.Seals.DependsOnPK7 || !a.Seals.DependsOnKappaE || a.Seals.PremisesNativelyDerived {
		t.Fatalf("bad seal dependence: %+v", a.Seals)
	}
	if a.Firewall.ClaimsScalarRuntimeTheorem || a.Firewall.ClaimsHiggsMassTheorem || a.Firewall.ClaimsNativeHistoryLoopUnit || a.Firewall.ClaimsNativeBoundaryHistory || a.Firewall.ClaimsNativeRadialSelector || a.Firewall.ClaimsYukawaOperatorTheorem || a.Firewall.ClaimsIndependentPrediction {
		t.Fatalf("firewall failed: %+v", a.Firewall)
	}

	res := Generation2DualEventExpectationScalarRuntimeTransportAssemblyAuditTheorem().Verify()
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
