package generation2observedpreflight

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2ObservedRankCompleteComparatorPreflightAirlockNonComputationAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 observed rank-complete comparator preflight airlock non-computation audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate469 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits synthetic socket and empirical airlock", Passed: a.Inheritance.Executed && a.Inheritance.Gate468SyntheticSocketValidated && a.Inheritance.Gate465AirlockAvailable && a.Inheritance.Gate464DUDSocketAvailable && a.Inheritance.NativeRegistryClean, Detail: FormatInheritance(a.Inheritance)},
			{Name: "accepts only bridge-only observed preflight schemas", Passed: a.Preflight.Executed && a.Preflight.AcceptedSchemaCases == 1 && a.Preflight.AllAcceptedBridgeOnlyObserved, Detail: FormatPreflight(a.Preflight)},
			{Name: "blocks redacted or incomplete observed d_ud execution", Passed: !a.Preflight.DUDComputed && a.Preflight.MissingNumericRejected && a.Preflight.MissingISpecRejected && a.Preflight.MissingIKRejected && a.Preflight.MissingBranchRejected && a.Preflight.MixedScaleRejected && a.Preflight.MissingUncertaintyRejected, Detail: FormatPreflight(a.Preflight)},
			{Name: "rejects Cabibbo-as-coordinate and native-promotion routes", Passed: a.Preflight.CabibboRayRejected && a.Preflight.NativePromotionRejected && a.Preflight.CKMNativePredictionRejected && a.Preflight.SwitchClosedRejected, Detail: FormatPreflight(a.Preflight)},
			{Name: "preserves native theorem firewall", Passed: a.Firewall.Executed && !a.Firewall.ObservedRowsNative && !a.Firewall.DUDNativePrediction && !a.Firewall.CKMNativePrediction && !a.Firewall.CKMMatrixConstructed && !a.Firewall.CKMEntryComputed && !a.Firewall.CabibboUsedAsRayInput && !a.Firewall.NativeRegistryWritten && a.Firewall.KGenStillForced && a.Firewall.XTriangleStillForced && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate explicit data-file run", Passed: a.Next.Gate == 470, Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusPreflightPolicyDefined, StatusSchemaAccepted, StatusPreflightValidated, StatusFailedMissingNumeric, StatusFailedCabibboAsRayInput, StatusFailedCKMPrediction, StatusFirewallPreserved, a.Truth}}
	}}
}
