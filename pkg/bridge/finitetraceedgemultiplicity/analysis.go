// Package finitetraceedgemultiplicity implements Gate 382:
// Finite Trace Edge Multiplicity / Effective Coefficient Sieve.
//
// Gate 381 proved that the finite Dirac graph has ten J-doubled edge slots,
// but correctly rejected the type error f0_CCM = Tr_E(P_edge). Gate 382 asks
// the repaired question: can f0 be kept as the normalized sharp-cutoff value
// f0=1 while the factor 10 enters lawfully through the finite Hilbert-space
// trace channel that controls Higgs kinetic/quartic normalization?
package finitetraceedgemultiplicity

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/ccmpfaffianf0closure"
	"github.com/bagherbal/asha-engine/pkg/bridge/spectralgraphf0index"
)

const (
	AuditID = "GATE382-FINITE-TRACE-EDGE-MULTIPLICITY-EFFECTIVE-COEFFICIENT-SIEVE"

	StatusMomentLockedToOne               = "CONDITIONAL_SUPPORT_CCM_MOMENT_F0_LOCKED_TO_UNIT_SHARP_CUTOFF"
	StatusFiniteTraceDecompositionAudited = "CONDITIONAL_SUPPORT_FINITE_TRACE_DECOMPOSITION_AUDITED"
	StatusEdgeMultiplicityTenInherited    = "CONDITIONAL_SUPPORT_J_DOUBLED_EDGE_MULTIPLICITY_TEN_INHERITED"
	StatusCoefficientLanesComputed        = "CONDITIONAL_SUPPORT_EFFECTIVE_COEFFICIENT_LANES_COMPUTED"
	StatusTenDenominatorLaneNearCloses    = "CONDITIONAL_SUPPORT_TEN_DENOMINATOR_LANE_REPRODUCES_NEAR_HIGGS_CLOSURE"
	StatusTenOverSevenGapIsolated         = "CONDITIONAL_SUPPORT_TEN_OVER_SEVEN_GAP_ISOLATED"
	StatusHiggsNearClosureInherited       = "CONDITIONAL_SUPPORT_HIGGS_NEAR_CLOSURE_INHERITED_FROM_GATE380"

	StatusTensionFiniteRatioAlreadyFullTrace         = "CONDITIONAL_TENSION_E_OVER_A2_ALREADY_APPEARS_AS_FULL_FINITE_TRACE_RATIO"
	StatusTensionMultiplyingTraceDoubleCounts        = "CONDITIONAL_TENSION_EDGE_MULTIPLICITY_IN_NUMERATOR_WOULD_DOUBLE_COUNT"
	StatusTensionDenominatorMultiplicityNeedsTheorem = "CONDITIONAL_TENSION_DENOMINATOR_EDGE_MULTIPLICITY_REQUIRES_KINETIC_TRACE_THEOREM"
	StatusTensionF0OneAloneOverpredictsHiggs         = "CONDITIONAL_TENSION_F0_ONE_ALONE_OVERPREDICTS_HIGGS_MASS"
	StatusTensionTenOverSevenOriginOpen              = "CONDITIONAL_TENSION_TEN_OVER_SEVEN_NORMALIZATION_ORIGIN_OPEN"

	StatusFailedEdgeMultiplicityNotExtractedAsNativeCoefficient = "FAILED_ROUTE_EDGE_MULTIPLICITY_NOT_EXTRACTED_AS_NATIVE_CCM_COEFFICIENT"
	StatusFailedF0MomentStillNotEdgeMultiplicity                = "FAILED_ROUTE_F0_MOMENT_STILL_NOT_EDGE_MULTIPLICITY"
	StatusFailedHiggsMassNotGeometricallySealed                 = "FAILED_ROUTE_HIGGS_MASS_NOT_GEOMETRICALLY_SEALED"
	StatusFailedTenOverSevenNotDerived                          = "FAILED_ROUTE_TEN_OVER_SEVEN_NOT_DERIVED"
	StatusFailedFullNumericalTOEClosureOpen                     = "FAILED_ROUTE_FULL_NUMERICAL_TOE_CLOSURE_STILL_NOT_REACHED"
)

const (
	TraceRatioEOverA2 = 1197.0 / 4624.0
	HiggsBoundaryGeV  = 125.10
	StandardVEVGeV    = 246.22
	UnitF0            = 1.0
)

type MomentNormalization struct {
	CCMF0Definition string
	LockedValue     float64
	Rationale       string
	Verdict         string
}

type EdgeMultiplicity struct {
	FundamentalEdges     []string
	FundamentalEdgeCount int
	JDoubledEdgeCount    int
	ProjectionTrace      float64
	InheritedFromGate381 bool
	CanReplaceF0         bool
	Verdict              string
}

type TraceDecomposition struct {
	FiniteRatioLabel         string
	FiniteRatioValue         float64
	IsAlreadyTraceRatio      bool
	GenericDecomposition     string
	UniformMultiplicityLaw   string
	CanPullExtraTenFromRatio bool
	Verdict                  string
}

type CoefficientLane struct {
	Name             string
	Formula          string
	F0               float64
	EffectiveDenom   float64
	RatioMultiplier  float64
	LambdaH          float64
	MassStandardGeV  float64
	MassPfaffianGeV  float64
	ErrorPfaffianGeV float64
	Native           bool
	CircularRisk     bool
	Verdict          string
}

type GapAudit struct {
	ContactF0         float64
	EdgeDenominator   float64
	RatioTenOverSeven float64
	TargetF0Standard  float64
	TargetF0Pfaffian  float64
	RecognizedNative  bool
	Verdict           string
}

type Calculation struct {
	Executed                  bool
	Moment                    MomentNormalization
	EdgeMultiplicity          EdgeMultiplicity
	TraceDecomposition        TraceDecomposition
	Lanes                     []CoefficientLane
	Gap                       GapAudit
	Statuses                  []string
	EdgeMultiplicityExtracted bool
	HiggsMassSealed           bool
	FullNumericalTOEClosure   bool
	Truth                     string
}

type Analysis struct{ Calculation Calculation }

var defaultOnce sync.Once
var defaultA Analysis
var defaultErr error

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	gate381, err := spectralgraphf0index.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("could not inherit Gate 381 edge projection: %w", err)
	}
	gate380 := ccmpfaffianf0closure.NativeConstants()
	if len(gate380) == 0 {
		return Analysis{}, fmt.Errorf("could not inherit Gate 380 Higgs closure constants")
	}

	pfVEV := gate380["pfaffian_vev_gev"]
	if pfVEV == 0 {
		pfVEV = ccmpfaffianf0closure.UnreducedPlanckGeV * math.Pow(2, 1.5) * math.Exp(-4*math.Pi*math.Pi)
	}
	edgeCount := gate381.Calculation.EdgeProjection.JDoubledEdgeSlotCount
	edges := append([]string(nil), gate381.Calculation.EdgeProjection.FundamentalEdges...)

	moment := MomentNormalization{
		CCMF0Definition: "CCM f₀ is the zeroth value/moment of the test function; for a unit sharp cutoff f(0)=1.",
		LockedValue:     UnitF0,
		Rationale:       "Gate 382 repairs the Gate 381 type mismatch by not setting f₀ equal to the edge count. The continuous moment is locked to the normalized cutoff value.",
		Verdict:         "The continuous part is clean: f₀=1. Any factor 10 must enter through finite trace normalization, not by redefining f₀.",
	}

	edge := EdgeMultiplicity{
		FundamentalEdges:     edges,
		FundamentalEdgeCount: len(edges),
		JDoubledEdgeCount:    edgeCount,
		ProjectionTrace:      float64(edgeCount),
		InheritedFromGate381: true,
		CanReplaceF0:         false,
		Verdict:              "The finite graph still natively contains five fundamental edge classes and ten J-doubled edge slots. This is a real finite multiplicity; it still cannot replace the scalar CCM moment f₀.",
	}

	trace := TraceDecomposition{
		FiniteRatioLabel:         "e/a²",
		FiniteRatioValue:         TraceRatioEOverA2,
		IsAlreadyTraceRatio:      true,
		GenericDecomposition:     "a=Tr(Y†Y), e=Tr((Y†Y)²). If N identical orthogonal channels exist, a=N·a₀ and e=N·e₀, so e/a²=(e₀/a₀²)/N.",
		UniformMultiplicityLaw:   "An edge multiplicity usually enters the canonical quartic through the finite trace ratio itself, not as an extra post-hoc multiplier. Adding N again risks double-counting unless e/a² was computed per edge rather than over the full finite representation.",
		CanPullExtraTenFromRatio: false,
		Verdict:                  "The ASHA value 1197/4624 is already recorded as the finite trace ratio e/a². Gate 382 finds no native theorem saying this ratio is a per-edge quantity that must be divided by, or multiplied by, an additional J-doubled edge count.",
	}

	lanes := []CoefficientLane{
		lane("unit f₀, no extra multiplicity", "λ=π²(e/a²)/(2·1)", 1, 1, 1, StandardVEVGeV, pfVEV, true, false, "With f₀=1 and the recorded full finite trace ratio, the Higgs mass is far too high. The unit sharp cutoff alone does not close the Higgs mass."),
		lane("wrong numerator multiplication", "λ=π²(10·e/a²)/(2·1)", 1, 1, 10, StandardVEVGeV, pfVEV, false, true, "Putting the edge count into the quartic numerator increases λ by 10 and badly overpredicts the Higgs mass. This is the wrong channel."),
		lane("denominator edge normalization witness", "λ=π²(e/a²)/(2·10)", 1, 10, 1, StandardVEVGeV, pfVEV, false, true, "If the edge count enters exactly as a kinetic/canonical-normalization denominator, this reproduces the Gate 380 f₀_eff≈10 near-closure. But that is precisely the missing theorem; it is not derived by finite trace decomposition alone."),
		lane("contact f₀=7 ledger", "λ=π²(e/a²)/(2·7)", 7, 7, 1, StandardVEVGeV, pfVEV, true, false, "The old contact f₀=7 ledger overpredicts the Higgs mass relative to the observed boundary."),
	}

	gap := GapAudit{
		ContactF0:         7,
		EdgeDenominator:   float64(edgeCount),
		RatioTenOverSeven: float64(edgeCount) / 7.0,
		TargetF0Standard:  gate380["f0_eff_standard_ew_vev"],
		TargetF0Pfaffian:  gate380["f0_eff_pfaffian_unreduced_planck"],
		RecognizedNative:  false,
		Verdict:           "The remaining 10/7 ratio is now isolated as a normalization mismatch between the contact-spectrum f₀=7 ledger and the edge-denominator value 10 that yields near-Higgs closure. Gate 382 does not derive this ratio from J-reality, moment-slot conversion, or trace decomposition; it remains the next exact gap.",
	}

	statuses := []string{
		StatusMomentLockedToOne,
		StatusFiniteTraceDecompositionAudited,
		StatusEdgeMultiplicityTenInherited,
		StatusCoefficientLanesComputed,
		StatusTenDenominatorLaneNearCloses,
		StatusTenOverSevenGapIsolated,
		StatusHiggsNearClosureInherited,
		StatusTensionFiniteRatioAlreadyFullTrace,
		StatusTensionMultiplyingTraceDoubleCounts,
		StatusTensionDenominatorMultiplicityNeedsTheorem,
		StatusTensionF0OneAloneOverpredictsHiggs,
		StatusTensionTenOverSevenOriginOpen,
		StatusFailedEdgeMultiplicityNotExtractedAsNativeCoefficient,
		StatusFailedF0MomentStillNotEdgeMultiplicity,
		StatusFailedHiggsMassNotGeometricallySealed,
		StatusFailedTenOverSevenNotDerived,
		StatusFailedFullNumericalTOEClosureOpen,
	}

	truth := "Gate 382 repairs the Gate 381 type mismatch by locking the continuous CCM moment to f₀=1 and asking whether the finite trace itself supplies the effective factor 10. The finite ASHA graph does contain ten J-doubled edge slots, but the recorded Higgs trace ratio e/a²=1197/4624 is already a finite trace ratio. Multiplying it by 10 is the wrong channel and greatly overpredicts the Higgs mass; putting 10 in the denominator reproduces the Gate 380 near-closure, but that requires a new kinetic/canonical-normalization theorem rather than following automatically from Tr_{H_F}(D_F⁴). Therefore Gate 382 isolates the exact remaining gap as the 10/7 normalization between the contact f₀=7 ledger and the edge-denominator value 10. Higgs mass closure remains a powerful near-closure, not a sealed theorem."

	return Analysis{Calculation: Calculation{
		Executed:                  true,
		Moment:                    moment,
		EdgeMultiplicity:          edge,
		TraceDecomposition:        trace,
		Lanes:                     lanes,
		Gap:                       gap,
		Statuses:                  statuses,
		EdgeMultiplicityExtracted: false,
		HiggsMassSealed:           false,
		FullNumericalTOEClosure:   false,
		Truth:                     truth,
	}}, nil
}

func lane(name, formula string, f0, effectiveDenom, ratioMultiplier, vStd, vPf float64, native, circularRisk bool, verdict string) CoefficientLane {
	lambda := math.Pi * math.Pi * (TraceRatioEOverA2 * ratioMultiplier) / (2.0 * effectiveDenom)
	// If the lane is the contact f0 lane, f0 itself is the denominator.
	if name == "contact f₀=7 ledger" {
		lambda = math.Pi * math.Pi * TraceRatioEOverA2 / (2.0 * f0)
	}
	stdMass := vStd * math.Sqrt(2.0*lambda)
	pfMass := vPf * math.Sqrt(2.0*lambda)
	return CoefficientLane{
		Name:             name,
		Formula:          formula,
		F0:               f0,
		EffectiveDenom:   effectiveDenom,
		RatioMultiplier:  ratioMultiplier,
		LambdaH:          lambda,
		MassStandardGeV:  stdMass,
		MassPfaffianGeV:  pfMass,
		ErrorPfaffianGeV: pfMass - HiggsBoundaryGeV,
		Native:           native,
		CircularRisk:     circularRisk,
		Verdict:          verdict,
	}
}

func StatusLine(c Calculation) string { return strings.Join(c.Statuses, "\n") }

func NativeConstants() map[string]float64 {
	a, err := BuildDefault()
	if err != nil {
		return map[string]float64{}
	}
	c := a.Calculation
	m := map[string]float64{
		"unit_ccm_f0":               c.Moment.LockedValue,
		"j_doubled_edge_count":      float64(c.EdgeMultiplicity.JDoubledEdgeCount),
		"edge_projection_trace":     c.EdgeMultiplicity.ProjectionTrace,
		"trace_ratio_e_over_a2":     c.TraceDecomposition.FiniteRatioValue,
		"ten_over_seven_gap":        c.Gap.RatioTenOverSeven,
		"target_f0_standard_ew_vev": c.Gap.TargetF0Standard,
		"target_f0_pfaffian_vev":    c.Gap.TargetF0Pfaffian,
	}
	for _, l := range c.Lanes {
		key := strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(l.Name, " ", "_"), "₀", "0"))
		m[key+"_lambda"] = l.LambdaH
		m[key+"_mh_pfaffian"] = l.MassPfaffianGeV
	}
	return m
}
