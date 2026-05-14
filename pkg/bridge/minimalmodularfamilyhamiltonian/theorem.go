package minimalmodularfamilyhamiltonian

import "github.com/bagherbal/asha-engine/pkg/theorem"

func MinimalModularFamilyHamiltonianAxiomConsistencySieveTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Minimal modular family Hamiltonian axiom consistency sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate412 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate411 lowest-cost axiom boundary", Passed: a.Inheritance.Executed && a.Inheritance.Gate411LeastCostKGen && a.Inheritance.Gate411NoAxiomPromoted && a.Inheritance.Gate410NoNativeFamilyBundle && a.Inheritance.Gate409TrivialU3Multiplicity && a.Inheritance.Gate408ScalarFlavorBlind && a.Inheritance.Gate372ChargedModuliDim == Gate372ChargedFlavorModuliDim && a.Inheritance.NoEmpiricalInputsImported, Detail: FormatInheritance(a.Inheritance)},
			{Name: "minimal K_gen axiom is formalized but not native", Passed: a.Hamiltonian.Executed && !a.Hamiltonian.NativeInCurrentAsha && a.Hamiltonian.ExplicitAxiom && a.Hamiltonian.Hermitian && a.Hamiltonian.Traceless && a.Hamiltonian.DistinctEigenvalues == FamilyRank && a.Hamiltonian.MinimalPolynomialDegree == FamilyRank && a.Hamiltonian.ProvidesThreeLevelOrder && a.Hamiltonian.DiagonalOnly && !a.Hamiltonian.CoefficientsEmpirical, Detail: FormatHamiltonian(a.Hamiltonian)},
			{Name: "nontracial KMS family state is activated", Passed: a.KMS.Executed && a.KMS.Positive && a.KMS.Normalized && !a.KMS.Tracial && a.KMS.ModularFlowActive && a.KMS.MaxWeightRatio > 1, Detail: FormatKMS(a.KMS)},
			{Name: "family Hamiltonian is gauge-compatible as an axiom", Passed: a.Compatibility.Executed && a.Compatibility.ActsOnlyOnFamilyFiber && a.Compatibility.CommutesWithAF && a.Compatibility.CommutesWithGaugeCharges && a.Compatibility.CommutesWithHypercharge && a.Compatibility.CommutesWithSU2L && a.Compatibility.CommutesWithBL && a.Compatibility.CompatibleWithGamma && a.Compatibility.JCompatibleIfMirrored && a.Compatibility.FirstOrderUnaffectedIfDFBroadcast && a.Compatibility.RequiresFamilyFiberAxiom, Detail: FormatCompatibility(a.Compatibility)},
			{Name: "single Hamiltonian remains diagonal-only", Passed: a.Mixing.Executed && a.Mixing.NativeNoncommutingPairs == 0 && a.Mixing.ConditionalNoncommutingPairs == 0 && a.Mixing.CommutatorKWithK2Norm == 0 && !a.Mixing.CKMNative && !a.Mixing.PMNSNative && !a.Mixing.CKMConditional && !a.Mixing.PMNSConditional && a.Mixing.DiagonalOnly, Detail: FormatMixing(a.Mixing)},
			{Name: "sector amplitude map still needs another rule", Passed: a.SectorMap.Executed && a.SectorMap.UniversalFamilyOrdering && !a.SectorMap.UpSectorMapNative && !a.SectorMap.DownSectorMapNative && !a.SectorMap.LeptonSectorMapNative && a.SectorMap.SectorSpecificMapsNeeded && !a.SectorMap.ObservedYukawasInserted && a.SectorMap.MassHierarchyCapacity, Detail: FormatSectorMap(a.SectorMap)},
			{Name: "empirical firewall is preserved", Passed: a.Firewall.Executed && a.Firewall.NoObservedMassesImported && a.Firewall.NoCKMImported && a.Firewall.NoPMNSImported && a.Firewall.NoYukawaMatricesInserted && a.Firewall.NoSectorAmplitudesInserted && a.Firewall.KGenPromotedAsAxiomOnly && a.Firewall.NoNativeDerivationClaimed, Detail: FormatFirewall(a.Firewall)},
			{Name: "13 charged flavor moduli remain native firewall", Passed: a.Moduli.StartDim == Gate372ChargedFlavorModuliDim && a.Moduli.BestNativeDim == Gate372ChargedFlavorModuliDim && !a.Moduli.NativeReductionBelow13 && a.Moduli.ConditionalHierarchy && !a.Moduli.ConditionalCKMPMNS && a.Moduli.FirewallPreserved, Detail: FormatModuli(a.Moduli)},
			{Name: "next gate targets a second noncommuting operator", Passed: a.Next.Gate == 413 && a.Next.Title == "Second Family Operator / Noncommuting Modular Pair Axiom Sieve", Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth}}
	}}
}
