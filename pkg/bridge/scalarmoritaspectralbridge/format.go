package scalarmoritaspectralbridge

import (
	"fmt"
	"strings"
)

func FormatInheritance(a Gate275Inheritance) string {
	return fmt.Sprintf("solved=%t twoBranch=%t unique=%t J=%t Y=%t a2a4=%t higgs=%t firewall=%t branches=%d verdict=%s", a.ScalarMoritaSolved, a.TwoBranchXYConstrained, a.UniqueXYLocked, a.PhysicalJDerived, a.HyperchargeDerived, a.A2A4Derived, a.HiggsRatioClaimed, a.FirewallPreserved, a.InheritedBranchCount, a.Verdict)
}

func FormatBridge(a ScalarMoritaBridgeTheorem) string {
	return fmt.Sprintf("lambda=%d/%d≈%.15g kappa=%d:%d eq=%q quadratic=%q roots=%q crossTower=%t shapeOnly=%t a2a4=%t verdict=%s", a.LambdaNumerator, a.LambdaDenominator, a.Lambda, a.KappaC, a.KappaQ, a.Equation, a.Quadratic, a.RootForm, a.CrossTowerBridge, a.ScaleFreeShapeOnly, a.EquivalentToA2A4, a.Verdict)
}

func FormatBranch(a SpectralMomentBranch) string {
	return fmt.Sprintf("%s r=%s≈%.15g |y/x|≈%.15g D2(x=1)=%.15g D4(x=1)=%.15g λ≈%.15g residual=%.3g D4/D2(x=1)=%.15g D2/D4(x=1)=%.15g scaleDependent=%t a2a4Claim=%t note=%q", a.Name, a.ExactRForm, a.R, a.AbsYOverX, a.D2ForXEqualsOne, a.D4ForXEqualsOne, a.ShapeLambda, a.ShapeResidualAbs, a.D4OverD2ForXEqualsOne, a.D2OverD4ForXEqualsOne, a.D4OverD2DependsOnScale, a.A2A4CandidateClaimed, a.Interpretation)
}

func FormatBranches(branches []SpectralMomentBranch) string {
	parts := make([]string, 0, len(branches))
	for _, b := range branches {
		parts = append(parts, FormatBranch(b))
	}
	return strings.Join(parts, "; ")
}

func FormatSelectorCandidate(a BranchSelectorCandidate) string {
	return fmt.Sprintf("%s inputs=%q tests=%t upper=%t lower=%t unique=%t selected=%s verdict=%s", a.Name, a.Inputs, a.TestsBranches, a.UpperPasses, a.LowerPasses, a.SelectsUnique, a.Selected, a.Verdict)
}

func FormatSelector(a BranchSelectorAudit) string {
	parts := make([]string, 0, len(a.Candidates))
	for _, c := range a.Candidates {
		parts = append(parts, FormatSelectorCandidate(c))
	}
	return fmt.Sprintf("upper=%t lower=%t unique=%t selected=%s finiteSelector=%t future=%t candidates=[%s] verdict=%s", a.UpperBranchAllowed, a.LowerBranchAllowed, a.UniqueBranch, a.SelectedBranch, a.FiniteCoreSelector, a.RequiresFutureInput, strings.Join(parts, "; "), a.Verdict)
}

func FormatHeatKernel(a HeatKernelProjectionAudit) string {
	return fmt.Sprintf("expansion=%q raw=%q a2=%q a4=%q cutoff=%t subtraction=%t gauge=%t scalar=%t J=%t Y=%t norm=%t map=%t verdict=%s", a.FormalExpansion, a.RawFiniteMomentShape, a.A2Candidate, a.A4Candidate, a.CutoffMomentsSpecified, a.SubtractionSchemeDerived, a.GaugeKineticProjection, a.ScalarFluctuationMap, a.PhysicalJAvailable, a.HyperchargeAvailable, a.FieldNormalizationDerived, a.CanMapRawTracesToA2A4, a.Verdict)
}

func FormatHiggs(a HiggsRatioAudit) string {
	return fmt.Sprintf("selectedBranch=%t absoluteScale=%t heatMap=%t fieldNorm=%t branches=%d a2a4=%t higgs=%t candidates=[%s] verdict=%s", a.UsesSelectedBranch, a.UsesAbsoluteDFScale, a.UsesHeatKernelMap, a.UsesFieldNormalization, a.BranchCount, a.InvariantA2A4Computed, a.HiggsMassRatioComputed, strings.Join(a.CandidateOnlyStatements, "; "), a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("noMass=%t noVEV=%t noCKM=%t noYukawaAmp=%t noTracePromotion=%t candidatesNotPredictions=%t empiricalSeal=%t polluted=%t verdict=%s", a.NoObservedMassInserted, a.NoVEVInserted, a.NoCKMPMNSInserted, a.NoEmpiricalYukawaAmplitudeInserted, a.RawTraceShapeNotPromoted, a.CandidateBranchesNotPredictions, a.EmpiricalYukawaSealPreserved, a.FiniteCorePolluted, a.Verdict)
}

func FormatFuture(a FutureMap) string {
	missing := []string{}
	for _, c := range a.Criteria {
		if c.Required && !c.Satisfied {
			missing = append(missing, c.Name+": "+c.Detail)
		}
	}
	return fmt.Sprintf("criteria=%d missing=[%s] branch=%t scale=%t J=%t Y=%t heat=%t norm=%t next=%q verdict=%q", len(a.Criteria), strings.Join(missing, "; "), a.NeedBranchSelector, a.NeedAbsoluteScale, a.NeedPhysicalJ, a.NeedHypercharge, a.NeedHeatKernelProjection, a.NeedFieldNormalization, a.RecommendedNextGate, a.Verdict)
}

func FormatSummary(a Summary) string {
	return fmt.Sprintf("inherit=%t bridge=%t branches=%t unique=%t heatFormal=%t heatDerived=%t a2a4=%t higgs=%t firewall=%t status=%s next=%q comment=%q", a.Gate275Inherited, a.BridgeFormalized, a.TwoBranchesCarried, a.UniqueBranchSelected, a.HeatKernelFormalized, a.HeatKernelDerived, a.A2A4Derived, a.HiggsRatioClaimed, a.FirewallPreserved, a.Status, a.NextGate, a.Comment)
}
