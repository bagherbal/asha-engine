// Package embedding tests whether the finite Higgs/contact active sector already
// contains enough canonical information to identify the physical 1+3 Fock-mode
// split: one temporal/lepton-like mode and three spatial/color-like modes.
//
// The answer is intentionally conservative. A four-dimensional active scalar
// sector is present, but its current finite spectrum is pair-degenerate. A
// pair-degenerate 2+2 spectral decomposition does not canonically determine a
// 1+3 charge split. Therefore the package records the obstruction and names the
// missing mathematical object: a charge-polarizing operator, such as a finite
// B-L/hypercharge/center action, that is compatible with the contact geometry.
package embedding

import (
	"sync"

	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/matter"
	matteraction "github.com/bagherbal/asha-engine/pkg/matter/action"
)

// Cluster records an eigenvalue degeneracy cluster of the active contact sector.
type Cluster struct {
	Multiplicity int
	Mean         float64
	Spread       float64
}

// Analysis summarizes the canonical embedding search.
type Analysis struct {
	Action matteraction.Analysis
	Bridge matter.FockContactBridge

	ActiveContactDimension int
	FockModeDimension      int
	FockTemporalModes      int
	FockSpatialModes       int

	Clusters []Cluster

	HasFourDimensionalActiveSector bool
	HasFockOnePlusThreeSplit       bool
	SpectrumDeterminesOnePlusThree bool
	CanonicalEmbeddingConstructed  bool

	DegeneracyFreedomDimension int
	RequiredChargeOperator     string
	Obstruction                string
}

var (
	embeddingDefaultOnce  sync.Once
	embeddingDefaultValue Analysis
	embeddingDefaultErr   error
)

func BuildDefault() (Analysis, error) {
	embeddingDefaultOnce.Do(func() {
		embeddingDefaultValue, embeddingDefaultErr = buildEmbeddingDefaultUncached()
	})
	return embeddingDefaultValue, embeddingDefaultErr
}

func buildEmbeddingDefaultUncached() (Analysis, error) {
	a, err := matteraction.BuildDefault()
	if err != nil {
		return Analysis{}, err
	}
	return Build(a, 1e-8)
}

func Build(a matteraction.Analysis, eps float64) (Analysis, error) {
	if eps <= 0 {
		eps = 1e-8
	}
	b := a.Bridge
	active := append([]float64(nil), b.Potential.ActiveContactSpectrum...)
	if len(active) == 0 {
		return Analysis{}, fmt.Errorf("active contact spectrum is empty")
	}
	clusters := clustersFromSpectrum(active, eps)

	// The active Higgs/contact sector has the correct total dimension when it has
	// four real active directions. The Fock side has the desired bookkeeping when
	// the four creation modes split as 1 temporal + 3 spatial/color seeds.
	hasActive4 := b.ActiveHiggsDirections == 4
	hasOnePlusThree := b.TemporalModeCount == 1 && b.SpatialModeCount == 3

	// A spectrum determines a 1+3 split only if its eigenspace multiplicities
	// contain one canonical one-dimensional block and one canonical three-
	// dimensional block. The present 2+2 pair-degenerate spectrum does not.
	spectrumDetermines13 := hasClusterMultiset(clusters, []int{1, 3})

	// Degenerate eigenspaces have internal O(m) basis freedom. The dimension of
	// O(m) is m(m-1)/2. This is the size of the unfixed orientation gauge inside
	// the active scalar eigenspaces.
	freedom := 0
	for _, c := range clusters {
		freedom += c.Multiplicity * (c.Multiplicity - 1) / 2
	}

	constructed := hasActive4 && hasOnePlusThree && spectrumDetermines13
	obstruction := "none"
	if !constructed {
		obstruction = "active sector is four-dimensional, but the current Higgs/contact spectrum splits as 2+2 rather than 1+3; eigenvectors are basis-ambiguous inside degenerate pairs"
	}

	return Analysis{
		Action:                         a,
		Bridge:                         b,
		ActiveContactDimension:         b.ActiveHiggsDirections,
		FockModeDimension:              b.FockModeCount,
		FockTemporalModes:              b.TemporalModeCount,
		FockSpatialModes:               b.SpatialModeCount,
		Clusters:                       clusters,
		HasFourDimensionalActiveSector: hasActive4,
		HasFockOnePlusThreeSplit:       hasOnePlusThree,
		SpectrumDeterminesOnePlusThree: spectrumDetermines13,
		CanonicalEmbeddingConstructed:  constructed,
		DegeneracyFreedomDimension:     freedom,
		RequiredChargeOperator:         "finite charge-polarizing operator compatible with K⊕K⊥, e.g. B-L / hypercharge / central U(1) action",
		Obstruction:                    obstruction,
	}, nil
}

func clustersFromSpectrum(values []float64, eps float64) []Cluster {
	clusters := make([]Cluster, 0)
	used := make([]bool, len(values))
	for i, v := range values {
		if used[i] {
			continue
		}
		used[i] = true
		count := 1
		sum := v
		minV, maxV := v, v
		for j := i + 1; j < len(values); j++ {
			if used[j] {
				continue
			}
			if math.Abs(values[j]-v) <= eps {
				used[j] = true
				count++
				sum += values[j]
				if values[j] < minV {
					minV = values[j]
				}
				if values[j] > maxV {
					maxV = values[j]
				}
			}
		}
		clusters = append(clusters, Cluster{Multiplicity: count, Mean: sum / float64(count), Spread: maxV - minV})
	}
	return clusters
}

func hasClusterMultiset(clusters []Cluster, wanted []int) bool {
	if len(clusters) != len(wanted) {
		return false
	}
	seen := make(map[int]int)
	for _, c := range clusters {
		seen[c.Multiplicity]++
	}
	for _, w := range wanted {
		seen[w]--
	}
	for _, left := range seen {
		if left != 0 {
			return false
		}
	}
	return true
}

func FormatClusters(clusters []Cluster) string {
	out := "["
	for i, c := range clusters {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("m=%d@%.10g", c.Multiplicity, c.Mean)
	}
	return out + "]"
}
