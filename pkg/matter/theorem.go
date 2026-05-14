package matter

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FockContactBridgeTheorem() theorem.Theorem {
	const id = "MATTER-FOCK-CONTACT-BRIDGE"
	const name = "Fock matter basis and contact/Higgs kinematic bridge"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerMatter,
		Status: theorem.BridgeRequired,
		Verify: func() theorem.Result {
			b, err := BuildDefaultFockContactBridge()
			if err != nil {
				return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct Fock/contact bridge", Passed: false, Detail: err.Error()}},
				}
			}

			checks := []theorem.Check{
				{
					Name:   "Fock generation seed",
					Passed: b.FockModeCount == 4 && b.FockStateCount == 16,
					Detail: fmt.Sprintf("four covariant creation modes generate %d states", b.FockStateCount),
				},
				{
					Name:   "sterile vacuum seed",
					Passed: b.VacuumSterilitySeedAvailable,
					Detail: fmt.Sprintf("sterile B−L-neutral vacuum candidates=%d", b.SterileVacuumCount),
				},
				{
					Name:   "one-particle seed split",
					Passed: b.OneParticleStateCount == 4 && b.QuarkSeedCount == 3 && b.LeptonSeedCount == 1,
					Detail: fmt.Sprintf("one-particle seeds=%d = %d color/quark seeds + %d lepton seed", b.OneParticleStateCount, b.QuarkSeedCount, b.LeptonSeedCount),
				},
				{
					Name:   "Fock/Higgs active-dimension match",
					Passed: b.ModeToActiveScalarMatch,
					Detail: fmt.Sprintf("Fock modes=%d and active contact-Higgs directions=%d", b.FockModeCount, b.ActiveHiggsDirections),
				},
				{
					Name:   "spatial/protected contact resonance",
					Passed: b.SpatialToProtectedMatch,
					Detail: fmt.Sprintf("spatial/color-seed modes=%d and protected unmixed contact directions=%d", b.SpatialModeCount, b.ProtectedContactDirections),
				},
				{
					Name:   "pair organization available",
					Passed: b.PairDegenerateHiggsSpectrum,
					Detail: "Higgs/contact active spectrum is pair-degenerate, supporting a complex two-component bridge candidate",
				},
				{
					Name:   "canonical embedding discipline",
					Passed: !b.CanonicalEmbeddingConstructed,
					Detail: "OPEN U-05: no canonical Fock |Ω⟩ → K₇ vector map is claimed in this gate",
				},
				{
					Name:   "Yukawa discipline",
					Passed: !b.YukawaOperatorConstructed,
					Detail: "OPEN U-07: no Yukawa texture or fermion mass matrix is claimed before a canonical representation action exists",
				},
			}
			notes := []string{
				"This gate upgrades the bridge from vague numerology to a typed kinematic resonance: 4 Fock modes, 16 Fock states, 4 active Higgs/contact directions, and 3 protected contact directions.",
				"The sterile vacuum candidate is exact inside the Fock bookkeeping, but dark-matter stability still requires a dynamical/CPT/contact bridge.",
				"The next hard theorem is a representation action: the finite Higgs operator must act on the Fock basis before masses or Yukawa eigenvalues can be discussed.",
			}
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
		},
	}
}
