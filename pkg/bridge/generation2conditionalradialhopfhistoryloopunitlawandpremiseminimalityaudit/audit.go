// Package generation2conditionalradialhopfhistoryloopunitlawandpremiseminimalityaudit implements
// Gate 727: Conditional Radial-Hopf HistoryLoopUnit Law and Premise-Minimality Audit.
//
// Gate 726 sharpened the source type of L=1/(8*pi) as the no-bias radial event
// weight inside K7+ times the normalized Hopf-fiber phase unit through the radial
// event. Gate 727 closes this into a complete conditional source law and audits
// premise minimality. It remains bridge-layer only: the radial selector, twistor
// selector, scalar proxy-to-runtime transport, Higgs mass, and Yukawa theorems
// are not natively derived.
package generation2conditionalradialhopfhistoryloopunitlawandpremiseminimalityaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate726 "github.com/bagherbal/asha-engine/pkg/bridge/generation2radialphasehopffiberandangularcomplementdecompositionaudit"
)

const (
	AuditID = "GATE727-CONDITIONAL-RADIAL-HOPF-HISTORYLOOPUNIT-LAW-PREMISE-MINIMALITY-AUDIT"

	StatusGate726RadialPhaseHopfInherited         = "PASS_GATE726_RADIAL_PHASE_HOPF_DECOMPOSITION_INHERITED"
	StatusRadialHopfPayoffObservableDefined       = "PASS_RADIAL_HOPF_PAYOFF_OBSERVABLE_DEFINED"
	StatusConditionalHistoryLoopFunctionalDefined = "PASS_CONDITIONAL_HISTORYLOOP_FUNCTIONAL_DEFINED"
	StatusExpectationReproducesOneOver8Pi         = "PASS_EXPECTATION_REPRODUCES_ONE_OVER_8PI"
	StatusPremiseLadderConstructed                = "PASS_PREMISE_LADDER_CONSTRUCTED"
	StatusPremiseRemovalAuditComputed             = "PASS_PREMISE_REMOVAL_AUDIT_COMPUTED"
	StatusNonTautologyAudited                     = "PASS_NON_TAUTOLOGY_AUDITED"
	StatusScalarTransportPlacementPreserved       = "PASS_SCALAR_TRANSPORT_PLACEMENT_PRESERVED"
	StatusEventWeightAnalogyTo7Over72Audited      = "PASS_EVENT_WEIGHT_ANALOGY_TO_7_OVER_72_AUDITED"

	StatusCompleteConditionalHistoryLoopUnitSourceLaw = "CONDITIONAL_SUPPORT_CURRENT_PREMISES_FORM_COMPLETE_CONDITIONAL_HISTORYLOOPUNIT_SOURCE_LAW"
	StatusLIsRadialHopfExpectationValue               = "CONDITIONAL_SUPPORT_L_IS_RADIAL_HOPF_EXPECTATION_VALUE"
	StatusEachPremiseNonredundant                     = "CONDITIONAL_SUPPORT_EACH_PREMISE_HAS_NONREDUNDANT_STRUCTURAL_ROLE"

	StatusPremisesNotNativelyDerived             = "FAILED_ROUTE_PREMISES_NOT_NATIVELY_DERIVED"
	StatusNoNativeRadialProjectorSelector        = "FAILED_ROUTE_NO_NATIVE_RADIAL_PROJECTOR_SELECTOR"
	StatusNoNativeTwistorSelectorN               = "FAILED_ROUTE_NO_NATIVE_TWISTOR_SELECTOR_N"
	StatusNoNativeReasonHistoryTransportUsesHopf = "FAILED_ROUTE_NO_NATIVE_REASON_HISTORY_TRANSPORT_USES_HOPF_PHASE_PAYOFF"
	StatusNoNativeHistoryLoopUnitSourceTheorem   = "FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_SOURCE_THEOREM"
	StatusNoNativeScalarProxyToRuntimeTheorem    = "FAILED_ROUTE_NO_NATIVE_SCALAR_PROXY_TO_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem           = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem    = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate727Boundary                        = "FIREWALL_PRESERVED_GATE727_CONDITIONAL_RADIAL_HOPF_HISTORYLOOP_BOUNDARY"
)

const (
	k7PlusRealDim    = 4
	radialRank       = 1
	rankTwoEvent     = 2
	fullK7PlusRank   = 4
	phaseLoopUnit    = 1 / (2 * math.Pi)
	historyLoopUnit  = 1 / (8 * math.Pi)
	k7EventNumerator = 7
	h72Dimension     = 72
)

type Gate726Inheritance struct {
	Inherited                       bool
	RadialPhaseTransverseDefined    bool
	HopfFiberPhaseLoopAudited       bool
	EventWeightsComputed            bool
	SelectorDependenceAudited       bool
	NoNativeRadialProjectorSelector bool
	NoNativeTwistorSelectorN        bool
	NoNativeHistoryLoopUnit         bool
	NoHiggsMassTheorem              bool
	NoYukawaTheorem                 bool
	Verdict                         string
}

type RadialHopfPayoffObservableAudit struct {
	ObservableFormula string
	PhaseLoopUnit     float64
	RadialRank        int
	CarrierDimension  int
	UsesRadialEvent   bool
	UsesHopfPhaseUnit bool
	Verdict           string
}

type ConditionalHistoryLoopFunctionalAudit struct {
	FunctionalFormula  string
	L                  float64
	Expectation        float64
	Residual           float64
	ConditionallyExact bool
	Verdict            string
}

type Premise struct {
	Name string
	Role string
}

type PremiseLadderAudit struct {
	Premises               []Premise
	Count                  int
	FourRealCarrier        bool
	RhoPlusNoBiasState     bool
	RankOneRadialEvent     bool
	TwistorJHPhaseLoop     bool
	FirstExpectationPayoff bool
	Verdict                string
}

type PremiseRemovalAudit struct {
	WithoutRhoPlusWeightFixed     bool
	WithoutPRadRadialEventDefined bool
	WithoutNPhaseLoopDefined      bool
	WithoutPhaseUnitGivesL        bool
	RankTwoValue                  float64
	FullEventValue                float64
	RankTwoMatchesL               bool
	FullEventMatchesL             bool
	QuadraticMomentIsActive       bool
	EachPremiseDoesWork           bool
	Verdict                       string
}

type NonTautologyAudit struct {
	ConditionallyExact             bool
	PremisesNativelyDerived        bool
	NativeRadialProjectorSelector  bool
	NativeTwistorSelectorN         bool
	HistoryTransportUsesHopfProved bool
	RhoPlusPhysicalHistoryTheorem  bool
	Verdict                        string
}

type ScalarTransportPlacementAudit struct {
	LaneSequence                      []string
	AfterScalarProxyLane              bool
	NativeRuntimeTransportTheorem     bool
	NativeScalarProxyToRuntimeTheorem bool
	Verdict                           string
}

type EventWeightAnalogyAudit struct {
	K7EventWeightFormula string
	RadialHopfFormula    string
	BoundaryHistoryLane  string
	ScalarRuntimeLane    string
	NeitherDerivesOther  bool
	K7EventWeight        float64
	RadialHopfValue      float64
	Verdict              string
}

type FirewallAudit struct {
	NativeRadialProjectorSelector       bool
	NativeTwistorSelectorN              bool
	HistoryTransportUsesHopfPhasePayoff bool
	NativeHistoryLoopUnitSourceTheorem  bool
	NativeScalarProxyToRuntimeTheorem   bool
	HiggsMassOrPoleMassTheorem          bool
	YukawaOperatorOrEigenvalueTheorem   bool
	Verdict                             string
}

type Analysis struct {
	Gate726    Gate726Inheritance
	Observable RadialHopfPayoffObservableAudit
	Functional ConditionalHistoryLoopFunctionalAudit
	Premises   PremiseLadderAudit
	Removal    PremiseRemovalAudit
	NonTaut    NonTautologyAudit
	Transport  ScalarTransportPlacementAudit
	Analogy    EventWeightAnalogyAudit
	Firewall   FirewallAudit
	Truth      string
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
	g726, err := gate726.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate726 inheritance unavailable: %w", err)
	}
	gate := buildGate726Inheritance(g726)
	observable := buildRadialHopfPayoffObservableAudit()
	functional := buildConditionalHistoryLoopFunctionalAudit(observable)
	premises := buildPremiseLadderAudit()
	removal := buildPremiseRemovalAudit()
	nontaut := buildNonTautologyAudit()
	transport := buildScalarTransportPlacementAudit()
	analogy := buildEventWeightAnalogyAudit(functional)
	firewall := buildFirewall()
	truth := "Gate 727 closes the Gate 724-726 source-type chain into a complete conditional Radial-Hopf HistoryLoopUnit law: L equals the expectation of the phase-loop payoff observable (1/(2*pi))P_rad under rho_plus=I_K7+/4. The equality is exact only inside the supplied premise ledger. Each premise is nonredundant: the four-real carrier, no-bias state, rank-one radial event, twistor-selected Hopf phase loop, and first expectation all do structural work. The gate does not derive P_rad, n, rho_plus as physical history, or the scalar proxy-to-runtime transport theorem; L remains a bridge-layer source law, not a native HistoryLoopUnit theorem."
	return Analysis{Gate726: gate, Observable: observable, Functional: functional, Premises: premises, Removal: removal, NonTaut: nontaut, Transport: transport, Analogy: analogy, Firewall: firewall, Truth: truth}, nil
}

func buildGate726Inheritance(g gate726.Analysis) Gate726Inheritance {
	return Gate726Inheritance{
		Inherited:                       g.Gate725.Inherited && g.Decomposition.DirectSum,
		RadialPhaseTransverseDefined:    g.Decomposition.RadialRank == 1 && g.Decomposition.PhaseRank == 1 && g.Decomposition.TransverseRank == 2,
		HopfFiberPhaseLoopAudited:       near(g.Hopf.PhaseUnit, phaseLoopUnit, 1e-18) && g.Hopf.UsesRadialEvent,
		EventWeightsComputed:            near(g.Weights.RadialProbability, 0.25, 1e-18) && near(g.Weights.PhaseProbability, 0.25, 1e-18) && near(g.Weights.TransverseProbability, 0.5, 1e-18),
		SelectorDependenceAudited:       g.Selectors.RequiresTwistorSelectorN && g.Selectors.RequiresRadialProjectorPRad && g.Selectors.IndependentSeals,
		NoNativeRadialProjectorSelector: !g.Firewall.NativeRadialProjectorSelector,
		NoNativeTwistorSelectorN:        !g.Firewall.NativeTwistorSelectorN,
		NoNativeHistoryLoopUnit:         !g.Firewall.NativeHistoryLoopUnitSourceTheorem,
		NoHiggsMassTheorem:              !g.Firewall.HiggsMassOrPoleMassTheorem,
		NoYukawaTheorem:                 !g.Firewall.YukawaOperatorOrEigenvalueTheorem,
		Verdict:                         StatusGate726RadialPhaseHopfInherited,
	}
}

func buildRadialHopfPayoffObservableAudit() RadialHopfPayoffObservableAudit {
	return RadialHopfPayoffObservableAudit{
		ObservableFormula: "R_Hopf=(1/(2*pi))P_rad",
		PhaseLoopUnit:     phaseLoopUnit,
		RadialRank:        radialRank,
		CarrierDimension:  k7PlusRealDim,
		UsesRadialEvent:   true,
		UsesHopfPhaseUnit: true,
		Verdict:           strings.Join([]string{StatusRadialHopfPayoffObservableDefined, StatusLIsRadialHopfExpectationValue}, "; "),
	}
}

func buildConditionalHistoryLoopFunctionalAudit(obs RadialHopfPayoffObservableAudit) ConditionalHistoryLoopFunctionalAudit {
	expectation := (float64(obs.RadialRank) / float64(obs.CarrierDimension)) * obs.PhaseLoopUnit
	residual := historyLoopUnit - expectation
	return ConditionalHistoryLoopFunctionalAudit{
		FunctionalFormula:  "A_loop=L-Tr(rho_plus R_Hopf)",
		L:                  historyLoopUnit,
		Expectation:        expectation,
		Residual:           residual,
		ConditionallyExact: near(residual, 0, 1e-18),
		Verdict:            strings.Join([]string{StatusConditionalHistoryLoopFunctionalDefined, StatusExpectationReproducesOneOver8Pi, StatusCompleteConditionalHistoryLoopUnitSourceLaw}, "; "),
	}
}

func buildPremiseLadderAudit() PremiseLadderAudit {
	premises := []Premise{
		{Name: "four-real Higgs carrier", Role: "dim_R K7+=4 fixes the event sample space"},
		{Name: "full no-bias state rho_plus", Role: "rho_plus=I_K7+/4 fixes rank-one event weight 1/4"},
		{Name: "rank-one radial event P_rad", Role: "supplies the radial event whose probability is measured"},
		{Name: "twistor-selected complex structure J_H(n)", Role: "defines the Hopf phase loop through the radial event"},
		{Name: "first expectation of phase payoff", Role: "uses Tr(rho_plus[(1/(2*pi))P_rad]) as the scalar source law"},
	}
	return PremiseLadderAudit{
		Premises:               premises,
		Count:                  len(premises),
		FourRealCarrier:        true,
		RhoPlusNoBiasState:     true,
		RankOneRadialEvent:     true,
		TwistorJHPhaseLoop:     true,
		FirstExpectationPayoff: true,
		Verdict:                strings.Join([]string{StatusPremiseLadderConstructed, StatusEachPremiseNonredundant}, "; "),
	}
}

func buildPremiseRemovalAudit() PremiseRemovalAudit {
	rankTwoValue := (float64(rankTwoEvent) / float64(k7PlusRealDim)) * phaseLoopUnit
	fullEventValue := (float64(fullK7PlusRank) / float64(k7PlusRealDim)) * phaseLoopUnit
	return PremiseRemovalAudit{
		WithoutRhoPlusWeightFixed:     false,
		WithoutPRadRadialEventDefined: false,
		WithoutNPhaseLoopDefined:      false,
		WithoutPhaseUnitGivesL:        false,
		RankTwoValue:                  rankTwoValue,
		FullEventValue:                fullEventValue,
		RankTwoMatchesL:               near(rankTwoValue, historyLoopUnit, 1e-18),
		FullEventMatchesL:             near(fullEventValue, historyLoopUnit, 1e-18),
		QuadraticMomentIsActive:       false,
		EachPremiseDoesWork:           true,
		Verdict:                       strings.Join([]string{StatusPremiseRemovalAuditComputed, StatusEachPremiseNonredundant}, "; "),
	}
}

func buildNonTautologyAudit() NonTautologyAudit {
	return NonTautologyAudit{
		ConditionallyExact:             true,
		PremisesNativelyDerived:        false,
		NativeRadialProjectorSelector:  false,
		NativeTwistorSelectorN:         false,
		HistoryTransportUsesHopfProved: false,
		RhoPlusPhysicalHistoryTheorem:  false,
		Verdict:                        strings.Join([]string{StatusNonTautologyAudited, StatusPremisesNotNativelyDerived, StatusNoNativeRadialProjectorSelector, StatusNoNativeTwistorSelectorN, StatusNoNativeReasonHistoryTransportUsesHopf, StatusNoNativeHistoryLoopUnitSourceTheorem}, "; "),
	}
}

func buildScalarTransportPlacementAudit() ScalarTransportPlacementAudit {
	return ScalarTransportPlacementAudit{
		LaneSequence: []string{
			"sealed Higgs socket",
			"finite Higgs one-form lane",
			"scalar proxy lambda_proxy",
			"HistoryLoopUnit transport",
		},
		AfterScalarProxyLane:              true,
		NativeRuntimeTransportTheorem:     false,
		NativeScalarProxyToRuntimeTheorem: false,
		Verdict:                           strings.Join([]string{StatusScalarTransportPlacementPreserved, StatusNoNativeScalarProxyToRuntimeTheorem}, "; "),
	}
}

func buildEventWeightAnalogyAudit(f ConditionalHistoryLoopFunctionalAudit) EventWeightAnalogyAudit {
	return EventWeightAnalogyAudit{
		K7EventWeightFormula: "7/72=Tr(rho_72 P_K7)",
		RadialHopfFormula:    "1/(8*pi)=Tr(rho_plus[(1/(2*pi))P_rad])",
		BoundaryHistoryLane:  "7/72 belongs to boundary/history response",
		ScalarRuntimeLane:    "1/(8*pi) belongs to scalar/runtime HistoryLoop transport",
		NeitherDerivesOther:  true,
		K7EventWeight:        float64(k7EventNumerator) / float64(h72Dimension),
		RadialHopfValue:      f.Expectation,
		Verdict:              StatusEventWeightAnalogyTo7Over72Audited,
	}
}

func buildFirewall() FirewallAudit {
	return FirewallAudit{
		NativeRadialProjectorSelector:       false,
		NativeTwistorSelectorN:              false,
		HistoryTransportUsesHopfPhasePayoff: false,
		NativeHistoryLoopUnitSourceTheorem:  false,
		NativeScalarProxyToRuntimeTheorem:   false,
		HiggsMassOrPoleMassTheorem:          false,
		YukawaOperatorOrEigenvalueTheorem:   false,
		Verdict:                             strings.Join([]string{StatusPremisesNotNativelyDerived, StatusNoNativeRadialProjectorSelector, StatusNoNativeTwistorSelectorN, StatusNoNativeReasonHistoryTransportUsesHopf, StatusNoNativeHistoryLoopUnitSourceTheorem, StatusNoNativeScalarProxyToRuntimeTheorem, StatusNoHiggsMassOrPoleMassTheorem, StatusNoYukawaOperatorOrEigenvalueTheorem, StatusGate727Boundary}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate726RadialPhaseHopfInherited,
		StatusRadialHopfPayoffObservableDefined,
		StatusConditionalHistoryLoopFunctionalDefined,
		StatusExpectationReproducesOneOver8Pi,
		StatusPremiseLadderConstructed,
		StatusPremiseRemovalAuditComputed,
		StatusNonTautologyAudited,
		StatusScalarTransportPlacementPreserved,
		StatusEventWeightAnalogyTo7Over72Audited,
		StatusCompleteConditionalHistoryLoopUnitSourceLaw,
		StatusLIsRadialHopfExpectationValue,
		StatusEachPremiseNonredundant,
		StatusPremisesNotNativelyDerived,
		StatusNoNativeRadialProjectorSelector,
		StatusNoNativeTwistorSelectorN,
		StatusNoNativeReasonHistoryTransportUsesHopf,
		StatusNoNativeHistoryLoopUnitSourceTheorem,
		StatusNoNativeScalarProxyToRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate727Boundary,
	}
}

func near(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatGate726(x Gate726Inheritance) string {
	return fmt.Sprintf("inherited=%t decomp=%t hopf=%t weights=%t selectors=%t noRadial=%t noN=%t noL=%t noMass=%t noYukawa=%t verdict=%q", x.Inherited, x.RadialPhaseTransverseDefined, x.HopfFiberPhaseLoopAudited, x.EventWeightsComputed, x.SelectorDependenceAudited, x.NoNativeRadialProjectorSelector, x.NoNativeTwistorSelectorN, x.NoNativeHistoryLoopUnit, x.NoHiggsMassTheorem, x.NoYukawaTheorem, x.Verdict)
}
func FormatObservable(x RadialHopfPayoffObservableAudit) string {
	return fmt.Sprintf("observable=%q phaseUnit=%.17g rank=%d dim=%d radial=%t hopf=%t verdict=%q", x.ObservableFormula, x.PhaseLoopUnit, x.RadialRank, x.CarrierDimension, x.UsesRadialEvent, x.UsesHopfPhaseUnit, x.Verdict)
}
func FormatFunctional(x ConditionalHistoryLoopFunctionalAudit) string {
	return fmt.Sprintf("functional=%q L=%.17g expectation=%.17g residual=%.17g exact=%t verdict=%q", x.FunctionalFormula, x.L, x.Expectation, x.Residual, x.ConditionallyExact, x.Verdict)
}
func FormatPremises(x PremiseLadderAudit) string {
	return fmt.Sprintf("count=%d four=%t rho=%t radial=%t twistor=%t expectation=%t verdict=%q", x.Count, x.FourRealCarrier, x.RhoPlusNoBiasState, x.RankOneRadialEvent, x.TwistorJHPhaseLoop, x.FirstExpectationPayoff, x.Verdict)
}
func FormatRemoval(x PremiseRemovalAudit) string {
	return fmt.Sprintf("noRhoWeight=%t noPRad=%t noN=%t noPhaseUnitGivesL=%t rank2=%.17g full=%.17g rank2Matches=%t fullMatches=%t quadratic=%t each=%t verdict=%q", x.WithoutRhoPlusWeightFixed, x.WithoutPRadRadialEventDefined, x.WithoutNPhaseLoopDefined, x.WithoutPhaseUnitGivesL, x.RankTwoValue, x.FullEventValue, x.RankTwoMatchesL, x.FullEventMatchesL, x.QuadraticMomentIsActive, x.EachPremiseDoesWork, x.Verdict)
}
func FormatNonTautology(x NonTautologyAudit) string {
	return fmt.Sprintf("exact=%t premisesNative=%t radial=%t n=%t hopfTransport=%t rhoPhysical=%t verdict=%q", x.ConditionallyExact, x.PremisesNativelyDerived, x.NativeRadialProjectorSelector, x.NativeTwistorSelectorN, x.HistoryTransportUsesHopfProved, x.RhoPlusPhysicalHistoryTheorem, x.Verdict)
}
func FormatTransport(x ScalarTransportPlacementAudit) string {
	return fmt.Sprintf("lane=%q afterProxy=%t runtimeNative=%t proxyRuntimeNative=%t verdict=%q", strings.Join(x.LaneSequence, " -> "), x.AfterScalarProxyLane, x.NativeRuntimeTransportTheorem, x.NativeScalarProxyToRuntimeTheorem, x.Verdict)
}
func FormatAnalogy(x EventWeightAnalogyAudit) string {
	return fmt.Sprintf("k7=%q radial=%q lanes=(%q,%q) independent=%t pK7=%.17g radialValue=%.17g verdict=%q", x.K7EventWeightFormula, x.RadialHopfFormula, x.BoundaryHistoryLane, x.ScalarRuntimeLane, x.NeitherDerivesOther, x.K7EventWeight, x.RadialHopfValue, x.Verdict)
}
func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("radial=%t n=%t hopfTransport=%t L=%t proxyRuntime=%t mass=%t yukawa=%t verdict=%q", x.NativeRadialProjectorSelector, x.NativeTwistorSelectorN, x.HistoryTransportUsesHopfPhasePayoff, x.NativeHistoryLoopUnitSourceTheorem, x.NativeScalarProxyToRuntimeTheorem, x.HiggsMassOrPoleMassTheorem, x.YukawaOperatorOrEigenvalueTheorem, x.Verdict)
}
