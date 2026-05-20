package generation2hodgestark7tow7leakagerankaudit

import (
	"fmt"
	"math"
	"strings"
)

func f64(x float64) string {
	if math.IsNaN(x) {
		return "NaN"
	}
	if math.IsInf(x, 1) {
		return "+Inf"
	}
	if math.IsInf(x, -1) {
		return "-Inf"
	}
	return fmt.Sprintf("%.15g", x)
}

func f64s(xs []float64) string {
	parts := make([]string, len(xs))
	for i, x := range xs {
		parts[i] = f64(x)
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func FormatInherited(i Gate631Inheritance) string {
	return fmt.Sprintf("H=%d U=%d V=%d K7=%d span=%d W7=%d indexZero=%t orthRep=%t sharpened=%t projectorFail=%t hodgeRankNeeded=%t boundaryMissing=%t firewall=%t verdict=%q", i.HDimension, i.UDimension, i.VDimension, i.K7Dimension, i.SpanDimension, i.W7Dimension, i.IndexZeroInherited, i.OrthogonalRepresentative, i.PairingProblemSharpened, i.ProjectorAlgebraFailed, i.HodgeRankTestRequired, i.BoundaryAssignmentMissing, i.Gate631FirewallPreserved, i.Verdict)
}

func FormatHodgeStar(h HodgeStarMatrixAudit) string {
	return fmt.Sprintf("basis=%q dim=%d typed=%t maps=%t convention=%q star2Residual=%s trace=%s self=%d anti=%d verdict=%q", h.Basis, h.MatrixDimension, h.TypedOnLambda4R8, h.MapsLambda4ToLambda4, h.OrientationConvention, f64(h.StarSquaredResidual), f64(h.Trace), h.SelfDualDimension, h.AntiSelfDualDimension, h.Verdict)
}

func FormatBasis(b BasisCertificate) string {
	return fmt.Sprintf("QK=%dx%d QW=%dx%d span=%dx%d QKiso=%s QWiso=%s QWQK=%s PBQK-QK=%s PGQK-QK=%s PBQW=%s PGQW=%s WorthUV=%t K7inUV=%t verdict=%q", b.QKRows, b.QKCols, b.QWRows, b.QWCols, b.SpanRows, b.SpanCols, f64(b.QKOrthonormalResidual), f64(b.QWOrthonormalResidual), f64(b.QWQKOrthogonalityResidual), f64(b.PBQKMinusQKResidual), f64(b.PGQKMinusQKResidual), f64(b.PBQWResidual), f64(b.PGQWResidual), b.QWOrthogonalToUAndV, b.K7ContainedInUAndV, b.Verdict)
}

func FormatLeakage(l LeakageRankTable) string {
	return fmt.Sprintf("formula=%q size=%dx%d rank=%d singulars=%s det=%s fro=%s minSV=%s cond=%s class=%q verdict=%q", l.Formula, l.Rows, l.Cols, l.Rank, f64s(l.SingularValues), f64(l.Determinant), f64(l.FrobeniusNorm), f64(l.MinimumSingularValue), f64(l.ConditionNumber), l.Classification, l.Verdict)
}

func FormatImageContainment(i ImageContainmentAudit) string {
	return fmt.Sprintf("||*K7||=%s ||PW*K7||=%s ||PUV*K7||=%s leakageRatio=%s spanRatio=%s contained=%t transverse=%t verdict=%q", f64(i.StarK7FrobeniusNorm), f64(i.PWStarK7FrobeniusNorm), f64(i.PUVStarK7FrobeniusNorm), f64(i.LeakageRatio), f64(i.SpanContainmentRatio), i.StarK7ContainedInUPlusV, i.TransverseComponentDetected, i.Verdict)
}

func FormatPairingMetric(p PairingMetricAudit) string {
	return fmt.Sprintf("formula=%q computed=%t rankFull=%t trace=%s scale=%s propIResidual=%s conformal=%t anisotropic=%t degenerate=%t verdict=%q", p.Formula, p.Computed, p.RankFull, f64(p.Trace), f64(p.ScaleCandidate), f64(p.ProportionalIdentityResidual), p.ConformalOrIsometric, p.AnisotropicNondegenerate, p.Degenerate, p.Verdict)
}

func FormatOrientation(o OrientationAudit) string {
	return fmt.Sprintf("det=%s nonzero=%t sign=%d basisDependent=%t physical=%t verdict=%q", f64(o.Determinant), o.NonzeroDeterminant, o.Sign, o.BasisOrientationDependent, o.PhysicalOrientationCertified, o.Verdict)
}

func FormatAlternativeComposites(a AlternativeCompositeAudit) string {
	rows := make([]string, len(a.Rows))
	for i, r := range a.Rows {
		rows[i] = fmt.Sprintf("%s formula=%q rank=%d fro=%s same=%t verdict=%q", r.Name, r.Formula, r.Rank, f64(r.FrobeniusNorm), r.SameAsPhiStar, r.Verdict)
	}
	return fmt.Sprintf("higher=%t nondegenerate=%t reason=%q rows={%s} verdict=%q", a.AnyHigherRankThanPhiStar, a.AnyNondegenerate, a.Reason, strings.Join(rows, "; "), a.Verdict)
}

func FormatBoundaryReadiness(b BoundaryReadinessAudit) string {
	return fmt.Sprintf("hodgePairing=%t k7w7=%t boundaryDim=%d needsW7Boundary=%t needsTrace=%t boundaryCertified=%t missing=%q verdict=%q", b.HodgePairingCertified, b.K7ToW7PairingFound, b.BoundaryPairDimension, b.StillRequiresW7ToBoundary, b.StillRequiresDefectTraceMap, b.BoundaryAssignmentCertified, b.MissingObject, b.Verdict)
}

func FormatNativeStatus(n NativeASHAStatus) string {
	return fmt.Sprintf("lambda4=%t PBPG=%t K7=%t W7=%t hodge=%t rank=%t nondegenerate=%t canonical=%t boundary=%t statement=%q verdict=%q", n.Lambda4Native, n.PBPGProjectorsConstructed, n.K7FrameConstructed, n.W7FrameConstructed, n.HodgeStarMatrixConstructed, n.HodgeRankComputed, n.HodgePairingNondegenerate, n.CanonicalK7ToW7PairingFound, n.BoundaryStressAssignmentNative, n.Statement, n.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("boundary=%t scalarRG=%t higgs=%t flavor=%t ckmPMNS=%t gauge=%t physicalOrientation=%t traceWeight=%t verdict=%q", f.ClaimsBoundaryStressAssignment, f.ClaimsScalarRGMatching, f.ClaimsHiggsMassDerivation, f.ClaimsFlavorDerivation, f.ClaimsCKMPMNSDerivation, f.ClaimsGaugeUnification, f.ClaimsPhysicalOrientation, f.ClaimsNativeTraceWeight, f.Verdict)
}
