package higgsprecisionroutesieve

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"sync"
)

const (
	GateNumber                = 337
	inheritedHighestGate      = 336
	precisionBits        uint = 768
	displayDigits             = 72

	StatusGate336Inherited              = "CONDITIONAL_SUPPORT_GATE336_EXACT_INVERSE_SHAPE_INHERITED"
	StatusRepairRoutesAudited           = "CONDITIONAL_SUPPORT_PRECISION_REPAIR_ROUTES_AUDITED"
	StatusOneLoopKernelComputed         = "CONDITIONAL_SUPPORT_ONE_LOOP_COMPONENT_KERNEL_RECOMPUTED_HIGH_PRECISION"
	StatusFiniteCountertermTargetSolved = "CONDITIONAL_SUPPORT_FINITE_COUNTERTERM_TARGET_SOLVED"
	StatusPoleCorrectionPreferred       = "CONDITIONAL_SUPPORT_POLE_CORRECTION_BRANCH_PREFERRED_OVER_CONTACT_SHAPE_FIT"
	StatusPrecisionLedgerCompiled       = "CONDITIONAL_SUPPORT_EXACT_EFFICIENT_PRECISION_LEDGER_COMPILED"

	StatusTensionRawKernelWrongSignMagnitude = "CONDITIONAL_TENSION_RAW_ONE_LOOP_KERNEL_HAS_WRONG_SIGN_AND_REQUIRES_RENORMALIZED_FINITE_PART"
	StatusTensionContactShapeFitForbidden    = "CONDITIONAL_TENSION_CONTACT_SHAPE_DEFORMATION_WOULD_DESTROY_NATIVE_RATIO"

	StatusFailedFullSMRenormalizedSelfEnergy = "FAILED_ROUTE_FULL_SM_RENORMALIZED_SELF_ENERGY_NOT_COMPUTED"
	StatusFailedCountertermsNotDerived       = "FAILED_ROUTE_FINITE_COUNTERTERMS_NOT_DERIVED_FROM_NATIVE_SCHEME"
	StatusFailedContactShapeNotModified      = "FAILED_ROUTE_CONTACT_SHAPE_NOT_MODIFIED_TO_FIT_DATA"
	StatusFailedColliderMassNotClaimed       = "FAILED_ROUTE_EXACT_COLLIDER_HIGGS_MASS_NOT_CLAIMED"
)

type Inputs struct {
	HighestInheritedGate int
	ContactShape         *big.Rat
	LambdaH              *big.Rat
	VEVGeV               *big.Rat
	TargetMassGeV        *big.Rat
	NativeMassGeV        *big.Float
	RequiredRePiGeV2     *big.Rat
	Status               string
}

type RepairRoute struct {
	Name                string
	Formula             string
	RequiredValue       string
	AllowedByNativeCore bool
	Interpretation      string
}

type RepairRoutes struct {
	Routes []RepairRoute
	Status string
}

type LoopInput struct {
	Name    string
	MassGeV *big.Rat
	Coeff   int64
	Source  string
}

type LoopComponent struct {
	Name             string
	Coeff            int64
	MassGeV          *big.Rat
	NumeratorGeV4    *big.Float
	ContributionGeV2 *big.Float
}

type OneLoopKernel struct {
	Formula          string
	ScaleDescription string
	Components       []LoopComponent
	RawKernelGeV2    *big.Float
	Status           string
}

type CountertermTarget struct {
	RequiredRePiGeV2      *big.Rat
	RawKernelGeV2         *big.Float
	FiniteRemainderGeV2   *big.Float
	RemainderOverRequired *big.Float
	RawOverRequired       *big.Float
	Interpretation        string
	Status                string
}

type Recommendation struct {
	BestRoute      string
	Reason         string
	NextGateTarget string
	Status         string
}

type Firewalls struct {
	NoFullRenormalizedSM bool
	NoNativeCounterterms bool
	NoShapeFit           bool
	NoColliderClaim      bool
	QuarantinedInputs    []string
	Statuses             []string
}

type Summary struct {
	NativeMassGeV       string
	RequiredRePiGeV2    string
	RawKernelGeV2       string
	FiniteRemainderGeV2 string
	ContactShapeDelta   string
	DirectAnswer        string
	Status              string
}

type Analysis struct {
	Inputs         Inputs
	Routes         RepairRoutes
	Kernel         OneLoopKernel
	Counterterm    CountertermTarget
	Recommendation Recommendation
	Firewalls      Firewalls
	Summary        Summary
	Truth          string
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
	routes := auditRepairRoutes(inputs)
	kernel := computeOneLoopKernel(inputs)
	counter := solveCounterterm(inputs, kernel)
	rec := recommendRoute(routes, counter)
	firewalls := preserveFirewalls()
	summary := compileSummary(inputs, routes, kernel, counter)
	truth := "Gate 337 audits the exact repair routes for the 43.604449567474 GeV² precision gap. Deforming the contact shape would fit the number but destroy the native 1197/4624 ratio, while the raw one-loop polynomial kernel is large and negative before renormalized finite parts. The mathematically preferred route is a real pole-correction calculation: compute the renormalized finite self-energy/counterterm ledger rather than modifying the contact geometry."
	return Analysis{Inputs: inputs, Routes: routes, Kernel: kernel, Counterterm: counter, Recommendation: rec, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func installInputs() Inputs {
	contact := rat(1197, 4624)
	lambda := rat(1197, 9248)
	v := rat(12311, 50)
	target := rat(1251, 10)
	nativeMass := sqrtRatMul(v, contact)
	nativeM2 := new(big.Rat).Mul(new(big.Rat).Mul(v, v), contact)
	targetM2 := new(big.Rat).Mul(target, target)
	required := new(big.Rat).Sub(nativeM2, targetM2)
	return Inputs{HighestInheritedGate: inheritedHighestGate, ContactShape: contact, LambdaH: lambda, VEVGeV: v, TargetMassGeV: target, NativeMassGeV: nativeMass, RequiredRePiGeV2: required, Status: StatusGate336Inherited}
}

func auditRepairRoutes(in Inputs) RepairRoutes {
	obsShape := new(big.Rat).Quo(new(big.Rat).Mul(in.TargetMassGeV, in.TargetMassGeV), new(big.Rat).Mul(in.VEVGeV, in.VEVGeV))
	deltaShape := new(big.Rat).Sub(obsShape, in.ContactShape)
	vReq := quoFloat(ratToFloat(in.TargetMassGeV, precisionBits), new(big.Float).SetPrec(precisionBits).Sqrt(ratToFloat(in.ContactShape, precisionBits)), precisionBits)
	vShift := subFloat(vReq, ratToFloat(in.VEVGeV, precisionBits), precisionBits)
	return RepairRoutes{Routes: []RepairRoute{
		{Name: "contact-shape deformation", Formula: "R_contact -> R_obs=(M_H/v)^2", RequiredValue: deltaShape.RatString(), AllowedByNativeCore: false, Interpretation: "Would erase the native 1197/4624 contact scalar shape; rejected except as diagnostic."},
		{Name: "electroweak VEV deformation", Formula: "v_required=M_H/sqrt(1197/4624)", RequiredValue: dec(vShift, 48) + " GeV shift", AllowedByNativeCore: false, Interpretation: "Moves the empirical electroweak input rather than computing the pole correction; diagnostic only."},
		{Name: "pole self-energy correction", Formula: "M_pole²-m_run²+ReΠ=0", RequiredValue: in.RequiredRePiGeV2.RatString() + " GeV²", AllowedByNativeCore: true, Interpretation: "Preserves contact geometry and asks for the missing renormalized finite pole correction."},
	}, Status: StatusRepairRoutesAudited}
}

func computeOneLoopKernel(in Inputs) OneLoopKernel {
	// Quarantined pole inputs used only for the continuum precision ledger.
	inputs := []LoopInput{
		{Name: "top", MassGeV: rat(4319, 25), Coeff: -12, Source: "quarantined m_t=172.76 GeV"},
		{Name: "W", MassGeV: rat(80379, 1000), Coeff: 6, Source: "quarantined m_W=80.379 GeV"},
		{Name: "Z", MassGeV: rat(455938, 5000), Coeff: 3, Source: "quarantined m_Z=91.1876 GeV"},
		{Name: "H-native", MassGeV: nil, Coeff: 3, Source: "native m_H=v sqrt(1197/4624)"},
	}
	pi := machinPi(precisionBits)
	v2 := mulFloat(ratToFloat(in.VEVGeV, precisionBits), ratToFloat(in.VEVGeV, precisionBits), precisionBits)
	denom := mulFloat(newFloat(16, precisionBits), mulFloat(pi, pi, precisionBits), precisionBits)
	denom.Mul(denom, v2)
	raw := new(big.Float).SetPrec(precisionBits).SetInt64(0)
	comps := make([]LoopComponent, 0, len(inputs))
	for _, li := range inputs {
		var mass *big.Float
		var massRat *big.Rat
		if li.MassGeV == nil {
			mass = in.NativeMassGeV
			massRat = nil
		} else {
			mass = ratToFloat(li.MassGeV, precisionBits)
			massRat = li.MassGeV
		}
		m2 := mulFloat(mass, mass, precisionBits)
		m4 := mulFloat(m2, m2, precisionBits)
		num := mulFloat(newFloat(li.Coeff, precisionBits), m4, precisionBits)
		contrib := quoFloat(num, denom, precisionBits)
		raw.Add(raw, contrib)
		comps = append(comps, LoopComponent{Name: li.Name, Coeff: li.Coeff, MassGeV: massRat, NumeratorGeV4: num, ContributionGeV2: contrib})
	}
	return OneLoopKernel{Formula: "Π_poly=(−12m_t^4+6m_W^4+3m_Z^4+3m_H^4)/(16π²v²)", ScaleDescription: "high-precision one-loop component kernel; not a renormalized on-shell SM self-energy", Components: comps, RawKernelGeV2: raw, Status: StatusOneLoopKernelComputed}
}

func solveCounterterm(in Inputs, k OneLoopKernel) CountertermTarget {
	req := ratToFloat(in.RequiredRePiGeV2, precisionBits)
	rem := subFloat(req, k.RawKernelGeV2, precisionBits)
	rover := quoFloat(rem, req, precisionBits)
	rawover := quoFloat(k.RawKernelGeV2, req, precisionBits)
	return CountertermTarget{RequiredRePiGeV2: in.RequiredRePiGeV2, RawKernelGeV2: k.RawKernelGeV2, FiniteRemainderGeV2: rem, RemainderOverRequired: rover, RawOverRequired: rawover, Interpretation: "The raw polynomial kernel is not the pole correction. A finite renormalized self-energy/counterterm contribution of this size and sign is required to close the collider comparison in the chosen scheme.", Status: StatusFiniteCountertermTargetSolved}
}

func recommendRoute(routes RepairRoutes, c CountertermTarget) Recommendation {
	return Recommendation{BestRoute: "renormalized pole-correction branch", Reason: "It preserves the exact contact shape 1197/4624 and converts the remaining 43.604449567474 GeV² into a standard precision self-energy/counterterm calculation. Contact-shape fitting and VEV shifting are rejected as less principled.", NextGateTarget: "install the full renormalized SM Higgs one-loop pole equation with Passarino-Veltman coefficient table and scheme-specific finite counterterms", Status: StatusPoleCorrectionPreferred}
}

func preserveFirewalls() Firewalls {
	return Firewalls{NoFullRenormalizedSM: true, NoNativeCounterterms: true, NoShapeFit: true, NoColliderClaim: true, QuarantinedInputs: []string{"M_H reference=125.10 GeV", "v=246.22 GeV", "m_t=172.76 GeV", "m_W=80.379 GeV", "m_Z=91.1876 GeV"}, Statuses: []string{StatusFailedFullSMRenormalizedSelfEnergy, StatusFailedCountertermsNotDerived, StatusFailedContactShapeNotModified, StatusFailedColliderMassNotClaimed}}
}

func compileSummary(in Inputs, routes RepairRoutes, k OneLoopKernel, c CountertermTarget) Summary {
	contactDelta := routes.Routes[0].RequiredValue
	direct := "The better repair route is pole correction, not contact-shape deformation. The exact native mass proxy remains 125.274157149699 GeV; the precise pole target is ReΠ=43.604449567474 GeV². A raw one-loop polynomial kernel gives a large negative value and therefore proves that the full renormalized on-shell coefficient/counterterm table is mandatory."
	return Summary{NativeMassGeV: dec(in.NativeMassGeV, displayDigits), RequiredRePiGeV2: in.RequiredRePiGeV2.RatString(), RawKernelGeV2: dec(k.RawKernelGeV2, displayDigits), FiniteRemainderGeV2: dec(c.FiniteRemainderGeV2, displayDigits), ContactShapeDelta: contactDelta, DirectAnswer: direct, Status: StatusPrecisionLedgerCompiled}
}

func Statuses(a Analysis) []string {
	out := []string{a.Inputs.Status, a.Routes.Status, a.Kernel.Status, a.Counterterm.Status, a.Recommendation.Status, a.Summary.Status, StatusTensionRawKernelWrongSignMagnitude, StatusTensionContactShapeFitForbidden}
	out = append(out, a.Firewalls.Statuses...)
	return out
}

func FormatInputs(x Inputs) string {
	return fmt.Sprintf("gate=%d R=%s λ=%s nativeMass=%s ReΠ=%s status=%s", x.HighestInheritedGate, x.ContactShape.RatString(), x.LambdaH.RatString(), dec(x.NativeMassGeV, 48), x.RequiredRePiGeV2.RatString(), x.Status)
}
func FormatRoutes(x RepairRoutes) string {
	parts := make([]string, 0, len(x.Routes))
	for _, r := range x.Routes {
		parts = append(parts, fmt.Sprintf("%s value=%s allowed=%v", r.Name, r.RequiredValue, r.AllowedByNativeCore))
	}
	return fmt.Sprintf("routes=[%s] status=%s", strings.Join(parts, "; "), x.Status)
}
func FormatKernel(x OneLoopKernel) string {
	parts := make([]string, 0, len(x.Components))
	for _, c := range x.Components {
		parts = append(parts, fmt.Sprintf("%s:%s", c.Name, dec(c.ContributionGeV2, 24)))
	}
	return fmt.Sprintf("formula=%s raw=%s components=[%s] status=%s", x.Formula, dec(x.RawKernelGeV2, 36), strings.Join(parts, "; "), x.Status)
}
func FormatCounterterm(x CountertermTarget) string {
	return fmt.Sprintf("target=%s raw=%s remainder=%s rem/target=%s raw/target=%s status=%s", x.RequiredRePiGeV2.RatString(), dec(x.RawKernelGeV2, 36), dec(x.FiniteRemainderGeV2, 36), dec(x.RemainderOverRequired, 24), dec(x.RawOverRequired, 24), x.Status)
}
func FormatRecommendation(x Recommendation) string {
	return fmt.Sprintf("best=%s reason=%s next=%s status=%s", x.BestRoute, x.Reason, x.NextGateTarget, x.Status)
}
func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("noFullSM=%v noCT=%v noShapeFit=%v noPole=%v inputs=[%s] statuses=[%s]", x.NoFullRenormalizedSM, x.NoNativeCounterterms, x.NoShapeFit, x.NoColliderClaim, strings.Join(x.QuarantinedInputs, ", "), strings.Join(x.Statuses, "; "))
}
func FormatSummary(x Summary) string {
	return fmt.Sprintf("m=%s ReΠ=%s raw=%s remainder=%s ΔR=%s answer=%s status=%s", x.NativeMassGeV, x.RequiredRePiGeV2, x.RawKernelGeV2, x.FiniteRemainderGeV2, x.ContactShapeDelta, x.DirectAnswer, x.Status)
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
func sqrtRatMul(v *big.Rat, r *big.Rat) *big.Float {
	sqrt := new(big.Float).SetPrec(precisionBits).Sqrt(ratToFloat(r, precisionBits))
	return mulFloat(ratToFloat(v, precisionBits), sqrt, precisionBits)
}

func machinPi(prec uint) *big.Float {
	a := atanInvInt(5, prec)
	b := atanInvInt(239, prec)
	return subFloat(mulFloat(newFloat(16, prec), a, prec), mulFloat(newFloat(4, prec), b, prec), prec)
}
func atanInvInt(q int64, prec uint) *big.Float {
	x := quoFloat(one(prec), newFloat(q, prec), prec)
	x2 := mulFloat(x, x, prec)
	sum := new(big.Float).SetPrec(prec).SetInt64(0)
	power := new(big.Float).SetPrec(prec).Set(x)
	eps := new(big.Float).SetPrec(prec).SetFloat64(math.Ldexp(1, -int(prec)-32))
	sign := 1
	for n := int64(0); n < 10000; n++ {
		add := quoFloat(power, newFloat(2*n+1, prec), prec)
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
