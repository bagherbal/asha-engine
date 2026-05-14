package embedding

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func CanonicalEmbeddingTheorem() theorem.Theorem {
	return theorem.Theorem{
		ID:     "MATTER-CANONICAL-EMBEDDING-SEARCH",
		Name:   "canonical Fock/contact charge-embedding search",
		Layer:  theorem.LayerMatter,
		Status: theorem.OpenTest,
		Verify: func() theorem.Result {
			a, err := BuildDefault()
			if err != nil {
				return theorem.Result{
					ID:     "MATTER-CANONICAL-EMBEDDING-SEARCH",
					Name:   "canonical Fock/contact charge-embedding search",
					Layer:  theorem.LayerMatter,
					Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "build analysis", Passed: false, Detail: err.Error()}},
				}
			}
			return theorem.Result{
				ID:     "MATTER-CANONICAL-EMBEDDING-SEARCH",
				Name:   "canonical Fock/contact charge-embedding search",
				Layer:  theorem.LayerMatter,
				Status: theorem.OpenTest,
				Checks: []theorem.Check{
					{Name: "four-dimensional active contact sector", Passed: a.HasFourDimensionalActiveSector, Detail: fmt.Sprintf("active contact directions=%d", a.ActiveContactDimension)},
					{Name: "Fock 1+3 mode split", Passed: a.HasFockOnePlusThreeSplit, Detail: fmt.Sprintf("temporal=%d, spatial/color seeds=%d", a.FockTemporalModes, a.FockSpatialModes)},
					{Name: "active spectral clusters", Passed: len(a.Clusters) > 0, Detail: FormatClusters(a.Clusters)},
					{Name: "spectrum determines 1+3 embedding", Passed: a.SpectrumDeterminesOnePlusThree, Detail: "requires spectral multiplicities {1,3}; current multiplicities do not fix lepton/color orientation"},
					{Name: "canonical embedding constructed", Passed: a.CanonicalEmbeddingConstructed, Detail: "requires an additional charge-polarizing operator before assigning eigenvectors to Fock modes"},
					{Name: "unfixed degeneracy freedom", Passed: a.DegeneracyFreedomDimension > 0, Detail: fmt.Sprintf("active eigenspaces retain O(2)×O(2) freedom; dimension=%d", a.DegeneracyFreedomDimension)},
				},
				Notes: []string{
					"This is an honest obstruction, not a numerical failure: a pair-degenerate 2+2 scalar spectrum cannot canonically produce a 1+3 charge split.",
					"Missing object: " + a.RequiredChargeOperator + ".",
					"Obstruction: " + a.Obstruction + ".",
				},
			}
		},
	}
}
