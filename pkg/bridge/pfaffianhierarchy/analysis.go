// Package pfaffianhierarchy implements Gate 341:
// Pfaffian Half-Action Hierarchy / Fermionic Fluctuation Determinant Derivation.
//
// Gate 340 cataloged exp(-4π²) and power-of-two prefactors as separate hierarchy
// near-misses. Gate 341 audits the combined semiclassical expression
//
//	ρ = v/M_P := 2^(N_gen/2) exp(-S_top/2)
//
// where N_gen=3 is inherited from triality and S_top=8π² from the topological
// action ledger. The gate distinguishes finite-core inputs from continuum path-
// integral measure inputs: Pfaffian half-action and zero-mode Gaussian factors
// are standard semiclassical structures, but still must be explicitly selected
// as the hierarchy mechanism before the VEV/Planck ratio is unconditionally
// derived.
package pfaffianhierarchy

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE341-PFAFFIAN-HALF-ACTION-HIERARCHY-FERMIONIC-FLUCTUATION-DETERMINANT"

	StatusGate340Inherited                = "CONDITIONAL_SUPPORT_GATE340_HIERARCHY_PROMOTION_AUDIT_INHERITED"
	StatusPfaffianHalfActionFormalized    = "CONDITIONAL_SUPPORT_PFAFFIAN_HALF_ACTION_FORMALIZED"
	StatusGenerationFluctuationFormalized = "CONDITIONAL_SUPPORT_GENERATION_FLUCTUATION_FACTOR_FORMALIZED"
	StatusCombinedHierarchyComputed       = "CONDITIONAL_SUPPORT_COMBINED_PFAFFIAN_HIERARCHY_FACTOR_COMPUTED"
	StatusGravityConnectionFormalized     = "CONDITIONAL_SUPPORT_GRAVITY_ELECTROWEAK_RATIO_FORMALIZED"
	StatusPrecisionComparisonExecuted     = "CONDITIONAL_SUPPORT_HIERARCHY_PRECISION_COMPARISON_EXECUTED"

	StatusTensionPfaffianMechanismExternal       = "CONDITIONAL_TENSION_PFAFFIAN_MEASURE_IS_CONTINUUM_PATH_INTEGRAL_INPUT"
	StatusTensionZeroModeNormalizationExternal   = "CONDITIONAL_TENSION_ZERO_MODE_SQRT2_PER_GENERATION_NOT_FINITE_CORE_DERIVED"
	StatusTensionReducedPlanckBranchNotSelected  = "CONDITIONAL_TENSION_REDUCED_PLANCK_BRANCH_NOT_SELECTED"
	StatusTensionF2MomentReinterpretedNotDerived = "CONDITIONAL_TENSION_F2_MOMENT_REINTERPRETED_BUT_NOT_DERIVED"

	StatusFailedUnconditionalHierarchyNotClaimed = "FAILED_ROUTE_UNCONDITIONAL_HIERARCHY_SCALING_FACTOR_NOT_CLAIMED"
	StatusFailedPfaffianActionNotFiniteCore      = "FAILED_ROUTE_PFAFFIAN_HALF_ACTION_NOT_DERIVED_FROM_FINITE_CORE"
	StatusFailedZeroModeCountNotDerived          = "FAILED_ROUTE_FERMIONIC_ZERO_MODE_NORMALIZATION_NOT_DERIVED_FROM_FINITE_CORE"
	StatusFailedF2MomentStillNotLocked           = "FAILED_ROUTE_F2_CUTOFF_MOMENT_STILL_NOT_LOCKED"
	StatusFailedElectroweakVEVNotClaimed         = "FAILED_ROUTE_ELECTROWEAK_VEV_NOT_DERIVED_UNCONDITIONALLY"
)

const (
	inheritedHighestGate = 340

	nGen               = 3
	electroweakVEVGeV  = 246.22
	unreducedPlanckGeV = 1.220890e19
	sTop               = 8 * math.Pi * math.Pi
)

type Inputs struct {
	HighestInheritedGate int
	NGen                 int
	STop                 float64
	ElectroweakVEVGeV    float64
	UnreducedPlanckGeV   float64
	ReducedPlanckGeV     float64
	Status               string
}

type PfaffianHalfAction struct {
	FullAction           float64
	HalfAction           float64
	FullExponential      float64
	HalfExponential      float64
	PfaffianRule         string
	HalfActionAuthorized bool
	FiniteCoreDerived    bool
	Status               string
}

type GenerationFluctuation struct {
	NGen                    int
	PerGenerationFactor     float64
	CombinedFactor          float64
	Rule                    string
	ZeroModeNormalizationOK bool
	FiniteCoreDerived       bool
	Status                  string
}

type HierarchyPrediction struct {
	PredictedRatio         float64
	ObservedUnreducedRatio float64
	ObservedReducedRatio   float64
	RatioToUnreducedTarget float64
	RelativeErrorUnreduced float64
	RatioToReducedTarget   float64
	RelativeErrorReduced   float64
	PredictedPlanckFromVEV float64
	PlanckDifferenceGeV    float64
	Status                 string
}

type GravityConnection struct {
	Formula                    string
	ElectroweakToGravityLinked bool
	F2MomentLocked             bool
	PlanckBranchSelected       string
	Interpretation             string
	Status                     string
}

type Summary struct {
	DirectAnswer string
	Prediction   string
	Agreement    string
	Caveat       string
	NextGate     string
	Status       string
}

type Analysis struct {
	Inputs     Inputs
	Pfaffian   PfaffianHalfAction
	Generation GenerationFluctuation
	Prediction HierarchyPrediction
	Gravity    GravityConnection
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
	pf := formalizePfaffian(inputs)
	gen := formalizeGenerationFluctuation(inputs)
	pred := computeHierarchyPrediction(inputs, pf, gen)
	grav := formalizeGravityConnection(pred)
	summary := compileSummary(pred)
	truth := "Gate 341 tests the combined semiclassical hierarchy law ρ=2^(N_gen/2)exp(-S_top/2). With N_gen=3 and S_top=8π², the prediction is 2.024352198454697e-17, within 0.378172% of v/M_P using the unreduced Planck mass. This is a strong conditional hierarchy witness. However, the half-action Pfaffian rule and sqrt(2)-per-generation fluctuation factor are continuum path-integral measure inputs, not yet finite-core theorems selecting f2 or Newton normalization. The gate therefore logs conditional support, not an unconditional electroweak VEV derivation."
	return Analysis{Inputs: inputs, Pfaffian: pf, Generation: gen, Prediction: pred, Gravity: grav, Summary: summary, Truth: truth}, nil
}

func compileInputs() Inputs {
	return Inputs{
		HighestInheritedGate: inheritedHighestGate,
		NGen:                 nGen,
		STop:                 sTop,
		ElectroweakVEVGeV:    electroweakVEVGeV,
		UnreducedPlanckGeV:   unreducedPlanckGeV,
		ReducedPlanckGeV:     unreducedPlanckGeV / math.Sqrt(8*math.Pi),
		Status:               StatusGate340Inherited,
	}
}

func formalizePfaffian(i Inputs) PfaffianHalfAction {
	return PfaffianHalfAction{
		FullAction:           i.STop,
		HalfAction:           i.STop / 2,
		FullExponential:      math.Exp(-i.STop),
		HalfExponential:      math.Exp(-i.STop / 2),
		PfaffianRule:         "for real/Majorana fermions: Z_F ∝ pf(D) = det(D)^(1/2), giving exp(-S_top/2) rather than exp(-S_top)",
		HalfActionAuthorized: true,
		FiniteCoreDerived:    false,
		Status:               StatusPfaffianHalfActionFormalized,
	}
}

func formalizeGenerationFluctuation(i Inputs) GenerationFluctuation {
	per := math.Sqrt2
	combined := math.Pow(2, float64(i.NGen)/2)
	return GenerationFluctuation{
		NGen:                    i.NGen,
		PerGenerationFactor:     per,
		CombinedFactor:          combined,
		Rule:                    "semiclassical Gaussian/zero-mode prefactor: each chiral generation contributes sqrt(2), so N_gen generations contribute 2^(N_gen/2)",
		ZeroModeNormalizationOK: true,
		FiniteCoreDerived:       false,
		Status:                  StatusGenerationFluctuationFormalized,
	}
}

func computeHierarchyPrediction(i Inputs, p PfaffianHalfAction, g GenerationFluctuation) HierarchyPrediction {
	pred := g.CombinedFactor * p.HalfExponential
	obsU := i.ElectroweakVEVGeV / i.UnreducedPlanckGeV
	obsR := i.ElectroweakVEVGeV / i.ReducedPlanckGeV
	predMP := i.ElectroweakVEVGeV / pred
	return HierarchyPrediction{
		PredictedRatio:         pred,
		ObservedUnreducedRatio: obsU,
		ObservedReducedRatio:   obsR,
		RatioToUnreducedTarget: pred / obsU,
		RelativeErrorUnreduced: (pred - obsU) / obsU,
		RatioToReducedTarget:   pred / obsR,
		RelativeErrorReduced:   (pred - obsR) / obsR,
		PredictedPlanckFromVEV: predMP,
		PlanckDifferenceGeV:    predMP - i.UnreducedPlanckGeV,
		Status:                 StatusCombinedHierarchyComputed,
	}
}

func formalizeGravityConnection(p HierarchyPrediction) GravityConnection {
	return GravityConnection{
		Formula:                    "v/M_P = 2^(N_gen/2) exp(-S_top/2)",
		ElectroweakToGravityLinked: true,
		F2MomentLocked:             false,
		PlanckBranchSelected:       "unreduced Planck mass branch; reduced branch is off by the expected sqrt(8π) conversion",
		Interpretation:             "the formula is a conditional finite-topology/fermionic-measure bridge from electroweak scale to gravitational scale, but it does not by itself derive the gravitational Seeley-de Witt f2 moment or Newton normalization",
		Status:                     StatusGravityConnectionFormalized,
	}
}

func compileSummary(p HierarchyPrediction) Summary {
	return Summary{
		DirectAnswer: "Gate 341 conditionally promotes the hierarchy near miss by combining the Pfaffian half-action exp(-S_top/2) with the three-generation Gaussian factor 2^(3/2).",
		Prediction:   fmt.Sprintf("rho_pred=%.15e; M_P_from_v=%.12e GeV", p.PredictedRatio, p.PredictedPlanckFromVEV),
		Agreement:    fmt.Sprintf("ratio_to_unreduced=%.12f; relative_error=%+.9f%%", p.RatioToUnreducedTarget, 100*p.RelativeErrorUnreduced),
		Caveat:       "Pfaffian and zero-mode factors are standard continuum measure inputs; native f2/Newton normalization remains firewalled.",
		NextGate:     "Derive the gravitational Seeley-de Witt f2/Newton coefficient or explicitly seal the Pfaffian hierarchy formula as the Phase-III scale law.",
		Status:       StatusCombinedHierarchyComputed,
	}
}

func Statuses(a Analysis) []string {
	return []string{
		a.Inputs.Status,
		a.Pfaffian.Status,
		a.Generation.Status,
		a.Prediction.Status,
		a.Gravity.Status,
		StatusPrecisionComparisonExecuted,
		StatusTensionPfaffianMechanismExternal,
		StatusTensionZeroModeNormalizationExternal,
		StatusTensionReducedPlanckBranchNotSelected,
		StatusTensionF2MomentReinterpretedNotDerived,
		StatusFailedUnconditionalHierarchyNotClaimed,
		StatusFailedPfaffianActionNotFiniteCore,
		StatusFailedZeroModeCountNotDerived,
		StatusFailedF2MomentStillNotLocked,
		StatusFailedElectroweakVEVNotClaimed,
	}
}

func FormatInputs(i Inputs) string {
	return fmt.Sprintf("highest_gate=%d; N_gen=%d; S_top=%.15f; v=%.12f GeV; M_P=%.12e GeV; Mbar_P=%.12e GeV; status=%s", i.HighestInheritedGate, i.NGen, i.STop, i.ElectroweakVEVGeV, i.UnreducedPlanckGeV, i.ReducedPlanckGeV, i.Status)
}

func FormatPfaffian(p PfaffianHalfAction) string {
	return fmt.Sprintf("S=%.15f; S/2=%.15f; exp(-S)=%.15e; exp(-S/2)=%.15e; rule=%s; authorized=%t; finite_core=%t; status=%s", p.FullAction, p.HalfAction, p.FullExponential, p.HalfExponential, p.PfaffianRule, p.HalfActionAuthorized, p.FiniteCoreDerived, p.Status)
}

func FormatGeneration(g GenerationFluctuation) string {
	return fmt.Sprintf("N_gen=%d; per_generation=%.15f; combined=%.15f; rule=%s; zero_mode_ok=%t; finite_core=%t; status=%s", g.NGen, g.PerGenerationFactor, g.CombinedFactor, g.Rule, g.ZeroModeNormalizationOK, g.FiniteCoreDerived, g.Status)
}

func FormatPrediction(p HierarchyPrediction) string {
	return fmt.Sprintf("rho_pred=%.15e; rho_unred=%.15e; ratio_unred=%.12f; rel_unred=%+.12f; rho_red=%.15e; ratio_red=%.12f; M_P_from_v=%.12e; delta_M_P=%.12e; status=%s", p.PredictedRatio, p.ObservedUnreducedRatio, p.RatioToUnreducedTarget, p.RelativeErrorUnreduced, p.ObservedReducedRatio, p.RatioToReducedTarget, p.PredictedPlanckFromVEV, p.PlanckDifferenceGeV, p.Status)
}

func FormatGravity(g GravityConnection) string {
	return fmt.Sprintf("formula=%s; ew_gravity_link=%t; f2_locked=%t; branch=%s; interpretation=%s; status=%s", g.Formula, g.ElectroweakToGravityLinked, g.F2MomentLocked, g.PlanckBranchSelected, g.Interpretation, g.Status)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("direct=%s; prediction=%s; agreement=%s; caveat=%s; next=%s; status=%s", s.DirectAnswer, s.Prediction, s.Agreement, s.Caveat, s.NextGate, s.Status)
}

func FormatStatuses(statuses []string) string { return strings.Join(statuses, "\n") }
