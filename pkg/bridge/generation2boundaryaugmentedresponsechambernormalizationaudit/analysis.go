// Package generation2boundaryaugmentedresponsechambernormalizationaudit implements
// Gate 917: BoundaryAugmented ResponseChamber Normalization Audit.
//
// Gate 917 follows Gate 916's S_split transport audit and addresses the fifth
// Gate 912 sub-object: whether the denominators in
//
//	alpha_B=(3/10)s+(7/72)s^2
//
// can be typed as boundary-augmented response chambers rather than convenient
// numerical normalizers. The audit supports the shape-level typing
//
//	H_10 = H_R^ambient ⊕ B_2, rank 8+2=10
//	H_72 = Lambda^4 V_8 ⊕ B_2, rank 70+2=72
//
// and reconstructs alpha_B from all five decomposed sub-objects. It does not
// certify a native normalization/activation theorem, native BoundaryAlpha, or
// native R3.
package generation2boundaryaugmentedresponsechambernormalizationaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE917-BOUNDARY-AUGMENTED-RESPONSE-CHAMBER-NORMALIZATION-AUDIT"

	SBoundary = 0.0012924448188162962
	AlphaB    = 0.0003878958469680527

	RankF1OverF0  = 3
	RankF2OverF0  = 7
	RankB2        = 2
	RankHRAmbient = 8
	RankLambda4V8 = 70
	RankH10       = 10
	RankH72       = 72

	LinearBareDenominator    = 8
	QuadraticBareDenominator = 70
	CommonDenominator        = 72

	AlphaLinear        = 0.00038773344564488885
	AlphaQuad          = 0.0000001624013231638281
	BareLinearAlpha    = 0.0004846668070561111
	BareQuadraticAlpha = 0.00000016704136154050832
	SameH72LinearAlpha = 0.00005385186745067901

	Gate913ShortStatus = "R3_ALPHA_SUBOBJECT_1_REDUCED_B2_RESPONSE_SHAPE_PASS_NATIVE_SELECTION_BLOCKED"
	Gate914ShortStatus = "R3_DEGREE_INDEXED_Z2_AIRLOCK_FLAG_SELECTOR_OBSTRUCTION"
	Gate915ShortStatus = "R3_Z2_BOUNDARYALPHA_CROSS_LANE_EXCLUSION_OBSTRUCTION"
	Gate916ShortStatus = "R3_S_SPLIT_TO_REDUCED_B2_RESPONSE_TRANSPORT_OBSTRUCTION"

	H10Definition        = "H_10=H_R^ambient plus B_2"
	H72Definition        = "H_72=Lambda^4 V_8 plus B_2"
	LinearLaneFormula    = "rank([F_1/F_0]_{Z2})/rank(H_10) * s = (3/10)s"
	QuadraticLaneFormula = "rank([F_2/F_0]_{Z2})/rank(H_72) * s^2 = (7/72)s^2"
	AlphaFormula         = "alpha_B=(3/10)s+(7/72)s^2"
	DecomposedFormula    = "alpha_B=rank([F_1/F_0]_{Z2})/rank(H_10)*s + rank([F_2/F_0]_{Z2})/rank(H_72)*s^2"

	NextGate            = "NEXT_PRESSURE_GATE918_Z2_BOUNDARYALPHA_DECOMPOSEDFUNCTOR_CONSOLIDATION_AND_NATIVE_THEOREM_GAP_AUDIT"
	Classification      = "R3_ALPHA_SUBOBJECT_5_RESPONSE_CHAMBER_NORMALIZATION_SHAPE_PASS_NATIVE_NORMALIZATION_BLOCKED"
	ShortStatus         = "R3_BOUNDARY_AUGMENTED_RESPONSE_CHAMBER_NORMALIZATION_OBSTRUCTION"
	FinalTruth          = "BOUNDARY_AUGMENTED_RESPONSE_CHAMBER_DENOMINATORS_TYPED_BUT_NATIVE_NORMALIZATION_THEOREM_MISSING"
	StrategicConclusion = "Gate 917 completes the five Gate 912 sub-object audits at shape level: response shape, selector shape, cross-lane exclusion shape, S_split insertion shape, and boundary-augmented chamber normalization shape. Reconstruction is bridge-level under seals, not a native alpha theorem."

	StatusInheritedFiveSubobjects     = "PASS_GATES913_TO916_FIRST_FOUR_ALPHA_SUBOBJECTS_INHERITED_AT_SHAPE_LEVEL"
	StatusH10Typed                    = "PASS_H10_TYPED_AS_BOUNDARY_AUGMENTED_RIGHT_RECTANGLE_RESPONSE_CHAMBER"
	StatusH72Typed                    = "PASS_H72_TYPED_AS_BOUNDARY_AUGMENTED_LAMBDA4V8_RESPONSE_CHAMBER"
	StatusLaneCompatibility           = "PASS_DENOMINATOR_PAIR_MATCHES_LOCAL_LINEAR_AND_GLOBAL_QUADRATIC_LANES"
	StatusBoundaryAugmentation        = "PASS_BOTH_RESPONSE_CHAMBERS_AUGMENTED_BY_BOUNDARY_PAIR_B2"
	StatusDenominatorContamination    = "PASS_WRONG_DENOMINATORS_DETECTED_AS_ALPHA_SEAL_MISMATCH"
	StatusFiveSubobjectReconstruction = "PASS_ALPHA_RECONSTRUCTED_FROM_ALL_FIVE_DECOMPOSED_SUBOBJECTS_AT_SHAPE_LEVEL"
	StatusNativeNormalizationBlocked  = "FIREWALL_PRESERVED_NATIVE_RESPONSE_CHAMBER_NORMALIZATION_NOT_CERTIFIED"

	SupportDegreeOneDenomHRAmbientPlusB2         = "CONDITIONAL_SUPPORT_DEGREE_ONE_DENOMINATOR_TYPED_AS_H_R_AMBIENT_PLUS_B2"
	SupportH10Rank8Plus2                         = "CONDITIONAL_SUPPORT_H10_CHAMBER_HAS_RANK_8_PLUS_2_EQUALS_10"
	SupportLinearLaneNormalizedByRightRectangle  = "CONDITIONAL_SUPPORT_LINEAR_ALPHA_LANE_NORMALIZED_BY_BOUNDARY_AUGMENTED_RIGHT_RECTANGLE"
	SupportDegreeTwoDenomLambda4V8PlusB2         = "CONDITIONAL_SUPPORT_DEGREE_TWO_DENOMINATOR_TYPED_AS_LAMBDA4V8_PLUS_B2"
	SupportH72Rank70Plus2                        = "CONDITIONAL_SUPPORT_H72_CHAMBER_HAS_RANK_70_PLUS_2_EQUALS_72"
	SupportQuadraticLaneNormalizedBy72           = "CONDITIONAL_SUPPORT_QUADRATIC_ALPHA_LANE_NORMALIZED_BY_AUGMENTED_72_CHAMBER"
	SupportLinearLaneLocalRightRectangle         = "CONDITIONAL_SUPPORT_LINEAR_LANE_USES_LOCAL_RIGHT_RECTANGLE_RESPONSE_CHAMBER"
	SupportQuadraticLaneGlobal72                 = "CONDITIONAL_SUPPORT_QUADRATIC_LANE_USES_GLOBAL_AUGMENTED_72_CHAMBER"
	SupportDenomPairMatchesLocality              = "CONDITIONAL_SUPPORT_DENOMINATOR_PAIR_10_72_MATCHES_LANE_LOCALITY_LEVELS"
	SupportBothChambersBoundaryAugmented         = "CONDITIONAL_SUPPORT_BOTH_RESPONSE_CHAMBERS_ARE_BOUNDARY_AUGMENTED_BY_B2"
	SupportUniformDenomAugmentation              = "CONDITIONAL_SUPPORT_DENOMINATOR_AUGMENTATION_IS_UNIFORM_ACROSS_ALPHA_LANES"
	SupportActiveAlphaRequiresLaneSpecificDenoms = "CONDITIONAL_SUPPORT_ACTIVE_ALPHA_REQUIRES_LANE_SPECIFIC_BOUNDARY_AUGMENTED_DENOMINATORS"
	SupportBareDenomsMismatch                    = "CONDITIONAL_SUPPORT_BARE_DENOMINATORS_8_AND_70_DO_NOT_MATCH_CURRENT_ALPHA_SEAL"
	SupportAllFiveSubobjectsAudited              = "CONDITIONAL_SUPPORT_ALL_FIVE_ALPHA_SUBOBJECTS_NOW_AUDITED_AT_SHAPE_LEVEL"
	SupportAlphaReconstructedFromComponents      = "CONDITIONAL_SUPPORT_ALPHA_B_RECONSTRUCTED_FROM_DECOMPOSED_Z2_BOUNDARY_ALPHA_COMPONENTS"

	FailureH10NotNativeActivation        = "FAILED_ROUTE_H10_NORMALIZATION_NOT_NATIVE_ACTIVATION_THEOREM"
	FailureH72NotNativeActivation        = "FAILED_ROUTE_H72_NORMALIZATION_NOT_NATIVE_ACTIVATION_THEOREM"
	FailureLocalVsGlobalNotNative        = "FAILED_ROUTE_LOCAL_VS_GLOBAL_CHAMBER_ASSIGNMENT_NOT_NATIVE_FUNCTOR_THEOREM"
	FailureBoundaryAugmentationNotNative = "FAILED_ROUTE_BOUNDARY_AUGMENTATION_NOT_NATIVE_NORMALIZATION_THEOREM"
	FailureNumericalMismatchNotNative    = "FAILED_ROUTE_NUMERICAL_MISMATCH_DETECTS_WRONG_DENOMINATORS_BUT_DOES_NOT_PROVE_NATIVE_NORMALIZATION"
	FailureReconstructionNotNativeAlpha  = "FAILED_ROUTE_RECONSTRUCTION_FROM_FIVE_SHAPE_LEVEL_SUBOBJECTS_NOT_NATIVE_ALPHA_THEOREM"
	FailureAlphaStillSealed              = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNotNativeR3                   = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked     = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap        = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap        = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues      = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator        = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type InheritedSubobjectLedger struct {
	Gate913ShortStatus     string
	Gate914ShortStatus     string
	Gate915ShortStatus     string
	Gate916ShortStatus     string
	ResponseShapeAudited   bool
	SelectorShapeAudited   bool
	CrossLaneShapeAudited  bool
	SSplitTransportAudited bool
	NativeAlpha            bool
	NativeR3               bool
	Supports, Failures     []string
}

type DegreeOneChamber struct {
	TargetRank            int
	AmbientRank           int
	BoundaryRank          int
	ChamberRank           int
	Definition            string
	Formula               string
	NormalizedCoefficient float64
	NativeActivation      bool
	Supports, Failures    []string
}

type DegreeTwoChamber struct {
	TargetRank            int
	Lambda4Rank           int
	BoundaryRank          int
	ChamberRank           int
	Definition            string
	Formula               string
	NormalizedCoefficient float64
	NativeActivation      bool
	Supports, Failures    []string
}

type LaneCompatibility struct {
	LinearLaneLocal      bool
	QuadraticLaneGlobal  bool
	DenominatorPair      [2]int
	MatchesLaneLocality  bool
	NativeFunctorTheorem bool
	Supports, Failures   []string
}

type BoundaryAugmentation struct {
	BareDenominators      [2]int
	AugmentedDenominators [2]int
	BoundaryRank          int
	BothAugmentedByB2     bool
	UniformAugmentation   bool
	NativeNormalization   bool
	Supports, Failures    []string
}

type DenominatorContamination struct {
	CorrectAlpha                float64
	BareLinearAlpha             float64
	BareQuadraticAlpha          float64
	SameH72LinearAlpha          float64
	BareLinearMismatches        bool
	BareQuadraticMismatches     bool
	CommonDenominatorMismatches bool
	NativeProof                 bool
	Supports, Failures          []string
}

type FiveSubobjectReconstruction struct {
	ResponseShape         bool
	DegreeSelector        bool
	CrossLaneExclusion    bool
	SSplitTransport       bool
	ChamberNormalization  bool
	Formula               string
	LinearContribution    float64
	QuadraticContribution float64
	TotalAlpha            float64
	NativeAlphaTheorem    bool
	Supports, Failures    []string
}

type FirewallLedger struct {
	H10Native            bool
	H72Native            bool
	LocalGlobalNative    bool
	BoundaryAugNative    bool
	MismatchNative       bool
	ReconstructionNative bool
	AlphaNative          bool
	R3Native             bool
	FullAFDescent        bool
	GenerationCarrier    bool
	FlavorOrientation    bool
	IndividualYukawa     bool
	NativeYukawa         bool
	Failures             []string
}

type Audit struct {
	ID                   string
	Truth                string
	Classification       string
	ShortStatus          string
	Inherited            InheritedSubobjectLedger
	DegreeOne            DegreeOneChamber
	DegreeTwo            DegreeTwoChamber
	LaneCompatibility    LaneCompatibility
	BoundaryAugmentation BoundaryAugmentation
	Contamination        DenominatorContamination
	Reconstruction       FiveSubobjectReconstruction
	Firewalls            FirewallLedger
	Final                string
}

func BuildDefault() (Audit, error) {
	linear := float64(RankF1OverF0) / float64(RankH10) * SBoundary
	quad := float64(RankF2OverF0) / float64(RankH72) * SBoundary * SBoundary
	if !near(linear, AlphaLinear) || !near(quad, AlphaQuad) || !near(linear+quad, AlphaB) {
		return Audit{}, fmt.Errorf("alpha lane arithmetic mismatch: linear %.18g quad %.18g total %.18g", linear, quad, linear+quad)
	}
	return Audit{
		ID:             AuditID,
		Truth:          FinalTruth,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		Inherited: InheritedSubobjectLedger{
			Gate913ShortStatus:     Gate913ShortStatus,
			Gate914ShortStatus:     Gate914ShortStatus,
			Gate915ShortStatus:     Gate915ShortStatus,
			Gate916ShortStatus:     Gate916ShortStatus,
			ResponseShapeAudited:   true,
			SelectorShapeAudited:   true,
			CrossLaneShapeAudited:  true,
			SSplitTransportAudited: true,
			NativeAlpha:            false,
			NativeR3:               false,
			Supports:               []string{StatusInheritedFiveSubobjects},
			Failures:               []string{FailureAlphaStillSealed, FailureNotNativeR3},
		},
		DegreeOne: DegreeOneChamber{
			TargetRank:            RankF1OverF0,
			AmbientRank:           RankHRAmbient,
			BoundaryRank:          RankB2,
			ChamberRank:           RankH10,
			Definition:            H10Definition,
			Formula:               LinearLaneFormula,
			NormalizedCoefficient: float64(RankF1OverF0) / float64(RankH10),
			NativeActivation:      false,
			Supports:              []string{SupportDegreeOneDenomHRAmbientPlusB2, SupportH10Rank8Plus2, SupportLinearLaneNormalizedByRightRectangle, StatusH10Typed},
			Failures:              []string{FailureH10NotNativeActivation},
		},
		DegreeTwo: DegreeTwoChamber{
			TargetRank:            RankF2OverF0,
			Lambda4Rank:           RankLambda4V8,
			BoundaryRank:          RankB2,
			ChamberRank:           RankH72,
			Definition:            H72Definition,
			Formula:               QuadraticLaneFormula,
			NormalizedCoefficient: float64(RankF2OverF0) / float64(RankH72),
			NativeActivation:      false,
			Supports:              []string{SupportDegreeTwoDenomLambda4V8PlusB2, SupportH72Rank70Plus2, SupportQuadraticLaneNormalizedBy72, StatusH72Typed},
			Failures:              []string{FailureH72NotNativeActivation},
		},
		LaneCompatibility: LaneCompatibility{
			LinearLaneLocal:      true,
			QuadraticLaneGlobal:  true,
			DenominatorPair:      [2]int{RankH10, RankH72},
			MatchesLaneLocality:  true,
			NativeFunctorTheorem: false,
			Supports:             []string{SupportLinearLaneLocalRightRectangle, SupportQuadraticLaneGlobal72, SupportDenomPairMatchesLocality, StatusLaneCompatibility},
			Failures:             []string{FailureLocalVsGlobalNotNative},
		},
		BoundaryAugmentation: BoundaryAugmentation{
			BareDenominators:      [2]int{LinearBareDenominator, QuadraticBareDenominator},
			AugmentedDenominators: [2]int{RankH10, RankH72},
			BoundaryRank:          RankB2,
			BothAugmentedByB2:     true,
			UniformAugmentation:   true,
			NativeNormalization:   false,
			Supports:              []string{SupportBothChambersBoundaryAugmented, SupportUniformDenomAugmentation, StatusBoundaryAugmentation},
			Failures:              []string{FailureBoundaryAugmentationNotNative},
		},
		Contamination: DenominatorContamination{
			CorrectAlpha:                AlphaB,
			BareLinearAlpha:             float64(RankF1OverF0) / float64(LinearBareDenominator) * SBoundary,
			BareQuadraticAlpha:          float64(RankF2OverF0) / float64(QuadraticBareDenominator) * SBoundary * SBoundary,
			SameH72LinearAlpha:          float64(RankF1OverF0) / float64(CommonDenominator) * SBoundary,
			BareLinearMismatches:        true,
			BareQuadraticMismatches:     true,
			CommonDenominatorMismatches: true,
			NativeProof:                 false,
			Supports:                    []string{SupportActiveAlphaRequiresLaneSpecificDenoms, SupportBareDenomsMismatch, StatusDenominatorContamination},
			Failures:                    []string{FailureNumericalMismatchNotNative},
		},
		Reconstruction: FiveSubobjectReconstruction{
			ResponseShape:         true,
			DegreeSelector:        true,
			CrossLaneExclusion:    true,
			SSplitTransport:       true,
			ChamberNormalization:  true,
			Formula:               DecomposedFormula,
			LinearContribution:    linear,
			QuadraticContribution: quad,
			TotalAlpha:            linear + quad,
			NativeAlphaTheorem:    false,
			Supports:              []string{SupportAllFiveSubobjectsAudited, SupportAlphaReconstructedFromComponents, StatusFiveSubobjectReconstruction},
			Failures:              []string{FailureReconstructionNotNativeAlpha, FailureAlphaStillSealed},
		},
		Firewalls: FirewallLedger{
			Failures: Failures(),
		},
		Final: NextGate,
	}, nil
}

func Statuses() []string {
	return []string{StatusInheritedFiveSubobjects, StatusH10Typed, StatusH72Typed, StatusLaneCompatibility, StatusBoundaryAugmentation, StatusDenominatorContamination, StatusFiveSubobjectReconstruction, StatusNativeNormalizationBlocked}
}

func Supports() []string {
	return []string{SupportDegreeOneDenomHRAmbientPlusB2, SupportH10Rank8Plus2, SupportLinearLaneNormalizedByRightRectangle, SupportDegreeTwoDenomLambda4V8PlusB2, SupportH72Rank70Plus2, SupportQuadraticLaneNormalizedBy72, SupportLinearLaneLocalRightRectangle, SupportQuadraticLaneGlobal72, SupportDenomPairMatchesLocality, SupportBothChambersBoundaryAugmented, SupportUniformDenomAugmentation, SupportActiveAlphaRequiresLaneSpecificDenoms, SupportBareDenomsMismatch, SupportAllFiveSubobjectsAudited, SupportAlphaReconstructedFromComponents}
}

func Failures() []string {
	return []string{FailureH10NotNativeActivation, FailureH72NotNativeActivation, FailureLocalVsGlobalNotNative, FailureBoundaryAugmentationNotNative, FailureNumericalMismatchNotNative, FailureReconstructionNotNativeAlpha, FailureAlphaStillSealed, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}
}

func (f FirewallLedger) List() []string { return append([]string{}, f.Failures...) }

func firewallsOK(f FirewallLedger) bool {
	return !f.H10Native && !f.H72Native && !f.LocalGlobalNative && !f.BoundaryAugNative && !f.MismatchNative && !f.ReconstructionNative && !f.AlphaNative && !f.R3Native && !f.FullAFDescent && !f.GenerationCarrier && !f.FlavorOrientation && !f.IndividualYukawa && !f.NativeYukawa && containsAll(f.Failures, Failures())
}

func FormatInherited(x InheritedSubobjectLedger) string {
	return fmt.Sprintf("inherited={913:%s 914:%s 915:%s 916:%s response:%t selector:%t crosslane:%t transport:%t nativeAlpha:%t nativeR3:%t supports:%s failures:%s}", x.Gate913ShortStatus, x.Gate914ShortStatus, x.Gate915ShortStatus, x.Gate916ShortStatus, x.ResponseShapeAudited, x.SelectorShapeAudited, x.CrossLaneShapeAudited, x.SSplitTransportAudited, x.NativeAlpha, x.NativeR3, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}
func FormatDegreeOne(x DegreeOneChamber) string {
	return fmt.Sprintf("degreeOne={targetRank:%d ambient:%d boundary:%d chamber:%d coeff:%g definition:%s native:%t supports:%s failures:%s}", x.TargetRank, x.AmbientRank, x.BoundaryRank, x.ChamberRank, x.NormalizedCoefficient, x.Definition, x.NativeActivation, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}
func FormatDegreeTwo(x DegreeTwoChamber) string {
	return fmt.Sprintf("degreeTwo={targetRank:%d lambda4:%d boundary:%d chamber:%d coeff:%g definition:%s native:%t supports:%s failures:%s}", x.TargetRank, x.Lambda4Rank, x.BoundaryRank, x.ChamberRank, x.NormalizedCoefficient, x.Definition, x.NativeActivation, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}
func FormatLaneCompatibility(x LaneCompatibility) string {
	return fmt.Sprintf("laneCompatibility={linearLocal:%t quadraticGlobal:%t denominators:%v matchesLocality:%t nativeFunctor:%t supports:%s failures:%s}", x.LinearLaneLocal, x.QuadraticLaneGlobal, x.DenominatorPair, x.MatchesLaneLocality, x.NativeFunctorTheorem, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}
func FormatBoundaryAugmentation(x BoundaryAugmentation) string {
	return fmt.Sprintf("boundaryAugmentation={bare:%v augmented:%v boundaryRank:%d bothB2:%t uniform:%t native:%t supports:%s failures:%s}", x.BareDenominators, x.AugmentedDenominators, x.BoundaryRank, x.BothAugmentedByB2, x.UniformAugmentation, x.NativeNormalization, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}
func FormatContamination(x DenominatorContamination) string {
	return fmt.Sprintf("contamination={correct:%g bareLinear:%g bareQuadratic:%g sameH72Linear:%g bareLinearMismatch:%t bareQuadraticMismatch:%t commonDenomMismatch:%t nativeProof:%t supports:%s failures:%s}", x.CorrectAlpha, x.BareLinearAlpha, x.BareQuadraticAlpha, x.SameH72LinearAlpha, x.BareLinearMismatches, x.BareQuadraticMismatches, x.CommonDenominatorMismatches, x.NativeProof, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}
func FormatReconstruction(x FiveSubobjectReconstruction) string {
	return fmt.Sprintf("reconstruction={response:%t selector:%t crosslane:%t transport:%t normalization:%t linear:%g quadratic:%g total:%g nativeAlpha:%t formula:%s supports:%s failures:%s}", x.ResponseShape, x.DegreeSelector, x.CrossLaneExclusion, x.SSplitTransport, x.ChamberNormalization, x.LinearContribution, x.QuadraticContribution, x.TotalAlpha, x.NativeAlphaTheorem, x.Formula, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}
func FormatFirewalls(f FirewallLedger) string {
	return fmt.Sprintf("firewalls={h10Native:%t h72Native:%t localGlobalNative:%t boundaryAugNative:%t reconstructionNative:%t alphaNative:%t r3Native:%t fullAF:%t generation:%t flavor:%t individualYukawa:%t nativeYukawa:%t failures:%s}", f.H10Native, f.H72Native, f.LocalGlobalNative, f.BoundaryAugNative, f.ReconstructionNative, f.AlphaNative, f.R3Native, f.FullAFDescent, f.GenerationCarrier, f.FlavorOrientation, f.IndividualYukawa, f.NativeYukawa, strings.Join(f.Failures, ","))
}

func containsAll(haystack, needles []string) bool {
	for _, n := range needles {
		found := false
		for _, h := range haystack {
			if h == n {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
func near(a, b float64) bool { return math.Abs(a-b) <= 1e-15 }
