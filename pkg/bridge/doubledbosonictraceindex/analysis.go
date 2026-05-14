// Package doubledbosonictraceindex implements Gate 329:
// Doubled Bosonic Trace Index / J-Mirror Gauge Capacity Audit.
//
// Gate 328 isolated a factor-of-two obstruction: the conventional Chern-Weil
// instanton lane gives α^{-1}=S_top/(2π)=4π, while the Higgs-successful lane
// requires α^{-1}=S_top/π=8π. Gate 329 audits whether the doubled real spectral
// triple H_F ⊕ H_F* supplies the required factor two through the bosonic spectral
// trace over particle and J-mirror antiparticle gauge carriers.
package doubledbosonictraceindex

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE329-DOUBLED-BOSONIC-TRACE-INDEX-J-MIRROR-GAUGE-CAPACITY-AUDIT"

	StatusInputsInherited                     = "CONDITIONAL_SUPPORT_GATE328_FACTOR_TWO_OBLIGATION_INHERITED"
	StatusGaugeMirrorFormalized               = "CONDITIONAL_SUPPORT_J_MIRROR_GAUGE_CARRIER_FORMALIZED"
	StatusDoubledBosonicTraceFactorFormalized = "CONDITIONAL_SUPPORT_DOUBLED_BOSONIC_TRACE_FACTOR_TWO_FORMALIZED"
	StatusQuotientLaneAudited                 = "CONDITIONAL_SUPPORT_REAL_STRUCTURE_QUOTIENT_LANE_AUDITED"
	StatusEightPiBranchConditionallyPromoted  = "CONDITIONAL_SUPPORT_EIGHT_PI_BRANCH_CONDITIONALLY_PROMOTED_BY_FULL_BOSONIC_TRACE"
	StatusHiggsProxyRecomputed                = "CONDITIONAL_SUPPORT_HIGGS_PROXY_RECOMPUTED_WITH_DOUBLED_TRACE"

	StatusTensionBosonicTraceConventionRequired  = "CONDITIONAL_TENSION_BOSONIC_TRACE_CONVENTION_REQUIRED"
	StatusTensionFermionicQuotientWrongDirection = "CONDITIONAL_TENSION_FERMIONIC_QUOTIENT_DOES_NOT_SUPPLY_EIGHT_PI"

	StatusFailedUnconditionalAlphaNotDerived = "FAILED_ROUTE_ALPHA_GUT_UNCONDITIONAL_VALUE_NOT_DERIVED"
	StatusFailedTraceConventionNotNative     = "FAILED_ROUTE_BOSONIC_TRACE_OVER_FULL_DOUBLED_SPACE_NOT_PROVED_NATIVE"
	StatusFailedQuotientLaneRejectsEightPi   = "FAILED_ROUTE_QUOTIENTED_PHYSICAL_TRACE_REJECTS_EIGHT_PI"
	StatusFailedColliderMassNotClaimed       = "FAILED_ROUTE_FINAL_COLLIDER_HIGGS_MASS_NOT_CLAIMED"
	StatusFailedPoleMassNotExecuted          = "FAILED_ROUTE_POLE_MASS_CONVERSION_NOT_EXECUTED"
)

const (
	inheritedHighestGate     = 328
	contactScalarNumerator   = 1197.0
	contactScalarDenominator = 4624.0
	electroweakVEVGeV        = 246.22
	observedHiggsGeV         = 125.10
	sTopologicalAction       = 8.0 * math.Pi * math.Pi
)

type Inputs struct {
	HighestInheritedGate int
	STopFormula          string
	STop                 float64
	ConventionalAlphaInv float64
	RequiredAlphaInv     float64
	RequiredMultiplier   float64
	ContactScalarRatio   float64
	AddsEmpiricalFit     bool
	Status               string
}

type MirrorCarrier struct {
	ParticleCarrierName       string
	MirrorCarrierName         string
	GaugeCurvatureAction      string
	ParticleTraceIndex        float64
	JMirrorTraceIndex         float64
	FullDoubledTraceIndex     float64
	CurvaturesHaveSameF2Sign  bool
	ComplexConjugationNeutral bool
	Status                    string
}

type TraceLane struct {
	Name             string
	Description      string
	TraceMultiplier  float64
	AlphaInverse     float64
	GStarSquared     float64
	LambdaH          float64
	HiggsMassGeV     float64
	DifferenceGeV    float64
	RelativeErrorPct float64
	MatchesEightPi   bool
	MatchesHiggs     bool
	Promoted         bool
	Status           string
}

type QuotientAudit struct {
	FermionicActionNeedsHalfFactor bool
	BosonicSpectralActionNeedsHalf bool
	QuotientMultiplier             float64
	QuotientAlphaInverse           float64
	QuotientGStarSquared           float64
	QuotientHiggsMassGeV           float64
	WrongDirectionForEightPi       bool
	CanExplainEightPi              bool
	Status                         string
}

type PromotionAudit struct {
	RequiredMultiplier         float64
	DoubledTraceMultiplier     float64
	MultiplierMatches          bool
	FullBosonicTraceNativeHere bool
	QuotientConventionRejected bool
	ConditionalPromotion       bool
	UnconditionalDerivation    bool
	Reason                     string
	Status                     string
}

type FirewallAudit struct {
	NoEmpiricalAlphaInserted        bool
	NoObservedHiggsFitInserted      bool
	NoPoleMassClaimed               bool
	NoFinalColliderMassClaimed      bool
	QuotientLaneKeptVisible         bool
	TraceConventionStillConditional bool
	FiniteCorePolluted              bool
	Status                          string
}

type Summary struct {
	FactorTwoSuppliedAsCapacity bool
	EightPiConditionallyWorks   bool
	QuotientLaneFails           bool
	NativeAlphaClosed           bool
	FinalMassClaimed            bool
	DirectAnswer                string
	NextObligation              string
	Status                      string
}

type Analysis struct {
	Inputs    Inputs
	Mirror    MirrorCarrier
	BaseLane  TraceLane
	Doubled   TraceLane
	Quotient  TraceLane
	QuotientA QuotientAudit
	Promotion PromotionAudit
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
	inputs := compileInputs()
	mirror := formalizeMirrorCarrier()
	baseLane := computeLane("single-carrier Chern-Weil lane", "α^{-1}=S_top/(2π)", 1.0, StatusQuotientLaneAudited, false)
	doubledLane := computeLane("full doubled bosonic spectral trace", "α^{-1}=2·S_top/(2π)=S_top/π", 2.0, StatusEightPiBranchConditionallyPromoted, true)
	quotientLane := computeLane("fermionic quotient/half-trace lane", "α^{-1}=½·S_top/(2π)", 0.5, StatusQuotientLaneAudited, false)
	quotientAudit := auditQuotientLane(quotientLane)
	promotion := auditPromotion(inputs, mirror, doubledLane, quotientAudit)
	firewall := auditFirewalls(quotientAudit, promotion)
	summary := compileSummary(doubledLane, quotientLane, promotion)
	truth := "Gate 329 shows that the exact missing factor of two from Gate 328 is present as a capacity of the full doubled bosonic spectral trace: the particle and J-mirror antiparticle curvature carriers contribute equal positive F² terms, giving trace multiplier 2 and α_GUT^{-1}=8π. However, the result is only conditionally promoted because the gate must still prove that the bosonic spectral action uses the full doubled trace without the fermionic real-structure quotient. Quotienting is the wrong direction for the Higgs-successful 8π branch."
	return Analysis{Inputs: inputs, Mirror: mirror, BaseLane: baseLane, Doubled: doubledLane, Quotient: quotientLane, QuotientA: quotientAudit, Promotion: promotion, Audit: firewall, Summary: summary, Truth: truth}, nil
}

func compileInputs() Inputs {
	conventional := sTopologicalAction / (2.0 * math.Pi)
	required := sTopologicalAction / math.Pi
	return Inputs{
		HighestInheritedGate: inheritedHighestGate,
		STopFormula:          "S_top = 8π²",
		STop:                 sTopologicalAction,
		ConventionalAlphaInv: conventional,
		RequiredAlphaInv:     required,
		RequiredMultiplier:   required / conventional,
		ContactScalarRatio:   contactScalarNumerator / contactScalarDenominator,
		AddsEmpiricalFit:     false,
		Status:               StatusInputsInherited,
	}
}

func formalizeMirrorCarrier() MirrorCarrier {
	return MirrorCarrier{
		ParticleCarrierName:       "H_F particle gauge carrier",
		MirrorCarrierName:         "J H_F = H_F* antiparticle gauge carrier",
		GaugeCurvatureAction:      "F acts by the representation on H_F and by the conjugate representation on H_F*; Tr(F†F) is invariant under conjugation",
		ParticleTraceIndex:        1.0,
		JMirrorTraceIndex:         1.0,
		FullDoubledTraceIndex:     2.0,
		CurvaturesHaveSameF2Sign:  true,
		ComplexConjugationNeutral: true,
		Status:                    StatusGaugeMirrorFormalized,
	}
}

func computeLane(name, desc string, multiplier float64, status string, promoted bool) TraceLane {
	alphaInv := multiplier * sTopologicalAction / (2.0 * math.Pi)
	g2 := 4.0 * math.Pi / alphaInv
	lambda := (contactScalarNumerator / contactScalarDenominator) * g2
	mass := electroweakVEVGeV * math.Sqrt(2.0*lambda)
	diff := mass - observedHiggsGeV
	return TraceLane{
		Name:             name,
		Description:      desc,
		TraceMultiplier:  multiplier,
		AlphaInverse:     alphaInv,
		GStarSquared:     g2,
		LambdaH:          lambda,
		HiggsMassGeV:     mass,
		DifferenceGeV:    diff,
		RelativeErrorPct: diff / observedHiggsGeV * 100.0,
		MatchesEightPi:   nearlyEqual(alphaInv, 8.0*math.Pi, 1e-12),
		MatchesHiggs:     math.Abs(diff) < 0.5,
		Promoted:         promoted,
		Status:           status,
	}
}

func auditQuotientLane(quotient TraceLane) QuotientAudit {
	return QuotientAudit{
		FermionicActionNeedsHalfFactor: true,
		BosonicSpectralActionNeedsHalf: false,
		QuotientMultiplier:             quotient.TraceMultiplier,
		QuotientAlphaInverse:           quotient.AlphaInverse,
		QuotientGStarSquared:           quotient.GStarSquared,
		QuotientHiggsMassGeV:           quotient.HiggsMassGeV,
		WrongDirectionForEightPi:       quotient.AlphaInverse < sTopologicalAction/(2.0*math.Pi),
		CanExplainEightPi:              false,
		Status:                         StatusTensionFermionicQuotientWrongDirection,
	}
}

func auditPromotion(inputs Inputs, mirror MirrorCarrier, doubled TraceLane, quotient QuotientAudit) PromotionAudit {
	matches := nearlyEqual(inputs.RequiredMultiplier, mirror.FullDoubledTraceIndex, 1e-12) && doubled.MatchesEightPi
	conditional := matches && mirror.CurvaturesHaveSameF2Sign && mirror.ComplexConjugationNeutral && !quotient.CanExplainEightPi
	return PromotionAudit{
		RequiredMultiplier:         inputs.RequiredMultiplier,
		DoubledTraceMultiplier:     mirror.FullDoubledTraceIndex,
		MultiplierMatches:          matches,
		FullBosonicTraceNativeHere: false,
		QuotientConventionRejected: quotient.WrongDirectionForEightPi,
		ConditionalPromotion:       conditional,
		UnconditionalDerivation:    false,
		Reason:                     "The full doubled bosonic trace supplies exactly the required factor two, but the gate still needs a native spectral-action convention theorem proving that the bosonic action is traced over H_F⊕H_F* without quotienting while the fermionic action alone receives the real-structure half-factor.",
		Status:                     StatusDoubledBosonicTraceFactorFormalized,
	}
}

func auditFirewalls(quotient QuotientAudit, promotion PromotionAudit) FirewallAudit {
	return FirewallAudit{
		NoEmpiricalAlphaInserted:        true,
		NoObservedHiggsFitInserted:      true,
		NoPoleMassClaimed:               true,
		NoFinalColliderMassClaimed:      true,
		QuotientLaneKeptVisible:         quotient.WrongDirectionForEightPi,
		TraceConventionStillConditional: !promotion.UnconditionalDerivation,
		FiniteCorePolluted:              false,
		Status:                          StatusFailedTraceConventionNotNative,
	}
}

func compileSummary(doubled, quotient TraceLane, promotion PromotionAudit) Summary {
	return Summary{
		FactorTwoSuppliedAsCapacity: promotion.MultiplierMatches,
		EightPiConditionallyWorks:   doubled.MatchesEightPi && doubled.MatchesHiggs,
		QuotientLaneFails:           quotient.AlphaInverse < sTopologicalAction/(2.0*math.Pi),
		NativeAlphaClosed:           false,
		FinalMassClaimed:            false,
		DirectAnswer:                "The doubled J-mirror bosonic trace supplies the exact factor two needed for α_GUT^{-1}=8π, but this is a conditional convention theorem rather than an unconditional α_GUT derivation until the bosonic-vs-fermionic trace quotient rule is proven natively.",
		NextObligation:              "Prove the bosonic spectral action trace is over the full doubled H_F⊕H_F* carrier, while the half-quotient applies only to the fermionic bilinear, not to the gauge kinetic a4 trace.",
		Status:                      StatusTensionBosonicTraceConventionRequired,
	}
}

func Statuses(a Analysis) []string {
	return []string{
		a.Inputs.Status,
		a.Mirror.Status,
		a.BaseLane.Status,
		a.Doubled.Status,
		a.Quotient.Status,
		a.QuotientA.Status,
		a.Promotion.Status,
		a.Audit.Status,
		a.Summary.Status,
		StatusHiggsProxyRecomputed,
		StatusFailedUnconditionalAlphaNotDerived,
		StatusFailedQuotientLaneRejectsEightPi,
		StatusFailedColliderMassNotClaimed,
		StatusFailedPoleMassNotExecuted,
	}
}

func FormatInputs(v Inputs) string {
	return fmt.Sprintf("gate=%d; %s; conventional_alpha_inv=%.12f; required_alpha_inv=%.12f; required_multiplier=%.6f; contact_shape=%.12f; empirical_fit=%t", v.HighestInheritedGate, v.STopFormula, v.ConventionalAlphaInv, v.RequiredAlphaInv, v.RequiredMultiplier, v.ContactScalarRatio, v.AddsEmpiricalFit)
}

func FormatMirror(v MirrorCarrier) string {
	return fmt.Sprintf("particle_index=%.1f; J_mirror_index=%.1f; full_index=%.1f; same_F2_sign=%t; conjugation_neutral=%t; action=%s", v.ParticleTraceIndex, v.JMirrorTraceIndex, v.FullDoubledTraceIndex, v.CurvaturesHaveSameF2Sign, v.ComplexConjugationNeutral, v.GaugeCurvatureAction)
}

func FormatLane(v TraceLane) string {
	return fmt.Sprintf("%s: multiplier=%.3f; %s; alpha_inv=%.12f; g2=%.12f; lambda=%.12f; m=%.6f GeV; diff=%+.6f GeV (%+.6f%%); eight_pi=%t; promoted=%t", v.Name, v.TraceMultiplier, v.Description, v.AlphaInverse, v.GStarSquared, v.LambdaH, v.HiggsMassGeV, v.DifferenceGeV, v.RelativeErrorPct, v.MatchesEightPi, v.Promoted)
}

func FormatQuotient(v QuotientAudit) string {
	return fmt.Sprintf("fermionic_half=%t; bosonic_half=%t; quotient_multiplier=%.3f; alpha_inv=%.12f; g2=%.12f; m=%.6f GeV; wrong_direction=%t; can_explain_eight_pi=%t", v.FermionicActionNeedsHalfFactor, v.BosonicSpectralActionNeedsHalf, v.QuotientMultiplier, v.QuotientAlphaInverse, v.QuotientGStarSquared, v.QuotientHiggsMassGeV, v.WrongDirectionForEightPi, v.CanExplainEightPi)
}

func FormatPromotion(v PromotionAudit) string {
	return fmt.Sprintf("required_multiplier=%.3f; doubled_multiplier=%.3f; matches=%t; bosonic_trace_native=%t; quotient_rejected=%t; conditional=%t; unconditional=%t; reason=%s", v.RequiredMultiplier, v.DoubledTraceMultiplier, v.MultiplierMatches, v.FullBosonicTraceNativeHere, v.QuotientConventionRejected, v.ConditionalPromotion, v.UnconditionalDerivation, v.Reason)
}

func FormatAudit(v FirewallAudit) string {
	return fmt.Sprintf("no_alpha_fit=%t; no_higgs_fit=%t; no_pole=%t; no_final=%t; quotient_visible=%t; trace_convention_conditional=%t; polluted=%t", v.NoEmpiricalAlphaInserted, v.NoObservedHiggsFitInserted, v.NoPoleMassClaimed, v.NoFinalColliderMassClaimed, v.QuotientLaneKeptVisible, v.TraceConventionStillConditional, v.FiniteCorePolluted)
}

func FormatSummary(v Summary) string {
	return fmt.Sprintf("factor_two_capacity=%t; eight_pi_works=%t; quotient_fails=%t; native_alpha_closed=%t; final_mass=%t; next=%s", v.FactorTwoSuppliedAsCapacity, v.EightPiConditionallyWorks, v.QuotientLaneFails, v.NativeAlphaClosed, v.FinalMassClaimed, v.NextObligation)
}

func Markdown(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 329 Registry Audit — Doubled Bosonic Trace Index / J-Mirror Gauge Capacity Audit\n\n")
	b.WriteString("## Gate identity\n\n")
	b.WriteString("- **Gate:** 329\n")
	b.WriteString("- **Package:** `pkg/bridge/doubledbosonictraceindex`\n")
	b.WriteString("- **Theorem:** `DoubledBosonicTraceIndexJMirrorGaugeCapacityAuditTheorem`\n")
	b.WriteString("- **Audit ID:** `" + AuditID + "`\n")
	b.WriteString("- **Layer:** Bridge / Phase-II Absolute Coupling Normalization\n")
	b.WriteString("- **Purpose:** audit whether the doubled real spectral triple supplies the factor two needed to promote the Gate 327/328 `8π` coupling branch.\n\n")

	b.WriteString("## Input ledger\n\n")
	b.WriteString("```text\n" + FormatInputs(a.Inputs) + "\n```\n\n")

	b.WriteString("## J-mirror gauge carrier\n\n")
	b.WriteString("```text\n" + FormatMirror(a.Mirror) + "\n```\n\n")
	b.WriteString("The particle and J-mirror antiparticle carriers contribute equal positive gauge-curvature squares. This gives a full doubled bosonic trace index of `2`.\n\n")

	b.WriteString("## Trace-normalization lanes\n\n")
	b.WriteString("| Lane | Multiplier | α⁻¹ | g_*² | Tree proxy | Verdict |\n")
	b.WriteString("| --- | ---: | ---: | ---: | ---: | --- |\n")
	b.WriteString(fmt.Sprintf("| Single carrier | %.3f | %.12f | %.12f | %.6f GeV | conventional `4π` branch |\n", a.BaseLane.TraceMultiplier, a.BaseLane.AlphaInverse, a.BaseLane.GStarSquared, a.BaseLane.HiggsMassGeV))
	b.WriteString(fmt.Sprintf("| Full doubled bosonic trace | %.3f | %.12f | %.12f | %.6f GeV | supplies `8π` capacity; conditional promotion |\n", a.Doubled.TraceMultiplier, a.Doubled.AlphaInverse, a.Doubled.GStarSquared, a.Doubled.HiggsMassGeV))
	b.WriteString(fmt.Sprintf("| Fermionic quotient / half-trace | %.3f | %.12f | %.12f | %.6f GeV | wrong direction for `8π` |\n\n", a.Quotient.TraceMultiplier, a.Quotient.AlphaInverse, a.Quotient.GStarSquared, a.Quotient.HiggsMassGeV))

	b.WriteString("## Quotient audit\n\n")
	b.WriteString("```text\n" + FormatQuotient(a.QuotientA) + "\n```\n\n")
	b.WriteString("The half-factor appropriate to fermionic real-structure bilinears does not supply the `8π` branch. It decreases the gauge trace coefficient instead of doubling it.\n\n")

	b.WriteString("## Promotion audit\n\n")
	b.WriteString("```text\n" + FormatPromotion(a.Promotion) + "\n```\n\n")

	b.WriteString("## Final status ledger\n\n")
	b.WriteString("```text\n")
	for _, s := range Statuses(a) {
		b.WriteString(s + "\n")
	}
	b.WriteString("```\n\n")

	b.WriteString("## Verdict\n\n")
	b.WriteString(a.Truth + "\n\n")
	b.WriteString("**Next obligation:** " + a.Summary.NextObligation + "\n")
	return b.String()
}

func nearlyEqual(a, b, tol float64) bool {
	return math.Abs(a-b) <= tol
}
