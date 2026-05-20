package generation2level1cpoleobservablesealanddiagnosticdeltaaudit

import (
	"strings"
	"testing"
)

func TestGate745DiagnosticDeltaAndObservableSeal(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate744.Inherited || !NearlyEqual(a.Gate744.TreeProxyGeV, 125.38000000298437, 1e-9) || a.Gate744.DeltaPoleObject != "Delta_pole" || !a.Gate744.DeltaPoleKeptSymbolic || a.Gate744.CorrectionLayerCount != 6 || !a.Gate744.Level1CAllowed || a.Gate744.Level2Allowed {
		t.Fatalf("bad Gate744 inheritance: %+v", a.Gate744)
	}
	if a.Observable.Name != "PoleMassObservableSeal" || a.Observable.Object != "m_H_pole_external" || !a.Observable.ExternalInput || a.Observable.ValueSupplied || a.Observable.NativeDerived || !a.Observable.AllowsDiagnostic {
		t.Fatalf("bad pole observable seal: %+v", a.Observable)
	}
	if a.Delta.Name != "Delta_pole_diag" || a.Delta.Expression != "m_H_pole_external - m_H_tree_proxy" || !a.Delta.RequiresPoleObservableSeal || !a.Delta.RequiresTreeProxy || a.Delta.NumericValueAssigned || a.Delta.NativeCorrectionTheorem || a.Delta.IndependentPrediction || !a.Delta.MeasuresProxyToPoleGapOnly {
		t.Fatalf("bad diagnostic delta: %+v", a.Delta)
	}
}

func TestGate745LayerWarningPackageAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.LayerWarning.TotalCorrectionOnly || !a.LayerWarning.CannotAssignPiecesWithoutPackage || a.LayerWarning.LayerCount != 6 || !a.LayerWarning.ExplanatoryPackageRequired {
		t.Fatalf("bad layer assignment warning: %+v", a.LayerWarning)
	}
	if a.NonFit.FittedFromExternalMassIsDerivedTheorem || !a.NonFit.ExternalObservableMeasuresGap || a.NonFit.ExternalObservableExplainsGap || a.NonFit.DiagnosticDeltaIsTreeToPoleTheorem || a.NonFit.DiagnosticDeltaIsPrediction {
		t.Fatalf("bad non-fit firewall: %+v", a.NonFit)
	}
	if a.Required.Count != 7 || !a.Required.AllRequired || a.Required.Native {
		t.Fatalf("bad required package: %+v", a.Required)
	}
	for _, want := range RequiredExplanatoryCorrectionPackageLabels {
		found := false
		for _, got := range a.Required.Labels {
			if got == want {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("missing required correction label %s", want)
		}
	}
	res := Generation2Level1CPoleObservableSealAndDiagnosticDeltaAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
