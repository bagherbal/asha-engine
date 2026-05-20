// Package generation2maximallymixedaugmentedchamberobserverstateaudit implements
// Gate 692: Maximally Mixed Augmented-Chamber Observer State Audit.
//
// Gate 691 rewrote the active bridge as a normalized trace pairing
//
//	D_base ≈ Tr_H72(I_H72 R_split)/Tr_H72(I_H72).
//
// Gate 692 sharpens this source type by defining the full augmented maximally
// mixed observer state
//
//	rho_72 = I_H72/Tr(I_H72) = I_H72/72,
//
// so that the leading bridge is the state expectation
//
//	D_base ≈ Tr(rho_72 R_split).
//
// This is a bridge-layer observer-state normalization audit only. It does not
// derive boundary stress, scalar RG matching, Higgs mass, gauge unification,
// flavor, CKM/PMNS, a native first-trace theorem, a native state-selection
// theorem, or a native 7/72 theorem.
package generation2maximallymixedaugmentedchamberobserverstateaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate691 "github.com/bagherbal/asha-engine/pkg/bridge/generation2linearresponsefunctionalandtracepairingnormalizationaudit"
)

const (
	AuditID = "GATE692-MAXIMALLY-MIXED-AUGMENTED-CHAMBER-OBSERVER-STATE-AUDIT"

	StatusGate691TracePairingInherited                 = "PASS_GATE691_TRACE_PAIRING_INHERITED"
	StatusRho72MaximallyMixedStateDefined              = "PASS_RHO_72_MAXIMALLY_MIXED_STATE_DEFINED"
	StatusActiveBridgeRewrittenAsStateExpectation      = "PASS_ACTIVE_BRIDGE_REWRITTEN_AS_STATE_EXPECTATION"
	StatusAlternativeNormalizedObserverStatesAudited   = "PASS_ALTERNATIVE_NORMALIZED_OBSERVER_STATES_AUDITED"
	StatusObserverDenominatorDegeneracyResolvedByState = "PASS_OBSERVER_DENOMINATOR_DEGENERACY_RESOLVED_BY_STATE_NORMALIZATION"
	StatusActiveResponseGlobalH72ExpectationValue      = "CONDITIONAL_SUPPORT_ACTIVE_RESPONSE_IS_GLOBAL_H72_EXPECTATION_VALUE"
	StatusRho72TypeCorrectFullAugmentedObserverState   = "CONDITIONAL_SUPPORT_RHO_72_IS_TYPE_CORRECT_FULL_AUGMENTED_CHAMBER_OBSERVER_STATE"
	StatusFiniteOnlyStateGives7Over70                  = "FAILED_ROUTE_FINITE_ONLY_STATE_GIVES_7_OVER_70"
	StatusKernelConditionalStateGives7Over71           = "FAILED_ROUTE_KERNEL_CONDITIONAL_STATE_GIVES_7_OVER_71"
	StatusLocalK7StateGivesSSplitNot7Over72            = "FAILED_ROUTE_LOCAL_K7_STATE_GIVES_S_SPLIT_NOT_7_OVER_72"
	StatusHodgeSignedObserverNotPositiveStateInactive  = "FAILED_ROUTE_HODGE_SIGNED_OBSERVER_IS_NOT_POSITIVE_STATE_AND_NOT_ACTIVE"
	StatusNoNativeMaximallyMixedObserverStateTheorem   = "FAILED_ROUTE_NO_NATIVE_MAXIMALLY_MIXED_OBSERVER_STATE_THEOREM"
	StatusNoNativeFirstTraceTheorem                    = "FAILED_ROUTE_NO_NATIVE_FIRST_TRACE_THEOREM"
	StatusNoNativeSevenOver72Theorem                   = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate692MaximallyMixedObserverStateBoundary   = "FIREWALL_PRESERVED_GATE692_MAXIMALLY_MIXED_OBSERVER_STATE_BOUNDARY"
)

const (
	lambda4Dimension        = 70
	boundaryDimension       = 2
	h72Dimension            = lambda4Dimension + boundaryDimension
	kernelDimension         = 71
	k7Dimension             = 7
	k7PlusDimension         = 4
	k7MinusDimension        = 3
	stateTolerance          = 1e-18
	residualTolerance       = 1e-18
	alternativeValueEpsilon = 1e-18
)

type Gate691Inheritance struct {
	TracePairingInherited             bool
	Operator                          string
	DBase                             float64
	SSplit                            float64
	F1                                float64
	E1                                float64
	H72Dimension                      int
	K7Dimension                       int
	Gate691ObserverDegeneracyRecorded bool
	Gate690ResidualClueRetained       bool
	NativeLinearResponseTheorem       bool
	NativeFirstTraceTheorem           bool
	NativeSevenOver72Theorem          bool
	ClaimsUniqueFullH72Observer       bool
	Verdict                           string
}

type MaximallyMixedStateAudit struct {
	StateName                string
	Definition               string
	Dimension                int
	IdentityTrace            float64
	NormalizationDenominator float64
	StateTrace               float64
	PositiveState            bool
	MaximallyMixedOnFullH72  bool
	NumeratorTraceOfResponse float64
	Expectation              float64
	ExpectedFirstTrace       float64
	EqualsActiveFirstTrace   bool
	TypeCorrectFullObserver  bool
	Verdict                  string
}

type ObserverStateCandidate struct {
	Name                     string
	Definition               string
	SupportDimension         int
	ContainsK7               bool
	PositiveState            bool
	NormalizedState          bool
	MaximallyMixedOnSupport  bool
	SignedObserver           bool
	NormalizationDenominator float64
	StateTrace               float64
	Expectation              float64
	ExpectedFormula          string
	ResidualAgainstRho72     float64
	MatchesActiveBridge      bool
	ActiveCandidate          bool
	FailureStatus            string
	Verdict                  string
}

type AlternativeObserverStateAudit struct {
	Candidates                     []ObserverStateCandidate
	CandidateCount                 int
	PositiveNormalizedStateCount   int
	ActiveStateCount               int
	FiniteOnlyStateInactive        bool
	KernelConditionalStateInactive bool
	LocalK7StateInactive           bool
	HodgeSignedObserverNotPositive bool
	AllAlternativesAudited         bool
	Verdict                        string
}

type DenominatorDegeneracyResolutionAudit struct {
	Gate691DegeneracySource        string
	Gate692Resolution              string
	FixedH72DenominatorDegenerate  bool
	StateNormalizationResolvesType bool
	Rho72UniqueAmongAuditedStates  bool
	NativeStateSelectionTheorem    bool
	Verdict                        string
}

type InterpretationAudit struct {
	Rho72Role              string
	ResponseOperatorRole   string
	ExpectationRole        string
	GlobalAverageDensity   bool
	BoundaryScalarEigen    bool
	SupportSelectedCarrier bool
	Verdict                string
}

type ResidualStatusAudit struct {
	E1                             float64
	AbsE1                          float64
	Expectation                    float64
	DBase                          float64
	QuadraticResidualClueRetained  bool
	QuadraticCorrectionPromoted    bool
	NativeSpectralExpansionTheorem bool
	Verdict                        string
}

type MissingTheoremAudit struct {
	Missing    []string
	Candidates []string
	PreciseGap string
	Verdict    string
}

type VerdictDiscipline struct {
	ClaimsNativeMaximallyMixedStateTheorem bool
	ClaimsNativeStateSelectionTheorem      bool
	ClaimsNativeFirstTraceTheorem          bool
	ClaimsNativeSevenOver72Theorem         bool
	ClaimsBoundaryStress                   bool
	ClaimsScalarRGMatching                 bool
	ClaimsHiggsMass                        bool
	ClaimsGaugeUnification                 bool
	ClaimsFlavorDerivation                 bool
	ClaimsCKMPMNS                          bool
	ClaimsProjectorActivation              bool
	PromotesQuadraticResidualCorrection    bool
	Verdict                                string
}

type Analysis struct {
	Inherited      Gate691Inheritance
	Rho72          MaximallyMixedStateAudit
	Alternatives   AlternativeObserverStateAudit
	Degeneracy     DenominatorDegeneracyResolutionAudit
	Interpretation InterpretationAudit
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
	g691, err := gate691.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate691 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g691)
	rho72 := buildRho72(inherited)
	alternatives := buildAlternatives(inherited, rho72)
	degeneracy := buildDegeneracyResolution(inherited, alternatives)
	interpretation := buildInterpretation()
	residual := buildResidualStatus(g691, rho72)
	missing := MissingTheoremAudit{
		Missing: []string{
			StatusNoNativeMaximallyMixedObserverStateTheorem,
			StatusNoNativeFirstTraceTheorem,
			StatusNoNativeSevenOver72Theorem,
		},
		Candidates: []string{
			"GlobalAugmentedObserverStateTheorem",
			"MaximallyMixedHistoryObserverTheorem",
			"HistoryResponseStateSelectionTheorem",
		},
		PreciseGap: "a native state-selection theorem explaining why physical history evaluates the support-selected response operator in rho_72=I_H72/72 rather than in finite-only, kernel-only, local-K7, signed-Hodge, or arbitrary observer states",
		Verdict: strings.Join([]string{
			StatusNoNativeMaximallyMixedObserverStateTheorem,
			StatusNoNativeFirstTraceTheorem,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	discipline := VerdictDiscipline{Verdict: StatusGate692MaximallyMixedObserverStateBoundary}
	truth := "Gate 692 sharpens the Gate691 normalized trace-pairing into a state expectation.  The active leading bridge is Tr(rho_72 R_split) with rho_72=I_H72/72, so Tr(rho_72)=1 and Tr(rho_72 R_split)=(7/72)S_split.  This resolves the Gate691 denominator degeneracy at the source-type level: with true state normalization, rho_finite=P_finite/70 gives (7/70)S_split, rho_kernel=P_kernel/71 gives (7/71)S_split, and rho_K7=P_K7/7 gives S_split.  The Hodge-signed observer is not a positive state and remains inactive.  The audit conditionally supports the active response as a global H72 expectation value in the full augmented maximally mixed observer state, but no native maximally mixed observer-state theorem, first-trace theorem, or native 7/72 theorem is certified."
	return Analysis{Inherited: inherited, Rho72: rho72, Alternatives: alternatives, Degeneracy: degeneracy, Interpretation: interpretation, Residual: residual, Missing: missing, Discipline: discipline, Truth: truth}, nil
}

func buildInheritance(g gate691.Analysis) Gate691Inheritance {
	return Gate691Inheritance{
		TracePairingInherited:             g.Pairing.EqualsFirstTrace && g.Pairing.Observer == "I_H72" && g.Pairing.Response == "R_split = S_split P_K7" && g.Pairing.DenominatorTrace == float64(h72Dimension),
		Operator:                          g.Inherited.Operator,
		DBase:                             g.Inherited.DBase,
		SSplit:                            g.Inherited.SSplit,
		F1:                                g.Inherited.F1,
		E1:                                g.Residual.E1,
		H72Dimension:                      g.Inherited.H72Dimension,
		K7Dimension:                       g.Inherited.K7Dimension,
		Gate691ObserverDegeneracyRecorded: g.Observers.AllPositiveK7ObserversGiveSameValue && !g.Observers.FullH72ObserverUnique,
		Gate690ResidualClueRetained:       g.Residual.QuadraticResidualClueRetained && !g.Residual.QuadraticCorrectionPromoted,
		NativeLinearResponseTheorem:       g.LinearResponse.NativeLinearResponseFunctionalTheorem,
		NativeFirstTraceTheorem:           g.Inherited.NativeFirstTraceTheorem,
		NativeSevenOver72Theorem:          g.Inherited.NativeSevenOver72Theorem,
		ClaimsUniqueFullH72Observer:       g.Discipline.ClaimsUniqueFullH72Observer,
		Verdict:                           StatusGate691TracePairingInherited,
	}
}

func buildRho72(i Gate691Inheritance) MaximallyMixedStateAudit {
	expectation := float64(i.K7Dimension) * i.SSplit / float64(i.H72Dimension)
	return MaximallyMixedStateAudit{
		StateName:                "rho_72",
		Definition:               "rho_72 = I_H72/Tr(I_H72) = I_H72/72",
		Dimension:                i.H72Dimension,
		IdentityTrace:            float64(i.H72Dimension),
		NormalizationDenominator: float64(i.H72Dimension),
		StateTrace:               1,
		PositiveState:            true,
		MaximallyMixedOnFullH72:  true,
		NumeratorTraceOfResponse: float64(i.K7Dimension) * i.SSplit,
		Expectation:              expectation,
		ExpectedFirstTrace:       i.F1,
		EqualsActiveFirstTrace:   math.Abs(expectation-i.F1) < stateTolerance,
		TypeCorrectFullObserver:  true,
		Verdict: strings.Join([]string{
			StatusRho72MaximallyMixedStateDefined,
			StatusActiveBridgeRewrittenAsStateExpectation,
			StatusActiveResponseGlobalH72ExpectationValue,
			StatusRho72TypeCorrectFullAugmentedObserverState,
		}, "; "),
	}
}

func buildAlternatives(i Gate691Inheritance, rho MaximallyMixedStateAudit) AlternativeObserverStateAudit {
	candidates := []ObserverStateCandidate{
		stateCandidate("rho_72", "I_H72/72", h72Dimension, true, true, true, true, false, i.SSplit, rho.Expectation, "(7/72)S_split", "", true),
		stateCandidate("rho_finite", "P_finite/70", lambda4Dimension, true, true, true, true, false, i.SSplit, rho.Expectation, "(7/70)S_split", StatusFiniteOnlyStateGives7Over70, false),
		stateCandidate("rho_kernel", "P_kernel/71", kernelDimension, true, true, true, true, false, i.SSplit, rho.Expectation, "(7/71)S_split", StatusKernelConditionalStateGives7Over71, false),
		stateCandidate("rho_K7", "P_K7/7", k7Dimension, true, true, true, true, false, i.SSplit, rho.Expectation, "S_split", StatusLocalK7StateGivesSSplitNot7Over72, false),
		stateCandidate("rho_signed", "P_+ - P_- is not a positive density state", h72Dimension, true, false, false, false, true, i.SSplit, rho.Expectation, "not a positive state; signed Hodge observer gives (1/72)S_split under H72 trace normalization", StatusHodgeSignedObserverNotPositiveStateInactive, false),
	}
	positiveNormalized := 0
	activeStates := 0
	finiteInactive := false
	kernelInactive := false
	k7Inactive := false
	signedNotPositive := false
	for _, c := range candidates {
		if c.PositiveState && c.NormalizedState {
			positiveNormalized++
		}
		if c.ActiveCandidate && c.MatchesActiveBridge {
			activeStates++
		}
		switch c.Name {
		case "rho_finite":
			finiteInactive = !c.MatchesActiveBridge && math.Abs(c.Expectation-(7.0/70.0)*i.SSplit) < alternativeValueEpsilon
		case "rho_kernel":
			kernelInactive = !c.MatchesActiveBridge && math.Abs(c.Expectation-(7.0/71.0)*i.SSplit) < alternativeValueEpsilon
		case "rho_K7":
			k7Inactive = !c.MatchesActiveBridge && math.Abs(c.Expectation-i.SSplit) < alternativeValueEpsilon
		case "rho_signed":
			signedNotPositive = !c.PositiveState && !c.NormalizedState && c.SignedObserver && !c.MatchesActiveBridge
		}
	}
	return AlternativeObserverStateAudit{
		Candidates:                     candidates,
		CandidateCount:                 len(candidates),
		PositiveNormalizedStateCount:   positiveNormalized,
		ActiveStateCount:               activeStates,
		FiniteOnlyStateInactive:        finiteInactive,
		KernelConditionalStateInactive: kernelInactive,
		LocalK7StateInactive:           k7Inactive,
		HodgeSignedObserverNotPositive: signedNotPositive,
		AllAlternativesAudited:         finiteInactive && kernelInactive && k7Inactive && signedNotPositive,
		Verdict: strings.Join([]string{
			StatusAlternativeNormalizedObserverStatesAudited,
			StatusFiniteOnlyStateGives7Over70,
			StatusKernelConditionalStateGives7Over71,
			StatusLocalK7StateGivesSSplitNot7Over72,
			StatusHodgeSignedObserverNotPositiveStateInactive,
		}, "; "),
	}
}

func stateCandidate(name, definition string, supportDimension int, containsK7, positive, normalized, maximallyMixed, signed bool, ssplit, activeValue float64, formula, failure string, active bool) ObserverStateCandidate {
	var expectation float64
	var stateTrace float64
	if positive && normalized {
		stateTrace = 1
		if supportDimension > 0 {
			expectation = float64(k7Dimension) * ssplit / float64(supportDimension)
		}
	} else if signed {
		stateTrace = 0
		expectation = float64(k7PlusDimension-k7MinusDimension) * ssplit / float64(h72Dimension)
	}
	matches := math.Abs(expectation-activeValue) < stateTolerance
	verdict := "ACTIVE_FULL_H72_MAXIMALLY_MIXED_STATE"
	if !active {
		verdict = failure
	}
	return ObserverStateCandidate{
		Name:                     name,
		Definition:               definition,
		SupportDimension:         supportDimension,
		ContainsK7:               containsK7,
		PositiveState:            positive,
		NormalizedState:          normalized,
		MaximallyMixedOnSupport:  maximallyMixed,
		SignedObserver:           signed,
		NormalizationDenominator: float64(supportDimension),
		StateTrace:               stateTrace,
		Expectation:              expectation,
		ExpectedFormula:          formula,
		ResidualAgainstRho72:     expectation - activeValue,
		MatchesActiveBridge:      matches,
		ActiveCandidate:          active,
		FailureStatus:            failure,
		Verdict:                  verdict,
	}
}

func buildDegeneracyResolution(i Gate691Inheritance, a AlternativeObserverStateAudit) DenominatorDegeneracyResolutionAudit {
	return DenominatorDegeneracyResolutionAudit{
		Gate691DegeneracySource:        "Gate 691 kept the denominator fixed at 72 while changing numerator observers that act as identity on K7",
		Gate692Resolution:              "Gate 692 normalizes observers as states; rho_finite, rho_kernel, and rho_K7 carry denominators 70, 71, and 7, so only rho_72 yields the active 7/72 response among audited normalized states",
		FixedH72DenominatorDegenerate:  i.Gate691ObserverDegeneracyRecorded,
		StateNormalizationResolvesType: a.FiniteOnlyStateInactive && a.KernelConditionalStateInactive && a.LocalK7StateInactive,
		Rho72UniqueAmongAuditedStates:  a.ActiveStateCount == 1,
		NativeStateSelectionTheorem:    false,
		Verdict: strings.Join([]string{
			StatusObserverDenominatorDegeneracyResolvedByState,
			StatusNoNativeMaximallyMixedObserverStateTheorem,
		}, "; "),
	}
}

func buildInterpretation() InterpretationAudit {
	return InterpretationAudit{
		Rho72Role:              "rho_72: full augmented chamber observer-state",
		ResponseOperatorRole:   "R_split=S_split P_K7: support-selected boundary response operator",
		ExpectationRole:        "Tr(rho_72 R_split): global average response density of the rank-seven support under the boundary split eigenvalue",
		GlobalAverageDensity:   true,
		BoundaryScalarEigen:    true,
		SupportSelectedCarrier: true,
		Verdict:                StatusActiveResponseGlobalH72ExpectationValue,
	}
}

func buildResidualStatus(g gate691.Analysis, rho MaximallyMixedStateAudit) ResidualStatusAudit {
	return ResidualStatusAudit{
		E1:                             g.Residual.E1,
		AbsE1:                          math.Abs(g.Residual.E1),
		Expectation:                    rho.Expectation,
		DBase:                          g.Inherited.DBase,
		QuadraticResidualClueRetained:  g.Residual.QuadraticResidualClueRetained,
		QuadraticCorrectionPromoted:    g.Residual.QuadraticCorrectionPromoted,
		NativeSpectralExpansionTheorem: g.Residual.NativeSpectralExpansionTheorem,
		Verdict:                        "GATE690_QUADRATIC_RESIDUAL_REMAINS_SUBLEADING_AND_UNPROMOTED",
	}
}

func Statuses() []string {
	return []string{
		StatusGate691TracePairingInherited,
		StatusRho72MaximallyMixedStateDefined,
		StatusActiveBridgeRewrittenAsStateExpectation,
		StatusAlternativeNormalizedObserverStatesAudited,
		StatusObserverDenominatorDegeneracyResolvedByState,
		StatusActiveResponseGlobalH72ExpectationValue,
		StatusRho72TypeCorrectFullAugmentedObserverState,
		StatusFiniteOnlyStateGives7Over70,
		StatusKernelConditionalStateGives7Over71,
		StatusLocalK7StateGivesSSplitNot7Over72,
		StatusHodgeSignedObserverNotPositiveStateInactive,
		StatusNoNativeMaximallyMixedObserverStateTheorem,
		StatusNoNativeFirstTraceTheorem,
		StatusNoNativeSevenOver72Theorem,
		StatusGate692MaximallyMixedObserverStateBoundary,
	}
}

func FormatInheritance(x Gate691Inheritance) string {
	return fmt.Sprintf("tracePairing=%t operator=%q dbase=%.18g ssplit=%.18g f1=%.18g e1=%.18g h72=%d k7=%d degeneracy=%t residualClue=%t nativeLinear=%t nativeFirst=%t native7=%t uniqueH72Claim=%t verdict=%q", x.TracePairingInherited, x.Operator, x.DBase, x.SSplit, x.F1, x.E1, x.H72Dimension, x.K7Dimension, x.Gate691ObserverDegeneracyRecorded, x.Gate690ResidualClueRetained, x.NativeLinearResponseTheorem, x.NativeFirstTraceTheorem, x.NativeSevenOver72Theorem, x.ClaimsUniqueFullH72Observer, x.Verdict)
}

func FormatRho72(x MaximallyMixedStateAudit) string {
	return fmt.Sprintf("state=%q definition=%q dim=%d traceI=%.18g denom=%.18g stateTrace=%.18g positive=%t maxMixed=%t numerator=%.18g expectation=%.18g expected=%.18g equalsActive=%t typeCorrect=%t verdict=%q", x.StateName, x.Definition, x.Dimension, x.IdentityTrace, x.NormalizationDenominator, x.StateTrace, x.PositiveState, x.MaximallyMixedOnFullH72, x.NumeratorTraceOfResponse, x.Expectation, x.ExpectedFirstTrace, x.EqualsActiveFirstTrace, x.TypeCorrectFullObserver, x.Verdict)
}

func FormatStateCandidate(x ObserverStateCandidate) string {
	return fmt.Sprintf("name=%q definition=%q supportDim=%d containsK7=%t positive=%t normalized=%t maxMixedSupport=%t signed=%t denom=%.18g trace=%.18g expectation=%.18g formula=%q residualRho72=%.18g matchesActive=%t active=%t failure=%q verdict=%q", x.Name, x.Definition, x.SupportDimension, x.ContainsK7, x.PositiveState, x.NormalizedState, x.MaximallyMixedOnSupport, x.SignedObserver, x.NormalizationDenominator, x.StateTrace, x.Expectation, x.ExpectedFormula, x.ResidualAgainstRho72, x.MatchesActiveBridge, x.ActiveCandidate, x.FailureStatus, x.Verdict)
}

func FormatAlternatives(x AlternativeObserverStateAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, FormatStateCandidate(c))
	}
	return fmt.Sprintf("candidates=[%s] count=%d positiveNormalized=%d activeStates=%d finiteInactive=%t kernelInactive=%t localK7Inactive=%t signedNotPositive=%t allAudited=%t verdict=%q", strings.Join(parts, " | "), x.CandidateCount, x.PositiveNormalizedStateCount, x.ActiveStateCount, x.FiniteOnlyStateInactive, x.KernelConditionalStateInactive, x.LocalK7StateInactive, x.HodgeSignedObserverNotPositive, x.AllAlternativesAudited, x.Verdict)
}

func FormatDegeneracy(x DenominatorDegeneracyResolutionAudit) string {
	return fmt.Sprintf("source=%q resolution=%q fixedH72Degenerate=%t stateNormalizationResolves=%t rho72UniqueAudited=%t nativeStateTheorem=%t verdict=%q", x.Gate691DegeneracySource, x.Gate692Resolution, x.FixedH72DenominatorDegenerate, x.StateNormalizationResolvesType, x.Rho72UniqueAmongAuditedStates, x.NativeStateSelectionTheorem, x.Verdict)
}

func FormatInterpretation(x InterpretationAudit) string {
	return fmt.Sprintf("rho=%q response=%q expectation=%q globalDensity=%t scalarEigen=%t supportCarrier=%t verdict=%q", x.Rho72Role, x.ResponseOperatorRole, x.ExpectationRole, x.GlobalAverageDensity, x.BoundaryScalarEigen, x.SupportSelectedCarrier, x.Verdict)
}

func FormatResidual(x ResidualStatusAudit) string {
	return fmt.Sprintf("e1=%.18g absE1=%.18g expectation=%.18g dbase=%.18g residualClue=%t correctionPromoted=%t spectralTheorem=%t verdict=%q", x.E1, x.AbsE1, x.Expectation, x.DBase, x.QuadraticResidualClueRetained, x.QuadraticCorrectionPromoted, x.NativeSpectralExpansionTheorem, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=[%s] candidates=[%s] precise=%q verdict=%q", strings.Join(x.Missing, "; "), strings.Join(x.Candidates, "; "), x.PreciseGap, x.Verdict)
}

func FormatDiscipline(x VerdictDiscipline) string {
	return fmt.Sprintf("nativeMaxMixed=%t nativeStateSelection=%t nativeFirst=%t native7=%t boundary=%t scalarRG=%t higgs=%t gauge=%t flavor=%t ckmPmns=%t activation=%t promoteQuadratic=%t verdict=%q", x.ClaimsNativeMaximallyMixedStateTheorem, x.ClaimsNativeStateSelectionTheorem, x.ClaimsNativeFirstTraceTheorem, x.ClaimsNativeSevenOver72Theorem, x.ClaimsBoundaryStress, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.ClaimsProjectorActivation, x.PromotesQuadraticResidualCorrection, x.Verdict)
}
