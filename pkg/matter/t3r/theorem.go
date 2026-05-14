package t3r

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func MatterT3RSearchTheorem() theorem.Theorem {
	const id = "MATTER-T3R-CHIRALITY-SEARCH"
	const name = "matter-side T3_R and physical chirality search"
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
					Checks: []theorem.Check{{Name: "construct matter T3_R search", Passed: false, Detail: err.Error()}},
				}
			}
			return theorem.Result{
				ID:     id,
				Name:   name,
				Layer:  theorem.LayerMatter,
				Status: theorem.BridgeRequired,
				Checks: []theorem.Check{
					{Name: "tensor domain", Passed: a.TensorDimension == a.MatterDimension*a.ScalarDimension, Detail: fmt.Sprintf("dim(H_Fock⊗H_Φ)=%d×%d=%d", a.MatterDimension, a.ScalarDimension, a.TensorDimension)},
					{Name: "temporal matter polarization", Passed: a.MatterSideOperatorFound && math.Abs(a.TemporalTrace) < eps, Detail: fmt.Sprintf("T0=1/2−N0 has Tr(T0)=%.3e, Tr(T0²)=%.10f", a.TemporalTrace, a.TemporalTraceSquared)},
					{Name: "compatible with B-L", Passed: a.TemporalCommutesWithBMinusLNorm < eps, Detail: fmt.Sprintf("||[T0,Q_B-L]||_F=%.3e", a.TemporalCommutesWithBMinusLNorm)},
					{Name: "vectorlike temporal test", Passed: !a.Vectorlike.FlippingAvailable, Detail: fmt.Sprintf("Y=T0+(B-L)/2 plus T_Φ gives Γ-flipping dim=%d; vectorlike temporal polarization is not physical chirality", a.Vectorlike.FlippingDim)},
					{Name: "even chiral restriction unlocks mixing", Passed: a.ChiralEven.FlippingAvailable, Detail: fmt.Sprintf("%s gives Γ-flipping dim=%d, preserving dim=%d", a.ChiralEven.Name, a.ChiralEven.FlippingDim, a.ChiralEven.PreservingDim)},
					{Name: "odd chiral restriction unlocks mixing", Passed: a.ChiralOdd.FlippingAvailable, Detail: fmt.Sprintf("%s gives Γ-flipping dim=%d, preserving dim=%d", a.ChiralOdd.Name, a.ChiralOdd.FlippingDim, a.ChiralOdd.PreservingDim)},
					{Name: "mirror ambiguity exposed", Passed: a.MirrorAmbiguity && !a.PhysicalOrientationSelected, Detail: "both even and odd chiral restrictions work algebraically; the physical left/right orientation is not selected yet"},
					{Name: "candidate hypercharge bridge", Passed: a.HyperchargeCandidateConstructed, Detail: "candidate form tested: Y_total=(T3_R+(B-L)/2)⊗I + I⊗T_Φ"},
					{Name: "Yukawa discipline", Passed: !a.ElectroweakYukawaDerived, Detail: "allowed Γ-flipping channels exist after chiral restriction, but no Yukawa texture or mass spectrum is claimed"},
					{Name: "remaining bridge unknowns", Passed: len(a.RemainingUnknowns) > 0, Detail: FormatUnknowns(a.RemainingUnknowns)},
				},
				Notes: []string{
					"The matter-side operator naturally available in the current Fock basis is temporal occupation polarization T0=1/2−N0.",
					"T0 is too vectorlike by itself. A chiral restriction creates charge-compatible grading-flipping channels, but the even/odd orientation is still a bridge choice.",
					"The engine has therefore found a viable finite T3_R candidate family, not a completed Standard Model hypercharge theorem.",
				},
			}
		},
	}
}
