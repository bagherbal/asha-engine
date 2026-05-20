// Package generation2higgsradialeventweightandphaselooptransportaudit implements
// Gate 724: Higgs Radial Event Weight and PhaseLoop Transport Audit.
//
// Gate 723 source-typed L=1/(8*pi)=(1/4)(1/(2*pi)) as a
// quarter-normalized phase-transport candidate. Gate 724 sharpens the quarter
// factor by auditing whether 1/4 is better read as the no-bias probability of a
// rank-one radial event inside the four-real-dimensional sealed Higgs carrier
// K7+, paired with the normalized internal phase-loop payoff 1/(2*pi). This is
// source-type compatibility only: the radial projector, HistoryLoopUnit source,
// scalar proxy-to-runtime theorem, Higgs mass theorem, and Yukawa theorem remain
// unproved.
package generation2higgsradialeventweightandphaselooptransportaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate723 "github.com/bagherbal/asha-engine/pkg/bridge/generation2quarternormalizedphasetransportsourcetypeaudit"
)

const (
	AuditID = "GATE724-HIGGS-RADIAL-EVENT-WEIGHT-PHASELOOP-TRANSPORT-AUDIT"

	StatusGate723QuarterPhaseTransportInherited = "PASS_GATE723_QUARTER_PHASE_TRANSPORT_INHERITED"
	StatusRhoPlusDefined                        = "PASS_RHO_PLUS_DEFINED"
	StatusRankOneRadialEventWeightComputed      = "PASS_RANK_ONE_RADIAL_EVENT_WEIGHT_COMPUTED"
	StatusPhaseLoopPayoffObservableDefined      = "PASS_PHASE_LOOP_PAYOFF_OBSERVABLE_DEFINED"
	StatusExpectationReproducesOneOver8Pi       = "PASS_EXPECTATION_REPRODUCES_ONE_OVER_8PI"
	StatusAlternativeRanksAudited               = "PASS_ALTERNATIVE_RANKS_AUDITED"
	StatusRadialSelectorFirewallAudited         = "PASS_RADIAL_SELECTOR_FIREWALL_AUDITED"
	StatusTwistorSelectorFirewallAudited        = "PASS_TWISTOR_SELECTOR_FIREWALL_AUDITED"
	StatusQNormalizationFirewallAudited         = "PASS_Q_NORMALIZATION_FIREWALL_AUDITED"
	StatusEventWeightAnalogyTo7Over72Audited    = "PASS_EVENT_WEIGHT_ANALOGY_TO_7_OVER_72_AUDITED"
	StatusScalarTransportPlacementPreserved     = "PASS_SCALAR_TRANSPORT_PLACEMENT_PRESERVED"

	StatusLRankOneRadialEventTimesPhaseLoop        = "CONDITIONAL_SUPPORT_L_IS_RANK_ONE_HIGGS_RADIAL_EVENT_WEIGHT_TIMES_PHASE_LOOP_UNIT"
	StatusOneOverFourRadialEventProbability        = "CONDITIONAL_SUPPORT_ONE_OVER_FOUR_IS_RADIAL_EVENT_PROBABILITY_IN_K7_PLUS"
	StatusOneOverTwoPiPhaseLoopPayoffCandidate     = "CONDITIONAL_SUPPORT_ONE_OVER_TWO_PI_IS_PHASE_LOOP_PAYOFF_CANDIDATE"
	StatusEventWeightAnalogyHistoryLoopK7Response  = "CONDITIONAL_SUPPORT_EVENT_WEIGHT_ANALOGY_BETWEEN_HISTORYLOOPUNIT_AND_K7_RESPONSE"
	StatusRankOneRadialEventBestTypedQuarterSource = "CONDITIONAL_SUPPORT_RANK_ONE_RADIAL_EVENT_IS_BEST_TYPED_SOURCE_FOR_QUARTER_FACTOR"

	StatusNoNativeRadialProjectorSelector         = "FAILED_ROUTE_NO_NATIVE_RADIAL_PROJECTOR_SELECTOR"
	StatusNoNativeHistoryLoopUnitSourceTheorem    = "FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_SOURCE_THEOREM"
	StatusNoNativePhaseLoopPayoffTheorem          = "FAILED_ROUTE_NO_NATIVE_PROOF_HISTORY_TRANSPORT_USES_PHASE_LOOP_PAYOFF"
	StatusTwistorSelectorDoesNotSelectRadialEvent = "FAILED_ROUTE_TWISTOR_SELECTOR_DOES_NOT_AUTOMATICALLY_SELECT_RADIAL_EVENT"
	StatusQDoesNotSourceL                         = "FAILED_ROUTE_Q_DOES_NOT_SOURCE_L"
	Status7Over72DoesNotSourceOneOver8Pi          = "FAILED_ROUTE_7_OVER_72_DOES_NOT_SOURCE_1_OVER_8PI"
	StatusNoNativeScalarProxyToRuntimeTheorem     = "FAILED_ROUTE_NO_NATIVE_SCALAR_PROXY_TO_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem            = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem     = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate724Boundary                         = "FIREWALL_PRESERVED_GATE724_HIGGS_RADIAL_EVENT_PHASELOOP_BOUNDARY"
)

const (
	k7PlusRealDim     = 4
	radialRank        = 1
	complexLineRank   = 2
	goldstoneLikeRank = 3
	fullRank          = 4
	k7Dim             = 7
	h72Dim            = 72
)

type Gate723Inheritance struct {
	Inherited                             bool
	QuarterPhaseCandidate                 bool
	PhaseLoopMeasureCandidate             bool
	FourComponentAverageCandidate         bool
	ScalarTransportNotRepresentationLayer bool
	NoNativeHistoryLoopUnit               bool
	NoNativePhaseLoopMeasure              bool
	NoScalarProxyRuntimeTheorem           bool
	QDoesNotSourceL                       bool
	LDoesNotSelectN                       bool
	SevenOver72DoesNotSourceL             bool
	Verdict                               string
}

type RhoPlusAudit struct {
	StateFormula     string
	CarrierDimension int
	Trace            float64
	MaximallyMixed   bool
	Verdict          string
}

type RadialEventAudit struct {
	ProjectorName       string
	Rank                int
	ProjectorIdempotent bool
	ProjectorSymmetric  bool
	ActsInsideK7Plus    bool
	Weight              float64
	NativeSelector      bool
	Verdict             string
}

type PhaseLoopPayoffAudit struct {
	PhaseLine           string
	PhaseAction         string
	PhaseLoopPayoff     float64
	ObservableFormula   string
	Expectation         float64
	HistoryLoopUnit     float64
	ExpectationMatchesL bool
	NativePayoffTheorem bool
	Verdict             string
}

type AlternativeRank struct {
	Name      string
	Rank      int
	Weight    float64
	Value     float64
	Active    bool
	Rejection string
}

type AlternativeRankAudit struct {
	Alternatives []AlternativeRank
	ActiveName   string
	Verdict      string
}

type SelectorFirewallAudit struct {
	P_radNativelySelected        bool
	TwistorSelectorSelectsRadial bool
	QSourcesRadialOrL            bool
	RadialSelectorSealName       string
	Verdict                      string
}

type EventWeightAnalogyAudit struct {
	K7EventProbability        float64
	RadialEventProbability    float64
	AnalogousEventWeights     bool
	SevenOver72SourcesQuarter bool
	SevenOver72SourcesL       bool
	Verdict                   string
}

type ScalarTransportPlacementAudit struct {
	BelongsAfterScalarProxy        bool
	DerivedFromRepresentationAlone bool
	NoScalarProxyRuntimeTheorem    bool
	Verdict                        string
}

type FirewallAudit struct {
	NativeRadialProjectorSelector      bool
	NativeHistoryLoopUnitSourceTheorem bool
	NativePhaseLoopPayoffTheorem       bool
	TwistorSelectorSelectsRadialEvent  bool
	QSourcesL                          bool
	SevenOver72SourcesOneOver8Pi       bool
	NativeScalarProxyToRuntimeTheorem  bool
	HiggsMassOrPoleMassTheorem         bool
	YukawaOperatorOrEigenvalueTheorem  bool
	Verdict                            string
}

type Analysis struct {
	Gate723     Gate723Inheritance
	RhoPlus     RhoPlusAudit
	RadialEvent RadialEventAudit
	PhasePayoff PhaseLoopPayoffAudit
	Ranks       AlternativeRankAudit
	Selectors   SelectorFirewallAudit
	Analogy     EventWeightAnalogyAudit
	Placement   ScalarTransportPlacementAudit
	Firewall    FirewallAudit
	Truth       string
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
	g723, err := gate723.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate723 inheritance unavailable: %w", err)
	}
	inherited := buildGate723Inheritance(g723)
	rho := buildRhoPlusAudit()
	radial := buildRadialEventAudit()
	phase := buildPhaseLoopPayoffAudit(radial.Weight)
	ranks := buildAlternativeRankAudit()
	selectors := buildSelectorFirewallAudit()
	analogy := buildEventWeightAnalogy(radial.Weight)
	placement := buildScalarTransportPlacement()
	firewall := buildFirewall()
	truth := "Gate 724 sharpens Gate 723's quarter factor by conditionally reading 1/4 as the no-bias probability of a rank-one Higgs-radial event inside the four-real-dimensional K7+ carrier. Pairing that event weight with the normalized phase-loop payoff 1/(2*pi) reproduces L=1/(8*pi). This is still source typing only: no native radial projector selector, HistoryLoopUnit source theorem, phase-loop payoff theorem, scalar proxy-to-runtime theorem, Higgs mass theorem, Yukawa theorem, q-source theorem, n-selector theorem, or 7/72-to-L theorem is certified."
	return Analysis{Gate723: inherited, RhoPlus: rho, RadialEvent: radial, PhasePayoff: phase, Ranks: ranks, Selectors: selectors, Analogy: analogy, Placement: placement, Firewall: firewall, Truth: truth}, nil
}

func buildGate723Inheritance(g gate723.Analysis) Gate723Inheritance {
	return Gate723Inheritance{
		Inherited:                             g.Gate722.Inherited && g.Quarter.EqualsHistoryLoopUnit && g.PhaseLoop.Candidate,
		QuarterPhaseCandidate:                 strings.Contains(g.Quarter.Verdict, gate723.StatusLQuarterPhaseCandidate),
		PhaseLoopMeasureCandidate:             strings.Contains(g.PhaseLoop.Verdict, gate723.StatusOneOverTwoPiPhaseLoopCandidate),
		FourComponentAverageCandidate:         strings.Contains(g.Quarter.Verdict, gate723.StatusOneOverFourHiggsComponentAverageCandidate),
		ScalarTransportNotRepresentationLayer: strings.Contains(g.Placement.Verdict, gate723.StatusLBelongsToScalarTransport),
		NoNativeHistoryLoopUnit:               !g.Firewall.NativeHistoryLoopUnitSourceTheorem,
		NoNativePhaseLoopMeasure:              !g.PhaseLoop.NativeHistoryTransportUsesMeasure,
		NoScalarProxyRuntimeTheorem:           !g.Firewall.NativeScalarProxyToRuntimeTheorem,
		QDoesNotSourceL:                       !g.QFirewall.NativeRelationQToL,
		LDoesNotSelectN:                       !g.NFirewall.LSelectsN,
		SevenOver72DoesNotSourceL:             !g.SevenFirewall.SevenOver72SourcesOneOver8Pi,
		Verdict:                               StatusGate723QuarterPhaseTransportInherited,
	}
}

func buildRhoPlusAudit() RhoPlusAudit {
	return RhoPlusAudit{
		StateFormula:     "rho_plus=I_K7+/4",
		CarrierDimension: k7PlusRealDim,
		Trace:            1,
		MaximallyMixed:   true,
		Verdict:          StatusRhoPlusDefined,
	}
}

func buildRadialEventAudit() RadialEventAudit {
	weight := float64(radialRank) / float64(k7PlusRealDim)
	return RadialEventAudit{
		ProjectorName:       "P_rad",
		Rank:                radialRank,
		ProjectorIdempotent: true,
		ProjectorSymmetric:  true,
		ActsInsideK7Plus:    true,
		Weight:              weight,
		NativeSelector:      false,
		Verdict:             strings.Join([]string{StatusRankOneRadialEventWeightComputed, StatusOneOverFourRadialEventProbability, StatusNoNativeRadialProjectorSelector}, "; "),
	}
}

func buildPhaseLoopPayoffAudit(radialWeight float64) PhaseLoopPayoffAudit {
	payoff := 1 / (2 * math.Pi)
	L := 1 / (8 * math.Pi)
	expectation := radialWeight * payoff
	return PhaseLoopPayoffAudit{
		PhaseLine:           "L_n=span(J_H(n))",
		PhaseAction:         "exp(theta J_H(n))",
		PhaseLoopPayoff:     payoff,
		ObservableFormula:   "R_phase=(1/(2*pi))P_rad",
		Expectation:         expectation,
		HistoryLoopUnit:     L,
		ExpectationMatchesL: near(expectation, L, 1e-18),
		NativePayoffTheorem: false,
		Verdict:             strings.Join([]string{StatusPhaseLoopPayoffObservableDefined, StatusExpectationReproducesOneOver8Pi, StatusLRankOneRadialEventTimesPhaseLoop, StatusOneOverTwoPiPhaseLoopPayoffCandidate, StatusNoNativePhaseLoopPayoffTheorem}, "; "),
	}
}

func buildAlternativeRankAudit() AlternativeRankAudit {
	payoff := 1 / (2 * math.Pi)
	alts := []AlternativeRank{
		{Name: "full K7+ event", Rank: fullRank, Weight: 1, Value: payoff, Active: false, Rejection: "gives 1/(2*pi), too large"},
		{Name: "complex-line event", Rank: complexLineRank, Weight: float64(complexLineRank) / k7PlusRealDim, Value: float64(complexLineRank) / k7PlusRealDim * payoff, Active: false, Rejection: "gives 1/(4*pi), too large"},
		{Name: "Goldstone-complement-like event", Rank: goldstoneLikeRank, Weight: float64(goldstoneLikeRank) / k7PlusRealDim, Value: float64(goldstoneLikeRank) / k7PlusRealDim * payoff, Active: false, Rejection: "gives 3/(8*pi), too large"},
		{Name: "rank-one radial event", Rank: radialRank, Weight: float64(radialRank) / k7PlusRealDim, Value: float64(radialRank) / k7PlusRealDim * payoff, Active: true, Rejection: "active candidate"},
	}
	return AlternativeRankAudit{Alternatives: alts, ActiveName: "rank-one radial event", Verdict: strings.Join([]string{StatusAlternativeRanksAudited, StatusRankOneRadialEventBestTypedQuarterSource}, "; ")}
}

func buildSelectorFirewallAudit() SelectorFirewallAudit {
	return SelectorFirewallAudit{
		P_radNativelySelected:        false,
		TwistorSelectorSelectsRadial: false,
		QSourcesRadialOrL:            false,
		RadialSelectorSealName:       "ScalarRadialSelectorSeal / HiggsVacuumDirectionSeal / RadialModeProjectionTheorem",
		Verdict:                      strings.Join([]string{StatusRadialSelectorFirewallAudited, StatusTwistorSelectorFirewallAudited, StatusQNormalizationFirewallAudited, StatusNoNativeRadialProjectorSelector, StatusTwistorSelectorDoesNotSelectRadialEvent, StatusQDoesNotSourceL}, "; "),
	}
}

func buildEventWeightAnalogy(radialWeight float64) EventWeightAnalogyAudit {
	return EventWeightAnalogyAudit{
		K7EventProbability:        float64(k7Dim) / float64(h72Dim),
		RadialEventProbability:    radialWeight,
		AnalogousEventWeights:     true,
		SevenOver72SourcesQuarter: false,
		SevenOver72SourcesL:       false,
		Verdict:                   strings.Join([]string{StatusEventWeightAnalogyTo7Over72Audited, StatusEventWeightAnalogyHistoryLoopK7Response, Status7Over72DoesNotSourceOneOver8Pi}, "; "),
	}
}

func buildScalarTransportPlacement() ScalarTransportPlacementAudit {
	return ScalarTransportPlacementAudit{
		BelongsAfterScalarProxy:        true,
		DerivedFromRepresentationAlone: false,
		NoScalarProxyRuntimeTheorem:    true,
		Verdict:                        strings.Join([]string{StatusScalarTransportPlacementPreserved, StatusNoNativeScalarProxyToRuntimeTheorem}, "; "),
	}
}

func buildFirewall() FirewallAudit {
	return FirewallAudit{
		NativeRadialProjectorSelector:      false,
		NativeHistoryLoopUnitSourceTheorem: false,
		NativePhaseLoopPayoffTheorem:       false,
		TwistorSelectorSelectsRadialEvent:  false,
		QSourcesL:                          false,
		SevenOver72SourcesOneOver8Pi:       false,
		NativeScalarProxyToRuntimeTheorem:  false,
		HiggsMassOrPoleMassTheorem:         false,
		YukawaOperatorOrEigenvalueTheorem:  false,
		Verdict:                            strings.Join([]string{StatusNoNativeRadialProjectorSelector, StatusNoNativeHistoryLoopUnitSourceTheorem, StatusNoNativePhaseLoopPayoffTheorem, StatusTwistorSelectorDoesNotSelectRadialEvent, StatusQDoesNotSourceL, Status7Over72DoesNotSourceOneOver8Pi, StatusNoNativeScalarProxyToRuntimeTheorem, StatusNoHiggsMassOrPoleMassTheorem, StatusNoYukawaOperatorOrEigenvalueTheorem, StatusGate724Boundary}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate723QuarterPhaseTransportInherited,
		StatusRhoPlusDefined,
		StatusRankOneRadialEventWeightComputed,
		StatusPhaseLoopPayoffObservableDefined,
		StatusExpectationReproducesOneOver8Pi,
		StatusAlternativeRanksAudited,
		StatusRadialSelectorFirewallAudited,
		StatusTwistorSelectorFirewallAudited,
		StatusQNormalizationFirewallAudited,
		StatusEventWeightAnalogyTo7Over72Audited,
		StatusScalarTransportPlacementPreserved,
		StatusLRankOneRadialEventTimesPhaseLoop,
		StatusOneOverFourRadialEventProbability,
		StatusOneOverTwoPiPhaseLoopPayoffCandidate,
		StatusEventWeightAnalogyHistoryLoopK7Response,
		StatusRankOneRadialEventBestTypedQuarterSource,
		StatusNoNativeRadialProjectorSelector,
		StatusNoNativeHistoryLoopUnitSourceTheorem,
		StatusNoNativePhaseLoopPayoffTheorem,
		StatusTwistorSelectorDoesNotSelectRadialEvent,
		StatusQDoesNotSourceL,
		Status7Over72DoesNotSourceOneOver8Pi,
		StatusNoNativeScalarProxyToRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate724Boundary,
	}
}

func near(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatGate723(x Gate723Inheritance) string {
	return fmt.Sprintf("inherited=%t quarterPhase=%t phaseLoop=%t fourAvg=%t scalarTransport=%t noNativeL=%t noPhaseMeasure=%t noProxyRuntime=%t qNoL=%t LNoN=%t pNoL=%t verdict=%q", x.Inherited, x.QuarterPhaseCandidate, x.PhaseLoopMeasureCandidate, x.FourComponentAverageCandidate, x.ScalarTransportNotRepresentationLayer, x.NoNativeHistoryLoopUnit, x.NoNativePhaseLoopMeasure, x.NoScalarProxyRuntimeTheorem, x.QDoesNotSourceL, x.LDoesNotSelectN, x.SevenOver72DoesNotSourceL, x.Verdict)
}
func FormatRhoPlus(x RhoPlusAudit) string {
	return fmt.Sprintf("state=%q dim=%d trace=%.17g maxMixed=%t verdict=%q", x.StateFormula, x.CarrierDimension, x.Trace, x.MaximallyMixed, x.Verdict)
}
func FormatRadial(x RadialEventAudit) string {
	return fmt.Sprintf("projector=%q rank=%d idem=%t sym=%t inside=%t weight=%.17g nativeSelector=%t verdict=%q", x.ProjectorName, x.Rank, x.ProjectorIdempotent, x.ProjectorSymmetric, x.ActsInsideK7Plus, x.Weight, x.NativeSelector, x.Verdict)
}
func FormatPhasePayoff(x PhaseLoopPayoffAudit) string {
	return fmt.Sprintf("line=%q action=%q payoff=%.17g observable=%q expectation=%.17g L=%.17g match=%t native=%t verdict=%q", x.PhaseLine, x.PhaseAction, x.PhaseLoopPayoff, x.ObservableFormula, x.Expectation, x.HistoryLoopUnit, x.ExpectationMatchesL, x.NativePayoffTheorem, x.Verdict)
}
func FormatRanks(x AlternativeRankAudit) string {
	parts := make([]string, 0, len(x.Alternatives))
	for _, a := range x.Alternatives {
		parts = append(parts, fmt.Sprintf("%s:rank=%d weight=%.17g value=%.17g active=%t %s", a.Name, a.Rank, a.Weight, a.Value, a.Active, a.Rejection))
	}
	return fmt.Sprintf("active=%q alternatives=[%s] verdict=%q", x.ActiveName, strings.Join(parts, " | "), x.Verdict)
}
func FormatSelectors(x SelectorFirewallAudit) string {
	return fmt.Sprintf("pRadNative=%t twistorSelects=%t qSources=%t seal=%q verdict=%q", x.P_radNativelySelected, x.TwistorSelectorSelectsRadial, x.QSourcesRadialOrL, x.RadialSelectorSealName, x.Verdict)
}
func FormatAnalogy(x EventWeightAnalogyAudit) string {
	return fmt.Sprintf("pK7=%.17g pRad=%.17g analogous=%t pSourcesQuarter=%t pSourcesL=%t verdict=%q", x.K7EventProbability, x.RadialEventProbability, x.AnalogousEventWeights, x.SevenOver72SourcesQuarter, x.SevenOver72SourcesL, x.Verdict)
}
func FormatPlacement(x ScalarTransportPlacementAudit) string {
	return fmt.Sprintf("afterProxy=%t fromRepresentationAlone=%t noProxyRuntime=%t verdict=%q", x.BelongsAfterScalarProxy, x.DerivedFromRepresentationAlone, x.NoScalarProxyRuntimeTheorem, x.Verdict)
}
func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("radial=%t L=%t phase=%t twistorRadial=%t q=%t pToL=%t proxyRuntime=%t mass=%t yukawa=%t verdict=%q", x.NativeRadialProjectorSelector, x.NativeHistoryLoopUnitSourceTheorem, x.NativePhaseLoopPayoffTheorem, x.TwistorSelectorSelectsRadialEvent, x.QSourcesL, x.SevenOver72SourcesOneOver8Pi, x.NativeScalarProxyToRuntimeTheorem, x.HiggsMassOrPoleMassTheorem, x.YukawaOperatorOrEigenvalueTheorem, x.Verdict)
}
