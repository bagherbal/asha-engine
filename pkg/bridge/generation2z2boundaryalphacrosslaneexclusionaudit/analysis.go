// Package generation2z2boundaryalphacrosslaneexclusionaudit implements
// Gate 915: Z2 BoundaryAlpha CrossLane Exclusion Audit.
//
// Gate 915 follows Gate 914's degree-indexed Z2 airlock flag selector shape.
// It audits whether the selector typing excludes the false cross-lanes
//
//	Lambda^1 B_2 -> [F_2/F_0]_{Z2}
//	Lambda^2 B_2 -> [F_1/F_0]_{Z2}
//
// which would contaminate the sealed alpha response with (7/72)s and
// (3/10)s^2. The gate supports cross-lane exclusion by exposure/enclosure
// type separation, degree-selector functionhood, rank-contamination detection,
// cumulative-enclosure consistency, and Z2 compatibility. It does not certify
// a native uniqueness theorem for I_B^Z2, does not transport S_split natively,
// and does not promote BoundaryAlpha or R3 to native status.
package generation2z2boundaryalphacrosslaneexclusionaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE915-Z2-BOUNDARYALPHA-CROSSLANE-EXCLUSION-AUDIT"

	SBoundary = 0.0012924448188162962
	AlphaB    = 0.0003878958469680527

	RankF1OverF0   = 3
	RankF2OverF0   = 7
	RankF2OverF1   = 4
	LinearDenom    = 10
	QuadraticDenom = 72
	Lambda1Dim     = 2
	Lambda2Dim     = 1

	Gate914Classification = "R3_ALPHA_SUBOBJECT_2_Z2_FLAG_SELECTOR_SHAPE_PASS_NATIVE_FUNCTOR_BLOCKED"
	Gate914ShortStatus    = "R3_DEGREE_INDEXED_Z2_AIRLOCK_FLAG_SELECTOR_OBSTRUCTION"
	Gate914Verdict        = "DEGREE_INDEXED_Z2_FLAG_SELECTOR_SHAPE_SUPPORTED_BUT_NATIVE_FUNCTOR_NOT_CERTIFIED"

	CorrectLinearLane    = "deg(Lambda^1 B_2)->[F_1/F_0]_{Z2}"
	CorrectQuadraticLane = "deg(Lambda^2 B_2)->[F_2/F_0]_{Z2}"
	FalseLinearLane      = "deg(Lambda^1 B_2)->[F_2/F_0]_{Z2}"
	FalseQuadraticLane   = "deg(Lambda^2 B_2)->[F_1/F_0]_{Z2}"

	LinearExposureType        = "Lambda^1 B_2 has exposure type only"
	QuadraticEnclosureType    = "Lambda^2 B_2 has enclosure type only"
	Z2ExposedFaceClass        = "[F_1/F_0]_{Z2}"
	Z2FullEnclosureClass      = "[F_2/F_0]_{Z2}"
	CorrectAlphaFormula       = "alpha_B=(3/10)s+(7/72)s^2"
	PollutedAlphaFormula      = "alpha_B_polluted=(3/10)s+(7/72)s^2+(7/72)s+(3/10)s^2=(143/360)(s+s^2)"
	FalseLinearTerm           = "(7/72)s"
	FalseQuadraticTerm        = "(3/10)s^2"
	CumulativeEnclosureTarget = "Lambda^2 B_2->[F_2/F_0]_{Z2}"
	AssociatedGradedRejected  = "Lambda^2 B_2 not -> F_2/F_1"
	NextGate                  = "NEXT_PRESSURE_GATE916_S_SPLIT_TO_REDUCED_BOUNDARYPAIR_RESPONSE_TRANSPORT_AUDIT"
	Classification            = "R3_ALPHA_SUBOBJECT_3_CROSS_LANE_EXCLUSION_SHAPE_PASS_NATIVE_UNIQUENESS_BLOCKED"
	ShortStatus               = "R3_Z2_BOUNDARYALPHA_CROSS_LANE_EXCLUSION_OBSTRUCTION"
	FinalTruth                = "Z2_CROSS_LANE_EXCLUSION_SUPPORTED_BY_DEGREE_SELECTOR_TYPING_BUT_NATIVE_UNIQUENESS_NOT_CERTIFIED"
	StrategicConclusion       = "wrong cross-lanes are type-incompatible with the degree selector, but native exclusion still awaits a unique I_B^Z2 theorem plus S_split transport."

	StatusGate914Inherited        = "PASS_GATE914_DEGREE_INDEXED_SELECTOR_SHAPE_INHERITED"
	StatusNoLoopBack              = "PASS_BRANCH_DOES_NOT_REOPEN_PHASE_SOCKET_OR_REPRESENTATIVE_AIRLOCK_WOUNDS"
	StatusExposureEnclosureTyped  = "PASS_EXPOSURE_ENCLOSURE_TYPE_SEPARATION_RECORDED"
	StatusSelectorDeterminism     = "PASS_DEGREE_INDEXED_SELECTOR_DETERMINISM_BLOCKS_FALSE_TARGETS_CONDITIONALLY"
	StatusRankContamination       = "PASS_FALSE_CROSS_LANES_PRODUCE_WRONG_ALPHA_RESPONSE"
	StatusCumulativeCompatible    = "PASS_CROSS_LANE_EXCLUSION_COMPATIBLE_WITH_CUMULATIVE_ENCLOSURE_CHOICE"
	StatusZ2Compatible            = "PASS_CROSS_LANE_EXCLUSION_IS_Z2_CLASS_COMPATIBLE"
	StatusNativeUniquenessBlocked = "FIREWALL_PRESERVED_NATIVE_CROSS_LANE_UNIQUENESS_NOT_CERTIFIED"

	SupportCrossLanesExcludedByType          = "CONDITIONAL_SUPPORT_CROSS_LANES_EXCLUDED_BY_EXPOSURE_ENCLOSURE_TYPE_SEPARATION"
	SupportLambda1ExposureOnly               = "CONDITIONAL_SUPPORT_LAMBDA1B2_HAS_EXPOSURE_TYPE_ONLY"
	SupportLambda2EnclosureOnly              = "CONDITIONAL_SUPPORT_LAMBDA2B2_HAS_ENCLOSURE_TYPE_ONLY"
	SupportCrossLanesExcludedIfFunction      = "CONDITIONAL_SUPPORT_CROSS_LANES_EXCLUDED_IF_I_B_Z2_IS_A_FUNCTION"
	SupportSelectorDeterminismBlocksFalse    = "CONDITIONAL_SUPPORT_DEGREE_INDEXED_SELECTOR_DETERMINISM_BLOCKS_FALSE_TARGETS"
	SupportFalseCrossLanesWrongAlpha         = "CONDITIONAL_SUPPORT_FALSE_CROSS_LANES_PRODUCE_WRONG_ALPHA_RESPONSE"
	SupportActiveAlphaRequiresExclusion      = "CONDITIONAL_SUPPORT_ACTIVE_ALPHA_REQUIRES_CROSS_LANE_EXCLUSION"
	SupportCumulativeCompatible              = "CONDITIONAL_SUPPORT_CROSS_LANE_EXCLUSION_COMPATIBLE_WITH_CUMULATIVE_ENCLOSURE_CHOICE"
	SupportDegreeTwoRemainsF2OverF0          = "CONDITIONAL_SUPPORT_DEGREE_TWO_REMAINS_F2_OVER_F0_NOT_F2_OVER_F1"
	SupportCrossLaneExclusionZ2Compatible    = "CONDITIONAL_SUPPORT_CROSS_LANE_EXCLUSION_IS_Z2_CLASS_COMPATIBLE"
	SupportFalseLanesRepresentativeFreeFalse = "CONDITIONAL_SUPPORT_FALSE_LANES_ARE_REPRESENTATIVE_INDEPENDENTLY_FALSE"
	SupportRemainingAlphaSubobjectsSharpened = "CONDITIONAL_SUPPORT_ALPHA_SUBOBJECTS_ONE_TWO_THREE_SHARPENED_UNDER_FIREWALL"

	FailureNoNativeZ2CrossLaneTheorem            = "FAILED_ROUTE_NO_NATIVE_Z2_CROSS_LANE_EXCLUSION_THEOREM"
	FailureTypeSeparationNotNativeFunctorTheorem = "FAILED_ROUTE_EXPOSURE_ENCLOSURE_TYPE_SEPARATION_NOT_NATIVE_FUNCTOR_THEOREM"
	FailureNoNativeUniqueDegreeSelector          = "FAILED_ROUTE_NO_NATIVE_PROOF_THAT_I_B_Z2_IS_THE_UNIQUE_DEGREE_SELECTOR"
	FailureNumericalMismatchNotNativeExclusion   = "FAILED_ROUTE_NUMERICAL_MISMATCH_DETECTS_CROSS_LANE_ERROR_BUT_DOES_NOT_PROVE_NATIVE_EXCLUSION"
	FailureNoNativeCumulativeReason              = "FAILED_ROUTE_NO_NATIVE_REASON_YET_FOR_CUMULATIVE_ENCLOSURE_OVER_ASSOCIATED_GRADED_SLICE"
	FailureZ2CompatibilityNotNativeExclusion     = "FAILED_ROUTE_Z2_COMPATIBILITY_OF_EXCLUSION_NOT_NATIVE_EXCLUSION_THEOREM"
	FailureAlphaStillSealed                      = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureDenominatorsSTransportStillExternal   = "FAILED_ROUTE_DENOMINATORS_AND_S_TRANSPORT_STILL_EXTERNAL"
	FailureNotNativeR3                           = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked             = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap                = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap                = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues              = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator                = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type InheritedSelector struct {
	Gate914Classification  string
	Gate914ShortStatus     string
	Gate914Verdict         string
	SelectorShapeSupported bool
	ReopensPhaseSign       bool
	ReopensSocketOrder     bool
	ReopensRepresentative  bool
	DerivesAlpha           bool
	UpdatesOfficialLedger  bool
	Supports, Failures     []string
}

type Lane struct {
	Name              string
	Degree            int
	SourceType        string
	Target            string
	TargetRank        int
	Correct           bool
	TypeCompatible    bool
	Z2Compatible      bool
	FalseContribution string
}

type TypeSeparation struct {
	Lambda1Type        string
	Lambda2Type        string
	LinearFalseLane    Lane
	QuadraticFalseLane Lane
	ExcludesByType     bool
	NativeTheorem      bool
	Supports, Failures []string
}

type SelectorDeterminism struct {
	ObjectName             string
	CorrectLinearTarget    string
	CorrectQuadraticTarget string
	FalseLinearTarget      string
	FalseQuadraticTarget   string
	IsFunction             bool
	UniqueNativeSelector   bool
	ExcludesFalseTargets   bool
	Supports, Failures     []string
}

type RankContamination struct {
	CorrectAlpha         float64
	PollutedAlpha        float64
	FalseDelta           float64
	CorrectFormula       string
	PollutedFormula      string
	LinearCoefficient    float64
	QuadraticCoefficient float64
	PollutedCoefficient  float64
	FalseTerms           []string
	MismatchDetected     bool
	NativeExclusion      bool
	Supports, Failures   []string
}

type CumulativeConsistency struct {
	DegreeTwoTarget              string
	F2OverF0Rank                 int
	F2OverF1Rank                 int
	KeepsCumulativeEnclosure     bool
	RejectsAssociatedGradedSlice bool
	NativeReasonForChoice        bool
	Supports, Failures           []string
}

type Z2Compatibility struct {
	CorrectLanesRepresentativeFree bool
	FalseLanesRepresentativeFree   bool
	CorrectMapToCorrect            bool
	FalseMapToFalse                bool
	NativeExclusionTheorem         bool
	Supports, Failures             []string
}

type Firewalls struct {
	NativeZ2CrossLaneTheorem         bool
	TypeSeparationNativeFunctor      bool
	UniqueDegreeSelector             bool
	NumericalMismatchNativeExclusion bool
	NativeCumulativeReason           bool
	Z2CompatibilityNativeExclusion   bool
	AlphaNative                      bool
	DenominatorsExternal             bool
	STransportExternal               bool
	NativeR3                         bool
	FullAFDescent                    bool
	GenerationCarrierMap             bool
	FlavorOrientationMap             bool
	IndividualYukawaValues           bool
	NativeYukawaOperator             bool
}

type Audit struct {
	ID             string
	Truth          string
	Classification string
	ShortStatus    string
	Inherited      InheritedSelector
	CorrectLanes   []Lane
	TypeSep        TypeSeparation
	Determinism    SelectorDeterminism
	Contamination  RankContamination
	Cumulative     CumulativeConsistency
	Z2             Z2Compatibility
	Firewalls      Firewalls
	Final          string
}

func BuildDefault() (Audit, error) {
	correctAlpha := alphaCorrect(SBoundary)
	pollutedAlpha := alphaPolluted(SBoundary)
	if !near(correctAlpha, AlphaB) {
		return Audit{}, fmt.Errorf("correct alpha mismatch: got %.18g want %.18g", correctAlpha, AlphaB)
	}
	return Audit{
		ID:             AuditID,
		Truth:          FinalTruth,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		Inherited: InheritedSelector{
			Gate914Classification:  Gate914Classification,
			Gate914ShortStatus:     Gate914ShortStatus,
			Gate914Verdict:         Gate914Verdict,
			SelectorShapeSupported: true,
			ReopensPhaseSign:       false,
			ReopensSocketOrder:     false,
			ReopensRepresentative:  false,
			DerivesAlpha:           false,
			UpdatesOfficialLedger:  false,
			Supports:               []string{StatusGate914Inherited, StatusNoLoopBack},
			Failures:               []string{FailureNoNativeZ2CrossLaneTheorem, FailureAlphaStillSealed},
		},
		CorrectLanes: []Lane{
			{Name: CorrectLinearLane, Degree: 1, SourceType: "exposure", Target: Z2ExposedFaceClass, TargetRank: RankF1OverF0, Correct: true, TypeCompatible: true, Z2Compatible: true},
			{Name: CorrectQuadraticLane, Degree: 2, SourceType: "enclosure", Target: Z2FullEnclosureClass, TargetRank: RankF2OverF0, Correct: true, TypeCompatible: true, Z2Compatible: true},
		},
		TypeSep: TypeSeparation{
			Lambda1Type:        LinearExposureType,
			Lambda2Type:        QuadraticEnclosureType,
			LinearFalseLane:    Lane{Name: FalseLinearLane, Degree: 1, SourceType: "exposure", Target: Z2FullEnclosureClass, TargetRank: RankF2OverF0, Correct: false, TypeCompatible: false, Z2Compatible: true, FalseContribution: FalseLinearTerm},
			QuadraticFalseLane: Lane{Name: FalseQuadraticLane, Degree: 2, SourceType: "enclosure", Target: Z2ExposedFaceClass, TargetRank: RankF1OverF0, Correct: false, TypeCompatible: false, Z2Compatible: true, FalseContribution: FalseQuadraticTerm},
			ExcludesByType:     true,
			NativeTheorem:      false,
			Supports:           []string{StatusExposureEnclosureTyped, SupportCrossLanesExcludedByType, SupportLambda1ExposureOnly, SupportLambda2EnclosureOnly},
			Failures:           []string{FailureTypeSeparationNotNativeFunctorTheorem, FailureNoNativeZ2CrossLaneTheorem},
		},
		Determinism: SelectorDeterminism{
			ObjectName:             "I_B^Z2(k)=[F_k/F_0]_{Z2}",
			CorrectLinearTarget:    Z2ExposedFaceClass,
			CorrectQuadraticTarget: Z2FullEnclosureClass,
			FalseLinearTarget:      Z2FullEnclosureClass,
			FalseQuadraticTarget:   Z2ExposedFaceClass,
			IsFunction:             true,
			UniqueNativeSelector:   false,
			ExcludesFalseTargets:   true,
			Supports:               []string{StatusSelectorDeterminism, SupportCrossLanesExcludedIfFunction, SupportSelectorDeterminismBlocksFalse},
			Failures:               []string{FailureNoNativeUniqueDegreeSelector, FailureNoNativeZ2CrossLaneTheorem},
		},
		Contamination: RankContamination{
			CorrectAlpha:         correctAlpha,
			PollutedAlpha:        pollutedAlpha,
			FalseDelta:           pollutedAlpha - correctAlpha,
			CorrectFormula:       CorrectAlphaFormula,
			PollutedFormula:      PollutedAlphaFormula,
			LinearCoefficient:    float64(RankF1OverF0) / float64(LinearDenom),
			QuadraticCoefficient: float64(RankF2OverF0) / float64(QuadraticDenom),
			PollutedCoefficient:  float64(143) / float64(360),
			FalseTerms:           []string{FalseLinearTerm, FalseQuadraticTerm},
			MismatchDetected:     !near(correctAlpha, pollutedAlpha),
			NativeExclusion:      false,
			Supports:             []string{StatusRankContamination, SupportFalseCrossLanesWrongAlpha, SupportActiveAlphaRequiresExclusion},
			Failures:             []string{FailureNumericalMismatchNotNativeExclusion},
		},
		Cumulative: CumulativeConsistency{
			DegreeTwoTarget:              CumulativeEnclosureTarget,
			F2OverF0Rank:                 RankF2OverF0,
			F2OverF1Rank:                 RankF2OverF1,
			KeepsCumulativeEnclosure:     true,
			RejectsAssociatedGradedSlice: true,
			NativeReasonForChoice:        false,
			Supports:                     []string{StatusCumulativeCompatible, SupportCumulativeCompatible, SupportDegreeTwoRemainsF2OverF0},
			Failures:                     []string{FailureNoNativeCumulativeReason},
		},
		Z2: Z2Compatibility{
			CorrectLanesRepresentativeFree: true,
			FalseLanesRepresentativeFree:   true,
			CorrectMapToCorrect:            true,
			FalseMapToFalse:                true,
			NativeExclusionTheorem:         false,
			Supports:                       []string{StatusZ2Compatible, SupportCrossLaneExclusionZ2Compatible, SupportFalseLanesRepresentativeFreeFalse},
			Failures:                       []string{FailureZ2CompatibilityNotNativeExclusion},
		},
		Firewalls: Firewalls{
			NativeZ2CrossLaneTheorem:         false,
			TypeSeparationNativeFunctor:      false,
			UniqueDegreeSelector:             false,
			NumericalMismatchNativeExclusion: false,
			NativeCumulativeReason:           false,
			Z2CompatibilityNativeExclusion:   false,
			AlphaNative:                      false,
			DenominatorsExternal:             true,
			STransportExternal:               true,
			NativeR3:                         false,
			FullAFDescent:                    false,
			GenerationCarrierMap:             false,
			FlavorOrientationMap:             false,
			IndividualYukawaValues:           false,
			NativeYukawaOperator:             false,
		},
		Final: NextGate,
	}, nil
}

func alphaCorrect(s float64) float64 {
	return (float64(RankF1OverF0)/float64(LinearDenom))*s + (float64(RankF2OverF0)/float64(QuadraticDenom))*s*s
}

func alphaPolluted(s float64) float64 {
	return alphaCorrect(s) + (float64(RankF2OverF0)/float64(QuadraticDenom))*s + (float64(RankF1OverF0)/float64(LinearDenom))*s*s
}

func near(a, b float64) bool { return math.Abs(a-b) <= 1e-15 }

func containsAll(have []string, want []string) bool {
	m := make(map[string]bool, len(have))
	for _, v := range have {
		m[v] = true
	}
	for _, v := range want {
		if !m[v] {
			return false
		}
	}
	return true
}

func firewallsOK(f Firewalls) bool {
	return !f.NativeZ2CrossLaneTheorem && !f.TypeSeparationNativeFunctor && !f.UniqueDegreeSelector && !f.NumericalMismatchNativeExclusion && !f.NativeCumulativeReason && !f.Z2CompatibilityNativeExclusion && !f.AlphaNative && f.DenominatorsExternal && f.STransportExternal && !f.NativeR3 && !f.FullAFDescent && !f.GenerationCarrierMap && !f.FlavorOrientationMap && !f.IndividualYukawaValues && !f.NativeYukawaOperator
}

func (a Audit) FirewallsList() []string {
	return []string{
		FailureNoNativeZ2CrossLaneTheorem,
		FailureTypeSeparationNotNativeFunctorTheorem,
		FailureNoNativeUniqueDegreeSelector,
		FailureNumericalMismatchNotNativeExclusion,
		FailureNoNativeCumulativeReason,
		FailureZ2CompatibilityNotNativeExclusion,
		FailureAlphaStillSealed,
		FailureDenominatorsSTransportStillExternal,
		FailureNotNativeR3,
		FailureFullAFDescentStillBlocked,
		FailureNoGenerationCarrierMap,
		FailureNoFlavorOrientationMap,
		FailureNoIndividualYukawaValues,
		FailureNoNativeYukawaOperator,
	}
}

func Statuses() []string {
	return []string{StatusGate914Inherited, StatusNoLoopBack, StatusExposureEnclosureTyped, StatusSelectorDeterminism, StatusRankContamination, StatusCumulativeCompatible, StatusZ2Compatible, StatusNativeUniquenessBlocked}
}

func Supports() []string {
	return []string{SupportCrossLanesExcludedByType, SupportLambda1ExposureOnly, SupportLambda2EnclosureOnly, SupportCrossLanesExcludedIfFunction, SupportSelectorDeterminismBlocksFalse, SupportFalseCrossLanesWrongAlpha, SupportActiveAlphaRequiresExclusion, SupportCumulativeCompatible, SupportDegreeTwoRemainsF2OverF0, SupportCrossLaneExclusionZ2Compatible, SupportFalseLanesRepresentativeFreeFalse, SupportRemainingAlphaSubobjectsSharpened}
}

func Failures() []string {
	return []string{FailureNoNativeZ2CrossLaneTheorem, FailureTypeSeparationNotNativeFunctorTheorem, FailureNoNativeUniqueDegreeSelector, FailureNumericalMismatchNotNativeExclusion, FailureNoNativeCumulativeReason, FailureZ2CompatibilityNotNativeExclusion, FailureAlphaStillSealed, FailureDenominatorsSTransportStillExternal, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}
}

func FormatInherited(x InheritedSelector) string {
	return fmt.Sprintf("Gate914={classification:%s short:%s verdict:%s selector_shape:%t reopen_phase:%t reopen_socket:%t reopen_rep:%t derives_alpha:%t updates_official:%t}", x.Gate914Classification, x.Gate914ShortStatus, x.Gate914Verdict, x.SelectorShapeSupported, x.ReopensPhaseSign, x.ReopensSocketOrder, x.ReopensRepresentative, x.DerivesAlpha, x.UpdatesOfficialLedger)
}

func FormatLane(l Lane) string {
	return fmt.Sprintf("lane={name:%s degree:%d source_type:%s target:%s rank:%d correct:%t type_compatible:%t z2:%t false_term:%s}", l.Name, l.Degree, l.SourceType, l.Target, l.TargetRank, l.Correct, l.TypeCompatible, l.Z2Compatible, l.FalseContribution)
}

func FormatTypeSeparation(x TypeSeparation) string {
	return fmt.Sprintf("type_separation={lambda1:%s lambda2:%s excludes_by_type:%t native:%t false_linear:%s false_quadratic:%s}", x.Lambda1Type, x.Lambda2Type, x.ExcludesByType, x.NativeTheorem, FormatLane(x.LinearFalseLane), FormatLane(x.QuadraticFalseLane))
}

func FormatDeterminism(x SelectorDeterminism) string {
	return fmt.Sprintf("selector_determinism={object:%s function:%t unique_native:%t excludes_false:%t correct1:%s correct2:%s false1:%s false2:%s}", x.ObjectName, x.IsFunction, x.UniqueNativeSelector, x.ExcludesFalseTargets, x.CorrectLinearTarget, x.CorrectQuadraticTarget, x.FalseLinearTarget, x.FalseQuadraticTarget)
}

func FormatContamination(x RankContamination) string {
	return fmt.Sprintf("rank_contamination={correct:%.18g polluted:%.18g delta:%.18g correct_formula:%s polluted_formula:%s polluted_coeff:%.18g mismatch:%t native:%t false_terms:%s}", x.CorrectAlpha, x.PollutedAlpha, x.FalseDelta, x.CorrectFormula, x.PollutedFormula, x.PollutedCoefficient, x.MismatchDetected, x.NativeExclusion, strings.Join(x.FalseTerms, ","))
}

func FormatCumulative(x CumulativeConsistency) string {
	return fmt.Sprintf("cumulative={target:%s f2_f0:%d f2_f1:%d keeps_cumulative:%t rejects_associated:%t native_reason:%t}", x.DegreeTwoTarget, x.F2OverF0Rank, x.F2OverF1Rank, x.KeepsCumulativeEnclosure, x.RejectsAssociatedGradedSlice, x.NativeReasonForChoice)
}

func FormatZ2(x Z2Compatibility) string {
	return fmt.Sprintf("z2={correct_rep_free:%t false_rep_free:%t correct_to_correct:%t false_to_false:%t native_exclusion:%t}", x.CorrectLanesRepresentativeFree, x.FalseLanesRepresentativeFree, x.CorrectMapToCorrect, x.FalseMapToFalse, x.NativeExclusionTheorem)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls={native_cross_lane:%t type_native:%t unique_selector:%t mismatch_native:%t cumulative_native:%t z2_native:%t alpha_native:%t denom_external:%t s_transport_external:%t native_r3:%t full_af:%t generation:%t flavor:%t individual_yukawa:%t native_yukawa:%t}", f.NativeZ2CrossLaneTheorem, f.TypeSeparationNativeFunctor, f.UniqueDegreeSelector, f.NumericalMismatchNativeExclusion, f.NativeCumulativeReason, f.Z2CompatibilityNativeExclusion, f.AlphaNative, f.DenominatorsExternal, f.STransportExternal, f.NativeR3, f.FullAFDescent, f.GenerationCarrierMap, f.FlavorOrientationMap, f.IndividualYukawaValues, f.NativeYukawaOperator)
}
