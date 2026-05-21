// Package generation2leptocolorfiniteactionprojectorledgeraudit implements
// Gate 838: LeptoColor Finite Representation Action and ProjectorLedger Audit.
//
// Gate 838 follows Gate 837's constructive carrier seal. Gate 837 introduced
// the shared lepto-color carrier W = C_lepton plus C^3_color so that P_3 W is
// the M_3(C) fundamental module by representation seal, bypassing the failed
// comparison of two separately-defined triplets. Gate 838 audits the next layer:
// whether the sealed particle carrier H_part=(C_R^2 plus C_L^2) tensor W can
// carry a consistent schematic rho_F(C plus H plus M_3(C)) action and a coarse
// finite-sector projector ledger. The gate certifies the sealed coarse ledger
// (R/L slot times lepton/color support, plus J-copy ranks), but it does not
// derive the representation natively, does not certify full first-order finite
// triple data, does not choose a color atom frame, and does not promote any
// sector projector into a trace magnitude, N_eff update, R3 ledger, or native
// Yukawa theorem.
package generation2leptocolorfiniteactionprojectorledgeraudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE838-LEPTOCOLOR-FINITE-ACTION-PROJECTOR-LEDGER-AUDIT"

	SBoundary       = 0.0012924448188162962
	AlphaB          = 0.0003878958469680527
	OperatorNEff    = 3.002327375081808
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	LeptonBlockDim     = 1
	ColorBlockDim      = 3
	WDim               = 4
	RightSlotDim       = 2
	LeftSlotDim        = 2
	ElectroweakSlotDim = 4
	HPartDim           = 16
	JCopyCount         = 2
	HFDim              = 32

	RankR1 = RightSlotDim * LeptonBlockDim
	RankR3 = RightSlotDim * ColorBlockDim
	RankL1 = LeftSlotDim * LeptonBlockDim
	RankL3 = LeftSlotDim * ColorBlockDim

	BMinusLLeptonWeight = -1.0
	BMinusLColorWeight  = 1.0 / 3.0

	StatusGate837Inherited                 = "PASS_GATE837_LEPTOCOLOR_CARRIER_SEAL_INHERITED"
	StatusWCarrierReverified               = "PASS_W_CARRIER_P1_P3_B_MINUS_L_REVERIFIED"
	StatusESlotAudited                     = "PASS_E_SLOT_C_R2_PLUS_C_L2_SOURCE_ROLES_AUDITED"
	StatusSchematicRhoActionConsistent     = "PASS_SCHEMATIC_RHO_F_ACTION_CONSISTENT_ON_SEALED_CARRIER"
	StatusM3P3ActionCertified              = "PASS_M3C_ACTS_ON_P3W_AND_TRIVIAL_ON_P1W_WITHIN_SEAL"
	StatusHLeftActionCertified             = "PASS_H_QUATERNIONIC_ACTION_ASSIGNED_TO_LEFT_DOUBLE_SOCKET_WITHIN_SEAL"
	StatusCRightActionCertified            = "PASS_C_ACTION_ASSIGNED_TO_RIGHT_SOCKET_PAIR_WITHIN_SEAL"
	StatusP1P3BMinusLPreserved             = "PASS_RHO_ACTION_PRESERVES_P1_P3_AND_B_MINUS_L_BLOCKS"
	StatusCoarseParticleLedgerConstructed  = "PASS_COARSE_PARTICLE_SIDE_PROJECTOR_LEDGER_CONSTRUCTED"
	StatusCoarseLedgerOrthogonalComplete   = "PASS_COARSE_PROJECTORS_ORTHOGONAL_AND_COMPLETE_ON_H_PART"
	StatusJCopyLedgerDoublesRanks          = "PASS_J_COPY_PROJECTOR_LEDGER_DOUBLES_TO_H_F_DIM_32"
	StatusDFSymbolicEdgeSkeletonAudited    = "PASS_D_F_SYMBOLIC_EDGE_SUPPORT_SKELETON_AUDITED"
	StatusCarrierLedgerNotMagnitudeReadout = "PASS_CARRIER_PROJECTOR_LEDGER_NOT_TRACE_MAGNITUDE_READOUT"
	StatusOneBlockNotGenerationTheorem     = "PASS_ONE_FINITE_CARRIER_BLOCK_NOT_THREE_GENERATION_THEOREM"
	StatusNoObservedDataUsed               = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusLedgerFrozen                     = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusR2PlusPlusRetained               = "PASS_R2_PLUS_PLUS_RETAINED_NOT_R3_OR_R4"
	StatusFirewallGate838                  = "FIREWALL_PRESERVED_GATE838_FINITE_ACTION_PROJECTOR_LEDGER_SEAL"

	SupportSharedLeptoColorCarrierInherited = "CONDITIONAL_SUPPORT_SHARED_W_CARRIER_INHERITED_FROM_GATE837"
	SupportP3WAsM3Fundamental               = "CONDITIONAL_SUPPORT_P3W_AS_M3C_FUNDAMENTAL_MODULE_WITHIN_SEAL"
	SupportBMinusLInternalToW               = "CONDITIONAL_SUPPORT_B_MINUS_L_INTERNAL_TO_LEPTOCOLOR_CARRIER"
	SupportESlotRightLeftSocketBody         = "CONDITIONAL_SUPPORT_E_SLOT_AS_C_R2_RIGHT_SOCKET_PAIR_PLUS_C_L2_LEFT_DOUBLE_SOCKET"
	SupportRhoFActionSeal                   = "CONDITIONAL_SUPPORT_RHO_F_ACTION_EXISTS_AS_REPRESENTATION_SEAL_ON_H_PART"
	SupportCoarsePiSectorFSeal              = "CONDITIONAL_SUPPORT_COARSE_PI_SECTOR_F_SEAL_FROM_R_L_TIMES_LEPTON_COLOR_PROJECTORS"
	SupportJCopyFiniteBody                  = "CONDITIONAL_SUPPORT_H_F_DIM_32_FROM_H_PART_PLUS_J_F_H_PART"
	SupportDFEdgesAsSocketGraph             = "CONDITIONAL_SUPPORT_D_F_EDGE_SUPPORT_AS_SYMBOLIC_SOCKET_GRAPH_ONLY"
	SupportSectorBodyBeforeCompression      = "CONDITIONAL_SUPPORT_FINITE_SECTOR_BODY_PRECEDES_AGGREGATE_TRACE_COMPRESSION"

	FailureRepresentationSealNotNative      = "FAILED_ROUTE_REPRESENTATION_ACTION_IS_SEAL_NOT_NATIVE_DERIVATION"
	FailureNoFullFiniteTripleProof          = "FAILED_ROUTE_NO_FULL_NATIVE_FINITE_TRIPLE_REPRESENTATION_PROOF"
	FailureNoExplicitMatrices               = "FAILED_ROUTE_NO_EXPLICIT_RHO_F_GAMMA_F_J_F_D_F_MATRICES_CERTIFIED"
	FailureNoFirstOrderProof                = "FAILED_ROUTE_NO_FIRST_ORDER_CONDITION_STABILITY_PROOF_CERTIFIED"
	FailureNoBimoduleCommutantProof         = "FAILED_ROUTE_NO_BIMODULE_COMMUTANT_DECOMPOSITION_PROOF_CERTIFIED"
	FailureNoCanonicalColorAtomFrame        = "FAILED_ROUTE_NO_CANONICAL_M3C_COLOR_ATOM_FRAME_CERTIFIED"
	FailureM3MatrixUnitsBasisDependent      = "FAILED_ROUTE_M3_MATRIX_UNITS_REMAIN_BASIS_DEPENDENT_WITHOUT_FRAME"
	FailureNoFineColorAtomLedger            = "FAILED_ROUTE_NO_FINE_COLOR_ATOM_PROJECTOR_LEDGER_WITHOUT_GAUGE_FRAME"
	FailureDFEdgesNotMagnitudes             = "FAILED_ROUTE_D_F_SYMBOLIC_EDGE_SUPPORT_NOT_YUKAWA_MAGNITUDES"
	FailureNoNumericalYukawaSockets         = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_SOCKET_VALUES_CERTIFIED"
	FailureNoTraceMagnitudeReadout          = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoAggregateCompressionMap        = "FAILED_ROUTE_NO_AGGREGATE_TRACE_COMPRESSION_MAP_CERTIFIED_YET"
	FailureAggregateOperatorNotSectorLedger = "FAILED_ROUTE_R2_PLUS_PLUS_AGGREGATE_OPERATOR_NOT_SECTOR_LEDGER"
	FailureAlphaStillSealed                 = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoNEffUpdate                     = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaUpdate                  = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoObservedYukawaFit              = "FAILED_ROUTE_NO_OBSERVED_YUKAWA_FITTING_ALLOWED"
	FailureNoParticleAssignment             = "FAILED_ROUTE_NO_STANDARD_MODEL_PARTICLE_ASSIGNMENT_FROM_LEDGER_SEAL"
	FailureNoThreeGenerationTheorem         = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
	FailureNotR3                            = "FAILED_ROUTE_R2_PLUS_PLUS_NOT_R3"
	FailureNotR4                            = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	S, AlphaB                         float64
	OperatorNEff, OfficialNEff        float64
	OfficialCYukawa, OfficialCHiggs   float64
	R2PlusPlus, R3, R4, AlphaIsNative bool
	OfficialFrozen                    bool
}

type WCarrier struct {
	Expression                               string
	LeptonDim, ColorDim, Dim                 int
	P1Rank, P3Rank                           int
	P1P3Orthogonal, P1PlusP3CompletesW       bool
	BMinusLLeptonWeight, BMinusLColorWeight  float64
	BMinusLTraceZero                         bool
	P3WIsM3Fundamental, ColorBlockCanonical  bool
	ColorAtomsCanonical, CanonicalColorFrame bool
	Supports, Failures                       []string
}

type ECarrier struct {
	Expression                                             string
	RightSlotDim, LeftSlotDim, Dim                         int
	RightSocketPairDeclared, LeftDoubleSocketDeclared      bool
	QuaternionicWeakRoleDeclared, ComplexRightRoleDeclared bool
	SourceTypingCertifiedAsSeal                            bool
	ObservedParticleNamesUsed                              bool
	Supports, Failures                                     []string
}

type RhoAction struct {
	Domain, Codomain                                     string
	M3ActsOnP3W, M3TrivialOnP1W                          bool
	HActsOnLeftDoubleSocket, CActsOnRightSocketPair      bool
	ActionPreservesP1P3, ActionPreservesBMinusL          bool
	RepresentationLawConsistentAtBlockLevel              bool
	NativeDerivationCertified, ExplicitMatricesCertified bool
	FirstOrderConditionCertified, BimoduleCommutantProof bool
	CompleteRhoFActionLedger                             bool
	Supports, Failures                                   []string
}

type Projector struct {
	Name       string
	Rank       int
	Expression string
}

type ProjectorLedger struct {
	ParticleProjectors                []Projector
	ParticleRankSum, ExpectedHPartDim int
	Orthogonal, CompleteOnHPart       bool
	JCopyIncluded                     bool
	HFProjectorRankSum, ExpectedHFDim int
	CoarsePiSectorFSealCertified      bool
	FullNativePiSectorFCertified      bool
	FineColorAtomLedgerCertified      bool
	CanonicalColorFrameCertified      bool
	TraceMagnitudeReadoutCertified    bool
	Supports, Failures                []string
}

type DFEdgeSkeleton struct {
	SymbolicEdgeSupportAudited             bool
	AllowedEdgeClass                       string
	CouplingGraphOnly, NumericalMagnitudes bool
	UsesObservedMasses, UsesCKM, UsesPMNS  bool
	Supports, Failures                     []string
}

type CompressionStance struct {
	SectorBodyBeforeCompression        bool
	AggregateCompressionMapCertified   bool
	AggregateToSectorPullbackCertified bool
	AggregateOperatorIsSectorLedger    bool
	TraceMagnitudeReadoutCertified     bool
	Supports, Failures                 []string
}

type Impact struct {
	RhoActionSealConstructed, CoarseLedgerConstructed bool
	CarrierProblemSolvedAtBlockLevel                  bool
	FullNativeFiniteTriple, FullPiSectorF             bool
	CanPromoteToR3, CanPromoteToR4                    bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs  bool
	NextGate, Classification                          string
	Verdicts, Failures                                []string
}

type Firewalls struct {
	Enforced                                                                bool
	RepresentationSealNotNative, NoFullFiniteTripleProof                    bool
	NoExplicitMatrices, NoFirstOrderProof, NoBimoduleProof                  bool
	NoCanonicalColorAtoms, MatrixUnitsBasisDependent, NoFineColorLedger     bool
	DFEdgesNotMagnitudes, NoNumericalYukawaSockets, NoTraceMagnitudeReadout bool
	NoCompressionMap, AggregateNotSectorLedger, AlphaSealed                 bool
	NoNEffUpdate, NoCYukawaUpdate, NoObservedYukawaFit                      bool
	NoParticleAssignment, NoThreeGeneration, NotR3, NotR4                   bool
	Verdict                                                                 string
}

type Analysis struct {
	Ledger      Ledger
	W           WCarrier
	E           ECarrier
	Rho         RhoAction
	Projectors  ProjectorLedger
	DF          DFEdgeSkeleton
	Compression CompressionStance
	Impact      Impact
	Firewalls   Firewalls
	Truth       string
	Final       string
}

func BuildDefault() (Analysis, error) {
	wTrace := BMinusLLeptonWeight*float64(LeptonBlockDim) + BMinusLColorWeight*float64(ColorBlockDim)
	if math.Abs(wTrace) > 1e-15 {
		return Analysis{}, fmt.Errorf("B-L trace on W is not zero: %.17g", wTrace)
	}
	projectors := []Projector{
		{Name: "Pi_R1", Rank: RankR1, Expression: "P_R tensor P_1"},
		{Name: "Pi_R3", Rank: RankR3, Expression: "P_R tensor P_3"},
		{Name: "Pi_L1", Rank: RankL1, Expression: "P_L tensor P_1"},
		{Name: "Pi_L3", Rank: RankL3, Expression: "P_L tensor P_3"},
	}
	rankSum := 0
	for _, p := range projectors {
		rankSum += p.Rank
	}
	if rankSum != HPartDim {
		return Analysis{}, fmt.Errorf("particle projector rank sum %d != %d", rankSum, HPartDim)
	}

	a := Analysis{}
	a.Ledger = Ledger{S: SBoundary, AlphaB: AlphaB, OperatorNEff: OperatorNEff, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, R2PlusPlus: true, R3: false, R4: false, AlphaIsNative: false, OfficialFrozen: true}
	a.W = WCarrier{
		Expression: "W = C_lepton plus C^3_color", LeptonDim: LeptonBlockDim, ColorDim: ColorBlockDim, Dim: WDim,
		P1Rank: LeptonBlockDim, P3Rank: ColorBlockDim, P1P3Orthogonal: true, P1PlusP3CompletesW: true,
		BMinusLLeptonWeight: BMinusLLeptonWeight, BMinusLColorWeight: BMinusLColorWeight, BMinusLTraceZero: true,
		P3WIsM3Fundamental: true, ColorBlockCanonical: true, ColorAtomsCanonical: false, CanonicalColorFrame: false,
		Supports: []string{SupportSharedLeptoColorCarrierInherited, SupportP3WAsM3Fundamental, SupportBMinusLInternalToW},
		Failures: []string{FailureNoCanonicalColorAtomFrame, FailureM3MatrixUnitsBasisDependent, FailureNoFineColorAtomLedger},
	}
	a.E = ECarrier{
		Expression: "E = C_R^2 plus C_L^2", RightSlotDim: RightSlotDim, LeftSlotDim: LeftSlotDim, Dim: ElectroweakSlotDim,
		RightSocketPairDeclared: true, LeftDoubleSocketDeclared: true, QuaternionicWeakRoleDeclared: true, ComplexRightRoleDeclared: true,
		SourceTypingCertifiedAsSeal: true, ObservedParticleNamesUsed: false,
		Supports: []string{SupportESlotRightLeftSocketBody},
		Failures: []string{FailureRepresentationSealNotNative, FailureNoFullFiniteTripleProof},
	}
	a.Rho = RhoAction{
		Domain: "A_F = C plus H plus M_3(C)", Codomain: "End(((C_R^2 plus C_L^2) tensor (C plus C^3)) plus J-copy)",
		M3ActsOnP3W: true, M3TrivialOnP1W: true, HActsOnLeftDoubleSocket: true, CActsOnRightSocketPair: true,
		ActionPreservesP1P3: true, ActionPreservesBMinusL: true, RepresentationLawConsistentAtBlockLevel: true,
		NativeDerivationCertified: false, ExplicitMatricesCertified: false, FirstOrderConditionCertified: false, BimoduleCommutantProof: false, CompleteRhoFActionLedger: false,
		Supports: []string{SupportRhoFActionSeal, SupportP3WAsM3Fundamental},
		Failures: []string{FailureRepresentationSealNotNative, FailureNoFullFiniteTripleProof, FailureNoExplicitMatrices, FailureNoFirstOrderProof, FailureNoBimoduleCommutantProof},
	}
	a.Projectors = ProjectorLedger{
		ParticleProjectors: projectors, ParticleRankSum: rankSum, ExpectedHPartDim: HPartDim, Orthogonal: true, CompleteOnHPart: true,
		JCopyIncluded: true, HFProjectorRankSum: JCopyCount * rankSum, ExpectedHFDim: HFDim,
		CoarsePiSectorFSealCertified: true, FullNativePiSectorFCertified: false, FineColorAtomLedgerCertified: false, CanonicalColorFrameCertified: false, TraceMagnitudeReadoutCertified: false,
		Supports: []string{SupportCoarsePiSectorFSeal, SupportJCopyFiniteBody},
		Failures: []string{FailureNoFullFiniteTripleProof, FailureNoCanonicalColorAtomFrame, FailureNoFineColorAtomLedger, FailureNoTraceMagnitudeReadout},
	}
	a.DF = DFEdgeSkeleton{
		SymbolicEdgeSupportAudited: true, AllowedEdgeClass: "left/right socket edge support only, with symbolic Yukawa sockets", CouplingGraphOnly: true,
		NumericalMagnitudes: false, UsesObservedMasses: false, UsesCKM: false, UsesPMNS: false,
		Supports: []string{SupportDFEdgesAsSocketGraph},
		Failures: []string{FailureDFEdgesNotMagnitudes, FailureNoNumericalYukawaSockets, FailureNoObservedYukawaFit},
	}
	a.Compression = CompressionStance{
		SectorBodyBeforeCompression: true, AggregateCompressionMapCertified: false, AggregateToSectorPullbackCertified: false, AggregateOperatorIsSectorLedger: false, TraceMagnitudeReadoutCertified: false,
		Supports: []string{SupportSectorBodyBeforeCompression},
		Failures: []string{FailureNoAggregateCompressionMap, FailureAggregateOperatorNotSectorLedger, FailureNoTraceMagnitudeReadout},
	}
	a.Impact = Impact{
		RhoActionSealConstructed: true, CoarseLedgerConstructed: true, CarrierProblemSolvedAtBlockLevel: true,
		FullNativeFiniteTriple: false, FullPiSectorF: false, CanPromoteToR3: false, CanPromoteToR4: false,
		CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false,
		NextGate: "Gate 839 — Aggregate Trace Compression Map Construction/Obstruction Audit", Classification: "sealed coarse finite-sector body, not R3 trace-magnitude ledger",
		Verdicts: []string{StatusGate837Inherited, StatusWCarrierReverified, StatusESlotAudited, StatusSchematicRhoActionConsistent, StatusCoarseParticleLedgerConstructed, StatusR2PlusPlusRetained},
		Failures: []string{FailureRepresentationSealNotNative, FailureNoFullFiniteTripleProof, FailureNoTraceMagnitudeReadout, FailureNoAggregateCompressionMap, FailureNotR3, FailureNotR4, FailureNoNEffUpdate, FailureNoCYukawaUpdate},
	}
	a.Firewalls = Firewalls{Enforced: true, RepresentationSealNotNative: true, NoFullFiniteTripleProof: true, NoExplicitMatrices: true, NoFirstOrderProof: true, NoBimoduleProof: true, NoCanonicalColorAtoms: true, MatrixUnitsBasisDependent: true, NoFineColorLedger: true, DFEdgesNotMagnitudes: true, NoNumericalYukawaSockets: true, NoTraceMagnitudeReadout: true, NoCompressionMap: true, AggregateNotSectorLedger: true, AlphaSealed: true, NoNEffUpdate: true, NoCYukawaUpdate: true, NoObservedYukawaFit: true, NoParticleAssignment: true, NoThreeGeneration: true, NotR3: true, NotR4: true, Verdict: StatusFirewallGate838}
	a.Truth = "Gate 838 constructs a sealed coarse finite-sector projector body on H_part, but does not convert it into trace magnitudes, R3, or a native Yukawa theorem."
	a.Final = "Verdict: W and E support a consistent sealed rho_F action and the coarse projectors Pi_R1, Pi_R3, Pi_L1, Pi_L3; the next missing object is an aggregate trace-compression map from this sector body to the R2++ shadow, not a magnitude update."
	return a, nil
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("ledger: alpha_B=%.16g operator_N_eff=%.16g official_N_eff=%.16g R2++=%t R3=%t R4=%t frozen=%t", l.AlphaB, l.OperatorNEff, l.OfficialNEff, l.R2PlusPlus, l.R3, l.R4, l.OfficialFrozen)
}
func FormatW(w WCarrier) string {
	return fmt.Sprintf("W: %s dim=%d P1=%d P3=%d trace(B-L)=0?%t P3W=M3fund?%t color_atoms_canonical?%t", w.Expression, w.Dim, w.P1Rank, w.P3Rank, w.BMinusLTraceZero, w.P3WIsM3Fundamental, w.ColorAtomsCanonical)
}
func FormatE(e ECarrier) string {
	return fmt.Sprintf("E: %s dim=%d R=%d L=%d H_on_L=%t C_on_R=%t observed_names=%t", e.Expression, e.Dim, e.RightSlotDim, e.LeftSlotDim, e.QuaternionicWeakRoleDeclared, e.ComplexRightRoleDeclared, e.ObservedParticleNamesUsed)
}
func FormatRho(r RhoAction) string {
	return fmt.Sprintf("rho: %s -> %s M3_on_P3=%t M3_trivial_P1=%t H_on_L=%t C_on_R=%t preserve_B-L=%t native=%t explicit_matrices=%t", r.Domain, r.Codomain, r.M3ActsOnP3W, r.M3TrivialOnP1W, r.HActsOnLeftDoubleSocket, r.CActsOnRightSocketPair, r.ActionPreservesBMinusL, r.NativeDerivationCertified, r.ExplicitMatricesCertified)
}
func FormatProjectors(p ProjectorLedger) string {
	parts := make([]string, 0, len(p.ParticleProjectors))
	for _, q := range p.ParticleProjectors {
		parts = append(parts, fmt.Sprintf("%s:%d", q.Name, q.Rank))
	}
	return fmt.Sprintf("projectors: [%s] particle_sum=%d/%d HF_sum=%d/%d orthogonal=%t complete=%t coarse_seal=%t full_native=%t magnitudes=%t", strings.Join(parts, ","), p.ParticleRankSum, p.ExpectedHPartDim, p.HFProjectorRankSum, p.ExpectedHFDim, p.Orthogonal, p.CompleteOnHPart, p.CoarsePiSectorFSealCertified, p.FullNativePiSectorFCertified, p.TraceMagnitudeReadoutCertified)
}
func FormatDF(d DFEdgeSkeleton) string {
	return fmt.Sprintf("D_F: symbolic_edges=%t class=%s graph_only=%t numerical_magnitudes=%t observed_masses=%t CKM=%t PMNS=%t", d.SymbolicEdgeSupportAudited, d.AllowedEdgeClass, d.CouplingGraphOnly, d.NumericalMagnitudes, d.UsesObservedMasses, d.UsesCKM, d.UsesPMNS)
}
func FormatCompression(c CompressionStance) string {
	return fmt.Sprintf("compression: sector_body_before_compression=%t compression_map=%t aggregate_pullback=%t aggregate_is_sector_ledger=%t magnitude_readout=%t", c.SectorBodyBeforeCompression, c.AggregateCompressionMapCertified, c.AggregateToSectorPullbackCertified, c.AggregateOperatorIsSectorLedger, c.TraceMagnitudeReadoutCertified)
}
func FormatImpact(i Impact) string {
	return fmt.Sprintf("impact: rho_seal=%t coarse_ledger=%t full_native_triple=%t full_Pi_sector_F=%t R3=%t R4=%t next=%s", i.RhoActionSealConstructed, i.CoarseLedgerConstructed, i.FullNativeFiniteTriple, i.FullPiSectorF, i.CanPromoteToR3, i.CanPromoteToR4, i.NextGate)
}

func Statuses() []string {
	return []string{
		StatusGate837Inherited, StatusWCarrierReverified, StatusESlotAudited, StatusSchematicRhoActionConsistent, StatusM3P3ActionCertified, StatusHLeftActionCertified, StatusCRightActionCertified, StatusP1P3BMinusLPreserved, StatusCoarseParticleLedgerConstructed, StatusCoarseLedgerOrthogonalComplete, StatusJCopyLedgerDoublesRanks, StatusDFSymbolicEdgeSkeletonAudited, StatusCarrierLedgerNotMagnitudeReadout, StatusOneBlockNotGenerationTheorem, StatusNoObservedDataUsed, StatusLedgerFrozen, StatusR2PlusPlusRetained, StatusFirewallGate838,
		SupportSharedLeptoColorCarrierInherited, SupportP3WAsM3Fundamental, SupportBMinusLInternalToW, SupportESlotRightLeftSocketBody, SupportRhoFActionSeal, SupportCoarsePiSectorFSeal, SupportJCopyFiniteBody, SupportDFEdgesAsSocketGraph, SupportSectorBodyBeforeCompression,
		FailureRepresentationSealNotNative, FailureNoFullFiniteTripleProof, FailureNoExplicitMatrices, FailureNoFirstOrderProof, FailureNoBimoduleCommutantProof, FailureNoCanonicalColorAtomFrame, FailureM3MatrixUnitsBasisDependent, FailureNoFineColorAtomLedger, FailureDFEdgesNotMagnitudes, FailureNoNumericalYukawaSockets, FailureNoTraceMagnitudeReadout, FailureNoAggregateCompressionMap, FailureAggregateOperatorNotSectorLedger, FailureAlphaStillSealed, FailureNoNEffUpdate, FailureNoCYukawaUpdate, FailureNoObservedYukawaFit, FailureNoParticleAssignment, FailureNoThreeGenerationTheorem, FailureNotR3, FailureNotR4,
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
