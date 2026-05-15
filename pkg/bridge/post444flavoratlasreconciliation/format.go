package post444flavoratlasreconciliation

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate420=%t acyclic=%t nativeFlavorDim=%d conditionalDim=%d familySealed=%t noReopen=%t noEmpirical=%t verdict=%s", x.Gate420PublicationAtlas, x.Gate420Acyclic, x.Gate420NativeFlavorDim, x.Gate420ConditionalDim, x.Gate420FamilyAxiomsSealed, x.Gate420NoFlavorReopening, x.NoEmpiricalInputsImported, x.Verdict)
}

func FormatGateDelta(x GateDelta) string {
	return fmt.Sprintf("G%d %s input=%s output=%s structural=%t firewall=%t observable=%t nativeDim=%d KXY=%d verdict=%s reason=%s", x.Gate, x.Name, x.InputStatus, x.OutputStatus, x.PromotesStructuralObject, x.PreservesFirewall, x.PredictsObservableValue, x.NativeChargedDimAfter, x.KXYCoeffDimAfter, x.Verdict, x.Reason)
}

func FormatReclassification(x Reclassification) string {
	return fmt.Sprintf("%s %s→%s status=%s→%s promoted=%t quarantined=%t valueBearing=%t verdict=%s reason=%s", x.Object, x.PreviousLayer, x.ReconciledLayer, x.PreviousStatus, x.ReconciledStatus, x.Promoted, x.Quarantined, x.ValueBearing, x.Verdict, x.Reason)
}

func FormatDelta(x ReconciliationDelta) string {
	return fmt.Sprintf("deltas=%d reclasses=%d promoted=%d quarantined=%d nativeDim=%d→%d KXY=%d→%d valuesAdded=%d selectorsAdded=%d verdict=%s reason=%s", len(x.Deltas), len(x.Reclassifications), x.PromotedObjects, x.QuarantinedObjects, x.NativeDimBefore, x.NativeDimAfter, x.KXYCoeffDimBefore, x.KXYCoeffDimAfter, x.FlavorObservableValuesAdded, x.CoefficientSelectorsAdded, x.Verdict, x.Reason)
}

func FormatNode(x AtlasNode) string {
	return fmt.Sprintf("%s [%s] gates=%s status=%s deps=%s claim=%s boundary=%s", x.ID, x.Layer, joinInts(x.Gates), x.Status, strings.Join(x.DependsOn, ","), x.Claim, x.Boundary)
}

func FormatAtlas(x ReconciledAtlas) string {
	return fmt.Sprintf("nodes=%d acyclic=%t nativeDim=%d KXY=%d K=%t G2Zero=%t X=%t YQuarantined=%t coeffQuarantined=%t noNewPhysics=%t verdict=%s", len(x.Nodes), x.Acyclic, x.NativeFlavorDim, x.KXYCoeffDim, x.KGenGeometric, x.Gen2BareZeroStructural, x.XTriangleSupportStructural, x.YPhaseQuarantined, x.CoefficientsQuarantined, x.NoNewPhysicsClaim, x.Verdict)
}

func FormatFirewall(x Firewall) string {
	return fmt.Sprintf("muonImported=%t charmImported=%t yukawaImported=%t CKM=%t PMNS=%t poleFit=%t curveFit=%t cosmology=%t nativeDim=%t KXY=%t verdict=%s reason=%s", !x.NoObservedMuonMassImported, !x.NoObservedCharmMassImported, !x.NoObservedYukawaImported, !x.NoCKMImported, !x.NoPMNSImported, !x.NoPoleMassFit, !x.NoCurveFit, !x.NoCosmologyInput, x.NativeFlavorDimPreserved, x.KXYCoeffDimPreserved, x.Verdict, x.Reason)
}

func FormatPatch(x RegistryPatch) string {
	return fmt.Sprintf("package=%s theorem=%s audit=%s runtime=%t atlasOverlay=%t reopensG420=%t rewrite=%t ready=%t verdict=%s reason=%s", x.Package, x.Theorem, x.AuditPath, x.RuntimeFamilyUpdated, x.PublicationAtlasOverlay, x.ReopensGate420, x.RequiresAtlasRewrite, x.Ready, x.Verdict, x.Reason)
}

func FormatFinal(x FinalStatus) string {
	return fmt.Sprintf("reconciled=%t K=%t G2Zero=%t X=%t YQuarantined=%t coeffQuarantined=%t nativeDim=%d KXY=%d noMass=%t noMixing=%t status=%s verdict=%s", x.Reconciled, x.KGenPromoted, x.Gen2ZeroPromoted, x.XSupportPromoted, x.YPhaseStillQuarantined, x.CoefficientsStillQuarantined, x.NativeFlavorDim, x.KXYCoeffDim, x.NoMassPrediction, x.NoMixingPrediction, x.Status, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Task=%s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func renderNodeTable(nodes []AtlasNode) string {
	var b strings.Builder
	b.WriteString("| ID | Layer | Gates | Status | Dependencies | Claim | Boundary |\n")
	b.WriteString("|---|---|---|---|---|---|---|\n")
	for _, n := range nodes {
		b.WriteString(fmt.Sprintf("| `%s` | %s | %s | `%s` | %s | %s | %s |\n", n.ID, n.Layer, joinInts(n.Gates), n.Status, strings.Join(n.DependsOn, ", "), n.Claim, n.Boundary))
	}
	return b.String()
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 448 Registry Audit — Post-444 Flavor Frontier Atlas Reconciliation\n\n")
	b.WriteString("## Scope\n\n")
	b.WriteString("Gate 448 reconciles the Gate-420 publication atlas with the later Generation-2 intersection sieves. It is an atlas/registry patch, not a phenomenological fit and not a rewrite of observed flavor physics.\n\n")

	b.WriteString("## Gate 420 inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Gate 444-447 delta\n\n")
	b.WriteString(FormatDelta(a.Delta) + "\n\n")
	b.WriteString("| Gate | Input status | Output status | Structural promotion | Firewall preserved | Observable predicted | Verdict | Reason |\n")
	b.WriteString("|---:|---|---|---:|---:|---:|---|---|\n")
	for _, x := range a.Delta.Deltas {
		b.WriteString(fmt.Sprintf("| %d | %s | %s | %t | %t | %t | `%s` | %s |\n", x.Gate, x.InputStatus, x.OutputStatus, x.PromotesStructuralObject, x.PreservesFirewall, x.PredictsObservableValue, x.Verdict, x.Reason))
	}
	b.WriteString("\n")

	b.WriteString("## Reclassification ledger\n\n")
	b.WriteString("| Object | Previous layer | Reconciled layer | Previous status | Reconciled status | Promoted | Quarantined | Value-bearing | Reason |\n")
	b.WriteString("|---|---|---|---|---|---:|---:|---:|---|\n")
	for _, x := range a.Delta.Reclassifications {
		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | `%s` | %t | %t | %t | %s |\n", x.Object, x.PreviousLayer, x.ReconciledLayer, x.PreviousStatus, x.ReconciledStatus, x.Promoted, x.Quarantined, x.ValueBearing, x.Reason))
	}
	b.WriteString("\n")

	b.WriteString("## Reconciled atlas overlay\n\n")
	b.WriteString(FormatAtlas(a.Atlas) + "\n\n")
	b.WriteString(renderNodeTable(a.Atlas.Nodes) + "\n")

	b.WriteString("The reconciled structural layer is:\n\n")
	b.WriteString("```text\n")
	b.WriteString("K_gen = diag(-1,0,1)                                  // Gate 444, geometric axis\n")
	b.WriteString("Gen2 bare level = 0                                    // Gate 444, structural zero\n")
	b.WriteString("support(B_lift) = complete endpoint-balanced triangle  // Gate 445, topology only\n")
	b.WriteString("Φ_cycle, Y_gen, ε, sector K/X/Y coefficients           // Gates 446-447, quarantined\n")
	b.WriteString("dim M_charged^native = 13; dim C_KXY^charged = 9       // preserved firewall\n")
	b.WriteString("```\n\n")

	b.WriteString("## Registry patch\n\n")
	b.WriteString(FormatPatch(a.Patch) + "\n\n")

	b.WriteString("## Empirical/firewall audit\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("No observed muon/charm mass, Yukawa entry, CKM angle, CKM phase, PMNS value, pole-mass fit, curve-fit, or cosmological coordinate is used.\n\n")

	b.WriteString("## Result statuses\n\n")
	for _, s := range Statuses() {
		b.WriteString("- `" + s + "`\n")
	}

	b.WriteString("\n## Final status\n\n")
	b.WriteString(FormatFinal(a.Final) + "\n\n")
	b.WriteString("## Next gate\n\n")
	b.WriteString(FormatNext(a.Next) + "\n\n")
	b.WriteString("## Truth statement\n\n")
	b.WriteString(a.Truth + "\n")
	return b.String()
}
