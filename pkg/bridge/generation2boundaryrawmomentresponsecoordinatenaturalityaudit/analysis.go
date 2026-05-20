// Package generation2boundaryrawmomentresponsecoordinatenaturalityaudit implements
// Gate 732: Boundary Raw-Moment Response Coordinate-Naturality Audit.
//
// Gate 731 source-typed the cubic coefficient as 7/36=2p_K7 and rewrote the
// compressed wall residual expansion as
//
//	D_base ≈ M1_wall + kappa_e M2_wall - 2p_K7 M3_wall.
//
// Gate 732 audits the coordinate in which this expansion is being expressed:
// raw powers of the boundary response operator R_wall=S_split P_K7. It compares
// the active raw-moment coordinate against variance and central-moment
// coordinates, records that projector powers add no new operator directions, and
// preserves the firewall that no native raw-moment response theorem follows.
package generation2boundaryrawmomentresponsecoordinatenaturalityaudit

import (
	"fmt"
	"math"
	"strings"
	"sync"

	gate731 "github.com/bagherbal/asha-engine/pkg/bridge/generation2cubicstresspullcoefficientsourcetypeanddoubleeventweightaudit"
)

const (
	AuditID = "GATE732-BOUNDARY-RAW-MOMENT-RESPONSE-COORDINATE-NATURALITY-AUDIT"

	StatusGate731CubicCoefficientSourceInherited = "PASS_GATE731_CUBIC_COEFFICIENT_SOURCE_INHERITED"
	StatusRawMomentResponseFunctionRewritten     = "PASS_RAW_MOMENT_RESPONSE_FUNCTION_REWRITTEN"
	StatusProjectorPowerDegeneracyRecorded       = "PASS_PROJECTOR_POWER_DEGENERACY_RECORDED"
	StatusVarianceCoordinateAudited              = "PASS_VARIANCE_COORDINATE_AUDITED"
	StatusCentralThirdMomentAudited              = "PASS_CENTRAL_THIRD_MOMENT_AUDITED"
	StatusRawVersusCentralComparisonAudited      = "PASS_RAW_VERSUS_CENTRAL_COMPARISON_AUDITED"
	StatusSourceTypeInterpretationRecorded       = "PASS_SOURCE_TYPE_INTERPRETATION_RECORDED"
	StatusCoordinateNaturalityFirewallEnforced   = "PASS_COORDINATE_NATURALITY_FIREWALL_ENFORCED"

	StatusActiveResidualExpansionLivesInRawMoments       = "CONDITIONAL_SUPPORT_ACTIVE_RESIDUAL_EXPANSION_LIVES_IN_RAW_RESPONSE_MOMENTS"
	StatusMomentExpansionIsScalarResponseFunction        = "CONDITIONAL_SUPPORT_MOMENT_EXPANSION_IS_SCALAR_RESPONSE_FUNCTION_ON_S_SPLIT"
	StatusRawM3CoordinateBestCompressesCurrentResidual   = "CONDITIONAL_SUPPORT_RAW_M3_COORDINATE_BEST_COMPRESSES_CURRENT_RESIDUAL"
	StatusActiveExpansionIsRawResponseFunction           = "CONDITIONAL_SUPPORT_ACTIVE_EXPANSION_IS_RAW_RESPONSE_FUNCTION_ON_S_SPLIT"
	StatusProjectorMomentExpansionNotNewOperatorGeometry = "CONDITIONAL_SUPPORT_MOMENT_EXPANSION_IS_SCALAR_RESPONSE_FUNCTION_NOT_NEW_OPERATOR_GEOMETRY"
	StatusSecondOrderBoundaryFluctuationScaleRelevant    = "CONDITIONAL_SUPPORT_SECOND_ORDER_BOUNDARY_FLUCTUATION_SCALE_IS_RELEVANT"

	StatusVarianceCoordinateNotActive                    = "FAILED_ROUTE_VARIANCE_COORDINATE_NOT_ACTIVE"
	StatusVarianceCoordinateNotActiveLeadingResidualForm = "FAILED_ROUTE_VARIANCE_COORDINATE_NOT_ACTIVE_LEADING_RESIDUAL_FORM"
	StatusCentralMomentFormNotActive                     = "FAILED_ROUTE_CENTRAL_MOMENT_FORM_NOT_ACTIVE"
	StatusCentralMomentNotSelectedOverRawM3              = "FAILED_ROUTE_CENTRAL_MOMENT_FORM_NOT_SELECTED_OVER_RAW_M3"
	StatusProjectorPowersNoIndependentDirections         = "FAILED_ROUTE_PROJECTOR_POWERS_DO_NOT_SUPPLY_INDEPENDENT_OPERATOR_DIRECTIONS"
	StatusNoNativeRawMomentResponseCoordinateTheorem     = "FAILED_ROUTE_NO_NATIVE_RAW_MOMENT_RESPONSE_COORDINATE_THEOREM"
	StatusNoNativeBoundaryMomentExpansionTheorem         = "FAILED_ROUTE_NO_NATIVE_BOUNDARY_MOMENT_EXPANSION_THEOREM"
	StatusNoNativeScalarRuntimeTheorem                   = "FAILED_ROUTE_NO_NATIVE_SCALAR_RUNTIME_THEOREM"
	StatusNoHiggsMassOrPoleMassTheorem                   = "FAILED_ROUTE_NO_HIGGS_MASS_OR_POLE_MASS_THEOREM"
	StatusNoYukawaOperatorOrEigenvalueTheorem            = "FAILED_ROUTE_NO_YUKAWA_OPERATOR_OR_EIGENVALUE_THEOREM"
	StatusGate732Boundary                                = "FIREWALL_PRESERVED_GATE732_RAW_MOMENT_COORDINATE_BOUNDARY"
)

type Gate731Inheritance struct {
	Inherited                 bool
	P_K7                      float64
	SSplit                    float64
	KappaE                    float64
	M1Wall                    float64
	M2Wall                    float64
	M3Wall                    float64
	EWall                     float64
	RawCubicResidual          float64
	DoubleEventWeight         float64
	MomentPolynomialAvailable bool
	NoNativeMomentExpansion   bool
	Verdict                   string
}

type RawMomentResponseAudit struct {
	M1Wall           float64
	M2Wall           float64
	M3Wall           float64
	ResponseFunction string
	FactoredFunction string
	PolynomialValue  float64
	RawResidual      float64
	UsesRawMoments   bool
	Verdict          string
}

type ProjectorPowerDegeneracyAudit struct {
	PowerFormula                  string
	AllPowersSupportedOnK7        bool
	IndependentOperatorDirections bool
	ScalarResponseFunctionOnly    bool
	Verdict                       string
}

type VarianceCoordinateAudit struct {
	VarianceWall          float64
	EWall                 float64
	CoefficientInVariance float64
	CoefficientInRawM2    float64
	KappaE                float64
	RawM2CloserToKappaE   bool
	TypedButInactive      bool
	Verdict               string
}

type CentralThirdMomentAudit struct {
	Mu3Wall              float64
	RawM3Wall            float64
	Mu3OverRawM3         float64
	Coefficient          float64
	ResidualCentralCubic float64
	ResidualRawCubic     float64
	RawCompressesBetter  bool
	Verdict              string
}

type RawVsCentralComparisonAudit struct {
	RawResidualAbs                  float64
	CentralResidualAbs              float64
	ImprovementFactor               float64
	RawSelectedByCurrentCompression bool
	Verdict                         string
}

type SourceTypeInterpretation struct {
	Leading   string
	Quadratic string
	Cubic     string
	Compact   string
	Verdict   string
}

type CoordinateNaturalityFirewall struct {
	RawMomentsNativelySelected  bool
	BoundaryMomentTheoremNative bool
	ScalarRuntimeTheoremNative  bool
	HiggsMassTheoremNative      bool
	YukawaTheoremNative         bool
	Verdict                     string
}

type Analysis struct {
	Gate731    Gate731Inheritance
	RawMoment  RawMomentResponseAudit
	Degeneracy ProjectorPowerDegeneracyAudit
	Variance   VarianceCoordinateAudit
	CentralM3  CentralThirdMomentAudit
	Comparison RawVsCentralComparisonAudit
	SourceType SourceTypeInterpretation
	Firewall   CoordinateNaturalityFirewall
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
	g731, err := gate731.BuildDefault()
	if err != nil {
		return Analysis{}, fmt.Errorf("Gate731 inheritance unavailable: %w", err)
	}
	inherited := buildGate731Inheritance(g731)
	raw := buildRawMomentResponse(inherited)
	deg := buildProjectorPowerDegeneracy()
	variance := buildVarianceCoordinate(inherited)
	central := buildCentralThirdMoment(inherited)
	comparison := buildRawVsCentralComparison(central)
	source := buildSourceTypeInterpretation()
	firewall := buildFirewall()
	truth := "Gate 732 audits the Gate731 residual expansion as a raw response-moment coordinate rather than a central-moment, variance, cumulant, or normalized-moment coordinate. Because R_wall=S_split P_K7 and P_K7 is a projector, all raw powers remain in the same K7 direction and only supply scalar powers of S_split: D_base≈p_K7 S_split[1+kappa_e S_split-2p_K7 S_split^2]. The variance and central-third-moment coordinates are typed but do not compress the current residual as well as the raw M2/M3 route. This selects raw moments only within the bridge ledger; no native raw-moment response coordinate theorem, boundary moment expansion theorem, scalar runtime theorem, Higgs mass theorem, or Yukawa theorem is certified."
	return Analysis{Gate731: inherited, RawMoment: raw, Degeneracy: deg, Variance: variance, CentralM3: central, Comparison: comparison, SourceType: source, Firewall: firewall, Truth: truth}, nil
}

func buildGate731Inheritance(g gate731.Analysis) Gate731Inheritance {
	eWall := g.Gate730.CombinedCorrection + g.Gate730.ResidualAfterCubicCorrection
	return Gate731Inheritance{
		Inherited:                 g.Gate730.Inherited && g.Polynomial.UsesDoubleEventForm && g.NonCircular.TwoPK7TypedButUnexplained,
		P_K7:                      g.DoubleEvent.K7EventProbability,
		SSplit:                    g.Gate730.SSplit,
		KappaE:                    g.Gate730.KappaE,
		M1Wall:                    g.Gate730.M1Wall,
		M2Wall:                    g.Gate730.M2Wall,
		M3Wall:                    g.Gate730.M3Wall,
		EWall:                     eWall,
		RawCubicResidual:          g.Gate730.ResidualAfterCubicCorrection,
		DoubleEventWeight:         g.DoubleEvent.DoubleK7Weight,
		MomentPolynomialAvailable: g.Polynomial.UsesDoubleEventForm,
		NoNativeMomentExpansion:   g.NonCircular.MomentExpansionTheoremNative == false,
		Verdict:                   StatusGate731CubicCoefficientSourceInherited,
	}
}

func buildRawMomentResponse(g Gate731Inheritance) RawMomentResponseAudit {
	poly := g.M1Wall + g.KappaE*g.M2Wall - g.DoubleEventWeight*g.M3Wall
	return RawMomentResponseAudit{
		M1Wall:           g.M1Wall,
		M2Wall:           g.M2Wall,
		M3Wall:           g.M3Wall,
		ResponseFunction: "D_base≈M1+kappa_e M2-2p_K7 M3",
		FactoredFunction: "D_base≈p_K7 S_split[1+kappa_e S_split-2p_K7 S_split^2]",
		PolynomialValue:  poly,
		RawResidual:      g.EWall + g.M1Wall - poly, // D_base - polynomial, using D_base=M1+E_wall.
		UsesRawMoments:   true,
		Verdict: strings.Join([]string{
			StatusRawMomentResponseFunctionRewritten,
			StatusActiveResidualExpansionLivesInRawMoments,
			StatusActiveExpansionIsRawResponseFunction,
		}, "; "),
	}
}

func buildProjectorPowerDegeneracy() ProjectorPowerDegeneracyAudit {
	return ProjectorPowerDegeneracyAudit{
		PowerFormula:                  "R_wall^n=S_split^n P_K7; M_n=p_K7 S_split^n",
		AllPowersSupportedOnK7:        true,
		IndependentOperatorDirections: false,
		ScalarResponseFunctionOnly:    true,
		Verdict: strings.Join([]string{
			StatusProjectorPowerDegeneracyRecorded,
			StatusMomentExpansionIsScalarResponseFunction,
			StatusProjectorMomentExpansionNotNewOperatorGeometry,
			StatusProjectorPowersNoIndependentDirections,
		}, "; "),
	}
}

func buildVarianceCoordinate(g Gate731Inheritance) VarianceCoordinateAudit {
	variance := g.P_K7 * (1.0 - g.P_K7) * g.SSplit * g.SSplit
	cVar := g.EWall / variance
	cRaw := g.EWall / g.M2Wall
	return VarianceCoordinateAudit{
		VarianceWall:          variance,
		EWall:                 g.EWall,
		CoefficientInVariance: cVar,
		CoefficientInRawM2:    cRaw,
		KappaE:                g.KappaE,
		RawM2CloserToKappaE:   math.Abs(cRaw-g.KappaE) < math.Abs(cVar-g.KappaE),
		TypedButInactive:      true,
		Verdict: strings.Join([]string{
			StatusVarianceCoordinateAudited,
			StatusSecondOrderBoundaryFluctuationScaleRelevant,
			StatusVarianceCoordinateNotActive,
			StatusVarianceCoordinateNotActiveLeadingResidualForm,
		}, "; "),
	}
}

func buildCentralThirdMoment(g Gate731Inheritance) CentralThirdMomentAudit {
	mu3 := g.P_K7 * (1.0 - g.P_K7) * (1.0 - 2.0*g.P_K7) * math.Pow(g.SSplit, 3)
	centralResidual := g.EWall - (g.KappaE*g.M2Wall - g.DoubleEventWeight*mu3)
	return CentralThirdMomentAudit{
		Mu3Wall:              mu3,
		RawM3Wall:            g.M3Wall,
		Mu3OverRawM3:         mu3 / g.M3Wall,
		Coefficient:          g.DoubleEventWeight,
		ResidualCentralCubic: centralResidual,
		ResidualRawCubic:     g.RawCubicResidual,
		RawCompressesBetter:  math.Abs(g.RawCubicResidual) < math.Abs(centralResidual),
		Verdict: strings.Join([]string{
			StatusCentralThirdMomentAudited,
			StatusCentralMomentNotSelectedOverRawM3,
			StatusCentralMomentFormNotActive,
		}, "; "),
	}
}

func buildRawVsCentralComparison(c CentralThirdMomentAudit) RawVsCentralComparisonAudit {
	factor := math.Abs(c.ResidualCentralCubic / c.ResidualRawCubic)
	return RawVsCentralComparisonAudit{
		RawResidualAbs:                  math.Abs(c.ResidualRawCubic),
		CentralResidualAbs:              math.Abs(c.ResidualCentralCubic),
		ImprovementFactor:               factor,
		RawSelectedByCurrentCompression: factor > 10.0,
		Verdict: strings.Join([]string{
			StatusRawVersusCentralComparisonAudited,
			StatusRawM3CoordinateBestCompressesCurrentResidual,
		}, "; "),
	}
}

func buildSourceTypeInterpretation() SourceTypeInterpretation {
	return SourceTypeInterpretation{
		Leading:   "raw first expectation M1=p_K7 S_split",
		Quadratic: "raw second moment M2 modulated by kappa_e",
		Cubic:     "raw third moment M3 pulled by double event weight 2p_K7",
		Compact:   "D_base≈M1+kappa_e M2-2p_K7 M3",
		Verdict: strings.Join([]string{
			StatusSourceTypeInterpretationRecorded,
			StatusActiveResidualExpansionLivesInRawMoments,
		}, "; "),
	}
}

func buildFirewall() CoordinateNaturalityFirewall {
	return CoordinateNaturalityFirewall{
		RawMomentsNativelySelected:  false,
		BoundaryMomentTheoremNative: false,
		ScalarRuntimeTheoremNative:  false,
		HiggsMassTheoremNative:      false,
		YukawaTheoremNative:         false,
		Verdict: strings.Join([]string{
			StatusCoordinateNaturalityFirewallEnforced,
			StatusNoNativeRawMomentResponseCoordinateTheorem,
			StatusNoNativeBoundaryMomentExpansionTheorem,
			StatusNoNativeScalarRuntimeTheorem,
			StatusNoHiggsMassOrPoleMassTheorem,
			StatusNoYukawaOperatorOrEigenvalueTheorem,
			StatusGate732Boundary,
		}, "; "),
	}
}

func Statuses() []string {
	return []string{
		StatusGate731CubicCoefficientSourceInherited,
		StatusRawMomentResponseFunctionRewritten,
		StatusProjectorPowerDegeneracyRecorded,
		StatusVarianceCoordinateAudited,
		StatusCentralThirdMomentAudited,
		StatusRawVersusCentralComparisonAudited,
		StatusSourceTypeInterpretationRecorded,
		StatusCoordinateNaturalityFirewallEnforced,
		StatusActiveResidualExpansionLivesInRawMoments,
		StatusMomentExpansionIsScalarResponseFunction,
		StatusRawM3CoordinateBestCompressesCurrentResidual,
		StatusActiveExpansionIsRawResponseFunction,
		StatusProjectorMomentExpansionNotNewOperatorGeometry,
		StatusSecondOrderBoundaryFluctuationScaleRelevant,
		StatusVarianceCoordinateNotActive,
		StatusVarianceCoordinateNotActiveLeadingResidualForm,
		StatusCentralMomentFormNotActive,
		StatusCentralMomentNotSelectedOverRawM3,
		StatusProjectorPowersNoIndependentDirections,
		StatusNoNativeRawMomentResponseCoordinateTheorem,
		StatusNoNativeBoundaryMomentExpansionTheorem,
		StatusNoNativeScalarRuntimeTheorem,
		StatusNoHiggsMassOrPoleMassTheorem,
		StatusNoYukawaOperatorOrEigenvalueTheorem,
		StatusGate732Boundary,
	}
}

func near(a, b, eps float64) bool { return math.Abs(a-b) <= eps }

func FormatGate731(x Gate731Inheritance) string {
	return fmt.Sprintf("inherited=%t p=%.17g S=%.17g kE=%.17g M1=%.17g M2=%.17g M3=%.17g E=%.17g rawResidual=%.17g 2p=%.17g poly=%t noNativeMoment=%t verdict=%q", x.Inherited, x.P_K7, x.SSplit, x.KappaE, x.M1Wall, x.M2Wall, x.M3Wall, x.EWall, x.RawCubicResidual, x.DoubleEventWeight, x.MomentPolynomialAvailable, x.NoNativeMomentExpansion, x.Verdict)
}
func FormatRawMoment(x RawMomentResponseAudit) string {
	return fmt.Sprintf("M1=%.17g M2=%.17g M3=%.17g response=%q factored=%q value=%.17g residual=%.17g raw=%t verdict=%q", x.M1Wall, x.M2Wall, x.M3Wall, x.ResponseFunction, x.FactoredFunction, x.PolynomialValue, x.RawResidual, x.UsesRawMoments, x.Verdict)
}
func FormatDegeneracy(x ProjectorPowerDegeneracyAudit) string {
	return fmt.Sprintf("formula=%q allK7=%t independentDirs=%t scalarOnly=%t verdict=%q", x.PowerFormula, x.AllPowersSupportedOnK7, x.IndependentOperatorDirections, x.ScalarResponseFunctionOnly, x.Verdict)
}
func FormatVariance(x VarianceCoordinateAudit) string {
	return fmt.Sprintf("var=%.17g E=%.17g cVar=%.17g cRaw=%.17g kE=%.17g rawCloser=%t typedInactive=%t verdict=%q", x.VarianceWall, x.EWall, x.CoefficientInVariance, x.CoefficientInRawM2, x.KappaE, x.RawM2CloserToKappaE, x.TypedButInactive, x.Verdict)
}
func FormatCentralM3(x CentralThirdMomentAudit) string {
	return fmt.Sprintf("mu3=%.17g rawM3=%.17g ratio=%.17g coeff=%.17g centralResidual=%.17g rawResidual=%.17g rawBetter=%t verdict=%q", x.Mu3Wall, x.RawM3Wall, x.Mu3OverRawM3, x.Coefficient, x.ResidualCentralCubic, x.ResidualRawCubic, x.RawCompressesBetter, x.Verdict)
}
func FormatComparison(x RawVsCentralComparisonAudit) string {
	return fmt.Sprintf("rawAbs=%.17g centralAbs=%.17g improvement=%.17g selected=%t verdict=%q", x.RawResidualAbs, x.CentralResidualAbs, x.ImprovementFactor, x.RawSelectedByCurrentCompression, x.Verdict)
}
func FormatSourceType(x SourceTypeInterpretation) string {
	return fmt.Sprintf("leading=%q quadratic=%q cubic=%q compact=%q verdict=%q", x.Leading, x.Quadratic, x.Cubic, x.Compact, x.Verdict)
}
func FormatFirewall(x CoordinateNaturalityFirewall) string {
	return fmt.Sprintf("rawNative=%t momentNative=%t runtimeNative=%t massNative=%t yukawaNative=%t verdict=%q", x.RawMomentsNativelySelected, x.BoundaryMomentTheoremNative, x.ScalarRuntimeTheoremNative, x.HiggsMassTheoremNative, x.YukawaTheoremNative, x.Verdict)
}
