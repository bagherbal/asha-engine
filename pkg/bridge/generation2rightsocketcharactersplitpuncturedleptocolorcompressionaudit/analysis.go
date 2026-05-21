// Package generation2rightsocketcharactersplitpuncturedleptocolorcompressionaudit implements
// Gate 840: RightSocket Character Split and Punctured LeptoColor Compression Audit.
//
// Gate 840 follows Gate 839's socket-compression obstruction. Gate 839 found
// that the R2++ aggregate rank 3+4 could live inside E tensor W as
// (e_t tensor P_3) plus (e_r tensor W), but it lacked rank-one socket
// projectors. Gate 840 audits the sharper candidate that the right socket pair
// C_R^2 is split by two represented C-characters, schematically lambda and
// conjugate(lambda). This can source an unordered rank-one socket pair
// e_+, e_- as a representation seal and exposes the punctured 2x4 right
// lepto-color rectangle. The gate still blocks orientation, a typed compression
// theorem, alpha derivation, trace magnitudes, R3, and R4.
package generation2rightsocketcharactersplitpuncturedleptocolorcompressionaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE840-RIGHT-SOCKET-CHARACTER-SPLIT-PUNCTURED-LEPTOCOLOR-COMPRESSION-AUDIT"

	SBoundary       = 0.0012924448188162962
	AlphaB          = 0.0003878958469680527
	OperatorNEff    = 3.002327375081808
	OfficialNEff    = 3.0023273474722147
	OfficialCYukawa = 0.9992248188812008
	OfficialCHiggs  = 1.0372205204048603

	LeptonBlockDim = 1
	ColorBlockDim  = 3
	WDim           = 4
	RightSlotDim   = 2
	LeftSlotDim    = 2
	EDim           = 4
	HPartDim       = 16
	HFDim          = 32

	CharacterSocketRank = 1
	RightRectangleRank  = RightSlotDim * WDim
	TopColorRank        = CharacterSocketRank * ColorBlockDim
	RestQuartetRank     = CharacterSocketRank * WDim
	SelectedRank        = TopColorRank + RestQuartetRank
	ExcludedRank        = CharacterSocketRank * LeptonBlockDim

	BMinusLLeptonWeight = -1.0
	BMinusLColorWeight  = 1.0 / 3.0

	StatusGate839Inherited              = "PASS_GATE839_SOCKET_COMPRESSION_CANDIDATE_INHERITED"
	StatusRightRectangleAudited         = "PASS_RIGHT_LEPTOCOLOR_RECTANGLE_AUDITED"
	StatusRightCharacterSplitAudited    = "PASS_RIGHT_SOCKET_CHARACTER_SPLIT_AUDITED"
	StatusRankOneCharacterSockets       = "PASS_RANK_ONE_CHARACTER_SOCKET_PAIR_SOURCE_TYPED_BY_SEAL"
	StatusOrientationFirewallAudited    = "PASS_DOMINANT_REST_ORIENTATION_FIREWALL_AUDITED"
	StatusPuncturedSupportAudited       = "PASS_PUNCTURED_LEPTOCOLOR_SUPPORT_AUDITED"
	StatusExcludedSingletonAudited      = "PASS_EXCLUDED_LEPTON_SINGLETON_AUDITED"
	StatusBMinusLConservationAudited    = "PASS_B_MINUS_L_PUNCTURE_CONSERVATION_AUDITED"
	StatusCompressionCandidateSharpened = "PASS_SOCKET_COMPRESSION_CANDIDATE_SHARPENED_TO_RIGHT_RECTANGLE_PUNCTURE"
	StatusMagnitudeFirewallPreserved    = "PASS_PUNCTURE_COMPRESSION_NOT_TRACE_MAGNITUDE_READOUT"
	StatusAlphaStillSealed              = "PASS_ALPHA_B_REMAINS_SEALED_AFTER_RIGHT_SOCKET_AUDIT"
	StatusOfficialLedgersFrozen         = "PASS_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusR2PlusPlusRetained            = "PASS_R2_PLUS_PLUS_RETAINED_NOT_R3_OR_R4"
	StatusNoObservedDataUsed            = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusFirewallGate840               = "FIREWALL_PRESERVED_GATE840_RIGHT_SOCKET_PUNCTURE_OBSTRUCTION"

	SupportGate839CompressionInherited            = "CONDITIONAL_SUPPORT_GATE839_COMPRESSION_CANDIDATE_INHERITED"
	SupportRightCharacterSplitSeal                = "CONDITIONAL_SUPPORT_RIGHT_SOCKET_CHARACTER_SPLIT_SEAL_CANDIDATE"
	SupportEPlusEMinusAsCharacterProjectors       = "CONDITIONAL_SUPPORT_E_PLUS_E_MINUS_AS_CHARACTER_PROJECTORS_IF_RHO_R_HAS_LAMBDA_BARLAMBDA_SPLIT"
	SupportRankOneSocketsNotArbitraryIfCharacters = "CONDITIONAL_SUPPORT_RANK_ONE_SOCKET_PROJECTORS_NOT_ARBITRARY_IF_CHARACTER_SPLIT_SEALED"
	SupportPuncturedRectangleCandidate            = "CONDITIONAL_SUPPORT_PUNCTURED_RIGHT_LEPTOCOLOR_RECTANGLE_CANDIDATE"
	SupportSelectedRankSeven                      = "CONDITIONAL_SUPPORT_SELECTED_SUPPORT_HAS_RANK_SEVEN_FROM_THREE_PLUS_FOUR"
	SupportExcludedSingletonRankOne               = "CONDITIONAL_SUPPORT_EXCLUDED_SINGLETON_HAS_RANK_ONE"
	SupportBMinusLSelectedPlusOne                 = "CONDITIONAL_SUPPORT_SELECTED_PUNCTURE_HAS_B_MINUS_L_TRACE_PLUS_ONE"
	SupportBMinusLExcludedMinusOne                = "CONDITIONAL_SUPPORT_EXCLUDED_SINGLETON_HAS_B_MINUS_L_TRACE_MINUS_ONE"
	SupportFullRightRectangleNeutral              = "CONDITIONAL_SUPPORT_FULL_RIGHT_RECTANGLE_B_MINUS_L_TRACE_ZERO"
	SupportCompressionShadowIfOrientation         = "CONDITIONAL_SUPPORT_AGGREGATE_SHADOW_IF_ORIENTATION_AND_COMPRESSION_MAP_CERTIFIED"

	FailureRightCharacterSplitSealNotNative = "FAILED_ROUTE_RIGHT_SOCKET_CHARACTER_SPLIT_IS_SEAL_NOT_NATIVE_DERIVATION"
	FailureNoExplicitRhoRMatrixProof        = "FAILED_ROUTE_NO_EXPLICIT_RHO_R_LAMBDA_BARLAMBDA_MATRIX_PROOF_CERTIFIED"
	FailureNoFullRhoFActionLedger           = "FAILED_ROUTE_NO_FULL_RHO_F_ACTION_LEDGER_CERTIFIED"
	FailureDominantRestOrientationMissing   = "FAILED_ROUTE_DOMINANT_REST_ORIENTATION_NOT_CERTIFIED"
	FailureEPlusNotTop                      = "FAILED_ROUTE_E_PLUS_NOT_IDENTIFIED_WITH_TOP_SOCKET"
	FailureEMinusNotRest                    = "FAILED_ROUTE_E_MINUS_NOT_IDENTIFIED_WITH_REST_SOCKET"
	FailureNoDominantColorSocketSelector    = "FAILED_ROUTE_NO_DOMINANT_COLOR_SOCKET_SELECTOR"
	FailureNoRestLeptoColorSocketSelector   = "FAILED_ROUTE_NO_REST_LEPTOCOLOR_SOCKET_SELECTOR"
	FailurePunctureNotCompressionTheorem    = "FAILED_ROUTE_PUNCTURED_RECTANGLE_IS_SUPPORT_ANATOMY_NOT_COMPRESSION_THEOREM"
	FailureNoTypedCompressionMap            = "FAILED_ROUTE_NO_TYPED_SOCKET_PAIR_COMPRESSION_MAP_CERTIFIED"
	FailureExcludedSingletonNotParticle     = "FAILED_ROUTE_EXCLUDED_SINGLETON_NOT_PHYSICAL_PARTICLE_ASSIGNMENT"
	FailureNoRightNeutrinoTheorem           = "FAILED_ROUTE_NO_RIGHT_NEUTRINO_OR_STERILE_PARTICLE_THEOREM"
	FailureNoDFOrHiggsOrientation           = "FAILED_ROUTE_NO_D_F_OR_HIGGS_EDGE_ORIENTATION_SELECTOR_CERTIFIED"
	FailureNoBoundaryRestOrientation        = "FAILED_ROUTE_NO_BOUNDARY_REST_PRESSURE_ORIENTATION_SELECTOR_CERTIFIED"
	FailureNoAlphaDerivation                = "FAILED_ROUTE_PUNCTURED_COMPRESSION_DOES_NOT_DERIVE_ALPHA_B"
	FailureAlphaStillSealed                 = "FAILED_ROUTE_ALPHA_B_STILL_SEALED_NOT_NATIVE_SOURCE"
	FailureNoTraceMagnitudeReadout          = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureCompressionNotYukawaMagnitude    = "FAILED_ROUTE_PUNCTURE_COMPRESSION_NOT_YUKAWA_MAGNITUDE_SOURCE"
	FailureNoNEffUpdate                     = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaUpdate                  = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoObservedYukawaFit              = "FAILED_ROUTE_NO_OBSERVED_YUKAWA_FITTING_ALLOWED"
	FailureNoThreeGenerationTheorem         = "FAILED_ROUTE_NO_THREE_GENERATION_THEOREM"
	FailureNotR3                            = "FAILED_ROUTE_R2_PLUS_PLUS_NOT_R3"
	FailureNotR4                            = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
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
	Gate839Inherited, Gate838BodyInherited    bool
	Supports, Failures                        []string
}

type RightCharacterSplit struct {
	RightSocketExpression                          string
	RhoRForm                                       string
	CharacterProjectorExpressions                  []string
	RightSocketRank, CharacterSocketRank           int
	CharacterPairOrthogonal, CharacterPairComplete bool
	CharacterSplitSealAudited                      bool
	CharacterProjectorsSourceTypedBySeal           bool
	NativeDerivationCertified                      bool
	ExplicitRhoRMatrixCertified                    bool
	FullRhoFActionLedgerCertified                  bool
	UnorderedPairCertified                         bool
	DominantRestOrientationCertified               bool
	Supports, Failures                             []string
}

type WCarrier struct {
	P1Rank, P3Rank, Dim                     int
	P1P3Orthogonal, P1PlusP3CompletesW      bool
	BMinusLLeptonWeight, BMinusLColorWeight float64
	BMinusLTraceZeroOnW                     bool
	Supports, Failures                      []string
}

type PuncturedSupport struct {
	FullRectangleExpression, SelectedExpression, ExcludedExpression string
	TopExpression, RestExpression                                   string
	FullRank, TopRank, RestRank, SelectedRank, ExcludedRank         int
	RanksCloseRightRectangle                                        bool
	UsesRightCharacterSockets                                       bool
	OrientationCertified                                            bool
	CompressionMapCertified                                         bool
	IsTheorem                                                       bool
	Supports, Failures                                              []string
}

type BMinusLBalance struct {
	TopTrace, RestTrace, SelectedTrace, ExcludedTrace, FullRectangleTrace float64
	TopTraceMatches, RestTraceMatches, SelectedExcludedCancel             bool
	FullRectangleNeutral, PunctureConservationPattern                     bool
	Supports, Failures                                                    []string
}

type ShadowOperator struct {
	CandidateTotalOperator                  string
	TopBlockLocatedInRightColorSocket       bool
	RestBlockLocatedInRightLeptoColorSocket bool
	OrientationNeeded                       bool
	CompressionMapCertified                 bool
	AlphaDerivedByCompression               bool
	TraceMagnitudeReadoutCertified          bool
	R3, R4                                  bool
	Supports, Failures                      []string
}

type Impact struct {
	FineSocketProblemPartiallyResolvedBySeal bool
	OrientationStillMissing                  bool
	CompressionMapStillMissing               bool
	BMinusLConservationPatternFound          bool
	CanPromoteToR3, CanPromoteToR4           bool
	CanUpdateNEff, CanUpdateCYukawa          bool
	CanUpdateCHiggs                          bool
	Classification                           string
	Supports, Failures                       []string
}

type Firewalls struct {
	Enforced                      bool
	CharacterSplitSealNotNative   bool
	NoExplicitRhoRProof           bool
	NoFullRhoFActionLedger        bool
	OrientationMissing            bool
	NoDominantSelector            bool
	NoRestSelector                bool
	PunctureNotCompressionTheorem bool
	NoCompressionMap              bool
	ExcludedSingletonNotParticle  bool
	NoRightNeutrinoTheorem        bool
	NoDFOrHiggsOrientation        bool
	NoBoundaryRestOrientation     bool
	NoAlphaDerivation             bool
	AlphaSealed                   bool
	NoTraceMagnitudeReadout       bool
	CompressionNotYukawaMagnitude bool
	NoNEffUpdate                  bool
	NoCYukawaUpdate               bool
	NoObservedYukawaFit           bool
	NoThreeGeneration             bool
	NotR3, NotR4                  bool
	Verdict                       string
}

type Analysis struct {
	AuditID    string
	Ledger     Ledger
	Body       FiniteBody
	RightSplit RightCharacterSplit
	W          WCarrier
	Puncture   PuncturedSupport
	BMinusL    BMinusLBalance
	Shadow     ShadowOperator
	Impact     Impact
	Firewalls  Firewalls
	Truth      string
	Final      string
}

func BuildDefault() (Analysis, error) {
	wTrace := float64(LeptonBlockDim)*BMinusLLeptonWeight + float64(ColorBlockDim)*BMinusLColorWeight
	if math.Abs(wTrace) > 1e-15 {
		return Analysis{}, fmt.Errorf("B-L trace on W is not zero: %.17g", wTrace)
	}
	fullRightTrace := float64(RightSlotDim) * wTrace
	if math.Abs(fullRightTrace) > 1e-15 {
		return Analysis{}, fmt.Errorf("B-L trace on right rectangle is not zero: %.17g", fullRightTrace)
	}
	if HPartDim != EDim*WDim || HFDim != 2*HPartDim || RightRectangleRank != 8 || SelectedRank != 7 || SelectedRank+ExcludedRank != RightRectangleRank {
		return Analysis{}, fmt.Errorf("dimension invariant failed")
	}

	ledger := Ledger{S: SBoundary, AlphaB: AlphaB, OperatorNEff: OperatorNEff, OfficialNEff: OfficialNEff, OfficialCYukawa: OfficialCYukawa, OfficialCHiggs: OfficialCHiggs, OfficialFrozen: true, R2PlusPlus: true, R3: false, R4: false, AlphaIsNative: false}
	body := FiniteBody{EExpression: "E=C_R^2 plus C_L^2", WExpression: "W=C_lepton plus C_color^3", HPartExpression: "H_part=E tensor W", EDim: EDim, WDim: WDim, HPartDim: HPartDim, HFDim: HFDim, RightSlotDim: RightSlotDim, LeftSlotDim: LeftSlotDim, Gate839Inherited: true, Gate838BodyInherited: true, Supports: []string{SupportGate839CompressionInherited}}
	rightSplit := RightCharacterSplit{
		RightSocketExpression: "C_R^2=e_+ plus e_-", RhoRForm: "rho_R(lambda)=diag(lambda, conjugate(lambda)) [sealed schematic]",
		CharacterProjectorExpressions: []string{"e_+", "e_-"}, RightSocketRank: RightSlotDim, CharacterSocketRank: CharacterSocketRank,
		CharacterPairOrthogonal: true, CharacterPairComplete: true, CharacterSplitSealAudited: true, CharacterProjectorsSourceTypedBySeal: true,
		NativeDerivationCertified: false, ExplicitRhoRMatrixCertified: false, FullRhoFActionLedgerCertified: false,
		UnorderedPairCertified: true, DominantRestOrientationCertified: false,
		Supports: []string{SupportRightCharacterSplitSeal, SupportEPlusEMinusAsCharacterProjectors, SupportRankOneSocketsNotArbitraryIfCharacters},
		Failures: []string{FailureRightCharacterSplitSealNotNative, FailureNoExplicitRhoRMatrixProof, FailureNoFullRhoFActionLedger, FailureDominantRestOrientationMissing, FailureEPlusNotTop, FailureEMinusNotRest},
	}
	w := WCarrier{P1Rank: LeptonBlockDim, P3Rank: ColorBlockDim, Dim: WDim, P1P3Orthogonal: true, P1PlusP3CompletesW: true, BMinusLLeptonWeight: BMinusLLeptonWeight, BMinusLColorWeight: BMinusLColorWeight, BMinusLTraceZeroOnW: true, Supports: []string{SupportFullRightRectangleNeutral}}
	puncture := PuncturedSupport{
		FullRectangleExpression: "C_R^2 tensor W", SelectedExpression: "(e_+ tensor P_3) plus (e_- tensor W)", ExcludedExpression: "e_+ tensor P_1",
		TopExpression: "e_+ tensor P_3", RestExpression: "e_- tensor W", FullRank: RightRectangleRank, TopRank: TopColorRank, RestRank: RestQuartetRank,
		SelectedRank: SelectedRank, ExcludedRank: ExcludedRank, RanksCloseRightRectangle: true, UsesRightCharacterSockets: true,
		OrientationCertified: false, CompressionMapCertified: false, IsTheorem: false,
		Supports: []string{SupportPuncturedRectangleCandidate, SupportSelectedRankSeven, SupportExcludedSingletonRankOne, SupportCompressionShadowIfOrientation},
		Failures: []string{FailureDominantRestOrientationMissing, FailureNoDominantColorSocketSelector, FailureNoRestLeptoColorSocketSelector, FailurePunctureNotCompressionTheorem, FailureNoTypedCompressionMap, FailureExcludedSingletonNotParticle, FailureNoRightNeutrinoTheorem},
	}
	bminusl := BMinusLBalance{
		TopTrace:           float64(ColorBlockDim) * BMinusLColorWeight,
		RestTrace:          float64(LeptonBlockDim)*BMinusLLeptonWeight + float64(ColorBlockDim)*BMinusLColorWeight,
		SelectedTrace:      float64(ColorBlockDim)*BMinusLColorWeight + (float64(LeptonBlockDim)*BMinusLLeptonWeight + float64(ColorBlockDim)*BMinusLColorWeight),
		ExcludedTrace:      float64(LeptonBlockDim) * BMinusLLeptonWeight,
		FullRectangleTrace: fullRightTrace,
		TopTraceMatches:    true, RestTraceMatches: true, SelectedExcludedCancel: true, FullRectangleNeutral: true, PunctureConservationPattern: true,
		Supports: []string{SupportBMinusLSelectedPlusOne, SupportBMinusLExcludedMinusOne, SupportFullRightRectangleNeutral},
	}
	if math.Abs(bminusl.TopTrace-1) > 1e-15 || math.Abs(bminusl.RestTrace) > 1e-15 || math.Abs(bminusl.SelectedTrace-1) > 1e-15 || math.Abs(bminusl.ExcludedTrace+1) > 1e-15 || math.Abs(bminusl.SelectedTrace+bminusl.ExcludedTrace) > 1e-15 {
		return Analysis{}, fmt.Errorf("B-L puncture balance failed")
	}
	shadow := ShadowOperator{CandidateTotalOperator: "I_{e_+ tensor P_3} plus [alpha_B P_3 - 3 alpha_B^2(B-L)] on e_- tensor W", TopBlockLocatedInRightColorSocket: true, RestBlockLocatedInRightLeptoColorSocket: true, OrientationNeeded: true, CompressionMapCertified: false, AlphaDerivedByCompression: false, TraceMagnitudeReadoutCertified: false, R3: false, R4: false, Supports: []string{SupportCompressionShadowIfOrientation}, Failures: []string{FailureNoDFOrHiggsOrientation, FailureNoBoundaryRestOrientation, FailureNoAlphaDerivation, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureCompressionNotYukawaMagnitude, FailureNotR3, FailureNotR4}}
	impact := Impact{FineSocketProblemPartiallyResolvedBySeal: true, OrientationStillMissing: true, CompressionMapStillMissing: true, BMinusLConservationPatternFound: true, CanPromoteToR3: false, CanPromoteToR4: false, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, Classification: "right-character socket seal plus punctured lepto-color support anatomy; orientation and compression map missing; R2++ retained", Supports: []string{SupportRightCharacterSplitSeal, SupportPuncturedRectangleCandidate, SupportBMinusLSelectedPlusOne}, Failures: []string{FailureDominantRestOrientationMissing, FailureNoTypedCompressionMap, FailureNoTraceMagnitudeReadout, FailureAlphaStillSealed, FailureNoNEffUpdate, FailureNoCYukawaUpdate, FailureNoObservedYukawaFit, FailureNoThreeGenerationTheorem, FailureNotR3, FailureNotR4}}
	firewalls := Firewalls{Enforced: true, CharacterSplitSealNotNative: true, NoExplicitRhoRProof: true, NoFullRhoFActionLedger: true, OrientationMissing: true, NoDominantSelector: true, NoRestSelector: true, PunctureNotCompressionTheorem: true, NoCompressionMap: true, ExcludedSingletonNotParticle: true, NoRightNeutrinoTheorem: true, NoDFOrHiggsOrientation: true, NoBoundaryRestOrientation: true, NoAlphaDerivation: true, AlphaSealed: true, NoTraceMagnitudeReadout: true, CompressionNotYukawaMagnitude: true, NoNEffUpdate: true, NoCYukawaUpdate: true, NoObservedYukawaFit: true, NoThreeGeneration: true, NotR3: true, NotR4: true, Verdict: StatusFirewallGate840}
	return Analysis{AuditID: AuditID, Ledger: ledger, Body: body, RightSplit: rightSplit, W: w, Puncture: puncture, BMinusL: bminusl, Shadow: shadow, Impact: impact, Firewalls: firewalls, Truth: "Gate 840 source-types the missing rank-one right socket pair by a sealed lambda/conjugate(lambda) character split and exposes the punctured right lepto-color rectangle.", Final: "Verdict: conditional right-character socket seal and B-L puncture conservation support; no dominant/rest orientation theorem, no typed compression map, no alpha derivation, no trace-magnitude readout, no R3/R4 promotion."}, nil
}

func Statuses() []string {
	return []string{StatusGate839Inherited, StatusRightRectangleAudited, StatusRightCharacterSplitAudited, StatusRankOneCharacterSockets, StatusOrientationFirewallAudited, StatusPuncturedSupportAudited, StatusExcludedSingletonAudited, StatusBMinusLConservationAudited, StatusCompressionCandidateSharpened, StatusMagnitudeFirewallPreserved, StatusAlphaStillSealed, StatusOfficialLedgersFrozen, StatusR2PlusPlusRetained, StatusNoObservedDataUsed, StatusFirewallGate840, SupportGate839CompressionInherited, SupportRightCharacterSplitSeal, SupportEPlusEMinusAsCharacterProjectors, SupportRankOneSocketsNotArbitraryIfCharacters, SupportPuncturedRectangleCandidate, SupportSelectedRankSeven, SupportExcludedSingletonRankOne, SupportBMinusLSelectedPlusOne, SupportBMinusLExcludedMinusOne, SupportFullRightRectangleNeutral, SupportCompressionShadowIfOrientation, FailureRightCharacterSplitSealNotNative, FailureNoExplicitRhoRMatrixProof, FailureNoFullRhoFActionLedger, FailureDominantRestOrientationMissing, FailureEPlusNotTop, FailureEMinusNotRest, FailureNoDominantColorSocketSelector, FailureNoRestLeptoColorSocketSelector, FailurePunctureNotCompressionTheorem, FailureNoTypedCompressionMap, FailureExcludedSingletonNotParticle, FailureNoRightNeutrinoTheorem, FailureNoDFOrHiggsOrientation, FailureNoBoundaryRestOrientation, FailureNoAlphaDerivation, FailureAlphaStillSealed, FailureNoTraceMagnitudeReadout, FailureCompressionNotYukawaMagnitude, FailureNoNEffUpdate, FailureNoCYukawaUpdate, FailureNoObservedYukawaFit, FailureNoThreeGenerationTheorem, FailureNotR3, FailureNotR4}
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("ledger: s=%.16g alpha_B=%.16g operator_N_eff=%.16g official_N_eff=%.16g frozen=%t R2++=%t R3=%t R4=%t alpha_native=%t", l.S, l.AlphaB, l.OperatorNEff, l.OfficialNEff, l.OfficialFrozen, l.R2PlusPlus, l.R3, l.R4, l.AlphaIsNative)
}

func FormatBody(b FiniteBody) string {
	return fmt.Sprintf("body: %s, %s, %s dim(E)=%d dim(W)=%d dim(H_part)=%d dim(H_F)=%d right=%d left=%d inherited839=%t", b.EExpression, b.WExpression, b.HPartExpression, b.EDim, b.WDim, b.HPartDim, b.HFDim, b.RightSlotDim, b.LeftSlotDim, b.Gate839Inherited)
}

func FormatRightSplit(r RightCharacterSplit) string {
	return fmt.Sprintf("right-split: %s rho=%s chars=[%s] rank=%d socket_rank=%d orthogonal=%t complete=%t seal=%t source_typed=%t native=%t explicit_rho=%t orientation=%t", r.RightSocketExpression, r.RhoRForm, strings.Join(r.CharacterProjectorExpressions, ","), r.RightSocketRank, r.CharacterSocketRank, r.CharacterPairOrthogonal, r.CharacterPairComplete, r.CharacterSplitSealAudited, r.CharacterProjectorsSourceTypedBySeal, r.NativeDerivationCertified, r.ExplicitRhoRMatrixCertified, r.DominantRestOrientationCertified)
}

func FormatW(w WCarrier) string {
	return fmt.Sprintf("W: rank(P1)=%d rank(P3)=%d dim=%d B-L=(%.12g,%.12g) trace_zero=%t", w.P1Rank, w.P3Rank, w.Dim, w.BMinusLLeptonWeight, w.BMinusLColorWeight, w.BMinusLTraceZeroOnW)
}

func FormatPuncture(p PuncturedSupport) string {
	return fmt.Sprintf("puncture: full=%s selected=%s excluded=%s ranks top/rest/selected/excluded/full=%d/%d/%d/%d/%d orientation=%t compression_map=%t theorem=%t", p.FullRectangleExpression, p.SelectedExpression, p.ExcludedExpression, p.TopRank, p.RestRank, p.SelectedRank, p.ExcludedRank, p.FullRank, p.OrientationCertified, p.CompressionMapCertified, p.IsTheorem)
}

func FormatBMinusL(b BMinusLBalance) string {
	return fmt.Sprintf("B-L balance: top=%.12g rest=%.12g selected=%.12g excluded=%.12g full=%.12g selected_plus_excluded_cancel=%t full_neutral=%t", b.TopTrace, b.RestTrace, b.SelectedTrace, b.ExcludedTrace, b.FullRectangleTrace, b.SelectedExcludedCancel, b.FullRectangleNeutral)
}

func FormatShadow(s ShadowOperator) string {
	return fmt.Sprintf("shadow: %s top_located=%t rest_located=%t orientation_needed=%t compression_map=%t alpha_derived=%t trace_readout=%t R3=%t R4=%t", s.CandidateTotalOperator, s.TopBlockLocatedInRightColorSocket, s.RestBlockLocatedInRightLeptoColorSocket, s.OrientationNeeded, s.CompressionMapCertified, s.AlphaDerivedByCompression, s.TraceMagnitudeReadoutCertified, s.R3, s.R4)
}

func FormatImpact(i Impact) string {
	return fmt.Sprintf("impact: fine_socket_partially_resolved=%t orientation_missing=%t compression_missing=%t B-L_pattern=%t promote_R3=%t promote_R4=%t update_N_eff=%t update_C_Yukawa=%t update_C_Higgs=%t classification=%q", i.FineSocketProblemPartiallyResolvedBySeal, i.OrientationStillMissing, i.CompressionMapStillMissing, i.BMinusLConservationPatternFound, i.CanPromoteToR3, i.CanPromoteToR4, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.Classification)
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
