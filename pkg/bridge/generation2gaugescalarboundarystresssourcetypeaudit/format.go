package generation2gaugescalarboundarystresssourcetypeaudit

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
	return fmt.Sprintf("Lambda12=%s R3MinusOne=%s lambda12=%s xi=%s eta3=%s 2xi=%s etaOver2xi=%s delta3=%s deltaLambda=%s EStress=%s EOverXi=%s betaLambda12=%s sharpens=%t scalarV1=%t threshold=%t matching=%t higherLoop=%t verdict=%q", f64(h.Lambda12GeV), f64(h.R3MinusOne), f64(h.LambdaLambda12), f64(h.XiBoundary), f64(h.Eta3), f64(h.TwoXiBoundary), f64(h.EtaOverTwoXi), f64(h.Delta3ColorBoundary), f64(h.DeltaLambdaBoundary), f64(h.BoundaryStressResidual), f64(h.BoundaryStressRelToXi), f64(h.BetaLambdaLambda12), h.PairingSharpensLambda12, h.ScalarV1Sensitive, h.ThresholdSensitive, h.MatchingSensitive, h.HigherLoopSensitive, h.Verdict)
}

func FormatSourceType(r SourceTypeClassification) string {
	return fmt.Sprintf("name=%q signCompatible=%t support=%q obstruction=%q required=[%s] verdict=%q description=%q", r.Name, r.SignCompatible, r.CurrentSupport, r.PrimaryObstruction, strings.Join(r.RequiredData, ","), r.Verdict, r.Description)
}
func FormatSourceTypes(rows []SourceTypeClassification) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatSourceType(r))
	}
	return strings.Join(parts, " | ")
}

func FormatSpectralActionLane(r SpectralActionLane) string {
	return fmt.Sprintf("lane=%q form=%q canHostXi=%t canPairGaugeScalar=%t nativeRelation=%t obstruction=%q verdict=%q", r.Lane, r.SymbolicForm, r.CanHostXiSlot, r.CanPairGaugeScalar, r.NativeRelation, r.Obstruction, r.Verdict)
}
func FormatSpectralActionLanes(rows []SpectralActionLane) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatSpectralActionLane(r))
	}
	return strings.Join(parts, " | ")
}

func FormatKineticQuarticPairing(k KineticQuarticPairingAudit) string {
	return fmt.Sprintf("strong=%q scalar=%q sameHeatKernel=%t sameF0=%t sameFiniteTrace=%t sameNormalization=%t symbolicSlot=%t nativeLaw=%t statement=%q verdict=%q", k.StrongLane, k.ScalarLane, k.SameHeatKernelCoeff, k.SameF0Dependence, k.SameFiniteTraceCoeff, k.SameNormalizationRule, k.SymbolicPairingSlot, k.NativeCoefficientLaw, k.Statement, k.Verdict)
}

func FormatBoundaryEquation(e BoundaryStressEquationResidual) string {
	return fmt.Sprintf("equation=%q R3MinusOne=%s lambda12=%s residual=%s xi=%s absResidualOverXi=%s halfResidualOverXi=%s interpretation=%q verdict=%q", e.Equation, f64(e.R3MinusOne), f64(e.LambdaLambda12), f64(e.Residual), f64(e.XiBoundary), f64(e.AbsResidualOverXi), f64(e.HalfResidualOverXi), e.Interpretation, e.Verdict)
}

func FormatEtaRelation(e EtaRelationAudit) string {
	return fmt.Sprintf("eta3=%s xi=%s 2xi=%s etaMinus2xi=%s etaOver2xi=%s relative=%s interpretation=%q verdict=%q", f64(e.Eta3), f64(e.XiBoundary), f64(e.TwoXiBoundary), f64(e.EtaMinusTwoXi), f64(e.EtaOverTwoXi), f64(e.RelativeResidual), e.Interpretation, e.Verdict)
}

func FormatSensitivity(s RGArtifactSensitivityLedger) string {
	return fmt.Sprintf("betaLambda12=%s scalarV1=%t topMass=%t alphaS=%t twoLoopScalar=%t gaugeTwoLoop=%t threshold=%t matching=%t Lambda12Choice=%t statement=%q verdict=%q", f64(s.BetaLambdaLambda12), s.ScalarV1Sensitive, s.TopMassSensitive, s.AlphaSSensitive, s.TwoLoopScalarSensitive, s.GaugeTwoLoopSensitive, s.ThresholdSensitive, s.MatchingSensitive, s.Lambda12ChoiceSensitive, s.Statement, s.Verdict)
}

func FormatNativeStatus(n NativeASHAStatus) string {
	return fmt.Sprintf("xi=%t colorKinetic=%t scalarQuartic=%t f0Split=%t colorScalarCoeffLaw=%t thresholdSpectrum=%t higgsStability=%t unification=%t statement=%q verdict=%q", n.NativeXiBoundary, n.NativeColorKineticCorrection, n.NativeScalarQuarticBoundary, n.NativeF0SectorSplit, n.NativeGaugeScalarCoefficientLaw, n.NativeThresholdSpectrum, n.NativeHiggsStability, n.NativeGaugeUnification, n.Statement, n.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("threshold=%t higgsStability=%t lambdaZero=%t higgsMass=%t unification=%t WZHiggs=%t nativeCorrection=%t verdict=%q", f.ClaimsThresholdExists, f.ClaimsHiggsStability, f.ClaimsLambdaZeroBoundary, f.ClaimsHiggsMassPrediction, f.ClaimsGaugeUnification, f.ClaimsWZHiggsPrediction, f.ClaimsNativeCorrection, f.Verdict)
}
