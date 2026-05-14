package publicationbundlepreflight

import (
	"fmt"
	"strings"
)

func FormatManifest(x PaperBundleManifest) string {
	return fmt.Sprintf("items=%d required=%d optional=%d ready=%d missing=%d path=%s verdict=%s", len(x.Items), x.RequiredCount, x.OptionalCount, x.ReadyCount, x.MissingCount, x.ManifestPath, x.Verdict)
}

func FormatSections(x SectionSourceMap) string {
	return fmt.Sprintf("sections=%d appendices=%d path=%s verdict=%s", x.SectionCount, x.AppendixCount, x.MapPath, x.Verdict)
}

func FormatFigures(x FigureLedger) string {
	return fmt.Sprintf("slots=%d ready=%d path=%s verdict=%s", x.SlotCount, x.ReadyCount, x.LedgerPath, x.Verdict)
}

func FormatFirewall(x FirewallChecklist) string {
	return fmt.Sprintf("rows=%d nativeFlavorDim=%d familyAxiomDim=%d path=%s verdict=%s", len(x.Rows), x.NativeFlavorDim, x.FamilyAxiomDim, x.ChecklistPath, x.Verdict)
}

func FormatExports(x ExportBundle) string {
	return fmt.Sprintf("preflight=%t manifest=%t sections=%t figures=%t firewall=%t assembly=%t ready=%t verdict=%s", x.PreflightMarkdown != "", x.ManifestMarkdown != "", x.SectionMapMarkdown != "", x.FigureLedgerMarkdown != "", x.FirewallMarkdown != "", x.AssemblyChecklistMD != "", x.Ready, x.Verdict)
}

func FormatFinal(x PreflightReport) string {
	return fmt.Sprintf("manifest=%t sections=%t figures=%t firewall=%t citation=%t rootPolicy=%t noNewPhysics=%t noAxiomPromotion=%t firewalls=%t ready=%t status=%s verdict=%s", x.ManifestReady, x.SectionMapReady, x.FigureLedgerReady, x.FirewallChecklistReady, x.CitationTemplatePreserved, x.RootCleanPolicyInherited, x.NoNewPhysicsClaim, x.NoAxiomPromotion, x.FirewallsPreserved, x.Ready, x.Status, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s", x.Gate, x.Title, x.Reason)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 425 Registry Audit — Final Paper Assembly / Publication Bundle Preflight\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Assemble the paper-facing publication bundle manifest, section source map, figure slot ledger, and claim-firewall checklist without adding new physics claims or reopening flavor/cosmology frontiers.\n\n")
	b.WriteString("## Gate 424 inheritance\n\n")
	b.WriteString(fmt.Sprintf("Gate %d artifact indexing is treated as the immediate predecessor. Gate 425 only prepares publication assembly surfaces.\n\n", PriorArtifactGate))
	b.WriteString("## Bundle manifest\n\n")
	b.WriteString(FormatManifest(a.Manifest) + "\n\n")
	b.WriteString("## Section source map\n\n")
	b.WriteString(FormatSections(a.Sections) + "\n\n")
	b.WriteString("## Figure slot ledger\n\n")
	b.WriteString(FormatFigures(a.Figures) + "\n\n")
	b.WriteString("## Firewall checklist\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString("## Export bundle\n\n")
	b.WriteString(FormatExports(a.Exports) + "\n\n")
	b.WriteString("## Publication bundle preflight\n\n")
	b.WriteString(a.Exports.PreflightMarkdown + "\n")
	b.WriteString("## Bundle manifest preview\n\n")
	b.WriteString(a.Exports.ManifestMarkdown + "\n")
	b.WriteString("## Section source map preview\n\n")
	b.WriteString(a.Exports.SectionMapMarkdown + "\n")
	b.WriteString("## Figure slot ledger preview\n\n")
	b.WriteString(a.Exports.FigureLedgerMarkdown + "\n")
	b.WriteString("## Claim firewall checklist preview\n\n")
	b.WriteString(a.Exports.FirewallMarkdown + "\n")
	b.WriteString("## Assembly checklist\n\n")
	b.WriteString(a.Exports.AssemblyChecklistMD + "\n")
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

func renderPreflight(m PaperBundleManifest, s SectionSourceMap, f FigureLedger, fw FirewallChecklist) string {
	var b strings.Builder
	b.WriteString("# Publication Bundle Preflight\n\n")
	b.WriteString("Gate 425 prepares the paper-facing bundle after the artifact index. It is an assembly and verification surface, not a new theorem gate.\n\n")
	b.WriteString("## Readiness summary\n\n")
	b.WriteString(fmt.Sprintf("- Bundle manifest: `%s`\n- Section source map: `%s`\n- Figure slot ledger: `%s`\n- Claim firewall checklist: `%s`\n- Required manifest items: `%d`\n- Missing required items: `%d`\n- Manuscript sections mapped: `%d`\n- Figure slots reserved: `%d`\n", m.ManifestPath, s.MapPath, f.LedgerPath, fw.ChecklistPath, m.RequiredCount, m.MissingCount, s.SectionCount, f.SlotCount))
	b.WriteString("\n## Boundary statement\n\n")
	b.WriteString("Native ASHA preserves `dim M_charged = 13`. The quarantined K/X/Y family axiom chain remains a nine-symbolic-coefficient capacity ledger. This preflight predicts no Yukawa value, CKM angle, CP phase, PMNS parameter, cosmology coordinate, or dark-sector observable.\n")
	return b.String()
}

func renderManifest(m PaperBundleManifest) string {
	var b strings.Builder
	b.WriteString("# Paper Bundle Manifest\n\n")
	b.WriteString("This manifest lists the files and workspaces needed to assemble the final paper without claim drift.\n\n")
	b.WriteString("| Path | Kind | Required | Source | Purpose | Claim rule | Readiness |\n")
	b.WriteString("|---|---|---:|---|---|---|---|\n")
	for _, item := range m.Items {
		b.WriteString(fmt.Sprintf("| `%s` | `%s` | `%t` | %s | %s | %s | %s |\n", item.Path, item.Kind, item.Required, esc(item.Source), esc(item.Purpose), esc(item.ClaimRule), esc(item.Readiness)))
	}
	return b.String()
}

func renderSectionMap(s SectionSourceMap) string {
	var b strings.Builder
	b.WriteString("# Section Source Map\n\n")
	b.WriteString("Each manuscript section must trace to gate outputs and include its boundary language.\n\n")
	b.WriteString("| Section | Source | Core claim | Proof input | Boundary | Status |\n")
	b.WriteString("|---|---|---|---|---|---|\n")
	for _, row := range s.Sections {
		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n", esc(row.Section), esc(row.Source), esc(row.CoreClaim), esc(row.ProofInput), esc(row.Boundary), esc(row.Status)))
	}
	b.WriteString("\n## Required appendices\n\n")
	b.WriteString("1. Theorem atlas / dependency graph.\n2. Failed-route and firewall ledger.\n3. Reproducibility and targeted-test policy.\n4. Figure and artifact manifest.\n")
	return b.String()
}

func renderFigureLedger(f FigureLedger) string {
	var b strings.Builder
	b.WriteString("# Figure Slot Ledger\n\n")
	b.WriteString("Figure files may be added later. This ledger reserves publication-safe slots and source expectations.\n\n")
	b.WriteString("| Figure | Target path | Source | Purpose | Status |\n")
	b.WriteString("|---|---|---|---|---|\n")
	for _, slot := range f.Slots {
		b.WriteString(fmt.Sprintf("| %s | `%s` | %s | %s | %s |\n", esc(slot.Name), slot.TargetPath, esc(slot.Source), esc(slot.Purpose), esc(slot.Status)))
	}
	return b.String()
}

func renderFirewall(fw FirewallChecklist) string {
	var b strings.Builder
	b.WriteString("# Claim Firewall Checklist\n\n")
	b.WriteString("Every paper draft should be checked against this table before moving from `docs/paper/drafts/` to `docs/paper/final/`.\n\n")
	b.WriteString(fmt.Sprintf("Native charged flavor dimension: `%d`\n\nConditional K/X/Y family source dimension: `%d` symbolic charged coefficients\n\n", fw.NativeFlavorDim, fw.FamilyAxiomDim))
	b.WriteString("| Topic | Allowed wording | Forbidden wording | Gate reference |\n")
	b.WriteString("|---|---|---|---|\n")
	for _, row := range fw.Rows {
		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s |\n", esc(row.Topic), esc(row.Allowed), esc(row.Forbidden), esc(row.GateReference)))
	}
	return b.String()
}

func renderAssemblyChecklist() string {
	return `# Paper Assembly Checklist

1. Start from the section source map in ` + "`docs/paper/SECTION_SOURCE_MAP.md`" + `.
2. Draft or place manuscript files under ` + "`docs/paper/drafts/`" + `.
3. For every section, cite the corresponding gate audit or summary source.
4. Add figures only through the slots in ` + "`docs/paper/FIGURE_SLOT_LEDGER.md`" + `.
5. Check every claim against ` + "`docs/paper/CLAIM_FIREWALL_CHECKLIST.md`" + `.
6. Preserve the README published-paper citation template until publication metadata is known.
7. Run targeted publication-support tests before finalizing:

` + "```bash" + `
go test -p=1 ./pkg/bridge/publicationbundlepreflight -count=1
go test -p=1 ./pkg/bridge/publicationbundlepreflight ./pkg/bridge/artifactindexexport ./pkg/bridge/reviewerobjectionmatrix ./pkg/bridge/executiveabstractclaimaudit ./pkg/bridge/manuscriptskeletonexport ./pkg/bridge/publicationtheorematlas -count=1
go list ./internal/app
` + "```" + `

Do not run ` + "`go test ./...`" + ` by default when timeout risk matters.
`
}
