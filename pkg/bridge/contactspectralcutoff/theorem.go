package contactspectralcutoff

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ContactSpectralCutoffIdentificationSTopBranchSelectorAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-SPECTRAL-CUTOFF-IDENTIFICATION-S-TOP-BRANCH-SELECTOR-AUDIT"
	const name = "Contact-Spectral Cutoff Identification / S_top Branch Selector Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 288 contact spectral cutoff audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "Gate 287 underdetermination is inherited", Passed: a.Inheritance.STopConstraintFormalized && a.Inheritance.FreeCutoffMomentsBlocked && !a.Inheritance.BranchPreviouslySelected, Detail: a.Inheritance.Verdict},
			{Name: "contact spectral cutoff moments are retrieved", Passed: a.Cutoff.IdentifiedAsF0F2F4 && a.Cutoff.Zeta0 == 7 && a.Cutoff.Zeta2 > 0 && a.Cutoff.Zeta4 > 0, Detail: FormatCutoff(a.Cutoff)},
			{Name: "Dirac scalar-Morita proxy and reduced a0 are explicit", Passed: a.Proxy.KappaC == 1 && a.Proxy.KappaQ == 3 && a.Proxy.A0Proxy == 4 && !a.Proxy.A0Final, Detail: FormatProxy(a.Proxy)},
			{Name: "quadratic S_top scale constraint is constructed", Passed: a.Constraint.UsesContactCutoff && a.Constraint.UsesTopologicalSTop && a.Constraint.CoeffC < 0, Detail: FormatConstraint(a.Constraint)},
			{Name: "both r branches admit positive real X", Passed: a.Sieve.BothBranchesSurvive && !a.Sieve.UniqueBranchSelected && len(a.Sieve.SurvivingBranches) == 2, Detail: FormatSieve(a.Sieve)},
			{Name: "contact cutoff locks total reduced trace moments but not branch distribution", Passed: a.MomentLock.ReducedProxyLocked && a.MomentLock.BranchIndependentD2D4 && a.MomentLock.BranchDependentX && !a.MomentLock.PhysicalHeatKernelLock, Detail: FormatMomentLock(a.MomentLock)},
			{Name: "Higgs observable remains unclaimed", Passed: !a.Higgs.HiggsMassRatioClaimed && !a.Higgs.HeatKernelProjectionDerived && !a.Higgs.DimensionlessObservableDefined, Detail: FormatHiggs(a.Higgs)},
			{Name: "firewalls are preserved", Passed: a.Firewalls.DoesNotDiscardSurvivingBranch && a.Firewalls.DoesNotClaimHiggsPrediction && !a.Firewalls.FiniteCorePolluted, Detail: FormatFirewalls(a.Firewalls)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth, FormatSummary(a.Summary)}}
	}}
}
