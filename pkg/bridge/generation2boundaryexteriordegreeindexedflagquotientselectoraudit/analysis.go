// Package generation2boundaryexteriordegreeindexedflagquotientselectoraudit implements
// Gate 878: BoundaryExterior Degree-Indexed FlagQuotient Selector Audit.
//
// Gate 878 follows Gate 877's puncture-complement flag audit and corrects the
// type of the remaining target-selection candidate.  The boundary exterior
// degree is not treated as a linear surjection
//
//	Lambda^k B_2 -> F_k/F_0,
//
// because the dimensions do not match: dim Lambda^1 B_2=2 while rank(F_1/F_0)=3,
// and dim Lambda^2 B_2=1 while rank(F_2/F_0)=7.  Instead, exterior degree is
// audited as a selector/index for a puncture-complement flag quotient:
//
//	Lambda^1 B_2 selects F_1/F_0 = Pi_top,
//	Lambda^2 B_2 selects F_2/F_0 = H_R^min.
//
// The second degree is explicitly cumulative enclosure over the puncture F_0,
// not the pure associated-graded slice F_2/F_1.  This prevents a false dimension
// theorem while preserving the conditional alpha reconstruction.  The gate does
// not certify a native degree-indexed selector functor, alpha theorem, R3 sector
// trace ledger, R4 Yukawa theorem, or official ledger update.
package generation2boundaryexteriordegreeindexedflagquotientselectoraudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE878-BOUNDARY-EXTERIOR-DEGREE-INDEXED-FLAG-QUOTIENT-SELECTOR-AUDIT"

	AlphaB    = 0.0003878958469680527
	SBoundary = 0.0012924448188162962

	BoundaryPairDim = 2
	Lambda1B2Dim    = 2
	Lambda2B2Dim    = 1
	Lambda3B2Dim    = 0

	PunctureRank    = 1
	P1Rank          = 1
	P3Rank          = 3
	WRank           = P1Rank + P3Rank
	RightSocketRank = 2

	F0Rank    = PunctureRank
	F1Rank    = WRank
	F2Rank    = RightSocketRank * WRank
	F1OverF0  = F1Rank - F0Rank
	F2OverF0  = F2Rank - F0Rank
	F2OverF1  = F2Rank - F1Rank
	PiTopRank = F1OverF0
	HRMinRank = F2OverF0
	WrongF2F1 = F2OverF1
	H10Dim    = F2Rank + BoundaryPairDim
	Lambda4V8 = 70
	H72Dim    = Lambda4V8 + BoundaryPairDim

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	Classification = "BOUNDARY_EXTERIOR_DEGREE_INDEXED_FLAG_QUOTIENT_SELECTOR_AUDIT"
	R2Status       = "R2+++++_BOUNDARY_EXTERIOR_DEGREE_INDEXED_FLAG_QUOTIENT_SELECTOR_OBSTRUCTION"

	StatusGate877Inherited         = "PASS_GATE877_FACE_FILTRATION_TARGET_FUNCTOR_INHERITED"
	StatusDimensionMismatchAudited = "PASS_DIMENSION_MISMATCH_SHOWS_SELECTOR_NOT_LINEAR_SURJECTION"
	StatusDegreeOneSelectorAudited = "PASS_DEGREE_ONE_FLAG_QUOTIENT_SELECTOR_AUDITED"
	StatusDegreeTwoSelectorAudited = "PASS_DEGREE_TWO_CUMULATIVE_FLAG_QUOTIENT_SELECTOR_AUDITED"
	StatusWrongGradedSliceRejected = "PASS_WRONG_ASSOCIATED_GRADED_SLICE_F2_OVER_F1_REJECTED"
	StatusCrossLaneReaudited       = "PASS_CROSS_LANE_EXCLUSION_REAUDITED_AS_SELECTOR_RULE"
	StatusAlphaReconstructed       = "PASS_ALPHA_B_RECONSTRUCTED_FROM_DEGREE_INDEXED_FLAG_SELECTORS"
	StatusOfficialFreezePreserved  = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusNoObservedDataUsed       = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusFirewallVerdict          = "FIREWALL_PRESERVED_GATE878_SELECTOR_NOT_NATIVE_FUNCTOR"

	SupportDegreesActAsSelectors           = "CONDITIONAL_SUPPORT_BOUNDARY_EXTERIOR_DEGREES_ACT_AS_FLAG_QUOTIENT_SELECTORS"
	SupportLambda1SelectsF1OverF0          = "CONDITIONAL_SUPPORT_LAMBDA1_SELECTS_F1_OVER_F0"
	SupportLambda2SelectsF2OverF0          = "CONDITIONAL_SUPPORT_LAMBDA2_SELECTS_F2_OVER_F0"
	SupportAlphaTargetsBySelector          = "CONDITIONAL_SUPPORT_ALPHA_TARGETS_RECONSTRUCTED_BY_DEGREE_INDEXED_FLAG_SELECTOR"
	SupportDimensionMismatchTypeCorrection = "CONDITIONAL_SUPPORT_DIMENSION_MISMATCH_SHOWS_TARGET_SELECTION_IS_SELECTOR_NOT_LINEAR_SURJECTION"
	SupportDegreeTwoCumulativeEnclosure    = "CONDITIONAL_SUPPORT_DEGREE_TWO_SELECTS_CUMULATIVE_ENCLOSURE_QUOTIENT_OVER_F0"
	SupportDegreeTwoNotPureSlice           = "CONDITIONAL_SUPPORT_DEGREE_TWO_IS_CUMULATIVE_ENCLOSURE_OVER_F0"
	SupportWrongF2F1Rejected               = "CONDITIONAL_SUPPORT_F2_OVER_F1_RANK_FOUR_REJECTED_AS_ALPHA_TARGET"
	SupportCrossLanesExcludedIfSelector    = "CONDITIONAL_SUPPORT_CROSS_LANES_EXCLUDED_IF_DEGREE_SELECTOR_FUNCTOR_CERTIFIED"
	SupportConditionalTraceProxyPlateau    = "CONDITIONAL_SUPPORT_CONDITIONAL_TRACE_PROXY_PLATEAU_REMAINS_COHERENT"
	SupportNoLedgerUpdate                  = "CONDITIONAL_SUPPORT_OPERATOR_DIAGNOSTIC_VALUES_REMAIN_SEPARATED_FROM_OFFICIAL_LEDGER"

	FailureNoNativeDegreeIndexedSelector = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_DEGREE_INDEXED_FLAG_SELECTOR_FUNCTOR"
	FailureNoNativeCrossLaneExclusion    = "FAILED_ROUTE_NO_NATIVE_CROSS_LANE_EXCLUSION_THEOREM"
	FailureNotLinearMap                  = "FAILED_ROUTE_LAMBDA_K_B2_NOT_LINEAR_SURJECTION_ONTO_FLAG_QUOTIENT"
	FailureNoNativeLambda1Selector       = "FAILED_ROUTE_NO_NATIVE_LAMBDA1B2_SELECTOR_FOR_F1_OVER_F0"
	FailureNoNativeLambda2Selector       = "FAILED_ROUTE_NO_NATIVE_LAMBDA2B2_SELECTOR_FOR_F2_OVER_F0"
	FailureDegreeTwoNotF2OverF1          = "FAILED_ROUTE_DEGREE_TWO_IS_NOT_PURE_GRADED_SLICE_F2_OVER_F1"
	FailureNoNativeTargetSelection       = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_TARGET_SELECTION_FUNCTOR"
	FailureAlphaStillSealed              = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNoNativeAlphaSource           = "FAILED_ROUTE_NO_NATIVE_ALPHA_B_SOURCE"
	FailureConditionalProxyNotR3         = "FAILED_ROUTE_CONDITIONAL_TRACE_PROXY_NOT_R3"
	FailureNoNativeR3                    = "FAILED_ROUTE_NO_NATIVE_R3_SECTOR_TRACE_LEDGER"
	FailureNoOfficialNEffUpdate          = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate         = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoSectorTraceMagnitude        = "FAILED_ROUTE_NO_NATIVE_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoNativeYukawaOperator        = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	OperatorNEff, OfficialNEff       float64
	OperatorCYukawa, OfficialCYukawa float64
	OperatorCHiggs, OfficialCHiggs   float64
	OfficialFrozen, CanUpdate        bool
	Supports, Failures               []string
}

type Flag struct {
	F0, F1, F2             string
	RankF0, RankF1, RankF2 int
	ValidNestedFlag        bool
	Supports, Failures     []string
}

type DimensionMismatch struct {
	Lambda1Dim, Lambda2Dim      int
	F1OverF0Rank, F2OverF0Rank  int
	SelectorNotLinearSurjection bool
	Supports, Failures          []string
}

type DegreeSelector struct {
	Degree, Mode, SelectedQuotient, Target string
	ExteriorDim, QuotientRank, TargetRank  int
	SelectorMode, LinearSurjection         bool
	NativeSelector                         bool
	Supports, Failures                     []string
}

type RejectedTarget struct {
	Degree, RejectedQuotient, Reason string
	RejectedRank, RequiredRank       int
	Rejected                         bool
	Supports, Failures               []string
}

type CrossLane struct {
	Lambda1Target, Lambda2Target string
	Lambda1Cross, Lambda2Cross   string
	ExcludedIfSelector           bool
	NativeExclusion              bool
	Supports, Failures           []string
}

type AlphaCandidate struct {
	LinearContribution, QuadraticContribution, ReconstructedAlpha float64
	ReconstructedFromSelectors                                    bool
	NativeSelectorFunctor                                         bool
	Supports, Failures                                            []string
}

type R3Assessment struct {
	ConditionalTraceProxyMature  bool
	AlphaNative                  bool
	TargetSelectionNative        bool
	EligibleForR3, EligibleForR4 bool
	Supports, Failures           []string
}

type Impact struct {
	Classification, Status                           string
	DegreeIndexedSelectorCandidate                   bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs bool
	CanPromoteToR3, CanPromoteToR4                   bool
}

type Firewalls struct {
	Enforced                      bool
	NoNativeDegreeIndexedSelector bool
	NoNativeCrossLaneExclusion    bool
	NotLinearSurjection           bool
	NoNativeLambda1Selector       bool
	NoNativeLambda2Selector       bool
	DegreeTwoNotF2OverF1          bool
	NoNativeTargetSelection       bool
	AlphaStillSealed              bool
	NoNativeAlphaSource           bool
	ConditionalNotR3              bool
	NoNativeR3                    bool
	NoOfficialNEffUpdate          bool
	NoCYukawaCHiggsUpdate         bool
	NoSectorTraceMagnitude        bool
	NoNativeYukawaOperator        bool
	Verdict                       string
}

type Audit struct {
	ID           string
	Ledger       Ledger
	Flag         Flag
	Mismatch     DimensionMismatch
	DegreeOne    DegreeSelector
	DegreeTwo    DegreeSelector
	WrongSlice   RejectedTarget
	CrossLane    CrossLane
	Alpha        AlphaCandidate
	R3           R3Assessment
	Impact       Impact
	Firewalls    Firewalls
	Truth, Final string
}

func BuildDefault() (Audit, error) {
	if F0Rank != 1 || F1Rank != 4 || F2Rank != 8 || F1OverF0 != 3 || F2OverF0 != 7 || F2OverF1 != 4 || H10Dim != 10 || H72Dim != 72 {
		return Audit{}, fmt.Errorf("unexpected Gate 878 flag quotient rank ledger")
	}
	if Lambda1B2Dim == F1OverF0 || Lambda2B2Dim == F2OverF0 {
		return Audit{}, fmt.Errorf("dimension mismatch vanished unexpectedly")
	}
	linear := float64(F1OverF0) / float64(H10Dim) * SBoundary
	quadratic := float64(F2OverF0) / float64(H72Dim) * SBoundary * SBoundary
	alpha := linear + quadratic
	if !near(alpha, AlphaB) {
		return Audit{}, fmt.Errorf("alpha reconstruction drift: got %.18g want %.18g", alpha, AlphaB)
	}
	if near(OperatorNEffDiagnostic, OfficialNEffFrozen) || near(OperatorCYukawaDiagnostic, OfficialCYukawaFrozen) || near(OperatorCHiggsDiagnostic, OfficialCHiggsFrozen) {
		return Audit{}, fmt.Errorf("operator and official ledgers unexpectedly collapsed")
	}

	ledger := Ledger{OperatorNEff: OperatorNEffDiagnostic, OfficialNEff: OfficialNEffFrozen, OperatorCYukawa: OperatorCYukawaDiagnostic, OfficialCYukawa: OfficialCYukawaFrozen, OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen, OfficialFrozen: true, CanUpdate: false, Supports: []string{SupportConditionalTraceProxyPlateau, SupportNoLedgerUpdate}, Failures: []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate}}
	flag := Flag{F0: "p=e_+ tensor P_1", F1: "e_+ tensor W", F2: "C_R^2 tensor W", RankF0: F0Rank, RankF1: F1Rank, RankF2: F2Rank, ValidNestedFlag: true, Supports: []string{SupportAlphaTargetsBySelector}, Failures: []string{FailureNoNativeDegreeIndexedSelector}}
	mismatch := DimensionMismatch{Lambda1Dim: Lambda1B2Dim, Lambda2Dim: Lambda2B2Dim, F1OverF0Rank: F1OverF0, F2OverF0Rank: F2OverF0, SelectorNotLinearSurjection: true, Supports: []string{SupportDimensionMismatchTypeCorrection}, Failures: []string{FailureNotLinearMap}}
	degreeOne := DegreeSelector{Degree: "Lambda^1 B_2", Mode: "degree-indexed selector", SelectedQuotient: "F_1/F_0", Target: "Pi_top", ExteriorDim: Lambda1B2Dim, QuotientRank: F1OverF0, TargetRank: PiTopRank, SelectorMode: true, LinearSurjection: false, NativeSelector: false, Supports: []string{SupportDegreesActAsSelectors, SupportLambda1SelectsF1OverF0, SupportDimensionMismatchTypeCorrection, SupportAlphaTargetsBySelector}, Failures: []string{FailureNoNativeDegreeIndexedSelector, FailureNoNativeLambda1Selector, FailureNotLinearMap}}
	degreeTwo := DegreeSelector{Degree: "Lambda^2 B_2", Mode: "cumulative enclosure selector", SelectedQuotient: "F_2/F_0", Target: "H_R^min", ExteriorDim: Lambda2B2Dim, QuotientRank: F2OverF0, TargetRank: HRMinRank, SelectorMode: true, LinearSurjection: false, NativeSelector: false, Supports: []string{SupportDegreesActAsSelectors, SupportLambda2SelectsF2OverF0, SupportDimensionMismatchTypeCorrection, SupportDegreeTwoCumulativeEnclosure, SupportDegreeTwoNotPureSlice, SupportAlphaTargetsBySelector}, Failures: []string{FailureNoNativeDegreeIndexedSelector, FailureNoNativeLambda2Selector, FailureNotLinearMap}}
	wrongSlice := RejectedTarget{Degree: "Lambda^2 B_2", RejectedQuotient: "F_2/F_1", Reason: "rank four pure associated-graded slice does not match required active rank-seven alpha target", RejectedRank: F2OverF1, RequiredRank: F2OverF0, Rejected: true, Supports: []string{SupportWrongF2F1Rejected, SupportDegreeTwoNotPureSlice}, Failures: []string{FailureDegreeTwoNotF2OverF1}}
	cross := CrossLane{Lambda1Target: "F_1/F_0", Lambda2Target: "F_2/F_0", Lambda1Cross: "F_2/F_0", Lambda2Cross: "F_1/F_0", ExcludedIfSelector: true, NativeExclusion: false, Supports: []string{SupportCrossLanesExcludedIfSelector}, Failures: []string{FailureNoNativeCrossLaneExclusion, FailureNoNativeDegreeIndexedSelector}}
	alphaCandidate := AlphaCandidate{LinearContribution: linear, QuadraticContribution: quadratic, ReconstructedAlpha: alpha, ReconstructedFromSelectors: true, NativeSelectorFunctor: false, Supports: []string{SupportAlphaTargetsBySelector, SupportLambda1SelectsF1OverF0, SupportLambda2SelectsF2OverF0}, Failures: []string{FailureNoNativeDegreeIndexedSelector, FailureAlphaStillSealed, FailureNoNativeAlphaSource}}
	r3 := R3Assessment{ConditionalTraceProxyMature: true, AlphaNative: false, TargetSelectionNative: false, EligibleForR3: false, EligibleForR4: false, Supports: []string{SupportConditionalTraceProxyPlateau}, Failures: []string{FailureConditionalProxyNotR3, FailureNoNativeR3, FailureNoSectorTraceMagnitude, FailureNoNativeYukawaOperator}}
	impact := Impact{Classification: Classification, Status: R2Status, DegreeIndexedSelectorCandidate: true, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, CanPromoteToR3: false, CanPromoteToR4: false}
	firewalls := Firewalls{Enforced: true, NoNativeDegreeIndexedSelector: true, NoNativeCrossLaneExclusion: true, NotLinearSurjection: true, NoNativeLambda1Selector: true, NoNativeLambda2Selector: true, DegreeTwoNotF2OverF1: true, NoNativeTargetSelection: true, AlphaStillSealed: true, NoNativeAlphaSource: true, ConditionalNotR3: true, NoNativeR3: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NoSectorTraceMagnitude: true, NoNativeYukawaOperator: true, Verdict: StatusFirewallVerdict}

	return Audit{ID: AuditID, Ledger: ledger, Flag: flag, Mismatch: mismatch, DegreeOne: degreeOne, DegreeTwo: degreeTwo, WrongSlice: wrongSlice, CrossLane: cross, Alpha: alphaCandidate, R3: r3, Impact: impact, Firewalls: firewalls, Truth: "Gate 878 corrects the type of Gate 877's target-selection candidate: boundary exterior degree is a selector/index for a puncture-complement flag quotient, not a linear surjection onto that quotient. Lambda^1 B_2 selects F_1/F_0=Pi_top, while Lambda^2 B_2 selects the cumulative enclosure F_2/F_0=H_R^min, not the pure slice F_2/F_1.", Final: "VERDICT: DEGREE_INDEXED_FLAG_QUOTIENT_SELECTOR_CANDIDATE_FOUND_BUT_NO_NATIVE_SELECTOR_FUNCTOR"}, nil
}

func Statuses() []string {
	return []string{StatusGate877Inherited, StatusDimensionMismatchAudited, StatusDegreeOneSelectorAudited, StatusDegreeTwoSelectorAudited, StatusWrongGradedSliceRejected, StatusCrossLaneReaudited, StatusAlphaReconstructed, StatusOfficialFreezePreserved, StatusNoObservedDataUsed, StatusFirewallVerdict, SupportDegreesActAsSelectors, SupportLambda1SelectsF1OverF0, SupportLambda2SelectsF2OverF0, SupportAlphaTargetsBySelector, SupportDimensionMismatchTypeCorrection, SupportDegreeTwoCumulativeEnclosure, SupportDegreeTwoNotPureSlice, SupportWrongF2F1Rejected, SupportCrossLanesExcludedIfSelector, FailureNoNativeDegreeIndexedSelector, FailureNotLinearMap, FailureNoNativeLambda1Selector, FailureNoNativeLambda2Selector, FailureDegreeTwoNotF2OverF1, FailureNoNativeCrossLaneExclusion, FailureAlphaStillSealed, FailureConditionalProxyNotR3, FailureNoNativeR3}
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("operator_N_eff=%.16g official_N_eff=%.16g operator_C_Yukawa=%.16g official_C_Yukawa=%.16g operator_C_Higgs=%.16g official_C_Higgs=%.16g frozen=%t update=%t supports=%s failures=%s", l.OperatorNEff, l.OfficialNEff, l.OperatorCYukawa, l.OfficialCYukawa, l.OperatorCHiggs, l.OfficialCHiggs, l.OfficialFrozen, l.CanUpdate, strings.Join(l.Supports, ","), strings.Join(l.Failures, ","))
}
func FormatFlag(f Flag) string {
	return fmt.Sprintf("F0=%q rank=%d F1=%q rank=%d F2=%q rank=%d nested=%t supports=%s failures=%s", f.F0, f.RankF0, f.F1, f.RankF1, f.F2, f.RankF2, f.ValidNestedFlag, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}
func FormatMismatch(m DimensionMismatch) string {
	return fmt.Sprintf("dim_Lambda1=%d rank_F1_F0=%d dim_Lambda2=%d rank_F2_F0=%d selector_not_linear_surjection=%t supports=%s failures=%s", m.Lambda1Dim, m.F1OverF0Rank, m.Lambda2Dim, m.F2OverF0Rank, m.SelectorNotLinearSurjection, strings.Join(m.Supports, ","), strings.Join(m.Failures, ","))
}
func FormatSelector(s DegreeSelector) string {
	return fmt.Sprintf("degree=%q mode=%q quotient=%q target=%q exterior_dim=%d quotient_rank=%d target_rank=%d selector=%t linear_surjection=%t native=%t supports=%s failures=%s", s.Degree, s.Mode, s.SelectedQuotient, s.Target, s.ExteriorDim, s.QuotientRank, s.TargetRank, s.SelectorMode, s.LinearSurjection, s.NativeSelector, strings.Join(s.Supports, ","), strings.Join(s.Failures, ","))
}
func FormatRejected(r RejectedTarget) string {
	return fmt.Sprintf("degree=%q rejected=%q rejected_rank=%d required_rank=%d is_rejected=%t reason=%q supports=%s failures=%s", r.Degree, r.RejectedQuotient, r.RejectedRank, r.RequiredRank, r.Rejected, r.Reason, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}
func FormatCrossLane(c CrossLane) string {
	return fmt.Sprintf("lambda1_target=%q lambda2_target=%q lambda1_cross=%q lambda2_cross=%q excluded_if_selector=%t native_exclusion=%t supports=%s failures=%s", c.Lambda1Target, c.Lambda2Target, c.Lambda1Cross, c.Lambda2Cross, c.ExcludedIfSelector, c.NativeExclusion, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}
func FormatAlpha(a AlphaCandidate) string {
	return fmt.Sprintf("linear=%.18g quadratic=%.18g alpha=%.18g from_selectors=%t native=%t supports=%s failures=%s", a.LinearContribution, a.QuadraticContribution, a.ReconstructedAlpha, a.ReconstructedFromSelectors, a.NativeSelectorFunctor, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatR3(r R3Assessment) string {
	return fmt.Sprintf("conditional_trace_proxy=%t alpha_native=%t target_native=%t r3=%t r4=%t supports=%s failures=%s", r.ConditionalTraceProxyMature, r.AlphaNative, r.TargetSelectionNative, r.EligibleForR3, r.EligibleForR4, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}
func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s status=%s selector_candidate=%t update_N_eff=%t update_C_Yukawa=%t update_C_Higgs=%t r3=%t r4=%t", i.Classification, i.Status, i.DegreeIndexedSelectorCandidate, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CanPromoteToR3, i.CanPromoteToR4)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%t no_native_degree_indexed_selector=%t no_cross_lane=%t not_linear_surjection=%t no_lambda1=%t no_lambda2=%t degree2_not_F2_F1=%t no_target_selection=%t alpha_sealed=%t no_native_alpha=%t conditional_not_r3=%t no_native_r3=%t no_N_eff_update=%t no_C_updates=%t no_sector_trace=%t no_yukawa_theorem=%t verdict=%s", f.Enforced, f.NoNativeDegreeIndexedSelector, f.NoNativeCrossLaneExclusion, f.NotLinearSurjection, f.NoNativeLambda1Selector, f.NoNativeLambda2Selector, f.DegreeTwoNotF2OverF1, f.NoNativeTargetSelection, f.AlphaStillSealed, f.NoNativeAlphaSource, f.ConditionalNotR3, f.NoNativeR3, f.NoOfficialNEffUpdate, f.NoCYukawaCHiggsUpdate, f.NoSectorTraceMagnitude, f.NoNativeYukawaOperator, f.Verdict)
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-15 }
func containsAll(have []string, want []string) bool {
	set := map[string]bool{}
	for _, h := range have {
		set[h] = true
	}
	for _, w := range want {
		if !set[w] {
			return false
		}
	}
	return true
}
func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NoNativeDegreeIndexedSelector && f.NoNativeCrossLaneExclusion && f.NotLinearSurjection && f.NoNativeLambda1Selector && f.NoNativeLambda2Selector && f.DegreeTwoNotF2OverF1 && f.NoNativeTargetSelection && f.AlphaStillSealed && f.NoNativeAlphaSource && f.ConditionalNotR3 && f.NoNativeR3 && f.NoOfficialNEffUpdate && f.NoCYukawaCHiggsUpdate && f.NoSectorTraceMagnitude && f.NoNativeYukawaOperator && f.Verdict == StatusFirewallVerdict
}
