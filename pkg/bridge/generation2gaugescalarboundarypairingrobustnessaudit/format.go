package generation2gaugescalarboundarypairingrobustnessaudit

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
	return fmt.Sprintf("%.15g", x)
}

func FormatInheritedPairing(p InheritedPairing) string {
	return fmt.Sprintf("Lambda12=%s R3MinusOne=%s eta3=%s lambdaLambda12=%s absLambda12=%s R3/absLambda=%s eta3/(2absLambda)=%s verdict=%q", f64(p.Lambda12GeV), f64(p.R3MinusOne), f64(p.Eta3), f64(p.LambdaLambda12), f64(p.AbsLambda12), f64(p.RatioR3ToAbsLam), f64(p.Eta3To2AbsLambda), p.Verdict)
}

func FormatCandidateScale(s CandidateScale) string {
	return fmt.Sprintf("scale=%q GeV=%s t=%s role=%q verdict=%q", s.Name, f64(s.ScaleGeV), f64(s.LogMuOverMZ), s.Role, s.Verdict)
}
func FormatCandidateScales(rows []CandidateScale) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatCandidateScale(r))
	}
	return strings.Join(parts, " | ")
}

func FormatGaugeResidual(r GaugeResidualByScale) string {
	return fmt.Sprintf("scale=%q GeV=%s g=[%s,%s,%s] u=[%s,%s,%s] rel=%s eta=%s relDef=%q etaDef=%q verdict=%q", r.ScaleName, f64(r.ScaleGeV), f64(r.G1), f64(r.G2), f64(r.G3), f64(r.U1), f64(r.U2), f64(r.U3), f64(r.GaugeRelativeResidual), f64(r.InverseFractionalResidual), r.ResidualDefinition, r.InverseResidualDefinition, r.Verdict)
}
func FormatGaugeResiduals(rows []GaugeResidualByScale) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatGaugeResidual(r))
	}
	return strings.Join(parts, " | ")
}

func FormatScalarByScale(r ScalarByScale) string {
	return fmt.Sprintf("scale=%q GeV=%s lambda=%s abs=%s yt=%s betaLambda=%s approximation=%q verdict=%q", r.ScaleName, f64(r.ScaleGeV), f64(r.Lambda), f64(r.AbsLambda), f64(r.YT), f64(r.BetaLambda), r.Approximation, r.Verdict)
}
func FormatScalarValues(rows []ScalarByScale) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatScalarByScale(r))
	}
	return strings.Join(parts, " | ")
}

func FormatPairing(r PairingByScale) string {
	return fmt.Sprintf("scale=%q GeV=%s A=%s absLambda=%s A-B=%s A/B=%s relVsB=%s eta=%s 2absLambda=%s eta/(2absLambda)=%s eta-2absLambda=%s score=%s interpretation=%q verdict=%q", r.ScaleName, f64(r.ScaleGeV), f64(r.GaugeRelativeResidual), f64(r.AbsLambda), f64(r.Difference), f64(r.RatioGaugeToAbsLambda), f64(r.RelativeResidualVsAbsLambda), f64(r.InverseFractionalResidual), f64(r.TwoAbsLambda), f64(r.EtaToTwoAbsLambda), f64(r.EtaMinusTwoAbsLambda), f64(r.ClosenessScore), r.Interpretation, r.Verdict)
}
func FormatPairings(rows []PairingByScale) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatPairing(r))
	}
	return strings.Join(parts, " | ")
}

func FormatUniqueness(u Lambda12UniquenessAudit) string {
	return fmt.Sprintf("bestScale=%q bestScore=%s lambda12Score=%s lambda12Unique=%t nextBest=%q gap=%s statement=%q verdict=%q", u.BestScaleByCloseness, f64(u.BestClosenessScore), f64(u.Lambda12ClosenessScore), u.Lambda12UniqueBest, u.NextBestScale, f64(u.GapToNextBest), u.Statement, u.Verdict)
}

func FormatLocalSensitivity(s LocalSensitivityAudit) string {
	return fmt.Sprintf("Lambda12=%s lambda12=%s betaLambda12=%s R3MinusOne=%s deltaLambdaToR3=%s deltaLogToR3=%s scaleFactorR3=%s eta3=%s deltaLambdaToEta=%s deltaLogToEta=%s scaleFactorEta=%s statement=%q verdict=%q", f64(s.Lambda12GeV), f64(s.LambdaLambda12), f64(s.BetaLambdaLambda12), f64(s.R3MinusOne), f64(s.DeltaLambdaToExactR3Pairing), f64(s.DeltaLogMuToExactR3Pairing), f64(s.ScaleFactorToExactR3Pairing), f64(s.Eta3), f64(s.DeltaLambdaToExactEtaPairing), f64(s.DeltaLogMuToExactEtaPairing), f64(s.ScaleFactorToExactEtaPairing), s.FragilityStatement, s.Verdict)
}

func FormatSensitivity(s SensitivityAndSchemeCautionLedger) string {
	return fmt.Sprintf("twoLoopGauge=%t twoLoopScalar=%t top=%t alphaS=%t higgsMatching=%t threshold=%t scale=%t scalarFragile=%t closure=%t statement=%q verdict=%q", s.TwoLoopGaugeSensitive, s.TwoLoopScalarSensitive, s.TopMassSensitive, s.AlphaSSensitive, s.HiggsMatchingSensitive, s.ThresholdSensitive, s.ScaleChoiceSensitive, s.ScalarSideFragile, s.ClosureCertified, s.Statement, s.Verdict)
}

func FormatNativeStatus(n NativeASHAStatus) string {
	return fmt.Sprintf("lambda12Selection=%t strongLambdaRelation=%t scalarBoundary=%t jointTheorem=%t unification=%t higgsPrediction=%t statement=%q verdict=%q", n.ProvidesNativeLambda12Selection, n.ProvidesNativeStrongLambdaRelation, n.ProvidesNativeScalarBoundaryCondition, n.ProvidesNativeJointCorrectionTheorem, n.ProvidesNativeGaugeUnification, n.ClaimsHiggsPrediction, n.Statement, n.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("lambdaZero=%t higgsMass=%t scalarStability=%t gaugeUnification=%t thresholdExistence=%t endpoint=%t verdict=%q", f.ClaimsLambdaZeroBoundary, f.ClaimsHiggsMass, f.ClaimsScalarStability, f.ClaimsGaugeUnification, f.ClaimsThresholdExistence, f.DerivesEndpoint, f.Verdict)
}
