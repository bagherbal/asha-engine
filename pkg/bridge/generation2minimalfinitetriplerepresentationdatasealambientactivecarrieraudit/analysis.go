// Package generation2minimalfinitetriplerepresentationdatasealambientactivecarrieraudit implements
// Gate 851: Minimal FiniteTriple Representation DataSeal and Ambient/Active Carrier Audit.
//
// Gate 851 follows Gate 850's first-order/J-opposite firewall.  Gate 850
// showed that the symbolic finite-Dirac support matrix has a coherent support
// shape, but cannot be promoted without a represented finite-triple data
// package.  Gate 851 therefore constructs the missing data as a bridge-level
// seal: the minimal active particle carrier H_part^min=H_L plus H_R^min, its
// real/opposite copy H_F^min, schematic rho_F/gamma_F/J_F/D_F^sym data, and the
// ambient-vs-active carrier distinction.  This is not a native finite-triple
// theorem and does not prove the first-order condition.
package generation2minimalfinitetriplerepresentationdatasealambientactivecarrieraudit

import (
	"strconv"
	"strings"
)

const (
	AuditID = "GATE851-MINIMAL-FINITE-TRIPLE-REPRESENTATION-DATA-SEAL-AMBIENT-ACTIVE-CARRIER-AUDIT"

	AlphaB          = 0.0003878958469680527
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	LeptonBlockDim     = 1
	ColorBlockDim      = 3
	WDim               = LeptonBlockDim + ColorBlockDim
	RightSocketPairDim = 2
	WeakDoubletDim     = 2
	HLRank             = WeakDoubletDim * WDim
	HRAmbientRank      = RightSocketPairDim * WDim
	RightPunctureRank  = 1
	HRMinRank          = HRAmbientRank - RightPunctureRank
	HPartAmbientRank   = HLRank + HRAmbientRank
	HPartMinRank       = HLRank + HRMinRank
	HFAmbientRank      = 2 * HPartAmbientRank
	HFMinRank          = 2 * HPartMinRank
	SymbolicYRank      = HRMinRank
	SymbolicDFRank     = 2 * SymbolicYRank
	SymbolicKernelDim  = HPartMinRank - SymbolicDFRank

	StatusGate850Inherited             = "PASS_GATE850_COMPATIBILITY_FIREWALL_INHERITED"
	StatusAmbientCarrierSeparated      = "PASS_AMBIENT_16_32_AND_ACTIVE_15_30_CARRIERS_SEPARATED"
	StatusMinimalParticleCarrier       = "PASS_MINIMAL_ACTIVE_CARRIER_H_PART_DIM_15_DEFINED"
	StatusMinimalRealCarrier           = "PASS_REAL_COPY_H_F_MIN_DIM_30_DEFINED"
	StatusPunctureOutsideRightCarrier  = "PASS_E_PLUS_TENSOR_P1_OUTSIDE_MINIMAL_RIGHT_CARRIER"
	StatusRhoFActionSealed             = "PASS_RHO_F_ACTION_SEALED_ON_MINIMAL_CARRIER"
	StatusRepresentationClosureAudited = "PASS_MINIMAL_CARRIER_REPRESENTATION_CLOSURE_AUDITED"
	StatusGammaFSealed                 = "PASS_GAMMA_F_CHIRALITY_SEALED"
	StatusJFSealed                     = "PASS_J_F_OPPOSITE_COPY_SEALED"
	StatusDFSymSealed                  = "PASS_SYMBOLIC_D_F_EXTENDED_TO_MINIMAL_H_F"
	StatusFirstOrderPrepared           = "PASS_FIRST_ORDER_TARGET_OBJECTS_PREPARED_FOR_NEXT_GATE"
	StatusNoObservedDataUsed           = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusOfficialLedgersFrozen        = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusDataSealVerdict              = "FIREWALL_PRESERVED_GATE851_MINIMAL_FINITE_TRIPLE_DATA_SEAL"

	SupportMinimalFiniteTripleDataSeal     = "CONDITIONAL_SUPPORT_MINIMAL_FINITE_TRIPLE_REPRESENTATION_DATA_SEAL_DEFINED"
	SupportAmbientToActiveForkTyped        = "CONDITIONAL_SUPPORT_AMBIENT_16_32_TO_ACTIVE_15_30_BRANCH_SEPARATED"
	SupportRhoFPreservesMinimalCarrier     = "CONDITIONAL_SUPPORT_RHO_F_SCHEMATIC_ACTION_PRESERVES_MINIMAL_ACTIVE_CARRIER"
	SupportRightCharacterProjectorsSealed  = "CONDITIONAL_SUPPORT_RHO_R_LAMBDA_BARLAMBDA_SEALS_RIGHT_SOCKET_PROJECTORS"
	SupportLeptoColorActionSealed          = "CONDITIONAL_SUPPORT_M3C_ACTS_ON_P3_AND_TRIVIALLY_ON_P1_AT_SEAL_LEVEL"
	SupportGammaChiralitySupport           = "CONDITIONAL_SUPPORT_GAMMA_F_LEFT_RIGHT_CHIRALITY_SUPPORT_SEALED"
	SupportJCopySupport                    = "CONDITIONAL_SUPPORT_J_F_PARTICLE_OPPOSITE_COPY_SEALED"
	SupportDFSymSupport                    = "CONDITIONAL_SUPPORT_D_F_SYM_SUPPORT_EXTENDED_TO_MINIMAL_H_F"
	SupportKernelStabilityQuestionIsolated = "CONDITIONAL_SUPPORT_LEFT_KERNEL_STABILITY_QUESTION_ISOLATED"
	SupportFirstOrderReadyForCalculation   = "CONDITIONAL_SUPPORT_FIRST_ORDER_EXPRESSION_READY_FOR_GATE852_CALCULATION"
	SupportR2DataSealStage                 = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_PLUS_PLUS_PLUS_DATA_SEAL_STAGE"

	FailureDataSealNotNativeFiniteTripleProof = "FAILED_ROUTE_MINIMAL_FINITE_TRIPLE_DATA_SEAL_IS_NOT_NATIVE_FINITE_TRIPLE_PROOF"
	FailureNoFullFirstOrderConditionProof     = "FAILED_ROUTE_NO_FULL_FIRST_ORDER_CONDITION_PROOF_YET"
	FailureNoJKOCompatibilityProof            = "FAILED_ROUTE_NO_J_F_KO_SIGN_OR_OPPOSITE_ACTION_PROOF_CERTIFIED"
	FailureNoBimoduleCommutantProof           = "FAILED_ROUTE_NO_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailureNoOperatorValuedDFMatrix           = "FAILED_ROUTE_NO_OPERATOR_VALUED_D_F_MATRIX_CERTIFIED"
	FailureDFSymNotYukawaMagnitudeSource      = "FAILED_ROUTE_SYMBOLIC_D_F_NOT_YUKAWA_MAGNITUDE_SOURCE"
	FailureNoNumericalYukawaValues            = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureKernelStabilityNotCertified        = "FAILED_ROUTE_LEFT_KERNEL_STABILITY_NOT_CERTIFIED_UNDER_FULL_RHO_F"
	FailureWeakSocketSplitNotNative           = "FAILED_ROUTE_WEAK_SOCKET_SPLIT_REMAINS_ORIENTATION_SEAL_NOT_NATIVE_H_ACTION_EIGENSPLIT"
	FailurePunctureAbsenceNotNative           = "FAILED_ROUTE_RIGHT_NEUTRAL_PUNCTURE_ABSENCE_REMAINS_SEAL_NOT_NATIVE_THEOREM"
	FailureNoPhysicalNeutrinoTheorem          = "FAILED_ROUTE_NO_PHYSICAL_NEUTRINO_THEOREM"
	FailureNoRightNeutrinoTheorem             = "FAILED_ROUTE_NO_RIGHT_NEUTRINO_OR_STERILE_PARTICLE_THEOREM"
	FailureNoMasslessnessTheorem              = "FAILED_ROUTE_SYMBOLIC_KERNEL_NOT_OBSERVED_MASSLESSNESS_THEOREM"
	FailureAlphaStillSealed                   = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoTraceMagnitudeReadout            = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoOfficialNEffUpdate               = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate              = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNotR3                              = "FAILED_ROUTE_R2_PLUS_PLUS_PLUS_PLUS_PLUS_DATA_SEAL_NOT_R3"
	FailureNotR4                              = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	AlphaB                        float64
	OfficialNEff, OfficialCYukawa float64
	OfficialCHiggs                float64
	OfficialFrozen                bool
	R2DataSealStage               bool
	R3, R4, AlphaNative           bool
}

type CarrierAudit struct {
	W, HL, HRAmbient, HRMin, HPartAmbient, HPartMin, HFAmbient, HFMin string
	WRank, HLRank, HRAmbientRank, HRMinRank                           int
	HPartAmbientRank, HPartMinRank, HFAmbientRank, HFMinRank          int
	Puncture, LeftKernel                                              Cell
	AmbientActiveSeparated, MinimalBranchSelected, ExtendedBranchKept bool
	Supports, Failures                                                []string
}

type Cell struct {
	Name, Support, Chirality, Socket, LeptoColor, Role string
	Rank                                               int
	Absent, Kernel, InMinimalCarrier, PhysicalName     bool
}

type RhoFSeal struct {
	Algebra, ActionSummary, RightAction, LeftAction, LeptoColorAction   string
	PreservesMinimalCarrier, CompleteActionLedger, NativeProof          bool
	RightSocketsSealed, LeptoColorBlocksPreserved, WeakDoubletPreserved bool
	AbsentCellClosureSafe, KernelStableUnderFullAction                  bool
	Supports, Failures                                                  []string
}

type GammaSeal struct {
	Expression                                      string
	LeftSign, RightSign                             int
	SupportLevel, NativeGammaMatrix, KOExtensionSet bool
	Supports, Failures                              []string
}

type JSeal struct {
	Expression                                         string
	ParticleRank, OppositeRank, TotalRank              int
	AntiunitaryExchangeSealed, OppositeActionCertified bool
	KODataCertified                                    bool
	Supports, Failures                                 []string
}

type DFSymSeal struct {
	Expression, YExpression, Domain, Codomain, Space string
	YRank, DFRank, KernelDim                         int
	ExtendedToJCopy, SupportOnly, OperatorValued     bool
	YukawaMagnitudeSource                            bool
	Supports, Failures                               []string
}

type FirstOrderPrep struct {
	Expression                                       string
	ObjectsPrepared, CanCalculateFirstOrderNow       bool
	HasRhoSeal, HasJSeal, HasGammaSeal, HasDFSymSeal bool
	NativeRhoF, NativeJ, NativeGamma, NativeDF       bool
	MissingForProof                                  []string
	Supports, Failures                               []string
}

type Impact struct {
	Classification                                                                   string
	Gate850Inherited, DataSealDefined, AmbientActiveForkSeparated                    bool
	RepresentationClosureSealed, FirstOrderReadyButNotProven                         bool
	NativeFiniteTriple, FirstOrderCertified, JOppositeCertified, KernelStable        bool
	PhysicalNeutrinoTheorem, MasslessTheorem                                         bool
	AlphaStillSealed, MagnitudesStillMissing                                         bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4 bool
}

type Firewalls struct {
	Enforced                                                                                     bool
	DataSealNotNative, NoFirstOrderProof, NoJKOProof, NoBimoduleProof, NoOperatorDF              bool
	DFSymNotMagnitudeSource, NoNumericalYukawas, KernelStabilityNotCertified, WeakSocketSealOnly bool
	PunctureAbsenceSealOnly, NoPhysicalNeutrino, NoRightNeutrino, NoMasslessness                 bool
	AlphaStillSealed, NoTraceMagnitudeReadout, NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate       bool
	NotR3, NotR4                                                                                 bool
	Verdict                                                                                      string
}

type Audit struct {
	Ledger     Ledger
	Carrier    CarrierAudit
	RhoF       RhoFSeal
	Gamma      GammaSeal
	J          JSeal
	DFSym      DFSymSeal
	FirstOrder FirstOrderPrep
	Impact     Impact
	Firewalls  Firewalls
	Truth      string
	Final      string
}

func BuildDefault() (Audit, error) {
	a := Audit{Ledger: buildLedger(), Carrier: buildCarrier(), RhoF: buildRhoF(), Gamma: buildGamma(), J: buildJ(), DFSym: buildDFSym(), FirstOrder: buildFirstOrderPrep(), Impact: buildImpact(), Firewalls: buildFirewalls(), Truth: "Gate 851 follows Gate 850's compatibility firewall by instantiating the missing represented finite-triple objects as a minimal data seal: H_part^min=H_L plus H_R^min has rank 15, H_F^min has rank 30, and the ambient 16/32 carrier remains distinct from the active minimal 15/30 branch.", Final: "Verdict: minimal finite-triple representation data seal defined at bridge level.  It prepares the objects needed for a first-order/J-opposite audit, but it is not a native finite-triple proof, not a Yukawa-magnitude source, not an alpha source, and not R3/R4."}
	return a, validate(a)
}

func validate(a Audit) error {
	if WDim != 4 || HLRank != 8 || HRAmbientRank != 8 || HRMinRank != 7 || HPartMinRank != 15 || HFMinRank != 30 || SymbolicKernelDim != 1 {
		return err("dimension constants inconsistent")
	}
	if a.Carrier.HPartMinRank != a.Carrier.HLRank+a.Carrier.HRMinRank || a.Carrier.HFMinRank != 2*a.Carrier.HPartMinRank {
		return err("carrier ranks inconsistent")
	}
	if !a.RhoF.PreservesMinimalCarrier || !a.RhoF.AbsentCellClosureSafe || a.RhoF.KernelStableUnderFullAction {
		return err("rho_F seal closure/kernel flags inconsistent")
	}
	if !a.FirstOrder.ObjectsPrepared || a.FirstOrder.CanCalculateFirstOrderNow {
		return err("first-order preparation flags inconsistent")
	}
	return nil
}

func buildLedger() Ledger {
	return Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true, R2DataSealStage: true, R3: false, R4: false, AlphaNative: false}
}

func buildCarrier() CarrierAudit {
	puncture := Cell{Name: "e_+ tensor P_1", Support: "right neutral puncture outside H_R^min", Chirality: "right", Socket: "e_+", LeptoColor: "P_1", Role: "absent ambient singleton removed by minimal right-neutral absence seal", Rank: 1, Absent: true, Kernel: false, InMinimalCarrier: false, PhysicalName: false}
	kernel := Cell{Name: "h_+ tensor P_1", Support: "left neutral kernel singleton inside H_L", Chirality: "left", Socket: "h_+", LeptoColor: "P_1", Role: "forced symbolic D_F kernel candidate pending rho_F/J_F stability proof", Rank: 1, Absent: false, Kernel: true, InMinimalCarrier: true, PhysicalName: false}
	return CarrierAudit{
		W: "W=C_l plus C_c^3", HL: "H_L=(h_+ plus h_-) tensor (P_1 plus P_3)", HRAmbient: "H_R^ambient=C_R^2 tensor W", HRMin: "H_R^min=(C_R^2 tensor W) minus (e_+ tensor P_1)", HPartAmbient: "H_part^ambient=H_L plus H_R^ambient", HPartMin: "H_part^min=H_L plus H_R^min", HFAmbient: "H_F^ambient=H_part^ambient plus J_F H_part^ambient", HFMin: "H_F^min=H_part^min plus J_F H_part^min",
		WRank: WDim, HLRank: HLRank, HRAmbientRank: HRAmbientRank, HRMinRank: HRMinRank, HPartAmbientRank: HPartAmbientRank, HPartMinRank: HPartMinRank, HFAmbientRank: HFAmbientRank, HFMinRank: HFMinRank,
		Puncture: puncture, LeftKernel: kernel, AmbientActiveSeparated: true, MinimalBranchSelected: true, ExtendedBranchKept: true,
		Supports: []string{SupportMinimalFiniteTripleDataSeal, SupportAmbientToActiveForkTyped},
		Failures: []string{FailureDataSealNotNativeFiniteTripleProof, FailurePunctureAbsenceNotNative},
	}
}

func buildRhoF() RhoFSeal {
	return RhoFSeal{
		Algebra:                 "A_F=C plus H plus M_3(C)",
		ActionSummary:           "rho_F(lambda,q,m) acts schematically by C on right character sockets, H on the left weak doublet, and M_3(C) on the P_3 color block",
		RightAction:             "rho_R(lambda)=diag(lambda,conjugate(lambda)) on e_+ plus e_-",
		LeftAction:              "q in H acts on C_L^2=h_+ plus h_- as weak-doublet module support",
		LeptoColorAction:        "m in M_3(C) acts on P_3W=C_c^3 and trivially on P_1W=C_l",
		PreservesMinimalCarrier: true, CompleteActionLedger: false, NativeProof: false, RightSocketsSealed: true, LeptoColorBlocksPreserved: true, WeakDoubletPreserved: true, AbsentCellClosureSafe: true, KernelStableUnderFullAction: false,
		Supports: []string{SupportRhoFPreservesMinimalCarrier, SupportRightCharacterProjectorsSealed, SupportLeptoColorActionSealed},
		Failures: []string{FailureDataSealNotNativeFiniteTripleProof, FailureWeakSocketSplitNotNative, FailureKernelStabilityNotCertified},
	}
}

func buildGamma() GammaSeal {
	return GammaSeal{Expression: "gamma_F=+1 on H_L and -1 on H_R^min, extended to J_F-copy by a pending KO-sign convention", LeftSign: 1, RightSign: -1, SupportLevel: true, NativeGammaMatrix: false, KOExtensionSet: false, Supports: []string{SupportGammaChiralitySupport}, Failures: []string{FailureDataSealNotNativeFiniteTripleProof, FailureNoJKOCompatibilityProof}}
}

func buildJ() JSeal {
	return JSeal{Expression: "J_F antiunitarily exchanges H_part^min with J_F H_part^min at seal level", ParticleRank: HPartMinRank, OppositeRank: HPartMinRank, TotalRank: HFMinRank, AntiunitaryExchangeSealed: true, OppositeActionCertified: false, KODataCertified: false, Supports: []string{SupportJCopySupport}, Failures: []string{FailureNoJKOCompatibilityProof, FailureDataSealNotNativeFiniteTripleProof}}
}

func buildDFSym() DFSymSeal {
	return DFSymSeal{Expression: "D_F^sym=[[0,Y_supp^dagger],[Y_supp,0]] plus J-copy support", YExpression: "Y_supp=y_+3Y_+3+y_-3Y_-3+y_-1Y_-1, y_+1=0", Domain: "H_R^min", Codomain: "H_L", Space: "H_part^min plus J_F-copy", YRank: SymbolicYRank, DFRank: SymbolicDFRank, KernelDim: SymbolicKernelDim, ExtendedToJCopy: true, SupportOnly: true, OperatorValued: false, YukawaMagnitudeSource: false, Supports: []string{SupportDFSymSupport}, Failures: []string{FailureNoOperatorValuedDFMatrix, FailureDFSymNotYukawaMagnitudeSource, FailureNoNumericalYukawaValues}}
}

func buildFirstOrderPrep() FirstOrderPrep {
	missing := []string{"operator-valued rho_F matrices", "certified J_F opposite action", "KO-sign convention", "operator-valued D_F matrix", "bimodule/commutant decomposition", "explicit first-order commutator evaluation"}
	return FirstOrderPrep{Expression: "[[D_F,rho_F(a)],J_F rho_F(b) J_F^{-1}]=0", ObjectsPrepared: true, CanCalculateFirstOrderNow: false, HasRhoSeal: true, HasJSeal: true, HasGammaSeal: true, HasDFSymSeal: true, NativeRhoF: false, NativeJ: false, NativeGamma: false, NativeDF: false, MissingForProof: missing, Supports: []string{SupportFirstOrderReadyForCalculation}, Failures: []string{FailureNoFullFirstOrderConditionProof, FailureNoJKOCompatibilityProof, FailureNoBimoduleCommutantProof, FailureNoOperatorValuedDFMatrix}}
}

func buildImpact() Impact {
	return Impact{Classification: "R2+++++_data_seal minimal finite-triple representation data seal; ambient 16/32 separated from active 15/30; not R3/R4", Gate850Inherited: true, DataSealDefined: true, AmbientActiveForkSeparated: true, RepresentationClosureSealed: true, FirstOrderReadyButNotProven: true, NativeFiniteTriple: false, FirstOrderCertified: false, JOppositeCertified: false, KernelStable: false, PhysicalNeutrinoTheorem: false, MasslessTheorem: false, AlphaStillSealed: true, MagnitudesStillMissing: true, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, CanPromoteToR3: false, CanPromoteToR4: false}
}

func buildFirewalls() Firewalls {
	return Firewalls{Enforced: true, DataSealNotNative: true, NoFirstOrderProof: true, NoJKOProof: true, NoBimoduleProof: true, NoOperatorDF: true, DFSymNotMagnitudeSource: true, NoNumericalYukawas: true, KernelStabilityNotCertified: true, WeakSocketSealOnly: true, PunctureAbsenceSealOnly: true, NoPhysicalNeutrino: true, NoRightNeutrino: true, NoMasslessness: true, AlphaStillSealed: true, NoTraceMagnitudeReadout: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NotR3: true, NotR4: true, Verdict: StatusDataSealVerdict}
}

func Statuses() []string {
	return []string{StatusGate850Inherited, StatusAmbientCarrierSeparated, StatusMinimalParticleCarrier, StatusMinimalRealCarrier, StatusPunctureOutsideRightCarrier, StatusRhoFActionSealed, StatusRepresentationClosureAudited, StatusGammaFSealed, StatusJFSealed, StatusDFSymSealed, StatusFirstOrderPrepared, StatusNoObservedDataUsed, StatusOfficialLedgersFrozen, StatusDataSealVerdict, SupportMinimalFiniteTripleDataSeal, SupportAmbientToActiveForkTyped, SupportRhoFPreservesMinimalCarrier, SupportRightCharacterProjectorsSealed, SupportLeptoColorActionSealed, SupportGammaChiralitySupport, SupportJCopySupport, SupportDFSymSupport, SupportKernelStabilityQuestionIsolated, SupportFirstOrderReadyForCalculation, SupportR2DataSealStage, FailureDataSealNotNativeFiniteTripleProof, FailureNoFullFirstOrderConditionProof, FailureNoJKOCompatibilityProof, FailureNoBimoduleCommutantProof, FailureNoOperatorValuedDFMatrix, FailureDFSymNotYukawaMagnitudeSource, FailureNoNumericalYukawaValues, FailureKernelStabilityNotCertified, FailureWeakSocketSplitNotNative, FailurePunctureAbsenceNotNative, FailureNoPhysicalNeutrinoTheorem, FailureNoRightNeutrinoTheorem, FailureNoMasslessnessTheorem, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNotR3, FailureNotR4}
}

func FormatLedger(l Ledger) string {
	return join("ledger", []string{"alpha_B=" + f64(l.AlphaB), "official_N_eff=" + f64(l.OfficialNEff), "official_C_Yukawa=" + f64(l.OfficialCYukawa), "official_C_Higgs=" + f64(l.OfficialCHiggs), "official_frozen=" + b(l.OfficialFrozen), "R2_data_seal_stage=" + b(l.R2DataSealStage), "R3=" + b(l.R3), "R4=" + b(l.R4), "alpha_native=" + b(l.AlphaNative)})
}

func FormatCell(c Cell) string {
	return join("cell", []string{c.Name, "support=" + c.Support, "chirality=" + c.Chirality, "socket=" + c.Socket, "lepto_color=" + c.LeptoColor, "role=" + c.Role, "rank=" + i(c.Rank), "absent=" + b(c.Absent), "kernel=" + b(c.Kernel), "in_minimal_carrier=" + b(c.InMinimalCarrier), "physical_name=" + b(c.PhysicalName)})
}

func FormatCarrier(c CarrierAudit) string {
	return join("carrier", []string{c.W, c.HL, c.HRAmbient, c.HRMin, c.HPartAmbient, c.HPartMin, c.HFAmbient, c.HFMin, "W_rank=" + i(c.WRank), "HL_rank=" + i(c.HLRank), "HR_ambient_rank=" + i(c.HRAmbientRank), "HR_min_rank=" + i(c.HRMinRank), "H_part_ambient_rank=" + i(c.HPartAmbientRank), "H_part_min_rank=" + i(c.HPartMinRank), "H_F_ambient_rank=" + i(c.HFAmbientRank), "H_F_min_rank=" + i(c.HFMinRank), "puncture=" + FormatCell(c.Puncture), "left_kernel=" + FormatCell(c.LeftKernel), "ambient_active_separated=" + b(c.AmbientActiveSeparated), "minimal_branch_selected=" + b(c.MinimalBranchSelected), "extended_branch_kept=" + b(c.ExtendedBranchKept), "supports=" + strings.Join(c.Supports, ","), "failures=" + strings.Join(c.Failures, ",")})
}

func FormatRhoF(r RhoFSeal) string {
	return join("rho_F_seal", []string{"algebra=" + r.Algebra, r.ActionSummary, "right=" + r.RightAction, "left=" + r.LeftAction, "lepto_color=" + r.LeptoColorAction, "preserves_minimal_carrier=" + b(r.PreservesMinimalCarrier), "complete_action_ledger=" + b(r.CompleteActionLedger), "native_proof=" + b(r.NativeProof), "right_sockets_sealed=" + b(r.RightSocketsSealed), "lepto_color_blocks_preserved=" + b(r.LeptoColorBlocksPreserved), "weak_doublet_preserved=" + b(r.WeakDoubletPreserved), "absent_cell_closure_safe=" + b(r.AbsentCellClosureSafe), "kernel_stable_under_full_action=" + b(r.KernelStableUnderFullAction), "supports=" + strings.Join(r.Supports, ","), "failures=" + strings.Join(r.Failures, ",")})
}

func FormatGamma(g GammaSeal) string {
	return join("gamma_F", []string{g.Expression, "left_sign=" + i(g.LeftSign), "right_sign=" + i(g.RightSign), "support_level=" + b(g.SupportLevel), "native_gamma_matrix=" + b(g.NativeGammaMatrix), "KO_extension_set=" + b(g.KOExtensionSet), "supports=" + strings.Join(g.Supports, ","), "failures=" + strings.Join(g.Failures, ",")})
}

func FormatJ(j JSeal) string {
	return join("J_F", []string{j.Expression, "particle_rank=" + i(j.ParticleRank), "opposite_rank=" + i(j.OppositeRank), "total_rank=" + i(j.TotalRank), "antiunitary_exchange_sealed=" + b(j.AntiunitaryExchangeSealed), "opposite_action_certified=" + b(j.OppositeActionCertified), "KO_data_certified=" + b(j.KODataCertified), "supports=" + strings.Join(j.Supports, ","), "failures=" + strings.Join(j.Failures, ",")})
}

func FormatDFSym(d DFSymSeal) string {
	return join("D_F_sym", []string{d.Expression, d.YExpression, "domain=" + d.Domain, "codomain=" + d.Codomain, "space=" + d.Space, "Y_rank=" + i(d.YRank), "D_F_rank=" + i(d.DFRank), "kernel_dim=" + i(d.KernelDim), "extended_to_J_copy=" + b(d.ExtendedToJCopy), "support_only=" + b(d.SupportOnly), "operator_valued=" + b(d.OperatorValued), "yukawa_magnitude_source=" + b(d.YukawaMagnitudeSource), "supports=" + strings.Join(d.Supports, ","), "failures=" + strings.Join(d.Failures, ",")})
}

func FormatFirstOrder(f FirstOrderPrep) string {
	return join("first_order_prep", []string{f.Expression, "objects_prepared=" + b(f.ObjectsPrepared), "can_calculate_now=" + b(f.CanCalculateFirstOrderNow), "rho_seal=" + b(f.HasRhoSeal), "J_seal=" + b(f.HasJSeal), "gamma_seal=" + b(f.HasGammaSeal), "D_F_sym_seal=" + b(f.HasDFSymSeal), "native_rho=" + b(f.NativeRhoF), "native_J=" + b(f.NativeJ), "native_gamma=" + b(f.NativeGamma), "native_D_F=" + b(f.NativeDF), "missing=" + strings.Join(f.MissingForProof, ","), "supports=" + strings.Join(f.Supports, ","), "failures=" + strings.Join(f.Failures, ",")})
}

func FormatImpact(im Impact) string {
	return join("impact", []string{im.Classification, "gate850_inherited=" + b(im.Gate850Inherited), "data_seal_defined=" + b(im.DataSealDefined), "ambient_active_fork=" + b(im.AmbientActiveForkSeparated), "representation_closure_sealed=" + b(im.RepresentationClosureSealed), "first_order_ready_not_proven=" + b(im.FirstOrderReadyButNotProven), "native_finite_triple=" + b(im.NativeFiniteTriple), "first_order=" + b(im.FirstOrderCertified), "J_opposite=" + b(im.JOppositeCertified), "kernel_stable=" + b(im.KernelStable), "physical_neutrino=" + b(im.PhysicalNeutrinoTheorem), "massless_theorem=" + b(im.MasslessTheorem), "alpha_sealed=" + b(im.AlphaStillSealed), "magnitudes_missing=" + b(im.MagnitudesStillMissing), "update_N_eff=" + b(im.CanUpdateNEff), "update_C_Yukawa=" + b(im.CanUpdateCYukawa), "update_C_Higgs=" + b(im.CanUpdateCHiggs), "R3=" + b(im.CanPromoteToR3), "R4=" + b(im.CanPromoteToR4)})
}

func FormatFirewalls(f Firewalls) string {
	return join("firewalls", []string{"enforced=" + b(f.Enforced), "data_seal_not_native=" + b(f.DataSealNotNative), "no_first_order_proof=" + b(f.NoFirstOrderProof), "no_J_KO_proof=" + b(f.NoJKOProof), "no_bimodule_proof=" + b(f.NoBimoduleProof), "no_operator_D_F=" + b(f.NoOperatorDF), "D_F_sym_not_magnitude_source=" + b(f.DFSymNotMagnitudeSource), "no_numerical_yukawas=" + b(f.NoNumericalYukawas), "kernel_stability_not_certified=" + b(f.KernelStabilityNotCertified), "weak_socket_seal_only=" + b(f.WeakSocketSealOnly), "puncture_absence_seal_only=" + b(f.PunctureAbsenceSealOnly), "no_physical_neutrino=" + b(f.NoPhysicalNeutrino), "no_right_neutrino=" + b(f.NoRightNeutrino), "no_masslessness=" + b(f.NoMasslessness), "alpha_sealed=" + b(f.AlphaStillSealed), "no_trace_magnitude_readout=" + b(f.NoTraceMagnitudeReadout), "no_official_N_eff_update=" + b(f.NoOfficialNEffUpdate), "no_C_Yukawa_C_Higgs_update=" + b(f.NoCYukawaCHiggsUpdate), "not_R3=" + b(f.NotR3), "not_R4=" + b(f.NotR4), "verdict=" + f.Verdict})
}

func containsAll(hay []string, needles []string) bool {
	m := map[string]bool{}
	for _, s := range hay {
		m[s] = true
	}
	for _, n := range needles {
		if !m[n] {
			return false
		}
	}
	return true
}

type err string

func (e err) Error() string                   { return string(e) }
func join(name string, parts []string) string { return name + "{" + strings.Join(parts, "; ") + "}" }
func f64(x float64) string                    { return strconv.FormatFloat(x, 'g', 17, 64) }
func i(x int) string                          { return strconv.Itoa(x) }
func b(x bool) string                         { return strconv.FormatBool(x) }
