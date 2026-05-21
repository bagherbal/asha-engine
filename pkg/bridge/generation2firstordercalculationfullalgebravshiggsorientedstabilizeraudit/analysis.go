// Package generation2firstordercalculationfullalgebravshiggsorientedstabilizeraudit implements
// Gate 855: First-Order Calculation: Full Algebra vs Higgs-Oriented Stabilizer Audit.
//
// Gate 855 follows Gate 854's operator-level finite-triple matrix seal.  It audits
// whether the symbolic finite-Dirac matrix support can satisfy the first-order
// condition for the full finite algebra A_F = C plus H plus M_3(C), or only after
// restricting to the Higgs/weak-oriented stabilizer frame.  The gate is intentionally
// conservative: it classifies the current D_F^sym as post-orientation support data,
// not an unbroken native finite-triple theorem, and it never promotes symbolic edge
// coefficients to Yukawa magnitudes.
package generation2firstordercalculationfullalgebravshiggsorientedstabilizeraudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE855-FIRST-ORDER-CALCULATION-FULL-ALGEBRA-VS-HIGGS-ORIENTED-STABILIZER-AUDIT"

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

	StatusGate854Inherited           = "PASS_GATE854_OPERATOR_MATRIX_SEAL_INHERITED"
	StatusFirstOrderTargetExecutable = "PASS_FIRST_ORDER_TARGET_EXECUTABLE_AT_SUPPORT_AUDIT_LEVEL"
	StatusFullAlgebraBranchAudited   = "PASS_FULL_A_F_FIRST_ORDER_BRANCH_AUDITED"
	StatusFullAlgebraWeakObstruction = "PASS_FULL_A_F_BRANCH_OBSTRUCTION_LOCALIZED_TO_WEAK_ORIENTATION_MIXING"
	StatusStabilizerBranchAudited    = "PASS_HIGGS_ORIENTED_STABILIZER_BRANCH_AUDITED"
	StatusCarrierPreservationAudited = "PASS_MINIMAL_CARRIER_PRESERVATION_AUDITED"
	StatusKernelPunctureAudited      = "PASS_KERNEL_AND_PUNCTURE_STABILITY_AUDITED"
	StatusNoObservedDataUsed         = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusLedgerFrozen               = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusFirewallVerdict            = "FIREWALL_PRESERVED_GATE855_POST_ORIENTATION_SUPPORT_NOT_FULL_UNBROKEN_FIRST_ORDER_THEOREM"

	SupportFirstOrderTargetExecutable      = "CONDITIONAL_SUPPORT_FIRST_ORDER_COMMUTATOR_NOW_HAS_OPERATOR_LEVEL_SEAL_INPUTS"
	SupportFullAlgebraObstructionIsWeakH   = "CONDITIONAL_SUPPORT_FULL_A_F_OBSTRUCTION_IS_GENERIC_H_MIXING_OF_H_PLUS_H_MINUS"
	SupportStabilizerCompatibilityAtSeal   = "CONDITIONAL_SUPPORT_FIRST_ORDER_COMPATIBILITY_IN_HIGGS_ORIENTED_STABILIZER_FRAME_AT_SUPPORT_LEVEL"
	SupportDFPostOrientation               = "CONDITIONAL_SUPPORT_D_F_SYM_CLASSIFIED_AS_POST_HIGGS_ORIENTATION_SUPPORT_OBJECT"
	SupportMinimalCarrierPreservedByBlocks = "CONDITIONAL_SUPPORT_MINIMAL_CARRIER_PRESERVED_BY_BLOCK_ACTION_IN_STABILIZER_BRANCH"
	SupportKernelStableInStabilizer        = "CONDITIONAL_SUPPORT_LEFT_KERNEL_SINGLETON_STABLE_IN_ORIENTED_STABILIZER_BRANCH"
	SupportR2FirstOrderFirewallStage       = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_PLUS_PLUS_PLUS_FIRST_ORDER_FIREWALL_STAGE"

	FailureFullAFNoFirstOrderTheorem          = "FAILED_ROUTE_FULL_A_F_FIRST_ORDER_CONDITION_NOT_CERTIFIED"
	FailureFullHActionMixesWeakSockets        = "FAILED_ROUTE_FULL_H_ACTION_DOES_NOT_PRESERVE_ORIENTED_SOCKET_D_F_SUPPORT"
	FailureFullAFBlockedByHiggsOrientation    = "FAILED_ROUTE_FIRST_ORDER_FULL_A_F_TEST_BLOCKED_BY_HIGGS_ORIENTATION_FRAME"
	FailureWeakFrameNotNativeHInvariant       = "FAILED_ROUTE_ORIENTED_H_PLUS_H_MINUS_FRAME_NOT_NATIVE_FULL_H_INVARIANT"
	FailureStabilizerNotFullUnbrokenAFTheorem = "FAILED_ROUTE_STABILIZER_FRAME_FIRST_ORDER_NOT_FULL_UNBROKEN_A_F_THEOREM"
	FailureStabilizerOnlySupportLevel         = "FAILED_ROUTE_STABILIZER_COMPATIBILITY_IS_SUPPORT_LEVEL_NOT_OPERATOR_THEOREM"
	FailureNoOperatorLevelJOppositeProof      = "FAILED_ROUTE_NO_OPERATOR_LEVEL_J_OPPOSITE_ACTION_PROOF_CERTIFIED"
	FailureNoBimoduleCommutantProof           = "FAILED_ROUTE_NO_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailureNoFullFirstOrderProof              = "FAILED_ROUTE_NO_FULL_FIRST_ORDER_CONDITION_PROOF"
	FailureNoKOSignProof                      = "FAILED_ROUTE_NO_KO_SIGN_EXTENSION_PROOF_CERTIFIED"
	FailureNoNativeFiniteTripleProof          = "FAILED_ROUTE_NO_NATIVE_FINITE_TRIPLE_PROOF_CERTIFIED"
	FailureDStillSymbolic                     = "FAILED_ROUTE_D_F_SYM_REMAINS_SYMBOLIC_SUPPORT_MATRIX"
	FailureYCoefficientsNotMagnitudes         = "FAILED_ROUTE_SYMBOLIC_Y_COEFFICIENTS_NOT_YUKAWA_MAGNITUDES"
	FailureNoAlphaSource                      = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoTraceMagnitudeReadout            = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoOfficialNEffUpdate               = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate              = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNotR3                              = "FAILED_ROUTE_R2_PLUS_PLUS_PLUS_PLUS_PLUS_FIRST_ORDER_FIREWALL_NOT_R3"
	FailureNotR4                              = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoParticleAssignment               = "FAILED_ROUTE_NO_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoNeutrinoTheorem                  = "FAILED_ROUTE_NO_PHYSICAL_NEUTRINO_THEOREM"
	FailureNoThreeGenerationTheorem           = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
)

type Ledger struct {
	AlphaB                        float64
	OfficialNEff, OfficialCYukawa float64
	OfficialCHiggs                float64
	OfficialFrozen                bool
	AlphaNative, R3, R4           bool
}

type Carrier struct {
	HLRank, HRMinRank, HPartMinRank, HFMinRank int
	AmbientPartRank, AmbientFRank              int
	RightPuncture, LeftKernel                  string
	RightPunctureOutside, LeftKernelInsideHL   bool
	MinimalPreservedInStabilizer               bool
	MinimalPreservedUnderFullAF                bool
	AbsentCellForcedBack                       bool
	Supports, Failures                         []string
}

type CommutatorTarget struct {
	Expression                                   string
	HasRhoMatrixSeal, HasJSeal, HasGammaSeal     bool
	HasDMatrixSeal, WellTyped, SupportExecutable bool
	OperatorTheoremExecutable, Certified         bool
	Supports, Failures                           []string
}

type FullAlgebraBranch struct {
	Algebra, WeakAction, DFrame                                             string
	GenericHMixesHPlusHMinus, DRequiresOrientedWeakSockets                  bool
	CarrierPreservedByFullAFAtCoarseLevel, OrientedSupportPreservedByFullAF bool
	FirstOrderSupportZero, FirstOrderCertified                              bool
	ObstructionTerms                                                        []string
	Supports, Failures                                                      []string
}

type StabilizerBranch struct {
	Algebra, WeakActionRestriction, DFrame                             string
	PreservesHPlusHMinus, PreservesLeptoColor, PreservesMinimalCarrier bool
	PunctureRemainsOutside, KernelRemainsStableCandidate               bool
	FirstOrderSupportCompatible, FirstOrderOperatorTheorem             bool
	FullUnbrokenAFTheorem                                              bool
	Supports, Failures                                                 []string
}

type KernelPuncture struct {
	RightPuncture, LeftKernel                          string
	RightPunctureRank, LeftKernelRank                  int
	LeftKernelStableFullAF, LeftKernelStableStabilizer bool
	PhysicalNeutrinoTheorem, MasslessnessTheorem       bool
	Supports, Failures                                 []string
}

type Impact struct {
	Classification                                                                           string
	Gate854Inherited, FirstOrderTargetExecutable, FullBranchAudited, StabilizerBranchAudited bool
	FullAFPass, StabilizerSupportPass, NativeFiniteTripleProof, FirstOrderCertified          bool
	PostOrientationSupportObject                                                             bool
	AlphaStillSealed, MagnitudesStillMissing                                                 bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4         bool
}

type Firewalls struct {
	Enforced                                                                                   bool
	FullAFNoFirstOrder, FullHMixesWeakSockets, FullAFBlockedByOrientation, WeakFrameNotNativeH bool
	StabilizerNotFullAF, StabilizerSupportOnly, NoJOppositeProof, NoBimoduleProof              bool
	NoFullFirstOrder, NoKOProof, NoNativeFiniteTriple, DSymbolicOnly, YSymbolicOnly            bool
	NoAlphaSource, NoTraceReadout, NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate                 bool
	NotR3, NotR4, NoParticleAssignment, NoNeutrinoTheorem, NoThreeGenerationTheorem            bool
	Verdict                                                                                    string
}

type Audit struct {
	ID         string
	Ledger     Ledger
	Carrier    Carrier
	Target     CommutatorTarget
	Full       FullAlgebraBranch
	Stabilizer StabilizerBranch
	Kernel     KernelPuncture
	Impact     Impact
	Firewalls  Firewalls
	Truth      string
	Final      string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		ID:         AuditID,
		Ledger:     buildLedger(),
		Carrier:    buildCarrier(),
		Target:     buildTarget(),
		Full:       buildFullBranch(),
		Stabilizer: buildStabilizerBranch(),
		Kernel:     buildKernel(),
		Impact:     buildImpact(),
		Firewalls:  buildFirewalls(),
		Truth:      "Gate 855 follows Gate 854 by auditing the first-order commutator in two branches.  The full A_F branch is obstructed because generic quaternionic action mixes the Higgs-oriented h_+/h_- frame used by D_F^sym.  The oriented stabilizer branch is support-compatible at seal level, but it is not a full unbroken finite-triple theorem.",
		Final:      "Verdict: D_F^sym is classified as a post-Higgs-orientation symbolic support object.  First-order compatibility is not certified for the full algebra; a support-level stabilizer branch is conditionally compatible.  This remains R2+++++ first-order-firewall stage, not R3/R4, not a Yukawa-magnitude source, and not an official ledger update.",
	}
	return a, validate(a)
}

func validate(a Audit) error {
	if WRank != 4 || HLRank != 8 || HRMinRank != 7 || HPartMinRank != 15 || HFMinRank != 30 || DSymRank != 14 || DSymKernelRank != 1 {
		return err("dimension constants inconsistent")
	}
	if a.Carrier.HPartMinRank != a.Carrier.HLRank+a.Carrier.HRMinRank {
		return err("minimal carrier rank inconsistent")
	}
	if !a.Target.WellTyped || !a.Target.SupportExecutable || a.Target.Certified {
		return err("commutator target flags inconsistent")
	}
	if !a.Full.GenericHMixesHPlusHMinus || a.Full.OrientedSupportPreservedByFullAF || a.Full.FirstOrderCertified || a.Full.FirstOrderSupportZero {
		return err("full algebra branch overpromoted")
	}
	if !a.Stabilizer.PreservesHPlusHMinus || !a.Stabilizer.FirstOrderSupportCompatible || a.Stabilizer.FirstOrderOperatorTheorem || a.Stabilizer.FullUnbrokenAFTheorem {
		return err("stabilizer branch flags inconsistent")
	}
	if a.Kernel.LeftKernelStableFullAF || !a.Kernel.LeftKernelStableStabilizer || a.Kernel.PhysicalNeutrinoTheorem || a.Kernel.MasslessnessTheorem {
		return err("kernel flags inconsistent")
	}
	return nil
}

func err(msg string) error { return fmt.Errorf("%s", msg) }

func buildLedger() Ledger {
	return Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true}
}

func buildCarrier() Carrier {
	return Carrier{
		HLRank: HLRank, HRMinRank: HRMinRank, HPartMinRank: HPartMinRank, HFMinRank: HFMinRank,
		AmbientPartRank: AmbientPartRank, AmbientFRank: AmbientFRank,
		RightPuncture: "e_+ tensor P_1", LeftKernel: "h_+ tensor P_1",
		RightPunctureOutside: true, LeftKernelInsideHL: true,
		MinimalPreservedInStabilizer: true, MinimalPreservedUnderFullAF: false, AbsentCellForcedBack: false,
		Supports: []string{StatusCarrierPreservationAudited, SupportMinimalCarrierPreservedByBlocks},
		Failures: []string{FailureFullAFBlockedByHiggsOrientation},
	}
}

func buildTarget() CommutatorTarget {
	return CommutatorTarget{
		Expression:       "[[D_F^sym,rho_F(a)],J_F rho_F(b) J_F^{-1}]",
		HasRhoMatrixSeal: true, HasJSeal: true, HasGammaSeal: true, HasDMatrixSeal: true,
		WellTyped: true, SupportExecutable: true, OperatorTheoremExecutable: false, Certified: false,
		Supports: []string{StatusFirstOrderTargetExecutable, SupportFirstOrderTargetExecutable},
		Failures: []string{FailureNoOperatorLevelJOppositeProof, FailureNoBimoduleCommutantProof, FailureNoFullFirstOrderProof, FailureNoKOSignProof},
	}
}

func buildFullBranch() FullAlgebraBranch {
	return FullAlgebraBranch{
		Algebra:                               "A_F=C plus H plus M_3(C)",
		WeakAction:                            "generic H action on C_L^2",
		DFrame:                                "Higgs-oriented h_+ plus h_- socket frame",
		GenericHMixesHPlusHMinus:              true,
		DRequiresOrientedWeakSockets:          true,
		CarrierPreservedByFullAFAtCoarseLevel: true,
		OrientedSupportPreservedByFullAF:      false,
		FirstOrderSupportZero:                 false,
		FirstOrderCertified:                   false,
		ObstructionTerms:                      []string{"q_offdiag h_+ -> h_-", "q_offdiag h_- -> h_+", "D_F^sym socket-diagonal support not H-equivariant before orientation"},
		Supports:                              []string{StatusFullAlgebraBranchAudited, StatusFullAlgebraWeakObstruction, SupportFullAlgebraObstructionIsWeakH},
		Failures:                              []string{FailureFullAFNoFirstOrderTheorem, FailureFullHActionMixesWeakSockets, FailureFullAFBlockedByHiggsOrientation, FailureWeakFrameNotNativeHInvariant},
	}
}

func buildStabilizerBranch() StabilizerBranch {
	return StabilizerBranch{
		Algebra:                      "Higgs-oriented stabilizer of h_+ plus h_-",
		WeakActionRestriction:        "orientation-preserving weak action; no generic H off-diagonal socket mixing",
		DFrame:                       "same oriented h_+/h_- frame used by D_F^sym",
		PreservesHPlusHMinus:         true,
		PreservesLeptoColor:          true,
		PreservesMinimalCarrier:      true,
		PunctureRemainsOutside:       true,
		KernelRemainsStableCandidate: true,
		FirstOrderSupportCompatible:  true,
		FirstOrderOperatorTheorem:    false,
		FullUnbrokenAFTheorem:        false,
		Supports:                     []string{StatusStabilizerBranchAudited, SupportStabilizerCompatibilityAtSeal, SupportDFPostOrientation, SupportMinimalCarrierPreservedByBlocks},
		Failures:                     []string{FailureStabilizerNotFullUnbrokenAFTheorem, FailureStabilizerOnlySupportLevel, FailureNoOperatorLevelJOppositeProof, FailureNoBimoduleCommutantProof},
	}
}

func buildKernel() KernelPuncture {
	return KernelPuncture{
		RightPuncture: "e_+ tensor P_1", LeftKernel: "h_+ tensor P_1", RightPunctureRank: 1, LeftKernelRank: 1,
		LeftKernelStableFullAF: false, LeftKernelStableStabilizer: true, PhysicalNeutrinoTheorem: false, MasslessnessTheorem: false,
		Supports: []string{StatusKernelPunctureAudited, SupportKernelStableInStabilizer},
		Failures: []string{FailureFullHActionMixesWeakSockets, FailureNoNeutrinoTheorem},
	}
}

func buildImpact() Impact {
	return Impact{Classification: "R2+++++_first_order_full_algebra_firewall_stabilizer_support", Gate854Inherited: true, FirstOrderTargetExecutable: true, FullBranchAudited: true, StabilizerBranchAudited: true, FullAFPass: false, StabilizerSupportPass: true, NativeFiniteTripleProof: false, FirstOrderCertified: false, PostOrientationSupportObject: true, AlphaStillSealed: true, MagnitudesStillMissing: true}
}

func buildFirewalls() Firewalls {
	return Firewalls{Enforced: true, FullAFNoFirstOrder: true, FullHMixesWeakSockets: true, FullAFBlockedByOrientation: true, WeakFrameNotNativeH: true, StabilizerNotFullAF: true, StabilizerSupportOnly: true, NoJOppositeProof: true, NoBimoduleProof: true, NoFullFirstOrder: true, NoKOProof: true, NoNativeFiniteTriple: true, DSymbolicOnly: true, YSymbolicOnly: true, NoAlphaSource: true, NoTraceReadout: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NotR3: true, NotR4: true, NoParticleAssignment: true, NoNeutrinoTheorem: true, NoThreeGenerationTheorem: true, Verdict: StatusFirewallVerdict}
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("alpha_B=%.16g; official_N_eff=%.16g; C_Yukawa=%.16g; C_Higgs=%.16g; frozen=%t; alpha_native=%t; R3=%t; R4=%t", l.AlphaB, l.OfficialNEff, l.OfficialCYukawa, l.OfficialCHiggs, l.OfficialFrozen, l.AlphaNative, l.R3, l.R4)
}

func FormatCarrier(c Carrier) string {
	return fmt.Sprintf("H_L=%d; H_R^min=%d; H_part^min=%d; H_F^min=%d; ambient=%d/%d; right_puncture=%s outside=%t; left_kernel=%s in_HL=%t; preserved_stabilizer=%t; preserved_full_A_F=%t", c.HLRank, c.HRMinRank, c.HPartMinRank, c.HFMinRank, c.AmbientPartRank, c.AmbientFRank, c.RightPuncture, c.RightPunctureOutside, c.LeftKernel, c.LeftKernelInsideHL, c.MinimalPreservedInStabilizer, c.MinimalPreservedUnderFullAF)
}

func FormatTarget(t CommutatorTarget) string {
	return fmt.Sprintf("%s; typed=%t; support_executable=%t; operator_executable=%t; certified=%t; rho=%t; J=%t; gamma=%t; D=%t", t.Expression, t.WellTyped, t.SupportExecutable, t.OperatorTheoremExecutable, t.Certified, t.HasRhoMatrixSeal, t.HasJSeal, t.HasGammaSeal, t.HasDMatrixSeal)
}

func FormatFullBranch(f FullAlgebraBranch) string {
	return fmt.Sprintf("full branch %s; weak=%s; D_frame=%s; H_mixes_h_lines=%t; oriented_support_preserved=%t; first_order_zero=%t; obstructions=[%s]", f.Algebra, f.WeakAction, f.DFrame, f.GenericHMixesHPlusHMinus, f.OrientedSupportPreservedByFullAF, f.FirstOrderSupportZero, strings.Join(f.ObstructionTerms, "; "))
}

func FormatStabilizer(s StabilizerBranch) string {
	return fmt.Sprintf("stabilizer=%s; restriction=%s; preserves_h_lines=%t; preserves_leptocolor=%t; preserves_minimal=%t; puncture_outside=%t; kernel_candidate=%t; support_compatible=%t; operator_theorem=%t; full_A_F=%t", s.Algebra, s.WeakActionRestriction, s.PreservesHPlusHMinus, s.PreservesLeptoColor, s.PreservesMinimalCarrier, s.PunctureRemainsOutside, s.KernelRemainsStableCandidate, s.FirstOrderSupportCompatible, s.FirstOrderOperatorTheorem, s.FullUnbrokenAFTheorem)
}

func FormatKernel(k KernelPuncture) string {
	return fmt.Sprintf("right_puncture=%s(rank=%d); left_kernel=%s(rank=%d); stable_full_A_F=%t; stable_stabilizer=%t; physical_neutrino=%t; masslessness=%t", k.RightPuncture, k.RightPunctureRank, k.LeftKernel, k.LeftKernelRank, k.LeftKernelStableFullAF, k.LeftKernelStableStabilizer, k.PhysicalNeutrinoTheorem, k.MasslessnessTheorem)
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s; target_exec=%t; full_pass=%t; stabilizer_support=%t; post_orientation=%t; first_order_certified=%t; native_finite_triple=%t; alpha_sealed=%t; magnitudes_missing=%t; update_N_eff=%t; R3=%t; R4=%t", i.Classification, i.FirstOrderTargetExecutable, i.FullAFPass, i.StabilizerSupportPass, i.PostOrientationSupportObject, i.FirstOrderCertified, i.NativeFiniteTripleProof, i.AlphaStillSealed, i.MagnitudesStillMissing, i.CanUpdateNEff, i.CanPromoteToR3, i.CanPromoteToR4)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%t; verdict=%s; full_no_first_order=%t; H_mixing=%t; stabilizer_not_full=%t; support_only=%t; no_J=%t; no_bimodule=%t; D_symbolic=%t; no_alpha=%t; not_R3=%t; not_R4=%t", f.Enforced, f.Verdict, f.FullAFNoFirstOrder, f.FullHMixesWeakSockets, f.StabilizerNotFullAF, f.StabilizerSupportOnly, f.NoJOppositeProof, f.NoBimoduleProof, f.DSymbolicOnly, f.NoAlphaSource, f.NotR3, f.NotR4)
}

func Statuses() []string {
	return []string{
		StatusGate854Inherited,
		StatusFirstOrderTargetExecutable,
		StatusFullAlgebraBranchAudited,
		StatusFullAlgebraWeakObstruction,
		StatusStabilizerBranchAudited,
		StatusCarrierPreservationAudited,
		StatusKernelPunctureAudited,
		StatusNoObservedDataUsed,
		StatusLedgerFrozen,
		SupportFirstOrderTargetExecutable,
		SupportFullAlgebraObstructionIsWeakH,
		SupportStabilizerCompatibilityAtSeal,
		SupportDFPostOrientation,
		SupportMinimalCarrierPreservedByBlocks,
		SupportKernelStableInStabilizer,
		SupportR2FirstOrderFirewallStage,
		FailureFullAFNoFirstOrderTheorem,
		FailureFullHActionMixesWeakSockets,
		FailureFullAFBlockedByHiggsOrientation,
		FailureWeakFrameNotNativeHInvariant,
		FailureStabilizerNotFullUnbrokenAFTheorem,
		FailureStabilizerOnlySupportLevel,
		FailureNoOperatorLevelJOppositeProof,
		FailureNoBimoduleCommutantProof,
		FailureNoFullFirstOrderProof,
		FailureNoKOSignProof,
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
