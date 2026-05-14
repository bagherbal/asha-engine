// Package hierarchyscalingaudit implements Gate 339:
// Gauge Hierarchy Scaling Audit / Planck Factor Sieve.
//
// Gate 338 closed the Higgs pole-matching precision ledger and left the scale
// hierarchy as a separate structural obligation. Gate 339 audits whether the
// finite Cℓ(1,7) topological invariants already contain a non-arbitrary
// suppression factor capable of explaining v/M_Pl. It computes the exact target
// ratios for unreduced and reduced Planck masses, evaluates native exponential,
// combinatorial, and trace-capacity candidates, and rejects arbitrary exponent
// fitting unless a prior theorem supplies the scale-control mechanism.
package hierarchyscalingaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE339-GAUGE-HIERARCHY-SCALING-AUDIT-PLANCK-FACTOR-SIEVE"

	StatusGate338Inherited             = "CONDITIONAL_SUPPORT_GATE338_POLE_MATCHING_AUDIT_INHERITED"
	StatusHierarchyRatioFormalized     = "CONDITIONAL_SUPPORT_HIERARCHY_RATIO_FORMALIZED"
	StatusTopologicalCandidatesAudited = "CONDITIONAL_SUPPORT_TOPOLOGICAL_SCALING_CANDIDATES_AUDITED"
	StatusScaleSynthesisSieveExecuted  = "CONDITIONAL_SUPPORT_SCALE_FACTOR_SYNTHESIS_SIEVE_EXECUTED"
	StatusNearMissesCataloged          = "CONDITIONAL_SUPPORT_HIERARCHY_NEAR_MISSES_CATALOGED"
	StatusFirewallsPreserved           = "CONDITIONAL_SUPPORT_HIERARCHY_FIREWALLS_PRESERVED"

	StatusTensionNoNativeMechanism       = "CONDITIONAL_TENSION_NO_CANONICAL_NATIVE_HIERARCHY_MECHANISM_FOUND"
	StatusTensionInstantonTooLarge       = "CONDITIONAL_TENSION_BGAP_INSTANTON_SUPPRESSION_TOO_LARGE"
	StatusTensionSTopTooSmall            = "CONDITIONAL_TENSION_STOP_EXPONENTIAL_SUPPRESSION_TOO_SMALL"
	StatusTensionPowerNearMissUnpromoted = "CONDITIONAL_TENSION_POWER_OF_TWO_NEAR_MISS_UNPROMOTED"

	StatusFailedHierarchyDerived         = "FAILED_ROUTE_HIERARCHY_SCALING_FACTOR_NOT_DERIVED"
	StatusFailedF2MomentUnlocked         = "FAILED_ROUTE_F2_CUTOFF_MOMENT_STILL_UNLOCKED"
	StatusFailedPlanckScaleNotNative     = "FAILED_ROUTE_PLANCK_SCALE_NORMALIZATION_NOT_DERIVED"
	StatusFailedArbitraryPowerFitting    = "FAILED_ROUTE_ARBITRARY_EXPONENT_FITTING_REJECTED"
	StatusFailedElectroweakVEVNotClaimed = "FAILED_ROUTE_ELECTROWEAK_VEV_NOT_DERIVED"
)

const (
	inheritedHighestGate = 338

	electroweakVEVGeV = 246.22
	// Conventional unreduced Planck mass in GeV. Used as a target ledger input,
	// not as a derived finite-algebra theorem.
	unreducedPlanckGeV = 1.220890e19

	bGap               = 0.102464921191
	contactResonance   = 4 / math.Pi
	sTop               = 8 * math.Pi * math.Pi
	f0Contact          = 7.0
	alphaGUTInvEightPi = 8 * math.Pi
	traceCapacity25    = 25.0
)

type Inputs struct {
	HighestInheritedGate int
	ElectroweakVEVGeV    float64
	UnreducedPlanckGeV   float64
	ReducedPlanckGeV     float64
	BGap                 float64
	SInst                float64
	STop                 float64
	Status               string
}

type HierarchyTargets struct {
	RhoUnreduced   float64
	RhoReduced     float64
	Log10Unreduced float64
	Log10Reduced   float64
	Status         string
}

type Candidate struct {
	Name                string
	Expression          string
	Value               float64
	RatioToUnreduced    float64
	RatioToReduced      float64
	Log10Value          float64
	Log10ErrorUnreduced float64
	Log10ErrorReduced   float64
	Native              bool
	Promotable          bool
	Verdict             string
}

type CandidateLedger struct {
	Candidates    []Candidate
	BestUnreduced Candidate
	BestReduced   Candidate
	Status        string
}

type SynthesisLane struct {
	Name                     string
	Expression               string
	Value                    float64
	RequiresFreeExponent     bool
	RequiresUnprovedScaleLaw bool
	Interpretation           string
}

type SynthesisSieve struct {
	Lanes              []SynthesisLane
	NativeDerived      bool
	ClosestNonPromoted string
	Status             string
}

type Firewalls struct {
	NoHierarchyScalingFactor         bool
	F2MomentUnlocked                 bool
	PlanckScaleNotNative             bool
	ArbitraryExponentFittingRejected bool
	ElectroweakVEVNotDerived         bool
	Status                           string
}

type Summary struct {
	RhoUnreduced        float64
	RhoReduced          float64
	BestNativeCandidate string
	BestNativeValue     float64
	DirectAnswer        string
	NextGate            string
	Status              string
}

type Analysis struct {
	Inputs     Inputs
	Targets    HierarchyTargets
	Candidates CandidateLedger
	Synthesis  SynthesisSieve
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
	inputs := compileInputs()
	targets := formalizeTargets(inputs)
	candidates := auditCandidates(inputs, targets)
	synthesis := executeSynthesisSieve(inputs, targets, candidates)
	firewalls := preserveFirewalls()
	summary := compileSummary(targets, candidates, synthesis)
	truth := "Gate 339 computes the hierarchy targets v/M_P and v/Mbar_P and audits native ASHA scaling candidates. B-gap instanton suppression, doubled-space powers, trace capacity, and the 8π coupling factor do not canonically generate the 10^-17 hierarchy. A rank-56 power-of-two near miss is cataloged but rejected as a derivation because no theorem states that this rank exponent controls the electroweak-to-Planck ratio. Therefore the gauge hierarchy remains a Phase-III f2/Planck-normalization obligation, not a solved finite-algebra result."
	return Analysis{Inputs: inputs, Targets: targets, Candidates: candidates, Synthesis: synthesis, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func compileInputs() Inputs {
	reduced := unreducedPlanckGeV / math.Sqrt(8*math.Pi)
	sInst := contactResonance / bGap
	return Inputs{
		HighestInheritedGate: inheritedHighestGate,
		ElectroweakVEVGeV:    electroweakVEVGeV,
		UnreducedPlanckGeV:   unreducedPlanckGeV,
		ReducedPlanckGeV:     reduced,
		BGap:                 bGap,
		SInst:                sInst,
		STop:                 sTop,
		Status:               StatusGate338Inherited,
	}
}

func formalizeTargets(inputs Inputs) HierarchyTargets {
	rhoU := inputs.ElectroweakVEVGeV / inputs.UnreducedPlanckGeV
	rhoR := inputs.ElectroweakVEVGeV / inputs.ReducedPlanckGeV
	return HierarchyTargets{
		RhoUnreduced:   rhoU,
		RhoReduced:     rhoR,
		Log10Unreduced: math.Log10(rhoU),
		Log10Reduced:   math.Log10(rhoR),
		Status:         StatusHierarchyRatioFormalized,
	}
}

func auditCandidates(inputs Inputs, targets HierarchyTargets) CandidateLedger {
	raw := []struct {
		name, expr, verdict string
		value               float64
		native, promotable  bool
	}{
		{name: "B-gap instanton exponential", expr: "exp[-(4/π)/B_gap]", value: math.Exp(-inputs.SInst), native: true, promotable: false, verdict: "too large by ~11 orders for unreduced Planck target"},
		{name: "topological action exponential", expr: "exp[-S_top] = exp[-8π²]", value: math.Exp(-inputs.STop), native: true, promotable: false, verdict: "too small by ~18 orders for unreduced Planck target"},
		{name: "topological action square-root", expr: "exp[-S_top/2] = exp[-4π²]", value: math.Exp(-inputs.STop / 2), native: true, promotable: false, verdict: "nearer but still not target and square-root rule is not derived"},
		{name: "doubled 16-bit state inverse", expr: "2^-16", value: math.Pow(2, -16), native: true, promotable: false, verdict: "far too large; state count alone cannot generate hierarchy"},
		{name: "three-generation doubled Hilbert inverse", expr: "1/96", value: 1.0 / 96.0, native: true, promotable: false, verdict: "far too large; Hilbert count is not hierarchy suppression"},
		{name: "contact cutoff inverse", expr: "1/f0 = 1/7", value: 1.0 / f0Contact, native: true, promotable: false, verdict: "far too large; cutoff moment is not scale hierarchy"},
		{name: "eight-pi coupling inverse", expr: "1/(8π)", value: 1.0 / alphaGUTInvEightPi, native: true, promotable: false, verdict: "far too large; explains coupling scale, not Planck hierarchy"},
		{name: "trace-capacity inverse", expr: "1/25", value: 1.0 / traceCapacity25, native: false, promotable: false, verdict: "far too large and C_trace=25 remains unproved weighted functional"},
		{name: "rank-56 Boolean near miss", expr: "2^-56", value: math.Pow(2, -56), native: true, promotable: false, verdict: "numerically close to unreduced target but exponent-to-scale law is not derived"},
		{name: "rank-70 exterior near miss", expr: "2^-70", value: math.Pow(2, -70), native: true, promotable: false, verdict: "too small; exterior dimension alone is not a scale law"},
	}
	candidates := make([]Candidate, 0, len(raw))
	for _, r := range raw {
		candidates = append(candidates, makeCandidate(r.name, r.expr, r.value, targets, r.native, r.promotable, r.verdict))
	}
	sort.Slice(candidates, func(i, j int) bool {
		return math.Abs(candidates[i].Log10ErrorUnreduced) < math.Abs(candidates[j].Log10ErrorUnreduced)
	})
	return CandidateLedger{Candidates: candidates, BestUnreduced: candidates[0], BestReduced: bestByReduced(candidates), Status: StatusTopologicalCandidatesAudited}
}

func makeCandidate(name, expr string, value float64, targets HierarchyTargets, native, promotable bool, verdict string) Candidate {
	return Candidate{
		Name:                name,
		Expression:          expr,
		Value:               value,
		RatioToUnreduced:    value / targets.RhoUnreduced,
		RatioToReduced:      value / targets.RhoReduced,
		Log10Value:          math.Log10(value),
		Log10ErrorUnreduced: math.Log10(value / targets.RhoUnreduced),
		Log10ErrorReduced:   math.Log10(value / targets.RhoReduced),
		Native:              native,
		Promotable:          promotable,
		Verdict:             verdict,
	}
}

func bestByReduced(candidates []Candidate) Candidate {
	best := candidates[0]
	bestErr := math.Abs(best.Log10ErrorReduced)
	for _, c := range candidates[1:] {
		if err := math.Abs(c.Log10ErrorReduced); err < bestErr {
			best = c
			bestErr = err
		}
	}
	return best
}

func executeSynthesisSieve(inputs Inputs, targets HierarchyTargets, candidates CandidateLedger) SynthesisSieve {
	lanes := []SynthesisLane{
		{
			Name:                     "single native invariant lane",
			Expression:               "candidate ∈ {exp(-S_inst), exp(-S_top), 2^-16, 1/(8π), ...}",
			Value:                    candidates.BestUnreduced.Value,
			RequiresFreeExponent:     false,
			RequiresUnprovedScaleLaw: true,
			Interpretation:           "No single native invariant is both canonical and close enough to derive the hierarchy.",
		},
		{
			Name:                     "rank-56 power lane",
			Expression:               "2^-56",
			Value:                    math.Pow(2, -56),
			RequiresFreeExponent:     false,
			RequiresUnprovedScaleLaw: true,
			Interpretation:           "Numerically near v/M_P but requires a theorem that rank-56 Boolean capacity exponentiates into the mass hierarchy.",
		},
		{
			Name:                     "fit exponent lane",
			Expression:               "2^-n with n = -log2(v/M_P)",
			Value:                    targets.RhoUnreduced,
			RequiresFreeExponent:     true,
			RequiresUnprovedScaleLaw: true,
			Interpretation:           "Always matches by construction; rejected as arbitrary exponent fitting.",
		},
		{
			Name:                     "B-gap instanton times combinatorics lane",
			Expression:               "exp(-S_inst) × 2^-n",
			Value:                    targets.RhoUnreduced,
			RequiresFreeExponent:     true,
			RequiresUnprovedScaleLaw: true,
			Interpretation:           "Can be tuned but no native n-selection theorem is available.",
		},
	}
	closest := fmt.Sprintf("%s = %.15e (ratio to unreduced target %.6g)", candidates.BestUnreduced.Name, candidates.BestUnreduced.Value, candidates.BestUnreduced.RatioToUnreduced)
	return SynthesisSieve{Lanes: lanes, NativeDerived: false, ClosestNonPromoted: closest, Status: StatusScaleSynthesisSieveExecuted}
}

func preserveFirewalls() Firewalls {
	return Firewalls{
		NoHierarchyScalingFactor:         true,
		F2MomentUnlocked:                 true,
		PlanckScaleNotNative:             true,
		ArbitraryExponentFittingRejected: true,
		ElectroweakVEVNotDerived:         true,
		Status:                           StatusFirewallsPreserved,
	}
}

func compileSummary(targets HierarchyTargets, candidates CandidateLedger, synthesis SynthesisSieve) Summary {
	return Summary{
		RhoUnreduced:        targets.RhoUnreduced,
		RhoReduced:          targets.RhoReduced,
		BestNativeCandidate: candidates.BestUnreduced.Name,
		BestNativeValue:     candidates.BestUnreduced.Value,
		DirectAnswer:        "No native, non-arbitrary hierarchy scaling factor is derived in Gate 339. The finite geometry supplies useful near-miss diagnostics, especially 2^-56, but does not yet prove that any candidate controls v/M_Pl.",
		NextGate:            "Derive or reject a native f2/Planck normalization theorem, preferably a weighted gravitational Seeley-de Witt a2 coefficient audit tying Newton's constant to the finite trace functional.",
		Status:              StatusFailedHierarchyDerived,
	}
}

func Statuses(a Analysis) []string {
	statuses := []string{
		a.Inputs.Status,
		a.Targets.Status,
		a.Candidates.Status,
		a.Synthesis.Status,
		StatusNearMissesCataloged,
		a.Firewalls.Status,
		StatusTensionNoNativeMechanism,
		StatusTensionInstantonTooLarge,
		StatusTensionSTopTooSmall,
		StatusTensionPowerNearMissUnpromoted,
		StatusFailedHierarchyDerived,
		StatusFailedF2MomentUnlocked,
		StatusFailedPlanckScaleNotNative,
		StatusFailedArbitraryPowerFitting,
		StatusFailedElectroweakVEVNotClaimed,
	}
	return statuses
}

func FormatInputs(i Inputs) string {
	return fmt.Sprintf("highest_gate=%d; v=%.12f GeV; M_P=%.12e GeV; Mbar_P=%.12e GeV; B_gap=%.12f; S_inst=%.12f; S_top=%.12f; status=%s", i.HighestInheritedGate, i.ElectroweakVEVGeV, i.UnreducedPlanckGeV, i.ReducedPlanckGeV, i.BGap, i.SInst, i.STop, i.Status)
}

func FormatTargets(t HierarchyTargets) string {
	return fmt.Sprintf("rho_unreduced=v/M_P=%.15e (log10 %.9f); rho_reduced=v/Mbar_P=%.15e (log10 %.9f); status=%s", t.RhoUnreduced, t.Log10Unreduced, t.RhoReduced, t.Log10Reduced, t.Status)
}

func FormatCandidate(c Candidate) string {
	return fmt.Sprintf("%s: %s = %.15e; ratio_unred=%.9e; ratio_red=%.9e; log10err_unred=%+.6f; native=%t; promotable=%t; verdict=%s", c.Name, c.Expression, c.Value, c.RatioToUnreduced, c.RatioToReduced, c.Log10ErrorUnreduced, c.Native, c.Promotable, c.Verdict)
}

func FormatCandidates(l CandidateLedger) string {
	parts := []string{fmt.Sprintf("status=%s", l.Status), fmt.Sprintf("best_unreduced={%s}", FormatCandidate(l.BestUnreduced)), fmt.Sprintf("best_reduced={%s}", FormatCandidate(l.BestReduced))}
	for _, c := range l.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return strings.Join(parts, "\n")
}

func FormatSynthesisLane(l SynthesisLane) string {
	return fmt.Sprintf("%s: %s -> %.15e; free_exponent=%t; unproved_scale_law=%t; %s", l.Name, l.Expression, l.Value, l.RequiresFreeExponent, l.RequiresUnprovedScaleLaw, l.Interpretation)
}

func FormatSynthesis(s SynthesisSieve) string {
	parts := []string{fmt.Sprintf("native_derived=%t; closest_non_promoted=%s; status=%s", s.NativeDerived, s.ClosestNonPromoted, s.Status)}
	for _, l := range s.Lanes {
		parts = append(parts, FormatSynthesisLane(l))
	}
	return strings.Join(parts, "\n")
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("no_hierarchy=%t; f2_unlocked=%t; planck_not_native=%t; reject_fit=%t; v_not_derived=%t; status=%s", f.NoHierarchyScalingFactor, f.F2MomentUnlocked, f.PlanckScaleNotNative, f.ArbitraryExponentFittingRejected, f.ElectroweakVEVNotDerived, f.Status)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("rho_unred=%.15e; rho_red=%.15e; best=%s %.15e; direct=%s; next=%s; status=%s", s.RhoUnreduced, s.RhoReduced, s.BestNativeCandidate, s.BestNativeValue, s.DirectAnswer, s.NextGate, s.Status)
}

func FormatStatuses(statuses []string) string { return strings.Join(statuses, "\n") }
