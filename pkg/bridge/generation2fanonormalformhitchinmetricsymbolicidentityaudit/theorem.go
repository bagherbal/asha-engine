package generation2fanonormalformhitchinmetricsymbolicidentityaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2FanoNormalFormHitchinMetricSymbolicIdentityAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 653 — Fano Normal-Form Hitchin Metric Symbolic Identity Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate653 Fano/Hitchin symbolic identity audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate652 Fano normal form and firewall", Passed: a.Inherited.FanoNormalFormInherited && a.Inherited.BVolumeForm && a.Inherited.ATwoFormTriple && a.Inherited.WedgeOrthonormality && a.Inherited.QuaternionicTriple && a.Inherited.AAAChannelFinite && a.Inherited.AABChannelsFinite && a.Inherited.FiniteNormalFormIdentities && !a.Inherited.FullBasisFreeFanoTheorem && !a.Inherited.ClaimsSplitG2 && !a.Inherited.ClaimsBoundaryStress && !a.Inherited.ClaimsSevenOver72 && !a.Inherited.ClaimsScalarFlavor && !a.Inherited.ClaimsPhysicalMetric && a.Inherited.Gate652FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "derive symbolic positive AAA block", Passed: a.Positive.UsesNormalForm && a.Positive.UsesWedgeIdentity && a.Positive.CPositive == unitC && a.Positive.ScalarMultipleOfP && a.Positive.AnisotropyResidual < tol && a.Positive.SymbolicDerivation, Detail: FormatPositive(a.Positive)},
			{Name: "derive symbolic negative AAB/ABA/BAA blocks", Passed: a.Negative.EachEqualsMinusC && a.Negative.CombinedCoefficient == -3 && a.Negative.CombinedTarget == -3 && a.Negative.CombinedResidual < tol && a.Negative.NegativeSignLocated && a.Negative.SymbolicDerivation && len(a.Negative.Rows) == 3, Detail: FormatNegative(a.Negative)},
			{Name: "derive symbolic mixed-block vanishing", Passed: a.Mixed.SymbolicallyZero && a.Mixed.MixedBlockNorm < tol && len(a.Mixed.Cases) == 2, Detail: FormatMixed(a.Mixed)},
			{Name: "audit equal c normalization across surviving channels", Passed: a.Normalization.CPositive == unitC && a.Normalization.CAAB == -unitC && a.Normalization.CABA == -unitC && a.Normalization.CBAA == -unitC && a.Normalization.AllEqualAbs && !a.Normalization.RequiresRescale && a.Normalization.Residual < tol, Detail: FormatNormalization(a.Normalization)},
			{Name: "reduce routes to a single Fano symbolic identity", Passed: a.Routes.AllRoutesReduce && a.Routes.SameSymbolicIdentity && !a.Routes.RouteDependentOnly && len(a.Routes.Rows) == 3, Detail: FormatRoutes(a.Routes)},
			{Name: "close internal Hitchin obstruction mechanism under normal-form assumptions", Passed: a.FinalIdentity.PositiveBlockPasses && a.FinalIdentity.NegativeBlockPasses && a.FinalIdentity.MixedBlockPasses && a.FinalIdentity.EqualNormalizationPasses && a.FinalIdentity.InternalMechanismClosed && !a.FinalIdentity.FullPGToFanoSourceTheorem, Detail: FormatFinalIdentity(a.FinalIdentity)},
			{Name: "preserve split-G2, boundary, scalar/flavor, physical, and 7/72 firewalls", Passed: !a.Firewalls.ClaimsSplitG2 && !a.Firewalls.ClaimsBoundaryStress && !a.Firewalls.ClaimsSevenOver72 && !a.Firewalls.ClaimsScalarFlavor && !a.Firewalls.ClaimsPhysicalMetric && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsCKMPMNS && !a.Firewalls.ClaimsGaugeUnification && a.Firewalls.Verdict == StatusGate653Boundary, Detail: FormatFirewalls(a.Firewalls)},
		}
		notes := append(Statuses(), a.Truth)
		notes = append(notes, "Gate653 proves the Fano normal-form-to-Hitchin-metric-ray implication while keeping the separate P_G-to-normal-form source theorem and all physical/boundary promotions firewalled.")
		if a.FinalIdentity.FullPGToFanoSourceTheorem {
			notes = append(notes, "WARNING_PG_TO_FANO_SOURCE_THEOREM_OVERPROMOTED")
		}
		if !strings.Contains(a.FinalIdentity.Verdict, StatusNoBasisFreePGToFanoTheorem) {
			notes = append(notes, "WARNING_MISSING_PG_TO_FANO_FIREWALL")
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
