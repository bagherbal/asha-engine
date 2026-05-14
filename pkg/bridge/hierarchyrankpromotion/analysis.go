// Package hierarchyrankpromotion implements Gate 340:
// Rank-56 / Half-Instanton Hierarchy Promotion Sieve.
//
// Gate 339 found numerical near misses for the electroweak-to-Planck hierarchy,
// especially 2^-56 and exp(-4π²). Gate 340 audits whether those near misses can
// be promoted into a native scale law, or whether they remain numerically
// suggestive but theoremically unlicensed. It computes the exact target ratios,
// effective exponents, required prefactors, and category checks needed to avoid
// arbitrary exponent fitting.
package hierarchyrankpromotion

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE340-RANK-56-HALF-INSTANTON-HIERARCHY-PROMOTION-SIEVE"

	StatusGate339Inherited            = "CONDITIONAL_SUPPORT_GATE339_HIERARCHY_AUDIT_INHERITED"
	StatusEffectiveExponentComputed   = "CONDITIONAL_SUPPORT_EFFECTIVE_HIERARCHY_EXPONENT_COMPUTED"
	StatusRank56PromotionAudited      = "CONDITIONAL_SUPPORT_RANK56_POWER_LAW_PROMOTION_AUDITED"
	StatusHalfInstantonPromotionAudit = "CONDITIONAL_SUPPORT_HALF_INSTANTON_PROMOTION_AUDITED"
	StatusPrefactorSieveExecuted      = "CONDITIONAL_SUPPORT_PREFactor_ALIGNMENT_SIEVE_EXECUTED"
	StatusCategoryFirewallPreserved   = "CONDITIONAL_SUPPORT_HIERARCHY_CATEGORY_FIREWALLS_PRESERVED"

	StatusTensionRank56NearButNotExact     = "CONDITIONAL_TENSION_RANK56_NEAR_BUT_NOT_EXACT"
	StatusTensionHalfActionNearButNotExact = "CONDITIONAL_TENSION_HALF_TOPOLOGICAL_ACTION_NEAR_BUT_NOT_EXACT"
	StatusTensionSqrtTwoPrefactorNearMiss  = "CONDITIONAL_TENSION_SQRT_TWO_PREFactor_NEAR_MISS_UNPROMOTED"
	StatusTensionNoScaleLawSelector        = "CONDITIONAL_TENSION_NO_NATIVE_SCALE_LAW_SELECTOR_FOUND"

	StatusFailedRank56ScaleLawNotDerived    = "FAILED_ROUTE_RANK56_SCALE_LAW_NOT_DERIVED"
	StatusFailedHalfInstantonRuleNotDerived = "FAILED_ROUTE_HALF_INSTANTON_RULE_NOT_DERIVED"
	StatusFailedPrefactorNotDerived         = "FAILED_ROUTE_HIERARCHY_PREFactor_NOT_DERIVED"
	StatusFailedHierarchyStillNotDerived    = "FAILED_ROUTE_HIERARCHY_SCALING_FACTOR_STILL_NOT_DERIVED"
	StatusFailedF2MomentStillUnlocked       = "FAILED_ROUTE_F2_CUTOFF_MOMENT_STILL_UNLOCKED"
	StatusFailedElectroweakVEVNotClaimed    = "FAILED_ROUTE_ELECTROWEAK_VEV_NOT_DERIVED"
)

const (
	inheritedHighestGate = 339

	electroweakVEVGeV  = 246.22
	unreducedPlanckGeV = 1.220890e19

	booleanProjectorRank = 56
	exteriorDim70        = 70

	sTop = 8 * math.Pi * math.Pi
)

type Inputs struct {
	HighestInheritedGate int
	ElectroweakVEVGeV    float64
	UnreducedPlanckGeV   float64
	ReducedPlanckGeV     float64
	STop                 float64
	BooleanProjectorRank int
	ExteriorDimension70  int
	Status               string
}

type Targets struct {
	RhoUnreduced  float64
	RhoReduced    float64
	Log2Unreduced float64
	Log2Reduced   float64
	LogEUnreduced float64
	LogEReduced   float64
	Status        string
}

type PromotionCandidate struct {
	Name                  string
	Expression            string
	Value                 float64
	RatioToUnreduced      float64
	RequiredPrefUnreduced float64
	RatioToReduced        float64
	RequiredPrefReduced   float64
	Log10ErrorUnreduced   float64
	Category              string
	HasNativeInvariant    bool
	HasScaleLawTheorem    bool
	HasPrefactorTheorem   bool
	Promotable            bool
	Verdict               string
}

type PromotionLedger struct {
	Candidates       []PromotionCandidate
	ClosestUnreduced PromotionCandidate
	Rank56           PromotionCandidate
	HalfTopological  PromotionCandidate
	Status           string
}

type PrefactorSieve struct {
	RequiredForRank56        float64
	RequiredForHalfAction    float64
	SqrtTwoRank56Value       float64
	SqrtTwoRank56Ratio       float64
	PiOverTwoRank56Value     float64
	PiOverTwoRank56Ratio     float64
	BestAccidentalExpression string
	BestAccidentalRatio      float64
	NativePrefactorDerived   bool
	Status                   string
}

type CategoryFirewall struct {
	RankExponentControlsMassScale     bool
	HalfTopologicalActionControlsVEV  bool
	PrefactorSelectedByFiniteGeometry bool
	ArbitraryExponentFittingRejected  bool
	F2MomentStillUnlocked             bool
	Status                            string
}

type Summary struct {
	DirectAnswer      string
	BestNearMiss      string
	EffectiveExponent string
	NextGate          string
	Status            string
}

type Analysis struct {
	Inputs     Inputs
	Targets    Targets
	Ledger     PromotionLedger
	Prefactors PrefactorSieve
	Firewalls  CategoryFirewall
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
	targets := computeTargets(inputs)
	ledger := auditPromotionCandidates(targets)
	prefactors := executePrefactorSieve(targets)
	firewalls := preserveFirewalls()
	summary := compileSummary(targets, ledger, prefactors)
	truth := "Gate 340 audits the two strongest hierarchy near misses from Gate 339: 2^-56 and exp(-4π²). The unreduced Planck target corresponds to an effective binary exponent n=55.46076288096928, so rank 56 is close but not exact; the required prefactor is 1.4532038761902069. The half-topological action exp(-4π²) requires an independent prefactor 2.817771098178961. No prior ASHA theorem states that Boolean rank, half-instanton action, or either prefactor controls the electroweak VEV. Therefore the near misses remain cataloged but unpromoted, and the hierarchy problem remains tied to an unresolved f2/Planck-normalization theorem."
	return Analysis{Inputs: inputs, Targets: targets, Ledger: ledger, Prefactors: prefactors, Firewalls: firewalls, Summary: summary, Truth: truth}, nil
}

func compileInputs() Inputs {
	reduced := unreducedPlanckGeV / math.Sqrt(8*math.Pi)
	return Inputs{
		HighestInheritedGate: inheritedHighestGate,
		ElectroweakVEVGeV:    electroweakVEVGeV,
		UnreducedPlanckGeV:   unreducedPlanckGeV,
		ReducedPlanckGeV:     reduced,
		STop:                 sTop,
		BooleanProjectorRank: booleanProjectorRank,
		ExteriorDimension70:  exteriorDim70,
		Status:               StatusGate339Inherited,
	}
}

func computeTargets(i Inputs) Targets {
	rhoU := i.ElectroweakVEVGeV / i.UnreducedPlanckGeV
	rhoR := i.ElectroweakVEVGeV / i.ReducedPlanckGeV
	return Targets{
		RhoUnreduced:  rhoU,
		RhoReduced:    rhoR,
		Log2Unreduced: -math.Log2(rhoU),
		Log2Reduced:   -math.Log2(rhoR),
		LogEUnreduced: -math.Log(rhoU),
		LogEReduced:   -math.Log(rhoR),
		Status:        StatusEffectiveExponentComputed,
	}
}

func auditPromotionCandidates(t Targets) PromotionLedger {
	raw := []struct {
		name, expr, cat, verdict string
		value                    float64
		native, scale, pref      bool
	}{
		{"rank-56 Boolean projector power", "2^-56", "finite rank / Boolean capacity", "closest native near miss but no theorem maps rank exponent to v/M_P", math.Pow(2, -56), true, false, false},
		{"half topological action exponential", "exp(-S_top/2)=exp(-4π²)", "topological action / instanton half-action", "near target but half-action square-root rule is not derived", math.Exp(-sTop / 2), true, false, false},
		{"full topological action exponential", "exp(-S_top)=exp(-8π²)", "topological action", "far too small; standard instanton exponential cannot be the hierarchy", math.Exp(-sTop), true, false, false},
		{"rank-70 exterior power", "2^-70", "exterior algebra dimension", "too small and no theorem maps Λ^4R^8 dimension to mass hierarchy", math.Pow(2, -70), true, false, false},
		{"required fitted binary power", "2^-55.46076288096928", "fit lane", "matches by construction and is rejected as arbitrary exponent fitting", t.RhoUnreduced, false, false, false},
		{"sqrt-two rank-56 repair", "sqrt(2)·2^-56", "prefactor repair", "very close but sqrt(2) prefactor is not selected by a hierarchy theorem", math.Sqrt2 * math.Pow(2, -56), true, false, false},
		{"pi-over-two rank-56 repair", "(π/2)·2^-56", "prefactor repair", "overshoots and lacks a native scale-law justification", (math.Pi / 2) * math.Pow(2, -56), true, false, false},
	}
	candidates := make([]PromotionCandidate, 0, len(raw))
	var rank56, half PromotionCandidate
	for _, r := range raw {
		c := makeCandidate(r.name, r.expr, r.cat, r.verdict, r.value, t, r.native, r.scale, r.pref)
		candidates = append(candidates, c)
		if r.name == "rank-56 Boolean projector power" {
			rank56 = c
		}
		if r.name == "half topological action exponential" {
			half = c
		}
	}
	sort.Slice(candidates, func(i, j int) bool {
		return math.Abs(candidates[i].Log10ErrorUnreduced) < math.Abs(candidates[j].Log10ErrorUnreduced)
	})
	return PromotionLedger{Candidates: candidates, ClosestUnreduced: closestNonFit(candidates), Rank56: rank56, HalfTopological: half, Status: StatusRank56PromotionAudited}
}

func closestNonFit(candidates []PromotionCandidate) PromotionCandidate {
	for _, c := range candidates {
		if c.Category != "fit lane" {
			return c
		}
	}
	return candidates[0]
}

func makeCandidate(name, expr, cat, verdict string, value float64, t Targets, native, scale, pref bool) PromotionCandidate {
	return PromotionCandidate{
		Name:                  name,
		Expression:            expr,
		Value:                 value,
		RatioToUnreduced:      value / t.RhoUnreduced,
		RequiredPrefUnreduced: t.RhoUnreduced / value,
		RatioToReduced:        value / t.RhoReduced,
		RequiredPrefReduced:   t.RhoReduced / value,
		Log10ErrorUnreduced:   math.Log10(value / t.RhoUnreduced),
		Category:              cat,
		HasNativeInvariant:    native,
		HasScaleLawTheorem:    scale,
		HasPrefactorTheorem:   pref,
		Promotable:            native && scale && pref,
		Verdict:               verdict,
	}
}

func executePrefactorSieve(t Targets) PrefactorSieve {
	rank56 := math.Pow(2, -56)
	halfAction := math.Exp(-sTop / 2)
	sqrtTwo := math.Sqrt2 * rank56
	piOverTwo := (math.Pi / 2) * rank56
	best := "sqrt(2)·2^-56"
	bestRatio := sqrtTwo / t.RhoUnreduced
	if math.Abs(math.Log(piOverTwo/t.RhoUnreduced)) < math.Abs(math.Log(bestRatio)) {
		best = "(π/2)·2^-56"
		bestRatio = piOverTwo / t.RhoUnreduced
	}
	return PrefactorSieve{
		RequiredForRank56:        t.RhoUnreduced / rank56,
		RequiredForHalfAction:    t.RhoUnreduced / halfAction,
		SqrtTwoRank56Value:       sqrtTwo,
		SqrtTwoRank56Ratio:       sqrtTwo / t.RhoUnreduced,
		PiOverTwoRank56Value:     piOverTwo,
		PiOverTwoRank56Ratio:     piOverTwo / t.RhoUnreduced,
		BestAccidentalExpression: best,
		BestAccidentalRatio:      bestRatio,
		NativePrefactorDerived:   false,
		Status:                   StatusPrefactorSieveExecuted,
	}
}

func preserveFirewalls() CategoryFirewall {
	return CategoryFirewall{
		RankExponentControlsMassScale:     false,
		HalfTopologicalActionControlsVEV:  false,
		PrefactorSelectedByFiniteGeometry: false,
		ArbitraryExponentFittingRejected:  true,
		F2MomentStillUnlocked:             true,
		Status:                            StatusCategoryFirewallPreserved,
	}
}

func compileSummary(t Targets, l PromotionLedger, p PrefactorSieve) Summary {
	return Summary{
		DirectAnswer:      "Gate 340 does not promote the hierarchy near misses. Rank 56 and half-topological action are numerically suggestive, but neither has a native theorem tying it to f2, Newton's constant, or the electroweak VEV.",
		BestNearMiss:      fmt.Sprintf("%s = %.15e; ratio_to_unreduced=%.12f", l.ClosestUnreduced.Name, l.ClosestUnreduced.Value, l.ClosestUnreduced.RatioToUnreduced),
		EffectiveExponent: fmt.Sprintf("target exponent n=-log2(v/M_P)=%.14f; rank56_delta=%.14f", t.Log2Unreduced, 56-t.Log2Unreduced),
		NextGate:          "Audit the gravitational Seeley-de Witt a2 coefficient and f2 moment: derive or reject a native Newton/electroweak scale relation rather than fitting hierarchy powers.",
		Status:            StatusFailedHierarchyStillNotDerived,
	}
}

func Statuses(a Analysis) []string {
	return []string{
		a.Inputs.Status,
		a.Targets.Status,
		a.Ledger.Status,
		StatusHalfInstantonPromotionAudit,
		a.Prefactors.Status,
		a.Firewalls.Status,
		StatusTensionRank56NearButNotExact,
		StatusTensionHalfActionNearButNotExact,
		StatusTensionSqrtTwoPrefactorNearMiss,
		StatusTensionNoScaleLawSelector,
		StatusFailedRank56ScaleLawNotDerived,
		StatusFailedHalfInstantonRuleNotDerived,
		StatusFailedPrefactorNotDerived,
		StatusFailedHierarchyStillNotDerived,
		StatusFailedF2MomentStillUnlocked,
		StatusFailedElectroweakVEVNotClaimed,
	}
}

func FormatInputs(i Inputs) string {
	return fmt.Sprintf("highest_gate=%d; v=%.12f GeV; M_P=%.12e GeV; Mbar_P=%.12e GeV; S_top=%.12f; rank=%d; dim70=%d; status=%s", i.HighestInheritedGate, i.ElectroweakVEVGeV, i.UnreducedPlanckGeV, i.ReducedPlanckGeV, i.STop, i.BooleanProjectorRank, i.ExteriorDimension70, i.Status)
}

func FormatTargets(t Targets) string {
	return fmt.Sprintf("rho_unreduced=%.15e; rho_reduced=%.15e; n2_unred=%.14f; n2_red=%.14f; e_unred=%.14f; e_red=%.14f; status=%s", t.RhoUnreduced, t.RhoReduced, t.Log2Unreduced, t.Log2Reduced, t.LogEUnreduced, t.LogEReduced, t.Status)
}

func FormatCandidate(c PromotionCandidate) string {
	return fmt.Sprintf("%s: %s = %.15e; ratio_unred=%.12f; pref_unred=%.12f; ratio_red=%.12f; log10err=%+.9f; category=%s; native=%t; scale_law=%t; prefactor_theorem=%t; promotable=%t; verdict=%s", c.Name, c.Expression, c.Value, c.RatioToUnreduced, c.RequiredPrefUnreduced, c.RatioToReduced, c.Log10ErrorUnreduced, c.Category, c.HasNativeInvariant, c.HasScaleLawTheorem, c.HasPrefactorTheorem, c.Promotable, c.Verdict)
}

func FormatLedger(l PromotionLedger) string {
	parts := []string{fmt.Sprintf("status=%s", l.Status), fmt.Sprintf("closest={%s}", FormatCandidate(l.ClosestUnreduced)), fmt.Sprintf("rank56={%s}", FormatCandidate(l.Rank56)), fmt.Sprintf("half={%s}", FormatCandidate(l.HalfTopological))}
	for _, c := range l.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return strings.Join(parts, "\n")
}

func FormatPrefactors(p PrefactorSieve) string {
	return fmt.Sprintf("pref_rank56=%.15f; pref_half_action=%.15f; sqrt2_rank56=%.15e ratio=%.12f; pi_over_two_rank56=%.15e ratio=%.12f; best_accidental=%s ratio=%.12f; native_prefactor=%t; status=%s", p.RequiredForRank56, p.RequiredForHalfAction, p.SqrtTwoRank56Value, p.SqrtTwoRank56Ratio, p.PiOverTwoRank56Value, p.PiOverTwoRank56Ratio, p.BestAccidentalExpression, p.BestAccidentalRatio, p.NativePrefactorDerived, p.Status)
}

func FormatFirewalls(f CategoryFirewall) string {
	return fmt.Sprintf("rank_controls_scale=%t; half_action_controls_vev=%t; prefactor_selected=%t; reject_fit=%t; f2_unlocked=%t; status=%s", f.RankExponentControlsMassScale, f.HalfTopologicalActionControlsVEV, f.PrefactorSelectedByFiniteGeometry, f.ArbitraryExponentFittingRejected, f.F2MomentStillUnlocked, f.Status)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("direct=%s; best=%s; exponent=%s; next=%s; status=%s", s.DirectAnswer, s.BestNearMiss, s.EffectiveExponent, s.NextGate, s.Status)
}

func FormatStatuses(statuses []string) string { return strings.Join(statuses, "\n") }
