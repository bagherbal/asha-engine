package contactidempotent

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ExactContactEigenprojectorNumberFieldSpectralIdempotentConstructionAttemptTheorem() theorem.Theorem {
	const id = "BRIDGE-EXACT-CONTACT-EIGENPROJECTOR-NUMBER-FIELD-SPECTRAL-IDEMPOTENT-ATTEMPT"
	const name = "Exact contact eigenprojector number-field / spectral idempotent construction attempt"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build exact contact spectral-idempotent attempt", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 150 exact matrix, charpoly, and root isolation are inherited", Passed: a.Previous.ExactRationalOverlapMatrix && a.Previous.ExactCharacteristicCertified && a.Previous.ExactRootIsolationCertified && a.Previous.RootIsolationCertificates == 7 && a.Previous.ContactBetaRowsAllowed == 0, Detail: fmt.Sprintf("matrix=%t char=%t rootIso=%t roots=%d beta=%d", a.Previous.ExactRationalOverlapMatrix, a.Previous.ExactCharacteristicCertified, a.Previous.ExactRootIsolationCertified, a.Previous.RootIsolationCertificates, a.Previous.ContactBetaRowsAllowed)},
			{Name: "exact rational primary spectral idempotent blocks are constructed", Passed: a.Decomposition.ExactMatrixInherited && a.Decomposition.ExactCharpolyInherited && a.Decomposition.ExactRootIsolation && a.Decomposition.FactorsPairwiseCoprime && a.Decomposition.SquarefreeMinimalFactors && a.Decomposition.RationalBlockProjectors == 5 && a.Decomposition.RationalSimpleProjectors == 4 && a.Decomposition.QuarticPrimaryProjectors == 1 && a.Decomposition.IndividualQuarticProjectors == 0 && a.Decomposition.TotalSpectralDimension == 14, Detail: FormatDecomposition(a.Decomposition)},
			{Name: "quartic individual eigenprojectors still require noncanonical number-field branch data", Passed: a.NumberField.RootIsolationAvailable && a.NumberField.CandidateNumberFieldDegree == 4 && !a.NumberField.RootChosenCanonically && !a.NumberField.FieldEmbeddingSelected && !a.NumberField.ExactAlgebraicRootSymbol && !a.NumberField.ExactEigenprojectorFormula && !a.NumberField.GaloisBranchChoiceFree && a.NumberField.IndividualQuarticProjectors == 0, Detail: FormatNumberField(a.NumberField)},
			{Name: "spectral idempotents do not assign contact row semantics", Passed: a.RowAudit.SpectralBlockIdempotents == 5 && a.RowAudit.EigenvalueRowsSeparated == 4 && a.RowAudit.ContactRows == 7 && a.RowAudit.RowwiseEigenprojectorAssignments == 0 && a.RowAudit.ContactRootToModeMap == 0 && a.RowAudit.ChargeSemanticRows == 0 && a.RowAudit.T3RRows == 0 && a.RowAudit.BMinusLRows == 0 && a.RowAudit.HyperchargeRows == 0 && a.RowAudit.RepresentationRows == 0 && a.RowAudit.BetaRowsAllowed == 0, Detail: FormatRowAudit(a.RowAudit)},
			{Name: "physics requirements remain incomplete", Passed: a.Requirements.ExactRationalMatrix && a.Requirements.ExactCharacteristicPolynomial && a.Requirements.ExactRootIsolation && a.Requirements.RationalPrimaryIdempotents && !a.Requirements.ExactNumberFieldArithmetic && !a.Requirements.CanonicalQuarticRootChoice && !a.Requirements.IndividualQuarticEigenprojectors && !a.Requirements.RowwiseContactAssignment && !a.Requirements.ChargeOperatorSelected && !a.Requirements.RepresentationRowsSelected && a.Requirements.ObservedInputFree && !a.Requirements.AllSatisfiedForPhysics, Detail: FormatRequirements(a.Requirements)},
			{Name: "contact beta and physical constants remain sealed", Passed: a.ContactRows == 7 && a.RationalPrimaryIdempotents == 5 && a.ExactNumberFieldProjectors == 0 && a.IndividualQuarticProjectors == 0 && a.RowwiseRootAssignmentProofs == 0 && a.ChargeSemanticRows == 0 && a.T3RRowsDerived == 0 && a.BMinusLRowsDerived == 0 && a.HyperchargeRowsDerived == 0 && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && a.BetaPermissionFirewallClosed && a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}}
	}}
}
