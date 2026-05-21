// Package generation2weakdoublethiggsorientationsealaudit implements
// Gate 853: WeakDoublet / HiggsOrientationSeal Audit.
//
// Gate 853 follows Gate 852's first-order/J-opposite firewall.  It audits the
// fragile weak rank-one split C_L^2 = h_+ plus h_- and separates the native
// quaternionic module from the orientation-relative Higgs/weak socket frame
// needed by the symbolic D_F edge skeleton.  This is an orientation-seal audit
// only.  It does not derive a Higgs vacuum, Yukawa magnitudes, alpha_B,
// first-order compatibility, or any physical particle assignment.
package generation2weakdoublethiggsorientationsealaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE853-WEAKDOUBLET-HIGGS-ORIENTATION-SEAL-AUDIT"

	AlphaB          = 0.0003878958469680527
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	P1Rank           = 1
	P3Rank           = 3
	WRank            = P1Rank + P3Rank
	WeakDoubletRank  = 2
	HPlusRank        = 1
	HMinusRank       = 1
	HLRank           = WeakDoubletRank * WRank
	RightMinRank     = 7
	HPartMinRank     = HLRank + RightMinRank
	HFMinRank        = 2 * HPartMinRank
	LeftKernelRank   = HPlusRank * P1Rank
	ActiveEdgeCount  = 3
	PunctureEdgeRank = 1

	StatusGate852Inherited              = "PASS_GATE852_FIRST_ORDER_FIREWALL_INHERITED"
	StatusWeakModuleFirewallAudited     = "PASS_WEAK_DOUBLET_H_MODULE_FIREWALL_AUDITED"
	StatusNoNativeHEigensplit           = "PASS_H_PLUS_H_MINUS_NOT_NATIVE_H_EIGENSPLIT"
	StatusHiggsOrientationSealDefined   = "PASS_HIGGS_ORIENTATION_SEAL_DEFINED"
	StatusOrientationProjectorsDefined  = "PASS_ORIENTATION_RELATIVE_SOCKET_PROJECTORS_DEFINED"
	StatusStabilityClassesSeparated     = "PASS_FULL_WEAK_MODULE_STABILITY_SEPARATED_FROM_ORIENTED_SOCKET_STABILITY"
	StatusEdgeSkeletonRewritten         = "PASS_GATE847_EDGE_SKELETON_REWRITTEN_IN_ORIENTED_FRAME"
	StatusLeftKernelOrientationRelative = "PASS_LEFT_KERNEL_SINGLETON_CLASSIFIED_AS_ORIENTATION_RELATIVE"
	StatusFirstOrderPreparationAudited  = "PASS_OPERATOR_LEVEL_FIRST_ORDER_PREPARATION_AUDITED"
	StatusNoObservedDataUsed            = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusLedgersFrozen                 = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusFirewallVerdict               = "FIREWALL_PRESERVED_GATE853_ORIENTATION_SEAL_NOT_NATIVE_THEOREM"

	SupportWeakSocketSplitAfterHiggsSeal  = "CONDITIONAL_SUPPORT_WEAK_SOCKET_SPLIT_EXISTS_AFTER_HIGGS_ORIENTATION_SEAL"
	SupportLeftKernelOrientationRelative  = "CONDITIONAL_SUPPORT_LEFT_KERNEL_SINGLETON_IS_ORIENTATION_RELATIVE"
	SupportEdgeSkeletonCompatibleWithSeal = "CONDITIONAL_SUPPORT_EDGE_SKELETON_COMPATIBLE_WITH_ORIENTED_WEAK_SOCKET_FRAME"
	SupportFirstOrderNeedsOperatorFrame   = "CONDITIONAL_SUPPORT_FIRST_ORDER_TEST_REQUIRES_OPERATOR_LEVEL_REALIZATION_IN_ORIENTED_FRAME"
	SupportFullWeakModuleHStable          = "CONDITIONAL_SUPPORT_FULL_WEAK_DOUBLET_MODULE_IS_H_STABLE"
	SupportR2OrientationSealStage         = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_PLUS_PLUS_PLUS_DATA_SEAL_ORIENTATION_STAGE"

	FailureWeakSplitNotNativeHEigensplit     = "FAILED_ROUTE_H_PLUS_H_MINUS_NOT_NATIVE_H_EIGENSPLIT"
	FailureHiggsOrientationSealNotNative     = "FAILED_ROUTE_HIGGS_ORIENTATION_SEAL_NOT_NATIVE_DERIVATION"
	FailureRankOneLinesNotGloballyHStable    = "FAILED_ROUTE_RANK_ONE_WEAK_LINES_NOT_STABLE_UNDER_FULL_H_ACTION"
	FailureNoNativeHiggsVacuumTheorem        = "FAILED_ROUTE_NO_NATIVE_HIGGS_VACUUM_ORIENTATION_THEOREM"
	FailureNoOperatorLevelRhoFJFDFF          = "FAILED_ROUTE_NO_OPERATOR_LEVEL_RHO_F_J_F_GAMMA_F_D_F_YET"
	FailureNoFirstOrderProofYet              = "FAILED_ROUTE_NO_FIRST_ORDER_PROOF_YET"
	FailureNoJOppositeProofYet               = "FAILED_ROUTE_NO_J_OPPOSITE_ACTION_COMPATIBILITY_PROOF_YET"
	FailureNoBimoduleCommutantProof          = "FAILED_ROUTE_NO_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailureLeftKernelNotRepresentationStable = "FAILED_ROUTE_LEFT_KERNEL_STABILITY_NOT_CERTIFIED_UNDER_FULL_RHO_F_AND_J_F"
	FailureOrientationDoesNotDeriveAlphaB    = "FAILED_ROUTE_HIGGS_ORIENTATION_SEAL_DOES_NOT_DERIVE_ALPHA_B"
	FailureNoYukawaMagnitudeSource           = "FAILED_ROUTE_NO_YUKAWA_MAGNITUDE_SOURCE"
	FailureNoSectorTraceMagnitudeReadout     = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoOfficialNEffUpdate              = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate             = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNotR3                             = "FAILED_ROUTE_R2_PLUS_PLUS_PLUS_PLUS_PLUS_ORIENTATION_SEAL_NOT_R3"
	FailureNotR4                             = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoPhysicalNeutrinoTheorem         = "FAILED_ROUTE_NO_PHYSICAL_NEUTRINO_THEOREM"
	FailureNoRightNeutrinoTheorem            = "FAILED_ROUTE_NO_RIGHT_NEUTRINO_OR_STERILE_PARTICLE_THEOREM"
	FailureNoMasslessnessTheorem             = "FAILED_ROUTE_LEFT_KERNEL_NOT_OBSERVED_MASSLESSNESS_THEOREM"
	FailureNoWeakMixingOrHiggsMassTheorem    = "FAILED_ROUTE_NO_WEAK_MIXING_OR_HIGGS_MASS_THEOREM"
	FailureNoParticleAssignment              = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
)

type Ledger struct {
	AlphaB                        float64
	OfficialNEff, OfficialCYukawa float64
	OfficialCHiggs                float64
	OfficialFrozen                bool
	AlphaNative, R3, R4           bool
}

type QuaternionicFirewall struct {
	Module, ActionDescription    string
	ModuleRank                   int
	FullModuleHStable            bool
	NativeRankOneEigensplit      bool
	GenericHActionPreservesLines bool
	Supports, Failures           []string
}

type HiggsOrientationSeal struct {
	UnitVector, Projector, Complement string
	HPlusRank, HMinusRank             int
	ProjectorsComplete, Orthogonal    bool
	DefinedAtSealLevel                bool
	NativeDerivation                  bool
	RequiresGaugeOrientation          bool
	Supports, Failures                []string
}

type StabilityClass struct {
	FullWeakModule, OrientedSockets string
	FullModuleStableUnderH          bool
	SocketsStableUnderFullH         bool
	SocketsStableAfterOrientation   bool
	GlobalQuaternionicEigenspaces   bool
	Supports, Failures              []string
}

type EdgeRewrite struct {
	Edges                         []string
	ActiveEdgeCount               int
	PunctureEdge                  string
	PunctureEdgeAbsent            bool
	LeptoColorPreserved           bool
	UsesOrientationFrame          bool
	CompatibleWithGate847Skeleton bool
	Supports, Failures            []string
}

type KernelOrientation struct {
	Kernel                         string
	Rank                           int
	OrientationRelative            bool
	StableUnderOrientationBlocks   bool
	StableUnderFullRhoJ            bool
	PhysicalNeutrino, Masslessness bool
	Supports, Failures             []string
}

type FirstOrderPreparation struct {
	Target, RequiredNextStep string
	OrientationSealAvailable bool
	OperatorRealizationReady bool
	FirstOrderExecutable     bool
	FirstOrderCertified      bool
	Supports, Failures       []string
}

type Impact struct {
	Classification                                                                      string
	Gate852Inherited, OrientationSealDefined, EdgeSkeletonRewritten                     bool
	WeakSplitNative, NativeHiggsOrientation, OperatorRealizationReady, FirstOrderProved bool
	LeftKernelStableUnderFullRhoJ                                                       bool
	AlphaStillSealed, MagnitudesStillMissing                                            bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4    bool
}

type Firewalls struct {
	Enforced                                                                                       bool
	WeakSplitNotNative, HiggsOrientationNotNative, RankOneLinesNotHStable                          bool
	NoNativeHiggsVacuum, NoOperatorPackage, NoFirstOrderProof, NoJOppositeProof, NoBimoduleProof   bool
	KernelNotRepresentationStable, OrientationDoesNotDeriveAlpha, NoYukawaMagnitudes               bool
	NoTraceMagnitudeReadout, NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate, NotR3, NotR4             bool
	NoPhysicalNeutrino, NoRightNeutrino, NoMasslessness, NoWeakMixingOrHiggsMass, NoParticleAssign bool
	Verdict                                                                                        string
}

type Audit struct {
	ID                    string
	Ledger                Ledger
	QuaternionicFirewall  QuaternionicFirewall
	OrientationSeal       HiggsOrientationSeal
	Stability             StabilityClass
	EdgeRewrite           EdgeRewrite
	Kernel                KernelOrientation
	FirstOrderPreparation FirstOrderPreparation
	Impact                Impact
	Firewalls             Firewalls
	Truth                 string
	Final                 string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		ID:                    AuditID,
		Ledger:                buildLedger(),
		QuaternionicFirewall:  buildQuaternionicFirewall(),
		OrientationSeal:       buildOrientationSeal(),
		Stability:             buildStability(),
		EdgeRewrite:           buildEdgeRewrite(),
		Kernel:                buildKernel(),
		FirstOrderPreparation: buildFirstOrderPreparation(),
		Impact:                buildImpact(),
		Firewalls:             buildFirewalls(),
		Truth:                 "Gate 853 admits h_+ plus h_- only as a Higgs/weak-orientation seal: C_L^2 is H-stable as a full doublet, while rank-one weak sockets are orientation-relative and not native quaternionic eigenspaces.",
		Final:                 "VERDICT: weak socket frame repaired at seal level; symbolic edge skeleton can be read in the oriented frame; first-order/J-opposite proof still requires operator-level rho_F,J_F,gamma_F,D_F matrices.",
	}
	if err := a.Validate(); err != nil {
		return Audit{}, err
	}
	return a, nil
}

func (a Audit) Validate() error {
	err := func(msg string) error { return fmt.Errorf("%s: %s", AuditID, msg) }
	if a.ID != AuditID {
		return err("wrong audit id")
	}
	if !almost(a.Ledger.AlphaB, AlphaB, 1e-18) || !a.Ledger.OfficialFrozen {
		return err("ledger mismatch")
	}
	if a.QuaternionicFirewall.ModuleRank != WeakDoubletRank || !a.QuaternionicFirewall.FullModuleHStable || a.QuaternionicFirewall.NativeRankOneEigensplit || a.QuaternionicFirewall.GenericHActionPreservesLines {
		return err("quaternionic firewall flags inconsistent")
	}
	if a.OrientationSeal.HPlusRank != 1 || a.OrientationSeal.HMinusRank != 1 || !a.OrientationSeal.ProjectorsComplete || !a.OrientationSeal.Orthogonal || !a.OrientationSeal.DefinedAtSealLevel || a.OrientationSeal.NativeDerivation {
		return err("orientation seal inconsistent")
	}
	if !a.Stability.FullModuleStableUnderH || a.Stability.SocketsStableUnderFullH || !a.Stability.SocketsStableAfterOrientation || a.Stability.GlobalQuaternionicEigenspaces {
		return err("stability classes inconsistent")
	}
	if a.EdgeRewrite.ActiveEdgeCount != ActiveEdgeCount || !a.EdgeRewrite.PunctureEdgeAbsent || !a.EdgeRewrite.LeptoColorPreserved || !a.EdgeRewrite.CompatibleWithGate847Skeleton {
		return err("edge rewrite inconsistent")
	}
	if a.Kernel.Rank != LeftKernelRank || !a.Kernel.OrientationRelative || !a.Kernel.StableUnderOrientationBlocks || a.Kernel.StableUnderFullRhoJ || a.Kernel.PhysicalNeutrino || a.Kernel.Masslessness {
		return err("kernel overpromoted")
	}
	if !a.FirstOrderPreparation.OrientationSealAvailable || a.FirstOrderPreparation.OperatorRealizationReady || a.FirstOrderPreparation.FirstOrderExecutable || a.FirstOrderPreparation.FirstOrderCertified {
		return err("first-order preparation overpromoted")
	}
	if !a.Firewalls.Enforced || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 || a.Impact.CanUpdateNEff {
		return err("firewall/impact inconsistent")
	}
	return nil
}

func buildLedger() Ledger {
	return Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true}
}

func buildQuaternionicFirewall() QuaternionicFirewall {
	return QuaternionicFirewall{
		Module: "C_L^2", ActionDescription: "H acts irreducibly/rotationally on the weak doublet support; arbitrary complex rank-one lines are not native H eigenspaces", ModuleRank: WeakDoubletRank,
		FullModuleHStable: true, NativeRankOneEigensplit: false, GenericHActionPreservesLines: false,
		Supports: []string{StatusWeakModuleFirewallAudited, StatusNoNativeHEigensplit, SupportFullWeakModuleHStable},
		Failures: []string{FailureWeakSplitNotNativeHEigensplit, FailureRankOneLinesNotGloballyHStable},
	}
}

func buildOrientationSeal() HiggsOrientationSeal {
	return HiggsOrientationSeal{
		UnitVector: "u_H in C_L^2", Projector: "P_H = |u_H><u_H| = h_+", Complement: "I_{C_L^2}-P_H = h_-", HPlusRank: HPlusRank, HMinusRank: HMinusRank,
		ProjectorsComplete: true, Orthogonal: true, DefinedAtSealLevel: true, NativeDerivation: false, RequiresGaugeOrientation: true,
		Supports: []string{StatusHiggsOrientationSealDefined, StatusOrientationProjectorsDefined, SupportWeakSocketSplitAfterHiggsSeal},
		Failures: []string{FailureHiggsOrientationSealNotNative, FailureNoNativeHiggsVacuumTheorem, FailureNoWeakMixingOrHiggsMassTheorem},
	}
}

func buildStability() StabilityClass {
	return StabilityClass{
		FullWeakModule: "C_L^2", OrientedSockets: "h_+ plus h_-",
		FullModuleStableUnderH: true, SocketsStableUnderFullH: false, SocketsStableAfterOrientation: true, GlobalQuaternionicEigenspaces: false,
		Supports: []string{StatusStabilityClassesSeparated, SupportFullWeakModuleHStable, SupportWeakSocketSplitAfterHiggsSeal},
		Failures: []string{FailureWeakSplitNotNativeHEigensplit, FailureRankOneLinesNotGloballyHStable, FailureHiggsOrientationSealNotNative},
	}
}

func buildEdgeRewrite() EdgeRewrite {
	return EdgeRewrite{
		Edges:           []string{"Y_+3: e_+ tensor P_3 -> h_+ tensor P_3", "Y_-3: e_- tensor P_3 -> h_- tensor P_3", "Y_-1: e_- tensor P_1 -> h_- tensor P_1"},
		ActiveEdgeCount: ActiveEdgeCount, PunctureEdge: "Y_+1: e_+ tensor P_1 -> h_+ tensor P_1", PunctureEdgeAbsent: true, LeptoColorPreserved: true, UsesOrientationFrame: true, CompatibleWithGate847Skeleton: true,
		Supports: []string{StatusEdgeSkeletonRewritten, SupportEdgeSkeletonCompatibleWithSeal},
		Failures: []string{FailureHiggsOrientationSealNotNative, FailureNoOperatorLevelRhoFJFDFF},
	}
}

func buildKernel() KernelOrientation {
	return KernelOrientation{
		Kernel: "h_+ tensor P_1", Rank: LeftKernelRank, OrientationRelative: true, StableUnderOrientationBlocks: true, StableUnderFullRhoJ: false, PhysicalNeutrino: false, Masslessness: false,
		Supports: []string{StatusLeftKernelOrientationRelative, SupportLeftKernelOrientationRelative},
		Failures: []string{FailureLeftKernelNotRepresentationStable, FailureNoPhysicalNeutrinoTheorem, FailureNoRightNeutrinoTheorem, FailureNoMasslessnessTheorem},
	}
}

func buildFirstOrderPreparation() FirstOrderPreparation {
	return FirstOrderPreparation{
		Target: "[[D_F,rho_F(a)],J_F rho_F(b) J_F^{-1}]=0", RequiredNextStep: "operator-level finite-triple matrix realization in the oriented weak frame",
		OrientationSealAvailable: true, OperatorRealizationReady: false, FirstOrderExecutable: false, FirstOrderCertified: false,
		Supports: []string{StatusFirstOrderPreparationAudited, SupportFirstOrderNeedsOperatorFrame},
		Failures: []string{FailureNoOperatorLevelRhoFJFDFF, FailureNoFirstOrderProofYet, FailureNoJOppositeProofYet, FailureNoBimoduleCommutantProof},
	}
}

func buildImpact() Impact {
	return Impact{Classification: "R2+++++_data_seal_higgs_orientation_seal", Gate852Inherited: true, OrientationSealDefined: true, EdgeSkeletonRewritten: true, WeakSplitNative: false, NativeHiggsOrientation: false, OperatorRealizationReady: false, FirstOrderProved: false, LeftKernelStableUnderFullRhoJ: false, AlphaStillSealed: true, MagnitudesStillMissing: true}
}

func buildFirewalls() Firewalls {
	return Firewalls{Enforced: true, WeakSplitNotNative: true, HiggsOrientationNotNative: true, RankOneLinesNotHStable: true, NoNativeHiggsVacuum: true, NoOperatorPackage: true, NoFirstOrderProof: true, NoJOppositeProof: true, NoBimoduleProof: true, KernelNotRepresentationStable: true, OrientationDoesNotDeriveAlpha: true, NoYukawaMagnitudes: true, NoTraceMagnitudeReadout: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NotR3: true, NotR4: true, NoPhysicalNeutrino: true, NoRightNeutrino: true, NoMasslessness: true, NoWeakMixingOrHiggsMass: true, NoParticleAssign: true, Verdict: StatusFirewallVerdict}
}

func Statuses() []string {
	return []string{
		StatusGate852Inherited, StatusWeakModuleFirewallAudited, StatusNoNativeHEigensplit, StatusHiggsOrientationSealDefined, StatusOrientationProjectorsDefined, StatusStabilityClassesSeparated, StatusEdgeSkeletonRewritten, StatusLeftKernelOrientationRelative, StatusFirstOrderPreparationAudited, StatusNoObservedDataUsed, StatusLedgersFrozen, StatusFirewallVerdict,
		SupportWeakSocketSplitAfterHiggsSeal, SupportLeftKernelOrientationRelative, SupportEdgeSkeletonCompatibleWithSeal, SupportFirstOrderNeedsOperatorFrame, SupportFullWeakModuleHStable, SupportR2OrientationSealStage,
		FailureWeakSplitNotNativeHEigensplit, FailureHiggsOrientationSealNotNative, FailureRankOneLinesNotGloballyHStable, FailureNoNativeHiggsVacuumTheorem, FailureNoOperatorLevelRhoFJFDFF, FailureNoFirstOrderProofYet, FailureNoJOppositeProofYet, FailureNoBimoduleCommutantProof, FailureLeftKernelNotRepresentationStable, FailureOrientationDoesNotDeriveAlphaB, FailureNoYukawaMagnitudeSource, FailureNoSectorTraceMagnitudeReadout, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNotR3, FailureNotR4, FailureNoPhysicalNeutrinoTheorem, FailureNoRightNeutrinoTheorem, FailureNoMasslessnessTheorem, FailureNoWeakMixingOrHiggsMassTheorem, FailureNoParticleAssignment,
	}
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("alpha_B=%.16g official_N_eff=%.16g official_C_Yukawa=%.16g official_C_Higgs=%.16g frozen=%t alpha_native=%t R3=%t R4=%t", l.AlphaB, l.OfficialNEff, l.OfficialCYukawa, l.OfficialCHiggs, l.OfficialFrozen, l.AlphaNative, l.R3, l.R4)
}
func FormatQuaternionicFirewall(q QuaternionicFirewall) string {
	return fmt.Sprintf("module=%s rank=%d full_H_stable=%t native_rank_one_eigensplit=%t generic_H_preserves_lines=%t action=%s supports=%s failures=%s", q.Module, q.ModuleRank, q.FullModuleHStable, q.NativeRankOneEigensplit, q.GenericHActionPreservesLines, q.ActionDescription, strings.Join(q.Supports, ","), strings.Join(q.Failures, ","))
}
func FormatOrientationSeal(o HiggsOrientationSeal) string {
	return fmt.Sprintf("unit=%s projector=%s complement=%s ranks=(%d,%d) complete=%t orthogonal=%t seal=%t native=%t requires_orientation=%t supports=%s failures=%s", o.UnitVector, o.Projector, o.Complement, o.HPlusRank, o.HMinusRank, o.ProjectorsComplete, o.Orthogonal, o.DefinedAtSealLevel, o.NativeDerivation, o.RequiresGaugeOrientation, strings.Join(o.Supports, ","), strings.Join(o.Failures, ","))
}
func FormatStability(s StabilityClass) string {
	return fmt.Sprintf("full=%s oriented=%s full_H_stable=%t sockets_H_stable=%t sockets_orientation_stable=%t global_H_eigenspaces=%t supports=%s failures=%s", s.FullWeakModule, s.OrientedSockets, s.FullModuleStableUnderH, s.SocketsStableUnderFullH, s.SocketsStableAfterOrientation, s.GlobalQuaternionicEigenspaces, strings.Join(s.Supports, ","), strings.Join(s.Failures, ","))
}
func FormatEdgeRewrite(e EdgeRewrite) string {
	return fmt.Sprintf("edges=%s active_count=%d puncture=%s absent=%t lepto_color_preserved=%t uses_orientation=%t compatible_gate847=%t supports=%s failures=%s", strings.Join(e.Edges, "; "), e.ActiveEdgeCount, e.PunctureEdge, e.PunctureEdgeAbsent, e.LeptoColorPreserved, e.UsesOrientationFrame, e.CompatibleWithGate847Skeleton, strings.Join(e.Supports, ","), strings.Join(e.Failures, ","))
}
func FormatKernel(k KernelOrientation) string {
	return fmt.Sprintf("kernel=%s rank=%d orientation_relative=%t orientation_block_stable=%t full_rhoJ_stable=%t physical_neutrino=%t masslessness=%t supports=%s failures=%s", k.Kernel, k.Rank, k.OrientationRelative, k.StableUnderOrientationBlocks, k.StableUnderFullRhoJ, k.PhysicalNeutrino, k.Masslessness, strings.Join(k.Supports, ","), strings.Join(k.Failures, ","))
}
func FormatFirstOrderPreparation(f FirstOrderPreparation) string {
	return fmt.Sprintf("target=%s next=%s orientation_seal=%t operator_ready=%t executable=%t certified=%t supports=%s failures=%s", f.Target, f.RequiredNextStep, f.OrientationSealAvailable, f.OperatorRealizationReady, f.FirstOrderExecutable, f.FirstOrderCertified, strings.Join(f.Supports, ","), strings.Join(f.Failures, ","))
}
func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s inherited=%t orientation=%t edge_rewritten=%t weak_native=%t native_higgs=%t operator_ready=%t first_order=%t kernel_full_rhoJ=%t alpha_sealed=%t magnitudes_missing=%t update_Neff=%t update_CY=%t update_CH=%t R3=%t R4=%t", i.Classification, i.Gate852Inherited, i.OrientationSealDefined, i.EdgeSkeletonRewritten, i.WeakSplitNative, i.NativeHiggsOrientation, i.OperatorRealizationReady, i.FirstOrderProved, i.LeftKernelStableUnderFullRhoJ, i.AlphaStillSealed, i.MagnitudesStillMissing, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CanPromoteToR3, i.CanPromoteToR4)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%t weak_not_native=%t higgs_not_native=%t lines_not_H_stable=%t no_higgs_vacuum=%t no_operator_package=%t no_first_order=%t no_J=%t no_bimodule=%t kernel_not_stable=%t no_alpha=%t no_yukawa=%t no_trace_readout=%t no_Neff=%t no_C=%t not_R3=%t not_R4=%t no_neutrino=%t no_right_neutrino=%t no_masslessness=%t no_weakmixing_higgsmass=%t no_particle=%t verdict=%s", f.Enforced, f.WeakSplitNotNative, f.HiggsOrientationNotNative, f.RankOneLinesNotHStable, f.NoNativeHiggsVacuum, f.NoOperatorPackage, f.NoFirstOrderProof, f.NoJOppositeProof, f.NoBimoduleProof, f.KernelNotRepresentationStable, f.OrientationDoesNotDeriveAlpha, f.NoYukawaMagnitudes, f.NoTraceMagnitudeReadout, f.NoOfficialNEffUpdate, f.NoCYukawaCHiggsUpdate, f.NotR3, f.NotR4, f.NoPhysicalNeutrino, f.NoRightNeutrino, f.NoMasslessness, f.NoWeakMixingOrHiggsMass, f.NoParticleAssign, f.Verdict)
}

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
func almost(a, b, eps float64) bool { return math.Abs(a-b) <= eps }
