package generation2scalarhiggsbridgesealinventoryandforecastboundaryaudit

import (
	"strings"
	"testing"
)

func TestGate735SealInventoryAndClassification(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate734.Inherited || !a.Gate734.ConsistencyClosure || a.Gate734.IndependentPrediction || !a.Gate734.NoNativeHistoryLoopUnit || !a.Gate734.NoNativeBoundaryGeneratingTheorem {
		t.Fatalf("bad Gate734 inheritance: %+v", a.Gate734)
	}
	if a.SealInventory.NonNativeOrSealedCount != 9 || !a.SealInventory.IncludesN || !a.SealInventory.IncludesQ || !a.SealInventory.IncludesPRad || !a.SealInventory.IncludesRhoPlus || !a.SealInventory.IncludesRho72 || !a.SealInventory.IncludesKappaE || !a.SealInventory.IncludesLambdaProxy || !a.SealInventory.IncludesHistoryLoopUnit || !a.SealInventory.IncludesFWall3 {
		t.Fatalf("bad seal inventory: %+v", a.SealInventory)
	}
	if !a.Classification.Separated || a.Classification.NativeCount < 5 || a.Classification.BridgeCount < 9 {
		t.Fatalf("bad native/bridge classification: %+v", a.Classification)
	}
	if !strings.Contains(a.Classification.BridgeObjects[0], "n") || !strings.Contains(strings.Join(a.Classification.BridgeObjects, " "), "lambda_proxy") {
		t.Fatalf("classification missing expected bridge objects: %+v", a.Classification)
	}
}

func TestGate735RequirementsForecastAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.RuntimeRequirements.Complete || a.RuntimeRequirements.NativeNow || len(a.RuntimeRequirements.Requirements) != 8 {
		t.Fatalf("bad scalar runtime requirements: %+v", a.RuntimeRequirements)
	}
	if !a.HiggsMassRequirements.Complete || a.HiggsMassRequirements.NativeNow || len(a.HiggsMassRequirements.Requirements) != 6 {
		t.Fatalf("bad Higgs requirements: %+v", a.HiggsMassRequirements)
	}
	if a.Forecast.Level0NativeTheoremAllowed || !a.Forecast.Level1BridgeConsistencyEstimateAllowed || a.Forecast.Level2PhysicalPredictionAllowed {
		t.Fatalf("bad forecast levels: %+v", a.Forecast)
	}
	if a.Firewall.ClaimsCubicBridgeIsHiggsMassTheorem || a.Firewall.ClaimsLambdaProxyIsPoleMassTheorem || a.Firewall.ClaimsLNativeLoopTheorem || a.Firewall.ClaimsPRadDerivedVacuum || a.Firewall.ClaimsKappaENativeFlavorTheorem || a.Firewall.ClaimsFWall3NativeGeneratingFunction || a.Firewall.ClaimsSealedHiggsSocketPhysicalScalarLaw {
		t.Fatalf("physical firewall failed: %+v", a.Firewall)
	}

	res := Generation2ScalarHiggsBridgeSealInventoryAndForecastBoundaryAuditTheorem().Verify()
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
