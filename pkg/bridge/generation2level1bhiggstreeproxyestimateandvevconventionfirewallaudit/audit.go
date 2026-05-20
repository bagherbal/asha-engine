// Package generation2level1bhiggstreeproxyestimateandvevconventionfirewallaudit implements
// Gate 741: Level-1B Higgs Tree Proxy Estimate and VEV-Convention Firewall Audit.
//
// Gate 740 allowed a tree-level Higgs proxy estimate only with explicit VEV and
// scalar-potential convention seals. Gate 741 performs that Level-1B numerical
// proxy using the Gate 739 runtime quartic bridge value, while preserving the
// firewall that this is not a Higgs pole-mass theorem or an independent physical
// prediction.
package generation2level1bhiggstreeproxyestimateandvevconventionfirewallaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate740 "github.com/bagherbal/asha-engine/pkg/bridge/generation2runtimequartictohiggsmasstranslationfirewallandrequiredinputsaudit"
)

const (
	AuditID = "GATE741-LEVEL1B-HIGGS-TREE-PROXY-ESTIMATE-VEV-CONVENTION-FIREWALL-AUDIT"

	StatusGate740HiggsTranslationFirewallInherited = "PASS_GATE740_HIGGS_TRANSLATION_FIREWALL_INHERITED"
	StatusRuntimeQuarticBridgeValueInherited       = "PASS_RUNTIME_QUARTIC_BRIDGE_VALUE_INHERITED"
	StatusVEVConventionSealDefined                 = "PASS_VEV_CONVENTION_SEAL_DEFINED"
	StatusTreeProxyRelationApplied                 = "PASS_TREE_PROXY_RELATION_APPLIED"
	StatusLevel1BTreeProxyEstimateComputed         = "PASS_LEVEL_1B_TREE_PROXY_ESTIMATE_COMPUTED"
	StatusSensitivityToVEVAndLambdaRecorded        = "PASS_SENSITIVITY_TO_VEV_AND_LAMBDA_RECORDED"
	StatusSealDependenceCarriedForward             = "PASS_SEAL_DEPENDENCE_CARRIED_FORWARD"
	StatusPoleMassFirewallEnforced                 = "PASS_POLE_MASS_FIREWALL_ENFORCED"

	StatusLevel1BTreeProxyEstimateAllowedWithExplicitSeals = "CONDITIONAL_SUPPORT_LEVEL_1B_TREE_PROXY_ESTIMATE_IS_ALLOWED_WITH_EXPLICIT_SEALS"
	StatusTreeProxyValueComputableUnderVEVConvention       = "CONDITIONAL_SUPPORT_TREE_PROXY_VALUE_IS_NUMERICALLY_COMPUTABLE_UNDER_VEV_CONVENTION"

	StatusVEVNotNativelyDerived                = "FAILED_ROUTE_VEV_NOT_NATIVELY_DERIVED"
	StatusRuntimeLambdaNotIndependentlyDerived = "FAILED_ROUTE_RUNTIME_LAMBDA_NOT_INDEPENDENTLY_DERIVED"
	StatusTreeProxyNotHiggsPoleMass            = "FAILED_ROUTE_TREE_PROXY_NOT_HIGGS_POLE_MASS"
	StatusNoNativeRGThresholdPoleCorrection    = "FAILED_ROUTE_NO_NATIVE_RG_THRESHOLD_POLE_CORRECTION_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem         = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem  = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate741Boundary                      = "FIREWALL_PRESERVED_GATE741_LEVEL1B_HIGGS_TREE_PROXY_BOUNDARY"
)

const (
	DefaultVEVGeV = 246.2196508
)

type Gate740Inheritance struct {
	Inherited                   bool
	LambdaRuntimeBridge         float64
	RuntimeQuarticBridgeLayer   bool
	Level1BAllowed              bool
	RuntimeLambdaNotPoleMass    bool
	TreeProxyNotPoleMassTheorem bool
	NoNativeVEVTheorem          bool
	NoPoleCorrectionTheorem     bool
	NoHiggsMassTheorem          bool
	NoYukawaTheorem             bool
	Verdict                     string
}

type RuntimeBridgeValue struct {
	LambdaRuntimeBridge        float64
	ClassifiedAsRuntimeQuartic bool
	NotIndependentlyDerived    bool
	NotPoleMass                bool
	Verdict                    string
}

type VEVConventionSeal struct {
	Name             string
	VGeV             float64
	SuppliedInput    bool
	NativeDerivation bool
	Convention       bool
	Verdict          string
}

type TreeProxyEstimate struct {
	Relation            string
	LambdaRuntimeBridge float64
	VGeV                float64
	SqrtTwoLambdaFactor float64
	TreeProxyGeV        float64
	Level               string
	PoleMassPrediction  bool
	Verdict             string
}

type SensitivityRecord struct {
	Formula                              string
	LinearInV                            bool
	HalfPowerInLambda                    bool
	DeltaMOverMFromDeltaVOverV           float64
	DeltaMOverMFromDeltaLambdaOverLambda float64
	Verdict                              string
}

type SealCarryForward struct {
	Labels             []string
	IncludesVEVSeal    bool
	IncludesConvention bool
	Explicit           bool
	ProxyRemainsSealed bool
	Verdict            string
}

type PoleMassFirewall struct {
	TreeProxyEqualsPoleMass           bool
	RuntimeLambdaIndependentlyDerived bool
	HasPoleCorrectionTheorem          bool
	HasHiggsMassTheorem               bool
	Level2PredictionAllowed           bool
	Verdict                           string
}

type LevelClassification struct {
	Level1BName    string
	Level1BAllowed bool
	Level2Name     string
	Level2Allowed  bool
	ExplicitSeals  bool
	Verdict        string
}

type Analysis struct {
	Gate740     Gate740Inheritance
	Runtime     RuntimeBridgeValue
	VEV         VEVConventionSeal
	Proxy       TreeProxyEstimate
	Sensitivity SensitivityRecord
	Seals       SealCarryForward
	Firewall    PoleMassFirewall
	Level       LevelClassification
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
	g740, err := gate740.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate740 inheritance unavailable: %w", err)
	}
	inherit := buildGate740Inheritance(g740)
	runtime := buildRuntimeBridgeValue(inherit)
	vev := buildVEVConventionSeal(DefaultVEVGeV)
	proxy := buildTreeProxyEstimate(runtime, vev)
	sensitivity := buildSensitivityRecord()
	seals := buildSealCarryForward(g740)
	firewall := buildPoleMassFirewall(g740)
	level := buildLevelClassification()
	truth := "Gate 741 performs only the permitted Level-1B tree-level proxy estimate m_H_tree_proxy=sqrt(2 lambda_runtime_bridge) v using an explicit VEVConventionSeal. The result is a sealed proxy value, not a Higgs pole mass, not an independent runtime-lambda derivation, and not a physical Higgs-mass theorem."
	return Analysis{Gate740: inherit, Runtime: runtime, VEV: vev, Proxy: proxy, Sensitivity: sensitivity, Seals: seals, Firewall: firewall, Level: level, Truth: truth}, nil
}

func buildGate740Inheritance(g gate740.Analysis) Gate740Inheritance {
	statuses := strings.Join(gate740.Statuses(), "\n")
	return Gate740Inheritance{
		Inherited:                   g.Gate739.Inherited && g.Quartic.BridgeLayer && g.Proxy.RequiresV && g.Forecast.Level1BAllowed && !g.Forecast.Level2Allowed,
		LambdaRuntimeBridge:         g.Quartic.LambdaRuntimeBridge,
		RuntimeQuarticBridgeLayer:   g.Quartic.BridgeLayer && !g.Quartic.PhysicalPoleMass,
		Level1BAllowed:              g.Forecast.Level1BAllowed,
		RuntimeLambdaNotPoleMass:    strings.Contains(statuses, gate740.StatusRuntimeLambdaNotPoleMass),
		TreeProxyNotPoleMassTheorem: strings.Contains(statuses, gate740.StatusTreeProxyNotHiggsPoleMassTheorem),
		NoNativeVEVTheorem:          strings.Contains(statuses, gate740.StatusNoNativeVEVOrElectroweakScaleTheorem),
		NoPoleCorrectionTheorem:     strings.Contains(statuses, gate740.StatusNoNativeRGThresholdPoleCorrectionTheorem),
		NoHiggsMassTheorem:          strings.Contains(statuses, gate740.StatusNoHiggsMassOrPoleMassTheorem),
		NoYukawaTheorem:             strings.Contains(statuses, gate740.StatusNoYukawaOperatorOrEigenvalueTheorem),
		Verdict:                     StatusGate740HiggsTranslationFirewallInherited,
	}
}

func buildRuntimeBridgeValue(g Gate740Inheritance) RuntimeBridgeValue {
	return RuntimeBridgeValue{
		LambdaRuntimeBridge:        g.LambdaRuntimeBridge,
		ClassifiedAsRuntimeQuartic: true,
		NotIndependentlyDerived:    true,
		NotPoleMass:                true,
		Verdict:                    strings.Join([]string{StatusRuntimeQuarticBridgeValueInherited, StatusRuntimeLambdaNotIndependentlyDerived}, "; "),
	}
}

func buildVEVConventionSeal(v float64) VEVConventionSeal {
	return VEVConventionSeal{
		Name:             "VEVConventionSeal",
		VGeV:             v,
		SuppliedInput:    true,
		NativeDerivation: false,
		Convention:       true,
		Verdict:          strings.Join([]string{StatusVEVConventionSealDefined, StatusVEVNotNativelyDerived}, "; "),
	}
}

func buildTreeProxyEstimate(r RuntimeBridgeValue, v VEVConventionSeal) TreeProxyEstimate {
	factor := math.Sqrt(2 * r.LambdaRuntimeBridge)
	return TreeProxyEstimate{
		Relation:            "m_H_tree_proxy=sqrt(2*lambda_runtime_bridge)*v",
		LambdaRuntimeBridge: r.LambdaRuntimeBridge,
		VGeV:                v.VGeV,
		SqrtTwoLambdaFactor: factor,
		TreeProxyGeV:        factor * v.VGeV,
		Level:               "Level 1B sealed tree-level proxy estimate",
		PoleMassPrediction:  false,
		Verdict:             strings.Join([]string{StatusTreeProxyRelationApplied, StatusLevel1BTreeProxyEstimateComputed, StatusTreeProxyValueComputableUnderVEVConvention, StatusTreeProxyNotHiggsPoleMass}, "; "),
	}
}

func buildSensitivityRecord() SensitivityRecord {
	return SensitivityRecord{
		Formula:                              "delta m_H / m_H = delta v / v + 0.5 delta lambda / lambda",
		LinearInV:                            true,
		HalfPowerInLambda:                    true,
		DeltaMOverMFromDeltaVOverV:           1,
		DeltaMOverMFromDeltaLambdaOverLambda: 0.5,
		Verdict:                              StatusSensitivityToVEVAndLambdaRecorded,
	}
}

func buildSealCarryForward(g gate740.Analysis) SealCarryForward {
	labels := append([]string{}, g.Seals.Seals...)
	labels = append(labels, "VEVConventionSeal", "scalar-potential convention")
	return SealCarryForward{
		Labels:             labels,
		IncludesVEVSeal:    true,
		IncludesConvention: true,
		Explicit:           len(labels) == 12,
		ProxyRemainsSealed: true,
		Verdict:            StatusSealDependenceCarriedForward,
	}
}

func buildPoleMassFirewall(g gate740.Analysis) PoleMassFirewall {
	return PoleMassFirewall{
		TreeProxyEqualsPoleMass:           false,
		RuntimeLambdaIndependentlyDerived: false,
		HasPoleCorrectionTheorem:          false,
		HasHiggsMassTheorem:               false,
		Level2PredictionAllowed:           false,
		Verdict:                           strings.Join([]string{StatusPoleMassFirewallEnforced, StatusNoNativeRGThresholdPoleCorrection, StatusNoHiggsMassOrPoleMassTheorem, StatusNoYukawaOperatorOrEigenvalueTheorem}, "; "),
	}
}

func buildLevelClassification() LevelClassification {
	return LevelClassification{
		Level1BName:    "Level 1B sealed tree-level Higgs proxy estimate",
		Level1BAllowed: true,
		Level2Name:     "Level 2 physical Higgs pole-mass prediction",
		Level2Allowed:  false,
		ExplicitSeals:  true,
		Verdict:        strings.Join([]string{StatusLevel1BTreeProxyEstimateAllowedWithExplicitSeals, StatusTreeProxyValueComputableUnderVEVConvention}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate740HiggsTranslationFirewallInherited,
		StatusRuntimeQuarticBridgeValueInherited,
		StatusVEVConventionSealDefined,
		StatusTreeProxyRelationApplied,
		StatusLevel1BTreeProxyEstimateComputed,
		StatusSensitivityToVEVAndLambdaRecorded,
		StatusSealDependenceCarriedForward,
		StatusPoleMassFirewallEnforced,
		StatusLevel1BTreeProxyEstimateAllowedWithExplicitSeals,
		StatusTreeProxyValueComputableUnderVEVConvention,
		StatusVEVNotNativelyDerived,
		StatusRuntimeLambdaNotIndependentlyDerived,
		StatusTreeProxyNotHiggsPoleMass,
		StatusNoNativeRGThresholdPoleCorrection,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate741Boundary,
	}
}

func FormatGate740(x Gate740Inheritance) string {
	return fmt.Sprintf("inherited=%t lambdaRuntime=%.17g bridgeLayer=%t level1B=%t notPole=%t treeNotPole=%t noVEV=%t noPoleCorrection=%t noMass=%t noYukawa=%t verdict=%q", x.Inherited, x.LambdaRuntimeBridge, x.RuntimeQuarticBridgeLayer, x.Level1BAllowed, x.RuntimeLambdaNotPoleMass, x.TreeProxyNotPoleMassTheorem, x.NoNativeVEVTheorem, x.NoPoleCorrectionTheorem, x.NoHiggsMassTheorem, x.NoYukawaTheorem, x.Verdict)
}

func FormatRuntime(x RuntimeBridgeValue) string {
	return fmt.Sprintf("lambdaRuntime=%.17g runtimeQuartic=%t notIndependent=%t notPole=%t verdict=%q", x.LambdaRuntimeBridge, x.ClassifiedAsRuntimeQuartic, x.NotIndependentlyDerived, x.NotPoleMass, x.Verdict)
}

func FormatVEV(x VEVConventionSeal) string {
	return fmt.Sprintf("name=%s vGeV=%.10f supplied=%t native=%t convention=%t verdict=%q", x.Name, x.VGeV, x.SuppliedInput, x.NativeDerivation, x.Convention, x.Verdict)
}

func FormatProxy(x TreeProxyEstimate) string {
	return fmt.Sprintf("relation=%q lambda=%.17g v=%.10f sqrt2lambda=%.17g proxyGeV=%.17g level=%q polePrediction=%t verdict=%q", x.Relation, x.LambdaRuntimeBridge, x.VGeV, x.SqrtTwoLambdaFactor, x.TreeProxyGeV, x.Level, x.PoleMassPrediction, x.Verdict)
}

func FormatSensitivity(x SensitivityRecord) string {
	return fmt.Sprintf("formula=%q linearV=%t halfLambda=%t coeffV=%.1f coeffLambda=%.1f verdict=%q", x.Formula, x.LinearInV, x.HalfPowerInLambda, x.DeltaMOverMFromDeltaVOverV, x.DeltaMOverMFromDeltaLambdaOverLambda, x.Verdict)
}

func FormatSeals(x SealCarryForward) string {
	return fmt.Sprintf("labels=[%s] vev=%t convention=%t explicit=%t proxySealed=%t verdict=%q", strings.Join(x.Labels, ","), x.IncludesVEVSeal, x.IncludesConvention, x.Explicit, x.ProxyRemainsSealed, x.Verdict)
}

func FormatFirewall(x PoleMassFirewall) string {
	return fmt.Sprintf("treeEqualsPole=%t runtimeIndependent=%t poleCorrection=%t massTheorem=%t level2=%t verdict=%q", x.TreeProxyEqualsPoleMass, x.RuntimeLambdaIndependentlyDerived, x.HasPoleCorrectionTheorem, x.HasHiggsMassTheorem, x.Level2PredictionAllowed, x.Verdict)
}

func FormatLevel(x LevelClassification) string {
	return fmt.Sprintf("level1B=%q allowed=%t level2=%q allowed=%t explicitSeals=%t verdict=%q", x.Level1BName, x.Level1BAllowed, x.Level2Name, x.Level2Allowed, x.ExplicitSeals, x.Verdict)
}
