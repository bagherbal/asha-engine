package publicationtheorematlas

import "github.com/bagherbal/asha-engine/pkg/theorem"

func PublicationGradeTheoremAtlasDependencyGraphExportTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Publication-grade theorem atlas / dependency graph export"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate420 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate419 final law-space board", Passed: a.Inheritance.Executed && a.Inheritance.Gate419BoardReady && a.Inheritance.NoFlavorReopening, Detail: FormatInheritance(a.Inheritance)},
			{Name: "publication theorem atlas compiled", Passed: a.Atlas.Executed && len(a.Atlas.Nodes) >= 20, Detail: FormatAtlas(a.Atlas)},
			{Name: "dependency graph exported", Passed: a.Exports.Executed && a.Exports.HasMermaid && a.Exports.HasDOT && a.Exports.HasMarkdown, Detail: FormatExports(a.Exports)},
			{Name: "atlas graph is acyclic", Passed: a.Atlas.Acyclic && len(a.Atlas.TopologicalOrder) == len(a.Atlas.Nodes), Detail: FormatAtlas(a.Atlas)},
			{Name: "layer classification preserved", Passed: a.Atlas.NativeCount >= 6 && a.Atlas.BridgeCount >= 5 && a.Atlas.QuarantinedCount >= 3 && a.Atlas.EnvironmentalCount >= 2, Detail: FormatAtlas(a.Atlas)},
			{Name: "failed routes indexed", Passed: a.FailedIndex.Executed && a.FailedIndex.Indexed && len(a.FailedIndex.Routes) >= 5, Detail: FormatFailedIndex(a.FailedIndex)},
			{Name: "firewalls exported", Passed: a.Firewalls.Executed && a.Firewalls.FlavorFirewallPreserved && a.Firewalls.CosmologyFirewallPreserved && a.Firewalls.NoEmpiricalDataInserted, Detail: FormatFirewalls(a.Firewalls)},
			{Name: "no new physics claim", Passed: a.Final.NoNewPhysicsClaim && a.Final.NoAxiomPromotion && a.Final.NativeFlavorDim == NativeChargedFlavorDim, Detail: FormatFinal(a.Final)},
			{Name: "publication atlas ready", Passed: a.Final.AtlasReady && a.Final.Status == StatusPublicationAtlasReady, Detail: FormatFinal(a.Final)},
			{Name: "next gate is manuscript skeleton", Passed: a.Next.Gate == 421, Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth}}
	}}
}
