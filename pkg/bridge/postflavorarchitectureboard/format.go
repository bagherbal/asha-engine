package postflavorarchitectureboard

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("G376=%t G387=%t scalarBlindness=%t fermionTriviality=%t familySeal=%t flavorDim=%d conditionalFamilyDim=%d verdict=%s", x.Gate376ProductGeometry, x.Gate387FinalArchitectureLedger, x.Gate398To408ScalarBlindness, x.Gate409To410FermionTriviality, x.Gate411To418FamilyAxiomSeal, x.FlavorFirewallDim, x.ConditionalFamilyAxiomDim, x.Verdict)
}

func FormatBoard(x ArchitectureBoard) string {
	return fmt.Sprintf("nodes=%d native=%d bridge=%d quarantined=%d environmental=%d ordered=%t verdict=%s", len(x.Nodes), x.NativeCount, x.BridgeCount, x.QuarantinedCount, x.EnvironmentalCount, x.Ordered, x.Verdict)
}

func FormatNode(n BoardNode) string {
	return fmt.Sprintf("F%d %s [%s] gates=%s native=%t bridge=%t quarantined=%t environmental=%t claim=%s boundary=%s", n.Floor, n.Slug, n.Layer, joinInts(n.CoreGates), n.Native, n.Bridge, n.Quarantined, n.Environmental, n.Claim, n.Boundary)
}

func FormatTheorems(x TheoremLedger) string {
	return fmt.Sprintf("items=%d nativeLawSpace=%t flavorCapacity=%t coefficientPrediction=%t verdict=%s", len(x.Items), x.NativeLawSpaceComplete, x.FlavorCapacityClassified, x.CoefficientPredictionClaimed, x.Verdict)
}

func FormatFrontiers(x FrontierLedger) string {
	return fmt.Sprintf("frontiers=%d flavorFirewall=%t cosmologyFirewall=%t noEmpirical=%t verdict=%s", len(x.Frontiers), x.FlavorFirewallPreserved, x.CosmologyFirewallPreserved, x.NoEmpiricalValuesInserted, x.Verdict)
}

func FormatPublication(x PublicationBoard) string {
	return fmt.Sprintf("%s: chain=%d ready=%t nextUse=%s verdict=%s", x.Title, len(x.EssentialChain), x.Ready, x.NextUse, x.Verdict)
}

func FormatFinal(x FinalStatus) string {
	return fmt.Sprintf("boardReady=%t nativeFlavorDim=%d conditionalFamilyDim=%d noPrediction=%t noAxiomPromotion=%t noFlavorReopening=%t status=%s verdict=%s", x.BoardReady, x.NativeFlavorDim, x.ConditionalFamilyDim, x.NoNativeFlavorPrediction, x.NoAxiomPromotion, x.NoFlavorReopening, x.Status, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s", x.Gate, x.Title, x.Reason)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 419 Registry Audit — Post-Flavor Architecture Consolidation / Final Law-Space Board\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Consolidate the mature ASHA architecture after the Gate-418 flavor seal. Gate 419 does not search for new flavor predictors; it builds the final law-space board separating native theorems, bridge lanes, quarantined family axioms, and environmental coordinates.\n\n")
	b.WriteString("## Gate 418 boundary inherited\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("## Final architecture board\n\n")
	b.WriteString(FormatBoard(a.Board) + "\n\n")
	b.WriteString("| Floor | Slug | Layer | Core gates | Claim | Boundary |\n")
	b.WriteString("|---:|---|---|---|---|---|\n")
	for _, n := range a.Board.Nodes {
		b.WriteString(fmt.Sprintf("| %d | `%s` | %s | %s | %s | %s |\n", n.Floor, n.Slug, n.Layer, joinInts(n.CoreGates), n.Claim, n.Boundary))
	}
	b.WriteString("\n## Consolidated theorem ledger\n\n")
	b.WriteString(FormatTheorems(a.Theorems) + "\n\n")
	b.WriteString("| Theorem block | Native | Conditional | Quarantined | Environmental | Claim | Firewall |\n")
	b.WriteString("|---|---:|---:|---:|---:|---|---|\n")
	for _, t := range a.Theorems.Items {
		b.WriteString(fmt.Sprintf("| %s | %t | %t | %t | %t | %s | %s |\n", t.Name, t.Native, t.Conditional, t.Quarantined, t.Environmental, t.Claim, t.Firewall))
	}
	b.WriteString("\n## Environmental frontier ledger\n\n")
	b.WriteString(FormatFrontiers(a.Frontiers) + "\n\n")
	b.WriteString("| Frontier | Native dim | Conditional dim | Status | Reopened? | Environmental coordinates |\n")
	b.WriteString("|---|---:|---:|---|---:|---|\n")
	for _, f := range a.Frontiers.Frontiers {
		b.WriteString(fmt.Sprintf("| %s | %d | %d | %s | %t | %s |\n", f.Name, f.NativeDim, f.ConditionalDim, f.Status, f.Reopened, strings.Join(f.EnvironmentalCoordinates, ", ")))
	}
	b.WriteString("\n## Publication board\n\n")
	b.WriteString(FormatPublication(a.Publication) + "\n\n")
	b.WriteString("Essential chain:\n\n")
	for i, x := range a.Publication.EssentialChain {
		b.WriteString(fmt.Sprintf("%d. %s\n", i+1, x))
	}
	b.WriteString("\n")
	b.WriteString(a.Publication.NativeLawSpaceStatement + "\n\n")
	b.WriteString(a.Publication.QuarantinedAxiomStatement + "\n\n")
	b.WriteString(a.Publication.EnvironmentalBoundaryStatement + "\n\n")
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
