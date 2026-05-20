package generation2hitchinnegativesectormultiplicitytraceidentityaudit

import (
	"fmt"
	"math"
	"strings"
)

func f64(x float64) string {
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return fmt.Sprintf("%g", x)
	}
	return fmt.Sprintf("%.15g", x)
}

func FormatInherited(x Gate645Inheritance) string {
	return fmt.Sprintf("negativeWeight=%t angle=%t components=%t routes=%d minusCandidate=%t symbolic=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physical=%t firewall=%t verdict=%q", x.NegativeWeightCertified, x.ProjectiveAngleDerived, x.ComponentAuditComputed, x.RouteCount, x.MinusThreeSourceCandidate, x.FullSymbolicTheoremCertified, x.SplitG2Certified, x.BoundaryStressAssignment, x.SevenOver72Theorem, x.ScalarFlavorTransport, x.PhysicalMetric, x.Gate645FirewallPreserved, x.Verdict)
}

func FormatComponentFamily(x ComponentFamilyContribution) string {
	return fmt.Sprintf("%s minus=%d hint=%q omega0=%s omega1Alt=%s omega2Alt=%s survives=%t blockContribution=%t symbolic=%t interpretation=%q", x.Family, x.MinusCount, x.DimensionHint, f64(x.Omega0NormSq), f64(x.Omega1AltNormSq), f64(x.Omega2AltNormSq), x.Survives, x.BlockContributionCertified, x.SymbolicContributionCertified, x.Interpretation)
}

func FormatComponents(x ComponentFamilyContributionAudit) string {
	parts := []string{fmt.Sprintf("families=%d all=%t anyContribution=%t symbolic=%t verdict=%q", len(x.Families), x.AllFamiliesAudited, x.AnyFamilyContributionCertified, x.SymbolicComponentTheoremCertified, x.Verdict)}
	for _, f := range x.Families {
		parts = append(parts, FormatComponentFamily(f))
	}
	return strings.Join(parts, "\n")
}

func FormatOffBlock(x OffBlockCancellationAudit) string {
	return fmt.Sprintf("maxOffNorm=%s numericCancel=%t structuralSource=%q symbolic=%t verdict=%q", f64(x.MaxOffBlockFrobeniusNorm), x.NumericalCancellation, x.StructuralCancellationSource, x.SymbolicCancellationCertified, x.Verdict)
}

func FormatPositive(x PositiveSectorUnitWeightAudit) string {
	return fmt.Sprintf("p=%d q=%d positiveWeight=%s maxSpread=%s unitCertified=%t symbolic=%t explanation=%q verdict=%q", x.PositiveDim, x.NegativeDim, f64(x.ObservedPositiveWeight), f64(x.MaxPositiveBlockSpread), x.UnitWeightCertified, x.SymbolicUnitWeightCertified, x.CandidateExplanation, x.Verdict)
}

func FormatNegative(x NegativeSectorMultiplicityAudit) string {
	return fmt.Sprintf("p=%d q=%d negativeWeight=%s ratio=%s maxDrift=%s certified=%t symbolic=%t formula=%q source=%q verdict=%q", x.PositiveDim, x.NegativeDim, f64(x.ObservedNegativeWeight), f64(x.ObservedMinusToPlusRatio), f64(x.MaxRatioDrift), x.MultiplicityWeightCertified, x.SymbolicMultiplicityCertified, x.CandidateFormula, x.CandidateSource, x.Verdict)
}

func FormatProjectorIdentity(x ProjectorPlaneIdentityAudit) string {
	return fmt.Sprintf("p=%d q=%d G=%q B=%q Gnorm2=%s Bnorm2=%s cosFormula=%q cos=%s cos2=%s rhoFormula=%q rho2=%s expCos=%s expRho2=%s matches=%t symbolic=%t verdict=%q", x.PositiveDim, x.NegativeDim, x.GHatFormula, x.BHatFormula, f64(x.GHatNormalizerSq), f64(x.BHatNormalizerSq), x.CosineFormula, f64(x.Cosine), f64(x.CosineSquared), x.ResidualSquaredFormula, f64(x.ResidualSquared), f64(x.ExpectedCosine), f64(x.ExpectedResidualSquared), x.IdentityMatchesRouteData, x.FullSymbolicTheoremCertified, x.Verdict)
}

func FormatRoute(x RouteTraceIdentityRow) string {
	return fmt.Sprintf("%s inertia=%s plus=%s minus=%s ratio=%s plusSpread=%s minusSpread=%s off=%s block=%t pq=%t", x.Name, x.Inertia, f64(x.PlusMean), f64(x.MinusMean), f64(x.MinusToPlusRatio), f64(x.PlusBlockSpread), f64(x.MinusBlockSpread), f64(x.OffBlockFrobeniusNorm), x.BlockFormCertified, x.MatchesPQIdentity)
}

func FormatRouteUniversality(x RouteUniversalityAudit) string {
	parts := []string{fmt.Sprintf("routes=%d all=%t universalCandidate=%t routeFailure=%t verdict=%q", len(x.Routes), x.AllRoutesPass, x.RouteUniversalCandidate, x.RouteDependentFailure, x.Verdict)}
	for _, r := range x.Routes {
		parts = append(parts, FormatRoute(r))
	}
	return strings.Join(parts, "\n")
}

func FormatInterpretation(x Interpretation) string {
	return fmt.Sprintf("gate645=%t components=%t pq=%t routeUniversal=%t symbolic=%t interpretation=%q verdict=%q", x.InheritedGate645, x.BlockComponentAuditComputed, x.PQIdentityMatches, x.RouteUniversal, x.SymbolicTheoremCertified, x.Interpretation, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("symbolic=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physicalMetric=%t flavor=%t higgs=%t ckm_pmns=%t gauge=%t verdict=%q", x.ClaimsFullSymbolicHitchinTheorem, x.ClaimsSplitG2, x.ClaimsBoundaryStress, x.ClaimsSevenOver72Theorem, x.ClaimsScalarFlavor, x.ClaimsPhysicalMetric, x.ClaimsFlavor, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.Verdict)
}
