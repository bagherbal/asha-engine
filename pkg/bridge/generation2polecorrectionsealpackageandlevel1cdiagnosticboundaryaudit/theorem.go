package generation2polecorrectionsealpackageandlevel1cdiagnosticboundaryaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2PoleCorrectionSealPackageAndLevel1CDiagnosticBoundaryAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 743 — Pole-Correction Seal Package and Level-1C Diagnostic Boundary Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate743 pole-correction seal package audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate742 tree-proxy to pole firewall", Passed: a.Gate742.Inherited && NearlyEqual(a.Gate742.TreeProxyGeV, 125.38000000298437, 1e-9) && a.Gate742.CorrectionObject == "Delta_pole" && !a.Gate742.CorrectionValueAssigned && a.Gate742.Level1CAllowed && a.Gate742.Level1CRequiresExternalPackage && !a.Gate742.Level2Allowed && a.Gate742.TreeProxyNotPoleMass && strings.Contains(a.Gate742.Verdict, StatusGate742TreeProxyToPoleFirewallInherited), Detail: FormatGate742(a.Gate742)},
			{Name: "define pole-correction seal package", Passed: a.Package.FullPackage && a.Package.Count == 9 && a.Package.HasPoleObservable && a.Package.HasPoleMassConvention && a.Package.HasRGScheme && a.Package.HasRenormalizationScale && a.Package.HasLoopOrder && a.Package.HasThresholdCorrection && a.Package.HasTopYukawaInput && a.Package.HasGaugeCouplingInput && a.Package.HasUncertaintyModel && !a.Package.DeltaPoleValueAssigned && strings.Contains(a.Package.Verdict, StatusPoleCorrectionSealPackageDefined), Detail: FormatPackage(a.Package)},
			{Name: "audit correction package minimality", Passed: a.Minimality.Minimal && a.Minimality.AllRequired && a.Minimality.Count == 9 && strings.Contains(a.Minimality.Verdict, StatusCorrectionPackageMinimalityAudited), Detail: FormatMinimality(a.Minimality)},
			{Name: "define Level-1C diagnostic boundary", Passed: a.Diagnostic.Level1BAllowed && a.Diagnostic.Level1CAllowed && a.Diagnostic.Level1CDiagnosticOnly && a.Diagnostic.Level1CRequiresAll && !a.Diagnostic.Level2Allowed && strings.Contains(a.Diagnostic.Verdict, StatusLevel1CDiagnosticBoundaryDefined), Detail: FormatDiagnostic(a.Diagnostic)},
			{Name: "separate tree proxy and pole observable", Passed: NearlyEqual(a.Separation.TreeProxyGeV, 125.38000000298437, 1e-9) && a.Separation.DeltaPoleObject == "Delta_pole" && !a.Separation.PoleObservableExternallySupplied && !a.Separation.ExternalPoleObservableASHADerived && !a.Separation.TreeProxyEqualsPoleObservable && a.Separation.DiagnosticCanComputeDeltaOnlyWithPackage && strings.Contains(a.Separation.Verdict, StatusTreeProxyAndPoleObservableSeparated), Detail: FormatSeparation(a.Separation)},
			{Name: "enforce physical firewalls", Passed: !a.Firewall.FittedDeltaIsDerivedTheorem && !a.Firewall.Level1CDiagnosticIsPrediction && !a.Firewall.TreeProxyProximityIsTheorem && !a.Firewall.ExternalObservableIsDerivation && !a.Firewall.IndependentPolePrediction && a.Firewall.NoYukawaTheorem && strings.Contains(a.Firewall.Verdict, StatusPhysicalFirewallsEnforced), Detail: FormatFirewall(a.Firewall)},
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
