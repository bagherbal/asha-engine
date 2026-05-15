package generation2observedelectroweakpreflight

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2ObservedElectroweakComparatorAirlockPreflightTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 observed electroweak comparator airlock preflight"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate506 observed electroweak preflight", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate505 synthetic adapter without observed or native data", Passed: a.Inheritance.Executed && a.Inheritance.SyntheticAdapterExecuted && a.Inheritance.SyntheticOnly && !a.Inheritance.Gate505ObservedDataImported && !a.Inheritance.Gate505NativeDataImported && a.Inheritance.Gate505NativeWriteBlocked && a.Inheritance.Gate506RedirectDefined, Detail: FormatInheritance(a.Inheritance)},
			{Name: "accept exactly one redacted bridge-only observed electroweak schema", Passed: a.Preflight.Executed && a.Preflight.AcceptedSchemaCases == 1 && a.Preflight.RejectedCases == 10 && a.Preflight.ReadyForNumericalAdapterCases == 0 && a.Preflight.AllAcceptedBridgeOnlyObserved, Detail: FormatPreflight(a.Preflight)},
			{Name: "reject missing metadata, missing rows, native promotion, kappa promotion, and observed-mass-native-input routes", Passed: a.Preflight.SwitchClosedRejected && a.Preflight.MissingVEVRejected && a.Preflight.MissingGaugeCouplingRejected && a.Preflight.MissingScaleSchemeRejected && a.Preflight.MissingSourceUncertaintyRejected && a.Preflight.ObservedMassAsNativeInputRejected && a.Preflight.WeakAngleNativePromotionRejected && a.Preflight.KappaPromotionRejected && a.Preflight.NativePromotionRejected, Detail: FormatPreflight(a.Preflight)},
			{Name: "do not run observed numerical adapter in preflight", Passed: !a.Preflight.NumericalAdapterRun && !a.Preflight.ObservedNumbersImported, Detail: FormatPreflight(a.Preflight)},
			{Name: "firewall blocks observed electroweak native writes", Passed: a.Firewall.Executed && !a.Firewall.ObservedNumbersImported && !a.Firewall.NumericalAdapterExecuted && !a.Firewall.NativeVEVWritten && !a.Firewall.NativeGaugeCouplingWritten && !a.Firewall.NativeWeakAngleWritten && !a.Firewall.NativeWZMassWritten && !a.Firewall.NativeKappaWritten && !a.Firewall.NativeRegistryWritten && !a.Firewall.NativeElectroweakPredictionMade, Detail: FormatFirewall(a.Firewall)},
			{Name: "Gate507 observed electroweak file adapter redirect is defined", Passed: a.Next.Gate == 507, Detail: a.Next.Title + ": " + a.Next.PrimaryTask},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusGate505SyntheticAdapterInherited, StatusObservedEWPreflightPolicyDefined, StatusObservedEWSchemaAccepted, StatusObservedEWPreflightValidated, StatusObservedEWBridgeAirlockAccepted, StatusObservedEWNoNumericImport, StatusObservedEWNoAdapterExecuted, StatusFailedSwitchClosed, StatusFailedMissingVEV, StatusFailedMissingGaugeCoupling, StatusFailedMissingScaleScheme, StatusFailedMissingSourceUncertainty, StatusFailedObservedMassAsNativeInput, StatusFailedWeakAngleNativePromotion, StatusFailedKappaPromotion, StatusFailedNativePromotion, StatusFailedNumericalAdapterNotRun, StatusFailedNoNativeElectroweakPrediction, StatusFirewallObservedNumbersNotImported, StatusFirewallNativeWriteBlocked, StatusGate507ObservedEWAdapterRedirect, a.Truth}}
	}}
}
