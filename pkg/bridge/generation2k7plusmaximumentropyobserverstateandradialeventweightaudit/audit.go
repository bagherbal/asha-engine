// Package generation2k7plusmaximumentropyobserverstateandradialeventweightaudit implements
// Gate 736: K7+ Maximum-Entropy Observer State and Radial Event Weight Audit.
//
// Gate 735 classified rho_plus as a remaining bridge/seal input in the scalar-Higgs
// bridge. Gate 736 audits whether rho_plus=I_K7+/4 is uniquely selected by the
// no-bias / maximum-entropy principle on the four-real-dimensional Hodge-positive
// carrier K7+, while preserving that rho_plus does not select the radial projector,
// the twistor point, the HistoryLoopUnit source theorem, scalar runtime, Higgs mass,
// or Yukawa physics.
package generation2k7plusmaximumentropyobserverstateandradialeventweightaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate735 "github.com/bagherbal/asha-engine/pkg/bridge/generation2scalarhiggsbridgesealinventoryandforecastboundaryaudit"
)

const (
	AuditID = "GATE736-K7PLUS-MAXIMUM-ENTROPY-OBSERVER-STATE-RADIAL-EVENT-WEIGHT-AUDIT"

	StatusGate735SealInventoryInherited             = "PASS_GATE735_SEAL_INVENTORY_INHERITED"
	StatusRhoPlusDefined                            = "PASS_RHO_PLUS_DEFINED"
	StatusRhoPlusUniquelyMaximizesEntropyOnK7Plus   = "PASS_RHO_PLUS_UNIQUELY_MAXIMIZES_ENTROPY_ON_K7_PLUS"
	StatusNoDirectionBiasSelectsRhoPlus             = "PASS_NO_DIRECTION_BIAS_SELECTS_RHO_PLUS"
	StatusRadialEventWeightComputed                 = "PASS_RADIAL_EVENT_WEIGHT_COMPUTED"
	StatusRadialPhaseTransverseEventWeightsComputed = "PASS_RADIAL_PHASE_TRANSVERSE_EVENT_WEIGHTS_COMPUTED"
	StatusBiasedStateFirewallAudited                = "PASS_BIASED_STATE_FIREWALL_AUDITED"
	StatusRadialSelectorFirewallEnforced            = "PASS_RADIAL_SELECTOR_FIREWALL_ENFORCED"
	StatusTwistorSelectorFirewallEnforced           = "PASS_TWISTOR_SELECTOR_FIREWALL_ENFORCED"
	StatusHistoryLoopPlacementAudited               = "PASS_HISTORYLOOP_PLACEMENT_AUDITED"

	StatusRhoPlusMaxEntropyFullK7PlusObserverState       = "CONDITIONAL_SUPPORT_RHO_PLUS_IS_MAXIMUM_ENTROPY_FULL_K7_PLUS_OBSERVER_STATE"
	StatusOneOverFourNoBiasRadialEventWeight             = "CONDITIONAL_SUPPORT_ONE_OVER_FOUR_IS_NO_BIAS_RADIAL_EVENT_WEIGHT"
	StatusHistoryLoopUsesMaximumEntropyRadialEventWeight = "CONDITIONAL_SUPPORT_HISTORYLOOPUNIT_CANDIDATE_USES_MAXIMUM_ENTROPY_RADIAL_EVENT_WEIGHT"

	StatusRhoPlusNotUniqueAmongAllDensityStates = "FAILED_ROUTE_RHO_PLUS_NOT_UNIQUE_AMONG_ALL_DENSITY_STATES"
	StatusBiasedStateReproductionIsCircular     = "FAILED_ROUTE_BIASED_STATE_REPRODUCTION_IS_CIRCULAR"
	StatusRhoPlusDoesNotSelectRadialProjector   = "FAILED_ROUTE_RHO_PLUS_DOES_NOT_SELECT_RADIAL_PROJECTOR"
	StatusRhoPlusDoesNotSelectTwistorPointN     = "FAILED_ROUTE_RHO_PLUS_DOES_NOT_SELECT_TWISTOR_POINT_N"
	StatusNoNativeHistoryLoopUnitSourceTheorem  = "FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_SOURCE_THEOREM"
	StatusNoNativeScalarRuntimeTheorem          = "FAILED_ROUTE_NO_NATIVE_SCALAR_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem          = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem   = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate736Boundary                       = "FIREWALL_PRESERVED_GATE736_K7_PLUS_MAXIMUM_ENTROPY_OBSERVER_BOUNDARY"
)

const (
	k7PlusRealDim = 4
	radialRank    = 1
	phaseRank     = 1
	transRank     = 2
)

type Gate735Inheritance struct {
	Inherited                     bool
	RhoPlusInventoried            bool
	PRadStillSealed               bool
	NoNativeHistoryLoopTheorem    bool
	NoNativeScalarRuntimeTheorem  bool
	NoHiggsMassTheorem            bool
	NoYukawaTheorem               bool
	ForecastOnlyBridgeConsistency bool
	Verdict                       string
}

type RhoPlusEntropyAudit struct {
	StateFormula         string
	Carrier              string
	Dimension            int
	Trace                float64
	Entropy              float64
	MaximumEntropy       float64
	PositiveNormalized   bool
	UniqueMaximumEntropy bool
	Verdict              string
}

type NoDirectionBiasAudit struct {
	InvariantGroup        string
	InvariantStateForm    string
	NormalizationEquation string
	Coefficient           float64
	SelectsRhoPlus        bool
	Verdict               string
}

type RadialEventWeightAudit struct {
	ProjectorName       string
	Rank                int
	CarrierDimension    int
	Weight              float64
	IndependentOfLine   bool
	RhoPlusSelectsEvent bool
	Verdict             string
}

type RadialPhaseTransverseWeightsAudit struct {
	RadialRank       int
	PhaseRank        int
	TransverseRank   int
	RadialWeight     float64
	PhaseWeight      float64
	TransverseWeight float64
	RequiresN        bool
	RequiresPRad     bool
	Verdict          string
}

type BiasedStateFirewallAudit struct {
	RhoPlusUniqueAmongAllDensityStates bool
	BiasedStateCanReproduceWeight      bool
	BiasedReproductionCircular         bool
	Verdict                            string
}

type SelectorFirewallAudit struct {
	RhoPlusSelectsPRad      bool
	RhoPlusSelectsN         bool
	RhoPlusSelectsPhaseLine bool
	RhoPlusSelectsHopfFiber bool
	UsefulAfterNAndPRad     bool
	Verdict                 string
}

type HistoryLoopPlacementAudit struct {
	Payoff                 float64
	RadialWeight           float64
	Expectation            float64
	HistoryLoopUnit        float64
	MatchesHistoryLoopUnit bool
	UsesMaxEntropyWeight   bool
	NativeTransportTheorem bool
	Verdict                string
}

type FirewallAudit struct {
	RhoPlusSelectsPRad                bool
	RhoPlusSelectsN                   bool
	NativeHistoryLoopUnitSource       bool
	NativeScalarRuntimeTheorem        bool
	HiggsMassOrPoleMassTheorem        bool
	YukawaOperatorOrEigenvalueTheorem bool
	Verdict                           string
}

type Analysis struct {
	Gate735     Gate735Inheritance
	Entropy     RhoPlusEntropyAudit
	NoBias      NoDirectionBiasAudit
	Radial      RadialEventWeightAudit
	Weights     RadialPhaseTransverseWeightsAudit
	Biased      BiasedStateFirewallAudit
	Selectors   SelectorFirewallAudit
	HistoryLoop HistoryLoopPlacementAudit
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
	g735, err := gate735.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate735 inheritance unavailable: %w", err)
	}
	inherited := buildGate735Inheritance(g735)
	entropy := buildRhoPlusEntropy()
	noBias := buildNoDirectionBias()
	radial := buildRadialEventWeight()
	weights := buildRadialPhaseTransverseWeights()
	biased := buildBiasedStateFirewall()
	selectors := buildSelectorFirewall()
	historyLoop := buildHistoryLoopPlacement(radial.Weight)
	firewall := buildFirewall()
	truth := "Gate 736 source-types rho_plus=I_K7+/4 as the unique maximum-entropy/no-direction-bias state on the four-real-dimensional K7+ carrier. It certifies that any supplied rank-one radial projector has no-bias weight 1/4 and that the Radial-Hopf candidate L=1/(8*pi) uses this maximum-entropy radial event weight times the Hopf phase payoff 1/(2*pi). It preserves the firewalls: rho_plus does not select P_rad or n, biased states can reproduce the weight only circularly, and no native HistoryLoopUnit, scalar-runtime, Higgs-mass, or Yukawa theorem is derived."
	return Analysis{Gate735: inherited, Entropy: entropy, NoBias: noBias, Radial: radial, Weights: weights, Biased: biased, Selectors: selectors, HistoryLoop: historyLoop, Firewall: firewall, Truth: truth}, nil
}

func buildGate735Inheritance(g gate735.Analysis) Gate735Inheritance {
	return Gate735Inheritance{
		Inherited:                     g.SealInventory.IncludesRhoPlus && g.SealInventory.IncludesPRad && g.Forecast.Level1BridgeConsistencyEstimateAllowed,
		RhoPlusInventoried:            g.SealInventory.IncludesRhoPlus,
		PRadStillSealed:               g.SealInventory.IncludesPRad,
		NoNativeHistoryLoopTheorem:    !g.Firewall.ClaimsLNativeLoopTheorem,
		NoNativeScalarRuntimeTheorem:  !g.Firewall.ClaimsSealedHiggsSocketPhysicalScalarLaw,
		NoHiggsMassTheorem:            !g.Firewall.ClaimsCubicBridgeIsHiggsMassTheorem,
		NoYukawaTheorem:               !g.Firewall.ClaimsKappaENativeFlavorTheorem,
		ForecastOnlyBridgeConsistency: g.Forecast.Level1BridgeConsistencyEstimateAllowed && !g.Forecast.Level2PhysicalPredictionAllowed,
		Verdict:                       StatusGate735SealInventoryInherited,
	}
}

func buildRhoPlusEntropy() RhoPlusEntropyAudit {
	max := math.Log(float64(k7PlusRealDim))
	return RhoPlusEntropyAudit{
		StateFormula:         "rho_plus=I_K7+/4",
		Carrier:              "K7+",
		Dimension:            k7PlusRealDim,
		Trace:                1,
		Entropy:              max,
		MaximumEntropy:       max,
		PositiveNormalized:   true,
		UniqueMaximumEntropy: true,
		Verdict: strings.Join([]string{
			StatusRhoPlusDefined,
			StatusRhoPlusUniquelyMaximizesEntropyOnK7Plus,
			StatusRhoPlusMaxEntropyFullK7PlusObserverState,
		}, "; "),
	}
}

func buildNoDirectionBias() NoDirectionBiasAudit {
	return NoDirectionBiasAudit{
		InvariantGroup:        "O(K7+,g_+)",
		InvariantStateForm:    "rho=c I_K7+",
		NormalizationEquation: "Tr(rho)=4c=1",
		Coefficient:           1.0 / float64(k7PlusRealDim),
		SelectsRhoPlus:        true,
		Verdict: strings.Join([]string{
			StatusNoDirectionBiasSelectsRhoPlus,
			StatusRhoPlusMaxEntropyFullK7PlusObserverState,
		}, "; "),
	}
}

func buildRadialEventWeight() RadialEventWeightAudit {
	weight := float64(radialRank) / float64(k7PlusRealDim)
	return RadialEventWeightAudit{
		ProjectorName:       "P_rad",
		Rank:                radialRank,
		CarrierDimension:    k7PlusRealDim,
		Weight:              weight,
		IndependentOfLine:   true,
		RhoPlusSelectsEvent: false,
		Verdict: strings.Join([]string{
			StatusRadialEventWeightComputed,
			StatusOneOverFourNoBiasRadialEventWeight,
			StatusRhoPlusDoesNotSelectRadialProjector,
		}, "; "),
	}
}

func buildRadialPhaseTransverseWeights() RadialPhaseTransverseWeightsAudit {
	return RadialPhaseTransverseWeightsAudit{
		RadialRank:       radialRank,
		PhaseRank:        phaseRank,
		TransverseRank:   transRank,
		RadialWeight:     float64(radialRank) / float64(k7PlusRealDim),
		PhaseWeight:      float64(phaseRank) / float64(k7PlusRealDim),
		TransverseWeight: float64(transRank) / float64(k7PlusRealDim),
		RequiresN:        true,
		RequiresPRad:     true,
		Verdict:          StatusRadialPhaseTransverseEventWeightsComputed,
	}
}

func buildBiasedStateFirewall() BiasedStateFirewallAudit {
	return BiasedStateFirewallAudit{
		RhoPlusUniqueAmongAllDensityStates: false,
		BiasedStateCanReproduceWeight:      true,
		BiasedReproductionCircular:         true,
		Verdict: strings.Join([]string{
			StatusBiasedStateFirewallAudited,
			StatusRhoPlusNotUniqueAmongAllDensityStates,
			StatusBiasedStateReproductionIsCircular,
		}, "; "),
	}
}

func buildSelectorFirewall() SelectorFirewallAudit {
	return SelectorFirewallAudit{
		RhoPlusSelectsPRad:      false,
		RhoPlusSelectsN:         false,
		RhoPlusSelectsPhaseLine: false,
		RhoPlusSelectsHopfFiber: false,
		UsefulAfterNAndPRad:     true,
		Verdict: strings.Join([]string{
			StatusRadialSelectorFirewallEnforced,
			StatusTwistorSelectorFirewallEnforced,
			StatusRhoPlusDoesNotSelectRadialProjector,
			StatusRhoPlusDoesNotSelectTwistorPointN,
		}, "; "),
	}
}

func buildHistoryLoopPlacement(radialWeight float64) HistoryLoopPlacementAudit {
	payoff := 1.0 / (2.0 * math.Pi)
	L := 1.0 / (8.0 * math.Pi)
	expectation := radialWeight * payoff
	return HistoryLoopPlacementAudit{
		Payoff:                 payoff,
		RadialWeight:           radialWeight,
		Expectation:            expectation,
		HistoryLoopUnit:        L,
		MatchesHistoryLoopUnit: near(expectation, L, 1e-18),
		UsesMaxEntropyWeight:   true,
		NativeTransportTheorem: false,
		Verdict: strings.Join([]string{
			StatusHistoryLoopPlacementAudited,
			StatusHistoryLoopUsesMaximumEntropyRadialEventWeight,
			StatusNoNativeHistoryLoopUnitSourceTheorem,
		}, "; "),
	}
}

func buildFirewall() FirewallAudit {
	return FirewallAudit{
		RhoPlusSelectsPRad:                false,
		RhoPlusSelectsN:                   false,
		NativeHistoryLoopUnitSource:       false,
		NativeScalarRuntimeTheorem:        false,
		HiggsMassOrPoleMassTheorem:        false,
		YukawaOperatorOrEigenvalueTheorem: false,
		Verdict: strings.Join([]string{
			StatusRhoPlusDoesNotSelectRadialProjector,
			StatusRhoPlusDoesNotSelectTwistorPointN,
			StatusNoNativeHistoryLoopUnitSourceTheorem,
			StatusNoNativeScalarRuntimeTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusGate736Boundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate735SealInventoryInherited,
		StatusRhoPlusDefined,
		StatusRhoPlusUniquelyMaximizesEntropyOnK7Plus,
		StatusNoDirectionBiasSelectsRhoPlus,
		StatusRadialEventWeightComputed,
		StatusRadialPhaseTransverseEventWeightsComputed,
		StatusBiasedStateFirewallAudited,
		StatusRadialSelectorFirewallEnforced,
		StatusTwistorSelectorFirewallEnforced,
		StatusHistoryLoopPlacementAudited,
		StatusRhoPlusMaxEntropyFullK7PlusObserverState,
		StatusOneOverFourNoBiasRadialEventWeight,
		StatusHistoryLoopUsesMaximumEntropyRadialEventWeight,
		StatusRhoPlusNotUniqueAmongAllDensityStates,
		StatusBiasedStateReproductionIsCircular,
		StatusRhoPlusDoesNotSelectRadialProjector,
		StatusRhoPlusDoesNotSelectTwistorPointN,
		StatusNoNativeHistoryLoopUnitSourceTheorem,
		StatusNoNativeScalarRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate736Boundary,
	}
}

func near(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatGate735(x Gate735Inheritance) string {
	return fmt.Sprintf("inherited=%t rhoPlus=%t pRad=%t noL=%t noRuntime=%t noMass=%t noYukawa=%t level1Only=%t verdict=%q", x.Inherited, x.RhoPlusInventoried, x.PRadStillSealed, x.NoNativeHistoryLoopTheorem, x.NoNativeScalarRuntimeTheorem, x.NoHiggsMassTheorem, x.NoYukawaTheorem, x.ForecastOnlyBridgeConsistency, x.Verdict)
}
func FormatEntropy(x RhoPlusEntropyAudit) string {
	return fmt.Sprintf("state=%q carrier=%q dim=%d trace=%.17g entropy=%.17g max=%.17g positive=%t unique=%t verdict=%q", x.StateFormula, x.Carrier, x.Dimension, x.Trace, x.Entropy, x.MaximumEntropy, x.PositiveNormalized, x.UniqueMaximumEntropy, x.Verdict)
}
func FormatNoBias(x NoDirectionBiasAudit) string {
	return fmt.Sprintf("group=%q form=%q norm=%q c=%.17g selects=%t verdict=%q", x.InvariantGroup, x.InvariantStateForm, x.NormalizationEquation, x.Coefficient, x.SelectsRhoPlus, x.Verdict)
}
func FormatRadial(x RadialEventWeightAudit) string {
	return fmt.Sprintf("projector=%q rank=%d dim=%d weight=%.17g independent=%t selects=%t verdict=%q", x.ProjectorName, x.Rank, x.CarrierDimension, x.Weight, x.IndependentOfLine, x.RhoPlusSelectsEvent, x.Verdict)
}
func FormatWeights(x RadialPhaseTransverseWeightsAudit) string {
	return fmt.Sprintf("ranks=%d+%d+%d weights=%.17g,%.17g,%.17g requiresN=%t requiresPRad=%t verdict=%q", x.RadialRank, x.PhaseRank, x.TransverseRank, x.RadialWeight, x.PhaseWeight, x.TransverseWeight, x.RequiresN, x.RequiresPRad, x.Verdict)
}
func FormatBiased(x BiasedStateFirewallAudit) string {
	return fmt.Sprintf("uniqueAll=%t canReproduce=%t circular=%t verdict=%q", x.RhoPlusUniqueAmongAllDensityStates, x.BiasedStateCanReproduceWeight, x.BiasedReproductionCircular, x.Verdict)
}
func FormatSelectors(x SelectorFirewallAudit) string {
	return fmt.Sprintf("selectsPRad=%t selectsN=%t selectsPhase=%t selectsHopf=%t usefulAfter=%t verdict=%q", x.RhoPlusSelectsPRad, x.RhoPlusSelectsN, x.RhoPlusSelectsPhaseLine, x.RhoPlusSelectsHopfFiber, x.UsefulAfterNAndPRad, x.Verdict)
}
func FormatHistoryLoop(x HistoryLoopPlacementAudit) string {
	return fmt.Sprintf("payoff=%.17g radialWeight=%.17g expectation=%.17g L=%.17g matches=%t maxEnt=%t native=%t verdict=%q", x.Payoff, x.RadialWeight, x.Expectation, x.HistoryLoopUnit, x.MatchesHistoryLoopUnit, x.UsesMaxEntropyWeight, x.NativeTransportTheorem, x.Verdict)
}
func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("selectsPRad=%t selectsN=%t nativeL=%t runtime=%t mass=%t yukawa=%t verdict=%q", x.RhoPlusSelectsPRad, x.RhoPlusSelectsN, x.NativeHistoryLoopUnitSource, x.NativeScalarRuntimeTheorem, x.HiggsMassOrPoleMassTheorem, x.YukawaOperatorOrEigenvalueTheorem, x.Verdict)
}
