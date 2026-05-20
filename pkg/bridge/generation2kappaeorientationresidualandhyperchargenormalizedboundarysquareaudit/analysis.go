// Package generation2kappaeorientationresidualandhyperchargenormalizedboundarysquareaudit implements
// Gate 747: Kappa_e Orientation Residual and Hypercharge-Normalized Boundary-Square Audit.
//
// Gate 746 showed that kappa_e is close to the flavor-orientation candidate
// sin^2(theta13)/4-J_CKM but not exact. Gate 747 audits whether the residual is
// source-typed by the mature 5/3 gauge/hypercharge normalization acting on the
// second-order boundary split S_split^2. The result is a strong source-type
// compression only, not a native flavor theorem or scalar-runtime derivation.
package generation2kappaeorientationresidualandhyperchargenormalizedboundarysquareaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate746 "github.com/bagherbal/asha-engine/pkg/bridge/generation2flavorwalldeficitkappaesourcetypeandscalarbridgedependencyaudit"
)

const (
	AuditID = "GATE747-KAPPA-E-ORIENTATION-RESIDUAL-HYPERCHARGE-NORMALIZED-BOUNDARY-SQUARE-AUDIT"

	StatusGate746KappaESourceAuditInherited          = "PASS_GATE746_KAPPA_E_SOURCE_AUDIT_INHERITED"
	StatusDeltaKappaEOverSSplitSquaredComputed       = "PASS_DELTA_KAPPA_E_OVER_S_SPLIT_SQUARED_COMPUTED"
	StatusTypedRatioCandidatesAudited                = "PASS_TYPED_RATIO_CANDIDATES_AUDITED"
	StatusHyperchargeBoundarySquareCorrectionDefined = "PASS_HYPERCHARGE_BOUNDARY_SQUARE_CORRECTION_DEFINED"
	StatusScalarRuntimeReplacementTested             = "PASS_SCALAR_RUNTIME_REPLACEMENT_TESTED"
	StatusSourceTypeInterpretationRecorded           = "PASS_SOURCE_TYPE_INTERPRETATION_RECORDED"
	StatusNoncircularityFirewallAudited              = "PASS_NONCIRCULARITY_FIREWALL_AUDITED"
	StatusPhysicalFirewallsEnforced                  = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusKappaEResidualSecondOrderBoundarySplitScale    = "CONDITIONAL_SUPPORT_KAPPA_E_RESIDUAL_IS_SECOND_ORDER_BOUNDARY_SPLIT_SCALE"
	StatusNegativeFiveThirdsBestTypedResidualCoefficient = "CONDITIONAL_SUPPORT_NEGATIVE_FIVE_OVER_THREE_IS_BEST_TYPED_RESIDUAL_COEFFICIENT"
	StatusKappaEOrientationHyperBoundaryCorrectionForm   = "CONDITIONAL_SUPPORT_KAPPA_E_HAS_ORIENTATION_PLUS_HYPERCHARGE_BOUNDARY_CORRECTION_FORM"
	StatusHyperBoundaryStronglyImprovesKappaESourceType  = "CONDITIONAL_SUPPORT_HYPERCHARGE_BOUNDARY_SQUARE_CORRECTION_STRONGLY_IMPROVES_KAPPA_E_SOURCE_TYPE"

	StatusCorrectionNotExact                                      = "FAILED_ROUTE_CORRECTION_NOT_EXACT"
	StatusNoNativeReasonFlavorResidualEqualsMinusFiveThirdsSquare = "FAILED_ROUTE_NO_NATIVE_REASON_FLAVOR_RESIDUAL_EQUALS_MINUS_FIVE_THIRDS_BOUNDARY_SQUARE"
	StatusNoNativeFlavorDeficitTheorem                            = "FAILED_ROUTE_NO_NATIVE_FLAVOR_DEFICIT_THEOREM"
	StatusNoNativePMNSOrCKMTheorem                                = "FAILED_ROUTE_NO_NATIVE_PMNS_OR_CKM_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem                     = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusNoIndependentScalarRuntimeTheorem                       = "FAILED_ROUTE_NO_INDEPENDENT_SCALAR_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem                            = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusGate747Boundary                                         = "FIREWALL_PRESERVED_GATE747_KAPPA_E_HYPERCHARGE_BOUNDARY_SQUARE_BOUNDARY"
)

const (
	HyperchargeGaugeCoefficient = 5.0 / 3.0
)

type Gate746Inheritance struct {
	Inherited           bool
	KappaEActiveInput   bool
	OrientationClose    bool
	OrientationNotExact bool
	FlavorFirewallKept  bool
	Verdict             string
}

type RatioCandidate struct {
	Name        string
	Value       float64
	AbsDistance float64
	Typed       bool
	Selected    bool
	Reason      string
}

type ResidualRatioAudit struct {
	KappaE                 float64
	KappaEOrient           float64
	DeltaKappaE            float64
	SSplit                 float64
	SSplitSquared          float64
	Ratio                  float64
	Candidates             []RatioCandidate
	BestCandidate          string
	CloseToMinusFiveThirds bool
	Verdict                string
}

type HyperchargeBoundarySquareCorrection struct {
	Formula                 string
	KappaEOrient            float64
	SSplitSquared           float64
	Coefficient             float64
	Correction              float64
	KappaEHyperBoundary     float64
	KappaE                  float64
	ResidualAfterCorrection float64
	CompressionFactor       float64
	CorrectionNotExact      bool
	Verdict                 string
}

type ScalarRuntimeReplacementTest struct {
	FormulaOrientOnly         string
	FormulaHyperBoundary      string
	SSplit                    float64
	P_K7                      float64
	KappaE                    float64
	KappaEOrient              float64
	KappaEHyperBoundary       float64
	FExact                    float64
	FOrient                   float64
	FHyperBoundary            float64
	RuntimeExact              float64
	RuntimeOrient             float64
	RuntimeHyperBoundary      float64
	RuntimeOrientShift        float64
	RuntimeHyperBoundaryShift float64
	ImprovementFactor         float64
	ReplacementNotNative      bool
	Verdict                   string
}

type SourceTypeInterpretation struct {
	Expression     string
	Terms          []string
	Interpretation string
	Verdict        string
}

type NoncircularityFirewall struct {
	Theta13EmpiricalBridgeInput     bool
	JCKMEmpiricalBridgeInput        bool
	FiveThirdsMatureButUncoupled    bool
	SSplitBoundaryNotFlavorOperator bool
	DerivesFlavorTheorem            bool
	DerivesScalarRuntime            bool
	DerivesHiggsMass                bool
	DerivesYukawa                   bool
	Verdict                         string
}

type Analysis struct {
	Gate746     Gate746Inheritance
	Ratio       ResidualRatioAudit
	Correction  HyperchargeBoundarySquareCorrection
	Replacement ScalarRuntimeReplacementTest
	SourceType  SourceTypeInterpretation
	Firewall    NoncircularityFirewall
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
	g746, err := gate746.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate746 inheritance unavailable: %w", err)
	}
	inherit := buildGate746Inheritance(g746)
	ratio := buildResidualRatioAudit(g746)
	correction := buildHyperchargeBoundaryCorrection(ratio)
	replacement := buildReplacementTest(g746, correction)
	sourceType := buildSourceTypeInterpretation()
	firewall := buildNoncircularityFirewall()
	truth := "Gate 747 source-types the kappa_e orientation residual as a second-order boundary-square correction. The ratio Delta_kappa_e/S_split^2 is close to -5/3, and kappa_e≈sin^2(theta13)/4-J_CKM-(5/3)S_split^2 improves the scalar-runtime replacement by roughly 340x relative to orientation-only. The correction is not exact and no native theorem couples the 5/3 gauge normalization to the flavor residual; kappa_e remains a bridge-layer flavor deficit seal."
	return Analysis{Gate746: inherit, Ratio: ratio, Correction: correction, Replacement: replacement, SourceType: sourceType, Firewall: firewall, Truth: truth}, nil
}

func buildGate746Inheritance(g gate746.Analysis) Gate746Inheritance {
	return Gate746Inheritance{
		Inherited:           g.Gate745.Inherited && g.Dependency.StructurallyActive && g.Orientation.CloseButNotExact && g.Firewall.KappaEStillBridgeSeal,
		KappaEActiveInput:   g.Dependency.StructurallyActive && g.Dependency.AppearsInBoundaryPolynomial && g.Dependency.AppearsInRuntimeTransport,
		OrientationClose:    g.Orientation.CloseButNotExact && !g.Orientation.NativeTheorem,
		OrientationNotExact: math.Abs(g.Orientation.DeltaKappaE) > 1e-9,
		FlavorFirewallKept:  !g.Firewall.DerivesPMNS && !g.Firewall.DerivesCKM && !g.Firewall.DerivesYukawaEigenvalues,
		Verdict:             StatusGate746KappaESourceAuditInherited,
	}
}

func buildResidualRatioAudit(g gate746.Analysis) ResidualRatioAudit {
	S := g.Replacement.SSplit
	delta := g.Orientation.DeltaKappaE
	S2 := S * S
	ratio := delta / S2
	cands := []RatioCandidate{
		candidate("-5/3 hypercharge/gauge normalization", -HyperchargeGaugeCoefficient, ratio, true, true, "mature ASHA gauge/hypercharge normalization coefficient"),
		candidate("-2 generic double-boundary control", -2.0, ratio, true, false, "typed control, but less close and less specific"),
		candidate("-3/2 simple nearby control", -1.5, ratio, true, false, "nearby simple control, but not the closest active coefficient"),
		candidate("-phi golden control", -(1+math.Sqrt(5))/2, ratio, false, false, "rejected without native fivefold carrier in this lane"),
	}
	return ResidualRatioAudit{
		KappaE:                 g.Orientation.KappaE,
		KappaEOrient:           g.Orientation.KappaEOrient,
		DeltaKappaE:            delta,
		SSplit:                 S,
		SSplitSquared:          S2,
		Ratio:                  ratio,
		Candidates:             cands,
		BestCandidate:          "-5/3 hypercharge/gauge normalization",
		CloseToMinusFiveThirds: math.Abs(ratio+HyperchargeGaugeCoefficient) < 0.01,
		Verdict: strings.Join([]string{
			StatusDeltaKappaEOverSSplitSquaredComputed,
			StatusTypedRatioCandidatesAudited,
			StatusKappaEResidualSecondOrderBoundarySplitScale,
			StatusNegativeFiveThirdsBestTypedResidualCoefficient,
		}, "; "),
	}
}

func candidate(name string, value, ratio float64, typed, selected bool, reason string) RatioCandidate {
	return RatioCandidate{Name: name, Value: value, AbsDistance: math.Abs(ratio - value), Typed: typed, Selected: selected, Reason: reason}
}

func buildHyperchargeBoundaryCorrection(r ResidualRatioAudit) HyperchargeBoundarySquareCorrection {
	corr := -HyperchargeGaugeCoefficient * r.SSplitSquared
	hyper := r.KappaEOrient + corr
	resid := r.KappaE - hyper
	return HyperchargeBoundarySquareCorrection{
		Formula:                 "kappa_e_hyper_boundary = kappa_e_orient - (5/3)S_split^2",
		KappaEOrient:            r.KappaEOrient,
		SSplitSquared:           r.SSplitSquared,
		Coefficient:             -HyperchargeGaugeCoefficient,
		Correction:              corr,
		KappaEHyperBoundary:     hyper,
		KappaE:                  r.KappaE,
		ResidualAfterCorrection: resid,
		CompressionFactor:       math.Abs(r.DeltaKappaE) / math.Abs(resid),
		CorrectionNotExact:      math.Abs(resid) > 1e-10,
		Verdict: strings.Join([]string{
			StatusHyperchargeBoundarySquareCorrectionDefined,
			StatusKappaEOrientationHyperBoundaryCorrectionForm,
			StatusCorrectionNotExact,
		}, "; "),
	}
}

func buildReplacementTest(g gate746.Analysis, c HyperchargeBoundarySquareCorrection) ScalarRuntimeReplacementTest {
	p := g.Replacement.P_K7
	S := g.Replacement.SSplit
	fExact := g.Replacement.FExact
	fOrient := g.Replacement.FOrient
	fHyper := rawPolynomial(p, S, c.KappaEHyperBoundary)
	absLambda := g.Replacement.W3Exact - g.Replacement.FExact
	wHyper := absLambda + fHyper
	runtimeExact := g.Replacement.RuntimeExactKappaE
	runtimeOrient := g.Replacement.RuntimeOrientKappaE
	runtimeHyper := g.Replacement.LambdaProxy * (1 + g.Replacement.L*(1-wHyper+c.KappaEHyperBoundary))
	orientShift := runtimeOrient - runtimeExact
	hyperShift := runtimeHyper - runtimeExact
	improvement := math.Abs(orientShift) / math.Abs(hyperShift)
	return ScalarRuntimeReplacementTest{
		FormulaOrientOnly:         "lambda_runtime≈lambda_proxy[1+L(1-W_3(kappa_e_orient)+kappa_e_orient)]",
		FormulaHyperBoundary:      "lambda_runtime≈lambda_proxy[1+L(1-W_3(kappa_e_orient-(5/3)S^2)+kappa_e_orient-(5/3)S^2)]",
		SSplit:                    S,
		P_K7:                      p,
		KappaE:                    c.KappaE,
		KappaEOrient:              c.KappaEOrient,
		KappaEHyperBoundary:       c.KappaEHyperBoundary,
		FExact:                    fExact,
		FOrient:                   fOrient,
		FHyperBoundary:            fHyper,
		RuntimeExact:              runtimeExact,
		RuntimeOrient:             runtimeOrient,
		RuntimeHyperBoundary:      runtimeHyper,
		RuntimeOrientShift:        orientShift,
		RuntimeHyperBoundaryShift: hyperShift,
		ImprovementFactor:         improvement,
		ReplacementNotNative:      true,
		Verdict: strings.Join([]string{
			StatusScalarRuntimeReplacementTested,
			StatusHyperBoundaryStronglyImprovesKappaESourceType,
			StatusCorrectionNotExact,
		}, "; "),
	}
}

func rawPolynomial(p, S, kappaE float64) float64 {
	return p*S + kappaE*p*S*S - 2*p*p*S*S*S
}

func buildSourceTypeInterpretation() SourceTypeInterpretation {
	terms := []string{
		"sin^2(theta13)/4: PMNS reactor leakage term",
		"-J_CKM: CKM orientation correction",
		"-(5/3)S_split^2: second-order hypercharge-normalized boundary split correction",
	}
	return SourceTypeInterpretation{
		Expression:     "kappa_e ≈ sin^2(theta13)/4 - J_CKM - (5/3)S_split^2",
		Terms:          terms,
		Interpretation: "orientation plus hypercharge-normalized boundary-square correction",
		Verdict:        StatusSourceTypeInterpretationRecorded,
	}
}

func buildNoncircularityFirewall() NoncircularityFirewall {
	return NoncircularityFirewall{
		Theta13EmpiricalBridgeInput:     true,
		JCKMEmpiricalBridgeInput:        true,
		FiveThirdsMatureButUncoupled:    true,
		SSplitBoundaryNotFlavorOperator: true,
		DerivesFlavorTheorem:            false,
		DerivesScalarRuntime:            false,
		DerivesHiggsMass:                false,
		DerivesYukawa:                   false,
		Verdict: strings.Join([]string{
			StatusNoncircularityFirewallAudited,
			StatusPhysicalFirewallsEnforced,
			StatusNoNativeReasonFlavorResidualEqualsMinusFiveThirdsSquare,
			StatusNoNativeFlavorDeficitTheorem,
			StatusNoNativePMNSOrCKMTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusNoIndependentScalarRuntimeTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusGate747Boundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate746KappaESourceAuditInherited,
		StatusDeltaKappaEOverSSplitSquaredComputed,
		StatusTypedRatioCandidatesAudited,
		StatusHyperchargeBoundarySquareCorrectionDefined,
		StatusScalarRuntimeReplacementTested,
		StatusSourceTypeInterpretationRecorded,
		StatusNoncircularityFirewallAudited,
		StatusPhysicalFirewallsEnforced,
		StatusKappaEResidualSecondOrderBoundarySplitScale,
		StatusNegativeFiveThirdsBestTypedResidualCoefficient,
		StatusKappaEOrientationHyperBoundaryCorrectionForm,
		StatusHyperBoundaryStronglyImprovesKappaESourceType,
		StatusCorrectionNotExact,
		StatusNoNativeReasonFlavorResidualEqualsMinusFiveThirdsSquare,
		StatusNoNativeFlavorDeficitTheorem,
		StatusNoNativePMNSOrCKMTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusNoIndependentScalarRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusGate747Boundary,
	}
}

func FormatGate746(x Gate746Inheritance) string {
	return fmt.Sprintf("inherited=%t active=%t orientClose=%t orientNotExact=%t flavorFirewall=%t verdict=%q", x.Inherited, x.KappaEActiveInput, x.OrientationClose, x.OrientationNotExact, x.FlavorFirewallKept, x.Verdict)
}

func FormatRatio(x ResidualRatioAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, fmt.Sprintf("%s=%.17g dist=%.17g typed=%t selected=%t", c.Name, c.Value, c.AbsDistance, c.Typed, c.Selected))
	}
	return fmt.Sprintf("kappaE=%.17g orient=%.17g delta=%.17g S=%.17g S2=%.17g ratio=%.17g best=%q close=%t candidates=[%s] verdict=%q", x.KappaE, x.KappaEOrient, x.DeltaKappaE, x.SSplit, x.SSplitSquared, x.Ratio, x.BestCandidate, x.CloseToMinusFiveThirds, strings.Join(parts, "; "), x.Verdict)
}

func FormatCorrection(x HyperchargeBoundarySquareCorrection) string {
	return fmt.Sprintf("formula=%q orient=%.17g S2=%.17g coeff=%.17g correction=%.17g hyper=%.17g kappaE=%.17g residual=%.17g compression=%.17g notExact=%t verdict=%q", x.Formula, x.KappaEOrient, x.SSplitSquared, x.Coefficient, x.Correction, x.KappaEHyperBoundary, x.KappaE, x.ResidualAfterCorrection, x.CompressionFactor, x.CorrectionNotExact, x.Verdict)
}

func FormatReplacement(x ScalarRuntimeReplacementTest) string {
	return fmt.Sprintf("S=%.17g p=%.17g kappaE=%.17g orient=%.17g hyper=%.17g FExact=%.17g FOrient=%.17g FHyper=%.17g runtimeExact=%.17g runtimeOrient=%.17g runtimeHyper=%.17g orientShift=%.17g hyperShift=%.17g improvement=%.17g notNative=%t verdict=%q", x.SSplit, x.P_K7, x.KappaE, x.KappaEOrient, x.KappaEHyperBoundary, x.FExact, x.FOrient, x.FHyperBoundary, x.RuntimeExact, x.RuntimeOrient, x.RuntimeHyperBoundary, x.RuntimeOrientShift, x.RuntimeHyperBoundaryShift, x.ImprovementFactor, x.ReplacementNotNative, x.Verdict)
}

func FormatSourceType(x SourceTypeInterpretation) string {
	return fmt.Sprintf("expression=%q terms=[%s] interpretation=%q verdict=%q", x.Expression, strings.Join(x.Terms, "; "), x.Interpretation, x.Verdict)
}

func FormatFirewall(x NoncircularityFirewall) string {
	return fmt.Sprintf("theta13Bridge=%t jckmBridge=%t fiveThirdsMatureButUncoupled=%t SBoundaryNotFlavor=%t derivesFlavor=%t derivesRuntime=%t derivesHiggs=%t derivesYukawa=%t verdict=%q", x.Theta13EmpiricalBridgeInput, x.JCKMEmpiricalBridgeInput, x.FiveThirdsMatureButUncoupled, x.SSplitBoundaryNotFlavorOperator, x.DerivesFlavorTheorem, x.DerivesScalarRuntime, x.DerivesHiggsMass, x.DerivesYukawa, x.Verdict)
}

func Near(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
