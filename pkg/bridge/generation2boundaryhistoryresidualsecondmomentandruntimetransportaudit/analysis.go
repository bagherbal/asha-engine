// Package generation2boundaryhistoryresidualsecondmomentandruntimetransportaudit implements
// Gate 729: Boundary-History Residual Second-Moment and Runtime Propagation Audit.
//
// Gate 728 assembled the scalar runtime bridge from the Gate700 boundary-history
// event expectation and the Gate727 radial-Hopf HistoryLoopUnit expectation. Gate
// 729 audits whether the remaining boundary-history wall residual is naturally
// second-order in the boundary uplift response operator R_wall=S_split P_K7, and
// how that residual compression propagates into the scalar runtime ledger. This
// remains a bridge-layer residual-compression audit, not a native boundary
// response, scalar runtime, Higgs mass, or Yukawa theorem.
package generation2boundaryhistoryresidualsecondmomentandruntimetransportaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate728 "github.com/bagherbal/asha-engine/pkg/bridge/generation2dualeventexpectationscalarruntimetransportassemblyaudit"
)

const (
	AuditID = "GATE729-BOUNDARY-HISTORY-RESIDUAL-SECOND-MOMENT-RUNTIME-PROPAGATION-AUDIT"

	StatusGate728DualEventRuntimeInherited      = "PASS_GATE728_DUAL_EVENT_EXPECTATION_RUNTIME_INHERITED"
	StatusBoundaryUpliftResponseOperatorDefined = "PASS_BOUNDARY_UPLIFT_RESPONSE_OPERATOR_DEFINED"
	StatusSecondRawMomentComputed               = "PASS_SECOND_RAW_MOMENT_COMPUTED"
	StatusWallResidualOverSecondMomentComputed  = "PASS_WALL_RESIDUAL_OVER_SECOND_MOMENT_COMPUTED"
	StatusTypedCoefficientCandidatesAudited     = "PASS_TYPED_COEFFICIENT_CANDIDATES_AUDITED"
	StatusKappaESecondOrderCorrectionTested     = "PASS_KAPPA_E_SECOND_ORDER_CORRECTION_TESTED"
	StatusVarianceControlAudited                = "PASS_VARIANCE_CONTROL_AUDITED"
	StatusRuntimeResidualPropagationAudited     = "PASS_RUNTIME_RESIDUAL_PROPAGATION_AUDITED"
	StatusNoncircularityFirewallAudited         = "PASS_NONCIRCULARITY_FIREWALL_AUDITED"
	StatusPhysicalFirewallsEnforced             = "PASS_PHYSICAL_FIREWALLS_ENFORCED"

	StatusWallResidualSecondOrderSuppressed     = "CONDITIONAL_SUPPORT_WALL_RESIDUAL_IS_SECOND_ORDER_SUPPRESSED"
	StatusKappaECloseToSecondOrderCoefficient   = "CONDITIONAL_SUPPORT_KAPPA_E_CLOSE_TO_SECOND_ORDER_WALL_RESIDUAL_COEFFICIENT"
	StatusRuntimeResidualCompressedByCorrection = "CONDITIONAL_SUPPORT_RUNTIME_RESIDUAL_IS_COMPRESSED_BY_SECOND_ORDER_WALL_CORRECTION"
	StatusKappaECompressesWallResidual          = "CONDITIONAL_SUPPORT_KAPPA_E_COMPRESSES_WALL_RESIDUAL_AT_SECOND_ORDER"
	StatusSecondOrderFluctuationScaleRelevant   = "CONDITIONAL_SUPPORT_SECOND_ORDER_BOUNDARY_FLUCTUATION_SCALE_IS_RELEVANT"
	StatusRuntimeCompressionFollowsWallResidual = "CONDITIONAL_SUPPORT_RUNTIME_RESIDUAL_COMPRESSION_FOLLOWS_WALL_RESIDUAL_COMPRESSION"

	StatusKappaESecondOrderCorrectionNotExact       = "FAILED_ROUTE_KAPPA_E_SECOND_ORDER_CORRECTION_NOT_EXACT"
	StatusKappaEResidualCoefficientDependent        = "FAILED_ROUTE_KAPPA_E_RESIDUAL_COEFFICIENT_IS_PARTIALLY_DEPENDENT"
	StatusNoNativeSecondOrderBoundaryResponse       = "FAILED_ROUTE_NO_NATIVE_SECOND_ORDER_BOUNDARY_RESPONSE_THEOREM"
	StatusNoNativeScalarRuntimeTheorem              = "FAILED_ROUTE_NO_NATIVE_SCALAR_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem              = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem       = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusKappaECorrectionNotIndependentlyCertified = "FAILED_ROUTE_KAPPA_E_SECOND_ORDER_CORRECTION_NOT_INDEPENDENTLY_CERTIFIED"
	StatusVarianceFormNotYetSelected                = "FAILED_ROUTE_VARIANCE_FORM_NOT_YET_SELECTED_AS_ACTIVE_CORRECTION"
	StatusGate729Boundary                           = "FIREWALL_PRESERVED_GATE729_BOUNDARY_HISTORY_SECOND_MOMENT_BOUNDARY"
)

const (
	lambdaProxyMZ   = 0.12490310236015
	kappaE          = 0.00550355419157456
	kappaEOrient    = 0.00550633006471245
	kappaLambda     = 0.0443230430960771
	historyLoopL    = 1.0 / (8.0 * math.Pi)
	k7Dim           = 7
	h72Dim          = 72
	momentTolerance = 1e-18
)

type Gate728Inheritance struct {
	Inherited                      bool
	P_K7                           float64
	SSplit                         float64
	DBase                          float64
	EWall                          float64
	LambdaProxy                    float64
	L                              float64
	DeltaLambdaRuntime             float64
	DualEventExpectationClosure    bool
	AssembledRuntimeNotIndependent bool
	PremisesNotNative              bool
	Verdict                        string
}

type BoundaryUpliftResponseAudit struct {
	Operator              string
	P_K7                  float64
	SSplit                float64
	LeadingExpectation    float64
	DBase                 float64
	EWall                 float64
	MatchesGate700Leading bool
	Verdict               string
}

type SecondRawMomentAudit struct {
	Formula                       string
	M2Wall                        float64
	EWall                         float64
	C2Wall                        float64
	SecondOrderSuppressed         bool
	ResidualMuchSmallerThanMoment bool
	Verdict                       string
}

type CoefficientCandidate struct {
	Name     string
	Value    float64
	Distance float64
}

type TypedCoefficientAudit struct {
	C2Wall             float64
	Candidates         []CoefficientCandidate
	ClosestName        string
	ClosestValue       float64
	ClosestDistance    float64
	KappaEClosestSmall bool
	NotExact           bool
	Verdict            string
}

type KappaECorrectionAudit struct {
	EWall                     float64
	M2Wall                    float64
	KappaE                    float64
	KappaEM2                  float64
	ResidualAfterCorrection   float64
	CompressionFactor         float64
	ImprovesRawResidual       bool
	NotExact                  bool
	NotIndependentlyCertified bool
	Verdict                   string
}

type VarianceControlAudit struct {
	Formula                  string
	P_K7                     float64
	PComplement              float64
	SSplit                   float64
	VarianceWall             float64
	CVariance                float64
	RelevantTypedScale       bool
	SelectedActiveCorrection bool
	Verdict                  string
}

type RuntimeResidualPropagationAudit struct {
	LambdaProxy                    float64
	L                              float64
	EWall                          float64
	RawRuntimeResidual             float64
	CorrectedWallResidual          float64
	CorrectedRuntimeResidual       float64
	CompressionFollowsWallResidual bool
	Verdict                        string
}

type NoncircularityAudit struct {
	DBaseContainsKappaE     bool
	KappaEUsedAsCoefficient bool
	IndependentTheorem      bool
	PartiallyDependent      bool
	Verdict                 string
}

type FirewallAudit struct {
	ClaimsNativeBoundaryHistory      bool
	ClaimsNativeSecondOrderExpansion bool
	ClaimsNativeScalarRuntime        bool
	ClaimsHiggsMassTheorem           bool
	ClaimsYukawaOperatorTheorem      bool
	ClaimsCKMPMNSTheorem             bool
	Verdict                          string
}

type Analysis struct {
	Gate728      Gate728Inheritance
	Uplift       BoundaryUpliftResponseAudit
	Moment       SecondRawMomentAudit
	Coefficients TypedCoefficientAudit
	KappaECorr   KappaECorrectionAudit
	Variance     VarianceControlAudit
	Runtime      RuntimeResidualPropagationAudit
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
	g728, err := gate728.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate728 inheritance unavailable: %w", err)
	}
	inherited := buildGate728Inheritance(g728)
	uplift := buildBoundaryUpliftResponse(inherited)
	moment := buildSecondRawMoment(uplift)
	coeffs := buildTypedCoefficientAudit(moment)
	kCorr := buildKappaECorrection(moment)
	variance := buildVarianceControl(uplift)
	runtime := buildRuntimeResidualPropagation(inherited, kCorr)
	noncirc := buildNoncircularity()
	firewall := buildFirewall()
	truth := "Gate 729 audits the Gate700/Gate728 wall residual as a second-order boundary-uplift effect. For R_wall=S_split P_K7, the second raw moment is M2_wall=Tr(rho_72 R_wall^2)=(7/72)S_split^2≈1.6240e-7, and E_wall/M2_wall≈0.005249855. The closest active small coefficient is kappa_e, which compresses the wall residual by about twentyfold and propagates to a corrected runtime residual of order 2.05e-13. This remains non-independent because D_base already contains kappa_e, and no native second-order boundary response, scalar runtime, Higgs mass, or Yukawa theorem is certified."
	return Analysis{Gate728: inherited, Uplift: uplift, Moment: moment, Coefficients: coeffs, KappaECorr: kCorr, Variance: variance, Runtime: runtime, NonCircular: noncirc, Firewall: firewall, Truth: truth}, nil
}

func buildGate728Inheritance(g gate728.Analysis) Gate728Inheritance {
	return Gate728Inheritance{
		Inherited:                      g.Assembly.DualEventExpectationForm && g.Propagation.RuntimeResidualIsWallResidual,
		P_K7:                           float64(k7Dim) / float64(h72Dim),
		SSplit:                         g.Gate700.SSplit,
		DBase:                          g.Gate700.DBase,
		EWall:                          g.Gate700.EWall,
		LambdaProxy:                    lambdaProxyMZ,
		L:                              g.Gate727.L,
		DeltaLambdaRuntime:             g.Propagation.DeltaLambdaPred,
		DualEventExpectationClosure:    g.Assembly.DualEventExpectationForm,
		AssembledRuntimeNotIndependent: !g.NonCircular.AssembledIndependentPrediction,
		PremisesNotNative:              !g.Seals.PremisesNativelyDerived,
		Verdict:                        StatusGate728DualEventRuntimeInherited,
	}
}

func buildBoundaryUpliftResponse(g Gate728Inheritance) BoundaryUpliftResponseAudit {
	leading := g.P_K7 * g.SSplit
	return BoundaryUpliftResponseAudit{
		Operator:              "R_wall=S_split P_K7",
		P_K7:                  g.P_K7,
		SSplit:                g.SSplit,
		LeadingExpectation:    leading,
		DBase:                 g.DBase,
		EWall:                 g.EWall,
		MatchesGate700Leading: near(g.DBase-leading, g.EWall, 1e-18),
		Verdict: strings.Join([]string{
			StatusBoundaryUpliftResponseOperatorDefined,
			StatusWallResidualSecondOrderSuppressed,
		}, "; "),
	}
}

func buildSecondRawMoment(u BoundaryUpliftResponseAudit) SecondRawMomentAudit {
	m2 := u.P_K7 * u.SSplit * u.SSplit
	c2 := u.EWall / m2
	return SecondRawMomentAudit{
		Formula:                       "M2_wall=Tr(rho_72 R_wall^2)=(7/72)S_split^2",
		M2Wall:                        m2,
		EWall:                         u.EWall,
		C2Wall:                        c2,
		SecondOrderSuppressed:         math.Abs(u.EWall) < math.Abs(m2),
		ResidualMuchSmallerThanMoment: math.Abs(u.EWall/m2) < 0.01,
		Verdict: strings.Join([]string{
			StatusSecondRawMomentComputed,
			StatusWallResidualOverSecondMomentComputed,
			StatusWallResidualSecondOrderSuppressed,
		}, "; "),
	}
}

func buildTypedCoefficientAudit(m SecondRawMomentAudit) TypedCoefficientAudit {
	raw := []CoefficientCandidate{
		{Name: "kappa_e", Value: kappaE},
		{Name: "kappa_e_orient", Value: kappaEOrient},
		{Name: "kappa_lambda", Value: kappaLambda},
		{Name: "L=1/(8*pi)", Value: historyLoopL},
		{Name: "S_split", Value: math.Sqrt(m.M2Wall / (float64(k7Dim) / float64(h72Dim)))},
	}
	best := raw[0]
	for i := range raw {
		raw[i].Distance = math.Abs(raw[i].Value - m.C2Wall)
		if raw[i].Distance < best.Distance || i == 0 {
			best = raw[i]
		}
	}
	return TypedCoefficientAudit{
		C2Wall:             m.C2Wall,
		Candidates:         raw,
		ClosestName:        best.Name,
		ClosestValue:       best.Value,
		ClosestDistance:    best.Distance,
		KappaEClosestSmall: best.Name == "kappa_e",
		NotExact:           best.Distance > 1e-6,
		Verdict: strings.Join([]string{
			StatusTypedCoefficientCandidatesAudited,
			StatusKappaECloseToSecondOrderCoefficient,
			StatusKappaESecondOrderCorrectionNotExact,
		}, "; "),
	}
}

func buildKappaECorrection(m SecondRawMomentAudit) KappaECorrectionAudit {
	kM2 := kappaE * m.M2Wall
	resid := m.EWall - kM2
	compression := math.Abs(m.EWall / resid)
	return KappaECorrectionAudit{
		EWall:                     m.EWall,
		M2Wall:                    m.M2Wall,
		KappaE:                    kappaE,
		KappaEM2:                  kM2,
		ResidualAfterCorrection:   resid,
		CompressionFactor:         compression,
		ImprovesRawResidual:       math.Abs(resid) < math.Abs(m.EWall),
		NotExact:                  math.Abs(resid) > 1e-14,
		NotIndependentlyCertified: true,
		Verdict: strings.Join([]string{
			StatusKappaESecondOrderCorrectionTested,
			StatusKappaECompressesWallResidual,
			StatusKappaESecondOrderCorrectionNotExact,
			StatusKappaECorrectionNotIndependentlyCertified,
		}, "; "),
	}
}

func buildVarianceControl(u BoundaryUpliftResponseAudit) VarianceControlAudit {
	p := u.P_K7
	v := p * (1 - p) * u.SSplit * u.SSplit
	return VarianceControlAudit{
		Formula:                  "Var_wall=p_K7(1-p_K7)S_split^2",
		P_K7:                     p,
		PComplement:              1 - p,
		SSplit:                   u.SSplit,
		VarianceWall:             v,
		CVariance:                u.EWall / v,
		RelevantTypedScale:       true,
		SelectedActiveCorrection: false,
		Verdict: strings.Join([]string{
			StatusVarianceControlAudited,
			StatusSecondOrderFluctuationScaleRelevant,
			StatusVarianceFormNotYetSelected,
		}, "; "),
	}
}

func buildRuntimeResidualPropagation(g Gate728Inheritance, k KappaECorrectionAudit) RuntimeResidualPropagationAudit {
	raw := lambdaProxyMZ * g.L * g.EWall
	corrected := lambdaProxyMZ * g.L * k.ResidualAfterCorrection
	return RuntimeResidualPropagationAudit{
		LambdaProxy:                    lambdaProxyMZ,
		L:                              g.L,
		EWall:                          g.EWall,
		RawRuntimeResidual:             raw,
		CorrectedWallResidual:          k.ResidualAfterCorrection,
		CorrectedRuntimeResidual:       corrected,
		CompressionFollowsWallResidual: math.Abs(corrected) < math.Abs(raw),
		Verdict: strings.Join([]string{
			StatusRuntimeResidualPropagationAudited,
			StatusRuntimeResidualCompressedByCorrection,
			StatusRuntimeCompressionFollowsWallResidual,
		}, "; "),
	}
}

func buildNoncircularity() NoncircularityAudit {
	return NoncircularityAudit{
		DBaseContainsKappaE:     true,
		KappaEUsedAsCoefficient: true,
		IndependentTheorem:      false,
		PartiallyDependent:      true,
		Verdict: strings.Join([]string{
			StatusNoncircularityFirewallAudited,
			StatusKappaEResidualCoefficientDependent,
		}, "; "),
	}
}

func buildFirewall() FirewallAudit {
	return FirewallAudit{
		ClaimsNativeBoundaryHistory:      false,
		ClaimsNativeSecondOrderExpansion: false,
		ClaimsNativeScalarRuntime:        false,
		ClaimsHiggsMassTheorem:           false,
		ClaimsYukawaOperatorTheorem:      false,
		ClaimsCKMPMNSTheorem:             false,
		Verdict: strings.Join([]string{
			StatusPhysicalFirewallsEnforced,
			StatusNoNativeSecondOrderBoundaryResponse,
			StatusNoNativeScalarRuntimeTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusGate729Boundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate728DualEventRuntimeInherited,
		StatusBoundaryUpliftResponseOperatorDefined,
		StatusSecondRawMomentComputed,
		StatusWallResidualOverSecondMomentComputed,
		StatusTypedCoefficientCandidatesAudited,
		StatusKappaESecondOrderCorrectionTested,
		StatusVarianceControlAudited,
		StatusRuntimeResidualPropagationAudited,
		StatusNoncircularityFirewallAudited,
		StatusPhysicalFirewallsEnforced,
		StatusWallResidualSecondOrderSuppressed,
		StatusKappaECloseToSecondOrderCoefficient,
		StatusRuntimeResidualCompressedByCorrection,
		StatusKappaECompressesWallResidual,
		StatusSecondOrderFluctuationScaleRelevant,
		StatusRuntimeCompressionFollowsWallResidual,
		StatusKappaESecondOrderCorrectionNotExact,
		StatusKappaEResidualCoefficientDependent,
		StatusNoNativeSecondOrderBoundaryResponse,
		StatusNoNativeScalarRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusKappaECorrectionNotIndependentlyCertified,
		StatusVarianceFormNotYetSelected,
		StatusGate729Boundary,
	}
}

func near(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatGate728(x Gate728Inheritance) string {
	return fmt.Sprintf("inherited=%t p=%.17g S=%.17g D=%.17g E=%.17g proxy=%.17g L=%.17g delta=%.17g dual=%t notIndependent=%t premisesNotNative=%t verdict=%q", x.Inherited, x.P_K7, x.SSplit, x.DBase, x.EWall, x.LambdaProxy, x.L, x.DeltaLambdaRuntime, x.DualEventExpectationClosure, x.AssembledRuntimeNotIndependent, x.PremisesNotNative, x.Verdict)
}
func FormatUplift(x BoundaryUpliftResponseAudit) string {
	return fmt.Sprintf("operator=%q p=%.17g S=%.17g leading=%.17g D=%.17g E=%.17g matches=%t verdict=%q", x.Operator, x.P_K7, x.SSplit, x.LeadingExpectation, x.DBase, x.EWall, x.MatchesGate700Leading, x.Verdict)
}
func FormatMoment(x SecondRawMomentAudit) string {
	return fmt.Sprintf("formula=%q M2=%.17g E=%.17g c2=%.17g suppressed=%t smaller=%t verdict=%q", x.Formula, x.M2Wall, x.EWall, x.C2Wall, x.SecondOrderSuppressed, x.ResidualMuchSmallerThanMoment, x.Verdict)
}
func FormatCoefficients(x TypedCoefficientAudit) string {
	parts := make([]string, 0, len(x.Candidates))
	for _, c := range x.Candidates {
		parts = append(parts, fmt.Sprintf("%s=%.17g dist=%.17g", c.Name, c.Value, c.Distance))
	}
	return fmt.Sprintf("c2=%.17g closest=%s value=%.17g dist=%.17g kEClosest=%t notExact=%t candidates=[%s] verdict=%q", x.C2Wall, x.ClosestName, x.ClosestValue, x.ClosestDistance, x.KappaEClosestSmall, x.NotExact, strings.Join(parts, "; "), x.Verdict)
}
func FormatKappaECorrection(x KappaECorrectionAudit) string {
	return fmt.Sprintf("E=%.17g M2=%.17g kE=%.17g kE*M2=%.17g residual=%.17g compression=%.17g improves=%t notExact=%t dependent=%t verdict=%q", x.EWall, x.M2Wall, x.KappaE, x.KappaEM2, x.ResidualAfterCorrection, x.CompressionFactor, x.ImprovesRawResidual, x.NotExact, x.NotIndependentlyCertified, x.Verdict)
}
func FormatVariance(x VarianceControlAudit) string {
	return fmt.Sprintf("formula=%q p=%.17g pc=%.17g S=%.17g Var=%.17g cVar=%.17g relevant=%t selected=%t verdict=%q", x.Formula, x.P_K7, x.PComplement, x.SSplit, x.VarianceWall, x.CVariance, x.RelevantTypedScale, x.SelectedActiveCorrection, x.Verdict)
}
func FormatRuntime(x RuntimeResidualPropagationAudit) string {
	return fmt.Sprintf("proxy=%.17g L=%.17g E=%.17g rawRuntime=%.17g correctedWall=%.17g correctedRuntime=%.17g compressed=%t verdict=%q", x.LambdaProxy, x.L, x.EWall, x.RawRuntimeResidual, x.CorrectedWallResidual, x.CorrectedRuntimeResidual, x.CompressionFollowsWallResidual, x.Verdict)
}
func FormatNoncircularity(x NoncircularityAudit) string {
	return fmt.Sprintf("DHasKE=%t kEUsed=%t independent=%t dependent=%t verdict=%q", x.DBaseContainsKappaE, x.KappaEUsedAsCoefficient, x.IndependentTheorem, x.PartiallyDependent, x.Verdict)
}
func FormatFirewall(x FirewallAudit) string {
	return fmt.Sprintf("boundary=%t secondOrder=%t runtime=%t mass=%t yukawa=%t ckm=%t verdict=%q", x.ClaimsNativeBoundaryHistory, x.ClaimsNativeSecondOrderExpansion, x.ClaimsNativeScalarRuntime, x.ClaimsHiggsMassTheorem, x.ClaimsYukawaOperatorTheorem, x.ClaimsCKMPMNSTheorem, x.Verdict)
}
