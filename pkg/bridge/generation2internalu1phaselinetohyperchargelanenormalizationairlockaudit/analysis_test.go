package generation2internalu1phaselinetohyperchargelanenormalizationairlockaudit

import (
	"strings"
	"testing"
)

func TestGate718PhaseLineTargetAndCompatibility(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.MovingPhaseInherited || !a.Inherited.CentralPhaseSocket || !a.Inherited.UniformPhaseAction || !a.Inherited.U1RequiresSelectorAndNorm || a.Inherited.SelectorIndependentU1Line || a.Inherited.NativeTwistorSelector || a.Inherited.HyperchargeAssigned || a.Inherited.HyperchargeNormalized || a.Inherited.FullPhysicalHiggsDoubletMap || a.Inherited.HiggsMassCertified || a.Inherited.YukawaCertified {
		t.Fatalf("bad Gate717 inheritance: %+v", a.Inherited)
	}
	if a.Shape.Dimension != 1 || !strings.Contains(a.Shape.LineDefinition, "L_n") || !strings.Contains(a.Shape.Generator, "q") || !a.Shape.HasCorrectLineShape || a.Shape.PhysicalU1Y {
		t.Fatalf("bad phase-line shape audit: %+v", a.Shape)
	}
	if !a.Uniform.GeneratorActsAsI || !a.Uniform.UniformOnC2 || a.Uniform.ComplexDimension != 2 || !strings.Contains(a.Uniform.ExponentialForm, "q J_H") {
		t.Fatalf("bad uniform phase inheritance: %+v", a.Uniform)
	}
	if !a.Target.TargetLaneIdentified || !a.Target.FiniteSpectralTripleLane || a.Target.TargetComplexDimension != 2 || a.Target.PhysicalIdentityClaimed || !strings.Contains(a.Target.TargetAction, "u(1)_Y") {
		t.Fatalf("bad hypercharge target lane audit: %+v", a.Target)
	}
	if !a.Compatibility.AbelianLieAlgebraTypesMatch || !a.Compatibility.NonzeroNormalizationNeeded || !a.Compatibility.RepresentationCompatible || a.Compatibility.NormalizationNative || !strings.Contains(a.Compatibility.ThetaYMap, "Theta_Y") {
		t.Fatalf("bad U1 compatibility audit: %+v", a.Compatibility)
	}
}

func TestGate718NormalizationSelectorCombinedAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Normalization.CandidateNormalizations) != 4 || !a.Normalization.SameLineDifferentQ || a.Normalization.ChargeUnitFixed || a.Normalization.PhysicalHyperchargeNorm {
		t.Fatalf("normalization firewall violated: %+v", a.Normalization)
	}
	if !a.Selector.PhaseLineDependsOnN || a.Selector.NativeTwistorSelector || a.Selector.CanonicalPhysicalMap || !a.Selector.RequiresSelector {
		t.Fatalf("selector firewall violated: %+v", a.Selector)
	}
	if !a.Combined.SU2SelectorIndependent || !a.Combined.U1SelectorDependent || !a.Combined.U1NormalizationDependent || len(a.Combined.RequiredChoices) != 2 || !a.Combined.FullU2CompatibleAfterNAndQ {
		t.Fatalf("bad combined electroweak airlock status: %+v", a.Combined)
	}
	if a.Physical.LnPhysicalU1Y || a.Physical.JHHyperchargeGenerator || a.Physical.QDerivedHiggsHypercharge || a.Physical.FullPhysicalHiggsDoublet || a.Physical.HiggsMass || a.Physical.ScalarRuntime || a.Physical.YukawaOperator || a.Physical.YukawaEigenvalues || len(a.Physical.MissingMaps) != 3 {
		t.Fatalf("physical firewall violated: %+v", a.Physical)
	}
	res := Generation2InternalU1PhaseLineToHyperchargeLaneNormalizationAirlockAuditTheorem().Verify()
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
