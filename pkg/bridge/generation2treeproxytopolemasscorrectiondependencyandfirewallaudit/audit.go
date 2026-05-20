// Package generation2treeproxytopolemasscorrectiondependencyandfirewallaudit implements
// Gate 742: Tree Proxy to Pole-Mass Correction Dependency and Firewall Audit.
//
// Gate 741 computed an allowed Level-1B sealed tree-level Higgs proxy. Gate 742
// audits the missing pole-translation layer: the formal correction object,
// required RG/threshold/top/gauge inputs, inherited seals, Level-1C diagnostic
// boundary, and the firewall against reading the tree proxy as a pole-mass
// prediction.
package generation2treeproxytopolemasscorrectiondependencyandfirewallaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate741 "github.com/bagherbal/asha-engine/pkg/bridge/generation2level1bhiggstreeproxyestimateandvevconventionfirewallaudit"
)

const (
	AuditID = "GATE742-TREE-PROXY-TO-POLE-MASS-CORRECTION-DEPENDENCY-FIREWALL-AUDIT"

	StatusGate741Level1BTreeProxyInherited        = "PASS_GATE741_LEVEL1B_TREE_PROXY_INHERITED"
	StatusPoleCorrectionObjectDefined             = "PASS_POLE_CORRECTION_OBJECT_DEFINED"
	StatusRequiredCorrectionIngredientsListed     = "PASS_REQUIRED_CORRECTION_INGREDIENTS_LISTED"
	StatusTreeProxyVersusPoleFirewallEnforced     = "PASS_TREE_PROXY_VERSUS_POLE_FIREWALL_ENFORCED"
	StatusSealInheritanceAudited                  = "PASS_SEAL_INHERITANCE_AUDITED"
	StatusForecastLevelsRefined                   = "PASS_FORECAST_LEVELS_REFINED"
	StatusTreeProxyInputToPolePipeline            = "CONDITIONAL_SUPPORT_TREE_PROXY_CAN_BE_INPUT_TO_POLE_TRANSLATION_PIPELINE"
	StatusLevel1CDiagnosticAllowedExternalPackage = "CONDITIONAL_SUPPORT_LEVEL_1C_DIAGNOSTIC_COMPARISON_ALLOWED_ONLY_WITH_EXTERNAL_CORRECTION_PACKAGE"
	StatusTreeProxyIsNotPoleMass                  = "FAILED_ROUTE_TREE_PROXY_IS_NOT_POLE_MASS"
	StatusNoNativeTreeToPoleCorrectionTheorem     = "FAILED_ROUTE_NO_NATIVE_TREE_TO_POLE_CORRECTION_THEOREM"
	StatusNoNativeRGThresholdMatchingTheorem      = "FAILED_ROUTE_NO_NATIVE_RG_THRESHOLD_MATCHING_THEOREM"
	StatusNoNativeTopYukawaOrGaugeInputTheorem    = "FAILED_ROUTE_NO_NATIVE_TOP_YUKAWA_OR_GAUGE_INPUT_THEOREM"
	StatusNoIndependentHiggsPoleMassPrediction    = "FAILED_ROUTE_NO_INDEPENDENT_HIGGS_POLE_MASS_PREDICTION"
	StatusNoYukawaOperatorOrEigenvalueTheorem     = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate742Boundary                         = "FIREWALL_PRESERVED_GATE742_TREE_PROXY_TO_POLE_BOUNDARY"
)

var RequiredPoleCorrectionIngredients = []string{
	"scalar-potential convention",
	"renormalization scheme",
	"renormalization scale",
	"RG transport of lambda",
	"top Yukawa / top mass input",
	"gauge coupling inputs",
	"electroweak threshold corrections",
	"loop-order convention",
	"matching between running mass, tree proxy, and pole observable",
	"uncertainty propagation",
}

var PoleLayerSealLabels = []string{
	"RGSchemeSeal",
	"PoleMassConventionSeal",
	"ThresholdCorrectionSeal",
	"TopYukawaInputSeal",
	"GaugeCouplingInputSeal",
}

type Gate741Inheritance struct {
	Inherited                bool
	LambdaRuntimeBridge      float64
	VEVGeV                   float64
	TreeProxyGeV             float64
	Level1B                  bool
	NotPoleMass              bool
	NotIndependentPrediction bool
	NoHiggsMassTheorem       bool
	NoYukawaTheorem          bool
	Verdict                  string
}

type PoleCorrectionObject struct {
	Name                       string
	Formula                    string
	ValueAssigned              bool
	RequiresPoleConvention     bool
	RequiresExternalCorrection bool
	Verdict                    string
}

type CorrectionIngredients struct {
	Items            []string
	Count            int
	AllListed        bool
	NoNativeRG       bool
	NoNativeTopGauge bool
	Verdict          string
}

type TreePoleFirewall struct {
	TreeProxyEqualsPoleMass            bool
	NearNumericalProximityIsPrediction bool
	PoleObservableNeedsLoopThreshold   bool
	TreeProxyConventionLevel           bool
	Verdict                            string
}

type SealInheritance struct {
	Gate741Seals      []string
	PoleLayerSeals    []string
	TotalCount        int
	IncludesRGScheme  bool
	IncludesPoleMass  bool
	IncludesThreshold bool
	IncludesTopYukawa bool
	IncludesGauge     bool
	Explicit          bool
	Verdict           string
}

type ForecastLevels struct {
	Level1BName             string
	Level1BAllowed          bool
	Level1CName             string
	Level1CAllowed          bool
	Level1CDiagnosticOnly   bool
	Level1CRequiresExternal bool
	Level2Name              string
	Level2Allowed           bool
	Verdict                 string
}

type Analysis struct {
	Gate741     Gate741Inheritance
	Correction  PoleCorrectionObject
	Ingredients CorrectionIngredients
	Firewall    TreePoleFirewall
	Seals       SealInheritance
	Forecast    ForecastLevels
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
	g741, err := gate741.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate741 inheritance unavailable: %w", err)
	}
	inherit := buildGate741Inheritance(g741)
	correction := buildPoleCorrectionObject()
	ingredients := buildCorrectionIngredients()
	firewall := buildTreePoleFirewall()
	seals := buildSealInheritance(g741)
	forecast := buildForecastLevels()
	truth := "Gate 742 defines the pole-translation layer as a missing correction package. The Gate741 tree proxy may feed a future tree-to-pole pipeline, but no pole correction, RG/threshold matching, top/gauge input theorem, or independent Higgs pole-mass prediction is certified."
	return Analysis{Gate741: inherit, Correction: correction, Ingredients: ingredients, Firewall: firewall, Seals: seals, Forecast: forecast, Truth: truth}, nil
}

func buildGate741Inheritance(g gate741.Analysis) Gate741Inheritance {
	statuses := strings.Join(gate741.Statuses(), "\n")
	return Gate741Inheritance{
		Inherited:                g.Gate740.Inherited && g.Proxy.Level == "Level 1B sealed tree-level proxy estimate" && !g.Proxy.PoleMassPrediction && g.Level.Level1BAllowed && !g.Level.Level2Allowed,
		LambdaRuntimeBridge:      g.Runtime.LambdaRuntimeBridge,
		VEVGeV:                   g.VEV.VGeV,
		TreeProxyGeV:             g.Proxy.TreeProxyGeV,
		Level1B:                  g.Level.Level1BAllowed,
		NotPoleMass:              strings.Contains(statuses, gate741.StatusTreeProxyNotHiggsPoleMass),
		NotIndependentPrediction: strings.Contains(statuses, gate741.StatusRuntimeLambdaNotIndependentlyDerived),
		NoHiggsMassTheorem:       strings.Contains(statuses, gate741.StatusNoHiggsMassOrPoleMassTheorem),
		NoYukawaTheorem:          strings.Contains(statuses, gate741.StatusNoYukawaOperatorOrEigenvalueTheorem),
		Verdict:                  StatusGate741Level1BTreeProxyInherited,
	}
}

func buildPoleCorrectionObject() PoleCorrectionObject {
	return PoleCorrectionObject{
		Name:                       "Delta_pole",
		Formula:                    "Delta_pole = m_H_pole - m_H_tree_proxy",
		ValueAssigned:              false,
		RequiresPoleConvention:     true,
		RequiresExternalCorrection: true,
		Verdict:                    strings.Join([]string{StatusPoleCorrectionObjectDefined, StatusNoNativeTreeToPoleCorrectionTheorem}, "; "),
	}
}

func buildCorrectionIngredients() CorrectionIngredients {
	items := append([]string{}, RequiredPoleCorrectionIngredients...)
	return CorrectionIngredients{
		Items:            items,
		Count:            len(items),
		AllListed:        len(items) == 10,
		NoNativeRG:       true,
		NoNativeTopGauge: true,
		Verdict:          strings.Join([]string{StatusRequiredCorrectionIngredientsListed, StatusNoNativeRGThresholdMatchingTheorem, StatusNoNativeTopYukawaOrGaugeInputTheorem}, "; "),
	}
}

func buildTreePoleFirewall() TreePoleFirewall {
	return TreePoleFirewall{
		TreeProxyEqualsPoleMass:            false,
		NearNumericalProximityIsPrediction: false,
		PoleObservableNeedsLoopThreshold:   true,
		TreeProxyConventionLevel:           true,
		Verdict:                            strings.Join([]string{StatusTreeProxyVersusPoleFirewallEnforced, StatusTreeProxyIsNotPoleMass, StatusNoIndependentHiggsPoleMassPrediction}, "; "),
	}
}

func buildSealInheritance(g gate741.Analysis) SealInheritance {
	gateSeals := append([]string{}, g.Seals.Labels...)
	poleSeals := append([]string{}, PoleLayerSealLabels...)
	return SealInheritance{
		Gate741Seals:      gateSeals,
		PoleLayerSeals:    poleSeals,
		TotalCount:        len(gateSeals) + len(poleSeals),
		IncludesRGScheme:  contains(poleSeals, "RGSchemeSeal"),
		IncludesPoleMass:  contains(poleSeals, "PoleMassConventionSeal"),
		IncludesThreshold: contains(poleSeals, "ThresholdCorrectionSeal"),
		IncludesTopYukawa: contains(poleSeals, "TopYukawaInputSeal"),
		IncludesGauge:     contains(poleSeals, "GaugeCouplingInputSeal"),
		Explicit:          len(gateSeals) == 12 && len(poleSeals) == 5,
		Verdict:           StatusSealInheritanceAudited,
	}
}

func buildForecastLevels() ForecastLevels {
	return ForecastLevels{
		Level1BName:             "Level 1B sealed tree-level proxy estimate",
		Level1BAllowed:          true,
		Level1CName:             "Level 1C tree-to-pole diagnostic comparison with external correction package",
		Level1CAllowed:          true,
		Level1CDiagnosticOnly:   true,
		Level1CRequiresExternal: true,
		Level2Name:              "Level 2 independent Higgs pole-mass prediction",
		Level2Allowed:           false,
		Verdict:                 strings.Join([]string{StatusForecastLevelsRefined, StatusTreeProxyInputToPolePipeline, StatusLevel1CDiagnosticAllowedExternalPackage, StatusNoIndependentHiggsPoleMassPrediction}, "; "),
	}
}

func contains(items []string, want string) bool {
	for _, item := range items {
		if item == want {
			return true
		}
	}
	return false
}

func Statuses() []string {
	return []string{
		StatusGate741Level1BTreeProxyInherited,
		StatusPoleCorrectionObjectDefined,
		StatusRequiredCorrectionIngredientsListed,
		StatusTreeProxyVersusPoleFirewallEnforced,
		StatusSealInheritanceAudited,
		StatusForecastLevelsRefined,
		StatusTreeProxyInputToPolePipeline,
		StatusLevel1CDiagnosticAllowedExternalPackage,
		StatusTreeProxyIsNotPoleMass,
		StatusNoNativeTreeToPoleCorrectionTheorem,
		StatusNoNativeRGThresholdMatchingTheorem,
		StatusNoNativeTopYukawaOrGaugeInputTheorem,
		StatusNoIndependentHiggsPoleMassPrediction,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate742Boundary,
	}
}

func FormatGate741(x Gate741Inheritance) string {
	return fmt.Sprintf("inherited=%t lambda=%.17g v=%.10f treeProxy=%.17g level1B=%t notPole=%t notIndependent=%t noMass=%t noYukawa=%t verdict=%q", x.Inherited, x.LambdaRuntimeBridge, x.VEVGeV, x.TreeProxyGeV, x.Level1B, x.NotPoleMass, x.NotIndependentPrediction, x.NoHiggsMassTheorem, x.NoYukawaTheorem, x.Verdict)
}

func FormatCorrection(x PoleCorrectionObject) string {
	return fmt.Sprintf("name=%s formula=%q valueAssigned=%t requiresPoleConvention=%t requiresExternalCorrection=%t verdict=%q", x.Name, x.Formula, x.ValueAssigned, x.RequiresPoleConvention, x.RequiresExternalCorrection, x.Verdict)
}

func FormatIngredients(x CorrectionIngredients) string {
	return fmt.Sprintf("count=%d allListed=%t noNativeRG=%t noNativeTopGauge=%t items=[%s] verdict=%q", x.Count, x.AllListed, x.NoNativeRG, x.NoNativeTopGauge, strings.Join(x.Items, ","), x.Verdict)
}

func FormatFirewall(x TreePoleFirewall) string {
	return fmt.Sprintf("treeEqualsPole=%t proximityPrediction=%t needsLoopThreshold=%t conventionLevel=%t verdict=%q", x.TreeProxyEqualsPoleMass, x.NearNumericalProximityIsPrediction, x.PoleObservableNeedsLoopThreshold, x.TreeProxyConventionLevel, x.Verdict)
}

func FormatSeals(x SealInheritance) string {
	return fmt.Sprintf("gate741=[%s] poleLayer=[%s] total=%d rg=%t pole=%t threshold=%t top=%t gauge=%t explicit=%t verdict=%q", strings.Join(x.Gate741Seals, ","), strings.Join(x.PoleLayerSeals, ","), x.TotalCount, x.IncludesRGScheme, x.IncludesPoleMass, x.IncludesThreshold, x.IncludesTopYukawa, x.IncludesGauge, x.Explicit, x.Verdict)
}

func FormatForecast(x ForecastLevels) string {
	return fmt.Sprintf("level1B=%q allowed=%t level1C=%q allowed=%t diagnostic=%t external=%t level2=%q allowed=%t verdict=%q", x.Level1BName, x.Level1BAllowed, x.Level1CName, x.Level1CAllowed, x.Level1CDiagnosticOnly, x.Level1CRequiresExternal, x.Level2Name, x.Level2Allowed, x.Verdict)
}

func NearlyEqual(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
