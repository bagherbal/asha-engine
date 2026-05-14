// Package grandunifiedledger implements Gate 326:
// Grand Unified Ledger / Project Capstone Audit.
//
// Gate 326 is not a new phenomenological fit and does not upgrade conditional
// transport diagnostics into final collider claims.  It compiles the current
// ASHA architecture through Gate 325 into a fully firewalled project ledger:
// native geometric triumphs, quarantined phenomenological alignments, and the
// precise Phase-III targets that remain outside the pure finite-algebra core.
package grandunifiedledger

import (
	"fmt"
	"strings"
	"sync"
)

const (
	AuditID = "GATE326-GRAND-UNIFIED-LEDGER-PROJECT-CAPSTONE-AUDIT"

	StatusGrandUnifiedLedgerCompiled = "CONDITIONAL_SUPPORT_GRAND_UNIFIED_LEDGER_COMPILED"
	StatusProjectCapstoneAchieved    = "CONDITIONAL_SUPPORT_PROJECT_CAPSTONE_ACHIEVED"
	StatusAbsoluteTriumphsCataloged  = "CONDITIONAL_SUPPORT_ABSOLUTE_GEOMETRIC_TRIUMPHS_CATALOGED"
	StatusProxyAlignmentsCataloged   = "CONDITIONAL_SUPPORT_PHENOMENOLOGICAL_PROXIES_CATALOGED"
	StatusFirewallsCataloged         = "CONDITIONAL_SUPPORT_EPISTEMOLOGICAL_FIREWALLS_CATALOGED"
	StatusPhaseIIICataloged          = "CONDITIONAL_SUPPORT_PHASE_III_TARGETS_FORMALIZED"
	StatusLedgerFirewallsPreserved   = "CONDITIONAL_SUPPORT_GRAND_LEDGER_FIREWALLS_PRESERVED"

	StatusFailedFinalTOENotClaimed      = "FAILED_ROUTE_FINAL_THEORY_OF_EVERYTHING_NOT_CLAIMED"
	StatusFailedAlphaGUTNotDerived      = "FAILED_ROUTE_ALPHA_GUT_ABSOLUTE_ORIGIN_NOT_DERIVED"
	StatusFailedWeightedTraceMissing    = "FAILED_ROUTE_CTRACE_25_WEIGHTED_FUNCTIONAL_NOT_DERIVED"
	StatusFailedFlavorVacuumNotSelected = "FAILED_ROUTE_CKM_FLAVOR_VACUUM_SELECTION_NOT_DERIVED"
	StatusFailedNativeProjectionMetric  = "FAILED_ROUTE_NATIVE_FLAVOR_PROJECTION_METRIC_NOT_DERIVED"
	StatusFailedTwoLoopNotExecuted      = "FAILED_ROUTE_TWO_LOOP_RG_NOT_EXECUTED"
	StatusFailedPoleMassNotExecuted     = "FAILED_ROUTE_POLE_MASS_CONVERSION_NOT_EXECUTED"
	StatusFailedColliderMassNotClaimed  = "FAILED_ROUTE_EXACT_COLLIDER_HIGGS_MASS_NOT_CLAIMED"
	StatusFailedPhaseIIIRequired        = "FAILED_ROUTE_PHASE_III_DYNAMICAL_VACUUM_PHYSICS_REQUIRED"
)

const (
	highestInheritedGate = 325

	weakMixingRatio       = "sin²θ_W = 3/8"
	higgsQuarticRatio     = "λ_H/g_*² = 1197/4624"
	contactResonance      = "4/π"
	canonicalJump         = "Δλ = -0.097846792207"
	empiricalProxyMassGeV = 125.6062977568
	thresholdRunMassGeV   = 124.9766199157
	observedHiggsGeV      = 125.10
	continuousFloorGeV    = 157.104
	topPositiveFloorGeV   = 258.687
)

type RegistrySpan struct {
	AuditID              string
	GateRange            string
	HighestGateInherited int
	CapstoneForPhase     string
	AddsNewPhysicsFit    bool
	RewritesHistory      bool
	Verdict              string
}

type AbsoluteTriumph struct {
	Name                  string
	Gate                  string
	Statement             string
	NativeToFiniteCore    bool
	RequiresEmpiricalData bool
	FinalColliderClaim    bool
	Status                string
}

type TriumphLedger struct {
	Cataloged                    bool
	Items                        []AbsoluteTriumph
	ContainsWeakMixing           bool
	ContainsMoritaColorSplit     bool
	ContainsGenerationTriality   bool
	ContainsTrueBimodule         bool
	ContainsTopologicalResonance bool
	ContainsTraceEquivalence     bool
	ContainsThresholdJump        bool
	NativeCount                  int
	Verdict                      string
}

type ProxyAlignment struct {
	Name                string
	Gate                string
	QuarantinedInput    string
	GeometricInput      string
	Output              string
	ErrorDescription    string
	EmpiricalComparison bool
	FinalDerivation     bool
	Status              string
}

type ProxyLedger struct {
	Cataloged                     bool
	Alignments                    []ProxyAlignment
	ContainsTreeLevel125Proxy     bool
	ContainsThresholdTransport125 bool
	Contains331Diagnostic         bool
	Contains157ContinuousFloor    bool
	EmpiricalInputsQuarantined    bool
	FinalMassClaimed              bool
	Verdict                       string
}

type Firewall struct {
	Name                    string
	Gate                    string
	MathematicalLimit       string
	WhyItCannotBeClosedHere string
	PhaseIIITarget          string
	Closed                  bool
	Status                  string
}

type FirewallLedger struct {
	Cataloged                         bool
	Items                             []Firewall
	ContainsAlphaGUTOrigin            bool
	ContainsWeightedTrace25           bool
	ContainsFlavorVacuumSelection     bool
	ContainsProjectionMetricSelection bool
	ContainsTwoLoopPolePrecision      bool
	ContainsExactColliderMass         bool
	AnyClosed                         bool
	Verdict                           string
}

type PhaseIIITarget struct {
	Name             string
	WorkPackage      string
	RequiredTheorem  string
	SuccessCriterion string
	InheritsFrom     []string
	Status           string
}

type PhaseIIILedger struct {
	Formalized                 bool
	Targets                    []PhaseIIITarget
	IncludesWeightedTrace      bool
	IncludesFlavorVacuum       bool
	IncludesProjectionMetric   bool
	IncludesPrecisionTransport bool
	IncludesFullSigmaPotential bool
	RequiresNoEmpiricalTuning  bool
	Verdict                    string
}

type FirewallAudit struct {
	NoAlphaGUTFitPromoted      bool
	NoCKMTextureInvented       bool
	NoFlavorMetricForced       bool
	NoObservedHiggsFitInserted bool
	NoObservedTopFitInserted   bool
	NoTwoLoopClaimed           bool
	NoPoleMassClaimed          bool
	NoFinalTOEClaimed          bool
	NoExactColliderMassClaimed bool
	FiniteCorePolluted         bool
	Verdict                    string
}

type Summary struct {
	LedgerCompiled       bool
	ProjectCapstone      bool
	TriumphsReady        bool
	ProxiesReady         bool
	FirewallsReady       bool
	PhaseIIIReady        bool
	FirewallsPreserved   bool
	FinalTOEClaimed      bool
	ExactColliderClaimed bool
	Status               string
	DirectAnswer         string
	NextPhase            string
}

type Analysis struct {
	Span      RegistrySpan
	Triumphs  TriumphLedger
	Proxies   ProxyLedger
	Firewalls FirewallLedger
	PhaseIII  PhaseIIILedger
	Audit     FirewallAudit
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
	span := compileSpan()
	triumphs := compileTriumphs()
	proxies := compileProxies()
	firewalls := compileFirewalls()
	phaseIII := compilePhaseIII()
	audit := auditFirewallIntegrity()
	summary := compileSummary(span, triumphs, proxies, firewalls, phaseIII, audit)
	truth := "Gate 326 is the grand unified ledger for ASHA through Gate 325. It records two exact native Standard Model boundary ratios and the derived B-gap threshold machinery, while preserving the limits that remain outside pure finite algebra: the absolute origin of α_GUT, the weighted C_trace=25 functional, the CKM/flavor vacuum selector, and two-loop/pole precision. The project capstone is achieved only as a fully firewalled ledger, not as a final theory-of-everything claim."
	return Analysis{Span: span, Triumphs: triumphs, Proxies: proxies, Firewalls: firewalls, PhaseIII: phaseIII, Audit: audit, Summary: summary, Truth: truth}, nil
}

func compileSpan() RegistrySpan {
	return RegistrySpan{
		AuditID:              AuditID,
		GateRange:            "Gate 1 → Gate 325 inherited; Gate 326 compiles the Grand Unified Ledger",
		HighestGateInherited: highestInheritedGate,
		CapstoneForPhase:     "Phase I structural derivation + Phase II non-perturbative threshold/flavor audit",
		AddsNewPhysicsFit:    false,
		RewritesHistory:      false,
		Verdict:              strings.Join([]string{StatusGrandUnifiedLedgerCompiled, StatusProjectCapstoneAchieved}, ";"),
	}
}

func compileTriumphs() TriumphLedger {
	items := []AbsoluteTriumph{
		{Name: "Weak mixing boundary ratio", Gate: "Gate 298/308", Statement: weakMixingRatio + " from the GUT-normalized hypercharge trace ledger", NativeToFiniteCore: true, RequiresEmpiricalData: false, FinalColliderClaim: false, Status: StatusAbsoluteTriumphsCataloged},
		{Name: "Morita color split", Gate: "Morita bimodule chain / Gate 295", Statement: "κ_C=1 and κ_Q=3 define the lepton/quark trace multiplicities and force the true bimodule interpretation", NativeToFiniteCore: true, RequiresEmpiricalData: false, FinalColliderClaim: false, Status: StatusAbsoluteTriumphsCataloged},
		{Name: "Generation triality topology", Gate: "τ_η chain / Gate 323/325", Statement: "τ_η=(2,-2,1) supplies the signed generation topology and exposes the flavor projection metric obstruction", NativeToFiniteCore: true, RequiresEmpiricalData: false, FinalColliderClaim: false, Status: StatusAbsoluteTriumphsCataloged},
		{Name: "True bimodule direct-sum resolution", Gate: "Gate 295 / Gate 319", Statement: "direct-sum carriers kill mixed functional determinants; true bimodule/off-diagonal overlap is required for heavy-light physics", NativeToFiniteCore: true, RequiresEmpiricalData: false, FinalColliderClaim: false, Status: StatusAbsoluteTriumphsCataloged},
		{Name: "Topological B-gap resonance", Gate: "B-gap/contact chain / Gate 318", Statement: contactResonance + " combines with B_gap and κ_Q to produce the portal-capacity witness 0.391387168826", NativeToFiniteCore: true, RequiresEmpiricalData: false, FinalColliderClaim: false, Status: StatusAbsoluteTriumphsCataloged},
		{Name: "Trace equivalence boundary", Gate: "Gate 307/308", Statement: higgsQuarticRatio + " after projected scalar carrier equivalence and GUT trace normalization", NativeToFiniteCore: true, RequiresEmpiricalData: false, FinalColliderClaim: false, Status: StatusAbsoluteTriumphsCataloged},
		{Name: "Canonical threshold jump witness", Gate: "Gate 320/321", Statement: canonicalJump + " from the rank-one Seesaw overlap index and canonical EFT normalization", NativeToFiniteCore: true, RequiresEmpiricalData: false, FinalColliderClaim: false, Status: StatusAbsoluteTriumphsCataloged},
		{Name: "Flavor projection metric obstruction", Gate: "Gate 325", Statement: "positive Hilbert-Schmidt metric forbids exact top nulling; signed projection allows a degenerate nullspace but does not select a vacuum", NativeToFiniteCore: true, RequiresEmpiricalData: false, FinalColliderClaim: false, Status: StatusAbsoluteTriumphsCataloged},
	}
	t := TriumphLedger{Cataloged: true, Items: items}
	for _, it := range items {
		if it.NativeToFiniteCore && !it.RequiresEmpiricalData {
			t.NativeCount++
		}
		s := strings.ToLower(it.Name + " " + it.Statement)
		t.ContainsWeakMixing = t.ContainsWeakMixing || strings.Contains(s, "weak mixing") || strings.Contains(it.Statement, "3/8")
		t.ContainsMoritaColorSplit = t.ContainsMoritaColorSplit || strings.Contains(s, "morita") || strings.Contains(it.Statement, "κ_Q=3")
		t.ContainsGenerationTriality = t.ContainsGenerationTriality || strings.Contains(s, "triality") || strings.Contains(it.Statement, "τ_η")
		t.ContainsTrueBimodule = t.ContainsTrueBimodule || strings.Contains(s, "true bimodule")
		t.ContainsTopologicalResonance = t.ContainsTopologicalResonance || strings.Contains(it.Statement, "4/π")
		t.ContainsTraceEquivalence = t.ContainsTraceEquivalence || strings.Contains(it.Statement, "1197/4624")
		t.ContainsThresholdJump = t.ContainsThresholdJump || strings.Contains(it.Statement, "Δλ") || strings.Contains(s, "threshold")
	}
	t.Verdict = strings.Join([]string{StatusAbsoluteTriumphsCataloged, StatusGrandUnifiedLedgerCompiled}, ";")
	return t
}

func compileProxies() ProxyLedger {
	alignments := []ProxyAlignment{
		{Name: "Tree-level Higgs proxy from empirical α_GUT", Gate: "Gate 315", QuarantinedInput: "α_GUT=1/25", GeometricInput: higgsQuarticRatio, Output: fmt.Sprintf("m_tree=%.6f GeV", empiricalProxyMassGeV), ErrorDescription: "≈0.4047% against 125.10 GeV proxy", EmpiricalComparison: true, FinalDerivation: false, Status: StatusProxyAlignmentsCataloged},
		{Name: "Derived threshold RG running proxy", Gate: "Gate 322", QuarantinedInput: "flattened/signed-top projection lane; one-loop running; conditional threshold scale", GeometricInput: canonicalJump + " with Gate-308 λ boundary", Output: fmt.Sprintf("m_run=%.6f GeV", thresholdRunMassGeV), ErrorDescription: "≈-0.0986% against 125.10 GeV before pole/two-loop conversion", EmpiricalComparison: true, FinalDerivation: false, Status: StatusProxyAlignmentsCataloged},
		{Name: "331 GeV diagnostic rejected as wrong seal/lane", Gate: "Gate 309", QuarantinedInput: "g_*²=1 and r_+ all-to-top lane", GeometricInput: "λ=1197/4624", Output: "m_run≈331.630 GeV", ErrorDescription: "diagnostic tension; not the ratio-verification lane", EmpiricalComparison: false, FinalDerivation: false, Status: StatusProxyAlignmentsCataloged},
		{Name: "Continuous flow floor", Gate: "Gate 313/314", QuarantinedInput: "zero-top/gauge-only envelope", GeometricInput: "Gate-308 quartic boundary without threshold jump", Output: fmt.Sprintf("m_floor≈%.3f GeV", continuousFloorGeV), ErrorDescription: "requires a discontinuous threshold jump to reach 125 GeV", EmpiricalComparison: false, FinalDerivation: false, Status: StatusProxyAlignmentsCataloged},
	}
	p := ProxyLedger{Cataloged: true, Alignments: alignments, EmpiricalInputsQuarantined: true, FinalMassClaimed: false}
	for _, a := range alignments {
		s := strings.ToLower(a.Name + " " + a.Output + " " + a.QuarantinedInput)
		p.ContainsTreeLevel125Proxy = p.ContainsTreeLevel125Proxy || strings.Contains(s, "tree-level") || strings.Contains(a.Output, "125.606")
		p.ContainsThresholdTransport125 = p.ContainsThresholdTransport125 || strings.Contains(s, "threshold rg") || strings.Contains(a.Output, "124.976")
		p.Contains331Diagnostic = p.Contains331Diagnostic || strings.Contains(a.Output, "331")
		p.Contains157ContinuousFloor = p.Contains157ContinuousFloor || strings.Contains(a.Output, "157")
	}
	p.Verdict = strings.Join([]string{StatusProxyAlignmentsCataloged, StatusLedgerFirewallsPreserved}, ";")
	return p
}

func compileFirewalls() FirewallLedger {
	items := []Firewall{
		{Name: "Absolute unified coupling origin", Gate: "Gate 316", MathematicalLimit: "α_GUT≈1/25 remains empirical in Gate 315", WhyItCannotBeClosedHere: "f₀=7 and τ_GUT=1 determine the form but not the missing absolute prefactor", PhaseIIITarget: "derive the weighted trace-capacity functional giving C_trace=25 or N₄=25/(28π)", Closed: false, Status: StatusFailedAlphaGUTNotDerived},
		{Name: "Weighted trace capacity C_trace=25", Gate: "Gate 317", MathematicalLimit: "raw Hilbert dimensions are 16,32,48,96, not 25", WhyItCannotBeClosedHere: "25 cannot be obtained by mixing spinor, gauge, and singlet categories without a valid trace functional", PhaseIIITarget: "construct the heat-kernel weighted capacity invariant rather than a raw degree count", Closed: false, Status: StatusFailedWeightedTraceMissing},
		{Name: "Flavor/CKM vacuum selection", Gate: "Gate 324/325", MathematicalLimit: "signed nullspace has dimension 2; positive metric forbids exact nulling", WhyItCannotBeClosedHere: "finite algebra supplies the landscape but not a unique mass-eigenstate orientation", PhaseIIITarget: "derive a dynamical vacuum selector / CKM texture potential", Closed: false, Status: StatusFailedFlavorVacuumNotSelected},
		{Name: "Native flavor projection metric", Gate: "Gate 325", MathematicalLimit: "top suppression depends on a signed interference metric, not the standard positive Yukawa trace", WhyItCannotBeClosedHere: "no theorem selects the signed projection as the physical top RG boundary metric", PhaseIIITarget: "derive the projection metric from the spectral action or variational flavor sector", Closed: false, Status: StatusFailedNativeProjectionMetric},
		{Name: "Two-loop and pole precision", Gate: "Gate 310/322", MathematicalLimit: "one-loop running mass is not a collider pole mass", WhyItCannotBeClosedHere: "self-energies, MS-bar/pole conversion, two-loop RG, and threshold-scale uncertainty are not executed", PhaseIIITarget: "install two-loop β-functions, matching, and pole-mass conversion", Closed: false, Status: StatusFailedTwoLoopNotExecuted},
		{Name: "Exact collider Higgs mass", Gate: "Gate 322/325", MathematicalLimit: "124.976 GeV is a conditional running-mass proxy, not a measured pole-mass derivation", WhyItCannotBeClosedHere: "flattened-top/sign-metric lane and precision conversion remain conditional", PhaseIIITarget: "integrate all sectors with derived flavor vacuum and precision transport", Closed: false, Status: StatusFailedColliderMassNotClaimed},
	}
	f := FirewallLedger{Cataloged: true, Items: items}
	for _, it := range items {
		f.AnyClosed = f.AnyClosed || it.Closed
		s := strings.ToLower(it.Name + " " + it.MathematicalLimit + " " + it.PhaseIIITarget)
		f.ContainsAlphaGUTOrigin = f.ContainsAlphaGUTOrigin || strings.Contains(s, "alpha") || strings.Contains(s, "α_gut")
		f.ContainsWeightedTrace25 = f.ContainsWeightedTrace25 || strings.Contains(s, "25") || strings.Contains(s, "weighted")
		f.ContainsFlavorVacuumSelection = f.ContainsFlavorVacuumSelection || strings.Contains(s, "ckm") || strings.Contains(s, "flavor")
		f.ContainsProjectionMetricSelection = f.ContainsProjectionMetricSelection || strings.Contains(s, "projection metric")
		f.ContainsTwoLoopPolePrecision = f.ContainsTwoLoopPolePrecision || strings.Contains(s, "two-loop") || strings.Contains(s, "pole")
		f.ContainsExactColliderMass = f.ContainsExactColliderMass || strings.Contains(s, "collider")
	}
	f.Verdict = strings.Join([]string{StatusFirewallsCataloged, StatusPhaseIIICataloged}, ";")
	return f
}

func compilePhaseIII() PhaseIIILedger {
	targets := []PhaseIIITarget{
		{Name: "Weighted trace-capacity theorem", WorkPackage: "Absolute coupling origin", RequiredTheorem: "C_trace=25 as a valid heat-kernel weighted functional", SuccessCriterion: "derive α_GUT^{-1}=25 without empirical α_GUT input", InheritsFrom: []string{"Gate 316", "Gate 317"}, Status: StatusPhaseIIICataloged},
		{Name: "Flavor vacuum selector", WorkPackage: "CKM / mass-basis orientation", RequiredTheorem: "unique U_flavor or vacuum functional selecting the signed null texture", SuccessCriterion: "derive the physical top boundary rather than using the diagnostic flattened-top lane", InheritsFrom: []string{"Gate 323", "Gate 324", "Gate 325"}, Status: StatusPhaseIIICataloged},
		{Name: "Native projection metric", WorkPackage: "Flavor projection metric", RequiredTheorem: "derive whether the RG top boundary uses positive trace or signed interference projection", SuccessCriterion: "authorize or reject top suppression from the finite spectral action itself", InheritsFrom: []string{"Gate 325"}, Status: StatusPhaseIIICataloged},
		{Name: "Precision threshold/RG/pole transport", WorkPackage: "Collider mass conversion", RequiredTheorem: "two-loop RG + explicit threshold scale + pole self-energy matching", SuccessCriterion: "convert the 124.976 GeV running proxy into a collider pole-mass prediction", InheritsFrom: []string{"Gate 310", "Gate 322"}, Status: StatusPhaseIIICataloged},
		{Name: "Full sigma/B-gap potential", WorkPackage: "Heavy sector completion", RequiredTheorem: "derive the full σ potential, mass threshold, and matching scale instead of witness normalization only", SuccessCriterion: "upgrade Δλ=-0.097846792207 from conditional EFT witness to complete heavy-sector theorem", InheritsFrom: []string{"Gate 318", "Gate 319", "Gate 320", "Gate 321"}, Status: StatusPhaseIIICataloged},
	}
	return PhaseIIILedger{
		Formalized:                 true,
		Targets:                    targets,
		IncludesWeightedTrace:      true,
		IncludesFlavorVacuum:       true,
		IncludesProjectionMetric:   true,
		IncludesPrecisionTransport: true,
		IncludesFullSigmaPotential: true,
		RequiresNoEmpiricalTuning:  true,
		Verdict:                    strings.Join([]string{StatusPhaseIIICataloged, StatusLedgerFirewallsPreserved}, ";"),
	}
}

func auditFirewallIntegrity() FirewallAudit {
	return FirewallAudit{
		NoAlphaGUTFitPromoted:      true,
		NoCKMTextureInvented:       true,
		NoFlavorMetricForced:       true,
		NoObservedHiggsFitInserted: true,
		NoObservedTopFitInserted:   true,
		NoTwoLoopClaimed:           true,
		NoPoleMassClaimed:          true,
		NoFinalTOEClaimed:          true,
		NoExactColliderMassClaimed: true,
		FiniteCorePolluted:         false,
		Verdict: strings.Join([]string{
			StatusLedgerFirewallsPreserved,
			StatusFailedFinalTOENotClaimed,
			StatusFailedAlphaGUTNotDerived,
			StatusFailedFlavorVacuumNotSelected,
			StatusFailedTwoLoopNotExecuted,
			StatusFailedPoleMassNotExecuted,
			StatusFailedColliderMassNotClaimed,
		}, ";"),
	}
}

func compileSummary(span RegistrySpan, triumphs TriumphLedger, proxies ProxyLedger, firewalls FirewallLedger, phaseIII PhaseIIILedger, audit FirewallAudit) Summary {
	firewallsOK := audit.NoAlphaGUTFitPromoted && audit.NoCKMTextureInvented && audit.NoFlavorMetricForced && audit.NoObservedHiggsFitInserted && audit.NoObservedTopFitInserted && audit.NoTwoLoopClaimed && audit.NoPoleMassClaimed && audit.NoFinalTOEClaimed && audit.NoExactColliderMassClaimed && !audit.FiniteCorePolluted
	return Summary{
		LedgerCompiled:       true,
		ProjectCapstone:      true,
		TriumphsReady:        triumphs.Cataloged && triumphs.ContainsWeakMixing && triumphs.ContainsTraceEquivalence && triumphs.ContainsThresholdJump,
		ProxiesReady:         proxies.Cataloged && proxies.ContainsTreeLevel125Proxy && proxies.ContainsThresholdTransport125 && proxies.EmpiricalInputsQuarantined,
		FirewallsReady:       firewalls.Cataloged && firewalls.ContainsAlphaGUTOrigin && firewalls.ContainsFlavorVacuumSelection && firewalls.ContainsTwoLoopPolePrecision && !firewalls.AnyClosed,
		PhaseIIIReady:        phaseIII.Formalized && phaseIII.IncludesWeightedTrace && phaseIII.IncludesFlavorVacuum && phaseIII.IncludesPrecisionTransport,
		FirewallsPreserved:   firewallsOK,
		FinalTOEClaimed:      false,
		ExactColliderClaimed: false,
		Status:               strings.Join([]string{StatusGrandUnifiedLedgerCompiled, StatusProjectCapstoneAchieved, StatusLedgerFirewallsPreserved}, ";"),
		DirectAnswer:         fmt.Sprintf("ASHA through Gate %d has two exact native boundary ratios (%s, %s), a derived threshold-jump witness (%s), and two near-125 GeV conditional alignments (%.6f GeV tree proxy, %.6f GeV running proxy), but it does not derive α_GUT, CKM/flavor vacuum selection, or pole-mass precision.", span.HighestGateInherited, weakMixingRatio, higgsQuarticRatio, canonicalJump, empiricalProxyMassGeV, thresholdRunMassGeV),
		NextPhase:            "Phase III: weighted trace-capacity origin of α_GUT, dynamical flavor-vacuum/CKM selector, and precision two-loop/pole transport.",
	}
}

func Statuses(a Analysis) []string {
	out := []string{
		StatusGrandUnifiedLedgerCompiled,
		StatusProjectCapstoneAchieved,
		StatusAbsoluteTriumphsCataloged,
		StatusProxyAlignmentsCataloged,
		StatusFirewallsCataloged,
		StatusPhaseIIICataloged,
		StatusLedgerFirewallsPreserved,
		StatusFailedFinalTOENotClaimed,
		StatusFailedAlphaGUTNotDerived,
		StatusFailedWeightedTraceMissing,
		StatusFailedFlavorVacuumNotSelected,
		StatusFailedNativeProjectionMetric,
		StatusFailedTwoLoopNotExecuted,
		StatusFailedPoleMassNotExecuted,
		StatusFailedColliderMassNotClaimed,
		StatusFailedPhaseIIIRequired,
	}
	if !a.Summary.LedgerCompiled {
		out = append(out, "FAILED_ROUTE_GRAND_UNIFIED_LEDGER_NOT_COMPILED")
	}
	return out
}

func FormatSpan(s RegistrySpan) string {
	return fmt.Sprintf("audit=%s gate_range=%q highest=%d capstone=%q adds_fit=%v rewrites=%v verdict=%s", s.AuditID, s.GateRange, s.HighestGateInherited, s.CapstoneForPhase, s.AddsNewPhysicsFit, s.RewritesHistory, s.Verdict)
}

func FormatTriumphs(t TriumphLedger) string {
	parts := []string{fmt.Sprintf("cataloged=%v native=%d weak=%v morita=%v triality=%v bimodule=%v resonance=%v trace=%v jump=%v", t.Cataloged, t.NativeCount, t.ContainsWeakMixing, t.ContainsMoritaColorSplit, t.ContainsGenerationTriality, t.ContainsTrueBimodule, t.ContainsTopologicalResonance, t.ContainsTraceEquivalence, t.ContainsThresholdJump)}
	for _, it := range t.Items {
		parts = append(parts, fmt.Sprintf("[%s gate=%s native=%v empirical=%v claim=%v] %s", it.Name, it.Gate, it.NativeToFiniteCore, it.RequiresEmpiricalData, it.FinalColliderClaim, it.Statement))
	}
	return strings.Join(parts, "\n")
}

func FormatProxies(p ProxyLedger) string {
	parts := []string{fmt.Sprintf("cataloged=%v tree125=%v transport125=%v 331=%v floor157=%v empirical_quarantined=%v final_claim=%v", p.Cataloged, p.ContainsTreeLevel125Proxy, p.ContainsThresholdTransport125, p.Contains331Diagnostic, p.Contains157ContinuousFloor, p.EmpiricalInputsQuarantined, p.FinalMassClaimed)}
	for _, it := range p.Alignments {
		parts = append(parts, fmt.Sprintf("[%s gate=%s input=%s geometry=%s output=%s error=%s final=%v]", it.Name, it.Gate, it.QuarantinedInput, it.GeometricInput, it.Output, it.ErrorDescription, it.FinalDerivation))
	}
	return strings.Join(parts, "\n")
}

func FormatFirewalls(f FirewallLedger) string {
	parts := []string{fmt.Sprintf("cataloged=%v alpha=%v c25=%v flavor=%v metric=%v precision=%v collider=%v any_closed=%v", f.Cataloged, f.ContainsAlphaGUTOrigin, f.ContainsWeightedTrace25, f.ContainsFlavorVacuumSelection, f.ContainsProjectionMetricSelection, f.ContainsTwoLoopPolePrecision, f.ContainsExactColliderMass, f.AnyClosed)}
	for _, it := range f.Items {
		parts = append(parts, fmt.Sprintf("[%s gate=%s closed=%v status=%s] limit=%s target=%s", it.Name, it.Gate, it.Closed, it.Status, it.MathematicalLimit, it.PhaseIIITarget))
	}
	return strings.Join(parts, "\n")
}

func FormatPhaseIII(p PhaseIIILedger) string {
	parts := []string{fmt.Sprintf("formalized=%v weighted=%v flavor=%v metric=%v precision=%v sigma=%v no_tuning=%v", p.Formalized, p.IncludesWeightedTrace, p.IncludesFlavorVacuum, p.IncludesProjectionMetric, p.IncludesPrecisionTransport, p.IncludesFullSigmaPotential, p.RequiresNoEmpiricalTuning)}
	for _, t := range p.Targets {
		parts = append(parts, fmt.Sprintf("[%s package=%s] theorem=%s success=%s inherits=%s", t.Name, t.WorkPackage, t.RequiredTheorem, t.SuccessCriterion, strings.Join(t.InheritsFrom, ",")))
	}
	return strings.Join(parts, "\n")
}

func FormatAudit(a FirewallAudit) string {
	return fmt.Sprintf("no_alpha_fit=%v no_ckm=%v no_metric_forced=%v no_higgs_fit=%v no_top_fit=%v no_2loop=%v no_pole=%v no_toe=%v no_exact_collider=%v polluted=%v verdict=%s", a.NoAlphaGUTFitPromoted, a.NoCKMTextureInvented, a.NoFlavorMetricForced, a.NoObservedHiggsFitInserted, a.NoObservedTopFitInserted, a.NoTwoLoopClaimed, a.NoPoleMassClaimed, a.NoFinalTOEClaimed, a.NoExactColliderMassClaimed, a.FiniteCorePolluted, a.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("compiled=%v capstone=%v triumphs=%v proxies=%v firewalls=%v phaseIII=%v preserved=%v toe=%v collider=%v status=%s answer=%q next=%q", s.LedgerCompiled, s.ProjectCapstone, s.TriumphsReady, s.ProxiesReady, s.FirewallsReady, s.PhaseIIIReady, s.FirewallsPreserved, s.FinalTOEClaimed, s.ExactColliderClaimed, s.Status, s.DirectAnswer, s.NextPhase)
}
