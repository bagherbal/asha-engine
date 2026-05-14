package contactquarticlocalfield

import (
	"fmt"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func QuarticBlockLocalFieldSpinStatisticsObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-QUARTIC-BLOCK-LOCAL-FIELD-SPIN-STATISTICS-OBSTRUCTION"
	const name = "Quartic block local-field / spin-statistics obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build quartic block local-field obstruction theorem", Passed: false, Detail: err.Error()}}}
		}
		candidateDetails := make([]string, 0, len(a.FieldCandidates))
		for _, c := range a.FieldCandidates {
			candidateDetails = append(candidateDetails, FormatFieldCandidate(c))
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 155 exact quartic multiplet audit is inherited", Passed: a.Previous.BetaPermissionFirewallClosed && a.Previous.QuarticOrbitRows == 4 && a.Previous.QuarticBlockInvariants == 4 && a.Previous.MultipletAudit.DimensionMatches == 4 && a.Previous.RepresentationCompleteRows == 0 && a.Previous.LocalFieldRows == 0 && a.Previous.SpinStatisticsRows == 0 && a.Previous.QuarticBlockBetaRows == 0 && a.Previous.ContactBetaRowsAllowed == 0, Detail: fmt.Sprintf("quarticRows=%d invariants=%d dimensionMatches=%d repr=%d local=%d spin=%d beta=%d", a.Previous.QuarticOrbitRows, a.Previous.QuarticBlockInvariants, a.Previous.MultipletAudit.DimensionMatches, a.Previous.RepresentationCompleteRows, a.Previous.LocalFieldRows, a.Previous.SpinStatisticsRows, a.Previous.ContactBetaRowsAllowed)},
			{Name: "local-field candidates match degree count but remain incomplete", Passed: len(a.FieldCandidates) == 5 && a.Summary.DegreeMatchingCandidates == 5 && a.Summary.LocalFieldCompleteRows == 0 && a.QuarticOrbitRows == 4, Detail: FormatSummary(a.Summary) + " :: " + strings.Join(candidateDetails, " | ")},
			{Name: "local-field permission requirements are not satisfied", Passed: a.RequirementAudit.QuarticBlockRows == 4 && a.RequirementAudit.BaseSpaceRows == 0 && a.RequirementAudit.LocalSectionRows == 0 && a.RequirementAudit.LorentzRepresentationRows == 0 && a.RequirementAudit.KineticOperatorRows == 0 && a.RequirementAudit.PoleResidueRows == 0 && a.RequirementAudit.SpinStatisticsRows == 0 && a.RequirementAudit.GaugeRepresentationRows == 0 && a.RequirementAudit.HyperchargeRows == 0 && a.RequirementAudit.MassActivationRows == 0 && a.RequirementAudit.DecouplingRows == 0 && !a.RequirementAudit.AllRequirementsSatisfied, Detail: FormatRequirementAudit(a.RequirementAudit)},
			{Name: "spin-statistics branch is not selected", Passed: a.SpinAudit.CandidateRows == 5 && a.SpinAudit.BosonicScalarRows == 0 && a.SpinAudit.FermionicSpinorRows == 0 && a.SpinAudit.GhostRegulatorRows == 0 && a.SpinAudit.AuxiliaryConstrainedRows == 0 && a.SpinAudit.CanonicalCommutationRows == 0 && a.SpinAudit.CanonicalAnticommutationRows == 0 && a.SpinAudit.BRSTGradingRows == 0 && a.SpinAudit.SpinStatisticsComplete == 0, Detail: FormatSpinAudit(a.SpinAudit)},
			{Name: "finite spectral block is exact but not a local propagator", Passed: a.KineticAudit.FiniteSpectralBlockExact && a.KineticAudit.SpectralOperatorRows == 4 && a.KineticAudit.LocalDifferentialOperatorRows == 0 && a.KineticAudit.LorentzSignatureRows == 0 && a.KineticAudit.HyperbolicEllipticClassRows == 0 && a.KineticAudit.PropagatorDenominatorRows == 0 && a.KineticAudit.PositiveResidueRows == 0 && !a.KineticAudit.LocalityComplete, Detail: FormatKineticAudit(a.KineticAudit)},
			{Name: "beta firewall remains closed after local-field audit", Passed: a.Firewall.ObservedInputFree && a.Firewall.QuarticBlockExact && a.Firewall.MultipletDimensionAudited && !a.Firewall.LocalFieldMap && !a.Firewall.SpinStatistics && !a.Firewall.KineticPoleResidue && !a.Firewall.GaugeRepresentation && !a.Firewall.Hypercharge && !a.Firewall.MassActivation && !a.Firewall.Decoupling && !a.Firewall.ThresholdBetaRows && !a.Firewall.PhysicalConstants && a.Firewall.FirewallClosed, Detail: FormatFirewall(a.Firewall)},
			{Name: "physical constants and threshold beta rows remain sealed", Passed: a.ContactRows == 7 && a.QuarticOrbitRows == 4 && a.IndividualQuarticRows == 0 && a.CanonicalQuarticBranches == 0 && a.RowwiseRootAssignmentProofs == 0 && a.ChargeSemanticRows == 0 && a.T3RRowsDerived == 0 && a.BMinusLRowsDerived == 0 && a.HyperchargeRowsDerived == 0 && a.GaugeRepresentationRows == 0 && a.SpinStatisticsRows == 0 && a.LocalFieldRows == 0 && a.KineticPoleResidueRows == 0 && a.MassActivationRows == 0 && a.DecouplingRows == 0 && a.DynkinIndexRows == 0 && a.RepresentationCompleteRows == 0 && a.QuarticBlockBetaRows == 0 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && a.BetaPermissionFirewallClosed && a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}}
	}}
}
