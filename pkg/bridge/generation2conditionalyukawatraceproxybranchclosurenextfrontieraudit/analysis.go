// Package generation2conditionalyukawatraceproxybranchclosurenextfrontieraudit implements
// Gate 881: Conditional Yukawa TraceProxy Branch Closure and Next-Frontier Audit.
//
// Gate 881 follows Gate 880's BoundaryAlpha IncidenceFlag Seal freeze. It does
// not attempt another alpha proof or a Yukawa-spectrum readout. It closes the
// current conditional trace-proxy branch, records the mature R2+++++ status, and
// selects the next lawful frontier: R3 preparation under the alpha seal, rather
// than looping on the missing native incidence functor or jumping to individual
// physical Yukawa values.
package generation2conditionalyukawatraceproxybranchclosurenextfrontieraudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE881-CONDITIONAL-YUKAWA-TRACEPROXY-BRANCH-CLOSURE-NEXT-FRONTIER-AUDIT"

	AlphaB    = 0.0003878958469680527
	SBoundary = 0.0012924448188162962

	F1OverF0 = 3
	F2OverF0 = 7
	H10Dim   = 10
	H72Dim   = 72

	OperatorNEffDiagnostic    = 3.002327375081808
	OfficialNEffFrozen        = 3.0023273474722147
	OperatorCYukawaDiagnostic = 0.9992248096922658
	OfficialCYukawaFrozen     = 0.9992248188812008
	OperatorCHiggsDiagnostic  = 1.037220510866514
	OfficialCHiggsFrozen      = 1.0372205204048603

	SealName       = "BOUNDARY_ALPHA_INCIDENCE_FLAG_SEAL"
	FullSealName   = "BOUNDARY_REDUCED_EXTERIOR_INCIDENCE_FLAG_ALPHA_SEAL"
	ClosureName    = "CONDITIONAL_YUKAWA_TRACE_PROXY_BRANCH_CLOSED"
	Classification = "CONDITIONAL_YUKAWA_TRACE_PROXY_BRANCH_CLOSURE"
	R2Status       = "R2+++++_CONDITIONAL_YUKAWA_TRACE_PROXY_BRANCH_CLOSED_NOT_R3"

	FrontierNativeIncidenceFunctor = "FRONTIER_A_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR_SEARCH"
	FrontierR3PreparationUnderSeal = "FRONTIER_B_R3_PREPARATION_UNDER_BOUNDARY_ALPHA_SEAL"
	FrontierIndividualYukawaBranch = "FRONTIER_C_INDIVIDUAL_YUKAWA_SPECTRUM_BRANCH"
	RecommendedNextFrontier        = FrontierR3PreparationUnderSeal

	MissingNativeTheoremName = "BoundaryExteriorIncidenceFlagFunctor"

	StatusGate880Inherited          = "PASS_GATE880_BOUNDARY_ALPHA_INCIDENCE_FLAG_SEAL_FREEZE_INHERITED"
	StatusBranchClosed              = "PASS_CONDITIONAL_TRACE_PROXY_BRANCH_CLOSED"
	StatusConditionalLedgerRecorded = "PASS_CONDITIONAL_OPERATOR_LEDGER_RECORDED"
	StatusOfficialFreezePreserved   = "PASS_OFFICIAL_LEDGER_FREEZE_PRESERVED"
	StatusNativeWallFiled           = "PASS_NATIVE_R3_WALL_EXACTLY_FILED"
	StatusFrontiersAudited          = "PASS_NEXT_FRONTIERS_AUDITED"
	StatusFrontierBSelected         = "PASS_FRONTIER_B_SELECTED_AS_NEXT_LAWFUL_PATH"
	StatusNoNewProofAttempt         = "PASS_NO_NEW_ALPHA_PROOF_ATTEMPT_IN_GATE881"
	StatusNoObservedDataUsed        = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusFirewallVerdict           = "FIREWALL_PRESERVED_GATE881_BRANCH_CLOSURE_NOT_R3_NOT_R4"

	SupportConditionalTraceProxyAchieved = "CONDITIONAL_SUPPORT_CONDITIONAL_YUKAWA_TRACE_PROXY_ACHIEVED"
	SupportMatureR2Closure               = "CONDITIONAL_SUPPORT_MATURE_R2_PLUS_PLUS_PLUS_PLUS_PLUS_BRANCH_CLOSED"
	SupportAlphaSealChainCoherent        = "CONDITIONAL_SUPPORT_BOUNDARY_ALPHA_INCIDENCE_FLAG_SEAL_CHAIN_COHERENT"
	SupportOperatorLedgerDiagnostic      = "CONDITIONAL_SUPPORT_OPERATOR_N_EFF_AND_C_YUKAWA_DIAGNOSTIC_ONLY"
	SupportNativeWallIsIncidenceFunctor  = "CONDITIONAL_SUPPORT_NATIVE_R3_WALL_IS_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	SupportFrontierAHardest              = "CONDITIONAL_SUPPORT_FRONTIER_A_IS_DIRECT_BUT_LIKELY_LOOPING_WITHOUT_NEW_OBJECT"
	SupportFrontierBSafest               = "CONDITIONAL_SUPPORT_FRONTIER_B_IS_RECOMMENDED_NEXT_FRONTIER"
	SupportFrontierCTooEarly             = "CONDITIONAL_SUPPORT_FRONTIER_C_IS_TOO_EARLY_FOR_PHYSICAL_YUKAWA_VALUES"

	FailureNoNativeIncidenceFunctor   = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_EXTERIOR_INCIDENCE_FLAG_FUNCTOR"
	FailureNoNativeCrossLaneExclusion = "FAILED_ROUTE_NO_NATIVE_CROSS_LANE_EXCLUSION_THEOREM"
	FailureAlphaStillSealed           = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureConditionalProxyNotR3      = "FAILED_ROUTE_CONDITIONAL_TRACE_PROXY_NOT_R3"
	FailureNoNativeR3                 = "FAILED_ROUTE_NO_NATIVE_R3_SECTOR_TRACE_LEDGER"
	FailureNoNativeSectorTraceLedger  = "FAILED_ROUTE_NO_NATIVE_SECTOR_TRACE_LEDGER"
	FailureNoIndividualYukawaSpectrum = "FAILED_ROUTE_NO_INDIVIDUAL_PHYSICAL_YUKAWA_SPECTRUM"
	FailureNoGenerationFlavorBranch   = "FAILED_ROUTE_NO_GENERATION_FLAVOR_SPLITTING_THEOREM"
	FailureNoOfficialNEffUpdate       = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate      = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNativeYukawaOperator     = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoR4                       = "FAILED_ROUTE_NO_R4_NATIVE_YUKAWA_THEOREM"
)

type ConditionalLedger struct {
	OperatorNEff, OfficialNEff       float64
	OperatorCYukawa, OfficialCYukawa float64
	OperatorCHiggs, OfficialCHiggs   float64
	OfficialFrozen, DiagnosticOnly   bool
	CanUpdate                        bool
	Supports, Failures               []string
}

type AlphaSeal struct {
	Name, FullName                            string
	Alpha                                     float64
	LinearContribution, QuadraticContribution float64
	RankF1OverF0, RankF2OverF0                int
	ReducedExteriorResponse                   bool
	IncidenceFlagSelector                     bool
	NativeFunctor                             bool
	Supports, Failures                        []string
}

type ConditionalChain struct {
	ReducedExteriorToAlpha  bool
	AlphaToSocketMagnitudes bool
	SocketMagnitudesToYDagY bool
	YDagYToHAgg             bool
	HAggToNEff              bool
	NEffToCYukawaProxy      bool
	CoherentGivenSeal       bool
	Supports, Failures      []string
}

type NativeWall struct {
	Name                         string
	DegreeOneRule, DegreeTwoRule string
	CrossLaneExclusion           string
	Native, RequiredForR3        bool
	BlocksOfficialLedger         bool
	Supports, Failures           []string
}

type Frontier struct {
	ID, Name, Description string
	DirectToNativeR3      bool
	RiskyWithoutNewObject bool
	Recommended           bool
	BlockedBy             []string
	Supports, Failures    []string
}

type FrontierDecision struct {
	Frontiers               []Frontier
	RecommendedFrontier     string
	Why                     string
	AvoidAlphaLoop          bool
	AvoidPhysicalYukawaJump bool
	Supports, Failures      []string
}

type Closure struct {
	Name, Status                       string
	BranchClosed                       bool
	ConditionalProxyMature             bool
	EligibleForNativeR3, EligibleForR4 bool
	NextFrontier                       string
	Supports, Failures                 []string
}

type Firewalls struct {
	Enforced                   bool
	AlphaStillSealed           bool
	NoNativeIncidenceFunctor   bool
	NoNativeCrossLaneExclusion bool
	ConditionalProxyNotR3      bool
	NoNativeR3                 bool
	NoIndividualYukawaSpectrum bool
	NoOfficialNEffUpdate       bool
	NoCYukawaCHiggsUpdate      bool
	NoNativeYukawaOperator     bool
	NoR4                       bool
	Verdict                    string
}

type Audit struct {
	ID        string
	Ledger    ConditionalLedger
	Alpha     AlphaSeal
	Chain     ConditionalChain
	Wall      NativeWall
	Decision  FrontierDecision
	Closure   Closure
	Firewalls Firewalls
	Truth     string
	Final     string
}

func BuildDefault() (Audit, error) {
	linear := float64(F1OverF0) / float64(H10Dim) * SBoundary
	quadratic := float64(F2OverF0) / float64(H72Dim) * SBoundary * SBoundary
	alpha := linear + quadratic
	if !near(alpha, AlphaB) {
		return Audit{}, fmt.Errorf("alpha reconstruction drift: got %.18g want %.18g", alpha, AlphaB)
	}
	if near(OperatorNEffDiagnostic, OfficialNEffFrozen) || near(OperatorCYukawaDiagnostic, OfficialCYukawaFrozen) || near(OperatorCHiggsDiagnostic, OfficialCHiggsFrozen) {
		return Audit{}, fmt.Errorf("operator and official diagnostic ledgers collapsed")
	}

	ledger := ConditionalLedger{
		OperatorNEff: OperatorNEffDiagnostic, OfficialNEff: OfficialNEffFrozen,
		OperatorCYukawa: OperatorCYukawaDiagnostic, OfficialCYukawa: OfficialCYukawaFrozen,
		OperatorCHiggs: OperatorCHiggsDiagnostic, OfficialCHiggs: OfficialCHiggsFrozen,
		OfficialFrozen: true, DiagnosticOnly: true, CanUpdate: false,
		Supports: []string{SupportConditionalTraceProxyAchieved, SupportOperatorLedgerDiagnostic},
		Failures: []string{FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureConditionalProxyNotR3},
	}
	alphaSeal := AlphaSeal{
		Name: SealName, FullName: FullSealName, Alpha: alpha, LinearContribution: linear, QuadraticContribution: quadratic,
		RankF1OverF0: F1OverF0, RankF2OverF0: F2OverF0, ReducedExteriorResponse: true, IncidenceFlagSelector: true, NativeFunctor: false,
		Supports: []string{SupportAlphaSealChainCoherent, SupportNativeWallIsIncidenceFunctor},
		Failures: []string{FailureAlphaStillSealed, FailureNoNativeIncidenceFunctor, FailureNoNativeCrossLaneExclusion},
	}
	chain := ConditionalChain{
		ReducedExteriorToAlpha: true, AlphaToSocketMagnitudes: true, SocketMagnitudesToYDagY: true, YDagYToHAgg: true, HAggToNEff: true, NEffToCYukawaProxy: true, CoherentGivenSeal: true,
		Supports: []string{SupportConditionalTraceProxyAchieved, SupportMatureR2Closure, SupportAlphaSealChainCoherent},
		Failures: []string{FailureAlphaStillSealed, FailureConditionalProxyNotR3, FailureNoNativeR3},
	}
	wall := NativeWall{
		Name:               MissingNativeTheoremName,
		DegreeOneRule:      "I_B(1)=F_1/F_0=Pi_top",
		DegreeTwoRule:      "I_B(2)=F_2/F_0=H_R^min",
		CrossLaneExclusion: "I_B(1)!=F_2/F_0 and I_B(2)!=F_1/F_0",
		Native:             false, RequiredForR3: true, BlocksOfficialLedger: true,
		Supports: []string{SupportNativeWallIsIncidenceFunctor},
		Failures: []string{FailureNoNativeIncidenceFunctor, FailureNoNativeCrossLaneExclusion, FailureAlphaStillSealed},
	}
	frontiers := []Frontier{
		{ID: FrontierNativeIncidenceFunctor, Name: "Native alpha theorem", Description: "prove BoundaryExteriorIncidenceFlagFunctor and cross-lane exclusion", DirectToNativeR3: true, RiskyWithoutNewObject: true, Recommended: false, BlockedBy: []string{FailureNoNativeIncidenceFunctor, FailureNoNativeCrossLaneExclusion}, Supports: []string{SupportFrontierAHardest}, Failures: []string{FailureNoNativeIncidenceFunctor}},
		{ID: FrontierR3PreparationUnderSeal, Name: "R3 preparation under seal", Description: "audit sector decomposition, generation carrier, trace-ledger requirements, and flavor firewalls given the frozen alpha seal", DirectToNativeR3: false, RiskyWithoutNewObject: false, Recommended: true, BlockedBy: []string{FailureAlphaStillSealed}, Supports: []string{SupportFrontierBSafest}, Failures: []string{FailureConditionalProxyNotR3}},
		{ID: FrontierIndividualYukawaBranch, Name: "Individual Yukawa spectrum branch", Description: "split aggregate trace proxy into individual physical values", DirectToNativeR3: false, RiskyWithoutNewObject: true, Recommended: false, BlockedBy: []string{FailureNoNativeSectorTraceLedger, FailureNoGenerationFlavorBranch, FailureNoIndividualYukawaSpectrum}, Supports: []string{SupportFrontierCTooEarly}, Failures: []string{FailureNoIndividualYukawaSpectrum, FailureNoGenerationFlavorBranch}},
	}
	decision := FrontierDecision{
		Frontiers: frontiers, RecommendedFrontier: RecommendedNextFrontier,
		Why:            "Frontier B avoids looping on the absent incidence functor and avoids premature physical Yukawa splitting while still moving toward R3 requirements.",
		AvoidAlphaLoop: true, AvoidPhysicalYukawaJump: true,
		Supports: []string{SupportFrontierAHardest, SupportFrontierBSafest, SupportFrontierCTooEarly},
		Failures: []string{FailureNoNativeIncidenceFunctor, FailureNoIndividualYukawaSpectrum},
	}
	closure := Closure{
		Name: ClosureName, Status: R2Status, BranchClosed: true, ConditionalProxyMature: true, EligibleForNativeR3: false, EligibleForR4: false, NextFrontier: RecommendedNextFrontier,
		Supports: []string{SupportConditionalTraceProxyAchieved, SupportMatureR2Closure, SupportFrontierBSafest},
		Failures: []string{FailureConditionalProxyNotR3, FailureNoNativeR3, FailureNoNativeYukawaOperator, FailureNoR4},
	}
	firewalls := Firewalls{Enforced: true, AlphaStillSealed: true, NoNativeIncidenceFunctor: true, NoNativeCrossLaneExclusion: true, ConditionalProxyNotR3: true, NoNativeR3: true, NoIndividualYukawaSpectrum: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NoNativeYukawaOperator: true, NoR4: true, Verdict: StatusFirewallVerdict}

	return Audit{ID: AuditID, Ledger: ledger, Alpha: alphaSeal, Chain: chain, Wall: wall, Decision: decision, Closure: closure, Firewalls: firewalls, Truth: "Gate 881 closes the mature conditional Yukawa trace-proxy branch without promoting it to native R3: the only native wall is the missing BoundaryExteriorIncidenceFlagFunctor and cross-lane exclusion theorem.", Final: "Recommended next frontier is R3 preparation under the BoundaryAlpha incidence-flag seal, not another alpha proof loop and not individual physical Yukawa splitting."}, nil
}

func FormatLedger(l ConditionalLedger) string {
	return fmt.Sprintf("ledger(operator_N_eff=%.16g official_N_eff=%.16g operator_CYukawa=%.16g official_CYukawa=%.16g operator_CHiggs=%.16g official_CHiggs=%.16g frozen=%t diagnostic_only=%t can_update=%t supports=%s failures=%s)", l.OperatorNEff, l.OfficialNEff, l.OperatorCYukawa, l.OfficialCYukawa, l.OperatorCHiggs, l.OfficialCHiggs, l.OfficialFrozen, l.DiagnosticOnly, l.CanUpdate, strings.Join(l.Supports, ","), strings.Join(l.Failures, ","))
}

func FormatAlpha(a AlphaSeal) string {
	return fmt.Sprintf("alpha(name=%s full=%s alpha=%.18g linear=%.18g quadratic=%.18g F1/F0=%d F2/F0=%d reduced=%t incidence=%t native=%t supports=%s failures=%s)", a.Name, a.FullName, a.Alpha, a.LinearContribution, a.QuadraticContribution, a.RankF1OverF0, a.RankF2OverF0, a.ReducedExteriorResponse, a.IncidenceFlagSelector, a.NativeFunctor, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatChain(c ConditionalChain) string {
	return fmt.Sprintf("chain(reduced_to_alpha=%t alpha_to_socket=%t socket_to_YdagY=%t YdagY_to_Hagg=%t Hagg_to_Neff=%t Neff_to_CYukawa=%t coherent_given_seal=%t supports=%s failures=%s)", c.ReducedExteriorToAlpha, c.AlphaToSocketMagnitudes, c.SocketMagnitudesToYDagY, c.YDagYToHAgg, c.HAggToNEff, c.NEffToCYukawaProxy, c.CoherentGivenSeal, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}

func FormatWall(w NativeWall) string {
	return fmt.Sprintf("native_wall(name=%s deg1=%s deg2=%s cross_lane=%s native=%t required_r3=%t blocks_ledger=%t supports=%s failures=%s)", w.Name, w.DegreeOneRule, w.DegreeTwoRule, w.CrossLaneExclusion, w.Native, w.RequiredForR3, w.BlocksOfficialLedger, strings.Join(w.Supports, ","), strings.Join(w.Failures, ","))
}

func FormatDecision(d FrontierDecision) string {
	parts := []string{}
	for _, f := range d.Frontiers {
		parts = append(parts, fmt.Sprintf("%s(recommended=%t risky=%t blocked_by=%s)", f.ID, f.Recommended, f.RiskyWithoutNewObject, strings.Join(f.BlockedBy, "+")))
	}
	return fmt.Sprintf("decision(recommended=%s avoid_alpha_loop=%t avoid_yukawa_jump=%t frontiers=%s why=%q supports=%s failures=%s)", d.RecommendedFrontier, d.AvoidAlphaLoop, d.AvoidPhysicalYukawaJump, strings.Join(parts, ";"), d.Why, strings.Join(d.Supports, ","), strings.Join(d.Failures, ","))
}

func FormatClosure(c Closure) string {
	return fmt.Sprintf("closure(name=%s status=%s closed=%t mature=%t r3=%t r4=%t next=%s supports=%s failures=%s)", c.Name, c.Status, c.BranchClosed, c.ConditionalProxyMature, c.EligibleForNativeR3, c.EligibleForR4, c.NextFrontier, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("firewalls(enforced=%t alpha_sealed=%t no_incidence=%t no_cross_lane=%t conditional_not_r3=%t no_r3=%t no_individual_yukawa=%t no_neff_update=%t no_c_update=%t no_yukawa=%t no_r4=%t verdict=%s)", f.Enforced, f.AlphaStillSealed, f.NoNativeIncidenceFunctor, f.NoNativeCrossLaneExclusion, f.ConditionalProxyNotR3, f.NoNativeR3, f.NoIndividualYukawaSpectrum, f.NoOfficialNEffUpdate, f.NoCYukawaCHiggsUpdate, f.NoNativeYukawaOperator, f.NoR4, f.Verdict)
}

func Statuses() []string {
	return []string{
		StatusGate880Inherited,
		StatusBranchClosed,
		StatusConditionalLedgerRecorded,
		StatusOfficialFreezePreserved,
		StatusNativeWallFiled,
		StatusFrontiersAudited,
		StatusFrontierBSelected,
		StatusNoNewProofAttempt,
		StatusNoObservedDataUsed,
		StatusFirewallVerdict,
		SupportConditionalTraceProxyAchieved,
		SupportMatureR2Closure,
		SupportAlphaSealChainCoherent,
		SupportOperatorLedgerDiagnostic,
		SupportNativeWallIsIncidenceFunctor,
		SupportFrontierAHardest,
		SupportFrontierBSafest,
		SupportFrontierCTooEarly,
		FailureNoNativeIncidenceFunctor,
		FailureNoNativeCrossLaneExclusion,
		FailureAlphaStillSealed,
		FailureConditionalProxyNotR3,
		FailureNoNativeR3,
		FailureNoNativeSectorTraceLedger,
		FailureNoIndividualYukawaSpectrum,
		FailureNoGenerationFlavorBranch,
		FailureNoOfficialNEffUpdate,
		FailureNoCYukawaCHiggsUpdate,
		FailureNoNativeYukawaOperator,
		FailureNoR4,
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
	return f.Enforced && f.AlphaStillSealed && f.NoNativeIncidenceFunctor && f.NoNativeCrossLaneExclusion && f.ConditionalProxyNotR3 && f.NoNativeR3 && f.NoIndividualYukawaSpectrum && f.NoOfficialNEffUpdate && f.NoCYukawaCHiggsUpdate && f.NoNativeYukawaOperator && f.NoR4 && f.Verdict == StatusFirewallVerdict
}
