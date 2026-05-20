package generation2fullaugmentedobserverstateselectionandbiasfirewallaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2FullAugmentedObserverStateSelectionAndBiasFirewallAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: AuditID, Name: "Gate 693 — Full Augmented Observer State Selection and Bias Firewall Audit", Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: AuditID, Name: "Gate 693 — Full Augmented Observer State Selection and Bias Firewall Audit", Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate692 state expectation", Passed: a.Inherited.StateExpectationInherited && a.Inherited.ResponseOperator == "R_split = S_split P_K7" && a.Inherited.H72Dimension == h72Dimension && a.Inherited.K7Dimension == k7Dimension && a.Inherited.Gate692Rho72TypeCorrect && a.Inherited.Gate692NoNativeStateSelection && a.Inherited.Gate692NoNativeFirstTraceTheorem && a.Inherited.Gate692NoNativeSevenOver72 && a.Inherited.Verdict == StatusGate692StateExpectationInherited, Detail: FormatInheritance(a.Inherited)},
			{Name: "reduce general state response to K7 weight", Passed: a.General.GeneralPositiveState && a.General.ReducesToK7Weight && a.General.ActiveRequiresK7Weight && a.General.DoesNotRequireFullRho72 && a.General.WarnsAboutStateDegeneracy && math.Abs(a.General.RequiredK7Weight-7.0/72.0) < tolerance && math.Abs(a.General.RequiredExpectation-a.Inherited.ActiveExpectation) < tolerance && strings.Contains(a.General.Verdict, StatusGeneralStateResponseReducedToK7Weight), Detail: FormatGeneral(a.General)},
			{Name: "audit typed observer-state alternatives", Passed: a.Alternatives.CandidateCount == 7 && len(a.Alternatives.Candidates) == 7 && a.Alternatives.PositiveNormalizedCount == 6 && a.Alternatives.MatchingActiveBridgeCount == 2 && a.Alternatives.UnbiasedMatchingCount == 1 && a.Alternatives.BiasedMatchingCount == 1 && a.Alternatives.FiniteOnlyRejected && a.Alternatives.KernelRejected && a.Alternatives.LocalK7Rejected && a.Alternatives.BoundaryOnlyRejected && a.Alternatives.HodgeSignedRejected && a.Alternatives.BiasedReproductionWitnessed && a.Alternatives.Rho72ActiveUnbiasedCandidate && a.Alternatives.AllTypedAlternativesAudited && strings.Contains(a.Alternatives.Verdict, StatusAlternativeTypedStatesAudited), Detail: FormatAlternatives(a.Alternatives)},
			{Name: "verify candidate K7 weights and responses", Passed: candidateWeightsPass(a), Detail: FormatAlternatives(a.Alternatives)},
			{Name: "classify rho_72 as minimal unbiased full H72 state", Passed: math.Abs(a.Rho72Selection.Rho72K7Weight-7.0/72.0) < tolerance && math.Abs(a.Rho72Selection.ActiveK7Weight-7.0/72.0) < tolerance && math.Abs(a.Rho72Selection.Rho72Expectation-a.Inherited.ActiveExpectation) < tolerance && a.Rho72Selection.FullSupport && a.Rho72Selection.Positive && a.Rho72Selection.Normalized && a.Rho72Selection.Unbiased && len(a.Rho72Selection.MinimalAssumptions) == 5 && a.Rho72Selection.UniqueUnderUnbiasedFullH72 && !a.Rho72Selection.UniqueAmongAllDensityStates && !a.Rho72Selection.NativeStateSelectionTheorem && strings.Contains(a.Rho72Selection.Verdict, StatusRho72MinimalUnbiasedFullAugmentedObserverState), Detail: FormatRho72Selection(a.Rho72Selection)},
			{Name: "preserve biased-state reproduction firewall", Passed: a.BiasFirewall.BiasedDensityStatesCanMatch && a.BiasFirewall.BiasedWitnessName == "rho_biased_weight_7_over_72" && math.Abs(a.BiasFirewall.BiasedWitnessK7Weight-7.0/72.0) < tolerance && math.Abs(a.BiasFirewall.BiasedWitnessExpectation-a.Inherited.ActiveExpectation) < tolerance && a.BiasFirewall.BiasedWitnessCircular && !a.BiasFirewall.ReproductionIsNativeSelection && a.BiasFirewall.Rho72UniquenessOverclaimed && strings.Contains(a.BiasFirewall.Verdict, StatusRho72NotUniqueAmongAllDensityStates), Detail: FormatBiasFirewall(a.BiasFirewall)},
			{Name: "retain residual as subleading clue only", Passed: math.Abs(a.Residual.E1-8.525834398014336e-10) < tolerance && math.Abs(a.Residual.Expectation-a.Inherited.ActiveExpectation) < tolerance && math.Abs(a.Residual.DBase-a.Inherited.DBase) < tolerance && a.Residual.QuadraticResidualClueRetained && !a.Residual.QuadraticCorrectionPromoted, Detail: FormatResidual(a.Residual)},
			{Name: "record missing state-selection, first-trace, and 7/72 theorems", Passed: len(a.Missing.Missing) == 3 && len(a.Missing.Candidates) == 3 && strings.Contains(a.Missing.PreciseGap, "state-selection") && strings.Contains(a.Missing.PreciseGap, "biased synthetic") && strings.Contains(a.Missing.Verdict, StatusNoNativeMaximallyMixedStateSelectionTheorem) && strings.Contains(a.Missing.Verdict, StatusNoNativeFirstTraceTheorem) && strings.Contains(a.Missing.Verdict, StatusNoNativeSevenOver72Theorem), Detail: FormatMissing(a.Missing)},
			{Name: "preserve Gate693 observer-state selection firewall", Passed: !a.Discipline.ClaimsNativeMaximallyMixedStateSelection && !a.Discipline.ClaimsNativeFirstTraceTheorem && !a.Discipline.ClaimsNativeSevenOver72Theorem && !a.Discipline.ClaimsRho72UniqueAmongAllStates && !a.Discipline.ClaimsBiasedStateNativeSelection && !a.Discipline.ClaimsBoundaryStress && !a.Discipline.ClaimsScalarRGMatching && !a.Discipline.ClaimsHiggsMass && !a.Discipline.ClaimsGaugeUnification && !a.Discipline.ClaimsFlavorDerivation && !a.Discipline.ClaimsCKMPMNS && !a.Discipline.ClaimsProjectorActivation && a.Discipline.Verdict == StatusGate693ObserverStateSelectionBoundary, Detail: FormatDiscipline(a.Discipline)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: AuditID, Name: "Gate 693 — Full Augmented Observer State Selection and Bias Firewall Audit", Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func candidateWeightsPass(a Analysis) bool {
	for _, c := range a.Alternatives.Candidates {
		switch c.Name {
		case "rho_72":
			if !c.PositiveState || !c.NormalizedState || !c.FullH72Support || !c.Unbiased || c.Biased || !c.MatchesActiveBridge || math.Abs(c.K7Weight-7.0/72.0) >= tolerance || math.Abs(c.Expectation-a.Inherited.ActiveExpectation) >= tolerance {
				return false
			}
		case "rho_finite":
			if !c.PositiveState || !c.NormalizedState || c.MatchesActiveBridge || math.Abs(c.K7Weight-7.0/70.0) >= tolerance || c.FailureStatus != StatusFiniteOnlyStateRejectedBy7Over70 {
				return false
			}
		case "rho_kernel":
			if !c.PositiveState || !c.NormalizedState || c.MatchesActiveBridge || math.Abs(c.K7Weight-7.0/71.0) >= tolerance || c.FailureStatus != StatusKernelStateRejectedBy7Over71 {
				return false
			}
		case "rho_K7":
			if !c.PositiveState || !c.NormalizedState || c.MatchesActiveBridge || math.Abs(c.K7Weight-1.0) >= tolerance || c.FailureStatus != StatusLocalK7StateRejectedByUnitWeight {
				return false
			}
		case "rho_boundary":
			if !c.PositiveState || !c.NormalizedState || c.MatchesActiveBridge || math.Abs(c.K7Weight) >= tolerance || c.FailureStatus != StatusBoundaryOnlyStateRejectedByZeroWeight {
				return false
			}
		case "rho_signed":
			if c.PositiveState || c.NormalizedState || !c.SignedObserver || c.MatchesActiveBridge || math.Abs(c.K7Weight-1.0/72.0) >= tolerance || c.FailureStatus != StatusHodgeSignedObserverRejectedAsNonPositiveState {
				return false
			}
		case "rho_biased_weight_7_over_72":
			if !c.PositiveState || !c.NormalizedState || !c.FullH72Support || !c.Biased || !c.Circular || !c.MatchesActiveBridge || math.Abs(c.K7Weight-7.0/72.0) >= tolerance || c.FailureStatus != StatusBiasedStatesCanReproduceWeightButAreCircular {
				return false
			}
		default:
			return false
		}
	}
	return true
}
