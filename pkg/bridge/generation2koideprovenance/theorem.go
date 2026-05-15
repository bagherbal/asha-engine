package generation2koideprovenance

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2KoideConstraintProvenanceTopologicalBaselineTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 Koide constraint provenance topological baseline"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate485 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate484 without promoting its empirical shadow", Passed: a.Inheritance.Executed && a.Inheritance.Gate480NullConeNative && a.Inheritance.Gate484C3BasisValidated && a.Inheritance.Gate484ChargedLeptonKoideShadowFound && a.Inheritance.ObservedMassesRemainBridgeData && a.Inheritance.NativeRegistryClean, Detail: FormatInheritance(a.Inheritance)},
			{Name: "proves C3 democratic/phase-plane orthogonality", Passed: a.Basis.Executed && a.Basis.DemocraticPhaseOrthogonality && a.Basis.PhaseNormEqualsThreeHalves && a.Basis.PhaseNormIndependentOfPsi && a.Basis.BasisNativeC3Orbit && a.Basis.NoEmpiricalMassesUsed, Detail: FormatBasis(a.Basis)},
			{Name: "derives R/S=sqrt(2) from the null boundary", Passed: a.Derivation.Executed && a.Derivation.PositiveFutureBranch && a.Derivation.NullForcesRatio && a.Derivation.KoideEquivalent && a.Derivation.ScaleFree && a.Derivation.PhaseFree, Detail: FormatDerivation(a.Derivation)},
			{Name: "collapses exactly one C3 shape coordinate", Passed: a.Collapse.Executed && a.Collapse.C3RawShadowDOF == 3 && a.Collapse.NullConstrainedShadowDOF == 2 && a.Collapse.CollapsedShapeDOF == 1 && a.Collapse.ScaleStillFree && a.Collapse.PsiStillFree && !a.Collapse.AbsoluteMassesDerived && a.Collapse.ChargedLeptonBridgeCompatible && !a.Collapse.FullFlavorModuliCollapsed && a.Collapse.NativeFlavorDimAfter == NativeFlavorDim && a.Collapse.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatCollapse(a.Collapse)},
			{Name: "keeps quark dressing and mixing outside the baseline", Passed: a.Sector.Executed && a.Sector.ColorlessLeptonBaselineEligible && !a.Sector.QuarkBaselineEligible && a.Sector.QuarkColorDressingDeclared && !a.Sector.PhysicalLeptonMassesDerived && !a.Sector.PhysicalQuarkMassesDerived && !a.Sector.CKMConstructed && !a.Sector.PMNSConstructed, Detail: FormatSector(a.Sector)},
			{Name: "preserves 13-moduli firewall", Passed: a.Firewall.Executed && !a.Firewall.ObservedMassImportedForProof && !a.Firewall.CKMImported && !a.Firewall.PMNSImported && a.Firewall.NullC3RatioNativeBaseline && !a.Firewall.KoideAsPhysicalMassPrediction && !a.Firewall.LeptonMassesDerived && !a.Firewall.QuarkMassesDerived && !a.Firewall.PhasePsiSelected && !a.Firewall.SectorPerturbationsNative && !a.Firewall.CKMMatrixConstructed && !a.Firewall.PMNSMatrixConstructed && !a.Firewall.NativeRegistryWritten && a.Firewall.NativeFlavorDimAfter == NativeFlavorDim && a.Firewall.KXYCoeffDimAfter == KXYCoeffDim, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{StatusGate484Inherited, StatusC3ShadowBasisProved, StatusNullKoideRatioDerived, StatusKoideQDerived, StatusLeptonBaselineCompatible, StatusFailedEmpiricalFitRejected, StatusFailedMassesNotDerived, StatusFailedPhaseNotSelected, StatusFailedQuarkPromotionRejected, StatusFailedMixingPredictionRejected, StatusFailedFullFlavorCollapseRejected, StatusFlavorFirewallPreserved, a.Truth}}
	}}
}
