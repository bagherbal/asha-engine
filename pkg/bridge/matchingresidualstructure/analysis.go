// Package matchingresidualstructure implements Gate 216: matching-residual
// structure audit / spectral heat-kernel coefficient search.
//
// Gate 215 reduced the single-scale two-loop threshold problem to one viable
// spectrum and one required matching residual vector. Gate 216 asks whether the
// existing finite spectral data already supplies that vector as a canonical
// heat-kernel/matching coefficient. The answer is deliberately strict: positive
// spectral anchors and signed eta-trace degrees exist, but no finite spectral
// triple, gauge-curvature projection, cutoff coefficient, or subtraction scheme
// maps them to threshold matching rows. Numerical near-matches and sign-only
// resonances are recorded as diagnostics, not promoted to theorems.
package matchingresidualstructure

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/singlescalematchingaudit"
	"github.com/bagherbal/asha-engine/pkg/bridge/threshold"
)

const (
	StatusFailedRoute        = "FAILED_ROUTE_SPECTRAL_MATCHING_RESIDUAL_DERIVATION"
	StatusConditionalSupport = "CONDITIONAL_SUPPORT_SPECTRAL_MATCHING_RESIDUAL"

	SignResonanceOnly = "SIGN_RESONANCE_ONLY_NO_HEAT_KERNEL_MAP"
	NoStructuralMatch = "NO_STRUCTURAL_MATCH"
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
func (t FloatTriple) String() string {
	return fmt.Sprintf("(%.12g,%.12g,%.12g)", t.U1GUT, t.SU2L, t.SU3C)
}

func (t FloatTriple) MaxAbs() float64 {
	return math.Max(math.Abs(t.U1GUT), math.Max(math.Abs(t.SU2L), math.Abs(t.SU3C)))
}
func (t FloatTriple) Norm() float64 {
	return math.Sqrt(t.U1GUT*t.U1GUT + t.SU2L*t.SU2L + t.SU3C*t.SU3C)
}
func (t FloatTriple) Scale(c float64) FloatTriple {
	return FloatTriple{c * t.U1GUT, c * t.SU2L, c * t.SU3C}
}
func (t FloatTriple) Sub(o FloatTriple) FloatTriple {
	return FloatTriple{t.U1GUT - o.U1GUT, t.SU2L - o.SU2L, t.SU3C - o.SU3C}
}
func (t FloatTriple) NormalizedMax() FloatTriple {
	m := t.MaxAbs()
	if m == 0 || math.IsNaN(m) || math.IsInf(m, 0) {
		return FloatTriple{}
	}
	return t.Scale(1 / m)
}

type Gate215Snapshot struct {
	Gate215Inherited               bool
	ThresholdSpectrumSealInherited bool
	MatchingEnvelopeInherited      bool
	SingleScaleCandidateUnique     bool
	BestRows                       string
	MBGeV                          float64
	MStarGeV                       float64
	EpsilonU                       float64
	RequiredDeltaMatch             FloatTriple
	RequiredSignPattern            string
	RequiredMaxAbs                 float64
	RequiredOverEnvelope           float64
	PlausibleClasses               int
	ClassesAudited                 int
	FiniteMatchingDerived          bool
	TruthStatement                 string
}

type SpectralDataAudit struct {
	BGapValue                         float64
	ContactPartialOverlap             []float64
	ContactPartialModeCount           int
	ContactMean                       float64
	ContactMin                        float64
	ContactMax                        float64
	ScalarActiveModeCount             int
	ContactZetaValues                 map[string]float64
	TauEtaDegrees                     FloatTriple
	TauEtaOrientationFlipped          FloatTriple
	FiniteFundamentalClassAvailable   bool
	ContactZetaLedgerAvailable        bool
	BGapAvailable                     bool
	AllSpectralScalarsDimensionless   bool
	GaugeSectorProjectionDerived      bool
	HeatKernelCoefficientMapDerived   bool
	CanonicalCutoffFunctionDerived    bool
	CanonicalSubtractionSchemeDerived bool
	FiniteMatchingRowsDerived         bool
	Verdict                           string
}

type ResidualStructureAudit struct {
	TargetVector                    FloatTriple
	TargetSignPattern               string
	TargetNormalized                FloatTriple
	EqualMagnitudeTolerance         float64
	AlternatingPattern              bool
	NearlyEqualMagnitudes           bool
	TraceTargetInterpretedAsDerived bool
	Verdict                         string
}

type TraceCandidate struct {
	Name                      string
	Source                    string
	Formula                   string
	ScalarValue               float64
	Vector                    FloatTriple
	Normalized                FloatTriple
	SignPattern               string
	SignsMatch                bool
	RelativeMagnitudesMatch   bool
	CanonicalMagnitudeMatches bool
	BestFittedScale           float64
	BestFittedMaxResidual     float64
	RequiresFittedCoefficient bool
	RequiresGaugeProjection   bool
	RequiresSpectralTriple    bool
	RequiresCutoffFunction    bool
	RequiresSubtractionScheme bool
	PromotedToMatchingRow     bool
	Verdict                   string
}

type HeatKernelMappingAudit struct {
	A2A4LanguageAudited               bool
	FiniteDiracOperatorDerived        bool
	SpectralTripleComplete            bool
	GaugeCurvatureProjectionRows      int
	GaugeKineticTraceMapDerived       bool
	CanonicalCutoffMomentsDerived     bool
	CanonicalMSbarSchemeImported      bool
	ThresholdSubtractionSchemeDerived bool
	DeltaMatchRowsDerived             int
	SignOnlyResonances                int
	FullStructuralMatches             int
	Verdict                           string
}

type CoefficientSearchAudit struct {
	CanonicalScalarsAudited        int
	CanonicalLoopScaledCandidates  int
	ExactMagnitudeMatches          int
	ClosestCandidateName           string
	ClosestCandidateFormula        string
	ClosestCandidateMagnitude      float64
	ClosestCandidateRatioToTarget  float64
	ClosestCandidateAbsError       float64
	ClosestCandidateAccepted       bool
	FittedCoefficientNeeded        bool
	ArbitraryNormalizationInserted bool
	Verdict                        string
}

type FirewallAudit struct {
	Gate215Inherited                bool
	ThresholdSpectrumSealInherited  bool
	EmpiricalCarrierSealInherited   bool
	LeptoquarkDynamicsSealInherited bool
	EmpiricalLedgerQuarantined      bool
	MatchingResidualPromoted        bool
	MatchingCorrectionsDerived      bool
	SpectralCoefficientTuned        bool
	HeatKernelMapImported           bool
	MSbarSchemeImported             bool
	PhysicalMassPredictionClaimed   bool
	PhysicalUnificationClaimed      bool
	ContactModesPromotedToParticles bool
	BGapPromotedToMass              bool
	ProtonLifetimeComputed          bool
	RecommendedNextGate             string
	OpenRequirements                []string
	Verdict                         string
}

type Summary struct {
	TestsAudited              int
	Gate215Inherited          bool
	SpectralDataAudited       bool
	TraceCandidatesAudited    int
	SignOnlyResonances        int
	StructuralMatches         int
	CanonicalMagnitudeMatches int
	Status                    string
	Comment                   string
}

type Analysis struct {
	Gate215           Gate215Snapshot
	Gate215Analysis   singlescalematchingaudit.Analysis
	SpectralData      SpectralDataAudit
	ResidualTarget    ResidualStructureAudit
	TraceCandidates   []TraceCandidate
	HeatKernelMap     HeatKernelMappingAudit
	CoefficientSearch CoefficientSearchAudit
	Firewall          FirewallAudit
	Summary           Summary
	TruthStatement    string
}

var (
	defaultOnce sync.Once
	defaultA    Analysis
	defaultErr  error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		g215, err := singlescalematchingaudit.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build Gate 215 input: %w", err)
			return
		}
		th, err := threshold.BuildDefault()
		if err != nil {
			defaultErr = fmt.Errorf("build finite threshold/spectral input: %w", err)
			return
		}
		defaultA, defaultErr = Build(g215, th)
	})
	return defaultA, defaultErr
}

func Build(g215 singlescalematchingaudit.Analysis, th threshold.Analysis) (Analysis, error) {
	snap, err := snapshotFromGate215(g215)
	if err != nil {
		return Analysis{}, err
	}
	if !snap.Gate215Inherited || !snap.ThresholdSpectrumSealInherited || !snap.MatchingEnvelopeInherited || !snap.SingleScaleCandidateUnique {
		return Analysis{}, fmt.Errorf("Gate 216 requires Gate 215 unique plausible single-scale target under inherited seals")
	}
	spec := auditSpectralData(th)
	target := auditResidualTarget(snap)
	cands := traceCandidates(snap, spec)
	hk := auditHeatKernelMap(cands, spec)
	coef := auditCoefficientSearch(snap, spec, cands)
	fw := auditFirewall(snap, spec, hk, coef)
	status := StatusFailedRoute
	if hk.FullStructuralMatches > 0 && coef.ExactMagnitudeMatches > 0 && !fw.MatchingCorrectionsDerived {
		status = StatusConditionalSupport
	}
	summary := Summary{
		TestsAudited:              6,
		Gate215Inherited:          snap.Gate215Inherited,
		SpectralDataAudited:       spec.BGapAvailable && spec.ContactPartialModeCount == 7 && spec.FiniteFundamentalClassAvailable,
		TraceCandidatesAudited:    len(cands),
		SignOnlyResonances:        hk.SignOnlyResonances,
		StructuralMatches:         hk.FullStructuralMatches,
		CanonicalMagnitudeMatches: coef.ExactMagnitudeMatches,
		Status:                    status,
		Comment:                   "Gate 216 tests whether existing finite spectral traces canonically produce the Gate-215 required alternating matching residual without fitted coefficients.",
	}
	truth := buildTruth(summary, snap, hk, coef)
	return Analysis{Gate215: snap, Gate215Analysis: g215, SpectralData: spec, ResidualTarget: target, TraceCandidates: cands, HeatKernelMap: hk, CoefficientSearch: coef, Firewall: fw, Summary: summary, TruthStatement: truth}, nil
}

func snapshotFromGate215(a singlescalematchingaudit.Analysis) (Gate215Snapshot, error) {
	if len(a.Fits) == 0 {
		return Gate215Snapshot{}, fmt.Errorf("Gate 215 has no degenerate fits")
	}
	best := a.Fits[0]
	dm := FloatTriple{best.RequiredDeltaMatch[0], best.RequiredDeltaMatch[1], best.RequiredDeltaMatch[2]}
	return Gate215Snapshot{
		Gate215Inherited:               a.Summary.Status == singlescalematchingaudit.StatusConditionalPhenomenology,
		ThresholdSpectrumSealInherited: a.Firewall.ThresholdSpectrumSealInherited,
		MatchingEnvelopeInherited:      a.Gate214.MatchingEnvelopeInherited && a.Config.EpsilonU > 0,
		SingleScaleCandidateUnique:     a.GlobalScan.PlausibleWithinEnvelope == 1 && best.MatchingPlausible,
		BestRows:                       best.Rows,
		MBGeV:                          best.MBGeV,
		MStarGeV:                       best.MStarGeV,
		EpsilonU:                       a.Config.EpsilonU,
		RequiredDeltaMatch:             dm,
		RequiredSignPattern:            signPattern(dm),
		RequiredMaxAbs:                 dm.MaxAbs(),
		RequiredOverEnvelope:           best.ResidualOverEpsilon,
		PlausibleClasses:               a.GlobalScan.PlausibleWithinEnvelope,
		ClassesAudited:                 a.GlobalScan.ClassesAudited,
		FiniteMatchingDerived:          a.MatchingAudit.NativeDeltaMatchRowsDerived || a.MatchingAudit.HeatKernelMatchingMapDerived || a.MatchingAudit.CanonicalSubtractionScheme,
		TruthStatement:                 a.TruthStatement,
	}, nil
}

func auditSpectralData(th threshold.Analysis) SpectralDataAudit {
	partial := append([]float64(nil), th.ContactPartialOverlap...)
	sort.Float64s(partial)
	mean, minv, maxv := meanMinMax(partial)
	zeta := map[string]float64{
		"zeta(0)":              7,
		"zeta(1)":              float64(7993) / float64(542),
		"zeta(2)":              float64(10529233) / float64(293764),
		"zeta(3)":              float64(15529024549) / float64(159220088),
		"zeta(4)":              float64(24783201328945) / float64(86297287696),
		"zeta(1)/7":            (float64(7993) / float64(542)) / 7,
		"inverse zeta(1)":      float64(542) / float64(7993),
		"mean contact overlap": mean,
	}
	return SpectralDataAudit{
		BGapValue:                         th.BGap,
		ContactPartialOverlap:             partial,
		ContactPartialModeCount:           len(partial),
		ContactMean:                       mean,
		ContactMin:                        minv,
		ContactMax:                        maxv,
		ScalarActiveModeCount:             len(th.ScalarActiveSpectrum),
		ContactZetaValues:                 zeta,
		TauEtaDegrees:                     FloatTriple{2, -2, 1},
		TauEtaOrientationFlipped:          FloatTriple{-2, 2, -1},
		FiniteFundamentalClassAvailable:   true,
		ContactZetaLedgerAvailable:        true,
		BGapAvailable:                     th.BGap > 0,
		AllSpectralScalarsDimensionless:   th.DimensionlessSpectralAnchorsAvailable && !th.PhysicalMassUnitDerived,
		GaugeSectorProjectionDerived:      false,
		HeatKernelCoefficientMapDerived:   false,
		CanonicalCutoffFunctionDerived:    false,
		CanonicalSubtractionSchemeDerived: false,
		FiniteMatchingRowsDerived:         false,
		Verdict:                           "finite spectral scalars and eta-signed degrees exist, but no gauge-sector heat-kernel projection or threshold matching row is derived",
	}
}

func auditResidualTarget(s Gate215Snapshot) ResidualStructureAudit {
	n := s.RequiredDeltaMatch.NormalizedMax()
	nearlyEqual := maxAbs3(FloatTriple{math.Abs(n.U1GUT) - 1, math.Abs(n.SU2L) - 1, math.Abs(n.SU3C) - 1}) < 0.003
	alt := s.RequiredSignPattern == "-+-"
	return ResidualStructureAudit{
		TargetVector:                    s.RequiredDeltaMatch,
		TargetSignPattern:               s.RequiredSignPattern,
		TargetNormalized:                n,
		EqualMagnitudeTolerance:         0.003,
		AlternatingPattern:              alt,
		NearlyEqualMagnitudes:           nearlyEqual,
		TraceTargetInterpretedAsDerived: false,
		Verdict:                         "Gate-215 residual is a precise target vector for matching theory; it is not itself a derived spectral trace",
	}
}

func traceCandidates(s Gate215Snapshot, spec SpectralDataAudit) []TraceCandidate {
	target := s.RequiredDeltaMatch
	eps := s.EpsilonU
	var c []TraceCandidate
	add := func(name, source, formula string, scalar float64, v FloatTriple, gauge, triple, cutoff, scheme bool) {
		cand := buildCandidate(name, source, formula, scalar, v, target, gauge, triple, cutoff, scheme)
		c = append(c, cand)
	}
	add("B-sector gap scalar", "B-sector", "gap_B · (1,1,1)", spec.BGapValue, FloatTriple{spec.BGapValue, spec.BGapValue, spec.BGapValue}, true, true, true, true)
	add("B-sector gap loop-scaled scalar", "B-sector", "gap_B/(16π²) · (1,1,1)", spec.BGapValue*eps, FloatTriple{spec.BGapValue * eps, spec.BGapValue * eps, spec.BGapValue * eps}, true, true, true, true)
	add("contact partial mean loop-scaled", "contact partial-overlap", "mean(lambda_partial)/(16π²) · (1,1,1)", spec.ContactMean*eps, FloatTriple{spec.ContactMean * eps, spec.ContactMean * eps, spec.ContactMean * eps}, true, true, true, true)
	add("contact zeta inverse trace loop-scaled", "contact zeta", "zeta(1)^-1/(16π²) · (1,1,1)", spec.ContactZetaValues["inverse zeta(1)"]*eps, uniform(spec.ContactZetaValues["inverse zeta(1)"]*eps), true, true, true, true)
	add("tau_eta native signed degrees", "scalar fundamental class", "(tau_eta(Q²),tau_eta(Z²),tau_eta(QZ))=(2,-2,1)", 1, spec.TauEtaDegrees, true, true, true, true)
	add("tau_eta orientation-flipped signed degrees", "scalar fundamental class", "-(2,-2,1)=(-2,2,-1)", 1, spec.TauEtaOrientationFlipped, true, true, true, true)
	add("tau_eta orientation-flipped loop-scaled degrees", "scalar fundamental class", "-(2,-2,1)/(16π²)", eps, spec.TauEtaOrientationFlipped.Scale(eps), true, true, true, true)
	return c
}

func buildCandidate(name, source, formula string, scalar float64, v, target FloatTriple, requiresGauge, requiresTriple, requiresCutoff, requiresScheme bool) TraceCandidate {
	bestScale, best := bestFittedScale(v, target)
	signs := signPattern(v)
	signsMatch := signs == signPattern(target)
	relMatch := vectorClose(v.NormalizedMax(), target.NormalizedMax(), 0.01)
	magMatch := math.Abs(v.MaxAbs()-target.MaxAbs()) <= 1e-9
	verdict := "incompatible with the Gate-215 residual vector"
	if signsMatch && !relMatch {
		verdict = "sign resonance only; relative magnitudes fail and no heat-kernel gauge projection selects this vector"
	} else if signsMatch && relMatch && !magMatch {
		verdict = "structural direction resonance only; magnitude would require a fitted coefficient or missing cutoff moment"
	} else if signsMatch && relMatch && magMatch && !(requiresGauge || requiresTriple || requiresCutoff || requiresScheme) {
		verdict = "canonical matching-row candidate"
	}
	return TraceCandidate{
		Name:                      name,
		Source:                    source,
		Formula:                   formula,
		ScalarValue:               scalar,
		Vector:                    v,
		Normalized:                v.NormalizedMax(),
		SignPattern:               signs,
		SignsMatch:                signsMatch,
		RelativeMagnitudesMatch:   relMatch,
		CanonicalMagnitudeMatches: magMatch,
		BestFittedScale:           bestScale,
		BestFittedMaxResidual:     best,
		RequiresFittedCoefficient: !magMatch,
		RequiresGaugeProjection:   requiresGauge,
		RequiresSpectralTriple:    requiresTriple,
		RequiresCutoffFunction:    requiresCutoff,
		RequiresSubtractionScheme: requiresScheme,
		PromotedToMatchingRow:     false,
		Verdict:                   verdict,
	}
}

func auditHeatKernelMap(c []TraceCandidate, spec SpectralDataAudit) HeatKernelMappingAudit {
	signOnly := 0
	full := 0
	for _, x := range c {
		if x.SignsMatch && !x.RelativeMagnitudesMatch {
			signOnly++
		}
		if x.SignsMatch && x.RelativeMagnitudesMatch && x.CanonicalMagnitudeMatches && !x.RequiresGaugeProjection && !x.RequiresSpectralTriple && !x.RequiresCutoffFunction && !x.RequiresSubtractionScheme {
			full++
		}
	}
	return HeatKernelMappingAudit{
		A2A4LanguageAudited:               true,
		FiniteDiracOperatorDerived:        false,
		SpectralTripleComplete:            false,
		GaugeCurvatureProjectionRows:      0,
		GaugeKineticTraceMapDerived:       spec.GaugeSectorProjectionDerived,
		CanonicalCutoffMomentsDerived:     spec.CanonicalCutoffFunctionDerived,
		CanonicalMSbarSchemeImported:      false,
		ThresholdSubtractionSchemeDerived: spec.CanonicalSubtractionSchemeDerived,
		DeltaMatchRowsDerived:             0,
		SignOnlyResonances:                signOnly,
		FullStructuralMatches:             full,
		Verdict:                           "heat-kernel coefficient language is only a preflight: the engine has no finite Dirac/spectral triple/cutoff/subtraction map that converts finite traces into δ_i^match rows",
	}
}

func auditCoefficientSearch(s Gate215Snapshot, spec SpectralDataAudit, c []TraceCandidate) CoefficientSearchAudit {
	target := s.RequiredMaxAbs
	type scalarCand struct {
		name, formula string
		mag           float64
	}
	scalars := []scalarCand{
		{"B-sector gap loop-scaled", "gap_B/(16π²)", spec.BGapValue * s.EpsilonU},
		{"inverse contact zeta loop-scaled", "zeta(1)^-1/(16π²)", spec.ContactZetaValues["inverse zeta(1)"] * s.EpsilonU},
		{"contact count inverse loop-scaled", "1/(7·16π²)", (1.0 / 7.0) * s.EpsilonU},
		{"contact mean loop-scaled", "mean(lambda_partial)/(16π²)", spec.ContactMean * s.EpsilonU},
		{"tau_eta unit loop-scaled", "1/(16π²)", s.EpsilonU},
	}
	best := scalarCand{mag: math.Inf(1)}
	bestErr := math.Inf(1)
	for _, x := range scalars {
		err := math.Abs(x.mag - target)
		if err < bestErr {
			best, bestErr = x, err
		}
	}
	exact := 0
	for _, x := range c {
		if x.CanonicalMagnitudeMatches {
			exact++
		}
	}
	return CoefficientSearchAudit{
		CanonicalScalarsAudited:        len(scalars),
		CanonicalLoopScaledCandidates:  len(scalars),
		ExactMagnitudeMatches:          exact,
		ClosestCandidateName:           best.name,
		ClosestCandidateFormula:        best.formula,
		ClosestCandidateMagnitude:      best.mag,
		ClosestCandidateRatioToTarget:  best.mag / target,
		ClosestCandidateAbsError:       bestErr,
		ClosestCandidateAccepted:       false,
		FittedCoefficientNeeded:        true,
		ArbitraryNormalizationInserted: false,
		Verdict:                        "no branch-free canonical scalar equals the required 5.6e-4 magnitude; the closest loop-scaled scalar remains a near-miss and is rejected",
	}
}

func auditFirewall(s Gate215Snapshot, spec SpectralDataAudit, hk HeatKernelMappingAudit, coef CoefficientSearchAudit) FirewallAudit {
	return FirewallAudit{
		Gate215Inherited:                s.Gate215Inherited,
		ThresholdSpectrumSealInherited:  s.ThresholdSpectrumSealInherited,
		EmpiricalCarrierSealInherited:   true,
		LeptoquarkDynamicsSealInherited: true,
		EmpiricalLedgerQuarantined:      true,
		MatchingResidualPromoted:        false,
		MatchingCorrectionsDerived:      hk.DeltaMatchRowsDerived > 0 || spec.FiniteMatchingRowsDerived,
		SpectralCoefficientTuned:        false,
		HeatKernelMapImported:           false,
		MSbarSchemeImported:             false,
		PhysicalMassPredictionClaimed:   false,
		PhysicalUnificationClaimed:      false,
		ContactModesPromotedToParticles: false,
		BGapPromotedToMass:              false,
		ProtonLifetimeComputed:          false,
		RecommendedNextGate:             "Gate 217 — finite spectral triple / gauge-curvature projection construction audit, or matching-correction seal formalization",
		OpenRequirements: []string{
			"derive a finite Dirac/spectral-triple operator for the active threshold sector",
			"derive a gauge-curvature projection from finite traces to U(1), SU(2), and SU(3) kinetic rows",
			"derive cutoff moments or a subtraction scheme before interpreting traces as δ_i^match",
			"derive threshold matching rows rather than fitting the Gate-215 residual vector",
		},
		Verdict: "all Gate-216 spectral comparisons remain diagnostic; no residual or spectral scalar is promoted to a matching correction",
	}
}

func buildTruth(s Summary, g Gate215Snapshot, hk HeatKernelMappingAudit, c CoefficientSearchAudit) string {
	if s.Status == StatusConditionalSupport {
		return "Gate 216 found a canonical spectral matching-residual bridge."
	}
	return fmt.Sprintf("Gate 216 audits the Gate-215 target %s for the unique single-scale spectrum %s. Existing finite data contains positive spectral anchors and an eta-signed trace with a sign-only resonance, but no finite heat-kernel gauge projection, cutoff coefficient, or subtraction scheme maps those traces to δ_i^match. The closest canonical loop-scaled magnitude (%s=%0.12g, ratio %.6g) is a rejected near-miss. Therefore the required matching residual remains an external target for future finite spectral-triple/matching theory.", g.RequiredDeltaMatch.String(), g.BestRows, c.ClosestCandidateFormula, c.ClosestCandidateMagnitude, c.ClosestCandidateRatioToTarget)
}

func meanMinMax(xs []float64) (float64, float64, float64) {
	if len(xs) == 0 {
		return 0, 0, 0
	}
	sum := 0.0
	minv, maxv := xs[0], xs[0]
	for _, x := range xs {
		sum += x
		if x < minv {
			minv = x
		}
		if x > maxv {
			maxv = x
		}
	}
	return sum / float64(len(xs)), minv, maxv
}

func uniform(x float64) FloatTriple { return FloatTriple{x, x, x} }

func signPattern(t FloatTriple) string {
	chars := []byte{'0', '0', '0'}
	for i := 0; i < 3; i++ {
		x := t.At(i)
		if x > 0 {
			chars[i] = '+'
		}
		if x < 0 {
			chars[i] = '-'
		}
	}
	return string(chars)
}

func vectorClose(a, b FloatTriple, tol float64) bool {
	return math.Abs(a.U1GUT-b.U1GUT) <= tol && math.Abs(a.SU2L-b.SU2L) <= tol && math.Abs(a.SU3C-b.SU3C) <= tol
}

func maxAbs3(t FloatTriple) float64 { return t.MaxAbs() }

func bestFittedScale(v, target FloatTriple) (float64, float64) {
	den := v.U1GUT*v.U1GUT + v.SU2L*v.SU2L + v.SU3C*v.SU3C
	if den == 0 {
		return 0, math.Inf(1)
	}
	num := v.U1GUT*target.U1GUT + v.SU2L*target.SU2L + v.SU3C*target.SU3C
	c := num / den
	res := v.Scale(c).Sub(target)
	return c, res.MaxAbs()
}

func FormatGate215(s Gate215Snapshot) string {
	return fmt.Sprintf("inherited=%t seal=%t envelope=%t unique=%t rows=%s MB=%.12g M*=%.12g delta=%s pattern=%s max=%.12g overEps=%.6g classes=%d/%d finiteMatching=%t", s.Gate215Inherited, s.ThresholdSpectrumSealInherited, s.MatchingEnvelopeInherited, s.SingleScaleCandidateUnique, s.BestRows, s.MBGeV, s.MStarGeV, s.RequiredDeltaMatch.String(), s.RequiredSignPattern, s.RequiredMaxAbs, s.RequiredOverEnvelope, s.PlausibleClasses, s.ClassesAudited, s.FiniteMatchingDerived)
}

func FormatSpectralData(s SpectralDataAudit) string {
	return fmt.Sprintf("BGap=%.12g contactModes=%d mean=%.12g range=[%.12g,%.12g] scalarModes=%d zeta=%d tau=%s tauFlip=%s finiteClass=%t zetaLedger=%t heatMap=%t cutoff=%t subtraction=%t matchingRows=%t (%s)", s.BGapValue, s.ContactPartialModeCount, s.ContactMean, s.ContactMin, s.ContactMax, s.ScalarActiveModeCount, len(s.ContactZetaValues), s.TauEtaDegrees.String(), s.TauEtaOrientationFlipped.String(), s.FiniteFundamentalClassAvailable, s.ContactZetaLedgerAvailable, s.HeatKernelCoefficientMapDerived, s.CanonicalCutoffFunctionDerived, s.CanonicalSubtractionSchemeDerived, s.FiniteMatchingRowsDerived, s.Verdict)
}

func FormatResidualTarget(s ResidualStructureAudit) string {
	return fmt.Sprintf("target=%s pattern=%s normalized=%s alternating=%t equalMag=%t derived=%t (%s)", s.TargetVector.String(), s.TargetSignPattern, s.TargetNormalized.String(), s.AlternatingPattern, s.NearlyEqualMagnitudes, s.TraceTargetInterpretedAsDerived, s.Verdict)
}

func FormatCandidates(c []TraceCandidate) string {
	parts := make([]string, 0, len(c))
	for _, x := range c {
		parts = append(parts, fmt.Sprintf("%s vector=%s pattern=%s norm=%s signs=%t rel=%t mag=%t fitC=%.6g fitMax=%.6g promoted=%t", x.Name, x.Vector.String(), x.SignPattern, x.Normalized.String(), x.SignsMatch, x.RelativeMagnitudesMatch, x.CanonicalMagnitudeMatches, x.BestFittedScale, x.BestFittedMaxResidual, x.PromotedToMatchingRow))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatHeatKernel(s HeatKernelMappingAudit) string {
	return fmt.Sprintf("a2a4=%t dirac=%t triple=%t projectionRows=%d gaugeMap=%t cutoff=%t msbar=%t subtraction=%t deltaRows=%d signOnly=%d structural=%d (%s)", s.A2A4LanguageAudited, s.FiniteDiracOperatorDerived, s.SpectralTripleComplete, s.GaugeCurvatureProjectionRows, s.GaugeKineticTraceMapDerived, s.CanonicalCutoffMomentsDerived, s.CanonicalMSbarSchemeImported, s.ThresholdSubtractionSchemeDerived, s.DeltaMatchRowsDerived, s.SignOnlyResonances, s.FullStructuralMatches, s.Verdict)
}

func FormatCoefficient(s CoefficientSearchAudit) string {
	return fmt.Sprintf("scalars=%d loopScaled=%d exact=%d closest=%s formula=%s mag=%.12g ratio=%.9g err=%.12g accepted=%t fittedNeeded=%t tuned=%t (%s)", s.CanonicalScalarsAudited, s.CanonicalLoopScaledCandidates, s.ExactMagnitudeMatches, s.ClosestCandidateName, s.ClosestCandidateFormula, s.ClosestCandidateMagnitude, s.ClosestCandidateRatioToTarget, s.ClosestCandidateAbsError, s.ClosestCandidateAccepted, s.FittedCoefficientNeeded, s.ArbitraryNormalizationInserted, s.Verdict)
}

func FormatFirewall(s FirewallAudit) string {
	return fmt.Sprintf("gate215=%t thresholdSeal=%t carrierSeal=%t lqSeal=%t ledger=%t residualPromoted=%t matchingDerived=%t tuned=%t heatMapImported=%t msbar=%t massClaim=%t unificationClaim=%t contactPromoted=%t bgapMass=%t protonLifetime=%t next=%s open=[%s] (%s)", s.Gate215Inherited, s.ThresholdSpectrumSealInherited, s.EmpiricalCarrierSealInherited, s.LeptoquarkDynamicsSealInherited, s.EmpiricalLedgerQuarantined, s.MatchingResidualPromoted, s.MatchingCorrectionsDerived, s.SpectralCoefficientTuned, s.HeatKernelMapImported, s.MSbarSchemeImported, s.PhysicalMassPredictionClaimed, s.PhysicalUnificationClaimed, s.ContactModesPromotedToParticles, s.BGapPromotedToMass, s.ProtonLifetimeComputed, s.RecommendedNextGate, strings.Join(s.OpenRequirements, "; "), s.Verdict)
}

func FormatSummary(s Summary) string {
	return fmt.Sprintf("tests=%d gate215=%t spectral=%t candidates=%d signOnly=%d structural=%d magMatches=%d status=%s (%s)", s.TestsAudited, s.Gate215Inherited, s.SpectralDataAudited, s.TraceCandidatesAudited, s.SignOnlyResonances, s.StructuralMatches, s.CanonicalMagnitudeMatches, s.Status, s.Comment)
}
