package generation2syntheticauthorizationmanifestadapter

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2SyntheticComparatorAuthorizationManifestAdapterDryRunTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 synthetic comparator authorization manifest adapter dry run"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate543 synthetic authorization manifest adapter", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate542 authorization-manifest airlock", Passed: a.Inheritance.Executed && a.Inheritance.Gate542RowsEnumerated && a.Inheritance.Gate542ComparatorBlocked && a.Inheritance.Gate542NativeWriteLocked && a.Inheritance.Gate542NoRealSource && a.Inheritance.Gate542RedirectsToGate543, Detail: FormatInheritance(a.Inheritance)},
			{Name: "parse all 14 synthetic authorization manifest rows", Passed: a.Import.Loaded && a.Import.AcceptedRows == 14 && a.Import.RejectedRows == 0 && len(a.Import.MissingRows) == 0 && len(a.Import.DuplicateRows) == 0 && a.Import.ChecksumVerified, Detail: FormatImport(a.Import)},
			{Name: "enforce synthetic dry-run metadata sieve", Passed: a.Import.AllBridgeOnly && a.Import.AllComparatorOnly && a.Import.AllQuarantineOnly && a.Import.AllDryRunOnly && a.Import.AllSynthetic && a.Import.AllNoTheorem && a.Import.AllSourceTagged && a.Import.AllConventionTagged && !a.Import.AnyNativePromotion && !a.Import.AnyNativeWrite && !a.Import.AnyPhysicalClaim && !a.Import.AnyObservedClaim, Detail: FormatImport(a.Import)},
			{Name: "arm only bridge-quarantine dry-run authorization", Passed: a.Authorization.DryRunAuthorizationArmed && !a.Authorization.LiveComparatorAuthorization && !a.Authorization.ComparatorExecutionPerformed && a.Authorization.BridgeQuarantineOnly && a.Authorization.QuarantineOutputTarget && a.Authorization.NativeWriteLocked && !a.Authorization.NativeWriteAuthorization && a.Authorization.SyntheticManifestOnly && !a.Authorization.CanImportRealSource, Detail: FormatAuthorization(a.Authorization)},
			{Name: "preserve physical-source and native-write firewalls", Passed: a.Firewall.Executed && !a.Firewall.RealSchwingerSourceImported && !a.Firewall.ObservedCorrelationImported && !a.Firewall.ConstructiveMeasureImported && !a.Firewall.PhysicalOSCertificateImported && !a.Firewall.PhysicalWickMapImported && !a.Firewall.PhysicalHamiltonianImported && !a.Firewall.ComparatorExecutionPerformed && !a.Firewall.LiveComparatorAuthorized && !a.Firewall.NativeSchwingerFunctionWrite && !a.Firewall.NativeEuclideanMeasureWrite && !a.Firewall.NativeOSPositivityWrite && !a.Firewall.NativeWickWrite && !a.Firewall.NativeHilbertWrite && !a.Firewall.NativeHamiltonianWrite && !a.Firewall.NativeUnitaryDynamicsWrite && !a.Firewall.NativeGlobalCausalWrite && !a.Firewall.NativeTimeArrowWrite && !a.Firewall.NativeRegistryWritten, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(statuses(), a.Truth)}
	}}
}
