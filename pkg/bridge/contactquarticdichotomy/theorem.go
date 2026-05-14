package contactquarticdichotomy

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func QuarticBlockConstraintOrPropagatorDichotomyFirewallTheorem() theorem.Theorem {
	const id = "BRIDGE-QUARTIC-BLOCK-CONSTRAINT-PROPAGATOR-DICHOTOMY-FIREWALL"
	const name = "Quartic block constraint-or-propagator dichotomy / BRST-locality firewall theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build quartic block dichotomy theorem", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 156 exact local-field obstruction is inherited", Passed: a.Previous.BetaPermissionFirewallClosed && a.Previous.QuarticOrbitRows == 4 && a.Previous.QuarticBlockInvariants == 4 && a.Previous.LocalFieldRows == 0 && a.Previous.SpinStatisticsRows == 0 && a.Previous.KineticPoleResidueRows == 0 && a.Previous.QuarticBlockBetaRows == 0 && a.Previous.ContactBetaRowsAllowed == 0, Detail: fmt.Sprintf("quarticRows=%d invariants=%d local=%d spin=%d kinetic=%d beta=%d", a.Previous.QuarticOrbitRows, a.Previous.QuarticBlockInvariants, a.Previous.LocalFieldRows, a.Previous.SpinStatisticsRows, a.Previous.KineticPoleResidueRows, a.Previous.ContactBetaRowsAllowed)},
			{Name: "propagator/locality branch is audited but incomplete", Passed: a.PropagatorBranch.QuarticBlockRows == 4 && !a.PropagatorBranch.LocalFieldMap && !a.PropagatorBranch.BaseSpaceSupport && !a.PropagatorBranch.LocalSections && !a.PropagatorBranch.LorentzRepresentation && !a.PropagatorBranch.KineticOperator && !a.PropagatorBranch.PropagatorDenominator && !a.PropagatorBranch.PoleResidueTheorem && !a.PropagatorBranch.PositiveResidue && !a.PropagatorBranch.GaugeRepresentation && !a.PropagatorBranch.HyperchargeRow && !a.PropagatorBranch.MassActivation && !a.PropagatorBranch.DecouplingRule && !a.PropagatorBranch.PropagatorBranchComplete, Detail: FormatPropagatorBranch(a.PropagatorBranch)},
			{Name: "constraint/BRST branch is audited but incomplete", Passed: a.ConstraintBranch.QuarticBlockRows == 4 && !a.ConstraintBranch.ConstraintEquations && a.ConstraintBranch.ConstraintRank == 0 && !a.ConstraintBranch.GaugeRedundancy && !a.ConstraintBranch.GhostGrading && !a.ConstraintBranch.BRSTOperator && !a.ConstraintBranch.NilpotentDifferential && !a.ConstraintBranch.BRSTPairing && !a.ConstraintBranch.ExactnessOrCohomology && !a.ConstraintBranch.SupertraceCancellation && !a.ConstraintBranch.ZeroBetaLedger && !a.ConstraintBranch.ConstraintBranchComplete, Detail: FormatConstraintBranch(a.ConstraintBranch)},
			{Name: "dichotomy remains unresolved", Passed: a.Dichotomy.ExactFiniteSpectralBlock && a.Dichotomy.PropagatorBranchAudited && a.Dichotomy.ConstraintBRSTBranchAudited && a.Dichotomy.CompleteBranches == 0 && a.Dichotomy.AcceptedPhysicalBranches == 0 && a.Dichotomy.AcceptedNonphysicalBranches == 0 && !a.Dichotomy.DichotomyResolved && a.Dichotomy.BetaRowsPermitted == 0 && a.Dichotomy.ZeroRowsProved == 0, Detail: FormatDichotomy(a.Dichotomy)},
			{Name: "beta firewall requires either locality route or cancellation route", Passed: a.Firewall.ObservedInputFree && a.Firewall.QuarticBlockExact && !a.Firewall.LocalityRouteComplete && !a.Firewall.ConstraintRouteComplete && a.Firewall.RepresentationRows == 0 && a.Firewall.HyperchargeRows == 0 && a.Firewall.KineticPoleResidueRows == 0 && a.Firewall.MassActivationRows == 0 && a.Firewall.DecouplingRows == 0 && a.Firewall.ThresholdBetaRows == 0 && a.Firewall.ProvenZeroRows == 0 && !a.Firewall.PhysicalConstants && a.Firewall.FirewallClosed, Detail: FormatFirewall(a.Firewall)},
			{Name: "physical constants and threshold rows remain sealed", Passed: a.ContactRows == 7 && a.QuarticOrbitRows == 4 && a.IndividualQuarticRows == 0 && a.CanonicalQuarticBranches == 0 && a.RowwiseRootAssignmentProofs == 0 && a.ChargeSemanticRows == 0 && a.T3RRowsDerived == 0 && a.BMinusLRowsDerived == 0 && a.HyperchargeRowsDerived == 0 && a.GaugeRepresentationRows == 0 && a.SpinStatisticsRows == 0 && a.LocalFieldRows == 0 && a.KineticPoleResidueRows == 0 && a.MassActivationRows == 0 && a.DecouplingRows == 0 && a.DynkinIndexRows == 0 && a.BRSTCancellationRows == 0 && a.ConstraintRows == 0 && a.PropagatorRows == 0 && a.RepresentationCompleteRows == 0 && a.QuarticBlockBetaRows == 0 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && a.BetaPermissionFirewallClosed && a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}}
	}}
}
