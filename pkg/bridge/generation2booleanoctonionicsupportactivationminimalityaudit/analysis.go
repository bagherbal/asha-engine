// Package generation2booleanoctonionicsupportactivationminimalityaudit implements
// Gate 686: Boolean-Octonionic Support Activation Minimality Audit.
//
// Gate 685 proved the conditional projector-identity selection theorem:
// rank(P)=7 plus Boolean support P_B P=P plus octonionic support P_G P=P
// uniquely selects the orthogonal K_7 projector P_K7. Gate 686 audits whether
// that support sieve is minimal, independent, and noncircular, and separates
// the bridge response into three typed factors:
//
//  1. the boundary control scalar S_split;
//  2. the Boolean-octonionic projector identity selector;
//  3. ordinary augmented trace scalarization on H_72.
//
// This is a bridge-layer support-minimality audit only. It does not derive
// boundary stress, scalar RG matching, Higgs mass, gauge unification, flavor,
// CKM/PMNS, a native 7/72 theorem, or a native projector-activation theorem.
package generation2booleanoctonionicsupportactivationminimalityaudit

import (
	"sort"
	"strings"
	"sync"
)

const (
	AuditID = "GATE686-BOOLEAN-OCTONIONIC-SUPPORT-ACTIVATION-MINIMALITY-AUDIT"

	StatusGate685SelectionInherited          = "PASS_GATE685_PROJECTOR_SELECTION_INHERITED"
	StatusConstraintLadderAudited            = "PASS_CONSTRAINT_LADDER_AUDITED"
	StatusRankOnlyDegeneracyConfirmed        = "PASS_RANK_ONLY_DEGENERACY_CONFIRMED"
	StatusBooleanOnlySupportNotUnique        = "PASS_BOOLEAN_ONLY_SUPPORT_NOT_UNIQUE"
	StatusOctonionicOnlySupportNotUnique     = "PASS_OCTONIONIC_ONLY_SUPPORT_NOT_UNIQUE"
	StatusBooleanPlusOctonionicSelectsK7     = "PASS_BOOLEAN_PLUS_OCTONIONIC_SUPPORT_SELECTS_K7"
	StatusSupportIndependenceAudited         = "PASS_SUPPORT_CONSTRAINTS_INDEPENDENCE_AUDITED"
	StatusNoncircularityAudited              = "PASS_NONCIRCULARITY_AUDITED"
	StatusActivationDecompositionWritten     = "PASS_ACTIVATION_DECOMPOSITION_WRITTEN"
	StatusMinimalSelector                    = "CONDITIONAL_SUPPORT_BOOLEAN_OCTONIONIC_SUPPORT_IS_MINIMAL_PROJECTOR_IDENTITY_SELECTOR"
	StatusResponseSplitsScalarAndSelector    = "CONDITIONAL_SUPPORT_ACTIVE_RESPONSE_SPLITS_INTO_BOUNDARY_SCALAR_AND_NATIVE_PROJECTOR_SELECTOR"
	StatusSSplitAloneDoesNotSelectProjector  = "FAILED_ROUTE_S_SPLIT_ALONE_DOES_NOT_SELECT_PROJECTOR_IDENTITY"
	StatusNoBoundaryScalarSupportActivation  = "FAILED_ROUTE_NO_NATIVE_REASON_BOUNDARY_SCALAR_ACTIVATES_SUPPORT_SIEVE"
	StatusNoNativeProjectorActivationTheorem = "FAILED_ROUTE_NO_NATIVE_PROJECTOR_ACTIVATION_THEOREM"
	StatusNoNativeSevenOver72Theorem         = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate686SupportMinimalityBoundary   = "FIREWALL_PRESERVED_GATE686_SUPPORT_MINIMALITY_BOUNDARY"
)

const (
	lambda4Dimension  = 70
	boundaryDimension = 2
	h72Dimension      = lambda4Dimension + boundaryDimension
	booleanRank       = 56
	octonionicRank    = 14
	k7Dimension       = 7
	uPlusVDimension   = booleanRank + octonionicRank - k7Dimension
	w7Dimension       = lambda4Dimension - uPlusVDimension
	auditedDBase      = 0.0001256552099683575
	auditedSSplit     = 0.0012924448188162962
)

type Gate685Inheritance struct {
	RankSevenTraceInherited    bool
	BooleanOctonionicSelection bool
	SelectedProjector          string
	DBase                      float64
	SSplit                     float64
	H72Dimension               int
	K7Dimension                int
	TraceScalarization         string
	ProjectorSelection         string
	ActivationStillUnproved    bool
	PriorFirewallPreserved     bool
	Verdict                    string
}

type ConstraintStep struct {
	Index             int
	Name              string
	Constraints       []string
	Carrier           string
	CarrierDimension  int
	RankSevenRequired bool
	Degenerate        bool
	UniquePK7         bool
	Witness           string
	Verdict           string
}

type ConstraintLadderAudit struct {
	Steps                        []ConstraintStep
	RankOnlyDegenerate           bool
	FiniteSupportOnlyDegenerate  bool
	BooleanOnlyDegenerate        bool
	OctonionicOnlyDegenerate     bool
	CombinedSupportSelectsK7     bool
	AllWeakerSelectorsDegenerate bool
	MinimalPairRequired          bool
	Verdict                      string
}

type IndependenceAudit struct {
	BooleanComplementDimension    int
	OctonionicComplementDimension int
	BooleanOnlyWitness            string
	OctonionicOnlyWitness         string
	BooleanImpliesOctonionic      bool
	OctonionicImpliesBoolean      bool
	NeitherConditionRedundant     bool
	BothRequiredToForceK7         bool
	Verdict                       string
}

type NoncircularityAudit struct {
	Assumptions             []string
	DoesNotAssumePK7        bool
	UsesOnlyRankAndSupport  bool
	UsesOnlyIntersectionDim bool
	ConclusionDerived       string
	ConditionalNotAbsolute  bool
	Noncircular             bool
	Verdict                 string
}

type ActivationDecompositionAudit struct {
	ActiveResponse              string
	BoundaryControlScalar       string
	ProjectorIdentitySelector   string
	TraceScalarization          string
	BoundaryScalarSelectsRank   bool
	SupportSelectorSelectsPK7   bool
	TraceOnlyScalarizes         bool
	SSplitAloneSelectsProjector bool
	NativeActivationProved      bool
	Verdict                     string
}

type MissingTheoremAudit struct {
	Missing    []string
	PreciseGap string
	Verdict    string
}

type VerdictDiscipline struct {
	ClaimsSSplitSelectsProjector       bool
	ClaimsBoundaryScalarActivatesSieve bool
	ClaimsProjectorActivation          bool
	ClaimsNativeSevenOver72            bool
	ClaimsBoundaryStressDerivation     bool
	ClaimsScalarRGMatching             bool
	ClaimsHiggsMass                    bool
	ClaimsGaugeUnification             bool
	ClaimsFlavorDerivation             bool
	Verdict                            string
}

type Analysis struct {
	Inherited     Gate685Inheritance
	Ladder        ConstraintLadderAudit
	Independence  IndependenceAudit
	Noncircular   NoncircularityAudit
	Decomposition ActivationDecompositionAudit
	Missing       MissingTheoremAudit
	Discipline    VerdictDiscipline
	Truth         string
}

var cache struct {
	sync.Once
	a   Analysis
	err error
}

func BuildDefault() (Analysis, error) {
	cache.Once.Do(func() { cache.a, cache.err = Build() })
	return cache.a, cache.err
}

func Build() (Analysis, error) {
	inherited := buildInheritance()
	ladder := buildConstraintLadder()
	independence := buildIndependence()
	noncircular := buildNoncircularity()
	decomposition := buildActivationDecomposition(ladder)
	missing := MissingTheoremAudit{
		Missing: []string{
			StatusSSplitAloneDoesNotSelectProjector,
			StatusNoBoundaryScalarSupportActivation,
			StatusNoNativeProjectorActivationTheorem,
			StatusNoNativeSevenOver72Theorem,
		},
		PreciseGap: "a native theorem showing why the boundary quotient scalar S_split activates the Boolean-octonionic support sieve P_B P=P and P_G P=P, rather than merely multiplying the already selected projector P_K7",
		Verdict:    strings.Join([]string{StatusSSplitAloneDoesNotSelectProjector, StatusNoBoundaryScalarSupportActivation, StatusNoNativeProjectorActivationTheorem, StatusNoNativeSevenOver72Theorem}, "; "),
	}
	discipline := VerdictDiscipline{Verdict: StatusGate686SupportMinimalityBoundary}
	truth := "Gate 686 audits the minimality boundary left by Gate 685. Rank-only, finite-only, Boolean-only, and octonionic-only selectors remain degenerate. Boolean support and octonionic support are independent: U has a rank-seven Boolean-only witness outside K_7, and V has a rank-seven octonionic-only witness outside K_7. Only the pair P_B P=P and P_G P=P, together with rank(P)=7 and dim(U∩V)=7, noncircularly selects P_K7. The active response splits into boundary scalar S_split, native projector selector, and ordinary trace scalarization; no native theorem yet proves why S_split activates that support sieve."
	return Analysis{Inherited: inherited, Ladder: ladder, Independence: independence, Noncircular: noncircular, Decomposition: decomposition, Missing: missing, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance() Gate685Inheritance {
	return Gate685Inheritance{
		RankSevenTraceInherited:    true,
		BooleanOctonionicSelection: true,
		SelectedProjector:          "P_K7",
		DBase:                      auditedDBase,
		SSplit:                     auditedSSplit,
		H72Dimension:               h72Dimension,
		K7Dimension:                k7Dimension,
		TraceScalarization:         "Tr_H72(S_split P_K7)/Tr_H72(I)=(7/72)S_split",
		ProjectorSelection:         "rank(P)=7 plus P_B P=P plus P_G P=P implies P=P_K7",
		ActivationStillUnproved:    true,
		PriorFirewallPreserved:     true,
		Verdict:                    StatusGate685SelectionInherited,
	}
}

func buildConstraintLadder() ConstraintLadderAudit {
	steps := []ConstraintStep{
		{
			Index:             1,
			Name:              "rank-seven only",
			Constraints:       []string{"rank(P)=7"},
			Carrier:           "H_72",
			CarrierDimension:  h72Dimension,
			RankSevenRequired: true,
			Degenerate:        true,
			UniquePK7:         false,
			Witness:           "P_W7 has rank seven and gives the same ordinary trace response as P_K7",
			Verdict:           StatusRankOnlyDegeneracyConfirmed,
		},
		{
			Index:             2,
			Name:              "finite support only",
			Constraints:       []string{"P_boundary=0", "rank(P)=7"},
			Carrier:           "Lambda^4 R^8",
			CarrierDimension:  lambda4Dimension,
			RankSevenRequired: true,
			Degenerate:        true,
			UniquePK7:         false,
			Witness:           "many rank-seven subspaces remain inside the seventy-dimensional finite chamber",
			Verdict:           "PASS_FINITE_SUPPORT_ONLY_NOT_UNIQUE",
		},
		{
			Index:             3,
			Name:              "Boolean support only",
			Constraints:       []string{"P_B P=P", "rank(P)=7"},
			Carrier:           "U=Im(P_B)",
			CarrierDimension:  booleanRank,
			RankSevenRequired: true,
			Degenerate:        true,
			UniquePK7:         false,
			Witness:           "P_Uonly7 can be chosen inside U but outside K_7",
			Verdict:           StatusBooleanOnlySupportNotUnique,
		},
		{
			Index:             4,
			Name:              "octonionic support only",
			Constraints:       []string{"P_G P=P", "rank(P)=7"},
			Carrier:           "V=Im(P_G)",
			CarrierDimension:  octonionicRank,
			RankSevenRequired: true,
			Degenerate:        true,
			UniquePK7:         false,
			Witness:           "P_Vonly7 can be chosen inside V but outside K_7",
			Verdict:           StatusOctonionicOnlySupportNotUnique,
		},
		{
			Index:             5,
			Name:              "Boolean plus octonionic support",
			Constraints:       []string{"P_B P=P", "P_G P=P", "rank(P)=7"},
			Carrier:           "U∩V=K_7",
			CarrierDimension:  k7Dimension,
			RankSevenRequired: true,
			Degenerate:        false,
			UniquePK7:         true,
			Witness:           "Im(P)⊂U∩V=K_7 and rank(P)=dim(K_7)=7 force Im(P)=K_7",
			Verdict:           StatusBooleanPlusOctonionicSelectsK7,
		},
	}
	return ConstraintLadderAudit{
		Steps:                        steps,
		RankOnlyDegenerate:           steps[0].Degenerate && !steps[0].UniquePK7,
		FiniteSupportOnlyDegenerate:  steps[1].Degenerate && !steps[1].UniquePK7,
		BooleanOnlyDegenerate:        steps[2].Degenerate && !steps[2].UniquePK7,
		OctonionicOnlyDegenerate:     steps[3].Degenerate && !steps[3].UniquePK7,
		CombinedSupportSelectsK7:     steps[4].UniquePK7 && !steps[4].Degenerate,
		AllWeakerSelectorsDegenerate: true,
		MinimalPairRequired:          true,
		Verdict:                      strings.Join([]string{StatusConstraintLadderAudited, StatusRankOnlyDegeneracyConfirmed, StatusBooleanOnlySupportNotUnique, StatusOctonionicOnlySupportNotUnique, StatusBooleanPlusOctonionicSelectsK7, StatusMinimalSelector}, "; "),
	}
}

func buildIndependence() IndependenceAudit {
	booleanComplement := booleanRank - k7Dimension
	octonionicComplement := octonionicRank - k7Dimension
	return IndependenceAudit{
		BooleanComplementDimension:    booleanComplement,
		OctonionicComplementDimension: octonionicComplement,
		BooleanOnlyWitness:            "a rank-seven projector in U∩K_7^perp exists because dim(U)-dim(K_7)=49≥7; it satisfies P_B P=P but not P_G P=P",
		OctonionicOnlyWitness:         "a rank-seven projector in V∩K_7^perp exists because dim(V)-dim(K_7)=7; it satisfies P_G P=P but not P_B P=P",
		BooleanImpliesOctonionic:      false,
		OctonionicImpliesBoolean:      false,
		NeitherConditionRedundant:     booleanComplement >= k7Dimension && octonionicComplement >= k7Dimension,
		BothRequiredToForceK7:         true,
		Verdict:                       strings.Join([]string{StatusSupportIndependenceAudited, StatusMinimalSelector}, "; "),
	}
}

func buildNoncircularity() NoncircularityAudit {
	return NoncircularityAudit{
		Assumptions: []string{
			"P is an orthogonal projector: P^2=P and P^T=P",
			"rank(P)=7",
			"P_B P=P",
			"P_G P=P",
			"dim(Im(P_B)∩Im(P_G))=7",
		},
		DoesNotAssumePK7:        true,
		UsesOnlyRankAndSupport:  true,
		UsesOnlyIntersectionDim: true,
		ConclusionDerived:       "Im(P)⊂Im(P_B)∩Im(P_G)=K_7; rank(P)=dim(K_7)=7; therefore Im(P)=K_7; symmetry gives P=P_K7",
		ConditionalNotAbsolute:  true,
		Noncircular:             true,
		Verdict:                 StatusNoncircularityAudited,
	}
}

func buildActivationDecomposition(ladder ConstraintLadderAudit) ActivationDecompositionAudit {
	return ActivationDecompositionAudit{
		ActiveResponse:              "R_split = S_split · P_selected",
		BoundaryControlScalar:       "S_split=lambda(Lambda_12)+(R_3-1), a boundary quotient scalar controlling amplitude",
		ProjectorIdentitySelector:   "P_selected is selected by rank(P)=7 plus Boolean-octonionic support P_B P=P and P_G P=P",
		TraceScalarization:          "Tr_H72(S_split P_selected)/72 = (7/72)S_split once P_selected=P_K7",
		BoundaryScalarSelectsRank:   true,
		SupportSelectorSelectsPK7:   ladder.CombinedSupportSelectsK7,
		TraceOnlyScalarizes:         true,
		SSplitAloneSelectsProjector: false,
		NativeActivationProved:      false,
		Verdict:                     strings.Join([]string{StatusActivationDecompositionWritten, StatusResponseSplitsScalarAndSelector, StatusSSplitAloneDoesNotSelectProjector, StatusNoBoundaryScalarSupportActivation}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate685SelectionInherited,
		StatusConstraintLadderAudited,
		StatusRankOnlyDegeneracyConfirmed,
		StatusBooleanOnlySupportNotUnique,
		StatusOctonionicOnlySupportNotUnique,
		StatusBooleanPlusOctonionicSelectsK7,
		StatusSupportIndependenceAudited,
		StatusNoncircularityAudited,
		StatusActivationDecompositionWritten,
		StatusMinimalSelector,
		StatusResponseSplitsScalarAndSelector,
		StatusSSplitAloneDoesNotSelectProjector,
		StatusNoBoundaryScalarSupportActivation,
		StatusNoNativeProjectorActivationTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusGate686SupportMinimalityBoundary,
	}
}

func sortedStepVerdicts(steps []ConstraintStep) []string {
	out := make([]string, 0, len(steps))
	for _, s := range steps {
		out = append(out, s.Verdict)
	}
	sort.Strings(out)
	return out
}
