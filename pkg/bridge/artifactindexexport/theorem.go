package artifactindexexport

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ArtifactIndexReproducibilityChecklistExportTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Artifact index / reproducibility checklist export"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate424 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate423 reviewer matrix boundary", Passed: PriorReviewerGate == 423, Detail: "prior gate = 423"},
			{Name: "artifact index compiled", Passed: a.Exports.ArtifactIndexMarkdown != "" && a.Tree.HasArtifactIndex, Detail: FormatExports(a.Exports)},
			{Name: "reproducibility checklist compiled", Passed: a.Exports.ReproducibilityMarkdown != "" && a.Repro.TargetedCount >= 3, Detail: FormatRepro(a.Repro)},
			{Name: "document tree indexed", Passed: a.Tree.Executed && len(a.Tree.Rows) >= 15, Detail: FormatTree(a.Tree)},
			{Name: "audit coverage indexed", Passed: a.Coverage.Executed && a.Coverage.GateAuditCount >= 227 && a.Coverage.LastGate == 424, Detail: FormatCoverage(a.Coverage)},
			{Name: "summaries/paper/visuals indexed", Passed: a.Tree.SummaryEntries >= 1 && a.Tree.PaperEntries >= 1 && a.Tree.VisualEntries >= 1, Detail: FormatTree(a.Tree)},
			{Name: "root cleanliness audited", Passed: a.Tree.RootIsClean && a.Final.RootClean, Detail: FormatFinal(a.Final)},
			{Name: "no new physics claim", Passed: a.Final.NoNewPhysicsClaim && a.Final.NoAxiomPromotion && a.Final.FirewallsPreserved, Detail: FormatFinal(a.Final)},
			{Name: "artifact index ready", Passed: a.Final.ArtifactIndexReady && a.Final.Status == StatusArtifactIndexReady, Detail: FormatFinal(a.Final)},
			{Name: "next gate is publication bundle preflight", Passed: a.Next.Gate == 425, Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth}}
	}}
}
