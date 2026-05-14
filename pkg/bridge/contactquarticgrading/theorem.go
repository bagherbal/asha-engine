package contactquarticgrading

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func QuarticGhostGradingGaloisInvarianceNontrivialParityObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-QUARTIC-GHOST-GRADING-GALOIS-INVARIANCE-PARITY-OBSTRUCTION"
	const name = "Quartic ghost-grading Galois invariance / nontrivial parity obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build quartic ghost-grading theorem", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 158 BRST obstruction is inherited", Passed: a.Previous.BetaPermissionFirewallClosed && a.Previous.QuarticOrbitRows == 4 && a.Previous.QuarticZeroBetaRows == 0 && a.Previous.QuarticBlockBetaRows == 0 && a.Previous.ContactBetaRowsAllowed == 0 && a.Previous.ContactZeroRowsProved == 0, Detail: fmt.Sprintf("quarticRows=%d quarticZero=%d quarticBeta=%d contactBeta=%d zero=%d", a.Previous.QuarticOrbitRows, a.Previous.QuarticZeroBetaRows, a.Previous.QuarticBlockBetaRows, a.Previous.ContactBetaRowsAllowed, a.Previous.ContactZeroRowsProved)},
			{Name: "quartic Galois orbit admits only constant invariant parity functions", Passed: a.GaloisAction.QuarticRows == 4 && a.GaloisAction.GaloisOrbitRows == 4 && a.GaloisAction.GaloisOrderCandidate == 24 && a.GaloisAction.TransitiveOrbit && a.GaloisAction.InvariantParityFunctions == 2 && a.GaloisAction.NontrivialInvariantParity == 0, Detail: FormatGaloisAction(a.GaloisAction)},
			{Name: "all sixteen parity assignments are audited", Passed: a.GaloisAction.AllParityAssignments == 16 && a.GaloisAction.NontrivialParityAssignments == 14 && len(a.ParityClasses) == 4 && a.Obstruction.AssignmentsAudited == 16 && a.Obstruction.NontrivialAssignments == 14, Detail: FormatObstruction(a.Obstruction)},
			{Name: "zero-count two/two gradings are branch-dependent", Passed: a.GaloisAction.ZeroSignedCountAssignments == 6 && a.GaloisAction.InvariantZeroSignedCount == 0 && len(a.ParityClasses) == 4 && a.ParityClasses[3].Name == "two-even/two-odd parity split" && a.ParityClasses[3].ZeroSignedCount && !a.ParityClasses[3].GaloisInvariant && a.ParityClasses[3].RequiresBranchChoice && a.ParityClasses[3].Assignments == 6 && a.ParityClasses[3].OrbitSize == 6 && !a.ParityClasses[3].SupertraceLedger && !a.ParityClasses[3].ZeroBetaLedger && !a.ParityClasses[3].CancellationComplete, Detail: FormatParityClass(a.ParityClasses[3])},
			{Name: "uniform gradings are invariant but non-cancelling", Passed: a.ParityClasses[0].GaloisInvariant && !a.ParityClasses[0].Nontrivial && !a.ParityClasses[0].ZeroSignedCount && a.ParityClasses[0].StabilizerOrder == 24 && a.ParityClasses[1].GaloisInvariant && !a.ParityClasses[1].Nontrivial && !a.ParityClasses[1].ZeroSignedCount && a.ParityClasses[1].StabilizerOrder == 24, Detail: FormatParityClass(a.ParityClasses[0]) + " :: " + FormatParityClass(a.ParityClasses[1])},
			{Name: "no canonical zero-supertrace or zero-beta grading is derived", Passed: !a.Obstruction.CanonicalZeroSupertrace && a.Obstruction.BranchChoiceRequired && !a.Obstruction.CompleteGhostGrading && !a.Obstruction.ZeroBetaLedger && a.Summary.CanonicalZeroSupertrace == false && a.Summary.QuarticZeroBetaRows == 0 && a.Summary.QuarticBlockBetaRows == 0, Detail: FormatObstruction(a.Obstruction)},
			{Name: "quartic grading firewall remains closed", Passed: a.Firewall.ObservedInputFree && a.Firewall.QuarticBlockExact && a.Firewall.BRSTRouteAudited && a.Firewall.GaloisActionAudited && !a.Firewall.GhostGrading && !a.Firewall.NontrivialParity && !a.Firewall.SupertraceCancellation && !a.Firewall.ZeroBetaLedger && a.Firewall.RepresentationRows == 0 && a.Firewall.HyperchargeRows == 0 && a.Firewall.ThresholdBetaRows == 0 && a.Firewall.ProvenZeroRows == 0 && !a.Firewall.PhysicalConstants && a.Firewall.FirewallClosed, Detail: FormatFirewall(a.Firewall)},
			{Name: "physical constants and contact beta rows remain sealed", Passed: a.ContactRows == 7 && a.QuarticOrbitRows == 4 && a.IndividualQuarticRows == 0 && a.CanonicalQuarticBranches == 0 && a.RowwiseRootAssignmentProofs == 0 && a.ChargeSemanticRows == 0 && a.T3RRowsDerived == 0 && a.BMinusLRowsDerived == 0 && a.HyperchargeRowsDerived == 0 && a.GaugeRepresentationRows == 0 && a.SpinStatisticsRows == 0 && a.LocalFieldRows == 0 && a.KineticPoleResidueRows == 0 && a.MassActivationRows == 0 && a.DecouplingRows == 0 && a.DynkinIndexRows == 0 && a.BRSTCancellationRows == 0 && a.ConstraintRows == 0 && a.PropagatorRows == 0 && a.RepresentationCompleteRows == 0 && a.QuarticZeroBetaRows == 0 && a.QuarticBlockBetaRows == 0 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && a.BetaPermissionFirewallClosed && a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}}
	}}
}
