package generation2comparatoroutputreleaseairlock

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2ComparatorOutputReleaseAirlockPreflightTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 comparator-output release airlock preflight"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate546 comparator-output release airlock", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate545 quarantined synthetic comparator output", Passed: a.Inheritance.Executed && a.Inheritance.Gate545BundleParsed && a.Inheritance.Gate545ChecksumVerified && a.Inheritance.Gate545DryRunExecuted && a.Inheritance.Gate545QuarantineOutput && a.Inheritance.Gate545HumanReviewRequired && a.Inheritance.Gate545RollbackTrace && a.Inheritance.Gate545NativeWriteLocked && a.Inheritance.Gate545NoRealSource && a.Inheritance.Gate545NoPhysicalClaims && a.Inheritance.Gate545RedirectsToGate546, Detail: FormatInheritance(a.Inheritance)},
			{Name: "enumerate release-review schema rows", Passed: a.Schema.Executed && a.Schema.RequiredRows == 15 && a.Schema.HumanReviewRows > 0 && a.Schema.ReproducibilityRows > 0 && a.Schema.SourceAuthenticityRows > 0 && a.Schema.CitationScopeRows > 0 && a.Schema.NativeWriteLockRows > 0 && a.Schema.RollbackRows > 0, Detail: FormatSchema(a.Schema)},
			{Name: "block comparator-output release in preflight", Passed: a.Guard.Executed && a.Guard.AirlockDefined && a.Guard.QuarantinedComparatorPresent && !a.Guard.ReleaseManifestImported && !a.Guard.HumanReviewCompleted && !a.Guard.ReproducibilityCompleted && !a.Guard.SourceChainAuthenticated && !a.Guard.ResidualThresholdAccepted && a.Guard.PhysicalClaimDiscriminator && !a.Guard.BridgeEvidenceReleaseAllowed && !a.Guard.BridgeEvidenceReleased && a.Guard.ReleaseTargetQuarantineOnly && a.Guard.NativeWriteLocked && !a.Guard.NativeWriteAuthorization && !a.Guard.NativeRegistryWrite && a.Guard.AbortConditionsDefined && a.Guard.AbortTriggeredBySynthetic, Detail: FormatGuard(a.Guard)},
			{Name: "preserve release and native-write firewalls", Passed: a.Firewall.Executed && a.Firewall.SyntheticComparatorOutputPresent && !a.Firewall.ComparatorOutputReleased && !a.Firewall.BridgeEvidenceClaimReleased && !a.Firewall.RealSchwingerSourceImported && !a.Firewall.AuthenticatedRealSource && !a.Firewall.PhysicalSchwingerFunctionsLoaded && !a.Firewall.OSPositivityCertificateLoaded && !a.Firewall.WickMapLoaded && !a.Firewall.HilbertSpaceReconstructed && !a.Firewall.HamiltonianSpectrumLoaded && !a.Firewall.UnitaryDynamicsLoaded && !a.Firewall.GlobalCausalityLoaded && !a.Firewall.TimeArrowLoaded && !a.Firewall.NativeSchwingerFunctionWrite && !a.Firewall.NativeOSPositivityWrite && !a.Firewall.NativeWickWrite && !a.Firewall.NativeHilbertWrite && !a.Firewall.NativeHamiltonianWrite && !a.Firewall.NativeUnitaryDynamicsWrite && !a.Firewall.NativeGlobalCausalWrite && !a.Firewall.NativeTimeArrowWrite && !a.Firewall.NativeRegistryWritten, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(Statuses(), a.Truth)}
	}}
}
