package generation2higgsradialprojectorandgoldstonecomplementorbitaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate725RadialProjectorGeometry(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate724.Inherited || !a.Gate724.RhoPlusDefined || !a.Gate724.RadialEventWeightComputed || !a.Gate724.PhaseLoopExpectationReproducesL || !a.Gate724.NoNativeRadialProjectorSelector {
		t.Fatalf("bad Gate724 inheritance: %+v", a.Gate724)
	}
	if a.Decomposition.RadialRank != 1 || a.Decomposition.AngularRank != 3 || a.Decomposition.CarrierDimension != 4 || !a.Decomposition.DirectSum || !a.Decomposition.Orthogonal {
		t.Fatalf("bad radial decomposition: %+v", a.Decomposition)
	}
	if math.Abs(a.Weights.RadialProbability-0.25) > 1e-18 || math.Abs(a.Weights.AngularProbability-0.75) > 1e-18 || math.Abs(a.Weights.Sum-1) > 1e-18 {
		t.Fatalf("bad event weights: %+v", a.Weights)
	}
	if a.Orbit.U2Dimension != 4 || a.Orbit.StabilizerDimension != 1 || a.Orbit.OrbitDimension != 3 || !a.Orbit.MatchesAngularComplementRank || a.Orbit.PhysicalEWSBTheorem || a.Orbit.PhysicalGoldstoneIdentification {
		t.Fatalf("bad orbit audit: %+v", a.Orbit)
	}
}

func TestGate725FirewallsAndTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Sources.NativeRadialSelectorFound || a.Sources.TwistorSelectorSelectsPRad || a.Sources.QSelectsPRad || a.Sources.ScalarWallSelectsPRad || a.Sources.BoundarySplitSelectsPRad || a.Sources.K7EventProjectorSelectsPRad || a.Sources.FanoHodgeSelectsPRad {
		t.Fatalf("selector source firewall failed: %+v", a.Sources)
	}
	if !a.Seal.TypeDistinctFromTwistorSelector || !a.Seal.TypeDistinctFromHyperchargeNormalization {
		t.Fatalf("seal classification failed: %+v", a.Seal)
	}
	if math.Abs(a.HistoryLoop.HistoryLoopUnit-1/(8*math.Pi)) > 1e-18 || !a.HistoryLoop.ReproducesL || a.HistoryLoop.NativeHistoryLoopUnitSourceTheorem || a.HistoryLoop.NativePhasePayoffTransportTheorem {
		t.Fatalf("bad HistoryLoop relation: %+v", a.HistoryLoop)
	}
	if a.Firewall.NativeRadialProjectorSelector || a.Firewall.TwistorSelectorSelectsPRad || a.Firewall.QSelectsPRad || a.Firewall.NativeEWSBTheorem || a.Firewall.PhysicalGoldstoneIdentification || a.Firewall.NativeHistoryLoopUnitSourceTheorem || a.Firewall.HiggsMassOrPoleMassTheorem || a.Firewall.YukawaOperatorOrEigenvalueTheorem {
		t.Fatalf("firewall failed: %+v", a.Firewall)
	}
	res := Generation2HiggsRadialProjectorAndGoldstoneComplementOrbitAuditTheorem().Verify()
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
