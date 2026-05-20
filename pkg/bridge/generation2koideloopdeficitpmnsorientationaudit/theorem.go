package generation2koideloopdeficitpmnsorientationaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2KoideLoopDeficitPMNSOrientationAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 Koide loop-deficit PMNS orientation audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate587 PMNS orientation audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate586 kappa deficit and runtime correction clues", Passed: a.Runtime.Kappa > 0.0054 && a.Runtime.Kappa < 0.0056 && a.Runtime.SqrtJCKM > 0.0055 && a.Runtime.Alpha2Over2Pi > 0.0053, Detail: FormatRuntime(a.Runtime)},
			{Name: "import NuFIT 6.0 PMNS data with version, ordering, convention, and uncertainties", Passed: a.PMNSInput.SourceVersion != "" && a.PMNSInput.MassOrdering == "Normal Ordering best fit" && a.PMNSInput.Sin2Theta13 > 0.022 && a.PMNSInput.DeltaCPDeg == 212, Detail: FormatPMNSInput(a.PMNSInput)},
			{Name: "compute PMNS Jarlskog and typed invariants", Passed: a.Invariants.JPMNS < 0 && a.Invariants.SqrtAbsJ > 0.13 && a.Invariants.S13Squared == a.PMNSInput.Sin2Theta13, Detail: FormatInvariants(a.Invariants)},
			{Name: "build PMNS candidate set and reject direct PMNS orientation match", Passed: a.Candidates.CandidateCount >= 12 && a.Candidates.SqrtJPMNS.RelativeResidual > 20 && a.Candidates.BestDirectPMNS.Name == "|J_PMNS|", Detail: FormatCandidateSet(a.Candidates)},
			{Name: "record PMNS-assisted alpha2/(2pi)/c13 as better typed candidate than sqrt(J_CKM) but not certified", Passed: a.Candidates.BestPMNSAssisted.Name == "alpha_2(M_Z)/(2π)/c13" && a.CKM.PMNSAssistedBetterThanSqrtJCKM && !a.Candidates.BestPMNSAssisted.Certified, Detail: FormatCKM(a.CKM)},
			{Name: "propagate PMNS uncertainty: broad |J_PMNS| can cover kappa, but no candidate certifies", Passed: !a.Uncertainty.CertifiedUnderUncertainty && a.Uncertainty.AnyCandidateCovers && !a.Uncertainty.Alpha2Over2PiDivC13.CoversKappa, Detail: FormatUncertainty(a.Uncertainty)},
			{Name: "preserve CKM midpoint only as numerical clue, not source", Passed: a.CKM.MidpointStillClosestNumeric && !a.CKM.CKMAlpha2Midpoint.Certified, Detail: FormatCKM(a.CKM)},
			{Name: "return kappa as remaining environmental seal", Passed: !a.Decision.AnyCandidateCertified && a.Decision.KappaRemainsSeal && a.Decision.PMNSProducesBetterTypedCandidate && a.Decision.CKMMidpointSurvives, Detail: FormatDecision(a.Decision)},
			{Name: "preserve ASHA flavor and root-trace firewalls", Passed: !a.Firewalls.DerivesKappa && !a.Firewalls.DerivesEpsilon && !a.Firewalls.DerivesKoide && !a.Firewalls.DerivesPMNS && !a.Firewalls.DerivesNeutrinoMasses && !a.Firewalls.DerivesFlavorTexture && !a.Firewalls.DerivesChargedLeptonMasses && !a.Firewalls.PromotesObservedAsNative && !a.Firewalls.AddsNewCarrier && a.Firewalls.PreservesGate352, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "return Gate587 final verdict", Passed: !a.Final.AnyCertified && a.Final.BestPMNSAssistedName == "alpha_2(M_Z)/(2π)/c13" && a.Final.RemainingSeal != "", Detail: FormatFinal(a.Final)},
		}
		notes := append(Statuses(), a.Truth, a.Decision.Decision)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: notes}
	}}
}
