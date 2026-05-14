package edgetohphiquotient

import "github.com/bagherbal/asha-engine/pkg/theorem"

func CanonicalEdgeToHphiQuotientContactEdgeIntertwinerSieveTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Canonical edge-to-Hphi quotient / contact-edge intertwiner sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate404 audit", Passed: false, Detail: err.Error()}}}
		}
		qy := findCandidate(a.Sieve.Candidates, "canonical Higgs/Yukawa edge restriction Q_Y: E_5 -> E_Y ~= H_phi")
		branch := findCandidate(a.Sieve.Candidates, "scalar branch quotient Q_branch: E_Y -> Phi_+ ⊕ Phi_-")
		j := findCandidate(a.Sieve.Candidates, "J-even/J-odd quotient from ten J-doubled edge slots")
		contact := findCandidate(a.Sieve.Candidates, "contact/scalar response quotient Q_contact from active contact sector")
		full := findCandidate(a.Sieve.Candidates, "full five-edge spectral quotient by chosen edge mode")
		sealed := findCandidate(a.Sieve.Candidates, "sealed q4 edge-to-Hphi companion quotient")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "inherits q4-scalar obstruction chain", Passed: a.Inheritance.Executed && a.Inheritance.Gate398NoQuarticBundleFunctor && a.Inheritance.Gate399QuaternionicPolynomialNo && a.Inheritance.Gate400NoMixedEdgeQ4 && a.Inheritance.Gate401ChargeWeightsDisjoint && a.Inheritance.Gate402GraphNoQ4 && a.Inheritance.Gate403OrientationNoQ4 && a.Inheritance.Gate403NeedsQuotient && a.Inheritance.Gate385OneFormEdges && a.Inheritance.Gate385JDoubledEdgeCount == JDoubledEdgeCount && a.Inheritance.Gate372ChargedModuliDim == Gate372ChargedModuliDim && a.Inheritance.NoEmpiricalInputsImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "q4 target remains irreducible quartic", Passed: a.Q4.Degree == Q4Degree && a.Q4.IrreducibleOverQ && len(a.Q4.MonicCoefficients) == 5, Detail: FormatQ4(a.Q4)},
			{Name: "quotient arena is formalized without empirical inputs", Passed: a.Arena.Formalized && a.Arena.StructuralEdgeDim == StructuralEdgeCount && a.Arena.JDoubledEdgeDim == JDoubledEdgeCount && a.Arena.HphiDim == HphiRealDim && !a.Arena.HasCanonicalFullEdgeQuotient && a.Arena.HasCanonicalYukawaRestriction && a.Arena.HasCanonicalBranchMap && a.Arena.HasCanonicalJEvenMap && !a.Arena.UsesObservedMasses && !a.Arena.UsesYukawaAmplitudes && !a.Arena.UsesManualQ4Placement, Detail: FormatArena(a.Arena)},
			{Name: "canonical Yukawa quotient exists but is pair-degenerate", Passed: qy.Native && qy.CanonicalQuotient && qy.OneFormDerived && qy.HphiEndomorphism && qy.Rank == HphiRealDim && qy.KernelDimension == 1 && qy.PairDegenerate && qy.MinimalDegree == 2 && !qy.Q4ExactMatch && !qy.PromotableAsQ4Selector, Detail: FormatCandidate(qy)},
			{Name: "scalar branch quotient is rank two", Passed: branch.Native && branch.CanonicalQuotient && branch.Rank == 2 && branch.KernelDimension == 2 && branch.MinimalDegree == 2 && branch.PairDegenerate && !branch.Q4ExactMatch && !branch.PromotableAsQ4Selector, Detail: FormatCandidate(branch)},
			{Name: "J quotient duplicates pair spectrum", Passed: j.Native && j.CanonicalQuotient && j.JCompatible && j.SourceDim == JDoubledEdgeCount && j.TargetDim == HphiRealDim && j.PairDegenerate && j.MinimalDegree == 2 && !j.Q4ExactMatch && !j.PromotableAsQ4Selector, Detail: FormatCandidate(j)},
			{Name: "contact/scalar response quotient remains quadratic", Passed: contact.Native && contact.ContactDerived && contact.HphiEndomorphism && contact.Rank == HphiRealDim && contact.MinimalDegree == 2 && contact.PairDegenerate && !contact.Q4ExactMatch && !contact.PromotableAsQ4Selector, Detail: FormatCandidate(contact)},
			{Name: "full five-edge quotient is noncanonical", Passed: full.Sealed && full.Circular && !full.Native && !full.CanonicalQuotient && full.MinimalDegree == 4 && full.IrreducibleQuarticCapacity && !full.PromotableAsQ4Selector, Detail: FormatCandidate(full)},
			{Name: "sealed q4 companion quotient is quarantined", Passed: sealed.Sealed && sealed.Circular && sealed.Q4ExactMatch && !sealed.Native && !sealed.CanonicalQuotient && !sealed.PromotableAsQ4Selector, Detail: FormatCandidate(sealed)},
			{Name: "no native q4 edge-to-Hphi intertwiner exists", Passed: a.Sieve.NativeQuotientCount > 0 && a.Sieve.NativeHphiEndomorphismCount > 0 && a.Sieve.NativeQuarticCapacityCount == 0 && a.Sieve.CanonicalHphiQ4MatchCount == 0, Detail: FormatSieve(a.Sieve)},
			{Name: "identity and flavor firewalls are preserved", Passed: !a.Impact.HphiQuarticIdentified && a.Impact.CanonicalQuotientFound && a.Impact.CanonicalYukawaQuotientFound && !a.Impact.NativeIntertwinerQ4Found && !a.Impact.YukawaCouplingsReduced && a.Impact.ChargedModuliResult == Gate372ChargedModuliDim && a.Impact.FlavorFirewallPreserved && a.Impact.HiggsLanePreserved, Detail: FormatImpact(a.Impact)},
			{Name: "empirical and arbitrary quotient firewalls remain clean", Passed: a.Firewall.NoObservedMassesImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoYukawaAmplitudesInserted && a.Firewall.NoManualQ4HphiID && a.Firewall.NoArbitraryFullEdgeQuotientPromoted && a.Firewall.NoCompanionOperatorPromoted && a.Firewall.NoFlavorModuliReductionClaimed, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate searches contact-to-edge pullback", Passed: a.Next.Gate == 405 && a.Next.Title != "", Detail: FormatNext(a.Next)},
		}, Notes: []string{a.Truth}}
	}}
}

func findCandidate(xs []QuotientCandidate, name string) QuotientCandidate {
	for _, x := range xs {
		if x.Name == name {
			return x
		}
	}
	return QuotientCandidate{Name: name, Verdict: "MISSING"}
}
