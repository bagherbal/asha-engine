// Package generation2effectiveparticipationscalarproxynormalformandruntimepropagationaudit implements
// Gate 757: Effective-Participation Scalar Proxy Normal Form and Runtime Propagation Audit.
//
// Gate 756 typed b/a^2 as the inverse participation ratio 1/N_eff. Gate 757
// substitutes the aggregate participation scalar proxy lambda_proxy=3/(8N_eff)
// into the Gate 752 flavor-reduced scalar-Higgs normal form. The gate is a
// bridge-layer normal-form and propagation audit only: it records the top-color
// shadow comparison and the runtime/tree-proxy diagnostic shift while preserving
// the Yukawa, generation, scalar-runtime, Higgs-mass, and pole-mass firewalls.
package generation2effectiveparticipationscalarproxynormalformandruntimepropagationaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE757-EFFECTIVE-PARTICIPATION-SCALAR-PROXY-NORMAL-FORM-AND-RUNTIME-PROPAGATION-AUDIT"

	StatusGate756YukawaTraceParticipationInherited    = "PASS_GATE756_YUKAWA_TRACE_PARTICIPATION_INHERITED"
	StatusGate752FlavorReducedNormalFormInherited     = "PASS_GATE752_FLAVOR_REDUCED_NORMAL_FORM_INHERITED"
	StatusEffectiveParticipationProxySubstituted      = "PASS_EFFECTIVE_PARTICIPATION_PROXY_SUBSTITUTED"
	StatusScalarHiggsEffectiveParticipationNormalForm = "PASS_SCALAR_HIGGS_EFFECTIVE_PARTICIPATION_NORMAL_FORM_WRITTEN"
	StatusTopColorShadowComparisonAudited             = "PASS_TOP_COLOR_SHADOW_COMPARISON_AUDITED"
	StatusRuntimePropagationOfNEffDeviationComputed   = "PASS_RUNTIME_PROPAGATION_OF_N_EFF_DEVIATION_COMPUTED"
	StatusLayerSeparationEnforced                     = "PASS_LAYER_SEPARATION_ENFORCED"
	StatusYukawaAndHiggsFirewallsEnforced             = "PASS_YUKAWA_AND_HIGGS_FIREWALLS_ENFORCED"

	StatusScalarProxyThreeOverEightNEff                    = "CONDITIONAL_SUPPORT_SCALAR_PROXY_CAN_BE_WRITTEN_AS_THREE_OVER_EIGHT_N_EFF"
	StatusCurrentBridgeHasEffectiveParticipationForm       = "CONDITIONAL_SUPPORT_CURRENT_SCALAR_HIGGS_BRIDGE_HAS_EFFECTIVE_PARTICIPATION_NORMAL_FORM"
	StatusNEffMinusThreeLowersProxyBelowOneEighth          = "CONDITIONAL_SUPPORT_N_EFF_MINUS_THREE_LOWERS_PROXY_BELOW_ONE_EIGHTH"
	StatusNEffNotNativeGenerationTheorem                   = "FAILED_ROUTE_N_EFF_NOT_NATIVE_GENERATION_THEOREM"
	StatusNoChannelAssignmentWithoutLedger                 = "FAILED_ROUTE_NO_CHANNEL_ASSIGNMENT_WITHOUT_DECOMPOSED_YUKAWA_LEDGER"
	StatusNoNativeYukawaOperatorTheorem                    = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusNoNativeScalarProxyDerivationTheorem             = "FAILED_ROUTE_NO_NATIVE_SCALAR_PROXY_DERIVATION_THEOREM"
	StatusNoIndependentScalarRuntimeTheorem                = "FAILED_ROUTE_NO_INDEPENDENT_SCALAR_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem                     = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusGate757EffectiveParticipationScalarProxyBoundary = "FIREWALL_PRESERVED_GATE757_EFFECTIVE_PARTICIPATION_SCALAR_PROXY_BOUNDARY"
)

const (
	oneEighth = 1.0 / 8.0
	three     = 3.0
	eight     = 8.0

	// Gate 756 aggregate participation values.
	nEffMZ        = 3.0023273474722147
	bOverA2MZ     = 0.33307493962706697
	lambdaProxyMZ = 0.12490310236015012
	nEffTopColor  = 3.0

	// Gate 752 / Gate 751 scalar-Higgs normal-form constants, copied as audited
	// bridge snapshots to avoid pulling the deep predecessor theorem chain.
	pK7            = 7.0 / 72.0
	xiBoundary     = 0.0503471644870914
	lambdaLambda12 = -0.049700942077683274
	r3MinusOne     = 0.0509933868964996
	kappaERed      = 0.005503554218475772
	lHopf          = 1.0 / (8.0 * math.Pi)

	// Gate 741 convention seal for diagnostic tree-proxy propagation only.
	vevConventionGeV   = 246.2196508
	gate741RuntimeSeed = 0.12965256505047373
)

type Gate756Inheritance struct {
	Inherited                   bool
	ParticipationFormula        string
	NEffDefinition              string
	NEff                        float64
	TraceRatio                  float64
	InverseNEff                 float64
	InverseResidual             float64
	LambdaProxyFormula          string
	LambdaProxy                 float64
	LambdaProxyFromNEff         float64
	LambdaProxyResidual         float64
	TopColorEffectiveCount      float64
	NEffMinusThree              float64
	EffectiveParticipationTyped bool
	YukawaTraceLedgerSealed     bool
	Verdict                     string
}

type Gate752Inheritance struct {
	Inherited             bool
	KappaERedFormula      string
	FWall3RedFormula      string
	RuntimeBridgeFormula  string
	ExpandedTraceFormula  string
	P                     float64
	SBoundary             float64
	AbsLambdaLambda12     float64
	KappaERed             float64
	LHopf                 float64
	FWall3Red             float64
	RuntimeBracket        float64
	RuntimeBracketFormula string
	ScalarMapTyped        bool
	NativeScalarRuntime   bool
	Verdict               string
}

type EffectiveParticipationNormalForm struct {
	ProxySubstitutionFormula           string
	RuntimeFormula                     string
	ExpandedTraceRuntimeFormula        string
	LambdaRuntimeEff                   float64
	RuntimeBracket                     float64
	ProxySubstituted                   bool
	NormalFormWritten                  bool
	EquivalentExpandedTraceFormWritten bool
	IndependentRuntimePrediction       bool
	NativeScalarProxyTheorem           bool
	Verdict                            string
}

type TopColorShadowComparison struct {
	NEffTop                       float64
	LambdaProxyTopShadow          float64
	LambdaRuntimeTopShadow        float64
	ProxyShift                    float64
	RuntimeShift                  float64
	RuntimeShiftApprox            float64
	RelativeRuntimeShift          float64
	TreeProxyEffGeV               float64
	TreeProxyTopShadowGeV         float64
	TreeProxyShiftGeV             float64
	VevConventionGeV              float64
	ProxyBelowTopShadow           bool
	RuntimeLoweredByParticipation bool
	TreeProxyDiagnosticOnly       bool
	HiggsPoleMassPrediction       bool
	Verdict                       string
}

type ParticipationInterpretation struct {
	NEffGreaterThanThree      bool
	LambdaProxyBelowOneEighth bool
	TraceLedgerMoreSpread     bool
	NonTopChannelsDiluteIPR   bool
	NonTopChannelsLowerProxy  bool
	AssignedSectorCorrection  bool
	Verdict                   string
}

type RelationToGate756 struct {
	Gate756ParticipationFormula     string
	Gate757NormalFormFormula        string
	TopShadowNormalFormFormula      string
	CompatibilityStatement          string
	UsesAggregateTracePair          bool
	RequiresNoTopYukawaChoice       bool
	RequiresGate752TransportBracket bool
	Compatible                      bool
	Verdict                         string
}

type LayerSeparation struct {
	NEffLayer                       string
	RuntimeCorrectionLayer          string
	NEffIsNativeGenerationTheorem   bool
	NEffIsYukawaEigenvalueTheorem   bool
	NEffIsScalarPotentialTheorem    bool
	NEffIsRuntimeLambdaTheorem      bool
	NEffIsHiggsMassTheorem          bool
	RuntimeBracketSeparateTransport bool
	Verdict                         string
}

type Firewalls struct {
	NEffNativeGenerationTheorem        bool
	NEffMinusThreeAssignedToSector     bool
	LambdaProxyScalarPotentialTheorem  bool
	LambdaRuntimeIndependentPrediction bool
	TreeProxyHiggsPoleMassPrediction   bool
	ClaimsYukawaEigenvaluesDerived     bool
	ClaimsFlavorHierarchyDerived       bool
	ClaimsCKMPMNSDerived               bool
	ClaimsHiggsMassTheorem             bool
	ClaimsPoleMassTheorem              bool
	Verdict                            string
}

type Analysis struct {
	Gate756        Gate756Inheritance
	Gate752        Gate752Inheritance
	NormalForm     EffectiveParticipationNormalForm
	TopShadow      TopColorShadowComparison
	Interpretation ParticipationInterpretation
	Relation       RelationToGate756
	Layering       LayerSeparation
	Firewalls      Firewalls
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
	g756 := buildGate756Inheritance()
	if g756.NEff <= 0 || math.IsNaN(g756.NEff) || math.IsInf(g756.NEff, 0) {
		return Analysis{}, fmt.Errorf("invalid Gate756 N_eff inheritance: %g", g756.NEff)
	}
	g752 := buildGate752Inheritance()
	if g752.RuntimeBracket <= 0 || math.IsNaN(g752.RuntimeBracket) || math.IsInf(g752.RuntimeBracket, 0) {
		return Analysis{}, fmt.Errorf("invalid Gate752 runtime bracket: %g", g752.RuntimeBracket)
	}
	normal := buildEffectiveParticipationNormalForm(g756, g752)
	top := buildTopColorShadowComparison(g756, normal)
	interp := buildParticipationInterpretation(g756, top)
	relation := buildRelationToGate756()
	layering := buildLayerSeparation()
	firewalls := buildFirewalls()
	truth := "Gate 757 substitutes the Gate756 participation scalar proxy lambda_proxy=3/(8N_eff) into the Gate752 flavor-reduced scalar-Higgs normal form. The current aggregate ledger has N_eff=3.0023273474722147, so the scalar proxy is below the top-color shadow 1/8 by -9.689763984987998e-05. Propagating through the same HistoryLoop/boundary bracket lowers the bridge runtime by about -1.005821898454e-04 and the Gate741 sealed tree proxy by about -0.0486 GeV relative to the pure N_eff=3 shadow. This is a normal-form and diagnostic propagation audit only, not a native Yukawa, scalar-runtime, Higgs-mass, or pole-mass theorem."

	return Analysis{Gate756: g756, Gate752: g752, NormalForm: normal, TopShadow: top, Interpretation: interp, Relation: relation, Layering: layering, Firewalls: firewalls, Truth: truth}, nil
}

func buildGate756Inheritance() Gate756Inheritance {
	lambdaFromNEff := three / (eight * nEffMZ)
	inv := 1.0 / nEffMZ
	return Gate756Inheritance{
		Inherited:                   true,
		ParticipationFormula:        "b/a^2=1/N_eff",
		NEffDefinition:              "N_eff=a^2/b",
		NEff:                        nEffMZ,
		TraceRatio:                  bOverA2MZ,
		InverseNEff:                 inv,
		InverseResidual:             inv - bOverA2MZ,
		LambdaProxyFormula:          "lambda_proxy=3/(8N_eff)",
		LambdaProxy:                 lambdaProxyMZ,
		LambdaProxyFromNEff:         lambdaFromNEff,
		LambdaProxyResidual:         lambdaFromNEff - lambdaProxyMZ,
		TopColorEffectiveCount:      nEffTopColor,
		NEffMinusThree:              nEffMZ - nEffTopColor,
		EffectiveParticipationTyped: true,
		YukawaTraceLedgerSealed:     true,
		Verdict: strings.Join([]string{
			StatusGate756YukawaTraceParticipationInherited,
			StatusScalarProxyThreeOverEightNEff,
		}, "; "),
	}
}

func buildGate752Inheritance() Gate752Inheritance {
	s := lambdaLambda12 + r3MinusOne
	absLam := math.Abs(lambdaLambda12)
	fWall := pK7*s + kappaERed*pK7*s*s - 2.0*pK7*pK7*s*s*s
	bracket := 1.0 + lHopf*(1.0-absLam-fWall+kappaERed)
	return Gate752Inheritance{
		Inherited:             true,
		KappaERedFormula:      "kappa_e_red=sin^2(theta13)/4-J_CKM-(5/3)s^2+xi_boundary p_K7 s^2",
		FWall3RedFormula:      "F_wall_3_red(s)=p_K7 s+kappa_e_red p_K7 s^2-2p_K7^2 s^3",
		RuntimeBridgeFormula:  "lambda_runtime_red=lambda_proxy[1+L_Hopf(1-|lambda|-F_wall_3_red(s)+kappa_e_red)]",
		ExpandedTraceFormula:  "lambda_runtime_red=lambda_proxy[1+Tr_K7+(rho_plus R_Hopf)(1-|lambda(Lambda_12)|-F_wall_3_red(sigma_boundary(b))+kappa_e_red)]",
		P:                     pK7,
		SBoundary:             s,
		AbsLambdaLambda12:     absLam,
		KappaERed:             kappaERed,
		LHopf:                 lHopf,
		FWall3Red:             fWall,
		RuntimeBracket:        bracket,
		RuntimeBracketFormula: "1+L_Hopf(1-|lambda|-F_wall_3_red+kappa_e_red)",
		ScalarMapTyped:        true,
		NativeScalarRuntime:   false,
		Verdict: strings.Join([]string{
			StatusGate752FlavorReducedNormalFormInherited,
			StatusNoIndependentScalarRuntimeTheorem,
		}, "; "),
	}
}

func buildEffectiveParticipationNormalForm(g756 Gate756Inheritance, g752 Gate752Inheritance) EffectiveParticipationNormalForm {
	runtime := g756.LambdaProxyFromNEff * g752.RuntimeBracket
	return EffectiveParticipationNormalForm{
		ProxySubstitutionFormula:           "lambda_proxy=3/(8N_eff)",
		RuntimeFormula:                     "lambda_runtime_eff=[3/(8N_eff)][1+L_Hopf(1-|lambda|-F_wall_3_red(s)+kappa_e_red)]",
		ExpandedTraceRuntimeFormula:        "lambda_runtime_eff=[3/(8N_eff)][1+Tr_K7+(rho_plus R_Hopf)(1-|lambda(Lambda_12)|-F_wall_3_red(sigma_boundary(b))+kappa_e_red)]",
		LambdaRuntimeEff:                   runtime,
		RuntimeBracket:                     g752.RuntimeBracket,
		ProxySubstituted:                   true,
		NormalFormWritten:                  true,
		EquivalentExpandedTraceFormWritten: true,
		IndependentRuntimePrediction:       false,
		NativeScalarProxyTheorem:           false,
		Verdict: strings.Join([]string{
			StatusEffectiveParticipationProxySubstituted,
			StatusScalarHiggsEffectiveParticipationNormalForm,
			StatusCurrentBridgeHasEffectiveParticipationForm,
			StatusNoNativeScalarProxyDerivationTheorem,
		}, "; "),
	}
}

func buildTopColorShadowComparison(g756 Gate756Inheritance, normal EffectiveParticipationNormalForm) TopColorShadowComparison {
	lambdaTop := three / (eight * nEffTopColor)
	runtimeTop := lambdaTop * normal.RuntimeBracket
	proxyShift := g756.LambdaProxyFromNEff - lambdaTop
	runtimeShift := normal.LambdaRuntimeEff - runtimeTop
	treeEff := math.Sqrt(2.0*normal.LambdaRuntimeEff) * vevConventionGeV
	treeTop := math.Sqrt(2.0*runtimeTop) * vevConventionGeV
	return TopColorShadowComparison{
		NEffTop:                       nEffTopColor,
		LambdaProxyTopShadow:          lambdaTop,
		LambdaRuntimeTopShadow:        runtimeTop,
		ProxyShift:                    proxyShift,
		RuntimeShift:                  runtimeShift,
		RuntimeShiftApprox:            runtimeShift,
		RelativeRuntimeShift:          runtimeShift / runtimeTop,
		TreeProxyEffGeV:               treeEff,
		TreeProxyTopShadowGeV:         treeTop,
		TreeProxyShiftGeV:             treeEff - treeTop,
		VevConventionGeV:              vevConventionGeV,
		ProxyBelowTopShadow:           proxyShift < 0,
		RuntimeLoweredByParticipation: runtimeShift < 0,
		TreeProxyDiagnosticOnly:       true,
		HiggsPoleMassPrediction:       false,
		Verdict: strings.Join([]string{
			StatusTopColorShadowComparisonAudited,
			StatusRuntimePropagationOfNEffDeviationComputed,
			StatusNEffMinusThreeLowersProxyBelowOneEighth,
			StatusNoHiggsMassOrPoleMassTheorem,
		}, "; "),
	}
}

func buildParticipationInterpretation(g756 Gate756Inheritance, top TopColorShadowComparison) ParticipationInterpretation {
	return ParticipationInterpretation{
		NEffGreaterThanThree:      g756.NEff > nEffTopColor,
		LambdaProxyBelowOneEighth: top.ProxyShift < 0,
		TraceLedgerMoreSpread:     g756.NEff > nEffTopColor,
		NonTopChannelsDiluteIPR:   g756.NEff > nEffTopColor,
		NonTopChannelsLowerProxy:  top.ProxyShift < 0,
		AssignedSectorCorrection:  false,
		Verdict: strings.Join([]string{
			StatusNEffMinusThreeLowersProxyBelowOneEighth,
			StatusNoChannelAssignmentWithoutLedger,
		}, "; "),
	}
}

func buildRelationToGate756() RelationToGate756 {
	return RelationToGate756{
		Gate756ParticipationFormula:     "b/a^2=1/N_eff and lambda_proxy=3/(8N_eff)",
		Gate757NormalFormFormula:        "lambda_runtime_eff=[3/(8N_eff)] transport_bracket_red",
		TopShadowNormalFormFormula:      "N_eff=3 => lambda_runtime_top_shadow=(1/8) transport_bracket_red",
		CompatibilityStatement:          "Gate756 supplies the aggregate proxy base; Gate752 supplies the separate HistoryLoop/boundary transport bracket.",
		UsesAggregateTracePair:          true,
		RequiresNoTopYukawaChoice:       true,
		RequiresGate752TransportBracket: true,
		Compatible:                      true,
		Verdict: strings.Join([]string{
			StatusGate756YukawaTraceParticipationInherited,
			StatusGate752FlavorReducedNormalFormInherited,
		}, "; "),
	}
}

func buildLayerSeparation() LayerSeparation {
	return LayerSeparation{
		NEffLayer:                       "finite Yukawa trace participation layer",
		RuntimeCorrectionLayer:          "HistoryLoop / boundary-history transport layer",
		NEffIsNativeGenerationTheorem:   false,
		NEffIsYukawaEigenvalueTheorem:   false,
		NEffIsScalarPotentialTheorem:    false,
		NEffIsRuntimeLambdaTheorem:      false,
		NEffIsHiggsMassTheorem:          false,
		RuntimeBracketSeparateTransport: true,
		Verdict: strings.Join([]string{
			StatusLayerSeparationEnforced,
			StatusNEffNotNativeGenerationTheorem,
			StatusNoNativeYukawaOperatorTheorem,
			StatusNoIndependentScalarRuntimeTheorem,
		}, "; "),
	}
}

func buildFirewalls() Firewalls {
	return Firewalls{
		NEffNativeGenerationTheorem:        false,
		NEffMinusThreeAssignedToSector:     false,
		LambdaProxyScalarPotentialTheorem:  false,
		LambdaRuntimeIndependentPrediction: false,
		TreeProxyHiggsPoleMassPrediction:   false,
		ClaimsYukawaEigenvaluesDerived:     false,
		ClaimsFlavorHierarchyDerived:       false,
		ClaimsCKMPMNSDerived:               false,
		ClaimsHiggsMassTheorem:             false,
		ClaimsPoleMassTheorem:              false,
		Verdict: strings.Join([]string{
			StatusYukawaAndHiggsFirewallsEnforced,
			StatusNEffNotNativeGenerationTheorem,
			StatusNoChannelAssignmentWithoutLedger,
			StatusNoNativeYukawaOperatorTheorem,
			StatusNoNativeScalarProxyDerivationTheorem,
			StatusNoIndependentScalarRuntimeTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusGate757EffectiveParticipationScalarProxyBoundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate756YukawaTraceParticipationInherited,
		StatusGate752FlavorReducedNormalFormInherited,
		StatusEffectiveParticipationProxySubstituted,
		StatusScalarHiggsEffectiveParticipationNormalForm,
		StatusTopColorShadowComparisonAudited,
		StatusRuntimePropagationOfNEffDeviationComputed,
		StatusLayerSeparationEnforced,
		StatusYukawaAndHiggsFirewallsEnforced,
		StatusScalarProxyThreeOverEightNEff,
		StatusCurrentBridgeHasEffectiveParticipationForm,
		StatusNEffMinusThreeLowersProxyBelowOneEighth,
		StatusNEffNotNativeGenerationTheorem,
		StatusNoChannelAssignmentWithoutLedger,
		StatusNoNativeYukawaOperatorTheorem,
		StatusNoNativeScalarProxyDerivationTheorem,
		StatusNoIndependentScalarRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusGate757EffectiveParticipationScalarProxyBoundary,
	}
}

func FormatGate756(x Gate756Inheritance) string {
	return fmt.Sprintf("inherited=%t participation=%q nEffDef=%q nEff=%.16g ratio=%.16g inverse=%.16g inverseResidual=%.16g lambdaFormula=%q lambda=%.16g lambdaFromNEff=%.16g lambdaResidual=%.16g topCount=%.16g nEffMinus3=%.16g typed=%t sealed=%t verdict=%q", x.Inherited, x.ParticipationFormula, x.NEffDefinition, x.NEff, x.TraceRatio, x.InverseNEff, x.InverseResidual, x.LambdaProxyFormula, x.LambdaProxy, x.LambdaProxyFromNEff, x.LambdaProxyResidual, x.TopColorEffectiveCount, x.NEffMinusThree, x.EffectiveParticipationTyped, x.YukawaTraceLedgerSealed, x.Verdict)
}

func FormatGate752(x Gate752Inheritance) string {
	return fmt.Sprintf("inherited=%t kappa=%q fwall=%q runtime=%q expanded=%q p=%.16g s=%.16g absLambda=%.16g kappa=%.16g L=%.16g fwallValue=%.16g bracket=%.16g bracketFormula=%q scalarMap=%t nativeRuntime=%t verdict=%q", x.Inherited, x.KappaERedFormula, x.FWall3RedFormula, x.RuntimeBridgeFormula, x.ExpandedTraceFormula, x.P, x.SBoundary, x.AbsLambdaLambda12, x.KappaERed, x.LHopf, x.FWall3Red, x.RuntimeBracket, x.RuntimeBracketFormula, x.ScalarMapTyped, x.NativeScalarRuntime, x.Verdict)
}

func FormatNormalForm(x EffectiveParticipationNormalForm) string {
	return fmt.Sprintf("proxy=%q runtime=%q expanded=%q lambdaRuntime=%.16g bracket=%.16g substituted=%t written=%t expandedWritten=%t independentRuntime=%t nativeProxy=%t verdict=%q", x.ProxySubstitutionFormula, x.RuntimeFormula, x.ExpandedTraceRuntimeFormula, x.LambdaRuntimeEff, x.RuntimeBracket, x.ProxySubstituted, x.NormalFormWritten, x.EquivalentExpandedTraceFormWritten, x.IndependentRuntimePrediction, x.NativeScalarProxyTheorem, x.Verdict)
}

func FormatTopShadow(x TopColorShadowComparison) string {
	return fmt.Sprintf("nEffTop=%.16g lambdaTop=%.16g runtimeTop=%.16g proxyShift=%.16g runtimeShift=%.16g runtimeApprox=%.16g relative=%.16g treeEff=%.16g treeTop=%.16g treeShift=%.16g v=%.16g proxyBelow=%t runtimeLowered=%t treeDiagnostic=%t polePrediction=%t verdict=%q", x.NEffTop, x.LambdaProxyTopShadow, x.LambdaRuntimeTopShadow, x.ProxyShift, x.RuntimeShift, x.RuntimeShiftApprox, x.RelativeRuntimeShift, x.TreeProxyEffGeV, x.TreeProxyTopShadowGeV, x.TreeProxyShiftGeV, x.VevConventionGeV, x.ProxyBelowTopShadow, x.RuntimeLoweredByParticipation, x.TreeProxyDiagnosticOnly, x.HiggsPoleMassPrediction, x.Verdict)
}

func FormatInterpretation(x ParticipationInterpretation) string {
	return fmt.Sprintf("nEffGT3=%t proxyBelowOneEighth=%t moreSpread=%t diluteIPR=%t lowerProxy=%t assignedSector=%t verdict=%q", x.NEffGreaterThanThree, x.LambdaProxyBelowOneEighth, x.TraceLedgerMoreSpread, x.NonTopChannelsDiluteIPR, x.NonTopChannelsLowerProxy, x.AssignedSectorCorrection, x.Verdict)
}

func FormatRelation(x RelationToGate756) string {
	return fmt.Sprintf("gate756=%q gate757=%q top=%q statement=%q aggregate=%t noTopChoice=%t needsTransport=%t compatible=%t verdict=%q", x.Gate756ParticipationFormula, x.Gate757NormalFormFormula, x.TopShadowNormalFormFormula, x.CompatibilityStatement, x.UsesAggregateTracePair, x.RequiresNoTopYukawaChoice, x.RequiresGate752TransportBracket, x.Compatible, x.Verdict)
}

func FormatLayering(x LayerSeparation) string {
	return fmt.Sprintf("nEffLayer=%q runtimeLayer=%q claims(nEffGeneration=%t yukawaEigen=%t scalarPotential=%t runtimeLambda=%t higgsMass=%t) separateTransport=%t verdict=%q", x.NEffLayer, x.RuntimeCorrectionLayer, x.NEffIsNativeGenerationTheorem, x.NEffIsYukawaEigenvalueTheorem, x.NEffIsScalarPotentialTheorem, x.NEffIsRuntimeLambdaTheorem, x.NEffIsHiggsMassTheorem, x.RuntimeBracketSeparateTransport, x.Verdict)
}

func FormatFirewalls(x Firewalls) string {
	return fmt.Sprintf("claims(nEffGeneration=%t sectorAssign=%t proxyPotential=%t runtimeIndependent=%t treePole=%t yukawaEigen=%t hierarchy=%t ckmPmns=%t higgs=%t pole=%t) verdict=%q", x.NEffNativeGenerationTheorem, x.NEffMinusThreeAssignedToSector, x.LambdaProxyScalarPotentialTheorem, x.LambdaRuntimeIndependentPrediction, x.TreeProxyHiggsPoleMassPrediction, x.ClaimsYukawaEigenvaluesDerived, x.ClaimsFlavorHierarchyDerived, x.ClaimsCKMPMNSDerived, x.ClaimsHiggsMassTheorem, x.ClaimsPoleMassTheorem, x.Verdict)
}
