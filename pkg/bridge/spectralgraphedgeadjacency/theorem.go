package spectralgraphedgeadjacency

import "github.com/bagherbal/asha-engine/pkg/theorem"

func SpectralGraphEdgeAdjacencyOperatorSearchTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Spectral graph edge-adjacency operator search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate402 audit", Passed: false, Detail: err.Error()}}}
		}
		yukawaAdj := findCandidate(a.Sieve.Candidates, "four Yukawa-edge adjacency graph K2 disjoint union K2")
		yukawaLap := findCandidate(a.Sieve.Candidates, "four Yukawa-edge graph Laplacian K2 disjoint union K2")
		fullLap := findCandidate(a.Sieve.Candidates, "full five-edge structural Laplacian P3 disjoint union K2")
		pos := findCandidate(a.Sieve.Candidates, "positive-spectrum quotient of full five-edge Laplacian")
		doubled := findCandidate(a.Sieve.Candidates, "J-doubled structural edge graph")
		sealed := findCandidate(a.Sieve.Candidates, "sealed q4 edge-graph companion quotient")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "inherits Gate401 obstruction and finite edge graph ledgers", Passed: a.Inheritance.Executed && a.Inheritance.Gate400NoNativeQ4Selector && a.Inheritance.Gate401AnisotropicWeightsFound && a.Inheritance.Gate401NoNativeWeightedLaplacian && a.Inheritance.Gate385OneFormEdges && a.Inheritance.Gate385JDoubledEdgeCount == JDoubledEdgeCount && a.Inheritance.Gate297FirstOrderEdgeGraph && a.Inheritance.Gate298InnerFluctuationFields && a.Inheritance.Gate372ChargedModuliDim == Gate372ChargedModuliDim && a.Inheritance.NoEmpiricalInputsImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "q4 target remains irreducible quartic", Passed: a.Q4.Degree == Q4Degree && a.Q4.IrreducibleOverQ && len(a.Q4.MonicCoefficients) == 5, Detail: FormatQ4(a.Q4)},
			{Name: "edge graph arena is formalized without charges, amplitudes, or masses", Passed: a.Arena.Formalized && a.Arena.StructuralEdgeCount == StructuralEdgeCount && a.Arena.YukawaEdgeCount == YukawaEdgeCount && a.Arena.JDoubledEdgeCount == JDoubledEdgeCount && a.Arena.HasCanonicalEndpointIncidence && !a.Arena.HasCanonicalEdgeOrientation && !a.Arena.HasCanonicalHphiQuotient && !a.Arena.UsesGaugeChargeWeights && !a.Arena.UsesYukawaAmplitudes && !a.Arena.UsesObservedMasses, Detail: FormatArena(a.Arena)},
			{Name: "four Yukawa-edge adjacency is pair-degenerate", Passed: yukawaAdj.Native && yukawaAdj.HphiEndomorphism && yukawaAdj.CanonicalQuotientToHphi && yukawaAdj.PairDegenerate && yukawaAdj.MinimalDegree == 2 && !yukawaAdj.Q4ExactMatch && !yukawaAdj.PromotableAsQ4Selector, Detail: FormatCandidate(yukawaAdj)},
			{Name: "four Yukawa-edge Laplacian is pair-degenerate", Passed: yukawaLap.Native && yukawaLap.HphiEndomorphism && yukawaLap.CanonicalQuotientToHphi && yukawaLap.PairDegenerate && yukawaLap.MinimalDegree == 2 && !yukawaLap.Q4ExactMatch && !yukawaLap.PromotableAsQ4Selector, Detail: FormatCandidate(yukawaLap)},
			{Name: "full five-edge graph has quartic capacity but is not Hphi and misses q4", Passed: fullLap.Native && fullLap.EdgeGraphDerived && !fullLap.HphiEndomorphism && !fullLap.CanonicalQuotientToHphi && fullLap.Dimension == 5 && fullLap.MinimalDegree == 4 && fullLap.IrreducibleQuarticCapacity && fullLap.MinimalResidualToQ4 > 1 && !fullLap.Q4ExactMatch && !fullLap.PromotableAsQ4Selector, Detail: FormatCandidate(fullLap)},
			{Name: "positive graph quotient is three-dimensional, not Hphi", Passed: pos.Native && !pos.HphiEndomorphism && pos.Dimension == 3 && pos.MinimalDegree == 3 && !pos.Q4ExactMatch && !pos.PromotableAsQ4Selector, Detail: FormatCandidate(pos)},
			{Name: "J-doubled graph only duplicates structural spectrum", Passed: doubled.Native && doubled.CompatibleWithJ && doubled.Dimension == JDoubledEdgeCount && doubled.MinimalDegree == 4 && !doubled.HphiEndomorphism && !doubled.Q4ExactMatch && !doubled.PromotableAsQ4Selector, Detail: FormatCandidate(doubled)},
			{Name: "sealed q4 companion is quarantined", Passed: sealed.Sealed && sealed.Circular && sealed.Q4ExactMatch && !sealed.Native && !sealed.CanonicalQuotientToHphi && !sealed.PromotableAsQ4Selector, Detail: FormatCandidate(sealed)},
			{Name: "no native q4 edge adjacency operator exists", Passed: a.Sieve.NativeGraphOperatorCount > 0 && a.Sieve.NativeHphiEndomorphismCount > 0 && a.Sieve.NativeQuarticCapacityCount > 0 && a.Sieve.CanonicalHphiQ4MatchCount == 0, Detail: FormatSieve(a.Sieve)},
			{Name: "identity and flavor firewalls are preserved", Passed: !a.Impact.HphiQuarticIdentified && a.Impact.NativeEdgeAdjacencyFound && !a.Impact.CanonicalGraphQuotientFound && !a.Impact.YukawaCouplingsReduced && a.Impact.ChargedModuliResult == Gate372ChargedModuliDim && a.Impact.FlavorFirewallPreserved && a.Impact.HiggsLanePreserved, Detail: FormatImpact(a.Impact)},
			{Name: "empirical and arbitrary-quotient firewalls remain clean", Passed: a.Firewall.NoObservedMassesImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoYukawaAmplitudesInserted && a.Firewall.NoGaugeChargeFitReused && a.Firewall.NoManualQ4HphiID && a.Firewall.NoArbitraryGraphQuotient && a.Firewall.NoCompanionOperatorPromoted && a.Firewall.NoFlavorModuliReductionClaimed, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate searches oriented incidence/boundary operator", Passed: a.Next.Gate == 403 && a.Next.Title != "", Detail: FormatNext(a.Next)},
		}, Notes: []string{a.Truth}}
	}}
}

func findCandidate(xs []GraphCandidate, name string) GraphCandidate {
	for _, x := range xs {
		if x.Name == name {
			return x
		}
	}
	return GraphCandidate{Name: name, Verdict: "MISSING"}
}
