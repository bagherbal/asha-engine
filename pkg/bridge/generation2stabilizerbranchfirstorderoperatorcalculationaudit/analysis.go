// Package generation2stabilizerbranchfirstorderoperatorcalculationaudit implements
// Gate 861: Stabilizer-Branch First-Order Operator Calculation Audit.
//
// Gate 861 follows Gate 860's scalar edge-socket operator realization.  Gate 860
// upgraded the symbolic edge support labels into operator-valued scalar socket
// maps:
//
//	Y = y_+3 |h_+><e_+| tensor I_{P_3}
//	  + y_-3 |h_-><e_-| tensor I_{P_3}
//	  + y_-1 |h_-><e_-| tensor I_{P_1},
//
// with Y_+1=0.  Gate 861 audits the stabilizer-branch first-order expression
//
//	[[D_F^sym,rho_F(a)],rho_F^op(b)] = 0,
//
// for a,b in A_F^orient = C_R plus C_H plus M_3(C).  The audit separates the
// allowed nonzero one-form commutator [D_F,rho_F(a)] from a genuine first-order
// obstruction.  It finds that the color obstruction is removed by the scalar
// identity action on P_3, while the remaining socket-character matching is still
// orientation-seal data rather than a native operator theorem.  No full unbroken
// A_F theorem, no Yukawa magnitude, no alpha_B source, and no R3/R4 promotion is
// certified.
package generation2stabilizerbranchfirstorderoperatorcalculationaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE861-STABILIZER-BRANCH-FIRST-ORDER-OPERATOR-CALCULATION-AUDIT"

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

	StatusGate860Inherited        = "PASS_GATE860_SCALAR_EDGE_SOCKET_OPERATOR_REALIZATION_INHERITED"
	StatusFirstOrderAttempted     = "PASS_OPERATOR_LEVEL_FIRST_ORDER_CALCULATION_ATTEMPTED_IN_A_F_ORIENT"
	StatusNonzeroDRhoAllowed      = "PASS_NONZERO_D_RHO_CLASSIFIED_AS_ONE_FORM_SOURCE"
	StatusColorObstructionRemoved = "PASS_COLOR_OBSTRUCTION_REMOVED_BY_EDGE_CENTRALITY"
	StatusPunctureZeroPreserved   = "PASS_PUNCTURE_EDGE_ZERO_PRESERVED"
	StatusKernelPreserved         = "PASS_LEFT_KERNEL_SINGLETON_PRESERVED"
	StatusNoObservedDataUsed      = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusLedgerFrozen            = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusFirewallVerdict         = "FIREWALL_PRESERVED_GATE861_STABILIZER_FIRST_ORDER_OPERATOR_CALCULATION_NOT_R3"

	SupportColorCentralitySolvesM3   = "CONDITIONAL_SUPPORT_COLOR_CENTRALITY_SOLVES_OPPOSITE_M3_PRESSURE"
	SupportOperatorFirstOrderIfChars = "CONDITIONAL_SUPPORT_STABILIZER_FIRST_ORDER_OPERATOR_COMPATIBILITY_GIVEN_SOCKET_CHARACTER_MATCHING"
	SupportPunctureKernelPersist     = "CONDITIONAL_SUPPORT_PUNCTURE_AND_LEFT_KERNEL_PERSIST_IN_A_F_ORIENT_BRANCH"
	SupportPostOrientationOperator   = "CONDITIONAL_SUPPORT_POST_ORIENTATION_OPERATOR_LEVEL_FINITE_TRIPLE_SEAL"
	SupportCharacterMatchNeeded      = "CONDITIONAL_SUPPORT_SOCKET_CHARACTER_MATCHING_IS_THE_REMAINING_OPERATOR_PRESSURE"

	FailureNotFullAFTheorem         = "FAILED_ROUTE_NOT_FULL_UNBROKEN_A_F_THEOREM"
	FailureSocketCharacterSeal      = "FAILED_ROUTE_SOCKET_CHARACTER_MATCH_REMAINS_ORIENTATION_SEAL"
	FailureNoNativeFiniteTriple     = "FAILED_ROUTE_NO_NATIVE_FINITE_TRIPLE_THEOREM"
	FailureNoFullUnbrokenAF         = "FAILED_ROUTE_A_F_ORIENT_NOT_FULL_A_F"
	FailureNoYukawaMagnitudes       = "FAILED_ROUTE_SYMBOLIC_Y_COEFFICIENTS_NOT_YUKAWA_MAGNITUDES"
	FailureNoNumericalYukawa        = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureNoAlphaSource            = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoTraceReadout           = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoOfficialNEffUpdate     = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate    = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoPhysicalParticleAssign = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoNeutrinoTheorem        = "FAILED_ROUTE_NO_PHYSICAL_NEUTRINO_THEOREM"
	FailureNoThreeGenerationTheorem = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
	FailureNoR3                     = "FAILED_ROUTE_R2_STABILIZER_OPERATOR_FIRST_ORDER_SEAL_NOT_R3"
	FailureNoR4                     = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoFullOperatorTheorem    = "FAILED_ROUTE_NO_FULL_UNBROKEN_OPERATOR_FIRST_ORDER_THEOREM"
	FailureNoCompleteJProof         = "FAILED_ROUTE_NO_COMPLETE_J_OPPOSITE_OPERATOR_PROOF_BEYOND_STABILIZER_SEAL"
	FailureNoBimoduleCommutantProof = "FAILED_ROUTE_NO_FULL_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
)

type Ledger struct {
	AlphaB                        float64
	OfficialNEff, OfficialCYukawa float64
	OfficialCHiggs                float64
	OfficialFrozen                bool
	AlphaNative, R3, R4           bool
}

type EdgeOperator struct {
	Name, Domain, Codomain, OperatorForm, VisibleFactor string
	Rank                                                int
	Present, CoefficientSymbolic                        bool
	ScalarOnVisibleFactor, IdentityOnVisibleFactor      bool
	ColorEdge, LeptonEdge, PunctureEdge                 bool
	NumericalValue, YukawaMagnitude                     bool
	Supports, Failures                                  []string
}

type FirstOrderTerm struct {
	Name, Expression, Sector, Verdict string
	ColorObstructionRemoved           bool
	CharacterMatchRequired            bool
	PunctureRegenerated               bool
	OperatorLevelCertified            bool
	Supports, Failures                []string
}

type FirstOrderCalculation struct {
	Algebra, Target                                   string
	Attempted, Gate860Inherited                       bool
	DRhoNonzeroAllowedOneForm                         bool
	ColorCentralityInstalled, ColorObstructionRemoved bool
	SocketCharacterMatchOperatorCertified             bool
	StabilizerOperatorCompatibilityCertified          bool
	FullUnbrokenCompatibilityCertified                bool
	PunctureEdgeZeroPreserved, LeftKernelPreserved    bool
	Terms                                             []FirstOrderTerm
	Supports, Failures                                []string
}

type DKernelLedger struct {
	DExpression                                                   string
	RankYIfActiveSocketsNonzero, RankDIfNonzero                   int
	KernelRankIfNonzero                                           int
	KernelSingleton                                               string
	PunctureSingleton                                             string
	PunctureZero, KernelPreserved, ZeroSocketBranchEnlargesKernel bool
	Supports, Failures                                            []string
}

type Impact struct {
	Classification                                                                       string
	OperatorFirstOrderAttempted, ColorCentralityInstalled, PunctureZero, KernelPreserved bool
	OperatorFirstOrderCertified, FullUnbrokenAFTheorem, NativeFiniteTripleProof          bool
	AlphaStillSealed, MagnitudesStillMissing                                             bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4     bool
}

type Firewalls struct {
	Enforced                                                                                             bool
	NotFullUnbrokenAFTheorem, SocketCharacterSeal, NoNativeFiniteTriple, AForientNotFullAF               bool
	NoYukawaMagnitudes, NoNumericalYukawa, NoAlphaSource, NoTraceReadout, NoOfficialNEffUpdate           bool
	NoCYukawaCHiggsUpdate, NoParticleAssignment, NoNeutrinoTheorem, NoThreeGenerationTheorem, NoR3, NoR4 bool
	NoFullUnbrokenOperatorTheorem, NoCompleteJProof, NoBimoduleCommutantProof                            bool
	Verdict                                                                                              string
}

type Audit struct {
	ID         string
	Ledger     Ledger
	Edges      []EdgeOperator
	FirstOrder FirstOrderCalculation
	Kernel     DKernelLedger
	Impact     Impact
	Firewalls  Firewalls
	Truth      string
	Final      string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		ID:     AuditID,
		Ledger: Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true},
		Edges: []EdgeOperator{
			{Name: "Y_+3", Domain: "e_+ tensor P_3", Codomain: "h_+ tensor P_3", OperatorForm: "y_+3 |h_+><e_+| tensor I_{P_3}", VisibleFactor: "P_3 color", Rank: P3Rank, Present: true, CoefficientSymbolic: true, ScalarOnVisibleFactor: true, IdentityOnVisibleFactor: true, ColorEdge: true, Supports: []string{SupportColorCentralitySolvesM3}, Failures: []string{FailureNoYukawaMagnitudes, FailureNoNumericalYukawa}},
			{Name: "Y_-3", Domain: "e_- tensor P_3", Codomain: "h_- tensor P_3", OperatorForm: "y_-3 |h_-><e_-| tensor I_{P_3}", VisibleFactor: "P_3 color", Rank: P3Rank, Present: true, CoefficientSymbolic: true, ScalarOnVisibleFactor: true, IdentityOnVisibleFactor: true, ColorEdge: true, Supports: []string{SupportColorCentralitySolvesM3}, Failures: []string{FailureNoYukawaMagnitudes, FailureNoNumericalYukawa}},
			{Name: "Y_-1", Domain: "e_- tensor P_1", Codomain: "h_- tensor P_1", OperatorForm: "y_-1 |h_-><e_-| tensor I_{P_1}", VisibleFactor: "P_1 lepton", Rank: P1Rank, Present: true, CoefficientSymbolic: true, ScalarOnVisibleFactor: true, IdentityOnVisibleFactor: true, LeptonEdge: true, Supports: []string{SupportPunctureKernelPersist}, Failures: []string{FailureNoYukawaMagnitudes, FailureNoNumericalYukawa}},
			{Name: "Y_+1", Domain: "e_+ tensor P_1", Codomain: "h_+ tensor P_1", OperatorForm: "0", VisibleFactor: "P_1 lepton", Rank: P1Rank, Present: false, CoefficientSymbolic: false, ScalarOnVisibleFactor: true, IdentityOnVisibleFactor: false, PunctureEdge: true, Supports: []string{StatusPunctureZeroPreserved, SupportPunctureKernelPersist}, Failures: []string{FailureNoNeutrinoTheorem, FailureNoPhysicalParticleAssign}},
		},
		FirstOrder: FirstOrderCalculation{
			Algebra:   "A_F^orient = C_R plus C_H plus M_3(C)",
			Target:    "[[D_F^sym,rho_F(a)],rho_F^op(b)] = 0 for a,b in A_F^orient",
			Attempted: true, Gate860Inherited: true, DRhoNonzeroAllowedOneForm: true, ColorCentralityInstalled: true, ColorObstructionRemoved: true,
			SocketCharacterMatchOperatorCertified: false, StabilizerOperatorCompatibilityCertified: false, FullUnbrokenCompatibilityCertified: false,
			PunctureEdgeZeroPreserved: true, LeftKernelPreserved: true,
			Terms: []FirstOrderTerm{
				{Name: "color opposite M_3 pressure", Expression: "Y_+3,Y_-3 are y I_{P_3}", Sector: "P_3", Verdict: "color obstruction removed by scalar identity action", ColorObstructionRemoved: true, Supports: []string{StatusColorObstructionRemoved, SupportColorCentralitySolvesM3}},
				{Name: "lepton color-trivial edge", Expression: "Y_-1 is y I_{P_1}", Sector: "P_1", Verdict: "M_3 action is trivial on lepton support", ColorObstructionRemoved: true, Supports: []string{SupportPunctureKernelPersist}},
				{Name: "socket character matching", Expression: "e_+ -> h_+, e_- -> h_-", Sector: "C_R/C_H", Verdict: "remaining operator pressure; currently orientation-seal matching", CharacterMatchRequired: true, OperatorLevelCertified: false, Supports: []string{SupportCharacterMatchNeeded, SupportOperatorFirstOrderIfChars}, Failures: []string{FailureSocketCharacterSeal}},
				{Name: "puncture edge", Expression: "Y_+1=0", Sector: "e_+ tensor P_1 to h_+ tensor P_1", Verdict: "zero edge preserved in stabilizer branch", PunctureRegenerated: false, Supports: []string{StatusPunctureZeroPreserved, SupportPunctureKernelPersist}},
			},
			Supports: []string{StatusGate860Inherited, StatusFirstOrderAttempted, StatusNonzeroDRhoAllowed, StatusColorObstructionRemoved, StatusPunctureZeroPreserved, StatusKernelPreserved, SupportColorCentralitySolvesM3, SupportOperatorFirstOrderIfChars, SupportPunctureKernelPersist, SupportPostOrientationOperator, SupportCharacterMatchNeeded},
			Failures: []string{FailureNotFullAFTheorem, FailureSocketCharacterSeal, FailureNoNativeFiniteTriple, FailureNoFullUnbrokenAF, FailureNoFullOperatorTheorem, FailureNoCompleteJProof, FailureNoBimoduleCommutantProof},
		},
		Kernel:    DKernelLedger{DExpression: "D_F^sym = [[0,Y^dagger],[Y,0]]", RankYIfActiveSocketsNonzero: YRankFull, RankDIfNonzero: DSymRankFull, KernelRankIfNonzero: KernelRank, KernelSingleton: "h_+ tensor P_1", PunctureSingleton: "e_+ tensor P_1", PunctureZero: true, KernelPreserved: true, ZeroSocketBranchEnlargesKernel: true, Supports: []string{StatusKernelPreserved, SupportPunctureKernelPersist}, Failures: []string{FailureNoNeutrinoTheorem, FailureNoPhysicalParticleAssign}},
		Impact:    Impact{Classification: "R2+++++_stabilizer_operator_first_order_seal", OperatorFirstOrderAttempted: true, ColorCentralityInstalled: true, PunctureZero: true, KernelPreserved: true, OperatorFirstOrderCertified: false, FullUnbrokenAFTheorem: false, NativeFiniteTripleProof: false, AlphaStillSealed: true, MagnitudesStillMissing: true},
		Firewalls: Firewalls{Enforced: true, NotFullUnbrokenAFTheorem: true, SocketCharacterSeal: true, NoNativeFiniteTriple: true, AForientNotFullAF: true, NoYukawaMagnitudes: true, NoNumericalYukawa: true, NoAlphaSource: true, NoTraceReadout: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NoParticleAssignment: true, NoNeutrinoTheorem: true, NoThreeGenerationTheorem: true, NoR3: true, NoR4: true, NoFullUnbrokenOperatorTheorem: true, NoCompleteJProof: true, NoBimoduleCommutantProof: true, Verdict: StatusFirewallVerdict},
		Truth:     "Gate 861 attempts the stabilizer-branch first-order operator calculation using the scalar edge-socket operator Y. Color-centrality removes the opposite M_3 pressure at the P_3 support, and Y_+1 remains zero.",
		Final:     "The branch is support/operator-seal compatible in the Higgs-oriented stabilizer layer only. Full unbroken A_F compatibility, native finite-triple proof, socket-character operator matching, Yukawa magnitudes, alpha_B source, R3/R4 promotion, and official ledger updates remain blocked.",
	}
	return a, a.Validate()
}

func (a Audit) Validate() error {
	err := func(msg string) error { return fmt.Errorf("%s: %s", AuditID, msg) }
	if !a.Ledger.OfficialFrozen || a.Ledger.AlphaNative || a.Ledger.R3 || a.Ledger.R4 {
		return err("ledger overpromoted")
	}
	if len(a.Edges) != 4 {
		return err("expected three active edges plus one puncture edge")
	}
	activeRank, activeCount := 0, 0
	for _, e := range a.Edges {
		if e.NumericalValue || e.YukawaMagnitude {
			return err("edge overpromoted to numerical/magnitude data")
		}
		if e.Present {
			activeCount++
			activeRank += e.Rank
			if !e.CoefficientSymbolic || !e.ScalarOnVisibleFactor || !e.IdentityOnVisibleFactor {
				return err("active edge must be symbolic scalar identity on visible support factor")
			}
		}
		if e.PunctureEdge && (e.Present || e.IdentityOnVisibleFactor || e.OperatorForm != "0") {
			return err("puncture edge must remain zero")
		}
	}
	if activeCount != 3 || activeRank != YRankFull {
		return err("active edge rank ledger inconsistent")
	}
	if !a.FirstOrder.Attempted || !a.FirstOrder.Gate860Inherited || !a.FirstOrder.DRhoNonzeroAllowedOneForm || !a.FirstOrder.ColorCentralityInstalled || !a.FirstOrder.ColorObstructionRemoved || a.FirstOrder.SocketCharacterMatchOperatorCertified || a.FirstOrder.StabilizerOperatorCompatibilityCertified || a.FirstOrder.FullUnbrokenCompatibilityCertified || !a.FirstOrder.PunctureEdgeZeroPreserved || !a.FirstOrder.LeftKernelPreserved {
		return err("first-order calculation flags inconsistent or overpromoted")
	}
	if !containsAll(a.FirstOrder.Supports, []string{StatusFirstOrderAttempted, StatusNonzeroDRhoAllowed, SupportColorCentralitySolvesM3, SupportOperatorFirstOrderIfChars}) {
		return err("first-order supports incomplete")
	}
	if !containsAll(a.FirstOrder.Failures, []string{FailureNotFullAFTheorem, FailureSocketCharacterSeal, FailureNoFullOperatorTheorem}) {
		return err("first-order firewalls incomplete")
	}
	if !a.Kernel.PunctureZero || !a.Kernel.KernelPreserved || a.Kernel.RankYIfActiveSocketsNonzero != YRankFull || a.Kernel.RankDIfNonzero != DSymRankFull || a.Kernel.KernelRankIfNonzero != KernelRank || a.Kernel.KernelSingleton != "h_+ tensor P_1" || a.Kernel.PunctureSingleton != "e_+ tensor P_1" {
		return err("kernel ledger inconsistent")
	}
	if !a.Impact.OperatorFirstOrderAttempted || !a.Impact.ColorCentralityInstalled || !a.Impact.PunctureZero || !a.Impact.KernelPreserved || a.Impact.OperatorFirstOrderCertified || a.Impact.FullUnbrokenAFTheorem || a.Impact.NativeFiniteTripleProof || !a.Impact.AlphaStillSealed || !a.Impact.MagnitudesStillMissing || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		return err("impact flags inconsistent")
	}
	if !a.Firewalls.Enforced || !a.Firewalls.NotFullUnbrokenAFTheorem || !a.Firewalls.SocketCharacterSeal || !a.Firewalls.NoNativeFiniteTriple || !a.Firewalls.AForientNotFullAF || !a.Firewalls.NoYukawaMagnitudes || !a.Firewalls.NoNumericalYukawa || !a.Firewalls.NoAlphaSource || !a.Firewalls.NoTraceReadout || !a.Firewalls.NoOfficialNEffUpdate || !a.Firewalls.NoCYukawaCHiggsUpdate || !a.Firewalls.NoParticleAssignment || !a.Firewalls.NoNeutrinoTheorem || !a.Firewalls.NoThreeGenerationTheorem || !a.Firewalls.NoR3 || !a.Firewalls.NoR4 || !a.Firewalls.NoFullUnbrokenOperatorTheorem || !a.Firewalls.NoCompleteJProof || !a.Firewalls.NoBimoduleCommutantProof || a.Firewalls.Verdict != StatusFirewallVerdict {
		return err("firewall flags inconsistent")
	}
	return nil
}

func Statuses() []string {
	return []string{StatusGate860Inherited, StatusFirstOrderAttempted, StatusNonzeroDRhoAllowed, StatusColorObstructionRemoved, StatusPunctureZeroPreserved, StatusKernelPreserved, StatusNoObservedDataUsed, StatusLedgerFrozen, StatusFirewallVerdict, SupportColorCentralitySolvesM3, SupportOperatorFirstOrderIfChars, SupportPunctureKernelPersist, SupportPostOrientationOperator, SupportCharacterMatchNeeded, FailureNotFullAFTheorem, FailureSocketCharacterSeal, FailureNoNativeFiniteTriple, FailureNoFullUnbrokenAF, FailureNoYukawaMagnitudes, FailureNoNumericalYukawa, FailureNoAlphaSource, FailureNoTraceReadout, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoPhysicalParticleAssign, FailureNoNeutrinoTheorem, FailureNoThreeGenerationTheorem, FailureNoR3, FailureNoR4, FailureNoFullOperatorTheorem, FailureNoCompleteJProof, FailureNoBimoduleCommutantProof}
}

func containsAll(haystack, needles []string) bool {
	for _, n := range needles {
		found := false
		for _, h := range haystack {
			if h == n {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("alpha_B=%.16g official_N_eff=%.16g frozen=%v alpha_native=%v R3=%v R4=%v", l.AlphaB, l.OfficialNEff, l.OfficialFrozen, l.AlphaNative, l.R3, l.R4)
}
func FormatEdges(edges []EdgeOperator) string {
	parts := []string{}
	for _, e := range edges {
		parts = append(parts, fmt.Sprintf("%s:%s->%s op=%s present=%v scalar=%v identity=%v magnitude=%v", e.Name, e.Domain, e.Codomain, e.OperatorForm, e.Present, e.ScalarOnVisibleFactor, e.IdentityOnVisibleFactor, e.YukawaMagnitude))
	}
	return strings.Join(parts, " | ")
}
func FormatFirstOrder(f FirstOrderCalculation) string {
	return fmt.Sprintf("%s for %s attempted=%v d_rho_oneform=%v color_removed=%v char_match_certified=%v stabilizer_certified=%v full_AF=%v failures=%s", f.Target, f.Algebra, f.Attempted, f.DRhoNonzeroAllowedOneForm, f.ColorObstructionRemoved, f.SocketCharacterMatchOperatorCertified, f.StabilizerOperatorCompatibilityCertified, f.FullUnbrokenCompatibilityCertified, strings.Join(f.Failures, ","))
}
func FormatTerms(ts []FirstOrderTerm) string {
	parts := []string{}
	for _, t := range ts {
		parts = append(parts, fmt.Sprintf("%s:%s sector=%s verdict=%s char_required=%v op_certified=%v", t.Name, t.Expression, t.Sector, t.Verdict, t.CharacterMatchRequired, t.OperatorLevelCertified))
	}
	return strings.Join(parts, " | ")
}
func FormatKernel(k DKernelLedger) string {
	return fmt.Sprintf("%s rankY=%d rankD=%d kernel_rank=%d kernel=%s puncture=%s puncture_zero=%v", k.DExpression, k.RankYIfActiveSocketsNonzero, k.RankDIfNonzero, k.KernelRankIfNonzero, k.KernelSingleton, k.PunctureSingleton, k.PunctureZero)
}
func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s first_order_attempted=%v color_centrality=%v puncture_zero=%v kernel=%v first_order_certified=%v alpha_sealed=%v magnitudes_missing=%v R3=%v R4=%v", i.Classification, i.OperatorFirstOrderAttempted, i.ColorCentralityInstalled, i.PunctureZero, i.KernelPreserved, i.OperatorFirstOrderCertified, i.AlphaStillSealed, i.MagnitudesStillMissing, i.CanPromoteToR3, i.CanPromoteToR4)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%v verdict=%s not_full_AF=%v char_seal=%v no_native_triple=%v no_yukawa=%v no_alpha=%v not_R3=%v", f.Enforced, f.Verdict, f.NotFullUnbrokenAFTheorem, f.SocketCharacterSeal, f.NoNativeFiniteTriple, f.NoYukawaMagnitudes, f.NoAlphaSource, f.NoR3)
}
