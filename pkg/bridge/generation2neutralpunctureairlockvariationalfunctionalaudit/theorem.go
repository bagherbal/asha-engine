package generation2neutralpunctureairlockvariationalfunctionalaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GENERATION2_NEUTRAL_PUNCTURE_AIRLOCK_VARIATIONAL_FUNCTIONAL_AUDIT"
	theoremName = "Gate 896 — NeutralPuncture Airlock Variational Functional Audit"
)

func Generation2NeutralPunctureAirlockVariationalFunctionalAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 895 neutral puncture airlock inherited", Passed: containsAll(Statuses(), []string{StatusGate895Inherited}) && a.Functional.Formulated, Detail: FormatFunctional(a.Functional)},
			{Name: "rank-one puncture candidates reduce to lepton socket cells but do not select plus", Passed: a.Rank.OnlyLeptonCellsAreRankOne && a.Rank.BothLeptonCandidatesPass && !a.Rank.SelectsEPlusUniquely && containsAll(a.Rank.Failures, []string{FailureRankDoesNotSelectPlus}), Detail: FormatRankTerm(a.Rank)},
			{Name: "alpha flag ranks reconstruct alpha for both plus and minus punctures", Passed: a.AlphaFlag.BothReconstructAlphaShape && !a.AlphaFlag.DistinguishesPlusFromMinus && near(a.AlphaFlag.PlusAlpha, AlphaB) && near(a.AlphaFlag.MinusAlpha, AlphaB) && containsAll(a.AlphaFlag.Failures, []string{FailureAlphaFlagDoesNotSelectPlus}), Detail: FormatAlphaFlagTerm(a.AlphaFlag)},
			{Name: "current oriented edge pattern selects e_+ puncture only circularly", Passed: a.Edge.SelectsEPlusAsNullEdge && !a.Edge.IndependentEdgeOrdering && a.Edge.CircularWithoutOrdering && containsAll(a.Edge.Failures, []string{FailureEdgeSelectionCircular, FailureNoNativeOrientedEdgeOrdering}), Detail: FormatEdgeSupportTerm(a.Edge)},
			{Name: "left kernel matches e_+ puncture but depends on preselected image", Passed: a.Kernel.MatchesEPlusPuncture && a.Kernel.DependsOnPreselectedImage && !a.Kernel.NativeKernelSelector && containsAll(a.Kernel.Failures, []string{FailureLeftKernelDependsOnImage}), Detail: FormatLeftKernelTerm(a.Kernel)},
			{Name: "B-L compensation does not distinguish plus from minus lepton puncture", Passed: a.BMinusL.FullRectangleNeutral && !a.BMinusL.DistinguishesPlusFromMinus && containsAll(a.BMinusL.Failures, []string{FailureBMinusLDoesNotSelectPlus}), Detail: FormatBMinusLTerm(a.BMinusL)},
			{Name: "airlock functional reduces obstruction to oriented edge ordering, not native theorem", Passed: a.Functional.Formulated && a.Functional.EPlusSatisfiesAllTerms && a.Functional.EMinusAlsoSatisfiesRankFlagBL && a.Functional.RequiresOrientedEdgeOrdering && !a.Functional.NativeFunctional && containsAll(a.Functional.Failures, []string{FailureNoNativeAirlockFunctional, FailureNoNativeOrientedEdgeOrdering}), Detail: FormatFunctional(a.Functional)},
			{Name: "official diagnostics remain frozen", Passed: a.Freeze.Frozen && a.Freeze.DiagnosticOnly && !a.Freeze.CanUpdate && near(a.Freeze.OperatorNEff, OperatorNEffDiagnostic) && !near(a.Freeze.OperatorNEff, a.Freeze.OfficialNEff), Detail: FormatFreeze(a.Freeze)},
			{Name: "native R3/R4, alpha, orientation, edge-ordering, and official-ledger firewalls preserved", Passed: firewallsOK(a.Firewalls) && containsAll(a.FirewallsList(), []string{FailureNotNativeR3, FailureAlphaStillSealed, FailureHiggsOrientationStillSealed, FailureNoNativeOrientedEdgeOrdering}), Detail: FormatFirewalls(a.Firewalls)},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := []string{a.Truth, FormatRankTerm(a.Rank), FormatAlphaFlagTerm(a.AlphaFlag), FormatEdgeSupportTerm(a.Edge), FormatLeftKernelTerm(a.Kernel), FormatBMinusLTerm(a.BMinusL), FormatFunctional(a.Functional), FormatFreeze(a.Freeze), FormatFirewalls(a.Firewalls), a.Final}
		notes = append(notes, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
