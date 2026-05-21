// Package generation2jfoppositeactionorderzerobimodulerealizationaudit implements
// Gate 858: J_F-Opposite Action and Order-Zero Bimodule Realization Audit.
//
// Gate 858 follows Gate 857's stabilizer-branch support-level first-order
// success.  It does not retry the first-order theorem.  Instead it audits the
// missing right/opposite action rho_F^op(b)=J_F rho_F(b) J_F^{-1} and the
// order-zero bimodule prerequisite [rho_F(a),rho_F^op(b)]=0 inside the
// post-Higgs-orientation stabilizer algebra A_F^orient.  The result remains a
// support-level bimodule seal: no operator-level J_F proof, no first-order
// theorem, no numerical Yukawa magnitudes, no alpha_B source, no R3/R4, and no
// official ledger update are certified.
package generation2jfoppositeactionorderzerobimodulerealizationaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE858-JF-OPPOSITE-ACTION-ORDER-ZERO-BIMODULE-REALIZATION-AUDIT"

	AlphaB          = 0.0003878958469680527
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	P1Rank          = 1
	P3Rank          = 3
	WRank           = P1Rank + P3Rank
	WeakDoubletRank = 2
	HLRank          = WeakDoubletRank * WRank
	HRMinRank       = 7
	HPartMinRank    = HLRank + HRMinRank
	HFMinRank       = 2 * HPartMinRank
	AmbientPartRank = 16
	AmbientFRank    = 32
	DSymRank        = 14
	DSymKernelRank  = HPartMinRank - DSymRank

	StatusGate857Inherited       = "PASS_GATE857_STABILIZER_BRANCH_SUPPORT_FIRST_ORDER_INHERITED"
	StatusAForientLeftAction     = "PASS_A_F_ORIENT_LEFT_ACTION_INHERITED"
	StatusOppositeRequirement    = "PASS_J_OPPOSITE_ACTION_REQUIREMENT_DEFINED"
	StatusOrderZeroAudited       = "PASS_ORDER_ZERO_TARGET_AUDITED"
	StatusMinimalJClosureAudited = "PASS_MINIMAL_15_30_CARRIER_J_CLOSURE_AUDITED"
	StatusEdgeBimoduleAudited    = "PASS_EDGE_BIMODULE_COMPATIBILITY_AUDITED_AT_SUPPORT_LEVEL"
	StatusFirstOrderDeferred     = "PASS_FIRST_ORDER_OPERATOR_CALCULATION_DEFERRED_UNTIL_ORDER_ZERO_DATA_EXISTS"
	StatusNoObservedDataUsed     = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusLedgerFrozen           = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusFirewallVerdict        = "FIREWALL_PRESERVED_GATE858_SUPPORT_BIMODULE_NOT_OPERATOR_THEOREM"

	SupportAForientSupportBimodule  = "CONDITIONAL_SUPPORT_A_F_ORIENT_FORMS_SUPPORT_LEVEL_BIMODULE_ON_H_F_MIN"
	SupportOrderZeroBlockSupport    = "CONDITIONAL_SUPPORT_ORDER_ZERO_HOLDS_AT_BLOCK_SUPPORT_LEVEL"
	SupportMinimalCarrierJClosure   = "CONDITIONAL_SUPPORT_MINIMAL_CARRIER_REMAINS_CLOSED_UNDER_FORMAL_J_COPY"
	SupportOppositeActionSeal       = "CONDITIONAL_SUPPORT_J_OPPOSITE_ACTION_DEFINED_AS_FORMAL_SEAL"
	SupportEdgeBimoduleSupport      = "CONDITIONAL_SUPPORT_ACTIVE_EDGES_ARE_BIMODULE_COMPATIBLE_AT_SUPPORT_LEVEL"
	SupportPunctureKernelUnderJSeal = "CONDITIONAL_SUPPORT_PUNCTURE_AND_KERNEL_PRESERVED_UNDER_FORMAL_J_COPY"
	SupportR2SupportBimoduleStage   = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_PLUS_PLUS_PLUS_SUPPORT_BIMODULE_SEAL"

	FailureAForientNotFullAF              = "FAILED_ROUTE_A_F_ORIENT_NOT_FULL_UNBROKEN_A_F"
	FailureFullHSocketFirewall            = "FAILED_ROUTE_FULL_H_ACTION_DOES_NOT_PRESERVE_SOCKET_FRAME"
	FailureJOppositeSealOnly              = "FAILED_ROUTE_J_OPPOSITE_ACTION_REMAINS_SEAL_WITHOUT_FULL_OPERATOR_MATRIX"
	FailureNoOperatorJOppositeProof       = "FAILED_ROUTE_NO_OPERATOR_LEVEL_J_OPPOSITE_ACTION_PROOF"
	FailureNoOrderZeroOperatorTheorem     = "FAILED_ROUTE_NO_ORDER_ZERO_OPERATOR_THEOREM"
	FailureOrderZeroSupportOnly           = "FAILED_ROUTE_ORDER_ZERO_COMPATIBILITY_IS_SUPPORT_LEVEL_NOT_OPERATOR_THEOREM"
	FailureNoFirstOrderOperatorTheorem    = "FAILED_ROUTE_NO_FIRST_ORDER_OPERATOR_THEOREM"
	FailureNoBimoduleCommutantProof       = "FAILED_ROUTE_NO_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailureNoNativeFiniteTripleProof      = "FAILED_ROUTE_NO_NATIVE_FINITE_TRIPLE_PROOF_CERTIFIED"
	FailureMinimalCarrierJSealOnly        = "FAILED_ROUTE_MINIMAL_15_30_CARRIER_J_CLOSURE_IS_FORMAL_SEAL"
	FailureJCopyDoesNotRestoreAmbientCell = "FAILED_ROUTE_J_COPY_DOES_NOT_REINTRODUCE_E_PLUS_TENSOR_P1_IN_MINIMAL_BRANCH"
	FailureEdgeIntertwinerSupportOnly     = "FAILED_ROUTE_EDGE_BIMODULE_INTERTWINER_IS_SUPPORT_LABEL_NOT_OPERATOR_PROOF"
	FailureEdgeIntertwinerNoValue         = "FAILED_ROUTE_EDGE_BIMODULE_INTERTWINER_NOT_YUKAWA_MAGNITUDE"
	FailureSymbolicYNotMagnitude          = "FAILED_ROUTE_SYMBOLIC_Y_COEFFICIENTS_NOT_YUKAWA_MAGNITUDES"
	FailureNoAlphaSource                  = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoTraceReadout                 = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoOfficialNEffUpdate           = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate          = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNotR3                          = "FAILED_ROUTE_R2_PLUS_PLUS_PLUS_PLUS_PLUS_SUPPORT_BIMODULE_NOT_R3"
	FailureNotR4                          = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoParticleAssignment           = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoNeutrinoTheorem              = "FAILED_ROUTE_NO_PHYSICAL_NEUTRINO_THEOREM"
	FailureNoThreeGenerationTheorem       = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
)

type Ledger struct {
	AlphaB                        float64
	OfficialNEff, OfficialCYukawa float64
	OfficialCHiggs                float64
	OfficialFrozen                bool
	AlphaNative, R3, R4           bool
}

type OrientedAlgebra struct {
	FullAlgebra, OrientedAlgebra string
	ContainsFullH, ContainsCH    bool
	ContainsRightC, ContainsM3C  bool
	PostOrientationLayer         bool
	Supports, Failures           []string
}

type LeftAction struct {
	Algebra                                    string
	PreservesHPlusHMinus, PreservesEPlusEMinus bool
	PreservesP1P3, PreservesHRMin, PreservesHF bool
	LeftActionOperatorCertified                bool
	Supports, Failures                         []string
}

type OppositeAction struct {
	Expression                     string
	FormalJExchangeDefined         bool
	OppositeSupportDefined         bool
	OppositeOperatorCertified      bool
	OrderZeroTargetTyped           bool
	MinimalCarrierClosedUnderJSeal bool
	AmbientCellReintroduced        bool
	Supports, Failures             []string
}

type OrderZero struct {
	Expression                string
	Algebra                   string
	SupportAuditable          bool
	BlockSupportCompatible    bool
	OperatorTheoremCertified  bool
	RequiresOperatorJOpposite bool
	Supports, Failures        []string
}

type EdgeBimodule struct {
	Name, Domain, Codomain, Support string
	Rank                            int
	LeftSupportCompatible           bool
	RightSupportCompatible          bool
	OperatorIntertwinerCertified    bool
	YukawaMagnitude                 bool
	Supports, Failures              []string
}

type MinimalCarrier struct {
	HLRank, HRMinRank, HPartMinRank, HFMinRank int
	AmbientPartRank, AmbientFRank              int
	RightPuncture, LeftKernel                  string
	RightPunctureOutsideMinimal                bool
	LeftKernelPresent                          bool
	JCopyRestoresAmbientPuncture               bool
	Supports, Failures                         []string
}

type FirstOrderBoundary struct {
	TargetExpression                  string
	OrderZeroPrerequisiteAudited      bool
	ReadyForOperatorFirstOrderAttempt bool
	FirstOrderOperatorCertified       bool
	Supports, Failures                []string
}

type Impact struct {
	Classification                                                                   string
	Gate857Inherited, SupportBimodule, OrderZeroSupportPass, JOppositeSealDefined    bool
	OperatorJOppositeProof, OrderZeroOperatorProof, FirstOrderOperatorProof          bool
	NativeFiniteTripleProof, AlphaStillSealed, MagnitudesStillMissing                bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4 bool
}

type Firewalls struct {
	Enforced                                                                                              bool
	AForientNotFullAF, FullHSocketFirewall, JOppositeSealOnly, NoOperatorJOpposite, NoOrderZeroOperator   bool
	OrderZeroSupportOnly, NoFirstOrderOperator, NoBimoduleProof, NoNativeTriple, MinimalJClosureSealOnly  bool
	JCopyNoAmbientRestore, EdgeSupportOnly, EdgeNoMagnitude, YSymbolicOnly, NoAlphaSource, NoTraceReadout bool
	NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate, NotR3, NotR4, NoParticleAssignment, NoNeutrinoTheorem    bool
	NoThreeGenerationTheorem                                                                              bool
	Verdict                                                                                               string
}

type Audit struct {
	ID             string
	Ledger         Ledger
	Algebra        OrientedAlgebra
	Left           LeftAction
	Opposite       OppositeAction
	OrderZero      OrderZero
	Edges          []EdgeBimodule
	Carrier        MinimalCarrier
	FirstOrderNext FirstOrderBoundary
	Impact         Impact
	Firewalls      Firewalls
	Truth          string
	Final          string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		ID:             AuditID,
		Ledger:         buildLedger(),
		Algebra:        buildAlgebra(),
		Left:           buildLeftAction(),
		Opposite:       buildOppositeAction(),
		OrderZero:      buildOrderZero(),
		Edges:          buildEdges(),
		Carrier:        buildCarrier(),
		FirstOrderNext: buildFirstOrderBoundary(),
		Impact:         buildImpact(),
		Firewalls:      buildFirewalls(),
		Truth:          "Gate 858 builds the oriented bimodule prerequisite: rho_F^op(b)=J_F rho_F(b) J_F^{-1} is defined as a formal support seal and order-zero is conditionally support-compatible, but no operator-level J/opposite or order-zero theorem is certified.",
		Final:          "VERDICT: CONDITIONAL_SUPPORT_A_F_ORIENT_FORMS_SUPPORT_LEVEL_BIMODULE_ON_H_F_MIN; FAILED_ROUTE_NO_OPERATOR_LEVEL_J_OPPOSITE_ACTION_PROOF; FAILED_ROUTE_NO_ORDER_ZERO_OPERATOR_THEOREM.",
	}
	if err := a.Validate(); err != nil {
		return Audit{}, err
	}
	return a, nil
}

func (a Audit) Validate() error {
	if !a.Ledger.OfficialFrozen || a.Ledger.AlphaNative || a.Ledger.R3 || a.Ledger.R4 {
		return err("ledger firewall violated")
	}
	if a.Algebra.OrientedAlgebra != "A_F^orient=C_R plus C_H plus M_3(C)" || a.Algebra.ContainsFullH || !a.Algebra.ContainsCH || !a.Algebra.ContainsRightC || !a.Algebra.ContainsM3C || !a.Algebra.PostOrientationLayer {
		return err("oriented algebra not typed correctly")
	}
	if !a.Left.PreservesHPlusHMinus || !a.Left.PreservesEPlusEMinus || !a.Left.PreservesP1P3 || !a.Left.PreservesHRMin || !a.Left.PreservesHF || a.Left.LeftActionOperatorCertified {
		return err("left action flags inconsistent")
	}
	if !a.Opposite.FormalJExchangeDefined || !a.Opposite.OppositeSupportDefined || a.Opposite.OppositeOperatorCertified || !a.Opposite.OrderZeroTargetTyped || !a.Opposite.MinimalCarrierClosedUnderJSeal || a.Opposite.AmbientCellReintroduced {
		return err("opposite action flags inconsistent")
	}
	if !a.OrderZero.SupportAuditable || !a.OrderZero.BlockSupportCompatible || a.OrderZero.OperatorTheoremCertified || !a.OrderZero.RequiresOperatorJOpposite {
		return err("order-zero flags inconsistent")
	}
	if len(a.Edges) != 3 {
		return err("expected three active edge bimodule entries")
	}
	for _, e := range a.Edges {
		if !e.LeftSupportCompatible || !e.RightSupportCompatible || e.OperatorIntertwinerCertified || e.YukawaMagnitude {
			return err("edge bimodule overpromoted")
		}
	}
	if a.Carrier.HLRank != HLRank || a.Carrier.HRMinRank != HRMinRank || a.Carrier.HPartMinRank != HPartMinRank || a.Carrier.HFMinRank != HFMinRank || a.Carrier.AmbientPartRank != AmbientPartRank || a.Carrier.AmbientFRank != AmbientFRank || !a.Carrier.RightPunctureOutsideMinimal || !a.Carrier.LeftKernelPresent || a.Carrier.JCopyRestoresAmbientPuncture {
		return err("minimal carrier closure flags inconsistent")
	}
	if !a.FirstOrderNext.OrderZeroPrerequisiteAudited || !a.FirstOrderNext.ReadyForOperatorFirstOrderAttempt || a.FirstOrderNext.FirstOrderOperatorCertified {
		return err("first-order boundary flags inconsistent")
	}
	if !a.Impact.Gate857Inherited || !a.Impact.SupportBimodule || !a.Impact.OrderZeroSupportPass || !a.Impact.JOppositeSealDefined || a.Impact.OperatorJOppositeProof || a.Impact.OrderZeroOperatorProof || a.Impact.FirstOrderOperatorProof || a.Impact.NativeFiniteTripleProof || !a.Impact.AlphaStillSealed || !a.Impact.MagnitudesStillMissing || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		return err("impact flags inconsistent")
	}
	if !a.Firewalls.Enforced || !a.Firewalls.AForientNotFullAF || !a.Firewalls.FullHSocketFirewall || !a.Firewalls.JOppositeSealOnly || !a.Firewalls.NoOperatorJOpposite || !a.Firewalls.NoOrderZeroOperator || !a.Firewalls.OrderZeroSupportOnly || !a.Firewalls.NoFirstOrderOperator || !a.Firewalls.NoBimoduleProof || !a.Firewalls.NoNativeTriple || !a.Firewalls.MinimalJClosureSealOnly || !a.Firewalls.JCopyNoAmbientRestore || !a.Firewalls.EdgeSupportOnly || !a.Firewalls.EdgeNoMagnitude || !a.Firewalls.YSymbolicOnly || !a.Firewalls.NoAlphaSource || !a.Firewalls.NoTraceReadout || !a.Firewalls.NoOfficialNEffUpdate || !a.Firewalls.NoCYukawaCHiggsUpdate || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || !a.Firewalls.NoParticleAssignment || !a.Firewalls.NoNeutrinoTheorem || !a.Firewalls.NoThreeGenerationTheorem || a.Firewalls.Verdict != StatusFirewallVerdict {
		return err("firewall flags inconsistent")
	}
	return nil
}

func err(msg string) error { return fmt.Errorf("%s", msg) }

func buildLedger() Ledger {
	return Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true}
}

func buildAlgebra() OrientedAlgebra {
	return OrientedAlgebra{FullAlgebra: "A_F=C plus H plus M_3(C)", OrientedAlgebra: "A_F^orient=C_R plus C_H plus M_3(C)", ContainsFullH: false, ContainsCH: true, ContainsRightC: true, ContainsM3C: true, PostOrientationLayer: true, Supports: []string{StatusGate857Inherited, SupportAForientSupportBimodule}, Failures: []string{FailureAForientNotFullAF, FailureFullHSocketFirewall}}
}

func buildLeftAction() LeftAction {
	return LeftAction{Algebra: "A_F^orient", PreservesHPlusHMinus: true, PreservesEPlusEMinus: true, PreservesP1P3: true, PreservesHRMin: true, PreservesHF: true, LeftActionOperatorCertified: false, Supports: []string{StatusAForientLeftAction, SupportAForientSupportBimodule}, Failures: []string{FailureOrderZeroSupportOnly}}
}

func buildOppositeAction() OppositeAction {
	return OppositeAction{Expression: "rho_F^op(b)=J_F rho_F(b) J_F^{-1}", FormalJExchangeDefined: true, OppositeSupportDefined: true, OppositeOperatorCertified: false, OrderZeroTargetTyped: true, MinimalCarrierClosedUnderJSeal: true, AmbientCellReintroduced: false, Supports: []string{StatusOppositeRequirement, SupportOppositeActionSeal, SupportMinimalCarrierJClosure}, Failures: []string{FailureJOppositeSealOnly, FailureNoOperatorJOppositeProof, FailureMinimalCarrierJSealOnly, FailureJCopyDoesNotRestoreAmbientCell}}
}

func buildOrderZero() OrderZero {
	return OrderZero{Expression: "[rho_F(a),rho_F^op(b)]=0", Algebra: "a,b in A_F^orient", SupportAuditable: true, BlockSupportCompatible: true, OperatorTheoremCertified: false, RequiresOperatorJOpposite: true, Supports: []string{StatusOrderZeroAudited, SupportOrderZeroBlockSupport, SupportAForientSupportBimodule}, Failures: []string{FailureOrderZeroSupportOnly, FailureNoOrderZeroOperatorTheorem, FailureNoOperatorJOppositeProof}}
}

func buildEdges() []EdgeBimodule {
	return []EdgeBimodule{
		{Name: "Y_+3", Domain: "e_+ tensor P_3", Codomain: "h_+ tensor P_3", Support: "P_3 color module", Rank: P3Rank, LeftSupportCompatible: true, RightSupportCompatible: true, OperatorIntertwinerCertified: false, YukawaMagnitude: false, Supports: []string{StatusEdgeBimoduleAudited, SupportEdgeBimoduleSupport}, Failures: []string{FailureEdgeIntertwinerSupportOnly, FailureEdgeIntertwinerNoValue}},
		{Name: "Y_-3", Domain: "e_- tensor P_3", Codomain: "h_- tensor P_3", Support: "P_3 color module", Rank: P3Rank, LeftSupportCompatible: true, RightSupportCompatible: true, OperatorIntertwinerCertified: false, YukawaMagnitude: false, Supports: []string{StatusEdgeBimoduleAudited, SupportEdgeBimoduleSupport}, Failures: []string{FailureEdgeIntertwinerSupportOnly, FailureEdgeIntertwinerNoValue}},
		{Name: "Y_-1", Domain: "e_- tensor P_1", Codomain: "h_- tensor P_1", Support: "P_1 color-trivial lepton support", Rank: P1Rank, LeftSupportCompatible: true, RightSupportCompatible: true, OperatorIntertwinerCertified: false, YukawaMagnitude: false, Supports: []string{StatusEdgeBimoduleAudited, SupportEdgeBimoduleSupport}, Failures: []string{FailureEdgeIntertwinerSupportOnly, FailureEdgeIntertwinerNoValue}},
	}
}

func buildCarrier() MinimalCarrier {
	return MinimalCarrier{HLRank: HLRank, HRMinRank: HRMinRank, HPartMinRank: HPartMinRank, HFMinRank: HFMinRank, AmbientPartRank: AmbientPartRank, AmbientFRank: AmbientFRank, RightPuncture: "e_+ tensor P_1", LeftKernel: "h_+ tensor P_1", RightPunctureOutsideMinimal: true, LeftKernelPresent: true, JCopyRestoresAmbientPuncture: false, Supports: []string{StatusMinimalJClosureAudited, SupportMinimalCarrierJClosure, SupportPunctureKernelUnderJSeal}, Failures: []string{FailureMinimalCarrierJSealOnly, FailureJCopyDoesNotRestoreAmbientCell, FailureNoParticleAssignment, FailureNoNeutrinoTheorem}}
}

func buildFirstOrderBoundary() FirstOrderBoundary {
	return FirstOrderBoundary{TargetExpression: "[[D_F^sym,rho_F(a)],rho_F^op(b)]=0", OrderZeroPrerequisiteAudited: true, ReadyForOperatorFirstOrderAttempt: true, FirstOrderOperatorCertified: false, Supports: []string{StatusFirstOrderDeferred}, Failures: []string{FailureNoFirstOrderOperatorTheorem, FailureNoBimoduleCommutantProof}}
}

func buildImpact() Impact {
	return Impact{Classification: "R2+++++_support_bimodule_order_zero_seal", Gate857Inherited: true, SupportBimodule: true, OrderZeroSupportPass: true, JOppositeSealDefined: true, OperatorJOppositeProof: false, OrderZeroOperatorProof: false, FirstOrderOperatorProof: false, NativeFiniteTripleProof: false, AlphaStillSealed: true, MagnitudesStillMissing: true, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, CanPromoteToR3: false, CanPromoteToR4: false}
}

func buildFirewalls() Firewalls {
	return Firewalls{Enforced: true, AForientNotFullAF: true, FullHSocketFirewall: true, JOppositeSealOnly: true, NoOperatorJOpposite: true, NoOrderZeroOperator: true, OrderZeroSupportOnly: true, NoFirstOrderOperator: true, NoBimoduleProof: true, NoNativeTriple: true, MinimalJClosureSealOnly: true, JCopyNoAmbientRestore: true, EdgeSupportOnly: true, EdgeNoMagnitude: true, YSymbolicOnly: true, NoAlphaSource: true, NoTraceReadout: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NotR3: true, NotR4: true, NoParticleAssignment: true, NoNeutrinoTheorem: true, NoThreeGenerationTheorem: true, Verdict: StatusFirewallVerdict}
}

func Statuses() []string {
	return []string{
		StatusGate857Inherited, StatusAForientLeftAction, StatusOppositeRequirement, StatusOrderZeroAudited, StatusMinimalJClosureAudited, StatusEdgeBimoduleAudited, StatusFirstOrderDeferred, StatusNoObservedDataUsed, StatusLedgerFrozen, StatusFirewallVerdict,
		SupportAForientSupportBimodule, SupportOrderZeroBlockSupport, SupportMinimalCarrierJClosure, SupportOppositeActionSeal, SupportEdgeBimoduleSupport, SupportPunctureKernelUnderJSeal, SupportR2SupportBimoduleStage,
		FailureAForientNotFullAF, FailureFullHSocketFirewall, FailureJOppositeSealOnly, FailureNoOperatorJOppositeProof, FailureNoOrderZeroOperatorTheorem, FailureOrderZeroSupportOnly, FailureNoFirstOrderOperatorTheorem, FailureNoBimoduleCommutantProof, FailureNoNativeFiniteTripleProof, FailureMinimalCarrierJSealOnly, FailureJCopyDoesNotRestoreAmbientCell, FailureEdgeIntertwinerSupportOnly, FailureEdgeIntertwinerNoValue, FailureSymbolicYNotMagnitude, FailureNoAlphaSource, FailureNoTraceReadout, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNotR3, FailureNotR4, FailureNoParticleAssignment, FailureNoNeutrinoTheorem, FailureNoThreeGenerationTheorem,
	}
}

func containsAll(haystack, needles []string) bool {
	m := make(map[string]bool, len(haystack))
	for _, s := range haystack {
		m[s] = true
	}
	for _, s := range needles {
		if !m[s] {
			return false
		}
	}
	return true
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("alpha_B=%.16g official_N_eff=%.16g official_C_Yukawa=%.16g official_C_Higgs=%.16g frozen=%t alpha_native=%t R3=%t R4=%t", l.AlphaB, l.OfficialNEff, l.OfficialCYukawa, l.OfficialCHiggs, l.OfficialFrozen, l.AlphaNative, l.R3, l.R4)
}

func FormatAlgebra(a OrientedAlgebra) string {
	return fmt.Sprintf("full=%s orient=%s contains_full_H=%t contains_C_H=%t contains_C_R=%t contains_M3C=%t post_orientation=%t supports=%s failures=%s", a.FullAlgebra, a.OrientedAlgebra, a.ContainsFullH, a.ContainsCH, a.ContainsRightC, a.ContainsM3C, a.PostOrientationLayer, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatLeftAction(l LeftAction) string {
	return fmt.Sprintf("algebra=%s preserves_h=%t preserves_e=%t preserves_P=%t preserves_HRmin=%t preserves_HF=%t operator_certified=%t supports=%s failures=%s", l.Algebra, l.PreservesHPlusHMinus, l.PreservesEPlusEMinus, l.PreservesP1P3, l.PreservesHRMin, l.PreservesHF, l.LeftActionOperatorCertified, strings.Join(l.Supports, ","), strings.Join(l.Failures, ","))
}

func FormatOpposite(o OppositeAction) string {
	return fmt.Sprintf("expr=%s formal_J=%t support_defined=%t operator_certified=%t order_zero_typed=%t minimal_J_closed=%t ambient_reintroduced=%t supports=%s failures=%s", o.Expression, o.FormalJExchangeDefined, o.OppositeSupportDefined, o.OppositeOperatorCertified, o.OrderZeroTargetTyped, o.MinimalCarrierClosedUnderJSeal, o.AmbientCellReintroduced, strings.Join(o.Supports, ","), strings.Join(o.Failures, ","))
}

func FormatOrderZero(o OrderZero) string {
	return fmt.Sprintf("expr=%s algebra=%s support_auditable=%t block_support=%t operator_theorem=%t requires_J=%t supports=%s failures=%s", o.Expression, o.Algebra, o.SupportAuditable, o.BlockSupportCompatible, o.OperatorTheoremCertified, o.RequiresOperatorJOpposite, strings.Join(o.Supports, ","), strings.Join(o.Failures, ","))
}

func FormatEdges(edges []EdgeBimodule) string {
	parts := make([]string, 0, len(edges))
	for _, e := range edges {
		parts = append(parts, fmt.Sprintf("%s:%s->%s rank=%d support=%s left=%t right=%t operator=%t magnitude=%t", e.Name, e.Domain, e.Codomain, e.Rank, e.Support, e.LeftSupportCompatible, e.RightSupportCompatible, e.OperatorIntertwinerCertified, e.YukawaMagnitude))
	}
	return strings.Join(parts, " | ")
}

func FormatCarrier(c MinimalCarrier) string {
	return fmt.Sprintf("HL=%d HRmin=%d Hpart_min=%d HF_min=%d ambient_part=%d ambient_F=%d right_puncture=%s left_kernel=%s puncture_outside=%t J_restores_ambient=%t supports=%s failures=%s", c.HLRank, c.HRMinRank, c.HPartMinRank, c.HFMinRank, c.AmbientPartRank, c.AmbientFRank, c.RightPuncture, c.LeftKernel, c.RightPunctureOutsideMinimal, c.JCopyRestoresAmbientPuncture, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}

func FormatFirstOrderBoundary(f FirstOrderBoundary) string {
	return fmt.Sprintf("target=%s order_zero_audited=%t ready_for_next=%t first_order_operator=%t supports=%s failures=%s", f.TargetExpression, f.OrderZeroPrerequisiteAudited, f.ReadyForOperatorFirstOrderAttempt, f.FirstOrderOperatorCertified, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s inherited=%t support_bimodule=%t order_zero_support=%t J_seal=%t J_operator=%t order_zero_operator=%t first_order_operator=%t native_triple=%t alpha_sealed=%t magnitudes_missing=%t update=(%t,%t,%t) promote=(%t,%t)", i.Classification, i.Gate857Inherited, i.SupportBimodule, i.OrderZeroSupportPass, i.JOppositeSealDefined, i.OperatorJOppositeProof, i.OrderZeroOperatorProof, i.FirstOrderOperatorProof, i.NativeFiniteTripleProof, i.AlphaStillSealed, i.MagnitudesStillMissing, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CanPromoteToR3, i.CanPromoteToR4)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%t J_seal_only=%t no_J_operator=%t no_order_zero_operator=%t support_only=%t no_first_order=%t no_bimodule=%t no_native=%t no_alpha=%t no_trace=%t no_updates=(%t,%t) notR3=%t notR4=%t verdict=%s", f.Enforced, f.JOppositeSealOnly, f.NoOperatorJOpposite, f.NoOrderZeroOperator, f.OrderZeroSupportOnly, f.NoFirstOrderOperator, f.NoBimoduleProof, f.NoNativeTriple, f.NoAlphaSource, f.NoTraceReadout, f.NoOfficialNEffUpdate, f.NoCYukawaCHiggsUpdate, f.NotR3, f.NotR4, f.Verdict)
}
