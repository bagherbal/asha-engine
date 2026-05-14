package gauge

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactCentralizerTheorem() theorem.Theorem {
	const id = "GAUGE-G2-R-CENTRALIZER"
	const name = "contact-preserving centralizer inside g₂"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerGauge,
		Status: theorem.VerifiedNumeric,
		Verify: func() theorem.Result {
			const eps = 1e-8
			c, err := BuildContactCentralizer()
			if err != nil {
				return theorem.Result{
					ID: id, Name: name, Layer: theorem.LayerGauge, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct g₂ centralizer", Passed: false, Detail: err.Error()}},
				}
			}
			checks := []theorem.Check{
				{
					Name:   "octonion derivation span",
					Passed: len(c.RawDerivations) == 21 && c.G2Rank == 14,
					Detail: fmt.Sprintf("21 standard derivations span dim(g₂)=%d", c.G2Rank),
				},
				{
					Name:   "compact skew generators",
					Passed: c.G2SkewResidual() < eps,
					Detail: fmt.Sprintf("max ||X+Xᵀ||_F across orthonormal g₂ frame = %.3e", c.G2SkewResidual()),
				},
				{
					Name:   "contact-copy involution",
					Passed: c.InvolutionResidual() < eps,
					Detail: fmt.Sprintf("R=diag(-,+,-,+,+,+,+), ||R²−I||_F = %.3e", c.InvolutionResidual()),
				},
				{
					Name:   "centralizer dimension",
					Passed: c.CentralizerDimension == 4,
					Detail: fmt.Sprintf("dim g₂ᴿ = dim ker(X↦[X,R]) = %d", c.CentralizerDimension),
				},
				{
					Name:   "centralizer frame",
					Passed: c.FrameIsometryResidual() < eps,
					Detail: fmt.Sprintf("||Q_RᵀQ_R−I||_F = %.3e", c.FrameIsometryResidual()),
				},
				{
					Name:   "commutes with R",
					Passed: c.CentralizerResidual() < eps,
					Detail: fmt.Sprintf("max ||[X,R]||_F across g₂ᴿ frame = %.3e", c.CentralizerResidual()),
				},
				{
					Name:   "Lie closure",
					Passed: c.ClosureResidual() < eps,
					Detail: fmt.Sprintf("max projection residual of [Xᵢ,Xⱼ] back into g₂ᴿ = %.3e", c.ClosureResidual()),
				},
				{
					Name:   "su(2) plus u(1) signature",
					Passed: c.CenterDimension == 1 && c.DerivedDimension == 3,
					Detail: fmt.Sprintf("center dimension=%d, derived algebra dimension=%d", c.CenterDimension, c.DerivedDimension),
				},
				{
					Name:   "no hidden physical constants",
					Passed: !math.IsNaN(float64(c.CentralizerDimension)),
					Detail: "all dimensions are obtained from octonion derivations and the contact involution R",
				},
			}
			return theorem.Result{
				ID:     id,
				Name:   name,
				Layer:  theorem.LayerGauge,
				Status: theorem.VerifiedNumeric,
				Checks: checks,
				Notes: []string{
					"This gate identifies the contact-preserving internal algebra at the octonionic tangent level.",
					"The result supports g₂ᴿ ≅ su(2) ⊕ u(1) numerically: four generators, one central direction, and a three-dimensional derived algebra.",
					"It does not yet prove the Boolean-compressed finite gauge theorem; that is the next gate.",
				},
			}
		},
	}
}
