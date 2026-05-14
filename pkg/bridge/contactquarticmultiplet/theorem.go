package contactquarticmultiplet

import (
	"fmt"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func QuarticBlockMultipletRepresentationBetaIndexObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-QUARTIC-BLOCK-MULTIPLET-REPRESENTATION-BETA-INDEX-OBSTRUCTION"
	const name = "Quartic block multiplet representation / beta-index obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build quartic block multiplet obstruction theorem", Passed: false, Detail: err.Error()}}}
		}
		candidateDetails := make([]string, 0, len(a.MultipletAudit.Candidates))
		for _, c := range a.MultipletAudit.Candidates {
			candidateDetails = append(candidateDetails, FormatCandidate(c))
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 154 exact quartic block compression is inherited", Passed: a.Previous.BetaPermissionFirewallClosed && a.Previous.QuarticCompressedBlocks == 1 && a.Previous.QuarticBlockInvariants == 4 && a.Previous.QuarticOrbitRows == 4 && a.Previous.QuarticBlockBetaRows == 0 && a.Previous.ContactBetaRowsAllowed == 0 && a.Previous.RepresentationCompleteRows == 0, Detail: fmt.Sprintf("quarticRows=%d blocks=%d invariants=%d quarticBeta=%d contactBeta=%d repr=%d", a.Previous.QuarticOrbitRows, a.Previous.QuarticCompressedBlocks, a.Previous.QuarticBlockInvariants, a.Previous.QuarticBlockBetaRows, a.Previous.ContactBetaRowsAllowed, a.Previous.RepresentationCompleteRows)},
			{Name: "dimension-four multiplet candidates are audited but not completed", Passed: a.MultipletAudit.QuarticBlockRows == 4 && len(a.MultipletAudit.Candidates) == 4 && a.MultipletAudit.DimensionMatches == 4 && a.MultipletAudit.RepresentationComplete == 0 && a.MultipletAudit.BetaPermitted == 0, Detail: FormatMultipletAudit(a.MultipletAudit) + " :: " + strings.Join(candidateDetails, " | ")},
			{Name: "Dynkin/beta index requirements are not satisfied", Passed: a.DynkinAudit.BlockRows == 4 && a.DynkinAudit.GroupActionRows == 0 && a.DynkinAudit.RepresentationRows == 0 && a.DynkinAudit.TraceNormalizationRows == 0 && a.DynkinAudit.SpinStatisticsRows == 0 && a.DynkinAudit.MultiplicityRows == 0 && a.DynkinAudit.MassThresholdRows == 0 && a.DynkinAudit.DecouplingRows == 0 && a.DynkinAudit.DynkinIndexRows == 0 && a.DynkinAudit.BetaIndexRows == 0 && !a.DynkinAudit.AllRequirementsSatisfied, Detail: FormatDynkinAudit(a.DynkinAudit)},
			{Name: "quartic invariants remain spectral, not charge or representation semantics", Passed: a.InvariantUse.SymmetricInvariants == 4 && a.InvariantUse.DegreeInvariant && a.InvariantUse.TraceInvariant && a.InvariantUse.DeterminantInvariant && a.InvariantUse.CharacteristicDataExact && !a.InvariantUse.RepresentationSemantics && !a.InvariantUse.DynkinIndexSemantics && !a.InvariantUse.ChargeSemantics, Detail: FormatInvariantUse(a.InvariantUse)},
			{Name: "beta firewall remains closed after multiplet audit", Passed: a.Firewall.ObservedInputFree && a.Firewall.QuarticBlockExact && a.Firewall.BlockCompressionBranchFree && !a.Firewall.MultipletRepresentation && !a.Firewall.DynkinIndex && !a.Firewall.LocalFieldMap && !a.Firewall.MassActivation && !a.Firewall.Decoupling && !a.Firewall.ThresholdBetaRows && !a.Firewall.PhysicalConstants && a.Firewall.FirewallClosed, Detail: FormatFirewall(a.Firewall)},
			{Name: "contact beta and physical constants remain sealed", Passed: a.ContactRows == 7 && a.QuarticOrbitRows == 4 && a.QuarticCompressedBlocks == 1 && a.QuarticBlockInvariants == 4 && a.IndividualQuarticRows == 0 && a.CanonicalQuarticBranches == 0 && a.RowwiseRootAssignmentProofs == 0 && a.ChargeSemanticRows == 0 && a.T3RRowsDerived == 0 && a.BMinusLRowsDerived == 0 && a.HyperchargeRowsDerived == 0 && a.GaugeRepresentationRows == 0 && a.SpinStatisticsRows == 0 && a.LocalFieldRows == 0 && a.MassActivationRows == 0 && a.DecouplingRows == 0 && a.DynkinIndexRows == 0 && a.RepresentationCompleteRows == 0 && a.QuarticBlockBetaRows == 0 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && a.BetaPermissionFirewallClosed && a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}}
	}}
}
