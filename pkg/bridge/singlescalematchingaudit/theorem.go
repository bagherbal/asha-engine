package singlescalematchingaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func SingleScaleDegenerateLimitMatchingAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-SINGLE-SCALE-DEGENERATE-LIMIT-MATCHING-AUDIT"
	const name = "Single-scale degenerate-limit matching audit / global two-loop class scan"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate 215 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 214 two-loop integration and matching envelope are inherited", Passed: a.Gate214.Gate214Inherited && a.Gate214.ThresholdSpectrumSealInherited && a.Gate214.MatchingEnvelopeInherited && a.Gate214.CentralTwoLoopConverged, Detail: FormatGate214(a.Gate214)},
			{Name: "all 22 unordered Gate-211 pair classes are scanned", Passed: a.GlobalScan.ClassesAudited == 22 && a.Config.ClassesExpected == 22 && a.Firewall.All22ClassesAudited, Detail: FormatGlobal(a.GlobalScan)},
			{Name: "forced single-threshold two-loop scan has a best ranked residual", Passed: len(a.Fits) == 22 && a.GlobalScan.BestClassRank > 0 && a.GlobalScan.BestMaxResidual > 0 && a.GlobalScan.BestResidualOverEpsilon > 0, Detail: FormatFit(a.Fits[0])},
			{Name: "the Gate-211 ranked witness is plausible inside the loop-factor envelope", Passed: a.Fits[0].ClassRank == 1 && a.Fits[0].MatchingPlausible && a.Fits[0].MaxAbsResidual < a.Config.EpsilonU, Detail: FormatFit(a.Fits[0])},
			{Name: "matching residuals are required corrections, not derived counterterms", Passed: a.MatchingAudit.Gate214MatchingEnvelopeInherited && !a.MatchingAudit.NativeDeltaMatchRowsDerived && !a.MatchingAudit.HeatKernelMatchingMapDerived && !a.MatchingAudit.CanonicalSubtractionScheme && a.MatchingAudit.EnvelopeUsedAsProxy && !a.MatchingAudit.ResidualInterpretedAsDerived, Detail: FormatMatching(a.MatchingAudit)},
			{Name: "the scan stays sealed and imports no Yukawa or proton-lifetime machinery", Passed: a.Firewall.ThresholdSpectrumSealInherited && a.Firewall.EmpiricalCarrierSealInherited && a.Firewall.LeptoquarkDynamicsSealInherited && a.Firewall.EmpiricalLedgerQuarantined && !a.Firewall.SingleScaleForcedAsFiniteCore && !a.Firewall.MatchingCorrectionsDerived && !a.Firewall.MatchingResidualPromoted && !a.Firewall.YukawaMatricesImported && !a.Firewall.PhysicalPredictionClaimed && !a.Firewall.ProtonLifetimeComputed, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Phenomenology, Checks: checks, Notes: []string{a.TruthStatement, "CONDITIONAL_PHENOMENOLOGY: Gate 215 ranks forced single-scale two-loop spectra by required matching residuals; it does not derive matching corrections or physical masses."}}
	}}
}
