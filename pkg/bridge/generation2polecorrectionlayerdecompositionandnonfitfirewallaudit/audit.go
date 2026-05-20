// Package generation2polecorrectionlayerdecompositionandnonfitfirewallaudit implements
// Gate 744: Pole-Correction Layer Decomposition and Non-Fit Firewall Audit.
//
// Gate 743 defined the minimal pole-correction seal package required for a
// lawful Level-1C diagnostic comparison. Gate 744 audits the internal layered
// structure of Delta_pole while preserving the firewall that it cannot be
// compressed into a fitted native theorem.
package generation2polecorrectionlayerdecompositionandnonfitfirewallaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate743 "github.com/bagherbal/asha-engine/pkg/bridge/generation2polecorrectionsealpackageandlevel1cdiagnosticboundaryaudit"
)

const (
	AuditID = "GATE744-POLE-CORRECTION-LAYER-DECOMPOSITION-NON-FIT-FIREWALL-AUDIT"

	StatusGate743PoleCorrectionSealPackageInherited = "PASS_GATE743_POLE_CORRECTION_SEAL_PACKAGE_INHERITED"
	StatusDeltaPoleKeptSymbolic                     = "PASS_DELTA_POLE_KEPT_SYMBOLIC"
	StatusCorrectionLayerDecompositionDefined       = "PASS_CORRECTION_LAYER_DECOMPOSITION_DEFINED"
	StatusCorrectionLayerMinimalityAudited          = "PASS_CORRECTION_LAYER_MINIMALITY_AUDITED"
	StatusNonFitFirewallEnforced                    = "PASS_NON_FIT_FIREWALL_ENFORCED"
	StatusForecastBoundaryPreserved                 = "PASS_FORECAST_BOUNDARY_PRESERVED"

	StatusDeltaPoleMultiLayerCorrectionObject        = "CONDITIONAL_SUPPORT_DELTA_POLE_IS_MULTI_LAYER_CORRECTION_OBJECT"
	StatusLevel1CDiagnosticRequiresLayeredCorrection = "CONDITIONAL_SUPPORT_LEVEL_1C_DIAGNOSTIC_REQUIRES_LAYERED_CORRECTION_PACKAGE"

	StatusDeltaPoleCannotBeFittedAsNativeTheorem = "FAILED_ROUTE_DELTA_POLE_CANNOT_BE_FITTED_AS_NATIVE_THEOREM"
	StatusNoNativeRGThresholdMatchingTheorem     = "FAILED_ROUTE_NO_NATIVE_RG_THRESHOLD_MATCHING_THEOREM"
	StatusNoNativeTopYukawaOrGaugeInputTheorem   = "FAILED_ROUTE_NO_NATIVE_TOP_YUKAWA_OR_GAUGE_INPUT_THEOREM"
	StatusNoNativeTreeToPoleCorrectionTheorem    = "FAILED_ROUTE_NO_NATIVE_TREE_TO_POLE_CORRECTION_THEOREM"
	StatusNoIndependentHiggsPoleMassPrediction   = "FAILED_ROUTE_NO_INDEPENDENT_HIGGS_POLE_MASS_PREDICTION"
	StatusNoYukawaOperatorOrEigenvalueTheorem    = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate744Boundary                        = "FIREWALL_PRESERVED_GATE744_POLE_CORRECTION_LAYER_BOUNDARY"
)

var CorrectionLayerNames = []string{
	"Delta_RG",
	"Delta_threshold",
	"Delta_scheme",
	"Delta_loop",
	"Delta_top/gauge",
	"Delta_uncertainty",
}

type Gate743Inheritance struct {
	Inherited                    bool
	TreeProxyGeV                 float64
	DeltaPoleObject              string
	DeltaPoleValueAssigned       bool
	FullCorrectionPackageDefined bool
	Level1CAllowed               bool
	Level1CDiagnosticOnly        bool
	Level2Allowed                bool
	Verdict                      string
}

type DeltaPoleSymbolic struct {
	Name                       string
	Expression                 string
	ValueAssigned              bool
	ExternalObservableSupplied bool
	CorrectionPackageSupplied  bool
	Verdict                    string
}

type CorrectionLayer struct {
	Name          string
	Role          string
	Required      bool
	NativeDerived bool
}

type LayerDecomposition struct {
	Layers            []CorrectionLayer
	Count             int
	FormalExpression  string
	AllRequired       bool
	AnyNativeDerived  bool
	CompressibleToFit bool
	Verdict           string
}

type MinimalityItem struct {
	Layer         string
	RemovalEffect string
	Required      bool
}

type MinimalityAudit struct {
	Items       []MinimalityItem
	Count       int
	AllRequired bool
	Minimal     bool
	Verdict     string
}

type NonFitFirewall struct {
	ObservedMinusProxyIsDerivedTheorem bool
	ExternalDiagnosticAllowed          bool
	ExternalDiagnosticIsPrediction     bool
	DeltaPoleKeptLayered               bool
	SingleFittedNumberLosesTypeInfo    bool
	Verdict                            string
}

type SourceTypeClassification struct {
	TreeProxyLevel       string
	TreeProxyStatus      string
	DeltaPoleStatus      string
	PoleObservableStatus string
	Level1BAllowed       bool
	Level1CAllowed       bool
	Level2Allowed        bool
	Verdict              string
}

type Analysis struct {
	Gate743        Gate743Inheritance
	DeltaPole      DeltaPoleSymbolic
	Decomposition  LayerDecomposition
	Minimality     MinimalityAudit
	NonFit         NonFitFirewall
	Classification SourceTypeClassification
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
	g743, err := gate743.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate743 inheritance unavailable: %w", err)
	}
	inheritance := buildGate743Inheritance(g743)
	delta := buildDeltaPoleSymbolic()
	decomposition := buildLayerDecomposition()
	minimality := buildMinimalityAudit()
	nonfit := buildNonFitFirewall()
	classification := buildSourceTypeClassification()
	truth := "Gate 744 keeps Delta_pole symbolic and decomposes it into RG, threshold, scheme, loop, top/gauge, and uncertainty layers. The correction remains a Level-1C diagnostic object only; fitting observed pole mass minus tree proxy is not a native ASHA tree-to-pole theorem."
	return Analysis{Gate743: inheritance, DeltaPole: delta, Decomposition: decomposition, Minimality: minimality, NonFit: nonfit, Classification: classification, Truth: truth}, nil
}

func buildGate743Inheritance(g gate743.Analysis) Gate743Inheritance {
	return Gate743Inheritance{
		Inherited:                    g.Gate742.Inherited && g.Package.FullPackage && g.Diagnostic.Level1CAllowed && g.Diagnostic.Level1CDiagnosticOnly && !g.Diagnostic.Level2Allowed,
		TreeProxyGeV:                 g.Gate742.TreeProxyGeV,
		DeltaPoleObject:              g.Gate742.CorrectionObject,
		DeltaPoleValueAssigned:       g.Gate742.CorrectionValueAssigned || g.Package.DeltaPoleValueAssigned,
		FullCorrectionPackageDefined: g.Package.FullPackage,
		Level1CAllowed:               g.Diagnostic.Level1CAllowed,
		Level1CDiagnosticOnly:        g.Diagnostic.Level1CDiagnosticOnly,
		Level2Allowed:                g.Diagnostic.Level2Allowed,
		Verdict:                      StatusGate743PoleCorrectionSealPackageInherited,
	}
}

func buildDeltaPoleSymbolic() DeltaPoleSymbolic {
	return DeltaPoleSymbolic{
		Name:                       "Delta_pole",
		Expression:                 "m_H_pole - m_H_tree_proxy",
		ValueAssigned:              false,
		ExternalObservableSupplied: false,
		CorrectionPackageSupplied:  false,
		Verdict:                    StatusDeltaPoleKeptSymbolic,
	}
}

func buildLayerDecomposition() LayerDecomposition {
	layers := []CorrectionLayer{
		{Name: "Delta_RG", Role: "correction from running lambda between chosen scales", Required: true, NativeDerived: false},
		{Name: "Delta_threshold", Role: "matching correction between effective running quantities and pole observable", Required: true, NativeDerived: false},
		{Name: "Delta_scheme", Role: "renormalization scheme and scalar-potential convention dependence", Required: true, NativeDerived: false},
		{Name: "Delta_loop", Role: "loop-order truncation correction", Required: true, NativeDerived: false},
		{Name: "Delta_top/gauge", Role: "top Yukawa, top mass, gauge coupling, and electroweak input dependence", Required: true, NativeDerived: false},
		{Name: "Delta_uncertainty", Role: "propagated uncertainty from bridge seals, inputs, and physical measurement", Required: true, NativeDerived: false},
	}
	all := true
	anyNative := false
	for _, layer := range layers {
		all = all && layer.Required && layer.Name != "" && layer.Role != ""
		anyNative = anyNative || layer.NativeDerived
	}
	return LayerDecomposition{
		Layers:            layers,
		Count:             len(layers),
		FormalExpression:  "Delta_pole = Delta_RG + Delta_threshold + Delta_scheme + Delta_loop + Delta_top/gauge + Delta_uncertainty",
		AllRequired:       all,
		AnyNativeDerived:  anyNative,
		CompressibleToFit: false,
		Verdict:           strings.Join([]string{StatusCorrectionLayerDecompositionDefined, StatusDeltaPoleMultiLayerCorrectionObject}, "; "),
	}
}

func buildMinimalityAudit() MinimalityAudit {
	items := []MinimalityItem{
		{Layer: "Delta_RG", RemovalEffect: "no scale-transport correction", Required: true},
		{Layer: "Delta_threshold", RemovalEffect: "no running-to-pole matching", Required: true},
		{Layer: "Delta_scheme", RemovalEffect: "comparison is convention-dependent and ill-typed", Required: true},
		{Layer: "Delta_loop", RemovalEffect: "perturbative order is undefined", Required: true},
		{Layer: "Delta_top/gauge", RemovalEffect: "dominant Standard Model correction dependencies are absent", Required: true},
		{Layer: "Delta_uncertainty", RemovalEffect: "diagnostic comparison has no error ledger", Required: true},
	}
	all := true
	for _, item := range items {
		all = all && item.Required && item.Layer != "" && item.RemovalEffect != ""
	}
	return MinimalityAudit{Items: items, Count: len(items), AllRequired: all, Minimal: all && len(items) == len(CorrectionLayerNames), Verdict: strings.Join([]string{StatusCorrectionLayerMinimalityAudited, StatusLevel1CDiagnosticRequiresLayeredCorrection}, "; ")}
}

func buildNonFitFirewall() NonFitFirewall {
	return NonFitFirewall{
		ObservedMinusProxyIsDerivedTheorem: false,
		ExternalDiagnosticAllowed:          true,
		ExternalDiagnosticIsPrediction:     false,
		DeltaPoleKeptLayered:               true,
		SingleFittedNumberLosesTypeInfo:    true,
		Verdict: strings.Join([]string{
			StatusNonFitFirewallEnforced,
			StatusDeltaPoleCannotBeFittedAsNativeTheorem,
			StatusNoNativeTreeToPoleCorrectionTheorem,
			StatusNoNativeRGThresholdMatchingTheorem,
			StatusNoNativeTopYukawaOrGaugeInputTheorem,
			StatusNoIndependentHiggsPoleMassPrediction,
		}, "; "),
	}
}

func buildSourceTypeClassification() SourceTypeClassification {
	return SourceTypeClassification{
		TreeProxyLevel:       "Level-1B",
		TreeProxyStatus:      "sealed scalar tree proxy",
		DeltaPoleStatus:      "sealed multi-layer pole-correction package object",
		PoleObservableStatus: "physical observable only after PoleMassObservableSeal and convention seals",
		Level1BAllowed:       true,
		Level1CAllowed:       true,
		Level2Allowed:        false,
		Verdict:              StatusForecastBoundaryPreserved,
	}
}

func Statuses() []string {
	return []string{
		StatusGate743PoleCorrectionSealPackageInherited,
		StatusDeltaPoleKeptSymbolic,
		StatusCorrectionLayerDecompositionDefined,
		StatusCorrectionLayerMinimalityAudited,
		StatusNonFitFirewallEnforced,
		StatusForecastBoundaryPreserved,
		StatusDeltaPoleMultiLayerCorrectionObject,
		StatusLevel1CDiagnosticRequiresLayeredCorrection,
		StatusDeltaPoleCannotBeFittedAsNativeTheorem,
		StatusNoNativeRGThresholdMatchingTheorem,
		StatusNoNativeTopYukawaOrGaugeInputTheorem,
		StatusNoNativeTreeToPoleCorrectionTheorem,
		StatusNoIndependentHiggsPoleMassPrediction,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate744Boundary,
	}
}

func FormatGate743(x Gate743Inheritance) string {
	return fmt.Sprintf("inherited=%t treeProxy=%.17g delta=%s assigned=%t fullPackage=%t level1C=%t diagnosticOnly=%t level2=%t verdict=%q", x.Inherited, x.TreeProxyGeV, x.DeltaPoleObject, x.DeltaPoleValueAssigned, x.FullCorrectionPackageDefined, x.Level1CAllowed, x.Level1CDiagnosticOnly, x.Level2Allowed, x.Verdict)
}

func FormatDeltaPole(x DeltaPoleSymbolic) string {
	return fmt.Sprintf("name=%s expression=%q assigned=%t externalObservable=%t packageSupplied=%t verdict=%q", x.Name, x.Expression, x.ValueAssigned, x.ExternalObservableSupplied, x.CorrectionPackageSupplied, x.Verdict)
}

func FormatDecomposition(x LayerDecomposition) string {
	parts := make([]string, 0, len(x.Layers))
	for _, layer := range x.Layers {
		parts = append(parts, layer.Name+":"+layer.Role)
	}
	return fmt.Sprintf("count=%d expression=%q allRequired=%t anyNative=%t compressible=%t layers=[%s] verdict=%q", x.Count, x.FormalExpression, x.AllRequired, x.AnyNativeDerived, x.CompressibleToFit, strings.Join(parts, " | "), x.Verdict)
}

func FormatMinimality(x MinimalityAudit) string {
	parts := make([]string, 0, len(x.Items))
	for _, item := range x.Items {
		parts = append(parts, item.Layer+":"+item.RemovalEffect)
	}
	return fmt.Sprintf("count=%d allRequired=%t minimal=%t items=[%s] verdict=%q", x.Count, x.AllRequired, x.Minimal, strings.Join(parts, " | "), x.Verdict)
}

func FormatNonFit(x NonFitFirewall) string {
	return fmt.Sprintf("observedMinusDerived=%t externalDiagnosticAllowed=%t diagnosticPrediction=%t layered=%t singleFitLosesTypeInfo=%t verdict=%q", x.ObservedMinusProxyIsDerivedTheorem, x.ExternalDiagnosticAllowed, x.ExternalDiagnosticIsPrediction, x.DeltaPoleKeptLayered, x.SingleFittedNumberLosesTypeInfo, x.Verdict)
}

func FormatClassification(x SourceTypeClassification) string {
	return fmt.Sprintf("treeLevel=%s treeStatus=%q deltaStatus=%q poleStatus=%q level1B=%t level1C=%t level2=%t verdict=%q", x.TreeProxyLevel, x.TreeProxyStatus, x.DeltaPoleStatus, x.PoleObservableStatus, x.Level1BAllowed, x.Level1CAllowed, x.Level2Allowed, x.Verdict)
}

func NearlyEqual(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
