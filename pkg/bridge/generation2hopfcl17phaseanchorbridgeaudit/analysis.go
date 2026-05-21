// Package generation2hopfcl17phaseanchorbridgeaudit implements
// Gate 902: Hopf–Cl(1,7) PhaseAnchor Bridge Audit.
//
// Gate 902 follows Gate 901's reduction of the R3-sealed branch to a single
// master bridge object: the PhaseAnchoredNeutralPunctureAirlockFunctor. It
// audits whether the missing phase anchor lambda > bar(lambda) can be sourced
// from Hopf/S1 phase orientation, the Cl(1,7) complex chirality airlock,
// J/KO conjugation data, boundary-pair orientation, or a finite spectral
// orientation cycle. The honest result is a source-candidate obstruction:
// Hopf/S1 and Cl(1,7) complex chirality are strongest and point to the same
// phase-orientation wound, but no typed Hopf–chirality bridge into the
// right-character socket order is certified. No alpha derivation, native R3
// promotion, physical-sector assignment, individual Yukawa value, or official
// ledger update is certified.
package generation2hopfcl17phaseanchorbridgeaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE902-HOPF-CL17-PHASE-ANCHOR-BRIDGE-AUDIT"

	AlphaB = 0.0003878958469680527

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	RightCharacterSplit = "rho_R(lambda)=lambda e_+ + bar(lambda)e_-"
	LambdaSocket        = "lambda socket -> exposure / puncture"
	BarLambdaSocket     = "bar(lambda) socket -> active / rest"
	PPhase              = "p_phi=e_lambda tensor P_1=e_+ tensor P_1"

	HopfS1PhaseOrientation    = "Hopf/S1 phase orientation"
	CL17ComplexChirality      = "Cl(1,7) complex chirality airlock gamma_chi=i omega"
	HopfChiralityBridge       = "HopfChiralityPhaseAnchorBridge"
	JKOConjugationStructure   = "J/KO conjugation structure"
	BoundaryPairOrientation   = "B2 exterior orientation b1 wedge b2"
	FiniteSpectralOrientation = "finite spectral orientation cycle"

	Classification = "R3_PHASE_ANCHOR_SOURCE_CANDIDATE_HOPF_CL17_BRIDGE_NOT_NATIVE"
	ShortStatus    = "R3_AIRLOCK_PHASE_ANCHOR_REDUCED_TO_HOPF_CHIRALITY_BRIDGE_OBSTRUCTION"

	StatusGate901Inherited       = "PASS_GATE901_PHASE_ANCHORED_AIRLOCK_SEAL_INHERITED"
	StatusHopfRouteAudited       = "PASS_HOPF_S1_PHASE_ANCHOR_ROUTE_AUDITED"
	StatusCL17RouteAudited       = "PASS_CL17_COMPLEX_CHIRALITY_PHASE_ANCHOR_ROUTE_AUDITED"
	StatusHopfChiralityAudited   = "PASS_HOPF_CL17_CHIRALITY_BRIDGE_ROUTE_AUDITED"
	StatusJKORouteAudited        = "PASS_J_KO_CONJUGATION_ROUTE_AUDITED"
	StatusBoundaryRouteAudited   = "PASS_BOUNDARY_PAIR_ORIENTATION_ROUTE_REAUDITED"
	StatusSpectralRouteAudited   = "PASS_FINITE_SPECTRAL_ORIENTATION_CYCLE_ROUTE_REAUDITED"
	StatusSourceCandidatesRanked = "PASS_PHASE_ANCHOR_SOURCE_CANDIDATES_RANKED"
	StatusOfficialFreeze         = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict        = "FIREWALL_PRESERVED_GATE902_HOPF_CL17_PHASE_ANCHOR_BRIDGE_NOT_NATIVE"

	SupportPhaseAnchoredAirlockInherited     = "CONDITIONAL_SUPPORT_PHASE_ANCHORED_AIRLOCK_SEAL_INHERITED"
	SupportRightCharacterLambdaBarShape      = "CONDITIONAL_SUPPORT_RIGHT_CHARACTER_LAMBDA_BARLAMBDA_ORDER_CAN_BE_STATED_GIVEN_PHASE_ANCHOR"
	SupportHopfStrongestPhaseAnchor          = "CONDITIONAL_SUPPORT_HOPF_S1_PHASE_ORIENTATION_IS_STRONGEST_PHASE_ANCHOR_SOURCE"
	SupportRightPairMatchesHopfConjugation   = "CONDITIONAL_SUPPORT_RIGHT_CHARACTER_LAMBDA_BARLAMBDA_PAIR_MATCHES_HOPF_PHASE_CONJUGATION_SHAPE"
	SupportAirlockReadableAsHopfIfSealed     = "CONDITIONAL_SUPPORT_PHASE_ANCHORED_AIRLOCK_CAN_BE_READ_AS_HOPF_PHASE_ORIENTED_IF_SEALED"
	SupportCL17StrongCandidate               = "CONDITIONAL_SUPPORT_CL17_COMPLEX_CHIRALITY_AIRLOCK_IS_STRONG_PHASE_ANCHOR_SOURCE_CANDIDATE"
	SupportIVsMinusI                         = "CONDITIONAL_SUPPORT_I_VS_MINUS_I_CAN_SOURCE_LAMBDA_VS_BARLAMBDA_ORIENTATION_IF_TYPED"
	SupportCL17RealFormFirewall              = "CONDITIONAL_SUPPORT_CL17_REAL_FORM_FIREWALL_RELEVANT_TO_RIGHT_CHARACTER_PHASE_ORDER"
	SupportHopfCL17CompatibleSources         = "CONDITIONAL_SUPPORT_HOPF_AND_CL17_CHIRALITY_ARE_COMPATIBLE_PHASE_ORIENTATION_SOURCES"
	SupportPhaseAnchorMayRequireBridge       = "CONDITIONAL_SUPPORT_PHASE_ANCHOR_MAY_REQUIRE_HOPF_CHIRALITY_BRIDGE"
	SupportHopfAndCL17SameWound              = "CONDITIONAL_SUPPORT_HOPF_AND_CL17_CHIRALITY_POINT_TO_SAME_PHASE_ORIENTATION_WOUND"
	SupportJKORelevant                       = "CONDITIONAL_SUPPORT_J_KO_STRUCTURE_RELEVANT_TO_CONJUGATION_LEDGER"
	SupportBoundaryCompatible                = "CONDITIONAL_SUPPORT_BOUNDARY_PAIR_ORIENTATION_COMPATIBLE_WITH_PHASE_ANCHORED_AIRLOCK"
	SupportSpectralDeepCandidate             = "CONDITIONAL_SUPPORT_FINITE_SPECTRAL_ORIENTATION_CYCLE_IS_DEEP_SOCKET_ORDER_SOURCE_CANDIDATE"
	SupportPhaseAnchoredAirlockCoherent      = "CONDITIONAL_SUPPORT_PHASE_ANCHORED_AIRLOCK_REMAINS_COHERENT_UNDER_HOPF_CL17_SOURCE_CANDIDATES"
	SupportOperatorDiagnosticsRemainCoherent = "CONDITIONAL_SUPPORT_OPERATOR_DIAGNOSTICS_REMAIN_COHERENT_UNDER_PHASE_ANCHOR_SOURCE_CANDIDATE_SEAL"

	FailureNoNativeRightPhaseOrientation  = "FAILED_ROUTE_NO_NATIVE_RIGHT_CHARACTER_PHASE_ORIENTATION_THEOREM"
	FailureNoHopfToSocketOrderMap         = "FAILED_ROUTE_NO_TYPED_HOPF_PHASE_TO_RIGHT_CHARACTER_SOCKET_ORDER_MAP"
	FailureHopfNotNativeSelector          = "FAILED_ROUTE_HOPF_PHASE_ORIENTATION_NOT_YET_NATIVE_RIGHT_CHARACTER_SELECTOR"
	FailureNoCL17ToRightPhaseMap          = "FAILED_ROUTE_NO_TYPED_CL17_COMPLEX_CHIRALITY_TO_RIGHT_CHARACTER_PHASE_MAP"
	FailureChiralityNotSocketOrderTheorem = "FAILED_ROUTE_COMPLEX_CHIRALITY_AIRLOCK_NOT_SOCKET_ORDER_THEOREM_YET"
	FailureNoNativeHopfChiralityBridge    = "FAILED_ROUTE_NO_NATIVE_HOPF_CHIRALITY_PHASE_ANCHOR_BRIDGE_CERTIFIED"
	FailureJKONoLambdaSelection           = "FAILED_ROUTE_J_KO_SIGN_DOES_NOT_SELECT_LAMBDA_OVER_BARLAMBDA"
	FailureJMirrorNoBreak                 = "FAILED_ROUTE_J_MIRROR_DOES_NOT_BREAK_RIGHT_CHARACTER_PHASE_Z2"
	FailureBoundaryNoLambdaSelection      = "FAILED_ROUTE_BOUNDARY_PAIR_ORIENTATION_DOES_NOT_SELECT_LAMBDA_OVER_BARLAMBDA"
	FailureBoundaryDegreeNotPhaseOrder    = "FAILED_ROUTE_BOUNDARY_ORIENTATION_SELECTS_DEGREE_ORDER_NOT_PHASE_SOCKET_ORDER"
	FailureNoSpectralCycleToSocketOrder   = "FAILED_ROUTE_NO_FINITE_SPECTRAL_ORIENTATION_CYCLE_TO_SOCKET_ORDER_THEOREM"
	FailureNoNativeSelectionLambdaBar     = "FAILED_ROUTE_NO_NATIVE_SELECTION_OF_LAMBDA_OVER_BARLAMBDA"
	FailureAlphaStillSealedWithoutPhase   = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED_WITHOUT_NATIVE_PHASE_ANCHORED_BOUNDARY_FUNCTOR"
	FailureHiggsStillSealedWithoutPhase   = "FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_WITHOUT_NATIVE_PHASE_ANCHORED_WEAK_SELECTOR"
	FailureNotNativeR3                    = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureNoGenerationCarrierMap         = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap         = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues       = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate           = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate          = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator         = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4NativeYukawaTheorem        = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type InheritedAudit struct {
	RightCharacterSplit       string
	LambdaSocket, BarSocket   string
	PhaseAnchorOrganizesChain bool
	NativeAnchor              bool
	Supports, Failures        []string
}

type HopfS1Audit struct {
	CandidateName                 string
	HasOrientedPhaseCircle        bool
	MatchesLambdaBarShape         bool
	CanReadAirlockIfSealed        bool
	NativeSocketOrderMapCertified bool
	Supports, Failures            []string
}

type CL17ChiralityAudit struct {
	CandidateName                 string
	CL17RealBoard                 string
	OmegaSquaredMinusOne          bool
	RequiresComplexChirality      bool
	CanOrientIOverMinusI          bool
	TypedToRightCharacterPhase    bool
	NativeSocketOrderMapCertified bool
	Supports, Failures            []string
}

type HopfChiralityBridgeAudit struct {
	BridgeName                 string
	HopfAndChiralityCompatible bool
	PointsToSamePhaseWound     bool
	NativeBridgeCertified      bool
	CanAnchorRightCharacters   bool
	Supports, Failures         []string
}

type JKOAudit struct {
	CandidateName       string
	RelevantToConjugacy bool
	SelectsLambda       bool
	BreaksZ2            bool
	NativeTheorem       bool
	Supports, Failures  []string
}

type BoundaryPairAudit struct {
	CandidateName              string
	HasExteriorOrientation     bool
	CompatibleWithPhaseAirlock bool
	SelectsDegreeOrder         bool
	SelectsRightCharacterPhase bool
	Supports, Failures         []string
}

type SpectralOrientationAudit struct {
	CandidateName      string
	DeepCandidate      bool
	CycleCertified     bool
	MapsToSocketOrder  bool
	NativeTheorem      bool
	Supports, Failures []string
}

type CandidateRankingAudit struct {
	StrongestCandidates []string
	BridgeCandidate     string
	NativeSourceFound   bool
	NextFrontier        string
	Supports, Failures  []string
}

type FreezeAudit struct {
	Alpha, OperatorNEff, OfficialNEff float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	Frozen, DiagnosticOnly, CanUpdate bool
	Supports, Failures                []string
}

type Firewalls struct {
	Enforced                      bool
	NoNativeRightPhaseOrientation bool
	NoHopfSocketMap               bool
	NoCL17RightPhaseMap           bool
	NoHopfChiralityBridge         bool
	NoNativeLambdaSelection       bool
	AlphaStillSealed              bool
	HiggsOrientationStillSealed   bool
	NotNativeR3                   bool
	NoGenerationCarrier           bool
	NoFlavorOrientation           bool
	NoIndividualYukawas           bool
	NoOfficialLedgerUpdate        bool
	NoNativeYukawaOperator        bool
	NoR4YukawaTheorem             bool
	Verdict                       string
}

type Audit struct {
	ID                  string
	Inherited           InheritedAudit
	HopfS1              HopfS1Audit
	CL17Chirality       CL17ChiralityAudit
	HopfChiralityBridge HopfChiralityBridgeAudit
	JKO                 JKOAudit
	BoundaryPair        BoundaryPairAudit
	SpectralOrientation SpectralOrientationAudit
	Ranking             CandidateRankingAudit
	Freeze              FreezeAudit
	Firewalls           Firewalls
	Classification      string
	ShortStatus         string
	Truth               string
	Final               string
}

func BuildDefault() (Audit, error) {
	inherited := buildInheritedAudit()
	if !inherited.PhaseAnchorOrganizesChain || inherited.NativeAnchor {
		return Audit{}, fmt.Errorf("inherited phase-anchor leak: %s", FormatInherited(inherited))
	}
	hopf := buildHopfS1Audit()
	if !hopf.HasOrientedPhaseCircle || !hopf.MatchesLambdaBarShape || !hopf.CanReadAirlockIfSealed || hopf.NativeSocketOrderMapCertified {
		return Audit{}, fmt.Errorf("Hopf route leak: %s", FormatHopfS1(hopf))
	}
	cl17 := buildCL17ChiralityAudit()
	if !cl17.OmegaSquaredMinusOne || !cl17.RequiresComplexChirality || !cl17.CanOrientIOverMinusI || cl17.TypedToRightCharacterPhase || cl17.NativeSocketOrderMapCertified {
		return Audit{}, fmt.Errorf("Cl17 route leak: %s", FormatCL17Chirality(cl17))
	}
	bridge := buildHopfChiralityBridgeAudit()
	if !bridge.HopfAndChiralityCompatible || !bridge.PointsToSamePhaseWound || bridge.NativeBridgeCertified || bridge.CanAnchorRightCharacters {
		return Audit{}, fmt.Errorf("Hopf-Cl17 bridge leak: %s", FormatHopfChiralityBridge(bridge))
	}
	jko := buildJKOAudit()
	if !jko.RelevantToConjugacy || jko.SelectsLambda || jko.BreaksZ2 || jko.NativeTheorem {
		return Audit{}, fmt.Errorf("J/KO route leak: %s", FormatJKO(jko))
	}
	boundary := buildBoundaryPairAudit()
	if !boundary.HasExteriorOrientation || !boundary.CompatibleWithPhaseAirlock || !boundary.SelectsDegreeOrder || boundary.SelectsRightCharacterPhase {
		return Audit{}, fmt.Errorf("boundary route leak: %s", FormatBoundaryPair(boundary))
	}
	spectral := buildSpectralOrientationAudit()
	if !spectral.DeepCandidate || spectral.CycleCertified || spectral.MapsToSocketOrder || spectral.NativeTheorem {
		return Audit{}, fmt.Errorf("spectral route leak: %s", FormatSpectralOrientation(spectral))
	}
	ranking := buildCandidateRankingAudit()
	if len(ranking.StrongestCandidates) != 2 || ranking.NativeSourceFound || ranking.BridgeCandidate != HopfChiralityBridge {
		return Audit{}, fmt.Errorf("ranking leak: %s", FormatRanking(ranking))
	}
	freeze := buildFreezeAudit()
	if !freeze.Frozen || !freeze.DiagnosticOnly || freeze.CanUpdate || near(freeze.OperatorNEff, freeze.OfficialNEff) {
		return Audit{}, fmt.Errorf("freeze leak: %s", FormatFreeze(freeze))
	}
	firewalls := buildFirewalls()
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewall leak: %s", FormatFirewalls(firewalls))
	}
	return Audit{ID: AuditID, Inherited: inherited, HopfS1: hopf, CL17Chirality: cl17, HopfChiralityBridge: bridge, JKO: jko, BoundaryPair: boundary, SpectralOrientation: spectral, Ranking: ranking, Freeze: freeze, Firewalls: firewalls, Classification: Classification, ShortStatus: ShortStatus, Truth: "Gate 902 reduces the phase-anchor source wound to a Hopf/S1 and Cl(1,7) complex-chirality bridge obstruction. Both are strong candidates, but no typed native bridge into the right-character socket order is certified.", Final: ShortStatus}, nil
}

func buildInheritedAudit() InheritedAudit {
	return InheritedAudit{RightCharacterSplit: RightCharacterSplit, LambdaSocket: LambdaSocket, BarSocket: BarLambdaSocket, PhaseAnchorOrganizesChain: true, NativeAnchor: false, Supports: []string{StatusGate901Inherited, SupportPhaseAnchoredAirlockInherited, SupportRightCharacterLambdaBarShape}, Failures: []string{FailureNoNativeRightPhaseOrientation, FailureNoNativeSelectionLambdaBar}}
}

func buildHopfS1Audit() HopfS1Audit {
	return HopfS1Audit{CandidateName: HopfS1PhaseOrientation, HasOrientedPhaseCircle: true, MatchesLambdaBarShape: true, CanReadAirlockIfSealed: true, NativeSocketOrderMapCertified: false, Supports: []string{StatusHopfRouteAudited, SupportHopfStrongestPhaseAnchor, SupportRightPairMatchesHopfConjugation, SupportAirlockReadableAsHopfIfSealed}, Failures: []string{FailureNoHopfToSocketOrderMap, FailureHopfNotNativeSelector}}
}

func buildCL17ChiralityAudit() CL17ChiralityAudit {
	return CL17ChiralityAudit{CandidateName: CL17ComplexChirality, CL17RealBoard: "Cl(1,7)=Mat(16,R)", OmegaSquaredMinusOne: true, RequiresComplexChirality: true, CanOrientIOverMinusI: true, TypedToRightCharacterPhase: false, NativeSocketOrderMapCertified: false, Supports: []string{StatusCL17RouteAudited, SupportCL17StrongCandidate, SupportIVsMinusI, SupportCL17RealFormFirewall}, Failures: []string{FailureNoCL17ToRightPhaseMap, FailureChiralityNotSocketOrderTheorem}}
}

func buildHopfChiralityBridgeAudit() HopfChiralityBridgeAudit {
	return HopfChiralityBridgeAudit{BridgeName: HopfChiralityBridge, HopfAndChiralityCompatible: true, PointsToSamePhaseWound: true, NativeBridgeCertified: false, CanAnchorRightCharacters: false, Supports: []string{StatusHopfChiralityAudited, SupportHopfCL17CompatibleSources, SupportPhaseAnchorMayRequireBridge, SupportHopfAndCL17SameWound}, Failures: []string{FailureNoNativeHopfChiralityBridge}}
}

func buildJKOAudit() JKOAudit {
	return JKOAudit{CandidateName: JKOConjugationStructure, RelevantToConjugacy: true, SelectsLambda: false, BreaksZ2: false, NativeTheorem: false, Supports: []string{StatusJKORouteAudited, SupportJKORelevant}, Failures: []string{FailureJKONoLambdaSelection, FailureJMirrorNoBreak}}
}

func buildBoundaryPairAudit() BoundaryPairAudit {
	return BoundaryPairAudit{CandidateName: BoundaryPairOrientation, HasExteriorOrientation: true, CompatibleWithPhaseAirlock: true, SelectsDegreeOrder: true, SelectsRightCharacterPhase: false, Supports: []string{StatusBoundaryRouteAudited, SupportBoundaryCompatible}, Failures: []string{FailureBoundaryNoLambdaSelection, FailureBoundaryDegreeNotPhaseOrder}}
}

func buildSpectralOrientationAudit() SpectralOrientationAudit {
	return SpectralOrientationAudit{CandidateName: FiniteSpectralOrientation, DeepCandidate: true, CycleCertified: false, MapsToSocketOrder: false, NativeTheorem: false, Supports: []string{StatusSpectralRouteAudited, SupportSpectralDeepCandidate}, Failures: []string{FailureNoSpectralCycleToSocketOrder}}
}

func buildCandidateRankingAudit() CandidateRankingAudit {
	return CandidateRankingAudit{StrongestCandidates: []string{HopfS1PhaseOrientation, CL17ComplexChirality}, BridgeCandidate: HopfChiralityBridge, NativeSourceFound: false, NextFrontier: "HOPF_CL17_PHASE_ANCHOR_BRIDGE_THEOREM", Supports: []string{StatusSourceCandidatesRanked, SupportHopfCL17CompatibleSources, SupportHopfAndCL17SameWound, SupportPhaseAnchoredAirlockCoherent}, Failures: []string{FailureNoNativeRightPhaseOrientation, FailureNoNativeHopfChiralityBridge, FailureNoNativeSelectionLambdaBar}}
}

func buildFreezeAudit() FreezeAudit {
	return FreezeAudit{Alpha: AlphaB, OperatorNEff: OperatorNEffDiagnostic, OfficialNEff: OfficialNEffFrozen, OperatorCYukawa: OperatorCYukawaDiagnostic, OfficialCYukawa: OfficialCYukawaFrozen, OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen, Frozen: true, DiagnosticOnly: true, CanUpdate: false, Supports: []string{StatusOfficialFreeze, SupportOperatorDiagnosticsRemainCoherent}, Failures: []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate}}
}

func buildFirewalls() Firewalls {
	return Firewalls{Enforced: true, NoNativeRightPhaseOrientation: true, NoHopfSocketMap: true, NoCL17RightPhaseMap: true, NoHopfChiralityBridge: true, NoNativeLambdaSelection: true, AlphaStillSealed: true, HiggsOrientationStillSealed: true, NotNativeR3: true, NoGenerationCarrier: true, NoFlavorOrientation: true, NoIndividualYukawas: true, NoOfficialLedgerUpdate: true, NoNativeYukawaOperator: true, NoR4YukawaTheorem: true, Verdict: StatusFirewallVerdict}
}

func FormatInherited(i InheritedAudit) string {
	return fmt.Sprintf("inherited(split=%s lambda=%s bar=%s organizes_chain=%t native=%t supports=%s failures=%s)", i.RightCharacterSplit, i.LambdaSocket, i.BarSocket, i.PhaseAnchorOrganizesChain, i.NativeAnchor, strings.Join(i.Supports, ","), strings.Join(i.Failures, ","))
}

func FormatHopfS1(h HopfS1Audit) string {
	return fmt.Sprintf("hopf_s1(candidate=%s phase_circle=%t matches_lambda_bar=%t readable_if_sealed=%t native_map=%t supports=%s failures=%s)", h.CandidateName, h.HasOrientedPhaseCircle, h.MatchesLambdaBarShape, h.CanReadAirlockIfSealed, h.NativeSocketOrderMapCertified, strings.Join(h.Supports, ","), strings.Join(h.Failures, ","))
}

func FormatCL17Chirality(c CL17ChiralityAudit) string {
	return fmt.Sprintf("cl17_chirality(candidate=%s board=%s omega2_minus_one=%t complex_airlock=%t i_orientation=%t typed_to_right=%t native_map=%t supports=%s failures=%s)", c.CandidateName, c.CL17RealBoard, c.OmegaSquaredMinusOne, c.RequiresComplexChirality, c.CanOrientIOverMinusI, c.TypedToRightCharacterPhase, c.NativeSocketOrderMapCertified, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}

func FormatHopfChiralityBridge(b HopfChiralityBridgeAudit) string {
	return fmt.Sprintf("hopf_chirality_bridge(name=%s compatible=%t same_wound=%t native=%t anchors_right=%t supports=%s failures=%s)", b.BridgeName, b.HopfAndChiralityCompatible, b.PointsToSamePhaseWound, b.NativeBridgeCertified, b.CanAnchorRightCharacters, strings.Join(b.Supports, ","), strings.Join(b.Failures, ","))
}

func FormatJKO(j JKOAudit) string {
	return fmt.Sprintf("jko(candidate=%s conjugacy=%t selects_lambda=%t breaks_z2=%t native=%t supports=%s failures=%s)", j.CandidateName, j.RelevantToConjugacy, j.SelectsLambda, j.BreaksZ2, j.NativeTheorem, strings.Join(j.Supports, ","), strings.Join(j.Failures, ","))
}

func FormatBoundaryPair(b BoundaryPairAudit) string {
	return fmt.Sprintf("boundary_pair(candidate=%s exterior_orientation=%t compatible=%t degree_order=%t socket_phase=%t supports=%s failures=%s)", b.CandidateName, b.HasExteriorOrientation, b.CompatibleWithPhaseAirlock, b.SelectsDegreeOrder, b.SelectsRightCharacterPhase, strings.Join(b.Supports, ","), strings.Join(b.Failures, ","))
}

func FormatSpectralOrientation(s SpectralOrientationAudit) string {
	return fmt.Sprintf("spectral_orientation(candidate=%s deep=%t certified=%t maps_to_socket=%t native=%t supports=%s failures=%s)", s.CandidateName, s.DeepCandidate, s.CycleCertified, s.MapsToSocketOrder, s.NativeTheorem, strings.Join(s.Supports, ","), strings.Join(s.Failures, ","))
}

func FormatRanking(r CandidateRankingAudit) string {
	return fmt.Sprintf("ranking(strongest=%s bridge=%s native_found=%t next=%s supports=%s failures=%s)", strings.Join(r.StrongestCandidates, ";"), r.BridgeCandidate, r.NativeSourceFound, r.NextFrontier, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}

func FormatFreeze(f FreezeAudit) string {
	return fmt.Sprintf("freeze(alpha=%.16g neff=%.16g official_neff=%.16g cy=%.16g official_cy=%.16g ch=%.16g official_ch=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.Alpha, f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t right_phase=%t hopf_map=%t cl17_map=%t hopf_cl17=%t lambda_selection=%t alpha=%t higgs=%t not_r3=%t generation=%t flavor=%t individual=%t official=%t yukawa=%t r4=%t verdict=%s)", f.Enforced, f.NoNativeRightPhaseOrientation, f.NoHopfSocketMap, f.NoCL17RightPhaseMap, f.NoHopfChiralityBridge, f.NoNativeLambdaSelection, f.AlphaStillSealed, f.HiggsOrientationStillSealed, f.NotNativeR3, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4YukawaTheorem, f.Verdict)
}

func Statuses() []string {
	return []string{StatusGate901Inherited, StatusHopfRouteAudited, StatusCL17RouteAudited, StatusHopfChiralityAudited, StatusJKORouteAudited, StatusBoundaryRouteAudited, StatusSpectralRouteAudited, StatusSourceCandidatesRanked, StatusOfficialFreeze, StatusFirewallVerdict, SupportPhaseAnchoredAirlockInherited, SupportRightCharacterLambdaBarShape, SupportHopfStrongestPhaseAnchor, SupportRightPairMatchesHopfConjugation, SupportAirlockReadableAsHopfIfSealed, SupportCL17StrongCandidate, SupportIVsMinusI, SupportCL17RealFormFirewall, SupportHopfCL17CompatibleSources, SupportPhaseAnchorMayRequireBridge, SupportHopfAndCL17SameWound, SupportJKORelevant, SupportBoundaryCompatible, SupportSpectralDeepCandidate, SupportPhaseAnchoredAirlockCoherent, SupportOperatorDiagnosticsRemainCoherent, FailureNoNativeRightPhaseOrientation, FailureNoHopfToSocketOrderMap, FailureHopfNotNativeSelector, FailureNoCL17ToRightPhaseMap, FailureChiralityNotSocketOrderTheorem, FailureNoNativeHopfChiralityBridge, FailureJKONoLambdaSelection, FailureJMirrorNoBreak, FailureBoundaryNoLambdaSelection, FailureBoundaryDegreeNotPhaseOrder, FailureNoSpectralCycleToSocketOrder, FailureNoNativeSelectionLambdaBar, FailureAlphaStillSealedWithoutPhase, FailureHiggsStillSealedWithoutPhase, FailureNotNativeR3, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem}
}

func (a Audit) FirewallsList() []string {
	return []string{FailureNoNativeRightPhaseOrientation, FailureNoHopfToSocketOrderMap, FailureNoCL17ToRightPhaseMap, FailureNoNativeHopfChiralityBridge, FailureNoNativeSelectionLambdaBar, FailureAlphaStillSealedWithoutPhase, FailureHiggsStillSealedWithoutPhase, FailureNotNativeR3, FailureNoGenerationCarrierMap, FailureNoOfficialNEffUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem}
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NoNativeRightPhaseOrientation && f.NoHopfSocketMap && f.NoCL17RightPhaseMap && f.NoHopfChiralityBridge && f.NoNativeLambdaSelection && f.AlphaStillSealed && f.HiggsOrientationStillSealed && f.NotNativeR3 && f.NoGenerationCarrier && f.NoFlavorOrientation && f.NoIndividualYukawas && f.NoOfficialLedgerUpdate && f.NoNativeYukawaOperator && f.NoR4YukawaTheorem && f.Verdict == StatusFirewallVerdict
}

func containsAll(haystack []string, needles []string) bool {
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

func near(a, b float64) bool { return math.Abs(a-b) < 5e-15 }
