package contactquarticscalaryukawabundle

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ContactQuarticPrimaryScalarYukawaBundleFunctorAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-QUARTIC-PRIMARY-SCALAR-YUKAWA-BUNDLE-FUNCTOR-AUDIT"
	const name = "Contact quartic primary to scalar/Yukawa bundle functor audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 398 audit", Passed: false, Detail: err.Error()}}}
		}
		abstract := findFunctor(a.Functors.Candidates, "abstract quartic primary module")
		dimOnly := findFunctor(a.Functors.Candidates, "dimension-only quartic to H_phi identification")
		sealed := findFunctor(a.Functors.Candidates, "sealed companion operator on H_phi stress test")
		edge := findFunctor(a.Functors.Candidates, "quartic primary to one-form edge module")
		yukawa := findFunctor(a.Functors.Candidates, "quartic primary weighting of Yukawa fibers")
		native := findScenario(a.Impact.Scenarios, "native Gate398 ledger")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "inherits singleton obstruction, quartic module, scalar carrier, one-form/Yukawa target, and flavor firewall", Passed: a.Inheritance.Executed && a.Inheritance.Gate397SingletonFlavorBlocked && a.Inheritance.Gate183QuarticPrimaryDim == 4 && a.Inheritance.Gate183ScalarCarrierDim == 4 && a.Inheritance.Gate37ActiveScalarDim == 4 && a.Inheritance.Gate385OneFormEdgeSupportDerived && a.Inheritance.Gate26MinimalYukawaChannels == 8 && a.Inheritance.Gate372ChargedModuliDim == 13 && a.Inheritance.NoEmpiricalFlavorValuesImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "quartic primary block is exact, branch-free, and abstractly modular", Passed: a.Quartic.Dimension == 4 && a.Quartic.GaloisSafePrimary && a.Quartic.BranchFreeBlock && a.Quartic.IndividualBranchesSelected == 0 && a.Quartic.CompanionRepresentation && a.Quartic.AbstractRankOneModule && a.Quartic.ExactAsContactSpectralDatum && !a.Quartic.CanonicalHphiIdentification && !a.Quartic.ScalarMinimalPolynomialDerived, Detail: FormatQuartic(a.Quartic)},
			{Name: "scalar carrier is four-real-dimensional but has no quartic action", Passed: a.Scalar.ActiveRealDim == 4 && a.Scalar.ComplexDoubletDim == 2 && a.Scalar.ProtectedDirections == 3 && a.Scalar.NormalFormAvailable && !a.Scalar.CanonicalQuarticAction && !a.Scalar.HiggsMassDerived && !a.Scalar.ElectroweakScaleDerived, Detail: FormatScalar(a.Scalar)},
			{Name: "one-form/Yukawa target is derived but unweighted by the quartic primary", Passed: a.Target.OneFormEdgeSupportDerived && a.Target.JDoubledEdgeCount == 10 && a.Target.YukawaChannels == 8 && a.Target.ScalarFiberEntries == 16 && !a.Target.MassMatrixDerived && !a.Target.QuarticActsOnEdges && !a.Target.QuarticWeightsYukawaFibers && !a.Target.YukawaBundleReduced, Detail: FormatTarget(a.Target)},
			{Name: "abstract quartic module is not the physical scalar carrier", Passed: abstract.Native && abstract.DimensionCompatible && abstract.AlgebraHomomorphism && abstract.ProjectiveModule && !abstract.PhysicalCarrierAction && !abstract.PromotableAsNativeFunctor, Detail: FormatFunctor(abstract)},
			{Name: "dimension-only H_phi identification is rejected", Passed: dimOnly.Native && dimOnly.DimensionCompatible && !dimOnly.AlgebraHomomorphism && !dimOnly.PhysicalCarrierAction && !dimOnly.ScalarMinimalPolynomial && !dimOnly.PromotableAsNativeFunctor, Detail: FormatFunctor(dimOnly)},
			{Name: "sealed companion stress test is quarantined", Passed: sealed.Sealed && sealed.Circular && sealed.DimensionCompatible && sealed.AlgebraHomomorphism && sealed.ScalarMinimalPolynomial && sealed.ArbitraryBasisIdentification && !sealed.CompatibleWithJ && !sealed.CompatibleWithFirstOrder && !sealed.PromotableAsNativeFunctor, Detail: FormatFunctor(sealed)},
			{Name: "quartic primary does not act on the one-form edge module", Passed: edge.Native && !edge.DimensionCompatible && edge.Rank == 10 && !edge.CompatibleWithOneFormEdges && !edge.PromotableAsNativeFunctor, Detail: FormatFunctor(edge)},
			{Name: "quartic primary does not reduce Yukawa coupling space", Passed: yukawa.Native && !yukawa.DimensionCompatible && yukawa.Rank == 16 && !yukawa.ReducesYukawaCouplings && !yukawa.ReducesFlavorModuli && !yukawa.PromotableAsNativeFunctor, Detail: FormatFunctor(yukawa)},
			{Name: "no promotable quartic scalar/Yukawa functor exists", Passed: a.Functors.PromotableNativeCount == 0 && a.Functors.PhysicalScalarActions == 0 && a.Functors.OneFormEdgeActions == 0 && a.Functors.YukawaReducingActions == 0, Detail: FormatFunctors(a.Functors)},
			{Name: "native impact preserves scalar lane and flavor firewall", Passed: native.Native && native.Failed && !native.ScalarBundleDerived && !native.YukawaCouplingsReduced && native.ChargedModuliResult == 13 && native.FlavorFirewallPreserved && !a.Impact.NativeFlavorReduction && a.Impact.BestNativeModuliDim == 13 && a.Impact.ScalarHiggsLanePreserved, Detail: FormatScenario(native) + " :: " + FormatImpact(a.Impact)},
			{Name: "firewalls remain clean", Passed: a.Firewall.NoMassesImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoEmpiricalOrderingImported && a.Firewall.NoObservedHiggsUsedForFunctor && a.Firewall.NoManualQuarticHphiID && a.Firewall.NoCompanionOperatorPromoted && a.Firewall.NoArbitraryBasisMapPromoted && a.Firewall.NoYukawaCouplingClaimed && a.Firewall.NoFlavorModuliReductionClaimed, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate is scalar-bundle identity selector/obstruction", Passed: a.Next.Gate == 399 && a.Next.Title != "", Detail: FormatNext(a.Next)},
		}, Notes: []string{a.Truth}}
	}}
}

func findFunctor(xs []FunctorCandidate, name string) FunctorCandidate {
	for _, x := range xs {
		if x.Name == name {
			return x
		}
	}
	return FunctorCandidate{Name: name, Reason: "not found", Verdict: "MISSING"}
}

func findScenario(xs []ImpactScenario, name string) ImpactScenario {
	for _, x := range xs {
		if x.Name == name {
			return x
		}
	}
	return ImpactScenario{Name: name, Reason: "not found", Verdict: "MISSING"}
}
