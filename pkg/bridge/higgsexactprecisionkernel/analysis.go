// Package higgsexactprecisionkernel implements Gate 335:
// Exact Native Higgs Prediction / Arbitrary-Precision Numerical Kernel Audit.
//
// Gate 334 installed finite Passarino-Veltman basis functions but preserved the
// full pole-mass firewall. Gate 335 responds by separating the closed-form native
// branch from the unresolved pole-kernel branch: every number that follows from
// the Gate 330 doubled bosonic trace convention is recomputed with exact rational
// arithmetic plus a high-precision Machin π kernel. No float64 approximation is
// used for the native boundary, Higgs proxy, or precision-gap ledgers.
package higgsexactprecisionkernel

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"sync"
)

const (
	AuditID = "GATE335-EXACT-NATIVE-HIGGS-PREDICTION-ARBITRARY-PRECISION-NUMERICAL-KERNEL-AUDIT"

	StatusGate334Inherited             = "CONDITIONAL_SUPPORT_GATE334_PV_KERNEL_INHERITED"
	StatusExactRationalInputsInstalled = "CONDITIONAL_SUPPORT_EXACT_RATIONAL_INPUTS_INSTALLED"
	StatusHighPrecisionPiComputed      = "CONDITIONAL_SUPPORT_HIGH_PRECISION_MACHIN_PI_COMPUTED"
	StatusNativeClosedFormComputed     = "CONDITIONAL_SUPPORT_NATIVE_CLOSED_FORM_HIGGS_PROXY_COMPUTED"
	StatusPrecisionGapComputed         = "CONDITIONAL_SUPPORT_EXACT_PRECISION_GAP_COMPUTED"
	StatusSelfEnergyTargetComputed     = "CONDITIONAL_SUPPORT_EXACT_SELF_ENERGY_TARGET_COMPUTED"
	StatusEfficiencyLedgerFormalized   = "CONDITIONAL_SUPPORT_EFFICIENT_DETERMINISTIC_NUMERIC_KERNEL_FORMALIZED"
	StatusFirewallsPreserved           = "CONDITIONAL_SUPPORT_FULL_PRECISION_FIREWALLS_PRESERVED"

	StatusTensionFullPrecisionNotFullPole  = "CONDITIONAL_TENSION_FULL_PRECISION_NATIVE_BRANCH_IS_NOT_FULL_POLE_CALCULATION"
	StatusTensionObservedInputsQuarantined = "CONDITIONAL_TENSION_V_AND_MH_OBSERVED_INPUTS_REMAIN_QUARANTINED"

	StatusFailedFullPVContraction = "FAILED_ROUTE_FULL_PASSARINO_VELTMAN_CONTRACTION_NOT_EXECUTED"
	StatusFailedCounterterms      = "FAILED_ROUTE_RENORMALIZED_COUNTERTERMS_NOT_DERIVED"
	StatusFailedTwoLoop           = "FAILED_ROUTE_TWO_LOOP_PRECISION_NOT_COMPUTED"
	StatusFailedColliderClaim     = "FAILED_ROUTE_EXACT_COLLIDER_HIGGS_MASS_NOT_CLAIMED"
)

const (
	inheritedHighestGate = 334
	precisionBits        = uint(512)
	displayDigits        = 72
)

type ExactInputs struct {
	HighestInheritedGate int
	ContactShape         string
	ContactShapeRational *big.Rat
	LambdaHRational      *big.Rat
	GStarSquaredRational *big.Rat
	VEVRationalGeV       *big.Rat
	ObservedMassRational *big.Rat
	Status               string
}

type PiKernel struct {
	Formula       string
	PrecisionBits uint
	Pi            *big.Float
	AlphaInverse  *big.Float
	Alpha         *big.Float
	Status        string
}

type NativePrediction struct {
	LambdaH         *big.Rat
	LambdaHDecimal  *big.Float
	MassClosedForm  string
	MassGeV         *big.Float
	MassSquaredGeV2 *big.Rat
	Status          string
}

type PrecisionGap struct {
	ObservedMassGeV      *big.Rat
	DeltaMassGeV         *big.Float
	RelativeMassError    *big.Float
	DeltaMassSquaredGeV2 *big.Rat
	RequiredRePiGeV2     *big.Rat
	Status               string
}

type EfficiencyLedger struct {
	UsesFloat64ForNativeBranch bool
	PiAlgorithm                string
	RationalOperations         []string
	Deterministic              bool
	FullPVContractionExecuted  bool
	Status                     string
}

type Firewalls struct {
	NoPVContraction     bool
	NoCounterterms      bool
	NoTwoLoop           bool
	NoColliderPoleClaim bool
	QuarantinedInputs   []string
	Status              string
}

type Summary struct {
	HiggsMassGeV     string
	DeltaMassGeV     string
	RequiredRePiGeV2 string
	AlphaInverse     string
	DirectAnswer     string
	NextObligation   string
	Status           string
}

type Analysis struct {
	Inputs     ExactInputs
	Pi         PiKernel
	Native     NativePrediction
	Gap        PrecisionGap
	Efficiency EfficiencyLedger
	Firewalls  Firewalls
	Summary    Summary
	Truth      string
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
	inputs := installExactInputs()
	pi := computePiKernel()
	native := computeNativePrediction(inputs)
	gap := computePrecisionGap(inputs, native)
	efficiency := formalizeEfficiencyLedger()
	firewalls := preserveFirewalls()
	summary := compileSummary(pi, native, gap)
	truth := "Gate 335 recomputes the closed-form Gate 330 Higgs branch with exact rational arithmetic and a 512-bit Machin π kernel. The native branch gives m_H=v√(1197/4624)=125.274157149698971935740602811547... GeV and an exact pole-gap target ReΠ=504067437/11560000 GeV². This is full precision for the closed-form native branch, not a completed one-loop collider pole calculation."
	return Analysis{Inputs: inputs, Pi: pi, Native: native, Gap: gap, Efficiency: efficiency, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func installExactInputs() ExactInputs {
	return ExactInputs{
		HighestInheritedGate: inheritedHighestGate,
		ContactShape:         "1197/4624",
		ContactShapeRational: rat(1197, 4624),
		LambdaHRational:      rat(1197, 9248),
		GStarSquaredRational: rat(1, 2),
		VEVRationalGeV:       rat(12311, 50), // 246.22 GeV as exact decimal rational.
		ObservedMassRational: rat(1251, 10),  // 125.10 GeV as exact decimal rational.
		Status:               StatusExactRationalInputsInstalled,
	}
}

func computePiKernel() PiKernel {
	pi := machinPi(precisionBits)
	alphaInv := mulFloat(newFloat(8, precisionBits), pi, precisionBits)
	alpha := quoFloat(one(precisionBits), alphaInv, precisionBits)
	return PiKernel{Formula: "π=16 atan(1/5)-4 atan(1/239); α_GUT⁻¹=8π", PrecisionBits: precisionBits, Pi: pi, AlphaInverse: alphaInv, Alpha: alpha, Status: StatusHighPrecisionPiComputed}
}

func computeNativePrediction(in ExactInputs) NativePrediction {
	lambdaDec := ratToFloat(in.LambdaHRational, precisionBits)
	shapeDec := ratToFloat(in.ContactShapeRational, precisionBits)
	sqrtShape := new(big.Float).SetPrec(precisionBits).Sqrt(shapeDec)
	mass := mulFloat(ratToFloat(in.VEVRationalGeV, precisionBits), sqrtShape, precisionBits)
	mass2 := new(big.Rat).Mul(new(big.Rat).Mul(in.VEVRationalGeV, in.VEVRationalGeV), in.ContactShapeRational)
	return NativePrediction{LambdaH: new(big.Rat).Set(in.LambdaHRational), LambdaHDecimal: lambdaDec, MassClosedForm: "m_H=v√(1197/4624)", MassGeV: mass, MassSquaredGeV2: mass2, Status: StatusNativeClosedFormComputed}
}

func computePrecisionGap(in ExactInputs, native NativePrediction) PrecisionGap {
	observedF := ratToFloat(in.ObservedMassRational, precisionBits)
	deltaMass := subFloat(native.MassGeV, observedF, precisionBits)
	rel := quoFloat(deltaMass, observedF, precisionBits)
	rel = mulFloat(rel, newFloat(100, precisionBits), precisionBits)
	observed2 := new(big.Rat).Mul(in.ObservedMassRational, in.ObservedMassRational)
	delta2 := new(big.Rat).Sub(native.MassSquaredGeV2, observed2)
	requiredPi := new(big.Rat).Set(delta2)
	return PrecisionGap{ObservedMassGeV: new(big.Rat).Set(in.ObservedMassRational), DeltaMassGeV: deltaMass, RelativeMassError: rel, DeltaMassSquaredGeV2: delta2, RequiredRePiGeV2: requiredPi, Status: StatusSelfEnergyTargetComputed}
}

func formalizeEfficiencyLedger() EfficiencyLedger {
	return EfficiencyLedger{
		UsesFloat64ForNativeBranch: false,
		PiAlgorithm:                "Machin arctangent series with termination below 2^-(precision+32); rational-only closed-form Higgs branch",
		RationalOperations:         []string{"1197/4624", "g_*²=1/2", "λ=1197/9248", "v=12311/50", "M_obs=1251/10", "ReΠ=(v²·1197/4624)-(1251/10)²"},
		Deterministic:              true,
		FullPVContractionExecuted:  false,
		Status:                     StatusEfficiencyLedgerFormalized,
	}
}

func preserveFirewalls() Firewalls {
	return Firewalls{
		NoPVContraction:     true,
		NoCounterterms:      true,
		NoTwoLoop:           true,
		NoColliderPoleClaim: true,
		QuarantinedInputs:   []string{"v=246.22 GeV", "M_H reference=125.10 GeV"},
		Status:              StatusFirewallsPreserved,
	}
}

func compileSummary(pi PiKernel, native NativePrediction, gap PrecisionGap) Summary {
	direct := "The closed-form native branch is now evaluated exactly: α_GUT⁻¹=8π, g_*²=1/2, λ=1197/9248, and m_H=v√(1197/4624)=125.274157149698971935740602811547... GeV."
	next := "Contract the finite PV basis with a renormalized one-loop SM coefficient/counterterm table if an exact collider pole-mass claim is desired."
	return Summary{HiggsMassGeV: dec(native.MassGeV, displayDigits), DeltaMassGeV: dec(gap.DeltaMassGeV, displayDigits), RequiredRePiGeV2: gap.RequiredRePiGeV2.RatString(), AlphaInverse: dec(pi.AlphaInverse, displayDigits), DirectAnswer: direct, NextObligation: next, Status: StatusPrecisionGapComputed}
}

func Statuses(a Analysis) []string {
	return []string{
		StatusGate334Inherited,
		a.Inputs.Status,
		a.Pi.Status,
		a.Native.Status,
		a.Gap.Status,
		a.Efficiency.Status,
		a.Firewalls.Status,
		a.Summary.Status,
		StatusTensionFullPrecisionNotFullPole,
		StatusTensionObservedInputsQuarantined,
		StatusFailedFullPVContraction,
		StatusFailedCounterterms,
		StatusFailedTwoLoop,
		StatusFailedColliderClaim,
	}
}

func FormatInputs(x ExactInputs) string {
	return fmt.Sprintf("gate=%d shape=%s λ=%s g²=%s v=%s observed=%s status=%s", x.HighestInheritedGate, x.ContactShape, x.LambdaHRational.RatString(), x.GStarSquaredRational.RatString(), x.VEVRationalGeV.RatString(), x.ObservedMassRational.RatString(), x.Status)
}
func FormatPi(x PiKernel) string {
	return fmt.Sprintf("formula=%s bits=%d π=%s αInv=%s α=%s status=%s", x.Formula, x.PrecisionBits, dec(x.Pi, displayDigits), dec(x.AlphaInverse, displayDigits), dec(x.Alpha, displayDigits), x.Status)
}
func FormatNative(x NativePrediction) string {
	return fmt.Sprintf("λ=%s λDec=%s massFormula=%s mass=%s mass²=%s status=%s", x.LambdaH.RatString(), dec(x.LambdaHDecimal, displayDigits), x.MassClosedForm, dec(x.MassGeV, displayDigits), x.MassSquaredGeV2.RatString(), x.Status)
}
func FormatGap(x PrecisionGap) string {
	return fmt.Sprintf("observed=%s Δm=%s rel=%s%% Δm²=%s ReΠ=%s status=%s", x.ObservedMassGeV.RatString(), dec(x.DeltaMassGeV, displayDigits), dec(x.RelativeMassError, 36), x.DeltaMassSquaredGeV2.RatString(), x.RequiredRePiGeV2.RatString(), x.Status)
}
func FormatEfficiency(x EfficiencyLedger) string {
	return fmt.Sprintf("float64Native=%v deterministic=%v fullPV=%v pi=%s rationals=[%s] status=%s", x.UsesFloat64ForNativeBranch, x.Deterministic, x.FullPVContractionExecuted, x.PiAlgorithm, strings.Join(x.RationalOperations, "; "), x.Status)
}
func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("noPV=%v noCT=%v noTwoLoop=%v noPole=%v quarantined=[%s] status=%s", x.NoPVContraction, x.NoCounterterms, x.NoTwoLoop, x.NoColliderPoleClaim, strings.Join(x.QuarantinedInputs, ", "), x.Status)
}
func FormatSummary(x Summary) string {
	return fmt.Sprintf("m=%s Δm=%s ReΠ=%s αInv=%s answer=%s next=%s status=%s", x.HiggsMassGeV, x.DeltaMassGeV, x.RequiredRePiGeV2, x.AlphaInverse, x.DirectAnswer, x.NextObligation, x.Status)
}
func FormatStatuses(ss []string) string { return "statuses=" + strings.Join(ss, "; ") }

func rat(n, d int64) *big.Rat                        { return new(big.Rat).SetFrac(big.NewInt(n), big.NewInt(d)) }
func one(prec uint) *big.Float                       { return new(big.Float).SetPrec(prec).SetInt64(1) }
func newFloat(v int64, prec uint) *big.Float         { return new(big.Float).SetPrec(prec).SetInt64(v) }
func ratToFloat(r *big.Rat, prec uint) *big.Float    { return new(big.Float).SetPrec(prec).SetRat(r) }
func mulFloat(a, b *big.Float, prec uint) *big.Float { return new(big.Float).SetPrec(prec).Mul(a, b) }
func quoFloat(a, b *big.Float, prec uint) *big.Float { return new(big.Float).SetPrec(prec).Quo(a, b) }
func subFloat(a, b *big.Float, prec uint) *big.Float { return new(big.Float).SetPrec(prec).Sub(a, b) }
func absFloat(x *big.Float, prec uint) *big.Float {
	z := new(big.Float).SetPrec(prec).Set(x)
	if z.Sign() < 0 {
		z.Neg(z)
	}
	return z
}
func dec(x *big.Float, digits int) string { return x.Text('f', digits) }

func machinPi(prec uint) *big.Float {
	a := atanInvInt(5, prec)
	b := atanInvInt(239, prec)
	term1 := mulFloat(newFloat(16, prec), a, prec)
	term2 := mulFloat(newFloat(4, prec), b, prec)
	return subFloat(term1, term2, prec)
}

func atanInvInt(q int64, prec uint) *big.Float {
	x := quoFloat(one(prec), newFloat(q, prec), prec)
	x2 := mulFloat(x, x, prec)
	sum := new(big.Float).SetPrec(prec).SetInt64(0)
	power := new(big.Float).SetPrec(prec).Set(x)
	eps := new(big.Float).SetPrec(prec).SetFloat64(math.Ldexp(1, -int(prec)-32))
	sign := 1
	for n := int64(0); n < 10000; n++ {
		denom := newFloat(2*n+1, prec)
		add := quoFloat(power, denom, prec)
		if sign > 0 {
			sum.Add(sum, add)
		} else {
			sum.Sub(sum, add)
		}
		if absFloat(add, prec).Cmp(eps) <= 0 {
			break
		}
		power.Mul(power, x2)
		sign *= -1
	}
	return sum
}

func nearlyFloat(a *big.Float, target float64, tol float64) bool {
	af, _ := a.Float64()
	return math.Abs(af-target) <= tol
}
func nearlyRat(a *big.Rat, target float64, tol float64) bool {
	af, _ := new(big.Float).SetPrec(256).SetRat(a).Float64()
	return math.Abs(af-target) <= tol
}
