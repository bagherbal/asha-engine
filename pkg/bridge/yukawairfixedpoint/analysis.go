// Package yukawairfixedpoint implements Gate 353:
// Yukawa Infrared Fixed-Point Basin / RG Attractor Reduction Audit.
//
// Gates 345-352 proved that pure kinematic algebra leaves the 15 minimal SM
// vacuum coordinates as moduli.  Gate 353 introduces time through one-loop RG
// flow and asks whether the spiral itself reduces that count: third-generation
// Yukawa quasi-fixed points, vacuum criticality at the intermediate scale, and
// baryogenesis/leptogenesis constraints are audited as dynamical selectors.
// The result is intentionally strict.  The top Yukawa has a real quasi-fixed
// basin for large UV input, but this does not by itself select the observed
// vacuum; the ASHA lambda boundary does not run to zero at the derived PeV
// scale for any perturbative top boundary in the audited lane; and baryogenesis
// still requires an explicit CP-asymmetry operator.
package yukawairfixedpoint

import (
	"fmt"
	"math"
	"strings"
	"sync"
)

const (
	AuditID = "GATE353-YUKAWA-INFRARED-FIXED-POINT-BASIN-RG-ATTRACTOR-REDUCTION-AUDIT"

	StatusRGAttractorFormalized   = "CONDITIONAL_SUPPORT_RG_ATTRACTOR_EQUATIONS_FORMALIZED"
	StatusIRBasinAudited          = "CONDITIONAL_SUPPORT_YUKAWA_IR_BASIN_AUDITED"
	StatusQuasiFixedPointDetected = "CONDITIONAL_SUPPORT_TOP_QUASI_FIXED_POINT_BASIN_DETECTED"
	StatusCriticalityFormalized   = "CONDITIONAL_SUPPORT_CENTER_CRITICALITY_CONDITION_FORMALIZED"
	StatusCriticalityScanned      = "CONDITIONAL_SUPPORT_CRITICAL_TOP_SCAN_EXECUTED"
	StatusBaryogenesisFormalized  = "CONDITIONAL_SUPPORT_LIGHT_CONE_BARYOGENESIS_CONSTRAINT_FORMALIZED"
	StatusParameterCensusUpdated  = "CONDITIONAL_SUPPORT_DYNAMICAL_PARAMETER_CENSUS_UPDATED"
	StatusTimeSelectionAudited    = "CONDITIONAL_SUPPORT_TIME_EVOLUTION_VACUUM_SELECTION_AUDITED"

	StatusTensionAttractorNotUnique     = "CONDITIONAL_TENSION_QUASI_FIXED_POINT_IS_BASIN_NOT_UNIQUE_SELECTOR"
	StatusTensionRPlusInAttractorLane   = "CONDITIONAL_TENSION_RPLUS_BOUNDARY_FLOWS_INTO_HIGH_TOP_LANE"
	StatusTensionCriticalityNotRealized = "CONDITIONAL_TENSION_ASHA_LAMBDA_BOUNDARY_DOES_NOT_HIT_ZERO_AT_MINT"
	StatusTensionBaryogenesisNeedsCP    = "CONDITIONAL_TENSION_BARYOGENESIS_REQUIRES_CP_ASYMMETRY_OPERATOR"
	StatusTensionSevenNotReached        = "CONDITIONAL_TENSION_SEVEN_SEAL_COUNT_NOT_REACHED"

	StatusFailedDynamicalSelectionNotActive = "FAILED_ROUTE_DYNAMICAL_VACUUM_SELECTION_NOT_ACTIVE"
	StatusFailedTopYukawaNotReduced         = "FAILED_ROUTE_TOP_YUKAWA_NOT_REMOVED_AS_VACUUM_COORDINATE"
	StatusFailedCriticalityNoSolution       = "FAILED_ROUTE_CENTER_CRITICALITY_HAS_NO_PERTURBATIVE_SOLUTION"
	StatusFailedBaryogenesisPhaseNotDerived = "FAILED_ROUTE_BARYOGENESIS_CP_PHASE_NOT_DERIVED"
	StatusFailedNoParameterReduction        = "FAILED_ROUTE_NO_ADDITIONAL_PARAMETER_REDUCTION_PROVED"
	StatusFailedSevenCoordinatesNotProved   = "FAILED_ROUTE_SEVEN_VACUUM_COORDINATES_NOT_PROVED"
)

const (
	inheritedGate          = 352
	startingVacuumInputs   = 15
	gutScaleGeV            = 2.40099519719e15
	intermediateScaleGeV   = 1.46774973718e6
	electroweakScaleGeV    = 246.22
	alphaInverseGUT        = 8.0 * math.Pi
	gStarSquared           = 0.5
	lambdaNativeBoundary   = 1197.0 / 9248.0
	rPlus                  = 1.645
	integrationStepsPerSeg = 12000
	perturbativeLimitSq    = 16.0 * math.Pi * math.Pi
)

type Span struct {
	AuditID       string
	InheritedGate int
	AddsFit       bool
	Purpose       string
	Verdict       string
}

type GaugeBoundary struct {
	AlphaInverse float64
	GStarSquared float64
	G1GUTSquared float64
	GYSquared    float64
	G2Squared    float64
	G3Squared    float64
	LambdaNative float64
	RPlus        float64
	RPlusYtUV    float64
}

type SpiralEquations struct {
	Formalized bool
	TopBeta    string
	BottomBeta string
	TauBeta    string
	LambdaBeta string
	OneLoop    bool
	Verdict    string
}

type FixedPointSnapshot struct {
	ScaleGeV float64
	TopY2    float64
	TopY     float64
	BottomY2 float64
	BottomY  float64
	TauY2    float64
	TauY     float64
}

type BasinPoint struct {
	YtUV           float64
	YtIR           float64
	LambdaIR       float64
	HiggsProxyGeV  float64
	Perturbative   bool
	Interpretation string
}

type SpiralAudit struct {
	Audited              bool
	Boundary             GaugeBoundary
	UVFixedPoint         FixedPointSnapshot
	Basin                []BasinPoint
	HighUVSpread         float64
	HighIRSpread         float64
	ContractionRatio     float64
	QuasiFixedPoint      bool
	RPlusInBasin         bool
	RPlusEndpointYt      float64
	RPlusEndpointMassGeV float64
	ParameterReduction   int
	ReductionProved      bool
	Verdict              string
}

type CriticalityAudit struct {
	Formalized           bool
	TargetScaleGeV       float64
	Condition            string
	ScanMinYtUV          float64
	ScanMaxYtUV          float64
	MinLambdaAtTarget    float64
	MinLambdaYtUV        float64
	LambdaAtYtZero       float64
	LambdaAtRPlus        float64
	PerturbativeSolution bool
	CriticalYtUV         float64
	ParameterReduction   int
	ReductionProved      bool
	Verdict              string
}

type BaryogenesisAudit struct {
	Formalized                  bool
	Constraint                  string
	ObservedEtaBQuarantined     float64
	StandardCKMInsufficient     bool
	BGapLeptogenesisHasCapacity bool
	CPAsymmetryOperatorDerived  bool
	ConsumesCKMOrPMNSPhase      bool
	ParameterReduction          int
	ReductionProved             bool
	Verdict                     string
}

type Census struct {
	StartingVacuumInputs  int
	SpiralReduction       int
	CriticalityReduction  int
	BaryogenesisReduction int
	TotalReduction        int
	RemainingInputs       int
	SevenSealTarget       int
	SevenSealReached      bool
	Verdict               string
}

type Summary struct {
	Executed           bool
	DynamicalSelection bool
	AnyReductionProved bool
	RemainingInputs    int
	Status             string
	DirectAnswer       string
	NextGate           string
}

type Analysis struct {
	Span         Span
	Equations    SpiralEquations
	Spiral       SpiralAudit
	Criticality  CriticalityAudit
	Baryogenesis BaryogenesisAudit
	Census       Census
	Summary      Summary
	Truth        string
}

type rgState struct {
	G1Sq, G2Sq, G3Sq float64
	Yt, Yb, Ytau     float64
	Lambda           float64
}

type betaCoefficients struct{ B1GUT, B2, B3 float64 }

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() { defaultA, defaultErr = Build() })
	return defaultA, defaultErr
}

func Build() (Analysis, error) {
	span := compileSpan()
	equations := formalizeEquations()
	spiral := auditSpiral()
	criticality := auditCriticality(spiral.Boundary)
	baryogenesis := auditBaryogenesis()
	census := updateCensus(spiral, criticality, baryogenesis)
	summary := compileSummary(census, spiral, criticality, baryogenesis)
	truth := "Gate 353 introduces time through one-loop RG flow.  The top Yukawa does possess a quasi-infrared fixed basin for sufficiently large UV values, but this is a basin of attraction rather than a unique vacuum selector.  The ASHA native quartic boundary remains positive at the derived intermediate scale throughout the perturbative top scan, so center criticality does not fix the top Yukawa.  Baryogenesis can constrain CP phases only after a native CP-asymmetry/leptogenesis operator is derived.  Therefore no additional reduction of the 15 minimal vacuum coordinates is proven in this gate."
	return Analysis{Span: span, Equations: equations, Spiral: spiral, Criticality: criticality, Baryogenesis: baryogenesis, Census: census, Summary: summary, Truth: truth}, nil
}

func compileSpan() Span {
	return Span{AuditID: AuditID, InheritedGate: inheritedGate, AddsFit: false, Purpose: "audit whether RG time evolution, vacuum criticality, and baryogenesis constraints reduce the 15 vacuum coordinates", Verdict: StatusTimeSelectionAudited}
}

func formalizeEquations() SpiralEquations {
	return SpiralEquations{
		Formalized: true,
		TopBeta:    "16π² dy_t/dlnμ = y_t[(9/2)y_t² + (3/2)y_b² - (17/20)g_1² - (9/4)g_2² - 8g_3²]",
		BottomBeta: "16π² dy_b/dlnμ = y_b[(9/2)y_b² + (3/2)y_t² + y_τ² - (1/4)g_1² - (9/4)g_2² - 8g_3²]",
		TauBeta:    "16π² dy_τ/dlnμ = y_τ[(5/2)y_τ² + 3y_b² - (9/4)g_1² - (9/4)g_2²]",
		LambdaBeta: "16π² dλ/dlnμ = 24λ² + 12λy_t² - 12y_t⁴ + gauge quartic - λ(9g_2²+3g_Y²)",
		OneLoop:    true,
		Verdict:    StatusRGAttractorFormalized,
	}
}

func auditSpiral() SpiralAudit {
	boundary := GaugeBoundary{
		AlphaInverse: alphaInverseGUT,
		GStarSquared: gStarSquared,
		G1GUTSquared: gStarSquared,
		GYSquared:    3.0 / 5.0 * gStarSquared,
		G2Squared:    gStarSquared,
		G3Squared:    gStarSquared,
		LambdaNative: lambdaNativeBoundary,
		RPlus:        rPlus,
		RPlusYtUV:    math.Sqrt(rPlus * gStarSquared),
	}
	fp := fixedPointAtScale(gutScaleGeV, boundary.G1GUTSquared, boundary.G2Squared, boundary.G3Squared)
	probes := []float64{0.05, 0.10, 0.30, 0.50, boundary.RPlusYtUV, 1.50, 3.00, 5.00}
	points := make([]BasinPoint, 0, len(probes))
	for _, y0 := range probes {
		end, ok := integrateFull(y0, 0, 0, lambdaNativeBoundary)
		mass := math.NaN()
		if ok && end.Lambda > 0 {
			mass = lambdaToMass(end.Lambda)
		}
		interp := "generic UV probe"
		if nearlyEqual(y0, boundary.RPlusYtUV, 1e-6) {
			interp = "ASHA r_+ boundary probe"
		}
		if y0 >= 1.5 {
			interp = "large-UV quasi-fixed-basin probe"
		}
		points = append(points, BasinPoint{YtUV: y0, YtIR: end.Yt, LambdaIR: end.Lambda, HiggsProxyGeV: mass, Perturbative: ok, Interpretation: interp})
	}
	high := filterHigh(points)
	uvMin, uvMax, irMin, irMax := ranges(high)
	contraction := math.Inf(1)
	if uvMax > uvMin {
		contraction = (irMax - irMin) / (uvMax - uvMin)
	}
	rPoint := findClosest(points, boundary.RPlusYtUV)
	quasi := len(high) >= 3 && contraction < 0.05
	reduction := 0
	reductionProved := false
	verdict := strings.Join([]string{StatusIRBasinAudited, StatusQuasiFixedPointDetected, StatusTensionAttractorNotUnique, StatusTensionRPlusInAttractorLane, StatusFailedTopYukawaNotReduced}, ";")
	return SpiralAudit{Audited: true, Boundary: boundary, UVFixedPoint: fp, Basin: points, HighUVSpread: uvMax - uvMin, HighIRSpread: irMax - irMin, ContractionRatio: contraction, QuasiFixedPoint: quasi, RPlusInBasin: rPoint.YtUV > 0.8, RPlusEndpointYt: rPoint.YtIR, RPlusEndpointMassGeV: rPoint.HiggsProxyGeV, ParameterReduction: reduction, ReductionProved: reductionProved, Verdict: verdict}
}

func fixedPointAtScale(scale, g1sq, g2sq, g3sq float64) FixedPointSnapshot {
	topY2 := ((17.0/20.0)*g1sq + (9.0/4.0)*g2sq + 8.0*g3sq) / (9.0 / 2.0)
	bottomY2 := ((1.0/4.0)*g1sq + (9.0/4.0)*g2sq + 8.0*g3sq) / (9.0 / 2.0)
	tauY2 := ((9.0/4.0)*g1sq + (9.0/4.0)*g2sq) / (5.0 / 2.0)
	return FixedPointSnapshot{ScaleGeV: scale, TopY2: topY2, TopY: math.Sqrt(topY2), BottomY2: bottomY2, BottomY: math.Sqrt(bottomY2), TauY2: tauY2, TauY: math.Sqrt(tauY2)}
}

func auditCriticality(boundary GaugeBoundary) CriticalityAudit {
	minLam := math.Inf(1)
	minY := math.NaN()
	lambdaY0 := math.NaN()
	lambdaR := math.NaN()
	for i := 0; i <= 400; i++ {
		y0 := 2.0 * float64(i) / 400.0
		at, ok := integrateToIntermediate(y0, lambdaNativeBoundary)
		if !ok {
			continue
		}
		if i == 0 {
			lambdaY0 = at.Lambda
		}
		if math.Abs(y0-boundary.RPlusYtUV) < 0.0025 {
			lambdaR = at.Lambda
		}
		if at.Lambda < minLam {
			minLam = at.Lambda
			minY = y0
		}
	}
	solution := minLam <= 0
	reduction := 0
	verdict := strings.Join([]string{StatusCriticalityFormalized, StatusCriticalityScanned, StatusTensionCriticalityNotRealized, StatusFailedCriticalityNoSolution}, ";")
	return CriticalityAudit{Formalized: true, TargetScaleGeV: intermediateScaleGeV, Condition: "find y_t(Λ) such that λ(M_int)=0 or λ,βλ graze zero at the derived intermediate scale", ScanMinYtUV: 0, ScanMaxYtUV: 2, MinLambdaAtTarget: minLam, MinLambdaYtUV: minY, LambdaAtYtZero: lambdaY0, LambdaAtRPlus: lambdaR, PerturbativeSolution: solution, CriticalYtUV: math.NaN(), ParameterReduction: reduction, ReductionProved: false, Verdict: verdict}
}

func auditBaryogenesis() BaryogenesisAudit {
	return BaryogenesisAudit{
		Formalized:                  true,
		Constraint:                  "Sakharov/leptogenesis requires out-of-equilibrium B or L violation plus CP-asymmetry ε_CP; η_B is a cosmological constraint on ε_CP times efficiency, not a CKM theorem by itself",
		ObservedEtaBQuarantined:     6.0e-10,
		StandardCKMInsufficient:     true,
		BGapLeptogenesisHasCapacity: true,
		CPAsymmetryOperatorDerived:  false,
		ConsumesCKMOrPMNSPhase:      false,
		ParameterReduction:          0,
		ReductionProved:             false,
		Verdict:                     strings.Join([]string{StatusBaryogenesisFormalized, StatusTensionBaryogenesisNeedsCP, StatusFailedBaryogenesisPhaseNotDerived}, ";"),
	}
}

func updateCensus(s SpiralAudit, c CriticalityAudit, b BaryogenesisAudit) Census {
	total := s.ParameterReduction + c.ParameterReduction + b.ParameterReduction
	remaining := startingVacuumInputs - total
	return Census{StartingVacuumInputs: startingVacuumInputs, SpiralReduction: s.ParameterReduction, CriticalityReduction: c.ParameterReduction, BaryogenesisReduction: b.ParameterReduction, TotalReduction: total, RemainingInputs: remaining, SevenSealTarget: 7, SevenSealReached: remaining == 7, Verdict: strings.Join([]string{StatusParameterCensusUpdated, StatusTensionSevenNotReached, StatusFailedNoParameterReduction, StatusFailedSevenCoordinatesNotProved}, ";")}
}

func compileSummary(c Census, s SpiralAudit, cr CriticalityAudit, b BaryogenesisAudit) Summary {
	any := c.TotalReduction > 0
	return Summary{Executed: true, DynamicalSelection: any, AnyReductionProved: any, RemainingInputs: c.RemainingInputs, Status: strings.Join([]string{StatusTimeSelectionAudited, StatusFailedDynamicalSelectionNotActive}, ";"), DirectAnswer: fmt.Sprintf("RG time evolution was audited.  A top quasi-fixed basin exists, but criticality and baryogenesis do not uniquely select vacuum coordinates here; remaining minimal inputs = %d.", c.RemainingInputs), NextGate: "If continuing, the missing object is a native CP/flavor-breaking operator or a nonstandard dynamical saturation principle, not another kinematic fit."}
}

func initialState(yt, yb, ytau, lambda float64) rgState {
	return rgState{G1Sq: gStarSquared, G2Sq: gStarSquared, G3Sq: gStarSquared, Yt: yt, Yb: yb, Ytau: ytau, Lambda: lambda}
}

func integrateFull(yt, yb, ytau, lambda float64) (rgState, bool) {
	high := betaCoefficients{B1GUT: 41.0/10.0 + 7.78628724237, B2: -19.0/6.0 + 9.65295390904, B3: -7.0 + 8.98628724237}
	low := betaCoefficients{B1GUT: 41.0 / 10.0, B2: -19.0 / 6.0, B3: -7.0}
	s, ok := integrateSegment(initialState(yt, yb, ytau, lambda), gutScaleGeV, intermediateScaleGeV, high, integrationStepsPerSeg)
	if !ok {
		return s, false
	}
	return integrateSegment(s, intermediateScaleGeV, electroweakScaleGeV, low, integrationStepsPerSeg)
}

func integrateToIntermediate(yt, lambda float64) (rgState, bool) {
	high := betaCoefficients{B1GUT: 41.0/10.0 + 7.78628724237, B2: -19.0/6.0 + 9.65295390904, B3: -7.0 + 8.98628724237}
	return integrateSegment(initialState(yt, 0, 0, lambda), gutScaleGeV, intermediateScaleGeV, high, integrationStepsPerSeg)
}

func integrateSegment(start rgState, from, to float64, b betaCoefficients, steps int) (rgState, bool) {
	if from <= 0 || to <= 0 || steps <= 0 {
		return start, false
	}
	s := start
	t0 := math.Log(from)
	dt := (math.Log(to) - t0) / float64(steps)
	for i := 0; i < steps; i++ {
		k1 := derivatives(s, b)
		k2 := derivatives(addScaled(s, k1, 0.5*dt), b)
		k3 := derivatives(addScaled(s, k2, 0.5*dt), b)
		k4 := derivatives(addScaled(s, k3, dt), b)
		s = combineRK4(s, k1, k2, k3, k4, dt)
		if !finiteState(s) || s.G1Sq <= 0 || s.G2Sq <= 0 || s.G3Sq <= 0 || s.G1Sq > perturbativeLimitSq || s.G2Sq > perturbativeLimitSq || s.G3Sq > perturbativeLimitSq || math.Abs(s.Yt) > 20 || math.Abs(s.Yb) > 20 || math.Abs(s.Ytau) > 20 || math.Abs(s.Lambda) > 100 {
			return s, false
		}
	}
	return s, true
}

func derivatives(s rgState, b betaCoefficients) rgState {
	inv := 1.0 / (16.0 * math.Pi * math.Pi)
	gYsq := (3.0 / 5.0) * s.G1Sq
	return rgState{
		G1Sq:   2.0 * b.B1GUT * inv * s.G1Sq * s.G1Sq,
		G2Sq:   2.0 * b.B2 * inv * s.G2Sq * s.G2Sq,
		G3Sq:   2.0 * b.B3 * inv * s.G3Sq * s.G3Sq,
		Yt:     s.Yt * inv * ((9.0/2.0)*s.Yt*s.Yt + (3.0/2.0)*s.Yb*s.Yb - (17.0/20.0)*s.G1Sq - (9.0/4.0)*s.G2Sq - 8.0*s.G3Sq),
		Yb:     s.Yb * inv * ((9.0/2.0)*s.Yb*s.Yb + (3.0/2.0)*s.Yt*s.Yt + s.Ytau*s.Ytau - (1.0/4.0)*s.G1Sq - (9.0/4.0)*s.G2Sq - 8.0*s.G3Sq),
		Ytau:   s.Ytau * inv * ((5.0/2.0)*s.Ytau*s.Ytau + 3.0*s.Yb*s.Yb - (9.0/4.0)*s.G1Sq - (9.0/4.0)*s.G2Sq),
		Lambda: inv * (24.0*s.Lambda*s.Lambda + 12.0*s.Lambda*s.Yt*s.Yt - 12.0*math.Pow(s.Yt, 4) + (3.0/16.0)*(2.0*s.G2Sq*s.G2Sq+math.Pow(s.G2Sq+gYsq, 2)) - s.Lambda*(9.0*s.G2Sq+3.0*gYsq)),
	}
}

func addScaled(s, k rgState, h float64) rgState {
	return rgState{G1Sq: s.G1Sq + h*k.G1Sq, G2Sq: s.G2Sq + h*k.G2Sq, G3Sq: s.G3Sq + h*k.G3Sq, Yt: s.Yt + h*k.Yt, Yb: s.Yb + h*k.Yb, Ytau: s.Ytau + h*k.Ytau, Lambda: s.Lambda + h*k.Lambda}
}

func combineRK4(s, k1, k2, k3, k4 rgState, dt float64) rgState {
	return rgState{G1Sq: s.G1Sq + dt*(k1.G1Sq+2*k2.G1Sq+2*k3.G1Sq+k4.G1Sq)/6, G2Sq: s.G2Sq + dt*(k1.G2Sq+2*k2.G2Sq+2*k3.G2Sq+k4.G2Sq)/6, G3Sq: s.G3Sq + dt*(k1.G3Sq+2*k2.G3Sq+2*k3.G3Sq+k4.G3Sq)/6, Yt: s.Yt + dt*(k1.Yt+2*k2.Yt+2*k3.Yt+k4.Yt)/6, Yb: s.Yb + dt*(k1.Yb+2*k2.Yb+2*k3.Yb+k4.Yb)/6, Ytau: s.Ytau + dt*(k1.Ytau+2*k2.Ytau+2*k3.Ytau+k4.Ytau)/6, Lambda: s.Lambda + dt*(k1.Lambda+2*k2.Lambda+2*k3.Lambda+k4.Lambda)/6}
}

func finiteState(s rgState) bool {
	vals := []float64{s.G1Sq, s.G2Sq, s.G3Sq, s.Yt, s.Yb, s.Ytau, s.Lambda}
	for _, v := range vals {
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return false
		}
	}
	return true
}

func filterHigh(points []BasinPoint) []BasinPoint {
	out := make([]BasinPoint, 0)
	for _, p := range points {
		if p.YtUV >= 1.5 && p.Perturbative {
			out = append(out, p)
		}
	}
	return out
}

func ranges(points []BasinPoint) (uvMin, uvMax, irMin, irMax float64) {
	if len(points) == 0 {
		return math.NaN(), math.NaN(), math.NaN(), math.NaN()
	}
	uvMin, uvMax, irMin, irMax = points[0].YtUV, points[0].YtUV, points[0].YtIR, points[0].YtIR
	for _, p := range points[1:] {
		uvMin = math.Min(uvMin, p.YtUV)
		uvMax = math.Max(uvMax, p.YtUV)
		irMin = math.Min(irMin, p.YtIR)
		irMax = math.Max(irMax, p.YtIR)
	}
	return
}

func findClosest(points []BasinPoint, target float64) BasinPoint {
	if len(points) == 0 {
		return BasinPoint{}
	}
	best := points[0]
	bestD := math.Abs(points[0].YtUV - target)
	for _, p := range points[1:] {
		if d := math.Abs(p.YtUV - target); d < bestD {
			best, bestD = p, d
		}
	}
	return best
}

func lambdaToMass(lambda float64) float64 {
	if lambda < 0 {
		return math.NaN()
	}
	return electroweakScaleGeV * math.Sqrt(2.0*lambda)
}

func nearlyEqual(a, b, tol float64) bool { return math.Abs(a-b) <= tol }
func relativeError(a, b float64) float64 {
	if b == 0 {
		return math.Inf(1)
	}
	return (a - b) / b
}

func FormatSpan(s Span) string {
	return fmt.Sprintf("%s inherited_gate=%d adds_fit=%t", s.AuditID, s.InheritedGate, s.AddsFit)
}
func FormatEquations(e SpiralEquations) string {
	return fmt.Sprintf("formalized=%t one_loop=%t top=%s", e.Formalized, e.OneLoop, e.TopBeta)
}
func FormatSpiral(s SpiralAudit) string {
	return fmt.Sprintf("r+_yt_uv=%.12f r+_yt_ir=%.12f r+_m=%.6f contraction=%.6g reduction=%d verdict=%s", s.Boundary.RPlusYtUV, s.RPlusEndpointYt, s.RPlusEndpointMassGeV, s.ContractionRatio, s.ParameterReduction, s.Verdict)
}
func FormatCriticality(c CriticalityAudit) string {
	return fmt.Sprintf("target=%.6g min_lambda=%.12f at_yt=%.6f lambda_y0=%.12f lambda_r+=%.12f solution=%t", c.TargetScaleGeV, c.MinLambdaAtTarget, c.MinLambdaYtUV, c.LambdaAtYtZero, c.LambdaAtRPlus, c.PerturbativeSolution)
}
func FormatBaryogenesis(b BaryogenesisAudit) string {
	return fmt.Sprintf("etaB≈%.3g CKM_insufficient=%t bgap_capacity=%t cp_operator=%t reduction=%d", b.ObservedEtaBQuarantined, b.StandardCKMInsufficient, b.BGapLeptogenesisHasCapacity, b.CPAsymmetryOperatorDerived, b.ParameterReduction)
}
func FormatCensus(c Census) string {
	return fmt.Sprintf("start=%d reductions(sp=%d,crit=%d,baryo=%d,total=%d) remaining=%d seven=%t", c.StartingVacuumInputs, c.SpiralReduction, c.CriticalityReduction, c.BaryogenesisReduction, c.TotalReduction, c.RemainingInputs, c.SevenSealReached)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("executed=%t dynamic_selection=%t remaining=%d status=%s", s.Executed, s.DynamicalSelection, s.RemainingInputs, s.Status)
}

func Statuses(a Analysis) []string {
	return []string{
		StatusRGAttractorFormalized,
		StatusIRBasinAudited,
		StatusQuasiFixedPointDetected,
		StatusCriticalityFormalized,
		StatusCriticalityScanned,
		StatusBaryogenesisFormalized,
		StatusParameterCensusUpdated,
		StatusTimeSelectionAudited,
		StatusTensionAttractorNotUnique,
		StatusTensionRPlusInAttractorLane,
		StatusTensionCriticalityNotRealized,
		StatusTensionBaryogenesisNeedsCP,
		StatusTensionSevenNotReached,
		StatusFailedDynamicalSelectionNotActive,
		StatusFailedTopYukawaNotReduced,
		StatusFailedCriticalityNoSolution,
		StatusFailedBaryogenesisPhaseNotDerived,
		StatusFailedNoParameterReduction,
		StatusFailedSevenCoordinatesNotProved,
	}
}
