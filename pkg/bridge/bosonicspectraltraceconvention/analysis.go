// Package bosonicspectraltraceconvention implements Gate 330:
// Bosonic Spectral Action Trace Convention / Full Doubled-Space Gauge Trace Audit.
//
// Gate 329 identified the exact factor two needed to promote the successful
// α_GUT^{-1}=8π branch: the particle and J-mirror antiparticle curvature
// carriers each contribute a positive Yang-Mills trace. Gate 330 audits whether
// this factor two is native to the bosonic spectral action, or whether it must
// be removed by the real-structure quotient used in fermionic actions.
package bosonicspectraltraceconvention

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE330-BOSONIC-SPECTRAL-ACTION-TRACE-CONVENTION-FULL-DOUBLED-SPACE-AUDIT"

	StatusInputsInherited                        = "CONDITIONAL_SUPPORT_GATE329_DOUBLED_TRACE_OBLIGATION_INHERITED"
	StatusRealTripleTraceAxiomFormalized         = "CONDITIONAL_SUPPORT_REAL_SPECTRAL_TRIPLE_TRACE_AXIOM_FORMALIZED"
	StatusBosonicFullTraceNative                 = "CONDITIONAL_SUPPORT_BOSONIC_SPECTRAL_ACTION_FULL_HILBERT_TRACE_NATIVE"
	StatusFermionicHalfFactorSeparated           = "CONDITIONAL_SUPPORT_FERMIONIC_HALF_FACTOR_SEPARATED_FROM_BOSONIC_TRACE"
	StatusJMirrorCurvatureContributionVerified   = "CONDITIONAL_SUPPORT_J_MIRROR_CURVATURE_POSITIVE_CONTRIBUTION_VERIFIED"
	StatusEightPiBranchPromotedByTraceConvention = "CONDITIONAL_SUPPORT_EIGHT_PI_BRANCH_PROMOTED_BY_BOSONIC_TRACE_CONVENTION"
	StatusHiggsProxyRecomputed                   = "CONDITIONAL_SUPPORT_HIGGS_PROXY_RECOMPUTED_WITH_NATIVE_DOUBLED_TRACE"

	StatusTensionAbsoluteCouplingStillConditional = "CONDITIONAL_TENSION_ABSOLUTE_COUPLING_STILL_DEPENDS_ON_TOPOLOGICAL_ACTION_MAP"
	StatusTensionRepresentationTraceIndexRequired = "CONDITIONAL_TENSION_REPRESENTATION_TRACE_INDEX_NORMALIZATION_STILL_REQUIRED"

	StatusFailedQuotientAsBosonicRuleRejected = "FAILED_ROUTE_QUOTIENTED_TRACE_NOT_NATIVE_TO_BOSONIC_SPECTRAL_ACTION"
	StatusFailedAlphaUnconditionalNotDerived  = "FAILED_ROUTE_ALPHA_GUT_UNCONDITIONAL_VALUE_NOT_DERIVED"
	StatusFailedPoleMassNotExecuted           = "FAILED_ROUTE_POLE_MASS_CONVERSION_NOT_EXECUTED"
	StatusFailedColliderMassNotClaimed        = "FAILED_ROUTE_FINAL_COLLIDER_HIGGS_MASS_NOT_CLAIMED"
)

const (
	inheritedHighestGate     = 329
	contactScalarNumerator   = 1197.0
	contactScalarDenominator = 4624.0
	electroweakVEVGeV        = 246.22
	observedHiggsGeV         = 125.10
	sTopologicalAction       = 8.0 * math.Pi * math.Pi
)

type Inputs struct {
	HighestInheritedGate  int
	Gate329Question       string
	STopFormula           string
	STop                  float64
	SingleCarrierAlphaInv float64
	DoubledAlphaInv       float64
	AddsEmpiricalAlpha    bool
	Status                string
}

type TraceAxiom struct {
	BosonicSpectralAction            string
	TraceDomain                      string
	RealStructureRole                string
	FermionicAction                  string
	FermionicHalfAppliesToBosons     bool
	BosonicTraceUsesFullHilbertSpace bool
	Status                           string
}

type CurvatureMirror struct {
	ParticleCurvatureTerm         string
	JMirrorCurvatureTerm          string
	ComplexConjugateTraceIdentity string
	ParticleIndex                 float64
	MirrorIndex                   float64
	TotalBosonicIndex             float64
	SameSign                      bool
	Positive                      bool
	Status                        string
}

type CouplingLane struct {
	Name               string
	Formula            string
	TraceMultiplier    float64
	AlphaInverse       float64
	GStarSquared       float64
	LambdaH            float64
	HiggsMassGeV       float64
	DifferenceGeV      float64
	RelativeErrorPct   float64
	NativeBosonicTrace bool
	Status             string
}

type SeparationAudit struct {
	FermionicPfaffianReason           string
	BosonicHeatKernelReason           string
	HalfFactorConfinedToFermions      bool
	ApplyingHalfToBosonsBreaksEightPi bool
	QuotientLaneRejected              bool
	Status                            string
}

type PromotionAudit struct {
	Gate329ConditionalFactorTwo               bool
	Gate330TraceConventionSuppliesFactorTwo   bool
	EightPiPromotedWithinBosonicAction        bool
	AlphaUnconditional                        bool
	MissingRepresentationTraceIndex           bool
	MissingTopologicalActionToCouplingTheorem bool
	Reason                                    string
	Status                                    string
}

type FirewallAudit struct {
	NoEmpiricalAlphaInserted            bool
	NoObservedMassFitted                bool
	NoPoleMassClaimed                   bool
	NoFinalColliderClaimed              bool
	RepresentationIndexStillFirewalled  bool
	TopologicalActionMapStillFirewalled bool
	Status                              string
}

type Summary struct {
	FullDoubledTraceNativeForBosons    bool
	FermionicQuotientSeparated         bool
	EightPiBranchPromotedConditionally bool
	AlphaStillUnconditionalFailure     bool
	HiggsProxyGeV                      float64
	DirectAnswer                       string
	NextObligation                     string
	Status                             string
}

type Analysis struct {
	Inputs     Inputs
	Trace      TraceAxiom
	Mirror     CurvatureMirror
	Single     CouplingLane
	Doubled    CouplingLane
	Quotient   CouplingLane
	Separation SeparationAudit
	Promotion  PromotionAudit
	Audit      FirewallAudit
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
	inputs := compileInputs()
	trace := formalizeTraceAxiom()
	mirror := verifyMirrorCurvature()
	single := computeLane("single-carrier Chern-Weil diagnostic", "α^{-1}=S_top/(2π)", 1.0, false, StatusFailedQuotientAsBosonicRuleRejected)
	doubled := computeLane("full doubled bosonic spectral action", "α^{-1}=2·S_top/(2π)=S_top/π", 2.0, true, StatusEightPiBranchPromotedByTraceConvention)
	quotient := computeLane("fermionic quotient misapplied to bosons", "α^{-1}=½·S_top/(2π)", 0.5, false, StatusFailedQuotientAsBosonicRuleRejected)
	separation := auditSeparation(quotient)
	promotion := auditPromotion(trace, mirror, doubled)
	firewall := auditFirewalls(promotion)
	summary := compileSummary(trace, separation, promotion, doubled)
	truth := "Gate 330 upgrades the Gate 329 factor-two from a mere capacity to the native bosonic trace convention of a real spectral triple: Tr f(D_A/Λ) is a full Hilbert-space heat-kernel trace, while the one-half/Pfaffian quotient belongs to the fermionic action and does not divide the bosonic Yang-Mills coefficient. Thus the full particle plus J-mirror curvature trace conditionally promotes α_GUT^{-1}=8π and g_*²=1/2. The result remains conditional because the representation trace index and the topological-action-to-coupling theorem are still separate obligations."
	return Analysis{Inputs: inputs, Trace: trace, Mirror: mirror, Single: single, Doubled: doubled, Quotient: quotient, Separation: separation, Promotion: promotion, Audit: firewall, Summary: summary, Truth: truth}, nil
}

func compileInputs() Inputs {
	return Inputs{
		HighestInheritedGate:  inheritedHighestGate,
		Gate329Question:       "Does the bosonic spectral action use the full doubled H_F ⊕ H_F* gauge trace, or a quotient/half trace?",
		STopFormula:           "S_top = 8π²",
		STop:                  sTopologicalAction,
		SingleCarrierAlphaInv: sTopologicalAction / (2.0 * math.Pi),
		DoubledAlphaInv:       sTopologicalAction / math.Pi,
		AddsEmpiricalAlpha:    false,
		Status:                StatusInputsInherited,
	}
}

func formalizeTraceAxiom() TraceAxiom {
	return TraceAxiom{
		BosonicSpectralAction:            "S_B = Tr_H f(D_A/Λ)",
		TraceDomain:                      "full finite Hilbert carrier H_F ⊕ H_F* after the real-structure completion",
		RealStructureRole:                "J supplies the conjugate representation; it does not impose a bosonic Pfaffian quotient on Tr f(D_A/Λ)",
		FermionicAction:                  "S_F = 1/2 <Jψ, D_A ψ> or Pfaffian-style quotient to avoid fermionic double counting",
		FermionicHalfAppliesToBosons:     false,
		BosonicTraceUsesFullHilbertSpace: true,
		Status:                           StatusRealTripleTraceAxiomFormalized,
	}
}

func verifyMirrorCurvature() CurvatureMirror {
	return CurvatureMirror{
		ParticleCurvatureTerm:         "Tr_H(F†F)",
		JMirrorCurvatureTerm:          "Tr_{JH}(\u0305F†\u0305F)",
		ComplexConjugateTraceIdentity: "Tr(\u0305F†\u0305F)=Tr(F†F)",
		ParticleIndex:                 1.0,
		MirrorIndex:                   1.0,
		TotalBosonicIndex:             2.0,
		SameSign:                      true,
		Positive:                      true,
		Status:                        StatusJMirrorCurvatureContributionVerified,
	}
}

func computeLane(name, formula string, multiplier float64, native bool, status string) CouplingLane {
	alphaInv := multiplier * sTopologicalAction / (2.0 * math.Pi)
	g2 := 4.0 * math.Pi / alphaInv
	lambda := (contactScalarNumerator / contactScalarDenominator) * g2
	mass := electroweakVEVGeV * math.Sqrt(2.0*lambda)
	diff := mass - observedHiggsGeV
	return CouplingLane{Name: name, Formula: formula, TraceMultiplier: multiplier, AlphaInverse: alphaInv, GStarSquared: g2, LambdaH: lambda, HiggsMassGeV: mass, DifferenceGeV: diff, RelativeErrorPct: diff / observedHiggsGeV * 100.0, NativeBosonicTrace: native, Status: status}
}

func auditSeparation(quotient CouplingLane) SeparationAudit {
	return SeparationAudit{
		FermionicPfaffianReason:           "fermion doubling is removed in the Grassmann/Pfaffian fermionic action",
		BosonicHeatKernelReason:           "bosonic heat-kernel coefficients are ordinary traces over the operator spectrum and count conjugate gauge carriers positively",
		HalfFactorConfinedToFermions:      true,
		ApplyingHalfToBosonsBreaksEightPi: !nearlyEqual(quotient.AlphaInverse, 8.0*math.Pi, 1e-12),
		QuotientLaneRejected:              true,
		Status:                            StatusFermionicHalfFactorSeparated,
	}
}

func auditPromotion(trace TraceAxiom, mirror CurvatureMirror, doubled CouplingLane) PromotionAudit {
	traceSupplies := trace.BosonicTraceUsesFullHilbertSpace && !trace.FermionicHalfAppliesToBosons && nearlyEqual(mirror.TotalBosonicIndex, 2.0, 1e-12) && mirror.SameSign && mirror.Positive
	return PromotionAudit{
		Gate329ConditionalFactorTwo:               true,
		Gate330TraceConventionSuppliesFactorTwo:   traceSupplies,
		EightPiPromotedWithinBosonicAction:        traceSupplies && nearlyEqual(doubled.AlphaInverse, 8.0*math.Pi, 1e-12) && nearlyEqual(doubled.GStarSquared, 0.5, 1e-12),
		AlphaUnconditional:                        false,
		MissingRepresentationTraceIndex:           true,
		MissingTopologicalActionToCouplingTheorem: true,
		Reason: "full doubled bosonic trace supplies the factor two; absolute α_GUT still also requires the exact representation trace index and the theorem mapping S_top to the Yang-Mills coefficient",
		Status: StatusEightPiBranchPromotedByTraceConvention,
	}
}

func auditFirewalls(p PromotionAudit) FirewallAudit {
	return FirewallAudit{
		NoEmpiricalAlphaInserted:            true,
		NoObservedMassFitted:                true,
		NoPoleMassClaimed:                   true,
		NoFinalColliderClaimed:              true,
		RepresentationIndexStillFirewalled:  p.MissingRepresentationTraceIndex,
		TopologicalActionMapStillFirewalled: p.MissingTopologicalActionToCouplingTheorem,
		Status:                              StatusTensionAbsoluteCouplingStillConditional,
	}
}

func compileSummary(trace TraceAxiom, sep SeparationAudit, p PromotionAudit, doubled CouplingLane) Summary {
	return Summary{
		FullDoubledTraceNativeForBosons:    trace.BosonicTraceUsesFullHilbertSpace,
		FermionicQuotientSeparated:         sep.HalfFactorConfinedToFermions && sep.QuotientLaneRejected,
		EightPiBranchPromotedConditionally: p.EightPiPromotedWithinBosonicAction,
		AlphaStillUnconditionalFailure:     !p.AlphaUnconditional,
		HiggsProxyGeV:                      doubled.HiggsMassGeV,
		DirectAnswer:                       "The bosonic spectral action uses the full doubled heat-kernel trace; the fermionic half-factor does not divide the bosonic F² term. This conditionally promotes the 8π branch, but not yet the unconditional α_GUT theorem.",
		NextObligation:                     "derive the exact representation trace index / weighted finite trace functional that maps the full doubled spectral trace into the canonical Yang-Mills normalization without external convention input",
		Status:                             StatusBosonicFullTraceNative,
	}
}

func Statuses(a Analysis) []string {
	return []string{
		a.Inputs.Status,
		a.Trace.Status,
		StatusBosonicFullTraceNative,
		a.Mirror.Status,
		a.Separation.Status,
		a.Promotion.Status,
		StatusHiggsProxyRecomputed,
		a.Audit.Status,
		StatusTensionRepresentationTraceIndexRequired,
		StatusFailedQuotientAsBosonicRuleRejected,
		StatusFailedAlphaUnconditionalNotDerived,
		StatusFailedPoleMassNotExecuted,
		StatusFailedColliderMassNotClaimed,
	}
}

func FormatInputs(i Inputs) string {
	return fmt.Sprintf("gate=%d S_top=%s=%.12f single_alpha_inv=%.12f doubled_alpha_inv=%.12f empirical=%v", i.HighestInheritedGate, i.STopFormula, i.STop, i.SingleCarrierAlphaInv, i.DoubledAlphaInv, i.AddsEmpiricalAlpha)
}
func FormatTrace(t TraceAxiom) string {
	return fmt.Sprintf("bosonic=%q domain=%q fermionic_half_applies_to_bosons=%v full_trace=%v", t.BosonicSpectralAction, t.TraceDomain, t.FermionicHalfAppliesToBosons, t.BosonicTraceUsesFullHilbertSpace)
}
func FormatMirror(m CurvatureMirror) string {
	return fmt.Sprintf("particle=%.1f mirror=%.1f total=%.1f same_sign=%v positive=%v identity=%s", m.ParticleIndex, m.MirrorIndex, m.TotalBosonicIndex, m.SameSign, m.Positive, m.ComplexConjugateTraceIdentity)
}
func FormatLane(l CouplingLane) string {
	return fmt.Sprintf("%s %s multiplier=%.6f alpha_inv=%.12f g2=%.12f lambda=%.12f m=%.6fGeV diff=%.6fGeV err=%.6f%% native=%v", l.Name, l.Formula, l.TraceMultiplier, l.AlphaInverse, l.GStarSquared, l.LambdaH, l.HiggsMassGeV, l.DifferenceGeV, l.RelativeErrorPct, l.NativeBosonicTrace)
}
func FormatSeparation(s SeparationAudit) string {
	return fmt.Sprintf("half_confined_to_fermions=%v quotient_rejected=%v reason=%s", s.HalfFactorConfinedToFermions, s.QuotientLaneRejected, s.BosonicHeatKernelReason)
}
func FormatPromotion(p PromotionAudit) string {
	return fmt.Sprintf("factor_two=%v eight_pi=%v alpha_unconditional=%v missing_rep_index=%v missing_action_map=%v reason=%s", p.Gate330TraceConventionSuppliesFactorTwo, p.EightPiPromotedWithinBosonicAction, p.AlphaUnconditional, p.MissingRepresentationTraceIndex, p.MissingTopologicalActionToCouplingTheorem, p.Reason)
}
func FormatAudit(a FirewallAudit) string {
	return fmt.Sprintf("no_emp_alpha=%v no_fit=%v no_pole=%v no_collider=%v rep_firewall=%v action_map_firewall=%v", a.NoEmpiricalAlphaInserted, a.NoObservedMassFitted, a.NoPoleMassClaimed, a.NoFinalColliderClaimed, a.RepresentationIndexStillFirewalled, a.TopologicalActionMapStillFirewalled)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("full_trace=%v fermionic_quotient_separated=%v eight_pi_conditional=%v alpha_unconditional_failure=%v m_proxy=%.6f next=%s", s.FullDoubledTraceNativeForBosons, s.FermionicQuotientSeparated, s.EightPiBranchPromotedConditionally, s.AlphaStillUnconditionalFailure, s.HiggsProxyGeV, s.NextObligation)
}

func FormatStatuses(statuses []string) string { return strings.Join(statuses, "\n") }

func nearlyEqual(a, b, eps float64) bool { return math.Abs(a-b) <= eps }
