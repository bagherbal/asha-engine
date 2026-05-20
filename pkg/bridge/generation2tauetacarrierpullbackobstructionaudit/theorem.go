package generation2tauetacarrierpullbackobstructionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2TauEtaCarrierPullbackObstructionAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 tau-eta carrier pullback obstruction audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate556 tau-eta pullback audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 555 selector theorem without promoting tau_eta", Passed: a.Inherited.SelectorTheoremProved && !a.Inherited.BMinusLUniqueWeakPlane && !a.Inherited.TauEtaPullbackValid && a.Inherited.TauEtaSealedCapacity, Detail: FormatInherited(a.Inherited)},
			{Name: "classify tau_eta as eta-graded trace vector rather than native operator", Passed: a.Type.IsTraceValueVector && a.Type.IsSealedBookkeepingDatum && !a.Type.IsSpectrumOfNativeOperator && !a.Type.IsDiagonalEndomorphism && !a.Type.IsCharacter && !a.Type.IsCoefficientVectorInNativeBasis, Detail: FormatType(a.Type)},
			{Name: "reject hand-inserted formal tau algebras as native source algebras", Passed: !a.SourceAlgebra.NativeSourceAlgebraExists && !a.SourceAlgebra.HasUnit && len(a.SourceAlgebra.Candidates) == 2 && a.SourceAlgebra.Candidates[0].InsertedByHand && !a.SourceAlgebra.Candidates[0].FoundAsNativeProjectAlgebra, Detail: FormatSource(a.SourceAlgebra)},
			{Name: "block unit-preserving carrier representation", Passed: !a.Representation.AnyValidUnitPreservingRepresentation && !a.Representation.RhoOneIsIdentity && len(a.Representation.Candidates) == 2, Detail: FormatRepresentation(a.Representation)},
			{Name: "record formal 2+1 selector capacity but reject canonical U12 selection", Passed: a.Selector.Gate555SelectorFormulaAppliesIfRepresentationExists && a.Selector.FormalCommutantDimension == 5 && !a.Selector.ValidRepresentationExists && !a.Selector.ProducesNativeSelector && !a.Selector.CanonicalU12Selected && a.Selector.BasisDependentIfForced, Detail: FormatSelector(a.Selector)},
			{Name: "separate formal B-L commutation from native refinement", Passed: a.BMinusL.FormalCommutatorWithBMinusLZero && !a.BMinusL.NativeCompatibilityVerified, Detail: FormatBMinusL(a.BMinusL)},
			{Name: "block spectral-triple promotion because compatibility data are missing", Passed: !a.SpectralTriple.NativeSpectralTriplePromotionAllowed && !a.SpectralTriple.GammaCompatibilityAvailable && !a.SpectralTriple.JCompatibilityAvailable && !a.SpectralTriple.DCompatibilityAvailable && !a.SpectralTriple.FirstOrderCompatibilityAvailable && len(a.SpectralTriple.MissingData) >= 6, Detail: FormatSpectralTriple(a.SpectralTriple)},
			{Name: "preserve weak/flavor/Higgs/Yukawa/CKM firewalls", Passed: !a.Firewall.PromotedToWeakIsospin && !a.Firewall.PromotedToGenerationMassHierarchy && !a.Firewall.PromotedToHiggs && !a.Firewall.PromotedToYukawa && !a.Firewall.PromotedToCKMPMNS && !a.Firewall.InsertedFormalAlgebraAsNative && !a.Firewall.InsertedDiagonalMatrixAsNative && !a.Firewall.NativeRegistryPolluted, Detail: FormatFirewall(a.Firewall)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(Statuses(), a.Truth, a.Final.MissingNextTheorem)}
	}}
}
