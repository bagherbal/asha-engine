package tauetaweakselector

import "github.com/bagherbal/asha-engine/pkg/theorem"

func SpatialS3SieveTauEtaTopologicalOrientationSelectorAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-SPATIAL-S3-SIEVE-TAU-ETA-TOPOLOGICAL-ORIENTATION-SELECTOR-AUDIT"
	const name = "Spatial S3 Sieve / tau_eta Topological Orientation Selector Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{{Name: "build Gate 259 tau_eta weak selector audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 258 B-L selector is inherited without broad historical execution", Passed: a.Inheritance.BMinusLLedgerRetrieved && a.Inheritance.ScalarSieveReduced && a.Inheritance.WeakFrameSieveReduced && a.Inheritance.RestrictedTrialityRescanned && !a.Inheritance.Gate258Neutral3PlaneDerived, Detail: FormatInheritance(a.Inheritance)},
			{Name: "tau_eta topological signature is retrieved as audited scalar fundamental-class data", Passed: a.TauEta.StableNativeDegrees && a.TauEta.ScalarTraceFunctionalOnly && a.TauEta.Sequence[0] == 2 && a.TauEta.Sequence[1] == -2 && a.TauEta.Sequence[2] == 1 && a.TauEta.TwoPlusOneMagnitudeSelector && !a.TauEta.NativeFockPullbackDerived, Detail: FormatTauEta(a.TauEta)},
			{Name: "SpontaneousCarrierSeal conditionally aligns tau_eta with spatial Fock modes", Passed: a.SpatialTag.ConditionalAlignmentApplied && a.SpatialTag.AlignmentSeal == "SpontaneousCarrierSeal" && a.SpatialTag.UniqueSpatialMode == 3 && a.SpatialTag.ComplementPlaneModes == [2]int{1, 2} && !a.SpatialTag.NativeTauToFockPullbackDerived && !a.SpatialTag.ManualUnsealedAxisAssignment, Detail: FormatSpatialTag(a.SpatialTag)},
			{Name: "tau_eta weak-plane sieve reduces B-L spatial frames to the complementary U12 orientation pair", Passed: a.WeakSieve.Reduced && a.WeakSieve.InputBMinusLSurvivorCount == 6 && a.WeakSieve.SurvivorCount == 2 && a.WeakSieve.UniqueUnorientedPlaneSelected && !a.WeakSieve.UniqueOrientedFrameSelected, Detail: FormatWeakSieve(a.WeakSieve)},
			{Name: "tau_eta does not select the uniform scalar sign mirror", Passed: !a.ScalarSieve.Reduced && a.ScalarSieve.InputBMinusLSurvivorCount == 2 && a.ScalarSieve.SurvivorCount == 2 && a.ScalarSieve.SignDegeneracyLeft, Detail: FormatScalarSieve(a.ScalarSieve)},
			{Name: "combined B-L/tau_eta witness space is reduced before kernel inspection", Passed: a.CombinedSieve.Reduced && a.CombinedSieve.InputBMinusLWitnessCount == 12 && a.CombinedSieve.SurvivingWitnessCount == 4 && !a.CombinedSieve.UniqueOrientation, Detail: FormatCombined(a.CombinedSieve)},
			{Name: "restricted all-branch triality scan is completed on tau_eta survivors", Passed: a.RestrictedScan.ScannedAfterTauEtaSelector && a.RestrictedScan.AllSurvivorsScanned && a.RestrictedScan.BranchCount == 3 && a.RestrictedScan.ResultCount == 12, Detail: FormatRestrictedScan(a.RestrictedScan)},
			{Name: "tau_eta sieve still does not derive an exact neutral 3-plane", Passed: a.RestrictedScan.ExactPolarized3PlaneResults == 0 && a.RestrictedScan.ExactFull3KernelResults == 0 && !a.Summary.Neutral3PlaneDerived, Detail: FormatRestrictedScan(a.RestrictedScan)},
			{Name: "firewall preserves Gate 242 and Gate 258 no-go results", Passed: a.Firewall.Gate258NoGoPreserved && a.Firewall.TauEtaRetrievedFromAudit && a.Firewall.TauEtaNativePullbackPreserved && a.Firewall.ConditionalSSBAlignmentUsed && a.Firewall.TauEtaUsedAsSelectorNotOutcome && a.Firewall.SelectorAppliedBeforeKernel && !a.Firewall.ForcedWeakPlaneWithoutSeal && !a.Firewall.ForcedScalarOrientation && !a.Firewall.SelectedTrialityByHand && !a.Firewall.SelectedTrialityByDesiredKernel && !a.Firewall.ForcedKernelDim3 && !a.Firewall.TreatedTauEtaAsFiniteFockOperator && !a.Firewall.PollutedFiniteCore, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{
			a.TruthStatement,
			"Gate 259 uses tau_eta only under the SpontaneousCarrierSeal: the spatial tag is a conditional vacuum alignment, not a native tau_eta-to-Fock pullback theorem.",
			"The selector is real: it reduces the B-L-compatible weak frames from six to the U12 orientation pair. It is still insufficient for the neutral three-plane in the Cartan electroweak route.",
		}}
	}}
}
