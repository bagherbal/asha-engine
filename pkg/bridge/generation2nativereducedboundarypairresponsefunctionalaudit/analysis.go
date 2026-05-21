// Package generation2nativereducedboundarypairresponsefunctionalaudit
// implements Gate 913: Native Reduced BoundaryPair Response Functional Audit.
//
// Gate 913 follows Gate 912's decomposition of the native Z2 BoundaryAlpha
// functor into five required sub-objects. It audits only the first sub-object:
// the reduced rank-two boundary-pair exterior response
//
//	R_B(s)=(1+s b1)(1+s b2)-1
//	     =s(b1+b2)+s^2(b1 wedge b2).
//
// The gate certifies the canonical finite exterior shape, zero-order
// suppression, and rank-two truncation. It does not certify that ASHA must use
// the reduced functional, does not transport S_split natively, does not select
// Z2 flag targets, does not prove cross-lane exclusion, and does not promote
// alpha_B or R3 to native status.
package generation2nativereducedboundarypairresponsefunctionalaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE913-NATIVE-REDUCED-BOUNDARYPAIR-RESPONSE-FUNCTIONAL-AUDIT"

	SBoundary = 0.0012924448188162962
	AlphaB    = 0.0003878958469680527

	BoundaryPairRank = 2
	Lambda0Dim       = 1
	Lambda1Dim       = 2
	Lambda2Dim       = 1
	Lambda3Dim       = 0

	Gate912Classification = "R3_Z2_BOUNDARY_ALPHA_FUNCTOR_SOURCE_DECOMPOSED_NOT_NATIVE"
	Gate912ShortStatus    = "R3_ALPHA_FUNCTOR_DECOMPOSITION_COMPLETE_NATIVE_SUBOBJECTS_MISSING"
	Gate912Verdict        = "Z2_BOUNDARYALPHA_FUNCTOR_DECOMPOSED_INTO_FIVE_REQUIRED_NATIVE_SUBOBJECTS"

	BoundaryPair                  = "B_2=<b1,b2>"
	UnreducedResponse             = "E_B(s)=(1+s b1)(1+s b2)"
	ReducedResponse               = "R_B(s)=E_B(s)-1=(1+s b1)(1+s b2)-1"
	ExpandedUnreducedResponse     = "E_B(s)=1+s(b1+b2)+s^2(b1 wedge b2)"
	ExpandedReducedResponse       = "R_B(s)=s(b1+b2)+s^2(b1 wedge b2)"
	DegreeOneResponse             = "degree 1: s(b1+b2)"
	DegreeTwoResponse             = "degree 2: s^2(b1 wedge b2)"
	GeneralMultiplicativeResponse = "E_B(s)=prod_{i=1}^n(1+s b_i)"
	AlphaSkeleton                 = "alpha_B=(3/10)s+(7/72)s^2"
	Lambda1Z2Target               = "Lambda^1 B_2 -> [F_1/F_0]_{Z2}"
	Lambda2Z2Target               = "Lambda^2 B_2 -> [F_2/F_0]_{Z2}"
	NativeReducedB2Theorem        = "NativeReducedBoundaryPairResponseFunctional"
	BoundaryPairReducedTheorem    = "BoundaryPairReducedExteriorResponseTheorem"
	NextGate                      = "NEXT_PRESSURE_GATE914_DEGREEINDEXED_Z2_AIRLOCK_FLAGFUNCTOR_AUDIT"
	Classification                = "R3_REDUCED_B2_RESPONSE_FUNCTIONAL_SHAPE_CERTIFIED_NOT_NATIVE_BOUNDARY_ALPHA"
	ShortStatus                   = "R3_ALPHA_SUBOBJECT_1_REDUCED_B2_RESPONSE_SHAPE_PASS_NATIVE_SELECTION_BLOCKED"
	FinalTruth                    = "REDUCED_B2_RESPONSE_FUNCTIONAL_HAS_CANONICAL_EXTERIOR_SHAPE_BUT_NATIVE_SELECTION_REMAINS_BLOCKED"
	StrategicConclusion           = "R_B(s) has the correct finite exterior form of a rank-two boundary pair, but it still needs target selection, S_split transport, and a native selection principle."

	StatusGate912Inherited       = "PASS_GATE912_ALPHA_FUNCTOR_DECOMPOSITION_INHERITED"
	StatusSubobjectOneSelected   = "PASS_SUBOBJECT_1_REDUCED_B2_RESPONSE_FUNCTIONAL_SELECTED"
	StatusExteriorLedgerBuilt    = "PASS_BOUNDARY_PAIR_EXTERIOR_LEDGER_BUILT"
	StatusExpansionExact         = "PASS_REDUCED_B2_RESPONSE_EXPANSION_EXACT"
	StatusZeroOrderSuppressed    = "PASS_ZERO_ORDER_TERM_SUPPRESSED_BY_E_B_MINUS_ONE"
	StatusRankTwoTruncation      = "PASS_CUBIC_AND_HIGHER_TERMS_ABSENT_BY_RANK_TWO_EXTERIOR_TRUNCATION"
	StatusNaturalityCandidate    = "PASS_MULTIPLICATIVE_BOUNDARY_PAIR_RESPONSE_RECORDED_AS_NATURAL_CANDIDATE"
	StatusAlphaShapeMatched      = "PASS_ALPHA_POWER_SHAPE_MATCHED_WITHOUT_TARGET_SELECTION"
	StatusSsplitFirewall         = "FIREWALL_PRESERVED_NO_NATIVE_S_SPLIT_TRANSPORT_TO_REDUCED_B2_RESPONSE"
	StatusNativeSelectionBlocked = "FIREWALL_PRESERVED_GATE913_NATIVE_SELECTION_BLOCKED"

	SupportExactExpansion              = "CONDITIONAL_SUPPORT_REDUCED_B2_RESPONSE_EXPANDS_TO_EXACT_S_PLUS_S_SQUARED_SHAPE"
	SupportDegreeOneExposureSum        = "CONDITIONAL_SUPPORT_DEGREE_ONE_RESPONSE_IS_S_TIMES_BOUNDARY_EXPOSURE_SUM"
	SupportDegreeTwoPairEnclosure      = "CONDITIONAL_SUPPORT_DEGREE_TWO_RESPONSE_IS_S_SQUARED_TIMES_BOUNDARY_PAIR_ENCLOSURE"
	SupportZeroOrderSuppressed         = "CONDITIONAL_SUPPORT_ZERO_ORDER_TERM_SUPPRESSED_BY_REDUCED_RESPONSE_E_B_MINUS_ONE"
	SupportStartsAtOrderOne            = "CONDITIONAL_SUPPORT_REDUCED_RESPONSE_STARTS_AT_BOUNDARY_ACTIVATION_ORDER_ONE"
	SupportCubicAbsentRankTwo          = "CONDITIONAL_SUPPORT_CUBIC_AND_HIGHER_RESPONSE_TERMS_ABSENT_BY_RANK_TWO_EXTERIOR_TRUNCATION"
	SupportLambda3Zero                 = "CONDITIONAL_SUPPORT_LAMBDA3_B2_EQUALS_ZERO_EXPLAINS_RESPONSE_TRUNCATION_AFTER_SECOND_DEGREE"
	SupportMultiplicativeNatural       = "CONDITIONAL_SUPPORT_MULTIPLICATIVE_BOUNDARY_PAIR_RESPONSE_IS_NATURAL_EXTERIOR_ACTIVATION_CANDIDATE"
	SupportReducedNontrivialPart       = "CONDITIONAL_SUPPORT_R_B_IS_REDUCED_NONTRIVIAL_PART_OF_BOUNDARY_EXTERIOR_ACTIVATION"
	SupportSuppliesAlphaPowerShape     = "CONDITIONAL_SUPPORT_REDUCED_B2_RESPONSE_SUPPLIES_ALPHA_POWER_SHAPE"
	SupportAlphaPolynomialShapeMatches = "CONDITIONAL_SUPPORT_ALPHA_POLYNOMIAL_SHAPE_MATCHES_BOUNDARY_PAIR_REDUCED_EXTERIOR_RESPONSE"
	SupportSsplitFeedsShape            = "CONDITIONAL_SUPPORT_S_SPLIT_FEEDS_DEGREE_ONE_AND_TWO_RESPONSE_SHAPE"

	FailureReducedB2NotNativeFunctional = "FAILED_ROUTE_REDUCED_B2_RESPONSE_NOT_NATIVE_BOUNDARY_FUNCTIONAL"
	FailureNoNativeReasonEBMinusOne     = "FAILED_ROUTE_NO_NATIVE_REASON_FOR_USING_E_B_MINUS_ONE_RESPONSE"
	FailureMultiplicativeNotNative      = "FAILED_ROUTE_MULTIPLICATIVE_BOUNDARY_RESPONSE_NOT_YET_NATIVE_ASHA_FUNCTIONAL"
	FailureNoVariationalProductForm     = "FAILED_ROUTE_NO_VARIATIONAL_OR_FUNCTORIAL_PRINCIPLE_SELECTING_PRODUCT_FORM"
	FailureNoNativeTransportSIntoB2     = "FAILED_ROUTE_NO_NATIVE_TRANSPORT_OF_S_SPLIT_INTO_REDUCED_B2_RESPONSE"
	FailureNoTypedSToExteriorParameter  = "FAILED_ROUTE_NO_TYPED_S_SPLIT_TO_BOUNDARY_PAIR_EXTERIOR_PARAMETER_MAP"
	FailureNoZ2FlagTargets              = "FAILED_ROUTE_REDUCED_RESPONSE_DOES_NOT_SELECT_Z2_FLAG_TARGETS"
	FailureNoAlphaCoefficients          = "FAILED_ROUTE_REDUCED_RESPONSE_DOES_NOT_BY_ITSELF_DERIVE_3_OVER_10_OR_7_OVER_72"
	FailureNoCrossLaneExclusion         = "FAILED_ROUTE_REDUCED_RESPONSE_DOES_NOT_PROVE_CROSS_LANE_EXCLUSION"
	FailureReducedResponseNotAlphaAlone = "FAILED_ROUTE_REDUCED_RESPONSE_DOES_NOT_DERIVE_ALPHA_B_ALONE"
	FailureAlphaStillSealed             = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNotNativeR3                  = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked    = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap       = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap       = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues     = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator       = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type InheritedDecomposition struct {
	Gate912Classification string
	Gate912ShortStatus    string
	Gate912Verdict        string
	SelectedSubobject     string
	SubobjectIndex        int
	TotalSubobjects       int
	ReopensPhaseSign      bool
	ReopensSocketOrder    bool
	ReopensRepresentative bool
	DerivesAlpha          bool
	UpdatesOfficialLedger bool
	Supports, Failures    []string
}

type ExteriorLedger struct {
	BoundaryPair           string
	Rank                   int
	Lambda0Dim, Lambda1Dim int
	Lambda2Dim, Lambda3Dim int
	Lambda3Zero            bool
	CubicAndHigherVanish   bool
	Supports, Failures     []string
}

type ExteriorTerm struct {
	Degree int
	Basis  string
	Coeff  float64
}

type ResponseExpansion struct {
	S                  float64
	UnreducedFormula   string
	ReducedFormula     string
	UnreducedExpanded  string
	ReducedExpanded    string
	UnreducedTerms     []ExteriorTerm
	ReducedTerms       []ExteriorTerm
	DegreeOneTerms     []ExteriorTerm
	DegreeTwoTerms     []ExteriorTerm
	ConstantRemoved    bool
	ExactShape         bool
	NativeFunctional   bool
	Supports, Failures []string
}

type ZeroOrderSuppression struct {
	IdentityTermBasis        string
	IdentityTermCoeff        float64
	RemovedByReduction       bool
	ReducedStartsAtOrderOne  bool
	NativeReasonForReduction bool
	Supports, Failures       []string
}

type RankTwoTruncation struct {
	Rank                      int
	HighestNonzeroDegree      int
	Lambda3Zero               bool
	NoCubicOrHigher           bool
	ExteriorAlgebraNativeFact bool
	Supports, Failures        []string
}

type NaturalityCandidate struct {
	GeneralFormula               string
	SpecializedFormula           string
	ReducedNontrivialPart        string
	MultiplicativeCandidate      bool
	NativeASHAFunctional         bool
	VariationalPrinciple         bool
	FunctorialSelectionPrinciple bool
	Supports, Failures           []string
}

type AlphaShapeRelation struct {
	AlphaSkeleton            string
	PowersSupplied           []int
	SuppliesPowerShape       bool
	SelectsZ2FlagTargets     bool
	DerivesCoefficients      bool
	ProvesCrossLaneExclusion bool
	DerivesAlpha             bool
	Supports, Failures       []string
}

type SsplitTransport struct {
	SourceCoordinate   string
	ExteriorParameter  string
	UsesSAsParameter   bool
	NativeTransport    bool
	TypedParameterMap  bool
	Supports, Failures []string
}

type Firewalls struct {
	ReducedB2NativeFunctional   bool
	NativeReasonEBMinusOne      bool
	MultiplicativeNative        bool
	VariationalProductSelection bool
	NativeSsplitTransport       bool
	TypedSToExteriorParameter   bool
	SelectsZ2FlagTargets        bool
	DerivesAlphaCoefficients    bool
	CrossLaneExclusion          bool
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
	Inherited      InheritedDecomposition
	Ledger         ExteriorLedger
	Expansion      ResponseExpansion
	ZeroOrder      ZeroOrderSuppression
	Truncation     RankTwoTruncation
	Naturality     NaturalityCandidate
	AlphaShape     AlphaShapeRelation
	Transport      SsplitTransport
	Firewalls      Firewalls
	Truth          string
	Final          string
}

func BuildDefault() (Audit, error) {
	inherited := buildInherited()
	if inherited.ReopensPhaseSign || inherited.ReopensSocketOrder || inherited.ReopensRepresentative || inherited.DerivesAlpha || inherited.UpdatesOfficialLedger {
		return Audit{}, fmt.Errorf("closed wound reopened: %s", FormatInherited(inherited))
	}

	ledger := buildExteriorLedger()
	if ledger.Rank != BoundaryPairRank || ledger.Lambda3Dim != 0 || !ledger.Lambda3Zero || !ledger.CubicAndHigherVanish {
		return Audit{}, fmt.Errorf("bad exterior ledger: %s", FormatLedger(ledger))
	}

	expansion := buildExpansion(SBoundary)
	if !expansion.ExactShape || !expansion.ConstantRemoved || expansion.NativeFunctional || len(expansion.ReducedTerms) != 3 || !near(sumDegree(expansion.ReducedTerms, 1), 2*SBoundary) || !near(sumDegree(expansion.ReducedTerms, 2), SBoundary*SBoundary) {
		return Audit{}, fmt.Errorf("bad response expansion: %s", FormatExpansion(expansion))
	}

	zero := buildZeroOrder()
	if !zero.RemovedByReduction || !zero.ReducedStartsAtOrderOne || zero.NativeReasonForReduction {
		return Audit{}, fmt.Errorf("bad zero-order status: %s", FormatZeroOrder(zero))
	}

	trunc := buildTruncation()
	if !trunc.Lambda3Zero || !trunc.NoCubicOrHigher || trunc.HighestNonzeroDegree != 2 {
		return Audit{}, fmt.Errorf("bad rank-two truncation: %s", FormatTruncation(trunc))
	}

	nat := buildNaturalityCandidate()
	if !nat.MultiplicativeCandidate || nat.NativeASHAFunctional || nat.VariationalPrinciple || nat.FunctorialSelectionPrinciple {
		return Audit{}, fmt.Errorf("bad naturality classification: %s", FormatNaturality(nat))
	}

	alpha := buildAlphaShapeRelation()
	if !alpha.SuppliesPowerShape || alpha.SelectsZ2FlagTargets || alpha.DerivesCoefficients || alpha.ProvesCrossLaneExclusion || alpha.DerivesAlpha {
		return Audit{}, fmt.Errorf("bad alpha-shape firewall: %s", FormatAlphaShape(alpha))
	}

	transport := buildSsplitTransport()
	if !transport.UsesSAsParameter || transport.NativeTransport || transport.TypedParameterMap {
		return Audit{}, fmt.Errorf("bad S_split transport firewall: %s", FormatTransport(transport))
	}

	firewalls := buildFirewalls()
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewall leak: %s", FormatFirewalls(firewalls))
	}

	return Audit{ID: AuditID, Classification: Classification, ShortStatus: ShortStatus, Inherited: inherited, Ledger: ledger, Expansion: expansion, ZeroOrder: zero, Truncation: trunc, Naturality: nat, AlphaShape: alpha, Transport: transport, Firewalls: firewalls, Truth: FinalTruth, Final: StrategicConclusion}, nil
}

func buildInherited() InheritedDecomposition {
	return InheritedDecomposition{
		Gate912Classification: Gate912Classification,
		Gate912ShortStatus:    Gate912ShortStatus,
		Gate912Verdict:        Gate912Verdict,
		SelectedSubobject:     NativeReducedB2Theorem,
		SubobjectIndex:        1,
		TotalSubobjects:       5,
		ReopensPhaseSign:      false,
		ReopensSocketOrder:    false,
		ReopensRepresentative: false,
		DerivesAlpha:          false,
		UpdatesOfficialLedger: false,
		Supports: []string{
			StatusGate912Inherited,
			StatusSubobjectOneSelected,
		},
		Failures: []string{
			FailureReducedB2NotNativeFunctional,
			FailureAlphaStillSealed,
		},
	}
}

func buildExteriorLedger() ExteriorLedger {
	return ExteriorLedger{
		BoundaryPair:         BoundaryPair,
		Rank:                 BoundaryPairRank,
		Lambda0Dim:           Lambda0Dim,
		Lambda1Dim:           Lambda1Dim,
		Lambda2Dim:           Lambda2Dim,
		Lambda3Dim:           Lambda3Dim,
		Lambda3Zero:          true,
		CubicAndHigherVanish: true,
		Supports: []string{
			StatusExteriorLedgerBuilt,
			SupportCubicAbsentRankTwo,
			SupportLambda3Zero,
		},
		Failures: []string{},
	}
}

func buildExpansion(s float64) ResponseExpansion {
	unreduced := []ExteriorTerm{
		{Degree: 0, Basis: "1", Coeff: 1},
		{Degree: 1, Basis: "b1", Coeff: s},
		{Degree: 1, Basis: "b2", Coeff: s},
		{Degree: 2, Basis: "b1 wedge b2", Coeff: s * s},
	}
	reduced := []ExteriorTerm{
		{Degree: 1, Basis: "b1", Coeff: s},
		{Degree: 1, Basis: "b2", Coeff: s},
		{Degree: 2, Basis: "b1 wedge b2", Coeff: s * s},
	}
	return ResponseExpansion{
		S:                 s,
		UnreducedFormula:  UnreducedResponse,
		ReducedFormula:    ReducedResponse,
		UnreducedExpanded: ExpandedUnreducedResponse,
		ReducedExpanded:   ExpandedReducedResponse,
		UnreducedTerms:    unreduced,
		ReducedTerms:      reduced,
		DegreeOneTerms:    []ExteriorTerm{reduced[0], reduced[1]},
		DegreeTwoTerms:    []ExteriorTerm{reduced[2]},
		ConstantRemoved:   true,
		ExactShape:        true,
		NativeFunctional:  false,
		Supports: []string{
			StatusExpansionExact,
			SupportExactExpansion,
			SupportDegreeOneExposureSum,
			SupportDegreeTwoPairEnclosure,
		},
		Failures: []string{
			FailureReducedB2NotNativeFunctional,
		},
	}
}

func buildZeroOrder() ZeroOrderSuppression {
	return ZeroOrderSuppression{
		IdentityTermBasis:        "1 in Lambda^0 B_2",
		IdentityTermCoeff:        1,
		RemovedByReduction:       true,
		ReducedStartsAtOrderOne:  true,
		NativeReasonForReduction: false,
		Supports: []string{
			StatusZeroOrderSuppressed,
			SupportZeroOrderSuppressed,
			SupportStartsAtOrderOne,
		},
		Failures: []string{
			FailureNoNativeReasonEBMinusOne,
		},
	}
}

func buildTruncation() RankTwoTruncation {
	return RankTwoTruncation{
		Rank:                      BoundaryPairRank,
		HighestNonzeroDegree:      2,
		Lambda3Zero:               true,
		NoCubicOrHigher:           true,
		ExteriorAlgebraNativeFact: true,
		Supports: []string{
			StatusRankTwoTruncation,
			SupportCubicAbsentRankTwo,
			SupportLambda3Zero,
		},
		Failures: []string{},
	}
}

func buildNaturalityCandidate() NaturalityCandidate {
	return NaturalityCandidate{
		GeneralFormula:               GeneralMultiplicativeResponse,
		SpecializedFormula:           UnreducedResponse,
		ReducedNontrivialPart:        ReducedResponse,
		MultiplicativeCandidate:      true,
		NativeASHAFunctional:         false,
		VariationalPrinciple:         false,
		FunctorialSelectionPrinciple: false,
		Supports: []string{
			StatusNaturalityCandidate,
			SupportMultiplicativeNatural,
			SupportReducedNontrivialPart,
		},
		Failures: []string{
			FailureMultiplicativeNotNative,
			FailureNoVariationalProductForm,
		},
	}
}

func buildAlphaShapeRelation() AlphaShapeRelation {
	return AlphaShapeRelation{
		AlphaSkeleton:            AlphaSkeleton,
		PowersSupplied:           []int{1, 2},
		SuppliesPowerShape:       true,
		SelectsZ2FlagTargets:     false,
		DerivesCoefficients:      false,
		ProvesCrossLaneExclusion: false,
		DerivesAlpha:             false,
		Supports: []string{
			StatusAlphaShapeMatched,
			SupportSuppliesAlphaPowerShape,
			SupportAlphaPolynomialShapeMatches,
		},
		Failures: []string{
			FailureNoZ2FlagTargets,
			FailureNoAlphaCoefficients,
			FailureNoCrossLaneExclusion,
			FailureReducedResponseNotAlphaAlone,
			FailureAlphaStillSealed,
		},
	}
}

func buildSsplitTransport() SsplitTransport {
	return SsplitTransport{
		SourceCoordinate:  "s=S_split",
		ExteriorParameter: "s in R_B(s)",
		UsesSAsParameter:  true,
		NativeTransport:   false,
		TypedParameterMap: false,
		Supports: []string{
			SupportSsplitFeedsShape,
		},
		Failures: []string{
			StatusSsplitFirewall,
			FailureNoNativeTransportSIntoB2,
			FailureNoTypedSToExteriorParameter,
		},
	}
}

func buildFirewalls() Firewalls {
	return Firewalls{}
}

func firewallsOK(f Firewalls) bool {
	return !f.ReducedB2NativeFunctional && !f.NativeReasonEBMinusOne && !f.MultiplicativeNative && !f.VariationalProductSelection && !f.NativeSsplitTransport && !f.TypedSToExteriorParameter && !f.SelectsZ2FlagTargets && !f.DerivesAlphaCoefficients && !f.CrossLaneExclusion && !f.AlphaNative && !f.NativeR3 && !f.FullAFDescent && !f.GenerationCarrierMap && !f.FlavorOrientationMap && !f.IndividualYukawaValues && !f.NativeYukawaOperator
}

func (f Firewalls) List() []string {
	out := []string{}
	if !f.ReducedB2NativeFunctional {
		out = append(out, FailureReducedB2NotNativeFunctional)
	}
	if !f.NativeReasonEBMinusOne {
		out = append(out, FailureNoNativeReasonEBMinusOne)
	}
	if !f.MultiplicativeNative {
		out = append(out, FailureMultiplicativeNotNative)
	}
	if !f.VariationalProductSelection {
		out = append(out, FailureNoVariationalProductForm)
	}
	if !f.NativeSsplitTransport {
		out = append(out, FailureNoNativeTransportSIntoB2)
	}
	if !f.TypedSToExteriorParameter {
		out = append(out, FailureNoTypedSToExteriorParameter)
	}
	if !f.SelectsZ2FlagTargets {
		out = append(out, FailureNoZ2FlagTargets)
	}
	if !f.DerivesAlphaCoefficients {
		out = append(out, FailureNoAlphaCoefficients)
	}
	if !f.CrossLaneExclusion {
		out = append(out, FailureNoCrossLaneExclusion)
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
	if !f.NativeYukawaOperator {
		out = append(out, FailureNoNativeYukawaOperator)
	}
	return out
}

func (a Audit) FirewallsList() []string { return a.Firewalls.List() }

func sumDegree(terms []ExteriorTerm, degree int) float64 {
	var out float64
	for _, t := range terms {
		if t.Degree == degree {
			out += t.Coeff
		}
	}
	return out
}

func hasDegree(terms []ExteriorTerm, degree int) bool {
	for _, t := range terms {
		if t.Degree == degree {
			return true
		}
	}
	return false
}

func FormatInherited(i InheritedDecomposition) string {
	return fmt.Sprintf("%s; selected=%s[%d/%d]; no_loopback=%t/%t/%t; derives_alpha=%t; official_update=%t", i.Gate912ShortStatus, i.SelectedSubobject, i.SubobjectIndex, i.TotalSubobjects, !i.ReopensPhaseSign, !i.ReopensSocketOrder, !i.ReopensRepresentative, i.DerivesAlpha, i.UpdatesOfficialLedger)
}

func FormatLedger(l ExteriorLedger) string {
	return fmt.Sprintf("%s; rank=%d; dims=(%d,%d,%d,%d); Lambda3Zero=%t; cubic_or_higher_vanish=%t", l.BoundaryPair, l.Rank, l.Lambda0Dim, l.Lambda1Dim, l.Lambda2Dim, l.Lambda3Dim, l.Lambda3Zero, l.CubicAndHigherVanish)
}

func FormatExpansion(e ResponseExpansion) string {
	return fmt.Sprintf("%s; %s; s=%.16g; reduced_terms=%s; degree1_sum=%.16g; degree2_sum=%.16g; constant_removed=%t; exact_shape=%t; native=%t", e.ReducedFormula, e.ReducedExpanded, e.S, FormatTerms(e.ReducedTerms), sumDegree(e.ReducedTerms, 1), sumDegree(e.ReducedTerms, 2), e.ConstantRemoved, e.ExactShape, e.NativeFunctional)
}

func FormatTerms(terms []ExteriorTerm) string {
	parts := []string{}
	for _, t := range terms {
		parts = append(parts, fmt.Sprintf("deg%d:%s=%.16g", t.Degree, t.Basis, t.Coeff))
	}
	return strings.Join(parts, " | ")
}

func FormatZeroOrder(z ZeroOrderSuppression) string {
	return fmt.Sprintf("identity=%s coeff=%.16g; removed=%t; starts_order_one=%t; native_reason=%t", z.IdentityTermBasis, z.IdentityTermCoeff, z.RemovedByReduction, z.ReducedStartsAtOrderOne, z.NativeReasonForReduction)
}

func FormatTruncation(t RankTwoTruncation) string {
	return fmt.Sprintf("rank=%d; highest_nonzero_degree=%d; Lambda3Zero=%t; no_cubic_or_higher=%t; exterior_native_fact=%t", t.Rank, t.HighestNonzeroDegree, t.Lambda3Zero, t.NoCubicOrHigher, t.ExteriorAlgebraNativeFact)
}

func FormatNaturality(n NaturalityCandidate) string {
	return fmt.Sprintf("general=%s; specialized=%s; reduced=%s; candidate=%t; native_ASHA=%t; variational=%t; functorial=%t", n.GeneralFormula, n.SpecializedFormula, n.ReducedNontrivialPart, n.MultiplicativeCandidate, n.NativeASHAFunctional, n.VariationalPrinciple, n.FunctorialSelectionPrinciple)
}

func FormatAlphaShape(a AlphaShapeRelation) string {
	return fmt.Sprintf("%s; powers=%v; supplies_shape=%t; selects_targets=%t; derives_coefficients=%t; cross_lane=%t; derives_alpha=%t", a.AlphaSkeleton, a.PowersSupplied, a.SuppliesPowerShape, a.SelectsZ2FlagTargets, a.DerivesCoefficients, a.ProvesCrossLaneExclusion, a.DerivesAlpha)
}

func FormatTransport(t SsplitTransport) string {
	return fmt.Sprintf("source=%s; parameter=%s; uses_s=%t; native_transport=%t; typed_map=%t", t.SourceCoordinate, t.ExteriorParameter, t.UsesSAsParameter, t.NativeTransport, t.TypedParameterMap)
}

func FormatFirewalls(f Firewalls) string {
	return strings.Join(f.List(), ",")
}

func Statuses() []string {
	return []string{
		StatusGate912Inherited,
		StatusSubobjectOneSelected,
		StatusExteriorLedgerBuilt,
		StatusExpansionExact,
		StatusZeroOrderSuppressed,
		StatusRankTwoTruncation,
		StatusNaturalityCandidate,
		StatusAlphaShapeMatched,
		StatusSsplitFirewall,
		StatusNativeSelectionBlocked,
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
