package structuralfamilyboardexport

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate448=%t K=%t G2Zero=%t X=%t YQuarantined=%t coeffQuarantined=%t nativeDim=%d KXY=%d noEmpirical=%t supportOnly=%t verdict=%s", x.Gate448Reconciled, x.KGenPromoted, x.Gen2ZeroPromoted, x.XSupportPromoted, x.YPhaseQuarantined, x.CoefficientsQuarantined, x.NativeFlavorDim, x.KXYCoeffDim, x.NoEmpiricalInputsImported, x.PublicationSupportOnly, x.Verdict)
}

func FormatBoard(x StructuralBoard) string {
	return fmt.Sprintf("rows=%d promoted=%d quarantined=%d nativeDim=%d KXY=%d K=%t G2Zero=%t X=%t YQuarantined=%t coeffQuarantined=%t noObservable=%t verdict=%s reason=%s", len(x.Rows), x.PromotedRows, x.QuarantinedRows, x.NativeFlavorDim, x.KXYCoeffDim, x.KGenRowPresent, x.Gen2ZeroRowPresent, x.XTriangleRowPresent, x.YPhaseRowQuarantined, x.CoeffRowsQuarantined, x.NoObservablePredicted, x.Verdict, x.Reason)
}

func FormatDelta(x ManuscriptDelta) string {
	return fmt.Sprintf("blocks=%d abstract=%t section9=%t conclusion=%t reviewer=%t appendix=%t noClaimDrift=%t noBinaryMutation=%t target=%s verdict=%s reason=%s", len(x.Blocks), x.AbstractInsertionReady, x.Section9ReplacementReady, x.ConclusionAddendumReady, x.ReviewerNoteReady, x.AppendixDeltaReady, x.NoClaimDrift, x.NoFinalDocumentMutation, x.RecommendedTargetPath, x.Verdict, x.Reason)
}

func FormatArtifacts(x FigureTableDelta) string {
	return fmt.Sprintf("tables=%d figures=%d required=%d ready=%d verdict=%s reason=%s", len(x.Tables), len(x.Figures), x.RequiredCount, x.ReadyCount, x.Verdict, x.Reason)
}

func FormatFirewall(x FirewallAddendum) string {
	return fmt.Sprintf("rows=%d nativeDim=%d KXY=%d allowK=%t allowX=%t forbidYukawa=%t forbidMixing=%t forbidMass=%t forbidCoeff=%t forbidCosmo=%t verdict=%s reason=%s", len(x.Rows), x.NativeFlavorDim, x.KXYCoeffDim, x.AllowsKGenPromotion, x.AllowsXSupportPromotion, x.ForbidsYukawaPrediction, x.ForbidsMixingPrediction, x.ForbidsMassPrediction, x.ForbidsCoefficientFit, x.ForbidsCosmologyUpdate, x.Verdict, x.Reason)
}

func FormatReviewer(x ReviewerPacket) string {
	return fmt.Sprintf("objections=%d ready=%d noClaimDrift=%t firewall=%t verdict=%s reason=%s", len(x.Objections), x.ReadyCount, x.NoClaimDrift, x.FirewallStated, x.Verdict, x.Reason)
}

func FormatExports(x ExportBundle) string {
	return fmt.Sprintf("target=%s board=%t delta=%t artifacts=%t firewall=%t reviewer=%t combined=%t ready=%t noNewPhysics=%t verdict=%s reason=%s", x.TargetPath, x.StructuralBoardMarkdown != "", x.ManuscriptDeltaMarkdown != "", x.FigureTableMarkdown != "", x.FirewallAddendumMarkdown != "", x.ReviewerPacketMarkdown != "", x.CombinedMarkdown != "", x.PublicationReady, x.NoNewPhysicsClaim, x.Verdict, x.Reason)
}

func FormatFinal(x FinalStatus) string {
	return fmt.Sprintf("ready=%t board=%t delta=%t firewall=%t reviewer=%t noNewPhysics=%t noMass=%t noYukawa=%t noCKM=%t noPMNS=%t nativeDim=%d KXY=%d status=%s verdict=%s", x.Ready, x.BoardReady, x.ManuscriptDeltaReady, x.FirewallReady, x.ReviewerReady, x.NoNewPhysicsClaim, x.NoObservedMassImported, x.NoYukawaImported, x.NoCKMImported, x.NoPMNSImported, x.NativeFlavorDim, x.KXYCoeffDim, x.Status, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s Task=%s", x.Gate, x.Title, x.Reason, x.PrimaryTask)
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 449 Registry Audit — Structural Family Board Export / Manuscript Delta Patch\n\n")
	b.WriteString("## Scope\n\n")
	b.WriteString("Gate 449 is a publication-support export. It converts the Gate-448 post-444 flavor reconciliation into guarded manuscript language, a structural family board, figure/table deltas, and reviewer-safe firewall wording. It is not a new physics derivation.\n\n")

	b.WriteString("## Gate 448 inheritance\n\n")
	b.WriteString(FormatInheritance(a.Inheritance) + "\n\n")

	b.WriteString("## Structural family board\n\n")
	b.WriteString(FormatBoard(a.Board) + "\n\n")
	b.WriteString(renderBoardTable(a.Board.Rows) + "\n")

	b.WriteString("## Manuscript delta\n\n")
	b.WriteString(FormatDelta(a.Delta) + "\n\n")
	b.WriteString("| Target section | Operation | Class | Source gates | Boundary | Ready | Patch text |\n")
	b.WriteString("|---|---|---|---|---|---:|---|\n")
	for _, p := range a.Delta.Blocks {
		b.WriteString(fmt.Sprintf("| %s | %s | `%s` | %s | %s | %t | %s |\n", esc(p.TargetSection), esc(p.Operation), p.ClaimClass, joinInts(p.SourceGates), esc(p.Boundary), p.Ready, esc(p.Markdown)))
	}
	b.WriteString("\n")

	b.WriteString("## Figure/table delta\n\n")
	b.WriteString(FormatArtifacts(a.Artifacts) + "\n\n")
	b.WriteString(renderArtifacts(a.Artifacts) + "\n")

	b.WriteString("## Claim firewall addendum\n\n")
	b.WriteString(FormatFirewall(a.Firewall) + "\n\n")
	b.WriteString(renderFirewallTable(a.Firewall.Rows) + "\n")

	b.WriteString("## Reviewer packet\n\n")
	b.WriteString(FormatReviewer(a.Reviewer) + "\n\n")
	b.WriteString(renderReviewerTable(a.Reviewer.Objections) + "\n")

	b.WriteString("## Export bundle\n\n")
	b.WriteString(FormatExports(a.Exports) + "\n\n")
	b.WriteString("Recommended target path: `" + a.Exports.TargetPath + "`\n\n")

	b.WriteString("## Export preview\n\n")
	b.WriteString(a.Exports.CombinedMarkdown + "\n")

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

func renderCombined(board StructuralBoard, delta ManuscriptDelta, artifacts FigureTableDelta, firewall FirewallAddendum, reviewer ReviewerPacket) string {
	var b strings.Builder
	b.WriteString("# Post-444 Manuscript Delta Patch\n\n")
	b.WriteString("This paper-facing patch updates the Gate-420 manuscript bundle after Gates 444--448. It promotes only structural family-board statements and preserves every value-bearing flavor/cosmology firewall.\n\n")
	b.WriteString("## Structural family board\n\n")
	b.WriteString(renderBoardTable(board.Rows) + "\n")
	b.WriteString("## Manuscript insertion blocks\n\n")
	for _, p := range delta.Blocks {
		b.WriteString(fmt.Sprintf("### %s\n\n", p.TargetSection))
		b.WriteString(fmt.Sprintf("- Operation: %s\n- Source gates: %s\n- Boundary: %s\n\n", p.Operation, joinInts(p.SourceGates), p.Boundary))
		b.WriteString(p.Markdown + "\n\n")
	}
	b.WriteString("## Figure/table delta\n\n")
	b.WriteString(renderArtifacts(artifacts) + "\n")
	b.WriteString("## Claim firewall addendum\n\n")
	b.WriteString(renderFirewallTable(firewall.Rows) + "\n")
	b.WriteString("## Reviewer packet\n\n")
	b.WriteString(renderReviewerTable(reviewer.Objections) + "\n")
	b.WriteString("## Non-claim boundary\n\n")
	b.WriteString("This patch predicts no observed muon/charm mass, Yukawa value, CKM angle, CKM phase, PMNS parameter, bridge amplitude, sector coefficient, dark-matter abundance, cosmological constant, or cosmological history.\n")
	return b.String()
}

func renderBoard(board StructuralBoard) string { return renderBoardTable(board.Rows) }
func renderDelta(delta ManuscriptDelta) string {
	var b strings.Builder
	for _, p := range delta.Blocks {
		b.WriteString(fmt.Sprintf("## %s\n\n%s\n\n", p.TargetSection, p.Markdown))
	}
	return b.String()
}
func renderArtifacts(artifacts FigureTableDelta) string {
	var b strings.Builder
	b.WriteString("| Kind | Name | Target path | Source | Purpose | Claim rule | Status |\n")
	b.WriteString("|---|---|---|---|---|---|---|\n")
	for _, x := range append(artifacts.Tables, artifacts.Figures...) {
		b.WriteString(fmt.Sprintf("| `%s` | %s | `%s` | %s | %s | %s | %s |\n", x.Kind, esc(x.Name), x.TargetPath, esc(x.Source), esc(x.Purpose), esc(x.ClaimRule), esc(x.Status)))
	}
	return b.String()
}
func renderFirewall(f FirewallAddendum) string { return renderFirewallTable(f.Rows) }
func renderReviewer(r ReviewerPacket) string   { return renderReviewerTable(r.Objections) }

func renderBoardTable(rows []BoardRow) string {
	var b strings.Builder
	b.WriteString("| Object | Layer | Formula | Gate | Claim | Boundary | Paper action | Status |\n")
	b.WriteString("|---|---|---|---:|---|---|---|---|\n")
	for _, r := range rows {
		b.WriteString(fmt.Sprintf("| %s | %s | `%s` | %d | %s | %s | %s | `%s` |\n", esc(r.Object), esc(r.Layer), esc(r.Formula), r.SourceGate, esc(r.Claim), esc(r.Boundary), esc(r.PaperAction), r.Status))
	}
	return b.String()
}

func renderFirewallTable(rows []FirewallRow) string {
	var b strings.Builder
	b.WriteString("| Topic | Allowed wording | Forbidden wording | Source | Status |\n")
	b.WriteString("|---|---|---|---|---|\n")
	for _, r := range rows {
		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | `%s` |\n", esc(r.Topic), esc(r.Allowed), esc(r.Forbidden), esc(r.Source), r.Status))
	}
	return b.String()
}

func renderReviewerTable(rows []ReviewerObjection) string {
	var b strings.Builder
	b.WriteString("| Reviewer concern | Response | Boundary | Status |\n")
	b.WriteString("|---|---|---|---|\n")
	for _, r := range rows {
		b.WriteString(fmt.Sprintf("| %s | %s | %s | `%s` |\n", esc(r.Objection), esc(r.Answer), esc(r.Boundary), r.Status))
	}
	return b.String()
}

func esc(s string) string {
	s = strings.ReplaceAll(s, "|", "\\|")
	s = strings.ReplaceAll(s, "\n", "<br>")
	return s
}
