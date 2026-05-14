package executiveabstractclaimaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ExecutiveAbstractClaimAuditSummaryExportTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Executive abstract / claim-audit summary export"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate422 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate421 manuscript skeleton", Passed: a.Inheritance.Executed && a.Inheritance.Gate421Ready, Detail: FormatInheritance(a.Inheritance)},
			{Name: "executive abstract compiled", Passed: a.Abstract.Executed && len(a.Abstract.NativeClaims) >= 5 && len(a.Abstract.Firewalls) >= 3, Detail: FormatAbstract(a.Abstract)},
			{Name: "claim-audit summary compiled", Passed: a.ClaimAudit.Executed && len(a.ClaimAudit.Rows) >= 10 && a.ClaimAudit.FirewallsExplicit, Detail: FormatClaimAudit(a.ClaimAudit)},
			{Name: "firewall language compiled", Passed: a.Exports.FirewallMarkdown != "" && a.Final.FirewallsPreserved, Detail: FormatExports(a.Exports)},
			{Name: "reviewer-safe warnings compiled", Passed: len(a.Abstract.ReviewerWarnings) >= 4 && a.Exports.ReviewerMarkdown != "", Detail: FormatAbstract(a.Abstract)},
			{Name: "explicit non-claim ledger compiled", Passed: len(a.Abstract.NonClaims) >= 5 && a.ClaimAudit.NonClaimCount >= 1, Detail: FormatClaimAudit(a.ClaimAudit)},
			{Name: "no new physics claim", Passed: a.Final.NoNewPhysicsClaim && a.Final.NoAxiomPromotion, Detail: FormatFinal(a.Final)},
			{Name: "executive claim-audit ready", Passed: a.Final.ExecutiveReady && a.Final.Status == StatusExecutiveSummaryReady, Detail: FormatFinal(a.Final)},
			{Name: "next gate is reviewer objection matrix", Passed: a.Next.Gate == 423, Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth}}
	}}
}
