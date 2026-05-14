// Package higgsonshellrenormalizationscheme implements Gate 338:
// On-Shell Renormalization Scheme / Passarino-Veltman Pole Matching Audit.
//
// Gate 337 isolated the exact pole precision target and rejected contact-shape
// fitting. Gate 338 installs the formal one-loop QFT machinery needed to turn
// the finite Passarino-Veltman basis into a renormalized Higgs pole mass:
// A0/B0 blocks, an on-shell counterterm ledger, the MS-bar comparison lane, and
// a geometric-alignment audit. It does not claim the final collider pole mass;
// the full Standard Model coefficient table and finite counterterms remain a
// scheme-dependent continuum calculation.
package higgsonshellrenormalizationscheme

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE338-ON-SHELL-RENORMALIZATION-SCHEME-PASSARINO-VELTMAN-POLE-MATCHING-AUDIT"

	StatusGate337Inherited             = "CONDITIONAL_SUPPORT_GATE337_PRECISION_ROUTE_SIEVE_INHERITED"
	StatusPVStructureFormalized        = "CONDITIONAL_SUPPORT_PASSARINO_VELTMAN_STRUCTURE_FORMALIZED"
	StatusPVFiniteBlocksComputed       = "CONDITIONAL_SUPPORT_FINITE_PV_BLOCKS_COMPUTED"
	StatusRenormalizationSchemeAudited = "CONDITIONAL_SUPPORT_RENORMALIZATION_SCHEME_DEPENDENCY_AUDITED"
	StatusCountertermTargetMapped      = "CONDITIONAL_SUPPORT_COUNTERTERM_TARGET_MAPPED_TO_ON_SHELL_LEDGER"
	StatusGeometricAlignmentAudited    = "CONDITIONAL_SUPPORT_GEOMETRIC_ALIGNMENT_AUDITED"
	StatusPrecisionFirewallsPreserved  = "CONDITIONAL_SUPPORT_PRECISION_FIREWALLS_PRESERVED"

	StatusTensionSchemeNotNative     = "CONDITIONAL_TENSION_IR_RENORMALIZATION_SCHEME_NOT_SELECTED_BY_FINITE_CORE"
	StatusTensionCountertermRequired = "CONDITIONAL_TENSION_FINITE_COUNTERTERM_REQUIRED_TO_REACH_POLE_TARGET"
	StatusTensionPVBlocksNotEnough   = "CONDITIONAL_TENSION_PV_BASIS_WITHOUT_COEFFICIENTS_DOES_NOT_CLOSE_POLE_MASS"

	StatusFailedFullSMCoefficientTable = "FAILED_ROUTE_FULL_SM_ONE_LOOP_COEFFICIENT_TABLE_NOT_INSTALLED"
	StatusFailedCountertermsDerived    = "FAILED_ROUTE_RENORMALIZED_COUNTERTERMS_NOT_DERIVED_FROM_NATIVE_SCHEME"
	StatusFailedGaugeInputScheme       = "FAILED_ROUTE_GAUGE_INPUT_SCHEME_NOT_DERIVED"
	StatusFailedNativeIRScheme         = "FAILED_ROUTE_ASHA_BOUNDARY_DOES_NOT_SELECT_IR_RENORMALIZATION_SCHEME"
	StatusFailedExactPoleMassClaim     = "FAILED_ROUTE_EXACT_COLLIDER_HIGGS_MASS_NOT_CLAIMED"
)

const (
	inheritedHighestGate = 337

	contactScalarNumerator   = 1197.0
	contactScalarDenominator = 4624.0
	electroweakVEVGeV        = 246.22
	observedHiggsGeV         = 125.10

	// Quarantined conventional pole inputs used only for the continuum precision
	// ledger. They are not derived by the finite Cℓ(1,7) core.
	nominalTopMassGeV = 172.76
	nominalWMassGeV   = 80.379
	nominalZMassGeV   = 91.1876
)

type Inputs struct {
	HighestInheritedGate int
	NativeMassGeV        float64
	ObservedPoleGeV      float64
	LambdaNative         float64
	RequiredRePiGeV2     float64
	RawPolynomialGeV2    float64
	FiniteRemainderGeV2  float64
	MuGeV                float64
	QuarantinedInputs    []string
	Status               string
}

type PVDefinition struct {
	Name       string
	Definition string
	Role       string
}

type PVStructure struct {
	Definitions  []PVDefinition
	SelfEnergy   string
	PoleEquation string
	Status       string
}

type PVBlock struct {
	Particle       string
	MassGeV        float64
	SGeV2          float64
	MuGeV          float64
	A0FiniteGeV2   float64
	B0Finite       float64
	ThresholdRatio float64
	RealBranch     bool
}

type PVBlockLedger struct {
	Blocks []PVBlock
	Status string
}

type SchemeLane struct {
	Name            string
	MassCondition   string
	CountertermRule string
	ScaleDependence string
	NativeSelected  bool
	CanHitTarget    bool
	Interpretation  string
}

type SchemeAudit struct {
	Lanes          []SchemeLane
	ChosenForAudit string
	Status         string
}

type CountertermMap struct {
	PoleEquation            string
	RequiredRePiGeV2        float64
	RawPolynomialGeV2       float64
	RequiredFiniteRemainder float64
	RemainderOverTarget     float64
	RemainderOverRawAbs     float64
	Interpretation          string
	Status                  string
}

type GeometricAlignment struct {
	UVBoundaryFixed        bool
	ContactShapeImmutable  bool
	F0CutoffFixed          bool
	IRSchemeSelected       bool
	NativeCountertermFound bool
	Interpretation         string
	Status                 string
}

type Firewalls struct {
	NoFullSMCoefficientTable bool
	NoNativeCounterterms     bool
	NoGaugeInputScheme       bool
	NoNativeIRScheme         bool
	NoExactPoleMassClaim     bool
	Status                   string
}

type Summary struct {
	RequiredRePiGeV2        float64
	RequiredFiniteRemainder float64
	PVBlocks                int
	DirectAnswer            string
	NextGate                string
	Status                  string
}

type Analysis struct {
	Inputs      Inputs
	PV          PVStructure
	Blocks      PVBlockLedger
	Schemes     SchemeAudit
	Counterterm CountertermMap
	Alignment   GeometricAlignment
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
	inputs := compileInputs()
	pv := formalizePVStructure()
	blocks := computePVBlocks(inputs)
	schemes := auditSchemes(inputs)
	counter := mapCountertermTarget(inputs)
	alignment := auditGeometricAlignment(inputs, schemes)
	firewalls := preserveFirewalls()
	summary := compileSummary(inputs, blocks, counter)
	truth := "Gate 338 installs the on-shell/MS-bar renormalization scheme ledger for the Higgs pole precision gap. The Passarino-Veltman A0/B0 structures and finite block values are explicit, and the required +1035.171 GeV² finite remainder is mapped to the on-shell counterterm ledger. The ASHA UV contact boundary remains immutable, but it does not itself select the IR renormalization scheme or derive the finite SM counterterms; therefore the exact collider pole mass is still firewalled."
	return Analysis{Inputs: inputs, PV: pv, Blocks: blocks, Schemes: schemes, Counterterm: counter, Alignment: alignment, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func compileInputs() Inputs {
	lambda := (contactScalarNumerator / contactScalarDenominator) * 0.5
	native := electroweakVEVGeV * math.Sqrt(contactScalarNumerator/contactScalarDenominator)
	observed := observedHiggsGeV
	req := native*native - observed*observed
	raw := (-12*math.Pow(nominalTopMassGeV, 4) + 6*math.Pow(nominalWMassGeV, 4) + 3*math.Pow(nominalZMassGeV, 4) + 3*math.Pow(native, 4)) / (16 * math.Pi * math.Pi * electroweakVEVGeV * electroweakVEVGeV)
	rem := req - raw
	return Inputs{
		HighestInheritedGate: inheritedHighestGate,
		NativeMassGeV:        native,
		ObservedPoleGeV:      observed,
		LambdaNative:         lambda,
		RequiredRePiGeV2:     req,
		RawPolynomialGeV2:    raw,
		FiniteRemainderGeV2:  rem,
		MuGeV:                native,
		QuarantinedInputs: []string{
			"m_t=172.76 GeV",
			"m_W=80.379 GeV",
			"m_Z=91.1876 GeV",
			"M_H reference=125.10 GeV",
			"μ=m_native for finite PV witness",
		},
		Status: StatusGate337Inherited,
	}
}

func formalizePVStructure() PVStructure {
	defs := []PVDefinition{
		{Name: "A0", Definition: "A0_fin(m²;μ²)=m²[1-ln(m²/μ²)] after Δ_UV subtraction", Role: "one-point tadpole scalar integral"},
		{Name: "B0", Definition: "B0_fin(s;m²,m²;μ²)=-ln(m²/μ²)+2-2√(4m²/s-1)atan(1/√(4m²/s-1)) for s<4m²", Role: "two-point scalar integral on real below-threshold branch"},
		{Name: "Π_HH", Definition: "ReΠ_HH(p²)=Σ_i C_i(p²,m_i,g,y,λ)·{A0_i,B0_i}+δM_H²+(p²-M_H²)δZ_H+...", Role: "renormalized Higgs self-energy contraction"},
	}
	return PVStructure{Definitions: defs, SelfEnergy: "Π_HH = Π_top + Π_W + Π_Z + Π_H + counterterms", PoleEquation: "M_H² - m_run² + ReΠ_HH(M_H²)=0", Status: StatusPVStructureFormalized}
}

func computePVBlocks(in Inputs) PVBlockLedger {
	s := in.ObservedPoleGeV * in.ObservedPoleGeV
	blocks := []PVBlock{
		computePVBlock("top", nominalTopMassGeV, s, in.MuGeV),
		computePVBlock("W", nominalWMassGeV, s, in.MuGeV),
		computePVBlock("Z", nominalZMassGeV, s, in.MuGeV),
		computePVBlock("H-native", in.NativeMassGeV, s, in.MuGeV),
	}
	return PVBlockLedger{Blocks: blocks, Status: StatusPVFiniteBlocksComputed}
}

func computePVBlock(name string, mass, s, mu float64) PVBlock {
	m2 := mass * mass
	mu2 := mu * mu
	ratio := 4 * m2 / s
	a0 := m2 * (1 - math.Log(m2/mu2))
	b0 := math.NaN()
	realBranch := ratio > 1
	if realBranch {
		x := math.Sqrt(ratio - 1)
		b0 = -math.Log(m2/mu2) + 2 - 2*x*math.Atan(1/x)
	}
	return PVBlock{Particle: name, MassGeV: mass, SGeV2: s, MuGeV: mu, A0FiniteGeV2: a0, B0Finite: b0, ThresholdRatio: ratio, RealBranch: realBranch}
}

func auditSchemes(in Inputs) SchemeAudit {
	lanes := []SchemeLane{
		{Name: "On-Shell", MassCondition: "M_H is defined by the real pole of the propagator", CountertermRule: "δM_H² cancels the chosen finite part so M_H²-m_run²+ReΠ_HH(M_H²)=0", ScaleDependence: "physical pole is μ-independent after all finite terms are included", NativeSelected: false, CanHitTarget: true, Interpretation: "Correct language for collider pole comparison, but finite counterterms require a full SM input scheme."},
		{Name: "MS-bar", MassCondition: "λ(μ), v(μ), y_t(μ), g_i(μ) define running masses", CountertermRule: "subtract only UV poles plus convention constants; finite residue remains μ-dependent", ScaleDependence: "explicit μ dependence remains until RG-improved matching is performed", NativeSelected: false, CanHitTarget: true, Interpretation: "Natural for UV-to-IR transport, but not itself the collider pole mass."},
		{Name: "Native ASHA contact boundary", MassCondition: "m_native²=v²·1197/4624", CountertermRule: "no finite IR counterterm theorem installed", ScaleDependence: "UV contact data fix the boundary ratio, not the IR renormalization prescription", NativeSelected: true, CanHitTarget: false, Interpretation: "Preserves the geometry but does not replace Standard Model pole matching."},
	}
	return SchemeAudit{Lanes: lanes, ChosenForAudit: "On-Shell pole ledger for comparison; not claimed native", Status: StatusRenormalizationSchemeAudited}
}

func mapCountertermTarget(in Inputs) CountertermMap {
	return CountertermMap{
		PoleEquation:            "M_H² - m_run² + ReΠ_HH(M_H²)=0",
		RequiredRePiGeV2:        in.RequiredRePiGeV2,
		RawPolynomialGeV2:       in.RawPolynomialGeV2,
		RequiredFiniteRemainder: in.FiniteRemainderGeV2,
		RemainderOverTarget:     in.FiniteRemainderGeV2 / in.RequiredRePiGeV2,
		RemainderOverRawAbs:     in.FiniteRemainderGeV2 / math.Abs(in.RawPolynomialGeV2),
		Interpretation:          "The raw polynomial component is not the pole correction. In an on-shell ledger, a finite scheme-dependent residue/counterterm of this size must convert the raw PV contraction into the +43.604 GeV² target.",
		Status:                  StatusCountertermTargetMapped,
	}
}

func auditGeometricAlignment(in Inputs, schemes SchemeAudit) GeometricAlignment {
	return GeometricAlignment{
		UVBoundaryFixed:        true,
		ContactShapeImmutable:  true,
		F0CutoffFixed:          true,
		IRSchemeSelected:       false,
		NativeCountertermFound: false,
		Interpretation:         "The ASHA contact boundary fixes λ/g² and the native tree proxy, but UV spectral data do not by themselves choose an on-shell versus MS-bar finite renormalization prescription at the electroweak scale.",
		Status:                 StatusGeometricAlignmentAudited,
	}
}

func preserveFirewalls() Firewalls {
	return Firewalls{NoFullSMCoefficientTable: true, NoNativeCounterterms: true, NoGaugeInputScheme: true, NoNativeIRScheme: true, NoExactPoleMassClaim: true, Status: StatusPrecisionFirewallsPreserved}
}

func compileSummary(in Inputs, blocks PVBlockLedger, c CountertermMap) Summary {
	direct := "PV A0/B0 structures and the on-shell/MS-bar scheme ledger are formalized. The required pole finite residue is exactly mapped, but no exact collider pole mass is claimed because the full SM coefficient table and finite counterterms are still external continuum data."
	next := "Install the complete renormalized SM Higgs one-loop coefficient table in a chosen input scheme, then contract it against the finite PV blocks and counterterms."
	return Summary{RequiredRePiGeV2: in.RequiredRePiGeV2, RequiredFiniteRemainder: c.RequiredFiniteRemainder, PVBlocks: len(blocks.Blocks), DirectAnswer: direct, NextGate: next, Status: StatusPrecisionFirewallsPreserved}
}

func Statuses(a Analysis) []string {
	return []string{
		a.Inputs.Status,
		a.PV.Status,
		a.Blocks.Status,
		a.Schemes.Status,
		a.Counterterm.Status,
		a.Alignment.Status,
		a.Firewalls.Status,
		a.Summary.Status,
		StatusTensionSchemeNotNative,
		StatusTensionCountertermRequired,
		StatusTensionPVBlocksNotEnough,
		StatusFailedFullSMCoefficientTable,
		StatusFailedCountertermsDerived,
		StatusFailedGaugeInputScheme,
		StatusFailedNativeIRScheme,
		StatusFailedExactPoleMassClaim,
	}
}

func FormatInputs(x Inputs) string {
	return fmt.Sprintf("gate=%d m_native=%.15f M_ref=%.15f λ=%.15f ReΠ_req=%.15f raw=%.15f rem=%.15f μ=%.15f inputs=[%s] status=%s", x.HighestInheritedGate, x.NativeMassGeV, x.ObservedPoleGeV, x.LambdaNative, x.RequiredRePiGeV2, x.RawPolynomialGeV2, x.FiniteRemainderGeV2, x.MuGeV, strings.Join(x.QuarantinedInputs, ", "), x.Status)
}

func FormatPVStructure(x PVStructure) string {
	parts := make([]string, 0, len(x.Definitions))
	for _, d := range x.Definitions {
		parts = append(parts, fmt.Sprintf("%s:%s role=%s", d.Name, d.Definition, d.Role))
	}
	return fmt.Sprintf("defs=[%s] selfEnergy=%s pole=%s status=%s", strings.Join(parts, " | "), x.SelfEnergy, x.PoleEquation, x.Status)
}

func FormatPVBlock(x PVBlock) string {
	return fmt.Sprintf("%s m=%.6f s=%.6f μ=%.6f A0=%.12f B0=%.12f 4m²/s=%.9f real=%v", x.Particle, x.MassGeV, x.SGeV2, x.MuGeV, x.A0FiniteGeV2, x.B0Finite, x.ThresholdRatio, x.RealBranch)
}

func FormatPVBlocks(x PVBlockLedger) string {
	parts := make([]string, 0, len(x.Blocks))
	for _, b := range x.Blocks {
		parts = append(parts, FormatPVBlock(b))
	}
	return fmt.Sprintf("blocks=[%s] status=%s", strings.Join(parts, " | "), x.Status)
}

func FormatSchemeLane(x SchemeLane) string {
	return fmt.Sprintf("%s mass=%s counterterm=%s scale=%s native=%v canHitTarget=%v role=%s", x.Name, x.MassCondition, x.CountertermRule, x.ScaleDependence, x.NativeSelected, x.CanHitTarget, x.Interpretation)
}

func FormatSchemes(x SchemeAudit) string {
	parts := make([]string, 0, len(x.Lanes))
	for _, l := range x.Lanes {
		parts = append(parts, FormatSchemeLane(l))
	}
	return fmt.Sprintf("chosen=%s lanes=[%s] status=%s", x.ChosenForAudit, strings.Join(parts, " | "), x.Status)
}

func FormatCounterterm(x CountertermMap) string {
	return fmt.Sprintf("pole=%s ReΠ_req=%.15f raw=%.15f finiteRemainder=%.15f rem/target=%.12f rem/|raw|=%.12f role=%s status=%s", x.PoleEquation, x.RequiredRePiGeV2, x.RawPolynomialGeV2, x.RequiredFiniteRemainder, x.RemainderOverTarget, x.RemainderOverRawAbs, x.Interpretation, x.Status)
}

func FormatAlignment(x GeometricAlignment) string {
	return fmt.Sprintf("uv=%v shape=%v f0=%v irScheme=%v nativeCT=%v role=%s status=%s", x.UVBoundaryFixed, x.ContactShapeImmutable, x.F0CutoffFixed, x.IRSchemeSelected, x.NativeCountertermFound, x.Interpretation, x.Status)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("noCoeff=%v noNativeCT=%v noGaugeInput=%v noNativeIR=%v noPole=%v status=%s", x.NoFullSMCoefficientTable, x.NoNativeCounterterms, x.NoGaugeInputScheme, x.NoNativeIRScheme, x.NoExactPoleMassClaim, x.Status)
}

func FormatSummary(x Summary) string {
	return fmt.Sprintf("ReΠ_req=%.15f rem=%.15f pvBlocks=%d answer=%s next=%s status=%s", x.RequiredRePiGeV2, x.RequiredFiniteRemainder, x.PVBlocks, x.DirectAnswer, x.NextGate, x.Status)
}

func FormatStatuses(ss []string) string  { return "statuses=" + strings.Join(ss, "; ") }
func nearlyEqual(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
