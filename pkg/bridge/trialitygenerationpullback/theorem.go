package trialitygenerationpullback

import "github.com/bagherbal/asha-engine/pkg/theorem"

func TrialityGenerationPullbackNativeTopYukawaBoundarySieveAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-TRIALITY-GENERATION-PULLBACK-NATIVE-TOP-YUKAWA-BOUNDARY-SIEVE"
	const name = "Triality Generation Pullback / Native Top-Yukawa Boundary Sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 323 triality pullback audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "tau_eta is pulled back onto the three-generation quark trace carrier", Passed: a.Carrier.Formalized && len(a.Carrier.TauEta) == 3 && a.Carrier.GenerationBreaking && a.Carrier.MagnitudeSpectrumDegenerate && !a.Carrier.PullbackUnique, Detail: FormatCarrier(a.Carrier)},
			{Name: "amplitude fractionalization yields normalized |tau_eta|^2 weights", Passed: a.Fractional.Formalized && len(a.Fractional.NormalizedWeights) == 3 && a.Fractional.GeneratesUniqueLowSlot && a.Fractional.GeneratesTwoHighSlots && a.Fractional.SumWeights > 0.999999 && a.Fractional.SumWeights < 1.000001, Detail: FormatFractionalization(a.Fractional)},
			{Name: "top-slot candidates are audited but not canonically derived", Passed: len(a.Candidates) >= 4 && !a.PullbackVerdict.CanonicalTopFractionDerived && a.PullbackVerdict.TopBoundaryStatus == StatusFailedNativeTopBoundaryNotDerived, Detail: FormatPullbackVerdict(a.PullbackVerdict)},
			{Name: "derived threshold transport is rerun for each candidate", Passed: len(a.Preflights) == len(a.Candidates) && a.Summary.PhysicalPreflightExecuted, Detail: a.Summary.DirectAnswer},
			{Name: "nonzero tau_eta top candidates do not preserve the Gate 322 near-125 GeV proxy", Passed: a.PullbackVerdict.NonzeroTauEtaTopSpoils125 && a.PullbackVerdict.GaugeOnlyStillRequired && !a.Summary.Gate322SuccessPreservedByCanonicalTop, Detail: FormatPullbackVerdict(a.PullbackVerdict)},
			{Name: "firewalls against final top/pole-mass claims are preserved", Passed: a.Firewalls.NoObservedTopMassInserted && a.Firewalls.NoCKMImported && a.Firewalls.NoFlavorTextureInvented && a.Firewalls.NoPoleMassClaimed && a.Firewalls.NoTwoLoopClaimed && a.Firewalls.NoFinalColliderMassClaimed && !a.Firewalls.FiniteCorePolluted && !a.Summary.FinalMassClaimed, Detail: FormatFirewalls(a.Firewalls)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks}
	}}
}
