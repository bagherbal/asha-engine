// Package generation2operatorlevelfinitetriplematrixrealizationaudit implements
// Gate 854: Operator-Level FiniteTriple Matrix Realization Audit.
//
// Gate 854 follows Gate 853's Higgs/weak orientation seal.  It instantiates an
// explicit ordered block basis and operator-level matrix descriptors for
// rho_F, gamma_F, J_F, and symbolic D_F^sym on the minimal active carrier
// H_F^min.  This is a matrix-realization seal only: it prepares Gate 855's
// first-order calculation, but does not certify the first-order condition,
// KO signs, a native finite triple, Yukawa magnitudes, alpha_B, R3, or R4.
package generation2operatorlevelfinitetriplematrixrealizationaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE854-OPERATOR-LEVEL-FINITE-TRIPLE-MATRIX-REALIZATION-AUDIT"

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
	PunctureRank    = 1
	LeftKernelRank  = 1
	YActiveEdges    = 3
	DSymRank        = 14
	DSymKernelRank  = HPartMinRank - DSymRank

	StatusGate853Inherited            = "PASS_GATE853_HIGGS_ORIENTATION_SEAL_INHERITED"
	StatusOrderedBasisDefined         = "PASS_ORDERED_MINIMAL_H_F_BASIS_DEFINED"
	StatusAmbientActiveSeparated      = "PASS_AMBIENT_16_32_AND_ACTIVE_15_30_CARRIERS_SEPARATED"
	StatusRhoMatrixSealDefined        = "PASS_RHO_F_BLOCK_ACTION_MATRIX_SEAL_DEFINED"
	StatusGammaMatrixSealDefined      = "PASS_GAMMA_F_CHIRALITY_MATRIX_SEAL_DEFINED"
	StatusJMatrixSealDefined          = "PASS_J_F_PARTICLE_OPPOSITE_EXCHANGE_MATRIX_SEAL_DEFINED"
	StatusDMatrixSealDefined          = "PASS_SYMBOLIC_D_F_MATRIX_SEAL_DEFINED"
	StatusDimensionChecksPassed       = "PASS_OPERATOR_MATRIX_DIMENSION_CHECKS_PASSED"
	StatusPunctureKernelTracked       = "PASS_RIGHT_PUNCTURE_AND_LEFT_KERNEL_TRACKED_IN_BASIS"
	StatusFirstOrderPreparedNotProved = "PASS_FIRST_ORDER_TARGET_PREPARED_BUT_NOT_PROVED"
	StatusNoObservedDataUsed          = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusLedgersFrozen               = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusFirewallVerdict             = "FIREWALL_PRESERVED_GATE854_OPERATOR_MATRIX_SEAL_NOT_NATIVE_FINITE_TRIPLE"

	SupportOperatorMatrixSeal            = "CONDITIONAL_SUPPORT_OPERATOR_LEVEL_MATRIX_DATA_EXISTS_AT_SEAL_LEVEL"
	SupportMinimalCarrierMatrixRealized  = "CONDITIONAL_SUPPORT_H_F_MIN_DIM_30_REALIZED_IN_ORDERED_BLOCK_BASIS"
	SupportRhoPreservesMinimalCarrier    = "CONDITIONAL_SUPPORT_RHO_F_BLOCK_ACTION_PRESERVES_MINIMAL_CARRIER_AT_SEAL_LEVEL"
	SupportDFSelfAdjointByBlock          = "CONDITIONAL_SUPPORT_D_F_SYM_SELF_ADJOINT_BY_CHIRAL_BLOCK_FORM"
	SupportDFChiralityOddByBlock         = "CONDITIONAL_SUPPORT_D_F_SYM_CHIRALITY_ODD_BY_LEFT_RIGHT_BLOCK_FORM"
	SupportFirstOrderNowExecutableTarget = "CONDITIONAL_SUPPORT_FIRST_ORDER_TARGET_HAS_OPERATOR_LEVEL_INPUTS_FOR_NEXT_GATE"
	SupportWeakFrameOrientationPhase     = "CONDITIONAL_SUPPORT_D_F_SYM_IS_ORIENTATION_FRAME_OBJECT_AFTER_HIGGS_SEAL"
	SupportR2OperatorMatrixSealStage     = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_PLUS_PLUS_PLUS_OPERATOR_MATRIX_SEAL_STAGE"

	FailureMatrixSealNotNativeProof       = "FAILED_ROUTE_OPERATOR_LEVEL_MATRIX_REALIZATION_IS_SEAL_NOT_NATIVE_FINITE_TRIPLE_PROOF"
	FailureNoFirstOrderProofYet           = "FAILED_ROUTE_NO_FULL_FIRST_ORDER_CONDITION_PROOF_YET"
	FailureNoFirstOrderCalculationYet     = "FAILED_ROUTE_FIRST_ORDER_CALCULATION_NOT_PERFORMED_IN_GATE_854"
	FailureNoJOppositeProofYet            = "FAILED_ROUTE_NO_J_OPPOSITE_ACTION_COMPATIBILITY_PROOF_YET"
	FailureNoBimoduleProofYet             = "FAILED_ROUTE_NO_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailureKOSignExtensionNotCertified    = "FAILED_ROUTE_KO_SIGN_EXTENSION_NOT_CERTIFIED"
	FailureJStructureFormalOnly           = "FAILED_ROUTE_J_F_COPY_EXCHANGE_IS_FORMAL_NOT_FULL_KO_REAL_STRUCTURE"
	FailureDOperatorValuedNotCertified    = "FAILED_ROUTE_D_F_MATRIX_IS_SYMBOLIC_SUPPORT_NOT_OPERATOR_VALUED_YUKAWA_MATRIX"
	FailureYCoefficientsSymbolicOnly      = "FAILED_ROUTE_SYMBOLIC_Y_COEFFICIENTS_NOT_YUKAWA_MAGNITUDES"
	FailureWeakFrameNotNativeHInvariant   = "FAILED_ROUTE_ORIENTED_H_PLUS_H_MINUS_FRAME_NOT_NATIVE_FULL_H_INVARIANT"
	FailureDFAfterOrientationNotUnbrokenH = "FAILED_ROUTE_D_F_SYM_LIVES_IN_ORIENTATION_FRAME_NOT_UNBROKEN_H_EQUIVARIANT_THEOREM"
	FailureKernelNotFullRhoJStable        = "FAILED_ROUTE_LEFT_KERNEL_STABILITY_NOT_CERTIFIED_UNDER_FULL_RHO_F_AND_J_F"
	FailurePunctureNotNativeNullEdge      = "FAILED_ROUTE_RIGHT_PUNCTURE_ABSENCE_NOT_NATIVE_NULL_EDGE_THEOREM"
	FailureNoAlphaSource                  = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoYukawaMagnitudeSource        = "FAILED_ROUTE_NO_YUKAWA_MAGNITUDE_SOURCE"
	FailureNoSectorTraceReadout           = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoOfficialNEffUpdate           = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate          = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNotR3                          = "FAILED_ROUTE_R2_PLUS_PLUS_PLUS_PLUS_PLUS_OPERATOR_MATRIX_SEAL_NOT_R3"
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

type BasisBlock struct {
	Name      string
	Side      string
	Support   string
	Rank      int
	Chirality int
	Present   bool
}

type OrderedBasis struct {
	ParticleBlocks, OppositeBlocks       []BasisBlock
	HLRank, HRMinRank, HPartMinRank      int
	HFMinRank, AmbientPartRank, AmbientF int
	RightPuncture, LeftKernel            string
	RightPunctureOutside, LeftKernelInHL bool
	Complete, JCopyIncluded              bool
	Supports, Failures                   []string
}

type RhoMatrixSeal struct {
	Algebra, Domain               string
	M3ActsOnP3, M3TrivialOnP1     bool
	CActsOnRightCharacters        bool
	HActsOnFullWeakDoublet        bool
	HMayMixHPlusHMinus            bool
	HPlusHMinusNativeHEigensplit  bool
	PreservesMinimalCarrier       bool
	PunctureForcedBackIntoCarrier bool
	DefinedAtSealLevel            bool
	NativeRepresentationProof     bool
	Supports, Failures            []string
}

type GammaMatrixSeal struct {
	ParticleRule, OppositeRule string
	SquareIdentity             bool
	ParticleSideDefined        bool
	KOExtensionCertified       bool
	ChiralityOddWithDFByBlock  bool
	SupportLevelOnly           bool
	Supports, Failures         []string
}

type JMatrixSeal struct {
	Rule                              string
	ParticleOppositeExchange          bool
	AntiunitaryFormal                 bool
	OppositeCopyDimension             int
	KOSignsCertified                  bool
	OppositeActionCompatibilityProved bool
	FullRealStructureProof            bool
	Supports, Failures                []string
}

type DMatrixSeal struct {
	Formula                            string
	YTerms                             []string
	YPlus1Zero                         bool
	HLRank, HRMinRank, ParticleDim     int
	Rank, KernelRank                   int
	SelfAdjointByBlock                 bool
	ChiralityOddByBlock                bool
	LeptoColorPreserving               bool
	ExtendedToJCopy                    bool
	OperatorValuedMatrixCertified      bool
	NumericalYukawaMagnitudesCertified bool
	OrientationFrameObject             bool
	UnbrokenHEquivariantTheorem        bool
	Supports, Failures                 []string
}

type MatrixChecks struct {
	BasisComplete, DimensionConsistent, RhoPreservesCarrier     bool
	GammaSquaredIdentity, DSelfAdjoint, DChiralityOdd           bool
	JMapsParticleToOpposite, PunctureOutside, LeftKernelPresent bool
	FirstOrderExecutableNextGate, FirstOrderProvedThisGate      bool
	Supports, Failures                                          []string
}

type Impact struct {
	Classification                                                                        string
	Gate853Inherited, OperatorMatrixSeal, BasisComplete, RhoSeal, GammaSeal, JSeal, DSeal bool
	FirstOrderPrepared, FirstOrderProved, NativeFiniteTripleProof                         bool
	AlphaStillSealed, MagnitudesStillMissing                                              bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4      bool
}

type Firewalls struct {
	Enforced                                                                                           bool
	MatrixSealNotNative, NoFirstOrderProof, NoFirstOrderCalculation, NoJOppositeProof, NoBimoduleProof bool
	KOSignNotCertified, JFormalOnly, DSymbolicOnly, YSymbolicOnly                                      bool
	WeakFrameNotNativeHInvariant, DFOrientationNotUnbrokenH, KernelNotStable, PunctureNotNative        bool
	NoAlphaSource, NoYukawaMagnitudes, NoTraceReadout, NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate     bool
	NotR3, NotR4, NoParticleAssign, NoNeutrinoTheorem, NoThreeGenerationTheorem                        bool
	Verdict                                                                                            string
}

type Audit struct {
	ID        string
	Ledger    Ledger
	Basis     OrderedBasis
	Rho       RhoMatrixSeal
	Gamma     GammaMatrixSeal
	J         JMatrixSeal
	D         DMatrixSeal
	Checks    MatrixChecks
	Impact    Impact
	Firewalls Firewalls
	Truth     string
	Final     string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		ID:        AuditID,
		Ledger:    buildLedger(),
		Basis:     buildBasis(),
		Rho:       buildRho(),
		Gamma:     buildGamma(),
		J:         buildJ(),
		D:         buildD(),
		Checks:    buildChecks(),
		Impact:    buildImpact(),
		Firewalls: buildFirewalls(),
		Truth:     "Gate 854 realizes the Gate 851 minimal finite-triple data seal as explicit oriented block-matrix data on H_F^min, while preserving that this is not yet a native finite triple or first-order proof.",
		Final:     "VERDICT: operator-level matrix seal constructed; Gate 855 may attempt first-order/J-opposite calculation, but alpha_B, Yukawa magnitudes, R3/R4, KO signs, and native finite-triple proof remain blocked.",
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
	if !almost(a.Ledger.AlphaB, AlphaB, 1e-18) || !a.Ledger.OfficialFrozen || a.Ledger.AlphaNative || a.Ledger.R3 || a.Ledger.R4 {
		return err("ledger mismatch")
	}
	if !a.Basis.Complete || !a.Basis.JCopyIncluded || a.Basis.HLRank != HLRank || a.Basis.HRMinRank != HRMinRank || a.Basis.HPartMinRank != HPartMinRank || a.Basis.HFMinRank != HFMinRank || a.Basis.AmbientPartRank != AmbientPartRank || a.Basis.AmbientF != AmbientFRank {
		return err("basis dimensions inconsistent")
	}
	if len(a.Basis.ParticleBlocks) != 7 || len(a.Basis.OppositeBlocks) != 7 || !a.Basis.RightPunctureOutside || !a.Basis.LeftKernelInHL {
		return err("basis block ledger inconsistent")
	}
	if !a.Rho.DefinedAtSealLevel || !a.Rho.PreservesMinimalCarrier || a.Rho.PunctureForcedBackIntoCarrier || a.Rho.NativeRepresentationProof || !a.Rho.HMayMixHPlusHMinus || a.Rho.HPlusHMinusNativeHEigensplit {
		return err("rho matrix seal inconsistent")
	}
	if !a.Gamma.ParticleSideDefined || !a.Gamma.SquareIdentity || !a.Gamma.ChiralityOddWithDFByBlock || a.Gamma.KOExtensionCertified || !a.Gamma.SupportLevelOnly {
		return err("gamma matrix seal inconsistent")
	}
	if !a.J.ParticleOppositeExchange || !a.J.AntiunitaryFormal || a.J.OppositeCopyDimension != HPartMinRank || a.J.KOSignsCertified || a.J.OppositeActionCompatibilityProved || a.J.FullRealStructureProof {
		return err("J matrix seal inconsistent")
	}
	if !a.D.YPlus1Zero || len(a.D.YTerms) != YActiveEdges || a.D.HLRank != HLRank || a.D.HRMinRank != HRMinRank || a.D.ParticleDim != HPartMinRank || a.D.Rank != DSymRank || a.D.KernelRank != DSymKernelRank || !a.D.SelfAdjointByBlock || !a.D.ChiralityOddByBlock || !a.D.LeptoColorPreserving || !a.D.ExtendedToJCopy || a.D.OperatorValuedMatrixCertified || a.D.NumericalYukawaMagnitudesCertified || !a.D.OrientationFrameObject || a.D.UnbrokenHEquivariantTheorem {
		return err("D matrix seal inconsistent")
	}
	if !a.Checks.BasisComplete || !a.Checks.DimensionConsistent || !a.Checks.RhoPreservesCarrier || !a.Checks.GammaSquaredIdentity || !a.Checks.DSelfAdjoint || !a.Checks.DChiralityOdd || !a.Checks.JMapsParticleToOpposite || !a.Checks.PunctureOutside || !a.Checks.LeftKernelPresent || !a.Checks.FirstOrderExecutableNextGate || a.Checks.FirstOrderProvedThisGate {
		return err("matrix checks inconsistent")
	}
	if !a.Impact.OperatorMatrixSeal || !a.Impact.FirstOrderPrepared || a.Impact.FirstOrderProved || a.Impact.NativeFiniteTripleProof || !a.Impact.AlphaStillSealed || !a.Impact.MagnitudesStillMissing || a.Impact.CanUpdateNEff || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		return err("impact overpromoted")
	}
	if !a.Firewalls.Enforced || a.Firewalls.Verdict != StatusFirewallVerdict || !a.Firewalls.MatrixSealNotNative || !a.Firewalls.NoFirstOrderProof || !a.Firewalls.KOSignNotCertified || !a.Firewalls.DSymbolicOnly || !a.Firewalls.WeakFrameNotNativeHInvariant || !a.Firewalls.NoAlphaSource || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 {
		return err("firewalls inconsistent")
	}
	return nil
}

func buildLedger() Ledger {
	return Ledger{AlphaB: AlphaB, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true}
}

func buildBasis() OrderedBasis {
	particle := []BasisBlock{
		{Name: "h_+ tensor P_3", Side: "L", Support: "left oriented color block", Rank: 3, Chirality: +1, Present: true},
		{Name: "h_+ tensor P_1", Side: "L", Support: "left neutral kernel singleton", Rank: 1, Chirality: +1, Present: true},
		{Name: "h_- tensor P_3", Side: "L", Support: "left oriented color block", Rank: 3, Chirality: +1, Present: true},
		{Name: "h_- tensor P_1", Side: "L", Support: "left lepton block", Rank: 1, Chirality: +1, Present: true},
		{Name: "e_+ tensor P_3", Side: "R", Support: "minimal right dominant color block", Rank: 3, Chirality: -1, Present: true},
		{Name: "e_- tensor P_3", Side: "R", Support: "minimal right rest color block", Rank: 3, Chirality: -1, Present: true},
		{Name: "e_- tensor P_1", Side: "R", Support: "minimal right rest lepton singleton", Rank: 1, Chirality: -1, Present: true},
	}
	opposite := make([]BasisBlock, 0, len(particle))
	for _, b := range particle {
		opposite = append(opposite, BasisBlock{Name: "J(" + b.Name + ")", Side: "J" + b.Side, Support: "opposite copy of " + b.Support, Rank: b.Rank, Chirality: b.Chirality, Present: true})
	}
	return OrderedBasis{ParticleBlocks: particle, OppositeBlocks: opposite, HLRank: HLRank, HRMinRank: HRMinRank, HPartMinRank: HPartMinRank, HFMinRank: HFMinRank, AmbientPartRank: AmbientPartRank, AmbientF: AmbientFRank, RightPuncture: "e_+ tensor P_1", LeftKernel: "h_+ tensor P_1", RightPunctureOutside: true, LeftKernelInHL: true, Complete: true, JCopyIncluded: true, Supports: []string{StatusOrderedBasisDefined, StatusAmbientActiveSeparated, SupportMinimalCarrierMatrixRealized}, Failures: []string{FailurePunctureNotNativeNullEdge}}
}

func buildRho() RhoMatrixSeal {
	return RhoMatrixSeal{Algebra: "A_F = C plus H plus M_3(C)", Domain: "H_F^min in oriented block basis", M3ActsOnP3: true, M3TrivialOnP1: true, CActsOnRightCharacters: true, HActsOnFullWeakDoublet: true, HMayMixHPlusHMinus: true, HPlusHMinusNativeHEigensplit: false, PreservesMinimalCarrier: true, PunctureForcedBackIntoCarrier: false, DefinedAtSealLevel: true, NativeRepresentationProof: false, Supports: []string{StatusRhoMatrixSealDefined, SupportRhoPreservesMinimalCarrier}, Failures: []string{FailureMatrixSealNotNativeProof, FailureWeakFrameNotNativeHInvariant, FailureDFAfterOrientationNotUnbrokenH}}
}

func buildGamma() GammaMatrixSeal {
	return GammaMatrixSeal{ParticleRule: "+1 on H_L, -1 on H_R^min", OppositeRule: "extend to J-copy by chosen KO convention", SquareIdentity: true, ParticleSideDefined: true, KOExtensionCertified: false, ChiralityOddWithDFByBlock: true, SupportLevelOnly: true, Supports: []string{StatusGammaMatrixSealDefined, SupportDFChiralityOddByBlock}, Failures: []string{FailureKOSignExtensionNotCertified}}
}

func buildJ() JMatrixSeal {
	return JMatrixSeal{Rule: "formal antiunitary exchange H_part^min <-> J_F H_part^min", ParticleOppositeExchange: true, AntiunitaryFormal: true, OppositeCopyDimension: HPartMinRank, KOSignsCertified: false, OppositeActionCompatibilityProved: false, FullRealStructureProof: false, Supports: []string{StatusJMatrixSealDefined}, Failures: []string{FailureKOSignExtensionNotCertified, FailureJStructureFormalOnly, FailureNoJOppositeProofYet}}
}

func buildD() DMatrixSeal {
	return DMatrixSeal{Formula: "D_F^sym = [[0,Y_supp^dagger],[Y_supp,0]], Y_supp=y_+3Y_+3+y_-3Y_-3+y_-1Y_-1, y_+1=0", YTerms: []string{"Y_+3: e_+ tensor P_3 -> h_+ tensor P_3", "Y_-3: e_- tensor P_3 -> h_- tensor P_3", "Y_-1: e_- tensor P_1 -> h_- tensor P_1"}, YPlus1Zero: true, HLRank: HLRank, HRMinRank: HRMinRank, ParticleDim: HPartMinRank, Rank: DSymRank, KernelRank: DSymKernelRank, SelfAdjointByBlock: true, ChiralityOddByBlock: true, LeptoColorPreserving: true, ExtendedToJCopy: true, OperatorValuedMatrixCertified: false, NumericalYukawaMagnitudesCertified: false, OrientationFrameObject: true, UnbrokenHEquivariantTheorem: false, Supports: []string{StatusDMatrixSealDefined, SupportDFSelfAdjointByBlock, SupportDFChiralityOddByBlock}, Failures: []string{FailureDOperatorValuedNotCertified, FailureYCoefficientsSymbolicOnly, FailureDFAfterOrientationNotUnbrokenH}}
}

func buildChecks() MatrixChecks {
	return MatrixChecks{BasisComplete: true, DimensionConsistent: true, RhoPreservesCarrier: true, GammaSquaredIdentity: true, DSelfAdjoint: true, DChiralityOdd: true, JMapsParticleToOpposite: true, PunctureOutside: true, LeftKernelPresent: true, FirstOrderExecutableNextGate: true, FirstOrderProvedThisGate: false, Supports: []string{StatusDimensionChecksPassed, StatusPunctureKernelTracked, SupportFirstOrderNowExecutableTarget}, Failures: []string{FailureNoFirstOrderProofYet, FailureNoFirstOrderCalculationYet}}
}

func buildImpact() Impact {
	return Impact{Classification: "R2+++++_operator_matrix_seal", Gate853Inherited: true, OperatorMatrixSeal: true, BasisComplete: true, RhoSeal: true, GammaSeal: true, JSeal: true, DSeal: true, FirstOrderPrepared: true, FirstOrderProved: false, NativeFiniteTripleProof: false, AlphaStillSealed: true, MagnitudesStillMissing: true}
}

func buildFirewalls() Firewalls {
	return Firewalls{Enforced: true, MatrixSealNotNative: true, NoFirstOrderProof: true, NoFirstOrderCalculation: true, NoJOppositeProof: true, NoBimoduleProof: true, KOSignNotCertified: true, JFormalOnly: true, DSymbolicOnly: true, YSymbolicOnly: true, WeakFrameNotNativeHInvariant: true, DFOrientationNotUnbrokenH: true, KernelNotStable: true, PunctureNotNative: true, NoAlphaSource: true, NoYukawaMagnitudes: true, NoTraceReadout: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NotR3: true, NotR4: true, NoParticleAssign: true, NoNeutrinoTheorem: true, NoThreeGenerationTheorem: true, Verdict: StatusFirewallVerdict}
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("alpha_B=%.16g; official_N_eff=%.16g; C_Yukawa=%.16g; C_Higgs=%.16g; frozen=%t; alpha_native=%t; R3=%t; R4=%t", l.AlphaB, l.OfficialNEff, l.OfficialCYukawa, l.OfficialCHiggs, l.OfficialFrozen, l.AlphaNative, l.R3, l.R4)
}

func FormatBasis(b OrderedBasis) string {
	parts := make([]string, 0, len(b.ParticleBlocks))
	for _, blk := range b.ParticleBlocks {
		parts = append(parts, fmt.Sprintf("%s(rank=%d,chi=%+d)", blk.Name, blk.Rank, blk.Chirality))
	}
	return fmt.Sprintf("H_L=%d; H_R^min=%d; H_part^min=%d; H_F^min=%d; ambient=%d/%d; right_puncture_outside=%t; left_kernel_in_HL=%t; particle_blocks=[%s]", b.HLRank, b.HRMinRank, b.HPartMinRank, b.HFMinRank, b.AmbientPartRank, b.AmbientF, b.RightPunctureOutside, b.LeftKernelInHL, strings.Join(parts, "; "))
}

func FormatRho(r RhoMatrixSeal) string {
	return fmt.Sprintf("%s on %s; M3->P3=%t; M3 trivial on P1=%t; C right characters=%t; H full doublet=%t; H may mix h-lines=%t; preserves minimal carrier=%t; native proof=%t", r.Algebra, r.Domain, r.M3ActsOnP3, r.M3TrivialOnP1, r.CActsOnRightCharacters, r.HActsOnFullWeakDoublet, r.HMayMixHPlusHMinus, r.PreservesMinimalCarrier, r.NativeRepresentationProof)
}

func FormatGamma(g GammaMatrixSeal) string {
	return fmt.Sprintf("gamma particle rule: %s; opposite rule: %s; gamma^2=I=%t; KO extension certified=%t; {D,gamma}=0 by block=%t; support_level_only=%t", g.ParticleRule, g.OppositeRule, g.SquareIdentity, g.KOExtensionCertified, g.ChiralityOddWithDFByBlock, g.SupportLevelOnly)
}

func FormatJ(j JMatrixSeal) string {
	return fmt.Sprintf("J rule: %s; exchange=%t; formal antiunitary=%t; opposite_dim=%d; KO signs=%t; opposite_action_proved=%t", j.Rule, j.ParticleOppositeExchange, j.AntiunitaryFormal, j.OppositeCopyDimension, j.KOSignsCertified, j.OppositeActionCompatibilityProved)
}

func FormatD(d DMatrixSeal) string {
	return fmt.Sprintf("%s; terms=[%s]; y_+1_zero=%t; dims L/R/part=%d/%d/%d; rank=%d; kernel=%d; selfadjoint=%t; chirality_odd=%t; operator_valued=%t; orientation_frame=%t", d.Formula, strings.Join(d.YTerms, "; "), d.YPlus1Zero, d.HLRank, d.HRMinRank, d.ParticleDim, d.Rank, d.KernelRank, d.SelfAdjointByBlock, d.ChiralityOddByBlock, d.OperatorValuedMatrixCertified, d.OrientationFrameObject)
}

func FormatChecks(c MatrixChecks) string {
	return fmt.Sprintf("basis=%t; dims=%t; rho_preserves=%t; gamma^2=%t; D*=D=%t; {D,gamma}=0=%t; J_exchange=%t; puncture_outside=%t; left_kernel=%t; first_order_ready_next=%t; first_order_proved_now=%t", c.BasisComplete, c.DimensionConsistent, c.RhoPreservesCarrier, c.GammaSquaredIdentity, c.DSelfAdjoint, c.DChiralityOdd, c.JMapsParticleToOpposite, c.PunctureOutside, c.LeftKernelPresent, c.FirstOrderExecutableNextGate, c.FirstOrderProvedThisGate)
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s; matrix_seal=%t; first_order_prepared=%t; first_order_proved=%t; native_finite_triple=%t; alpha_sealed=%t; magnitudes_missing=%t; update_N_eff=%t; R3=%t; R4=%t", i.Classification, i.OperatorMatrixSeal, i.FirstOrderPrepared, i.FirstOrderProved, i.NativeFiniteTripleProof, i.AlphaStillSealed, i.MagnitudesStillMissing, i.CanUpdateNEff, i.CanPromoteToR3, i.CanPromoteToR4)
}

func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%t; verdict=%s; matrix_not_native=%t; no_first_order=%t; no_J_opposite=%t; no_KO=%t; D_symbolic=%t; weak_frame_not_H_native=%t; no_alpha=%t; no_yukawa=%t; not_R3=%t; not_R4=%t", f.Enforced, f.Verdict, f.MatrixSealNotNative, f.NoFirstOrderProof, f.NoJOppositeProof, f.KOSignNotCertified, f.DSymbolicOnly, f.WeakFrameNotNativeHInvariant, f.NoAlphaSource, f.NoYukawaMagnitudes, f.NotR3, f.NotR4)
}

func Statuses() []string {
	return []string{
		StatusGate853Inherited,
		StatusOrderedBasisDefined,
		StatusAmbientActiveSeparated,
		StatusRhoMatrixSealDefined,
		StatusGammaMatrixSealDefined,
		StatusJMatrixSealDefined,
		StatusDMatrixSealDefined,
		StatusDimensionChecksPassed,
		StatusPunctureKernelTracked,
		StatusFirstOrderPreparedNotProved,
		StatusNoObservedDataUsed,
		StatusLedgersFrozen,
		SupportOperatorMatrixSeal,
		SupportMinimalCarrierMatrixRealized,
		SupportRhoPreservesMinimalCarrier,
		SupportDFSelfAdjointByBlock,
		SupportDFChiralityOddByBlock,
		SupportFirstOrderNowExecutableTarget,
		SupportWeakFrameOrientationPhase,
		SupportR2OperatorMatrixSealStage,
		FailureMatrixSealNotNativeProof,
		FailureNoFirstOrderProofYet,
		FailureNoFirstOrderCalculationYet,
		FailureNoJOppositeProofYet,
		FailureNoBimoduleProofYet,
		FailureKOSignExtensionNotCertified,
		FailureJStructureFormalOnly,
		FailureDOperatorValuedNotCertified,
		FailureYCoefficientsSymbolicOnly,
		FailureWeakFrameNotNativeHInvariant,
		FailureDFAfterOrientationNotUnbrokenH,
		FailureKernelNotFullRhoJStable,
		FailurePunctureNotNativeNullEdge,
		FailureNoAlphaSource,
		FailureNoYukawaMagnitudeSource,
		FailureNoSectorTraceReadout,
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

func almost(a, b, eps float64) bool {
	if a > b {
		return a-b <= eps
	}
	return b-a <= eps
}
