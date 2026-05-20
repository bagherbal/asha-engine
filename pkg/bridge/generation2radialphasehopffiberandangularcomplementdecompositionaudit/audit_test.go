package generation2radialphasehopffiberandangularcomplementdecompositionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate726RadialPhaseHopfGeometry(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate725.Inherited || !a.Gate725.RadialDecompositionDefined || !a.Gate725.RadialEventWeightComputed || !a.Gate725.U2OrbitShadowAudited {
		t.Fatalf("bad Gate725 inheritance: %+v", a.Gate725)
	}
	if !a.Phase.JHSquaresMinusIdentity || !a.Phase.JHSkewOrthogonal || !a.Phase.OrthogonalToRadial || !a.Phase.LiesInAngularComplement || a.Phase.PhaseRank != 1 || !a.Phase.ProjectorOrthogonalToRadial {
		t.Fatalf("bad phase direction audit: %+v", a.Phase)
	}
	if a.Decomposition.RadialRank != 1 || a.Decomposition.PhaseRank != 1 || a.Decomposition.TransverseRank != 2 || a.Decomposition.AngularRank != 3 || a.Decomposition.CarrierDimension != 4 || !a.Decomposition.DirectSum || !a.Decomposition.AngularSplitsOneTwo {
		t.Fatalf("bad radial/phase/transverse decomposition: %+v", a.Decomposition)
	}
	if math.Abs(a.Hopf.PhaseUnit-1/(2*math.Pi)) > 1e-18 || !a.Hopf.UsesRadialEvent || a.Hopf.AbstractPhaseOnly {
		t.Fatalf("bad Hopf fiber audit: %+v", a.Hopf)
	}
}

func TestGate726WeightsOrbitAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Weights.RadialProbability-0.25) > 1e-18 || math.Abs(a.Weights.PhaseProbability-0.25) > 1e-18 || math.Abs(a.Weights.TransverseProbability-0.5) > 1e-18 || math.Abs(a.Weights.TotalProbability-1) > 1e-18 || !a.Weights.HistoryLoopUsesRadialWeight {
		t.Fatalf("bad event weights: %+v", a.Weights)
	}
	if a.Orbit.U2Dimension != 4 || a.Orbit.StabilizerDimension != 1 || a.Orbit.OrbitDimension != 3 || a.Orbit.PhaseFiberDimension != 1 || a.Orbit.ProjectiveTransverseDimension != 2 || !a.Orbit.SplitsAsOnePlusTwo {
		t.Fatalf("bad U2 Hopf orbit audit: %+v", a.Orbit)
	}
	if !a.Selectors.RequiresTwistorSelectorN || !a.Selectors.RequiresRadialProjectorPRad || a.Selectors.NAloneSelectsRadialLine || a.Selectors.PRadAloneSelectsComplexPhase || !a.Selectors.IndependentSeals {
		t.Fatalf("selector firewall failed: %+v", a.Selectors)
	}
	if a.Firewall.NativeRadialProjectorSelector || a.Firewall.NativeTwistorSelectorN || a.Firewall.NativeEWSBTheorem || a.Firewall.PhysicalGoldstoneIdentification || a.Firewall.HopfFiberAsPhysicalTimeOrRG || a.Firewall.NativeHistoryLoopUnitSourceTheorem || a.Firewall.HiggsMassOrPoleMassTheorem || a.Firewall.YukawaOperatorOrEigenvalueTheorem {
		t.Fatalf("physical firewall failed: %+v", a.Firewall)
	}

	res := Generation2RadialPhaseHopfFiberAndAngularComplementDecompositionAuditTheorem().Verify()
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
