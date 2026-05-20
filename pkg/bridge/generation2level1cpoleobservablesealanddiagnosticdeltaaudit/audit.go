// Package generation2level1cpoleobservablesealanddiagnosticdeltaaudit implements
// Gate 745: Level-1C Pole Observable Seal and Diagnostic Delta Audit.
//
// Gate 744 decomposed the symbolic pole correction Delta_pole into typed
// correction layers. Gate 745 audits the lawful Level-1C diagnostic form that
// becomes available only when an external pole observable is supplied through a
// PoleMassObservableSeal, while preserving that the diagnostic gap is not an
// ASHA tree-to-pole theorem.
package generation2level1cpoleobservablesealanddiagnosticdeltaaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate744 "github.com/bagherbal/asha-engine/pkg/bridge/generation2polecorrectionlayerdecompositionandnonfitfirewallaudit"
)

const (
	AuditID = "GATE745-LEVEL-1C-POLE-OBSERVABLE-SEAL-DIAGNOSTIC-DELTA-AUDIT"

	StatusGate744PoleCorrectionLayerInherited          = "PASS_GATE744_POLE_CORRECTION_LAYER_INHERITED"
	StatusPoleObservableSealDefined                    = "PASS_POLE_OBSERVABLE_SEAL_DEFINED"
	StatusLevel1CDiagnosticDeltaFormDefined            = "PASS_LEVEL_1C_DIAGNOSTIC_DELTA_FORM_DEFINED"
	StatusLayerAssignmentWarningRecorded               = "PASS_LAYER_ASSIGNMENT_WARNING_RECORDED"
	StatusNonFitFirewallEnforced                       = "PASS_NON_FIT_FIREWALL_ENFORCED"
	StatusRequiredExplanatoryCorrectionPackageRecorded = "PASS_REQUIRED_EXPLANATORY_CORRECTION_PACKAGE_RECORDED"

	StatusLevel1CDiagnosticDeltaAllowedExternalPole = "CONDITIONAL_SUPPORT_LEVEL_1C_DIAGNOSTIC_DELTA_IS_ALLOWED_WITH_EXTERNAL_POLE_OBSERVABLE"
	StatusDeltaPoleDiagMeasuresGapOnly              = "CONDITIONAL_SUPPORT_DELTA_POLE_DIAG_MEASURES_PROXY_TO_POLE_GAP_ONLY"

	StatusExternalPoleObservableNotASHADerived = "FAILED_ROUTE_EXTERNAL_POLE_OBSERVABLE_IS_NOT_ASHA_DERIVATION"
	StatusDiagnosticDeltaNotTreeToPoleTheorem  = "FAILED_ROUTE_DIAGNOSTIC_DELTA_IS_NOT_TREE_TO_POLE_THEOREM"
	StatusNoNativeRGThresholdMatchingTheorem   = "FAILED_ROUTE_NO_NATIVE_RG_THRESHOLD_MATCHING_THEOREM"
	StatusNoNativeTopYukawaOrGaugeInputTheorem = "FAILED_ROUTE_NO_NATIVE_TOP_YUKAWA_OR_GAUGE_INPUT_THEOREM"
	StatusNoIndependentHiggsPoleMassPrediction = "FAILED_ROUTE_NO_INDEPENDENT_HIGGS_POLE_MASS_PREDICTION"
	StatusNoYukawaOperatorOrEigenvalueTheorem  = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate745Boundary                      = "FIREWALL_PRESERVED_GATE745_LEVEL1C_POLE_DIAGNOSTIC_BOUNDARY"
)

var RequiredExplanatoryCorrectionPackageLabels = []string{
	"RGSchemeSeal",
	"RenormalizationScaleSeal",
	"LoopOrderSeal",
	"ThresholdCorrectionSeal",
	"TopYukawaInputSeal",
	"GaugeCouplingInputSeal",
	"UncertaintyModelSeal",
}

type Gate744Inheritance struct {
	Inherited               bool
	TreeProxyGeV            float64
	DeltaPoleObject         string
	DeltaPoleKeptSymbolic   bool
	CorrectionLayerCount    int
	HasLayeredCorrection    bool
	Level1CAllowed          bool
	Level2Allowed           bool
	NonFitFirewallPreserved bool
	Verdict                 string
}

type PoleObservableSeal struct {
	Name             string
	Object           string
	ExternalInput    bool
	ValueSupplied    bool
	Unit             string
	NativeDerived    bool
	AllowsDiagnostic bool
	Verdict          string
}

type DiagnosticDelta struct {
	Name                       string
	Expression                 string
	Level                      string
	RequiresPoleObservableSeal bool
	RequiresTreeProxy          bool
	NumericValueAssigned       bool
	NativeCorrectionTheorem    bool
	IndependentPrediction      bool
	MeasuresProxyToPoleGapOnly bool
	Verdict                    string
}

type LayerAssignmentWarning struct {
	TotalCorrectionOnly              bool
	CannotAssignPiecesWithoutPackage bool
	Layers                           []string
	LayerCount                       int
	ExplanatoryPackageRequired       bool
	Verdict                          string
}

type NonFitFirewall struct {
	FittedFromExternalMassIsDerivedTheorem bool
	ExternalObservableMeasuresGap          bool
	ExternalObservableExplainsGap          bool
	DiagnosticDeltaIsTreeToPoleTheorem     bool
	DiagnosticDeltaIsPrediction            bool
	Verdict                                string
}

type ExplanatoryCorrectionPackage struct {
	Labels      []string
	Count       int
	AllRequired bool
	Native      bool
	Verdict     string
}

type Analysis struct {
	Gate744      Gate744Inheritance
	Observable   PoleObservableSeal
	Delta        DiagnosticDelta
	LayerWarning LayerAssignmentWarning
	NonFit       NonFitFirewall
	Required     ExplanatoryCorrectionPackage
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
	g744, err := gate744.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate744 inheritance unavailable: %w", err)
	}
	inherited := buildGate744Inheritance(g744)
	observable := buildPoleObservableSeal()
	delta := buildDiagnosticDelta()
	warning := buildLayerAssignmentWarning()
	nonfit := buildNonFitFirewall()
	required := buildExplanatoryCorrectionPackage()
	truth := "Gate 745 defines PoleMassObservableSeal and the Level-1C diagnostic delta form Delta_pole_diag=m_H_pole_external-m_H_tree_proxy. The delta may measure the tree-proxy to pole gap when supplied externally, but it neither explains the gap nor becomes a native tree-to-pole correction theorem."
	return Analysis{Gate744: inherited, Observable: observable, Delta: delta, LayerWarning: warning, NonFit: nonfit, Required: required, Truth: truth}, nil
}

func buildGate744Inheritance(g gate744.Analysis) Gate744Inheritance {
	return Gate744Inheritance{
		Inherited:               g.Gate743.Inherited && g.DeltaPole.Name == "Delta_pole" && !g.DeltaPole.ValueAssigned && g.Decomposition.Count == 6 && g.Decomposition.AllRequired && g.NonFit.ExternalDiagnosticAllowed && !g.NonFit.ExternalDiagnosticIsPrediction && g.Classification.Level1CAllowed && !g.Classification.Level2Allowed,
		TreeProxyGeV:            g.Gate743.TreeProxyGeV,
		DeltaPoleObject:         g.DeltaPole.Name,
		DeltaPoleKeptSymbolic:   !g.DeltaPole.ValueAssigned,
		CorrectionLayerCount:    g.Decomposition.Count,
		HasLayeredCorrection:    g.Decomposition.Count == 6 && g.Decomposition.AllRequired,
		Level1CAllowed:          g.Classification.Level1CAllowed,
		Level2Allowed:           g.Classification.Level2Allowed,
		NonFitFirewallPreserved: !g.NonFit.ObservedMinusProxyIsDerivedTheorem && !g.NonFit.ExternalDiagnosticIsPrediction && g.NonFit.DeltaPoleKeptLayered,
		Verdict:                 StatusGate744PoleCorrectionLayerInherited,
	}
}

func buildPoleObservableSeal() PoleObservableSeal {
	return PoleObservableSeal{
		Name:             "PoleMassObservableSeal",
		Object:           "m_H_pole_external",
		ExternalInput:    true,
		ValueSupplied:    false,
		Unit:             "GeV",
		NativeDerived:    false,
		AllowsDiagnostic: true,
		Verdict:          StatusPoleObservableSealDefined,
	}
}

func buildDiagnosticDelta() DiagnosticDelta {
	return DiagnosticDelta{
		Name:                       "Delta_pole_diag",
		Expression:                 "m_H_pole_external - m_H_tree_proxy",
		Level:                      "Level-1C diagnostic comparison",
		RequiresPoleObservableSeal: true,
		RequiresTreeProxy:          true,
		NumericValueAssigned:       false,
		NativeCorrectionTheorem:    false,
		IndependentPrediction:      false,
		MeasuresProxyToPoleGapOnly: true,
		Verdict:                    strings.Join([]string{StatusLevel1CDiagnosticDeltaFormDefined, StatusLevel1CDiagnosticDeltaAllowedExternalPole, StatusDeltaPoleDiagMeasuresGapOnly}, "; "),
	}
}

func buildLayerAssignmentWarning() LayerAssignmentWarning {
	layers := append([]string{}, gate744.CorrectionLayerNames...)
	return LayerAssignmentWarning{
		TotalCorrectionOnly:              true,
		CannotAssignPiecesWithoutPackage: true,
		Layers:                           layers,
		LayerCount:                       len(layers),
		ExplanatoryPackageRequired:       true,
		Verdict:                          StatusLayerAssignmentWarningRecorded,
	}
}

func buildNonFitFirewall() NonFitFirewall {
	return NonFitFirewall{
		FittedFromExternalMassIsDerivedTheorem: false,
		ExternalObservableMeasuresGap:          true,
		ExternalObservableExplainsGap:          false,
		DiagnosticDeltaIsTreeToPoleTheorem:     false,
		DiagnosticDeltaIsPrediction:            false,
		Verdict: strings.Join([]string{
			StatusNonFitFirewallEnforced,
			StatusExternalPoleObservableNotASHADerived,
			StatusDiagnosticDeltaNotTreeToPoleTheorem,
			StatusNoNativeRGThresholdMatchingTheorem,
			StatusNoNativeTopYukawaOrGaugeInputTheorem,
			StatusNoIndependentHiggsPoleMassPrediction,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
		}, "; "),
	}
}

func buildExplanatoryCorrectionPackage() ExplanatoryCorrectionPackage {
	labels := append([]string{}, RequiredExplanatoryCorrectionPackageLabels...)
	return ExplanatoryCorrectionPackage{
		Labels:      labels,
		Count:       len(labels),
		AllRequired: true,
		Native:      false,
		Verdict:     StatusRequiredExplanatoryCorrectionPackageRecorded,
	}
}

func Statuses() []string {
	return []string{
		StatusGate744PoleCorrectionLayerInherited,
		StatusPoleObservableSealDefined,
		StatusLevel1CDiagnosticDeltaFormDefined,
		StatusLayerAssignmentWarningRecorded,
		StatusNonFitFirewallEnforced,
		StatusRequiredExplanatoryCorrectionPackageRecorded,
		StatusLevel1CDiagnosticDeltaAllowedExternalPole,
		StatusDeltaPoleDiagMeasuresGapOnly,
		StatusExternalPoleObservableNotASHADerived,
		StatusDiagnosticDeltaNotTreeToPoleTheorem,
		StatusNoNativeRGThresholdMatchingTheorem,
		StatusNoNativeTopYukawaOrGaugeInputTheorem,
		StatusNoIndependentHiggsPoleMassPrediction,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate745Boundary,
	}
}

func FormatGate744(x Gate744Inheritance) string {
	return fmt.Sprintf("inherited=%t treeProxy=%.17g delta=%s symbolic=%t layerCount=%d layered=%t level1C=%t level2=%t nonFit=%t verdict=%q", x.Inherited, x.TreeProxyGeV, x.DeltaPoleObject, x.DeltaPoleKeptSymbolic, x.CorrectionLayerCount, x.HasLayeredCorrection, x.Level1CAllowed, x.Level2Allowed, x.NonFitFirewallPreserved, x.Verdict)
}

func FormatObservable(x PoleObservableSeal) string {
	return fmt.Sprintf("name=%s object=%s external=%t valueSupplied=%t unit=%s native=%t allowsDiagnostic=%t verdict=%q", x.Name, x.Object, x.ExternalInput, x.ValueSupplied, x.Unit, x.NativeDerived, x.AllowsDiagnostic, x.Verdict)
}

func FormatDelta(x DiagnosticDelta) string {
	return fmt.Sprintf("name=%s expression=%q level=%q requiresObservable=%t requiresTreeProxy=%t valueAssigned=%t nativeTheorem=%t prediction=%t gapOnly=%t verdict=%q", x.Name, x.Expression, x.Level, x.RequiresPoleObservableSeal, x.RequiresTreeProxy, x.NumericValueAssigned, x.NativeCorrectionTheorem, x.IndependentPrediction, x.MeasuresProxyToPoleGapOnly, x.Verdict)
}

func FormatLayerWarning(x LayerAssignmentWarning) string {
	return fmt.Sprintf("totalOnly=%t cannotAssign=%t layerCount=%d packageRequired=%t layers=[%s] verdict=%q", x.TotalCorrectionOnly, x.CannotAssignPiecesWithoutPackage, x.LayerCount, x.ExplanatoryPackageRequired, strings.Join(x.Layers, ", "), x.Verdict)
}

func FormatNonFit(x NonFitFirewall) string {
	return fmt.Sprintf("fitDerived=%t measuresGap=%t explainsGap=%t treeToPoleTheorem=%t prediction=%t verdict=%q", x.FittedFromExternalMassIsDerivedTheorem, x.ExternalObservableMeasuresGap, x.ExternalObservableExplainsGap, x.DiagnosticDeltaIsTreeToPoleTheorem, x.DiagnosticDeltaIsPrediction, x.Verdict)
}

func FormatRequired(x ExplanatoryCorrectionPackage) string {
	return fmt.Sprintf("count=%d allRequired=%t native=%t labels=[%s] verdict=%q", x.Count, x.AllRequired, x.Native, strings.Join(x.Labels, ", "), x.Verdict)
}

func NearlyEqual(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
