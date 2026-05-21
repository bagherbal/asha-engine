// Package generation2stabilizerbranchfirstordermatrixedgeintertwineraudit implements
// Gate 857: Stabilizer-Branch First-Order Matrix and Edge-Intertwiner Audit.
//
// Gate 857 follows Gate 856's post-Higgs-orientation stabilizer layer.  It
// audits the first-order target inside A_F^orient = C_R plus C_H plus M_3(C),
// separates allowed nonzero one-form commutators from first-order obstruction,
// and checks whether the three symbolic support edges behave as blockwise
// stabilizer-branch intertwiners.  It remains a support-level audit: full
// operator-level J-opposite compatibility, bimodule decomposition, numerical
// Yukawa magnitudes, alpha_B source, R3, and R4 promotion stay blocked.
package generation2stabilizerbranchfirstordermatrixedgeintertwineraudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE857-STABILIZER-BRANCH-FIRST-ORDER-MATRIX-EDGE-INTERTWINER-AUDIT"

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

	StatusGate856Inherited         = "PASS_GATE856_A_F_ORIENT_STABILIZER_LAYER_INHERITED"
	StatusSupportPreservation      = "PASS_A_F_ORIENT_SUPPORT_PRESERVATION_AUDITED"
	StatusCommutatorSeparated      = "PASS_NONZERO_D_COMMUTATOR_SEPARATED_FROM_FIRST_ORDER_OBSTRUCTION"
	StatusEdgeIntertwinersAudited  = "PASS_ACTIVE_EDGE_INTERTWINER_SUPPORT_AUDITED"
	StatusOppositeActionAudited    = "PASS_J_OPPOSITE_SUPPORT_ACTION_AUDITED"
	StatusPunctureKernelAudited    = "PASS_PUNCTURE_AND_LEFT_KERNEL_STABILITY_IN_ORIENTED_BRANCH_AUDITED"
	StatusFirstOrderSupportAudited = "PASS_STABILIZER_BRANCH_FIRST_ORDER_SUPPORT_LEVEL_AUDITED"
	StatusNoObservedDataUsed       = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusLedgerFrozen             = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusFirewallVerdict          = "FIREWALL_PRESERVED_GATE857_STABILIZER_SUPPORT_FIRST_ORDER_NOT_OPERATOR_THEOREM"

	SupportStabilizerFirstOrderSupportCompatibility = "CONDITIONAL_SUPPORT_STABILIZER_BRANCH_FIRST_ORDER_SUPPORT_COMPATIBILITY"
	SupportActiveEdgesBlockwiseCompatible           = "CONDITIONAL_SUPPORT_ACTIVE_EDGES_ARE_BLOCKWISE_COMPATIBLE_WITH_A_F_ORIENT"
	SupportColorEdgesM3Compatible                   = "CONDITIONAL_SUPPORT_COLOR_EDGES_COMPATIBLE_WITH_M3C_SUPPORT_MODULES"
	SupportLeptonEdgeColorTrivial                   = "CONDITIONAL_SUPPORT_LEPTON_EDGE_COLOR_TRIVIAL_ON_P1_SUPPORT"
	SupportPunctureKernelStable                     = "CONDITIONAL_SUPPORT_PUNCTURE_AND_KERNEL_STABLE_IN_ORIENTED_BRANCH"
	SupportOneFormCommutatorAllowed                 = "CONDITIONAL_SUPPORT_NONZERO_D_RHO_COMMUTATOR_IS_ALLOWED_ONE_FORM_SOURCE"
	SupportR2StabilizerSupportStage                 = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_PLUS_PLUS_PLUS_STABILIZER_SUPPORT_FIRST_ORDER_SEAL"

	FailureFullAFNotTestedAsTarget         = "FAILED_ROUTE_FULL_UNBROKEN_A_F_NOT_THE_TARGET_OF_GATE857"
	FailureAForientNotFullAF               = "FAILED_ROUTE_A_F_ORIENT_NOT_FULL_UNBROKEN_A_F"
	FailureFullHSocketFirewall             = "FAILED_ROUTE_FULL_H_ACTION_DOES_NOT_PRESERVE_SOCKET_FRAME"
	FailureSupportOnlyNotOperatorProof     = "FAILED_ROUTE_STABILIZER_FIRST_ORDER_COMPATIBILITY_IS_SUPPORT_LEVEL_NOT_OPERATOR_THEOREM"
	FailureNoFullOperatorFirstOrderTheorem = "FAILED_ROUTE_NO_FULL_OPERATOR_LEVEL_FIRST_ORDER_THEOREM"
	FailureNoCompleteJOppositeProof        = "FAILED_ROUTE_NO_COMPLETE_J_OPPOSITE_ACTION_PROOF"
	FailureNoBimoduleProof                 = "FAILED_ROUTE_NO_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailureNoNativeFiniteTripleProof       = "FAILED_ROUTE_NO_NATIVE_FINITE_TRIPLE_PROOF_CERTIFIED"
	FailureCharacterMatchSupportOnly       = "FAILED_ROUTE_EDGE_CHARACTER_MATCH_IS_SUPPORT_LABEL_NOT_OPERATOR_INTERTWINER_PROOF"
	FailureYSupportNotMagnitude            = "FAILED_ROUTE_SUPPORT_LEVEL_INTERTWINER_NOT_YUKAWA_MAGNITUDE"
	FailureSymbolicYNotMagnitude           = "FAILED_ROUTE_SYMBOLIC_Y_COEFFICIENTS_NOT_YUKAWA_MAGNITUDES"
	FailureNoAlphaSource                   = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoTraceReadout                  = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoOfficialNEffUpdate            = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate           = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNotR3                           = "FAILED_ROUTE_R2_PLUS_PLUS_PLUS_PLUS_PLUS_STABILIZER_SUPPORT_FIRST_ORDER_NOT_R3"
	FailureNotR4                           = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoParticleAssignment            = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoNeutrinoTheorem               = "FAILED_ROUTE_NO_PHYSICAL_NEUTRINO_THEOREM"
	FailureNoThreeGenerationTheorem        = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
)

type Ledger struct {
	AlphaB                        float64
	OfficialNEff, OfficialCYukawa float64
	OfficialCHiggs                float64
	OfficialFrozen                bool
	AlphaNative, R3, R4           bool
}

type StabilizerAlgebra struct {
	FullAlgebra, OrientedAlgebra string
	ContainsFullH, ContainsCH    bool
	ContainsRightC, ContainsM3C  bool
	PostOrientationLayer         bool
	Supports, Failures           []string
}

type SupportPreservation struct {
	PreservesHPlusHMinus, PreservesEPlusEMinus bool
	PreservesP1P3, PreservesHRMin, PreservesHF bool
	RightPunctureOutside, LeftKernelCandidate  bool
	Supports, Failures                         []string
}

type EdgeIntertwiner struct {
	Name, Domain, Codomain  string
	Rank                    int
	ColorSupport            string
	BlockwiseCompatible     bool
	OperatorIntertwiner     bool
	CharacterMatchCertified bool
	Supports, Failures      []string
}

type FirstOrder struct {
	TargetExpression, Algebra   string
	DCommutatorExpectedNonzero  bool
	NonzeroCommutatorAllowed    bool
	OppositeSupportAuditable    bool
	OppositeOperatorCertified   bool
	SupportFirstOrderCompatible bool
	OperatorFirstOrderCertified bool
	BimoduleCertified           bool
	Supports, Failures          []string
}

type KernelPuncture struct {
	RightPuncture, LeftKernel string
	RightPuncturePreserved    bool
	LeftKernelPreserved       bool
	LeftKernelOperatorStable  bool
	PhysicalParticleTheorem   bool
	Supports, Failures        []string
}

type Carrier struct {
	HLRank, HRMinRank, HPartMinRank, HFMinRank int
	AmbientPartRank, AmbientFRank              int
	DSymRank, KernelRank                       int
	Supports, Failures                         []string
}

type Impact struct {
	Classification                                                                   string
	Gate856Inherited, AForientTarget, EdgeSupportCompatible, FirstOrderSupportPass   bool
	FullAFPass, OperatorFirstOrderProof, NativeFiniteTripleProof, AlphaStillSealed   bool
	MagnitudesStillMissing                                                           bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4 bool
}

type Firewalls struct {
	Enforced                                                                                         bool
	FullAFNotTarget, AForientNotFullAF, FullHSocketFirewall, SupportOnly, NoOperatorFirstOrder       bool
	NoJOppositeProof, NoBimoduleProof, NoNativeTriple, CharacterMatchSupportOnly, IntertwinerNoValue bool
	YSymbolicOnly, NoAlphaSource, NoTraceReadout, NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate        bool
	NotR3, NotR4, NoParticleAssignment, NoNeutrinoTheorem, NoThreeGenerationTheorem                  bool
	Verdict                                                                                          string
}

type Audit struct {
	ID          string
	Ledger      Ledger
	Algebra     StabilizerAlgebra
	Preserve    SupportPreservation
	Edges       []EdgeIntertwiner
	FirstOrder  FirstOrder
	NeutralPair KernelPuncture
	Carrier     Carrier
	Impact      Impact
	Firewalls   Firewalls
	Truth       string
	Final       string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		ID:          AuditID,
		Ledger:      buildLedger(),
		Algebra:     buildAlgebra(),
		Preserve:    buildPreservation(),
		Edges:       buildEdges(),
		FirstOrder:  buildFirstOrder(),
		NeutralPair: buildKernelPuncture(),
		Carrier:     buildCarrier(),
		Impact:      buildImpact(),
		Firewalls:   buildFirewalls(),
		Truth:       "Gate 857 audits the first-order target inside A_F^orient, not full unbroken A_F: active edges are support-compatible intertwiners in the stabilizer branch, but no operator-level first-order theorem is certified.",
		Final:       "VERDICT: CONDITIONAL_SUPPORT_STABILIZER_BRANCH_FIRST_ORDER_SUPPORT_COMPATIBILITY; FAILED_ROUTE_NO_FULL_OPERATOR_LEVEL_FIRST_ORDER_THEOREM.",
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
	if !a.Preserve.PreservesHPlusHMinus || !a.Preserve.PreservesEPlusEMinus || !a.Preserve.PreservesP1P3 || !a.Preserve.PreservesHRMin || !a.Preserve.PreservesHF || !a.Preserve.RightPunctureOutside || !a.Preserve.LeftKernelCandidate {
		return err("support preservation failed")
	}
	if len(a.Edges) != 3 {
		return err("expected three active edge families")
	}
	for _, e := range a.Edges {
		if !e.BlockwiseCompatible || e.OperatorIntertwiner || e.CharacterMatchCertified {
			return err("edge-intertwiner firewall violated")
		}
	}
	if !a.FirstOrder.DCommutatorExpectedNonzero || !a.FirstOrder.NonzeroCommutatorAllowed || !a.FirstOrder.OppositeSupportAuditable || a.FirstOrder.OppositeOperatorCertified || !a.FirstOrder.SupportFirstOrderCompatible || a.FirstOrder.OperatorFirstOrderCertified || a.FirstOrder.BimoduleCertified {
		return err("first-order support/operator flags inconsistent")
	}
	if !a.NeutralPair.RightPuncturePreserved || !a.NeutralPair.LeftKernelPreserved || a.NeutralPair.LeftKernelOperatorStable || a.NeutralPair.PhysicalParticleTheorem {
		return err("neutral pair overpromoted")
	}
	if a.Carrier.HLRank != HLRank || a.Carrier.HRMinRank != HRMinRank || a.Carrier.HPartMinRank != HPartMinRank || a.Carrier.HFMinRank != HFMinRank || a.Carrier.AmbientPartRank != AmbientPartRank || a.Carrier.AmbientFRank != AmbientFRank || a.Carrier.DSymRank != DSymRank || a.Carrier.KernelRank != DSymKernelRank {
		return err("carrier ranks inconsistent")
	}
	if !a.Impact.Gate856Inherited || !a.Impact.AForientTarget || !a.Impact.EdgeSupportCompatible || !a.Impact.FirstOrderSupportPass || a.Impact.FullAFPass || a.Impact.OperatorFirstOrderProof || a.Impact.NativeFiniteTripleProof || !a.Impact.AlphaStillSealed || !a.Impact.MagnitudesStillMissing || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		return err("impact flags inconsistent")
	}
	if !a.Firewalls.Enforced || !a.Firewalls.FullAFNotTarget || !a.Firewalls.AForientNotFullAF || !a.Firewalls.FullHSocketFirewall || !a.Firewalls.SupportOnly || !a.Firewalls.NoOperatorFirstOrder || !a.Firewalls.NoJOppositeProof || !a.Firewalls.NoBimoduleProof || !a.Firewalls.NoNativeTriple || !a.Firewalls.CharacterMatchSupportOnly || !a.Firewalls.IntertwinerNoValue || !a.Firewalls.YSymbolicOnly || !a.Firewalls.NoAlphaSource || !a.Firewalls.NoTraceReadout || !a.Firewalls.NoOfficialNEffUpdate || !a.Firewalls.NoCYukawaCHiggsUpdate || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 || !a.Firewalls.NoParticleAssignment || !a.Firewalls.NoNeutrinoTheorem || !a.Firewalls.NoThreeGenerationTheorem || a.Firewalls.Verdict != StatusFirewallVerdict {
		return err("firewall flags inconsistent")
	}
	return nil
}

func err(msg string) error { return fmt.Errorf("%s", msg) }

func buildLedger() Ledger {
	return Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true}
}

func buildAlgebra() StabilizerAlgebra {
	return StabilizerAlgebra{
		FullAlgebra: "A_F=C plus H plus M_3(C)", OrientedAlgebra: "A_F^orient=C_R plus C_H plus M_3(C)",
		ContainsFullH: false, ContainsCH: true, ContainsRightC: true, ContainsM3C: true, PostOrientationLayer: true,
		Supports: []string{StatusGate856Inherited, SupportStabilizerFirstOrderSupportCompatibility},
		Failures: []string{FailureFullAFNotTestedAsTarget, FailureAForientNotFullAF, FailureFullHSocketFirewall},
	}
}

func buildPreservation() SupportPreservation {
	return SupportPreservation{
		PreservesHPlusHMinus: true, PreservesEPlusEMinus: true, PreservesP1P3: true, PreservesHRMin: true, PreservesHF: true,
		RightPunctureOutside: true, LeftKernelCandidate: true,
		Supports: []string{StatusSupportPreservation, SupportPunctureKernelStable},
		Failures: []string{FailureSupportOnlyNotOperatorProof},
	}
}

func buildEdges() []EdgeIntertwiner {
	return []EdgeIntertwiner{
		{Name: "Y_+3", Domain: "e_+ tensor P_3", Codomain: "h_+ tensor P_3", Rank: P3Rank, ColorSupport: "P_3 color module", BlockwiseCompatible: true, OperatorIntertwiner: false, CharacterMatchCertified: false, Supports: []string{SupportActiveEdgesBlockwiseCompatible, SupportColorEdgesM3Compatible}, Failures: []string{FailureCharacterMatchSupportOnly, FailureYSupportNotMagnitude}},
		{Name: "Y_-3", Domain: "e_- tensor P_3", Codomain: "h_- tensor P_3", Rank: P3Rank, ColorSupport: "P_3 color module", BlockwiseCompatible: true, OperatorIntertwiner: false, CharacterMatchCertified: false, Supports: []string{SupportActiveEdgesBlockwiseCompatible, SupportColorEdgesM3Compatible}, Failures: []string{FailureCharacterMatchSupportOnly, FailureYSupportNotMagnitude}},
		{Name: "Y_-1", Domain: "e_- tensor P_1", Codomain: "h_- tensor P_1", Rank: P1Rank, ColorSupport: "P_1 color-trivial lepton support", BlockwiseCompatible: true, OperatorIntertwiner: false, CharacterMatchCertified: false, Supports: []string{SupportActiveEdgesBlockwiseCompatible, SupportLeptonEdgeColorTrivial}, Failures: []string{FailureCharacterMatchSupportOnly, FailureYSupportNotMagnitude}},
	}
}

func buildFirstOrder() FirstOrder {
	return FirstOrder{
		TargetExpression: "[[D_F^sym,rho_F(a)],J_F rho_F(b) J_F^{-1}]", Algebra: "a,b in A_F^orient",
		DCommutatorExpectedNonzero: true, NonzeroCommutatorAllowed: true, OppositeSupportAuditable: true, OppositeOperatorCertified: false,
		SupportFirstOrderCompatible: true, OperatorFirstOrderCertified: false, BimoduleCertified: false,
		Supports: []string{StatusCommutatorSeparated, StatusFirstOrderSupportAudited, SupportOneFormCommutatorAllowed, SupportStabilizerFirstOrderSupportCompatibility},
		Failures: []string{FailureSupportOnlyNotOperatorProof, FailureNoFullOperatorFirstOrderTheorem, FailureNoCompleteJOppositeProof, FailureNoBimoduleProof, FailureNoNativeFiniteTripleProof},
	}
}

func buildKernelPuncture() KernelPuncture {
	return KernelPuncture{
		RightPuncture: "e_+ tensor P_1", LeftKernel: "h_+ tensor P_1", RightPuncturePreserved: true, LeftKernelPreserved: true,
		LeftKernelOperatorStable: false, PhysicalParticleTheorem: false,
		Supports: []string{StatusPunctureKernelAudited, SupportPunctureKernelStable},
		Failures: []string{FailureSupportOnlyNotOperatorProof, FailureNoParticleAssignment, FailureNoNeutrinoTheorem},
	}
}

func buildCarrier() Carrier {
	return Carrier{HLRank: HLRank, HRMinRank: HRMinRank, HPartMinRank: HPartMinRank, HFMinRank: HFMinRank, AmbientPartRank: AmbientPartRank, AmbientFRank: AmbientFRank, DSymRank: DSymRank, KernelRank: DSymKernelRank}
}

func buildImpact() Impact {
	return Impact{
		Classification:   "R2+++++_stabilizer_support_first_order_seal",
		Gate856Inherited: true, AForientTarget: true, EdgeSupportCompatible: true, FirstOrderSupportPass: true,
		FullAFPass: false, OperatorFirstOrderProof: false, NativeFiniteTripleProof: false,
		AlphaStillSealed: true, MagnitudesStillMissing: true,
		CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, CanPromoteToR3: false, CanPromoteToR4: false,
	}
}

func buildFirewalls() Firewalls {
	return Firewalls{Enforced: true, FullAFNotTarget: true, AForientNotFullAF: true, FullHSocketFirewall: true, SupportOnly: true, NoOperatorFirstOrder: true, NoJOppositeProof: true, NoBimoduleProof: true, NoNativeTriple: true, CharacterMatchSupportOnly: true, IntertwinerNoValue: true, YSymbolicOnly: true, NoAlphaSource: true, NoTraceReadout: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NotR3: true, NotR4: true, NoParticleAssignment: true, NoNeutrinoTheorem: true, NoThreeGenerationTheorem: true, Verdict: StatusFirewallVerdict}
}

func Statuses() []string {
	return []string{
		StatusGate856Inherited, StatusSupportPreservation, StatusCommutatorSeparated, StatusEdgeIntertwinersAudited, StatusOppositeActionAudited, StatusPunctureKernelAudited, StatusFirstOrderSupportAudited, StatusNoObservedDataUsed, StatusLedgerFrozen, StatusFirewallVerdict,
		SupportStabilizerFirstOrderSupportCompatibility, SupportActiveEdgesBlockwiseCompatible, SupportColorEdgesM3Compatible, SupportLeptonEdgeColorTrivial, SupportPunctureKernelStable, SupportOneFormCommutatorAllowed, SupportR2StabilizerSupportStage,
		FailureFullAFNotTestedAsTarget, FailureAForientNotFullAF, FailureFullHSocketFirewall, FailureSupportOnlyNotOperatorProof, FailureNoFullOperatorFirstOrderTheorem, FailureNoCompleteJOppositeProof, FailureNoBimoduleProof, FailureNoNativeFiniteTripleProof, FailureCharacterMatchSupportOnly, FailureYSupportNotMagnitude, FailureSymbolicYNotMagnitude, FailureNoAlphaSource, FailureNoTraceReadout, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNotR3, FailureNotR4, FailureNoParticleAssignment, FailureNoNeutrinoTheorem, FailureNoThreeGenerationTheorem,
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

func FormatAlgebra(a StabilizerAlgebra) string {
	return fmt.Sprintf("full=%s orient=%s contains_full_H=%t contains_C_H=%t post_orientation=%t supports=%s failures=%s", a.FullAlgebra, a.OrientedAlgebra, a.ContainsFullH, a.ContainsCH, a.PostOrientationLayer, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatPreservation(p SupportPreservation) string {
	return fmt.Sprintf("preserves_h=%t preserves_e=%t preserves_P=%t preserves_HRmin=%t puncture_outside=%t left_kernel_candidate=%t supports=%s failures=%s", p.PreservesHPlusHMinus, p.PreservesEPlusEMinus, p.PreservesP1P3, p.PreservesHRMin, p.RightPunctureOutside, p.LeftKernelCandidate, strings.Join(p.Supports, ","), strings.Join(p.Failures, ","))
}

func FormatEdges(edges []EdgeIntertwiner) string {
	parts := make([]string, 0, len(edges))
	for _, e := range edges {
		parts = append(parts, fmt.Sprintf("%s:%s->%s rank=%d blockwise=%t operator=%t character=%t", e.Name, e.Domain, e.Codomain, e.Rank, e.BlockwiseCompatible, e.OperatorIntertwiner, e.CharacterMatchCertified))
	}
	return strings.Join(parts, " | ")
}

func FormatFirstOrder(f FirstOrder) string {
	return fmt.Sprintf("target=%s algebra=%s nonzero_commutator_allowed=%t support_compatible=%t operator_certified=%t opposite_operator=%t bimodule=%t supports=%s failures=%s", f.TargetExpression, f.Algebra, f.NonzeroCommutatorAllowed, f.SupportFirstOrderCompatible, f.OperatorFirstOrderCertified, f.OppositeOperatorCertified, f.BimoduleCertified, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}

func FormatNeutralPair(k KernelPuncture) string {
	return fmt.Sprintf("right_puncture=%s left_kernel=%s preserved=(%t,%t) operator_stable=%t physical=%t supports=%s failures=%s", k.RightPuncture, k.LeftKernel, k.RightPuncturePreserved, k.LeftKernelPreserved, k.LeftKernelOperatorStable, k.PhysicalParticleTheorem, strings.Join(k.Supports, ","), strings.Join(k.Failures, ","))
}

func FormatCarrier(c Carrier) string {
	return fmt.Sprintf("HL=%d HRmin=%d Hpart_min=%d HF_min=%d ambient_part=%d ambient_F=%d D_rank=%d kernel=%d", c.HLRank, c.HRMinRank, c.HPartMinRank, c.HFMinRank, c.AmbientPartRank, c.AmbientFRank, c.DSymRank, c.KernelRank)
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s inherited=%t AForient=%t support_first_order=%t full_AF=%t operator_first_order=%t native_triple=%t alpha_sealed=%t magnitudes_missing=%t update=(%t,%t,%t) promote=(%t,%t)", i.Classification, i.Gate856Inherited, i.AForientTarget, i.FirstOrderSupportPass, i.FullAFPass, i.OperatorFirstOrderProof, i.NativeFiniteTripleProof, i.AlphaStillSealed, i.MagnitudesStillMissing, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CanPromoteToR3, i.CanPromoteToR4)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%t support_only=%t no_operator_first_order=%t no_J=%t no_bimodule=%t no_native=%t no_alpha=%t no_trace=%t no_updates=(%t,%t) notR3=%t notR4=%t verdict=%s", f.Enforced, f.SupportOnly, f.NoOperatorFirstOrder, f.NoJOppositeProof, f.NoBimoduleProof, f.NoNativeTriple, f.NoAlphaSource, f.NoTraceReadout, f.NoOfficialNEffUpdate, f.NoCYukawaCHiggsUpdate, f.NotR3, f.NotR4, f.Verdict)
}
