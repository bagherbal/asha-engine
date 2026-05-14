package publicationtheorematlas

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("G419=%t G418FlavorSeal=%t nativeFlavorDim=%d conditionalFamilyDim=%d noFlavorReopening=%t verdict=%s", x.Gate419BoardReady, x.Gate418FlavorSealInherited, x.NativeFlavorDim, x.ConditionalFamilyDim, x.NoFlavorReopening, x.Verdict)
}

func FormatAtlas(x Atlas) string {
	return fmt.Sprintf("nodes=%d edges=%d native=%d bridge=%d quarantined=%d environmental=%d failed=%d acyclic=%t verdict=%s", len(x.Nodes), len(x.Edges), x.NativeCount, x.BridgeCount, x.QuarantinedCount, x.EnvironmentalCount, x.FailedRouteCount, x.Acyclic, x.Verdict)
}

func FormatNode(n AtlasNode) string {
	return fmt.Sprintf("%s [%s] gates=%s package=%s status=%s deps=%s claim=%s boundary=%s", n.ID, n.Layer, joinInts(n.Gates), n.Package, n.Status, strings.Join(n.DependsOn, ","), n.Claim, n.Boundary)
}

func FormatExports(x ExportBundle) string {
	return fmt.Sprintf("mermaid=%t dot=%t markdown=%t machineLedger=%d ready=%t verdict=%s", x.HasMermaid, x.HasDOT, x.HasMarkdown, len(x.MachineLedger), x.PublicationReady, x.Verdict)
}

func FormatFirewalls(x FirewallLedger) string {
	return fmt.Sprintf("firewalls=%d flavor=%t cosmology=%t noEmpirical=%t verdict=%s", len(x.Firewalls), x.FlavorFirewallPreserved, x.CosmologyFirewallPreserved, x.NoEmpiricalDataInserted, x.Verdict)
}

func FormatFailedIndex(x FailedRouteIndex) string {
	return fmt.Sprintf("routes=%d scalar=%d fermion=%d familyAxiom=%d indexed=%t verdict=%s", len(x.Routes), x.ScalarRoutes, x.FermionRoutes, x.FamilyAxiomRoutes, x.Indexed, x.Verdict)
}

func FormatFinal(x FinalStatus) string {
	return fmt.Sprintf("atlasReady=%t acyclic=%t firewalls=%t noNewPhysics=%t noAxiomPromotion=%t nativeFlavorDim=%d conditionalFamilyDim=%d status=%s verdict=%s", x.AtlasReady, x.GraphAcyclic, x.FirewallsPreserved, x.NoNewPhysicsClaim, x.NoAxiomPromotion, x.NativeFlavorDim, x.ConditionalFamilyDim, x.Status, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s", x.Gate, x.Title, x.Reason)
}

func renderNodeTable(a Atlas) string {
	var b strings.Builder
	b.WriteString("| ID | Layer | Gates | Package | Status | Dependencies | Claim | Boundary |\n")
	b.WriteString("|---|---|---|---|---|---|---|---|\n")
	for _, n := range a.Nodes {
		b.WriteString(fmt.Sprintf("| `%s` | %s | %s | `%s` | %s | %s | %s | %s |\n", n.ID, n.Layer, joinInts(n.Gates), n.Package, n.Status, strings.Join(n.DependsOn, ", "), n.Claim, n.Boundary))
	}
	return b.String()
}

func renderMachineLedger(a Atlas) []string {
	rows := make([]string, 0, len(a.Nodes))
	for _, n := range a.Nodes {
		rows = append(rows, fmt.Sprintf("%s|%s|%s|%s|%s", n.ID, n.Layer, joinInts(n.Gates), n.Status, strings.Join(n.DependsOn, ",")))
	}
	return rows
}

func renderMermaid(a Atlas) string {
	var b strings.Builder
	b.WriteString("graph TD\n")
	for _, n := range a.Nodes {
		label := strings.ReplaceAll(n.ID, "-", "_")
		b.WriteString(fmt.Sprintf("  %s[\"%s\\n%s\"]\n", label, n.ID, n.Layer))
	}
	for _, e := range a.Edges {
		from := strings.ReplaceAll(e.From, "-", "_")
		to := strings.ReplaceAll(e.To, "-", "_")
		b.WriteString(fmt.Sprintf("  %s --> %s\n", from, to))
	}
	return b.String()
}

func renderDOT(a Atlas) string {
	var b strings.Builder
	b.WriteString("digraph ASHA_Gate420_Atlas {\n")
	b.WriteString("  rankdir=LR;\n")
	for _, n := range a.Nodes {
		b.WriteString(fmt.Sprintf("  \"%s\" [label=\"%s\\n%s\"];\n", n.ID, n.ID, n.Layer))
	}
	for _, e := range a.Edges {
		b.WriteString(fmt.Sprintf("  \"%s\" -> \"%s\";\n", e.From, e.To))
	}
	b.WriteString("}\n")
	return b.String()
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 420 Registry Audit — Publication-Grade Theorem Atlas / Dependency Graph Export\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Export the Gate-419 post-flavor ASHA law-space board into a peer-reviewable theorem atlas and dependency graph. Gate 420 is an artifact/export gate: it adds no physics claim, predicts no flavor coefficient, and promotes no quarantined axiom.\n\n")
	b.WriteString("## Gate 419 inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("## Atlas summary\n\n")
	b.WriteString(FormatAtlas(a.Atlas) + "\n\n")
	b.WriteString(renderNodeTable(a.Atlas) + "\n")
	b.WriteString("## Dependency graph exports\n\n")
	b.WriteString(FormatExports(a.Exports) + "\n\n")
	b.WriteString("### Mermaid\n\n```mermaid\n")
	b.WriteString(a.Exports.Mermaid)
	b.WriteString("```\n\n")
	b.WriteString("### DOT\n\n```dot\n")
	b.WriteString(a.Exports.DOT)
	b.WriteString("```\n\n")
	b.WriteString("## Topological order\n\n")
	for i, id := range a.Atlas.TopologicalOrder {
		b.WriteString(fmt.Sprintf("%d. `%s`\n", i+1, id))
	}
	b.WriteString("\n## Firewall ledger\n\n")
	b.WriteString(FormatFirewalls(a.Firewalls) + "\n\n")
	b.WriteString("| Firewall | Native dim | Conditional dim | Status | Preserved | Coordinates | Claim |\n")
	b.WriteString("|---|---:|---:|---|---:|---|---|\n")
	for _, f := range a.Firewalls.Firewalls {
		b.WriteString(fmt.Sprintf("| %s | %d | %d | `%s` | %t | %s | %s |\n", f.Name, f.NativeDimension, f.ConditionalDimension, f.Status, f.Preserved, strings.Join(f.Coordinates, ", "), f.Claim))
	}
	b.WriteString("\n## Failed-route index\n\n")
	b.WriteString(FormatFailedIndex(a.FailedIndex) + "\n\n")
	b.WriteString("| Gate range | Route | Reason | Lesson |\n")
	b.WriteString("|---|---|---|---|\n")
	for _, r := range a.FailedIndex.Routes {
		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n", r.GateRange, r.Route, r.Reason, r.Lesson))
	}
	b.WriteString("\n## Result statuses\n\n")
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
