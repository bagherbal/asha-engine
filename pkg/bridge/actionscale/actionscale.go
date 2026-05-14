// Package actionscale audits whether finite topological/action normalization can
// supply the missing dimensionful scalar scale.
//
// The finite source architecture has strong dimensionless anchors: the contact
// index I_BG=1, a scalar radius, a B-sector gap, and a contact leakage norm. The
// natural topological action normalization associated with a unit index is the
// Yang--Mills/instanton-shaped seal 8*pi^2*I_BG. That is valuable, but it is
// still dimensionless in natural units. It can normalize a coupling or a path
// integral weight; it cannot by itself choose a physical mass unit in GeV.
//
// This package therefore implements the scale firewall for gravity/action
// normalization: derive all dimensionless action data available, then reject any
// claim that the electroweak scale, Planck scale, or Higgs mass has been obtained
// unless a genuine dimensional bridge is present.
package actionscale

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/scalarscale"
)

type Analysis struct {
	ScalarScale scalarscale.Analysis

	ContactIndex float64

	// TopologicalActionSeal is the dimensionless action normalization that would
	// correspond to a unit instanton/topological index in continuum Yang--Mills
	// normalization. It is an action/coupling normalization candidate, not a mass.
	TopologicalActionSeal float64
	ActionWeight          float64

	RadiusSquared float64
	Radius        float64
	RadialMassSq  float64
	BGap          float64
	LeakageSq     float64

	ActionToRadiusRatio  float64
	ActionToGapRatio     float64
	ActionToLeakageRatio float64

	UnitIndexAvailable          bool
	DimensionlessActionDerived  bool
	ContinuumIndexBridgeDerived bool
	CouplingNormalizationOpen   bool
	DimensionfulUnitDerived     bool
	GravityScaleDerived         bool
	ScalarScaleFixed            bool
	HiddenObservedScaleInserted bool

	TruthStatement    string
	RemainingUnknowns []string
}

var (
	defaultOnce  sync.Once
	defaultValue Analysis
	defaultErr   error
)

func BuildDefault() (Analysis, error) {
	defaultOnce.Do(func() {
		ss, err := scalarscale.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(ss, 1e-12)
	})
	return defaultValue, defaultErr
}

func Build(ss scalarscale.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-12
	}
	if ss.ContactIndex <= eps {
		return Analysis{}, fmt.Errorf("contact index must be positive")
	}
	if ss.FiniteRadiusSquared <= eps || ss.BGap <= eps || ss.ContactLeakageNormSquared <= eps {
		return Analysis{}, fmt.Errorf("scalar-scale anchors are incomplete")
	}

	seal := 8 * math.Pi * math.Pi * ss.ContactIndex
	weight := math.Exp(-seal)

	return Analysis{
		ScalarScale:                 ss,
		ContactIndex:                ss.ContactIndex,
		TopologicalActionSeal:       seal,
		ActionWeight:                weight,
		RadiusSquared:               ss.FiniteRadiusSquared,
		Radius:                      ss.FiniteRadius,
		RadialMassSq:                ss.DimensionlessRadialMassSq,
		BGap:                        ss.BGap,
		LeakageSq:                   ss.ContactLeakageNormSquared,
		ActionToRadiusRatio:         seal / ss.FiniteRadiusSquared,
		ActionToGapRatio:            seal / ss.BGap,
		ActionToLeakageRatio:        seal / ss.ContactLeakageNormSquared,
		UnitIndexAvailable:          math.Abs(ss.ContactIndex-1) < 1e-8,
		DimensionlessActionDerived:  true,
		ContinuumIndexBridgeDerived: false,
		CouplingNormalizationOpen:   true,
		DimensionfulUnitDerived:     false,
		GravityScaleDerived:         false,
		ScalarScaleFixed:            false,
		HiddenObservedScaleInserted: false,
		TruthStatement:              "The finite contact index supports a dimensionless topological action normalization, but action normalization is not a mass unit. The scalar scale family v(mu)=mu*r0 remains free until a separate dimensional bridge fixes mu.",
		RemainingUnknowns: []string{
			"U-19A-INDEX-CONTINUUM-BRIDGE: prove that the finite contact index maps to a continuum topological charge without changing normalization",
			"U-19B-COUPLING-NORMALIZATION: derive whether the action seal fixes a gauge coupling, a spectral cutoff, or only a dimensionless path-integral weight",
			"U-19C-DIMENSIONFUL-GRAVITY-BRIDGE: derive a mass/length unit from gravity or a spectral cutoff before comparing to Planck or electroweak scales",
			"U-19D-NO-FITTED-MU: reject choosing mu from v=246 GeV, m_H=125 GeV, or M_Pl unless independently derived",
		},
	}, nil
}
