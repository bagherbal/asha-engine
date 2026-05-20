package generation2boundaryweighteddeficitclosurerobustnessaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2BoundaryWeightedDeficitClosureRobustnessAndNoncircularityAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 661 — BoundaryWeightedDeficitClosure Robustness and Noncircularity Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate661 robustness audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate660 active boundary weight", Passed: a.Inherited.ActiveWeightInherited && math.Abs(a.Inherited.W72-0.04982659643506822) < 5e-15 && math.Abs(a.Inherited.KSum-0.04982659728765166) < 5e-15 && math.Abs(a.Inherited.WeightedResidual-8.525834413464217e-10) < 5e-18 && math.Abs(a.Inherited.FormulaLiftResidualExact-4.2369718844526005e-12) < 5e-17 && a.Inherited.FormulaLiftBridgeLayerOnly && a.Inherited.NoNativeSevenOver72Theorem && a.Inherited.NoNativeK7BoundaryMap && a.Inherited.NoNativeTransportTheorem && a.Inherited.NoFanoHitchinBoundaryRevival && a.Inherited.FirewallPreserved, Detail: FormatInherited(a.Inherited)},
			{Name: "audit dependency graph and circularity", Passed: len(a.Dependency.Nodes) == 9 && a.Dependency.KappaLambdaDefinedFromRuntime && a.Dependency.LambdaLambda12DependsOnRuntime && a.Dependency.W72DependsOnBoundaryEndpoints && a.Dependency.FormulaLiftPartlyTautological && strings.Contains(a.Dependency.NontrivialStatement, "kappa_lambda + kappa_e"), Detail: FormatDependencyGraph(a.Dependency)},
			{Name: "isolate nontrivial closure", Passed: math.Abs(a.Closure.ClosureResidualExact-8.525834413464217e-10) < 5e-18 && a.Closure.RelativeToBoundarySplit < 1e-6 && !a.Closure.FormulaLiftIndependent && strings.Contains(a.Closure.NontrivialBridgeStatement, "7/72"), Detail: FormatClosure(a.Closure)},
			{Name: "audit orientation approximation", Passed: math.Abs(a.Orientation.KappaEDifference-2.775873137889728e-06) < 5e-18 && math.Abs(a.Orientation.ClosureResidualOrientation-2.7767257213331953e-06) < 5e-18 && a.Orientation.RelativeResidualOrientationToW72 < 6e-5 && a.Orientation.RelativeResidualOrientationToSplit < 0.003 && a.Orientation.ExactToOrientationResidualRatio > 3000, Detail: FormatOrientation(a.Orientation)},
			{Name: "define uncertainty slots without invented propagation", Passed: len(a.Uncertainty.Slots) == 5 && !a.Uncertainty.FullPropagationAvailable && !a.Uncertainty.InventedUncertainties && !a.Uncertainty.ClosureSignificanceCertified, Detail: FormatUncertainty(a.Uncertainty)},
			{Name: "define scale sensitivity slots", Passed: len(a.Scale.Rows) == 5 && a.Scale.Lambda12OnlyComputed && !a.Scale.NearbyScaleSweepAvailable && !a.Scale.EndpointIndependenceCertified, Detail: FormatScale(a.Scale)},
			{Name: "audit typed weight uniqueness", Passed: len(a.Weights.Rows) == 6 && a.Weights.BestName == "7/72" && a.Weights.BestResidual < 1e-9 && a.Weights.SecondBestName == "1/10" && a.Weights.ImprovementOverSecond > 4000 && a.Weights.NoArbitrarySearch, Detail: FormatWeights(a.Weights)},
			{Name: "preserve robustness/noncircularity firewalls", Passed: a.Discipline.ClassifiesRobustV1ExactLedger && a.Discipline.ClassifiesPendingUncertaintySweep && !a.Discipline.ClaimsNativeSevenOver72Theorem && !a.Discipline.ClaimsNativeTransportTheorem && !a.Discipline.ClaimsIndependentEndpointDerivation && !a.Discipline.ClaimsBoundaryStressDerivation && !a.Discipline.ClaimsHiggsPrediction && !a.Discipline.ClaimsScalarStability && !a.Discipline.ClaimsFlavorDerivation && !a.Discipline.ClaimsCKMPMNSDerivation && !a.Discipline.ClaimsGaugeUnification && !a.Discipline.ClaimsFanoHitchinBoundaryMap && a.Discipline.Verdict == StatusGate661Boundary, Detail: FormatDiscipline(a.Discipline)},
		}
		notes := append(Statuses(), a.Truth)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
