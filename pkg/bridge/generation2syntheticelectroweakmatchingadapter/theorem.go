package generation2syntheticelectroweakmatchingadapter

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2SyntheticElectroweakMatchingAdapterDryRunTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 synthetic electroweak matching adapter dry-run"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate505 synthetic electroweak adapter", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate504 bridge permission ledger without native electroweak rows", Passed: a.Inheritance.Executed && a.Inheritance.PermissionLedgerAccepted && a.Inheritance.BridgeInputSchemaDefined && a.Inheritance.NativeRows == 0 && a.Inheritance.BridgeRows == 6 && a.Inheritance.FormulaBridgeOnly && a.Inheritance.PermissionAllowsExplicitAdapter && !a.Inheritance.Gate504NumericalAdapterExecuted && !a.Inheritance.Gate504ObservedEWDataImported && a.Inheritance.Gate504NativeWriteBlocked, Detail: FormatInheritance(a.Inheritance)},
			{Name: "use explicit fake 3-4-5 synthetic bridge inputs only", Passed: a.Input.Synthetic && !a.Input.Observed && !a.Input.Native && a.Input.V == 2 && a.Input.G2 == 3 && a.Input.GY == 4 && a.Input.RenormalizationScaleTag != "" && a.Input.Scheme != "", Detail: FormatInput(a.Input)},
			{Name: "compute tree-level bridge outputs and photon zero mode from fake inputs", Passed: a.Output.Executed && a.Output.UsedTreeLevelFormula && nearly(a.Output.MW, 3, 1e-12) && nearly(a.Output.MZ, 5, 1e-12) && nearly(a.Output.Sin2ThetaW, 16.0/25.0, 1e-12) && nearly(a.Output.Cos2ThetaW, 9.0/25.0, 1e-12) && a.Output.MGamma == 0 && nearly(a.Output.RhoTree, 1, 1e-12), Detail: FormatOutput(a.Output)},
			{Name: "classify computed quantities as bridge-only and reject observed/native promotion", Passed: a.Adapter.Executed && a.Adapter.SyntheticOnly && !a.Adapter.ObservedDataImported && !a.Adapter.NativeDataImported && a.Adapter.InputsPositive && a.Adapter.InputsFinite && a.Adapter.ScaleSchemeMetadataPresent && a.Adapter.ComputedWithExplicitInputs && a.Adapter.WeakAngleBridgeOutputOnly && a.Adapter.WZBridgeOutputOnly && a.Adapter.PhotonZeroPreserved && a.Adapter.RhoTreeIdentityConfirmed && !a.Adapter.ObservedMassesClaimed && !a.Adapter.NativeWeakAngleDerived && !a.Adapter.NativeWZMassesDerived && !a.Adapter.NativeGaugeCouplingsDerived && !a.Adapter.NativeVEVDerived && !a.Adapter.NativeKappaPromoted && !a.Adapter.NativeYukawaTraceDerived, Detail: FormatAdapter(a.Adapter)},
			{Name: "firewall blocks observed imports and synthetic-output native writes", Passed: a.Firewall.Executed && !a.Firewall.ObservedVEVImported && !a.Firewall.ObservedGaugeCouplingsImported && !a.Firewall.ObservedWeakAngleImported && !a.Firewall.ObservedWMassImported && !a.Firewall.ObservedZMassImported && !a.Firewall.ObservedYukawaImported && !a.Firewall.NativeVEVWritten && !a.Firewall.NativeGaugeCouplingWritten && !a.Firewall.NativeWeakAngleWritten && !a.Firewall.NativeWZMassWritten && !a.Firewall.NativeKappaWritten && !a.Firewall.SyntheticOutputWrittenNative, Detail: FormatFirewall(a.Firewall)},
			{Name: "Gate506 observed electroweak comparator airlock redirect is defined", Passed: a.Next.Gate == 506, Detail: a.Next.Title + ": " + a.Next.PrimaryTask},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusGate504PermissionLedgerInherited, StatusSyntheticAdapterExecuted, StatusSyntheticInputsExplicitlyFake, StatusBridgeTreeWZComputed, StatusPhotonZeroSyntheticPreserved, StatusTreeRhoIdentitySyntheticConfirmed, StatusWeakAngleComputedAsBridgeOutput, StatusNoObservedElectroweakDataImported, StatusFailedSyntheticNotNativePrediction, StatusFailedSyntheticNotObservedMasses, StatusFailedVEVCouplingsWeakAngleNotDerived, StatusFailedKappaStillBridge, StatusFailedYukawaTraceStillSealed, StatusFirewallNoObservedDataImported, StatusFirewallSyntheticNativeWriteBlocked, StatusGate506ObservedComparatorRedirect, a.Truth}}
	}}
}
