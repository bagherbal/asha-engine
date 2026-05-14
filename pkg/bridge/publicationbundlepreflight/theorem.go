package publicationbundlepreflight

import "github.com/bagherbal/asha-engine/pkg/theorem"

func FinalPaperAssemblyPublicationBundlePreflightTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Final paper assembly / publication bundle preflight"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate425 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate424 artifact index", Passed: PriorArtifactGate == 424, Detail: "prior gate = 424"},
			{Name: "publication bundle preflight compiled", Passed: a.Exports.PreflightMarkdown != "" && a.Final.Ready, Detail: FormatFinal(a.Final)},
			{Name: "paper manifest compiled", Passed: a.Manifest.RequiredCount >= 12 && a.Manifest.MissingCount == 0, Detail: FormatManifest(a.Manifest)},
			{Name: "section source map compiled", Passed: a.Sections.SectionCount == ManuscriptSectionCount && a.Sections.AppendixCount >= RequiredAppendixCount, Detail: FormatSections(a.Sections)},
			{Name: "figure slot ledger compiled", Passed: a.Figures.SlotCount >= 6 && a.Figures.ReadyCount == a.Figures.SlotCount, Detail: FormatFigures(a.Figures)},
			{Name: "firewall checklist compiled", Passed: len(a.Firewall.Rows) >= 8 && a.Firewall.NativeFlavorDim == NativeChargedFlavorDim, Detail: FormatFirewall(a.Firewall)},
			{Name: "citation template preserved", Passed: a.Final.CitationTemplatePreserved, Detail: FormatFinal(a.Final)},
			{Name: "publication readiness audited", Passed: a.Final.Ready && a.Final.Status == StatusBundlePreflightReady, Detail: FormatFinal(a.Final)},
			{Name: "no new physics claim", Passed: a.Final.NoNewPhysicsClaim && a.Final.NoAxiomPromotion && a.Final.FirewallsPreserved, Detail: FormatFinal(a.Final)},
			{Name: "next gate is paper draft integration", Passed: a.Next.Gate == 426, Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth}}
	}}
}
