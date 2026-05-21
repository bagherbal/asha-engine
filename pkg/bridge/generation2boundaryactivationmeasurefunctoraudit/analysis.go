// Package generation2boundaryactivationmeasurefunctoraudit implements
// Gate 920: BoundaryActivationMeasure Functor Audit.
//
// Gate 920 follows Gate 919's collapse route. It defines the formal bridge
// measure mu_B on the reduced rank-two boundary-pair response and audits
// whether that measure lawfully reassembles the five alpha-side subobjects into
// alpha_B^Z2. The result is intentionally limited: mu_B is a bridge-lawful
// candidate and organization principle, not a certified native ASHA measure,
// not native alpha, and not native R3.
package generation2boundaryactivationmeasurefunctoraudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE920-BOUNDARYACTIVATIONMEASURE-FUNCTOR-AUDIT"

	Gate919ShortStatus = "R3_ALPHA_GAPS_COLLAPSE_TO_BOUNDARY_MEASURE_OBSTRUCTION"

	SBoundary = 0.0012924448188162962
	AlphaB    = 0.0003878958469680527

	RankI1  = 3
	RankI2  = 7
	RankH10 = 10
	RankH72 = 72

	AlphaLinear = 0.00038773344564488885
	AlphaQuad   = 0.0000001624013231638281

	ReducedResponse             = "R_B(s)=(1+s b1)(1+s b2)-1=s(b1+b2)+s^2(b1 wedge b2)"
	BoundaryActivationMeasure   = "BoundaryActivationMeasure"
	BoundaryResponseFunctor     = "Z2BoundaryResponseFunctor"
	MeasureFormula              = "mu_B(R_B(S_split))=sum_{k=1}^2 rank(I_B^Z2(k))/rank(H_k)*S_split^k"
	BranchMeasureFormula        = "mu_B(R_B(S_split))=(3/10)S_split+(7/72)S_split^2"
	BoundaryAlphaMeasureFormula = "alpha_B^Z2=mu_B(R_B(S_split))"

	Classification      = "R3_BOUNDARY_ACTIVATION_MEASURE_FUNCTOR_CANDIDATE_NOT_NATIVE"
	ShortStatus         = "R3_ALPHA_BOUNDARY_MEASURE_CANDIDATE_NATIVE_MEASURE_MISSING"
	FinalTruth          = "BOUNDARY_ACTIVATION_MEASURE_FORMALLY_REASSEMBLES_ALPHA_B_Z2_BUT_NATIVE_MEASURE_THEOREM_MISSING"
	StrategicConclusion = "Gate 920 compresses the alpha branch from five shape-level subobjects to one formal BoundaryActivationMeasure candidate. The key native question is now why this measure is forced or unique; mu_B remains bridge-lawful notation, not a native ASHA theorem."
	NextGate            = "NEXT_PRESSURE_GATE921_BOUNDARYACTIVATIONMEASURE_NATURALITY_AND_UNIQUENESS_AUDIT"

	StatusInheritedGate919      = "PASS_GATE919_BOUNDARY_MEASURE_COLLAPSE_ROUTE_INHERITED"
	StatusDomainReducedResponse = "PASS_MU_B_DOMAIN_REDUCED_BOUNDARY_PAIR_RESPONSE"
	StatusDegreeExtraction      = "PASS_MU_B_EXTRACTS_ACTIVE_EXTERIOR_DEGREES"
	StatusSelectorIntegrated    = "PASS_MU_B_INTEGRATES_Z2_AIRLOCK_SELECTOR"
	StatusChamberNormalization  = "PASS_MU_B_INTEGRATES_BOUNDARY_AUGMENTED_CHAMBERS"
	StatusCrossLaneAbsorbed     = "PASS_CROSS_LANE_EXCLUSION_ABSORBED_IF_SELECTOR_FUNCTIONAL"
	StatusAlphaReconstructed    = "PASS_MU_B_RECONSTRUCTS_ALPHA_B_Z2"
	StatusNativeMeasureMissing  = "FIREWALL_PRESERVED_NATIVE_MEASURE_THEOREM_MISSING"
	StatusNoNativePromotion     = "FIREWALL_PRESERVED_ALPHA_AND_R3_NOT_NATIVE"

	SupportMuBDomainReducedResponse          = "CONDITIONAL_SUPPORT_MU_B_DOMAIN_IS_REDUCED_BOUNDARY_PAIR_RESPONSE"
	SupportMuBActsOnActiveDegrees            = "CONDITIONAL_SUPPORT_MU_B_ACTS_ON_ACTIVE_NONZERO_EXTERIOR_DEGREES"
	SupportMuBIgnoresLambda0Basepoint        = "CONDITIONAL_SUPPORT_MU_B_IGNORES_LAMBDA0_BASEPOINT_AFTER_REDUCTION"
	SupportMuBExtractsByExteriorDegree       = "CONDITIONAL_SUPPORT_MU_B_EXTRACTS_RESPONSE_BY_EXTERIOR_DEGREE"
	SupportDegreeKCarriesSPowerK             = "CONDITIONAL_SUPPORT_DEGREE_K_COMPONENT_CARRIES_S_SPLIT_POWER_K"
	SupportSPowerFollowsExteriorResponse     = "CONDITIONAL_SUPPORT_S_SPLIT_POWER_STRUCTURE_FOLLOWS_FROM_REDUCED_EXTERIOR_RESPONSE"
	SupportMuBIntegratesSelector             = "CONDITIONAL_SUPPORT_MU_B_INTEGRATES_DEGREE_INDEXED_Z2_AIRLOCK_SELECTOR"
	SupportMuBRecoversRankPair               = "CONDITIONAL_SUPPORT_MU_B_RECOVERS_TARGET_RANK_PAIR_3_7"
	SupportSelectorRepresentativeIndependent = "CONDITIONAL_SUPPORT_SELECTOR_INPUTS_ARE_REPRESENTATIVE_INDEPENDENT"
	SupportMuBIntegratesChamberNormalization = "CONDITIONAL_SUPPORT_MU_B_INTEGRATES_BOUNDARY_AUGMENTED_CHAMBER_NORMALIZATION"
	SupportMuBRecoversCoefficients           = "CONDITIONAL_SUPPORT_MU_B_RECOVERS_COEFFICIENTS_3_OVER_10_AND_7_OVER_72"
	SupportLaneNormalizationExplicit         = "CONDITIONAL_SUPPORT_LANE_NORMALIZATION_IS_EXPLICIT_IN_MEASURE_FORM"
	SupportMuBExcludesCrossLanesIfFunctional = "CONDITIONAL_SUPPORT_MU_B_EXCLUDES_CROSS_LANES_IF_SELECTOR_IS_FUNCTIONAL"
	SupportCrossLaneAbsorbedInIndexing       = "CONDITIONAL_SUPPORT_CROSS_LANE_EXCLUSION_IS_ABSORBED_IN_MEASURE_INDEXING"
	SupportMuBReconstructsAlpha              = "CONDITIONAL_SUPPORT_MU_B_RECONSTRUCTS_ALPHA_B_Z2"
	SupportMeasureReassemblesFive            = "CONDITIONAL_SUPPORT_BOUNDARY_ACTIVATION_MEASURE_REASSEMBLES_ALL_FIVE_ALPHA_SUBOBJECTS"
	SupportAlphaAsMeasure                    = "CONDITIONAL_SUPPORT_ALPHA_B_CAN_BE_EXPRESSED_AS_MEASURE_OF_REDUCED_BOUNDARY_RESPONSE"

	FailureNoNativeBoundaryActivationMeasure = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_ACTIVATION_MEASURE_CERTIFIED"
	FailureMuBFormalNotNative                = "FAILED_ROUTE_MU_B_IS_FORMAL_BRIDGE_MEASURE_NOT_NATIVE_THEOREM"
	FailureNoNativeMeasureUniqueness         = "FAILED_ROUTE_NO_NATIVE_MEASURE_UNIQUENESS_THEOREM"
	FailureNoNativeMuBDomain                 = "FAILED_ROUTE_NO_NATIVE_PROOF_THAT_MU_B_MUST_ACT_ON_REDUCED_RESPONSE"
	FailureReducedResponseBridgeSelected     = "FAILED_ROUTE_REDUCED_RESPONSE_REMAINS_BRIDGE_SELECTED_NOT_NATIVE"
	FailureNoNativeDegreeExtraction          = "FAILED_ROUTE_NO_NATIVE_DEGREE_EXTRACTION_FUNCTIONAL_CERTIFIED"
	FailureNoNativeSSplitTransportMap        = "FAILED_ROUTE_NO_NATIVE_S_SPLIT_TRANSPORT_MAP"
	FailureNoNativeDegreeToZ2FlagFunctor     = "FAILED_ROUTE_NO_NATIVE_DEGREE_TO_Z2_FLAG_CLASS_FUNCTOR"
	FailureNoNativeUniqueSelector            = "FAILED_ROUTE_NO_NATIVE_PROOF_THAT_I_B_Z2_IS_UNIQUE_SELECTOR"
	FailureNoNativeChamberNormalization      = "FAILED_ROUTE_NO_NATIVE_RESPONSE_CHAMBER_NORMALIZATION_THEOREM"
	FailureNoNativeH1H2Reason                = "FAILED_ROUTE_NO_NATIVE_REASON_H1_EQUALS_H10_AND_H2_EQUALS_H72"
	FailureNoNativeZ2CrossLaneExclusion      = "FAILED_ROUTE_NO_NATIVE_Z2_CROSS_LANE_EXCLUSION_THEOREM"
	FailureSelectorFunctionhoodNotNative     = "FAILED_ROUTE_FUNCTIONHOOD_OF_I_B_Z2_NOT_NATIVE_CERTIFIED"
	FailureAlphaByMuBNotNative               = "FAILED_ROUTE_ALPHA_RECONSTRUCTION_BY_MU_B_NOT_NATIVE_ALPHA_THEOREM"
	FailureAlphaBridgeCandidateNotNative     = "FAILED_ROUTE_ALPHA_B_REMAINS_BRIDGE_CANDIDATE_NOT_NATIVE"
	FailureNotNativeR3                       = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked         = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap            = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap            = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues          = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator            = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type MeasureDomainAudit struct {
	InheritedStatus  string
	Domain           string
	NonzeroDegrees   []int
	IncludesLambda0  bool
	ReducedBasepoint bool
	NativeTheorem    bool
	Supports         []string
	Failures         []string
}

type DegreeExtractionAudit struct {
	DegreeCoefficients  map[int]float64
	DegreePowers        map[int]int
	SeparateS2Transport bool
	ExteriorGeneratedS2 bool
	NativeTheorem       bool
	Supports            []string
	Failures            []string
}

type SelectorIntegrationAudit struct {
	Targets                   map[int]string
	Ranks                     map[int]int
	RepresentativeIndependent bool
	UniqueNativeSelector      bool
	Supports                  []string
	Failures                  []string
}

type ChamberNormalizationAudit struct {
	Chambers            map[int]string
	Ranks               map[int]int
	Weights             map[int]float64
	ExplicitLaneWeights bool
	NativeTheorem       bool
	Supports            []string
	Failures            []string
}

type CrossLaneMeasureAudit struct {
	CorrectTargets       map[int]string
	FalseTargets         map[int]string
	ExcludedIfFunctional bool
	AbsorbedInIndexing   bool
	FunctionhoodNative   bool
	Supports             []string
	Failures             []string
}

type AlphaReconstructionAudit struct {
	Formula               string
	BoundaryFormula       string
	S                     float64
	RankI1                int
	RankI2                int
	RankH1                int
	RankH2                int
	LinearContribution    float64
	QuadraticContribution float64
	Alpha                 float64
	ReassemblesFive       bool
	NativeAlpha           bool
	Supports              []string
	Failures              []string
}

type NativeStatusAudit struct {
	BridgeMeasureCandidate bool
	NativeMeasure          bool
	NativeAlpha            bool
	NativeR3               bool
	MissingNativeTheorems  []string
	Supports               []string
	Failures               []string
}

type FirewallLedger struct {
	BoundaryActivationMeasureNative bool
	MuBNative                       bool
	MeasureUniquenessNative         bool
	DomainNative                    bool
	DegreeExtractionNative          bool
	SSplitTransportNative           bool
	SelectorNative                  bool
	ChamberNormalizationNative      bool
	CrossLaneNative                 bool
	AlphaNative                     bool
	NativeR3                        bool
	FullAFDescent                   bool
	GenerationCarrierMap            bool
	FlavorOrientationMap            bool
	IndividualYukawaValues          bool
	NativeYukawaOperator            bool
}

func (f FirewallLedger) List() []string {
	out := []string{}
	if !f.BoundaryActivationMeasureNative {
		out = append(out, FailureNoNativeBoundaryActivationMeasure)
	}
	if !f.MuBNative {
		out = append(out, FailureMuBFormalNotNative)
	}
	if !f.MeasureUniquenessNative {
		out = append(out, FailureNoNativeMeasureUniqueness)
	}
	if !f.DomainNative {
		out = append(out, FailureNoNativeMuBDomain, FailureReducedResponseBridgeSelected)
	}
	if !f.DegreeExtractionNative {
		out = append(out, FailureNoNativeDegreeExtraction)
	}
	if !f.SSplitTransportNative {
		out = append(out, FailureNoNativeSSplitTransportMap)
	}
	if !f.SelectorNative {
		out = append(out, FailureNoNativeDegreeToZ2FlagFunctor, FailureNoNativeUniqueSelector)
	}
	if !f.ChamberNormalizationNative {
		out = append(out, FailureNoNativeChamberNormalization, FailureNoNativeH1H2Reason)
	}
	if !f.CrossLaneNative {
		out = append(out, FailureNoNativeZ2CrossLaneExclusion, FailureSelectorFunctionhoodNotNative)
	}
	if !f.AlphaNative {
		out = append(out, FailureAlphaByMuBNotNative, FailureAlphaBridgeCandidateNotNative)
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

type Analysis struct {
	ID               string
	Truth            string
	Classification   string
	ShortStatus      string
	Domain           MeasureDomainAudit
	DegreeExtraction DegreeExtractionAudit
	Selector         SelectorIntegrationAudit
	Chambers         ChamberNormalizationAudit
	CrossLanes       CrossLaneMeasureAudit
	Alpha            AlphaReconstructionAudit
	NativeStatus     NativeStatusAudit
	Firewalls        FirewallLedger
	Final            string
}

func BuildDefault() (Analysis, error) {
	linear := float64(RankI1) / float64(RankH10) * SBoundary
	quad := float64(RankI2) / float64(RankH72) * SBoundary * SBoundary
	alpha := linear + quad
	if !near(linear, AlphaLinear) || !near(quad, AlphaQuad) || !near(alpha, AlphaB) {
		return Analysis{}, fmt.Errorf("alpha reconstruction mismatch: linear %.17g quad %.17g alpha %.17g", linear, quad, alpha)
	}

	domain := MeasureDomainAudit{
		InheritedStatus:  Gate919ShortStatus,
		Domain:           ReducedResponse,
		NonzeroDegrees:   []int{1, 2},
		IncludesLambda0:  false,
		ReducedBasepoint: true,
		NativeTheorem:    false,
		Supports:         []string{SupportMuBDomainReducedResponse, SupportMuBActsOnActiveDegrees, SupportMuBIgnoresLambda0Basepoint},
		Failures:         []string{FailureNoNativeMuBDomain, FailureReducedResponseBridgeSelected},
	}

	degree := DegreeExtractionAudit{
		DegreeCoefficients:  map[int]float64{1: SBoundary, 2: SBoundary * SBoundary},
		DegreePowers:        map[int]int{1: 1, 2: 2},
		SeparateS2Transport: false,
		ExteriorGeneratedS2: true,
		NativeTheorem:       false,
		Supports:            []string{SupportMuBExtractsByExteriorDegree, SupportDegreeKCarriesSPowerK, SupportSPowerFollowsExteriorResponse},
		Failures:            []string{FailureNoNativeDegreeExtraction, FailureNoNativeSSplitTransportMap},
	}

	selector := SelectorIntegrationAudit{
		Targets:                   map[int]string{1: "[F_1/F_0]_{Z2}", 2: "[F_2/F_0]_{Z2}"},
		Ranks:                     map[int]int{1: RankI1, 2: RankI2},
		RepresentativeIndependent: true,
		UniqueNativeSelector:      false,
		Supports:                  []string{SupportMuBIntegratesSelector, SupportMuBRecoversRankPair, SupportSelectorRepresentativeIndependent},
		Failures:                  []string{FailureNoNativeDegreeToZ2FlagFunctor, FailureNoNativeUniqueSelector},
	}

	chambers := ChamberNormalizationAudit{
		Chambers:            map[int]string{1: "H_10=H_R^ambient+B_2", 2: "H_72=Lambda^4 V_8+B_2"},
		Ranks:               map[int]int{1: RankH10, 2: RankH72},
		Weights:             map[int]float64{1: float64(RankI1) / float64(RankH10), 2: float64(RankI2) / float64(RankH72)},
		ExplicitLaneWeights: true,
		NativeTheorem:       false,
		Supports:            []string{SupportMuBIntegratesChamberNormalization, SupportMuBRecoversCoefficients, SupportLaneNormalizationExplicit},
		Failures:            []string{FailureNoNativeChamberNormalization, FailureNoNativeH1H2Reason},
	}

	cross := CrossLaneMeasureAudit{
		CorrectTargets:       map[int]string{1: "[F_1/F_0]_{Z2}", 2: "[F_2/F_0]_{Z2}"},
		FalseTargets:         map[int]string{1: "[F_2/F_0]_{Z2}", 2: "[F_1/F_0]_{Z2}"},
		ExcludedIfFunctional: true,
		AbsorbedInIndexing:   true,
		FunctionhoodNative:   false,
		Supports:             []string{SupportMuBExcludesCrossLanesIfFunctional, SupportCrossLaneAbsorbedInIndexing},
		Failures:             []string{FailureNoNativeZ2CrossLaneExclusion, FailureSelectorFunctionhoodNotNative},
	}

	alphaAudit := AlphaReconstructionAudit{
		Formula:               MeasureFormula,
		BoundaryFormula:       BoundaryAlphaMeasureFormula,
		S:                     SBoundary,
		RankI1:                RankI1,
		RankI2:                RankI2,
		RankH1:                RankH10,
		RankH2:                RankH72,
		LinearContribution:    linear,
		QuadraticContribution: quad,
		Alpha:                 alpha,
		ReassemblesFive:       true,
		NativeAlpha:           false,
		Supports:              []string{SupportMuBReconstructsAlpha, SupportMeasureReassemblesFive, SupportAlphaAsMeasure},
		Failures:              []string{FailureAlphaByMuBNotNative},
	}

	native := NativeStatusAudit{
		BridgeMeasureCandidate: true,
		NativeMeasure:          false,
		NativeAlpha:            false,
		NativeR3:               false,
		MissingNativeTheorems:  []string{"native domain theorem", "native degree extraction theorem", "native selector theorem", "native chamber normalization theorem", "native S_split transport theorem", "native measure uniqueness theorem"},
		Supports:               []string{SupportMeasureReassemblesFive, SupportAlphaAsMeasure},
		Failures:               []string{FailureNoNativeBoundaryActivationMeasure, FailureMuBFormalNotNative, FailureNoNativeMeasureUniqueness, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3},
	}

	return Analysis{
		ID:               AuditID,
		Truth:            FinalTruth,
		Classification:   Classification,
		ShortStatus:      ShortStatus,
		Domain:           domain,
		DegreeExtraction: degree,
		Selector:         selector,
		Chambers:         chambers,
		CrossLanes:       cross,
		Alpha:            alphaAudit,
		NativeStatus:     native,
		Firewalls:        FirewallLedger{},
		Final:            StrategicConclusion,
	}, nil
}

func Statuses() []string {
	return []string{StatusInheritedGate919, StatusDomainReducedResponse, StatusDegreeExtraction, StatusSelectorIntegrated, StatusChamberNormalization, StatusCrossLaneAbsorbed, StatusAlphaReconstructed, StatusNativeMeasureMissing, StatusNoNativePromotion}
}

func Supports() []string {
	return []string{SupportMuBDomainReducedResponse, SupportMuBActsOnActiveDegrees, SupportMuBIgnoresLambda0Basepoint, SupportMuBExtractsByExteriorDegree, SupportDegreeKCarriesSPowerK, SupportSPowerFollowsExteriorResponse, SupportMuBIntegratesSelector, SupportMuBRecoversRankPair, SupportSelectorRepresentativeIndependent, SupportMuBIntegratesChamberNormalization, SupportMuBRecoversCoefficients, SupportLaneNormalizationExplicit, SupportMuBExcludesCrossLanesIfFunctional, SupportCrossLaneAbsorbedInIndexing, SupportMuBReconstructsAlpha, SupportMeasureReassemblesFive, SupportAlphaAsMeasure}
}

func Failures() []string {
	return []string{FailureNoNativeBoundaryActivationMeasure, FailureMuBFormalNotNative, FailureNoNativeMeasureUniqueness, FailureNoNativeMuBDomain, FailureReducedResponseBridgeSelected, FailureNoNativeDegreeExtraction, FailureNoNativeSSplitTransportMap, FailureNoNativeDegreeToZ2FlagFunctor, FailureNoNativeUniqueSelector, FailureNoNativeChamberNormalization, FailureNoNativeH1H2Reason, FailureNoNativeZ2CrossLaneExclusion, FailureSelectorFunctionhoodNotNative, FailureAlphaByMuBNotNative, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}
}

func FormatDomain(d MeasureDomainAudit) string {
	return fmt.Sprintf("domain inherited=%s reduced=%t active_degrees=%v lambda0_included=%t native=%t domain=%s", d.InheritedStatus, d.ReducedBasepoint, d.NonzeroDegrees, d.IncludesLambda0, d.NativeTheorem, d.Domain)
}

func FormatDegreeExtraction(d DegreeExtractionAudit) string {
	return fmt.Sprintf("degree_extraction coeffs=%v powers=%v separate_s2_transport=%t exterior_s2=%t native=%t", d.DegreeCoefficients, d.DegreePowers, d.SeparateS2Transport, d.ExteriorGeneratedS2, d.NativeTheorem)
}

func FormatSelector(s SelectorIntegrationAudit) string {
	return fmt.Sprintf("selector targets=%v ranks=%v representative_independent=%t unique_native=%t", s.Targets, s.Ranks, s.RepresentativeIndependent, s.UniqueNativeSelector)
}

func FormatChambers(c ChamberNormalizationAudit) string {
	return fmt.Sprintf("chambers=%v ranks=%v weights=%v explicit_lane_weights=%t native=%t", c.Chambers, c.Ranks, c.Weights, c.ExplicitLaneWeights, c.NativeTheorem)
}

func FormatCrossLanes(c CrossLaneMeasureAudit) string {
	return fmt.Sprintf("cross_lanes correct=%v false=%v excluded_if_functional=%t absorbed_in_indexing=%t functionhood_native=%t", c.CorrectTargets, c.FalseTargets, c.ExcludedIfFunctional, c.AbsorbedInIndexing, c.FunctionhoodNative)
}

func FormatAlpha(a AlphaReconstructionAudit) string {
	return fmt.Sprintf("alpha formula=%s boundary=%s linear=%.17g quad=%.17g alpha=%.17g ranks=(%d,%d;%d,%d) reassembles_five=%t native_alpha=%t", a.Formula, a.BoundaryFormula, a.LinearContribution, a.QuadraticContribution, a.Alpha, a.RankI1, a.RankI2, a.RankH1, a.RankH2, a.ReassemblesFive, a.NativeAlpha)
}

func FormatNativeStatus(n NativeStatusAudit) string {
	return fmt.Sprintf("native_status bridge_measure_candidate=%t native_measure=%t native_alpha=%t native_r3=%t missing=%s", n.BridgeMeasureCandidate, n.NativeMeasure, n.NativeAlpha, n.NativeR3, strings.Join(n.MissingNativeTheorems, "; "))
}

func FormatFirewalls(f FirewallLedger) string { return strings.Join(f.List(), ";") }

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

func near(a, b float64) bool { return math.Abs(a-b) < 1e-18 }

func nearLoose(a, b float64) bool { return math.Abs(a-b) < 1e-12 }

func firewallsOK(f FirewallLedger) bool {
	return !f.BoundaryActivationMeasureNative && !f.MuBNative && !f.MeasureUniquenessNative && !f.DomainNative && !f.DegreeExtractionNative && !f.SSplitTransportNative && !f.SelectorNative && !f.ChamberNormalizationNative && !f.CrossLaneNative && !f.AlphaNative && !f.NativeR3 && !f.FullAFDescent && !f.GenerationCarrierMap && !f.FlavorOrientationMap && !f.IndividualYukawaValues && !f.NativeYukawaOperator
}
