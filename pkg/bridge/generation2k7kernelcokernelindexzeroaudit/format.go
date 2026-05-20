package generation2k7kernelcokernelindexzeroaudit

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

func FormatInherited(i Gate629Inheritance) string {
	return fmt.Sprintf("U=%d V=%d directSum=%d lambda4=%d intersection=%d span=%d cokernel=%d boundaryPair=%d chamber=%d boundaryWeight=%s residual=%s dualCandidate=%t isoMissing=%t boundaryAssignmentMissing=%t firewall=%t verdict=%q", i.UDimension, i.VDimension, i.DirectSumDimension, i.Lambda4Dimension, i.IntersectionDimension, i.SpanDimension, i.CokernelDimension, i.BoundaryPairDimension, i.AugmentedChamberDimension, f64(i.BoundaryWeight), f64(i.WeightedClosureResidual), i.Gate629DualCandidate, i.Gate629IsomorphismMissing, i.Gate629BoundaryAssignmentMissing, i.Gate629FirewallPreserved, i.Verdict)
}

func FormatAdditionMap(a AdditionMapAudit) string {
	return fmt.Sprintf("map=%q domain=%q codomain=%q formula=%q domainDim=%d codomainDim=%d square=%t kernel=%q kernelDim=%d kernelIsK7=%t image=%q imageDim=%d imageIsSpan=%t cokernel=%q cokernelDim=%d cokernelMatchesK7=%t rankDefect=%d index=%d indexZero=%t verdict=%q", a.MapName, a.Domain, a.Codomain, a.Formula, a.DomainDimension, a.CodomainDimension, a.SquareOperator, a.KernelExpression, a.KernelDimension, a.KernelIsK7, a.ImageExpression, a.ImageDimension, a.ImageIsSpan, a.CokernelExpression, a.CokernelDimension, a.CokernelMatchesK7, a.RankDefect, a.Index, a.IndexZero, a.Verdict)
}

func FormatDefect(d KernelCokernelDefectAudit) string {
	return fmt.Sprintf("kernel=%q:%d cokernel=%q:%d balanced=%t index=%d fredholmAnalogyOnly=%t candidatePair=%t missing=%q interpretation=%q verdict=%q", d.KernelCarrier, d.KernelDimension, d.CokernelCarrier, d.CokernelDimension, d.DefectsBalanced, d.Index, d.FredholmAnalogyOnly, d.CandidateDefectPair, d.MissingPairing, d.Interpretation, d.Verdict)
}

func FormatBlockCompression(b K7BlockCompressionAudit) string {
	return fmt.Sprintf("K7block=%d PBblocks=%d PGblocks=%d spanBlocks=%d lambda4Blocks=%d boundaryCoords=%d augmented=%q weightExpr=%q weight=%s exact=%t defectCandidate=%t interpretation=%q verdict=%q", b.K7BlockDimension, b.PBBlocks, b.PGBlocks, b.SpanBlocks, b.Lambda4Blocks, b.BoundaryCoordinates, b.AugmentedExpression, b.BoundaryWeightExpression, f64(b.BoundaryWeight), b.CompressionExact, b.DefectBlockCandidate, b.Interpretation, b.Verdict)
}

func FormatPairingCandidate(c PairingCandidate) string {
	return fmt.Sprintf("name=%q lane=%q couldTouch=%t certified=%t failure=%q", c.Name, c.SourceLane, c.CouldTouchDefects, c.Certified, c.FailureReason)
}

func FormatPairingCandidates(rows []PairingCandidate) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatPairingCandidate(r))
	}
	return strings.Join(parts, " | ")
}

func FormatPairing(p PairingCandidateAudit) string {
	return fmt.Sprintf("candidates=[%s] canonical=%t metric=%t hodge=%t eta=%t projector=%t missing=%q verdict=%q", FormatPairingCandidates(p.Candidates), p.CanonicalPairingFound, p.MetricPairingCertified, p.HodgeStarPairingCertified, p.EtaPairingCertified, p.ProjectorPairingCertified, p.MissingObject, p.Verdict)
}

func FormatBoundaryAssignment(b BoundaryStressAssignmentAudit) string {
	return fmt.Sprintf("defectCanSupplySeven=%t boundaryPair=%d weight=%s stressLine=%q assignment=%t transport=%t missing=%q verdict=%q", b.DefectBlockCanSupplySeven, b.BoundaryPairDimension, f64(b.BoundaryWeight), b.BoundaryStressLine, b.AssignmentCertified, b.NativeTransportTheorem, b.MissingObject, b.Verdict)
}

func FormatNativeStatus(n NativeASHAStatus) string {
	return fmt.Sprintf("lambda4Native=%t U=%t V=%t K7=%t mapTyped=%t kernelTyped=%t cokernelTyped=%t indexZeroTyped=%t pairing=%t boundaryAssignment=%t traceTheorem=%t statement=%q verdict=%q", n.Lambda4Native, n.UImageRankNative, n.VImageRankNative, n.K7IntersectionNative, n.AdditionMapTyped, n.KernelDimensionTyped, n.CokernelDimensionTyped, n.IndexZeroTyped, n.CanonicalKernelCokernelPairing, n.BoundaryStressAssignmentNative, n.K7DefectBoundaryTraceTheorem, n.Statement, n.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("canonicalPairing=%t boundaryAssignment=%t defectTrace=%t boundaryPairNative=%t scalarRG=%t flavor=%t gaugeUnification=%t higgsMass=%t endpoint=%t verdict=%q", f.ClaimsCanonicalPairing, f.ClaimsBoundaryStressAssignment, f.ClaimsK7DefectTraceTheorem, f.ClaimsBoundaryPairNative, f.ClaimsScalarRGMatching, f.ClaimsFlavorOrientation, f.ClaimsGaugeUnification, f.ClaimsHiggsMassDerived, f.ClaimsEndpointDerivation, f.Verdict)
}
