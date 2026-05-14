package manuscriptskeletonexport

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("Gate420=%t acyclic=%t nativeFlavorDim=%d conditionalFamilyDim=%d noFlavorReopening=%t verdict=%s", x.Gate420AtlasReady, x.AtlasGraphAcyclic, x.NativeFlavorDim, x.ConditionalFamilyDim, x.NoFlavorReopening, x.Verdict)
}

func FormatManuscript(x Manuscript) string {
	return fmt.Sprintf("sections=%d native=%d bridge=%d boundary=%d appendix=%d proofs=%d appendices=%d flavorFirewall=%t cosmologyFirewall=%t noNewPhysics=%t verdict=%s", len(x.Sections), x.NativeSections, x.BridgeSections, x.BoundarySections, x.AppendixSections, len(x.ProofObligations), len(x.Appendices), x.HasFlavorFirewall, x.HasCosmologyFirewall, x.NoNewPhysicsClaim, x.Verdict)
}

func FormatExports(x ExportBundle) string {
	return fmt.Sprintf("outline=%t proofMatrix=%t appendices=%t checklist=%d ready=%t verdict=%s", x.OutlineMarkdown != "", x.ProofMatrixMarkdown != "", x.AppendixMarkdown != "", len(x.ArtifactChecklist), x.PublicationReady, x.Verdict)
}

func FormatFinal(x FinalStatus) string {
	return fmt.Sprintf("skeletonReady=%t proofExport=%t firewalls=%t noNewPhysics=%t noAxiomPromotion=%t nativeFlavorDim=%d conditionalFamilyDim=%d status=%s verdict=%s", x.SkeletonReady, x.ProofExportReady, x.FirewallsPreserved, x.NoNewPhysicsClaim, x.NoAxiomPromotion, x.NativeFlavorDim, x.ConditionalFamilyDim, x.Status, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s", x.Gate, x.Title, x.Reason)
}

func renderOutline(m Manuscript) string {
	var b strings.Builder
	b.WriteString("# Manuscript Skeleton\n\n")
	b.WriteString("**Title:** " + m.Title + "\n\n")
	b.WriteString("## Abstract\n\n" + m.Abstract + "\n\n")
	for i, s := range m.Sections {
		b.WriteString(fmt.Sprintf("## %d. %s\n\n", i, s.Title))
		b.WriteString("- Kind: `" + string(s.Kind) + "`\n")
		b.WriteString("- Gates: " + joinInts(s.Gates) + "\n")
		b.WriteString("- Claim: " + s.Claim + "\n")
		b.WriteString("- Proof task: " + s.ProofTask + "\n")
		b.WriteString("- Boundary: " + s.Boundary + "\n")
		if len(s.DependsOn) > 0 {
			b.WriteString("- Depends on: `" + strings.Join(s.DependsOn, "`, `") + "`\n")
		}
		b.WriteString("\n")
	}
	return b.String()
}

func renderProofMatrix(m Manuscript) string {
	var b strings.Builder
	b.WriteString("| ID | Section | Type | Statement | Evidence | Status |\n")
	b.WriteString("|---|---|---|---|---|---|\n")
	for _, p := range m.ProofObligations {
		b.WriteString(fmt.Sprintf("| `%s` | `%s` | %s | %s | %s | `%s` |\n", p.ID, p.SectionID, p.Type, p.Statement, p.Evidence, p.Status))
	}
	return b.String()
}

func renderAppendices(m Manuscript) string {
	var b strings.Builder
	b.WriteString("| ID | Title | Purpose | Inputs | Status |\n")
	b.WriteString("|---|---|---|---|---|\n")
	for _, a := range m.Appendices {
		b.WriteString(fmt.Sprintf("| `%s` | %s | %s | %s | `%s` |\n", a.ID, a.Title, a.Purpose, strings.Join(a.Inputs, ", "), a.Status))
	}
	return b.String()
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 421 Registry Audit — Manuscript Skeleton / Section-by-Section Proof Export\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Convert the Gate-420 publication theorem atlas into a manuscript skeleton with section-level proof obligations, appendices, firewalls, and artifact checklist. Gate 421 is an exposition/export gate and adds no new physics claim.\n\n")
	b.WriteString("## Gate 420 inheritance\n\n" + FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("## Manuscript summary\n\n" + FormatManuscript(a.Manuscript) + "\n\n")
	b.WriteString("## Manuscript outline\n\n")
	b.WriteString(a.Exports.OutlineMarkdown + "\n")
	b.WriteString("## Proof obligation matrix\n\n")
	b.WriteString(a.Exports.ProofMatrixMarkdown + "\n")
	b.WriteString("## Appendices\n\n")
	b.WriteString(a.Exports.AppendixMarkdown + "\n")
	b.WriteString("## Artifact checklist\n\n")
	for _, item := range a.Exports.ArtifactChecklist {
		b.WriteString("- " + item + "\n")
	}
	b.WriteString("\n## Result statuses\n\n")
	for _, s := range Statuses() {
		b.WriteString("- `" + s + "`\n")
	}
	b.WriteString("\n## Final status\n\n" + FormatFinal(a.Final) + "\n\n")
	b.WriteString("## Next gate\n\n" + FormatNext(a.Next) + "\n\n")
	b.WriteString("## Truth statement\n\n" + a.Truth + "\n")
	return b.String()
}
