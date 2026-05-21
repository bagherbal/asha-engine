// Package generation2neutralpunctureairlockunificationaudit implements
// Gate 895: NeutralPuncture Airlock Unification Audit.
//
// Gate 895 follows Gate 894's minimal null-edge orientation candidate. It
// audits whether the two remaining R3 seals — the BoundaryAlpha incidence-flag
// seal and the Higgs/post-orientation weak-frame seal — are two projections of
// one deeper neutral puncture airlock centered on p=e_+ tensor P_1. The gate
// source-types the common puncture object but does not certify a native
// airlock functor, does not promote to R3, and does not update official ledgers.
package generation2neutralpunctureairlockunificationaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE895-NEUTRAL-PUNCTURE-AIRLOCK-UNIFICATION-AUDIT"

	AlphaB = 0.0003878958469680527
	Ssplit = 0.001292444818816423

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	RightCharacterSplit = "C_R^2=e_+ plus e_-"
	LeptoColorSplit     = "W=P_1 plus P_3"
	WeakFrame           = "C_L^2=h_+ plus h_-"

	Puncture    = "p=e_+ tensor P_1"
	F0          = "F_0=p=e_+ tensor P_1"
	F1          = "F_1=e_+ tensor W"
	F2          = "F_2=C_R^2 tensor W"
	PiTop       = "F_1/F_0=Pi_top=e_+ tensor P_3"
	HRMin       = "F_2/F_0=H_R^min"
	LeftKernel  = "h_+ tensor P_1"
	ImageY      = "Im(Y)=(h_+ tensor P_3) plus (h_- tensor P_3) plus (h_- tensor P_1)"
	MissingEdge = "Y_+1:e_+ tensor P_1 -> h_+ tensor P_1"

	RankPuncture = 1
	RankF1       = 4
	RankF2       = 8
	RankPiTop    = 3
	RankHRMin    = 7
	RankHLeft    = 8
	RankImageY   = 7
	RankKernel   = 1
	DimH10       = 10
	DimH72       = 72

	Classification = "R3_DUALSEAL_NEUTRAL_PUNCTURE_AIRLOCK_UNIFICATION_CANDIDATE_NOT_NATIVE"
	ShortStatus    = "R3_CANDIDATE_TWO_SEALS_COLLAPSE_TO_NEUTRAL_PUNCTURE_AIRLOCK_SEAL"
	NextFrontier   = "NEUTRAL_PUNCTURE_AIRLOCK_VARIATIONAL_FUNCTIONAL_AUDIT"

	StatusGate894Inherited        = "PASS_GATE894_MINIMAL_NULL_EDGE_ORIENTATION_CANDIDATE_INHERITED"
	StatusPunctureIndependence    = "PASS_NEUTRAL_PUNCTURE_INDEPENDENCE_AUDITED"
	StatusAlphaFlagReconstructed  = "PASS_PUNCTURE_FLAG_ALPHA_TARGETS_RECONSTRUCTED"
	StatusWeakKernelReconstructed = "PASS_PUNCTURE_KERNEL_WEAK_SOCKET_RECONSTRUCTED"
	StatusCrossSealUnification    = "PASS_ALPHA_AND_ORIENTATION_SEALS_UNIFICATION_AUDITED"
	StatusOfficialFreeze          = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict         = "FIREWALL_PRESERVED_GATE895_AIRLOCK_NOT_NATIVE"

	SupportPunctureDefinedBeforeWeakOrientation = "CONDITIONAL_SUPPORT_NEUTRAL_PUNCTURE_DEFINED_BEFORE_WEAK_SOCKET_ORIENTATION"
	SupportPunctureFlagReconstructsAlphaTargets = "CONDITIONAL_SUPPORT_PUNCTURE_FLAG_RECONSTRUCTS_BOUNDARY_ALPHA_TARGETS"
	SupportAlphaReconstructedFromPunctureFlag   = "CONDITIONAL_SUPPORT_ALPHA_B_RECONSTRUCTED_FROM_NEUTRAL_PUNCTURE_FLAG"
	SupportWeakSocketFromLeftKernelCandidate    = "CONDITIONAL_SUPPORT_WEAK_SOCKET_FRAME_CAN_BE_RECONSTRUCTED_FROM_LEFT_KERNEL_CANDIDATE"
	SupportHPlusMissingLeftLeptonLine           = "CONDITIONAL_SUPPORT_H_PLUS_IS_THE_MISSING_LEFT_LEPTON_LINE_OF_MINIMAL_IMAGE"
	SupportCommonSourceOfAlphaAndOrientation    = "CONDITIONAL_SUPPORT_NEUTRAL_PUNCTURE_IS_COMMON_SOURCE_OF_ALPHA_FLAG_AND_WEAK_ORIENTATION"
	SupportDualSealReducesToAirlock             = "CONDITIONAL_SUPPORT_R3_DUAL_SEAL_WOUND_REDUCES_TO_SINGLE_PUNCTURE_AIRLOCK_FUNCTOR"
	SupportHiggsSealWeakenedToAirlock           = "CONDITIONAL_SUPPORT_HIGGS_ORIENTATION_SEAL_WEAKENED_TO_PUNCTURE_AIRLOCK_SEAL"
	SupportOperatorNEffReproduced               = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_REPRODUCED_BY_ORIENTED_LEDGER"

	FailureNotNativeR3                       = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureAlphaStillSealed                  = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureHiggsOrientationStillSealed       = "FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED"
	FailureNoNeutralPunctureAirlockFunctor   = "FAILED_ROUTE_NO_NATIVE_NEUTRAL_PUNCTURE_AIRLOCK_FUNCTOR"
	FailureNoNativeBoundaryIncidenceFunctor  = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureNoNativeWeakSocketSelector        = "FAILED_ROUTE_NO_NATIVE_WEAK_SOCKET_SELECTOR_FUNCTIONAL"
	FailureNoNativeVariationalMinimality     = "FAILED_ROUTE_NO_NATIVE_VARIATIONAL_MINIMALITY_THEOREM"
	FailurePunctureFlagNotNativeAlphaFunctor = "FAILED_ROUTE_PUNCTURE_FLAG_NOT_NATIVE_BOUNDARY_ALPHA_FUNCTOR_YET"
	FailureNoNativeMinimalImageSelection     = "FAILED_ROUTE_NO_NATIVE_VARIATIONAL_RULE_SELECTING_THIS_IMAGE"
	FailureWeakReconstructionDependsOnImage  = "FAILED_ROUTE_WEAK_SOCKET_RECONSTRUCTION_STILL_DEPENDS_ON_MINIMAL_IMAGE_CHOICE"
	FailureFullHMixesWeakSockets             = "FAILED_ROUTE_FULL_H_ACTION_MIXES_H_PLUS_H_MINUS"
	FailureNoNativeDescentFullToOrient       = "FAILED_ROUTE_NO_NATIVE_DESCENT_FROM_FULL_A_F_TO_A_F_ORIENT"
	FailureNoNativeR3SectorLedger            = "FAILED_ROUTE_NO_NATIVE_R3_SECTOR_TRACE_LEDGER"
	FailureNoGenerationCarrierMap            = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap            = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues          = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate              = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate             = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator            = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4NativeYukawaTheorem           = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type PunctureIndependenceAudit struct {
	Puncture                         string
	RequiresRightCharacterSplit      bool
	RequiresLeptoColorSplit          bool
	RequiresWeakSocketFrame          bool
	DefinedBeforeWeakOrientation     bool
	UpstreamOfAlphaFlagAndWeakKernel bool
	Supports, Failures               []string
}

type AlphaFlagAudit struct {
	F0, F1, F2             string
	F0SubsetF1SubsetF2     bool
	RankF0, RankF1, RankF2 int
	Q1, Q2                 string
	RankQ1, RankQ2         int
	DimH10, DimH72         int
	Alpha                  float64
	ReconstructsAlpha      bool
	NativeAlphaFunctor     bool
	Supports, Failures     []string
}

type WeakKernelAudit struct {
	HLeftRank, ImageRank, KernelRank int
	Image                            string
	Kernel                           string
	QuotientIsKernel                 bool
	CanReconstructWeakFrameCandidate bool
	NativeMinimalImageRule           bool
	DependsOnMinimalImageChoice      bool
	Supports, Failures               []string
}

type AirlockUnificationAudit struct {
	CommonObject               string
	ControlsAlphaFlag          bool
	ControlsWeakKernel         bool
	TwoSealProblemReducesToOne bool
	NativeAirlockFunctor       bool
	NextMissingObject          string
	Supports, Failures         []string
}

type FreezeAudit struct {
	OperatorNEff, OfficialNEff        float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	Frozen, DiagnosticOnly, CanUpdate bool
	Supports, Failures                []string
}

type Firewalls struct {
	Enforced                        bool
	NotNativeR3                     bool
	AlphaStillSealed                bool
	HiggsOrientationStillSealed     bool
	NoNeutralPunctureAirlockFunctor bool
	NoBoundaryIncidenceFunctor      bool
	NoWeakSocketSelector            bool
	NoVariationalMinimality         bool
	NoNativeDescentFullToOrient     bool
	FullHMixesWeakSockets           bool
	NoNativeR3SectorLedger          bool
	NoGenerationCarrier             bool
	NoFlavorOrientation             bool
	NoIndividualYukawas             bool
	NoOfficialLedgerUpdate          bool
	NoNativeYukawaOperator          bool
	NoR4NativeYukawaTheorem         bool
	Verdict                         string
}

type Audit struct {
	ID                   string
	PunctureIndependence PunctureIndependenceAudit
	AlphaFlag            AlphaFlagAudit
	WeakKernel           WeakKernelAudit
	Airlock              AirlockUnificationAudit
	Freeze               FreezeAudit
	Firewalls            Firewalls
	Truth                string
	Final                string
}

func BuildDefault() (Audit, error) {
	puncture := buildPunctureIndependenceAudit()
	if !puncture.DefinedBeforeWeakOrientation || puncture.RequiresWeakSocketFrame {
		return Audit{}, fmt.Errorf("puncture independence failed: %s", FormatPunctureIndependence(puncture))
	}

	alpha := buildAlphaFlagAudit()
	if !alpha.F0SubsetF1SubsetF2 || !alpha.ReconstructsAlpha || alpha.NativeAlphaFunctor || alpha.RankQ1 != RankPiTop || alpha.RankQ2 != RankHRMin {
		return Audit{}, fmt.Errorf("alpha flag promoted incorrectly: %s", FormatAlphaFlag(alpha))
	}

	weak := buildWeakKernelAudit()
	if !weak.QuotientIsKernel || !weak.CanReconstructWeakFrameCandidate || weak.NativeMinimalImageRule || !weak.DependsOnMinimalImageChoice {
		return Audit{}, fmt.Errorf("weak kernel audit promoted incorrectly: %s", FormatWeakKernel(weak))
	}

	airlock := buildAirlockUnificationAudit()
	if !airlock.ControlsAlphaFlag || !airlock.ControlsWeakKernel || !airlock.TwoSealProblemReducesToOne || airlock.NativeAirlockFunctor {
		return Audit{}, fmt.Errorf("airlock unification leak: %s", FormatAirlock(airlock))
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
		NoNeutralPunctureAirlockFunctor: true, NoBoundaryIncidenceFunctor: true, NoWeakSocketSelector: true,
		NoVariationalMinimality: true, NoNativeDescentFullToOrient: true, FullHMixesWeakSockets: true,
		NoNativeR3SectorLedger: true, NoGenerationCarrier: true, NoFlavorOrientation: true, NoIndividualYukawas: true,
		NoOfficialLedgerUpdate: true, NoNativeYukawaOperator: true, NoR4NativeYukawaTheorem: true,
		Verdict: StatusFirewallVerdict,
	}
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewalls not preserved: %s", FormatFirewalls(firewalls))
	}

	return Audit{
		ID: AuditID, PunctureIndependence: puncture, AlphaFlag: alpha, WeakKernel: weak, Airlock: airlock, Freeze: freeze, Firewalls: firewalls,
		Truth: "Gate 895 collapses the two remaining R3 seals to a common neutral puncture source candidate p=e_+ tensor P_1. The same p is the basepoint of the BoundaryAlpha incidence flag and the right puncture whose missing edge leaves h_+ tensor P_1 as the left kernel candidate.",
		Final: "The branch becomes R3_DUALSEAL_NEUTRAL_PUNCTURE_AIRLOCK_UNIFICATION_CANDIDATE_NOT_NATIVE. The dual-seal wound is reduced to a single NeutralPunctureAirlockFunctor candidate, but no native airlock functor, incidence functor, weak-socket selector functional, variational minimality theorem, native R3 ledger, or official ledger update is certified.",
	}, nil
}

func buildPunctureIndependenceAudit() PunctureIndependenceAudit {
	return PunctureIndependenceAudit{
		Puncture: Puncture, RequiresRightCharacterSplit: true, RequiresLeptoColorSplit: true, RequiresWeakSocketFrame: false,
		DefinedBeforeWeakOrientation: true, UpstreamOfAlphaFlagAndWeakKernel: true,
		Supports: []string{StatusPunctureIndependence, SupportPunctureDefinedBeforeWeakOrientation, SupportCommonSourceOfAlphaAndOrientation},
		Failures: []string{},
	}
}

func buildAlphaFlagAudit() AlphaFlagAudit {
	alpha := float64(RankPiTop)/float64(DimH10)*Ssplit + float64(RankHRMin)/float64(DimH72)*Ssplit*Ssplit
	return AlphaFlagAudit{
		F0: F0, F1: F1, F2: F2, F0SubsetF1SubsetF2: true,
		RankF0: RankPuncture, RankF1: RankF1, RankF2: RankF2,
		Q1: PiTop, Q2: HRMin, RankQ1: RankPiTop, RankQ2: RankHRMin, DimH10: DimH10, DimH72: DimH72,
		Alpha: alpha, ReconstructsAlpha: near(alpha, AlphaB), NativeAlphaFunctor: false,
		Supports: []string{StatusAlphaFlagReconstructed, SupportPunctureFlagReconstructsAlphaTargets, SupportAlphaReconstructedFromPunctureFlag},
		Failures: []string{FailurePunctureFlagNotNativeAlphaFunctor, FailureNoNativeBoundaryIncidenceFunctor, FailureAlphaStillSealed},
	}
}

func buildWeakKernelAudit() WeakKernelAudit {
	return WeakKernelAudit{
		HLeftRank: RankHLeft, ImageRank: RankImageY, KernelRank: RankKernel, Image: ImageY, Kernel: LeftKernel,
		QuotientIsKernel: true, CanReconstructWeakFrameCandidate: true, NativeMinimalImageRule: false, DependsOnMinimalImageChoice: true,
		Supports: []string{StatusWeakKernelReconstructed, SupportWeakSocketFromLeftKernelCandidate, SupportHPlusMissingLeftLeptonLine, SupportHiggsSealWeakenedToAirlock},
		Failures: []string{FailureNoNativeMinimalImageSelection, FailureWeakReconstructionDependsOnImage, FailureNoNativeWeakSocketSelector, FailureHiggsOrientationStillSealed},
	}
}

func buildAirlockUnificationAudit() AirlockUnificationAudit {
	return AirlockUnificationAudit{
		CommonObject: Puncture, ControlsAlphaFlag: true, ControlsWeakKernel: true, TwoSealProblemReducesToOne: true, NativeAirlockFunctor: false,
		NextMissingObject: "NeutralPunctureAirlockFunctor or NeutralPuncture Airlock Variational Functional",
		Supports:          []string{StatusCrossSealUnification, SupportCommonSourceOfAlphaAndOrientation, SupportDualSealReducesToAirlock, SupportHiggsSealWeakenedToAirlock},
		Failures:          []string{FailureNoNeutralPunctureAirlockFunctor, FailureNoNativeBoundaryIncidenceFunctor, FailureNoNativeWeakSocketSelector, FailureNoNativeVariationalMinimality},
	}
}

func FormatPunctureIndependence(p PunctureIndependenceAudit) string {
	return fmt.Sprintf("puncture_independence(p=%s right_split=%t leptocolor_split=%t weak_frame=%t before_weak_orientation=%t upstream=%t supports=%s failures=%s)", p.Puncture, p.RequiresRightCharacterSplit, p.RequiresLeptoColorSplit, p.RequiresWeakSocketFrame, p.DefinedBeforeWeakOrientation, p.UpstreamOfAlphaFlagAndWeakKernel, strings.Join(p.Supports, ","), strings.Join(p.Failures, ","))
}

func FormatAlphaFlag(a AlphaFlagAudit) string {
	return fmt.Sprintf("alpha_flag(F0=%s F1=%s F2=%s flag=%t ranks=%d,%d,%d Q1=%s rankQ1=%d Q2=%s rankQ2=%d dims=%d,%d alpha=%.16g reconstructs=%t native=%t supports=%s failures=%s)", a.F0, a.F1, a.F2, a.F0SubsetF1SubsetF2, a.RankF0, a.RankF1, a.RankF2, a.Q1, a.RankQ1, a.Q2, a.RankQ2, a.DimH10, a.DimH72, a.Alpha, a.ReconstructsAlpha, a.NativeAlphaFunctor, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatWeakKernel(w WeakKernelAudit) string {
	return fmt.Sprintf("weak_kernel(H_L_rank=%d image_rank=%d kernel=%s kernel_rank=%d image=%s quotient_is_kernel=%t reconstructs_frame_candidate=%t native_minimal_image=%t depends_on_image=%t supports=%s failures=%s)", w.HLeftRank, w.ImageRank, w.Kernel, w.KernelRank, w.Image, w.QuotientIsKernel, w.CanReconstructWeakFrameCandidate, w.NativeMinimalImageRule, w.DependsOnMinimalImageChoice, strings.Join(w.Supports, ","), strings.Join(w.Failures, ","))
}

func FormatAirlock(a AirlockUnificationAudit) string {
	return fmt.Sprintf("airlock(common=%s controls_alpha=%t controls_weak=%t reduces_two_seals=%t native=%t next=%s supports=%s failures=%s)", a.CommonObject, a.ControlsAlphaFlag, a.ControlsWeakKernel, a.TwoSealProblemReducesToOne, a.NativeAirlockFunctor, a.NextMissingObject, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatFreeze(f FreezeAudit) string {
	return fmt.Sprintf("official_freeze(operator_Neff=%.16g official_Neff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t not_native_r3=%t alpha_sealed=%t higgs_sealed=%t no_airlock=%t no_incidence=%t no_selector=%t no_variational=%t no_descent=%t full_H_mixes=%t no_r3=%t no_generation=%t no_flavor=%t no_individual=%t no_official=%t no_yukawa=%t no_r4=%t verdict=%s)", f.Enforced, f.NotNativeR3, f.AlphaStillSealed, f.HiggsOrientationStillSealed, f.NoNeutralPunctureAirlockFunctor, f.NoBoundaryIncidenceFunctor, f.NoWeakSocketSelector, f.NoVariationalMinimality, f.NoNativeDescentFullToOrient, f.FullHMixesWeakSockets, f.NoNativeR3SectorLedger, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4NativeYukawaTheorem, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate894Inherited, StatusPunctureIndependence, StatusAlphaFlagReconstructed, StatusWeakKernelReconstructed, StatusCrossSealUnification, StatusOfficialFreeze, StatusFirewallVerdict,
		SupportPunctureDefinedBeforeWeakOrientation, SupportPunctureFlagReconstructsAlphaTargets, SupportAlphaReconstructedFromPunctureFlag, SupportWeakSocketFromLeftKernelCandidate, SupportHPlusMissingLeftLeptonLine, SupportCommonSourceOfAlphaAndOrientation, SupportDualSealReducesToAirlock, SupportHiggsSealWeakenedToAirlock, SupportOperatorNEffReproduced,
		FailureNotNativeR3, FailureAlphaStillSealed, FailureHiggsOrientationStillSealed, FailureNoNeutralPunctureAirlockFunctor, FailureNoNativeBoundaryIncidenceFunctor, FailureNoNativeWeakSocketSelector, FailureNoNativeVariationalMinimality, FailurePunctureFlagNotNativeAlphaFunctor, FailureNoNativeMinimalImageSelection, FailureWeakReconstructionDependsOnImage, FailureFullHMixesWeakSockets, FailureNoNativeDescentFullToOrient, FailureNoNativeR3SectorLedger, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem,
	}
}

func (a Audit) FirewallsList() []string {
	return []string{FailureNotNativeR3, FailureAlphaStillSealed, FailureHiggsOrientationStillSealed, FailureNoNeutralPunctureAirlockFunctor, FailureNoNativeBoundaryIncidenceFunctor, FailureNoNativeWeakSocketSelector, FailureNoNativeVariationalMinimality, FailureFullHMixesWeakSockets, FailureNoNativeDescentFullToOrient, FailureNoNativeR3SectorLedger, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem}
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NotNativeR3 && f.AlphaStillSealed && f.HiggsOrientationStillSealed &&
		f.NoNeutralPunctureAirlockFunctor && f.NoBoundaryIncidenceFunctor && f.NoWeakSocketSelector &&
		f.NoVariationalMinimality && f.NoNativeDescentFullToOrient && f.FullHMixesWeakSockets &&
		f.NoNativeR3SectorLedger && f.NoGenerationCarrier && f.NoFlavorOrientation && f.NoIndividualYukawas &&
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
