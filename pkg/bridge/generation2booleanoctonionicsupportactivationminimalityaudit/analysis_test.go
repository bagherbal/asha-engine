package generation2booleanoctonionicsupportactivationminimalityaudit

import (
	"strings"
	"testing"
)

func TestGate686Build(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.BooleanOctonionicSelection || a.Inherited.SelectedProjector != "P_K7" || !a.Inherited.ActivationStillUnproved {
		t.Fatalf("bad Gate685 inheritance: %+v", a.Inherited)
	}
	if len(a.Ladder.Steps) != 5 || !a.Ladder.CombinedSupportSelectsK7 || !a.Ladder.MinimalPairRequired {
		t.Fatalf("bad selector ladder: %+v", a.Ladder)
	}
	if !a.Decomposition.SupportSelectorSelectsPK7 || a.Decomposition.SSplitAloneSelectsProjector || a.Decomposition.NativeActivationProved {
		t.Fatalf("bad activation decomposition: %+v", a.Decomposition)
	}
}

func TestConstraintLadderMinimality(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ladder.RankOnlyDegenerate || !a.Ladder.FiniteSupportOnlyDegenerate || !a.Ladder.BooleanOnlyDegenerate || !a.Ladder.OctonionicOnlyDegenerate {
		t.Fatalf("weaker selectors should remain degenerate: %+v", a.Ladder)
	}
	last := a.Ladder.Steps[len(a.Ladder.Steps)-1]
	if last.Carrier != "U∩V=K_7" || last.CarrierDimension != 7 || !last.UniquePK7 || last.Degenerate {
		t.Fatalf("combined support should uniquely select K7: %+v", last)
	}
}

func TestSupportIndependence(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Independence.BooleanComplementDimension != 49 || a.Independence.OctonionicComplementDimension != 7 {
		t.Fatalf("bad complement dimensions: %+v", a.Independence)
	}
	if a.Independence.BooleanImpliesOctonionic || a.Independence.OctonionicImpliesBoolean || !a.Independence.NeitherConditionRedundant || !a.Independence.BothRequiredToForceK7 {
		t.Fatalf("support conditions should be independent and both required: %+v", a.Independence)
	}
}

func TestNoncircularityAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Noncircular.Noncircular || !a.Noncircular.DoesNotAssumePK7 || !a.Noncircular.ConditionalNotAbsolute {
		t.Fatalf("selection proof must be noncircular and conditional: %+v", a.Noncircular)
	}
	if a.Discipline.ClaimsSSplitSelectsProjector || a.Discipline.ClaimsBoundaryScalarActivatesSieve || a.Discipline.ClaimsProjectorActivation || a.Discipline.ClaimsNativeSevenOver72 {
		t.Fatalf("firewall violation: %+v", a.Discipline)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2BooleanOctonionicSupportActivationMinimalityAuditTheorem().Verify()
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
