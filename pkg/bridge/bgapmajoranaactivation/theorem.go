package bgapmajoranaactivation

import "github.com/bagherbal/asha-engine/pkg/theorem"

func BGapMajoranaActivationSpectralActionSigmaHMixedQuarticCorrectionAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-BGAP-MAJORANA-ACTIVATION-SPECTRAL-ACTION-SIGMA-H-MIXED-QUARTIC-CORRECTION-AUDIT"
	const name = "B-Gap Majorana Activation in the Spectral Action / σ-H Mixed Quartic Correction"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 312 B-gap Majorana activation audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "B-gap Majorana / σ carrier is conditionally activated without claiming a native mass theorem", Passed: a.Activation.Formalized && a.Activation.ActivatedAsConditionalSeal && !a.Activation.ActivatedAsPhysicalMass && !a.Activation.NativeDerivation && a.Activation.BGap > 0 && a.Activation.KappaM == 1, Detail: FormatActivation(a.Activation)},
			{Name: "majorana trace extension records κM and separates diagonal Majorana terms from the mixed Hσ projector", Passed: a.Trace.Formalized && a.Trace.KappaM == 1 && a.Trace.MajoranaTrace2 != "" && a.Trace.MajoranaTrace4 != "" && a.Trace.MixedQuarticTerm != "" && !a.Trace.CrossTermsDerived, Detail: FormatTrace(a.Trace)},
			{Name: "σ-H mixed quartic correction is formalized as λeff=λHH-λHσ²/(4λσσ)", Passed: a.Potential.Formalized && a.Potential.LambdaHH > 0.25 && a.Potential.LambdaSigmaSigma > 0 && a.Potential.RequiresPortal && a.Potential.RequiresSigmaVEV, Detail: FormatPotential(a.Potential)},
			{Name: "correction lanes include control, B-gap overlap witness, and maximal stable cancellation", Passed: len(a.Lanes) == 3 && a.Lanes[0].Name == "unactivated_control_gate308_boundary" && a.Lanes[1].PortalChi > 0.10 && a.Lanes[1].PortalChi < 0.103 && a.Lanes[2].EffectiveLambdaUV == 0 && a.Lanes[2].StableNonNegative, Detail: FormatLane(a.Lanes[0]) + " || " + FormatLane(a.Lanes[1]) + " || " + FormatLane(a.Lanes[2])},
			{Name: "one-loop r_plus RG reruns are computed for every stable correction lane", Passed: len(a.Results) == 3 && a.Results[0].Computed && a.Results[1].Computed && a.Results[2].Computed && a.Results[0].HiggsMassGeV > 331 && a.Results[2].HiggsMassGeV > 331, Detail: FormatRGResult(a.Results[0]) + " || " + FormatRGResult(a.Results[1]) + " || " + FormatRGResult(a.Results[2])},
			{Name: "capacity audit shows boundary-only B-gap σ correction does not resolve the 331 GeV one-loop tension", Passed: a.Capacity.Formalized && !a.Capacity.BoundaryCorrectionCanResolve && a.Capacity.TopSectorDominates && a.Capacity.BestStableMassGapGeV > 200 && a.Capacity.BoundaryCorrectionMovesMassGeV < 0.01, Detail: FormatCapacity(a.Capacity)},
			{Name: "firewalls block portal fitting, threshold jumps, two-loop transport, pole conversion, and final mass claims", Passed: a.Firewalls.NoObservedMassFitInserted && a.Firewalls.NoPortalCouplingFitted && a.Firewalls.NoSigmaVEVFitted && a.Firewalls.NoThresholdJumpInserted && a.Firewalls.NoTwoLoopRGExecuted && a.Firewalls.NoPoleMassConversionInserted && a.Firewalls.MajoranaEdgeRemainsConditional && a.Firewalls.NoFinalMassClaimed && !a.Firewalls.FiniteCorePolluted, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "summary activates B-gap but diagnoses that threshold/top-sector physics is still required", Passed: a.Summary.MajoranaActivationFormalized && a.Summary.TraceExtensionFormalized && a.Summary.SigmaCorrectionFormalized && a.Summary.RGRerunComputed && !a.Summary.BGapBoundaryCorrectionSolves && a.Summary.TopSectorStillDominates && !a.Summary.FinalMassClaimed && a.Summary.FirewallPreserved, Detail: FormatSummary(a.Summary)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, "Gate 312 activates the B-gap/σ idea as a conditional spectral-action correction, but it does not fit λHσ or σ's VEV to the observed Higgs mass.", "Under the inherited Gate-309 one-loop r_+ lane, boundary quartic correction alone is insufficient; the legal next step is a native B-gap threshold-matching tensor and/or top-sector tensor refinement."}}
	}}
}
