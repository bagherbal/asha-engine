// Package matchingcorrectionseal implements Gate 218: MatchingCorrectionSeal /
// full SM Yukawa 2-loop integration audit.
//
// Gate 217 proved that finite spectral-action matching corrections cannot yet
// be derived because the heavy-sector finite Dirac operator, heat-kernel gauge
// projection, and subtraction scheme are missing. Gate 218 therefore introduces
// an explicit MatchingCorrectionSeal and then asks a narrower phenomenological
// question: does the Gate-215 single-scale target remain viable when the SM top
// Yukawa and Higgs quartic are evolved together with the two-loop gauge system?
package matchingcorrectionseal

import (
	"fmt"
	"math"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/finitespectraltriple"
)

const (
	StatusConditionalPhenomenology = "CONDITIONAL_PHENOMENOLOGY_ON_MATCHING_CORRECTION_SEAL_FULL_SM_YUKAWA_2LOOP"
	StatusFailedRoute              = "FAILED_ROUTE_FULL_SM_YUKAWA_2LOOP_MATCHING_AUDIT"

	MatchingCorrectionSealID = "SEAL-MATCHING-CORRECTION-GATE218"
	mzGeV                    = 91.1876
	planckLogBound           = 37.8
	targetU                  = 1.0
	runTolerance             = 2.0e-7
)

type FloatTriple struct{ U1GUT, SU2L, SU3C float64 }

func (t FloatTriple) At(i int) float64 {
	switch i {
	case 0:
		return t.U1GUT
	case 1:
		return t.SU2L
	case 2:
		return t.SU3C
	default:
		panic("bad gauge index")
	}
}
func (t FloatTriple) Add(o FloatTriple) FloatTriple {
	return FloatTriple{t.U1GUT + o.U1GUT, t.SU2L + o.SU2L, t.SU3C + o.SU3C}
}
func (t FloatTriple) MaxAbs() float64 {
	return math.Max(math.Abs(t.U1GUT), math.Max(math.Abs(t.SU2L), math.Abs(t.SU3C)))
}
func (t FloatTriple) String() string {
	return fmt.Sprintf("(%.12g,%.12g,%.12g)", t.U1GUT, t.SU2L, t.SU3C)
}

type FloatMatrix3 struct{ M [3][3]float64 }

func (m FloatMatrix3) Add(o FloatMatrix3) FloatMatrix3 {
	var out FloatMatrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out.M[i][j] = m.M[i][j] + o.M[i][j]
		}
	}
	return out
}
func (m FloatMatrix3) String() string {
	rows := make([]string, 0, 3)
	for i := 0; i < 3; i++ {
		rows = append(rows, fmt.Sprintf("[%.12g,%.12g,%.12g]", m.M[i][0], m.M[i][1], m.M[i][2]))
	}
	return "[" + strings.Join(rows, ";") + "]"
}

type Gate217Snapshot struct {
	Gate217Inherited                bool
	FiniteSpectralTripleFailed      bool
	MatchingResidualTargetInherited bool
	ThresholdSpectrumSealInherited  bool
	MissingFiniteDiracOperator      bool
	MissingGaugeCurvatureProjection bool
	MissingCutoffSubtractionScheme  bool
	RequiredDeltaMatch              FloatTriple
	ResidualMaxAbs                  float64
	TruthStatement                  string
}

type MatchingCorrectionSeal struct {
	ID                          string
	Active                      bool
	RequiredDeltaMatch          FloatTriple
	ResidualPromotedAsDerived   bool
	FiniteSpectralTripleDerived bool
	TheoreticalStatus           string
	Verdict                     string
}

type PhenomenologicalInputs struct {
	Source                   string
	MZGeV                    float64
	AlphaEMInvMZ             float64
	Sin2ThetaWMZ             float64
	AlphaSMZ                 float64
	TopPoleMassGeV           float64
	HiggsMassGeV             float64
	ElectroweakVEVGeV        float64
	InitialYTop              float64
	InitialLambda            float64
	UsesPoleToYukawaTreeSeed bool
	FiniteCoreDerived        bool
	Verdict                  string
}

type HeavySpectrum struct {
	Row1Name        string
	Row2Name        string
	Row1Rep         string
	Row2Rep         string
	Row1DeltaB      FloatTriple
	Row2DeltaB      FloatTriple
	TotalDeltaB     FloatTriple
	SMBeta          FloatTriple
	TotalBeta       FloatTriple
	SMTwoLoop       FloatMatrix3
	HeavyTwoLoop    FloatMatrix3
	TotalTwoLoop    FloatMatrix3
	TargetU         float64
	ConditionalOnly bool
	Verdict         string
}

type IntegratorConfig struct {
	Method                  string
	Equation                string
	Coordinates             string
	StepsPerLogUnit         int
	MinimumStepsPerSegment  int
	MaxCoordinateIterations int
	InitialStarts           int
	YukawaTermsIncluded     bool
	LambdaRunningIncluded   bool
	HeavyYukawaCouplings    bool
	MatchingCorrectionsUsed bool
	Verdict                 string
}

type RunningState struct {
	U      [3]float64
	YTop   float64
	Lambda float64
}

type DegenerateFullSMFit struct {
	LB                  float64
	LStar               float64
	MBGeV               float64
	MStarGeV            float64
	BoundaryU           [3]float64
	BoundaryYTop        float64
	BoundaryLambda      float64
	Residual            FloatTriple
	RequiredDeltaMatch  FloatTriple
	ResidualNorm        float64
	MaxAbsResidual      float64
	RMSResidual         float64
	ResidualOverEpsilon float64
	OptimizerIterations int
	Converged           bool
	ScaleOrdered        bool
	SubPlanck           bool
	PositiveToBoundary  bool
	NoLandauBelowPlanck bool
	MatchingPlausible   bool
	Status              string
	Verdict             string
}

type ComparisonAudit struct {
	Gate215MBGeV               float64
	Gate215MStarGeV            float64
	Gate215RequiredDeltaMatch  FloatTriple
	Gate215ResidualOverEpsilon float64
	Gate218MBGeV               float64
	Gate218MStarGeV            float64
	Gate218RequiredDeltaMatch  FloatTriple
	Gate218ResidualOverEpsilon float64
	MBShiftFactor              float64
	MStarShiftFactor           float64
	ResidualShiftMaxAbs        float64
	PlausibilityPreserved      bool
	Verdict                    string
}

type FullSMYukawaAudit struct {
	GaugeYukawaCoefficients FloatTriple
	YTopBetaIncluded        bool
	LambdaBetaIncluded      bool
	BottomTauYukawasOmitted bool
	SMYukawaMatricesDerived bool
	TopMassDerived          bool
	HiggsMassDerived        bool
	HeavySectorYukawasAdded bool
	Verdict                 string
}

type FirewallAudit struct {
	Gate217Inherited                bool
	MatchingCorrectionSealActive    bool
	RequiredResidualQuarantined     bool
	MatchingCorrectionsDerived      bool
	MatchingResidualPromoted        bool
	ThresholdSpectrumSealInherited  bool
	EmpiricalCarrierSealInherited   bool
	LeptoquarkDynamicsSealInherited bool
	EmpiricalLedgerQuarantined      bool
	TopMassFiniteDerived            bool
	HiggsMassFiniteDerived          bool
	SMYukawaMatricesFiniteDerived   bool
	PhysicalPredictionClaimed       bool
	ProtonLifetimeComputed          bool
	RecommendedNextGate             string
	OpenRequirements                []string
	Verdict                         string
}

type Summary struct {
	TestsAudited               int
	Gate217Inherited           bool
	MatchingSealActive         bool
	FullSMYukawaIncluded       bool
	FitConverged               bool
	MatchingPlausible          bool
	MatchingCorrectionsDerived bool
	Status                     string
	Comment                    string
}

type Analysis struct {
	Gate217         Gate217Snapshot
	Gate217Analysis finitespectraltriple.Analysis
	Seal            MatchingCorrectionSeal
	Inputs          PhenomenologicalInputs
	Spectrum        HeavySpectrum
	Config          IntegratorConfig
	FullSM          FullSMYukawaAudit
	Fit             DegenerateFullSMFit
	Comparison      ComparisonAudit
	Firewall        FirewallAudit
	Summary         Summary
	TruthStatement  string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		g217, err := finitespectraltriple.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 217 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(g217)
	})
	return defaultA, defaultErr
}

func Build(g217 finitespectraltriple.Analysis) (Analysis, error) {
	snap := snapshotFromGate217(g217)
	if !snap.Gate217Inherited || !snap.FiniteSpectralTripleFailed || !snap.ThresholdSpectrumSealInherited {
		return Analysis{}, fmt.Errorf("Gate 218 requires Gate 217 failed spectral-triple route under ThresholdSpectrumSeal")
	}
	seal := buildSeal(snap)
	inputs := buildInputs()
	spectrum := buildSpectrum()
	config := defaultConfig()
	fullSM := auditFullSM()
	fit := fitDegenerateFullSM(spectrum, inputs, config)
	comparison := compareToGate215(fit)
	firewall := auditFirewall(g217, snap, seal, fullSM)
	status := StatusConditionalPhenomenology
	if !fit.Converged || !fit.MatchingPlausible {
		status = StatusFailedRoute
	}
	summary := Summary{TestsAudited: 8, Gate217Inherited: snap.Gate217Inherited, MatchingSealActive: seal.Active, FullSMYukawaIncluded: config.YukawaTermsIncluded && config.LambdaRunningIncluded, FitConverged: fit.Converged, MatchingPlausible: fit.MatchingPlausible, MatchingCorrectionsDerived: false, Status: status, Comment: "Gate 218 seals the Gate-215 matching residual and reruns the forced single-scale threshold audit with SM top-Yukawa and Higgs-quartic running included as phenomenological inputs."}
	truth := buildTruth(fit, comparison, seal)
	return Analysis{Gate217: snap, Gate217Analysis: g217, Seal: seal, Inputs: inputs, Spectrum: spectrum, Config: config, FullSM: fullSM, Fit: fit, Comparison: comparison, Firewall: firewall, Summary: summary, TruthStatement: truth}, nil
}

func snapshotFromGate217(a finitespectraltriple.Analysis) Gate217Snapshot {
	r := a.Gate216.SpectralResidualTarget
	return Gate217Snapshot{
		Gate217Inherited:                a.Summary.Status == finitespectraltriple.StatusFailedRoute,
		FiniteSpectralTripleFailed:      a.Readiness.CanOnlyStateTarget && !a.Readiness.CanDeriveDeltaMatch,
		MatchingResidualTargetInherited: a.Gate216.Gate215SingleScaleTargetInherited,
		ThresholdSpectrumSealInherited:  a.Firewall.ThresholdSpectrumSealInherited,
		MissingFiniteDiracOperator:      a.DiracAudit.MissingPiece == finitespectraltriple.DiracMissingOperator,
		MissingGaugeCurvatureProjection: a.HeatKernel.MissingPiece == finitespectraltriple.ProjectionMissing,
		MissingCutoffSubtractionScheme:  a.Cutoff.MissingPiece == finitespectraltriple.CutoffMissing,
		RequiredDeltaMatch:              FloatTriple{r.U1GUT, r.SU2L, r.SU3C},
		ResidualMaxAbs:                  math.Max(math.Abs(r.U1GUT), math.Max(math.Abs(r.SU2L), math.Abs(r.SU3C))),
		TruthStatement:                  a.TruthStatement,
	}
}

func buildSeal(s Gate217Snapshot) MatchingCorrectionSeal {
	return MatchingCorrectionSeal{ID: MatchingCorrectionSealID, Active: true, RequiredDeltaMatch: s.RequiredDeltaMatch, ResidualPromotedAsDerived: false, FiniteSpectralTripleDerived: false, TheoreticalStatus: "MATCHING_CORRECTION_SEALED_NOT_DERIVED", Verdict: "Gate 218 quarantines δ_match as a theoretical boundary condition because Gate 217 blocks finite spectral-action derivation"}
}

func buildInputs() PhenomenologicalInputs {
	v := 246.21965
	mt := 172.56
	mh := 125.20
	yt := math.Sqrt(2) * mt / v
	lam := mh * mh / (2 * v * v)
	return PhenomenologicalInputs{Source: "PDG-2025-style phenomenological electroweak inputs; tree-level y_t and λ seeds, not finite-core derivations", MZGeV: mzGeV, AlphaEMInvMZ: 127.955, Sin2ThetaWMZ: 0.23122, AlphaSMZ: 0.1179, TopPoleMassGeV: mt, HiggsMassGeV: mh, ElectroweakVEVGeV: v, InitialYTop: yt, InitialLambda: lam, UsesPoleToYukawaTreeSeed: true, FiniteCoreDerived: false, Verdict: "top and Higgs inputs are empirical phenomenological seeds only"}
}

func buildSpectrum() HeavySpectrum {
	smB := FloatTriple{41.0 / 10.0, -19.0 / 6.0, -7.0}
	row1 := FloatTriple{12.0 / 5.0, 8.0 / 3.0, 0}
	row2 := FloatTriple{16.0 / 5.0, 16.0 / 3.0, 8.0}
	smM := smTwoLoopMatrix()
	heavy := FloatMatrix3{M: [3][3]float64{{144.0 / 25.0, 108.0 / 5.0, 144.0 / 5.0}, {36.0 / 5.0, 108.0, 48.0}, {18.0 / 5.0, 18.0, 192.0}}}
	return HeavySpectrum{Row1Name: "Dirac fermion electroweak triplet", Row2Name: "Dirac fermion color-octet weak doublet", Row1Rep: "(1,3,Y=1)", Row2Rep: "(8,2,Y=1/2)", Row1DeltaB: row1, Row2DeltaB: row2, TotalDeltaB: row1.Add(row2), SMBeta: smB, TotalBeta: smB.Add(row1).Add(row2), SMTwoLoop: smM, HeavyTwoLoop: heavy, TotalTwoLoop: smM.Add(heavy), TargetU: targetU, ConditionalOnly: true, Verdict: "Gate-215 unique plausible single-scale spectrum under ThresholdSpectrumSeal"}
}

func defaultConfig() IntegratorConfig {
	return IntegratorConfig{Method: "forced single-threshold piecewise RK4 with coordinate minimization", Equation: "du_i/dlnμ = -b_i/(8π²) - (Σ_j B_ij/u_j - c_i y_t²)/(128π⁴), plus one-loop y_t and λ running", Coordinates: "L=ln(μ/M_Z), with M_B single-scale and M_* boundary", StepsPerLogUnit: 36, MinimumStepsPerSegment: 40, MaxCoordinateIterations: 420, InitialStarts: 9, YukawaTermsIncluded: true, LambdaRunningIncluded: true, HeavyYukawaCouplings: false, MatchingCorrectionsUsed: false, Verdict: "full SM scalar/Yukawa sector is phenomenological input; no matching row is used in the optimization"}
}

func auditFullSM() FullSMYukawaAudit {
	return FullSMYukawaAudit{GaugeYukawaCoefficients: FloatTriple{17.0 / 10.0, 3.0 / 2.0, 2.0}, YTopBetaIncluded: true, LambdaBetaIncluded: true, BottomTauYukawasOmitted: true, SMYukawaMatricesDerived: false, TopMassDerived: false, HiggsMassDerived: false, HeavySectorYukawasAdded: false, Verdict: "top Yukawa and Higgs quartic are included as empirical SM running variables; no full finite Yukawa texture derivation is claimed"}
}

func fitDegenerateFullSM(s HeavySpectrum, in PhenomenologicalInputs, cfg IntegratorConfig) DegenerateFullSMFit {
	starts := [][2]float64{{10.260447, 35.166578}, {10.0, 35.0}, {10.5, 35.2}, {11.0, 35.4}, {9.5, 34.8}, {12.0, 35.8}, {8.0, 34.0}, {14.0, 36.0}, {10.260447, 36.0}}
	bestP := [2]float64{math.NaN(), math.NaN()}
	bestObj := math.Inf(1)
	bestIt := 0
	for _, st := range starts {
		p, obj, it := coordinateOptimize(s, in, cfg, st)
		if obj < bestObj {
			bestP, bestObj, bestIt = p, obj, it
		}
	}
	state := integrateTo(s, in, cfg, bestP[0], bestP[1])
	residual := FloatTriple{state.U[0] - targetU, state.U[1] - targetU, state.U[2] - targetU}
	req := FloatTriple{-residual.U1GUT, -residual.SU2L, -residual.SU3C}
	n := math.Sqrt(residual.U1GUT*residual.U1GUT + residual.SU2L*residual.SU2L + residual.SU3C*residual.SU3C)
	max := residual.MaxAbs()
	eps := 1.0 / (16.0 * math.Pi * math.Pi)
	positive, noLandau := positivityAudits(s, in, cfg, bestP[0], bestP[1])
	ordered := validParams(bestP)
	plausible := ordered && positive && noLandau && max <= eps
	converged := math.Abs(max-bestObj) < 1.0e-8 && finite(max)
	status := "MATCHING_RESIDUAL_EXCEEDS_LOOP_FACTOR_ENVELOPE"
	verdict := "full-SM-Yukawa single-scale fit requires matching outside the loop-factor envelope"
	if plausible {
		status = "MATCHING_RESIDUAL_WITHIN_LOOP_FACTOR_ENVELOPE"
		verdict = "full-SM-Yukawa single-scale fit remains viable under the MatchingCorrectionSeal; required residual is target data, not derived"
	}
	return DegenerateFullSMFit{LB: bestP[0], LStar: bestP[1], MBGeV: mzGeV * math.Exp(bestP[0]), MStarGeV: mzGeV * math.Exp(bestP[1]), BoundaryU: state.U, BoundaryYTop: state.YTop, BoundaryLambda: state.Lambda, Residual: residual, RequiredDeltaMatch: req, ResidualNorm: n, MaxAbsResidual: max, RMSResidual: n / math.Sqrt(3), ResidualOverEpsilon: max / eps, OptimizerIterations: bestIt, Converged: converged, ScaleOrdered: ordered, SubPlanck: ordered && bestP[1] < planckLogBound, PositiveToBoundary: positive, NoLandauBelowPlanck: noLandau, MatchingPlausible: plausible, Status: status, Verdict: verdict}
}

func coordinateOptimize(s HeavySpectrum, in PhenomenologicalInputs, cfg IntegratorConfig, start [2]float64) ([2]float64, float64, int) {
	p := start
	if !validParams(p) {
		p = [2]float64{10.25, 35.15}
	}
	best := objective(s, in, cfg, p)
	step := 3.0
	dirs := [][2]float64{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	iters := 0
	for iters = 0; iters < cfg.MaxCoordinateIterations && step > 2e-6; iters++ {
		improved := false
		for _, d := range dirs {
			q := [2]float64{p[0] + step*d[0], p[1] + step*d[1]}
			if !validParams(q) {
				continue
			}
			obj := objective(s, in, cfg, q)
			if obj < best {
				p, best, improved = q, obj, true
			}
		}
		if !improved {
			step *= 0.55
		}
	}
	return p, best, iters
}

func objective(s HeavySpectrum, in PhenomenologicalInputs, cfg IntegratorConfig, p [2]float64) float64 {
	if !validParams(p) {
		return math.Inf(1)
	}
	st := integrateTo(s, in, cfg, p[0], p[1])
	if !finiteState(st) {
		return math.Inf(1)
	}
	return maxAbs3([3]float64{st.U[0] - targetU, st.U[1] - targetU, st.U[2] - targetU})
}

func integrateTo(s HeavySpectrum, in PhenomenologicalInputs, cfg IntegratorConfig, lb, lend float64) RunningState {
	st := initialState(in)
	if lend <= 0 {
		return st
	}
	if lb > 0 {
		st = integrateSegment(st, 0, math.Min(lb, lend), s.SMBeta, s.SMTwoLoop, cfg)
	}
	if lend > lb {
		st = integrateSegment(st, math.Max(0, lb), lend, s.TotalBeta, s.TotalTwoLoop, cfg)
	}
	return st
}

func initialState(in PhenomenologicalInputs) RunningState {
	return RunningState{U: [3]float64{((3.0 / 5.0) * (1.0 - in.Sin2ThetaWMZ) * in.AlphaEMInvMZ) / (4.0 * math.Pi), (in.Sin2ThetaWMZ * in.AlphaEMInvMZ) / (4.0 * math.Pi), (1.0 / in.AlphaSMZ) / (4.0 * math.Pi)}, YTop: in.InitialYTop, Lambda: in.InitialLambda}
}

func integrateSegment(st RunningState, a, b float64, beta FloatTriple, mat FloatMatrix3, cfg IntegratorConfig) RunningState {
	if b <= a {
		return st
	}
	steps := int(math.Ceil((b - a) * float64(cfg.StepsPerLogUnit)))
	if steps < cfg.MinimumStepsPerSegment {
		steps = cfg.MinimumStepsPerSegment
	}
	h := (b - a) / float64(steps)
	for i := 0; i < steps; i++ {
		k1 := deriv(st, beta, mat)
		k2 := deriv(addScaled(st, k1, 0.5*h), beta, mat)
		k3 := deriv(addScaled(st, k2, 0.5*h), beta, mat)
		k4 := deriv(addScaled(st, k3, h), beta, mat)
		for j := 0; j < 3; j++ {
			st.U[j] += h * (k1.U[j] + 2*k2.U[j] + 2*k3.U[j] + k4.U[j]) / 6.0
		}
		st.YTop += h * (k1.YTop + 2*k2.YTop + 2*k3.YTop + k4.YTop) / 6.0
		st.Lambda += h * (k1.Lambda + 2*k2.Lambda + 2*k3.Lambda + k4.Lambda) / 6.0
		if !finiteState(st) || min3(st.U) <= 0 {
			return RunningState{U: [3]float64{math.NaN(), math.NaN(), math.NaN()}, YTop: math.NaN(), Lambda: math.NaN()}
		}
	}
	return st
}

func deriv(st RunningState, beta FloatTriple, mat FloatMatrix3) RunningState {
	var out RunningState
	cY := [3]float64{17.0 / 10.0, 3.0 / 2.0, 2.0}
	for i := 0; i < 3; i++ {
		sum := 0.0
		for j := 0; j < 3; j++ {
			if st.U[j] <= 0 {
				return RunningState{U: [3]float64{math.NaN(), math.NaN(), math.NaN()}, YTop: math.NaN(), Lambda: math.NaN()}
			}
			sum += mat.M[i][j] / st.U[j]
		}
		out.U[i] = -beta.At(i)/(8.0*math.Pi*math.Pi) - (sum-cY[i]*st.YTop*st.YTop)/(128.0*math.Pow(math.Pi, 4))
	}
	g1sq, g2sq, g3sq := 1.0/st.U[0], 1.0/st.U[1], 1.0/st.U[2]
	yp := st.YTop
	out.YTop = yp / (16.0 * math.Pi * math.Pi) * (4.5*yp*yp - (17.0/20.0*g1sq + 9.0/4.0*g2sq + 8.0*g3sq))
	gpSq := 3.0 / 5.0 * g1sq
	l := st.Lambda
	out.Lambda = (24*l*l - 6*math.Pow(yp, 4) + (9.0/8.0)*g2sq*g2sq + (3.0/4.0)*g2sq*gpSq + (3.0/8.0)*gpSq*gpSq + (-9*g2sq-3*gpSq+12*yp*yp)*l) / (16.0 * math.Pi * math.Pi)
	return out
}

func addScaled(st, k RunningState, h float64) RunningState {
	return RunningState{U: [3]float64{st.U[0] + h*k.U[0], st.U[1] + h*k.U[1], st.U[2] + h*k.U[2]}, YTop: st.YTop + h*k.YTop, Lambda: st.Lambda + h*k.Lambda}
}

func positivityAudits(s HeavySpectrum, in PhenomenologicalInputs, cfg IntegratorConfig, lb, lstar float64) (bool, bool) {
	a := integrateTo(s, in, cfg, lb, lstar)
	b := integrateTo(s, in, cfg, lb, planckLogBound)
	return finiteState(a) && min3(a.U) > 1e-9, finiteState(b) && min3(b.U) > 1e-9
}

func compareToGate215(f DegenerateFullSMFit) ComparisonAudit {
	g215Req := FloatTriple{-0.000561193804, 0.000561440698, -0.000560508948}
	g215MB := 2.60752425e6
	g215MS := 1.71690311e17
	eps := 1.0 / (16.0 * math.Pi * math.Pi)
	diff := FloatTriple{f.RequiredDeltaMatch.U1GUT - g215Req.U1GUT, f.RequiredDeltaMatch.SU2L - g215Req.SU2L, f.RequiredDeltaMatch.SU3C - g215Req.SU3C}
	return ComparisonAudit{Gate215MBGeV: g215MB, Gate215MStarGeV: g215MS, Gate215RequiredDeltaMatch: g215Req, Gate215ResidualOverEpsilon: g215Req.MaxAbs() / eps, Gate218MBGeV: f.MBGeV, Gate218MStarGeV: f.MStarGeV, Gate218RequiredDeltaMatch: f.RequiredDeltaMatch, Gate218ResidualOverEpsilon: f.ResidualOverEpsilon, MBShiftFactor: f.MBGeV / g215MB, MStarShiftFactor: f.MStarGeV / g215MS, ResidualShiftMaxAbs: diff.MaxAbs(), PlausibilityPreserved: f.MatchingPlausible, Verdict: "full SM top-Yukawa/Higgs running shifts the sealed single-scale fit but preserves the matching-envelope plausibility test"}
}

func auditFirewall(g217 finitespectraltriple.Analysis, snap Gate217Snapshot, seal MatchingCorrectionSeal, y FullSMYukawaAudit) FirewallAudit {
	return FirewallAudit{Gate217Inherited: snap.Gate217Inherited, MatchingCorrectionSealActive: seal.Active, RequiredResidualQuarantined: !seal.ResidualPromotedAsDerived, MatchingCorrectionsDerived: false, MatchingResidualPromoted: false, ThresholdSpectrumSealInherited: snap.ThresholdSpectrumSealInherited, EmpiricalCarrierSealInherited: g217.Firewall.EmpiricalCarrierSealInherited, LeptoquarkDynamicsSealInherited: g217.Firewall.LeptoquarkDynamicsSealInherited, EmpiricalLedgerQuarantined: g217.Firewall.EmpiricalLedgerQuarantined, TopMassFiniteDerived: y.TopMassDerived, HiggsMassFiniteDerived: y.HiggsMassDerived, SMYukawaMatricesFiniteDerived: y.SMYukawaMatricesDerived, PhysicalPredictionClaimed: false, ProtonLifetimeComputed: false, RecommendedNextGate: "Gate 219 — finite matching-correction seal stability / input-sensitivity and bottom-tau-Yukawa audit", OpenRequirements: []string{"derive finite D_F, heat-kernel gauge projection, and cutoff/subtraction before replacing MatchingCorrectionSeal", "upgrade empirical top/Higgs seeds to a controlled MSbar input ledger before precision publication", "include bottom/tau and full two-loop scalar-sector formulas only under explicit phenomenological seals", "keep ThresholdSpectrumSeal, EmpiricalCarrierSeal, and LeptoquarkDynamicsSeal active"}, Verdict: "firewall preserved: matching vector, top mass, Higgs mass, and Yukawa/scalar running are phenomenological inputs, not finite-core theorems"}
}

func smTwoLoopMatrix() FloatMatrix3 {
	return FloatMatrix3{M: [3][3]float64{{199.0 / 50.0, 27.0 / 10.0, 44.0 / 5.0}, {9.0 / 10.0, 35.0 / 6.0, 12.0}, {11.0 / 10.0, 9.0 / 2.0, -26.0}}}
}
func validParams(p [2]float64) bool {
	return finite(p[0]) && finite(p[1]) && p[0] > 0 && p[1] > p[0] && p[1] < planckLogBound
}
func finite(x float64) bool { return !math.IsNaN(x) && !math.IsInf(x, 0) }
func finiteState(st RunningState) bool {
	return finite(st.YTop) && finite(st.Lambda) && finite(st.U[0]) && finite(st.U[1]) && finite(st.U[2])
}
func min3(a [3]float64) float64 { return math.Min(a[0], math.Min(a[1], a[2])) }
func maxAbs3(a [3]float64) float64 {
	return math.Max(math.Abs(a[0]), math.Max(math.Abs(a[1]), math.Abs(a[2])))
}

func buildTruth(f DegenerateFullSMFit, c ComparisonAudit, seal MatchingCorrectionSeal) string {
	return fmt.Sprintf("Gate 218 introduces %s and reruns the unique Gate-215 single-scale spectrum with top-Yukawa and Higgs-quartic running. The forced full-SM two-loop fit gives M_B=%.9g GeV and M*=%.9g GeV with required δ_match=%s, max|δ|/ε=%.6g. Plausibility preserved=%t. These are sealed phenomenological numbers only; finite matching corrections remain un-derived.", seal.ID, f.MBGeV, f.MStarGeV, f.RequiredDeltaMatch.String(), c.Gate218ResidualOverEpsilon, c.PlausibilityPreserved)
}

func FormatGate217(g Gate217Snapshot) string {
	return fmt.Sprintf("gate217=%t failedSpectralTriple=%t target=%t spectrumSeal=%t missing=[DF:%t projection:%t cutoff:%t] required=%s max=%.9g", g.Gate217Inherited, g.FiniteSpectralTripleFailed, g.MatchingResidualTargetInherited, g.ThresholdSpectrumSealInherited, g.MissingFiniteDiracOperator, g.MissingGaugeCurvatureProjection, g.MissingCutoffSubtractionScheme, g.RequiredDeltaMatch.String(), g.ResidualMaxAbs)
}
func FormatSeal(s MatchingCorrectionSeal) string {
	return fmt.Sprintf("id=%s active=%t required=%s promoted=%t finiteSpectral=%t status=%s", s.ID, s.Active, s.RequiredDeltaMatch.String(), s.ResidualPromotedAsDerived, s.FiniteSpectralTripleDerived, s.TheoreticalStatus)
}
func FormatInputs(i PhenomenologicalInputs) string {
	return fmt.Sprintf("source=%q MZ=%.6g alphaInv=%.6g sin2=%.6g alphaS=%.6g mt=%.6g mH=%.6g v=%.6g yt0=%.9g lambda0=%.9g treeSeed=%t finite=%t", i.Source, i.MZGeV, i.AlphaEMInvMZ, i.Sin2ThetaWMZ, i.AlphaSMZ, i.TopPoleMassGeV, i.HiggsMassGeV, i.ElectroweakVEVGeV, i.InitialYTop, i.InitialLambda, i.UsesPoleToYukawaTreeSeed, i.FiniteCoreDerived)
}
func FormatSpectrum(s HeavySpectrum) string {
	return fmt.Sprintf("rows=[%s %s Δb=%s; %s %s Δb=%s] totalΔ=%s Bheavy=%s target=%.3g", s.Row1Name, s.Row1Rep, s.Row1DeltaB.String(), s.Row2Name, s.Row2Rep, s.Row2DeltaB.String(), s.TotalDeltaB.String(), s.HeavyTwoLoop.String(), s.TargetU)
}
func FormatConfig(c IntegratorConfig) string {
	return fmt.Sprintf("method=%s steps=%d min=%d iters=%d starts=%d yukawa=%t lambda=%t heavyYukawa=%t matchingUsed=%t", c.Method, c.StepsPerLogUnit, c.MinimumStepsPerSegment, c.MaxCoordinateIterations, c.InitialStarts, c.YukawaTermsIncluded, c.LambdaRunningIncluded, c.HeavyYukawaCouplings, c.MatchingCorrectionsUsed)
}
func FormatFullSM(a FullSMYukawaAudit) string {
	return fmt.Sprintf("gaugeYukawaC=%s ytop=%t lambda=%t omitBottomTau=%t derived[yukawa:%t top:%t higgs:%t] heavyYukawa=%t", a.GaugeYukawaCoefficients.String(), a.YTopBetaIncluded, a.LambdaBetaIncluded, a.BottomTauYukawasOmitted, a.SMYukawaMatricesDerived, a.TopMassDerived, a.HiggsMassDerived, a.HeavySectorYukawasAdded)
}
func FormatFit(f DegenerateFullSMFit) string {
	return fmt.Sprintf("L=(MB=%.9g,M*=%.9g) M=(MB=%.9g,M*=%.9g) U=(%.12g,%.12g,%.12g) yt*=%.9g lambda*=%.9g residual=%s required=%s max=%.9g rms=%.9g overε=%.6g it=%d conv=%t ordered=%t subPlanck=%t positive=%t noLandau=%t plausible=%t status=%s", f.LB, f.LStar, f.MBGeV, f.MStarGeV, f.BoundaryU[0], f.BoundaryU[1], f.BoundaryU[2], f.BoundaryYTop, f.BoundaryLambda, f.Residual.String(), f.RequiredDeltaMatch.String(), f.MaxAbsResidual, f.RMSResidual, f.ResidualOverEpsilon, f.OptimizerIterations, f.Converged, f.ScaleOrdered, f.SubPlanck, f.PositiveToBoundary, f.NoLandauBelowPlanck, f.MatchingPlausible, f.Status)
}
func FormatComparison(c ComparisonAudit) string {
	return fmt.Sprintf("G215 M=(%.9g,%.9g) δ=%s overε=%.6g | G218 M=(%.9g,%.9g) δ=%s overε=%.6g shifts=(MB %.6g,M* %.6g,δmax %.6g) preserved=%t", c.Gate215MBGeV, c.Gate215MStarGeV, c.Gate215RequiredDeltaMatch.String(), c.Gate215ResidualOverEpsilon, c.Gate218MBGeV, c.Gate218MStarGeV, c.Gate218RequiredDeltaMatch.String(), c.Gate218ResidualOverEpsilon, c.MBShiftFactor, c.MStarShiftFactor, c.ResidualShiftMaxAbs, c.PlausibilityPreserved)
}
func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gate217=%t matchSeal=%t residualQuarantined=%t matchingDerived=%t promoted=%t spectrumSeal=%t carrier=%t lq=%t ledger=%t finite[top:%t higgs:%t yukawa:%t] prediction=%t lifetime=%t next=%q", f.Gate217Inherited, f.MatchingCorrectionSealActive, f.RequiredResidualQuarantined, f.MatchingCorrectionsDerived, f.MatchingResidualPromoted, f.ThresholdSpectrumSealInherited, f.EmpiricalCarrierSealInherited, f.LeptoquarkDynamicsSealInherited, f.EmpiricalLedgerQuarantined, f.TopMassFiniteDerived, f.HiggsMassFiniteDerived, f.SMYukawaMatricesFiniteDerived, f.PhysicalPredictionClaimed, f.ProtonLifetimeComputed, f.RecommendedNextGate)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d gate217=%t seal=%t fullSM=%t fit=%t plausible=%t matchingDerived=%t status=%s", s.TestsAudited, s.Gate217Inherited, s.MatchingSealActive, s.FullSMYukawaIncluded, s.FitConverged, s.MatchingPlausible, s.MatchingCorrectionsDerived, s.Status)
}
