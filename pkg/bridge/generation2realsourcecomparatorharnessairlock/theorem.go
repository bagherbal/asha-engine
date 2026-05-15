package generation2realsourcecomparatorharnessairlock

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2RealSourceComparatorExecutionHarnessAirlockPreflightTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 real-source comparator execution harness airlock preflight"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate544 comparator harness airlock", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate543 synthetic authorization manifest lock", Passed: a.Inheritance.Executed && a.Inheritance.Gate543ManifestParsed && a.Inheritance.Gate543ChecksumVerified && a.Inheritance.Gate543DryRunArmed && !a.Inheritance.Gate543LiveComparator && !a.Inheritance.Gate543RealSourceImported && a.Inheritance.Gate543NativeWriteBlocked && a.Inheritance.Gate543RedirectsToGate544, Detail: FormatInheritance(a.Inheritance)},
			{Name: "enumerate comparator execution harness contracts", Passed: a.Schema.Executed && a.Schema.RequiredRows == 16 && a.Schema.SourceRows == 7 && a.Schema.InputContractRows == 5 && a.Schema.OutputContractRows == 2 && a.Schema.QuarantineRows == 4 && a.Schema.AbortRows >= 6 && a.Schema.NativeWriteLockRows == 1, Detail: FormatSchema(a.Schema)},
			{Name: "block comparator execution in preflight", Passed: a.Guard.Executed && a.Guard.HarnessDefined && !a.Guard.RealSourceLoaded && !a.Guard.AuthorizationManifestLoaded && !a.Guard.ComparatorExecutionAuthorized && !a.Guard.ComparatorExecutionPerformed && !a.Guard.DryRunComparatorExecution && !a.Guard.LiveComparatorExecution && a.Guard.QuarantineOutputSchemaAvailable && !a.Guard.QuarantineOutputWritten && a.Guard.NativeWriteLocked && !a.Guard.NativeWriteAuthorization && a.Guard.AbortConditionsDefined && a.Guard.AbortTriggeredByNoSource, Detail: FormatGuard(a.Guard)},
			{Name: "preserve OS/Wick/Hilbert/Hamiltonian native firewalls", Passed: a.Firewall.Executed && !a.Firewall.RealSchwingerSourceImported && !a.Firewall.PhysicalSchwingerFunctionsLoaded && !a.Firewall.ConstructiveMeasureLoaded && !a.Firewall.OSPositivityCertificateLoaded && !a.Firewall.WickMapLoaded && !a.Firewall.HilbertSpaceReconstructed && !a.Firewall.HamiltonianSpectrumLoaded && !a.Firewall.ComparatorExecutionPerformed && !a.Firewall.QuarantineOutputWritten && !a.Firewall.NativeSchwingerFunctionWrite && !a.Firewall.NativeOSPositivityWrite && !a.Firewall.NativeWickWrite && !a.Firewall.NativeHilbertWrite && !a.Firewall.NativeHamiltonianWrite && !a.Firewall.NativeUnitaryDynamicsWrite && !a.Firewall.NativeGlobalCausalWrite && !a.Firewall.NativeTimeArrowWrite && !a.Firewall.NativeRegistryWritten, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(statuses(), a.Truth)}
	}}
}
