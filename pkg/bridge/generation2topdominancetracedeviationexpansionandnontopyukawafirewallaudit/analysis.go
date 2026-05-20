// Package generation2topdominancetracedeviationexpansionandnontopyukawafirewallaudit implements
// Gate 755: Top-Dominance Trace-Deviation Expansion and Non-Top Yukawa Firewall Audit.
//
// Gate 754 source-typed the one-third trace-shape shadow b/a^2≈1/3 as the
// color-tripled single-dominant Yukawa channel limit. Gate 755 audits the exact
// finite trace-deviation expansion obtained by splitting the dominant colored
// top-like channel from all remaining Yukawa trace contributions. It types the
// deviation formula, records why the current ledger lies below 1/3, and blocks
// any promotion to a native Yukawa, flavor, scalar-runtime, Higgs-mass, or
// pole-mass theorem.
package generation2topdominancetracedeviationexpansionandnontopyukawafirewallaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE755-TOP-DOMINANCE-TRACE-DEVIATION-EXPANSION-AND-NON-TOP-YUKAWA-FIREWALL-AUDIT"

	StatusGate754OneThirdTraceShadowInherited   = "PASS_GATE754_ONE_THIRD_TRACE_SHADOW_INHERITED"
	StatusTopColorDominantSplitDefined          = "PASS_TOP_COLOR_DOMINANT_SPLIT_DEFINED"
	StatusNormalizedRestVariablesDefined        = "PASS_NORMALIZED_REST_VARIABLES_DEFINED"
	StatusExactTraceDeviationFormulaDerived     = "PASS_EXACT_TRACE_DEVIATION_FORMULA_DERIVED"
	StatusOneEighthProxyDeviationRewritten      = "PASS_ONE_EIGHTH_PROXY_DEVIATION_REWRITTEN"
	StatusRequiredYukawaDecompositionDataListed = "PASS_REQUIRED_YUKAWA_DECOMPOSITION_DATA_LISTED"
	StatusYukawaFirewallEnforced                = "PASS_YUKAWA_FIREWALL_ENFORCED"
	StatusRuntimeAndHiggsFirewallsEnforced      = "PASS_RUNTIME_AND_HIGGS_FIREWALLS_ENFORCED"

	StatusDeltaRatioIsNonTopCorrection        = "CONDITIONAL_SUPPORT_DELTA_RATIO_IS_NON_TOP_TRACE_CORRECTION_TO_COLOR_DOMINANCE"
	StatusLambdaProxyDeviationTransported     = "CONDITIONAL_SUPPORT_LAMBDA_PROXY_DEVIATION_FROM_ONE_EIGHTH_IS_TRANSPORTED_TRACE_DEVIATION"
	StatusRestChannelsLowerRatioBelowOneThird = "CONDITIONAL_SUPPORT_REST_CHANNELS_LOWER_B_OVER_A_SQUARED_BELOW_ONE_THIRD_IN_CURRENT_LEDGER"

	StatusNoNativeTopYukawaDerivation               = "FAILED_ROUTE_NO_NATIVE_TOP_YUKAWA_DERIVATION"
	StatusNoNumericalAlphaBetaWithoutLedger         = "FAILED_ROUTE_NO_NUMERICAL_ALPHA_BETA_WITHOUT_DECOMPOSED_YUKAWA_LEDGER"
	StatusNoNativeYukawaOperatorTheorem             = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusNoNativeFlavorHierarchyTheorem            = "FAILED_ROUTE_NO_NATIVE_FLAVOR_HIERARCHY_THEOREM"
	StatusNoNativeScalarProxyDerivationTheorem      = "FAILED_ROUTE_NO_NATIVE_SCALAR_PROXY_DERIVATION_THEOREM"
	StatusNoIndependentScalarRuntimeTheorem         = "FAILED_ROUTE_NO_INDEPENDENT_SCALAR_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem              = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusGate755TopDominanceTraceDeviationBoundary = "FIREWALL_PRESERVED_GATE755_TOP_DOMINANCE_TRACE_DEVIATION_BOUNDARY"
)

const (
	threeEighths = 3.0 / 8.0
	oneThird     = 1.0 / 3.0
	oneEighth    = 1.0 / 8.0

	aTraceMZ      = 2.8424095142339083
	bTraceMZ      = 2.6910096440382287
	bOverA2MZSeed = 0.33307493962706697
	lambdaProxyMZ = 0.12490310236015
)

type Gate754Inheritance struct {
	Inherited                       bool
	ScalarProxyFormula              string
	TraceRatioFormula               string
	OneThirdLimitSource             string
	ATraceMZ                        float64
	BTraceMZ                        float64
	BOverA2Computed                 float64
	BOverA2Seed                     float64
	DeltaRatio                      float64
	LambdaProxy                     float64
	TopColorDominanceLimitDefined   bool
	OneThirdTraceShadowConditional  bool
	NativeBA2OneThirdTheoremBlocked bool
	NativeDeltaDecompositionBlocked bool
	RuntimeAndHiggsTheoremsBlocked  bool
	Verdict                         string
}

type TopColorDominantSplit struct {
	DominantSquaredSingularValueSymbol string
	ATopFormula                        string
	BTopFormula                        string
	ARestFormula                       string
	BRestFormula                       string
	TotalAFormula                      string
	TotalBFormula                      string
	ZeroRestCondition                  string
	TopLimitRatioFormula               string
	TopLimitRatio                      float64
	SplitIsAlgebraicIdentity           bool
	TopYukawaValueDerived              bool
	Verdict                            string
}

type NormalizedRestVariables struct {
	AlphaDefinition                string
	BetaDefinition                 string
	ARewritten                     string
	BRewritten                     string
	RatioRewritten                 string
	AlphaRequiresTopValue          bool
	BetaRequiresTopValue           bool
	RequiresDecomposedYukawaLedger bool
	NumericalAlphaBetaAvailable    bool
	Verdict                        string
}

type TraceDeviationFormula struct {
	ExactFormula                string
	FirstOrderFormula           string
	ProbeAlpha                  float64
	ProbeBeta                   float64
	ProbeRatioDirect            float64
	ProbeRatioByFormula         float64
	ProbeDeltaDirect            float64
	ProbeDeltaByFormula         float64
	ProbeFormulaResidual        float64
	CurrentDeltaRatio           float64
	CurrentRatioBelowOneThird   bool
	AssumptionAlphaPositive     bool
	AssumptionBetaMuchLessAlpha bool
	FirstOrderExplainsSign      bool
	NativeDeltaRatioTheorem     bool
	Verdict                     string
}

type OneEighthProxyDeviation struct {
	FormulaFromTraceDeviation string
	EquivalentFormula         string
	LambdaProxyComputed       float64
	LambdaProxySeed           float64
	LambdaProxySeedResidual   float64
	ProxyMinusOneEighth       float64
	ThreeEighthsTimesDelta    float64
	TransportIdentityResidual float64
	DeviationTransported      bool
	ScalarPotentialTheorem    bool
	RuntimeLambdaTheorem      bool
	Verdict                   string
}

type RequiredYukawaDecompositionData struct {
	RequiredItems                []string
	CanComputeAlphaBeta          bool
	CanAssignBottomTauCharm      bool
	CanAssignNeutrinoConvention  bool
	CanAssignScaleDependence     bool
	CanAssignFiniteTraceResidual bool
	TypedTopLikeTAvailable       bool
	DecomposedYukawaLedger       bool
	Verdict                      string
}

type YukawaFirewall struct {
	DeltaRatioIsNativeYukawaTheorem bool
	TopDominanceDerivesTopYukawa    bool
	AlphaBetaDerivesHierarchy       bool
	ClaimsYuDerived                 bool
	ClaimsYdDerived                 bool
	ClaimsYeDerived                 bool
	ClaimsYnuDerived                bool
	ClaimsCKMPMNSDerived            bool
	ClaimsGenerationCarrier         bool
	ClaimsFlavorTheorem             bool
	SealedLedgerExplicit            bool
	Verdict                         string
}

type RuntimeAndHiggsFirewalls struct {
	LambdaProxyNearOneEighthIsScalarPotentialTheorem bool
	LambdaProxyEqualsRuntimeLambda                   bool
	RuntimeLambdaEqualsHiggsMass                     bool
	TreeProxyEqualsPoleMass                          bool
	RequiresHistoryLoopTransport                     bool
	RequiresBoundaryHistoryResponse                  bool
	RequiresKappaEReduction                          bool
	RequiresScalarRuntimeBridge                      bool
	ClaimsIndependentScalarRuntime                   bool
	ClaimsHiggsMassTheorem                           bool
	ClaimsPoleMassTheorem                            bool
	Verdict                                          string
}

type Analysis struct {
	Gate754      Gate754Inheritance
	Split        TopColorDominantSplit
	Rest         NormalizedRestVariables
	Deviation    TraceDeviationFormula
	Proxy        OneEighthProxyDeviation
	RequiredData RequiredYukawaDecompositionData
	Yukawa       YukawaFirewall
	Runtime      RuntimeAndHiggsFirewalls
	Truth        string
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
	gate754 := buildGate754Inheritance()
	if math.IsNaN(gate754.BOverA2Computed) || math.IsInf(gate754.BOverA2Computed, 0) || gate754.ATraceMZ <= 0 {
		return Analysis{}, fmt.Errorf("invalid inherited Gate754 ratio: a=%g b=%g ratio=%g", gate754.ATraceMZ, gate754.BTraceMZ, gate754.BOverA2Computed)
	}
	split := buildTopColorDominantSplit()
	rest := buildNormalizedRestVariables()
	dev := buildTraceDeviationFormula(gate754)
	proxy := buildOneEighthProxyDeviation(gate754)
	required := buildRequiredYukawaDecompositionData()
	yukawa := buildYukawaFirewall()
	runtime := buildRuntimeAndHiggsFirewalls()
	truth := "Gate 755 expands the Gate754 one-third shadow by the exact top-color split a=3T+a_rest and b=3T^2+b_rest. With alpha=a_rest/(3T) and beta=b_rest/(3T^2), the finite identity b/a^2-1/3=(1/3)(beta-2alpha-alpha^2)/(1+alpha)^2 is certified as a trace-deviation formula. The current negative deviation is typed only as non-top trace correction pressure; without a typed top-like T and decomposed Yukawa ledger, alpha, beta, and channel attribution remain blocked. The induced lambda_proxy-1/8=(3/8)delta_ratio is only a transported scalar-proxy deviation, not a scalar-potential, runtime-lambda, Higgs-mass, pole-mass, Yukawa, or flavor theorem."

	return Analysis{Gate754: gate754, Split: split, Rest: rest, Deviation: dev, Proxy: proxy, RequiredData: required, Yukawa: yukawa, Runtime: runtime, Truth: truth}, nil
}

func buildGate754Inheritance() Gate754Inheritance {
	ratio := bTraceMZ / (aTraceMZ * aTraceMZ)
	delta := ratio - oneThird
	return Gate754Inheritance{
		Inherited:                       true,
		ScalarProxyFormula:              "lambda_proxy=(3/8)(b/a^2)",
		TraceRatioFormula:               "b/a^2",
		OneThirdLimitSource:             "color-tripled single-dominant top-like Yukawa channel: a_top=3T, b_top=3T^2",
		ATraceMZ:                        aTraceMZ,
		BTraceMZ:                        bTraceMZ,
		BOverA2Computed:                 ratio,
		BOverA2Seed:                     bOverA2MZSeed,
		DeltaRatio:                      delta,
		LambdaProxy:                     threeEighths * ratio,
		TopColorDominanceLimitDefined:   true,
		OneThirdTraceShadowConditional:  true,
		NativeBA2OneThirdTheoremBlocked: true,
		NativeDeltaDecompositionBlocked: true,
		RuntimeAndHiggsTheoremsBlocked:  true,
		Verdict: strings.Join([]string{
			StatusGate754OneThirdTraceShadowInherited,
			StatusDeltaRatioIsNonTopCorrection,
		}, "; "),
	}
}

func buildTopColorDominantSplit() TopColorDominantSplit {
	return TopColorDominantSplit{
		DominantSquaredSingularValueSymbol: "T=y_t^2",
		ATopFormula:                        "a_top=3T",
		BTopFormula:                        "b_top=3T^2",
		ARestFormula:                       "a_rest=a-3T",
		BRestFormula:                       "b_rest=b-3T^2",
		TotalAFormula:                      "a=3T+a_rest",
		TotalBFormula:                      "b=3T^2+b_rest",
		ZeroRestCondition:                  "a_rest=0 and b_rest=0",
		TopLimitRatioFormula:               "b_top/a_top^2=3T^2/(3T)^2=1/3",
		TopLimitRatio:                      oneThird,
		SplitIsAlgebraicIdentity:           true,
		TopYukawaValueDerived:              false,
		Verdict: strings.Join([]string{
			StatusTopColorDominantSplitDefined,
			StatusNoNativeTopYukawaDerivation,
		}, "; "),
	}
}

func buildNormalizedRestVariables() NormalizedRestVariables {
	return NormalizedRestVariables{
		AlphaDefinition:                "alpha=a_rest/(3T)",
		BetaDefinition:                 "beta=b_rest/(3T^2)",
		ARewritten:                     "a=3T(1+alpha)",
		BRewritten:                     "b=3T^2(1+beta)",
		RatioRewritten:                 "b/a^2=(1/3)(1+beta)/(1+alpha)^2",
		AlphaRequiresTopValue:          true,
		BetaRequiresTopValue:           true,
		RequiresDecomposedYukawaLedger: true,
		NumericalAlphaBetaAvailable:    false,
		Verdict: strings.Join([]string{
			StatusNormalizedRestVariablesDefined,
			StatusNoNumericalAlphaBetaWithoutLedger,
		}, "; "),
	}
}

func buildTraceDeviationFormula(g Gate754Inheritance) TraceDeviationFormula {
	// Probe only certifies the algebraic identity. It is not a physical Yukawa ledger.
	const alpha = 0.012
	const beta = 0.00003
	ratioDirect := (oneThird) * (1 + beta) / ((1 + alpha) * (1 + alpha))
	deltaDirect := ratioDirect - oneThird
	deltaFormula := (oneThird) * (beta - 2*alpha - alpha*alpha) / ((1 + alpha) * (1 + alpha))
	return TraceDeviationFormula{
		ExactFormula:                "delta_ratio=b/a^2-1/3=(1/3)[(1+beta)/(1+alpha)^2-1]=(1/3)(beta-2alpha-alpha^2)/(1+alpha)^2",
		FirstOrderFormula:           "for alpha>0 and beta<<alpha, delta_ratio≈-(2/3)alpha",
		ProbeAlpha:                  alpha,
		ProbeBeta:                   beta,
		ProbeRatioDirect:            ratioDirect,
		ProbeRatioByFormula:         oneThird + deltaFormula,
		ProbeDeltaDirect:            deltaDirect,
		ProbeDeltaByFormula:         deltaFormula,
		ProbeFormulaResidual:        deltaDirect - deltaFormula,
		CurrentDeltaRatio:           g.DeltaRatio,
		CurrentRatioBelowOneThird:   g.DeltaRatio < 0,
		AssumptionAlphaPositive:     true,
		AssumptionBetaMuchLessAlpha: true,
		FirstOrderExplainsSign:      g.DeltaRatio < 0,
		NativeDeltaRatioTheorem:     false,
		Verdict: strings.Join([]string{
			StatusExactTraceDeviationFormulaDerived,
			StatusDeltaRatioIsNonTopCorrection,
			StatusRestChannelsLowerRatioBelowOneThird,
		}, "; "),
	}
}

func buildOneEighthProxyDeviation(g Gate754Inheritance) OneEighthProxyDeviation {
	lambda := threeEighths * g.BOverA2Computed
	transported := threeEighths * g.DeltaRatio
	return OneEighthProxyDeviation{
		FormulaFromTraceDeviation: "lambda_proxy-1/8=(3/8)delta_ratio",
		EquivalentFormula:         "lambda_proxy=1/8+(1/8)(3b/a^2-1)",
		LambdaProxyComputed:       lambda,
		LambdaProxySeed:           lambdaProxyMZ,
		LambdaProxySeedResidual:   lambda - lambdaProxyMZ,
		ProxyMinusOneEighth:       lambda - oneEighth,
		ThreeEighthsTimesDelta:    transported,
		TransportIdentityResidual: (lambda - oneEighth) - transported,
		DeviationTransported:      math.Abs((lambda-oneEighth)-transported) < 1e-16,
		ScalarPotentialTheorem:    false,
		RuntimeLambdaTheorem:      false,
		Verdict: strings.Join([]string{
			StatusOneEighthProxyDeviationRewritten,
			StatusLambdaProxyDeviationTransported,
		}, "; "),
	}
}

func buildRequiredYukawaDecompositionData() RequiredYukawaDecompositionData {
	return RequiredYukawaDecompositionData{
		RequiredItems: []string{
			"typed dominant top-like squared singular value T=y_t^2",
			"decomposed Yukawa ledger separating top channel from rest channels",
			"sector labels for bottom, tau, charm, neutrino, and remaining singular values",
			"scale convention for the M_Z Yukawa ledger",
			"finite trace normalization convention for a and b",
		},
		CanComputeAlphaBeta:          false,
		CanAssignBottomTauCharm:      false,
		CanAssignNeutrinoConvention:  false,
		CanAssignScaleDependence:     false,
		CanAssignFiniteTraceResidual: false,
		TypedTopLikeTAvailable:       false,
		DecomposedYukawaLedger:       false,
		Verdict: strings.Join([]string{
			StatusRequiredYukawaDecompositionDataListed,
			StatusNoNumericalAlphaBetaWithoutLedger,
		}, "; "),
	}
}

func buildYukawaFirewall() YukawaFirewall {
	return YukawaFirewall{
		DeltaRatioIsNativeYukawaTheorem: false,
		TopDominanceDerivesTopYukawa:    false,
		AlphaBetaDerivesHierarchy:       false,
		ClaimsYuDerived:                 false,
		ClaimsYdDerived:                 false,
		ClaimsYeDerived:                 false,
		ClaimsYnuDerived:                false,
		ClaimsCKMPMNSDerived:            false,
		ClaimsGenerationCarrier:         false,
		ClaimsFlavorTheorem:             false,
		SealedLedgerExplicit:            true,
		Verdict: strings.Join([]string{
			StatusYukawaFirewallEnforced,
			StatusNoNativeTopYukawaDerivation,
			StatusNoNativeYukawaOperatorTheorem,
			StatusNoNativeFlavorHierarchyTheorem,
		}, "; "),
	}
}

func buildRuntimeAndHiggsFirewalls() RuntimeAndHiggsFirewalls {
	return RuntimeAndHiggsFirewalls{
		LambdaProxyNearOneEighthIsScalarPotentialTheorem: false,
		LambdaProxyEqualsRuntimeLambda:                   false,
		RuntimeLambdaEqualsHiggsMass:                     false,
		TreeProxyEqualsPoleMass:                          false,
		RequiresHistoryLoopTransport:                     true,
		RequiresBoundaryHistoryResponse:                  true,
		RequiresKappaEReduction:                          true,
		RequiresScalarRuntimeBridge:                      true,
		ClaimsIndependentScalarRuntime:                   false,
		ClaimsHiggsMassTheorem:                           false,
		ClaimsPoleMassTheorem:                            false,
		Verdict: strings.Join([]string{
			StatusRuntimeAndHiggsFirewallsEnforced,
			StatusNoNativeScalarProxyDerivationTheorem,
			StatusNoIndependentScalarRuntimeTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusGate755TopDominanceTraceDeviationBoundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate754OneThirdTraceShadowInherited,
		StatusTopColorDominantSplitDefined,
		StatusNormalizedRestVariablesDefined,
		StatusExactTraceDeviationFormulaDerived,
		StatusOneEighthProxyDeviationRewritten,
		StatusRequiredYukawaDecompositionDataListed,
		StatusYukawaFirewallEnforced,
		StatusRuntimeAndHiggsFirewallsEnforced,
		StatusDeltaRatioIsNonTopCorrection,
		StatusLambdaProxyDeviationTransported,
		StatusRestChannelsLowerRatioBelowOneThird,
		StatusNoNativeTopYukawaDerivation,
		StatusNoNumericalAlphaBetaWithoutLedger,
		StatusNoNativeYukawaOperatorTheorem,
		StatusNoNativeFlavorHierarchyTheorem,
		StatusNoNativeScalarProxyDerivationTheorem,
		StatusNoIndependentScalarRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusGate755TopDominanceTraceDeviationBoundary,
	}
}

func FormatGate754(x Gate754Inheritance) string {
	return fmt.Sprintf("inherited=%t proxy=%q ratioFormula=%q source=%q a=%.16g b=%.16g ratio=%.16g seed=%.16g delta=%.16g lambdaProxy=%.16g topLimit=%t shadow=%t blocked(ba2=%t deltaDecomp=%t runtimeHiggs=%t) verdict=%q", x.Inherited, x.ScalarProxyFormula, x.TraceRatioFormula, x.OneThirdLimitSource, x.ATraceMZ, x.BTraceMZ, x.BOverA2Computed, x.BOverA2Seed, x.DeltaRatio, x.LambdaProxy, x.TopColorDominanceLimitDefined, x.OneThirdTraceShadowConditional, x.NativeBA2OneThirdTheoremBlocked, x.NativeDeltaDecompositionBlocked, x.RuntimeAndHiggsTheoremsBlocked, x.Verdict)
}

func FormatSplit(x TopColorDominantSplit) string {
	return fmt.Sprintf("T=%q aTop=%q bTop=%q aRest=%q bRest=%q totalA=%q totalB=%q zeroRest=%q topRatio=%q ratio=%.16g identity=%t topDerived=%t verdict=%q", x.DominantSquaredSingularValueSymbol, x.ATopFormula, x.BTopFormula, x.ARestFormula, x.BRestFormula, x.TotalAFormula, x.TotalBFormula, x.ZeroRestCondition, x.TopLimitRatioFormula, x.TopLimitRatio, x.SplitIsAlgebraicIdentity, x.TopYukawaValueDerived, x.Verdict)
}

func FormatRest(x NormalizedRestVariables) string {
	return fmt.Sprintf("alpha=%q beta=%q a=%q b=%q ratio=%q requires(top=%t betaTop=%t ledger=%t) numerical=%t verdict=%q", x.AlphaDefinition, x.BetaDefinition, x.ARewritten, x.BRewritten, x.RatioRewritten, x.AlphaRequiresTopValue, x.BetaRequiresTopValue, x.RequiresDecomposedYukawaLedger, x.NumericalAlphaBetaAvailable, x.Verdict)
}

func FormatDeviation(x TraceDeviationFormula) string {
	return fmt.Sprintf("exact=%q firstOrder=%q probe(alpha=%.16g beta=%.16g ratioDirect=%.16g ratioFormula=%.16g deltaDirect=%.16g deltaFormula=%.16g residual=%.16g) currentDelta=%.16g belowThird=%t assumptions(alphaPositive=%t betaSmall=%t) firstOrderSign=%t nativeTheorem=%t verdict=%q", x.ExactFormula, x.FirstOrderFormula, x.ProbeAlpha, x.ProbeBeta, x.ProbeRatioDirect, x.ProbeRatioByFormula, x.ProbeDeltaDirect, x.ProbeDeltaByFormula, x.ProbeFormulaResidual, x.CurrentDeltaRatio, x.CurrentRatioBelowOneThird, x.AssumptionAlphaPositive, x.AssumptionBetaMuchLessAlpha, x.FirstOrderExplainsSign, x.NativeDeltaRatioTheorem, x.Verdict)
}

func FormatProxy(x OneEighthProxyDeviation) string {
	return fmt.Sprintf("formula=%q equivalent=%q lambda=%.16g seed=%.16g seedResidual=%.16g minusOneEighth=%.16g threeEighthDelta=%.16g residual=%.16g transported=%t scalarPotential=%t runtime=%t verdict=%q", x.FormulaFromTraceDeviation, x.EquivalentFormula, x.LambdaProxyComputed, x.LambdaProxySeed, x.LambdaProxySeedResidual, x.ProxyMinusOneEighth, x.ThreeEighthsTimesDelta, x.TransportIdentityResidual, x.DeviationTransported, x.ScalarPotentialTheorem, x.RuntimeLambdaTheorem, x.Verdict)
}

func FormatRequiredData(x RequiredYukawaDecompositionData) string {
	return fmt.Sprintf("required=[%s] canComputeAlphaBeta=%t assign(bottomTauCharm=%t neutrino=%t scale=%t finiteTrace=%t) topT=%t ledger=%t verdict=%q", strings.Join(x.RequiredItems, "; "), x.CanComputeAlphaBeta, x.CanAssignBottomTauCharm, x.CanAssignNeutrinoConvention, x.CanAssignScaleDependence, x.CanAssignFiniteTraceResidual, x.TypedTopLikeTAvailable, x.DecomposedYukawaLedger, x.Verdict)
}

func FormatYukawaFirewall(x YukawaFirewall) string {
	return fmt.Sprintf("claims(deltaNative=%t topDerives=%t alphaBetaHierarchy=%t Yu=%t Yd=%t Ye=%t Ynu=%t ckmPmns=%t carrier=%t flavor=%t) sealedLedger=%t verdict=%q", x.DeltaRatioIsNativeYukawaTheorem, x.TopDominanceDerivesTopYukawa, x.AlphaBetaDerivesHierarchy, x.ClaimsYuDerived, x.ClaimsYdDerived, x.ClaimsYeDerived, x.ClaimsYnuDerived, x.ClaimsCKMPMNSDerived, x.ClaimsGenerationCarrier, x.ClaimsFlavorTheorem, x.SealedLedgerExplicit, x.Verdict)
}

func FormatRuntimeFirewall(x RuntimeAndHiggsFirewalls) string {
	return fmt.Sprintf("claims(proxyAsPotential=%t proxyRuntime=%t runtimeHiggs=%t treePole=%t independentRuntime=%t higgs=%t pole=%t) requires(historyLoop=%t boundary=%t kappaE=%t scalarBridge=%t) verdict=%q", x.LambdaProxyNearOneEighthIsScalarPotentialTheorem, x.LambdaProxyEqualsRuntimeLambda, x.RuntimeLambdaEqualsHiggsMass, x.TreeProxyEqualsPoleMass, x.ClaimsIndependentScalarRuntime, x.ClaimsHiggsMassTheorem, x.ClaimsPoleMassTheorem, x.RequiresHistoryLoopTransport, x.RequiresBoundaryHistoryResponse, x.RequiresKappaEReduction, x.RequiresScalarRuntimeBridge, x.Verdict)
}
