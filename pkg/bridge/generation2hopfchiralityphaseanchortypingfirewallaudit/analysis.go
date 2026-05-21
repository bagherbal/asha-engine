// Package generation2hopfchiralityphaseanchortypingfirewallaudit implements
// Gate 903: HopfChirality PhaseAnchor Typing and Firewall Audit.
//
// Gate 903 follows Gate 902's reduction of the phase-anchor source wound to a
// Hopf/S1 and Cl(1,7) complex-chirality bridge obstruction. It audits whether
// the Hopf phase circle and the Cl(1,7) complex chirality airlock can be typed
// as the same phase-orientation source for the right-character order
// lambda > bar(lambda). The honest result is a shape-support obstruction:
// Hopf phase, Cl(1,7) complex chirality, and the right-character split share a
// conjugate-pair phase shape, but no typed transport into the socket order is
// certified. No alpha derivation, native R3 promotion, physical-sector
// assignment, individual Yukawa value, or official ledger update is certified.
package generation2hopfchiralityphaseanchortypingfirewallaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE903-HOPF-CHIRALITY-PHASE-ANCHOR-TYPING-FIREWALL-AUDIT"

	AlphaB = 0.0003878958469680527

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	RightCharacterSplit = "rho_R(lambda)=lambda e_+ + bar(lambda)e_-"
	HopfPhaseCircle     = "Hopf/S1 phase orientation"
	CL17Chirality       = "Cl(1,7) complex chirality airlock gamma_chi=i omega"
	PhaseTransportMap   = "HopfChiralityRightCharacterTransportMap"

	Classification = "R3_PHASE_ANCHOR_HOPF_CHIRALITY_TYPING_OBSTRUCTION"
	ShortStatus    = "R3_AIRLOCK_PHASE_ANCHOR_SHAPE_SUPPORTED_TRANSPORT_MISSING"

	StatusGate902Inherited          = "PASS_GATE902_HOPF_CL17_BRIDGE_OBSTRUCTION_INHERITED"
	StatusHopfTypingAudited         = "PASS_HOPF_PHASE_TO_RIGHT_CHARACTER_TYPING_AUDITED"
	StatusCL17TypingAudited         = "PASS_CL17_CHIRALITY_TO_PHASE_TYPING_AUDITED"
	StatusAlignmentAudited          = "PASS_HOPF_CHIRALITY_ALIGNMENT_AUDITED"
	StatusFirewallAudited           = "PASS_SYMBOLIC_SHAPE_VS_TYPED_TRANSPORT_FIREWALL_AUDITED"
	StatusTransportObjectIdentified = "PASS_PHASE_TRANSPORT_MAP_IDENTIFIED_AS_NEXT_MISSING_OBJECT"
	StatusOfficialFreeze            = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict           = "FIREWALL_PRESERVED_GATE903_SHAPE_SUPPORT_BUT_TRANSPORT_MISSING"

	SupportGate902Inherited            = "CONDITIONAL_SUPPORT_GATE902_PHASE_ANCHOR_SOURCE_WOUND_INHERITED"
	SupportHopfHasRightShape           = "CONDITIONAL_SUPPORT_HOPF_PHASE_HAS_RIGHT_CHARACTER_CONJUGATION_SHAPE"
	SupportHopfPositiveLabelsEPlus     = "CONDITIONAL_SUPPORT_HOPF_POSITIVE_PHASE_CAN_LABEL_E_PLUS_IF_PHASE_ANCHOR_SEALED"
	SupportHopfConjugateLabelsEMinus   = "CONDITIONAL_SUPPORT_HOPF_CONJUGATE_PHASE_CAN_LABEL_E_MINUS_IF_PHASE_ANCHOR_SEALED"
	SupportCL17HasConjugationShape     = "CONDITIONAL_SUPPORT_CL17_COMPLEX_CHIRALITY_HAS_PHASE_CONJUGATION_SHAPE"
	SupportCL17IOrientation            = "CONDITIONAL_SUPPORT_CL17_COMPLEX_CHIRALITY_SUPPLIES_I_VS_MINUS_I_ORIENTATION_CANDIDATE"
	SupportIVsMinusIShape              = "CONDITIONAL_SUPPORT_I_VS_MINUS_I_HAS_CORRECT_PHASE_CONJUGATION_SHAPE"
	SupportHopfCL17Align               = "CONDITIONAL_SUPPORT_HOPF_AND_CL17_CHIRALITY_ALIGN_AS_PHASE_ORIENTATION_CANDIDATES"
	SupportCompatibleOrientationTypes  = "CONDITIONAL_SUPPORT_HOPF_PHASE_AND_CL17_CHIRALITY_HAVE_COMPATIBLE_PHASE_ORIENTATION_TYPE"
	SupportAlignmentStrongestCandidate = "CONDITIONAL_SUPPORT_HOPF_CHIRALITY_ALIGNMENT_IS_STRONGEST_PHASE_ANCHOR_BRIDGE_CANDIDATE"
	SupportRightAnchorIfTransportTyped = "CONDITIONAL_SUPPORT_RIGHT_CHARACTER_PHASE_ANCHOR_COULD_BE_SOURCED_IF_ALIGNMENT_MAP_CERTIFIED"
	SupportSourceSharpenedToTransport  = "CONDITIONAL_SUPPORT_PHASE_ANCHOR_SOURCE_CANDIDATE_SHARPENED_TO_TYPED_TRANSPORT_MAP"
	SupportR3SealReducesToTransport    = "CONDITIONAL_SUPPORT_R3_AIRLOCK_SEAL_REDUCES_TO_PHASE_TRANSPORT_THEOREM"
	SupportDiagnosticsRemainCoherent   = "CONDITIONAL_SUPPORT_OPERATOR_DIAGNOSTICS_REMAIN_COHERENT_UNDER_PHASE_TRANSPORT_SEAL"

	FailureNoHopfToRightCharacterMap     = "FAILED_ROUTE_NO_NATIVE_HOPF_PHASE_TO_RIGHT_CHARACTER_REPRESENTATION_MAP"
	FailureHopfLabelingStillSeal         = "FAILED_ROUTE_HOPF_PHASE_LABELING_STILL_SEAL_WITHOUT_TYPED_ACTION_ON_C_R2"
	FailureNoGammaChiToRightAction       = "FAILED_ROUTE_NO_TYPED_GAMMA_CHI_TO_RIGHT_CHARACTER_ACTION_MAP"
	FailureChiralityDoesNotSelectSocket  = "FAILED_ROUTE_COMPLEX_CHIRALITY_ORIENTATION_DOES_NOT_YET_SELECT_SOCKET_ORDER"
	FailureNoHopfChiralityAlignment      = "FAILED_ROUTE_NO_NATIVE_HOPF_CHIRALITY_ALIGNMENT_MAP"
	FailureNoGammaTransportToRhoR        = "FAILED_ROUTE_NO_TYPED_TRANSPORT_FROM_GAMMA_CHI_ORIENTATION_TO_RHO_R_CHARACTER_ORDER"
	FailurePhaseAnchorStillSealed        = "FAILED_ROUTE_PHASE_ANCHOR_REMAINS_SEALED"
	FailureShapeMatchNotTheorem          = "FAILED_ROUTE_PHASE_SHAPE_MATCH_NOT_PHASE_ANCHOR_THEOREM"
	FailureConjugationResonanceNotMap    = "FAILED_ROUTE_SYMBOLIC_CONJUGATION_RESONANCE_NOT_TYPED_SOCKET_ORDER_MAP"
	FailureNoTypedPhaseTransport         = "FAILED_ROUTE_NO_TYPED_PHASE_TRANSPORT_TO_SOCKET_ORDER"
	FailureNoNativeRightPhaseOrientation = "FAILED_ROUTE_NO_NATIVE_RIGHT_CHARACTER_PHASE_ORIENTATION_THEOREM"
	FailureNoNativeLambdaSelection       = "FAILED_ROUTE_NO_NATIVE_SELECTION_OF_LAMBDA_OVER_BARLAMBDA"
	FailureAlphaStillSealedWithoutPhase  = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED_WITHOUT_NATIVE_PHASE_ANCHORED_BOUNDARY_FUNCTOR"
	FailureHiggsStillSealedWithoutPhase  = "FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_WITHOUT_NATIVE_PHASE_ANCHORED_WEAK_SELECTOR"
	FailureNotNativeR3                   = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureNoGenerationCarrierMap        = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap        = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues      = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate          = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate         = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator        = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4NativeYukawaTheorem       = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type InheritedAudit struct {
	Gate902Classification string
	RightCharacterSplit   string
	PhaseAnchorMissing    bool
	StrongSources         []string
	NativeBridgeCertified bool
	Supports, Failures    []string
}

type HopfTypingAudit struct {
	CandidateName             string
	HasS1Phase                bool
	HasPositiveConjugatePair  bool
	LabelsEPlusIfSealed       bool
	LabelsEMinusIfSealed      bool
	TypedActionOnCR2Certified bool
	NativeRepresentationMap   bool
	Supports, Failures        []string
}

type CL17TypingAudit struct {
	CandidateName           string
	RealBoard               string
	OmegaSquaredMinusOne    bool
	GammaChi                string
	SuppliesIOverMinusI     bool
	CorrectConjugationShape bool
	TypedToRightCharacter   bool
	SelectsSocketOrder      bool
	Supports, Failures      []string
}

type AlignmentAudit struct {
	BridgeName                    string
	CompatiblePhaseTypes          bool
	StrongestBridgeCandidate      bool
	NativeAlignmentMap            bool
	TransportToRhoR               bool
	CanSourceRightCharacterAnchor bool
	Supports, Failures            []string
}

type ShapeFirewallAudit struct {
	SamePhaseShape           bool
	TypedTransportCertified  bool
	ConjugationResonanceOnly bool
	NativeLambdaSelection    bool
	Supports, Failures       []string
}

type TransportTargetAudit struct {
	MissingObject      string
	Domain             string
	Codomain           string
	RequiredAction     string
	Sharpened          bool
	NativeMapCertified bool
	Supports, Failures []string
}

type FreezeAudit struct {
	Alpha, OperatorNEff, OfficialNEff float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	Frozen, DiagnosticOnly, CanUpdate bool
	Supports, Failures                []string
}

type Firewalls struct {
	Enforced                 bool
	NoHopfRightCharacterMap  bool
	NoGammaChiRightAction    bool
	NoHopfChiralityAlignment bool
	NoTypedPhaseTransport    bool
	ShapeMatchNotTheorem     bool
	PhaseAnchorStillSealed   bool
	NoNativeRightPhase       bool
	AlphaStillSealed         bool
	HiggsStillSealed         bool
	NotNativeR3              bool
	NoGenerationCarrier      bool
	NoFlavorOrientation      bool
	NoIndividualYukawas      bool
	NoOfficialLedgerUpdate   bool
	NoNativeYukawaOperator   bool
	NoR4YukawaTheorem        bool
	Verdict                  string
}

type Audit struct {
	ID              string
	Inherited       InheritedAudit
	HopfTyping      HopfTypingAudit
	CL17Typing      CL17TypingAudit
	Alignment       AlignmentAudit
	ShapeFirewall   ShapeFirewallAudit
	TransportTarget TransportTargetAudit
	Freeze          FreezeAudit
	Firewalls       Firewalls
	Classification  string
	ShortStatus     string
	Truth           string
	Final           string
}

func BuildDefault() (Audit, error) {
	inherited := buildInheritedAudit()
	if !inherited.PhaseAnchorMissing || inherited.NativeBridgeCertified {
		return Audit{}, fmt.Errorf("inherited leak: %s", FormatInherited(inherited))
	}
	hopf := buildHopfTypingAudit()
	if !hopf.HasS1Phase || !hopf.HasPositiveConjugatePair || !hopf.LabelsEPlusIfSealed || !hopf.LabelsEMinusIfSealed || hopf.TypedActionOnCR2Certified || hopf.NativeRepresentationMap {
		return Audit{}, fmt.Errorf("Hopf typing leak: %s", FormatHopfTyping(hopf))
	}
	cl17 := buildCL17TypingAudit()
	if !cl17.OmegaSquaredMinusOne || !cl17.SuppliesIOverMinusI || !cl17.CorrectConjugationShape || cl17.TypedToRightCharacter || cl17.SelectsSocketOrder {
		return Audit{}, fmt.Errorf("Cl17 typing leak: %s", FormatCL17Typing(cl17))
	}
	alignment := buildAlignmentAudit()
	if !alignment.CompatiblePhaseTypes || !alignment.StrongestBridgeCandidate || alignment.NativeAlignmentMap || alignment.TransportToRhoR || alignment.CanSourceRightCharacterAnchor {
		return Audit{}, fmt.Errorf("alignment leak: %s", FormatAlignment(alignment))
	}
	shape := buildShapeFirewallAudit()
	if !shape.SamePhaseShape || shape.TypedTransportCertified || !shape.ConjugationResonanceOnly || shape.NativeLambdaSelection {
		return Audit{}, fmt.Errorf("shape firewall leak: %s", FormatShapeFirewall(shape))
	}
	transport := buildTransportTargetAudit()
	if !transport.Sharpened || transport.NativeMapCertified || transport.MissingObject != PhaseTransportMap {
		return Audit{}, fmt.Errorf("transport target leak: %s", FormatTransportTarget(transport))
	}
	freeze := buildFreezeAudit()
	if !freeze.Frozen || !freeze.DiagnosticOnly || freeze.CanUpdate || near(freeze.OperatorNEff, freeze.OfficialNEff) {
		return Audit{}, fmt.Errorf("freeze leak: %s", FormatFreeze(freeze))
	}
	firewalls := buildFirewalls()
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewall leak: %s", FormatFirewalls(firewalls))
	}
	return Audit{
		ID:              AuditID,
		Inherited:       inherited,
		HopfTyping:      hopf,
		CL17Typing:      cl17,
		Alignment:       alignment,
		ShapeFirewall:   shape,
		TransportTarget: transport,
		Freeze:          freeze,
		Firewalls:       firewalls,
		Classification:  Classification,
		ShortStatus:     ShortStatus,
		Truth:           "Gate 903 certifies a Hopf/Cl(1,7)/right-character phase-shape alignment, but blocks promotion because no typed phase transport map into rho_R(lambda)=lambda e_+ + bar(lambda)e_- is certified.",
		Final:           ShortStatus,
	}, nil
}

func buildInheritedAudit() InheritedAudit {
	return InheritedAudit{Gate902Classification: "R3_AIRLOCK_PHASE_ANCHOR_REDUCED_TO_HOPF_CHIRALITY_BRIDGE_OBSTRUCTION", RightCharacterSplit: RightCharacterSplit, PhaseAnchorMissing: true, StrongSources: []string{HopfPhaseCircle, CL17Chirality}, NativeBridgeCertified: false, Supports: []string{StatusGate902Inherited, SupportGate902Inherited}, Failures: []string{FailureNoNativeRightPhaseOrientation, FailurePhaseAnchorStillSealed}}
}

func buildHopfTypingAudit() HopfTypingAudit {
	return HopfTypingAudit{CandidateName: HopfPhaseCircle, HasS1Phase: true, HasPositiveConjugatePair: true, LabelsEPlusIfSealed: true, LabelsEMinusIfSealed: true, TypedActionOnCR2Certified: false, NativeRepresentationMap: false, Supports: []string{StatusHopfTypingAudited, SupportHopfHasRightShape, SupportHopfPositiveLabelsEPlus, SupportHopfConjugateLabelsEMinus}, Failures: []string{FailureNoHopfToRightCharacterMap, FailureHopfLabelingStillSeal}}
}

func buildCL17TypingAudit() CL17TypingAudit {
	return CL17TypingAudit{CandidateName: CL17Chirality, RealBoard: "Cl(1,7)=Mat(16,R)", OmegaSquaredMinusOne: true, GammaChi: "gamma_chi=i omega", SuppliesIOverMinusI: true, CorrectConjugationShape: true, TypedToRightCharacter: false, SelectsSocketOrder: false, Supports: []string{StatusCL17TypingAudited, SupportCL17HasConjugationShape, SupportCL17IOrientation, SupportIVsMinusIShape}, Failures: []string{FailureNoGammaChiToRightAction, FailureChiralityDoesNotSelectSocket}}
}

func buildAlignmentAudit() AlignmentAudit {
	return AlignmentAudit{BridgeName: "HopfChiralityPhaseAnchorBridge", CompatiblePhaseTypes: true, StrongestBridgeCandidate: true, NativeAlignmentMap: false, TransportToRhoR: false, CanSourceRightCharacterAnchor: false, Supports: []string{StatusAlignmentAudited, SupportHopfCL17Align, SupportCompatibleOrientationTypes, SupportAlignmentStrongestCandidate, SupportRightAnchorIfTransportTyped}, Failures: []string{FailureNoHopfChiralityAlignment, FailureNoGammaTransportToRhoR, FailurePhaseAnchorStillSealed}}
}

func buildShapeFirewallAudit() ShapeFirewallAudit {
	return ShapeFirewallAudit{SamePhaseShape: true, TypedTransportCertified: false, ConjugationResonanceOnly: true, NativeLambdaSelection: false, Supports: []string{StatusFirewallAudited, SupportSourceSharpenedToTransport, SupportR3SealReducesToTransport}, Failures: []string{FailureShapeMatchNotTheorem, FailureConjugationResonanceNotMap, FailureNoNativeLambdaSelection}}
}

func buildTransportTargetAudit() TransportTargetAudit {
	return TransportTargetAudit{MissingObject: PhaseTransportMap, Domain: "Hopf S1 / gamma_chi=i omega orientation", Codomain: RightCharacterSplit, RequiredAction: "transport positive phase to lambda socket and conjugate phase to bar(lambda) socket", Sharpened: true, NativeMapCertified: false, Supports: []string{StatusTransportObjectIdentified, SupportSourceSharpenedToTransport, SupportR3SealReducesToTransport}, Failures: []string{FailureNoTypedPhaseTransport, FailureNoNativeRightPhaseOrientation}}
}

func buildFreezeAudit() FreezeAudit {
	return FreezeAudit{Alpha: AlphaB, OperatorNEff: OperatorNEffDiagnostic, OfficialNEff: OfficialNEffFrozen, OperatorCYukawa: OperatorCYukawaDiagnostic, OfficialCYukawa: OfficialCYukawaFrozen, OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen, Frozen: true, DiagnosticOnly: true, CanUpdate: false, Supports: []string{StatusOfficialFreeze, SupportDiagnosticsRemainCoherent}, Failures: []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate}}
}

func buildFirewalls() Firewalls {
	return Firewalls{Enforced: true, NoHopfRightCharacterMap: true, NoGammaChiRightAction: true, NoHopfChiralityAlignment: true, NoTypedPhaseTransport: true, ShapeMatchNotTheorem: true, PhaseAnchorStillSealed: true, NoNativeRightPhase: true, AlphaStillSealed: true, HiggsStillSealed: true, NotNativeR3: true, NoGenerationCarrier: true, NoFlavorOrientation: true, NoIndividualYukawas: true, NoOfficialLedgerUpdate: true, NoNativeYukawaOperator: true, NoR4YukawaTheorem: true, Verdict: StatusFirewallVerdict}
}

func FormatInherited(i InheritedAudit) string {
	return fmt.Sprintf("inherited(gate902=%s split=%s missing=%t sources=%s native_bridge=%t supports=%s failures=%s)", i.Gate902Classification, i.RightCharacterSplit, i.PhaseAnchorMissing, strings.Join(i.StrongSources, ","), i.NativeBridgeCertified, strings.Join(i.Supports, ","), strings.Join(i.Failures, ","))
}

func FormatHopfTyping(h HopfTypingAudit) string {
	return fmt.Sprintf("hopf_typing(candidate=%s s1=%t conjugate_pair=%t labels_eplus=%t labels_eminus=%t typed_action_cr2=%t native_map=%t supports=%s failures=%s)", h.CandidateName, h.HasS1Phase, h.HasPositiveConjugatePair, h.LabelsEPlusIfSealed, h.LabelsEMinusIfSealed, h.TypedActionOnCR2Certified, h.NativeRepresentationMap, strings.Join(h.Supports, ","), strings.Join(h.Failures, ","))
}

func FormatCL17Typing(c CL17TypingAudit) string {
	return fmt.Sprintf("cl17_typing(candidate=%s board=%s omega2_minus_one=%t gamma=%s i_orientation=%t conjugation_shape=%t typed_right=%t selects_socket=%t supports=%s failures=%s)", c.CandidateName, c.RealBoard, c.OmegaSquaredMinusOne, c.GammaChi, c.SuppliesIOverMinusI, c.CorrectConjugationShape, c.TypedToRightCharacter, c.SelectsSocketOrder, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}

func FormatAlignment(a AlignmentAudit) string {
	return fmt.Sprintf("alignment(bridge=%s compatible_types=%t strongest=%t native_alignment=%t transport_rhor=%t anchors_right=%t supports=%s failures=%s)", a.BridgeName, a.CompatiblePhaseTypes, a.StrongestBridgeCandidate, a.NativeAlignmentMap, a.TransportToRhoR, a.CanSourceRightCharacterAnchor, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatShapeFirewall(s ShapeFirewallAudit) string {
	return fmt.Sprintf("shape_firewall(same_shape=%t typed_transport=%t resonance_only=%t native_lambda_selection=%t supports=%s failures=%s)", s.SamePhaseShape, s.TypedTransportCertified, s.ConjugationResonanceOnly, s.NativeLambdaSelection, strings.Join(s.Supports, ","), strings.Join(s.Failures, ","))
}

func FormatTransportTarget(t TransportTargetAudit) string {
	return fmt.Sprintf("transport_target(missing=%s domain=%s codomain=%s required=%s sharpened=%t native=%t supports=%s failures=%s)", t.MissingObject, t.Domain, t.Codomain, t.RequiredAction, t.Sharpened, t.NativeMapCertified, strings.Join(t.Supports, ","), strings.Join(t.Failures, ","))
}

func FormatFreeze(f FreezeAudit) string {
	return fmt.Sprintf("freeze(alpha=%.16g operator_neff=%.16g official_neff=%.16g operator_cy=%.16g official_cy=%.16g operator_ch=%.16g official_ch=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.Alpha, f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t hopf_map=%t gamma_action=%t alignment=%t transport=%t shape_not_theorem=%t phase_sealed=%t native_right_phase=%t alpha=%t higgs=%t native_r3=%t generation=%t flavor=%t individual=%t official=%t yukawa=%t r4=%t verdict=%s)", f.Enforced, f.NoHopfRightCharacterMap, f.NoGammaChiRightAction, f.NoHopfChiralityAlignment, f.NoTypedPhaseTransport, f.ShapeMatchNotTheorem, f.PhaseAnchorStillSealed, f.NoNativeRightPhase, f.AlphaStillSealed, f.HiggsStillSealed, f.NotNativeR3, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4YukawaTheorem, f.Verdict)
}

func Statuses() []string {
	return []string{StatusGate902Inherited, StatusHopfTypingAudited, StatusCL17TypingAudited, StatusAlignmentAudited, StatusFirewallAudited, StatusTransportObjectIdentified, StatusOfficialFreeze, StatusFirewallVerdict, SupportHopfHasRightShape, SupportCL17HasConjugationShape, SupportHopfCL17Align, SupportSourceSharpenedToTransport, FailureNoHopfToRightCharacterMap, FailureNoGammaChiToRightAction, FailureNoHopfChiralityAlignment, FailureNoTypedPhaseTransport, FailureShapeMatchNotTheorem, FailureNotNativeR3}
}

func (a Audit) FirewallsList() []string {
	return []string{FailureNoHopfToRightCharacterMap, FailureNoGammaChiToRightAction, FailureNoHopfChiralityAlignment, FailureNoTypedPhaseTransport, FailureShapeMatchNotTheorem, FailureNoNativeRightPhaseOrientation, FailureNoNativeLambdaSelection, FailureAlphaStillSealedWithoutPhase, FailureHiggsStillSealedWithoutPhase, FailureNotNativeR3, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem}
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NoHopfRightCharacterMap && f.NoGammaChiRightAction && f.NoHopfChiralityAlignment && f.NoTypedPhaseTransport && f.ShapeMatchNotTheorem && f.PhaseAnchorStillSealed && f.NoNativeRightPhase && f.AlphaStillSealed && f.HiggsStillSealed && f.NotNativeR3 && f.NoGenerationCarrier && f.NoFlavorOrientation && f.NoIndividualYukawas && f.NoOfficialLedgerUpdate && f.NoNativeYukawaOperator && f.NoR4YukawaTheorem && f.Verdict == StatusFirewallVerdict
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
