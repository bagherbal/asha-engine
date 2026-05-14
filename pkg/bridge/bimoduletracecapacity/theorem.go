package bimoduletracecapacity

import "github.com/bagherbal/asha-engine/pkg/theorem"

func BimoduleTraceCapacitySieveSectorHierarchyAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-BIMODULE-TRACE-CAPACITY-SIEVE-SECTOR-HIERARCHY-AUDIT"
	const name = "Bimodule Trace Capacity Sieve / Sector Hierarchy Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 290 trace capacity audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 289 sector sensitivity is inherited", Passed: a.Inheritance.SectorProjectedTracesBranchSensitive && !a.Inheritance.SelectionPrincipleDerived && len(a.Inheritance.BranchesSurvived) == 2, Detail: a.Inheritance.Verdict},
			{Name: "Morita capacity candidates are formalized without native promotion", Passed: len(a.Capacity.CandidateBounds) == 2 && !a.Capacity.NativeBoundDerived, Detail: FormatCapacity(a.Capacity)},
			{Name: "branch stress test evaluates both r branches", Passed: len(a.Stress.Results) == 2, Detail: FormatStress(a.Stress)},
			{Name: "weak total capacity bound is too weak to select a branch", Passed: a.Stress.BothPassTotalCapacity, Detail: FormatStress(a.Stress)},
			{Name: "strong per-slot monotonic diagnostic would select r_plus but is unsealed", Passed: a.Stress.ExactlyOnePassPerSlot && a.Stress.PerSlotSelectedBranch == "r_plus" && !a.Veto.PerSlotCapacityIsNativeTheorem, Detail: FormatVeto(a.Veto)},
			{Name: "r_minus is not vetoed by derived geometry", Passed: !a.Veto.RMinusViolatesDerivedGeometry && !a.Selection.UniqueBranchSelected && len(a.Selection.SurvivingBranches) == 2, Detail: FormatSelection(a.Selection)},
			{Name: "Higgs prediction remains firewalled", Passed: !a.Higgs.HiggsPredictionClaimed && !a.Higgs.HeatKernelProjectionDerived, Detail: FormatHiggs(a.Higgs)},
			{Name: "firewalls preserve multiplicity/amplitude distinction", Passed: a.Firewalls.DoesNotPromoteMultiplicityToAmplitude && a.Firewalls.DoesNotUsePerSlotBoundAsTheorem && !a.Firewalls.FiniteCorePolluted, Detail: FormatFirewalls(a.Firewalls)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary)}}
	}}
}
