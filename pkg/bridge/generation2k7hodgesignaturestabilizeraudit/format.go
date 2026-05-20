package generation2k7hodgesignaturestabilizeraudit

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

func FormatInherited(i Gate633Inheritance) string {
	return fmt.Sprintf("H=%d K7=%d preserves=%t stable=%t noCompanion=%t noK7W7=%t boundaryMissing=%t firewall=%t verdict=%q", i.HDimension, i.K7Dimension, i.StarPreservesK7, i.K7HodgeStable, i.NoNewCompanionSevenPlane, i.NoK7ToW7Pairing, i.NoBoundaryAssignment, i.Gate633FirewallPreserved, i.Verdict)
}

func FormatRestrictedOperator(r RestrictedHodgeOperatorAudit) string {
	return fmt.Sprintf("formula=%q size=%dx%d symRes=%s orthRes=%s invRes=%s trace=%s det=%s maxAbs=%s orth=%t sym=%t involutive=%t verdict=%q", r.Formula, r.Rows, r.Cols, f64(r.SymmetryResidual), f64(r.OrthogonalityResidual), f64(r.InvolutionResidual), f64(r.Trace), f64(r.Determinant), f64(r.OperatorMaxAbs), r.Orthogonal, r.Symmetric, r.Involutive, r.Verdict)
}

func FormatSpectrum(s HodgeSpectrumAudit) string {
	return fmt.Sprintf("eig=%s plus=%d minus=%d trace=%s det=%s signature=%q self=%t anti=%t mixed=%t verdict=%q", f64s(s.Eigenvalues), s.PlusRank, s.MinusRank, f64(s.Trace), f64(s.Determinant), s.Signature, s.FullySelfDual, s.FullyAntiSelfDual, s.Mixed, s.Verdict)
}

func FormatInternalProjectors(p InternalProjectorAudit) string {
	return fmt.Sprintf("plusRank=%d minusRank=%d plusTrace=%s minusTrace=%s plusIdem=%s minusIdem=%s plusSym=%s minusSym=%s comp=%s orth=%s certified=%t verdict=%q", p.PlusProjectorRank, p.MinusProjectorRank, f64(p.PlusProjectorTrace), f64(p.MinusProjectorTrace), f64(p.PlusProjectorIdempotence), f64(p.MinusProjectorIdempotence), f64(p.PlusProjectorSymmetry), f64(p.MinusProjectorSymmetry), f64(p.ComplementarityResidual), f64(p.OrthogonalityResidual), p.ProjectorsCertified, p.Verdict)
}

func FormatAmbientProjection(a AmbientProjectionAudit) string {
	return fmt.Sprintf("star2Res=%s ambientTrace=%s ambient+=%d ambient-=%d ||P+QK||^2=%s ||P-QK||^2=%s frac+=%s frac-=%s res+=%s res-=%s verdict=%q", f64(a.AmbientHodgeStarSquaredResidual), f64(a.AmbientTrace), a.AmbientSelfDualRank, a.AmbientAntiSelfDualRank, f64(a.K7SelfDualFrobeniusSquared), f64(a.K7AntiSelfDualFrobeniusSquared), f64(a.K7SelfDualFraction), f64(a.K7AntiSelfDualFraction), f64(a.SelfDualContainmentResidual), f64(a.AntiSelfDualContainmentResidual), a.Verdict)
}

func FormatClassification(c SignatureClassification) string {
	return fmt.Sprintf("self=%t anti=%t mixed=%t plus=%d minus=%d trace=%s det=%s statement=%q verdict=%q", c.K7FullySelfDual, c.K7FullyAntiSelfDual, c.K7MixedHodgePolarity, c.PlusDimension, c.MinusDimension, f64(c.Trace), f64(c.Determinant), c.Statement, c.Verdict)
}

func FormatConsequences(c ConsequenceForPriorRoutes) string {
	return fmt.Sprintf("reopenK7W7=%t reopenV0=%t boundaryPromoted=%t sevenOver72Promoted=%t native=%q missing=%q boundaryVerdict=%q sevenOver72Verdict=%q", c.K7ToW7PairingReopened, c.OctonionicResidualReopened, c.BoundaryAssignmentPromoted, c.SevenOver72Promoted, c.NativeObjectDiscovered, c.RemainingMissingObject, c.VerdictBoundary, c.VerdictSevenOver72)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("boundary=%t sevenOver72=%t scalarRG=%t higgs=%t flavor=%t ckmPMNS=%t gauge=%t physicalOrientation=%t verdict=%q", f.ClaimsBoundaryStressAssignment, f.ClaimsSevenOver72Theorem, f.ClaimsScalarRGMatching, f.ClaimsHiggsMassDerivation, f.ClaimsFlavorDerivation, f.ClaimsCKMPMNSDerivation, f.ClaimsGaugeUnification, f.ClaimsPhysicalOrientation, f.Verdict)
}
