package reviewerobjectionmatrix

import "github.com/bagherbal/asha-engine/pkg/theorem"

func ReviewerObjectionMatrixRebuttalReadinessExportTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Reviewer objection matrix / rebuttal readiness export"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate423 audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherits Gate422 claim audit", Passed: a.Inheritance.Executed && a.Inheritance.Gate422Ready, Detail: FormatInheritance(a.Inheritance)},
			{Name: "reviewer objection matrix compiled", Passed: a.Matrix.Executed && len(a.Matrix.Rows) >= 10, Detail: FormatMatrix(a.Matrix)},
			{Name: "rebuttal boundaries compiled", Passed: a.Matrix.AllRowsHaveBoundaries && a.Final.BoundariesReady, Detail: FormatMatrix(a.Matrix)},
			{Name: "gate reference map compiled", Passed: a.Matrix.AllRowsHaveReferences && a.Exports.GateReferenceMarkdown != "", Detail: FormatExports(a.Exports)},
			{Name: "risk ranking compiled", Passed: a.Matrix.HighRiskCount >= 3 && a.Exports.RiskMarkdown != "", Detail: FormatMatrix(a.Matrix)},
			{Name: "claim wording audited", Passed: len(a.Guide.RequiredPhrases) >= 5 && len(a.Guide.ForbiddenPhrases) >= 5, Detail: FormatGuide(a.Guide)},
			{Name: "no new physics claim", Passed: a.Final.NoNewPhysicsClaim && a.Final.NoAxiomPromotion, Detail: FormatFinal(a.Final)},
			{Name: "reviewer matrix ready", Passed: a.Final.MatrixReady && a.Final.Status == StatusReviewerMatrixReady, Detail: FormatFinal(a.Final)},
			{Name: "next gate is artifact index", Passed: a.Next.Gate == 424, Detail: FormatNext(a.Next)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: []string{a.Truth}}
	}}
}
