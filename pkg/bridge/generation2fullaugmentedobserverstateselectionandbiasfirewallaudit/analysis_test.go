package generation2fullaugmentedobserverstateselectionandbiasfirewallaudit

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
	if !a.Inherited.StateExpectationInherited || a.Inherited.ResponseOperator != "R_split = S_split P_K7" || a.Inherited.H72Dimension != h72Dimension || a.Inherited.K7Dimension != k7Dimension {
		t.Fatalf("bad Gate692 inheritance: %+v", a.Inherited)
	}
	if !a.Inherited.Gate692Rho72TypeCorrect || !a.Inherited.Gate692NoNativeStateSelection || !a.Inherited.Gate692NoNativeFirstTraceTheorem || !a.Inherited.Gate692NoNativeSevenOver72 {
		t.Fatalf("Gate692 firewall not inherited: %+v", a.Inherited)
	}
	if !a.General.ReducesToK7Weight || !a.General.ActiveRequiresK7Weight || !a.General.DoesNotRequireFullRho72 || !a.General.WarnsAboutStateDegeneracy {
		t.Fatalf("bad general state-response audit: %+v", a.General)
	}
	if math.Abs(a.General.RequiredK7Weight-7.0/72.0) > tolerance || math.Abs(a.General.RequiredExpectation-a.Inherited.ActiveExpectation) > tolerance {
		t.Fatalf("wrong active K7 weight: %+v", a.General)
	}
}

func TestTypedStateAlternativesAndBiasWitness(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Alternatives.CandidateCount != 7 || a.Alternatives.PositiveNormalizedCount != 6 || a.Alternatives.MatchingActiveBridgeCount != 2 || a.Alternatives.UnbiasedMatchingCount != 1 || a.Alternatives.BiasedMatchingCount != 1 {
		t.Fatalf("bad alternative counts: %+v", a.Alternatives)
	}
	if !a.Alternatives.FiniteOnlyRejected || !a.Alternatives.KernelRejected || !a.Alternatives.LocalK7Rejected || !a.Alternatives.BoundaryOnlyRejected || !a.Alternatives.HodgeSignedRejected || !a.Alternatives.BiasedReproductionWitnessed || !a.Alternatives.Rho72ActiveUnbiasedCandidate || !a.Alternatives.AllTypedAlternativesAudited {
		t.Fatalf("bad alternative classification: %+v", a.Alternatives)
	}
	if !candidateWeightsPass(a) {
		t.Fatalf("candidate weights failed: %+v", a.Alternatives)
	}
	for _, want := range []string{StatusFiniteOnlyStateRejectedBy7Over70, StatusKernelStateRejectedBy7Over71, StatusLocalK7StateRejectedByUnitWeight, StatusBoundaryOnlyStateRejectedByZeroWeight, StatusHodgeSignedObserverRejectedAsNonPositiveState, StatusBiasedStatesCanReproduceWeightButAreCircular} {
		if !strings.Contains(a.Alternatives.Verdict, want) {
			t.Fatalf("missing alternative verdict %s in %s", want, a.Alternatives.Verdict)
		}
	}
}

func TestRho72SelectionAndBiasFirewall(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Rho72Selection.Rho72K7Weight-7.0/72.0) > tolerance || math.Abs(a.Rho72Selection.Rho72Expectation-a.Inherited.ActiveExpectation) > tolerance {
		t.Fatalf("bad rho72 selection values: %+v", a.Rho72Selection)
	}
	if !a.Rho72Selection.FullSupport || !a.Rho72Selection.Positive || !a.Rho72Selection.Normalized || !a.Rho72Selection.Unbiased || len(a.Rho72Selection.MinimalAssumptions) != 5 {
		t.Fatalf("bad rho72 assumptions: %+v", a.Rho72Selection)
	}
	if !a.Rho72Selection.UniqueUnderUnbiasedFullH72 || a.Rho72Selection.UniqueAmongAllDensityStates || a.Rho72Selection.NativeStateSelectionTheorem {
		t.Fatalf("rho72 uniqueness firewall violated: %+v", a.Rho72Selection)
	}
	if !a.BiasFirewall.BiasedDensityStatesCanMatch || a.BiasFirewall.BiasedWitnessName != "rho_biased_weight_7_over_72" || math.Abs(a.BiasFirewall.BiasedWitnessK7Weight-7.0/72.0) > tolerance || math.Abs(a.BiasFirewall.BiasedWitnessExpectation-a.Inherited.ActiveExpectation) > tolerance || !a.BiasFirewall.BiasedWitnessCircular || a.BiasFirewall.ReproductionIsNativeSelection || !a.BiasFirewall.Rho72UniquenessOverclaimed {
		t.Fatalf("bad bias firewall: %+v", a.BiasFirewall)
	}
}

func TestResidualMissingAndDiscipline(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Residual.E1-8.525834398014336e-10) > tolerance || math.Abs(a.Residual.Expectation-a.Inherited.ActiveExpectation) > tolerance || !a.Residual.QuadraticResidualClueRetained || a.Residual.QuadraticCorrectionPromoted {
		t.Fatalf("bad residual status: %+v", a.Residual)
	}
	if len(a.Missing.Missing) != 3 || len(a.Missing.Candidates) != 3 || !strings.Contains(a.Missing.PreciseGap, "biased synthetic") || !strings.Contains(a.Missing.Verdict, StatusNoNativeMaximallyMixedStateSelectionTheorem) {
		t.Fatalf("bad missing-theorem ledger: %+v", a.Missing)
	}
	if a.Discipline.ClaimsNativeMaximallyMixedStateSelection || a.Discipline.ClaimsNativeFirstTraceTheorem || a.Discipline.ClaimsNativeSevenOver72Theorem || a.Discipline.ClaimsRho72UniqueAmongAllStates || a.Discipline.ClaimsBiasedStateNativeSelection || a.Discipline.ClaimsBoundaryStress || a.Discipline.ClaimsScalarRGMatching || a.Discipline.ClaimsHiggsMass || a.Discipline.ClaimsGaugeUnification || a.Discipline.ClaimsFlavorDerivation || a.Discipline.ClaimsCKMPMNS || a.Discipline.ClaimsProjectorActivation {
		t.Fatalf("discipline firewall violated: %+v", a.Discipline)
	}
}

func TestTheorem(t *testing.T) {
	res := Generation2FullAugmentedObserverStateSelectionAndBiasFirewallAuditTheorem().Verify()
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
