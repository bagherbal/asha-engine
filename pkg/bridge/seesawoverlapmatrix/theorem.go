package seesawoverlapmatrix

import "github.com/bagherbal/asha-engine/pkg/theorem"

func SeesawOverlapMatrixConstructionMajoranaHiggsMixingSieveAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-SEESAW-OVERLAP-MATRIX-CONSTRUCTION-MAJORANA-HIGGS-MIXING-SIEVE-AUDIT"
	const name = "Seesaw Overlap Matrix Construction / Majorana-Higgs Mixing Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 320 seesaw overlap audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "doubled-space L_L, nu_R, nu_R^c block is formalized with J_swap Majorana edge", Passed: a.Block.Formalized && a.Block.JSwapInstalled && len(a.Block.Basis) == 3 && a.Block.HiggsEdge.From == "L_L" && a.Block.HiggsEdge.To == "nu_R" && a.Block.MajoranaEdge.To == "nu_R^c" && a.Block.DirectSumOverlap == 0, Detail: FormatBlock(a.Block)},
			{Name: "sequential seesaw path L_L -> nu_R -> nu_R^c is explicitly constructed", Passed: a.Path.Constructed && a.Path.UsesJSwap && a.Path.HiggsEdgeExists && a.Path.MajoranaEdgeExists && a.Path.PathCount == 1 && a.Path.DirectSumPathCount == 0 && a.Path.PathMatrixRank == 1 && a.Path.PathWeight > 0.10 && a.Path.PathWeight < 0.103, Detail: FormatPath(a.Path)},
			{Name: "Omega_Hsigma support matrix is derived and has canonical overlap index one", Passed: a.Overlap.Derived && a.Overlap.IndexVerified && a.Overlap.TraceOmegaDagOmega == 1 && a.Overlap.CanonicalOverlapIndex == 1 && a.Overlap.DirectSumIndex == 0 && a.Overlap.UniquePathNormalized, Detail: FormatOverlap(a.Overlap)},
			{Name: "B-gap portal weight is enabled and still matches Gate-314 target within one percent", Passed: a.Portal.Enabled && a.Portal.WeightsMultiplicative && !a.Portal.ThresholdPromoted && a.Portal.Coefficient > 0.39 && a.Portal.Coefficient < 0.392 && a.Portal.WithinOnePercent && a.Portal.ImpliedDeltaLambda < -0.097 && a.Portal.ImpliedDeltaLambda > -0.099, Detail: FormatPortal(a.Portal)},
			{Name: "promotion remains blocked by heavy propagator, self-quartic, and coupling normalization", Passed: a.Promotion.ExplicitMatrixDerived && a.Promotion.OverlapIndexDerived && !a.Promotion.HeavyPropagatorDerived && !a.Promotion.HeavySelfQuarticDerived && !a.Promotion.LambdaMixNormalized && !a.Promotion.ThresholdJumpDerived && !a.Promotion.PortalPromotionAuthorized, Detail: FormatPromotion(a.Promotion)},
			{Name: "firewalls preserve no final Higgs mass or threshold theorem claim", Passed: a.Firewalls.NoFinalMassClaimed && a.Firewalls.NoPoleMassClaimed && a.Firewalls.NoThresholdClaimed && a.Firewalls.NoHeavyPropagatorClaimed && a.Firewalls.NoHeavyQuarticClaimed && a.Firewalls.NoLambdaMixClaimed && !a.Firewalls.FiniteCorePolluted, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "summary records explicit overlap success without over-promoting the portal", Passed: a.Summary.DoubledBlockFormalized && a.Summary.SeesawPathConstructed && a.Summary.ExplicitMatrixDerived && a.Summary.OverlapIndexVerified && a.Summary.PortalWeightEnabled && !a.Summary.ThresholdPromoted && a.Summary.FirewallsPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 320 resolves the Gate-319 overlap-index obstruction: the doubled seesaw support graph supplies a unique normalized sigma-H path with Omega_Hsigma=1.", "The portal resonance is structurally enabled but not yet a full threshold theorem until heavy propagator and heavy self-quartic normalizations are derived."}}
	}}
}
