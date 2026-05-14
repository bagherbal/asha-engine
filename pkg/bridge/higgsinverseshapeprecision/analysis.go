package higgsinverseshapeprecision

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"sync"
)

const (
	GateNumber                = 336
	inheritedHighestGate      = 335
	precisionBits        uint = 768
	displayDigits             = 72

	StatusGate335Inherited              = "CONDITIONAL_SUPPORT_GATE335_EXACT_NATIVE_PRECISION_INHERITED"
	StatusExactInverseShapeComputed     = "CONDITIONAL_SUPPORT_EXACT_INVERSE_HIGGS_SHAPE_COMPUTED"
	StatusExactDeviationComputed        = "CONDITIONAL_SUPPORT_EXACT_CONTACT_SHAPE_DEVIATION_COMPUTED"
	StatusSelfEnergyEquivalenceComputed = "CONDITIONAL_SUPPORT_SELF_ENERGY_EQUIVALENCE_RECOMPUTED"
	StatusRequiredVEVComputed           = "CONDITIONAL_SUPPORT_REQUIRED_VEV_FOR_EXACT_NATIVE_SHAPE_COMPUTED"
	StatusPrecisionLedgerCompiled       = "CONDITIONAL_SUPPORT_FULL_PRECISION_INVERSE_LEDGER_COMPILED"

	StatusTensionNativeShapeSlightlyHigh = "CONDITIONAL_TENSION_NATIVE_CONTACT_SHAPE_EXCEEDS_OBSERVED_PROXY_BY_0_2786_PERCENT"
	StatusTensionMassProxyAboveTarget    = "CONDITIONAL_TENSION_NATIVE_MASS_PROXY_ABOVE_125_10_BY_0_174_GEV"

	StatusFailedPoleCorrectionNotComputed = "FAILED_ROUTE_POLE_CORRECTION_NOT_COMPUTED"
	StatusFailedNativeShapeNotAltered     = "FAILED_ROUTE_CONTACT_SHAPE_NOT_MODIFIED_TO_FIT_DATA"
	StatusFailedColliderMassNotClaimed    = "FAILED_ROUTE_EXACT_COLLIDER_HIGGS_MASS_NOT_CLAIMED"
)

type ExactInputs struct {
	HighestInheritedGate int
	NativeShape          *big.Rat
	LambdaH              *big.Rat
	VEVGeV               *big.Rat
	TargetMassGeV        *big.Rat
	Status               string
}

type InverseShape struct {
	ObservedShape     *big.Rat
	ObservedLambda    *big.Rat
	Formula           string
	ObservedShapeDec  *big.Float
	ObservedLambdaDec *big.Float
	Status            string
}

type Deviation struct {
	ShapeDelta         *big.Rat
	LambdaDelta        *big.Rat
	RelativeShapeError *big.Float
	RelativeMassError  *big.Float
	MassDeltaGeV       *big.Float
	NativeMassGeV      *big.Float
	TargetMassGeV      *big.Float
	Status             string
}

type SelfEnergyEquivalence struct {
	NativeMassSquaredGeV2 *big.Rat
	TargetMassSquaredGeV2 *big.Rat
	DifferenceGeV2        *big.Rat
	RequiredRePiGeV2      *big.Rat
	Formula               string
	Status                string
}

type RequiredVEV struct {
	VRequiredForTargetGeV *big.Float
	VShiftGeV             *big.Float
	RelativeVShiftPercent *big.Float
	Formula               string
	Status                string
}

type PrecisionLedger struct {
	UsesRationalCore  bool
	UsesFloat64Core   bool
	SqrtPrecisionBits uint
	ExactRationals    []string
	Interpretation    string
	Status            string
}

type Firewalls struct {
	NoPoleCorrection  bool
	NoFittingShape    bool
	NoColliderClaim   bool
	QuarantinedInputs []string
	Status            string
}

type Summary struct {
	NativeShape               string
	ObservedShape             string
	ShapeDelta                string
	RelativeShapeErrorPercent string
	NativeMassGeV             string
	MassDeltaGeV              string
	RequiredRePiGeV2          string
	DirectAnswer              string
	NextObligation            string
	Status                    string
}

type Analysis struct {
	Inputs      ExactInputs
	Inverse     InverseShape
	Deviation   Deviation
	SelfEnergy  SelfEnergyEquivalence
	RequiredVEV RequiredVEV
	Ledger      PrecisionLedger
	Firewalls   Firewalls
	Summary     Summary
	Truth       string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	inputs := installInputs()
	inv := computeInverseShape(inputs)
	dev := computeDeviation(inputs, inv)
	se := computeSelfEnergyEquivalence(inputs)
	reqv := computeRequiredVEV(inputs)
	ledger := compilePrecisionLedger()
	firewalls := preserveFirewalls()
	summary := compileSummary(inputs, inv, dev, se)
	truth := "Gate 336 performs the exact inverse comparison: the collider proxy shape is (125.10/246.22)^2=39125025/151560721, while the native contact shape is 1197/4624. The exact gap is 504067437/700816773904, a 0.2786225029% shape excess, equivalent to a +0.174157149699 GeV tree-mass excess and the same ReΠ target 504067437/11560000 GeV². This is a full-precision diagnostic ledger, not a fit and not a pole-mass calculation."
	return Analysis{Inputs: inputs, Inverse: inv, Deviation: dev, SelfEnergy: se, RequiredVEV: reqv, Ledger: ledger, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func installInputs() ExactInputs {
	return ExactInputs{HighestInheritedGate: inheritedHighestGate, NativeShape: rat(1197, 4624), LambdaH: rat(1197, 9248), VEVGeV: rat(12311, 50), TargetMassGeV: rat(1251, 10), Status: StatusGate335Inherited}
}

func computeInverseShape(in ExactInputs) InverseShape {
	m2 := new(big.Rat).Mul(in.TargetMassGeV, in.TargetMassGeV)
	v2 := new(big.Rat).Mul(in.VEVGeV, in.VEVGeV)
	obsShape := new(big.Rat).Quo(m2, v2)
	obsLambda := new(big.Rat).Quo(obsShape, rat(2, 1))
	return InverseShape{ObservedShape: obsShape, ObservedLambda: obsLambda, Formula: "R_obs=(M_H/v)^2; λ_obs=R_obs/2", ObservedShapeDec: ratToFloat(obsShape, precisionBits), ObservedLambdaDec: ratToFloat(obsLambda, precisionBits), Status: StatusExactInverseShapeComputed}
}

func computeDeviation(in ExactInputs, inv InverseShape) Deviation {
	shapeDelta := new(big.Rat).Sub(in.NativeShape, inv.ObservedShape)
	lambdaDelta := new(big.Rat).Sub(in.LambdaH, inv.ObservedLambda)
	relShape := quoFloat(ratToFloat(shapeDelta, precisionBits), ratToFloat(inv.ObservedShape, precisionBits), precisionBits)
	relShape.Mul(relShape, newFloat(100, precisionBits))
	nativeMass := sqrtRatMul(in.VEVGeV, in.NativeShape)
	targetMass := ratToFloat(in.TargetMassGeV, precisionBits)
	massDelta := subFloat(nativeMass, targetMass, precisionBits)
	relMass := quoFloat(massDelta, targetMass, precisionBits)
	relMass.Mul(relMass, newFloat(100, precisionBits))
	return Deviation{ShapeDelta: shapeDelta, LambdaDelta: lambdaDelta, RelativeShapeError: relShape, RelativeMassError: relMass, MassDeltaGeV: massDelta, NativeMassGeV: nativeMass, TargetMassGeV: targetMass, Status: StatusExactDeviationComputed}
}

func computeSelfEnergyEquivalence(in ExactInputs) SelfEnergyEquivalence {
	nativeM2 := new(big.Rat).Mul(new(big.Rat).Mul(in.VEVGeV, in.VEVGeV), in.NativeShape)
	targetM2 := new(big.Rat).Mul(in.TargetMassGeV, in.TargetMassGeV)
	diff := new(big.Rat).Sub(nativeM2, targetM2)
	return SelfEnergyEquivalence{NativeMassSquaredGeV2: nativeM2, TargetMassSquaredGeV2: targetM2, DifferenceGeV2: diff, RequiredRePiGeV2: new(big.Rat).Set(diff), Formula: "M_pole²-m_run²+ReΠ=0 => ReΠ=m_run²-M_pole²", Status: StatusSelfEnergyEquivalenceComputed}
}

func computeRequiredVEV(in ExactInputs) RequiredVEV {
	target := ratToFloat(in.TargetMassGeV, precisionBits)
	sqrtShape := new(big.Float).SetPrec(precisionBits).Sqrt(ratToFloat(in.NativeShape, precisionBits))
	vRequired := quoFloat(target, sqrtShape, precisionBits)
	vObserved := ratToFloat(in.VEVGeV, precisionBits)
	shift := subFloat(vRequired, vObserved, precisionBits)
	rel := quoFloat(shift, vObserved, precisionBits)
	rel.Mul(rel, newFloat(100, precisionBits))
	return RequiredVEV{VRequiredForTargetGeV: vRequired, VShiftGeV: shift, RelativeVShiftPercent: rel, Formula: "v_required=M_H/sqrt(1197/4624)", Status: StatusRequiredVEVComputed}
}

func compilePrecisionLedger() PrecisionLedger {
	return PrecisionLedger{UsesRationalCore: true, UsesFloat64Core: false, SqrtPrecisionBits: precisionBits, ExactRationals: []string{"R_native=1197/4624", "λ_native=1197/9248", "v=12311/50", "M=1251/10", "R_obs=39125025/151560721", "ΔR=504067437/700816773904", "Δλ=504067437/1401633547808", "ReΠ=504067437/11560000"}, Interpretation: "The derived contact shape is held fixed; the discrepancy is interpreted as a precision/pole conversion target, not as permission to retune the geometry.", Status: StatusPrecisionLedgerCompiled}
}

func preserveFirewalls() Firewalls {
	return Firewalls{NoPoleCorrection: true, NoFittingShape: true, NoColliderClaim: true, QuarantinedInputs: []string{"v=246.22 GeV", "M_H reference=125.10 GeV"}, Status: StatusFailedPoleCorrectionNotComputed}
}

func compileSummary(in ExactInputs, inv InverseShape, dev Deviation, se SelfEnergyEquivalence) Summary {
	direct := "The exact inverse diagnostic confirms the native contact shape is 0.2786225029% above the proxy shape extracted from 125.10 GeV and v=246.22 GeV; the corresponding closed-form mass is 125.274157149699 GeV."
	next := "The remaining mathematical obligation is a renormalized pole-mass correction ReΠ=504067437/11560000 GeV² or an equivalent precision scheme, not a modification of the contact shape."
	return Summary{NativeShape: in.NativeShape.RatString(), ObservedShape: inv.ObservedShape.RatString(), ShapeDelta: dev.ShapeDelta.RatString(), RelativeShapeErrorPercent: dec(dev.RelativeShapeError, 48), NativeMassGeV: dec(dev.NativeMassGeV, displayDigits), MassDeltaGeV: dec(dev.MassDeltaGeV, displayDigits), RequiredRePiGeV2: se.RequiredRePiGeV2.RatString(), DirectAnswer: direct, NextObligation: next, Status: StatusPrecisionLedgerCompiled}
}

func Statuses(a Analysis) []string {
	return []string{StatusGate335Inherited, a.Inverse.Status, a.Deviation.Status, a.SelfEnergy.Status, a.RequiredVEV.Status, a.Ledger.Status, StatusTensionNativeShapeSlightlyHigh, StatusTensionMassProxyAboveTarget, StatusFailedPoleCorrectionNotComputed, StatusFailedNativeShapeNotAltered, StatusFailedColliderMassNotClaimed}
}

func FormatInputs(x ExactInputs) string {
	return fmt.Sprintf("gate=%d R=%s λ=%s v=%s M=%s status=%s", x.HighestInheritedGate, x.NativeShape.RatString(), x.LambdaH.RatString(), x.VEVGeV.RatString(), x.TargetMassGeV.RatString(), x.Status)
}
func FormatInverse(x InverseShape) string {
	return fmt.Sprintf("formula=%s R_obs=%s R_obsDec=%s λ_obs=%s λ_obsDec=%s status=%s", x.Formula, x.ObservedShape.RatString(), dec(x.ObservedShapeDec, displayDigits), x.ObservedLambda.RatString(), dec(x.ObservedLambdaDec, displayDigits), x.Status)
}
func FormatDeviation(x Deviation) string {
	return fmt.Sprintf("ΔR=%s Δλ=%s relR=%s%% mNative=%s target=%s Δm=%s relM=%s%% status=%s", x.ShapeDelta.RatString(), x.LambdaDelta.RatString(), dec(x.RelativeShapeError, 48), dec(x.NativeMassGeV, displayDigits), dec(x.TargetMassGeV, displayDigits), dec(x.MassDeltaGeV, displayDigits), dec(x.RelativeMassError, 48), x.Status)
}
func FormatSelfEnergy(x SelfEnergyEquivalence) string {
	return fmt.Sprintf("formula=%s nativeM²=%s targetM²=%s diff=%s ReΠ=%s status=%s", x.Formula, x.NativeMassSquaredGeV2.RatString(), x.TargetMassSquaredGeV2.RatString(), x.DifferenceGeV2.RatString(), x.RequiredRePiGeV2.RatString(), x.Status)
}
func FormatRequiredVEV(x RequiredVEV) string {
	return fmt.Sprintf("formula=%s vRequired=%s shift=%s rel=%s%% status=%s", x.Formula, dec(x.VRequiredForTargetGeV, displayDigits), dec(x.VShiftGeV, displayDigits), dec(x.RelativeVShiftPercent, 48), x.Status)
}
func FormatLedger(x PrecisionLedger) string {
	return fmt.Sprintf("rational=%v float64=%v bits=%d exact=[%s] interpretation=%s status=%s", x.UsesRationalCore, x.UsesFloat64Core, x.SqrtPrecisionBits, strings.Join(x.ExactRationals, "; "), x.Interpretation, x.Status)
}
func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("noPole=%v noFit=%v noCollider=%v quarantined=[%s] status=%s", x.NoPoleCorrection, x.NoFittingShape, x.NoColliderClaim, strings.Join(x.QuarantinedInputs, ", "), x.Status)
}
func FormatSummary(x Summary) string {
	return fmt.Sprintf("R=%s R_obs=%s ΔR=%s rel=%s%% m=%s Δm=%s ReΠ=%s answer=%s next=%s status=%s", x.NativeShape, x.ObservedShape, x.ShapeDelta, x.RelativeShapeErrorPercent, x.NativeMassGeV, x.MassDeltaGeV, x.RequiredRePiGeV2, x.DirectAnswer, x.NextObligation, x.Status)
}
func FormatStatuses(ss []string) string { return "statuses=" + strings.Join(ss, "; ") }

func rat(n, d int64) *big.Rat                        { return new(big.Rat).SetFrac(big.NewInt(n), big.NewInt(d)) }
func ratToFloat(r *big.Rat, prec uint) *big.Float    { return new(big.Float).SetPrec(prec).SetRat(r) }
func newFloat(v int64, prec uint) *big.Float         { return new(big.Float).SetPrec(prec).SetInt64(v) }
func quoFloat(a, b *big.Float, prec uint) *big.Float { return new(big.Float).SetPrec(prec).Quo(a, b) }
func subFloat(a, b *big.Float, prec uint) *big.Float { return new(big.Float).SetPrec(prec).Sub(a, b) }
func dec(x *big.Float, digits int) string            { return x.Text('f', digits) }
func sqrtRatMul(v *big.Rat, r *big.Rat) *big.Float {
	return new(big.Float).SetPrec(precisionBits).Mul(ratToFloat(v, precisionBits), new(big.Float).SetPrec(precisionBits).Sqrt(ratToFloat(r, precisionBits)))
}
func nearlyFloat(a *big.Float, target float64, tol float64) bool {
	af, _ := a.Float64()
	return math.Abs(af-target) <= tol
}
func nearlyRat(a *big.Rat, target float64, tol float64) bool {
	af, _ := new(big.Float).SetPrec(256).SetRat(a).Float64()
	return math.Abs(af-target) <= tol
}
