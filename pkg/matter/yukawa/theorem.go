package yukawa

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func IntertwinerSelectionTheorem() theorem.Theorem {
	const id = "MATTER-YUKAWA-INTERTWINER-SELECTION"
	const name = "Yukawa/intertwiner charge-selection rule on H_Fock⊗H_Φ"
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
					Checks: []theorem.Check{{Name: "construct Yukawa/intertwiner selection gate", Passed: false, Detail: err.Error()}},
				}
			}
			return theorem.Result{
				ID:     id,
				Name:   name,
				Layer:  theorem.LayerMatter,
				Status: theorem.BridgeRequired,
				Checks: []theorem.Check{
					{Name: "tensor domain", Passed: a.TensorDimension == a.MatterDimension*a.ScalarDimension, Detail: fmt.Sprintf("dim(H_Fock⊗H_Φ)=%d×%d=%d", a.MatterDimension, a.ScalarDimension, a.TensorDimension)},
					{Name: "neutral selection rule formulated", Passed: a.YukawaSelectionRuleFormulated, Detail: "neutral intertwiners satisfy [Q_B-L⊗I_Φ, Y]=0"},
					{Name: "charge-sector decomposition", Passed: len(a.ChargeSectors) > 0, Detail: FormatChargeSectors(a.ChargeSectors)},
					{Name: "charge-preserving linear space", Passed: a.ChargePreservingDimension > 0, Detail: fmt.Sprintf("dim{Y:[Q,Y]=0}=%d inside full End(H) dim=%d", a.ChargePreservingDimension, a.FullLinearDimension)},
					{Name: "charge-changing maps rejected for neutral scalar", Passed: a.ChargeChangingDimension > 0 && a.NeutralScalarOnly, Detail: fmt.Sprintf("neutral B-L scalar forbids %d charge-changing entries unless scalar charge/hypercharge is added", a.ChargeChangingDimension)},
					{Name: "neutral selection fraction", Passed: a.NeutralSelectionFraction > 0 && a.NeutralSelectionFraction < 1, Detail: fmt.Sprintf("allowed neutral fraction %.6f", a.NeutralSelectionFraction)},
					{Name: "one-particle selection space", Passed: a.OneParticleChargePreservingDimension > 0, Detail: fmt.Sprintf("one-particle charge-preserving dim=%d inside %dx%d linear space", a.OneParticleChargePreservingDimension, a.OneParticleTotalDimension, a.OneParticleTotalDimension)},
					{Name: "parity/chirality placeholder balance", Passed: a.ParityBalanceResidual == 0, Detail: fmt.Sprintf("parity-preserving dim=%d, parity-flipping dim=%d; true chirality operator still open", a.ParityPreservingDimension, a.ParityFlippingDimension)},
					{Name: "sample neutral witness", Passed: a.ChargeRuleResidual < eps, Detail: fmt.Sprintf("using S_total as a neutral witness gives ||[Q,Y]||_F=%.3e", a.ChargeRuleResidual)},
					{Name: "physical Yukawa texture discipline", Passed: !a.PhysicalYukawaTextureDerived && !a.MassMatrixDerived, Detail: "selection rule exists; no SM Yukawa texture, fermion masses, or Higgs mass are claimed"},
					{Name: "remaining bridge unknowns", Passed: len(a.RemainingUnknowns) > 0, Detail: FormatUnknowns(a.RemainingUnknowns)},
				},
				Notes: []string{
					"This gate turns U-07 from a vague missing mass matrix into a precise selection-rule problem.",
					"With the current neutral scalar factor, Yukawa/intertwiner maps must preserve B-L charge blocks. Charge-changing electroweak couplings require a scalar charge/hypercharge bridge.",
					"The next missing theorem is not numeric fitting; it is constructing the correct finite hypercharge/chirality operator on the tensor product.",
				},
			}
		},
	}
}
