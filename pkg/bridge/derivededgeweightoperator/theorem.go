package derivededgeweightoperator

import "github.com/bagherbal/asha-engine/pkg/theorem"

func DerivedEdgeWeightOperatorHyperchargeLaplacianSieveTheorem() theorem.Theorem {
	const id = "GATE401-DERIVED-EDGE-WEIGHT-OPERATOR-HYPERCHARGE-LAPLACIAN-SIEVE"
	const name = "Derived edge-weight operator / hypercharge Laplacian sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate401 audit", Passed: false, Detail: err.Error()}}}
		}
		uniform := findCandidate(a.Sieve.Candidates, "uniform J-doubled edge measure")
		t3 := findCandidate(a.Sieve.Candidates, "scalar branch T3/hypercharge weight")
		branchY := findCandidate(a.Sieve.Candidates, "branch-averaged right-hypercharge edge Laplacian")
		edgeY := findCandidate(a.Sieve.Candidates, "edge-resolved right-hypercharge four-channel stress test")
		edgeY2 := findCandidate(a.Sieve.Candidates, "edge-resolved squared-hypercharge stress test")
		bl := findCandidate(a.Sieve.Candidates, "edge-resolved B-L four-channel stress test")
		sealed := findCandidate(a.Sieve.Candidates, "sealed q4-weighted edge companion")
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{
			{Name: "inherits Gate400 obstruction, Gate385 one-form edges, charge ledger, and flavor firewall", Passed: a.Inheritance.Executed && a.Inheritance.Gate400UniformCentral && a.Inheritance.Gate400PairDegenerateCompression && a.Inheritance.Gate400NoNativeQ4Selector && a.Inheritance.Gate385OneFormEdges && a.Inheritance.Gate385JDoubledEdgeCount == JDoubledEdgeCount && a.Inheritance.Gate26YukawaChannelsDerived && a.Inheritance.Gate41HyperchargeNormalization && a.Inheritance.Gate372ChargedModuliDim == Gate372ChargedModuliDim && a.Inheritance.NoEmpiricalInputsImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "q4 target remains irreducible quartic", Passed: a.Q4.Degree == Q4Degree && a.Q4.IrreducibleOverQ && len(a.Q4.MonicCoefficients) == 5, Detail: FormatQ4(a.Q4)},
			{Name: "edge-weight arena has native charge weights and no empirical amplitudes", Passed: a.Arena.Formalized && a.Arena.StructuralEdgeCount == StructuralEdgeCount && a.Arena.JDoubledEdgeCount == JDoubledEdgeCount && a.Arena.HphiDimension == HphiRealDim && a.Arena.NativeElectroweakWeights && a.Arena.NativeBMinusLWeights && a.Arena.NativeT3Weights && !a.Arena.ExplicitYukawaAmplitudesUsed && !a.Arena.ObservedMassesUsed, Detail: FormatArena(a.Arena)},
			{Name: "uniform edge weight remains central", Passed: uniform.NativeWeights && uniform.CanonicalCompressionToHphi && uniform.CentralOnHphi && uniform.MinimalDegree == 1 && !uniform.Q4ExactMatch && !uniform.PromotableAsQ4Selector, Detail: FormatCandidate(uniform)},
			{Name: "scalar T3/hypercharge branch weight is pair-degenerate, not q4", Passed: t3.NativeWeights && t3.CanonicalCompressionToHphi && t3.PairDegenerate && t3.MinimalDegree == 2 && !t3.Q4ExactMatch && !t3.PromotableAsQ4Selector, Detail: FormatCandidate(t3)},
			{Name: "branch-averaged hypercharge compression remains pair-degenerate", Passed: branchY.NativeWeights && branchY.CanonicalCompressionToHphi && branchY.PairDegenerate && branchY.MinimalDegree == 2 && !branchY.Q4ExactMatch && !branchY.PromotableAsQ4Selector, Detail: FormatCandidate(branchY)},
			{Name: "edge-resolved hypercharge has degree-four capacity but is noncanonical and polynomially disjoint from q4", Passed: edgeY.NativeWeights && edgeY.EdgeResolved && !edgeY.CanonicalCompressionToHphi && edgeY.MinimalDegree == 4 && edgeY.CharacteristicResidualToQ4 > 0.1 && !edgeY.Q4ExactMatch && !edgeY.PromotableAsQ4Selector, Detail: FormatCandidate(edgeY)},
			{Name: "edge-resolved squared hypercharge also misses q4", Passed: edgeY2.NativeWeights && edgeY2.EdgeResolved && !edgeY2.CanonicalCompressionToHphi && edgeY2.MinimalDegree == 4 && edgeY2.CharacteristicResidualToQ4 > 0.1 && !edgeY2.Q4ExactMatch && !edgeY2.PromotableAsQ4Selector, Detail: FormatCandidate(edgeY2)},
			{Name: "B-L edge weights do not produce q4", Passed: bl.NativeWeights && bl.PairDegenerate && bl.MinimalDegree == 2 && !bl.Q4ExactMatch && !bl.PromotableAsQ4Selector, Detail: FormatCandidate(bl)},
			{Name: "sealed q4-weighted companion is quarantined", Passed: sealed.Sealed && sealed.Circular && sealed.Q4ExactMatch && !sealed.NativeWeights && !sealed.CanonicalCompressionToHphi && !sealed.PromotableAsQ4Selector, Detail: FormatCandidate(sealed)},
			{Name: "no native q4 weighted Laplacian exists", Passed: a.Sieve.NativeAnisotropicCount > 0 && a.Sieve.NativeQuarticCapacityCount > 0 && a.Sieve.CanonicalHphiQ4MatchCount == 0, Detail: FormatSieve(a.Sieve)},
			{Name: "identity and flavor firewalls are preserved", Passed: !a.Impact.HphiQuarticIdentified && !a.Impact.CanonicalWeightedLaplacianFound && !a.Impact.YukawaCouplingsReduced && a.Impact.ChargedModuliResult == Gate372ChargedModuliDim && a.Impact.FlavorFirewallPreserved && a.Impact.HiggsLanePreserved, Detail: FormatImpact(a.Impact)},
			{Name: "empirical and arbitrary-identification firewalls remain clean", Passed: a.Firewall.NoObservedMassesImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoYukawaAmplitudesInserted && a.Firewall.NoManualQ4HphiID && a.Firewall.NoArbitraryEdgeComponentMap && a.Firewall.NoAffineChargeFitPromoted && a.Firewall.NoFlavorModuliReductionClaimed, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate searches spectral graph edge adjacency", Passed: a.Next.Gate == 402 && a.Next.Title != "", Detail: FormatNext(a.Next)},
		}, Notes: []string{a.Truth}}
	}}
}

func findCandidate(xs []WeightedCandidate, name string) WeightedCandidate {
	for _, x := range xs {
		if x.Name == name {
			return x
		}
	}
	return WeightedCandidate{Name: name, Verdict: "MISSING"}
}
