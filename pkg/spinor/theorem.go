package spinor

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FockSpaceTheorem() theorem.Theorem {
	const id = "SPINOR-WITT-FOCK-16"
	const name = "Witt/Fock matter basis from four covariant modes"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerMatter,
		Status: theorem.ExactFinite,
		Verify: func() theorem.Result {
			const eps = 1e-12
			f, err := NewCovariantPhaseFockSpace(4)
			if err != nil {
				return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct Fock space", Passed: false, Detail: err.Error()}},
				}
			}
			vacuum, err := f.Vacuum()
			if err != nil {
				return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "locate vacuum", Passed: false, Detail: err.Error()}},
				}
			}
			oneParticle := f.OneParticleStates()
			quarkSeeds := 0
			leptonSeeds := 0
			for _, s := range oneParticle {
				charge := s.BMinusL()
				if math.Abs(charge-(1.0/3.0)) < eps {
					quarkSeeds++
				}
				if math.Abs(charge+1.0) < eps {
					leptonSeeds++
				}
			}

			checks := []theorem.Check{
				{Name: "covariant mode count", Passed: f.ModeCount() == 4, Detail: fmt.Sprintf("four creation modes a†_μ from μ=0..3; got %d", f.ModeCount())},
				{Name: "finite Fock dimension", Passed: f.StateCount() == f.ExpectedStateCount() && f.StateCount() == 16, Detail: fmt.Sprintf("2^%d=%d occupation states", f.ModeCount(), f.StateCount())},
				{Name: "1+3 mode split", Passed: f.TemporalModeCount() == 1 && f.SpatialModeCount() == 3, Detail: fmt.Sprintf("temporal modes=%d, spatial/color-seed modes=%d", f.TemporalModeCount(), f.SpatialModeCount())},
				{Name: "sterile vacuum candidate", Passed: vacuum.IsSterileVacuumCandidate(eps), Detail: fmt.Sprintf("|Ω⟩ has excitation number=%d and B−L=%.1f", vacuum.ExcitationNumber(), vacuum.BMinusL())},
				{Name: "one-particle seed charges", Passed: quarkSeeds == 3 && leptonSeeds == 1, Detail: fmt.Sprintf("one-particle states split into %d quark-color seeds with B−L=1/3 and %d lepton seed with B−L=-1", quarkSeeds, leptonSeeds)},
			}
			notes := []string{
				"This theorem only constructs the finite Fock bookkeeping. It does not yet assign full Standard Model hypercharge or Yukawa masses.",
				"The vacuum state is sterile under this B−L bookkeeping, but dark-matter stability remains a bridge theorem.",
			}
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.ExactFinite, Checks: checks, Notes: notes}
		},
	}
}
