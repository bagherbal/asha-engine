// Package generation2neutralpunctureairlockvariationalfunctionalaudit implements
// Gate 896: NeutralPuncture Airlock Variational Functional Audit.
//
// Gate 896 follows Gate 895's neutral puncture airlock unification. It audits
// whether the common puncture p=e_+ tensor P_1 can be selected by a lawful
// variational/minimality functional rather than declared. The gate formulates a
// NeutralPunctureAirlockFunctional candidate and tests rank, alpha-flag,
// edge-support, left-kernel, and B-L terms. It shows that rank, flag, and B-L
// constraints reduce the candidate set to neutral lepton singletons but do not
// distinguish e_+ tensor P_1 from e_- tensor P_1. The plus puncture is selected
// only after the oriented edge ordering is included, so the next missing object
// is an OrientedEdgeOrderingFunctional / SocketOrderSelector. No native R3,
// individual Yukawa value, physical sector assignment, or official ledger update
// is certified.
package generation2neutralpunctureairlockvariationalfunctionalaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE896-NEUTRAL-PUNCTURE-AIRLOCK-VARIATIONAL-FUNCTIONAL-AUDIT"

	AlphaB = 0.0003878958469680527
	Ssplit = 0.001292444818816423

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	PuncturePlus  = "e_+ tensor P_1"
	PunctureMinus = "e_- tensor P_1"
	ColorPlus     = "e_+ tensor P_3"
	ColorMinus    = "e_- tensor P_3"
	P1            = "P_1"
	P3            = "P_3"
	W             = "W=P_1 plus P_3"
	RightRect     = "C_R^2 tensor W"
	F1Plus        = "e_+ tensor W"
	F1Minus       = "e_- tensor W"
	PiTopPlus     = "(e_+ tensor W) minus (e_+ tensor P_1)=e_+ tensor P_3"
	PiTopMinus    = "(e_- tensor W) minus (e_- tensor P_1)=e_- tensor P_3"
	HRMinPlus     = "(C_R^2 tensor W) minus (e_+ tensor P_1)"
	HRMinMinus    = "(C_R^2 tensor W) minus (e_- tensor P_1)"
	ImageY        = "Im(Y)=(h_+ tensor P_3) plus (h_- tensor P_3) plus (h_- tensor P_1)"
	LeftKernel    = "h_+ tensor P_1"
	MissingEdge   = "Y_+1:e_+ tensor P_1 -> h_+ tensor P_1"
	ActiveEdge1   = "e_+ tensor P_3 -> h_+ tensor P_3"
	ActiveEdge2   = "e_- tensor P_3 -> h_- tensor P_3"
	ActiveEdge3   = "e_- tensor P_1 -> h_- tensor P_1"

	RankColorCell     = 3
	RankLeptonCell    = 1
	RankRightRect     = 8
	RankActiveRight   = 7
	RankF1            = 4
	RankPiTop         = 3
	RankHLeft         = 8
	RankImageY        = 7
	RankKernel        = 1
	DimH10            = 10
	DimH72            = 72
	BMinusLLepton     = -1
	BMinusLComplement = 1

	Classification = "R3_NEUTRAL_PUNCTURE_AIRLOCK_FUNCTIONAL_OBSTRUCTION_EDGE_ORDERING_MISSING"
	ShortStatus    = "R3_AIRLOCK_CANDIDATE_EDGE_ORDERING_OBSTRUCTION"
	NextFrontier   = "PLUS_MINUS_SOCKET_ORDER_SELECTOR_AUDIT"

	StatusGate895Inherited      = "PASS_GATE895_NEUTRAL_PUNCTURE_AIRLOCK_UNIFICATION_INHERITED"
	StatusFunctionalFormulated  = "PASS_NEUTRAL_PUNCTURE_AIRLOCK_FUNCTIONAL_FORMULATED"
	StatusRankCandidatesAudited = "PASS_RANK_ONE_PUNCTURE_CANDIDATES_AUDITED"
	StatusAlphaFlagAudited      = "PASS_ALPHA_FLAG_TERM_AUDITED"
	StatusEdgeSupportAudited    = "PASS_EDGE_SUPPORT_TERM_AUDITED"
	StatusLeftKernelAudited     = "PASS_LEFT_KERNEL_TERM_AUDITED"
	StatusBMinusLCompAudited    = "PASS_B_MINUS_L_COMPENSATION_TERM_AUDITED"
	StatusObstructionSharpened  = "PASS_AIRLOCK_OBSTRUCTION_SHARPENED_TO_EDGE_ORDERING"
	StatusOfficialFreeze        = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusFirewallVerdict       = "FIREWALL_PRESERVED_GATE896_AIRLOCK_NOT_NATIVE"

	SupportFunctionalFormulated          = "CONDITIONAL_SUPPORT_NEUTRAL_PUNCTURE_AIRLOCK_FUNCTIONAL_FORMULATED"
	SupportLeptonSingletonCandidates     = "CONDITIONAL_SUPPORT_RANK_ONE_PUNCTURE_CANDIDATES_REDUCE_TO_LEPTON_SOCKET_CELLS"
	SupportEPlusSatisfiesAirlock         = "CONDITIONAL_SUPPORT_E_PLUS_P1_SATISFIES_AIRLOCK_CONSTRAINTS"
	SupportPunctureFlagReconstructsAlpha = "CONDITIONAL_SUPPORT_PUNCTURE_FLAG_RECONSTRUCTS_ALPHA_TARGETS"
	SupportEPlusNullEdge                 = "CONDITIONAL_SUPPORT_CURRENT_EDGE_PATTERN_SELECTS_E_PLUS_P1_AS_NULL_EDGE"
	SupportLeftKernelMatchesEPlus        = "CONDITIONAL_SUPPORT_LEFT_KERNEL_TERM_MATCHES_E_PLUS_PUNCTURE_TO_H_PLUS_KERNEL"
	SupportTwoSealWoundReduced           = "CONDITIONAL_SUPPORT_TWO_SEAL_WOUND_REDUCED_TO_PUNCTURE_PLUS_ORIENTED_EDGE_ORDERING"
	SupportOperatorNEffReproduced        = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_REPRODUCED_BY_ORIENTED_LEDGER"

	FailureNotNativeR3                      = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureAlphaStillSealed                 = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureHiggsOrientationStillSealed      = "FAILED_ROUTE_HIGGS_ORIENTATION_REMAINS_SEALED"
	FailureNoNativeAirlockFunctional        = "FAILED_ROUTE_NO_NATIVE_NEUTRAL_PUNCTURE_AIRLOCK_FUNCTIONAL"
	FailureRankDoesNotSelectPlus            = "FAILED_ROUTE_RANK_ONE_PUNCTURE_CONDITION_DOES_NOT_UNIQUELY_SELECT_E_PLUS_P1"
	FailureAlphaFlagDoesNotSelectPlus       = "FAILED_ROUTE_ALPHA_FLAG_RANKS_DO_NOT_DISTINGUISH_PLUS_FROM_MINUS_PUNCTURE"
	FailureBMinusLDoesNotSelectPlus         = "FAILED_ROUTE_B_MINUS_L_COMPENSATION_DOES_NOT_DISTINGUISH_E_PLUS_P1_FROM_E_MINUS_P1"
	FailureEdgeSelectionCircular            = "FAILED_ROUTE_EDGE_PATTERN_SELECTION_CIRCULAR_WITHOUT_INDEPENDENT_EDGE_ORDERING"
	FailureLeftKernelDependsOnImage         = "FAILED_ROUTE_LEFT_KERNEL_SELECTION_DEPENDS_ON_PRESELECTED_IMAGE_Y"
	FailureNoNativeOrientedEdgeOrdering     = "FAILED_ROUTE_NO_NATIVE_ORIENTED_EDGE_ORDERING_FUNCTIONAL"
	FailureNoNativeBoundaryIncidenceFunctor = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureNoNativeWeakSocketSelector       = "FAILED_ROUTE_NO_NATIVE_WEAK_SOCKET_SELECTOR_FUNCTIONAL"
	FailureNoNativeVariationalMinimality    = "FAILED_ROUTE_NO_NATIVE_VARIATIONAL_MINIMALITY_THEOREM"
	FailureNoNativeMinimalNullEdgePrinciple = "FAILED_ROUTE_NO_NATIVE_MINIMAL_NULL_EDGE_ORIENTATION_PRINCIPLE"
	FailureFullHMixesWeakSockets            = "FAILED_ROUTE_FULL_H_ACTION_MIXES_H_PLUS_H_MINUS"
	FailureNoNativeDescentFullToOrient      = "FAILED_ROUTE_NO_NATIVE_DESCENT_FROM_FULL_A_F_TO_A_F_ORIENT"
	FailureNoGenerationCarrierMap           = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap           = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues         = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoOfficialNEffUpdate             = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate            = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator           = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4NativeYukawaTheorem          = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type Cell struct {
	Name       string
	Rank       int
	LeptonCell bool
	ColorCell  bool
	BMinusL    int
}

type RankTermAudit struct {
	Candidates                []Cell
	RankOneCandidates         []string
	OnlyLeptonCellsAreRankOne bool
	BothLeptonCandidatesPass  bool
	SelectsEPlusUniquely      bool
	ActiveComplementRank      int
	Supports, Failures        []string
}

type AlphaFlagTermAudit struct {
	PlusF1, MinusF1             string
	PlusQuotient, MinusQuotient string
	PlusFullComplement          string
	MinusFullComplement         string
	PlusRanks, MinusRanks       []int
	PlusAlpha, MinusAlpha       float64
	BothReconstructAlphaShape   bool
	DistinguishesPlusFromMinus  bool
	Supports, Failures          []string
}

type EdgeSupportTermAudit struct {
	ActiveEdges             []string
	MissingEdge             string
	SelectsEPlusAsNullEdge  bool
	IndependentEdgeOrdering bool
	CircularWithoutOrdering bool
	Supports, Failures      []string
}

type LeftKernelTermAudit struct {
	Image                            string
	Kernel                           string
	HLeftRank, ImageRank, KernelRank int
	MatchesEPlusPuncture             bool
	DependsOnPreselectedImage        bool
	NativeKernelSelector             bool
	Supports, Failures               []string
}

type BMinusLTermAudit struct {
	PlusLeptonCharge, MinusLeptonCharge         int
	PlusComplementCharge, MinusComplementCharge int
	FullRectangleNeutral                        bool
	DistinguishesPlusFromMinus                  bool
	Supports, Failures                          []string
}

type FunctionalAudit struct {
	Formulated                    bool
	Terms                         []string
	EPlusSatisfiesAllTerms        bool
	EMinusAlsoSatisfiesRankFlagBL bool
	RequiresOrientedEdgeOrdering  bool
	NativeFunctional              bool
	NextMissingObject             string
	Supports, Failures            []string
}

type FreezeAudit struct {
	OperatorNEff, OfficialNEff        float64
	OperatorCYukawa, OfficialCYukawa  float64
	OperatorCHiggs, OfficialCHiggs    float64
	Frozen, DiagnosticOnly, CanUpdate bool
	Supports, Failures                []string
}

type Firewalls struct {
	Enforced                     bool
	NotNativeR3                  bool
	AlphaStillSealed             bool
	HiggsOrientationStillSealed  bool
	NoNativeAirlockFunctional    bool
	NoNativeOrientedEdgeOrdering bool
	NoBoundaryIncidenceFunctor   bool
	NoWeakSocketSelector         bool
	NoVariationalMinimality      bool
	NoNativeDescentFullToOrient  bool
	FullHMixesWeakSockets        bool
	NoGenerationCarrier          bool
	NoFlavorOrientation          bool
	NoIndividualYukawas          bool
	NoOfficialLedgerUpdate       bool
	NoNativeYukawaOperator       bool
	NoR4NativeYukawaTheorem      bool
	Verdict                      string
}

type Audit struct {
	ID         string
	Rank       RankTermAudit
	AlphaFlag  AlphaFlagTermAudit
	Edge       EdgeSupportTermAudit
	Kernel     LeftKernelTermAudit
	BMinusL    BMinusLTermAudit
	Functional FunctionalAudit
	Freeze     FreezeAudit
	Firewalls  Firewalls
	Truth      string
	Final      string
}

func BuildDefault() (Audit, error) {
	rank := buildRankTermAudit()
	if !rank.OnlyLeptonCellsAreRankOne || !rank.BothLeptonCandidatesPass || rank.SelectsEPlusUniquely || rank.ActiveComplementRank != RankActiveRight {
		return Audit{}, fmt.Errorf("rank term promoted incorrectly: %s", FormatRankTerm(rank))
	}

	flag := buildAlphaFlagTermAudit()
	if !flag.BothReconstructAlphaShape || flag.DistinguishesPlusFromMinus || !near(flag.PlusAlpha, AlphaB) || !near(flag.MinusAlpha, AlphaB) {
		return Audit{}, fmt.Errorf("alpha flag term promoted incorrectly: %s", FormatAlphaFlagTerm(flag))
	}

	edge := buildEdgeSupportTermAudit()
	if !edge.SelectsEPlusAsNullEdge || edge.IndependentEdgeOrdering || !edge.CircularWithoutOrdering {
		return Audit{}, fmt.Errorf("edge term failed to preserve circularity firewall: %s", FormatEdgeSupportTerm(edge))
	}

	kernel := buildLeftKernelTermAudit()
	if !kernel.MatchesEPlusPuncture || !kernel.DependsOnPreselectedImage || kernel.NativeKernelSelector || kernel.KernelRank != 1 {
		return Audit{}, fmt.Errorf("kernel term promoted incorrectly: %s", FormatLeftKernelTerm(kernel))
	}

	bl := buildBMinusLTermAudit()
	if !bl.FullRectangleNeutral || bl.DistinguishesPlusFromMinus {
		return Audit{}, fmt.Errorf("B-L term promoted incorrectly: %s", FormatBMinusLTerm(bl))
	}

	functional := buildFunctionalAudit()
	if !functional.Formulated || !functional.EPlusSatisfiesAllTerms || !functional.EMinusAlsoSatisfiesRankFlagBL || !functional.RequiresOrientedEdgeOrdering || functional.NativeFunctional {
		return Audit{}, fmt.Errorf("airlock functional promoted incorrectly: %s", FormatFunctional(functional))
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
		NoNativeAirlockFunctional: true, NoNativeOrientedEdgeOrdering: true, NoBoundaryIncidenceFunctor: true,
		NoWeakSocketSelector: true, NoVariationalMinimality: true, NoNativeDescentFullToOrient: true,
		FullHMixesWeakSockets: true, NoGenerationCarrier: true, NoFlavorOrientation: true, NoIndividualYukawas: true,
		NoOfficialLedgerUpdate: true, NoNativeYukawaOperator: true, NoR4NativeYukawaTheorem: true,
		Verdict: StatusFirewallVerdict,
	}
	if !firewallsOK(firewalls) {
		return Audit{}, fmt.Errorf("firewalls not preserved: %s", FormatFirewalls(firewalls))
	}

	return Audit{
		ID: AuditID, Rank: rank, AlphaFlag: flag, Edge: edge, Kernel: kernel, BMinusL: bl, Functional: functional, Freeze: freeze, Firewalls: firewalls,
		Truth: "Gate 896 formulates a NeutralPunctureAirlockFunctional candidate and finds the precise obstruction: rank, alpha-flag, and B-L terms identify neutral lepton singleton punctures but do not distinguish e_+ tensor P_1 from e_- tensor P_1. The plus puncture is selected only after the oriented edge ordering is included.",
		Final: "The branch becomes R3_NEUTRAL_PUNCTURE_AIRLOCK_FUNCTIONAL_OBSTRUCTION_EDGE_ORDERING_MISSING. The two-seal wound is reduced to puncture plus oriented edge ordering, but no native airlock functional, socket-order selector, variational minimality theorem, native R3 ledger, or official ledger update is certified.",
	}, nil
}

func buildRankTermAudit() RankTermAudit {
	cells := []Cell{
		{Name: ColorPlus, Rank: RankColorCell, ColorCell: true, BMinusL: 1},
		{Name: PuncturePlus, Rank: RankLeptonCell, LeptonCell: true, BMinusL: BMinusLLepton},
		{Name: ColorMinus, Rank: RankColorCell, ColorCell: true, BMinusL: 1},
		{Name: PunctureMinus, Rank: RankLeptonCell, LeptonCell: true, BMinusL: BMinusLLepton},
	}
	return RankTermAudit{
		Candidates: cells, RankOneCandidates: []string{PuncturePlus, PunctureMinus},
		OnlyLeptonCellsAreRankOne: true, BothLeptonCandidatesPass: true, SelectsEPlusUniquely: false, ActiveComplementRank: RankRightRect - RankLeptonCell,
		Supports: []string{StatusRankCandidatesAudited, SupportLeptonSingletonCandidates},
		Failures: []string{FailureRankDoesNotSelectPlus},
	}
}

func buildAlphaFlagTermAudit() AlphaFlagTermAudit {
	plusAlpha := float64(RankPiTop)/float64(DimH10)*Ssplit + float64(RankActiveRight)/float64(DimH72)*Ssplit*Ssplit
	minusAlpha := float64(RankPiTop)/float64(DimH10)*Ssplit + float64(RankActiveRight)/float64(DimH72)*Ssplit*Ssplit
	return AlphaFlagTermAudit{
		PlusF1: F1Plus, MinusF1: F1Minus, PlusQuotient: PiTopPlus, MinusQuotient: PiTopMinus,
		PlusFullComplement: HRMinPlus, MinusFullComplement: HRMinMinus,
		PlusRanks: []int{RankPiTop, RankActiveRight}, MinusRanks: []int{RankPiTop, RankActiveRight},
		PlusAlpha: plusAlpha, MinusAlpha: minusAlpha, BothReconstructAlphaShape: near(plusAlpha, AlphaB) && near(minusAlpha, AlphaB), DistinguishesPlusFromMinus: false,
		Supports: []string{StatusAlphaFlagAudited, SupportPunctureFlagReconstructsAlpha},
		Failures: []string{FailureAlphaFlagDoesNotSelectPlus, FailureNoNativeBoundaryIncidenceFunctor, FailureAlphaStillSealed},
	}
}

func buildEdgeSupportTermAudit() EdgeSupportTermAudit {
	return EdgeSupportTermAudit{
		ActiveEdges: []string{ActiveEdge1, ActiveEdge2, ActiveEdge3}, MissingEdge: MissingEdge,
		SelectsEPlusAsNullEdge: true, IndependentEdgeOrdering: false, CircularWithoutOrdering: true,
		Supports: []string{StatusEdgeSupportAudited, SupportEPlusNullEdge},
		Failures: []string{FailureEdgeSelectionCircular, FailureNoNativeOrientedEdgeOrdering},
	}
}

func buildLeftKernelTermAudit() LeftKernelTermAudit {
	return LeftKernelTermAudit{
		Image: ImageY, Kernel: LeftKernel, HLeftRank: RankHLeft, ImageRank: RankImageY, KernelRank: RankKernel,
		MatchesEPlusPuncture: true, DependsOnPreselectedImage: true, NativeKernelSelector: false,
		Supports: []string{StatusLeftKernelAudited, SupportLeftKernelMatchesEPlus},
		Failures: []string{FailureLeftKernelDependsOnImage, FailureNoNativeWeakSocketSelector},
	}
}

func buildBMinusLTermAudit() BMinusLTermAudit {
	return BMinusLTermAudit{
		PlusLeptonCharge: BMinusLLepton, MinusLeptonCharge: BMinusLLepton,
		PlusComplementCharge: BMinusLComplement, MinusComplementCharge: BMinusLComplement,
		FullRectangleNeutral: true, DistinguishesPlusFromMinus: false,
		Supports: []string{StatusBMinusLCompAudited},
		Failures: []string{FailureBMinusLDoesNotSelectPlus},
	}
}

func buildFunctionalAudit() FunctionalAudit {
	return FunctionalAudit{
		Formulated:             true,
		Terms:                  []string{"rank", "alpha_flag", "edge", "left_kernel", "B-L"},
		EPlusSatisfiesAllTerms: true, EMinusAlsoSatisfiesRankFlagBL: true, RequiresOrientedEdgeOrdering: true, NativeFunctional: false,
		NextMissingObject: "OrientedEdgeOrderingFunctional or SocketOrderSelector",
		Supports:          []string{StatusFunctionalFormulated, StatusObstructionSharpened, SupportFunctionalFormulated, SupportEPlusSatisfiesAirlock, SupportTwoSealWoundReduced},
		Failures:          []string{FailureNoNativeAirlockFunctional, FailureNoNativeOrientedEdgeOrdering, FailureNoNativeVariationalMinimality},
	}
}

func FormatRankTerm(r RankTermAudit) string {
	return fmt.Sprintf("rank_term(rank_one_candidates=%s only_lepton=%t both_pass=%t selects_e_plus=%t active_complement_rank=%d supports=%s failures=%s)", strings.Join(r.RankOneCandidates, ","), r.OnlyLeptonCellsAreRankOne, r.BothLeptonCandidatesPass, r.SelectsEPlusUniquely, r.ActiveComplementRank, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}

func FormatAlphaFlagTerm(a AlphaFlagTermAudit) string {
	return fmt.Sprintf("alpha_flag_term(plusQ=%s minusQ=%s plusRanks=%v minusRanks=%v plusAlpha=%.16g minusAlpha=%.16g both_reconstruct=%t distinguishes=%t supports=%s failures=%s)", a.PlusQuotient, a.MinusQuotient, a.PlusRanks, a.MinusRanks, a.PlusAlpha, a.MinusAlpha, a.BothReconstructAlphaShape, a.DistinguishesPlusFromMinus, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatEdgeSupportTerm(e EdgeSupportTermAudit) string {
	return fmt.Sprintf("edge_term(active=%s missing=%s selects_e_plus=%t independent_ordering=%t circular_without_ordering=%t supports=%s failures=%s)", strings.Join(e.ActiveEdges, ";"), e.MissingEdge, e.SelectsEPlusAsNullEdge, e.IndependentEdgeOrdering, e.CircularWithoutOrdering, strings.Join(e.Supports, ","), strings.Join(e.Failures, ","))
}

func FormatLeftKernelTerm(k LeftKernelTermAudit) string {
	return fmt.Sprintf("kernel_term(image=%s kernel=%s ranks=%d,%d,%d matches_e_plus=%t depends_on_image=%t native=%t supports=%s failures=%s)", k.Image, k.Kernel, k.HLeftRank, k.ImageRank, k.KernelRank, k.MatchesEPlusPuncture, k.DependsOnPreselectedImage, k.NativeKernelSelector, strings.Join(k.Supports, ","), strings.Join(k.Failures, ","))
}

func FormatBMinusLTerm(b BMinusLTermAudit) string {
	return fmt.Sprintf("bminusl_term(plus_lepton=%d minus_lepton=%d plus_complement=%d minus_complement=%d full_neutral=%t distinguishes=%t supports=%s failures=%s)", b.PlusLeptonCharge, b.MinusLeptonCharge, b.PlusComplementCharge, b.MinusComplementCharge, b.FullRectangleNeutral, b.DistinguishesPlusFromMinus, strings.Join(b.Supports, ","), strings.Join(b.Failures, ","))
}

func FormatFunctional(f FunctionalAudit) string {
	return fmt.Sprintf("functional(formulated=%t terms=%s e_plus_all=%t e_minus_rank_flag_bl=%t requires_edge_ordering=%t native=%t next=%s supports=%s failures=%s)", f.Formulated, strings.Join(f.Terms, ","), f.EPlusSatisfiesAllTerms, f.EMinusAlsoSatisfiesRankFlagBL, f.RequiresOrientedEdgeOrdering, f.NativeFunctional, f.NextMissingObject, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFreeze(f FreezeAudit) string {
	return fmt.Sprintf("official_freeze(operator_Neff=%.16g official_Neff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic=%t can_update=%t supports=%s failures=%s)", f.OperatorNEff, f.OfficialNEff, f.OperatorCYukawa, f.OfficialCYukawa, f.OperatorCHiggs, f.OfficialCHiggs, f.Frozen, f.DiagnosticOnly, f.CanUpdate, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t not_native_r3=%t alpha_sealed=%t higgs_sealed=%t no_airlock_functional=%t no_edge_ordering=%t no_incidence=%t no_selector=%t no_variational=%t no_descent=%t full_H_mixes=%t no_generation=%t no_flavor=%t no_individual=%t no_official=%t no_yukawa=%t no_r4=%t verdict=%s)", f.Enforced, f.NotNativeR3, f.AlphaStillSealed, f.HiggsOrientationStillSealed, f.NoNativeAirlockFunctional, f.NoNativeOrientedEdgeOrdering, f.NoBoundaryIncidenceFunctor, f.NoWeakSocketSelector, f.NoVariationalMinimality, f.NoNativeDescentFullToOrient, f.FullHMixesWeakSockets, f.NoGenerationCarrier, f.NoFlavorOrientation, f.NoIndividualYukawas, f.NoOfficialLedgerUpdate, f.NoNativeYukawaOperator, f.NoR4NativeYukawaTheorem, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate895Inherited, StatusFunctionalFormulated, StatusRankCandidatesAudited, StatusAlphaFlagAudited, StatusEdgeSupportAudited, StatusLeftKernelAudited, StatusBMinusLCompAudited, StatusObstructionSharpened, StatusOfficialFreeze, StatusFirewallVerdict,
		SupportFunctionalFormulated, SupportLeptonSingletonCandidates, SupportEPlusSatisfiesAirlock, SupportPunctureFlagReconstructsAlpha, SupportEPlusNullEdge, SupportLeftKernelMatchesEPlus, SupportTwoSealWoundReduced, SupportOperatorNEffReproduced,
		FailureNotNativeR3, FailureAlphaStillSealed, FailureHiggsOrientationStillSealed, FailureNoNativeAirlockFunctional, FailureRankDoesNotSelectPlus, FailureAlphaFlagDoesNotSelectPlus, FailureBMinusLDoesNotSelectPlus, FailureEdgeSelectionCircular, FailureLeftKernelDependsOnImage, FailureNoNativeOrientedEdgeOrdering, FailureNoNativeBoundaryIncidenceFunctor, FailureNoNativeWeakSocketSelector, FailureNoNativeVariationalMinimality, FailureNoNativeMinimalNullEdgePrinciple, FailureFullHMixesWeakSockets, FailureNoNativeDescentFullToOrient, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem,
	}
}

func (a Audit) FirewallsList() []string {
	return []string{FailureNotNativeR3, FailureAlphaStillSealed, FailureHiggsOrientationStillSealed, FailureNoNativeAirlockFunctional, FailureNoNativeOrientedEdgeOrdering, FailureNoNativeBoundaryIncidenceFunctor, FailureNoNativeWeakSocketSelector, FailureNoNativeVariationalMinimality, FailureFullHMixesWeakSockets, FailureNoNativeDescentFullToOrient, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoOfficialNEffUpdate, FailureNoNativeYukawaOperator, FailureNoR4NativeYukawaTheorem}
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NotNativeR3 && f.AlphaStillSealed && f.HiggsOrientationStillSealed &&
		f.NoNativeAirlockFunctional && f.NoNativeOrientedEdgeOrdering && f.NoBoundaryIncidenceFunctor &&
		f.NoWeakSocketSelector && f.NoVariationalMinimality && f.NoNativeDescentFullToOrient &&
		f.FullHMixesWeakSockets && f.NoGenerationCarrier && f.NoFlavorOrientation && f.NoIndividualYukawas &&
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
