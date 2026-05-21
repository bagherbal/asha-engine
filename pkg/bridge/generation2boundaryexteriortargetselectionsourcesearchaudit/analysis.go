// Package generation2boundaryexteriortargetselectionsourcesearchaudit implements
// Gate 875: BoundaryExterior Target-Selection Source Search Audit.
//
// Gate 875 follows Gate 874's conditional Yukawa trace-proxy ledger freeze.
// It does not recompute the conditional trace proxy and does not attempt an
// official ledger update.  It searches the now-isolated remaining R3 wound:
// a native/source-typed mechanism for the target assignments
//
//	Lambda^1 B_2 -> Pi_top,
//	Lambda^2 B_2 -> H_R^min,
//
// together with the cross-lane exclusions
//
//	Lambda^1 B_2 not -> H_R^min,
//	Lambda^2 B_2 not -> Pi_top.
//
// Three candidate source routes are audited: the puncture/complement route, the
// boundary degree/support-codimension route, and the trace-normalization chamber
// route.  The result is intentionally conservative: the puncture/complement
// route is identified as the strongest internal source candidate, but no native
// BoundaryExteriorTargetSelectionFunctor, cross-lane exclusion theorem, alpha
// theorem, R3 sector trace ledger, or official ledger update is certified.
package generation2boundaryexteriortargetselectionsourcesearchaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE875-BOUNDARY-EXTERIOR-TARGET-SELECTION-SOURCE-SEARCH-AUDIT"

	AlphaB    = 0.0003878958469680527
	SBoundary = 0.0012924448188162962

	BoundaryPairDim = 2
	Lambda1B2Dim    = 2
	Lambda2B2Dim    = 1
	Lambda3B2Dim    = 0

	PiTopRank      = 3
	PunctureRank   = 1
	HRambientRank  = 8
	HRminRank      = HRambientRank - PunctureRank
	H10Dim         = HRambientRank + BoundaryPairDim
	Lambda4V8Rank  = 70
	H72Dim         = Lambda4V8Rank + BoundaryPairDim
	RightFullRank  = 8
	RightActiveDim = 7

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	Classification = "BOUNDARY_EXTERIOR_TARGET_SELECTION_SOURCE_SEARCH_AUDIT"
	R2Status       = "R2+++++_BOUNDARY_EXTERIOR_TARGET_SELECTION_SOURCE_SEARCH_OBSTRUCTION"

	StatusGate874Inherited        = "PASS_GATE874_CONDITIONAL_PROXY_LEDGER_FREEZE_INHERITED"
	StatusRemainingWoundInherited = "PASS_REMAINING_WOUND_IDENTIFIED_AS_BOUNDARY_EXTERIOR_TARGET_SELECTION"
	StatusPunctureRouteAudited    = "PASS_PUNCTURE_COMPLEMENT_SOURCE_ROUTE_AUDITED"
	StatusCodimRouteAudited       = "PASS_BOUNDARY_DEGREE_SUPPORT_CODIMENSION_ROUTE_AUDITED"
	StatusChamberRouteAudited     = "PASS_TRACE_NORMALIZATION_CHAMBER_ROUTE_AUDITED"
	StatusCrossLaneAudited        = "PASS_CROSS_LANE_EXCLUSION_REQUIREMENT_REAUDITED"
	StatusAlphaStillReconstructs  = "PASS_ALPHA_B_RECONSTRUCTED_FROM_CURRENT_TARGET_CANDIDATES"
	StatusOfficialFreezePreserved = "PASS_OFFICIAL_N_EFF_C_YUKAWA_C_HIGGS_FREEZE_PRESERVED"
	StatusNoObservedDataUsed      = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusFirewallVerdict         = "FIREWALL_PRESERVED_GATE875_SOURCE_SEARCH_NOT_TARGET_SELECTION_THEOREM"

	SupportConditionalTraceProxyPlateau      = "CONDITIONAL_SUPPORT_CONDITIONAL_TRACE_PROXY_PLATEAU_INHERITED"
	SupportPunctureComplementStrongestRoute  = "CONDITIONAL_SUPPORT_PUNCTURE_COMPLEMENT_IS_STRONGEST_TARGET_SELECTION_SOURCE"
	SupportExposureVisibleComplementPiTop    = "CONDITIONAL_SUPPORT_EXPOSURE_TARGETS_VISIBLE_COLOR_COMPLEMENT_PI_TOP"
	SupportEnclosurePuncturedActiveDomain    = "CONDITIONAL_SUPPORT_ENCLOSURE_TARGETS_ACTIVE_PUNCTURED_DOMAIN_H_R_MIN"
	SupportDegreeCodimRouteCandidate         = "CONDITIONAL_SUPPORT_BOUNDARY_DEGREE_SUPPORT_CODIMENSION_ROUTE_AUDITED_AS_CANDIDATE"
	SupportH10H72TypedResponseChambers       = "CONDITIONAL_SUPPORT_H10_AND_H72_ARE_TYPED_RESPONSE_CHAMBERS"
	SupportR3WoundSharpened                  = "CONDITIONAL_SUPPORT_NATIVE_R3_WOUND_REDUCES_TO_BOUNDARY_EXTERIOR_TARGET_SELECTION_FUNCTOR"
	SupportAlphaSealStillCoherent            = "CONDITIONAL_SUPPORT_ALPHA_B_EXTERIOR_SEAL_REMAINS_COHERENT_UNDER_SOURCE_SEARCH"
	SupportCrossLaneExclusionCandidateByType = "CONDITIONAL_SUPPORT_CROSS_LANE_EXCLUSION_HAS_PUNCTURE_EXPOSURE_ENCLOSURE_TYPE_CANDIDATE"
	SupportNoLedgerUpdate                    = "CONDITIONAL_SUPPORT_OPERATOR_DIAGNOSTIC_VALUES_REMAIN_SEPARATED_FROM_OFFICIAL_LEDGER"

	FailureNoNativeTargetSelectionFunctor    = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_TARGET_SELECTION_FUNCTOR"
	FailureNoNativeExposureToPiTopMap        = "FAILED_ROUTE_NO_NATIVE_LAMBDA1B2_TO_PI_TOP_MAP"
	FailureNoNativeEnclosureToHRMinMap       = "FAILED_ROUTE_NO_NATIVE_LAMBDA2B2_TO_H_R_MIN_MAP"
	FailureNoNativeCrossLaneExclusionTheorem = "FAILED_ROUTE_NO_NATIVE_CROSS_LANE_EXCLUSION_THEOREM"
	FailurePunctureComplementNotFunctor      = "FAILED_ROUTE_PUNCTURE_COMPLEMENT_ROUTE_NOT_NATIVE_TARGET_SELECTION_FUNCTOR"
	FailureCodimRouteNotFunctor              = "FAILED_ROUTE_BOUNDARY_DEGREE_SUPPORT_CODIMENSION_ROUTE_NOT_NATIVE_FUNCTOR"
	FailureChamberRouteNotFunctor            = "FAILED_ROUTE_TRACE_NORMALIZATION_CHAMBER_ROUTE_NOT_TARGET_SELECTION_FUNCTOR"
	FailureResponseChamberTypingNotTheorem   = "FAILED_ROUTE_RESPONSE_CHAMBER_TYPING_NOT_TARGET_SELECTION_THEOREM"
	FailureAlphaStillSealed                  = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNoNativeAlphaSource               = "FAILED_ROUTE_NO_NATIVE_ALPHA_B_SOURCE"
	FailureNoNativeSocketMagnitudeSource     = "FAILED_ROUTE_NO_NATIVE_SOCKET_MAGNITUDE_SOURCE"
	FailureNoSectorTraceMagnitude            = "FAILED_ROUTE_NO_NATIVE_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureConditionalProxyNotR3             = "FAILED_ROUTE_CONDITIONAL_TRACE_PROXY_NOT_R3"
	FailureNoNativeR3                        = "FAILED_ROUTE_NO_NATIVE_R3_SECTOR_TRACE_LEDGER"
	FailureNoOfficialNEffUpdate              = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate             = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNumericalYukawaValues           = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureNotPhysicalYukawaSpectrum         = "FAILED_ROUTE_CONDITIONAL_TRACE_PROXY_NOT_PHYSICAL_YUKAWA_SPECTRUM"
	FailureNotR4                             = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	OperatorNEff, OfficialNEff       float64
	OperatorCYukawa, OfficialCYukawa float64
	OperatorCHiggs, OfficialCHiggs   float64
	OfficialFrozen                   bool
	CanUpdateOfficial                bool
	Supports, Failures               []string
}

type RemainingWound struct {
	NeedLambda1ToPiTop, NeedLambda2ToHRMin     bool
	NeedNoLambda1ToHRMin, NeedNoLambda2ToPiTop bool
	AlphaSealed, R3Blocked                     bool
	ExactWound, MissingObject                  string
	Supports, Failures                         []string
}

type SourceRoute struct {
	Name               string
	Premise            string
	ExposureTarget     string
	EnclosureTarget    string
	CrossLaneExclusion string
	NativeFunctor      bool
	Strength           string
	Supports, Failures []string
}

type AlphaReconstruction struct {
	LinearContribution, QuadraticContribution float64
	ReconstructedAlpha                        float64
	ShapeCoherent                             bool
	TargetSelectionNative                     bool
	Supports, Failures                        []string
}

type R3Assessment struct {
	ConditionalTraceProxyMature  bool
	TargetSelectionFunctor       bool
	AlphaNative                  bool
	SocketMagnitudeNative        bool
	SectorTraceMagnitudeNative   bool
	EligibleForR3, EligibleForR4 bool
	Supports, Failures           []string
}

type Impact struct {
	Classification, Status                                                           string
	StrongestRoute                                                                   string
	PunctureRouteStrongest, CodimRouteCandidate, ChamberRouteCandidate               bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4 bool
}

type Firewalls struct {
	Enforced                                                                                                       bool
	NoNativeTargetSelectionFunctor, NoNativeLambda1ToPiTop, NoNativeLambda2ToHRMin, NoNativeCrossLaneExclusion     bool
	PunctureComplementNotFunctor, CodimRouteNotFunctor, ChamberRouteNotFunctor, ResponseChamberTypingNotTheorem    bool
	AlphaStillSealed, NoNativeAlphaSource, NoNativeSocketMagnitudeSource, NoSectorTraceMagnitude, ConditionalNotR3 bool
	NoNativeR3, NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate, NoNumericalYukawa, NotPhysicalYukawaSpectrum, NotR4   bool
	Verdict                                                                                                        string
}

type Audit struct {
	ID            string
	Ledger        Ledger
	Wound         RemainingWound
	PunctureRoute SourceRoute
	CodimRoute    SourceRoute
	ChamberRoute  SourceRoute
	Candidate     AlphaReconstruction
	R3            R3Assessment
	Impact        Impact
	Firewalls     Firewalls
	Truth, Final  string
}

func BuildDefault() (Audit, error) {
	if PiTopRank != 3 || HRminRank != 7 || HRambientRank != 8 || H10Dim != 10 || H72Dim != 72 {
		return Audit{}, fmt.Errorf("unexpected Gate 875 rank/chamber ledger")
	}
	linear := float64(PiTopRank) / float64(H10Dim) * SBoundary
	quadratic := float64(HRminRank) / float64(H72Dim) * SBoundary * SBoundary
	alpha := linear + quadratic
	if !near(alpha, AlphaB) {
		return Audit{}, fmt.Errorf("alpha reconstruction drift: got %.18g want %.18g", alpha, AlphaB)
	}
	if near(OperatorNEffDiagnostic, OfficialNEffFrozen) || near(OperatorCYukawaDiagnostic, OfficialCYukawaFrozen) || near(OperatorCHiggsDiagnostic, OfficialCHiggsFrozen) {
		return Audit{}, fmt.Errorf("operator and official ledgers unexpectedly collapsed")
	}

	ledger := Ledger{
		OperatorNEff:      OperatorNEffDiagnostic,
		OfficialNEff:      OfficialNEffFrozen,
		OperatorCYukawa:   OperatorCYukawaDiagnostic,
		OfficialCYukawa:   OfficialCYukawaFrozen,
		OperatorCHiggs:    OperatorCHiggsDiagnostic,
		OfficialCHiggs:    OfficialCHiggsFrozen,
		OfficialFrozen:    true,
		CanUpdateOfficial: false,
		Supports:          []string{SupportConditionalTraceProxyPlateau, SupportNoLedgerUpdate},
		Failures:          []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureConditionalProxyNotR3},
	}

	wound := RemainingWound{
		NeedLambda1ToPiTop:   true,
		NeedLambda2ToHRMin:   true,
		NeedNoLambda1ToHRMin: true,
		NeedNoLambda2ToPiTop: true,
		AlphaSealed:          true,
		R3Blocked:            true,
		ExactWound:           "Lambda^1 B_2 -> Pi_top; Lambda^2 B_2 -> H_R^min; cross-lane exclusions",
		MissingObject:        "BoundaryExteriorTargetSelectionFunctor",
		Supports:             []string{SupportR3WoundSharpened},
		Failures:             []string{FailureNoNativeTargetSelectionFunctor, FailureNoNativeCrossLaneExclusionTheorem, FailureAlphaStillSealed},
	}

	puncture := SourceRoute{
		Name:               "puncture/complement route",
		Premise:            "e_+ tensor P_1 puncture splits visible exposed complement Pi_top from enclosed active complement H_R^min",
		ExposureTarget:     "Pi_top=e_+ tensor P_3",
		EnclosureTarget:    "H_R^min=(C_R^2 tensor W) minus (e_+ tensor P_1)",
		CrossLaneExclusion: "candidate by exposure/complement versus enclosed puncture-complement type",
		NativeFunctor:      false,
		Strength:           "strongest internal source candidate",
		Supports:           []string{SupportPunctureComplementStrongestRoute, SupportExposureVisibleComplementPiTop, SupportEnclosurePuncturedActiveDomain, SupportCrossLaneExclusionCandidateByType},
		Failures:           []string{FailurePunctureComplementNotFunctor, FailureNoNativeTargetSelectionFunctor, FailureNoNativeCrossLaneExclusionTheorem},
	}

	codim := SourceRoute{
		Name:               "boundary degree / support codimension route",
		Premise:            "Lambda^1 B_2 as boundary exposure face; Lambda^2 B_2 as full boundary-pair enclosure volume",
		ExposureTarget:     "visible rank-three boundary-exposed socket Pi_top",
		EnclosureTarget:    "rank-seven active punctured right domain H_R^min",
		CrossLaneExclusion: "candidate by exposure degree versus enclosure degree",
		NativeFunctor:      false,
		Strength:           "plausible type candidate, weaker than puncture/complement route",
		Supports:           []string{SupportDegreeCodimRouteCandidate},
		Failures:           []string{FailureCodimRouteNotFunctor, FailureNoNativeTargetSelectionFunctor},
	}

	chamber := SourceRoute{
		Name:               "trace-normalization chamber route",
		Premise:            "degree one reads in H10=H_R^ambient plus B_2; degree two reads in H72=Lambda^4 V_8 plus B_2",
		ExposureTarget:     "Pi_top in H10",
		EnclosureTarget:    "H_R^min normalized against H72",
		CrossLaneExclusion: "not supplied by chamber typing alone",
		NativeFunctor:      false,
		Strength:           "denominator/chamber typing only",
		Supports:           []string{SupportH10H72TypedResponseChambers},
		Failures:           []string{FailureChamberRouteNotFunctor, FailureResponseChamberTypingNotTheorem},
	}

	candidate := AlphaReconstruction{
		LinearContribution:    linear,
		QuadraticContribution: quadratic,
		ReconstructedAlpha:    alpha,
		ShapeCoherent:         true,
		TargetSelectionNative: false,
		Supports:              []string{SupportAlphaSealStillCoherent, SupportR3WoundSharpened},
		Failures:              []string{FailureNoNativeTargetSelectionFunctor, FailureAlphaStillSealed, FailureNoNativeAlphaSource},
	}

	r3 := R3Assessment{
		ConditionalTraceProxyMature: true,
		TargetSelectionFunctor:      false,
		AlphaNative:                 false,
		SocketMagnitudeNative:       false,
		SectorTraceMagnitudeNative:  false,
		EligibleForR3:               false,
		EligibleForR4:               false,
		Supports:                    []string{SupportConditionalTraceProxyPlateau, SupportR3WoundSharpened},
		Failures:                    []string{FailureNoNativeTargetSelectionFunctor, FailureNoNativeSocketMagnitudeSource, FailureNoSectorTraceMagnitude, FailureNoNativeR3, FailureNotR4},
	}

	impact := Impact{
		Classification:         Classification,
		Status:                 R2Status,
		StrongestRoute:         "puncture/complement route",
		PunctureRouteStrongest: true,
		CodimRouteCandidate:    true,
		ChamberRouteCandidate:  true,
		CanUpdateNEff:          false,
		CanUpdateCYukawa:       false,
		CanUpdateCHiggs:        false,
		CanPromoteToR3:         false,
		CanPromoteToR4:         false,
	}

	firewalls := Firewalls{
		Enforced:                        true,
		NoNativeTargetSelectionFunctor:  true,
		NoNativeLambda1ToPiTop:          true,
		NoNativeLambda2ToHRMin:          true,
		NoNativeCrossLaneExclusion:      true,
		PunctureComplementNotFunctor:    true,
		CodimRouteNotFunctor:            true,
		ChamberRouteNotFunctor:          true,
		ResponseChamberTypingNotTheorem: true,
		AlphaStillSealed:                true,
		NoNativeAlphaSource:             true,
		NoNativeSocketMagnitudeSource:   true,
		NoSectorTraceMagnitude:          true,
		ConditionalNotR3:                true,
		NoNativeR3:                      true,
		NoOfficialNEffUpdate:            true,
		NoCYukawaCHiggsUpdate:           true,
		NoNumericalYukawa:               true,
		NotPhysicalYukawaSpectrum:       true,
		NotR4:                           true,
		Verdict:                         StatusFirewallVerdict,
	}

	return Audit{
		ID:            AuditID,
		Ledger:        ledger,
		Wound:         wound,
		PunctureRoute: puncture,
		CodimRoute:    codim,
		ChamberRoute:  chamber,
		Candidate:     candidate,
		R3:            r3,
		Impact:        impact,
		Firewalls:     firewalls,
		Truth:         "Gate 875 searches for native sources of the remaining boundary exterior target-selection theorem. The puncture/complement route is the strongest internal candidate, but it is not a certified functor and cannot promote the conditional trace proxy to R3.",
		Final:         "VERDICT: PUNCTURE_COMPLEMENT_STRONGEST_TARGET_SELECTION_SOURCE_CANDIDATE_BUT_NO_NATIVE_BOUNDARY_EXTERIOR_TARGET_SELECTION_FUNCTOR",
	}, nil
}

func Statuses() []string {
	return []string{
		StatusGate874Inherited,
		StatusRemainingWoundInherited,
		StatusPunctureRouteAudited,
		StatusCodimRouteAudited,
		StatusChamberRouteAudited,
		StatusCrossLaneAudited,
		StatusAlphaStillReconstructs,
		StatusOfficialFreezePreserved,
		StatusNoObservedDataUsed,
		StatusFirewallVerdict,
		SupportPunctureComplementStrongestRoute,
		SupportExposureVisibleComplementPiTop,
		SupportEnclosurePuncturedActiveDomain,
		SupportDegreeCodimRouteCandidate,
		SupportH10H72TypedResponseChambers,
		SupportR3WoundSharpened,
		FailureNoNativeTargetSelectionFunctor,
		FailureNoNativeExposureToPiTopMap,
		FailureNoNativeEnclosureToHRMinMap,
		FailureNoNativeCrossLaneExclusionTheorem,
		FailurePunctureComplementNotFunctor,
		FailureCodimRouteNotFunctor,
		FailureChamberRouteNotFunctor,
		FailureAlphaStillSealed,
		FailureConditionalProxyNotR3,
		FailureNoNativeR3,
	}
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("operator_N_eff=%.16g official_N_eff=%.16g operator_C_Yukawa=%.16g official_C_Yukawa=%.16g operator_C_Higgs=%.16g official_C_Higgs=%.16g official_frozen=%t update_official=%t supports=%s failures=%s", l.OperatorNEff, l.OfficialNEff, l.OperatorCYukawa, l.OfficialCYukawa, l.OperatorCHiggs, l.OfficialCHiggs, l.OfficialFrozen, l.CanUpdateOfficial, strings.Join(l.Supports, ","), strings.Join(l.Failures, ","))
}

func FormatWound(w RemainingWound) string {
	return fmt.Sprintf("need_Lambda1_to_Pi_top=%t need_Lambda2_to_H_R_min=%t need_cross_lane_exclusions=%t/%t alpha_sealed=%t R3_blocked=%t exact_wound=%q missing=%q supports=%s failures=%s", w.NeedLambda1ToPiTop, w.NeedLambda2ToHRMin, w.NeedNoLambda1ToHRMin, w.NeedNoLambda2ToPiTop, w.AlphaSealed, w.R3Blocked, w.ExactWound, w.MissingObject, strings.Join(w.Supports, ","), strings.Join(w.Failures, ","))
}

func FormatRoute(r SourceRoute) string {
	return fmt.Sprintf("route=%q premise=%q exposure_target=%q enclosure_target=%q cross_lane=%q native_functor=%t strength=%q supports=%s failures=%s", r.Name, r.Premise, r.ExposureTarget, r.EnclosureTarget, r.CrossLaneExclusion, r.NativeFunctor, r.Strength, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}

func FormatCandidate(c AlphaReconstruction) string {
	return fmt.Sprintf("linear=%.18g quadratic=%.18g alpha=%.18g shape=%t target_native=%t supports=%s failures=%s", c.LinearContribution, c.QuadraticContribution, c.ReconstructedAlpha, c.ShapeCoherent, c.TargetSelectionNative, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}

func FormatR3(r R3Assessment) string {
	return fmt.Sprintf("conditional_proxy_mature=%t target_functor=%t alpha_native=%t socket_magnitude_native=%t trace_readout_native=%t R3=%t R4=%t supports=%s failures=%s", r.ConditionalTraceProxyMature, r.TargetSelectionFunctor, r.AlphaNative, r.SocketMagnitudeNative, r.SectorTraceMagnitudeNative, r.EligibleForR3, r.EligibleForR4, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s status=%s strongest_route=%q puncture_strongest=%t codim_candidate=%t chamber_candidate=%t update_N_eff=%t update_C_Yukawa=%t update_C_Higgs=%t promote_R3=%t promote_R4=%t", i.Classification, i.Status, i.StrongestRoute, i.PunctureRouteStrongest, i.CodimRouteCandidate, i.ChamberRouteCandidate, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CanPromoteToR3, i.CanPromoteToR4)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%t no_target_functor=%t no_L1_to_Pi_top=%t no_L2_to_HRmin=%t no_cross_lane=%t puncture_not_functor=%t codim_not_functor=%t chamber_not_functor=%t chamber_typing_not_theorem=%t alpha_sealed=%t no_alpha_source=%t no_socket_magnitude=%t no_trace_readout=%t conditional_not_R3=%t no_R3=%t no_official_update=%t no_C_updates=%t no_yukawa_values=%t not_physical_yukawa=%t not_R4=%t verdict=%s", f.Enforced, f.NoNativeTargetSelectionFunctor, f.NoNativeLambda1ToPiTop, f.NoNativeLambda2ToHRMin, f.NoNativeCrossLaneExclusion, f.PunctureComplementNotFunctor, f.CodimRouteNotFunctor, f.ChamberRouteNotFunctor, f.ResponseChamberTypingNotTheorem, f.AlphaStillSealed, f.NoNativeAlphaSource, f.NoNativeSocketMagnitudeSource, f.NoSectorTraceMagnitude, f.ConditionalNotR3, f.NoNativeR3, f.NoOfficialNEffUpdate, f.NoCYukawaCHiggsUpdate, f.NoNumericalYukawa, f.NotPhysicalYukawaSpectrum, f.NotR4, f.Verdict)
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NoNativeTargetSelectionFunctor && f.NoNativeLambda1ToPiTop && f.NoNativeLambda2ToHRMin && f.NoNativeCrossLaneExclusion && f.PunctureComplementNotFunctor && f.CodimRouteNotFunctor && f.ChamberRouteNotFunctor && f.ResponseChamberTypingNotTheorem && f.AlphaStillSealed && f.NoNativeAlphaSource && f.NoNativeSocketMagnitudeSource && f.NoSectorTraceMagnitude && f.ConditionalNotR3 && f.NoNativeR3 && f.NoOfficialNEffUpdate && f.NoCYukawaCHiggsUpdate && f.NoNumericalYukawa && f.NotPhysicalYukawaSpectrum && f.NotR4
}

func containsAll(haystack, needles []string) bool {
	seen := make(map[string]bool, len(haystack))
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

func near(a, b float64) bool { return math.Abs(a-b) < 1e-12 }
