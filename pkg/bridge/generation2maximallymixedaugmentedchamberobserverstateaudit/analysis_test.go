package generation2maximallymixedaugmentedchamberobserverstateaudit

import (
	"math"
	"strings"
	"testing"
)

func TestInheritanceAndRho72State(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.TracePairingInherited || a.Inherited.Operator != "R_split = S_split P_K7" || a.Inherited.H72Dimension != h72Dimension || a.Inherited.K7Dimension != k7Dimension {
		t.Fatalf("bad inheritance: %+v", a.Inherited)
	}
	if !a.Inherited.Gate691ObserverDegeneracyRecorded || !a.Inherited.Gate690ResidualClueRetained || a.Inherited.NativeLinearResponseTheorem || a.Inherited.NativeFirstTraceTheorem || a.Inherited.NativeSevenOver72Theorem || a.Inherited.ClaimsUniqueFullH72Observer {
		t.Fatalf("inherited firewall violated: %+v", a.Inherited)
	}
	if a.Rho72.StateName != "rho_72" || !a.Rho72.PositiveState || !a.Rho72.MaximallyMixedOnFullH72 || math.Abs(a.Rho72.StateTrace-1) > stateTolerance {
		t.Fatalf("bad rho72 state: %+v", a.Rho72)
	}
	want := (7.0 / 72.0) * a.Inherited.SSplit
	if math.Abs(a.Rho72.Expectation-want) > stateTolerance || math.Abs(a.Rho72.Expectation-a.Inherited.F1) > stateTolerance || !a.Rho72.EqualsActiveFirstTrace {
		t.Fatalf("bad rho72 expectation: %+v", a.Rho72)
	}
}

func TestAlternativeNormalizedObserverStates(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Alternatives.CandidateCount != 5 || a.Alternatives.PositiveNormalizedStateCount != 4 || a.Alternatives.ActiveStateCount != 1 {
		t.Fatalf("bad alternative counts: %+v", a.Alternatives)
	}
	if !a.Alternatives.FiniteOnlyStateInactive || !a.Alternatives.KernelConditionalStateInactive || !a.Alternatives.LocalK7StateInactive || !a.Alternatives.HodgeSignedObserverNotPositive || !a.Alternatives.AllAlternativesAudited {
		t.Fatalf("bad alternative inactive classification: %+v", a.Alternatives)
	}
	if !observerStateValuesPass(a) {
		t.Fatalf("alternative values failed: %+v", a.Alternatives)
	}
	joined := a.Alternatives.Verdict
	for _, want := range []string{StatusFiniteOnlyStateGives7Over70, StatusKernelConditionalStateGives7Over71, StatusLocalK7StateGivesSSplitNot7Over72, StatusHodgeSignedObserverNotPositiveStateInactive} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing alternative verdict %s in %s", want, joined)
		}
	}
}

func TestDenominatorDegeneracyResolutionAndInterpretation(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Degeneracy.FixedH72DenominatorDegenerate || !a.Degeneracy.StateNormalizationResolvesType || !a.Degeneracy.Rho72UniqueAmongAuditedStates || a.Degeneracy.NativeStateSelectionTheorem {
		t.Fatalf("bad degeneracy resolution: %+v", a.Degeneracy)
	}
	if !strings.Contains(a.Degeneracy.Gate692Resolution, "70, 71, and 7") || !strings.Contains(a.Degeneracy.Verdict, StatusObserverDenominatorDegeneracyResolvedByState) {
		t.Fatalf("degeneracy resolution text/status incomplete: %+v", a.Degeneracy)
	}
	if !a.Interpretation.GlobalAverageDensity || !a.Interpretation.BoundaryScalarEigen || !a.Interpretation.SupportSelectedCarrier {
		t.Fatalf("bad interpretation: %+v", a.Interpretation)
	}
}

func TestResidualAndFirewall(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Residual.E1-8.525834398014336e-10) > residualTolerance || math.Abs(a.Residual.DBase-a.Inherited.DBase) > residualTolerance || math.Abs(a.Residual.Expectation-a.Rho72.Expectation) > stateTolerance {
		t.Fatalf("bad residual ledger: %+v", a.Residual)
	}
	if !a.Residual.QuadraticResidualClueRetained || a.Residual.QuadraticCorrectionPromoted || a.Residual.NativeSpectralExpansionTheorem {
		t.Fatalf("quadratic clue promoted incorrectly: %+v", a.Residual)
	}
	if a.Discipline.ClaimsNativeMaximallyMixedStateTheorem || a.Discipline.ClaimsNativeStateSelectionTheorem || a.Discipline.ClaimsNativeFirstTraceTheorem || a.Discipline.ClaimsNativeSevenOver72Theorem || a.Discipline.PromotesQuadraticResidualCorrection {
		t.Fatalf("discipline firewall violated: %+v", a.Discipline)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2MaximallyMixedAugmentedChamberObserverStateAuditTheorem().Verify()
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
