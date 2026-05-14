package electroweak

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func OperatorSearchTheorem() theorem.Theorem {
	const id = "MATTER-ELECTROWEAK-OPERATOR-SEARCH"
	const name = "hypercharge/chirality operator search on H_Fock⊗H_Φ"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerMatter,
		Status: theorem.BridgeRequired,
		Verify: func() theorem.Result {
			const eps = 1e-8
			a, err := BuildDefault()
			if err != nil {
				return theorem.Result{
					ID:     id,
					Name:   name,
					Layer:  theorem.LayerMatter,
					Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct electroweak operator search", Passed: false, Detail: err.Error()}},
				}
			}
			return theorem.Result{
				ID:     id,
				Name:   name,
				Layer:  theorem.LayerMatter,
				Status: theorem.BridgeRequired,
				Checks: []theorem.Check{
					{Name: "tensor domain", Passed: a.TensorDimension == a.MatterDimension*a.ScalarDimension, Detail: fmt.Sprintf("dim(H_Fock⊗H_Φ)=%d×%d=%d", a.MatterDimension, a.ScalarDimension, a.TensorDimension)},
					{Name: "Fock grading candidate", Passed: a.FockGradingPresent, Detail: "Γ_F|n⟩=(-1)^N|n⟩ from finite occupation parity"},
					{Name: "grading involution", Passed: a.GradingSquareResidual < eps, Detail: fmt.Sprintf("||Γ²−I||_max = %.3e", a.GradingSquareResidual)},
					{Name: "balanced grading trace", Passed: math.Abs(a.GradingTrace) < eps && a.GradingBalanceResidual == 0, Detail: fmt.Sprintf("Tr(Γ)=%.3e, dims(+/-)=%d/%d", a.GradingTrace, a.PositiveGradingDimension, a.NegativeGradingDimension)},
					{Name: "commutes with B-L", Passed: a.CommutesWithBMinusLNorm < eps, Detail: fmt.Sprintf("||[Γ,Q_B-L]||_F=%.3e", a.CommutesWithBMinusLNorm)},
					{Name: "commutes with scalar response", Passed: a.CommutesWithScalarResponseNorm < eps, Detail: fmt.Sprintf("||[Γ,S_Φ]||_F=%.3e", a.CommutesWithScalarResponseNorm)},
					{Name: "charge/parity sector decomposition", Passed: len(a.ChargeParitySectors) > 0, Detail: FormatChargeParitySectors(a.ChargeParitySectors)},
					{Name: "neutral chirality-flipping obstruction", Passed: !a.NeutralChiralityFlippingAvailable, Detail: fmt.Sprintf("dim neutral Γ-flipping maps=%d; Γ-preserving maps=%d; scalar charge/hypercharge is required for EW chirality-changing couplings", a.NeutralChiralityFlippingDimension, a.NeutralChiralityPreservingDimension)},
					{Name: "hypercharge not derived from current data", Passed: !a.HyperchargeDerived && !a.T3ROperatorPresent && !a.ScalarHyperchargePresent, Detail: a.HyperchargeFormula},
					{Name: "electroweak Yukawa discipline", Passed: !a.PhysicalChiralityDerived && !a.ElectroweakYukawaDerived, Detail: "finite grading exists; physical chirality, scalar charge, hypercharge, and EW Yukawa texture remain bridge theorems"},
					{Name: "remaining bridge unknowns", Passed: len(a.RemainingUnknowns) > 0, Detail: FormatUnknowns(a.RemainingUnknowns)},
				},
				Notes: []string{
					"This gate upgrades the vague chirality placeholder into a concrete finite grading candidate Γ_F=(-1)^N.",
					"The obstruction is now precise: B-L-preserving neutral maps cannot flip this finite grading in the current tensor model.",
					"The result is useful but not final: occupation parity is a grading, not yet the full Standard Model left/right chirality operator.",
					"Hypercharge cannot be extracted from B-L and neutral scalar response alone; the missing finite operator is T3_R or an equivalent scalar-charge bridge.",
				},
			}
		},
	}
}
