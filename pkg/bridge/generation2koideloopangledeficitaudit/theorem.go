package generation2koideloopangledeficitaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2KoideLoopAngleDeficitAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 Koide loop-angle deficit audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate586 Koide loop-angle deficit audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate585 loop-angle clue and runtime quantities", Passed: a.Runtime.Gate585BestCandidate == "1/(8π)" && a.Runtime.EpsilonRad > 0.0395 && a.Runtime.LoopUnit > 0.0397 && a.Runtime.SqrtJCKM > 0.0055, Detail: FormatRuntime(a.Runtime)},
			{Name: "define kappa_e deficit in epsilon_e=(1/(8pi))(1-kappa_e)", Passed: a.Definition.Kappa > 0.0054 && a.Definition.Kappa < 0.0056 && abs(a.Definition.ReconstructionError) < 1e-16, Detail: FormatDefinition(a.Definition)},
			{Name: "build typed kappa candidate set", Passed: a.Candidates.CandidateCount >= 15 && a.Candidates.Best.Name == "sqrt(J_CKM)" && len(a.Candidates.CertifiedCandidates) == 0, Detail: FormatCandidateSet(a.Candidates)},
			{Name: "record sqrt(J_CKM) as nearest orientation-sized clue but not source theorem", Passed: a.Orientation.SqrtJCKM.Near && !a.Orientation.SqrtJCKM.Certified && a.Orientation.NearButNotSource && !a.Orientation.PMNSRuntimeInput, Detail: FormatOrientation(a.Orientation)},
			{Name: "reject transport drift and Koide R defect as kappa source", Passed: !a.Transport.Certified && abs(a.Transport.BestTransport.RelativeResidual) > 0.99, Detail: FormatTransport(a.Transport)},
			{Name: "record alpha2 over 2pi as percent-close correction but not certified", Passed: a.Corrections.Alpha2Over2Pi.Near && !a.Corrections.Alpha2Over2Pi.Certified && !a.Corrections.Certified, Detail: FormatCorrections(a.Corrections)},
			{Name: "return no-certified-kappa-source decision", Passed: !a.Decision.CertifiedSource && a.Decision.NearClue && a.Decision.BestCandidateName == "sqrt(J_CKM)" && a.Decision.MinimalNextRequirement != "", Detail: FormatDecision(a.Decision)},
			{Name: "preserve root-trace and flavor firewalls", Passed: !a.Firewalls.DerivesKappa && !a.Firewalls.DerivesEpsilon && !a.Firewalls.DerivesKoide && !a.Firewalls.DerivesLeptonMasses && !a.Firewalls.DerivesYukawaEigenvalues && !a.Firewalls.DerivesCKM && !a.Firewalls.DerivesPMNS && !a.Firewalls.AddsNewCarrier && !a.Firewalls.PromotesObservedAsNative && a.Firewalls.PreservesGate352, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "return Gate586 final loop-angle deficit verdict", Passed: !a.Final.CandidateCertified && a.Final.NearOrientationClue && a.Final.NearCouplingClue && !a.Final.NativeDerivationCertified && a.Final.RemainingSeal != "", Detail: FormatFinal(a.Final)},
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
