package generation2k7overlambda4boundarypairprojectionaudit

import (
	"fmt"
	"math"
	"strings"
)

func f64(x float64) string {
	if math.IsNaN(x) {
		return "symbolic"
	}
	if math.IsInf(x, 1) {
		return "+Inf"
	}
	if math.IsInf(x, -1) {
		return "-Inf"
	}
	return fmt.Sprintf("%.15g", x)
}

func FormatInherited(i Gate627Inheritance) string {
	return fmt.Sprintf("K7=%d lambda4=%d denom=%d weight=%s complementNum=%d numNative=%t certified72=%t projection=%t sourceTheorem=%t firewall=%t mixture=%s weightedResidual=%s scalarResidual=%s absLambda12=%s R3MinusOne=%s split=%s xi=%s verdict=%q", i.K7Dimension, i.Lambda4Dimension, i.Gate627Denominator, f64(i.Gate627Weight), i.Gate627ComplementNumerator, i.Gate627NumeratorIsNative, i.Gate627Certified72Carrier, i.Gate627ProjectionExists, i.Gate627SourceTheorem, i.Gate627FirewallPreserved, f64(i.Gate626WeightedMixture), f64(i.Gate626WeightedResidual), f64(i.Gate626ScalarResidual), f64(i.AbsLambda12), f64(i.R3MinusOne), f64(i.BoundarySplit), f64(i.XiBoundary), i.Verdict)
}

func FormatChamber(c ChamberDimensionAudit) string {
	return fmt.Sprintf("lambda4=%q:%d boundaryPair=%q:%d expression=%q chamber=%d target=%d equals=%t nativeLambda4=%t bridgePair=%t nativeDirectSum=%t betterThan8x9=%t interpretation=%q verdict=%q", c.Lambda4Carrier, c.Lambda4Dimension, c.BoundaryPairCarrier, c.BoundaryPairDimension, c.AugmentedChamberExpression, c.AugmentedChamberDimension, c.TargetDenominator, c.EqualsTargetDenominator, c.UsesNativeLambda4Carrier, c.UsesBridgeBoundaryPair, c.DirectSumCertifiedNative, c.BetterThan8Times9, c.Interpretation, c.Verdict)
}

func FormatDenominatorComparisonRow(r DenominatorComparisonRow) string {
	return fmt.Sprintf("name=%q expression=%q value=%d nativeFinite=%t envPair=%t quarantined=%t certifiedNative=%t rank=%d comment=%q", r.Name, r.Expression, r.Value, r.NativeFiniteCarrier, r.UsesEnvironmentalPair, r.UsesQuarantinedLedger, r.CertifiedAsNative, r.StrengthRank, r.Comment)
}

func FormatDenominatorComparisonRows(rows []DenominatorComparisonRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatDenominatorComparisonRow(r))
	}
	return strings.Join(parts, " | ")
}

func FormatDenominatorComparison(d DenominatorComparisonAudit) string {
	return fmt.Sprintf("best=%q bestValue=%d best70plus2=%t nativeDenom=%t bridgeCandidate=%t rows=[%s] verdict=%q", d.BestExpression, d.BestValue, d.BestIs70Plus2, d.AnyNativeDenominator, d.AnyBridgeCandidate, FormatDenominatorComparisonRows(d.Rows), d.Verdict)
}

func FormatBoundaryPair(b BoundaryStressPairAudit) string {
	return fmt.Sprintf("pair=[%s] absLambda12=%s R3MinusOne=%s xi=%s split=%s dim=%d gate613=%t from626=%t nativeFinite=%t bridgeOnly=%t interpretation=%q verdict=%q", strings.Join(b.PairCoordinates, "; "), f64(b.AbsLambda12), f64(b.R3MinusOne), f64(b.XiBoundary), f64(b.BoundarySplit), b.PairDimension, b.PairIsGate613Boundary, b.PairInheritedFromGate626, b.PairNativeFiniteObject, b.BridgeCoordinateOnly, b.Interpretation, b.Verdict)
}

func FormatK7Embedding(k K7Lambda4EmbeddingAudit) string {
	return fmt.Sprintf("K7=%q:%d lambda4=%q:%d fits=%t rankPB=%d rankPG=%d native=%t projection=%t interpretation=%q verdict=%q", k.K7Carrier, k.K7Dimension, k.Lambda4Carrier, k.Lambda4Dimension, k.K7FitsInsideLambda4, k.RankPB, k.RankPG, k.NativeCarrierCertified, k.ProjectionToBoundaryFound, k.Interpretation, k.Verdict)
}

func FormatComplement(c ComplementChamberAudit) string {
	return fmt.Sprintf("lambda4=%d K7=%d nonK7=%d boundaryPair=%d augComplement=%d chamber=%d weight=%s equals65over72=%t structured=%t nativeProjection=%t equation=%q verdict=%q", c.Lambda4Dimension, c.K7Dimension, c.NonK7Lambda4ComplementDimension, c.BoundaryPairDimension, c.AugmentedComplementDimension, c.AugmentedChamberDimension, f64(c.ComplementWeight), c.Equals65Over72, c.HasStructuredComplementReading, c.NativeComplementProjection, c.Equation, c.Verdict)
}

func FormatProjectionTrace(p ProjectionTraceAudit) string {
	return fmt.Sprintf("domain=%q dim=%d K7=%q:%d pullLine=%q fraction=%s expected=%s matches=%t operator=%t idempotent=%t traceFn=%t intertwiner=%t missing=%q verdict=%q", p.DomainChamber, p.DomainDimension, p.K7TraceCarrier, p.K7TraceDimension, p.BoundaryPullLine, f64(p.TraceFraction), f64(p.ExpectedWeight), p.TraceFractionMatches, p.ProjectionOperatorExists, p.IdempotentCertified, p.TraceFunctionalCertified, p.IntertwinerCertified, p.MissingObject, p.Verdict)
}

func FormatWeightedClosure(w WeightedClosureCarryAudit) string {
	return fmt.Sprintf("kappaSum=%s mixture=%s residual=%s boundaryWeight=%s scalarWeight=%s ratio=%s ratioMatchesGate626=%t equation=%q verdict=%q", f64(w.KappaSum), f64(w.WeightedMixture), f64(w.WeightedClosureResidual), f64(w.BoundaryWeight), f64(w.ScalarWeight), f64(w.WeightFromChamberRatio), w.ChamberRatioMatchesGate626, w.MixtureEquation, w.Verdict)
}

func FormatNativeStatus(n NativeASHAStatus) string {
	return fmt.Sprintf("lambda4Native=%t K7Native=%t boundaryPairNative=%t chamberNative=%t airlockNative=%t projectorNative=%t traceTheoremNative=%t transportNative=%t statement=%q verdict=%q", n.Lambda4Native, n.K7Native, n.BoundaryPairNativeFinite, n.AugmentedChamberNative, n.ProductAirlockNative, n.K7BoundaryPullProjectorNative, n.TraceFractionTheoremNative, n.GaugeScalarFlavorTransportNative, n.Statement, n.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("nativeChamber=%t nativePair=%t nativeProjection=%t nativeTrace=%t scalarRG=%t flavor=%t gaugeUnification=%t higgsMass=%t endpoint=%t verdict=%q", f.ClaimsNativeAugmentedChamber, f.ClaimsNativeBoundaryPair, f.ClaimsNativeProjection, f.ClaimsNativeTraceTheorem, f.ClaimsScalarRGMatching, f.ClaimsFlavorOrientation, f.ClaimsGaugeUnification, f.ClaimsHiggsMassDerived, f.ClaimsEndpointDerivation, f.Verdict)
}
