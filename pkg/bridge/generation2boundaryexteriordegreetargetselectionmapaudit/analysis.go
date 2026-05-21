// Package generation2boundaryexteriordegreetargetselectionmapaudit implements
// Gate 871: BoundaryExterior Degree-Target Selection Map Audit.
//
// Gate 871 follows Gate 870's reduced boundary-pair exterior response audit. Gate
// 870 conditionally explained the alpha_B power shape through the reduced
// response
//
//	R_B(s)=(1+s b_1)(1+s b_2)-1=s(b_1+b_2)+s^2(b_1 wedge b_2),
//
// suppressing the zero-order term and truncating after degree two because
// Lambda^3 B_2=0. Gate 871 audits the remaining wound: the degree-target
// selection maps Lambda^1 B_2 -> Pi_top and Lambda^2 B_2 -> H_R^min, plus the
// required cross-lane exclusions. The result is intentionally conservative:
// target assignments are coherent candidates, but no native target-selection
// theorem or alpha_B theorem is certified.
package generation2boundaryexteriordegreetargetselectionmapaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE871-BOUNDARY-EXTERIOR-DEGREE-TARGET-SELECTION-MAP-AUDIT"

	AlphaB    = 0.0003878958469680527
	SBoundary = 0.0012924448188162962

	BoundaryPairDim = 2
	Lambda1B2Dim    = 2
	Lambda2B2Dim    = 1
	Lambda3B2Dim    = 0

	PiTopRank     = 3
	HRambientRank = 8
	H10Dim        = HRambientRank + BoundaryPairDim
	HRminRank     = 7
	Lambda4V8Rank = 70
	H72Dim        = Lambda4V8Rank + BoundaryPairDim

	Classification = "BOUNDARY_EXTERIOR_DEGREE_TARGET_SELECTION_MAP_OBSTRUCTION"
	R2Status       = "R2+++++_BOUNDARY_EXTERIOR_DEGREE_TARGET_SELECTION_OBSTRUCTION"

	StatusGate870Inherited       = "PASS_GATE870_REDUCED_BOUNDARY_EXTERIOR_RESPONSE_INHERITED"
	StatusDegreeTargetWound      = "PASS_DEGREE_TARGET_SELECTION_WOUND_ISOLATED"
	StatusDegreeOneTargetAudited = "PASS_LAMBDA1_B2_TO_PI_TOP_TARGET_AUDITED"
	StatusDegreeTwoTargetAudited = "PASS_LAMBDA2_B2_TO_H_R_MIN_TARGET_AUDITED"
	StatusCrossLaneAudited       = "PASS_CROSS_LANE_EXCLUSION_AUDITED"
	StatusChambersAudited        = "PASS_H10_AND_H72_RESPONSE_CHAMBERS_AUDITED"
	StatusAlphaReconstructed     = "PASS_ALPHA_B_RECONSTRUCTED_FROM_CANDIDATE_TARGETS"
	StatusFirewallEnforced       = "PASS_DEGREE_TARGET_SELECTION_FIREWALL_ENFORCED"
	StatusLedgerFrozen           = "PASS_OFFICIAL_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusNoObservedDataUsed     = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusNextWound              = "PASS_NEXT_WOUND_IDENTIFIED_AS_NATIVE_DEGREE_TARGET_SELECTION_MAP"
	StatusFirewallVerdict        = "FIREWALL_PRESERVED_GATE871_DEGREE_TARGET_SELECTION_NOT_ALPHA_THEOREM"

	SupportReducedB2ShapeInherited             = "CONDITIONAL_SUPPORT_REDUCED_B2_RESPONSE_GIVES_CORRECT_ALPHA_POWER_SHAPE"
	SupportDegreeOneDominantSocketCandidate    = "CONDITIONAL_SUPPORT_DEGREE_ONE_TARGETS_DOMINANT_SOCKET_CANDIDATE"
	SupportDegreeTwoActiveRightDomainCandidate = "CONDITIONAL_SUPPORT_DEGREE_TWO_TARGETS_ACTIVE_PUNCTURED_RIGHT_DOMAIN_CANDIDATE"
	SupportDegreeOneSingleBoundaryResponse     = "CONDITIONAL_SUPPORT_DEGREE_ONE_AS_SINGLE_BOUNDARY_DOMINANT_SOCKET_RESPONSE"
	SupportDegreeTwoFullBoundaryPairResponse   = "CONDITIONAL_SUPPORT_DEGREE_TWO_AS_FULL_BOUNDARY_PAIR_ACTIVE_DOMAIN_RESPONSE"
	SupportH10H72TypedResponseChambers         = "CONDITIONAL_SUPPORT_H10_AND_H72_ARE_TYPED_RESPONSE_CHAMBERS"
	SupportAlphaPowerShapeStillCoherent        = "CONDITIONAL_SUPPORT_ALPHA_B_REDUCED_EXTERIOR_POWER_SHAPE_REMAINS_COHERENT"
	SupportDegreeTargetWoundSharpened          = "CONDITIONAL_SUPPORT_ALPHA_B_WOUND_REDUCED_TO_DEGREE_TARGET_SELECTION"

	FailureNoNativeDegreeTargetSelectionMap  = "FAILED_ROUTE_NO_NATIVE_DEGREE_TARGET_SELECTION_MAP"
	FailureNoNativeLambda1ToPiTopMap         = "FAILED_ROUTE_NO_NATIVE_LAMBDA1B2_TO_PI_TOP_MAP"
	FailureNoNativeLambda2ToHRMinMap         = "FAILED_ROUTE_NO_NATIVE_LAMBDA2B2_TO_H_R_MIN_MAP"
	FailureNoCrossLaneExclusionTheorem       = "FAILED_ROUTE_NO_CROSS_LANE_EXCLUSION_THEOREM"
	FailureNoLinearHRMinExclusion            = "FAILED_ROUTE_NO_NATIVE_EXCLUSION_OF_LINEAR_H_R_MIN_TERM"
	FailureNoQuadraticPiTopExclusion         = "FAILED_ROUTE_NO_NATIVE_EXCLUSION_OF_QUADRATIC_PI_TOP_TERM"
	FailureResponseChamberTypingNotTheorem   = "FAILED_ROUTE_RESPONSE_CHAMBER_TYPING_NOT_TARGET_SELECTION_THEOREM"
	FailureNoNativeReducedExteriorFunctional = "FAILED_ROUTE_NO_NATIVE_REDUCED_BOUNDARY_EXTERIOR_RESPONSE_FUNCTIONAL"
	FailureNoSharedDegreeTargetFunctor       = "FAILED_ROUTE_NO_SHARED_DEGREE_TARGET_SELECTION_FUNCTOR_CERTIFIED"
	FailureAlphaStillSealed                  = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNoNativeAlphaSource               = "FAILED_ROUTE_NO_NATIVE_ALPHA_B_SOURCE"
	FailureSocketRankRatiosNotActivation     = "FAILED_ROUTE_SOCKET_RANK_RATIOS_NOT_ACTIVATION_THEOREM"
	FailureNoBoundaryAlphaTransport          = "FAILED_ROUTE_NO_BOUNDARY_ALPHA_DOMAIN_TRANSPORT_THEOREM"
	FailureNoOfficialNEffUpdate              = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate             = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNumericalYukawa                 = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureNoSectorTraceMagnitude            = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoFullUnbrokenAF                  = "FAILED_ROUTE_NO_FULL_UNBROKEN_A_F_THEOREM"
	FailureAForientNotFullAF                 = "FAILED_ROUTE_A_F_ORIENT_NOT_FULL_A_F"
	FailureNoNativeFiniteTriple              = "FAILED_ROUTE_NO_NATIVE_FINITE_TRIPLE_THEOREM"
	FailureNotR3                             = "FAILED_ROUTE_NOT_R3_BOUNDARY_EXTERIOR_DEGREE_TARGET_SELECTION_OBSTRUCTION"
	FailureNotR4                             = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	AlphaB, SBoundary float64
	AlphaNative       bool
	OfficialFrozen    bool
	R3, R4            bool
}

type ReducedResponseInherited struct {
	Expression                           string
	ZeroOrderSuppressed, HigherTruncated bool
	DegreeOnePresent, DegreeTwoPresent   bool
	NativeFunctionalCertified            bool
	Supports, Failures                   []string
}

type TargetMap struct {
	Degree             int
	Source             string
	CandidateTarget    string
	Chamber            string
	TargetRank         int
	ChamberDim         int
	Contribution       float64
	MapCertified       bool
	Interpretation     string
	Supports, Failures []string
}

type CrossLaneExclusion struct {
	LinearToHRMinExcludedCandidate, QuadraticToPiTopExcludedCandidate   bool
	CrossLaneExclusionTheorem                                           bool
	WouldAddLinearHRMinContribution, WouldAddQuadraticPiTopContribution float64
	Supports, Failures                                                  []string
}

type ChamberAudit struct {
	H10Typed, H72Typed             bool
	ResponseChamberTypingTheorem   bool
	H10Description, H72Description string
	Supports, Failures             []string
}

type AlphaCandidate struct {
	Expression                      string
	LinearContribution              float64
	QuadraticContribution           float64
	CrossLinearHRMinContribution    float64
	CrossQuadraticPiTopContribution float64
	ReconstructedAlpha              float64
	ShapeCoherent, Native           bool
	Supports, Failures              []string
}

type Obstruction struct {
	ReducedExteriorShapeInherited  bool
	DegreeTargetSelectionCertified bool
	CrossLaneExclusionCertified    bool
	AlphaNative                    bool
	RemainingWound, NextGate       string
	Supports, Failures             []string
}

type R3Assessment struct {
	PostOrientationFiniteTripleSeal bool
	YDagYReadoutCarrierReady        bool
	SocketMagnitudeSourceTyped      bool
	SocketRankAlphaSourceTyped      bool
	ReducedB2ResponseShapeTyped     bool
	DegreeTargetSelectionCertified  bool
	AlphaNative                     bool
	SectorTraceMagnitudeReadout     bool
	EligibleForR3, EligibleForR4    bool
	Supports, Failures              []string
}

type Impact struct {
	Classification, Status                                                           string
	ReducedShapeSolved, DegreeTargetsSolved, CrossLanesSolved, AlphaNative           bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4 bool
}

type Firewalls struct {
	Enforced                                                                                        bool
	NoNativeDegreeTargetSelectionMap, NoNativeLambda1ToPiTopMap, NoNativeLambda2ToHRMinMap          bool
	NoCrossLaneExclusionTheorem, NoLinearHRMinExclusion, NoQuadraticPiTopExclusion                  bool
	ResponseChamberTypingNotTheorem, NoNativeReducedExteriorFunctional, NoSharedDegreeTargetFunctor bool
	AlphaStillSealed, NoNativeAlphaSource, SocketRankRatiosNotActivation, NoBoundaryAlphaTransport  bool
	NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate, NoNumericalYukawa, NoSectorTraceMagnitude          bool
	NoFullUnbrokenAF, AForientNotFullAF, NoNativeFiniteTriple, NotR3, NotR4                         bool
	Verdict                                                                                         string
}

type Audit struct {
	ID           string
	Ledger       Ledger
	Response     ReducedResponseInherited
	DegreeOne    TargetMap
	DegreeTwo    TargetMap
	CrossLane    CrossLaneExclusion
	Chambers     ChamberAudit
	Candidate    AlphaCandidate
	Obstruction  Obstruction
	R3           R3Assessment
	Impact       Impact
	Firewalls    Firewalls
	Truth, Final string
}

func BuildDefault() (Audit, error) {
	if H10Dim != 10 || H72Dim != 72 || PiTopRank != 3 || HRminRank != 7 || Lambda3B2Dim != 0 {
		return Audit{}, fmt.Errorf("unexpected Gate 871 dimension ledger")
	}
	linear := float64(PiTopRank) / float64(H10Dim) * SBoundary
	quadratic := float64(HRminRank) / float64(H72Dim) * SBoundary * SBoundary
	crossLinearHRMin := float64(HRminRank) / float64(H72Dim) * SBoundary
	crossQuadraticPiTop := float64(PiTopRank) / float64(H10Dim) * SBoundary * SBoundary
	alpha := linear + quadratic
	if !near(alpha, AlphaB) {
		return Audit{}, fmt.Errorf("alpha reconstruction mismatch got %.18g want %.18g", alpha, AlphaB)
	}
	return Audit{
		ID:     AuditID,
		Ledger: Ledger{AlphaB: AlphaB, SBoundary: SBoundary, AlphaNative: false, OfficialFrozen: true, R3: false, R4: false},
		Response: ReducedResponseInherited{
			Expression:          "R_B(s)=(1+s b1)(1+s b2)-1=s(b1+b2)+s^2(b1 wedge b2)",
			ZeroOrderSuppressed: true, HigherTruncated: true, DegreeOnePresent: true, DegreeTwoPresent: true, NativeFunctionalCertified: false,
			Supports: []string{StatusGate870Inherited, SupportReducedB2ShapeInherited, SupportAlphaPowerShapeStillCoherent},
			Failures: []string{FailureNoNativeReducedExteriorFunctional},
		},
		DegreeOne: TargetMap{
			Degree: 1, Source: "Lambda^1 B_2 reduced exterior term s(b1+b2)", CandidateTarget: "Pi_top=e_+ tensor P_3", Chamber: "H10=H_R^ambient plus B_2", TargetRank: PiTopRank, ChamberDim: H10Dim, Contribution: linear, MapCertified: false,
			Interpretation: "single-boundary dominant socket response candidate",
			Supports:       []string{StatusDegreeOneTargetAudited, SupportDegreeOneDominantSocketCandidate, SupportDegreeOneSingleBoundaryResponse},
			Failures:       []string{FailureNoNativeDegreeTargetSelectionMap, FailureNoNativeLambda1ToPiTopMap},
		},
		DegreeTwo: TargetMap{
			Degree: 2, Source: "Lambda^2 B_2 reduced exterior term s^2(b1 wedge b2)", CandidateTarget: "H_R^min active punctured right domain", Chamber: "H72=Lambda^4 V_8 plus B_2", TargetRank: HRminRank, ChamberDim: H72Dim, Contribution: quadratic, MapCertified: false,
			Interpretation: "full boundary-pair active-domain response candidate",
			Supports:       []string{StatusDegreeTwoTargetAudited, SupportDegreeTwoActiveRightDomainCandidate, SupportDegreeTwoFullBoundaryPairResponse},
			Failures:       []string{FailureNoNativeDegreeTargetSelectionMap, FailureNoNativeLambda2ToHRMinMap},
		},
		CrossLane: CrossLaneExclusion{
			LinearToHRMinExcludedCandidate: true, QuadraticToPiTopExcludedCandidate: true, CrossLaneExclusionTheorem: false,
			WouldAddLinearHRMinContribution: crossLinearHRMin, WouldAddQuadraticPiTopContribution: crossQuadraticPiTop,
			Supports: []string{StatusCrossLaneAudited},
			Failures: []string{FailureNoCrossLaneExclusionTheorem, FailureNoLinearHRMinExclusion, FailureNoQuadraticPiTopExclusion},
		},
		Chambers: ChamberAudit{
			H10Typed: true, H72Typed: true, ResponseChamberTypingTheorem: false,
			H10Description: "H10 is the ambient right lepto-color rectangle plus boundary pair: 8+2", H72Description: "H72 is the Lambda^4 V8 chamber plus boundary pair: 70+2",
			Supports: []string{StatusChambersAudited, SupportH10H72TypedResponseChambers},
			Failures: []string{FailureResponseChamberTypingNotTheorem},
		},
		Candidate: AlphaCandidate{
			Expression:         "alpha_B=[rank(Pi_top)/10]s+[rank(H_R^min)/72]s^2 with degree targets Lambda1B2->Pi_top and Lambda2B2->H_R^min",
			LinearContribution: linear, QuadraticContribution: quadratic, CrossLinearHRMinContribution: crossLinearHRMin, CrossQuadraticPiTopContribution: crossQuadraticPiTop,
			ReconstructedAlpha: alpha, ShapeCoherent: true, Native: false,
			Supports: []string{StatusAlphaReconstructed, SupportReducedB2ShapeInherited, SupportDegreeTargetWoundSharpened},
			Failures: []string{FailureNoNativeDegreeTargetSelectionMap, FailureNoNativeLambda1ToPiTopMap, FailureNoNativeLambda2ToHRMinMap, FailureAlphaStillSealed},
		},
		Obstruction: Obstruction{
			ReducedExteriorShapeInherited: true, DegreeTargetSelectionCertified: false, CrossLaneExclusionCertified: false, AlphaNative: false,
			RemainingWound: "native degree-target selection map Lambda^1 B2 -> Pi_top and Lambda^2 B2 -> H_R^min, including cross-lane exclusions",
			NextGate:       "Gate 872 — BoundaryExterior Target-Selection Seal / Theorem Eligibility Audit",
			Supports:       []string{StatusDegreeTargetWound, SupportDegreeTargetWoundSharpened, SupportReducedB2ShapeInherited},
			Failures:       []string{FailureNoNativeDegreeTargetSelectionMap, FailureNoCrossLaneExclusionTheorem, FailureAlphaStillSealed},
		},
		R3: R3Assessment{
			PostOrientationFiniteTripleSeal: true, YDagYReadoutCarrierReady: true, SocketMagnitudeSourceTyped: true, SocketRankAlphaSourceTyped: true, ReducedB2ResponseShapeTyped: true,
			DegreeTargetSelectionCertified: false, AlphaNative: false, SectorTraceMagnitudeReadout: false, EligibleForR3: false, EligibleForR4: false,
			Supports: []string{SupportReducedB2ShapeInherited, SupportDegreeOneDominantSocketCandidate, SupportDegreeTwoActiveRightDomainCandidate},
			Failures: []string{FailureNoNativeDegreeTargetSelectionMap, FailureAlphaStillSealed, FailureNoSectorTraceMagnitude, FailureNotR3, FailureNotR4},
		},
		Impact: Impact{Classification: Classification, Status: R2Status, ReducedShapeSolved: true, DegreeTargetsSolved: false, CrossLanesSolved: false, AlphaNative: false, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, CanPromoteToR3: false, CanPromoteToR4: false},
		Firewalls: Firewalls{
			Enforced: true, NoNativeDegreeTargetSelectionMap: true, NoNativeLambda1ToPiTopMap: true, NoNativeLambda2ToHRMinMap: true, NoCrossLaneExclusionTheorem: true, NoLinearHRMinExclusion: true, NoQuadraticPiTopExclusion: true,
			ResponseChamberTypingNotTheorem: true, NoNativeReducedExteriorFunctional: true, NoSharedDegreeTargetFunctor: true, AlphaStillSealed: true, NoNativeAlphaSource: true, SocketRankRatiosNotActivation: true, NoBoundaryAlphaTransport: true,
			NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NoNumericalYukawa: true, NoSectorTraceMagnitude: true, NoFullUnbrokenAF: true, AForientNotFullAF: true, NoNativeFiniteTriple: true, NotR3: true, NotR4: true,
			Verdict: StatusFirewallVerdict,
		},
		Truth: "Gate 871 narrows the alpha_B wound to degree-target selection: reduced B2 exterior response gives the s+s^2 shape, but no native map sends Lambda^1 B2 to Pi_top and Lambda^2 B2 to H_R^min.",
		Final: "VERDICT: CONDITIONAL_SUPPORT_DEGREE_TARGET_CANDIDATES, FAILED_ROUTE_NO_NATIVE_DEGREE_TARGET_SELECTION_MAP, FAILED_ROUTE_ALPHA_B_REMAINS_SEALED, NOT_R3.",
	}, nil
}

func Statuses() []string {
	return []string{
		StatusGate870Inherited, StatusDegreeTargetWound, StatusDegreeOneTargetAudited, StatusDegreeTwoTargetAudited, StatusCrossLaneAudited, StatusChambersAudited, StatusAlphaReconstructed, StatusFirewallEnforced, StatusLedgerFrozen, StatusNoObservedDataUsed, StatusNextWound, StatusFirewallVerdict,
		SupportReducedB2ShapeInherited, SupportDegreeOneDominantSocketCandidate, SupportDegreeTwoActiveRightDomainCandidate, SupportDegreeOneSingleBoundaryResponse, SupportDegreeTwoFullBoundaryPairResponse, SupportH10H72TypedResponseChambers, SupportAlphaPowerShapeStillCoherent, SupportDegreeTargetWoundSharpened,
		FailureNoNativeDegreeTargetSelectionMap, FailureNoNativeLambda1ToPiTopMap, FailureNoNativeLambda2ToHRMinMap, FailureNoCrossLaneExclusionTheorem, FailureNoLinearHRMinExclusion, FailureNoQuadraticPiTopExclusion, FailureResponseChamberTypingNotTheorem, FailureNoNativeReducedExteriorFunctional, FailureNoSharedDegreeTargetFunctor, FailureAlphaStillSealed, FailureNoNativeAlphaSource, FailureSocketRankRatiosNotActivation, FailureNoBoundaryAlphaTransport, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNumericalYukawa, FailureNoSectorTraceMagnitude, FailureNoFullUnbrokenAF, FailureAForientNotFullAF, FailureNoNativeFiniteTriple, FailureNotR3, FailureNotR4,
	}
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("alpha_B=%.18g s=%.18g alphaNative=%t officialFrozen=%t R3=%t R4=%t", l.AlphaB, l.SBoundary, l.AlphaNative, l.OfficialFrozen, l.R3, l.R4)
}
func FormatResponse(r ReducedResponseInherited) string {
	return fmt.Sprintf("response=%s zeroRemoved=%t higherTruncated=%t degree1=%t degree2=%t nativeFunctional=%t supports=%s failures=%s", r.Expression, r.ZeroOrderSuppressed, r.HigherTruncated, r.DegreeOnePresent, r.DegreeTwoPresent, r.NativeFunctionalCertified, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}
func FormatTarget(t TargetMap) string {
	return fmt.Sprintf("degree=%d source=%s target=%s chamber=%s rank=%d dim=%d contribution=%.18g certified=%t interpretation=%s supports=%s failures=%s", t.Degree, t.Source, t.CandidateTarget, t.Chamber, t.TargetRank, t.ChamberDim, t.Contribution, t.MapCertified, t.Interpretation, strings.Join(t.Supports, ","), strings.Join(t.Failures, ","))
}
func FormatCrossLane(c CrossLaneExclusion) string {
	return fmt.Sprintf("linearHRminExcludedCandidate=%t quadraticPiTopExcludedCandidate=%t theorem=%t wouldAddLinearHRmin=%.18g wouldAddQuadraticPiTop=%.18g supports=%s failures=%s", c.LinearToHRMinExcludedCandidate, c.QuadraticToPiTopExcludedCandidate, c.CrossLaneExclusionTheorem, c.WouldAddLinearHRMinContribution, c.WouldAddQuadraticPiTopContribution, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}
func FormatChambers(c ChamberAudit) string {
	return fmt.Sprintf("H10Typed=%t H72Typed=%t typingTheorem=%t H10=%s H72=%s supports=%s failures=%s", c.H10Typed, c.H72Typed, c.ResponseChamberTypingTheorem, c.H10Description, c.H72Description, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}
func FormatCandidate(c AlphaCandidate) string {
	return fmt.Sprintf("expr=%s linear=%.18g quadratic=%.18g crossLinearHRmin=%.18g crossQuadraticPiTop=%.18g alpha=%.18g shape=%t native=%t supports=%s failures=%s", c.Expression, c.LinearContribution, c.QuadraticContribution, c.CrossLinearHRMinContribution, c.CrossQuadraticPiTopContribution, c.ReconstructedAlpha, c.ShapeCoherent, c.Native, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}
func FormatObstruction(o Obstruction) string {
	return fmt.Sprintf("reducedShape=%t targetCertified=%t crossCertified=%t alphaNative=%t wound=%s next=%s supports=%s failures=%s", o.ReducedExteriorShapeInherited, o.DegreeTargetSelectionCertified, o.CrossLaneExclusionCertified, o.AlphaNative, o.RemainingWound, o.NextGate, strings.Join(o.Supports, ","), strings.Join(o.Failures, ","))
}
func FormatR3(r R3Assessment) string {
	return fmt.Sprintf("finiteTripleSeal=%t YdagY=%t socketMagnitudeTyped=%t socketRankAlphaTyped=%t reducedB2Typed=%t targetCertified=%t alphaNative=%t sectorReadout=%t R3=%t R4=%t supports=%s failures=%s", r.PostOrientationFiniteTripleSeal, r.YDagYReadoutCarrierReady, r.SocketMagnitudeSourceTyped, r.SocketRankAlphaSourceTyped, r.ReducedB2ResponseShapeTyped, r.DegreeTargetSelectionCertified, r.AlphaNative, r.SectorTraceMagnitudeReadout, r.EligibleForR3, r.EligibleForR4, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}
func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s status=%s reducedShape=%t degreeTargets=%t crossLanes=%t alphaNative=%t updateNEff=%t updateCYukawa=%t updateCHiggs=%t R3=%t R4=%t", i.Classification, i.Status, i.ReducedShapeSolved, i.DegreeTargetsSolved, i.CrossLanesSolved, i.AlphaNative, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CanPromoteToR3, i.CanPromoteToR4)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%t targetMap=%t lambda1=%t lambda2=%t cross=%t linearHRmin=%t quadraticPiTop=%t chamberTyping=%t reducedFunctional=%t sharedFunctor=%t alphaSealed=%t nativeAlpha=%t rankRatios=%t transport=%t officialNEff=%t CYukawaCHiggs=%t numericalYukawa=%t sectorTrace=%t fullAF=%t AForient=%t finiteTriple=%t notR3=%t notR4=%t verdict=%s", f.Enforced, f.NoNativeDegreeTargetSelectionMap, f.NoNativeLambda1ToPiTopMap, f.NoNativeLambda2ToHRMinMap, f.NoCrossLaneExclusionTheorem, f.NoLinearHRMinExclusion, f.NoQuadraticPiTopExclusion, f.ResponseChamberTypingNotTheorem, f.NoNativeReducedExteriorFunctional, f.NoSharedDegreeTargetFunctor, f.AlphaStillSealed, f.NoNativeAlphaSource, f.SocketRankRatiosNotActivation, f.NoBoundaryAlphaTransport, f.NoOfficialNEffUpdate, f.NoCYukawaCHiggsUpdate, f.NoNumericalYukawa, f.NoSectorTraceMagnitude, f.NoFullUnbrokenAF, f.AForientNotFullAF, f.NoNativeFiniteTriple, f.NotR3, f.NotR4, f.Verdict)
}

func containsAll(haystack, needles []string) bool {
	m := make(map[string]bool, len(haystack))
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
func near(a, b float64) bool { return math.Abs(a-b) <= 1e-15 }

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NoNativeDegreeTargetSelectionMap && f.NoNativeLambda1ToPiTopMap && f.NoNativeLambda2ToHRMinMap && f.NoCrossLaneExclusionTheorem && f.NoLinearHRMinExclusion && f.NoQuadraticPiTopExclusion && f.ResponseChamberTypingNotTheorem && f.NoNativeReducedExteriorFunctional && f.NoSharedDegreeTargetFunctor && f.AlphaStillSealed && f.NoNativeAlphaSource && f.SocketRankRatiosNotActivation && f.NoBoundaryAlphaTransport && f.NoOfficialNEffUpdate && f.NoCYukawaCHiggsUpdate && f.NoNumericalYukawa && f.NoSectorTraceMagnitude && f.NoFullUnbrokenAF && f.AForientNotFullAF && f.NoNativeFiniteTriple && f.NotR3 && f.NotR4 && f.Verdict == StatusFirewallVerdict
}
