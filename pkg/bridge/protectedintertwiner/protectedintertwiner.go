// Package protectedintertwiner implements Gate 87: protected-contact / broken-
// generator intertwiner search.
//
// Gate 48 established a count-level resonance: the scalar/contact doublet has
// three angular directions, the electroweak breaking pattern has three broken
// generator directions, and the finite contact vacuum has three protected
// unmixed directions.  Gates 84-86 strengthened the broken-generator side and
// the scalar-vacuum side.
//
// This gate asks the stricter question: is there a canonical isometry or
// intertwiner from the three protected contact directions to the three broken
// generator image directions?  A true gauge-eating theorem needs such a map.
// A mere dimension match leaves an O(3) family of possible identifications and
// is therefore only a bridge diagnostic.
package protectedintertwiner

import (
	"fmt"
	"math"
	"sync"

	"github.com/bagherbal/asha-engine/pkg/bridge/gaugeeating"
	"github.com/bagherbal/asha-engine/pkg/bridge/scalarvacuum"
)

type Analysis struct {
	ScalarVacuum scalarvacuum.Analysis
	GaugeEating  gaugeeating.Analysis

	ProtectedContactDirections  int
	ScalarAngularDirections     int
	BrokenGeneratorDirections   int
	BrokenImageRank             int
	UnbrokenGeneratorDirections int

	CountLevelResonance             bool
	BrokenImageMetricPositive       bool
	BrokenImageMinEigen             float64
	BrokenImageMaxEigen             float64
	BrokenImageCondition            float64
	BrokenImageMetricAvailable      bool
	ProtectedContactMetricAvailable bool

	AbstractIsometryExists           bool
	AbstractIsometryGroup            string
	AbstractIsometryFreedomDimension int
	CanonicalIntertwinerDerived      bool
	ProtectedToBrokenMapDerived      bool
	GaugeEatingTheoremDerived        bool

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
		sv, err := scalarvacuum.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		ge, err := gaugeeating.BuildDefault()
		if err != nil {
			defaultErr = err
			return
		}
		defaultValue, defaultErr = Build(sv, ge, 1e-10)
	})
	return defaultValue, defaultErr
}

func Build(sv scalarvacuum.Analysis, ge gaugeeating.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-10
	}
	protected := sv.ScalarPotential.ProtectedDirectionCount
	angular := sv.ActiveRealDimension - 1
	broken := ge.BrokenGeneratorCount
	if protected <= 0 || angular <= 0 || broken <= 0 {
		return Analysis{}, fmt.Errorf("protected intertwiner search requires positive protected/angular/broken dimensions")
	}
	count := protected == 3 && angular == 3 && broken == 3 && ge.BrokenImageRank == 3
	min, max := ge.BrokenImageMinEigen, ge.BrokenImageMaxEigen
	cond := ge.BrokenImageCondition
	positive := min > eps && max > eps && !math.IsInf(cond, 1) && !math.IsNaN(cond)

	truth := "The protected-contact and broken-generator sectors match at dimension three, and the broken-generator images form a positive three-dimensional metric. This is exactly the count-level shape needed for gauge eating. However, the engine still has no canonical map from the protected contact frame into the broken-generator image frame. An abstract isometry exists, but without a finite contact/BF action selecting it, the identification carries an O(3) freedom and remains a bridge-level diagnostic."

	return Analysis{
		ScalarVacuum:                     sv,
		GaugeEating:                      ge,
		ProtectedContactDirections:       protected,
		ScalarAngularDirections:          angular,
		BrokenGeneratorDirections:        broken,
		BrokenImageRank:                  ge.BrokenImageRank,
		UnbrokenGeneratorDirections:      ge.UnbrokenGeneratorCount,
		CountLevelResonance:              count,
		BrokenImageMetricPositive:        positive,
		BrokenImageMinEigen:              min,
		BrokenImageMaxEigen:              max,
		BrokenImageCondition:             cond,
		BrokenImageMetricAvailable:       ge.GoldstoneImageTheoremDiagnostic,
		ProtectedContactMetricAvailable:  false,
		AbstractIsometryExists:           count && positive,
		AbstractIsometryGroup:            "O(3)",
		AbstractIsometryFreedomDimension: 3,
		CanonicalIntertwinerDerived:      false,
		ProtectedToBrokenMapDerived:      false,
		GaugeEatingTheoremDerived:        false,
		TruthStatement:                   truth,
		RecommendedNextGate:              "Gate 88 — Protected-Contact Metric / Connection Form Search",
		RemainingUnknowns: []string{
			"U-18C7A-PROTECTED-CONTACT-METRIC: derive a metric/connection form on the three protected contact directions, not only their dimension",
			"U-18C7B-CANONICAL-O3-REDUCTION: derive an action term that selects one O(3) isometry from protected contact space to broken-generator images",
			"U-18C7C-CONTACT-BF-INTERTWINER: compute the protected-to-broken map from finite BF/contact curvature or prove the O(3) freedom is pure gauge",
			"U-19C1-SCALAR-KINETIC-ACTION: derive scalar kinetic normalization before interpreting the gauge-eating diagnostic as a theorem",
		},
	}, nil
}

func FormatFloat(x float64) string { return fmt.Sprintf("%.10f", x) }
