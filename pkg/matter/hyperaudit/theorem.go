package hyperaudit

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func HyperchargeTableAuditTheorem() theorem.Theorem {
	const id = "MATTER-HYPERCHARGE-TABLE-AUDIT"
	const name = "chiral orientation and hypercharge table audit"
	return theorem.Theorem{
		ID:     id,
		Name:   name,
		Layer:  theorem.LayerMatter,
		Status: theorem.BridgeRequired,
		Verify: func() theorem.Result {
			a, err := BuildDefault()
			if err != nil {
				return theorem.Result{ID: id, Name: name, Layer: theorem.LayerMatter, Status: theorem.FailedRoute,
					Checks: []theorem.Check{{Name: "construct hypercharge table audit", Passed: false, Detail: err.Error()}},
				}
			}
			return theorem.Result{
				ID:     id,
				Name:   name,
				Layer:  theorem.LayerMatter,
				Status: theorem.BridgeRequired,
				Checks: []theorem.Check{
					{Name: "even-branch hypercharge spectrum", Passed: true, Detail: FormatCounts(a.Even.HyperchargeCounts)},
					{Name: "even branch rejected", Passed: !a.Even.MatchesRightSingletConjugateTable && a.Even.ExoticChargePresent, Detail: fmt.Sprintf("right-singlet score=%d/16; exotic charges %s", a.Even.RightSingletScore, FormatCounts(a.Even.ExoticCharges))},
					{Name: "odd-branch hypercharge spectrum", Passed: true, Detail: FormatCounts(a.Odd.HyperchargeCounts)},
					{Name: "odd branch matches right-singlet/conjugate table", Passed: a.Odd.MatchesRightSingletConjugateTable, Detail: fmt.Sprintf("right-singlet score=%d/16; no exotic charges=%v", a.Odd.RightSingletScore, !a.Odd.ExoticChargePresent)},
					{Name: "chiral orientation selected", Passed: a.ChiralOrientationSelected && a.PreferredBranchName == a.Odd.Name, Detail: fmt.Sprintf("preferred branch: %s", a.PreferredBranchName)},
					{Name: "full SM left-handed table not yet derived", Passed: !a.FullStandardModelTableDerived && a.SU2LDoubletBridgeMissing, Detail: "neither branch yields Q_L Y=1/6×6 and L_L Y=-1/2×2; SU(2)_L doublet bridge is still missing"},
					{Name: "conjugation convention still open", Passed: a.ConjugationConventionMissing, Detail: "odd branch gives a right-singlet/conjugate multiset; particle vs left-handed-conjugate orientation remains a bridge choice"},
					{Name: "remaining bridge unknowns", Passed: len(a.RemainingUnknowns) > 0, Detail: FormatUnknowns(a.RemainingUnknowns)},
				},
				Notes: []string{
					"The finite audit selects the odd Fock-parity restriction as the viable T3_R orientation for right-singlet/conjugate hypercharge values.",
					"This is stronger than Gate 21: the mirror ambiguity is broken at the hypercharge-table level.",
					"It is still not the full Standard Model hypercharge theorem, because SU(2)_L doublets and conjugation conventions are not yet derived.",
				},
			}
		},
	}
}
