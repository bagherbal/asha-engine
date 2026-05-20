// Package generation2yukawatraceparticipationratioandeffectivetopcolorchannelcountaudit implements
// Gate 756: Yukawa Trace Participation Ratio and Effective Top-Color Channel Count Audit.
//
// Gate 755 expanded the b/a^2 deviation from the one-third top-color shadow by
// choosing a top/rest split. Gate 756 audits the complementary basis-clean
// diagnostic available already from the aggregate trace pair (a,b): expand the
// color-weighted Yukawa trace ledger into positive trace atoms x_i, rewrite
// b/a^2 as sum_i (x_i/a)^2, and type the effective channel count
// N_eff=a^2/b. The gate preserves the Yukawa, flavor, scalar-runtime,
// Higgs-mass, and pole-mass firewalls.
package generation2yukawatraceparticipationratioandeffectivetopcolorchannelcountaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE756-YUKAWA-TRACE-PARTICIPATION-RATIO-AND-EFFECTIVE-TOP-COLOR-CHANNEL-COUNT-AUDIT"

	StatusGate755TopDominanceTraceDeviationInherited = "PASS_GATE755_TOP_DOMINANCE_TRACE_DEVIATION_INHERITED"
	StatusTraceAtomExpansionDefined                  = "PASS_TRACE_ATOM_EXPANSION_DEFINED"
	StatusBA2TypedAsInverseParticipationRatio        = "PASS_B_OVER_A_SQUARED_TYPED_AS_INVERSE_PARTICIPATION_RATIO"
	StatusEffectiveChannelCountComputed              = "PASS_EFFECTIVE_CHANNEL_COUNT_COMPUTED"
	StatusTopColorEffectiveCountComparisonAudited    = "PASS_TOP_COLOR_EFFECTIVE_COUNT_COMPARISON_AUDITED"
	StatusOneEighthProxyRewrittenUsingNEff           = "PASS_ONE_EIGHTH_PROXY_REWRITTEN_USING_N_EFF"
	StatusRelationToGate755AlphaBetaRecorded         = "PASS_RELATION_TO_GATE755_ALPHA_BETA_FORM_RECORDED"
	StatusYukawaFirewallEnforced                     = "PASS_YUKAWA_FIREWALL_ENFORCED"
	StatusRuntimeAndHiggsFirewallsEnforced           = "PASS_RUNTIME_AND_HIGGS_FIREWALLS_ENFORCED"

	StatusBA2IsYukawaTraceParticipationRatio       = "CONDITIONAL_SUPPORT_B_OVER_A_SQUARED_IS_YUKAWA_TRACE_PARTICIPATION_RATIO"
	StatusNEffNearThreeTopColorDominance           = "CONDITIONAL_SUPPORT_N_EFF_NEAR_THREE_SOURCE_TYPES_TOP_COLOR_DOMINANCE"
	StatusNonTopChannelsTinyEffectiveParticipation = "CONDITIONAL_SUPPORT_NON_TOP_CHANNELS_APPEAR_AS_TINY_EFFECTIVE_TRACE_PARTICIPATION"
	StatusLambdaProxyEqualsThreeOverEightNEff      = "CONDITIONAL_SUPPORT_LAMBDA_PROXY_EQUALS_THREE_OVER_EIGHT_N_EFF"
	StatusNEffNotNativeGenerationTheorem           = "FAILED_ROUTE_N_EFF_NOT_NATIVE_GENERATION_THEOREM"
	StatusNoChannelAssignmentWithoutLedger         = "FAILED_ROUTE_NO_CHANNEL_ASSIGNMENT_WITHOUT_DECOMPOSED_YUKAWA_LEDGER"
	StatusNoNativeYukawaOperatorTheorem            = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusNoNativeFlavorHierarchyTheorem           = "FAILED_ROUTE_NO_NATIVE_FLAVOR_HIERARCHY_THEOREM"
	StatusNoNativeScalarProxyDerivationTheorem     = "FAILED_ROUTE_NO_NATIVE_SCALAR_PROXY_DERIVATION_THEOREM"
	StatusNoIndependentScalarRuntimeTheorem        = "FAILED_ROUTE_NO_INDEPENDENT_SCALAR_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem             = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusGate756YukawaTraceParticipationBoundary  = "FIREWALL_PRESERVED_GATE756_YUKAWA_TRACE_PARTICIPATION_BOUNDARY"
)

const (
	threeEighths = 3.0 / 8.0
	oneThird     = 1.0 / 3.0
	oneEighth    = 1.0 / 8.0

	aTraceMZ      = 2.8424095142339083
	bTraceMZ      = 2.6910096440382287
	bOverA2MZSeed = 0.33307493962706697
	lambdaProxyMZ = 0.12490310236015012
)

type Gate755Inheritance struct {
	Inherited                            bool
	TopRestFormula                       string
	TraceDeviationFormula                string
	AggregateTracePairAvailable          bool
	RequiresNoTopYukawaChoice            bool
	ATraceMZ                             float64
	BTraceMZ                             float64
	BOverA2Computed                      float64
	BOverA2Seed                          float64
	DeltaRatio                           float64
	LambdaProxy                          float64
	TopDominanceDeviationTyped           bool
	NumericalAlphaBetaBlocked            bool
	NativeYukawaAndScalarTheoremsBlocked bool
	Verdict                              string
}

type TraceAtomExpansion struct {
	AtomDefinition                         string
	QuarkColorMultiplicityRule             string
	AFormula                               string
	BFormula                               string
	NormalizedWeightDefinition             string
	WeightSumFormula                       string
	IPRFormula                             string
	ColorFactorExpandedAsRepeatedAtoms     bool
	AtomsPositive                          bool
	RequiresDecomposedYukawaLedgerForAtoms bool
	UsesOnlyAggregatePairForNEff           bool
	Verdict                                string
}

type InverseParticipationRatio struct {
	TraceRatioFormula             string
	IPRFormula                    string
	ComputedRatio                 float64
	SeedRatio                     float64
	SeedResidual                  float64
	SyntheticTopColorWeights      []float64
	SyntheticTopColorIPR          float64
	TopColorIPR                   float64
	RatioIsIPR                    bool
	BasisCleanAggregateDiagnostic bool
	NativeYukawaTheorem           bool
	Verdict                       string
}

type EffectiveChannelCount struct {
	Definition                     string
	ComputedFromRatio              float64
	ComputedFromTracePair          float64
	TracePairResidual              float64
	TopColorValue                  float64
	DeviationFromThree             float64
	RelativeDeviationFromThree     float64
	CurrentLedgerAboveThree        bool
	NearThree                      bool
	InterpretedAsTinyTraceSpread   bool
	NativeGenerationTheorem        bool
	ChannelAssignmentWithoutLedger bool
	Verdict                        string
}

type OneEighthProxyRewrite struct {
	LambdaProxyFormula         string
	ParticipationFormula       string
	OneEighthLimitFormula      string
	EquivalentFormula          string
	LambdaProxyComputed        float64
	LambdaProxySeed            float64
	LambdaProxySeedResidual    float64
	TopColorProxyLimit         float64
	ProxyBelowOneEighth        bool
	ThreeOverEightNEffIdentity bool
	ScalarPotentialTheorem     bool
	RuntimeLambdaTheorem       bool
	Verdict                    string
}

type RelationToGate755 struct {
	Gate755TopRestFormula              string
	Gate756ParticipationFormula        string
	NEffAlphaBetaFormula               string
	CompatibilityProbeAlpha            float64
	CompatibilityProbeBeta             float64
	ProbeRatioGate755                  float64
	ProbeNEffGate756                   float64
	ProbeInverseNEff                   float64
	ProbeCompatibilityResidual         float64
	Gate755NeedsDecomposedLedger       bool
	Gate756WorksFromAggregateTracePair bool
	CompatibleDiagnostics              bool
	Verdict                            string
}

type YukawaFirewall struct {
	NEffIsNativeGenerationTheorem   bool
	NEffDerivesFlavorHierarchy      bool
	NEffMinusThreeAssignedToChannel bool
	ClaimsYuDerived                 bool
	ClaimsYdDerived                 bool
	ClaimsYeDerived                 bool
	ClaimsYnuDerived                bool
	ClaimsCKMPMNSDerived            bool
	ClaimsNativeFlavorTheorem       bool
	SealedLedgerExplicit            bool
	Verdict                         string
}

type RuntimeAndHiggsFirewalls struct {
	LambdaProxyNearOneEighthIsScalarPotentialTheorem bool
	LambdaProxyEqualsRuntimeLambda                   bool
	RuntimeLambdaEqualsHiggsMass                     bool
	TreeProxyEqualsPoleMass                          bool
	ClaimsIndependentScalarRuntime                   bool
	ClaimsHiggsMassTheorem                           bool
	ClaimsPoleMassTheorem                            bool
	RequiresHistoryLoopTransport                     bool
	RequiresBoundaryHistoryResponse                  bool
	RequiresKappaEReduction                          bool
	RequiresScalarRuntimeBridge                      bool
	Verdict                                          string
}

type Analysis struct {
	Gate755   Gate755Inheritance
	Atoms     TraceAtomExpansion
	IPR       InverseParticipationRatio
	Effective EffectiveChannelCount
	Proxy     OneEighthProxyRewrite
	Relation  RelationToGate755
	Yukawa    YukawaFirewall
	Runtime   RuntimeAndHiggsFirewalls
	Truth     string
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
	gate755 := buildGate755Inheritance()
	if gate755.ATraceMZ <= 0 || gate755.BTraceMZ <= 0 || math.IsNaN(gate755.BOverA2Computed) || math.IsInf(gate755.BOverA2Computed, 0) {
		return Analysis{}, fmt.Errorf("invalid Gate755 aggregate trace inheritance: a=%g b=%g ratio=%g", gate755.ATraceMZ, gate755.BTraceMZ, gate755.BOverA2Computed)
	}
	atoms := buildTraceAtomExpansion()
	ipr := buildInverseParticipationRatio(gate755)
	effective := buildEffectiveChannelCount(gate755)
	proxy := buildOneEighthProxyRewrite(effective)
	relation := buildRelationToGate755()
	yukawa := buildYukawaFirewall()
	runtime := buildRuntimeAndHiggsFirewalls()
	truth := "Gate 756 rewrites the sealed Yukawa trace-shape ratio as an inverse participation ratio over positive trace atoms after expanding quark color multiplicities into repeated atoms: b/a^2=sum_i(x_i/a)^2. The aggregate ledger gives N_eff=a^2/b=3.0023273474722147, just above the top-color value 3, so the current scalar proxy is lambda_proxy=3/(8N_eff)=0.12490310236015012, slightly below 1/8. This is a participation diagnostic only: it source-types the near-three effective count as top-color dominance plus tiny non-top trace spread, while blocking channel assignment, generation/flavor derivation, native Yukawa eigenvalue claims, runtime scalar lambda, Higgs mass, and pole mass."

	return Analysis{Gate755: gate755, Atoms: atoms, IPR: ipr, Effective: effective, Proxy: proxy, Relation: relation, Yukawa: yukawa, Runtime: runtime, Truth: truth}, nil
}

func buildGate755Inheritance() Gate755Inheritance {
	ratio := bTraceMZ / (aTraceMZ * aTraceMZ)
	return Gate755Inheritance{
		Inherited:                            true,
		TopRestFormula:                       "b/a^2=(1/3)(1+beta)/(1+alpha)^2",
		TraceDeviationFormula:                "delta_ratio=b/a^2-1/3=(1/3)(beta-2alpha-alpha^2)/(1+alpha)^2",
		AggregateTracePairAvailable:          true,
		RequiresNoTopYukawaChoice:            true,
		ATraceMZ:                             aTraceMZ,
		BTraceMZ:                             bTraceMZ,
		BOverA2Computed:                      ratio,
		BOverA2Seed:                          bOverA2MZSeed,
		DeltaRatio:                           ratio - oneThird,
		LambdaProxy:                          threeEighths * ratio,
		TopDominanceDeviationTyped:           true,
		NumericalAlphaBetaBlocked:            true,
		NativeYukawaAndScalarTheoremsBlocked: true,
		Verdict: strings.Join([]string{
			StatusGate755TopDominanceTraceDeviationInherited,
			StatusNoChannelAssignmentWithoutLedger,
		}, "; "),
	}
}

func buildTraceAtomExpansion() TraceAtomExpansion {
	return TraceAtomExpansion{
		AtomDefinition:                         "positive Yukawa trace atoms x_i contributed by squared singular values",
		QuarkColorMultiplicityRule:             "quark color factor 3 is represented as three repeated atoms with the same squared singular value",
		AFormula:                               "a=sum_i x_i",
		BFormula:                               "b=sum_i x_i^2",
		NormalizedWeightDefinition:             "w_i=x_i/a",
		WeightSumFormula:                       "sum_i w_i=1",
		IPRFormula:                             "b/a^2=sum_i w_i^2",
		ColorFactorExpandedAsRepeatedAtoms:     true,
		AtomsPositive:                          true,
		RequiresDecomposedYukawaLedgerForAtoms: true,
		UsesOnlyAggregatePairForNEff:           true,
		Verdict: strings.Join([]string{
			StatusTraceAtomExpansionDefined,
			StatusBA2IsYukawaTraceParticipationRatio,
		}, "; "),
	}
}

func buildInverseParticipationRatio(g Gate755Inheritance) InverseParticipationRatio {
	weights := []float64{oneThird, oneThird, oneThird}
	synthetic := 0.0
	for _, w := range weights {
		synthetic += w * w
	}
	return InverseParticipationRatio{
		TraceRatioFormula:             "b/a^2",
		IPRFormula:                    "IPR=sum_i w_i^2 where w_i=x_i/a",
		ComputedRatio:                 g.BOverA2Computed,
		SeedRatio:                     g.BOverA2Seed,
		SeedResidual:                  g.BOverA2Computed - g.BOverA2Seed,
		SyntheticTopColorWeights:      weights,
		SyntheticTopColorIPR:          synthetic,
		TopColorIPR:                   oneThird,
		RatioIsIPR:                    true,
		BasisCleanAggregateDiagnostic: true,
		NativeYukawaTheorem:           false,
		Verdict: strings.Join([]string{
			StatusBA2TypedAsInverseParticipationRatio,
			StatusBA2IsYukawaTraceParticipationRatio,
		}, "; "),
	}
}

func buildEffectiveChannelCount(g Gate755Inheritance) EffectiveChannelCount {
	fromRatio := 1.0 / g.BOverA2Computed
	fromTracePair := (g.ATraceMZ * g.ATraceMZ) / g.BTraceMZ
	dev := fromRatio - 3.0
	return EffectiveChannelCount{
		Definition:                     "N_eff=1/(b/a^2)=a^2/b",
		ComputedFromRatio:              fromRatio,
		ComputedFromTracePair:          fromTracePair,
		TracePairResidual:              fromRatio - fromTracePair,
		TopColorValue:                  3.0,
		DeviationFromThree:             dev,
		RelativeDeviationFromThree:     dev / 3.0,
		CurrentLedgerAboveThree:        fromRatio > 3.0,
		NearThree:                      math.Abs(dev) < 0.003,
		InterpretedAsTinyTraceSpread:   fromRatio > 3.0 && math.Abs(dev/3.0) < 0.001,
		NativeGenerationTheorem:        false,
		ChannelAssignmentWithoutLedger: false,
		Verdict: strings.Join([]string{
			StatusEffectiveChannelCountComputed,
			StatusTopColorEffectiveCountComparisonAudited,
			StatusNEffNearThreeTopColorDominance,
			StatusNonTopChannelsTinyEffectiveParticipation,
			StatusNEffNotNativeGenerationTheorem,
		}, "; "),
	}
}

func buildOneEighthProxyRewrite(e EffectiveChannelCount) OneEighthProxyRewrite {
	lambda := 3.0 / (8.0 * e.ComputedFromRatio)
	return OneEighthProxyRewrite{
		LambdaProxyFormula:         "lambda_proxy=(3/8)(b/a^2)",
		ParticipationFormula:       "lambda_proxy=3/(8N_eff)",
		OneEighthLimitFormula:      "N_eff=3 => lambda_proxy=1/8",
		EquivalentFormula:          "lambda_proxy=(1/8)(3/N_eff)",
		LambdaProxyComputed:        lambda,
		LambdaProxySeed:            lambdaProxyMZ,
		LambdaProxySeedResidual:    lambda - lambdaProxyMZ,
		TopColorProxyLimit:         oneEighth,
		ProxyBelowOneEighth:        lambda < oneEighth,
		ThreeOverEightNEffIdentity: math.Abs(lambda-(oneEighth*(3.0/e.ComputedFromRatio))) < 1e-16,
		ScalarPotentialTheorem:     false,
		RuntimeLambdaTheorem:       false,
		Verdict: strings.Join([]string{
			StatusOneEighthProxyRewrittenUsingNEff,
			StatusLambdaProxyEqualsThreeOverEightNEff,
		}, "; "),
	}
}

func buildRelationToGate755() RelationToGate755 {
	// Probe only certifies algebraic compatibility between the alpha/beta and N_eff forms.
	// It is not a physical Yukawa ledger.
	const alpha = 0.012
	const beta = 0.00003
	ratio755 := oneThird * (1.0 + beta) / ((1.0 + alpha) * (1.0 + alpha))
	neff756 := 3.0 * (1.0 + alpha) * (1.0 + alpha) / (1.0 + beta)
	return RelationToGate755{
		Gate755TopRestFormula:              "b/a^2=(1/3)(1+beta)/(1+alpha)^2",
		Gate756ParticipationFormula:        "b/a^2=1/N_eff",
		NEffAlphaBetaFormula:               "N_eff=3(1+alpha)^2/(1+beta)",
		CompatibilityProbeAlpha:            alpha,
		CompatibilityProbeBeta:             beta,
		ProbeRatioGate755:                  ratio755,
		ProbeNEffGate756:                   neff756,
		ProbeInverseNEff:                   1.0 / neff756,
		ProbeCompatibilityResidual:         ratio755 - (1.0 / neff756),
		Gate755NeedsDecomposedLedger:       true,
		Gate756WorksFromAggregateTracePair: true,
		CompatibleDiagnostics:              math.Abs(ratio755-(1.0/neff756)) < 1e-16,
		Verdict: strings.Join([]string{
			StatusRelationToGate755AlphaBetaRecorded,
			StatusNoChannelAssignmentWithoutLedger,
		}, "; "),
	}
}

func buildYukawaFirewall() YukawaFirewall {
	return YukawaFirewall{
		NEffIsNativeGenerationTheorem:   false,
		NEffDerivesFlavorHierarchy:      false,
		NEffMinusThreeAssignedToChannel: false,
		ClaimsYuDerived:                 false,
		ClaimsYdDerived:                 false,
		ClaimsYeDerived:                 false,
		ClaimsYnuDerived:                false,
		ClaimsCKMPMNSDerived:            false,
		ClaimsNativeFlavorTheorem:       false,
		SealedLedgerExplicit:            true,
		Verdict: strings.Join([]string{
			StatusYukawaFirewallEnforced,
			StatusNEffNotNativeGenerationTheorem,
			StatusNoChannelAssignmentWithoutLedger,
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
		ClaimsIndependentScalarRuntime:                   false,
		ClaimsHiggsMassTheorem:                           false,
		ClaimsPoleMassTheorem:                            false,
		RequiresHistoryLoopTransport:                     true,
		RequiresBoundaryHistoryResponse:                  true,
		RequiresKappaEReduction:                          true,
		RequiresScalarRuntimeBridge:                      true,
		Verdict: strings.Join([]string{
			StatusRuntimeAndHiggsFirewallsEnforced,
			StatusNoNativeScalarProxyDerivationTheorem,
			StatusNoIndependentScalarRuntimeTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusGate756YukawaTraceParticipationBoundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate755TopDominanceTraceDeviationInherited,
		StatusTraceAtomExpansionDefined,
		StatusBA2TypedAsInverseParticipationRatio,
		StatusEffectiveChannelCountComputed,
		StatusTopColorEffectiveCountComparisonAudited,
		StatusOneEighthProxyRewrittenUsingNEff,
		StatusRelationToGate755AlphaBetaRecorded,
		StatusYukawaFirewallEnforced,
		StatusRuntimeAndHiggsFirewallsEnforced,
		StatusBA2IsYukawaTraceParticipationRatio,
		StatusNEffNearThreeTopColorDominance,
		StatusNonTopChannelsTinyEffectiveParticipation,
		StatusLambdaProxyEqualsThreeOverEightNEff,
		StatusNEffNotNativeGenerationTheorem,
		StatusNoChannelAssignmentWithoutLedger,
		StatusNoNativeYukawaOperatorTheorem,
		StatusNoNativeFlavorHierarchyTheorem,
		StatusNoNativeScalarProxyDerivationTheorem,
		StatusNoIndependentScalarRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusGate756YukawaTraceParticipationBoundary,
	}
}

func FormatGate755(x Gate755Inheritance) string {
	return fmt.Sprintf("inherited=%t topRest=%q deviation=%q aggregatePair=%t noTopChoice=%t a=%.16g b=%.16g ratio=%.16g seed=%.16g delta=%.16g lambda=%.16g typed=%t alphaBetaBlocked=%t nativeBlocked=%t verdict=%q", x.Inherited, x.TopRestFormula, x.TraceDeviationFormula, x.AggregateTracePairAvailable, x.RequiresNoTopYukawaChoice, x.ATraceMZ, x.BTraceMZ, x.BOverA2Computed, x.BOverA2Seed, x.DeltaRatio, x.LambdaProxy, x.TopDominanceDeviationTyped, x.NumericalAlphaBetaBlocked, x.NativeYukawaAndScalarTheoremsBlocked, x.Verdict)
}

func FormatAtoms(x TraceAtomExpansion) string {
	return fmt.Sprintf("atom=%q colorRule=%q a=%q b=%q w=%q sum=%q ipr=%q expanded=%t positive=%t requiresLedger=%t aggregateNEff=%t verdict=%q", x.AtomDefinition, x.QuarkColorMultiplicityRule, x.AFormula, x.BFormula, x.NormalizedWeightDefinition, x.WeightSumFormula, x.IPRFormula, x.ColorFactorExpandedAsRepeatedAtoms, x.AtomsPositive, x.RequiresDecomposedYukawaLedgerForAtoms, x.UsesOnlyAggregatePairForNEff, x.Verdict)
}

func FormatIPR(x InverseParticipationRatio) string {
	parts := make([]string, 0, len(x.SyntheticTopColorWeights))
	for _, w := range x.SyntheticTopColorWeights {
		parts = append(parts, fmt.Sprintf("%.16g", w))
	}
	return fmt.Sprintf("ratioFormula=%q ipr=%q ratio=%.16g seed=%.16g residual=%.16g topWeights=[%s] topIPR=%.16g topColorIPR=%.16g isIPR=%t basisClean=%t nativeYukawa=%t verdict=%q", x.TraceRatioFormula, x.IPRFormula, x.ComputedRatio, x.SeedRatio, x.SeedResidual, strings.Join(parts, ", "), x.SyntheticTopColorIPR, x.TopColorIPR, x.RatioIsIPR, x.BasisCleanAggregateDiagnostic, x.NativeYukawaTheorem, x.Verdict)
}

func FormatEffective(x EffectiveChannelCount) string {
	return fmt.Sprintf("definition=%q nEffRatio=%.16g nEffTrace=%.16g residual=%.16g top=%.16g deviation=%.16g relative=%.16g aboveThree=%t near=%t tinySpread=%t nativeGeneration=%t channelAssignment=%t verdict=%q", x.Definition, x.ComputedFromRatio, x.ComputedFromTracePair, x.TracePairResidual, x.TopColorValue, x.DeviationFromThree, x.RelativeDeviationFromThree, x.CurrentLedgerAboveThree, x.NearThree, x.InterpretedAsTinyTraceSpread, x.NativeGenerationTheorem, x.ChannelAssignmentWithoutLedger, x.Verdict)
}

func FormatProxy(x OneEighthProxyRewrite) string {
	return fmt.Sprintf("formula=%q participation=%q limit=%q equivalent=%q lambda=%.16g seed=%.16g seedResidual=%.16g topLimit=%.16g belowOneEighth=%t identity=%t scalarPotential=%t runtime=%t verdict=%q", x.LambdaProxyFormula, x.ParticipationFormula, x.OneEighthLimitFormula, x.EquivalentFormula, x.LambdaProxyComputed, x.LambdaProxySeed, x.LambdaProxySeedResidual, x.TopColorProxyLimit, x.ProxyBelowOneEighth, x.ThreeOverEightNEffIdentity, x.ScalarPotentialTheorem, x.RuntimeLambdaTheorem, x.Verdict)
}

func FormatRelation(x RelationToGate755) string {
	return fmt.Sprintf("gate755=%q gate756=%q nEff=%q probe(alpha=%.16g beta=%.16g ratio755=%.16g nEff756=%.16g invNEff=%.16g residual=%.16g) needsLedger=%t aggregatePair=%t compatible=%t verdict=%q", x.Gate755TopRestFormula, x.Gate756ParticipationFormula, x.NEffAlphaBetaFormula, x.CompatibilityProbeAlpha, x.CompatibilityProbeBeta, x.ProbeRatioGate755, x.ProbeNEffGate756, x.ProbeInverseNEff, x.ProbeCompatibilityResidual, x.Gate755NeedsDecomposedLedger, x.Gate756WorksFromAggregateTracePair, x.CompatibleDiagnostics, x.Verdict)
}

func FormatYukawaFirewall(x YukawaFirewall) string {
	return fmt.Sprintf("claims(nEffGeneration=%t hierarchy=%t channelAssign=%t Yu=%t Yd=%t Ye=%t Ynu=%t ckmPmns=%t flavor=%t) sealedLedger=%t verdict=%q", x.NEffIsNativeGenerationTheorem, x.NEffDerivesFlavorHierarchy, x.NEffMinusThreeAssignedToChannel, x.ClaimsYuDerived, x.ClaimsYdDerived, x.ClaimsYeDerived, x.ClaimsYnuDerived, x.ClaimsCKMPMNSDerived, x.ClaimsNativeFlavorTheorem, x.SealedLedgerExplicit, x.Verdict)
}

func FormatRuntimeFirewall(x RuntimeAndHiggsFirewalls) string {
	return fmt.Sprintf("claims(proxyAsPotential=%t proxyRuntime=%t runtimeHiggs=%t treePole=%t independentRuntime=%t higgs=%t pole=%t) requires(historyLoop=%t boundary=%t kappaE=%t scalarBridge=%t) verdict=%q", x.LambdaProxyNearOneEighthIsScalarPotentialTheorem, x.LambdaProxyEqualsRuntimeLambda, x.RuntimeLambdaEqualsHiggsMass, x.TreeProxyEqualsPoleMass, x.ClaimsIndependentScalarRuntime, x.ClaimsHiggsMassTheorem, x.ClaimsPoleMassTheorem, x.RequiresHistoryLoopTransport, x.RequiresBoundaryHistoryResponse, x.RequiresKappaEReduction, x.RequiresScalarRuntimeBridge, x.Verdict)
}
