package contactquotient

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactSpectralInvariantQuotientOrbitCollapseTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-SPECTRAL-INVARIANT-QUOTIENT-ORBIT-COLLAPSE"
	const name = "contact spectral-invariant quotient / orbit-collapse theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact quotient/orbit-collapse search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 119 contact action obstruction inherited", Passed: a.ActionObstructionInherited && a.ContactAction.ContactWeightedAutomorphismGroupOrder == 1 && a.ContactAction.FanoAutomorphismGroupOrder == 168 && a.ContactAction.OrderMismatchObstructionDerived && !a.ContactAction.CanonicalContactFanoAssignmentDerived, Detail: "identity-only weighted contact action vs transitive Fano symmetry"},
			{Name: "weighted spectral quotient is canonical but singleton-only", Passed: a.ContactSpectrumQuotientSearchAttempted && a.WeightedSpectrumQuotientDerived && a.WeightedQuotientCanonical && a.WeightedQuotientIsIdentity && a.WeightedQuotientPreservesAllRows && !a.WeightedQuotientProducesFanoOrbit, Detail: fmt.Sprintf("weighted orbit sizes=%v", a.WeightedQuotientOrbitSizes)},
			{Name: "anonymous quotient collapses all rows and loses spectral row data", Passed: a.AnonymousSpectrumQuotientDerived && a.AnonymousQuotientCanonical && a.AnonymousQuotientCollapsesAllRows && a.AnonymousQuotientDestroysSpectralRows && !a.AnonymousQuotientRepresentationUsable && a.OrbitCollapseObstructionDerived && a.SpectralInformationLossDerived, Detail: fmt.Sprintf("anonymous orbit sizes=%v", a.AnonymousQuotientOrbitSizes)},
			{Name: "transported Fano quotient remains convention-dependent", Passed: a.TransportedFanoQuotientPossibleAfterChoice && !a.TransportedFanoQuotientCanonical && a.Summary.CompatibleBijectionCount == 5040 && a.Summary.CanonicalTransportedQuotients == 0 && a.QuotientForkObstructionDerived, Detail: FormatSummary(a.Summary)},
			{Name: "contact quotient does not open representation or beta rows", Passed: a.ContactRows == 7 && a.OpenContactRowsAfter == 7 && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullFiniteBetaMatchingTensorDerived && !a.RepresentationRowFromQuotientDerived, Detail: FormatRows(a.Rows, 10)},
			{Name: "contact quotient search does not leak physical constants", Passed: a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, physical thetaW, threshold-corrected beta tensor, W/Z/Higgs/fermion masses, M*, or g_* are used"},
		}, Notes: []string{a.TruthStatement, "criteria: " + FormatCriteria(a.Criteria), "attempts: " + FormatAttempts(a.Attempts), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
