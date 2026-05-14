// Package sectorspectrum performs Gate 65: current-sector spectral assignment
// search.
//
// Gate 64 exposed positive finite spectral anchors that could be used as
// propagator denominator families in diagnostics. Gate 65 asks a stricter
// question: can any of those spectral anchors be assigned to the current
// sectors central, color-su3, B-L, and leptoquark by a representation-level
// theorem rather than by numerology?
//
// The result is deliberately conservative. Multiplicity/count matches are
// audited, but they are not enough. A valid assignment requires a finite
// operator map from the current-sector representation into the spectral carrier.
// At this stage no such map exists. Therefore the propagator denominators rho_A
// remain unassigned and the exchange/NJL kernel remains open.
package sectorspectrum

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/propagatorspectrum"
)

type CurrentSector struct {
	Name           string
	GeneratorCount int
	Role           string
}

type SpectralCarrier struct {
	Name          string
	Multiplicity  int
	Dimension     int
	ValuesSummary string
	Kind          string
}

type AssignmentAttempt struct {
	Sector            CurrentSector
	BestCarrier       SpectralCarrier
	MultiplicityGap   int
	CountMatch        bool
	RepresentationMap bool
	Assigned          bool
	Reason            string
}

type Analysis struct {
	Propagator propagatorspectrum.Analysis

	Sectors  []CurrentSector
	Carriers []SpectralCarrier
	Attempts []AssignmentAttempt

	ExactCountMatches               int
	PotentialScalarSingletonMatches int
	RequiredSectorAssignments       int
	RepresentationMapsDerived       int
	AssignedSectorCount             int
	AllSectorsAssigned              bool
	PropagatorDenominatorsDerived   bool
	ExchangeKernelUpdated           bool
	AttractiveScalarChannelDerived  bool
	UpDownSplittingDerived          bool
	RegulatorCriticalityDerived     bool
	CondensationClaimAllowed        bool
	HiddenObservedInputUsed         bool

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
		p, err := propagatorspectrum.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(p)
	})
	return defaultValue, defaultErr
}

func Build(p propagatorspectrum.Analysis) (Analysis, error) {
	if len(p.Exchange.SectorDiagnostics) == 0 {
		return Analysis{}, fmt.Errorf("Gate 65 requires Gate 63/64 current-sector diagnostics")
	}
	if len(p.Families) == 0 {
		return Analysis{}, fmt.Errorf("Gate 65 requires finite denominator families from Gate 64")
	}

	sectors := []CurrentSector{
		{Name: "central", GeneratorCount: 1, Role: "overall u(1) current"},
		{Name: "color-su3", GeneratorCount: 8, Role: "color adjoint current sector"},
		{Name: "b-minus-l", GeneratorCount: 1, Role: "B-L abelian current"},
		{Name: "leptoquark", GeneratorCount: 6, Role: "off-diagonal SU(4)/(SU(3)xU(1)) current sector"},
	}

	carriers := []SpectralCarrier{
		{Name: "B-sector positive spectrum", Multiplicity: len(p.Threshold.BPositiveEigenvalues), Dimension: len(p.Threshold.BPositiveEigenvalues), Kind: "B-sector", ValuesSummary: fmt.Sprintf("%d positive modes, first gap %.10f", len(p.Threshold.BPositiveEigenvalues), p.Threshold.BGap)},
		{Name: "contact partial-overlap spectrum", Multiplicity: len(p.Threshold.ContactPartialOverlap), Dimension: len(p.Threshold.ContactPartialOverlap), Kind: "contact-partial", ValuesSummary: propagatorspectrumFloatList(p.Threshold.ContactPartialOverlap, 7)},
		{Name: "scalar/contact active spectrum", Multiplicity: len(p.Threshold.ScalarActiveSpectrum), Dimension: len(p.Threshold.ScalarActiveSpectrum), Kind: "scalar-active", ValuesSummary: propagatorspectrumFloatList(p.Threshold.ScalarActiveSpectrum, 4)},
		{Name: "contact leakage invariant", Multiplicity: 1, Dimension: 1, Kind: "vacuum-frustration", ValuesSummary: "single scalar invariant L_BG^2"},
		{Name: "scalar radial curvature", Multiplicity: 1, Dimension: 1, Kind: "scalar-radial", ValuesSummary: "single scalar-sector radial invariant"},
	}

	attempts := make([]AssignmentAttempt, 0, len(sectors))
	exactMatches := 0
	singletonMatches := 0
	for _, s := range sectors {
		best := bestMultiplicityCarrier(s, carriers)
		gap := abs(s.GeneratorCount - best.Multiplicity)
		countMatch := gap == 0
		if countMatch {
			exactMatches++
		}
		if s.GeneratorCount == 1 {
			singletonMatches += 2 // leakage and radial are singleton scalar anchors, but ambiguous and not current representations.
		}
		reason := reasonForRejection(s, best, countMatch)
		attempts = append(attempts, AssignmentAttempt{
			Sector:            s,
			BestCarrier:       best,
			MultiplicityGap:   gap,
			CountMatch:        countMatch,
			RepresentationMap: false,
			Assigned:          false,
			Reason:            reason,
		})
	}

	truth := "Gate 65 finds no canonical current-sector spectral assignment. The finite spectra contain real anchors, but the multiplicities do not match the current-sector decomposition in a representation-theoretic way: color-su3 needs an adjoint 8-carrier, leptoquark needs a 6-carrier, and the two abelian sectors cannot be distinguished by scalar singleton invariants. Count proximity is not a propagator theorem. Therefore rho_A for central, color, B-L, and leptoquark remains unassigned."

	return Analysis{
		Propagator:                      p,
		Sectors:                         sectors,
		Carriers:                        carriers,
		Attempts:                        attempts,
		ExactCountMatches:               exactMatches,
		PotentialScalarSingletonMatches: singletonMatches,
		RequiredSectorAssignments:       len(sectors),
		RepresentationMapsDerived:       0,
		AssignedSectorCount:             0,
		AllSectorsAssigned:              false,
		PropagatorDenominatorsDerived:   false,
		ExchangeKernelUpdated:           false,
		AttractiveScalarChannelDerived:  false,
		UpDownSplittingDerived:          false,
		RegulatorCriticalityDerived:     false,
		CondensationClaimAllowed:        false,
		HiddenObservedInputUsed:         false,
		TruthStatement:                  truth,
		RecommendedNextGate:             "Gate 66 — Current-Sector Operator Construction Search",
		RemainingUnknowns: []string{
			"U-20D2B1A-COLOR-SPECTRAL-CARRIER: derive an adjoint 8-dimensional spectral carrier for color-su3, not a 7-mode near miss",
			"U-20D2B1B-LEPTOQUARK-CARRIER: derive a 6-dimensional leptoquark spectral carrier or reject leptoquark exchange from the NJL kernel",
			"U-20D2B1C-ABELIAN-SEPARATION: distinguish central and B-L propagators by a kinetic/action theorem, not singleton scalar anchors",
			"U-20D2B2-PROPAGATOR-DENOMINATORS: derive rho_A after representation carriers exist",
			"U-20D4-UP-DOWN-SPLITTING: derive the finite operator selecting top-like up over bottom-like down",
		},
	}, nil
}

func bestMultiplicityCarrier(s CurrentSector, carriers []SpectralCarrier) SpectralCarrier {
	best := carriers[0]
	bestGap := abs(s.GeneratorCount - best.Multiplicity)
	for _, c := range carriers[1:] {
		gap := abs(s.GeneratorCount - c.Multiplicity)
		if gap < bestGap {
			best = c
			bestGap = gap
		}
	}
	return best
}

func reasonForRejection(s CurrentSector, c SpectralCarrier, countMatch bool) string {
	switch s.Name {
	case "color-su3":
		return "color requires an adjoint 8-carrier with SU(3) action; current best finite carrier is not an adjoint representation"
	case "leptoquark":
		return "leptoquark requires a 6-carrier with off-diagonal SU(4)/(SU(3)xU(1)) action; no 6-mode spectral carrier is derived"
	case "central", "b-minus-l":
		return "abelian singleton spectral anchors exist, but no theorem separates central from B-L propagator denominators"
	default:
		if countMatch {
			return "count match alone is insufficient; representation map is missing"
		}
		return fmt.Sprintf("multiplicity mismatch with best carrier %s", c.Name)
	}
}

func propagatorspectrumFloatList(xs []float64, max int) string {
	if max <= 0 || max > len(xs) {
		max = len(xs)
	}
	parts := make([]string, 0, max)
	for i := 0; i < max; i++ {
		parts = append(parts, fmt.Sprintf("%.10f", xs[i]))
	}
	if max < len(xs) {
		parts = append(parts, fmt.Sprintf("... +%d", len(xs)-max))
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func FormatAttempts(xs []AssignmentAttempt) string {
	ys := append([]AssignmentAttempt(nil), xs...)
	sort.Slice(ys, func(i, j int) bool { return ys[i].Sector.Name < ys[j].Sector.Name })
	parts := make([]string, 0, len(ys))
	for _, a := range ys {
		parts = append(parts, fmt.Sprintf("%s(n=%d) -> best=%s(m=%d, gap=%d, assigned=%v)", a.Sector.Name, a.Sector.GeneratorCount, a.BestCarrier.Name, a.BestCarrier.Multiplicity, a.MultiplicityGap, a.Assigned))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatCarriers(xs []SpectralCarrier) string {
	ys := append([]SpectralCarrier(nil), xs...)
	sort.Slice(ys, func(i, j int) bool { return ys[i].Name < ys[j].Name })
	parts := make([]string, 0, len(ys))
	for _, x := range ys {
		parts = append(parts, fmt.Sprintf("%s(m=%d, kind=%s, %s)", x.Name, x.Multiplicity, x.Kind, x.ValuesSummary))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatUnknowns(xs []string) string {
	ys := append([]string(nil), xs...)
	sort.Strings(ys)
	return "[" + strings.Join(ys, "; ") + "]"
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
