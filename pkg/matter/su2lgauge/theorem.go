package su2lgauge

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func GeneratorAuditTheorem() theorem.Theorem {
	const id = "MATTER-SU2L-GENERATOR-AUDIT"
	const name = "finite SU(2)_L doublet generator audit"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerMatter,
		Status: theorem.BridgeRequired,
		Verify: func() theorem.Result {
			a, err := BuildDefault()
			if err != nil {
				return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct SU(2)_L generator audit", Passed: false, Detail: err.Error()}},
				}
			}
			eps := 1e-10
			return theorem.Result{
				ID:     id,
				Name:   name,
				Layer:  theorem.LayerMatter,
				Status: theorem.BridgeRequired,
				Checks: []theorem.Check{
					{Name: "left-doublet state space", Passed: a.Dimension == 8, Detail: fmt.Sprintf("dim(H_L)=8 = Q_L(3 colors×2) + L_L(2); states=%s", FormatStates(a.States))},
					{Name: "multiplet organization", Passed: a.QuarkColorDiagonal && a.LeptonDoubletFound, Detail: FormatMultiplets(a.Multiplets)},
					{Name: "SU(2) ladder relation [T3,T+]", Passed: a.CommutatorT3TPlusNorm < eps, Detail: fmt.Sprintf("||[T3,T+]−T+||_F = %.3e", a.CommutatorT3TPlusNorm)},
					{Name: "SU(2) ladder relation [T3,T-]", Passed: a.CommutatorT3TMinusNorm < eps, Detail: fmt.Sprintf("||[T3,T-]+T-||_F = %.3e", a.CommutatorT3TMinusNorm)},
					{Name: "SU(2) ladder relation [T+,T-]", Passed: a.CommutatorTPlusTMinusNorm < eps, Detail: fmt.Sprintf("||[T+,T-]−2T3||_F = %.3e", a.CommutatorTPlusTMinusNorm)},
					{Name: "hypercharge commutes with SU(2)_L", Passed: a.CommutesWithHyperchargeNorm < eps, Detail: fmt.Sprintf("||[Y,T+]+[Y,T-]||_F = %.3e", a.CommutesWithHyperchargeNorm)},
					{Name: "Gell-Mann–Nishijima charge identity", Passed: a.GellMannNishijimaNorm < eps, Detail: fmt.Sprintf("||Q−(T3+Y)||_F = %.3e", a.GellMannNishijimaNorm)},
					{Name: "nonabelian SU(2)_L generators derived at charge-table level", Passed: a.NonabelianSU2LGeneratorsDerived, Detail: "explicit T3,T+,T- matrices act on the derived left doublet space"},
					{Name: "finite geometric origin still open", Passed: !a.ContinuumGaugeFieldDerived, Detail: "this gate derives the left-doublet representation, not yet a continuum gauge field or a direct Boolean-compressed origin"},
					{Name: "Yukawa intertwiner still open", Passed: !a.YukawaIntertwinerDerived, Detail: "explicit gauge-compatible mass/intertwiner maps are the next gate"},
					{Name: "remaining bridge unknowns", Passed: len(a.RemainingUnknowns) > 0, Detail: FormatUnknowns(a.RemainingUnknowns)},
				},
				Notes: []string{
					"Gate 23 solved SU(2)_L at the hypercharge-selection level. Gate 24 adds explicit finite ladder generators on the derived Q_L and L_L doublet space.",
					"The SU(2)_L action is now a real representation theorem on the audited left-doublet table. The remaining harder problem is to derive the same action directly from finite Boolean/contact connection geometry.",
				},
			}
		},
	}
}
