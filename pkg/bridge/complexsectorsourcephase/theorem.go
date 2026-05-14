package complexsectorsourcephase

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ComplexSectorSourceCPPhaseAxiomSieveTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Complex sector-source CP-phase axiom sieve"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate417 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate416 real-sector CP boundary", Passed: a.Inheritance.Executed && a.Inheritance.Gate416RealLedgerDim == 6 && a.Inheritance.Gate416ComplexLedgerDim == 9 && a.Inheritance.Gate416RealNoCKMCP && a.Inheritance.Gate416NativeFirewallPreserved, Detail: FormatInheritance(a.Inheritance)},
			{Name: "minimal complex phase axiom formalized", Passed: a.Axiom.Executed && a.Axiom.RealCoefficientsPerSector == 3 && a.Axiom.HermitianTextures && !a.Axiom.NativeToCurrentAsha && !a.Axiom.EmpiricalYukawaImported && !a.Axiom.PromotedToTheorem, Detail: FormatAxiom(a.Axiom)},
			{Name: "gauge/J/Gamma/first-order compatibility audited", Passed: a.Compatibility.Executed && a.Compatibility.GaugeCompatible && a.Compatibility.CompatibleWithJReal && a.Compatibility.CompatibleWithGamma && a.Compatibility.FirstOrderCompatible && a.Compatibility.HermiticityPreserved && a.Compatibility.RequiresNewPhaseAxiom && !a.Compatibility.BreaksSMGaugeAction && !a.Compatibility.ObservedDataImported, Detail: FormatCompatibility(a.Compatibility)},
			{Name: "K/X/Y full family texture algebra audited", Passed: a.Algebra.Executed && a.Algebra.HermitianBasisDimension == 9 && a.Algebra.GeneratedComplexAlgebraDim == 9 && a.Algebra.SpansFullComplexMatrixSpace && a.Algebra.KXCommutatorNorm > 0 && a.Algebra.KYCommutatorNorm > 0 && !a.Algebra.Native, Detail: FormatAlgebra(a.Algebra)},
			{Name: "CP-odd capacity activated conditionally", Passed: a.CPSample.Executed && a.CPSample.UpDownCommutatorNorm > 0 && a.CPSample.NonzeroCPCapacity && !a.CPSample.CoefficientValuesFixed && !a.CPSample.CKMAnglesPredicted && !a.CPSample.CPPhasePredicted, Detail: FormatCPSample(a.CPSample)},
			{Name: "complex phase parameter count completed", Passed: a.Parameters.Executed && a.Parameters.StartDim == Gate372ChargedFlavorModuliDim && a.Parameters.BestNativeDim == Gate372ChargedFlavorModuliDim && a.Parameters.BestConditionalRealDim == 6 && a.Parameters.BestConditionalComplexDim == 9 && !a.Parameters.NativeReductionBelow13 && a.Parameters.ConditionalCPBelow13 && a.Parameters.CoefficientValuesFree && a.Parameters.CKMAnglesUnderdetermined, Detail: FormatParameters(a.Parameters)},
			{Name: "empirical firewall preserved", Passed: a.Empirical.Executed && a.Empirical.NoObservedMassesImported && a.Empirical.NoCKMImported && a.Empirical.NoPMNSImported && a.Empirical.NoYukawaMatricesInserted && a.Empirical.CoefficientSymbolsOnly && a.Empirical.AxiomQuarantined, Detail: FormatEmpirical(a.Empirical)},
			{Name: "native 13-moduli firewall preserved", Passed: a.Firewall.Executed && a.Firewall.NativeDim == Gate372ChargedFlavorModuliDim && a.Firewall.FirewallPreserved && a.Firewall.NoNativeDerivationClaimed && a.Firewall.AxiomStatusPreserved, Detail: FormatFirewall(a.Firewall)},
			{Name: "next gate closes family axiom frontier", Passed: a.Next.Gate == 418 && a.Next.Title == "Family-Axiom Closure Ledger / Flavor Frontier Seal", Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth}}
	}}
}
