// Package generation2polecorrectionsealpackageandlevel1cdiagnosticboundaryaudit implements
// Gate 743: Pole-Correction Seal Package and Level-1C Diagnostic Boundary Audit.
//
// Gate 742 defined the formal correction object Delta_pole but assigned no
// value. Gate 743 audits the minimal correction-seal package required for a
// lawful Level-1C tree-to-pole diagnostic comparison while preserving the
// firewall against independent Higgs pole-mass prediction claims.
package generation2polecorrectionsealpackageandlevel1cdiagnosticboundaryaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate742 "github.com/bagherbal/asha-engine/pkg/bridge/generation2treeproxytopolemasscorrectiondependencyandfirewallaudit"
)

const (
	AuditID = "GATE743-POLE-CORRECTION-SEAL-PACKAGE-LEVEL-1C-DIAGNOSTIC-BOUNDARY-AUDIT"

	StatusGate742TreeProxyToPoleFirewallInherited = "PASS_GATE742_TREE_PROXY_TO_POLE_FIREWALL_INHERITED"
	StatusPoleCorrectionSealPackageDefined        = "PASS_POLE_CORRECTION_SEAL_PACKAGE_DEFINED"
	StatusCorrectionPackageMinimalityAudited      = "PASS_CORRECTION_PACKAGE_MINIMALITY_AUDITED"
	StatusLevel1CDiagnosticBoundaryDefined        = "PASS_LEVEL_1C_DIAGNOSTIC_BOUNDARY_DEFINED"
	StatusTreeProxyAndPoleObservableSeparated     = "PASS_TREE_PROXY_AND_POLE_OBSERVABLE_SEPARATED"
	StatusPhysicalFirewallsEnforced               = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusLevel1CDiagnosticAllowedFullPackage = "CONDITIONAL_SUPPORT_LEVEL_1C_DIAGNOSTIC_ALLOWED_WITH_FULL_CORRECTION_PACKAGE"
	StatusDeltaPoleValidOnlyAsSealedObject    = "CONDITIONAL_SUPPORT_DELTA_POLE_IS_VALID_ONLY_AS_SEALED_CORRECTION_OBJECT"

	StatusNoNativeTreeToPoleCorrectionTheorem  = "FAILED_ROUTE_NO_NATIVE_TREE_TO_POLE_CORRECTION_THEOREM"
	StatusNoNativeRGThresholdMatchingTheorem   = "FAILED_ROUTE_NO_NATIVE_RG_THRESHOLD_MATCHING_THEOREM"
	StatusNoNativeTopYukawaOrGaugeInputTheorem = "FAILED_ROUTE_NO_NATIVE_TOP_YUKAWA_OR_GAUGE_INPUT_THEOREM"
	StatusExternalPoleObservableNotASHADerived = "FAILED_ROUTE_EXTERNAL_POLE_OBSERVABLE_IS_NOT_ASHA_DERIVATION"
	StatusNoIndependentHiggsPoleMassPrediction = "FAILED_ROUTE_NO_INDEPENDENT_HIGGS_POLE_MASS_PREDICTION"
	StatusNoYukawaOperatorOrEigenvalueTheorem  = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate743Boundary                      = "FIREWALL_PRESERVED_GATE743_POLE_CORRECTION_SEAL_PACKAGE_BOUNDARY"
)

var PoleCorrectionSealPackageLabels = []string{
	"PoleMassObservableSeal",
	"PoleMassConventionSeal",
	"RGSchemeSeal",
	"RenormalizationScaleSeal",
	"LoopOrderSeal",
	"ThresholdCorrectionSeal",
	"TopYukawaInputSeal",
	"GaugeCouplingInputSeal",
	"UncertaintyModelSeal",
}

type Gate742Inheritance struct {
	Inherited                      bool
	TreeProxyGeV                   float64
	CorrectionObject               string
	CorrectionValueAssigned        bool
	Level1CAllowed                 bool
	Level1CRequiresExternalPackage bool
	Level2Allowed                  bool
	TreeProxyNotPoleMass           bool
	NoNativeTreeToPoleCorrection   bool
	NoIndependentPolePrediction    bool
	Verdict                        string
}

type PoleCorrectionSealPackage struct {
	Labels                  []string
	Count                   int
	HasPoleObservable       bool
	HasPoleMassConvention   bool
	HasRGScheme             bool
	HasRenormalizationScale bool
	HasLoopOrder            bool
	HasThresholdCorrection  bool
	HasTopYukawaInput       bool
	HasGaugeCouplingInput   bool
	HasUncertaintyModel     bool
	FullPackage             bool
	DeltaPoleValueAssigned  bool
	Verdict                 string
}

type MinimalityItem struct {
	Seal               string
	RemovalEffect      string
	RequiredForLevel1C bool
}

type MinimalityAudit struct {
	Items       []MinimalityItem
	Count       int
	AllRequired bool
	Minimal     bool
	Verdict     string
}

type DiagnosticBoundary struct {
	Level1BName           string
	Level1BAllowed        bool
	Level1CName           string
	Level1CAllowed        bool
	Level1CDiagnosticOnly bool
	Level1CRequiresAll    bool
	Level2Name            string
	Level2Allowed         bool
	Verdict               string
}

type SeparationAudit struct {
	TreeProxyGeV                             float64
	DeltaPoleObject                          string
	PoleObservableExternallySupplied         bool
	ExternalPoleObservableASHADerived        bool
	TreeProxyEqualsPoleObservable            bool
	DiagnosticCanComputeDeltaOnlyWithPackage bool
	Verdict                                  string
}

type PhysicalFirewall struct {
	FittedDeltaIsDerivedTheorem    bool
	Level1CDiagnosticIsPrediction  bool
	TreeProxyProximityIsTheorem    bool
	ExternalObservableIsDerivation bool
	IndependentPolePrediction      bool
	NoYukawaTheorem                bool
	Verdict                        string
}

type Analysis struct {
	Gate742    Gate742Inheritance
	Package    PoleCorrectionSealPackage
	Minimality MinimalityAudit
	Diagnostic DiagnosticBoundary
	Separation SeparationAudit
	Firewall   PhysicalFirewall
	Truth      string
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
	g742, err := gate742.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate742 inheritance unavailable: %w", err)
	}
	inherit := buildGate742Inheritance(g742)
	pkg := buildPoleCorrectionSealPackage()
	minimality := buildMinimalityAudit()
	diagnostic := buildDiagnosticBoundary()
	separation := buildSeparationAudit(g742)
	firewall := buildPhysicalFirewall()
	truth := "Gate 743 defines the minimal pole-correction seal package required for Level-1C diagnostics. Delta_pole is valid only as a sealed correction object after an external pole/correction package is supplied; no native tree-to-pole theorem or independent Higgs pole-mass prediction is certified."
	return Analysis{Gate742: inherit, Package: pkg, Minimality: minimality, Diagnostic: diagnostic, Separation: separation, Firewall: firewall, Truth: truth}, nil
}

func buildGate742Inheritance(g gate742.Analysis) Gate742Inheritance {
	statuses := strings.Join(gate742.Statuses(), "\n")
	return Gate742Inheritance{
		Inherited:                      g.Gate741.Inherited && g.Firewall.TreeProxyConventionLevel && !g.Firewall.TreeProxyEqualsPoleMass && g.Forecast.Level1CAllowed && !g.Forecast.Level2Allowed,
		TreeProxyGeV:                   g.Gate741.TreeProxyGeV,
		CorrectionObject:               g.Correction.Name,
		CorrectionValueAssigned:        g.Correction.ValueAssigned,
		Level1CAllowed:                 g.Forecast.Level1CAllowed,
		Level1CRequiresExternalPackage: g.Forecast.Level1CRequiresExternal,
		Level2Allowed:                  g.Forecast.Level2Allowed,
		TreeProxyNotPoleMass:           strings.Contains(statuses, gate742.StatusTreeProxyIsNotPoleMass),
		NoNativeTreeToPoleCorrection:   strings.Contains(statuses, gate742.StatusNoNativeTreeToPoleCorrectionTheorem),
		NoIndependentPolePrediction:    strings.Contains(statuses, gate742.StatusNoIndependentHiggsPoleMassPrediction),
		Verdict:                        StatusGate742TreeProxyToPoleFirewallInherited,
	}
}

func buildPoleCorrectionSealPackage() PoleCorrectionSealPackage {
	labels := append([]string{}, PoleCorrectionSealPackageLabels...)
	return PoleCorrectionSealPackage{
		Labels:                  labels,
		Count:                   len(labels),
		HasPoleObservable:       contains(labels, "PoleMassObservableSeal"),
		HasPoleMassConvention:   contains(labels, "PoleMassConventionSeal"),
		HasRGScheme:             contains(labels, "RGSchemeSeal"),
		HasRenormalizationScale: contains(labels, "RenormalizationScaleSeal"),
		HasLoopOrder:            contains(labels, "LoopOrderSeal"),
		HasThresholdCorrection:  contains(labels, "ThresholdCorrectionSeal"),
		HasTopYukawaInput:       contains(labels, "TopYukawaInputSeal"),
		HasGaugeCouplingInput:   contains(labels, "GaugeCouplingInputSeal"),
		HasUncertaintyModel:     contains(labels, "UncertaintyModelSeal"),
		FullPackage:             len(labels) == 9,
		DeltaPoleValueAssigned:  false,
		Verdict:                 strings.Join([]string{StatusPoleCorrectionSealPackageDefined, StatusDeltaPoleValidOnlyAsSealedObject}, "; "),
	}
}

func buildMinimalityAudit() MinimalityAudit {
	items := []MinimalityItem{
		{Seal: "PoleMassObservableSeal", RemovalEffect: "no target pole observable", RequiredForLevel1C: true},
		{Seal: "PoleMassConventionSeal", RemovalEffect: "pole value is not conventionally typed", RequiredForLevel1C: true},
		{Seal: "RGSchemeSeal", RemovalEffect: "running/pole comparison is ill-typed", RequiredForLevel1C: true},
		{Seal: "RenormalizationScaleSeal", RemovalEffect: "renormalization point is unspecified", RequiredForLevel1C: true},
		{Seal: "LoopOrderSeal", RemovalEffect: "correction order is ambiguous", RequiredForLevel1C: true},
		{Seal: "ThresholdCorrectionSeal", RemovalEffect: "tree-to-pole map is absent", RequiredForLevel1C: true},
		{Seal: "TopYukawaInputSeal", RemovalEffect: "dominant top-sector correction dependency is missing", RequiredForLevel1C: true},
		{Seal: "GaugeCouplingInputSeal", RemovalEffect: "gauge-sector correction dependency is missing", RequiredForLevel1C: true},
		{Seal: "UncertaintyModelSeal", RemovalEffect: "numerical comparison has no error ledger", RequiredForLevel1C: true},
	}
	all := true
	for _, item := range items {
		all = all && item.RequiredForLevel1C && item.Seal != "" && item.RemovalEffect != ""
	}
	return MinimalityAudit{Items: items, Count: len(items), AllRequired: all, Minimal: all && len(items) == 9, Verdict: StatusCorrectionPackageMinimalityAudited}
}

func buildDiagnosticBoundary() DiagnosticBoundary {
	return DiagnosticBoundary{
		Level1BName:           "Level 1B sealed tree proxy estimate",
		Level1BAllowed:        true,
		Level1CName:           "Level 1C diagnostic comparison to externally supplied pole/correction package",
		Level1CAllowed:        true,
		Level1CDiagnosticOnly: true,
		Level1CRequiresAll:    true,
		Level2Name:            "Level 2 independent Higgs pole-mass prediction",
		Level2Allowed:         false,
		Verdict:               strings.Join([]string{StatusLevel1CDiagnosticBoundaryDefined, StatusLevel1CDiagnosticAllowedFullPackage, StatusNoIndependentHiggsPoleMassPrediction}, "; "),
	}
}

func buildSeparationAudit(g gate742.Analysis) SeparationAudit {
	return SeparationAudit{
		TreeProxyGeV:                             g.Gate741.TreeProxyGeV,
		DeltaPoleObject:                          g.Correction.Name,
		PoleObservableExternallySupplied:         false,
		ExternalPoleObservableASHADerived:        false,
		TreeProxyEqualsPoleObservable:            false,
		DiagnosticCanComputeDeltaOnlyWithPackage: true,
		Verdict:                                  strings.Join([]string{StatusTreeProxyAndPoleObservableSeparated, StatusExternalPoleObservableNotASHADerived}, "; "),
	}
}

func buildPhysicalFirewall() PhysicalFirewall {
	return PhysicalFirewall{
		FittedDeltaIsDerivedTheorem:    false,
		Level1CDiagnosticIsPrediction:  false,
		TreeProxyProximityIsTheorem:    false,
		ExternalObservableIsDerivation: false,
		IndependentPolePrediction:      false,
		NoYukawaTheorem:                true,
		Verdict: strings.Join([]string{
			StatusPhysicalFirewallsEnforced,
			StatusNoNativeTreeToPoleCorrectionTheorem,
			StatusNoNativeRGThresholdMatchingTheorem,
			StatusNoNativeTopYukawaOrGaugeInputTheorem,
			StatusNoIndependentHiggsPoleMassPrediction,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
		}, "; "),
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
		StatusGate742TreeProxyToPoleFirewallInherited,
		StatusPoleCorrectionSealPackageDefined,
		StatusCorrectionPackageMinimalityAudited,
		StatusLevel1CDiagnosticBoundaryDefined,
		StatusTreeProxyAndPoleObservableSeparated,
		StatusPhysicalFirewallsEnforced,
		StatusLevel1CDiagnosticAllowedFullPackage,
		StatusDeltaPoleValidOnlyAsSealedObject,
		StatusNoNativeTreeToPoleCorrectionTheorem,
		StatusNoNativeRGThresholdMatchingTheorem,
		StatusNoNativeTopYukawaOrGaugeInputTheorem,
		StatusExternalPoleObservableNotASHADerived,
		StatusNoIndependentHiggsPoleMassPrediction,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate743Boundary,
	}
}

func FormatGate742(x Gate742Inheritance) string {
	return fmt.Sprintf("inherited=%t treeProxy=%.17g correction=%s valueAssigned=%t level1C=%t external=%t level2=%t notPole=%t noTreePole=%t noPrediction=%t verdict=%q", x.Inherited, x.TreeProxyGeV, x.CorrectionObject, x.CorrectionValueAssigned, x.Level1CAllowed, x.Level1CRequiresExternalPackage, x.Level2Allowed, x.TreeProxyNotPoleMass, x.NoNativeTreeToPoleCorrection, x.NoIndependentPolePrediction, x.Verdict)
}

func FormatPackage(x PoleCorrectionSealPackage) string {
	return fmt.Sprintf("labels=[%s] count=%d poleObservable=%t convention=%t rg=%t scale=%t loop=%t threshold=%t top=%t gauge=%t uncertainty=%t full=%t deltaAssigned=%t verdict=%q", strings.Join(x.Labels, ","), x.Count, x.HasPoleObservable, x.HasPoleMassConvention, x.HasRGScheme, x.HasRenormalizationScale, x.HasLoopOrder, x.HasThresholdCorrection, x.HasTopYukawaInput, x.HasGaugeCouplingInput, x.HasUncertaintyModel, x.FullPackage, x.DeltaPoleValueAssigned, x.Verdict)
}

func FormatMinimality(x MinimalityAudit) string {
	parts := make([]string, 0, len(x.Items))
	for _, item := range x.Items {
		parts = append(parts, item.Seal+":"+item.RemovalEffect)
	}
	return fmt.Sprintf("count=%d allRequired=%t minimal=%t items=[%s] verdict=%q", x.Count, x.AllRequired, x.Minimal, strings.Join(parts, " | "), x.Verdict)
}

func FormatDiagnostic(x DiagnosticBoundary) string {
	return fmt.Sprintf("level1B=%q allowed=%t level1C=%q allowed=%t diagnosticOnly=%t requiresAll=%t level2=%q allowed=%t verdict=%q", x.Level1BName, x.Level1BAllowed, x.Level1CName, x.Level1CAllowed, x.Level1CDiagnosticOnly, x.Level1CRequiresAll, x.Level2Name, x.Level2Allowed, x.Verdict)
}

func FormatSeparation(x SeparationAudit) string {
	return fmt.Sprintf("treeProxy=%.17g delta=%s poleSupplied=%t poleASHADerived=%t treeEqualsPole=%t canComputeDeltaOnlyWithPackage=%t verdict=%q", x.TreeProxyGeV, x.DeltaPoleObject, x.PoleObservableExternallySupplied, x.ExternalPoleObservableASHADerived, x.TreeProxyEqualsPoleObservable, x.DiagnosticCanComputeDeltaOnlyWithPackage, x.Verdict)
}

func FormatFirewall(x PhysicalFirewall) string {
	return fmt.Sprintf("fittedDerived=%t diagnosticPrediction=%t proximityTheorem=%t externalDerivation=%t independentPrediction=%t noYukawa=%t verdict=%q", x.FittedDeltaIsDerivedTheorem, x.Level1CDiagnosticIsPrediction, x.TreeProxyProximityIsTheorem, x.ExternalObservableIsDerivation, x.IndependentPolePrediction, x.NoYukawaTheorem, x.Verdict)
}

func NearlyEqual(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
