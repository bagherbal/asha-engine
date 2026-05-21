// Package generation2z2boundaryalphadecomposedfunctorconsolidationaudit implements
// Gate 918: Z2 BoundaryAlpha DecomposedFunctor Consolidation and Native-Theorem Gap Audit.
//
// Gate 918 follows Gates 913--917, which audited the five Gate 912
// BoundaryAlpha_Z2 sub-objects at shape level:
//
//  1. reduced B2 response shape
//  2. degree-indexed Z2 flag selector shape
//  3. cross-lane exclusion shape
//  4. S_split uniform insertion shape
//  5. boundary-augmented response-chamber normalization shape
//
// The audit reassembles these pieces into the strongest current
// BoundaryAlpha_Z2 candidate. It upgrades the alpha branch from an opaque seal
// to a decomposed bridge-theorem candidate, while preserving that no native
// reduced response, degree functor, cross-lane theorem, S_split transport,
// chamber normalization, native alpha theorem, or native R3 theorem has been
// certified.
package generation2z2boundaryalphadecomposedfunctorconsolidationaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE918-Z2-BOUNDARYALPHA-DECOMPOSED-FUNCTOR-CONSOLIDATION-AUDIT"

	SBoundary = 0.0012924448188162962
	AlphaB    = 0.0003878958469680527

	RankF1OverF0  = 3
	RankF2OverF0  = 7
	RankH10       = 10
	RankH72       = 72
	RankB2        = 2
	RankHRAmbient = 8
	RankLambda4V8 = 70

	AlphaLinear = 0.00038773344564488885
	AlphaQuad   = 0.0000001624013231638281

	NEffOperator    = 3.002327375081808
	CYukawaOperator = 0.9992248096922658
	CHiggsOperator  = 1.037220510866514

	Gate913ShortStatus = "R3_ALPHA_SUBOBJECT_1_REDUCED_B2_RESPONSE_SHAPE_PASS_NATIVE_SELECTION_BLOCKED"
	Gate914ShortStatus = "R3_DEGREE_INDEXED_Z2_AIRLOCK_FLAG_SELECTOR_OBSTRUCTION"
	Gate915ShortStatus = "R3_Z2_BOUNDARYALPHA_CROSS_LANE_EXCLUSION_OBSTRUCTION"
	Gate916ShortStatus = "R3_S_SPLIT_TO_REDUCED_B2_RESPONSE_TRANSPORT_OBSTRUCTION"
	Gate917ShortStatus = "R3_BOUNDARY_AUGMENTED_RESPONSE_CHAMBER_NORMALIZATION_OBSTRUCTION"

	ReducedResponseShape = "R_B(s)=(1+s b1)(1+s b2)-1=s(b1+b2)+s^2(b1 wedge b2)"
	DegreeSelectorShape  = "deg(Lambda^1 B_2)->[F_1/F_0]_{Z2}; deg(Lambda^2 B_2)->[F_2/F_0]_{Z2}"
	CrossLaneShape       = "Lambda^1 B_2 not->[F_2/F_0]_{Z2}; Lambda^2 B_2 not->[F_1/F_0]_{Z2}"
	SSplitTransportShape = "T_s(S_split)=s uniformly inserted into (1+s b1) and (1+s b2)"
	NormalizationShape   = "H_10=H_R^ambient plus B_2; H_72=Lambda^4 V_8 plus B_2"
	ConsolidatedFormula  = "alpha_B^Z2=rank([F_1/F_0]_{Z2})/rank(H_10)*s + rank([F_2/F_0]_{Z2})/rank(H_72)*s^2"

	NextGate            = "NEXT_PRESSURE_GATE919_BOUNDARYALPHA_NATIVEGAP_PRIORITY_AND_COLLAPSEROUTE_AUDIT"
	Classification      = "R3_Z2_BOUNDARYALPHA_DECOMPOSED_BRIDGE_CANDIDATE_NOT_NATIVE"
	ShortStatus         = "R3_ALPHA_DECOMPOSED_BRIDGE_CANDIDATE_NATIVE_GAPS_EXPLICIT"
	FinalTruth          = "Z2_BOUNDARYALPHA_REASSEMBLED_AS_DECOMPOSED_BRIDGE_THEOREM_CANDIDATE_BUT_NATIVE_THEOREM_GAPS_REMAIN"
	StrategicConclusion = "Gate 918 promotes BoundaryAlpha_Z2 from an opaque class seal to a decomposed bridge-theorem candidate. The five shape-level subobjects compose coherently and reconstruct alpha_B, but native R3 remains blocked by native alpha theorem gaps plus full A_F descent / lawful spontaneous-orientation status."

	StatusFiveSubobjectsInherited        = "PASS_GATES913_TO917_FIVE_ALPHA_SUBOBJECTS_INHERITED_AT_SHAPE_LEVEL"
	StatusInternalCoherence              = "PASS_FIVE_SUBOBJECTS_COMPOSE_WITHOUT_SHAPE_LEVEL_CONTRADICTION"
	StatusAlphaNoLongerOpaqueSeal        = "PASS_ALPHA_B_NO_LONGER_SINGLE_OPAQUE_SEAL"
	StatusRepresentativeIndependent      = "PASS_BOUNDARYALPHA_Z2_REPRESENTATIVE_INDEPENDENT_ON_Z2_AIRLOCK_CLASS"
	StatusBridgeCandidatePromoted        = "PASS_BOUNDARYALPHA_Z2_PROMOTED_TO_DECOMPOSED_BRIDGE_CANDIDATE"
	StatusR3TraceLedgerUnderCandidate    = "PASS_R3_TRACE_LEDGER_RESTS_ON_DECOMPOSED_ALPHA_BRIDGE_CANDIDATE"
	StatusNativeGapsExplicit             = "PASS_NATIVE_R3_BLOCKERS_EXPLICIT_ALPHA_GAPS_PLUS_FULL_AF_DESCENT"
	StatusNativeTheoremFirewallsPreserve = "FIREWALL_PRESERVED_DECOMPOSED_BRIDGE_CANDIDATE_NOT_NATIVE_THEOREM"

	SupportDecomposedFunctorCoherent       = "CONDITIONAL_SUPPORT_Z2_BOUNDARYALPHA_DECOMPOSED_FUNCTOR_IS_INTERNALLY_COHERENT"
	SupportFiveSubobjectsCompose           = "CONDITIONAL_SUPPORT_FIVE_SUBOBJECTS_COMPOSE_TO_RECONSTRUCT_ALPHA_B"
	SupportAlphaNoLongerOpaque             = "CONDITIONAL_SUPPORT_ALPHA_B_NO_LONGER_SINGLE_OPAQUE_SEAL"
	SupportAlphaBridgeCandidateForm        = "CONDITIONAL_SUPPORT_ALPHA_B_HAS_DECOMPOSED_BRIDGE_THEOREM_CANDIDATE_FORM"
	SupportBoundaryAlphaRepresentativeFree = "CONDITIONAL_SUPPORT_BOUNDARYALPHA_Z2_IS_REPRESENTATIVE_INDEPENDENT"
	SupportPhaseSignAbsentAfterQuotient    = "CONDITIONAL_SUPPORT_PHASE_SIGN_DOES_NOT_ENTER_ALPHA_B_AFTER_Z2_QUOTIENT"
	SupportZ2AirlockCorrectDomain          = "CONDITIONAL_SUPPORT_Z2_AIRLOCK_CLASS_IS_THE_CORRECT_ALPHA_DOMAIN"
	SupportPromotedFromOpaqueSeal          = "CONDITIONAL_SUPPORT_BOUNDARYALPHA_Z2_CAN_BE_PROMOTED_FROM_OPAQUE_SEAL_TO_DECOMPOSED_BRIDGE_CANDIDATE"
	SupportAlphaTheoremObligationLedger    = "CONDITIONAL_SUPPORT_ALPHA_BRANCH_NOW_HAS_THEOREM_OBLIGATION_LEDGER"
	SupportR3TraceLedgerOnCandidate        = "CONDITIONAL_SUPPORT_R3_TRACE_LEDGER_NOW_RESTS_ON_DECOMPOSED_ALPHA_BRIDGE_CANDIDATE"
	SupportOperatorsRemainReconstructed    = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_C_YUKAWA_C_HIGGS_REMAIN_RECONSTRUCTED_UNDER_DECOMPOSED_ALPHA_CANDIDATE"
	SupportNativeR3BlockersExplicit        = "CONDITIONAL_SUPPORT_NATIVE_R3_BLOCKERS_ARE_NOW_ALPHA_THEOREM_GAPS_PLUS_FULL_AF_DESCENT"

	FailureNotNativeR3                          = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureInternalCoherenceNotNative           = "FAILED_ROUTE_INTERNAL_COHERENCE_NOT_NATIVE_THEOREM"
	FailureRepresentativeIndependenceNotNative  = "FAILED_ROUTE_REPRESENTATIVE_INDEPENDENCE_NOT_NATIVE_AIRLOCK_FUNCTOR"
	FailureNoNativeReducedB2ResponseFunctional  = "FAILED_ROUTE_NO_NATIVE_REDUCED_B2_RESPONSE_FUNCTIONAL"
	FailureNoNativeDegreeToZ2FlagClassFunctor   = "FAILED_ROUTE_NO_NATIVE_DEGREE_TO_Z2_FLAG_CLASS_FUNCTOR"
	FailureNoNativeZ2CrossLaneExclusionTheorem  = "FAILED_ROUTE_NO_NATIVE_Z2_CROSS_LANE_EXCLUSION_THEOREM"
	FailureNoNativeSSplitTransportMap           = "FAILED_ROUTE_NO_NATIVE_S_SPLIT_TRANSPORT_MAP"
	FailureNoNativeResponseChamberNormalization = "FAILED_ROUTE_NO_NATIVE_RESPONSE_CHAMBER_NORMALIZATION_THEOREM"
	FailureReconstructionNotNativeAlpha         = "FAILED_ROUTE_RECONSTRUCTION_FROM_FIVE_SHAPE_LEVEL_SUBOBJECTS_NOT_NATIVE_ALPHA_THEOREM"
	FailureAlphaBridgeCandidateNotNative        = "FAILED_ROUTE_ALPHA_B_REMAINS_BRIDGE_CANDIDATE_NOT_NATIVE"
	FailureFullAFDescentStillBlocked            = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureHiggsOrientationSealed               = "FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_AS_ORIENTATION_CLASS"
	FailureNoGenerationCarrierMap               = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap               = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues             = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator               = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoOfficialNEffUpdateAllowed          = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaOrCHiggsUpdateAllowed       = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
)

type SubobjectStack struct {
	Gate913ShortStatus string
	Gate914ShortStatus string
	Gate915ShortStatus string
	Gate916ShortStatus string
	Gate917ShortStatus string
	ResponseShape      string
	SelectorShape      string
	CrossLaneShape     string
	TransportShape     string
	NormalizationShape string
	AllAuditedAtShape  bool
	NativeTheorem      bool
	Supports, Failures []string
}

type ReassembledBoundaryAlpha struct {
	Formula               string
	RankF1OverF0          int
	RankF2OverF0          int
	RankH10               int
	RankH72               int
	S                     float64
	LinearContribution    float64
	QuadraticContribution float64
	TotalAlpha            float64
	InternalCoherent      bool
	OpaqueSeal            bool
	BridgeCandidate       bool
	NativeTheorem         bool
	Supports, Failures    []string
}

type RepresentativeIndependence struct {
	Domain                  string
	RankPair                [2]int
	TauPhiPreservesRankPair bool
	PhaseSignEntersAlpha    bool
	CorrectAlphaDomain      bool
	NativeAirlockFunctor    bool
	Supports, Failures      []string
}

type BridgeCandidateStatus struct {
	PreviousStatus       string
	NewStatus            string
	TheoremObligations   []string
	AllVisibleComponents bool
	NativeTheorem        bool
	Supports, Failures   []string
}

type R3TraceLedgerStatus struct {
	TraceRows             []TraceRow
	ATotalOverT           float64
	BTotalOverT2          float64
	NEffOperator          float64
	CYukawaOperator       float64
	CHiggsOperator        float64
	DiagnosticOnly        bool
	OfficialUpdateAllowed bool
	Supports, Failures    []string
}

type TraceRow struct {
	Rank   int
	Weight string
}

type NativeGapLedger struct {
	AlphaGaps                 []string
	FiniteLayerGaps           []string
	GenerationFlavorR4OrLater bool
	NativeR3                  bool
	Supports, Failures        []string
}

type FirewallLedger struct {
	NativeR3                         bool
	InternalCoherenceNative          bool
	RepresentativeIndependenceNative bool
	NativeReducedB2Response          bool
	NativeDegreeFunctor              bool
	NativeCrossLane                  bool
	NativeSSplitTransport            bool
	NativeNormalization              bool
	ReconstructionNativeAlpha        bool
	AlphaNative                      bool
	FullAFDescent                    bool
	HiggsOrientationNative           bool
	GenerationCarrier                bool
	FlavorOrientation                bool
	IndividualYukawa                 bool
	NativeYukawa                     bool
	OfficialNEffUpdate               bool
	OfficialCoefficientUpdate        bool
	Failures                         []string
}

type Audit struct {
	ID                         string
	Truth                      string
	Classification             string
	ShortStatus                string
	Subobjects                 SubobjectStack
	BoundaryAlpha              ReassembledBoundaryAlpha
	RepresentativeIndependence RepresentativeIndependence
	BridgeCandidate            BridgeCandidateStatus
	R3TraceLedger              R3TraceLedgerStatus
	NativeGaps                 NativeGapLedger
	Firewalls                  FirewallLedger
	Final                      string
}

func BuildDefault() (Audit, error) {
	linear := float64(RankF1OverF0) / float64(RankH10) * SBoundary
	quadratic := float64(RankF2OverF0) / float64(RankH72) * SBoundary * SBoundary
	alpha := linear + quadratic
	if !near(linear, AlphaLinear) || !near(quadratic, AlphaQuad) || !near(alpha, AlphaB) {
		return Audit{}, fmt.Errorf("alpha reconstruction mismatch: linear %.18g quadratic %.18g total %.18g", linear, quadratic, alpha)
	}
	aTotal := 3 + 3*AlphaB
	bTotal := 3 + 3*AlphaB*AlphaB - 6*math.Pow(AlphaB, 3) + 12*math.Pow(AlphaB, 4)
	computedNEff := 3 * math.Pow(1+AlphaB, 2) / (1 + AlphaB*AlphaB - 2*math.Pow(AlphaB, 3) + 4*math.Pow(AlphaB, 4))
	if !nearLoose(computedNEff, NEffOperator) {
		return Audit{}, fmt.Errorf("N_eff mismatch: got %.18g want %.18g", computedNEff, NEffOperator)
	}

	alphaGaps := []string{
		"native reduced B2 response functional",
		"native degree-indexed Z2 flag functor",
		"native Z2 cross-lane exclusion theorem",
		"native S_split transport map",
		"native response-chamber normalization theorem",
	}
	finiteLayerGaps := []string{
		"full A_F descent",
		"lawful spontaneous-orientation interpretation of A_F^orient",
	}

	return Audit{
		ID:             AuditID,
		Truth:          FinalTruth,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		Subobjects: SubobjectStack{
			Gate913ShortStatus: Gate913ShortStatus,
			Gate914ShortStatus: Gate914ShortStatus,
			Gate915ShortStatus: Gate915ShortStatus,
			Gate916ShortStatus: Gate916ShortStatus,
			Gate917ShortStatus: Gate917ShortStatus,
			ResponseShape:      ReducedResponseShape,
			SelectorShape:      DegreeSelectorShape,
			CrossLaneShape:     CrossLaneShape,
			TransportShape:     SSplitTransportShape,
			NormalizationShape: NormalizationShape,
			AllAuditedAtShape:  true,
			NativeTheorem:      false,
			Supports:           []string{StatusFiveSubobjectsInherited},
			Failures:           []string{FailureReconstructionNotNativeAlpha, FailureAlphaBridgeCandidateNotNative},
		},
		BoundaryAlpha: ReassembledBoundaryAlpha{
			Formula:               ConsolidatedFormula,
			RankF1OverF0:          RankF1OverF0,
			RankF2OverF0:          RankF2OverF0,
			RankH10:               RankH10,
			RankH72:               RankH72,
			S:                     SBoundary,
			LinearContribution:    linear,
			QuadraticContribution: quadratic,
			TotalAlpha:            alpha,
			InternalCoherent:      true,
			OpaqueSeal:            false,
			BridgeCandidate:       true,
			NativeTheorem:         false,
			Supports:              []string{SupportDecomposedFunctorCoherent, SupportFiveSubobjectsCompose, SupportAlphaNoLongerOpaque, SupportAlphaBridgeCandidateForm, StatusInternalCoherence, StatusAlphaNoLongerOpaqueSeal},
			Failures:              []string{FailureInternalCoherenceNotNative, FailureReconstructionNotNativeAlpha},
		},
		RepresentativeIndependence: RepresentativeIndependence{
			Domain:                  "[p]_{Z2}={e_lambda tensor P_1, e_barlambda tensor P_1}",
			RankPair:                [2]int{RankF1OverF0, RankF2OverF0},
			TauPhiPreservesRankPair: true,
			PhaseSignEntersAlpha:    false,
			CorrectAlphaDomain:      true,
			NativeAirlockFunctor:    false,
			Supports:                []string{SupportBoundaryAlphaRepresentativeFree, SupportPhaseSignAbsentAfterQuotient, SupportZ2AirlockCorrectDomain, StatusRepresentativeIndependent},
			Failures:                []string{FailureRepresentativeIndependenceNotNative},
		},
		BridgeCandidate: BridgeCandidateStatus{
			PreviousStatus:       "BoundaryAlpha_Z2 class seal",
			NewStatus:            "BoundaryAlpha_Z2 decomposed bridge theorem candidate",
			TheoremObligations:   alphaGaps,
			AllVisibleComponents: true,
			NativeTheorem:        false,
			Supports:             []string{SupportPromotedFromOpaqueSeal, SupportAlphaTheoremObligationLedger, StatusBridgeCandidatePromoted},
			Failures:             []string{FailureNoNativeReducedB2ResponseFunctional, FailureNoNativeDegreeToZ2FlagClassFunctor, FailureNoNativeZ2CrossLaneExclusionTheorem, FailureNoNativeSSplitTransportMap, FailureNoNativeResponseChamberNormalization},
		},
		R3TraceLedger: R3TraceLedgerStatus{
			TraceRows:             []TraceRow{{Rank: 3, Weight: "1"}, {Rank: 3, Weight: "alpha_B(1-alpha_B)"}, {Rank: 1, Weight: "3 alpha_B^2"}},
			ATotalOverT:           aTotal,
			BTotalOverT2:          bTotal,
			NEffOperator:          computedNEff,
			CYukawaOperator:       CYukawaOperator,
			CHiggsOperator:        CHiggsOperator,
			DiagnosticOnly:        true,
			OfficialUpdateAllowed: false,
			Supports:              []string{SupportR3TraceLedgerOnCandidate, SupportOperatorsRemainReconstructed, StatusR3TraceLedgerUnderCandidate},
			Failures:              []string{FailureNotNativeR3, FailureNoOfficialNEffUpdateAllowed, FailureNoCYukawaOrCHiggsUpdateAllowed},
		},
		NativeGaps: NativeGapLedger{
			AlphaGaps:                 alphaGaps,
			FiniteLayerGaps:           finiteLayerGaps,
			GenerationFlavorR4OrLater: true,
			NativeR3:                  false,
			Supports:                  []string{SupportNativeR3BlockersExplicit, StatusNativeGapsExplicit},
			Failures:                  []string{FailureFullAFDescentStillBlocked, FailureHiggsOrientationSealed, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator},
		},
		Firewalls: FirewallLedger{Failures: Failures()},
		Final:     NextGate,
	}, nil
}

func Statuses() []string {
	return []string{StatusFiveSubobjectsInherited, StatusInternalCoherence, StatusAlphaNoLongerOpaqueSeal, StatusRepresentativeIndependent, StatusBridgeCandidatePromoted, StatusR3TraceLedgerUnderCandidate, StatusNativeGapsExplicit, StatusNativeTheoremFirewallsPreserve}
}

func Supports() []string {
	return []string{SupportDecomposedFunctorCoherent, SupportFiveSubobjectsCompose, SupportAlphaNoLongerOpaque, SupportAlphaBridgeCandidateForm, SupportBoundaryAlphaRepresentativeFree, SupportPhaseSignAbsentAfterQuotient, SupportZ2AirlockCorrectDomain, SupportPromotedFromOpaqueSeal, SupportAlphaTheoremObligationLedger, SupportR3TraceLedgerOnCandidate, SupportOperatorsRemainReconstructed, SupportNativeR3BlockersExplicit}
}

func Failures() []string {
	return []string{FailureNotNativeR3, FailureInternalCoherenceNotNative, FailureRepresentativeIndependenceNotNative, FailureNoNativeReducedB2ResponseFunctional, FailureNoNativeDegreeToZ2FlagClassFunctor, FailureNoNativeZ2CrossLaneExclusionTheorem, FailureNoNativeSSplitTransportMap, FailureNoNativeResponseChamberNormalization, FailureReconstructionNotNativeAlpha, FailureAlphaBridgeCandidateNotNative, FailureFullAFDescentStillBlocked, FailureHiggsOrientationSealed, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator, FailureNoOfficialNEffUpdateAllowed, FailureNoCYukawaOrCHiggsUpdateAllowed}
}

func (f FirewallLedger) List() []string { return append([]string{}, f.Failures...) }

func firewallsOK(f FirewallLedger) bool {
	return !f.NativeR3 && !f.InternalCoherenceNative && !f.RepresentativeIndependenceNative && !f.NativeReducedB2Response && !f.NativeDegreeFunctor && !f.NativeCrossLane && !f.NativeSSplitTransport && !f.NativeNormalization && !f.ReconstructionNativeAlpha && !f.AlphaNative && !f.FullAFDescent && !f.HiggsOrientationNative && !f.GenerationCarrier && !f.FlavorOrientation && !f.IndividualYukawa && !f.NativeYukawa && !f.OfficialNEffUpdate && !f.OfficialCoefficientUpdate && containsAll(f.Failures, Failures())
}

func FormatSubobjects(x SubobjectStack) string {
	return fmt.Sprintf("subobjects={913:%s 914:%s 915:%s 916:%s 917:%s allShape:%t native:%t response:%s selector:%s crosslane:%s transport:%s normalization:%s supports:%s failures:%s}", x.Gate913ShortStatus, x.Gate914ShortStatus, x.Gate915ShortStatus, x.Gate916ShortStatus, x.Gate917ShortStatus, x.AllAuditedAtShape, x.NativeTheorem, x.ResponseShape, x.SelectorShape, x.CrossLaneShape, x.TransportShape, x.NormalizationShape, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}

func FormatBoundaryAlpha(x ReassembledBoundaryAlpha) string {
	return fmt.Sprintf("boundaryAlpha={formula:%s ranks:%d,%d denominators:%d,%d s:%g linear:%g quadratic:%g total:%g coherent:%t opaqueSeal:%t bridgeCandidate:%t native:%t supports:%s failures:%s}", x.Formula, x.RankF1OverF0, x.RankF2OverF0, x.RankH10, x.RankH72, x.S, x.LinearContribution, x.QuadraticContribution, x.TotalAlpha, x.InternalCoherent, x.OpaqueSeal, x.BridgeCandidate, x.NativeTheorem, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}

func FormatRepresentativeIndependence(x RepresentativeIndependence) string {
	return fmt.Sprintf("representativeIndependence={domain:%s rankPair:%v tauPhiPreserves:%t phaseSignEntersAlpha:%t correctDomain:%t nativeAirlock:%t supports:%s failures:%s}", x.Domain, x.RankPair, x.TauPhiPreservesRankPair, x.PhaseSignEntersAlpha, x.CorrectAlphaDomain, x.NativeAirlockFunctor, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}

func FormatBridgeCandidate(x BridgeCandidateStatus) string {
	return fmt.Sprintf("bridgeCandidate={previous:%s new:%s obligations:%s allVisible:%t native:%t supports:%s failures:%s}", x.PreviousStatus, x.NewStatus, strings.Join(x.TheoremObligations, ";"), x.AllVisibleComponents, x.NativeTheorem, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}

func FormatR3TraceLedger(x R3TraceLedgerStatus) string {
	rows := make([]string, 0, len(x.TraceRows))
	for _, r := range x.TraceRows {
		rows = append(rows, fmt.Sprintf("(%d,%s)", r.Rank, r.Weight))
	}
	return fmt.Sprintf("r3TraceLedger={rows:%s aTotal:%g bTotal:%g nEff:%g cYukawa:%g cHiggs:%g diagnosticOnly:%t officialUpdate:%t supports:%s failures:%s}", strings.Join(rows, ";"), x.ATotalOverT, x.BTotalOverT2, x.NEffOperator, x.CYukawaOperator, x.CHiggsOperator, x.DiagnosticOnly, x.OfficialUpdateAllowed, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}

func FormatNativeGaps(x NativeGapLedger) string {
	return fmt.Sprintf("nativeGaps={alpha:%s finiteLayer:%s r4orLater:%t nativeR3:%t supports:%s failures:%s}", strings.Join(x.AlphaGaps, ";"), strings.Join(x.FiniteLayerGaps, ";"), x.GenerationFlavorR4OrLater, x.NativeR3, strings.Join(x.Supports, ","), strings.Join(x.Failures, ","))
}

func FormatFirewalls(f FirewallLedger) string {
	return fmt.Sprintf("firewalls={nativeR3:%t internalNative:%t reprNative:%t reducedB2:%t degreeFunctor:%t crossLane:%t sTransport:%t normalization:%t alphaNative:%t fullAF:%t higgsOrientation:%t generation:%t flavor:%t individualYukawa:%t nativeYukawa:%t officialNEff:%t officialCoeff:%t failures:%s}", f.NativeR3, f.InternalCoherenceNative, f.RepresentativeIndependenceNative, f.NativeReducedB2Response, f.NativeDegreeFunctor, f.NativeCrossLane, f.NativeSSplitTransport, f.NativeNormalization, f.AlphaNative, f.FullAFDescent, f.HiggsOrientationNative, f.GenerationCarrier, f.FlavorOrientation, f.IndividualYukawa, f.NativeYukawa, f.OfficialNEffUpdate, f.OfficialCoefficientUpdate, strings.Join(f.Failures, ","))
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

func near(a, b float64) bool      { return math.Abs(a-b) <= 1e-15 }
func nearLoose(a, b float64) bool { return math.Abs(a-b) <= 1e-12 }
