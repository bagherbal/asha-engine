// Package generation2degreeindexedselectorfunctionhoodsourceaudit implements
// Gate 923: DegreeIndexed Selector Functionhood Source Audit.
//
// Gate 923 follows Gate 922's constraint-source audit and focuses on the
// primary remaining native gap for the BoundaryActivationMeasure: why exterior
// degree k should select the cumulative Z2 airlock quotient [F_k/F_0]_{Z2}.
// It source-types the selector by the exposure/enclosure interpretation while
// preserving the firewall that this is still bridge-level typing, not a native
// selector theorem.
package generation2degreeindexedselectorfunctionhoodsourceaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE923-DEGREEINDEXED-SELECTOR-FUNCTIONHOOD-SOURCE-AUDIT"

	Gate922ShortStatus = "R3_ALPHA_MEASURE_CONSTRAINTS_PARTLY_SOURCED_SELECTOR_STILL_MISSING"

	SBoundary = 0.0012924448188162962
	AlphaB    = 0.0003878958469680527

	RankExposureFace    = 3
	RankFullEnclosure   = 7
	RankAssociatedSlice = 4
	RankH10             = 10
	RankH72             = 72

	AlphaLinear = 0.00038773344564488885
	AlphaQuad   = 0.0000001624013231638281

	PunctureClass         = "[p]_{Z2}={e_lambda tensor P_1,e_barlambda tensor P_1}"
	BoundaryResponse      = "R_B(s)=s(b1+b2)+s^2(b1 wedge b2)"
	SelectorFormula       = "I_B^Z2(k)=[F_k/F_0]_{Z2}"
	SelectorOneFormula    = "I_B^Z2(1)=[F_1/F_0]_{Z2}"
	SelectorTwoFormula    = "I_B^Z2(2)=[F_2/F_0]_{Z2}"
	MuBFormula            = "mu_B(R_B(S_split))=rank(I_B^Z2(1))/rank(H_10)*S_split+rank(I_B^Z2(2))/rank(H_72)*S_split^2"
	BoundaryAlphaFormula  = "alpha_B^Z2=(3/10)S_split+(7/72)S_split^2"
	ExposureEnclosureRule = "Lambda^1 B_2=single-boundary exposure; Lambda^2 B_2=full boundary-pair enclosure"
	NextGate              = "NEXT_PRESSURE_GATE924_BOUNDARYDEGREE_EXPOSUREENCLOSURE_FUNCTOR_AUDIT"
	StrategicConclusion   = "Gate 923 weakens the selector wound: selector functionhood is source-typed by exposure/enclosure, but exposure/enclosure itself is not yet native. The next pressure point is whether Lambda^1 B_2 is natively single-boundary exposure and Lambda^2 B_2 is natively full boundary-pair enclosure."

	SourceBridgeTypedNotNative = "BRIDGE_TYPED_NOT_NATIVE"
	SourceDependent            = "DEPENDENT_ON_SELECTOR_FUNCTIONHOOD"
	SourceZ2Compatible         = "Z2_CLASS_COMPATIBLE_NOT_NATIVE"
	PrimaryGapExposureFunctor  = "PRIMARY_GAP_EXPOSURE_ENCLOSURE_FUNCTOR"

	Classification = "R3_SELECTOR_FUNCTIONHOOD_SOURCE_TYPED_NOT_NATIVE"
	ShortStatus    = "R3_ALPHA_SELECTOR_GAP_WEAKENED_TO_EXPOSURE_ENCLOSURE_FUNCTOR"
	FinalTruth     = "DEGREE_INDEXED_SELECTOR_FUNCTIONHOOD_SOURCE_TYPED_BY_EXPOSURE_ENCLOSURE_BUT_NATIVE_SELECTOR_THEOREM_MISSING"

	StatusInheritedGate922      = "PASS_GATE922_SELECTOR_PRIMARY_GAP_INHERITED"
	StatusDegreeOneSourceTyped  = "PASS_DEGREE_ONE_SOURCE_TYPED_AS_SINGLE_BOUNDARY_EXPOSURE"
	StatusDegreeTwoSourceTyped  = "PASS_DEGREE_TWO_SOURCE_TYPED_AS_FULL_BOUNDARY_PAIR_ENCLOSURE"
	StatusCumulativeEnclosure   = "PASS_DEGREE_TWO_CUMULATIVE_ENCLOSURE_OVER_ASSOCIATED_GRADED_SLICE"
	StatusFunctionhoodSupported = "PASS_SELECTOR_FUNCTIONHOOD_SUPPORTED_IF_EXPOSURE_ENCLOSURE_ACCEPTED"
	StatusCrossLaneConsequence  = "PASS_CROSS_LANE_EXCLUSION_FOLLOWS_FROM_SELECTOR_FUNCTIONHOOD"
	StatusZ2Compatible          = "PASS_SELECTOR_FUNCTIONHOOD_Z2_CLASS_COMPATIBLE"
	StatusMuBGapReduced         = "PASS_SELECTOR_SUPPLIES_BOUNDARY_ACTIVATION_MEASURE_TARGET_RANKS"
	StatusNativeSelectorMissing = "FIREWALL_PRESERVED_NATIVE_SELECTOR_THEOREM_MISSING"
	StatusNativeR3Blocked       = "FIREWALL_PRESERVED_ALPHA_AND_R3_NOT_NATIVE"

	SupportDegreeOneExposureSource            = "CONDITIONAL_SUPPORT_DEGREE_ONE_HAS_SINGLE_BOUNDARY_EXPOSURE_SOURCE"
	SupportSingleExposureTargetsFaceQuotient  = "CONDITIONAL_SUPPORT_SINGLE_BOUNDARY_EXPOSURE_TARGETS_EXPOSED_PHASE_FACE_QUOTIENT"
	SupportIBZ2OneExposedFace                 = "CONDITIONAL_SUPPORT_I_B_Z2_OF_ONE_EQUALS_EXPOSED_FACE_CLASS"
	SupportDegreeTwoEnclosureSource           = "CONDITIONAL_SUPPORT_DEGREE_TWO_HAS_FULL_BOUNDARY_PAIR_ENCLOSURE_SOURCE"
	SupportFullEnclosureTargetsComplement     = "CONDITIONAL_SUPPORT_FULL_BOUNDARY_PAIR_ENCLOSURE_TARGETS_FULL_PUNCTURE_COMPLEMENT_QUOTIENT"
	SupportIBZ2TwoFullEnclosure               = "CONDITIONAL_SUPPORT_I_B_Z2_OF_TWO_EQUALS_FULL_ENCLOSURE_CLASS"
	SupportDegreeTwoCumulativeNotGraded       = "CONDITIONAL_SUPPORT_DEGREE_TWO_IS_CUMULATIVE_ENCLOSURE_NOT_ASSOCIATED_GRADED_SLICE"
	SupportFullPairRequiresF2OverF0           = "CONDITIONAL_SUPPORT_FULL_PAIR_ENCLOSURE_REQUIRES_F2_OVER_F0"
	SupportF2OverF1Rejected                   = "CONDITIONAL_SUPPORT_F2_OVER_F1_REJECTED_AS_ALPHA_SELECTOR_TARGET"
	SupportExposureEnclosureGivesFunctionhood = "CONDITIONAL_SUPPORT_EXPOSURE_ENCLOSURE_TYPING_GIVES_SELECTOR_FUNCTIONHOOD"
	SupportEachDegreeUniqueTargetType         = "CONDITIONAL_SUPPORT_EACH_BOUNDARY_DEGREE_HAS_UNIQUE_TARGET_TYPE"
	SupportIBZ2FunctionalIfAccepted           = "CONDITIONAL_SUPPORT_I_B_Z2_IS_FUNCTIONAL_IF_EXPOSURE_ENCLOSURE_SELECTOR_IS_ACCEPTED"
	SupportCrossLaneFollowsFunctionhood       = "CONDITIONAL_SUPPORT_CROSS_LANE_EXCLUSION_FOLLOWS_FROM_SELECTOR_FUNCTIONHOOD"
	SupportFalseTermsBlocked                  = "CONDITIONAL_SUPPORT_FALSE_ALPHA_TERMS_BLOCKED_IF_I_B_Z2_IS_FUNCTIONAL"
	SupportSelectorFunctionhoodZ2Compatible   = "CONDITIONAL_SUPPORT_SELECTOR_FUNCTIONHOOD_IS_Z2_CLASS_COMPATIBLE"
	SupportIBZ2CommutesWithPhaseFlip          = "CONDITIONAL_SUPPORT_I_B_Z2_COMMUTES_WITH_PHASE_REPRESENTATIVE_FLIP"
	SupportSelectorRanksRepresentativeFree    = "CONDITIONAL_SUPPORT_SELECTOR_TARGET_RANKS_ARE_REPRESENTATIVE_INDEPENDENT"
	SupportSelectorSuppliesMuBRanks           = "CONDITIONAL_SUPPORT_SELECTOR_FUNCTIONHOOD_SUPPLIES_MU_B_TARGET_RANKS"
	SupportMuBNativeGapReduced                = "CONDITIONAL_SUPPORT_BOUNDARY_ACTIVATION_MEASURE_NATIVE_GAP_REDUCED_BY_SELECTOR_SOURCE_TYPING"

	FailureNoNativeSelectorFunctionhood      = "FAILED_ROUTE_NO_NATIVE_SELECTOR_FUNCTIONHOOD_THEOREM"
	FailureNoNativeDegreeToZ2FlagFunctor     = "FAILED_ROUTE_NO_NATIVE_DEGREE_TO_Z2_FLAG_CLASS_FUNCTOR"
	FailureExposureToF1NotNative             = "FAILED_ROUTE_EXPOSURE_TO_F1_OVER_F0_NOT_NATIVE_SELECTOR_THEOREM"
	FailureEnclosureToF2NotNative            = "FAILED_ROUTE_ENCLOSURE_TO_F2_OVER_F0_NOT_NATIVE_SELECTOR_THEOREM"
	FailureNoNativeCumulativeTheorem         = "FAILED_ROUTE_NO_NATIVE_THEOREM_FOR_CUMULATIVE_ENCLOSURE_OVER_GRADED_SLICE"
	FailureSelectorDependsOnBridgeRule       = "FAILED_ROUTE_SELECTOR_FUNCTIONHOOD_STILL_DEPENDS_ON_BRIDGE_EXPOSURE_ENCLOSURE_RULE"
	FailureCrossLaneNotNativeWithoutSelector = "FAILED_ROUTE_CROSS_LANE_EXCLUSION_NOT_NATIVE_WITHOUT_NATIVE_SELECTOR_FUNCTIONHOOD"
	FailureZ2CompatibilityNotNativeSelector  = "FAILED_ROUTE_Z2_CLASS_COMPATIBILITY_NOT_NATIVE_SELECTOR_THEOREM"
	FailureMuBStillNotNativeWithoutSelector  = "FAILED_ROUTE_MU_B_STILL_NOT_NATIVE_WITHOUT_NATIVE_SELECTOR_THEOREM"
	FailureAlphaBridgeCandidateNotNative     = "FAILED_ROUTE_ALPHA_B_REMAINS_BRIDGE_CANDIDATE_NOT_NATIVE"
	FailureNotNativeR3                       = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked         = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap            = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap            = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues          = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator            = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type SelectorLaneAudit struct {
	Degree       int
	ResponseTerm string
	SourceType   string
	Target       string
	Rank         int
	SourceStatus string
	BridgeTyped  bool
	Native       bool
	Supports     []string
	Failures     []string
}

type CumulativeEnclosureAudit struct {
	DegreeTwoTarget          string
	AssociatedGradedRejected string
	RankF2OverF0             int
	RankF2OverF1             int
	CumulativeRequired       bool
	Native                   bool
	Supports                 []string
	Failures                 []string
}

type FunctionhoodAudit struct {
	DomainDegrees                []int
	TargetRule                   string
	ExposureEnclosureAccepted    bool
	FunctionalIfExposureAccepted bool
	NativeFunctionhood           bool
	PrimaryGap                   string
	Supports                     []string
	Failures                     []string
}

type CrossLaneAudit struct {
	FalseLinearLane      string
	FalseQuadraticLane   string
	FalseLinearTerm      string
	FalseQuadraticTerm   string
	ExcludedIfFunctional bool
	NativeExclusion      bool
	Supports             []string
	Failures             []string
}

type Z2CompatibilityAudit struct {
	LambdaRepresentativeOne string
	LambdaRepresentativeTwo string
	BarRepresentativeOne    string
	BarRepresentativeTwo    string
	CommutesWithPhaseFlip   bool
	RanksRepresentativeFree bool
	NativeZ2Selector        bool
	Supports                []string
	Failures                []string
}

type MuBContributionAudit struct {
	Formula               string
	AlphaFormula          string
	S                     float64
	Rank1                 int
	Rank2                 int
	RankH10               int
	RankH72               int
	LinearContribution    float64
	QuadraticContribution float64
	Alpha                 float64
	NativeMuB             bool
	Supports              []string
	Failures              []string
}

type FirewallLedger struct {
	NoNativeSelectorFunctionhood      bool
	NoNativeDegreeToZ2FlagFunctor     bool
	ExposureToF1NotNative             bool
	EnclosureToF2NotNative            bool
	NoNativeCumulativeTheorem         bool
	SelectorDependsOnBridgeRule       bool
	CrossLaneNotNativeWithoutSelector bool
	Z2CompatibilityNotNativeSelector  bool
	MuBStillNotNativeWithoutSelector  bool
	AlphaBridgeCandidateNotNative     bool
	NotNativeR3                       bool
	FullAFDescentStillBlocked         bool
	NoGenerationCarrierMap            bool
	NoFlavorOrientationMap            bool
	NoIndividualYukawaValues          bool
	NoNativeYukawaOperator            bool
}

func (f FirewallLedger) List() []string {
	out := []string{}
	if f.NoNativeSelectorFunctionhood {
		out = append(out, FailureNoNativeSelectorFunctionhood)
	}
	if f.NoNativeDegreeToZ2FlagFunctor {
		out = append(out, FailureNoNativeDegreeToZ2FlagFunctor)
	}
	if f.ExposureToF1NotNative {
		out = append(out, FailureExposureToF1NotNative)
	}
	if f.EnclosureToF2NotNative {
		out = append(out, FailureEnclosureToF2NotNative)
	}
	if f.NoNativeCumulativeTheorem {
		out = append(out, FailureNoNativeCumulativeTheorem)
	}
	if f.SelectorDependsOnBridgeRule {
		out = append(out, FailureSelectorDependsOnBridgeRule)
	}
	if f.CrossLaneNotNativeWithoutSelector {
		out = append(out, FailureCrossLaneNotNativeWithoutSelector)
	}
	if f.Z2CompatibilityNotNativeSelector {
		out = append(out, FailureZ2CompatibilityNotNativeSelector)
	}
	if f.MuBStillNotNativeWithoutSelector {
		out = append(out, FailureMuBStillNotNativeWithoutSelector)
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

type Analysis struct {
	ID                  string
	InheritedStatus     string
	Classification      string
	ShortStatus         string
	Truth               string
	DegreeOne           SelectorLaneAudit
	DegreeTwo           SelectorLaneAudit
	CumulativeEnclosure CumulativeEnclosureAudit
	Functionhood        FunctionhoodAudit
	CrossLane           CrossLaneAudit
	Z2Compatibility     Z2CompatibilityAudit
	MuB                 MuBContributionAudit
	Firewalls           FirewallLedger
	Final               string
}

func BuildDefault() (Analysis, error) {
	linear := float64(RankExposureFace) / float64(RankH10) * SBoundary
	quad := float64(RankFullEnclosure) / float64(RankH72) * SBoundary * SBoundary
	alpha := linear + quad
	if !near(linear, AlphaLinear) || !near(quad, AlphaQuad) || !near(alpha, AlphaB) {
		return Analysis{}, fmt.Errorf("alpha reconstruction drift: linear=%.18g quad=%.18g alpha=%.18g", linear, quad, alpha)
	}

	firewalls := FirewallLedger{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}
	return Analysis{
		ID:              AuditID,
		InheritedStatus: Gate922ShortStatus,
		Classification:  Classification,
		ShortStatus:     ShortStatus,
		Truth:           FinalTruth,
		DegreeOne: SelectorLaneAudit{
			Degree: 1, ResponseTerm: "s(b1+b2)", SourceType: "single-boundary exposure", Target: "[F_1/F_0]_{Z2}=exposed phase-face quotient", Rank: RankExposureFace,
			SourceStatus: SourceBridgeTypedNotNative, BridgeTyped: true, Native: false,
			Supports: []string{SupportDegreeOneExposureSource, SupportSingleExposureTargetsFaceQuotient, SupportIBZ2OneExposedFace}, Failures: []string{FailureExposureToF1NotNative},
		},
		DegreeTwo: SelectorLaneAudit{
			Degree: 2, ResponseTerm: "s^2(b1 wedge b2)", SourceType: "full boundary-pair enclosure", Target: "[F_2/F_0]_{Z2}=full puncture-complement quotient", Rank: RankFullEnclosure,
			SourceStatus: SourceBridgeTypedNotNative, BridgeTyped: true, Native: false,
			Supports: []string{SupportDegreeTwoEnclosureSource, SupportFullEnclosureTargetsComplement, SupportIBZ2TwoFullEnclosure}, Failures: []string{FailureEnclosureToF2NotNative},
		},
		CumulativeEnclosure: CumulativeEnclosureAudit{
			DegreeTwoTarget: "[F_2/F_0]_{Z2}", AssociatedGradedRejected: "F_2/F_1", RankF2OverF0: RankFullEnclosure, RankF2OverF1: RankAssociatedSlice, CumulativeRequired: true, Native: false,
			Supports: []string{SupportDegreeTwoCumulativeNotGraded, SupportFullPairRequiresF2OverF0, SupportF2OverF1Rejected}, Failures: []string{FailureNoNativeCumulativeTheorem},
		},
		Functionhood: FunctionhoodAudit{
			DomainDegrees: []int{1, 2}, TargetRule: SelectorFormula, ExposureEnclosureAccepted: true, FunctionalIfExposureAccepted: true, NativeFunctionhood: false, PrimaryGap: PrimaryGapExposureFunctor,
			Supports: []string{SupportExposureEnclosureGivesFunctionhood, SupportEachDegreeUniqueTargetType, SupportIBZ2FunctionalIfAccepted}, Failures: []string{FailureNoNativeSelectorFunctionhood, FailureNoNativeDegreeToZ2FlagFunctor, FailureSelectorDependsOnBridgeRule},
		},
		CrossLane: CrossLaneAudit{
			FalseLinearLane: "Lambda^1 B_2 -> [F_2/F_0]_{Z2}", FalseQuadraticLane: "Lambda^2 B_2 -> [F_1/F_0]_{Z2}", FalseLinearTerm: "(7/72)s", FalseQuadraticTerm: "(3/10)s^2", ExcludedIfFunctional: true, NativeExclusion: false,
			Supports: []string{SupportCrossLaneFollowsFunctionhood, SupportFalseTermsBlocked}, Failures: []string{FailureCrossLaneNotNativeWithoutSelector},
		},
		Z2Compatibility: Z2CompatibilityAudit{
			LambdaRepresentativeOne: "I_B(1)=e_lambda tensor P_3", LambdaRepresentativeTwo: "I_B(2)=(C_R^2 tensor W)-(e_lambda tensor P_1)", BarRepresentativeOne: "I_B(1)=e_barlambda tensor P_3", BarRepresentativeTwo: "I_B(2)=(C_R^2 tensor W)-(e_barlambda tensor P_1)", CommutesWithPhaseFlip: true, RanksRepresentativeFree: true, NativeZ2Selector: false,
			Supports: []string{SupportSelectorFunctionhoodZ2Compatible, SupportIBZ2CommutesWithPhaseFlip, SupportSelectorRanksRepresentativeFree}, Failures: []string{FailureZ2CompatibilityNotNativeSelector},
		},
		MuB: MuBContributionAudit{
			Formula: MuBFormula, AlphaFormula: BoundaryAlphaFormula, S: SBoundary, Rank1: RankExposureFace, Rank2: RankFullEnclosure, RankH10: RankH10, RankH72: RankH72, LinearContribution: linear, QuadraticContribution: quad, Alpha: alpha, NativeMuB: false,
			Supports: []string{SupportSelectorSuppliesMuBRanks, SupportMuBNativeGapReduced}, Failures: []string{FailureMuBStillNotNativeWithoutSelector, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3},
		},
		Firewalls: firewalls,
		Final:     StrategicConclusion,
	}, nil
}

func Statuses() []string {
	return []string{StatusInheritedGate922, StatusDegreeOneSourceTyped, StatusDegreeTwoSourceTyped, StatusCumulativeEnclosure, StatusFunctionhoodSupported, StatusCrossLaneConsequence, StatusZ2Compatible, StatusMuBGapReduced, StatusNativeSelectorMissing, StatusNativeR3Blocked}
}
func Supports() []string {
	return []string{SupportDegreeOneExposureSource, SupportSingleExposureTargetsFaceQuotient, SupportIBZ2OneExposedFace, SupportDegreeTwoEnclosureSource, SupportFullEnclosureTargetsComplement, SupportIBZ2TwoFullEnclosure, SupportDegreeTwoCumulativeNotGraded, SupportFullPairRequiresF2OverF0, SupportF2OverF1Rejected, SupportExposureEnclosureGivesFunctionhood, SupportEachDegreeUniqueTargetType, SupportIBZ2FunctionalIfAccepted, SupportCrossLaneFollowsFunctionhood, SupportFalseTermsBlocked, SupportSelectorFunctionhoodZ2Compatible, SupportIBZ2CommutesWithPhaseFlip, SupportSelectorRanksRepresentativeFree, SupportSelectorSuppliesMuBRanks, SupportMuBNativeGapReduced}
}
func Failures() []string {
	return []string{FailureNoNativeSelectorFunctionhood, FailureNoNativeDegreeToZ2FlagFunctor, FailureExposureToF1NotNative, FailureEnclosureToF2NotNative, FailureNoNativeCumulativeTheorem, FailureSelectorDependsOnBridgeRule, FailureCrossLaneNotNativeWithoutSelector, FailureZ2CompatibilityNotNativeSelector, FailureMuBStillNotNativeWithoutSelector, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}
}

func FormatLane(l SelectorLaneAudit) string {
	return fmt.Sprintf("degree=%d|term=%s|type=%s|target=%s|rank=%d|status=%s|bridge=%v|native=%v", l.Degree, l.ResponseTerm, l.SourceType, l.Target, l.Rank, l.SourceStatus, l.BridgeTyped, l.Native)
}
func FormatCumulative(c CumulativeEnclosureAudit) string {
	return fmt.Sprintf("degree2_target=%s|reject=%s|rank_f2_f0=%d|rank_f2_f1=%d|cumulative=%v|native=%v", c.DegreeTwoTarget, c.AssociatedGradedRejected, c.RankF2OverF0, c.RankF2OverF1, c.CumulativeRequired, c.Native)
}
func FormatFunctionhood(f FunctionhoodAudit) string {
	return fmt.Sprintf("domain=%v|rule=%s|exposure_enclosure=%v|functional_if_accepted=%v|native=%v|primary_gap=%s", f.DomainDegrees, f.TargetRule, f.ExposureEnclosureAccepted, f.FunctionalIfExposureAccepted, f.NativeFunctionhood, f.PrimaryGap)
}
func FormatCrossLane(c CrossLaneAudit) string {
	return fmt.Sprintf("false_linear=%s|false_quad=%s|terms=%s,%s|excluded_if_functional=%v|native=%v", c.FalseLinearLane, c.FalseQuadraticLane, c.FalseLinearTerm, c.FalseQuadraticTerm, c.ExcludedIfFunctional, c.NativeExclusion)
}
func FormatZ2(z Z2CompatibilityAudit) string {
	return fmt.Sprintf("lambda=(%s;%s)|bar=(%s;%s)|commutes=%v|rank_free=%v|native=%v", z.LambdaRepresentativeOne, z.LambdaRepresentativeTwo, z.BarRepresentativeOne, z.BarRepresentativeTwo, z.CommutesWithPhaseFlip, z.RanksRepresentativeFree, z.NativeZ2Selector)
}
func FormatMuB(m MuBContributionAudit) string {
	return fmt.Sprintf("formula=%s|alpha=%.18g|linear=%.18g|quadratic=%.18g|native_mu_b=%v", m.Formula, m.Alpha, m.LinearContribution, m.QuadraticContribution, m.NativeMuB)
}
func FormatFirewalls(f FirewallLedger) string { return strings.Join(f.List(), ";") }
func containsAll(haystack, needles []string) bool {
	seen := map[string]bool{}
	for _, h := range haystack {
		seen[h] = true
	}
	for _, n := range needles {
		if !seen[n] {
			return false
		}
	}
	return true
}
func near(a, b float64) bool { return math.Abs(a-b) < 1e-18 }
func firewallsOK(f FirewallLedger) bool {
	return f.NoNativeSelectorFunctionhood && f.NoNativeDegreeToZ2FlagFunctor && f.ExposureToF1NotNative && f.EnclosureToF2NotNative && f.NoNativeCumulativeTheorem && f.SelectorDependsOnBridgeRule && f.CrossLaneNotNativeWithoutSelector && f.Z2CompatibilityNotNativeSelector && f.MuBStillNotNativeWithoutSelector && f.AlphaBridgeCandidateNotNative && f.NotNativeR3 && f.FullAFDescentStillBlocked && f.NoGenerationCarrierMap && f.NoFlavorOrientationMap && f.NoIndividualYukawaValues && f.NoNativeYukawaOperator
}
