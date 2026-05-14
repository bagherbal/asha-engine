package tauetaweakselector

import "testing"

func TestGate259TauEtaWeakSelectorAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if !a.Summary.Gate258Inherited || !a.TauEta.StableNativeDegrees || a.TauEta.NativeFockPullbackDerived {
		t.Fatalf("expected Gate 258 inheritance and sealed tau_eta retrieval: %s / %s", FormatSummary(a.Summary), FormatTauEta(a.TauEta))
	}
	if !a.SpatialTag.ConditionalAlignmentApplied || a.SpatialTag.UniqueSpatialMode != 3 || a.SpatialTag.ComplementPlaneModes != [2]int{1, 2} || a.SpatialTag.NativeTauToFockPullbackDerived {
		t.Fatalf("unexpected spatial tag: %s", FormatSpatialTag(a.SpatialTag))
	}
	if a.WeakSieve.InputBMinusLSurvivorCount != 6 || a.WeakSieve.SurvivorCount != 2 || !a.WeakSieve.Reduced || !a.WeakSieve.UniqueUnorientedPlaneSelected || a.WeakSieve.UniqueOrientedFrameSelected {
		t.Fatalf("unexpected tau_eta weak sieve: %s", FormatWeakSieve(a.WeakSieve))
	}
	if a.ScalarSieve.InputBMinusLSurvivorCount != 2 || a.ScalarSieve.SurvivorCount != 2 || a.ScalarSieve.Reduced || !a.ScalarSieve.SignDegeneracyLeft {
		t.Fatalf("unexpected scalar sieve: %s", FormatScalarSieve(a.ScalarSieve))
	}
	if a.CombinedSieve.InputBMinusLWitnessCount != 12 || a.CombinedSieve.SurvivingWitnessCount != 4 || !a.CombinedSieve.Reduced || a.CombinedSieve.UniqueOrientation {
		t.Fatalf("unexpected combined witness sieve: %s", FormatCombined(a.CombinedSieve))
	}
	if !a.RestrictedScan.AllSurvivorsScanned || a.RestrictedScan.ResultCount != 12 || a.RestrictedScan.BranchCount != 3 {
		t.Fatalf("unexpected restricted scan inventory: %s", FormatRestrictedScan(a.RestrictedScan))
	}
	if a.RestrictedScan.ExactPolarized3PlaneResults != 0 || a.RestrictedScan.ExactFull3KernelResults != 0 || a.Summary.Neutral3PlaneDerived {
		t.Fatalf("tau_eta sieve must not force a three-plane: %s", FormatRestrictedScan(a.RestrictedScan))
	}
	if a.Firewall.ForcedWeakPlaneWithoutSeal || a.Firewall.TreatedTauEtaAsFiniteFockOperator || a.Firewall.SelectedTrialityByHand || a.Firewall.SelectedTrialityByDesiredKernel || a.Firewall.ForcedKernelDim3 || a.Firewall.PollutedFiniteCore {
		t.Fatalf("firewall violation: %s", FormatFirewall(a.Firewall))
	}
}
