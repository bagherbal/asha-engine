package generation2level1scalarruntimebridgeconsistencyestimateandnonpredictionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate739Level1RuntimeEstimateAndSeals(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate738.Inherited || !a.Gate738.PackageMinimal || !a.Gate738.SealsIndependent || !a.Gate738.RequiresThreeSealPackage {
		t.Fatalf("bad Gate738 inheritance: %+v", a.Gate738)
	}
	if !a.Gate734.Inherited || !a.Gate734.NotIndependentPrediction || !a.Gate734.NoNativeRuntimeTheorem || !a.Gate734.NoMassTheorem || !a.Gate734.NoYukawaTheorem {
		t.Fatalf("bad Gate734 inheritance: %+v", a.Gate734)
	}
	if !a.Estimate.Level1Allowed || !a.Estimate.NearFloatScale {
		t.Fatalf("estimate not allowed/near-float: %+v", a.Estimate)
	}
	if math.Abs(a.Estimate.W3-0.049826597288039835) > 1e-16 {
		t.Fatalf("unexpected W3: %.17g", a.Estimate.W3)
	}
	if math.Abs(a.Estimate.KappaLambdaBridge-0.04432304309646527) > 1e-16 {
		t.Fatalf("unexpected kappaLambdaBridge: %.17g", a.Estimate.KappaLambdaBridge)
	}
	if math.Abs(a.Estimate.RuntimeBridge-0.12965256505047373) > 1e-15 {
		t.Fatalf("unexpected runtime bridge: %.17g", a.Estimate.RuntimeBridge)
	}
	if math.Abs(a.Estimate.RuntimeResidual) > nearFloatRuntimeTolerance {
		t.Fatalf("runtime residual too large: %.17g", a.Estimate.RuntimeResidual)
	}
	if len(a.Seals.Labels) != 10 || !a.Seals.AllExplicit || !a.Seals.AllRequiredByBridge {
		t.Fatalf("bad seals: %+v", a.Seals)
	}
}

func TestGate739FirewallsAndTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.NonPrediction.KappaLambdaDefinedFromRuntimeLedger || a.NonPrediction.IndependentRuntimePrediction || !a.NonPrediction.ConsistencyClosure {
		t.Fatalf("non-prediction firewall failed: %+v", a.NonPrediction)
	}
	if a.HiggsMass.RuntimeLambdaBridgeIsHiggsMassTheorem || a.HiggsMass.HasScalarPotentialTheorem || a.HiggsMass.HasVEVOrScaleTheorem || a.HiggsMass.HasPoleMassCorrectionTheorem || a.HiggsMass.HasUncertaintyPropagation || !a.HiggsMass.HasPhysicalMassConventionFirewall {
		t.Fatalf("Higgs mass firewall failed: %+v", a.HiggsMass)
	}
	res := Generation2Level1ScalarRuntimeBridgeConsistencyEstimateAndNonPredictionAuditTheorem().Verify()
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
