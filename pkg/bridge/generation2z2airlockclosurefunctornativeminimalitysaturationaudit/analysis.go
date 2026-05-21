// Package generation2z2airlockclosurefunctornativeminimalitysaturationaudit implements
// Gate 928: Z2 AirlockClosureFunctor Native Minimality and Saturation Audit.
//
// Gate 928 follows Gate 927's closure-incidence factorization of Theta_B^Z2.
// It audits whether the candidate closure ladder Cl(0)=F_0, Cl(1)=F_1,
// Cl(2)=F_2 is forced by basepoint, monotonicity, minimality, saturation,
// fixed-base quotienting, and Z2 representative-independence. The result is a
// uniqueness-under-axioms audit only: the closure axioms themselves remain
// bridge-level and are not promoted to native ASHA theorems.
package generation2z2airlockclosurefunctornativeminimalitysaturationaudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE928-Z2-AIRLOCK-CLOSURE-FUNCTOR-NATIVE-MINIMALITY-SATURATION-AUDIT"

	Gate927ShortStatus = "R3_ALPHA_TARGET_FUNCTOR_REDUCED_TO_AIRLOCK_CLOSURE_THEOREM"

	BoundarySubsetChain = "0<1<2 boundary subset cardinality chain"
	ExteriorSourceChain = "Lambda^0 B_2=basepoint; Lambda^1 B_2=one-boundary exposure; Lambda^2 B_2=full boundary-pair enclosure"
	AirlockFlagChain    = "F_0 subset F_1 subset F_2"
	ClosureFunctor      = "Cl_{B->A}^{Z2}"
	ClosureZero         = "Cl(0)=F_0"
	ClosureOne          = "Cl(1)=F_1"
	ClosureTwo          = "Cl(2)=F_2"
	ThetaViaClosure     = "Theta_B^Z2(k)=[Cl_{B->A}^{Z2}(k)/F_0]_{Z2}"
	ThetaOne            = "Theta_B^Z2(1)=[F_1/F_0]_{Z2}"
	ThetaTwo            = "Theta_B^Z2(2)=[F_2/F_0]_{Z2}"
	AlphaViaClosure     = "mu_B(R_B(S_split))=(3/10)S_split+(7/72)S_split^2"
	NextGate            = "NEXT_PRESSURE_GATE929_AIRLOCKCLOSURE_AXIOM_SOURCE_AND_FLAG_GENERATED_MINIMALITY_AUDIT"

	RankF1OverF0 = 3
	RankF2OverF0 = 7
	RankH10      = 10
	RankH72      = 72

	Classification = "R3_AIRLOCK_CLOSURE_UNIQUENESS_CANDIDATE_NATIVE_AXIOM_SOURCE_MISSING"
	ShortStatus    = "R3_ALPHA_CLOSURE_UNIQUE_UNDER_AXIOMS_NATIVE_SOURCE_MISSING"
	FinalTruth     = "Z2_AIRLOCK_CLOSURE_UNIQUE_UNDER_MINIMALITY_MONOTONICITY_SATURATION_AXIOMS_BUT_NATIVE_AXIOM_SOURCE_MISSING"

	StatusInheritedGate927         = "PASS_GATE927_CLOSURE_INCIDENCE_FACTORIZATION_INHERITED"
	StatusBasepointAxiom           = "PASS_BASEPOINT_AXIOM_FORCES_CL_0_EQUALS_F0_UNDER_AXIOM"
	StatusMonotonicityAxiom        = "PASS_CLOSURE_MONOTONICITY_COMPATIBLE_WITH_BOUNDARY_AND_AIRLOCK_CHAINS"
	StatusMinimalNontrivialClosure = "PASS_MINIMAL_NONTRIVIAL_CLOSURE_FORCES_CL_1_EQUALS_F1_UNDER_AXIOM"
	StatusSaturationClosure        = "PASS_SATURATION_AXIOM_FORCES_CL_2_EQUALS_F2_UNDER_AXIOM"
	StatusZ2ClosureCompatibility   = "PASS_Z2_REPRESENTATIVE_INDEPENDENCE_OF_CLOSURE_LADDER"
	StatusFixedBaseQuotient        = "PASS_FIXED_BASE_QUOTIENT_FORCES_CUMULATIVE_TARGETS_UNDER_RULE"
	StatusUniqueClosureUnderAxioms = "PASS_AIRLOCK_CLOSURE_UNIQUE_UNDER_BASEPOINT_MONOTONE_MINIMAL_SATURATED_Z2_AXIOMS"
	StatusMeasureReconstruction    = "PASS_BOUNDARY_ACTIVATION_MEASURE_TARGETS_FIXED_BY_UNIQUE_CLOSURE_CANDIDATE"
	StatusNativeAxiomSourceMissing = "FIREWALL_PRESERVED_CLOSURE_AXIOMS_NOT_NATIVE_ASHA_THEOREMS"
	StatusAlphaR3StillBlocked      = "FIREWALL_PRESERVED_ALPHA_AND_R3_NOT_NATIVE"

	SupportBasepointForcesCl0             = "CONDITIONAL_SUPPORT_BASEPOINT_AXIOM_FORCES_CL_0_EQUALS_F0"
	SupportEmptySubsetNoActiveResponse    = "CONDITIONAL_SUPPORT_EMPTY_BOUNDARY_SUBSET_HAS_NO_ACTIVE_RESPONSE"
	SupportReducedMatchesBasepointClosure = "CONDITIONAL_SUPPORT_REDUCED_RESPONSE_MATCHES_BASEPOINT_CLOSURE"
	SupportClosureMustBeMonotone          = "CONDITIONAL_SUPPORT_CLOSURE_MUST_BE_MONOTONE_FROM_BOUNDARY_CHAIN_TO_AIRLOCK_FLAG"
	SupportAirlockFlagThreeLevels         = "CONDITIONAL_SUPPORT_AIRLOCK_FLAG_HAS_MATCHING_MONOTONE_THREE_LEVEL_STRUCTURE"
	SupportMinimalForcesCl1               = "CONDITIONAL_SUPPORT_MINIMAL_NONTRIVIAL_CLOSURE_FORCES_CL_1_EQUALS_F1"
	SupportSingletonCannotSkipF2          = "CONDITIONAL_SUPPORT_SINGLETON_BOUNDARY_ACTIVATION_CANNOT_SKIP_TO_SATURATED_F2_UNDER_MINIMALITY"
	SupportExposureFirstNonbase           = "CONDITIONAL_SUPPORT_EXPOSURE_LEVEL_IS_FIRST_NONBASE_AIRLOCK_CLOSURE"
	SupportSaturationForcesCl2            = "CONDITIONAL_SUPPORT_SATURATION_AXIOM_FORCES_CL_2_EQUALS_F2"
	SupportTopDegreeClosesF2              = "CONDITIONAL_SUPPORT_TOP_BOUNDARY_DEGREE_CLOSES_TO_FULL_AIRLOCK_RECTANGLE"
	SupportFullPairCannotRemainF1         = "CONDITIONAL_SUPPORT_FULL_PAIR_ACTIVATION_CANNOT_REMAIN_AT_F1_UNDER_SATURATION"
	SupportClosureZ2Independent           = "CONDITIONAL_SUPPORT_CLOSURE_LADDER_IS_Z2_REPRESENTATIVE_INDEPENDENT"
	SupportPhaseFlipCommutesClosure       = "CONDITIONAL_SUPPORT_GLOBAL_PHASE_FLIP_COMMUTES_WITH_CLOSURE_LEVELS"
	SupportClosureRanksZ2Invariant        = "CONDITIONAL_SUPPORT_CLOSURE_TARGET_RANKS_ARE_Z2_INVARIANT"
	SupportFixedBaseForcesCumulative      = "CONDITIONAL_SUPPORT_FIXED_BASE_QUOTIENT_RULE_FORCES_CUMULATIVE_TARGETS"
	SupportF2OverF0FromPunctureBase       = "CONDITIONAL_SUPPORT_F2_OVER_F0_FOLLOWS_FROM_CLOSURE_OVER_PUNCTURE_BASE"
	SupportGradedRejectedByFixedBase      = "CONDITIONAL_SUPPORT_ASSOCIATED_GRADED_F2_OVER_F1_REJECTED_BY_FIXED_BASE_QUOTIENTING"
	SupportClosureUniqueUnderAxioms       = "CONDITIONAL_SUPPORT_AIRLOCK_CLOSURE_IS_UNIQUE_UNDER_BASEPOINT_MONOTONE_MINIMAL_SATURATED_Z2_AXIOMS"
	SupportCl1Cl2Forced                   = "CONDITIONAL_SUPPORT_CL_1_EQUALS_F1_AND_CL_2_EQUALS_F2_ARE_FORCED_BY_CLOSURE_AXIOMS"
	SupportAlternativesFail               = "CONDITIONAL_SUPPORT_ALTERNATIVE_CLOSURES_FAIL_MINIMALITY_SATURATION_OR_CLOSURE_LEVEL_TYPING"
	SupportUniqueClosureSuppliesTheta     = "CONDITIONAL_SUPPORT_UNIQUE_CLOSURE_SUPPLIES_THETA_B_Z2"
	SupportMeasureTargetsFixed            = "CONDITIONAL_SUPPORT_BOUNDARY_ACTIVATION_MEASURE_TARGETS_FIXED_BY_UNIQUE_AIRLOCK_CLOSURE"
	SupportAlphaViaUniqueClosure          = "CONDITIONAL_SUPPORT_ALPHA_B_RECONSTRUCTED_THROUGH_UNIQUE_CLOSURE_CANDIDATE"

	FailureBasepointAxiomNotNative        = "FAILED_ROUTE_BASEPOINT_AXIOM_NOT_NATIVE_AIRLOCK_THEOREM"
	FailureMonotonicityNotNative          = "FAILED_ROUTE_MONOTONICITY_NOT_NATIVE_BOUNDARY_AIRLOCK_ACTION_THEOREM"
	FailureMinimalClosureNotNative        = "FAILED_ROUTE_MINIMAL_NONTRIVIAL_CLOSURE_AXIOM_NOT_NATIVE_CERTIFIED"
	FailureSaturationNotNative            = "FAILED_ROUTE_SATURATION_AXIOM_NOT_NATIVE_CERTIFIED"
	FailureZ2ClosureNotNative             = "FAILED_ROUTE_Z2_EQUIVARIANT_CLOSURE_NOT_NATIVE_GLOBAL_PHASE_THEOREM"
	FailureFixedBaseQuotientNotNative     = "FAILED_ROUTE_FIXED_BASE_QUOTIENT_RULE_NOT_NATIVE_CERTIFIED"
	FailureClosureAxiomsNotNative         = "FAILED_ROUTE_CLOSURE_AXIOMS_NOT_NATIVE_ASHA_THEOREMS"
	FailureAlphaViaClosureNotNative       = "FAILED_ROUTE_ALPHA_RECONSTRUCTION_THROUGH_UNIQUE_CLOSURE_NOT_NATIVE_ALPHA_THEOREM"
	FailureMuBNotNativeWithoutAxiomSource = "FAILED_ROUTE_MU_B_STILL_NOT_NATIVE_WITHOUT_NATIVE_CLOSURE_AXIOM_SOURCE"
	FailureAlphaBridgeCandidateNotNative  = "FAILED_ROUTE_ALPHA_B_REMAINS_BRIDGE_CANDIDATE_NOT_NATIVE"
	FailureNotNativeR3                    = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked      = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap         = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap         = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues       = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator         = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type BasepointAxiomAudit struct {
	SubsetCardinality    int
	ClosureTarget        string
	ForcedByAxiom        bool
	EmptyHasNoActivation bool
	MatchesReducedForm   bool
	NativeAxiomCertified bool
	Supports             []string
	Failures             []string
}

type MonotonicityAudit struct {
	SourceChain         string
	TargetChain         string
	Monotone            bool
	MatchingThreeLevels bool
	NativeActionTheorem bool
	Supports            []string
	Failures            []string
}

type MinimalClosureAudit struct {
	SingletonLevel       int
	ClosureTarget        string
	ForcedByMinimality   bool
	SkipsSaturation      bool
	FirstNonbaseClosure  bool
	NativeAxiomCertified bool
	Supports             []string
	Failures             []string
}

type SaturationAudit struct {
	FullPairLevel        int
	ClosureTarget        string
	ForcedBySaturation   bool
	TopBoundaryDegree    bool
	CannotRemainAtF1     bool
	NativeAxiomCertified bool
	Supports             []string
	Failures             []string
}

type Z2ClosureAudit struct {
	LambdaLadder       string
	BarLambdaLadder    string
	PhaseFlipCommutes  bool
	RanksInvariant     bool
	RepresentativeFree bool
	NativeZ2Theorem    bool
	Supports           []string
	Failures           []string
}

type FixedBaseQuotientAudit struct {
	ThetaFormula        string
	TopTarget           string
	RejectedGraded      string
	ForcesCumulative    bool
	RejectsAssociated   bool
	NativeRuleCertified bool
	Supports            []string
	Failures            []string
}

type ClosureUniquenessAudit struct {
	Cl0                  string
	Cl1                  string
	Cl2                  string
	Basepoint            bool
	Monotone             bool
	Minimal              bool
	Saturated            bool
	Z2Invariant          bool
	UniqueUnderAxioms    bool
	NativeAxiomSource    bool
	RejectedAlternatives []string
	Supports             []string
	Failures             []string
}

type MeasureConsequenceAudit struct {
	ThetaOneRank               int
	ThetaTwoRank               int
	H10Rank                    int
	H72Rank                    int
	AlphaFormula               string
	UniqueClosureSuppliesTheta bool
	TargetsFixedByClosure      bool
	AlphaReconstructed         bool
	NativeAlphaTheorem         bool
	Supports                   []string
	Failures                   []string
}

type FirewallLedger struct {
	BasepointAxiomNotNative        bool
	MonotonicityNotNative          bool
	MinimalClosureNotNative        bool
	SaturationNotNative            bool
	Z2ClosureNotNative             bool
	FixedBaseQuotientNotNative     bool
	ClosureAxiomsNotNative         bool
	AlphaViaClosureNotNative       bool
	MuBNotNativeWithoutAxiomSource bool
	AlphaBridgeCandidateNotNative  bool
	NotNativeR3                    bool
	FullAFDescentStillBlocked      bool
	NoGenerationCarrierMap         bool
	NoFlavorOrientationMap         bool
	NoIndividualYukawaValues       bool
	NoNativeYukawaOperator         bool
}

type Audit struct {
	ID             string
	Inherited      string
	Basepoint      BasepointAxiomAudit
	Monotonicity   MonotonicityAudit
	Minimality     MinimalClosureAudit
	Saturation     SaturationAudit
	Z2             Z2ClosureAudit
	FixedBase      FixedBaseQuotientAudit
	Uniqueness     ClosureUniquenessAudit
	Measure        MeasureConsequenceAudit
	Firewalls      FirewallLedger
	Truth          string
	Classification string
	ShortStatus    string
	Final          string
}

func BuildDefault() (Audit, error) {
	return Audit{
		ID:             AuditID,
		Inherited:      Gate927ShortStatus,
		Basepoint:      BasepointAxiomAudit{SubsetCardinality: 0, ClosureTarget: "F_0", ForcedByAxiom: true, EmptyHasNoActivation: true, MatchesReducedForm: true, NativeAxiomCertified: false, Supports: []string{SupportBasepointForcesCl0, SupportEmptySubsetNoActiveResponse, SupportReducedMatchesBasepointClosure}, Failures: []string{FailureBasepointAxiomNotNative}},
		Monotonicity:   MonotonicityAudit{SourceChain: BoundarySubsetChain, TargetChain: AirlockFlagChain, Monotone: true, MatchingThreeLevels: true, NativeActionTheorem: false, Supports: []string{SupportClosureMustBeMonotone, SupportAirlockFlagThreeLevels}, Failures: []string{FailureMonotonicityNotNative}},
		Minimality:     MinimalClosureAudit{SingletonLevel: 1, ClosureTarget: "F_1", ForcedByMinimality: true, SkipsSaturation: false, FirstNonbaseClosure: true, NativeAxiomCertified: false, Supports: []string{SupportMinimalForcesCl1, SupportSingletonCannotSkipF2, SupportExposureFirstNonbase}, Failures: []string{FailureMinimalClosureNotNative}},
		Saturation:     SaturationAudit{FullPairLevel: 2, ClosureTarget: "F_2", ForcedBySaturation: true, TopBoundaryDegree: true, CannotRemainAtF1: true, NativeAxiomCertified: false, Supports: []string{SupportSaturationForcesCl2, SupportTopDegreeClosesF2, SupportFullPairCannotRemainF1}, Failures: []string{FailureSaturationNotNative}},
		Z2:             Z2ClosureAudit{LambdaLadder: "F_0^lambda subset F_1^lambda subset F_2", BarLambdaLadder: "F_0^barlambda subset F_1^barlambda subset F_2", PhaseFlipCommutes: true, RanksInvariant: true, RepresentativeFree: true, NativeZ2Theorem: false, Supports: []string{SupportClosureZ2Independent, SupportPhaseFlipCommutesClosure, SupportClosureRanksZ2Invariant}, Failures: []string{FailureZ2ClosureNotNative}},
		FixedBase:      FixedBaseQuotientAudit{ThetaFormula: ThetaViaClosure, TopTarget: "[F_2/F_0]_{Z2}", RejectedGraded: "F_2/F_1", ForcesCumulative: true, RejectsAssociated: true, NativeRuleCertified: false, Supports: []string{SupportFixedBaseForcesCumulative, SupportF2OverF0FromPunctureBase, SupportGradedRejectedByFixedBase}, Failures: []string{FailureFixedBaseQuotientNotNative}},
		Uniqueness:     ClosureUniquenessAudit{Cl0: ClosureZero, Cl1: ClosureOne, Cl2: ClosureTwo, Basepoint: true, Monotone: true, Minimal: true, Saturated: true, Z2Invariant: true, UniqueUnderAxioms: true, NativeAxiomSource: false, RejectedAlternatives: []string{"Cl(1)=F_2 fails minimality", "Cl(2)=F_1 fails saturation", "Cl(2)=F_2/F_1 is not a closure level"}, Supports: []string{SupportClosureUniqueUnderAxioms, SupportCl1Cl2Forced, SupportAlternativesFail}, Failures: []string{FailureClosureAxiomsNotNative}},
		Measure:        MeasureConsequenceAudit{ThetaOneRank: RankF1OverF0, ThetaTwoRank: RankF2OverF0, H10Rank: RankH10, H72Rank: RankH72, AlphaFormula: AlphaViaClosure, UniqueClosureSuppliesTheta: true, TargetsFixedByClosure: true, AlphaReconstructed: true, NativeAlphaTheorem: false, Supports: []string{SupportUniqueClosureSuppliesTheta, SupportMeasureTargetsFixed, SupportAlphaViaUniqueClosure}, Failures: []string{FailureAlphaViaClosureNotNative, FailureMuBNotNativeWithoutAxiomSource, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3}},
		Firewalls:      FirewallLedger{BasepointAxiomNotNative: true, MonotonicityNotNative: true, MinimalClosureNotNative: true, SaturationNotNative: true, Z2ClosureNotNative: true, FixedBaseQuotientNotNative: true, ClosureAxiomsNotNative: true, AlphaViaClosureNotNative: true, MuBNotNativeWithoutAxiomSource: true, AlphaBridgeCandidateNotNative: true, NotNativeR3: true, FullAFDescentStillBlocked: true, NoGenerationCarrierMap: true, NoFlavorOrientationMap: true, NoIndividualYukawaValues: true, NoNativeYukawaOperator: true},
		Truth:          FinalTruth,
		Classification: Classification,
		ShortStatus:    ShortStatus,
		Final:          "Gate 928 proves uniqueness under closure axioms, not native origin of those axioms.",
	}, nil
}

func Statuses() []string {
	return []string{StatusInheritedGate927, StatusBasepointAxiom, StatusMonotonicityAxiom, StatusMinimalNontrivialClosure, StatusSaturationClosure, StatusZ2ClosureCompatibility, StatusFixedBaseQuotient, StatusUniqueClosureUnderAxioms, StatusMeasureReconstruction, StatusNativeAxiomSourceMissing, StatusAlphaR3StillBlocked}
}
func Supports() []string {
	return []string{SupportBasepointForcesCl0, SupportEmptySubsetNoActiveResponse, SupportReducedMatchesBasepointClosure, SupportClosureMustBeMonotone, SupportAirlockFlagThreeLevels, SupportMinimalForcesCl1, SupportSingletonCannotSkipF2, SupportExposureFirstNonbase, SupportSaturationForcesCl2, SupportTopDegreeClosesF2, SupportFullPairCannotRemainF1, SupportClosureZ2Independent, SupportPhaseFlipCommutesClosure, SupportClosureRanksZ2Invariant, SupportFixedBaseForcesCumulative, SupportF2OverF0FromPunctureBase, SupportGradedRejectedByFixedBase, SupportClosureUniqueUnderAxioms, SupportCl1Cl2Forced, SupportAlternativesFail, SupportUniqueClosureSuppliesTheta, SupportMeasureTargetsFixed, SupportAlphaViaUniqueClosure}
}
func Failures() []string {
	return []string{FailureBasepointAxiomNotNative, FailureMonotonicityNotNative, FailureMinimalClosureNotNative, FailureSaturationNotNative, FailureZ2ClosureNotNative, FailureFixedBaseQuotientNotNative, FailureClosureAxiomsNotNative, FailureAlphaViaClosureNotNative, FailureMuBNotNativeWithoutAxiomSource, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}
}

func (f FirewallLedger) List() []string {
	out := []string{}
	if f.BasepointAxiomNotNative {
		out = append(out, FailureBasepointAxiomNotNative)
	}
	if f.MonotonicityNotNative {
		out = append(out, FailureMonotonicityNotNative)
	}
	if f.MinimalClosureNotNative {
		out = append(out, FailureMinimalClosureNotNative)
	}
	if f.SaturationNotNative {
		out = append(out, FailureSaturationNotNative)
	}
	if f.Z2ClosureNotNative {
		out = append(out, FailureZ2ClosureNotNative)
	}
	if f.FixedBaseQuotientNotNative {
		out = append(out, FailureFixedBaseQuotientNotNative)
	}
	if f.ClosureAxiomsNotNative {
		out = append(out, FailureClosureAxiomsNotNative)
	}
	if f.AlphaViaClosureNotNative {
		out = append(out, FailureAlphaViaClosureNotNative)
	}
	if f.MuBNotNativeWithoutAxiomSource {
		out = append(out, FailureMuBNotNativeWithoutAxiomSource)
	}
	if f.AlphaBridgeCandidateNotNative {
		out = append(out, FailureAlphaBridgeCandidateNotNative)
	}
	if f.NotNativeR3 {
		out = append(out, FailureNotNativeR3)
	}
	if f.FullAFDescentStillBlocked {
		out = append(out, FailureFullAFDescentStillBlocked)
	}
	if f.NoGenerationCarrierMap {
		out = append(out, FailureNoGenerationCarrierMap)
	}
	if f.NoFlavorOrientationMap {
		out = append(out, FailureNoFlavorOrientationMap)
	}
	if f.NoIndividualYukawaValues {
		out = append(out, FailureNoIndividualYukawaValues)
	}
	if f.NoNativeYukawaOperator {
		out = append(out, FailureNoNativeYukawaOperator)
	}
	return out
}

func FormatBasepoint(a BasepointAxiomAudit) string {
	return fmt.Sprintf("level=%d target=%s forced=%t empty_no_activation=%t reduced_match=%t native_axiom=%t supports=%s failures=%s", a.SubsetCardinality, a.ClosureTarget, a.ForcedByAxiom, a.EmptyHasNoActivation, a.MatchesReducedForm, a.NativeAxiomCertified, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatMonotonicity(a MonotonicityAudit) string {
	return fmt.Sprintf("source=%s target=%s monotone=%t matching_three_levels=%t native_action=%t supports=%s failures=%s", a.SourceChain, a.TargetChain, a.Monotone, a.MatchingThreeLevels, a.NativeActionTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatMinimality(a MinimalClosureAudit) string {
	return fmt.Sprintf("level=%d target=%s forced=%t skips_saturation=%t first_nonbase=%t native_axiom=%t supports=%s failures=%s", a.SingletonLevel, a.ClosureTarget, a.ForcedByMinimality, a.SkipsSaturation, a.FirstNonbaseClosure, a.NativeAxiomCertified, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatSaturation(a SaturationAudit) string {
	return fmt.Sprintf("level=%d target=%s forced=%t top_degree=%t cannot_remain_f1=%t native_axiom=%t supports=%s failures=%s", a.FullPairLevel, a.ClosureTarget, a.ForcedBySaturation, a.TopBoundaryDegree, a.CannotRemainAtF1, a.NativeAxiomCertified, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatZ2(a Z2ClosureAudit) string {
	return fmt.Sprintf("lambda=%s barlambda=%s phase_commutes=%t ranks_invariant=%t representative_free=%t native_z2=%t supports=%s failures=%s", a.LambdaLadder, a.BarLambdaLadder, a.PhaseFlipCommutes, a.RanksInvariant, a.RepresentativeFree, a.NativeZ2Theorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatFixedBase(a FixedBaseQuotientAudit) string {
	return fmt.Sprintf("theta=%s top=%s rejected=%s cumulative=%t rejects_associated=%t native_rule=%t supports=%s failures=%s", a.ThetaFormula, a.TopTarget, a.RejectedGraded, a.ForcesCumulative, a.RejectsAssociated, a.NativeRuleCertified, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatUniqueness(a ClosureUniquenessAudit) string {
	return fmt.Sprintf("cl0=%s cl1=%s cl2=%s base=%t monotone=%t minimal=%t saturated=%t z2=%t unique=%t native_axioms=%t rejected=%s supports=%s failures=%s", a.Cl0, a.Cl1, a.Cl2, a.Basepoint, a.Monotone, a.Minimal, a.Saturated, a.Z2Invariant, a.UniqueUnderAxioms, a.NativeAxiomSource, strings.Join(a.RejectedAlternatives, ";"), strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatMeasure(a MeasureConsequenceAudit) string {
	return fmt.Sprintf("rank1=%d rank2=%d h10=%d h72=%d alpha=%s supplies_theta=%t targets_fixed=%t alpha_reconstructed=%t native_alpha=%t supports=%s failures=%s", a.ThetaOneRank, a.ThetaTwoRank, a.H10Rank, a.H72Rank, a.AlphaFormula, a.UniqueClosureSuppliesTheta, a.TargetsFixedByClosure, a.AlphaReconstructed, a.NativeAlphaTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}
func FormatFirewalls(f FirewallLedger) string { return strings.Join(f.List(), ",") }

func firewallsOK(f FirewallLedger) bool { return containsAll(f.List(), Failures()) }
func containsAll(got, want []string) bool {
	m := map[string]bool{}
	for _, g := range got {
		m[g] = true
	}
	for _, w := range want {
		if !m[w] {
			return false
		}
	}
	return true
}
