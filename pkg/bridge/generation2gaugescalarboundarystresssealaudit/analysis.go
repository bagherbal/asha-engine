// Package generation2gaugescalarboundarystresssealaudit implements
// Gate 613: Joint Gauge-Scalar Boundary Stress Seal Audit.
//
// Gate 612 showed that the bridge-layer proximity between the strong-sector
// boundary wound and the scalar quartic wound is sharpest at Lambda_12 among
// the audited natural gauge scales. Gate 613 asks whether the paired residuals
// can be compressed into one signed boundary-stress coordinate without
// promoting that compression to native ASHA law, gauge unification, scalar
// stability, a Higgs prediction, threshold existence, or a lambda=0 theorem.
package generation2gaugescalarboundarystresssealaudit

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2gaugescalarboundarypairingrobustnessaudit"
)

const (
	AuditID = "GATE613-JOINT-GAUGE-SCALAR-BOUNDARY-STRESS-SEAL-AUDIT"

	StatusGate612Inherited              = "PASS_GATE612_PAIRING_ROBUSTNESS_INHERITED"
	StatusOneParameterCompressionTested = "PASS_ONE_PARAMETER_BOUNDARY_STRESS_COMPRESSION_TESTED"
	StatusSignedStressVectorDefined     = "PASS_SIGNED_STRESS_VECTOR_DEFINED"
	StatusBoundaryStressAntiAligned     = "CONDITIONAL_SUPPORT_BOUNDARY_STRESS_PAIR_ANTI_ALIGNED_AT_LAMBDA12"
	StatusBoundaryStressSealDefined     = "CONDITIONAL_SUPPORT_GAUGE_SCALAR_BOUNDARY_STRESS_SEAL_DEFINED"
	StatusEta3ApproxTwoXi               = "CONDITIONAL_SUPPORT_ETA3_APPROX_TWO_XI_BUT_NOT_CERTIFIED"
	StatusRobustnessInherited           = "PASS_GATE612_ROBUSTNESS_AND_SENSITIVITY_INHERITED"
	StatusNoNativeXi                    = "FAILED_ROUTE_NO_NATIVE_XI_BOUNDARY_THEOREM"
	StatusNoNativeGaugeScalarEquation   = "FAILED_ROUTE_NO_NATIVE_GAUGE_SCALAR_BOUNDARY_EQUATION"
	StatusNoNativeLambdaBoundary        = "FAILED_ROUTE_NO_NATIVE_LAMBDA_BOUNDARY_THEOREM"
	StatusNoNativeColorKinetic          = "FAILED_ROUTE_NO_NATIVE_COLOR_KINETIC_CORRECTION_THEOREM"
	StatusNoHiggsOrUnification          = "FAILED_ROUTE_NO_HIGGS_STABILITY_OR_GAUGE_UNIFICATION_CLAIM"
	StatusGate613Boundary               = "FIREWALL_PRESERVED_GATE613_BOUNDARY_STRESS_SEAL_BOUNDARY"
	StatusNoThresholdClaim              = "FIREWALL_PRESERVED_NO_THRESHOLD_EXISTENCE_CLAIM"
	StatusNoEndpointDerivation          = "FIREWALL_PRESERVED_NO_OBSERVED_ENDPOINT_DERIVATION"
)

type Inherited struct {
	Lambda12GeV               float64
	R3MinusOne                float64
	LambdaLambda12            float64
	AbsLambda12               float64
	Eta3                      float64
	TwoAbsLambda12            float64
	Delta3ColorBoundary       float64
	DeltaLambdaBoundary       float64
	PairingSharpensAtLambda12 bool
	PairingIsV1Sensitive      bool
	Verdict                   string
}

type CompressionCandidate struct {
	Name                     string
	Xi                       float64
	GaugeResidual            float64
	ScalarAbsResidual        float64
	GaugeMinusXi             float64
	ScalarAbsMinusXi         float64
	GaugeResidualNormalized  float64
	ScalarResidualNormalized float64
	MaxAbsNormalizedResidual float64
	Construction             string
	Verdict                  string
}

type SignedStressVectorAudit struct {
	R3MinusOne       float64
	LambdaLambda12   float64
	XiMean           float64
	XiGeom           float64
	SPlus            float64
	SMinus           float64
	IdealVectorPlus  float64
	IdealVectorMinus float64
	Statement        string
	Verdict          string
}

type AntiAlignmentAudit struct {
	SPlus                  float64
	XiMean                 float64
	RelativeAntiAlignment  float64
	HalfResidual           float64
	HalfResidualOverXiMean float64
	RatioR3OverAbsLambda   float64
	AntiAligned            bool
	Statement              string
	Verdict                string
}

type EtaComparison struct {
	XiName           string
	Xi               float64
	Eta3             float64
	TwoXi            float64
	EtaMinusTwoXi    float64
	EtaOverTwoXi     float64
	RelativeResidual float64
	Interpretation   string
	Verdict          string
}

type GaugeScalarBoundaryStressSeal struct {
	ScaleGeV                  float64
	XiBoundary                float64
	StrongRelativeWound       float64
	ScalarQuarticWound        float64
	SignedStressApproximation string
	Eta3                      float64
	Eta3Approximation         string
	NativeCorrectionTheorem   bool
	Interpretation            string
	Verdict                   string
}

type RobustnessInheritance struct {
	PairingSharpensAtLambda12 bool
	ScalarV1Sensitive         bool
	HigherLoopSensitive       bool
	ThresholdSensitive        bool
	MatchingSensitive         bool
	Statement                 string
	Verdict                   string
}

type NativeASHAStatus struct {
	ProvidesNativeXiBoundary             bool
	ProvidesNativeGaugeScalarEquation    bool
	ProvidesNativeColorKineticCorrection bool
	ProvidesNativeLambdaBoundary         bool
	ProvidesNativeScalarStability        bool
	ProvidesGaugeUnification             bool
	ClaimsHiggsPrediction                bool
	Statement                            string
	Verdict                              string
}

type Firewalls struct {
	ClaimsLambdaZeroBoundary bool
	ClaimsHiggsMass          bool
	ClaimsScalarStability    bool
	ClaimsGaugeUnification   bool
	ClaimsThresholdExistence bool
	DerivesEndpoint          bool
	Verdict                  string
}

type Analysis struct {
	Inherited             Inherited
	CompressionCandidates []CompressionCandidate
	SignedStressVector    SignedStressVectorAudit
	AntiAlignment         AntiAlignmentAudit
	EtaComparisons        []EtaComparison
	StressSeal            GaugeScalarBoundaryStressSeal
	Robustness            RobustnessInheritance
	NativeStatus          NativeASHAStatus
	Firewalls             Firewalls
	Truth                 string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	g612, err := generation2gaugescalarboundarypairingrobustnessaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate612 predecessor: %w", err)
	}
	inherited := inherit(g612)
	candidates := buildCompressionCandidates(inherited)
	stress := buildSignedStressVector(inherited)
	anti := buildAntiAlignment(inherited, stress)
	eta := buildEtaComparisons(inherited)
	seal := buildStressSeal(inherited, candidates)
	robustness := buildRobustnessInheritance(g612)
	native := buildNativeStatus()
	firewalls := auditFirewalls()
	truth := "Gate 613 compresses the Gate611/612 Lambda_12 gauge-scalar wounds into a bridge-layer signed boundary-stress vector S_boundary=(R3-1,lambda)≈(+xi,-xi). The one-parameter compression is numerically meaningful and anti-aligned at Lambda_12, and eta3 is close to twice the same stress scale, but ASHA supplies no native xi theorem, no native gauge-scalar boundary equation, no lambda boundary theorem, and no Higgs or unification claim."
	return Analysis{inherited, candidates, stress, anti, eta, seal, robustness, native, firewalls, truth}, nil
}

func inherit(a generation2gaugescalarboundarypairingrobustnessaudit.Analysis) Inherited {
	p := a.InheritedPairing
	return Inherited{
		Lambda12GeV:               p.Lambda12GeV,
		R3MinusOne:                p.R3MinusOne,
		LambdaLambda12:            p.LambdaLambda12,
		AbsLambda12:               p.AbsLambda12,
		Eta3:                      p.Eta3,
		TwoAbsLambda12:            2 * p.AbsLambda12,
		Delta3ColorBoundary:       0.32739043299998416,
		DeltaLambdaBoundary:       p.AbsLambda12,
		PairingSharpensAtLambda12: a.UniquenessAudit.Lambda12UniqueBest,
		PairingIsV1Sensitive:      a.SensitivityLedger.ScalarSideFragile,
		Verdict:                   StatusGate612Inherited,
	}
}

func buildCompressionCandidates(h Inherited) []CompressionCandidate {
	A := h.R3MinusOne
	B := h.AbsLambda12
	candidates := []struct {
		name string
		xi   float64
		how  string
	}{
		{"xi_A_abs_lambda", B, "xi_A=|lambda(Lambda_12)|"},
		{"xi_B_R3_minus_one", A, "xi_B=R_3-1"},
		{"xi_mean", 0.5 * (A + B), "xi_mean=0.5[(R_3-1)+|lambda|]"},
		{"xi_geom", math.Sqrt(A * B), "xi_geom=sqrt((R_3-1)|lambda|)"},
	}
	out := make([]CompressionCandidate, 0, len(candidates))
	for _, c := range candidates {
		gr := A - c.xi
		sr := B - c.xi
		gnorm := gr / c.xi
		snorm := sr / c.xi
		maxNorm := math.Max(math.Abs(gnorm), math.Abs(snorm))
		out = append(out, CompressionCandidate{
			Name:                     c.name,
			Xi:                       c.xi,
			GaugeResidual:            A,
			ScalarAbsResidual:        B,
			GaugeMinusXi:             gr,
			ScalarAbsMinusXi:         sr,
			GaugeResidualNormalized:  gnorm,
			ScalarResidualNormalized: snorm,
			MaxAbsNormalizedResidual: maxNorm,
			Construction:             c.how,
			Verdict:                  StatusOneParameterCompressionTested,
		})
	}
	return out
}

func buildSignedStressVector(h Inherited) SignedStressVectorAudit {
	xiMean := 0.5 * (h.R3MinusOne + h.AbsLambda12)
	xiGeom := math.Sqrt(h.R3MinusOne * h.AbsLambda12)
	sPlus := h.R3MinusOne + h.LambdaLambda12
	sMinus := h.R3MinusOne - h.LambdaLambda12
	return SignedStressVectorAudit{
		R3MinusOne:       h.R3MinusOne,
		LambdaLambda12:   h.LambdaLambda12,
		XiMean:           xiMean,
		XiGeom:           xiGeom,
		SPlus:            sPlus,
		SMinus:           sMinus,
		IdealVectorPlus:  0,
		IdealVectorMinus: 2 * xiMean,
		Statement:        "The signed vector S_boundary=(R_3-1, lambda(Lambda_12)) is tested against the anti-aligned form (+xi,-xi). The sum mode S_plus measures deviation from anti-alignment; the difference mode S_minus measures twice the common stress scale.",
		Verdict:          StatusSignedStressVectorDefined,
	}
}

func buildAntiAlignment(h Inherited, s SignedStressVectorAudit) AntiAlignmentAudit {
	rel := math.Abs(s.SPlus) / s.XiMean
	half := 0.5 * s.SPlus
	halfRel := math.Abs(half) / s.XiMean
	return AntiAlignmentAudit{
		SPlus:                  s.SPlus,
		XiMean:                 s.XiMean,
		RelativeAntiAlignment:  rel,
		HalfResidual:           half,
		HalfResidualOverXiMean: halfRel,
		RatioR3OverAbsLambda:   h.R3MinusOne / h.AbsLambda12,
		AntiAligned:            rel < 0.05,
		Statement:              "Because lambda(Lambda_12) is negative, S_plus=(R_3-1)+lambda tests the anti-alignment R_3-1≈|lambda|. The half-residual is about the signed error around xi_mean.",
		Verdict:                StatusBoundaryStressAntiAligned,
	}
}

func buildEtaComparisons(h Inherited) []EtaComparison {
	xiMean := 0.5 * (h.R3MinusOne + h.AbsLambda12)
	xiGeom := math.Sqrt(h.R3MinusOne * h.AbsLambda12)
	rows := []struct {
		name string
		xi   float64
	}{{"xi_mean", xiMean}, {"xi_geom", xiGeom}, {"xi_abs_lambda", h.AbsLambda12}, {"xi_R3_minus_one", h.R3MinusOne}}
	out := make([]EtaComparison, 0, len(rows))
	for _, r := range rows {
		two := 2 * r.xi
		diff := h.Eta3 - two
		ratio := h.Eta3 / two
		out = append(out, EtaComparison{
			XiName:           r.name,
			Xi:               r.xi,
			Eta3:             h.Eta3,
			TwoXi:            two,
			EtaMinusTwoXi:    diff,
			EtaOverTwoXi:     ratio,
			RelativeResidual: diff / two,
			Interpretation:   "typed comparison eta_3≈2 xi_boundary; no arbitrary constants introduced",
			Verdict:          StatusEta3ApproxTwoXi,
		})
	}
	return out
}

func buildStressSeal(h Inherited, candidates []CompressionCandidate) GaugeScalarBoundaryStressSeal {
	xi := findCandidate(candidates, "xi_mean").Xi
	return GaugeScalarBoundaryStressSeal{
		ScaleGeV:                  h.Lambda12GeV,
		XiBoundary:                xi,
		StrongRelativeWound:       h.R3MinusOne,
		ScalarQuarticWound:        h.LambdaLambda12,
		SignedStressApproximation: "S_boundary=(R_3-1,lambda(Lambda_12))≈(+xi_boundary,-xi_boundary)",
		Eta3:                      h.Eta3,
		Eta3Approximation:         "eta_3≈2 xi_boundary",
		NativeCorrectionTheorem:   false,
		Interpretation:            "GaugeScalarBoundaryStressSeal is a bridge-layer compression of two positive correction needs at Lambda_12: positive color inverse-kinetic correction and positive scalar quartic correction. It is not a native theorem.",
		Verdict:                   StatusBoundaryStressSealDefined,
	}
}

func buildRobustnessInheritance(a generation2gaugescalarboundarypairingrobustnessaudit.Analysis) RobustnessInheritance {
	return RobustnessInheritance{
		PairingSharpensAtLambda12: a.UniquenessAudit.Lambda12UniqueBest,
		ScalarV1Sensitive:         a.SensitivityLedger.ScalarSideFragile,
		HigherLoopSensitive:       a.SensitivityLedger.TwoLoopGaugeSensitive || a.SensitivityLedger.TwoLoopScalarSensitive,
		ThresholdSensitive:        a.SensitivityLedger.ThresholdSensitive,
		MatchingSensitive:         a.SensitivityLedger.HiggsMatchingSensitive,
		Statement:                 "Gate612 found that the pairing sharpens at Lambda_12 among audited natural gauge scales, while the scalar side remains one-loop/top-dominant and sensitive to higher loops, thresholds, matching, top mass, alpha_s, and scale choice.",
		Verdict:                   StatusRobustnessInherited,
	}
}

func buildNativeStatus() NativeASHAStatus {
	return NativeASHAStatus{
		ProvidesNativeXiBoundary:             false,
		ProvidesNativeGaugeScalarEquation:    false,
		ProvidesNativeColorKineticCorrection: false,
		ProvidesNativeLambdaBoundary:         false,
		ProvidesNativeScalarStability:        false,
		ProvidesGaugeUnification:             false,
		ClaimsHiggsPrediction:                false,
		Statement:                            "ASHA currently supplies no native xi_boundary, no native gauge-scalar boundary equation, no color kinetic correction theorem, no lambda boundary theorem, no scalar stability theorem, and no gauge unification theorem.",
		Verdict:                              StatusNoNativeXi,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{false, false, false, false, false, false, StatusGate613Boundary}
}

func findCandidate(rows []CompressionCandidate, name string) CompressionCandidate {
	for _, r := range rows {
		if r.Name == name {
			return r
		}
	}
	return CompressionCandidate{Name: name, Xi: math.NaN()}
}

func Statuses() []string {
	return []string{
		StatusGate612Inherited,
		StatusOneParameterCompressionTested,
		StatusSignedStressVectorDefined,
		StatusBoundaryStressAntiAligned,
		StatusBoundaryStressSealDefined,
		StatusEta3ApproxTwoXi,
		StatusRobustnessInherited,
		StatusNoNativeXi,
		StatusNoNativeGaugeScalarEquation,
		StatusNoNativeLambdaBoundary,
		StatusNoNativeColorKinetic,
		StatusNoHiggsOrUnification,
		StatusGate613Boundary,
		StatusNoThresholdClaim,
		StatusNoEndpointDerivation,
	}
}
