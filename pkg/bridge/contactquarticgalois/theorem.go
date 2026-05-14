package contactquarticgalois

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func QuarticContactNumberFieldBranchGaloisSymmetryObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-QUARTIC-CONTACT-NUMBER-FIELD-BRANCH-GALOIS-OBSTRUCTION"
	const name = "Quartic contact number-field branch / Galois symmetry obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build quartic contact Galois obstruction", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 151 Q-primary idempotent certificate is inherited", Passed: a.Previous.ExactRationalOverlapMatrix && a.Previous.ExactCharacteristicCertified && a.Previous.ExactRootIsolationCertified && a.Previous.RationalPrimaryIdempotents == 5 && a.Previous.IndividualQuarticProjectors == 0 && a.Previous.ContactBetaRowsAllowed == 0, Detail: fmt.Sprintf("matrix=%t char=%t rootIso=%t Qidempotents=%d individualQuartic=%d beta=%d", a.Previous.ExactRationalOverlapMatrix, a.Previous.ExactCharacteristicCertified, a.Previous.ExactRootIsolationCertified, a.Previous.RationalPrimaryIdempotents, a.Previous.IndividualQuarticProjectors, a.Previous.ContactBetaRowsAllowed)},
			{Name: "quartic block is exact but branch-rich", Passed: a.QuarticField.Degree == 4 && a.QuarticField.RationalRootCount == 0 && a.QuarticField.IrreducibleOverQ && a.QuarticField.RealRootsIsolated == 4 && a.QuarticField.RootIntervalsInherited && !a.QuarticField.DiscriminantSquare && a.QuarticField.GaloisOrderCandidate == 24 && a.QuarticField.GaloisTransitive && a.QuarticField.Branches == 4 && !a.QuarticField.CanonicalBranchSelector, Detail: FormatQuarticField(a.QuarticField)},
			{Name: "Galois-invariant block does not give individual quartic projectors", Passed: a.BranchAudit.QuarticRoots == 4 && a.BranchAudit.IndividualNumberFieldBranches == 4 && a.BranchAudit.BranchChoicesRequired == 4 && a.BranchAudit.CanonicalRootChoices == 0 && a.BranchAudit.FieldEmbeddingsSelected == 0 && a.BranchAudit.ExactIndividualProjectors == 0 && a.BranchAudit.GaloisInvariantQuarticBlock && !a.BranchAudit.GaloisInvariantIndividualRoots, Detail: FormatBranchAudit(a.BranchAudit)},
			{Name: "projector semantics stop at rational primary decomposition", Passed: a.ProjectorAudit.RationalPrimaryBlocks == 5 && a.ProjectorAudit.RationalSimpleProjectors == 4 && a.ProjectorAudit.QuarticPrimaryBlocks == 1 && a.ProjectorAudit.IndividualQuarticProjectors == 0 && a.ProjectorAudit.RowwiseRootAssignments == 0 && a.ProjectorAudit.ContactRootToModeMap == 0 && a.ProjectorAudit.ChargeSemanticRows == 0 && a.ProjectorAudit.T3RRows == 0 && a.ProjectorAudit.BMinusLRows == 0 && a.ProjectorAudit.HyperchargeRows == 0 && a.ProjectorAudit.RepresentationRows == 0 && a.ProjectorAudit.ContactBetaRowsAllowed == 0, Detail: FormatProjectorAudit(a.ProjectorAudit)},
			{Name: "physics firewall remains closed", Passed: a.Firewall.ObservedInputFree && a.Firewall.ExactMatrix && a.Firewall.ExactCharpoly && a.Firewall.ExactRootIsolation && a.Firewall.RationalPrimaryIdempotents && !a.Firewall.QuarticBranchSelector && !a.Firewall.IndividualQuarticSplit && !a.Firewall.RowSemanticMap && !a.Firewall.ContactCharges && !a.Firewall.RepresentationRows && !a.Firewall.BetaRows && !a.Firewall.PhysicalConstants && !a.Firewall.AllSatisfiedForPhysics, Detail: FormatFirewall(a.Firewall)},
			{Name: "contact beta and physical constants remain sealed", Passed: a.ContactRows == 7 && a.QuarticNumberFieldDegree == 4 && a.QuarticGaloisOrderCandidate == 24 && a.QuarticBranches == 4 && a.CanonicalQuarticBranches == 0 && a.ExactNumberFieldProjectors == 0 && a.IndividualQuarticProjectors == 0 && a.RowwiseRootAssignmentProofs == 0 && a.ChargeSemanticRows == 0 && a.T3RRowsDerived == 0 && a.BMinusLRowsDerived == 0 && a.HyperchargeRowsDerived == 0 && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && a.BetaPermissionFirewallClosed && a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}}
	}}
}
