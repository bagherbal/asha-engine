package complexsectorsourcephase

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate416RealDim=%d Gate416ComplexDim=%d realNoCP=%t valuesBoundary=%t nativeFirewall=%t chargedDim=%d verdict=%s", x.Gate416RealLedgerDim, x.Gate416ComplexLedgerDim, x.Gate416RealNoCKMCP, x.Gate416ValuesRemainBoundaryData, x.Gate416NativeFirewallPreserved, x.ChargedFlavorModuliDim, x.Verdict)
}

func FormatAxiom(x ComplexPhaseAxiom) string {
	return fmt.Sprintf("%s: %s; charged=%v realCoeff/sector=%d hermitian=%t native=%t empirical=%t verdict=%s", x.Name, x.TextureExpression, x.ChargedSectors, x.RealCoefficientsPerSector, x.HermitianTextures, x.NativeToCurrentAsha, x.EmpiricalYukawaImported, x.Verdict)
}

func FormatCompatibility(x CompatibilityAudit) string {
	return fmt.Sprintf("gauge=%t J=%t Gamma=%t firstOrder=%t hermitian=%t residual=%.3g requiresPhaseAxiom=%t verdict=%s", x.GaugeCompatible, x.CompatibleWithJReal, x.CompatibleWithGamma, x.FirstOrderCompatible, x.HermiticityPreserved, x.CompatibilityResidual, x.RequiresNewPhaseAxiom, x.Verdict)
}

func FormatAlgebra(x AlgebraAudit) string {
	return fmt.Sprintf("rank=%d hermitianBasisDim=%d generatedComplexDim=%d ||[K,X]||=%.12f ||[K,Y]||=%.12f ||[X,Y]||=%.12f fullM3=%t native=%t verdict=%s", x.FamilyRank, x.HermitianBasisDimension, x.GeneratedComplexAlgebraDim, x.KXCommutatorNorm, x.KYCommutatorNorm, x.XYCommutatorNorm, x.SpansFullComplexMatrixSpace, x.Native, x.Verdict)
}

func FormatCPSample(x CPSample) string {
	return fmt.Sprintf("up=%v down=%v ||[Mu,Md]||=%.12f %s=%.12f CPCapacity=%t anglesPredicted=%t phasePredicted=%t valuesFixed=%t", x.UpCoefficients, x.DownCoefficients, x.UpDownCommutatorNorm, x.CPOddInvariantFormula, x.CPOddInvariant, x.NonzeroCPCapacity, x.CKMAnglesPredicted, x.CPPhasePredicted, x.CoefficientValuesFixed)
}

func FormatParameters(x ParameterCount) string {
	parts := []string{fmt.Sprintf("start=%d bestNative=%d bestReal=%d bestComplex=%d valuesFree=%t anglesUnderdetermined=%t", x.StartDim, x.BestNativeDim, x.BestConditionalRealDim, x.BestConditionalComplexDim, x.CoefficientValuesFree, x.CKMAnglesUnderdetermined)}
	for _, s := range x.Scenarios {
		parts = append(parts, fmt.Sprintf("%s[%s]: charged=%d totalNu=%d CKM=%t CP=%t fixed=%t native=%t conditional=%t", s.Name, s.Status, s.ChargedParameterCount, s.TotalWithNeutrinoCount, s.CKMCapacity, s.CPCapable, s.CoefficientValuesFixed, s.Native, s.Conditional))
	}
	return strings.Join(parts, " | ")
}

func FormatEmpirical(x EmpiricalIndependence) string {
	return fmt.Sprintf("massesImported=%t CKMImported=%t PMNSImported=%t YukawaMatrices=%t symbolicOnly=%t quarantined=%t verdict=%s", !x.NoObservedMassesImported, !x.NoCKMImported, !x.NoPMNSImported, !x.NoYukawaMatricesInserted, x.CoefficientSymbolsOnly, x.AxiomQuarantined, x.Verdict)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("nativeDim=%d conditionalDims=%v noNativeClaim=%t axiomStatus=%t preserved=%t verdict=%s", x.NativeDim, x.ConditionalAxiomDims, x.NoNativeDerivationClaimed, x.AxiomStatusPreserved, x.FirewallPreserved, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s", x.Gate, x.Title, x.Reason)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 417 Registry Audit — Complex Sector-Source CP-Phase Axiom Sieve\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Treat the smallest complex/phase extension of the Gate-416 sector-source axiom as an explicit quarantined axiom: add `Y_gen=i(S_gen-S_gen^T)` and audit whether CP-capable texture capacity appears without importing observed Yukawa data.\n\n")
	b.WriteString("## Prior boundary inherited\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("## Complex phase axiom\n\n")
	b.WriteString(FormatAxiom(a.Axiom) + "\n\n")
	b.WriteString("## Compatibility audit\n\n")
	b.WriteString(FormatCompatibility(a.Compatibility) + "\n\n")
	b.WriteString("## Family texture algebra\n\n")
	b.WriteString(FormatAlgebra(a.Algebra) + "\n\n")
	b.WriteString("## CP-capacity sample\n\n")
	b.WriteString(FormatCPSample(a.CPSample) + "\n\n")
	b.WriteString("## Parameter-count table\n\n")
	b.WriteString("| Scenario | Status | Charged parameters | With neutrino | CKM | CP-capable | CKM angles predicted | CP phase predicted | Coefficients fixed | Native? |\n")
	b.WriteString("|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|\n")
	for _, s := range a.Parameters.Scenarios {
		b.WriteString(fmt.Sprintf("| %s | %s | %d | %d | %t | %t | %t | %t | %t | %t |\n", s.Name, s.Status, s.ChargedParameterCount, s.TotalWithNeutrinoCount, s.CKMCapacity, s.CPCapable, s.CKMAnglesPredicted, s.CPPhasePredicted, s.CoefficientValuesFixed, s.Native))
	}
	b.WriteString("\n")
	b.WriteString("## Empirical firewall\n\n")
	b.WriteString(FormatEmpirical(a.Empirical) + "\n\n")
	b.WriteString("## Final firewall status\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("## Result statuses\n\n")
	for _, s := range []string{StatusGate416Inherited, StatusComplexPhaseAxiomFormalized, StatusShiftQuadratureAudited, StatusHermitianAlgebraDim9, StatusCPCapacityActivated, StatusParameterCountingCompleted, StatusAxiomQuarantinedNotNative, StatusFailedPhaseNotNative, StatusFailedPhaseCoefficientsFree, StatusFailedCPValueNotPredicted, StatusFailedAnglesUnderdetermined, StatusFailedNoNativeModuliReduction, StatusFirewallPreserved13Moduli} {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Next gate\n\n")
	b.WriteString(FormatNext(a.Next) + "\n\n")
	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}
