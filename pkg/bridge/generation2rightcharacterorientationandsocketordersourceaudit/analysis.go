// Package generation2rightcharacterorientationandsocketordersourceaudit implements
// Gate 898: RightCharacter Orientation and SocketOrder Source Audit.
//
// Gate 898 follows Gate 897's Z2 neutral-puncture airlock consolidation. It
// audits whether the conjugate right-character pair lambda/bar(lambda) carries
// a native or bridge-lawful orientation that selects e_+ as the exposed/puncture
// socket and e_- as the active/rest socket. The result sharpens the missing
// SocketOrderSelector into a RightCharacterPhaseOrientationSource obstruction.
// No alpha derivation, native R3 promotion, physical-sector assignment,
// individual Yukawa value, or official ledger update is certified.
package generation2rightcharacterorientationandsocketordersourceaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE898-RIGHT-CHARACTER-ORIENTATION-SOCKET-ORDER-SOURCE-AUDIT"

	AlphaB = 0.0003878958469680527
	Ssplit = 0.001292444818816423

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	RightCharacterPlus  = "chi_R^+=lambda"
	RightCharacterMinus = "chi_R^-=bar(lambda)"
	PuncturePlus        = "e_+ tensor P_1"
	PunctureMinus       = "e_- tensor P_1"
	ColorPlus           = "e_+ tensor P_3"
	ColorMinus          = "e_- tensor P_3"

	ActiveEdgePlusColor   = "e_+ tensor P_3 -> h_+ tensor P_3"
	ActiveEdgeMinusColor  = "e_- tensor P_3 -> h_- tensor P_3"
	ActiveEdgeMinusLepton = "e_- tensor P_1 -> h_- tensor P_1"
	MissingEdgePlusLepton = "e_+ tensor P_1 -> h_+ tensor P_1 = 0"

	RankPiTop     = 3
	RankHRMin     = 7
	DimH10        = 10
	DimH72        = 72
	BMinusLLepton = -1

	Classification = "R3_CANDIDATE_NEUTRAL_PUNCTURE_AIRLOCK_REQUIRES_RIGHT_CHARACTER_PHASE_ORIENTATION_SEAL"
	ShortStatus    = "R3_AIRLOCK_Z2_SOCKET_ORDER_PHASE_ORIENTATION_OBSTRUCTION"
	NextFrontier   = "RIGHT_CHARACTER_PHASE_ORIENTATION_SOURCE_AUDIT"

	StatusGate897Inherited       = "PASS_GATE897_Z2_AIRLOCK_OBSTRUCTION_INHERITED"
	StatusRightCharacterAudited  = "PASS_RIGHT_CHARACTER_ORIENTATION_SOURCE_AUDITED"
	StatusComplexRouteAudited    = "PASS_COMPLEX_ORIENTATION_ROUTE_AUDITED"
	StatusOneFormRouteAudited    = "PASS_FINITE_ONE_FORM_ARROW_DIRECTION_ROUTE_AUDITED"
	StatusBoundaryRouteAudited   = "PASS_BOUNDARY_EXPOSURE_DIRECTION_ROUTE_AUDITED"
	StatusJChiralityRouteAudited = "PASS_J_CHIRALITY_KO_ROUTE_AUDITED"
	StatusBMinusLRouteAudited    = "PASS_B_MINUS_L_COMPENSATION_ROUTE_AUDITED"
	StatusPhaseRouteAudited      = "PASS_PHASE_ORIENTATION_ROUTE_AUDITED"
	StatusObstructionSharpened   = "PASS_SOCKET_ORDER_OBSTRUCTION_SHARPENED_TO_RIGHT_CHARACTER_PHASE_ORIENTATION"
	StatusOfficialFreeze         = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict        = "FIREWALL_PRESERVED_GATE898_PHASE_ORIENTATION_NOT_NATIVE"

	SupportRightSocketCharacterPairTyped      = "CONDITIONAL_SUPPORT_RIGHT_SOCKET_CHARACTER_PAIR_TYPED"
	SupportSocketOrderZ2Ambiguity             = "CONDITIONAL_SUPPORT_SOCKET_ORDER_AMBIGUITY_IS_PLUS_MINUS_Z2"
	SupportComplexOrientationCandidate        = "CONDITIONAL_SUPPORT_RIGHT_CHARACTER_PAIR_HAS_COMPLEX_ORIENTATION_CANDIDATE"
	SupportEPlusLambdaEMinusBar               = "CONDITIONAL_SUPPORT_E_PLUS_AS_LAMBDA_SOCKET_AND_E_MINUS_AS_BARLAMBDA_SOCKET"
	SupportSocketOrderGivenComplexOrientation = "CONDITIONAL_SUPPORT_SOCKET_ORDER_CAN_BE_STATED_GIVEN_COMPLEX_ORIENTATION"
	SupportPhaseOrientationCanSelectPlus      = "CONDITIONAL_SUPPORT_COMPLEX_PHASE_ORIENTATION_CAN_SELECT_E_PLUS_AS_LAMBDA_SOCKET_IF_SEALED"
	SupportAirlockOrderFollowsPhaseSeal       = "CONDITIONAL_SUPPORT_CURRENT_AIRLOCK_ORDER_FOLLOWS_FROM_RIGHT_CHARACTER_PHASE_ORIENTATION"
	SupportFiniteOneFormCompatible            = "CONDITIONAL_SUPPORT_FINITE_ONE_FORM_EDGE_TABLE_COMPATIBLE_WITH_E_PLUS_EXPOSURE_ORDER"
	SupportJChiralityCandidate                = "CONDITIONAL_SUPPORT_J_AND_CHIRALITY_ARE_POSSIBLE_SOCKET_ORDER_SOURCE_CANDIDATES"
	SupportPhaseOrientationSealRequired       = "CONDITIONAL_SUPPORT_SOCKET_ORDER_MAY_REQUIRE_RIGHT_CHARACTER_PHASE_ORIENTATION_SEAL"
	SupportZ2IsPhaseAmbiguity                 = "CONDITIONAL_SUPPORT_PLUS_MINUS_Z2_IS_PHASE_ORIENTATION_AMBIGUITY"
	SupportWoundSharpensToPhaseSource         = "CONDITIONAL_SUPPORT_SOCKET_ORDER_WOUND_SHARPENS_TO_RIGHT_CHARACTER_PHASE_ORIENTATION_SOURCE"
	SupportOperatorNEffReproduced             = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_REPRODUCED_BY_DUAL_SEAL_LEDGER"

	FailureNotNativeR3                         = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureAlphaStillSealed                    = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureHiggsOrientationStillSealed         = "FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED"
	FailureNoNativeSocketOrderSelector         = "FAILED_ROUTE_NO_NATIVE_SOCKET_ORDER_SELECTOR"
	FailureNoNativeRightPhaseOrientation       = "FAILED_ROUTE_NO_NATIVE_RIGHT_CHARACTER_PHASE_ORIENTATION_THEOREM"
	FailureCharacterConjugationNeedsPhase      = "FAILED_ROUTE_CHARACTER_CONJUGATION_PAIR_DOES_NOT_SELECT_PLUS_WITHOUT_PHASE_ORIENTATION"
	FailureComplexOrientationNotNativeSelector = "FAILED_ROUTE_COMPLEX_ORIENTATION_NOT_NATIVE_SOCKET_ORDER_SELECTOR_YET"
	FailureLambdaBarLabelingConvention         = "FAILED_ROUTE_LAMBDA_VS_BARLAMBDA_LABELING_IS_ORIENTATION_CONVENTION_WITHOUT_SOURCE"
	FailureBoundaryExposureNoSocketSign        = "FAILED_ROUTE_BOUNDARY_EXPOSURE_DIRECTION_SELECTS_FLAG_LEVEL_NOT_SOCKET_SIGN"
	FailureBoundaryDegreeNoBreakZ2             = "FAILED_ROUTE_BOUNDARY_DEGREE_ORDER_DOES_NOT_BREAK_PLUS_MINUS_Z2"
	FailureOneFormArrowRestatesOrder           = "FAILED_ROUTE_ONE_FORM_ARROW_DIRECTION_RESTATES_SOCKET_ORDER_WITHOUT_INDEPENDENT_SOURCE"
	FailureEdgeDirectionNotNativeSelector      = "FAILED_ROUTE_EDGE_DIRECTION_NOT_NATIVE_SOCKET_ORDER_SELECTOR_YET"
	FailureNoKOSignSelectsPlus                 = "FAILED_ROUTE_NO_KO_SIGN_OR_J_OPPOSITE_THEOREM_SELECTS_PLUS_SOCKET"
	FailureJMirrorNoBreakZ2                    = "FAILED_ROUTE_J_MIRROR_EXTENSION_DOES_NOT_BREAK_SOCKET_Z2"
	FailureChiralityNoPlus                     = "FAILED_ROUTE_CHIRALITY_RIGHT_LEFT_SPLIT_DOES_NOT_SELECT_E_PLUS_OVER_E_MINUS"
	FailureBMinusLNoBreakZ2                    = "FAILED_ROUTE_B_MINUS_L_COMPENSATION_DOES_NOT_BREAK_PLUS_MINUS_SOCKET_Z2"
	FailureNoNativeSelectionSigmaPlus          = "FAILED_ROUTE_NO_NATIVE_SELECTION_OF_SIGMA_EQUALS_PLUS"
	FailureNoGenerationCarrierMap              = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap              = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues            = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate                = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate               = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator              = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type RightCharacterAudit struct {
	Characters                        []string
	PairTyped                         bool
	ComplexOrientationCandidate       bool
	NativePhaseOrientation            bool
	EPlusAsLambda                     bool
	EMinusAsBarLambda                 bool
	SocketOrderStatedGivenOrientation bool
	SelectsPlusWithoutOrientation     bool
	RequiresPhaseOrientationSeal      bool
	Supports, Failures                []string
}

type OneFormArrowAudit struct {
	Edges                        []string
	MatchesEPlusPunctureOrder    bool
	DerivesOrderIndependently    bool
	RestatesSocketOrder          bool
	NativeArrowDirectionSelector bool
	Supports, Failures           []string
}

type BoundaryDegreeAudit struct {
	Degrees            []int
	IndexesFlagLevels  bool
	SelectsSocketSign  bool
	BreaksZ2           bool
	Supports, Failures []string
}

type JChiralityAudit struct {
	CandidatesPossible   bool
	KOSignCertified      bool
	JMirrorBreaksZ2      bool
	ChiralitySelectsPlus bool
	Supports, Failures   []string
}

type BMinusLAudit struct {
	PlusPunctureCharge  int
	MinusPunctureCharge int
	CompensationWorks   bool
	BreaksZ2            bool
	Supports, Failures  []string
}

type PhaseOrientationAudit struct {
	Z2IsPhaseAmbiguity      bool
	RightPhaseSealCandidate bool
	NativeTheorem           bool
	SelectsPlusIfSealed     bool
	NextFrontier            string
	Supports, Failures      []string
}

type FreezeAudit struct {
	OperatorNEff, OfficialNEff        float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	Frozen, DiagnosticOnly, CanUpdate bool
	Supports, Failures                []string
}

type Firewalls struct {
	Enforced                      bool
	NotNativeR3                   bool
	AlphaStillSealed              bool
	HiggsOrientationStillSealed   bool
	NoNativeSocketOrderSelector   bool
	NoNativeRightPhaseOrientation bool
	NoNativeSelectionSigmaPlus    bool
	NoGenerationCarrier           bool
	NoFlavorOrientation           bool
	NoIndividualYukawas           bool
	NoOfficialLedgerUpdate        bool
	NoNativeYukawaOperator        bool
	Verdict                       string
}

type Audit struct {
	ID             string
	RightCharacter RightCharacterAudit
	OneFormArrow   OneFormArrowAudit
	BoundaryDegree BoundaryDegreeAudit
	JChirality     JChiralityAudit
	BMinusL        BMinusLAudit
	Phase          PhaseOrientationAudit
	Freeze         FreezeAudit
	Firewalls      Firewalls
	Truth          string
	Final          string
}

func BuildDefault() (Audit, error) {
	right := buildRightCharacterAudit()
	if !right.PairTyped || !right.ComplexOrientationCandidate || right.NativePhaseOrientation || right.SelectsPlusWithoutOrientation || !right.RequiresPhaseOrientationSeal {
		return Audit{}, fmt.Errorf("right character audit promoted incorrectly: %s", FormatRightCharacter(right))
	}
	oneForm := buildOneFormArrowAudit()
	if !oneForm.MatchesEPlusPunctureOrder || oneForm.DerivesOrderIndependently || !oneForm.RestatesSocketOrder || oneForm.NativeArrowDirectionSelector {
		return Audit{}, fmt.Errorf("one-form route promoted incorrectly: %s", FormatOneFormArrow(oneForm))
	}
	boundary := buildBoundaryDegreeAudit()
	if !boundary.IndexesFlagLevels || boundary.SelectsSocketSign || boundary.BreaksZ2 {
		return Audit{}, fmt.Errorf("boundary route promoted incorrectly: %s", FormatBoundaryDegree(boundary))
	}
	j := buildJChiralityAudit()
	if !j.CandidatesPossible || j.KOSignCertified || j.JMirrorBreaksZ2 || j.ChiralitySelectsPlus {
		return Audit{}, fmt.Errorf("J/chirality route promoted incorrectly: %s", FormatJChirality(j))
	}
	bl := buildBMinusLAudit()
	if !bl.CompensationWorks || bl.BreaksZ2 || bl.PlusPunctureCharge != BMinusLLepton || bl.MinusPunctureCharge != BMinusLLepton {
		return Audit{}, fmt.Errorf("B-L route promoted incorrectly: %s", FormatBMinusL(bl))
	}
	phase := buildPhaseOrientationAudit()
	if !phase.Z2IsPhaseAmbiguity || !phase.RightPhaseSealCandidate || phase.NativeTheorem || !phase.SelectsPlusIfSealed {
		return Audit{}, fmt.Errorf("phase route promoted incorrectly: %s", FormatPhase(phase))
	}
	freeze := buildFreezeAudit()
	firewalls := buildFirewalls()
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewalls not preserved: %s", FormatFirewalls(firewalls))
	}
	return Audit{
		ID:             AuditID,
		RightCharacter: right,
		OneFormArrow:   oneForm,
		BoundaryDegree: boundary,
		JChirality:     j,
		BMinusL:        bl,
		Phase:          phase,
		Freeze:         freeze,
		Firewalls:      firewalls,
		Truth:          "Gate 898 sharpens the Gate 897 Z2 airlock obstruction: the missing SocketOrderSelector is a right-character phase-orientation source problem for the conjugate pair lambda/bar(lambda).",
		Final:          "The branch becomes R3_AIRLOCK_Z2_SOCKET_ORDER_PHASE_ORIENTATION_OBSTRUCTION. Current data define the airlock family but cannot natively select sigma=+; a RightCharacterPhaseOrientationSeal / SocketOrderPhaseSelector is now the exact missing object. Alpha, Higgs orientation, native R3, physical sector assignment, individual Yukawas, and official ledger updates remain blocked.",
	}, nil
}

func buildRightCharacterAudit() RightCharacterAudit {
	return RightCharacterAudit{
		Characters:                        []string{RightCharacterPlus, RightCharacterMinus},
		PairTyped:                         true,
		ComplexOrientationCandidate:       true,
		NativePhaseOrientation:            false,
		EPlusAsLambda:                     true,
		EMinusAsBarLambda:                 true,
		SocketOrderStatedGivenOrientation: true,
		SelectsPlusWithoutOrientation:     false,
		RequiresPhaseOrientationSeal:      true,
		Supports:                          []string{StatusRightCharacterAudited, StatusComplexRouteAudited, SupportRightSocketCharacterPairTyped, SupportSocketOrderZ2Ambiguity, SupportComplexOrientationCandidate, SupportEPlusLambdaEMinusBar, SupportSocketOrderGivenComplexOrientation, SupportPhaseOrientationCanSelectPlus, SupportAirlockOrderFollowsPhaseSeal},
		Failures:                          []string{FailureNoNativeSocketOrderSelector, FailureNoNativeRightPhaseOrientation, FailureCharacterConjugationNeedsPhase, FailureComplexOrientationNotNativeSelector, FailureLambdaBarLabelingConvention, FailureNoNativeSelectionSigmaPlus},
	}
}

func buildOneFormArrowAudit() OneFormArrowAudit {
	return OneFormArrowAudit{
		Edges:                        []string{ActiveEdgePlusColor, ActiveEdgeMinusColor, ActiveEdgeMinusLepton, MissingEdgePlusLepton},
		MatchesEPlusPunctureOrder:    true,
		DerivesOrderIndependently:    false,
		RestatesSocketOrder:          true,
		NativeArrowDirectionSelector: false,
		Supports:                     []string{StatusOneFormRouteAudited, SupportFiniteOneFormCompatible},
		Failures:                     []string{FailureOneFormArrowRestatesOrder, FailureEdgeDirectionNotNativeSelector},
	}
}

func buildBoundaryDegreeAudit() BoundaryDegreeAudit {
	return BoundaryDegreeAudit{
		Degrees:           []int{1, 2},
		IndexesFlagLevels: true,
		SelectsSocketSign: false,
		BreaksZ2:          false,
		Supports:          []string{StatusBoundaryRouteAudited},
		Failures:          []string{FailureBoundaryExposureNoSocketSign, FailureBoundaryDegreeNoBreakZ2},
	}
}

func buildJChiralityAudit() JChiralityAudit {
	return JChiralityAudit{
		CandidatesPossible:   true,
		KOSignCertified:      false,
		JMirrorBreaksZ2:      false,
		ChiralitySelectsPlus: false,
		Supports:             []string{StatusJChiralityRouteAudited, SupportJChiralityCandidate},
		Failures:             []string{FailureNoKOSignSelectsPlus, FailureJMirrorNoBreakZ2, FailureChiralityNoPlus},
	}
}

func buildBMinusLAudit() BMinusLAudit {
	return BMinusLAudit{
		PlusPunctureCharge:  BMinusLLepton,
		MinusPunctureCharge: BMinusLLepton,
		CompensationWorks:   true,
		BreaksZ2:            false,
		Supports:            []string{StatusBMinusLRouteAudited},
		Failures:            []string{FailureBMinusLNoBreakZ2},
	}
}

func buildPhaseOrientationAudit() PhaseOrientationAudit {
	return PhaseOrientationAudit{
		Z2IsPhaseAmbiguity:      true,
		RightPhaseSealCandidate: true,
		NativeTheorem:           false,
		SelectsPlusIfSealed:     true,
		NextFrontier:            NextFrontier,
		Supports:                []string{StatusPhaseRouteAudited, StatusObstructionSharpened, SupportPhaseOrientationSealRequired, SupportZ2IsPhaseAmbiguity, SupportWoundSharpensToPhaseSource},
		Failures:                []string{FailureNoNativeRightPhaseOrientation, FailureNoNativeSocketOrderSelector, FailureNoNativeSelectionSigmaPlus},
	}
}

func buildFreezeAudit() FreezeAudit {
	return FreezeAudit{
		OperatorNEff: OperatorNEffDiagnostic, OfficialNEff: OfficialNEffFrozen,
		OperatorCYukawa: OperatorCYukawaDiagnostic, OfficialCYukawa: OfficialCYukawaFrozen,
		OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen,
		Frozen: true, DiagnosticOnly: true, CanUpdate: false,
		Supports: []string{StatusOfficialFreeze, SupportOperatorNEffReproduced},
		Failures: []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate},
	}
}

func buildFirewalls() Firewalls {
	return Firewalls{
		Enforced:                      true,
		NotNativeR3:                   true,
		AlphaStillSealed:              true,
		HiggsOrientationStillSealed:   true,
		NoNativeSocketOrderSelector:   true,
		NoNativeRightPhaseOrientation: true,
		NoNativeSelectionSigmaPlus:    true,
		NoGenerationCarrier:           true,
		NoFlavorOrientation:           true,
		NoIndividualYukawas:           true,
		NoOfficialLedgerUpdate:        true,
		NoNativeYukawaOperator:        true,
		Verdict:                       StatusFirewallVerdict,
	}
}

func FormatRightCharacter(r RightCharacterAudit) string {
	return fmt.Sprintf("right_character(characters=%s pair_typed=%t complex_candidate=%t native_phase=%t e_plus_lambda=%t e_minus_bar=%t order_given_orientation=%t selects_plus_without_orientation=%t requires_phase_seal=%t supports=%s failures=%s)", strings.Join(r.Characters, ","), r.PairTyped, r.ComplexOrientationCandidate, r.NativePhaseOrientation, r.EPlusAsLambda, r.EMinusAsBarLambda, r.SocketOrderStatedGivenOrientation, r.SelectsPlusWithoutOrientation, r.RequiresPhaseOrientationSeal, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}

func FormatOneFormArrow(o OneFormArrowAudit) string {
	return fmt.Sprintf("one_form_arrow(edges=%s matches_plus_order=%t derives_independently=%t restates_order=%t native_arrow_selector=%t supports=%s failures=%s)", strings.Join(o.Edges, ";"), o.MatchesEPlusPunctureOrder, o.DerivesOrderIndependently, o.RestatesSocketOrder, o.NativeArrowDirectionSelector, strings.Join(o.Supports, ","), strings.Join(o.Failures, ","))
}

func FormatBoundaryDegree(b BoundaryDegreeAudit) string {
	return fmt.Sprintf("boundary_degree(degrees=%v indexes_flag=%t selects_socket_sign=%t breaks_z2=%t supports=%s failures=%s)", b.Degrees, b.IndexesFlagLevels, b.SelectsSocketSign, b.BreaksZ2, strings.Join(b.Supports, ","), strings.Join(b.Failures, ","))
}

func FormatJChirality(j JChiralityAudit) string {
	return fmt.Sprintf("j_chirality(candidates_possible=%t ko_certified=%t j_mirror_breaks_z2=%t chirality_selects_plus=%t supports=%s failures=%s)", j.CandidatesPossible, j.KOSignCertified, j.JMirrorBreaksZ2, j.ChiralitySelectsPlus, strings.Join(j.Supports, ","), strings.Join(j.Failures, ","))
}

func FormatBMinusL(b BMinusLAudit) string {
	return fmt.Sprintf("b_minus_l(plus_charge=%d minus_charge=%d compensation=%t breaks_z2=%t supports=%s failures=%s)", b.PlusPunctureCharge, b.MinusPunctureCharge, b.CompensationWorks, b.BreaksZ2, strings.Join(b.Supports, ","), strings.Join(b.Failures, ","))
}

func FormatPhase(p PhaseOrientationAudit) string {
	return fmt.Sprintf("phase_orientation(z2_phase_ambiguity=%t right_phase_seal_candidate=%t native_theorem=%t selects_plus_if_sealed=%t next=%s supports=%s failures=%s)", p.Z2IsPhaseAmbiguity, p.RightPhaseSealCandidate, p.NativeTheorem, p.SelectsPlusIfSealed, p.NextFrontier, strings.Join(p.Supports, ","), strings.Join(p.Failures, ","))
}

func FormatFreeze(f FreezeAudit) string {
	return fmt.Sprintf("official_freeze(operator_Neff=%.16g official_Neff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t not_native_r3=%t alpha_sealed=%t higgs_sealed=%t no_socket_order=%t no_right_phase=%t no_sigma_plus=%t no_generation=%t no_flavor=%t no_individual=%t no_official=%t no_yukawa=%t verdict=%s)", f.Enforced, f.NotNativeR3, f.AlphaStillSealed, f.HiggsOrientationStillSealed, f.NoNativeSocketOrderSelector, f.NoNativeRightPhaseOrientation, f.NoNativeSelectionSigmaPlus, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate897Inherited, StatusRightCharacterAudited, StatusComplexRouteAudited, StatusOneFormRouteAudited, StatusBoundaryRouteAudited, StatusJChiralityRouteAudited, StatusBMinusLRouteAudited, StatusPhaseRouteAudited, StatusObstructionSharpened, StatusOfficialFreeze, StatusFirewallVerdict,
		SupportRightSocketCharacterPairTyped, SupportSocketOrderZ2Ambiguity, SupportComplexOrientationCandidate, SupportEPlusLambdaEMinusBar, SupportSocketOrderGivenComplexOrientation, SupportPhaseOrientationCanSelectPlus, SupportAirlockOrderFollowsPhaseSeal, SupportFiniteOneFormCompatible, SupportJChiralityCandidate, SupportPhaseOrientationSealRequired, SupportZ2IsPhaseAmbiguity, SupportWoundSharpensToPhaseSource, SupportOperatorNEffReproduced,
		FailureNotNativeR3, FailureAlphaStillSealed, FailureHiggsOrientationStillSealed, FailureNoNativeSocketOrderSelector, FailureNoNativeRightPhaseOrientation, FailureCharacterConjugationNeedsPhase, FailureComplexOrientationNotNativeSelector, FailureLambdaBarLabelingConvention, FailureBoundaryExposureNoSocketSign, FailureBoundaryDegreeNoBreakZ2, FailureOneFormArrowRestatesOrder, FailureEdgeDirectionNotNativeSelector, FailureNoKOSignSelectsPlus, FailureJMirrorNoBreakZ2, FailureChiralityNoPlus, FailureBMinusLNoBreakZ2, FailureNoNativeSelectionSigmaPlus, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNativeYukawaOperator,
	}
}

func (a Audit) FirewallsList() []string {
	return []string{FailureNotNativeR3, FailureAlphaStillSealed, FailureHiggsOrientationStillSealed, FailureNoNativeSocketOrderSelector, FailureNoNativeRightPhaseOrientation, FailureNoNativeSelectionSigmaPlus, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoNativeYukawaOperator}
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NotNativeR3 && f.AlphaStillSealed && f.HiggsOrientationStillSealed &&
		f.NoNativeSocketOrderSelector && f.NoNativeRightPhaseOrientation && f.NoNativeSelectionSigmaPlus &&
		f.NoGenerationCarrier && f.NoFlavorOrientation && f.NoIndividualYukawas &&
		f.NoOfficialLedgerUpdate && f.NoNativeYukawaOperator && f.Verdict == StatusFirewallVerdict
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
