package contact

import (
	"fmt"
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactSpaceTheorem() theorem.Theorem {
	const id = "GEO-CONTACT-K-BG-L4-R8"
	const name = "Boolean–Octonionic contact space K inside Λ⁴R⁸"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerGeometry,
		Status: theorem.ExactFinite,
		Verify: func() theorem.Result {
			const eps = 1e-8
			space, err := BuildDefault()
			if err != nil {
				return theorem.Result{
					ID: id, Name: name, Layer: theorem.LayerGeometry, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct contact space", Passed: false, Detail: err.Error()}},
				}
			}

			projectorResidual, _ := space.ContactProjector.IdempotenceResidual()
			symmetryResidual, _ := space.ContactProjector.SymmetryResidual()
			trace, _ := space.ContactProjector.Trace()
			frameResidual, _ := space.FrameIsometryResidual()
			booleanResidual, _ := space.BooleanContainmentResidual()
			g2Residual, _ := space.G2ContainmentResidual()
			denom := space.ExpectedContactDenominator()
			index := space.ContactIndex()
			contactDim := space.Dimension()

			checks := []theorem.Check{
				{
					Name:   "contact overlap spectrum",
					Passed: contactDim == denom && space.OverlapMultiplicityNear(1, eps) == denom,
					Detail: fmt.Sprintf("Q_GᵀP_BQ_G has %d eigenvalues near 1 inside a %dD G₂ sector", space.OverlapMultiplicityNear(1, eps), space.G2Support.SectorDimension()),
				},
				{
					Name:   "contact dimension",
					Passed: contactDim == denom,
					Detail: fmt.Sprintf("dim K=%d, expected one G₂ contact copy=%d", contactDim, denom),
				},
				{
					Name:   "orthonormal contact frame",
					Passed: frameResidual < eps,
					Detail: fmt.Sprintf("||Q_KᵀQ_K−I||_F = %.3e", frameResidual),
				},
				{
					Name:   "contact projector",
					Passed: projectorResidual < eps && symmetryResidual < eps,
					Detail: fmt.Sprintf("||P_K²−P_K||_F = %.3e, ||P_K−P_Kᵀ||_F = %.3e", projectorResidual, symmetryResidual),
				},
				{
					Name:   "projector trace rank",
					Passed: math.Abs(trace-float64(contactDim)) < eps,
					Detail: fmt.Sprintf("Tr(P_K)=%.10f", trace),
				},
				{
					Name:   "Boolean containment",
					Passed: booleanResidual < eps,
					Detail: fmt.Sprintf("||P_BP_K−P_K||_F = %.3e", booleanResidual),
				},
				{
					Name:   "G₂ containment",
					Passed: g2Residual < eps,
					Detail: fmt.Sprintf("||P_GP_K−P_K||_F = %.3e", g2Residual),
				},
				{
					Name:   "finite contact index",
					Passed: math.Abs(index-1) < eps,
					Detail: fmt.Sprintf("I_BG = dim(K)/%d = %.10f", denom, index),
				},
				{
					Name:   "bare contact leakage invariant",
					Passed: space.BareLeakageNorm() > 0 && !math.IsNaN(space.BareLeakageNorm()),
					Detail: fmt.Sprintf("L_BG=||P_BP_G−P_K||_F = %.10f; this is not Λ", space.BareLeakageNorm()),
				},
			}

			return theorem.Result{
				ID:     id,
				Name:   name,
				Layer:  theorem.LayerGeometry,
				Status: theorem.ExactFinite,
				Checks: checks,
				Notes: []string{
					"This gate computes K as the eigenvalue-1 intersection of the Boolean support and the G₂ calibration sector.",
					"The leakage invariant is intentionally named L_BG; it is a finite projector residual, not the observed cosmological constant.",
					fmt.Sprintf("overlap spectrum signature: %s", compactSpectrum(space.OverlapEigenvalues, eps)),
				},
			}
		},
	}
}

func compactSpectrum(values []float64, eps float64) string {
	buckets := make([]string, 0)
	used := make([]bool, len(values))
	for i, value := range values {
		if used[i] {
			continue
		}
		count := 1
		used[i] = true
		for j := i + 1; j < len(values); j++ {
			if math.Abs(values[j]-value) < eps {
				used[j] = true
				count++
			}
		}
		buckets = append(buckets, fmt.Sprintf("%.10g×%d", value, count))
	}
	return strings.Join(buckets, ", ")
}
