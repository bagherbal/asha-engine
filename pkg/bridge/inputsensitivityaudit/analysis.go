// Package inputsensitivityaudit implements Gate 219: input-sensitivity and
// bottom-tau-Yukawa completeness audit.
//
// Gate 218 introduced the MatchingCorrectionSeal and showed that the sealed
// single-scale heavy spectrum remains viable after adding top-Yukawa and Higgs
// quartic running. Gate 219 adds bottom/tau Yukawa completeness and propagates
// low-energy empirical input uncertainties through the same forced single-scale
// two-loop audit. The outputs are conditional phenomenology only.
package inputsensitivityaudit

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/matchingcorrectionseal"
)

const (
	StatusConditionalPhenomenology = "CONDITIONAL_PHENOMENOLOGY_INPUT_SENSITIVITY_BOTTOM_TAU_COMPLETE_2LOOP"
	StatusFailedRoute              = "FAILED_ROUTE_INPUT_SENSITIVITY_BREAKS_MATCHING_ENVELOPE"

	InputSensitivitySealID = "SEAL-INPUT-SENSITIVITY-GATE219"
	mzGeV                  = 91.1876
	planckLogBound         = 37.8
	targetU                = 1.0
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

type Gate218Snapshot struct {
	Gate218Inherited             bool
	MatchingCorrectionSealActive bool
	ThresholdSpectrumSealActive  bool
	MatchingPlausible            bool
	BottomTauOmitted             bool
	MBGeV                        float64
	MStarGeV                     float64
	RequiredDeltaMatch           FloatTriple
	ResidualOverEpsilon          float64
	TruthStatement               string
}

type EmpiricalInput struct {
	Name              string
	Symbol            string
	Central           float64
	Sigma             float64
	Unit              string
	FiniteCoreDerived bool
	ScanEnabled       bool
	Verdict           string
}

type PhenomenologicalInputs struct {
	Source             string
	MZGeV              float64
	AlphaEMInvMZ       float64
	Sin2ThetaWMZ       float64
	AlphaSMZ           float64
	TopPoleMassGeV     float64
	BottomMassGeV      float64
	TauMassGeV         float64
	HiggsMassGeV       float64
	ElectroweakVEVGeV  float64
	InitialYTop        float64
	InitialYBottom     float64
	InitialYTau        float64
	InitialLambda      float64
	UsesTreeLevelSeeds bool
	FiniteCoreDerived  bool
	Verdict            string
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
	YukawaTermsIncluded     bool
	BottomTauIncluded       bool
	LambdaRunningIncluded   bool
	MatchingCorrectionsUsed bool
	Verdict                 string
}

type RunningState struct {
	U       [3]float64
	YTop    float64
	YBottom float64
	YTau    float64
	Lambda  float64
}

type DegenerateFit struct {
	CaseName            string
	LB                  float64
	LStar               float64
	MBGeV               float64
	MStarGeV            float64
	BoundaryU           [3]float64
	BoundaryYTop        float64
	BoundaryYBottom     float64
	BoundaryYTau        float64
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

type ScanCase struct {
	Name             string
	Perturbation     string
	Inputs           PhenomenologicalInputs
	Fit              DegenerateFit
	DeltaLB          float64
	DeltaLStar       float64
	DeltaMBGeV       float64
	DeltaMStarGeV    float64
	DeltaMaxResidual float64
	Verdict          string
}

type SensitivitySummary struct {
	CasesAudited             int
	ConvergedCases           int
	PlausibleCases           int
	BrokenEnvelopeCases      int
	CentralMBGeV             float64
	CentralMStarGeV          float64
	MBMinGeV                 float64
	MBMaxGeV                 float64
	MStarMinGeV              float64
	MStarMaxGeV              float64
	MBMinusGeV               float64
	MBPlusGeV                float64
	MStarMinusGeV            float64
	MStarPlusGeV             float64
	ResidualOverEpsilonMin   float64
	ResidualOverEpsilonMax   float64
	WorstCaseName            string
	WorstResidualOverEpsilon float64
	DominantScaleDriver      string
	DominantResidualDriver   string
	Verdict                  string
}

type CompletenessAudit struct {
	TopYukawaIncluded             bool
	BottomYukawaIncluded          bool
	TauYukawaIncluded             bool
	HiggsQuarticIncluded          bool
	FullYukawaMatricesDerived     bool
	OtherFermionYukawasIgnored    bool
	HeavyYukawaCouplingsAdded     bool
	GaugeYukawaCoefficientsTop    FloatTriple
	GaugeYukawaCoefficientsBottom FloatTriple
	GaugeYukawaCoefficientsTau    FloatTriple
	Verdict                       string
}

type MatchingSealAudit struct {
	MatchingCorrectionSealInherited bool
	RequiredResidualQuarantined     bool
	MatchingCorrectionsDerived      bool
	MatchingResidualPromoted        bool
	EpsilonU                        float64
	Verdict                         string
}

type FirewallAudit struct {
	Gate218Inherited                bool
	MatchingCorrectionSealActive    bool
	ThresholdSpectrumSealInherited  bool
	EmpiricalCarrierSealInherited   bool
	LeptoquarkDynamicsSealInherited bool
	EmpiricalLedgerQuarantined      bool
	InputUncertaintiesFiniteDerived bool
	InputsTunedToForceZeroResidual  bool
	MatchingCorrectionsDerived      bool
	MatchingResidualPromoted        bool
	PhysicalPredictionClaimed       bool
	ProtonLifetimeComputed          bool
	RecommendedNextGate             string
	OpenRequirements                []string
	Verdict                         string
}

type Summary struct {
	TestsAudited              int
	Gate218Inherited          bool
	BottomTauComplete         bool
	CasesAudited              int
	PlausibleCases            int
	EnvelopePreservedAt1Sigma bool
	Status                    string
	Comment                   string
}

type Analysis struct {
	Gate218         Gate218Snapshot
	Gate218Analysis matchingcorrectionseal.Analysis
	Inputs          []EmpiricalInput
	CentralInputs   PhenomenologicalInputs
	Spectrum        HeavySpectrum
	Config          IntegratorConfig
	Completeness    CompletenessAudit
	Seal            MatchingSealAudit
	CentralFit      DegenerateFit
	ScanCases       []ScanCase
	Sensitivity     SensitivitySummary
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
		g218, err := matchingcorrectionseal.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 218 input: %w", err)
			return
		}
		defaultA, defaultErr = Build(g218)
	})
	return defaultA, defaultErr
}

func Build(g218 matchingcorrectionseal.Analysis) (Analysis, error) {
	snap := snapshotFromGate218(g218)
	if !snap.Gate218Inherited || !snap.MatchingCorrectionSealActive || !snap.MatchingPlausible {
		return Analysis{}, fmt.Errorf("Gate 219 requires the Gate 218 matching-sealed plausible full-SM-Yukawa audit")
	}
	inputs := empiricalInputs()
	central := centralPhenomenologicalInputs(inputs)
	spectrum := buildSpectrum()
	cfg := defaultConfig()
	completeness := auditCompleteness()
	seal := auditSeal()
	centralFit := fitDegenerate(spectrum, central, cfg, "central")
	cases := runSensitivityCases(spectrum, cfg, central, inputs, centralFit)
	sensitivity := summarizeSensitivity(cases, centralFit)
	firewall := auditFirewall(g218, snap)
	status := StatusConditionalPhenomenology
	if sensitivity.BrokenEnvelopeCases > 0 || !centralFit.MatchingPlausible {
		status = StatusFailedRoute
	}
	summary := Summary{TestsAudited: 9, Gate218Inherited: snap.Gate218Inherited, BottomTauComplete: completeness.BottomYukawaIncluded && completeness.TauYukawaIncluded, CasesAudited: sensitivity.CasesAudited, PlausibleCases: sensitivity.PlausibleCases, EnvelopePreservedAt1Sigma: sensitivity.BrokenEnvelopeCases == 0, Status: status, Comment: "Gate 219 adds bottom/tau Yukawa completeness and propagates bounded 1σ empirical input uncertainty through the sealed single-scale two-loop audit."}
	truth := buildTruth(centralFit, sensitivity)
	return Analysis{Gate218: snap, Gate218Analysis: g218, Inputs: inputs, CentralInputs: central, Spectrum: spectrum, Config: cfg, Completeness: completeness, Seal: seal, CentralFit: centralFit, ScanCases: cases, Sensitivity: sensitivity, Firewall: firewall, Summary: summary, TruthStatement: truth}, nil
}

func snapshotFromGate218(a matchingcorrectionseal.Analysis) Gate218Snapshot {
	return Gate218Snapshot{
		Gate218Inherited:             a.Summary.Status == matchingcorrectionseal.StatusConditionalPhenomenology,
		MatchingCorrectionSealActive: a.Seal.Active,
		ThresholdSpectrumSealActive:  a.Firewall.ThresholdSpectrumSealInherited,
		MatchingPlausible:            a.Fit.MatchingPlausible,
		BottomTauOmitted:             a.FullSM.BottomTauYukawasOmitted,
		MBGeV:                        a.Fit.MBGeV,
		MStarGeV:                     a.Fit.MStarGeV,
		RequiredDeltaMatch:           FloatTriple{a.Fit.RequiredDeltaMatch.U1GUT, a.Fit.RequiredDeltaMatch.SU2L, a.Fit.RequiredDeltaMatch.SU3C},
		ResidualOverEpsilon:          a.Fit.ResidualOverEpsilon,
		TruthStatement:               a.TruthStatement,
	}
}

func empiricalInputs() []EmpiricalInput {
	return []EmpiricalInput{
		{Name: "strong coupling at Z", Symbol: "alpha_s(M_Z)", Central: 0.1179, Sigma: 0.0009, Unit: "dimensionless", ScanEnabled: true, Verdict: "primary gauge-input uncertainty; scanned at ±1σ"},
		{Name: "top pole mass", Symbol: "m_t", Central: 172.56, Sigma: 0.70, Unit: "GeV", ScanEnabled: true, Verdict: "primary Yukawa-input uncertainty; scanned at conservative ±1σ"},
		{Name: "Higgs mass", Symbol: "m_H", Central: 125.20, Sigma: 0.11, Unit: "GeV", ScanEnabled: true, Verdict: "quartic seed uncertainty; scanned at ±1σ"},
		{Name: "bottom mass", Symbol: "m_b", Central: 4.18, Sigma: 0.03, Unit: "GeV", ScanEnabled: true, Verdict: "bottom-Yukawa completeness input; scanned at ±1σ"},
		{Name: "tau mass", Symbol: "m_tau", Central: 1.77686, Sigma: 0.00012, Unit: "GeV", ScanEnabled: true, Verdict: "tau-Yukawa completeness input; scanned at ±1σ"},
		{Name: "inverse electromagnetic coupling at Z", Symbol: "alpha_em^-1(M_Z)", Central: 127.955, Sigma: 0, Unit: "dimensionless", ScanEnabled: false, Verdict: "held fixed to isolate the requested dominant uncertainties"},
		{Name: "weak mixing at Z", Symbol: "sin^2(theta_W)", Central: 0.23122, Sigma: 0, Unit: "dimensionless", ScanEnabled: false, Verdict: "held fixed to isolate the requested dominant uncertainties"},
	}
}

func centralPhenomenologicalInputs(xs []EmpiricalInput) PhenomenologicalInputs {
	vals := map[string]float64{}
	for _, x := range xs {
		vals[x.Symbol] = x.Central
	}
	return buildInputs(vals["alpha_s(M_Z)"], vals["m_t"], vals["m_b"], vals["m_tau"], vals["m_H"], vals["alpha_em^-1(M_Z)"], vals["sin^2(theta_W)"])
}

func buildInputs(alphaS, mt, mb, mtau, mh, alphaInv, sin2 float64) PhenomenologicalInputs {
	v := 246.21965
	return PhenomenologicalInputs{
		Source: "PDG-style empirical input ledger; uncertainties are scan data, not finite-core derivations",
		MZGeV:  mzGeV, AlphaEMInvMZ: alphaInv, Sin2ThetaWMZ: sin2, AlphaSMZ: alphaS,
		TopPoleMassGeV: mt, BottomMassGeV: mb, TauMassGeV: mtau, HiggsMassGeV: mh, ElectroweakVEVGeV: v,
		InitialYTop: math.Sqrt(2) * mt / v, InitialYBottom: math.Sqrt(2) * mb / v, InitialYTau: math.Sqrt(2) * mtau / v, InitialLambda: mh * mh / (2 * v * v),
		UsesTreeLevelSeeds: true, FiniteCoreDerived: false,
		Verdict: "tree-level Yukawa/quartic seeds are empirical phenomenology; they are not ASHA finite-core predictions",
	}
}

func buildSpectrum() HeavySpectrum {
	smB := FloatTriple{41.0 / 10.0, -19.0 / 6.0, -7.0}
	row1 := FloatTriple{12.0 / 5.0, 8.0 / 3.0, 0}
	row2 := FloatTriple{16.0 / 5.0, 16.0 / 3.0, 8.0}
	smM := smTwoLoopMatrix()
	heavy := FloatMatrix3{M: [3][3]float64{{144.0 / 25.0, 108.0 / 5.0, 144.0 / 5.0}, {36.0 / 5.0, 108, 48}, {18.0 / 5.0, 18, 192}}}
	return HeavySpectrum{Row1Name: "Dirac fermion electroweak triplet", Row2Name: "Dirac fermion color-octet weak doublet", Row1Rep: "(1,3,Y=1)", Row2Rep: "(8,2,Y=1/2)", Row1DeltaB: row1, Row2DeltaB: row2, TotalDeltaB: row1.Add(row2), SMBeta: smB, TotalBeta: smB.Add(row1).Add(row2), SMTwoLoop: smM, HeavyTwoLoop: heavy, TotalTwoLoop: smM.Add(heavy), TargetU: targetU, ConditionalOnly: true, Verdict: "Gate-215 unique plausible single-scale spectrum remains sealed by ThresholdSpectrumSeal"}
}

func smTwoLoopMatrix() FloatMatrix3 {
	return FloatMatrix3{M: [3][3]float64{{199.0 / 50.0, 27.0 / 10.0, 44.0 / 5.0}, {9.0 / 10.0, 35.0 / 6.0, 12.0}, {11.0 / 10.0, 9.0 / 2.0, -26.0}}}
}

func defaultConfig() IntegratorConfig {
	return IntegratorConfig{Method: "forced single-threshold RK4 plus coordinate minimization and ±1σ bounding scan", Equation: "du_i/dlnμ = -b_i/(8π²) - (Σ_j B_ij/u_j - c_t y_t² - c_b y_b² - c_tau y_tau²)/(128π⁴); one-loop y_t,y_b,y_tau,λ", Coordinates: "L=ln(μ/M_Z), M_B single-scale, M_* boundary", StepsPerLogUnit: 36, MinimumStepsPerSegment: 40, MaxCoordinateIterations: 420, YukawaTermsIncluded: true, BottomTauIncluded: true, LambdaRunningIncluded: true, MatchingCorrectionsUsed: false, Verdict: "bottom/tau completeness and input sensitivity are phenomenological upgrades, not finite derivations"}
}

func auditCompleteness() CompletenessAudit {
	return CompletenessAudit{TopYukawaIncluded: true, BottomYukawaIncluded: true, TauYukawaIncluded: true, HiggsQuarticIncluded: true, FullYukawaMatricesDerived: false, OtherFermionYukawasIgnored: true, HeavyYukawaCouplingsAdded: false, GaugeYukawaCoefficientsTop: FloatTriple{17.0 / 10.0, 3.0 / 2.0, 2}, GaugeYukawaCoefficientsBottom: FloatTriple{1.0 / 2.0, 3.0 / 2.0, 2}, GaugeYukawaCoefficientsTau: FloatTriple{3.0 / 2.0, 1.0 / 2.0, 0}, Verdict: "third-family top/bottom/tau Yukawa completeness is included; lighter Yukawas remain negligible phenomenological omissions"}
}

func auditSeal() MatchingSealAudit {
	return MatchingSealAudit{MatchingCorrectionSealInherited: true, RequiredResidualQuarantined: true, MatchingCorrectionsDerived: false, MatchingResidualPromoted: false, EpsilonU: epsilonU(), Verdict: "δ_match remains sealed target data; the scan cannot tune it to zero"}
}

func fitDegenerate(s HeavySpectrum, in PhenomenologicalInputs, cfg IntegratorConfig, name string) DegenerateFit {
	starts := [][2]float64{{10.242, 35.17}, {10.25, 35.17}, {10.0, 35.0}, {10.5, 35.2}}
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
	positive, noLandau := positivityAudits(s, in, cfg, bestP[0], bestP[1])
	ordered := validParams(bestP)
	plausible := ordered && positive && noLandau && max <= epsilonU()
	converged := finite(max) && bestObj < epsilonU()
	status := "MATCHING_RESIDUAL_EXCEEDS_LOOP_FACTOR_ENVELOPE"
	verdict := "required matching residual exceeds the phenomenological loop-factor envelope"
	if plausible {
		status = "MATCHING_RESIDUAL_WITHIN_LOOP_FACTOR_ENVELOPE"
		verdict = "required matching residual remains inside the loop-factor envelope; this is conditional phenomenology only"
	}
	return DegenerateFit{CaseName: name, LB: bestP[0], LStar: bestP[1], MBGeV: mzGeV * math.Exp(bestP[0]), MStarGeV: mzGeV * math.Exp(bestP[1]), BoundaryU: state.U, BoundaryYTop: state.YTop, BoundaryYBottom: state.YBottom, BoundaryYTau: state.YTau, BoundaryLambda: state.Lambda, Residual: residual, RequiredDeltaMatch: req, ResidualNorm: n, MaxAbsResidual: max, RMSResidual: n / math.Sqrt(3), ResidualOverEpsilon: max / epsilonU(), OptimizerIterations: bestIt, Converged: converged, ScaleOrdered: ordered, SubPlanck: ordered && bestP[1] < planckLogBound, PositiveToBoundary: positive, NoLandauBelowPlanck: noLandau, MatchingPlausible: plausible, Status: status, Verdict: verdict}
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
	for iters = 0; iters < cfg.MaxCoordinateIterations && step > 1e-5; iters++ {
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
	return RunningState{U: [3]float64{((3.0 / 5.0) * (1.0 - in.Sin2ThetaWMZ) * in.AlphaEMInvMZ) / (4.0 * math.Pi), (in.Sin2ThetaWMZ * in.AlphaEMInvMZ) / (4.0 * math.Pi), (1.0 / in.AlphaSMZ) / (4.0 * math.Pi)}, YTop: in.InitialYTop, YBottom: in.InitialYBottom, YTau: in.InitialYTau, Lambda: in.InitialLambda}
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
		st.YBottom += h * (k1.YBottom + 2*k2.YBottom + 2*k3.YBottom + k4.YBottom) / 6.0
		st.YTau += h * (k1.YTau + 2*k2.YTau + 2*k3.YTau + k4.YTau) / 6.0
		st.Lambda += h * (k1.Lambda + 2*k2.Lambda + 2*k3.Lambda + k4.Lambda) / 6.0
		if !finiteState(st) || min3(st.U) <= 0 {
			return RunningState{U: [3]float64{math.NaN(), math.NaN(), math.NaN()}, YTop: math.NaN(), YBottom: math.NaN(), YTau: math.NaN(), Lambda: math.NaN()}
		}
	}
	return st
}

func deriv(st RunningState, beta FloatTriple, mat FloatMatrix3) RunningState {
	var out RunningState
	ct := [3]float64{17.0 / 10.0, 3.0 / 2.0, 2}
	cb := [3]float64{1.0 / 2.0, 3.0 / 2.0, 2}
	ce := [3]float64{3.0 / 2.0, 1.0 / 2.0, 0}
	for i := 0; i < 3; i++ {
		sum := 0.0
		for j := 0; j < 3; j++ {
			if st.U[j] <= 0 {
				return RunningState{U: [3]float64{math.NaN(), math.NaN(), math.NaN()}, YTop: math.NaN(), YBottom: math.NaN(), YTau: math.NaN(), Lambda: math.NaN()}
			}
			sum += mat.M[i][j] / st.U[j]
		}
		yuk := ct[i]*st.YTop*st.YTop + cb[i]*st.YBottom*st.YBottom + ce[i]*st.YTau*st.YTau
		out.U[i] = -beta.At(i)/(8.0*math.Pi*math.Pi) - (sum-yuk)/(128.0*math.Pow(math.Pi, 4))
	}
	g1sq, g2sq, g3sq := 1.0/st.U[0], 1.0/st.U[1], 1.0/st.U[2]
	yt, yb, ye := st.YTop, st.YBottom, st.YTau
	trace := 3*yt*yt + 3*yb*yb + ye*ye
	out.YTop = yt / (16.0 * math.Pi * math.Pi) * (1.5*(yt*yt-yb*yb) + trace - (17.0/20.0*g1sq + 9.0/4.0*g2sq + 8.0*g3sq))
	out.YBottom = yb / (16.0 * math.Pi * math.Pi) * (1.5*(yb*yb-yt*yt) + trace - (1.0/4.0*g1sq + 9.0/4.0*g2sq + 8.0*g3sq))
	out.YTau = ye / (16.0 * math.Pi * math.Pi) * (1.5*ye*ye + trace - (9.0/4.0*g1sq + 9.0/4.0*g2sq))
	gpSq := 3.0 / 5.0 * g1sq
	l := st.Lambda
	out.Lambda = (24*l*l - 6*math.Pow(yt, 4) - 6*math.Pow(yb, 4) - 2*math.Pow(ye, 4) + (9.0/8.0)*g2sq*g2sq + (3.0/4.0)*g2sq*gpSq + (3.0/8.0)*gpSq*gpSq + (-9*g2sq-3*gpSq+12*yt*yt+12*yb*yb+4*ye*ye)*l) / (16.0 * math.Pi * math.Pi)
	return out
}

func addScaled(st, k RunningState, h float64) RunningState {
	return RunningState{U: [3]float64{st.U[0] + h*k.U[0], st.U[1] + h*k.U[1], st.U[2] + h*k.U[2]}, YTop: st.YTop + h*k.YTop, YBottom: st.YBottom + h*k.YBottom, YTau: st.YTau + h*k.YTau, Lambda: st.Lambda + h*k.Lambda}
}

func runSensitivityCases(s HeavySpectrum, cfg IntegratorConfig, central PhenomenologicalInputs, inputs []EmpiricalInput, centralFit DegenerateFit) []ScanCase {
	out := []ScanCase{{Name: "central", Perturbation: "none", Inputs: central, Fit: centralFit, Verdict: "central bottom/tau-complete fit"}}
	for _, inp := range inputs {
		if !inp.ScanEnabled || inp.Sigma == 0 {
			continue
		}
		for _, sign := range []float64{-1, 1} {
			vals := map[string]float64{}
			for _, x := range inputs {
				vals[x.Symbol] = x.Central
			}
			vals[inp.Symbol] = inp.Central + sign*inp.Sigma
			caseName := fmt.Sprintf("%s %+gσ", inp.Symbol, sign)
			ph := buildInputs(vals["alpha_s(M_Z)"], vals["m_t"], vals["m_b"], vals["m_tau"], vals["m_H"], vals["alpha_em^-1(M_Z)"], vals["sin^2(theta_W)"])
			fit := fitDegenerate(s, ph, cfg, caseName)
			out = append(out, ScanCase{Name: caseName, Perturbation: fmt.Sprintf("%s shifted by %+g %s", inp.Symbol, sign*inp.Sigma, inp.Unit), Inputs: ph, Fit: fit, DeltaLB: fit.LB - centralFit.LB, DeltaLStar: fit.LStar - centralFit.LStar, DeltaMBGeV: fit.MBGeV - centralFit.MBGeV, DeltaMStarGeV: fit.MStarGeV - centralFit.MStarGeV, DeltaMaxResidual: fit.MaxAbsResidual - centralFit.MaxAbsResidual, Verdict: "single-parameter 1σ perturbation; no tuning or refitting of empirical input beyond stated uncertainty"})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Fit.ResidualOverEpsilon > out[j].Fit.ResidualOverEpsilon })
	return out
}

func summarizeSensitivity(cases []ScanCase, central DegenerateFit) SensitivitySummary {
	s := SensitivitySummary{CasesAudited: len(cases), CentralMBGeV: central.MBGeV, CentralMStarGeV: central.MStarGeV, MBMinGeV: math.Inf(1), MStarMinGeV: math.Inf(1), ResidualOverEpsilonMin: math.Inf(1)}
	maxMBDelta, maxResDelta := -1.0, -1.0
	for _, c := range cases {
		f := c.Fit
		if f.Converged {
			s.ConvergedCases++
		}
		if f.MatchingPlausible {
			s.PlausibleCases++
		} else {
			s.BrokenEnvelopeCases++
		}
		if f.MBGeV < s.MBMinGeV {
			s.MBMinGeV = f.MBGeV
		}
		if f.MBGeV > s.MBMaxGeV {
			s.MBMaxGeV = f.MBGeV
		}
		if f.MStarGeV < s.MStarMinGeV {
			s.MStarMinGeV = f.MStarGeV
		}
		if f.MStarGeV > s.MStarMaxGeV {
			s.MStarMaxGeV = f.MStarGeV
		}
		if f.ResidualOverEpsilon < s.ResidualOverEpsilonMin {
			s.ResidualOverEpsilonMin = f.ResidualOverEpsilon
		}
		if f.ResidualOverEpsilon > s.ResidualOverEpsilonMax {
			s.ResidualOverEpsilonMax = f.ResidualOverEpsilon
			s.WorstCaseName = c.Name
			s.WorstResidualOverEpsilon = f.ResidualOverEpsilon
		}
		if d := math.Abs(f.MBGeV-central.MBGeV) / central.MBGeV; d > maxMBDelta && c.Name != "central" {
			maxMBDelta = d
			s.DominantScaleDriver = c.Name
		}
		if d := math.Abs(f.MaxAbsResidual - central.MaxAbsResidual); d > maxResDelta && c.Name != "central" {
			maxResDelta = d
			s.DominantResidualDriver = c.Name
		}
	}
	s.MBMinusGeV = central.MBGeV - s.MBMinGeV
	s.MBPlusGeV = s.MBMaxGeV - central.MBGeV
	s.MStarMinusGeV = central.MStarGeV - s.MStarMinGeV
	s.MStarPlusGeV = s.MStarMaxGeV - central.MStarGeV
	s.Verdict = "all 1σ scan cases preserve the matching plausibility envelope"
	if s.BrokenEnvelopeCases > 0 {
		s.Verdict = "at least one 1σ scan case breaks the matching plausibility envelope"
	}
	return s
}

func auditFirewall(g218 matchingcorrectionseal.Analysis, snap Gate218Snapshot) FirewallAudit {
	return FirewallAudit{Gate218Inherited: snap.Gate218Inherited, MatchingCorrectionSealActive: snap.MatchingCorrectionSealActive, ThresholdSpectrumSealInherited: snap.ThresholdSpectrumSealActive, EmpiricalCarrierSealInherited: g218.Firewall.EmpiricalCarrierSealInherited, LeptoquarkDynamicsSealInherited: g218.Firewall.LeptoquarkDynamicsSealInherited, EmpiricalLedgerQuarantined: g218.Firewall.EmpiricalLedgerQuarantined, InputUncertaintiesFiniteDerived: false, InputsTunedToForceZeroResidual: false, MatchingCorrectionsDerived: false, MatchingResidualPromoted: false, PhysicalPredictionClaimed: false, ProtonLifetimeComputed: false, RecommendedNextGate: "Gate 220 — experimental observability / PeV-threshold indirect-signature audit", OpenRequirements: []string{"finite derivation of matching corrections", "full precision conversion between pole and running masses", "higher-loop and scheme comparison", "experimental signature model for sealed heavy spectrum"}, Verdict: "input uncertainties are empirical scan parameters; no finite-core derivation or tuning is claimed"}
}

func buildTruth(f DegenerateFit, s SensitivitySummary) string {
	return fmt.Sprintf("Gate 219 upgrades the sealed single-scale two-loop audit with bottom/tau Yukawa completeness and propagates ±1σ empirical input uncertainties. The central fit gives M_B=%.9g GeV, M*=%.9g GeV, δ_req=%s, max|δ|/ε=%.6g. Across %d scan cases the envelope is preserved=%t; M_B range=[%.9g,%.9g] GeV and M* range=[%.9g,%.9g] GeV. These are conditional phenomenology, not finite-core predictions.", f.MBGeV, f.MStarGeV, f.RequiredDeltaMatch.String(), f.ResidualOverEpsilon, s.CasesAudited, s.BrokenEnvelopeCases == 0, s.MBMinGeV, s.MBMaxGeV, s.MStarMinGeV, s.MStarMaxGeV)
}

func positivityAudits(s HeavySpectrum, in PhenomenologicalInputs, cfg IntegratorConfig, lb, lstar float64) (bool, bool) {
	a := integrateTo(s, in, cfg, lb, lstar)
	b := integrateTo(s, in, cfg, lb, planckLogBound)
	return finiteState(a) && min3(a.U) > 1e-9, finiteState(b) && min3(b.U) > 1e-9
}
func validParams(p [2]float64) bool {
	return finite(p[0]) && finite(p[1]) && p[0] > 0.1 && p[1] > p[0]+0.01 && p[1] < planckLogBound
}
func finite(x float64) bool { return !math.IsNaN(x) && !math.IsInf(x, 0) }
func finiteState(st RunningState) bool {
	return finite(st.U[0]) && finite(st.U[1]) && finite(st.U[2]) && finite(st.YTop) && finite(st.YBottom) && finite(st.YTau) && finite(st.Lambda)
}
func maxAbs3(x [3]float64) float64 {
	return math.Max(math.Abs(x[0]), math.Max(math.Abs(x[1]), math.Abs(x[2])))
}
func min3(x [3]float64) float64 { return math.Min(x[0], math.Min(x[1], x[2])) }
func epsilonU() float64         { return 1.0 / (16.0 * math.Pi * math.Pi) }

func FormatGate218(g Gate218Snapshot) string {
	return fmt.Sprintf("gate218=%t matchSeal=%t spectrumSeal=%t plausible=%t bottomTauOmitted=%t M=(%.9g,%.9g) δ=%s overε=%.6g", g.Gate218Inherited, g.MatchingCorrectionSealActive, g.ThresholdSpectrumSealActive, g.MatchingPlausible, g.BottomTauOmitted, g.MBGeV, g.MStarGeV, g.RequiredDeltaMatch.String(), g.ResidualOverEpsilon)
}
func FormatInputs(xs []EmpiricalInput) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s=%.9g±%.3g %s scan=%t finite=%t", x.Symbol, x.Central, x.Sigma, x.Unit, x.ScanEnabled, x.FiniteCoreDerived))
	}
	return strings.Join(parts, "; ")
}
func FormatPhenomenologicalInputs(i PhenomenologicalInputs) string {
	return fmt.Sprintf("alphaInv=%.6g sin2=%.6g alphaS=%.6g mt=%.6g mb=%.6g mtau=%.6g mh=%.6g yt=%.9g yb=%.9g ytau=%.9g lambda=%.9g finite=%t", i.AlphaEMInvMZ, i.Sin2ThetaWMZ, i.AlphaSMZ, i.TopPoleMassGeV, i.BottomMassGeV, i.TauMassGeV, i.HiggsMassGeV, i.InitialYTop, i.InitialYBottom, i.InitialYTau, i.InitialLambda, i.FiniteCoreDerived)
}
func FormatSpectrum(s HeavySpectrum) string {
	return fmt.Sprintf("rows=[%s %s Δb=%s; %s %s Δb=%s] totalΔ=%s Bheavy=%s target=%.3g", s.Row1Name, s.Row1Rep, s.Row1DeltaB.String(), s.Row2Name, s.Row2Rep, s.Row2DeltaB.String(), s.TotalDeltaB.String(), s.HeavyTwoLoop.String(), s.TargetU)
}
func FormatConfig(c IntegratorConfig) string {
	return fmt.Sprintf("method=%s steps=%d min=%d iters=%d y=%t bottomTau=%t lambda=%t matchingUsed=%t", c.Method, c.StepsPerLogUnit, c.MinimumStepsPerSegment, c.MaxCoordinateIterations, c.YukawaTermsIncluded, c.BottomTauIncluded, c.LambdaRunningIncluded, c.MatchingCorrectionsUsed)
}
func FormatCompleteness(c CompletenessAudit) string {
	return fmt.Sprintf("top=%t bottom=%t tau=%t lambda=%t fullMatrices=%t lightIgnored=%t heavyYukawa=%t coeffs=[t%s b%s tau%s]", c.TopYukawaIncluded, c.BottomYukawaIncluded, c.TauYukawaIncluded, c.HiggsQuarticIncluded, c.FullYukawaMatricesDerived, c.OtherFermionYukawasIgnored, c.HeavyYukawaCouplingsAdded, c.GaugeYukawaCoefficientsTop.String(), c.GaugeYukawaCoefficientsBottom.String(), c.GaugeYukawaCoefficientsTau.String())
}
func FormatSeal(s MatchingSealAudit) string {
	return fmt.Sprintf("inherited=%t quarantined=%t derived=%t promoted=%t epsilon=%.12g", s.MatchingCorrectionSealInherited, s.RequiredResidualQuarantined, s.MatchingCorrectionsDerived, s.MatchingResidualPromoted, s.EpsilonU)
}
func FormatFit(f DegenerateFit) string {
	return fmt.Sprintf("case=%s L=(MB=%.9g,M*=%.9g) M=(MB=%.9g,M*=%.9g) U=(%.12g,%.12g,%.12g) y*=(%.6g,%.6g,%.6g) λ*=%.6g residual=%s required=%s max=%.9g overε=%.6g it=%d conv=%t ordered=%t subPlanck=%t positive=%t noLandau=%t plausible=%t", f.CaseName, f.LB, f.LStar, f.MBGeV, f.MStarGeV, f.BoundaryU[0], f.BoundaryU[1], f.BoundaryU[2], f.BoundaryYTop, f.BoundaryYBottom, f.BoundaryYTau, f.BoundaryLambda, f.Residual.String(), f.RequiredDeltaMatch.String(), f.MaxAbsResidual, f.ResidualOverEpsilon, f.OptimizerIterations, f.Converged, f.ScaleOrdered, f.SubPlanck, f.PositiveToBoundary, f.NoLandauBelowPlanck, f.MatchingPlausible)
}
func FormatScanCase(c ScanCase) string {
	return fmt.Sprintf("%s [%s] %s ΔM=(MB %.6g,M* %.6g) ΔL=(%.6g,%.6g) Δmaxδ=%.6g", c.Name, c.Perturbation, FormatFit(c.Fit), c.DeltaMBGeV, c.DeltaMStarGeV, c.DeltaLB, c.DeltaLStar, c.DeltaMaxResidual)
}
func FormatScanTop(cases []ScanCase, n int) string {
	if n > len(cases) {
		n = len(cases)
	}
	parts := make([]string, 0, n)
	for i := 0; i < n; i++ {
		parts = append(parts, FormatScanCase(cases[i]))
	}
	return strings.Join(parts, "\n")
}
func FormatSensitivity(s SensitivitySummary) string {
	return fmt.Sprintf("cases=%d conv=%d plausible=%d broken=%d MB=[%.9g,%.9g] central=%.9g -%.6g +%.6g M*=[%.9g,%.9g] central=%.9g -%.6g +%.6g overε=[%.6g,%.6g] worst=%s %.6g drivers=(scale %s,residual %s)", s.CasesAudited, s.ConvergedCases, s.PlausibleCases, s.BrokenEnvelopeCases, s.MBMinGeV, s.MBMaxGeV, s.CentralMBGeV, s.MBMinusGeV, s.MBPlusGeV, s.MStarMinGeV, s.MStarMaxGeV, s.CentralMStarGeV, s.MStarMinusGeV, s.MStarPlusGeV, s.ResidualOverEpsilonMin, s.ResidualOverEpsilonMax, s.WorstCaseName, s.WorstResidualOverEpsilon, s.DominantScaleDriver, s.DominantResidualDriver)
}
func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("gate218=%t matchSeal=%t spectrumSeal=%t carrier=%t lq=%t ledger=%t inputFinite=%t tuned=%t matchingDerived=%t promoted=%t prediction=%t lifetime=%t next=%q", f.Gate218Inherited, f.MatchingCorrectionSealActive, f.ThresholdSpectrumSealInherited, f.EmpiricalCarrierSealInherited, f.LeptoquarkDynamicsSealInherited, f.EmpiricalLedgerQuarantined, f.InputUncertaintiesFiniteDerived, f.InputsTunedToForceZeroResidual, f.MatchingCorrectionsDerived, f.MatchingResidualPromoted, f.PhysicalPredictionClaimed, f.ProtonLifetimeComputed, f.RecommendedNextGate)
}
func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d gate218=%t bottomTau=%t cases=%d plausible=%d preserved=%t status=%s", s.TestsAudited, s.Gate218Inherited, s.BottomTauComplete, s.CasesAudited, s.PlausibleCases, s.EnvelopePreservedAt1Sigma, s.Status)
}
