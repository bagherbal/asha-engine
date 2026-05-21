// Package generation2boundaryactivationmeasurenativeconstraintsourceaudit implements
// Gate 922: BoundaryActivationMeasure NativeConstraint Source Audit.
//
// Gate 922 follows Gate 921's uniqueness-under-constraints result and asks
// whether the constraints that made mu_B unique are themselves native, bridge-
// lawful, dependent, or merely compatibility conditions. The result is a mixed
// source ledger: exterior-degree respect has the strongest native-shape source;
// basepoint deviation, Z2 orientation-class invariance, and local/global chamber
// typing are strong bridge sources; selector functionhood remains the primary
// native gap for the BoundaryActivationMeasure.
package generation2boundaryactivationmeasurenativeconstraintsourceaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE922-BOUNDARYACTIVATIONMEASURE-NATIVECONSTRAINT-SOURCE-AUDIT"

	Gate921ShortStatus = "R3_ALPHA_MEASURE_UNIQUENESS_SUPPORTED_NATIVE_THEOREM_MISSING"

	SBoundary = 0.0012924448188162962
	AlphaB    = 0.0003878958469680527

	RankI1  = 3
	RankI2  = 7
	RankH10 = 10
	RankH72 = 72

	AlphaLinear = 0.00038773344564488885
	AlphaQuad   = 0.0000001624013231638281

	ReducedResponse       = "R_B(s)=E_B(s)-1=(1+s b1)(1+s b2)-1"
	ExteriorResponse      = "R_B(s)=s(b1+b2)+s^2(b1 wedge b2)"
	MeasureFormula        = "mu_B(R_B(S_split))=sum_{k=1}^2 rank(I_B^Z2(k))/rank(H_k)*S_split^k"
	BranchMeasureFormula  = "mu_B(R_B(S_split))=(3/10)S_split+(7/72)S_split^2"
	BoundaryAlphaFormula  = "alpha_B^Z2=mu_B(R_B(S_split))"
	BoundaryMeasureObject = "BoundaryActivationMeasure"

	SourceBridgeStrong             = "BRIDGE_STRONG_NOT_NATIVE"
	SourceNativeShapeStrong        = "NATIVE_SHAPE_STRONG"
	SourceBridgeCandidateNotNative = "BRIDGE_CANDIDATE_NOT_NATIVE"
	SourceDependentOnSelector      = "DEPENDENT_ON_SELECTOR"
	SourceBridgeStrongOrientation  = "BRIDGE_STRONG_ORIENTATION_CLASS"
	SourceCompatibilityOnly        = "COMPATIBILITY_ONLY"
	PrimaryGapSelectorFunctionhood = "PRIMARY_GAP_SELECTOR_FUNCTIONHOOD"
	StrongestNativeShapeSource     = "STRONGEST_NATIVE_SHAPE_SOURCE_DEGREE_RESPECT_FROM_EXTERIOR_ALGEBRA"
	StrongestBridgeSources         = "STRONGEST_BRIDGE_SOURCES_BASEPOINT_Z2_ORIENTATION_AND_LOCAL_GLOBAL_CHAMBERS"

	Classification      = "R3_BOUNDARY_MEASURE_CONSTRAINT_SOURCE_AUDIT_SELECTOR_FUNCTOR_PRIMARY_GAP"
	ShortStatus         = "R3_ALPHA_MEASURE_CONSTRAINTS_PARTLY_SOURCED_SELECTOR_STILL_MISSING"
	FinalTruth          = "BOUNDARY_ACTIVATION_MEASURE_CONSTRAINT_SOURCES_AUDITED_SELECTOR_FUNCTIONHOOD_REMAINS_PRIMARY_NATIVE_GAP"
	StrategicConclusion = "Gate 922 refines the native wound: most BoundaryActivationMeasure constraints have bridge or native-shape support, but selector functionhood remains the primary native gap. The next pressure point is why exterior degree k must select [F_k/F_0]_{Z2}."
	NextGate            = "NEXT_PRESSURE_GATE923_DEGREEINDEXED_SELECTOR_FUNCTIONHOOD_SOURCE_AUDIT"

	StatusInheritedGate921          = "PASS_GATE921_UNIQUENESS_UNDER_CONSTRAINTS_INHERITED"
	StatusReducedSourceAudited      = "PASS_REDUCED_RESPONSE_CONSTRAINT_HAS_BASEPOINT_DEVIATION_SOURCE"
	StatusDegreeSourceAudited       = "PASS_DEGREE_RESPECT_HAS_NATIVE_EXTERIOR_ALGEBRA_SOURCE"
	StatusSelectorSourceAudited     = "PASS_SELECTOR_FUNCTIONHOOD_HAS_EXPOSURE_ENCLOSURE_SOURCE_CANDIDATE"
	StatusCrossLaneDependent        = "PASS_CROSS_LANE_EXCLUSION_DEPENDS_ON_SELECTOR_FUNCTIONHOOD"
	StatusChamberSourceAudited      = "PASS_CHAMBER_NORMALIZATION_HAS_LOCAL_GLOBAL_LANE_SOURCE_CANDIDATE"
	StatusZ2SourceAudited           = "PASS_Z2_REPRESENTATIVE_INDEPENDENCE_HAS_ORIENTATION_CLASS_SOURCE"
	StatusPositivitySourceAudited   = "PASS_POSITIVITY_IS_COMPATIBILITY_NOT_SELECTION_THEOREM"
	StatusPrimaryGapIdentified      = "PASS_SELECTOR_FUNCTIONHOOD_IDENTIFIED_AS_PRIMARY_NATIVE_GAP_FOR_MU_B"
	StatusNativeMeasureStillMissing = "FIREWALL_PRESERVED_NATIVE_BOUNDARY_ACTIVATION_MEASURE_MISSING"
	StatusNoNativeR3Promotion       = "FIREWALL_PRESERVED_ALPHA_AND_R3_NOT_NATIVE"

	SupportReducedBasepointDeviation     = "CONDITIONAL_SUPPORT_REDUCED_RESPONSE_HAS_BASEPOINT_DEVIATION_SOURCE"
	SupportLambda0InactiveBasepoint      = "CONDITIONAL_SUPPORT_LAMBDA0_TERM_IS_INACTIVE_BOUNDARY_BASEPOINT"
	SupportAlphaStartsAfterActivation    = "CONDITIONAL_SUPPORT_ALPHA_RESPONSE_STARTS_ONLY_AFTER_BOUNDARY_ACTIVATION"
	SupportDegreeNativeExteriorSource    = "CONDITIONAL_SUPPORT_DEGREE_RESPECT_HAS_NATIVE_EXTERIOR_ALGEBRA_SOURCE"
	SupportSPowerFromExteriorDegree      = "CONDITIONAL_SUPPORT_S_POWER_K_FOLLOWS_FROM_EXTERIOR_MULTIPLICATIVE_DEGREE"
	SupportDegreeTwoNotSeparateTransport = "CONDITIONAL_SUPPORT_DEGREE_TWO_POWER_IS_NOT_SEPARATELY_INSERTED"
	SupportSelectorExposureSource        = "CONDITIONAL_SUPPORT_SELECTOR_FUNCTIONHOOD_HAS_EXPOSURE_ENCLOSURE_SOURCE_CANDIDATE"
	SupportDegreeOneTwoSourceTargets     = "CONDITIONAL_SUPPORT_DEGREE_ONE_AS_EXPOSURE_AND_DEGREE_TWO_AS_ENCLOSURE_SOURCE_TYPES_TARGETS"
	SupportCrossLaneSourceSelector       = "CONDITIONAL_SUPPORT_CROSS_LANE_EXCLUSION_SOURCE_IS_SELECTOR_FUNCTIONHOOD"
	SupportCrossLaneDependent            = "CONDITIONAL_SUPPORT_CROSS_LANE_EXCLUSION_IS_DEPENDENT_NOT_PRIMARY"
	SupportChamberLocalGlobalSource      = "CONDITIONAL_SUPPORT_CHAMBER_NORMALIZATION_HAS_LOCAL_GLOBAL_LANE_SOURCE_CANDIDATE"
	SupportH10LocalRightChamber          = "CONDITIONAL_SUPPORT_H10_IS_LOCAL_RIGHT_RECTANGLE_BOUNDARY_CHAMBER"
	SupportH72GlobalLambda4Chamber       = "CONDITIONAL_SUPPORT_H72_IS_GLOBAL_LAMBDA4_BOUNDARY_CHAMBER"
	SupportUniformBoundaryAugmentation   = "CONDITIONAL_SUPPORT_BOUNDARY_AUGMENTATION_IS_UNIFORM_IN_BOTH_LANES"
	SupportZ2StrongOrientationSource     = "CONDITIONAL_SUPPORT_Z2_REPRESENTATIVE_INDEPENDENCE_HAS_STRONG_ORIENTATION_CLASS_SOURCE"
	SupportPhaseGaugeForAlphaTrace       = "CONDITIONAL_SUPPORT_PHASE_SIGN_IS_GAUGE_FOR_ALPHA_AND_TRACE_LEDGER"
	SupportAlphaRankPairZ2Invariant      = "CONDITIONAL_SUPPORT_ALPHA_RANK_PAIR_3_7_IS_Z2_CLASS_INVARIANT"
	SupportPositiveActiveLanes           = "CONDITIONAL_SUPPORT_ALPHA_LANES_ARE_POSITIVE_FOR_POSITIVE_S_SPLIT"
	SupportPositiveMeasureActiveResponse = "CONDITIONAL_SUPPORT_BOUNDARY_ACTIVATION_MEASURE_IS_POSITIVE_ON_ACTIVE_RESPONSE"
	SupportSelectorPrimaryGap            = "CONDITIONAL_SUPPORT_SELECTOR_FUNCTIONHOOD_IS_PRIMARY_REMAINING_NATIVE_GAP_FOR_MU_B"

	FailureNoNativeBoundaryActivationMeasure     = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_ACTIVATION_MEASURE_CERTIFIED"
	FailureNoNativeSelectorFunctionhood          = "FAILED_ROUTE_NO_NATIVE_SELECTOR_FUNCTIONHOOD_THEOREM"
	FailureNoNativeDegreeToZ2FlagFunctor         = "FAILED_ROUTE_NO_NATIVE_DEGREE_TO_Z2_FLAG_CLASS_FUNCTOR"
	FailureExposureEnclosureNotNativeFunctor     = "FAILED_ROUTE_EXPOSURE_ENCLOSURE_TYPING_NOT_NATIVE_FUNCTOR"
	FailureNoNativeCrossLaneWithoutSelector      = "FAILED_ROUTE_NO_NATIVE_CROSS_LANE_EXCLUSION_WITHOUT_NATIVE_SELECTOR_FUNCTIONHOOD"
	FailureNoNativeLaneLocalityToChamber         = "FAILED_ROUTE_NO_NATIVE_LANE_LOCALITY_TO_CHAMBER_THEOREM"
	FailureNoNativeResponseChamberNormalization  = "FAILED_ROUTE_NO_NATIVE_RESPONSE_CHAMBER_NORMALIZATION_THEOREM"
	FailureNoNativeGlobalPhaseZ2Equivariance     = "FAILED_ROUTE_NO_NATIVE_GLOBAL_PHASE_Z2_EQUIVARIANCE_THEOREM"
	FailureNoNativeBasepointDeviation            = "FAILED_ROUTE_NO_NATIVE_BASEPOINT_DEVIATION_THEOREM_CERTIFIED"
	FailurePositivityNotSelectionTheorem         = "FAILED_ROUTE_POSITIVITY_NOT_NATIVE_SELECTION_THEOREM"
	FailurePositivityNotUniqueMuB                = "FAILED_ROUTE_POSITIVITY_DOES_NOT_UNIQUELY_DEFINE_MU_B"
	FailureNoFullNativeMeasureFromExteriorDegree = "FAILED_ROUTE_NO_FULL_NATIVE_MEASURE_THEOREM_FROM_EXTERIOR_DEGREE_ALONE"
	FailureAlphaBridgeCandidateNotNative         = "FAILED_ROUTE_ALPHA_B_REMAINS_BRIDGE_CANDIDATE_NOT_NATIVE"
	FailureNotNativeR3                           = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked             = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap                = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap                = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues              = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator                = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type ConstraintSourceAudit struct {
	Name         string
	Constraint   string
	Candidate    string
	SourceStatus string
	Verdict      string
	Primary      bool
	Native       bool
	BridgeLawful bool
	Dependent    bool
	Compatible   bool
	Supports     []string
	Failures     []string
}

type ConstraintSourceLedger struct {
	ReducedResponse      ConstraintSourceAudit
	DegreeRespect        ConstraintSourceAudit
	SelectorFunctionhood ConstraintSourceAudit
	CrossLaneExclusion   ConstraintSourceAudit
	ChamberNormalization ConstraintSourceAudit
	Z2Independence       ConstraintSourceAudit
	Positivity           ConstraintSourceAudit
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
	InheritedStatus                   string
	ConstraintsPartlySourced          bool
	DegreeRespectStrongestNativeShape bool
	SelectorFunctionhoodPrimaryGap    bool
	NativeBoundaryActivationMeasure   bool
	NativeAlpha                       bool
	NativeR3                          bool
	Supports                          []string
	Failures                          []string
}

type FirewallLedger struct {
	NoNativeBoundaryActivationMeasure    bool
	NoNativeSelectorFunctionhood         bool
	NoNativeDegreeToZ2FlagFunctor        bool
	ExposureEnclosureNotNativeFunctor    bool
	NoNativeCrossLaneWithoutSelector     bool
	NoNativeLaneLocalityToChamber        bool
	NoNativeResponseChamberNormalization bool
	NoNativeGlobalPhaseZ2Equivariance    bool
	NoNativeBasepointDeviation           bool
	PositivityNotSelectionTheorem        bool
	AlphaBridgeCandidateNotNative        bool
	NotNativeR3                          bool
	FullAFDescentStillBlocked            bool
	NoGenerationCarrierMap               bool
	NoFlavorOrientationMap               bool
	NoIndividualYukawaValues             bool
	NoNativeYukawaOperator               bool
}

func (f FirewallLedger) List() []string {
	out := []string{}
	if f.NoNativeBoundaryActivationMeasure {
		out = append(out, FailureNoNativeBoundaryActivationMeasure)
	}
	if f.NoNativeSelectorFunctionhood {
		out = append(out, FailureNoNativeSelectorFunctionhood)
	}
	if f.NoNativeDegreeToZ2FlagFunctor {
		out = append(out, FailureNoNativeDegreeToZ2FlagFunctor)
	}
	if f.ExposureEnclosureNotNativeFunctor {
		out = append(out, FailureExposureEnclosureNotNativeFunctor)
	}
	if f.NoNativeCrossLaneWithoutSelector {
		out = append(out, FailureNoNativeCrossLaneWithoutSelector)
	}
	if f.NoNativeLaneLocalityToChamber {
		out = append(out, FailureNoNativeLaneLocalityToChamber)
	}
	if f.NoNativeResponseChamberNormalization {
		out = append(out, FailureNoNativeResponseChamberNormalization)
	}
	if f.NoNativeGlobalPhaseZ2Equivariance {
		out = append(out, FailureNoNativeGlobalPhaseZ2Equivariance)
	}
	if f.NoNativeBasepointDeviation {
		out = append(out, FailureNoNativeBasepointDeviation)
	}
	if f.PositivityNotSelectionTheorem {
		out = append(out, FailurePositivityNotSelectionTheorem)
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
	ID             string
	Classification string
	ShortStatus    string
	Truth          string
	Ledger         ConstraintSourceLedger
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
		return Analysis{}, fmt.Errorf("alpha reconstruction drift: linear=%.18g quad=%.18g alpha=%.18g", linear, quad, alpha)
	}

	ledger := ConstraintSourceLedger{
		ReducedResponse: ConstraintSourceAudit{
			Name:         "reduced response",
			Constraint:   "mu_B acts on R_B(s)=E_B(s)-1, not E_B(s)",
			Candidate:    "activation measures deviation from the Lambda^0 basepoint",
			SourceStatus: SourceBridgeStrong,
			Verdict:      "basepoint deviation candidate",
			BridgeLawful: true,
			Supports:     []string{SupportReducedBasepointDeviation, SupportLambda0InactiveBasepoint, SupportAlphaStartsAfterActivation},
			Failures:     []string{FailureNoNativeBasepointDeviation},
		},
		DegreeRespect: ConstraintSourceAudit{
			Name:         "degree respect",
			Constraint:   "degree k component contributes S_split^k",
			Candidate:    "exterior multiplicative degree of R_B(s)",
			SourceStatus: SourceNativeShapeStrong,
			Verdict:      "exterior algebra supplies powers",
			Native:       true,
			BridgeLawful: true,
			Supports:     []string{SupportDegreeNativeExteriorSource, SupportSPowerFromExteriorDegree, SupportDegreeTwoNotSeparateTransport},
			Failures:     []string{FailureNoFullNativeMeasureFromExteriorDegree},
		},
		SelectorFunctionhood: ConstraintSourceAudit{
			Name:         "selector functionhood",
			Constraint:   "I_B^Z2 is a function of exterior degree",
			Candidate:    "degree one exposure / degree two enclosure incidence typing",
			SourceStatus: SourceBridgeCandidateNotNative,
			Verdict:      "exposure/enclosure typing only",
			Primary:      true,
			BridgeLawful: true,
			Supports:     []string{SupportSelectorExposureSource, SupportDegreeOneTwoSourceTargets},
			Failures:     []string{FailureNoNativeSelectorFunctionhood, FailureNoNativeDegreeToZ2FlagFunctor, FailureExposureEnclosureNotNativeFunctor},
		},
		CrossLaneExclusion: ConstraintSourceAudit{
			Name:         "cross-lane exclusion",
			Constraint:   "degree one cannot target enclosure and degree two cannot target exposure",
			Candidate:    "downstream consequence of selector functionhood",
			SourceStatus: SourceDependentOnSelector,
			Verdict:      "follows if selector is native",
			Dependent:    true,
			Supports:     []string{SupportCrossLaneSourceSelector, SupportCrossLaneDependent},
			Failures:     []string{FailureNoNativeCrossLaneWithoutSelector},
		},
		ChamberNormalization: ConstraintSourceAudit{
			Name:         "chamber normalization",
			Constraint:   "degree one normalizes by H_10 and degree two by H_72",
			Candidate:    "local right-rectangle exposure vs global Lambda4 enclosure chamber typing",
			SourceStatus: SourceBridgeStrong,
			Verdict:      "local/global chamber typing",
			BridgeLawful: true,
			Supports:     []string{SupportChamberLocalGlobalSource, SupportH10LocalRightChamber, SupportH72GlobalLambda4Chamber, SupportUniformBoundaryAugmentation},
			Failures:     []string{FailureNoNativeLaneLocalityToChamber, FailureNoNativeResponseChamberNormalization},
		},
		Z2Independence: ConstraintSourceAudit{
			Name:         "Z2 independence",
			Constraint:   "mu_B is invariant under lambda <-> bar(lambda)",
			Candidate:    "orientation-class invariant rank pair (3,7)",
			SourceStatus: SourceBridgeStrongOrientation,
			Verdict:      "orientation-class invariant",
			BridgeLawful: true,
			Supports:     []string{SupportZ2StrongOrientationSource, SupportPhaseGaugeForAlphaTrace, SupportAlphaRankPairZ2Invariant},
			Failures:     []string{FailureNoNativeGlobalPhaseZ2Equivariance},
		},
		Positivity: ConstraintSourceAudit{
			Name:         "positivity",
			Constraint:   "response contributions are nonnegative for positive S_split",
			Candidate:    "small positive boundary activation compatibility",
			SourceStatus: SourceCompatibilityOnly,
			Verdict:      "not selective alone",
			Compatible:   true,
			Supports:     []string{SupportPositiveActiveLanes, SupportPositiveMeasureActiveResponse},
			Failures:     []string{FailurePositivityNotSelectionTheorem, FailurePositivityNotUniqueMuB},
		},
	}

	firewalls := FirewallLedger{true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true, true}
	analysis := Analysis{
		ID:             AuditID,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		Truth:          FinalTruth,
		Ledger:         ledger,
		Alpha:          AlphaMeasureAudit{Formula: MeasureFormula, BoundaryFormula: BoundaryAlphaFormula, S: SBoundary, LinearContribution: linear, QuadraticContribution: quad, Alpha: alpha, NativeAlpha: false},
		NativeStatus: NativeStatusAudit{
			InheritedStatus:                   Gate921ShortStatus,
			ConstraintsPartlySourced:          true,
			DegreeRespectStrongestNativeShape: true,
			SelectorFunctionhoodPrimaryGap:    true,
			NativeBoundaryActivationMeasure:   false,
			NativeAlpha:                       false,
			NativeR3:                          false,
			Supports:                          []string{SupportSelectorPrimaryGap, SupportDegreeNativeExteriorSource, SupportZ2StrongOrientationSource, SupportChamberLocalGlobalSource, SupportReducedBasepointDeviation},
			Failures:                          []string{FailureNoNativeBoundaryActivationMeasure, FailureNoNativeSelectorFunctionhood, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3},
		},
		Firewalls: firewalls,
		Final:     StrategicConclusion,
	}
	return analysis, nil
}

func Statuses() []string {
	return []string{StatusInheritedGate921, StatusReducedSourceAudited, StatusDegreeSourceAudited, StatusSelectorSourceAudited, StatusCrossLaneDependent, StatusChamberSourceAudited, StatusZ2SourceAudited, StatusPositivitySourceAudited, StatusPrimaryGapIdentified, StatusNativeMeasureStillMissing, StatusNoNativeR3Promotion}
}
func Supports() []string {
	return []string{SupportReducedBasepointDeviation, SupportLambda0InactiveBasepoint, SupportAlphaStartsAfterActivation, SupportDegreeNativeExteriorSource, SupportSPowerFromExteriorDegree, SupportDegreeTwoNotSeparateTransport, SupportSelectorExposureSource, SupportDegreeOneTwoSourceTargets, SupportCrossLaneSourceSelector, SupportCrossLaneDependent, SupportChamberLocalGlobalSource, SupportH10LocalRightChamber, SupportH72GlobalLambda4Chamber, SupportUniformBoundaryAugmentation, SupportZ2StrongOrientationSource, SupportPhaseGaugeForAlphaTrace, SupportAlphaRankPairZ2Invariant, SupportPositiveActiveLanes, SupportPositiveMeasureActiveResponse, SupportSelectorPrimaryGap}
}
func Failures() []string {
	return []string{FailureNoNativeBoundaryActivationMeasure, FailureNoNativeSelectorFunctionhood, FailureNoNativeDegreeToZ2FlagFunctor, FailureExposureEnclosureNotNativeFunctor, FailureNoNativeCrossLaneWithoutSelector, FailureNoNativeLaneLocalityToChamber, FailureNoNativeResponseChamberNormalization, FailureNoNativeGlobalPhaseZ2Equivariance, FailureNoNativeBasepointDeviation, FailurePositivityNotSelectionTheorem, FailurePositivityNotUniqueMuB, FailureNoFullNativeMeasureFromExteriorDegree, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}
}

func FormatConstraint(c ConstraintSourceAudit) string {
	flags := []string{}
	if c.Primary {
		flags = append(flags, "primary")
	}
	if c.Native {
		flags = append(flags, "native-shape")
	}
	if c.BridgeLawful {
		flags = append(flags, "bridge")
	}
	if c.Dependent {
		flags = append(flags, "dependent")
	}
	if c.Compatible {
		flags = append(flags, "compatibility")
	}
	return fmt.Sprintf("%s|status=%s|verdict=%s|constraint=%s|candidate=%s|flags=%s", c.Name, c.SourceStatus, c.Verdict, c.Constraint, c.Candidate, strings.Join(flags, ","))
}

func FormatLedger(l ConstraintSourceLedger) string {
	return strings.Join([]string{FormatConstraint(l.ReducedResponse), FormatConstraint(l.DegreeRespect), FormatConstraint(l.SelectorFunctionhood), FormatConstraint(l.CrossLaneExclusion), FormatConstraint(l.ChamberNormalization), FormatConstraint(l.Z2Independence), FormatConstraint(l.Positivity)}, " || ")
}

func FormatAlpha(a AlphaMeasureAudit) string {
	return fmt.Sprintf("formula=%s|alpha=%.18g|linear=%.18g|quadratic=%.18g|native=%v", a.Formula, a.Alpha, a.LinearContribution, a.QuadraticContribution, a.NativeAlpha)
}

func FormatNativeStatus(n NativeStatusAudit) string {
	return fmt.Sprintf("inherited=%s|partly_sourced=%v|degree_native_shape=%v|selector_primary_gap=%v|native_measure=%v|native_alpha=%v|native_r3=%v", n.InheritedStatus, n.ConstraintsPartlySourced, n.DegreeRespectStrongestNativeShape, n.SelectorFunctionhoodPrimaryGap, n.NativeBoundaryActivationMeasure, n.NativeAlpha, n.NativeR3)
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
	return f.NoNativeBoundaryActivationMeasure && f.NoNativeSelectorFunctionhood && f.NoNativeDegreeToZ2FlagFunctor && f.ExposureEnclosureNotNativeFunctor && f.NoNativeCrossLaneWithoutSelector && f.NoNativeLaneLocalityToChamber && f.NoNativeResponseChamberNormalization && f.NoNativeGlobalPhaseZ2Equivariance && f.NoNativeBasepointDeviation && f.PositivityNotSelectionTheorem && f.AlphaBridgeCandidateNotNative && f.NotNativeR3 && f.FullAFDescentStillBlocked && f.NoGenerationCarrierMap && f.NoFlavorOrientationMap && f.NoIndividualYukawaValues && f.NoNativeYukawaOperator
}
