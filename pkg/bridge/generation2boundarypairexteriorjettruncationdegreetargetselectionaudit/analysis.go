// Package generation2boundarypairexteriorjettruncationdegreetargetselectionaudit implements
// Gate 869: BoundaryPair Exterior-Jet Truncation and Degree-Target Selection Audit.
//
// Gate 869 follows Gate 868's boundary jet-response obstruction. It audits the
// next finite-calculus candidate for the alpha_B power shape: the exterior
// calculus of the rank-two boundary pair B_2. Since Lambda^3 B_2=0, the
// exterior-degree candidate can explain why only degree-one and degree-two
// response orders are available. The gate is conservative: exterior truncation
// is a strong shape candidate, but it does not by itself certify the target
// maps Lambda^1 B_2 -> Pi_top, Lambda^2 B_2 -> H_R^min, zero-order suppression,
// or the native BoundaryAlpha response functional.
package generation2boundarypairexteriorjettruncationdegreetargetselectionaudit

import (
	"fmt"
	"math"
	"strings"
)

const (
	AuditID = "GATE869-BOUNDARY-PAIR-EXTERIOR-JET-TRUNCATION-DEGREE-TARGET-SELECTION-AUDIT"

	AlphaB       = 0.0003878958469680527
	SBoundary    = 0.0012924448188162962
	OfficialNEff = 3.0023273474722147

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

	Classification = "BOUNDARY_PAIR_EXTERIOR_JET_TRUNCATION_DEGREE_TARGET_SELECTION_OBSTRUCTION"
	R2Status       = "R2+++++_BOUNDARY_PAIR_EXTERIOR_JET_TRUNCATION_OBSTRUCTION"

	StatusGate868Inherited       = "PASS_GATE868_BOUNDARY_JET_RESPONSE_OBSTRUCTION_INHERITED"
	StatusBoundaryPairAudited    = "PASS_BOUNDARY_PAIR_B2_EXTERIOR_CALCULUS_AUDITED"
	StatusExteriorTruncation     = "PASS_EXTERIOR_TRUNCATION_LAMBDA3_B2_ZERO_AUDITED"
	StatusDegreeOneTargetAudited = "PASS_DEGREE_ONE_DOMINANT_SOCKET_TARGET_AUDITED"
	StatusDegreeTwoTargetAudited = "PASS_DEGREE_TWO_ACTIVE_RIGHT_DOMAIN_TARGET_AUDITED"
	StatusZeroOrderAudited       = "PASS_ZERO_ORDER_SUPPRESSION_REQUIREMENT_AUDITED"
	StatusCrossLaneAudited       = "PASS_CROSS_LANE_EXCLUSION_REQUIREMENT_AUDITED"
	StatusAlphaReconstructed     = "PASS_ALPHA_B_FORMALLY_RECONSTRUCTED_FROM_EXTERIOR_DEGREE_CANDIDATE"
	StatusFirewallEnforced       = "PASS_BOUNDARY_PAIR_EXTERIOR_RESPONSE_FIREWALL_ENFORCED"
	StatusLedgerFrozen           = "PASS_OFFICIAL_N_EFF_C_YUKAWA_C_HIGGS_LEDGER_UPDATES_BLOCKED"
	StatusNoObservedDataUsed     = "PASS_NO_OBSERVED_MASS_CKM_PMNS_HIGGS_DATA_USED"
	StatusNextWound              = "PASS_NEXT_WOUND_IDENTIFIED_AS_DEGREE_TARGET_AND_BOUNDARY_RESPONSE_FUNCTIONAL"
	StatusFirewallVerdict        = "FIREWALL_PRESERVED_GATE869_EXTERIOR_TRUNCATION_NOT_ALPHA_THEOREM"

	SupportAlphaPowerMatchesB2Exterior       = "CONDITIONAL_SUPPORT_ALPHA_B_POWER_STRUCTURE_MATCHES_B2_EXTERIOR_JET_TRUNCATION"
	SupportFirstSecondOrdersHaveB2Candidate  = "CONDITIONAL_SUPPORT_FIRST_AND_SECOND_RESPONSE_ORDERS_HAVE_BOUNDARY_PAIR_SOURCE_CANDIDATE"
	SupportNoCubicFromLambda3B2Zero          = "CONDITIONAL_SUPPORT_NO_CUBIC_OR_HIGHER_TERMS_FROM_LAMBDA3_B2_EQUALS_ZERO"
	SupportDegreeOneDominantSocketCandidate  = "CONDITIONAL_SUPPORT_DEGREE_ONE_BOUNDARY_RESPONSE_LANDS_ON_DOMINANT_SOCKET_CANDIDATE"
	SupportDegreeTwoActiveRightCandidate     = "CONDITIONAL_SUPPORT_DEGREE_TWO_BOUNDARY_RESPONSE_LANDS_ON_ACTIVE_RIGHT_DOMAIN_CANDIDATE"
	SupportGate866SocketRanksCompatible      = "CONDITIONAL_SUPPORT_GATE866_SOCKET_RANK_SOURCES_COMPATIBLE_WITH_B2_EXTERIOR_DEGREES"
	SupportExteriorCalculusExplainsStopShape = "CONDITIONAL_SUPPORT_B2_EXTERIOR_CALCULUS_EXPLAINS_S_AND_S_SQUARED_STOP_SHAPE"
	SupportBoundaryJetSealSharpened          = "CONDITIONAL_SUPPORT_BOUNDARY_ALPHA_JET_SEAL_SHARPENED_TO_BOUNDARY_PAIR_EXTERIOR_DEGREE_SEAL"

	FailureNoNativeB2ExteriorFunctional  = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_PAIR_EXTERIOR_RESPONSE_FUNCTIONAL"
	FailureNoDegreeTargetSelection       = "FAILED_ROUTE_NO_DEGREE_TARGET_SELECTION_THEOREM"
	FailureExteriorTruncationNotAlpha    = "FAILED_ROUTE_EXTERIOR_TRUNCATION_DOES_NOT_BY_ITSELF_DERIVE_ALPHA_RESPONSE"
	FailureNoDegreeOneToPiTopMap         = "FAILED_ROUTE_NO_TYPED_DEGREE_ONE_TO_PI_TOP_MAP_CERTIFIED"
	FailureNoDegreeTwoToHRMinMap         = "FAILED_ROUTE_NO_TYPED_DEGREE_TWO_TO_H_R_MIN_MAP_CERTIFIED"
	FailureNoZeroOrderSuppression        = "FAILED_ROUTE_NO_NATIVE_ZERO_ORDER_SUPPRESSION_THEOREM"
	FailureNoLinearHRMinExclusion        = "FAILED_ROUTE_NO_NATIVE_EXCLUSION_OF_LINEAR_H_R_MIN_TERM"
	FailureNoQuadraticPiTopExclusion     = "FAILED_ROUTE_NO_NATIVE_EXCLUSION_OF_QUADRATIC_PI_TOP_TERM"
	FailureNoDegreeFunctor               = "FAILED_ROUTE_NO_SHARED_BOUNDARY_EXTERIOR_DEGREE_FUNCTOR_CERTIFIED"
	FailureNoTruncationResponseTheorem   = "FAILED_ROUTE_NO_NATIVE_TRUNCATION_RESPONSE_THEOREM_FOR_ALPHA_POLYNOMIAL"
	FailureAlphaStillSealed              = "FAILED_ROUTE_ALPHA_B_REMAINS_SEALED"
	FailureNoNativeAlphaSource           = "FAILED_ROUTE_NO_NATIVE_ALPHA_B_SOURCE"
	FailureSocketRankRatiosNotActivation = "FAILED_ROUTE_SOCKET_RANK_RATIOS_NOT_ACTIVATION_THEOREM"
	FailureNoBoundaryAlphaTransport      = "FAILED_ROUTE_NO_BOUNDARY_ALPHA_DOMAIN_TRANSPORT_THEOREM"
	FailureNoOfficialNEffUpdate          = "FAILED_ROUTE_NO_OFFICIAL_N_EFF_UPDATE_ALLOWED"
	FailureNoCYukawaCHiggsUpdate         = "FAILED_ROUTE_NO_C_YUKAWA_OR_C_HIGGS_UPDATE_ALLOWED"
	FailureNoNumericalYukawa             = "FAILED_ROUTE_NO_NUMERICAL_YUKAWA_VALUES_CERTIFIED"
	FailureNoSectorTraceMagnitude        = "FAILED_ROUTE_NO_SECTOR_TRACE_MAGNITUDE_READOUT_MAP"
	FailureNoFullUnbrokenAF              = "FAILED_ROUTE_NO_FULL_UNBROKEN_A_F_THEOREM"
	FailureAForientNotFullAF             = "FAILED_ROUTE_A_F_ORIENT_NOT_FULL_A_F"
	FailureNoNativeFiniteTriple          = "FAILED_ROUTE_NO_NATIVE_FINITE_TRIPLE_THEOREM"
	FailureNotR3                         = "FAILED_ROUTE_NOT_R3_BOUNDARY_PAIR_EXTERIOR_TRUNCATION_OBSTRUCTION"
	FailureNotR4                         = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type Ledger struct {
	AlphaB, SBoundary, OfficialNEff float64
	OfficialFrozen                  bool
	AlphaNative, ExteriorFunctional bool
	R3, R4                          bool
}

type ExteriorDegree struct {
	Degree               int
	Dimension            int
	ResponsePower        string
	CandidateTarget      string
	TargetRank           int
	TargetChamberDim     int
	Contribution         float64
	TargetMapCertified   bool
	SuppressionCertified bool
	Supports, Failures   []string
}

type BoundaryPairExterior struct {
	Carrier                        string
	Dimension                      int
	Lambda0Dim, Lambda1Dim         int
	Lambda2Dim, Lambda3Dim         int
	ExteriorCalculusTyped          bool
	TruncatesAfterSecondDegree     bool
	TruncationDerivesAlphaResponse bool
	Supports, Failures             []string
}

type DegreeTargetAudit struct {
	DegreeOne, DegreeTwo                              ExteriorDegree
	DegreeTargetSelectionTheorem, SharedDegreeFunctor bool
	Supports, Failures                                []string
}

type ZeroAndCrossLaneAudit struct {
	ZeroOrderPresent, ZeroOrderContributes                 bool
	LinearHRMinTermAbsent, QuadraticPiTopTermAbsent        bool
	ZeroOrderSuppressionTheorem, CrossLaneExclusionTheorem bool
	CubicAndHigherAbsent, CubicStopDerivedByLambda3B2Zero  bool
	Supports, Failures                                     []string
}

type AlphaExteriorCandidate struct {
	Expression, DegreeExpression string
	S                            float64
	ReconstructedAlpha           float64
	ShapeCoherent, Native        bool
	ExteriorFunctionalCertified  bool
	Supports, Failures           []string
}

type Obstruction struct {
	ExteriorTruncationWorks           bool
	NativeExteriorFunctionalCertified bool
	DegreeTargetSelectionCertified    bool
	ZeroOrderSuppressionCertified     bool
	AlphaNative                       bool
	RemainingWound, NextGate          string
	Supports, Failures                []string
}

type R3Assessment struct {
	YDaggerYReadoutCarrierReady, SocketMagnitudeTransferTyped bool
	SocketRankAlphaSourceTyped, BoundaryJetShapeTyped         bool
	B2ExteriorTruncationTyped, AlphaNative                    bool
	SectorTraceMagnitudeReadout                               bool
	EligibleForR3, EligibleForR4                              bool
	Supports, Failures                                        []string
}

type Impact struct {
	Classification, Status                                                           string
	B2ExteriorStopShape, NativeExteriorFunctionalSolved, DegreeTargetsSolved         bool
	AlphaNative, SectorTraceReadout                                                  bool
	CanUpdateNEff, CanUpdateCYukawa, CanUpdateCHiggs, CanPromoteToR3, CanPromoteToR4 bool
}

type Firewalls struct {
	Enforced                                                                                        bool
	NoNativeB2ExteriorFunctional, NoDegreeTargetSelection, ExteriorTruncationNotAlpha               bool
	NoDegreeOneToPiTopMap, NoDegreeTwoToHRMinMap, NoZeroOrderSuppression                            bool
	NoLinearHRMinExclusion, NoQuadraticPiTopExclusion, NoDegreeFunctor, NoTruncationResponseTheorem bool
	AlphaStillSealed, NoNativeAlphaSource, SocketRankRatiosNotActivation, NoBoundaryAlphaTransport  bool
	NoOfficialNEffUpdate, NoCYukawaCHiggsUpdate, NoNumericalYukawa, NoSectorTraceMagnitude          bool
	NoFullUnbrokenAF, AForientNotFullAF, NoNativeFiniteTriple, NotR3, NotR4                         bool
	Verdict                                                                                         string
}

type Audit struct {
	ID           string
	Ledger       Ledger
	BoundaryPair BoundaryPairExterior
	Targets      DegreeTargetAudit
	ZeroCross    ZeroAndCrossLaneAudit
	Candidate    AlphaExteriorCandidate
	Obstruction  Obstruction
	R3           R3Assessment
	Impact       Impact
	Firewalls    Firewalls
	Truth, Final string
}

func BuildDefault() (Audit, error) {
	degreeOneContribution := float64(PiTopRank) / float64(H10Dim) * SBoundary
	degreeTwoContribution := float64(HRminRank) / float64(H72Dim) * SBoundary * SBoundary
	a := Audit{
		ID:     AuditID,
		Ledger: Ledger{AlphaB: AlphaB, SBoundary: SBoundary, OfficialNEff: OfficialNEff, OfficialFrozen: true},
		BoundaryPair: BoundaryPairExterior{
			Carrier: "B_2 boundary pair", Dimension: BoundaryPairDim, Lambda0Dim: Lambda0B2Dim, Lambda1Dim: Lambda1B2Dim, Lambda2Dim: Lambda2B2Dim, Lambda3Dim: Lambda3B2Dim,
			ExteriorCalculusTyped: true, TruncatesAfterSecondDegree: true, TruncationDerivesAlphaResponse: false,
			Supports: []string{StatusBoundaryPairAudited, StatusExteriorTruncation, SupportAlphaPowerMatchesB2Exterior, SupportNoCubicFromLambda3B2Zero, SupportExteriorCalculusExplainsStopShape},
			Failures: []string{FailureExteriorTruncationNotAlpha, FailureNoNativeB2ExteriorFunctional, FailureNoTruncationResponseTheorem},
		},
		Targets: DegreeTargetAudit{
			DegreeOne:                    ExteriorDegree{Degree: 1, Dimension: Lambda1B2Dim, ResponsePower: "s", CandidateTarget: "Pi_top=e_+ tensor P_3", TargetRank: PiTopRank, TargetChamberDim: H10Dim, Contribution: degreeOneContribution, TargetMapCertified: false, Supports: []string{StatusDegreeOneTargetAudited, SupportDegreeOneDominantSocketCandidate}, Failures: []string{FailureNoDegreeOneToPiTopMap}},
			DegreeTwo:                    ExteriorDegree{Degree: 2, Dimension: Lambda2B2Dim, ResponsePower: "s^2", CandidateTarget: "H_R^min active punctured right edge-domain", TargetRank: HRminRank, TargetChamberDim: H72Dim, Contribution: degreeTwoContribution, TargetMapCertified: false, Supports: []string{StatusDegreeTwoTargetAudited, SupportDegreeTwoActiveRightCandidate}, Failures: []string{FailureNoDegreeTwoToHRMinMap}},
			DegreeTargetSelectionTheorem: false, SharedDegreeFunctor: false,
			Supports: []string{SupportFirstSecondOrdersHaveB2Candidate, SupportGate866SocketRanksCompatible},
			Failures: []string{FailureNoDegreeTargetSelection, FailureNoDegreeFunctor, FailureNoDegreeOneToPiTopMap, FailureNoDegreeTwoToHRMinMap},
		},
		ZeroCross: ZeroAndCrossLaneAudit{
			ZeroOrderPresent: true, ZeroOrderContributes: false, LinearHRMinTermAbsent: true, QuadraticPiTopTermAbsent: true, ZeroOrderSuppressionTheorem: false, CrossLaneExclusionTheorem: false, CubicAndHigherAbsent: true, CubicStopDerivedByLambda3B2Zero: true,
			Supports: []string{StatusZeroOrderAudited, StatusCrossLaneAudited, SupportNoCubicFromLambda3B2Zero},
			Failures: []string{FailureNoZeroOrderSuppression, FailureNoLinearHRMinExclusion, FailureNoQuadraticPiTopExclusion, FailureNoDegreeTargetSelection},
		},
		Candidate: AlphaExteriorCandidate{
			Expression:       "alpha_B = (1/10)Tr_H10(Pi_top J_B^(1)) + (1/72)Tr_H72(P_HRmin J_B^(2))",
			DegreeExpression: "Lambda^1 B_2 -> s on Pi_top; Lambda^2 B_2 -> s^2 on H_R^min; Lambda^k B_2=0 for k>=3",
			S:                SBoundary, ReconstructedAlpha: degreeOneContribution + degreeTwoContribution, ShapeCoherent: true, Native: false, ExteriorFunctionalCertified: false,
			Supports: []string{StatusGate868Inherited, StatusAlphaReconstructed, SupportAlphaPowerMatchesB2Exterior, SupportBoundaryJetSealSharpened},
			Failures: []string{FailureNoNativeB2ExteriorFunctional, FailureNoDegreeTargetSelection, FailureAlphaStillSealed},
		},
		Obstruction: Obstruction{ExteriorTruncationWorks: true, NativeExteriorFunctionalCertified: false, DegreeTargetSelectionCertified: false, ZeroOrderSuppressionCertified: false, AlphaNative: false, RemainingWound: "BoundaryPairExteriorResponseFunctional and degree-target selection theorem", NextGate: "Gate 870 — BoundaryAlpha ExteriorDegree Response Seal or Native Functional Audit", Supports: []string{StatusNextWound, SupportBoundaryJetSealSharpened}, Failures: []string{FailureNoNativeB2ExteriorFunctional, FailureNoDegreeTargetSelection, FailureNoZeroOrderSuppression, FailureAlphaStillSealed}},
		R3:          R3Assessment{YDaggerYReadoutCarrierReady: true, SocketMagnitudeTransferTyped: true, SocketRankAlphaSourceTyped: true, BoundaryJetShapeTyped: true, B2ExteriorTruncationTyped: true, AlphaNative: false, SectorTraceMagnitudeReadout: false, EligibleForR3: false, EligibleForR4: false, Supports: []string{SupportBoundaryJetSealSharpened}, Failures: []string{FailureAlphaStillSealed, FailureNoSectorTraceMagnitude, FailureNotR3, FailureNotR4}},
		Impact:      Impact{Classification: Classification, Status: R2Status, B2ExteriorStopShape: true, NativeExteriorFunctionalSolved: false, DegreeTargetsSolved: false, AlphaNative: false, SectorTraceReadout: false},
		Firewalls:   Firewalls{Enforced: true, NoNativeB2ExteriorFunctional: true, NoDegreeTargetSelection: true, ExteriorTruncationNotAlpha: true, NoDegreeOneToPiTopMap: true, NoDegreeTwoToHRMinMap: true, NoZeroOrderSuppression: true, NoLinearHRMinExclusion: true, NoQuadraticPiTopExclusion: true, NoDegreeFunctor: true, NoTruncationResponseTheorem: true, AlphaStillSealed: true, NoNativeAlphaSource: true, SocketRankRatiosNotActivation: true, NoBoundaryAlphaTransport: true, NoOfficialNEffUpdate: true, NoCYukawaCHiggsUpdate: true, NoNumericalYukawa: true, NoSectorTraceMagnitude: true, NoFullUnbrokenAF: true, AForientNotFullAF: true, NoNativeFiniteTriple: true, NotR3: true, NotR4: true, Verdict: StatusFirewallVerdict},
		Truth:       "Gate 869 audits Lambda^bullet B_2 as the finite exterior-jet calculus candidate behind the s+s^2 alpha_B shape. Lambda^3 B_2=0 gives a real truncation candidate, but degree-target maps, zero-order suppression, and the native response functional are not certified.",
		Final:       "B_2 exterior calculus explains the possible stop after second degree, but alpha_B remains a sealed boundary response because the degree-one and degree-two targets are not natively selected.",
	}
	if err := a.Validate(); err != nil {
		return Audit{}, err
	}
	return a, nil
}

func (a Audit) Validate() error {
	err := func(msg string) error { return fmt.Errorf("%s: %s", AuditID, msg) }
	if a.ID != AuditID || !a.Ledger.OfficialFrozen || a.Ledger.AlphaNative || a.Ledger.ExteriorFunctional || a.Ledger.R3 || a.Ledger.R4 {
		return err("ledger malformed or overpromoted")
	}
	if a.BoundaryPair.Dimension != 2 || a.BoundaryPair.Lambda3Dim != 0 || !a.BoundaryPair.ExteriorCalculusTyped || !a.BoundaryPair.TruncatesAfterSecondDegree || a.BoundaryPair.TruncationDerivesAlphaResponse {
		return err("boundary pair exterior audit malformed")
	}
	if a.Targets.DegreeOne.Degree != 1 || a.Targets.DegreeOne.TargetRank != PiTopRank || a.Targets.DegreeOne.TargetChamberDim != H10Dim || a.Targets.DegreeOne.TargetMapCertified {
		return err("degree-one target malformed")
	}
	if a.Targets.DegreeTwo.Degree != 2 || a.Targets.DegreeTwo.TargetRank != HRminRank || a.Targets.DegreeTwo.TargetChamberDim != H72Dim || a.Targets.DegreeTwo.TargetMapCertified {
		return err("degree-two target malformed")
	}
	if a.Targets.DegreeTargetSelectionTheorem || a.Targets.SharedDegreeFunctor || !containsAll(a.Targets.Failures, []string{FailureNoDegreeTargetSelection, FailureNoDegreeFunctor}) {
		return err("degree target overpromoted")
	}
	if !a.ZeroCross.ZeroOrderPresent || a.ZeroCross.ZeroOrderContributes || !a.ZeroCross.CubicAndHigherAbsent || !a.ZeroCross.CubicStopDerivedByLambda3B2Zero || a.ZeroCross.ZeroOrderSuppressionTheorem || a.ZeroCross.CrossLaneExclusionTheorem {
		return err("zero/cross-lane audit malformed")
	}
	if !a.Candidate.ShapeCoherent || a.Candidate.Native || a.Candidate.ExteriorFunctionalCertified || !near(a.Candidate.ReconstructedAlpha, AlphaB) {
		return err("alpha exterior candidate malformed")
	}
	if !a.Obstruction.ExteriorTruncationWorks || a.Obstruction.NativeExteriorFunctionalCertified || a.Obstruction.DegreeTargetSelectionCertified || a.Obstruction.ZeroOrderSuppressionCertified || a.Obstruction.AlphaNative {
		return err("obstruction overpromoted")
	}
	if !a.R3.YDaggerYReadoutCarrierReady || !a.R3.SocketMagnitudeTransferTyped || !a.R3.SocketRankAlphaSourceTyped || !a.R3.BoundaryJetShapeTyped || !a.R3.B2ExteriorTruncationTyped || a.R3.AlphaNative || a.R3.SectorTraceMagnitudeReadout || a.R3.EligibleForR3 || a.R3.EligibleForR4 {
		return err("R3 assessment overpromoted")
	}
	if a.Impact.NativeExteriorFunctionalSolved || a.Impact.DegreeTargetsSolved || a.Impact.AlphaNative || a.Impact.SectorTraceReadout || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs || a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 {
		return err("impact overpromoted")
	}
	if !firewallsOK(a.Firewalls) {
		return err("firewalls not enforced")
	}
	return nil
}

func near(a, b float64) bool { return math.Abs(a-b) < 1e-15 }

func containsAll(haystack, needles []string) bool {
	set := make(map[string]bool, len(haystack))
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
	return f.Enforced && f.NoNativeB2ExteriorFunctional && f.NoDegreeTargetSelection && f.ExteriorTruncationNotAlpha && f.NoDegreeOneToPiTopMap && f.NoDegreeTwoToHRMinMap && f.NoZeroOrderSuppression && f.NoLinearHRMinExclusion && f.NoQuadraticPiTopExclusion && f.NoDegreeFunctor && f.NoTruncationResponseTheorem && f.AlphaStillSealed && f.NoNativeAlphaSource && f.SocketRankRatiosNotActivation && f.NoBoundaryAlphaTransport && f.NoOfficialNEffUpdate && f.NoCYukawaCHiggsUpdate && f.NoNumericalYukawa && f.NoSectorTraceMagnitude && f.NoFullUnbrokenAF && f.AForientNotFullAF && f.NoNativeFiniteTriple && f.NotR3 && f.NotR4 && f.Verdict == StatusFirewallVerdict
}

func FormatLedger(l Ledger) string {
	return fmt.Sprintf("ledger alpha_B=%.16g s=%.16g official_N_eff=%.16g frozen=%t alpha_native=%t R3=%t R4=%t", l.AlphaB, l.SBoundary, l.OfficialNEff, l.OfficialFrozen, l.AlphaNative, l.R3, l.R4)
}
func FormatBoundaryPair(b BoundaryPairExterior) string {
	return fmt.Sprintf("%s dim=%d exterior_dims=[%d,%d,%d,%d] lambda3_zero=%t truncates_after_degree2=%t native_alpha=%t supports=%s failures=%s", b.Carrier, b.Dimension, b.Lambda0Dim, b.Lambda1Dim, b.Lambda2Dim, b.Lambda3Dim, b.Lambda3Dim == 0, b.TruncatesAfterSecondDegree, b.TruncationDerivesAlphaResponse, strings.Join(b.Supports, ","), strings.Join(b.Failures, ","))
}
func FormatDegree(d ExteriorDegree) string {
	return fmt.Sprintf("degree=%d dim=%d power=%s target=%s rank=%d chamber=%d contribution=%.16g target_map_certified=%t supports=%s failures=%s", d.Degree, d.Dimension, d.ResponsePower, d.CandidateTarget, d.TargetRank, d.TargetChamberDim, d.Contribution, d.TargetMapCertified, strings.Join(d.Supports, ","), strings.Join(d.Failures, ","))
}
func FormatTargets(t DegreeTargetAudit) string {
	return fmt.Sprintf("degree_targets selection_theorem=%t shared_functor=%t | d1={%s} | d2={%s} supports=%s failures=%s", t.DegreeTargetSelectionTheorem, t.SharedDegreeFunctor, FormatDegree(t.DegreeOne), FormatDegree(t.DegreeTwo), strings.Join(t.Supports, ","), strings.Join(t.Failures, ","))
}
func FormatZeroCross(z ZeroAndCrossLaneAudit) string {
	return fmt.Sprintf("zero_order_present=%t zero_contributes=%t linear_HRmin_absent=%t quadratic_PiTop_absent=%t cubic_and_higher_absent=%t cubic_stop_from_lambda3=%t zero_theorem=%t cross_lane_theorem=%t failures=%s", z.ZeroOrderPresent, z.ZeroOrderContributes, z.LinearHRMinTermAbsent, z.QuadraticPiTopTermAbsent, z.CubicAndHigherAbsent, z.CubicStopDerivedByLambda3B2Zero, z.ZeroOrderSuppressionTheorem, z.CrossLaneExclusionTheorem, strings.Join(z.Failures, ","))
}
func FormatCandidate(c AlphaExteriorCandidate) string {
	return fmt.Sprintf("%s ; %s ; reconstructed_alpha=%.16g shape=%t native=%t functional=%t supports=%s failures=%s", c.Expression, c.DegreeExpression, c.ReconstructedAlpha, c.ShapeCoherent, c.Native, c.ExteriorFunctionalCertified, strings.Join(c.Supports, ","), strings.Join(c.Failures, ","))
}
func FormatObstruction(o Obstruction) string {
	return fmt.Sprintf("exterior_truncation_works=%t native_functional=%t degree_targets=%t zero_suppression=%t alpha_native=%t wound=%s next=%s failures=%s", o.ExteriorTruncationWorks, o.NativeExteriorFunctionalCertified, o.DegreeTargetSelectionCertified, o.ZeroOrderSuppressionCertified, o.AlphaNative, o.RemainingWound, o.NextGate, strings.Join(o.Failures, ","))
}
func FormatR3(r R3Assessment) string {
	return fmt.Sprintf("YdagY=%t socket_transfer=%t socket_rank_alpha=%t boundary_jet_shape=%t B2_exterior=%t alpha_native=%t sector_trace=%t R3=%t R4=%t failures=%s", r.YDaggerYReadoutCarrierReady, r.SocketMagnitudeTransferTyped, r.SocketRankAlphaSourceTyped, r.BoundaryJetShapeTyped, r.B2ExteriorTruncationTyped, r.AlphaNative, r.SectorTraceMagnitudeReadout, r.EligibleForR3, r.EligibleForR4, strings.Join(r.Failures, ","))
}
func FormatImpact(i Impact) string {
	return fmt.Sprintf("classification=%s status=%s B2_stop=%t native_functional=%t degree_targets=%t alpha_native=%t sector_trace=%t update_N_eff=%t update_CYukawa=%t update_CHiggs=%t R3=%t R4=%t", i.Classification, i.Status, i.B2ExteriorStopShape, i.NativeExteriorFunctionalSolved, i.DegreeTargetsSolved, i.AlphaNative, i.SectorTraceReadout, i.CanUpdateNEff, i.CanUpdateCYukawa, i.CanUpdateCHiggs, i.CanPromoteToR3, i.CanPromoteToR4)
}
func FormatFirewalls(f Firewalls) string {
	return fmt.Sprintf("enforced=%t verdict=%s alpha_sealed=%t not_R3=%t not_R4=%t no_degree_targets=%t no_zero_suppression=%t", f.Enforced, f.Verdict, f.AlphaStillSealed, f.NotR3, f.NotR4, f.NoDegreeTargetSelection, f.NoZeroOrderSuppression)
}

func Statuses() []string {
	return []string{
		StatusGate868Inherited, StatusBoundaryPairAudited, StatusExteriorTruncation, StatusDegreeOneTargetAudited, StatusDegreeTwoTargetAudited, StatusZeroOrderAudited, StatusCrossLaneAudited, StatusAlphaReconstructed, StatusFirewallEnforced, StatusLedgerFrozen, StatusNoObservedDataUsed, StatusNextWound, StatusFirewallVerdict,
		SupportAlphaPowerMatchesB2Exterior, SupportFirstSecondOrdersHaveB2Candidate, SupportNoCubicFromLambda3B2Zero, SupportDegreeOneDominantSocketCandidate, SupportDegreeTwoActiveRightCandidate, SupportGate866SocketRanksCompatible, SupportExteriorCalculusExplainsStopShape, SupportBoundaryJetSealSharpened,
		FailureNoNativeB2ExteriorFunctional, FailureNoDegreeTargetSelection, FailureExteriorTruncationNotAlpha, FailureNoDegreeOneToPiTopMap, FailureNoDegreeTwoToHRMinMap, FailureNoZeroOrderSuppression, FailureNoLinearHRMinExclusion, FailureNoQuadraticPiTopExclusion, FailureNoDegreeFunctor, FailureNoTruncationResponseTheorem, FailureAlphaStillSealed, FailureNotR3, FailureNotR4,
	}
}
