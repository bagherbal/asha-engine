package generation2spectralactioncoefficientgrammaraudit

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
	return fmt.Sprintf("Lambda12=%s R3MinusOne=%s lambda12=%s xi=%s delta3=%s deltaLambda=%s eta3=%s EStress=%s EOverXi=%s 2xi=%s etaOver2xi=%s verdict=%q", f64(h.Lambda12GeV), f64(h.R3MinusOne), f64(h.LambdaLambda12), f64(h.XiBoundary), f64(h.Delta3ColorBoundary), f64(h.DeltaLambdaBoundary), f64(h.Eta3), f64(h.BoundaryResidual), f64(h.ResidualOverXi), f64(h.TwoXiBoundary), f64(h.EtaOverTwoXi), h.Verdict)
}

func FormatDependency(r CoefficientDependency) string {
	return fmt.Sprintf("coefficient=%q form=%q lane=%q depends=[%s] native=%t bridge=%t environmental=%t hostStress=%t obstruction=%q verdict=%q", r.Coefficient, r.SymbolicForm, r.Lane, strings.Join(r.DependsOn, ","), r.Native, r.AlgebraicBridge, r.Environmental, r.CanHostStressSlot, r.Obstruction, r.Verdict)
}
func FormatDependencies(rows []CoefficientDependency) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatDependency(r))
	}
	return strings.Join(parts, " | ")
}

func FormatSharedAudit(r SharedCoefficientAudit) string {
	return fmt.Sprintf("question=%q answer=%q shared=%t nativeRelation=%t requiredSeal=%q obstruction=%q verdict=%q", r.Question, r.Answer, r.Shared, r.NativeRelation, r.RequiredSeal, r.Obstruction, r.Verdict)
}
func FormatSharedAudits(rows []SharedCoefficientAudit) string {
	parts := make([]string, 0, len(rows))
	for _, r := range rows {
		parts = append(parts, FormatSharedAudit(r))
	}
	return strings.Join(parts, " | ")
}

func FormatColorDeformation(c ColorSpecificDeformationAudit) string {
	return fmt.Sprintf("deformation=%q bridge=%t nativeTrace=%t sectorF0=%t threshold=%t algebraExtension=%t obstruction=%q verdict=%q", c.Deformation, c.BridgeExpressible, c.NativeRepresentationTrace, c.RequiresSectorSplitF0, c.RequiresThresholdMatching, c.RequiresAlgebraExtension, c.Obstruction, c.Verdict)
}

func FormatScalarCorrection(s ScalarQuarticCorrectionAudit) string {
	return fmt.Sprintf("correction=%q bridge=%t viaBA2=%t viaF0=%t viaYukawaThreshold=%t viaMetric=%t viaMatching=%t native=%t obstruction=%q verdict=%q", s.Correction, s.BridgeExpressible, s.ViaBA2, s.ViaF0, s.ViaYukawaTraceThreshold, s.ViaScalarMetric, s.ViaMatching, s.Native, s.Obstruction, s.Verdict)
}

func FormatJointDeformation(j JointDeformationAudit) string {
	return fmt.Sprintf("deltaCoeff=%q shadow=%q bridge=%t forcesRatio=%t forcesEquation=%t native=%t residual=%s residualOverXi=%s statement=%q verdict=%q", j.DeltaCoeff, j.NormalizedShadow, j.BridgeExpressible, j.ForcesDeltaLambdaOverC3, j.ForcesStressEquation, j.KnownNativeRelation, f64(j.StressResidual), f64(j.ResidualOverXi), j.Statement, j.Verdict)
}

func FormatTypeConsistency(t TypeConsistencyLedger) string {
	return fmt.Sprintf("rawColor=%q rawScalar=%q rawSafe=%t normalized=[%s] normalizedSafe=%t statement=%q verdict=%q", t.RawColorType, t.RawScalarType, t.RawComparisonSafe, strings.Join(t.NormalizedForms, ";"), t.NormalizedSafe, t.Statement, t.Verdict)
}

func FormatNativeObstructions(n NativeObstructionLedger) string {
	return fmt.Sprintf("missing=[%s] nativeXi=%t nativeSU3=%t nativeC3Lambda=%t nativeF0=%t nativeLambdaBC=%t nativeThresholds=%t statement=%q verdict=%q", strings.Join(n.MissingStructures, ","), n.NativeXi, n.NativeSU3Only, n.NativeC3LambdaLaw, n.NativeF0Split, n.NativeLambdaBC, n.NativeThresholds, n.Statement, n.Verdict)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("xiNative=%t lambdaZero=%t higgsMass=%t higgsStability=%t unification=%t threshold=%t nativeCorrection=%t endpointDerivation=%t verdict=%q", f.ClaimsXiNative, f.ClaimsLambdaZero, f.ClaimsHiggsMass, f.ClaimsHiggsStability, f.ClaimsGaugeUnification, f.ClaimsThresholdExistence, f.ClaimsNativeCorrection, f.ClaimsObservedEndpointDerivation, f.Verdict)
}
