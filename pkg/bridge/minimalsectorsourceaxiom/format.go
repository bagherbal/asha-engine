package minimalsectorsourceaxiom

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate413Capacity=%t Gate414NoSelector=%t Gate415LeastCostCKMCapable=%t valuesBoundary=%t chargedDim=%d verdict=%s", x.Gate413NoncommutingCapacity, x.Gate414NoCoefficientSelector, x.Gate415LeastCostCKMCapableAxiom, x.Gate415ValuesRemainBoundaryData, x.ChargedModuliDim, x.Verdict)
}

func FormatAxiom(x SectorSourceAxiom) string {
	return fmt.Sprintf("%s: charged=%v neutral=%v realCoeff/sector=%d phaseCoeff/sector=%d gaugeBlind=%t native=%t empirical=%t verdict=%s", x.Name, x.ChargedSectors, x.NeutralSectors, x.RealCoefficientsPerSector, x.PhaseCoefficientsPerSector, x.GaugeBlindFamilyFiber, x.NativeToCurrentAsha, x.EmpiricalYukawaImported, x.Verdict)
}

func FormatCompatibility(x CompatibilityAudit) string {
	return fmt.Sprintf("gauge=%t J=%t Gamma=%t firstOrder=%t residual=%.3g requiresAxiom=%t verdict=%s", x.GaugeCompatible, x.CompatibleWithJReal, x.CompatibleWithGamma, x.FirstOrderCompatible, x.CompatibilityResidual, x.RequiresNewSourceAxiom, x.Verdict)
}

func RenderFamilies(xs []TextureFamily) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		parts = append(parts, fmt.Sprintf("%s{params=%d noncomm=%t CP=%t native=%t verdict=%s}", x.Name, x.RealParameterCount, x.NoncommutingCapacity, x.CPCapable, x.Native, x.Verdict))
	}
	return strings.Join(parts, "; ")
}

func FormatCommutator(x CommutatorAudit) string {
	return fmt.Sprintf("criterion=%q ||[K,X]||=%.12f wedge=%.12f sample||[Mu,Md]||=%.12f valuesFixed=%t", x.Criterion, x.KXCommutatorNorm, x.SectorWedgeDeterminant, x.SampleMassCommutatorNorm, x.CoefficientsFixedByCriterion)
}

func FormatParameters(x ParameterCount) string {
	parts := []string{fmt.Sprintf("start=%d bestNative=%d bestConditionalReal=%d bestConditionalCP=%d valuesFree=%t", x.StartDim, x.BestNativeDim, x.BestConditionalRealDim, x.BestConditionalCPDim, x.CoefficientValuesFree)}
	for _, s := range x.Scenarios {
		parts = append(parts, fmt.Sprintf("%s[%s]: charged=%d totalNu=%d CKM=%t CP=%t fixed=%t native=%t conditional=%t", s.Name, s.Status, s.ChargedParameterCount, s.TotalWithNeutrinoCount, s.CKMCapacity, s.CKMCPPhaseCapacity, s.CoefficientsFixed, s.Native, s.Conditional))
	}
	return strings.Join(parts, " | ")
}

func FormatEmpirical(x EmpiricalIndependence) string {
	return fmt.Sprintf("masses=%t CKM=%t PMNS=%t YukawaMatrices=%t symbolicOnly=%t quarantined=%t verdict=%s", !x.NoObservedMassesImported, !x.NoCKMImported, !x.NoPMNSImported, !x.NoYukawaMatricesInserted, x.CoefficientSymbolsOnly, x.AxiomQuarantined, x.Verdict)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("nativeDim=%d conditionalDims=%v axiomStatus=%t noNativeClaim=%t preserved=%t verdict=%s", x.NativeDim, x.ConditionalAxiomDims, x.AxiomStatusPreserved, x.NoNativeDerivationClaimed, x.FirewallPreserved, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s", x.Gate, x.Title, x.Reason)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 416 Registry Audit — Minimal Sector-Source Axiom Consistency / Parameter-Counting Sieve\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Treat the Gate-415 charge-sector source boundary as an explicit quarantined axiom and count the remaining charged flavor parameters under gauge, J, Gamma, and first-order compatibility.\n\n")
	b.WriteString("## Prior boundary inherited\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("## Axiom formalization\n\n")
	b.WriteString(FormatAxiom(a.Axiom) + "\n\n")
	b.WriteString("## Compatibility audit\n\n")
	b.WriteString(FormatCompatibility(a.Compatibility) + "\n\n")
	b.WriteString("## Texture-family ledger\n\n")
	for _, f := range a.Families {
		b.WriteString(fmt.Sprintf("- **%s** — `%s`; parameters=%d; noncommuting=%t; CP-capable=%t; native=%t. %s\n", f.Name, f.Expression, f.RealParameterCount, f.NoncommutingCapacity, f.CPCapable, f.Native, f.Reason))
	}
	b.WriteString("\n## Noncommuting criterion\n\n")
	b.WriteString(FormatCommutator(a.Commutator) + "\n\n")
	b.WriteString("## Parameter-count table\n\n")
	b.WriteString("| Scenario | Status | Charged parameters | With neutrino | CKM | CP phase | Coefficients fixed | Native? |\n")
	b.WriteString("|---|---:|---:|---:|---:|---:|---:|---:|\n")
	for _, s := range a.Parameters.Scenarios {
		b.WriteString(fmt.Sprintf("| %s | %s | %d | %d | %t | %t | %t | %t |\n", s.Name, s.Status, s.ChargedParameterCount, s.TotalWithNeutrinoCount, s.CKMCapacity, s.CKMCPPhaseCapacity, s.CoefficientsFixed, s.Native))
	}
	b.WriteString("\n")
	b.WriteString("## Empirical firewall\n\n")
	b.WriteString(FormatEmpirical(a.Empirical) + "\n\n")
	b.WriteString("## Final firewall status\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("## Result statuses\n\n")
	for _, s := range []string{StatusGate415BoundaryInherited, StatusMinimalSectorSourceFormalized, StatusCompatibilityAudited, StatusNoncommutingCriterionDerived, StatusParameterCountingCompleted, StatusRealSourceSixParameterLedger, StatusComplexPhaseExtensionAudited, StatusAxiomQuarantinedNotNative, StatusFailedSourceNotNative, StatusFailedCoefficientsRemainFree, StatusFailedRealNoCPPhase, StatusFailedFullCKMNeedsPhaseAxiom, StatusFailedNoNativeModuliReduction, StatusFirewallPreserved13Moduli} {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Next gate\n\n")
	b.WriteString(FormatNext(a.Next) + "\n\n")
	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}
