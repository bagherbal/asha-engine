package generation2gaugescalarboundarystresssealaudit

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

func FormatInherited(h Inherited) string {
	return fmt.Sprintf("Lambda12=%s R3MinusOne=%s lambda12=%s absLambda=%s eta3=%s 2absLambda=%s delta3=%s deltaLambda=%s sharpensAtLambda12=%t v1Sensitive=%t verdict=%q", f64(h.Lambda12GeV), f64(h.R3MinusOne), f64(h.LambdaLambda12), f64(h.AbsLambda12), f64(h.Eta3), f64(h.TwoAbsLambda12), f64(h.Delta3ColorBoundary), f64(h.DeltaLambdaBoundary), h.PairingSharpensAtLambda12, h.PairingIsV1Sensitive, h.Verdict)
}

func FormatCompressionCandidate(r CompressionCandidate) string {
	return fmt.Sprintf("name=%q xi=%s gauge=%s scalarAbs=%s gaugeMinusXi=%s scalarMinusXi=%s gaugeNorm=%s scalarNorm=%s maxNorm=%s construction=%q verdict=%q", r.Name, f64(r.Xi), f64(r.GaugeResidual), f64(r.ScalarAbsResidual), f64(r.GaugeMinusXi), f64(r.ScalarAbsMinusXi), f64(r.GaugeResidualNormalized), f64(r.ScalarResidualNormalized), f64(r.MaxAbsNormalizedResidual), r.Construction, r.Verdict)
}
func FormatCompressionCandidates(rows []CompressionCandidate) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatCompressionCandidate(r))
	}
	return strings.Join(parts, " | ")
}

func FormatSignedStress(s SignedStressVectorAudit) string {
	return fmt.Sprintf("S=(%s,%s) xiMean=%s xiGeom=%s SPlus=%s SMinus=%s idealPlus=%s idealMinus=%s statement=%q verdict=%q", f64(s.R3MinusOne), f64(s.LambdaLambda12), f64(s.XiMean), f64(s.XiGeom), f64(s.SPlus), f64(s.SMinus), f64(s.IdealVectorPlus), f64(s.IdealVectorMinus), s.Statement, s.Verdict)
}

func FormatAntiAlignment(a AntiAlignmentAudit) string {
	return fmt.Sprintf("SPlus=%s xiMean=%s relative=%s halfResidual=%s halfOverXi=%s ratioR3AbsLambda=%s antiAligned=%t statement=%q verdict=%q", f64(a.SPlus), f64(a.XiMean), f64(a.RelativeAntiAlignment), f64(a.HalfResidual), f64(a.HalfResidualOverXiMean), f64(a.RatioR3OverAbsLambda), a.AntiAligned, a.Statement, a.Verdict)
}

func FormatEtaComparison(r EtaComparison) string {
	return fmt.Sprintf("xiName=%q xi=%s eta3=%s 2xi=%s etaMinus2xi=%s etaOver2xi=%s relative=%s interpretation=%q verdict=%q", r.XiName, f64(r.Xi), f64(r.Eta3), f64(r.TwoXi), f64(r.EtaMinusTwoXi), f64(r.EtaOverTwoXi), f64(r.RelativeResidual), r.Interpretation, r.Verdict)
}
func FormatEtaComparisons(rows []EtaComparison) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatEtaComparison(r))
	}
	return strings.Join(parts, " | ")
}

func FormatStressSeal(s GaugeScalarBoundaryStressSeal) string {
	return fmt.Sprintf("scale=%s xi=%s strong=%s scalar=%s signed=%q eta3=%s etaApprox=%q native=%t interpretation=%q verdict=%q", f64(s.ScaleGeV), f64(s.XiBoundary), f64(s.StrongRelativeWound), f64(s.ScalarQuarticWound), s.SignedStressApproximation, f64(s.Eta3), s.Eta3Approximation, s.NativeCorrectionTheorem, s.Interpretation, s.Verdict)
}

func FormatRobustness(r RobustnessInheritance) string {
	return fmt.Sprintf("sharpensLambda12=%t scalarV1=%t higherLoop=%t threshold=%t matching=%t statement=%q verdict=%q", r.PairingSharpensAtLambda12, r.ScalarV1Sensitive, r.HigherLoopSensitive, r.ThresholdSensitive, r.MatchingSensitive, r.Statement, r.Verdict)
}

func FormatNativeStatus(n NativeASHAStatus) string {
	return fmt.Sprintf("xi=%t equation=%t colorKinetic=%t lambdaBoundary=%t scalarStability=%t unification=%t higgsPrediction=%t statement=%q verdict=%q", n.ProvidesNativeXiBoundary, n.ProvidesNativeGaugeScalarEquation, n.ProvidesNativeColorKineticCorrection, n.ProvidesNativeLambdaBoundary, n.ProvidesNativeScalarStability, n.ProvidesGaugeUnification, n.ClaimsHiggsPrediction, n.Statement, n.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("lambdaZero=%t higgsMass=%t scalarStability=%t unification=%t threshold=%t endpoint=%t verdict=%q", f.ClaimsLambdaZeroBoundary, f.ClaimsHiggsMass, f.ClaimsScalarStability, f.ClaimsGaugeUnification, f.ClaimsThresholdExistence, f.DerivesEndpoint, f.Verdict)
}
