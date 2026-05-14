package bgapmajoranaactivation

import "testing"

func TestMajoranaActivationAndTrace(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Activation.Formalized || !a.Activation.ActivatedAsConditionalSeal || a.Activation.ActivatedAsPhysicalMass || a.Activation.NativeDerivation || a.Activation.BGap <= 0 || a.Activation.KappaM != 1 {
		t.Fatalf("bad activation: %s", FormatActivation(a.Activation))
	}
	if !a.Trace.Formalized || a.Trace.KappaM != 1 || a.Trace.MajoranaTrace2 == "" || a.Trace.MajoranaTrace4 == "" || a.Trace.CrossTermsDerived {
		t.Fatalf("bad trace extension: %s", FormatTrace(a.Trace))
	}
}

func TestSigmaPotentialAndLanes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Potential.Formalized || a.Potential.LambdaHH < 0.25 || a.Potential.LambdaSigmaSigma <= 0 || !a.Potential.RequiresPortal || !a.Potential.RequiresSigmaVEV {
		t.Fatalf("bad potential: %s", FormatPotential(a.Potential))
	}
	if len(a.Lanes) != 3 || a.Lanes[0].EffectiveLambdaUV <= a.Lanes[1].EffectiveLambdaUV || a.Lanes[2].EffectiveLambdaUV != 0 || !a.Lanes[2].StableNonNegative {
		t.Fatalf("bad correction lanes: %s || %s || %s", FormatLane(a.Lanes[0]), FormatLane(a.Lanes[1]), FormatLane(a.Lanes[2]))
	}
}

func TestRGRerunCapacity(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if len(a.Results) != 3 {
		t.Fatalf("expected three results, got %d", len(a.Results))
	}
	for _, r := range a.Results {
		if !r.Computed || !r.Perturbative || r.HiggsMassGeV < 331 || r.FinalLambdaAtV < 0.90 {
			t.Fatalf("bad RG result: %s", FormatRGResult(r))
		}
	}
	if !a.Capacity.Formalized || a.Capacity.BoundaryCorrectionCanResolve || !a.Capacity.TopSectorDominates || a.Capacity.BestStableMassGapGeV < 200 || a.Capacity.BoundaryCorrectionMovesMassGeV > 0.01 {
		t.Fatalf("capacity audit should show unresolved top-dominated tension: %s", FormatCapacity(a.Capacity))
	}
}

func TestFirewallsAndSummary(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault returned error: %v", err)
	}
	if !a.Firewalls.NoObservedMassFitInserted || !a.Firewalls.NoPortalCouplingFitted || !a.Firewalls.NoSigmaVEVFitted || !a.Firewalls.NoThresholdJumpInserted || !a.Firewalls.NoTwoLoopRGExecuted || !a.Firewalls.NoPoleMassConversionInserted || !a.Firewalls.MajoranaEdgeRemainsConditional || !a.Firewalls.NoFinalMassClaimed || a.Firewalls.FiniteCorePolluted {
		t.Fatalf("firewall failed: %s", FormatFirewalls(a.Firewalls))
	}
	if !a.Summary.MajoranaActivationFormalized || !a.Summary.TraceExtensionFormalized || !a.Summary.SigmaCorrectionFormalized || !a.Summary.RGRerunComputed || a.Summary.BGapBoundaryCorrectionSolves || !a.Summary.TopSectorStillDominates || a.Summary.FinalMassClaimed || !a.Summary.FirewallPreserved {
		t.Fatalf("summary failed: %s", FormatSummary(a.Summary))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := BGapMajoranaActivationSpectralActionSigmaHMixedQuarticCorrectionAuditTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem did not pass:\n%s", res.Details())
	}
}
