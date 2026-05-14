package mixededgelaplaciansieve

import "github.com/bagherbal/asha-engine/pkg/theorem"

func NonQuaternionicScalarIdentityMixedEdgeLaplacianSieveTheorem() theorem.Theorem {
	const id = "GATE400-NON-QUATERNIONIC-SCALAR-IDENTITY-MIXED-EDGE-LAPLACIAN-SIEVE"
	const name = "Non-quaternionic scalar identity / mixed edge Laplacian sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate400 audit", Passed: false, Detail: err.Error()}}}
		}
		uniform := findCandidate(a.Mixed.Candidates, "uniform one-form edge Laplacian projected to H_phi")
		raw := findCandidate(a.Mixed.Candidates, "raw contact-to-scalar compression P_C Delta_E P_K")
		squared := findCandidate(a.Mixed.Candidates, "squared contact/edge compression scalar response")
		sealed := findCandidate(a.Mixed.Candidates, "sealed q4 companion operator declared on H_phi")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "inherits Gate 398/399 scalar identity obstructions and Gate 385 one-form support", Passed: a.Inheritance.Executed && a.Inheritance.Gate398NoCanonicalHphiID && a.Inheritance.Gate399QuaternionicDisjoint && a.Inheritance.Gate385OneFormEdgeSupportDerived && a.Inheritance.Gate385JDoubledEdgeCount == JDoubledEdgeCount && a.Inheritance.Gate37HphiRealDim == HphiRealDim && a.Inheritance.Gate372ChargedModuliDim == Gate372ChargedModuliDim && a.Inheritance.NoEmpiricalInputsImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "q4 remains irreducible quartic contact primary target", Passed: a.Q4.Degree == 4 && a.Q4.IrreducibleOverQ && a.Q4.ContactPrimary && a.Q4.BranchFree, Detail: FormatQ4(a.Q4)},
			{Name: "one-form edge Laplacian arena is formalized without empirical masses", Passed: a.Arena.Formalized && a.Arena.EdgeSupportDimension == JDoubledEdgeCount && a.Arena.HphiDimension == HphiRealDim && a.Arena.OneFormEdgeMeasureDerived && a.Arena.UniformEdgeMetric && !a.Arena.ExplicitDFEdgeWeightsDerived && !a.Arena.PhysicalMassesInserted, Detail: FormatArena(a.Arena)},
			{Name: "uniform edge Laplacian is central on H_phi", Passed: uniform.Native && uniform.HphiEndomorphism && uniform.CentralOnHphi && uniform.MinimalDegree == 1 && !uniform.Q4ExactMatch && !uniform.PromotableAsQ4Selector, Detail: FormatCandidate(uniform)},
			{Name: "raw contact compression is not an H_phi endomorphism", Passed: raw.Native && raw.ContactCompressed && !raw.HphiEndomorphism && raw.MinimalDegree == 0 && !raw.Q4ExactMatch && !raw.PromotableAsQ4Selector, Detail: FormatCandidate(raw)},
			{Name: "squared mixed compression recovers pair-degenerate scalar response, not q4", Passed: squared.Native && squared.HphiEndomorphism && squared.PairDegenerate && squared.MinimalDegree == 2 && !squared.IrreducibleQuartic && !squared.Q4ExactMatch && !squared.PromotableAsQ4Selector, Detail: FormatCandidate(squared)},
			{Name: "sealed q4 companion stress test is quarantined", Passed: sealed.Sealed && sealed.Circular && sealed.Q4ExactMatch && sealed.MinimalDegree == 4 && !sealed.Native && !sealed.CompatibleWithJ && !sealed.CompatibleWithFirstOrder && !sealed.PromotableAsQ4Selector, Detail: FormatCandidate(sealed)},
			{Name: "no native q4 scalar selector exists", Passed: a.Mixed.NativeQ4MatchCount == 0 && a.Mixed.PromotableNativeCount == 0 && a.Mixed.SealedQ4MatchCount == 1, Detail: FormatMixed(a.Mixed)},
			{Name: "identity and flavor firewalls are preserved", Passed: !a.Impact.HphiQuarticIdentified && !a.Impact.OneFormEdgeFunctorDerived && !a.Impact.YukawaCouplingsReduced && a.Impact.ChargedModuliResult == Gate372ChargedModuliDim && a.Impact.FlavorFirewallPreserved && a.Impact.HiggsLanePreserved, Detail: FormatImpact(a.Impact)},
			{Name: "empirical and arbitrary-identification firewalls remain clean", Passed: a.Firewall.NoObservedMassesImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoObservedHiggsInserted && a.Firewall.NoManualQ4HphiID && a.Firewall.NoCompanionOperatorPromoted && a.Firewall.NoArbitraryBasisMapPromoted && a.Firewall.NoYukawaCouplingClaimed && a.Firewall.NoFlavorModuliReductionClaimed, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate searches derived nonuniform edge weights", Passed: a.Next.Gate == 401 && a.Next.Title != "", Detail: FormatNext(a.Next)},
		}, Notes: []string{a.Truth}}
	}}
}

func findCandidate(xs []MixedOperatorCandidate, name string) MixedOperatorCandidate {
	for _, x := range xs {
		if x.Name == name {
			return x
		}
	}
	return MixedOperatorCandidate{Name: name, Verdict: "MISSING"}
}
