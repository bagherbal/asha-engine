package generation2boundaryweighteddeficitclosurerobustnessaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate660Inheritance) string {
	return fmt.Sprintf("activeInherited=%t W72=%.15g Ksum=%.15g weightedResidual=%.15g formulaExact=%.15g formulaOrient=%.15g bridgeOnly=%t no7=%t noK7Map=%t noTransport=%t noFano=%t firewall=%t verdict=%q", x.ActiveWeightInherited, x.W72, x.KSum, x.WeightedResidual, x.FormulaLiftResidualExact, x.FormulaLiftResidualOrient, x.FormulaLiftBridgeLayerOnly, x.NoNativeSevenOver72Theorem, x.NoNativeK7BoundaryMap, x.NoNativeTransportTheorem, x.NoFanoHitchinBoundaryRevival, x.FirewallPreserved, x.Verdict)
}

func FormatDependencyNode(x DependencyNode) string {
	return fmt.Sprintf("%s=%.15g role=%q source=%q depends=[%s] independent=%t derived=%t circularity=%q", x.Name, x.Value, x.Role, x.Source, strings.Join(x.DependsOn, ","), x.Independent, x.Derived, x.Circularity)
}

func FormatDependencyGraph(x DependencyGraphAudit) string {
	parts := make([]string, 0, len(x.Nodes))
	for _, n := range x.Nodes {
		parts = append(parts, FormatDependencyNode(n))
	}
	return fmt.Sprintf("nodes=%d kappaFromRuntime=%t lambda12FromRuntime=%t W72FromEndpoints=%t formulaTautological=%t nontrivial=%q verdict=%q ledger=[%s]", len(x.Nodes), x.KappaLambdaDefinedFromRuntime, x.LambdaLambda12DependsOnRuntime, x.W72DependsOnBoundaryEndpoints, x.FormulaLiftPartlyTautological, x.NontrivialStatement, x.Verdict, strings.Join(parts, "; "))
}

func FormatClosure(x NontrivialClosureAudit) string {
	return fmt.Sprintf("kappaLambda=%.15g kappaE=%.15g Ksum=%.15g W72=%.15g residual=%.15g relW72=%.15g relSplit=%.15g formulaResidual=%.15g formulaIndependent=%t statement=%q verdict=%q", x.KappaLambda, x.KappaEExact, x.KSumExact, x.W72, x.ClosureResidualExact, x.RelativeToW72, x.RelativeToBoundarySplit, x.ScalarFormulaLiftResidual, x.FormulaLiftIndependent, x.NontrivialBridgeStatement, x.Verdict)
}

func FormatOrientation(x OrientationApproximationAudit) string {
	return fmt.Sprintf("kappaExact=%.15g kappaOrient=%.15g delta=%.15g KsumOrient=%.15g residualOrient=%.15g relW72=%.15g relSplit=%.15g orient/exact=%.15g verdict=%q", x.KappaEExact, x.KappaEOrientation, x.KappaEDifference, x.KSumOrientation, x.ClosureResidualOrientation, x.RelativeResidualOrientationToW72, x.RelativeResidualOrientationToSplit, x.ExactToOrientationResidualRatio, x.Verdict)
}

func FormatUncertaintySlot(x UncertaintySlot) string {
	return fmt.Sprintf("%s neededFor=%q available=%t treatment=%q perturb=%q", x.Quantity, x.NeededFor, x.Available, x.Treatment, x.WouldPerturb)
}

func FormatUncertainty(x UncertaintyAudit) string {
	parts := make([]string, 0, len(x.Slots))
	for _, s := range x.Slots {
		parts = append(parts, FormatUncertaintySlot(s))
	}
	return fmt.Sprintf("slots=%d fullPropagation=%t invented=%t significance=%t verdict=%q ledger=[%s]", len(x.Slots), x.FullPropagationAvailable, x.InventedUncertainties, x.ClosureSignificanceCertified, x.Verdict, strings.Join(parts, "; "))
}

func FormatScaleRow(x ScaleSensitivityRow) string {
	return fmt.Sprintf("%s available=%t required=%q treatment=%q compare=%t", x.Scale, x.Available, x.RequiredData, x.CurrentTreatment, x.CanCompareClosure)
}

func FormatScale(x ScaleSensitivityAudit) string {
	parts := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		parts = append(parts, FormatScaleRow(r))
	}
	return fmt.Sprintf("rows=%d lambda12Only=%t nearbySweep=%t endpointIndependence=%t verdict=%q ledger=[%s]", len(x.Rows), x.Lambda12OnlyComputed, x.NearbyScaleSweepAvailable, x.EndpointIndependenceCertified, x.Verdict, strings.Join(parts, "; "))
}

func FormatWeightRow(x TypedWeightRow) string {
	return fmt.Sprintf("%s weight=%.15g target=%.15g residual=%.15g abs=%.15g role=%q allowed=%t", x.Name, x.Weight, x.Target, x.Residual, x.AbsResidual, x.TypedRole, x.Allowed)
}

func FormatWeights(x TypedWeightUniquenessAudit) string {
	parts := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		parts = append(parts, FormatWeightRow(r))
	}
	return fmt.Sprintf("best=%q bestResidual=%.15g second=%q secondResidual=%.15g improvement=%.15g noArbitrary=%t verdict=%q rows=[%s]", x.BestName, x.BestResidual, x.SecondBestName, x.SecondBestResidual, x.ImprovementOverSecond, x.NoArbitrarySearch, x.Verdict, strings.Join(parts, "; "))
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("robustV1=%t pendingUncertainty=%t claims7=%t claimsTransport=%t claimsEndpoint=%t claimsBoundary=%t claimsHiggs=%t claimsStability=%t claimsFlavor=%t claimsCKM=%t claimsGauge=%t claimsFano=%t verdict=%q", x.ClassifiesRobustV1ExactLedger, x.ClassifiesPendingUncertaintySweep, x.ClaimsNativeSevenOver72Theorem, x.ClaimsNativeTransportTheorem, x.ClaimsIndependentEndpointDerivation, x.ClaimsBoundaryStressDerivation, x.ClaimsHiggsPrediction, x.ClaimsScalarStability, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNSDerivation, x.ClaimsGaugeUnification, x.ClaimsFanoHitchinBoundaryMap, x.Verdict)
}
