package generation2treeproxytopolemasscorrectiondependencyandfirewallaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2TreeProxyToPoleMassCorrectionDependencyAndFirewallAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 742 — Tree Proxy to Pole-Mass Correction Dependency and Firewall Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate742 tree-proxy to pole audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate741 Level-1B tree proxy", Passed: a.Gate741.Inherited && a.Gate741.Level1B && a.Gate741.NotPoleMass && a.Gate741.NotIndependentPrediction && NearlyEqual(a.Gate741.TreeProxyGeV, 125.38000000298437, 1e-9) && strings.Contains(a.Gate741.Verdict, StatusGate741Level1BTreeProxyInherited), Detail: FormatGate741(a.Gate741)},
			{Name: "define formal pole correction object", Passed: a.Correction.Name == "Delta_pole" && !a.Correction.ValueAssigned && a.Correction.RequiresPoleConvention && a.Correction.RequiresExternalCorrection && strings.Contains(a.Correction.Verdict, StatusPoleCorrectionObjectDefined), Detail: FormatCorrection(a.Correction)},
			{Name: "list required correction ingredients", Passed: a.Ingredients.AllListed && a.Ingredients.Count == 10 && a.Ingredients.NoNativeRG && a.Ingredients.NoNativeTopGauge && strings.Contains(a.Ingredients.Verdict, StatusRequiredCorrectionIngredientsListed), Detail: FormatIngredients(a.Ingredients)},
			{Name: "enforce tree proxy versus pole firewall", Passed: !a.Firewall.TreeProxyEqualsPoleMass && !a.Firewall.NearNumericalProximityIsPrediction && a.Firewall.PoleObservableNeedsLoopThreshold && a.Firewall.TreeProxyConventionLevel && strings.Contains(a.Firewall.Verdict, StatusTreeProxyVersusPoleFirewallEnforced), Detail: FormatFirewall(a.Firewall)},
			{Name: "audit seal inheritance", Passed: a.Seals.Explicit && a.Seals.TotalCount == 17 && a.Seals.IncludesRGScheme && a.Seals.IncludesPoleMass && a.Seals.IncludesThreshold && a.Seals.IncludesTopYukawa && a.Seals.IncludesGauge && strings.Contains(a.Seals.Verdict, StatusSealInheritanceAudited), Detail: FormatSeals(a.Seals)},
			{Name: "refine forecast levels", Passed: a.Forecast.Level1BAllowed && a.Forecast.Level1CAllowed && a.Forecast.Level1CDiagnosticOnly && a.Forecast.Level1CRequiresExternal && !a.Forecast.Level2Allowed && strings.Contains(a.Forecast.Verdict, StatusForecastLevelsRefined), Detail: FormatForecast(a.Forecast)},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := append([]string{a.Truth}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
