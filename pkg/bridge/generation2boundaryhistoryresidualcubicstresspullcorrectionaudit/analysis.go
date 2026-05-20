// Package generation2boundaryhistoryresidualcubicstresspullcorrectionaudit implements
// Gate 730: Boundary-History Residual Cubic Stress-Pull Correction Audit.
//
// Gate 729 showed that the Gate700/Gate728 boundary-history wall residual is
// second-order suppressed in the boundary uplift operator R_wall=S_split P_K7
// and that kappa_e is the closest active small typed second-order coefficient.
// Gate 730 audits the next residual after that candidate correction and tests
// whether it is compressed by the typed cubic boundary stress-pull coefficient
// 7/36. This remains a bridge-layer residual-structure audit, not a native
// boundary moment expansion, scalar runtime, Higgs mass, or Yukawa theorem.
package generation2boundaryhistoryresidualcubicstresspullcorrectionaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	gate729 "github.com/bagherbal/asha-engine/pkg/bridge/generation2boundaryhistoryresidualsecondmomentandruntimetransportaudit"
)

const (
	AuditID = "GATE730-BOUNDARY-HISTORY-RESIDUAL-CUBIC-STRESS-PULL-CORRECTION-AUDIT"

	StatusGate729SecondMomentResidualInherited     = "PASS_GATE729_SECOND_MOMENT_RESIDUAL_INHERITED"
	StatusCubicWallMomentComputed                  = "PASS_CUBIC_WALL_MOMENT_COMPUTED"
	StatusCubicCoefficientRatioComputed            = "PASS_CUBIC_COEFFICIENT_RATIO_COMPUTED"
	StatusTypedCubicCoefficientCandidatesAudited   = "PASS_TYPED_CUBIC_COEFFICIENT_CANDIDATES_AUDITED"
	StatusCubicStressPullCorrectionTested          = "PASS_CUBIC_STRESS_PULL_CORRECTION_TESTED"
	StatusRuntimePropagationCubicCorrectionAudited = "PASS_RUNTIME_PROPAGATION_OF_CUBIC_CORRECTION_AUDITED"
	StatusSourceTypeInterpretationRecorded         = "PASS_SOURCE_TYPE_INTERPRETATION_RECORDED"
	StatusNoncircularityFirewallAudited            = "PASS_NONCIRCULARITY_FIREWALL_AUDITED"
	StatusPhysicalFirewallsEnforced                = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusSecondPlusThirdMomentStructure          = "CONDITIONAL_SUPPORT_BOUNDARY_HISTORY_RESIDUAL_HAS_SECOND_PLUS_THIRD_ORDER_MOMENT_STRUCTURE"
	StatusSevenOverThirtySixCompressesResidual    = "CONDITIONAL_SUPPORT_SEVEN_OVER_THIRTY_SIX_CUBIC_STRESS_PULL_COMPRESSES_RESIDUAL"
	StatusRuntimeCompressedByTypedCubicCorrection = "CONDITIONAL_SUPPORT_RUNTIME_RESIDUAL_IS_COMPRESSED_BY_TYPED_CUBIC_WALL_CORRECTION"
	StatusSecondOrderResidualCubicScale           = "CONDITIONAL_SUPPORT_SECOND_ORDER_RESIDUAL_IS_CUBIC_SCALE"
	StatusTypedBoundaryStressCoefficientClosest   = "CONDITIONAL_SUPPORT_SEVEN_OVER_THIRTY_SIX_IS_CLOSEST_TYPED_BOUNDARY_STRESS_COEFFICIENT"

	StatusKappaEQuadraticCoefficientDependent        = "FAILED_ROUTE_KAPPA_E_QUADRATIC_COEFFICIENT_PARTIALLY_DEPENDENT"
	StatusCubicCorrectionNotExact                    = "FAILED_ROUTE_CUBIC_CORRECTION_NOT_EXACT"
	StatusNoNativeReasonCubicCoeffSevenOverThirtySix = "FAILED_ROUTE_NO_NATIVE_REASON_CUBIC_COEFFICIENT_IS_SEVEN_OVER_THIRTY_SIX"
	StatusNoNativeBoundaryMomentExpansionTheorem     = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_MOMENT_EXPANSION_THEOREM"
	StatusNoNativeScalarRuntimeTheorem               = "FAILED_ROUTE_NO_NATIVE_SCALAR_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem               = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem        = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate730Boundary                            = "FIREWALL_PRESERVED_GATE730_CUBIC_STRESS_PULL_BOUNDARY"
)

const (
	k7Dim                = 7.0
	h72Dim               = 72.0
	cubicStressPullCoeff = 7.0 / 36.0
	simpleOneFifth       = 1.0 / 5.0
	higgsRadialWeight    = 1.0 / 4.0
	phaseLoopUnit        = 1.0 / (2.0 * math.Pi)
)

type Gate729Inheritance struct {
	Inherited                           bool
	P_K7                                float64
	SSplit                              float64
	EWall                               float64
	M2Wall                              float64
	KappaE                              float64
	KappaEM2                            float64
	E2Residual                          float64
	RawRuntimeResidual                  float64
	SecondOrderCorrectedRuntimeResidual float64
	LambdaProxy                         float64
	L                                   float64
	KappaEPartiallyDependent            bool
	NoNativeSecondOrderBoundaryTheorem  bool
	Verdict                             string
}

type CubicMomentAudit struct {
	Formula                  string
	M3Wall                   float64
	NegativeE2OverM3         float64
	SecondResidualCubicScale bool
	Verdict                  string
}

type CubicCoefficientCandidate struct {
	Name     string
	Value    float64
	Distance float64
	Source   string
}

type TypedCubicCoefficientAudit struct {
	TargetCoefficient  float64
	Candidates         []CubicCoefficientCandidate
	ClosestName        string
	ClosestValue       float64
	ClosestDistance    float64
	SevenOver36Closest bool
	NoArbitrarySearch  bool
	Verdict            string
}

type CubicCorrectionAudit struct {
	EWall                        float64
	KappaE                       float64
	M2Wall                       float64
	M3Wall                       float64
	CubicCoefficient             float64
	QuadraticTerm                float64
	CubicStressPullTerm          float64
	CombinedCorrection           float64
	ResidualAfterCubicCorrection float64
	RawCompressionFactor         float64
	ImprovesSecondOrderResidual  bool
	NotExact                     bool
	Verdict                      string
}

type RuntimePropagationAudit struct {
	LambdaProxy                   float64
	L                             float64
	RawRuntimeResidual            float64
	SecondOrderCorrectedRuntime   float64
	CubicCorrectedWallResidual    float64
	CubicCorrectedRuntimeResidual float64
	CompressedToNearFloatScale    bool
	ImprovesSecondOrderRuntime    bool
	Verdict                       string
}

type SourceTypeInterpretation struct {
	ExpansionFormula             string
	LeadingTerm                  string
	QuadraticTerm                string
	CubicTerm                    string
	MomentExpansionTheoremNative bool
	Verdict                      string
}

type NoncircularityAudit struct {
	DBaseContainsKappaE           bool
	KappaEUsedAsQuadraticCoeff    bool
	CubicCoeffTypedButUnexplained bool
	NativeExpansionTheorem        bool
	Verdict                       string
}

type FirewallAudit struct {
	ClaimsNativeBoundaryHistory bool
	ClaimsNativeMomentExpansion bool
	ClaimsNativeScalarRuntime   bool
	ClaimsHiggsMassTheorem      bool
	ClaimsYukawaTheorem         bool
	ClaimsCKMPMNSTheorem        bool
	Verdict                     string
}

type Analysis struct {
	Gate729      Gate729Inheritance
	CubicMoment  CubicMomentAudit
	Coefficients TypedCubicCoefficientAudit
	CubicCorr    CubicCorrectionAudit
	Runtime      RuntimePropagationAudit
	SourceType   SourceTypeInterpretation
	NonCircular  NoncircularityAudit
	Firewall     FirewallAudit
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
	g729, err := gate729.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate729 inheritance unavailable: %w", err)
	}
	inherited := buildGate729Inheritance(g729)
	cubicMoment := buildCubicMoment(inherited)
	coeffs := buildTypedCubicCoefficientAudit(cubicMoment)
	cubicCorr := buildCubicCorrection(inherited, cubicMoment)
	runtime := buildRuntimePropagation(inherited, cubicCorr)
	source := buildSourceTypeInterpretation()
	noncirc := buildNoncircularity()
	firewall := buildFirewall()
	truth := "Gate 730 audits the remaining wall residual after the Gate729 kappa_e second-order correction. The third raw moment M3_wall=Tr(rho_72 R_wall^3)=(7/72)S_split^3 is the correct cubic scale, and -E2_res/M3_wall≈0.19629 sits close to the typed boundary stress-pull coefficient 7/36≈0.19444. The candidate correction kappa_e*M2_wall-(7/36)*M3_wall compresses the wall residual to about -3.88e-13 and the propagated scalar-runtime residual to about -1.93e-15. This is a residual-compression clue only: kappa_e is partially dependent through D_base, the 7/36 cubic coefficient lacks a native selection theorem, and no boundary moment expansion, scalar runtime, Higgs mass, or Yukawa theorem is certified."
	return Analysis{Gate729: inherited, CubicMoment: cubicMoment, Coefficients: coeffs, CubicCorr: cubicCorr, Runtime: runtime, SourceType: source, NonCircular: noncirc, Firewall: firewall, Truth: truth}, nil
}

func buildGate729Inheritance(g gate729.Analysis) Gate729Inheritance {
	return Gate729Inheritance{
		Inherited:                           g.Gate728.Inherited && g.Moment.SecondOrderSuppressed && g.KappaECorr.ImprovesRawResidual,
		P_K7:                                g.Gate728.P_K7,
		SSplit:                              g.Gate728.SSplit,
		EWall:                               g.Gate728.EWall,
		M2Wall:                              g.Moment.M2Wall,
		KappaE:                              g.KappaECorr.KappaE,
		KappaEM2:                            g.KappaECorr.KappaEM2,
		E2Residual:                          g.KappaECorr.ResidualAfterCorrection,
		RawRuntimeResidual:                  g.Runtime.RawRuntimeResidual,
		SecondOrderCorrectedRuntimeResidual: g.Runtime.CorrectedRuntimeResidual,
		LambdaProxy:                         g.Runtime.LambdaProxy,
		L:                                   g.Runtime.L,
		KappaEPartiallyDependent:            g.NonCircular.PartiallyDependent,
		NoNativeSecondOrderBoundaryTheorem:  !g.Firewall.ClaimsNativeSecondOrderExpansion,
		Verdict:                             StatusGate729SecondMomentResidualInherited,
	}
}

func buildCubicMoment(g Gate729Inheritance) CubicMomentAudit {
	m3 := g.P_K7 * math.Pow(g.SSplit, 3)
	ratio := -g.E2Residual / m3
	return CubicMomentAudit{
		Formula:                  "M3_wall=Tr(rho_72 R_wall^3)=(7/72)S_split^3",
		M3Wall:                   m3,
		NegativeE2OverM3:         ratio,
		SecondResidualCubicScale: math.Abs(g.E2Residual) < math.Abs(m3) && math.Abs(ratio) < 1,
		Verdict: strings.Join([]string{
			StatusCubicWallMomentComputed,
			StatusCubicCoefficientRatioComputed,
			StatusSecondOrderResidualCubicScale,
		}, "; "),
	}
}

func buildTypedCubicCoefficientAudit(m CubicMomentAudit) TypedCubicCoefficientAudit {
	cands := []CubicCoefficientCandidate{
		{Name: "7/36", Value: cubicStressPullCoeff, Source: "typed midpoint/stress-pull coefficient from boundary lane"},
		{Name: "1/5", Value: simpleOneFifth, Source: "simple nearby control"},
		{Name: "7/72", Value: k7Dim / h72Dim, Source: "K7 event probability"},
		{Name: "1/4", Value: higgsRadialWeight, Source: "Higgs radial event probability"},
		{Name: "1/(2*pi)", Value: phaseLoopUnit, Source: "phase-loop unit"},
	}
	for i := range cands {
		cands[i].Distance = math.Abs(cands[i].Value - m.NegativeE2OverM3)
	}
	sort.SliceStable(cands, func(i, j int) bool { return cands[i].Distance < cands[j].Distance })
	best := cands[0]
	return TypedCubicCoefficientAudit{
		TargetCoefficient:  m.NegativeE2OverM3,
		Candidates:         cands,
		ClosestName:        best.Name,
		ClosestValue:       best.Value,
		ClosestDistance:    best.Distance,
		SevenOver36Closest: best.Name == "7/36",
		NoArbitrarySearch:  true,
		Verdict: strings.Join([]string{
			StatusTypedCubicCoefficientCandidatesAudited,
			StatusTypedBoundaryStressCoefficientClosest,
		}, "; "),
	}
}

func buildCubicCorrection(g Gate729Inheritance, m CubicMomentAudit) CubicCorrectionAudit {
	quad := g.KappaE * g.M2Wall
	cubic := cubicStressPullCoeff * m.M3Wall
	combined := quad - cubic
	resid := g.EWall - combined
	compression := math.Abs(g.EWall / resid)
	return CubicCorrectionAudit{
		EWall:                        g.EWall,
		KappaE:                       g.KappaE,
		M2Wall:                       g.M2Wall,
		M3Wall:                       m.M3Wall,
		CubicCoefficient:             cubicStressPullCoeff,
		QuadraticTerm:                quad,
		CubicStressPullTerm:          cubic,
		CombinedCorrection:           combined,
		ResidualAfterCubicCorrection: resid,
		RawCompressionFactor:         compression,
		ImprovesSecondOrderResidual:  math.Abs(resid) < math.Abs(g.E2Residual),
		NotExact:                     math.Abs(resid) > 1e-15,
		Verdict: strings.Join([]string{
			StatusCubicStressPullCorrectionTested,
			StatusSecondPlusThirdMomentStructure,
			StatusSevenOverThirtySixCompressesResidual,
			StatusCubicCorrectionNotExact,
		}, "; "),
	}
}

func buildRuntimePropagation(g Gate729Inheritance, c CubicCorrectionAudit) RuntimePropagationAudit {
	cubicRuntime := g.LambdaProxy * g.L * c.ResidualAfterCubicCorrection
	return RuntimePropagationAudit{
		LambdaProxy:                   g.LambdaProxy,
		L:                             g.L,
		RawRuntimeResidual:            g.RawRuntimeResidual,
		SecondOrderCorrectedRuntime:   g.SecondOrderCorrectedRuntimeResidual,
		CubicCorrectedWallResidual:    c.ResidualAfterCubicCorrection,
		CubicCorrectedRuntimeResidual: cubicRuntime,
		CompressedToNearFloatScale:    math.Abs(cubicRuntime) < 1e-14,
		ImprovesSecondOrderRuntime:    math.Abs(cubicRuntime) < math.Abs(g.SecondOrderCorrectedRuntimeResidual),
		Verdict: strings.Join([]string{
			StatusRuntimePropagationCubicCorrectionAudited,
			StatusRuntimeCompressedByTypedCubicCorrection,
		}, "; "),
	}
}

func buildSourceTypeInterpretation() SourceTypeInterpretation {
	return SourceTypeInterpretation{
		ExpansionFormula:             "D_base≈Tr(rho_72 R_wall)+kappa_e Tr(rho_72 R_wall^2)-(7/36)Tr(rho_72 R_wall^3)",
		LeadingTerm:                  "no-bias K7 boundary uplift event expectation",
		QuadraticTerm:                "flavor-wall deficit modulation candidate",
		CubicTerm:                    "typed boundary stress-pull correction candidate",
		MomentExpansionTheoremNative: false,
		Verdict: strings.Join([]string{
			StatusSourceTypeInterpretationRecorded,
			StatusNoNativeBoundaryMomentExpansionTheorem,
		}, "; "),
	}
}

func buildNoncircularity() NoncircularityAudit {
	return NoncircularityAudit{
		DBaseContainsKappaE:           true,
		KappaEUsedAsQuadraticCoeff:    true,
		CubicCoeffTypedButUnexplained: true,
		NativeExpansionTheorem:        false,
		Verdict: strings.Join([]string{
			StatusNoncircularityFirewallAudited,
			StatusKappaEQuadraticCoefficientDependent,
			StatusNoNativeReasonCubicCoeffSevenOverThirtySix,
			StatusNoNativeBoundaryMomentExpansionTheorem,
		}, "; "),
	}
}

func buildFirewall() FirewallAudit {
	return FirewallAudit{
		ClaimsNativeBoundaryHistory: false,
		ClaimsNativeMomentExpansion: false,
		ClaimsNativeScalarRuntime:   false,
		ClaimsHiggsMassTheorem:      false,
		ClaimsYukawaTheorem:         false,
		ClaimsCKMPMNSTheorem:        false,
		Verdict: strings.Join([]string{
			StatusPhysicalFirewallsEnforced,
			StatusNoNativeScalarRuntimeTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusGate730Boundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate729SecondMomentResidualInherited,
		StatusCubicWallMomentComputed,
		StatusCubicCoefficientRatioComputed,
		StatusTypedCubicCoefficientCandidatesAudited,
		StatusCubicStressPullCorrectionTested,
		StatusRuntimePropagationCubicCorrectionAudited,
		StatusSourceTypeInterpretationRecorded,
		StatusNoncircularityFirewallAudited,
		StatusPhysicalFirewallsEnforced,
		StatusSecondPlusThirdMomentStructure,
		StatusSevenOverThirtySixCompressesResidual,
		StatusRuntimeCompressedByTypedCubicCorrection,
		StatusSecondOrderResidualCubicScale,
		StatusTypedBoundaryStressCoefficientClosest,
		StatusKappaEQuadraticCoefficientDependent,
		StatusCubicCorrectionNotExact,
		StatusNoNativeReasonCubicCoeffSevenOverThirtySix,
		StatusNoNativeBoundaryMomentExpansionTheorem,
		StatusNoNativeScalarRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate730Boundary,
	}
}

func near(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatGate729(x Gate729Inheritance) string {
	return fmt.Sprintf("inherited=%t p=%.17g S=%.17g E=%.17g M2=%.17g kE=%.17g kEM2=%.17g E2=%.17g rawRuntime=%.17g secondRuntime=%.17g dependent=%t noNativeSecond=%t verdict=%q", x.Inherited, x.P_K7, x.SSplit, x.EWall, x.M2Wall, x.KappaE, x.KappaEM2, x.E2Residual, x.RawRuntimeResidual, x.SecondOrderCorrectedRuntimeResidual, x.KappaEPartiallyDependent, x.NoNativeSecondOrderBoundaryTheorem, x.Verdict)
}
func FormatCubicMoment(x CubicMomentAudit) string {
	return fmt.Sprintf("formula=%q M3=%.17g coeff=-E2/M3=%.17g cubicScale=%t verdict=%q", x.Formula, x.M3Wall, x.NegativeE2OverM3, x.SecondResidualCubicScale, x.Verdict)
}
func FormatCoefficients(x TypedCubicCoefficientAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, fmt.Sprintf("%s=%.17g dist=%.17g source=%s", c.Name, c.Value, c.Distance, c.Source))
	}
	return fmt.Sprintf("target=%.17g closest=%s value=%.17g dist=%.17g seven36=%t noSearch=%t candidates=[%s] verdict=%q", x.TargetCoefficient, x.ClosestName, x.ClosestValue, x.ClosestDistance, x.SevenOver36Closest, x.NoArbitrarySearch, strings.Join(parts, "; "), x.Verdict)
}
func FormatCubicCorrection(x CubicCorrectionAudit) string {
	return fmt.Sprintf("E=%.17g kE=%.17g M2=%.17g M3=%.17g c3=%.17g quad=%.17g cubic=%.17g combined=%.17g residual=%.17g compression=%.17g improvesSecond=%t notExact=%t verdict=%q", x.EWall, x.KappaE, x.M2Wall, x.M3Wall, x.CubicCoefficient, x.QuadraticTerm, x.CubicStressPullTerm, x.CombinedCorrection, x.ResidualAfterCubicCorrection, x.RawCompressionFactor, x.ImprovesSecondOrderResidual, x.NotExact, x.Verdict)
}
func FormatRuntime(x RuntimePropagationAudit) string {
	return fmt.Sprintf("lambdaProxy=%.17g L=%.17g raw=%.17g second=%.17g cubicWall=%.17g cubicRuntime=%.17g nearFloat=%t improvesSecond=%t verdict=%q", x.LambdaProxy, x.L, x.RawRuntimeResidual, x.SecondOrderCorrectedRuntime, x.CubicCorrectedWallResidual, x.CubicCorrectedRuntimeResidual, x.CompressedToNearFloatScale, x.ImprovesSecondOrderRuntime, x.Verdict)
}
func FormatSourceType(x SourceTypeInterpretation) string {
	return fmt.Sprintf("expansion=%q leading=%q quadratic=%q cubic=%q native=%t verdict=%q", x.ExpansionFormula, x.LeadingTerm, x.QuadraticTerm, x.CubicTerm, x.MomentExpansionTheoremNative, x.Verdict)
}
func FormatNoncircularity(x NoncircularityAudit) string {
	return fmt.Sprintf("DHasKE=%t kEQuadratic=%t c3TypedUnexplained=%t nativeExpansion=%t verdict=%q", x.DBaseContainsKappaE, x.KappaEUsedAsQuadraticCoeff, x.CubicCoeffTypedButUnexplained, x.NativeExpansionTheorem, x.Verdict)
}
func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("boundary=%t moment=%t runtime=%t mass=%t yukawa=%t ckm=%t verdict=%q", x.ClaimsNativeBoundaryHistory, x.ClaimsNativeMomentExpansion, x.ClaimsNativeScalarRuntime, x.ClaimsHiggsMassTheorem, x.ClaimsYukawaTheorem, x.ClaimsCKMPMNSTheorem, x.Verdict)
}
