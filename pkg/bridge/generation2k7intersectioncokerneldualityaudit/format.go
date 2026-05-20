package generation2k7intersectioncokerneldualityaudit

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

func FormatInherited(i Gate628Inheritance) string {
	return fmt.Sprintf("lambda4=%d boundaryPair=%d chamber=%d K7=%d nonK7=%d augmentedComplement=%d weight=%d/%d boundaryWeight=%s scalarWeight=%s weightedResidual=%s mixture=%s absLambda12=%s R3MinusOne=%s split=%s chamberCandidate=%t projectionMissing=%t airlockMissing=%t firewall=%t verdict=%q", i.Lambda4Dimension, i.BoundaryPairDimension, i.AugmentedChamberDimension, i.K7Dimension, i.NonK7Lambda4Complement, i.AugmentedComplementDimension, i.WeightNumerator, i.WeightDenominator, f64(i.BoundaryWeight), f64(i.ScalarWeight), f64(i.WeightedClosureResidual), f64(i.WeightedMixture), f64(i.AbsLambda12), f64(i.R3MinusOne), f64(i.BoundarySplit), i.Gate628ChamberCandidate, i.Gate628ProjectionMissing, i.Gate628ProductAirlockMissing, i.Gate628FirewallPreserved, i.Verdict)
}

func FormatSpan(s BooleanOctonionicSpanAudit) string {
	return fmt.Sprintf("U=%q V=%q intersection=%q rankPB=%d rankPG=%d dimIntersection=%d span=%d expectedSpan=%d formula=%q spanMatches=%t lambda4=%d cokernel=%d cokernelFormula=%q cokernelMatchesK7=%t rankCertified=%t verdict=%q", s.UCarrier, s.VCarrier, s.IntersectionCarrier, s.RankPB, s.RankPG, s.IntersectionDimension, s.SpanDimension, s.ExpectedSpanDimension, s.SpanFormula, s.SpanMatchesExpected, s.Lambda4Dimension, s.CokernelDimension, s.CokernelFormula, s.CokernelMatchesK7Dimension, s.SpanDimensionCertifiedByRank, s.Verdict)
}

func FormatDuality(d IntersectionCokernelDualityAudit) string {
	return fmt.Sprintf("intersection=%q:%d cokernel=%q:%d dimensionsEqual=%t dimensionalOnly=%t isoFound=%t pairingFound=%t candidate=%t missing=%q interpretation=%q verdict=%q", d.IntersectionCarrier, d.IntersectionDimension, d.CokernelCarrier, d.CokernelDimension, d.DimensionsEqual, d.EqualityIsOnlyDimensional, d.CanonicalIsomorphismFound, d.CanonicalPairingFound, d.DualityCandidate, d.MissingMap, d.Interpretation, d.Verdict)
}

func FormatChamberSplit(c ChamberSplitAudit) string {
	return fmt.Sprintf("intersectionOrGap=%d span=%d boundaryPair=%d chamber=%d split=%q matches72=%t nativeSpan=%t bridgePair=%t sharperThan70Plus2=%t interpretation=%q verdict=%q", c.IntersectionOrGapDimension, c.SpanDimension, c.BoundaryPairDimension, c.AugmentedChamberDimension, c.SplitExpression, c.SplitMatches72, c.NativeSpanDimension, c.BoundaryPairBridgeOnly, c.SharperThan70Plus2, c.Interpretation, c.Verdict)
}

func FormatComplementRole(c ComplementRoleAudit) string {
	return fmt.Sprintf("span=%d boundaryPair=%d spanBoundary=%d chamber=%d weight=%s equals65over72=%t previous=%q sharpened=%q role=%q nativeRoleTheorem=%t verdict=%q", c.SpanDimension, c.BoundaryPairDimension, c.SpanBoundaryComplement, c.AugmentedChamberDimension, f64(c.ComplementWeight), c.Equals65Over72, c.PreviousComplementEquation, c.SharpenedComplementEquation, c.RoleReading, c.NativeRoleTheoremFound, c.Verdict)
}

func FormatBoundaryPullCandidate(c BoundaryPullCandidate) string {
	return fmt.Sprintf("name=%q dim=%d source=%q canSupplySeven=%t assigned=%t nativeTheorem=%t comment=%q", c.Name, c.Dimension, c.SourceType, c.CanSupplySeven, c.BoundaryAssignment, c.NativeTheorem, c.Comment)
}

func FormatBoundaryPullCandidates(rows []BoundaryPullCandidate) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatBoundaryPullCandidate(r))
	}
	return strings.Join(parts, " | ")
}

func FormatBoundaryPullAssignment(b BoundaryPullAssignmentAudit) string {
	return fmt.Sprintf("candidates=[%s] line=%q weight=%s intersectionAssigned=%t cokernelAssigned=%t dualAssigned=%t certified=%t missing=%q verdict=%q", FormatBoundaryPullCandidates(b.Candidates), b.BoundaryPullLine, f64(b.BoundaryWeight), b.IntersectionAssigned, b.CokernelAssigned, b.DualPairAssigned, b.AssignmentCertified, b.MissingObject, b.Verdict)
}

func FormatWeightedMixture(w WeightedMixtureReinterpretationAudit) string {
	return fmt.Sprintf("kappaSum=%s weightedMixture=%s residual=%s boundaryWeight=%s scalarWeight=%s scalarWeightAs63Plus2=%t boundaryWeightAsSeven=%t equation=%q interpretation=%q verdict=%q", f64(w.KappaSum), f64(w.WeightedMixture), f64(w.Residual), f64(w.BoundaryWeight), f64(w.ScalarWeight), w.ScalarWeightAs63Plus2, w.BoundaryWeightAsSeven, w.MixtureEquation, w.Interpretation, w.Verdict)
}

func FormatNativeStatus(n NativeASHAStatus) string {
	return fmt.Sprintf("lambda4Native=%t PBnative=%t PGnative=%t K7native=%t spanTyped=%t cokernelTyped=%t iso=%t boundaryPairNative=%t assignmentNative=%t projectorNative=%t transportNative=%t statement=%q verdict=%q", n.Lambda4Native, n.PBImageRankNative, n.PGImageRankNative, n.K7IntersectionNative, n.BooleanOctonionicSpanDimensionTyped, n.Lambda4CokernelDimensionTyped, n.IntersectionCokernelIsomorphism, n.BoundaryPairNativeFinite, n.BoundaryPullAssignmentNative, n.DualBoundaryProjectorNative, n.GaugeScalarFlavorTransportNative, n.Statement, n.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("k7CokernelIso=%t boundaryAssignment=%t dualProjector=%t boundaryPairNative=%t scalarRG=%t flavor=%t gaugeUnification=%t higgsMass=%t endpoint=%t verdict=%q", f.ClaimsK7CokernelIsomorphism, f.ClaimsBoundaryPullAssignment, f.ClaimsDualBoundaryProjector, f.ClaimsBoundaryPairNative, f.ClaimsScalarRGMatching, f.ClaimsFlavorOrientation, f.ClaimsGaugeUnification, f.ClaimsHiggsMassDerived, f.ClaimsEndpointDerivation, f.Verdict)
}
