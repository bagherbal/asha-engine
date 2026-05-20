package generation2polecorrectionlayerdecompositionandnonfitfirewallaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2PoleCorrectionLayerDecompositionAndNonFitFirewallAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 744 — Pole-Correction Layer Decomposition and Non-Fit Firewall Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate744 pole-correction layer decomposition audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate743 pole-correction seal package", Passed: a.Gate743.Inherited && NearlyEqual(a.Gate743.TreeProxyGeV, 125.38000000298437, 1e-9) && a.Gate743.DeltaPoleObject == "Delta_pole" && !a.Gate743.DeltaPoleValueAssigned && a.Gate743.FullCorrectionPackageDefined && a.Gate743.Level1CAllowed && a.Gate743.Level1CDiagnosticOnly && !a.Gate743.Level2Allowed && strings.Contains(a.Gate743.Verdict, StatusGate743PoleCorrectionSealPackageInherited), Detail: FormatGate743(a.Gate743)},
			{Name: "keep Delta_pole symbolic", Passed: a.DeltaPole.Name == "Delta_pole" && a.DeltaPole.Expression == "m_H_pole - m_H_tree_proxy" && !a.DeltaPole.ValueAssigned && !a.DeltaPole.ExternalObservableSupplied && !a.DeltaPole.CorrectionPackageSupplied && strings.Contains(a.DeltaPole.Verdict, StatusDeltaPoleKeptSymbolic), Detail: FormatDeltaPole(a.DeltaPole)},
			{Name: "define correction layer decomposition", Passed: a.Decomposition.Count == 6 && a.Decomposition.AllRequired && !a.Decomposition.AnyNativeDerived && !a.Decomposition.CompressibleToFit && strings.Contains(a.Decomposition.FormalExpression, "Delta_RG") && strings.Contains(a.Decomposition.FormalExpression, "Delta_uncertainty") && strings.Contains(a.Decomposition.Verdict, StatusCorrectionLayerDecompositionDefined), Detail: FormatDecomposition(a.Decomposition)},
			{Name: "audit correction layer minimality", Passed: a.Minimality.Minimal && a.Minimality.AllRequired && a.Minimality.Count == 6 && strings.Contains(a.Minimality.Verdict, StatusCorrectionLayerMinimalityAudited), Detail: FormatMinimality(a.Minimality)},
			{Name: "enforce non-fit firewall", Passed: !a.NonFit.ObservedMinusProxyIsDerivedTheorem && a.NonFit.ExternalDiagnosticAllowed && !a.NonFit.ExternalDiagnosticIsPrediction && a.NonFit.DeltaPoleKeptLayered && a.NonFit.SingleFittedNumberLosesTypeInfo && strings.Contains(a.NonFit.Verdict, StatusNonFitFirewallEnforced), Detail: FormatNonFit(a.NonFit)},
			{Name: "preserve forecast boundary", Passed: a.Classification.Level1BAllowed && a.Classification.Level1CAllowed && !a.Classification.Level2Allowed && a.Classification.TreeProxyLevel == "Level-1B" && strings.Contains(a.Classification.DeltaPoleStatus, "multi-layer") && strings.Contains(a.Classification.Verdict, StatusForecastBoundaryPreserved), Detail: FormatClassification(a.Classification)},
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
