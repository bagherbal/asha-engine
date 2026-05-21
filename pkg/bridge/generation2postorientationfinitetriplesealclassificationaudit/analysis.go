// Package generation2postorientationfinitetriplesealclassificationaudit implements
// Gate 863: Post-Orientation FiniteTriple Seal Classification Audit.
//
// Gate 863 follows Gate 862's socket-character intertwiner seal. The stabilizer
// branch now has a minimal active carrier, a Higgs-oriented algebra, scalar
// operator-valued edge sockets, color centrality, a character-identification
// seal, a puncture/kernel pair, and support/order-zero/first-order compatibility
// in the post-orientation layer. This gate does not try to prove a new theorem.
// It classifies the branch by layer and records exactly what is certified, what
// remains sealed, and why the construction is still not an R3 sector
// trace-ledger or R4 native Yukawa theorem.
package generation2postorientationfinitetriplesealclassificationaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE863-POST-ORIENTATION-FINITE-TRIPLE-SEAL-CLASSIFICATION-AUDIT"

	AlphaB          = 0.0003878958469680527
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	HPartAmbientRank = 16
	HFAmbientRank    = 32
	HLRank           = 8
	HRMinRank        = 7
	HPartMinRank     = 15
	HFMinRank        = 30
	YRank            = 7
	DSymRank         = 14
	KernelRank       = 1

	Classification = "POST_ORIENTATION_FINITE_TRIPLE_SEAL"
	Subtype        = "STABILIZER_BRANCH_FIRST_ORDER_COMPATIBLE_GIVEN_SOCKET_CHARACTER_SEAL"

	StatusGate862Inherited   = "PASS_GATE862_SOCKET_CHARACTER_INTERTWINER_SEAL_INHERITED"
	StatusLayerStackAudited  = "PASS_POST_ORIENTATION_FINITE_TRIPLE_LAYER_STACK_AUDITED"
	StatusFullAFBlocked      = "PASS_FULL_UNBROKEN_A_F_LAYER_BLOCKED_FOR_CURRENT_D_F_SYM"
	StatusOrientLayerSuccess = "PASS_A_F_ORIENT_STABILIZER_LAYER_CLASSIFIED_AS_SEAL_SUCCESS"
	StatusMinimalCarrier     = "PASS_MINIMAL_15_30_CARRIER_BRANCH_CLASSIFIED"
	StatusEdgeOperator       = "PASS_OPERATOR_VALUED_SCALAR_EDGE_SOCKET_MATRIX_CLASSIFIED"
	StatusFirstOrderSeal     = "PASS_STABILIZER_FIRST_ORDER_COMPATIBILITY_CLASSIFIED_AS_CONDITIONAL_SEAL"
	StatusPunctureKernel     = "PASS_CHIRAL_NEUTRAL_PUNCTURE_KERNEL_PAIR_CLASSIFIED"
	StatusR3Eligibility      = "PASS_R3_ELIGIBILITY_AUDITED_AND_BLOCKED"
	StatusNextWound          = "PASS_NEXT_WOUND_IDENTIFIED_AS_Y_DAGGER_Y_TRACE_MAGNITUDE_READOUT"
	StatusLedgerFrozen       = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusNoObservedDataUsed = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusFirewallVerdict    = "FIREWALL_PRESERVED_GATE863_CLASSIFICATION_NOT_R3_NOT_R4"

	SupportPostOrientationSeal  = "CONDITIONAL_SUPPORT_POST_ORIENTATION_FINITE_TRIPLE_SEAL"
	SupportStabilizerFirstOrder = "CONDITIONAL_SUPPORT_STABILIZER_BRANCH_FIRST_ORDER_COMPATIBILITY_GIVEN_SOCKET_CHARACTER_IDENTIFICATION"
	SupportScalarEdgeMatrix     = "CONDITIONAL_SUPPORT_OPERATOR_VALUED_SCALAR_EDGE_SOCKET_MATRIX"
	SupportMinimalCarrier       = "CONDITIONAL_SUPPORT_MINIMAL_15_30_CARRIER_BRANCH"
	SupportPunctureKernelPair   = "CONDITIONAL_SUPPORT_CHIRAL_NEUTRAL_PUNCTURE_KERNEL_PAIR"
	SupportAForientHome         = "CONDITIONAL_SUPPORT_D_F_SYM_LIVES_IN_A_F_ORIENT_POST_ORIENTATION_LAYER"
	SupportYdaggerYNext         = "CONDITIONAL_SUPPORT_NEXT_R3_PRESSURE_IS_Y_DAGGER_Y_TRACE_MAGNITUDE_READOUT"

	FailureCurrentDFNotFullAF        = "FAILED_ROUTE_CURRENT_D_F_SYM_NOT_FULL_UNBROKEN_A_F_THEOREM"
	FailureSocketIDNotNative         = "FAILED_ROUTE_SOCKET_CHARACTER_IDENTIFICATION_NOT_NATIVE"
	FailureHiggsOrientationNotNative = "FAILED_ROUTE_HIGGS_ORIENTATION_NOT_NATIVE"
	FailureAForientNotFullAF         = "FAILED_ROUTE_A_F_ORIENT_NOT_FULL_A_F"
	FailureNoNativeFiniteTriple      = "FAILED_ROUTE_NO_NATIVE_FINITE_TRIPLE_THEOREM"
	FailureNoFullUnbrokenFO          = "FAILED_ROUTE_NO_FULL_UNBROKEN_OPERATOR_FIRST_ORDER_THEOREM"
	FailureSymbolicYNotMagnitude     = "FAILED_ROUTE_SYMBOLIC_Y_NOT_YUKAWA_MAGNITUDE"
	FailureNoNumericalYukawa         = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureNoTraceMagnitudeReadout   = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoAlphaSource             = "FAILED_ROUTE_NO_ALPHA_B_SOURCE"
	FailureNoYToAggTheorem           = "FAILED_ROUTE_NO_Y_DAGGER_Y_TO_H_AGG_TRACE_MAGNITUDE_THEOREM"
	FailureNoOfficialNEffUpdate      = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate     = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoParticleAssignment      = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoNeutrinoTheorem         = "FAILED_ROUTE_NO_PHYSICAL_NEUTRINO_THEOREM"
	FailureNoThreeGenerationTheorem  = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
	FailureNoR3                      = "FAILED_ROUTE_POST_ORIENTATION_FINITE_TRIPLE_SEAL_NOT_R3"
	FailureNoR4                      = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	AlphaB                        float64
	OfficialNEff, OfficialCYukawa float64
	OfficialCHiggs                float64
	OfficialFrozen                bool
	AlphaNative, R3, R4           bool
}

type Layer struct {
	Name, Object, Classification, Status string
	Native, SealSuccess, FullUnbroken    bool
	Dimension                            int
	Supports, Failures                   []string
}

type CarrierClassification struct {
	AmbientPartRank, AmbientFullRank int
	MinimalPartRank, MinimalFullRank int
	LeftRank, RightMinRank           int
	RightPuncture, LeftKernel        string
	MinimalBranchSealed              bool
	AmbientBranchNative              bool
	Supports, Failures               []string
}

type EdgeOperatorClassification struct {
	YExpression, DSymExpression                                               string
	ColorCentral, ScalarSocket, CharacterMatchedBySeal, FirstOrderConditional bool
	NumericalYukawa, NativeFiniteTriple                                       bool
	YRank, DSymRank, KernelRank                                               int
	Supports, Failures                                                        []string
}

type R3Assessment struct {
	FiniteBodySealPresent, SectorTraceLedgerPresent      bool
	TraceMagnitudeReadoutPresent, YDaggerYShapeCandidate bool
	YSocketMagnitudesDerived, AlphaNative                bool
	EligibleForR3, EligibleForR4                         bool
	NextGate, NextQuestion                               string
	Supports, Failures                                   []string
}

type Impact struct {
	Classification, Subtype                                                              string
	PostOrientationFiniteTripleSeal, StabilizerFirstOrderConditional, NativeFiniteTriple bool
	YukawaMagnitudes, SectorTraceReadout, AlphaNative                                    bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4     bool
}

type Firewalls struct {
	Enforced                                                                                                  bool
	CurrentDFNotFullAF, SocketIDNotNative, HiggsOrientationNotNative, AForientNotFullAF, NoNativeFiniteTriple bool
	NoFullUnbrokenFirstOrder, SymbolicYNotMagnitude, NoNumericalYukawa, NoTraceReadout, NoAlphaSource         bool
	NoYToAggTheorem, NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate, NoParticleAssignment                        bool
	NoNeutrinoTheorem, NoThreeGenerationTheorem, NoR3, NoR4                                                   bool
	Verdict                                                                                                   string
}

type Audit struct {
	ID        string
	Ledger    Ledger
	Layers    []Layer
	Carrier   CarrierClassification
	Edge      EdgeOperatorClassification
	R3        R3Assessment
	Impact    Impact
	Firewalls Firewalls
	Truth     string
	Final     string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		ID:     AuditID,
		Ledger: Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true},
		Layers: []Layer{
			{Name: "native/full unbroken finite algebra", Object: "A_F=C plus H plus M_3(C)", Classification: "blocked for current D_F^sym", Status: StatusFullAFBlocked, Native: false, SealSuccess: false, FullUnbroken: true, Supports: []string{StatusFullAFBlocked}, Failures: []string{FailureCurrentDFNotFullAF}},
			{Name: "post-orientation stabilizer algebra", Object: "A_F^orient=C_R plus C_H plus M_3(C)", Classification: "post-Higgs-orientation stabilizer layer", Status: StatusOrientLayerSuccess, Native: false, SealSuccess: true, FullUnbroken: false, Supports: []string{StatusOrientLayerSuccess, SupportAForientHome}, Failures: []string{FailureAForientNotFullAF, FailureHiggsOrientationNotNative}},
			{Name: "minimal finite carrier", Object: "H_F^min=(H_L plus H_R^min) plus J(H_L plus H_R^min)", Classification: "minimal 15/30 representation seal", Status: StatusMinimalCarrier, Native: false, SealSuccess: true, FullUnbroken: false, Dimension: HFMinRank, Supports: []string{StatusMinimalCarrier, SupportMinimalCarrier}, Failures: []string{FailureNoNativeFiniteTriple}},
			{Name: "edge operator", Object: "D_F^sym=[[0,Y^dagger],[Y,0]]", Classification: "operator-valued symbolic support matrix", Status: StatusEdgeOperator, Native: false, SealSuccess: true, FullUnbroken: false, Supports: []string{StatusEdgeOperator, SupportScalarEdgeMatrix}, Failures: []string{FailureSymbolicYNotMagnitude, FailureNoNumericalYukawa}},
			{Name: "first-order compatibility", Object: "[[D_F^sym,rho_F(a)],rho_F^op(b)] for a,b in A_F^orient", Classification: "stabilizer-compatible given socket-character seal", Status: StatusFirstOrderSeal, Native: false, SealSuccess: true, FullUnbroken: false, Supports: []string{StatusFirstOrderSeal, SupportStabilizerFirstOrder}, Failures: []string{FailureSocketIDNotNative, FailureNoFullUnbrokenFO}},
		},
		Carrier:   CarrierClassification{AmbientPartRank: HPartAmbientRank, AmbientFullRank: HFAmbientRank, MinimalPartRank: HPartMinRank, MinimalFullRank: HFMinRank, LeftRank: HLRank, RightMinRank: HRMinRank, RightPuncture: "e_+ tensor P_1", LeftKernel: "h_+ tensor P_1", MinimalBranchSealed: true, AmbientBranchNative: false, Supports: []string{StatusMinimalCarrier, StatusPunctureKernel, SupportMinimalCarrier, SupportPunctureKernelPair}, Failures: []string{FailureNoNativeFiniteTriple, FailureNoParticleAssignment, FailureNoNeutrinoTheorem}},
		Edge:      EdgeOperatorClassification{YExpression: "Y=y_+3 |h_+><e_+| tensor I_P3 + y_-3 |h_-><e_-| tensor I_P3 + y_-1 |h_-><e_-| tensor I_P1, Y_+1=0", DSymExpression: "D_F^sym=[[0,Y^dagger],[Y,0]]", ColorCentral: true, ScalarSocket: true, CharacterMatchedBySeal: true, FirstOrderConditional: true, YRank: YRank, DSymRank: DSymRank, KernelRank: KernelRank, Supports: []string{StatusGate862Inherited, StatusEdgeOperator, StatusFirstOrderSeal, SupportScalarEdgeMatrix, SupportStabilizerFirstOrder}, Failures: []string{FailureSocketIDNotNative, FailureSymbolicYNotMagnitude, FailureNoTraceMagnitudeReadout}},
		R3:        R3Assessment{FiniteBodySealPresent: true, SectorTraceLedgerPresent: false, TraceMagnitudeReadoutPresent: false, YDaggerYShapeCandidate: true, YSocketMagnitudesDerived: false, AlphaNative: false, EligibleForR3: false, EligibleForR4: false, NextGate: "Gate 864 — Y^dagger Y TraceMagnitude Readout Obstruction Audit", NextQuestion: "Does Y^dagger Y have the right carrier shape, and are the socket magnitudes derived or inserted?", Supports: []string{StatusR3Eligibility, StatusNextWound, SupportYdaggerYNext}, Failures: []string{FailureNoTraceMagnitudeReadout, FailureNoYToAggTheorem, FailureNoAlphaSource, FailureNoR3, FailureNoR4}},
		Impact:    Impact{Classification: Classification, Subtype: Subtype, PostOrientationFiniteTripleSeal: true, StabilizerFirstOrderConditional: true},
		Firewalls: Firewalls{Enforced: true, CurrentDFNotFullAF: true, SocketIDNotNative: true, HiggsOrientationNotNative: true, AForientNotFullAF: true, NoNativeFiniteTriple: true, NoFullUnbrokenFirstOrder: true, SymbolicYNotMagnitude: true, NoNumericalYukawa: true, NoTraceReadout: true, NoAlphaSource: true, NoYToAggTheorem: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NoParticleAssignment: true, NoNeutrinoTheorem: true, NoThreeGenerationTheorem: true, NoR3: true, NoR4: true, Verdict: StatusFirewallVerdict},
		Truth:     "Gate 863 classifies the Gate 862 branch as a post-orientation finite-triple operator seal: coherent in the Higgs-oriented stabilizer layer, conditional on socket-character identification, and still below R3/R4.",
		Final:     "CLASSIFICATION: POST_ORIENTATION_FINITE_TRIPLE_SEAL; NEXT WOUND: Y^dagger Y -> trace magnitudes remains unproved.",
	}
	a.Impact.YukawaMagnitudes = false
	a.Impact.SectorTraceReadout = false
	a.Impact.AlphaNative = false
	a.Impact.CanUpdateNEff = false
	a.Impact.CanUpdateCYukawa = false
	a.Impact.CanUpdateCHiggs = false
	a.Impact.CanPromoteToR3 = false
	a.Impact.CanPromoteToR4 = false
	if err := a.Validate(); err != nil {
		return Audit{}, err
	}
	return a, nil
}

func (a Audit) Validate() error {
	err := func(msg string) error { return fmt.Errorf("%s: %s", AuditID, msg) }
	if a.ID != AuditID {
		return err("bad audit id")
	}
	if len(a.Layers) != 5 {
		return err("expected five classification layers")
	}
	for _, l := range a.Layers {
		if l.Native {
			return err("classification layer overpromoted to native")
		}
	}
	if a.Carrier.AmbientPartRank != HPartAmbientRank || a.Carrier.AmbientFullRank != HFAmbientRank || a.Carrier.MinimalPartRank != HPartMinRank || a.Carrier.MinimalFullRank != HFMinRank || !a.Carrier.MinimalBranchSealed || a.Carrier.AmbientBranchNative {
		return err("carrier classification malformed")
	}
	if !a.Edge.ColorCentral || !a.Edge.ScalarSocket || !a.Edge.CharacterMatchedBySeal || !a.Edge.FirstOrderConditional || a.Edge.NumericalYukawa || a.Edge.NativeFiniteTriple || a.Edge.YRank != YRank || a.Edge.DSymRank != DSymRank || a.Edge.KernelRank != KernelRank {
		return err("edge classification malformed or overpromoted")
	}
	if !a.R3.FiniteBodySealPresent || a.R3.SectorTraceLedgerPresent || a.R3.TraceMagnitudeReadoutPresent || !a.R3.YDaggerYShapeCandidate || a.R3.YSocketMagnitudesDerived || a.R3.AlphaNative || a.R3.EligibleForR3 || a.R3.EligibleForR4 {
		return err("R3 assessment malformed or overpromoted")
	}
	if !a.Impact.PostOrientationFiniteTripleSeal || !a.Impact.StabilizerFirstOrderConditional || a.Impact.NativeFiniteTriple || a.Impact.YukawaMagnitudes || a.Impact.SectorTraceReadout || a.Impact.AlphaNative || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		return err("impact overpromoted")
	}
	if !a.Firewalls.Enforced || !a.Firewalls.CurrentDFNotFullAF || !a.Firewalls.SocketIDNotNative || !a.Firewalls.HiggsOrientationNotNative || !a.Firewalls.AForientNotFullAF || !a.Firewalls.NoNativeFiniteTriple || !a.Firewalls.NoFullUnbrokenFirstOrder || !a.Firewalls.SymbolicYNotMagnitude || !a.Firewalls.NoNumericalYukawa || !a.Firewalls.NoTraceReadout || !a.Firewalls.NoAlphaSource || !a.Firewalls.NoYToAggTheorem || !a.Firewalls.NoOfficialNEffUpdate || !a.Firewalls.NoCYukawaCHiggsUpdate || !a.Firewalls.NoParticleAssignment || !a.Firewalls.NoNeutrinoTheorem || !a.Firewalls.NoThreeGenerationTheorem || !a.Firewalls.NoR3 || !a.Firewalls.NoR4 || a.Firewalls.Verdict != StatusFirewallVerdict {
		return err("firewalls not preserved")
	}
	return nil
}

func Statuses() []string {
	return []string{
		StatusGate862Inherited, StatusLayerStackAudited, StatusFullAFBlocked, StatusOrientLayerSuccess, StatusMinimalCarrier,
		StatusEdgeOperator, StatusFirstOrderSeal, StatusPunctureKernel, StatusR3Eligibility, StatusNextWound,
		StatusLedgerFrozen, StatusNoObservedDataUsed, StatusFirewallVerdict,
		SupportPostOrientationSeal, SupportStabilizerFirstOrder, SupportScalarEdgeMatrix, SupportMinimalCarrier,
		SupportPunctureKernelPair, SupportAForientHome, SupportYdaggerYNext,
		FailureCurrentDFNotFullAF, FailureSocketIDNotNative, FailureHiggsOrientationNotNative, FailureAForientNotFullAF,
		FailureNoNativeFiniteTriple, FailureNoFullUnbrokenFO, FailureSymbolicYNotMagnitude, FailureNoNumericalYukawa,
		FailureNoTraceMagnitudeReadout, FailureNoAlphaSource, FailureNoYToAggTheorem, FailureNoOfficialNEffUpdate,
		FailureNoCYukawaCHiggsUpdate, FailureNoParticleAssignment, FailureNoNeutrinoTheorem, FailureNoThreeGenerationTheorem,
		FailureNoR3, FailureNoR4,
	}
}

func containsAll(hay []string, needles []string) bool {
	set := map[string]bool{}
	for _, h := range hay {
		set[h] = true
	}
	for _, n := range needles {
		if !set[n] {
			return false
		}
	}
	return true
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("alpha_B=%.16g official_N_eff=%.16g official_C_Yukawa=%.16g official_C_Higgs=%.16g frozen=%t alpha_native=%t R3=%t R4=%t", l.AlphaB, l.OfficialNEff, l.OfficialCYukawa, l.OfficialCHiggs, l.OfficialFrozen, l.AlphaNative, l.R3, l.R4)
}

func FormatLayers(ls []Layer) string {
	parts := make([]string, 0, len(ls))
	for _, l := range ls {
		parts = append(parts, fmt.Sprintf("%s => %s [%s] native=%t seal_success=%t full_unbroken=%t", l.Name, l.Object, l.Classification, l.Native, l.SealSuccess, l.FullUnbroken))
	}
	return strings.Join(parts, " | ")
}

func FormatCarrier(c CarrierClassification) string {
	return fmt.Sprintf("ambient=%d/%d minimal=%d/%d H_L=%d H_R_min=%d right_puncture=%s left_kernel=%s minimal_sealed=%t", c.AmbientPartRank, c.AmbientFullRank, c.MinimalPartRank, c.MinimalFullRank, c.LeftRank, c.RightMinRank, c.RightPuncture, c.LeftKernel, c.MinimalBranchSealed)
}

func FormatEdge(e EdgeOperatorClassification) string {
	return fmt.Sprintf("Y=%s D=%s color_central=%t scalar_socket=%t character_seal=%t first_order_conditional=%t rankY=%d rankD=%d kernel=%d", e.YExpression, e.DSymExpression, e.ColorCentral, e.ScalarSocket, e.CharacterMatchedBySeal, e.FirstOrderConditional, e.YRank, e.DSymRank, e.KernelRank)
}

func FormatR3(r R3Assessment) string {
	return fmt.Sprintf("finite_body=%t sector_ledger=%t trace_readout=%t YdaggerY_shape=%t y_magnitudes_derived=%t alpha_native=%t R3=%t R4=%t next=%s question=%s", r.FiniteBodySealPresent, r.SectorTraceLedgerPresent, r.TraceMagnitudeReadoutPresent, r.YDaggerYShapeCandidate, r.YSocketMagnitudesDerived, r.AlphaNative, r.EligibleForR3, r.EligibleForR4, r.NextGate, r.NextQuestion)
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s subtype=%s post_orientation_seal=%t stabilizer_FO_conditional=%t native=%t yukawa_magnitudes=%t trace_readout=%t alpha_native=%t update_Neff=%t update_CYukawa=%t update_CHiggs=%t R3=%t R4=%t", i.Classification, i.Subtype, i.PostOrientationFiniteTripleSeal, i.StabilizerFirstOrderConditional, i.NativeFiniteTriple, i.YukawaMagnitudes, i.SectorTraceReadout, i.AlphaNative, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CanPromoteToR3, i.CanPromoteToR4)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%t current_D_not_full_AF=%t socket_ID_not_native=%t Higgs_orientation_not_native=%t AForient_not_full_AF=%t no_native_finite_triple=%t no_full_unbroken_FO=%t symbolic_Y_not_magnitude=%t no_trace_readout=%t no_alpha_source=%t no_YdaggerY_theorem=%t no_R3=%t no_R4=%t verdict=%s", f.Enforced, f.CurrentDFNotFullAF, f.SocketIDNotNative, f.HiggsOrientationNotNative, f.AForientNotFullAF, f.NoNativeFiniteTriple, f.NoFullUnbrokenFirstOrder, f.SymbolicYNotMagnitude, f.NoTraceReadout, f.NoAlphaSource, f.NoYToAggTheorem, f.NoR3, f.NoR4, f.Verdict)
}
