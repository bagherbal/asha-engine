package generation2basisinvariance

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2FamilyBasisInvarianceTextureGaugeArtifactAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 family basis-invariance texture gauge-artifact audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate452 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate 451 no-selector result", Passed: a.Inheritance.Executed && a.Inheritance.Gate444KGenForced && a.Inheritance.Gate445TriangleForced && a.Inheritance.Gate450TextureZeroSumRule && a.Inheritance.Gate451FullTrianglePreserved && a.Inheritance.Gate451NoNativePhaseRaySelector && a.Inheritance.Gate451GSTFritzschQuarantined && a.Inheritance.NoEmpiricalInputsImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "computes native K-preserving basis group", Passed: a.BasisAudit.Executed && a.BasisAudit.GeneralUnitaryRejected && a.BasisAudit.AllNativeAllowedPreserve13 && !a.BasisAudit.AnyNativeAllowedDeletes13, Detail: FormatBasisAudit(a.BasisAudit)},
			{Name: "proves support cannot rephase to nearest-neighbor chain", Passed: a.Support.Executed && a.Support.SupportPatternInvariant && !a.Support.CanRephaseToNN && a.Support.EdgeCountTriangle == 3 && a.Support.EdgeCountNearestNeighbor == 2 && a.Support.TriangleCycleCount == 1 && a.Support.NearestNeighborCycleCount == 0, Detail: FormatSupportAudit(a.Support)},
			{Name: "separates triangle and chain by invariant ledgers", Passed: a.Spectral.Executed && !a.Spectral.SameInvariantClass && a.Spectral.TriangleDetLiftCoeff == 2 && a.Spectral.NearestNeighborDetLiftCoeff == 0 && a.Spectral.TriangleCommutatorNorm2 == 12 && a.Spectral.NearestNeighborCommutatorNorm2 == 4, Detail: FormatSpectralAudit(a.Spectral)},
			{Name: "rejects nearest-neighbor branch as native gauge artifact", Passed: a.Verdict.Executed && !a.Verdict.NearestNeighborCanBeNativeGaugeArtifact && a.Verdict.RequiresNonNativeGeneralUnitary && a.Verdict.KGenAddressDestroyed && a.Verdict.TextureZeroAddressDestroyed && a.Verdict.GSTFritzschStillEmpiricalAssumption && !a.Verdict.ReevaluateRatios, Detail: FormatGaugeArtifactVerdict(a.Verdict)},
			{Name: "13-moduli firewall preserved", Passed: a.Firewall.Executed && a.Firewall.NoObservedMuonMassImported && a.Firewall.NoObservedCharmMassImported && a.Firewall.NoObservedYukawaImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoCurveFit && a.Firewall.KGenStillForced && a.Firewall.Generation2ZeroStillForced && a.Firewall.XTriangleStillForced && a.Firewall.YPhaseStillQuarantined && a.Firewall.SectorCoefficientsStillSealed && a.Firewall.GSTFritzschRelationsQuarantined && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate defines explicit empirical interface only", Passed: a.Next.Gate == 453, Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: checks, Notes: []string{StatusFailedNoBasisSuppression, StatusFailedNNTextureNotGaugeEquivalent, StatusFailedGeneralFamilyRotationBreaksKAddress, StatusTextureGaugeArtifactQuarantined, StatusEmpiricalFirewallPreserved, a.Truth}}
	}}
}
