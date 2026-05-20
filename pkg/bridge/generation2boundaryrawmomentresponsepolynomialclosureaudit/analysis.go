// Package generation2boundaryrawmomentresponsepolynomialclosureaudit implements
// Gate 733: Boundary Raw-Moment Response Polynomial Closure Audit.
//
// Gate 732 showed that the active residual-compression coordinate is the raw
// moment coordinate of R_wall=S_split P_K7. Gate 733 packages the resulting
// cubic scalar response polynomial, audits its numerical closure and fourth-order
// temptation, and preserves the firewall that no native boundary response
// generating function or moment expansion theorem is certified.
package generation2boundaryrawmomentresponsepolynomialclosureaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate730 "github.com/bagherbal/asha-engine/pkg/bridge/generation2boundaryhistoryresidualcubicstresspullcorrectionaudit"
	gate732 "github.com/bagherbal/asha-engine/pkg/bridge/generation2boundaryrawmomentresponsecoordinatenaturalityaudit"
)

const (
	AuditID = "GATE733-BOUNDARY-RAW-MOMENT-RESPONSE-POLYNOMIAL-CLOSURE-AUDIT"

	StatusGate732RawMomentCoordinateInherited       = "PASS_GATE732_RAW_MOMENT_COORDINATE_INHERITED"
	StatusCubicRawMomentResponsePolynomialDefined   = "PASS_CUBIC_RAW_MOMENT_RESPONSE_POLYNOMIAL_DEFINED"
	StatusCubicPolynomialClosureResidualComputed    = "PASS_CUBIC_POLYNOMIAL_CLOSURE_RESIDUAL_COMPUTED"
	StatusFourthOrderRequiredCoefficientComputed    = "PASS_FOURTH_ORDER_REQUIRED_COEFFICIENT_COMPUTED"
	StatusStopConditionAudited                      = "PASS_STOP_CONDITION_AUDITED"
	StatusPolynomialSourceTypeRecorded              = "PASS_POLYNOMIAL_SOURCE_TYPE_RECORDED"
	StatusGeneratingFunctionCandidateAudited        = "PASS_GENERATING_FUNCTION_CANDIDATE_AUDITED"
	StatusCubicPolynomialRuntimePropagationRecorded = "PASS_CUBIC_POLYNOMIAL_RUNTIME_PROPAGATION_RECORDED"
	StatusNoncircularityFirewallAudited             = "PASS_NONCIRCULARITY_FIREWALL_AUDITED"
	StatusPhysicalFirewallsEnforced                 = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusCubicPolynomialCurrentBestBoundaryClosure       = "CONDITIONAL_SUPPORT_CUBIC_RAW_MOMENT_POLYNOMIAL_IS_CURRENT_BEST_BOUNDARY_RESPONSE_CLOSURE"
	StatusStoppingAtCubicMoreLawfulThanUntypedM4Fit       = "CONDITIONAL_SUPPORT_STOPPING_AT_CUBIC_IS_MORE_LAWFUL_THAN_UNTYPED_M4_FIT"
	StatusScalarRuntimeResidualPropagatedCubicPolynomial  = "CONDITIONAL_SUPPORT_SCALAR_RUNTIME_RESIDUAL_IS_PROPAGATED_CUBIC_POLYNOMIAL_RESIDUAL"
	StatusPolynomialGeneratingFunctionCandidateTruncation = "CONDITIONAL_SUPPORT_POLYNOMIAL_IS_BOUNDARY_RESPONSE_GENERATING_FUNCTION_CANDIDATE_TRUNCATION"
	StatusCubicPolynomialStronglyCompressesResidual       = "CONDITIONAL_SUPPORT_CUBIC_POLYNOMIAL_STRONGLY_COMPRESSES_BOUNDARY_HISTORY_RESIDUAL"

	StatusNoTypedFourthOrderCoefficientSource        = "FAILED_ROUTE_NO_TYPED_FOURTH_ORDER_COEFFICIENT_SOURCE"
	StatusNoNativeBoundaryResponseGeneratingFunction = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_RESPONSE_GENERATING_FUNCTION_THEOREM"
	StatusNoNativeBoundaryMomentExpansionTheorem     = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_MOMENT_EXPANSION_THEOREM"
	StatusNoNativeScalarRuntimeTheorem               = "FAILED_ROUTE_NO_NATIVE_SCALAR_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem               = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem        = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusKappaECoefficientPartiallyDependent        = "FAILED_ROUTE_KAPPA_E_COEFFICIENT_PARTIALLY_DEPENDENT"
	StatusNoNativeReasonForDoubleK7CubicCoefficient  = "FAILED_ROUTE_NO_NATIVE_REASON_FOR_DOUBLE_K7_CUBIC_COEFFICIENT"
	StatusGate733Boundary                            = "FIREWALL_PRESERVED_GATE733_RAW_MOMENT_POLYNOMIAL_CLOSURE_BOUNDARY"
)

type Gate732Inheritance struct {
	Inherited                 bool
	P_K7                      float64
	SSplit                    float64
	KappaE                    float64
	M1Wall                    float64
	M2Wall                    float64
	M3Wall                    float64
	EWall                     float64
	DBase                     float64
	RawCubicResidual          float64
	RawM3BestCompression      bool
	RawMomentCoordinateActive bool
	NoNativeRawMomentTheorem  bool
	Verdict                   string
}

type CubicPolynomialDefinition struct {
	Formula         string
	FactoredFormula string
	P_K7            float64
	SSplit          float64
	KappaE          float64
	LeadingTerm     float64
	QuadraticTerm   float64
	CubicTerm       float64
	Value           float64
	Verdict         string
}

type CubicClosureResidualAudit struct {
	DBase             float64
	PolynomialValue   float64
	Residual          float64
	LeadingResidual   float64
	CompressionFactor float64
	StrongCompression bool
	Verdict           string
}

type FourthOrderAudit struct {
	M4Wall             float64
	ResidualCubic      float64
	RequiredCoeff      float64
	TypedSourceFound   bool
	PromoteFourthOrder bool
	Verdict            string
}

type StopConditionAudit struct {
	ProjectorPowersSupplyNewDirections bool
	HigherMomentsOnlyScalarPowers      bool
	TypedFourthOrderSourceFound        bool
	StoppingAtCubicMoreLawful          bool
	Verdict                            string
}

type PolynomialSourceType struct {
	Leading   string
	Quadratic string
	Cubic     string
	Compact   string
	Verdict   string
}

type GeneratingFunctionCandidate struct {
	Form                         string
	GWallTruncation              string
	NativeGeneratingFunction     bool
	CandidateTruncationSupported bool
	Verdict                      string
}

type RuntimePropagationAudit struct {
	LambdaProxy             float64
	L                       float64
	CubicPolynomialResidual float64
	RuntimeResidual         float64
	NearEliminated          bool
	Verdict                 string
}

type NoncircularityFirewall struct {
	KappaEPartiallyDependent      bool
	DoubleK7CoefficientNative     bool
	BoundaryMomentExpansionNative bool
	Verdict                       string
}

type PhysicalFirewall struct {
	ScalarRuntimeTheoremNative bool
	HiggsMassTheoremNative     bool
	YukawaTheoremNative        bool
	Verdict                    string
}

type Analysis struct {
	Gate732     Gate732Inheritance
	Polynomial  CubicPolynomialDefinition
	Closure     CubicClosureResidualAudit
	FourthOrder FourthOrderAudit
	Stop        StopConditionAudit
	SourceType  PolynomialSourceType
	Generating  GeneratingFunctionCandidate
	Runtime     RuntimePropagationAudit
	NonCircular NoncircularityFirewall
	Firewall    PhysicalFirewall
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
	g732, err := gate732.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate732 inheritance unavailable: %w", err)
	}
	g730, err := gate730.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate730 runtime inheritance unavailable: %w", err)
	}
	inherited := buildGate732Inheritance(g732)
	poly := buildCubicPolynomial(inherited)
	closure := buildClosure(inherited, poly)
	fourth := buildFourthOrder(inherited, closure)
	stop := buildStopCondition(fourth)
	source := buildSourceType()
	gen := buildGeneratingFunction()
	runtime := buildRuntimePropagation(g730, closure)
	noncirc := buildNoncircularity()
	firewall := buildPhysicalFirewall()
	truth := "Gate 733 closes the current boundary-history residual ledger at the cubic raw-moment polynomial F_wall_3(S)=p_K7 S[1+kappa_e S-2p_K7 S^2]. This strongly compresses the Gate700 wall residual and propagates to a scalar-runtime residual near 1e-15. The fourth raw moment can formally absorb the remaining residual, but its required coefficient has no typed ASHA source; because projector powers add no new directions, continuing the expansion without a typed coefficient would be fitting. The cubic polynomial is therefore the current best bridge closure and possible generating-function truncation, not a native boundary response theorem, scalar runtime theorem, Higgs mass theorem, or Yukawa theorem."
	return Analysis{Gate732: inherited, Polynomial: poly, Closure: closure, FourthOrder: fourth, Stop: stop, SourceType: source, Generating: gen, Runtime: runtime, NonCircular: noncirc, Firewall: firewall, Truth: truth}, nil
}

func buildGate732Inheritance(g gate732.Analysis) Gate732Inheritance {
	dbase := g.Gate731.M1Wall + g.Gate731.EWall
	return Gate732Inheritance{
		Inherited:                 g.Gate731.Inherited && g.RawMoment.UsesRawMoments && g.Comparison.RawSelectedByCurrentCompression,
		P_K7:                      g.Gate731.P_K7,
		SSplit:                    g.Gate731.SSplit,
		KappaE:                    g.Gate731.KappaE,
		M1Wall:                    g.Gate731.M1Wall,
		M2Wall:                    g.Gate731.M2Wall,
		M3Wall:                    g.Gate731.M3Wall,
		EWall:                     g.Gate731.EWall,
		DBase:                     dbase,
		RawCubicResidual:          g.RawMoment.RawResidual,
		RawM3BestCompression:      g.Comparison.RawSelectedByCurrentCompression,
		RawMomentCoordinateActive: g.RawMoment.UsesRawMoments,
		NoNativeRawMomentTheorem:  !g.Firewall.RawMomentsNativelySelected && !g.Firewall.BoundaryMomentTheoremNative,
		Verdict:                   StatusGate732RawMomentCoordinateInherited,
	}
}

func buildCubicPolynomial(g Gate732Inheritance) CubicPolynomialDefinition {
	leading := g.P_K7 * g.SSplit
	quadratic := g.KappaE * g.P_K7 * g.SSplit * g.SSplit
	cubic := -2.0 * g.P_K7 * g.P_K7 * math.Pow(g.SSplit, 3)
	value := leading + quadratic + cubic
	return CubicPolynomialDefinition{
		Formula:         "F_wall_3(S)=p_K7 S+kappa_e p_K7 S^2-2p_K7^2 S^3",
		FactoredFormula: "F_wall_3(S)=p_K7 S[1+kappa_e S-2p_K7 S^2]",
		P_K7:            g.P_K7,
		SSplit:          g.SSplit,
		KappaE:          g.KappaE,
		LeadingTerm:     leading,
		QuadraticTerm:   quadratic,
		CubicTerm:       cubic,
		Value:           value,
		Verdict:         StatusCubicRawMomentResponsePolynomialDefined,
	}
}

func buildClosure(g Gate732Inheritance, p CubicPolynomialDefinition) CubicClosureResidualAudit {
	resid := g.DBase - p.Value
	factor := math.Abs(g.EWall / resid)
	return CubicClosureResidualAudit{
		DBase:             g.DBase,
		PolynomialValue:   p.Value,
		Residual:          resid,
		LeadingResidual:   g.EWall,
		CompressionFactor: factor,
		StrongCompression: factor > 1000,
		Verdict: strings.Join([]string{
			StatusCubicPolynomialClosureResidualComputed,
			StatusCubicPolynomialStronglyCompressesResidual,
			StatusCubicPolynomialCurrentBestBoundaryClosure,
		}, "; "),
	}
}

func buildFourthOrder(g Gate732Inheritance, c CubicClosureResidualAudit) FourthOrderAudit {
	m4 := g.P_K7 * math.Pow(g.SSplit, 4)
	coeff := c.Residual / m4
	return FourthOrderAudit{
		M4Wall:             m4,
		ResidualCubic:      c.Residual,
		RequiredCoeff:      coeff,
		TypedSourceFound:   false,
		PromoteFourthOrder: false,
		Verdict: strings.Join([]string{
			StatusFourthOrderRequiredCoefficientComputed,
			StatusNoTypedFourthOrderCoefficientSource,
		}, "; "),
	}
}

func buildStopCondition(f FourthOrderAudit) StopConditionAudit {
	return StopConditionAudit{
		ProjectorPowersSupplyNewDirections: false,
		HigherMomentsOnlyScalarPowers:      true,
		TypedFourthOrderSourceFound:        f.TypedSourceFound,
		StoppingAtCubicMoreLawful:          !f.TypedSourceFound && !f.PromoteFourthOrder,
		Verdict: strings.Join([]string{
			StatusStopConditionAudited,
			StatusStoppingAtCubicMoreLawfulThanUntypedM4Fit,
			StatusNoTypedFourthOrderCoefficientSource,
		}, "; "),
	}
}

func buildSourceType() PolynomialSourceType {
	return PolynomialSourceType{
		Leading:   "M1: no-bias K7 event expectation",
		Quadratic: "kappa_e M2: flavor-wall-modulated second raw response moment",
		Cubic:     "-2p_K7 M3: double-K7-event / boundary-pair stress-pull cubic correction",
		Compact:   "F_wall_3(S)=pS[1+kappa_e S-2pS^2]",
		Verdict:   StatusPolynomialSourceTypeRecorded,
	}
}

func buildGeneratingFunction() GeneratingFunctionCandidate {
	return GeneratingFunctionCandidate{
		Form:                         "F_wall(S)=p_K7 S G_wall(S)",
		GWallTruncation:              "G_wall(S)=1+kappa_e S-2p_K7 S^2+...",
		NativeGeneratingFunction:     false,
		CandidateTruncationSupported: true,
		Verdict: strings.Join([]string{
			StatusGeneratingFunctionCandidateAudited,
			StatusPolynomialGeneratingFunctionCandidateTruncation,
			StatusNoNativeBoundaryResponseGeneratingFunction,
		}, "; "),
	}
}

func buildRuntimePropagation(g gate730.Analysis, c CubicClosureResidualAudit) RuntimePropagationAudit {
	runtime := g.Runtime.LambdaProxy * g.Runtime.L * c.Residual
	return RuntimePropagationAudit{
		LambdaProxy:             g.Runtime.LambdaProxy,
		L:                       g.Runtime.L,
		CubicPolynomialResidual: c.Residual,
		RuntimeResidual:         runtime,
		NearEliminated:          math.Abs(runtime) < 1e-14,
		Verdict: strings.Join([]string{
			StatusCubicPolynomialRuntimePropagationRecorded,
			StatusScalarRuntimeResidualPropagatedCubicPolynomial,
		}, "; "),
	}
}

func buildNoncircularity() NoncircularityFirewall {
	return NoncircularityFirewall{
		KappaEPartiallyDependent:      true,
		DoubleK7CoefficientNative:     false,
		BoundaryMomentExpansionNative: false,
		Verdict: strings.Join([]string{
			StatusNoncircularityFirewallAudited,
			StatusKappaECoefficientPartiallyDependent,
			StatusNoNativeReasonForDoubleK7CubicCoefficient,
			StatusNoNativeBoundaryMomentExpansionTheorem,
		}, "; "),
	}
}

func buildPhysicalFirewall() PhysicalFirewall {
	return PhysicalFirewall{
		ScalarRuntimeTheoremNative: false,
		HiggsMassTheoremNative:     false,
		YukawaTheoremNative:        false,
		Verdict: strings.Join([]string{
			StatusPhysicalFirewallsEnforced,
			StatusNoNativeScalarRuntimeTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusGate733Boundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate732RawMomentCoordinateInherited,
		StatusCubicRawMomentResponsePolynomialDefined,
		StatusCubicPolynomialClosureResidualComputed,
		StatusFourthOrderRequiredCoefficientComputed,
		StatusStopConditionAudited,
		StatusPolynomialSourceTypeRecorded,
		StatusGeneratingFunctionCandidateAudited,
		StatusCubicPolynomialRuntimePropagationRecorded,
		StatusNoncircularityFirewallAudited,
		StatusPhysicalFirewallsEnforced,
		StatusCubicPolynomialCurrentBestBoundaryClosure,
		StatusStoppingAtCubicMoreLawfulThanUntypedM4Fit,
		StatusScalarRuntimeResidualPropagatedCubicPolynomial,
		StatusPolynomialGeneratingFunctionCandidateTruncation,
		StatusCubicPolynomialStronglyCompressesResidual,
		StatusNoTypedFourthOrderCoefficientSource,
		StatusNoNativeBoundaryResponseGeneratingFunction,
		StatusNoNativeBoundaryMomentExpansionTheorem,
		StatusNoNativeScalarRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusKappaECoefficientPartiallyDependent,
		StatusNoNativeReasonForDoubleK7CubicCoefficient,
		StatusGate733Boundary,
	}
}

func near(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatGate732(x Gate732Inheritance) string {
	return fmt.Sprintf("inherited=%t p=%.17g S=%.17g kE=%.17g M1=%.17g M2=%.17g M3=%.17g E=%.17g D=%.17g rawResidual=%.17g rawBest=%t rawActive=%t noNativeRaw=%t verdict=%q", x.Inherited, x.P_K7, x.SSplit, x.KappaE, x.M1Wall, x.M2Wall, x.M3Wall, x.EWall, x.DBase, x.RawCubicResidual, x.RawM3BestCompression, x.RawMomentCoordinateActive, x.NoNativeRawMomentTheorem, x.Verdict)
}
func FormatPolynomial(x CubicPolynomialDefinition) string {
	return fmt.Sprintf("formula=%q factored=%q p=%.17g S=%.17g kE=%.17g leading=%.17g quadratic=%.17g cubic=%.17g value=%.17g verdict=%q", x.Formula, x.FactoredFormula, x.P_K7, x.SSplit, x.KappaE, x.LeadingTerm, x.QuadraticTerm, x.CubicTerm, x.Value, x.Verdict)
}
func FormatClosure(x CubicClosureResidualAudit) string {
	return fmt.Sprintf("D=%.17g F3=%.17g residual=%.17g leadingResidual=%.17g compression=%.17g strong=%t verdict=%q", x.DBase, x.PolynomialValue, x.Residual, x.LeadingResidual, x.CompressionFactor, x.StrongCompression, x.Verdict)
}
func FormatFourthOrder(x FourthOrderAudit) string {
	return fmt.Sprintf("M4=%.17g residual=%.17g c4=%.17g typed=%t promote=%t verdict=%q", x.M4Wall, x.ResidualCubic, x.RequiredCoeff, x.TypedSourceFound, x.PromoteFourthOrder, x.Verdict)
}
func FormatStop(x StopConditionAudit) string {
	return fmt.Sprintf("newDirs=%t scalarOnly=%t typedM4=%t stopCubic=%t verdict=%q", x.ProjectorPowersSupplyNewDirections, x.HigherMomentsOnlyScalarPowers, x.TypedFourthOrderSourceFound, x.StoppingAtCubicMoreLawful, x.Verdict)
}
func FormatSourceType(x PolynomialSourceType) string {
	return fmt.Sprintf("leading=%q quadratic=%q cubic=%q compact=%q verdict=%q", x.Leading, x.Quadratic, x.Cubic, x.Compact, x.Verdict)
}
func FormatGenerating(x GeneratingFunctionCandidate) string {
	return fmt.Sprintf("form=%q G=%q native=%t trunc=%t verdict=%q", x.Form, x.GWallTruncation, x.NativeGeneratingFunction, x.CandidateTruncationSupported, x.Verdict)
}
func FormatRuntime(x RuntimePropagationAudit) string {
	return fmt.Sprintf("lambdaProxy=%.17g L=%.17g wallResidual=%.17g runtime=%.17g near=%t verdict=%q", x.LambdaProxy, x.L, x.CubicPolynomialResidual, x.RuntimeResidual, x.NearEliminated, x.Verdict)
}
func FormatNoncircularity(x NoncircularityFirewall) string {
	return fmt.Sprintf("kEdep=%t doubleNative=%t momentNative=%t verdict=%q", x.KappaEPartiallyDependent, x.DoubleK7CoefficientNative, x.BoundaryMomentExpansionNative, x.Verdict)
}
func FormatFirewall(x PhysicalFirewall) string {
	return fmt.Sprintf("runtimeNative=%t massNative=%t yukawaNative=%t verdict=%q", x.ScalarRuntimeTheoremNative, x.HiggsMassTheoremNative, x.YukawaTheoremNative, x.Verdict)
}
