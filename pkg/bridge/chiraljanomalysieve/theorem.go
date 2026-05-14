package chiraljanomalysieve

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ChiralJStructureAnomalySieveAsymmetricTraceAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-CHIRAL-J-STRUCTURE-ANOMALY-SIEVE-ASYMMETRIC-TRACE-AUDIT"
	const name = "Chiral/J-Structure Anomaly Sieve / Asymmetric Trace Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 289 chiral/J anomaly sieve audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 288 branch masking is inherited", Passed: a.Inheritance.ContactCutoffIdentified && a.Inheritance.BothBranchesSurvived && !a.Inheritance.BranchSelected, Detail: a.Inheritance.Verdict},
			{Name: "gamma proxy is formalized but not promoted to completed physical gamma", Passed: a.Gamma.ProxyDefined && !a.Gamma.PhysicalGammaDerived && !a.Gamma.CompletedHFDerived, Detail: FormatGamma(a.Gamma)},
			{Name: "chiral gamma traces are computed and branch-blind", Passed: a.Chiral.AllGammaD2Zero && a.Chiral.AllGammaD4Zero && !a.Chiral.BranchSensitiveViaGamma, Detail: FormatChiral(a.Chiral)},
			{Name: "sector-projected traces expose branch sensitivity only as diagnostics", Passed: a.Sector.BranchSensitive && !a.Sector.NativeSelectionFunctional && a.Sector.SelectedBranch == "", Detail: FormatSector(a.Sector)},
			{Name: "physical J and opposite action remain missing", Passed: a.J.AntiLinearCandidateAvailable && !a.J.PhysicalJDerived && !a.J.OppositeActionConstructed, Detail: FormatJ(a.J)},
			{Name: "anomaly conditions do not select r branch", Passed: !a.Anomaly.FullHyperchargeRepresentationDerived && !a.Anomaly.AnomalyEquationsDependOnR && !a.Anomaly.CanEliminateBranch, Detail: FormatAnomaly(a.Anomaly)},
			{Name: "both amplitude branches survive the asymmetric trace audit", Passed: !a.Sieve.UniqueBranchSelected && len(a.Sieve.SurvivingBranches) == 2, Detail: FormatSieve(a.Sieve)},
			{Name: "Higgs prediction remains firewalled", Passed: !a.Higgs.HiggsPredictionClaimed && !a.Higgs.HeatKernelProjectionDerived && !a.Higgs.DimensionlessObservableDefined, Detail: FormatHiggs(a.Higgs)},
			{Name: "firewalls are preserved", Passed: a.Firewalls.DoesNotDiscardSurvivingBranch && a.Firewalls.DoesNotClaimHiggsPrediction && !a.Firewalls.FiniteCorePolluted, Detail: FormatFirewalls(a.Firewalls)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary)}}
	}}
}
