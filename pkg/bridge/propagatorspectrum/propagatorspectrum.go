// Package propagatorspectrum performs Gate 64: finite propagator denominator
// search from B-sector, contact, and scalar spectra.
//
// Gate 63 exposed diagnostic exchange kernels, but no finite action selected
// propagator denominators rho_A for the current sectors. This gate asks whether
// the spectra already computed by the finite engine can supply those
// denominators. The answer is deliberately strict: many positive spectral
// anchors exist, but no representation-level map assigns any anchor to the
// central, color, B-L, or leptoquark current sectors. Therefore all spectral
// propagator kernels remain diagnostics, not derived four-fermion couplings.
package propagatorspectrum

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/exchangeaction"
	"github.com/bagherbal/asha-engine/pkg/bridge/threshold"
)

type DenominatorFamily struct {
	Name           string
	Value          float64
	Formula        string
	KernelIfGlobal float64
	Selected       bool
}

type Analysis struct {
	Exchange  exchangeaction.Analysis
	Threshold threshold.Analysis

	CurrentSectorCount          int
	ThresholdCandidateCount     int
	PositiveSpectralAnchorCount int

	Families []DenominatorFamily

	SmallestDenominator       float64
	LargestDenominator        float64
	StrongestDiagnosticKernel float64
	StrongestDiagnosticFamily string

	SpectralDenominatorsAvailable         bool
	SectorSpectralAssignmentDerived       bool
	CurrentSectorRepresentationMapDerived bool
	PropagatorDenominatorsDerived         bool
	ExchangeKernelUpdated                 bool
	AttractiveKernelDerived               bool
	UpDownSplittingDerived                bool
	RegulatorCriticalityDerived           bool
	CondensationClaimAllowed              bool
	HiddenObservedInputUsed               bool

	TruthStatement      string
	RecommendedNextGate string
	RemainingUnknowns   []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		ex, err := exchangeaction.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		th, err := threshold.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(ex, th, 1e-12)
	})
	return defaultValue, defaultErr
}

func Build(ex exchangeaction.Analysis, th threshold.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-12
	}
	if ex.UnitAttractiveKernel <= eps {
		return Analysis{}, fmt.Errorf("Gate 64 requires Gate 63 unit attractive diagnostic")
	}
	if len(ex.SectorDiagnostics) == 0 {
		return Analysis{}, fmt.Errorf("no current-sector diagnostics available")
	}
	if len(th.Candidates) == 0 {
		return Analysis{}, fmt.Errorf("no finite threshold/spectral candidates available")
	}

	activeMean := meanPositive(th.ScalarActiveSpectrum, eps)
	partialMean := meanPositive(th.ContactPartialOverlap, eps)
	bGap := th.BGap
	leakage := findCandidate(th.Candidates, threshold.LeakageCandidate)
	radial := findCandidate(th.Candidates, threshold.RadialCandidate)

	families := make([]DenominatorFamily, 0)
	addFamily := func(name string, value float64, formula string) {
		if value > eps && !math.IsNaN(value) && !math.IsInf(value, 0) {
			families = append(families, DenominatorFamily{
				Name:           name,
				Value:          value,
				Formula:        formula,
				KernelIfGlobal: ex.UnitAttractiveKernel / value,
				Selected:       false,
			})
		}
	}
	addFamily("B-sector first gap", bGap, "rho_A = first positive eigenvalue of O_B")
	addFamily("mean scalar/contact active eigenvalue", activeMean, "rho_A = mean(lambda_active)")
	addFamily("mean contact partial-overlap", partialMean, "rho_A = mean(lambda_partial)")
	addFamily("contact leakage norm squared", leakage, "rho_A = L_BG^2")
	addFamily("scalar radial curvature", radial, "rho_A = m_radial_hat^2")

	if len(families) == 0 {
		return Analysis{}, fmt.Errorf("no positive denominator family could be formed")
	}
	sort.Slice(families, func(i, j int) bool { return families[i].Value < families[j].Value })

	minDen := families[0].Value
	maxDen := families[len(families)-1].Value
	strongest := math.Inf(-1)
	strongestName := ""
	for _, f := range families {
		if f.KernelIfGlobal > strongest {
			strongest = f.KernelIfGlobal
			strongestName = f.Name
		}
	}

	truth := "Gate 64 finds real finite spectral anchors that could be used as propagator denominators in diagnostic branches. However, no representation-level theorem maps the B-sector gap, scalar/contact eigenvalues, contact partial modes, leakage invariant, or radial curvature to the central, color, B-L, or leptoquark current sectors. Therefore spectral denominators are available, but current-sector propagator weights remain unassigned and the NJL exchange kernel remains open."

	return Analysis{
		Exchange:                              ex,
		Threshold:                             th,
		CurrentSectorCount:                    len(ex.SectorDiagnostics),
		ThresholdCandidateCount:               len(th.Candidates),
		PositiveSpectralAnchorCount:           len(th.BPositiveEigenvalues) + len(th.ContactPartialOverlap) + len(th.ScalarActiveSpectrum),
		Families:                              families,
		SmallestDenominator:                   minDen,
		LargestDenominator:                    maxDen,
		StrongestDiagnosticKernel:             strongest,
		StrongestDiagnosticFamily:             strongestName,
		SpectralDenominatorsAvailable:         true,
		SectorSpectralAssignmentDerived:       false,
		CurrentSectorRepresentationMapDerived: false,
		PropagatorDenominatorsDerived:         false,
		ExchangeKernelUpdated:                 false,
		AttractiveKernelDerived:               false,
		UpDownSplittingDerived:                false,
		RegulatorCriticalityDerived:           false,
		CondensationClaimAllowed:              false,
		HiddenObservedInputUsed:               false,
		TruthStatement:                        truth,
		RecommendedNextGate:                   "Gate 65 — Current-Sector Spectral Assignment Search",
		RemainingUnknowns: []string{
			"U-20D2B1-SECTOR-SPECTRAL-MAP: map current sectors to finite kinetic/spectral operators",
			"U-20D2B2-PROPAGATOR-DENOMINATORS: derive rho_A for central, color, B-L, and leptoquark sectors",
			"U-20D3-RELATIVE-COUPLINGS: derive relative current-sector coupling strengths",
			"U-20D4-UP-DOWN-SPLITTING: derive the operator selecting top-like up over bottom-like down",
			"U-20D6-CRITICAL-REGULATOR: derive the finite NJL threshold C_reg",
		},
	}, nil
}

func meanPositive(xs []float64, eps float64) float64 {
	sum := 0.0
	n := 0
	for _, x := range xs {
		if x > eps {
			sum += x
			n++
		}
	}
	if n == 0 {
		return 0
	}
	return sum / float64(n)
}

func findCandidate(c []threshold.Candidate, kind threshold.CandidateKind) float64 {
	for _, x := range c {
		if x.Kind == kind {
			return x.Value
		}
	}
	return 0
}

func FormatFamilies(xs []DenominatorFamily) string {
	ys := append([]DenominatorFamily(nil), xs...)
	sort.Slice(ys, func(i, j int) bool { return ys[i].Value < ys[j].Value })
	parts := make([]string, 0, len(ys))
	for _, x := range ys {
		parts = append(parts, fmt.Sprintf("%s(rho=%.10f, Gdiag=%.10f, selected=%v)", x.Name, x.Value, x.KernelIfGlobal, x.Selected))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUnknowns(xs []string) string {
	ys := append([]string(nil), xs...)
	sort.Strings(ys)
	return "[" + strings.Join(ys, "; ") + "]"
}
