package twoloopintegration

import "github.com/bagherbal/asha-engine/pkg/theorem"

func SealedTwoLoopRGIntegrationMatchingEnvelopeTheorem() theorem.Theorem {
	const id = "BRIDGE-SEALED-TWO-LOOP-RG-INTEGRATION-MATCHING-ENVELOPE-AUDIT"
	const name = "Sealed two-loop RG integration / matching-correction uncertainty envelope audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 214 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 213 ThresholdSpectrumSeal and two-loop warning are inherited", Passed: a.Gate213.Gate213Inherited && a.Gate213.ThresholdSpectrumSealInherited && a.Gate213.MatchingCorrectionsObstructed && a.Gate213.TwoLoopWarningInherited, Detail: FormatGate213(a.Gate213)},
			{Name: "sealed no-Yukawa two-loop spectrum is constructed without finite-core promotion", Passed: a.Spectrum.Row1Rep == "(1,3,Y=1)" && a.Spectrum.Row2Rep == "(8,2,Y=1/2)" && !a.Spectrum.YukawaIncluded && !a.Spectrum.FiniteDerived, Detail: FormatSpectrum(a.Spectrum)},
			{Name: "central two-loop RK4/Newton fit converges to u*=1", Passed: a.Central.Converged && a.Central.ResidualNorm < 1e-8 && a.Central.ScaleOrdered && a.Central.DistinctThresholds && a.Central.SubPlanck && a.Central.PositiveToBoundary && a.Central.NoLandauBelowPlanck, Detail: FormatSolution(a.Central)},
			{Name: "two-loop fit is allowed to reorder the two sealed thresholds", Passed: a.Central.LB2 < a.Central.LB1 && a.Central.DeltaL > 0 && a.Central.DeltaL < 0.1, Detail: FormatSolution(a.Central)},
			{Name: "matching uncertainty envelope is explicit phenomenological proxy only", Passed: a.Envelope.Status == MatchingEnvelopeStatus && a.Envelope.CasesAudited == 8 && a.Envelope.ConvergedCases == 8 && a.Envelope.EpsilonU > 0 && a.Envelope.MStarMinGeV < a.Central.MStarGeV && a.Envelope.MStarMaxGeV > a.Central.MStarGeV, Detail: FormatEnvelope(a.Envelope)},
			{Name: "derived matching corrections remain obstructed", Passed: a.MatchingAudit.Gate213MatchingObstructionInherited && !a.MatchingAudit.NativeDeltaMatchRowsDerived && !a.MatchingAudit.CanonicalSubtractionSchemeDerived && !a.MatchingAudit.EnvelopeImportedAsFiniteCore && a.MatchingAudit.EnvelopeUsedAsPhenomenologicalProxy, Detail: FormatMatching(a.MatchingAudit)},
			{Name: "firewalls remain closed", Passed: a.Firewall.ThresholdSpectrumSealInherited && a.Firewall.EmpiricalCarrierSealInherited && a.Firewall.LeptoquarkDynamicsSealInherited && a.Firewall.EmpiricalLedgerQuarantined && !a.Firewall.MatchingCorrectionsDerived && !a.Firewall.MatchingEnvelopeFiniteCore && !a.Firewall.YukawaMatricesImported && !a.Firewall.PhysicalPredictionClaimed && !a.Firewall.FiniteMassPredictionClaimed && !a.Firewall.ProtonLifetimeComputed && !a.Firewall.OneLoopScalesOverwrittenAsCore, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Checks: checks, Notes: []string{a.TruthStatement, "CONDITIONAL_PHENOMENOLOGY: Gate 214 produces sealed two-loop numerical fits and matching-envelope error bars only; no finite-core mass or matching theorem is claimed."}}
	}}
}
