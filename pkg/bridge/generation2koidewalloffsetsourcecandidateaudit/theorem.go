package generation2koidewalloffsetsourcecandidateaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2KoideWallOffsetSourceCandidateAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 Koide wall-offset source candidate audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate585 Koide wall-offset source candidate audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate584 wall-offset ratio closure and Gate583 epsilon target", Passed: a.Runtime.Gate583EpsilonRad > 0.0395 && a.Runtime.Gate583EpsilonRad < 0.0397 && a.Runtime.Gate584SolvedEpsilonRad > 0.0395 && a.Runtime.Gate584SolvedEpsilonRad < 0.0397, Detail: FormatRuntime(a.Runtime)},
			{Name: "define source-candidate epsilon target without replacing the seal", Passed: a.Target.PrimaryEpsilonRad > 0 && a.Target.CertifiedToleranceRelative < a.Target.NearToleranceRelative && a.Target.UseForCandidateSieve != "", Detail: FormatTarget(a.Target)},
			{Name: "build typed dimensionless candidate set", Passed: a.Candidates.CandidateCount >= 20 && a.Candidates.Best.Name == "1/(8π)" && len(a.Candidates.CertifiedCandidates) == 0, Detail: FormatCandidateSet(a.Candidates)},
			{Name: "record one-over-eight-pi as near loop clue but not source theorem", Passed: a.Loop.OneOver8Pi.Near && !a.Loop.OneOver8Pi.Certified && a.Loop.NearButNotCertified && abs(a.Loop.RequiredCorrectionPct) > 0.5 && abs(a.Loop.RequiredCorrectionPct) < 0.6, Detail: FormatLoop(a.Loop)},
			{Name: "reject direct electroweak coupling candidate source", Passed: !a.Couplings.Certified && a.Couplings.BestCoupling.Name != "1/(8π)" && abs(a.Couplings.BestCoupling.RelativeResidual) > 0.1, Detail: FormatCouplings(a.Couplings)},
			{Name: "reject gauge scalar CKM residuals as epsilon source", Passed: !a.Residuals.Certified && abs(a.Residuals.BestResidual.RelativeResidual) > 0.25, Detail: FormatResiduals(a.Residuals)},
			{Name: "return no-certified-source decision", Passed: !a.Decision.CertifiedSource && a.Decision.NearClue && a.Decision.BestCandidateName == "1/(8π)" && a.Decision.MinimalNextRequirement != "", Detail: FormatDecision(a.Decision)},
			{Name: "preserve flavor and root-trace firewalls", Passed: !a.Firewalls.DerivesEpsilon && !a.Firewalls.DerivesKoide && !a.Firewalls.DerivesLeptonMasses && !a.Firewalls.DerivesYukawaEigenvalues && !a.Firewalls.DerivesCKM && !a.Firewalls.DerivesPMNS && !a.Firewalls.DerivesGenerationHierarchy && !a.Firewalls.AddsNewCarrier && !a.Firewalls.PromotesObservedAsNative && a.Firewalls.PreservesGate352, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "return Gate585 final source-candidate verdict", Passed: !a.Final.CandidateCertified && a.Final.NearLoopScaleClue && !a.Final.NativeDerivationCertified && a.Final.RemainingSeal != "", Detail: FormatFinal(a.Final)},
		}
		notes := append(Statuses(), a.Truth, a.Decision.Decision)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
