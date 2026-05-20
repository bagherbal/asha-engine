package generation2historytransportbracketnormalformandscalarwoundcomplementaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate759InheritanceAndHistoryBracket(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate758.Inherited || !a.Gate758.ThreeFactorNormalFormAvailable || a.Gate758.IndependentScalarRuntimeTheorem {
		t.Fatalf("bad Gate758 inheritance: %+v", a.Gate758)
	}
	if math.Abs(a.Gate758.CYukawa-cYukawaMZ) > 1e-15 || math.Abs(a.Gate758.CHistory-cHistoryMZ) > 1e-12 || math.Abs(a.Gate758.LambdaRuntimeEff-lambdaRuntimeEffMZ) > 1e-15 {
		t.Fatalf("bad inherited numerics: %+v", a.Gate758)
	}
	if !a.Bracket.BracketDefined || !a.Bracket.OmegaComputed || a.Bracket.PhysicalTimeOrRGScale {
		t.Fatalf("bad bracket typing: %+v", a.Bracket)
	}
	if math.Abs(a.Bracket.LHopf-lHopf) > 1e-18 || math.Abs(a.Bracket.OmegaHistory-0.9556769569304386) > 1e-12 || math.Abs(a.Bracket.OmegaResidual) > 1e-18 {
		t.Fatalf("bad bracket numerics: %+v", a.Bracket)
	}
}

func TestGate759ReducedDeficitAndNormalForm(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Deficit.Defined || !a.Deficit.RewrittenAsComplement || a.Deficit.NativeScalarTheorem {
		t.Fatalf("bad deficit typing: %+v", a.Deficit)
	}
	if math.Abs(a.Deficit.KappaLambdaRed-0.04432304306956136) > 1e-12 || math.Abs(a.Deficit.Complement-a.Bracket.OmegaHistory) > 1e-15 || math.Abs(a.Deficit.ComplementResidual) > 1e-18 {
		t.Fatalf("bad deficit numerics: %+v", a.Deficit)
	}
	if !a.NormalForm.CHistoryWritten || !a.NormalForm.FullFormRewritten || !a.NormalForm.ThreeFactorNormalForm || a.NormalForm.IndependentRuntimeTheorem {
		t.Fatalf("bad normal form typing: %+v", a.NormalForm)
	}
	if math.Abs(a.NormalForm.CHistory-cHistoryMZ) > 1e-12 || math.Abs(a.NormalForm.LambdaRuntimeFromNormalForm-lambdaRuntimeEffMZ) > 1e-15 || math.Abs(a.NormalForm.NormalFormResidual) > 1e-15 {
		t.Fatalf("bad normal form numerics: %+v", a.NormalForm)
	}
}

func TestGate759InterpretationLayersAndIllegalRejection(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Interpretation.Recorded || a.Interpretation.NativeHistoryLoopTheorem || a.Interpretation.NativeTransportTheorem {
		t.Fatalf("bad interpretation: %+v", a.Interpretation)
	}
	if !a.Layers.LayerSeparationAudited || !a.Layers.FactorsMultiplyAfterScalarCollapse || a.Layers.OperatorsOnSameNativeBoard {
		t.Fatalf("bad layer audit: %+v", a.Layers)
	}
	if !a.Illegal.Audited || a.Illegal.KappaLambdaRedNativeScalarTheorem || a.Illegal.OmegaHistoryPhysicalTimeOrRGScale || a.Illegal.LHopfBoundaryEventProbability || a.Illegal.CHistoryNativeHistoryLoopTheorem || a.Illegal.LambdaRuntimeEffIndependentPrediction || a.Illegal.TreeProxyPoleMassPrediction || a.Illegal.ClaimsYukawaEigenvaluesDerived || a.Illegal.ClaimsHiggsMassOrPoleMassTheorem {
		t.Fatalf("bad illegal rejection firewall: %+v", a.Illegal)
	}
}

func TestGate759TheoremVerdictStatuses(t *testing.T) {
	res := Generation2HistoryTransportBracketNormalFormAndScalarWoundComplementAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status note %s", want)
		}
	}
}
