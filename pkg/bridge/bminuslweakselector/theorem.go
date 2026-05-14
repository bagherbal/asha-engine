package bminuslweakselector

import "github.com/bagherbal/asha-engine/pkg/theorem"

func WeakPlaneSelectorBMinusLEmbeddingOrientationConstraintAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-WEAK-PLANE-SELECTOR-B-MINUS-L-EMBEDDING-ORIENTATION-CONSTRAINT-AUDIT"
	const name = "Weak-Plane Selector / B-L Embedding Orientation Constraint Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{{Name: "build Gate 258 B-L weak selector audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 257 witness scan is inherited without rerunning broad historical gates", Passed: a.Inheritance.NativeChargeEigenvaluesExtracted && a.Inheritance.EmbeddingWitnessesScanned && a.Inheritance.AllTrialityBranchesScanned && !a.Inheritance.Gate257Neutral3PlaneDerived, Detail: FormatInheritance(a.Inheritance)},
			{Name: "native B-L 1⊕3 Fock ledger is retrieved as selector", Passed: a.BMinusL.DerivedFiniteFockLedger && a.BMinusL.OnePlusThreeSplit && a.BMinusL.SpatialIsotropy && !a.BMinusL.UsesObservedInput, Detail: FormatBMinusL(a.BMinusL)},
			{Name: "B-L scalar/contact sieve reduces scalar embeddings before any kernel test", Passed: a.ScalarSieve.Reduced && a.ScalarSieve.InputCount == 8 && a.ScalarSieve.SurvivorCount == 2 && !a.ScalarSieve.UniqueSelected, Detail: FormatScalarSieve(a.ScalarSieve)},
			{Name: "B-L weak-frame sieve rejects temporal-spatial weak planes and preserves only spatial-spatial planes", Passed: a.WeakSieve.Reduced && a.WeakSieve.InputCount == 12 && a.WeakSieve.SurvivorCount == 6 && !a.WeakSieve.UniqueSelected, Detail: FormatWeakSieve(a.WeakSieve)},
			{Name: "combined witness space is reduced independently of the desired 3-plane", Passed: a.CombinedSieve.Reduced && a.CombinedSieve.InputWitnessCount == 96 && a.CombinedSieve.SurvivingWitnessCount == 12 && !a.CombinedSieve.UniqueOrientation, Detail: FormatCombined(a.CombinedSieve)},
			{Name: "restricted all-branch triality scan is completed on the B-L survivors", Passed: a.RestrictedScan.ScannedAfterSelector && a.RestrictedScan.AllSurvivorsScanned && a.RestrictedScan.BranchCount == 3 && a.RestrictedScan.ResultCount == 36, Detail: FormatRestrictedScan(a.RestrictedScan)},
			{Name: "B-L sieve still does not derive an exact neutral 3-plane", Passed: a.RestrictedScan.ExactPolarized3PlaneResults == 0 && a.RestrictedScan.ExactFull3KernelResults == 0 && !a.Summary.Neutral3PlaneDerived, Detail: FormatRestrictedScan(a.RestrictedScan)},
			{Name: "firewall prevents B-L from being used as an outcome-tuned selector", Passed: a.Firewall.Gate257NoGoPreserved && a.Firewall.BMinusLAppliedBeforeKernel && a.Firewall.BMinusLUsedAsSelectorNotOutcome && !a.Firewall.ForcedWeakPlane && !a.Firewall.ForcedScalarOrientation && !a.Firewall.SelectedTrialityByHand && !a.Firewall.ForcedKernelDim3 && !a.Firewall.AcceptedYOnlyAsQ && !a.Firewall.PollutedFiniteCore, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{
			a.TruthStatement,
			"Gate 258 proves that B-L is a genuine native selector: it cuts away temporal-spatial weak frames and non-isotropic contact scalar embeddings.",
			"The selector is not strong enough to choose one weak plane, one scalar sign, one triality branch, or the neutral three-plane; the Yukawa texture remains sealed.",
		}}
	}}
}
