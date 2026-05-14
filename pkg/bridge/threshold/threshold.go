// Package threshold audits whether the finite engine has derived a heavy
// threshold spectrum suitable for RG matching.
//
// The package is intentionally strict.  The engine has many dimensionless
// spectral anchors: B-sector gaps, contact partial-overlap modes, scalar/contact
// active eigenvalues, and leakage invariants.  Those are real finite data.  They
// are not yet physical threshold masses, because no dimensionful unit, activation
// rule, or finite-to-continuum matching prescription has been derived.
package threshold

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/betacoeff"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarscale"
	"github.com/bagherbal/asha-engine/pkg/dynamics/bsector"
	"github.com/bagherbal/asha-engine/pkg/dynamics/scalarpotential"
)

type CandidateKind string

const (
	ScalarActiveCandidate   CandidateKind = "scalar-active"
	BGapCandidate           CandidateKind = "B-sector-gap"
	ContactOverlapCandidate CandidateKind = "contact-overlap"
	LeakageCandidate        CandidateKind = "contact-leakage"
	RadialCandidate         CandidateKind = "scalar-radial"
)

type Candidate struct {
	Name                string
	Kind                CandidateKind
	Value               float64
	Multiplicity        int
	Formula             string
	PhysicalMassDerived bool
}

type Cluster struct {
	Value        float64
	Multiplicity int
}

type Analysis struct {
	Beta   betacoeff.Analysis
	Scale  scalarscale.Analysis
	Scalar scalarpotential.Analysis
	Vacuum bsector.Vacuum

	BPositiveEigenvalues  []float64
	BGap                  float64
	ContactPartialOverlap []float64
	ScalarActiveSpectrum  []float64
	ScalarClusters        []Cluster

	Candidates []Candidate

	DimensionlessSpectralAnchorsAvailable bool
	PhysicalMassUnitDerived               bool
	ThresholdActivationRuleDerived        bool
	FiniteToContinuumMatchingDerived      bool
	ThresholdCorrectedBetaDerived         bool
	HiddenThresholdScaleInserted          bool
	ObservedMassesUsed                    bool

	ThresholdMassFamily string
	TruthStatement      string
	MinimumMissingData  []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		beta, err := betacoeff.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		scale, err := scalarscale.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		scalar, err := scalarpotential.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		vac, err := bsector.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(beta, scale, scalar, vac, 1e-8)
	})
	return defaultValue, defaultErr
}

func Build(beta betacoeff.Analysis, scale scalarscale.Analysis, scalar scalarpotential.Analysis, vac bsector.Vacuum, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-8
	}
	bpos := positive(vac.Eigenvalues, eps)
	if len(bpos) == 0 {
		return Analysis{}, fmt.Errorf("B-sector has no positive spectrum")
	}
	partial := contactPartialModes(vac.Contact.OverlapEigenvalues, eps)
	if len(partial) == 0 {
		return Analysis{}, fmt.Errorf("contact overlap has no partial modes")
	}
	active := append([]float64(nil), scalar.ActiveSpectrum...)
	if len(active) != 4 {
		return Analysis{}, fmt.Errorf("expected four scalar active values, got %d", len(active))
	}

	candidates := make([]Candidate, 0)
	for i, v := range active {
		candidates = append(candidates, Candidate{
			Name:         fmt.Sprintf("scalar active eigenvalue %d", i+1),
			Kind:         ScalarActiveCandidate,
			Value:        v,
			Multiplicity: 1,
			Formula:      "M_i(mu)=mu*sqrt(lambda_i)",
		})
	}
	candidates = append(candidates, Candidate{Name: "B-sector first spectral gap", Kind: BGapCandidate, Value: vac.FirstPositiveEigenvalue(eps), Multiplicity: 1, Formula: "M_gap(mu)=mu*sqrt(gap)"})
	candidates = append(candidates, Candidate{Name: "contact leakage norm squared", Kind: LeakageCandidate, Value: vac.Contact.BareLeakageNormSquared(), Multiplicity: 1, Formula: "dimensionless frustration invariant, not a mass"})
	candidates = append(candidates, Candidate{Name: "scalar radial curvature", Kind: RadialCandidate, Value: scalar.DimensionlessRadialMassSq, Multiplicity: 1, Formula: "m_radial(mu)=mu*sqrt(m_radial_hat^2)"})
	for i, v := range partial {
		candidates = append(candidates, Candidate{Name: fmt.Sprintf("contact partial-overlap mode %d", i+1), Kind: ContactOverlapCandidate, Value: v, Multiplicity: 1, Formula: "dimensionless overlap mode; activation rule not derived"})
	}

	for i := range candidates {
		candidates[i].PhysicalMassDerived = false
	}

	return Analysis{
		Beta:                                  beta,
		Scale:                                 scale,
		Scalar:                                scalar,
		Vacuum:                                vac,
		BPositiveEigenvalues:                  bpos,
		BGap:                                  vac.FirstPositiveEigenvalue(eps),
		ContactPartialOverlap:                 partial,
		ScalarActiveSpectrum:                  active,
		ScalarClusters:                        clusters(active, eps),
		Candidates:                            candidates,
		DimensionlessSpectralAnchorsAvailable: true,
		PhysicalMassUnitDerived:               false,
		ThresholdActivationRuleDerived:        false,
		FiniteToContinuumMatchingDerived:      false,
		ThresholdCorrectedBetaDerived:         false,
		HiddenThresholdScaleInserted:          false,
		ObservedMassesUsed:                    false,
		ThresholdMassFamily:                   "threshold candidates remain families M_i(mu)=mu*f_i until a non-fitted physical unit and activation rule are derived",
		TruthStatement:                        "The finite engine exposes several dimensionless spectral threshold anchors, but no physical heavy threshold spectrum is yet derived. Threshold matching needs a mass unit, an activation/decoupling rule, and a finite-to-continuum map assigning which modes enter each gauge beta function.",
		MinimumMissingData: []string{
			"derive the physical unit mu or boundary scale M* without fitting observed masses",
			"derive which finite modes are continuum-active and which are integrated out",
			"derive representation assignments for each threshold mode under SU(3)c, SU(2)L, and U(1)Y",
			"derive matching/decoupling rules for threshold corrections to b1, b2, and b3",
			"derive whether contact partial-overlap modes are physical thresholds, regulator artifacts, or vacuum-frustration modes",
		},
	}, nil
}

func positive(values []float64, eps float64) []float64 {
	out := make([]float64, 0)
	for _, v := range values {
		if v > eps {
			out = append(out, v)
		}
	}
	sort.Float64s(out)
	return out
}

func contactPartialModes(values []float64, eps float64) []float64 {
	out := make([]float64, 0)
	for _, v := range values {
		if v > eps && math.Abs(v-1) > eps {
			out = append(out, v)
		}
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(out)))
	return out
}

func clusters(values []float64, eps float64) []Cluster {
	if len(values) == 0 {
		return nil
	}
	vals := append([]float64(nil), values...)
	sort.Sort(sort.Reverse(sort.Float64Slice(vals)))
	out := make([]Cluster, 0)
	for i := 0; i < len(vals); {
		j := i + 1
		sum := vals[i]
		for j < len(vals) && math.Abs(vals[j]-vals[i]) <= eps {
			sum += vals[j]
			j++
		}
		out = append(out, Cluster{Value: sum / float64(j-i), Multiplicity: j - i})
		i = j
	}
	return out
}

func FormatCandidates(c []Candidate, max int) string {
	if max <= 0 || max > len(c) {
		max = len(c)
	}
	parts := make([]string, 0, max)
	for i := 0; i < max; i++ {
		parts = append(parts, fmt.Sprintf("%s=%.10f", c[i].Name, c[i].Value))
	}
	if max < len(c) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(c)-max))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func FormatClusters(c []Cluster) string {
	parts := make([]string, 0, len(c))
	for _, cl := range c {
		parts = append(parts, fmt.Sprintf("m=%d@%.10f", cl.Multiplicity, cl.Value))
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func FormatFloats(values []float64, max int) string {
	if max <= 0 || max > len(values) {
		max = len(values)
	}
	parts := make([]string, 0, max)
	for i := 0; i < max; i++ {
		parts = append(parts, fmt.Sprintf("%.10f", values[i]))
	}
	if max < len(values) {
		parts = append(parts, fmt.Sprintf("... +%d more", len(values)-max))
	}
	return "[" + strings.Join(parts, ", ") + "]"
}
