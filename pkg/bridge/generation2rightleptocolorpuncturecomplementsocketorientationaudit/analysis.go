// Package generation2rightleptocolorpuncturecomplementsocketorientationaudit implements
// Gate 841: Right LeptoColor Puncture Complement and Socket-Orientation Audit.
//
// Gate 841 follows Gate 840's right-socket character split. Gate 840 exposed
// the punctured right lepto-color rectangle
//
//	Pi_7 = (e_+ tensor P_3) plus (e_- tensor W)
//
// with excluded singleton e_+ tensor P_1. Gate 841 audits the sharper 8=7+1
// complement law, the B-L compensating trace pattern, and whether the excluded
// singleton can be promoted to a sterile/null-edge puncture that orients the
// dominant/rest sockets. The answer remains conservative: support anatomy and
// B-L compensation are certified; sterile/null-edge status, dominant/rest
// orientation, typed compression, alpha derivation, trace magnitudes, R3, and
// R4 remain blocked.
package generation2rightleptocolorpuncturecomplementsocketorientationaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE841-RIGHT-LEPTOCOLOR-PUNCTURE-COMPLEMENT-SOCKET-ORIENTATION-AUDIT"

	SBoundary       = 0.0012924448188162962
	AlphaB          = 0.0003878958469680527
	OperatorNEff    = 3.002327375081808
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	LeptonBlockDim = 1
	ColorBlockDim  = 3
	WDim           = LeptonBlockDim + ColorBlockDim
	RightSlotDim   = 2
	LeftSlotDim    = 2
	EDim           = RightSlotDim + LeftSlotDim
	HPartDim       = EDim * WDim
	HFDim          = 2 * HPartDim

	CharacterSocketRank = 1
	RightRectangleRank  = RightSlotDim * WDim
	DominantColorRank   = CharacterSocketRank * ColorBlockDim
	RestQuartetRank     = CharacterSocketRank * WDim
	ActiveSupportRank   = DominantColorRank + RestQuartetRank
	PunctureRank        = CharacterSocketRank * LeptonBlockDim

	BMinusLLeptonWeight = -1.0
	BMinusLColorWeight  = 1.0 / 3.0

	StatusGate840Inherited                  = "PASS_GATE840_PUNCTURED_RIGHT_RECTANGLE_INHERITED"
	StatusComplementLawAudited              = "PASS_RIGHT_LEPTOCOLOR_PUNCTURE_COMPLEMENT_LAW_AUDITED"
	StatusEightEqualsSevenPlusOne           = "PASS_RIGHT_RECTANGLE_EIGHT_EQUALS_SEVEN_PLUS_ONE_CERTIFIED"
	StatusPunctureRankAudited               = "PASS_EXCLUDED_SINGLETON_PUNCTURE_RANK_AUDITED"
	StatusBMinusLCompensationAudited        = "PASS_B_MINUS_L_COMPENSATING_PUNCTURE_BALANCE_AUDITED"
	StatusSterileNullEdgeAudited            = "PASS_STERILE_NULL_EDGE_PUNCTURE_ROUTE_AUDITED"
	StatusDominantOrientationAudited        = "PASS_DOMINANT_COLOR_SOCKET_ORIENTATION_ROUTE_AUDITED"
	StatusRestOrientationAudited            = "PASS_REST_QUARTET_SOCKET_ORIENTATION_ROUTE_AUDITED"
	StatusAggregateLocationCandidateAudited = "PASS_AGGREGATE_FINITE_BODY_LOCATION_CANDIDATE_AUDITED"
	StatusMagnitudeFirewallPreserved        = "PASS_PUNCTURE_ORIENTATION_NOT_TRACE_MAGNITUDE_READOUT"
	StatusAlphaStillSealed                  = "PASS_ALPHA_B_REMAINS_SEALED_AFTER_PUNCTURE_ORIENTATION_AUDIT"
	StatusOfficialLedgersFrozen             = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusR2PlusPlusRetained                = "PASS_R2_PLUS_PLUS_RETAINED_NOT_R3_OR_R4"
	StatusNoObservedDataUsed                = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusFirewallGate841                   = "FIREWALL_PRESERVED_GATE841_PUNCTURE_ORIENTATION_OBSTRUCTION"

	SupportGate840PunctureInherited              = "CONDITIONAL_SUPPORT_GATE840_PUNCTURED_RIGHT_RECTANGLE_INHERITED"
	SupportRightRectangleComplement              = "CONDITIONAL_SUPPORT_RIGHT_RECTANGLE_DECOMPOSES_AS_ACTIVE_SEVEN_PLUS_PUNCTURE_ONE"
	SupportActiveSupportRankSeven                = "CONDITIONAL_SUPPORT_ACTIVE_SUPPORT_HAS_RANK_SEVEN"
	SupportPunctureSingletonRankOne              = "CONDITIONAL_SUPPORT_PUNCTURE_SINGLETON_HAS_RANK_ONE"
	SupportBMinusLActivePlusOne                  = "CONDITIONAL_SUPPORT_ACTIVE_SUPPORT_HAS_B_MINUS_L_TRACE_PLUS_ONE"
	SupportBMinusLPunctureMinusOne               = "CONDITIONAL_SUPPORT_PUNCTURE_SINGLETON_HAS_B_MINUS_L_TRACE_MINUS_ONE"
	SupportBMinusLFullRightRectangleNeutral      = "CONDITIONAL_SUPPORT_FULL_RIGHT_RECTANGLE_B_MINUS_L_TRACE_ZERO"
	SupportPunctureIsCompensatingSingleton       = "CONDITIONAL_SUPPORT_PUNCTURE_IS_B_MINUS_L_COMPENSATING_SINGLETON"
	SupportSterileNullEdgeCandidate              = "CONDITIONAL_SUPPORT_EXCLUDED_SINGLETON_IS_STERILE_NULL_EDGE_CANDIDATE_ONLY"
	SupportDominantColorOrientationCandidate     = "CONDITIONAL_SUPPORT_E_PLUS_COLOR_BLOCK_COULD_LOCATE_DOMINANT_I3_IF_ORIENTATION_CERTIFIED"
	SupportRestQuartetOrientationCandidate       = "CONDITIONAL_SUPPORT_E_MINUS_W_BLOCK_COULD_LOCATE_REST_QUARTET_IF_ORIENTATION_CERTIFIED"
	SupportAggregateShadowIfOrientationCertified = "CONDITIONAL_SUPPORT_AGGREGATE_SHADOW_FINITE_BODY_LOCATION_IF_ORIENTATION_AND_COMPRESSION_CERTIFIED"

	FailureRightCharacterSplitStillSeal      = "FAILED_ROUTE_RIGHT_SOCKET_CHARACTER_SPLIT_REMAINS_SEAL_NOT_NATIVE_DERIVATION"
	FailureNoFullRhoFActionLedger            = "FAILED_ROUTE_NO_FULL_RHO_F_ACTION_LEDGER_CERTIFIED"
	FailureNoDFEdgeData                      = "FAILED_ROUTE_NO_D_F_EDGE_DATA_TO_CERTIFY_STERILE_PUNCTURE"
	FailureNoNullEdgeTheorem                 = "FAILED_ROUTE_NO_NULL_EDGE_THEOREM_FOR_EXCLUDED_SINGLETON"
	FailureNoSterilePunctureTheorem          = "FAILED_ROUTE_EXCLUDED_SINGLETON_NOT_CERTIFIED_AS_STERILE_PUNCTURE"
	FailurePunctureNotPhysicalParticle       = "FAILED_ROUTE_EXCLUDED_SINGLETON_NOT_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoRightNeutrinoTheorem            = "FAILED_ROUTE_NO_RIGHT_NEUTRINO_OR_STERILE_PARTICLE_THEOREM"
	FailureNoDominantColorOrientationTheorem = "FAILED_ROUTE_NO_DOMINANT_COLOR_ORIENTATION_THEOREM"
	FailureNoRestQuartetOrientationTheorem   = "FAILED_ROUTE_NO_REST_QUARTET_ORIENTATION_THEOREM"
	FailureNoDFOrHiggsOrientationSelector    = "FAILED_ROUTE_NO_D_F_OR_HIGGS_EDGE_ORIENTATION_SELECTOR_CERTIFIED"
	FailureNoBoundaryRestOrientationSelector = "FAILED_ROUTE_NO_BOUNDARY_REST_PRESSURE_ORIENTATION_SELECTOR_CERTIFIED"
	FailureNoTypedSocketOrientationMap       = "FAILED_ROUTE_NO_TYPED_SOCKET_ORIENTATION_MAP_CERTIFIED"
	FailureNoTypedCompressionMap             = "FAILED_ROUTE_NO_TYPED_PUNCTURED_SOCKET_COMPRESSION_MAP_CERTIFIED"
	FailureCompressionCandidateNotTheorem    = "FAILED_ROUTE_PUNCTURE_COMPLEMENT_IS_SUPPORT_ANATOMY_NOT_COMPRESSION_THEOREM"
	FailureNoAlphaDerivation                 = "FAILED_ROUTE_PUNCTURE_ORIENTATION_DOES_NOT_DERIVE_ALPHA_B"
	FailureAlphaStillSealed                  = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoTraceMagnitudeReadout           = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureCompressionNotYukawaMagnitude     = "FAILED_ROUTE_PUNCTURE_ORIENTATION_NOT_YUKAWA_MAGNITUDE_SOURCE"
	FailureNoAggregateCompressionMap         = "FAILED_ROUTE_NO_AGGREGATE_TRACE_COMPRESSION_MAP_CERTIFIED"
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

type InheritedCarrier struct {
	EExpression, WExpression, RightRectangleExpression string
	EDim, WDim, RightRectangleRank, HPartDim, HFDim    int
	Gate838BodyInherited, Gate839CompressionInherited  bool
	Gate840PunctureInherited                           bool
	Supports, Failures                                 []string
}

type ComplementLaw struct {
	ActiveExpression, PunctureExpression, FullExpression string
	DominantColorExpression, RestQuartetExpression       string
	ActiveRank, PunctureRank, FullRank                   int
	DominantColorRank, RestQuartetRank                   int
	Orthogonal, Complete, EightEqualsSevenPlusOne        bool
	SupportAnatomyCertified, CompressionTheoremCertified bool
	Supports, Failures                                   []string
}

type BMinusLCompensation struct {
	DominantColorTrace, RestQuartetTrace  float64
	ActiveTrace, PunctureTrace, FullTrace float64
	ActivePlusPunctureCancel              bool
	FullNeutral                           bool
	CompensatingPuncturePattern           bool
	Supports, Failures                    []string
}

type SterileNullEdgeAudit struct {
	CandidateExpression                                  string
	Rank                                                 int
	RightSocket, Leptonic, Colorless, ExcludedFromActive bool
	BMinusLTrace                                         float64
	DFEdgeDataAvailable                                  bool
	NullEdgeCertified                                    bool
	SterilePunctureCertified                             bool
	PhysicalParticleAssignmentCertified                  bool
	Supports, Failures                                   []string
}

type SocketOrientationAudit struct {
	DominantCandidateExpression, RestCandidateExpression      string
	DominantWouldSourceI3, RestWouldCarryW                    bool
	SameSocketContainsDominantColorAndPuncture                bool
	DominantOrientationCertified, RestOrientationCertified    bool
	DFOrHiggsSelectorCertified, BoundaryRestSelectorCertified bool
	OrientationMapCertified                                   bool
	Supports, Failures                                        []string
}

type AggregateLocation struct {
	CandidateOperator                                           string
	FiniteBodyLocationCandidate                                 bool
	TopBlockLocatedIfOrientation, RestBlockLocatedIfOrientation bool
	CompressionMapCertified                                     bool
	AlphaDerivedByCompression                                   bool
	TraceMagnitudeReadoutCertified                              bool
	R3, R4                                                      bool
	Supports, Failures                                          []string
}

type Impact struct {
	PunctureComplementLawCertified                   bool
	BMinusLCompensationFound                         bool
	SterileNullEdgeStillUncertified                  bool
	OrientationStillMissing                          bool
	CompressionMapStillMissing                       bool
	CanPromoteToR3, CanPromoteToR4                   bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs bool
	Classification                                   string
	Supports, Failures                               []string
}

type Firewalls struct {
	Enforced                               bool
	RightCharacterSplitStillSeal           bool
	NoFullRhoFActionLedger                 bool
	NoDFEdgeData                           bool
	NoNullEdgeTheorem                      bool
	NoSterilePunctureTheorem               bool
	PunctureNotPhysicalParticle            bool
	NoRightNeutrinoTheorem                 bool
	NoDominantOrientationTheorem           bool
	NoRestOrientationTheorem               bool
	NoDFOrHiggsOrientationSelector         bool
	NoBoundaryRestOrientationSelector      bool
	NoTypedSocketOrientationMap            bool
	NoTypedCompressionMap                  bool
	CompressionCandidateNotTheorem         bool
	NoAlphaDerivation                      bool
	AlphaSealed                            bool
	NoTraceMagnitudeReadout                bool
	CompressionNotYukawaMagnitude          bool
	NoAggregateCompressionMap              bool
	NoNEffUpdate, NoCYukawaUpdate          bool
	NoObservedYukawaFit, NoThreeGeneration bool
	NotR3, NotR4                           bool
	Verdict                                string
}

type Analysis struct {
	AuditID      string
	Ledger       Ledger
	Carrier      InheritedCarrier
	Complement   ComplementLaw
	BMinusL      BMinusLCompensation
	Sterile      SterileNullEdgeAudit
	Orientation  SocketOrientationAudit
	Location     AggregateLocation
	Impact       Impact
	Firewalls    Firewalls
	Truth, Final string
}

func BuildDefault() (Analysis, error) {
	fullTrace := float64(RightSlotDim) * (float64(LeptonBlockDim)*BMinusLLeptonWeight + float64(ColorBlockDim)*BMinusLColorWeight)
	if math.Abs(fullTrace) > 1e-15 {
		return Analysis{}, fmt.Errorf("full right lepto-color rectangle B-L trace not neutral: %.16g", fullTrace)
	}
	if ActiveSupportRank+PunctureRank != RightRectangleRank {
		return Analysis{}, fmt.Errorf("8=7+1 complement rank check failed")
	}

	ledger := Ledger{S: SBoundary, AlphaB: AlphaB, OperatorNEff: OperatorNEff, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true, R2PlusPlus: true, R3: false, R4: false, AlphaIsNative: false}
	carrier := InheritedCarrier{
		EExpression: "E=C_R^2 plus C_L^2", WExpression: "W=C_lepton plus C_color^3", RightRectangleExpression: "C_R^2 tensor W",
		EDim: EDim, WDim: WDim, RightRectangleRank: RightRectangleRank, HPartDim: HPartDim, HFDim: HFDim,
		Gate838BodyInherited: true, Gate839CompressionInherited: true, Gate840PunctureInherited: true,
		Supports: []string{SupportGate840PunctureInherited}, Failures: []string{FailureRightCharacterSplitStillSeal, FailureNoFullRhoFActionLedger},
	}
	complement := ComplementLaw{
		ActiveExpression: "Pi_7=(e_+ tensor P_3) plus (e_- tensor W)", PunctureExpression: "Pi_puncture=e_+ tensor P_1", FullExpression: "C_R^2 tensor W=Pi_7 plus Pi_puncture",
		DominantColorExpression: "e_+ tensor P_3", RestQuartetExpression: "e_- tensor W",
		ActiveRank: ActiveSupportRank, PunctureRank: PunctureRank, FullRank: RightRectangleRank, DominantColorRank: DominantColorRank, RestQuartetRank: RestQuartetRank,
		Orthogonal: true, Complete: true, EightEqualsSevenPlusOne: true, SupportAnatomyCertified: true, CompressionTheoremCertified: false,
		Supports: []string{SupportRightRectangleComplement, SupportActiveSupportRankSeven, SupportPunctureSingletonRankOne, SupportAggregateShadowIfOrientationCertified},
		Failures: []string{FailureCompressionCandidateNotTheorem, FailureNoTypedCompressionMap, FailureNoAggregateCompressionMap},
	}
	bminusl := BMinusLCompensation{
		DominantColorTrace:       float64(ColorBlockDim) * BMinusLColorWeight,
		RestQuartetTrace:         float64(LeptonBlockDim)*BMinusLLeptonWeight + float64(ColorBlockDim)*BMinusLColorWeight,
		ActiveTrace:              float64(ColorBlockDim)*BMinusLColorWeight + (float64(LeptonBlockDim)*BMinusLLeptonWeight + float64(ColorBlockDim)*BMinusLColorWeight),
		PunctureTrace:            float64(LeptonBlockDim) * BMinusLLeptonWeight,
		FullTrace:                fullTrace,
		ActivePlusPunctureCancel: true, FullNeutral: true, CompensatingPuncturePattern: true,
		Supports: []string{SupportBMinusLActivePlusOne, SupportBMinusLPunctureMinusOne, SupportBMinusLFullRightRectangleNeutral, SupportPunctureIsCompensatingSingleton},
	}
	if math.Abs(bminusl.DominantColorTrace-1) > 1e-15 || math.Abs(bminusl.RestQuartetTrace) > 1e-15 || math.Abs(bminusl.ActiveTrace-1) > 1e-15 || math.Abs(bminusl.PunctureTrace+1) > 1e-15 || math.Abs(bminusl.ActiveTrace+bminusl.PunctureTrace) > 1e-15 {
		return Analysis{}, fmt.Errorf("B-L compensating puncture balance failed")
	}
	sterile := SterileNullEdgeAudit{
		CandidateExpression: "e_+ tensor P_1", Rank: PunctureRank, RightSocket: true, Leptonic: true, Colorless: true, ExcludedFromActive: true,
		BMinusLTrace: bminusl.PunctureTrace, DFEdgeDataAvailable: false, NullEdgeCertified: false, SterilePunctureCertified: false, PhysicalParticleAssignmentCertified: false,
		Supports: []string{SupportSterileNullEdgeCandidate},
		Failures: []string{FailureNoDFEdgeData, FailureNoNullEdgeTheorem, FailureNoSterilePunctureTheorem, FailurePunctureNotPhysicalParticle, FailureNoRightNeutrinoTheorem},
	}
	orientation := SocketOrientationAudit{
		DominantCandidateExpression: "I_3 on e_+ tensor P_3", RestCandidateExpression: "alpha_B P_3 - 3 alpha_B^2(B-L) on e_- tensor W",
		DominantWouldSourceI3: true, RestWouldCarryW: true, SameSocketContainsDominantColorAndPuncture: true,
		DominantOrientationCertified: false, RestOrientationCertified: false, DFOrHiggsSelectorCertified: false, BoundaryRestSelectorCertified: false, OrientationMapCertified: false,
		Supports: []string{SupportDominantColorOrientationCandidate, SupportRestQuartetOrientationCandidate, SupportAggregateShadowIfOrientationCertified},
		Failures: []string{FailureNoDominantColorOrientationTheorem, FailureNoRestQuartetOrientationTheorem, FailureNoDFOrHiggsOrientationSelector, FailureNoBoundaryRestOrientationSelector, FailureNoTypedSocketOrientationMap},
	}
	location := AggregateLocation{
		CandidateOperator:           "H_total/T = I_{e_+ tensor P_3} plus [alpha_B P_3 - 3 alpha_B^2(B-L)]_{e_- tensor W}",
		FiniteBodyLocationCandidate: true, TopBlockLocatedIfOrientation: true, RestBlockLocatedIfOrientation: true,
		CompressionMapCertified: false, AlphaDerivedByCompression: false, TraceMagnitudeReadoutCertified: false, R3: false, R4: false,
		Supports: []string{SupportAggregateShadowIfOrientationCertified},
		Failures: []string{FailureNoTypedCompressionMap, FailureNoAlphaDerivation, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureCompressionNotYukawaMagnitude, FailureNoAggregateCompressionMap, FailureNotR3, FailureNotR4},
	}
	impact := Impact{PunctureComplementLawCertified: true, BMinusLCompensationFound: true, SterileNullEdgeStillUncertified: true, OrientationStillMissing: true, CompressionMapStillMissing: true, CanPromoteToR3: false, CanPromoteToR4: false, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, Classification: "right lepto-color 8=7+1 puncture complement certified as support anatomy; sterile/null-edge and socket orientation remain blocked; R2++ retained", Supports: []string{SupportRightRectangleComplement, SupportPunctureIsCompensatingSingleton, SupportSterileNullEdgeCandidate}, Failures: []string{FailureNoDFEdgeData, FailureNoDominantColorOrientationTheorem, FailureNoRestQuartetOrientationTheorem, FailureNoTypedCompressionMap, FailureNoTraceMagnitudeReadout, FailureNoNEffUpdate, FailureNoCYukawaUpdate, FailureNotR3, FailureNotR4}}
	firewalls := Firewalls{Enforced: true, RightCharacterSplitStillSeal: true, NoFullRhoFActionLedger: true, NoDFEdgeData: true, NoNullEdgeTheorem: true, NoSterilePunctureTheorem: true, PunctureNotPhysicalParticle: true, NoRightNeutrinoTheorem: true, NoDominantOrientationTheorem: true, NoRestOrientationTheorem: true, NoDFOrHiggsOrientationSelector: true, NoBoundaryRestOrientationSelector: true, NoTypedSocketOrientationMap: true, NoTypedCompressionMap: true, CompressionCandidateNotTheorem: true, NoAlphaDerivation: true, AlphaSealed: true, NoTraceMagnitudeReadout: true, CompressionNotYukawaMagnitude: true, NoAggregateCompressionMap: true, NoNEffUpdate: true, NoCYukawaUpdate: true, NoObservedYukawaFit: true, NoThreeGeneration: true, NotR3: true, NotR4: true, Verdict: StatusFirewallGate841}
	return Analysis{AuditID: AuditID, Ledger: ledger, Carrier: carrier, Complement: complement, BMinusL: bminusl, Sterile: sterile, Orientation: orientation, Location: location, Impact: impact, Firewalls: firewalls, Truth: "Gate 841 upgrades Gate 840's punctured right rectangle into an explicit 8=7+1 complement with a B-L compensating singleton, while refusing to certify sterile/null-edge status or dominant/rest orientation without D_F edge data or a native orientation theorem.", Final: "Verdict: puncture-complement anatomy and B-L compensation certified; excluded singleton remains a sterile/null-edge candidate only; no orientation theorem, no typed compression map, no alpha derivation, no trace-magnitude readout, no R3/R4 promotion."}, nil
}

func Statuses() []string {
	return []string{StatusGate840Inherited, StatusComplementLawAudited, StatusEightEqualsSevenPlusOne, StatusPunctureRankAudited, StatusBMinusLCompensationAudited, StatusSterileNullEdgeAudited, StatusDominantOrientationAudited, StatusRestOrientationAudited, StatusAggregateLocationCandidateAudited, StatusMagnitudeFirewallPreserved, StatusAlphaStillSealed, StatusOfficialLedgersFrozen, StatusR2PlusPlusRetained, StatusNoObservedDataUsed, StatusFirewallGate841, SupportGate840PunctureInherited, SupportRightRectangleComplement, SupportActiveSupportRankSeven, SupportPunctureSingletonRankOne, SupportBMinusLActivePlusOne, SupportBMinusLPunctureMinusOne, SupportBMinusLFullRightRectangleNeutral, SupportPunctureIsCompensatingSingleton, SupportSterileNullEdgeCandidate, SupportDominantColorOrientationCandidate, SupportRestQuartetOrientationCandidate, SupportAggregateShadowIfOrientationCertified, FailureRightCharacterSplitStillSeal, FailureNoFullRhoFActionLedger, FailureNoDFEdgeData, FailureNoNullEdgeTheorem, FailureNoSterilePunctureTheorem, FailurePunctureNotPhysicalParticle, FailureNoRightNeutrinoTheorem, FailureNoDominantColorOrientationTheorem, FailureNoRestQuartetOrientationTheorem, FailureNoDFOrHiggsOrientationSelector, FailureNoBoundaryRestOrientationSelector, FailureNoTypedSocketOrientationMap, FailureNoTypedCompressionMap, FailureCompressionCandidateNotTheorem, FailureNoAlphaDerivation, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureCompressionNotYukawaMagnitude, FailureNoAggregateCompressionMap, FailureNoNEffUpdate, FailureNoCYukawaUpdate, FailureNoObservedYukawaFit, FailureNoThreeGenerationTheorem, FailureNotR3, FailureNotR4}
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("ledger: s=%.16g alpha_B=%.16g operator_N_eff=%.16g official_N_eff=%.16g frozen=%t R2++=%t R3=%t R4=%t alpha_native=%t", l.S, l.AlphaB, l.OperatorNEff, l.OfficialNEff, l.OfficialFrozen, l.R2PlusPlus, l.R3, l.R4, l.AlphaIsNative)
}

func FormatCarrier(c InheritedCarrier) string {
	return fmt.Sprintf("carrier: %s, %s, right=%s dim(E)=%d dim(W)=%d rank(right rectangle)=%d dim(H_part)=%d dim(H_F)=%d inherited838=%t inherited839=%t inherited840=%t", c.EExpression, c.WExpression, c.RightRectangleExpression, c.EDim, c.WDim, c.RightRectangleRank, c.HPartDim, c.HFDim, c.Gate838BodyInherited, c.Gate839CompressionInherited, c.Gate840PunctureInherited)
}

func FormatComplement(c ComplementLaw) string {
	return fmt.Sprintf("complement: full=%s active=%s puncture=%s ranks active/puncture/full=%d/%d/%d dominant/rest=%d/%d orthogonal=%t complete=%t 8=7+1=%t theorem=%t", c.FullExpression, c.ActiveExpression, c.PunctureExpression, c.ActiveRank, c.PunctureRank, c.FullRank, c.DominantColorRank, c.RestQuartetRank, c.Orthogonal, c.Complete, c.EightEqualsSevenPlusOne, c.CompressionTheoremCertified)
}

func FormatBMinusL(b BMinusLCompensation) string {
	return fmt.Sprintf("B-L compensation: dominant=%.12g rest=%.12g active=%.12g puncture=%.12g full=%.12g cancel=%t neutral=%t compensating=%t", b.DominantColorTrace, b.RestQuartetTrace, b.ActiveTrace, b.PunctureTrace, b.FullTrace, b.ActivePlusPunctureCancel, b.FullNeutral, b.CompensatingPuncturePattern)
}

func FormatSterile(s SterileNullEdgeAudit) string {
	return fmt.Sprintf("sterile-null-edge: candidate=%s rank=%d right=%t leptonic=%t colorless=%t excluded=%t B-L=%.12g D_F_data=%t null_edge=%t sterile_certified=%t physical_assignment=%t", s.CandidateExpression, s.Rank, s.RightSocket, s.Leptonic, s.Colorless, s.ExcludedFromActive, s.BMinusLTrace, s.DFEdgeDataAvailable, s.NullEdgeCertified, s.SterilePunctureCertified, s.PhysicalParticleAssignmentCertified)
}

func FormatOrientation(o SocketOrientationAudit) string {
	return fmt.Sprintf("orientation: dominant=%s rest=%s dominant_I3=%t rest_W=%t same_socket_dominant_and_puncture=%t dominant_cert=%t rest_cert=%t D_F/Higgs_selector=%t boundary_selector=%t orientation_map=%t", o.DominantCandidateExpression, o.RestCandidateExpression, o.DominantWouldSourceI3, o.RestWouldCarryW, o.SameSocketContainsDominantColorAndPuncture, o.DominantOrientationCertified, o.RestOrientationCertified, o.DFOrHiggsSelectorCertified, o.BoundaryRestSelectorCertified, o.OrientationMapCertified)
}

func FormatLocation(l AggregateLocation) string {
	return fmt.Sprintf("location: %s finite_body_candidate=%t top_if_oriented=%t rest_if_oriented=%t compression_map=%t alpha_derived=%t trace_readout=%t R3=%t R4=%t", l.CandidateOperator, l.FiniteBodyLocationCandidate, l.TopBlockLocatedIfOrientation, l.RestBlockLocatedIfOrientation, l.CompressionMapCertified, l.AlphaDerivedByCompression, l.TraceMagnitudeReadoutCertified, l.R3, l.R4)
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("impact: complement=%t B-L=%t sterile_uncertified=%t orientation_missing=%t compression_missing=%t promote_R3=%t promote_R4=%t update_N_eff=%t update_C_Yukawa=%t update_C_Higgs=%t classification=%q", i.PunctureComplementLawCertified, i.BMinusLCompensationFound, i.SterileNullEdgeStillUncertified, i.OrientationStillMissing, i.CompressionMapStillMissing, i.CanPromoteToR3, i.CanPromoteToR4, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.Classification)
}

func containsAll(got, want []string) bool {
	seen := map[string]bool{}
	for _, g := range got {
		seen[g] = true
	}
	for _, w := range want {
		if !seen[w] {
			return false
		}
	}
	return true
}

func containsText(xs []string, needle string) bool {
	return strings.Contains(strings.Join(xs, "\n"), needle)
}
