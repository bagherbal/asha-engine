// Package generation2boundarydegreetoairlockflagtargetfunctoraudit implements
// Gate 925: BoundaryDegree-to-AirlockFlag TargetFunctor Audit.
//
// Gate 925 follows Gate 924's result that Lambda^1 B_2 has native exterior
// exposure shape and Lambda^2 B_2 has native exterior enclosure shape. It
// audits whether this two-level boundary-degree chain can be transported to the
// two-level Z2 puncture-airlock flag quotient chain as a bridge-level target
// functor Theta_B^Z2. The gate supports the order-preserving target-functor
// shape while preserving the firewall that no native theorem yet forces that
// functor.
package generation2boundarydegreetoairlockflagtargetfunctoraudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE925-BOUNDARYDEGREE-TO-AIRLOCKFLAG-TARGETFUNCTOR-AUDIT"

	Gate924ShortStatus = "R3_ALPHA_EXPOSURE_ENCLOSURE_NATIVE_SHAPE_TARGET_FUNCTOR_BLOCKED"

	BoundaryDegreeChain = "deg(Lambda^1 B_2)<deg(Lambda^2 B_2)"
	AirlockFlagChain    = "F_0 subset F_1 subset F_2"
	Z2PunctureClass     = "[p]_{Z2}={e_lambda tensor P_1,e_barlambda tensor P_1}"
	ThetaFunctor        = "Theta_B^Z2:{1,2}->{[F_1/F_0]_{Z2},[F_2/F_0]_{Z2}}"
	ThetaOne            = "Theta_B^Z2(1)=[F_1/F_0]_{Z2}"
	ThetaTwo            = "Theta_B^Z2(2)=[F_2/F_0]_{Z2}"
	ExposureTarget      = "exposure -> [F_1/F_0]_{Z2}"
	EnclosureTarget     = "enclosure -> [F_2/F_0]_{Z2}"
	MeasureFormula      = "mu_B(R_B(S_split))=rank(Theta_B^Z2(1))/rank(H_10)*S_split+rank(Theta_B^Z2(2))/rank(H_72)*S_split^2"
	AlphaFormula        = "alpha_B=(3/10)S_split+(7/72)S_split^2"
	NextGate            = "NEXT_PRESSURE_GATE926_TARGETFUNCTOR_NATURALITY_UNIQUENESS_AUDIT"

	RankF1OverF0 = 3
	RankF2OverF0 = 7
	RankF2OverF1 = 4
	RankH10      = 10
	RankH72      = 72

	Classification = "R3_DEGREE_TO_AIRLOCK_FLAG_TARGET_FUNCTOR_SUPPORTED_NOT_NATIVE"
	ShortStatus    = "R3_ALPHA_TARGET_FUNCTOR_SHAPE_SUPPORTED_NATIVE_THETA_MISSING"
	FinalTruth     = "BOUNDARY_DEGREE_TO_AIRLOCK_FLAG_TARGET_FUNCTOR_SHAPE_SUPPORTED_BUT_NATIVE_FUNCTOR_THEOREM_MISSING"

	StatusInheritedGate924          = "PASS_GATE924_NATIVE_EXTERIOR_EXPOSURE_ENCLOSURE_SHAPE_INHERITED"
	StatusTwoLevelOrderMatch        = "PASS_SOURCE_AND_TARGET_HAVE_MATCHING_TWO_LEVEL_ORDER_TYPE"
	StatusExposureTargetShape       = "PASS_EXPOSURE_TARGETS_FIRST_NONBASE_FLAG_QUOTIENT"
	StatusEnclosureTargetShape      = "PASS_ENCLOSURE_TARGETS_CUMULATIVE_FULL_FLAG_QUOTIENT"
	StatusAssociatedGradedRejected  = "PASS_ASSOCIATED_GRADED_SLICE_REJECTED_FOR_TOP_BOUNDARY_DEGREE"
	StatusThetaFunctorShapeDefined  = "PASS_THETA_B_Z2_TARGET_FUNCTOR_SHAPE_DEFINED"
	StatusSelectorFunctionhood      = "PASS_THETA_B_Z2_SUPPLIES_SELECTOR_FUNCTIONHOOD_SHAPE"
	StatusMuBTargetRanks            = "PASS_THETA_B_Z2_SUPPLIES_BOUNDARY_MEASURE_TARGET_RANKS"
	StatusNativeThetaMissing        = "FIREWALL_PRESERVED_NATIVE_THETA_B_Z2_THEOREM_MISSING"
	StatusNativeAlphaR3StillBlocked = "FIREWALL_PRESERVED_ALPHA_AND_R3_NOT_NATIVE"

	SupportSourceDegreesTwoLevel              = "CONDITIONAL_SUPPORT_SOURCE_BOUNDARY_DEGREES_FORM_TWO_LEVEL_ACTIVE_CHAIN"
	SupportTargetFlagTwoNonbaseLevels         = "CONDITIONAL_SUPPORT_TARGET_AIRLOCK_FLAG_HAS_TWO_NONBASE_QUOTIENT_LEVELS"
	SupportDegreeAndFlagMatchingOrderType     = "CONDITIONAL_SUPPORT_DEGREE_CHAIN_AND_FLAG_CHAIN_HAVE_MATCHING_ORDER_TYPE"
	SupportExposureTargetsFirstQuotient       = "CONDITIONAL_SUPPORT_EXPOSURE_TARGETS_FIRST_NONBASE_AIRLOCK_FLAG_QUOTIENT"
	SupportDegreeOneMapsByMinimalExposure     = "CONDITIONAL_SUPPORT_DEGREE_ONE_MAPS_TO_F1_OVER_F0_BY_MINIMAL_EXPOSURE_LEVEL"
	SupportThetaOneExposedFace                = "CONDITIONAL_SUPPORT_THETA_B_Z2_OF_ONE_EQUALS_EXPOSED_FACE_CLASS"
	SupportEnclosureTargetsFullQuotient       = "CONDITIONAL_SUPPORT_ENCLOSURE_TARGETS_CUMULATIVE_FULL_AIRLOCK_FLAG_QUOTIENT"
	SupportDegreeTwoMapsByFullEnclosure       = "CONDITIONAL_SUPPORT_DEGREE_TWO_MAPS_TO_F2_OVER_F0_BY_FULL_PAIR_ENCLOSURE_LEVEL"
	SupportThetaTwoFullEnclosure              = "CONDITIONAL_SUPPORT_THETA_B_Z2_OF_TWO_EQUALS_FULL_ENCLOSURE_CLASS"
	SupportAssociatedGradedRejected           = "CONDITIONAL_SUPPORT_ASSOCIATED_GRADED_SLICE_F2_OVER_F1_REJECTED"
	SupportTopDegreeSelectsCumulative         = "CONDITIONAL_SUPPORT_TOP_BOUNDARY_DEGREE_SELECTS_CUMULATIVE_QUOTIENT_NOT_INCREMENTAL_SLICE"
	SupportFullPairActivationRequiresF2OverF0 = "CONDITIONAL_SUPPORT_FULL_PAIR_ACTIVATION_REQUIRES_F2_OVER_F0"
	SupportThetaShapeDefined                  = "CONDITIONAL_SUPPORT_BOUNDARY_DEGREE_TO_AIRLOCK_FLAG_TARGET_FUNCTOR_SHAPE_DEFINED"
	SupportThetaOrderPreserving               = "CONDITIONAL_SUPPORT_THETA_B_Z2_IS_ORDER_PRESERVING"
	SupportThetaZ2RepresentativeIndependent   = "CONDITIONAL_SUPPORT_THETA_B_Z2_IS_Z2_REPRESENTATIVE_INDEPENDENT"
	SupportThetaExposureEnclosureTyped        = "CONDITIONAL_SUPPORT_THETA_B_Z2_IS_EXPOSURE_ENCLOSURE_TYPED"
	SupportThetaCumulativeTopDegree           = "CONDITIONAL_SUPPORT_THETA_B_Z2_IS_CUMULATIVE_AT_TOP_DEGREE"
	SupportThetaSuppliesSelectorFunctionhood  = "CONDITIONAL_SUPPORT_THETA_B_Z2_SUPPLIES_SELECTOR_FUNCTIONHOOD"
	SupportIBZ2EqualsTheta                    = "CONDITIONAL_SUPPORT_I_B_Z2_EQUALS_THETA_B_Z2"
	SupportCrossLaneFollowsFromTheta          = "CONDITIONAL_SUPPORT_CROSS_LANE_EXCLUSION_FOLLOWS_FROM_THETA_B_Z2_FUNCTIONHOOD"
	SupportThetaSuppliesMuBRanks              = "CONDITIONAL_SUPPORT_THETA_B_Z2_SUPPLIES_BOUNDARY_ACTIVATION_MEASURE_TARGET_RANKS"
	SupportAlphaReconstructedGivenTheta       = "CONDITIONAL_SUPPORT_ALPHA_B_RECONSTRUCTED_GIVEN_THETA_B_Z2_TARGET_FUNCTOR"
	SupportMuBNativeGapReducedToTheta         = "CONDITIONAL_SUPPORT_BOUNDARY_MEASURE_NATIVE_GAP_REDUCED_TO_NATIVE_THETA_B_Z2_THEOREM"

	FailureOrderTypeNotNative            = "FAILED_ROUTE_ORDER_TYPE_MATCH_DOES_NOT_BY_ITSELF_CERTIFY_NATIVE_TARGET_FUNCTOR"
	FailureMinimalExposureLevelNotNative = "FAILED_ROUTE_MINIMAL_EXPOSURE_LEVEL_RULE_NOT_NATIVE_TARGET_FUNCTOR_THEOREM"
	FailureFullEnclosureLevelNotNative   = "FAILED_ROUTE_FULL_ENCLOSURE_LEVEL_RULE_NOT_NATIVE_TARGET_FUNCTOR_THEOREM"
	FailureCumulativeOverGradedNotNative = "FAILED_ROUTE_CUMULATIVE_OVER_ASSOCIATED_GRADED_RULE_NOT_NATIVE_THEOREM"
	FailureThetaShapeNotNative           = "FAILED_ROUTE_THETA_B_Z2_FUNCTOR_SHAPE_NOT_NATIVE_FUNCTOR_THEOREM"
	FailureSelectorNonNativeWithoutTheta = "FAILED_ROUTE_SELECTOR_FUNCTIONHOOD_REMAINS_NON_NATIVE_WITHOUT_NATIVE_THETA_B_Z2"
	FailureMuBNotNativeWithoutTheta      = "FAILED_ROUTE_MU_B_STILL_NOT_NATIVE_WITHOUT_NATIVE_THETA_B_Z2_THEOREM"
	FailureAlphaBridgeCandidateNotNative = "FAILED_ROUTE_ALPHA_B_REMAINS_BRIDGE_CANDIDATE_NOT_NATIVE"
	FailureNotNativeR3                   = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked     = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap        = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap        = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues      = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator        = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type ChainMatchAudit struct {
	SourceChain         string
	TargetChain         string
	SourceLevels        int
	TargetLevels        int
	MatchingOrderType   bool
	NativeTargetFunctor bool
	Supports            []string
	Failures            []string
}

type TargetLaneAudit struct {
	Degree       int
	BoundaryType string
	Target       string
	Rank         int
	Rule         string
	Z2ClassLevel bool
	NativeRule   bool
	Supports     []string
	Failures     []string
}

type AssociatedGradedAudit struct {
	RejectedTarget       string
	RejectedRank         int
	CumulativeTarget     string
	CumulativeRank       int
	TopDegreeCumulative  bool
	NativeCumulativeRule bool
	Supports             []string
	Failures             []string
}

type ThetaFunctorAudit struct {
	Formula                     string
	OrderPreserving             bool
	Z2RepresentativeIndependent bool
	ExposureEnclosureTyped      bool
	CumulativeTopDegree         bool
	NativeFunctor               bool
	Supports                    []string
	Failures                    []string
}

type SelectorConsequenceAudit struct {
	SelectorFormula            string
	IBEqualsTheta              bool
	CrossLaneExcluded          bool
	NativeSelectorFunctionhood bool
	Supports                   []string
	Failures                   []string
}

type MeasureConsequenceAudit struct {
	MeasureFormula      string
	ThetaRankOne        int
	ThetaRankTwo        int
	H10Rank             int
	H72Rank             int
	AlphaFormula        string
	TargetRanksSupplied bool
	NativeMeasure       bool
	Supports            []string
	Failures            []string
}

type FirewallLedger struct {
	OrderTypeNotNative            bool
	MinimalExposureLevelNotNative bool
	FullEnclosureLevelNotNative   bool
	CumulativeOverGradedNotNative bool
	ThetaShapeNotNative           bool
	SelectorNonNativeWithoutTheta bool
	MuBNotNativeWithoutTheta      bool
	AlphaBridgeCandidateNotNative bool
	NotNativeR3                   bool
	FullAFDescentStillBlocked     bool
	NoGenerationCarrierMap        bool
	NoFlavorOrientationMap        bool
	NoIndividualYukawaValues      bool
	NoNativeYukawaOperator        bool
}

func (f FirewallLedger) List() []string {
	var out []string
	if f.OrderTypeNotNative {
		out = append(out, FailureOrderTypeNotNative)
	}
	if f.MinimalExposureLevelNotNative {
		out = append(out, FailureMinimalExposureLevelNotNative)
	}
	if f.FullEnclosureLevelNotNative {
		out = append(out, FailureFullEnclosureLevelNotNative)
	}
	if f.CumulativeOverGradedNotNative {
		out = append(out, FailureCumulativeOverGradedNotNative)
	}
	if f.ThetaShapeNotNative {
		out = append(out, FailureThetaShapeNotNative)
	}
	if f.SelectorNonNativeWithoutTheta {
		out = append(out, FailureSelectorNonNativeWithoutTheta)
	}
	if f.MuBNotNativeWithoutTheta {
		out = append(out, FailureMuBNotNativeWithoutTheta)
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
	ChainMatch      ChainMatchAudit
	ExposureTarget  TargetLaneAudit
	EnclosureTarget TargetLaneAudit
	Graded          AssociatedGradedAudit
	Theta           ThetaFunctorAudit
	Selector        SelectorConsequenceAudit
	Measure         MeasureConsequenceAudit
	Firewalls       FirewallLedger
	Final           string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		ID:              AuditID,
		InheritedStatus: Gate924ShortStatus,
		Truth:           FinalTruth,
		Classification:  Classification,
		ShortStatus:     ShortStatus,
		ChainMatch: ChainMatchAudit{
			SourceChain:         BoundaryDegreeChain,
			TargetChain:         AirlockFlagChain,
			SourceLevels:        2,
			TargetLevels:        2,
			MatchingOrderType:   true,
			NativeTargetFunctor: false,
			Supports:            []string{SupportSourceDegreesTwoLevel, SupportTargetFlagTwoNonbaseLevels, SupportDegreeAndFlagMatchingOrderType},
			Failures:            []string{FailureOrderTypeNotNative},
		},
		ExposureTarget: TargetLaneAudit{
			Degree:       1,
			BoundaryType: "one-factor boundary exposure",
			Target:       "[F_1/F_0]_{Z2}",
			Rank:         RankF1OverF0,
			Rule:         "minimal exposure level -> first non-base airlock flag quotient",
			Z2ClassLevel: true,
			NativeRule:   false,
			Supports:     []string{SupportExposureTargetsFirstQuotient, SupportDegreeOneMapsByMinimalExposure, SupportThetaOneExposedFace},
			Failures:     []string{FailureMinimalExposureLevelNotNative},
		},
		EnclosureTarget: TargetLaneAudit{
			Degree:       2,
			BoundaryType: "two-factor full boundary-pair enclosure",
			Target:       "[F_2/F_0]_{Z2}",
			Rank:         RankF2OverF0,
			Rule:         "full pair enclosure level -> cumulative full airlock flag quotient",
			Z2ClassLevel: true,
			NativeRule:   false,
			Supports:     []string{SupportEnclosureTargetsFullQuotient, SupportDegreeTwoMapsByFullEnclosure, SupportThetaTwoFullEnclosure},
			Failures:     []string{FailureFullEnclosureLevelNotNative},
		},
		Graded: AssociatedGradedAudit{
			RejectedTarget:       "F_2/F_1",
			RejectedRank:         RankF2OverF1,
			CumulativeTarget:     "[F_2/F_0]_{Z2}",
			CumulativeRank:       RankF2OverF0,
			TopDegreeCumulative:  true,
			NativeCumulativeRule: false,
			Supports:             []string{SupportAssociatedGradedRejected, SupportTopDegreeSelectsCumulative, SupportFullPairActivationRequiresF2OverF0},
			Failures:             []string{FailureCumulativeOverGradedNotNative},
		},
		Theta: ThetaFunctorAudit{
			Formula:                     ThetaFunctor + "; " + ThetaOne + "; " + ThetaTwo,
			OrderPreserving:             true,
			Z2RepresentativeIndependent: true,
			ExposureEnclosureTyped:      true,
			CumulativeTopDegree:         true,
			NativeFunctor:               false,
			Supports:                    []string{SupportThetaShapeDefined, SupportThetaOrderPreserving, SupportThetaZ2RepresentativeIndependent, SupportThetaExposureEnclosureTyped, SupportThetaCumulativeTopDegree},
			Failures:                    []string{FailureThetaShapeNotNative},
		},
		Selector: SelectorConsequenceAudit{
			SelectorFormula:            "I_B^Z2(k)=Theta_B^Z2(k)",
			IBEqualsTheta:              true,
			CrossLaneExcluded:          true,
			NativeSelectorFunctionhood: false,
			Supports:                   []string{SupportThetaSuppliesSelectorFunctionhood, SupportIBZ2EqualsTheta, SupportCrossLaneFollowsFromTheta},
			Failures:                   []string{FailureSelectorNonNativeWithoutTheta},
		},
		Measure: MeasureConsequenceAudit{
			MeasureFormula:      MeasureFormula,
			ThetaRankOne:        RankF1OverF0,
			ThetaRankTwo:        RankF2OverF0,
			H10Rank:             RankH10,
			H72Rank:             RankH72,
			AlphaFormula:        AlphaFormula,
			TargetRanksSupplied: true,
			NativeMeasure:       false,
			Supports:            []string{SupportThetaSuppliesMuBRanks, SupportAlphaReconstructedGivenTheta, SupportMuBNativeGapReducedToTheta},
			Failures:            []string{FailureMuBNotNativeWithoutTheta, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3},
		},
		Firewalls: FirewallLedger{
			OrderTypeNotNative:            true,
			MinimalExposureLevelNotNative: true,
			FullEnclosureLevelNotNative:   true,
			CumulativeOverGradedNotNative: true,
			ThetaShapeNotNative:           true,
			SelectorNonNativeWithoutTheta: true,
			MuBNotNativeWithoutTheta:      true,
			AlphaBridgeCandidateNotNative: true,
			NotNativeR3:                   true,
			FullAFDescentStillBlocked:     true,
			NoGenerationCarrierMap:        true,
			NoFlavorOrientationMap:        true,
			NoIndividualYukawaValues:      true,
			NoNativeYukawaOperator:        true,
		},
		Final: "Gate 925 supports Theta_B^Z2 as an order-preserving Z2 class target-functor shape from the active boundary-degree chain to the puncture-airlock flag quotient chain, but no native theorem yet forces that target functor.",
	}
	if err := validate(a); err != nil {
		return Audit{}, err
	}
	return a, nil
}

func validate(a Audit) error {
	if a.InheritedStatus != Gate924ShortStatus {
		return fmt.Errorf("unexpected inherited status: %s", a.InheritedStatus)
	}
	if !a.ChainMatch.MatchingOrderType || a.ChainMatch.NativeTargetFunctor || a.ChainMatch.SourceLevels != 2 || a.ChainMatch.TargetLevels != 2 {
		return fmt.Errorf("bad chain match audit: %s", FormatChain(a.ChainMatch))
	}
	if a.ExposureTarget.Degree != 1 || a.ExposureTarget.Rank != RankF1OverF0 || !a.ExposureTarget.Z2ClassLevel || a.ExposureTarget.NativeRule {
		return fmt.Errorf("bad exposure target: %s", FormatTarget(a.ExposureTarget))
	}
	if a.EnclosureTarget.Degree != 2 || a.EnclosureTarget.Rank != RankF2OverF0 || !a.EnclosureTarget.Z2ClassLevel || a.EnclosureTarget.NativeRule {
		return fmt.Errorf("bad enclosure target: %s", FormatTarget(a.EnclosureTarget))
	}
	if !a.Graded.TopDegreeCumulative || a.Graded.NativeCumulativeRule || a.Graded.RejectedRank != RankF2OverF1 || a.Graded.CumulativeRank != RankF2OverF0 {
		return fmt.Errorf("bad associated-graded audit: %s", FormatGraded(a.Graded))
	}
	if !a.Theta.OrderPreserving || !a.Theta.Z2RepresentativeIndependent || !a.Theta.ExposureEnclosureTyped || !a.Theta.CumulativeTopDegree || a.Theta.NativeFunctor {
		return fmt.Errorf("bad theta audit: %s", FormatTheta(a.Theta))
	}
	if !a.Selector.IBEqualsTheta || !a.Selector.CrossLaneExcluded || a.Selector.NativeSelectorFunctionhood {
		return fmt.Errorf("bad selector consequence: %s", FormatSelector(a.Selector))
	}
	if !a.Measure.TargetRanksSupplied || a.Measure.NativeMeasure || a.Measure.ThetaRankOne != RankF1OverF0 || a.Measure.ThetaRankTwo != RankF2OverF0 || a.Measure.H10Rank != RankH10 || a.Measure.H72Rank != RankH72 {
		return fmt.Errorf("bad measure consequence: %s", FormatMeasure(a.Measure))
	}
	if !firewallsOK(a.Firewalls) {
		return fmt.Errorf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	return nil
}

func Statuses() []string {
	return []string{StatusInheritedGate924, StatusTwoLevelOrderMatch, StatusExposureTargetShape, StatusEnclosureTargetShape, StatusAssociatedGradedRejected, StatusThetaFunctorShapeDefined, StatusSelectorFunctionhood, StatusMuBTargetRanks, StatusNativeThetaMissing, StatusNativeAlphaR3StillBlocked}
}

func Supports() []string {
	return []string{SupportSourceDegreesTwoLevel, SupportTargetFlagTwoNonbaseLevels, SupportDegreeAndFlagMatchingOrderType, SupportExposureTargetsFirstQuotient, SupportDegreeOneMapsByMinimalExposure, SupportThetaOneExposedFace, SupportEnclosureTargetsFullQuotient, SupportDegreeTwoMapsByFullEnclosure, SupportThetaTwoFullEnclosure, SupportAssociatedGradedRejected, SupportTopDegreeSelectsCumulative, SupportFullPairActivationRequiresF2OverF0, SupportThetaShapeDefined, SupportThetaOrderPreserving, SupportThetaZ2RepresentativeIndependent, SupportThetaExposureEnclosureTyped, SupportThetaCumulativeTopDegree, SupportThetaSuppliesSelectorFunctionhood, SupportIBZ2EqualsTheta, SupportCrossLaneFollowsFromTheta, SupportThetaSuppliesMuBRanks, SupportAlphaReconstructedGivenTheta, SupportMuBNativeGapReducedToTheta}
}

func Failures() []string {
	return []string{FailureOrderTypeNotNative, FailureMinimalExposureLevelNotNative, FailureFullEnclosureLevelNotNative, FailureCumulativeOverGradedNotNative, FailureThetaShapeNotNative, FailureSelectorNonNativeWithoutTheta, FailureMuBNotNativeWithoutTheta, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}
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

func FormatChain(c ChainMatchAudit) string {
	return fmt.Sprintf("source=%s target=%s source_levels=%d target_levels=%d matching_order_type=%t native_target_functor=%t", c.SourceChain, c.TargetChain, c.SourceLevels, c.TargetLevels, c.MatchingOrderType, c.NativeTargetFunctor)
}
func FormatTarget(t TargetLaneAudit) string {
	return fmt.Sprintf("degree=%d boundary_type=%s target=%s rank=%d rule=%s z2_class=%t native_rule=%t", t.Degree, t.BoundaryType, t.Target, t.Rank, t.Rule, t.Z2ClassLevel, t.NativeRule)
}
func FormatGraded(g AssociatedGradedAudit) string {
	return fmt.Sprintf("rejected=%s rejected_rank=%d cumulative=%s cumulative_rank=%d top_degree_cumulative=%t native_cumulative_rule=%t", g.RejectedTarget, g.RejectedRank, g.CumulativeTarget, g.CumulativeRank, g.TopDegreeCumulative, g.NativeCumulativeRule)
}
func FormatTheta(t ThetaFunctorAudit) string {
	return fmt.Sprintf("formula=%s order_preserving=%t z2_independent=%t exposure_enclosure_typed=%t cumulative_top=%t native_functor=%t", t.Formula, t.OrderPreserving, t.Z2RepresentativeIndependent, t.ExposureEnclosureTyped, t.CumulativeTopDegree, t.NativeFunctor)
}
func FormatSelector(s SelectorConsequenceAudit) string {
	return fmt.Sprintf("selector=%s ib_equals_theta=%t cross_lane_excluded=%t native_selector_functionhood=%t", s.SelectorFormula, s.IBEqualsTheta, s.CrossLaneExcluded, s.NativeSelectorFunctionhood)
}
func FormatMeasure(m MeasureConsequenceAudit) string {
	return fmt.Sprintf("formula=%s ranks=(%d,%d) chambers=(%d,%d) alpha=%s target_ranks_supplied=%t native_measure=%t", m.MeasureFormula, m.ThetaRankOne, m.ThetaRankTwo, m.H10Rank, m.H72Rank, m.AlphaFormula, m.TargetRanksSupplied, m.NativeMeasure)
}
func FormatFirewalls(f FirewallLedger) string { return strings.Join(f.List(), ";") }
