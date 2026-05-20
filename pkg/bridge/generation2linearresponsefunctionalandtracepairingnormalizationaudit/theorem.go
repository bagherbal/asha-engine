package generation2linearresponsefunctionalandtracepairingnormalizationaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2LinearResponseFunctionalAndTracePairingNormalizationAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: AuditID, Name: "Gate 691 — Linear Response Functional and Trace-Pairing Normalization Audit", Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: AuditID, Name: "Gate 691 — Linear Response Functional and Trace-Pairing Normalization Audit", Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate689 first trace and Gate690 residual status", Passed: a.Inherited.Gate689FirstTraceSelectionInherited && a.Inherited.Gate690ResidualStatusInherited && a.Inherited.Operator == "R_split = S_split P_K7" && a.Inherited.H72Dimension == h72Dimension && a.Inherited.K7Dimension == k7Dimension && a.Inherited.QuadraticResidualClueRetained && !a.Inherited.QuadraticCorrectionPromoted && !a.Inherited.NativeSpectralExpansionTheorem && !a.Inherited.NativeFirstTraceTheorem && !a.Inherited.NativeSevenOver72Theorem && strings.Contains(a.Inherited.Verdict, StatusGate689FirstTraceSelectionInherited) && strings.Contains(a.Inherited.Verdict, StatusGate690ResidualStatusInherited), Detail: FormatInheritance(a.Inherited)},
			{Name: "define normalized trace pairing and recover active bridge", Passed: a.Pairing.Definition == "<A,B>_tr,norm = Tr_H72(A B)/Tr_H72(I_H72)" && a.Pairing.Observer == "I_H72" && a.Pairing.Response == "R_split = S_split P_K7" && math.Abs(a.Pairing.NumeratorTrace-7*a.Inherited.SSplit) < pairingTolerance && math.Abs(a.Pairing.DenominatorTrace-72) < pairingTolerance && math.Abs(a.Pairing.Value-a.Inherited.F1) < pairingTolerance && a.Pairing.EqualsFirstTrace && a.Pairing.LinearInResponse && a.Pairing.LinearInSSplit && a.Pairing.BilinearInArguments && strings.Contains(a.Pairing.Verdict, StatusNormalizedTracePairingDefined) && strings.Contains(a.Pairing.Verdict, StatusActiveBridgeRewrittenAsTracePairing), Detail: FormatPairing(a.Pairing)},
			{Name: "classify observer response roles", Passed: a.Roles.FullObserverTypeCorrect && a.Roles.ResponseSupportSelected && a.Roles.BoundaryScalarIsEigenvalue && strings.Contains(a.Roles.FullChamberObserverRole, "I_H72") && strings.Contains(a.Roles.ResponseOperatorRole, "R_split") && strings.Contains(a.Roles.SupportCarrierRole, "P_K7") && strings.Contains(a.Roles.BoundaryScalarRole, "S_split") && strings.Contains(a.Roles.Verdict, StatusObserverResponseRoleClassified) && strings.Contains(a.Roles.Verdict, StatusFullChamberIdentityObserverTypeCorrect), Detail: FormatRoles(a.Roles)},
			{Name: "audit alternative observer pairings", Passed: a.Observers.CandidateCount == 5 && len(a.Observers.Candidates) == 5 && a.Observers.PositiveIdentityOnK7Count == 4 && a.Observers.AllPositiveK7ObserversGiveSameValue && a.Observers.SignedPolarityObserverInactive && !a.Observers.FullH72ObserverUnique && strings.Contains(a.Observers.DegeneracyWarning, "does not uniquely select I_H72") && strings.Contains(a.Observers.Verdict, StatusAlternativeObserverPairingsAudited) && strings.Contains(a.Observers.Verdict, StatusTracePairingDoesNotUniquelySelectH72), Detail: FormatObservers(a.Observers)},
			{Name: "verify observer values", Passed: observerValuesPass(a), Detail: FormatObservers(a.Observers)},
			{Name: "audit linear response order", Passed: a.LinearResponse.DBaseLinearInWallCoordinates && a.LinearResponse.TracePairingLinearInResponse && a.LinearResponse.TracePairingLinearInSSplit && a.LinearResponse.MatchesWallCoordinateOrder && !a.LinearResponse.NativeLinearResponseFunctionalTheorem && strings.Contains(a.LinearResponse.Verdict, StatusActiveBridgeLinearTracePairingResponse) && strings.Contains(a.LinearResponse.Verdict, StatusNoNativeLinearResponseFunctionalTheorem), Detail: FormatLinearResponse(a.LinearResponse)},
			{Name: "retain Gate690 residual as subleading clue only", Passed: math.Abs(a.Residual.E1-8.525834398014336e-10) < residualTolerance && math.Abs(a.Residual.QuadraticF2-1.624013231638281e-7) < 1e-21 && math.Abs(a.Residual.QuadraticCoefficient-0.005249855254820553) < 1e-15 && a.Residual.QuadraticResidualClueRetained && !a.Residual.QuadraticCorrectionPromoted && !a.Residual.NativeSpectralExpansionTheorem && strings.Contains(a.Residual.Verdict, StatusQuadraticResidualRemainsSubleadingClue), Detail: FormatResidual(a.Residual)},
			{Name: "record missing linear response and first trace theorems", Passed: len(a.Missing.Missing) == 3 && strings.Contains(a.Missing.PreciseGap, "LinearResponseFunctionalTheorem") && strings.Contains(a.Missing.PreciseGap, "positive K7-containing observer") && strings.Contains(a.Missing.Verdict, StatusNoNativeLinearResponseFunctionalTheorem) && strings.Contains(a.Missing.Verdict, StatusNoNativeFirstTraceTheorem) && strings.Contains(a.Missing.Verdict, StatusNoNativeSevenOver72Theorem), Detail: FormatMissing(a.Missing)},
			{Name: "preserve Gate691 trace-pairing firewall", Passed: !a.Discipline.ClaimsUniqueFullH72Observer && !a.Discipline.ClaimsNativeLinearResponseTheorem && !a.Discipline.ClaimsNativeFirstTraceTheorem && !a.Discipline.ClaimsNativeSpectralExpansion && !a.Discipline.ClaimsNativeSevenOver72 && !a.Discipline.ClaimsBoundaryStress && !a.Discipline.ClaimsScalarRGMatching && !a.Discipline.ClaimsHiggsMass && !a.Discipline.ClaimsGaugeUnification && !a.Discipline.ClaimsFlavorDerivation && !a.Discipline.ClaimsCKMPMNS && !a.Discipline.ClaimsProjectorActivation && !a.Discipline.PromotesQuadraticResidualCorrection && a.Discipline.Verdict == StatusGate691TracePairingLinearResponseBoundary, Detail: FormatDiscipline(a.Discipline)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: AuditID, Name: "Gate 691 — Linear Response Functional and Trace-Pairing Normalization Audit", Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func observerValuesPass(a Analysis) bool {
	for _, c := range a.Observers.Candidates {
		switch c.Name {
		case "I_H72", "P_finite", "P_kernel", "P_K7":
			if !c.PositiveObserver || !c.ActsAsIdentityOnK7 || !c.EquivalentToActiveFirstTrace || math.Abs(c.Value-a.Inherited.F1) >= pairingTolerance {
				return false
			}
		case "S_K":
			if !c.SignedObserver || c.EquivalentToActiveFirstTrace || math.Abs(c.Value-(1.0/72.0)*a.Inherited.SSplit) >= pairingTolerance {
				return false
			}
		default:
			return false
		}
	}
	return true
}
