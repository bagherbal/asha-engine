// Package twoloopintegration implements Gate 214: sealed two-loop RG
// integration / matching-correction uncertainty envelope audit.
//
// Gate 213 introduced the ThresholdSpectrumSeal, selected the Gate-211 ranked
// witness only as a quarantined test subject, and proved that its one-loop
// scales are not precision-stable under a two-loop preflight. Gate 214 performs
// the next legal move: a numerical, piecewise two-loop integration and a
// loop-factor matching-envelope audit. The outputs are conditional
// phenomenology, not finite-core derivations.
package twoloopintegration

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/thresholdspectrumseal"
)

const (
	StatusConditionalPhenomenology = "CONDITIONAL_PHENOMENOLOGY_ON_THRESHOLD_SPECTRUM_AND_MATCHING_ENVELOPE"
	StatusFailedRoute              = "FAILED_ROUTE_TWO_LOOP_INTEGRATION"

	MatchingEnvelopeStatus = "MATCHING_UNCERTAINTY_ENVELOPE_PHENOMENOLOGICAL_PROXY"
	IntegratorStatus       = "TWO_LOOP_RK4_INTEGRATION_CONVERGED"

	mzGeV          = 91.1876
	planckLogBound = 37.8
	runTolerance   = 1.0e-9
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

func (t FloatTriple) Array() [3]float64 { return [3]float64{t.U1GUT, t.SU2L, t.SU3C} }

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

type Gate213Snapshot struct {
	Gate213Inherited               bool
	ThresholdSpectrumSealInherited bool
	MatchingCorrectionsObstructed  bool
	TwoLoopWarningInherited        bool
	OneLoopReferenceScalesOnly     bool
	SelectedRow1                   string
	SelectedRow2                   string
	OneLoopLB1                     float64
	OneLoopLB2                     float64
	OneLoopLStar                   float64
	OneLoopMB1GeV                  float64
	OneLoopMB2GeV                  float64
	OneLoopMStarGeV                float64
	TruthStatement                 string
}

type SealedSpectrum struct {
	Row1Name       string
	Row2Name       string
	Row1Rep        string
	Row2Rep        string
	Row1DeltaB     FloatTriple
	Row2DeltaB     FloatTriple
	SMBeta         FloatTriple
	SMTwoLoop      FloatMatrix3
	Row1TwoLoop    FloatMatrix3
	Row2TwoLoop    FloatMatrix3
	TotalTwoLoop   FloatMatrix3
	TargetU        float64
	MZGeV          float64
	PlanckLogBound float64
	YukawaIncluded bool
	FiniteDerived  bool
	Verdict        string
}

type IntegratorConfig struct {
	Method                  string
	Equation                string
	Coordinates             string
	StepsPerLogUnit         int
	MinimumStepsPerSegment  int
	NewtonTolerance         float64
	FiniteDifferenceScale   float64
	MaxNewtonIterations     int
	YukawaTermsIncluded     bool
	MatchingCorrectionsUsed bool
	Verdict                 string
}

type ScaleSolution struct {
	Name                string
	TargetU             [3]float64
	LB1                 float64
	LB2                 float64
	LStar               float64
	MB1GeV              float64
	MB2GeV              float64
	MStarGeV            float64
	DeltaL              float64
	ThresholdOrder      string
	BoundaryU           [3]float64
	Residual            [3]float64
	ResidualNorm        float64
	Iterations          int
	Converged           bool
	ScaleOrdered        bool
	DistinctThresholds  bool
	SubPlanck           bool
	PositiveToBoundary  bool
	NoLandauBelowPlanck bool
	PhysicalPrediction  bool
	FiniteCoreDerived   bool
	Status              string
	Verdict             string
}

type EnvelopeCase struct {
	Name         string
	Shift        [3]float64
	Solution     ScaleSolution
	Converged    bool
	ResidualNorm float64
}

type MatchingUncertaintyEnvelope struct {
	EpsilonU                    float64
	Interpretation              string
	CasesAudited                int
	ConvergedCases              int
	CentralSolutionName         string
	LB1Min, LB1Max              float64
	LB2Min, LB2Max              float64
	LStarMin, LStarMax          float64
	MB1MinGeV, MB1MaxGeV        float64
	MB2MinGeV, MB2MaxGeV        float64
	MStarMinGeV, MStarMaxGeV    float64
	MB1MinusGeV, MB1PlusGeV     float64
	MB2MinusGeV, MB2PlusGeV     float64
	MStarMinusGeV, MStarPlusGeV float64
	Status                      string
	Verdict                     string
}

type MatchingObstructionAudit struct {
	Gate213MatchingObstructionInherited bool
	NativeDeltaMatchRowsDerived         bool
	CanonicalSubtractionSchemeDerived   bool
	EnvelopeImportedAsFiniteCore        bool
	EnvelopeUsedAsPhenomenologicalProxy bool
	Status                              string
	Verdict                             string
}

type FirewallAudit struct {
	ThresholdSpectrumSealInherited  bool
	EmpiricalCarrierSealInherited   bool
	LeptoquarkDynamicsSealInherited bool
	EmpiricalLedgerQuarantined      bool
	MatchingCorrectionsDerived      bool
	MatchingEnvelopeFiniteCore      bool
	YukawaMatricesImported          bool
	PhysicalPredictionClaimed       bool
	FiniteMassPredictionClaimed     bool
	ProtonLifetimeComputed          bool
	OneLoopScalesOverwrittenAsCore  bool
	RecommendedNextGate             string
	OpenRequirements                []string
	Verdict                         string
}

type Summary struct {
	TestsAudited                int
	Gate213Inherited            bool
	CentralIntegrationConverged bool
	EnvelopeCases               int
	EnvelopeConvergedCases      int
	MatchingCorrectionsDerived  bool
	Status                      string
	Comment                     string
}

type Analysis struct {
	Gate213         Gate213Snapshot
	Gate213Analysis thresholdspectrumseal.Analysis
	Spectrum        SealedSpectrum
	Config          IntegratorConfig
	Central         ScaleSolution
	EnvelopeCases   []EnvelopeCase
	Envelope        MatchingUncertaintyEnvelope
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
		g213, err := thresholdspectrumseal.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultA, defaultErr = Build(g213)
	})
	return defaultA, defaultErr
}

func Build(g213 thresholdspectrumseal.Analysis) (Analysis, error) {
	snap := snapshotFromGate213(g213)
	if !snap.Gate213Inherited || !snap.ThresholdSpectrumSealInherited || !snap.MatchingCorrectionsObstructed || !snap.TwoLoopWarningInherited {
		return Analysis{}, fmt.Errorf("Gate 214 requires Gate 213 seal, matching obstruction, and two-loop warning")
	}
	spectrum := buildSpectrum(g213)
	config := defaultConfig()
	initial := [3]float64{g213.Subject.LB1, g213.Subject.LB2, g213.Subject.LStar}
	central, err := solveScales("central two-loop corrected fit", spectrum, config, [3]float64{1, 1, 1}, initial)
	if err != nil {
		return Analysis{}, err
	}
	cases, env := buildMatchingEnvelope(spectrum, config, central)
	matchingAudit := auditMatching(g213, env)
	fw := auditFirewall(g213, matchingAudit)
	status := StatusConditionalPhenomenology
	if !central.Converged || env.ConvergedCases != env.CasesAudited {
		status = StatusFailedRoute
	}
	summary := Summary{
		TestsAudited:                7,
		Gate213Inherited:            snap.Gate213Inherited,
		CentralIntegrationConverged: central.Converged,
		EnvelopeCases:               env.CasesAudited,
		EnvelopeConvergedCases:      env.ConvergedCases,
		MatchingCorrectionsDerived:  matchingAudit.NativeDeltaMatchRowsDerived,
		Status:                      status,
		Comment:                     "Gate 214 runs a sealed piecewise two-loop RK4 integration for the Gate-213 ThresholdSpectrumSeal test subject, solves corrected scales numerically, and wraps them in a phenomenological loop-factor matching envelope because δ_i^match remains un-derived.",
	}
	truth := buildTruth(central, env, matchingAudit)
	return Analysis{Gate213: snap, Gate213Analysis: g213, Spectrum: spectrum, Config: config, Central: central, EnvelopeCases: cases, Envelope: env, MatchingAudit: matchingAudit, Firewall: fw, Summary: summary, TruthStatement: truth}, nil
}

func snapshotFromGate213(a thresholdspectrumseal.Analysis) Gate213Snapshot {
	return Gate213Snapshot{
		Gate213Inherited:               a.Summary.Status == thresholdspectrumseal.StatusConditionalSealPreflight,
		ThresholdSpectrumSealInherited: a.Seal.DegeneracyQuarantined && a.Subject.ConditionalOnly,
		MatchingCorrectionsObstructed:  a.Matching.Status == thresholdspectrumseal.MatchingCorrectionsFailed && !a.Matching.ThresholdMatchingCoefficientsDerived,
		TwoLoopWarningInherited:        a.Stability.Status == thresholdspectrumseal.TwoLoopWarning,
		OneLoopReferenceScalesOnly:     a.Stability.OneLoopScalesValidOnlyAtOneLoop,
		SelectedRow1:                   a.Subject.Row1Rep,
		SelectedRow2:                   a.Subject.Row2Rep,
		OneLoopLB1:                     a.Subject.LB1,
		OneLoopLB2:                     a.Subject.LB2,
		OneLoopLStar:                   a.Subject.LStar,
		OneLoopMB1GeV:                  a.Subject.MB1GeV,
		OneLoopMB2GeV:                  a.Subject.MB2GeV,
		OneLoopMStarGeV:                a.Subject.MStarGeV,
		TruthStatement:                 a.TruthStatement,
	}
}

func buildSpectrum(a thresholdspectrumseal.Analysis) SealedSpectrum {
	return SealedSpectrum{
		Row1Name:       a.Subject.Row1Name,
		Row2Name:       a.Subject.Row2Name,
		Row1Rep:        a.Subject.Row1Rep,
		Row2Rep:        a.Subject.Row2Rep,
		Row1DeltaB:     ft(a.Subject.Row1DeltaB.B1.Float(), a.Subject.Row1DeltaB.B2.Float(), a.Subject.Row1DeltaB.B3.Float()),
		Row2DeltaB:     ft(a.Subject.Row2DeltaB.B1.Float(), a.Subject.Row2DeltaB.B2.Float(), a.Subject.Row2DeltaB.B3.Float()),
		SMBeta:         ft(41.0/10.0, -19.0/6.0, -7.0),
		SMTwoLoop:      matrixFromRational(a.TwoLoop.SMTwoLoopMatrixNoYukawa),
		Row1TwoLoop:    matrixFromRational(a.TwoLoop.CarrierMatrices[0].Matrix),
		Row2TwoLoop:    matrixFromRational(a.TwoLoop.CarrierMatrices[1].Matrix),
		TotalTwoLoop:   matrixFromRational(a.TwoLoop.TotalTwoLoopMatrixNoYukawa),
		TargetU:        1.0,
		MZGeV:          mzGeV,
		PlanckLogBound: planckLogBound,
		YukawaIncluded: false,
		FiniteDerived:  false,
		Verdict:        "sealed Gate-213 test subject; two-loop integration is conditional phenomenology and omits un-derived Yukawa and matching data",
	}
}

func defaultConfig() IntegratorConfig {
	return IntegratorConfig{
		Method:                  "piecewise fixed-step RK4 with damped Newton finite-difference solve",
		Equation:                "du_i/dlnμ = -b_i/(8π²) - Σ_j B_ij/u_j /(128π⁴), with u_i=1/g_i² and GUT-normalized U(1)",
		Coordinates:             "logarithmic scales L=ln(μ/M_Z); thresholds are continuous phenomenological parameters",
		StepsPerLogUnit:         30,
		MinimumStepsPerSegment:  32,
		NewtonTolerance:         runTolerance,
		FiniteDifferenceScale:   1.0e-5,
		MaxNewtonIterations:     32,
		YukawaTermsIncluded:     false,
		MatchingCorrectionsUsed: false,
		Verdict:                 "numerical QFT integrator only; no finite spectral matching scheme is imported",
	}
}

func solveScales(name string, s SealedSpectrum, cfg IntegratorConfig, target [3]float64, initial [3]float64) (ScaleSolution, error) {
	p := initial
	best := p
	bestR := residual(s, cfg, p, target)
	bestNorm := norm3(bestR)
	iters := 0
	converged := false
	for iter := 0; iter < cfg.MaxNewtonIterations; iter++ {
		iters = iter + 1
		r := residual(s, cfg, p, target)
		n := norm3(r)
		if n < bestNorm {
			best, bestR, bestNorm = p, r, n
		}
		if n < cfg.NewtonTolerance {
			best, bestR, bestNorm = p, r, n
			converged = true
			break
		}
		j := jacobian(s, cfg, p, target, r)
		delta, ok := solve3(j, [3]float64{-r[0], -r[1], -r[2]})
		if !ok {
			break
		}
		improved := false
		lambda := 1.0
		for lambda > 1.0e-6 {
			q := [3]float64{p[0] + lambda*delta[0], p[1] + lambda*delta[1], p[2] + lambda*delta[2]}
			if validParameters(q) {
				qr := residual(s, cfg, q, target)
				qn := norm3(qr)
				if qn < n {
					p = q
					improved = true
					break
				}
			}
			lambda *= 0.5
		}
		if !improved {
			break
		}
	}
	if !converged && bestNorm < cfg.NewtonTolerance*10 {
		converged = true
	}
	boundary := integrateTo(s, cfg, best, best[2])
	sol := buildSolution(name, s, cfg, target, best, boundary, bestR, bestNorm, iters, converged)
	if !sol.Converged {
		return sol, fmt.Errorf("two-loop scale solve did not converge: %s", FormatSolution(sol))
	}
	return sol, nil
}

func buildSolution(name string, s SealedSpectrum, cfg IntegratorConfig, target, p, boundary, r [3]float64, n float64, iters int, converged bool) ScaleSolution {
	mb1 := s.MZGeV * math.Exp(p[0])
	mb2 := s.MZGeV * math.Exp(p[1])
	ms := s.MZGeV * math.Exp(p[2])
	order := fmt.Sprintf("%s before %s", s.Row1Rep, s.Row2Rep)
	if p[1] < p[0] {
		order = fmt.Sprintf("%s before %s", s.Row2Rep, s.Row1Rep)
	}
	positiveBoundary, noLandau := positivityAudits(s, cfg, p)
	status := IntegratorStatus
	if !converged {
		status = StatusFailedRoute
	}
	return ScaleSolution{
		Name:                name,
		TargetU:             target,
		LB1:                 p[0],
		LB2:                 p[1],
		LStar:               p[2],
		MB1GeV:              mb1,
		MB2GeV:              mb2,
		MStarGeV:            ms,
		DeltaL:              math.Abs(p[0] - p[1]),
		ThresholdOrder:      order,
		BoundaryU:           boundary,
		Residual:            r,
		ResidualNorm:        n,
		Iterations:          iters,
		Converged:           converged,
		ScaleOrdered:        p[0] > 0 && p[1] > 0 && p[2] > math.Max(p[0], p[1]),
		DistinctThresholds:  math.Abs(p[0]-p[1]) > 1.0e-8,
		SubPlanck:           p[2] < s.PlanckLogBound,
		PositiveToBoundary:  positiveBoundary,
		NoLandauBelowPlanck: noLandau,
		PhysicalPrediction:  false,
		FiniteCoreDerived:   false,
		Status:              status,
		Verdict:             "two-loop corrected scale fit under ThresholdSpectrumSeal; not a finite-core mass prediction and not precision-complete without derived matching corrections",
	}
}

func buildMatchingEnvelope(s SealedSpectrum, cfg IntegratorConfig, central ScaleSolution) ([]EnvelopeCase, MatchingUncertaintyEnvelope) {
	eps := 1.0 / (16.0 * math.Pi * math.Pi)
	cases := []EnvelopeCase{}
	minL := [3]float64{central.LB1, central.LB2, central.LStar}
	maxL := minL
	minM := [3]float64{central.MB1GeV, central.MB2GeV, central.MStarGeV}
	maxM := minM
	converged := 0
	initial := [3]float64{central.LB1, central.LB2, central.LStar}
	for _, s1 := range []float64{-1, 1} {
		for _, s2 := range []float64{-1, 1} {
			for _, s3 := range []float64{-1, 1} {
				shift := [3]float64{s1 * eps, s2 * eps, s3 * eps}
				target := [3]float64{1 + shift[0], 1 + shift[1], 1 + shift[2]}
				sol, err := solveScales(fmt.Sprintf("matching corner shift=(%+.0f,%+.0f,%+.0f)ε", s1, s2, s3), s, cfg, target, initial)
				ok := err == nil && sol.Converged
				if ok {
					converged++
					ls := [3]float64{sol.LB1, sol.LB2, sol.LStar}
					ms := [3]float64{sol.MB1GeV, sol.MB2GeV, sol.MStarGeV}
					for i := 0; i < 3; i++ {
						if ls[i] < minL[i] {
							minL[i] = ls[i]
						}
						if ls[i] > maxL[i] {
							maxL[i] = ls[i]
						}
						if ms[i] < minM[i] {
							minM[i] = ms[i]
						}
						if ms[i] > maxM[i] {
							maxM[i] = ms[i]
						}
					}
				}
				cases = append(cases, EnvelopeCase{Name: sol.Name, Shift: shift, Solution: sol, Converged: ok, ResidualNorm: sol.ResidualNorm})
			}
		}
	}
	env := MatchingUncertaintyEnvelope{
		EpsilonU:            eps,
		Interpretation:      "deterministic ±1/(16π²) u-space corner scan standing in for un-derived threshold matching corrections; this is a phenomenological uncertainty proxy, not a finite theorem",
		CasesAudited:        len(cases),
		ConvergedCases:      converged,
		CentralSolutionName: central.Name,
		LB1Min:              minL[0], LB1Max: maxL[0],
		LB2Min: minL[1], LB2Max: maxL[1],
		LStarMin: minL[2], LStarMax: maxL[2],
		MB1MinGeV: minM[0], MB1MaxGeV: maxM[0],
		MB2MinGeV: minM[1], MB2MaxGeV: maxM[1],
		MStarMinGeV: minM[2], MStarMaxGeV: maxM[2],
		MB1MinusGeV: central.MB1GeV - minM[0], MB1PlusGeV: maxM[0] - central.MB1GeV,
		MB2MinusGeV: central.MB2GeV - minM[1], MB2PlusGeV: maxM[1] - central.MB2GeV,
		MStarMinusGeV: central.MStarGeV - minM[2], MStarPlusGeV: maxM[2] - central.MStarGeV,
		Status:  MatchingEnvelopeStatus,
		Verdict: "matching-envelope scan converged for all deterministic loop-factor corners; scale bands quantify theory uncertainty from the still-sealed δ_i^match rows",
	}
	return cases, env
}

func auditMatching(g213 thresholdspectrumseal.Analysis, env MatchingUncertaintyEnvelope) MatchingObstructionAudit {
	return MatchingObstructionAudit{
		Gate213MatchingObstructionInherited: g213.Matching.Status == thresholdspectrumseal.MatchingCorrectionsFailed,
		NativeDeltaMatchRowsDerived:         false,
		CanonicalSubtractionSchemeDerived:   false,
		EnvelopeImportedAsFiniteCore:        false,
		EnvelopeUsedAsPhenomenologicalProxy: env.CasesAudited > 0,
		Status:                              thresholdspectrumseal.MatchingCorrectionsFailed,
		Verdict:                             "Gate 214 still does not derive δ_i^match; the envelope is an explicit phenomenological loop-factor proxy around the sealed two-loop fit",
	}
}

func auditFirewall(g213 thresholdspectrumseal.Analysis, m MatchingObstructionAudit) FirewallAudit {
	return FirewallAudit{
		ThresholdSpectrumSealInherited:  g213.Seal.DegeneracyQuarantined,
		EmpiricalCarrierSealInherited:   g213.Firewall.EmpiricalCarrierSealInherited,
		LeptoquarkDynamicsSealInherited: g213.Firewall.LeptoquarkDynamicsSealInherited,
		EmpiricalLedgerQuarantined:      g213.Firewall.EmpiricalLedgerQuarantined,
		MatchingCorrectionsDerived:      m.NativeDeltaMatchRowsDerived,
		MatchingEnvelopeFiniteCore:      m.EnvelopeImportedAsFiniteCore,
		YukawaMatricesImported:          false,
		PhysicalPredictionClaimed:       false,
		FiniteMassPredictionClaimed:     false,
		ProtonLifetimeComputed:          false,
		OneLoopScalesOverwrittenAsCore:  false,
		RecommendedNextGate:             "Gate 215 — threshold-spectrum envelope ranking / Yukawa and matching-scheme sensitivity audit",
		OpenRequirements: []string{
			"derive a finite spectral matching map before replacing the loop-factor envelope with exact δ_i^match",
			"include SM Yukawa and scalar-sector terms only under a separate seal or finite derivation",
			"test all 22 unordered Gate-211 spectra under the same sealed two-loop envelope before claiming uniqueness",
			"preserve ThresholdSpectrumSeal, EmpiricalCarrierSeal, and LeptoquarkDynamicsSeal",
		},
		Verdict: "firewall preserved: corrected scales and error bars are conditional numerical phenomenology, not finite algebraic predictions",
	}
}

func residual(s SealedSpectrum, cfg IntegratorConfig, p, target [3]float64) [3]float64 {
	if !validParameters(p) {
		return [3]float64{1e6, 1e6, 1e6}
	}
	u := integrateTo(s, cfg, p, p[2])
	return [3]float64{u[0] - target[0], u[1] - target[1], u[2] - target[2]}
}

func integrateTo(s SealedSpectrum, cfg IntegratorConfig, p [3]float64, end float64) [3]float64 {
	u := zPoleU()
	if end <= 0 {
		return u
	}
	points := []float64{0, end}
	if p[0] > 0 && p[0] < end {
		points = append(points, p[0])
	}
	if p[1] > 0 && p[1] < end {
		points = append(points, p[1])
	}
	sort.Float64s(points)
	uniq := []float64{points[0]}
	for _, x := range points[1:] {
		if math.Abs(x-uniq[len(uniq)-1]) > 1e-12 {
			uniq = append(uniq, x)
		}
	}
	for k := 0; k < len(uniq)-1; k++ {
		start, stop := uniq[k], uniq[k+1]
		if stop <= start {
			continue
		}
		mid := 0.5 * (start + stop)
		b, bm := activeCoefficients(s, p, mid)
		u = rk4Segment(u, b, bm, stop-start, cfg)
	}
	return u
}

func activeCoefficients(s SealedSpectrum, p [3]float64, l float64) (FloatTriple, FloatMatrix3) {
	b := s.SMBeta
	bm := s.SMTwoLoop
	if l >= p[0] {
		b = addTriple(b, s.Row1DeltaB)
		bm = bm.Add(s.Row1TwoLoop)
	}
	if l >= p[1] {
		b = addTriple(b, s.Row2DeltaB)
		bm = bm.Add(s.Row2TwoLoop)
	}
	return b, bm
}

func rk4Segment(u [3]float64, b FloatTriple, bm FloatMatrix3, length float64, cfg IntegratorConfig) [3]float64 {
	steps := int(math.Ceil(length * float64(cfg.StepsPerLogUnit)))
	if steps < cfg.MinimumStepsPerSegment {
		steps = cfg.MinimumStepsPerSegment
	}
	h := length / float64(steps)
	for i := 0; i < steps; i++ {
		k1 := derivativeU(u, b, bm)
		k2 := derivativeU(addScaled(u, k1, 0.5*h), b, bm)
		k3 := derivativeU(addScaled(u, k2, 0.5*h), b, bm)
		k4 := derivativeU(addScaled(u, k3, h), b, bm)
		for j := 0; j < 3; j++ {
			u[j] += h * (k1[j] + 2*k2[j] + 2*k3[j] + k4[j]) / 6.0
		}
	}
	return u
}

func derivativeU(u [3]float64, b FloatTriple, bm FloatMatrix3) [3]float64 {
	var out [3]float64
	for i := 0; i < 3; i++ {
		sum := 0.0
		for j := 0; j < 3; j++ {
			if u[j] <= 0 || math.IsNaN(u[j]) || math.IsInf(u[j], 0) {
				return [3]float64{math.NaN(), math.NaN(), math.NaN()}
			}
			sum += bm.M[i][j] / u[j]
		}
		out[i] = -b.At(i)/(8.0*math.Pi*math.Pi) - sum/(128.0*math.Pow(math.Pi, 4))
	}
	return out
}

func jacobian(s SealedSpectrum, cfg IntegratorConfig, p, target, baseR [3]float64) [3][3]float64 {
	var j [3][3]float64
	for col := 0; col < 3; col++ {
		eps := cfg.FiniteDifferenceScale * math.Max(1, math.Abs(p[col]))
		q := p
		q[col] += eps
		qr := residual(s, cfg, q, target)
		for row := 0; row < 3; row++ {
			j[row][col] = (qr[row] - baseR[row]) / eps
		}
	}
	return j
}

func solve3(a [3][3]float64, b [3]float64) ([3]float64, bool) {
	m := [3][4]float64{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			m[i][j] = a[i][j]
		}
		m[i][3] = b[i]
	}
	for col := 0; col < 3; col++ {
		pivot := col
		for r := col + 1; r < 3; r++ {
			if math.Abs(m[r][col]) > math.Abs(m[pivot][col]) {
				pivot = r
			}
		}
		if math.Abs(m[pivot][col]) < 1e-14 {
			return [3]float64{}, false
		}
		if pivot != col {
			m[pivot], m[col] = m[col], m[pivot]
		}
		pv := m[col][col]
		for c := col; c < 4; c++ {
			m[col][c] /= pv
		}
		for r := 0; r < 3; r++ {
			if r == col {
				continue
			}
			fac := m[r][col]
			for c := col; c < 4; c++ {
				m[r][c] -= fac * m[col][c]
			}
		}
	}
	return [3]float64{m[0][3], m[1][3], m[2][3]}, true
}

func positivityAudits(s SealedSpectrum, cfg IntegratorConfig, p [3]float64) (bool, bool) {
	positiveBoundary := true
	noLandau := true
	checkpoints := []float64{0, p[0], p[1], p[2], s.PlanckLogBound}
	sort.Float64s(checkpoints)
	for _, l := range checkpoints {
		if l < 0 {
			continue
		}
		u := integrateTo(s, cfg, p, l)
		for _, x := range u {
			if x <= 0 || math.IsNaN(x) || math.IsInf(x, 0) {
				if l <= p[2] {
					positiveBoundary = false
				}
				noLandau = false
			}
		}
	}
	return positiveBoundary, noLandau
}

func validParameters(p [3]float64) bool {
	return p[0] > 0 && p[1] > 0 && p[2] > math.Max(p[0], p[1]) && p[2] < planckLogBound && math.Abs(p[0]-p[1]) > 1e-8
}

func zPoleU() [3]float64 {
	alphaInv := 127.955
	sin2 := 0.23122
	alphaS := 0.1179
	return [3]float64{
		((3.0 / 5.0) * (1.0 - sin2) * alphaInv) / (4.0 * math.Pi),
		(sin2 * alphaInv) / (4.0 * math.Pi),
		(1.0 / alphaS) / (4.0 * math.Pi),
	}
}

func matrixFromRational(m thresholdspectrumseal.RationalMatrix3) FloatMatrix3 {
	var out FloatMatrix3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			out.M[i][j] = m.FloatAt(i, j)
		}
	}
	return out
}

func ft(a, b, c float64) FloatTriple { return FloatTriple{a, b, c} }

func addTriple(a, b FloatTriple) FloatTriple {
	return ft(a.U1GUT+b.U1GUT, a.SU2L+b.SU2L, a.SU3C+b.SU3C)
}

func addScaled(a, b [3]float64, h float64) [3]float64 {
	return [3]float64{a[0] + h*b[0], a[1] + h*b[1], a[2] + h*b[2]}
}

func norm3(x [3]float64) float64 { return math.Sqrt(x[0]*x[0] + x[1]*x[1] + x[2]*x[2]) }

func buildTruth(c ScaleSolution, e MatchingUncertaintyEnvelope, m MatchingObstructionAudit) string {
	return fmt.Sprintf("Gate 214 conditionally solves the sealed two-loop no-Yukawa RG system for the Gate-213 spectrum. The corrected central scales are M_B(row1)=%.9g GeV, M_B(row2)=%.9g GeV, M*=%.9g GeV, with threshold order %s and residual %.3g. Because %s remains active, the ±1/(16π²) matching-envelope proxy expands the scale ranges to M_B1∈[%.9g,%.9g], M_B2∈[%.9g,%.9g], M*∈[%.9g,%.9g] GeV. These are conditional phenomenological fits, not finite-core predictions.", c.MB1GeV, c.MB2GeV, c.MStarGeV, c.ThresholdOrder, c.ResidualNorm, m.Status, e.MB1MinGeV, e.MB1MaxGeV, e.MB2MinGeV, e.MB2MaxGeV, e.MStarMinGeV, e.MStarMaxGeV)
}

func FormatGate213(s Gate213Snapshot) string {
	return fmt.Sprintf("gate213=%t spectrumSeal=%t matchingObstructed=%t twoLoopWarning=%t oneLoopOnly=%t rows=[%s,%s] oneLoopL=(%.9g,%.9g,%.9g)", s.Gate213Inherited, s.ThresholdSpectrumSealInherited, s.MatchingCorrectionsObstructed, s.TwoLoopWarningInherited, s.OneLoopReferenceScalesOnly, s.SelectedRow1, s.SelectedRow2, s.OneLoopLB1, s.OneLoopLB2, s.OneLoopLStar)
}

func FormatSpectrum(s SealedSpectrum) string {
	return fmt.Sprintf("rows=[%s %s Δb=%s; %s %s Δb=%s] Btotal=%s targetU=%.9g yukawa=%t finite=%t", s.Row1Name, s.Row1Rep, s.Row1DeltaB.String(), s.Row2Name, s.Row2Rep, s.Row2DeltaB.String(), s.TotalTwoLoop.String(), s.TargetU, s.YukawaIncluded, s.FiniteDerived)
}

func FormatConfig(c IntegratorConfig) string {
	return fmt.Sprintf("method=%s coords=%s stepsPerLog=%d minSteps=%d tol=%.3g yukawa=%t matchingUsed=%t", c.Method, c.Coordinates, c.StepsPerLogUnit, c.MinimumStepsPerSegment, c.NewtonTolerance, c.YukawaTermsIncluded, c.MatchingCorrectionsUsed)
}

func FormatSolution(s ScaleSolution) string {
	return fmt.Sprintf("%s L=(%.9g,%.9g,%.9g) M=(%.9g,%.9g,%.9g) ΔL=%.9g order=%s u=(%.12g,%.12g,%.12g) r=(%.3g,%.3g,%.3g) norm=%.3g it=%d conv=%t ordered=%t distinct=%t subPlanck=%t positive=%t noLandau=%t status=%s", s.Name, s.LB1, s.LB2, s.LStar, s.MB1GeV, s.MB2GeV, s.MStarGeV, s.DeltaL, s.ThresholdOrder, s.BoundaryU[0], s.BoundaryU[1], s.BoundaryU[2], s.Residual[0], s.Residual[1], s.Residual[2], s.ResidualNorm, s.Iterations, s.Converged, s.ScaleOrdered, s.DistinctThresholds, s.SubPlanck, s.PositiveToBoundary, s.NoLandauBelowPlanck, s.Status)
}

func FormatEnvelope(e MatchingUncertaintyEnvelope) string {
	return fmt.Sprintf("eps=%.12g cases=%d converged=%d L1=[%.9g,%.9g] L2=[%.9g,%.9g] L*=[%.9g,%.9g] MB1=[%.9g,%.9g] MB2=[%.9g,%.9g] M*=[%.9g,%.9g] status=%s", e.EpsilonU, e.CasesAudited, e.ConvergedCases, e.LB1Min, e.LB1Max, e.LB2Min, e.LB2Max, e.LStarMin, e.LStarMax, e.MB1MinGeV, e.MB1MaxGeV, e.MB2MinGeV, e.MB2MaxGeV, e.MStarMinGeV, e.MStarMaxGeV, e.Status)
}

func FormatMatching(a MatchingObstructionAudit) string {
	return fmt.Sprintf("gate213Obstruction=%t nativeRows=%t scheme=%t envelopeCore=%t proxy=%t status=%s", a.Gate213MatchingObstructionInherited, a.NativeDeltaMatchRowsDerived, a.CanonicalSubtractionSchemeDerived, a.EnvelopeImportedAsFiniteCore, a.EnvelopeUsedAsPhenomenologicalProxy, a.Status)
}

func FormatFirewall(f FirewallAudit) string {
	return fmt.Sprintf("spectrumSeal=%t carrierSeal=%t lqSeal=%t ledger=%t matching=%t envelopeCore=%t yukawa=%t physicalPrediction=%t finiteMass=%t lifetime=%t oneLoopCoreOverwrite=%t next=%s", f.ThresholdSpectrumSealInherited, f.EmpiricalCarrierSealInherited, f.LeptoquarkDynamicsSealInherited, f.EmpiricalLedgerQuarantined, f.MatchingCorrectionsDerived, f.MatchingEnvelopeFiniteCore, f.YukawaMatricesImported, f.PhysicalPredictionClaimed, f.FiniteMassPredictionClaimed, f.ProtonLifetimeComputed, f.OneLoopScalesOverwrittenAsCore, f.RecommendedNextGate)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d gate213=%t central=%t envelope=%d/%d matchingDerived=%t status=%s", s.TestsAudited, s.Gate213Inherited, s.CentralIntegrationConverged, s.EnvelopeConvergedCases, s.EnvelopeCases, s.MatchingCorrectionsDerived, s.Status)
}
