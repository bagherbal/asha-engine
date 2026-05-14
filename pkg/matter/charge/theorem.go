package charge

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ChargePolarizationTheorem() theorem.Theorem {
	return theorem.Theorem{
		ID:     "MATTER-CHARGE-POLARIZATION-BL",
		Name:   "B-L charge-polarizing bridge for the 1+3 Fock split",
		Layer:  theorem.LayerMatter,
		Status: theorem.BridgeRequired,
		Verify: func() theorem.Result {
			a, err := BuildDefault()
			if err != nil {
				return theorem.Result{
					ID:     "MATTER-CHARGE-POLARIZATION-BL",
					Name:   "B-L charge-polarizing bridge for the 1+3 Fock split",
					Layer:  theorem.LayerMatter,
					Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "build charge polarization", Passed: false, Detail: err.Error()}},
				}
			}
			eps := 1e-8
			return theorem.Result{
				ID:     "MATTER-CHARGE-POLARIZATION-BL",
				Name:   "B-L charge-polarizing bridge for the 1+3 Fock split",
				Layer:  theorem.LayerMatter,
				Status: theorem.BridgeRequired,
				Checks: []theorem.Check{
					{Name: "one-particle B-L spectrum", Passed: len(a.OneParticleChargeSpectrum) == 4, Detail: FormatFloatSlice(a.OneParticleChargeSpectrum)},
					{Name: "charge spectrum gives 1+3 split", Passed: a.ChargePolarizesOnePlusThree, Detail: FormatClusters(a.OneParticleChargeClusters)},
					{Name: "traceless one-particle charge", Passed: math.Abs(a.TraceOneParticleCharge) < eps, Detail: fmt.Sprintf("Tr(Q_B-L)=%.3e", a.TraceOneParticleCharge)},
					{Name: "charge normalization invariant", Passed: math.Abs(a.TraceOneParticleChargeSquared-(4.0/3.0)) < eps, Detail: fmt.Sprintf("Tr(Q_B-L²)=%.10f", a.TraceOneParticleChargeSquared)},
					{Name: "sterile vacuum remains neutral", Passed: math.Abs(a.VacuumCharge) < eps, Detail: fmt.Sprintf("Q_B-L(|Ω⟩)=%.3e", a.VacuumCharge)},
					{Name: "commutes with current diagonal Fock response", Passed: a.CommutatorWithFockResponseNorm < eps, Detail: fmt.Sprintf("||[Q_B-L,H_F]||_F = %.3e", a.CommutatorWithFockResponseNorm)},
					{Name: "direct scalar-to-color identification rejected", Passed: !a.DirectScalarToColorIsotropyPossible, Detail: fmt.Sprintf("best spatial scalar anisotropy=%.10g with spatial weights %s", a.BestSpatialScalarAnisotropy, FormatFloatSlice(a.BestSpatialScalarWeights))},
					{Name: "resolved missing object", Passed: a.ChargePolarizesOnePlusThree && !a.DirectScalarToColorIsotropyPossible, Detail: "charge polarization is separate from scalar/Higgs mixing; do not identify Higgs eigenvalues as color charges"},
				},
				Notes: []string{
					a.Resolution,
					a.RemainingUnknown,
					"This gate turns the Gate-15 obstruction into a standard bridge: B-L lives in the Witt/Fock matter sector, while Φ remains the finite scalar/vacuum-mixing sector.",
				},
			}
		},
	}
}
