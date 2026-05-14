package bsector

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactVacuumTheorem() theorem.Theorem {
	const id = "DYN-BSECTOR-CONTACT-VACUUM"
	const name = "B-sector quadratic action has K as its exact zero-energy sector"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerDynamics,
		Status: theorem.Variational,
		Verify: func() theorem.Result {
			const eps = 1e-8
			vacuum, err := BuildDefault()
			if err != nil {
				return theorem.Result{
					ID: id, Name: name, Layer: theorem.LayerDynamics, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct B-sector vacuum", Passed: false, Detail: err.Error()}},
				}
			}

			symResidual, _ := vacuum.OperatorSymmetryResidual()
			zeroActionResidual, _ := vacuum.ZeroModeActionResidual()
			contactActionResidual, _ := vacuum.ContactActionResidual()
			kernelContactResidual, _ := vacuum.KernelEqualsContactResidual()
			contactFrameResidual, _ := vacuum.ContactFrameIsometryResidual()
			zeroProjectorResidual, _ := vacuum.ZeroModeProjector.IdempotenceResidual()
			contactProjectorResidual, _ := vacuum.ContactProjector.IdempotenceResidual()
			zeroTrace, _ := vacuum.ZeroModeProjector.Trace()
			gap := vacuum.FirstPositiveEigenvalue(eps)
			zeroDim := vacuum.ZeroModeDimension(eps)
			contactDim := vacuum.Contact.Dimension()
			negativeCount := vacuum.NegativeEigenvalueCount(eps)

			checks := []theorem.Check{
				{
					Name:   "Boolean-coordinate action operator",
					Passed: vacuum.BooleanDimension() == vacuum.Contact.BooleanSupport.LowerDimension(),
					Detail: fmt.Sprintf("O_B = Wᵀ(I−P_G)W is %dx%d on Boolean coordinates", vacuum.Operator.Rows(), vacuum.Operator.Cols()),
				},
				{
					Name:   "operator symmetry",
					Passed: symResidual < eps,
					Detail: fmt.Sprintf("||O_B−O_Bᵀ||_F = %.3e", symResidual),
				},
				{
					Name:   "positive semidefinite spectrum",
					Passed: negativeCount == 0,
					Detail: fmt.Sprintf("negative eigenvalues below −eps: %d", negativeCount),
				},
				{
					Name:   "zero-mode multiplicity",
					Passed: zeroDim == contactDim,
					Detail: fmt.Sprintf("dim ker(O_B)=%d, dim K=%d", zeroDim, contactDim),
				},
				{
					Name:   "zero modes have zero action",
					Passed: zeroActionResidual < eps,
					Detail: fmt.Sprintf("||O_B Q_0||_F = %.3e", zeroActionResidual),
				},
				{
					Name:   "contact frame has zero action",
					Passed: contactActionResidual < eps,
					Detail: fmt.Sprintf("||O_B WᵀQ_K||_F = %.3e", contactActionResidual),
				},
				{
					Name:   "contact frame remains orthonormal in Boolean coordinates",
					Passed: contactFrameResidual < eps,
					Detail: fmt.Sprintf("||(WᵀQ_K)ᵀ(WᵀQ_K)−I||_F = %.3e", contactFrameResidual),
				},
				{
					Name:   "kernel equals contact projector",
					Passed: kernelContactResidual < eps,
					Detail: fmt.Sprintf("||P_ker(O_B)−P_K^Boolean||_F = %.3e", kernelContactResidual),
				},
				{
					Name:   "projector hygiene",
					Passed: zeroProjectorResidual < eps && contactProjectorResidual < eps && math.Abs(zeroTrace-float64(contactDim)) < eps,
					Detail: fmt.Sprintf("||P_0²−P_0||_F = %.3e, ||P_KB²−P_KB||_F = %.3e, Tr(P_0)=%.10f", zeroProjectorResidual, contactProjectorResidual, zeroTrace),
				},
				{
					Name:   "spectral gap above contact vacuum",
					Passed: !math.IsNaN(gap) && gap > eps,
					Detail: fmt.Sprintf("first positive eigenvalue of O_B = %.10f", gap),
				},
			}

			return theorem.Result{
				ID:     id,
				Name:   name,
				Layer:  theorem.LayerDynamics,
				Status: theorem.Variational,
				Checks: checks,
				Notes: []string{
					"This gate upgrades K from a static intersection to the vacuum of a finite quadratic action.",
					"The action is evaluated on Boolean fields B=Wb: S_B[b]=||(I−P_G)Wb||²=bᵀO_Bb.",
					"No physical mass, Λ, or coupling is inferred from this gate; it proves the finite zero-energy sector only.",
					fmt.Sprintf("low spectrum signature: %s", lowSpectrum(vacuum.Eigenvalues, 12)),
				},
			}
		},
	}
}

func lowSpectrum(values []float64, n int) string {
	if n > len(values) {
		n = len(values)
	}
	out := ""
	for i := 0; i < n; i++ {
		if i > 0 {
			out += ", "
		}
		out += fmt.Sprintf("%.10g", values[i])
	}
	return out
}
