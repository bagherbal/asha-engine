package generation2halftraceboundarycoordinateweightaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(x Gate655Inheritance) string {
	return fmt.Sprintf("seal=%t internal=%t numerator7=%t noBoundary=%t no7over72=%t noStress=%t noScalarFlavor=%t noHistory=%t claimsBoundary=%t claims7=%t claimsScalarFlavor=%t claimsHistory=%t firewall=%t verdict=%q", x.FanoSealDefined, x.FanoSealInternalOnly, x.FanoStructuresNumerator, x.NoBoundaryInterface, x.NoSevenOver72Theorem, x.NoBoundaryStress, x.NoScalarFlavorMap, x.NoHistoryLoopSource, x.ClaimsBoundaryStress, x.ClaimsSevenOver72, x.ClaimsScalarFlavor, x.ClaimsHistoryLoopUnit, x.Gate655Firewall, x.Verdict)
}

func FormatSourceType(x SourceTypeAudit) string {
	parts := make([]string, 0, len(x.Factors))
	for _, f := range x.Factors {
		parts = append(parts, fmt.Sprintf("%s=%g source=%q typed=%t native=%t bridge=%t certified=%t", f.Factor, f.Value, f.Source, f.Typed, f.Native, f.BridgeOnly, f.CertifiedRoute))
	}
	return fmt.Sprintf("wFull=%.15g wHalf=%.15g seven=%t seventyTwo=%t halfTyped=%t halfNative=%t allTyped=%t map=%t verdict=%q factors=%s", x.FullWeight, x.HalfWeight, x.SevenTyped, x.SeventyTwoTyped, x.HalfTyped, x.HalfNative, x.AllFactorsTyped, x.CertifiedHalfTraceMap, x.Verdict, strings.Join(parts, "; "))
}

func FormatBoundaryComparison(x BoundaryComparisonAudit) string {
	parts := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		parts = append(parts, fmt.Sprintf("%s target=%.15g cand=%.15g signed=%.15g abs=%.15g rel=%.5g rank=%d class=%q", r.Target, r.TargetValue, r.Candidate, r.SignedResidual, r.AbsResidual, r.RelativeResidual, r.Rank, r.Classification))
	}
	return fmt.Sprintf("closest=%s residual=%.15g certified=%t noProximity=%t verdict=%q rows=%s", x.ClosestTarget, x.ClosestResidual, x.CertifiedMatch, x.NoProximityCertification, x.Verdict, strings.Join(parts, "; "))
}

func FormatMeanStress(x MeanStressAudit) string {
	return fmt.Sprintf("xi=%.15g wHalf=%.15g signed=%.15g rel=%.5g boundarySplit=%.15g existingMeanBetter=%t antiAlignStronger=%t verdict=%q", x.XiBoundary, x.HalfWeight, x.SignedResidual, x.RelativeResidual, x.BoundarySplit, x.ExistingMeanStressBetter, x.AntiAlignmentSealStronger, x.Verdict)
}

func FormatSplit(x TwoCoordinateSplitAudit) string {
	return fmt.Sprintf("full=%.15g half=%.15g signed=(%.15g,%.15g) mean=%.15g pair=%q fullTyped=%t perCoord=%t signedTyped=%t meanTyped=%t map=%t trace=%t verdict=%q", x.FullWeight, x.HalfWeight, x.SignedPair[0], x.SignedPair[1], x.MeanStressCandidate, x.BoundaryPair, x.FullWeightTyped, x.PerCoordinateTyped, x.SignedPairTyped, x.MeanStressTyped, x.SuppliesBoundaryMap, x.SuppliesTraceTheorem, x.Verdict)
}

func FormatRelations(x PreviousSealRelationAudit) string {
	parts := make([]string, 0, len(x.Rows))
	for _, r := range x.Rows {
		parts = append(parts, fmt.Sprintf("%s candidate=%q lawful=%t typedArithmetic=%t class=%q", r.Seal, r.Candidate, r.LawfulRelation, r.TypedArithmetic, r.Classification))
	}
	return fmt.Sprintf("history=%t boundary=%t orientation=%t fano=%t verdict=%q rows=%s", x.HistoryLoopSource, x.BoundaryStressSource, x.OrientationBalanceSource, x.FanoHitchinSource, x.Verdict, strings.Join(parts, "; "))
}

func FormatBoundaryMap(x BoundaryMapObstructionAudit) string {
	return fmt.Sprintf("missingMap=%q missingTrace=%q halfMap=%t sevenMap=%t stressMap=%t deriveStress=%t deriveLambdaR3=%t verdict=%q", x.MissingMap, x.MissingTraceTheorem, x.HasHalfTraceMap, x.HasSevenOver72Map, x.HasBoundaryStressMap, x.CanDeriveBoundaryStress, x.CanDeriveLambdaOrR3, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("stress=%t lambdaR3=%t seven144=%t seven72=%t history=%t scalarFlavor=%t physical=%t higgs=%t ckmPmns=%t gauge=%t verdict=%q", x.ClaimsBoundaryStress, x.ClaimsLambdaR3, x.ClaimsSevenOver144, x.ClaimsSevenOver72, x.ClaimsHistoryLoopUnit, x.ClaimsScalarFlavor, x.ClaimsPhysicalMetric, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.Verdict)
}
