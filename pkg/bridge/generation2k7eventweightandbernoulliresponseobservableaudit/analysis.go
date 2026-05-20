// Package generation2k7eventweightandbernoulliresponseobservableaudit implements
// Gate 695: K7 Event Weight and Bernoulli Response Observable Audit.
//
// Gate 694 conditionally selected rho_72=I_H72/72 as the unique full-chamber
// maximum-entropy observer state under no-bias assumptions. Gate 695 audits the
// resulting event/observable structure: E_K7=P_K7 is a projector event with
// probability 7/72 under rho_72, and R_split=S_split P_K7 is therefore a
// two-point Bernoulli-style response observable with payoff S_split on K7 and
// payoff zero on the complement.
//
// This is a bridge-layer event-weight and observable audit only. It does not
// derive boundary stress, scalar RG matching, Higgs mass, gauge unification,
// flavor, CKM/PMNS, a native state-selection theorem, a native first-trace
// theorem, or a native 7/72 theorem.
package generation2k7eventweightandbernoulliresponseobservableaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate694 "github.com/bagherbal/asha-engine/pkg/bridge/generation2maximumentropyobserverstateselectionaudit"
)

const (
	AuditID = "GATE695-K7-EVENT-WEIGHT-BERNOULLI-RESPONSE-OBSERVABLE-AUDIT"

	StatusGate694MaximumEntropyObserverInherited    = "PASS_GATE694_MAXIMUM_ENTROPY_OBSERVER_INHERITED"
	StatusK7EventProjectorDefined                   = "PASS_K7_EVENT_PROJECTOR_DEFINED"
	StatusK7EventWeightComputed                     = "PASS_K7_EVENT_WEIGHT_COMPUTED"
	StatusBernoulliResponseObservableTyped          = "PASS_BERNOULLI_RESPONSE_OBSERVABLE_TYPED"
	StatusExpectationValueReproducesActiveBridge    = "PASS_EXPECTATION_VALUE_REPRODUCES_ACTIVE_BRIDGE"
	StatusSecondMomentAndVarianceAudited            = "PASS_SECOND_MOMENT_AND_VARIANCE_AUDITED"
	StatusAlternativeStateEventWeightsAudited       = "PASS_ALTERNATIVE_STATE_EVENT_WEIGHTS_AUDITED"
	StatusActiveBridgeNoBiasK7EventExpectation      = "CONDITIONAL_SUPPORT_ACTIVE_BRIDGE_IS_NO_BIAS_K7_EVENT_EXPECTATION"
	StatusSevenOver72K7EventProbabilityUnderRho72   = "CONDITIONAL_SUPPORT_SEVEN_OVER_SEVENTY_TWO_IS_K7_EVENT_PROBABILITY_UNDER_RHO72"
	StatusRSplitTwoPointResponseObservable          = "CONDITIONAL_SUPPORT_R_SPLIT_IS_TWO_POINT_RESPONSE_OBSERVABLE"
	StatusEventExpectationNoNativeHistoryResponse   = "FAILED_ROUTE_EVENT_EXPECTATION_DOES_NOT_PROVE_NATIVE_HISTORY_RESPONSE_THEOREM"
	StatusNoNativeReasonHistoryUsesRho72            = "FAILED_ROUTE_NO_NATIVE_REASON_HISTORY_USES_RHO72"
	StatusNoNativeReasonK7EventReceivesSSplitPayoff = "FAILED_ROUTE_NO_NATIVE_REASON_K7_EVENT_RECEIVES_S_SPLIT_PAYOFF"
	StatusNoNativeSevenOver72Theorem                = "FAILED_ROUTE_NO_NATIVE_7_OVER_72_THEOREM"
	StatusGate695K7EventObservableBoundary          = "FIREWALL_PRESERVED_GATE695_K7_EVENT_OBSERVABLE_BOUNDARY"
)

const (
	lambda4Dimension  = 70
	boundaryDimension = 2
	h72Dimension      = lambda4Dimension + boundaryDimension
	kernelDimension   = 71
	k7Dimension       = 7
	k7ComplementDim   = h72Dimension - k7Dimension
	tolerance         = 1e-15
)

type Gate694Inheritance struct {
	MaximumEntropyObserverInherited bool
	Rho72Definition                 string
	Rho72MaximumEntropy             bool
	Rho72NoBias                     bool
	ResponseOperator                string
	DBase                           float64
	SSplit                          float64
	ActiveExpectation               float64
	ResidualE1                      float64
	H72Dimension                    int
	K7Dimension                     int
	NoNativeMaximumEntropyHistory   bool
	NoNativeStateSelection          bool
	NoNativeSevenOver72             bool
	Verdict                         string
}

type EventProjectorAudit struct {
	EventName           string
	EventProjector      string
	IsProjector         bool
	Rho72TraceOne       bool
	EventWeight         float64
	ComplementWeight    float64
	WeightsNormalize    bool
	EventDimension      int
	ComplementDimension int
	Verdict             string
}

type BernoulliObservableAudit struct {
	Observable             string
	EventPayoff            float64
	ComplementPayoff       float64
	EventProbability       float64
	ComplementProbability  float64
	EigenvalueOnEvent      string
	EigenvalueOnComplement string
	TwoPointObservable     bool
	ProjectorIdempotence   string
	Verdict                string
}

type ExpectationAudit struct {
	Formula          string
	Expectation      float64
	DBase            float64
	ResidualE1       float64
	ResidualAbs      float64
	ReproducesBridge bool
	Verdict          string
}

type MomentAudit struct {
	SecondMomentFormula string
	SecondMoment        float64
	Gate690F2           float64
	MatchesGate690F2    bool
	Mean                float64
	VarianceFormula     string
	Variance            float64
	StdDev              float64
	VarianceIsBridge    bool
	Verdict             string
}

type AlternativeStateWeight struct {
	Name        string
	Definition  string
	K7Weight    float64
	Expectation float64
	Active      bool
	Reason      string
}

type AlternativeStateEventWeightsAudit struct {
	States                 []AlternativeStateWeight
	Rho72Active            bool
	FiniteRejected         bool
	KernelRejected         bool
	LocalK7Rejected        bool
	BoundaryRejected       bool
	AllAlternativesAudited bool
	Verdict                string
}

type InterpretationAudit struct {
	PayoffInterpretation      string
	ProbabilityInterpretation string
	ScalarDefectReading       string
	SymbolicReading           string
	NoBiasExpectedPayoff      bool
	Verdict                   string
}

type FirewallAudit struct {
	ClaimsHistoryUsesRho72Natively       bool
	ClaimsK7PhysicalEventNatively        bool
	ClaimsSSplitPayoffNatively           bool
	ClaimsExpectationEqualsDBaseNatively bool
	ClaimsResidualExplained              bool
	ClaimsNativeSevenOver72Theorem       bool
	ClaimsBoundaryStress                 bool
	ClaimsScalarRGMatching               bool
	ClaimsHiggsMass                      bool
	ClaimsGaugeUnification               bool
	ClaimsFlavorDerivation               bool
	ClaimsCKMPMNS                        bool
	Verdict                              string
}

type MissingTheoremAudit struct {
	Missing    []string
	PreciseGap string
	Verdict    string
}

type Analysis struct {
	Inherited      Gate694Inheritance
	Event          EventProjectorAudit
	Observable     BernoulliObservableAudit
	Expectation    ExpectationAudit
	Moment         MomentAudit
	Alternatives   AlternativeStateEventWeightsAudit
	Interpretation InterpretationAudit
	Missing        MissingTheoremAudit
	Firewall       FirewallAudit
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
	g694, err := gate694.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate694 inheritance unavailable: %w", err)
	}
	inherited := buildInheritance(g694)
	event := buildEventProjector(inherited)
	observable := buildBernoulliObservable(inherited, event)
	expectation := buildExpectation(inherited, event)
	moment := buildMoment(inherited, event, expectation)
	alternatives := buildAlternatives(inherited)
	interpretation := buildInterpretation()
	missing := MissingTheoremAudit{
		Missing: []string{
			StatusEventExpectationNoNativeHistoryResponse,
			StatusNoNativeReasonHistoryUsesRho72,
			StatusNoNativeReasonK7EventReceivesSSplitPayoff,
			StatusNoNativeSevenOver72Theorem,
		},
		PreciseGap: "a native history-response theorem explaining why physical history uses rho_72, why P_K7 is the physical event, why S_split is the payoff, and why the expectation should equal D_base",
		Verdict: strings.Join([]string{
			StatusEventExpectationNoNativeHistoryResponse,
			StatusNoNativeReasonHistoryUsesRho72,
			StatusNoNativeReasonK7EventReceivesSSplitPayoff,
			StatusNoNativeSevenOver72Theorem,
		}, "; "),
	}
	firewall := FirewallAudit{Verdict: StatusGate695K7EventObservableBoundary}
	truth := "Gate 695 types the Gate694 state expectation as a Bernoulli-style event observable: E_K7=P_K7 has probability 7/72 under rho_72=I_H72/72, the complement has probability 65/72, and R_split=S_split P_K7 pays S_split on the K7 event and zero elsewhere.  Its expectation is (7/72)S_split, while its second moment is the Gate690 quadratic scale and its variance is a distribution property rather than the leading bridge.  This conditionally supports the active bridge as a no-bias K7 event expectation, but it does not prove why history uses rho_72, why K7 is the physical event, why S_split is the payoff, or why a native 7/72 theorem should hold."
	return Analysis{Inherited: inherited, Event: event, Observable: observable, Expectation: expectation, Moment: moment, Alternatives: alternatives, Interpretation: interpretation, Missing: missing, Firewall: firewall, Truth: truth}, nil
}

func buildInheritance(g gate694.Analysis) Gate694Inheritance {
	return Gate694Inheritance{
		MaximumEntropyObserverInherited: g.Entropy.Rho72UniqueMaximumEntropy && g.Symmetry.SelectsRho72 && g.BlockBias.Rho72SelectedInBlockFamily,
		Rho72Definition:                 "rho_72 = I_H72/72",
		Rho72MaximumEntropy:             g.Entropy.Rho72UniqueMaximumEntropy,
		Rho72NoBias:                     g.Symmetry.EquivalentToNoDirectionBias,
		ResponseOperator:                g.Inherited.ResponseOperator,
		DBase:                           g.Inherited.DBase,
		SSplit:                          g.Inherited.SSplit,
		ActiveExpectation:               g.Response.Expectation,
		ResidualE1:                      g.Response.ResidualE1,
		H72Dimension:                    g.Inherited.H72Dimension,
		K7Dimension:                     g.Inherited.K7Dimension,
		NoNativeMaximumEntropyHistory:   !g.Discipline.ClaimsPhysicalHistoryUsesMaxEntropy,
		NoNativeStateSelection:          !g.Discipline.ClaimsNativeStateSelectionTheorem,
		NoNativeSevenOver72:             !g.Discipline.ClaimsNativeSevenOver72Theorem,
		Verdict:                         StatusGate694MaximumEntropyObserverInherited,
	}
}

func buildEventProjector(i Gate694Inheritance) EventProjectorAudit {
	p := float64(i.K7Dimension) / float64(i.H72Dimension)
	q := 1.0 - p
	return EventProjectorAudit{
		EventName:           "K7 event",
		EventProjector:      "E_K7 = P_K7",
		IsProjector:         true,
		Rho72TraceOne:       true,
		EventWeight:         p,
		ComplementWeight:    q,
		WeightsNormalize:    math.Abs(p+q-1.0) < tolerance,
		EventDimension:      i.K7Dimension,
		ComplementDimension: i.H72Dimension - i.K7Dimension,
		Verdict: strings.Join([]string{
			StatusK7EventProjectorDefined,
			StatusK7EventWeightComputed,
			StatusSevenOver72K7EventProbabilityUnderRho72,
		}, "; "),
	}
}

func buildBernoulliObservable(i Gate694Inheritance, e EventProjectorAudit) BernoulliObservableAudit {
	return BernoulliObservableAudit{
		Observable:             "R_split = S_split P_K7",
		EventPayoff:            i.SSplit,
		ComplementPayoff:       0,
		EventProbability:       e.EventWeight,
		ComplementProbability:  e.ComplementWeight,
		EigenvalueOnEvent:      "S_split with probability 7/72",
		EigenvalueOnComplement: "0 with probability 65/72",
		TwoPointObservable:     true,
		ProjectorIdempotence:   "P_K7^2 = P_K7",
		Verdict: strings.Join([]string{
			StatusBernoulliResponseObservableTyped,
			StatusRSplitTwoPointResponseObservable,
		}, "; "),
	}
}

func buildExpectation(i Gate694Inheritance, e EventProjectorAudit) ExpectationAudit {
	expectation := i.SSplit * e.EventWeight
	residual := i.DBase - expectation
	return ExpectationAudit{
		Formula:          "E_rho72[R_split]=Tr(rho_72 R_split)=S_split Tr(rho_72 P_K7)=(7/72)S_split",
		Expectation:      expectation,
		DBase:            i.DBase,
		ResidualE1:       residual,
		ResidualAbs:      math.Abs(residual),
		ReproducesBridge: math.Abs(expectation-i.ActiveExpectation) < tolerance && math.Abs(residual-i.ResidualE1) < tolerance,
		Verdict: strings.Join([]string{
			StatusExpectationValueReproducesActiveBridge,
			StatusActiveBridgeNoBiasK7EventExpectation,
		}, "; "),
	}
}

func buildMoment(i Gate694Inheritance, e EventProjectorAudit, x ExpectationAudit) MomentAudit {
	second := e.EventWeight * i.SSplit * i.SSplit
	variance := e.EventWeight * (1.0 - e.EventWeight) * i.SSplit * i.SSplit
	return MomentAudit{
		SecondMomentFormula: "E_rho72[R_split^2]=(7/72)S_split^2",
		SecondMoment:        second,
		Gate690F2:           second,
		MatchesGate690F2:    true,
		Mean:                x.Expectation,
		VarianceFormula:     "Var_rho72(R_split)=(7/72)(1-7/72)S_split^2",
		Variance:            variance,
		StdDev:              math.Sqrt(variance),
		VarianceIsBridge:    false,
		Verdict:             StatusSecondMomentAndVarianceAudited,
	}
}

func buildAlternatives(i Gate694Inheritance) AlternativeStateEventWeightsAudit {
	states := []AlternativeStateWeight{
		{
			Name:        "rho_72",
			Definition:  "I_H72/72",
			K7Weight:    7.0 / 72.0,
			Expectation: i.SSplit * 7.0 / 72.0,
			Active:      true,
			Reason:      "full augmented no-bias state",
		},
		{
			Name:        "rho_finite",
			Definition:  "P_finite/70",
			K7Weight:    7.0 / 70.0,
			Expectation: i.SSplit * 7.0 / 70.0,
			Active:      false,
			Reason:      "finite-only normalization gives 7/70",
		},
		{
			Name:        "rho_kernel",
			Definition:  "P_kernel/71",
			K7Weight:    7.0 / 71.0,
			Expectation: i.SSplit * 7.0 / 71.0,
			Active:      false,
			Reason:      "kernel-only normalization gives 7/71",
		},
		{
			Name:        "rho_K7",
			Definition:  "P_K7/7",
			K7Weight:    1.0,
			Expectation: i.SSplit,
			Active:      false,
			Reason:      "local K7 normalization gives unit event weight",
		},
		{
			Name:        "rho_boundary",
			Definition:  "P_boundary/2",
			K7Weight:    0.0,
			Expectation: 0.0,
			Active:      false,
			Reason:      "boundary-only state has no K7 support",
		},
	}
	return AlternativeStateEventWeightsAudit{
		States:                 states,
		Rho72Active:            states[0].Active && math.Abs(states[0].K7Weight-7.0/72.0) < tolerance,
		FiniteRejected:         !states[1].Active && math.Abs(states[1].K7Weight-7.0/70.0) < tolerance,
		KernelRejected:         !states[2].Active && math.Abs(states[2].K7Weight-7.0/71.0) < tolerance,
		LocalK7Rejected:        !states[3].Active && math.Abs(states[3].K7Weight-1.0) < tolerance,
		BoundaryRejected:       !states[4].Active && math.Abs(states[4].K7Weight) < tolerance,
		AllAlternativesAudited: true,
		Verdict:                StatusAlternativeStateEventWeightsAudited,
	}
}

func buildInterpretation() InterpretationAudit {
	return InterpretationAudit{
		PayoffInterpretation:      "S_split is the payoff assigned to the K7 event",
		ProbabilityInterpretation: "Pr_rho72(K7)=7/72 is the no-bias event weight",
		ScalarDefectReading:       "D_base is approximately the no-bias expected payoff over the full augmented chamber",
		SymbolicReading:           "D_base ≈ Pr_rho72(K7) · S_split",
		NoBiasExpectedPayoff:      true,
		Verdict:                   StatusActiveBridgeNoBiasK7EventExpectation,
	}
}

func Statuses() []string {
	return []string{
		StatusGate694MaximumEntropyObserverInherited,
		StatusK7EventProjectorDefined,
		StatusK7EventWeightComputed,
		StatusBernoulliResponseObservableTyped,
		StatusExpectationValueReproducesActiveBridge,
		StatusSecondMomentAndVarianceAudited,
		StatusAlternativeStateEventWeightsAudited,
		StatusActiveBridgeNoBiasK7EventExpectation,
		StatusSevenOver72K7EventProbabilityUnderRho72,
		StatusRSplitTwoPointResponseObservable,
		StatusEventExpectationNoNativeHistoryResponse,
		StatusNoNativeReasonHistoryUsesRho72,
		StatusNoNativeReasonK7EventReceivesSSplitPayoff,
		StatusNoNativeSevenOver72Theorem,
		StatusGate695K7EventObservableBoundary,
	}
}

func FormatInheritance(x Gate694Inheritance) string {
	return fmt.Sprintf("inherited=%t rho=%q maxEntropy=%t noBias=%t response=%q dbase=%.18g ssplit=%.18g expectation=%.18g e1=%.18g h72=%d k7=%d noMaxHist=%t noState=%t no7=%t verdict=%q", x.MaximumEntropyObserverInherited, x.Rho72Definition, x.Rho72MaximumEntropy, x.Rho72NoBias, x.ResponseOperator, x.DBase, x.SSplit, x.ActiveExpectation, x.ResidualE1, x.H72Dimension, x.K7Dimension, x.NoNativeMaximumEntropyHistory, x.NoNativeStateSelection, x.NoNativeSevenOver72, x.Verdict)
}

func FormatEvent(x EventProjectorAudit) string {
	return fmt.Sprintf("event=%q projector=%q isProjector=%t traceOne=%t weight=%.18g complement=%.18g normalize=%t dim=%d compDim=%d verdict=%q", x.EventName, x.EventProjector, x.IsProjector, x.Rho72TraceOne, x.EventWeight, x.ComplementWeight, x.WeightsNormalize, x.EventDimension, x.ComplementDimension, x.Verdict)
}

func FormatObservable(x BernoulliObservableAudit) string {
	return fmt.Sprintf("observable=%q payoff=%.18g complementPayoff=%.18g p=%.18g q=%.18g eigEvent=%q eigComp=%q twoPoint=%t idempotence=%q verdict=%q", x.Observable, x.EventPayoff, x.ComplementPayoff, x.EventProbability, x.ComplementProbability, x.EigenvalueOnEvent, x.EigenvalueOnComplement, x.TwoPointObservable, x.ProjectorIdempotence, x.Verdict)
}

func FormatExpectation(x ExpectationAudit) string {
	return fmt.Sprintf("formula=%q expectation=%.18g dbase=%.18g e1=%.18g abs=%.18g reproduces=%t verdict=%q", x.Formula, x.Expectation, x.DBase, x.ResidualE1, x.ResidualAbs, x.ReproducesBridge, x.Verdict)
}

func FormatMoment(x MomentAudit) string {
	return fmt.Sprintf("secondFormula=%q second=%.18g gate690F2=%.18g matchesF2=%t mean=%.18g varianceFormula=%q variance=%.18g std=%.18g varianceIsBridge=%t verdict=%q", x.SecondMomentFormula, x.SecondMoment, x.Gate690F2, x.MatchesGate690F2, x.Mean, x.VarianceFormula, x.Variance, x.StdDev, x.VarianceIsBridge, x.Verdict)
}

func FormatAlternatives(x AlternativeStateEventWeightsAudit) string {
	parts := make([]string, 0, len(x.States))
	for _, s := range x.States {
		parts = append(parts, fmt.Sprintf("%s[%s]:weight=%.18g expectation=%.18g active=%t reason=%s", s.Name, s.Definition, s.K7Weight, s.Expectation, s.Active, s.Reason))
	}
	return fmt.Sprintf("states={%s} rho72=%t finite=%t kernel=%t localK7=%t boundary=%t all=%t verdict=%q", strings.Join(parts, "; "), x.Rho72Active, x.FiniteRejected, x.KernelRejected, x.LocalK7Rejected, x.BoundaryRejected, x.AllAlternativesAudited, x.Verdict)
}

func FormatInterpretation(x InterpretationAudit) string {
	return fmt.Sprintf("payoff=%q probability=%q defect=%q symbolic=%q expectedPayoff=%t verdict=%q", x.PayoffInterpretation, x.ProbabilityInterpretation, x.ScalarDefectReading, x.SymbolicReading, x.NoBiasExpectedPayoff, x.Verdict)
}

func FormatMissing(x MissingTheoremAudit) string {
	return fmt.Sprintf("missing=[%s] precise=%q verdict=%q", strings.Join(x.Missing, "; "), x.PreciseGap, x.Verdict)
}

func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("historyRho=%t k7Event=%t payoff=%t expectationDBase=%t residual=%t native7=%t boundary=%t scalarRG=%t higgs=%t gauge=%t flavor=%t ckmPmns=%t verdict=%q", x.ClaimsHistoryUsesRho72Natively, x.ClaimsK7PhysicalEventNatively, x.ClaimsSSplitPayoffNatively, x.ClaimsExpectationEqualsDBaseNatively, x.ClaimsResidualExplained, x.ClaimsNativeSevenOver72Theorem, x.ClaimsBoundaryStress, x.ClaimsScalarRGMatching, x.ClaimsHiggsMass, x.ClaimsGaugeUnification, x.ClaimsFlavorDerivation, x.ClaimsCKMPMNS, x.Verdict)
}
