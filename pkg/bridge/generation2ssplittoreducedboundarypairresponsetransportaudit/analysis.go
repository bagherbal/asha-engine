// Package generation2ssplittoreducedboundarypairresponsetransportaudit implements
// Gate 916: S_split to Reduced BoundaryPair Response Transport Audit.
//
// Gate 916 follows Gate 915's Z2 BoundaryAlpha cross-lane exclusion audit.
// It audits the fourth Gate 912 sub-object: whether the boundary split
// coordinate S_split can be transported as the scalar response parameter s
// inside the reduced rank-two boundary-pair response
//
//	R_B(s)=(1+s b1)(1+s b2)-1.
//
// The key sharpening is that s^2 is not independently transported. The scalar
// s is inserted once into each boundary factor, and the exterior product
// (s b1)(s b2)=s^2(b1 wedge b2) produces the quadratic term. The audit supports
// this transport shape but does not certify a native T_s map, a native reason
// for uniform insertion into both boundary generators, denominator normalization,
// BoundaryAlpha, or native R3.
package generation2ssplittoreducedboundarypairresponsetransportaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE916-S-SPLIT-TO-REDUCED-BOUNDARYPAIR-RESPONSE-TRANSPORT-AUDIT"

	SBoundary   = 0.0012924448188162962
	AlphaB      = 0.0003878958469680527
	AlphaLinear = 0.00038773344564488885
	AlphaQuad   = 0.0000001624013231638281

	RankF1OverF0   = 3
	RankF2OverF0   = 7
	LinearDenom    = 10
	QuadraticDenom = 72

	Gate913Classification = "R3_REDUCED_B2_RESPONSE_FUNCTIONAL_SHAPE_CERTIFIED_NOT_NATIVE_BOUNDARY_ALPHA"
	Gate913ShortStatus    = "R3_ALPHA_SUBOBJECT_1_REDUCED_B2_RESPONSE_SHAPE_PASS_NATIVE_SELECTION_BLOCKED"
	Gate914Classification = "R3_ALPHA_SUBOBJECT_2_Z2_FLAG_SELECTOR_SHAPE_PASS_NATIVE_FUNCTOR_BLOCKED"
	Gate914ShortStatus    = "R3_DEGREE_INDEXED_Z2_AIRLOCK_FLAG_SELECTOR_OBSTRUCTION"
	Gate915Classification = "R3_ALPHA_SUBOBJECT_3_CROSS_LANE_EXCLUSION_SHAPE_PASS_NATIVE_UNIQUENESS_BLOCKED"
	Gate915ShortStatus    = "R3_Z2_BOUNDARYALPHA_CROSS_LANE_EXCLUSION_OBSTRUCTION"

	TransportSource          = "S_split"
	TransportTarget          = "scalar parameter s in R_B(s)"
	TransportMapCandidate    = "T_s(S_split)=s"
	ReducedResponse          = "R_B(s)=(1+s b1)(1+s b2)-1"
	ExpandedReducedResponse  = "R_B(s)=s(b1+b2)+s^2(b1 wedge b2)"
	BoundaryFactor1          = "1+s b1"
	BoundaryFactor2          = "1+s b2"
	QuadraticExteriorProduct = "(s b1)(s b2)=s^2(b1 wedge b2)"
	IdentityTerm             = "1 in Lambda^0 B_2"
	ActiveBoundaryTerms      = "s b1, s b2"
	AlphaFormula             = "alpha_B=(3/10)s+(7/72)s^2"
	SelectorPipeline         = "S_split -> R_B(s) powers -> degree selector -> alpha rank lanes"
	NextGate                 = "NEXT_PRESSURE_GATE917_BOUNDARYAUGMENTED_RESPONSECHAMBER_NORMALIZATION_AUDIT"
	Classification           = "R3_ALPHA_SUBOBJECT_4_S_SPLIT_TRANSPORT_SHAPE_PASS_NATIVE_TRANSPORT_BLOCKED"
	ShortStatus              = "R3_S_SPLIT_TO_REDUCED_B2_RESPONSE_TRANSPORT_OBSTRUCTION"
	FinalTruth               = "S_SPLIT_TRANSPORT_COMPATIBLE_WITH_REDUCED_B2_RESPONSE_AS_SINGLE_UNIFORM_BOUNDARY_FACTOR_INSERTION_BUT_NATIVE_TRANSPORT_MAP_MISSING"
	StrategicConclusion      = "S_split transport compresses to one scalar insertion into each boundary factor; s^2 is produced by exterior multiplication, but the native T_s map and uniform-insertion law remain missing."

	StatusGates913To915Inherited       = "PASS_GATES913_914_915_SHAPE_SELECTOR_CROSSLANE_RESULTS_INHERITED"
	StatusTransportTargetTyped         = "PASS_S_SPLIT_TARGET_TYPED_AS_REDUCED_B2_RESPONSE_PARAMETER"
	StatusSingleInsertion              = "PASS_SINGLE_S_SPLIT_INSERTION_GENERATES_LINEAR_AND_QUADRATIC_RESPONSE"
	StatusDimensionlessCompatibility   = "PASS_S_SPLIT_SCALAR_RESPONSE_PARAMETER_COMPATIBLE_WITH_BOUNDARY_GENERATORS"
	StatusActiveBasepointReduction     = "PASS_S_SPLIT_APPLIES_TO_ACTIVE_TERMS_IDENTITY_IS_BASEPOINT_REMOVED_BY_REDUCTION"
	StatusSelectorCompatibility        = "PASS_S_SPLIT_RESPONSE_COMPATIBLE_WITH_SELECTOR_AND_CROSSLANE_FIREWALL"
	StatusAlphaReconstruction          = "PASS_ALPHA_RECONSTRUCTED_UNDER_TRANSPORT_SEAL_AND_PRIOR_SUBOBJECTS"
	StatusNativeTransportMapStillBlock = "FIREWALL_PRESERVED_NATIVE_S_SPLIT_TRANSPORT_MAP_NOT_CERTIFIED"

	SupportTransportTargetIsResponseParameter       = "CONDITIONAL_SUPPORT_S_SPLIT_TRANSPORT_TARGET_IS_REDUCED_B2_RESPONSE_PARAMETER"
	SupportSSplitScalarResponseParameter            = "CONDITIONAL_SUPPORT_S_SPLIT_IS_USED_AS_SCALAR_BOUNDARY_RESPONSE_PARAMETER"
	SupportAlphaUsesSSplitOnlyThroughRB             = "CONDITIONAL_SUPPORT_ALPHA_USES_S_SPLIT_ONLY_THROUGH_R_B_OF_S"
	SupportSingleInsertionGeneratesPowers           = "CONDITIONAL_SUPPORT_SINGLE_S_SPLIT_INSERTION_INTO_EACH_BOUNDARY_FACTOR_GENERATES_S_AND_S_SQUARED"
	SupportS2FromExteriorProduct                    = "CONDITIONAL_SUPPORT_S_SQUARED_TERM_ARISES_FROM_EXTERIOR_PRODUCT_NOT_SEPARATE_TRANSPORT"
	SupportPowerResponseUniformInsertion            = "CONDITIONAL_SUPPORT_POWER_RESPONSE_REDUCES_TO_UNIFORM_BOUNDARY_FACTOR_INSERTION"
	SupportSSplitScalarType                         = "CONDITIONAL_SUPPORT_S_SPLIT_HAS_SCALAR_RESPONSE_PARAMETER_TYPE"
	SupportDimensionlessMultipliesExteriorGenerator = "CONDITIONAL_SUPPORT_DIMENSIONLESS_S_SPLIT_CAN_MULTIPLY_BOUNDARY_EXTERIOR_GENERATORS"
	SupportTransportAppliesToActiveTerms            = "CONDITIONAL_SUPPORT_S_SPLIT_TRANSPORT_APPLIES_TO_ACTIVE_BOUNDARY_GENERATOR_TERMS"
	SupportIdentityBasepointRemoved                 = "CONDITIONAL_SUPPORT_IDENTITY_TERM_IS_BASEPOINT_AND_REMOVED_BY_REDUCTION"
	SupportCompatibleWithDegreeSelector             = "CONDITIONAL_SUPPORT_S_SPLIT_TRANSPORT_IS_COMPATIBLE_WITH_DEGREE_INDEXED_SELECTOR"
	SupportSingleInsertionFeedsCorrectAlphaLanes    = "CONDITIONAL_SUPPORT_SINGLE_INSERTION_RESPONSE_FEEDS_CORRECT_ALPHA_LANES_UNDER_SELECTOR"
	SupportTransportDoesNotReopenCrossLanes         = "CONDITIONAL_SUPPORT_S_SPLIT_TRANSPORT_DOES_NOT_REOPEN_CROSS_LANE_POLLUTION"
	SupportAlphaReconstructedGivenTransport         = "CONDITIONAL_SUPPORT_ALPHA_B_RECONSTRUCTED_GIVEN_S_SPLIT_TRANSPORT_AND_PRIOR_SUBOBJECTS"
	SupportTransportWoundReducedToNativeTs          = "CONDITIONAL_SUPPORT_S_SPLIT_TRANSPORT_WOUND_REDUCES_TO_NATIVE_SOURCE_OF_T_S"

	FailureNoNativeTsMap                     = "FAILED_ROUTE_NO_NATIVE_T_S_MAP_FROM_S_SPLIT_TO_BOUNDARY_RESPONSE_PARAMETER"
	FailureNoNativeTransportToZ2Airlock      = "FAILED_ROUTE_NO_NATIVE_TRANSPORT_OF_S_SPLIT_TO_Z2_AIRLOCK_CLASS"
	FailureNoTypedSSplitExteriorParameterMap = "FAILED_ROUTE_NO_TYPED_S_SPLIT_TO_BOUNDARY_PAIR_EXTERIOR_PARAMETER_MAP"
	FailureNoNativeUniformInsertionReason    = "FAILED_ROUTE_NO_NATIVE_REASON_YET_FOR_UNIFORM_INSERTION_INTO_BOTH_BOUNDARY_FACTORS"
	FailureSSplitScalarTypeSealed            = "FAILED_ROUTE_S_SPLIT_SCALAR_TYPE_IS_SEALED_NOT_NATIVE_BOUNDARY_TRANSPORT_THEOREM"
	FailureNoNativeBasepointReduction        = "FAILED_ROUTE_NO_NATIVE_BASEPOINT_REDUCTION_THEOREM_FOR_BOUNDARY_ALPHA"
	FailureSelectorCompatibilityNotNative    = "FAILED_ROUTE_COMPATIBILITY_WITH_SELECTOR_NOT_NATIVE_TRANSPORT_THEOREM"
	FailureAlphaReconstructionNotNative      = "FAILED_ROUTE_ALPHA_RECONSTRUCTION_UNDER_TRANSPORT_SEAL_NOT_NATIVE_ALPHA_THEOREM"
	FailureDenominatorNormalizationExternal  = "FAILED_ROUTE_DENOMINATOR_NORMALIZATION_STILL_EXTERNAL"
	FailureAlphaStillSealed                  = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNotNativeR3                       = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked         = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap            = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap            = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues          = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator            = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type InheritedSubobjects struct {
	Gate913Classification  string
	Gate913ShortStatus     string
	Gate914Classification  string
	Gate914ShortStatus     string
	Gate915Classification  string
	Gate915ShortStatus     string
	ResponseShapeCertified bool
	SelectorShapeTyped     bool
	CrossLanesBlocked      bool
	DerivesAlpha           bool
	UpdatesOfficialLedger  bool
	PromotesNativeR3       bool
	Supports, Failures     []string
}

type ScalarTransportTarget struct {
	Source               string
	Target               string
	MapCandidate         string
	TargetsAlphaDirectly bool
	TargetsSocketMag     bool
	UsesRBOnly           bool
	NativeMap            bool
	Supports, Failures   []string
}

type UniformInsertion struct {
	Factor1                    string
	Factor2                    string
	Insertions                 int
	ScalarInsertedPerFactor    bool
	SeparateQuadraticTransport bool
	QuadraticFromProduct       bool
	ExpandedResponse           string
	NativeUniformLaw           bool
	Supports, Failures         []string
}

type ScalarCompatibility struct {
	Parameter             float64
	Dimensionless         bool
	CanMultiplyGenerators bool
	ScalarTypeNative      bool
	Supports, Failures    []string
}

type ActiveResponseReduction struct {
	UnreducedResponse        string
	IdentityTerm             string
	ActiveTerms              string
	Reduction                string
	TransportAppliesToActive bool
	IdentityTransported      bool
	BasepointRemoved         bool
	NativeBasepointTheorem   bool
	Supports, Failures       []string
}

type SelectorCompatibility struct {
	Pipeline                  string
	DegreeOnePower            string
	DegreeTwoPower            string
	DegreeSelectorCompatible  bool
	FeedsCorrectAlphaLanes    bool
	ReopensCrossLanePollution bool
	NativeTransportTheorem    bool
	Supports, Failures        []string
}

type AlphaUnderTransport struct {
	S                      float64
	AlphaLinear            float64
	AlphaQuadratic         float64
	AlphaTotal             float64
	Formula                string
	RankPair               [2]int
	TransportSealAssumed   bool
	PriorSubobjectsAssumed bool
	NativeAlphaTheorem     bool
	Supports, Failures     []string
}

type Firewalls struct {
	NativeTsMap                  bool
	NativeTransportToZ2Airlock   bool
	TypedSSplitExteriorMap       bool
	NativeUniformInsertionReason bool
	NativeScalarType             bool
	NativeBasepointReduction     bool
	NativeSelectorCompatibility  bool
	NativeAlphaTheorem           bool
	DenominatorNormalization     bool
	AlphaNative                  bool
	NativeR3                     bool
	FullAFDescent                bool
	GenerationCarrierMap         bool
	FlavorOrientationMap         bool
	IndividualYukawaValues       bool
	NativeYukawaOperator         bool
}

type Audit struct {
	ID             string
	Truth          string
	Classification string
	ShortStatus    string
	Inherited      InheritedSubobjects
	Target         ScalarTransportTarget
	Insertion      UniformInsertion
	Scalar         ScalarCompatibility
	Reduction      ActiveResponseReduction
	Selector       SelectorCompatibility
	Alpha          AlphaUnderTransport
	Firewalls      Firewalls
	Final          string
}

func BuildDefault() (Audit, error) {
	lin := alphaLinear(SBoundary)
	quad := alphaQuadratic(SBoundary)
	total := lin + quad
	if !near(lin, AlphaLinear) {
		return Audit{}, fmt.Errorf("alpha linear mismatch: got %.18g want %.18g", lin, AlphaLinear)
	}
	if !near(quad, AlphaQuad) {
		return Audit{}, fmt.Errorf("alpha quadratic mismatch: got %.18g want %.18g", quad, AlphaQuad)
	}
	if !near(total, AlphaB) {
		return Audit{}, fmt.Errorf("alpha total mismatch: got %.18g want %.18g", total, AlphaB)
	}

	return Audit{
		ID:             AuditID,
		Truth:          FinalTruth,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		Inherited: InheritedSubobjects{
			Gate913Classification:  Gate913Classification,
			Gate913ShortStatus:     Gate913ShortStatus,
			Gate914Classification:  Gate914Classification,
			Gate914ShortStatus:     Gate914ShortStatus,
			Gate915Classification:  Gate915Classification,
			Gate915ShortStatus:     Gate915ShortStatus,
			ResponseShapeCertified: true,
			SelectorShapeTyped:     true,
			CrossLanesBlocked:      true,
			DerivesAlpha:           false,
			UpdatesOfficialLedger:  false,
			PromotesNativeR3:       false,
			Supports:               []string{StatusGates913To915Inherited},
			Failures:               []string{FailureNoNativeTsMap, FailureAlphaStillSealed, FailureNotNativeR3},
		},
		Target: ScalarTransportTarget{
			Source:               TransportSource,
			Target:               TransportTarget,
			MapCandidate:         TransportMapCandidate,
			TargetsAlphaDirectly: false,
			TargetsSocketMag:     false,
			UsesRBOnly:           true,
			NativeMap:            false,
			Supports:             []string{StatusTransportTargetTyped, SupportTransportTargetIsResponseParameter, SupportSSplitScalarResponseParameter, SupportAlphaUsesSSplitOnlyThroughRB},
			Failures:             []string{FailureNoNativeTsMap},
		},
		Insertion: UniformInsertion{
			Factor1:                    BoundaryFactor1,
			Factor2:                    BoundaryFactor2,
			Insertions:                 2,
			ScalarInsertedPerFactor:    true,
			SeparateQuadraticTransport: false,
			QuadraticFromProduct:       true,
			ExpandedResponse:           ExpandedReducedResponse,
			NativeUniformLaw:           false,
			Supports:                   []string{StatusSingleInsertion, SupportSingleInsertionGeneratesPowers, SupportS2FromExteriorProduct, SupportPowerResponseUniformInsertion},
			Failures:                   []string{FailureNoNativeUniformInsertionReason},
		},
		Scalar: ScalarCompatibility{
			Parameter:             SBoundary,
			Dimensionless:         true,
			CanMultiplyGenerators: true,
			ScalarTypeNative:      false,
			Supports:              []string{StatusDimensionlessCompatibility, SupportSSplitScalarType, SupportDimensionlessMultipliesExteriorGenerator},
			Failures:              []string{FailureSSplitScalarTypeSealed, FailureNoTypedSSplitExteriorParameterMap},
		},
		Reduction: ActiveResponseReduction{
			UnreducedResponse:        "E_B(s)=(1+s b1)(1+s b2)",
			IdentityTerm:             IdentityTerm,
			ActiveTerms:              ActiveBoundaryTerms,
			Reduction:                ReducedResponse,
			TransportAppliesToActive: true,
			IdentityTransported:      false,
			BasepointRemoved:         true,
			NativeBasepointTheorem:   false,
			Supports:                 []string{StatusActiveBasepointReduction, SupportTransportAppliesToActiveTerms, SupportIdentityBasepointRemoved},
			Failures:                 []string{FailureNoNativeBasepointReduction},
		},
		Selector: SelectorCompatibility{
			Pipeline:                  SelectorPipeline,
			DegreeOnePower:            "s(b1+b2)",
			DegreeTwoPower:            "s^2(b1 wedge b2)",
			DegreeSelectorCompatible:  true,
			FeedsCorrectAlphaLanes:    true,
			ReopensCrossLanePollution: false,
			NativeTransportTheorem:    false,
			Supports:                  []string{StatusSelectorCompatibility, SupportCompatibleWithDegreeSelector, SupportSingleInsertionFeedsCorrectAlphaLanes, SupportTransportDoesNotReopenCrossLanes},
			Failures:                  []string{FailureSelectorCompatibilityNotNative},
		},
		Alpha: AlphaUnderTransport{
			S:                      SBoundary,
			AlphaLinear:            lin,
			AlphaQuadratic:         quad,
			AlphaTotal:             total,
			Formula:                AlphaFormula,
			RankPair:               [2]int{RankF1OverF0, RankF2OverF0},
			TransportSealAssumed:   true,
			PriorSubobjectsAssumed: true,
			NativeAlphaTheorem:     false,
			Supports:               []string{StatusAlphaReconstruction, SupportAlphaReconstructedGivenTransport, SupportTransportWoundReducedToNativeTs},
			Failures:               []string{FailureAlphaReconstructionNotNative, FailureAlphaStillSealed},
		},
		Firewalls: Firewalls{
			NativeTsMap:                  false,
			NativeTransportToZ2Airlock:   false,
			TypedSSplitExteriorMap:       false,
			NativeUniformInsertionReason: false,
			NativeScalarType:             false,
			NativeBasepointReduction:     false,
			NativeSelectorCompatibility:  false,
			NativeAlphaTheorem:           false,
			DenominatorNormalization:     false,
			AlphaNative:                  false,
			NativeR3:                     false,
			FullAFDescent:                false,
			GenerationCarrierMap:         false,
			FlavorOrientationMap:         false,
			IndividualYukawaValues:       false,
			NativeYukawaOperator:         false,
		},
		Final: NextGate,
	}, nil
}

func alphaLinear(s float64) float64 { return (float64(RankF1OverF0) / float64(LinearDenom)) * s }
func alphaQuadratic(s float64) float64 {
	return (float64(RankF2OverF0) / float64(QuadraticDenom)) * s * s
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
	return !f.NativeTsMap && !f.NativeTransportToZ2Airlock && !f.TypedSSplitExteriorMap && !f.NativeUniformInsertionReason && !f.NativeScalarType && !f.NativeBasepointReduction && !f.NativeSelectorCompatibility && !f.NativeAlphaTheorem && !f.DenominatorNormalization && !f.AlphaNative && !f.NativeR3 && !f.FullAFDescent && !f.GenerationCarrierMap && !f.FlavorOrientationMap && !f.IndividualYukawaValues && !f.NativeYukawaOperator
}

func (a Audit) FirewallsList() []string {
	return []string{
		FailureNoNativeTsMap,
		FailureNoNativeTransportToZ2Airlock,
		FailureNoTypedSSplitExteriorParameterMap,
		FailureNoNativeUniformInsertionReason,
		FailureSSplitScalarTypeSealed,
		FailureNoNativeBasepointReduction,
		FailureSelectorCompatibilityNotNative,
		FailureAlphaReconstructionNotNative,
		FailureDenominatorNormalizationExternal,
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
	return []string{StatusGates913To915Inherited, StatusTransportTargetTyped, StatusSingleInsertion, StatusDimensionlessCompatibility, StatusActiveBasepointReduction, StatusSelectorCompatibility, StatusAlphaReconstruction, StatusNativeTransportMapStillBlock}
}

func Supports() []string {
	return []string{SupportTransportTargetIsResponseParameter, SupportSSplitScalarResponseParameter, SupportAlphaUsesSSplitOnlyThroughRB, SupportSingleInsertionGeneratesPowers, SupportS2FromExteriorProduct, SupportPowerResponseUniformInsertion, SupportSSplitScalarType, SupportDimensionlessMultipliesExteriorGenerator, SupportTransportAppliesToActiveTerms, SupportIdentityBasepointRemoved, SupportCompatibleWithDegreeSelector, SupportSingleInsertionFeedsCorrectAlphaLanes, SupportTransportDoesNotReopenCrossLanes, SupportAlphaReconstructedGivenTransport, SupportTransportWoundReducedToNativeTs}
}

func Failures() []string {
	return []string{FailureNoNativeTsMap, FailureNoNativeTransportToZ2Airlock, FailureNoTypedSSplitExteriorParameterMap, FailureNoNativeUniformInsertionReason, FailureSSplitScalarTypeSealed, FailureNoNativeBasepointReduction, FailureSelectorCompatibilityNotNative, FailureAlphaReconstructionNotNative, FailureDenominatorNormalizationExternal, FailureAlphaStillSealed, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}
}

func FormatInherited(x InheritedSubobjects) string {
	return fmt.Sprintf("inherited={gate913:%s/%s gate914:%s/%s gate915:%s/%s response_shape:%t selector:%t cross_lane:%t derives_alpha:%t updates_official:%t native_r3:%t}", x.Gate913Classification, x.Gate913ShortStatus, x.Gate914Classification, x.Gate914ShortStatus, x.Gate915Classification, x.Gate915ShortStatus, x.ResponseShapeCertified, x.SelectorShapeTyped, x.CrossLanesBlocked, x.DerivesAlpha, x.UpdatesOfficialLedger, x.PromotesNativeR3)
}

func FormatTarget(x ScalarTransportTarget) string {
	return fmt.Sprintf("target={source:%s target:%s map:%s alpha_direct:%t socket_mag:%t rb_only:%t native_map:%t}", x.Source, x.Target, x.MapCandidate, x.TargetsAlphaDirectly, x.TargetsSocketMag, x.UsesRBOnly, x.NativeMap)
}

func FormatInsertion(x UniformInsertion) string {
	return fmt.Sprintf("insertion={factor1:%s factor2:%s insertions:%d per_factor:%t separate_s2_transport:%t quadratic_from_product:%t expanded:%s native_uniform:%t}", x.Factor1, x.Factor2, x.Insertions, x.ScalarInsertedPerFactor, x.SeparateQuadraticTransport, x.QuadraticFromProduct, x.ExpandedResponse, x.NativeUniformLaw)
}

func FormatScalar(x ScalarCompatibility) string {
	return fmt.Sprintf("scalar={s:%.18g dimensionless:%t multiply_generators:%t scalar_type_native:%t}", x.Parameter, x.Dimensionless, x.CanMultiplyGenerators, x.ScalarTypeNative)
}

func FormatReduction(x ActiveResponseReduction) string {
	return fmt.Sprintf("reduction={unreduced:%s identity:%s active:%s reduction:%s transport_active:%t identity_transported:%t basepoint_removed:%t native_basepoint:%t}", x.UnreducedResponse, x.IdentityTerm, x.ActiveTerms, x.Reduction, x.TransportAppliesToActive, x.IdentityTransported, x.BasepointRemoved, x.NativeBasepointTheorem)
}

func FormatSelector(x SelectorCompatibility) string {
	return fmt.Sprintf("selector={pipeline:%s degree1:%s degree2:%s compatible:%t feeds_correct:%t reopens_cross_lanes:%t native_transport:%t}", x.Pipeline, x.DegreeOnePower, x.DegreeTwoPower, x.DegreeSelectorCompatible, x.FeedsCorrectAlphaLanes, x.ReopensCrossLanePollution, x.NativeTransportTheorem)
}

func FormatAlpha(x AlphaUnderTransport) string {
	return fmt.Sprintf("alpha={s:%.18g linear:%.18g quadratic:%.18g total:%.18g formula:%s rank_pair:%d,%d transport_seal:%t prior_subobjects:%t native_alpha:%t}", x.S, x.AlphaLinear, x.AlphaQuadratic, x.AlphaTotal, x.Formula, x.RankPair[0], x.RankPair[1], x.TransportSealAssumed, x.PriorSubobjectsAssumed, x.NativeAlphaTheorem)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls={native_ts:%t native_z2_transport:%t typed_s_map:%t uniform_reason:%t scalar_native:%t basepoint_native:%t selector_native:%t alpha_theorem:%t denominator_norm:%t alpha_native:%t native_r3:%t full_af:%t generation:%t flavor:%t individual_yukawa:%t native_yukawa:%t}", f.NativeTsMap, f.NativeTransportToZ2Airlock, f.TypedSSplitExteriorMap, f.NativeUniformInsertionReason, f.NativeScalarType, f.NativeBasepointReduction, f.NativeSelectorCompatibility, f.NativeAlphaTheorem, f.DenominatorNormalization, f.AlphaNative, f.NativeR3, f.FullAFDescent, f.GenerationCarrierMap, f.FlavorOrientationMap, f.IndividualYukawaValues, f.NativeYukawaOperator)
}

func JoinMarkers(xs []string) string { return strings.Join(xs, ",") }
