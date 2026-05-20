package generation2maximumentropyobserverstateselectionaudit

import (
	"math"
	"strings"
	"testing"
)

func TestInheritanceAndGeneralStateResponse(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.ObserverStateSelectionInherited || a.Inherited.Rho72Definition != "rho_72 = I_H72/72" || a.Inherited.ResponseOperator != "R_split = S_split P_K7" || a.Inherited.H72Dimension != h72Dimension || a.Inherited.K7Dimension != k7Dimension {
		t.Fatalf("bad Gate693 inheritance: %+v", a.Inherited)
	}
	if !a.Inherited.Gate693MinimalUnbiased || !a.Inherited.Gate693BiasedStatesCanMatch || !a.Inherited.Gate693NoNativeStateSelection || !a.Inherited.Gate693NoNativeFirstTrace || !a.Inherited.Gate693NoNativeSevenOver72 {
		t.Fatalf("Gate693 firewall not inherited: %+v", a.Inherited)
	}
	if !a.General.ReducesToK7Weight || !a.General.ActiveRequiresK7Weight || math.Abs(a.General.RequiredK7Weight-7.0/72.0) > tolerance || math.Abs(a.General.RequiredExpectation-a.Inherited.ActiveExpectation) > tolerance {
		t.Fatalf("bad general response: %+v", a.General)
	}
}

func TestMaximumEntropyAndSymmetryNoBias(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Entropy.Dimension != h72Dimension || math.Abs(a.Entropy.Rho72Eigenvalue-1.0/72.0) > tolerance || math.Abs(a.Entropy.Rho72Entropy-math.Log(72.0)) > tolerance || !a.Entropy.StrictConcavityUsed || !a.Entropy.Rho72UniqueMaximumEntropy {
		t.Fatalf("bad entropy audit: %+v", a.Entropy)
	}
	if !a.Entropy.ExampleBiasedEntropyLowerThanMax || !a.Entropy.AnyBiasedStateEntropyLower || !(a.Entropy.ExampleBiasedEntropy < a.Entropy.MaximumEntropy) {
		t.Fatalf("biased entropy firewall failed: %+v", a.Entropy)
	}
	if a.Symmetry.InvariantStateForm != "rho = c I_H72" || math.Abs(a.Symmetry.ScalarCoefficient-1.0/72.0) > tolerance || !a.Symmetry.SelectsRho72 || !a.Symmetry.EquivalentToNoDirectionBias {
		t.Fatalf("bad symmetry audit: %+v", a.Symmetry)
	}
}

func TestBlockBiasFamilyAndBiasFirewall(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.BlockBias.Family != "rho(a,b)=a P_finite + b P_boundary" || a.BlockBias.NormalizationEquation != "70a+2b=1" || a.BlockBias.K7WeightFormula != "Tr(rho(a,b) P_K7)=7a" {
		t.Fatalf("bad block family: %+v", a.BlockBias)
	}
	if math.Abs(a.BlockBias.SolvedA-1.0/72.0) > tolerance || math.Abs(a.BlockBias.SolvedB-1.0/72.0) > tolerance || !a.BlockBias.EqualPerDimensionWeight || !a.BlockBias.FiniteBoundaryBiasRejected || !a.BlockBias.Rho72SelectedInBlockFamily {
		t.Fatalf("block family did not select rho72: %+v", a.BlockBias)
	}
	if !a.Bias.BiasedDensityStatesCanMatch || math.Abs(a.Bias.BiasedWitnessK7Weight-7.0/72.0) > tolerance || math.Abs(a.Bias.BiasedWitnessExpectation-a.Inherited.ActiveExpectation) > tolerance || !a.Bias.BiasedWitnessCircular || a.Bias.ReproductionIsNativeSelection || !a.Bias.PreservesGate693Firewall {
		t.Fatalf("bias firewall failed: %+v", a.Bias)
	}
}

func TestResponseMissingAndDiscipline(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Response.Rho72K7Weight-7.0/72.0) > tolerance || math.Abs(a.Response.Expectation-a.Inherited.ActiveExpectation) > tolerance || math.Abs(a.Response.ResidualE1-8.525834398014336e-10) > tolerance {
		t.Fatalf("bad response value: %+v", a.Response)
	}
	if len(a.Missing.Missing) != 3 || len(a.Missing.Candidates) != 3 || !strings.Contains(a.Missing.PreciseGap, "native physical-history") || !strings.Contains(a.Missing.Verdict, StatusNoNativeMaximumEntropyHistoryObserverTheorem) {
		t.Fatalf("bad missing theorem audit: %+v", a.Missing)
	}
	if a.Discipline.ClaimsPhysicalHistoryUsesMaxEntropy || a.Discipline.ClaimsNativeStateSelectionTheorem || a.Discipline.ClaimsNativeFirstTraceTheorem || a.Discipline.ClaimsNativeSevenOver72Theorem || a.Discipline.ClaimsBiasedStateNativeSelection || a.Discipline.ClaimsBoundaryStress || a.Discipline.ClaimsScalarRGMatching || a.Discipline.ClaimsHiggsMass || a.Discipline.ClaimsGaugeUnification || a.Discipline.ClaimsFlavorDerivation || a.Discipline.ClaimsCKMPMNS || a.Discipline.ClaimsProjectorActivation {
		t.Fatalf("discipline firewall violated: %+v", a.Discipline)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2MaximumEntropyObserverStateSelectionAuditTheorem().Verify()
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
