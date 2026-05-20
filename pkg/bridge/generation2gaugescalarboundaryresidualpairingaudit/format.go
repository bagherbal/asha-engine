package generation2gaugescalarboundaryresidualpairingaudit

import (
	"fmt"
	"math"
	"strings"
)

func f64(x float64) string {
	if math.IsNaN(x) {
		return "symbolic"
	}
	return fmt.Sprintf("%.15g", x)
}

func FormatInheritedGauge(g InheritedGate610) string {
	return fmt.Sprintf("Lambda12=%s R3MinusOne=%s delta3=%s uStar=%s eta3=%s deltaAlpha=%s verdict=%q", f64(g.Lambda12GeV), f64(g.R3MinusOne), f64(g.Delta3Required), f64(g.UStar), f64(g.Eta3), f64(g.DeltaAlpha3Inv), g.Verdict)
}
func FormatInheritedScalar(s InheritedScalarTransport) string {
	return fmt.Sprintf("lambdaMZ=%s lambdaLambda12=%s abs=%s betaLambdaMZ=%s zeroCrossing=%s hasZero=%t yT_MZ=%s yT_Lambda12=%s approximation=%q verdict=%q", f64(s.LambdaMZ), f64(s.LambdaLambda12), f64(s.AbsLambdaLambda12), f64(s.BetaLambdaMZ), f64(s.ZeroCrossingScaleGeV), s.HasZeroCrossing, f64(s.YT_MZ), f64(s.YT_Lambda12), s.Approximation, s.Verdict)
}
func FormatResidualComparison(r StrongScalarResidualComparison) string {
	return fmt.Sprintf("A=R3-1=%s B=abs(lambdaLambda12)=%s A-B=%s A/B=%s relVsB=%s uncertaintyClaim=%t interpretation=%q verdict=%q", f64(r.StrongR3MinusOne), f64(r.AbsLambdaLambda12), f64(r.DifferenceAminusB), f64(r.RatioAOverB), f64(r.RelativeResidualVsB), r.WithinV1UncertaintyClaim, r.Interpretation, r.Verdict)
}
func FormatCoefficientRow(r BoundaryCoefficientComparisonRow) string {
	return fmt.Sprintf("quantity=%q formula=%q value=%s eta3=%s diff=%s rel=%s interpretation=%q verdict=%q", r.Quantity, r.Formula, f64(r.Value), f64(r.CompareToEta3), f64(r.Difference), f64(r.RelativeDelta), r.Interpretation, r.Verdict)
}
func FormatCoefficientRows(rows []BoundaryCoefficientComparisonRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatCoefficientRow(r))
	}
	return strings.Join(parts, " | ")
}
func FormatSignRow(r SignCompatibilityRow) string {
	return fmt.Sprintf("sector=%q wound=%q variable=%q shift=%s positive=%t interpretation=%q verdict=%q", r.Sector, r.RuntimeWound, r.NaturalVariable, f64(r.RequiredShift), r.PositiveShift, r.Interpretation, r.Verdict)
}
func FormatSignRows(rows []SignCompatibilityRow) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatSignRow(r))
	}
	return strings.Join(parts, " | ")
}
func FormatCorrectionSlot(s BoundaryCorrectionSlot) string {
	return fmt.Sprintf("slot=%q formula=%q required=%s target=%q diagnostic=%t verdict=%q", s.SlotName, s.Formula, f64(s.RequiredValue), s.Target, s.DiagnosticOnly, s.Verdict)
}
func FormatCorrectionSlots(rows []BoundaryCorrectionSlot) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatCorrectionSlot(r))
	}
	return strings.Join(parts, " | ")
}
func FormatJointVector(v JointBoundaryCorrectionVector) string {
	return fmt.Sprintf("delta3=%s deltaLambda=%s eta3=%s etaLambda=%s scalarNorm=%q meaningful=%t certified=%t interpretation=%q verdict=%q", f64(v.Delta3ColorBoundary), f64(v.DeltaLambdaBoundary), f64(v.Eta3), f64(v.EtaLambda), v.ScalarNormalization, v.MeaningfulLedger, v.CertifiedRelation, v.Interpretation, v.Verdict)
}
func FormatSensitivity(s SensitivityAndSchemeCautionLedger) string {
	return fmt.Sprintf("twoLoopGauge=%t twoLoopScalar=%t top=%t alphaS=%t higgsMatching=%t threshold=%t lambda12=%t scalarMore=%t closure=%t statement=%q verdict=%q", s.TwoLoopGaugeSensitive, s.TwoLoopScalarSensitive, s.TopMassSensitive, s.AlphaSSensitive, s.HiggsMatchingSensitive, s.ThresholdSensitive, s.Lambda12ChoiceSensitive, s.ScalarMoreSensitive, s.ClosureCertified, s.Statement, s.Verdict)
}
func FormatNativeStatus(n NativeASHAStatus) string {
	return fmt.Sprintf("colorKinetic=%t scalarBoundary=%t relation=%t thresholdTheorem=%t stability=%t higgsMass=%t statement=%q verdict=%q", n.ProvesNativeColorKineticCorrection, n.ProvesNativeScalarQuarticBoundary, n.ProvesDeltaLambdaR3Relation, n.ProvesGaugeScalarThresholdTheorem, n.ProvesHiggsStabilityTheorem, n.ClaimsHiggsMassPrediction, n.Statement, n.Verdict)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("lambdaZero=%t stability=%t relation=%t higgsMass=%t unification=%t verdict=%q", f.ClaimsLambdaZeroBoundaryDerived, f.ClaimsScalarStabilityDerived, f.ClaimsGaugeScalarRelationDerived, f.ClaimsHiggsMassPredicted, f.ClaimsGaugeUnification, f.Verdict)
}
