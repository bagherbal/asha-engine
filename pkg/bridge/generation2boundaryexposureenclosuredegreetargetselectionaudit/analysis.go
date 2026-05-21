// Package generation2boundaryexposureenclosuredegreetargetselectionaudit implements
// Gate 872: Boundary Exposure/Enclosure Degree-Target Selection Audit.
//
// Gate 872 follows Gate 871's degree-target selection obstruction. Gate 870
// conditionally explained the reduced boundary-pair response shape
//
//	R_B(s)=(1+s b_1)(1+s b_2)-1=s(b_1+b_2)+s^2(b_1 wedge b_2),
//
// and Gate 871 isolated the remaining target assignment wound:
// Lambda^1 B_2 -> Pi_top and Lambda^2 B_2 -> H_R^min, with cross-lane
// exclusions. Gate 872 audits the sharper exposure/enclosure candidate:
// Lambda^1 B_2 is single-boundary exposure targeting the exposed dominant
// socket Pi_top, while Lambda^2 B_2 is full boundary-pair enclosure targeting
// the punctured active right domain H_R^min. The result is intentionally
// conservative: the exposure/enclosure interpretation is a coherent target
// selection candidate, but no native target-selection theorem, cross-lane
// exclusion theorem, or alpha_B theorem is certified.
package generation2boundaryexposureenclosuredegreetargetselectionaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE872-BOUNDARY-EXPOSURE-ENCLOSURE-DEGREE-TARGET-SELECTION-AUDIT"

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

	Classification = "BOUNDARY_EXPOSURE_ENCLOSURE_DEGREE_TARGET_SELECTION_OBSTRUCTION"
	R2Status       = "R2+++++_BOUNDARY_EXPOSURE_ENCLOSURE_DEGREE_TARGET_SELECTION_OBSTRUCTION"

	StatusGate871Inherited          = "PASS_GATE871_DEGREE_TARGET_SELECTION_WOUND_INHERITED"
	StatusReducedResponseInherited  = "PASS_REDUCED_B2_RESPONSE_SHAPE_INHERITED"
	StatusExposureCandidateAudited  = "PASS_LAMBDA1_B2_AS_EXPOSURE_CANDIDATE_AUDITED"
	StatusEnclosureCandidateAudited = "PASS_LAMBDA2_B2_AS_ENCLOSURE_CANDIDATE_AUDITED"
	StatusExposureTargetAudited     = "PASS_EXPOSURE_TO_PI_TOP_TARGET_AUDITED"
	StatusEnclosureTargetAudited    = "PASS_ENCLOSURE_TO_H_R_MIN_TARGET_AUDITED"
	StatusCrossLaneTypeAudited      = "PASS_CROSS_LANE_EXCLUSION_BY_EXPOSURE_ENCLOSURE_TYPE_AUDITED"
	StatusPunctureRoleAudited       = "PASS_PUNCTURE_ROLE_IN_ENCLOSED_DOMAIN_AUDITED"
	StatusAlphaReconstructed        = "PASS_ALPHA_B_RECONSTRUCTED_FROM_EXPOSURE_ENCLOSURE_TARGETS"
	StatusFirewallEnforced          = "PASS_EXPOSURE_ENCLOSURE_TARGET_SELECTION_FIREWALL_ENFORCED"
	StatusLedgerFrozen              = "PASS_OFFICIAL_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusNoObservedDataUsed        = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusNextWound                 = "PASS_NEXT_WOUND_IDENTIFIED_AS_BOUNDARY_ALPHA_SEAL_OR_THEOREM_ELIGIBILITY"
	StatusFirewallVerdict           = "FIREWALL_PRESERVED_GATE872_EXPOSURE_ENCLOSURE_SELECTION_NOT_ALPHA_THEOREM"

	SupportReducedB2ShapeInherited             = "CONDITIONAL_SUPPORT_REDUCED_B2_RESPONSE_GIVES_CORRECT_ALPHA_POWER_SHAPE"
	SupportLambda1AsExposure                   = "CONDITIONAL_SUPPORT_LAMBDA1B2_AS_SINGLE_BOUNDARY_EXPOSURE"
	SupportExposureTargetsDominantSocket       = "CONDITIONAL_SUPPORT_SINGLE_EXPOSURE_TARGETS_DOMINANT_SOCKET_CANDIDATE"
	SupportLambda2AsEnclosure                  = "CONDITIONAL_SUPPORT_LAMBDA2B2_AS_FULL_BOUNDARY_PAIR_ENCLOSURE"
	SupportEnclosureTargetsPuncturedDomain     = "CONDITIONAL_SUPPORT_FULL_ENCLOSURE_TARGETS_PUNCTURED_ACTIVE_RIGHT_DOMAIN_CANDIDATE"
	SupportCrossLaneTypeCandidate              = "CONDITIONAL_SUPPORT_CROSS_LANE_EXCLUSION_BY_EXPOSURE_ENCLOSURE_TYPE"
	SupportPunctureRequiredForRankSeven        = "CONDITIONAL_SUPPORT_PUNCTURE_SUBTRACTION_IS_REQUIRED_FOR_DEGREE_TWO_TARGET_RANK_SEVEN"
	SupportDegreeTargetsHaveTypeInterpretation = "CONDITIONAL_SUPPORT_DEGREE_TARGETS_HAVE_EXPOSURE_ENCLOSURE_TYPE_INTERPRETATION"
	SupportAlphaPowerShapeStillCoherent        = "CONDITIONAL_SUPPORT_ALPHA_B_REDUCED_EXTERIOR_POWER_SHAPE_REMAINS_COHERENT"
	SupportAlphaSealCandidateMatured           = "CONDITIONAL_SUPPORT_BOUNDARY_ALPHA_CAN_BE_CLASSIFIED_AS_EXPOSURE_ENCLOSURE_RESPONSE_SEAL_CANDIDATE"

	FailureNoNativeExposureEnclosureMap      = "FAILED_ROUTE_NO_NATIVE_EXPOSURE_ENCLOSURE_TARGET_SELECTION_MAP"
	FailureNoNativeExposureToPiTopMap        = "FAILED_ROUTE_NO_NATIVE_EXPOSURE_TO_PI_TOP_MAP"
	FailureNoNativeEnclosureToHRMinMap       = "FAILED_ROUTE_NO_NATIVE_ENCLOSURE_TO_H_R_MIN_MAP"
	FailureNoNativeCrossLaneExclusionTheorem = "FAILED_ROUTE_NO_NATIVE_CROSS_LANE_EXCLUSION_THEOREM"
	FailureNoNativePuncturedEnclosure        = "FAILED_ROUTE_NO_NATIVE_PUNCTURED_ENCLOSURE_SELECTION_THEOREM"
	FailureNoNativeLambda1ToPiTopMap         = "FAILED_ROUTE_NO_NATIVE_LAMBDA1B2_TO_PI_TOP_MAP"
	FailureNoNativeLambda2ToHRMinMap         = "FAILED_ROUTE_NO_NATIVE_LAMBDA2B2_TO_H_R_MIN_MAP"
	FailureNoNativeDegreeTargetSelectionMap  = "FAILED_ROUTE_NO_NATIVE_DEGREE_TARGET_SELECTION_MAP"
	FailureNoLinearHRMinExclusion            = "FAILED_ROUTE_NO_NATIVE_EXCLUSION_OF_LINEAR_H_R_MIN_TERM"
	FailureNoQuadraticPiTopExclusion         = "FAILED_ROUTE_NO_NATIVE_EXCLUSION_OF_QUADRATIC_PI_TOP_TERM"
	FailureNoNativeReducedExteriorFunctional = "FAILED_ROUTE_NO_NATIVE_REDUCED_BOUNDARY_EXTERIOR_RESPONSE_FUNCTIONAL"
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
	FailureNotR3                             = "FAILED_ROUTE_NOT_R3_BOUNDARY_EXPOSURE_ENCLOSURE_TARGET_SELECTION_OBSTRUCTION"
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

type ExposureEnclosureCandidate struct {
	Degree             int
	ExteriorObject     string
	Type               string
	CandidateTarget    string
	TargetRank         int
	Chamber            string
	ChamberDim         int
	Contribution       float64
	MapCertified       bool
	NativeTheorem      bool
	Interpretation     string
	Supports, Failures []string
}

type CrossLaneTypeAudit struct {
	ExposureToHRMinExcludedCandidate, EnclosureToPiTopExcludedCandidate   bool
	CrossLaneExclusionTheorem                                             bool
	WouldAddExposureHRMinContribution, WouldAddEnclosurePiTopContribution float64
	TypeReason, RemainingWound                                            string
	Supports, Failures                                                    []string
}

type PunctureRole struct {
	AmbientRightRank, PunctureRank, ActiveRightRank int
	PunctureRequiredForRankSeven                    bool
	PuncturedEnclosureTheorem                       bool
	Puncture, ActiveDomain                          string
	Supports, Failures                              []string
}

type AlphaCandidate struct {
	Expression                      string
	LinearContribution              float64
	QuadraticContribution           float64
	CrossExposureHRMinContribution  float64
	CrossEnclosurePiTopContribution float64
	ReconstructedAlpha              float64
	ShapeCoherent, Native           bool
	Supports, Failures              []string
}

type Obstruction struct {
	ReducedExteriorShapeInherited bool
	ExposureEnclosureTyped        bool
	TargetSelectionCertified      bool
	CrossLaneExclusionCertified   bool
	AlphaNative                   bool
	RemainingWound, NextGate      string
	Supports, Failures            []string
}

type R3Assessment struct {
	PostOrientationFiniteTripleSeal bool
	YDagYReadoutCarrierReady        bool
	SocketMagnitudeSourceTyped      bool
	SocketRankAlphaSourceTyped      bool
	ReducedB2ResponseShapeTyped     bool
	ExposureEnclosureCandidate      bool
	TargetSelectionCertified        bool
	AlphaNative                     bool
	SectorTraceMagnitudeReadout     bool
	EligibleForR3, EligibleForR4    bool
	Supports, Failures              []string
}

type Impact struct {
	Classification, Status                                                                         string
	ReducedShapeSolved, ExposureEnclosureTyped, DegreeTargetsSolved, CrossLanesSolved, AlphaNative bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4               bool
}

type Firewalls struct {
	Enforced                                                                                                           bool
	NoNativeExposureEnclosureMap, NoNativeExposureToPiTopMap, NoNativeEnclosureToHRMinMap, NoNativeCrossLaneExclusion  bool
	NoNativePuncturedEnclosure, NoNativeLambda1ToPiTopMap, NoNativeLambda2ToHRMinMap, NoNativeDegreeTargetSelectionMap bool
	NoLinearHRMinExclusion, NoQuadraticPiTopExclusion, NoNativeReducedExteriorFunctional                               bool
	AlphaStillSealed, NoNativeAlphaSource, SocketRankRatiosNotActivation, NoBoundaryAlphaTransport                     bool
	NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate, NoNumericalYukawa, NoSectorTraceMagnitude                             bool
	NoFullUnbrokenAF, AForientNotFullAF, NoNativeFiniteTriple, NotR3, NotR4                                            bool
	Verdict                                                                                                            string
}

type Audit struct {
	ID           string
	Ledger       Ledger
	Response     ReducedResponseInherited
	Exposure     ExposureEnclosureCandidate
	Enclosure    ExposureEnclosureCandidate
	CrossLane    CrossLaneTypeAudit
	Puncture     PunctureRole
	Candidate    AlphaCandidate
	Obstruction  Obstruction
	R3           R3Assessment
	Impact       Impact
	Firewalls    Firewalls
	Truth, Final string
}

func BuildDefault() (Audit, error) {
	if H10Dim != 10 || H72Dim != 72 || PiTopRank != 3 || HRminRank != 7 || Lambda3B2Dim != 0 {
		return Audit{}, fmt.Errorf("unexpected Gate 872 dimension ledger")
	}
	linear := float64(PiTopRank) / float64(H10Dim) * SBoundary
	quadratic := float64(HRminRank) / float64(H72Dim) * SBoundary * SBoundary
	crossExposureHRMin := float64(HRminRank) / float64(H72Dim) * SBoundary
	crossEnclosurePiTop := float64(PiTopRank) / float64(H10Dim) * SBoundary * SBoundary
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
			Supports: []string{StatusReducedResponseInherited, SupportReducedB2ShapeInherited, SupportAlphaPowerShapeStillCoherent},
			Failures: []string{FailureNoNativeReducedExteriorFunctional},
		},
		Exposure: ExposureEnclosureCandidate{
			Degree: 1, ExteriorObject: "Lambda^1 B_2", Type: "single-boundary exposure", CandidateTarget: "Pi_top=e_+ tensor P_3 exposed dominant socket", TargetRank: PiTopRank, Chamber: "H10=H_R^ambient plus B_2", ChamberDim: H10Dim,
			Contribution: linear, MapCertified: false, NativeTheorem: false, Interpretation: "degree-one exposure targets the normalized visible dominant color socket candidate",
			Supports: []string{StatusExposureCandidateAudited, StatusExposureTargetAudited, SupportLambda1AsExposure, SupportExposureTargetsDominantSocket, SupportDegreeTargetsHaveTypeInterpretation},
			Failures: []string{FailureNoNativeExposureEnclosureMap, FailureNoNativeExposureToPiTopMap, FailureNoNativeLambda1ToPiTopMap},
		},
		Enclosure: ExposureEnclosureCandidate{
			Degree: 2, ExteriorObject: "Lambda^2 B_2", Type: "full boundary-pair enclosure", CandidateTarget: "H_R^min punctured enclosed active right domain", TargetRank: HRminRank, Chamber: "H72=Lambda^4 V_8 plus B_2", ChamberDim: H72Dim,
			Contribution: quadratic, MapCertified: false, NativeTheorem: false, Interpretation: "degree-two enclosure targets the punctured active right edge-domain candidate",
			Supports: []string{StatusEnclosureCandidateAudited, StatusEnclosureTargetAudited, SupportLambda2AsEnclosure, SupportEnclosureTargetsPuncturedDomain, SupportDegreeTargetsHaveTypeInterpretation},
			Failures: []string{FailureNoNativeExposureEnclosureMap, FailureNoNativeEnclosureToHRMinMap, FailureNoNativeLambda2ToHRMinMap},
		},
		CrossLane: CrossLaneTypeAudit{
			ExposureToHRMinExcludedCandidate: true, EnclosureToPiTopExcludedCandidate: true, CrossLaneExclusionTheorem: false,
			WouldAddExposureHRMinContribution: crossExposureHRMin, WouldAddEnclosurePiTopContribution: crossEnclosurePiTop,
			TypeReason: "exposure should target exposed dominant socket; enclosure should target enclosed active right domain", RemainingWound: "native cross-lane exclusion theorem",
			Supports: []string{StatusCrossLaneTypeAudited, SupportCrossLaneTypeCandidate},
			Failures: []string{FailureNoNativeCrossLaneExclusionTheorem, FailureNoLinearHRMinExclusion, FailureNoQuadraticPiTopExclusion},
		},
		Puncture: PunctureRole{
			AmbientRightRank: HRambientRank, PunctureRank: 1, ActiveRightRank: HRminRank, PunctureRequiredForRankSeven: true, PuncturedEnclosureTheorem: false,
			Puncture: "e_+ tensor P_1 neutral right puncture", ActiveDomain: "H_R^min=(C_R^2 tensor W) minus (e_+ tensor P_1)",
			Supports: []string{StatusPunctureRoleAudited, SupportPunctureRequiredForRankSeven},
			Failures: []string{FailureNoNativePuncturedEnclosure},
		},
		Candidate: AlphaCandidate{
			Expression:         "alpha_B=[rank(Pi_top)/10]s+[rank(H_R^min)/72]s^2 with Lambda1B2 as exposure and Lambda2B2 as enclosure",
			LinearContribution: linear, QuadraticContribution: quadratic, CrossExposureHRMinContribution: crossExposureHRMin, CrossEnclosurePiTopContribution: crossEnclosurePiTop,
			ReconstructedAlpha: alpha, ShapeCoherent: true, Native: false,
			Supports: []string{StatusAlphaReconstructed, SupportReducedB2ShapeInherited, SupportDegreeTargetsHaveTypeInterpretation, SupportAlphaSealCandidateMatured},
			Failures: []string{FailureNoNativeExposureEnclosureMap, FailureNoNativeCrossLaneExclusionTheorem, FailureAlphaStillSealed},
		},
		Obstruction: Obstruction{
			ReducedExteriorShapeInherited: true, ExposureEnclosureTyped: true, TargetSelectionCertified: false, CrossLaneExclusionCertified: false, AlphaNative: false,
			RemainingWound: "native exposure/enclosure degree-target selection map and cross-lane exclusion theorem",
			NextGate:       "Gate 873 — BoundaryAlpha Seal / R3 Eligibility Audit",
			Supports:       []string{StatusGate871Inherited, SupportDegreeTargetsHaveTypeInterpretation, SupportAlphaSealCandidateMatured},
			Failures:       []string{FailureNoNativeExposureEnclosureMap, FailureNoNativeCrossLaneExclusionTheorem, FailureAlphaStillSealed},
		},
		R3: R3Assessment{
			PostOrientationFiniteTripleSeal: true, YDagYReadoutCarrierReady: true, SocketMagnitudeSourceTyped: true, SocketRankAlphaSourceTyped: true, ReducedB2ResponseShapeTyped: true,
			ExposureEnclosureCandidate: true, TargetSelectionCertified: false, AlphaNative: false, SectorTraceMagnitudeReadout: false, EligibleForR3: false, EligibleForR4: false,
			Supports: []string{SupportReducedB2ShapeInherited, SupportDegreeTargetsHaveTypeInterpretation, SupportAlphaSealCandidateMatured},
			Failures: []string{FailureNoNativeExposureEnclosureMap, FailureAlphaStillSealed, FailureNoSectorTraceMagnitude, FailureNotR3, FailureNotR4},
		},
		Impact: Impact{Classification: Classification, Status: R2Status, ReducedShapeSolved: true, ExposureEnclosureTyped: true, DegreeTargetsSolved: false, CrossLanesSolved: false, AlphaNative: false, CanUpdateNEff: false, CanUpdateCYukawa: false, CanUpdateCHiggs: false, CanPromoteToR3: false, CanPromoteToR4: false},
		Firewalls: Firewalls{
			Enforced: true, NoNativeExposureEnclosureMap: true, NoNativeExposureToPiTopMap: true, NoNativeEnclosureToHRMinMap: true, NoNativeCrossLaneExclusion: true, NoNativePuncturedEnclosure: true,
			NoNativeLambda1ToPiTopMap: true, NoNativeLambda2ToHRMinMap: true, NoNativeDegreeTargetSelectionMap: true, NoLinearHRMinExclusion: true, NoQuadraticPiTopExclusion: true, NoNativeReducedExteriorFunctional: true,
			AlphaStillSealed: true, NoNativeAlphaSource: true, SocketRankRatiosNotActivation: true, NoBoundaryAlphaTransport: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true,
			NoNumericalYukawa: true, NoSectorTraceMagnitude: true, NoFullUnbrokenAF: true, AForientNotFullAF: true, NoNativeFiniteTriple: true, NotR3: true, NotR4: true,
			Verdict: StatusFirewallVerdict,
		},
		Truth: "Gate 872 gives the remaining degree-target wound an exposure/enclosure type interpretation: Lambda^1 B2 is a single-boundary exposure candidate for Pi_top and Lambda^2 B2 is a full boundary-pair enclosure candidate for H_R^min. The target maps and cross-lane exclusions remain unproved.",
		Final: "VERDICT: CONDITIONAL_SUPPORT_EXPOSURE_ENCLOSURE_TARGET_SELECTION_CANDIDATE, FAILED_ROUTE_NO_NATIVE_EXPOSURE_ENCLOSURE_TARGET_SELECTION_MAP, FAILED_ROUTE_ALPHA_B_REMAINS_SEALED, NOT_R3.",
	}, nil
}

func Statuses() []string {
	return []string{
		StatusGate871Inherited, StatusReducedResponseInherited, StatusExposureCandidateAudited, StatusEnclosureCandidateAudited, StatusExposureTargetAudited, StatusEnclosureTargetAudited, StatusCrossLaneTypeAudited, StatusPunctureRoleAudited, StatusAlphaReconstructed, StatusFirewallEnforced, StatusLedgerFrozen, StatusNoObservedDataUsed, StatusNextWound, StatusFirewallVerdict,
		SupportReducedB2ShapeInherited, SupportLambda1AsExposure, SupportExposureTargetsDominantSocket, SupportLambda2AsEnclosure, SupportEnclosureTargetsPuncturedDomain, SupportCrossLaneTypeCandidate, SupportPunctureRequiredForRankSeven, SupportDegreeTargetsHaveTypeInterpretation, SupportAlphaPowerShapeStillCoherent, SupportAlphaSealCandidateMatured,
		FailureNoNativeExposureEnclosureMap, FailureNoNativeExposureToPiTopMap, FailureNoNativeEnclosureToHRMinMap, FailureNoNativeCrossLaneExclusionTheorem, FailureNoNativePuncturedEnclosure, FailureNoNativeLambda1ToPiTopMap, FailureNoNativeLambda2ToHRMinMap, FailureNoNativeDegreeTargetSelectionMap, FailureNoLinearHRMinExclusion, FailureNoQuadraticPiTopExclusion, FailureNoNativeReducedExteriorFunctional, FailureAlphaStillSealed, FailureNoNativeAlphaSource, FailureSocketRankRatiosNotActivation, FailureNoBoundaryAlphaTransport, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNumericalYukawa, FailureNoSectorTraceMagnitude, FailureNoFullUnbrokenAF, FailureAForientNotFullAF, FailureNoNativeFiniteTriple, FailureNotR3, FailureNotR4,
	}
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("alpha_B=%.18g s=%.18g alphaNative=%t officialFrozen=%t R3=%t R4=%t", l.AlphaB, l.SBoundary, l.AlphaNative, l.OfficialFrozen, l.R3, l.R4)
}
func FormatResponse(r ReducedResponseInherited) string {
	return fmt.Sprintf("response=%s zeroRemoved=%t higherTruncated=%t degree1=%t degree2=%t nativeFunctional=%t supports=%s failures=%s", r.Expression, r.ZeroOrderSuppressed, r.HigherTruncated, r.DegreeOnePresent, r.DegreeTwoPresent, r.NativeFunctionalCertified, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}
func FormatExposureEnclosure(c ExposureEnclosureCandidate) string {
	return fmt.Sprintf("degree=%d exterior=%s type=%s target=%s rank=%d chamber=%s dim=%d contribution=%.18g mapCertified=%t native=%t interpretation=%s supports=%s failures=%s", c.Degree, c.ExteriorObject, c.Type, c.CandidateTarget, c.TargetRank, c.Chamber, c.ChamberDim, c.Contribution, c.MapCertified, c.NativeTheorem, c.Interpretation, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}
func FormatCrossLane(c CrossLaneTypeAudit) string {
	return fmt.Sprintf("exposureToHRminExcludedCandidate=%t enclosureToPiTopExcludedCandidate=%t theorem=%t wouldAddExposureHRmin=%.18g wouldAddEnclosurePiTop=%.18g typeReason=%s wound=%s supports=%s failures=%s", c.ExposureToHRMinExcludedCandidate, c.EnclosureToPiTopExcludedCandidate, c.CrossLaneExclusionTheorem, c.WouldAddExposureHRMinContribution, c.WouldAddEnclosurePiTopContribution, c.TypeReason, c.RemainingWound, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}
func FormatPuncture(p PunctureRole) string {
	return fmt.Sprintf("ambientRightRank=%d punctureRank=%d activeRightRank=%d punctureRequired=%t theorem=%t puncture=%s active=%s supports=%s failures=%s", p.AmbientRightRank, p.PunctureRank, p.ActiveRightRank, p.PunctureRequiredForRankSeven, p.PuncturedEnclosureTheorem, p.Puncture, p.ActiveDomain, strings.Join(p.Supports, ","), strings.Join(p.Failures, ","))
}
func FormatCandidate(c AlphaCandidate) string {
	return fmt.Sprintf("expr=%s linear=%.18g quadratic=%.18g crossExposureHRmin=%.18g crossEnclosurePiTop=%.18g alpha=%.18g shape=%t native=%t supports=%s failures=%s", c.Expression, c.LinearContribution, c.QuadraticContribution, c.CrossExposureHRMinContribution, c.CrossEnclosurePiTopContribution, c.ReconstructedAlpha, c.ShapeCoherent, c.Native, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}
func FormatObstruction(o Obstruction) string {
	return fmt.Sprintf("reducedShape=%t exposureEnclosureTyped=%t targetCertified=%t crossCertified=%t alphaNative=%t wound=%s next=%s supports=%s failures=%s", o.ReducedExteriorShapeInherited, o.ExposureEnclosureTyped, o.TargetSelectionCertified, o.CrossLaneExclusionCertified, o.AlphaNative, o.RemainingWound, o.NextGate, strings.Join(o.Supports, ","), strings.Join(o.Failures, ","))
}
func FormatR3(r R3Assessment) string {
	return fmt.Sprintf("finiteTripleSeal=%t YdagY=%t socketMagnitudeTyped=%t socketRankAlphaTyped=%t reducedB2Typed=%t exposureEnclosure=%t targetCertified=%t alphaNative=%t sectorReadout=%t R3=%t R4=%t supports=%s failures=%s", r.PostOrientationFiniteTripleSeal, r.YDagYReadoutCarrierReady, r.SocketMagnitudeSourceTyped, r.SocketRankAlphaSourceTyped, r.ReducedB2ResponseShapeTyped, r.ExposureEnclosureCandidate, r.TargetSelectionCertified, r.AlphaNative, r.SectorTraceMagnitudeReadout, r.EligibleForR3, r.EligibleForR4, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}
func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s status=%s reducedShape=%t exposureEnclosure=%t degreeTargets=%t crossLanes=%t alphaNative=%t updateNEff=%t updateCYukawa=%t updateCHiggs=%t R3=%t R4=%t", i.Classification, i.Status, i.ReducedShapeSolved, i.ExposureEnclosureTyped, i.DegreeTargetsSolved, i.CrossLanesSolved, i.AlphaNative, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CanPromoteToR3, i.CanPromoteToR4)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%t exposureEnclosureMap=%t exposureToPiTop=%t enclosureToHRmin=%t cross=%t puncture=%t lambda1=%t lambda2=%t targetMap=%t linearHRmin=%t quadraticPiTop=%t reducedFunctional=%t alphaSealed=%t nativeAlpha=%t rankRatios=%t transport=%t officialNEff=%t CYukawaCHiggs=%t numericalYukawa=%t sectorTrace=%t fullAF=%t AForient=%t finiteTriple=%t notR3=%t notR4=%t verdict=%s", f.Enforced, f.NoNativeExposureEnclosureMap, f.NoNativeExposureToPiTopMap, f.NoNativeEnclosureToHRMinMap, f.NoNativeCrossLaneExclusion, f.NoNativePuncturedEnclosure, f.NoNativeLambda1ToPiTopMap, f.NoNativeLambda2ToHRMinMap, f.NoNativeDegreeTargetSelectionMap, f.NoLinearHRMinExclusion, f.NoQuadraticPiTopExclusion, f.NoNativeReducedExteriorFunctional, f.AlphaStillSealed, f.NoNativeAlphaSource, f.SocketRankRatiosNotActivation, f.NoBoundaryAlphaTransport, f.NoOfficialNEffUpdate, f.NoCYukawaCHiggsUpdate, f.NoNumericalYukawa, f.NoSectorTraceMagnitude, f.NoFullUnbrokenAF, f.AForientNotFullAF, f.NoNativeFiniteTriple, f.NotR3, f.NotR4, f.Verdict)
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
	return f.Enforced && f.NoNativeExposureEnclosureMap && f.NoNativeExposureToPiTopMap && f.NoNativeEnclosureToHRMinMap && f.NoNativeCrossLaneExclusion && f.NoNativePuncturedEnclosure && f.NoNativeLambda1ToPiTopMap && f.NoNativeLambda2ToHRMinMap && f.NoNativeDegreeTargetSelectionMap && f.NoLinearHRMinExclusion && f.NoQuadraticPiTopExclusion && f.NoNativeReducedExteriorFunctional && f.AlphaStillSealed && f.NoNativeAlphaSource && f.SocketRankRatiosNotActivation && f.NoBoundaryAlphaTransport && f.NoOfficialNEffUpdate && f.NoCYukawaCHiggsUpdate && f.NoNumericalYukawa && f.NoSectorTraceMagnitude && f.NoFullUnbrokenAF && f.AForientNotFullAF && f.NoNativeFiniteTriple && f.NotR3 && f.NotR4 && f.Verdict == StatusFirewallVerdict
}
