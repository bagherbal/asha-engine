package manuscriptskeletonexport

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ManuscriptSkeletonSectionBySectionProofExportTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Manuscript skeleton / section-by-section proof export"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate421 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate420 theorem atlas", Passed: a.Inheritance.Executed && a.Inheritance.Gate420AtlasReady && a.Inheritance.AtlasGraphAcyclic, Detail: FormatInheritance(a.Inheritance)},
			{Name: "manuscript skeleton compiled", Passed: a.Manuscript.Executed && len(a.Manuscript.Sections) >= 10, Detail: FormatManuscript(a.Manuscript)},
			{Name: "section-by-section proof export ready", Passed: len(a.Manuscript.ProofObligations) >= len(a.Manuscript.Sections)*2, Detail: FormatManuscript(a.Manuscript)},
			{Name: "appendices compiled", Passed: len(a.Manuscript.Appendices) >= 4 && a.Exports.AppendixMarkdown != "", Detail: FormatExports(a.Exports)},
			{Name: "firewalls preserved in manuscript", Passed: a.Final.FirewallsPreserved && a.Final.NativeFlavorDim == NativeChargedFlavorDim, Detail: FormatFinal(a.Final)},
			{Name: "no new physics claim", Passed: a.Final.NoNewPhysicsClaim && a.Final.NoAxiomPromotion, Detail: FormatFinal(a.Final)},
			{Name: "manuscript skeleton ready", Passed: a.Final.SkeletonReady && a.Final.Status == StatusManuscriptSkeletonReady, Detail: FormatFinal(a.Final)},
			{Name: "next gate is executive abstract", Passed: a.Next.Gate == 422, Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth}}
	}}
}
