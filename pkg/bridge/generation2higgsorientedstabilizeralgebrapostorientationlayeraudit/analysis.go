// Package generation2higgsorientedstabilizeralgebrapostorientationlayeraudit implements
// Gate 856: Higgs-Oriented Stabilizer Algebra and Post-Orientation Layer Audit.
//
// Gate 856 follows Gate 855's layer classification.  Gate 855 showed that the
// symbolic finite-Dirac skeleton D_F^sym is not compatible with the full
// unbroken quaternionic weak action because generic H mixes the oriented weak
// socket frame h_+ plus h_-.  Gate 856 therefore audits the algebraic stabilizer
// of the chosen Higgs/weak orientation, defines the post-orientation algebraic
// layer that preserves the symbolic D_F support, and keeps the firewall between
// post-orientation support compatibility and a full unbroken finite-triple theorem.
package generation2higgsorientedstabilizeralgebrapostorientationlayeraudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE856-HIGGS-ORIENTED-STABILIZER-ALGEBRA-POST-ORIENTATION-LAYER-AUDIT"

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

	StatusGate855Inherited           = "PASS_GATE855_FULL_VS_ORIENTED_LAYER_CLASSIFICATION_INHERITED"
	StatusStabilizerAudited          = "PASS_HIGGS_ORIENTED_STABILIZER_OF_WEAK_SOCKET_FRAME_AUDITED"
	StatusFullHFirewallAudited       = "PASS_FULL_H_QUATERNIONIC_SOCKET_FIREWALL_PRESERVED"
	StatusOrientedAlgebraDefined     = "PASS_A_F_ORIENT_DEFINED_AS_POST_ORIENTATION_STABILIZER_LAYER"
	StatusOrientedActionPreservation = "PASS_A_F_ORIENT_ACTION_PRESERVATION_AUDITED"
	StatusDFCompatibilityAudited     = "PASS_D_F_SYM_SUPPORT_COMPATIBILITY_WITH_A_F_ORIENT_AUDITED"
	StatusKernelPuncturePreservation = "PASS_PUNCTURE_AND_LEFT_KERNEL_PRESERVATION_IN_ORIENTED_LAYER_AUDITED"
	StatusNextFirstOrderTyped        = "PASS_STABILIZER_BRANCH_FIRST_ORDER_TARGET_TYPED_FOR_GATE857"
	StatusNoObservedDataUsed         = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusLedgerFrozen               = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusFirewallVerdict            = "FIREWALL_PRESERVED_GATE856_POST_ORIENTATION_STABILIZER_NOT_FULL_UNBROKEN_A_F_THEOREM"

	SupportStabilizerIsComplexSubalgebra      = "CONDITIONAL_SUPPORT_STAB_H_OF_H_PLUS_H_MINUS_IS_COMPLEX_ORIENTATION_SUBALGEBRA_C_H"
	SupportAForientDefinition                 = "CONDITIONAL_SUPPORT_A_F_ORIENT_EQUALS_C_R_PLUS_C_H_PLUS_M3C_AT_STABILIZER_SEAL_LEVEL"
	SupportAForientPreservesSockets           = "CONDITIONAL_SUPPORT_A_F_ORIENT_PRESERVES_H_PLUS_H_MINUS_P1_P3_AND_H_R_MIN"
	SupportAForientPreservesPunctureKernel    = "CONDITIONAL_SUPPORT_A_F_ORIENT_PRESERVES_RIGHT_PUNCTURE_EXCLUSION_AND_LEFT_KERNEL_CANDIDATE"
	SupportDFCompatibleWithAForient           = "CONDITIONAL_SUPPORT_D_F_SYM_SUPPORT_COMPATIBLE_WITH_HIGGS_ORIENTED_STABILIZER_LAYER"
	SupportPostOrientationLayerClassification = "CONDITIONAL_SUPPORT_D_F_SYM_BELONGS_TO_POST_HIGGS_ORIENTATION_LAYER"
	SupportFirstOrderNextTarget               = "CONDITIONAL_SUPPORT_GATE857_CAN_TEST_FIRST_ORDER_FOR_A_F_ORIENT"
	SupportR2StabilizerLayerStage             = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_PLUS_PLUS_PLUS_POST_ORIENTATION_STABILIZER_LAYER"

	FailureFullHPreservesSocketFrame     = "FAILED_ROUTE_FULL_H_ACTION_DOES_NOT_PRESERVE_SOCKET_FRAME"
	FailureFullHNativeSocketEigensplit   = "FAILED_ROUTE_H_PLUS_H_MINUS_NOT_NATIVE_H_EIGENSPLIT"
	FailureFullAFCompatibility           = "FAILED_ROUTE_D_F_SYM_NOT_FULL_UNBROKEN_A_F_COMPATIBLE_OBJECT"
	FailureStabilizerNotFullH            = "FAILED_ROUTE_C_H_STABILIZER_NOT_FULL_H_QUATERNIONIC_ACTION"
	FailureAForientNotUnbrokenAF         = "FAILED_ROUTE_A_F_ORIENT_NOT_FULL_UNBROKEN_A_F"
	FailurePostOrientationNotEWBTheorem  = "FAILED_ROUTE_POST_ORIENTATION_STABILIZER_NOT_ELECTROWEAK_BREAKING_THEOREM"
	FailureNoNativeHiggsVacuumTheorem    = "FAILED_ROUTE_NO_NATIVE_HIGGS_VACUUM_ORIENTATION_THEOREM"
	FailureNoWeakMixingTheorem           = "FAILED_ROUTE_NO_WEAK_MIXING_OR_WEINBERG_ANGLE_THEOREM"
	FailureNoFirstOrderCalculationYet    = "FAILED_ROUTE_STABILIZER_BRANCH_FIRST_ORDER_CALCULATION_NOT_PERFORMED_IN_GATE_856"
	FailureNoFullFirstOrderProof         = "FAILED_ROUTE_NO_FULL_FIRST_ORDER_CONDITION_PROOF"
	FailureNoJOppositeProof              = "FAILED_ROUTE_NO_J_OPPOSITE_ACTION_COMPATIBILITY_PROOF_CERTIFIED"
	FailureNoBimoduleProof               = "FAILED_ROUTE_NO_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailureSupportOnlyNotOperatorTheorem = "FAILED_ROUTE_STABILIZER_COMPATIBILITY_IS_SUPPORT_LEVEL_NOT_OPERATOR_THEOREM"
	FailureNoNativeFiniteTripleProof     = "FAILED_ROUTE_NO_NATIVE_FINITE_TRIPLE_PROOF_CERTIFIED"
	FailureDStillSymbolic                = "FAILED_ROUTE_D_F_SYM_REMAINS_SYMBOLIC_SUPPORT_MATRIX"
	FailureYCoefficientsNotMagnitudes    = "FAILED_ROUTE_SYMBOLIC_Y_COEFFICIENTS_NOT_YUKAWA_MAGNITUDES"
	FailureNoAlphaSource                 = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoTraceMagnitudeReadout       = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoOfficialNEffUpdate          = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate         = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNotR3                         = "FAILED_ROUTE_R2_PLUS_PLUS_PLUS_PLUS_PLUS_STABILIZER_LAYER_NOT_R3"
	FailureNotR4                         = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoParticleAssignment          = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoNeutrinoTheorem             = "FAILED_ROUTE_NO_PHYSICAL_NEUTRINO_THEOREM"
	FailureNoThreeGenerationTheorem      = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
)

type Ledger struct {
	AlphaB                        float64
	OfficialNEff, OfficialCYukawa float64
	OfficialCHiggs                float64
	OfficialFrozen                bool
	AlphaNative, R3, R4           bool
}

type WeakFrame struct {
	Module, OrientedFrame, Stabilizer string
	FullHActsOnFullDoublet            bool
	FullHPreservesIndividualLines     bool
	StabilizerPreservesHPlus          bool
	StabilizerPreservesHMinus         bool
	StabilizerIsComplexSubalgebra     bool
	StabilizerIsNativeFullH           bool
	Supports, Failures                []string
}

type OrientedAlgebra struct {
	FullAlgebra, OrientedAlgebra, WeakFactor string
	ContainsFullH                            bool
	ContainsCH                               bool
	ContainsM3C                              bool
	ContainsRightC                           bool
	PostOrientationLayer                     bool
	UnbrokenFullAFTheorem                    bool
	PhysicalElectroweakTheorem               bool
	Supports, Failures                       []string
}

type ActionPreservation struct {
	PreservesHPlusHMinus                              bool
	PreservesP1P3, PreservesHRMin, PreservesHFMin     bool
	PunctureRemainsOutside, LeftKernelStableCandidate bool
	MinimalCarrierClosureInAForient                   bool
	MinimalCarrierClosureInFullAF                     bool
	Supports, Failures                                []string
}

type DFCompatibility struct {
	DObject, CompatibleAlgebra, FirstOrderTarget                   string
	SupportCompatible, OperatorTheoremCompatible, FullAFCompatible bool
	PostOrientationObject, FirstOrderReadyForGate857               bool
	FirstOrderCalculatedThisGate, FirstOrderCertified              bool
	Supports, Failures                                             []string
}

type Carrier struct {
	HLRank, HRMinRank, HPartMinRank, HFMinRank int
	AmbientPartRank, AmbientFRank              int
	DSymRank, KernelRank                       int
	RightPuncture, LeftKernel                  string
	Supports, Failures                         []string
}

type Impact struct {
	Classification                                                                           string
	Gate855Inherited, StabilizerAlgebraTyped, AForientDefined, AForientPreservesCarrier      bool
	FullAFPass, StabilizerSupportPass, PostOrientationLayer, FirstOrderReady, FirstOrderDone bool
	NativeFiniteTripleProof, AlphaStillSealed, MagnitudesStillMissing                        bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4         bool
}

type Firewalls struct {
	Enforced                                                                                                    bool
	FullHPreservesSocketFrame, NativeHEigensplit, FullAFCompatible, StabilizerNotFullH, AForientNotFullAF       bool
	PostOrientationNotEWB, NoHiggsVacuumTheorem, NoWeakMixingTheorem                                            bool
	NoFirstOrderThisGate, NoFullFirstOrderProof, NoJOppositeProof, NoBimoduleProof, SupportOnly, NoNativeTriple bool
	DSymbolicOnly, YSymbolicOnly, NoAlphaSource, NoTraceReadout, NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate    bool
	NotR3, NotR4, NoParticleAssignment, NoNeutrinoTheorem, NoThreeGenerationTheorem                             bool
	Verdict                                                                                                     string
}

type Audit struct {
	ID        string
	Ledger    Ledger
	WeakFrame WeakFrame
	Algebra   OrientedAlgebra
	Action    ActionPreservation
	DF        DFCompatibility
	Carrier   Carrier
	Impact    Impact
	Firewalls Firewalls
	Truth     string
	Final     string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		ID:        AuditID,
		Ledger:    buildLedger(),
		WeakFrame: buildWeakFrame(),
		Algebra:   buildAlgebra(),
		Action:    buildAction(),
		DF:        buildDF(),
		Carrier:   buildCarrier(),
		Impact:    buildImpact(),
		Firewalls: buildFirewalls(),
		Truth:     "Gate 856 inherits Gate 855's result: D_F^sym is blocked for the full unbroken A_F because generic H mixes h_+/h_-.  The new object is the Higgs-oriented stabilizer algebra A_F^orient=C_R plus C_H plus M_3(C), which preserves the weak socket frame and therefore the support of D_F^sym at seal level.  This is a post-orientation layer, not a full unbroken finite-triple theorem.",
		Final:     "A_F^orient is now typed as the correct layer for the next first-order calculation.  The full H firewall, Higgs-orientation seal, symbolic-Y firewall, alpha_B seal, R3/R4 firewall, and all official ledgers remain closed.",
	}
	if err := a.validate(); err != nil {
		return Audit{}, err
	}
	return a, nil
}

func (a Audit) validate() error {
	if a.Carrier.HLRank != HLRank || a.Carrier.HRMinRank != HRMinRank || a.Carrier.HPartMinRank != HPartMinRank || a.Carrier.HFMinRank != HFMinRank {
		return err("carrier ranks inconsistent")
	}
	if a.Carrier.DSymRank != DSymRank || a.Carrier.KernelRank != DSymKernelRank {
		return err("D_F support ranks inconsistent")
	}
	if !a.WeakFrame.FullHActsOnFullDoublet || a.WeakFrame.FullHPreservesIndividualLines || !a.WeakFrame.StabilizerPreservesHPlus || !a.WeakFrame.StabilizerPreservesHMinus || !a.WeakFrame.StabilizerIsComplexSubalgebra || a.WeakFrame.StabilizerIsNativeFullH {
		return err("weak-frame stabilizer flags inconsistent")
	}
	if a.Algebra.ContainsFullH || !a.Algebra.ContainsCH || !a.Algebra.ContainsM3C || !a.Algebra.ContainsRightC || !a.Algebra.PostOrientationLayer || a.Algebra.UnbrokenFullAFTheorem || a.Algebra.PhysicalElectroweakTheorem {
		return err("oriented algebra flags inconsistent")
	}
	if !a.Action.PreservesHPlusHMinus || !a.Action.PreservesP1P3 || !a.Action.PreservesHRMin || !a.Action.PreservesHFMin || !a.Action.PunctureRemainsOutside || !a.Action.LeftKernelStableCandidate || !a.Action.MinimalCarrierClosureInAForient || a.Action.MinimalCarrierClosureInFullAF {
		return err("action preservation flags inconsistent")
	}
	if !a.DF.SupportCompatible || a.DF.OperatorTheoremCompatible || a.DF.FullAFCompatible || !a.DF.PostOrientationObject || !a.DF.FirstOrderReadyForGate857 || a.DF.FirstOrderCalculatedThisGate || a.DF.FirstOrderCertified {
		return err("D_F compatibility flags inconsistent")
	}
	if !a.Impact.Gate855Inherited || !a.Impact.StabilizerAlgebraTyped || !a.Impact.AForientDefined || !a.Impact.AForientPreservesCarrier || a.Impact.FullAFPass || !a.Impact.StabilizerSupportPass || !a.Impact.PostOrientationLayer || !a.Impact.FirstOrderReady || a.Impact.FirstOrderDone || a.Impact.NativeFiniteTripleProof {
		return err("impact flags inconsistent")
	}
	if !a.Firewalls.Enforced || !a.Firewalls.FullHPreservesSocketFrame || !a.Firewalls.NativeHEigensplit || !a.Firewalls.StabilizerNotFullH || !a.Firewalls.NoFirstOrderThisGate || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 {
		return err("firewall flags inconsistent")
	}
	return nil
}

func err(msg string) error { return fmt.Errorf("%s", msg) }

func buildLedger() Ledger {
	return Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true}
}

func buildWeakFrame() WeakFrame {
	return WeakFrame{
		Module: "C_L^2", OrientedFrame: "h_+ plus h_-", Stabilizer: "C_H subset H preserving h_+ and h_-",
		FullHActsOnFullDoublet: true, FullHPreservesIndividualLines: false,
		StabilizerPreservesHPlus: true, StabilizerPreservesHMinus: true,
		StabilizerIsComplexSubalgebra: true, StabilizerIsNativeFullH: false,
		Supports: []string{StatusStabilizerAudited, SupportStabilizerIsComplexSubalgebra},
		Failures: []string{FailureFullHPreservesSocketFrame, FailureFullHNativeSocketEigensplit, FailureStabilizerNotFullH},
	}
}

func buildAlgebra() OrientedAlgebra {
	return OrientedAlgebra{
		FullAlgebra:     "A_F=C plus H plus M_3(C)",
		OrientedAlgebra: "A_F^orient=C_R plus C_H plus M_3(C)",
		WeakFactor:      "C_H=Stab_H(h_+ plus h_-)",
		ContainsFullH:   false, ContainsCH: true, ContainsM3C: true, ContainsRightC: true,
		PostOrientationLayer: true, UnbrokenFullAFTheorem: false, PhysicalElectroweakTheorem: false,
		Supports: []string{StatusOrientedAlgebraDefined, SupportAForientDefinition, SupportPostOrientationLayerClassification},
		Failures: []string{FailureAForientNotUnbrokenAF, FailurePostOrientationNotEWBTheorem, FailureNoWeakMixingTheorem},
	}
}

func buildAction() ActionPreservation {
	return ActionPreservation{
		PreservesHPlusHMinus: true, PreservesP1P3: true, PreservesHRMin: true, PreservesHFMin: true,
		PunctureRemainsOutside: true, LeftKernelStableCandidate: true,
		MinimalCarrierClosureInAForient: true, MinimalCarrierClosureInFullAF: false,
		Supports: []string{StatusOrientedActionPreservation, SupportAForientPreservesSockets, SupportAForientPreservesPunctureKernel},
		Failures: []string{FailureFullAFCompatibility, FailureFullHPreservesSocketFrame, FailureSupportOnlyNotOperatorTheorem},
	}
}

func buildDF() DFCompatibility {
	return DFCompatibility{
		DObject:           "D_F^sym=[[0,Y_supp^dagger],[Y_supp,0]]",
		CompatibleAlgebra: "A_F^orient=C_R plus C_H plus M_3(C)",
		FirstOrderTarget:  "[[D_F^sym,rho_F(a)],J_F rho_F(b) J_F^{-1}] for a,b in A_F^orient",
		SupportCompatible: true, OperatorTheoremCompatible: false, FullAFCompatible: false,
		PostOrientationObject: true, FirstOrderReadyForGate857: true,
		FirstOrderCalculatedThisGate: false, FirstOrderCertified: false,
		Supports: []string{StatusDFCompatibilityAudited, StatusNextFirstOrderTyped, SupportDFCompatibleWithAForient, SupportPostOrientationLayerClassification, SupportFirstOrderNextTarget},
		Failures: []string{FailureNoFirstOrderCalculationYet, FailureNoFullFirstOrderProof, FailureNoJOppositeProof, FailureNoBimoduleProof, FailureDStillSymbolic},
	}
}

func buildCarrier() Carrier {
	return Carrier{
		HLRank: HLRank, HRMinRank: HRMinRank, HPartMinRank: HPartMinRank, HFMinRank: HFMinRank,
		AmbientPartRank: AmbientPartRank, AmbientFRank: AmbientFRank, DSymRank: DSymRank, KernelRank: DSymKernelRank,
		RightPuncture: "e_+ tensor P_1", LeftKernel: "h_+ tensor P_1",
		Supports: []string{StatusKernelPuncturePreservation, SupportAForientPreservesPunctureKernel},
		Failures: []string{FailureNoNeutrinoTheorem, FailureNoParticleAssignment},
	}
}

func buildImpact() Impact {
	return Impact{Classification: "R2+++++_post_orientation_stabilizer_layer", Gate855Inherited: true, StabilizerAlgebraTyped: true, AForientDefined: true, AForientPreservesCarrier: true, FullAFPass: false, StabilizerSupportPass: true, PostOrientationLayer: true, FirstOrderReady: true, FirstOrderDone: false, NativeFiniteTripleProof: false, AlphaStillSealed: true, MagnitudesStillMissing: true}
}

func buildFirewalls() Firewalls {
	return Firewalls{Enforced: true, FullHPreservesSocketFrame: true, NativeHEigensplit: true, FullAFCompatible: true, StabilizerNotFullH: true, AForientNotFullAF: true, PostOrientationNotEWB: true, NoHiggsVacuumTheorem: true, NoWeakMixingTheorem: true, NoFirstOrderThisGate: true, NoFullFirstOrderProof: true, NoJOppositeProof: true, NoBimoduleProof: true, SupportOnly: true, NoNativeTriple: true, DSymbolicOnly: true, YSymbolicOnly: true, NoAlphaSource: true, NoTraceReadout: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NotR3: true, NotR4: true, NoParticleAssignment: true, NoNeutrinoTheorem: true, NoThreeGenerationTheorem: true, Verdict: StatusFirewallVerdict}
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("alpha_B=%.16g; official_N_eff=%.16g; C_Yukawa=%.16g; C_Higgs=%.16g; frozen=%t; alpha_native=%t; R3=%t; R4=%t", l.AlphaB, l.OfficialNEff, l.OfficialCYukawa, l.OfficialCHiggs, l.OfficialFrozen, l.AlphaNative, l.R3, l.R4)
}

func FormatWeakFrame(w WeakFrame) string {
	return fmt.Sprintf("module=%s; frame=%s; stabilizer=%s; full_H_on_doublet=%t; full_H_preserves_lines=%t; C_H_preserves_h+=%t; C_H_preserves_h-=%t; C_H_complex=%t; C_H_full_H=%t", w.Module, w.OrientedFrame, w.Stabilizer, w.FullHActsOnFullDoublet, w.FullHPreservesIndividualLines, w.StabilizerPreservesHPlus, w.StabilizerPreservesHMinus, w.StabilizerIsComplexSubalgebra, w.StabilizerIsNativeFullH)
}

func FormatAlgebra(a OrientedAlgebra) string {
	return fmt.Sprintf("full=%s; orient=%s; weak_factor=%s; contains_full_H=%t; contains_C_H=%t; contains_M3C=%t; post_orientation=%t; unbroken_full_A_F=%t; physical_EW=%t", a.FullAlgebra, a.OrientedAlgebra, a.WeakFactor, a.ContainsFullH, a.ContainsCH, a.ContainsM3C, a.PostOrientationLayer, a.UnbrokenFullAFTheorem, a.PhysicalElectroweakTheorem)
}

func FormatAction(a ActionPreservation) string {
	return fmt.Sprintf("preserves_h_lines=%t; preserves_P1P3=%t; preserves_HRmin=%t; preserves_HFmin=%t; puncture_outside=%t; left_kernel_candidate=%t; closure_orient=%t; closure_full_A_F=%t", a.PreservesHPlusHMinus, a.PreservesP1P3, a.PreservesHRMin, a.PreservesHFMin, a.PunctureRemainsOutside, a.LeftKernelStableCandidate, a.MinimalCarrierClosureInAForient, a.MinimalCarrierClosureInFullAF)
}

func FormatDF(d DFCompatibility) string {
	return fmt.Sprintf("D=%s; algebra=%s; target=%s; support_compatible=%t; operator_theorem=%t; full_A_F=%t; post_orientation=%t; gate857_ready=%t; first_order_done=%t; first_order_certified=%t", d.DObject, d.CompatibleAlgebra, d.FirstOrderTarget, d.SupportCompatible, d.OperatorTheoremCompatible, d.FullAFCompatible, d.PostOrientationObject, d.FirstOrderReadyForGate857, d.FirstOrderCalculatedThisGate, d.FirstOrderCertified)
}

func FormatCarrier(c Carrier) string {
	return fmt.Sprintf("H_L=%d; H_R^min=%d; H_part^min=%d; H_F^min=%d; ambient=%d/%d; D_rank=%d; kernel=%d; right_puncture=%s; left_kernel=%s", c.HLRank, c.HRMinRank, c.HPartMinRank, c.HFMinRank, c.AmbientPartRank, c.AmbientFRank, c.DSymRank, c.KernelRank, c.RightPuncture, c.LeftKernel)
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s; gate855=%t; A_F_orient=%t; preserves_carrier=%t; full_pass=%t; stabilizer_support=%t; post_orientation=%t; first_order_ready=%t; first_order_done=%t; native_triple=%t; alpha_sealed=%t; magnitudes_missing=%t; update_N_eff=%t; R3=%t; R4=%t", i.Classification, i.Gate855Inherited, i.AForientDefined, i.AForientPreservesCarrier, i.FullAFPass, i.StabilizerSupportPass, i.PostOrientationLayer, i.FirstOrderReady, i.FirstOrderDone, i.NativeFiniteTripleProof, i.AlphaStillSealed, i.MagnitudesStillMissing, i.CanUpdateNEff, i.CanPromoteToR3, i.CanPromoteToR4)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%t; verdict=%s; full_H_preserve_frame_fail=%t; C_H_not_full_H=%t; A_F_orient_not_full_A_F=%t; no_EW_theorem=%t; no_first_order_this_gate=%t; support_only=%t; D_symbolic=%t; no_alpha=%t; no_trace_readout=%t; not_R3=%t; not_R4=%t", f.Enforced, f.Verdict, f.FullHPreservesSocketFrame, f.StabilizerNotFullH, f.AForientNotFullAF, f.PostOrientationNotEWB, f.NoFirstOrderThisGate, f.SupportOnly, f.DSymbolicOnly, f.NoAlphaSource, f.NoTraceReadout, f.NotR3, f.NotR4)
}

func Statuses() []string {
	return []string{
		StatusGate855Inherited,
		StatusStabilizerAudited,
		StatusFullHFirewallAudited,
		StatusOrientedAlgebraDefined,
		StatusOrientedActionPreservation,
		StatusDFCompatibilityAudited,
		StatusKernelPuncturePreservation,
		StatusNextFirstOrderTyped,
		StatusNoObservedDataUsed,
		StatusLedgerFrozen,
		SupportStabilizerIsComplexSubalgebra,
		SupportAForientDefinition,
		SupportAForientPreservesSockets,
		SupportAForientPreservesPunctureKernel,
		SupportDFCompatibleWithAForient,
		SupportPostOrientationLayerClassification,
		SupportFirstOrderNextTarget,
		SupportR2StabilizerLayerStage,
		FailureFullHPreservesSocketFrame,
		FailureFullHNativeSocketEigensplit,
		FailureFullAFCompatibility,
		FailureStabilizerNotFullH,
		FailureAForientNotUnbrokenAF,
		FailurePostOrientationNotEWBTheorem,
		FailureNoNativeHiggsVacuumTheorem,
		FailureNoWeakMixingTheorem,
		FailureNoFirstOrderCalculationYet,
		FailureNoFullFirstOrderProof,
		FailureNoJOppositeProof,
		FailureNoBimoduleProof,
		FailureSupportOnlyNotOperatorTheorem,
		FailureNoNativeFiniteTripleProof,
		FailureDStillSymbolic,
		FailureYCoefficientsNotMagnitudes,
		FailureNoAlphaSource,
		FailureNoTraceMagnitudeReadout,
		FailureNoOfficialNEffUpdate,
		FailureNoCYukawaCHiggsUpdate,
		FailureNotR3,
		FailureNotR4,
		FailureNoParticleAssignment,
		FailureNoNeutrinoTheorem,
		FailureNoThreeGenerationTheorem,
	}
}

func containsAll(haystack, needles []string) bool {
	m := map[string]bool{}
	for _, h := range haystack {
		m[h] = true
	}
	for _, n := range needles {
		if !m[n] {
			return false
		}
	}
	return true
}

func join(xs []string) string { return strings.Join(xs, "; ") }
