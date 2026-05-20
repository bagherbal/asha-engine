// Package generation2runtimequartictohiggsmasstranslationfirewallandrequiredinputsaudit implements
// Gate 740: Runtime Quartic to Higgs-Mass Translation Firewall and Required Inputs Audit.
//
// Gate 739 produced a Level-1 scalar-runtime bridge consistency estimate for
// lambda_runtime. Gate 740 audits what extra mathematical/physical inputs are
// required before a runtime quartic can be translated into a Higgs-mass
// statement. It records the tree-level proxy relation only as a convention-
// dependent candidate and preserves the firewall that no Higgs pole-mass theorem
// is certified.
package generation2runtimequartictohiggsmasstranslationfirewallandrequiredinputsaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate739 "github.com/bagherbal/asha-engine/pkg/bridge/generation2level1scalarruntimebridgeconsistencyestimateandnonpredictionaudit"
)

const (
	AuditID = "GATE740-RUNTIME-QUARTIC-TO-HIGGS-MASS-TRANSLATION-FIREWALL-REQUIRED-INPUTS-AUDIT"

	StatusGate739Level1ScalarRuntimeEstimateInherited = "PASS_GATE739_LEVEL1_SCALAR_RUNTIME_ESTIMATE_INHERITED"
	StatusRuntimeQuarticClassified                    = "PASS_RUNTIME_QUARTIC_CLASSIFIED"
	StatusTreeLevelProxyRelationAudited               = "PASS_TREE_LEVEL_PROXY_RELATION_AUDITED"
	StatusHiggsMassRequiredInputsListed               = "PASS_HIGGS_MASS_REQUIRED_INPUTS_LISTED"
	StatusProxyVersusPoleFirewallEnforced             = "PASS_PROXY_VERSUS_POLE_FIREWALL_ENFORCED"
	StatusSealDependenceCarriedIntoMassTranslation    = "PASS_SEAL_DEPENDENCE_CARRIED_INTO_MASS_TRANSLATION"
	StatusForecastLevelsRefined                       = "PASS_FORECAST_LEVELS_REFINED"

	StatusRuntimeLambdaCanEnterTreeLevelProxyForm          = "CONDITIONAL_SUPPORT_RUNTIME_LAMBDA_CAN_ENTER_TREE_LEVEL_PROXY_FORM"
	StatusLevel1BTreeProxyEstimateAllowedWithExplicitSeals = "CONDITIONAL_SUPPORT_LEVEL_1B_TREE_PROXY_ESTIMATE_ALLOWED_WITH_EXPLICIT_SEALS"

	StatusRuntimeLambdaNotIndependentlyDerived     = "FAILED_ROUTE_RUNTIME_LAMBDA_NOT_INDEPENDENTLY_DERIVED"
	StatusRuntimeLambdaNotPoleMass                 = "FAILED_ROUTE_RUNTIME_LAMBDA_NOT_POLE_MASS"
	StatusTreeProxyNotHiggsPoleMassTheorem         = "FAILED_ROUTE_TREE_PROXY_NOT_HIGGS_POLE_MASS_THEOREM"
	StatusNoNativeVEVOrElectroweakScaleTheorem     = "FAILED_ROUTE_NO_NATIVE_VEV_OR_ELECTROWEAK_SCALE_THEOREM"
	StatusNoNativeRGThresholdPoleCorrectionTheorem = "FAILED_ROUTE_NO_NATIVE_RG_THRESHOLD_POLE_CORRECTION_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem             = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem      = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusLevel2HiggsPoleMassPredictionNotAllowed  = "FAILED_ROUTE_LEVEL_2_HIGGS_POLE_MASS_PREDICTION_NOT_ALLOWED"
	StatusGate740Boundary                          = "FIREWALL_PRESERVED_GATE740_HIGGS_MASS_TRANSLATION_BOUNDARY"
)

type Gate739Inheritance struct {
	Inherited                    bool
	LambdaRuntimeBridge          float64
	RuntimeResidual              float64
	Level1Allowed                bool
	NotIndependentPrediction     bool
	NoNativeScalarRuntimeTheorem bool
	RuntimeBridgeNotHiggsMass    bool
	NoHiggsMassTheorem           bool
	NoYukawaTheorem              bool
	Verdict                      string
}

type RuntimeQuarticClassification struct {
	LambdaRuntimeBridge        float64
	ClassifiedAsRuntimeQuartic bool
	BridgeLayer                bool
	PhysicalPoleMass           bool
	Verdict                    string
}

type TreeLevelProxyRelation struct {
	Relation            string
	RequiresV           bool
	RequiresConvention  bool
	LambdaRuntime       float64
	SqrtTwoLambdaFactor float64
	ConventionDependent bool
	PoleMassTheorem     bool
	Verdict             string
}

type RequiredInput struct {
	Name   string
	Reason string
}

type RequiredInputsAudit struct {
	Inputs            []RequiredInput
	AllListed         bool
	HasVEV            bool
	HasConvention     bool
	HasScaleMatching  bool
	HasRGTransport    bool
	HasThresholdLoop  bool
	HasGaugeYukawaTop bool
	HasUncertainty    bool
	Verdict           string
}

type ProxyPoleFirewall struct {
	RuntimeLambdaEqualsPoleMass          bool
	TreeProxyEqualsPoleMass              bool
	NearAgreementIsIndependentPrediction bool
	RuntimeLambdaNotPoleMass             bool
	TreeProxyNotPoleMassTheorem          bool
	Verdict                              string
}

type SealCarryover struct {
	Seals                      []string
	Explicit                   bool
	TreeProxyWouldRemainLevel1 bool
	NoSealReduction            bool
	Verdict                    string
}

type ForecastLevels struct {
	Level1AName    string
	Level1AAllowed bool
	Level1BName    string
	Level1BAllowed bool
	Level2Name     string
	Level2Allowed  bool
	Verdict        string
}

type Analysis struct {
	Gate739  Gate739Inheritance
	Quartic  RuntimeQuarticClassification
	Proxy    TreeLevelProxyRelation
	Required RequiredInputsAudit
	Firewall ProxyPoleFirewall
	Seals    SealCarryover
	Forecast ForecastLevels
	Truth    string
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
	g739, err := gate739.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate739 inheritance unavailable: %w", err)
	}
	inherit := buildGate739Inheritance(g739)
	quartic := buildQuarticClassification(inherit)
	proxy := buildTreeProxyRelation(inherit)
	required := buildRequiredInputsAudit()
	firewall := buildProxyPoleFirewall()
	seals := buildSealCarryover(g739)
	forecast := buildForecastLevels()
	truth := "Gate 740 inherits the Level-1 scalar-runtime bridge estimate lambda_runtime_bridge and audits only the translation boundary from runtime quartic to Higgs mass. The tree-level relation m_H_tree_proxy=sqrt(2 lambda_runtime) v is recorded as a convention-dependent proxy requiring v, scalar-potential normalization, scale matching, RG/threshold corrections, gauge/Yukawa inputs, and uncertainty propagation. The gate blocks promotion of runtime lambda or a tree proxy to a physical Higgs pole-mass theorem."
	return Analysis{Gate739: inherit, Quartic: quartic, Proxy: proxy, Required: required, Firewall: firewall, Seals: seals, Forecast: forecast, Truth: truth}, nil
}

func buildGate739Inheritance(g gate739.Analysis) Gate739Inheritance {
	return Gate739Inheritance{
		Inherited:                    g.Estimate.Level1Allowed && g.Estimate.NearFloatScale && g.NonPrediction.ConsistencyClosure && !g.HiggsMass.RuntimeLambdaBridgeIsHiggsMassTheorem,
		LambdaRuntimeBridge:          g.Estimate.RuntimeBridge,
		RuntimeResidual:              g.Estimate.RuntimeResidual,
		Level1Allowed:                g.Estimate.Level1Allowed,
		NotIndependentPrediction:     !g.NonPrediction.IndependentRuntimePrediction,
		NoNativeScalarRuntimeTheorem: strings.Contains(strings.Join(gate739.Statuses(), "\n"), gate739.StatusNoNativeScalarRuntimeTheorem),
		RuntimeBridgeNotHiggsMass:    !g.HiggsMass.RuntimeLambdaBridgeIsHiggsMassTheorem,
		NoHiggsMassTheorem:           !g.HiggsMass.RuntimeLambdaBridgeIsHiggsMassTheorem && !g.HiggsMass.HasPoleMassCorrectionTheorem,
		NoYukawaTheorem:              strings.Contains(strings.Join(gate739.Statuses(), "\n"), gate739.StatusNoYukawaOperatorOrEigenvalueTheorem),
		Verdict:                      StatusGate739Level1ScalarRuntimeEstimateInherited,
	}
}

func buildQuarticClassification(g Gate739Inheritance) RuntimeQuarticClassification {
	return RuntimeQuarticClassification{
		LambdaRuntimeBridge:        g.LambdaRuntimeBridge,
		ClassifiedAsRuntimeQuartic: true,
		BridgeLayer:                true,
		PhysicalPoleMass:           false,
		Verdict:                    StatusRuntimeQuarticClassified,
	}
}

func buildTreeProxyRelation(g Gate739Inheritance) TreeLevelProxyRelation {
	return TreeLevelProxyRelation{
		Relation:            "m_H_tree_proxy=sqrt(2*lambda_runtime)*v",
		RequiresV:           true,
		RequiresConvention:  true,
		LambdaRuntime:       g.LambdaRuntimeBridge,
		SqrtTwoLambdaFactor: math.Sqrt(2 * g.LambdaRuntimeBridge),
		ConventionDependent: true,
		PoleMassTheorem:     false,
		Verdict:             strings.Join([]string{StatusTreeLevelProxyRelationAudited, StatusRuntimeLambdaCanEnterTreeLevelProxyForm}, "; "),
	}
}

func buildRequiredInputsAudit() RequiredInputsAudit {
	inputs := []RequiredInput{
		{Name: "VEV / electroweak scale", Reason: "v must be supplied or derived before sqrt(2 lambda_runtime) v has mass dimension"},
		{Name: "scalar-potential convention", Reason: "normalization of the quartic and mass relation must be fixed"},
		{Name: "scale matching", Reason: "lambda_runtime must be located at the correct physical scale"},
		{Name: "RG transport", Reason: "movement between Lambda12, M_Z, pole scale, or other scales must be justified"},
		{Name: "threshold / loop corrections", Reason: "tree proxy must be converted into a pole observable"},
		{Name: "gauge/Yukawa/top-sector dependence", Reason: "pole corrections depend on other Standard Model inputs"},
		{Name: "uncertainty propagation", Reason: "bridge and measurement uncertainties must remain separated"},
	}
	return RequiredInputsAudit{Inputs: inputs, AllListed: len(inputs) == 7, HasVEV: true, HasConvention: true, HasScaleMatching: true, HasRGTransport: true, HasThresholdLoop: true, HasGaugeYukawaTop: true, HasUncertainty: true, Verdict: StatusHiggsMassRequiredInputsListed}
}

func buildProxyPoleFirewall() ProxyPoleFirewall {
	return ProxyPoleFirewall{
		RuntimeLambdaEqualsPoleMass:          false,
		TreeProxyEqualsPoleMass:              false,
		NearAgreementIsIndependentPrediction: false,
		RuntimeLambdaNotPoleMass:             true,
		TreeProxyNotPoleMassTheorem:          true,
		Verdict:                              strings.Join([]string{StatusProxyVersusPoleFirewallEnforced, StatusRuntimeLambdaNotPoleMass, StatusTreeProxyNotHiggsPoleMassTheorem}, "; "),
	}
}

func buildSealCarryover(g gate739.Analysis) SealCarryover {
	seals := make([]string, 0, len(g.Seals.Labels))
	for _, label := range g.Seals.Labels {
		seals = append(seals, label.Name)
	}
	return SealCarryover{Seals: seals, Explicit: len(seals) == 10, TreeProxyWouldRemainLevel1: true, NoSealReduction: true, Verdict: StatusSealDependenceCarriedIntoMassTranslation}
}

func buildForecastLevels() ForecastLevels {
	return ForecastLevels{
		Level1AName:    "Level 1A scalar-runtime bridge consistency estimate",
		Level1AAllowed: true,
		Level1BName:    "Level 1B tree-level Higgs proxy estimate with supplied v and convention",
		Level1BAllowed: true,
		Level2Name:     "Level 2 physical Higgs pole-mass prediction",
		Level2Allowed:  false,
		Verdict:        strings.Join([]string{StatusForecastLevelsRefined, StatusLevel1BTreeProxyEstimateAllowedWithExplicitSeals, StatusLevel2HiggsPoleMassPredictionNotAllowed}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate739Level1ScalarRuntimeEstimateInherited,
		StatusRuntimeQuarticClassified,
		StatusTreeLevelProxyRelationAudited,
		StatusHiggsMassRequiredInputsListed,
		StatusProxyVersusPoleFirewallEnforced,
		StatusSealDependenceCarriedIntoMassTranslation,
		StatusForecastLevelsRefined,
		StatusRuntimeLambdaCanEnterTreeLevelProxyForm,
		StatusLevel1BTreeProxyEstimateAllowedWithExplicitSeals,
		StatusRuntimeLambdaNotIndependentlyDerived,
		StatusRuntimeLambdaNotPoleMass,
		StatusTreeProxyNotHiggsPoleMassTheorem,
		StatusNoNativeVEVOrElectroweakScaleTheorem,
		StatusNoNativeRGThresholdPoleCorrectionTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusLevel2HiggsPoleMassPredictionNotAllowed,
		StatusGate740Boundary,
	}
}

func FormatGate739(x Gate739Inheritance) string {
	return fmt.Sprintf("inherited=%t lambdaRuntime=%.17g residual=%.17g level1=%t notIndependent=%t noRuntime=%t notMass=%t noMass=%t noYukawa=%t verdict=%q", x.Inherited, x.LambdaRuntimeBridge, x.RuntimeResidual, x.Level1Allowed, x.NotIndependentPrediction, x.NoNativeScalarRuntimeTheorem, x.RuntimeBridgeNotHiggsMass, x.NoHiggsMassTheorem, x.NoYukawaTheorem, x.Verdict)
}

func FormatQuartic(x RuntimeQuarticClassification) string {
	return fmt.Sprintf("lambdaRuntime=%.17g runtimeQuartic=%t bridgeLayer=%t physicalPoleMass=%t verdict=%q", x.LambdaRuntimeBridge, x.ClassifiedAsRuntimeQuartic, x.BridgeLayer, x.PhysicalPoleMass, x.Verdict)
}

func FormatProxy(x TreeLevelProxyRelation) string {
	return fmt.Sprintf("relation=%q requiresV=%t requiresConvention=%t lambda=%.17g sqrt2lambda=%.17g conventionDependent=%t poleTheorem=%t verdict=%q", x.Relation, x.RequiresV, x.RequiresConvention, x.LambdaRuntime, x.SqrtTwoLambdaFactor, x.ConventionDependent, x.PoleMassTheorem, x.Verdict)
}

func FormatRequired(x RequiredInputsAudit) string {
	names := make([]string, 0, len(x.Inputs))
	for _, input := range x.Inputs {
		names = append(names, input.Name)
	}
	return fmt.Sprintf("inputs=[%s] allListed=%t vev=%t convention=%t scale=%t rg=%t threshold=%t top=%t uncertainty=%t verdict=%q", strings.Join(names, " | "), x.AllListed, x.HasVEV, x.HasConvention, x.HasScaleMatching, x.HasRGTransport, x.HasThresholdLoop, x.HasGaugeYukawaTop, x.HasUncertainty, x.Verdict)
}

func FormatFirewall(x ProxyPoleFirewall) string {
	return fmt.Sprintf("runtimeEqualsPole=%t treeEqualsPole=%t nearAgreementPrediction=%t runtimeNotPole=%t treeNotPole=%t verdict=%q", x.RuntimeLambdaEqualsPoleMass, x.TreeProxyEqualsPoleMass, x.NearAgreementIsIndependentPrediction, x.RuntimeLambdaNotPoleMass, x.TreeProxyNotPoleMassTheorem, x.Verdict)
}

func FormatSeals(x SealCarryover) string {
	return fmt.Sprintf("seals=[%s] explicit=%t treeProxyLevel1=%t noReduction=%t verdict=%q", strings.Join(x.Seals, ","), x.Explicit, x.TreeProxyWouldRemainLevel1, x.NoSealReduction, x.Verdict)
}

func FormatForecast(x ForecastLevels) string {
	return fmt.Sprintf("1A=%q allowed=%t 1B=%q allowed=%t level2=%q allowed=%t verdict=%q", x.Level1AName, x.Level1AAllowed, x.Level1BName, x.Level1BAllowed, x.Level2Name, x.Level2Allowed, x.Verdict)
}
