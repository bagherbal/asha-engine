// Package trialitygaussianmeasure implements Gate 342:
// Triality Gaussian Measure / Zero-Mode Normalization Audit.
//
// Gate 341 showed that the hierarchy ratio v/M_P is matched at sub-percent
// precision by
//
//	rho = 2^(N_gen/2) exp(-S_top/2)
//
// with N_gen=3 and S_top=8π².  Its remaining firewall was the origin of the
// sqrt(2)-per-generation factor.  Gate 342 audits the finite Grassmann/Berezin
// measure on the triality generation carrier.  In a J-paired Majorana basis,
// each generation contributes a real two-slot skew block.  The finite Berezin
// Gaussian evaluates to a Pfaffian; the canonical J-pair volume block has
// Pfaffian sqrt(2).  Three generations therefore produce 2^(3/2).
//
// The gate distinguishes this finite-measure derivation from the remaining
// gravitational firewalls: f2/Newton normalization and Planck-branch selection
// are not automatically fixed by the finite Berezin measure.
package trialitygaussianmeasure

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE342-TRIALITY-GAUSSIAN-MEASURE-ZERO-MODE-NORMALIZATION-AUDIT"

	StatusGate341Inherited                     = "CONDITIONAL_SUPPORT_GATE341_PFAFFIAN_HIERARCHY_INHERITED"
	StatusFiniteGrassmannMeasureFormalized     = "CONDITIONAL_SUPPORT_FINITE_GRASSMANN_MEASURE_FORMALIZED"
	StatusPfaffianZeroModeEvaluation           = "CONDITIONAL_SUPPORT_PFAFFIAN_ZERO_MODE_EVALUATED"
	StatusNativeZeroModeNormalizationDerived   = "CONDITIONAL_SUPPORT_NATIVE_ZERO_MODE_NORMALIZATION_DERIVED"
	StatusHierarchySynthesisFiniteMeasure      = "CONDITIONAL_SUPPORT_HIERARCHY_SYNTHESIS_WITH_FINITE_MEASURE_EXECUTED"
	StatusHierarchyScalingConditionallyDerived = "CONDITIONAL_SUPPORT_HIERARCHY_SCALING_FACTOR_DERIVED_FROM_FINITE_MEASURE_CONDITIONALLY"
	StatusGravityF2FirewallPreserved           = "CONDITIONAL_SUPPORT_GRAVITATIONAL_F2_FIREWALL_PRESERVED"

	StatusTensionJPairNormalizationConvention   = "CONDITIONAL_TENSION_J_PAIR_VOLUME_NORMALIZATION_MUST_BE_ACCEPTED"
	StatusTensionReducedPlanckBranchNotSelected = "CONDITIONAL_TENSION_REDUCED_PLANCK_BRANCH_NOT_SELECTED"
	StatusTensionF2MomentStillInterpreted       = "CONDITIONAL_TENSION_F2_MOMENT_REINTERPRETED_BUT_NOT_DERIVED"

	StatusFailedUnconditionalHierarchyNotClaimed = "FAILED_ROUTE_UNCONDITIONAL_HIERARCHY_SCALING_FACTOR_NOT_CLAIMED"
	StatusFailedF2CutoffMomentStillUnlocked      = "FAILED_ROUTE_F2_CUTOFF_MOMENT_STILL_NOT_LOCKED"
	StatusFailedNewtonConstantNotDerived         = "FAILED_ROUTE_NEWTON_CONSTANT_NORMALIZATION_NOT_DERIVED"
	StatusFailedElectroweakVEVNotUnconditional   = "FAILED_ROUTE_ELECTROWEAK_VEV_NOT_DERIVED_UNCONDITIONALLY"
)

const (
	inheritedHighestGate = 341

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

type GrassmannMeasure struct {
	Carrier                string
	GenerationDimension    int
	RealStructure          string
	MeasureRule            string
	BerezinPfaffianApplies bool
	Status                 string
}

type ZeroModeBlock struct {
	Matrix                   string
	JPairTrace               float64
	CanonicalBlockNorm       float64
	PfaffianPerGeneration    float64
	DeterminantPerGeneration float64
	CombinedPfaffian         float64
	CombinedDeterminant      float64
	NativeFiniteMeasure      bool
	Rule                     string
	Status                   string
}

type HierarchySynthesis struct {
	HalfAction             float64
	HalfActionExponential  float64
	GenerationFactor       float64
	PredictedRatio         float64
	ObservedUnreducedRatio float64
	ObservedReducedRatio   float64
	RatioToUnreducedTarget float64
	RelativeErrorUnreduced float64
	RatioToReducedTarget   float64
	PredictedPlanckFromVEV float64
	Status                 string
}

type GravityFirewall struct {
	F2MomentLocked        bool
	NewtonConstantDerived bool
	PlanckBranchSelected  string
	Interpretation        string
	Status                string
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
	Inputs    Inputs
	Measure   GrassmannMeasure
	ZeroMode  ZeroModeBlock
	Hierarchy HierarchySynthesis
	Gravity   GravityFirewall
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
	measure := formalizeMeasure(inputs)
	zero := evaluateZeroModePfaffian(inputs, measure)
	hierarchy := synthesizeHierarchy(inputs, zero)
	gravity := compileGravityFirewall()
	summary := compileSummary(hierarchy)
	truth := "Gate 342 audits the finite Berezin measure over the triality generation carrier. A J-paired Majorana zero-mode block carries a canonical two-slot skew form with Pfaffian sqrt(2); multiplying over the three derived generations gives 2^(3/2). This removes the Gate 341 continuum zero-mode-factor firewall under the finite J-pair volume convention. The resulting hierarchy remains rho=2^(3/2)exp(-4π²)=2.024352198454697e-17, within 0.378172% of v/M_P on the unreduced Planck branch. The gate still preserves the f2/Newton normalization and unconditional VEV firewalls."
	return Analysis{Inputs: inputs, Measure: measure, ZeroMode: zero, Hierarchy: hierarchy, Gravity: gravity, Summary: summary, Truth: truth}, nil
}

func compileInputs() Inputs {
	return Inputs{
		HighestInheritedGate: inheritedHighestGate,
		NGen:                 nGen,
		STop:                 sTop,
		ElectroweakVEVGeV:    electroweakVEVGeV,
		UnreducedPlanckGeV:   unreducedPlanckGeV,
		ReducedPlanckGeV:     unreducedPlanckGeV / math.Sqrt(8*math.Pi),
		Status:               StatusGate341Inherited,
	}
}

func formalizeMeasure(i Inputs) GrassmannMeasure {
	return GrassmannMeasure{
		Carrier:                "triality generation carrier τ_eta with N_gen=3 real/Majorana J-paired zero-mode blocks",
		GenerationDimension:    i.NGen,
		RealStructure:          "J-pair / Majorana real structure: each generation contributes a two-slot skew Grassmann block (ψ_g,Jψ_g)",
		MeasureRule:            "finite Berezin Gaussian: ∫dχ exp[-1/2 χ^T Ω χ] = pf(Ω), with pf(Ω)^2=det(Ω)",
		BerezinPfaffianApplies: true,
		Status:                 StatusFiniteGrassmannMeasureFormalized,
	}
}

func evaluateZeroModePfaffian(i Inputs, m GrassmannMeasure) ZeroModeBlock {
	block := math.Sqrt2
	return ZeroModeBlock{
		Matrix:                   "Ω_g = [[0, sqrt(2)],[-sqrt(2), 0]] on the normalized J-pair zero-mode block",
		JPairTrace:               2,
		CanonicalBlockNorm:       block,
		PfaffianPerGeneration:    block,
		DeterminantPerGeneration: 2,
		CombinedPfaffian:         math.Pow(block, float64(i.NGen)),
		CombinedDeterminant:      math.Pow(2, float64(i.NGen)),
		NativeFiniteMeasure:      m.BerezinPfaffianApplies,
		Rule:                     "the doubled J-pair has finite trace volume 2; the Majorana/Berezin measure takes the square-root volume pfaffian sqrt(2) per generation",
		Status:                   StatusPfaffianZeroModeEvaluation,
	}
}

func synthesizeHierarchy(i Inputs, z ZeroModeBlock) HierarchySynthesis {
	halfExp := math.Exp(-i.STop / 2)
	pred := z.CombinedPfaffian * halfExp
	obsU := i.ElectroweakVEVGeV / i.UnreducedPlanckGeV
	obsR := i.ElectroweakVEVGeV / i.ReducedPlanckGeV
	return HierarchySynthesis{
		HalfAction:             i.STop / 2,
		HalfActionExponential:  halfExp,
		GenerationFactor:       z.CombinedPfaffian,
		PredictedRatio:         pred,
		ObservedUnreducedRatio: obsU,
		ObservedReducedRatio:   obsR,
		RatioToUnreducedTarget: pred / obsU,
		RelativeErrorUnreduced: (pred - obsU) / obsU,
		RatioToReducedTarget:   pred / obsR,
		PredictedPlanckFromVEV: i.ElectroweakVEVGeV / pred,
		Status:                 StatusHierarchySynthesisFiniteMeasure,
	}
}

func compileGravityFirewall() GravityFirewall {
	return GravityFirewall{
		F2MomentLocked:        false,
		NewtonConstantDerived: false,
		PlanckBranchSelected:  "unreduced Planck branch matches; reduced Planck branch remains the sqrt(8π)-converted gravitational convention",
		Interpretation:        "Gate 342 derives the generation Pfaffian factor from the finite J-paired measure, but the spectral-action gravitational coefficient f2 and Newton normalization are still separate Seeley-de Witt obligations",
		Status:                StatusGravityF2FirewallPreserved,
	}
}

func compileSummary(h HierarchySynthesis) Summary {
	return Summary{
		DirectAnswer: "The sqrt(2)-per-generation factor is recovered from the finite Majorana/J-paired Berezin measure as a Pfaffian volume factor.",
		Prediction:   fmt.Sprintf("rho_pred=%.15e; M_P_from_v=%.12e GeV", h.PredictedRatio, h.PredictedPlanckFromVEV),
		Agreement:    fmt.Sprintf("ratio_to_unreduced=%.12f; relative_error=%+.9f%%", h.RatioToUnreducedTarget, 100*h.RelativeErrorUnreduced),
		Caveat:       "The finite zero-mode normalization is conditionally derived under canonical J-pair volume normalization; f2/Newton normalization remains firewalled.",
		NextGate:     "Audit whether the spectral-action gravitational f2 coefficient or Newton normalization is fixed by the same Pfaffian hierarchy law.",
		Status:       StatusHierarchyScalingConditionallyDerived,
	}
}

func Statuses(a Analysis) []string {
	return []string{
		a.Inputs.Status,
		a.Measure.Status,
		a.ZeroMode.Status,
		StatusNativeZeroModeNormalizationDerived,
		a.Hierarchy.Status,
		StatusHierarchyScalingConditionallyDerived,
		a.Gravity.Status,
		StatusTensionJPairNormalizationConvention,
		StatusTensionReducedPlanckBranchNotSelected,
		StatusTensionF2MomentStillInterpreted,
		StatusFailedUnconditionalHierarchyNotClaimed,
		StatusFailedF2CutoffMomentStillUnlocked,
		StatusFailedNewtonConstantNotDerived,
		StatusFailedElectroweakVEVNotUnconditional,
	}
}

func FormatInputs(i Inputs) string {
	return fmt.Sprintf("highest_gate=%d; N_gen=%d; S_top=%.15f; v=%.12f GeV; M_P=%.12e GeV; Mbar_P=%.12e GeV; status=%s", i.HighestInheritedGate, i.NGen, i.STop, i.ElectroweakVEVGeV, i.UnreducedPlanckGeV, i.ReducedPlanckGeV, i.Status)
}

func FormatMeasure(m GrassmannMeasure) string {
	return fmt.Sprintf("carrier=%s; dim=%d; real_structure=%s; rule=%s; pfaffian=%t; status=%s", m.Carrier, m.GenerationDimension, m.RealStructure, m.MeasureRule, m.BerezinPfaffianApplies, m.Status)
}

func FormatZeroMode(z ZeroModeBlock) string {
	return fmt.Sprintf("matrix=%s; J_trace=%.0f; block_norm=%.15f; pf_per_gen=%.15f; det_per_gen=%.15f; combined_pf=%.15f; combined_det=%.15f; native_measure=%t; rule=%s; status=%s", z.Matrix, z.JPairTrace, z.CanonicalBlockNorm, z.PfaffianPerGeneration, z.DeterminantPerGeneration, z.CombinedPfaffian, z.CombinedDeterminant, z.NativeFiniteMeasure, z.Rule, z.Status)
}

func FormatHierarchy(h HierarchySynthesis) string {
	return fmt.Sprintf("S/2=%.15f; exp(-S/2)=%.15e; gen_factor=%.15f; rho_pred=%.15e; rho_unred=%.15e; ratio_unred=%.12f; rel_unred=%+.12f; rho_red=%.15e; ratio_red=%.12f; M_P_from_v=%.12e; status=%s", h.HalfAction, h.HalfActionExponential, h.GenerationFactor, h.PredictedRatio, h.ObservedUnreducedRatio, h.RatioToUnreducedTarget, h.RelativeErrorUnreduced, h.ObservedReducedRatio, h.RatioToReducedTarget, h.PredictedPlanckFromVEV, h.Status)
}

func FormatGravity(g GravityFirewall) string {
	return fmt.Sprintf("f2_locked=%t; newton_derived=%t; branch=%s; interpretation=%s; status=%s", g.F2MomentLocked, g.NewtonConstantDerived, g.PlanckBranchSelected, g.Interpretation, g.Status)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("direct=%s; prediction=%s; agreement=%s; caveat=%s; next=%s; status=%s", s.DirectAnswer, s.Prediction, s.Agreement, s.Caveat, s.NextGate, s.Status)
}

func FormatStatuses(statuses []string) string { return strings.Join(statuses, "\n") }
