// Package gravityspectralactionf2 implements Gate 343:
// Gravitational Spectral Action / f2 Cutoff Moment Sieve.
//
// Gate 342 conditionally derived the electroweak-to-unreduced-Planck hierarchy
//
//	v/M_P = 2^(3/2) exp(-4π²)
//
// from the topological half-action and finite triality Pfaffian measure.  Gate
// 343 maps that ratio back into the gravitational sector of the spectral action.
// The important distinction is that the Einstein-Hilbert coefficient fixes the
// dimensionful product f2 Λ²; it does not fix f2 alone unless a native theorem
// selects the cutoff scale Λ.  This gate therefore extracts the exact invariant
// obligation and audits common scale choices without forcing them.
package gravityspectralactionf2

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE343-GRAVITATIONAL-SPECTRAL-ACTION-F2-CUTOFF-MOMENT-SIEVE"

	StatusGate342Inherited                = "CONDITIONAL_SUPPORT_GATE342_HIERARCHY_RATIO_INHERITED"
	StatusEinsteinHilbertFormalized       = "CONDITIONAL_SUPPORT_EINSTEIN_HILBERT_SPECTRAL_ACTION_FORMALIZED"
	StatusF2MomentTargetExtracted         = "CONDITIONAL_SUPPORT_F2_MOMENT_TARGET_EXTRACTED"
	StatusScaleChoiceSieveExecuted        = "CONDITIONAL_SUPPORT_CUTOFF_SCALE_CHOICE_SIEVE_EXECUTED"
	StatusGeometricResonanceAuditExecuted = "CONDITIONAL_SUPPORT_GEOMETRIC_RESONANCE_AUDIT_EXECUTED"
	StatusF2LambdaInvariantDerived        = "CONDITIONAL_SUPPORT_F2_LAMBDA_PRODUCT_INVARIANT_DERIVED"

	StatusTensionF2NotAlone                   = "CONDITIONAL_TENSION_F2_NOT_ISOLATED_WITHOUT_LAMBDA_SCALE_THEOREM"
	StatusTensionPlanckCutoffGivesPiOver64    = "CONDITIONAL_TENSION_PLANCK_CUTOFF_GIVES_SIMPLE_PI_OVER_64_TARGET"
	StatusTensionNoKnownInvariantMatch        = "CONDITIONAL_TENSION_NO_KNOWN_NATIVE_INVARIANT_MATCHES_F2_TARGET"
	StatusTensionCurvatureEndomorphismIgnored = "CONDITIONAL_TENSION_CURVATURE_ENDOMORPHISM_C_TERM_IGNORED_IN_LEADING_LEDGER"

	StatusFailedF2MomentStillUnlocked       = "FAILED_ROUTE_F2_CUTOFF_MOMENT_STILL_UNLOCKED"
	StatusFailedLambdaCutoffScaleNotDerived = "FAILED_ROUTE_CUTOFF_SCALE_LAMBDA_NOT_DERIVED"
	StatusFailedNewtonConstantNotDerived    = "FAILED_ROUTE_NEWTON_CONSTANT_NORMALIZATION_NOT_DERIVED_UNCONDITIONALLY"
	StatusFailedNativeF2ResonanceNotFound   = "FAILED_ROUTE_NATIVE_F2_RESONANCE_NOT_FOUND"
	StatusFailedCosmologicalF4Firewall      = "FAILED_ROUTE_COSMOLOGICAL_CONSTANT_F4_CHANNEL_STILL_FIREWALLED"
)

const (
	inheritedHighestGate = 342

	nGen               = 3
	electroweakVEVGeV  = 246.22
	unreducedPlanckGeV = 1.220890e19
	sTop               = 8 * math.Pi * math.Pi
	f0Contact          = 7.0
	bGap               = 0.102464921191
	contactOmega2      = 61.0 / 25.0
)

type Inputs struct {
	HighestInheritedGate int
	NGen                 int
	STop                 float64
	VEVGeV               float64
	UnreducedPlanckGeV   float64
	ReducedPlanckGeV     float64
	HierarchyPredicted   float64
	HierarchyObserved    float64
	Status               string
}

type EinsteinHilbertLedger struct {
	SpectralCoefficientCG float64
	FormulaReduced        string
	FormulaUnreduced      string
	LeadingAssumptions    string
	Status                string
}

type F2Target struct {
	InvariantF2LambdaOverReducedPlanck float64
	InvariantF2LambdaOverUnreduced     float64
	ProductF2LambdaSquaredGeV2         float64
	ProductF2LambdaSquaredOverV2       float64
	PlanckCutoffF2Target               float64
	ReducedPlanckCutoffF2Target        float64
	VEVCutoffF2Target                  float64
	Rule                               string
	Status                             string
}

type ScaleCandidate struct {
	Name           string
	LambdaGeV      float64
	F2Required     float64
	Dimensionless  bool
	Interpretation string
	Status         string
}

type ResonanceCandidate struct {
	Name          string
	Value         float64
	RelativeError float64
	Promoted      bool
	Reason        string
}

type GeometricAudit struct {
	Candidates       []ResonanceCandidate
	BestCandidate    ResonanceCandidate
	NativeMatchFound bool
	Status           string
}

type Firewall struct {
	F2Locked              bool
	LambdaLocked          bool
	NewtonConstantDerived bool
	CosmologicalF4Locked  bool
	Explanation           string
	Status                string
}

type Summary struct {
	DirectAnswer string
	Invariant    string
	PlanckChoice string
	Caveat       string
	NextGate     string
	Status       string
}

type Analysis struct {
	Inputs     Inputs
	EH         EinsteinHilbertLedger
	Target     F2Target
	ScaleSieve []ScaleCandidate
	Resonance  GeometricAudit
	Firewall   Firewall
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
	eh := formalizeEinsteinHilbert()
	target := extractF2Target(inputs, eh)
	scaleSieve := auditScaleChoices(inputs, eh, target)
	resonance := auditGeometricResonances(target)
	firewall := compileFirewall(resonance)
	summary := compileSummary(target, resonance)
	truth := "Gate 343 maps the Gate 342 hierarchy ratio into the gravitational spectral-action ledger. In the leading Einstein-Hilbert channel, Mbar_P^2=(8/π²) f2 Λ², equivalently M_P^2=(64/π) f2 Λ². Thus the geometry fixes the invariant f2(Λ/M_P)^2=π/64, or f2=π/64 only if Λ is natively identified with the unreduced Planck scale. No existing finite invariant in the ledger uniquely equals this target, so f2 and Newton normalization remain firewalled while the precise mathematical obligation is now explicit."
	return Analysis{Inputs: inputs, EH: eh, Target: target, ScaleSieve: scaleSieve, Resonance: resonance, Firewall: firewall, Summary: summary, Truth: truth}, nil
}

func compileInputs() Inputs {
	reduced := unreducedPlanckGeV / math.Sqrt(8*math.Pi)
	hierarchyPred := math.Pow(2, float64(nGen)/2) * math.Exp(-sTop/2)
	hierarchyObs := electroweakVEVGeV / unreducedPlanckGeV
	return Inputs{HighestInheritedGate: inheritedHighestGate, NGen: nGen, STop: sTop, VEVGeV: electroweakVEVGeV, UnreducedPlanckGeV: unreducedPlanckGeV, ReducedPlanckGeV: reduced, HierarchyPredicted: hierarchyPred, HierarchyObserved: hierarchyObs, Status: StatusGate342Inherited}
}

func formalizeEinsteinHilbert() EinsteinHilbertLedger {
	cG := 8 / (math.Pi * math.Pi)
	return EinsteinHilbertLedger{
		SpectralCoefficientCG: cG,
		FormulaReduced:        "Mbar_P^2 = (8/π²) f2 Λ² in the leading a2 Einstein-Hilbert ledger",
		FormulaUnreduced:      "M_P^2 = 8π Mbar_P² = (64/π) f2 Λ²",
		LeadingAssumptions:    "Euclidean-to-Lorentzian sign fixed; curvature-endomorphism c/Higgs-trace correction ignored in this leading gravitational coefficient; cosmological f4 channel excluded",
		Status:                StatusEinsteinHilbertFormalized,
	}
}

func extractF2Target(i Inputs, eh EinsteinHilbertLedger) F2Target {
	product := (i.ReducedPlanckGeV * i.ReducedPlanckGeV) / eh.SpectralCoefficientCG
	return F2Target{
		InvariantF2LambdaOverReducedPlanck: 1 / eh.SpectralCoefficientCG, // f2 (Λ/Mbar)^2
		InvariantF2LambdaOverUnreduced:     math.Pi / 64,                 // f2 (Λ/M)^2
		ProductF2LambdaSquaredGeV2:         product,                      // f2 Λ²
		ProductF2LambdaSquaredOverV2:       product / (i.VEVGeV * i.VEVGeV),
		PlanckCutoffF2Target:               math.Pi / 64,
		ReducedPlanckCutoffF2Target:        1 / eh.SpectralCoefficientCG,
		VEVCutoffF2Target:                  product / (i.VEVGeV * i.VEVGeV),
		Rule:                               "the spectral action determines f2Λ²; f2 alone requires a native Λ selection theorem",
		Status:                             StatusF2MomentTargetExtracted,
	}
}

func auditScaleChoices(i Inputs, eh EinsteinHilbertLedger, t F2Target) []ScaleCandidate {
	return []ScaleCandidate{
		{
			Name:           "Λ = unreduced Planck mass",
			LambdaGeV:      i.UnreducedPlanckGeV,
			F2Required:     t.PlanckCutoffF2Target,
			Dimensionless:  true,
			Interpretation: "simple target f2=π/64; attractive but depends on selecting the unreduced Planck cutoff",
			Status:         StatusTensionPlanckCutoffGivesPiOver64,
		},
		{
			Name:           "Λ = reduced Planck mass",
			LambdaGeV:      i.ReducedPlanckGeV,
			F2Required:     t.ReducedPlanckCutoffF2Target,
			Dimensionless:  true,
			Interpretation: "target f2=π²/8 under reduced Planck cutoff convention",
			Status:         StatusScaleChoiceSieveExecuted,
		},
		{
			Name:           "Λ = electroweak VEV",
			LambdaGeV:      i.VEVGeV,
			F2Required:     t.VEVCutoffF2Target,
			Dimensionless:  true,
			Interpretation: "enormous f2; not a natural cutoff interpretation",
			Status:         StatusScaleChoiceSieveExecuted,
		},
	}
}

func auditGeometricResonances(t F2Target) GeometricAudit {
	target := t.PlanckCutoffF2Target
	candidates := []ResonanceCandidate{
		candidate("π/64 target itself", target, target, false, "identity of the extracted obligation, not an independent geometric invariant"),
		candidate("1/(8π) coupling scale", 1/(8*math.Pi), target, false, "near same circle scale but misses by about 19% and belongs to gauge coupling ledger"),
		candidate("B_gap", bGap, target, false, "B-gap is roughly twice the Planck-cutoff f2 target and belongs to Majorana hierarchy ledger"),
		candidate("f0 contact volume", f0Contact, target, false, "kinetic a4 cutoff moment, wrong order for f2 target"),
		candidate("contact Ω² = 61/25", contactOmega2, target, false, "curvature/contact diagnostic, wrong order"),
		candidate("exp(-S_top/2)", math.Exp(-sTop/2), target, false, "hierarchy exponential, many orders below f2 target"),
	}
	best := candidates[0]
	for _, c := range candidates[1:] {
		if math.Abs(c.RelativeError) < math.Abs(best.RelativeError) && c.Name != "π/64 target itself" {
			best = c
		}
	}
	if best.Name == "π/64 target itself" && len(candidates) > 1 {
		best = candidates[1]
		for _, c := range candidates[2:] {
			if math.Abs(c.RelativeError) < math.Abs(best.RelativeError) && c.Name != "π/64 target itself" {
				best = c
			}
		}
	}
	return GeometricAudit{Candidates: candidates, BestCandidate: best, NativeMatchFound: false, Status: StatusGeometricResonanceAuditExecuted}
}

func candidate(name string, value, target float64, promoted bool, reason string) ResonanceCandidate {
	return ResonanceCandidate{Name: name, Value: value, RelativeError: (value - target) / target, Promoted: promoted, Reason: reason}
}

func compileFirewall(g GeometricAudit) Firewall {
	return Firewall{F2Locked: false, LambdaLocked: false, NewtonConstantDerived: false, CosmologicalF4Locked: false, Explanation: "Gate 343 extracts f2Λ² and the dimensionless Planck-cutoff target π/64, but no independent finite-core theorem selects Λ or identifies π/64 as a native f2 moment; the f4 cosmological channel is untouched.", Status: StatusFailedF2MomentStillUnlocked}
}

func compileSummary(t F2Target, g GeometricAudit) Summary {
	return Summary{
		DirectAnswer: "The gravitational spectral action fixes f2Λ², not f2 by itself.",
		Invariant:    fmt.Sprintf("f2(Λ/M_P)^2=π/64=%.15f; f2Λ²=%.12e GeV²", t.InvariantF2LambdaOverUnreduced, t.ProductF2LambdaSquaredGeV2),
		PlanckChoice: fmt.Sprintf("if Λ=M_P then f2=%.15f; if Λ=Mbar_P then f2=%.15f", t.PlanckCutoffF2Target, t.ReducedPlanckCutoffF2Target),
		Caveat:       fmt.Sprintf("best non-identity native candidate audited: %s = %.15e with relative error %+.6f%%", g.BestCandidate.Name, g.BestCandidate.Value, 100*g.BestCandidate.RelativeError),
		NextGate:     "Derive or reject a native cutoff-scale selector Λ=M_P and a circle-normalized f2=π/64 moment theorem.",
		Status:       StatusF2LambdaInvariantDerived,
	}
}

func Statuses(a Analysis) []string {
	return []string{
		a.Inputs.Status,
		a.EH.Status,
		a.Target.Status,
		StatusF2LambdaInvariantDerived,
		StatusScaleChoiceSieveExecuted,
		a.Resonance.Status,
		StatusTensionF2NotAlone,
		StatusTensionPlanckCutoffGivesPiOver64,
		StatusTensionNoKnownInvariantMatch,
		StatusTensionCurvatureEndomorphismIgnored,
		StatusFailedF2MomentStillUnlocked,
		StatusFailedLambdaCutoffScaleNotDerived,
		StatusFailedNewtonConstantNotDerived,
		StatusFailedNativeF2ResonanceNotFound,
		StatusFailedCosmologicalF4Firewall,
	}
}

func FormatInputs(i Inputs) string {
	return fmt.Sprintf("highest_gate=%d; N_gen=%d; S_top=%.15f; v=%.12f GeV; M_P=%.12e GeV; Mbar_P=%.12e GeV; rho_pred=%.15e; rho_obs=%.15e; status=%s", i.HighestInheritedGate, i.NGen, i.STop, i.VEVGeV, i.UnreducedPlanckGeV, i.ReducedPlanckGeV, i.HierarchyPredicted, i.HierarchyObserved, i.Status)
}

func FormatEH(e EinsteinHilbertLedger) string {
	return fmt.Sprintf("C_G=%.15f; reduced=%s; unreduced=%s; assumptions=%s; status=%s", e.SpectralCoefficientCG, e.FormulaReduced, e.FormulaUnreduced, e.LeadingAssumptions, e.Status)
}

func FormatTarget(t F2Target) string {
	return fmt.Sprintf("f2(Λ/Mbar)^2=%.15f; f2(Λ/M_P)^2=%.15f; f2Λ²=%.12e GeV²; f2Λ²/v²=%.12e; f2@M_P=%.15f; f2@Mbar=%.15f; f2@v=%.12e; rule=%s; status=%s", t.InvariantF2LambdaOverReducedPlanck, t.InvariantF2LambdaOverUnreduced, t.ProductF2LambdaSquaredGeV2, t.ProductF2LambdaSquaredOverV2, t.PlanckCutoffF2Target, t.ReducedPlanckCutoffF2Target, t.VEVCutoffF2Target, t.Rule, t.Status)
}

func FormatScale(s ScaleCandidate) string {
	return fmt.Sprintf("%s: Λ=%.12e GeV; f2_required=%.15e; dimensionless=%t; interpretation=%s; status=%s", s.Name, s.LambdaGeV, s.F2Required, s.Dimensionless, s.Interpretation, s.Status)
}

func FormatScaleSieve(sc []ScaleCandidate) string {
	parts := make([]string, 0, len(sc))
	for _, s := range sc {
		parts = append(parts, FormatScale(s))
	}
	return strings.Join(parts, "\n")
}

func FormatCandidate(c ResonanceCandidate) string {
	return fmt.Sprintf("%s=%.15e; rel_error=%+.9f%%; promoted=%t; reason=%s", c.Name, c.Value, 100*c.RelativeError, c.Promoted, c.Reason)
}

func FormatResonance(g GeometricAudit) string {
	parts := make([]string, 0, len(g.Candidates)+2)
	for _, c := range g.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	parts = append(parts, fmt.Sprintf("best_nonidentity=%s; native_match=%t; status=%s", FormatCandidate(g.BestCandidate), g.NativeMatchFound, g.Status))
	return strings.Join(parts, "\n")
}

func FormatFirewall(f Firewall) string {
	return fmt.Sprintf("f2_locked=%t; lambda_locked=%t; newton=%t; f4_locked=%t; explanation=%s; status=%s", f.F2Locked, f.LambdaLocked, f.NewtonConstantDerived, f.CosmologicalF4Locked, f.Explanation, f.Status)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("direct=%s; invariant=%s; planck_choice=%s; caveat=%s; next=%s; status=%s", s.DirectAnswer, s.Invariant, s.PlanckChoice, s.Caveat, s.NextGate, s.Status)
}

func FormatStatuses(statuses []string) string { return strings.Join(statuses, "\n") }
