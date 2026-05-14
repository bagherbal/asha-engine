package executiveabstractclaimaudit

import (
	"fmt"
	"strings"
)

func FormatInheritance(x Inheritance) string {
	return fmt.Sprintf("gate421=%t sections=%d proofs=%d nativeFlavorDim=%d conditionalFamilyDim=%d noFlavorReopening=%t verdict=%s", x.Gate421Ready, x.SectionCount, x.ProofObligationCount, x.NativeFlavorDim, x.ConditionalFamilyDim, x.NoFlavorReopening, x.Verdict)
}

func FormatAbstract(x ExecutiveAbstract) string {
	return fmt.Sprintf("nativeClaims=%d conditionalClaims=%d firewalls=%d nonClaims=%d warnings=%d nativeFlavorDim=%d conditionalFamilyDim=%d noNewPhysics=%t verdict=%s", len(x.NativeClaims), len(x.ConditionalClaims), len(x.Firewalls), len(x.NonClaims), len(x.ReviewerWarnings), x.NativeFlavorDim, x.ConditionalFamilyDim, x.NoNewPhysicsClaim, x.Verdict)
}

func FormatClaimAudit(x ClaimAudit) string {
	return fmt.Sprintf("rows=%d native=%d bridge=%d conditional=%d firewall=%d failed=%d nonClaim=%d firewallsExplicit=%t noAxiomPromotion=%t verdict=%s", len(x.Rows), x.NativeCount, x.BridgeCount, x.ConditionalCount, x.FirewallCount, x.FailedRouteCount, x.NonClaimCount, x.FirewallsExplicit, x.NoAxiomPromotion, x.Verdict)
}

func FormatExports(x ExportBundle) string {
	return fmt.Sprintf("executive=%t claimAudit=%t firewalls=%t reviewerWarnings=%t frontMatter=%d ready=%t verdict=%s", x.ExecutiveMarkdown != "", x.ClaimAuditMarkdown != "", x.FirewallMarkdown != "", x.ReviewerMarkdown != "", len(x.RecommendedFrontMatter), x.PublicationReady, x.Verdict)
}

func FormatFinal(x FinalStatus) string {
	return fmt.Sprintf("executiveReady=%t claimAuditReady=%t firewalls=%t noNewPhysics=%t noAxiomPromotion=%t nativeFlavorDim=%d conditionalFamilyDim=%d status=%s verdict=%s", x.ExecutiveReady, x.ClaimAuditReady, x.FirewallsPreserved, x.NoNewPhysicsClaim, x.NoAxiomPromotion, x.NativeFlavorDim, x.ConditionalFamilyDim, x.Status, x.Verdict)
}

func FormatNext(x NextStep) string {
	return fmt.Sprintf("Gate %d — %s: %s", x.Gate, x.Title, x.Reason)
}

func renderExecutive(x ExecutiveAbstract) string {
	var b strings.Builder
	b.WriteString("# Executive Abstract\n\n")
	b.WriteString("**Title:** " + x.Title + "\n\n")
	b.WriteString("**One-sentence claim:** " + x.OneSentence + "\n\n")
	b.WriteString(x.ShortAbstract + "\n\n")
	b.WriteString("## Native claims\n\n")
	for _, c := range x.NativeClaims {
		b.WriteString("- " + c + "\n")
	}
	b.WriteString("\n## Conditional/quarantined claims\n\n")
	for _, c := range x.ConditionalClaims {
		b.WriteString("- " + c + "\n")
	}
	b.WriteString("\n## Firewalls\n\n")
	for _, c := range x.Firewalls {
		b.WriteString("- " + c + "\n")
	}
	b.WriteString("\n## Explicit non-claims\n\n")
	for _, c := range x.NonClaims {
		b.WriteString("- " + c + "\n")
	}
	return b.String()
}

func renderClaimAudit(x ClaimAudit) string {
	var b strings.Builder
	b.WriteString("| ID | Class | Safe verb | Claim | Evidence | Boundary |\n")
	b.WriteString("|---|---|---|---|---|---|\n")
	for _, r := range x.Rows {
		b.WriteString(fmt.Sprintf("| `%s` | `%s` | `%s` | %s | %s | %s |\n", r.ID, r.Class, r.SafeVerb, r.Claim, r.Evidence, r.Boundary))
	}
	return b.String()
}

func renderFirewalls(x ExecutiveAbstract) string {
	var b strings.Builder
	b.WriteString("# Firewall Language\n\n")
	b.WriteString("Use the following exact-style boundary language in abstracts, introductions, and replies to reviewers.\n\n")
	b.WriteString("## Flavor firewall\n\n")
	b.WriteString(fmt.Sprintf("Native ASHA preserves the charged flavor frontier at dimension %d. The quarantined K/X/Y family axiom chain supplies hierarchy, mixing, and CP capacity with %d symbolic charged coefficients, but their values are environmental boundary coordinates rather than native ASHA predictions.\n\n", x.NativeFlavorDim, x.ConditionalFamilyDim))
	b.WriteString("## Cosmology firewall\n\n")
	b.WriteString("Cosmological observables such as dark matter abundance, cosmological constant, and universe age are treated as environmental/history-sensitive coordinates unless a future gate derives a native dynamical bridge.\n\n")
	b.WriteString("## Axiom quarantine\n\n")
	b.WriteString("K_gen, X_gen, and Y_gen are explicit family-capacity axioms. They are not promoted to native finite-geometry theorems.\n")
	return b.String()
}

func renderReviewerWarnings(x ExecutiveAbstract) string {
	var b strings.Builder
	b.WriteString("# Reviewer-Safe Warnings\n\n")
	for _, w := range x.ReviewerWarnings {
		b.WriteString("- " + w + "\n")
	}
	return b.String()
}

func RenderAudit(a Analysis) string {
	var b strings.Builder
	b.WriteString("# Gate 422 Registry Audit — Executive Abstract / Claim-Audit Summary Export\n\n")
	b.WriteString("## Claim tested\n\n")
	b.WriteString("Convert the Gate-421 manuscript skeleton into concise executive front matter: abstract language, claim-audit table, firewall language, reviewer-safe warnings, and explicit non-claims. Gate 422 is an exposition/export gate and adds no new physics claim.\n\n")
	b.WriteString("## Gate 421 inheritance\n\n" + FormatInheritance(a.Inheritance) + "\n\n")
	b.WriteString("## Executive abstract summary\n\n" + FormatAbstract(a.Abstract) + "\n\n")
	b.WriteString(a.Exports.ExecutiveMarkdown + "\n")
	b.WriteString("## Claim-audit table\n\n")
	b.WriteString("Summary: " + FormatClaimAudit(a.ClaimAudit) + "\n\n")
	b.WriteString(a.Exports.ClaimAuditMarkdown + "\n")
	b.WriteString("## Firewall language\n\n")
	b.WriteString(a.Exports.FirewallMarkdown + "\n")
	b.WriteString("## Reviewer warnings\n\n")
	b.WriteString(a.Exports.ReviewerMarkdown + "\n")
	b.WriteString("## Recommended front matter checklist\n\n")
	for _, item := range a.Exports.RecommendedFrontMatter {
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
