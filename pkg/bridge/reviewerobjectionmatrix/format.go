package reviewerobjectionmatrix

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate422=%t nativeFlavorDim=%d conditionalFamilyDim=%d noFlavorReopening=%t verdict=%s", x.Gate422Ready, x.NativeFlavorDim, x.ConditionalFamilyDim, x.NoFlavorReopening, x.Verdict)
}

func FormatMatrix(x ObjectionMatrix) string {
	return fmt.Sprintf("rows=%d high=%d medium=%d low=%d flavor=%d axiom=%d cosmology=%d failedRoute=%d reproducibility=%d refs=%t boundaries=%t verdict=%s", len(x.Rows), x.HighRiskCount, x.MediumRiskCount, x.LowRiskCount, x.FlavorRows, x.AxiomRows, x.CosmologyRows, x.FailedRouteRows, x.ReproducibilityRows, x.AllRowsHaveReferences, x.AllRowsHaveBoundaries, x.Verdict)
}

func FormatGuide(x RebuttalGuide) string {
	return fmt.Sprintf("rules=%d required=%d forbidden=%d evidence=%d verdict=%s", len(x.Rules), len(x.RequiredPhrases), len(x.ForbiddenPhrases), len(x.EvidenceChecklist), x.Verdict)
}

func FormatExports(x ExportBundle) string {
	return fmt.Sprintf("matrix=%t guide=%t refs=%t risk=%t ready=%t verdict=%s", x.ObjectionMarkdown != "", x.RebuttalMarkdown != "", x.GateReferenceMarkdown != "", x.RiskMarkdown != "", x.PublicationReady, x.Verdict)
}

func FormatFinal(x FinalStatus) string {
	return fmt.Sprintf("matrixReady=%t boundariesReady=%t firewalls=%t noNewPhysics=%t noAxiomPromotion=%t nativeFlavorDim=%d conditionalFamilyDim=%d status=%s verdict=%s", x.MatrixReady, x.BoundariesReady, x.FirewallsPreserved, x.NoNewPhysicsClaim, x.NoAxiomPromotion, x.NativeFlavorDim, x.ConditionalFamilyDim, x.Status, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s", x.Gate, x.Title, x.Reason)
}

func renderObjectionMatrix(x ObjectionMatrix) string {
	var b strings.Builder
	b.WriteString("| ID | Class | Risk | Objection | Rebuttal | Boundary | Gates | Safe wording | Forbidden wording |\n")
	b.WriteString("|---|---|---|---|---|---|---|---|---|\n")
	for _, r := range x.Rows {
		b.WriteString(fmt.Sprintf("| `%s` | `%s` | `%s` | %s | %s | %s | %s | %s | %s |\n", r.ID, r.Class, r.Severity, esc(r.Objection), esc(r.Rebuttal), esc(r.Boundary), formatGates(r.GateReferences), esc(r.SafeWording), esc(r.UnsafeWording)))
	}
	return b.String()
}

func renderRebuttalGuide(x RebuttalGuide) string {
	var b strings.Builder
	b.WriteString("# Rebuttal Readiness Guide\n\n")
	b.WriteString(x.OpeningParagraph + "\n\n")
	b.WriteString("## Rules\n\n")
	for _, r := range x.Rules {
		b.WriteString("- " + r + "\n")
	}
	b.WriteString("\n## Required phrases\n\n")
	for _, r := range x.RequiredPhrases {
		b.WriteString("- `" + r + "`\n")
	}
	b.WriteString("\n## Forbidden phrases\n\n")
	for _, r := range x.ForbiddenPhrases {
		b.WriteString("- `" + r + "`\n")
	}
	b.WriteString("\n## Evidence checklist\n\n")
	for _, r := range x.EvidenceChecklist {
		b.WriteString("- " + r + "\n")
	}
	return b.String()
}

func renderGateReferenceMap(x ObjectionMatrix) string {
	refs := map[int][]string{}
	for _, r := range x.Rows {
		for _, g := range r.GateReferences {
			refs[g] = append(refs[g], r.ID)
		}
	}
	gates := sortedKeys(refs)
	var b strings.Builder
	b.WriteString("# Gate Reference Map\n\n")
	for _, g := range gates {
		b.WriteString(fmt.Sprintf("- Gate %d: %s\n", g, strings.Join(refs[g], ", ")))
	}
	return b.String()
}

func renderRiskSummary(x ObjectionMatrix) string {
	return fmt.Sprintf("# Reviewer Risk Summary\n\n- High-risk objections: %d\n- Medium-risk objections: %d\n- Low-risk objections: %d\n- Flavor rows: %d\n- Axiom rows: %d\n- Cosmology rows: %d\n- Failed-route rows: %d\n- Reproducibility rows: %d\n", x.HighRiskCount, x.MediumRiskCount, x.LowRiskCount, x.FlavorRows, x.AxiomRows, x.CosmologyRows, x.FailedRouteRows, x.ReproducibilityRows)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 423 Registry Audit — Reviewer Objection Matrix / Rebuttal Readiness Export\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Convert the Gate-422 executive claim audit into a reviewer-facing objection matrix with exact rebuttal boundaries and source-to-gate references. Gate 423 is an exposition/export gate and adds no new physics claim.\n\n")
	b.WriteString("## Gate 422 inheritance\n\n" + FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("## Objection matrix summary\n\n" + FormatMatrix(a.Matrix) + "\n\n")
	b.WriteString(a.Exports.ObjectionMarkdown + "\n")
	b.WriteString("## Rebuttal guide\n\n")
	b.WriteString("Summary: " + FormatGuide(a.Guide) + "\n\n")
	b.WriteString(a.Exports.RebuttalMarkdown + "\n")
	b.WriteString("## Gate reference map\n\n")
	b.WriteString(a.Exports.GateReferenceMarkdown + "\n")
	b.WriteString("## Risk summary\n\n")
	b.WriteString(a.Exports.RiskMarkdown + "\n")
	b.WriteString("## Result statuses\n\n")
	for _, s := range Statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Final status\n\n" + FormatFinal(a.Final) + "\n\n")
	b.WriteString("## Next gate\n\n" + FormatNext(a.Next) + "\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}

func formatGates(gs []int) string {
	parts := make([]string, len(gs))
	for i, g := range gs {
		parts[i] = fmt.Sprintf("G%d", g)
	}
	return strings.Join(parts, ", ")
}

func esc(s string) string { return strings.ReplaceAll(s, "|", "\\|") }

func sortedKeys(m map[int][]string) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	for i := 1; i < len(keys); i++ {
		for j := i; j > 0 && keys[j] < keys[j-1]; j-- {
			keys[j], keys[j-1] = keys[j-1], keys[j]
		}
	}
	return keys
}
