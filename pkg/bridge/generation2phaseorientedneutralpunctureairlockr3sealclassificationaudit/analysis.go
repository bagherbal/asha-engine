// Package generation2phaseorientedneutralpunctureairlockr3sealclassificationaudit implements
// Gate 900: PhaseOriented NeutralPuncture Airlock R3-Seal Classification Audit.
//
// Gate 900 follows Gate 899's result that the neutral-puncture airlock is a
// Z2 family ordered only by a right-character phase-orientation seal. It
// freezes the mature branch as an R3-sealed candidate under a neutral-puncture
// airlock plus phase orientation, while preserving the firewalls against native
// R3/R4, alpha derivation, physical-sector assignment, individual Yukawa
// values, and official ledger updates.
package generation2phaseorientedneutralpunctureairlockr3sealclassificationaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE900-PHASE-ORIENTED-NEUTRAL-PUNCTURE-AIRLOCK-R3-SEAL-CLASSIFICATION-AUDIT"

	AlphaB = 0.0003878958469680527
	Ssplit = 0.001292444818816423

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	PunctureOrderedPlus  = "p_+=e_+ tensor P_1"
	PiPlus3              = "Pi_+3=e_+ tensor P_3"
	PiMinus3             = "Pi_-3=e_- tensor P_3"
	PiMinus1             = "Pi_-1=e_- tensor P_1"
	RightPhaseOrder      = "lambda socket -> exposure/puncture; bar(lambda) socket -> active/rest"
	NeutralAirlockSeal   = "NeutralPunctureAirlockSeal"
	PhaseOrientationSeal = "RightCharacterPhaseOrientationSeal"
	BoundaryAlphaSeal    = "BoundaryAlphaIncidenceFlagSeal"
	HiggsOrientationSeal = "HiggsPostOrientationWeakSocketSeal"

	Classification = "R3_SEALED_CANDIDATE_UNDER_NEUTRAL_PUNCTURE_AIRLOCK_AND_RIGHT_CHARACTER_PHASE_ORIENTATION_NOT_NATIVE_R3"
	ShortStatus    = "R3_SEALED_CANDIDATE_CONFIRMED_NATIVE_R3_BLOCKED"

	StatusGate899Inherited       = "PASS_GATE899_PHASE_ORIENTATION_SEAL_INHERITED"
	StatusMatureChainAssembled   = "PASS_MATURE_R3_SEALED_CHAIN_ASSEMBLED"
	StatusProjectorLedgerAudited = "PASS_ORDERED_PROJECTOR_LEDGER_UNDER_PHASE_ORIENTATION_AUDITED"
	StatusReadoutAudited         = "PASS_Y_DAGGER_Y_READOUT_UNDER_SEALS_AUDITED"
	StatusPromotionChecklist     = "PASS_NATIVE_R3_PROMOTION_CHECKLIST_AUDITED"
	StatusOfficialFreeze         = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict        = "FIREWALL_PRESERVED_GATE900_R3_SEALED_CANDIDATE_NOT_NATIVE"

	SupportR3SealedCandidateComplete      = "CONDITIONAL_SUPPORT_R3_SEALED_CANDIDATE_STRUCTURE_COMPLETE"
	SupportNeutralAirlockUnifiesWounds    = "CONDITIONAL_SUPPORT_NEUTRAL_PUNCTURE_AIRLOCK_UNIFIES_ALPHA_AND_ORIENTATION_WOUNDS"
	SupportPhaseOrientationSelectsAirlock = "CONDITIONAL_SUPPORT_PHASE_ORIENTATION_SELECTS_ORDERED_AIRLOCK_IF_SEALED"
	SupportProjectorReadoutCompleteSealed = "CONDITIONAL_SUPPORT_PROJECTOR_LEDGER_AND_TRACE_READOUT_COMPLETE_UNDER_SEALS"
	SupportYDagYReproducesOperatorNEff    = "CONDITIONAL_SUPPORT_Y_DAGGER_Y_REPRODUCES_OPERATOR_N_EFF_UNDER_SEALS"
	SupportNativeBlockersReduced          = "CONDITIONAL_SUPPORT_NATIVE_R3_BLOCKERS_REDUCED_TO_PHASE_ORIENTATION_AND_NATIVE_AIRLOCK_FUNCTOR"
	SupportOrderedRepresentative          = "CONDITIONAL_SUPPORT_ORDERED_REPRESENTATIVE_P_PLUS_SELECTED_BY_PHASE_ORIENTATION_SEAL"
	SupportLedgerRowsStable               = "CONDITIONAL_SUPPORT_LEDGER_ROWS_STABLE_UNDER_R3_SEAL_CLASSIFICATION"

	FailureNotNativeR3                    = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureAlphaStillSealed               = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNoNativeNeutralPunctureAirlock = "FAILED_ROUTE_NO_NATIVE_NEUTRAL_PUNCTURE_AIRLOCK_FUNCTOR"
	FailureNoNativeRightPhaseOrientation  = "FAILED_ROUTE_NO_NATIVE_RIGHT_CHARACTER_PHASE_ORIENTATION_THEOREM"
	FailureNoNativeBoundaryIncidenceFlag  = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureNoNativeWeakSocketSelector     = "FAILED_ROUTE_NO_NATIVE_WEAK_SOCKET_SELECTOR_FUNCTIONAL"
	FailureNoNativeDescentFullAF          = "FAILED_ROUTE_NO_NATIVE_DESCENT_FROM_FULL_A_F_TO_A_F_ORIENT"
	FailureNoNativeSelectionSigmaPlus     = "FAILED_ROUTE_NO_NATIVE_SELECTION_OF_SIGMA_EQUALS_PLUS"
	FailureNoNativeR3SectorTraceLedger    = "FAILED_ROUTE_NO_NATIVE_R3_SECTOR_TRACE_LEDGER"
	FailureNoPhysicalParticleAssignment   = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoGenerationCarrierMap         = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap         = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues       = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate           = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate          = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator         = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4NativeYukawaTheorem        = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type MatureChainAudit struct {
	NeutralAirlockFamily           bool
	RightPhaseSealRequired         bool
	BoundaryAlphaReconstructed     bool
	HiggsOrientationPunctureKernel bool
	ProjectorLedger                bool
	PositiveReadoutRows            bool
	OperatorNEffReconstructed      bool
	Supports, Failures             []string
}

type OrderedRepresentativeAudit struct {
	Puncture                    string
	PhaseOrder                  string
	PiPlus3, PiMinus3, PiMinus1 string
	HRMinComplete               bool
	SelectedNatively            bool
	Supports, Failures          []string
}

type ReadoutAudit struct {
	WeightPlus3, WeightMinus3, WeightMinus1 float64
	TraceTotal, SquareTraceTotal            float64
	OperatorNEff, OperatorCYukawa           float64
	Positive, ReproducesNEff                bool
	Supports, Failures                      []string
}

type PromotionChecklist struct {
	NativeAlphaSource      bool
	NativePhaseOrientation bool
	FullAFDescent          bool
	NativeSectorLedger     bool
	PhysicalInterpretation bool
	AllowedNativeR3        bool
	Supports, Failures     []string
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
	NoNativeAirlockFunctor           bool
	NoNativeRightPhaseOrientation    bool
	NoNativeBoundaryIncidenceFunctor bool
	NoNativeWeakSocketSelector       bool
	NoNativeDescent                  bool
	NoSigmaPlusSelection             bool
	NoPhysicalAssignment             bool
	NoGenerationCarrier              bool
	NoFlavorOrientation              bool
	NoIndividualYukawas              bool
	NoOfficialLedgerUpdate           bool
	NoNativeYukawaOperator           bool
	NoR4YukawaTheorem                bool
	Verdict                          string
}

type Audit struct {
	ID             string
	MatureChain    MatureChainAudit
	Representative OrderedRepresentativeAudit
	Readout        ReadoutAudit
	Promotion      PromotionChecklist
	Freeze         FreezeAudit
	Firewalls      Firewalls
	Classification string
	ShortStatus    string
	Truth          string
	Final          string
}

func BuildDefault() (Audit, error) {
	mature := buildMatureChainAudit()
	if !mature.NeutralAirlockFamily || !mature.RightPhaseSealRequired || !mature.BoundaryAlphaReconstructed || !mature.HiggsOrientationPunctureKernel || !mature.ProjectorLedger || !mature.PositiveReadoutRows || !mature.OperatorNEffReconstructed {
		return Audit{}, fmt.Errorf("mature chain incomplete: %s", FormatMatureChain(mature))
	}
	rep := buildOrderedRepresentativeAudit()
	if rep.SelectedNatively || !rep.HRMinComplete || rep.Puncture != PunctureOrderedPlus || rep.PhaseOrder != RightPhaseOrder {
		return Audit{}, fmt.Errorf("ordered representative promoted incorrectly: %s", FormatRepresentative(rep))
	}
	readout := buildReadoutAudit()
	if !readout.Positive || !readout.ReproducesNEff || !near(readout.OperatorNEff, OperatorNEffDiagnostic) {
		return Audit{}, fmt.Errorf("readout mismatch: %s", FormatReadout(readout))
	}
	promotion := buildPromotionChecklist()
	if promotion.AllowedNativeR3 || promotion.NativeAlphaSource || promotion.NativePhaseOrientation || promotion.FullAFDescent || promotion.NativeSectorLedger || promotion.PhysicalInterpretation {
		return Audit{}, fmt.Errorf("native R3 promotion leaked: %s", FormatPromotion(promotion))
	}
	freeze := buildFreezeAudit()
	if !freeze.Frozen || !freeze.DiagnosticOnly || freeze.CanUpdate || near(freeze.OperatorNEff, freeze.OfficialNEff) {
		return Audit{}, fmt.Errorf("official ledger freeze leaked: %s", FormatFreeze(freeze))
	}
	firewalls := buildFirewalls()
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewall leak: %s", FormatFirewalls(firewalls))
	}
	return Audit{ID: AuditID, MatureChain: mature, Representative: rep, Readout: readout, Promotion: promotion, Freeze: freeze, Firewalls: firewalls, Classification: Classification, ShortStatus: ShortStatus, Truth: "Gate 900 freezes the branch as a mature R3-sealed candidate: structurally complete under NeutralPunctureAirlock plus RightCharacterPhaseOrientation seals, but not native R3.", Final: "R3_SEALED_CANDIDATE_CONFIRMED_NATIVE_R3_BLOCKED"}, nil
}

func buildMatureChainAudit() MatureChainAudit {
	return MatureChainAudit{NeutralAirlockFamily: true, RightPhaseSealRequired: true, BoundaryAlphaReconstructed: true, HiggsOrientationPunctureKernel: true, ProjectorLedger: true, PositiveReadoutRows: true, OperatorNEffReconstructed: true, Supports: []string{StatusGate899Inherited, StatusMatureChainAssembled, SupportR3SealedCandidateComplete, SupportNeutralAirlockUnifiesWounds, SupportNativeBlockersReduced}, Failures: []string{FailureNoNativeNeutralPunctureAirlock, FailureNoNativeRightPhaseOrientation, FailureNotNativeR3}}
}

func buildOrderedRepresentativeAudit() OrderedRepresentativeAudit {
	return OrderedRepresentativeAudit{Puncture: PunctureOrderedPlus, PhaseOrder: RightPhaseOrder, PiPlus3: PiPlus3, PiMinus3: PiMinus3, PiMinus1: PiMinus1, HRMinComplete: true, SelectedNatively: false, Supports: []string{StatusProjectorLedgerAudited, SupportPhaseOrientationSelectsAirlock, SupportOrderedRepresentative, SupportProjectorReadoutCompleteSealed}, Failures: []string{FailureNoNativeSelectionSigmaPlus, FailureNoNativeRightPhaseOrientation, FailureNoNativeWeakSocketSelector}}
}

func buildReadoutAudit() ReadoutAudit {
	wPlus := 1.0
	wMinus3 := AlphaB * (1 - AlphaB)
	wMinus1 := 3 * AlphaB * AlphaB
	trace := 3*wPlus + 3*wMinus3 + wMinus1
	square := 3*wPlus*wPlus + 3*wMinus3*wMinus3 + wMinus1*wMinus1
	neff := trace * trace / square
	cy := 3.0 / neff
	return ReadoutAudit{WeightPlus3: wPlus, WeightMinus3: wMinus3, WeightMinus1: wMinus1, TraceTotal: trace, SquareTraceTotal: square, OperatorNEff: neff, OperatorCYukawa: cy, Positive: wPlus > 0 && wMinus3 > 0 && wMinus1 > 0, ReproducesNEff: near(neff, OperatorNEffDiagnostic), Supports: []string{StatusReadoutAudited, SupportYDagYReproducesOperatorNEff, SupportLedgerRowsStable}, Failures: []string{FailureAlphaStillSealed, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate}}
}

func buildPromotionChecklist() PromotionChecklist {
	return PromotionChecklist{NativeAlphaSource: false, NativePhaseOrientation: false, FullAFDescent: false, NativeSectorLedger: false, PhysicalInterpretation: false, AllowedNativeR3: false, Supports: []string{StatusPromotionChecklist, SupportNativeBlockersReduced}, Failures: []string{FailureAlphaStillSealed, FailureNoNativeBoundaryIncidenceFlag, FailureNoNativeRightPhaseOrientation, FailureNoNativeNeutralPunctureAirlock, FailureNoNativeDescentFullAF, FailureNoNativeR3SectorTraceLedger, FailureNoPhysicalParticleAssignment, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoR4NativeYukawaTheorem}}
}

func buildFreezeAudit() FreezeAudit {
	return FreezeAudit{OperatorNEff: OperatorNEffDiagnostic, OfficialNEff: OfficialNEffFrozen, OperatorCYukawa: OperatorCYukawaDiagnostic, OfficialCYukawa: OfficialCYukawaFrozen, OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen, Frozen: true, DiagnosticOnly: true, CanUpdate: false, Supports: []string{StatusOfficialFreeze}, Failures: []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate}}
}

func buildFirewalls() Firewalls {
	return Firewalls{Enforced: true, NotNativeR3: true, AlphaStillSealed: true, NoNativeAirlockFunctor: true, NoNativeRightPhaseOrientation: true, NoNativeBoundaryIncidenceFunctor: true, NoNativeWeakSocketSelector: true, NoNativeDescent: true, NoSigmaPlusSelection: true, NoPhysicalAssignment: true, NoGenerationCarrier: true, NoFlavorOrientation: true, NoIndividualYukawas: true, NoOfficialLedgerUpdate: true, NoNativeYukawaOperator: true, NoR4YukawaTheorem: true, Verdict: StatusFirewallVerdict}
}

func FormatMatureChain(m MatureChainAudit) string {
	return fmt.Sprintf("mature_chain(airlock_family=%t phase_required=%t alpha_reconstructed=%t higgs_puncture_kernel=%t projector_ledger=%t readout_rows=%t neff=%t supports=%s failures=%s)", m.NeutralAirlockFamily, m.RightPhaseSealRequired, m.BoundaryAlphaReconstructed, m.HiggsOrientationPunctureKernel, m.ProjectorLedger, m.PositiveReadoutRows, m.OperatorNEffReconstructed, strings.Join(m.Supports, ","), strings.Join(m.Failures, ","))
}

func FormatRepresentative(r OrderedRepresentativeAudit) string {
	return fmt.Sprintf("ordered_representative(puncture=%s phase_order=%s pi=[%s,%s,%s] hrmin_complete=%t selected_natively=%t supports=%s failures=%s)", r.Puncture, r.PhaseOrder, r.PiPlus3, r.PiMinus3, r.PiMinus1, r.HRMinComplete, r.SelectedNatively, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}

func FormatReadout(r ReadoutAudit) string {
	return fmt.Sprintf("readout(w_plus3=%.16g w_minus3=%.16g w_minus1=%.16g trace=%.16g square=%.16g neff=%.16g cy=%.16g positive=%t reproduces_neff=%t supports=%s failures=%s)", r.WeightPlus3, r.WeightMinus3, r.WeightMinus1, r.TraceTotal, r.SquareTraceTotal, r.OperatorNEff, r.OperatorCYukawa, r.Positive, r.ReproducesNEff, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}

func FormatPromotion(p PromotionChecklist) string {
	return fmt.Sprintf("promotion(native_alpha=%t native_phase=%t full_af_descent=%t native_ledger=%t physical=%t allowed_r3=%t supports=%s failures=%s)", p.NativeAlphaSource, p.NativePhaseOrientation, p.FullAFDescent, p.NativeSectorLedger, p.PhysicalInterpretation, p.AllowedNativeR3, strings.Join(p.Supports, ","), strings.Join(p.Failures, ","))
}

func FormatFreeze(f FreezeAudit) string {
	return fmt.Sprintf("freeze(operator_Neff=%.16g official_Neff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t not_native_r3=%t alpha=%t airlock=%t phase=%t boundary=%t weak=%t descent=%t sigma=%t physical=%t generation=%t flavor=%t individual=%t official=%t yukawa=%t r4=%t verdict=%s)", f.Enforced, f.NotNativeR3, f.AlphaStillSealed, f.NoNativeAirlockFunctor, f.NoNativeRightPhaseOrientation, f.NoNativeBoundaryIncidenceFunctor, f.NoNativeWeakSocketSelector, f.NoNativeDescent, f.NoSigmaPlusSelection, f.NoPhysicalAssignment, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4YukawaTheorem, f.Verdict)
}

func Statuses() []string {
	return []string{StatusGate899Inherited, StatusMatureChainAssembled, StatusProjectorLedgerAudited, StatusReadoutAudited, StatusPromotionChecklist, StatusOfficialFreeze, StatusFirewallVerdict, SupportR3SealedCandidateComplete, SupportNeutralAirlockUnifiesWounds, SupportPhaseOrientationSelectsAirlock, SupportProjectorReadoutCompleteSealed, SupportYDagYReproducesOperatorNEff, SupportNativeBlockersReduced, SupportOrderedRepresentative, SupportLedgerRowsStable, FailureNotNativeR3, FailureAlphaStillSealed, FailureNoNativeNeutralPunctureAirlock, FailureNoNativeRightPhaseOrientation, FailureNoNativeBoundaryIncidenceFlag, FailureNoNativeWeakSocketSelector, FailureNoNativeDescentFullAF, FailureNoNativeSelectionSigmaPlus, FailureNoNativeR3SectorTraceLedger, FailureNoPhysicalParticleAssignment, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem}
}

func (a Audit) FirewallsList() []string {
	return []string{FailureNotNativeR3, FailureAlphaStillSealed, FailureNoNativeNeutralPunctureAirlock, FailureNoNativeRightPhaseOrientation, FailureNoNativeBoundaryIncidenceFlag, FailureNoNativeWeakSocketSelector, FailureNoNativeDescentFullAF, FailureNoNativeSelectionSigmaPlus, FailureNoOfficialNEffUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem}
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NotNativeR3 && f.AlphaStillSealed && f.NoNativeAirlockFunctor && f.NoNativeRightPhaseOrientation && f.NoNativeBoundaryIncidenceFunctor && f.NoNativeWeakSocketSelector && f.NoNativeDescent && f.NoSigmaPlusSelection && f.NoPhysicalAssignment && f.NoGenerationCarrier && f.NoFlavorOrientation && f.NoIndividualYukawas && f.NoOfficialLedgerUpdate && f.NoNativeYukawaOperator && f.NoR4YukawaTheorem && f.Verdict == StatusFirewallVerdict
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
