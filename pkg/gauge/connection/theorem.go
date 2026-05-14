package connection

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ProjectedConnectionTheorem() theorem.Theorem {
	const id = "GAUGE-PROJECTED-CONNECTION-CURVATURE"
	const name = "projected gauge connection and second-fundamental curvature identity"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerGauge,
		Status: theorem.VerifiedNumeric,
		Verify: func() theorem.Result {
			const eps = 1e-8
			a, err := BuildDefault()
			if err != nil {
				return theorem.Result{
					ID: id, Name: name, Layer: theorem.LayerGauge, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct projected connection analysis", Passed: false, Detail: err.Error()}},
				}
			}
			projectorsOK := ValidateProjectionPair(a.Compression.BooleanComplementProjector, a.Compression.BooleanContactProjector, 1e-7) == nil
			identityOK := a.MaxCurvatureIdentityResidual < eps && a.MaxCurvatureIdentityRelative < eps
			reconstructionOK := a.MaxBlockReconstructionResidual < eps
			status := theorem.VerifiedNumeric
			if !identityOK || !reconstructionOK || !projectorsOK {
				status = theorem.OpenTest
			}
			checks := []theorem.Check{
				{
					Name:   "contact/complement projector split",
					Passed: projectorsOK,
					Detail: "P_C and P_K form an orthogonal decomposition of Boolean coordinates",
				},
				{
					Name:   "block reconstruction",
					Passed: reconstructionOK,
					Detail: fmt.Sprintf("max ||A−(PAP+PAQ+QAP+QAQ)||_F = %.3e", a.MaxBlockReconstructionResidual),
				},
				{
					Name:   "off-diagonal second fundamental sector",
					Passed: a.MaxOffDiagonalNorm > eps,
					Detail: fmt.Sprintf("max sqrt(||P_CAQ||²+||QAP_C||²) = %.6e", a.MaxOffDiagonalNorm),
				},
				{
					Name:   "projection is not a Lie homomorphism",
					Passed: a.MaxProjectionDefectNorm > eps,
					Detail: fmt.Sprintf("max ||P[A,B]P−[PAP,PBP]||_F = %.6e", a.MaxProjectionDefectNorm),
				},
				{
					Name:   "second fundamental curvature size",
					Passed: a.MaxSecondFundamentalNorm > eps,
					Detail: fmt.Sprintf("max ||PAQBP−PBQAP||_F = %.6e", a.MaxSecondFundamentalNorm),
				},
				{
					Name:   "Gauss-type projection identity",
					Passed: identityOK,
					Detail: fmt.Sprintf("max residual of P[A,B]P−[PAP,PBP]=PAQBP−PBQAP is %.3e, relative %.3e", a.MaxCurvatureIdentityResidual, a.MaxCurvatureIdentityRelative),
				},
			}
			notes := []string{
				"This resolves the Gate 8/9 problem mathematically: strict compression discarded the off-diagonal connection blocks.",
				"The missing implementation is a finite connection with a second fundamental/Higgs-like sector, not a smaller forced Lie algebra on K⊥ alone.",
				"Future bridge gates should use curvature and covariant derivatives on the K⊕K⊥ vacuum bundle instead of requiring PAP to be a standalone Lie representation.",
			}
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerGauge, Status: status, Checks: checks, Notes: notes}
		},
	}
}
