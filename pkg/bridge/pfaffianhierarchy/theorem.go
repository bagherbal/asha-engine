package pfaffianhierarchy

import "github.com/bagherbal/asha-engine/pkg/theorem"

func PfaffianHalfActionHierarchyFermionicFluctuationDeterminantTheorem() theorem.Theorem {
	const id = "BRIDGE-PFAFFIAN-HALF-ACTION-HIERARCHY-FERMIONIC-FLUCTUATION-DETERMINANT"
	const name = "Pfaffian Half-Action Hierarchy / Fermionic Fluctuation Determinant Derivation"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 341 Pfaffian hierarchy audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 340 hierarchy promotion audit inherited", Passed: a.Inputs.HighestInheritedGate == inheritedHighestGate && a.Inputs.NGen == 3 && a.Inputs.STop > 78, Detail: FormatInputs(a.Inputs)},
			{Name: "Pfaffian half-action formalized", Passed: a.Pfaffian.HalfActionAuthorized && !a.Pfaffian.FiniteCoreDerived && a.Pfaffian.HalfExponential > 7.15e-18 && a.Pfaffian.HalfExponential < 7.16e-18, Detail: FormatPfaffian(a.Pfaffian)},
			{Name: "three-generation fluctuation factor formalized", Passed: a.Generation.NGen == 3 && a.Generation.CombinedFactor > 2.828 && a.Generation.CombinedFactor < 2.829 && !a.Generation.FiniteCoreDerived, Detail: FormatGeneration(a.Generation)},
			{Name: "combined hierarchy prediction matches unreduced branch at sub-percent level", Passed: a.Prediction.RatioToUnreducedTarget > 1.003 && a.Prediction.RatioToUnreducedTarget < 1.0045 && a.Prediction.RatioToReducedTarget > 0.20 && a.Prediction.RatioToReducedTarget < 0.201, Detail: FormatPrediction(a.Prediction)},
			{Name: "gravity connection formalized but f2 remains unlocked", Passed: a.Gravity.ElectroweakToGravityLinked && !a.Gravity.F2MomentLocked, Detail: FormatGravity(a.Gravity)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary), FormatStatuses(Statuses(a))}}
	}}
}
