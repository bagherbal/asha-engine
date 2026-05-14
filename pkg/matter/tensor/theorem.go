package tensor

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/matter/action"
	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func TensorFactorBridgeTheorem() theorem.Theorem {
	const id = "MATTER-TENSOR-FACTOR-BRIDGE"
	const name = "tensor-factor bridge separating Fock charge from finite scalar mixing"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerMatter,
		Status: theorem.BridgeRequired,
		Verify: func() theorem.Result {
			const eps = 1e-8
			a, err := BuildDefault()
			if err != nil {
				return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct tensor-factor bridge", Passed: false, Detail: err.Error()}},
				}
			}
			return theorem.Result{
				ID:     id,
				Name:   name,
				Layer:  theorem.LayerMatter,
				Status: theorem.BridgeRequired,
				Checks: []theorem.Check{
					{Name: "tensor Hilbert space", Passed: a.TensorDimension == a.MatterDimension*a.ScalarDimension, Detail: fmt.Sprintf("dim(H_Fock⊗H_Φ)=%d×%d=%d", a.MatterDimension, a.ScalarDimension, a.TensorDimension)},
					{Name: "charge acts only on matter factor", Passed: a.MatterChargeOperator.Rows() == a.TensorDimension, Detail: fmt.Sprintf("Q_total=Q_B-L⊗I_Φ is %dx%d", a.MatterChargeOperator.Rows(), a.MatterChargeOperator.Cols())},
					{Name: "scalar response acts only on Higgs/contact factor", Passed: a.ScalarResponse.Rows() == a.TensorDimension, Detail: fmt.Sprintf("S_total=I_Fock⊗S_Φ with weights %s", action.FormatFloatSlice(a.ScalarOneParticleWeights))},
					{Name: "charge/scalar independence", Passed: a.ChargeScalarCommutatorNorm < eps, Detail: fmt.Sprintf("||[Q_total,S_total]||_F = %.3e", a.ChargeScalarCommutatorNorm)},
					{Name: "charge/matter response compatibility", Passed: a.ChargeMatterCommutatorNorm < eps, Detail: fmt.Sprintf("||[Q_total,H_F⊗I_Φ]||_F = %.3e", a.ChargeMatterCommutatorNorm)},
					{Name: "matter/scalar tensor compatibility", Passed: a.MatterScalarCommutatorNorm < eps, Detail: fmt.Sprintf("||[H_F⊗I_Φ,I_Fock⊗S_Φ]||_F = %.3e", a.MatterScalarCommutatorNorm)},
					{Name: "scalar trace tensor identity", Passed: a.TensorScalarTraceResidual < eps, Detail: fmt.Sprintf("Tr(I⊗S)=%.10f, dim(Fock)Tr(S)=%.10f", a.TensorScalarTrace, a.ExpectedTensorScalarTrace)},
					{Name: "charge trace tensor identity", Passed: a.TensorChargeTraceResidual < eps && math.Abs(a.TensorChargeTrace) < eps, Detail: fmt.Sprintf("Tr(Q⊗I)=%.3e, expected %.3e", a.TensorChargeTrace, a.ExpectedTensorChargeTrace)},
					{Name: "neutral vacuum scalar fiber", Passed: math.Abs(a.VacuumFiberCharge) < eps && a.VacuumScalarFiberDimension == a.ScalarDimension, Detail: fmt.Sprintf("|Ω⟩⊗H_Φ has charge %.3e and scalar fiber dimension %d", a.VacuumFiberCharge, a.VacuumScalarFiberDimension)},
					{Name: "charge-sector scalar fibers", Passed: len(a.ChargeFibers) > 0, Detail: FormatFibers(a.ChargeFibers)},
					{Name: "direct eigenvalue collapse rejected", Passed: a.DirectEigenvalueIdentificationRejected, Detail: "scalar/contact eigenvalues remain a separate factor; they are not color charges"},
					{Name: "Yukawa discipline", Passed: !a.YukawaIntertwinerConstructed && !a.SelectionRuleFormulated, Detail: "OPEN U-07: next gate must construct an intertwiner/selection rule, not a diagonal lookup"},
				},
				Notes: []string{
					"This gate resolves the Gate-15/16 tension by using the standard tensor product H_total=H_Fock⊗H_Φ.",
					"B-L supplies the 1+3 matter charge polarization; the finite Higgs/contact spectrum supplies the scalar response on a separate factor.",
					"The next missing object is a gauge-compatible Yukawa/intertwiner map between matter states and the scalar factor.",
				},
			}
		},
	}
}
