// Package generation2firstorderjoppositecompatibilitycalculationaudit implements
// Gate 852: First-Order / J-Opposite Compatibility Calculation Audit.
//
// Gate 852 follows Gate 851's minimal finite-triple representation data seal.
// The first-order expression is now well typed, but the audit must determine
// whether it can actually be calculated from seal-level data.  The gate tracks
// the separate obstruction sources: incomplete rho_F/J_F operator data, the
// non-native weak rank-one socket split, minimal-carrier closure, symbolic
// D_F support, and kernel stability.  It does not promote the symbolic support
// matrix to a native finite-triple theorem.
package generation2firstorderjoppositecompatibilitycalculationaudit

import (
	"strconv"
	"strings"
)

const (
	AuditID = "GATE852-FIRST-ORDER-J-OPPOSITE-COMPATIBILITY-CALCULATION-AUDIT"

	AlphaB          = 0.0003878958469680527
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	P1Rank            = 1
	P3Rank            = 3
	WRank             = P1Rank + P3Rank
	WeakDoubletRank   = 2
	RightAmbientRank  = 2 * WRank
	RightPunctureRank = 1
	RightMinRank      = RightAmbientRank - RightPunctureRank
	HLRank            = WeakDoubletRank * WRank
	HPartMinRank      = HLRank + RightMinRank
	HFMinRank         = 2 * HPartMinRank
	YSupportRank      = RightMinRank
	DFSymRank         = 2 * YSupportRank
	DFSymKernelDim    = HPartMinRank - DFSymRank

	StatusGate851DataSealInherited     = "PASS_GATE851_MINIMAL_DATA_SEAL_INHERITED"
	StatusFirstOrderTargetTyped        = "PASS_FIRST_ORDER_TARGET_NOW_TYPED"
	StatusMinimalCarrierClosureAudited = "PASS_MINIMAL_CARRIER_CLOSURE_AUDITED"
	StatusWeakOrientationAudited       = "PASS_WEAK_H_PLUS_H_MINUS_ORIENTATION_STABILITY_AUDITED"
	StatusJOppositeRequirementAudited  = "PASS_J_OPPOSITE_ACTION_REQUIREMENT_AUDITED"
	StatusFirstOrderExpressionAudited  = "PASS_FIRST_ORDER_COMMUTATOR_EXPRESSION_AUDITED"
	StatusKernelStabilityAudited       = "PASS_LEFT_KERNEL_STABILITY_REQUIREMENT_AUDITED"
	StatusNoObservedDataUsed           = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusOfficialLedgersFrozen        = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusCompatibilityFirewallVerdict = "FIREWALL_PRESERVED_GATE852_FIRST_ORDER_COMPATIBILITY_NOT_CERTIFIED"

	SupportFirstOrderTargetWellTyped        = "CONDITIONAL_SUPPORT_FIRST_ORDER_TARGET_WELL_TYPED_AFTER_DATA_SEAL"
	SupportMinimalCarrierClosedByBlockSeal  = "CONDITIONAL_SUPPORT_MINIMAL_CARRIER_CLOSED_UNDER_BLOCK_ACTION_SEAL"
	SupportWeakOrientationRequiresHiggsSeal = "CONDITIONAL_SUPPORT_WEAK_SOCKET_SPLIT_REQUIRES_HIGGS_ORIENTATION_SEAL"
	SupportChiralityOddnessSupportLevel     = "CONDITIONAL_SUPPORT_CHIRALITY_ODDNESS_REMAINS_SUPPORT_LEVEL_ONLY"
	SupportKernelStableCandidate            = "CONDITIONAL_SUPPORT_LEFT_NEUTRAL_KERNEL_SINGLETON_IS_STABILITY_CANDIDATE"
	SupportFirstOrderObstructionClassified  = "CONDITIONAL_SUPPORT_FIRST_ORDER_OBSTRUCTION_SOURCE_CLASSIFIED"
	SupportR2CompatibilityFirewallStage     = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_PLUS_PLUS_PLUS_DATA_SEAL_COMPATIBILITY_FIREWALL_STAGE"

	FailureDataSealNotNativeFiniteTripleProof  = "FAILED_ROUTE_MINIMAL_FINITE_TRIPLE_DATA_SEAL_IS_NOT_NATIVE_FINITE_TRIPLE_PROOF"
	FailureNoCompleteRhoFActionLedger          = "FAILED_ROUTE_NO_COMPLETE_RHO_F_ACTION_LEDGER_CERTIFIED"
	FailureNoCompletePackage                   = "FAILED_ROUTE_NO_COMPLETE_RHO_F_J_F_GAMMA_F_D_F_PACKAGE_CERTIFIED"
	FailureWeakSocketSplitNotNative            = "FAILED_ROUTE_H_PLUS_H_MINUS_NOT_NATIVE_H_EIGENSPLIT"
	FailureWeakOrientationNeedsHiggsSeal       = "FAILED_ROUTE_WEAK_SOCKET_SPLIT_REQUIRES_HIGGS_ORIENTATION_SEAL"
	FailureMinimalCarrierClosureNotNative      = "FAILED_ROUTE_MINIMAL_CARRIER_CLOSURE_REMAINS_BLOCK_ACTION_SEAL_NOT_NATIVE_PROOF"
	FailureNoOperatorLevelJOppositeAction      = "FAILED_ROUTE_NO_OPERATOR_LEVEL_J_OPPOSITE_ACTION"
	FailureNoJOppositeCompatibilityProof       = "FAILED_ROUTE_NO_J_OPPOSITE_ACTION_COMPATIBILITY_PROOF_CERTIFIED"
	FailureFirstOrderNotExecutableWithSealData = "FAILED_ROUTE_FIRST_ORDER_CALCULATION_NOT_EXECUTABLE_WITH_SEAL_LEVEL_DATA"
	FailureNoFullFirstOrderConditionProof      = "FAILED_ROUTE_NO_FULL_FIRST_ORDER_CONDITION_PROOF"
	FailureNoBimoduleCommutantProof            = "FAILED_ROUTE_NO_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailureNoOperatorValuedDFMatrix            = "FAILED_ROUTE_NO_OPERATOR_VALUED_D_F_MATRIX_CERTIFIED"
	FailureSymbolicDFNotNative                 = "FAILED_ROUTE_SYMBOLIC_D_F_SUPPORT_MATRIX_NOT_NATIVE_FINITE_TRIPLE"
	FailureKernelStabilityNotCertified         = "FAILED_ROUTE_KERNEL_SINGLETON_STABILITY_NOT_CERTIFIED_WITHOUT_OPERATOR_RHO_F_AND_J_F"
	FailureDFSymNotYukawaMagnitudeSource       = "FAILED_ROUTE_SYMBOLIC_D_F_NOT_YUKAWA_MAGNITUDE_SOURCE"
	FailureNoNumericalYukawaValues             = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureAlphaStillSealed                    = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoTraceMagnitudeReadout             = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoOfficialNEffUpdate                = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate               = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNotR3                               = "FAILED_ROUTE_R2_PLUS_PLUS_PLUS_PLUS_PLUS_DATA_SEAL_COMPATIBILITY_NOT_R3"
	FailureNotR4                               = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
	FailureNoPhysicalNeutrinoTheorem           = "FAILED_ROUTE_NO_PHYSICAL_NEUTRINO_THEOREM"
	FailureNoRightNeutrinoTheorem              = "FAILED_ROUTE_NO_RIGHT_NEUTRINO_OR_STERILE_PARTICLE_THEOREM"
	FailureNoMasslessnessTheorem               = "FAILED_ROUTE_SYMBOLIC_KERNEL_NOT_OBSERVED_MASSLESSNESS_THEOREM"
)

type Ledger struct {
	AlphaB                        float64
	OfficialNEff, OfficialCYukawa float64
	OfficialCHiggs                float64
	OfficialFrozen                bool
	R3, R4, AlphaNative           bool
}

type CarrierClosure struct {
	Carrier, AmbientCell, MinimalBranch string
	HLRank, HRMinRank, HPartMinRank     int
	HFMinRank                           int
	PreservedByBlockActionSeal          bool
	ClosedNatively                      bool
	AbsentCellForcedBackBySchematicRhoF bool
	Supports, Failures                  []string
}

type WeakOrientationAudit struct {
	WeakModule, Split, HAction string
	RankHPlus, RankHMinus      int
	SplitDefinedAtSealLevel    bool
	StableUnderFullHAction     bool
	NativeHEigensplit          bool
	RequiresHiggsOrientation   bool
	PrimaryFragilePoint        bool
	Supports, Failures         []string
}

type JOppositeAudit struct {
	Expression                                        string
	JSealAvailable, OperatorLevelJ, OppositeCertified bool
	CanBuildOppositeCommutator                        bool
	Supports, Failures                                []string
}

type FirstOrderAudit struct {
	Expression                                                    string
	TypedAfterDataSeal, ExecutableNow, Certified                  bool
	HasRhoSeal, HasJSeal, HasGammaSeal, HasDFSupport              bool
	HasOperatorRho, HasOperatorJ, HasOperatorGamma, HasOperatorDF bool
	ObstructionSources                                            []string
	Supports, Failures                                            []string
}

type KernelAudit struct {
	Kernel, RightPuncture                            string
	KernelRank, DFSymKernelDim                       int
	KernelInsideMinimalCarrier, RightPunctureOutside bool
	StableUnderSchematicBlocks, StableUnderFullRhoJ  bool
	PhysicalNeutrinoTheorem, MasslessnessTheorem     bool
	Supports, Failures                               []string
}

type ChiralityAudit struct {
	Expression                                         string
	LeftSign, RightSign                                int
	OddnessSupportLevel, OperatorGamma, KOExtensionSet bool
	Supports, Failures                                 []string
}

type Impact struct {
	Classification                                                                        string
	DataSealInherited, FirstOrderTargetTyped, CompatibilityAudited                        bool
	WeakOrientationIsBlocker, MissingJOppositeIsBlocker, MissingOperatorMatricesIsBlocker bool
	NativeFiniteTriple, FirstOrderCertified, JOppositeCertified, KernelStable             bool
	PhysicalNeutrinoTheorem, MasslessTheorem                                              bool
	AlphaStillSealed, MagnitudesStillMissing                                              bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4      bool
}

type Firewalls struct {
	Enforced                                                                                   bool
	DataSealNotNative, NoCompleteRhoF, NoCompletePackage, WeakOrientationSealOnly              bool
	NoOperatorJ, NoJOppositeProof, FirstOrderNotExecutable, NoFirstOrderProof, NoBimoduleProof bool
	NoOperatorDF, SymbolicDFNotNative, KernelStabilityNotCertified                             bool
	DFSymNotMagnitudeSource, NoNumericalYukawas, AlphaStillSealed, NoTraceMagnitudeReadout     bool
	NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate, NotR3, NotR4                                  bool
	NoPhysicalNeutrino, NoRightNeutrino, NoMasslessness                                        bool
	Verdict                                                                                    string
}

type Audit struct {
	Ledger          Ledger
	CarrierClosure  CarrierClosure
	WeakOrientation WeakOrientationAudit
	JOpposite       JOppositeAudit
	FirstOrder      FirstOrderAudit
	Kernel          KernelAudit
	Chirality       ChiralityAudit
	Impact          Impact
	Firewalls       Firewalls
	Truth           string
	Final           string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		Ledger:          buildLedger(),
		CarrierClosure:  buildCarrierClosure(),
		WeakOrientation: buildWeakOrientation(),
		JOpposite:       buildJOpposite(),
		FirstOrder:      buildFirstOrder(),
		Kernel:          buildKernel(),
		Chirality:       buildChirality(),
		Impact:          buildImpact(),
		Firewalls:       buildFirewalls(),
		Truth:           "Gate 852 follows Gate 851 by auditing the first-order/J-opposite target with the minimal finite-triple data seal in place.  The target is now well typed, but the calculation remains blocked by seal-level rho_F/J_F/D_F data and by the non-native weak h_+/h_- orientation.",
		Final:           "Verdict: first-order and J-opposite compatibility are audited but not certified.  The primary pressure points are weak-doublet orientation, missing operator-level J-opposite action, missing operator-valued rho_F/D_F matrices, and unproven kernel stability.  This remains R2+++++ data-seal compatibility, not R3/R4.",
	}
	return a, validate(a)
}

func validate(a Audit) error {
	if WRank != 4 || HLRank != 8 || RightMinRank != 7 || HPartMinRank != 15 || HFMinRank != 30 || DFSymKernelDim != 1 {
		return err("dimension constants inconsistent")
	}
	if a.CarrierClosure.HPartMinRank != a.CarrierClosure.HLRank+a.CarrierClosure.HRMinRank {
		return err("minimal carrier ranks inconsistent")
	}
	if !a.FirstOrder.TypedAfterDataSeal || a.FirstOrder.ExecutableNow || a.FirstOrder.Certified {
		return err("first-order executable/certification flags inconsistent")
	}
	if !a.WeakOrientation.RequiresHiggsOrientation || a.WeakOrientation.StableUnderFullHAction || a.WeakOrientation.NativeHEigensplit {
		return err("weak-orientation flags inconsistent")
	}
	if a.Kernel.StableUnderFullRhoJ || a.Kernel.PhysicalNeutrinoTheorem || a.Kernel.MasslessnessTheorem {
		return err("kernel overpromoted")
	}
	return nil
}

func buildLedger() Ledger {
	return Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true, R3: false, R4: false, AlphaNative: false}
}

func buildCarrierClosure() CarrierClosure {
	return CarrierClosure{
		Carrier: "H_part^min=H_L plus H_R^min", AmbientCell: "e_+ tensor P_1 outside H_R^min", MinimalBranch: "15/30 active branch inside 16/32 ambient carrier",
		HLRank: HLRank, HRMinRank: RightMinRank, HPartMinRank: HPartMinRank, HFMinRank: HFMinRank,
		PreservedByBlockActionSeal: true, ClosedNatively: false, AbsentCellForcedBackBySchematicRhoF: false,
		Supports: []string{SupportMinimalCarrierClosedByBlockSeal},
		Failures: []string{FailureMinimalCarrierClosureNotNative, FailureDataSealNotNativeFiniteTripleProof},
	}
}

func buildWeakOrientation() WeakOrientationAudit {
	return WeakOrientationAudit{
		WeakModule: "C_L^2", Split: "C_L^2=h_+ plus h_-", HAction: "H acts on C_L^2 as weak-doublet module; generic H action does not natively preserve arbitrary complex rank-one lines",
		RankHPlus: 1, RankHMinus: 1, SplitDefinedAtSealLevel: true, StableUnderFullHAction: false, NativeHEigensplit: false, RequiresHiggsOrientation: true, PrimaryFragilePoint: true,
		Supports: []string{SupportWeakOrientationRequiresHiggsSeal},
		Failures: []string{FailureWeakSocketSplitNotNative, FailureWeakOrientationNeedsHiggsSeal},
	}
}

func buildJOpposite() JOppositeAudit {
	return JOppositeAudit{Expression: "J_F rho_F(a) J_F^{-1} on H_F^min", JSealAvailable: true, OperatorLevelJ: false, OppositeCertified: false, CanBuildOppositeCommutator: false, Supports: []string{SupportFirstOrderTargetWellTyped}, Failures: []string{FailureNoOperatorLevelJOppositeAction, FailureNoJOppositeCompatibilityProof, FailureNoCompletePackage}}
}

func buildFirstOrder() FirstOrderAudit {
	obstructions := []string{"rho_F is schematic, not operator-level", "J_F opposite action is not operator-level", "gamma_F is support-level only", "D_F^sym is a support matrix, not operator-valued", "h_+/h_- weak orientation is a seal, not an H eigensplit", "bimodule/commutant decomposition is absent"}
	return FirstOrderAudit{Expression: "[[D_F,rho_F(a)],J_F rho_F(b) J_F^{-1}]=0", TypedAfterDataSeal: true, ExecutableNow: false, Certified: false, HasRhoSeal: true, HasJSeal: true, HasGammaSeal: true, HasDFSupport: true, HasOperatorRho: false, HasOperatorJ: false, HasOperatorGamma: false, HasOperatorDF: false, ObstructionSources: obstructions, Supports: []string{SupportFirstOrderTargetWellTyped, SupportFirstOrderObstructionClassified}, Failures: []string{FailureFirstOrderNotExecutableWithSealData, FailureNoFullFirstOrderConditionProof, FailureNoCompleteRhoFActionLedger, FailureNoOperatorLevelJOppositeAction, FailureNoBimoduleCommutantProof, FailureNoOperatorValuedDFMatrix}}
}

func buildKernel() KernelAudit {
	return KernelAudit{Kernel: "h_+ tensor P_1", RightPuncture: "e_+ tensor P_1", KernelRank: 1, DFSymKernelDim: DFSymKernelDim, KernelInsideMinimalCarrier: true, RightPunctureOutside: true, StableUnderSchematicBlocks: true, StableUnderFullRhoJ: false, PhysicalNeutrinoTheorem: false, MasslessnessTheorem: false, Supports: []string{SupportKernelStableCandidate}, Failures: []string{FailureKernelStabilityNotCertified, FailureNoPhysicalNeutrinoTheorem, FailureNoRightNeutrinoTheorem, FailureNoMasslessnessTheorem}}
}

func buildChirality() ChiralityAudit {
	return ChiralityAudit{Expression: "gamma_F=+1 on H_L and -1 on H_R^min; D_F^sym is left/right off-diagonal", LeftSign: 1, RightSign: -1, OddnessSupportLevel: true, OperatorGamma: false, KOExtensionSet: false, Supports: []string{SupportChiralityOddnessSupportLevel}, Failures: []string{FailureDataSealNotNativeFiniteTripleProof, FailureNoCompletePackage}}
}

func buildImpact() Impact {
	return Impact{Classification: "R2+++++_data_seal_compatibility_firewall: first-order target typed but not executable/certified", DataSealInherited: true, FirstOrderTargetTyped: true, CompatibilityAudited: true, WeakOrientationIsBlocker: true, MissingJOppositeIsBlocker: true, MissingOperatorMatricesIsBlocker: true, NativeFiniteTriple: false, FirstOrderCertified: false, JOppositeCertified: false, KernelStable: false, PhysicalNeutrinoTheorem: false, MasslessTheorem: false, AlphaStillSealed: true, MagnitudesStillMissing: true, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, CanPromoteToR3: false, CanPromoteToR4: false}
}

func buildFirewalls() Firewalls {
	return Firewalls{Enforced: true, DataSealNotNative: true, NoCompleteRhoF: true, NoCompletePackage: true, WeakOrientationSealOnly: true, NoOperatorJ: true, NoJOppositeProof: true, FirstOrderNotExecutable: true, NoFirstOrderProof: true, NoBimoduleProof: true, NoOperatorDF: true, SymbolicDFNotNative: true, KernelStabilityNotCertified: true, DFSymNotMagnitudeSource: true, NoNumericalYukawas: true, AlphaStillSealed: true, NoTraceMagnitudeReadout: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NotR3: true, NotR4: true, NoPhysicalNeutrino: true, NoRightNeutrino: true, NoMasslessness: true, Verdict: StatusCompatibilityFirewallVerdict}
}

func Statuses() []string {
	return []string{StatusGate851DataSealInherited, StatusFirstOrderTargetTyped, StatusMinimalCarrierClosureAudited, StatusWeakOrientationAudited, StatusJOppositeRequirementAudited, StatusFirstOrderExpressionAudited, StatusKernelStabilityAudited, StatusNoObservedDataUsed, StatusOfficialLedgersFrozen, StatusCompatibilityFirewallVerdict, SupportFirstOrderTargetWellTyped, SupportMinimalCarrierClosedByBlockSeal, SupportWeakOrientationRequiresHiggsSeal, SupportChiralityOddnessSupportLevel, SupportKernelStableCandidate, SupportFirstOrderObstructionClassified, SupportR2CompatibilityFirewallStage, FailureDataSealNotNativeFiniteTripleProof, FailureNoCompleteRhoFActionLedger, FailureNoCompletePackage, FailureWeakSocketSplitNotNative, FailureWeakOrientationNeedsHiggsSeal, FailureMinimalCarrierClosureNotNative, FailureNoOperatorLevelJOppositeAction, FailureNoJOppositeCompatibilityProof, FailureFirstOrderNotExecutableWithSealData, FailureNoFullFirstOrderConditionProof, FailureNoBimoduleCommutantProof, FailureNoOperatorValuedDFMatrix, FailureSymbolicDFNotNative, FailureKernelStabilityNotCertified, FailureDFSymNotYukawaMagnitudeSource, FailureNoNumericalYukawaValues, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNotR3, FailureNotR4, FailureNoPhysicalNeutrinoTheorem, FailureNoRightNeutrinoTheorem, FailureNoMasslessnessTheorem}
}

func FormatLedger(l Ledger) string {
	return join("ledger", []string{"alpha_B=" + f64(l.AlphaB), "official_N_eff=" + f64(l.OfficialNEff), "official_C_Yukawa=" + f64(l.OfficialCYukawa), "official_C_Higgs=" + f64(l.OfficialCHiggs), "official_frozen=" + b(l.OfficialFrozen), "R3=" + b(l.R3), "R4=" + b(l.R4), "alpha_native=" + b(l.AlphaNative)})
}

func FormatCarrierClosure(c CarrierClosure) string {
	return join("carrier_closure", []string{c.Carrier, c.AmbientCell, c.MinimalBranch, "HL_rank=" + i(c.HLRank), "HR_min_rank=" + i(c.HRMinRank), "H_part_min_rank=" + i(c.HPartMinRank), "H_F_min_rank=" + i(c.HFMinRank), "block_action_closure_seal=" + b(c.PreservedByBlockActionSeal), "closed_natively=" + b(c.ClosedNatively), "absent_cell_forced_back=" + b(c.AbsentCellForcedBackBySchematicRhoF), "supports=" + strings.Join(c.Supports, ","), "failures=" + strings.Join(c.Failures, ",")})
}

func FormatWeakOrientation(w WeakOrientationAudit) string {
	return join("weak_orientation", []string{w.WeakModule, w.Split, w.HAction, "rank_h_plus=" + i(w.RankHPlus), "rank_h_minus=" + i(w.RankHMinus), "seal_level=" + b(w.SplitDefinedAtSealLevel), "stable_under_full_H=" + b(w.StableUnderFullHAction), "native_H_eigensplit=" + b(w.NativeHEigensplit), "requires_Higgs_orientation=" + b(w.RequiresHiggsOrientation), "primary_fragile_point=" + b(w.PrimaryFragilePoint), "supports=" + strings.Join(w.Supports, ","), "failures=" + strings.Join(w.Failures, ",")})
}

func FormatJOpposite(j JOppositeAudit) string {
	return join("J_opposite", []string{j.Expression, "J_seal_available=" + b(j.JSealAvailable), "operator_level_J=" + b(j.OperatorLevelJ), "opposite_certified=" + b(j.OppositeCertified), "can_build_opposite_commutator=" + b(j.CanBuildOppositeCommutator), "supports=" + strings.Join(j.Supports, ","), "failures=" + strings.Join(j.Failures, ",")})
}

func FormatFirstOrder(f FirstOrderAudit) string {
	return join("first_order", []string{f.Expression, "typed_after_data_seal=" + b(f.TypedAfterDataSeal), "executable_now=" + b(f.ExecutableNow), "certified=" + b(f.Certified), "rho_seal=" + b(f.HasRhoSeal), "J_seal=" + b(f.HasJSeal), "gamma_seal=" + b(f.HasGammaSeal), "D_F_support=" + b(f.HasDFSupport), "operator_rho=" + b(f.HasOperatorRho), "operator_J=" + b(f.HasOperatorJ), "operator_gamma=" + b(f.HasOperatorGamma), "operator_D_F=" + b(f.HasOperatorDF), "obstructions=" + strings.Join(f.ObstructionSources, ","), "supports=" + strings.Join(f.Supports, ","), "failures=" + strings.Join(f.Failures, ",")})
}

func FormatKernel(k KernelAudit) string {
	return join("kernel", []string{"kernel=" + k.Kernel, "right_puncture=" + k.RightPuncture, "kernel_rank=" + i(k.KernelRank), "D_F_sym_kernel_dim=" + i(k.DFSymKernelDim), "inside_minimal=" + b(k.KernelInsideMinimalCarrier), "right_puncture_outside=" + b(k.RightPunctureOutside), "stable_under_schematic_blocks=" + b(k.StableUnderSchematicBlocks), "stable_under_full_rho_J=" + b(k.StableUnderFullRhoJ), "physical_neutrino=" + b(k.PhysicalNeutrinoTheorem), "masslessness=" + b(k.MasslessnessTheorem), "supports=" + strings.Join(k.Supports, ","), "failures=" + strings.Join(k.Failures, ",")})
}

func FormatChirality(c ChiralityAudit) string {
	return join("chirality", []string{c.Expression, "left_sign=" + i(c.LeftSign), "right_sign=" + i(c.RightSign), "oddness_support_level=" + b(c.OddnessSupportLevel), "operator_gamma=" + b(c.OperatorGamma), "KO_extension_set=" + b(c.KOExtensionSet), "supports=" + strings.Join(c.Supports, ","), "failures=" + strings.Join(c.Failures, ",")})
}

func FormatImpact(im Impact) string {
	return join("impact", []string{im.Classification, "data_seal_inherited=" + b(im.DataSealInherited), "first_order_target_typed=" + b(im.FirstOrderTargetTyped), "compatibility_audited=" + b(im.CompatibilityAudited), "weak_orientation_blocker=" + b(im.WeakOrientationIsBlocker), "missing_J_blocker=" + b(im.MissingJOppositeIsBlocker), "missing_operator_matrices_blocker=" + b(im.MissingOperatorMatricesIsBlocker), "native_finite_triple=" + b(im.NativeFiniteTriple), "first_order=" + b(im.FirstOrderCertified), "J_opposite=" + b(im.JOppositeCertified), "kernel_stable=" + b(im.KernelStable), "physical_neutrino=" + b(im.PhysicalNeutrinoTheorem), "massless_theorem=" + b(im.MasslessTheorem), "alpha_sealed=" + b(im.AlphaStillSealed), "magnitudes_missing=" + b(im.MagnitudesStillMissing), "update_N_eff=" + b(im.CanUpdateNEff), "update_C_Yukawa=" + b(im.CanUpdateCYukawa), "update_C_Higgs=" + b(im.CanUpdateCHiggs), "R3=" + b(im.CanPromoteToR3), "R4=" + b(im.CanPromoteToR4)})
}

func FormatFirewalls(f Firewalls) string {
	return join("firewalls", []string{"enforced=" + b(f.Enforced), "data_seal_not_native=" + b(f.DataSealNotNative), "no_complete_rho=" + b(f.NoCompleteRhoF), "no_complete_package=" + b(f.NoCompletePackage), "weak_orientation_seal_only=" + b(f.WeakOrientationSealOnly), "no_operator_J=" + b(f.NoOperatorJ), "no_J_opposite_proof=" + b(f.NoJOppositeProof), "first_order_not_executable=" + b(f.FirstOrderNotExecutable), "no_first_order_proof=" + b(f.NoFirstOrderProof), "no_bimodule_proof=" + b(f.NoBimoduleProof), "no_operator_D_F=" + b(f.NoOperatorDF), "symbolic_D_F_not_native=" + b(f.SymbolicDFNotNative), "kernel_stability_not_certified=" + b(f.KernelStabilityNotCertified), "D_F_sym_not_magnitude_source=" + b(f.DFSymNotMagnitudeSource), "no_numerical_yukawas=" + b(f.NoNumericalYukawas), "alpha_sealed=" + b(f.AlphaStillSealed), "no_trace_magnitude_readout=" + b(f.NoTraceMagnitudeReadout), "no_official_N_eff_update=" + b(f.NoOfficialNEffUpdate), "no_C_Yukawa_C_Higgs_update=" + b(f.NoCYukawaCHiggsUpdate), "not_R3=" + b(f.NotR3), "not_R4=" + b(f.NotR4), "no_physical_neutrino=" + b(f.NoPhysicalNeutrino), "no_right_neutrino=" + b(f.NoRightNeutrino), "no_masslessness=" + b(f.NoMasslessness), "verdict=" + f.Verdict})
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
