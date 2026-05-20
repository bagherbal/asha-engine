package generation2hodgestarinternaldestinationaudit

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

func FormatInherited(i Gate632Inheritance) string {
	return fmt.Sprintf("H=%d U=%d V=%d K7=%d span=%d W7=%d leakageRank=%d ||PW*K7||=%s ||PUV*K7||=%s transverseFail=%t insideUV=%t boundaryMissing=%t firewall=%t verdict=%q", i.HDimension, i.UDimension, i.VDimension, i.K7Dimension, i.SpanDimension, i.W7Dimension, i.LeakageRank, f64(i.PWStarK7FrobeniusNorm), f64(i.PUVStarK7FrobeniusNorm), i.TransverseFailureCertified, i.StarK7InsideUPlusV, i.NoBoundaryAssignment, i.Gate632FirewallPreserved, i.Verdict)
}

func FormatL7(l L7Certificate) string {
	return fmt.Sprintf("def=%q size=%dx%d rank=%d iso=%s ||PW QL||=%s PUVfrac=%s PUVres=%s star2res=%s contained=%t verdict=%q", l.Definition, l.Rows, l.Cols, l.Rank, f64(l.QLOthonormalResidual), f64(l.PWQLFrobeniusNorm), f64(l.PUVProjectionFraction), f64(l.PUVContainmentResidual), f64(l.StarTwoCycleResidual), l.InternalContainmentCertified, l.Verdict)
}

func FormatK7Preservation(k K7PreservationAudit) string {
	return fmt.Sprintf("formula=%q rank=%d sv=%s det=%s frac=%s residual=%s preserves=%t stable=%t verdict=%q", k.MatrixFormula, k.Rank, f64s(k.SingularValues), f64(k.Determinant), f64(k.ProjectionFraction), f64(k.ContainmentResidual), k.StarPreservesK7, k.HodgeStable, k.Verdict)
}

func FormatT56(t T56InternalComplementAudit) string {
	return fmt.Sprintf("def=%q dim=%d rank=%d sv=%s frac=%s residual=%s inside=%t verdict=%q", t.Definition, t.Dimension, t.Rank, f64s(t.SingularValues), f64(t.ProjectionFraction), f64(t.ContainmentResidual), t.L7InsideT56, t.Verdict)
}

func FormatOctonionicResidual(o OctonionicResidualAudit) string {
	return fmt.Sprintf("def=%q V=%d K7=%d V0=%d rank=%d sv=%s frac=%s residual=%s equalsV0=%t Vsplit=%t verdict=%q", o.Definition, o.VDimension, o.K7Dimension, o.V0Dimension, o.Rank, f64s(o.SingularValues), f64(o.ProjectionFraction), f64(o.ContainmentResidual), o.L7EqualsV0, o.VDecomposesAsK7StarK7, o.Verdict)
}

func FormatBooleanResidual(b BooleanResidualAudit) string {
	return fmt.Sprintf("def=%q U=%d K7=%d U0=%d rank=%d sv=%s frac=%s residual=%s insideU0=%t verdict=%q", b.Definition, b.UDimension, b.K7Dimension, b.U0Dimension, b.Rank, f64s(b.SingularValues), f64(b.ProjectionFraction), f64(b.ContainmentResidual), b.L7InsideU0, b.Verdict)
}

func FormatOblique(o ObliqueDecompositionAudit) string {
	return fmt.Sprintf("def=%q T56=%d frac=%s residual=%s coordNorm=%s directResidual=%s oblique=%t reason=%q verdict=%q", o.Definition, o.CandidateT56Dimension, f64(o.ProjectionFraction), f64(o.ContainmentResidual), f64(o.DirectSumCoordinateNorm), f64(o.DirectSumResidual), o.ObliquePlaneDetected, o.Reason, o.Verdict)
}

func FormatStarTwoCycle(s StarTwoCycleAudit) string {
	return fmt.Sprintf("formula=%q star2res=%s equalsK7=%t equalsV0=%t equalsU0=%t carrier=%q verdict=%q", s.Formula, f64(s.StarSquaredResidual), s.L7EqualsK7, s.L7EqualsV0, s.L7EqualsU0, s.CarrierClassification, s.Verdict)
}

func FormatConsequence(c ConsequenceFor7Over72) string {
	return fmt.Sprintf("K7stable=%t oct14split=%t newPlane=%t boundaryDim=%d tracePromoted=%t boundaryMissing=%t statement=%q verdict=%q", c.K7Stable, c.Octonionic14HodgeSplit, c.NewSevenPlaneDiscovered, c.BoundaryPairDimension, c.TraceWeightPromoted, c.BoundaryAssignmentMissing, c.Statement, c.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("boundary=%t sevenOver72=%t scalarRG=%t higgs=%t flavor=%t ckmPMNS=%t gauge=%t verdict=%q", f.ClaimsBoundaryStressAssignment, f.ClaimsSevenOver72Theorem, f.ClaimsScalarRGMatching, f.ClaimsHiggsMassDerivation, f.ClaimsFlavorDerivation, f.ClaimsCKMPMNSDerivation, f.ClaimsGaugeUnification, f.Verdict)
}
