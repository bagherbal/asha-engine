// Package generation2z2boundaryalphafunctorsourcedecompositionaudit
// implements Gate 912: Z2 BoundaryAlpha Functor Source Decomposition Audit.
//
// Gate 912 follows Gate 911's rail selection. It does not attempt to prove the
// native Z2 BoundaryAlpha functor in one step. Instead, it decomposes the
// missing native theorem into five exact sub-objects: reduced B2 response
// functional, degree-to-Z2-flag-class selector, cross-lane exclusion theorem,
// S_split transport law, and boundary-augmented chamber normalization. The gate
// preserves the R3 class seal and leaves alpha_B sealed until all sub-objects are
// certified.
package generation2z2boundaryalphafunctorsourcedecompositionaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE912-Z2-BOUNDARY-ALPHA-FUNCTOR-SOURCE-DECOMPOSITION-AUDIT"

	SBoundary = 0.0012924448188162962
	AlphaB    = 0.0003878958469680527

	RankF1OverF0 = 3
	RankF2OverF0 = 7
	LinearDenom  = 10
	QuadDenom    = 72
	AmbientRank8 = 8
	BoundaryRank = 2
	Lambda4Rank  = 70

	Gate911Classification = "R3_FRONTIER_SELECTED_Z2_BOUNDARY_ALPHA_FUNCTOR_BEFORE_FULL_AF_DESCENT"
	Gate911ShortStatus    = "R3_NEXT_RAIL_NATIVE_Z2_BOUNDARY_ALPHA_FUNCTOR"
	Gate911Verdict        = "FRONTIER_A_SELECTED_NATIVE_Z2_BOUNDARY_ALPHA_FUNCTOR_FIRST"

	PunctureClass             = "[p]_{Z2}={e_lambda tensor P_1,e_barlambda tensor P_1}"
	BoundaryAlphaTarget       = "BoundaryAlpha_Z2:[p]_{Z2}->(3/10)s+(7/72)s^2"
	BoundaryAlphaFormula      = "alpha_B^Z2=[rank([F_1/F_0]_{Z2})/10]s+[rank([F_2/F_0]_{Z2})/72]s^2"
	ReducedB2Response         = "R_B(s)=(1+s b1)(1+s b2)-1=s(b1+b2)+s^2(b1 wedge b2)"
	Lambda1Target             = "deg(Lambda^1 B_2)->[F_1/F_0]_{Z2}"
	Lambda2Target             = "deg(Lambda^2 B_2)->[F_2/F_0]_{Z2}"
	ForbiddenLinearFull       = "Lambda^1 B_2 not -> [F_2/F_0]_{Z2}"
	ForbiddenQuadraticExposed = "Lambda^2 B_2 not -> [F_1/F_0]_{Z2}"
	WrongLinearFullTerm       = "(7/72)s"
	WrongQuadraticExposedTerm = "(3/10)s^2"
	H10Chamber                = "H_10=H_R^{ambient} plus B_2, rank=8+2=10"
	H72Chamber                = "H_72=Lambda^4 V_8 plus B_2, rank=70+2=72"
	NativeFunctorObject       = "Z2EquivariantNeutralPunctureAirlockFunctor"
	NativeReducedB2Theorem    = "NativeReducedBoundaryPairResponseFunctional"
	DegreeSelectorTheorem     = "DegreeIndexedZ2AirlockFlagFunctor"
	CrossLaneTheorem          = "Z2BoundaryAlphaCrossLaneExclusionTheorem"
	SsplitTransportTheorem    = "SsplitToZ2BoundaryResponseTransportLaw"
	DenominatorTypingTheorem  = "BoundaryAugmentedResponseChamberNormalizationTheorem"
	NextGate                  = "NEXT_PRESSURE_GATE913_NATIVE_REDUCED_BOUNDARYPAIR_RESPONSE_FUNCTIONAL_AUDIT"
	Classification            = "R3_Z2_BOUNDARY_ALPHA_FUNCTOR_SOURCE_DECOMPOSED_NOT_NATIVE"
	ShortStatus               = "R3_ALPHA_FUNCTOR_DECOMPOSITION_COMPLETE_NATIVE_SUBOBJECTS_MISSING"
	FinalTruth                = "Z2_BOUNDARYALPHA_FUNCTOR_DECOMPOSED_INTO_FIVE_REQUIRED_NATIVE_SUBOBJECTS"
	StrategicConclusion       = "The native Z2 BoundaryAlpha wound is decomposed into five exact theorem requirements; Gate 913 should attack the reduced B2 response functional first."

	StatusGate911Inherited        = "PASS_GATE911_NATIVE_Z2_BOUNDARY_ALPHA_RAIL_INHERITED"
	StatusNoLoopBack              = "PASS_BRANCH_DOES_NOT_REOPEN_PHASE_SOCKET_OR_REPRESENTATIVE_ALPHA_WOUNDS"
	StatusSealedFormulaReproduced = "PASS_BOUNDARY_ALPHA_Z2_SEALED_FORMULA_REPRODUCED"
	StatusDecompositionComplete   = "PASS_NATIVE_Z2_BOUNDARYALPHA_FUNCTOR_DECOMPOSED_INTO_FIVE_SUBOBJECTS"
	StatusReducedB2Shape          = "PASS_REDUCED_B2_RESPONSE_SHAPE_IDENTIFIED_AS_REQUIRED_SUBOBJECT"
	StatusDegreeSelectorRequired  = "PASS_DEGREE_TO_Z2_FLAG_CLASS_SELECTOR_IDENTIFIED_AS_REQUIRED_SUBOBJECT"
	StatusCrossLaneRequired       = "PASS_CROSS_LANE_EXCLUSION_IDENTIFIED_AS_REQUIRED_SUBOBJECT"
	StatusSsplitTransportRequired = "PASS_S_SPLIT_TRANSPORT_IDENTIFIED_AS_REQUIRED_SUBOBJECT"
	StatusDenominatorTyping       = "PASS_DENOMINATOR_CHAMBERS_TYPED_AS_REQUIRED_SUBOBJECT"
	StatusNativeMissing           = "FIREWALL_PRESERVED_GATE912_NATIVE_SUBOBJECTS_MISSING"

	SupportNativeZ2AlphaDecomposed     = "CONDITIONAL_SUPPORT_NATIVE_Z2_BOUNDARYALPHA_FUNCTOR_DECOMPOSED"
	SupportReducedB2Required           = "CONDITIONAL_SUPPORT_REDUCED_B2_RESPONSE_IS_REQUIRED_SUBOBJECT"
	SupportReducedB2CorrectShape       = "CONDITIONAL_SUPPORT_REDUCED_B2_RESPONSE_HAS_CORRECT_S_PLUS_S_SQUARED_SHAPE"
	SupportZeroOrderSuppressed         = "CONDITIONAL_SUPPORT_ZERO_ORDER_SUPPRESSED_BY_REDUCED_RESPONSE"
	SupportCubicAbsent                 = "CONDITIONAL_SUPPORT_CUBIC_AND_HIGHER_TERMS_ABSENT_BY_LAMBDA3_B2_ZERO"
	SupportDegreeSelectorRequired      = "CONDITIONAL_SUPPORT_DEGREE_TO_Z2_FLAG_CLASS_SELECTOR_IS_REQUIRED_SUBOBJECT"
	SupportDegreeOneExposed            = "CONDITIONAL_SUPPORT_DEGREE_ONE_TARGETS_Z2_EXPOSED_FACE_CLASS"
	SupportDegreeTwoFull               = "CONDITIONAL_SUPPORT_DEGREE_TWO_TARGETS_Z2_FULL_ENCLOSURE_CLASS"
	SupportRankPairRepresentativeFree  = "CONDITIONAL_SUPPORT_TARGET_RANK_PAIR_3_7_IS_Z2_REPRESENTATIVE_INDEPENDENT"
	SupportCrossLaneRequired           = "CONDITIONAL_SUPPORT_CROSS_LANE_EXCLUSION_IS_REQUIRED_SUBOBJECT"
	SupportCrossLanesExcludedIfFunctor = "CONDITIONAL_SUPPORT_CROSS_LANES_EXCLUDED_IF_DEGREE_INDEXED_FUNCTOR_CERTIFIED"
	SupportSsplitTransportRequired     = "CONDITIONAL_SUPPORT_S_SPLIT_TRANSPORT_IS_REQUIRED_SUBOBJECT"
	SupportSsplitFeedsDegreeShape      = "CONDITIONAL_SUPPORT_S_SPLIT_FEEDS_DEGREE_ONE_AND_TWO_RESPONSE_SHAPE"
	SupportDenominatorChambersRequired = "CONDITIONAL_SUPPORT_BOUNDARY_AUGMENTED_DENOMINATOR_CHAMBERS_ARE_REQUIRED_SUBOBJECTS"
	SupportDenominatorsTyped           = "CONDITIONAL_SUPPORT_DENOMINATORS_TYPED_AS_BOUNDARY_AUGMENTED_RESPONSE_CHAMBERS"
	SupportAlphaReducedToFive          = "CONDITIONAL_SUPPORT_ALPHA_B_NATIVE_STATUS_REDUCED_TO_FIVE_EXACT_THEOREM_REQUIREMENTS"

	FailureNoNativeZ2BoundaryAlphaFunctor = "FAILED_ROUTE_NO_NATIVE_Z2_BOUNDARY_ALPHA_FUNCTOR"
	FailureReducedB2NotNativeFunctional   = "FAILED_ROUTE_REDUCED_B2_RESPONSE_NOT_NATIVE_BOUNDARY_FUNCTIONAL"
	FailureNoNativeReasonEBMinusOne       = "FAILED_ROUTE_NO_NATIVE_REASON_FOR_USING_E_B_MINUS_ONE_RESPONSE"
	FailureNoNativeTransportSInB2         = "FAILED_ROUTE_NO_NATIVE_TRANSPORT_OF_S_SPLIT_INTO_REDUCED_B2_RESPONSE"
	FailureNoNativeDegreeToZ2FlagFunctor  = "FAILED_ROUTE_NO_NATIVE_DEGREE_TO_Z2_FLAG_CLASS_FUNCTOR"
	FailureNoNativeLambda1ExposedMap      = "FAILED_ROUTE_NO_NATIVE_LAMBDA1B2_TO_EXPOSED_FACE_CLASS_MAP"
	FailureNoNativeLambda2FullMap         = "FAILED_ROUTE_NO_NATIVE_LAMBDA2B2_TO_FULL_ENCLOSURE_CLASS_MAP"
	FailureNoNativeZ2CrossLane            = "FAILED_ROUTE_NO_NATIVE_Z2_CROSS_LANE_EXCLUSION_THEOREM"
	FailureNoNativeLinearDomainExclusion  = "FAILED_ROUTE_NO_NATIVE_EXCLUSION_OF_LINEAR_ACTIVE_DOMAIN_CLASS"
	FailureNoNativeQuadraticFaceExclusion = "FAILED_ROUTE_NO_NATIVE_EXCLUSION_OF_QUADRATIC_EXPOSED_FACE_CLASS"
	FailureNoNativeTransportS             = "FAILED_ROUTE_NO_NATIVE_TRANSPORT_OF_S_SPLIT_TO_Z2_AIRLOCK_CLASS"
	FailureNoTypedSToLambda1              = "FAILED_ROUTE_NO_TYPED_S_SPLIT_TO_LAMBDA1B2_RESPONSE"
	FailureNoTypedS2ToLambda2             = "FAILED_ROUTE_NO_TYPED_S_SPLIT_SQUARED_TO_LAMBDA2B2_RESPONSE"
	FailureDenominatorNotActivation       = "FAILED_ROUTE_DENOMINATOR_TYPING_NOT_BOUNDARY_ALPHA_ACTIVATION_THEOREM"
	FailureAlphaStillSealed               = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNotNativeR3                    = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked      = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap         = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap         = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues       = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate           = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator         = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type InheritedRail struct {
	Gate911Classification string
	Gate911ShortStatus    string
	Gate911Verdict        string
	SelectedRail          string
	ReopensPhaseSign      bool
	ReopensSocketOrder    bool
	ReopensRepAlpha       bool
	DerivesAlpha          bool
	UpdatesOfficialLedger bool
	Supports, Failures    []string
}

type SealedFormula struct {
	PunctureClass         string
	Target                string
	Formula               string
	S                     float64
	Alpha                 float64
	LinearContribution    float64
	QuadraticContribution float64
	RankPair              [2]int
	Denominators          [2]int
	RepresentativeFree    bool
	Native                bool
	Supports, Failures    []string
}

type RequiredSubobject struct {
	Index              int
	Name               string
	RequiredTheorem    string
	Description        string
	CertifiedNative    bool
	Required           bool
	CorrectShape       bool
	RankInputs         []int
	Denominators       []int
	Targets            []string
	ForbiddenTargets   []string
	WrongTerms         []string
	Supports, Failures []string
}

type Decomposition struct {
	NativeFunctorTarget string
	Subobjects          []RequiredSubobject
	RequiredCount       int
	CertifiedCount      int
	MissingCount        int
	NativeFunctor       bool
	AlphaNative         bool
	R3Native            bool
	Supports, Failures  []string
}

type Firewalls struct {
	NativeZ2BoundaryAlphaFunctor bool
	ReducedB2NativeFunctional    bool
	DegreeToZ2FlagFunctor        bool
	NativeZ2CrossLane            bool
	NativeTransportS             bool
	DenominatorActivationLaw     bool
	AlphaNative                  bool
	NativeR3                     bool
	FullAFDescent                bool
	GenerationCarrierMap         bool
	FlavorOrientationMap         bool
	IndividualYukawaValues       bool
	OfficialLedgerUpdate         bool
	NativeYukawaOperator         bool
}

type Audit struct {
	ID             string
	Classification string
	ShortStatus    string
	Inherited      InheritedRail
	Formula        SealedFormula
	Decomposition  Decomposition
	Firewalls      Firewalls
	Truth          string
	Final          string
}

func LinearContribution(s float64) float64 {
	return float64(RankF1OverF0) / float64(LinearDenom) * s
}

func QuadraticContribution(s float64) float64 {
	return float64(RankF2OverF0) / float64(QuadDenom) * s * s
}

func BoundaryAlphaZ2(s float64) float64 { return LinearContribution(s) + QuadraticContribution(s) }

func BuildDefault() (Audit, error) {
	inherited := buildInheritedRail()
	if inherited.ReopensPhaseSign || inherited.ReopensSocketOrder || inherited.ReopensRepAlpha || inherited.DerivesAlpha || inherited.UpdatesOfficialLedger {
		return Audit{}, fmt.Errorf("inherited rail reopened a closed wound: %s", FormatInherited(inherited))
	}

	formula := buildSealedFormula()
	if !near(formula.Alpha, AlphaB) || !formula.RepresentativeFree || formula.Native {
		return Audit{}, fmt.Errorf("sealed formula classification leak: %s", FormatFormula(formula))
	}

	decomp := buildDecomposition()
	if decomp.RequiredCount != 5 || decomp.CertifiedCount != 0 || decomp.MissingCount != 5 || decomp.NativeFunctor || decomp.AlphaNative || decomp.R3Native {
		return Audit{}, fmt.Errorf("bad decomposition accounting: %s", FormatDecomposition(decomp))
	}
	if !allRequiredSubobjectsPresent(decomp.Subobjects) {
		return Audit{}, fmt.Errorf("missing required subobject: %s", FormatDecomposition(decomp))
	}

	firewalls := buildFirewalls()
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewall leak: %s", FormatFirewalls(firewalls))
	}

	return Audit{ID: AuditID, Classification: Classification, ShortStatus: ShortStatus, Inherited: inherited, Formula: formula, Decomposition: decomp, Firewalls: firewalls, Truth: FinalTruth, Final: StrategicConclusion}, nil
}

func buildInheritedRail() InheritedRail {
	return InheritedRail{
		Gate911Classification: Gate911Classification,
		Gate911ShortStatus:    Gate911ShortStatus,
		Gate911Verdict:        Gate911Verdict,
		SelectedRail:          NativeFunctorObject,
		ReopensPhaseSign:      false,
		ReopensSocketOrder:    false,
		ReopensRepAlpha:       false,
		DerivesAlpha:          false,
		UpdatesOfficialLedger: false,
		Supports: []string{
			StatusGate911Inherited,
			StatusNoLoopBack,
		},
		Failures: []string{
			FailureAlphaStillSealed,
			FailureNoNativeZ2BoundaryAlphaFunctor,
		},
	}
}

func buildSealedFormula() SealedFormula {
	return SealedFormula{
		PunctureClass:         PunctureClass,
		Target:                BoundaryAlphaTarget,
		Formula:               BoundaryAlphaFormula,
		S:                     SBoundary,
		Alpha:                 BoundaryAlphaZ2(SBoundary),
		LinearContribution:    LinearContribution(SBoundary),
		QuadraticContribution: QuadraticContribution(SBoundary),
		RankPair:              [2]int{RankF1OverF0, RankF2OverF0},
		Denominators:          [2]int{LinearDenom, QuadDenom},
		RepresentativeFree:    true,
		Native:                false,
		Supports: []string{
			StatusSealedFormulaReproduced,
			SupportRankPairRepresentativeFree,
		},
		Failures: []string{
			FailureAlphaStillSealed,
			FailureNoNativeZ2BoundaryAlphaFunctor,
		},
	}
}

func buildDecomposition() Decomposition {
	subobjects := []RequiredSubobject{
		buildReducedB2Subobject(),
		buildDegreeSelectorSubobject(),
		buildCrossLaneSubobject(),
		buildSsplitTransportSubobject(),
		buildDenominatorSubobject(),
	}
	certified := 0
	for _, s := range subobjects {
		if s.CertifiedNative {
			certified++
		}
	}
	return Decomposition{
		NativeFunctorTarget: NativeFunctorObject,
		Subobjects:          subobjects,
		RequiredCount:       len(subobjects),
		CertifiedCount:      certified,
		MissingCount:        len(subobjects) - certified,
		NativeFunctor:       false,
		AlphaNative:         false,
		R3Native:            false,
		Supports: []string{
			StatusDecompositionComplete,
			SupportNativeZ2AlphaDecomposed,
			SupportAlphaReducedToFive,
		},
		Failures: []string{
			FailureNoNativeZ2BoundaryAlphaFunctor,
			FailureAlphaStillSealed,
			FailureNotNativeR3,
		},
	}
}

func buildReducedB2Subobject() RequiredSubobject {
	return RequiredSubobject{
		Index:           1,
		Name:            "Reduced B2 response functional",
		RequiredTheorem: NativeReducedB2Theorem,
		Description:     ReducedB2Response,
		CertifiedNative: false,
		Required:        true,
		CorrectShape:    true,
		Targets:         []string{"Lambda^0 suppressed", "Lambda^1 B_2", "Lambda^2 B_2", "Lambda^3 B_2=0"},
		Supports: []string{
			StatusReducedB2Shape,
			SupportReducedB2Required,
			SupportReducedB2CorrectShape,
			SupportZeroOrderSuppressed,
			SupportCubicAbsent,
		},
		Failures: []string{
			FailureReducedB2NotNativeFunctional,
			FailureNoNativeReasonEBMinusOne,
			FailureNoNativeTransportSInB2,
		},
	}
}

func buildDegreeSelectorSubobject() RequiredSubobject {
	return RequiredSubobject{
		Index:           2,
		Name:            "Degree-to-Z2-flag-class selector",
		RequiredTheorem: DegreeSelectorTheorem,
		Description:     "degree-indexed incidence selector from exterior boundary degrees to Z2 airlock quotient classes",
		CertifiedNative: false,
		Required:        true,
		CorrectShape:    true,
		RankInputs:      []int{RankF1OverF0, RankF2OverF0},
		Targets:         []string{Lambda1Target, Lambda2Target},
		Supports: []string{
			StatusDegreeSelectorRequired,
			SupportDegreeSelectorRequired,
			SupportDegreeOneExposed,
			SupportDegreeTwoFull,
			SupportRankPairRepresentativeFree,
		},
		Failures: []string{
			FailureNoNativeDegreeToZ2FlagFunctor,
			FailureNoNativeLambda1ExposedMap,
			FailureNoNativeLambda2FullMap,
		},
	}
}

func buildCrossLaneSubobject() RequiredSubobject {
	return RequiredSubobject{
		Index:            3,
		Name:             "Cross-lane exclusion theorem",
		RequiredTheorem:  CrossLaneTheorem,
		Description:      "degree one must not activate full-enclosure class and degree two must not activate exposed-face class",
		CertifiedNative:  false,
		Required:         true,
		CorrectShape:     true,
		Targets:          []string{Lambda1Target, Lambda2Target},
		ForbiddenTargets: []string{ForbiddenLinearFull, ForbiddenQuadraticExposed},
		WrongTerms:       []string{WrongLinearFullTerm, WrongQuadraticExposedTerm},
		Supports: []string{
			StatusCrossLaneRequired,
			SupportCrossLaneRequired,
			SupportCrossLanesExcludedIfFunctor,
		},
		Failures: []string{
			FailureNoNativeZ2CrossLane,
			FailureNoNativeLinearDomainExclusion,
			FailureNoNativeQuadraticFaceExclusion,
		},
	}
}

func buildSsplitTransportSubobject() RequiredSubobject {
	return RequiredSubobject{
		Index:           4,
		Name:            "S_split transport law",
		RequiredTheorem: SsplitTransportTheorem,
		Description:     "same boundary coordinate s=S_split feeds s into degree one and s^2 into degree two",
		CertifiedNative: false,
		Required:        true,
		CorrectShape:    true,
		Targets:         []string{"S_split -> Lambda^1 B_2 response coefficient s", "S_split^2 -> Lambda^2 B_2 response coefficient s^2"},
		Supports: []string{
			StatusSsplitTransportRequired,
			SupportSsplitTransportRequired,
			SupportSsplitFeedsDegreeShape,
		},
		Failures: []string{
			FailureNoNativeTransportS,
			FailureNoTypedSToLambda1,
			FailureNoTypedS2ToLambda2,
		},
	}
}

func buildDenominatorSubobject() RequiredSubobject {
	return RequiredSubobject{
		Index:           5,
		Name:            "Boundary-augmented chamber normalization",
		RequiredTheorem: DenominatorTypingTheorem,
		Description:     "denominator chambers typed as H_10 and H_72 but not yet activation laws",
		CertifiedNative: false,
		Required:        true,
		CorrectShape:    true,
		RankInputs:      []int{AmbientRank8, BoundaryRank, Lambda4Rank, BoundaryRank},
		Denominators:    []int{LinearDenom, QuadDenom},
		Targets:         []string{H10Chamber, H72Chamber},
		Supports: []string{
			StatusDenominatorTyping,
			SupportDenominatorChambersRequired,
			SupportDenominatorsTyped,
		},
		Failures: []string{
			FailureDenominatorNotActivation,
		},
	}
}

func buildFirewalls() Firewalls {
	return Firewalls{
		NativeZ2BoundaryAlphaFunctor: false,
		ReducedB2NativeFunctional:    false,
		DegreeToZ2FlagFunctor:        false,
		NativeZ2CrossLane:            false,
		NativeTransportS:             false,
		DenominatorActivationLaw:     false,
		AlphaNative:                  false,
		NativeR3:                     false,
		FullAFDescent:                false,
		GenerationCarrierMap:         false,
		FlavorOrientationMap:         false,
		IndividualYukawaValues:       false,
		OfficialLedgerUpdate:         false,
		NativeYukawaOperator:         false,
	}
}

func firewallsOK(f Firewalls) bool {
	return !f.NativeZ2BoundaryAlphaFunctor && !f.ReducedB2NativeFunctional && !f.DegreeToZ2FlagFunctor && !f.NativeZ2CrossLane && !f.NativeTransportS && !f.DenominatorActivationLaw && !f.AlphaNative && !f.NativeR3 && !f.FullAFDescent && !f.GenerationCarrierMap && !f.FlavorOrientationMap && !f.IndividualYukawaValues && !f.OfficialLedgerUpdate && !f.NativeYukawaOperator
}

func (f Firewalls) List() []string {
	out := []string{}
	if !f.NativeZ2BoundaryAlphaFunctor {
		out = append(out, FailureNoNativeZ2BoundaryAlphaFunctor)
	}
	if !f.ReducedB2NativeFunctional {
		out = append(out, FailureReducedB2NotNativeFunctional)
	}
	if !f.DegreeToZ2FlagFunctor {
		out = append(out, FailureNoNativeDegreeToZ2FlagFunctor)
	}
	if !f.NativeZ2CrossLane {
		out = append(out, FailureNoNativeZ2CrossLane)
	}
	if !f.NativeTransportS {
		out = append(out, FailureNoNativeTransportS)
	}
	if !f.DenominatorActivationLaw {
		out = append(out, FailureDenominatorNotActivation)
	}
	if !f.AlphaNative {
		out = append(out, FailureAlphaStillSealed)
	}
	if !f.NativeR3 {
		out = append(out, FailureNotNativeR3)
	}
	if !f.FullAFDescent {
		out = append(out, FailureFullAFDescentStillBlocked)
	}
	if !f.GenerationCarrierMap {
		out = append(out, FailureNoGenerationCarrierMap)
	}
	if !f.FlavorOrientationMap {
		out = append(out, FailureNoFlavorOrientationMap)
	}
	if !f.IndividualYukawaValues {
		out = append(out, FailureNoIndividualYukawaValues)
	}
	if !f.OfficialLedgerUpdate {
		out = append(out, FailureNoOfficialNEffUpdate)
	}
	if !f.NativeYukawaOperator {
		out = append(out, FailureNoNativeYukawaOperator)
	}
	return out
}

func (a Audit) FirewallsList() []string { return a.Firewalls.List() }

func allRequiredSubobjectsPresent(subs []RequiredSubobject) bool {
	if len(subs) != 5 {
		return false
	}
	seen := map[int]bool{}
	for _, s := range subs {
		if !s.Required || s.CertifiedNative {
			return false
		}
		seen[s.Index] = true
	}
	for i := 1; i <= 5; i++ {
		if !seen[i] {
			return false
		}
	}
	return true
}

func FormatInherited(i InheritedRail) string {
	return fmt.Sprintf("%s; rail=%s; no_loopback=%t/%t/%t; derives_alpha=%t; official_update=%t", i.Gate911ShortStatus, i.SelectedRail, !i.ReopensPhaseSign, !i.ReopensSocketOrder, !i.ReopensRepAlpha, i.DerivesAlpha, i.UpdatesOfficialLedger)
}

func FormatFormula(f SealedFormula) string {
	return fmt.Sprintf("%s; s=%.16g; alpha=%.16g; linear=%.16g; quadratic=%.16g; ranks=(%d,%d); denoms=(%d,%d); representative_free=%t; native=%t", f.Formula, f.S, f.Alpha, f.LinearContribution, f.QuadraticContribution, f.RankPair[0], f.RankPair[1], f.Denominators[0], f.Denominators[1], f.RepresentativeFree, f.Native)
}

func FormatSubobject(s RequiredSubobject) string {
	return fmt.Sprintf("subobject_%d=%s; theorem=%s; required=%t; certified_native=%t; shape_ok=%t; targets=%s; forbidden=%s; failures=%s", s.Index, s.Name, s.RequiredTheorem, s.Required, s.CertifiedNative, s.CorrectShape, strings.Join(s.Targets, " | "), strings.Join(s.ForbiddenTargets, " | "), strings.Join(s.Failures, ","))
}

func FormatDecomposition(d Decomposition) string {
	parts := []string{}
	for _, s := range d.Subobjects {
		parts = append(parts, fmt.Sprintf("%d:%s/native=%t", s.Index, s.Name, s.CertifiedNative))
	}
	return fmt.Sprintf("target=%s; required=%d; certified=%d; missing=%d; native_functor=%t; alpha_native=%t; r3_native=%t; subobjects=[%s]", d.NativeFunctorTarget, d.RequiredCount, d.CertifiedCount, d.MissingCount, d.NativeFunctor, d.AlphaNative, d.R3Native, strings.Join(parts, "; "))
}

func FormatFirewalls(f Firewalls) string {
	return strings.Join(f.List(), ",")
}

func Statuses() []string {
	return []string{
		StatusGate911Inherited,
		StatusNoLoopBack,
		StatusSealedFormulaReproduced,
		StatusDecompositionComplete,
		StatusReducedB2Shape,
		StatusDegreeSelectorRequired,
		StatusCrossLaneRequired,
		StatusSsplitTransportRequired,
		StatusDenominatorTyping,
		StatusNativeMissing,
		FinalTruth,
		Classification,
		ShortStatus,
		NextGate,
	}
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-15 }

func containsAll(haystack []string, needles []string) bool {
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
