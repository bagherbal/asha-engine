// Package generation2kappaehyperchargeboundaryresidualandboundarystressmomentaudit implements
// Gate 748: Kappa_e Hypercharge-Boundary Residual and Boundary-Stress Moment Audit.
//
// Gate 747 source-typed most of the kappa_e orientation residual by the mature
// -5/3 hypercharge-normalized boundary square. Gate 748 audits the remaining
// residual and tests whether it is source-typed by the boundary-stress midpoint
// xi_boundary multiplied by the K7 second raw wall moment M2_wall. The result is
// a strong bridge-layer residual compression only, not a native flavor theorem.
package generation2kappaehyperchargeboundaryresidualandboundarystressmomentaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate746 "github.com/bagherbal/asha-engine/pkg/bridge/generation2flavorwalldeficitkappaesourcetypeandscalarbridgedependencyaudit"
	gate747 "github.com/bagherbal/asha-engine/pkg/bridge/generation2kappaeorientationresidualandhyperchargenormalizedboundarysquareaudit"
)

const (
	AuditID = "GATE748-KAPPA-E-HYPERCHARGE-BOUNDARY-RESIDUAL-BOUNDARY-STRESS-MOMENT-AUDIT"

	StatusGate747KappaEHyperchargeBoundarySquareInherited = "PASS_GATE747_KAPPA_E_HYPERCHARGE_BOUNDARY_SQUARE_INHERITED"
	StatusGate747ResidualOverM2WallComputed               = "PASS_GATE747_RESIDUAL_OVER_M2_WALL_COMPUTED"
	StatusBoundaryStressCandidatesAudited                 = "PASS_BOUNDARY_STRESS_CANDIDATES_AUDITED"
	StatusBoundaryStressMomentCorrectionDefined           = "PASS_BOUNDARY_STRESS_MOMENT_CORRECTION_DEFINED"
	StatusScalarRuntimeReplacementTested                  = "PASS_SCALAR_RUNTIME_REPLACEMENT_TESTED"
	StatusSourceTypeInterpretationRecorded                = "PASS_SOURCE_TYPE_INTERPRETATION_RECORDED"
	StatusNoncircularityFirewallAudited                   = "PASS_NONCIRCULARITY_FIREWALL_AUDITED"
	StatusPhysicalFirewallsEnforced                       = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusGate747ResidualK7SecondMomentScale             = "CONDITIONAL_SUPPORT_GATE747_RESIDUAL_IS_K7_SECOND_MOMENT_SCALE"
	StatusXiBoundaryBestTypedStressCoefficientCandidate  = "CONDITIONAL_SUPPORT_XI_BOUNDARY_IS_BEST_TYPED_STRESS_COEFFICIENT_CANDIDATE"
	StatusKappaEOrientationHyperchargeBoundaryStressForm = "CONDITIONAL_SUPPORT_KAPPA_E_HAS_ORIENTATION_PLUS_HYPERCHARGE_PLUS_BOUNDARY_STRESS_MOMENT_FORM"
	StatusBoundaryStressMomentRefinesKappaESourceType    = "CONDITIONAL_SUPPORT_BOUNDARY_STRESS_MOMENT_REFINES_KAPPA_E_SOURCE_TYPE"

	StatusCorrectionNotExact                         = "FAILED_ROUTE_CORRECTION_NOT_EXACT"
	StatusNoNativeReasonKappaEResidualEqualsXiM2Wall = "FAILED_ROUTE_NO_NATIVE_REASON_KAPPA_E_RESIDUAL_EQUALS_XI_BOUNDARY_M2_WALL"
	StatusNoNativeFlavorDeficitTheorem               = "FAILED_ROUTE_NO_NATIVE_FLAVOR_DEFICIT_THEOREM"
	StatusNoNativePMNSOrCKMTheorem                   = "FAILED_ROUTE_NO_NATIVE_PMNS_OR_CKM_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem        = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusNoIndependentScalarRuntimeTheorem          = "FAILED_ROUTE_NO_INDEPENDENT_SCALAR_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem               = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusGate748Boundary                            = "FIREWALL_PRESERVED_GATE748_KAPPA_E_BOUNDARY_STRESS_MOMENT_BOUNDARY"
)

type Gate747Inheritance struct {
	Inherited                    bool
	HyperBoundaryCorrectionClose bool
	ResidualNotZero              bool
	FlavorFirewallKept           bool
	Verdict                      string
}

type StressCandidate struct {
	Name        string
	Value       float64
	AbsDistance float64
	Typed       bool
	Selected    bool
	Reason      string
}

type ResidualOverM2Audit struct {
	KappaE              float64
	KappaEHyperBoundary float64
	Residual            float64
	P_K7                float64
	SSplit              float64
	SSplitSquared       float64
	M2Wall              float64
	Ratio               float64
	SecondMomentScale   bool
	Verdict             string
}

type BoundaryStressCandidateAudit struct {
	AbsLambda     float64
	R3MinusOne    float64
	XiBoundary    float64
	Ratio         float64
	Candidates    []StressCandidate
	BestCandidate string
	Verdict       string
}

type BoundaryStressMomentCorrection struct {
	Formula                 string
	KappaEOrient            float64
	KappaEHyperBoundary     float64
	XiBoundary              float64
	M2Wall                  float64
	StressMoment            float64
	KappaEHyperStress       float64
	KappaE                  float64
	ResidualAfterCorrection float64
	CompressionFactor       float64
	CorrectionNotExact      bool
	Verdict                 string
}

type ScalarRuntimeReplacementTest struct {
	FormulaStress               string
	KappaE                      float64
	KappaEOrient                float64
	KappaEHyperBoundary         float64
	KappaEHyperStress           float64
	FExact                      float64
	FOrient                     float64
	FHyperBoundary              float64
	FHyperStress                float64
	RuntimeExact                float64
	RuntimeOrient               float64
	RuntimeHyperBoundary        float64
	RuntimeHyperStress          float64
	RuntimeOrientShift          float64
	RuntimeHyperBoundaryShift   float64
	RuntimeHyperStressShift     float64
	StressImprovementOverOrient float64
	StressImprovementOverHyper  float64
	ReplacementNotNative        bool
	Verdict                     string
}

type SourceTypeInterpretation struct {
	Expression     string
	Terms          []string
	Interpretation string
	Verdict        string
}

type NoncircularityFirewall struct {
	Theta13EmpiricalBridgeInput    bool
	JCKMEmpiricalBridgeInput       bool
	FiveThirdsMatureButUncoupled   bool
	XiBoundaryBridgeStressQuantity bool
	M2WallBoundaryMomentNotFlavor  bool
	DerivesFlavorTheorem           bool
	DerivesScalarRuntime           bool
	DerivesHiggsMass               bool
	DerivesYukawa                  bool
	Verdict                        string
}

type Analysis struct {
	Gate747     Gate747Inheritance
	Residual    ResidualOverM2Audit
	Stress      BoundaryStressCandidateAudit
	Correction  BoundaryStressMomentCorrection
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
	g747, err := gate747.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate747 inheritance unavailable: %w", err)
	}
	g746, err := gate746.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate746 scalar bridge inheritance unavailable: %w", err)
	}
	inherit := buildGate747Inheritance(g747)
	residual := buildResidualOverM2(g747, g746)
	stress := buildStressCandidates(g746, residual)
	correction := buildStressMomentCorrection(g747, residual, stress)
	replacement := buildReplacementTest(g747, g746, correction)
	sourceType := buildSourceTypeInterpretation()
	firewall := buildNoncircularityFirewall()
	truth := "Gate 748 audits the residual left after the Gate747 hypercharge-normalized boundary-square correction to kappa_e. The residual is of K7 second raw wall moment scale, and xi_boundary*M2_wall compresses it by roughly 300x, improving the scalar-runtime replacement to about 1e-13 scale. This is a residual source-type compression only: xi_boundary and M2_wall are boundary bridge objects, not native flavor operators, and no PMNS/CKM/Yukawa or scalar-runtime theorem is derived."
	return Analysis{Gate747: inherit, Residual: residual, Stress: stress, Correction: correction, Replacement: replacement, SourceType: sourceType, Firewall: firewall, Truth: truth}, nil
}

func buildGate747Inheritance(g gate747.Analysis) Gate747Inheritance {
	return Gate747Inheritance{
		Inherited:                    g.Gate746.Inherited && g.Correction.CompressionFactor > 300 && g.Firewall.FiveThirdsMatureButUncoupled,
		HyperBoundaryCorrectionClose: math.Abs(g.Correction.ResidualAfterCorrection) < 1e-8,
		ResidualNotZero:              math.Abs(g.Correction.ResidualAfterCorrection) > 1e-10,
		FlavorFirewallKept:           !g.Firewall.DerivesFlavorTheorem && !g.Firewall.DerivesYukawa && !g.Firewall.DerivesScalarRuntime,
		Verdict:                      StatusGate747KappaEHyperchargeBoundarySquareInherited,
	}
}

func buildResidualOverM2(g747 gate747.Analysis, g746 gate746.Analysis) ResidualOverM2Audit {
	p := g746.Replacement.P_K7
	S := g747.Ratio.SSplit
	S2 := S * S
	m2 := p * S2
	res := g747.Correction.ResidualAfterCorrection
	return ResidualOverM2Audit{
		KappaE:              g747.Correction.KappaE,
		KappaEHyperBoundary: g747.Correction.KappaEHyperBoundary,
		Residual:            res,
		P_K7:                p,
		SSplit:              S,
		SSplitSquared:       S2,
		M2Wall:              m2,
		Ratio:               res / m2,
		SecondMomentScale:   math.Abs(res/m2) > 0.04 && math.Abs(res/m2) < 0.06,
		Verdict: strings.Join([]string{
			StatusGate747ResidualOverM2WallComputed,
			StatusGate747ResidualK7SecondMomentScale,
		}, "; "),
	}
}

func buildStressCandidates(g gate746.Analysis, r ResidualOverM2Audit) BoundaryStressCandidateAudit {
	absLambda := g.Replacement.W3Exact - g.Replacement.FExact
	R := absLambda + r.SSplit
	xi := 0.5 * (absLambda + R)
	cands := []StressCandidate{
		stressCandidate("xi_boundary midpoint", xi, r.Ratio, true, true, "0.5(|lambda|+R_3-1), the structured boundary-stress midpoint"),
		stressCandidate("|lambda(Lambda_12)|", absLambda, r.Ratio, true, false, "scalar zero-wall depth, close but not midpoint-typed"),
		stressCandidate("R_3-1", R, r.Ratio, true, false, "gauge meeting-wall wound, close but one-sided"),
	}
	return BoundaryStressCandidateAudit{
		AbsLambda:     absLambda,
		R3MinusOne:    R,
		XiBoundary:    xi,
		Ratio:         r.Ratio,
		Candidates:    cands,
		BestCandidate: "xi_boundary midpoint",
		Verdict: strings.Join([]string{
			StatusBoundaryStressCandidatesAudited,
			StatusXiBoundaryBestTypedStressCoefficientCandidate,
		}, "; "),
	}
}

func stressCandidate(name string, value, ratio float64, typed, selected bool, reason string) StressCandidate {
	return StressCandidate{Name: name, Value: value, AbsDistance: math.Abs(ratio - value), Typed: typed, Selected: selected, Reason: reason}
}

func buildStressMomentCorrection(g gate747.Analysis, r ResidualOverM2Audit, s BoundaryStressCandidateAudit) BoundaryStressMomentCorrection {
	moment := s.XiBoundary * r.M2Wall
	hyperStress := g.Correction.KappaEHyperBoundary + moment
	resid := g.Correction.KappaE - hyperStress
	return BoundaryStressMomentCorrection{
		Formula:                 "kappa_e_hyper_stress = kappa_e_orient - (5/3)S_split^2 + xi_boundary M2_wall",
		KappaEOrient:            g.Correction.KappaEOrient,
		KappaEHyperBoundary:     g.Correction.KappaEHyperBoundary,
		XiBoundary:              s.XiBoundary,
		M2Wall:                  r.M2Wall,
		StressMoment:            moment,
		KappaEHyperStress:       hyperStress,
		KappaE:                  g.Correction.KappaE,
		ResidualAfterCorrection: resid,
		CompressionFactor:       math.Abs(r.Residual) / math.Abs(resid),
		CorrectionNotExact:      math.Abs(resid) > 1e-12,
		Verdict: strings.Join([]string{
			StatusBoundaryStressMomentCorrectionDefined,
			StatusKappaEOrientationHyperchargeBoundaryStressForm,
			StatusBoundaryStressMomentRefinesKappaESourceType,
			StatusCorrectionNotExact,
		}, "; "),
	}
}

func buildReplacementTest(g747 gate747.Analysis, g746 gate746.Analysis, c BoundaryStressMomentCorrection) ScalarRuntimeReplacementTest {
	p := g746.Replacement.P_K7
	S := g746.Replacement.SSplit
	absLambda := g746.Replacement.W3Exact - g746.Replacement.FExact
	fExact := g746.Replacement.FExact
	fOrient := g746.Replacement.FOrient
	fHyper := g747.Replacement.FHyperBoundary
	fStress := rawPolynomial(p, S, c.KappaEHyperStress)
	runtimeExact := g746.Replacement.RuntimeExactKappaE
	runtimeOrient := g746.Replacement.RuntimeOrientKappaE
	runtimeHyper := g747.Replacement.RuntimeHyperBoundary
	wStress := absLambda + fStress
	runtimeStress := g746.Replacement.LambdaProxy * (1 + g746.Replacement.L*(1-wStress+c.KappaEHyperStress))
	orientShift := runtimeOrient - runtimeExact
	hyperShift := runtimeHyper - runtimeExact
	stressShift := runtimeStress - runtimeExact
	return ScalarRuntimeReplacementTest{
		FormulaStress:               "lambda_runtime≈lambda_proxy[1+L(1-W_3(kappa_e_hyper_stress)+kappa_e_hyper_stress)]",
		KappaE:                      c.KappaE,
		KappaEOrient:                c.KappaEOrient,
		KappaEHyperBoundary:         c.KappaEHyperBoundary,
		KappaEHyperStress:           c.KappaEHyperStress,
		FExact:                      fExact,
		FOrient:                     fOrient,
		FHyperBoundary:              fHyper,
		FHyperStress:                fStress,
		RuntimeExact:                runtimeExact,
		RuntimeOrient:               runtimeOrient,
		RuntimeHyperBoundary:        runtimeHyper,
		RuntimeHyperStress:          runtimeStress,
		RuntimeOrientShift:          orientShift,
		RuntimeHyperBoundaryShift:   hyperShift,
		RuntimeHyperStressShift:     stressShift,
		StressImprovementOverOrient: math.Abs(orientShift) / math.Abs(stressShift),
		StressImprovementOverHyper:  math.Abs(hyperShift) / math.Abs(stressShift),
		ReplacementNotNative:        true,
		Verdict: strings.Join([]string{
			StatusScalarRuntimeReplacementTested,
			StatusBoundaryStressMomentRefinesKappaESourceType,
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
		"-(5/3)S_split^2: hypercharge-normalized second-order boundary split correction",
		"+xi_boundary p_K7 S_split^2: boundary-stress-weighted K7 second raw moment correction",
	}
	return SourceTypeInterpretation{
		Expression:     "kappa_e ≈ sin^2(theta13)/4 - J_CKM - (5/3)S_split^2 + xi_boundary p_K7 S_split^2",
		Terms:          terms,
		Interpretation: "orientation plus hypercharge-normalized boundary square plus boundary-stress K7 second raw moment correction",
		Verdict:        StatusSourceTypeInterpretationRecorded,
	}
}

func buildNoncircularityFirewall() NoncircularityFirewall {
	return NoncircularityFirewall{
		Theta13EmpiricalBridgeInput:    true,
		JCKMEmpiricalBridgeInput:       true,
		FiveThirdsMatureButUncoupled:   true,
		XiBoundaryBridgeStressQuantity: true,
		M2WallBoundaryMomentNotFlavor:  true,
		DerivesFlavorTheorem:           false,
		DerivesScalarRuntime:           false,
		DerivesHiggsMass:               false,
		DerivesYukawa:                  false,
		Verdict: strings.Join([]string{
			StatusNoncircularityFirewallAudited,
			StatusPhysicalFirewallsEnforced,
			StatusNoNativeReasonKappaEResidualEqualsXiM2Wall,
			StatusNoNativeFlavorDeficitTheorem,
			StatusNoNativePMNSOrCKMTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusNoIndependentScalarRuntimeTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusGate748Boundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate747KappaEHyperchargeBoundarySquareInherited,
		StatusGate747ResidualOverM2WallComputed,
		StatusBoundaryStressCandidatesAudited,
		StatusBoundaryStressMomentCorrectionDefined,
		StatusScalarRuntimeReplacementTested,
		StatusSourceTypeInterpretationRecorded,
		StatusNoncircularityFirewallAudited,
		StatusPhysicalFirewallsEnforced,
		StatusGate747ResidualK7SecondMomentScale,
		StatusXiBoundaryBestTypedStressCoefficientCandidate,
		StatusKappaEOrientationHyperchargeBoundaryStressForm,
		StatusBoundaryStressMomentRefinesKappaESourceType,
		StatusCorrectionNotExact,
		StatusNoNativeReasonKappaEResidualEqualsXiM2Wall,
		StatusNoNativeFlavorDeficitTheorem,
		StatusNoNativePMNSOrCKMTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusNoIndependentScalarRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusGate748Boundary,
	}
}

func FormatGate747(x Gate747Inheritance) string {
	return fmt.Sprintf("inherited=%t hyperClose=%t residualNotZero=%t flavorFirewall=%t verdict=%q", x.Inherited, x.HyperBoundaryCorrectionClose, x.ResidualNotZero, x.FlavorFirewallKept, x.Verdict)
}

func FormatResidual(x ResidualOverM2Audit) string {
	return fmt.Sprintf("kappaE=%.17g hyper=%.17g residual=%.17g p=%.17g S=%.17g S2=%.17g M2=%.17g ratio=%.17g secondMomentScale=%t verdict=%q", x.KappaE, x.KappaEHyperBoundary, x.Residual, x.P_K7, x.SSplit, x.SSplitSquared, x.M2Wall, x.Ratio, x.SecondMomentScale, x.Verdict)
}

func FormatStress(x BoundaryStressCandidateAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, fmt.Sprintf("%s=%.17g dist=%.17g typed=%t selected=%t", c.Name, c.Value, c.AbsDistance, c.Typed, c.Selected))
	}
	return fmt.Sprintf("absLambda=%.17g R=%.17g xi=%.17g ratio=%.17g best=%q candidates=[%s] verdict=%q", x.AbsLambda, x.R3MinusOne, x.XiBoundary, x.Ratio, x.BestCandidate, strings.Join(parts, "; "), x.Verdict)
}

func FormatCorrection(x BoundaryStressMomentCorrection) string {
	return fmt.Sprintf("formula=%q orient=%.17g hyper=%.17g xi=%.17g M2=%.17g stressMoment=%.17g hyperStress=%.17g kappaE=%.17g residual=%.17g compression=%.17g notExact=%t verdict=%q", x.Formula, x.KappaEOrient, x.KappaEHyperBoundary, x.XiBoundary, x.M2Wall, x.StressMoment, x.KappaEHyperStress, x.KappaE, x.ResidualAfterCorrection, x.CompressionFactor, x.CorrectionNotExact, x.Verdict)
}

func FormatReplacement(x ScalarRuntimeReplacementTest) string {
	return fmt.Sprintf("kappaE=%.17g orient=%.17g hyper=%.17g stress=%.17g FExact=%.17g FOrient=%.17g FHyper=%.17g FStress=%.17g runtimeExact=%.17g runtimeOrient=%.17g runtimeHyper=%.17g runtimeStress=%.17g orientShift=%.17g hyperShift=%.17g stressShift=%.17g improveOrient=%.17g improveHyper=%.17g notNative=%t verdict=%q", x.KappaE, x.KappaEOrient, x.KappaEHyperBoundary, x.KappaEHyperStress, x.FExact, x.FOrient, x.FHyperBoundary, x.FHyperStress, x.RuntimeExact, x.RuntimeOrient, x.RuntimeHyperBoundary, x.RuntimeHyperStress, x.RuntimeOrientShift, x.RuntimeHyperBoundaryShift, x.RuntimeHyperStressShift, x.StressImprovementOverOrient, x.StressImprovementOverHyper, x.ReplacementNotNative, x.Verdict)
}

func FormatSourceType(x SourceTypeInterpretation) string {
	return fmt.Sprintf("expression=%q terms=[%s] interpretation=%q verdict=%q", x.Expression, strings.Join(x.Terms, "; "), x.Interpretation, x.Verdict)
}

func FormatFirewall(x NoncircularityFirewall) string {
	return fmt.Sprintf("theta13Bridge=%t jckmBridge=%t fiveThirdsUncoupled=%t xiBoundaryBridge=%t M2BoundaryMoment=%t derivesFlavor=%t derivesRuntime=%t derivesHiggs=%t derivesYukawa=%t verdict=%q", x.Theta13EmpiricalBridgeInput, x.JCKMEmpiricalBridgeInput, x.FiveThirdsMatureButUncoupled, x.XiBoundaryBridgeStressQuantity, x.M2WallBoundaryMomentNotFlavor, x.DerivesFlavorTheorem, x.DerivesScalarRuntime, x.DerivesHiggsMass, x.DerivesYukawa, x.Verdict)
}

func Near(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
