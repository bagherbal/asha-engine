package contactchargelattice

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactChargeLatticeEmbeddingRationalSpectrumObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-CHARGE-LATTICE-EMBEDDING-RATIONAL-SPECTRUM-OBSTRUCTION"
	const name = "Contact charge lattice embedding / rational-spectrum obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact charge lattice obstruction", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 145 centered current is inherited as diagnostic only", Passed: a.ContactRows == 7 && a.CenteredPositiveRows == 3 && a.CenteredNegativeRows == 4 && a.CenteredZeroRows == 0 && a.BetaPermissionFirewallClosed && a.ChargeSemanticEmbeddings == 0, Detail: fmt.Sprintf("centered=[%s]", FormatApproximations(a.BoundedRationalAudit.Approximations))},
			{Name: "half-integer and sixth-integer charge lattices do not contain the raw centered spectrum", Passed: a.HalfIntegerAudit.AppliesToRawSpectrum && !a.HalfIntegerAudit.ExactEmbedding && a.HalfIntegerAudit.ExactRows < a.ContactRows && a.SixthIntegerAudit.AppliesToRawSpectrum && !a.SixthIntegerAudit.ExactEmbedding && a.SixthIntegerAudit.ExactRows < a.ContactRows, Detail: FormatCandidates([]LatticeCandidate{a.HalfIntegerAudit, a.SixthIntegerAudit})},
			{Name: "balanced seventh lattice exists only for the summary split, not for charge semantics", Passed: a.SeventhBalancedAudit.AppliesToBalancedSummary && a.SeventhBalancedAudit.ExactEmbedding && !a.SeventhBalancedAudit.AppliesToRawSpectrum && !a.SeventhBalancedAudit.ChargeOperatorSemantic && a.BalancedExactEmbeddings == 1, Detail: FormatCandidate(a.SeventhBalancedAudit)},
			{Name: "bounded rational approximation does not provide an exact selected raw charge lattice", Passed: a.BoundedRationalAudit.AppliesToRawSpectrum && a.BoundedRationalAudit.RequiresDenominatorFit && !a.BoundedRationalAudit.ExactEmbedding && a.BoundedRationalAudit.ExactRows < a.ContactRows && !a.BoundedRationalAudit.ChargeOperatorSemantic, Detail: FormatCandidate(a.BoundedRationalAudit) + " approximations=" + FormatApproximations(a.BoundedRationalAudit.Approximations)},
			{Name: "free scaling and observed-charge fitting are blocked", Passed: a.FreeScaledAudit.RequiresScaleChoice && a.FreeScaledAudit.RequiresDenominatorFit && a.FreeScaledAudit.ApproximateEmbedding && !a.FreeScaledAudit.ExactEmbedding && !a.FreeScaledAudit.ChargeOperatorSemantic && a.ObservedFitAudit.RequiresObservedInput && !a.ObservedFitAudit.Available && !a.HiddenObservedInputUsed, Detail: FormatCandidates([]LatticeCandidate{a.FreeScaledAudit, a.ObservedFitAudit})},
			{Name: "charge-lattice requirements remain unsatisfied and beta firewall stays closed", Passed: !a.Requirements.FiniteSelectedLattice && !a.Requirements.RawSpectrumEmbedded && !a.Requirements.PhysicalChargeSemantics && !a.Requirements.OperatorPullback && !a.Requirements.LocalFieldMap && !a.Requirements.MassActivation && !a.Requirements.DecouplingRule && a.Requirements.ObservedInputFree && !a.Requirements.AllSatisfied && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0, Detail: FormatRequirements(a.Requirements)},
			{Name: "no physical constants or threshold beta rows leak through rational-spectrum search", Passed: a.LatticeCandidatesAudited == 6 && a.AvailableCandidates == 5 && a.RawExactEmbeddings == 0 && a.RawApproxEmbeddings == 1 && a.ChargeSemanticEmbeddings == 0 && a.ScaleDependentCandidates == 2 && a.ObservedFitCandidates == 1 && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}}
	}}
}
