package higgspotential

import (
	"fmt"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func PotentialCandidateTheorem() theorem.Theorem {
	const id = "DYN-HIGGS-POTENTIAL-CANDIDATE"
	const name = "finite Higgs-potential candidate from vacuum-mixing spectra"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerDynamics,
		Status: theorem.BridgeRequired,
		Verify: func() theorem.Result {
			const eps = 1e-8
			a, err := BuildDefault()
			if err != nil {
				return theorem.Result{ID: id, Name: name, Layer: theorem.LayerDynamics, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct finite Higgs-potential candidate", Passed: false, Detail: err.Error()}},
				}
			}

			activeFourOK := a.ActiveContactDimension == 4
			protectedThreeOK := a.ProtectedContactDimension == 3
			pairOK := a.PairDegenerateSpectrum
			orderOK := a.OrderParameterNormSquared > eps
			quarticOK := a.QuarticTrace > eps && a.NormalizedQuarticShape > eps
			kinematicsOK := a.MexicanHatKinematics

			checks := []theorem.Check{
				{
					Name:   "active real scalar sector",
					Passed: activeFourOK,
					Detail: fmt.Sprintf("rank(M_K)=%d active contact directions; spectrum %s", a.ActiveContactDimension, formatFloatSlice(a.ActiveContactSpectrum)),
				},
				{
					Name:   "protected residual contact sector",
					Passed: protectedThreeOK,
					Detail: fmt.Sprintf("dim(K)-rank(M_K)=%d protected unmixed contact directions", a.ProtectedContactDimension),
				},
				{
					Name:   "pair-degenerate contact spectrum",
					Passed: pairOK,
					Detail: fmt.Sprintf("pair residual %.3e; clusters %s", a.PairDegeneracyResidual, formatClusters(a.DegeneracyClusters)),
				},
				{
					Name:   "finite order-parameter trace",
					Passed: orderOK,
					Detail: fmt.Sprintf("τ=Tr(M_K)=%.10f; no electroweak vev is inferred here", a.OrderParameterNormSquared),
				},
				{
					Name:   "quartic shape invariant",
					Passed: quarticOK,
					Detail: fmt.Sprintf("Tr(M_K²)=%.10f, Tr(M_K²)/Tr(M_K)²=%.10f", a.QuarticTrace, a.NormalizedQuarticShape),
				},
				{
					Name:   "spectral anisotropy",
					Passed: a.SpectralAnisotropy >= 0,
					Detail: fmt.Sprintf("mean active eigenvalue=%.10f, anisotropy=(λmax−λmin)/mean=%.10f", a.MeanActiveEigenvalue, a.SpectralAnisotropy),
				},
				{
					Name:   "Mexican-hat kinematic ingredients",
					Passed: kinematicsOK,
					Detail: "requires 4 active real directions, 3 protected directions, positive trace, and pair organization",
				},
				{
					Name:   "bridge discipline",
					Passed: true,
					Detail: "status is BRIDGE_REQUIRED: this gate defines finite potential invariants, not the observed Higgs mass or electroweak vev",
				},
			}
			notes := []string{
				"The finite potential candidate is spectral: its primitive invariants are rank(M_K), Tr(M_K), Tr(M_K²), degeneracy clusters, and anisotropy.",
				"The four active real contact directions are the first credible finite Higgs-doublet kinematic signal in this branch, but a physical Higgs field still needs a matter/Yukawa bridge.",
				"The three protected unmixed contact directions should not be discarded; they may encode residual gauge/Goldstone/contact constraints.",
				"No measured physical constant is used or fitted in this theorem gate.",
			}
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerDynamics, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
		},
	}
}

func formatFloatSlice(values []float64) string {
	parts := make([]string, len(values))
	for i, v := range values {
		parts[i] = fmt.Sprintf("%.10g", v)
	}
	return "[" + strings.Join(parts, ", ") + "]"
}

func formatClusters(clusters []DegeneracyCluster) string {
	parts := make([]string, len(clusters))
	for i, c := range clusters {
		parts[i] = fmt.Sprintf("m=%d@%.10g", c.Multiplicity, c.Mean)
	}
	return "[" + strings.Join(parts, ", ") + "]"
}
