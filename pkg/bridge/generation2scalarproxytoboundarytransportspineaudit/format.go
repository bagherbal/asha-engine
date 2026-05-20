package generation2scalarproxytoboundarytransportspineaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate657Inheritance) string {
	return fmt.Sprintf("pivot=%t fanoClosed=%t activeVector=%t primaryRG=%t scalar=%t boundary=%t history=%t k7Blocked=%t noFanoMap=%t noTrace=%t firewall=%t verdict=%q", x.TransportPivotInherited, x.FanoBoundaryClosed, x.ActiveBridgeVectorBuilt, x.PrimaryWasRGTransport, x.ScalarMatchingActive, x.BoundaryStressActive, x.HistoryLoopActive, x.K7BoundaryBlocked, x.NoFanoBoundaryMap, x.NoSevenTraceTheorem, x.FirewallPreserved, x.Verdict)
}

func FormatProxy(x ProxyLane) string {
	return fmt.Sprintf("formula=%q lambda_proxy=%.15g b/a2=%.15g one8=%.15g diff=%.15g rel=%.15g close=%t treeOnly=%t cannotReplaceRuntime=%t verdict=%q", x.Formula, x.LambdaProxyMZ, x.BA2Ratio, x.OneEighth, x.DifferenceFromOne8, x.RelativeFromOne8, x.CloseToOneEighth, x.TreeProxyOnly, x.CannotReplaceRuntime, x.Verdict)
}

func FormatMatching(x LowScaleMatchingLane) string {
	return fmt.Sprintf("proxy=%.15g runtime=%.15g delta=%.15g rho=%.15g L=%.15g kappa=%.15g oneMinusKappa=%.15g reconstructed=%.15g residual=%.15g rawL=%.15g rawResidual=%.15g loopSized=%t verdict=%q", x.LambdaProxyMZ, x.LambdaRuntimeMZ, x.DeltaLambdaMatch, x.RelativeMatch, x.HistoryLoopUnit, x.KappaLambda, x.OneMinusKappaLambda, x.ReconstructedRuntimeMZ, x.ReconstructionResidual, x.RawLAnsatzRuntime, x.RawLAnsatzResidual, x.LoopSized, x.Verdict)
}

func FormatRG(x RGTransportLane) string {
	return fmt.Sprintf("start=%q boundary=%q lambdaStart=%.15g lambdaBoundary=%.15g absBoundary=%.15g negative=%t v1=%t claimsThreshold=%t verdict=%q", x.StartScale, x.BoundaryScale, x.LambdaRuntimeStart, x.LambdaBoundary, x.AbsLambdaBoundary, x.RuntimeTurnsNegative, x.UsesCurrentV1RG, x.ClaimsThresholdLaw, x.Verdict)
}

func FormatBoundary(x BoundaryStressLane) string {
	return fmt.Sprintf("absLambda=%.15g r3=%.15g xi=%.15g mean=%.15g split=%.15g half=%.15g absMinusXi=%.15g r3MinusXi=%.15g form=%q xiPreferred=%t verdict=%q", x.AbsLambdaBoundary, x.R3Minus1, x.XiBoundary, x.MeanStressRecomputed, x.BoundarySplit, x.HalfSplit, x.AbsLambdaResidualToXi, x.R3ResidualToXi, x.AntiAlignmentForm, x.XiPreferredOverHalfTrace, x.Verdict)
}

func FormatResiduals(x ResidualDecomposition) string {
	parts := make([]string, 0, len(x.Slots))
	for _, s := range x.Slots {
		parts = append(parts, fmt.Sprintf("%s=%.15g scale=%q status=%q requires=%q", s.Name, s.Value, s.Scale, s.TypedStatus, s.RequiresSource))
	}
	return fmt.Sprintf("match=%t rg=%t boundary=%t thresholdOpen=%t verdict=%q slots=%s", x.MatchSlotSeparated, x.RGSlotSeparated, x.BoundarySlotSeparated, x.ThresholdSlotsOpen, x.Verdict, strings.Join(parts, "; "))
}

func FormatSources(x SourceAudit) string {
	return fmt.Sprintf("kappa=%t xi=%t L=%t proxyRuntime=%t rgThreshold=%t boundary=%t random=%t typedOnly=%t verdict=%q", x.KappaLambdaSourceCertified, x.XiBoundarySourceCertified, x.HistoryLoopSourceCertified, x.ProxyRuntimeTheorem, x.RGThresholdTheorem, x.BoundaryStressTheorem, x.SearchedRandomConstants, x.TypedQuantitiesOnly, x.Verdict)
}

func FormatSpine(x SpineClassification) string {
	return fmt.Sprintf("name=%q active=%t bridgeOnly=%t merged=%t touches=%s next=%q verdict=%q", x.Name, x.Active, x.BridgeLayerOnly, x.MergesScalarBoundary, strings.Join(x.Touches, ","), x.NextPressurePoint, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("higgs=%t stability=%t gauge=%t thresholds=%t nativeScalar=%t boundary=%t spacetime=%t flavor=%t verdict=%q", x.ClaimsHiggsMass, x.ClaimsScalarStability, x.ClaimsGaugeUnification, x.ClaimsThresholdExistence, x.ClaimsNativeScalarTheorem, x.ClaimsBoundaryStressDerived, x.ClaimsPhysicalSpacetime, x.ClaimsFlavorTheorem, x.Verdict)
}
