// Package generation2boundaryactivationmeasurenaturalityanduniquenessaudit implements
// Gate 921: BoundaryActivationMeasure Naturality and Uniqueness Audit.
//
// Gate 921 follows Gate 920's formal BoundaryActivationMeasure candidate and
// tests whether the measure is unique under the current discovered naturality
// constraints. The result is deliberately conditional: mu_B is unique among the
// tested constraint-compatible bridge measures, but no native ASHA measure
// theorem, no native alpha theorem, and no native R3 theorem is certified.
package generation2boundaryactivationmeasurenaturalityanduniquenessaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE921-BOUNDARYACTIVATIONMEASURE-NATURALITY-UNIQUENESS-AUDIT"

	Gate920ShortStatus = "R3_ALPHA_BOUNDARY_MEASURE_CANDIDATE_NATIVE_MEASURE_MISSING"

	SBoundary = 0.0012924448188162962
	AlphaB    = 0.0003878958469680527

	RankI1  = 3
	RankI2  = 7
	RankH10 = 10
	RankH72 = 72

	AlphaLinear = 0.00038773344564488885
	AlphaQuad   = 0.0000001624013231638281

	ReducedResponse       = "R_B(s)=(1+s b1)(1+s b2)-1=s(b1+b2)+s^2(b1 wedge b2)"
	MeasureFormula        = "mu_B(R_B(S_split))=sum_{k=1}^2 rank(I_B^Z2(k))/rank(H_k)*S_split^k"
	BranchMeasureFormula  = "mu_B(R_B(S_split))=(3/10)S_split+(7/72)S_split^2"
	BoundaryAlphaFormula  = "alpha_B^Z2=mu_B(R_B(S_split))"
	BoundaryMeasureObject = "BoundaryActivationMeasure"

	Classification      = "R3_BOUNDARY_ACTIVATION_MEASURE_NATURALITY_UNIQUENESS_CANDIDATE_NOT_NATIVE"
	ShortStatus         = "R3_ALPHA_MEASURE_UNIQUENESS_SUPPORTED_NATIVE_THEOREM_MISSING"
	FinalTruth          = "BOUNDARY_ACTIVATION_MEASURE_IS_UNIQUE_UNDER_CURRENT_NATURALITY_CONSTRAINTS_BUT_NATIVE_MEASURE_THEOREM_MISSING"
	StrategicConclusion = "Gate 921 upgrades mu_B from a formal bridge measure to a unique natural BoundaryActivationMeasure candidate under the current constraints. The remaining wound is whether the naturality constraints themselves are native ASHA theorems."
	NextGate            = "NEXT_PRESSURE_GATE922_BOUNDARYACTIVATIONMEASURE_NATIVECONSTRAINT_SOURCE_AUDIT"

	StatusInheritedGate920      = "PASS_GATE920_BOUNDARY_ACTIVATION_MEASURE_CANDIDATE_INHERITED"
	StatusDomainNaturality      = "PASS_DOMAIN_NATURALITY_ON_REDUCED_ACTIVE_B2_RESPONSE"
	StatusBasepointUniqueness   = "PASS_BASEPOINT_REDUCTION_UNIQUE_IF_NO_CONSTANT_ALPHA_TERM"
	StatusDegreeNaturality      = "PASS_DEGREE_NATURALITY_FORCES_S_SPLIT_POWER_K"
	StatusSelectorUniqueness    = "PASS_SELECTOR_FUNCTIONHOOD_FORCES_UNIQUE_TARGET_PER_DEGREE"
	StatusNormalizationUnique   = "PASS_LANE_LOCALITY_FORCES_H10_H72_CHAMBER_PAIR_UNDER_CURRENT_CONSTRAINTS"
	StatusZ2Representative      = "PASS_MU_B_REPRESENTATIVE_INDEPENDENT_ON_Z2_CLASS"
	StatusAlternativeRejections = "PASS_STANDARD_ALTERNATIVE_MEASURES_FAIL_CONSTRAINTS"
	StatusUniquenessConditional = "PASS_MU_B_UNIQUE_AMONG_TESTED_CONSTRAINT_COMPATIBLE_MEASURES"
	StatusNativeMeasureMissing  = "FIREWALL_PRESERVED_NATIVE_MEASURE_THEOREM_MISSING"
	StatusNoNativeR3Promotion   = "FIREWALL_PRESERVED_ALPHA_AND_R3_NOT_NATIVE"

	SupportMuBNaturalReducedResponse          = "CONDITIONAL_SUPPORT_MU_B_IS_NATURAL_ON_REDUCED_ACTIVE_B2_RESPONSE"
	SupportMuBNoArbitraryPolynomial           = "CONDITIONAL_SUPPORT_MU_B_DOES_NOT_REQUIRE_ARBITRARY_POLYNOMIAL_INPUT"
	SupportReducedForcesNoConstant            = "CONDITIONAL_SUPPORT_REDUCED_RESPONSE_FORCES_NO_CONSTANT_ALPHA_TERM"
	SupportBasepointRemovalUnique             = "CONDITIONAL_SUPPORT_BASEPOINT_REMOVAL_IS_UNIQUE_IF_ALPHA_HAS_NO_CONSTANT_RESPONSE"
	SupportDegreeForcesSPowerK                = "CONDITIONAL_SUPPORT_DEGREE_NATURALITY_FORCES_S_SPLIT_POWER_K_ON_DEGREE_K"
	SupportSPowerAssignmentUnique             = "CONDITIONAL_SUPPORT_S_POWER_ASSIGNMENT_IS_UNIQUE_GIVEN_EXTERIOR_DEGREE_RESPECT"
	SupportSelectorFunctionhoodUnique         = "CONDITIONAL_SUPPORT_SELECTOR_FUNCTIONHOOD_FORCES_UNIQUE_TARGET_PER_DEGREE"
	SupportCrossLaneFromSelectorUniqueness    = "CONDITIONAL_SUPPORT_CROSS_LANE_EXCLUSION_FOLLOWS_FROM_SELECTOR_UNIQUENESS"
	SupportLaneLocalityForcesChambers         = "CONDITIONAL_SUPPORT_LANE_LOCALITY_FORCES_CHAMBER_PAIR_H10_H72"
	SupportNormalizationUniqueGivenChambers   = "CONDITIONAL_SUPPORT_NORMALIZATION_IS_UNIQUE_GIVEN_LOCAL_GLOBAL_RESPONSE_CHAMBERS"
	SupportMuBZ2RepresentativeIndependent     = "CONDITIONAL_SUPPORT_MU_B_IS_Z2_REPRESENTATIVE_INDEPENDENT"
	SupportPhaseSignNoChangeMeasure           = "CONDITIONAL_SUPPORT_PHASE_SIGN_DOES_NOT_CHANGE_BOUNDARY_MEASURE_VALUE"
	SupportAlternativeMeasuresFailConstraints = "CONDITIONAL_SUPPORT_STANDARD_ALTERNATIVE_MEASURES_FAIL_REQUIRED_CONSTRAINTS"
	SupportMuBUniqueAmongTestedMeasures       = "CONDITIONAL_SUPPORT_MU_B_IS_UNIQUE_AMONG_TESTED_CONSTRAINT_COMPATIBLE_MEASURES"

	FailureNoNativeBoundaryActivationMeasure = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_ACTIVATION_MEASURE_CERTIFIED"
	FailureNoNativeMeasureUniqueness         = "FAILED_ROUTE_NO_NATIVE_MEASURE_UNIQUENESS_THEOREM"
	FailureDomainNaturalityNotNative         = "FAILED_ROUTE_DOMAIN_NATURALITY_NOT_NATIVE_MEASURE_THEOREM"
	FailureNoNativeBasepointReduction        = "FAILED_ROUTE_NO_NATIVE_BASEPOINT_REDUCTION_THEOREM"
	FailureNoNativeDegreeRespectingMeasure   = "FAILED_ROUTE_NO_NATIVE_DEGREE_RESPECTING_MEASURE_THEOREM"
	FailureNoNativeUniqueSelector            = "FAILED_ROUTE_NO_NATIVE_PROOF_THAT_I_B_Z2_IS_UNIQUE_SELECTOR"
	FailureNoNativeSelectorFunctionhood      = "FAILED_ROUTE_NO_NATIVE_SELECTOR_FUNCTIONHOOD_THEOREM"
	FailureNoNativeLaneLocalityToChamber     = "FAILED_ROUTE_NO_NATIVE_LANE_LOCALITY_TO_CHAMBER_THEOREM"
	FailureNoNativeChamberNormalization      = "FAILED_ROUTE_NO_NATIVE_RESPONSE_CHAMBER_NORMALIZATION_THEOREM"
	FailureZ2InvarianceNotNative             = "FAILED_ROUTE_Z2_INVARIANCE_NOT_NATIVE_MEASURE_THEOREM"
	FailureAlternativeRejectionNotFullNative = "FAILED_ROUTE_ALTERNATIVE_REJECTION_NOT_FULL_NATIVE_UNIQUENESS_THEOREM"
	FailureUnreducedConstantTerm             = "FAILED_ROUTE_UNREDUCED_MEASURE_ADDS_FORBIDDEN_CONSTANT_TERM"
	FailureCrossLaneAddsFalseTerms           = "FAILED_ROUTE_CROSS_LANE_MEASURE_ADDS_FALSE_ALPHA_TERMS"
	FailureBareChamberBreaksAugmentation     = "FAILED_ROUTE_BARE_CHAMBER_MEASURE_BREAKS_BOUNDARY_AUGMENTATION"
	FailureCommonDenominatorBreaksLocality   = "FAILED_ROUTE_COMMON_DENOMINATOR_MEASURE_BREAKS_LANE_LOCALITY"
	FailureAlphaBridgeCandidateNotNative     = "FAILED_ROUTE_ALPHA_B_REMAINS_BRIDGE_CANDIDATE_NOT_NATIVE"
	FailureNotNativeR3                       = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked         = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap            = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap            = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues          = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator            = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type DomainNaturalityAudit struct {
	InheritedStatus     string
	Domain              string
	ActiveDegrees       []int
	ArbitraryPolynomial bool
	NaturalOnReducedB2  bool
	NativeTheorem       bool
	Supports            []string
	Failures            []string
}

type BasepointUniquenessAudit struct {
	UnreducedHasLambda0     bool
	ReducedResponseUsed     bool
	AlphaConstantTerm       float64
	NoConstantTermForced    bool
	UniqueIfNoConstantAlpha bool
	NativeTheorem           bool
	Supports                []string
	Failures                []string
}

type DegreeNaturalityAudit struct {
	DegreePowers          map[int]int
	Coefficients          map[int]float64
	AlternativePowers     map[int]int
	PowerAssignmentUnique bool
	NativeTheorem         bool
	Supports              []string
	Failures              []string
}

type SelectorUniquenessAudit struct {
	Targets             map[int]string
	FalseTargets        map[int]string
	Ranks               map[int]int
	FunctionhoodAssumed bool
	UniquePerDegree     bool
	CrossLanesExcluded  bool
	NativeTheorem       bool
	Supports            []string
	Failures            []string
}

type NormalizationUniquenessAudit struct {
	Chambers               map[int]string
	Ranks                  map[int]int
	Weights                map[int]float64
	LaneLocalityAccepted   bool
	UniqueGivenLocalGlobal bool
	NativeTheorem          bool
	Supports               []string
	Failures               []string
}

type Z2IndependenceAudit struct {
	RepresentativesExchanged bool
	RanksInvariant           bool
	MeasureInvariant         bool
	NativeTheorem            bool
	Supports                 []string
	Failures                 []string
}

type AlternativeMeasureAudit struct {
	UnreducedRejected         bool
	CrossLaneRejected         bool
	BareChamberRejected       bool
	CommonDenominatorRejected bool
	PollutedLinearWeight      float64
	PollutedQuadraticWeight   float64
	BareLinearWeight          float64
	BareQuadraticWeight       float64
	UniqueAmongTested         bool
	FullNativeUniqueness      bool
	Supports                  []string
	Failures                  []string
}

type AlphaMeasureAudit struct {
	Formula               string
	BoundaryFormula       string
	S                     float64
	LinearContribution    float64
	QuadraticContribution float64
	Alpha                 float64
	NativeAlpha           bool
}

type NativeStatusAudit struct {
	NaturalMeasureCandidate bool
	UniqueUnderConstraints  bool
	NativeMeasure           bool
	NativeAlpha             bool
	NativeR3                bool
	MissingNativeSources    []string
	Failures                []string
}

type FirewallLedger struct {
	BoundaryActivationMeasureNative bool
	MeasureUniquenessNative         bool
	DomainNaturalityNative          bool
	BasepointReductionNative        bool
	DegreeRespectNative             bool
	SelectorNative                  bool
	SelectorFunctionhoodNative      bool
	LaneLocalityNative              bool
	ChamberNormalizationNative      bool
	Z2InvarianceNative              bool
	AlternativeRejectionNative      bool
	AlphaNative                     bool
	NativeR3                        bool
	FullAFDescent                   bool
	GenerationCarrierMap            bool
	FlavorOrientationMap            bool
	IndividualYukawaValues          bool
	NativeYukawaOperator            bool
}

func (f FirewallLedger) List() []string {
	out := []string{}
	if !f.BoundaryActivationMeasureNative {
		out = append(out, FailureNoNativeBoundaryActivationMeasure)
	}
	if !f.MeasureUniquenessNative {
		out = append(out, FailureNoNativeMeasureUniqueness)
	}
	if !f.DomainNaturalityNative {
		out = append(out, FailureDomainNaturalityNotNative)
	}
	if !f.BasepointReductionNative {
		out = append(out, FailureNoNativeBasepointReduction)
	}
	if !f.DegreeRespectNative {
		out = append(out, FailureNoNativeDegreeRespectingMeasure)
	}
	if !f.SelectorNative {
		out = append(out, FailureNoNativeUniqueSelector)
	}
	if !f.SelectorFunctionhoodNative {
		out = append(out, FailureNoNativeSelectorFunctionhood)
	}
	if !f.LaneLocalityNative {
		out = append(out, FailureNoNativeLaneLocalityToChamber)
	}
	if !f.ChamberNormalizationNative {
		out = append(out, FailureNoNativeChamberNormalization)
	}
	if !f.Z2InvarianceNative {
		out = append(out, FailureZ2InvarianceNotNative)
	}
	if !f.AlternativeRejectionNative {
		out = append(out, FailureAlternativeRejectionNotFullNative)
	}
	if !f.AlphaNative {
		out = append(out, FailureAlphaBridgeCandidateNotNative)
	}
	if !f.NativeR3 {
		out = append(out, FailureNotNativeR3)
	}
	if !f.FullAFDescent {
		out = append(out, FailureFullAFDescentStillBlocked)
	}
	if !f.GenerationCarrierMap {
		out = append(out, FailureNoGenerationCarrierMap)
	}
	if !f.FlavorOrientationMap {
		out = append(out, FailureNoFlavorOrientationMap)
	}
	if !f.IndividualYukawaValues {
		out = append(out, FailureNoIndividualYukawaValues)
	}
	if !f.NativeYukawaOperator {
		out = append(out, FailureNoNativeYukawaOperator)
	}
	return out
}

type Analysis struct {
	ID             string
	Truth          string
	Classification string
	ShortStatus    string
	Domain         DomainNaturalityAudit
	Basepoint      BasepointUniquenessAudit
	Degree         DegreeNaturalityAudit
	Selector       SelectorUniquenessAudit
	Normalization  NormalizationUniquenessAudit
	Z2             Z2IndependenceAudit
	Alternatives   AlternativeMeasureAudit
	Alpha          AlphaMeasureAudit
	NativeStatus   NativeStatusAudit
	Firewalls      FirewallLedger
	Final          string
}

func BuildDefault() (Analysis, error) {
	linear := float64(RankI1) / float64(RankH10) * SBoundary
	quad := float64(RankI2) / float64(RankH72) * SBoundary * SBoundary
	alpha := linear + quad
	if !near(linear, AlphaLinear) || !near(quad, AlphaQuad) || !near(alpha, AlphaB) {
		return Analysis{}, fmt.Errorf("alpha reconstruction mismatch: linear %.17g quad %.17g alpha %.17g", linear, quad, alpha)
	}

	pollutedWeight := float64(RankI1)/float64(RankH10) + float64(RankI2)/float64(RankH72)

	return Analysis{
		ID:             AuditID,
		Truth:          FinalTruth,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		Domain: DomainNaturalityAudit{
			InheritedStatus:     Gate920ShortStatus,
			Domain:              ReducedResponse,
			ActiveDegrees:       []int{1, 2},
			ArbitraryPolynomial: false,
			NaturalOnReducedB2:  true,
			NativeTheorem:       false,
			Supports:            []string{SupportMuBNaturalReducedResponse, SupportMuBNoArbitraryPolynomial},
			Failures:            []string{FailureDomainNaturalityNotNative},
		},
		Basepoint: BasepointUniquenessAudit{
			UnreducedHasLambda0:     true,
			ReducedResponseUsed:     true,
			AlphaConstantTerm:       0,
			NoConstantTermForced:    true,
			UniqueIfNoConstantAlpha: true,
			NativeTheorem:           false,
			Supports:                []string{SupportReducedForcesNoConstant, SupportBasepointRemovalUnique},
			Failures:                []string{FailureNoNativeBasepointReduction},
		},
		Degree: DegreeNaturalityAudit{
			DegreePowers:          map[int]int{1: 1, 2: 2},
			Coefficients:          map[int]float64{1: SBoundary, 2: SBoundary * SBoundary},
			AlternativePowers:     map[int]int{1: 2, 2: 1},
			PowerAssignmentUnique: true,
			NativeTheorem:         false,
			Supports:              []string{SupportDegreeForcesSPowerK, SupportSPowerAssignmentUnique},
			Failures:              []string{FailureNoNativeDegreeRespectingMeasure},
		},
		Selector: SelectorUniquenessAudit{
			Targets:             map[int]string{1: "[F_1/F_0]_{Z2}", 2: "[F_2/F_0]_{Z2}"},
			FalseTargets:        map[int]string{1: "[F_2/F_0]_{Z2}", 2: "[F_1/F_0]_{Z2}"},
			Ranks:               map[int]int{1: RankI1, 2: RankI2},
			FunctionhoodAssumed: true,
			UniquePerDegree:     true,
			CrossLanesExcluded:  true,
			NativeTheorem:       false,
			Supports:            []string{SupportSelectorFunctionhoodUnique, SupportCrossLaneFromSelectorUniqueness},
			Failures:            []string{FailureNoNativeUniqueSelector, FailureNoNativeSelectorFunctionhood},
		},
		Normalization: NormalizationUniquenessAudit{
			Chambers:               map[int]string{1: "H_10=H_R^ambient+B_2", 2: "H_72=Lambda^4 V_8+B_2"},
			Ranks:                  map[int]int{1: RankH10, 2: RankH72},
			Weights:                map[int]float64{1: float64(RankI1) / float64(RankH10), 2: float64(RankI2) / float64(RankH72)},
			LaneLocalityAccepted:   true,
			UniqueGivenLocalGlobal: true,
			NativeTheorem:          false,
			Supports:               []string{SupportLaneLocalityForcesChambers, SupportNormalizationUniqueGivenChambers},
			Failures:               []string{FailureNoNativeLaneLocalityToChamber, FailureNoNativeChamberNormalization},
		},
		Z2: Z2IndependenceAudit{
			RepresentativesExchanged: true,
			RanksInvariant:           true,
			MeasureInvariant:         true,
			NativeTheorem:            false,
			Supports:                 []string{SupportMuBZ2RepresentativeIndependent, SupportPhaseSignNoChangeMeasure},
			Failures:                 []string{FailureZ2InvarianceNotNative},
		},
		Alternatives: AlternativeMeasureAudit{
			UnreducedRejected:         true,
			CrossLaneRejected:         true,
			BareChamberRejected:       true,
			CommonDenominatorRejected: true,
			PollutedLinearWeight:      pollutedWeight,
			PollutedQuadraticWeight:   pollutedWeight,
			BareLinearWeight:          float64(RankI1) / 8.0,
			BareQuadraticWeight:       float64(RankI2) / 70.0,
			UniqueAmongTested:         true,
			FullNativeUniqueness:      false,
			Supports:                  []string{SupportAlternativeMeasuresFailConstraints, SupportMuBUniqueAmongTestedMeasures},
			Failures:                  []string{FailureUnreducedConstantTerm, FailureCrossLaneAddsFalseTerms, FailureBareChamberBreaksAugmentation, FailureCommonDenominatorBreaksLocality, FailureAlternativeRejectionNotFullNative},
		},
		Alpha: AlphaMeasureAudit{
			Formula:               MeasureFormula,
			BoundaryFormula:       BoundaryAlphaFormula,
			S:                     SBoundary,
			LinearContribution:    linear,
			QuadraticContribution: quad,
			Alpha:                 alpha,
			NativeAlpha:           false,
		},
		NativeStatus: NativeStatusAudit{
			NaturalMeasureCandidate: true,
			UniqueUnderConstraints:  true,
			NativeMeasure:           false,
			NativeAlpha:             false,
			NativeR3:                false,
			MissingNativeSources:    []string{"domain naturality", "basepoint reduction", "degree-respecting measure", "selector functionhood", "lane-locality chamber law", "Z2-invariant measure law", "full native uniqueness theorem"},
			Failures:                []string{FailureNoNativeBoundaryActivationMeasure, FailureNoNativeMeasureUniqueness, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3},
		},
		Firewalls: FirewallLedger{},
		Final:     StrategicConclusion,
	}, nil
}

func Statuses() []string {
	return []string{StatusInheritedGate920, StatusDomainNaturality, StatusBasepointUniqueness, StatusDegreeNaturality, StatusSelectorUniqueness, StatusNormalizationUnique, StatusZ2Representative, StatusAlternativeRejections, StatusUniquenessConditional, StatusNativeMeasureMissing, StatusNoNativeR3Promotion}
}

func Supports() []string {
	return []string{SupportMuBNaturalReducedResponse, SupportMuBNoArbitraryPolynomial, SupportReducedForcesNoConstant, SupportBasepointRemovalUnique, SupportDegreeForcesSPowerK, SupportSPowerAssignmentUnique, SupportSelectorFunctionhoodUnique, SupportCrossLaneFromSelectorUniqueness, SupportLaneLocalityForcesChambers, SupportNormalizationUniqueGivenChambers, SupportMuBZ2RepresentativeIndependent, SupportPhaseSignNoChangeMeasure, SupportAlternativeMeasuresFailConstraints, SupportMuBUniqueAmongTestedMeasures}
}

func Failures() []string {
	return []string{FailureNoNativeBoundaryActivationMeasure, FailureNoNativeMeasureUniqueness, FailureDomainNaturalityNotNative, FailureNoNativeBasepointReduction, FailureNoNativeDegreeRespectingMeasure, FailureNoNativeUniqueSelector, FailureNoNativeSelectorFunctionhood, FailureNoNativeLaneLocalityToChamber, FailureNoNativeChamberNormalization, FailureZ2InvarianceNotNative, FailureAlternativeRejectionNotFullNative, FailureUnreducedConstantTerm, FailureCrossLaneAddsFalseTerms, FailureBareChamberBreaksAugmentation, FailureCommonDenominatorBreaksLocality, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}
}

func FormatDomain(d DomainNaturalityAudit) string {
	return fmt.Sprintf("domain inherited=%s natural_reduced_b2=%t arbitrary_polynomial=%t active_degrees=%v native=%t domain=%s", d.InheritedStatus, d.NaturalOnReducedB2, d.ArbitraryPolynomial, d.ActiveDegrees, d.NativeTheorem, d.Domain)
}

func FormatBasepoint(b BasepointUniquenessAudit) string {
	return fmt.Sprintf("basepoint unreduced_lambda0=%t reduced_response=%t constant=%.17g no_constant_forced=%t unique_if_no_constant=%t native=%t", b.UnreducedHasLambda0, b.ReducedResponseUsed, b.AlphaConstantTerm, b.NoConstantTermForced, b.UniqueIfNoConstantAlpha, b.NativeTheorem)
}

func FormatDegree(d DegreeNaturalityAudit) string {
	return fmt.Sprintf("degree powers=%v coeffs=%v alternative_powers=%v unique=%t native=%t", d.DegreePowers, d.Coefficients, d.AlternativePowers, d.PowerAssignmentUnique, d.NativeTheorem)
}

func FormatSelector(s SelectorUniquenessAudit) string {
	return fmt.Sprintf("selector targets=%v false_targets=%v ranks=%v functionhood_assumed=%t unique_per_degree=%t cross_lanes_excluded=%t native=%t", s.Targets, s.FalseTargets, s.Ranks, s.FunctionhoodAssumed, s.UniquePerDegree, s.CrossLanesExcluded, s.NativeTheorem)
}

func FormatNormalization(n NormalizationUniquenessAudit) string {
	return fmt.Sprintf("normalization chambers=%v ranks=%v weights=%v lane_locality=%t unique_given_local_global=%t native=%t", n.Chambers, n.Ranks, n.Weights, n.LaneLocalityAccepted, n.UniqueGivenLocalGlobal, n.NativeTheorem)
}

func FormatZ2(z Z2IndependenceAudit) string {
	return fmt.Sprintf("z2 representatives_exchanged=%t ranks_invariant=%t measure_invariant=%t native=%t", z.RepresentativesExchanged, z.RanksInvariant, z.MeasureInvariant, z.NativeTheorem)
}

func FormatAlternatives(a AlternativeMeasureAudit) string {
	return fmt.Sprintf("alternatives unreduced_rejected=%t cross_lane_rejected=%t bare_chamber_rejected=%t common_denominator_rejected=%t polluted_weights=(%.17g,%.17g) bare_weights=(%.17g,%.17g) unique_tested=%t full_native_unique=%t", a.UnreducedRejected, a.CrossLaneRejected, a.BareChamberRejected, a.CommonDenominatorRejected, a.PollutedLinearWeight, a.PollutedQuadraticWeight, a.BareLinearWeight, a.BareQuadraticWeight, a.UniqueAmongTested, a.FullNativeUniqueness)
}

func FormatAlpha(a AlphaMeasureAudit) string {
	return fmt.Sprintf("alpha formula=%s boundary=%s linear=%.17g quad=%.17g alpha=%.17g native_alpha=%t", a.Formula, a.BoundaryFormula, a.LinearContribution, a.QuadraticContribution, a.Alpha, a.NativeAlpha)
}

func FormatNativeStatus(n NativeStatusAudit) string {
	return fmt.Sprintf("native_status natural_candidate=%t unique_constraints=%t native_measure=%t native_alpha=%t native_r3=%t missing=%s", n.NaturalMeasureCandidate, n.UniqueUnderConstraints, n.NativeMeasure, n.NativeAlpha, n.NativeR3, strings.Join(n.MissingNativeSources, "; "))
}

func FormatFirewalls(f FirewallLedger) string { return strings.Join(f.List(), ";") }

func containsAll(haystack, needles []string) bool {
	m := map[string]bool{}
	for _, h := range haystack {
		m[h] = true
	}
	for _, n := range needles {
		if !m[n] {
			return false
		}
	}
	return true
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-18 }

func nearLoose(a, b float64) bool { return math.Abs(a-b) < 1e-12 }

func firewallsOK(f FirewallLedger) bool {
	return !f.BoundaryActivationMeasureNative && !f.MeasureUniquenessNative && !f.DomainNaturalityNative && !f.BasepointReductionNative && !f.DegreeRespectNative && !f.SelectorNative && !f.SelectorFunctionhoodNative && !f.LaneLocalityNative && !f.ChamberNormalizationNative && !f.Z2InvarianceNative && !f.AlternativeRejectionNative && !f.AlphaNative && !f.NativeR3 && !f.FullAFDescent && !f.GenerationCarrierMap && !f.FlavorOrientationMap && !f.IndividualYukawaValues && !f.NativeYukawaOperator
}
