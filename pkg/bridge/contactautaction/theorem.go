package contactautaction

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactSideAutomorphismActionEquivariantAssignmentSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-SIDE-AUTOMORPHISM-ACTION-EQUIVARIANT-ASSIGNMENT-SEARCH"
	const name = "contact-side automorphism action construction / equivariant assignment search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact-side automorphism action search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 118 symmetry selector obstruction inherited", Passed: a.SymmetrySelectorObstructionInherited && a.FanoAutomorphismGroupDerived && a.FanoAutomorphismGroupOrder == 168 && a.FanoPointActionTransitive && a.FanoLineActionTransitive, Detail: "the transitive 168-element Fano action remains the target symmetry"},
			{Name: "contact weighted automorphism group is identity-only", Passed: a.ContactSideActionSearchAttempted && a.ContactWeightedAutomorphismGroupDerived && a.ContactWeightedAutomorphismGroupOrder == 1 && a.ContactWeightedActionIdentityOnly && a.ContactSpectralValuesAllDistinct, Detail: FormatSummary(a.Summary)},
			{Name: "Aut(Fano) cannot act canonically while preserving contact overlap data", Passed: a.OrderMismatchObstructionDerived && !a.AutFanoActionOnContactDerived && !a.AutFanoActionPreservingContactData && a.TrivialContactActionRejected, Detail: fmt.Sprintf("weighted contact order=%d vs Fano order=%d", a.ContactWeightedAutomorphismGroupOrder, a.FanoAutomorphismGroupOrder)},
			{Name: "transported actions exist only after choosing one of 7! bijections", Passed: a.TransportedFanoActionsPossibleAfterChoice && !a.TransportedFanoActionCanonical && !a.EquivariantAssignmentDerived && !a.CanonicalContactFanoAssignmentDerived && !a.NaturalitySquareFormulable, Detail: fmt.Sprintf("compatible bijections=%d; canonical transported actions=%d", a.Summary.CompatibleBijectionCount, a.Summary.CanonicalTransportedActions)},
			{Name: "contact representation and beta permission remain closed", Passed: a.ContactRows == 7 && a.OpenContactRowsAfter == 7 && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullFiniteBetaMatchingTensorDerived, Detail: FormatRows(a.Rows, 10)},
			{Name: "contact action search does not leak physical constants", Passed: a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ResidualSymmetryBroken && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived && !a.HiddenObservedInputUsed, Detail: "no alpha, physical thetaW, W/Z/Higgs/fermion masses, M*, g_*, or observed thresholds are used"},
		}, Notes: []string{a.TruthStatement, "criteria: " + FormatCriteria(a.Criteria), "attempts: " + FormatAttempts(a.Attempts), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
