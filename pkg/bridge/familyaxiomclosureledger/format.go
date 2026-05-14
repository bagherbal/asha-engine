package familyaxiomclosureledger

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("G411Ledger=%t G412Hierarchy=%t G413Mixing=%t G414NoSelector=%t G415SourceLedger=%t G416RealDim=%d G417ComplexDim=%d CP=%t valuesFree=%t chargedDim=%d verdict=%s", x.Gate411AxiomLedgerCompiled, x.Gate412HierarchyCapacity, x.Gate413MixingCapacity, x.Gate414NoCoefficientSelector, x.Gate415SourceLedgerCompiled, x.Gate416RealLedgerDim, x.Gate417ComplexLedgerDim, x.Gate417CPCapacity, x.Gate417CoefficientValuesFree, x.ChargedFlavorModuliDim, x.Verdict)
}

func FormatProgression(x AxiomProgression) string {
	return fmt.Sprintf("steps=%d minimalHierarchy=G%d minimalMixing=G%d minimalCP=G%d allQuarantined=%t verdict=%s", len(x.Steps), x.MinimalHierarchyGate, x.MinimalMixingGate, x.MinimalCPGate, x.AllAxiomsQuarantined, x.Verdict)
}

func FormatStep(s AxiomStep) string {
	return fmt.Sprintf("G%d %s: ops=[%s] hierarchy=%t mixing=%t CP=%t gauge=%t J/Gamma=%t firstOrder=%t native=%t theorem=%t valuesImported=%t fixed=%t remaining=%d verdict=%s", s.Gate, s.Name, joinOps(s.Operators), s.AddsHierarchy, s.AddsRealMixing, s.AddsCPCapacity, s.GaugeCompatible, s.CompatibleWithJGamma, s.FirstOrderCompatible, s.NativeToCurrentAsha, s.PromotedToTheorem, s.EmpiricalValuesImported, s.CoefficientsFixed, s.RemainingFreeLedger, s.Verdict)
}

func FormatParameters(x ParameterReduction) string {
	parts := []string{fmt.Sprintf("start=%d nativeDim=%d conditionalDim=%d conditionalCompression=%t nativeCompression=%t nineSymbolic=%t valuesPredicted=%t", x.StartDim, x.NativeDim, x.ConditionalCompressedDim, x.ConditionalCompression, x.NativeCompression, x.NineCoefficientsSymbolic, x.CKMAndPhaseValuesPredicted)}
	for _, r := range x.Rows {
		parts = append(parts, fmt.Sprintf("%s[%s]: coords=%d hierarchy=%t CKM=%t CP=%t predicted=%t environmental=%t native=%t conditional=%t", r.Name, r.Status, r.ChargedCoordinates, r.HierarchyCapacity, r.CKMCapacity, r.CPCapacity, r.ValuesPredicted, r.ValuesEnvironmental, r.Native, r.Conditional))
	}
	return strings.Join(parts, " | ")
}

func FormatSeal(x EnvironmentalSeal) string {
	return fmt.Sprintf("%s: nativeLawSpace=%t flavorCapacity=%t valuesPredicted=%t environmental=%t historical=%t noFitting=%t noNativeCollapse=%t verdict=%s", x.Name, x.NativeLawSpaceComplete, x.FlavorCapacityComplete, x.CoefficientValuesPredicted, x.CoefficientsEnvironmental, x.RequiresHistoricalData, x.NoEmpiricalFitting, x.NoNativeCollapseClaimed, x.Verdict)
}

func FormatEmpirical(x EmpiricalFirewall) string {
	return fmt.Sprintf("massesImported=%t CKMImported=%t PMNSImported=%t YukawaMatrices=%t symbolicOnly=%t rejectsFitting=%t verdict=%s", !x.NoObservedMassesImported, !x.NoCKMImported, !x.NoPMNSImported, !x.NoYukawaMatricesInserted, x.SymbolicCoefficientsOnly, x.RejectsCurveFitting, x.Verdict)
}

func FormatFinal(x FinalSeal) string {
	return fmt.Sprintf("nativeDim=%d conditionalComplexDim=%d formallySealed=%t completeAsLedger=%t noNativePrediction=%t firewall=%t status=%s verdict=%s", x.NativeChargedDim, x.ConditionalComplexDim, x.FlavorSectorFormallySealed, x.ProjectFlavorCompleteAsLedger, x.NoNativePredictionClaimed, x.FirewallPreserved, x.FinalStatus, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s", x.Gate, x.Title, x.Reason)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 418 Registry Audit — Family-Axiom Closure Ledger / Flavor Frontier Seal\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Compile the capstone ledger for Gates 411–417, separating native ASHA law-space from quarantined family axioms and environmental flavor coefficients. Gate 418 does not search for another hidden selector; it closes the current flavor frontier with explicit firewalls.\n\n")
	b.WriteString("## Prior boundary inherited\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("## Axiom progression ledger\n\n")
	b.WriteString(FormatProgression(a.Progression) + "\n\n")
	b.WriteString("| Gate | Axiom / test | Operators | Hierarchy | Mixing | CP | Native? | Coefficients fixed? | Remaining ledger | Boundary |\n")
	b.WriteString("|---:|---|---|---:|---:|---:|---:|---:|---:|---|\n")
	for _, s := range a.Progression.Steps {
		b.WriteString(fmt.Sprintf("| %d | %s | `%s` | %t | %t | %t | %t | %t | %d | %s |\n", s.Gate, s.Name, joinOps(s.Operators), s.AddsHierarchy, s.AddsRealMixing, s.AddsCPCapacity, s.NativeToCurrentAsha, s.CoefficientsFixed, s.RemainingFreeLedger, s.Boundary))
	}
	b.WriteString("\n## Parameter reduction summary\n\n")
	b.WriteString(FormatParameters(a.Parameters) + "\n\n")
	b.WriteString("| Scenario | Status | Charged coordinates | Hierarchy | CKM | CP | Values predicted | Environmental | Native? |\n")
	b.WriteString("|---|---|---:|---:|---:|---:|---:|---:|---:|\n")
	for _, r := range a.Parameters.Rows {
		b.WriteString(fmt.Sprintf("| %s | %s | %d | %t | %t | %t | %t | %t | %t |\n", r.Name, r.Status, r.ChargedCoordinates, r.HierarchyCapacity, r.CKMCapacity, r.CPCapacity, r.ValuesPredicted, r.ValuesEnvironmental, r.Native))
	}
	b.WriteString("\n## Epistemological boundary seal\n\n")
	b.WriteString(FormatSeal(a.Seal) + "\n\n")
	b.WriteString(a.Seal.Statement + "\n\n")
	b.WriteString("## Empirical firewall\n\n")
	b.WriteString(FormatEmpirical(a.Empirical) + "\n\n")
	b.WriteString("## Final status\n\n")
	b.WriteString(FormatFinal(a.Final) + "\n\n")
	b.WriteString("## Result statuses\n\n")
	for _, s := range Statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Next gate\n\n")
	b.WriteString(FormatNext(a.Next) + "\n\n")
	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}
