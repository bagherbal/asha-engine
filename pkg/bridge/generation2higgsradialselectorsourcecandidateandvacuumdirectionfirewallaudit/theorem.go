package generation2higgsradialselectorsourcecandidateandvacuumdirectionfirewallaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HiggsRadialSelectorSourceCandidateAndVacuumDirectionFirewallAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 737 — Higgs Radial Selector Source-Candidate and Vacuum-Direction Firewall Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate737 radial selector audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate736 rho_plus and selector firewalls", Passed: a.Gate736.Inherited && a.Gate736.RhoPlusMaxEntropy && a.Gate736.RadialWeightCertified && a.Gate736.RhoPlusDoesNotSelectPRad && a.Gate736.RhoPlusDoesNotSelectN && a.Gate736.HistoryLoopConditional && strings.Contains(a.Gate736.Verdict, StatusGate736K7PlusMaximumEntropyObserverInherited), Detail: FormatGate736(a.Gate736)},
			{Name: "define radial selector problem", Passed: a.Problem.Rank == 1 && a.Problem.NeedsLineInsideK7Plus && !a.Problem.CurrentlyNative && strings.Contains(a.Problem.Verdict, StatusRadialSelectorProblemDefined), Detail: FormatProblem(a.Problem)},
			{Name: "audit candidate sources and reject native selector", Passed: len(a.Candidates.Candidates) == 8 && !a.Candidates.AnyNativeSelectorFound && !a.Candidates.BoundaryScalarsContainVector && strings.Contains(a.Candidates.Verdict, StatusNoNativeRadialProjectorSelectorFound) && strings.Contains(a.Candidates.Verdict, StatusBoundaryScalarDataDoNotSelectPRad), Detail: FormatCandidates(a.Candidates)},
			{Name: "audit symmetry obstruction", Passed: a.Symmetry.RequiresVacuumSelector && !a.Symmetry.CurrentDataSelectsLine && strings.Contains(a.Symmetry.Verdict, StatusPRadRequiresSymmetryBreakingOrVacuumSelector), Detail: FormatSymmetry(a.Symmetry)},
			{Name: "classify P_rad as type-distinct seal", Passed: len(a.Seal.SealNames) == 3 && a.Seal.DistinctFromN && a.Seal.DistinctFromQ && a.Seal.DistinctFromRhoPlus && strings.Contains(a.Seal.Verdict, StatusPRadTypeDistinctScalarVacuumDirectionSeal), Detail: FormatSeal(a.Seal)},
			{Name: "record HistoryLoop dependence on P_rad", Passed: a.HistoryLoop.RhoPlusSuppliesWeight && a.HistoryLoop.PRadSuppliesEvent && a.HistoryLoop.NSuppliesPhaseLoop && !a.HistoryLoop.QSuppliesPRad && a.HistoryLoop.ConditionalWithoutPRad && strings.Contains(a.HistoryLoop.Verdict, StatusHistoryLoopUnitSourceConditionalWithoutPRad), Detail: FormatHistoryLoop(a.HistoryLoop)},
			{Name: "preserve physical firewalls", Passed: !a.Firewall.PRadIsHiggsVacuumTheorem && !a.Firewall.PRadIsElectroweakBreakingTheorem && !a.Firewall.PhaseTransverseGoldstoneTheorem && !a.Firewall.PRadIsHiggsMassTheorem && !a.Firewall.YukawaOperatorOrEigenvalueTheorem && strings.Contains(a.Firewall.Verdict, StatusGate737Boundary), Detail: FormatFirewall(a.Firewall)},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := append([]string{a.Truth}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
