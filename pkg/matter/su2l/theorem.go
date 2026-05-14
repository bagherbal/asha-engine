package su2l

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func DoubletAuditTheorem() theorem.Theorem {
	const id = "MATTER-SU2L-DOUBLET-AUDIT"
	const name = "Yukawa-induced SU(2)_L doublet hypercharge audit"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerMatter,
		Status: theorem.BridgeRequired,
		Verify: func() theorem.Result {
			a, err := BuildDefault()
			if err != nil {
				return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct SU(2)_L doublet audit", Passed: false, Detail: err.Error()}},
				}
			}
			return theorem.Result{
				ID:     id,
				Name:   name,
				Layer:  theorem.LayerMatter,
				Status: theorem.BridgeRequired,
				Checks: []theorem.Check{
					{Name: "odd right-singlet table available", Passed: a.OddRightTableAvailable, Detail: "uses Gate 22 preferred odd T3_R branch"},
					{Name: "finite scalar charge pair", Passed: len(a.ScalarCharges) == 2, Detail: fmt.Sprintf("Y_Φ candidates = [%.6g, %.6g] from pair weight %.6g", a.ScalarCharges[0], a.ScalarCharges[1], a.ScalarWeight)},
					{Name: "standard-orientation charge balance", Passed: a.Standard.ChargeBalanceExact && a.Standard.MatchesStandardOrientation, Detail: FormatOrientation(a.Standard)},
					{Name: "conjugate-orientation charge balance", Passed: a.Conjugate.ChargeBalanceExact && a.Conjugate.MatchesConjugateOrientation, Detail: FormatOrientation(a.Conjugate)},
					{Name: "charge-level SU(2)_L doublets derived", Passed: a.ChargeLevelSU2LDoubletsDerived, Detail: "Y_R − Y_Φ yields Q_L Y=1/6×6 and L_L Y=-1/2×2, plus the conjugate mirror"},
					{Name: "neutral seed ambiguity exposed", Passed: a.NeutralSeedAmbiguity, Detail: "odd branch has two neutral seeds; selecting ν_R vs ν^c still requires a finite reality/CPT convention"},
					{Name: "nonabelian SU(2)_L generators not yet derived", Passed: !a.NonabelianSU2LGeneratorsDerived, Detail: "this gate proves charge-level doublet hypercharges, not raising/lowering generators"},
					{Name: "Yukawa intertwiner still open", Passed: !a.YukawaIntertwinerDerived, Detail: "charge channels are fixed; explicit finite intertwiner matrices remain open"},
					{Name: "remaining bridge unknowns", Passed: len(a.RemainingUnknowns) > 0, Detail: FormatUnknowns(a.RemainingUnknowns)},
				},
				Notes: []string{
					"Gate 22 gave the right-singlet/conjugate table. Gate 23 shows that adding the finite scalar doublet charge reconstructs the left doublet hypercharges by the charge-balance equation Y_L = Y_R − Y_Φ.",
					"This upgrades U-13 from missing to charge-level solved: Q_L and L_L hypercharges appear with the correct multiplicities, together with the conjugate mirror orientation.",
					"The theorem is intentionally not yet a full SU(2)_L gauge theorem; the nonabelian generators and explicit Yukawa intertwiners are the next targets.",
				},
			}
		},
	}
}
