// Package nativeunifiedcouplingorigin implements Gate 316:
// Native Unified Coupling Origin / Absolute Gauge Coupling Trace-Capacity Audit.
//
// Gate 315 verified the algebraic Higgs-to-gauge ratio
//
//	lambda_H / g_*^2 = 1197/4624
//
// against a quarantined empirical comparison input alpha_GUT=1/25.  Gate 316
// audits whether the absolute unified coupling itself can be derived from the
// finite Cℓ(1,7) spectral geometry instead of being borrowed from phenomenology.
//
// The result is deliberately strict.  The gate formalizes the absolute gauge
// normalization equation
//
//	1/g_*^2 = N4 f0 tau_GUT,
//	alpha_GUT^{-1} = 4*pi*N4*f0*tau_GUT,
//
// inserts the previously promoted contact cutoff moment f0=7 and the Gate-308
// GUT trace index tau_GUT=1, and computes the exact remaining prefactor required
// to obtain alpha_GUT^{-1}=25.  It then audits finite trace-capacity candidates.
// The package reconstructs the target and the missing mathematical obligation;
// it does not declare alpha_GUT derived from the finite core.
package nativeunifiedcouplingorigin

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE316-NATIVE-UNIFIED-COUPLING-ORIGIN"

	StatusGate315RatioContextInherited      = "CONDITIONAL_SUPPORT_GATE315_RATIO_CONTEXT_INHERITED"
	StatusGaugeKineticAbsoluteMapFormalized = "CONDITIONAL_SUPPORT_GAUGE_KINETIC_ABSOLUTE_NORMALIZATION_MAP_FORMALIZED"
	StatusContactF0AndTauLedgerApplied      = "CONDITIONAL_SUPPORT_CONTACT_F0_AND_TAU_GUT_LEDGER_APPLIED"
	StatusAlphaGUTTargetReconstructed       = "CONDITIONAL_SUPPORT_ALPHA_GUT_TARGET_RECONSTRUCTED"
	StatusRequiredN4CapacityComputed        = "CONDITIONAL_SUPPORT_REQUIRED_N4_CAPACITY_COMPUTED"
	StatusTraceCapacityCandidatesAudited    = "CONDITIONAL_SUPPORT_TRACE_CAPACITY_CANDIDATES_AUDITED"
	StatusNativeCouplingAuditCompleted      = "CONDITIONAL_SUPPORT_NATIVE_UNIFIED_COUPLING_ORIGIN_AUDIT_COMPLETED"
	StatusGate316FirewallsPreserved         = "CONDITIONAL_SUPPORT_GATE316_FIREWALLS_PRESERVED"

	StatusTensionContactF0AloneInsufficient    = "CONDITIONAL_TENSION_CONTACT_F0_ALONE_DOES_NOT_FIX_ALPHA_GUT"
	StatusTensionCapacity25CandidateUnselected = "CONDITIONAL_TENSION_TRACE_CAPACITY_25_CANDIDATE_UNSELECTED"

	StatusFailedAlphaGUTAbsoluteStillSealed     = "FAILED_ROUTE_ALPHA_GUT_ABSOLUTE_VALUE_STILL_SEALED"
	StatusFailedNativeTraceCapacity25NotDerived = "FAILED_ROUTE_NATIVE_TRACE_CAPACITY_25_NOT_DERIVED"
	StatusFailedSeeleyPrefactorNotFiniteNative  = "FAILED_ROUTE_SEELEY_DEWITT_PREFACTOR_NOT_NATIVE_FINITE_CORE"
	StatusFailedRenormalizationSchemeRequired   = "FAILED_ROUTE_CONTINUUM_RENORMALIZATION_SCHEME_REQUIRED"
	StatusFailedHiggsProxyNotUpgraded           = "FAILED_ROUTE_HIGGS_PROXY_NOT_UPGRADED_TO_DERIVATION"
)

const (
	contactF0             = 7.0
	tauGUT                = 1.0
	alphaInverseTarget    = 25.0
	alphaTarget           = 1.0 / alphaInverseTarget
	traceRatioNumerator   = 1197.0
	traceRatioDenominator = 4624.0
)

type RatioContext struct {
	HiggsGaugeRatioEquation string
	RatioNumerator          int
	RatioDenominator        int
	Ratio                   float64
	AlphaGUTFromGate315     string
	EmpiricalOnly           bool
	Verdict                 string
}

type GaugeKineticMap struct {
	CouplingEquation      string
	AlphaInverseEquation  string
	F0                    float64
	TauGUT                float64
	N4Symbol              string
	ContactMomentPromoted bool
	TauLedgerApplied      bool
	Verdict               string
}

type AbsoluteTarget struct {
	TargetAlphaExpression         string
	TargetAlpha                   float64
	TargetAlphaInverse            float64
	TargetGStarSquaredExpression  string
	TargetGStarSquared            float64
	TargetInverseGStarSquared     float64
	ReconstructedFromGate315Input bool
	DerivedFromFiniteCore         bool
	Verdict                       string
}

type PrefactorRequirement struct {
	RequiredN4Expression          string
	RequiredN4                    float64
	RequiredN4TimesF0Tau          float64
	AlphaInverseIfN4EqualsOne     float64
	ContactF0AloneMatchesTarget   bool
	MissingPrefactorMatchesTarget bool
	Verdict                       string
}

type CapacityCandidate struct {
	Name                string
	Formula             string
	Value               float64
	Positive            bool
	Equals25            bool
	CanonicallySelected bool
	Comment             string
	Status              string
}

type TraceCapacityAudit struct {
	Candidates                   []CapacityCandidate
	HasInteger25Candidate        bool
	HasCanonicalNativeDerivation bool
	RequiredCapacity             string
	Verdict                      string
}

type HiggsProxyRecheck struct {
	LambdaRatio                float64
	TargetGStarSquared         float64
	LambdaFromTargetAlpha      float64
	TreeMassProxyGeV           float64
	SameAsGate315Proxy         bool
	UpgradedToNativeDerivation bool
	Verdict                    string
}

type FirewallAudit struct {
	NoAlphaGUTDerivationClaimed  bool
	NoForcedCapacitySelection    bool
	NoContinuumPrefactorInvented bool
	NoHiggsMassDerivationClaimed bool
	FiniteCorePolluted           bool
	Obligations                  []string
	Verdict                      string
}

type Summary struct {
	AbsoluteMapFormalized      bool
	F0AndTauApplied            bool
	TargetReconstructed        bool
	RequiredPrefactorComputed  bool
	TraceCandidatesAudited     bool
	NativeAlphaDerived         bool
	HiggsProxyStillConditional bool
	FirewallsPreserved         bool
	Status                     string
	DirectAnswer               string
	NextGate                   string
}

type Analysis struct {
	Context     RatioContext
	Map         GaugeKineticMap
	Target      AbsoluteTarget
	Requirement PrefactorRequirement
	Capacity    TraceCapacityAudit
	HiggsProxy  HiggsProxyRecheck
	Firewalls   FirewallAudit
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
	context := inheritRatioContext()
	gmap := buildGaugeKineticMap()
	target := reconstructTarget()
	requirement := computePrefactorRequirement(gmap, target)
	capacity := auditTraceCapacity(requirement)
	proxy := recheckHiggsProxy(context, target, capacity)
	firewalls := auditFirewalls(target, requirement, capacity, proxy)
	summary := buildSummary(gmap, target, requirement, capacity, proxy, firewalls)
	truth := "Gate 316 formalizes the absolute unified-coupling equation alpha_GUT^{-1}=4*pi*N4*f0*tau_GUT.  With f0=7 and tau_GUT=1, matching the empirical comparison value alpha_GUT^{-1}=25 requires N4=25/(28*pi).  The audit reconstructs the target exactly and identifies trace-capacity 25 as the missing finite-core obligation, but no existing finite theorem canonically derives that capacity or the continuum prefactor.  Therefore alpha_GUT remains sealed, while the Gate-315 Higgs-ratio proxy remains a strong empirical comparison rather than a native derivation."
	return Analysis{Context: context, Map: gmap, Target: target, Requirement: requirement, Capacity: capacity, HiggsProxy: proxy, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func inheritRatioContext() RatioContext {
	ratio := traceRatioNumerator / traceRatioDenominator
	return RatioContext{
		HiggsGaugeRatioEquation: "lambda_H/g_*^2 = 1197/4624",
		RatioNumerator:          int(traceRatioNumerator),
		RatioDenominator:        int(traceRatioDenominator),
		Ratio:                   ratio,
		AlphaGUTFromGate315:     "alpha_GUT = 1/25",
		EmpiricalOnly:           true,
		Verdict:                 StatusGate315RatioContextInherited,
	}
}

func buildGaugeKineticMap() GaugeKineticMap {
	return GaugeKineticMap{
		CouplingEquation:      "1/g_*^2 = N4 * f0 * tau_GUT",
		AlphaInverseEquation:  "alpha_GUT^{-1} = 4*pi*N4*f0*tau_GUT",
		F0:                    contactF0,
		TauGUT:                tauGUT,
		N4Symbol:              "N4",
		ContactMomentPromoted: true,
		TauLedgerApplied:      true,
		Verdict:               strings.Join([]string{StatusGaugeKineticAbsoluteMapFormalized, StatusContactF0AndTauLedgerApplied}, ";"),
	}
}

func reconstructTarget() AbsoluteTarget {
	g2 := 4.0 * math.Pi * alphaTarget
	return AbsoluteTarget{
		TargetAlphaExpression:         "alpha_GUT = 1/25",
		TargetAlpha:                   alphaTarget,
		TargetAlphaInverse:            alphaInverseTarget,
		TargetGStarSquaredExpression:  "g_*^2 = 4*pi/25",
		TargetGStarSquared:            g2,
		TargetInverseGStarSquared:     1.0 / g2,
		ReconstructedFromGate315Input: true,
		DerivedFromFiniteCore:         false,
		Verdict:                       strings.Join([]string{StatusAlphaGUTTargetReconstructed, StatusFailedAlphaGUTAbsoluteStillSealed}, ";"),
	}
}

func computePrefactorRequirement(m GaugeKineticMap, t AbsoluteTarget) PrefactorRequirement {
	requiredN4 := t.TargetAlphaInverse / (4.0 * math.Pi * m.F0 * m.TauGUT)
	n4f0tau := requiredN4 * m.F0 * m.TauGUT
	alphaInvN4One := 4.0 * math.Pi * m.F0 * m.TauGUT
	return PrefactorRequirement{
		RequiredN4Expression:          "N4_required = 25/(4*pi*f0*tau_GUT) = 25/(28*pi)",
		RequiredN4:                    requiredN4,
		RequiredN4TimesF0Tau:          n4f0tau,
		AlphaInverseIfN4EqualsOne:     alphaInvN4One,
		ContactF0AloneMatchesTarget:   nearlyEqual(alphaInvN4One, t.TargetAlphaInverse, 1e-12),
		MissingPrefactorMatchesTarget: nearlyEqual(4.0*math.Pi*n4f0tau, t.TargetAlphaInverse, 1e-12),
		Verdict:                       strings.Join([]string{StatusRequiredN4CapacityComputed, StatusTensionContactF0AloneInsufficient}, ";"),
	}
}

func auditTraceCapacity(r PrefactorRequirement) TraceCapacityAudit {
	candidates := []CapacityCandidate{
		{Name: "contact cutoff moment", Formula: "f0", Value: contactF0, Positive: true, Equals25: false, CanonicallySelected: true, Comment: "locks the a4 moment but gives only seven units before the continuum prefactor", Status: StatusContactF0AndTauLedgerApplied},
		{Name: "target inverse alpha capacity", Formula: "alpha_GUT^{-1}", Value: alphaInverseTarget, Positive: true, Equals25: true, CanonicallySelected: false, Comment: "exact capacity required by empirical unification comparison; not yet derived from finite trace capacity", Status: StatusTensionCapacity25CandidateUnselected},
		{Name: "required continuum-normalized prefactor", Formula: "N4_required = 25/(28*pi)", Value: r.RequiredN4, Positive: true, Equals25: false, CanonicallySelected: false, Comment: "contains the continuum 4*pi convention and cannot be pure finite integer data by itself", Status: StatusFailedSeeleyPrefactorNotFiniteNative},
	}
	has25 := false
	canonical := false
	for _, c := range candidates {
		has25 = has25 || c.Equals25
		canonical = canonical || (c.Equals25 && c.CanonicallySelected)
	}
	return TraceCapacityAudit{
		Candidates:                   candidates,
		HasInteger25Candidate:        has25,
		HasCanonicalNativeDerivation: canonical,
		RequiredCapacity:             "derive C_trace=25 or equivalently N4=25/(28*pi) from finite trace capacity plus continuum normalization",
		Verdict:                      strings.Join([]string{StatusTraceCapacityCandidatesAudited, StatusFailedNativeTraceCapacity25NotDerived, StatusFailedAlphaGUTAbsoluteStillSealed}, ";"),
	}
}

func recheckHiggsProxy(c RatioContext, t AbsoluteTarget, a TraceCapacityAudit) HiggsProxyRecheck {
	lambda := c.Ratio * t.TargetGStarSquared
	mass := 246.22 * math.Sqrt(2.0*lambda)
	return HiggsProxyRecheck{
		LambdaRatio:                c.Ratio,
		TargetGStarSquared:         t.TargetGStarSquared,
		LambdaFromTargetAlpha:      lambda,
		TreeMassProxyGeV:           mass,
		SameAsGate315Proxy:         mass > 125.0 && mass < 126.5,
		UpgradedToNativeDerivation: t.DerivedFromFiniteCore && a.HasCanonicalNativeDerivation,
		Verdict:                    strings.Join([]string{StatusAlphaGUTTargetReconstructed, StatusFailedHiggsProxyNotUpgraded}, ";"),
	}
}

func auditFirewalls(t AbsoluteTarget, r PrefactorRequirement, c TraceCapacityAudit, p HiggsProxyRecheck) FirewallAudit {
	obligations := []string{
		StatusFailedAlphaGUTAbsoluteStillSealed + ": alpha_GUT=1/25 is reconstructed as the required target, not derived",
		StatusFailedNativeTraceCapacity25NotDerived + ": no theorem selects trace capacity 25 from the finite Hilbert trace",
		StatusFailedSeeleyPrefactorNotFiniteNative + ": N4 includes continuum heat-kernel and 4*pi conventions",
		StatusFailedRenormalizationSchemeRequired + ": absolute couplings require scheme and scale selection",
		StatusFailedHiggsProxyNotUpgraded + ": the 125.6 GeV proxy remains conditional on empirical alpha_GUT",
	}
	return FirewallAudit{
		NoAlphaGUTDerivationClaimed:  !t.DerivedFromFiniteCore,
		NoForcedCapacitySelection:    !c.HasCanonicalNativeDerivation,
		NoContinuumPrefactorInvented: r.RequiredN4 > 0 && r.MissingPrefactorMatchesTarget,
		NoHiggsMassDerivationClaimed: !p.UpgradedToNativeDerivation,
		FiniteCorePolluted:           false,
		Obligations:                  obligations,
		Verdict:                      strings.Join([]string{StatusGate316FirewallsPreserved, StatusFailedAlphaGUTAbsoluteStillSealed, StatusFailedHiggsProxyNotUpgraded}, ";"),
	}
}

func buildSummary(m GaugeKineticMap, t AbsoluteTarget, r PrefactorRequirement, c TraceCapacityAudit, p HiggsProxyRecheck, f FirewallAudit) Summary {
	native := t.DerivedFromFiniteCore && c.HasCanonicalNativeDerivation
	status := StatusFailedAlphaGUTAbsoluteStillSealed
	if native {
		status = "CONDITIONAL_SUPPORT_NATIVE_UNIFIED_COUPLING_DERIVED"
	}
	return Summary{
		AbsoluteMapFormalized:      m.ContactMomentPromoted && m.TauLedgerApplied,
		F0AndTauApplied:            m.F0 == contactF0 && m.TauGUT == tauGUT,
		TargetReconstructed:        t.ReconstructedFromGate315Input && t.TargetAlphaInverse == alphaInverseTarget,
		RequiredPrefactorComputed:  r.RequiredN4 > 0 && r.MissingPrefactorMatchesTarget,
		TraceCandidatesAudited:     len(c.Candidates) >= 3 && c.HasInteger25Candidate,
		NativeAlphaDerived:         native,
		HiggsProxyStillConditional: !p.UpgradedToNativeDerivation,
		FirewallsPreserved:         f.NoAlphaGUTDerivationClaimed && f.NoForcedCapacitySelection && f.NoContinuumPrefactorInvented && f.NoHiggsMassDerivationClaimed && !f.FiniteCorePolluted,
		Status:                     status,
		DirectAnswer:               "Gate 316 derives the required absolute-normalization target, not the absolute coupling itself: with f0=7 and tau_GUT=1, alpha_GUT^{-1}=25 requires N4=25/(28*pi), so a native trace-capacity theorem selecting 25 is still missing.",
		NextGate:                   "Phase-II should either derive the trace-capacity 25 theorem or move to the heavy-threshold portal tensor with alpha_GUT kept quarantined.",
	}
}

func nearlyEqual(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatRatioContext(x RatioContext) string {
	return fmt.Sprintf("equation=%s; ratio=%d/%d=%.12f; alpha_source=%s; empirical_only=%t; verdict=%s", x.HiggsGaugeRatioEquation, x.RatioNumerator, x.RatioDenominator, x.Ratio, x.AlphaGUTFromGate315, x.EmpiricalOnly, x.Verdict)
}

func FormatGaugeKineticMap(x GaugeKineticMap) string {
	return fmt.Sprintf("%s; %s; f0=%.0f; tau_GUT=%.0f; contact_promoted=%t; tau_applied=%t; verdict=%s", x.CouplingEquation, x.AlphaInverseEquation, x.F0, x.TauGUT, x.ContactMomentPromoted, x.TauLedgerApplied, x.Verdict)
}

func FormatTarget(x AbsoluteTarget) string {
	return fmt.Sprintf("%s=%.12f; alpha_inv=%.12f; %s=%.12f; inv_g2=%.12f; reconstructed=%t; derived_from_core=%t; verdict=%s", x.TargetAlphaExpression, x.TargetAlpha, x.TargetAlphaInverse, x.TargetGStarSquaredExpression, x.TargetGStarSquared, x.TargetInverseGStarSquared, x.ReconstructedFromGate315Input, x.DerivedFromFiniteCore, x.Verdict)
}

func FormatRequirement(x PrefactorRequirement) string {
	return fmt.Sprintf("%s=%.12f; N4*f0*tau=%.12f; alpha_inv_if_N4_1=%.12f; f0_alone_matches=%t; required_prefactor_matches=%t; verdict=%s", x.RequiredN4Expression, x.RequiredN4, x.RequiredN4TimesF0Tau, x.AlphaInverseIfN4EqualsOne, x.ContactF0AloneMatchesTarget, x.MissingPrefactorMatchesTarget, x.Verdict)
}

func FormatCapacity(x TraceCapacityAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, fmt.Sprintf("%s:%s=%.12f canon=%t status=%s", c.Name, c.Formula, c.Value, c.CanonicallySelected, c.Status))
	}
	return fmt.Sprintf("candidates=[%s]; has25=%t; canonical25=%t; required=%s; verdict=%s", strings.Join(parts, "; "), x.HasInteger25Candidate, x.HasCanonicalNativeDerivation, x.RequiredCapacity, x.Verdict)
}

func FormatHiggsProxy(x HiggsProxyRecheck) string {
	return fmt.Sprintf("ratio=%.12f; gstar2=%.12f; lambda=%.12f; m_tree=%.6f GeV; same_as_gate315=%t; native_derivation=%t; verdict=%s", x.LambdaRatio, x.TargetGStarSquared, x.LambdaFromTargetAlpha, x.TreeMassProxyGeV, x.SameAsGate315Proxy, x.UpgradedToNativeDerivation, x.Verdict)
}

func FormatFirewalls(x FirewallAudit) string {
	return fmt.Sprintf("no_alpha_claim=%t; no_forced_capacity=%t; no_prefactor_invented=%t; no_mass_derivation=%t; polluted=%t; obligations=[%s]; verdict=%s", x.NoAlphaGUTDerivationClaimed, x.NoForcedCapacitySelection, x.NoContinuumPrefactorInvented, x.NoHiggsMassDerivationClaimed, x.FiniteCorePolluted, strings.Join(x.Obligations, "; "), x.Verdict)
}

func FormatSummary(x Summary) string {
	return fmt.Sprintf("map=%t; f0_tau=%t; target=%t; prefactor=%t; candidates=%t; native_alpha=%t; proxy_conditional=%t; firewalls=%t; status=%s; answer=%s; next=%s", x.AbsoluteMapFormalized, x.F0AndTauApplied, x.TargetReconstructed, x.RequiredPrefactorComputed, x.TraceCandidatesAudited, x.NativeAlphaDerived, x.HiggsProxyStillConditional, x.FirewallsPreserved, x.Status, x.DirectAnswer, x.NextGate)
}
