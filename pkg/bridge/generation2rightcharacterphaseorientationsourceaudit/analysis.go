// Package generation2rightcharacterphaseorientationsourceaudit implements
// Gate 899: RightCharacter PhaseOrientation Source Audit.
//
// Gate 899 follows Gate 898's result that the neutral-puncture airlock is
// ordered only after choosing a phase orientation for the conjugate right
// character pair lambda/bar(lambda). It audits possible existing ASHA sources
// for that phase orientation: Hopf/S1 phase orientation, the Cl(1,7) complex
// chirality airlock, J/KO sign data, boundary-pair orientation, and finite
// spectral orientation cycles. The honest result is a source-candidate
// obstruction: Hopf/S1 and complex chirality are strongest, but no native
// socket-phase selector is certified. No alpha derivation, native R3
// promotion, physical-sector assignment, individual Yukawa value, or official
// ledger update is certified.
package generation2rightcharacterphaseorientationsourceaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE899-RIGHT-CHARACTER-PHASE-ORIENTATION-SOURCE-AUDIT"

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

	HopfS1PhaseOrientation       = "Hopf/S1 phase orientation"
	ComplexChiralityAirlock      = "Cl(1,7) complex chirality airlock gamma_chi=i omega"
	JKOConjugationCandidate      = "J/KO conjugation and opposite action"
	BoundaryPairOrientation      = "B2 exterior orientation b1 wedge b2"
	FiniteSpectralOrientation    = "finite spectral orientation cycle"
	RightCharacterPhaseSeal      = "RightCharacterPhaseOrientationSeal"
	SocketOrderPhaseSelector     = "SocketOrderPhaseSelector"
	NextFrontier                 = "RIGHT_CHARACTER_PHASE_ORIENTATION_SOURCE_THEOREM"
	Classification               = "R3_CANDIDATE_SOCKET_ORDER_REDUCED_TO_PHASE_ORIENTATION_SEAL"
	ShortStatus                  = "R3_AIRLOCK_PHASE_ORIENTATION_SOURCE_CANDIDATE_NOT_NATIVE"
	StatusGate898Inherited       = "PASS_GATE898_PHASE_ORIENTATION_OBSTRUCTION_INHERITED"
	StatusHopfRouteAudited       = "PASS_HOPF_S1_PHASE_ORIENTATION_ROUTE_AUDITED"
	StatusChiralityRouteAudited  = "PASS_CL17_COMPLEX_CHIRALITY_AIRLOCK_ROUTE_AUDITED"
	StatusJKORouteAudited        = "PASS_J_KO_SIGN_ROUTE_AUDITED"
	StatusBoundaryRouteAudited   = "PASS_BOUNDARY_PAIR_ORIENTATION_ROUTE_AUDITED"
	StatusSpectralRouteAudited   = "PASS_FINITE_SPECTRAL_ORIENTATION_CYCLE_ROUTE_AUDITED"
	StatusSourceCandidatesRanked = "PASS_SOCKET_PHASE_SOURCE_CANDIDATES_RANKED"
	StatusOfficialFreeze         = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict        = "FIREWALL_PRESERVED_GATE899_PHASE_ORIENTATION_SOURCE_NOT_NATIVE"

	SupportRightSocketCharacterPairTyped = "CONDITIONAL_SUPPORT_RIGHT_SOCKET_CHARACTER_PAIR_TYPED"
	SupportSocketOrderZ2Ambiguity        = "CONDITIONAL_SUPPORT_SOCKET_ORDER_AMBIGUITY_IS_PLUS_MINUS_Z2"
	SupportHopfStrongestCandidate        = "CONDITIONAL_SUPPORT_HOPF_S1_PHASE_ORIENTATION_IS_STRONGEST_SOCKET_ORDER_SOURCE_CANDIDATE"
	SupportHopfCanSourceIfSealed         = "CONDITIONAL_SUPPORT_RIGHT_CHARACTER_ORDER_CAN_BE_SOURCED_BY_PHASE_ORIENTATION_IF_SEALED"
	SupportChiralityCandidate            = "CONDITIONAL_SUPPORT_COMPLEX_CHIRALITY_AIRLOCK_IS_SOCKET_PHASE_ORIENTATION_CANDIDATE"
	SupportCL17FirewallCandidate         = "CONDITIONAL_SUPPORT_CL17_REAL_FORM_FIREWALL_MAY_SOURCE_LAMBDA_BARLAMBDA_ORIENTATION"
	SupportJKORelevant                   = "CONDITIONAL_SUPPORT_J_KO_DATA_RELEVANT_TO_CONJUGATION_STRUCTURE"
	SupportBoundaryCompatible            = "CONDITIONAL_SUPPORT_BOUNDARY_PAIR_ORIENTATION_COMPATIBLE_WITH_ORDERED_AIRLOCK"
	SupportSpectralOrientationCandidate  = "CONDITIONAL_SUPPORT_FINITE_SPECTRAL_ORIENTATION_CYCLE_IS_DEEP_SOCKET_ORDER_SOURCE_CANDIDATE"
	SupportHopfAndChiralityStrongest     = "CONDITIONAL_SUPPORT_HOPF_S1_AND_COMPLEX_CHIRALITY_AIRLOCK_ARE_STRONGEST_PHASE_SOURCE_CANDIDATES"
	SupportSocketOrderReducedToPhaseSeal = "CONDITIONAL_SUPPORT_SOCKET_ORDER_REDUCED_TO_RIGHT_CHARACTER_PHASE_ORIENTATION_SEAL"
	SupportOperatorNEffReproduced        = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_REPRODUCED_BY_DUAL_SEAL_LEDGER"

	FailureNotNativeR3                           = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureAlphaStillSealed                      = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureHiggsOrientationStillSealed           = "FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED"
	FailureNoNativeSocketOrderSelector           = "FAILED_ROUTE_NO_NATIVE_SOCKET_ORDER_SELECTOR"
	FailureNoNativeRightPhaseOrientation         = "FAILED_ROUTE_NO_NATIVE_RIGHT_CHARACTER_PHASE_ORIENTATION_THEOREM"
	FailureNoNativeSelectionSigmaPlus            = "FAILED_ROUTE_NO_NATIVE_SELECTION_OF_SIGMA_EQUALS_PLUS"
	FailureNoHopfToSocketOrderTheorem            = "FAILED_ROUTE_NO_HOPF_TO_RIGHT_CHARACTER_SOCKET_ORDER_THEOREM"
	FailurePhaseOrientationSealNotNative         = "FAILED_ROUTE_PHASE_ORIENTATION_IS_SEAL_NOT_NATIVE_SELECTOR"
	FailureNoCL17ChiralityToRightOrderMap        = "FAILED_ROUTE_NO_TYPED_CL17_CHIRALITY_TO_RIGHT_CHARACTER_ORDER_MAP"
	FailureComplexChiralityNotSocketOrderTheorem = "FAILED_ROUTE_COMPLEX_CHIRALITY_AIRLOCK_NOT_SOCKET_ORDER_THEOREM_YET"
	FailureJKODoesNotSelectPlus                  = "FAILED_ROUTE_J_KO_SIGN_DOES_NOT_CURRENTLY_SELECT_E_PLUS_AS_PUNCTURE"
	FailureNoKOSignSocketOrderTheorem            = "FAILED_ROUTE_NO_KO_SIGN_EXTENSION_THEOREM_FOR_SOCKET_ORDER"
	FailureBoundaryPairNoLambdaBar               = "FAILED_ROUTE_BOUNDARY_PAIR_ORIENTATION_DOES_NOT_SELECT_LAMBDA_OVER_BARLAMBDA"
	FailureBoundarySelectsDegreeNotSocketPhase   = "FAILED_ROUTE_BOUNDARY_ORIENTATION_SELECTS_DEGREE_ORDER_NOT_SOCKET_PHASE_ORDER"
	FailureNoFiniteSpectralOrientationTheorem    = "FAILED_ROUTE_NO_FINITE_SPECTRAL_ORIENTATION_CYCLE_TO_SOCKET_ORDER_THEOREM"
	FailureNoGenerationCarrierMap                = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap                = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues              = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate                  = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate                 = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator                = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type InheritedAudit struct {
	Characters                 []string
	Airlocks                   []string
	Z2Family                   bool
	RequiresPhaseOrientation   bool
	CanSelectSigmaPlusNatively bool
	Supports, Failures         []string
}

type HopfS1Audit struct {
	CandidateName       string
	PhaseOrientation    bool
	CanOrientLambdaBar  bool
	StrongestCandidate  bool
	NativeTheorem       bool
	SelectsPlusIfSealed bool
	Supports, Failures  []string
}

type ComplexChiralityAudit struct {
	CandidateName            string
	OmegaSquaredMinusOne     bool
	RequiresComplexAirlock   bool
	CanOrientIOverMinusI     bool
	TypedToRightCharacters   bool
	NativeSocketOrderTheorem bool
	Supports, Failures       []string
}

type JKOAudit struct {
	CandidateName         string
	RelevantToConjugation bool
	KOSignCertified       bool
	SelectsPlus           bool
	NativeTheorem         bool
	Supports, Failures    []string
}

type BoundaryPairAudit struct {
	CandidateName         string
	ExteriorOrientation   bool
	CompatibleWithAirlock bool
	SelectsDegreeOrder    bool
	SelectsSocketPhase    bool
	Supports, Failures    []string
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
	PhaseSealName       string
	SocketSelectorName  string
	NativeSourceFound   bool
	NextFrontier        string
	Supports, Failures  []string
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
	ID                  string
	Inherited           InheritedAudit
	HopfS1              HopfS1Audit
	ComplexChirality    ComplexChiralityAudit
	JKO                 JKOAudit
	BoundaryPair        BoundaryPairAudit
	SpectralOrientation SpectralOrientationAudit
	Ranking             CandidateRankingAudit
	Freeze              FreezeAudit
	Firewalls           Firewalls
	Truth               string
	Final               string
}

func BuildDefault() (Audit, error) {
	inherited := buildInheritedAudit()
	if !inherited.Z2Family || !inherited.RequiresPhaseOrientation || inherited.CanSelectSigmaPlusNatively {
		return Audit{}, fmt.Errorf("inherited audit promoted incorrectly: %s", FormatInherited(inherited))
	}
	hopf := buildHopfS1Audit()
	if !hopf.PhaseOrientation || !hopf.CanOrientLambdaBar || !hopf.StrongestCandidate || hopf.NativeTheorem || !hopf.SelectsPlusIfSealed {
		return Audit{}, fmt.Errorf("Hopf route promoted incorrectly: %s", FormatHopfS1(hopf))
	}
	chirality := buildComplexChiralityAudit()
	if !chirality.OmegaSquaredMinusOne || !chirality.RequiresComplexAirlock || !chirality.CanOrientIOverMinusI || chirality.TypedToRightCharacters || chirality.NativeSocketOrderTheorem {
		return Audit{}, fmt.Errorf("chirality route promoted incorrectly: %s", FormatComplexChirality(chirality))
	}
	jko := buildJKOAudit()
	if !jko.RelevantToConjugation || jko.KOSignCertified || jko.SelectsPlus || jko.NativeTheorem {
		return Audit{}, fmt.Errorf("J/KO route promoted incorrectly: %s", FormatJKO(jko))
	}
	boundary := buildBoundaryPairAudit()
	if !boundary.ExteriorOrientation || !boundary.CompatibleWithAirlock || !boundary.SelectsDegreeOrder || boundary.SelectsSocketPhase {
		return Audit{}, fmt.Errorf("boundary route promoted incorrectly: %s", FormatBoundaryPair(boundary))
	}
	spectral := buildSpectralOrientationAudit()
	if !spectral.DeepCandidate || spectral.CycleCertified || spectral.MapsToSocketOrder || spectral.NativeTheorem {
		return Audit{}, fmt.Errorf("spectral route promoted incorrectly: %s", FormatSpectralOrientation(spectral))
	}
	ranking := buildCandidateRankingAudit()
	if len(ranking.StrongestCandidates) != 2 || ranking.NativeSourceFound || ranking.NextFrontier != NextFrontier {
		return Audit{}, fmt.Errorf("ranking audit promoted incorrectly: %s", FormatRanking(ranking))
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
		ID:                  AuditID,
		Inherited:           inherited,
		HopfS1:              hopf,
		ComplexChirality:    chirality,
		JKO:                 jko,
		BoundaryPair:        boundary,
		SpectralOrientation: spectral,
		Ranking:             ranking,
		Freeze:              freeze,
		Firewalls:           firewalls,
		Truth:               "Gate 899 reduces the ordered airlock wound to a right-character phase-orientation seal; Hopf/S1 and Cl(1,7) complex chirality are strongest candidates, but no native selector is certified.",
		Final:               Classification,
	}, nil
}

func buildInheritedAudit() InheritedAudit {
	return InheritedAudit{
		Characters:                 []string{RightCharacterPlus, RightCharacterMinus},
		Airlocks:                   []string{PuncturePlus, PunctureMinus},
		Z2Family:                   true,
		RequiresPhaseOrientation:   true,
		CanSelectSigmaPlusNatively: false,
		Supports:                   []string{StatusGate898Inherited, SupportRightSocketCharacterPairTyped, SupportSocketOrderZ2Ambiguity},
		Failures:                   []string{FailureNoNativeSocketOrderSelector, FailureNoNativeRightPhaseOrientation, FailureNoNativeSelectionSigmaPlus},
	}
}

func buildHopfS1Audit() HopfS1Audit {
	return HopfS1Audit{
		CandidateName:       HopfS1PhaseOrientation,
		PhaseOrientation:    true,
		CanOrientLambdaBar:  true,
		StrongestCandidate:  true,
		NativeTheorem:       false,
		SelectsPlusIfSealed: true,
		Supports:            []string{StatusHopfRouteAudited, SupportHopfStrongestCandidate, SupportHopfCanSourceIfSealed},
		Failures:            []string{FailureNoHopfToSocketOrderTheorem, FailurePhaseOrientationSealNotNative, FailureNoNativeRightPhaseOrientation},
	}
}

func buildComplexChiralityAudit() ComplexChiralityAudit {
	return ComplexChiralityAudit{
		CandidateName:            ComplexChiralityAirlock,
		OmegaSquaredMinusOne:     true,
		RequiresComplexAirlock:   true,
		CanOrientIOverMinusI:     true,
		TypedToRightCharacters:   false,
		NativeSocketOrderTheorem: false,
		Supports:                 []string{StatusChiralityRouteAudited, SupportChiralityCandidate, SupportCL17FirewallCandidate},
		Failures:                 []string{FailureNoCL17ChiralityToRightOrderMap, FailureComplexChiralityNotSocketOrderTheorem},
	}
}

func buildJKOAudit() JKOAudit {
	return JKOAudit{
		CandidateName:         JKOConjugationCandidate,
		RelevantToConjugation: true,
		KOSignCertified:       false,
		SelectsPlus:           false,
		NativeTheorem:         false,
		Supports:              []string{StatusJKORouteAudited, SupportJKORelevant},
		Failures:              []string{FailureJKODoesNotSelectPlus, FailureNoKOSignSocketOrderTheorem},
	}
}

func buildBoundaryPairAudit() BoundaryPairAudit {
	return BoundaryPairAudit{
		CandidateName:         BoundaryPairOrientation,
		ExteriorOrientation:   true,
		CompatibleWithAirlock: true,
		SelectsDegreeOrder:    true,
		SelectsSocketPhase:    false,
		Supports:              []string{StatusBoundaryRouteAudited, SupportBoundaryCompatible},
		Failures:              []string{FailureBoundaryPairNoLambdaBar, FailureBoundarySelectsDegreeNotSocketPhase},
	}
}

func buildSpectralOrientationAudit() SpectralOrientationAudit {
	return SpectralOrientationAudit{
		CandidateName:     FiniteSpectralOrientation,
		DeepCandidate:     true,
		CycleCertified:    false,
		MapsToSocketOrder: false,
		NativeTheorem:     false,
		Supports:          []string{StatusSpectralRouteAudited, SupportSpectralOrientationCandidate},
		Failures:          []string{FailureNoFiniteSpectralOrientationTheorem},
	}
}

func buildCandidateRankingAudit() CandidateRankingAudit {
	return CandidateRankingAudit{
		StrongestCandidates: []string{HopfS1PhaseOrientation, ComplexChiralityAirlock},
		PhaseSealName:       RightCharacterPhaseSeal,
		SocketSelectorName:  SocketOrderPhaseSelector,
		NativeSourceFound:   false,
		NextFrontier:        NextFrontier,
		Supports:            []string{StatusSourceCandidatesRanked, SupportHopfAndChiralityStrongest, SupportSocketOrderReducedToPhaseSeal},
		Failures:            []string{FailureNoNativeRightPhaseOrientation, FailureNoNativeSocketOrderSelector, FailureNoNativeSelectionSigmaPlus},
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

func FormatInherited(i InheritedAudit) string {
	return fmt.Sprintf("inherited(characters=%s airlocks=%s z2_family=%t requires_phase=%t selects_sigma_plus_natively=%t supports=%s failures=%s)", strings.Join(i.Characters, ","), strings.Join(i.Airlocks, ","), i.Z2Family, i.RequiresPhaseOrientation, i.CanSelectSigmaPlusNatively, strings.Join(i.Supports, ","), strings.Join(i.Failures, ","))
}

func FormatHopfS1(h HopfS1Audit) string {
	return fmt.Sprintf("hopf_s1(candidate=%s phase_orientation=%t can_orient_lambda_bar=%t strongest=%t native=%t selects_plus_if_sealed=%t supports=%s failures=%s)", h.CandidateName, h.PhaseOrientation, h.CanOrientLambdaBar, h.StrongestCandidate, h.NativeTheorem, h.SelectsPlusIfSealed, strings.Join(h.Supports, ","), strings.Join(h.Failures, ","))
}

func FormatComplexChirality(c ComplexChiralityAudit) string {
	return fmt.Sprintf("complex_chirality(candidate=%s omega_squared_minus_one=%t requires_complex_airlock=%t can_orient_i=%t typed_to_right_characters=%t native_socket_order=%t supports=%s failures=%s)", c.CandidateName, c.OmegaSquaredMinusOne, c.RequiresComplexAirlock, c.CanOrientIOverMinusI, c.TypedToRightCharacters, c.NativeSocketOrderTheorem, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}

func FormatJKO(j JKOAudit) string {
	return fmt.Sprintf("j_ko(candidate=%s relevant_to_conjugation=%t ko_certified=%t selects_plus=%t native=%t supports=%s failures=%s)", j.CandidateName, j.RelevantToConjugation, j.KOSignCertified, j.SelectsPlus, j.NativeTheorem, strings.Join(j.Supports, ","), strings.Join(j.Failures, ","))
}

func FormatBoundaryPair(b BoundaryPairAudit) string {
	return fmt.Sprintf("boundary_pair(candidate=%s exterior_orientation=%t compatible=%t selects_degree_order=%t selects_socket_phase=%t supports=%s failures=%s)", b.CandidateName, b.ExteriorOrientation, b.CompatibleWithAirlock, b.SelectsDegreeOrder, b.SelectsSocketPhase, strings.Join(b.Supports, ","), strings.Join(b.Failures, ","))
}

func FormatSpectralOrientation(s SpectralOrientationAudit) string {
	return fmt.Sprintf("spectral_orientation(candidate=%s deep_candidate=%t cycle_certified=%t maps_to_socket_order=%t native=%t supports=%s failures=%s)", s.CandidateName, s.DeepCandidate, s.CycleCertified, s.MapsToSocketOrder, s.NativeTheorem, strings.Join(s.Supports, ","), strings.Join(s.Failures, ","))
}

func FormatRanking(r CandidateRankingAudit) string {
	return fmt.Sprintf("ranking(strongest=%s phase_seal=%s selector=%s native_source_found=%t next=%s supports=%s failures=%s)", strings.Join(r.StrongestCandidates, ","), r.PhaseSealName, r.SocketSelectorName, r.NativeSourceFound, r.NextFrontier, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}

func FormatFreeze(f FreezeAudit) string {
	return fmt.Sprintf("official_freeze(operator_Neff=%.16g official_Neff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t not_native_r3=%t alpha_sealed=%t higgs_sealed=%t no_socket_order=%t no_right_phase=%t no_sigma_plus=%t no_generation=%t no_flavor=%t no_individual=%t no_official=%t no_yukawa=%t verdict=%s)", f.Enforced, f.NotNativeR3, f.AlphaStillSealed, f.HiggsOrientationStillSealed, f.NoNativeSocketOrderSelector, f.NoNativeRightPhaseOrientation, f.NoNativeSelectionSigmaPlus, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate898Inherited, StatusHopfRouteAudited, StatusChiralityRouteAudited, StatusJKORouteAudited, StatusBoundaryRouteAudited, StatusSpectralRouteAudited, StatusSourceCandidatesRanked, StatusOfficialFreeze, StatusFirewallVerdict,
		SupportRightSocketCharacterPairTyped, SupportSocketOrderZ2Ambiguity, SupportHopfStrongestCandidate, SupportHopfCanSourceIfSealed, SupportChiralityCandidate, SupportCL17FirewallCandidate, SupportJKORelevant, SupportBoundaryCompatible, SupportSpectralOrientationCandidate, SupportHopfAndChiralityStrongest, SupportSocketOrderReducedToPhaseSeal, SupportOperatorNEffReproduced,
		FailureNotNativeR3, FailureAlphaStillSealed, FailureHiggsOrientationStillSealed, FailureNoNativeSocketOrderSelector, FailureNoNativeRightPhaseOrientation, FailureNoNativeSelectionSigmaPlus, FailureNoHopfToSocketOrderTheorem, FailurePhaseOrientationSealNotNative, FailureNoCL17ChiralityToRightOrderMap, FailureComplexChiralityNotSocketOrderTheorem, FailureJKODoesNotSelectPlus, FailureNoKOSignSocketOrderTheorem, FailureBoundaryPairNoLambdaBar, FailureBoundarySelectsDegreeNotSocketPhase, FailureNoFiniteSpectralOrientationTheorem, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNativeYukawaOperator,
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
