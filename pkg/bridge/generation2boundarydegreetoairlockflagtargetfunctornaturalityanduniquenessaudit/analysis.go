// Package generation2boundarydegreetoairlockflagtargetfunctornaturalityanduniquenessaudit implements
// Gate 926: BoundaryDegree-to-AirlockFlag TargetFunctor Naturality and Uniqueness Audit.
//
// Gate 926 follows Gate 925's target-functor shape support. It audits whether
// Theta_B^Z2(k)=[F_k/F_0]_{Z2} is the unique natural selector under the current
// order, exposure/enclosure, Z2-invariance, cumulative-enclosure, and alpha-rank
// constraints. The gate supports uniqueness under these constraints while
// preserving the firewall that the constraints do not yet amount to a native
// target-functor theorem.
package generation2boundarydegreetoairlockflagtargetfunctornaturalityanduniquenessaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE926-BOUNDARYDEGREE-TO-AIRLOCKFLAG-TARGETFUNCTOR-NATURALITY-UNIQUENESS-AUDIT"

	Gate925ShortStatus = "R3_ALPHA_TARGET_FUNCTOR_SHAPE_SUPPORTED_NATIVE_THETA_MISSING"

	SourceChain             = "deg(Lambda^1 B_2)<deg(Lambda^2 B_2)"
	TargetChain             = "[F_1/F_0]_{Z2}<[F_2/F_0]_{Z2}"
	ThetaFunctor            = "Theta_B^Z2(k)=[F_k/F_0]_{Z2}"
	ThetaOne                = "Theta_B^Z2(1)=[F_1/F_0]_{Z2}"
	ThetaTwo                = "Theta_B^Z2(2)=[F_2/F_0]_{Z2}"
	SwappedAssignment       = "1->[F_2/F_0]_{Z2};2->[F_1/F_0]_{Z2}"
	AssociatedGradedTarget  = "Theta_alt(2)=F_2/F_1"
	MeasureFormula          = "mu_B(R_B(S_split))=rank(Theta_B^Z2(1))/rank(H_10)*S_split+rank(Theta_B^Z2(2))/rank(H_72)*S_split^2"
	AlphaFormula            = "alpha_B=(3/10)S_split+(7/72)S_split^2"
	NextGate                = "NEXT_PRESSURE_GATE927_BOUNDARYDEGREE_AIRLOCKFLAG_INCIDENCE_SOURCE_AUDIT"
	ExposureType            = "Lambda^1 B_2=one-factor boundary exposure"
	EnclosureType           = "Lambda^2 B_2=two-factor top boundary enclosure"
	LambdaRepresentativeOne = "Theta_B(1)=e_lambda tensor P_3"
	LambdaRepresentativeTwo = "Theta_B(2)=(C_R^2 tensor W)-(e_lambda tensor P_1)"
	BarRepresentativeOne    = "Theta_B(1)=e_barlambda tensor P_3"
	BarRepresentativeTwo    = "Theta_B(2)=(C_R^2 tensor W)-(e_barlambda tensor P_1)"

	RankF1OverF0 = 3
	RankF2OverF0 = 7
	RankF2OverF1 = 4
	RankH10      = 10
	RankH72      = 72

	Classification = "R3_THETA_B_Z2_NATURALITY_UNIQUENESS_CANDIDATE_NOT_NATIVE"
	ShortStatus    = "R3_ALPHA_TARGET_FUNCTOR_UNIQUE_UNDER_CONSTRAINTS_NATIVE_SOURCE_MISSING"
	FinalTruth     = "THETA_B_Z2_IS_UNIQUE_NATURAL_TARGET_FUNCTOR_UNDER_ORDER_EXPOSURE_ENCLOSURE_AND_Z2_CONSTRAINTS_BUT_NATIVE_SOURCE_MISSING"

	StatusInheritedGate925            = "PASS_GATE925_THETA_B_Z2_TARGET_FUNCTOR_SHAPE_INHERITED"
	StatusOrderPreservingUnique       = "PASS_ORDER_PRESERVING_SELECTOR_UNIQUE_BETWEEN_TWO_LEVEL_CHAINS"
	StatusExposureEnclosureUnique     = "PASS_EXPOSURE_ENCLOSURE_TYPING_FORCES_THETA_ASSIGNMENT_UNDER_CONSTRAINTS"
	StatusZ2RepresentativeIndependent = "PASS_THETA_B_Z2_REPRESENTATIVE_INDEPENDENT_AT_CLASS_LEVEL"
	StatusAssociatedGradedRejected    = "PASS_ASSOCIATED_GRADED_ALTERNATIVE_REJECTED_BY_TYPE_AND_RANK"
	StatusConstantCrossLaneRejected   = "PASS_CONSTANT_AND_CROSS_LANE_ALTERNATIVES_REJECTED"
	StatusMuBStrengthened             = "PASS_UNIQUE_THETA_STRENGTHENS_BOUNDARY_ACTIVATION_MEASURE"
	StatusNativeThetaSourceMissing    = "FIREWALL_PRESERVED_NATIVE_THETA_B_Z2_SOURCE_MISSING"
	StatusNativeAlphaR3StillBlocked   = "FIREWALL_PRESERVED_ALPHA_AND_R3_NOT_NATIVE"

	SupportOrderPreservingUnique            = "CONDITIONAL_SUPPORT_ORDER_PRESERVING_SELECTOR_IS_UNIQUE_BETWEEN_TWO_LEVEL_CHAINS"
	SupportSwappedOrderReversing            = "CONDITIONAL_SUPPORT_SWAPPED_TARGET_ASSIGNMENT_IS_ORDER_REVERSING"
	SupportThetaUniqueOrderPreserving       = "CONDITIONAL_SUPPORT_THETA_B_Z2_IS_THE_UNIQUE_ORDER_PRESERVING_TWO_LEVEL_SELECTOR"
	SupportExposureTargetsF1                = "CONDITIONAL_SUPPORT_EXPOSURE_TYPE_UNIQUELY_TARGETS_F1_OVER_F0"
	SupportEnclosureTargetsF2               = "CONDITIONAL_SUPPORT_ENCLOSURE_TYPE_UNIQUELY_TARGETS_F2_OVER_F0"
	SupportExposureEnclosureForcesTheta     = "CONDITIONAL_SUPPORT_EXPOSURE_ENCLOSURE_TYPING_FORCES_THETA_B_Z2_ASSIGNMENT"
	SupportThetaRepresentativeIndependent   = "CONDITIONAL_SUPPORT_THETA_B_Z2_IS_REPRESENTATIVE_INDEPENDENT"
	SupportThetaCommutesZ2                  = "CONDITIONAL_SUPPORT_THETA_B_Z2_COMMUTES_WITH_GLOBAL_PHASE_Z2_FLIP"
	SupportThetaRanksZ2Invariant            = "CONDITIONAL_SUPPORT_THETA_B_Z2_TARGET_RANKS_ARE_Z2_INVARIANT"
	SupportAssociatedGradedFailsType        = "CONDITIONAL_SUPPORT_ASSOCIATED_GRADED_TARGET_FAILS_CUMULATIVE_ENCLOSURE_TYPE"
	SupportAssociatedGradedFailsRank        = "CONDITIONAL_SUPPORT_ASSOCIATED_GRADED_TARGET_FAILS_ALPHA_RANK_REQUIREMENT"
	SupportF2OverF1RejectedByTypeAndRank    = "CONDITIONAL_SUPPORT_F2_OVER_F1_REJECTED_BY_BOTH_TYPE_AND_RANK"
	SupportDegreeZeroAbsentByReduction      = "CONDITIONAL_SUPPORT_DEGREE_ZERO_TARGET_ABSENT_BY_REDUCED_RESPONSE"
	SupportCrossLaneViolatesType            = "CONDITIONAL_SUPPORT_CROSS_LANE_ASSIGNMENTS_VIOLATE_EXPOSURE_ENCLOSURE_TYPE"
	SupportCrossLaneFailsAlphaShape         = "CONDITIONAL_SUPPORT_CROSS_LANE_ASSIGNMENTS_FAIL_ALPHA_RESPONSE_SHAPE"
	SupportThetaUniquenessStrengthensMuB    = "CONDITIONAL_SUPPORT_THETA_B_Z2_UNIQUENESS_STRENGTHENS_BOUNDARY_ACTIVATION_MEASURE"
	SupportSelectorAndCrossLaneFollowUnique = "CONDITIONAL_SUPPORT_SELECTOR_FUNCTIONHOOD_AND_CROSS_LANE_EXCLUSION_FOLLOW_FROM_UNIQUE_THETA_B_Z2"
	SupportMuBTargetRanksFixedByUniqueTheta = "CONDITIONAL_SUPPORT_MU_B_TARGET_RANKS_ARE_FIXED_BY_UNIQUE_THETA_B_Z2"

	FailureThetaNotNative                    = "FAILED_ROUTE_THETA_B_Z2_NOT_NATIVE_TARGET_FUNCTOR"
	FailureOrderPreservationNotNative        = "FAILED_ROUTE_ORDER_PRESERVATION_NOT_YET_NATIVE_TARGET_FUNCTOR_THEOREM"
	FailureExposureEnclosureUniquenessNative = "FAILED_ROUTE_EXPOSURE_ENCLOSURE_UNIQUENESS_NOT_NATIVE_FUNCTOR_THEOREM"
	FailureZ2IndependenceNotNative           = "FAILED_ROUTE_Z2_REPRESENTATIVE_INDEPENDENCE_NOT_NATIVE_FUNCTOR_THEOREM"
	FailureF2OverF1RejectionNotNative        = "FAILED_ROUTE_REJECTION_OF_F2_OVER_F1_NOT_NATIVE_CUMULATIVE_ENCLOSURE_THEOREM"
	FailureAlternativeRejectionNotNative     = "FAILED_ROUTE_ALTERNATIVE_REJECTION_NOT_FULL_NATIVE_UNIQUENESS_THEOREM"
	FailureMuBNotNativeWithoutThetaSource    = "FAILED_ROUTE_MU_B_STILL_NOT_NATIVE_WITHOUT_NATIVE_THETA_B_Z2_SOURCE"
	FailureAlphaBridgeCandidateNotNative     = "FAILED_ROUTE_ALPHA_B_REMAINS_BRIDGE_CANDIDATE_NOT_NATIVE"
	FailureNotNativeR3                       = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked         = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap            = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap            = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues          = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator            = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type OrderUniquenessAudit struct {
	SourceChain           string
	TargetChain           string
	SourceLevels          int
	TargetLevels          int
	NontrivialBijection   bool
	OrderPreserving       bool
	SwappedOrderReversing bool
	UniqueUnderOrder      bool
	NativeOrderTheorem    bool
	Supports              []string
	Failures              []string
}

type TypeUniquenessAudit struct {
	DegreeOneType         string
	DegreeTwoType         string
	ExposureTarget        string
	EnclosureTarget       string
	ExposureUnique        bool
	EnclosureUnique       bool
	ForcesThetaAssignment bool
	NativeTypeTheorem     bool
	Supports              []string
	Failures              []string
}

type Z2IndependenceAudit struct {
	LambdaOne                 string
	LambdaTwo                 string
	BarOne                    string
	BarTwo                    string
	RankOne                   int
	RankTwo                   int
	RepresentativeIndependent bool
	CommutesWithFlip          bool
	NativeZ2Theorem           bool
	Supports                  []string
	Failures                  []string
}

type AssociatedGradedRejectionAudit struct {
	AlternativeTarget       string
	AlternativeRank         int
	RequiredTarget          string
	RequiredRank            int
	FailsCumulativeType     bool
	FailsAlphaRank          bool
	RejectedByTypeAndRank   bool
	NativeCumulativeTheorem bool
	Supports                []string
	Failures                []string
}

type AlternativeRejectionAudit struct {
	DegreeZeroAbsent           bool
	CrossLaneExposureToEnclose bool
	CrossLaneEnclosureToExpose bool
	FalseLinearTerm            string
	FalseQuadraticTerm         string
	ViolatesType               bool
	FailsAlphaShape            bool
	NativeUniquenessTheorem    bool
	Supports                   []string
	Failures                   []string
}

type MeasureConsequenceAudit struct {
	MeasureFormula       string
	ThetaRankOne         int
	ThetaRankTwo         int
	H10Rank              int
	H72Rank              int
	AlphaFormula         string
	SelectorFunctionhood bool
	CrossLaneExclusion   bool
	TargetRanksFixed     bool
	NativeMeasure        bool
	Supports             []string
	Failures             []string
}

type FirewallLedger struct {
	ThetaNotNative                    bool
	OrderPreservationNotNative        bool
	ExposureEnclosureUniquenessNative bool
	Z2IndependenceNotNative           bool
	F2OverF1RejectionNotNative        bool
	AlternativeRejectionNotNative     bool
	MuBNotNativeWithoutThetaSource    bool
	AlphaBridgeCandidateNotNative     bool
	NotNativeR3                       bool
	FullAFDescentStillBlocked         bool
	NoGenerationCarrierMap            bool
	NoFlavorOrientationMap            bool
	NoIndividualYukawaValues          bool
	NoNativeYukawaOperator            bool
}

func (f FirewallLedger) List() []string {
	var out []string
	if f.ThetaNotNative {
		out = append(out, FailureThetaNotNative)
	}
	if f.OrderPreservationNotNative {
		out = append(out, FailureOrderPreservationNotNative)
	}
	if f.ExposureEnclosureUniquenessNative {
		out = append(out, FailureExposureEnclosureUniquenessNative)
	}
	if f.Z2IndependenceNotNative {
		out = append(out, FailureZ2IndependenceNotNative)
	}
	if f.F2OverF1RejectionNotNative {
		out = append(out, FailureF2OverF1RejectionNotNative)
	}
	if f.AlternativeRejectionNotNative {
		out = append(out, FailureAlternativeRejectionNotNative)
	}
	if f.MuBNotNativeWithoutThetaSource {
		out = append(out, FailureMuBNotNativeWithoutThetaSource)
	}
	if f.AlphaBridgeCandidateNotNative {
		out = append(out, FailureAlphaBridgeCandidateNotNative)
	}
	if f.NotNativeR3 {
		out = append(out, FailureNotNativeR3)
	}
	if f.FullAFDescentStillBlocked {
		out = append(out, FailureFullAFDescentStillBlocked)
	}
	if f.NoGenerationCarrierMap {
		out = append(out, FailureNoGenerationCarrierMap)
	}
	if f.NoFlavorOrientationMap {
		out = append(out, FailureNoFlavorOrientationMap)
	}
	if f.NoIndividualYukawaValues {
		out = append(out, FailureNoIndividualYukawaValues)
	}
	if f.NoNativeYukawaOperator {
		out = append(out, FailureNoNativeYukawaOperator)
	}
	return out
}

type Audit struct {
	ID              string
	InheritedStatus string
	Truth           string
	Classification  string
	ShortStatus     string
	Order           OrderUniquenessAudit
	Types           TypeUniquenessAudit
	Z2              Z2IndependenceAudit
	Graded          AssociatedGradedRejectionAudit
	Alternatives    AlternativeRejectionAudit
	Measure         MeasureConsequenceAudit
	Firewalls       FirewallLedger
	Final           string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		ID:              AuditID,
		InheritedStatus: Gate925ShortStatus,
		Truth:           FinalTruth,
		Classification:  Classification,
		ShortStatus:     ShortStatus,
		Order: OrderUniquenessAudit{
			SourceChain:           SourceChain,
			TargetChain:           TargetChain,
			SourceLevels:          2,
			TargetLevels:          2,
			NontrivialBijection:   true,
			OrderPreserving:       true,
			SwappedOrderReversing: true,
			UniqueUnderOrder:      true,
			NativeOrderTheorem:    false,
			Supports:              []string{SupportOrderPreservingUnique, SupportSwappedOrderReversing, SupportThetaUniqueOrderPreserving},
			Failures:              []string{FailureOrderPreservationNotNative},
		},
		Types: TypeUniquenessAudit{
			DegreeOneType:         ExposureType,
			DegreeTwoType:         EnclosureType,
			ExposureTarget:        "[F_1/F_0]_{Z2}",
			EnclosureTarget:       "[F_2/F_0]_{Z2}",
			ExposureUnique:        true,
			EnclosureUnique:       true,
			ForcesThetaAssignment: true,
			NativeTypeTheorem:     false,
			Supports:              []string{SupportExposureTargetsF1, SupportEnclosureTargetsF2, SupportExposureEnclosureForcesTheta},
			Failures:              []string{FailureExposureEnclosureUniquenessNative},
		},
		Z2: Z2IndependenceAudit{
			LambdaOne:                 LambdaRepresentativeOne,
			LambdaTwo:                 LambdaRepresentativeTwo,
			BarOne:                    BarRepresentativeOne,
			BarTwo:                    BarRepresentativeTwo,
			RankOne:                   RankF1OverF0,
			RankTwo:                   RankF2OverF0,
			RepresentativeIndependent: true,
			CommutesWithFlip:          true,
			NativeZ2Theorem:           false,
			Supports:                  []string{SupportThetaRepresentativeIndependent, SupportThetaCommutesZ2, SupportThetaRanksZ2Invariant},
			Failures:                  []string{FailureZ2IndependenceNotNative},
		},
		Graded: AssociatedGradedRejectionAudit{
			AlternativeTarget:       "F_2/F_1",
			AlternativeRank:         RankF2OverF1,
			RequiredTarget:          "[F_2/F_0]_{Z2}",
			RequiredRank:            RankF2OverF0,
			FailsCumulativeType:     true,
			FailsAlphaRank:          true,
			RejectedByTypeAndRank:   true,
			NativeCumulativeTheorem: false,
			Supports:                []string{SupportAssociatedGradedFailsType, SupportAssociatedGradedFailsRank, SupportF2OverF1RejectedByTypeAndRank},
			Failures:                []string{FailureF2OverF1RejectionNotNative},
		},
		Alternatives: AlternativeRejectionAudit{
			DegreeZeroAbsent:           true,
			CrossLaneExposureToEnclose: true,
			CrossLaneEnclosureToExpose: true,
			FalseLinearTerm:            "(7/72)S_split",
			FalseQuadraticTerm:         "(3/10)S_split^2",
			ViolatesType:               true,
			FailsAlphaShape:            true,
			NativeUniquenessTheorem:    false,
			Supports:                   []string{SupportDegreeZeroAbsentByReduction, SupportCrossLaneViolatesType, SupportCrossLaneFailsAlphaShape},
			Failures:                   []string{FailureAlternativeRejectionNotNative},
		},
		Measure: MeasureConsequenceAudit{
			MeasureFormula:       MeasureFormula,
			ThetaRankOne:         RankF1OverF0,
			ThetaRankTwo:         RankF2OverF0,
			H10Rank:              RankH10,
			H72Rank:              RankH72,
			AlphaFormula:         AlphaFormula,
			SelectorFunctionhood: true,
			CrossLaneExclusion:   true,
			TargetRanksFixed:     true,
			NativeMeasure:        false,
			Supports:             []string{SupportThetaUniquenessStrengthensMuB, SupportSelectorAndCrossLaneFollowUnique, SupportMuBTargetRanksFixedByUniqueTheta},
			Failures:             []string{FailureMuBNotNativeWithoutThetaSource, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3},
		},
		Firewalls: FirewallLedger{
			ThetaNotNative:                    true,
			OrderPreservationNotNative:        true,
			ExposureEnclosureUniquenessNative: true,
			Z2IndependenceNotNative:           true,
			F2OverF1RejectionNotNative:        true,
			AlternativeRejectionNotNative:     true,
			MuBNotNativeWithoutThetaSource:    true,
			AlphaBridgeCandidateNotNative:     true,
			NotNativeR3:                       true,
			FullAFDescentStillBlocked:         true,
			NoGenerationCarrierMap:            true,
			NoFlavorOrientationMap:            true,
			NoIndividualYukawaValues:          true,
			NoNativeYukawaOperator:            true,
		},
		Final: "Gate 926 supports Theta_B^Z2 as the unique natural target functor under order, exposure/enclosure, Z2, cumulative-enclosure, and alpha-rank constraints, but no native source theorem yet forces those constraints.",
	}
	if err := validate(a); err != nil {
		return Audit{}, err
	}
	return a, nil
}

func validate(a Audit) error {
	if a.InheritedStatus != Gate925ShortStatus {
		return fmt.Errorf("unexpected inherited status: %s", a.InheritedStatus)
	}
	if !a.Order.NontrivialBijection || !a.Order.OrderPreserving || !a.Order.SwappedOrderReversing || !a.Order.UniqueUnderOrder || a.Order.NativeOrderTheorem || a.Order.SourceLevels != 2 || a.Order.TargetLevels != 2 {
		return fmt.Errorf("bad order audit: %s", FormatOrder(a.Order))
	}
	if !a.Types.ExposureUnique || !a.Types.EnclosureUnique || !a.Types.ForcesThetaAssignment || a.Types.NativeTypeTheorem {
		return fmt.Errorf("bad type audit: %s", FormatTypes(a.Types))
	}
	if !a.Z2.RepresentativeIndependent || !a.Z2.CommutesWithFlip || a.Z2.NativeZ2Theorem || a.Z2.RankOne != RankF1OverF0 || a.Z2.RankTwo != RankF2OverF0 {
		return fmt.Errorf("bad z2 audit: %s", FormatZ2(a.Z2))
	}
	if !a.Graded.FailsCumulativeType || !a.Graded.FailsAlphaRank || !a.Graded.RejectedByTypeAndRank || a.Graded.NativeCumulativeTheorem || a.Graded.AlternativeRank != RankF2OverF1 || a.Graded.RequiredRank != RankF2OverF0 {
		return fmt.Errorf("bad graded audit: %s", FormatGraded(a.Graded))
	}
	if !a.Alternatives.DegreeZeroAbsent || !a.Alternatives.CrossLaneExposureToEnclose || !a.Alternatives.CrossLaneEnclosureToExpose || !a.Alternatives.ViolatesType || !a.Alternatives.FailsAlphaShape || a.Alternatives.NativeUniquenessTheorem {
		return fmt.Errorf("bad alternative audit: %s", FormatAlternatives(a.Alternatives))
	}
	if !a.Measure.SelectorFunctionhood || !a.Measure.CrossLaneExclusion || !a.Measure.TargetRanksFixed || a.Measure.NativeMeasure || a.Measure.ThetaRankOne != RankF1OverF0 || a.Measure.ThetaRankTwo != RankF2OverF0 {
		return fmt.Errorf("bad measure consequence: %s", FormatMeasure(a.Measure))
	}
	if !firewallsOK(a.Firewalls) {
		return fmt.Errorf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	return nil
}

func Statuses() []string {
	return []string{StatusInheritedGate925, StatusOrderPreservingUnique, StatusExposureEnclosureUnique, StatusZ2RepresentativeIndependent, StatusAssociatedGradedRejected, StatusConstantCrossLaneRejected, StatusMuBStrengthened, StatusNativeThetaSourceMissing, StatusNativeAlphaR3StillBlocked}
}
func Supports() []string {
	return []string{SupportOrderPreservingUnique, SupportSwappedOrderReversing, SupportThetaUniqueOrderPreserving, SupportExposureTargetsF1, SupportEnclosureTargetsF2, SupportExposureEnclosureForcesTheta, SupportThetaRepresentativeIndependent, SupportThetaCommutesZ2, SupportThetaRanksZ2Invariant, SupportAssociatedGradedFailsType, SupportAssociatedGradedFailsRank, SupportF2OverF1RejectedByTypeAndRank, SupportDegreeZeroAbsentByReduction, SupportCrossLaneViolatesType, SupportCrossLaneFailsAlphaShape, SupportThetaUniquenessStrengthensMuB, SupportSelectorAndCrossLaneFollowUnique, SupportMuBTargetRanksFixedByUniqueTheta}
}
func Failures() []string {
	return []string{FailureThetaNotNative, FailureOrderPreservationNotNative, FailureExposureEnclosureUniquenessNative, FailureZ2IndependenceNotNative, FailureF2OverF1RejectionNotNative, FailureAlternativeRejectionNotNative, FailureMuBNotNativeWithoutThetaSource, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}
}

func firewallsOK(f FirewallLedger) bool { return containsAll(f.List(), Failures()) }
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

func FormatOrder(o OrderUniquenessAudit) string {
	return fmt.Sprintf("source=%s target=%s levels=(%d,%d) nontrivial_bijection=%t order_preserving=%t swapped_order_reversing=%t unique_under_order=%t native_order_theorem=%t", o.SourceChain, o.TargetChain, o.SourceLevels, o.TargetLevels, o.NontrivialBijection, o.OrderPreserving, o.SwappedOrderReversing, o.UniqueUnderOrder, o.NativeOrderTheorem)
}
func FormatTypes(t TypeUniquenessAudit) string {
	return fmt.Sprintf("degree_one_type=%s degree_two_type=%s exposure_target=%s enclosure_target=%s exposure_unique=%t enclosure_unique=%t forces_theta=%t native_type_theorem=%t", t.DegreeOneType, t.DegreeTwoType, t.ExposureTarget, t.EnclosureTarget, t.ExposureUnique, t.EnclosureUnique, t.ForcesThetaAssignment, t.NativeTypeTheorem)
}
func FormatZ2(z Z2IndependenceAudit) string {
	return fmt.Sprintf("lambda=(%s;%s) barlambda=(%s;%s) ranks=(%d,%d) representative_independent=%t commutes_with_flip=%t native_z2_theorem=%t", z.LambdaOne, z.LambdaTwo, z.BarOne, z.BarTwo, z.RankOne, z.RankTwo, z.RepresentativeIndependent, z.CommutesWithFlip, z.NativeZ2Theorem)
}
func FormatGraded(g AssociatedGradedRejectionAudit) string {
	return fmt.Sprintf("alt=%s alt_rank=%d required=%s required_rank=%d fails_type=%t fails_rank=%t rejected_by_type_and_rank=%t native_cumulative_theorem=%t", g.AlternativeTarget, g.AlternativeRank, g.RequiredTarget, g.RequiredRank, g.FailsCumulativeType, g.FailsAlphaRank, g.RejectedByTypeAndRank, g.NativeCumulativeTheorem)
}
func FormatAlternatives(a AlternativeRejectionAudit) string {
	return fmt.Sprintf("degree_zero_absent=%t cross_lanes=(%t,%t) false_terms=(%s,%s) violates_type=%t fails_alpha_shape=%t native_uniqueness=%t", a.DegreeZeroAbsent, a.CrossLaneExposureToEnclose, a.CrossLaneEnclosureToExpose, a.FalseLinearTerm, a.FalseQuadraticTerm, a.ViolatesType, a.FailsAlphaShape, a.NativeUniquenessTheorem)
}
func FormatMeasure(m MeasureConsequenceAudit) string {
	return fmt.Sprintf("formula=%s ranks=(%d,%d) chambers=(%d,%d) alpha=%s selector_functionhood=%t cross_lane_exclusion=%t target_ranks_fixed=%t native_measure=%t", m.MeasureFormula, m.ThetaRankOne, m.ThetaRankTwo, m.H10Rank, m.H72Rank, m.AlphaFormula, m.SelectorFunctionhood, m.CrossLaneExclusion, m.TargetRanksFixed, m.NativeMeasure)
}
func FormatFirewalls(f FirewallLedger) string { return strings.Join(f.List(), ";") }
