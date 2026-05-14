package heavylightoverlapoperator

import "github.com/bagherbal/asha-engine/pkg/theorem"

func FunctionalDeterminantSieveHeavyLightOverlapOperatorAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-FUNCTIONAL-DETERMINANT-SIEVE-HEAVY-LIGHT-OVERLAP-OPERATOR-AUDIT"
	const name = "Functional Determinant Sieve / Heavy-Light Overlap Operator Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 319 heavy-light overlap audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "functional determinant expansion is formalized with explicit heavy-propagator and overlap obligations", Passed: a.Determinant.Formalized && a.Determinant.PortalTermOrder == 2 && a.Determinant.NeedsHeavyPropagator && a.Determinant.NeedsOverlapInsertion && a.Determinant.DirectSumFactorizes, Detail: FormatDeterminant(a.Determinant)},
			{Name: "direct-sum determinant lane has zero sigma-H mixing", Passed: a.Operator.DirectSumLane.CrossTermsVanish && !a.Operator.DirectSumLane.OverlapOperatorExists && a.Operator.DirectSumLane.Coefficient == 0 && a.Operator.DirectSumLane.DerivedFromMatrices, Detail: FormatLane(a.Operator.DirectSumLane)},
			{Name: "true-bimodule conditional overlap lane reproduces the near-target kappa_Q*(4/pi)*B_gap coefficient", Passed: a.Operator.TrueBimoduleConditionalLane.OverlapOperatorExists && !a.Operator.TrueBimoduleConditionalLane.CrossTermsVanish && a.Operator.TrueBimoduleConditionalLane.WithinOnePercent && a.Operator.TrueBimoduleConditionalLane.Coefficient > 0.39 && a.Operator.TrueBimoduleConditionalLane.Coefficient < 0.392 && !a.Operator.TrueBimoduleConditionalLane.DerivedFromMatrices, Detail: FormatLane(a.Operator.TrueBimoduleConditionalLane)},
			{Name: "multiplicative coefficient sieve identifies the witness but refuses to force the weights", Passed: a.Sieve.Formalized && a.Sieve.WithinOnePercent && a.Sieve.ConditionalDeltaLambda < -0.097 && a.Sieve.ConditionalDeltaLambda > -0.099 && !a.Sieve.FactorsForcedMultiplicative && a.Sieve.DirectSumValue == 0, Detail: FormatSieve(a.Sieve)},
			{Name: "promotion audit blocks threshold theorem until explicit sigma-H matrix and heavy self-quartic are derived", Passed: a.Promotion.FunctionalDeterminantInstalled && !a.Promotion.SigmaHOverlapOperatorDerived && !a.Promotion.MultiplicativeWeightsForced && !a.Promotion.HeavySelfQuarticDerived && !a.Promotion.ThresholdJumpDerived && !a.Promotion.PromotionAuthorized, Detail: FormatPromotion(a.Promotion)},
			{Name: "firewalls preserve no portal, threshold, RG, pole-mass, or final Higgs-mass claim", Passed: a.Firewalls.NoPortalCouplingClaimed && a.Firewalls.NoThresholdJumpClaimed && a.Firewalls.NoFinalMassClaimed && a.Firewalls.NoRGReexecutionClaimed && a.Firewalls.NoPoleMassClaimed && a.Firewalls.NoExplicitMatrixClaimed && a.Firewalls.NoHeavyQuarticClaimed && !a.Firewalls.FiniteCorePolluted, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "summary records operator target without overclaiming physical promotion", Passed: a.Summary.DeterminantFormalized && a.Summary.DirectSumRejected && a.Summary.TemplateFormalized && a.Summary.WitnessMatchesTarget && !a.Summary.OperatorDerived && !a.Summary.PortalPromoted && a.Summary.FirewallsPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 319 proves that direct-sum carriers cannot generate the portal; the required path is a true-bimodule sigma-H overlap operator.", "The kappa_Q*(4/pi)*B_gap resonance is now a precise overlap-operator target, not yet a derived threshold matching theorem."}}
	}}
}
