package generation2twistorinvariantsu2socketandmovingu1phaseaudit

import (
	"strings"
	"testing"
)

func TestGate714CommonCommutantIntersectionAndMovingPhaseLine(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.TwistorSocketBundleInherited || a.Inherited.K7PlusDimension != 4 || a.Inherited.K7MinusDimension != 3 || a.Inherited.TwistorSphereDimension != 2 || !a.Inherited.FamilyValuedSocketBundle || a.Inherited.SingleSocketPromoted || a.Inherited.NativeTwistorPointSelector {
		t.Fatalf("bad Gate713 inheritance: %+v", a.Inherited)
	}
	if a.CommonCommutant.Dimension != 3 || !a.CommonCommutant.IncludedInAllSockets || !a.CommonCommutant.SelectorIndependent || a.CommonCommutant.PhysicalSU2LCertified || !strings.Contains(a.CommonCommutant.Verdict, StatusCommonCommutantTwistorInvariantSU2Candidate) {
		t.Fatalf("bad common commutant audit: %+v", a.CommonCommutant)
	}
	if !a.Intersection.EqualsCommonCommutant || a.Intersection.Dimension != 3 || !a.Intersection.ProofUsesBasisDirections || a.Intersection.ContainsMovingPhaseLine || !strings.Contains(a.Intersection.Verdict, StatusTwistorIntersectionEqualsCommonCommutant) {
		t.Fatalf("bad twistor intersection audit: %+v", a.Intersection)
	}
	if a.PhaseLine.Dimension != 1 || !a.PhaseLine.MovesWithSelectorN || a.PhaseLine.CommonToAllSockets || a.PhaseLine.SelectorIndependentLine || a.PhaseLine.HyperchargeCertified || !strings.Contains(a.PhaseLine.Verdict, StatusNoSelectorIndependentU1PhaseLine) {
		t.Fatalf("bad moving phase line audit: %+v", a.PhaseLine)
	}
}

func TestGate714SocketSeparationPhysicalFirewallAndTheorem(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.LieAlgebra.Dimension != 3 || !a.LieAlgebra.ClosesAsSU2Like || !a.LieAlgebra.NormalizationRequired || a.LieAlgebra.PhysicalSU2LCertified || !strings.Contains(a.LieAlgebra.Verdict, StatusLieAlgebraStructureOfCommutantAudited) {
		t.Fatalf("bad Lie algebra audit: %+v", a.LieAlgebra)
	}
	if a.Separation.IndependentCount != 3 || a.Separation.DependentCount != 3 || !a.Separation.SeparationValid || !strings.Contains(a.Separation.Verdict, StatusElectroweakAirlockSplits) {
		t.Fatalf("bad separation audit: %+v", a.Separation)
	}
	f := a.PhysicalFirewall
	if f.InternalCommutantPhysicalSU2L || f.MovingPhasePhysicalU1Y || f.HyperchargeNormalization || f.TypedHiggsDoubletMap || f.YukawaOperator || f.YukawaEigenvalues || f.HiggsMass || f.ScalarRuntime || len(f.MissingMaps) != 4 || !strings.Contains(f.Verdict, StatusInternalCommutantNotPhysicalSU2L) || !strings.Contains(f.Verdict, StatusNoHyperchargeAssignmentOrNormalization) {
		t.Fatalf("physical firewall violated: %+v", f)
	}
	if !strings.Contains(a.Strategy.SU2Problem, "SU(2)_L") || !strings.Contains(a.Strategy.U1Problem, "phase line") || !strings.Contains(a.Strategy.Verdict, StatusElectroweakAirlockSplits) {
		t.Fatalf("bad strategy: %+v", a.Strategy)
	}
	res := Generation2TwistorInvariantSU2SocketAndMovingU1PhaseAuditTheorem().Verify()
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
