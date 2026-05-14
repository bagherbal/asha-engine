package minimalsectorsourceaxiom

import "github.com/bagherbal/asha-engine/pkg/theorem"

func MinimalSectorSourceAxiomConsistencyParameterCountingSieveTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Minimal sector-source axiom consistency / parameter-counting sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate416 audit", Passed: false, Detail: err.Error()}}}
		}
		var realOK, complexOK, observedRejected bool
		for _, f := range a.Families {
			if f.Name == "minimal real charge-sector source" && f.RealParameterCount == 6 && f.NoncommutingCapacity && !f.CPCapable && !f.Native {
				realOK = true
			}
			if f.Name == "minimal complex/phase sector source" && f.RealParameterCount == 9 && f.NoncommutingCapacity && f.CPCapable && !f.Native {
				complexOK = true
			}
			if f.Name == "unconstrained observed Yukawa source" && f.RealParameterCount == Gate372ChargedFlavorModuliDim && !f.Native {
				observedRejected = true
			}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate415 source-boundary ledger", Passed: a.Inheritance.Executed && a.Inheritance.Gate415LeastCostCKMCapableAxiom && a.Inheritance.Gate415ValuesRemainBoundaryData && a.Inheritance.ChargedModuliDim == Gate372ChargedFlavorModuliDim, Detail: FormatInheritance(a.Inheritance)},
			{Name: "minimal sector-source axiom formalized", Passed: a.Axiom.Executed && len(a.Axiom.ChargedSectors) == 3 && a.Axiom.RealCoefficientsPerSector == 2 && !a.Axiom.NativeToCurrentAsha && !a.Axiom.EmpiricalYukawaImported && !a.Axiom.PromotedToTheorem, Detail: FormatAxiom(a.Axiom)},
			{Name: "gauge/J/Gamma/first-order compatibility audited", Passed: a.Compatibility.Executed && a.Compatibility.GaugeCompatible && a.Compatibility.CompatibleWithJReal && a.Compatibility.CompatibleWithGamma && a.Compatibility.FirstOrderCompatible && a.Compatibility.RequiresNewSourceAxiom && !a.Compatibility.BreaksSMGaugeAction && !a.Compatibility.ObservedDataImported, Detail: FormatCompatibility(a.Compatibility)},
			{Name: "texture families counted", Passed: realOK && complexOK && observedRejected, Detail: RenderFamilies(a.Families)},
			{Name: "noncommuting sector criterion derived", Passed: a.Commutator.Executed && a.Commutator.KXCommutatorNorm > 0 && a.Commutator.NonzeroIfSectorRaysDiffer && a.Commutator.ZeroIfSectorRaysParallel && !a.Commutator.CoefficientsFixedByCriterion, Detail: FormatCommutator(a.Commutator)},
			{Name: "parameter-counting completed", Passed: a.Parameters.Executed && a.Parameters.StartDim == Gate372ChargedFlavorModuliDim && a.Parameters.BestNativeDim == Gate372ChargedFlavorModuliDim && a.Parameters.BestConditionalRealDim == 6 && a.Parameters.BestConditionalCPDim == 9 && !a.Parameters.NativeReductionBelow13 && a.Parameters.ConditionalReductionBelow13 && a.Parameters.CoefficientValuesFree, Detail: FormatParameters(a.Parameters)},
			{Name: "empirical firewall preserved", Passed: a.Empirical.Executed && a.Empirical.NoObservedMassesImported && a.Empirical.NoCKMImported && a.Empirical.NoPMNSImported && a.Empirical.NoYukawaMatricesInserted && a.Empirical.CoefficientSymbolsOnly && a.Empirical.AxiomQuarantined, Detail: FormatEmpirical(a.Empirical)},
			{Name: "native 13-moduli firewall preserved", Passed: a.Firewall.Executed && a.Firewall.NativeDim == Gate372ChargedFlavorModuliDim && a.Firewall.FirewallPreserved && a.Firewall.NoNativeDerivationClaimed && a.Firewall.AxiomStatusPreserved, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate targets CP/phase extension", Passed: a.Next.Gate == 417 && a.Next.Title == "Complex Sector-Source CP-Phase Axiom Sieve", Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth}}
	}}
}
