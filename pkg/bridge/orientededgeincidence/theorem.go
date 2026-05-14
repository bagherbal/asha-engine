package orientededgeincidence

import "github.com/bagherbal/asha-engine/pkg/theorem"

func OrientedEdgeIncidenceBoundaryOperatorSieveTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Oriented edge-incidence boundary operator sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate403 audit", Passed: false, Detail: err.Error()}}}
		}
		yy := findCandidate(a.Sieve.Candidates, "four Yukawa oriented edge Gram d_Y^T d_Y")
		full := findCandidate(a.Sieve.Candidates, "full five-edge oriented incidence Gram d_E^T d_E")
		quotient := findCandidate(a.Sieve.Candidates, "noncanonical four-mode quotient of full oriented incidence Gram")
		twist := findCandidate(a.Sieve.Candidates, "J-twisted complex Majorana boundary d^†d")
		doubled := findCandidate(a.Sieve.Candidates, "J-doubled oriented boundary Gram")
		sealed := findCandidate(a.Sieve.Candidates, "sealed q4 oriented-boundary companion quotient")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "inherits edge-graph and q4 obstruction history", Passed: a.Inheritance.Executed && a.Inheritance.Gate399QuaternionicPolynomialNo && a.Inheritance.Gate400NoMixedEdgeQ4 && a.Inheritance.Gate401ChargeWeightsDisjoint && a.Inheritance.Gate402UndirectedGraphNative && a.Inheritance.Gate402FullGraphQuarticCapacity && a.Inheritance.Gate402NoGraphQ4 && a.Inheritance.Gate385OneFormEdges && a.Inheritance.Gate385JDoubledEdgeCount == JDoubledEdgeCount && a.Inheritance.Gate297FirstOrderEdgeGraph && a.Inheritance.Gate372ChargedModuliDim == Gate372ChargedModuliDim && a.Inheritance.NoEmpiricalInputsImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "q4 target remains irreducible quartic", Passed: a.Q4.Degree == Q4Degree && a.Q4.IrreducibleOverQ && len(a.Q4.MonicCoefficients) == 5, Detail: FormatQ4(a.Q4)},
			{Name: "oriented boundary arena is formalized", Passed: a.Arena.Formalized && a.Arena.StructuralEdgeCount == StructuralEdgeCount && a.Arena.YukawaEdgeCount == YukawaEdgeCount && a.Arena.JDoubledEdgeCount == JDoubledEdgeCount && a.Arena.ChiralOrientationAvailable && !a.Arena.MajoranaOrientationCanonical && !a.Arena.HasCanonicalHphiQuotient && !a.Arena.UsesGaugeChargeWeights && !a.Arena.UsesYukawaAmplitudes && !a.Arena.UsesObservedMasses, Detail: FormatArena(a.Arena)},
			{Name: "four Yukawa oriented incidence Gram remains pair-degenerate", Passed: yy.Native && yy.HphiEndomorphism && yy.CanonicalQuotientToHphi && yy.OrientationSignsCancel && yy.PairDegenerate && yy.MinimalDegree == 2 && !yy.Q4ExactMatch && !yy.PromotableAsQ4Selector, Detail: FormatCandidate(yy)},
			{Name: "full five-edge oriented incidence Gram is not Hphi and not q4", Passed: full.Native && full.BoundaryDerived && !full.HphiEndomorphism && !full.CanonicalQuotientToHphi && full.Dimension == StructuralEdgeCount && full.MinimalDegree == 5 && full.OrientationSignsCancel && !full.Q4ExactMatch && !full.PromotableAsQ4Selector, Detail: FormatCandidate(full)},
			{Name: "forced four-mode quotient is sealed and disjoint from q4", Passed: quotient.Sealed && quotient.Circular && !quotient.Native && quotient.Dimension == HphiRealDim && quotient.MinimalDegree == 4 && quotient.MinimalResidualToQ4 > 1 && !quotient.Q4ExactMatch && !quotient.PromotableAsQ4Selector, Detail: FormatCandidate(quotient)},
			{Name: "Majorana phase/twist is not spectrally new", Passed: twist.Sealed && twist.CompatibleWithJ && twist.OrientationSignsCancel && !twist.HphiEndomorphism && twist.MinimalDegree == 5 && !twist.Q4ExactMatch && !twist.PromotableAsQ4Selector, Detail: FormatCandidate(twist)},
			{Name: "J-doubled oriented boundary only duplicates spectrum", Passed: doubled.Native && doubled.CompatibleWithJ && doubled.Dimension == JDoubledEdgeCount && doubled.MinimalDegree == 5 && !doubled.HphiEndomorphism && !doubled.Q4ExactMatch && !doubled.PromotableAsQ4Selector, Detail: FormatCandidate(doubled)},
			{Name: "sealed q4 companion is quarantined", Passed: sealed.Sealed && sealed.Circular && sealed.Q4ExactMatch && !sealed.Native && !sealed.CanonicalQuotientToHphi && !sealed.PromotableAsQ4Selector, Detail: FormatCandidate(sealed)},
			{Name: "no native oriented q4 selector exists", Passed: a.Sieve.NativeBoundaryOperatorCount > 0 && a.Sieve.NativeHphiEndomorphismCount > 0 && a.Sieve.NativeQuarticCapacityCount > 0 && a.Sieve.CanonicalHphiQ4MatchCount == 0, Detail: FormatSieve(a.Sieve)},
			{Name: "identity and flavor firewalls are preserved", Passed: !a.Impact.HphiQuarticIdentified && a.Impact.NativeBoundaryOperatorFound && !a.Impact.CanonicalBoundaryQuotientFound && a.Impact.OrientedIncidenceLaneOpened && !a.Impact.YukawaCouplingsReduced && a.Impact.ChargedModuliResult == Gate372ChargedModuliDim && a.Impact.FlavorFirewallPreserved && a.Impact.HiggsLanePreserved, Detail: FormatImpact(a.Impact)},
			{Name: "empirical and arbitrary-quotient firewalls remain clean", Passed: a.Firewall.NoObservedMassesImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoYukawaAmplitudesInserted && a.Firewall.NoGaugeChargeFitReused && a.Firewall.NoManualQ4HphiID && a.Firewall.NoArbitraryBoundaryQuotient && a.Firewall.NoCompanionOperatorPromoted && a.Firewall.NoFlavorModuliReductionClaimed, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate searches canonical edge-to-Hphi quotient", Passed: a.Next.Gate == 404 && a.Next.Title != "", Detail: FormatNext(a.Next)},
		}, Notes: []string{a.Truth}}
	}}
}

func findCandidate(xs []BoundaryCandidate, name string) BoundaryCandidate {
	for _, x := range xs {
		if x.Name == name {
			return x
		}
	}
	return BoundaryCandidate{Name: name, Verdict: "MISSING"}
}
