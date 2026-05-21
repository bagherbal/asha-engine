// Package generation2airlocksupportclosureoperatorexistenceidempotenceaudit implements
// Gate 930: AirlockSupportClosureOperator Existence and Idempotence Audit.
//
// Gate 930 follows Gate 929's result that the airlock closure axioms are
// source-typed by the puncture-airlock flag, but the operator itself remains
// missing. This gate turns the least-support closure into an explicit finite
// bridge-level operator on the candidate support chain and audits whether it is
// extensive, monotone, idempotent, basepoint-preserving, minimally nontrivial,
// saturated at full boundary-pair activation, and Z2-equivariant. It still does
// not certify a native ASHA closure theorem.
package generation2airlocksupportclosureoperatorexistenceidempotenceaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE930-AIRLOCK-SUPPORT-CLOSURE-OPERATOR-EXISTENCE-IDEMPOTENCE-AUDIT"

	Gate929ShortStatus = "R3_ALPHA_CLOSURE_AXIOMS_SOURCED_TO_LEAST_SUPPORT_OPERATOR_GAP"

	AdmissibleSupportFamily   = "A_airlock={F_0,F_1,F_2} finite puncture-rooted support chain"
	Z2AdmissibleSupportFamily = "A_airlock^Z2={[F_0]_{Z2},[F_1]_{Z2},[F_2]_{Z2}}"
	BoundaryDemandChain       = "0<1<2 boundary activation demand chain"
	ClosureOperatorName       = "AirlockSupportClosureOperator"
	ClosureDefinition         = "Cl_airlock(k)=least admissible support satisfying boundary activation demand k"
	ClosureZero               = "Cl_airlock(0)=F_0"
	ClosureOne                = "Cl_airlock(1)=F_1"
	ClosureTwo                = "Cl_airlock(2)=F_2"
	ThetaViaClosure           = "Theta_B^Z2(k)=[Cl_airlock^Z2(k)/F_0]_{Z2}"
	MuBViaClosure             = "mu_B(R_B(S_split))=sum_{k=1}^2 rank([Cl_airlock^Z2(k)/F_0]_{Z2})/rank(H_k)*S_split^k"
	AlphaViaClosureOperator   = "mu_B(R_B(S_split))=(3/10)S_split+(7/72)S_split^2"
	NextGate                  = "NEXT_PRESSURE_GATE931_AIRLOCK_ADMISSIBLESUPPORT_LATTICE_SOURCE_AUDIT"

	RankF0       = 1
	RankF1       = 4
	RankF2       = 8
	RankF1OverF0 = 3
	RankF2OverF0 = 7
	RankH10      = 10
	RankH72      = 72

	Classification = "R3_AIRLOCK_SUPPORT_CLOSURE_OPERATOR_EXISTS_AS_BRIDGE_CLOSURE_NOT_NATIVE"
	ShortStatus    = "R3_ALPHA_CLOSURE_OPERATOR_EXISTS_NATIVE_SOURCE_MISSING"
	FinalTruth     = "AIRLOCK_SUPPORT_CLOSURE_OPERATOR_EXISTS_AND_IS_EXTENSIVE_MONOTONE_IDEMPOTENT_Z2_EQUIVARIANT_ON_CANDIDATE_SUPPORT_CHAIN_BUT_NATIVE_SOURCE_REMAINS_MISSING"

	StatusInheritedGate929       = "PASS_GATE929_FLAG_GENERATED_AXIOM_SOURCE_INHERITED"
	StatusFiniteOperatorExists   = "PASS_AIRLOCK_SUPPORT_CLOSURE_OPERATOR_EXISTS_ON_FINITE_FLAG_CHAIN"
	StatusExtensiveDemandTyped   = "PASS_CLOSURE_EXTENSIVE_AT_BOUNDARY_DEMAND_SUPPORT_LEVEL"
	StatusMonotone               = "PASS_AIRLOCK_CLOSURE_MONOTONE_ON_CANDIDATE_CHAIN"
	StatusIdempotent             = "PASS_AIRLOCK_CLOSURE_IDEMPOTENT_ON_ADMISSIBLE_SUPPORTS"
	StatusMinimalNonbase         = "PASS_CL_1_EQUALS_F1_BY_LEAST_NONBASE_ADMISSIBLE_SUPPORT"
	StatusSaturatedFullPair      = "PASS_CL_2_EQUALS_F2_BY_FULL_PAIR_SATURATION"
	StatusZ2Equivariant          = "PASS_AIRLOCK_CLOSURE_Z2_EQUIVARIANT_AT_CLASS_LEVEL"
	StatusTargetFunctorRecovered = "PASS_THETA_B_Z2_RECOVERED_FROM_AIRLOCK_CLOSURE_OPERATOR"
	StatusMeasureRewritten       = "PASS_BOUNDARY_ACTIVATION_MEASURE_REWRITTEN_USING_AIRLOCK_SUPPORT_CLOSURE"
	StatusNativeSourceMissing    = "FIREWALL_PRESERVED_NATIVE_AIRLOCK_SUPPORT_CLOSURE_OPERATOR_MISSING"
	StatusAlphaR3StillBlocked    = "FIREWALL_PRESERVED_ALPHA_AND_R3_NOT_NATIVE"

	SupportClosureExistsOnFiniteFlagChain      = "CONDITIONAL_SUPPORT_AIRLOCK_SUPPORT_CLOSURE_OPERATOR_EXISTS_ON_FINITE_FLAG_CHAIN"
	SupportEachDemandHasLeastSupport           = "CONDITIONAL_SUPPORT_EACH_BOUNDARY_DEMAND_HAS_LEAST_ADMISSIBLE_SUPPORT"
	SupportExistenceFromFiniteChain            = "CONDITIONAL_SUPPORT_CLOSURE_EXISTENCE_FOLLOWS_FROM_FINITE_CHAIN_STRUCTURE"
	SupportClosureExtensiveDemandLevel         = "CONDITIONAL_SUPPORT_CLOSURE_IS_EXTENSIVE_AT_BOUNDARY_DEMAND_SUPPORT_LEVEL"
	SupportCl0ContainsBasepointDemand          = "CONDITIONAL_SUPPORT_CL_0_CONTAINS_BASEPOINT_DEMAND"
	SupportCl1ContainsExposureDemand           = "CONDITIONAL_SUPPORT_CL_1_CONTAINS_EXPOSURE_DEMAND"
	SupportCl2ContainsFullEnclosureDemand      = "CONDITIONAL_SUPPORT_CL_2_CONTAINS_FULL_ENCLOSURE_DEMAND"
	SupportAirlockClosureMonotone              = "CONDITIONAL_SUPPORT_AIRLOCK_CLOSURE_IS_MONOTONE"
	SupportDemandOrderMatchesSupportInclusion  = "CONDITIONAL_SUPPORT_BOUNDARY_DEMAND_ORDER_MATCHES_SUPPORT_INCLUSION_ORDER"
	SupportCl0SubsetCl1SubsetCl2               = "CONDITIONAL_SUPPORT_CL_0_SUBSET_CL_1_SUBSET_CL_2"
	SupportAirlockClosureIdempotent            = "CONDITIONAL_SUPPORT_AIRLOCK_CLOSURE_IS_IDEMPOTENT"
	SupportF0F1F2ClosedSupports                = "CONDITIONAL_SUPPORT_F0_F1_F2_ARE_CLOSED_SUPPORTS"
	SupportClosureTwiceNoChange                = "CONDITIONAL_SUPPORT_APPLYING_CLOSURE_TWICE_DOES_NOT_CHANGE_TARGET"
	SupportCl1ByLeastNonbaseSupport            = "CONDITIONAL_SUPPORT_CL_1_EQUALS_F1_BY_LEAST_NONBASE_ADMISSIBLE_SUPPORT"
	SupportSingletonDoesNotJumpToF2            = "CONDITIONAL_SUPPORT_SINGLETON_BOUNDARY_ACTIVATION_DOES_NOT_JUMP_TO_F2"
	SupportMinimalExposureForced               = "CONDITIONAL_SUPPORT_MINIMAL_EXPOSURE_TARGET_IS_FORCED_WITHIN_ADMISSIBLE_CHAIN"
	SupportCl2ByFullPairSaturation             = "CONDITIONAL_SUPPORT_CL_2_EQUALS_F2_BY_FULL_PAIR_SATURATION"
	SupportTopDemandRequiresSaturatedRectangle = "CONDITIONAL_SUPPORT_TOP_BOUNDARY_DEMAND_REQUIRES_SATURATED_RIGHT_RECTANGLE"
	SupportFullPairCannotCloseToF1             = "CONDITIONAL_SUPPORT_FULL_PAIR_ACTIVATION_CANNOT_CLOSE_TO_F1"
	SupportAirlockClosureZ2Equivariant         = "CONDITIONAL_SUPPORT_AIRLOCK_CLOSURE_IS_Z2_EQUIVARIANT"
	SupportPhaseFlipCommutesWithClosure        = "CONDITIONAL_SUPPORT_GLOBAL_PHASE_FLIP_COMMUTES_WITH_CLOSURE_OPERATOR"
	SupportClosureDescendsToZ2Class            = "CONDITIONAL_SUPPORT_CLOSURE_DESCENDS_TO_Z2_SUPPORT_CLASS"
	SupportThetaRecoveredFromClosure           = "CONDITIONAL_SUPPORT_THETA_B_Z2_RECOVERED_FROM_AIRLOCK_CLOSURE_OPERATOR"
	SupportFixedBaseQuotientCumulativeTargets  = "CONDITIONAL_SUPPORT_FIXED_BASE_QUOTIENT_RECOVERS_CUMULATIVE_TARGETS"
	SupportAssociatedGradedRejectedByF0Root    = "CONDITIONAL_SUPPORT_ASSOCIATED_GRADED_SLICE_REJECTED_BY_CLOSURE_ROOTED_AT_F0"
	SupportMuBRewrittenUsingClosure            = "CONDITIONAL_SUPPORT_BOUNDARY_ACTIVATION_MEASURE_REWRITTEN_USING_AIRLOCK_SUPPORT_CLOSURE"
	SupportAlphaReconstructedThroughClosure    = "CONDITIONAL_SUPPORT_ALPHA_B_RECONSTRUCTED_THROUGH_CLOSURE_OPERATOR"
	SupportMuBTargetGapReducedToClosureStatus  = "CONDITIONAL_SUPPORT_MU_B_TARGET_FUNCTOR_GAP_REDUCED_TO_NATIVE_STATUS_OF_CLOSURE_OPERATOR"

	FailureExistenceBridgeNotNative         = "FAILED_ROUTE_EXISTENCE_ON_BRIDGE_SUPPORT_CHAIN_NOT_NATIVE_ASHA_CLOSURE_THEOREM"
	FailureExtensivityDemandTypedNotNative  = "FAILED_ROUTE_EXTENSIVITY_IS_DEMAND_TYPED_NOT_NATIVE_SUBSPACE_CLOSURE_YET"
	FailureMonotonicityCandidateNotNative   = "FAILED_ROUTE_MONOTONICITY_HOLDS_ON_CANDIDATE_CHAIN_NOT_NATIVE_BOUNDARY_ACTION_THEOREM"
	FailureIdempotenceCandidateNotNative    = "FAILED_ROUTE_IDEMPOTENCE_ON_ADMISSIBLE_FAMILY_NOT_NATIVE_CLOSURE_OPERATOR_THEOREM"
	FailureLeastNonbaseSupportNotNative     = "FAILED_ROUTE_LEAST_NONBASE_SUPPORT_RULE_NOT_NATIVE_ASHA_THEOREM"
	FailureFullPairSaturationNotNative      = "FAILED_ROUTE_FULL_PAIR_SATURATION_RULE_NOT_NATIVE_ASHA_THEOREM"
	FailureZ2EquivarianceClosureNotNative   = "FAILED_ROUTE_Z2_EQUIVARIANCE_OF_CLOSURE_NOT_NATIVE_GLOBAL_PHASE_THEOREM"
	FailureFixedBaseQuotientBridgeNotNative = "FAILED_ROUTE_FIXED_BASE_QUOTIENT_STILL_BRIDGE_MEASURE_RULE_NOT_NATIVE"
	FailureNoNativeAirlockSupportClosure    = "FAILED_ROUTE_NO_NATIVE_AIRLOCK_SUPPORT_CLOSURE_OPERATOR"
	FailureAlphaBridgeCandidateNotNative    = "FAILED_ROUTE_ALPHA_B_REMAINS_BRIDGE_CANDIDATE_NOT_NATIVE"
	FailureNotNativeR3                      = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked        = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap           = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap           = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues         = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator           = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type SupportChainAudit struct {
	AdmissibleFamily     string
	Z2AdmissibleFamily   string
	FiniteChain          bool
	LeastSupportsExist   bool
	NativeClosureTheorem bool
	Supports             []string
	Failures             []string
}

type ExtensivityAudit struct {
	DemandTyped           bool
	Cl0ContainsBasepoint  bool
	Cl1ContainsExposure   bool
	Cl2ContainsEnclosure  bool
	NativeSubspaceClosure bool
	Supports              []string
	Failures              []string
}

type MonotonicityAudit struct {
	DemandOrder         string
	SupportOrder        string
	Monotone            bool
	SupportInclusion    bool
	NativeActionTheorem bool
	Supports            []string
	Failures            []string
}

type IdempotenceAudit struct {
	ImageInAdmissibleFamily bool
	F0Closed                bool
	F1Closed                bool
	F2Closed                bool
	Idempotent              bool
	NativeClosureTheorem    bool
	Supports                []string
	Failures                []string
}

type MinimalityAudit struct {
	DemandK             int
	ClosureTarget       string
	LeastNonbaseSupport bool
	RejectsJumpToF2     bool
	NativeRule          bool
	Supports            []string
	Failures            []string
}

type SaturationAudit struct {
	DemandK            int
	ClosureTarget      string
	FullPairSaturation bool
	RejectsCloseToF1   bool
	NativeRule         bool
	Supports           []string
	Failures           []string
}

type Z2EquivarianceAudit struct {
	LambdaLadder       string
	BarLambdaLadder    string
	PhaseFlipCommutes  bool
	DescendsToZ2Class  bool
	NativePhaseTheorem bool
	Supports           []string
	Failures           []string
}

type TargetRecoveryAudit struct {
	ThetaFormula            string
	ThetaOne                string
	ThetaTwo                string
	ThetaOneRank            int
	ThetaTwoRank            int
	CumulativeTargets       bool
	RejectsAssociatedGraded bool
	NativeFixedBaseTheorem  bool
	Supports                []string
	Failures                []string
}

type MeasureConsequenceAudit struct {
	MeasureFormula        string
	AlphaFormula          string
	ThetaOneRank          int
	ThetaTwoRank          int
	H10Rank               int
	H72Rank               int
	RewrittenUsingClosure bool
	AlphaReconstructed    bool
	NativeAlpha           bool
	Supports              []string
	Failures              []string
}

type FirewallLedger struct {
	ExistenceBridgeNotNative         bool
	ExtensivityDemandTypedNotNative  bool
	MonotonicityCandidateNotNative   bool
	IdempotenceCandidateNotNative    bool
	LeastNonbaseSupportNotNative     bool
	FullPairSaturationNotNative      bool
	Z2EquivarianceClosureNotNative   bool
	FixedBaseQuotientBridgeNotNative bool
	NoNativeAirlockSupportClosure    bool
	AlphaBridgeCandidateNotNative    bool
	NotNativeR3                      bool
	FullAFDescentStillBlocked        bool
	NoGenerationCarrierMap           bool
	NoFlavorOrientationMap           bool
	NoIndividualYukawaValues         bool
	NoNativeYukawaOperator           bool
}

type Audit struct {
	ID             string
	Inherited      string
	SupportChain   SupportChainAudit
	Extensivity    ExtensivityAudit
	Monotonicity   MonotonicityAudit
	Idempotence    IdempotenceAudit
	Minimality     MinimalityAudit
	Saturation     SaturationAudit
	Z2             Z2EquivarianceAudit
	TargetRecovery TargetRecoveryAudit
	Measure        MeasureConsequenceAudit
	Firewalls      FirewallLedger
	Truth          string
	Classification string
	ShortStatus    string
	Final          string
}

func BuildDefault() (Audit, error) {
	return Audit{
		ID:             AuditID,
		Inherited:      Gate929ShortStatus,
		SupportChain:   SupportChainAudit{AdmissibleFamily: AdmissibleSupportFamily, Z2AdmissibleFamily: Z2AdmissibleSupportFamily, FiniteChain: true, LeastSupportsExist: true, NativeClosureTheorem: false, Supports: []string{SupportClosureExistsOnFiniteFlagChain, SupportEachDemandHasLeastSupport, SupportExistenceFromFiniteChain}, Failures: []string{FailureExistenceBridgeNotNative}},
		Extensivity:    ExtensivityAudit{DemandTyped: true, Cl0ContainsBasepoint: true, Cl1ContainsExposure: true, Cl2ContainsEnclosure: true, NativeSubspaceClosure: false, Supports: []string{SupportClosureExtensiveDemandLevel, SupportCl0ContainsBasepointDemand, SupportCl1ContainsExposureDemand, SupportCl2ContainsFullEnclosureDemand}, Failures: []string{FailureExtensivityDemandTypedNotNative}},
		Monotonicity:   MonotonicityAudit{DemandOrder: BoundaryDemandChain, SupportOrder: "F_0 subset F_1 subset F_2", Monotone: true, SupportInclusion: true, NativeActionTheorem: false, Supports: []string{SupportAirlockClosureMonotone, SupportDemandOrderMatchesSupportInclusion, SupportCl0SubsetCl1SubsetCl2}, Failures: []string{FailureMonotonicityCandidateNotNative}},
		Idempotence:    IdempotenceAudit{ImageInAdmissibleFamily: true, F0Closed: true, F1Closed: true, F2Closed: true, Idempotent: true, NativeClosureTheorem: false, Supports: []string{SupportAirlockClosureIdempotent, SupportF0F1F2ClosedSupports, SupportClosureTwiceNoChange}, Failures: []string{FailureIdempotenceCandidateNotNative}},
		Minimality:     MinimalityAudit{DemandK: 1, ClosureTarget: "F_1", LeastNonbaseSupport: true, RejectsJumpToF2: true, NativeRule: false, Supports: []string{SupportCl1ByLeastNonbaseSupport, SupportSingletonDoesNotJumpToF2, SupportMinimalExposureForced}, Failures: []string{FailureLeastNonbaseSupportNotNative}},
		Saturation:     SaturationAudit{DemandK: 2, ClosureTarget: "F_2", FullPairSaturation: true, RejectsCloseToF1: true, NativeRule: false, Supports: []string{SupportCl2ByFullPairSaturation, SupportTopDemandRequiresSaturatedRectangle, SupportFullPairCannotCloseToF1}, Failures: []string{FailureFullPairSaturationNotNative}},
		Z2:             Z2EquivarianceAudit{LambdaLadder: "F_0^lambda -> F_1^lambda -> F_2", BarLambdaLadder: "F_0^barlambda -> F_1^barlambda -> F_2", PhaseFlipCommutes: true, DescendsToZ2Class: true, NativePhaseTheorem: false, Supports: []string{SupportAirlockClosureZ2Equivariant, SupportPhaseFlipCommutesWithClosure, SupportClosureDescendsToZ2Class}, Failures: []string{FailureZ2EquivarianceClosureNotNative}},
		TargetRecovery: TargetRecoveryAudit{ThetaFormula: ThetaViaClosure, ThetaOne: "[F_1/F_0]_{Z2}", ThetaTwo: "[F_2/F_0]_{Z2}", ThetaOneRank: RankF1OverF0, ThetaTwoRank: RankF2OverF0, CumulativeTargets: true, RejectsAssociatedGraded: true, NativeFixedBaseTheorem: false, Supports: []string{SupportThetaRecoveredFromClosure, SupportFixedBaseQuotientCumulativeTargets, SupportAssociatedGradedRejectedByF0Root}, Failures: []string{FailureFixedBaseQuotientBridgeNotNative}},
		Measure:        MeasureConsequenceAudit{MeasureFormula: MuBViaClosure, AlphaFormula: AlphaViaClosureOperator, ThetaOneRank: RankF1OverF0, ThetaTwoRank: RankF2OverF0, H10Rank: RankH10, H72Rank: RankH72, RewrittenUsingClosure: true, AlphaReconstructed: true, NativeAlpha: false, Supports: []string{SupportMuBRewrittenUsingClosure, SupportAlphaReconstructedThroughClosure, SupportMuBTargetGapReducedToClosureStatus}, Failures: []string{FailureNoNativeAirlockSupportClosure, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3}},
		Firewalls:      FirewallLedger{ExistenceBridgeNotNative: true, ExtensivityDemandTypedNotNative: true, MonotonicityCandidateNotNative: true, IdempotenceCandidateNotNative: true, LeastNonbaseSupportNotNative: true, FullPairSaturationNotNative: true, Z2EquivarianceClosureNotNative: true, FixedBaseQuotientBridgeNotNative: true, NoNativeAirlockSupportClosure: true, AlphaBridgeCandidateNotNative: true, NotNativeR3: true, FullAFDescentStillBlocked: true, NoGenerationCarrierMap: true, NoFlavorOrientationMap: true, NoIndividualYukawaValues: true, NoNativeYukawaOperator: true},
		Truth:          FinalTruth, Classification: Classification, ShortStatus: ShortStatus,
		Final: "Gate 930 constructs an explicit bridge-level AirlockSupportClosureOperator on the finite support chain, verifies closure-operator laws, and preserves the missing native source for the admissible support lattice.",
	}, nil
}

func Statuses() []string {
	return []string{StatusInheritedGate929, StatusFiniteOperatorExists, StatusExtensiveDemandTyped, StatusMonotone, StatusIdempotent, StatusMinimalNonbase, StatusSaturatedFullPair, StatusZ2Equivariant, StatusTargetFunctorRecovered, StatusMeasureRewritten, StatusNativeSourceMissing, StatusAlphaR3StillBlocked}
}
func Supports() []string {
	return []string{SupportClosureExistsOnFiniteFlagChain, SupportEachDemandHasLeastSupport, SupportExistenceFromFiniteChain, SupportClosureExtensiveDemandLevel, SupportCl0ContainsBasepointDemand, SupportCl1ContainsExposureDemand, SupportCl2ContainsFullEnclosureDemand, SupportAirlockClosureMonotone, SupportDemandOrderMatchesSupportInclusion, SupportCl0SubsetCl1SubsetCl2, SupportAirlockClosureIdempotent, SupportF0F1F2ClosedSupports, SupportClosureTwiceNoChange, SupportCl1ByLeastNonbaseSupport, SupportSingletonDoesNotJumpToF2, SupportMinimalExposureForced, SupportCl2ByFullPairSaturation, SupportTopDemandRequiresSaturatedRectangle, SupportFullPairCannotCloseToF1, SupportAirlockClosureZ2Equivariant, SupportPhaseFlipCommutesWithClosure, SupportClosureDescendsToZ2Class, SupportThetaRecoveredFromClosure, SupportFixedBaseQuotientCumulativeTargets, SupportAssociatedGradedRejectedByF0Root, SupportMuBRewrittenUsingClosure, SupportAlphaReconstructedThroughClosure, SupportMuBTargetGapReducedToClosureStatus}
}
func Failures() []string {
	return []string{FailureExistenceBridgeNotNative, FailureExtensivityDemandTypedNotNative, FailureMonotonicityCandidateNotNative, FailureIdempotenceCandidateNotNative, FailureLeastNonbaseSupportNotNative, FailureFullPairSaturationNotNative, FailureZ2EquivarianceClosureNotNative, FailureFixedBaseQuotientBridgeNotNative, FailureNoNativeAirlockSupportClosure, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}
}

func (f FirewallLedger) List() []string {
	out := []string{}
	if f.ExistenceBridgeNotNative {
		out = append(out, FailureExistenceBridgeNotNative)
	}
	if f.ExtensivityDemandTypedNotNative {
		out = append(out, FailureExtensivityDemandTypedNotNative)
	}
	if f.MonotonicityCandidateNotNative {
		out = append(out, FailureMonotonicityCandidateNotNative)
	}
	if f.IdempotenceCandidateNotNative {
		out = append(out, FailureIdempotenceCandidateNotNative)
	}
	if f.LeastNonbaseSupportNotNative {
		out = append(out, FailureLeastNonbaseSupportNotNative)
	}
	if f.FullPairSaturationNotNative {
		out = append(out, FailureFullPairSaturationNotNative)
	}
	if f.Z2EquivarianceClosureNotNative {
		out = append(out, FailureZ2EquivarianceClosureNotNative)
	}
	if f.FixedBaseQuotientBridgeNotNative {
		out = append(out, FailureFixedBaseQuotientBridgeNotNative)
	}
	if f.NoNativeAirlockSupportClosure {
		out = append(out, FailureNoNativeAirlockSupportClosure)
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

func FormatSupportChain(a SupportChainAudit) string {
	return fmt.Sprintf("family=%s z2_family=%s finite=%t least_supports=%t native=%t supports=%s failures=%s", a.AdmissibleFamily, a.Z2AdmissibleFamily, a.FiniteChain, a.LeastSupportsExist, a.NativeClosureTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatExtensivity(a ExtensivityAudit) string {
	return fmt.Sprintf("demand_typed=%t cl0_base=%t cl1_exposure=%t cl2_enclosure=%t native_subspace=%t supports=%s failures=%s", a.DemandTyped, a.Cl0ContainsBasepoint, a.Cl1ContainsExposure, a.Cl2ContainsEnclosure, a.NativeSubspaceClosure, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatMonotonicity(a MonotonicityAudit) string {
	return fmt.Sprintf("demand_order=%s support_order=%s monotone=%t inclusion=%t native=%t supports=%s failures=%s", a.DemandOrder, a.SupportOrder, a.Monotone, a.SupportInclusion, a.NativeActionTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatIdempotence(a IdempotenceAudit) string {
	return fmt.Sprintf("image_admissible=%t closed=(%t,%t,%t) idempotent=%t native=%t supports=%s failures=%s", a.ImageInAdmissibleFamily, a.F0Closed, a.F1Closed, a.F2Closed, a.Idempotent, a.NativeClosureTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatMinimality(a MinimalityAudit) string {
	return fmt.Sprintf("k=%d target=%s least_nonbase=%t rejects_f2_jump=%t native=%t supports=%s failures=%s", a.DemandK, a.ClosureTarget, a.LeastNonbaseSupport, a.RejectsJumpToF2, a.NativeRule, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatSaturation(a SaturationAudit) string {
	return fmt.Sprintf("k=%d target=%s full_pair_saturation=%t rejects_f1=%t native=%t supports=%s failures=%s", a.DemandK, a.ClosureTarget, a.FullPairSaturation, a.RejectsCloseToF1, a.NativeRule, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatZ2(a Z2EquivarianceAudit) string {
	return fmt.Sprintf("lambda=%s barlambda=%s commutes=%t descends=%t native=%t supports=%s failures=%s", a.LambdaLadder, a.BarLambdaLadder, a.PhaseFlipCommutes, a.DescendsToZ2Class, a.NativePhaseTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatTargetRecovery(a TargetRecoveryAudit) string {
	return fmt.Sprintf("theta=%s targets=(%s,%s) ranks=(%d,%d) cumulative=%t rejects_associated=%t native_fixed_base=%t supports=%s failures=%s", a.ThetaFormula, a.ThetaOne, a.ThetaTwo, a.ThetaOneRank, a.ThetaTwoRank, a.CumulativeTargets, a.RejectsAssociatedGraded, a.NativeFixedBaseTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatMeasure(a MeasureConsequenceAudit) string {
	return fmt.Sprintf("measure=%s alpha=%s ranks=(%d,%d) chambers=(%d,%d) rewritten=%t reconstructed=%t native_alpha=%t supports=%s failures=%s", a.MeasureFormula, a.AlphaFormula, a.ThetaOneRank, a.ThetaTwoRank, a.H10Rank, a.H72Rank, a.RewrittenUsingClosure, a.AlphaReconstructed, a.NativeAlpha, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatFirewalls(f FirewallLedger) string { return strings.Join(f.List(), ",") }

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
func firewallsOK(f FirewallLedger) bool {
	return f.ExistenceBridgeNotNative && f.ExtensivityDemandTypedNotNative && f.MonotonicityCandidateNotNative && f.IdempotenceCandidateNotNative && f.LeastNonbaseSupportNotNative && f.FullPairSaturationNotNative && f.Z2EquivarianceClosureNotNative && f.FixedBaseQuotientBridgeNotNative && f.NoNativeAirlockSupportClosure && f.AlphaBridgeCandidateNotNative && f.NotNativeR3 && f.FullAFDescentStillBlocked && f.NoGenerationCarrierMap && f.NoFlavorOrientationMap && f.NoIndividualYukawaValues && f.NoNativeYukawaOperator
}
