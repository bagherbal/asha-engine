// Package generation2socketcharacteridentificationedgeintertwinerpromotionaudit implements
// Gate 862: SocketCharacter Identification and Edge-Intertwiner Promotion Audit.
//
// Gate 862 follows Gate 861's stabilizer-branch first-order operator seal. Gate
// 861 removed the color obstruction by requiring the color edges to be scalar on
// P_3, but it left the socket matching
//
//	e_+ -> h_+,
//	e_- -> h_-
//
// as orientation-seal data. Gate 862 audits whether the right character sockets
// of C_R and the oriented weak sockets of C_H can be identified strongly enough
// for the scalar edge maps to be operator-level intertwiners inside
//
//	A_F^orient = C_R plus C_H plus M_3(C).
//
// The result is constructive only at seal level: the edge intertwiners hold if a
// C_R -> C_H character-identification seal is admitted. This does not promote
// the construction to a native finite triple, a full unbroken A_F theorem, a
// Yukawa-magnitude theorem, R3/R4, or an official ledger update.
package generation2socketcharacteridentificationedgeintertwinerpromotionaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE862-SOCKET-CHARACTER-IDENTIFICATION-EDGE-INTERTWINER-PROMOTION-AUDIT"

	AlphaB          = 0.0003878958469680527
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	P1Rank       = 1
	P3Rank       = 3
	WRank        = P1Rank + P3Rank
	HLRank       = 2 * WRank
	HRMinRank    = 7
	HPartMinRank = HLRank + HRMinRank
	YRankFull    = HRMinRank
	DSymRankFull = 2 * YRankFull
	KernelRank   = HPartMinRank - DSymRankFull

	StatusGate861Inherited       = "PASS_GATE861_STABILIZER_FIRST_ORDER_OPERATOR_SEAL_INHERITED"
	StatusRightCharacterLedger   = "PASS_RIGHT_SOCKET_CHARACTER_LEDGER_AUDITED"
	StatusWeakCharacterLedger    = "PASS_ORIENTED_WEAK_SOCKET_CHARACTER_LEDGER_AUDITED"
	StatusCharacterMapFormulated = "PASS_C_R_TO_C_H_CHARACTER_IDENTIFICATION_MAP_FORMULATED"
	StatusIntertwinerConditions  = "PASS_EDGE_INTERTWINER_CONDITIONS_FORMULATED"
	StatusActiveIntertwiners     = "PASS_ACTIVE_EDGE_INTERTWINERS_HOLD_GIVEN_CHARACTER_IDENTIFICATION_SEAL"
	StatusPunctureZeroReaudited  = "PASS_PUNCTURE_EDGE_ZERO_REAUDITED"
	StatusFirstOrderSharpened    = "PASS_STABILIZER_FIRST_ORDER_OPERATOR_COMPATIBILITY_SHARPENED"
	StatusNoObservedDataUsed     = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusLedgerFrozen           = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusFirewallVerdict        = "FIREWALL_PRESERVED_GATE862_CHARACTER_IDENTIFICATION_NOT_R3"

	SupportEdgeIntertwinersGivenID = "CONDITIONAL_SUPPORT_EDGE_INTERTWINERS_HOLD_GIVEN_C_R_TO_C_H_CHARACTER_IDENTIFICATION"
	SupportSocketMatchOrientation  = "CONDITIONAL_SUPPORT_SOCKET_CHARACTER_MATCHING_BY_ORIENTATION_SEAL"
	SupportFirstOrderSharpened     = "CONDITIONAL_SUPPORT_STABILIZER_FIRST_ORDER_OPERATOR_COMPATIBILITY_SHARPENED"
	SupportPunctureNotForced       = "CONDITIONAL_SUPPORT_CHARACTER_MATCHING_DOES_NOT_FORCE_Y_PLUS1"
	SupportRightCharsTyped         = "CONDITIONAL_SUPPORT_RIGHT_SOCKET_CHARACTERS_TYPED_BY_LAMBDA_AND_BARLAMBDA"
	SupportWeakCharsOriented       = "CONDITIONAL_SUPPORT_WEAK_SOCKET_CHARACTERS_TYPED_BY_ORIENTED_C_H_STABILIZER"
	SupportCharacterPressureClosed = "CONDITIONAL_SUPPORT_SOCKET_CHARACTER_PRESSURE_CLOSED_ONLY_AT_SEAL_LEVEL"

	FailureCRToCHNotNative       = "FAILED_ROUTE_C_R_TO_C_H_CHARACTER_IDENTIFICATION_NOT_NATIVE"
	FailureSocketMatchSeal       = "FAILED_ROUTE_SOCKET_CHARACTER_MATCHING_REMAINS_SEAL_NOT_NATIVE_THEOREM"
	FailureNoFullUnbrokenAF      = "FAILED_ROUTE_NO_FULL_UNBROKEN_A_F_THEOREM"
	FailureAForientNotFullAF     = "FAILED_ROUTE_A_F_ORIENT_NOT_FULL_A_F"
	FailureNoNativeFiniteTriple  = "FAILED_ROUTE_NO_NATIVE_FINITE_TRIPLE_THEOREM"
	FailureNoFullOperatorFO      = "FAILED_ROUTE_NO_FULL_UNBROKEN_OPERATOR_FIRST_ORDER_THEOREM"
	FailureNoCompleteJProof      = "FAILED_ROUTE_NO_COMPLETE_J_OPPOSITE_OPERATOR_PROOF_BEYOND_STABILIZER_SEAL"
	FailureNoBimoduleProof       = "FAILED_ROUTE_NO_FULL_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailureNoYukawaMagnitudes    = "FAILED_ROUTE_SYMBOLIC_Y_COEFFICIENTS_NOT_YUKAWA_MAGNITUDES"
	FailureNoNumericalYukawa     = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureNoAlphaSource         = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoTraceReadout        = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoOfficialNEffUpdate  = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoParticleAssignment  = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoNeutrinoTheorem     = "FAILED_ROUTE_NO_PHYSICAL_NEUTRINO_THEOREM"
	FailureNoThreeGenTheorem     = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
	FailureNoR3                  = "FAILED_ROUTE_R2_CHARACTER_INTERTWINER_SEAL_NOT_R3"
	FailureNoR4                  = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailurePunctureNotParticle   = "FAILED_ROUTE_Y_PLUS1_ZERO_NOT_PHYSICAL_NEUTRINO_THEOREM"
)

type Ledger struct {
	AlphaB                        float64
	OfficialNEff, OfficialCYukawa float64
	OfficialCHiggs                float64
	OfficialFrozen                bool
	AlphaNative, R3, R4           bool
}

type SocketCharacter struct {
	Name, Side, Algebra, Socket, CharacterExpression string
	Rank                                             int
	Typed, Native, OrientationRelative               bool
	Supports, Failures                               []string
}

type CharacterIdentification struct {
	Name, MapExpression, PlusMatch, MinusMatch string
	Defined, OrientationSeal, Native           bool
	ForcesPunctureEdge, OperatorCertified      bool
	Supports, Failures                         []string
}

type EdgeIntertwiner struct {
	Name, EdgeExpression, IntertwinerEquation       string
	DomainCharacter, CodomainCharacter              string
	ColorCentral, LeptonTrivial, PunctureEdge       bool
	HoldsGivenIdentification, OperatorCertified     bool
	PunctureForced, NumericalValue, YukawaMagnitude bool
	Supports, Failures                              []string
}

type FirstOrderPosition struct {
	Algebra, Target                                              string
	Gate861Inherited, ColorCentralityInstalled                   bool
	CharacterIdentificationNeeded, CharacterIdentificationSealed bool
	StabilizerOperatorCompatibilitySharpened                     bool
	StabilizerOperatorCompatibilityNative                        bool
	FullUnbrokenCompatibilityCertified                           bool
	Supports, Failures                                           []string
}

type PunctureKernel struct {
	PunctureEdge, RightPuncture, LeftKernel      string
	PunctureZero, PunctureNotForcedByCharacterID bool
	KernelRank                                   int
	KernelPhysicalTheorem                        bool
	Supports, Failures                           []string
}

type Impact struct {
	Classification                                                                   string
	CharacterMapDefined, EdgeIntertwinersHoldGivenSeal, PunctureZero                 bool
	FirstOrderSharpened, NativeCharacterTheorem, NativeFiniteTripleProof             bool
	AlphaStillSealed, MagnitudesStillMissing                                         bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4 bool
}

type Firewalls struct {
	Enforced                                                                                                bool
	CRToCHNotNative, SocketMatchSeal, NoFullUnbrokenAF, AForientNotFullAF, NoNativeFiniteTriple             bool
	NoFullOperatorFirstOrder, NoCompleteJProof, NoBimoduleProof                                             bool
	NoYukawaMagnitudes, NoNumericalYukawa, NoAlphaSource, NoTraceReadout                                    bool
	NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate, NoParticleAssignment, NoNeutrinoTheorem, NoThreeGenTheorem bool
	NoR3, NoR4                                                                                              bool
	Verdict                                                                                                 string
}

type Audit struct {
	ID         string
	Ledger     Ledger
	Characters []SocketCharacter
	IDMap      CharacterIdentification
	Edges      []EdgeIntertwiner
	FirstOrder FirstOrderPosition
	Kernel     PunctureKernel
	Impact     Impact
	Firewalls  Firewalls
	Truth      string
	Final      string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		ID:     AuditID,
		Ledger: Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true},
		Characters: []SocketCharacter{
			{Name: "chi_R^+", Side: "right", Algebra: "C_R", Socket: "e_+", CharacterExpression: "lambda", Rank: 1, Typed: true, Native: false, OrientationRelative: false, Supports: []string{StatusRightCharacterLedger, SupportRightCharsTyped}, Failures: []string{FailureCRToCHNotNative}},
			{Name: "chi_R^-", Side: "right", Algebra: "C_R", Socket: "e_-", CharacterExpression: "bar(lambda)", Rank: 1, Typed: true, Native: false, OrientationRelative: false, Supports: []string{StatusRightCharacterLedger, SupportRightCharsTyped}, Failures: []string{FailureCRToCHNotNative}},
			{Name: "chi_H^+", Side: "weak-oriented", Algebra: "C_H", Socket: "h_+", CharacterExpression: "z", Rank: 1, Typed: true, Native: false, OrientationRelative: true, Supports: []string{StatusWeakCharacterLedger, SupportWeakCharsOriented}, Failures: []string{FailureSocketMatchSeal}},
			{Name: "chi_H^-", Side: "weak-oriented", Algebra: "C_H", Socket: "h_-", CharacterExpression: "bar(z)", Rank: 1, Typed: true, Native: false, OrientationRelative: true, Supports: []string{StatusWeakCharacterLedger, SupportWeakCharsOriented}, Failures: []string{FailureSocketMatchSeal}},
		},
		IDMap: CharacterIdentification{Name: "iota_RH", MapExpression: "iota:C_R -> C_H, iota(lambda)=z", PlusMatch: "chi_R^+ <-> chi_H^+", MinusMatch: "chi_R^- <-> chi_H^-", Defined: true, OrientationSeal: true, Native: false, ForcesPunctureEdge: false, OperatorCertified: false, Supports: []string{StatusCharacterMapFormulated, SupportSocketMatchOrientation, SupportEdgeIntertwinersGivenID, SupportPunctureNotForced, SupportCharacterPressureClosed}, Failures: []string{FailureCRToCHNotNative, FailureSocketMatchSeal}},
		Edges: []EdgeIntertwiner{
			{Name: "Y_+3", EdgeExpression: "y_+3 |h_+><e_+| tensor I_{P_3}", IntertwinerEquation: "Y_+3 rho_R(lambda)=rho_H(iota(lambda)) Y_+3", DomainCharacter: "chi_R^+", CodomainCharacter: "chi_H^+", ColorCentral: true, HoldsGivenIdentification: true, Supports: []string{StatusIntertwinerConditions, StatusActiveIntertwiners, SupportEdgeIntertwinersGivenID}, Failures: []string{FailureSocketMatchSeal, FailureNoYukawaMagnitudes}},
			{Name: "Y_-3", EdgeExpression: "y_-3 |h_-><e_-| tensor I_{P_3}", IntertwinerEquation: "Y_-3 rho_R(lambda)=rho_H(iota(lambda)) Y_-3", DomainCharacter: "chi_R^-", CodomainCharacter: "chi_H^-", ColorCentral: true, HoldsGivenIdentification: true, Supports: []string{StatusIntertwinerConditions, StatusActiveIntertwiners, SupportEdgeIntertwinersGivenID}, Failures: []string{FailureSocketMatchSeal, FailureNoYukawaMagnitudes}},
			{Name: "Y_-1", EdgeExpression: "y_-1 |h_-><e_-| tensor I_{P_1}", IntertwinerEquation: "Y_-1 rho_R(lambda)=rho_H(iota(lambda)) Y_-1", DomainCharacter: "chi_R^-", CodomainCharacter: "chi_H^-", LeptonTrivial: true, HoldsGivenIdentification: true, Supports: []string{StatusIntertwinerConditions, StatusActiveIntertwiners, SupportEdgeIntertwinersGivenID}, Failures: []string{FailureSocketMatchSeal, FailureNoYukawaMagnitudes}},
			{Name: "Y_+1", EdgeExpression: "0", IntertwinerEquation: "Y_+1=0 remains absent", DomainCharacter: "chi_R^+", CodomainCharacter: "chi_H^+", PunctureEdge: true, HoldsGivenIdentification: true, PunctureForced: false, Supports: []string{StatusPunctureZeroReaudited, SupportPunctureNotForced}, Failures: []string{FailurePunctureNotParticle, FailureNoNeutrinoTheorem}},
		},
		FirstOrder: FirstOrderPosition{Algebra: "A_F^orient=C_R plus C_H plus M_3(C)", Target: "[[D_F^sym,rho_F(a)],rho_F^op(b)]=0", Gate861Inherited: true, ColorCentralityInstalled: true, CharacterIdentificationNeeded: true, CharacterIdentificationSealed: true, StabilizerOperatorCompatibilitySharpened: true, Supports: []string{StatusGate861Inherited, StatusFirstOrderSharpened, SupportFirstOrderSharpened, SupportEdgeIntertwinersGivenID}, Failures: []string{FailureSocketMatchSeal, FailureNoFullUnbrokenAF, FailureNoNativeFiniteTriple, FailureNoFullOperatorFO}},
		Kernel:     PunctureKernel{PunctureEdge: "Y_+1=0", RightPuncture: "e_+ tensor P_1", LeftKernel: "h_+ tensor P_1", PunctureZero: true, PunctureNotForcedByCharacterID: true, KernelRank: KernelRank, Supports: []string{StatusPunctureZeroReaudited, SupportPunctureNotForced}, Failures: []string{FailureNoNeutrinoTheorem, FailurePunctureNotParticle}},
		Impact:     Impact{Classification: "R2+++++_socket_character_intertwiner_seal", CharacterMapDefined: true, EdgeIntertwinersHoldGivenSeal: true, PunctureZero: true, FirstOrderSharpened: true, NativeCharacterTheorem: false, NativeFiniteTripleProof: false, AlphaStillSealed: true, MagnitudesStillMissing: true},
		Firewalls:  Firewalls{Enforced: true, CRToCHNotNative: true, SocketMatchSeal: true, NoFullUnbrokenAF: true, AForientNotFullAF: true, NoNativeFiniteTriple: true, NoFullOperatorFirstOrder: true, NoCompleteJProof: true, NoBimoduleProof: true, NoYukawaMagnitudes: true, NoNumericalYukawa: true, NoAlphaSource: true, NoTraceReadout: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NoParticleAssignment: true, NoNeutrinoTheorem: true, NoThreeGenTheorem: true, NoR3: true, NoR4: true, Verdict: StatusFirewallVerdict},
		Truth:      "Gate 862 audits the remaining socket-character pressure left by Gate 861. The right characters lambda/bar(lambda) and the oriented weak characters z/bar(z) can be paired by an orientation identification seal iota:C_R -> C_H, making the active scalar edge sockets intertwiners only at seal level.",
		Final:      "The stabilizer-branch first-order story is sharpened: color centrality is already installed and the active edges intertwine if the C_R/C_H character identification is admitted. The identification remains an orientation seal, not a native theorem; no full unbroken A_F theorem, native finite triple, Yukawa magnitude, alpha source, R3/R4 promotion, particle theorem, or official ledger update is certified.",
	}
	return a, a.Validate()
}

func (a Audit) Validate() error {
	err := func(msg string) error { return fmt.Errorf("%s: %s", AuditID, msg) }
	if !a.Ledger.OfficialFrozen || a.Ledger.AlphaNative || a.Ledger.R3 || a.Ledger.R4 {
		return err("ledger overpromoted")
	}
	if len(a.Characters) != 4 {
		return err("expected two right and two oriented weak characters")
	}
	for _, c := range a.Characters {
		if !c.Typed || c.Rank != 1 || c.Native {
			return err("character overpromoted or malformed")
		}
	}
	if !a.IDMap.Defined || !a.IDMap.OrientationSeal || a.IDMap.Native || a.IDMap.ForcesPunctureEdge || a.IDMap.OperatorCertified {
		return err("character identification map must be orientation seal only and must not force puncture")
	}
	if len(a.Edges) != 4 {
		return err("expected three active edge intertwiners plus one puncture edge")
	}
	active := 0
	for _, e := range a.Edges {
		if e.NumericalValue || e.YukawaMagnitude || e.OperatorCertified {
			return err("edge overpromoted")
		}
		if e.PunctureEdge {
			if e.EdgeExpression != "0" || e.PunctureForced {
				return err("puncture edge must remain zero and not forced")
			}
			continue
		}
		active++
		if !e.HoldsGivenIdentification {
			return err("active edge must hold given character identification")
		}
	}
	if active != 3 {
		return err("expected three active edge intertwiners")
	}
	if !a.FirstOrder.Gate861Inherited || !a.FirstOrder.ColorCentralityInstalled || !a.FirstOrder.CharacterIdentificationNeeded || !a.FirstOrder.CharacterIdentificationSealed || !a.FirstOrder.StabilizerOperatorCompatibilitySharpened || a.FirstOrder.StabilizerOperatorCompatibilityNative || a.FirstOrder.FullUnbrokenCompatibilityCertified {
		return err("first-order position not correctly sharpened")
	}
	if !a.Kernel.PunctureZero || !a.Kernel.PunctureNotForcedByCharacterID || a.Kernel.KernelRank != KernelRank || a.Kernel.KernelPhysicalTheorem {
		return err("puncture/kernel ledger overpromoted or malformed")
	}
	if !a.Impact.CharacterMapDefined || !a.Impact.EdgeIntertwinersHoldGivenSeal || !a.Impact.PunctureZero || !a.Impact.FirstOrderSharpened || a.Impact.NativeCharacterTheorem || a.Impact.NativeFiniteTripleProof || !a.Impact.AlphaStillSealed || !a.Impact.MagnitudesStillMissing || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		return err("impact overpromoted")
	}
	if !a.Firewalls.Enforced || !a.Firewalls.CRToCHNotNative || !a.Firewalls.SocketMatchSeal || !a.Firewalls.NoFullUnbrokenAF || !a.Firewalls.AForientNotFullAF || !a.Firewalls.NoNativeFiniteTriple || !a.Firewalls.NoFullOperatorFirstOrder || !a.Firewalls.NoCompleteJProof || !a.Firewalls.NoBimoduleProof || !a.Firewalls.NoYukawaMagnitudes || !a.Firewalls.NoNumericalYukawa || !a.Firewalls.NoAlphaSource || !a.Firewalls.NoTraceReadout || !a.Firewalls.NoOfficialNEffUpdate || !a.Firewalls.NoCYukawaCHiggsUpdate || !a.Firewalls.NoParticleAssignment || !a.Firewalls.NoNeutrinoTheorem || !a.Firewalls.NoThreeGenTheorem || !a.Firewalls.NoR3 || !a.Firewalls.NoR4 || a.Firewalls.Verdict != StatusFirewallVerdict {
		return err("firewall not enforced")
	}
	return nil
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("alpha_B=%.16g official_N_eff=%.16g official_C_Yukawa=%.16g official_C_Higgs=%.16g frozen=%t alpha_native=%t R3=%t R4=%t", l.AlphaB, l.OfficialNEff, l.OfficialCYukawa, l.OfficialCHiggs, l.OfficialFrozen, l.AlphaNative, l.R3, l.R4)
}

func FormatCharacters(chars []SocketCharacter) string {
	parts := make([]string, 0, len(chars))
	for _, c := range chars {
		parts = append(parts, fmt.Sprintf("%s[%s:%s:%s]=%s typed=%t native=%t orientation_relative=%t", c.Name, c.Algebra, c.Side, c.Socket, c.CharacterExpression, c.Typed, c.Native, c.OrientationRelative))
	}
	return strings.Join(parts, "; ")
}

func FormatIdentification(m CharacterIdentification) string {
	return fmt.Sprintf("%s %s plus=%s minus=%s defined=%t orientation_seal=%t native=%t forces_Y+1=%t operator_certified=%t supports=%s failures=%s", m.Name, m.MapExpression, m.PlusMatch, m.MinusMatch, m.Defined, m.OrientationSeal, m.Native, m.ForcesPunctureEdge, m.OperatorCertified, strings.Join(m.Supports, ","), strings.Join(m.Failures, ","))
}

func FormatEdges(edges []EdgeIntertwiner) string {
	parts := make([]string, 0, len(edges))
	for _, e := range edges {
		parts = append(parts, fmt.Sprintf("%s:%s equation=%s domain_char=%s codomain_char=%s holds_given_id=%t operator_certified=%t puncture_forced=%t", e.Name, e.EdgeExpression, e.IntertwinerEquation, e.DomainCharacter, e.CodomainCharacter, e.HoldsGivenIdentification, e.OperatorCertified, e.PunctureForced))
	}
	return strings.Join(parts, "; ")
}

func FormatFirstOrder(f FirstOrderPosition) string {
	return fmt.Sprintf("algebra=%s target=%s gate861=%t color_centrality=%t char_id_needed=%t char_id_sealed=%t sharpened=%t native=%t full_unbroken=%t supports=%s failures=%s", f.Algebra, f.Target, f.Gate861Inherited, f.ColorCentralityInstalled, f.CharacterIdentificationNeeded, f.CharacterIdentificationSealed, f.StabilizerOperatorCompatibilitySharpened, f.StabilizerOperatorCompatibilityNative, f.FullUnbrokenCompatibilityCertified, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatKernel(k PunctureKernel) string {
	return fmt.Sprintf("puncture_edge=%s right=%s left_kernel=%s puncture_zero=%t not_forced_by_char_id=%t kernel_rank=%d physical_theorem=%t", k.PunctureEdge, k.RightPuncture, k.LeftKernel, k.PunctureZero, k.PunctureNotForcedByCharacterID, k.KernelRank, k.KernelPhysicalTheorem)
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("class=%s char_map=%t edge_intertwiners_given_seal=%t puncture_zero=%t first_order_sharpened=%t native_char=%t native_triple=%t alpha_sealed=%t magnitudes_missing=%t update_Neff=%t update_CY=%t update_CH=%t R3=%t R4=%t", i.Classification, i.CharacterMapDefined, i.EdgeIntertwinersHoldGivenSeal, i.PunctureZero, i.FirstOrderSharpened, i.NativeCharacterTheorem, i.NativeFiniteTripleProof, i.AlphaStillSealed, i.MagnitudesStillMissing, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CanPromoteToR3, i.CanPromoteToR4)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%t verdict=%s no_native_char=%t socket_seal=%t no_full_AF=%t AForient_not_full=%t no_native_triple=%t no_first_order=%t no_yukawa=%t no_R3=%t no_R4=%t", f.Enforced, f.Verdict, f.CRToCHNotNative, f.SocketMatchSeal, f.NoFullUnbrokenAF, f.AForientNotFullAF, f.NoNativeFiniteTriple, f.NoFullOperatorFirstOrder, f.NoYukawaMagnitudes, f.NoR3, f.NoR4)
}

func containsAll(haystack, needles []string) bool {
	set := map[string]struct{}{}
	for _, h := range haystack {
		set[h] = struct{}{}
	}
	for _, n := range needles {
		if _, ok := set[n]; !ok {
			return false
		}
	}
	return true
}

func Statuses() []string {
	return []string{
		StatusGate861Inherited, StatusRightCharacterLedger, StatusWeakCharacterLedger, StatusCharacterMapFormulated, StatusIntertwinerConditions, StatusActiveIntertwiners, StatusPunctureZeroReaudited, StatusFirstOrderSharpened, StatusNoObservedDataUsed, StatusLedgerFrozen, StatusFirewallVerdict,
		SupportEdgeIntertwinersGivenID, SupportSocketMatchOrientation, SupportFirstOrderSharpened, SupportPunctureNotForced, SupportRightCharsTyped, SupportWeakCharsOriented, SupportCharacterPressureClosed,
		FailureCRToCHNotNative, FailureSocketMatchSeal, FailureNoFullUnbrokenAF, FailureAForientNotFullAF, FailureNoNativeFiniteTriple, FailureNoFullOperatorFO, FailureNoCompleteJProof, FailureNoBimoduleProof, FailureNoYukawaMagnitudes, FailureNoNumericalYukawa, FailureNoAlphaSource, FailureNoTraceReadout, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoParticleAssignment, FailureNoNeutrinoTheorem, FailureNoThreeGenTheorem, FailureNoR3, FailureNoR4, FailurePunctureNotParticle,
	}
}
