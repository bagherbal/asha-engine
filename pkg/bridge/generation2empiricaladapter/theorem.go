package generation2empiricaladapter

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2EmpiricalTextureAdapterDryRunFirewallTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 empirical texture adapter dry-run firewall test"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate455 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate454 coefficient-ray rank protocol", Passed: a.Inheritance.Executed && a.Inheritance.Gate444KGenForced && a.Inheritance.Gate445TriangleForced && a.Inheritance.Gate453EmpiricalInterfaceDefined && a.Inheritance.Gate454ProjectiveRayDOF == ProjectiveRayDOF && a.Inheritance.Gate454SpectrumOnlyRank == 1 && a.Inheritance.Gate454MinimumLocalScalars == MinLocalScalars && a.Inheritance.Gate454CPBranchTagRequired && a.Inheritance.Gate454NativeSelectorAbsent && a.Inheritance.NoEmpiricalInputsImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "defines fail-closed dry-run adapter schema", Passed: a.Schema.Executed && a.Schema.DefaultValueMode == ValueModeDummy && !a.Schema.AllowsObservedValuesByDefault && !a.Schema.AllowsNativeCoefficientExport && !a.Schema.AllowsGSTAsNativeLaw && !a.Schema.AllowsCKMPMNSAsNativeSelectors && len(a.Schema.RequiredLabels) >= 4, Detail: FormatSchema(a.Schema)},
			{Name: "accepts only native ledgers and labelled symbolic bridge comparators", Passed: a.Sieve.Executed && a.Sieve.AllowedCount == 4 && a.Sieve.NativeLedgerAllowed && a.Sieve.LocalRayDryRunAllowed && a.Sieve.OrientedRayDryRunAllowed, Detail: FormatSieve(a.Sieve)},
			{Name: "rejects all native-promotion and metadata-failure routes", Passed: a.Sieve.RejectedCount == 5 && a.Sieve.SpectrumOnlyNativePromotionRejected && a.Sieve.MissingMetadataRejected && a.Sieve.GSTNativePromotionRejected && a.Sieve.CKMPMNSNativeSelectorRejected && a.Sieve.ObservedValuesRejectedByDefault && a.Sieve.NativeCoefficientExportRejected && !a.Sieve.AnyForbiddenAccepted, Detail: FormatSieve(a.Sieve)},
			{Name: "dry-run export imports no observed values and no native coefficient", Passed: a.Export.Executed && a.Export.ActualObservedValueCount == 0 && a.Export.NativeExportCount == 0 && a.Export.BridgeExportCount >= 3 && a.Export.NativePromotionBlocked && a.Export.SchemaFailuresFailClosed, Detail: FormatExport(a.Export)},
			{Name: "13-moduli and coefficient firewall preserved", Passed: a.Firewall.Executed && a.Firewall.NoObservedMuonMassImported && a.Firewall.NoObservedCharmMassImported && a.Firewall.NoObservedYukawaImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoCurveFitPromoted && a.Firewall.NoGSTPromotion && a.Firewall.NoNativeCoefficientRayValue && a.Firewall.KGenStillForced && a.Firewall.XTriangleStillForced && a.Firewall.TextureZeroSumRuleStillBridge && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.CPOrientationStillBranchTagged && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate maps symbolic inverse and caustics", Passed: a.Next.Gate == 456, Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusGate454Inherited, StatusAdapterSchemaDefined, StatusDryRunFirewallValidated, StatusLabelledLocalRayDryRunAllowed, StatusLabelledOrientedRayDryRunAllowed, StatusNoObservedValuesImportedDefault, StatusBridgeOnlyExportsValidated, StatusEmpiricalFirewallPreserved, StatusFailedSpectrumOnlyNativePromotionRejected, StatusFailedMissingMetadataRejected, StatusFailedGSTNativePromotionRejected, StatusFailedCKMPMNSNativeSelectorRejected, StatusFailedObservedValuesRejectedByDefault, StatusFailedNativeCoefficientExportAbsent, a.Truth}}
	}}
}
