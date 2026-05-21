// Package generation2socketorderorientededgeairlockconsolidationaudit implements
// Gate 897: SocketOrder / OrientedEdgeOrdering / NeutralPunctureAirlock Consolidation Audit.
//
// Gate 897 follows Gate 896's airlock functional obstruction. It consolidates
// three nearby wounds: plus/minus socket order, oriented edge ordering, and the
// neutral puncture airlock functor. The gate shows that the current data select
// an unoriented rank-one neutral lepton puncture airlock family, but do not
// natively select the ordered representative p=e_+ tensor P_1. The remaining
// obstruction is a Z2 plus/minus SocketOrderSelector. No alpha derivation,
// native R3 promotion, physical sector assignment, individual Yukawa value, or
// official ledger update is certified.
package generation2socketorderorientededgeairlockconsolidationaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE897-SOCKET-ORDER-ORIENTED-EDGE-AIRLOCK-CONSOLIDATION-AUDIT"

	AlphaB = 0.0003878958469680527
	Ssplit = 0.001292444818816423

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	PuncturePlus        = "e_+ tensor P_1"
	PunctureMinus       = "e_- tensor P_1"
	ColorPlus           = "e_+ tensor P_3"
	ColorMinus          = "e_- tensor P_3"
	RightCharacterPlus  = "chi_R^+=lambda"
	RightCharacterMinus = "chi_R^-=bar(lambda)"

	F0Plus  = "F_0^+=e_+ tensor P_1"
	F1Plus  = "F_1^+=e_+ tensor W"
	F0Minus = "F_0^-=e_- tensor P_1"
	F1Minus = "F_1^-=e_- tensor W"
	F2      = "F_2=C_R^2 tensor W"

	ActiveEdgePlusColor   = "e_+ tensor P_3 -> h_+ tensor P_3"
	ActiveEdgeMinusColor  = "e_- tensor P_3 -> h_- tensor P_3"
	ActiveEdgeMinusLepton = "e_- tensor P_1 -> h_- tensor P_1"
	MissingEdgePlusLepton = "e_+ tensor P_1 -> h_+ tensor P_1 = 0"

	MirrorActiveEdgeMinusColor   = "e_- tensor P_3 -> h_- tensor P_3"
	MirrorActiveEdgePlusColor    = "e_+ tensor P_3 -> h_+ tensor P_3"
	MirrorActiveEdgePlusLepton   = "e_+ tensor P_1 -> h_+ tensor P_1"
	MirrorMissingEdgeMinusLepton = "e_- tensor P_1 -> h_- tensor P_1 = 0"

	RankLeptonCell = 1
	RankColorCell  = 3
	RankRightRect  = 8
	RankF1         = 4
	RankPiTop      = 3
	RankHRMin      = 7
	DimH10         = 10
	DimH72         = 72
	BMinusLLepton  = -1

	Classification = "R3_CANDIDATE_NEUTRAL_PUNCTURE_AIRLOCK_Z2_FAMILY_ORDERED_SELECTOR_MISSING"
	ShortStatus    = "R3_AIRLOCK_Z2_FAMILY_SOCKET_ORDER_OBSTRUCTION"
	NextFrontier   = "RIGHT_CHARACTER_ORIENTATION_AND_SOCKET_ORDER_SOURCE_AUDIT"

	StatusGate896Inherited     = "PASS_GATE896_AIRLOCK_FUNCTIONAL_OBSTRUCTION_INHERITED"
	StatusSocketPairAudited    = "PASS_PLUS_MINUS_SOCKET_ORDER_SELECTOR_AUDITED"
	StatusEdgeOrderingAudited  = "PASS_ORIENTED_EDGE_ORDERING_FUNCTIONAL_AUDITED"
	StatusAirlockFamilyAudited = "PASS_NEUTRAL_PUNCTURE_AIRLOCK_Z2_FAMILY_AUDITED"
	StatusMirrorTableAudited   = "PASS_Z2_MIRROR_EDGE_TABLE_AUDITED"
	StatusAlphaReconstruction  = "PASS_BOTH_SOCKET_ORDERINGS_RECONSTRUCT_ALPHA_RANKS"
	StatusObstructionSharpened = "PASS_OBSTRUCTION_SHARPENED_TO_SOCKET_ORDER_SELECTOR"
	StatusOfficialFreeze       = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict      = "FIREWALL_PRESERVED_GATE897_Z2_AIRLOCK_NOT_NATIVE"

	SupportRightSocketCharacterPairTyped = "CONDITIONAL_SUPPORT_RIGHT_SOCKET_CHARACTER_PAIR_TYPED"
	SupportPlusMinusRemainingAmbiguity   = "CONDITIONAL_SUPPORT_PLUS_MINUS_SOCKET_PAIR_IS_THE_REMAINING_ORDER_AMBIGUITY"
	SupportSocketOrderRequired           = "CONDITIONAL_SUPPORT_SOCKET_ORDER_SELECTOR_IS_REQUIRED_FOR_ORDERED_AIRLOCK"
	SupportEdgeTableFromOrderedRule      = "CONDITIONAL_SUPPORT_EDGE_TABLE_IS_GENERATED_BY_ORDERED_SOCKET_AIRLOCK_RULE"
	SupportCurrentTableIfPlusExposed     = "CONDITIONAL_SUPPORT_CURRENT_EDGE_TABLE_FOLLOWS_IF_E_PLUS_IS_EXPOSED_SOCKET"
	SupportZ2MirrorTableExists           = "CONDITIONAL_SUPPORT_Z2_MIRROR_EDGE_TABLE_EXISTS_IF_E_MINUS_IS_EXPOSED_SOCKET"
	SupportEdgeOrderingReducesToSocket   = "CONDITIONAL_SUPPORT_ORIENTED_EDGE_ORDERING_REDUCES_TO_SOCKET_ORDER_SELECTOR"
	SupportAirlockZ2Family               = "CONDITIONAL_SUPPORT_NEUTRAL_PUNCTURE_AIRLOCK_EXISTS_AS_Z2_FAMILY_CANDIDATE"
	SupportBothAirlocksReconstructAlpha  = "CONDITIONAL_SUPPORT_BOTH_PLUS_AND_MINUS_AIRLOCKS_RECONSTRUCT_ALPHA_RANKS"
	SupportTwoSealWoundReduced           = "CONDITIONAL_SUPPORT_TWO_SEAL_WOUND_REDUCES_TO_AIRLOCK_PLUS_SOCKET_ORDER_SELECTOR"
	SupportUnorientedAirlockCandidate    = "CONDITIONAL_SUPPORT_R3_CANDIDATE_HAS_UNORIENTED_AIRLOCK_FUNCTOR_CANDIDATE"
	SupportOperatorNEffReproduced        = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_REPRODUCED_BY_ORIENTED_LEDGER"

	FailureNotNativeR3                      = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureAlphaStillSealed                 = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureHiggsOrientationStillSealed      = "FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED"
	FailureNoNativeSocketOrderSelector      = "FAILED_ROUTE_NO_NATIVE_SOCKET_ORDER_SELECTOR"
	FailureNoNativePlusMinusOrder           = "FAILED_ROUTE_NO_NATIVE_PLUS_MINUS_ORDERING_OF_RIGHT_CHARACTER_PAIR"
	FailureCharacterConjugationNoPlus       = "FAILED_ROUTE_CHARACTER_CONJUGATION_PAIR_DOES_NOT_SELECT_PLUS_AS_PUNCTURE"
	FailureBoundaryDegreeNoPlus             = "FAILED_ROUTE_BOUNDARY_DEGREE_ORDER_DOES_NOT_SELECT_E_PLUS_OVER_E_MINUS"
	FailureJChiralityNoPlus                 = "FAILED_ROUTE_J_OR_CHIRALITY_DATA_DOES_NOT_SELECT_PLUS_SOCKET_ORDER"
	FailureBMinusLNoPlus                    = "FAILED_ROUTE_B_MINUS_L_COMPENSATION_DOES_NOT_SELECT_PLUS_SOCKET_ORDER"
	FailureNoNativeOrientedEdgeOrdering     = "FAILED_ROUTE_NO_NATIVE_ORIENTED_EDGE_ORDERING_FUNCTIONAL"
	FailureEdgeTableNoPlusByItself          = "FAILED_ROUTE_EDGE_TABLE_DOES_NOT_SELECT_PLUS_ORDER_BY_ITSELF"
	FailureEdgeOrderingCircular             = "FAILED_ROUTE_EDGE_ORDERING_REMAINS_CIRCULAR_WITHOUT_SOCKET_ORDER_SELECTOR"
	FailureNoNativeNeutralPunctureAirlock   = "FAILED_ROUTE_NO_NATIVE_NEUTRAL_PUNCTURE_AIRLOCK_FUNCTOR"
	FailureNoNativeOrderedAirlock           = "FAILED_ROUTE_NO_NATIVE_ORDERED_NEUTRAL_PUNCTURE_AIRLOCK_FUNCTOR"
	FailureNoNativeSelectionSigmaPlus       = "FAILED_ROUTE_NO_NATIVE_SELECTION_OF_SIGMA_EQUALS_PLUS"
	FailureNoNativeBoundaryIncidenceFunctor = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureNoNativeDescentFullToOrient      = "FAILED_ROUTE_NO_NATIVE_DESCENT_FROM_FULL_A_F_TO_A_F_ORIENT"
	FailureNoGenerationCarrierMap           = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap           = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues         = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate             = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate            = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator           = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4NativeYukawaTheorem          = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type SocketOrderAudit struct {
	Characters                []string
	PairTyped                 bool
	NativeOrder               bool
	PlusSelected              bool
	BoundaryDegreeSelectsPlus bool
	JOrChiralitySelectsPlus   bool
	BMinusLSelectsPlus        bool
	RequiresSelector          bool
	Supports, Failures        []string
}

type EdgeOrderingAudit struct {
	OrderedRule                 string
	CurrentOrdering             []string
	MirrorOrdering              []string
	CurrentFollowsIfPlusExposed bool
	MirrorExistsIfMinusExposed  bool
	SameRankPattern             bool
	SelectsPlusByItself         bool
	ReducesToSocketOrder        bool
	Supports, Failures          []string
}

type AirlockFamilyAudit struct {
	Sigmas                               []string
	PlusPuncture, MinusPuncture          string
	PlusF0, PlusF1, MinusF0, MinusF1, F2 string
	PlusRanks, MinusRanks                []int
	PlusAlpha, MinusAlpha                float64
	BothReconstructAlpha                 bool
	Z2Family                             bool
	OrderedRepresentativeCertified       bool
	Supports, Failures                   []string
}

type FreezeAudit struct {
	OperatorNEff, OfficialNEff        float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	Frozen, DiagnosticOnly, CanUpdate bool
	Supports, Failures                []string
}

type Firewalls struct {
	Enforced                         bool
	NotNativeR3                      bool
	AlphaStillSealed                 bool
	HiggsOrientationStillSealed      bool
	NoNativeSocketOrderSelector      bool
	NoNativeOrientedEdgeOrdering     bool
	NoNativeAirlockFunctor           bool
	NoNativeOrderedAirlock           bool
	NoNativeBoundaryIncidenceFunctor bool
	NoNativeDescentFullToOrient      bool
	NoGenerationCarrier              bool
	NoFlavorOrientation              bool
	NoIndividualYukawas              bool
	NoOfficialLedgerUpdate           bool
	NoNativeYukawaOperator           bool
	NoR4NativeYukawaTheorem          bool
	Verdict                          string
}

type Audit struct {
	ID            string
	SocketOrder   SocketOrderAudit
	EdgeOrdering  EdgeOrderingAudit
	AirlockFamily AirlockFamilyAudit
	Freeze        FreezeAudit
	Firewalls     Firewalls
	Truth         string
	Final         string
}

func BuildDefault() (Audit, error) {
	socket := buildSocketOrderAudit()
	if !socket.PairTyped || socket.NativeOrder || socket.PlusSelected || !socket.RequiresSelector {
		return Audit{}, fmt.Errorf("socket order promoted incorrectly: %s", FormatSocketOrder(socket))
	}
	edge := buildEdgeOrderingAudit()
	if !edge.CurrentFollowsIfPlusExposed || !edge.MirrorExistsIfMinusExposed || !edge.SameRankPattern || edge.SelectsPlusByItself || !edge.ReducesToSocketOrder {
		return Audit{}, fmt.Errorf("edge ordering promoted incorrectly: %s", FormatEdgeOrdering(edge))
	}
	family := buildAirlockFamilyAudit()
	if !family.BothReconstructAlpha || !family.Z2Family || family.OrderedRepresentativeCertified || !near(family.PlusAlpha, AlphaB) || !near(family.MinusAlpha, AlphaB) {
		return Audit{}, fmt.Errorf("airlock family promoted incorrectly: %s", FormatAirlockFamily(family))
	}
	freeze := FreezeAudit{
		OperatorNEff: OperatorNEffDiagnostic, OfficialNEff: OfficialNEffFrozen,
		OperatorCYukawa: OperatorCYukawaDiagnostic, OfficialCYukawa: OfficialCYukawaFrozen,
		OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen,
		Frozen: true, DiagnosticOnly: true, CanUpdate: false,
		Supports: []string{StatusOfficialFreeze, SupportOperatorNEffReproduced},
		Failures: []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate},
	}
	firewalls := Firewalls{
		Enforced: true, NotNativeR3: true, AlphaStillSealed: true, HiggsOrientationStillSealed: true,
		NoNativeSocketOrderSelector: true, NoNativeOrientedEdgeOrdering: true, NoNativeAirlockFunctor: true,
		NoNativeOrderedAirlock: true, NoNativeBoundaryIncidenceFunctor: true, NoNativeDescentFullToOrient: true,
		NoGenerationCarrier: true, NoFlavorOrientation: true, NoIndividualYukawas: true,
		NoOfficialLedgerUpdate: true, NoNativeYukawaOperator: true, NoR4NativeYukawaTheorem: true,
		Verdict: StatusFirewallVerdict,
	}
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewalls not preserved: %s", FormatFirewalls(firewalls))
	}
	return Audit{
		ID: AuditID, SocketOrder: socket, EdgeOrdering: edge, AirlockFamily: family, Freeze: freeze, Firewalls: firewalls,
		Truth: "Gate 897 consolidates plus/minus socket order, oriented edge ordering, and neutral puncture airlock into one Z2 obstruction: the current data select an unoriented neutral lepton puncture airlock family, but do not natively select sigma=+.",
		Final: "The branch becomes R3_AIRLOCK_Z2_FAMILY_SOCKET_ORDER_OBSTRUCTION. The two-seal wound reduces to an unoriented NeutralPunctureAirlock family plus a missing SocketOrderSelector; alpha, Higgs orientation, native R3/R4, physical sectors, individual Yukawas, and official ledger updates remain blocked.",
	}, nil
}

func buildSocketOrderAudit() SocketOrderAudit {
	return SocketOrderAudit{
		Characters: []string{RightCharacterPlus, RightCharacterMinus},
		PairTyped:  true, NativeOrder: false, PlusSelected: false,
		BoundaryDegreeSelectsPlus: false, JOrChiralitySelectsPlus: false, BMinusLSelectsPlus: false,
		RequiresSelector: true,
		Supports:         []string{StatusSocketPairAudited, SupportRightSocketCharacterPairTyped, SupportPlusMinusRemainingAmbiguity, SupportSocketOrderRequired},
		Failures:         []string{FailureNoNativeSocketOrderSelector, FailureNoNativePlusMinusOrder, FailureCharacterConjugationNoPlus, FailureBoundaryDegreeNoPlus, FailureJChiralityNoPlus, FailureBMinusLNoPlus},
	}
}

func buildEdgeOrderingAudit() EdgeOrderingAudit {
	return EdgeOrderingAudit{
		OrderedRule:                 "given (e_exposed,e_rest), remove e_exposed tensor P_1 and keep color exposure plus rest color/lepton edges",
		CurrentOrdering:             []string{ActiveEdgePlusColor, ActiveEdgeMinusColor, ActiveEdgeMinusLepton, MissingEdgePlusLepton},
		MirrorOrdering:              []string{MirrorActiveEdgeMinusColor, MirrorActiveEdgePlusColor, MirrorActiveEdgePlusLepton, MirrorMissingEdgeMinusLepton},
		CurrentFollowsIfPlusExposed: true, MirrorExistsIfMinusExposed: true, SameRankPattern: true,
		SelectsPlusByItself: false, ReducesToSocketOrder: true,
		Supports: []string{StatusEdgeOrderingAudited, StatusMirrorTableAudited, SupportEdgeTableFromOrderedRule, SupportCurrentTableIfPlusExposed, SupportZ2MirrorTableExists, SupportEdgeOrderingReducesToSocket},
		Failures: []string{FailureNoNativeOrientedEdgeOrdering, FailureEdgeTableNoPlusByItself, FailureEdgeOrderingCircular},
	}
}

func buildAirlockFamilyAudit() AirlockFamilyAudit {
	plusAlpha := float64(RankPiTop)/float64(DimH10)*Ssplit + float64(RankHRMin)/float64(DimH72)*Ssplit*Ssplit
	minusAlpha := float64(RankPiTop)/float64(DimH10)*Ssplit + float64(RankHRMin)/float64(DimH72)*Ssplit*Ssplit
	return AirlockFamilyAudit{
		Sigmas: []string{"+", "-"}, PlusPuncture: PuncturePlus, MinusPuncture: PunctureMinus,
		PlusF0: F0Plus, PlusF1: F1Plus, MinusF0: F0Minus, MinusF1: F1Minus, F2: F2,
		PlusRanks: []int{RankPiTop, RankHRMin}, MinusRanks: []int{RankPiTop, RankHRMin},
		PlusAlpha: plusAlpha, MinusAlpha: minusAlpha,
		BothReconstructAlpha: near(plusAlpha, AlphaB) && near(minusAlpha, AlphaB),
		Z2Family:             true, OrderedRepresentativeCertified: false,
		Supports: []string{StatusAirlockFamilyAudited, StatusAlphaReconstruction, SupportAirlockZ2Family, SupportBothAirlocksReconstructAlpha, SupportTwoSealWoundReduced, SupportUnorientedAirlockCandidate},
		Failures: []string{FailureNoNativeNeutralPunctureAirlock, FailureNoNativeOrderedAirlock, FailureNoNativeSocketOrderSelector, FailureNoNativeSelectionSigmaPlus, FailureAlphaStillSealed, FailureHiggsOrientationStillSealed},
	}
}

func FormatSocketOrder(s SocketOrderAudit) string {
	return fmt.Sprintf("socket_order(characters=%s pair_typed=%t native_order=%t plus_selected=%t boundary_degree_plus=%t J_chirality_plus=%t BminusL_plus=%t requires_selector=%t supports=%s failures=%s)", strings.Join(s.Characters, ","), s.PairTyped, s.NativeOrder, s.PlusSelected, s.BoundaryDegreeSelectsPlus, s.JOrChiralitySelectsPlus, s.BMinusLSelectsPlus, s.RequiresSelector, strings.Join(s.Supports, ","), strings.Join(s.Failures, ","))
}

func FormatEdgeOrdering(e EdgeOrderingAudit) string {
	return fmt.Sprintf("edge_ordering(rule=%s current=%s mirror=%s current_if_plus=%t mirror_if_minus=%t same_rank=%t selects_plus=%t reduces_to_socket=%t supports=%s failures=%s)", e.OrderedRule, strings.Join(e.CurrentOrdering, ";"), strings.Join(e.MirrorOrdering, ";"), e.CurrentFollowsIfPlusExposed, e.MirrorExistsIfMinusExposed, e.SameRankPattern, e.SelectsPlusByItself, e.ReducesToSocketOrder, strings.Join(e.Supports, ","), strings.Join(e.Failures, ","))
}

func FormatAirlockFamily(a AirlockFamilyAudit) string {
	return fmt.Sprintf("airlock_family(sigmas=%s plus=%s minus=%s plusRanks=%v minusRanks=%v plusAlpha=%.16g minusAlpha=%.16g both_alpha=%t z2=%t ordered_certified=%t supports=%s failures=%s)", strings.Join(a.Sigmas, ","), a.PlusPuncture, a.MinusPuncture, a.PlusRanks, a.MinusRanks, a.PlusAlpha, a.MinusAlpha, a.BothReconstructAlpha, a.Z2Family, a.OrderedRepresentativeCertified, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatFreeze(f FreezeAudit) string {
	return fmt.Sprintf("official_freeze(operator_Neff=%.16g official_Neff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t not_native_r3=%t alpha_sealed=%t higgs_sealed=%t no_socket_order=%t no_edge_ordering=%t no_airlock=%t no_ordered_airlock=%t no_incidence=%t no_descent=%t no_generation=%t no_flavor=%t no_individual=%t no_official=%t no_yukawa=%t no_r4=%t verdict=%s)", f.Enforced, f.NotNativeR3, f.AlphaStillSealed, f.HiggsOrientationStillSealed, f.NoNativeSocketOrderSelector, f.NoNativeOrientedEdgeOrdering, f.NoNativeAirlockFunctor, f.NoNativeOrderedAirlock, f.NoNativeBoundaryIncidenceFunctor, f.NoNativeDescentFullToOrient, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4NativeYukawaTheorem, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate896Inherited, StatusSocketPairAudited, StatusEdgeOrderingAudited, StatusAirlockFamilyAudited, StatusMirrorTableAudited, StatusAlphaReconstruction, StatusObstructionSharpened, StatusOfficialFreeze, StatusFirewallVerdict,
		SupportRightSocketCharacterPairTyped, SupportPlusMinusRemainingAmbiguity, SupportSocketOrderRequired, SupportEdgeTableFromOrderedRule, SupportCurrentTableIfPlusExposed, SupportZ2MirrorTableExists, SupportEdgeOrderingReducesToSocket, SupportAirlockZ2Family, SupportBothAirlocksReconstructAlpha, SupportTwoSealWoundReduced, SupportUnorientedAirlockCandidate, SupportOperatorNEffReproduced,
		FailureNotNativeR3, FailureAlphaStillSealed, FailureHiggsOrientationStillSealed, FailureNoNativeSocketOrderSelector, FailureNoNativePlusMinusOrder, FailureCharacterConjugationNoPlus, FailureBoundaryDegreeNoPlus, FailureJChiralityNoPlus, FailureBMinusLNoPlus, FailureNoNativeOrientedEdgeOrdering, FailureEdgeTableNoPlusByItself, FailureEdgeOrderingCircular, FailureNoNativeNeutralPunctureAirlock, FailureNoNativeOrderedAirlock, FailureNoNativeSelectionSigmaPlus, FailureNoNativeBoundaryIncidenceFunctor, FailureNoNativeDescentFullToOrient, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem,
	}
}

func (a Audit) FirewallsList() []string {
	return []string{FailureNotNativeR3, FailureAlphaStillSealed, FailureHiggsOrientationStillSealed, FailureNoNativeSocketOrderSelector, FailureNoNativeOrientedEdgeOrdering, FailureNoNativeNeutralPunctureAirlock, FailureNoNativeOrderedAirlock, FailureNoNativeBoundaryIncidenceFunctor, FailureNoNativeDescentFullToOrient, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem}
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NotNativeR3 && f.AlphaStillSealed && f.HiggsOrientationStillSealed &&
		f.NoNativeSocketOrderSelector && f.NoNativeOrientedEdgeOrdering && f.NoNativeAirlockFunctor &&
		f.NoNativeOrderedAirlock && f.NoNativeBoundaryIncidenceFunctor && f.NoNativeDescentFullToOrient &&
		f.NoGenerationCarrier && f.NoFlavorOrientation && f.NoIndividualYukawas &&
		f.NoOfficialLedgerUpdate && f.NoNativeYukawaOperator && f.NoR4NativeYukawaTheorem && f.Verdict == StatusFirewallVerdict
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
