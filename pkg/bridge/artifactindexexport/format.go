package artifactindexexport

import (
	"fmt"
	"strings"
)

func FormatTree(x DocumentationTree) string {
	return fmt.Sprintf("rows=%d rootEntries=%d docsEntries=%d auditEntries=%d summaries=%d paper=%d visuals=%d code=%d rootClean=%t quickStart=%t artifactIndex=%t repro=%t verdict=%s", len(x.Rows), x.RootEntries, x.DocsEntries, x.AuditEntries, x.SummaryEntries, x.PaperEntries, x.VisualEntries, x.CodeEntries, x.RootIsClean, x.HasQuickStart, x.HasArtifactIndex, x.HasReproChecklist, x.Verdict)
}

func FormatCoverage(x AuditCoverage) string {
	return fmt.Sprintf("gateAudits=%d range=G%d-G%d missing=[%s] phenomenology=%d final=%d index=%s verdict=%s", x.GateAuditCount, x.FirstGate, x.LastGate, formatInts(x.KnownMissing), x.PhenomenologyCount, x.FinalReportCount, x.IndexPath, x.Verdict)
}

func FormatRepro(x ReproChecklist) string {
	return fmt.Sprintf("commands=%d targeted=%d avoided=%d policy=%d path=%s verdict=%s", len(x.Commands), x.TargetedCount, x.AvoidedCount, len(x.PolicyBullets), x.ChecklistPath, x.Verdict)
}

func FormatExports(x ExportBundle) string {
	return fmt.Sprintf("artifactIndex=%t reproducibility=%t maintenance=%t publicationWorkspace=%t ready=%t verdict=%s", x.ArtifactIndexMarkdown != "", x.ReproducibilityMarkdown != "", x.MaintenanceChecklistMarkdown != "", x.PublicationWorkspaceMarkdown != "", x.Ready, x.Verdict)
}

func FormatFinal(x FinalStatus) string {
	return fmt.Sprintf("artifactIndexReady=%t reproReady=%t rootClean=%t firewalls=%t noNewPhysics=%t noAxiomPromotion=%t nativeFlavorDim=%d conditionalFamilyDim=%d status=%s verdict=%s", x.ArtifactIndexReady, x.ReproChecklistReady, x.RootClean, x.FirewallsPreserved, x.NoNewPhysicsClaim, x.NoAxiomPromotion, x.NativeFlavorDim, x.ConditionalFamilyDim, x.Status, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s", x.Gate, x.Title, x.Reason)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 424 Registry Audit — Artifact Index / Reproducibility Checklist Export\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Compile the canonical artifact index and reproducibility checklist after the repository cleanup, without adding new physics claims or reopening any sealed frontier.\n\n")
	b.WriteString("## Gate 423 inheritance\n\n")
	b.WriteString(fmt.Sprintf("Gate %d reviewer matrix is treated as the immediate predecessor. Gate 424 only exports artifact navigation and reproducibility policy.\n\n", PriorReviewerGate))
	b.WriteString("## Documentation tree\n\n")
	b.WriteString(FormatTree(a.Tree) + "\n\n")
	b.WriteString("## Audit coverage\n\n")
	b.WriteString(FormatCoverage(a.Coverage) + "\n\n")
	b.WriteString("## Reproducibility checklist\n\n")
	b.WriteString(FormatRepro(a.Repro) + "\n\n")
	b.WriteString("## Export bundle\n\n")
	b.WriteString(FormatExports(a.Exports) + "\n\n")
	b.WriteString("## Artifact index preview\n\n")
	b.WriteString(a.Exports.ArtifactIndexMarkdown + "\n")
	b.WriteString("## Reproducibility checklist preview\n\n")
	b.WriteString(a.Exports.ReproducibilityMarkdown + "\n")
	b.WriteString("## Maintenance checklist\n\n")
	b.WriteString(a.Exports.MaintenanceChecklistMarkdown + "\n")
	b.WriteString("## Publication workspace guide\n\n")
	b.WriteString(a.Exports.PublicationWorkspaceMarkdown + "\n")
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
