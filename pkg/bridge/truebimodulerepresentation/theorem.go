package truebimodulerepresentation

import "github.com/bagherbal/asha-engine/pkg/theorem"

func TrueBimoduleAssemblyLeftRightRepresentationAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-TRUE-BIMODULE-ASSEMBLY-LEFT-RIGHT-REPRESENTATION-AUDIT"
	const name = "True Bimodule Assembly / Left-Right Representation Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 295 true bimodule audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 294 direct-sum paradox is inherited", Passed: a.Input.JSwapKOSix && a.Input.NaiveLeftActionFailed, Detail: FormatInput(a.Input)},
			{Name: "weak H action is isolated on the left", Passed: a.Left.ActsOnQuarks && a.Left.ActsOnLeptons && a.Left.NonTrivialResidual > 0.1, Detail: FormatLeft(a.Left)},
			{Name: "color M3 action is isolated on the right/opposite side", Passed: a.Right.ActsOnQuarks && !a.Right.ActsOnLeptons && a.Right.NonTrivialResidual > 0.1, Detail: FormatRight(a.Right)},
			{Name: "weak-left/color-right actions commute on Q_L", Passed: a.Bimodule.ZeroOrderVerified && a.Bimodule.WeakColorCommutatorNorm < 1e-12 && a.Bimodule.NaiveLeftCrossTermNorm > 0.1, Detail: FormatBimodule(a.Bimodule)},
			{Name: "true bimodule resolves the zero-order direct-sum paradox", Passed: a.Summary.TrueBimoduleDerived && a.Summary.ZeroOrderResolved, Detail: FormatSummary(a.Summary)},
			{Name: "hypercharge splitting remains a separate un-derived ledger", Passed: !a.Hypercharge.DerivedByGate && !a.Hypercharge.FractionalChargesGenerated, Detail: FormatHypercharge(a.Hypercharge)},
			{Name: "first-order Dirac theorem remains blocked without canonical D_F", Passed: a.OrderOne.ZeroOrderConditionVerified && !a.OrderOne.FirstOrderVerified && !a.OrderOne.CanonicalDiracAvailable, Detail: FormatOrderOne(a.OrderOne)},
			{Name: "firewalls preserve Higgs/B-gap dynamics", Passed: !a.Firewalls.FiniteCorePolluted && a.Firewalls.DoesNotUnlockHiggs && a.Firewalls.DoesNotUnlockBGap, Detail: FormatFirewalls(a.Firewalls)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary)}}
	}}
}
