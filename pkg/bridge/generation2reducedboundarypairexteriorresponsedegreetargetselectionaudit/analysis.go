// Package generation2reducedboundarypairexteriorresponsedegreetargetselectionaudit implements
// Gate 870: Reduced BoundaryPair Exterior Response and Degree-Target Selection Audit.
//
// Gate 870 follows Gate 869's boundary-pair exterior truncation obstruction. Gate
// 869 showed that Lambda^bullet B_2 has the right finite degree shape, but the
// full exterior algebra still contains Lambda^0 B_2 and therefore does not by
// itself explain the absence of a constant response. Gate 870 audits the sharper
// reduced response candidate
//
//	R_B(s) = (1+s b_1)(1+s b_2)-1
//	       = s(b_1+b_2)+s^2(b_1 wedge b_2),
//
// which suppresses the zero-order term and stops after second degree because
// Lambda^3 B_2=0. The gate remains conservative: the reduced exterior response
// has the exact s+s^2 shape, but no native boundary response functional, no
// degree-target selection theorem, and no alpha_B theorem are certified.
package generation2reducedboundarypairexteriorresponsedegreetargetselectionaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE870-REDUCED-BOUNDARY-PAIR-EXTERIOR-RESPONSE-DEGREE-TARGET-SELECTION-AUDIT"

	AlphaB    = 0.0003878958469680527
	SBoundary = 0.0012924448188162962

	BoundaryPairDim = 2
	Lambda0B2Dim    = 1
	Lambda1B2Dim    = 2
	Lambda2B2Dim    = 1
	Lambda3B2Dim    = 0

	PiTopRank     = 3
	HRambientRank = 8
	H10Dim        = HRambientRank + BoundaryPairDim
	HRminRank     = 7
	Lambda4V8Rank = 70
	H72Dim        = Lambda4V8Rank + BoundaryPairDim

	Classification = "REDUCED_BOUNDARY_PAIR_EXTERIOR_RESPONSE_DEGREE_TARGET_SELECTION_OBSTRUCTION"
	R2Status       = "R2+++++_REDUCED_BOUNDARY_PAIR_EXTERIOR_RESPONSE_OBSTRUCTION"

	StatusGate869Inherited       = "PASS_GATE869_BOUNDARY_PAIR_EXTERIOR_TRUNCATION_OBSTRUCTION_INHERITED"
	StatusReducedResponseDefined = "PASS_REDUCED_BOUNDARY_EXTERIOR_RESPONSE_DEFINED"
	StatusZeroOrderSuppressed    = "PASS_ZERO_ORDER_TERM_SUPPRESSED_BY_REDUCED_RESPONSE"
	StatusCubicSuppressed        = "PASS_CUBIC_AND_HIGHER_TERMS_SUPPRESSED_BY_LAMBDA3_B2_ZERO"
	StatusDegreeOneTargetAudited = "PASS_DEGREE_ONE_TO_PI_TOP_TARGET_AUDITED"
	StatusDegreeTwoTargetAudited = "PASS_DEGREE_TWO_TO_H_R_MIN_TARGET_AUDITED"
	StatusCrossLaneAudited       = "PASS_CROSS_LANE_EXCLUSION_REQUIREMENT_AUDITED"
	StatusAlphaReconstructed     = "PASS_ALPHA_B_FORMALLY_RECONSTRUCTED_FROM_REDUCED_EXTERIOR_RESPONSE"
	StatusFirewallEnforced       = "PASS_REDUCED_BOUNDARY_EXTERIOR_RESPONSE_FIREWALL_ENFORCED"
	StatusLedgerFrozen           = "PASS_OFFICIAL_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusNoObservedDataUsed     = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusNextWound              = "PASS_NEXT_WOUND_IDENTIFIED_AS_DEGREE_TARGET_SELECTION_AND_NATIVE_RESPONSE_FUNCTIONAL"
	StatusFirewallVerdict        = "FIREWALL_PRESERVED_GATE870_REDUCED_EXTERIOR_RESPONSE_NOT_ALPHA_THEOREM"

	SupportReducedResponseExactShape       = "CONDITIONAL_SUPPORT_REDUCED_BOUNDARY_EXTERIOR_RESPONSE_HAS_EXACT_S_PLUS_S_SQUARED_SHAPE"
	SupportZeroOrderSuppressedByReduction  = "CONDITIONAL_SUPPORT_ZERO_ORDER_TERM_SUPPRESSED_BY_REDUCED_RESPONSE"
	SupportNoCubicFromLambda3B2Zero        = "CONDITIONAL_SUPPORT_NO_CUBIC_OR_HIGHER_TERMS_FROM_LAMBDA3_B2_EQUALS_ZERO"
	SupportGate868JetsHaveReducedCandidate = "CONDITIONAL_SUPPORT_GATE868_J1_AND_J2_INSERTIONS_HAVE_REDUCED_EXTERIOR_RESPONSE_CANDIDATE"
	SupportDegreeOnePiTopCandidate         = "CONDITIONAL_SUPPORT_DEGREE_ONE_REDUCED_RESPONSE_MATCHES_DOMINANT_SOCKET_LANE"
	SupportDegreeTwoHRMinCandidate         = "CONDITIONAL_SUPPORT_DEGREE_TWO_WEDGE_RESPONSE_MATCHES_ACTIVE_RIGHT_DOMAIN_LANE"
	SupportConstantAndHigherExplained      = "CONDITIONAL_SUPPORT_CONSTANT_AND_HIGHER_ORDER_STRUCTURE_EXPLAINED_BY_REDUCED_B2_RESPONSE"
	SupportAlphaShapeSharpened             = "CONDITIONAL_SUPPORT_ALPHA_B_POWER_SHAPE_SHARPENED_TO_REDUCED_BOUNDARY_PAIR_RESPONSE"

	FailureNoNativeReducedExteriorFunctional = "FAILED_ROUTE_NO_NATIVE_REDUCED_BOUNDARY_EXTERIOR_RESPONSE_FUNCTIONAL"
	FailureNoBoundaryFunctional              = "FAILED_ROUTE_REDUCED_RESPONSE_NOT_NATIVE_WITHOUT_BOUNDARY_FUNCTIONAL"
	FailureNoDegreeTargetSelection           = "FAILED_ROUTE_NO_DEGREE_TARGET_SELECTION_THEOREM"
	FailureNoTypedDegreeOneToPiTopMap        = "FAILED_ROUTE_NO_TYPED_DEGREE_ONE_TO_PI_TOP_MAP"
	FailureNoTypedDegreeTwoToHRMinMap        = "FAILED_ROUTE_NO_TYPED_DEGREE_TWO_TO_H_R_MIN_MAP"
	FailureNoLinearHRMinExclusion            = "FAILED_ROUTE_NO_NATIVE_EXCLUSION_OF_LINEAR_H_R_MIN_TERM"
	FailureNoQuadraticPiTopExclusion         = "FAILED_ROUTE_NO_NATIVE_EXCLUSION_OF_QUADRATIC_PI_TOP_TERM"
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
	FailureNotR3                             = "FAILED_ROUTE_NOT_R3_REDUCED_BOUNDARY_PAIR_RESPONSE_OBSTRUCTION"
	FailureNotR4                             = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	AlphaB, SBoundary float64
	AlphaNative       bool
	OfficialFrozen    bool
	R3, R4            bool
}

type ReducedResponse struct {
	Expression, Expansion                          string
	Lambda0Removed, DegreeOnePresent               bool
	DegreeTwoPresent, CubicAndHigherVanish         bool
	NativeFunctionalCertified, ExactShapeCandidate bool
	Supports, Failures                             []string
}

type DegreeTarget struct {
	Degree             int
	Source             string
	Target             string
	TargetRank         int
	ChamberDim         int
	Contribution       float64
	TargetMapCertified bool
	Supports, Failures []string
}

type DegreeTargetAudit struct {
	DegreeOne, DegreeTwo                              DegreeTarget
	DegreeTargetSelectionTheorem, SharedTargetFunctor bool
	Supports, Failures                                []string
}

type CrossLaneAudit struct {
	LinearHRMinAbsent, QuadraticPiTopAbsent       bool
	CrossLaneExclusionTheorem                     bool
	ZeroOrderSuppressed, CubicAndHigherSuppressed bool
	NativeReducedFunctional, AlphaNative          bool
	Supports, Failures                            []string
}

type AlphaCandidate struct {
	Expression, ReducedExpression string
	LinearContribution            float64
	QuadraticContribution         float64
	ReconstructedAlpha            float64
	ShapeCoherent, Native         bool
	Supports, Failures            []string
}

type Obstruction struct {
	ReducedResponseExplainsShape     bool
	NativeReducedFunctionalCertified bool
	DegreeTargetSelectionCertified   bool
	AlphaNative                      bool
	RemainingWound, NextGate         string
	Supports, Failures               []string
}

type R3Assessment struct {
	PostOrientationFiniteTripleSeal bool
	YDagYReadoutCarrierReady        bool
	SocketMagnitudeSourceTyped      bool
	SocketRankAlphaSourceTyped      bool
	ReducedB2ResponseShapeTyped     bool
	AlphaNative                     bool
	SectorTraceMagnitudeReadout     bool
	EligibleForR3, EligibleForR4    bool
	Supports, Failures              []string
}

type Impact struct {
	Classification, Status                                                           string
	ZeroOrderSolvedAtShapeLevel, CubicStopSolvedAtShapeLevel                         bool
	DegreeTargetsSolved, NativeReducedFunctionalSolved, AlphaNative                  bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4 bool
}

type Firewalls struct {
	Enforced                                                                                      bool
	NoNativeReducedExteriorFunctional, NoBoundaryFunctional, NoDegreeTargetSelection              bool
	NoTypedDegreeOneToPiTopMap, NoTypedDegreeTwoToHRMinMap, NoLinearHRMinExclusion                bool
	NoQuadraticPiTopExclusion, NoSharedDegreeTargetFunctor, AlphaStillSealed, NoNativeAlphaSource bool
	SocketRankRatiosNotActivation, NoBoundaryAlphaTransport, NoOfficialNEffUpdate                 bool
	NoCYukawaCHiggsUpdate, NoNumericalYukawa, NoSectorTraceMagnitude, NoFullUnbrokenAF            bool
	AForientNotFullAF, NoNativeFiniteTriple, NotR3, NotR4                                         bool
	Verdict                                                                                       string
}

type Audit struct {
	ID           string
	Ledger       Ledger
	Response     ReducedResponse
	Targets      DegreeTargetAudit
	CrossLane    CrossLaneAudit
	Candidate    AlphaCandidate
	Obstruction  Obstruction
	R3           R3Assessment
	Impact       Impact
	Firewalls    Firewalls
	Truth, Final string
}

func BuildDefault() (Audit, error) {
	linear := float64(PiTopRank) / float64(H10Dim) * SBoundary
	quadratic := float64(HRminRank) / float64(H72Dim) * SBoundary * SBoundary
	a := Audit{
		ID:     AuditID,
		Ledger: Ledger{AlphaB: AlphaB, SBoundary: SBoundary, OfficialFrozen: true},
		Response: ReducedResponse{
			Expression:     "R_B(s)=(1+s b1)(1+s b2)-1",
			Expansion:      "s(b1+b2)+s^2(b1 wedge b2)",
			Lambda0Removed: true, DegreeOnePresent: true, DegreeTwoPresent: true, CubicAndHigherVanish: true,
			NativeFunctionalCertified: false, ExactShapeCandidate: true,
			Supports: []string{StatusReducedResponseDefined, StatusZeroOrderSuppressed, StatusCubicSuppressed, SupportReducedResponseExactShape, SupportZeroOrderSuppressedByReduction, SupportNoCubicFromLambda3B2Zero, SupportConstantAndHigherExplained},
			Failures: []string{FailureNoNativeReducedExteriorFunctional, FailureNoBoundaryFunctional},
		},
		Targets: DegreeTargetAudit{
			DegreeOne:                    DegreeTarget{Degree: 1, Source: "s(b1+b2) in Lambda^1 B_2", Target: "Pi_top=e_+ tensor P_3 in H10", TargetRank: PiTopRank, ChamberDim: H10Dim, Contribution: linear, TargetMapCertified: false, Supports: []string{StatusDegreeOneTargetAudited, SupportDegreeOnePiTopCandidate}, Failures: []string{FailureNoTypedDegreeOneToPiTopMap}},
			DegreeTwo:                    DegreeTarget{Degree: 2, Source: "s^2(b1 wedge b2) in Lambda^2 B_2", Target: "H_R^min active punctured right domain in H72", TargetRank: HRminRank, ChamberDim: H72Dim, Contribution: quadratic, TargetMapCertified: false, Supports: []string{StatusDegreeTwoTargetAudited, SupportDegreeTwoHRMinCandidate}, Failures: []string{FailureNoTypedDegreeTwoToHRMinMap}},
			DegreeTargetSelectionTheorem: false, SharedTargetFunctor: false,
			Supports: []string{SupportGate868JetsHaveReducedCandidate, SupportAlphaShapeSharpened},
			Failures: []string{FailureNoDegreeTargetSelection, FailureNoSharedDegreeTargetFunctor, FailureNoTypedDegreeOneToPiTopMap, FailureNoTypedDegreeTwoToHRMinMap},
		},
		CrossLane: CrossLaneAudit{
			LinearHRMinAbsent: true, QuadraticPiTopAbsent: true, CrossLaneExclusionTheorem: false,
			ZeroOrderSuppressed: true, CubicAndHigherSuppressed: true, NativeReducedFunctional: false, AlphaNative: false,
			Supports: []string{StatusCrossLaneAudited, SupportZeroOrderSuppressedByReduction, SupportNoCubicFromLambda3B2Zero},
			Failures: []string{FailureNoDegreeTargetSelection, FailureNoLinearHRMinExclusion, FailureNoQuadraticPiTopExclusion},
		},
		Candidate: AlphaCandidate{
			Expression:         "alpha_B=(1/10)Tr_H10(Pi_top tau1(R_B))+(1/72)Tr_H72(P_HRmin tau2(R_B))",
			ReducedExpression:  "(3/10)s+(7/72)s^2",
			LinearContribution: linear, QuadraticContribution: quadratic, ReconstructedAlpha: linear + quadratic,
			ShapeCoherent: true, Native: false,
			Supports: []string{StatusGate869Inherited, StatusAlphaReconstructed, SupportReducedResponseExactShape, SupportGate868JetsHaveReducedCandidate, SupportAlphaShapeSharpened},
			Failures: []string{FailureNoNativeReducedExteriorFunctional, FailureNoDegreeTargetSelection, FailureAlphaStillSealed, FailureNoNativeAlphaSource},
		},
		Obstruction: Obstruction{
			ReducedResponseExplainsShape: true, NativeReducedFunctionalCertified: false, DegreeTargetSelectionCertified: false, AlphaNative: false,
			RemainingWound: "degree-target maps Lambda^1 B2 -> Pi_top and Lambda^2 B2 -> H_R^min, plus native reduced boundary response functional",
			NextGate:       "Gate 871 — Boundary Exterior Degree-Target Map Construction/Obstruction Audit",
			Supports:       []string{SupportReducedResponseExactShape, SupportZeroOrderSuppressedByReduction, SupportNoCubicFromLambda3B2Zero, SupportConstantAndHigherExplained},
			Failures:       []string{FailureNoNativeReducedExteriorFunctional, FailureNoBoundaryFunctional, FailureNoDegreeTargetSelection, FailureNoTypedDegreeOneToPiTopMap, FailureNoTypedDegreeTwoToHRMinMap, FailureAlphaStillSealed},
		},
		R3: R3Assessment{
			PostOrientationFiniteTripleSeal: true, YDagYReadoutCarrierReady: true, SocketMagnitudeSourceTyped: true, SocketRankAlphaSourceTyped: true, ReducedB2ResponseShapeTyped: true,
			AlphaNative: false, SectorTraceMagnitudeReadout: false, EligibleForR3: false, EligibleForR4: false,
			Supports: []string{SupportReducedResponseExactShape, SupportZeroOrderSuppressedByReduction, SupportNoCubicFromLambda3B2Zero},
			Failures: []string{FailureAlphaStillSealed, FailureNoSectorTraceMagnitude, FailureNotR3, FailureNotR4},
		},
		Impact: Impact{Classification: Classification, Status: R2Status, ZeroOrderSolvedAtShapeLevel: true, CubicStopSolvedAtShapeLevel: true, DegreeTargetsSolved: false, NativeReducedFunctionalSolved: false, AlphaNative: false},
		Firewalls: Firewalls{
			Enforced: true, NoNativeReducedExteriorFunctional: true, NoBoundaryFunctional: true, NoDegreeTargetSelection: true,
			NoTypedDegreeOneToPiTopMap: true, NoTypedDegreeTwoToHRMinMap: true, NoLinearHRMinExclusion: true, NoQuadraticPiTopExclusion: true, NoSharedDegreeTargetFunctor: true,
			AlphaStillSealed: true, NoNativeAlphaSource: true, SocketRankRatiosNotActivation: true, NoBoundaryAlphaTransport: true,
			NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NoNumericalYukawa: true, NoSectorTraceMagnitude: true,
			NoFullUnbrokenAF: true, AForientNotFullAF: true, NoNativeFiniteTriple: true, NotR3: true, NotR4: true,
			Verdict: StatusFirewallVerdict,
		},
		Truth: "Gate 870 conditionally upgrades the B2 exterior calculus into a reduced response R_B(s)=(1+s b1)(1+s b2)-1, explaining the zero-order suppression and cubic stop at shape level only.",
		Final: "Verdict: reduced exterior response has the exact s+s^2 form, but alpha_B remains sealed because degree-target maps and a native boundary response functional are not certified.",
	}
	if !near(a.Candidate.ReconstructedAlpha, AlphaB) {
		return Audit{}, fmt.Errorf("alpha reconstruction mismatch: got %.18g want %.18g", a.Candidate.ReconstructedAlpha, AlphaB)
	}
	return a, nil
}

func near(a, b float64) bool { return math.Abs(a-b) <= 1e-15 }

func containsAll(haystack []string, needles []string) bool {
	set := map[string]bool{}
	for _, h := range haystack {
		set[h] = true
	}
	for _, n := range needles {
		if !set[n] {
			return false
		}
	}
	return true
}

func firewallsOK(f Firewalls) bool {
	return f.Enforced && f.NoNativeReducedExteriorFunctional && f.NoBoundaryFunctional && f.NoDegreeTargetSelection && f.NoTypedDegreeOneToPiTopMap && f.NoTypedDegreeTwoToHRMinMap && f.NoLinearHRMinExclusion && f.NoQuadraticPiTopExclusion && f.NoSharedDegreeTargetFunctor && f.AlphaStillSealed && f.NoNativeAlphaSource && f.SocketRankRatiosNotActivation && f.NoBoundaryAlphaTransport && f.NoOfficialNEffUpdate && f.NoCYukawaCHiggsUpdate && f.NoNumericalYukawa && f.NoSectorTraceMagnitude && f.NoFullUnbrokenAF && f.AForientNotFullAF && f.NoNativeFiniteTriple && f.NotR3 && f.NotR4 && f.Verdict == StatusFirewallVerdict
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("alpha_B=%.18g s=%.18g official_frozen=%t alpha_native=%t R3=%t R4=%t", l.AlphaB, l.SBoundary, l.OfficialFrozen, l.AlphaNative, l.R3, l.R4)
}
func FormatResponse(r ReducedResponse) string {
	return fmt.Sprintf("%s => %s lambda0_removed=%t degree1=%t degree2=%t cubic_plus_vanish=%t native=%t supports=%s failures=%s", r.Expression, r.Expansion, r.Lambda0Removed, r.DegreeOnePresent, r.DegreeTwoPresent, r.CubicAndHigherVanish, r.NativeFunctionalCertified, strings.Join(r.Supports, ","), strings.Join(r.Failures, ","))
}
func FormatDegree(d DegreeTarget) string {
	return fmt.Sprintf("degree=%d source=%s target=%s rank=%d chamber=%d contribution=%.18g target_map=%t", d.Degree, d.Source, d.Target, d.TargetRank, d.ChamberDim, d.Contribution, d.TargetMapCertified)
}
func FormatTargets(t DegreeTargetAudit) string {
	return fmt.Sprintf("degree_one={%s} degree_two={%s} target_theorem=%t shared_functor=%t failures=%s", FormatDegree(t.DegreeOne), FormatDegree(t.DegreeTwo), t.DegreeTargetSelectionTheorem, t.SharedTargetFunctor, strings.Join(t.Failures, ","))
}
func FormatCrossLane(c CrossLaneAudit) string {
	return fmt.Sprintf("zero_suppressed=%t cubic_suppressed=%t linear_HRmin_absent=%t quadratic_PiTop_absent=%t cross_theorem=%t native_response=%t failures=%s", c.ZeroOrderSuppressed, c.CubicAndHigherSuppressed, c.LinearHRMinAbsent, c.QuadraticPiTopAbsent, c.CrossLaneExclusionTheorem, c.NativeReducedFunctional, strings.Join(c.Failures, ","))
}
func FormatCandidate(c AlphaCandidate) string {
	return fmt.Sprintf("%s => %s linear=%.18g quadratic=%.18g alpha=%.18g coherent=%t native=%t failures=%s", c.Expression, c.ReducedExpression, c.LinearContribution, c.QuadraticContribution, c.ReconstructedAlpha, c.ShapeCoherent, c.Native, strings.Join(c.Failures, ","))
}
func FormatObstruction(o Obstruction) string {
	return fmt.Sprintf("reduced_shape=%t native_functional=%t degree_targets=%t alpha_native=%t remaining=%s next=%s failures=%s", o.ReducedResponseExplainsShape, o.NativeReducedFunctionalCertified, o.DegreeTargetSelectionCertified, o.AlphaNative, o.RemainingWound, o.NextGate, strings.Join(o.Failures, ","))
}
func FormatR3(r R3Assessment) string {
	return fmt.Sprintf("post_orientation=%t YdagY=%t socket_magnitude=%t socket_rank_alpha=%t reduced_B2=%t alpha_native=%t sector_readout=%t R3=%t R4=%t failures=%s", r.PostOrientationFiniteTripleSeal, r.YDagYReadoutCarrierReady, r.SocketMagnitudeSourceTyped, r.SocketRankAlphaSourceTyped, r.ReducedB2ResponseShapeTyped, r.AlphaNative, r.SectorTraceMagnitudeReadout, r.EligibleForR3, r.EligibleForR4, strings.Join(r.Failures, ","))
}
func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s status=%s zero_order_shape=%t cubic_shape=%t degree_targets=%t native_functional=%t alpha_native=%t update_Neff=%t promote_R3=%t", i.Classification, i.Status, i.ZeroOrderSolvedAtShapeLevel, i.CubicStopSolvedAtShapeLevel, i.DegreeTargetsSolved, i.NativeReducedFunctionalSolved, i.AlphaNative, i.CanUpdateNEff, i.CanPromoteToR3)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%t verdict=%s native_reduced=%t target_selection=%t alpha_sealed=%t no_R3=%t no_R4=%t", f.Enforced, f.Verdict, f.NoNativeReducedExteriorFunctional, f.NoDegreeTargetSelection, f.AlphaStillSealed, f.NotR3, f.NotR4)
}

func Statuses() []string {
	return []string{
		StatusGate869Inherited, StatusReducedResponseDefined, StatusZeroOrderSuppressed, StatusCubicSuppressed, StatusDegreeOneTargetAudited, StatusDegreeTwoTargetAudited, StatusCrossLaneAudited, StatusAlphaReconstructed, StatusFirewallEnforced, StatusLedgerFrozen, StatusNoObservedDataUsed, StatusNextWound, StatusFirewallVerdict,
		SupportReducedResponseExactShape, SupportZeroOrderSuppressedByReduction, SupportNoCubicFromLambda3B2Zero, SupportGate868JetsHaveReducedCandidate, SupportDegreeOnePiTopCandidate, SupportDegreeTwoHRMinCandidate, SupportConstantAndHigherExplained, SupportAlphaShapeSharpened,
		FailureNoNativeReducedExteriorFunctional, FailureNoBoundaryFunctional, FailureNoDegreeTargetSelection, FailureNoTypedDegreeOneToPiTopMap, FailureNoTypedDegreeTwoToHRMinMap, FailureNoLinearHRMinExclusion, FailureNoQuadraticPiTopExclusion, FailureNoSharedDegreeTargetFunctor, FailureAlphaStillSealed, FailureNoNativeAlphaSource, FailureSocketRankRatiosNotActivation, FailureNoBoundaryAlphaTransport, FailureNoOfficialNEffUpdate, FailureNoCYukawaCHiggsUpdate, FailureNoNumericalYukawa, FailureNoSectorTraceMagnitude, FailureNoFullUnbrokenAF, FailureAForientNotFullAF, FailureNoNativeFiniteTriple, FailureNotR3, FailureNotR4,
	}
}
