package contactedgepullback

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ContactToEdgeNaturalTransformationPullbackSieveTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Contact-to-edge natural transformation / pullback sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate405 audit", Passed: false, Detail: err.Error()}}}
		}
		native := findCandidate(a.Sieve.Candidates, "native contact projector to one-form edge ledger")
		yukawa := findCandidate(a.Sieve.Candidates, "reverse of canonical Yukawa edge restriction")
		ext := findCandidate(a.Sieve.Candidates, "sealed q4 extension to five structural edge slots")
		jdup := findCandidate(a.Sieve.Candidates, "sealed J-doubled q4 pullback")
		intertwiner := findCandidate(a.Sieve.Candidates, "contact q4 as edge weight/intertwiner with native D_F edge graph")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "inherits quotient obstruction chain", Passed: a.Inheritance.Executed && a.Inheritance.Gate398NoQuarticBundleFunctor && a.Inheritance.Gate399QuaternionicPolynomialNo && a.Inheritance.Gate400NoMixedEdgeQ4 && a.Inheritance.Gate401ChargeWeightsDisjoint && a.Inheritance.Gate402GraphNoQ4 && a.Inheritance.Gate403OrientationNoQ4 && a.Inheritance.Gate404QuotientNoQ4 && a.Inheritance.Gate404NeedsPullback && a.Inheritance.Gate385OneFormEdges && a.Inheritance.Gate385JDoubledEdgeCount == JDoubledEdgeCount && a.Inheritance.Gate372ChargedModuliDim == Gate372ChargedModuliDim && a.Inheritance.NoEmpiricalInputsImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "q4 target is exact contact primary", Passed: a.Q4.Degree == Q4Degree && a.Q4.Dimension == ContactPrimaryDim && a.Q4.IrreducibleOverQ && len(a.Q4.MonicCoefficients) == 5, Detail: FormatQ4(a.Q4)},
			{Name: "pullback arena is typed but no native functor is known", Passed: a.Arena.Formalized && a.Arena.ContactPrimaryDim == ContactPrimaryDim && a.Arena.StructuralEdgeDim == StructuralEdgeCount && a.Arena.JDoubledEdgeDim == JDoubledEdgeCount && a.Arena.HphiDim == HphiRealDim && !a.Arena.NativeFunctorKnown && !a.Arena.ContactEdgeActionDerived && !a.Arena.UsesObservedMasses && !a.Arena.UsesYukawaAmplitudes && !a.Arena.UsesManualRootPlacement, Detail: FormatArena(a.Arena)},
			{Name: "native contact projector has no edge action", Passed: !native.Native && native.ContactDerived && native.EdgeDerived && !native.Typed && !native.PullbackConstructed && !native.PreservesQ4Polynomial && native.Verdict == StatusFailedNoNativeContactToEdgeMap, Detail: FormatCandidate(native)},
			{Name: "Yukawa restriction is wrong-direction and circular", Passed: yukawa.Native && yukawa.Typed && yukawa.EdgeDerived && yukawa.Circular && !yukawa.ContactDerived && !yukawa.PullbackConstructed && !yukawa.PreservesQ4Polynomial && yukawa.Verdict == StatusFailedYukawaRestrictionWrongDirection, Detail: FormatCandidate(yukawa)},
			{Name: "sealed q4 E5 extension is manual", Passed: ext.Sealed && ext.Circular && ext.PullbackConstructed && ext.PreservesQ4Polynomial && !ext.Native && !ext.Canonical && !ext.JCompatible && !ext.FirstOrderCompatible && !ext.DFIntertwiner && !ext.PromotableAsQ4EdgeWeight, Detail: FormatCandidate(ext)},
			{Name: "sealed J-doubled q4 duplicates manual placement", Passed: jdup.Sealed && jdup.Circular && jdup.PullbackConstructed && jdup.PreservesQ4Polynomial && jdup.JCompatible && !jdup.Native && !jdup.Canonical && !jdup.FirstOrderCompatible && !jdup.DFIntertwiner && !jdup.PromotableAsQ4EdgeWeight, Detail: FormatCandidate(jdup)},
			{Name: "q4 block does not intertwine native edge graph", Passed: intertwiner.ContactDerived && intertwiner.EdgeDerived && !intertwiner.DFIntertwiner && !intertwiner.NaturalitySquareFormed && !intertwiner.CommutatorZero && !intertwiner.PromotableAsQ4EdgeWeight, Detail: FormatCandidate(intertwiner)},
			{Name: "no canonical natural transformation exists", Passed: a.Sieve.NativePullbackCount == 0 && a.Sieve.NativeQ4PreservingCount == 0 && a.Sieve.NativeDFIntertwinerCount == 0 && a.Sieve.CanonicalNaturalTransformCount == 0 && a.Sieve.SealedOrManualCount > 0, Detail: FormatSieve(a.Sieve)},
			{Name: "identity and flavor firewalls are preserved", Passed: !a.Impact.ContactPullbackAchieved && !a.Impact.Q4OnEdgeSpacePreserved && !a.Impact.CanonicalNaturalTransformation && !a.Impact.HphiQuarticIdentified && !a.Impact.YukawaCouplingsReduced && a.Impact.ChargedModuliResult == Gate372ChargedModuliDim && a.Impact.FlavorFirewallPreserved && a.Impact.HiggsLanePreserved, Detail: FormatImpact(a.Impact)},
			{Name: "empirical and manual-placement firewalls remain clean", Passed: a.Firewall.NoObservedMassesImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoYukawaAmplitudesInserted && a.Firewall.NoManualQ4HphiID && a.Firewall.NoManualRootPlacementPromoted && a.Firewall.NoArbitraryEdgeBasisPromoted && a.Firewall.NoCompanionOperatorPromoted && a.Firewall.NoFlavorModuliReductionClaimed, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate classifies q4 as contact-internal or searches a different scalar selector", Passed: a.Next.Gate == 406 && a.Next.Title != "", Detail: FormatNext(a.Next)},
		}, Notes: []string{a.Truth}}
	}}
}

func findCandidate(xs []PullbackCandidate, name string) PullbackCandidate {
	for _, x := range xs {
		if x.Name == name {
			return x
		}
	}
	return PullbackCandidate{Name: name, Verdict: "MISSING"}
}
