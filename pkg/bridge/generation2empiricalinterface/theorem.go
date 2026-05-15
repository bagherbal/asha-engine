package generation2empiricalinterface

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2TextureZeroInvariantLedgerAllowedEmpiricalInterfaceTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 texture-zero invariant ledger allowed empirical interface"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate453 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gates 444-452 texture-zero/no-GST boundary", Passed: a.Inheritance.Executed && a.Inheritance.Gate444KGenForced && a.Inheritance.Gate445TriangleForced && a.Inheritance.Gate450TextureZeroSumRule && a.Inheritance.Gate450RatiosRequireAmplitudes && a.Inheritance.Gate451FullTrianglePreserved && a.Inheritance.Gate451NoNativePhaseRaySelector && a.Inheritance.Gate452NearestNeighborNotGauge && a.Inheritance.NoEmpiricalInputsImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "seals native invariant ledger", Passed: a.Native.Executed && !a.Native.NativeOnlyPredictsGST && contains(a.Native.PromotedStructuralObjects, "M_22=0 spectral sum rule") && contains(a.Native.QuarantinedObjects, "GST/Fritzsch relation") && contains(a.Native.QuarantinedObjects, "physical muon/charm masses"), Detail: FormatNativeLedger(a.Native)},
			{Name: "validates explicit empirical import contract", Passed: a.Contract.Executed && a.Contract.RequiresExplicitLabel && a.Contract.RequiresRenormalizationTag && a.Contract.RequiresSectorTag && !a.Contract.AllowsNativeClaim && a.Contract.RejectedPromotionCount >= 2, Detail: FormatImportContract(a.Contract)},
			{Name: "allows residual comparators but not native GST/coefficient claims", Passed: a.Residuals.Executed && a.Residuals.AllowsTextureResiduals && a.Residuals.AllowsGSTResidual && !a.Residuals.AllowsNativeGSTRatioClaim && !a.Residuals.AllowsCoefficientFittingAsNative, Detail: FormatResidualLedger(a.Residuals)},
			{Name: "sieve accepts labelled comparator use and rejects hidden promotion", Passed: a.Sieve.Executed && a.Sieve.NativeOnlyAllowed && a.Sieve.EmpiricalFitAllowed && a.Sieve.PromotionRejected && !a.Sieve.AnyForbiddenAccepted, Detail: FormatInterfaceSieve(a.Sieve)},
			{Name: "13-moduli firewall preserved", Passed: a.Firewall.Executed && a.Firewall.NoObservedMuonMassImported && a.Firewall.NoObservedCharmMassImported && a.Firewall.NoObservedYukawaImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoCurveFit && a.Firewall.NoGSTPromotion && a.Firewall.KGenStillForced && a.Firewall.Generation2ZeroStillForced && a.Firewall.XTriangleStillForced && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.GSTFritzschRelationsQuarantined && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate asks only for observability rank", Passed: a.Next.Gate == 454, Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusNativeInvariantLedgerSealed, StatusEmpiricalInterfaceDefined, StatusImportContractValidated, StatusPromotionFirewallValidated, StatusResidualComputationsQuarantined, StatusEmpiricalFirewallPreserved, StatusFailedNativeRatioNotRestored, StatusFailedGSTRequiresEmpiricalBranchInput, StatusFailedObservablePromotionRejected, a.Truth}}
	}}
}

func contains(xs []string, want string) bool {
	for _, x := range xs {
		if x == want {
			return true
		}
	}
	return false
}
