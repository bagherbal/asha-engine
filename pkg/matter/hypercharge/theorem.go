package hypercharge

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ScalarHyperchargeBridgeTheorem() theorem.Theorem {
	const id = "MATTER-SCALAR-HYPERCHARGE-BRIDGE"
	const name = "T3_R / scalar-hypercharge bridge search"
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
					Checks: []theorem.Check{{Name: "construct scalar hypercharge bridge", Passed: false, Detail: err.Error()}},
				}
			}
			return theorem.Result{
				ID:     id,
				Name:   name,
				Layer:  theorem.LayerMatter,
				Status: theorem.BridgeRequired,
				Checks: []theorem.Check{
					{Name: "active scalar doublet support", Passed: a.ScalarDimension == 4 && a.ScalarDoubletCandidate, Detail: fmt.Sprintf("active scalar/contact dimension=%d with degeneracy clusters %s", a.ScalarDimension, FormatIntSlice(a.ScalarChargeClusters))},
					{Name: "fundamental weak weight derived", Passed: math.Abs(a.FundamentalWeight-0.5) < eps, Detail: fmt.Sprintf("weight=1/pairMultiplicity=1/%d=%.10f", a.PairMultiplicity, a.FundamentalWeight)},
					{Name: "scalar T3 candidate", Passed: a.ScalarChargeBridgeConstructed, Detail: fmt.Sprintf("T_Φ spectrum %s", FormatFloatSlice(a.ScalarChargeSpectrum))},
					{Name: "trace-zero scalar charge", Passed: math.Abs(a.ScalarChargeTrace) < eps, Detail: fmt.Sprintf("Tr(T_Φ)=%.3e, Tr(T_Φ²)=%.10f", a.ScalarChargeTrace, a.ScalarChargeTrace2)},
					{Name: "compatible with pair-degenerate scalar response", Passed: a.CommutesWithScalarResponseNorm < eps, Detail: fmt.Sprintf("||[T_Φ,S_Φ]||_F=%.3e", a.CommutesWithScalarResponseNorm)},
					{Name: "independent from B-L matter charge", Passed: a.CommutesWithBMinusLNorm < eps, Detail: fmt.Sprintf("||[I⊗T_Φ,Q_B-L⊗I]||_F=%.3e", a.CommutesWithBMinusLNorm)},
					{Name: "independent from current grading", Passed: a.CommutesWithFockGradingNorm < eps, Detail: fmt.Sprintf("||[I⊗T_Φ,Γ_F⊗I]||_F=%.3e", a.CommutesWithFockGradingNorm)},
					{Name: "Pati-Salam template diagnostic", Passed: a.PatiSalamFlippingDim == 0, Detail: fmt.Sprintf("Q_PS=(B-L)/2+T_Φ sectors %s; Γ-flipping dim=%d, preserving dim=%d", FormatSectors(a.PatiSalamSectors), a.PatiSalamFlippingDim, a.PatiSalamPreservingDim)},
					{Name: "raw compensator diagnostic", Passed: a.RawFlippingDim > 0, Detail: fmt.Sprintf("Q_raw=(B-L)+T_Φ has Γ-flipping dim=%d; diagnostic only, not SM hypercharge", a.RawFlippingDim)},
					{Name: "matter T3_R not yet constructed", Passed: !a.MatterT3ROperatorConstructed && !a.StandardModelHyperchargeDerived, Detail: "scalar doublet charge exists, but matter-side T3_R / physical hypercharge is still open"},
					{Name: "physical chirality discipline", Passed: !a.PhysicalChiralityDerived && !a.ElectroweakYukawaDerived, Detail: "Γ_F remains a grading candidate; no electroweak Yukawa or mass texture is claimed"},
					{Name: "remaining bridge unknowns", Passed: len(a.RemainingUnknowns) > 0, Detail: FormatUnknowns(a.RemainingUnknowns)},
				},
				Notes: []string{
					"This gate extracts the part the finite scalar/contact spectrum really gives: a 2+2 real scalar doublet with a canonical trace-zero T_Φ charge.",
					"The Pati-Salam-shaped template still does not make Γ_F a physical chirality operator. That is valuable: occupation parity is not enough.",
					"The next missing object is matter-side T3_R, or an equivalent finite operator that supplies physical left/right hypercharge before Yukawa textures are attempted.",
				},
			}
		},
	}
}
