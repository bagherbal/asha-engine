package generation2level1cpoleobservablesealanddiagnosticdeltaaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2Level1CPoleObservableSealAndDiagnosticDeltaAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 745 — Level-1C Pole Observable Seal and Diagnostic Delta Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate745 pole observable diagnostic audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate744 pole-correction layer firewall", Passed: a.Gate744.Inherited && NearlyEqual(a.Gate744.TreeProxyGeV, 125.38000000298437, 1e-9) && a.Gate744.DeltaPoleObject == "Delta_pole" && a.Gate744.DeltaPoleKeptSymbolic && a.Gate744.CorrectionLayerCount == 6 && a.Gate744.HasLayeredCorrection && a.Gate744.Level1CAllowed && !a.Gate744.Level2Allowed && a.Gate744.NonFitFirewallPreserved && strings.Contains(a.Gate744.Verdict, StatusGate744PoleCorrectionLayerInherited), Detail: FormatGate744(a.Gate744)},
			{Name: "define external pole observable seal", Passed: a.Observable.Name == "PoleMassObservableSeal" && a.Observable.Object == "m_H_pole_external" && a.Observable.ExternalInput && !a.Observable.ValueSupplied && !a.Observable.NativeDerived && a.Observable.AllowsDiagnostic && strings.Contains(a.Observable.Verdict, StatusPoleObservableSealDefined), Detail: FormatObservable(a.Observable)},
			{Name: "define Level-1C diagnostic delta form", Passed: a.Delta.Name == "Delta_pole_diag" && a.Delta.Expression == "m_H_pole_external - m_H_tree_proxy" && a.Delta.RequiresPoleObservableSeal && a.Delta.RequiresTreeProxy && !a.Delta.NumericValueAssigned && !a.Delta.NativeCorrectionTheorem && !a.Delta.IndependentPrediction && a.Delta.MeasuresProxyToPoleGapOnly && strings.Contains(a.Delta.Verdict, StatusLevel1CDiagnosticDeltaFormDefined), Detail: FormatDelta(a.Delta)},
			{Name: "record layer assignment warning", Passed: a.LayerWarning.TotalCorrectionOnly && a.LayerWarning.CannotAssignPiecesWithoutPackage && a.LayerWarning.LayerCount == 6 && a.LayerWarning.ExplanatoryPackageRequired && strings.Contains(a.LayerWarning.Verdict, StatusLayerAssignmentWarningRecorded), Detail: FormatLayerWarning(a.LayerWarning)},
			{Name: "enforce non-fit firewall", Passed: !a.NonFit.FittedFromExternalMassIsDerivedTheorem && a.NonFit.ExternalObservableMeasuresGap && !a.NonFit.ExternalObservableExplainsGap && !a.NonFit.DiagnosticDeltaIsTreeToPoleTheorem && !a.NonFit.DiagnosticDeltaIsPrediction && strings.Contains(a.NonFit.Verdict, StatusNonFitFirewallEnforced), Detail: FormatNonFit(a.NonFit)},
			{Name: "record required explanatory correction package", Passed: a.Required.Count == 7 && a.Required.AllRequired && !a.Required.Native && strings.Contains(a.Required.Verdict, StatusRequiredExplanatoryCorrectionPackageRecorded), Detail: FormatRequired(a.Required)},
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
