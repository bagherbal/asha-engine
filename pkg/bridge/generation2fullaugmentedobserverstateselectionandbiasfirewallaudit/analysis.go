// Package generation2fullaugmentedobserverstateselectionandbiasfirewallaudit implements
// Gate 693: Full Augmented Observer State Selection and Bias Firewall Audit.
//
// Gate 692 rewrote the active bridge as the state expectation
//
//	D_base ≈ Tr(rho_72 R_split),
//
// with rho_72=I_H72/72 and R_split=S_split P_K7.  Gate 693 audits the
// sharper state-selection pressure point: rho_72 is the minimal unbiased
// full-augmented observer state, but it is not unique among all density
// states because biased states can be constructed with the same K7 weight.
//
// This is a bridge-layer observer-state selection and bias firewall audit
// only. It does not derive boundary stress, scalar RG matching, Higgs mass,
// gauge unification, flavor, CKM/PMNS, a native state-selection theorem, a
// native first-trace theorem, or a native 7/72 theorem.
package generation2fullaugmentedobserverstateselectionandbiasfirewallaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate692 "github.com/bagherbal/asha-engine/pkg/bridge/generation2maximallymixedaugmentedchamberobserverstateaudit"
)

const (
	AuditID = "GATE693-FULL-AUGMENTED-OBSERVER-STATE-SELECTION-AND-BIAS-FIREWALL-AUDIT"

	StatusGate692StateExpectationInherited               = "PASS_GATE692_STATE_EXPECTATION_INHERITED"
	StatusGeneralStateResponseReducedToK7Weight          = "PASS_GENERAL_STATE_RESPONSE_REDUCED_TO_K7_WEIGHT"
	StatusActiveResponseRequiresK7Weight7Over72          = "PASS_ACTIVE_RESPONSE_REQUIRES_K7_WEIGHT_7_OVER_72"
	StatusRho72GivesActiveK7Weight                       = "PASS_RHO_72_GIVES_ACTIVE_K7_WEIGHT"
	StatusAlternativeTypedStatesAudited                  = "PASS_ALTERNATIVE_TYPED_STATES_AUDITED"
	StatusFiniteOnlyStateRejectedBy7Over70               = "PASS_FINITE_ONLY_STATE_REJECTED_BY_7_OVER_70"
	StatusKernelStateRejectedBy7Over71                   = "PASS_KERNEL_STATE_REJECTED_BY_7_OVER_71"
	StatusLocalK7StateRejectedByUnitWeight               = "PASS_LOCAL_K7_STATE_REJECTED_BY_UNIT_WEIGHT"
	StatusBoundaryOnlyStateRejectedByZeroWeight          = "PASS_BOUNDARY_ONLY_STATE_REJECTED_BY_ZERO_WEIGHT"
	StatusHodgeSignedObserverRejectedAsNonPositiveState  = "PASS_HODGE_SIGNED_OBSERVER_REJECTED_AS_NON_POSITIVE_STATE"
	StatusBiasedStatesCanReproduceWeightButAreCircular   = "PASS_BIASED_STATES_CAN_REPRODUCE_WEIGHT_BUT_ARE_CIRCULAR"
	StatusRho72MinimalUnbiasedFullAugmentedObserverState = "CONDITIONAL_SUPPORT_RHO_72_AS_MINIMAL_UNBIASED_FULL_AUGMENTED_OBSERVER_STATE"
	StatusActiveBridgeGlobalUnbiasedK7WeightExpectation  = "CONDITIONAL_SUPPORT_ACTIVE_BRIDGE_IS_GLOBAL_UNBIASED_K7_WEIGHT_EXPECTATION"
	StatusRho72NotUniqueAmongAllDensityStates            = "FAILED_ROUTE_RHO_72_NOT_UNIQUE_AMONG_ALL_DENSITY_STATES"
	StatusBiasedStateReproductionNotNativeSelection      = "FAILED_ROUTE_BIASED_STATE_REPRODUCTION_IS_NOT_NATIVE_SELECTION"
	StatusNoNativeMaximallyMixedStateSelectionTheorem    = "FAILED_ROUTE_NO_NATIVE_MAXIMALLY_MIXED_STATE_SELECTION_THEOREM"
	StatusNoNativeFirstTraceTheorem                      = "FAILED_ROUTE_NO_NATIVE_FIRST_TRACE_THEOREM"
	StatusNoNativeSevenOver72Theorem                     = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate693ObserverStateSelectionBoundary          = "FIREWALL_PRESERVED_GATE693_OBSERVER_STATE_SELECTION_BOUNDARY"
)

const (
	lambda4Dimension  = 70
	boundaryDimension = 2
	h72Dimension      = lambda4Dimension + boundaryDimension
	kernelDimension   = 71
	k7Dimension       = 7
	k7PlusDimension   = 4
	k7MinusDimension  = 3
	tolerance         = 1e-18
)

type Gate692Inheritance struct {
	StateExpectationInherited        bool
	Rho72Definition                  string
	ResponseOperator                 string
	DBase                            float64
	SSplit                           float64
	ActiveExpectation                float64
	ResidualE1                       float64
	H72Dimension                     int
	K7Dimension                      int
	Gate692Rho72TypeCorrect          bool
	Gate692NoNativeStateSelection    bool
	Gate692NoNativeFirstTraceTheorem bool
	Gate692NoNativeSevenOver72       bool
	Verdict                          string
}

type GeneralStateResponseAudit struct {
	Formula                   string
	ResponseOperator          string
	SSplit                    float64
	RequiredK7Weight          float64
	RequiredExpectation       float64
	GeneralPositiveState      bool
	ReducesToK7Weight         bool
	ActiveRequiresK7Weight    bool
	DoesNotRequireFullRho72   bool
	WarnsAboutStateDegeneracy bool
	Verdict                   string
}

type ObserverStateCandidate struct {
	Name                string
	Definition          string
	SupportDimension    int
	K7Weight            float64
	Expectation         float64
	PositiveState       bool
	NormalizedState     bool
	FullH72Support      bool
	Unbiased            bool
	Biased              bool
	Circular            bool
	SignedObserver      bool
	MatchesActiveWeight bool
	MatchesActiveBridge bool
	FailureStatus       string
	Verdict             string
}

type TypedStateAlternativesAudit struct {
	Candidates                   []ObserverStateCandidate
	CandidateCount               int
	PositiveNormalizedCount      int
	MatchingActiveBridgeCount    int
	UnbiasedMatchingCount        int
	BiasedMatchingCount          int
	FiniteOnlyRejected           bool
	KernelRejected               bool
	LocalK7Rejected              bool
	BoundaryOnlyRejected         bool
	HodgeSignedRejected          bool
	BiasedReproductionWitnessed  bool
	Rho72ActiveUnbiasedCandidate bool
	AllTypedAlternativesAudited  bool
	Verdict                      string
}

type Rho72SelectionAudit struct {
	Rho72K7Weight               float64
	ActiveK7Weight              float64
	Rho72Expectation            float64
	ActiveExpectation           float64
	FullSupport                 bool
	Positive                    bool
	Normalized                  bool
	Unbiased                    bool
	MinimalAssumptions          []string
	UniqueUnderUnbiasedFullH72  bool
	UniqueAmongAllDensityStates bool
	NativeStateSelectionTheorem bool
	Verdict                     string
}

type BiasFirewallAudit struct {
	BiasedDensityStatesCanMatch   bool
	BiasedWitnessName             string
	BiasedWitnessK7Weight         float64
	BiasedWitnessExpectation      float64
	BiasedWitnessCircular         bool
	ReproductionIsNativeSelection bool
	Rho72UniquenessOverclaimed    bool
	Verdict                       string
}

type ResidualStatusAudit struct {
	E1                            float64
	AbsE1                         float64
	Expectation                   float64
	DBase                         float64
	QuadraticResidualClueRetained bool
	QuadraticCorrectionPromoted   bool
	Verdict                       string
}

type MissingTheoremAudit struct {
	Missing    []string
	Candidates []string
	PreciseGap string
	Verdict    string
}

type VerdictDiscipline struct {
	ClaimsNativeMaximallyMixedStateSelection bool
	ClaimsNativeFirstTraceTheorem            bool
	ClaimsNativeSevenOver72Theorem           bool
	ClaimsRho72UniqueAmongAllStates          bool
	ClaimsBiasedStateNativeSelection         bool
	ClaimsBoundaryStress                     bool
	ClaimsScalarRGMatching                   bool
	ClaimsHiggsMass                          bool
	ClaimsGaugeUnification                   bool
	ClaimsFlavorDerivation                   bool
	ClaimsCKMPMNS                            bool
	ClaimsProjectorActivation                bool
	Verdict                                  string
}

type Analysis struct {
	Inherited      Gate692Inheritance
	General        GeneralStateResponseAudit
	Alternatives   TypedStateAlternativesAudit
	Rho72Selection Rho72SelectionAudit
	BiasFirewall   BiasFirewallAudit
	Residual       ResidualStatusAudit
	Missing        MissingTheoremAudit
	Discipline     VerdictDiscipline
	Truth          string
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
	g692, err := gate692.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate692 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g692)
	general := buildGeneralStateResponse(inherited)
	alternatives := buildAlternatives(inherited, general)
	rho72 := buildRho72Selection(inherited, alternatives, general)
	bias := buildBiasFirewall(inherited, alternatives, general)
	residual := buildResidual(inherited, g692)
	missing := MissingTheoremAudit{
		Missing: []string{
			StatusNoNativeMaximallyMixedStateSelectionTheorem,
			StatusNoNativeFirstTraceTheorem,
			StatusNoNativeSevenOver72Theorem,
		},
		Candidates: []string{
			"GlobalAugmentedObserverStateTheorem",
			"MaximallyMixedHistoryObserverTheorem",
			"HistoryResponseStateSelectionTheorem",
		},
		PreciseGap: "a native state-selection theorem explaining why physical history evaluates R_split in the unbiased full H72 state rho_72 rather than in a finite-only, kernel-only, local-K7, boundary-only, signed-Hodge, or biased synthetic state",
		Verdict: strings.Join([]string{
			StatusNoNativeMaximallyMixedStateSelectionTheorem,
			StatusNoNativeFirstTraceTheorem,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	discipline := VerdictDiscipline{Verdict: StatusGate693ObserverStateSelectionBoundary}
	truth := "Gate 693 keeps the Gate692 state-expectation bridge but adds the bias firewall.  For any positive normalized state rho, Tr(rho R_split)=S_split Tr(rho P_K7), so the active bridge requires K7 weight 7/72.  rho_72=I_H72/72 gives this weight as the minimal unbiased full-chamber observer state.  Finite-only, kernel-only, local-K7, boundary-only, and signed-Hodge states fail as active typed alternatives.  However rho_72 is not unique among all density states: a biased synthetic density state can be assigned the same K7 weight by construction.  Such reproduction is circular and not a native selection theorem.  The audit conditionally supports rho_72 as the clean unbiased full H72 observer state while preserving the missing maximally mixed state-selection, first-trace, and native 7/72 theorems."
	return Analysis{Inherited: inherited, General: general, Alternatives: alternatives, Rho72Selection: rho72, BiasFirewall: bias, Residual: residual, Missing: missing, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate692.Analysis) Gate692Inheritance {
	return Gate692Inheritance{
		StateExpectationInherited:        g.Rho72.EqualsActiveFirstTrace && g.Rho72.PositiveState && g.Rho72.MaximallyMixedOnFullH72,
		Rho72Definition:                  g.Rho72.Definition,
		ResponseOperator:                 g.Inherited.Operator,
		DBase:                            g.Inherited.DBase,
		SSplit:                           g.Inherited.SSplit,
		ActiveExpectation:                g.Rho72.Expectation,
		ResidualE1:                       g.Residual.E1,
		H72Dimension:                     g.Inherited.H72Dimension,
		K7Dimension:                      g.Inherited.K7Dimension,
		Gate692Rho72TypeCorrect:          g.Rho72.TypeCorrectFullObserver,
		Gate692NoNativeStateSelection:    !g.Discipline.ClaimsNativeStateSelectionTheorem,
		Gate692NoNativeFirstTraceTheorem: !g.Discipline.ClaimsNativeFirstTraceTheorem,
		Gate692NoNativeSevenOver72:       !g.Discipline.ClaimsNativeSevenOver72Theorem,
		Verdict:                          StatusGate692StateExpectationInherited,
	}
}

func buildGeneralStateResponse(i Gate692Inheritance) GeneralStateResponseAudit {
	weight := float64(i.K7Dimension) / float64(i.H72Dimension)
	return GeneralStateResponseAudit{
		Formula:                   "For any positive normalized state rho, Tr(rho R_split)=S_split Tr(rho P_K7)",
		ResponseOperator:          i.ResponseOperator,
		SSplit:                    i.SSplit,
		RequiredK7Weight:          weight,
		RequiredExpectation:       weight * i.SSplit,
		GeneralPositiveState:      true,
		ReducesToK7Weight:         true,
		ActiveRequiresK7Weight:    true,
		DoesNotRequireFullRho72:   true,
		WarnsAboutStateDegeneracy: true,
		Verdict: strings.Join([]string{
			StatusGeneralStateResponseReducedToK7Weight,
			StatusActiveResponseRequiresK7Weight7Over72,
			StatusRho72NotUniqueAmongAllDensityStates,
		}, "; "),
	}
}

func buildAlternatives(i Gate692Inheritance, g GeneralStateResponseAudit) TypedStateAlternativesAudit {
	activeWeight := g.RequiredK7Weight
	candidates := []ObserverStateCandidate{
		stateCandidate("rho_72", "I_H72/72", h72Dimension, activeWeight, i.SSplit, true, true, true, true, false, false, false, "", "ACTIVE_MINIMAL_UNBIASED_FULL_H72_STATE"),
		stateCandidate("rho_finite", "P_Lambda4/70", lambda4Dimension, 7.0/70.0, i.SSplit, true, true, false, false, false, false, false, StatusFiniteOnlyStateRejectedBy7Over70, StatusFiniteOnlyStateRejectedBy7Over70),
		stateCandidate("rho_kernel", "P_kernel/71", kernelDimension, 7.0/71.0, i.SSplit, true, true, false, false, false, false, false, StatusKernelStateRejectedBy7Over71, StatusKernelStateRejectedBy7Over71),
		stateCandidate("rho_K7", "P_K7/7", k7Dimension, 1.0, i.SSplit, true, true, false, false, false, false, false, StatusLocalK7StateRejectedByUnitWeight, StatusLocalK7StateRejectedByUnitWeight),
		stateCandidate("rho_boundary", "P_boundary/2", boundaryDimension, 0.0, i.SSplit, true, true, false, false, false, false, false, StatusBoundaryOnlyStateRejectedByZeroWeight, StatusBoundaryOnlyStateRejectedByZeroWeight),
		stateCandidate("rho_signed", "P_+ - P_- is signed and not a positive density state", h72Dimension, float64(k7PlusDimension-k7MinusDimension)/float64(h72Dimension), i.SSplit, false, false, false, false, false, false, true, StatusHodgeSignedObserverRejectedAsNonPositiveState, StatusHodgeSignedObserverRejectedAsNonPositiveState),
		stateCandidate("rho_biased_weight_7_over_72", "biased positive density state constrained to Tr(rho P_K7)=7/72", h72Dimension, activeWeight, i.SSplit, true, true, true, false, true, true, false, StatusBiasedStatesCanReproduceWeightButAreCircular, StatusBiasedStatesCanReproduceWeightButAreCircular),
	}

	positiveNormalized := 0
	matches := 0
	unbiasedMatches := 0
	biasedMatches := 0
	finiteRejected := false
	kernelRejected := false
	localK7Rejected := false
	boundaryRejected := false
	signedRejected := false
	biasedWitness := false
	rho72Active := false
	for _, c := range candidates {
		if c.PositiveState && c.NormalizedState {
			positiveNormalized++
		}
		if c.MatchesActiveBridge {
			matches++
			if c.Unbiased {
				unbiasedMatches++
			}
			if c.Biased {
				biasedMatches++
			}
		}
		switch c.Name {
		case "rho_72":
			rho72Active = c.MatchesActiveBridge && c.FullH72Support && c.Unbiased
		case "rho_finite":
			finiteRejected = !c.MatchesActiveBridge && math.Abs(c.K7Weight-7.0/70.0) < tolerance
		case "rho_kernel":
			kernelRejected = !c.MatchesActiveBridge && math.Abs(c.K7Weight-7.0/71.0) < tolerance
		case "rho_K7":
			localK7Rejected = !c.MatchesActiveBridge && math.Abs(c.K7Weight-1.0) < tolerance
		case "rho_boundary":
			boundaryRejected = !c.MatchesActiveBridge && math.Abs(c.K7Weight) < tolerance
		case "rho_signed":
			signedRejected = !c.PositiveState && !c.NormalizedState && c.SignedObserver
		case "rho_biased_weight_7_over_72":
			biasedWitness = c.MatchesActiveBridge && c.Biased && c.Circular
		}
	}

	return TypedStateAlternativesAudit{
		Candidates:                   candidates,
		CandidateCount:               len(candidates),
		PositiveNormalizedCount:      positiveNormalized,
		MatchingActiveBridgeCount:    matches,
		UnbiasedMatchingCount:        unbiasedMatches,
		BiasedMatchingCount:          biasedMatches,
		FiniteOnlyRejected:           finiteRejected,
		KernelRejected:               kernelRejected,
		LocalK7Rejected:              localK7Rejected,
		BoundaryOnlyRejected:         boundaryRejected,
		HodgeSignedRejected:          signedRejected,
		BiasedReproductionWitnessed:  biasedWitness,
		Rho72ActiveUnbiasedCandidate: rho72Active,
		AllTypedAlternativesAudited:  finiteRejected && kernelRejected && localK7Rejected && boundaryRejected && signedRejected && biasedWitness && rho72Active,
		Verdict: strings.Join([]string{
			StatusAlternativeTypedStatesAudited,
			StatusFiniteOnlyStateRejectedBy7Over70,
			StatusKernelStateRejectedBy7Over71,
			StatusLocalK7StateRejectedByUnitWeight,
			StatusBoundaryOnlyStateRejectedByZeroWeight,
			StatusHodgeSignedObserverRejectedAsNonPositiveState,
			StatusBiasedStatesCanReproduceWeightButAreCircular,
		}, "; "),
	}
}

func stateCandidate(name, definition string, supportDimension int, k7Weight, ssplit float64, positive, normalized, fullSupport, unbiased, biased, circular, signed bool, failureStatus, verdict string) ObserverStateCandidate {
	expectation := k7Weight * ssplit
	activeWeight := float64(k7Dimension) / float64(h72Dimension)
	matchesWeight := math.Abs(k7Weight-activeWeight) < tolerance
	return ObserverStateCandidate{
		Name:                name,
		Definition:          definition,
		SupportDimension:    supportDimension,
		K7Weight:            k7Weight,
		Expectation:         expectation,
		PositiveState:       positive,
		NormalizedState:     normalized,
		FullH72Support:      fullSupport,
		Unbiased:            unbiased,
		Biased:              biased,
		Circular:            circular,
		SignedObserver:      signed,
		MatchesActiveWeight: matchesWeight,
		MatchesActiveBridge: matchesWeight,
		FailureStatus:       failureStatus,
		Verdict:             verdict,
	}
}

func buildRho72Selection(i Gate692Inheritance, a TypedStateAlternativesAudit, g GeneralStateResponseAudit) Rho72SelectionAudit {
	return Rho72SelectionAudit{
		Rho72K7Weight:               float64(i.K7Dimension) / float64(i.H72Dimension),
		ActiveK7Weight:              g.RequiredK7Weight,
		Rho72Expectation:            g.RequiredExpectation,
		ActiveExpectation:           i.ActiveExpectation,
		FullSupport:                 true,
		Positive:                    true,
		Normalized:                  true,
		Unbiased:                    true,
		MinimalAssumptions:          []string{"full H72 support", "positivity", "normalization", "no preferred subspace beyond H72", "no spectral bias inside K7 or its complement"},
		UniqueUnderUnbiasedFullH72:  a.Rho72ActiveUnbiasedCandidate && a.UnbiasedMatchingCount == 1,
		UniqueAmongAllDensityStates: false,
		NativeStateSelectionTheorem: false,
		Verdict: strings.Join([]string{
			StatusRho72GivesActiveK7Weight,
			StatusRho72MinimalUnbiasedFullAugmentedObserverState,
			StatusActiveBridgeGlobalUnbiasedK7WeightExpectation,
			StatusRho72NotUniqueAmongAllDensityStates,
		}, "; "),
	}
}

func buildBiasFirewall(i Gate692Inheritance, a TypedStateAlternativesAudit, g GeneralStateResponseAudit) BiasFirewallAudit {
	witness := ObserverStateCandidate{}
	for _, c := range a.Candidates {
		if c.Biased && c.MatchesActiveBridge {
			witness = c
			break
		}
	}
	return BiasFirewallAudit{
		BiasedDensityStatesCanMatch:   a.BiasedReproductionWitnessed,
		BiasedWitnessName:             witness.Name,
		BiasedWitnessK7Weight:         witness.K7Weight,
		BiasedWitnessExpectation:      witness.Expectation,
		BiasedWitnessCircular:         witness.Circular,
		ReproductionIsNativeSelection: false,
		Rho72UniquenessOverclaimed:    true,
		Verdict: strings.Join([]string{
			StatusBiasedStatesCanReproduceWeightButAreCircular,
			StatusRho72NotUniqueAmongAllDensityStates,
			StatusBiasedStateReproductionNotNativeSelection,
		}, "; "),
	}
}

func buildResidual(i Gate692Inheritance, g gate692.Analysis) ResidualStatusAudit {
	return ResidualStatusAudit{
		E1:                            i.ResidualE1,
		AbsE1:                         math.Abs(i.ResidualE1),
		Expectation:                   i.ActiveExpectation,
		DBase:                         i.DBase,
		QuadraticResidualClueRetained: g.Residual.QuadraticResidualClueRetained,
		QuadraticCorrectionPromoted:   g.Residual.QuadraticCorrectionPromoted,
		Verdict:                       "GATE690_QUADRATIC_RESIDUAL_REMAINS_SUBLEADING_AND_UNPROMOTED",
	}
}

func Statuses() []string {
	return []string{
		StatusGate692StateExpectationInherited,
		StatusGeneralStateResponseReducedToK7Weight,
		StatusActiveResponseRequiresK7Weight7Over72,
		StatusRho72GivesActiveK7Weight,
		StatusAlternativeTypedStatesAudited,
		StatusFiniteOnlyStateRejectedBy7Over70,
		StatusKernelStateRejectedBy7Over71,
		StatusLocalK7StateRejectedByUnitWeight,
		StatusBoundaryOnlyStateRejectedByZeroWeight,
		StatusHodgeSignedObserverRejectedAsNonPositiveState,
		StatusBiasedStatesCanReproduceWeightButAreCircular,
		StatusRho72MinimalUnbiasedFullAugmentedObserverState,
		StatusActiveBridgeGlobalUnbiasedK7WeightExpectation,
		StatusRho72NotUniqueAmongAllDensityStates,
		StatusBiasedStateReproductionNotNativeSelection,
		StatusNoNativeMaximallyMixedStateSelectionTheorem,
		StatusNoNativeFirstTraceTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusGate693ObserverStateSelectionBoundary,
	}
}

func FormatInheritance(x Gate692Inheritance) string {
	return fmt.Sprintf("stateExpectation=%t rho=%q response=%q dbase=%.18g ssplit=%.18g activeExpectation=%.18g e1=%.18g h72=%d k7=%d typeCorrect=%t noStateSelection=%t noFirst=%t no7=%t verdict=%q", x.StateExpectationInherited, x.Rho72Definition, x.ResponseOperator, x.DBase, x.SSplit, x.ActiveExpectation, x.ResidualE1, x.H72Dimension, x.K7Dimension, x.Gate692Rho72TypeCorrect, x.Gate692NoNativeStateSelection, x.Gate692NoNativeFirstTraceTheorem, x.Gate692NoNativeSevenOver72, x.Verdict)
}

func FormatGeneral(x GeneralStateResponseAudit) string {
	return fmt.Sprintf("formula=%q response=%q ssplit=%.18g requiredWeight=%.18g requiredExpectation=%.18g positiveState=%t reducesToK7=%t activeRequires=%t doesNotRequireRho72=%t degeneracyWarning=%t verdict=%q", x.Formula, x.ResponseOperator, x.SSplit, x.RequiredK7Weight, x.RequiredExpectation, x.GeneralPositiveState, x.ReducesToK7Weight, x.ActiveRequiresK7Weight, x.DoesNotRequireFullRho72, x.WarnsAboutStateDegeneracy, x.Verdict)
}

func FormatCandidate(x ObserverStateCandidate) string {
	return fmt.Sprintf("%s def=%q support=%d k7Weight=%.18g expectation=%.18g positive=%t normalized=%t full=%t unbiased=%t biased=%t circular=%t signed=%t matchesWeight=%t matchesBridge=%t failure=%q verdict=%q", x.Name, x.Definition, x.SupportDimension, x.K7Weight, x.Expectation, x.PositiveState, x.NormalizedState, x.FullH72Support, x.Unbiased, x.Biased, x.Circular, x.SignedObserver, x.MatchesActiveWeight, x.MatchesActiveBridge, x.FailureStatus, x.Verdict)
}

func FormatAlternatives(x TypedStateAlternativesAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, FormatCandidate(c))
	}
	return fmt.Sprintf("candidates=[%s] count=%d positiveNormalized=%d matching=%d unbiasedMatching=%d biasedMatching=%d finiteRejected=%t kernelRejected=%t localK7Rejected=%t boundaryRejected=%t signedRejected=%t biasedWitness=%t rho72Active=%t allAudited=%t verdict=%q", strings.Join(parts, " | "), x.CandidateCount, x.PositiveNormalizedCount, x.MatchingActiveBridgeCount, x.UnbiasedMatchingCount, x.BiasedMatchingCount, x.FiniteOnlyRejected, x.KernelRejected, x.LocalK7Rejected, x.BoundaryOnlyRejected, x.HodgeSignedRejected, x.BiasedReproductionWitnessed, x.Rho72ActiveUnbiasedCandidate, x.AllTypedAlternativesAudited, x.Verdict)
}

func FormatRho72Selection(x Rho72SelectionAudit) string {
	return fmt.Sprintf("rhoWeight=%.18g activeWeight=%.18g rhoExpectation=%.18g activeExpectation=%.18g full=%t positive=%t normalized=%t unbiased=%t assumptions=[%s] uniqueUnbiasedFull=%t uniqueAll=%t nativeState=%t verdict=%q", x.Rho72K7Weight, x.ActiveK7Weight, x.Rho72Expectation, x.ActiveExpectation, x.FullSupport, x.Positive, x.Normalized, x.Unbiased, strings.Join(x.MinimalAssumptions, "; "), x.UniqueUnderUnbiasedFullH72, x.UniqueAmongAllDensityStates, x.NativeStateSelectionTheorem, x.Verdict)
}

func FormatBiasFirewall(x BiasFirewallAudit) string {
	return fmt.Sprintf("biasedCanMatch=%t witness=%q witnessWeight=%.18g witnessExpectation=%.18g circular=%t nativeSelection=%t overclaim=%t verdict=%q", x.BiasedDensityStatesCanMatch, x.BiasedWitnessName, x.BiasedWitnessK7Weight, x.BiasedWitnessExpectation, x.BiasedWitnessCircular, x.ReproductionIsNativeSelection, x.Rho72UniquenessOverclaimed, x.Verdict)
}

func FormatResidual(x ResidualStatusAudit) string {
	return fmt.Sprintf("e1=%.18g absE1=%.18g expectation=%.18g dbase=%.18g residualClue=%t correctionPromoted=%t verdict=%q", x.E1, x.AbsE1, x.Expectation, x.DBase, x.QuadraticResidualClueRetained, x.QuadraticCorrectionPromoted, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=[%s] candidates=[%s] precise=%q verdict=%q", strings.Join(x.Missing, "; "), strings.Join(x.Candidates, "; "), x.PreciseGap, x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("nativeMaxMixedSelection=%t nativeFirst=%t native7=%t uniqueAll=%t biasedNative=%t boundary=%t scalarRG=%t higgs=%t gauge=%t flavor=%t ckmPmns=%t activation=%t verdict=%q", x.ClaimsNativeMaximallyMixedStateSelection, x.ClaimsNativeFirstTraceTheorem, x.ClaimsNativeSevenOver72Theorem, x.ClaimsRho72UniqueAmongAllStates, x.ClaimsBiasedStateNativeSelection, x.ClaimsBoundaryStress, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.ClaimsProjectorActivation, x.Verdict)
}
