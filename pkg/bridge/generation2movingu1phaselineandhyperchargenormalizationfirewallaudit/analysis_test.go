package generation2movingu1phaselineandhyperchargenormalizationfirewallaudit

import (
	"strings"
	"testing"
)

func TestGate717MovingU1PhaseLineAndCentrality(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.SU2AirlockInherited || !a.Inherited.SU2SideStructurallyReady || !a.Inherited.InternalCSelectorIndependent || !a.Inherited.ComplexCarrierSelectorDependent || !a.Inherited.U1PhaseSelectorDependent || a.Inherited.HyperchargeDerived || a.Inherited.HyperchargeNormalized || a.Inherited.FullHiggsDoubletMap || a.Inherited.HiggsMassCertified || a.Inherited.YukawaCertified {
		t.Fatalf("bad Gate716 inheritance: %+v", a.Inherited)
	}
	if a.PhaseLine.Dimension != 1 || !a.PhaseLine.DependsOnSelectorN || !a.PhaseLine.FixedJHRequired || a.PhaseLine.SelectorIndependent || a.PhaseLine.PhysicalHypercharge || !strings.Contains(a.PhaseLine.Definition, "span") {
		t.Fatalf("bad phase line audit: %+v", a.PhaseLine)
	}
	if !a.Central.CommutesWithC || !a.Central.LiesInCenterOfU2 || !a.Central.FixedJHOnly || a.Central.PhysicalU1Y || !strings.Contains(a.Central.Commutator, "[J_H") {
		t.Fatalf("bad central phase-line audit: %+v", a.Central)
	}
	if !a.Uniform.ActsAsMultiplicationByI || !a.Uniform.UniformOnFullC2 || a.Uniform.ComplexDimension != 2 || a.Uniform.PhysicalChargeFixed || !strings.Contains(a.Uniform.ExponentialAction, "exp") {
		t.Fatalf("bad uniform phase action: %+v", a.Uniform)
	}
}

func TestGate717NormalizationSelectorAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Charge.PhaseLineFixed || a.Charge.NaturalDirection != "J_H(n)" || len(a.Charge.CandidateNormalizations) != 3 || !a.Charge.SameLineDifferentCharges || a.Charge.PhysicalHyperchargeNormalization || !a.Charge.ThetaYRequired {
		t.Fatalf("bad charge normalization audit: %+v", a.Charge)
	}
	if !a.Selector.PhaseLineDependsOnN || a.Selector.NativeTwistorPointSelector || a.Selector.SelectorIndependentU1Line || a.Selector.ComplexStructureSelected {
		t.Fatalf("selector dependence firewall violated: %+v", a.Selector)
	}
	if !a.Asymmetry.SU2SelectorIndependent || !a.Asymmetry.U1SelectorDependent || !a.Asymmetry.U1NormalizationOpen || !strings.Contains(a.Asymmetry.SU2Side, "twistor-invariant") {
		t.Fatalf("bad SU2/U1 asymmetry audit: %+v", a.Asymmetry)
	}
	if a.Physical.LnPhysicalU1Y || a.Physical.JHHyperchargeGenerator || a.Physical.InternalPhaseChargePhysicalHiggsHypercharge || a.Physical.FullPhysicalHiggsDoublet || a.Physical.HyperchargeAssignment || a.Physical.HyperchargeNormalization || a.Physical.HiggsMass || a.Physical.ScalarRuntime || a.Physical.YukawaOperator || a.Physical.YukawaEigenvalues || len(a.Physical.MissingMaps) != 3 {
		t.Fatalf("physical hypercharge firewall violated: %+v", a.Physical)
	}
	res := Generation2MovingU1PhaseLineAndHyperchargeNormalizationFirewallAuditTheorem().Verify()
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
