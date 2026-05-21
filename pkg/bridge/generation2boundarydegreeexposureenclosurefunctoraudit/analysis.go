// Package generation2boundarydegreeexposureenclosurefunctoraudit implements
// Gate 924: BoundaryDegree ExposureEnclosure Functor Audit.
//
// Gate 924 follows Gate 923's selector-source audit and asks whether the
// exposure/enclosure language is native to the exterior calculus of the rank-two
// boundary pair B_2. It certifies native exterior-degree shape for Lambda^1 B_2
// as one-factor exposure and Lambda^2 B_2 as top-degree full-pair enclosure,
// while preserving the firewall that this still does not give a native target
// functor into the Z2 airlock flag quotients.
package generation2boundarydegreeexposureenclosurefunctoraudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE924-BOUNDARYDEGREE-EXPOSUREENCLOSURE-FUNCTOR-AUDIT"

	Gate923ShortStatus = "R3_ALPHA_SELECTOR_GAP_WEAKENED_TO_EXPOSURE_ENCLOSURE_FUNCTOR"

	BoundaryPair     = "B_2=<b1,b2>"
	ExteriorLedger   = "Lambda^0 B_2=basepoint; Lambda^1 B_2=span{b1,b2}; Lambda^2 B_2=span{b1 wedge b2}; Lambda^3 B_2=0"
	ReducedResponse  = "R_B(s)=(1+s b1)(1+s b2)-1=s(b1+b2)+s^2(b1 wedge b2)"
	DegreeOneTerm    = "s(b1+b2)"
	DegreeTwoTerm    = "s^2(b1 wedge b2)"
	TargetFunctorGap = "TARGET_FUNCTOR_GAP_EXPOSURE_TO_F1_OVER_F0_AND_ENCLOSURE_TO_F2_OVER_F0"
	NextGate         = "NEXT_PRESSURE_GATE925_BOUNDARYDEGREE_TO_AIRLOCKFLAG_TARGETFUNCTOR_AUDIT"

	RankF1OverF0 = 3
	RankF2OverF0 = 7
	RankF2OverF1 = 4

	Classification = "R3_EXPOSURE_ENCLOSURE_FUNCTOR_NATIVE_SHAPE_SUPPORTED_TARGET_MAP_MISSING"
	ShortStatus    = "R3_ALPHA_EXPOSURE_ENCLOSURE_NATIVE_SHAPE_TARGET_FUNCTOR_BLOCKED"
	FinalTruth     = "BOUNDARY_DEGREE_EXPOSURE_ENCLOSURE_TYPING_HAS_NATIVE_EXTERIOR_SHAPE_BUT_TARGET_FUNCTOR_REMAINS_UNCERTIFIED"

	StatusInheritedGate923         = "PASS_GATE923_EXPOSURE_ENCLOSURE_SELECTOR_SOURCE_INHERITED"
	StatusDegreeOneNativeShape     = "PASS_LAMBDA1B2_SINGLE_BOUNDARY_EXPOSURE_NATIVE_EXTERIOR_SHAPE"
	StatusDegreeTwoNativeShape     = "PASS_LAMBDA2B2_FULL_BOUNDARY_PAIR_ENCLOSURE_NATIVE_EXTERIOR_SHAPE"
	StatusDegreeContrastGrounded   = "PASS_EXPOSURE_ENCLOSURE_CONTRAST_GROUNDED_IN_EXTERIOR_DEGREE"
	StatusCumulativePressure       = "PASS_TOP_DEGREE_PAIR_ACTIVATION_SOURCE_TYPES_CUMULATIVE_ENCLOSURE"
	StatusSelectorStrengthened     = "PASS_SELECTOR_SOURCE_REDUCED_TO_EXTERIOR_DEGREE_TYPE"
	StatusMuBSelectorInputStronger = "PASS_BOUNDARY_ACTIVATION_MEASURE_SELECTOR_INPUT_HAS_EXTERIOR_DEGREE_SOURCE"
	StatusTargetFunctorMissing     = "FIREWALL_PRESERVED_TARGET_FUNCTOR_REMAINS_UNCERTIFIED"
	StatusNativeR3Blocked          = "FIREWALL_PRESERVED_ALPHA_AND_R3_NOT_NATIVE"

	SourceNativeExteriorShape = "NATIVE_EXTERIOR_SHAPE_SUPPORTED"
	SourceBridgeTargetBlocked = "TARGET_FUNCTOR_NOT_NATIVE"

	SupportLambda1SingleGenerator               = "CONDITIONAL_SUPPORT_LAMBDA1B2_IS_SINGLE_BOUNDARY_GENERATOR_SPACE"
	SupportDegreeOneActivatesOneFactor          = "CONDITIONAL_SUPPORT_DEGREE_ONE_RESPONSE_ACTIVATES_ONE_BOUNDARY_FACTOR_AT_A_TIME"
	SupportDegreeOneExposureFromExteriorDegree  = "CONDITIONAL_SUPPORT_DEGREE_ONE_HAS_EXPOSURE_TYPE_FROM_EXTERIOR_DEGREE"
	SupportLambda1ExposureNativeShape           = "CONDITIONAL_SUPPORT_LAMBDA1B2_AS_SINGLE_BOUNDARY_EXPOSURE_HAS_NATIVE_EXTERIOR_SHAPE"
	SupportLambda2TopPairSpace                  = "CONDITIONAL_SUPPORT_LAMBDA2B2_IS_TOP_BOUNDARY_PAIR_EXTERIOR_SPACE"
	SupportDegreeTwoRequiresBothFactors         = "CONDITIONAL_SUPPORT_DEGREE_TWO_REQUIRES_BOTH_BOUNDARY_FACTORS"
	SupportDegreeTwoFullEnclosureFromTopDegree  = "CONDITIONAL_SUPPORT_DEGREE_TWO_HAS_FULL_ENCLOSURE_TYPE_FROM_TOP_EXTERIOR_DEGREE"
	SupportLambda2EnclosureNativeShape          = "CONDITIONAL_SUPPORT_LAMBDA2B2_AS_FULL_BOUNDARY_PAIR_ENCLOSURE_HAS_NATIVE_EXTERIOR_SHAPE"
	SupportExposureEnclosureGrounded            = "CONDITIONAL_SUPPORT_EXPOSURE_ENCLOSURE_CONTRAST_IS_GROUNDED_IN_EXTERIOR_DEGREE"
	SupportOneVsTwoFactorDistinguishesDegrees   = "CONDITIONAL_SUPPORT_ONE_FACTOR_VS_TWO_FACTOR_BOUNDARY_ACTIVATION_DISTINGUISHES_DEGREE_ONE_AND_TWO"
	SupportBoundaryDegreeTypesNotRandom         = "CONDITIONAL_SUPPORT_BOUNDARY_DEGREE_TYPES_ARE_NOT_RANDOM_LABELS"
	SupportTopDegreeCumulativeEnclosure         = "CONDITIONAL_SUPPORT_TOP_DEGREE_PAIR_ACTIVATION_SOURCE_TYPES_CUMULATIVE_ENCLOSURE"
	SupportFullEnclosurePointsToF2OverF0        = "CONDITIONAL_SUPPORT_FULL_ENCLOSURE_POINTS_TO_F2_OVER_F0_RATHER_THAN_F2_OVER_F1"
	SupportCumulativeHasTopDegreeSource         = "CONDITIONAL_SUPPORT_CUMULATIVE_ENCLOSURE_HAS_EXTERIOR_TOP_DEGREE_SOURCE_CANDIDATE"
	SupportExposureEnclosureStrengthensSelector = "CONDITIONAL_SUPPORT_EXPOSURE_ENCLOSURE_TYPING_STRENGTHENS_SELECTOR_FUNCTIONHOOD"
	SupportSelectorReducedToExteriorDegree      = "CONDITIONAL_SUPPORT_SELECTOR_SOURCE_NOW_REDUCED_TO_EXTERIOR_DEGREE_TYPE"
	SupportIBZ2TargetsCompatible                = "CONDITIONAL_SUPPORT_I_B_Z2_TARGETS_ARE_COMPATIBLE_WITH_BOUNDARY_DEGREE_TYPES"
	SupportMuBSelectorInputExteriorSource       = "CONDITIONAL_SUPPORT_BOUNDARY_ACTIVATION_MEASURE_SELECTOR_INPUT_HAS_EXTERIOR_DEGREE_SOURCE"
	SupportMuBFunctionhoodGapWeakened           = "CONDITIONAL_SUPPORT_MU_B_FUNCTIONHOOD_GAP_WEAKENED_BY_NATIVE_EXTERIOR_DEGREE_TYPING"
	SupportAlphaMeasureStrongerSelectorTyping   = "CONDITIONAL_SUPPORT_ALPHA_MEASURE_NOW_HAS_STRONGER_SELECTOR_SOURCE_TYPING"

	FailureExposureLanguageNotTargetFunctor  = "FAILED_ROUTE_EXPOSURE_LANGUAGE_IS_TYPE_INTERPRETATION_NOT_FULL_NATIVE_TARGET_FUNCTOR"
	FailureEnclosureLanguageNotTargetFunctor = "FAILED_ROUTE_ENCLOSURE_LANGUAGE_IS_TYPE_INTERPRETATION_NOT_FULL_NATIVE_TARGET_FUNCTOR"
	FailureExteriorContrastDoesNotSelectZ2   = "FAILED_ROUTE_EXTERIOR_DEGREE_TYPE_CONTRAST_DOES_NOT_BY_ITSELF_SELECT_Z2_FLAG_TARGETS"
	FailureNoNativeTopDegreeToF2OverF0       = "FAILED_ROUTE_NO_NATIVE_FUNCTOR_FROM_TOP_EXTERIOR_DEGREE_TO_F2_OVER_F0_CERTIFIED"
	FailureSelectorFunctionhoodTargetFunctor = "FAILED_ROUTE_SELECTOR_FUNCTIONHOOD_STILL_NOT_NATIVE_WITHOUT_DEGREE_TO_FLAG_FUNCTOR"
	FailureMuBStillNotNativeTargetFunctor    = "FAILED_ROUTE_MU_B_STILL_NOT_NATIVE_WITHOUT_FULL_DEGREE_TO_FLAG_FUNCTOR"
	FailureAlphaBridgeCandidateNotNative     = "FAILED_ROUTE_ALPHA_B_REMAINS_BRIDGE_CANDIDATE_NOT_NATIVE"
	FailureNotNativeR3                       = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked         = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap            = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap            = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues          = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator            = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type DegreeExteriorTypingAudit struct {
	Degree          int
	ExteriorSpace   string
	ResponseTerm    string
	BoundaryFactors int
	TopDegree       bool
	Interpretation  string
	SourceStatus    string
	NativeShape     bool
	NativeTargetMap bool
	Supports        []string
	Failures        []string
}

type ExteriorContrastAudit struct {
	DegreeOneType    string
	DegreeTwoType    string
	GroundedInDegree bool
	ArbitraryLabels  bool
	SelectsZ2Targets bool
	Supports         []string
	Failures         []string
}

type CumulativeEnclosureAudit struct {
	TopDegreeTerm            string
	CumulativeTarget         string
	AssociatedGradedRejected string
	RankCumulative           int
	RankAssociatedGraded     int
	TopDegreeSourceCandidate bool
	NativeTargetFunctor      bool
	Supports                 []string
	Failures                 []string
}

type SelectorConsequenceAudit struct {
	SelectorOne              string
	SelectorTwo              string
	StrengthenedByDegreeType bool
	TargetFunctorNative      bool
	Supports                 []string
	Failures                 []string
}

type MuBConsequenceAudit struct {
	MeasureFormula             string
	SelectorInputExteriorTyped bool
	FunctionhoodGapWeakened    bool
	NativeMuB                  bool
	Supports                   []string
	Failures                   []string
}

type FirewallLedger struct {
	ExposureLanguageNotTargetFunctor  bool
	EnclosureLanguageNotTargetFunctor bool
	ExteriorContrastDoesNotSelectZ2   bool
	NoNativeTopDegreeToF2OverF0       bool
	SelectorFunctionhoodTargetFunctor bool
	MuBStillNotNativeTargetFunctor    bool
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
	if f.ExposureLanguageNotTargetFunctor {
		out = append(out, FailureExposureLanguageNotTargetFunctor)
	}
	if f.EnclosureLanguageNotTargetFunctor {
		out = append(out, FailureEnclosureLanguageNotTargetFunctor)
	}
	if f.ExteriorContrastDoesNotSelectZ2 {
		out = append(out, FailureExteriorContrastDoesNotSelectZ2)
	}
	if f.NoNativeTopDegreeToF2OverF0 {
		out = append(out, FailureNoNativeTopDegreeToF2OverF0)
	}
	if f.SelectorFunctionhoodTargetFunctor {
		out = append(out, FailureSelectorFunctionhoodTargetFunctor)
	}
	if f.MuBStillNotNativeTargetFunctor {
		out = append(out, FailureMuBStillNotNativeTargetFunctor)
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
	DegreeOne       DegreeExteriorTypingAudit
	DegreeTwo       DegreeExteriorTypingAudit
	Contrast        ExteriorContrastAudit
	Cumulative      CumulativeEnclosureAudit
	Selector        SelectorConsequenceAudit
	MuB             MuBConsequenceAudit
	Firewalls       FirewallLedger
	Final           string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		ID:              AuditID,
		InheritedStatus: Gate923ShortStatus,
		Truth:           FinalTruth,
		Classification:  Classification,
		ShortStatus:     ShortStatus,
		DegreeOne: DegreeExteriorTypingAudit{
			Degree:          1,
			ExteriorSpace:   "Lambda^1 B_2=span{b1,b2}",
			ResponseTerm:    DegreeOneTerm,
			BoundaryFactors: 1,
			TopDegree:       false,
			Interpretation:  "single-boundary exposure",
			SourceStatus:    SourceNativeExteriorShape,
			NativeShape:     true,
			NativeTargetMap: false,
			Supports:        []string{SupportLambda1SingleGenerator, SupportDegreeOneActivatesOneFactor, SupportDegreeOneExposureFromExteriorDegree, SupportLambda1ExposureNativeShape},
			Failures:        []string{FailureExposureLanguageNotTargetFunctor},
		},
		DegreeTwo: DegreeExteriorTypingAudit{
			Degree:          2,
			ExteriorSpace:   "Lambda^2 B_2=span{b1 wedge b2}",
			ResponseTerm:    DegreeTwoTerm,
			BoundaryFactors: 2,
			TopDegree:       true,
			Interpretation:  "full boundary-pair enclosure",
			SourceStatus:    SourceNativeExteriorShape,
			NativeShape:     true,
			NativeTargetMap: false,
			Supports:        []string{SupportLambda2TopPairSpace, SupportDegreeTwoRequiresBothFactors, SupportDegreeTwoFullEnclosureFromTopDegree, SupportLambda2EnclosureNativeShape},
			Failures:        []string{FailureEnclosureLanguageNotTargetFunctor},
		},
		Contrast: ExteriorContrastAudit{
			DegreeOneType:    "one-factor boundary activation / exposure",
			DegreeTwoType:    "two-factor top exterior activation / enclosure",
			GroundedInDegree: true,
			ArbitraryLabels:  false,
			SelectsZ2Targets: false,
			Supports:         []string{SupportExposureEnclosureGrounded, SupportOneVsTwoFactorDistinguishesDegrees, SupportBoundaryDegreeTypesNotRandom},
			Failures:         []string{FailureExteriorContrastDoesNotSelectZ2},
		},
		Cumulative: CumulativeEnclosureAudit{
			TopDegreeTerm:            DegreeTwoTerm,
			CumulativeTarget:         "[F_2/F_0]_{Z2}",
			AssociatedGradedRejected: "F_2/F_1",
			RankCumulative:           RankF2OverF0,
			RankAssociatedGraded:     RankF2OverF1,
			TopDegreeSourceCandidate: true,
			NativeTargetFunctor:      false,
			Supports:                 []string{SupportTopDegreeCumulativeEnclosure, SupportFullEnclosurePointsToF2OverF0, SupportCumulativeHasTopDegreeSource},
			Failures:                 []string{FailureNoNativeTopDegreeToF2OverF0},
		},
		Selector: SelectorConsequenceAudit{
			SelectorOne:              "I_B^Z2(1)=[F_1/F_0]_{Z2}",
			SelectorTwo:              "I_B^Z2(2)=[F_2/F_0]_{Z2}",
			StrengthenedByDegreeType: true,
			TargetFunctorNative:      false,
			Supports:                 []string{SupportExposureEnclosureStrengthensSelector, SupportSelectorReducedToExteriorDegree, SupportIBZ2TargetsCompatible},
			Failures:                 []string{FailureSelectorFunctionhoodTargetFunctor},
		},
		MuB: MuBConsequenceAudit{
			MeasureFormula:             "mu_B(R_B(S_split))=sum_{k=1}^2 rank(I_B^Z2(k))/rank(H_k)*S_split^k",
			SelectorInputExteriorTyped: true,
			FunctionhoodGapWeakened:    true,
			NativeMuB:                  false,
			Supports:                   []string{SupportMuBSelectorInputExteriorSource, SupportMuBFunctionhoodGapWeakened, SupportAlphaMeasureStrongerSelectorTyping},
			Failures:                   []string{FailureMuBStillNotNativeTargetFunctor, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3},
		},
		Firewalls: FirewallLedger{
			ExposureLanguageNotTargetFunctor:  true,
			EnclosureLanguageNotTargetFunctor: true,
			ExteriorContrastDoesNotSelectZ2:   true,
			NoNativeTopDegreeToF2OverF0:       true,
			SelectorFunctionhoodTargetFunctor: true,
			MuBStillNotNativeTargetFunctor:    true,
			AlphaBridgeCandidateNotNative:     true,
			NotNativeR3:                       true,
			FullAFDescentStillBlocked:         true,
			NoGenerationCarrierMap:            true,
			NoFlavorOrientationMap:            true,
			NoIndividualYukawaValues:          true,
			NoNativeYukawaOperator:            true,
		},
		Final: "Gate 924 supports exposure/enclosure as native exterior-degree shape of B_2, but the native target functor from exterior degree types to Z2 airlock flag quotients remains missing.",
	}
	if err := validate(a); err != nil {
		return Audit{}, err
	}
	return a, nil
}

func validate(a Audit) error {
	if a.InheritedStatus != Gate923ShortStatus {
		return fmt.Errorf("unexpected inherited status: %s", a.InheritedStatus)
	}
	if !a.DegreeOne.NativeShape || a.DegreeOne.NativeTargetMap {
		return fmt.Errorf("degree one must be native shape but not native target map: %s", FormatDegree(a.DegreeOne))
	}
	if !a.DegreeTwo.NativeShape || !a.DegreeTwo.TopDegree || a.DegreeTwo.NativeTargetMap {
		return fmt.Errorf("degree two must be top-degree native shape but not native target map: %s", FormatDegree(a.DegreeTwo))
	}
	if !a.Contrast.GroundedInDegree || a.Contrast.ArbitraryLabels || a.Contrast.SelectsZ2Targets {
		return fmt.Errorf("bad contrast audit: %s", FormatContrast(a.Contrast))
	}
	if !a.Cumulative.TopDegreeSourceCandidate || a.Cumulative.NativeTargetFunctor || a.Cumulative.RankCumulative != RankF2OverF0 || a.Cumulative.RankAssociatedGraded != RankF2OverF1 {
		return fmt.Errorf("bad cumulative audit: %s", FormatCumulative(a.Cumulative))
	}
	if !a.Selector.StrengthenedByDegreeType || a.Selector.TargetFunctorNative {
		return fmt.Errorf("bad selector consequence: %s", FormatSelector(a.Selector))
	}
	if !a.MuB.SelectorInputExteriorTyped || !a.MuB.FunctionhoodGapWeakened || a.MuB.NativeMuB {
		return fmt.Errorf("bad mu_B consequence: %s", FormatMuB(a.MuB))
	}
	if !firewallsOK(a.Firewalls) {
		return fmt.Errorf("firewall leak: %s", FormatFirewalls(a.Firewalls))
	}
	return nil
}

func Statuses() []string {
	return []string{StatusInheritedGate923, StatusDegreeOneNativeShape, StatusDegreeTwoNativeShape, StatusDegreeContrastGrounded, StatusCumulativePressure, StatusSelectorStrengthened, StatusMuBSelectorInputStronger, StatusTargetFunctorMissing, StatusNativeR3Blocked}
}

func Supports() []string {
	return []string{SupportLambda1SingleGenerator, SupportDegreeOneActivatesOneFactor, SupportDegreeOneExposureFromExteriorDegree, SupportLambda1ExposureNativeShape, SupportLambda2TopPairSpace, SupportDegreeTwoRequiresBothFactors, SupportDegreeTwoFullEnclosureFromTopDegree, SupportLambda2EnclosureNativeShape, SupportExposureEnclosureGrounded, SupportOneVsTwoFactorDistinguishesDegrees, SupportBoundaryDegreeTypesNotRandom, SupportTopDegreeCumulativeEnclosure, SupportFullEnclosurePointsToF2OverF0, SupportCumulativeHasTopDegreeSource, SupportExposureEnclosureStrengthensSelector, SupportSelectorReducedToExteriorDegree, SupportIBZ2TargetsCompatible, SupportMuBSelectorInputExteriorSource, SupportMuBFunctionhoodGapWeakened, SupportAlphaMeasureStrongerSelectorTyping}
}

func Failures() []string {
	return []string{FailureExposureLanguageNotTargetFunctor, FailureEnclosureLanguageNotTargetFunctor, FailureExteriorContrastDoesNotSelectZ2, FailureNoNativeTopDegreeToF2OverF0, FailureSelectorFunctionhoodTargetFunctor, FailureMuBStillNotNativeTargetFunctor, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}
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

func FormatDegree(d DegreeExteriorTypingAudit) string {
	return fmt.Sprintf("degree=%d exterior=%s term=%s factors=%d top=%t interpretation=%s source=%s native_shape=%t native_target_map=%t", d.Degree, d.ExteriorSpace, d.ResponseTerm, d.BoundaryFactors, d.TopDegree, d.Interpretation, d.SourceStatus, d.NativeShape, d.NativeTargetMap)
}
func FormatContrast(c ExteriorContrastAudit) string {
	return fmt.Sprintf("degree_one=%s degree_two=%s grounded=%t arbitrary=%t selects_z2_targets=%t", c.DegreeOneType, c.DegreeTwoType, c.GroundedInDegree, c.ArbitraryLabels, c.SelectsZ2Targets)
}
func FormatCumulative(c CumulativeEnclosureAudit) string {
	return fmt.Sprintf("top_term=%s cumulative=%s rank_cumulative=%d rejected=%s rank_rejected=%d top_degree_source=%t native_target_functor=%t", c.TopDegreeTerm, c.CumulativeTarget, c.RankCumulative, c.AssociatedGradedRejected, c.RankAssociatedGraded, c.TopDegreeSourceCandidate, c.NativeTargetFunctor)
}
func FormatSelector(s SelectorConsequenceAudit) string {
	return fmt.Sprintf("selector_one=%s selector_two=%s strengthened_by_degree_type=%t native_target_functor=%t", s.SelectorOne, s.SelectorTwo, s.StrengthenedByDegreeType, s.TargetFunctorNative)
}
func FormatMuB(m MuBConsequenceAudit) string {
	return fmt.Sprintf("formula=%s selector_input_exterior_typed=%t functionhood_gap_weakened=%t native_muB=%t", m.MeasureFormula, m.SelectorInputExteriorTyped, m.FunctionhoodGapWeakened, m.NativeMuB)
}
func FormatFirewalls(f FirewallLedger) string { return strings.Join(f.List(), ";") }
