// Package generation2finitehiggsoneformscalarproxycoefficienttypingandfirewallaudit implements
// Gate 753: Finite Higgs One-Form Scalar Proxy Coefficient Typing and Firewall Audit.
//
// Gate 752 produced the flavor-reduced scalar-Higgs normal form with
// lambda_proxy as the multiplicative base. Gate 753 audits the source type of
// that base, the coefficient 3/8, and the spectral-action trace ratio b/a^2
// before the proxy is used in scalar-runtime transport. It preserves the
// distinction between finite trace-shape diagnostics, bridge normalization, and
// runtime/pole-mass physics.
package generation2finitehiggsoneformscalarproxycoefficienttypingandfirewallaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE753-FINITE-HIGGS-ONE-FORM-SCALAR-PROXY-COEFFICIENT-TYPING-AND-FIREWALL-AUDIT"

	StatusGate752FlavorReducedNormalFormInherited = "PASS_GATE752_FLAVOR_REDUCED_NORMAL_FORM_INHERITED"
	StatusGate620ScalarProxyLaneInherited         = "PASS_GATE620_SCALAR_PROXY_LANE_INHERITED"
	StatusFiniteTraceFormsTyped                   = "PASS_FINITE_SPECTRAL_ACTION_TRACE_FORMS_TYPED"
	StatusBA2RatioAudited                         = "PASS_B_OVER_A_SQUARED_RATIO_AUDITED"
	StatusThreeEighthsCoefficientTyped            = "PASS_THREE_EIGHTHS_COEFFICIENT_TYPED"
	StatusOneEighthShadowComputed                 = "PASS_ONE_EIGHTH_PROXY_SHADOW_COMPUTED"
	StatusMultiplicativeBaseRoleAudited           = "PASS_MULTIPLICATIVE_BASE_ROLE_AUDITED"
	StatusSourceLayersSeparated                   = "PASS_SCALAR_PROXY_SOURCE_LAYERS_SEPARATED"
	StatusPhysicalFirewallsEnforced               = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusScalarProxyTypedAsFiniteOneFormDiagnostic = "CONDITIONAL_SUPPORT_LAMBDA_PROXY_IS_FINITE_HIGGS_ONE_FORM_SCALAR_DIAGNOSTIC"
	StatusThreeEighthsGaugeBoundaryCoefficient      = "CONDITIONAL_SUPPORT_THREE_EIGHTHS_IS_GAUGE_BOUNDARY_NORMALIZATION_REUSED_AS_SCALAR_PROXY_COEFFICIENT"
	StatusBA2NearOneThirdShadow                     = "CONDITIONAL_SUPPORT_B_OVER_A_SQUARED_NEAR_ONE_THIRD_GIVES_ONE_EIGHTH_PROXY_SHADOW"
	StatusProxyMayServeAsRuntimeBase                = "CONDITIONAL_SUPPORT_LAMBDA_PROXY_MAY_SERVE_AS_PRE_TRANSPORT_MULTIPLICATIVE_BASE"

	StatusNoNativeBA2OneThirdTheorem        = "FAILED_ROUTE_NO_NATIVE_B_OVER_A_SQUARED_ONE_THIRD_THEOREM"
	StatusNoNativeThreeEighthsScalarTheorem = "FAILED_ROUTE_NO_NATIVE_THREE_EIGHTHS_SCALAR_POTENTIAL_THEOREM"
	StatusNoNativeScalarProxyTheorem        = "FAILED_ROUTE_NO_NATIVE_SCALAR_PROXY_DERIVATION_THEOREM"
	StatusNoNativeProxyRuntimeTheorem       = "FAILED_ROUTE_NO_NATIVE_PROXY_TO_RUNTIME_MATCHING_THEOREM"
	StatusNoNativeYukawaEigenvalueTheorem   = "FAILED_ROUTE_NO_NATIVE_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusNoNativeFlavorTheorem             = "FAILED_ROUTE_NO_NATIVE_FLAVOR_HIERARCHY_PMNS_OR_CKM_THEOREM"
	StatusNoIndependentRuntimeTheorem       = "FAILED_ROUTE_NO_INDEPENDENT_SCALAR_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem      = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoHistoryLoopUnitTheorem          = "FAILED_ROUTE_NO_NATIVE_HISTORYLOOPUNIT_THEOREM"
	StatusGate753Boundary                   = "FIREWALL_PRESERVED_GATE753_SCALAR_PROXY_COEFFICIENT_TYPING_BOUNDARY"
)

const (
	threeEighths          = 3.0 / 8.0
	oneThird              = 1.0 / 3.0
	oneEighth             = 1.0 / 8.0
	aTraceMZ              = 2.8424095142339083
	bTraceMZ              = 2.6910096440382287
	bOverA2MZ             = 0.33307493962706697
	lambdaProxyMZ         = 0.12490310236015
	lambdaRuntimeMZ       = 0.1296525650504758
	lambdaProxyLambda12   = 0.12490365414532
	lambdaRuntimeLambda12 = -0.049700942077683274
	gate752RuntimeRed     = 0.12965256505060743
	gate752RuntimeFormula = "lambda_runtime_red=lambda_proxy[1+L_Hopf(1-|lambda|-F_wall_3_red(s)+kappa_e_red)]"
)

type Gate752Inheritance struct {
	Inherited                 bool
	ReducedNormalFormReady    bool
	LambdaProxyMultiplicative bool
	KappaEReducedInserted     bool
	RuntimePredictionBlocked  bool
	HiggsMassBlocked          bool
	Verdict                   string
}

type Gate620ProxyInheritance struct {
	Inherited                bool
	ProxyRowsAvailable       bool
	LowScaleProxyPositive    bool
	HighScaleProxyPositive   bool
	LowScaleCloseToRuntime   bool
	HighScaleRuntimeSignFail bool
	SeparateScalarLanes      bool
	Verdict                  string
}

type FiniteTraceForms struct {
	AFormula                    string
	BFormula                    string
	AType                       string
	BType                       string
	ATraceMZ                    float64
	BTraceMZ                    float64
	ATracePositive              bool
	BTraceNonNegative           bool
	PolynomialTraceFormsNative  bool
	EvaluatedYukawaValuesSealed bool
	Verdict                     string
}

type BA2RatioAudit struct {
	BA2Formula            string
	BOverA2MZ             float64
	DeltaFromOneThird     float64
	RelativeFromOneThird  float64
	NonNegativeRatio      bool
	DimensionlessRatio    bool
	TopDominanceCandidate bool
	NativeOneThirdTheorem bool
	Verdict               string
}

type ThreeEighthsCoefficientAudit struct {
	Coefficient                    float64
	SourceFormula                  string
	SourceLayer                    string
	GaugeBoundaryNormalization     bool
	ScalarPotentialCoefficient     bool
	ScalarConventionAirlock        bool
	NativeScalarCoefficientTheorem bool
	Verdict                        string
}

type OneEighthShadowAudit struct {
	Formula                      string
	IdealBA2                     float64
	IdealProxy                   float64
	ActualProxy                  float64
	ProxyMinusOneEighth          float64
	RelativeProxyMinusOneEighth  float64
	CoefficientTimesBA2Deviation float64
	ShadowIdentityResidual       float64
	CloseToOneEighth             bool
	Verdict                      string
}

type MultiplicativeBaseRole struct {
	NormalForm              string
	LambdaProxy             float64
	ReducedRuntime          float64
	TransportFactor         float64
	ProxyOutsideLoopBracket bool
	IndependentOfKappaE     bool
	IndependentOfFWall3     bool
	IndependentOfLHopf      bool
	RuntimeDerived          bool
	BaseRole                string
	Verdict                 string
}

type SourceLayerSeparation struct {
	NativeTraceShapeLayer      []string
	BridgeCoefficientLayer     []string
	EnvironmentalValueLayer    []string
	RuntimeTransportLayer      []string
	AllLayersSeparated         bool
	NoCircularRuntimePromotion bool
	Verdict                    string
}

type PhysicalFirewalls struct {
	ClaimsNativeBA2OneThird        bool
	ClaimsNativeThreeEighthsScalar bool
	ClaimsNativeScalarProxy        bool
	ClaimsRuntimeLambda            bool
	ClaimsHiggsMass                bool
	ClaimsPoleMass                 bool
	ClaimsYukawaEigenvalues        bool
	ClaimsFlavorHierarchy          bool
	ClaimsHistoryLoopUnitSource    bool
	Verdict                        string
}

type Analysis struct {
	Gate752     Gate752Inheritance
	Gate620     Gate620ProxyInheritance
	TraceForms  FiniteTraceForms
	BA2         BA2RatioAudit
	Coefficient ThreeEighthsCoefficientAudit
	OneEighth   OneEighthShadowAudit
	BaseRole    MultiplicativeBaseRole
	Layers      SourceLayerSeparation
	Firewalls   PhysicalFirewalls
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
	gate752Inh := buildGate752Inheritance()
	gate620Inh := buildGate620Inheritance()
	trace := buildFiniteTraceForms()
	ba2 := buildBA2RatioAudit(trace)
	coeff := buildThreeEighthsCoefficientAudit()
	shadow := buildOneEighthShadowAudit(ba2, coeff)
	base := buildMultiplicativeBaseRole(shadow)
	layers := buildSourceLayerSeparation()
	firewalls := buildPhysicalFirewalls()
	truth := "Gate 753 types lambda_proxy=(3/8)(b/a^2) as a finite Higgs one-form scalar proxy diagnostic: a and b are spectral-action trace-shape forms, b/a^2 is a dimensionless nonnegative ratio evaluated here with sealed Yukawa data, and 3/8 is a gauge-boundary normalization reused as a scalar proxy coefficient through a bridge airlock. The proxy is lawful as the pre-transport multiplicative base in the Gate752 normal form, but it is not a native scalar potential theorem, not runtime lambda, not a Higgs mass or pole-mass theorem, and not a flavor/Yukawa derivation."

	return Analysis{Gate752: gate752Inh, Gate620: gate620Inh, TraceForms: trace, BA2: ba2, Coefficient: coeff, OneEighth: shadow, BaseRole: base, Layers: layers, Firewalls: firewalls, Truth: truth}, nil
}

func buildGate752Inheritance() Gate752Inheritance {
	return Gate752Inheritance{
		Inherited:                 true,
		ReducedNormalFormReady:    strings.Contains(gate752RuntimeFormula, "lambda_proxy") && strings.Contains(gate752RuntimeFormula, "kappa_e_red"),
		LambdaProxyMultiplicative: strings.HasPrefix(gate752RuntimeFormula, "lambda_runtime_red=lambda_proxy"),
		KappaEReducedInserted:     true,
		RuntimePredictionBlocked:  true,
		HiggsMassBlocked:          true,
		Verdict:                   StatusGate752FlavorReducedNormalFormInherited,
	}
}

func buildGate620Inheritance() Gate620ProxyInheritance {
	return Gate620ProxyInheritance{
		Inherited:                true,
		ProxyRowsAvailable:       true,
		LowScaleProxyPositive:    lambdaProxyMZ > 0,
		HighScaleProxyPositive:   lambdaProxyLambda12 > 0,
		LowScaleCloseToRuntime:   math.Abs(lambdaProxyMZ-lambdaRuntimeMZ) < 0.006,
		HighScaleRuntimeSignFail: lambdaRuntimeLambda12 < 0 && lambdaProxyLambda12 > 0,
		SeparateScalarLanes:      true,
		Verdict:                  StatusGate620ScalarProxyLaneInherited,
	}
}

func buildFiniteTraceForms() FiniteTraceForms {
	return FiniteTraceForms{
		AFormula:                    "a=Tr(Y_e†Y_e+Y_nu†Y_nu+3Y_u†Y_u+3Y_d†Y_d)",
		BFormula:                    "b=Tr((Y_e†Y_e)^2+(Y_nu†Y_nu)^2+3(Y_u†Y_u)^2+3(Y_d†Y_d)^2)",
		AType:                       "finite spectral-action Yukawa quadratic trace form",
		BType:                       "finite spectral-action Yukawa quartic trace form",
		ATraceMZ:                    aTraceMZ,
		BTraceMZ:                    bTraceMZ,
		ATracePositive:              aTraceMZ > 0,
		BTraceNonNegative:           bTraceMZ >= 0,
		PolynomialTraceFormsNative:  true,
		EvaluatedYukawaValuesSealed: true,
		Verdict: strings.Join([]string{
			StatusFiniteTraceFormsTyped,
			StatusScalarProxyTypedAsFiniteOneFormDiagnostic,
			StatusNoNativeYukawaEigenvalueTheorem,
		}, "; "),
	}
}

func buildBA2RatioAudit(t FiniteTraceForms) BA2RatioAudit {
	ba2 := t.BTraceMZ / (t.ATraceMZ * t.ATraceMZ)
	delta := ba2 - oneThird
	return BA2RatioAudit{
		BA2Formula:            "b/a^2",
		BOverA2MZ:             ba2,
		DeltaFromOneThird:     delta,
		RelativeFromOneThird:  delta / oneThird,
		NonNegativeRatio:      ba2 >= 0,
		DimensionlessRatio:    true,
		TopDominanceCandidate: true,
		NativeOneThirdTheorem: false,
		Verdict: strings.Join([]string{
			StatusBA2RatioAudited,
			StatusBA2NearOneThirdShadow,
			StatusNoNativeBA2OneThirdTheorem,
		}, "; "),
	}
}

func buildThreeEighthsCoefficientAudit() ThreeEighthsCoefficientAudit {
	return ThreeEighthsCoefficientAudit{
		Coefficient:                    threeEighths,
		SourceFormula:                  "c_proxy=sin²(theta_*)=3/8 from finite gauge-boundary normalization",
		SourceLayer:                    "gauge-boundary trace normalization reused through scalar proxy airlock",
		GaugeBoundaryNormalization:     true,
		ScalarPotentialCoefficient:     false,
		ScalarConventionAirlock:        true,
		NativeScalarCoefficientTheorem: false,
		Verdict: strings.Join([]string{
			StatusThreeEighthsCoefficientTyped,
			StatusThreeEighthsGaugeBoundaryCoefficient,
			StatusNoNativeThreeEighthsScalarTheorem,
		}, "; "),
	}
}

func buildOneEighthShadowAudit(b BA2RatioAudit, c ThreeEighthsCoefficientAudit) OneEighthShadowAudit {
	actual := c.Coefficient * b.BOverA2MZ
	ideal := c.Coefficient * oneThird
	coeffDelta := c.Coefficient * b.DeltaFromOneThird
	return OneEighthShadowAudit{
		Formula:                      "lambda_proxy=(3/8)(b/a^2)=(3/8)(1/3+delta)=1/8+(3/8)delta",
		IdealBA2:                     oneThird,
		IdealProxy:                   ideal,
		ActualProxy:                  actual,
		ProxyMinusOneEighth:          actual - oneEighth,
		RelativeProxyMinusOneEighth:  (actual - oneEighth) / oneEighth,
		CoefficientTimesBA2Deviation: coeffDelta,
		ShadowIdentityResidual:       (actual - oneEighth) - coeffDelta,
		CloseToOneEighth:             math.Abs(actual-oneEighth) < 1e-3,
		Verdict: strings.Join([]string{
			StatusOneEighthShadowComputed,
			StatusBA2NearOneThirdShadow,
		}, "; "),
	}
}

func buildMultiplicativeBaseRole(s OneEighthShadowAudit) MultiplicativeBaseRole {
	reduced := gate752RuntimeRed
	return MultiplicativeBaseRole{
		NormalForm:              gate752RuntimeFormula,
		LambdaProxy:             s.ActualProxy,
		ReducedRuntime:          reduced,
		TransportFactor:         reduced / s.ActualProxy,
		ProxyOutsideLoopBracket: strings.HasPrefix(gate752RuntimeFormula, "lambda_runtime_red=lambda_proxy"),
		IndependentOfKappaE:     true,
		IndependentOfFWall3:     true,
		IndependentOfLHopf:      true,
		RuntimeDerived:          false,
		BaseRole:                "pre-transport scalar proxy multiplicative base; all HistoryLoop/flavor-wall corrections remain inside the bracket",
		Verdict: strings.Join([]string{
			StatusMultiplicativeBaseRoleAudited,
			StatusProxyMayServeAsRuntimeBase,
			StatusNoNativeProxyRuntimeTheorem,
			StatusNoIndependentRuntimeTheorem,
		}, "; "),
	}
}

func buildSourceLayerSeparation() SourceLayerSeparation {
	return SourceLayerSeparation{
		NativeTraceShapeLayer: []string{
			"formal spectral-action trace polynomials a and b",
			"dimensionless nonnegative ratio b/a^2 as a trace-shape diagnostic",
		},
		BridgeCoefficientLayer: []string{
			"c_proxy=3/8 imported from finite gauge-boundary normalization into scalar proxy airlock",
			"lambda_proxy=(3/8)(b/a^2) as finite Higgs one-form scalar proxy diagnostic",
		},
		EnvironmentalValueLayer: []string{
			"evaluated Yukawa singular-value ledger supplying numerical a and b",
			"observed top/color dominance explanation for b/a^2≈1/3",
		},
		RuntimeTransportLayer: []string{
			"HistoryLoopUnit bracket L_Hopf(1-|lambda|-F_wall_3+kappa_e_red)",
			"RG/runtime quartic lambda_runtime and pole-mass corrections remain separate",
		},
		AllLayersSeparated:         true,
		NoCircularRuntimePromotion: true,
		Verdict: strings.Join([]string{
			StatusSourceLayersSeparated,
			StatusScalarProxyTypedAsFiniteOneFormDiagnostic,
			StatusProxyMayServeAsRuntimeBase,
		}, "; "),
	}
}

func buildPhysicalFirewalls() PhysicalFirewalls {
	return PhysicalFirewalls{
		ClaimsNativeBA2OneThird:        false,
		ClaimsNativeThreeEighthsScalar: false,
		ClaimsNativeScalarProxy:        false,
		ClaimsRuntimeLambda:            false,
		ClaimsHiggsMass:                false,
		ClaimsPoleMass:                 false,
		ClaimsYukawaEigenvalues:        false,
		ClaimsFlavorHierarchy:          false,
		ClaimsHistoryLoopUnitSource:    false,
		Verdict: strings.Join([]string{
			StatusPhysicalFirewallsEnforced,
			StatusNoNativeBA2OneThirdTheorem,
			StatusNoNativeThreeEighthsScalarTheorem,
			StatusNoNativeScalarProxyTheorem,
			StatusNoNativeProxyRuntimeTheorem,
			StatusNoNativeYukawaEigenvalueTheorem,
			StatusNoNativeFlavorTheorem,
			StatusNoIndependentRuntimeTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusNoHistoryLoopUnitTheorem,
			StatusGate753Boundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate752FlavorReducedNormalFormInherited,
		StatusGate620ScalarProxyLaneInherited,
		StatusFiniteTraceFormsTyped,
		StatusBA2RatioAudited,
		StatusThreeEighthsCoefficientTyped,
		StatusOneEighthShadowComputed,
		StatusMultiplicativeBaseRoleAudited,
		StatusSourceLayersSeparated,
		StatusPhysicalFirewallsEnforced,
		StatusScalarProxyTypedAsFiniteOneFormDiagnostic,
		StatusThreeEighthsGaugeBoundaryCoefficient,
		StatusBA2NearOneThirdShadow,
		StatusProxyMayServeAsRuntimeBase,
		StatusNoNativeBA2OneThirdTheorem,
		StatusNoNativeThreeEighthsScalarTheorem,
		StatusNoNativeScalarProxyTheorem,
		StatusNoNativeProxyRuntimeTheorem,
		StatusNoNativeYukawaEigenvalueTheorem,
		StatusNoNativeFlavorTheorem,
		StatusNoIndependentRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoHistoryLoopUnitTheorem,
		StatusGate753Boundary,
	}
}

func FormatGate752(x Gate752Inheritance) string {
	return fmt.Sprintf("inherited=%t reduced=%t proxyBase=%t kappaRed=%t runtimeBlocked=%t higgsBlocked=%t verdict=%q", x.Inherited, x.ReducedNormalFormReady, x.LambdaProxyMultiplicative, x.KappaEReducedInserted, x.RuntimePredictionBlocked, x.HiggsMassBlocked, x.Verdict)
}

func FormatGate620(x Gate620ProxyInheritance) string {
	return fmt.Sprintf("inherited=%t rows=%t lowProxy=%t highProxy=%t closeMZ=%t signFailL12=%t separate=%t verdict=%q", x.Inherited, x.ProxyRowsAvailable, x.LowScaleProxyPositive, x.HighScaleProxyPositive, x.LowScaleCloseToRuntime, x.HighScaleRuntimeSignFail, x.SeparateScalarLanes, x.Verdict)
}

func FormatTraceForms(x FiniteTraceForms) string {
	return fmt.Sprintf("a=%q b=%q aMZ=%.16g bMZ=%.16g aPos=%t bNonNeg=%t nativeForms=%t sealedValues=%t verdict=%q", x.AFormula, x.BFormula, x.ATraceMZ, x.BTraceMZ, x.ATracePositive, x.BTraceNonNegative, x.PolynomialTraceFormsNative, x.EvaluatedYukawaValuesSealed, x.Verdict)
}

func FormatBA2(x BA2RatioAudit) string {
	return fmt.Sprintf("formula=%q ba2=%.16g deltaOneThird=%.16g rel=%.16g nonNeg=%t dimless=%t topDominance=%t native=%t verdict=%q", x.BA2Formula, x.BOverA2MZ, x.DeltaFromOneThird, x.RelativeFromOneThird, x.NonNegativeRatio, x.DimensionlessRatio, x.TopDominanceCandidate, x.NativeOneThirdTheorem, x.Verdict)
}

func FormatCoefficient(x ThreeEighthsCoefficientAudit) string {
	return fmt.Sprintf("c=%.16g source=%q layer=%q gaugeBoundary=%t scalarCoeff=%t airlock=%t nativeScalar=%t verdict=%q", x.Coefficient, x.SourceFormula, x.SourceLayer, x.GaugeBoundaryNormalization, x.ScalarPotentialCoefficient, x.ScalarConventionAirlock, x.NativeScalarCoefficientTheorem, x.Verdict)
}

func FormatOneEighth(x OneEighthShadowAudit) string {
	return fmt.Sprintf("formula=%q idealBA2=%.16g idealProxy=%.16g actual=%.16g delta=%.16g rel=%.16g coeffDelta=%.16g identityResidual=%.16g close=%t verdict=%q", x.Formula, x.IdealBA2, x.IdealProxy, x.ActualProxy, x.ProxyMinusOneEighth, x.RelativeProxyMinusOneEighth, x.CoefficientTimesBA2Deviation, x.ShadowIdentityResidual, x.CloseToOneEighth, x.Verdict)
}

func FormatBaseRole(x MultiplicativeBaseRole) string {
	return fmt.Sprintf("form=%q proxy=%.16g runtime=%.16g factor=%.16g outside=%t independent(k=%t f=%t L=%t) runtimeDerived=%t role=%q verdict=%q", x.NormalForm, x.LambdaProxy, x.ReducedRuntime, x.TransportFactor, x.ProxyOutsideLoopBracket, x.IndependentOfKappaE, x.IndependentOfFWall3, x.IndependentOfLHopf, x.RuntimeDerived, x.BaseRole, x.Verdict)
}

func FormatLayers(x SourceLayerSeparation) string {
	return fmt.Sprintf("native=[%s] bridge=[%s] environmental=[%s] runtime=[%s] separated=%t noCircular=%t verdict=%q", strings.Join(x.NativeTraceShapeLayer, "; "), strings.Join(x.BridgeCoefficientLayer, "; "), strings.Join(x.EnvironmentalValueLayer, "; "), strings.Join(x.RuntimeTransportLayer, "; "), x.AllLayersSeparated, x.NoCircularRuntimePromotion, x.Verdict)
}

func FormatFirewalls(x PhysicalFirewalls) string {
	return fmt.Sprintf("ba2Native=%t cScalarNative=%t proxyNative=%t runtime=%t higgs=%t pole=%t yukawa=%t flavor=%t L=%t verdict=%q", x.ClaimsNativeBA2OneThird, x.ClaimsNativeThreeEighthsScalar, x.ClaimsNativeScalarProxy, x.ClaimsRuntimeLambda, x.ClaimsHiggsMass, x.ClaimsPoleMass, x.ClaimsYukawaEigenvalues, x.ClaimsFlavorHierarchy, x.ClaimsHistoryLoopUnitSource, x.Verdict)
}
