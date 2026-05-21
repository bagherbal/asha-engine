// Package generation2degreeindexedz2airlockflagfunctoraudit
// implements Gate 914: DegreeIndexed Z2 Airlock FlagFunctor Audit.
//
// Gate 914 follows Gate 913's certification of the reduced boundary-pair
// response shape. It audits only the second sub-object from Gate 912: whether
// the exterior degrees of R_B(s) can be typed as selectors for the Z2 airlock
// flag quotient classes. The gate supports the selector shape
//
//	deg(Lambda^1 B_2) -> [F_1/F_0]_{Z2}
//	deg(Lambda^2 B_2) -> [F_2/F_0]_{Z2}
//
// but it does not certify a native functor, does not prove independent
// cross-lane exclusion, does not transport S_split, and does not promote
// BoundaryAlpha or R3 to native status.
package generation2degreeindexedz2airlockflagfunctoraudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE914-DEGREEINDEXED-Z2-AIRLOCK-FLAGFUNCTOR-AUDIT"

	SBoundary = 0.0012924448188162962
	AlphaB    = 0.0003878958469680527

	Lambda1Dim       = 2
	Lambda2Dim       = 1
	RankF1OverF0     = 3
	RankF2OverF0     = 7
	RankF2OverF1     = 4
	LinearDenom      = 10
	QuadraticDenom   = 72
	BoundaryPairRank = 2

	Gate913Classification = "R3_REDUCED_B2_RESPONSE_FUNCTIONAL_SHAPE_CERTIFIED_NOT_NATIVE_BOUNDARY_ALPHA"
	Gate913ShortStatus    = "R3_ALPHA_SUBOBJECT_1_REDUCED_B2_RESPONSE_SHAPE_PASS_NATIVE_SELECTION_BLOCKED"
	Gate913Verdict        = "REDUCED_B2_RESPONSE_FUNCTIONAL_HAS_CANONICAL_EXTERIOR_SHAPE_BUT_NATIVE_SELECTION_REMAINS_BLOCKED"

	PunctureClass          = "[p]_{Z2}={e_lambda tensor P_1,e_barlambda tensor P_1}"
	F0Definition           = "F_0=p"
	F1Definition           = "F_1=e_phase tensor W"
	F2Definition           = "F_2=C_R^2 tensor W"
	Z2ExposedFaceClass     = "[F_1/F_0]_{Z2}={e_lambda tensor P_3,e_barlambda tensor P_3}"
	Z2FullEnclosureClass   = "[F_2/F_0]_{Z2}={(C_R^2 tensor W)-(e_lambda tensor P_1),(C_R^2 tensor W)-(e_barlambda tensor P_1)}"
	AssociatedGradedSlice  = "F_2/F_1"
	DegreeOneTerm          = "s(b1+b2)"
	DegreeTwoTerm          = "s^2(b1 wedge b2)"
	Lambda1ToExposed       = "deg(Lambda^1 B_2)->[F_1/F_0]_{Z2}"
	Lambda2ToFull          = "deg(Lambda^2 B_2)->[F_2/F_0]_{Z2}"
	ForbiddenLinearFull    = "Lambda^1 B_2 not -> [F_2/F_0]_{Z2}"
	ForbiddenQuadraticFace = "Lambda^2 B_2 not -> [F_1/F_0]_{Z2}"
	FalseLinearFullTerm    = "(7/72)s"
	FalseQuadraticFaceTerm = "(3/10)s^2"
	AlphaSkeleton          = "alpha_B=(3/10)s+(7/72)s^2"
	DegreeSelectorObject   = "DegreeIndexedZ2AirlockFlagFunctor"
	NextGate               = "NEXT_PRESSURE_GATE915_Z2_BOUNDARYALPHA_CROSSLANE_EXCLUSION_AUDIT"
	Classification         = "R3_ALPHA_SUBOBJECT_2_Z2_FLAG_SELECTOR_SHAPE_PASS_NATIVE_FUNCTOR_BLOCKED"
	ShortStatus            = "R3_DEGREE_INDEXED_Z2_AIRLOCK_FLAG_SELECTOR_OBSTRUCTION"
	FinalTruth             = "DEGREE_INDEXED_Z2_FLAG_SELECTOR_SHAPE_SUPPORTED_BUT_NATIVE_FUNCTOR_NOT_CERTIFIED"
	StrategicConclusion    = "R_B(s) supplies exterior degrees and the degree-indexed selector supplies target ranks, but cross-lane exclusion and S_split transport remain unresolved."

	StatusGate913Inherited       = "PASS_GATE913_REDUCED_B2_RESPONSE_SHAPE_INHERITED"
	StatusNoLoopBack             = "PASS_BRANCH_DOES_NOT_REOPEN_PHASE_SOCKET_OR_REPRESENTATIVE_AIRLOCK_WOUNDS"
	StatusDegreeOneTarget        = "PASS_DEGREE_ONE_TARGET_RECORDED_AS_Z2_EXPOSED_FACE_CLASS"
	StatusDegreeTwoTarget        = "PASS_DEGREE_TWO_TARGET_RECORDED_AS_Z2_FULL_ENCLOSURE_CLASS"
	StatusSelectorNotSurjection  = "PASS_DEGREE_TO_FLAG_OBJECT_TYPED_AS_SELECTOR_NOT_LINEAR_SURJECTION"
	StatusCumulativeChoice       = "PASS_DEGREE_TWO_CUMULATIVE_ENCLOSURE_CHOICE_RECORDED"
	StatusAssociatedGradedReject = "PASS_ASSOCIATED_GRADED_SLICE_REJECTED_FOR_ALPHA_TARGET"
	StatusAlphaRankReconstruct   = "PASS_SELECTOR_RECONSTRUCTS_ALPHA_RANK_PAIR_UNDER_SEAL"
	StatusCrossLaneConditional   = "PASS_CROSS_LANE_EXCLUSION_RECORDED_AS_CONDITIONAL_ON_CERTIFIED_SELECTOR"
	StatusNativeFunctorBlocked   = "FIREWALL_PRESERVED_DEGREE_INDEXED_SELECTOR_NOT_NATIVE_FUNCTOR"

	SupportDegreeOneTargetsExposed      = "CONDITIONAL_SUPPORT_DEGREE_ONE_BOUNDARY_RESPONSE_TARGETS_Z2_EXPOSED_FACE_CLASS"
	SupportLambda1SingleExposure        = "CONDITIONAL_SUPPORT_LAMBDA1B2_AS_SINGLE_BOUNDARY_EXPOSURE"
	SupportExposedFaceRankThree         = "CONDITIONAL_SUPPORT_EXPOSED_FACE_CLASS_HAS_RANK_THREE"
	SupportDegreeTwoTargetsFull         = "CONDITIONAL_SUPPORT_DEGREE_TWO_BOUNDARY_RESPONSE_TARGETS_Z2_FULL_ENCLOSURE_CLASS"
	SupportLambda2FullEnclosure         = "CONDITIONAL_SUPPORT_LAMBDA2B2_AS_FULL_BOUNDARY_PAIR_ENCLOSURE"
	SupportFullEnclosureRankSeven       = "CONDITIONAL_SUPPORT_FULL_ENCLOSURE_CLASS_HAS_RANK_SEVEN"
	SupportSelectorNotLinearSurjection  = "CONDITIONAL_SUPPORT_DEGREE_TO_FLAG_OBJECT_IS_SELECTOR_NOT_LINEAR_SURJECTION"
	SupportDimensionMismatchSelector    = "CONDITIONAL_SUPPORT_DIMENSION_MISMATCH_FORCES_SELECTOR_TYPING"
	SupportDegreeTwoCumulative          = "CONDITIONAL_SUPPORT_DEGREE_TWO_SELECTS_CUMULATIVE_ENCLOSURE_CLASS_F2_OVER_F0"
	SupportAssociatedGradedRejected     = "CONDITIONAL_SUPPORT_ASSOCIATED_GRADED_SLICE_F2_OVER_F1_REJECTED_FOR_ALPHA_TARGET"
	SupportSelectorReconstructsRankPair = "CONDITIONAL_SUPPORT_DEGREE_INDEXED_Z2_FLAG_SELECTOR_RECONSTRUCTS_ALPHA_RANK_PAIR"
	SupportRankPairFromSelectedClasses  = "CONDITIONAL_SUPPORT_ALPHA_RANK_PAIR_3_7_FOLLOWS_FROM_SELECTED_Z2_FLAG_CLASSES"
	SupportCrossLaneWouldFollow         = "CONDITIONAL_SUPPORT_CROSS_LANE_EXCLUSION_WOULD_FOLLOW_FROM_CERTIFIED_DEGREE_INDEXED_SELECTOR"

	FailureNoNativeDegreeToZ2FlagFunctor  = "FAILED_ROUTE_NO_NATIVE_DEGREE_TO_Z2_FLAG_CLASS_FUNCTOR"
	FailureNoNativeLambda1ExposedMap      = "FAILED_ROUTE_NO_NATIVE_LAMBDA1B2_TO_Z2_EXPOSED_FACE_CLASS_MAP"
	FailureNoNativeLambda2FullMap         = "FAILED_ROUTE_NO_NATIVE_LAMBDA2B2_TO_Z2_FULL_ENCLOSURE_CLASS_MAP"
	FailureLambdaKB2NotSurjection         = "FAILED_ROUTE_LAMBDAK_B2_NOT_LINEAR_SURJECTION_ONTO_Z2_FLAG_QUOTIENT"
	FailureNoNativeCumulativeReason       = "FAILED_ROUTE_NO_NATIVE_REASON_YET_FOR_CUMULATIVE_OVER_ASSOCIATED_GRADED_CHOICE"
	FailureNoIndependentCrossLane         = "FAILED_ROUTE_NO_INDEPENDENT_NATIVE_Z2_CROSS_LANE_EXCLUSION_THEOREM_YET"
	FailureSelectorRanksNotAlphaSource    = "FAILED_ROUTE_SELECTOR_RECONSTRUCTS_ALPHA_RANKS_BUT_NOT_NATIVE_ALPHA_SOURCE"
	FailureDenominatorsSTransportExternal = "FAILED_ROUTE_DENOMINATORS_AND_S_TRANSPORT_STILL_EXTERNAL_TO_SELECTOR"
	FailureAlphaStillSealed               = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNotNativeR3                    = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked      = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap         = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap         = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues       = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator         = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type InheritedResponse struct {
	Gate913Classification string
	Gate913ShortStatus    string
	Gate913Verdict        string
	ReducedShapeCertified bool
	ReopensPhaseSign      bool
	ReopensSocketOrder    bool
	ReopensRepresentative bool
	DerivesAlpha          bool
	UpdatesOfficialLedger bool
	Supports, Failures    []string
}

type Z2FlagClass struct {
	Name                 string
	Definition           string
	Representatives      []string
	Rank                 int
	RepresentativeFree   bool
	AssociatedGradedRank int
}

type DegreeTarget struct {
	Degree             int
	BoundaryTerm       string
	Source             string
	Target             Z2FlagClass
	Interpretation     string
	Selector           bool
	LinearSurjection   bool
	NativeMap          bool
	Supports, Failures []string
}

type SelectorTyping struct {
	Lambda1Dim, Lambda2Dim int
	ExposedRank, FullRank  int
	DimensionMismatch      bool
	SelectorNotSurjection  bool
	Supports, Failures     []string
}

type CumulativeEnclosureChoice struct {
	DegreeTwoTarget              string
	F2OverF0Rank                 int
	F2OverF1Rank                 int
	SelectsCumulativeEnclosure   bool
	RejectsAssociatedGradedSlice bool
	NativeReasonForChoice        bool
	Supports, Failures           []string
}

type AlphaRankReconstruction struct {
	Formula              string
	S                    float64
	Alpha                float64
	RankPair             [2]int
	Denominators         [2]int
	ReconstructsRankPair bool
	NativeAlphaSource    bool
	DenominatorsExternal bool
	STransportExternal   bool
	Supports, Failures   []string
}

type CrossLaneStatus struct {
	ForbiddenLanes           []string
	FalseTerms               []string
	WouldFollowFromSelector  bool
	IndependentNativeTheorem bool
	ProvesCrossLaneExclusion bool
	Supports, Failures       []string
}

type Firewalls struct {
	NativeDegreeToZ2FlagFunctor bool
	NativeLambda1ExposedMap     bool
	NativeLambda2FullMap        bool
	LinearSurjection            bool
	NativeCumulativeReason      bool
	IndependentCrossLane        bool
	NativeAlphaSource           bool
	DenominatorsAndSTransport   bool
	AlphaNative                 bool
	NativeR3                    bool
	FullAFDescent               bool
	GenerationCarrierMap        bool
	FlavorOrientationMap        bool
	IndividualYukawaValues      bool
	NativeYukawaOperator        bool
}

type Audit struct {
	ID             string
	Classification string
	ShortStatus    string
	Inherited      InheritedResponse
	DegreeOne      DegreeTarget
	DegreeTwo      DegreeTarget
	Typing         SelectorTyping
	Cumulative     CumulativeEnclosureChoice
	AlphaRanks     AlphaRankReconstruction
	CrossLane      CrossLaneStatus
	Firewalls      Firewalls
	Truth          string
	Final          string
}

func BoundaryAlpha(s float64) float64 {
	return float64(RankF1OverF0)/float64(LinearDenom)*s + float64(RankF2OverF0)/float64(QuadraticDenom)*s*s
}

func BuildDefault() (Audit, error) {
	inherited := buildInherited()
	if !inherited.ReducedShapeCertified || inherited.ReopensPhaseSign || inherited.ReopensSocketOrder || inherited.ReopensRepresentative || inherited.DerivesAlpha || inherited.UpdatesOfficialLedger {
		return Audit{}, fmt.Errorf("bad inherited response: %s", FormatInherited(inherited))
	}

	degreeOne := buildDegreeOneTarget()
	degreeTwo := buildDegreeTwoTarget()
	typing := buildSelectorTyping()
	cumulative := buildCumulativeChoice()
	alpha := buildAlphaReconstruction()
	cross := buildCrossLaneStatus()
	firewalls := buildFirewalls()

	if degreeOne.Target.Rank != RankF1OverF0 || degreeTwo.Target.Rank != RankF2OverF0 {
		return Audit{}, fmt.Errorf("bad rank targets: %s | %s", FormatDegreeTarget(degreeOne), FormatDegreeTarget(degreeTwo))
	}
	if !degreeOne.Selector || !degreeTwo.Selector || degreeOne.LinearSurjection || degreeTwo.LinearSurjection || degreeOne.NativeMap || degreeTwo.NativeMap {
		return Audit{}, fmt.Errorf("selector typing leaked into map/surjection: %s | %s", FormatDegreeTarget(degreeOne), FormatDegreeTarget(degreeTwo))
	}
	if !typing.DimensionMismatch || !typing.SelectorNotSurjection {
		return Audit{}, fmt.Errorf("dimension mismatch did not force selector typing: %s", FormatTyping(typing))
	}
	if !cumulative.SelectsCumulativeEnclosure || !cumulative.RejectsAssociatedGradedSlice || cumulative.NativeReasonForChoice || cumulative.F2OverF1Rank == RankF2OverF0 {
		return Audit{}, fmt.Errorf("bad cumulative choice: %s", FormatCumulative(cumulative))
	}
	if !alpha.ReconstructsRankPair || !near(alpha.Alpha, AlphaB) || alpha.NativeAlphaSource || !alpha.DenominatorsExternal || !alpha.STransportExternal {
		return Audit{}, fmt.Errorf("bad alpha rank reconstruction: %s", FormatAlpha(alpha))
	}
	if !cross.WouldFollowFromSelector || cross.IndependentNativeTheorem || cross.ProvesCrossLaneExclusion {
		return Audit{}, fmt.Errorf("cross-lane firewall leaked: %s", FormatCrossLane(cross))
	}
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewall leak: %s", FormatFirewalls(firewalls))
	}

	return Audit{
		ID:             AuditID,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		Inherited:      inherited,
		DegreeOne:      degreeOne,
		DegreeTwo:      degreeTwo,
		Typing:         typing,
		Cumulative:     cumulative,
		AlphaRanks:     alpha,
		CrossLane:      cross,
		Firewalls:      firewalls,
		Truth:          FinalTruth,
		Final:          NextGate,
	}, nil
}

func buildInherited() InheritedResponse {
	return InheritedResponse{
		Gate913Classification: Gate913Classification,
		Gate913ShortStatus:    Gate913ShortStatus,
		Gate913Verdict:        Gate913Verdict,
		ReducedShapeCertified: true,
		Supports:              []string{StatusGate913Inherited, StatusNoLoopBack},
		Failures:              []string{FailureNoNativeDegreeToZ2FlagFunctor, FailureAlphaStillSealed, FailureNotNativeR3},
	}
}

func buildDegreeOneTarget() DegreeTarget {
	return DegreeTarget{
		Degree:       1,
		BoundaryTerm: DegreeOneTerm,
		Source:       "Lambda^1 B_2",
		Target: Z2FlagClass{
			Name:               "Z2 exposed-face class",
			Definition:         Z2ExposedFaceClass,
			Representatives:    []string{"e_lambda tensor P_3", "e_barlambda tensor P_3"},
			Rank:               RankF1OverF0,
			RepresentativeFree: true,
		},
		Interpretation:   "single-boundary exposure selects the exposed socket-face complement F_1/F_0",
		Selector:         true,
		LinearSurjection: false,
		NativeMap:        false,
		Supports:         []string{StatusDegreeOneTarget, SupportDegreeOneTargetsExposed, SupportLambda1SingleExposure, SupportExposedFaceRankThree},
		Failures:         []string{FailureNoNativeLambda1ExposedMap, FailureNoNativeDegreeToZ2FlagFunctor},
	}
}

func buildDegreeTwoTarget() DegreeTarget {
	return DegreeTarget{
		Degree:       2,
		BoundaryTerm: DegreeTwoTerm,
		Source:       "Lambda^2 B_2",
		Target: Z2FlagClass{
			Name:                 "Z2 full-enclosure class",
			Definition:           Z2FullEnclosureClass,
			Representatives:      []string{"(C_R^2 tensor W)-(e_lambda tensor P_1)", "(C_R^2 tensor W)-(e_barlambda tensor P_1)"},
			Rank:                 RankF2OverF0,
			RepresentativeFree:   true,
			AssociatedGradedRank: RankF2OverF1,
		},
		Interpretation:   "full boundary-pair enclosure selects the full puncture-complement active domain F_2/F_0",
		Selector:         true,
		LinearSurjection: false,
		NativeMap:        false,
		Supports:         []string{StatusDegreeTwoTarget, SupportDegreeTwoTargetsFull, SupportLambda2FullEnclosure, SupportFullEnclosureRankSeven},
		Failures:         []string{FailureNoNativeLambda2FullMap, FailureNoNativeDegreeToZ2FlagFunctor},
	}
}

func buildSelectorTyping() SelectorTyping {
	return SelectorTyping{
		Lambda1Dim:            Lambda1Dim,
		Lambda2Dim:            Lambda2Dim,
		ExposedRank:           RankF1OverF0,
		FullRank:              RankF2OverF0,
		DimensionMismatch:     Lambda1Dim != RankF1OverF0 || Lambda2Dim != RankF2OverF0,
		SelectorNotSurjection: true,
		Supports:              []string{StatusSelectorNotSurjection, SupportSelectorNotLinearSurjection, SupportDimensionMismatchSelector},
		Failures:              []string{FailureLambdaKB2NotSurjection, FailureNoNativeDegreeToZ2FlagFunctor},
	}
}

func buildCumulativeChoice() CumulativeEnclosureChoice {
	return CumulativeEnclosureChoice{
		DegreeTwoTarget:              Lambda2ToFull,
		F2OverF0Rank:                 RankF2OverF0,
		F2OverF1Rank:                 RankF2OverF1,
		SelectsCumulativeEnclosure:   true,
		RejectsAssociatedGradedSlice: true,
		NativeReasonForChoice:        false,
		Supports:                     []string{StatusCumulativeChoice, StatusAssociatedGradedReject, SupportDegreeTwoCumulative, SupportAssociatedGradedRejected},
		Failures:                     []string{FailureNoNativeCumulativeReason},
	}
}

func buildAlphaReconstruction() AlphaRankReconstruction {
	return AlphaRankReconstruction{
		Formula:              AlphaSkeleton,
		S:                    SBoundary,
		Alpha:                BoundaryAlpha(SBoundary),
		RankPair:             [2]int{RankF1OverF0, RankF2OverF0},
		Denominators:         [2]int{LinearDenom, QuadraticDenom},
		ReconstructsRankPair: true,
		NativeAlphaSource:    false,
		DenominatorsExternal: true,
		STransportExternal:   true,
		Supports:             []string{StatusAlphaRankReconstruct, SupportSelectorReconstructsRankPair, SupportRankPairFromSelectedClasses},
		Failures:             []string{FailureSelectorRanksNotAlphaSource, FailureDenominatorsSTransportExternal, FailureAlphaStillSealed},
	}
}

func buildCrossLaneStatus() CrossLaneStatus {
	return CrossLaneStatus{
		ForbiddenLanes:           []string{ForbiddenLinearFull, ForbiddenQuadraticFace},
		FalseTerms:               []string{FalseLinearFullTerm, FalseQuadraticFaceTerm},
		WouldFollowFromSelector:  true,
		IndependentNativeTheorem: false,
		ProvesCrossLaneExclusion: false,
		Supports:                 []string{StatusCrossLaneConditional, SupportCrossLaneWouldFollow},
		Failures:                 []string{FailureNoIndependentCrossLane},
	}
}

func buildFirewalls() Firewalls { return Firewalls{} }

func firewallsOK(f Firewalls) bool {
	return !f.NativeDegreeToZ2FlagFunctor &&
		!f.NativeLambda1ExposedMap &&
		!f.NativeLambda2FullMap &&
		!f.LinearSurjection &&
		!f.NativeCumulativeReason &&
		!f.IndependentCrossLane &&
		!f.NativeAlphaSource &&
		!f.DenominatorsAndSTransport &&
		!f.AlphaNative &&
		!f.NativeR3 &&
		!f.FullAFDescent &&
		!f.GenerationCarrierMap &&
		!f.FlavorOrientationMap &&
		!f.IndividualYukawaValues &&
		!f.NativeYukawaOperator
}

func (a Audit) FirewallsList() []string {
	return []string{
		FailureNoNativeDegreeToZ2FlagFunctor,
		FailureNoNativeLambda1ExposedMap,
		FailureNoNativeLambda2FullMap,
		FailureLambdaKB2NotSurjection,
		FailureNoNativeCumulativeReason,
		FailureNoIndependentCrossLane,
		FailureSelectorRanksNotAlphaSource,
		FailureDenominatorsSTransportExternal,
		FailureAlphaStillSealed,
		FailureNotNativeR3,
		FailureFullAFDescentStillBlocked,
		FailureNoGenerationCarrierMap,
		FailureNoFlavorOrientationMap,
		FailureNoIndividualYukawaValues,
		FailureNoNativeYukawaOperator,
	}
}

func Statuses() []string {
	return []string{
		StatusGate913Inherited,
		StatusNoLoopBack,
		StatusDegreeOneTarget,
		StatusDegreeTwoTarget,
		StatusSelectorNotSurjection,
		StatusCumulativeChoice,
		StatusAssociatedGradedReject,
		StatusAlphaRankReconstruct,
		StatusCrossLaneConditional,
		StatusNativeFunctorBlocked,
	}
}

func Supports() []string {
	return []string{
		SupportDegreeOneTargetsExposed,
		SupportLambda1SingleExposure,
		SupportExposedFaceRankThree,
		SupportDegreeTwoTargetsFull,
		SupportLambda2FullEnclosure,
		SupportFullEnclosureRankSeven,
		SupportSelectorNotLinearSurjection,
		SupportDimensionMismatchSelector,
		SupportDegreeTwoCumulative,
		SupportAssociatedGradedRejected,
		SupportSelectorReconstructsRankPair,
		SupportRankPairFromSelectedClasses,
		SupportCrossLaneWouldFollow,
	}
}

func Failures() []string {
	return []string{
		FailureNoNativeDegreeToZ2FlagFunctor,
		FailureNoNativeLambda1ExposedMap,
		FailureNoNativeLambda2FullMap,
		FailureLambdaKB2NotSurjection,
		FailureNoNativeCumulativeReason,
		FailureNoIndependentCrossLane,
		FailureSelectorRanksNotAlphaSource,
		FailureDenominatorsSTransportExternal,
		FailureAlphaStillSealed,
		FailureNotNativeR3,
		FailureFullAFDescentStillBlocked,
		FailureNoGenerationCarrierMap,
		FailureNoFlavorOrientationMap,
		FailureNoIndividualYukawaValues,
		FailureNoNativeYukawaOperator,
	}
}

func FormatInherited(x InheritedResponse) string {
	return fmt.Sprintf("Gate913(classification=%s short=%s verdict=%s reduced_shape_certified=%t reopens_phase=%t reopens_socket=%t reopens_representative=%t derives_alpha=%t updates_official=%t supports=%s failures=%s)", x.Gate913Classification, x.Gate913ShortStatus, x.Gate913Verdict, x.ReducedShapeCertified, x.ReopensPhaseSign, x.ReopensSocketOrder, x.ReopensRepresentative, x.DerivesAlpha, x.UpdatesOfficialLedger, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}

func FormatFlagClass(x Z2FlagClass) string {
	return fmt.Sprintf("%s(definition=%s rank=%d representative_free=%t reps=%s associated_graded_rank=%d)", x.Name, x.Definition, x.Rank, x.RepresentativeFree, strings.Join(x.Representatives, ";"), x.AssociatedGradedRank)
}

func FormatDegreeTarget(x DegreeTarget) string {
	return fmt.Sprintf("DegreeTarget(degree=%d term=%s source=%s target=%s interpretation=%s selector=%t linear_surjection=%t native_map=%t supports=%s failures=%s)", x.Degree, x.BoundaryTerm, x.Source, FormatFlagClass(x.Target), x.Interpretation, x.Selector, x.LinearSurjection, x.NativeMap, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}

func FormatTyping(x SelectorTyping) string {
	return fmt.Sprintf("SelectorTyping(dim_Lambda1=%d dim_Lambda2=%d exposed_rank=%d full_rank=%d dimension_mismatch=%t selector_not_surjection=%t supports=%s failures=%s)", x.Lambda1Dim, x.Lambda2Dim, x.ExposedRank, x.FullRank, x.DimensionMismatch, x.SelectorNotSurjection, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}

func FormatCumulative(x CumulativeEnclosureChoice) string {
	return fmt.Sprintf("CumulativeChoice(target=%s rank_F2_over_F0=%d rank_F2_over_F1=%d cumulative=%t reject_associated_graded=%t native_reason=%t supports=%s failures=%s)", x.DegreeTwoTarget, x.F2OverF0Rank, x.F2OverF1Rank, x.SelectsCumulativeEnclosure, x.RejectsAssociatedGradedSlice, x.NativeReasonForChoice, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}

func FormatAlpha(x AlphaRankReconstruction) string {
	return fmt.Sprintf("AlphaRankReconstruction(formula=%s s=%.16g alpha=%.16g rank_pair=%v denominators=%v reconstructs=%t native_source=%t denominators_external=%t s_transport_external=%t supports=%s failures=%s)", x.Formula, x.S, x.Alpha, x.RankPair, x.Denominators, x.ReconstructsRankPair, x.NativeAlphaSource, x.DenominatorsExternal, x.STransportExternal, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}

func FormatCrossLane(x CrossLaneStatus) string {
	return fmt.Sprintf("CrossLane(forbidden=%s false_terms=%s would_follow_from_selector=%t independent_native=%t proves=%t supports=%s failures=%s)", strings.Join(x.ForbiddenLanes, ";"), strings.Join(x.FalseTerms, ";"), x.WouldFollowFromSelector, x.IndependentNativeTheorem, x.ProvesCrossLaneExclusion, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("Firewalls(native_degree_to_z2_flag=%t lambda1_exposed=%t lambda2_full=%t linear_surjection=%t cumulative_reason=%t independent_cross_lane=%t native_alpha_source=%t denom_s_transport=%t alpha_native=%t r3_native=%t full_af=%t generation=%t flavor=%t individual_yukawa=%t native_yukawa=%t)", x.NativeDegreeToZ2FlagFunctor, x.NativeLambda1ExposedMap, x.NativeLambda2FullMap, x.LinearSurjection, x.NativeCumulativeReason, x.IndependentCrossLane, x.NativeAlphaSource, x.DenominatorsAndSTransport, x.AlphaNative, x.NativeR3, x.FullAFDescent, x.GenerationCarrierMap, x.FlavorOrientationMap, x.IndividualYukawaValues, x.NativeYukawaOperator)
}

func containsAll(got, want []string) bool {
	set := map[string]bool{}
	for _, g := range got {
		set[g] = true
	}
	for _, w := range want {
		if !set[w] {
			return false
		}
	}
	return true
}

func near(x, y float64) bool { return math.Abs(x-y) <= 1e-15 }
