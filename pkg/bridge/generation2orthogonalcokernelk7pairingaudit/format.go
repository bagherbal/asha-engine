package generation2orthogonalcokernelk7pairingaudit

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

func FormatInherited(i Gate630Inheritance) string {
	return fmt.Sprintf("H=%d U=%d V=%d directSum=%d K7=%d span=%d cokernel=%d index=%d boundaryPair=%d chamber=%d weight=%s indexZero=%t pairingMissing=%t boundaryMissing=%t firewall=%t verdict=%q", i.HDimension, i.UDimension, i.VDimension, i.DirectSumDimension, i.K7Dimension, i.SpanDimension, i.CokernelDimension, i.Index, i.BoundaryPairDimension, i.AugmentedChamberDimension, f64(i.BoundaryWeight), i.Gate630IndexZero, i.Gate630PairingMissing, i.Gate630BoundaryMissing, i.Gate630FirewallPreserved, i.Verdict)
}

func FormatOrthogonalW7(w OrthogonalCokernelRepresentativeTable) string {
	return fmt.Sprintf("ambient=%q:%d span=%q:%d Pspan=%q PW=%q W=%q def=%q Wdim=%d orthU=%t orthV=%t directSum=%q certified=%t quotient=%q represents=%t metricDependent=%t nativeComplement=%t verdict=%q", w.AmbientSpace, w.AmbientDimension, w.SpanSpace, w.SpanDimension, w.ProjectorOntoSpan, w.ProjectorOntoW, w.WName, w.WDefinition, w.WDimension, w.WOrthogonalToU, w.WOrthogonalToV, w.DirectSumDecomposition, w.DirectSumCertified, w.QuotientRepresentative, w.RepresentsCokernel, w.MetricDependent, w.NativeComplementCertified, w.Verdict)
}

func FormatExactSequence(e ExactDefectSequence) string {
	return fmt.Sprintf("sequence=%q injection=%q A=%q projection=%q exactK=%t exactDirectSum=%t exactH=%t exactW=%t altSum=%d rankNullity=%t verdict=%q", e.Sequence, e.KernelInjection, e.AdditionMap, e.ProjectionMap, e.ExactAtK7, e.ExactAtDirectSum, e.ExactAtH, e.ExactAtW7, e.DimensionAlternatingSum, e.ExactByRankNullity, e.Verdict)
}

func FormatPairingCandidate(c PairingCandidate) string {
	return fmt.Sprintf("name=%q formula=%q lane=%q touches=%t rankTest=%t nondegenerate=%t canonical=%t condition=%q verdict=%q", c.Name, c.Formula, c.SourceLane, c.TouchesK7AndW7, c.RankTestAvailable, c.NondegenerateCertified, c.CanonicalCertified, c.FailureOrCondition, c.Verdict)
}

func FormatPairingCandidates(rows []PairingCandidate) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatPairingCandidate(r))
	}
	return strings.Join(parts, " | ")
}

func FormatCandidatePairingTable(c CandidatePairingTable) string {
	return fmt.Sprintf("rows=[%s] canonical=%t nondegenerate=%t sharpened=%t missing=%q verdict=%q", FormatPairingCandidates(c.Candidates), c.CanonicalPairingFound, c.NondegeneratePairingFound, c.PairingProblemSharpened, c.MissingObject, c.Verdict)
}

func FormatHodgeStar(h HodgeStarPairingAudit) string {
	return fmt.Sprintf("formula=%q typed=%t mapsLambda4=%t orientationChoice=%t rankTest=%t nondegenerate=%t uvw=%q condition=%q verdict=%q", h.Formula, h.HodgeStarTypedOnLambda4, h.MapsLambda4ToLambda4, h.RequiresOrientationChoice, h.RankTestImplemented, h.NondegenerateCertified, h.PreservesOrExchangesUVW, h.Condition, h.Verdict)
}

func FormatProjectorCandidate(c ProjectorAlgebraCandidate) string {
	return fmt.Sprintf("operator=%q action=%q afterPW=%q pairs=%t", c.Operator, c.ActionOnK7, c.AfterPW, c.PairsK7ToW7)
}

func FormatProjectorCandidates(rows []ProjectorAlgebraCandidate) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatProjectorCandidate(r))
	}
	return strings.Join(parts, " | ")
}

func FormatProjectorAlgebra(p ProjectorAlgebraPairingAudit) string {
	return fmt.Sprintf("rows=[%s] K7fixedPB=%t K7fixedPG=%t PWkillsSpan=%t any=%t reason=%q verdict=%q", FormatProjectorCandidates(p.Rows), p.K7FixedByPB, p.K7FixedByPG, p.PWKillsUPlusV, p.AnyPairingCertified, p.Reason, p.Verdict)
}

func FormatEta(e EtaPairingAudit) string {
	return fmt.Sprintf("formula=%q typedEtaOnLambda4=%t rankTest=%t pairing=%t compatibility=%t reason=%q verdict=%q", e.Formula, e.TypedEtaOnLambda4Available, e.RankTestImplemented, e.PairingCertified, e.CompatibilityCertified, e.Reason, e.Verdict)
}

func FormatDeterminantLine(d DeterminantLineAudit) string {
	return fmt.Sprintf("sequence=%q relation=%q canonicalLine=%t pointwise=%t orientationDependent=%t volumeBookkeeping=%t traceByItself=%t interpretation=%q verdict=%q", d.ExactSequence, d.DeterminantRelation, d.CanonicalLineRelation, d.PointwiseIsomorphism, d.OrientationDependent, d.CanSupportVolumeBookkeeping, d.CanSupportNormalizedTraceByItself, d.Interpretation, d.Verdict)
}

func FormatBoundaryReadiness(b BoundaryReadinessAudit) string {
	return fmt.Sprintf("K7toW7=%t detLine=%t boundaryPair=%d needsW7Boundary=%t needsDefectTrace=%t assignment=%t missing=%q verdict=%q", b.K7ToW7PairingCertified, b.DeterminantLineRelationAvailable, b.BoundaryPairDimension, b.StillRequiresW7ToBoundary, b.StillRequiresDefectTraceToBoundary, b.BoundaryAssignmentCertified, b.MissingObject, b.Verdict)
}

func FormatNativeStatus(n NativeASHAStatus) string {
	return fmt.Sprintf("lambda4=%t metric=%t W7rep=%t K7=%t W7dim=%t exact=%t hodgeRank=%t projector=%t eta=%t detLine=%t canonicalK7W7=%t boundaryNative=%t statement=%q verdict=%q", n.Lambda4Native, n.AmbientMetricAdmitted, n.OrthogonalRepresentativeTyped, n.K7Native, n.W7DimensionTyped, n.ExactDefectSequenceTyped, n.HodgeStarRankCertified, n.ProjectorPairingCertified, n.EtaPairingCertified, n.DeterminantLineRelationTyped, n.CanonicalK7ToW7Pairing, n.BoundaryStressAssignmentNative, n.Statement, n.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("canonicalK7W7=%t boundaryAssignment=%t scalarRG=%t higgs=%t flavor=%t ckmPmns=%t gauge=%t boundaryPairNative=%t traceTheorem=%t verdict=%q", f.ClaimsCanonicalK7W7Pairing, f.ClaimsBoundaryStressAssignment, f.ClaimsScalarRGMatching, f.ClaimsHiggsMassDerivation, f.ClaimsFlavorDerivation, f.ClaimsCKMPMNSDerivation, f.ClaimsGaugeUnification, f.ClaimsBoundaryPairNative, f.ClaimsNativeTraceTheorem, f.Verdict)
}
