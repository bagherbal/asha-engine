package generation2hitchincubicsectorcontractionmultiplicityaudit

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

func FormatInherited(x Gate646Inheritance) string {
	return fmt.Sprintf("projector=%t universal=%t p=%d q=%d G=%q B=%q cos=%s rho2=%s symbolic=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physical=%t firewall=%t verdict=%q", x.ProjectorPlaneIdentityInherited, x.RouteUniversal, x.PositiveDim, x.NegativeDim, x.GHatFormula, x.BHatFormula, f64(x.Cosine), f64(x.ResidualSquared), x.FullSymbolicTheoremCertified, x.SplitG2Certified, x.BoundaryStressAssignment, x.SevenOver72Theorem, x.ScalarFlavorTransport, x.PhysicalMetric, x.Gate646FirewallPreserved, x.Verdict)
}

func FormatFamily(x FamilyLedgerRow) string {
	return fmt.Sprintf("%s minus=%d parity=%+d hint=%q omega1=%s omega2=%s omegaB=%s plus=%t minusBlock=%t mixed=%t symbolic=%t comment=%q", x.Family, x.MinusCount, x.ParityUnderSK, x.DimensionHint, f64(x.Omega1AltNormSq), f64(x.Omega2AltNormSq), f64(x.OmegaBAltNormSq), x.ContributesToPlus, x.ContributesToMinus, x.ContributesToMixed, x.SymbolicSourceFound, x.Comment)
}

func FormatFamilies(x ComponentFamilyLedger) string {
	parts := []string{fmt.Sprintf("families=%d all=%t antisymTwists=%t symbolic=%t verdict=%q", len(x.Families), x.AllFamiliesAudited, x.AntisymmetrizedTwistsAudited, x.AnySymbolicFamilyTheorem, x.Verdict)}
	for _, f := range x.Families {
		parts = append(parts, FormatFamily(f))
	}
	return strings.Join(parts, "\n")
}

func FormatTriple(x TripleContribution) string {
	return fmt.Sprintf("%s %s totalMinus=%d plusTrace=%s minusTrace=%s mixed=%s plusUnit=%s minusUnit=%s significant=%t", x.RouteName, x.Families, x.TotalMinusSlots, f64(x.PlusTrace), f64(x.MinusTrace), f64(x.MixedFrobenius), f64(x.PlusMeanUnit), f64(x.MinusMeanUnit), x.Significant)
}

func FormatRouteContribution(x RouteContributionLedger) string {
	parts := []string{fmt.Sprintf("%s formula=%q signAligned=%t rawPlus=%s rawMinus=%s ratio=%s normPlus=%s normMinus=%s off=%s recon=%s significant=%d/%d block=%t minusQ=%t offCancel=%t symbolic=%t comment=%q", x.RouteName, x.Formula, x.SignAligned, f64(x.RawPositiveMean), f64(x.RawNegativeMean), f64(x.RawMinusToPlusRatio), f64(x.NormalizedPositiveWeight), f64(x.NormalizedNegativeWeight), f64(x.OffBlockFrobeniusNorm), f64(x.AdditiveReconstructionError), x.SignificantTripleCount, x.TotalTripleCount, x.BlockRayCertified, x.MinusQCertified, x.OffBlockCancelledAtTotal, x.SymbolicContractionCertified, x.Comment)}
	for _, t := range x.TopContributions {
		parts = append(parts, FormatTriple(t))
	}
	return strings.Join(parts, "\n")
}

func FormatContributions(x HitchinBlockContributionDecomposition) string {
	parts := []string{fmt.Sprintf("routes=%d reconstruct=%t block=%t anySymbolic=%t samePlane=%t verdict=%q", len(x.Routes), x.AllRoutesReconstruct, x.AllRoutesBlockRayCertified, x.AnyRouteSymbolicCertified, x.SameProjectorPlaneShadow, x.Verdict)}
	for _, r := range x.Routes {
		parts = append(parts, FormatRouteContribution(r))
	}
	return strings.Join(parts, "\n")
}

func FormatPositive(x PositiveSectorUnitAudit) string {
	return fmt.Sprintf("p=%d q=%d allUnit=%t maxDrift=%s symbolic=%t source=%q verdict=%q", x.PositiveDim, x.NegativeDim, x.AllRoutesUnitPositive, f64(x.MaxPositiveWeightDrift), x.SymbolicUnitTheoremCertified, x.SourceClassification, x.Verdict)
}

func FormatNegative(x NegativeSectorMultiplicityAudit) string {
	return fmt.Sprintf("p=%d q=%d target=%s allMinusQ=%t maxDrift=%s cubicSupported=%t symbolic=%t source=%q verdict=%q", x.PositiveDim, x.NegativeDim, f64(x.TargetRatio), x.AllRoutesMinusQ, f64(x.MaxRatioDrift), x.CubicSectorMultiplicitySupported, x.SymbolicMultiplicityTheoremCertified, x.CandidateSource, x.Verdict)
}

func FormatOffBlock(x OffBlockCancellationSourceAudit) string {
	return fmt.Sprintf("maxTotalOff=%s maxTopOffSum=%s structural=%t mechanism=%q verdict=%q", f64(x.MaxTotalOffBlockNorm), f64(x.MaxSumSignificantOffNorms), x.StructuralCancellationCertified, x.CancellationMechanism, x.Verdict)
}

func FormatRouteUniversality(x RouteUniversalityComparison) string {
	return fmt.Sprintf("routes=%d sameFinalRay=%t componentLedgersEqual=%t mechanism=%q verdict=%q", x.RouteCount, x.AllRoutesSameFinalRay, x.ComponentContributionLedgersEqual, x.InternalMechanismClassification, x.Verdict)
}

func FormatTheoremReadiness(x SymbolicTheoremReadiness) string {
	return fmt.Sprintf("candidate=%q finiteSupports=%t components=%t blocks=%t symbolic=%t missing=%q verdict=%q", x.CandidateTheorem, x.FiniteLedgerSupportsTheorem, x.ComponentLedgerComputed, x.BlockContributionComputed, x.FullSymbolicTheoremCertified, x.MissingProofObject, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("symbolic=%t splitG2=%t boundary=%t sevenOver72=%t scalarFlavor=%t physicalMetric=%t flavor=%t higgs=%t ckm_pmns=%t gauge=%t verdict=%q", x.ClaimsFullSymbolicHitchinTheorem, x.ClaimsSplitG2, x.ClaimsBoundaryStress, x.ClaimsSevenOver72Theorem, x.ClaimsScalarFlavor, x.ClaimsPhysicalMetric, x.ClaimsFlavor, x.ClaimsHiggsMass, x.ClaimsCKMPMNS, x.ClaimsGaugeUnification, x.Verdict)
}
