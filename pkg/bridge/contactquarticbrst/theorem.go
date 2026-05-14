package contactquarticbrst

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func QuarticBRSTCandidateDifferentialZeroSupertraceAttemptTheorem() theorem.Theorem {
	const id = "BRIDGE-QUARTIC-BRST-CANDIDATE-DIFFERENTIAL-ZERO-SUPERTRACE-ATTEMPT"
	const name = "Quartic BRST candidate differential / zero-supertrace construction attempt"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build quartic BRST candidate theorem", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 157 dichotomy firewall is inherited", Passed: a.Previous.BetaPermissionFirewallClosed && a.Previous.QuarticOrbitRows == 4 && a.Previous.BRSTCancellationRows == 0 && a.Previous.ConstraintRows == 0 && a.Previous.QuarticBlockBetaRows == 0 && a.Previous.ContactBetaRowsAllowed == 0 && a.Previous.ContactZeroRowsProved == 0, Detail: fmt.Sprintf("quarticRows=%d brst=%d constraint=%d beta=%d contactBeta=%d", a.Previous.QuarticOrbitRows, a.Previous.BRSTCancellationRows, a.Previous.ConstraintRows, a.Previous.QuarticBlockBetaRows, a.Previous.ContactBetaRowsAllowed)},
			{Name: "zero differential is canonical but inert", Passed: len(a.DifferentialCandidates) >= 1 && a.DifferentialCandidates[0].Name == "zero differential Q=0" && a.DifferentialCandidates[0].Nilpotent && a.DifferentialCandidates[0].SquareZero && a.DifferentialCandidates[0].Canonical && a.DifferentialCandidates[0].GaloisInvariant && a.DifferentialCandidates[0].CohomologyDimension == 4 && !a.DifferentialCandidates[0].ZeroSupertraceLedger && !a.DifferentialCandidates[0].ZeroBetaLedger && !a.DifferentialCandidates[0].BRSTCancellationComplete, Detail: FormatDifferentialCandidate(a.DifferentialCandidates[0])},
			{Name: "nonzero BRST candidates are not canonical zero-beta ledgers", Passed: a.Construction.DifferentialsAudited == 4 && a.Construction.NilpotentCandidates == 3 && a.Construction.CanonicalNilpotent == 1 && a.Construction.NonzeroCanonicalBRST == 0 && a.Construction.CompleteBRSTCancellations == 0 && a.Construction.ZeroBetaLedgers == 0 && !a.Construction.ConstructionComplete, Detail: FormatConstruction(a.Construction)},
			{Name: "Galois-invariant gradings do not yield zero supertrace", Passed: a.Supertrace.QuarticBlockRows == 4 && a.Supertrace.GaloisOrbitRows == 4 && a.Supertrace.GaloisInvariantGradings == 2 && a.Supertrace.NontrivialInvariantGradings == 0 && a.Supertrace.ZeroSupertraceGradings == 0 && !a.Supertrace.CanonicalZeroSupertrace && a.Supertrace.PointwiseCancellationPairs == 0 && !a.Supertrace.ZeroBetaLedger, Detail: FormatSupertrace(a.Supertrace)},
			{Name: "two-even/two-odd grading is branch-dependent", Passed: len(a.GhostGradings) == 3 && a.GhostGradings[2].Name == "two-even/two-odd split" && a.GhostGradings[2].Nontrivial && a.GhostGradings[2].TraceZeroPossible && !a.GhostGradings[2].GaloisInvariant && a.GhostGradings[2].RequiresBranchChoices && !a.GhostGradings[2].ZeroSupertraceLedger && !a.GhostGradings[2].BRSTCancellationComplete, Detail: FormatGhostGrading(a.GhostGradings[2])},
			{Name: "BRST cancellation firewall remains closed", Passed: a.Firewall.ObservedInputFree && a.Firewall.QuarticBlockExact && a.Firewall.ConstraintRouteAudited && !a.Firewall.BRSTOperator && !a.Firewall.GhostGrading && !a.Firewall.NilpotentDifferential && !a.Firewall.ExactnessOrCohomology && !a.Firewall.SupertraceCancellation && !a.Firewall.ZeroBetaLedger && a.Firewall.RepresentationRows == 0 && a.Firewall.HyperchargeRows == 0 && a.Firewall.ThresholdBetaRows == 0 && a.Firewall.ProvenZeroRows == 0 && !a.Firewall.PhysicalConstants && a.Firewall.FirewallClosed, Detail: FormatFirewall(a.Firewall)},
			{Name: "physical constants and beta rows remain sealed", Passed: a.ContactRows == 7 && a.QuarticOrbitRows == 4 && a.IndividualQuarticRows == 0 && a.CanonicalQuarticBranches == 0 && a.RowwiseRootAssignmentProofs == 0 && a.ChargeSemanticRows == 0 && a.T3RRowsDerived == 0 && a.BMinusLRowsDerived == 0 && a.HyperchargeRowsDerived == 0 && a.GaugeRepresentationRows == 0 && a.SpinStatisticsRows == 0 && a.LocalFieldRows == 0 && a.KineticPoleResidueRows == 0 && a.MassActivationRows == 0 && a.DecouplingRows == 0 && a.DynkinIndexRows == 0 && a.BRSTCancellationRows == 0 && a.ConstraintRows == 0 && a.PropagatorRows == 0 && a.RepresentationCompleteRows == 0 && a.QuarticZeroBetaRows == 0 && a.QuarticBlockBetaRows == 0 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && a.BetaPermissionFirewallClosed && a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}}
	}}
}
