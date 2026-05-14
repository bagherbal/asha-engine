package quarticexternalselector

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func QuarticParityBranchBreakingExternalSelectorFirewallTheorem() theorem.Theorem {
	const id = "BRIDGE-QUARTIC-EXTERNAL-SELECTOR-FIREWALL"
	const name = "Quartic parity branch-breaking external-selector firewall theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build quartic external-selector theorem", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 159 quartic grading obstruction is inherited", Passed: a.Previous.BetaPermissionFirewallClosed && a.Previous.QuarticOrbitRows == 4 && a.Previous.CanonicalQuarticBranches == 0 && a.Previous.QuarticBlockBetaRows == 0 && a.Previous.ContactBetaRowsAllowed == 0 && a.Previous.ContactZeroRowsProved == 0, Detail: fmt.Sprintf("quarticRows=%d branches=%d quarticBeta=%d contactBeta=%d zero=%d", a.Previous.QuarticOrbitRows, a.Previous.CanonicalQuarticBranches, a.Previous.QuarticBlockBetaRows, a.Previous.ContactBetaRowsAllowed, a.Previous.ContactZeroRowsProved)},
			{Name: "all five external selector candidates are audited", Passed: len(a.Candidates) == 5 && a.SourceAudit.CandidatesAudited == 5 && a.SourceAudit.SourcesAvailable == 5, Detail: FormatAudit(a.SourceAudit)},
			{Name: "scalar vacuum orientation does not canonically reach the quartic block", Passed: a.Candidates[0].Kind == ScalarVacuumOrientation && a.Candidates[0].SourceAvailable && !a.Candidates[0].SourceCanonical && !a.Candidates[0].CanonicalMapToQuartic && !a.Candidates[0].ReachesQuarticBlock && !a.Candidates[0].TwoTwoSplit, Detail: FormatCandidate(a.Candidates[0])},
			{Name: "broken gauge images lack a protected-contact/quartic intertwiner", Passed: a.Candidates[1].Kind == BrokenGaugeImages && a.Candidates[1].SourceAvailable && a.Candidates[1].SourceCanonical && !a.Candidates[1].CanonicalMapToContact && !a.Candidates[1].CanonicalMapToQuartic && !a.Candidates[1].ReachesQuarticBlock && !a.Candidates[1].TwoTwoSplit, Detail: FormatCandidate(a.Candidates[1])},
			{Name: "matter B−L has no canonical Fock-to-contact quartic pullback", Passed: a.Candidates[2].Kind == MatterBLCharge && a.Candidates[2].SourceAvailable && !a.Candidates[2].CanonicalMapToContact && a.FockContactKernel.BMinusLPullbackRowsDerived == 0 && a.FockContactKernel.FullOperatorIntertwinersDerived == 0 && !a.Candidates[2].ReachesQuarticBlock && !a.Candidates[2].TwoTwoSplit, Detail: FormatCandidate(a.Candidates[2])},
			{Name: "action second variation restricts only isotropically on the quartic block", Passed: a.Candidates[3].Kind == ActionSecondVariation && a.Candidates[3].SourceAvailable && a.Candidates[3].SourceCanonical && a.Candidates[3].CanonicalMapToQuartic && a.Candidates[3].ReachesQuarticBlock && a.Candidates[3].ProjectionRank == 4 && a.Candidates[3].DistinctEigenvalues == 1 && !a.Candidates[3].NonDegenerateSpectrum && !a.Candidates[3].TwoTwoSplit && !a.Candidates[3].RequiresBranchChoice, Detail: FormatCandidate(a.Candidates[3])},
			{Name: "rational-quartic cross-coupling vanishes by spectral orthogonality", Passed: a.Candidates[4].Kind == RationalQuarticCoupling && a.Candidates[4].SourceAvailable && a.Candidates[4].CanonicalMapToQuartic && a.CrossAudit.OrthogonalProjectors && a.CrossAudit.CrossTermRank == 0 && a.CrossAudit.CrossTermFrobeniusNorm == 0 && !a.CrossAudit.ProvidesSelector && !a.Candidates[4].TwoTwoSplit, Detail: FormatCross(a.CrossAudit)},
			{Name: "no canonical branch-breaking selector is derived", Passed: a.SourceAudit.SuccessfulSelectors == 0 && a.SourceAudit.NonDegenerateSpectra == 0 && a.SourceAudit.TwoTwoSplits == 0 && a.CanonicalTwoTwoSplits == 0 && a.BranchBreakingSources == 0 && a.ExternalSelectorRows == 0, Detail: FormatAudit(a.SourceAudit)},
			{Name: "external-selector firewall remains closed and physical constants remain sealed", Passed: a.Firewall.ObservedInputFree && a.Firewall.Gate159Inherited && a.Firewall.QuarticBlockExact && a.Firewall.ExternalSourcesAudited == 5 && !a.Firewall.BranchBreakingSourceDerived && !a.Firewall.GhostGradingDerived && !a.Firewall.BRSTCancellationDerived && a.Firewall.RepresentationRows == 0 && a.Firewall.ThresholdBetaRows == 0 && a.Firewall.ProvenZeroRows == 0 && !a.Firewall.PhysicalConstantsDerived && a.Firewall.FirewallClosed && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: FormatFirewall(a.Firewall) + " :: " + FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}}
	}}
}
