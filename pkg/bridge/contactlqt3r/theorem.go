package contactlqt3r

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactT3RChiralitySourceSearchLeptoquarkHyperchargeTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-T3R-CHIRALITY-SOURCE-LEPTOQUARK-HYPERCHARGE-SEARCH"
	const name = "contact T3R/chirality source search for leptoquark hypercharge"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build contact T3R/chirality source audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "matter T3R candidate family exists", Passed: a.MatterT3ROperatorFound && a.MatterChiralRestricted && a.MatterT3RDiagnosticAvailable && math.Abs(a.T3R.TemporalTrace) < 1e-10, Detail: fmt.Sprintf("T0 trace=%.3e trace2=%.10f chiral=%t mirror=%t", a.T3R.TemporalTrace, a.T3R.TemporalTraceSquared, a.MatterChiralRestricted, a.MatterMirrorAmbiguity)},
			{Name: "matter hyperaudit branch is diagnostic only", Passed: a.MatterHyperauditSelectsBranch && !a.MatterFullSMTableDerived && a.Hyperaudit.PreferredBranchName != "", Detail: fmt.Sprintf("preferred=%q fullSM=%t", a.Hyperaudit.PreferredBranchName, a.MatterFullSMTableDerived)},
			{Name: "B-L plus T3R produces only hypothetical contact Y values", Passed: math.Abs(a.Summary.HalfBMinusLDifference-2.0/3.0) < 1e-10 && a.Summary.HypotheticalYValueCount == 4 && a.HyperchargeRowsDerived == 0, Detail: FormatRows(a.Rows)},
			{Name: "no contact pullback, T3R, chirality, SU2, or signed B-L rows", Passed: a.ContactPullbackRowsDerived == 0 && a.ContactT3RRowsDerived == 0 && a.ContactChiralityRowsDerived == 0 && a.SignedBLRowsDerived == 0 && a.WeakSU2RowsDerived == 0, Detail: FormatSources(a.Sources)},
			{Name: "contact representation and beta permission remain closed", Passed: a.BetaPermissionFirewallClosed && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullBetaMatchingTensorDerived, Detail: fmt.Sprintf("rep=%d open=%d beta=%d zero=%d", a.RepresentationCompleteRows, a.RepresentationOpenRows, a.ContactBetaRowsAllowed, a.ContactZeroRowsProved)},
			{Name: "S6 ambiguity and physical-flow nullity are preserved", Passed: a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3, Detail: FormatSummary(a.Summary)},
			{Name: "no observed constants, scales, or masses leak in", Passed: !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: "no alpha/thetaW/Qobs/masses/M*/g* input"},
		}, Notes: []string{a.TruthStatement, "sources: " + FormatSources(a.Sources), "rows: " + FormatRows(a.Rows), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
