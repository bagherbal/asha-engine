package contactrootiso

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ExactContactRootIsolationRowwiseEigenprojectorAssignmentTheorem() theorem.Theorem {
	const id = "BRIDGE-EXACT-CONTACT-ROOT-ISOLATION-ROWWISE-EIGENPROJECTOR-ASSIGNMENT"
	const name = "Exact contact root-isolation / row-wise eigenprojector assignment theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build exact contact root-isolation certificate", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 149 exact matrix and charpoly certificates are inherited", Passed: a.Previous.ExactRationalOverlapMatrix && a.Previous.ExactCharacteristicCertified && a.Previous.ExactAnnihilationCertified && a.Previous.RootIsolationCertificates == 0 && a.Previous.ContactBetaRowsAllowed == 0, Detail: fmt.Sprintf("matrix=%t char=%t annihilation=%t prevRootIso=%d beta=%d", a.Previous.ExactRationalOverlapMatrix, a.Previous.ExactCharacteristicCertified, a.Previous.ExactAnnihilationCertified, a.Previous.RootIsolationCertificates, a.Previous.ContactBetaRowsAllowed)},
			{Name: "quartic roots are isolated by exact rational sign-change intervals", Passed: a.Quartic.IntervalsDisjoint && a.Quartic.SignChanges == 4 && a.Quartic.Degree == 4 && a.Quartic.AllQuarticRootsCovered, Detail: FormatQuartic(a.Quartic)},
			{Name: "all seven non-unit contact spectral roots are isolated", Passed: a.ExactRootIsolationCertified && a.Certificate.ExactCharpolyInherited && a.Certificate.RationalRootsCertified == 3 && a.Certificate.QuarticRootsCertified == 4 && a.Certificate.TotalPartialRoots == 7 && a.Certificate.UnitRootMultiplicity == 7 && a.RootIsolationCertificates == 7, Detail: FormatCertificate(a.Certificate)},
			{Name: "root isolation still does not construct number-field projectors or row assignments", Passed: a.Projectors.ExactRootIsolationAvailable && a.Projectors.ExactNumberFieldProjectors == 0 && a.Projectors.RationalProjectors == 0 && a.Projectors.QuarticProjectors == 0 && a.Projectors.RowwiseProjectorAssignments == 0 && a.Projectors.RootToContactRowSemantics == 0 && a.Projectors.ChargeSemanticRows == 0 && a.Projectors.RepresentationRows == 0 && a.Projectors.BetaRowsAllowed == 0, Detail: FormatProjectors(a.Projectors)},
			{Name: "physics requirements remain incomplete", Passed: a.Requirements.ExactRationalMatrix && a.Requirements.ExactCharacteristicPolynomial && a.Requirements.ExactRootIsolation && !a.Requirements.ExactNumberFieldArithmetic && !a.Requirements.EigenprojectorFormulaInField && !a.Requirements.RowwiseProjectorAssignment && !a.Requirements.ChargeOperatorSelected && !a.Requirements.RepresentationRowsSelected && a.Requirements.ObservedInputFree && !a.Requirements.AllSatisfiedForPhysics, Detail: FormatRequirements(a.Requirements)},
			{Name: "contact beta and physical constants remain sealed", Passed: a.ContactRows == 7 && a.ChargeSemanticRows == 0 && a.T3RRowsDerived == 0 && a.BMinusLRowsDerived == 0 && a.HyperchargeRowsDerived == 0 && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && a.BetaPermissionFirewallClosed && a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}}
	}}
}
