package contactcohomology

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteContactConstraintDifferentialCohomologyObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-CONTACT-CONSTRAINT-DIFFERENTIAL-COHOMOLOGY-OBSTRUCTION"
	const name = "finite contact constraint differential / cohomology obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact cohomology obstruction", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 113 unresolved contact carrier inherited", Passed: a.BranchSelector.BranchSelectorAttempted && !a.BranchSelector.BranchSelectorDerived && a.OpenContactRowsBefore == 7 && a.BranchSelector.BetaCorrectionRowsAllowed == 0, Detail: fmt.Sprintf("unresolved before=%d; beta rows=%d", a.OpenContactRowsBefore, a.BranchSelector.BetaCorrectionRowsAllowed)},
			{Name: "seven positive contact modes form only a candidate chain carrier", Passed: a.ChainGroupCarrierConstructed && a.ContactRows == 7 && a.PositiveFiniteContactRows == 7 && a.CandidateChainGroupDimension == 7, Detail: FormatRows(a.Rows, 10)},
			{Name: "zero differential audited but not a cancellation proof", Passed: a.ZeroDifferentialConstructed && a.ZeroDifferentialSquareZero && a.ZeroDifferentialCohomologyDimension == 7 && !a.ZeroDifferentialProvesCancellation, Detail: FormatDifferentials(a.DifferentialAttempts)},
			{Name: "nontrivial canonical nilpotent differential absent", Passed: a.AnyNontrivialCandidateConstructed && !a.AnyNontrivialCandidateCanonical && !a.AnyNontrivialCandidateSquareZero && !a.CanonicalDifferentialDerived && !a.NontrivialNilpotentDifferentialDerived, Detail: FormatDifferentials(a.DifferentialAttempts)},
			{Name: "BRST grading/pairing/exactness absent", Passed: !a.GhostGradingDerived && !a.PairingDerived && !a.ExactnessOrCohomologyDerived && !a.AcyclicComplexDerived && !a.CancellationLedgerDerived, Detail: FormatCriteria(a.ObstructionCriteria)},
			{Name: "constraint complex completes no contact rows", Passed: a.ConstraintComplexCompleteRows == 0 && a.ContactZeroRowsProved == 0 && a.ContactBetaRowsAllowed == 0 && a.OpenContactRowsAfter == 7, Detail: fmt.Sprintf("complete=%d; zero rows=%d; beta rows=%d; open after=%d", a.ConstraintComplexCompleteRows, a.ContactZeroRowsProved, a.ContactBetaRowsAllowed, a.OpenContactRowsAfter)},
			{Name: "cohomology obstruction recorded", Passed: a.CohomologyObstructionDerived && a.NoCanonicalBRSTDifferentialUnderCurrentData && !a.RepresentationOrConstraintDichotomyDerived && !a.BranchSelectorDerived, Detail: "zero differential leaves H dimension 7; nontrivial alternatives are noncanonical under current data"},
			{Name: "threshold beta tensor remains sealed", Passed: a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullFiniteBetaMatchingTensorDerived, Detail: "no contact Δb_i row and no contact zero-row cancellation is proven"},
			{Name: "residual physical-flow nullity unchanged", Passed: a.ResidualNullityAfter == a.ResidualNullityBefore && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken, Detail: fmt.Sprintf("nullity before=%d; after=%d", a.ResidualNullityBefore, a.ResidualNullityAfter)},
			{Name: "physical predictions remain sealed", Passed: !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, physical thetaW, W/Z/Higgs/fermion masses, M*, g_*, or observed thresholds are used"},
		}, Notes: []string{a.TruthStatement, "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
