// Package higgspassarinoveltmankernel implements Gate 334:
// Passarino-Veltman Higgs Pole Kernel / Finite One-Loop Integral Installation Audit.
//
// Gate 333 installed the sign/multiplicity component ledger and showed that a
// raw polynomial one-loop kernel is not a renormalized pole-mass prediction.
// Gate 334 installs the finite Passarino-Veltman basis functions needed by a
// real one-loop Higgs pole calculation. It evaluates A0 and B0 witnesses under
// a quarantined scale choice, but it does not claim an exact collider mass: the
// full Standard Model coefficient table, counterterms, gauge convention, and
// input scheme remain firewalled.
package higgspassarinoveltmankernel

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE334-HIGGS-PASSARINO-VELTMAN-POLE-KERNEL-FINITE-INTEGRAL-INSTALLATION-AUDIT"

	StatusGate333Inherited             = "CONDITIONAL_SUPPORT_GATE333_ONE_LOOP_COMPONENT_LEDGER_INHERITED"
	StatusPVBasisFormalized            = "CONDITIONAL_SUPPORT_PASSARINO_VELTMAN_BASIS_FORMALIZED"
	StatusFinitePVFunctionsComputed    = "CONDITIONAL_SUPPORT_FINITE_PV_FUNCTIONS_COMPUTED"
	StatusOnShellKernelSlotsInstalled  = "CONDITIONAL_SUPPORT_ON_SHELL_KERNEL_SLOTS_INSTALLED"
	StatusCoefficientFirewallPreserved = "CONDITIONAL_SUPPORT_COEFFICIENT_TABLE_FIREWALL_PRESERVED"
	StatusPoleClaimFirewallPreserved   = "CONDITIONAL_SUPPORT_POLE_MASS_FIREWALL_PRESERVED"

	StatusTensionPVFunctionsNotEnough     = "CONDITIONAL_TENSION_PV_FUNCTIONS_ALONE_DO_NOT_FIX_POLE_MASS"
	StatusTensionScaleSchemeQuarantined   = "CONDITIONAL_TENSION_RENORMALIZATION_SCALE_AND_SCHEME_QUARANTINED"
	StatusTensionThresholdBranchesMissing = "CONDITIONAL_TENSION_FULL_THRESHOLD_BRANCHES_NOT_INSTALLED"

	StatusFailedFullCoefficientTable = "FAILED_ROUTE_FULL_SM_ONE_LOOP_COEFFICIENT_TABLE_NOT_INSTALLED"
	StatusFailedCounterterms         = "FAILED_ROUTE_RENORMALIZED_COUNTERTERMS_NOT_DERIVED"
	StatusFailedGaugeScheme          = "FAILED_ROUTE_GAUGE_AND_INPUT_SCHEME_NOT_DERIVED"
	StatusFailedExactPoleMass        = "FAILED_ROUTE_EXACT_COLLIDER_HIGGS_MASS_NOT_CLAIMED"
)

const (
	inheritedHighestGate = 333

	contactScalarNumerator   = 1197.0
	contactScalarDenominator = 4624.0
	electroweakVEVGeV        = 246.22
	observedHiggsGeV         = 125.10

	// Quarantined conventional SM central values used only to evaluate the PV
	// basis. They are not derived from the finite algebra.
	nominalTopMassGeV = 172.76
	nominalWMassGeV   = 80.379
	nominalZMassGeV   = 91.1876
)

type Inputs struct {
	HighestInheritedGate int
	NativeRunMassGeV     float64
	ObservedPoleGeV      float64
	ExternalMassInputs   []string
	ScaleChoice          string
	MuGeV                float64
	Status               string
}

type PVBasis struct {
	A0FiniteDefinition string
	B0FiniteDefinition string
	EqualMassLane      bool
	BelowThresholdOnly bool
	Status             string
}

type PVValue struct {
	Particle       string
	MassGeV        float64
	SGeV2          float64
	MuGeV          float64
	A0FiniteGeV2   float64
	B0Finite       float64
	ThresholdRatio float64
	Regime         string
	Status         string
}

type PVLedger struct {
	Values []PVValue
	Status string
}

type KernelSlot struct {
	Name                    string
	RequiredPVBlocks        []string
	InstalledBasisAvailable bool
	FullCoefficientKnown    bool
	FiniteCountertermKnown  bool
	Interpretation          string
}

type KernelSlots struct {
	Slots             []KernelSlot
	AllBasisAvailable bool
	FullKernelClosed  bool
	Status            string
}

type Firewalls struct {
	NoCoefficientTable bool
	NoCounterterms     bool
	NoGaugeScheme      bool
	NoExactPoleClaim   bool
	Status             string
}

type Summary struct {
	PVBlockCount int
	KernelClosed bool
	DirectAnswer string
	NextGate     string
	Status       string
}

type Analysis struct {
	Inputs    Inputs
	Basis     PVBasis
	PV        PVLedger
	Slots     KernelSlots
	Firewalls Firewalls
	Summary   Summary
	Truth     string
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
	inputs := compileInputs()
	basis := formalizePVBasis()
	pv := computePVLedger(inputs)
	slots := installKernelSlots(pv)
	firewalls := preserveFirewalls()
	summary := compileSummary(pv, slots)
	truth := "Gate 334 installs the finite Passarino-Veltman basis required for a one-loop Higgs pole-mass calculation. The A0 and B0 blocks are computable and finite under an explicit scale choice, but they are only basis functions. A physical pole mass still requires the full SM coefficient table, gauge/input scheme, and renormalized finite counterterms."
	return Analysis{Inputs: inputs, Basis: basis, PV: pv, Slots: slots, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func compileInputs() Inputs {
	native := electroweakVEVGeV * math.Sqrt(contactScalarNumerator/contactScalarDenominator)
	return Inputs{
		HighestInheritedGate: inheritedHighestGate,
		NativeRunMassGeV:     native,
		ObservedPoleGeV:      observedHiggsGeV,
		ExternalMassInputs:   []string{"m_t=172.76 GeV", "m_W=80.379 GeV", "m_Z=91.1876 GeV"},
		ScaleChoice:          "quarantined μ=m_native proxy; not a native renormalization theorem",
		MuGeV:                native,
		Status:               StatusGate333Inherited,
	}
}

func formalizePVBasis() PVBasis {
	return PVBasis{
		A0FiniteDefinition: "A0_fin(m²;μ²)=m²[1-ln(m²/μ²)] after divergent piece is subtracted",
		B0FiniteDefinition: "B0_fin(s;m²,m²;μ²)=-ln(m²/μ²)+2-2√(4m²/s-1) atan(1/√(4m²/s-1)) for s<4m²",
		EqualMassLane:      true,
		BelowThresholdOnly: true,
		Status:             StatusPVBasisFormalized,
	}
}

func computePVLedger(in Inputs) PVLedger {
	s := in.ObservedPoleGeV * in.ObservedPoleGeV
	values := []PVValue{
		computePVValue("top", nominalTopMassGeV, s, in.MuGeV),
		computePVValue("W", nominalWMassGeV, s, in.MuGeV),
		computePVValue("Z", nominalZMassGeV, s, in.MuGeV),
		computePVValue("H", in.NativeRunMassGeV, s, in.MuGeV),
	}
	return PVLedger{Values: values, Status: StatusFinitePVFunctionsComputed}
}

func computePVValue(name string, mass, s, mu float64) PVValue {
	m2 := mass * mass
	mu2 := mu * mu
	ratio := 4.0 * m2 / s
	regime := "below pair threshold; real equal-mass B0 branch"
	if ratio <= 1.0 {
		regime = "above threshold; complex branch required and not installed"
	}
	a0 := m2 * (1.0 - math.Log(m2/mu2))
	b0 := math.NaN()
	if ratio > 1.0 {
		x := math.Sqrt(ratio - 1.0)
		b0 = -math.Log(m2/mu2) + 2.0 - 2.0*x*math.Atan(1.0/x)
	}
	return PVValue{Particle: name, MassGeV: mass, SGeV2: s, MuGeV: mu, A0FiniteGeV2: a0, B0Finite: b0, ThresholdRatio: ratio, Regime: regime, Status: StatusFinitePVFunctionsComputed}
}

func installKernelSlots(pv PVLedger) KernelSlots {
	basisAvailable := len(pv.Values) == 4
	slots := []KernelSlot{
		{Name: "fermion/top contribution", RequiredPVBlocks: []string{"A0(m_t)", "B0(s;m_t,m_t)"}, InstalledBasisAvailable: basisAvailable, FullCoefficientKnown: false, FiniteCountertermKnown: false, Interpretation: "dominant Yukawa loop requires exact on-shell/MS-bar coefficient and counterterm"},
		{Name: "charged gauge/W contribution", RequiredPVBlocks: []string{"A0(m_W)", "B0(s;m_W,m_W)"}, InstalledBasisAvailable: basisAvailable, FullCoefficientKnown: false, FiniteCountertermKnown: false, Interpretation: "gauge-loop coefficient depends on gauge/input scheme"},
		{Name: "neutral gauge/Z contribution", RequiredPVBlocks: []string{"A0(m_Z)", "B0(s;m_Z,m_Z)"}, InstalledBasisAvailable: basisAvailable, FullCoefficientKnown: false, FiniteCountertermKnown: false, Interpretation: "neutral-gauge loop requires scheme-fixed coefficient"},
		{Name: "scalar/H contribution", RequiredPVBlocks: []string{"A0(m_H)", "B0(s;m_H,m_H)"}, InstalledBasisAvailable: basisAvailable, FullCoefficientKnown: false, FiniteCountertermKnown: false, Interpretation: "scalar self-energy requires renormalized λ and field-strength counterterms"},
	}
	return KernelSlots{Slots: slots, AllBasisAvailable: basisAvailable, FullKernelClosed: false, Status: StatusOnShellKernelSlotsInstalled}
}

func preserveFirewalls() Firewalls {
	return Firewalls{NoCoefficientTable: true, NoCounterterms: true, NoGaugeScheme: true, NoExactPoleClaim: true, Status: StatusPoleClaimFirewallPreserved}
}

func compileSummary(pv PVLedger, slots KernelSlots) Summary {
	direct := "Finite A0/B0 Passarino-Veltman building blocks are now installed for the top, W, Z, and Higgs equal-mass below-threshold lanes, but the physical one-loop pole kernel is not closed."
	next := "Install the renormalized SM Higgs self-energy coefficient table and counterterm scheme that contracts these PV basis functions into ReΠ_HH(p²)."
	return Summary{PVBlockCount: len(pv.Values), KernelClosed: slots.FullKernelClosed, DirectAnswer: direct, NextGate: next, Status: StatusCoefficientFirewallPreserved}
}

func Statuses(a Analysis) []string {
	return []string{
		a.Inputs.Status,
		a.Basis.Status,
		a.PV.Status,
		a.Slots.Status,
		a.Firewalls.Status,
		a.Summary.Status,
		StatusTensionPVFunctionsNotEnough,
		StatusTensionScaleSchemeQuarantined,
		StatusTensionThresholdBranchesMissing,
		StatusFailedFullCoefficientTable,
		StatusFailedCounterterms,
		StatusFailedGaugeScheme,
		StatusFailedExactPoleMass,
	}
}

func FormatInputs(x Inputs) string {
	return fmt.Sprintf("gate=%d m_native=%.12f observed=%.12f μ=%.12f scale=%s external=[%s] status=%s", x.HighestInheritedGate, x.NativeRunMassGeV, x.ObservedPoleGeV, x.MuGeV, x.ScaleChoice, strings.Join(x.ExternalMassInputs, ", "), x.Status)
}
func FormatBasis(x PVBasis) string {
	return fmt.Sprintf("A0=%s B0=%s equalMass=%v belowThreshold=%v status=%s", x.A0FiniteDefinition, x.B0FiniteDefinition, x.EqualMassLane, x.BelowThresholdOnly, x.Status)
}
func FormatPVValue(x PVValue) string {
	return fmt.Sprintf("%s mass=%.6f s=%.6f μ=%.6f A0_fin=%+.12f B0_fin=%+.12f 4m²/s=%.9f regime=%s status=%s", x.Particle, x.MassGeV, x.SGeV2, x.MuGeV, x.A0FiniteGeV2, x.B0Finite, x.ThresholdRatio, x.Regime, x.Status)
}
func FormatPVLedger(x PVLedger) string {
	parts := make([]string, 0, len(x.Values))
	for _, v := range x.Values {
		parts = append(parts, FormatPVValue(v))
	}
	return fmt.Sprintf("values=[%s] status=%s", strings.Join(parts, " | "), x.Status)
}
func FormatKernelSlot(x KernelSlot) string {
	return fmt.Sprintf("%s blocks=[%s] basis=%v coeff=%v CT=%v role=%s", x.Name, strings.Join(x.RequiredPVBlocks, ","), x.InstalledBasisAvailable, x.FullCoefficientKnown, x.FiniteCountertermKnown, x.Interpretation)
}
func FormatSlots(x KernelSlots) string {
	parts := make([]string, 0, len(x.Slots))
	for _, s := range x.Slots {
		parts = append(parts, FormatKernelSlot(s))
	}
	return fmt.Sprintf("allBasis=%v fullKernel=%v slots=[%s] status=%s", x.AllBasisAvailable, x.FullKernelClosed, strings.Join(parts, " | "), x.Status)
}
func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("noCoeff=%v noCT=%v noGaugeScheme=%v noExactPole=%v status=%s", x.NoCoefficientTable, x.NoCounterterms, x.NoGaugeScheme, x.NoExactPoleClaim, x.Status)
}
func FormatSummary(x Summary) string {
	return fmt.Sprintf("PVBlocks=%d kernelClosed=%v answer=%s next=%s status=%s", x.PVBlockCount, x.KernelClosed, x.DirectAnswer, x.NextGate, x.Status)
}
func FormatStatuses(ss []string) string  { return "statuses=" + strings.Join(ss, "; ") }
func nearlyEqual(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
