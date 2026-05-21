// Package generation2rightleptocolorneutralsingletonpunctureedgesupportaudit implements
// Gate 842: Right LeptoColor Neutral Singleton Puncture / Edge-Support Audit.
//
// Gate 842 follows Gate 841's 8=7+1 punctured right lepto-color rectangle.
// It expands the right rectangle into its four cells, audits the excluded
// e_+ tensor P_1 cell as a neutral/right-lepton puncture candidate, and asks
// whether the current project data can certify a null-edge/sterile absence
// theorem that orients the dominant/rest socket compression. The answer is
// deliberately conservative: the four-cell anatomy and B-L compensation are
// certified; null-edge status, physical particle assignment, socket
// orientation, typed aggregate compression, alpha derivation, trace-magnitude
// readout, R3, and R4 remain blocked without an explicit D_F edge graph and
// full finite representation ledger.
package generation2rightleptocolorneutralsingletonpunctureedgesupportaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE842-RIGHT-LEPTOCOLOR-NEUTRAL-SINGLETON-PUNCTURE-EDGE-SUPPORT-AUDIT"

	SBoundary       = 0.0012924448188162962
	AlphaB          = 0.0003878958469680527
	OperatorNEff    = 3.002327375081808
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	LeptonBlockDim       = 1
	ColorBlockDim        = 3
	WDim                 = LeptonBlockDim + ColorBlockDim
	RightSocketPairDim   = 2
	RightRectangleRank   = RightSocketPairDim * WDim
	CharacterSocketRank  = 1
	EPlusP3Rank          = CharacterSocketRank * ColorBlockDim
	EPlusP1Rank          = CharacterSocketRank * LeptonBlockDim
	EMinusP3Rank         = CharacterSocketRank * ColorBlockDim
	EMinusP1Rank         = CharacterSocketRank * LeptonBlockDim
	ActiveSupportRank    = EPlusP3Rank + EMinusP3Rank + EMinusP1Rank
	PunctureRank         = EPlusP1Rank
	DominantTripletRank  = EPlusP3Rank
	RestQuartetRank      = EMinusP3Rank + EMinusP1Rank
	ParticleSideBodyRank = 16
	FiniteBodyRank       = 32

	BMinusLLeptonWeight = -1.0
	BMinusLColorWeight  = 1.0 / 3.0

	StatusGate841Inherited            = "PASS_GATE841_PUNCTURE_COMPLEMENT_INHERITED"
	StatusFourCellLedgerAudited       = "PASS_RIGHT_RECTANGLE_FOUR_CELL_LEDGER_AUDITED"
	StatusEightEqualsThreeThreeOneOne = "PASS_RIGHT_RECTANGLE_EIGHT_EQUALS_THREE_PLUS_THREE_PLUS_ONE_PLUS_ONE_CERTIFIED"
	StatusSevenEqualsThreeThreeOne    = "PASS_ACTIVE_SUPPORT_SEVEN_EQUALS_THREE_PLUS_THREE_PLUS_ONE_CERTIFIED"
	StatusPunctureSingletonAudited    = "PASS_NEUTRAL_RIGHT_LEPTON_SINGLETON_PUNCTURE_AUDITED"
	StatusCharacterOrientationAudited = "PASS_RIGHT_CHARACTER_ORIENTATION_ROUTE_AUDITED"
	StatusBMinusLPatternAudited       = "PASS_B_MINUS_L_PUNCTURE_COMPENSATION_PATTERN_AUDITED"
	StatusMinimalEdgeSupportAudited   = "PASS_MINIMAL_EDGE_SUPPORT_ROUTE_AUDITED"
	StatusAggregatePlacementAudited   = "PASS_ORIENTED_AGGREGATE_LOCATION_ROUTE_AUDITED"
	StatusMagnitudeFirewallPreserved  = "PASS_CARRIER_PUNCTURE_NOT_TRACE_MAGNITUDE_READOUT"
	StatusAlphaStillSealed            = "PASS_ALPHA_B_REMAINS_SEALED_AFTER_GATE842"
	StatusOfficialLedgersFrozen       = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusR2PlusPlusRetained          = "PASS_R2_PLUS_PLUS_RETAINED_NOT_R3_OR_R4"
	StatusNoObservedDataUsed          = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusFirewallGate842             = "FIREWALL_PRESERVED_GATE842_NEUTRAL_SINGLETON_PUNCTURE_OBSTRUCTION"

	SupportGate841ComplementInherited       = "CONDITIONAL_SUPPORT_GATE841_8_EQUALS_7_PLUS_1_COMPLEMENT_INHERITED"
	SupportRightRectangleFourCellLedger     = "CONDITIONAL_SUPPORT_RIGHT_RECTANGLE_HAS_FOUR_CELL_LEDGER_3_1_3_1"
	SupportActiveSupportRankSeven           = "CONDITIONAL_SUPPORT_ACTIVE_SUPPORT_IS_THREE_PLUS_THREE_PLUS_ONE_RANK_SEVEN"
	SupportPunctureSingletonRankOne         = "CONDITIONAL_SUPPORT_PUNCTURE_SINGLETON_IS_RANK_ONE"
	SupportPunctureIsRightLeptonColorless   = "CONDITIONAL_SUPPORT_PUNCTURE_IS_RIGHT_SOCKET_LEPTON_COLORLESS_CELL"
	SupportPunctureIsAbsentSterileCandidate = "CONDITIONAL_SUPPORT_PUNCTURE_IS_ABSENT_STERILE_SINGLETON_CANDIDATE_ONLY"
	SupportCharacterPairLambdaConjugate     = "CONDITIONAL_SUPPORT_RIGHT_SOCKET_PAIR_HAS_LAMBDA_CONJUGATE_CHARACTER_SEAL"
	SupportOrientationCandidate             = "CONDITIONAL_SUPPORT_E_PLUS_COLOR_AND_E_MINUS_QUARTET_ORIENTATION_CANDIDATE"
	SupportBMinusLActivePlusOne             = "CONDITIONAL_SUPPORT_ACTIVE_SUPPORT_HAS_B_MINUS_L_TRACE_PLUS_ONE"
	SupportBMinusLPunctureMinusOne          = "CONDITIONAL_SUPPORT_PUNCTURE_HAS_B_MINUS_L_TRACE_MINUS_ONE"
	SupportBMinusLFullRightRectangleNeutral = "CONDITIONAL_SUPPORT_FULL_RIGHT_RECTANGLE_B_MINUS_L_TRACE_ZERO"
	SupportDominantColorFiniteLocation      = "CONDITIONAL_SUPPORT_DOMINANT_I3_COULD_BE_LOCATED_ON_E_PLUS_TENSOR_P3_IF_ORIENTATION_CERTIFIED"
	SupportRestQuartetFiniteLocation        = "CONDITIONAL_SUPPORT_REST_W_COULD_BE_LOCATED_ON_E_MINUS_TENSOR_W_IF_ORIENTATION_CERTIFIED"
	SupportAggregateShadowIfEdgeAndOrient   = "CONDITIONAL_SUPPORT_AGGREGATE_SHADOW_LOCATION_IF_NULL_EDGE_AND_ORIENTATION_CERTIFIED"
	SupportPiActiveIsMinimalRightRectangle  = "CONDITIONAL_SUPPORT_ACTIVE_SUPPORT_IS_RIGHT_RECTANGLE_MINUS_NEUTRAL_SINGLETON"

	FailureRightCharacterSplitStillSeal      = "FAILED_ROUTE_RIGHT_SOCKET_CHARACTER_SPLIT_REMAINS_SEAL_NOT_NATIVE_DERIVATION"
	FailureNoExplicitRhoRMatrixProof         = "FAILED_ROUTE_NO_EXPLICIT_RHO_R_LAMBDA_BARLAMBDA_MATRIX_PROOF_CERTIFIED"
	FailureNoFullRhoFActionLedger            = "FAILED_ROUTE_NO_FULL_RHO_F_ACTION_LEDGER_CERTIFIED"
	FailureNoDFEdgeGraph                     = "FAILED_ROUTE_NO_D_F_EDGE_GRAPH_TO_CERTIFY_NULL_EDGE_PUNCTURE"
	FailureNoNullEdgeTheorem                 = "FAILED_ROUTE_NO_NULL_EDGE_THEOREM_FOR_NEUTRAL_SINGLETON"
	FailureNoMinimalAbsenceTheorem           = "FAILED_ROUTE_NO_MINIMAL_RIGHT_NEUTRAL_ABSENCE_THEOREM_CERTIFIED"
	FailureNoSterilePunctureTheorem          = "FAILED_ROUTE_NEUTRAL_SINGLETON_NOT_CERTIFIED_AS_STERILE_PUNCTURE"
	FailurePunctureNotPhysicalParticle       = "FAILED_ROUTE_NEUTRAL_SINGLETON_NOT_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoRightNeutrinoTheorem            = "FAILED_ROUTE_NO_RIGHT_NEUTRINO_OR_STERILE_PARTICLE_THEOREM"
	FailureNoDominantColorOrientationTheorem = "FAILED_ROUTE_NO_DOMINANT_COLOR_ORIENTATION_THEOREM"
	FailureNoRestQuartetOrientationTheorem   = "FAILED_ROUTE_NO_REST_QUARTET_ORIENTATION_THEOREM"
	FailureNoTypedSocketOrientationMap       = "FAILED_ROUTE_NO_TYPED_SOCKET_ORIENTATION_MAP_CERTIFIED"
	FailureNoTypedCompressionMap             = "FAILED_ROUTE_NO_TYPED_PUNCTURED_SOCKET_COMPRESSION_MAP_CERTIFIED"
	FailureCompressionCandidateNotTheorem    = "FAILED_ROUTE_ACTIVE_RIGHT_RECTANGLE_MINUS_PUNCTURE_IS_SUPPORT_ANATOMY_NOT_COMPRESSION_THEOREM"
	FailureNoAggregateCompressionMap         = "FAILED_ROUTE_NO_AGGREGATE_TRACE_COMPRESSION_MAP_CERTIFIED"
	FailureNoAlphaDerivation                 = "FAILED_ROUTE_PUNCTURE_EDGE_AUDIT_DOES_NOT_DERIVE_ALPHA_B"
	FailureAlphaStillSealed                  = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoTraceMagnitudeReadout           = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureCompressionNotYukawaMagnitude     = "FAILED_ROUTE_PUNCTURE_SUPPORT_NOT_YUKAWA_MAGNITUDE_SOURCE"
	FailureNoNEffUpdate                      = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaUpdate                   = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoObservedYukawaFit               = "FAILED_ROUTE_NO_OBSERVED_YUKAWA_FITTING_ALLOWED"
	FailureNoThreeGenerationTheorem          = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
	FailureNotR3                             = "FAILED_ROUTE_R2_PLUS_PLUS_NOT_R3"
	FailureNotR4                             = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	S, AlphaB                       float64
	OperatorNEff, OfficialNEff      float64
	OfficialCYukawa, OfficialCHiggs float64
	OfficialFrozen                  bool
	R2PlusPlus, R3, R4              bool
	AlphaIsNative                   bool
}

type Cell struct {
	Name, Expression, Socket, LeptoColorRole, SafeStructuralLabel string
	Rank                                                          int
	BMinusLTrace                                                  float64
	Active, Puncture, Colorless, Leptonic, Colored                bool
}

type FourCellLedger struct {
	Cells                                            []Cell
	FullRank, ActiveRank, PunctureRank               int
	RankPattern, ActivePattern                       string
	Orthogonal, Complete, Gate841Inherited           bool
	FourCellLedgerCertified, ActiveMinusPunctureForm bool
	Supports, Failures                               []string
}

type CharacterOrientation struct {
	CharacterModel                      string
	EPlusCharacter, EMinusCharacter     string
	UnorderedPairCertified              bool
	OrderedPhysicalOrientationCertified bool
	DominantColorOrientationCertified   bool
	RestQuartetOrientationCertified     bool
	EPlusColorWithEPlusLeptonPuncture   bool
	EMinusFullWQuartetCandidate         bool
	Supports, Failures                  []string
}

type BMinusLPattern struct {
	EPlusP3Trace, EPlusP1Trace, EMinusP3Trace, EMinusP1Trace float64
	ActiveTrace, PunctureTrace, FullTrace                    float64
	ActivePlusPunctureCancel, FullNeutral                    bool
	CompensatingSingletonPattern                             bool
	Supports, Failures                                       []string
}

type EdgeSupportAudit struct {
	PunctureExpression                                    string
	DFEdgeGraphAvailable                                  bool
	DFEdgesIntoPuncture, DFEdgesOutOfPuncture             int
	NullEdgeCertified, MinimalAbsenceCertified            bool
	SterilePunctureCertified, PhysicalAssignmentCertified bool
	SafeLabel                                             string
	Supports, Failures                                    []string
}

type AggregatePlacement struct {
	CandidateExpression, TopBlockExpression, RestBlockExpression string
	TopRank, RestRank, TotalRank                                 int
	FiniteBodyLocationCandidate                                  bool
	OrientedByNullEdgeCertified                                  bool
	CompressionMapCertified                                      bool
	TraceCompressionShadowCertified                              bool
	AlphaDerivedByCompression                                    bool
	TraceMagnitudeReadoutCertified                               bool
	R3, R4                                                       bool
	Supports, Failures                                           []string
}

type Firewall struct {
	Enforced                       bool
	RightCharacterSplitStillSeal   bool
	NoExplicitRhoRMatrixProof      bool
	NoFullRhoFActionLedger         bool
	NoDFEdgeGraph                  bool
	NoNullEdgeTheorem              bool
	NoMinimalAbsenceTheorem        bool
	NoSterilePunctureTheorem       bool
	PunctureNotPhysicalParticle    bool
	NoRightNeutrinoTheorem         bool
	NoDominantOrientationTheorem   bool
	NoRestOrientationTheorem       bool
	NoTypedSocketOrientationMap    bool
	NoTypedCompressionMap          bool
	CompressionCandidateNotTheorem bool
	NoAggregateCompressionMap      bool
	NoAlphaDerivation              bool
	AlphaSealed                    bool
	NoTraceMagnitudeReadout        bool
	CompressionNotYukawaMagnitude  bool
	NoNEffUpdate                   bool
	NoCYukawaUpdate                bool
	NoObservedYukawaFit            bool
	NoThreeGeneration              bool
	NotR3, NotR4                   bool
	Verdict                        string
}

type Impact struct {
	FourCellLedgerCertified                          bool
	NeutralSingletonPunctureIsolated                 bool
	BMinusLCompensationFound                         bool
	DFEdgeGraphStillMissing                          bool
	NullEdgeStillUncertified                         bool
	OrientationStillMissing                          bool
	CompressionMapStillMissing                       bool
	CanPromoteToR3, CanPromoteToR4                   bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs bool
	Classification                                   string
}

type Analysis struct {
	ID          string
	Truth       string
	Ledger      Ledger
	Cells       FourCellLedger
	Orientation CharacterOrientation
	BMinusL     BMinusLPattern
	Edge        EdgeSupportAudit
	Placement   AggregatePlacement
	Firewalls   Firewall
	Impact      Impact
	Final       string
}

func BuildDefault() (Analysis, error) {
	cells := buildFourCellLedger()
	orient := buildCharacterOrientation()
	bminusl := buildBMinusLPattern()
	edge := buildEdgeAudit()
	placement := buildAggregatePlacement()
	fw := buildFirewall()
	impact := buildImpact()
	ledger := Ledger{S: SBoundary, AlphaB: AlphaB, OperatorNEff: OperatorNEff, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true, R2PlusPlus: true, R3: false, R4: false, AlphaIsNative: false}
	if err := validate(cells, bminusl, placement); err != nil {
		return Analysis{}, err
	}
	return Analysis{
		ID:          AuditID,
		Truth:       "Gate 842 isolates the neutral/right-lepton singleton puncture e_+ tensor P_1 inside the four-cell right lepto-color rectangle, but refuses to certify null-edge absence, physical sterile assignment, socket orientation, aggregate compression, alpha source, trace magnitudes, R3, or R4 without explicit D_F and finite-representation data.",
		Ledger:      ledger,
		Cells:       cells,
		Orientation: orient,
		BMinusL:     bminusl,
		Edge:        edge,
		Placement:   placement,
		Firewalls:   fw,
		Impact:      impact,
		Final:       "Result: active rank-seven support is conditionally the right rectangle minus the neutral singleton puncture, but the puncture remains an absent/sterile candidate only; the aggregate operator does not yet become a certified finite-body trace compression theorem.",
	}, nil
}

func buildFourCellLedger() FourCellLedger {
	cells := []Cell{
		{Name: "e_plus_P3", Expression: "e_+ tensor P_3", Socket: "e_+", LeptoColorRole: "color triplet", SafeStructuralLabel: "dominant-character color triplet candidate", Rank: EPlusP3Rank, BMinusLTrace: float64(ColorBlockDim) * BMinusLColorWeight, Active: true, Colored: true},
		{Name: "e_plus_P1", Expression: "e_+ tensor P_1", Socket: "e_+", LeptoColorRole: "lepton singleton", SafeStructuralLabel: "neutral right-lepton puncture / absent sterile singleton candidate", Rank: EPlusP1Rank, BMinusLTrace: BMinusLLeptonWeight, Puncture: true, Colorless: true, Leptonic: true},
		{Name: "e_minus_P3", Expression: "e_- tensor P_3", Socket: "e_-", LeptoColorRole: "color triplet", SafeStructuralLabel: "rest-character color triplet candidate", Rank: EMinusP3Rank, BMinusLTrace: float64(ColorBlockDim) * BMinusLColorWeight, Active: true, Colored: true},
		{Name: "e_minus_P1", Expression: "e_- tensor P_1", Socket: "e_-", LeptoColorRole: "lepton singleton", SafeStructuralLabel: "rest-character lepton singleton candidate", Rank: EMinusP1Rank, BMinusLTrace: BMinusLLeptonWeight, Active: true, Colorless: true, Leptonic: true},
	}
	return FourCellLedger{
		Cells:                   cells,
		FullRank:                RightRectangleRank,
		ActiveRank:              ActiveSupportRank,
		PunctureRank:            PunctureRank,
		RankPattern:             "8 = 3 + 1 + 3 + 1",
		ActivePattern:           "7 = 3 + 3 + 1 = (e_+ tensor P_3) plus (e_- tensor P_3) plus (e_- tensor P_1)",
		Orthogonal:              true,
		Complete:                true,
		Gate841Inherited:        true,
		FourCellLedgerCertified: true,
		ActiveMinusPunctureForm: true,
		Supports:                []string{SupportGate841ComplementInherited, SupportRightRectangleFourCellLedger, SupportActiveSupportRankSeven, SupportPunctureSingletonRankOne, SupportPunctureIsRightLeptonColorless, SupportPunctureIsAbsentSterileCandidate, SupportPiActiveIsMinimalRightRectangle},
		Failures:                []string{FailureCompressionCandidateNotTheorem, FailureNoTypedCompressionMap, FailureNoAggregateCompressionMap},
	}
}

func buildCharacterOrientation() CharacterOrientation {
	return CharacterOrientation{
		CharacterModel:                      "rho_R(lambda)=diag(lambda,conjugate(lambda)) [sealed schematic inherited from Gate 840]",
		EPlusCharacter:                      "lambda-character socket candidate",
		EMinusCharacter:                     "conjugate-character socket candidate",
		UnorderedPairCertified:              true,
		OrderedPhysicalOrientationCertified: false,
		DominantColorOrientationCertified:   false,
		RestQuartetOrientationCertified:     false,
		EPlusColorWithEPlusLeptonPuncture:   true,
		EMinusFullWQuartetCandidate:         true,
		Supports:                            []string{SupportCharacterPairLambdaConjugate, SupportOrientationCandidate, SupportDominantColorFiniteLocation, SupportRestQuartetFiniteLocation},
		Failures:                            []string{FailureRightCharacterSplitStillSeal, FailureNoExplicitRhoRMatrixProof, FailureNoFullRhoFActionLedger, FailureNoDominantColorOrientationTheorem, FailureNoRestQuartetOrientationTheorem, FailureNoTypedSocketOrientationMap},
	}
}

func buildBMinusLPattern() BMinusLPattern {
	epp3 := float64(ColorBlockDim) * BMinusLColorWeight
	epp1 := BMinusLLeptonWeight
	emp3 := float64(ColorBlockDim) * BMinusLColorWeight
	emp1 := BMinusLLeptonWeight
	active := epp3 + emp3 + emp1
	puncture := epp1
	full := active + puncture
	return BMinusLPattern{
		EPlusP3Trace:                 epp3,
		EPlusP1Trace:                 epp1,
		EMinusP3Trace:                emp3,
		EMinusP1Trace:                emp1,
		ActiveTrace:                  active,
		PunctureTrace:                puncture,
		FullTrace:                    full,
		ActivePlusPunctureCancel:     nearlyEqual(full, 0),
		FullNeutral:                  nearlyEqual(full, 0),
		CompensatingSingletonPattern: nearlyEqual(active, 1) && nearlyEqual(puncture, -1),
		Supports:                     []string{SupportBMinusLActivePlusOne, SupportBMinusLPunctureMinusOne, SupportBMinusLFullRightRectangleNeutral},
		Failures:                     []string{},
	}
}

func buildEdgeAudit() EdgeSupportAudit {
	return EdgeSupportAudit{
		PunctureExpression:          "e_+ tensor P_1",
		DFEdgeGraphAvailable:        false,
		DFEdgesIntoPuncture:         0,
		DFEdgesOutOfPuncture:        0,
		NullEdgeCertified:           false,
		MinimalAbsenceCertified:     false,
		SterilePunctureCertified:    false,
		PhysicalAssignmentCertified: false,
		SafeLabel:                   "neutral right-lepton puncture / absent sterile singleton candidate only",
		Supports:                    []string{SupportPunctureIsAbsentSterileCandidate, SupportAggregateShadowIfEdgeAndOrient},
		Failures:                    []string{FailureNoDFEdgeGraph, FailureNoNullEdgeTheorem, FailureNoMinimalAbsenceTheorem, FailureNoSterilePunctureTheorem, FailurePunctureNotPhysicalParticle, FailureNoRightNeutrinoTheorem},
	}
}

func buildAggregatePlacement() AggregatePlacement {
	return AggregatePlacement{
		CandidateExpression:             "Pi_active = (e_+ tensor P_3) plus (e_- tensor W) = right rectangle minus (e_+ tensor P_1)",
		TopBlockExpression:              "I_{e_+ tensor P_3}",
		RestBlockExpression:             "[alpha_B P_3 - 3 alpha_B^2(B-L)] on e_- tensor W",
		TopRank:                         DominantTripletRank,
		RestRank:                        RestQuartetRank,
		TotalRank:                       ActiveSupportRank,
		FiniteBodyLocationCandidate:     true,
		OrientedByNullEdgeCertified:     false,
		CompressionMapCertified:         false,
		TraceCompressionShadowCertified: false,
		AlphaDerivedByCompression:       false,
		TraceMagnitudeReadoutCertified:  false,
		R3:                              false,
		R4:                              false,
		Supports:                        []string{SupportDominantColorFiniteLocation, SupportRestQuartetFiniteLocation, SupportAggregateShadowIfEdgeAndOrient},
		Failures:                        []string{FailureNoDFEdgeGraph, FailureNoNullEdgeTheorem, FailureNoDominantColorOrientationTheorem, FailureNoRestQuartetOrientationTheorem, FailureNoTypedCompressionMap, FailureNoAggregateCompressionMap, FailureNoAlphaDerivation, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureNotR3, FailureNotR4},
	}
}

func buildFirewall() Firewall {
	return Firewall{Enforced: true, RightCharacterSplitStillSeal: true, NoExplicitRhoRMatrixProof: true, NoFullRhoFActionLedger: true, NoDFEdgeGraph: true, NoNullEdgeTheorem: true, NoMinimalAbsenceTheorem: true, NoSterilePunctureTheorem: true, PunctureNotPhysicalParticle: true, NoRightNeutrinoTheorem: true, NoDominantOrientationTheorem: true, NoRestOrientationTheorem: true, NoTypedSocketOrientationMap: true, NoTypedCompressionMap: true, CompressionCandidateNotTheorem: true, NoAggregateCompressionMap: true, NoAlphaDerivation: true, AlphaSealed: true, NoTraceMagnitudeReadout: true, CompressionNotYukawaMagnitude: true, NoNEffUpdate: true, NoCYukawaUpdate: true, NoObservedYukawaFit: true, NoThreeGeneration: true, NotR3: true, NotR4: true, Verdict: StatusFirewallGate842}
}

func buildImpact() Impact {
	return Impact{FourCellLedgerCertified: true, NeutralSingletonPunctureIsolated: true, BMinusLCompensationFound: true, DFEdgeGraphStillMissing: true, NullEdgeStillUncertified: true, OrientationStillMissing: true, CompressionMapStillMissing: true, CanPromoteToR3: false, CanPromoteToR4: false, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, Classification: "right lepto-color 8=3+1+3+1 ledger with active 7=3+3+1 support; R2++ finite-body location candidate only"}
}

func validate(c FourCellLedger, b BMinusLPattern, p AggregatePlacement) error {
	if c.FullRank != RightRectangleRank || c.ActiveRank != 7 || c.PunctureRank != 1 {
		return fmt.Errorf("invalid cell ranks: full=%d active=%d puncture=%d", c.FullRank, c.ActiveRank, c.PunctureRank)
	}
	if len(c.Cells) != 4 {
		return fmt.Errorf("expected four cells, got %d", len(c.Cells))
	}
	var total, active, puncture int
	for _, cell := range c.Cells {
		total += cell.Rank
		if cell.Active {
			active += cell.Rank
		}
		if cell.Puncture {
			puncture += cell.Rank
		}
	}
	if total != c.FullRank || active != c.ActiveRank || puncture != c.PunctureRank {
		return fmt.Errorf("cell ledger totals mismatch: total=%d active=%d puncture=%d", total, active, puncture)
	}
	if !nearlyEqual(b.ActiveTrace, 1) || !nearlyEqual(b.PunctureTrace, -1) || !nearlyEqual(b.FullTrace, 0) {
		return fmt.Errorf("invalid B-L traces: active=%g puncture=%g full=%g", b.ActiveTrace, b.PunctureTrace, b.FullTrace)
	}
	if p.TopRank+p.RestRank != p.TotalRank || p.TotalRank != c.ActiveRank {
		return fmt.Errorf("placement rank mismatch: top=%d rest=%d total=%d", p.TopRank, p.RestRank, p.TotalRank)
	}
	return nil
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("s=%.16g alpha_B=%.16g operator_N_eff=%.16g official_N_eff=%.16g frozen=%t R2++=%t R3=%t R4=%t", l.S, l.AlphaB, l.OperatorNEff, l.OfficialNEff, l.OfficialFrozen, l.R2PlusPlus, l.R3, l.R4)
}

func FormatCells(c FourCellLedger) string {
	parts := make([]string, 0, len(c.Cells))
	for _, cell := range c.Cells {
		mark := "inactive"
		if cell.Active {
			mark = "active"
		}
		if cell.Puncture {
			mark = "puncture"
		}
		parts = append(parts, fmt.Sprintf("%s:%s rank=%d B-L=%.0f %s", cell.Name, cell.Expression, cell.Rank, cell.BMinusLTrace, mark))
	}
	return fmt.Sprintf("%s; active %s; full=%d active=%d puncture=%d cells=[%s]", c.RankPattern, c.ActivePattern, c.FullRank, c.ActiveRank, c.PunctureRank, strings.Join(parts, "; "))
}

func FormatOrientation(o CharacterOrientation) string {
	return fmt.Sprintf("%s; e_+=%s e_-=%s unordered=%t dominantCertified=%t restCertified=%t", o.CharacterModel, o.EPlusCharacter, o.EMinusCharacter, o.UnorderedPairCertified, o.DominantColorOrientationCertified, o.RestQuartetOrientationCertified)
}

func FormatBMinusL(b BMinusLPattern) string {
	return fmt.Sprintf("B-L traces e+P3=%.0f e+P1=%.0f e-P3=%.0f e-P1=%.0f active=%.0f puncture=%.0f full=%.0f neutral=%t", b.EPlusP3Trace, b.EPlusP1Trace, b.EMinusP3Trace, b.EMinusP1Trace, b.ActiveTrace, b.PunctureTrace, b.FullTrace, b.FullNeutral)
}

func FormatEdge(e EdgeSupportAudit) string {
	return fmt.Sprintf("puncture=%s label=%q D_F_edges_available=%t nullEdgeCertified=%t minimalAbsenceCertified=%t sterileCertified=%t physicalAssignment=%t", e.PunctureExpression, e.SafeLabel, e.DFEdgeGraphAvailable, e.NullEdgeCertified, e.MinimalAbsenceCertified, e.SterilePunctureCertified, e.PhysicalAssignmentCertified)
}

func FormatPlacement(p AggregatePlacement) string {
	return fmt.Sprintf("%s top=%s(rank %d) rest=%s(rank %d) total=%d locationCandidate=%t compressionCertified=%t traceCompressionCertified=%t R3=%t R4=%t", p.CandidateExpression, p.TopBlockExpression, p.TopRank, p.RestBlockExpression, p.RestRank, p.TotalRank, p.FiniteBodyLocationCandidate, p.CompressionMapCertified, p.TraceCompressionShadowCertified, p.R3, p.R4)
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("%s; fourCell=%t puncture=%t B-L=%t D_F_missing=%t nullEdge=%t orientationMissing=%t compressionMissing=%t R3=%t R4=%t", i.Classification, i.FourCellLedgerCertified, i.NeutralSingletonPunctureIsolated, i.BMinusLCompensationFound, i.DFEdgeGraphStillMissing, i.NullEdgeStillUncertified, i.OrientationStillMissing, i.CompressionMapStillMissing, i.CanPromoteToR3, i.CanPromoteToR4)
}

func Statuses() []string {
	return []string{
		StatusGate841Inherited, StatusFourCellLedgerAudited, StatusEightEqualsThreeThreeOneOne, StatusSevenEqualsThreeThreeOne, StatusPunctureSingletonAudited, StatusCharacterOrientationAudited, StatusBMinusLPatternAudited, StatusMinimalEdgeSupportAudited, StatusAggregatePlacementAudited, StatusMagnitudeFirewallPreserved, StatusAlphaStillSealed, StatusOfficialLedgersFrozen, StatusR2PlusPlusRetained, StatusNoObservedDataUsed, StatusFirewallGate842,
		SupportGate841ComplementInherited, SupportRightRectangleFourCellLedger, SupportActiveSupportRankSeven, SupportPunctureSingletonRankOne, SupportPunctureIsRightLeptonColorless, SupportPunctureIsAbsentSterileCandidate, SupportCharacterPairLambdaConjugate, SupportOrientationCandidate, SupportBMinusLActivePlusOne, SupportBMinusLPunctureMinusOne, SupportBMinusLFullRightRectangleNeutral, SupportDominantColorFiniteLocation, SupportRestQuartetFiniteLocation, SupportAggregateShadowIfEdgeAndOrient, SupportPiActiveIsMinimalRightRectangle,
		FailureRightCharacterSplitStillSeal, FailureNoExplicitRhoRMatrixProof, FailureNoFullRhoFActionLedger, FailureNoDFEdgeGraph, FailureNoNullEdgeTheorem, FailureNoMinimalAbsenceTheorem, FailureNoSterilePunctureTheorem, FailurePunctureNotPhysicalParticle, FailureNoRightNeutrinoTheorem, FailureNoDominantColorOrientationTheorem, FailureNoRestQuartetOrientationTheorem, FailureNoTypedSocketOrientationMap, FailureNoTypedCompressionMap, FailureCompressionCandidateNotTheorem, FailureNoAggregateCompressionMap, FailureNoAlphaDerivation, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureCompressionNotYukawaMagnitude, FailureNoNEffUpdate, FailureNoCYukawaUpdate, FailureNoObservedYukawaFit, FailureNoThreeGenerationTheorem, FailureNotR3, FailureNotR4,
	}
}

func containsAll(haystack, needles []string) bool {
	m := make(map[string]struct{}, len(haystack))
	for _, s := range haystack {
		m[s] = struct{}{}
	}
	for _, s := range needles {
		if _, ok := m[s]; !ok {
			return false
		}
	}
	return true
}

func containsText(values []string, needle string) bool {
	for _, v := range values {
		if strings.Contains(v, needle) {
			return true
		}
	}
	return false
}

func nearlyEqual(a, b float64) bool { return math.Abs(a-b) < 1e-12 }
