// Package higgsoneloopselfenergyledger implements Gate 333:
// One-Loop Higgs Self-Energy Component Ledger / Renormalized Pole Kernel Audit.
//
// Gate 332 converted the native 125.274 GeV tree/running Higgs proxy into the
// finite self-energy target required by a pole-mass conversion. Gate 333 installs
// the actual one-loop Standard-Model component ledger at the level of signs,
// multiplicities, natural magnitudes, and renormalization obligations. It does
// not claim an exact collider pole mass: a complete on-shell/MS-bar conversion
// requires Passarino-Veltman functions, counterterms, scale choice, and input
// scheme selection.
package higgsoneloopselfenergyledger

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE333-HIGGS-ONE-LOOP-SELF-ENERGY-COMPONENT-LEDGER-RENORMALIZED-POLE-KERNEL-AUDIT"

	StatusGate332Inherited            = "CONDITIONAL_SUPPORT_GATE332_SELF_ENERGY_TARGET_INHERITED"
	StatusComponentLedgerFormalized   = "CONDITIONAL_SUPPORT_ONE_LOOP_COMPONENT_LEDGER_FORMALIZED"
	StatusVeltmanKernelAudited        = "CONDITIONAL_SUPPORT_UNRENORMALIZED_ONE_LOOP_KERNEL_AUDITED"
	StatusCountertermLedgerFormalized = "CONDITIONAL_SUPPORT_RENORMALIZED_COUNTERTERM_LEDGER_FORMALIZED"
	StatusSchemeDependencyFormalized  = "CONDITIONAL_SUPPORT_POLE_SCHEME_DEPENDENCY_FORMALIZED"
	StatusPrecisionFirewallPreserved  = "CONDITIONAL_SUPPORT_PRECISION_FIREWALL_PRESERVED"

	StatusTensionRawKernelNotTarget   = "CONDITIONAL_TENSION_RAW_ONE_LOOP_KERNEL_NOT_EQUAL_TO_POLE_TARGET"
	StatusTensionCountertermMandatory = "CONDITIONAL_TENSION_FINITE_COUNTERTERM_AND_SCHEME_CHOICE_MANDATORY"

	StatusFailedPassarinoVeltmanNotComputed = "FAILED_ROUTE_PASSARINO_VELTMAN_FUNCTIONS_NOT_COMPUTED"
	StatusFailedCountertermsNotDerived      = "FAILED_ROUTE_FINITE_COUNTERTERMS_NOT_DERIVED"
	StatusFailedInputSchemeNotNative        = "FAILED_ROUTE_SM_INPUT_SCHEME_NOT_NATIVE"
	StatusFailedExactPoleMassNotClaimed     = "FAILED_ROUTE_EXACT_COLLIDER_HIGGS_MASS_NOT_CLAIMED"
)

const (
	inheritedHighestGate = 332

	contactScalarNumerator   = 1197.0
	contactScalarDenominator = 4624.0
	gStarSquaredNative       = 0.5

	electroweakVEVGeV = 246.22
	observedHiggsGeV  = 125.10

	// Quarantined conventional SM central values used only for one-loop capacity
	// auditing. They are not derived by the finite algebra and are not used to
	// claim a precision collider mass.
	nominalTopMassGeV = 172.76
	nominalWMassGeV   = 80.379
	nominalZMassGeV   = 91.1876
)

type Inputs struct {
	HighestInheritedGate int
	LambdaH              float64
	NativeProxyMassGeV   float64
	ObservedPoleGeV      float64
	RequiredRePiGeV2     float64
	Status               string
}

type Component struct {
	Name             string
	Multiplicity     float64
	Sign             float64
	MassGeV          float64
	ContributionGeV2 float64
	Interpretation   string
}

type ComponentLedger struct {
	Formula    string
	Components []Component
	Status     string
}

type KernelAudit struct {
	RawKernelGeV2                float64
	RequiredRePiGeV2             float64
	DifferenceRawMinusTargetGeV2 float64
	RawOverRequired              float64
	MatchesTarget                bool
	Status                       string
}

type CountertermLedger struct {
	RequiredFiniteCountertermGeV2 float64
	CountertermOverRawAbs         float64
	CountertermOverTarget         float64
	CountertermMandatory          bool
	Interpretation                string
	Status                        string
}

type SchemeLedger struct {
	NeedsGaugeChoice          bool
	NeedsRenormalizationScale bool
	NeedsMassInputScheme      bool
	NeedsPVIntegrals          bool
	NeedsFiniteCounterterms   bool
	ExactPoleMassComputed     bool
	Status                    string
}

type FirewallAudit struct {
	NoExactPVFunctions      bool
	NoCountertermDerivation bool
	NoExactColliderClaim    bool
	NoNativeSMInputClaim    bool
	Status                  string
}

type Summary struct {
	RequiredRePiGeV2              float64
	RawKernelGeV2                 float64
	RequiredFiniteCountertermGeV2 float64
	DirectAnswer                  string
	NextObligation                string
	Status                        string
}

type Analysis struct {
	Inputs       Inputs
	Ledger       ComponentLedger
	Kernel       KernelAudit
	Counterterms CountertermLedger
	Scheme       SchemeLedger
	Audit        FirewallAudit
	Summary      Summary
	Truth        string
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
	ledger := formalizeComponentLedger(inputs)
	kernel := auditRawKernel(inputs, ledger)
	counterterms := formalizeCountertermLedger(kernel)
	scheme := formalizeSchemeLedger()
	audit := auditFirewalls()
	summary := compileSummary(inputs, kernel, counterterms)
	truth := "Gate 333 installs the one-loop Higgs self-energy component ledger. The raw Veltman/Coleman-Weinberg-like kernel has the expected large top-dominated negative sign and is not the finite pole target from Gate 332. Therefore the remaining +43.604 GeV² pole conversion cannot be read directly from an unrenormalized component sum; the finite on-shell/MS-bar counterterm, scale, and Passarino-Veltman kernel must be fixed before an exact collider Higgs pole mass is claimed."
	return Analysis{Inputs: inputs, Ledger: ledger, Kernel: kernel, Counterterms: counterterms, Scheme: scheme, Audit: audit, Summary: summary, Truth: truth}, nil
}

func compileInputs() Inputs {
	lambda := (contactScalarNumerator / contactScalarDenominator) * gStarSquaredNative
	nativeMass := electroweakVEVGeV * math.Sqrt(contactScalarNumerator/contactScalarDenominator)
	run2 := nativeMass * nativeMass
	pole2 := observedHiggsGeV * observedHiggsGeV
	requiredPi := run2 - pole2
	return Inputs{HighestInheritedGate: inheritedHighestGate, LambdaH: lambda, NativeProxyMassGeV: nativeMass, ObservedPoleGeV: observedHiggsGeV, RequiredRePiGeV2: requiredPi, Status: StatusGate332Inherited}
}

func formalizeComponentLedger(in Inputs) ComponentLedger {
	// This is a capacity kernel, not a full pole self-energy. It records the
	// standard relative one-loop signs/multiplicities in the Veltman/CW mass^4
	// structure: (1/16π²v²)[-12m_t^4 + 6m_W^4 + 3m_Z^4 + 3m_H^4].
	comps := []Component{
		makeComponent("top-quark fermion loop", -12, nominalTopMassGeV, "dominant negative fermionic contribution"),
		makeComponent("W-boson loop", +6, nominalWMassGeV, "positive charged gauge contribution"),
		makeComponent("Z-boson loop", +3, nominalZMassGeV, "positive neutral gauge contribution"),
		makeComponent("Higgs/scalar loop", +3, in.NativeProxyMassGeV, "positive scalar self-interaction contribution"),
	}
	return ComponentLedger{Formula: "Π_kernel≈(1/16π²v²)[-12m_t⁴+6m_W⁴+3m_Z⁴+3m_H⁴] before PV functions/counterterms", Components: comps, Status: StatusComponentLedgerFormalized}
}

func makeComponent(name string, signedMultiplicity float64, mass float64, interpretation string) Component {
	factor := math.Pow(mass, 4) / (16.0 * math.Pi * math.Pi * electroweakVEVGeV * electroweakVEVGeV)
	sign := 1.0
	if signedMultiplicity < 0 {
		sign = -1.0
	}
	return Component{Name: name, Multiplicity: math.Abs(signedMultiplicity), Sign: sign, MassGeV: mass, ContributionGeV2: signedMultiplicity * factor, Interpretation: interpretation}
}

func auditRawKernel(in Inputs, ledger ComponentLedger) KernelAudit {
	raw := 0.0
	for _, c := range ledger.Components {
		raw += c.ContributionGeV2
	}
	diff := raw - in.RequiredRePiGeV2
	ratio := raw / in.RequiredRePiGeV2
	matches := math.Abs(diff) < 1.0
	return KernelAudit{RawKernelGeV2: raw, RequiredRePiGeV2: in.RequiredRePiGeV2, DifferenceRawMinusTargetGeV2: diff, RawOverRequired: ratio, MatchesTarget: matches, Status: StatusVeltmanKernelAudited}
}

func formalizeCountertermLedger(k KernelAudit) CountertermLedger {
	requiredCT := k.RequiredRePiGeV2 - k.RawKernelGeV2
	interp := "A finite renormalized pole scheme must transform the raw top-dominated one-loop kernel into the small Gate-332 pole target. This is normal QFT bookkeeping, but the counterterm is not derived here."
	return CountertermLedger{RequiredFiniteCountertermGeV2: requiredCT, CountertermOverRawAbs: requiredCT / math.Abs(k.RawKernelGeV2), CountertermOverTarget: requiredCT / k.RequiredRePiGeV2, CountertermMandatory: math.Abs(requiredCT) > 100, Interpretation: interp, Status: StatusCountertermLedgerFormalized}
}

func formalizeSchemeLedger() SchemeLedger {
	return SchemeLedger{NeedsGaugeChoice: true, NeedsRenormalizationScale: true, NeedsMassInputScheme: true, NeedsPVIntegrals: true, NeedsFiniteCounterterms: true, ExactPoleMassComputed: false, Status: StatusSchemeDependencyFormalized}
}

func auditFirewalls() FirewallAudit {
	return FirewallAudit{NoExactPVFunctions: true, NoCountertermDerivation: true, NoExactColliderClaim: true, NoNativeSMInputClaim: true, Status: StatusPrecisionFirewallPreserved}
}

func compileSummary(in Inputs, k KernelAudit, c CountertermLedger) Summary {
	direct := "The one-loop component ledger is installed, but the raw top-dominated kernel is not the pole target; an explicit renormalized self-energy calculation is still required."
	next := "Compute the finite one-loop Passarino-Veltman Higgs self-energy in a chosen on-shell/MS-bar scheme and verify whether the required +43.604 GeV² target emerges without fitting."
	return Summary{RequiredRePiGeV2: in.RequiredRePiGeV2, RawKernelGeV2: k.RawKernelGeV2, RequiredFiniteCountertermGeV2: c.RequiredFiniteCountertermGeV2, DirectAnswer: direct, NextObligation: next, Status: StatusCountertermLedgerFormalized}
}

func Statuses(a Analysis) []string {
	return []string{
		a.Inputs.Status,
		a.Ledger.Status,
		a.Kernel.Status,
		a.Counterterms.Status,
		a.Scheme.Status,
		a.Audit.Status,
		StatusTensionRawKernelNotTarget,
		StatusTensionCountertermMandatory,
		StatusFailedPassarinoVeltmanNotComputed,
		StatusFailedCountertermsNotDerived,
		StatusFailedInputSchemeNotNative,
		StatusFailedExactPoleMassNotClaimed,
	}
}

func FormatInputs(x Inputs) string {
	return fmt.Sprintf("gate=%d λ=%.15f m_native=%.12f observed=%.12f ReΠ_target=%.12f status=%s", x.HighestInheritedGate, x.LambdaH, x.NativeProxyMassGeV, x.ObservedPoleGeV, x.RequiredRePiGeV2, x.Status)
}
func FormatComponent(x Component) string {
	return fmt.Sprintf("%s sign=%+.0f multiplicity=%.0f mass=%.6f contribution=%+.12f role=%s", x.Name, x.Sign, x.Multiplicity, x.MassGeV, x.ContributionGeV2, x.Interpretation)
}
func FormatLedger(x ComponentLedger) string {
	parts := make([]string, 0, len(x.Components))
	for _, c := range x.Components {
		parts = append(parts, FormatComponent(c))
	}
	return fmt.Sprintf("formula=%s components=[%s] status=%s", x.Formula, strings.Join(parts, " | "), x.Status)
}
func FormatKernel(x KernelAudit) string {
	return fmt.Sprintf("raw=%+.12f target=%+.12f diff=%+.12f raw/target=%+.9f matches=%v status=%s", x.RawKernelGeV2, x.RequiredRePiGeV2, x.DifferenceRawMinusTargetGeV2, x.RawOverRequired, x.MatchesTarget, x.Status)
}
func FormatCounterterms(x CountertermLedger) string {
	return fmt.Sprintf("CT_required=%+.12f CT/|raw|=%.9f CT/target=%.9f mandatory=%v status=%s reason=%s", x.RequiredFiniteCountertermGeV2, x.CountertermOverRawAbs, x.CountertermOverTarget, x.CountertermMandatory, x.Status, x.Interpretation)
}
func FormatScheme(x SchemeLedger) string {
	return fmt.Sprintf("gauge=%v scale=%v inputScheme=%v PV=%v counterterms=%v exactPole=%v status=%s", x.NeedsGaugeChoice, x.NeedsRenormalizationScale, x.NeedsMassInputScheme, x.NeedsPVIntegrals, x.NeedsFiniteCounterterms, x.ExactPoleMassComputed, x.Status)
}
func FormatAudit(x FirewallAudit) string {
	return fmt.Sprintf("noPV=%v noCTDerivation=%v noExactPole=%v noNativeInputs=%v status=%s", x.NoExactPVFunctions, x.NoCountertermDerivation, x.NoExactColliderClaim, x.NoNativeSMInputClaim, x.Status)
}
func FormatSummary(x Summary) string {
	return fmt.Sprintf("target=%.12f raw=%+.12f CT=%+.12f answer=%s next=%s status=%s", x.RequiredRePiGeV2, x.RawKernelGeV2, x.RequiredFiniteCountertermGeV2, x.DirectAnswer, x.NextObligation, x.Status)
}
func FormatStatuses(ss []string) string  { return "statuses=" + strings.Join(ss, "; ") }
func nearlyEqual(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
