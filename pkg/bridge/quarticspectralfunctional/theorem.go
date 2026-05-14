package quarticspectralfunctional

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CollectiveQuarticSpectralFunctionalActionLevelContributionTheorem() theorem.Theorem {
	const id = "BRIDGE-QUARTIC-COLLECTIVE-SPECTRAL-FUNCTIONAL-ACTION-LEVEL-FIREWALL"
	const name = "Collective quartic spectral functional / action-level contribution theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build quartic spectral functional theorem", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 160 external-selector firewall is inherited", Passed: a.Previous.BetaPermissionFirewallClosed && a.Previous.QuarticOrbitRows == 4 && a.Previous.CanonicalTwoTwoSplits == 0 && a.Previous.BranchBreakingSources == 0 && a.Previous.ContactBetaRowsAllowed == 0 && a.Previous.ContactZeroRowsProved == 0, Detail: fmt.Sprintf("quarticRows=%d twoTwo=%d branchSources=%d beta=%d zero=%d", a.Previous.QuarticOrbitRows, a.Previous.CanonicalTwoTwoSplits, a.Previous.BranchBreakingSources, a.Previous.ContactBetaRowsAllowed, a.Previous.ContactZeroRowsProved)},
			{Name: "quartic polynomial has exact Q-symmetric invariants", Passed: a.Polynomial.Degree == 4 && a.Polynomial.Leading == 3240 && a.Polynomial.Cubic == -7668 && a.Polynomial.Quadratic == 6426 && a.Polynomial.Linear == -2235 && a.Polynomial.Constant == 271 && a.Polynomial.Sum.Equal(NewRational(71, 30)) && a.Polynomial.PairSum.Equal(NewRational(119, 60)) && a.Polynomial.TripleSum.Equal(NewRational(149, 216)) && a.Polynomial.Product.Equal(NewRational(271, 3240)) && a.Polynomial.ExactOverQ && a.Polynomial.GaloisInvariant, Detail: FormatPolynomial(a.Polynomial)},
			{Name: "Newton power sums are exact and branch-free", Passed: len(a.Moments) == 10 && a.Moments[0].Value.Equal(NewRational(71, 30)) && a.Moments[1].Value.Equal(NewRational(1471, 900)) && a.Moments[2].Value.Equal(NewRational(33581, 27000)) && a.Moments[3].Value.Equal(NewRational(809891, 810000)) && a.Moments[4].Value.Equal(NewRational(2235, 271)) && a.MomentAudit.MomentsComputed == 10 && a.MomentAudit.ExactRationalMoments == 10 && a.MomentAudit.GaloisInvariantMoments == 10 && a.MomentAudit.BranchFreeMoments == 10 && a.MomentAudit.BranchChoicesUsed == 0, Detail: FormatMomentAudit(a.MomentAudit) + " :: " + FormatMoments(a.Moments)},
			{Name: "full contact spectral moments combine rational singletons and quartic block", Passed: a.Moments[5].Value.Equal(NewRational(58, 15)) && a.Moments[6].Value.Equal(NewRational(61, 25)) && a.Moments[7].Value.Equal(NewRational(11489, 6750)) && a.Moments[8].Value.Equal(NewRational(257629, 202500)) && a.Moments[9].Value.Equal(NewRational(7993, 542)), Detail: FormatMoments(a.Moments[5:])},
			{Name: "collective candidate functionals do not reproduce existing boundary selectors", Passed: a.BoundaryAudit.CandidatesAudited == 7 && !a.BoundaryAudit.ObservedInputsUsed && a.BoundaryAudit.MatchesKappaU1 == 0 && a.BoundaryAudit.MatchesEmbeddedBoundary == 0 && a.BoundaryAudit.MatchesContactWeakAngle == 0 && a.BoundaryAudit.MatchesGeneratorWeakAngle == 0 && a.BoundaryAudit.BoundaryConstraintsDerived == 0 && !a.BoundaryAudit.GaugeKineticHessianDerived && !a.BoundaryAudit.U1CompletionDerived && !a.BoundaryAudit.NewWeakAngleDerived, Detail: FormatBoundaryAudit(a.BoundaryAudit) + " :: " + FormatCandidateList(a.FunctionalCandidates)},
			{Name: "collective spectral data preserve the beta-permission firewall", Passed: a.BetaFirewall.CollectiveSpectralData && a.BetaFirewall.IndividualQuarticRows == 0 && a.BetaFirewall.GaugeRepresentationRows == 0 && a.BetaFirewall.SpinStatisticsRows == 0 && a.BetaFirewall.LocalFieldRows == 0 && a.BetaFirewall.MassActivationRows == 0 && a.BetaFirewall.DecouplingRows == 0 && a.BetaFirewall.DynkinIndexRows == 0 && a.BetaFirewall.ThresholdBetaRows == 0 && a.BetaFirewall.ProvenZeroRows == 0 && !a.BetaFirewall.PhysicalConstantsDerived && a.BetaFirewall.BetaPermissionFirewallClosed, Detail: FormatBetaFirewall(a.BetaFirewall)},
			{Name: "contact beta rows and physical constants remain sealed", Passed: a.ContactRows == 7 && a.RationalSingletonRows == 3 && a.QuarticOrbitRows == 4 && a.QuarticCollectiveBlocks == 1 && a.QuarticSpectralMoments == 10 && a.IndividualQuarticRows == 0 && a.CanonicalQuarticBranches == 0 && a.ExternalSelectorRows == 0 && a.CanonicalTwoTwoSplits == 0 && a.BranchBreakingSources == 0 && a.GaugeRepresentationRows == 0 && a.LocalFieldRows == 0 && a.MassActivationRows == 0 && a.DecouplingRows == 0 && a.DynkinIndexRows == 0 && a.QuarticBlockBetaRows == 0 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && a.BetaPermissionFirewallClosed && !a.ThresholdCorrectedBeta && !a.FullBetaMatchingTensor && a.BoundaryConstraintsDerived == 0 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3 && !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: FormatSummary(a.Summary) + " :: " + a.TruthStatement},
		}}
	}}
}
