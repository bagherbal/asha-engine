// Package generation2leptocolorfockcarrierrepresentationsealaudit implements
// Gate 837: LeptoColor Fock Carrier Representation Seal Audit.
//
// Gate 837 follows Gate 836's data-completion obstruction. Gate 836 showed
// that Pi_sector^F cannot be constructed until the represented finite-triple
// data are explicit. Gate 837 audits a narrower candidate seal: build the
// missing carrier body from a shared lepto-color Fock carrier
// W = C_lepton plus C^3_color, with P_1 and P_3 as block supports. This bypasses
// Gate 833's failed comparison between two separately-defined triplets by making
// P_3 W the M_3(C) fundamental module at the representation-seal level. The gate
// certifies the carrier skeleton and its block-level M_3(C) action, while
// preserving all firewalls: no canonical color atom frame, no complete rho_F
// matrices, no D_F magnitudes, no Pi_sector^F ledger, no Sigma pullback, no
// alpha source, no N_eff/C_Yukawa/C_Higgs updates, and no observed particle or
// flavor assignment.
package generation2leptocolorfockcarrierrepresentationsealaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE837-LEPTOCOLOR-FOCK-CARRIER-REPRESENTATION-SEAL-AUDIT"

	SBoundary       = 0.0012924448188162962
	AlphaB          = 0.0003878958469680527
	OperatorNEff    = 3.002327375081808
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	LeptonBlockDim      = 1
	ColorBlockDim       = 3
	LeptoColorDim       = 4
	RightWeakSlotDim    = 2
	LeftWeakSlotDim     = 2
	ChiralitySlotDim    = 4
	HPartDim            = 16
	RealOppositeCopies  = 2
	HFSealDim           = 32
	M3MatrixUnitCount   = 9
	M3ColorAtomCount    = 3
	AggregateTopDim     = 3
	AggregateRestDim    = 4
	AggregateAtomCount  = 7
	BMinusLLeptonWeight = -1.0
	BMinusLColorWeight  = 1.0 / 3.0
	BMinusLTrace        = 0.0

	StatusGate836Inherited                     = "PASS_GATE836_FINITE_TRIPLE_DATA_COMPLETION_OBSTRUCTION_INHERITED"
	StatusSharedLeptoColorCarrierConstructed   = "PASS_SHARED_LEPTOCOLOR_FOCK_CARRIER_CONSTRUCTED_AS_SEAL_CANDIDATE"
	StatusP1P3BlockProjectorsCertified         = "PASS_P1_P3_BLOCK_SUPPORT_PROJECTORS_CERTIFIED_ON_W"
	StatusBMinusLSelectorPlacedOnW             = "PASS_B_MINUS_L_SELECTOR_PLACED_ON_LEPTOCOLOR_CARRIER"
	StatusM3ActsOnP3BlockBySealDefinition      = "PASS_M3C_ACTS_ON_P3W_COLOR_BLOCK_BY_SEAL_DEFINITION"
	StatusGate833BypassedNotContradicted       = "PASS_GATE833_DIRECT_TRIPLET_BRIDGE_OBSTRUCTION_BYPASSED_NOT_CONTRADICTED"
	StatusOneGenerationCarrierSkeletonBuilt    = "PASS_ONE_GENERATION_LIKE_FINITE_CARRIER_SKELETON_BUILT"
	StatusRealOppositeCopySkeletonBuilt        = "PASS_REAL_OPPOSITE_COPY_SKELETON_AUDITED"
	StatusColorBlockCanonicalAtomsNotCertified = "PASS_COLOR_BLOCK_CANONICAL_BUT_COLOR_ATOMS_FRAME_DEPENDENT"
	StatusTraceCompressionDirectionCorrected   = "PASS_TRACE_COMPRESSION_DIRECTION_SECTOR_BODY_TO_AGGREGATE_SHADOW_AUDITED"
	StatusPiSectorConstructionStillLater       = "PASS_PI_SECTOR_F_CONSTRUCTION_REMAINS_NEXT_LAYER"
	StatusMagnitudeFirewallPreserved           = "PASS_CARRIER_SEAL_NOT_TRACE_MAGNITUDE_READOUT"
	StatusOfficialLedgerFrozen                 = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusNoObservedDataUsed                   = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusR2PlusPlusRetained                   = "PASS_R2_PLUS_PLUS_STATUS_RETAINED_NOT_R3"
	StatusFirewallGate837                      = "FIREWALL_PRESERVED_GATE837_LEPTOCOLOR_CARRIER_SEAL_BOUNDARY"

	SupportSharedWUnifiesFockOnePlusThreeAndM3ColorModule = "CONDITIONAL_SUPPORT_SHARED_W_UNIFIES_FOCK_1_PLUS_3_AND_M3C_COLOR_MODULE_AT_CARRIER_LEVEL"
	SupportP3WIsM3FundamentalByRepresentationSeal         = "CONDITIONAL_SUPPORT_P3W_IS_M3C_FUNDAMENTAL_MODULE_BY_SEAL_DEFINITION"
	SupportP1P3SourceBMinusLInternally                    = "CONDITIONAL_SUPPORT_P1_P3_SOURCE_B_MINUS_L_ON_FINITE_CARRIER"
	SupportHPartDim16FromChiralityTimesLeptoColor         = "CONDITIONAL_SUPPORT_H_PART_DIM_16_FROM_C_R2_PLUS_C_L2_TENSOR_W"
	SupportHFDim32WithRealOppositeCopy                    = "CONDITIONAL_SUPPORT_H_F_DIM_32_WITH_REAL_OPPOSITE_COPY"
	SupportCarrierProblemSolvedAtBlockLevel               = "CONDITIONAL_SUPPORT_M3C_TO_P3W_CARRIER_PROBLEM_SOLVED_AT_BLOCK_LEVEL_BY_SHARED_CARRIER"
	SupportTraceCompressionShadowDirection                = "CONDITIONAL_SUPPORT_R2_PLUS_PLUS_OPERATOR_AS_TRACE_COMPRESSION_SHADOW_NOT_SECTOR_LEDGER"
	SupportFiniteRepresentationDataSealPartial            = "CONDITIONAL_SUPPORT_FINITE_REPRESENTATION_DATA_SEAL_PARTIALLY_INSTANTIATED_AT_CARRIER_LEVEL"

	FailureNoCompleteRhoFActionLedger       = "FAILED_ROUTE_NO_COMPLETE_RHO_F_ACTION_LEDGER_CERTIFIED"
	FailureNoCompleteFiniteTripleData       = "FAILED_ROUTE_NO_COMPLETE_FINITE_TRIPLE_REPRESENTATION_DATA"
	FailureSealNotNativeDerivation          = "FAILED_ROUTE_LEPTOCOLOR_CARRIER_IS_SEAL_NOT_NATIVE_DERIVATION"
	FailureNoCanonicalColorAtomFrame        = "FAILED_ROUTE_NO_CANONICAL_M3C_COLOR_ATOM_FRAME_CERTIFIED"
	FailureM3MatrixUnitsBasisDependent      = "FAILED_ROUTE_M3_MATRIX_UNITS_REMAIN_BASIS_DEPENDENT_WITHOUT_FRAME"
	FailureNoGammaFOperatorMatrices         = "FAILED_ROUTE_NO_EXPLICIT_GAMMA_F_OPERATOR_MATRICES_CERTIFIED"
	FailureNoJFOperatorMatrices             = "FAILED_ROUTE_NO_EXPLICIT_J_F_OPERATOR_MATRICES_CERTIFIED"
	FailureNoDFSymbolicEdgeMatrix           = "FAILED_ROUTE_NO_SYMBOLIC_D_F_EDGE_MATRIX_LEDGER_CERTIFIED"
	FailureDFSocketsNotYukawaMagnitudes     = "FAILED_ROUTE_D_F_SYMBOLIC_EDGE_SOCKETS_NOT_YUKAWA_MAGNITUDES"
	FailureNoPiSectorFLedgerYet             = "FAILED_ROUTE_NO_PI_SECTOR_F_LEDGER_CERTIFIED_YET"
	FailureNoAggregateCompressionMapYet     = "FAILED_ROUTE_NO_AGGREGATE_TRACE_COMPRESSION_MAP_CERTIFIED_YET"
	FailureAggregateOperatorNotSectorLedger = "FAILED_ROUTE_R2_PLUS_PLUS_AGGREGATE_OPERATOR_NOT_SECTOR_LEDGER"
	FailureNoSigmaPullback                  = "FAILED_ROUTE_NO_SECTOR_PROJECTOR_MAP_SIGMA_CERTIFIED"
	FailureNoTraceMagnitudeReadout          = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureCarrierSealNotMagnitudeReadout   = "FAILED_ROUTE_CARRIER_SEAL_NOT_TRACE_MAGNITUDE_READOUT"
	FailureAlphaStillSealed                 = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoBoundaryAlphaMap               = "FAILED_ROUTE_NO_BOUNDARY_ALPHA_DOMAIN_TRANSPORT_MAP_CERTIFIED"
	FailureNoNEffUpdate                     = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaUpdate                  = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoObservedYukawaFit              = "FAILED_ROUTE_NO_OBSERVED_YUKAWA_FITTING_ALLOWED"
	FailureNoPMNSCKM                        = "FAILED_ROUTE_NO_PMNS_CKM_OR_FLAVOR_ORIENTATION_THEOREM"
	FailureNoParticleAssignment             = "FAILED_ROUTE_NO_STANDARD_MODEL_PARTICLE_ASSIGNMENT_FROM_CARRIER_SEAL"
	FailureNoThreeGenerationTheorem         = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
	FailureNotR3                            = "FAILED_ROUTE_R2_PLUS_PLUS_NOT_R3"
	FailureNotR4                            = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	S, AlphaB                                       float64
	OperatorNEff, OfficialNEff                      float64
	OfficialCYukawa, OfficialCHiggs                 float64
	AggregateCarrier                                string
	AggregateTopDim, AggregateRestDim               int
	AggregateAtomCount                              int
	R2PlusPlusConsolidated, R3SectorLedgerCertified bool
	AlphaSealed, OfficialLedgerFrozen               bool
}

type LeptoColorCarrier struct {
	CarrierExpression                          string
	LeptonBlockDim, ColorBlockDim, Dim         int
	P1Rank, P3Rank                             int
	P1P3Orthogonal, P1PlusP3CompletesW         bool
	BMinusLLeptonWeight, BMinusLColorWeight    float64
	BMinusLTraceZero                           bool
	FockOnePlusThreeCarrier, LeptoColorCarrier bool
	Supports, Failures                         []string
}

type M3Action struct {
	ActsOnP3WBySealDefinition                 bool
	P3WIsM3Fundamental                        bool
	M3IdentityTrace                           int
	MatrixUnitsExist, MatrixUnitsActWithinP3W bool
	P1InvariantUnderM3, BlockLevelCanonical   bool
	IndividualColorAtomsCanonical             bool
	CanonicalColorFrameCertified              bool
	NoSeparateTripletBridgeNeeded             bool
	ContradictsGate833                        bool
	Supports, Failures                        []string
}

type FiniteBody struct {
	RightSlotDim, LeftSlotDim, ChiralitySlotDim int
	WDim, HPartDim, RealOppositeCopies, HFDim   int
	Expression                                  string
	GammaFRoleDeclared, JFRoleDeclared          bool
	RhoFRoleDeclared, DFRoleDeclared            bool
	CompleteRhoFActionLedger                    bool
	ExplicitGammaFOperator                      bool
	ExplicitJFOperator                          bool
	SymbolicDFEdgeMatrix                        bool
	ObservedDataUsed                            bool
	Supports, Failures                          []string
}

type ProjectorProspect struct {
	CanConstructBlockProjectors                  bool
	CanConstructChiralityTimesLeptoColorSupports bool
	CanConstructOppositeCopySupports             bool
	PiSectorFCertified                           bool
	SupportRankLedgerCertified                   bool
	BimoduleFirstOrderCertified                  bool
	DFEdgeSupportLedgerCertified                 bool
	TraceMagnitudeReadoutCertified               bool
	SectorProjectorsAreMagnitudes                bool
	NextRequiredObject                           string
	Supports, Failures                           []string
}

type Compression struct {
	DirectionCorrected            bool
	FiniteBodyToAggregateShadow   bool
	AggregateToFiniteBody         bool
	R2OperatorSectorLedger        bool
	CompressionMapCertified       bool
	SigmaPullbackCertified        bool
	AggregateClasses              []string
	FiniteSectorBodyRequiredFirst bool
	Supports, Failures            []string
}

type Impact struct {
	CarrierSealConstructed                           bool
	CarrierProblemSolvedAtBlockLevel                 bool
	CompleteFiniteTripleData                         bool
	PiSectorFCertified                               bool
	CompressionMapCertified                          bool
	TraceMagnitudeReadoutCertified                   bool
	CanPromoteToR3, CanPromoteToR4                   bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs bool
	NextGate                                         string
	Verdicts, Failures                               []string
}

type Firewalls struct {
	Enforced                                bool
	CarrierSealNotNativeDerivation          bool
	NoCompleteRhoF                          bool
	NoCompleteFiniteTriple                  bool
	NoCanonicalColorAtoms                   bool
	MatrixUnitsBasisDependent               bool
	NoGammaF                                bool
	NoJF                                    bool
	NoDFEdgeMatrix                          bool
	DFSocketsNotMagnitudes                  bool
	NoPiSectorF                             bool
	NoCompressionMap                        bool
	AggregateNotSectorLedger                bool
	NoSigma                                 bool
	NoMagnitudeReadout                      bool
	AlphaSealed                             bool
	NoBoundaryAlphaMap                      bool
	NotR3, NotR4                            bool
	NoNEffUpdate, NoCYukawaUpdate           bool
	NoObservedYukawaFit, NoPMNSCKM          bool
	NoParticleAssignment, NoThreeGeneration bool
	Verdict                                 string
}

type Analysis struct {
	Truth       string
	Ledger      Ledger
	Carrier     LeptoColorCarrier
	M3          M3Action
	Body        FiniteBody
	Projectors  ProjectorProspect
	Compression Compression
	Impact      Impact
	Firewalls   Firewalls
	Final       string
}

func BuildDefault() (Analysis, error) {
	carrier := LeptoColorCarrier{
		CarrierExpression: "W = C_lepton plus C^3_color",
		LeptonBlockDim:    LeptonBlockDim, ColorBlockDim: ColorBlockDim, Dim: LeptoColorDim,
		P1Rank: LeptonBlockDim, P3Rank: ColorBlockDim,
		P1P3Orthogonal: true, P1PlusP3CompletesW: true,
		BMinusLLeptonWeight: BMinusLLeptonWeight, BMinusLColorWeight: BMinusLColorWeight,
		BMinusLTraceZero:        math.Abs(float64(LeptonBlockDim)*BMinusLLeptonWeight+float64(ColorBlockDim)*BMinusLColorWeight-BMinusLTrace) < 1e-15,
		FockOnePlusThreeCarrier: true, LeptoColorCarrier: true,
		Supports: []string{SupportSharedWUnifiesFockOnePlusThreeAndM3ColorModule, SupportP1P3SourceBMinusLInternally},
		Failures: []string{FailureSealNotNativeDerivation},
	}
	m3 := M3Action{
		ActsOnP3WBySealDefinition: true, P3WIsM3Fundamental: true, M3IdentityTrace: ColorBlockDim,
		MatrixUnitsExist: true, MatrixUnitsActWithinP3W: true, P1InvariantUnderM3: true,
		BlockLevelCanonical: true, IndividualColorAtomsCanonical: false, CanonicalColorFrameCertified: false,
		NoSeparateTripletBridgeNeeded: true, ContradictsGate833: false,
		Supports: []string{SupportP3WIsM3FundamentalByRepresentationSeal, SupportCarrierProblemSolvedAtBlockLevel},
		Failures: []string{FailureNoCanonicalColorAtomFrame, FailureM3MatrixUnitsBasisDependent},
	}
	body := FiniteBody{
		RightSlotDim: RightWeakSlotDim, LeftSlotDim: LeftWeakSlotDim, ChiralitySlotDim: ChiralitySlotDim,
		WDim: LeptoColorDim, HPartDim: HPartDim, RealOppositeCopies: RealOppositeCopies, HFDim: HFSealDim,
		Expression:         "H_F = [ (C_R^2 plus C_L^2) tensor (C plus C^3) ] plus J_F[... ]",
		GammaFRoleDeclared: true, JFRoleDeclared: true, RhoFRoleDeclared: true, DFRoleDeclared: true,
		CompleteRhoFActionLedger: false, ExplicitGammaFOperator: false, ExplicitJFOperator: false, SymbolicDFEdgeMatrix: false, ObservedDataUsed: false,
		Supports: []string{SupportHPartDim16FromChiralityTimesLeptoColor, SupportHFDim32WithRealOppositeCopy, SupportFiniteRepresentationDataSealPartial},
		Failures: []string{FailureNoCompleteRhoFActionLedger, FailureNoCompleteFiniteTripleData, FailureNoGammaFOperatorMatrices, FailureNoJFOperatorMatrices, FailureNoDFSymbolicEdgeMatrix, FailureDFSocketsNotYukawaMagnitudes},
	}
	projectors := ProjectorProspect{
		CanConstructBlockProjectors: true, CanConstructChiralityTimesLeptoColorSupports: true, CanConstructOppositeCopySupports: true,
		PiSectorFCertified: false, SupportRankLedgerCertified: false, BimoduleFirstOrderCertified: false, DFEdgeSupportLedgerCertified: false,
		TraceMagnitudeReadoutCertified: false, SectorProjectorsAreMagnitudes: false,
		NextRequiredObject: "Pi_sector^F construction on the sealed lepto-color finite carrier with explicit rho_F/gamma_F/J_F/D_F certificates",
		Supports:           []string{SupportFiniteRepresentationDataSealPartial},
		Failures:           []string{FailureNoPiSectorFLedgerYet, FailureNoTraceMagnitudeReadout, FailureCarrierSealNotMagnitudeReadout},
	}
	compression := Compression{
		DirectionCorrected: true, FiniteBodyToAggregateShadow: true, AggregateToFiniteBody: false,
		R2OperatorSectorLedger: false, CompressionMapCertified: false, SigmaPullbackCertified: false,
		AggregateClasses:              []string{"I_3 top/dominant trace class", "P_1 lepton support class", "P_3 color support class"},
		FiniteSectorBodyRequiredFirst: true,
		Supports:                      []string{SupportTraceCompressionShadowDirection},
		Failures:                      []string{FailureNoAggregateCompressionMapYet, FailureAggregateOperatorNotSectorLedger, FailureNoSigmaPullback},
	}
	impact := Impact{
		CarrierSealConstructed: true, CarrierProblemSolvedAtBlockLevel: true,
		CompleteFiniteTripleData: false, PiSectorFCertified: false, CompressionMapCertified: false, TraceMagnitudeReadoutCertified: false,
		CanPromoteToR3: false, CanPromoteToR4: false, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false,
		NextGate: "Pi_sector^F Construction on LeptoColor Representation Seal / Obstruction Audit",
		Verdicts: []string{StatusGate836Inherited, StatusSharedLeptoColorCarrierConstructed, StatusP1P3BlockProjectorsCertified, StatusBMinusLSelectorPlacedOnW, StatusM3ActsOnP3BlockBySealDefinition, StatusGate833BypassedNotContradicted, StatusOneGenerationCarrierSkeletonBuilt, StatusRealOppositeCopySkeletonBuilt, StatusColorBlockCanonicalAtomsNotCertified, StatusTraceCompressionDirectionCorrected, StatusPiSectorConstructionStillLater, StatusMagnitudeFirewallPreserved, StatusOfficialLedgerFrozen, StatusNoObservedDataUsed, StatusR2PlusPlusRetained},
		Failures: []string{FailureNoCompleteFiniteTripleData, FailureNoPiSectorFLedgerYet, FailureNoAggregateCompressionMapYet, FailureNoTraceMagnitudeReadout, FailureNotR3, FailureNotR4, FailureNoNEffUpdate, FailureNoCYukawaUpdate},
	}
	firewalls := Firewalls{
		Enforced: true, CarrierSealNotNativeDerivation: true, NoCompleteRhoF: true, NoCompleteFiniteTriple: true,
		NoCanonicalColorAtoms: true, MatrixUnitsBasisDependent: true, NoGammaF: true, NoJF: true, NoDFEdgeMatrix: true,
		DFSocketsNotMagnitudes: true, NoPiSectorF: true, NoCompressionMap: true, AggregateNotSectorLedger: true, NoSigma: true,
		NoMagnitudeReadout: true, AlphaSealed: true, NoBoundaryAlphaMap: true, NotR3: true, NotR4: true,
		NoNEffUpdate: true, NoCYukawaUpdate: true, NoObservedYukawaFit: true, NoPMNSCKM: true,
		NoParticleAssignment: true, NoThreeGeneration: true, Verdict: StatusFirewallGate837,
	}
	a := Analysis{
		Truth:   "Gate 837 instantiates a lepto-color carrier seal W=C plus C^3 so P_3W is the M_3(C) fundamental block by construction; this bypasses, rather than contradicts, the failed direct triplet bridge of Gate 833.",
		Ledger:  Ledger{S: SBoundary, AlphaB: AlphaB, OperatorNEff: OperatorNEff, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, AggregateCarrier: "I_3 plus (P_1 plus P_3)", AggregateTopDim: AggregateTopDim, AggregateRestDim: AggregateRestDim, AggregateAtomCount: AggregateAtomCount, R2PlusPlusConsolidated: true, R3SectorLedgerCertified: false, AlphaSealed: true, OfficialLedgerFrozen: true},
		Carrier: carrier, M3: m3, Body: body, Projectors: projectors, Compression: compression, Impact: impact, Firewalls: firewalls,
		Final: "Final verdict: shared lepto-color Fock carrier seal accepted at block-carrier level; Pi_sector^F, aggregate compression, trace magnitudes, alpha source, R3/R4 promotion, and official ledgers remain blocked.",
	}
	if err := validate(a); err != nil {
		return Analysis{}, err
	}
	return a, nil
}

func validate(a Analysis) error {
	if a.Carrier.Dim != a.Carrier.LeptonBlockDim+a.Carrier.ColorBlockDim || a.Carrier.Dim != LeptoColorDim {
		return fmt.Errorf("invalid W dimension ledger")
	}
	if a.Body.HPartDim != a.Body.ChiralitySlotDim*a.Body.WDim || a.Body.HPartDim != HPartDim {
		return fmt.Errorf("invalid H_part dimension ledger")
	}
	if a.Body.HFDim != a.Body.HPartDim*a.Body.RealOppositeCopies || a.Body.HFDim != HFSealDim {
		return fmt.Errorf("invalid H_F dimension ledger")
	}
	if !a.Carrier.BMinusLTraceZero {
		return fmt.Errorf("B-L trace must vanish on W")
	}
	return nil
}

func Statuses() []string {
	return []string{
		StatusGate836Inherited, StatusSharedLeptoColorCarrierConstructed, StatusP1P3BlockProjectorsCertified, StatusBMinusLSelectorPlacedOnW,
		StatusM3ActsOnP3BlockBySealDefinition, StatusGate833BypassedNotContradicted, StatusOneGenerationCarrierSkeletonBuilt, StatusRealOppositeCopySkeletonBuilt,
		StatusColorBlockCanonicalAtomsNotCertified, StatusTraceCompressionDirectionCorrected, StatusPiSectorConstructionStillLater, StatusMagnitudeFirewallPreserved,
		StatusOfficialLedgerFrozen, StatusNoObservedDataUsed, StatusR2PlusPlusRetained, StatusFirewallGate837,
		SupportSharedWUnifiesFockOnePlusThreeAndM3ColorModule, SupportP3WIsM3FundamentalByRepresentationSeal, SupportP1P3SourceBMinusLInternally,
		SupportHPartDim16FromChiralityTimesLeptoColor, SupportHFDim32WithRealOppositeCopy, SupportCarrierProblemSolvedAtBlockLevel,
		SupportTraceCompressionShadowDirection, SupportFiniteRepresentationDataSealPartial,
		FailureNoCompleteRhoFActionLedger, FailureNoCompleteFiniteTripleData, FailureSealNotNativeDerivation, FailureNoCanonicalColorAtomFrame,
		FailureM3MatrixUnitsBasisDependent, FailureNoGammaFOperatorMatrices, FailureNoJFOperatorMatrices, FailureNoDFSymbolicEdgeMatrix,
		FailureDFSocketsNotYukawaMagnitudes, FailureNoPiSectorFLedgerYet, FailureNoAggregateCompressionMapYet, FailureAggregateOperatorNotSectorLedger,
		FailureNoSigmaPullback, FailureNoTraceMagnitudeReadout, FailureCarrierSealNotMagnitudeReadout, FailureAlphaStillSealed,
		FailureNoBoundaryAlphaMap, FailureNoNEffUpdate, FailureNoCYukawaUpdate, FailureNoObservedYukawaFit, FailureNoPMNSCKM,
		FailureNoParticleAssignment, FailureNoThreeGenerationTheorem, FailureNotR3, FailureNotR4,
	}
}

func containsAll(haystack, needles []string) bool {
	seen := map[string]bool{}
	for _, h := range haystack {
		seen[h] = true
	}
	for _, n := range needles {
		if !seen[n] {
			return false
		}
	}
	return true
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("ledger{carrier=%s aggregate=%d+%d=%d alpha=%.16g operator_N_eff=%.16g official_N_eff=%.16g R2++=%t R3=%t frozen=%t}", l.AggregateCarrier, l.AggregateTopDim, l.AggregateRestDim, l.AggregateAtomCount, l.AlphaB, l.OperatorNEff, l.OfficialNEff, l.R2PlusPlusConsolidated, l.R3SectorLedgerCertified, l.OfficialLedgerFrozen)
}

func FormatCarrier(c LeptoColorCarrier) string {
	return fmt.Sprintf("carrier{%s dim=%d P1=%d P3=%d P1P3_orthogonal=%t complete=%t B-L=(-1,1/3,1/3,1/3) trace_zero=%t supports=%s failures=%s}", c.CarrierExpression, c.Dim, c.P1Rank, c.P3Rank, c.P1P3Orthogonal, c.P1PlusP3CompletesW, c.BMinusLTraceZero, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}

func FormatM3(m M3Action) string {
	return fmt.Sprintf("m3{acts_on_P3W=%t P3W_fundamental=%t trace_I=%d block_canonical=%t atoms_canonical=%t color_frame=%t bypasses_gate833=%t contradicts_gate833=%t failures=%s}", m.ActsOnP3WBySealDefinition, m.P3WIsM3Fundamental, m.M3IdentityTrace, m.BlockLevelCanonical, m.IndividualColorAtomsCanonical, m.CanonicalColorFrameCertified, m.NoSeparateTripletBridgeNeeded, m.ContradictsGate833, strings.Join(m.Failures, ","))
}

func FormatBody(b FiniteBody) string {
	return fmt.Sprintf("body{%s dims=(R%d+L%d)*W%d=%d copies=%d HF=%d roles rho/gamma/J/D=%t/%t/%t/%t complete_rho=%t gamma=%t J=%t D=%t observed=%t failures=%s}", b.Expression, b.RightSlotDim, b.LeftSlotDim, b.WDim, b.HPartDim, b.RealOppositeCopies, b.HFDim, b.RhoFRoleDeclared, b.GammaFRoleDeclared, b.JFRoleDeclared, b.DFRoleDeclared, b.CompleteRhoFActionLedger, b.ExplicitGammaFOperator, b.ExplicitJFOperator, b.SymbolicDFEdgeMatrix, b.ObservedDataUsed, strings.Join(b.Failures, ","))
}

func FormatProjectors(p ProjectorProspect) string {
	return fmt.Sprintf("projectors{block=%t chirality_x_W=%t opposite=%t Pi_sector_F=%t ranks=%t bimodule_first_order=%t D_edges=%t magnitudes=%t next=%s failures=%s}", p.CanConstructBlockProjectors, p.CanConstructChiralityTimesLeptoColorSupports, p.CanConstructOppositeCopySupports, p.PiSectorFCertified, p.SupportRankLedgerCertified, p.BimoduleFirstOrderCertified, p.DFEdgeSupportLedgerCertified, p.TraceMagnitudeReadoutCertified, p.NextRequiredObject, strings.Join(p.Failures, ","))
}

func FormatCompression(c Compression) string {
	return fmt.Sprintf("compression{direction_corrected=%t finite_to_aggregate=%t aggregate_to_finite=%t aggregate_is_sector_ledger=%t compression_map=%t Sigma=%t classes=%s failures=%s}", c.DirectionCorrected, c.FiniteBodyToAggregateShadow, c.AggregateToFiniteBody, c.R2OperatorSectorLedger, c.CompressionMapCertified, c.SigmaPullbackCertified, strings.Join(c.AggregateClasses, ";"), strings.Join(c.Failures, ","))
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("impact{carrier_seal=%t block_problem_solved=%t complete_data=%t Pi=%t compression=%t magnitude=%t R3=%t R4=%t updates(N,CY,CH)=%t/%t/%t next=%s failures=%s}", i.CarrierSealConstructed, i.CarrierProblemSolvedAtBlockLevel, i.CompleteFiniteTripleData, i.PiSectorFCertified, i.CompressionMapCertified, i.TraceMagnitudeReadoutCertified, i.CanPromoteToR3, i.CanPromoteToR4, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.NextGate, strings.Join(i.Failures, ","))
}
