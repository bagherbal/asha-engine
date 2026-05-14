package contactreconstruction

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactSpectralReconstructionInvariantToRowLiftingObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-SPECTRAL-RECONSTRUCTION-INVARIANT-ROW-LIFTING-OBSTRUCTION"
	const name = "contact spectral reconstruction / invariant-to-row lifting obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact spectral reconstruction search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 120 quotient fork inherited", Passed: a.QuotientForkObstructionInherited && a.OrbitCollapseObstructionInherited && a.SpectralInformationLossInherited && a.Quotient.Summary.CompatibleBijectionCount == 5040 && !a.Quotient.TransportedFanoQuotientCanonical, Detail: "weighted quotient preserves rows; anonymous quotient restores symmetry only by losing row data"},
			{Name: "weighted singleton lift is canonical but not Fano-like", Passed: a.InvariantToRowLiftingSearchAttempted && a.WeightedSingletonLiftDerived && a.WeightedSingletonLiftCanonical && a.WeightedSingletonLiftPreservesRows && !a.WeightedSingletonLiftFanoLike && !a.WeightedSingletonLiftRepUsable, Detail: fmt.Sprintf("weighted orbit sizes=%v", a.Summary.WeightedOrbitSizes)},
			{Name: "anonymous invariant lift has 7! noncanonical row reconstructions", Passed: a.AnonymousInvariantLiftAttempted && a.AnonymousInvariantLiftConstructed && !a.AnonymousInvariantLiftCanonical && a.AnonymousInvariantLiftNeedsChoice && a.AnonymousInvariantLiftPossibleRows == 5040 && a.AnonymousInvariantLiftCanonicalRows == 0 && a.RowLiftingAmbiguityDerived, Detail: FormatSummary(a.Summary)},
			{Name: "spectral multiset recovers values but not row semantics", Passed: a.SpectralMultisetRecovered && a.SpectralMultisetRecoversValues && !a.SpectralMultisetRecoversRows && !a.SpectralMultisetRecoversFanoRows && !a.InvariantToRowReconstructionDerived && a.ReconstructionObstructionDerived && a.InformationChoiceNoGoDerived && !a.NoLossNoChoiceLiftExists, Detail: FormatAttempts(a.Attempts)},
			{Name: "contact reconstruction does not open representation or beta rows", Passed: a.ContactRows == 7 && a.OpenContactRowsAfter == 7 && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullFiniteBetaMatchingTensorDerived && !a.FanoEquivariantRowLiftDerived, Detail: FormatRows(a.Rows, 10)},
			{Name: "contact reconstruction search does not leak physical constants", Passed: a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, physical thetaW, threshold-corrected beta tensor, W/Z/Higgs/fermion masses, M*, or g_* are used"},
		}, Notes: []string{a.TruthStatement, "criteria: " + FormatCriteria(a.Criteria), "attempts: " + FormatAttempts(a.Attempts), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
