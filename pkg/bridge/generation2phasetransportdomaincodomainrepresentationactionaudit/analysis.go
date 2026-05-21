// Package generation2phasetransportdomaincodomainrepresentationactionaudit implements
// Gate 904: PhaseTransport Domain/Codomain and Representation Action Audit.
//
// Gate 904 follows Gate 903's result that Hopf, Cl(1,7) chirality, and the
// right-character pair share a conjugate phase shape, but lack typed transport.
// It audits the exact domain, codomain, and representation action required for
// a HopfChiralityRightCharacterTransportMap. The honest result is that the
// domain and codomain are now typed, but the action map into C_R^2 / rho_R is
// missing. No alpha derivation, native R3 promotion, particle assignment,
// individual Yukawa value, or official ledger update is certified.
package generation2phasetransportdomaincodomainrepresentationactionaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE904-PHASE-TRANSPORT-DOMAIN-CODOMAIN-REPRESENTATION-ACTION-AUDIT"

	AlphaB = 0.0003878958469680527

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	RightCharacterSplit = "rho_R(lambda)=lambda e_+ + bar(lambda)e_-"
	HopfDomain          = "Hopf/S1 phase orientation"
	CL17Domain          = "Cl(1,7) complex chirality orientation gamma_chi=i omega"
	TransportMap        = "HopfChiralityRightCharacterTransportMap"
	CodomainCR2         = "End(C_R^2) with ordered right-character projector pair (e_lambda,e_barlambda)"

	Classification = "R3_PHASE_TRANSPORT_MAP_TYPED_ACTION_OBSTRUCTION"
	ShortStatus    = "R3_AIRLOCK_PHASE_TRANSPORT_DOMAIN_CODOMAIN_TYPED_BUT_NOT_NATIVE"

	StatusGate903Inherited       = "PASS_GATE903_PHASE_SHAPE_SUPPORTED_TRANSPORT_MISSING_INHERITED"
	StatusDomainTyped            = "PASS_PHASE_TRANSPORT_DOMAIN_TYPED"
	StatusCodomainTyped          = "PASS_PHASE_TRANSPORT_CODOMAIN_TYPED"
	StatusActionCompatibility    = "PASS_PHASE_TRANSPORT_ACTION_COMPATIBILITY_AUDITED"
	StatusNonCircularityAudited  = "PASS_PHASE_TRANSPORT_NONCIRCULARITY_FIREWALL_AUDITED"
	StatusAirlockEffectAudited   = "PASS_PHASE_TRANSPORT_SEAL_EFFECT_ON_AIRLOCK_AUDITED"
	StatusMissingObjectSharpened = "PASS_R3_MASTER_WOUND_REDUCED_TO_TYPED_PHASE_ACTION_ON_C_R2"
	StatusOfficialFreeze         = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict        = "FIREWALL_PRESERVED_GATE904_DOMAIN_CODOMAIN_TYPED_ACTION_MISSING"

	SupportGate903Inherited          = "CONDITIONAL_SUPPORT_GATE903_PHASE_TRANSPORT_WOUND_INHERITED"
	SupportDomainTyped               = "CONDITIONAL_SUPPORT_PHASE_TRANSPORT_DOMAIN_TYPED_AS_HOPF_S1_PLUS_CL17_CHIRALITY_ORIENTATION"
	SupportHopfDomainTyped           = "CONDITIONAL_SUPPORT_HOPF_PHASE_DOMAIN_HAS_CORRECT_ORIENTED_S1_TYPE"
	SupportCL17DomainTyped           = "CONDITIONAL_SUPPORT_CL17_CHIRALITY_DOMAIN_HAS_INTERNAL_COMPLEX_ORIENTATION"
	SupportCodomainTyped             = "CONDITIONAL_SUPPORT_PHASE_TRANSPORT_CODOMAIN_TYPED_AS_RIGHT_CHARACTER_PROJECTOR_PAIR"
	SupportRhoRTarget                = "CONDITIONAL_SUPPORT_RHO_R_IS_THE_REQUIRED_TARGET_REPRESENTATION"
	SupportTransportSealSelectsEPlus = "CONDITIONAL_SUPPORT_TRANSPORT_SEAL_WOULD_SELECT_E_PLUS_AS_LAMBDA_SOCKET"
	SupportTransportOrdersAirlock    = "CONDITIONAL_SUPPORT_TRANSPORT_SEAL_WOULD_ORDER_NEUTRAL_PUNCTURE_AIRLOCK"
	SupportTransportCollapsesWounds  = "CONDITIONAL_SUPPORT_PHASE_TRANSPORT_SEAL_WOULD_COLLAPSE_SOCKET_ORDER_EDGE_ORDER_ALPHA_AND_WEAK_KERNEL"
	SupportMasterWoundToActionCR2    = "CONDITIONAL_SUPPORT_R3_MASTER_WOUND_REDUCES_TO_TYPED_PHASE_ACTION_ON_C_R2"
	SupportDiagnosticsRemainCoherent = "CONDITIONAL_SUPPORT_OPERATOR_DIAGNOSTICS_REMAIN_COHERENT_UNDER_PHASE_TRANSPORT_SEAL"

	FailureNoNativePhaseTransport       = "FAILED_ROUTE_NO_NATIVE_PHASE_TRANSPORT_MAP"
	FailureNoTypedHopfActionCR2         = "FAILED_ROUTE_NO_TYPED_HOPF_PHASE_ACTION_ON_C_R2"
	FailureNoTypedGammaChiActionCR2     = "FAILED_ROUTE_NO_TYPED_GAMMA_CHI_ACTION_ON_C_R2"
	FailureNoHopfChiralityRhoRAction    = "FAILED_ROUTE_NO_HOPF_CHIRALITY_TO_RHO_R_REPRESENTATION_ACTION"
	FailureNoTransportToCR2Projectors   = "FAILED_ROUTE_NO_TRANSPORT_MAP_FROM_PHASE_DOMAIN_TO_C_R2_PROJECTORS"
	FailureTransportNotActionCompatible = "FAILED_ROUTE_PHASE_TRANSPORT_NOT_ACTION_COMPATIBLE_WITH_RHO_R_YET"
	FailureNoTypedPhaseActionRightPair  = "FAILED_ROUTE_NO_TYPED_PHASE_ACTION_ON_RIGHT_SOCKET_PAIR"
	FailureRhoRRestatesOrder            = "FAILED_ROUTE_RHO_R_LABELING_RESTATES_SOCKET_ORDER_WITHOUT_SOURCE"
	FailureTransportByTargetLabelsOnly  = "FAILED_ROUTE_PHASE_TRANSPORT_CANNOT_BE_DEFINED_BY_TARGET_LABELS_ONLY"
	FailureShapeNotTransport            = "FAILED_ROUTE_PHASE_SHAPE_MATCH_NOT_PHASE_TRANSPORT"
	FailureNoLambdaSelection            = "FAILED_ROUTE_NO_NATIVE_SELECTION_OF_LAMBDA_OVER_BARLAMBDA"
	FailurePhaseAnchorSealed            = "FAILED_ROUTE_PHASE_ANCHOR_REMAINS_SEALED"
	FailureTransportSealNotNativeR3     = "FAILED_ROUTE_TRANSPORT_SEAL_NOT_NATIVE_R3"
	FailureAlphaStillSealedWithoutPhase = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED_WITHOUT_NATIVE_PHASE_ANCHORED_BOUNDARY_FUNCTOR"
	FailureHiggsStillSealedWithoutPhase = "FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED_WITHOUT_NATIVE_PHASE_ANCHORED_WEAK_SELECTOR"
	FailureNotNativeR3                  = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureNoGenerationCarrierMap       = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap       = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues     = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate         = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate        = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator       = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4NativeYukawaTheorem      = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type InheritedAudit struct {
	Gate903Classification string
	ShapeSupported        bool
	TransportMissing      bool
	MissingObject         string
	NativeTransport       bool
	Supports, Failures    []string
}

type DomainAudit struct {
	SourceDomain          string
	HopfS1Typed           bool
	CL17ChiralityTyped    bool
	HopfActsOnCR2         bool
	GammaChiActsOnCR2     bool
	NativeDomainActionMap bool
	Supports, Failures    []string
}

type CodomainAudit struct {
	Codomain                    string
	RightCharacterSplit         string
	ProjectorPairTyped          bool
	OutputsOrderedPair          bool
	NativeTransportToProjectors bool
	Supports, Failures          []string
}

type ActionAudit struct {
	RequiredPositiveAction   string
	RequiredConjugateAction  string
	PositiveToEPlus          bool
	ConjugateToEMinus        bool
	ActionCompatibleWithRhoR bool
	TypedActionOnRightPair   bool
	Supports, Failures       []string
}

type NonCircularityAudit struct {
	RhoRLabelsSockets          bool
	RhoRExplainsOrdering       bool
	TransportDefinedByLabels   bool
	TargetLabelRestatement     bool
	NonCircularSourceCertified bool
	Supports, Failures         []string
}

type AirlockEffectAudit struct {
	IfTransportSealed     bool
	SelectsEPlusAsLambda  bool
	OrdersNeutralPuncture bool
	CollapsesLocalWounds  bool
	NativeR3Promotion     bool
	Supports, Failures    []string
}

type MissingObjectAudit struct {
	MissingObject      string
	Domain             string
	Codomain           string
	RequiredAction     string
	NowFullyTyped      bool
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
	Enforced                bool
	NoNativePhaseTransport  bool
	NoHopfActionCR2         bool
	NoGammaActionCR2        bool
	NoHopfChiralityRhoR     bool
	NoActionCompatibility   bool
	NoNonCircularRhoRSource bool
	ShapeNotTransport       bool
	PhaseAnchorStillSealed  bool
	NotNativeR3             bool
	AlphaStillSealed        bool
	HiggsStillSealed        bool
	NoGenerationCarrier     bool
	NoFlavorOrientation     bool
	NoIndividualYukawas     bool
	NoOfficialLedgerUpdate  bool
	NoNativeYukawaOperator  bool
	NoR4YukawaTheorem       bool
	Verdict                 string
}

type Audit struct {
	ID             string
	Inherited      InheritedAudit
	Domain         DomainAudit
	Codomain       CodomainAudit
	Action         ActionAudit
	NonCircularity NonCircularityAudit
	AirlockEffect  AirlockEffectAudit
	MissingObject  MissingObjectAudit
	Freeze         FreezeAudit
	Firewalls      Firewalls
	Classification string
	ShortStatus    string
	Truth          string
	Final          string
}

func BuildDefault() (Audit, error) {
	inherited := buildInheritedAudit()
	if !inherited.ShapeSupported || !inherited.TransportMissing || inherited.NativeTransport || inherited.MissingObject != TransportMap {
		return Audit{}, fmt.Errorf("inherited leak: %s", FormatInherited(inherited))
	}
	domain := buildDomainAudit()
	if !domain.HopfS1Typed || !domain.CL17ChiralityTyped || domain.HopfActsOnCR2 || domain.GammaChiActsOnCR2 || domain.NativeDomainActionMap {
		return Audit{}, fmt.Errorf("domain leak: %s", FormatDomain(domain))
	}
	codomain := buildCodomainAudit()
	if !codomain.ProjectorPairTyped || !codomain.OutputsOrderedPair || codomain.NativeTransportToProjectors {
		return Audit{}, fmt.Errorf("codomain leak: %s", FormatCodomain(codomain))
	}
	action := buildActionAudit()
	if !action.PositiveToEPlus || !action.ConjugateToEMinus || action.ActionCompatibleWithRhoR || action.TypedActionOnRightPair {
		return Audit{}, fmt.Errorf("action leak: %s", FormatAction(action))
	}
	noncircular := buildNonCircularityAudit()
	if !noncircular.RhoRLabelsSockets || noncircular.RhoRExplainsOrdering || noncircular.TransportDefinedByLabels || !noncircular.TargetLabelRestatement || noncircular.NonCircularSourceCertified {
		return Audit{}, fmt.Errorf("noncircularity leak: %s", FormatNonCircularity(noncircular))
	}
	airlock := buildAirlockEffectAudit()
	if !airlock.IfTransportSealed || !airlock.SelectsEPlusAsLambda || !airlock.OrdersNeutralPuncture || !airlock.CollapsesLocalWounds || airlock.NativeR3Promotion {
		return Audit{}, fmt.Errorf("airlock effect leak: %s", FormatAirlockEffect(airlock))
	}
	missing := buildMissingObjectAudit()
	if missing.MissingObject != TransportMap || !missing.NowFullyTyped || missing.NativeMapCertified {
		return Audit{}, fmt.Errorf("missing object leak: %s", FormatMissingObject(missing))
	}
	freeze := buildFreezeAudit()
	if !freeze.Frozen || !freeze.DiagnosticOnly || freeze.CanUpdate || near(freeze.OperatorNEff, freeze.OfficialNEff) {
		return Audit{}, fmt.Errorf("freeze leak: %s", FormatFreeze(freeze))
	}
	firewalls := buildFirewalls()
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewall leak: %s", FormatFirewalls(firewalls))
	}
	return Audit{ID: AuditID, Inherited: inherited, Domain: domain, Codomain: codomain, Action: action, NonCircularity: noncircular, AirlockEffect: airlock, MissingObject: missing, Freeze: freeze, Firewalls: firewalls, Classification: Classification, ShortStatus: ShortStatus, Truth: "Gate 904 types the phase-transport domain and codomain, but blocks promotion because no action-compatible map from Hopf/S1 or gamma_chi orientation into C_R^2 / rho_R is certified.", Final: ShortStatus}, nil
}

func buildInheritedAudit() InheritedAudit {
	return InheritedAudit{Gate903Classification: "R3_AIRLOCK_PHASE_ANCHOR_SHAPE_SUPPORTED_TRANSPORT_MISSING", ShapeSupported: true, TransportMissing: true, MissingObject: TransportMap, NativeTransport: false, Supports: []string{StatusGate903Inherited, SupportGate903Inherited}, Failures: []string{FailureShapeNotTransport, FailurePhaseAnchorSealed}}
}

func buildDomainAudit() DomainAudit {
	return DomainAudit{SourceDomain: HopfDomain + " plus " + CL17Domain, HopfS1Typed: true, CL17ChiralityTyped: true, HopfActsOnCR2: false, GammaChiActsOnCR2: false, NativeDomainActionMap: false, Supports: []string{StatusDomainTyped, SupportDomainTyped, SupportHopfDomainTyped, SupportCL17DomainTyped}, Failures: []string{FailureNoTypedHopfActionCR2, FailureNoTypedGammaChiActionCR2}}
}

func buildCodomainAudit() CodomainAudit {
	return CodomainAudit{Codomain: CodomainCR2, RightCharacterSplit: RightCharacterSplit, ProjectorPairTyped: true, OutputsOrderedPair: true, NativeTransportToProjectors: false, Supports: []string{StatusCodomainTyped, SupportCodomainTyped, SupportRhoRTarget}, Failures: []string{FailureNoTransportToCR2Projectors}}
}

func buildActionAudit() ActionAudit {
	return ActionAudit{RequiredPositiveAction: "T_phase(+)=e_+ and positive phase acts by lambda on e_+", RequiredConjugateAction: "T_phase(-)=e_- and conjugate phase acts by bar(lambda) on e_-", PositiveToEPlus: true, ConjugateToEMinus: true, ActionCompatibleWithRhoR: false, TypedActionOnRightPair: false, Supports: []string{StatusActionCompatibility, SupportTransportSealSelectsEPlus}, Failures: []string{FailureTransportNotActionCompatible, FailureNoTypedPhaseActionRightPair, FailureNoHopfChiralityRhoRAction}}
}

func buildNonCircularityAudit() NonCircularityAudit {
	return NonCircularityAudit{RhoRLabelsSockets: true, RhoRExplainsOrdering: false, TransportDefinedByLabels: false, TargetLabelRestatement: true, NonCircularSourceCertified: false, Supports: []string{StatusNonCircularityAudited, SupportRhoRTarget}, Failures: []string{FailureRhoRRestatesOrder, FailureTransportByTargetLabelsOnly}}
}

func buildAirlockEffectAudit() AirlockEffectAudit {
	return AirlockEffectAudit{IfTransportSealed: true, SelectsEPlusAsLambda: true, OrdersNeutralPuncture: true, CollapsesLocalWounds: true, NativeR3Promotion: false, Supports: []string{StatusAirlockEffectAudited, SupportTransportSealSelectsEPlus, SupportTransportOrdersAirlock, SupportTransportCollapsesWounds}, Failures: []string{FailureTransportSealNotNativeR3}}
}

func buildMissingObjectAudit() MissingObjectAudit {
	return MissingObjectAudit{MissingObject: TransportMap, Domain: HopfDomain + " plus " + CL17Domain, Codomain: CodomainCR2, RequiredAction: "positive phase acts on C_R^2 as the lambda-character socket and conjugate phase as the bar(lambda)-character socket", NowFullyTyped: true, NativeMapCertified: false, Supports: []string{StatusMissingObjectSharpened, SupportMasterWoundToActionCR2}, Failures: []string{FailureNoNativePhaseTransport, FailureNoHopfChiralityRhoRAction}}
}

func buildFreezeAudit() FreezeAudit {
	return FreezeAudit{Alpha: AlphaB, OperatorNEff: OperatorNEffDiagnostic, OfficialNEff: OfficialNEffFrozen, OperatorCYukawa: OperatorCYukawaDiagnostic, OfficialCYukawa: OfficialCYukawaFrozen, OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen, Frozen: true, DiagnosticOnly: true, CanUpdate: false, Supports: []string{StatusOfficialFreeze, SupportDiagnosticsRemainCoherent}, Failures: []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate}}
}

func buildFirewalls() Firewalls {
	return Firewalls{Enforced: true, NoNativePhaseTransport: true, NoHopfActionCR2: true, NoGammaActionCR2: true, NoHopfChiralityRhoR: true, NoActionCompatibility: true, NoNonCircularRhoRSource: true, ShapeNotTransport: true, PhaseAnchorStillSealed: true, NotNativeR3: true, AlphaStillSealed: true, HiggsStillSealed: true, NoGenerationCarrier: true, NoFlavorOrientation: true, NoIndividualYukawas: true, NoOfficialLedgerUpdate: true, NoNativeYukawaOperator: true, NoR4YukawaTheorem: true, Verdict: StatusFirewallVerdict}
}

func FormatInherited(i InheritedAudit) string {
	return fmt.Sprintf("inherited(gate903=%s shape=%t transport_missing=%t missing=%s native=%t supports=%s failures=%s)", i.Gate903Classification, i.ShapeSupported, i.TransportMissing, i.MissingObject, i.NativeTransport, strings.Join(i.Supports, ","), strings.Join(i.Failures, ","))
}

func FormatDomain(d DomainAudit) string {
	return fmt.Sprintf("domain(source=%s hopf_s1=%t cl17=%t hopf_action_cr2=%t gamma_action_cr2=%t native_action=%t supports=%s failures=%s)", d.SourceDomain, d.HopfS1Typed, d.CL17ChiralityTyped, d.HopfActsOnCR2, d.GammaChiActsOnCR2, d.NativeDomainActionMap, strings.Join(d.Supports, ","), strings.Join(d.Failures, ","))
}

func FormatCodomain(c CodomainAudit) string {
	return fmt.Sprintf("codomain(target=%s split=%s projectors=%t ordered_output=%t native_transport=%t supports=%s failures=%s)", c.Codomain, c.RightCharacterSplit, c.ProjectorPairTyped, c.OutputsOrderedPair, c.NativeTransportToProjectors, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}

func FormatAction(a ActionAudit) string {
	return fmt.Sprintf("action(positive=%s conjugate=%s positive_to_eplus=%t conjugate_to_eminus=%t compatible_rhor=%t typed_right_pair=%t supports=%s failures=%s)", a.RequiredPositiveAction, a.RequiredConjugateAction, a.PositiveToEPlus, a.ConjugateToEMinus, a.ActionCompatibleWithRhoR, a.TypedActionOnRightPair, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatNonCircularity(n NonCircularityAudit) string {
	return fmt.Sprintf("noncircularity(rhor_labels=%t rhor_explains_order=%t defined_by_labels=%t target_label_restatement=%t noncircular_source=%t supports=%s failures=%s)", n.RhoRLabelsSockets, n.RhoRExplainsOrdering, n.TransportDefinedByLabels, n.TargetLabelRestatement, n.NonCircularSourceCertified, strings.Join(n.Supports, ","), strings.Join(n.Failures, ","))
}

func FormatAirlockEffect(a AirlockEffectAudit) string {
	return fmt.Sprintf("airlock_effect(if_sealed=%t selects_eplus=%t orders_puncture=%t collapses_wounds=%t native_r3=%t supports=%s failures=%s)", a.IfTransportSealed, a.SelectsEPlusAsLambda, a.OrdersNeutralPuncture, a.CollapsesLocalWounds, a.NativeR3Promotion, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatMissingObject(m MissingObjectAudit) string {
	return fmt.Sprintf("missing_object(name=%s domain=%s codomain=%s required=%s typed=%t native=%t supports=%s failures=%s)", m.MissingObject, m.Domain, m.Codomain, m.RequiredAction, m.NowFullyTyped, m.NativeMapCertified, strings.Join(m.Supports, ","), strings.Join(m.Failures, ","))
}

func FormatFreeze(f FreezeAudit) string {
	return fmt.Sprintf("freeze(alpha=%.16g operator_neff=%.16g official_neff=%.16g operator_cy=%.16g official_cy=%.16g operator_ch=%.16g official_ch=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.Alpha, f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t native_transport=%t hopf_cr2=%t gamma_cr2=%t hopfchirality_rhor=%t action=%t noncircular=%t shape=%t phase_sealed=%t native_r3=%t alpha=%t higgs=%t generation=%t flavor=%t individual=%t official=%t yukawa=%t r4=%t verdict=%s)", f.Enforced, f.NoNativePhaseTransport, f.NoHopfActionCR2, f.NoGammaActionCR2, f.NoHopfChiralityRhoR, f.NoActionCompatibility, f.NoNonCircularRhoRSource, f.ShapeNotTransport, f.PhaseAnchorStillSealed, f.NotNativeR3, f.AlphaStillSealed, f.HiggsStillSealed, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4YukawaTheorem, f.Verdict)
}

func Statuses() []string {
	return []string{StatusGate903Inherited, StatusDomainTyped, StatusCodomainTyped, StatusActionCompatibility, StatusNonCircularityAudited, StatusAirlockEffectAudited, StatusMissingObjectSharpened, StatusOfficialFreeze, StatusFirewallVerdict, SupportDomainTyped, SupportCodomainTyped, SupportMasterWoundToActionCR2, FailureNoNativePhaseTransport, FailureNoTypedHopfActionCR2, FailureNoTypedGammaChiActionCR2, FailureNoHopfChiralityRhoRAction, FailureRhoRRestatesOrder, FailureNotNativeR3}
}

func (a Audit) FirewallsList() []string {
	return []string{FailureNoNativePhaseTransport, FailureNoTypedHopfActionCR2, FailureNoTypedGammaChiActionCR2, FailureNoHopfChiralityRhoRAction, FailureTransportNotActionCompatible, FailureRhoRRestatesOrder, FailureShapeNotTransport, FailureNoLambdaSelection, FailurePhaseAnchorSealed, FailureAlphaStillSealedWithoutPhase, FailureHiggsStillSealedWithoutPhase, FailureNotNativeR3, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem}
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NoNativePhaseTransport && f.NoHopfActionCR2 && f.NoGammaActionCR2 && f.NoHopfChiralityRhoR && f.NoActionCompatibility && f.NoNonCircularRhoRSource && f.ShapeNotTransport && f.PhaseAnchorStillSealed && f.NotNativeR3 && f.AlphaStillSealed && f.HiggsStillSealed && f.NoGenerationCarrier && f.NoFlavorOrientation && f.NoIndividualYukawas && f.NoOfficialLedgerUpdate && f.NoNativeYukawaOperator && f.NoR4YukawaTheorem && f.Verdict == StatusFirewallVerdict
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
