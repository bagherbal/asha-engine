// Package generation2dominantsocketleptocolortracecompressionmapaudit implements
// Gate 839: DominantSocket LeptoColor Trace-Compression Map Audit.
//
// Gate 839 follows Gate 838's sealed finite-sector body
// H_part=(C_R^2 plus C_L^2) tensor (C_lepton plus C^3_color). Gate 838
// constructed the coarse particle-side projector ledger with ranks 2,6,2,6.
// Gate 839 audits the next, sharper question: can the R2++ aggregate carrier
// I_3 plus W be obtained as a socket-level trace compression of E tensor W by
// selecting one dominant color socket e_t tensor P_3 and one rest lepto-color
// socket e_r tensor W? The gate certifies the candidate rank anatomy 3+4 if
// rank-one socket selectors exist, but it does not certify those selectors, the
// compression map, the alpha source, trace magnitudes, R3, or a native Yukawa
// theorem.
package generation2dominantsocketleptocolortracecompressionmapaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE839-DOMINANT-SOCKET-LEPTOCOLOR-TRACE-COMPRESSION-MAP-AUDIT"

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
	HFDim              = 32

	CandidateSocketRank = 1
	TopCompressionRank  = CandidateSocketRank * ColorBlockDim
	RestCompressionRank = CandidateSocketRank * WDim
	AggregateRank       = TopCompressionRank + RestCompressionRank

	BMinusLLeptonWeight = -1.0
	BMinusLColorWeight  = 1.0 / 3.0

	StatusGate838Inherited                  = "PASS_GATE838_SEALED_FINITE_SECTOR_BODY_INHERITED"
	StatusFiniteBodyReverified              = "PASS_H_PART_EQUALS_E_TENSOR_W_BODY_REVERIFIED"
	StatusWCarrierReverified                = "PASS_W_CARRIER_AND_B_MINUS_L_REVERIFIED"
	StatusCompressionCandidateFormulated    = "PASS_SOCKET_COMPRESSION_CANDIDATE_FORMULATED"
	StatusRanksMatchI3PlusWIfSelectorsExist = "PASS_RANKS_MATCH_I3_PLUS_W_IF_RANK_ONE_SOCKETS_EXIST"
	StatusTopCandidateAudited               = "PASS_DOMINANT_COLOR_TRIPLET_CANDIDATE_AUDITED"
	StatusRestCandidateAudited              = "PASS_REST_LEPTOCOLOR_QUARTET_CANDIDATE_AUDITED"
	StatusFineSocketRequirementAudited      = "PASS_FINE_SOCKET_SELECTOR_REQUIREMENT_AUDITED"
	StatusNoK7Promotion                     = "PASS_SEVEN_COUNT_REINTERPRETED_AS_COMPRESSION_CANDIDATE_NOT_K7_THEOREM"
	StatusCompressionNonCircularityAudited  = "PASS_COMPRESSION_MAP_NONCIRCULARITY_AUDITED"
	StatusAggregateShadowOnly               = "PASS_AGGREGATE_OPERATOR_CLASSIFIED_AS_POSSIBLE_TRACE_SHADOW_ONLY"
	StatusMagnitudeFirewallPreserved        = "PASS_SOCKET_COMPRESSION_NOT_TRACE_MAGNITUDE_READOUT"
	StatusAlphaStillSealed                  = "PASS_ALPHA_B_REMAINS_SEALED_AFTER_COMPRESSION_AUDIT"
	StatusOfficialLedgersFrozen             = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusR2PlusPlusRetained                = "PASS_R2_PLUS_PLUS_RETAINED_NOT_R3_OR_R4"
	StatusNoObservedDataUsed                = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusFirewallGate839                   = "FIREWALL_PRESERVED_GATE839_SOCKET_COMPRESSION_OBSTRUCTION"

	SupportGate838SectorBodyInherited        = "CONDITIONAL_SUPPORT_GATE838_SEALED_COARSE_SECTOR_BODY_INHERITED"
	SupportSocketCompressionCandidate        = "CONDITIONAL_SUPPORT_SOCKET_COMPRESSION_CANDIDATE_ON_E_TENSOR_W"
	SupportTopAsEtTensorP3IfSelectorExists   = "CONDITIONAL_SUPPORT_TOP_BLOCK_AS_E_TENSOR_P3_IF_DOMINANT_SOCKET_SELECTOR_EXISTS"
	SupportRestAsErTensorWIfSelectorExists   = "CONDITIONAL_SUPPORT_REST_BLOCK_AS_E_TENSOR_W_IF_REST_SOCKET_SELECTOR_EXISTS"
	SupportRanksThreePlusFour                = "CONDITIONAL_SUPPORT_RANK_3_PLUS_4_FROM_ONE_COLOR_SOCKET_PLUS_ONE_LEPTOCOLOR_SOCKET"
	SupportBMinusLActsOnRestW                = "CONDITIONAL_SUPPORT_B_MINUS_L_REST_TRANSFER_ACTS_NATURALLY_ON_W_FACTOR"
	SupportAggregateAsTraceShadowIfSelectors = "CONDITIONAL_SUPPORT_AGGREGATE_OPERATOR_COULD_BE_TRACE_COMPRESSION_SHADOW_IF_SELECTORS_CERTIFIED"
	SupportSevenCompressionNotK7             = "CONDITIONAL_SUPPORT_SEVEN_COUNT_HAS_FINITE_CARRIER_COMPRESSION_CANDIDATE_NOT_K7"
	SupportSectorBodyBeforeCompression       = "CONDITIONAL_SUPPORT_FINITE_SECTOR_BODY_PRECEDES_AGGREGATE_COMPRESSION"

	FailureNoFineSocketProjectors         = "FAILED_ROUTE_NO_FINE_SOCKET_PROJECTORS_CERTIFIED"
	FailureSocketAtomsBasisDependent      = "FAILED_ROUTE_RANK_ONE_SOCKET_PROJECTORS_ARE_BASIS_DEPENDENT_WITHOUT_SELECTOR"
	FailureNoDominantColorSocketSelector  = "FAILED_ROUTE_NO_DOMINANT_COLOR_SOCKET_SELECTOR"
	FailureNoRestLeptoColorSocketSelector = "FAILED_ROUTE_NO_REST_LEPTOCOLOR_SOCKET_SELECTOR"
	FailureEtErNotCanonical               = "FAILED_ROUTE_E_T_AND_E_R_NOT_CANONICALLY_SELECTED_BY_GATE839"
	FailureNoSocketPairCompressionMap     = "FAILED_ROUTE_NO_TYPED_SOCKET_PAIR_COMPRESSION_MAP_CERTIFIED"
	FailureCompressionCandidateNotTheorem = "FAILED_ROUTE_COMPRESSION_CANDIDATE_IS_RANK_ANATOMY_NOT_THEOREM"
	FailureNoDFOrHiggsSocketSelector      = "FAILED_ROUTE_NO_D_F_OR_HIGGS_EDGE_SELECTOR_FOR_DOMINANT_SOCKET_CERTIFIED"
	FailureNoBoundaryRestSocketSelector   = "FAILED_ROUTE_NO_BOUNDARY_REST_PRESSURE_SOCKET_SELECTOR_CERTIFIED"
	FailureAggregateNotCompressionTheorem = "FAILED_ROUTE_AGGREGATE_OPERATOR_NOT_FINITE_BODY_COMPRESSION_THEOREM"
	FailureSevenNotK7                     = "FAILED_ROUTE_SEVEN_COMPRESSION_RANK_NOT_K7_PROJECTOR_THEOREM"
	FailureNoAggregateToK7                = "FAILED_ROUTE_NO_AGGREGATE_COMPRESSION_TO_K7_MAP_CERTIFIED"
	FailureNoTraceMagnitudeReadout        = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureCompressionNotYukawaMagnitude  = "FAILED_ROUTE_SOCKET_COMPRESSION_NOT_YUKAWA_MAGNITUDE_SOURCE"
	FailureAlphaStillSealed               = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoAlphaDerivation              = "FAILED_ROUTE_SOCKET_COMPRESSION_DOES_NOT_DERIVE_ALPHA_B"
	FailureNoNEffUpdate                   = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaUpdate                = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoObservedYukawaFit            = "FAILED_ROUTE_NO_OBSERVED_YUKAWA_FITTING_ALLOWED"
	FailureNoParticleAssignment           = "FAILED_ROUTE_NO_STANDARD_MODEL_PARTICLE_ASSIGNMENT_FROM_SOCKET_COMPRESSION"
	FailureNoThreeGenerationTheorem       = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
	FailureNotR3                          = "FAILED_ROUTE_R2_PLUS_PLUS_NOT_R3"
	FailureNotR4                          = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	S, AlphaB                       float64
	OperatorNEff, OfficialNEff      float64
	OfficialCYukawa, OfficialCHiggs float64
	OfficialFrozen                  bool
	R2PlusPlus, R3, R4              bool
	AlphaIsNative                   bool
}

type FiniteBody struct {
	EExpression, WExpression, HPartExpression string
	EDim, WDim, HPartDim, HFDim               int
	RightSlotDim, LeftSlotDim                 int
	LeptonDim, ColorDim                       int
	Gate838Inherited, CoarseLedgerExists      bool
	Supports, Failures                        []string
}

type WCarrier struct {
	P1Rank, P3Rank, Dim                     int
	P1P3Orthogonal, P1PlusP3CompletesW      bool
	BMinusLLeptonWeight, BMinusLColorWeight float64
	BMinusLTraceZero                        bool
	BMinusLRestActionOnW                    bool
	Supports, Failures                      []string
}

type SocketRefinement struct {
	EExpression                              string
	RankOneSocketProjectorsPossible          bool
	FineSocketProjectorsCertified            bool
	SocketAtomsBasisDependentWithoutSelector bool
	DominantSelectorCertified                bool
	RestSelectorCertified                    bool
	EtErCanonical                            bool
	PotentialSourcesAudited                  []string
	Supports, Failures                       []string
}

type CompressionCandidate struct {
	TopExpression, RestExpression, AggregateExpression string
	SocketRank, TopRank, RestRank, AggregateRank       int
	MatchesI3PlusW                                     bool
	TopSelectorCertified, RestSelectorCertified        bool
	CompressionMapCertified                            bool
	NonCircular                                        bool
	UsesObservedData                                   bool
	IsTheorem                                          bool
	Supports, Failures                                 []string
}

type ShadowOperator struct {
	TopOperator, RestOperator, TotalOperator string
	TopIsIdentityI3Candidate                 bool
	RestUsesBMinusLTransferOnW               bool
	AlphaB                                   float64
	AlphaDerivedByCompression                bool
	TraceMagnitudeReadoutCertified           bool
	AggregateOperatorIsSectorLedger          bool
	R3, R4                                   bool
	Supports, Failures                       []string
}

type Impact struct {
	CompressionCandidateFormulated          bool
	FiniteBodyLocationForAggregateSuggested bool
	RankAnatomyExplainedConditionally       bool
	SelectorsMissing                        bool
	CompressionMapMissing                   bool
	CanPromoteToR3, CanPromoteToR4          bool
	CanUpdateNEff, CanUpdateCYukawa         bool
	CanUpdateCHiggs                         bool
	Classification                          string
	Supports, Failures                      []string
}

type Firewalls struct {
	Enforced                       bool
	NoFineSocketProjectors         bool
	SocketAtomsBasisDependent      bool
	NoDominantSelector             bool
	NoRestSelector                 bool
	NoCompressionMap               bool
	CompressionCandidateNotTheorem bool
	NoDFOrHiggsSelector            bool
	NoBoundaryRestSelector         bool
	AggregateNotCompressionTheorem bool
	SevenNotK7                     bool
	NoTraceMagnitudeReadout        bool
	CompressionNotYukawaMagnitude  bool
	AlphaSealed                    bool
	NoAlphaDerivation              bool
	NoNEffUpdate                   bool
	NoCYukawaUpdate                bool
	NoObservedYukawaFit            bool
	NoParticleAssignment           bool
	NoThreeGeneration              bool
	NotR3, NotR4                   bool
	Verdict                        string
}

type Analysis struct {
	AuditID     string
	Ledger      Ledger
	Body        FiniteBody
	W           WCarrier
	Sockets     SocketRefinement
	Compression CompressionCandidate
	Shadow      ShadowOperator
	Impact      Impact
	Firewalls   Firewalls
	Truth       string
	Final       string
}

func BuildDefault() (Analysis, error) {
	wTrace := float64(LeptonBlockDim)*BMinusLLeptonWeight + float64(ColorBlockDim)*BMinusLColorWeight
	if math.Abs(wTrace) > 1e-15 {
		return Analysis{}, fmt.Errorf("B-L trace on W is not zero: %.17g", wTrace)
	}
	if HPartDim != ElectroweakSlotDim*WDim || HFDim != 2*HPartDim || AggregateRank != 7 {
		return Analysis{}, fmt.Errorf("dimension invariant failed")
	}

	ledger := Ledger{
		S: SBoundary, AlphaB: AlphaB,
		OperatorNEff: OperatorNEff, OfficialNEff: OfficialNEff,
		OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs,
		OfficialFrozen: true, R2PlusPlus: true, R3: false, R4: false, AlphaIsNative: false,
	}
	body := FiniteBody{
		EExpression: "E=C_R^2 plus C_L^2", WExpression: "W=C_lepton plus C_color^3",
		HPartExpression: "H_part=E tensor W", EDim: ElectroweakSlotDim, WDim: WDim,
		HPartDim: HPartDim, HFDim: HFDim, RightSlotDim: RightSlotDim, LeftSlotDim: LeftSlotDim,
		LeptonDim: LeptonBlockDim, ColorDim: ColorBlockDim, Gate838Inherited: true, CoarseLedgerExists: true,
		Supports: []string{SupportGate838SectorBodyInherited, SupportSectorBodyBeforeCompression},
	}
	w := WCarrier{
		P1Rank: LeptonBlockDim, P3Rank: ColorBlockDim, Dim: WDim,
		P1P3Orthogonal: true, P1PlusP3CompletesW: true,
		BMinusLLeptonWeight: BMinusLLeptonWeight, BMinusLColorWeight: BMinusLColorWeight,
		BMinusLTraceZero: true, BMinusLRestActionOnW: true,
		Supports: []string{SupportBMinusLActsOnRestW},
	}
	sockets := SocketRefinement{
		EExpression:                              "E=C_R^2 plus C_L^2",
		RankOneSocketProjectorsPossible:          true,
		FineSocketProjectorsCertified:            false,
		SocketAtomsBasisDependentWithoutSelector: true,
		DominantSelectorCertified:                false,
		RestSelectorCertified:                    false,
		EtErCanonical:                            false,
		PotentialSourcesAudited: []string{
			"D_F symbolic edge skeleton",
			"chirality/right-left socket structure",
			"finite one-form Higgs edge support",
			"top-dominant trace atom seal",
			"boundary/rest-pressure split",
		},
		Failures: []string{FailureNoFineSocketProjectors, FailureSocketAtomsBasisDependent, FailureNoDominantColorSocketSelector, FailureNoRestLeptoColorSocketSelector, FailureEtErNotCanonical, FailureNoDFOrHiggsSocketSelector, FailureNoBoundaryRestSocketSelector},
	}
	compression := CompressionCandidate{
		TopExpression: "Pi_top=e_t tensor P_3", RestExpression: "Pi_rest=e_r tensor I_W",
		AggregateExpression: "C_agg(E tensor W)=(e_t tensor P_3) plus (e_r tensor W)",
		SocketRank:          CandidateSocketRank, TopRank: TopCompressionRank, RestRank: RestCompressionRank,
		AggregateRank: AggregateRank, MatchesI3PlusW: true,
		TopSelectorCertified: false, RestSelectorCertified: false, CompressionMapCertified: false,
		NonCircular: true, UsesObservedData: false, IsTheorem: false,
		Supports: []string{SupportSocketCompressionCandidate, SupportTopAsEtTensorP3IfSelectorExists, SupportRestAsErTensorWIfSelectorExists, SupportRanksThreePlusFour, SupportSevenCompressionNotK7, SupportAggregateAsTraceShadowIfSelectors},
		Failures: []string{FailureNoDominantColorSocketSelector, FailureNoRestLeptoColorSocketSelector, FailureNoSocketPairCompressionMap, FailureCompressionCandidateNotTheorem, FailureAggregateNotCompressionTheorem, FailureSevenNotK7, FailureNoAggregateToK7},
	}
	shadow := ShadowOperator{
		TopOperator: "I_{e_t tensor P_3} congruent I_3", RestOperator: "alpha_B P_3 - 3 alpha_B^2(B-L) on e_r tensor W",
		TotalOperator: "I_3 plus [alpha_B P_3 - 3 alpha_B^2(B-L)]", TopIsIdentityI3Candidate: true,
		RestUsesBMinusLTransferOnW: true, AlphaB: AlphaB, AlphaDerivedByCompression: false,
		TraceMagnitudeReadoutCertified: false, AggregateOperatorIsSectorLedger: false, R3: false, R4: false,
		Supports: []string{SupportAggregateAsTraceShadowIfSelectors, SupportBMinusLActsOnRestW},
		Failures: []string{FailureNoTraceMagnitudeReadout, FailureCompressionNotYukawaMagnitude, FailureAlphaStillSealed, FailureNoAlphaDerivation, FailureNotR3, FailureNotR4},
	}
	impact := Impact{
		CompressionCandidateFormulated:          true,
		FiniteBodyLocationForAggregateSuggested: true,
		RankAnatomyExplainedConditionally:       true,
		SelectorsMissing:                        true,
		CompressionMapMissing:                   true,
		CanPromoteToR3:                          false, CanPromoteToR4: false,
		CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false,
		Classification: "socket-compression rank anatomy candidate; selectors and compression map missing; R2++ retained",
		Supports:       []string{SupportSocketCompressionCandidate, SupportAggregateAsTraceShadowIfSelectors},
		Failures:       []string{FailureNoFineSocketProjectors, FailureNoSocketPairCompressionMap, FailureNoTraceMagnitudeReadout, FailureAlphaStillSealed, FailureNoNEffUpdate, FailureNoCYukawaUpdate, FailureNoObservedYukawaFit, FailureNoParticleAssignment, FailureNoThreeGenerationTheorem, FailureNotR3, FailureNotR4},
	}
	firewalls := Firewalls{
		Enforced:               true,
		NoFineSocketProjectors: true, SocketAtomsBasisDependent: true,
		NoDominantSelector: true, NoRestSelector: true, NoCompressionMap: true,
		CompressionCandidateNotTheorem: true, NoDFOrHiggsSelector: true, NoBoundaryRestSelector: true,
		AggregateNotCompressionTheorem: true, SevenNotK7: true,
		NoTraceMagnitudeReadout: true, CompressionNotYukawaMagnitude: true,
		AlphaSealed: true, NoAlphaDerivation: true,
		NoNEffUpdate: true, NoCYukawaUpdate: true, NoObservedYukawaFit: true,
		NoParticleAssignment: true, NoThreeGeneration: true, NotR3: true, NotR4: true,
		Verdict: StatusFirewallGate839,
	}
	return Analysis{
		AuditID: AuditID, Ledger: ledger, Body: body, W: w, Sockets: sockets,
		Compression: compression, Shadow: shadow, Impact: impact, Firewalls: firewalls,
		Truth: "Gate 839 identifies the finite-body compression candidate E tensor W -> (e_t tensor P_3) plus (e_r tensor W), but rank-one socket selectors are not certified.",
		Final: "Verdict: conditional socket-compression rank anatomy support; no DominantSocketSelector, no RestSocketSelector, no typed compression map, no alpha derivation, no trace-magnitude readout, no R3/R4 promotion.",
	}, nil
}

func Statuses() []string {
	return []string{
		StatusGate838Inherited, StatusFiniteBodyReverified, StatusWCarrierReverified,
		StatusCompressionCandidateFormulated, StatusRanksMatchI3PlusWIfSelectorsExist,
		StatusTopCandidateAudited, StatusRestCandidateAudited, StatusFineSocketRequirementAudited,
		StatusNoK7Promotion, StatusCompressionNonCircularityAudited, StatusAggregateShadowOnly,
		StatusMagnitudeFirewallPreserved, StatusAlphaStillSealed, StatusOfficialLedgersFrozen,
		StatusR2PlusPlusRetained, StatusNoObservedDataUsed, StatusFirewallGate839,
		SupportGate838SectorBodyInherited, SupportSocketCompressionCandidate, SupportTopAsEtTensorP3IfSelectorExists,
		SupportRestAsErTensorWIfSelectorExists, SupportRanksThreePlusFour, SupportBMinusLActsOnRestW,
		SupportAggregateAsTraceShadowIfSelectors, SupportSevenCompressionNotK7, SupportSectorBodyBeforeCompression,
		FailureNoFineSocketProjectors, FailureSocketAtomsBasisDependent, FailureNoDominantColorSocketSelector,
		FailureNoRestLeptoColorSocketSelector, FailureEtErNotCanonical, FailureNoSocketPairCompressionMap,
		FailureCompressionCandidateNotTheorem, FailureNoDFOrHiggsSocketSelector, FailureNoBoundaryRestSocketSelector,
		FailureAggregateNotCompressionTheorem, FailureSevenNotK7, FailureNoAggregateToK7, FailureNoTraceMagnitudeReadout,
		FailureCompressionNotYukawaMagnitude, FailureAlphaStillSealed, FailureNoAlphaDerivation,
		FailureNoNEffUpdate, FailureNoCYukawaUpdate, FailureNoObservedYukawaFit, FailureNoParticleAssignment,
		FailureNoThreeGenerationTheorem, FailureNotR3, FailureNotR4,
	}
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("ledger: s=%.16g alpha_B=%.16g operator_N_eff=%.16g official_N_eff=%.16g frozen=%t R2++=%t R3=%t R4=%t alpha_native=%t", l.S, l.AlphaB, l.OperatorNEff, l.OfficialNEff, l.OfficialFrozen, l.R2PlusPlus, l.R3, l.R4, l.AlphaIsNative)
}

func FormatBody(b FiniteBody) string {
	return fmt.Sprintf("body: %s, %s, %s dim(E)=%d dim(W)=%d dim(H_part)=%d dim(H_F)=%d inherited=%t coarse_ledger=%t", b.EExpression, b.WExpression, b.HPartExpression, b.EDim, b.WDim, b.HPartDim, b.HFDim, b.Gate838Inherited, b.CoarseLedgerExists)
}

func FormatW(w WCarrier) string {
	return fmt.Sprintf("W: rank(P1)=%d rank(P3)=%d dim=%d B-L=(%.12g,%.12g) trace_zero=%t rest_action_on_W=%t", w.P1Rank, w.P3Rank, w.Dim, w.BMinusLLeptonWeight, w.BMinusLColorWeight, w.BMinusLTraceZero, w.BMinusLRestActionOnW)
}

func FormatSockets(s SocketRefinement) string {
	return fmt.Sprintf("sockets: %s rank_one_possible=%t fine_certified=%t basis_dependent=%t dominant_selector=%t rest_selector=%t canonical_et_er=%t sources=[%s]", s.EExpression, s.RankOneSocketProjectorsPossible, s.FineSocketProjectorsCertified, s.SocketAtomsBasisDependentWithoutSelector, s.DominantSelectorCertified, s.RestSelectorCertified, s.EtErCanonical, strings.Join(s.PotentialSourcesAudited, ";"))
}

func FormatCompression(c CompressionCandidate) string {
	return fmt.Sprintf("compression: %s, %s -> ranks %d+%d=%d matches_I3_plus_W=%t top_selector=%t rest_selector=%t map_certified=%t noncircular=%t observed_data=%t theorem=%t", c.TopExpression, c.RestExpression, c.TopRank, c.RestRank, c.AggregateRank, c.MatchesI3PlusW, c.TopSelectorCertified, c.RestSelectorCertified, c.CompressionMapCertified, c.NonCircular, c.UsesObservedData, c.IsTheorem)
}

func FormatShadow(s ShadowOperator) string {
	return fmt.Sprintf("shadow: top=%s rest=%s total=%s alpha=%.16g alpha_derived=%t trace_readout=%t sector_ledger=%t R3=%t R4=%t", s.TopOperator, s.RestOperator, s.TotalOperator, s.AlphaB, s.AlphaDerivedByCompression, s.TraceMagnitudeReadoutCertified, s.AggregateOperatorIsSectorLedger, s.R3, s.R4)
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("impact: candidate=%t body_location=%t rank_anatomy=%t selectors_missing=%t compression_missing=%t promote_R3=%t promote_R4=%t update_N_eff=%t update_C_Yukawa=%t update_C_Higgs=%t classification=%q", i.CompressionCandidateFormulated, i.FiniteBodyLocationForAggregateSuggested, i.RankAnatomyExplainedConditionally, i.SelectorsMissing, i.CompressionMapMissing, i.CanPromoteToR3, i.CanPromoteToR4, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.Classification)
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
