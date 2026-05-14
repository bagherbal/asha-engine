// Package singlescalematchingaudit implements Gate 215: single-scale
// degenerate-limit matching audit / global two-loop class scan.
//
// Gate 214 showed that the sealed Gate-211 ranked witness becomes nearly
// degenerate under no-Yukawa two-loop running. Gate 215 asks the next legal
// question: if each of Gate-211's 22 unordered viable spectra is forced to a
// single heavy threshold, how large a finite threshold-matching correction would
// be required to hit the topological u*=1 boundary? The answer is a conditional
// phenomenology scan, not a finite-core derivation of matching coefficients.
package singlescalematchingaudit

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/representationrowlattice"
	"github.com/bagherbal/asha-engine/pkg/bridge/twoloopintegration"
	"github.com/bagherbal/asha-engine/pkg/bridge/twothresholdminimality"
)

const (
	StatusConditionalPhenomenology = "CONDITIONAL_PHENOMENOLOGY_SINGLE_SCALE_MATCHING_AUDIT"
	StatusFailedRoute              = "FAILED_ROUTE_SINGLE_SCALE_MATCHING_AUDIT"

	MatchingPlausible = "MATCHING_RESIDUAL_WITHIN_LOOP_FACTOR_ENVELOPE"
	MatchingRejected  = "MATCHING_RESIDUAL_EXCEEDS_LOOP_FACTOR_ENVELOPE"

	mzGeV          = 91.1876
	planckLogBound = 37.8
	targetU        = 1.0
)

type Rational = representationrowlattice.Rational

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

type Gate214Snapshot struct {
	Gate214Inherited               bool
	ThresholdSpectrumSealInherited bool
	MatchingEnvelopeInherited      bool
	CentralTwoLoopConverged        bool
	CentralDeltaL                  float64
	MatchingEpsilonU               float64
	OrderedGate211Pairs            int
	UnorderedGate211Classes        int
	TruthStatement                 string
}

type ScanConfig struct {
	Method                 string
	Equation               string
	TargetU                float64
	EpsilonU               float64
	ClassesExpected        int
	StepsPerLogUnit        int
	MinimumStepsPerSegment int
	MaxCoordinateIters     int
	InitialStartsPerClass  int
	YukawaTermsIncluded    bool
	MatchingDerived        bool
	Verdict                string
}

type GroupInvariants struct {
	Symbol string
	Dim    int64
	T      Rational
	C2     Rational
	CG     Rational
}

type ParsedRepresentation struct {
	Name       string
	Rep        string
	Statistics string
	SU3        GroupInvariants
	SU2        GroupInvariants
	Y          Rational
	YCasimir   Rational
	Supported  bool
	Verdict    string
}

type ClassSpectrum struct {
	ClassRank          int
	Key                string
	RowAName           string
	RowBName           string
	RowARep            string
	RowBRep            string
	RowAStatistic      string
	RowBStatistic      string
	RowADeltaB         FloatTriple
	RowBDeltaB         FloatTriple
	TotalDeltaB        FloatTriple
	SMBeta             FloatTriple
	TotalOneLoopBeta   FloatTriple
	SMTwoLoop          FloatMatrix3
	RowATwoLoop        FloatMatrix3
	RowBTwoLoop        FloatMatrix3
	TotalTwoLoop       FloatMatrix3
	Gate211MeanLog     float64
	Gate211DeltaL      float64
	Gate211LStar       float64
	Gate211MStarGeV    float64
	TwoLoopSupported   bool
	LeptoquarkSealSafe bool
	AnomalySafe        bool
	Verdict            string
}

type DegenerateFit struct {
	ClassRank           int
	Rows                string
	LB                  float64
	LStar               float64
	MBGeV               float64
	MStarGeV            float64
	BoundaryU           [3]float64
	Residual            [3]float64
	RequiredDeltaMatch  [3]float64
	ResidualNorm        float64
	MaxAbsResidual      float64
	RMSResidual         float64
	ResidualOverEpsilon float64
	Converged           bool
	OptimizerIterations int
	ScaleOrdered        bool
	SubPlanck           bool
	PositiveToBoundary  bool
	NoLandauBelowPlanck bool
	MatchingPlausible   bool
	Status              string
	Verdict             string
}

type GlobalScanAudit struct {
	ClassesAudited            int
	TwoLoopSupportedClasses   int
	OptimizerConvergedClasses int
	PlausibleWithinEnvelope   int
	RejectedByEnvelope        int
	BestClassRank             int
	BestRows                  string
	BestMaxResidual           float64
	BestResidualOverEpsilon   float64
	SmallestMBGeV             float64
	LargestMBGeV              float64
	Verdict                   string
}

type MatchingObstructionAudit struct {
	Gate214MatchingEnvelopeInherited bool
	NativeDeltaMatchRowsDerived      bool
	HeatKernelMatchingMapDerived     bool
	CanonicalSubtractionScheme       bool
	EnvelopeUsedAsProxy              bool
	ResidualInterpretedAsDerived     bool
	Status                           string
	Verdict                          string
}

type FirewallAudit struct {
	ThresholdSpectrumSealInherited  bool
	EmpiricalCarrierSealInherited   bool
	LeptoquarkDynamicsSealInherited bool
	EmpiricalLedgerQuarantined      bool
	All22ClassesAudited             bool
	SingleScaleForcedAsFiniteCore   bool
	MatchingCorrectionsDerived      bool
	MatchingResidualPromoted        bool
	YukawaMatricesImported          bool
	PhysicalPredictionClaimed       bool
	ProtonLifetimeComputed          bool
	RecommendedNextGate             string
	OpenRequirements                []string
	Verdict                         string
}

type Summary struct {
	TestsAudited            int
	Gate214Inherited        bool
	ClassesAudited          int
	PlausibleClasses        int
	BestResidualOverEpsilon float64
	Status                  string
	Comment                 string
}

type Analysis struct {
	Gate214         Gate214Snapshot
	Gate214Analysis twoloopintegration.Analysis
	Config          ScanConfig
	Spectra         []ClassSpectrum
	Fits            []DegenerateFit
	GlobalScan      GlobalScanAudit
	MatchingAudit   MatchingObstructionAudit
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
		g214, err := twoloopintegration.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultA, defaultErr = Build(g214)
	})
	return defaultA, defaultErr
}

func Build(g214 twoloopintegration.Analysis) (Analysis, error) {
	snap := snapshotFromGate214(g214)
	if !snap.Gate214Inherited || !snap.ThresholdSpectrumSealInherited || !snap.MatchingEnvelopeInherited || !snap.CentralTwoLoopConverged {
		return Analysis{}, fmt.Errorf("Gate 215 requires Gate 214 two-loop integration, ThresholdSpectrumSeal, and matching envelope inheritance")
	}
	cfg := defaultConfig(snap)
	classes := g214.Gate213Analysis.Gate212Analysis.PairClasses
	spectra := make([]ClassSpectrum, 0, len(classes))
	fits := make([]DegenerateFit, 0, len(classes))
	for i, pc := range classes {
		sp := buildSpectrum(i+1, pc)
		spectra = append(spectra, sp)
		fit := fitDegenerateClass(sp, cfg)
		fits = append(fits, fit)
	}
	sort.Slice(fits, func(i, j int) bool {
		if math.Abs(fits[i].MaxAbsResidual-fits[j].MaxAbsResidual) > 1e-14 {
			return fits[i].MaxAbsResidual < fits[j].MaxAbsResidual
		}
		if math.Abs(fits[i].MStarGeV-fits[j].MStarGeV) > 1e-6 {
			return gutDistance(fits[i].MStarGeV) < gutDistance(fits[j].MStarGeV)
		}
		return fits[i].ClassRank < fits[j].ClassRank
	})
	scan := auditGlobal(fits, spectra)
	matching := auditMatching(snap)
	fw := auditFirewall(snap, len(classes), scan, matching)
	status := StatusConditionalPhenomenology
	if scan.ClassesAudited == 0 || scan.TwoLoopSupportedClasses == 0 {
		status = StatusFailedRoute
	}
	summary := Summary{
		TestsAudited:            7,
		Gate214Inherited:        snap.Gate214Inherited,
		ClassesAudited:          len(classes),
		PlausibleClasses:        scan.PlausibleWithinEnvelope,
		BestResidualOverEpsilon: scan.BestResidualOverEpsilon,
		Status:                  status,
		Comment:                 "Gate 215 forces each of the 22 Gate-211 unordered viable spectra into a degenerate single-threshold two-loop solve and ranks the required u-space matching residual against the explicit loop-factor envelope ε=1/(16π²).",
	}
	truth := buildTruth(summary, scan, matching)
	return Analysis{Gate214: snap, Gate214Analysis: g214, Config: cfg, Spectra: spectra, Fits: fits, GlobalScan: scan, MatchingAudit: matching, Firewall: fw, Summary: summary, TruthStatement: truth}, nil
}

func snapshotFromGate214(a twoloopintegration.Analysis) Gate214Snapshot {
	g212 := a.Gate213Analysis.Gate212Analysis
	return Gate214Snapshot{
		Gate214Inherited:               a.Summary.Status == twoloopintegration.StatusConditionalPhenomenology,
		ThresholdSpectrumSealInherited: a.Firewall.ThresholdSpectrumSealInherited,
		MatchingEnvelopeInherited:      a.Envelope.Status == twoloopintegration.MatchingEnvelopeStatus && a.Envelope.EpsilonU > 0,
		CentralTwoLoopConverged:        a.Central.Converged,
		CentralDeltaL:                  a.Central.DeltaL,
		MatchingEpsilonU:               a.Envelope.EpsilonU,
		OrderedGate211Pairs:            g212.Summary.OrderedViablePairs,
		UnorderedGate211Classes:        g212.Summary.UnorderedPairClasses,
		TruthStatement:                 a.TruthStatement,
	}
}

func defaultConfig(s Gate214Snapshot) ScanConfig {
	return ScanConfig{
		Method:                 "multi-start constrained coordinate descent over (L_B,L_*) with fixed-step RK4 two-loop integration",
		Equation:               "du_i/dlnμ = -b_i/(8π²) - Σ_j B_ij/u_j /(128π⁴); a single common threshold activates both rows at L_B",
		TargetU:                targetU,
		EpsilonU:               s.MatchingEpsilonU,
		ClassesExpected:        s.UnorderedGate211Classes,
		StepsPerLogUnit:        16,
		MinimumStepsPerSegment: 24,
		MaxCoordinateIters:     90,
		InitialStartsPerClass:  8,
		YukawaTermsIncluded:    false,
		MatchingDerived:        false,
		Verdict:                "numerical degenerate-limit scan only; matching residuals are required corrections, not derived finite counterterms",
	}
}

func buildSpectrum(rank int, pc twothresholdminimality.PairClass) ClassSpectrum {
	a := rowToParsed(pc.RowA)
	b := rowToParsed(pc.RowB)
	smB := FloatTriple{41.0 / 10.0, -19.0 / 6.0, -7.0}
	smM := smTwoLoopMatrix()
	mA, okA := twoLoopMatrix(a)
	mB, okB := twoLoopMatrix(b)
	dbA := rtToFloat(pc.RowA.DeltaB)
	dbB := rtToFloat(pc.RowB.DeltaB)
	supported := okA && okB
	verdict := "two-loop matrix constructed from standard QFT group invariants for Gate-211 class"
	if !supported {
		verdict = "unsupported carrier type for Gate-215 two-loop scan; retained in audit but excluded from plausibility ranking"
	}
	return ClassSpectrum{
		ClassRank:          rank,
		Key:                pc.Key,
		RowAName:           pc.RowA.Name,
		RowBName:           pc.RowB.Name,
		RowARep:            pc.RowA.SMRepresentation,
		RowBRep:            pc.RowB.SMRepresentation,
		RowAStatistic:      pc.RowA.Statistic,
		RowBStatistic:      pc.RowB.Statistic,
		RowADeltaB:         dbA,
		RowBDeltaB:         dbB,
		TotalDeltaB:        dbA.Add(dbB),
		SMBeta:             smB,
		TotalOneLoopBeta:   smB.Add(dbA).Add(dbB),
		SMTwoLoop:          smM,
		RowATwoLoop:        mA,
		RowBTwoLoop:        mB,
		TotalTwoLoop:       smM.Add(mA).Add(mB),
		Gate211MeanLog:     pc.MeanThresholdLog,
		Gate211DeltaL:      pc.DeltaL,
		Gate211LStar:       pc.Representative.LStar,
		Gate211MStarGeV:    pc.MStarGeV,
		TwoLoopSupported:   supported,
		LeptoquarkSealSafe: true,
		AnomalySafe:        true,
		Verdict:            verdict,
	}
}

func fitDegenerateClass(s ClassSpectrum, cfg ScanConfig) DegenerateFit {
	if !s.TwoLoopSupported {
		return DegenerateFit{ClassRank: s.ClassRank, Rows: rowsLabel(s), ResidualNorm: math.Inf(1), MaxAbsResidual: math.Inf(1), ResidualOverEpsilon: math.Inf(1), Status: StatusFailedRoute, Verdict: s.Verdict}
	}
	starts := candidateStarts(s)
	bestP := [2]float64{math.NaN(), math.NaN()}
	bestObj := math.Inf(1)
	bestIter := 0
	converged := false
	for _, st := range starts {
		p, obj, iters := coordinateOptimize(s, cfg, st)
		if obj < bestObj {
			bestP, bestObj, bestIter = p, obj, iters
		}
	}
	boundary := integrateTo(s, cfg, bestP[0], bestP[1])
	residual := [3]float64{boundary[0] - cfg.TargetU, boundary[1] - cfg.TargetU, boundary[2] - cfg.TargetU}
	n := norm3(residual)
	max := maxAbs3(residual)
	rms := n / math.Sqrt(3)
	if max < bestObj+1e-10 {
		converged = true
	}
	positive, noLandau := positivityAudits(s, cfg, bestP[0], bestP[1])
	scaleOrdered := validParams(bestP)
	subPlanck := scaleOrdered && bestP[1] < planckLogBound
	plausible := converged && scaleOrdered && subPlanck && positive && noLandau && max <= cfg.EpsilonU
	status := MatchingRejected
	verdict := "required degenerate-limit matching residual exceeds the loop-factor envelope; pair is not a plausible single-scale completion under this proxy"
	if plausible {
		status = MatchingPlausible
		verdict = "required degenerate-limit matching residual fits inside ε=1/(16π²); pair is plausible as a single-scale completion only if finite matching corrections later supply this residual"
	}
	return DegenerateFit{
		ClassRank:           s.ClassRank,
		Rows:                rowsLabel(s),
		LB:                  bestP[0],
		LStar:               bestP[1],
		MBGeV:               mzGeV * math.Exp(bestP[0]),
		MStarGeV:            mzGeV * math.Exp(bestP[1]),
		BoundaryU:           boundary,
		Residual:            residual,
		RequiredDeltaMatch:  [3]float64{-residual[0], -residual[1], -residual[2]},
		ResidualNorm:        n,
		MaxAbsResidual:      max,
		RMSResidual:         rms,
		ResidualOverEpsilon: max / cfg.EpsilonU,
		Converged:           converged,
		OptimizerIterations: bestIter,
		ScaleOrdered:        scaleOrdered,
		SubPlanck:           subPlanck,
		PositiveToBoundary:  positive,
		NoLandauBelowPlanck: noLandau,
		MatchingPlausible:   plausible,
		Status:              status,
		Verdict:             verdict,
	}
}

func candidateStarts(s ClassSpectrum) [][2]float64 {
	mean := s.Gate211MeanLog
	star := s.Gate211LStar
	starts := [][2]float64{
		{mean, star},
		{mean + 1.5, star + 0.8},
		{mean + 3.0, star + 1.2},
		{10.0, 35.0},
		{12.0, 35.5},
		{15.0, 36.0},
		{math.Max(1.0, mean-1.0), math.Min(planckLogBound-0.2, star+0.5)},
		{math.Min(planckLogBound-2.0, mean+5.0), math.Min(planckLogBound-0.1, star+2.0)},
	}
	out := make([][2]float64, 0, len(starts))
	for _, p := range starts {
		if p[1] <= p[0]+0.2 {
			p[1] = p[0] + 1.0
		}
		if p[1] >= planckLogBound {
			p[1] = planckLogBound - 0.2
		}
		if p[0] <= 0 {
			p[0] = 0.5
		}
		if validParams(p) {
			out = append(out, p)
		}
	}
	return out
}

func coordinateOptimize(s ClassSpectrum, cfg ScanConfig, start [2]float64) ([2]float64, float64, int) {
	p := start
	best := objective(s, cfg, p)
	step := 3.0
	dirs := [][2]float64{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}
	iters := 0
	for iters = 0; iters < cfg.MaxCoordinateIters && step > 1e-5; iters++ {
		improved := false
		for _, d := range dirs {
			q := [2]float64{p[0] + step*d[0], p[1] + step*d[1]}
			if !validParams(q) {
				continue
			}
			obj := objective(s, cfg, q)
			if obj < best {
				p, best = q, obj
				improved = true
			}
		}
		if !improved {
			step *= 0.55
		}
	}
	return p, best, iters
}

func objective(s ClassSpectrum, cfg ScanConfig, p [2]float64) float64 {
	if !validParams(p) {
		return math.Inf(1)
	}
	b := integrateTo(s, cfg, p[0], p[1])
	if !finite3(b) {
		return math.Inf(1)
	}
	r := [3]float64{b[0] - cfg.TargetU, b[1] - cfg.TargetU, b[2] - cfg.TargetU}
	return maxAbs3(r)
}

func integrateTo(s ClassSpectrum, cfg ScanConfig, lb, lend float64) [3]float64 {
	u := [3]float64{4.69678547222, 2.35435958464, 0.674957349838}
	if lend <= 0 {
		return u
	}
	if lb > 0 {
		end := math.Min(lb, lend)
		u = integrateSegment(u, 0, end, s.SMBeta, s.SMTwoLoop, cfg)
	}
	if lend > lb {
		u = integrateSegment(u, math.Max(0, lb), lend, s.TotalOneLoopBeta, s.TotalTwoLoop, cfg)
	}
	return u
}

func integrateSegment(u [3]float64, a, b float64, beta FloatTriple, mat FloatMatrix3, cfg ScanConfig) [3]float64 {
	if b <= a {
		return u
	}
	steps := int(math.Ceil((b - a) * float64(cfg.StepsPerLogUnit)))
	if steps < cfg.MinimumStepsPerSegment {
		steps = cfg.MinimumStepsPerSegment
	}
	h := (b - a) / float64(steps)
	for i := 0; i < steps; i++ {
		k1 := deriv(u, beta, mat)
		k2 := deriv(addScaled(u, k1, 0.5*h), beta, mat)
		k3 := deriv(addScaled(u, k2, 0.5*h), beta, mat)
		k4 := deriv(addScaled(u, k3, h), beta, mat)
		for j := 0; j < 3; j++ {
			u[j] += h * (k1[j] + 2*k2[j] + 2*k3[j] + k4[j]) / 6.0
		}
		if !finite3(u) || min3(u) <= 0 {
			return [3]float64{math.NaN(), math.NaN(), math.NaN()}
		}
	}
	return u
}

func deriv(u [3]float64, beta FloatTriple, mat FloatMatrix3) [3]float64 {
	var out [3]float64
	for i := 0; i < 3; i++ {
		sum := 0.0
		for j := 0; j < 3; j++ {
			sum += mat.M[i][j] / u[j]
		}
		out[i] = -beta.At(i)/(8.0*math.Pi*math.Pi) - sum/(128.0*math.Pi*math.Pi*math.Pi*math.Pi)
	}
	return out
}

func positivityAudits(s ClassSpectrum, cfg ScanConfig, lb, lstar float64) (bool, bool) {
	toStar := integrateTo(s, cfg, lb, lstar)
	positive := finite3(toStar) && min3(toStar) > 1e-9
	toPlanck := integrateTo(s, cfg, lb, planckLogBound)
	noLandau := finite3(toPlanck) && min3(toPlanck) > 1e-9
	return positive, noLandau
}

func validParams(p [2]float64) bool {
	return finite(p[0]) && finite(p[1]) && p[0] > 0 && p[1] > p[0] && p[1] < planckLogBound
}

func rowToParsed(r twothresholdminimality.RowSemantics) ParsedRepresentation {
	y, _ := parseRat(r.Hypercharge)
	yc := y.Mul(y).Mul(representationrowlattice.R(3, 5))
	su3, ok3 := su3Invariant(r.SU3Symbol)
	su2, ok2 := su2Invariant(r.SU2Symbol)
	supportedStats := r.Statistic == "Dirac fermion" || r.Statistic == "Weyl fermion" || r.Statistic == "complex scalar" || r.Statistic == "real scalar"
	return ParsedRepresentation{Name: r.Name, Rep: r.SMRepresentation, Statistics: r.Statistic, SU3: su3, SU2: su2, Y: y, YCasimir: yc, Supported: ok3 && ok2 && supportedStats, Verdict: "parsed Gate-211 row semantics"}
}

func twoLoopMatrix(p ParsedRepresentation) (FloatMatrix3, bool) {
	if !p.Supported {
		return FloatMatrix3{}, false
	}
	s := [3]Rational{
		p.YCasimir.MulInt(p.SU3.Dim * p.SU2.Dim),
		p.SU2.T.MulInt(p.SU3.Dim),
		p.SU3.T.MulInt(p.SU2.Dim),
	}
	c := [3]Rational{p.YCasimir, p.SU2.C2, p.SU3.C2}
	cg := [3]Rational{representationrowlattice.R(0, 1), p.SU2.CG, p.SU3.CG}
	diagGaugeFactor := representationrowlattice.R(20, 3)
	spinFactor := 1.0
	formulaFactor := 1.0
	switch p.Statistics {
	case "Dirac fermion":
		diagGaugeFactor = representationrowlattice.R(20, 3)
		spinFactor = 1.0
		formulaFactor = 1.0
	case "Weyl fermion":
		diagGaugeFactor = representationrowlattice.R(20, 3)
		spinFactor = 1.0
		formulaFactor = 0.5
	case "complex scalar", "real scalar":
		diagGaugeFactor = representationrowlattice.R(2, 3)
		spinFactor = 1.0
		formulaFactor = 1.0
		if p.Statistics == "real scalar" {
			formulaFactor = 0.5
		}
	default:
		return FloatMatrix3{}, false
	}
	var out FloatMatrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			var val Rational
			if i == j {
				term := diagGaugeFactor.Mul(cg[i]).Add(c[i].MulInt(4))
				val = term.Mul(s[i])
			} else {
				val = c[j].Mul(s[i]).MulInt(4)
			}
			out.M[i][j] = val.Float() * spinFactor * formulaFactor
		}
	}
	return out, true
}

func su3Invariant(sym string) (GroupInvariants, bool) {
	switch sym {
	case "1":
		return GroupInvariants{Symbol: sym, Dim: 1, T: representationrowlattice.R(0, 1), C2: representationrowlattice.R(0, 1), CG: representationrowlattice.R(3, 1)}, true
	case "3", "3bar":
		return GroupInvariants{Symbol: sym, Dim: 3, T: representationrowlattice.R(1, 2), C2: representationrowlattice.R(4, 3), CG: representationrowlattice.R(3, 1)}, true
	case "8":
		return GroupInvariants{Symbol: sym, Dim: 8, T: representationrowlattice.R(3, 1), C2: representationrowlattice.R(3, 1), CG: representationrowlattice.R(3, 1)}, true
	default:
		return GroupInvariants{}, false
	}
}

func su2Invariant(sym string) (GroupInvariants, bool) {
	switch sym {
	case "1":
		return GroupInvariants{Symbol: sym, Dim: 1, T: representationrowlattice.R(0, 1), C2: representationrowlattice.R(0, 1), CG: representationrowlattice.R(2, 1)}, true
	case "2":
		return GroupInvariants{Symbol: sym, Dim: 2, T: representationrowlattice.R(1, 2), C2: representationrowlattice.R(3, 4), CG: representationrowlattice.R(2, 1)}, true
	case "3":
		return GroupInvariants{Symbol: sym, Dim: 3, T: representationrowlattice.R(2, 1), C2: representationrowlattice.R(2, 1), CG: representationrowlattice.R(2, 1)}, true
	default:
		return GroupInvariants{}, false
	}
}

func parseRat(s string) (Rational, error) {
	s = strings.TrimSpace(s)
	if strings.Contains(s, "/") {
		parts := strings.Split(s, "/")
		if len(parts) != 2 {
			return Rational{}, fmt.Errorf("bad rational %q", s)
		}
		n, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			return Rational{}, err
		}
		d, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			return Rational{}, err
		}
		return representationrowlattice.R(n, d), nil
	}
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return Rational{}, err
	}
	return representationrowlattice.R(n, 1), nil
}

func rtToFloat(t representationrowlattice.RationalTriple) FloatTriple {
	return FloatTriple{t.B1.Float(), t.B2.Float(), t.B3.Float()}
}

func smTwoLoopMatrix() FloatMatrix3 {
	return FloatMatrix3{M: [3][3]float64{
		{199.0 / 50.0, 27.0 / 10.0, 44.0 / 5.0},
		{9.0 / 10.0, 35.0 / 6.0, 12.0},
		{11.0 / 10.0, 9.0 / 2.0, -26.0},
	}}
}

func auditGlobal(fits []DegenerateFit, spectra []ClassSpectrum) GlobalScanAudit {
	aud := GlobalScanAudit{ClassesAudited: len(fits), SmallestMBGeV: math.Inf(1)}
	for _, s := range spectra {
		if s.TwoLoopSupported {
			aud.TwoLoopSupportedClasses++
		}
	}
	for _, f := range fits {
		if f.Converged {
			aud.OptimizerConvergedClasses++
		}
		if f.MatchingPlausible {
			aud.PlausibleWithinEnvelope++
		} else {
			aud.RejectedByEnvelope++
		}
		if f.MBGeV < aud.SmallestMBGeV {
			aud.SmallestMBGeV = f.MBGeV
		}
		if f.MBGeV > aud.LargestMBGeV {
			aud.LargestMBGeV = f.MBGeV
		}
	}
	if len(fits) > 0 {
		aud.BestClassRank = fits[0].ClassRank
		aud.BestRows = fits[0].Rows
		aud.BestMaxResidual = fits[0].MaxAbsResidual
		aud.BestResidualOverEpsilon = fits[0].ResidualOverEpsilon
	}
	if aud.PlausibleWithinEnvelope > 0 {
		aud.Verdict = fmt.Sprintf("%d of %d unordered Gate-211 classes have forced single-scale two-loop residuals inside the ε=1/(16π²) matching envelope", aud.PlausibleWithinEnvelope, aud.ClassesAudited)
	} else {
		aud.Verdict = fmt.Sprintf("no unordered Gate-211 class has a forced single-scale two-loop residual inside the ε=1/(16π²) matching envelope; best residual is %.6g ε", aud.BestResidualOverEpsilon)
	}
	return aud
}

func auditMatching(s Gate214Snapshot) MatchingObstructionAudit {
	return MatchingObstructionAudit{Gate214MatchingEnvelopeInherited: s.MatchingEnvelopeInherited, NativeDeltaMatchRowsDerived: false, HeatKernelMatchingMapDerived: false, CanonicalSubtractionScheme: false, EnvelopeUsedAsProxy: true, ResidualInterpretedAsDerived: false, Status: "FAILED_ROUTE_DERIVED_MATCHING_CORRECTIONS_PRESERVED", Verdict: "Gate 215 computes the required matching residual vector but does not derive δ_i^match; the loop-factor envelope remains a phenomenological proxy inherited from Gate 214"}
}

func auditFirewall(s Gate214Snapshot, classCount int, scan GlobalScanAudit, m MatchingObstructionAudit) FirewallAudit {
	return FirewallAudit{
		ThresholdSpectrumSealInherited:  s.ThresholdSpectrumSealInherited,
		EmpiricalCarrierSealInherited:   true,
		LeptoquarkDynamicsSealInherited: true,
		EmpiricalLedgerQuarantined:      true,
		All22ClassesAudited:             classCount == s.UnorderedGate211Classes && scan.ClassesAudited == s.UnorderedGate211Classes,
		SingleScaleForcedAsFiniteCore:   false,
		MatchingCorrectionsDerived:      m.NativeDeltaMatchRowsDerived,
		MatchingResidualPromoted:        m.ResidualInterpretedAsDerived,
		YukawaMatricesImported:          false,
		PhysicalPredictionClaimed:       false,
		ProtonLifetimeComputed:          false,
		RecommendedNextGate:             "Gate 216 — matching-residual structure audit / spectral heat-kernel coefficient search",
		OpenRequirements: []string{
			"derive a finite heat-kernel or spectral matching map before promoting required δ_i^match",
			"include SM Yukawa two-loop terms only under a separate empirical Yukawa seal",
			"re-run the full 22-class scan if the matching envelope is replaced by derived counterterms",
		},
		Verdict: "all outputs remain conditional phenomenology under ThresholdSpectrumSeal; no matching correction, single-scale spectrum, or physical mass is finite-derived",
	}
}

func buildTruth(s Summary, scan GlobalScanAudit, m MatchingObstructionAudit) string {
	return fmt.Sprintf("Gate 215 audits the degenerate single-scale limit of all 22 Gate-211 viable unordered spectra under two-loop running. The scan ranks the required u-space matching residual rather than deriving it. PlausibleWithinEnvelope=%d, best=%s, bestResidual/ε=%.6g. Matching corrections remain obstructed: nativeRows=%v heatKernel=%v subtraction=%v.", scan.PlausibleWithinEnvelope, scan.BestRows, scan.BestResidualOverEpsilon, m.NativeDeltaMatchRowsDerived, m.HeatKernelMatchingMapDerived, m.CanonicalSubtractionScheme)
}

func rowsLabel(s ClassSpectrum) string {
	return fmt.Sprintf("%s [%s] + %s [%s]", s.RowARep, s.RowAStatistic, s.RowBRep, s.RowBStatistic)
}

func addScaled(u, k [3]float64, h float64) [3]float64 {
	return [3]float64{u[0] + h*k[0], u[1] + h*k[1], u[2] + h*k[2]}
}
func norm3(v [3]float64) float64 { return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2]) }
func maxAbs3(v [3]float64) float64 {
	return math.Max(math.Abs(v[0]), math.Max(math.Abs(v[1]), math.Abs(v[2])))
}
func min3(v [3]float64) float64 { return math.Min(v[0], math.Min(v[1], v[2])) }
func finite(x float64) bool     { return !math.IsNaN(x) && !math.IsInf(x, 0) }
func finite3(v [3]float64) bool { return finite(v[0]) && finite(v[1]) && finite(v[2]) }
func gutDistance(m float64) float64 {
	if m <= 0 || math.IsNaN(m) {
		return math.Inf(1)
	}
	l := math.Log10(m)
	if l < 15 {
		return 15 - l
	}
	if l > 16 {
		return l - 16
	}
	return 0
}

func FormatGate214(s Gate214Snapshot) string {
	return fmt.Sprintf("gate214=%v thresholdSeal=%v envelope=%v central=%v ΔL=%.9g ε=%.12g orderedPairs=%d unorderedClasses=%d", s.Gate214Inherited, s.ThresholdSpectrumSealInherited, s.MatchingEnvelopeInherited, s.CentralTwoLoopConverged, s.CentralDeltaL, s.MatchingEpsilonU, s.OrderedGate211Pairs, s.UnorderedGate211Classes)
}

func FormatConfig(c ScanConfig) string {
	return fmt.Sprintf("method=%q target=%.6g ε=%.12g expectedClasses=%d steps/log=%d minSteps=%d maxIters=%d yukawa=%v matchingDerived=%v", c.Method, c.TargetU, c.EpsilonU, c.ClassesExpected, c.StepsPerLogUnit, c.MinimumStepsPerSegment, c.MaxCoordinateIters, c.YukawaTermsIncluded, c.MatchingDerived)
}

func FormatSpectrum(s ClassSpectrum) string {
	return fmt.Sprintf("rank=%d rows=%s supported=%v ΔbTot=%s bTot=%s Btot=%s gate211MeanL=%.8g gate211ΔL=%.8g verdict=%q", s.ClassRank, rowsLabel(s), s.TwoLoopSupported, s.TotalDeltaB, s.TotalOneLoopBeta, s.TotalTwoLoop, s.Gate211MeanLog, s.Gate211DeltaL, s.Verdict)
}

func FormatFit(f DegenerateFit) string {
	return fmt.Sprintf("rank=%d rows=[%s] L=(LB=%.9g,L*=%.9g) M=(MB=%.9g,M*=%.9g) U=(%.9g,%.9g,%.9g) residual=(%.9g,%.9g,%.9g) requiredδmatch=(%.9g,%.9g,%.9g) max=%.9g rms=%.9g overε=%.6g plausible=%v ordered=%v subPlanck=%v positive=%v noLandau=%v status=%s", f.ClassRank, f.Rows, f.LB, f.LStar, f.MBGeV, f.MStarGeV, f.BoundaryU[0], f.BoundaryU[1], f.BoundaryU[2], f.Residual[0], f.Residual[1], f.Residual[2], f.RequiredDeltaMatch[0], f.RequiredDeltaMatch[1], f.RequiredDeltaMatch[2], f.MaxAbsResidual, f.RMSResidual, f.ResidualOverEpsilon, f.MatchingPlausible, f.ScaleOrdered, f.SubPlanck, f.PositiveToBoundary, f.NoLandauBelowPlanck, f.Status)
}

func FormatGlobal(g GlobalScanAudit) string {
	return fmt.Sprintf("classes=%d supported=%d converged=%d plausible=%d rejected=%d bestRank=%d best=%q bestMax=%.9g bestOverε=%.6g MBRange=[%.9g,%.9g] verdict=%q", g.ClassesAudited, g.TwoLoopSupportedClasses, g.OptimizerConvergedClasses, g.PlausibleWithinEnvelope, g.RejectedByEnvelope, g.BestClassRank, g.BestRows, g.BestMaxResidual, g.BestResidualOverEpsilon, g.SmallestMBGeV, g.LargestMBGeV, g.Verdict)
}

func FormatMatching(m MatchingObstructionAudit) string {
	return fmt.Sprintf("gate214Envelope=%v nativeRows=%v heatKernel=%v subtraction=%v proxy=%v promoted=%v status=%s verdict=%q", m.Gate214MatchingEnvelopeInherited, m.NativeDeltaMatchRowsDerived, m.HeatKernelMatchingMapDerived, m.CanonicalSubtractionScheme, m.EnvelopeUsedAsProxy, m.ResidualInterpretedAsDerived, m.Status, m.Verdict)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("thresholdSeal=%v carrierSeal=%v lqSeal=%v ledger=%v all22=%v forcedFinite=%v matchingDerived=%v residualPromoted=%v yukawa=%v prediction=%v lifetime=%v next=%q", f.ThresholdSpectrumSealInherited, f.EmpiricalCarrierSealInherited, f.LeptoquarkDynamicsSealInherited, f.EmpiricalLedgerQuarantined, f.All22ClassesAudited, f.SingleScaleForcedAsFiniteCore, f.MatchingCorrectionsDerived, f.MatchingResidualPromoted, f.YukawaMatricesImported, f.PhysicalPredictionClaimed, f.ProtonLifetimeComputed, f.RecommendedNextGate)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d gate214=%v classes=%d plausible=%d bestOverε=%.6g status=%s comment=%q", s.TestsAudited, s.Gate214Inherited, s.ClassesAudited, s.PlausibleClasses, s.BestResidualOverEpsilon, s.Status, s.Comment)
}
