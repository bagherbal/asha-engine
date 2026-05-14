// Package crosssectorreductionaudit implements Gate 349:
// Cross-Sector Reduction Audit / Vacuum Parameter Compression Sieve.
//
// Gate 349 does not attempt to fit fermion masses.  It asks whether the
// ASHA landscape relations found through Gate 348 reduce the empirical vacuum
// coordinates further by cross-sector constraints: seesaw reduction, vacuum
// stability saturation, and B-gap / triality mass-ratio laws.
package crosssectorreductionaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE349-CROSS-SECTOR-REDUCTION-AUDIT-VACUUM-PARAMETER-COMPRESSION-SIEVE"

	StatusCrossSectorReductionAuditExecuted = "CONDITIONAL_SUPPORT_CROSS_SECTOR_REDUCTION_AUDIT_EXECUTED"
	StatusSeesawDependencyFormalized        = "CONDITIONAL_SUPPORT_SEESAW_DEPENDENCY_FORMALIZED"
	StatusStabilityBoundFormalized          = "CONDITIONAL_SUPPORT_VACUUM_STABILITY_BOUND_FORMALIZED"
	StatusBGapPowerLawTested                = "CONDITIONAL_SUPPORT_BGAP_POWER_LAW_TESTED"
	StatusParameterCensusUpdated            = "CONDITIONAL_SUPPORT_PARAMETER_CENSUS_UPDATED"
	StatusReductionTargetsCataloged         = "CONDITIONAL_SUPPORT_REDUCTION_TARGETS_CATALOGED"

	StatusTensionSeesawNeedsDiracTexture  = "CONDITIONAL_TENSION_SEESAW_NEEDS_DIRAC_TEXTURE"
	StatusTensionStabilityBoundNotUnique  = "CONDITIONAL_TENSION_STABILITY_BOUND_NOT_UNIQUE"
	StatusTensionBGapPowerLawNotUniversal = "CONDITIONAL_TENSION_BGAP_POWER_LAW_NOT_UNIVERSAL"
	StatusTensionSevenSealCountNotReached = "CONDITIONAL_TENSION_SEVEN_SEAL_COUNT_NOT_REACHED"

	StatusFailedNoParameterReductionProved      = "FAILED_ROUTE_NO_ADDITIONAL_PARAMETER_REDUCTION_PROVED"
	StatusFailedNeutrinoMassesNotDerived        = "FAILED_ROUTE_NEUTRINO_MASS_RATIOS_NOT_DERIVED"
	StatusFailedTopYukawaNotPredicted           = "FAILED_ROUTE_TOP_YUKAWA_NOT_PREDICTED_BY_STABILITY"
	StatusFailedMassPowerLawNotDerived          = "FAILED_ROUTE_BGAP_MASS_POWER_LAW_NOT_DERIVED"
	StatusFailedSevenVacuumCoordinatesNotProved = "FAILED_ROUTE_SEVEN_VACUUM_COORDINATES_NOT_PROVED"
)

const (
	baselineSMParameters      = 19
	derivedBoundaryParameters = 4
	baselineVacuumInputs      = baselineSMParameters - derivedBoundaryParameters

	bGap  = 0.102464921191
	rPlus = 1.64547046301119

	// Quarantined reference values used only to test whether the proposed
	// relations have numerical capacity.  They are not imported as derivations.
	deltaM21EV2 = 7.49e-5
	deltaM31EV2 = 2.513e-3

	electronMeV = 0.510998950
	muonMeV     = 105.6583755
	tauMeV      = 1776.86
)

type AuditSpan struct {
	AuditID       string
	InheritedGate int
	AddsNewFit    bool
	Purpose       string
	Verdict       string
}

type SeesawAudit struct {
	Formalized                       bool
	MajoranaScaleSymbol              string
	Formula                          string
	CommonMajoranaScaleCancelsRatios bool
	RequiresDiracSingularValues      bool
	DeltaM21EV2                      float64
	DeltaM31EV2                      float64
	ObservedDeltaRatio               float64
	SimpleBGapExponent               float64
	RatioPredicted                   bool
	ReductionProved                  bool
	Verdict                          string
}

type StabilityAudit struct {
	Formalized              bool
	Boundary                string
	StabilityCondition      string
	DependsOnTopYukawa      bool
	BoundIsInequality       bool
	RequiresSaturationAxiom bool
	PredictsTopMass         bool
	ReductionProved         bool
	Verdict                 string
}

type PowerLawDatum struct {
	Name                          string
	ObservedRatio                 float64
	BGapExponent                  float64
	NearestHalfInteger            float64
	NearestInteger                float64
	RelativeErrorToNearestHalf    float64
	RelativeErrorToNearestInteger float64
	SimpleLawAccepted             bool
}

type PowerLawAudit struct {
	Formalized                     bool
	Law                            string
	Data                           []PowerLawDatum
	MaxRelativeErrorNearestHalf    float64
	MaxRelativeErrorNearestInteger float64
	UniversalSimplePowerLawFound   bool
	ReductionProved                bool
	Verdict                        string
}

type CensusAudit struct {
	BaselineSMParameters       int
	DerivedBoundaryConstraints int
	StartingVacuumInputs       int
	SeesawReduction            int
	StabilityReduction         int
	PowerLawReduction          int
	TotalAdditionalReduction   int
	RemainingVacuumInputs      int
	SevenSealTargetReached     bool
	Verdict                    string
}

type Summary struct {
	Executed              bool
	AnyReductionProved    bool
	RemainingVacuumInputs int
	OneLine               string
	Status                string
}

type Analysis struct {
	Span      AuditSpan
	Seesaw    SeesawAudit
	Stability StabilityAudit
	PowerLaw  PowerLawAudit
	Census    CensusAudit
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
	seesaw := auditSeesaw()
	stability := auditStability()
	power := auditPowerLaw()
	census := compileCensus(seesaw, stability, power)
	summary := compileSummary(census)
	truth := "Gate 349 tests whether ASHA's cross-sector structures reduce the remaining vacuum coordinates.  The seesaw, stability, and B-gap power-law lanes have strong diagnostic value, but none becomes a theorem without an additional Dirac texture, saturation axiom, or native mass-ratio operator."
	return Analysis{Span: span, Seesaw: seesaw, Stability: stability, PowerLaw: power, Census: census, Summary: summary, Truth: truth}, nil
}

func compileSpan() AuditSpan {
	return AuditSpan{AuditID: AuditID, InheritedGate: 348, AddsNewFit: false, Purpose: "test cross-sector reductions of the Gate-345 minimal 15 vacuum coordinates", Verdict: StatusCrossSectorReductionAuditExecuted}
}

func auditSeesaw() SeesawAudit {
	observed := deltaM31EV2 / deltaM21EV2
	exponent := math.Log(observed) / math.Log(1.0/bGap)
	return SeesawAudit{
		Formalized:                       true,
		MajoranaScaleSymbol:              "M_R ∝ B_gap · M_* or intermediate threshold scale",
		Formula:                          "m_ν,i ≈ m_D,i² / M_R",
		CommonMajoranaScaleCancelsRatios: true,
		RequiresDiracSingularValues:      true,
		DeltaM21EV2:                      deltaM21EV2,
		DeltaM31EV2:                      deltaM31EV2,
		ObservedDeltaRatio:               observed,
		SimpleBGapExponent:               exponent,
		RatioPredicted:                   false,
		ReductionProved:                  false,
		Verdict:                          strings.Join([]string{StatusSeesawDependencyFormalized, StatusTensionSeesawNeedsDiracTexture, StatusFailedNeutrinoMassesNotDerived}, ";"),
	}
}

func auditStability() StabilityAudit {
	return StabilityAudit{
		Formalized:              true,
		Boundary:                "λ_H/g_*² = 1197/4624 with α_GUT⁻¹ = 8π branch",
		StabilityCondition:      "λ(μ) ≥ 0 for v ≤ μ ≤ M_P",
		DependsOnTopYukawa:      true,
		BoundIsInequality:       true,
		RequiresSaturationAxiom: true,
		PredictsTopMass:         false,
		ReductionProved:         false,
		Verdict:                 strings.Join([]string{StatusStabilityBoundFormalized, StatusTensionStabilityBoundNotUnique, StatusFailedTopYukawaNotPredicted}, ";"),
	}
}

func auditPowerLaw() PowerLawAudit {
	data := []PowerLawDatum{
		evaluatePowerLaw("muon/electron charged-lepton ratio", muonMeV/electronMeV),
		evaluatePowerLaw("tau/muon charged-lepton ratio", tauMeV/muonMeV),
		evaluatePowerLaw("atmospheric/solar neutrino Δm² ratio", deltaM31EV2/deltaM21EV2),
		evaluatePowerLaw("r_plus amplitude ratio", rPlus),
	}
	maxHalf := 0.0
	maxInt := 0.0
	accepted := true
	for _, d := range data {
		maxHalf = math.Max(maxHalf, math.Abs(d.RelativeErrorToNearestHalf))
		maxInt = math.Max(maxInt, math.Abs(d.RelativeErrorToNearestInteger))
		accepted = accepted && d.SimpleLawAccepted
	}
	return PowerLawAudit{Formalized: true, Law: "ratio ?= B_gap^{-n} for canonical integer or half-integer n", Data: data, MaxRelativeErrorNearestHalf: maxHalf, MaxRelativeErrorNearestInteger: maxInt, UniversalSimplePowerLawFound: accepted, ReductionProved: false, Verdict: strings.Join([]string{StatusBGapPowerLawTested, StatusTensionBGapPowerLawNotUniversal, StatusFailedMassPowerLawNotDerived}, ";")}
}

func evaluatePowerLaw(name string, observed float64) PowerLawDatum {
	exponent := math.Log(observed) / math.Log(1.0/bGap)
	nearestHalf := math.Round(exponent*2.0) / 2.0
	nearestInt := math.Round(exponent)
	halfPred := math.Pow(1.0/bGap, nearestHalf)
	intPred := math.Pow(1.0/bGap, nearestInt)
	halfErr := (halfPred - observed) / observed
	intErr := (intPred - observed) / observed
	return PowerLawDatum{Name: name, ObservedRatio: observed, BGapExponent: exponent, NearestHalfInteger: nearestHalf, NearestInteger: nearestInt, RelativeErrorToNearestHalf: halfErr, RelativeErrorToNearestInteger: intErr, SimpleLawAccepted: math.Abs(halfErr) < 0.05 || math.Abs(intErr) < 0.05}
}

func compileCensus(seesaw SeesawAudit, stability StabilityAudit, power PowerLawAudit) CensusAudit {
	seesawReduction := 0
	if seesaw.ReductionProved {
		seesawReduction = 3
	}
	stabilityReduction := 0
	if stability.ReductionProved {
		stabilityReduction = 1
	}
	powerReduction := 0
	if power.ReductionProved {
		powerReduction = 4
	}
	total := seesawReduction + stabilityReduction + powerReduction
	remaining := baselineVacuumInputs - total
	return CensusAudit{BaselineSMParameters: baselineSMParameters, DerivedBoundaryConstraints: derivedBoundaryParameters, StartingVacuumInputs: baselineVacuumInputs, SeesawReduction: seesawReduction, StabilityReduction: stabilityReduction, PowerLawReduction: powerReduction, TotalAdditionalReduction: total, RemainingVacuumInputs: remaining, SevenSealTargetReached: remaining == 7, Verdict: strings.Join([]string{StatusParameterCensusUpdated, StatusReductionTargetsCataloged, StatusTensionSevenSealCountNotReached, StatusFailedNoParameterReductionProved, StatusFailedSevenVacuumCoordinatesNotProved}, ";")}
}

func compileSummary(c CensusAudit) Summary {
	reduced := c.TotalAdditionalReduction > 0
	status := StatusFailedNoParameterReductionProved
	if reduced {
		status = StatusParameterCensusUpdated
	}
	line := fmt.Sprintf("Cross-sector relations were audited but not promoted: minimal vacuum inputs remain %d, not 7.", c.RemainingVacuumInputs)
	return Summary{Executed: true, AnyReductionProved: reduced, RemainingVacuumInputs: c.RemainingVacuumInputs, OneLine: line, Status: status}
}

func Statuses(a Analysis) []string {
	return []string{
		StatusCrossSectorReductionAuditExecuted,
		StatusSeesawDependencyFormalized,
		StatusStabilityBoundFormalized,
		StatusBGapPowerLawTested,
		StatusParameterCensusUpdated,
		StatusReductionTargetsCataloged,
		StatusTensionSeesawNeedsDiracTexture,
		StatusTensionStabilityBoundNotUnique,
		StatusTensionBGapPowerLawNotUniversal,
		StatusTensionSevenSealCountNotReached,
		StatusFailedNoParameterReductionProved,
		StatusFailedNeutrinoMassesNotDerived,
		StatusFailedTopYukawaNotPredicted,
		StatusFailedMassPowerLawNotDerived,
		StatusFailedSevenVacuumCoordinatesNotProved,
	}
}

func FormatSpan(s AuditSpan) string {
	return fmt.Sprintf("%s | inherited gate=%d | adds_new_fit=%t | %s", s.AuditID, s.InheritedGate, s.AddsNewFit, s.Purpose)
}

func FormatSeesaw(s SeesawAudit) string {
	return fmt.Sprintf("%s; observed Δm31²/Δm21²=%.9g; equivalent B_gap exponent=%.6f; ratio_predicted=%t", s.Formula, s.ObservedDeltaRatio, s.SimpleBGapExponent, s.RatioPredicted)
}

func FormatStability(s StabilityAudit) string {
	return fmt.Sprintf("%s; condition=%s; inequality=%t; saturation_axiom_required=%t; predicts_top=%t", s.Boundary, s.StabilityCondition, s.BoundIsInequality, s.RequiresSaturationAxiom, s.PredictsTopMass)
}

func FormatPowerLaw(p PowerLawAudit) string {
	parts := []string{p.Law}
	for _, d := range p.Data {
		parts = append(parts, fmt.Sprintf("%s ratio=%.9g exponent=%.6f nearest_half=%.1f err_half=%+.3f nearest_int=%.0f err_int=%+.3f", d.Name, d.ObservedRatio, d.BGapExponent, d.NearestHalfInteger, d.RelativeErrorToNearestHalf, d.NearestInteger, d.RelativeErrorToNearestInteger))
	}
	return strings.Join(parts, " | ")
}

func FormatCensus(c CensusAudit) string {
	return fmt.Sprintf("SM=%d, native_constraints=%d, starting_vacuum=%d, extra_reduction=%d, remaining=%d, seven_target=%t", c.BaselineSMParameters, c.DerivedBoundaryConstraints, c.StartingVacuumInputs, c.TotalAdditionalReduction, c.RemainingVacuumInputs, c.SevenSealTargetReached)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("executed=%t; any_reduction=%t; remaining=%d; %s", s.Executed, s.AnyReductionProved, s.RemainingVacuumInputs, s.OneLine)
}
