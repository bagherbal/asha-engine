package fockcontactkernel

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FockContactKernelSelectionOperatorIntertwiningObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-FOCK-CONTACT-KERNEL-SELECTION-OPERATOR-INTERTWINING-OBSTRUCTION"
	const name = "Fock-contact kernel selection / operator-intertwining obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Fock-contact kernel selection search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "rank-seven Fock quotient needs a nine-dimensional kernel", Passed: a.MatterDimension == 16 && a.ContactRows == 7 && a.RequiredKernelDim == 9 && a.UnconstrainedGrassmannDimension == 63, Detail: FormatSummary(a.Summary)},
			{Name: "T3R-invariant kernels are still a family", Passed: a.Summary.T3RInvariantSplitPatterns == 8 && a.Summary.T3RInvariantResidualDimMin == 7 && a.Summary.T3RInvariantResidualDimMax == 31, Detail: FormatSplits(a.Splits)},
			{Name: "T3R/chirality joint-invariant kernels are still noncanonical", Passed: a.Summary.T3RChiralityInvariantSplitPatterns == 80 && a.Summary.T3RChiralityResidualDimMin == 3 && a.Summary.T3RChiralityResidualDimMax == 15 && a.CanonicalKernelCandidates == 1, Detail: FormatCandidates(a.Candidates)},
			{Name: "no full operator-intertwiner or target contact operator is derived", Passed: a.FullOperatorIntertwinersDerived == 0 && a.TargetContactOperatorsDerived == 0 && a.T3RPullbackRowsDerived == 0 && a.ChiralityPullbackRowsDerived == 0 && a.BMinusLPullbackRowsDerived == 0 && a.SU2LPullbackRowsDerived == 0 && a.HyperchargeRowsDerived == 0, Detail: fmt.Sprintf("intertwiners=%d targetOps=%d Y=%d", a.FullOperatorIntertwinersDerived, a.TargetContactOperatorsDerived, a.HyperchargeRowsDerived)},
			{Name: "contact representation and beta firewall remain closed", Passed: a.BetaPermissionFirewallClosed && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullBetaMatchingTensorDerived, Detail: fmt.Sprintf("rep=%d open=%d beta=%d zero=%d", a.RepresentationCompleteRows, a.RepresentationOpenRows, a.ContactBetaRowsAllowed, a.ContactZeroRowsProved)},
			{Name: "S6 ambiguity, physical-flow nullity, and no-observed-input discipline are preserved", Passed: a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: fmt.Sprintf("s6=%d nullity=%d→%d", a.ResidualS6Choices, a.ResidualNullityBefore, a.ResidualNullityAfter)},
		}, Notes: []string{a.TruthStatement, "splits: " + FormatSplits(a.Splits), "candidates: " + FormatCandidates(a.Candidates), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
