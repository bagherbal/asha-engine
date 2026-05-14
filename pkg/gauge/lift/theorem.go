package lift

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func BooleanCompressionTheorem() theorem.Theorem {
	const id = "GAUGE-BOOLEAN-LIFT-COMPRESSION"
	const name = "Boolean lift/compression of contact-preserving g₂ᴿ"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerGauge,
		Status: theorem.OpenTest,
		Verify: func() theorem.Result {
			const eps = 1e-8
			c, err := BuildDefault()
			if err != nil {
				return theorem.Result{
					ID: id, Name: name, Layer: theorem.LayerGauge, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct Boolean lift/compression", Passed: false, Detail: err.Error()}},
				}
			}
			strictClosure := c.ClosureRelativeResidual < eps && c.MaxBoundaryLeakage < eps
			status := theorem.OpenTest
			if strictClosure {
				status = theorem.VerifiedNumeric
			}
			checks := []theorem.Check{
				{
					Name:   "input tangent algebra",
					Passed: c.Centralizer.CentralizerDimension == 4 && c.Centralizer.CenterDimension == 1 && c.Centralizer.DerivedDimension == 3,
					Detail: fmt.Sprintf("g₂ᴿ tangent data: dim=%d, center=%d, derived=%d", c.Centralizer.CentralizerDimension, c.Centralizer.CenterDimension, c.Centralizer.DerivedDimension),
				},
				{
					Name:   "exterior lift target",
					Passed: len(c.ExteriorGenerators) == c.Centralizer.CentralizerDimension && c.Contact.G2Support.MiddleDimension() == 70,
					Detail: fmt.Sprintf("lifted %d generators to Λ⁴R⁸ operators of size 70x70", len(c.ExteriorGenerators)),
				},
				{
					Name:   "Boolean compression target",
					Passed: len(c.BooleanGenerators) == c.Centralizer.CentralizerDimension && c.Contact.BooleanSupport.LowerDimension() == 56,
					Detail: fmt.Sprintf("compressed to Boolean coordinates Wᵀρ(X)W of size 56x56"),
				},
				{
					Name:   "contact complement dimension",
					Passed: c.Contact.Dimension() == 7 && c.Contact.BooleanSupport.LowerDimension()-c.Contact.Dimension() == 49,
					Detail: fmt.Sprintf("dim K=%d, dim K⊥ inside Boolean support=%d", c.Contact.Dimension(), c.Contact.BooleanSupport.LowerDimension()-c.Contact.Dimension()),
				},
				{
					Name:   "restricted generator span",
					Passed: c.CompressedFrameRank > 0 && c.CompressedFrameRank <= c.Centralizer.CentralizerDimension,
					Detail: fmt.Sprintf("rank span{P_C Wᵀρ(X)W P_C} = %d", c.CompressedFrameRank),
				},
				{
					Name:   "restricted skew symmetry",
					Passed: c.MaxSkewResidual < eps,
					Detail: fmt.Sprintf("max ||J+Jᵀ||_F after restriction = %.3e", c.MaxSkewResidual),
				},
				{
					Name:   "contact-boundary leakage",
					Passed: !math.IsNaN(c.MaxBoundaryLeakage),
					Detail: fmt.Sprintf("max boundary leakage ||P_KJP_C|| or ||P_CJP_K|| = %.6e", c.MaxBoundaryLeakage),
				},
				{
					Name:   "finite Lie closure residual",
					Passed: !math.IsNaN(c.ClosureRelativeResidual),
					Detail: fmt.Sprintf("max relative residual of [Jᵢ,Jⱼ] projected back to restricted span = %.6e", c.ClosureRelativeResidual),
				},
				{
					Name:   "strict finite gauge theorem",
					Passed: strictClosure,
					Detail: fmt.Sprintf("requires leakage < %.1e and closure residual < %.1e", eps, eps),
				},
			}
			notes := []string{
				"This gate is intentionally harsher than the tangent-level g₂ᴿ centralizer theorem.",
				"A nonzero residual means the naive Boolean compression is not yet a complete finite gauge theorem; it becomes a measured open problem, not a hidden failure.",
			}
			if strictClosure {
				notes = append(notes, "The contact-preserving algebra survives Boolean lift/compression at the configured tolerance.")
			} else {
				notes = append(notes, "Current result: the tangent algebra exists, but the strict Boolean-compressed closure theorem is not yet proven.")
			}
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerGauge, Status: status, Checks: checks, Notes: notes}
		},
	}
}
