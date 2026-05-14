// Package vacuumcriticalityradiative implements Gate 350:
// Vacuum Criticality & Radiative Hierarchy Sieve.
//
// Gate 350 answers the Gate-349 challenge without fitting masses: it audits
// whether a dynamical criticality principle or a tree-level-radiative ansatz
// can reduce the remaining ASHA vacuum coordinates.  The theorem is strict:
// a boundary condition is only promoted if it follows from an installed native
// operator, not because it has useful phenomenological capacity.
package vacuumcriticalityradiative

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE350-VACUUM-CRITICALITY-RADIATIVE-HIERARCHY-SIEVE"

	StatusVacuumCriticalitySieveExecuted     = "CONDITIONAL_SUPPORT_VACUUM_CRITICALITY_SIEVE_EXECUTED"
	StatusCriticalityEquationFormalized      = "CONDITIONAL_SUPPORT_CRITICALITY_EQUATION_FORMALIZED"
	StatusCriticalTopYukawaComputed          = "CONDITIONAL_SUPPORT_CRITICAL_TOP_YUKAWA_BOUNDARY_COMPUTED"
	StatusRadiativeHierarchyAnsatzFormalized = "CONDITIONAL_SUPPORT_RADIATIVE_HIERARCHY_ANSATZ_FORMALIZED"
	StatusSMYukawaZeroFixedPointAudited      = "CONDITIONAL_SUPPORT_SM_YUKAWA_ZERO_FIXED_POINT_AUDITED"
	StatusParameterCensusUpdated             = "CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED"
	StatusNextInvariantProgramIdentified     = "CONDITIONAL_SUPPORT_MATRIX_INVARIANT_PROGRAM_IDENTIFIED"

	StatusTensionNativeBoundaryNotCritical     = "CONDITIONAL_TENSION_NATIVE_LAMBDA_BOUNDARY_NOT_MULTIPLE_POINT_CRITICAL"
	StatusTensionCriticalityRequiresNewAxiom   = "CONDITIONAL_TENSION_CRITICALITY_REQUIRES_SATURATION_AXIOM"
	StatusTensionRadiativeZeroesStayZero       = "CONDITIONAL_TENSION_STANDARD_RG_PRESERVES_TREE_LEVEL_ZERO_YUKAWAS"
	StatusTensionLightMassesNeedFlavorOperator = "CONDITIONAL_TENSION_LIGHT_MASSES_REQUIRE_EXTRA_FLAVOR_BREAKING_OPERATOR"
	StatusTensionSevenCountNotReached          = "CONDITIONAL_TENSION_SEVEN_SEAL_COUNT_NOT_REACHED"

	StatusFailedCriticalityNotDerived           = "FAILED_ROUTE_VACUUM_CRITICALITY_PRINCIPLE_NOT_DERIVED"
	StatusFailedNativeLambdaNoBetaTangency      = "FAILED_ROUTE_NATIVE_LAMBDA_BOUNDARY_HAS_NO_REAL_BETA_ZERO_TOP_SOLUTION"
	StatusFailedTopYukawaNotPredicted           = "FAILED_ROUTE_TOP_YUKAWA_NOT_PREDICTED"
	StatusFailedRadiativeMassesNotGenerated     = "FAILED_ROUTE_RADIATIVE_LIGHT_MASSES_NOT_GENERATED_BY_STANDARD_RG"
	StatusFailedLightYukawasStillVacuumCoords   = "FAILED_ROUTE_LIGHT_YUKAWAS_REMAIN_VACUUM_COORDINATES"
	StatusFailedSevenVacuumCoordinatesNotProved = "FAILED_ROUTE_SEVEN_VACUUM_COORDINATES_NOT_PROVED"
)

const (
	baselineVacuumInputs = 15

	contactShape = 1197.0 / 4624.0
	gStarSquared = 0.5

	// The one-loop beta function supplied in the project uses the Standard
	// Model hypercharge coupling g_Y, not the GUT-normalized coupling.
	// Gate 308/330 normalization gives g_*² = g₂² = (5/3) g_Y².
	hyperchargeSquared = (3.0 / 5.0) * gStarSquared
)

type Span struct {
	AuditID       string
	InheritedGate int
	AddsNewFit    bool
	Purpose       string
	Verdict       string
}

type CriticalityAudit struct {
	Formalized                bool
	NativeLambdaBoundary      float64
	G2Squared                 float64
	GYSquared                 float64
	BetaNumeratorAtNativeY0   float64
	NativeBoundaryHasBetaZero bool
	MultiplePointLambda       float64
	CriticalYukawaSquared     float64
	CriticalYukawa            float64
	CriticalYukawaOverGStar   float64
	RequiresSaturationAxiom   bool
	ReductionProved           bool
	Verdict                   string
}

type RadiativeAudit struct {
	Formalized                     bool
	TreeLevelAnsatz                string
	StandardOneLoopYukawaForm      string
	ZeroYukawaIsFixedPoint         bool
	GaugeLoopsGenerateYukawas      bool
	RequiresFlavorBreakingOperator bool
	LightMassesGenerated           bool
	CandidateReductionIfPromoted   int
	ReductionProved                bool
	Verdict                        string
}

type MatrixInvariantProgram struct {
	Identified          bool
	Reason              string
	CandidateInvariants []string
	PromotedThisGate    bool
	Verdict             string
}

type Census struct {
	StartingVacuumInputs     int
	CriticalityReduction     int
	RadiativeReduction       int
	TotalAdditionalReduction int
	RemainingVacuumInputs    int
	SevenSealTargetReached   bool
	Verdict                  string
}

type Summary struct {
	Executed              bool
	AnyReductionProved    bool
	RemainingVacuumInputs int
	OneLine               string
	Status                string
}

type Analysis struct {
	Span        Span
	Criticality CriticalityAudit
	Radiative   RadiativeAudit
	Invariants  MatrixInvariantProgram
	Census      Census
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
	span := compileSpan()
	crit := auditCriticality()
	rad := auditRadiativeHierarchy()
	inv := auditMatrixInvariantProgram()
	census := compileCensus(crit, rad)
	summary := compileSummary(census)
	truth := "Gate 350 tests two dynamical first-principle candidates.  Multiple-point criticality computes a sharp top-Yukawa target only after adding a saturation axiom, and tree-level zeroes remain zero under standard multiplicative SM Yukawa RG.  Therefore no new vacuum-coordinate reduction is promoted."
	return Analysis{Span: span, Criticality: crit, Radiative: rad, Invariants: inv, Census: census, Summary: summary, Truth: truth}, nil
}

func compileSpan() Span {
	return Span{AuditID: AuditID, InheritedGate: 349, AddsNewFit: false, Purpose: "test whether vacuum criticality or radiative hierarchy dynamics reduce the 15 ASHA vacuum coordinates", Verdict: StatusVacuumCriticalitySieveExecuted}
}

func auditCriticality() CriticalityAudit {
	nativeLambda := contactShape * gStarSquared
	betaAtY0 := betaLambdaNumerator(nativeLambda, 0, gStarSquared, hyperchargeSquared)

	// Multiple-point lane: λ=0 and βλ=0.  At λ=0 the one-loop equation is
	// 12 y_t^4 = (3/16)(2g₂⁴ + (g₂²+g_Y²)²).
	y2crit := math.Sqrt((2*sq(gStarSquared) + sq(gStarSquared+hyperchargeSquared)) / 64.0)
	ycrit := math.Sqrt(y2crit)
	y2OverG := y2crit / gStarSquared

	// Native λ boundary beta-zero check.  With the native positive λ, the
	// beta numerator is a downward quadratic in q=y².  A real root would be
	// necessary for βλ=0 at the native boundary.
	hasNativeRoot := betaZeroHasRealYukawaSolution(nativeLambda, gStarSquared, hyperchargeSquared)

	return CriticalityAudit{
		Formalized:                true,
		NativeLambdaBoundary:      nativeLambda,
		G2Squared:                 gStarSquared,
		GYSquared:                 hyperchargeSquared,
		BetaNumeratorAtNativeY0:   betaAtY0,
		NativeBoundaryHasBetaZero: hasNativeRoot,
		MultiplePointLambda:       0,
		CriticalYukawaSquared:     y2crit,
		CriticalYukawa:            ycrit,
		CriticalYukawaOverGStar:   y2OverG,
		RequiresSaturationAxiom:   true,
		ReductionProved:           false,
		Verdict: strings.Join([]string{
			StatusCriticalityEquationFormalized,
			StatusCriticalTopYukawaComputed,
			StatusTensionNativeBoundaryNotCritical,
			StatusTensionCriticalityRequiresNewAxiom,
			StatusFailedCriticalityNotDerived,
			StatusFailedNativeLambdaNoBetaTangency,
			StatusFailedTopYukawaNotPredicted,
		}, ";"),
	}
}

func auditRadiativeHierarchy() RadiativeAudit {
	return RadiativeAudit{
		Formalized:                     true,
		TreeLevelAnsatz:                "rank-one Yukawa boundary: only third generation nonzero at Λ_GUT; first and second generation singular values set to zero",
		StandardOneLoopYukawaForm:      "dY/dt = Y·F(Y†Y,g) plus matrix products with at least one Yukawa insertion",
		ZeroYukawaIsFixedPoint:         true,
		GaugeLoopsGenerateYukawas:      false,
		RequiresFlavorBreakingOperator: true,
		LightMassesGenerated:           false,
		CandidateReductionIfPromoted:   6,
		ReductionProved:                false,
		Verdict: strings.Join([]string{
			StatusRadiativeHierarchyAnsatzFormalized,
			StatusSMYukawaZeroFixedPointAudited,
			StatusTensionRadiativeZeroesStayZero,
			StatusTensionLightMassesNeedFlavorOperator,
			StatusFailedRadiativeMassesNotGenerated,
			StatusFailedLightYukawasStillVacuumCoords,
		}, ";"),
	}
}

func auditMatrixInvariantProgram() MatrixInvariantProgram {
	return MatrixInvariantProgram{
		Identified: true,
		Reason:     "Gate 349 rejected single-eigenvalue power laws; any future reduction must constrain matrix invariants, not individual fitted ratios.",
		CandidateInvariants: []string{
			"Tr(Y_f†Y_f)",
			"Tr((Y_f†Y_f)^2)",
			"det(Y_f†Y_f)",
			"discriminant/characteristic polynomial of Y_f†Y_f",
			"Koide-like root-trace functional for charged leptons",
		},
		PromotedThisGate: false,
		Verdict:          StatusNextInvariantProgramIdentified,
	}
}

func compileCensus(c CriticalityAudit, r RadiativeAudit) Census {
	critReduction := 0
	if c.ReductionProved {
		critReduction = 1
	}
	radReduction := 0
	if r.ReductionProved {
		radReduction = r.CandidateReductionIfPromoted
	}
	total := critReduction + radReduction
	remaining := baselineVacuumInputs - total
	return Census{
		StartingVacuumInputs:     baselineVacuumInputs,
		CriticalityReduction:     critReduction,
		RadiativeReduction:       radReduction,
		TotalAdditionalReduction: total,
		RemainingVacuumInputs:    remaining,
		SevenSealTargetReached:   remaining == 7,
		Verdict: strings.Join([]string{
			StatusParameterCensusUpdated,
			StatusTensionSevenCountNotReached,
			StatusFailedSevenVacuumCoordinatesNotProved,
		}, ";"),
	}
}

func compileSummary(c Census) Summary {
	reduced := c.TotalAdditionalReduction > 0
	status := StatusFailedSevenVacuumCoordinatesNotProved
	if reduced {
		status = StatusParameterCensusUpdated
	}
	return Summary{Executed: true, AnyReductionProved: reduced, RemainingVacuumInputs: c.RemainingVacuumInputs, OneLine: fmt.Sprintf("Gate 350 audits criticality and radiative hierarchy but promotes no new reduction: vacuum inputs remain %d.", c.RemainingVacuumInputs), Status: status}
}

func sq(x float64) float64 { return x * x }

// betaLambdaNumerator returns the one-loop bracket before division by 16π²,
// using q = y_t² and the beta function supplied in the Gate-309 lineage.
func betaLambdaNumerator(lambda, q, g2sq, gYsq float64) float64 {
	return 24*lambda*lambda + 12*lambda*q - 12*q*q + (3.0/16.0)*(2*sq(g2sq)+sq(g2sq+gYsq)) - lambda*(9*g2sq+3*gYsq)
}

func betaZeroHasRealYukawaSolution(lambda, g2sq, gYsq float64) bool {
	// β = -12q² + 12λq + c.  Real q roots require discriminant ≥ 0.
	c := 24*lambda*lambda + (3.0/16.0)*(2*sq(g2sq)+sq(g2sq+gYsq)) - lambda*(9*g2sq+3*gYsq)
	disc := sq(12*lambda) + 48*c
	if disc < 0 {
		return false
	}
	q1 := (12*lambda + math.Sqrt(disc)) / 24.0
	q2 := (12*lambda - math.Sqrt(disc)) / 24.0
	return q1 >= 0 || q2 >= 0
}

func Statuses(a Analysis) []string {
	return []string{
		StatusVacuumCriticalitySieveExecuted,
		StatusCriticalityEquationFormalized,
		StatusCriticalTopYukawaComputed,
		StatusRadiativeHierarchyAnsatzFormalized,
		StatusSMYukawaZeroFixedPointAudited,
		StatusParameterCensusUpdated,
		StatusNextInvariantProgramIdentified,
		StatusTensionNativeBoundaryNotCritical,
		StatusTensionCriticalityRequiresNewAxiom,
		StatusTensionRadiativeZeroesStayZero,
		StatusTensionLightMassesNeedFlavorOperator,
		StatusTensionSevenCountNotReached,
		StatusFailedCriticalityNotDerived,
		StatusFailedNativeLambdaNoBetaTangency,
		StatusFailedTopYukawaNotPredicted,
		StatusFailedRadiativeMassesNotGenerated,
		StatusFailedLightYukawasStillVacuumCoords,
		StatusFailedSevenVacuumCoordinatesNotProved,
	}
}

func FormatSpan(s Span) string {
	return fmt.Sprintf("%s | inherited gate=%d | adds_new_fit=%t | %s", s.AuditID, s.InheritedGate, s.AddsNewFit, s.Purpose)
}

func FormatCriticality(c CriticalityAudit) string {
	return fmt.Sprintf("λ_native=%.15f; g₂²=%.12f; gY²=%.12f; β_native(y=0)=%.12f; native_beta_zero=%t; MPP λ=0 => y_t²=%.12f, y_t=%.12f, y_t²/g_*²=%.12f; saturation_axiom=%t", c.NativeLambdaBoundary, c.G2Squared, c.GYSquared, c.BetaNumeratorAtNativeY0, c.NativeBoundaryHasBetaZero, c.CriticalYukawaSquared, c.CriticalYukawa, c.CriticalYukawaOverGStar, c.RequiresSaturationAxiom)
}

func FormatRadiative(r RadiativeAudit) string {
	return fmt.Sprintf("%s; one_loop_form=%s; zero_fixed_point=%t; gauge_generates_yukawa=%t; flavor_operator_required=%t; reduction_proved=%t", r.TreeLevelAnsatz, r.StandardOneLoopYukawaForm, r.ZeroYukawaIsFixedPoint, r.GaugeLoopsGenerateYukawas, r.RequiresFlavorBreakingOperator, r.ReductionProved)
}

func FormatInvariants(i MatrixInvariantProgram) string {
	return fmt.Sprintf("identified=%t; promoted=%t; candidates=%s; reason=%s", i.Identified, i.PromotedThisGate, strings.Join(i.CandidateInvariants, ", "), i.Reason)
}

func FormatCensus(c Census) string {
	return fmt.Sprintf("starting=%d; criticality_reduction=%d; radiative_reduction=%d; total_reduction=%d; remaining=%d; seven_target=%t", c.StartingVacuumInputs, c.CriticalityReduction, c.RadiativeReduction, c.TotalAdditionalReduction, c.RemainingVacuumInputs, c.SevenSealTargetReached)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("executed=%t; any_reduction=%t; remaining=%d; %s", s.Executed, s.AnyReductionProved, s.RemainingVacuumInputs, s.OneLine)
}
