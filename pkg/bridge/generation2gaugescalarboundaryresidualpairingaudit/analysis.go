// Package generation2gaugescalarboundaryresidualpairingaudit implements
// Gate 611: Gauge-Scalar Boundary Residual Pairing Audit.
//
// Gate 610 typed the strong-sector mismatch at Lambda_12 as a sign-compatible
// boundary inverse-coupling / color kinetic normalization slot. Gate 606 also
// exposed a negative scalar quartic at Lambda_12 in the v1 one-loop top-dominant
// scalar transport. Gate 611 asks whether those two residuals should be treated
// as a structured gauge-scalar boundary ledger or merely as a v1 numerical
// proximity, without deriving unification, Higgs stability, or a native ASHA
// boundary correction.
package generation2gaugescalarboundaryresidualpairingaudit

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/generation2colorkineticboundarycorrectionnormalizationaudit"
	"github.com/bagherbal/asha-engine/pkg/historytransport"
)

const (
	AuditID = "GATE611-GAUGE-SCALAR-BOUNDARY-RESIDUAL-PAIRING-AUDIT"

	StatusGate610Inherited              = "PASS_GATE610_COLOR_BOUNDARY_SLOT_INHERITED"
	StatusGate606ScalarInherited        = "PASS_GATE606_SCALAR_TRANSPORT_INHERITED"
	StatusResidualScalesCompared        = "PASS_STRONG_SCALAR_RESIDUAL_SCALES_COMPARED"
	StatusR3CloseToAbsLambdaConditional = "CONDITIONAL_SUPPORT_R3_MINUS_ONE_CLOSE_TO_ABS_LAMBDA_LAMBDA12_BUT_NOT_CERTIFIED"
	StatusScalarSlotDefined             = "PASS_SCALAR_BOUNDARY_CORRECTION_SLOT_DEFINED"
	StatusBothPositiveCorrections       = "CONDITIONAL_SUPPORT_BOTH_WOUNDS_REQUIRE_POSITIVE_BOUNDARY_CORRECTIONS"
	StatusJointLedgerDefined            = "CONDITIONAL_SUPPORT_JOINT_GAUGE_SCALAR_BOUNDARY_LEDGER_DEFINED"
	StatusSensitivityCaution            = "PASS_SCHEME_AND_HIGHER_LOOP_SENSITIVITY_LEDGER_RECORDED"
	StatusNoNativeGaugeScalarRelation   = "FAILED_ROUTE_NO_NATIVE_GAUGE_SCALAR_BOUNDARY_RELATION"
	StatusNoNativeLambdaZeroBoundary    = "FAILED_ROUTE_NO_NATIVE_LAMBDA_ZERO_BOUNDARY_THEOREM"
	StatusNoHiggsStabilityOrMassClaim   = "FAILED_ROUTE_NO_HIGGS_STABILITY_OR_MASS_CLAIM"
	StatusNoNativeColorKinetic          = "FAILED_ROUTE_NO_NATIVE_COLOR_KINETIC_CORRECTION_THEOREM"
	StatusGate611Boundary               = "FIREWALL_PRESERVED_GATE611_GAUGE_SCALAR_PAIRING_BOUNDARY"
	StatusNoUnificationClaim            = "FIREWALL_PRESERVED_NO_GAUGE_UNIFICATION_CLAIM"
	StatusNoScalarStabilityClaim        = "FIREWALL_PRESERVED_NO_FINAL_SCALAR_STABILITY_CLAIM"
	StatusNoEndpointDerivation          = "FIREWALL_PRESERVED_NO_OBSERVED_ENDPOINT_DERIVATION"
)

const (
	boundarySin2Theta = 3.0 / 8.0
	boundaryMWMZRatio = 5.0 / 8.0
	boundaryKY        = 5.0 / 3.0
)

type InheritedGate610 struct {
	Lambda12GeV    float64
	R3MinusOne     float64
	Delta3Required float64
	UStar          float64
	Eta3           float64
	DeltaAlpha3Inv float64
	Verdict        string
}

type InheritedScalarTransport struct {
	LambdaMZ             float64
	LambdaLambda12       float64
	AbsLambdaLambda12    float64
	BetaLambdaMZ         float64
	ZeroCrossingScaleGeV float64
	HasZeroCrossing      bool
	YT_MZ                float64
	YT_Lambda12          float64
	Approximation        string
	Verdict              string
}

type StrongScalarResidualComparison struct {
	StrongR3MinusOne         float64
	AbsLambdaLambda12        float64
	DifferenceAminusB        float64
	RatioAOverB              float64
	RelativeResidualVsB      float64
	WithinV1UncertaintyClaim bool
	Interpretation           string
	Verdict                  string
}

type BoundaryCoefficientComparisonRow struct {
	Quantity       string
	Formula        string
	Value          float64
	CompareToEta3  float64
	Difference     float64
	RelativeDelta  float64
	Interpretation string
	Verdict        string
}

type SignCompatibilityRow struct {
	Sector          string
	RuntimeWound    string
	NaturalVariable string
	RequiredShift   float64
	PositiveShift   bool
	Interpretation  string
	Verdict         string
}

type BoundaryCorrectionSlot struct {
	SlotName       string
	Formula        string
	RequiredValue  float64
	Target         string
	DiagnosticOnly bool
	Verdict        string
}

type JointBoundaryCorrectionVector struct {
	Delta3ColorBoundary float64
	DeltaLambdaBoundary float64
	Eta3                float64
	EtaLambda           float64
	ScalarNormalization string
	MeaningfulLedger    bool
	CertifiedRelation   bool
	Interpretation      string
	Verdict             string
}

type SensitivityAndSchemeCautionLedger struct {
	TwoLoopGaugeSensitive   bool
	TwoLoopScalarSensitive  bool
	TopMassSensitive        bool
	AlphaSSensitive         bool
	HiggsMatchingSensitive  bool
	ThresholdSensitive      bool
	Lambda12ChoiceSensitive bool
	ScalarMoreSensitive     bool
	ClosureCertified        bool
	Statement               string
	Verdict                 string
}

type NativeASHAStatus struct {
	ProvesNativeColorKineticCorrection bool
	ProvesNativeScalarQuarticBoundary  bool
	ProvesDeltaLambdaR3Relation        bool
	ProvesGaugeScalarThresholdTheorem  bool
	ProvesHiggsStabilityTheorem        bool
	ClaimsHiggsMassPrediction          bool
	Statement                          string
	Verdict                            string
}

type Firewalls struct {
	ClaimsLambdaZeroBoundaryDerived  bool
	ClaimsScalarStabilityDerived     bool
	ClaimsGaugeScalarRelationDerived bool
	ClaimsHiggsMassPredicted         bool
	ClaimsGaugeUnification           bool
	Verdict                          string
}

type Analysis struct {
	InheritedGauge         InheritedGate610
	InheritedScalar        InheritedScalarTransport
	ResidualComparison     StrongScalarResidualComparison
	CoefficientComparisons []BoundaryCoefficientComparisonRow
	SignCompatibility      []SignCompatibilityRow
	CorrectionSlots        []BoundaryCorrectionSlot
	JointVector            JointBoundaryCorrectionVector
	SensitivityLedger      SensitivityAndSchemeCautionLedger
	NativeStatus           NativeASHAStatus
	Firewalls              Firewalls
	Truth                  string
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
	g610, err := generation2colorkineticboundarycorrectionnormalizationaudit.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build Gate610 predecessor: %w", err)
	}
	b, err := historytransport.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("build history transport bundle: %w", err)
	}
	ig := inheritGauge(g610, b)
	is := inheritScalar(b)
	rc := compareResidualScales(ig, is)
	coeffs := compareBoundaryCoefficients(ig, is)
	signs := buildSignCompatibility(ig, is)
	slots := buildCorrectionSlots(ig, is)
	joint := buildJointVector(ig, is)
	sens := buildSensitivityLedger()
	native := buildNativeStatus()
	fw := auditFirewalls()
	truth := "Gate 611 pairs the sign-compatible color inverse-coupling boundary wound with the v1 scalar quartic boundary wound. R3-1 is numerically close to |lambda(Lambda12)| at the few-percent level, and eta3 is close to 2|lambda(Lambda12)| at the several-percent level, but this is not certified: scalar running is more scheme/top-mass/matching sensitive, and ASHA currently supplies no native gauge-scalar boundary relation, lambda=0 boundary theorem, Higgs stability theorem, or Higgs mass prediction."
	return Analysis{ig, is, rc, coeffs, signs, slots, joint, sens, native, fw, truth}, nil
}

func inheritGauge(g610 generation2colorkineticboundarycorrectionnormalizationaudit.Analysis, b historytransport.Bundle) InheritedGate610 {
	eta := g610.FractionalCorrection.EtaAgainstUStar
	return InheritedGate610{
		Lambda12GeV:    b.GaugeBoundary.Lambda12GeV,
		R3MinusOne:     b.GaugeBoundary.R3 - 1,
		Delta3Required: g610.Inherited.Delta3Required,
		UStar:          g610.Inherited.UStar,
		Eta3:           eta,
		DeltaAlpha3Inv: g610.Inherited.DeltaAlpha3Inv,
		Verdict:        StatusGate610Inherited,
	}
}

func inheritScalar(b historytransport.Bundle) InheritedScalarTransport {
	zero := math.NaN()
	has := b.ScalarTransport.ZeroCrossingScaleGeV != nil
	if has {
		zero = *b.ScalarTransport.ZeroCrossingScaleGeV
	}
	return InheritedScalarTransport{
		LambdaMZ:             b.ScalarTransport.LambdaMZ,
		LambdaLambda12:       b.ScalarTransport.LambdaLambda12,
		AbsLambdaLambda12:    math.Abs(b.ScalarTransport.LambdaLambda12),
		BetaLambdaMZ:         b.ScalarTransport.BetaLambdaMZ,
		ZeroCrossingScaleGeV: zero,
		HasZeroCrossing:      has,
		YT_MZ:                b.ScalarTransport.YT_MZ,
		YT_Lambda12:          b.ScalarTransport.YT_Lambda12,
		Approximation:        b.ScalarTransport.Approximation,
		Verdict:              StatusGate606ScalarInherited,
	}
}

func compareResidualScales(g InheritedGate610, s InheritedScalarTransport) StrongScalarResidualComparison {
	diff := g.R3MinusOne - s.AbsLambdaLambda12
	ratio := g.R3MinusOne / s.AbsLambdaLambda12
	rel := diff / s.AbsLambdaLambda12
	return StrongScalarResidualComparison{
		StrongR3MinusOne:         g.R3MinusOne,
		AbsLambdaLambda12:        s.AbsLambdaLambda12,
		DifferenceAminusB:        diff,
		RatioAOverB:              ratio,
		RelativeResidualVsB:      rel,
		WithinV1UncertaintyClaim: false,
		Interpretation:           "R3-1 and |lambda(Lambda12)| are close at the few-percent scale, but v1 scalar transport is top-dominant one-loop and threshold/scheme sensitive, so the proximity is recorded as a pairing clue only.",
		Verdict:                  StatusR3CloseToAbsLambdaConditional,
	}
}

func compareBoundaryCoefficients(g InheritedGate610, s InheritedScalarTransport) []BoundaryCoefficientComparisonRow {
	b := s.AbsLambdaLambda12
	rows := []struct {
		q, f string
		v    float64
		note string
	}{
		{"2|lambda(Lambda12)|", "2*abs(lambda(Lambda12))", 2 * b, "natural comparison to eta3 because the scalar quartic wound is negative and a zero-boundary diagnostic would require +|lambda|"},
		{"|lambda|/sin2(theta_*)", "abs(lambda(Lambda12))/(3/8)", b / boundarySin2Theta, "uses the ASHA boundary weak-angle value; included as typed comparison, not a fit"},
		{"|lambda|/(mW^2/mZ^2)_*", "abs(lambda(Lambda12))/(5/8)", b / boundaryMWMZRatio, "uses the symbolic boundary W/Z ratio; included as typed comparison, not a fit"},
		{"k_Y*|lambda|", "(5/3)*abs(lambda(Lambda12))", boundaryKY * b, "uses hypercharge trace normalization; included as typed comparison, not a fit"},
	}
	out := make([]BoundaryCoefficientComparisonRow, 0, len(rows))
	for _, r := range rows {
		diff := g.Eta3 - r.v
		out = append(out, BoundaryCoefficientComparisonRow{
			Quantity: r.q, Formula: r.f, Value: r.v, CompareToEta3: g.Eta3, Difference: diff, RelativeDelta: diff / r.v, Interpretation: r.note, Verdict: StatusResidualScalesCompared,
		})
	}
	return out
}

func buildSignCompatibility(g InheritedGate610, s InheritedScalarTransport) []SignCompatibilityRow {
	return []SignCompatibilityRow{
		{"strong gauge", "g3 too large at Lambda12; inverse coupling too small", "u3=1/g3^2", g.Delta3Required, g.Delta3Required > 0, "requires a positive inverse-kinetic/color boundary correction", StatusBothPositiveCorrections},
		{"scalar quartic", "lambda(Lambda12)<0 in v1 scalar transport", "lambda", s.AbsLambdaLambda12, s.AbsLambdaLambda12 > 0, "a lambda=0 diagnostic boundary target would require a positive quartic boundary correction", StatusBothPositiveCorrections},
	}
}

func buildCorrectionSlots(g InheritedGate610, s InheritedScalarTransport) []BoundaryCorrectionSlot {
	return []BoundaryCorrectionSlot{
		{"delta_3^color_boundary", "u3_eff = u3_runtime + delta_3^color_boundary", g.Delta3Required, "u3_eff=u_star at Lambda12", true, StatusGate610Inherited},
		{"delta_lambda_boundary", "lambda_eff(Lambda12)=lambda_runtime(Lambda12)+delta_lambda_boundary", s.AbsLambdaLambda12, "lambda_eff=0 diagnostic target only", true, StatusScalarSlotDefined},
	}
}

func buildJointVector(g InheritedGate610, s InheritedScalarTransport) JointBoundaryCorrectionVector {
	return JointBoundaryCorrectionVector{
		Delta3ColorBoundary: g.Delta3Required,
		DeltaLambdaBoundary: s.AbsLambdaLambda12,
		Eta3:                g.Eta3,
		EtaLambda:           s.AbsLambdaLambda12,
		ScalarNormalization: "no native scalar normalization analogous to u_star is currently available; eta_lambda is therefore recorded as the raw quartic correction slot",
		MeaningfulLedger:    true,
		CertifiedRelation:   false,
		Interpretation:      "The joint vector records two positive boundary correction slots: a color inverse-kinetic shift and a scalar quartic shift. It is meaningful as a history ledger but not a native relation.",
		Verdict:             StatusJointLedgerDefined,
	}
}

func buildSensitivityLedger() SensitivityAndSchemeCautionLedger {
	return SensitivityAndSchemeCautionLedger{
		TwoLoopGaugeSensitive:   true,
		TwoLoopScalarSensitive:  true,
		TopMassSensitive:        true,
		AlphaSSensitive:         true,
		HiggsMatchingSensitive:  true,
		ThresholdSensitive:      true,
		Lambda12ChoiceSensitive: true,
		ScalarMoreSensitive:     true,
		ClosureCertified:        false,
		Statement:               "The scalar wound is especially sensitive to top mass, alpha_s, Higgs pole/MSbar matching, loop order, threshold corrections, and the chosen boundary scale. The gauge wound is also threshold/scheme sensitive but currently cleaner as an inverse-coupling residual.",
		Verdict:                 StatusSensitivityCaution,
	}
}

func buildNativeStatus() NativeASHAStatus {
	return NativeASHAStatus{
		false, false, false, false, false, false,
		"ASHA currently supplies no native color kinetic correction, no scalar quartic boundary condition, no delta_lambda~R3 relation, no gauge-scalar threshold theorem, no Higgs stability theorem, and no Higgs mass prediction in this gate.",
		StatusNoNativeGaugeScalarRelation,
	}
}

func auditFirewalls() Firewalls {
	return Firewalls{false, false, false, false, false, StatusGate611Boundary}
}

func Statuses() []string {
	return []string{
		StatusGate610Inherited, StatusGate606ScalarInherited, StatusResidualScalesCompared, StatusR3CloseToAbsLambdaConditional, StatusScalarSlotDefined, StatusBothPositiveCorrections, StatusJointLedgerDefined, StatusSensitivityCaution, StatusNoNativeGaugeScalarRelation, StatusNoNativeLambdaZeroBoundary, StatusNoHiggsStabilityOrMassClaim, StatusNoNativeColorKinetic, StatusGate611Boundary, StatusNoUnificationClaim, StatusNoScalarStabilityClaim, StatusNoEndpointDerivation,
	}
}
