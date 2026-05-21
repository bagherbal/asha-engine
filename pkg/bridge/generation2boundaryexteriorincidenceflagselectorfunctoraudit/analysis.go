// Package generation2boundaryexteriorincidenceflagselectorfunctoraudit implements
// Gate 879: BoundaryExterior IncidenceFlag Selector Functor Audit.
//
// Gate 879 follows Gate 878's correction that boundary exterior degree is a
// selector/index, not a linear surjection onto a target quotient.  It audits the
// sharper candidate that the nonzero degree poset of the reduced boundary-pair
// exterior response,
//
//	deg Lambda^1 B_2 < deg Lambda^2 B_2,
//
// is an incidence skeleton selecting the puncture-complement flag quotients:
//
//	I_B(1)=F_1/F_0 = Pi_top,
//	I_B(2)=F_2/F_0 = H_R^min.
//
// The gate reconstructs alpha_B from the incidence selector without pretending
// that Lambda^k B_2 spans the quotient spaces.  It does not certify a native
// BoundaryExteriorIncidenceFlagFunctor, cross-lane exclusion theorem, alpha
// theorem, R3 sector trace ledger, R4 Yukawa theorem, or official ledger update.
package generation2boundaryexteriorincidenceflagselectorfunctoraudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE879-BOUNDARY-EXTERIOR-INCIDENCE-FLAG-SELECTOR-FUNCTOR-AUDIT"

	AlphaB    = 0.0003878958469680527
	SBoundary = 0.0012924448188162962

	BoundaryPairDim = 2
	Lambda0B2Dim    = 1
	Lambda1B2Dim    = 2
	Lambda2B2Dim    = 1
	Lambda3B2Dim    = 0

	PunctureRank    = 1
	P1Rank          = 1
	P3Rank          = 3
	WRank           = P1Rank + P3Rank
	RightSocketRank = 2

	F0Rank   = PunctureRank
	F1Rank   = WRank
	F2Rank   = RightSocketRank * WRank
	F1OverF0 = F1Rank - F0Rank
	F2OverF0 = F2Rank - F0Rank
	F2OverF1 = F2Rank - F1Rank
	H10Dim   = F2Rank + BoundaryPairDim
	H72Dim   = 70 + BoundaryPairDim

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	Classification = "BOUNDARY_EXTERIOR_INCIDENCE_FLAG_SELECTOR_FUNCTOR_AUDIT"
	R2Status       = "R2+++++_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_SELECTOR_FUNCTOR_OBSTRUCTION"

	StatusGate878Inherited             = "PASS_GATE878_DEGREE_INDEXED_SELECTOR_TYPE_CORRECTION_INHERITED"
	StatusSourceIncidenceAudited       = "PASS_SOURCE_EXTERIOR_DEGREE_INCIDENCE_SKELETON_AUDITED"
	StatusTargetFlagAudited            = "PASS_TARGET_PUNCTURE_COMPLEMENT_FLAG_AUDITED"
	StatusIncidenceSelectorAudited     = "PASS_INCIDENCE_SELECTOR_CANDIDATE_AUDITED"
	StatusNotLinearSurjectionPreserved = "PASS_SELECTOR_NOT_LINEAR_SURJECTION_FIREWALL_PRESERVED"
	StatusAlphaReconstructed           = "PASS_ALPHA_B_RECONSTRUCTED_BY_INCIDENCE_FLAG_SELECTOR"
	StatusCrossLaneAudited             = "PASS_CROSS_LANE_EXCLUSION_AUDITED_AS_INCIDENCE_SELECTOR_RULE"
	StatusOfficialFreezePreserved      = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusNoObservedDataUsed           = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusFirewallVerdict              = "FIREWALL_PRESERVED_GATE879_INCIDENCE_SELECTOR_NOT_NATIVE_FUNCTOR"

	SupportExteriorDegreeIncidenceShape       = "CONDITIONAL_SUPPORT_BOUNDARY_EXTERIOR_DEGREE_HAS_INCIDENCE_SELECTOR_SHAPE"
	SupportDegreeOneSelectsF1OverF0           = "CONDITIONAL_SUPPORT_DEGREE_ONE_SELECTS_F1_OVER_F0"
	SupportDegreeTwoSelectsF2OverF0           = "CONDITIONAL_SUPPORT_DEGREE_TWO_SELECTS_F2_OVER_F0"
	SupportAlphaReconstructedByIncidence      = "CONDITIONAL_SUPPORT_ALPHA_B_RECONSTRUCTED_BY_INCIDENCE_FLAG_SELECTOR"
	SupportDimensionMismatchSelectorNotLinear = "CONDITIONAL_SUPPORT_DIMENSION_MISMATCH_CORRECTLY_CLASSIFIES_SELECTOR_NOT_LINEAR_MAP"
	SupportSourcePosetOneLessTwo              = "CONDITIONAL_SUPPORT_REDUCED_EXTERIOR_DEGREES_FORM_SOURCE_INCIDENCE_POSET_ONE_LESS_TWO"
	SupportTargetFlagQuotients                = "CONDITIONAL_SUPPORT_TARGETS_ARE_PUNCTURE_COMPLEMENT_FLAG_QUOTIENTS"
	SupportCrossLaneExcludedIfFunctor         = "CONDITIONAL_SUPPORT_CROSS_LANES_EXCLUDED_IF_INCIDENCE_FUNCTOR_CERTIFIED"
	SupportConditionalTraceProxyPlateau       = "CONDITIONAL_SUPPORT_CONDITIONAL_TRACE_PROXY_PLATEAU_REMAINS_COHERENT"
	SupportNoLedgerUpdate                     = "CONDITIONAL_SUPPORT_OPERATOR_DIAGNOSTIC_VALUES_REMAIN_SEPARATED_FROM_OFFICIAL_LEDGER"

	FailureNoNativeIncidenceFunctor   = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureNoNativeCrossLaneExclusion = "FAILED_ROUTE_NO_NATIVE_CROSS_LANE_EXCLUSION_THEOREM"
	FailureNotLinearMap               = "FAILED_ROUTE_LAMBDA_K_B2_NOT_LINEAR_SURJECTION_ONTO_FLAG_QUOTIENT"
	FailureNoNativeDegreeOneIncidence = "FAILED_ROUTE_NO_NATIVE_DEGREE_ONE_TO_F1_OVER_F0_INCIDENCE_SELECTOR"
	FailureNoNativeDegreeTwoIncidence = "FAILED_ROUTE_NO_NATIVE_DEGREE_TWO_TO_F2_OVER_F0_INCIDENCE_SELECTOR"
	FailureNoNativeTargetSelection    = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_TARGET_SELECTION_FUNCTOR"
	FailureAlphaStillSealed           = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNoNativeAlphaSource        = "FAILED_ROUTE_NO_NATIVE_ALPHA_B_SOURCE"
	FailureConditionalProxyNotR3      = "FAILED_ROUTE_CONDITIONAL_TRACE_PROXY_NOT_R3"
	FailureNoNativeR3                 = "FAILED_ROUTE_NO_NATIVE_R3_SECTOR_TRACE_LEDGER"
	FailureNoOfficialNEffUpdate       = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate      = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoSectorTraceMagnitude     = "FAILED_ROUTE_NO_NATIVE_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoNativeYukawaOperator     = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	OperatorNEff, OfficialNEff       float64
	OperatorCYukawa, OfficialCYukawa float64
	OperatorCHiggs, OfficialCHiggs   float64
	OfficialFrozen, CanUpdate        bool
	Supports, Failures               []string
}

type SourceIncidence struct {
	NonzeroDegrees            []int
	DegreeOneBeforeDegreeTwo  bool
	Lambda3Zero               bool
	SelectorIndexNotGenerator bool
	Supports, Failures        []string
}

type TargetFlag struct {
	F0, F1, F2                 string
	RankF0, RankF1, RankF2     int
	F1OverF0Rank, F2OverF0Rank int
	Nested, QuotientsValid     bool
	Supports, Failures         []string
}

type IncidenceSelector struct {
	Degree, ExteriorDim            int
	SelectedQuotient, Target       string
	QuotientRank                   int
	SelectorMode, LinearSurjection bool
	NativeFunctor                  bool
	Supports, Failures             []string
}

type CrossLane struct {
	DegreeOneTarget, DegreeTwoTarget string
	DegreeOneCross, DegreeTwoCross   string
	ExcludedIfFunctor                bool
	NativeExclusion                  bool
	Supports, Failures               []string
}

type AlphaCandidate struct {
	LinearContribution, QuadraticContribution, ReconstructedAlpha float64
	ReconstructedByIncidence                                      bool
	NativeIncidenceFunctor                                        bool
	Supports, Failures                                            []string
}

type R3Assessment struct {
	ConditionalTraceProxyMature  bool
	AlphaNative                  bool
	IncidenceFunctorNative       bool
	EligibleForR3, EligibleForR4 bool
	Supports, Failures           []string
}

type Impact struct {
	Classification, Status                           string
	IncidenceSelectorCandidate                       bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs bool
	CanPromoteToR3, CanPromoteToR4                   bool
}

type Firewalls struct {
	Enforced                   bool
	NoNativeIncidenceFunctor   bool
	NoNativeCrossLaneExclusion bool
	NotLinearSurjection        bool
	NoNativeDegreeOneSelector  bool
	NoNativeDegreeTwoSelector  bool
	NoNativeTargetSelection    bool
	AlphaStillSealed           bool
	NoNativeAlphaSource        bool
	ConditionalNotR3           bool
	NoNativeR3                 bool
	NoOfficialNEffUpdate       bool
	NoCYukawaCHiggsUpdate      bool
	NoSectorTraceMagnitude     bool
	NoNativeYukawaOperator     bool
	Verdict                    string
}

type Audit struct {
	ID        string
	Ledger    Ledger
	Source    SourceIncidence
	Flag      TargetFlag
	DegreeOne IncidenceSelector
	DegreeTwo IncidenceSelector
	CrossLane CrossLane
	Alpha     AlphaCandidate
	R3        R3Assessment
	Impact    Impact
	Firewalls Firewalls
	Truth     string
	Final     string
}

func BuildDefault() (Audit, error) {
	if Lambda1B2Dim == F1OverF0 || Lambda2B2Dim == F2OverF0 {
		return Audit{}, fmt.Errorf("Gate 879 requires exterior degree to be selector, not dimension-matched surjection")
	}
	if F0Rank != 1 || F1Rank != 4 || F2Rank != 8 || F1OverF0 != 3 || F2OverF0 != 7 || H10Dim != 10 || H72Dim != 72 {
		return Audit{}, fmt.Errorf("unexpected Gate 879 flag quotient ledger")
	}
	if Lambda3B2Dim != 0 {
		return Audit{}, fmt.Errorf("boundary pair exterior truncation failed: Lambda^3 B2 dim=%d", Lambda3B2Dim)
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
	source := SourceIncidence{NonzeroDegrees: []int{1, 2}, DegreeOneBeforeDegreeTwo: true, Lambda3Zero: true, SelectorIndexNotGenerator: true, Supports: []string{SupportExteriorDegreeIncidenceShape, SupportSourcePosetOneLessTwo, SupportDimensionMismatchSelectorNotLinear}, Failures: []string{FailureNotLinearMap, FailureNoNativeIncidenceFunctor}}
	flag := TargetFlag{F0: "F_0=p=e_+ tensor P_1", F1: "F_1=e_+ tensor W", F2: "F_2=C_R^2 tensor W", RankF0: F0Rank, RankF1: F1Rank, RankF2: F2Rank, F1OverF0Rank: F1OverF0, F2OverF0Rank: F2OverF0, Nested: true, QuotientsValid: true, Supports: []string{SupportTargetFlagQuotients}, Failures: []string{}}
	degreeOne := IncidenceSelector{Degree: 1, ExteriorDim: Lambda1B2Dim, SelectedQuotient: "F_1/F_0", Target: "Pi_top", QuotientRank: F1OverF0, SelectorMode: true, LinearSurjection: false, NativeFunctor: false, Supports: []string{SupportDegreeOneSelectsF1OverF0, SupportExteriorDegreeIncidenceShape}, Failures: []string{FailureNoNativeDegreeOneIncidence, FailureNoNativeIncidenceFunctor, FailureNotLinearMap}}
	degreeTwo := IncidenceSelector{Degree: 2, ExteriorDim: Lambda2B2Dim, SelectedQuotient: "F_2/F_0", Target: "H_R^min", QuotientRank: F2OverF0, SelectorMode: true, LinearSurjection: false, NativeFunctor: false, Supports: []string{SupportDegreeTwoSelectsF2OverF0, SupportExteriorDegreeIncidenceShape}, Failures: []string{FailureNoNativeDegreeTwoIncidence, FailureNoNativeIncidenceFunctor, FailureNotLinearMap}}
	cross := CrossLane{DegreeOneTarget: "F_1/F_0", DegreeTwoTarget: "F_2/F_0", DegreeOneCross: "F_2/F_0", DegreeTwoCross: "F_1/F_0", ExcludedIfFunctor: true, NativeExclusion: false, Supports: []string{SupportCrossLaneExcludedIfFunctor}, Failures: []string{FailureNoNativeCrossLaneExclusion, FailureNoNativeIncidenceFunctor}}
	alphaCandidate := AlphaCandidate{LinearContribution: linear, QuadraticContribution: quadratic, ReconstructedAlpha: alpha, ReconstructedByIncidence: true, NativeIncidenceFunctor: false, Supports: []string{SupportAlphaReconstructedByIncidence, SupportDegreeOneSelectsF1OverF0, SupportDegreeTwoSelectsF2OverF0}, Failures: []string{FailureNoNativeIncidenceFunctor, FailureAlphaStillSealed, FailureNoNativeAlphaSource}}
	r3 := R3Assessment{ConditionalTraceProxyMature: true, AlphaNative: false, IncidenceFunctorNative: false, EligibleForR3: false, EligibleForR4: false, Supports: []string{SupportConditionalTraceProxyPlateau}, Failures: []string{FailureConditionalProxyNotR3, FailureNoNativeR3, FailureNoSectorTraceMagnitude, FailureNoNativeYukawaOperator}}
	impact := Impact{Classification: Classification, Status: R2Status, IncidenceSelectorCandidate: true, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, CanPromoteToR3: false, CanPromoteToR4: false}
	firewalls := Firewalls{Enforced: true, NoNativeIncidenceFunctor: true, NoNativeCrossLaneExclusion: true, NotLinearSurjection: true, NoNativeDegreeOneSelector: true, NoNativeDegreeTwoSelector: true, NoNativeTargetSelection: true, AlphaStillSealed: true, NoNativeAlphaSource: true, ConditionalNotR3: true, NoNativeR3: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NoSectorTraceMagnitude: true, NoNativeYukawaOperator: true, Verdict: StatusFirewallVerdict}

	return Audit{ID: AuditID, Ledger: ledger, Source: source, Flag: flag, DegreeOne: degreeOne, DegreeTwo: degreeTwo, CrossLane: cross, Alpha: alphaCandidate, R3: r3, Impact: impact, Firewalls: firewalls, Truth: "Gate 879 types boundary exterior degree as an incidence selector into the puncture-complement flag, not as a linear map onto the target quotient.", Final: "BoundaryExteriorIncidenceFlagSelectorFunctor remains a candidate only; alpha_B is still sealed and the conditional trace proxy is not R3."}, nil
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("ledger(operator_N_eff=%.16g official_N_eff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t can_update=%t supports=%s failures=%s)", l.OperatorNEff, l.OfficialNEff, l.OperatorCYukawa, l.OfficialCYukawa, l.OperatorCHiggs, l.OfficialCHiggs, l.OfficialFrozen, l.CanUpdate, strings.Join(l.Supports, ","), strings.Join(l.Failures, ","))
}

func FormatSource(s SourceIncidence) string {
	return fmt.Sprintf("source_incidence(nonzero_degrees=%v one_less_two=%t lambda3_zero=%t selector_index_not_generator=%t supports=%s failures=%s)", s.NonzeroDegrees, s.DegreeOneBeforeDegreeTwo, s.Lambda3Zero, s.SelectorIndexNotGenerator, strings.Join(s.Supports, ","), strings.Join(s.Failures, ","))
}

func FormatFlag(f TargetFlag) string {
	return fmt.Sprintf("flag(%s rank=%d subset %s rank=%d subset %s rank=%d F1/F0=%d F2/F0=%d nested=%t supports=%s failures=%s)", f.F0, f.RankF0, f.F1, f.RankF1, f.F2, f.RankF2, f.F1OverF0Rank, f.F2OverF0Rank, f.Nested, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatSelector(s IncidenceSelector) string {
	return fmt.Sprintf("selector(degree=%d exterior_dim=%d quotient=%s target=%s quotient_rank=%d selector_mode=%t linear_surjection=%t native=%t supports=%s failures=%s)", s.Degree, s.ExteriorDim, s.SelectedQuotient, s.Target, s.QuotientRank, s.SelectorMode, s.LinearSurjection, s.NativeFunctor, strings.Join(s.Supports, ","), strings.Join(s.Failures, ","))
}

func FormatCrossLane(c CrossLane) string {
	return fmt.Sprintf("cross_lane(deg1_target=%s deg2_target=%s deg1_cross=%s deg2_cross=%s excluded_if_functor=%t native=%t supports=%s failures=%s)", c.DegreeOneTarget, c.DegreeTwoTarget, c.DegreeOneCross, c.DegreeTwoCross, c.ExcludedIfFunctor, c.NativeExclusion, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}

func FormatAlpha(a AlphaCandidate) string {
	return fmt.Sprintf("alpha(linear=%.18g quadratic=%.18g reconstructed=%.18g target=%.18g by_incidence=%t native_functor=%t supports=%s failures=%s)", a.LinearContribution, a.QuadraticContribution, a.ReconstructedAlpha, AlphaB, a.ReconstructedByIncidence, a.NativeIncidenceFunctor, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatR3(r R3Assessment) string {
	return fmt.Sprintf("r3(conditional_trace_proxy_mature=%t alpha_native=%t incidence_functor_native=%t eligible_r3=%t eligible_r4=%t supports=%s failures=%s)", r.ConditionalTraceProxyMature, r.AlphaNative, r.IncidenceFunctorNative, r.EligibleForR3, r.EligibleForR4, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("impact(classification=%s status=%s incidence_candidate=%t update_Neff=%t update_CYukawa=%t update_CHiggs=%t promote_R3=%t promote_R4=%t)", i.Classification, i.Status, i.IncidenceSelectorCandidate, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CanPromoteToR3, i.CanPromoteToR4)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t no_incidence_functor=%t no_cross_lane=%t not_linear_surjection=%t no_deg1=%t no_deg2=%t no_target_selection=%t alpha_sealed=%t no_alpha_source=%t conditional_not_r3=%t no_r3=%t no_neff_update=%t no_c_updates=%t no_trace_magnitude=%t no_yukawa=%t verdict=%s)", f.Enforced, f.NoNativeIncidenceFunctor, f.NoNativeCrossLaneExclusion, f.NotLinearSurjection, f.NoNativeDegreeOneSelector, f.NoNativeDegreeTwoSelector, f.NoNativeTargetSelection, f.AlphaStillSealed, f.NoNativeAlphaSource, f.ConditionalNotR3, f.NoNativeR3, f.NoOfficialNEffUpdate, f.NoCYukawaCHiggsUpdate, f.NoSectorTraceMagnitude, f.NoNativeYukawaOperator, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate878Inherited,
		StatusSourceIncidenceAudited,
		StatusTargetFlagAudited,
		StatusIncidenceSelectorAudited,
		StatusNotLinearSurjectionPreserved,
		StatusAlphaReconstructed,
		StatusCrossLaneAudited,
		StatusOfficialFreezePreserved,
		StatusNoObservedDataUsed,
		StatusFirewallVerdict,
		SupportExteriorDegreeIncidenceShape,
		SupportDegreeOneSelectsF1OverF0,
		SupportDegreeTwoSelectsF2OverF0,
		SupportAlphaReconstructedByIncidence,
		SupportDimensionMismatchSelectorNotLinear,
		SupportSourcePosetOneLessTwo,
		SupportTargetFlagQuotients,
		SupportCrossLaneExcludedIfFunctor,
		SupportConditionalTraceProxyPlateau,
		SupportNoLedgerUpdate,
		FailureNoNativeIncidenceFunctor,
		FailureNoNativeCrossLaneExclusion,
		FailureNotLinearMap,
		FailureNoNativeDegreeOneIncidence,
		FailureNoNativeDegreeTwoIncidence,
		FailureNoNativeTargetSelection,
		FailureAlphaStillSealed,
		FailureNoNativeAlphaSource,
		FailureConditionalProxyNotR3,
		FailureNoNativeR3,
		FailureNoOfficialNEffUpdate,
		FailureNoCYukawaCHiggsUpdate,
		FailureNoSectorTraceMagnitude,
		FailureNoNativeYukawaOperator,
	}
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-15 }

func containsAll(have []string, want []string) bool {
	m := map[string]bool{}
	for _, h := range have {
		m[h] = true
	}
	for _, w := range want {
		if !m[w] {
			return false
		}
	}
	return true
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NoNativeIncidenceFunctor && f.NoNativeCrossLaneExclusion && f.NotLinearSurjection && f.NoNativeDegreeOneSelector && f.NoNativeDegreeTwoSelector && f.NoNativeTargetSelection && f.AlphaStillSealed && f.NoNativeAlphaSource && f.ConditionalNotR3 && f.NoNativeR3 && f.NoOfficialNEffUpdate && f.NoCYukawaCHiggsUpdate && f.NoSectorTraceMagnitude && f.NoNativeYukawaOperator && f.Verdict == StatusFirewallVerdict
}
