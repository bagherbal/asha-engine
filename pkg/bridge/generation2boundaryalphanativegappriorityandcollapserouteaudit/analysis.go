// Package generation2boundaryalphanativegappriorityandcollapserouteaudit implements
// Gate 919: BoundaryAlpha NativeGap Priority and CollapseRoute Audit.
//
// Gate 919 follows Gate 918's status that BoundaryAlpha_Z2 is no longer an
// opaque seal: it is a decomposed bridge-theorem candidate with five explicit
// native gaps. This audit does not solve those gaps. It asks whether they are
// truly independent or whether they are projections of one deeper missing
// object: a BoundaryActivationMeasure / Z2 BoundaryResponseFunctor.
//
// The audit ranks the remaining alpha-side gaps, identifies S_split transport
// as the highest-priority subgap, and records the strongest collapse route as a
// formal measure-like functor. It preserves that no native BoundaryActivation
// measure, native alpha theorem, native R3 theorem, full A_F descent,
// generation/flavor carrier, physical Yukawa spectrum, or official ledger
// update has been certified.
package generation2boundaryalphanativegappriorityandcollapserouteaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE919-BOUNDARYALPHA-NATIVEGAP-PRIORITY-COLLAPSEROUTE-AUDIT"

	SBoundary = 0.0012924448188162962
	AlphaB    = 0.0003878958469680527

	RankF1OverF0 = 3
	RankF2OverF0 = 7
	RankH10      = 10
	RankH72      = 72

	AlphaLinear = 0.00038773344564488885
	AlphaQuad   = 0.0000001624013231638281

	Gate918ShortStatus = "R3_ALPHA_DECOMPOSED_BRIDGE_CANDIDATE_NATIVE_GAPS_EXPLICIT"

	BoundaryActivationMeasureName = "BoundaryActivationMeasure"
	BoundaryResponseFunctorName   = "Z2BoundaryResponseFunctor"
	MuBFormula                    = "mu_B(R_B(S_split))=sum_{k=1}^2 rank(I_B^Z2(k))/rank(H_k)*S_split^k"
	MuBBranchFormula              = "mu_B(R_B(S_split))=(3/10)S_split+(7/72)S_split^2"
	CollapsedGap                  = "one candidate master native gap: BoundaryActivationMeasure"

	NextGate            = "NEXT_PRESSURE_GATE920_BOUNDARYACTIVATIONMEASURE_FUNCTOR_AUDIT"
	Classification      = "R3_BOUNDARYALPHA_NATIVE_GAPS_COLLAPSE_TO_BOUNDARY_ACTIVATION_MEASURE_CANDIDATE"
	ShortStatus         = "R3_ALPHA_GAPS_COLLAPSE_TO_BOUNDARY_MEASURE_OBSTRUCTION"
	FinalTruth          = "BOUNDARYALPHA_NATIVE_GAPS_PRIORITIZED_AND_COLLAPSE_ROUTE_IDENTIFIED"
	StrategicConclusion = "Gate 919 redirects the branch away from five disconnected proof searches. The five alpha-side native gaps share a boundary response measure structure, so the strongest next pressure object is a BoundaryActivationMeasure / Z2 BoundaryResponseFunctor. This is a collapse-route candidate only: alpha_B remains a decomposed bridge candidate, not a native theorem."

	StatusInheritedGate918              = "PASS_GATE918_DECOMPOSED_ALPHA_BRIDGE_CANDIDATE_INHERITED"
	StatusSharedBoundaryMeasurePattern  = "PASS_FIVE_NATIVE_GAPS_SHARE_BOUNDARY_RESPONSE_MEASURE_PATTERN"
	StatusPriorityRankingCompleted      = "PASS_NATIVE_GAP_PRIORITY_RANKING_COMPLETED"
	StatusSSplitHighestPriority         = "PASS_S_SPLIT_TRANSPORT_IDENTIFIED_AS_HIGHEST_PRIORITY_SUBGAP"
	StatusCollapseRouteIdentified       = "PASS_BOUNDARY_ACTIVATION_MEASURE_SELECTED_AS_COLLAPSE_ROUTE_CANDIDATE"
	StatusMuBFormalReassemblesAlpha     = "PASS_FORMAL_MU_B_REASSEMBLES_ALPHA_FROM_DEGREE_TARGETS_AND_CHAMBERS"
	StatusMasterMeasureRequirementsOpen = "FIREWALL_PRESERVED_MASTER_MEASURE_REQUIREMENTS_OPEN"
	StatusAlphaNotPromoted              = "FIREWALL_PRESERVED_ALPHA_NOT_PROMOTED_TO_NATIVE"

	SupportFiveGapsShareMeasureStructure          = "CONDITIONAL_SUPPORT_FIVE_NATIVE_GAPS_SHARE_BOUNDARY_RESPONSE_MEASURE_STRUCTURE"
	SupportAlphaGapsMayCollapseToMeasure          = "CONDITIONAL_SUPPORT_ALPHA_GAPS_MAY_COLLAPSE_TO_BOUNDARY_ACTIVATION_MEASURE"
	SupportSingleMasterFunctorCouldGeneratePieces = "CONDITIONAL_SUPPORT_SINGLE_MASTER_FUNCTOR_COULD_GENERATE_RESPONSE_SELECTOR_TRANSPORT_AND_NORMALIZATION"
	SupportNativeGapPriorityRankingCompleted      = "CONDITIONAL_SUPPORT_NATIVE_GAP_PRIORITY_RANKING_COMPLETED"
	SupportSSplitTransportHighestPriority         = "CONDITIONAL_SUPPORT_S_SPLIT_TRANSPORT_IS_HIGHEST_PRIORITY_SUBGAP"
	SupportBoundaryActivationMeasureStrongest     = "CONDITIONAL_SUPPORT_BOUNDARY_ACTIVATION_MEASURE_IS_STRONGEST_COLLAPSE_ROUTE"
	SupportBoundaryActivationMeasureReassembles   = "CONDITIONAL_SUPPORT_BOUNDARY_ACTIVATION_MEASURE_CAN_REASSEMBLE_ALL_FIVE_SUBOBJECTS"
	SupportMuBReconstructsAlpha                   = "CONDITIONAL_SUPPORT_MU_B_FORMULA_RECONSTRUCTS_ALPHA_B_FROM_DEGREE_TARGETS_AND_CHAMBERS"
	SupportNativeAlphaMayNeedMeasureFunctor       = "CONDITIONAL_SUPPORT_NATIVE_ALPHA_THEOREM_MAY_REQUIRE_MEASURE_FUNCTOR_NOT_FIVE_SEPARATE_THEOREMS"
	SupportNativeBoundaryAlphaAsMeasure           = "CONDITIONAL_SUPPORT_NATIVE_BOUNDARY_ALPHA_COULD_BE_FORMULATED_AS_BOUNDARY_ACTIVATION_MEASURE"
	SupportMeasureAbsorbsPieces                   = "CONDITIONAL_SUPPORT_MEASURE_FORM_ABSORBS_RESPONSE_SELECTOR_TRANSPORT_NORMALIZATION_AND_EXCLUSION"
	SupportAlphaGapsCollapseCandidate             = "CONDITIONAL_SUPPORT_ALPHA_NATIVE_GAPS_COLLAPSE_TO_BOUNDARY_ACTIVATION_MEASURE_CANDIDATE"

	FailureNoNativeBoundaryActivationMeasureCertified = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_ACTIVATION_MEASURE_CERTIFIED"
	FailureNoNativeBoundaryResponseMeasure            = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_RESPONSE_MEASURE"
	FailureMuBFormalNotNative                         = "FAILED_ROUTE_MU_B_IS_FORMAL_REASSEMBLY_NOT_NATIVE_MEASURE_THEOREM"
	FailureNoNativeSSplitTransportMap                 = "FAILED_ROUTE_NO_NATIVE_S_SPLIT_TRANSPORT_MAP"
	FailureNoNativeReducedB2ResponseFunctional        = "FAILED_ROUTE_NO_NATIVE_REDUCED_B2_RESPONSE_FUNCTIONAL"
	FailureNoNativeDegreeToZ2FlagClassFunctor         = "FAILED_ROUTE_NO_NATIVE_DEGREE_TO_Z2_FLAG_CLASS_FUNCTOR"
	FailureNoNativeZ2CrossLaneExclusionTheorem        = "FAILED_ROUTE_NO_NATIVE_Z2_CROSS_LANE_EXCLUSION_THEOREM"
	FailureNoNativeResponseChamberNormalization       = "FAILED_ROUTE_NO_NATIVE_RESPONSE_CHAMBER_NORMALIZATION_THEOREM"
	FailureAlphaBridgeCandidateNotNative              = "FAILED_ROUTE_ALPHA_B_REMAINS_BRIDGE_CANDIDATE_NOT_NATIVE"
	FailureNotNativeR3                                = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked                  = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap                     = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap                     = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues                   = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator                     = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type NativeGap struct {
	Name                string
	Priority            int
	PriorityLabel       string
	Reason              string
	MayFollowFromMaster bool
	Failure             string
}

type GapIndependenceAudit struct {
	InheritedStatus       string
	GapNames              []string
	SharedStructure       []string
	InitiallyIndependent  bool
	MayCollapseToMeasure  bool
	CandidateMasterObject string
	NativeCertified       bool
	Supports, Failures    []string
}

type PriorityRanking struct {
	OrderedGaps            []NativeGap
	HighestPriorityGap     string
	BoundaryActivationRank int
	CrossLaneDependent     bool
	RankingComplete        bool
	Supports, Failures     []string
}

type CollapseRoute struct {
	MasterObject          string
	AlternateName         string
	Formula               string
	BranchFormula         string
	I1Target              string
	I2Target              string
	H1                    string
	H2                    string
	RankI1                int
	RankI2                int
	RankH1                int
	RankH2                int
	S                     float64
	LinearContribution    float64
	QuadraticContribution float64
	Alpha                 float64
	ReassemblesAllFive    bool
	NativeTheorem         bool
	Supports, Failures    []string
}

type MasterMeasureRequirements struct {
	Requirements       []string
	AllRequired        bool
	AllCertified       bool
	Supports, Failures []string
}

type PromotionStatus struct {
	PreviousStatus     string
	NewStatus          string
	BridgeCandidate    bool
	NativeAlpha        bool
	NativeR3           bool
	OfficialUpdate     bool
	Supports, Failures []string
}

type FirewallLedger struct {
	BoundaryActivationMeasureCertified bool
	BoundaryResponseMeasureCertified   bool
	MuBNativeMeasureTheorem            bool
	SSplitTransportNative              bool
	ReducedB2ResponseNative            bool
	DegreeToZ2FlagFunctorNative        bool
	CrossLaneExclusionNative           bool
	ResponseChamberNormalizationNative bool
	AlphaNative                        bool
	NativeR3                           bool
	FullAFDescent                      bool
	GenerationCarrierMap               bool
	FlavorOrientationMap               bool
	IndividualYukawaValues             bool
	NativeYukawaOperator               bool
}

func (f FirewallLedger) List() []string {
	out := []string{}
	if !f.BoundaryActivationMeasureCertified {
		out = append(out, FailureNoNativeBoundaryActivationMeasureCertified)
	}
	if !f.BoundaryResponseMeasureCertified {
		out = append(out, FailureNoNativeBoundaryResponseMeasure)
	}
	if !f.MuBNativeMeasureTheorem {
		out = append(out, FailureMuBFormalNotNative)
	}
	if !f.SSplitTransportNative {
		out = append(out, FailureNoNativeSSplitTransportMap)
	}
	if !f.ReducedB2ResponseNative {
		out = append(out, FailureNoNativeReducedB2ResponseFunctional)
	}
	if !f.DegreeToZ2FlagFunctorNative {
		out = append(out, FailureNoNativeDegreeToZ2FlagClassFunctor)
	}
	if !f.CrossLaneExclusionNative {
		out = append(out, FailureNoNativeZ2CrossLaneExclusionTheorem)
	}
	if !f.ResponseChamberNormalizationNative {
		out = append(out, FailureNoNativeResponseChamberNormalization)
	}
	if !f.AlphaNative {
		out = append(out, FailureAlphaBridgeCandidateNotNative)
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

type Audit struct {
	ID              string
	Truth           string
	Classification  string
	ShortStatus     string
	GapIndependence GapIndependenceAudit
	PriorityRanking PriorityRanking
	CollapseRoute   CollapseRoute
	Requirements    MasterMeasureRequirements
	Promotion       PromotionStatus
	Firewalls       FirewallLedger
	Final           string
}

func BuildDefault() (Audit, error) {
	linear := float64(RankF1OverF0) / float64(RankH10) * SBoundary
	quad := float64(RankF2OverF0) / float64(RankH72) * SBoundary * SBoundary
	alpha := linear + quad
	if !near(linear, AlphaLinear) || !near(quad, AlphaQuad) || !near(alpha, AlphaB) {
		return Audit{}, fmt.Errorf("alpha reconstruction mismatch: linear %.18g quad %.18g total %.18g", linear, quad, alpha)
	}

	gapNames := []string{
		"native reduced B2 response functional",
		"native degree-to-Z2-flag-class functor",
		"native Z2 cross-lane exclusion theorem",
		"native S_split transport map",
		"native response-chamber normalization theorem",
	}

	ordered := []NativeGap{
		{Name: "S_split -> s transport map", Priority: 1, PriorityLabel: "highest", Reason: "R_B(s) cannot become an activation law unless the scalar input s is typed from S_split.", MayFollowFromMaster: false, Failure: FailureNoNativeSSplitTransportMap},
		{Name: "BoundaryActivationMeasure / response measure", Priority: 2, PriorityLabel: "very high", Reason: "A master measure could unify response shape, selector, transport, normalization, and exclusion.", MayFollowFromMaster: false, Failure: FailureNoNativeBoundaryResponseMeasure},
		{Name: "degree-to-Z2-flag selector", Priority: 3, PriorityLabel: "high", Reason: "Essential for target ranks, but likely a component of the master measure.", MayFollowFromMaster: true, Failure: FailureNoNativeDegreeToZ2FlagClassFunctor},
		{Name: "boundary response-chamber normalization", Priority: 4, PriorityLabel: "medium-high", Reason: "Required for denominators 10 and 72, but likely subordinate to the measure's lane normalization rule.", MayFollowFromMaster: true, Failure: FailureNoNativeResponseChamberNormalization},
		{Name: "Z2 cross-lane exclusion", Priority: 5, PriorityLabel: "dependent", Reason: "Expected to follow from a native degree-indexed selector/functionhood theorem.", MayFollowFromMaster: true, Failure: FailureNoNativeZ2CrossLaneExclusionTheorem},
	}

	return Audit{
		ID:             AuditID,
		Truth:          FinalTruth,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		GapIndependence: GapIndependenceAudit{
			InheritedStatus:       Gate918ShortStatus,
			GapNames:              gapNames,
			SharedStructure:       []string{"boundary degree", "boundary response", "boundary target", "boundary normalization", "boundary scalar insertion"},
			InitiallyIndependent:  true,
			MayCollapseToMeasure:  true,
			CandidateMasterObject: BoundaryActivationMeasureName,
			NativeCertified:       false,
			Supports:              []string{SupportFiveGapsShareMeasureStructure, SupportAlphaGapsMayCollapseToMeasure, SupportSingleMasterFunctorCouldGeneratePieces},
			Failures:              []string{FailureNoNativeBoundaryActivationMeasureCertified},
		},
		PriorityRanking: PriorityRanking{
			OrderedGaps:            ordered,
			HighestPriorityGap:     ordered[0].Name,
			BoundaryActivationRank: 2,
			CrossLaneDependent:     true,
			RankingComplete:        true,
			Supports:               []string{SupportNativeGapPriorityRankingCompleted, SupportSSplitTransportHighestPriority, SupportBoundaryActivationMeasureStrongest},
			Failures:               []string{FailureNoNativeSSplitTransportMap, FailureNoNativeBoundaryResponseMeasure, FailureNoNativeDegreeToZ2FlagClassFunctor, FailureNoNativeResponseChamberNormalization, FailureNoNativeZ2CrossLaneExclusionTheorem},
		},
		CollapseRoute: CollapseRoute{
			MasterObject:          BoundaryActivationMeasureName,
			AlternateName:         BoundaryResponseFunctorName,
			Formula:               MuBFormula,
			BranchFormula:         MuBBranchFormula,
			I1Target:              "[F_1/F_0]_{Z2}",
			I2Target:              "[F_2/F_0]_{Z2}",
			H1:                    "H_10",
			H2:                    "H_72",
			RankI1:                RankF1OverF0,
			RankI2:                RankF2OverF0,
			RankH1:                RankH10,
			RankH2:                RankH72,
			S:                     SBoundary,
			LinearContribution:    linear,
			QuadraticContribution: quad,
			Alpha:                 alpha,
			ReassemblesAllFive:    true,
			NativeTheorem:         false,
			Supports:              []string{SupportBoundaryActivationMeasureReassembles, SupportMuBReconstructsAlpha, SupportNativeAlphaMayNeedMeasureFunctor},
			Failures:              []string{FailureMuBFormalNotNative},
		},
		Requirements: MasterMeasureRequirements{
			Requirements: []string{
				"source: R_B(s) is the reduced active boundary response",
				"parameter: S_split transports to s",
				"degree: exterior degree k indexes response order k",
				"target: degree k selects the Z2 flag quotient I_B^Z2(k)",
				"normalizer: each lane is divided by its response chamber H_k",
				"exclusion: no non-degree-matched lanes appear",
			},
			AllRequired:  true,
			AllCertified: false,
			Supports:     []string{SupportNativeBoundaryAlphaAsMeasure, SupportMeasureAbsorbsPieces},
			Failures:     []string{FailureNoNativeBoundaryActivationMeasureCertified},
		},
		Promotion: PromotionStatus{
			PreviousStatus:  "five explicit native gaps",
			NewStatus:       CollapsedGap,
			BridgeCandidate: true,
			NativeAlpha:     false,
			NativeR3:        false,
			OfficialUpdate:  false,
			Supports:        []string{SupportAlphaGapsCollapseCandidate},
			Failures:        []string{FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3},
		},
		Firewalls: FirewallLedger{},
		Final:     StrategicConclusion,
	}, nil
}

func Statuses() []string {
	return []string{StatusInheritedGate918, StatusSharedBoundaryMeasurePattern, StatusPriorityRankingCompleted, StatusSSplitHighestPriority, StatusCollapseRouteIdentified, StatusMuBFormalReassemblesAlpha, StatusMasterMeasureRequirementsOpen, StatusAlphaNotPromoted}
}

func Supports() []string {
	return []string{SupportFiveGapsShareMeasureStructure, SupportAlphaGapsMayCollapseToMeasure, SupportSingleMasterFunctorCouldGeneratePieces, SupportNativeGapPriorityRankingCompleted, SupportSSplitTransportHighestPriority, SupportBoundaryActivationMeasureStrongest, SupportBoundaryActivationMeasureReassembles, SupportMuBReconstructsAlpha, SupportNativeAlphaMayNeedMeasureFunctor, SupportNativeBoundaryAlphaAsMeasure, SupportMeasureAbsorbsPieces, SupportAlphaGapsCollapseCandidate}
}

func Failures() []string {
	return []string{FailureNoNativeBoundaryActivationMeasureCertified, FailureNoNativeBoundaryResponseMeasure, FailureMuBFormalNotNative, FailureNoNativeSSplitTransportMap, FailureNoNativeReducedB2ResponseFunctional, FailureNoNativeDegreeToZ2FlagClassFunctor, FailureNoNativeZ2CrossLaneExclusionTheorem, FailureNoNativeResponseChamberNormalization, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}
}

func FormatGapIndependence(g GapIndependenceAudit) string {
	return fmt.Sprintf("inherited=%s; gaps=%s; shared=%s; initially_independent=%t; may_collapse_to_measure=%t; candidate=%s; native_certified=%t", g.InheritedStatus, strings.Join(g.GapNames, " | "), strings.Join(g.SharedStructure, " | "), g.InitiallyIndependent, g.MayCollapseToMeasure, g.CandidateMasterObject, g.NativeCertified)
}

func FormatPriorityRanking(p PriorityRanking) string {
	parts := make([]string, 0, len(p.OrderedGaps))
	for _, g := range p.OrderedGaps {
		parts = append(parts, fmt.Sprintf("%d:%s[%s]", g.Priority, g.Name, g.PriorityLabel))
	}
	return fmt.Sprintf("ranking=%s; highest=%s; boundary_activation_rank=%d; cross_lane_dependent=%t; complete=%t", strings.Join(parts, " > "), p.HighestPriorityGap, p.BoundaryActivationRank, p.CrossLaneDependent, p.RankingComplete)
}

func FormatCollapseRoute(c CollapseRoute) string {
	return fmt.Sprintf("master=%s/%s; formula=%s; branch=%s; I1=%s rank=%d; I2=%s rank=%d; H1=%s rank=%d; H2=%s rank=%d; linear=%.18g; quad=%.18g; alpha=%.18g; reassembles_all_five=%t; native=%t", c.MasterObject, c.AlternateName, c.Formula, c.BranchFormula, c.I1Target, c.RankI1, c.I2Target, c.RankI2, c.H1, c.RankH1, c.H2, c.RankH2, c.LinearContribution, c.QuadraticContribution, c.Alpha, c.ReassemblesAllFive, c.NativeTheorem)
}

func FormatRequirements(r MasterMeasureRequirements) string {
	return fmt.Sprintf("requirements=%s; all_required=%t; all_certified=%t", strings.Join(r.Requirements, " | "), r.AllRequired, r.AllCertified)
}

func FormatPromotion(p PromotionStatus) string {
	return fmt.Sprintf("previous=%s; new=%s; bridge_candidate=%t; native_alpha=%t; native_r3=%t; official_update=%t", p.PreviousStatus, p.NewStatus, p.BridgeCandidate, p.NativeAlpha, p.NativeR3, p.OfficialUpdate)
}

func FormatFirewalls(f FirewallLedger) string { return strings.Join(f.List(), " | ") }

func near(a, b float64) bool { return math.Abs(a-b) <= 1e-18 }

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

func firewallsOK(f FirewallLedger) bool {
	return !f.BoundaryActivationMeasureCertified && !f.BoundaryResponseMeasureCertified && !f.MuBNativeMeasureTheorem && !f.SSplitTransportNative && !f.ReducedB2ResponseNative && !f.DegreeToZ2FlagFunctorNative && !f.CrossLaneExclusionNative && !f.ResponseChamberNormalizationNative && !f.AlphaNative && !f.NativeR3 && !f.FullAFDescent && !f.GenerationCarrierMap && !f.FlavorOrientationMap && !f.IndividualYukawaValues && !f.NativeYukawaOperator
}
