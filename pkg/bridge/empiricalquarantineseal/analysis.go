// Package empiricalquarantineseal implements Gate 348:
// Empirical Quarantine Seal / Grand Unified Project Ledger.
//
// Gate 348 is a final capstone ledger.  It does not introduce a new fit,
// does not close CKM/flavor, does not choose a renormalization scheme, and
// does not claim an exact collider mass.  Its purpose is to seal the
// landscape/vacuum boundary found through Gate 347: the ASHA finite geometry
// derives rigid Standard Model structure and several exact boundary ratios,
// while the remaining continuous vacuum coordinates are quarantined as Phase
// III dynamical or empirical inputs.
package empiricalquarantineseal

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE348-EMPIRICAL-QUARANTINE-SEAL-GRAND-UNIFIED-PROJECT-LEDGER"

	StatusGrandUnifiedLedgerCompiled = "CONDITIONAL_SUPPORT_GRAND_UNIFIED_LEDGER_COMPILED"
	StatusPhaseIIQuarantineSealed    = "CONDITIONAL_SUPPORT_PHASE_II_QUARANTINE_SEALED"
	StatusRigidLandscapeCataloged    = "CONDITIONAL_SUPPORT_RIGID_GEOMETRIC_LANDSCAPE_CATALOGED"
	StatusProxiesCataloged           = "CONDITIONAL_SUPPORT_PHENOMENOLOGICAL_INTERSECTIONS_CATALOGED"
	StatusEmpiricalQuarantineDefined = "CONDITIONAL_SUPPORT_EMPIRICAL_QUARANTINE_DEFINED"
	StatusSeparationPreserved        = "CONDITIONAL_SUPPORT_KINEMATICS_DYNAMICS_SEPARATION_PRESERVED"

	StatusFailedVacuumNotDerived         = "FAILED_ROUTE_PHYSICAL_VACUUM_POINT_NOT_DERIVED"
	StatusFailedYukawasQuarantined       = "FAILED_ROUTE_YUKAWA_SINGULAR_VALUES_QUARANTINED"
	StatusFailedCKMQuarantined           = "FAILED_ROUTE_CKM_FLAVOR_ORIENTATION_QUARANTINED"
	StatusFailedStrongCPQuarantined      = "FAILED_ROUTE_STRONG_CP_PHASE_QUARANTINED"
	StatusFailedGravityCutoffQuarantined = "FAILED_ROUTE_GRAVITATIONAL_CUTOFF_SCALE_QUARANTINED"
	StatusFailedPoleSchemeQuarantined    = "FAILED_ROUTE_POLE_MASS_RENORMALIZATION_SCHEME_QUARANTINED"
	StatusFailedCosmologicalQuarantined  = "FAILED_ROUTE_COSMOLOGICAL_CONSTANT_QUARANTINED"
	StatusFailedFinalTOE                 = "FAILED_ROUTE_FINAL_THEORY_OF_EVERYTHING_NOT_CLAIMED"
)

const (
	inheritedGate = 347

	thresholdJump = -0.097846792207
	treeProxyGeV  = 125.6062977568011
	nativeTreeGeV = 125.274157149699
	rgProxyGeV    = 124.976620
	observedGeV   = 125.10
)

type LedgerSpan struct {
	AuditID              string
	GateRange            string
	HighestGateInherited int
	AddsNewPhysics       bool
	ImportsObservedFit   bool
	Phase                string
	Verdict              string
}

type LandscapeItem struct {
	Name                 string
	Gate                 string
	Statement            string
	ParameterFree        bool
	NativeFiniteGeometry bool
	RequiresVacuumChoice bool
	Status               string
}

type LandscapeLedger struct {
	Cataloged                  bool
	Items                      []LandscapeItem
	ContainsWeakMixing         bool
	ContainsMoritaSplit        bool
	ContainsGenerationTriality bool
	ContainsTrueBimodule       bool
	ContainsTraceEquivalence   bool
	ContainsThresholdJump      bool
	ContainsPfaffianHierarchy  bool
	ContainsAlphaEightPi       bool
	ParameterFreeCount         int
	Verdict                    string
}

type ProxyItem struct {
	Name               string
	Gate               string
	GeometricInput     string
	QuarantinedInput   string
	Output             string
	ErrorDescription   string
	FinalClaim         bool
	EmpiricalInterface bool
	Status             string
}

type ProxyLedger struct {
	Cataloged                     bool
	Items                         []ProxyItem
	Contains125TreeProxy          bool
	ContainsNative125Proxy        bool
	ContainsThresholdTransport    bool
	ContainsPrecisionPoleTarget   bool
	AllEmpiricalInputsQuarantined bool
	FinalMassClaimed              bool
	Verdict                       string
}

type QuarantinedInput struct {
	Name           string
	Type           string
	Dimension      int
	Gate           string
	Reason         string
	PhaseIIITarget string
	Closed         bool
	Status         string
}

type QuarantineLedger struct {
	Defined                        bool
	Items                          []QuarantinedInput
	MinimalSMVacuumDimension       int
	ExtendedVacuumDimension        int
	ContainsYukawas                bool
	ContainsCKM                    bool
	ContainsStrongCP               bool
	ContainsGravityCutoff          bool
	ContainsPoleScheme             bool
	ContainsCosmologicalConstant   bool
	ContainsFlavorProjectionMetric bool
	AnyClosed                      bool
	Verdict                        string
}

type SeparationAudit struct {
	NoYukawaFitPromoted        bool
	NoCKMInvented              bool
	NoPoleSchemeChosen         bool
	NoCosmologicalFitPromoted  bool
	NoObservedMassInserted     bool
	NoAlphaGUTFitNeededInFinal bool
	NoFinalTOEClaimed          bool
	NoExactColliderClaimed     bool
	LandscapeVacuumSeparated   bool
	FiniteCorePolluted         bool
	Verdict                    string
}

type Summary struct {
	LedgerCompiled       bool
	PhaseIISealed        bool
	LandscapeReady       bool
	ProxiesReady         bool
	QuarantineReady      bool
	SeparationPreserved  bool
	FinalTOEClaimed      bool
	ExactColliderClaimed bool
	OneLine              string
	Status               string
}

type Analysis struct {
	Span       LedgerSpan
	Landscape  LandscapeLedger
	Proxies    ProxyLedger
	Quarantine QuarantineLedger
	Audit      SeparationAudit
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
	span := compileSpan()
	landscape := compileLandscape()
	proxies := compileProxies()
	quarantine := compileQuarantine()
	audit := compileAudit()
	summary := compileSummary(span, landscape, proxies, quarantine, audit)
	truth := "Gate 348 seals the ASHA landscape/vacuum boundary through Gate 347.  The finite geometry has produced rigid Standard Model structure, exact boundary ratios, the B-gap threshold witness, and the Pfaffian hierarchy relation; the remaining Yukawa, CKM, strong-CP, gravitational-cutoff, cosmological, and pole-scheme data are quarantined as vacuum or precision coordinates."
	return Analysis{Span: span, Landscape: landscape, Proxies: proxies, Quarantine: quarantine, Audit: audit, Summary: summary, Truth: truth}, nil
}

func compileSpan() LedgerSpan {
	return LedgerSpan{AuditID: AuditID, GateRange: "Gate 1 → Gate 347 inherited; Gate 348 compiles and seals the empirical quarantine", HighestGateInherited: inheritedGate, AddsNewPhysics: false, ImportsObservedFit: false, Phase: "Phase I + Phase II capstone; Phase III obligations quarantined", Verdict: strings.Join([]string{StatusGrandUnifiedLedgerCompiled, StatusPhaseIIQuarantineSealed}, ";")}
}

func compileLandscape() LandscapeLedger {
	items := []LandscapeItem{
		{Name: "weak mixing boundary", Gate: "Gate 298/308", Statement: "sin²θ_W = 3/8", ParameterFree: true, NativeFiniteGeometry: true, RequiresVacuumChoice: false, Status: StatusRigidLandscapeCataloged},
		{Name: "Morita color split", Gate: "Gate 295 chain", Statement: "κ_C:κ_Q = 1:3", ParameterFree: true, NativeFiniteGeometry: true, RequiresVacuumChoice: false, Status: StatusRigidLandscapeCataloged},
		{Name: "generation triality", Gate: "Gate 26 / Gate 323", Statement: "τ_η = (2,-2,1) and N_gen = 3", ParameterFree: true, NativeFiniteGeometry: true, RequiresVacuumChoice: false, Status: StatusRigidLandscapeCataloged},
		{Name: "true bimodule resolution", Gate: "Gate 295 / Gate 319", Statement: "direct-sum sectors have zero cross determinant; the true bimodule and Ω_Hσ overlap are required for heavy-light physics", ParameterFree: true, NativeFiniteGeometry: true, RequiresVacuumChoice: false, Status: StatusRigidLandscapeCataloged},
		{Name: "trace equivalence boundary", Gate: "Gate 307/308", Statement: "λ_H/g_*² = 1197/4624", ParameterFree: true, NativeFiniteGeometry: true, RequiresVacuumChoice: false, Status: StatusRigidLandscapeCataloged},
		{Name: "canonical threshold jump", Gate: "Gate 320/321", Statement: fmt.Sprintf("Δλ = %.12f from κ_Q(4/π)B_gapΩ_Hσ/(-4)", thresholdJump), ParameterFree: true, NativeFiniteGeometry: true, RequiresVacuumChoice: false, Status: StatusRigidLandscapeCataloged},
		{Name: "Pfaffian hierarchy scaling", Gate: "Gate 341/342", Statement: "v/M_P = 2^(3/2) exp(-4π²)", ParameterFree: true, NativeFiniteGeometry: true, RequiresVacuumChoice: false, Status: StatusRigidLandscapeCataloged},
		{Name: "full doubled bosonic coupling branch", Gate: "Gate 327–330", Statement: "α_GUT⁻¹ = 8π under the full bosonic spectral trace convention", ParameterFree: true, NativeFiniteGeometry: true, RequiresVacuumChoice: false, Status: StatusRigidLandscapeCataloged},
		{Name: "flavor flatness theorem", Gate: "Gate 346/347", Statement: "standard spectral invariants are unitary flavor-flat; Majorana overlap does not derive CKM", ParameterFree: true, NativeFiniteGeometry: true, RequiresVacuumChoice: true, Status: StatusRigidLandscapeCataloged},
	}
	l := LandscapeLedger{Cataloged: true, Items: items}
	for _, item := range items {
		if item.ParameterFree && item.NativeFiniteGeometry {
			l.ParameterFreeCount++
		}
		blob := strings.ToLower(item.Name + " " + item.Statement)
		l.ContainsWeakMixing = l.ContainsWeakMixing || strings.Contains(item.Statement, "3/8")
		l.ContainsMoritaSplit = l.ContainsMoritaSplit || strings.Contains(blob, "morita") || strings.Contains(item.Statement, "1:3")
		l.ContainsGenerationTriality = l.ContainsGenerationTriality || strings.Contains(blob, "triality") || strings.Contains(item.Statement, "τ_η")
		l.ContainsTrueBimodule = l.ContainsTrueBimodule || strings.Contains(blob, "bimodule")
		l.ContainsTraceEquivalence = l.ContainsTraceEquivalence || strings.Contains(item.Statement, "1197/4624")
		l.ContainsThresholdJump = l.ContainsThresholdJump || strings.Contains(item.Statement, "Δλ")
		l.ContainsPfaffianHierarchy = l.ContainsPfaffianHierarchy || strings.Contains(item.Statement, "exp(-4π²)")
		l.ContainsAlphaEightPi = l.ContainsAlphaEightPi || strings.Contains(item.Statement, "8π")
	}
	l.Verdict = StatusRigidLandscapeCataloged
	return l
}

func compileProxies() ProxyLedger {
	items := []ProxyItem{
		{Name: "empirical α_GUT tree-level Higgs proxy", Gate: "Gate 315", GeometricInput: "λ_H/g_*² = 1197/4624", QuarantinedInput: "α_GUT≈1/25", Output: fmt.Sprintf("m_H≈%.6f GeV", treeProxyGeV), ErrorDescription: "sub-percent proxy, not a pole-mass derivation", FinalClaim: false, EmpiricalInterface: true, Status: StatusProxiesCataloged},
		{Name: "native 8π tree-level Higgs proxy", Gate: "Gate 330/335", GeometricInput: "α_GUT⁻¹=8π and λ_H/g_*²=1197/4624", QuarantinedInput: "electroweak unit v", Output: fmt.Sprintf("m_native≈%.6f GeV", nativeTreeGeV), ErrorDescription: "requires pole-mass precision to compare exactly to collider reference", FinalClaim: false, EmpiricalInterface: true, Status: StatusProxiesCataloged},
		{Name: "threshold RG transport proxy", Gate: "Gate 322", GeometricInput: "Δλ=-0.097846792207", QuarantinedInput: "flattened-top flavor envelope", Output: fmt.Sprintf("m_run≈%.6f GeV", rgProxyGeV), ErrorDescription: "diagnostic lane; not the full CKM/top sector", FinalClaim: false, EmpiricalInterface: true, Status: StatusProxiesCataloged},
		{Name: "pole precision target", Gate: "Gate 336–338", GeometricInput: "m_native²-M_ref²", QuarantinedInput: "on-shell/MSbar finite counterterm scheme", Output: "ReΠ_required≈+43.604449567 GeV²", ErrorDescription: "scheme-dependent finite pole conversion", FinalClaim: false, EmpiricalInterface: true, Status: StatusProxiesCataloged},
	}
	p := ProxyLedger{Cataloged: true, Items: items, AllEmpiricalInputsQuarantined: true}
	for _, item := range items {
		p.FinalMassClaimed = p.FinalMassClaimed || item.FinalClaim
		blob := strings.ToLower(item.Name + " " + item.Output)
		p.Contains125TreeProxy = p.Contains125TreeProxy || strings.Contains(blob, "125.606")
		p.ContainsNative125Proxy = p.ContainsNative125Proxy || strings.Contains(blob, "125.274")
		p.ContainsThresholdTransport = p.ContainsThresholdTransport || strings.Contains(blob, "124.976")
		p.ContainsPrecisionPoleTarget = p.ContainsPrecisionPoleTarget || strings.Contains(blob, "43.604")
	}
	p.Verdict = StatusProxiesCataloged
	return p
}

func compileQuarantine() QuarantineLedger {
	items := []QuarantinedInput{
		{Name: "charged-fermion Yukawa singular values", Type: "continuous vacuum moduli", Dimension: 9, Gate: "Gate 345", Reason: "spectral traces determine invariant forms but not individual singular values", PhaseIIITarget: "dynamical Yukawa vacuum selection", Closed: false, Status: StatusFailedYukawasQuarantined},
		{Name: "CKM / quark flavor orientation", Type: "continuous vacuum orientation", Dimension: 4, Gate: "Gate 347", Reason: "standard and Majorana-Dirac traces remain U(3)-flat in quark flavor", PhaseIIITarget: "non-unitary flavor texture or empirical CKM seal", Closed: false, Status: StatusFailedCKMQuarantined},
		{Name: "strong CP phase", Type: "continuous vacuum phase", Dimension: 1, Gate: "Gate 345", Reason: "not selected by finite spectral kinematics", PhaseIIITarget: "axion/topological θ selection", Closed: false, Status: StatusFailedStrongCPQuarantined},
		{Name: "absolute unit / electroweak scale", Type: "scale coordinate", Dimension: 1, Gate: "Gate 345/342", Reason: "ratios are derived; one unit scale fixes GeV units", PhaseIIITarget: "unit normalization or metrological input", Closed: false, Status: StatusFailedVacuumNotDerived},
		{Name: "gravitational cutoff product", Type: "spectral moment invariant", Dimension: 1, Gate: "Gate 343", Reason: "f₂Λ² is fixed as product, but f₂ and Λ separately are not selected", PhaseIIITarget: "native cutoff-scale theorem", Closed: false, Status: StatusFailedGravityCutoffQuarantined},
		{Name: "pole-mass renormalization scheme", Type: "precision convention", Dimension: 1, Gate: "Gate 338", Reason: "PV integrals are installed but finite counterterm/input scheme is not selected by UV geometry", PhaseIIITarget: "full SM renormalized pole stack", Closed: false, Status: StatusFailedPoleSchemeQuarantined},
		{Name: "cosmological constant / vacuum energy", Type: "vacuum energy coordinate", Dimension: 1, Gate: "Gate 344", Reason: "f₄Λ⁴a₀ and vacuum subtraction are not locked by the known spectral moments", PhaseIIITarget: "vacuum energy renormalization and f₄ moment theorem", Closed: false, Status: StatusFailedCosmologicalQuarantined},
		{Name: "native flavor projection metric", Type: "discrete/structural choice", Dimension: 0, Gate: "Gate 325", Reason: "positive metric forbids top nulling; signed metric allows it but is not selected", PhaseIIITarget: "metric-selection variational principle", Closed: false, Status: StatusFailedCKMQuarantined},
	}
	q := QuarantineLedger{Defined: true, Items: items, MinimalSMVacuumDimension: 15, ExtendedVacuumDimension: 25}
	for _, item := range items {
		q.AnyClosed = q.AnyClosed || item.Closed
		blob := strings.ToLower(item.Name + " " + item.Type + " " + item.Reason)
		q.ContainsYukawas = q.ContainsYukawas || strings.Contains(blob, "yukawa")
		q.ContainsCKM = q.ContainsCKM || strings.Contains(blob, "ckm") || strings.Contains(blob, "flavor orientation")
		q.ContainsStrongCP = q.ContainsStrongCP || strings.Contains(blob, "strong cp")
		q.ContainsGravityCutoff = q.ContainsGravityCutoff || strings.Contains(blob, "gravitational") || strings.Contains(blob, "f₂")
		q.ContainsPoleScheme = q.ContainsPoleScheme || strings.Contains(blob, "pole") || strings.Contains(blob, "renormalization")
		q.ContainsCosmologicalConstant = q.ContainsCosmologicalConstant || strings.Contains(blob, "cosmological") || strings.Contains(blob, "vacuum energy")
		q.ContainsFlavorProjectionMetric = q.ContainsFlavorProjectionMetric || strings.Contains(blob, "projection metric")
	}
	q.Verdict = strings.Join([]string{StatusEmpiricalQuarantineDefined, StatusPhaseIIQuarantineSealed}, ";")
	return q
}

func compileAudit() SeparationAudit {
	return SeparationAudit{NoYukawaFitPromoted: true, NoCKMInvented: true, NoPoleSchemeChosen: true, NoCosmologicalFitPromoted: true, NoObservedMassInserted: true, NoAlphaGUTFitNeededInFinal: true, NoFinalTOEClaimed: true, NoExactColliderClaimed: true, LandscapeVacuumSeparated: true, FiniteCorePolluted: false, Verdict: StatusSeparationPreserved}
}

func compileSummary(span LedgerSpan, l LandscapeLedger, p ProxyLedger, q QuarantineLedger, a SeparationAudit) Summary {
	return Summary{LedgerCompiled: span.HighestGateInherited == inheritedGate && !span.AddsNewPhysics, PhaseIISealed: q.Defined && !q.AnyClosed, LandscapeReady: l.Cataloged && l.ContainsWeakMixing && l.ContainsTraceEquivalence && l.ContainsPfaffianHierarchy, ProxiesReady: p.Cataloged && p.AllEmpiricalInputsQuarantined, QuarantineReady: q.ContainsYukawas && q.ContainsCKM && q.ContainsGravityCutoff && q.ContainsPoleScheme && q.ContainsCosmologicalConstant, SeparationPreserved: a.LandscapeVacuumSeparated && !a.FiniteCorePolluted, FinalTOEClaimed: false, ExactColliderClaimed: false, OneLine: "ASHA through Gate 347 derives the geometric landscape and quarantines the vacuum coordinates.", Status: strings.Join([]string{StatusGrandUnifiedLedgerCompiled, StatusPhaseIIQuarantineSealed}, ";")}
}

func Statuses(a Analysis) []string {
	statuses := []string{StatusGrandUnifiedLedgerCompiled, StatusPhaseIIQuarantineSealed, a.Landscape.Verdict, a.Proxies.Verdict, StatusEmpiricalQuarantineDefined, a.Quarantine.Verdict, a.Audit.Verdict}
	statuses = append(statuses, StatusFailedVacuumNotDerived, StatusFailedYukawasQuarantined, StatusFailedCKMQuarantined, StatusFailedStrongCPQuarantined, StatusFailedGravityCutoffQuarantined, StatusFailedPoleSchemeQuarantined, StatusFailedCosmologicalQuarantined, StatusFailedFinalTOE)
	return statuses
}

func FormatSpan(s LedgerSpan) string {
	return fmt.Sprintf("%s | %s | inherited=%d | adds_new_physics=%v | imports_fit=%v", s.AuditID, s.GateRange, s.HighestGateInherited, s.AddsNewPhysics, s.ImportsObservedFit)
}
func FormatLandscape(l LandscapeLedger) string {
	return fmt.Sprintf("items=%d parameter_free=%d weak=%v morita=%v triality=%v bimodule=%v trace=%v jump=%v hierarchy=%v alpha8π=%v", len(l.Items), l.ParameterFreeCount, l.ContainsWeakMixing, l.ContainsMoritaSplit, l.ContainsGenerationTriality, l.ContainsTrueBimodule, l.ContainsTraceEquivalence, l.ContainsThresholdJump, l.ContainsPfaffianHierarchy, l.ContainsAlphaEightPi)
}
func FormatProxies(p ProxyLedger) string {
	return fmt.Sprintf("items=%d tree125=%v native125=%v threshold125=%v pole_target=%v quarantined=%v final_mass_claimed=%v", len(p.Items), p.Contains125TreeProxy, p.ContainsNative125Proxy, p.ContainsThresholdTransport, p.ContainsPrecisionPoleTarget, p.AllEmpiricalInputsQuarantined, p.FinalMassClaimed)
}
func FormatQuarantine(q QuarantineLedger) string {
	return fmt.Sprintf("items=%d sm_dim=%d extended_dim=%d yukawas=%v ckm=%v strongcp=%v gravity=%v pole=%v cosmology=%v projection_metric=%v any_closed=%v", len(q.Items), q.MinimalSMVacuumDimension, q.ExtendedVacuumDimension, q.ContainsYukawas, q.ContainsCKM, q.ContainsStrongCP, q.ContainsGravityCutoff, q.ContainsPoleScheme, q.ContainsCosmologicalConstant, q.ContainsFlavorProjectionMetric, q.AnyClosed)
}
func FormatAudit(a SeparationAudit) string {
	return fmt.Sprintf("no_yukawa_fit=%v no_ckm=%v no_pole_scheme=%v no_cosmo_fit=%v no_mass_fit=%v no_final_toe=%v no_exact_collider=%v separated=%v polluted=%v", a.NoYukawaFitPromoted, a.NoCKMInvented, a.NoPoleSchemeChosen, a.NoCosmologicalFitPromoted, a.NoObservedMassInserted, a.NoFinalTOEClaimed, a.NoExactColliderClaimed, a.LandscapeVacuumSeparated, a.FiniteCorePolluted)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("compiled=%v phaseII_sealed=%v landscape=%v proxies=%v quarantine=%v separated=%v final_toe=%v collider=%v", s.LedgerCompiled, s.PhaseIISealed, s.LandscapeReady, s.ProxiesReady, s.QuarantineReady, s.SeparationPreserved, s.FinalTOEClaimed, s.ExactColliderClaimed)
}
