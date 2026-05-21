// Package generation2boundarysubsetairlockclosureincidencefunctoraudit implements
// Gate 927: BoundarySubset AirlockClosure IncidenceFunctor Audit.
//
// Gate 927 follows Gate 926's result that Theta_B^Z2 is the unique natural
// target functor under current constraints, but still lacks a native source.
// This gate factors Theta_B^Z2 through a finite boundary-subset/airlock-closure
// incidence candidate: exterior degree -> activated subset cardinality ->
// airlock closure level -> Z2 quotient over F_0. The factorization sharpens the
// wound to a candidate Z2 AirlockClosureFunctor while preserving that no native
// closure theorem is certified.
package generation2boundarysubsetairlockclosureincidencefunctoraudit

import (
	"fmt"
	"strings"
)

const (
	AuditID = "GATE927-BOUNDARY-SUBSET-AIRLOCK-CLOSURE-INCIDENCE-FUNCTOR-AUDIT"

	Gate926ShortStatus = "R3_ALPHA_TARGET_FUNCTOR_UNIQUE_UNDER_CONSTRAINTS_NATIVE_SOURCE_MISSING"

	BoundaryPair           = "B_2=<b1,b2>"
	BoundarySubsetLattice  = "nonempty subsets {b1},{b2},{b1,b2}; cardinality classes |S|=1,2"
	SourceDegreeChain      = "deg(Lambda^1 B_2)<deg(Lambda^2 B_2)"
	CardinalityChain       = "|S|=1<|S|=2"
	AirlockFlagChain       = "F_0 subset F_1 subset F_2"
	Z2PunctureClass        = "[p]_{Z2}={e_lambda tensor P_1,e_barlambda tensor P_1}"
	ClosureFunctor         = "Cl_{B->A}^{Z2}:{|S|=0,1,2}->{F_0,F_1,F_2}_{Z2}"
	ClosureZero            = "Cl_{B->A}^{Z2}(0)=F_0"
	ClosureOne             = "Cl_{B->A}^{Z2}(1)=F_1"
	ClosureTwo             = "Cl_{B->A}^{Z2}(2)=F_2"
	ThetaViaClosure        = "Theta_B^Z2(k)=[Cl_{B->A}^{Z2}(k)/F_0]_{Z2}"
	ThetaOne               = "Theta_B^Z2(1)=[F_1/F_0]_{Z2}"
	ThetaTwo               = "Theta_B^Z2(2)=[F_2/F_0]_{Z2}"
	MeasureViaClosure      = "mu_B(R_B(S_split))=sum_{k=1}^{2} rank([Cl_{B->A}^{Z2}(k)/F_0]_{Z2})/rank(H_k)*S_split^k"
	AlphaFormula           = "alpha_B=(3/10)S_split+(7/72)S_split^2"
	NextGate               = "NEXT_PRESSURE_GATE928_Z2_AIRLOCKCLOSUREFUNCTOR_NATIVE_MINIMALITY_AND_SATURATION_AUDIT"
	AssociatedGradedTarget = "F_2/F_1"

	RankF1OverF0 = 3
	RankF2OverF0 = 7
	RankF2OverF1 = 4
	RankH10      = 10
	RankH72      = 72

	Classification = "R3_THETA_B_Z2_INCIDENT_CLOSURE_FUNCTOR_CANDIDATE_NOT_NATIVE"
	ShortStatus    = "R3_ALPHA_TARGET_FUNCTOR_REDUCED_TO_AIRLOCK_CLOSURE_THEOREM"
	FinalTruth     = "BOUNDARY_DEGREE_TO_AIRLOCK_FLAG_TARGET_FUNCTOR_FACTORS_THROUGH_CLOSURE_INCIDENT_FUNCTOR_BUT_NATIVE_CLOSURE_THEOREM_MISSING"

	StatusInheritedGate926             = "PASS_GATE926_THETA_B_Z2_UNIQUENESS_UNDER_CONSTRAINTS_INHERITED"
	StatusSubsetLatticeSource          = "PASS_BOUNDARY_EXTERIOR_DEGREE_FACTORS_THROUGH_SUBSET_CARDINALITY"
	StatusAirlockClosureLadderTyped    = "PASS_AIRLOCK_FLAG_TYPED_AS_CLOSURE_LADDER_CANDIDATE"
	StatusThetaFactorsThroughClosure   = "PASS_THETA_B_Z2_FACTORS_THROUGH_AIRLOCK_CLOSURE_CANDIDATE"
	StatusBasepointClosure             = "PASS_EMPTY_BOUNDARY_SUBSET_MAPS_TO_PUNCTURE_BASEPOINT"
	StatusSingletonClosure             = "PASS_SINGLETON_BOUNDARY_ACTIVATION_CLOSES_TO_F1_CANDIDATE"
	StatusFullPairClosure              = "PASS_FULL_BOUNDARY_PAIR_ACTIVATION_CLOSES_TO_F2_CANDIDATE"
	StatusCumulativeQuotientFromBase   = "PASS_CUMULATIVE_QUOTIENT_F2_OVER_F0_FROM_FIXED_BASEPOINT_QUOTIENT"
	StatusClosureUniqueUnderRules      = "PASS_CLOSURE_FUNCTOR_UNIQUE_UNDER_MONOTONE_MINIMAL_SATURATED_RULES"
	StatusClosureMeasureReconstruction = "PASS_BOUNDARY_ACTIVATION_MEASURE_REWRITTEN_USING_AIRLOCK_CLOSURE"
	StatusNativeClosureMissing         = "FIREWALL_PRESERVED_NATIVE_AIRLOCK_CLOSURE_FUNCTOR_MISSING"
	StatusAlphaR3StillBlocked          = "FIREWALL_PRESERVED_ALPHA_AND_R3_NOT_NATIVE"

	SupportExteriorDegreeEqualsSubsetCardinality = "CONDITIONAL_SUPPORT_EXTERIOR_DEGREE_EQUALS_BOUNDARY_SUBSET_CARDINALITY_FOR_B2"
	SupportLambda1SingletonSubsets               = "CONDITIONAL_SUPPORT_LAMBDA1B2_CORRESPONDS_TO_SINGLETON_BOUNDARY_SUBSETS"
	SupportLambda2FullPairSubset                 = "CONDITIONAL_SUPPORT_LAMBDA2B2_CORRESPONDS_TO_FULL_BOUNDARY_PAIR_SUBSET"
	SupportSourceChainSubsetLattice              = "CONDITIONAL_SUPPORT_SOURCE_CHAIN_HAS_NATIVE_FINITE_SUBSET_LATTICE_SOURCE"
	SupportAirlockFlagClosureLadder              = "CONDITIONAL_SUPPORT_AIRLOCK_FLAG_HAS_CLOSURE_LADDER_TYPE"
	SupportF0PunctureBasepoint                   = "CONDITIONAL_SUPPORT_F0_IS_PUNCTURE_BASEPOINT"
	SupportF1MinimalExposedClosure               = "CONDITIONAL_SUPPORT_F1_IS_MINIMAL_EXPOSED_CLOSURE_ABOVE_PUNCTURE"
	SupportF2SaturatedFullRightRectangle         = "CONDITIONAL_SUPPORT_F2_IS_SATURATED_FULL_RIGHT_RECTANGLE_CLOSURE"
	SupportThetaFactorsThroughClosure            = "CONDITIONAL_SUPPORT_THETA_B_Z2_FACTORS_THROUGH_AIRLOCK_CLOSURE_FUNCTOR"
	SupportDegreeToFlagHasClosureSource          = "CONDITIONAL_SUPPORT_DEGREE_TO_FLAG_TARGET_MAP_HAS_INCIDENT_CLOSURE_SOURCE_CANDIDATE"
	SupportClosureQuotientReconstructsTheta      = "CONDITIONAL_SUPPORT_CLOSURE_QUOTIENT_RECONSTRUCTS_THETA_B_Z2"
	SupportEmptySubsetMapsF0                     = "CONDITIONAL_SUPPORT_EMPTY_BOUNDARY_SUBSET_MAPS_TO_PUNCTURE_BASEPOINT"
	SupportBasepointClosureMatchesReduction      = "CONDITIONAL_SUPPORT_BASEPOINT_CLOSURE_MATCHES_REDUCED_RESPONSE"
	SupportDegreeZeroNoActiveAlpha               = "CONDITIONAL_SUPPORT_DEGREE_ZERO_HAS_NO_ACTIVE_ALPHA_QUOTIENT"
	SupportSingletonClosesF1                     = "CONDITIONAL_SUPPORT_SINGLETON_BOUNDARY_ACTIVATION_CLOSES_TO_F1"
	SupportOneBoundaryGeneratesF1                = "CONDITIONAL_SUPPORT_ONE_BOUNDARY_FACTOR_GENERATES_MINIMAL_EXPOSED_AIRLOCK_FACE"
	SupportThetaOneFromSingletonClosure          = "CONDITIONAL_SUPPORT_THETA_ONE_FOLLOWS_FROM_SINGLETON_CLOSURE"
	SupportFullPairClosesF2                      = "CONDITIONAL_SUPPORT_FULL_BOUNDARY_PAIR_ACTIVATION_CLOSES_TO_F2"
	SupportTwoBoundaryGeneratesF2                = "CONDITIONAL_SUPPORT_TWO_BOUNDARY_FACTORS_GENERATE_SATURATED_AIRLOCK_RECTANGLE"
	SupportThetaTwoFromFullPairClosure           = "CONDITIONAL_SUPPORT_THETA_TWO_FOLLOWS_FROM_FULL_PAIR_CLOSURE"
	SupportCumulativeF2OverF0FromBase            = "CONDITIONAL_SUPPORT_CUMULATIVE_QUOTIENT_F2_OVER_F0_FOLLOWS_FROM_FIXED_BASEPOINT_QUOTIENT"
	SupportGradedRejectedByBasepointClosure      = "CONDITIONAL_SUPPORT_ASSOCIATED_GRADED_F2_OVER_F1_REJECTED_BY_BASEPOINT_CLOSURE_FORM"
	SupportTopDegreeClosureOverPuncture          = "CONDITIONAL_SUPPORT_TOP_DEGREE_TARGET_IS_CLOSURE_OVER_PUNCTURE_BASE"
	SupportClosureUniqueRules                    = "CONDITIONAL_SUPPORT_AIRLOCK_CLOSURE_FUNCTOR_IS_UNIQUE_UNDER_MONOTONE_MINIMAL_SATURATED_RULES"
	SupportMinimalityForcesF1                    = "CONDITIONAL_SUPPORT_MINIMALITY_FORCES_SINGLETON_TO_F1"
	SupportSaturationForcesF2                    = "CONDITIONAL_SUPPORT_SATURATION_FORCES_FULL_PAIR_TO_F2"
	SupportZ2ClassInvarianceClosure              = "CONDITIONAL_SUPPORT_Z2_CLASS_INVARIANCE_PRESERVED_BY_CLOSURE"
	SupportClosureSuppliesThetaTargets           = "CONDITIONAL_SUPPORT_CLOSURE_FACTORIZATION_SUPPLIES_THETA_B_Z2_TARGETS"
	SupportMeasureUsingAirlockClosure            = "CONDITIONAL_SUPPORT_BOUNDARY_ACTIVATION_MEASURE_CAN_BE_WRITTEN_USING_AIRLOCK_CLOSURE"
	SupportAlphaViaClosureMeasure                = "CONDITIONAL_SUPPORT_ALPHA_B_RECONSTRUCTED_THROUGH_CLOSURE_FACTORED_MEASURE"

	FailureSubsetCardinalityAloneNoTargets  = "FAILED_ROUTE_BOUNDARY_SUBSET_CARDINALITY_ALONE_DOES_NOT_SELECT_AIRLOCK_TARGETS"
	FailureAirlockClosureNotNative          = "FAILED_ROUTE_AIRLOCK_CLOSURE_LADDER_NOT_NATIVE_CLOSURE_OPERATOR_YET"
	FailureClosureFunctorCandidateNotNative = "FAILED_ROUTE_CLOSURE_FUNCTOR_IS_CANDIDATE_NOT_NATIVE_THEOREM"
	FailureBasepointClosureNotNative        = "FAILED_ROUTE_BASEPOINT_CLOSURE_NOT_NATIVE_AIRLOCK_THEOREM"
	FailureSingletonToF1ClosureNotNative    = "FAILED_ROUTE_SINGLETON_TO_F1_CLOSURE_NOT_NATIVE_THEOREM"
	FailureFullPairToF2ClosureNotNative     = "FAILED_ROUTE_FULL_PAIR_TO_F2_CLOSURE_NOT_NATIVE_THEOREM"
	FailureFixedBasepointQuotientNotNative  = "FAILED_ROUTE_FIXED_BASEPOINT_QUOTIENT_RULE_NOT_NATIVE_THEOREM"
	FailureMinimalSaturationNotNative       = "FAILED_ROUTE_MINIMALITY_AND_SATURATION_RULES_NOT_NATIVE_CERTIFIED"
	FailureMuBNotNativeWithoutClosure       = "FAILED_ROUTE_MU_B_STILL_NOT_NATIVE_WITHOUT_NATIVE_CLOSURE_FUNCTOR"
	FailureAlphaBridgeCandidateNotNative    = "FAILED_ROUTE_ALPHA_B_REMAINS_BRIDGE_CANDIDATE_NOT_NATIVE"
	FailureNotNativeR3                      = "FAILED_ROUTE_NOT_NATIVE_R3"
	FailureFullAFDescentStillBlocked        = "FAILED_ROUTE_FULL_A_F_DESCENT_STILL_BLOCKED"
	FailureNoGenerationCarrierMap           = "FAILED_ROUTE_NO_GENERATION_CARRIER_MAP"
	FailureNoFlavorOrientationMap           = "FAILED_ROUTE_NO_FLAVOR_ORIENTATION_MAP"
	FailureNoIndividualYukawaValues         = "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"
	FailureNoNativeYukawaOperator           = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_THEOREM"
)

type BoundarySubsetAudit struct {
	BoundaryPair                  string
	SourceChain                   string
	CardinalityChain              string
	Lambda1MatchesSingletons      bool
	Lambda2MatchesFullPair        bool
	ExteriorDegreeEqualsCardinal  bool
	NativeFiniteSubsetSource      bool
	SelectsAirlockTargetsByItself bool
	Supports                      []string
	Failures                      []string
}

type AirlockClosureLadderAudit struct {
	FlagChain             string
	F0Role                string
	F1Role                string
	F2Role                string
	ClosureLadderType     bool
	NativeClosureOperator bool
	Supports              []string
	Failures              []string
}

type ClosureFactorizationAudit struct {
	ClosureFunctor            string
	ThetaViaClosure           string
	ClosureZero               string
	ClosureOne                string
	ClosureTwo                string
	ThetaOne                  string
	ThetaTwo                  string
	FactorsTheta              bool
	QuotientReconstructsTheta bool
	NativeClosureTheorem      bool
	Supports                  []string
	Failures                  []string
}

type ClosureLevelAudit struct {
	Level                int
	BoundarySubset       string
	ClosureTarget        string
	ThetaTarget          string
	MatchesReducedForm   bool
	ClosureSupported     bool
	NativeClosureTheorem bool
	Supports             []string
	Failures             []string
}

type CumulativeQuotientAudit struct {
	TopDegreeClosure            string
	FixedBasepoint              string
	CumulativeTarget            string
	AssociatedGradedAlternative string
	CumulativeRank              int
	AssociatedGradedRank        int
	FollowsFromBasepoint        bool
	RejectsAssociatedGraded     bool
	NativeBasepointRule         bool
	Supports                    []string
	Failures                    []string
}

type ClosureUniquenessAudit struct {
	Cl0                     string
	Cl1                     string
	Cl2                     string
	Monotone                bool
	MinimalSingleton        bool
	SaturatedFullPair       bool
	Z2Invariant             bool
	UniqueUnderRules        bool
	NativeMinimalSaturation bool
	Supports                []string
	Failures                []string
}

type MeasureClosureAudit struct {
	MeasureFormula         string
	ThetaRankOne           int
	ThetaRankTwo           int
	H10Rank                int
	H72Rank                int
	AlphaFormula           string
	ClosureSuppliesTargets bool
	MeasureUsesClosure     bool
	AlphaReconstructed     bool
	NativeMeasureByClosure bool
	Supports               []string
	Failures               []string
}

type FirewallLedger struct {
	SubsetCardinalityAloneNoTargets  bool
	AirlockClosureNotNative          bool
	ClosureFunctorCandidateNotNative bool
	BasepointClosureNotNative        bool
	SingletonToF1ClosureNotNative    bool
	FullPairToF2ClosureNotNative     bool
	FixedBasepointQuotientNotNative  bool
	MinimalSaturationNotNative       bool
	MuBNotNativeWithoutClosure       bool
	AlphaBridgeCandidateNotNative    bool
	NotNativeR3                      bool
	FullAFDescentStillBlocked        bool
	NoGenerationCarrierMap           bool
	NoFlavorOrientationMap           bool
	NoIndividualYukawaValues         bool
	NoNativeYukawaOperator           bool
}

func (f FirewallLedger) List() []string {
	var out []string
	if f.SubsetCardinalityAloneNoTargets {
		out = append(out, FailureSubsetCardinalityAloneNoTargets)
	}
	if f.AirlockClosureNotNative {
		out = append(out, FailureAirlockClosureNotNative)
	}
	if f.ClosureFunctorCandidateNotNative {
		out = append(out, FailureClosureFunctorCandidateNotNative)
	}
	if f.BasepointClosureNotNative {
		out = append(out, FailureBasepointClosureNotNative)
	}
	if f.SingletonToF1ClosureNotNative {
		out = append(out, FailureSingletonToF1ClosureNotNative)
	}
	if f.FullPairToF2ClosureNotNative {
		out = append(out, FailureFullPairToF2ClosureNotNative)
	}
	if f.FixedBasepointQuotientNotNative {
		out = append(out, FailureFixedBasepointQuotientNotNative)
	}
	if f.MinimalSaturationNotNative {
		out = append(out, FailureMinimalSaturationNotNative)
	}
	if f.MuBNotNativeWithoutClosure {
		out = append(out, FailureMuBNotNativeWithoutClosure)
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

type Audit struct {
	ID              string
	InheritedStatus string
	Truth           string
	Classification  string
	ShortStatus     string
	Subset          BoundarySubsetAudit
	Ladder          AirlockClosureLadderAudit
	Factorization   ClosureFactorizationAudit
	Basepoint       ClosureLevelAudit
	Singleton       ClosureLevelAudit
	FullPair        ClosureLevelAudit
	Cumulative      CumulativeQuotientAudit
	Uniqueness      ClosureUniquenessAudit
	Measure         MeasureClosureAudit
	Firewalls       FirewallLedger
	Final           string
}

func BuildDefault() (Audit, error) {
	a := Audit{
		ID:              AuditID,
		InheritedStatus: Gate926ShortStatus,
		Truth:           FinalTruth,
		Classification:  Classification,
		ShortStatus:     ShortStatus,
		Subset: BoundarySubsetAudit{
			BoundaryPair:                  BoundaryPair,
			SourceChain:                   SourceDegreeChain,
			CardinalityChain:              CardinalityChain,
			Lambda1MatchesSingletons:      true,
			Lambda2MatchesFullPair:        true,
			ExteriorDegreeEqualsCardinal:  true,
			NativeFiniteSubsetSource:      true,
			SelectsAirlockTargetsByItself: false,
			Supports:                      []string{SupportExteriorDegreeEqualsSubsetCardinality, SupportLambda1SingletonSubsets, SupportLambda2FullPairSubset, SupportSourceChainSubsetLattice},
			Failures:                      []string{FailureSubsetCardinalityAloneNoTargets},
		},
		Ladder: AirlockClosureLadderAudit{
			FlagChain:             AirlockFlagChain,
			F0Role:                "puncture basepoint",
			F1Role:                "minimal exposed closure above puncture",
			F2Role:                "saturated full right rectangle closure",
			ClosureLadderType:     true,
			NativeClosureOperator: false,
			Supports:              []string{SupportAirlockFlagClosureLadder, SupportF0PunctureBasepoint, SupportF1MinimalExposedClosure, SupportF2SaturatedFullRightRectangle},
			Failures:              []string{FailureAirlockClosureNotNative},
		},
		Factorization: ClosureFactorizationAudit{
			ClosureFunctor:            ClosureFunctor,
			ThetaViaClosure:           ThetaViaClosure,
			ClosureZero:               ClosureZero,
			ClosureOne:                ClosureOne,
			ClosureTwo:                ClosureTwo,
			ThetaOne:                  ThetaOne,
			ThetaTwo:                  ThetaTwo,
			FactorsTheta:              true,
			QuotientReconstructsTheta: true,
			NativeClosureTheorem:      false,
			Supports:                  []string{SupportThetaFactorsThroughClosure, SupportDegreeToFlagHasClosureSource, SupportClosureQuotientReconstructsTheta},
			Failures:                  []string{FailureClosureFunctorCandidateNotNative},
		},
		Basepoint: ClosureLevelAudit{
			Level:                0,
			BoundarySubset:       "empty boundary subset |S|=0",
			ClosureTarget:        "F_0",
			ThetaTarget:          "no active alpha quotient",
			MatchesReducedForm:   true,
			ClosureSupported:     true,
			NativeClosureTheorem: false,
			Supports:             []string{SupportEmptySubsetMapsF0, SupportBasepointClosureMatchesReduction, SupportDegreeZeroNoActiveAlpha},
			Failures:             []string{FailureBasepointClosureNotNative},
		},
		Singleton: ClosureLevelAudit{
			Level:                1,
			BoundarySubset:       "singleton boundary subset |S|=1",
			ClosureTarget:        "F_1",
			ThetaTarget:          ThetaOne,
			MatchesReducedForm:   true,
			ClosureSupported:     true,
			NativeClosureTheorem: false,
			Supports:             []string{SupportSingletonClosesF1, SupportOneBoundaryGeneratesF1, SupportThetaOneFromSingletonClosure},
			Failures:             []string{FailureSingletonToF1ClosureNotNative},
		},
		FullPair: ClosureLevelAudit{
			Level:                2,
			BoundarySubset:       "full boundary subset |S|=2",
			ClosureTarget:        "F_2",
			ThetaTarget:          ThetaTwo,
			MatchesReducedForm:   true,
			ClosureSupported:     true,
			NativeClosureTheorem: false,
			Supports:             []string{SupportFullPairClosesF2, SupportTwoBoundaryGeneratesF2, SupportThetaTwoFromFullPairClosure},
			Failures:             []string{FailureFullPairToF2ClosureNotNative},
		},
		Cumulative: CumulativeQuotientAudit{
			TopDegreeClosure:            "Cl(2)=F_2",
			FixedBasepoint:              "F_0",
			CumulativeTarget:            "F_2/F_0",
			AssociatedGradedAlternative: AssociatedGradedTarget,
			CumulativeRank:              RankF2OverF0,
			AssociatedGradedRank:        RankF2OverF1,
			FollowsFromBasepoint:        true,
			RejectsAssociatedGraded:     true,
			NativeBasepointRule:         false,
			Supports:                    []string{SupportCumulativeF2OverF0FromBase, SupportGradedRejectedByBasepointClosure, SupportTopDegreeClosureOverPuncture},
			Failures:                    []string{FailureFixedBasepointQuotientNotNative},
		},
		Uniqueness: ClosureUniquenessAudit{
			Cl0:                     ClosureZero,
			Cl1:                     ClosureOne,
			Cl2:                     ClosureTwo,
			Monotone:                true,
			MinimalSingleton:        true,
			SaturatedFullPair:       true,
			Z2Invariant:             true,
			UniqueUnderRules:        true,
			NativeMinimalSaturation: false,
			Supports:                []string{SupportClosureUniqueRules, SupportMinimalityForcesF1, SupportSaturationForcesF2, SupportZ2ClassInvarianceClosure},
			Failures:                []string{FailureMinimalSaturationNotNative},
		},
		Measure: MeasureClosureAudit{
			MeasureFormula:         MeasureViaClosure,
			ThetaRankOne:           RankF1OverF0,
			ThetaRankTwo:           RankF2OverF0,
			H10Rank:                RankH10,
			H72Rank:                RankH72,
			AlphaFormula:           AlphaFormula,
			ClosureSuppliesTargets: true,
			MeasureUsesClosure:     true,
			AlphaReconstructed:     true,
			NativeMeasureByClosure: false,
			Supports:               []string{SupportClosureSuppliesThetaTargets, SupportMeasureUsingAirlockClosure, SupportAlphaViaClosureMeasure},
			Failures:               []string{FailureMuBNotNativeWithoutClosure, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3},
		},
		Firewalls: FirewallLedger{
			SubsetCardinalityAloneNoTargets:  true,
			AirlockClosureNotNative:          true,
			ClosureFunctorCandidateNotNative: true,
			BasepointClosureNotNative:        true,
			SingletonToF1ClosureNotNative:    true,
			FullPairToF2ClosureNotNative:     true,
			FixedBasepointQuotientNotNative:  true,
			MinimalSaturationNotNative:       true,
			MuBNotNativeWithoutClosure:       true,
			AlphaBridgeCandidateNotNative:    true,
			NotNativeR3:                      true,
			FullAFDescentStillBlocked:        true,
			NoGenerationCarrierMap:           true,
			NoFlavorOrientationMap:           true,
			NoIndividualYukawaValues:         true,
			NoNativeYukawaOperator:           true,
		},
		Final: NextGate,
	}
	if err := a.Validate(); err != nil {
		return Audit{}, err
	}
	return a, nil
}

func (a Audit) Validate() error {
	if a.InheritedStatus != Gate926ShortStatus {
		return fmt.Errorf("unexpected inherited status %q", a.InheritedStatus)
	}
	if a.Truth != FinalTruth || a.Classification != Classification || a.ShortStatus != ShortStatus {
		return fmt.Errorf("unexpected verdict/classification/status")
	}
	if !a.Subset.ExteriorDegreeEqualsCardinal || !a.Subset.NativeFiniteSubsetSource || a.Subset.SelectsAirlockTargetsByItself {
		return fmt.Errorf("boundary subset audit leaked or failed: %s", FormatSubset(a.Subset))
	}
	if !a.Ladder.ClosureLadderType || a.Ladder.NativeClosureOperator {
		return fmt.Errorf("airlock ladder audit leaked or failed: %s", FormatLadder(a.Ladder))
	}
	if !a.Factorization.FactorsTheta || !a.Factorization.QuotientReconstructsTheta || a.Factorization.NativeClosureTheorem {
		return fmt.Errorf("closure factorization audit leaked or failed: %s", FormatFactorization(a.Factorization))
	}
	if !a.Basepoint.ClosureSupported || !a.Basepoint.MatchesReducedForm || a.Basepoint.NativeClosureTheorem {
		return fmt.Errorf("basepoint closure audit leaked or failed: %s", FormatClosureLevel(a.Basepoint))
	}
	if !a.Singleton.ClosureSupported || !a.Singleton.MatchesReducedForm || a.Singleton.NativeClosureTheorem || a.Singleton.ClosureTarget != "F_1" {
		return fmt.Errorf("singleton closure audit leaked or failed: %s", FormatClosureLevel(a.Singleton))
	}
	if !a.FullPair.ClosureSupported || !a.FullPair.MatchesReducedForm || a.FullPair.NativeClosureTheorem || a.FullPair.ClosureTarget != "F_2" {
		return fmt.Errorf("full-pair closure audit leaked or failed: %s", FormatClosureLevel(a.FullPair))
	}
	if !a.Cumulative.FollowsFromBasepoint || !a.Cumulative.RejectsAssociatedGraded || a.Cumulative.NativeBasepointRule || a.Cumulative.CumulativeRank != RankF2OverF0 || a.Cumulative.AssociatedGradedRank != RankF2OverF1 {
		return fmt.Errorf("cumulative quotient audit leaked or failed: %s", FormatCumulative(a.Cumulative))
	}
	if !a.Uniqueness.UniqueUnderRules || !a.Uniqueness.Monotone || !a.Uniqueness.MinimalSingleton || !a.Uniqueness.SaturatedFullPair || !a.Uniqueness.Z2Invariant || a.Uniqueness.NativeMinimalSaturation {
		return fmt.Errorf("closure uniqueness audit leaked or failed: %s", FormatUniqueness(a.Uniqueness))
	}
	if !a.Measure.ClosureSuppliesTargets || !a.Measure.MeasureUsesClosure || !a.Measure.AlphaReconstructed || a.Measure.NativeMeasureByClosure || a.Measure.ThetaRankOne != RankF1OverF0 || a.Measure.ThetaRankTwo != RankF2OverF0 || a.Measure.H10Rank != RankH10 || a.Measure.H72Rank != RankH72 {
		return fmt.Errorf("measure closure audit leaked or failed: %s", FormatMeasure(a.Measure))
	}
	if !firewallsOK(a.Firewalls) {
		return fmt.Errorf("firewall ledger incomplete: %s", FormatFirewalls(a.Firewalls))
	}
	return nil
}

func Statuses() []string {
	return []string{StatusInheritedGate926, StatusSubsetLatticeSource, StatusAirlockClosureLadderTyped, StatusThetaFactorsThroughClosure, StatusBasepointClosure, StatusSingletonClosure, StatusFullPairClosure, StatusCumulativeQuotientFromBase, StatusClosureUniqueUnderRules, StatusClosureMeasureReconstruction, StatusNativeClosureMissing, StatusAlphaR3StillBlocked}
}

func Supports() []string {
	return []string{SupportExteriorDegreeEqualsSubsetCardinality, SupportLambda1SingletonSubsets, SupportLambda2FullPairSubset, SupportSourceChainSubsetLattice, SupportAirlockFlagClosureLadder, SupportF0PunctureBasepoint, SupportF1MinimalExposedClosure, SupportF2SaturatedFullRightRectangle, SupportThetaFactorsThroughClosure, SupportDegreeToFlagHasClosureSource, SupportClosureQuotientReconstructsTheta, SupportEmptySubsetMapsF0, SupportBasepointClosureMatchesReduction, SupportDegreeZeroNoActiveAlpha, SupportSingletonClosesF1, SupportOneBoundaryGeneratesF1, SupportThetaOneFromSingletonClosure, SupportFullPairClosesF2, SupportTwoBoundaryGeneratesF2, SupportThetaTwoFromFullPairClosure, SupportCumulativeF2OverF0FromBase, SupportGradedRejectedByBasepointClosure, SupportTopDegreeClosureOverPuncture, SupportClosureUniqueRules, SupportMinimalityForcesF1, SupportSaturationForcesF2, SupportZ2ClassInvarianceClosure, SupportClosureSuppliesThetaTargets, SupportMeasureUsingAirlockClosure, SupportAlphaViaClosureMeasure}
}

func Failures() []string {
	return []string{FailureSubsetCardinalityAloneNoTargets, FailureAirlockClosureNotNative, FailureClosureFunctorCandidateNotNative, FailureBasepointClosureNotNative, FailureSingletonToF1ClosureNotNative, FailureFullPairToF2ClosureNotNative, FailureFixedBasepointQuotientNotNative, FailureMinimalSaturationNotNative, FailureMuBNotNativeWithoutClosure, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}
}

func FormatSubset(a BoundarySubsetAudit) string {
	return fmt.Sprintf("boundary_pair=%s source_chain=%s cardinality_chain=%s lambda1_singletons=%t lambda2_full_pair=%t degree_equals_cardinality=%t native_subset_source=%t selects_airlock_targets_by_itself=%t supports=%s failures=%s", a.BoundaryPair, a.SourceChain, a.CardinalityChain, a.Lambda1MatchesSingletons, a.Lambda2MatchesFullPair, a.ExteriorDegreeEqualsCardinal, a.NativeFiniteSubsetSource, a.SelectsAirlockTargetsByItself, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatLadder(a AirlockClosureLadderAudit) string {
	return fmt.Sprintf("flag=%s F0=%s F1=%s F2=%s closure_ladder_type=%t native_closure_operator=%t supports=%s failures=%s", a.FlagChain, a.F0Role, a.F1Role, a.F2Role, a.ClosureLadderType, a.NativeClosureOperator, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatFactorization(a ClosureFactorizationAudit) string {
	return fmt.Sprintf("closure=%s theta_via_closure=%s cl0=%s cl1=%s cl2=%s theta1=%s theta2=%s factors_theta=%t quotient_reconstructs_theta=%t native_closure=%t supports=%s failures=%s", a.ClosureFunctor, a.ThetaViaClosure, a.ClosureZero, a.ClosureOne, a.ClosureTwo, a.ThetaOne, a.ThetaTwo, a.FactorsTheta, a.QuotientReconstructsTheta, a.NativeClosureTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatClosureLevel(a ClosureLevelAudit) string {
	return fmt.Sprintf("level=%d subset=%s closure_target=%s theta_target=%s matches_reduced_form=%t closure_supported=%t native_closure=%t supports=%s failures=%s", a.Level, a.BoundarySubset, a.ClosureTarget, a.ThetaTarget, a.MatchesReducedForm, a.ClosureSupported, a.NativeClosureTheorem, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatCumulative(a CumulativeQuotientAudit) string {
	return fmt.Sprintf("top_closure=%s basepoint=%s cumulative_target=%s graded_alt=%s cumulative_rank=%d graded_rank=%d follows_from_basepoint=%t rejects_graded=%t native_basepoint_rule=%t supports=%s failures=%s", a.TopDegreeClosure, a.FixedBasepoint, a.CumulativeTarget, a.AssociatedGradedAlternative, a.CumulativeRank, a.AssociatedGradedRank, a.FollowsFromBasepoint, a.RejectsAssociatedGraded, a.NativeBasepointRule, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatUniqueness(a ClosureUniquenessAudit) string {
	return fmt.Sprintf("cl0=%s cl1=%s cl2=%s monotone=%t minimal_singleton=%t saturated_full_pair=%t z2_invariant=%t unique_under_rules=%t native_min_sat=%t supports=%s failures=%s", a.Cl0, a.Cl1, a.Cl2, a.Monotone, a.MinimalSingleton, a.SaturatedFullPair, a.Z2Invariant, a.UniqueUnderRules, a.NativeMinimalSaturation, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatMeasure(a MeasureClosureAudit) string {
	return fmt.Sprintf("measure=%s rank1=%d rank2=%d h10=%d h72=%d alpha=%s closure_supplies_targets=%t uses_closure=%t alpha_reconstructed=%t native_measure=%t supports=%s failures=%s", a.MeasureFormula, a.ThetaRankOne, a.ThetaRankTwo, a.H10Rank, a.H72Rank, a.AlphaFormula, a.ClosureSuppliesTargets, a.MeasureUsesClosure, a.AlphaReconstructed, a.NativeMeasureByClosure, strings.Join(a.Supports, ","), strings.Join(a.Failures, ","))
}

func FormatFirewalls(f FirewallLedger) string {
	return strings.Join(f.List(), ",")
}

func firewallsOK(f FirewallLedger) bool {
	required := []string{FailureSubsetCardinalityAloneNoTargets, FailureAirlockClosureNotNative, FailureClosureFunctorCandidateNotNative, FailureBasepointClosureNotNative, FailureSingletonToF1ClosureNotNative, FailureFullPairToF2ClosureNotNative, FailureFixedBasepointQuotientNotNative, FailureMinimalSaturationNotNative, FailureMuBNotNativeWithoutClosure, FailureAlphaBridgeCandidateNotNative, FailureNotNativeR3, FailureFullAFDescentStillBlocked, FailureNoGenerationCarrierMap, FailureNoFlavorOrientationMap, FailureNoIndividualYukawaValues, FailureNoNativeYukawaOperator}
	return containsAll(f.List(), required)
}

func containsAll(got []string, want []string) bool {
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
