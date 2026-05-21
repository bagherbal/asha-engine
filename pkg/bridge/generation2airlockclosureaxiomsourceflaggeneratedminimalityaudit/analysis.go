// Package generation2airlockclosureaxiomsourceflaggeneratedminimalityaudit implements
// Gate 929: AirlockClosure Axiom Source and Flag-Generated Minimality Audit.
//
// Gate 929 follows Gate 928's result that the Z2 airlock closure ladder is
// unique if basepoint, monotonicity, minimality, saturation, fixed-base
// quotienting, and Z2 representative-independence are accepted. This gate asks
// whether those axioms are sourced by the puncture-airlock flag itself. The
// result is a flag-generated least-support source typing, not a native closure
// operator theorem.
package generation2airlockclosureaxiomsourceflaggeneratedminimalityaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE929-AIRLOCK-CLOSURE-AXIOM-SOURCE-FLAG-GENERATED-MINIMALITY-AUDIT"

	Gate928ShortStatus = "R3_ALPHA_CLOSURE_UNIQUE_UNDER_AXIOMS_NATIVE_SOURCE_MISSING"

	BoundarySubsetChain = "0<1<2 boundary subset cardinality chain"
	AirlockFlagChain    = "F_0 subset F_1 subset F_2"
	Z2PunctureClass     = "[p]_{Z2}={e_lambda tensor P_1,e_barlambda tensor P_1}"
	CandidateOperator   = "AirlockSupportClosureOperator"
	LeastSupportRule    = "least admissible support closure above the neutral puncture"
	ClosureZero         = "Cl_airlock(0)=F_0"
	ClosureOne          = "Cl_airlock(1)=F_1"
	ClosureTwo          = "Cl_airlock(2)=F_2"
	ThetaViaClosure     = "Theta_B^Z2(k)=[Cl_airlock^Z2(k)/F_0]_{Z2}"
	AlphaViaClosure     = "mu_B(R_B(S_split))=(3/10)S_split+(7/72)S_split^2"
	NextGate            = "NEXT_PRESSURE_GATE930_AIRLOCKSUPPORTCLOSUREOPERATOR_EXISTENCE_AND_IDEMPOTENCE_AUDIT"

	RankF0       = 1
	RankF1       = 4
	RankF2       = 8
	RankF1OverF0 = 3
	RankF2OverF0 = 7
	RankH10      = 10
	RankH72      = 72

	Classification = "R3_AIRLOCK_CLOSURE_AXIOMS_FLAG_SOURCED_NOT_NATIVE"
	ShortStatus    = "R3_ALPHA_CLOSURE_AXIOMS_SOURCED_TO_LEAST_SUPPORT_OPERATOR_GAP"
	FinalTruth     = "AIRLOCK_CLOSURE_AXIOMS_SOURCE_TYPED_BY_FLAG_GENERATED_LEAST_SUPPORT_COMPLETION_BUT_NATIVE_CLOSURE_OPERATOR_MISSING"

	StatusInheritedGate928          = "PASS_GATE928_CLOSURE_UNIQUE_UNDER_AXIOMS_INHERITED"
	StatusBasepointFlagGenerated    = "PASS_BASEPOINT_AXIOM_SOURCED_BY_PUNCTURE_INITIALITY"
	StatusMonotonicityFlagGenerated = "PASS_MONOTONICITY_SOURCED_BY_SUPPORT_INCLUSION"
	StatusMinimalityFlagGenerated   = "PASS_MINIMALITY_SOURCED_BY_LEAST_SAME_SOCKET_COMPLETION"
	StatusSaturationFlagGenerated   = "PASS_SATURATION_SOURCED_BY_FULL_RIGHT_RECTANGLE_COMPLETION"
	StatusFixedBaseFlagGenerated    = "PASS_FIXED_BASE_QUOTIENT_SOURCED_BY_RELATIVE_ACTIVATION_ABOVE_PUNCTURE"
	StatusZ2FlagGenerated           = "PASS_Z2_INVARIANCE_SOURCED_BY_CLASS_LEVEL_SUPPORT_CLOSURE"
	StatusAxiomLedger               = "PASS_CLOSURE_AXIOMS_ARE_FLAG_GENERATED_NOT_ARBITRARY"
	StatusMeasureConsequence        = "PASS_BOUNDARY_ACTIVATION_MEASURE_TARGETS_FIXED_BY_FLAG_GENERATED_CLOSURE"
	StatusNativeClosureMissing      = "FIREWALL_PRESERVED_NATIVE_AIRLOCK_SUPPORT_CLOSURE_OPERATOR_MISSING"
	StatusAlphaR3StillBlocked       = "FIREWALL_PRESERVED_ALPHA_AND_R3_NOT_NATIVE"

	SupportBasepointByPunctureInitiality       = "CONDITIONAL_SUPPORT_BASEPOINT_AXIOM_SOURCED_BY_PUNCTURE_INITIALITY"
	SupportEmptyClosesMinimalSupport           = "CONDITIONAL_SUPPORT_EMPTY_BOUNDARY_SUBSET_CLOSES_TO_MINIMAL_AIRLOCK_SUPPORT"
	SupportCl0FlagGenerated                    = "CONDITIONAL_SUPPORT_CL_0_EQUALS_F0_IS_FLAG_GENERATED"
	SupportMonotonicityBySupportInclusion      = "CONDITIONAL_SUPPORT_MONOTONICITY_SOURCED_BY_SUPPORT_INCLUSION"
	SupportBoundaryAirlockOrdersCompatible     = "CONDITIONAL_SUPPORT_BOUNDARY_SUBSET_ORDER_AND_AIRLOCK_FLAG_ORDER_ARE_COMPATIBLE"
	SupportClosureMonotonicityFlagNatural      = "CONDITIONAL_SUPPORT_CLOSURE_MONOTONICITY_IS_FLAG_NATURAL"
	SupportMinimalityBySameSocketCompletion    = "CONDITIONAL_SUPPORT_MINIMALITY_AXIOM_SOURCED_BY_LEAST_SAME_SOCKET_COMPLETION"
	SupportSingletonClosesExposedFace          = "CONDITIONAL_SUPPORT_SINGLETON_ACTIVATION_CLOSES_TO_EXPOSED_PHASE_FACE"
	SupportCl1ForcedByMinimalSupport           = "CONDITIONAL_SUPPORT_CL_1_EQUALS_F1_IS_FORCED_BY_MINIMAL_NONTRIVIAL_SUPPORT_ABOVE_PUNCTURE"
	SupportSaturationByFullPairActivation      = "CONDITIONAL_SUPPORT_SATURATION_AXIOM_SOURCED_BY_FULL_BOUNDARY_PAIR_ACTIVATION"
	SupportTopDegreeClosesFullRectangle        = "CONDITIONAL_SUPPORT_TOP_EXTERIOR_DEGREE_CLOSES_TO_FULL_RIGHT_RECTANGLE"
	SupportCl2ForcedBySaturatedCompletion      = "CONDITIONAL_SUPPORT_CL_2_EQUALS_F2_IS_FORCED_BY_SATURATED_SUPPORT_COMPLETION"
	SupportFixedBaseByRelativeActivation       = "CONDITIONAL_SUPPORT_FIXED_BASE_QUOTIENT_SOURCED_BY_RELATIVE_ACTIVATION_ABOVE_PUNCTURE"
	SupportCumulativeQuotientFlagGenerated     = "CONDITIONAL_SUPPORT_CUMULATIVE_QUOTIENT_IS_FLAG_GENERATED"
	SupportAssociatedGradedRejectedByBasepoint = "CONDITIONAL_SUPPORT_ASSOCIATED_GRADED_SLICE_REJECTED_BY_RELATIVE_BASEPOINT_MEASURE"
	SupportZ2ByClassLevelClosure               = "CONDITIONAL_SUPPORT_Z2_INVARIANCE_SOURCED_BY_CLASS_LEVEL_SUPPORT_CLOSURE"
	SupportPhaseFlipExchangesRepresentatives   = "CONDITIONAL_SUPPORT_PHASE_FLIP_EXCHANGES_CLOSURE_REPRESENTATIVES"
	SupportClosureQuotientRanksZ2Invariant     = "CONDITIONAL_SUPPORT_CLOSURE_QUOTIENT_RANKS_ARE_Z2_INVARIANT"
	SupportAxiomsFlagGenerated                 = "CONDITIONAL_SUPPORT_CLOSURE_AXIOMS_ARE_FLAG_GENERATED_NOT_ARBITRARY"
	SupportThetaReconstructedFromFlagClosure   = "CONDITIONAL_SUPPORT_THETA_B_Z2_RECONSTRUCTED_FROM_FLAG_GENERATED_CLOSURE"
	SupportMeasureTargetsFixedByFlagClosure    = "CONDITIONAL_SUPPORT_BOUNDARY_ACTIVATION_MEASURE_TARGETS_FIXED_BY_FLAG_GENERATED_CLOSURE"

	FailureNoNativeAirlockSupportClosure       = "FAILED_ROUTE_NO_NATIVE_AIRLOCK_SUPPORT_CLOSURE_OPERATOR"
	FailurePunctureInitialityNotNative         = "FAILED_ROUTE_PUNCTURE_INITIALITY_NOT_YET_NATIVE_CLOSURE_OPERATOR_THEOREM"
	FailureSupportInclusionNotNative           = "FAILED_ROUTE_SUPPORT_INCLUSION_COMPATIBILITY_NOT_NATIVE_BOUNDARY_ACTION_THEOREM"
	FailureLeastSameSocketNotNative            = "FAILED_ROUTE_LEAST_SAME_SOCKET_COMPLETION_NOT_NATIVE_AIRLOCK_CLOSURE_THEOREM"
	FailureFullRightRectangleNotNative         = "FAILED_ROUTE_FULL_RIGHT_RECTANGLE_SATURATION_NOT_NATIVE_CLOSURE_THEOREM"
	FailureRelativeActivationQuotientNotNative = "FAILED_ROUTE_RELATIVE_ACTIVATION_QUOTIENT_NOT_NATIVE_MEASURE_THEOREM"
	FailureZ2ClassClosureNotNative             = "FAILED_ROUTE_Z2_CLASS_SUPPORT_CLOSURE_NOT_NATIVE_GLOBAL_PHASE_THEOREM"
	FailureAxiomsFlagSourcedNotNative          = "FAILED_ROUTE_CLOSURE_AXIOMS_FLAG_SOURCED_BUT_NOT_NATIVE_ASHA_THEOREMS"
	FailureMuBNotNativeWithoutSupportClosure   = "FAILED_ROUTE_MU_B_STILL_NOT_NATIVE_WITHOUT_NATIVE_AIRLOCK_SUPPORT_CLOSURE_OPERATOR"
	FailureAlphaBridgeCandidateNotNative       = "FAILED_ROUTE_ALPHA_B_REMAINS_BRIDGE_CANDIDATE_NOT_NATIVE"
	FailureNotNativeR3                         = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked           = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap              = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap              = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues            = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator              = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type BasepointSourceAudit struct {
	PunctureInitiality   bool
	EmptyBoundarySubset  bool
	ClosureTarget        string
	FlagGenerated        bool
	NativeClosureTheorem bool
	Supports             []string
	Failures             []string
}

type MonotonicitySourceAudit struct {
	BoundaryOrder       string
	AirlockOrder        string
	SupportInclusion    bool
	OrdersCompatible    bool
	FlagNatural         bool
	NativeActionTheorem bool
	Supports            []string
	Failures            []string
}

type MinimalitySourceAudit struct {
	SingletonActivation    bool
	SameSocketCompletion   string
	ClosureTarget          string
	ExposedFace            bool
	ForcedByMinimalSupport bool
	NativeClosureTheorem   bool
	Supports               []string
	Failures               []string
}

type SaturationSourceAudit struct {
	FullPairActivation   bool
	TopExteriorDegree    bool
	SaturatedCompletion  string
	ClosureTarget        string
	ForcedBySaturation   bool
	NativeClosureTheorem bool
	Supports             []string
	Failures             []string
}

type FixedBaseSourceAudit struct {
	AbsoluteClosures        string
	RelativeTargets         string
	UsesFixedBaseF0         bool
	CumulativeQuotient      bool
	RejectsAssociatedGraded bool
	NativeMeasureTheorem    bool
	Supports                []string
	Failures                []string
}

type Z2SourceAudit struct {
	LambdaClosureLadder      string
	BarLambdaClosureLadder   string
	PhaseFlipExchanges       bool
	ClassLevelClosure        bool
	RanksInvariant           bool
	NativeGlobalPhaseTheorem bool
	Supports                 []string
	Failures                 []string
}

type ClosureAxiomLedgerAudit struct {
	BasepointStatus      string
	MonotonicityStatus   string
	MinimalityStatus     string
	SaturationStatus     string
	FixedBaseStatus      string
	Z2Status             string
	FlagGenerated        bool
	NativeOperatorExists bool
	Supports             []string
	Failures             []string
}

type MeasureConsequenceAudit struct {
	ThetaFormula       string
	ThetaOneRank       int
	ThetaTwoRank       int
	H10Rank            int
	H72Rank            int
	AlphaFormula       string
	ThetaReconstructed bool
	TargetsFixed       bool
	NativeAlpha        bool
	Supports           []string
	Failures           []string
}

type FirewallLedger struct {
	NoNativeAirlockSupportClosure       bool
	PunctureInitialityNotNative         bool
	SupportInclusionNotNative           bool
	LeastSameSocketNotNative            bool
	FullRightRectangleNotNative         bool
	RelativeActivationQuotientNotNative bool
	Z2ClassClosureNotNative             bool
	AxiomsFlagSourcedNotNative          bool
	MuBNotNativeWithoutSupportClosure   bool
	AlphaBridgeCandidateNotNative       bool
	NotNativeR3                         bool
	FullAFDescentStillBlocked           bool
	NoGenerationCarrierMap              bool
	NoFlavorOrientationMap              bool
	NoIndividualYukawaValues            bool
	NoNativeYukawaOperator              bool
}

type Audit struct {
	ID             string
	Inherited      string
	Basepoint      BasepointSourceAudit
	Monotonicity   MonotonicitySourceAudit
	Minimality     MinimalitySourceAudit
	Saturation     SaturationSourceAudit
	FixedBase      FixedBaseSourceAudit
	Z2             Z2SourceAudit
	Ledger         ClosureAxiomLedgerAudit
	Measure        MeasureConsequenceAudit
	Firewalls      FirewallLedger
	Truth          string
	Classification string
	ShortStatus    string
	Final          string
}

func BuildDefault() (Audit, error) {
	return Audit{
		ID:           AuditID,
		Inherited:    Gate928ShortStatus,
		Basepoint:    BasepointSourceAudit{PunctureInitiality: true, EmptyBoundarySubset: true, ClosureTarget: "F_0", FlagGenerated: true, NativeClosureTheorem: false, Supports: []string{SupportBasepointByPunctureInitiality, SupportEmptyClosesMinimalSupport, SupportCl0FlagGenerated}, Failures: []string{FailurePunctureInitialityNotNative}},
		Monotonicity: MonotonicitySourceAudit{BoundaryOrder: BoundarySubsetChain, AirlockOrder: AirlockFlagChain, SupportInclusion: true, OrdersCompatible: true, FlagNatural: true, NativeActionTheorem: false, Supports: []string{SupportMonotonicityBySupportInclusion, SupportBoundaryAirlockOrdersCompatible, SupportClosureMonotonicityFlagNatural}, Failures: []string{FailureSupportInclusionNotNative}},
		Minimality:   MinimalitySourceAudit{SingletonActivation: true, SameSocketCompletion: "F_1=e_phase tensor W=P_1 plus P_3 in the same phase socket", ClosureTarget: "F_1", ExposedFace: true, ForcedByMinimalSupport: true, NativeClosureTheorem: false, Supports: []string{SupportMinimalityBySameSocketCompletion, SupportSingletonClosesExposedFace, SupportCl1ForcedByMinimalSupport}, Failures: []string{FailureLeastSameSocketNotNative}},
		Saturation:   SaturationSourceAudit{FullPairActivation: true, TopExteriorDegree: true, SaturatedCompletion: "F_2=C_R^2 tensor W", ClosureTarget: "F_2", ForcedBySaturation: true, NativeClosureTheorem: false, Supports: []string{SupportSaturationByFullPairActivation, SupportTopDegreeClosesFullRectangle, SupportCl2ForcedBySaturatedCompletion}, Failures: []string{FailureFullRightRectangleNotNative}},
		FixedBase:    FixedBaseSourceAudit{AbsoluteClosures: "F_0,F_1,F_2", RelativeTargets: "[F_1/F_0]_{Z2},[F_2/F_0]_{Z2}", UsesFixedBaseF0: true, CumulativeQuotient: true, RejectsAssociatedGraded: true, NativeMeasureTheorem: false, Supports: []string{SupportFixedBaseByRelativeActivation, SupportCumulativeQuotientFlagGenerated, SupportAssociatedGradedRejectedByBasepoint}, Failures: []string{FailureRelativeActivationQuotientNotNative}},
		Z2:           Z2SourceAudit{LambdaClosureLadder: "Cl_lambda(0)=F_0^lambda; Cl_lambda(1)=F_1^lambda; Cl_lambda(2)=F_2", BarLambdaClosureLadder: "Cl_barlambda(0)=F_0^barlambda; Cl_barlambda(1)=F_1^barlambda; Cl_barlambda(2)=F_2", PhaseFlipExchanges: true, ClassLevelClosure: true, RanksInvariant: true, NativeGlobalPhaseTheorem: false, Supports: []string{SupportZ2ByClassLevelClosure, SupportPhaseFlipExchangesRepresentatives, SupportClosureQuotientRanksZ2Invariant}, Failures: []string{FailureZ2ClassClosureNotNative}},
		Ledger:       ClosureAxiomLedgerAudit{BasepointStatus: "bridge-strong", MonotonicityStatus: "bridge-strong", MinimalityStatus: "strongest new source", SaturationStatus: "strong", FixedBaseStatus: "bridge-strong", Z2Status: "bridge-strong", FlagGenerated: true, NativeOperatorExists: false, Supports: []string{SupportAxiomsFlagGenerated}, Failures: []string{FailureNoNativeAirlockSupportClosure, FailureAxiomsFlagSourcedNotNative}},
		Measure:      MeasureConsequenceAudit{ThetaFormula: ThetaViaClosure, ThetaOneRank: RankF1OverF0, ThetaTwoRank: RankF2OverF0, H10Rank: RankH10, H72Rank: RankH72, AlphaFormula: AlphaViaClosure, ThetaReconstructed: true, TargetsFixed: true, NativeAlpha: false, Supports: []string{SupportThetaReconstructedFromFlagClosure, SupportMeasureTargetsFixedByFlagClosure}, Failures: []string{FailureMuBNotNativeWithoutSupportClosure, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3}},
		Firewalls:    FirewallLedger{NoNativeAirlockSupportClosure: true, PunctureInitialityNotNative: true, SupportInclusionNotNative: true, LeastSameSocketNotNative: true, FullRightRectangleNotNative: true, RelativeActivationQuotientNotNative: true, Z2ClassClosureNotNative: true, AxiomsFlagSourcedNotNative: true, MuBNotNativeWithoutSupportClosure: true, AlphaBridgeCandidateNotNative: true, NotNativeR3: true, FullAFDescentStillBlocked: true, NoGenerationCarrierMap: true, NoFlavorOrientationMap: true, NoIndividualYukawaValues: true, NoNativeYukawaOperator: true},
		Truth:        FinalTruth, Classification: Classification, ShortStatus: ShortStatus,
		Final: "Gate 929 source-types the closure axioms by flag-generated least-support completion but preserves the missing native AirlockSupportClosureOperator.",
	}, nil
}

func Statuses() []string {
	return []string{StatusInheritedGate928, StatusBasepointFlagGenerated, StatusMonotonicityFlagGenerated, StatusMinimalityFlagGenerated, StatusSaturationFlagGenerated, StatusFixedBaseFlagGenerated, StatusZ2FlagGenerated, StatusAxiomLedger, StatusMeasureConsequence, StatusNativeClosureMissing, StatusAlphaR3StillBlocked}
}
func Supports() []string {
	return []string{SupportBasepointByPunctureInitiality, SupportEmptyClosesMinimalSupport, SupportCl0FlagGenerated, SupportMonotonicityBySupportInclusion, SupportBoundaryAirlockOrdersCompatible, SupportClosureMonotonicityFlagNatural, SupportMinimalityBySameSocketCompletion, SupportSingletonClosesExposedFace, SupportCl1ForcedByMinimalSupport, SupportSaturationByFullPairActivation, SupportTopDegreeClosesFullRectangle, SupportCl2ForcedBySaturatedCompletion, SupportFixedBaseByRelativeActivation, SupportCumulativeQuotientFlagGenerated, SupportAssociatedGradedRejectedByBasepoint, SupportZ2ByClassLevelClosure, SupportPhaseFlipExchangesRepresentatives, SupportClosureQuotientRanksZ2Invariant, SupportAxiomsFlagGenerated, SupportThetaReconstructedFromFlagClosure, SupportMeasureTargetsFixedByFlagClosure}
}
func Failures() []string {
	return []string{FailureNoNativeAirlockSupportClosure, FailurePunctureInitialityNotNative, FailureSupportInclusionNotNative, FailureLeastSameSocketNotNative, FailureFullRightRectangleNotNative, FailureRelativeActivationQuotientNotNative, FailureZ2ClassClosureNotNative, FailureAxiomsFlagSourcedNotNative, FailureMuBNotNativeWithoutSupportClosure, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}
}

func (f FirewallLedger) List() []string {
	out := []string{}
	if f.NoNativeAirlockSupportClosure {
		out = append(out, FailureNoNativeAirlockSupportClosure)
	}
	if f.PunctureInitialityNotNative {
		out = append(out, FailurePunctureInitialityNotNative)
	}
	if f.SupportInclusionNotNative {
		out = append(out, FailureSupportInclusionNotNative)
	}
	if f.LeastSameSocketNotNative {
		out = append(out, FailureLeastSameSocketNotNative)
	}
	if f.FullRightRectangleNotNative {
		out = append(out, FailureFullRightRectangleNotNative)
	}
	if f.RelativeActivationQuotientNotNative {
		out = append(out, FailureRelativeActivationQuotientNotNative)
	}
	if f.Z2ClassClosureNotNative {
		out = append(out, FailureZ2ClassClosureNotNative)
	}
	if f.AxiomsFlagSourcedNotNative {
		out = append(out, FailureAxiomsFlagSourcedNotNative)
	}
	if f.MuBNotNativeWithoutSupportClosure {
		out = append(out, FailureMuBNotNativeWithoutSupportClosure)
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

func FormatBasepoint(a BasepointSourceAudit) string {
	return fmt.Sprintf("target=%s puncture_initial=%t empty_subset=%t flag_generated=%t native=%t supports=%s failures=%s", a.ClosureTarget, a.PunctureInitiality, a.EmptyBoundarySubset, a.FlagGenerated, a.NativeClosureTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatMonotonicity(a MonotonicitySourceAudit) string {
	return fmt.Sprintf("boundary_order=%s airlock_order=%s support_inclusion=%t compatible=%t flag_natural=%t native=%t supports=%s failures=%s", a.BoundaryOrder, a.AirlockOrder, a.SupportInclusion, a.OrdersCompatible, a.FlagNatural, a.NativeActionTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatMinimality(a MinimalitySourceAudit) string {
	return fmt.Sprintf("singleton=%t same_socket=%s target=%s exposed=%t forced=%t native=%t supports=%s failures=%s", a.SingletonActivation, a.SameSocketCompletion, a.ClosureTarget, a.ExposedFace, a.ForcedByMinimalSupport, a.NativeClosureTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatSaturation(a SaturationSourceAudit) string {
	return fmt.Sprintf("full_pair=%t top_degree=%t completion=%s target=%s forced=%t native=%t supports=%s failures=%s", a.FullPairActivation, a.TopExteriorDegree, a.SaturatedCompletion, a.ClosureTarget, a.ForcedBySaturation, a.NativeClosureTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatFixedBase(a FixedBaseSourceAudit) string {
	return fmt.Sprintf("closures=%s targets=%s fixed_f0=%t cumulative=%t rejects_associated=%t native=%t supports=%s failures=%s", a.AbsoluteClosures, a.RelativeTargets, a.UsesFixedBaseF0, a.CumulativeQuotient, a.RejectsAssociatedGraded, a.NativeMeasureTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatZ2(a Z2SourceAudit) string {
	return fmt.Sprintf("lambda=%s barlambda=%s phase_exchanges=%t class_level=%t ranks_invariant=%t native=%t supports=%s failures=%s", a.LambdaClosureLadder, a.BarLambdaClosureLadder, a.PhaseFlipExchanges, a.ClassLevelClosure, a.RanksInvariant, a.NativeGlobalPhaseTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatLedger(a ClosureAxiomLedgerAudit) string {
	return fmt.Sprintf("basepoint=%s monotonicity=%s minimality=%s saturation=%s fixed_base=%s z2=%s flag_generated=%t native_operator=%t supports=%s failures=%s", a.BasepointStatus, a.MonotonicityStatus, a.MinimalityStatus, a.SaturationStatus, a.FixedBaseStatus, a.Z2Status, a.FlagGenerated, a.NativeOperatorExists, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatMeasure(a MeasureConsequenceAudit) string {
	return fmt.Sprintf("theta=%s ranks=(%d,%d) chambers=(%d,%d) alpha=%s theta_reconstructed=%t targets_fixed=%t native_alpha=%t supports=%s failures=%s", a.ThetaFormula, a.ThetaOneRank, a.ThetaTwoRank, a.H10Rank, a.H72Rank, a.AlphaFormula, a.ThetaReconstructed, a.TargetsFixed, a.NativeAlpha, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
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
	return f.NoNativeAirlockSupportClosure && f.PunctureInitialityNotNative && f.SupportInclusionNotNative && f.LeastSameSocketNotNative && f.FullRightRectangleNotNative && f.RelativeActivationQuotientNotNative && f.Z2ClassClosureNotNative && f.AxiomsFlagSourcedNotNative && f.MuBNotNativeWithoutSupportClosure && f.AlphaBridgeCandidateNotNative && f.NotNativeR3 && f.FullAFDescentStillBlocked && f.NoGenerationCarrierMap && f.NoFlavorOrientationMap && f.NoIndividualYukawaValues && f.NoNativeYukawaOperator
}
