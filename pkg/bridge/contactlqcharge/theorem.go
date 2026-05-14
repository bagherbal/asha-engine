package contactlqcharge

import (
	"fmt"
	"math"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func LeptoquarkContactHyperchargeSourceBLChargeLatticeObstructionTheorem() theorem.Theorem {
	const id = "BRIDGE-LEPTOQUARK-CONTACT-HYPERCHARGE-SOURCE-BL-CHARGE-LATTICE-OBSTRUCTION"
	const name = "leptoquark contact hypercharge source / B-L and charge-lattice obstruction theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build B-L charge-lattice obstruction audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "B-L charge bridge is valid but matter-side", Passed: a.BMinusLChargeBridgeValid && a.BMinusLPolarizesOnePlusThree && math.Abs(a.Summary.BMinusLOneParticleTrace) < 1e-10 && math.Abs(a.Summary.BMinusLOneParticleTrace2-4.0/3.0) < 1e-10, Detail: FormatSummary(a.Summary)},
			{Name: "lepton-color B-L difference is only a current diagnostic", Passed: a.BLDifferenceDiagnostic && math.Abs(a.LeptonColorBLDifference-4.0/3.0) < 1e-10 && a.Summary.BLDiagnosticRows == 6 && a.SignedBLRowsDerived == 0, Detail: FormatRows(a.Rows)},
			{Name: "T3R, weak chirality, SU(2)L, and hypercharge rows are absent", Passed: a.T3RRowsDerived == 0 && a.WeakChiralityRowsDerived == 0 && a.WeakSU2RowsDerived == 0 && a.HyperchargeRowsDerived == 0 && a.ElectricChargeRowsDerived == 0, Detail: FormatSources(a.Sources)},
			{Name: "contact representation and beta permission remain closed", Passed: a.BetaPermissionFirewallClosed && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullBetaMatchingTensorDerived, Detail: fmt.Sprintf("rep=%d open=%d beta=%d zero=%d", a.RepresentationCompleteRows, a.RepresentationOpenRows, a.ContactBetaRowsAllowed, a.ContactZeroRowsProved)},
			{Name: "S6 ambiguity and physical-flow nullity are preserved", Passed: a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3, Detail: fmt.Sprintf("s6=%d nullity=%d->%d", a.ResidualS6Choices, a.ResidualNullityBefore, a.ResidualNullityAfter)},
			{Name: "no observed charges, constants, scales, or masses leak in", Passed: !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: "no alpha/thetaW/Qobs/masses/M*/g* input"},
		}, Notes: []string{a.TruthStatement, "sources: " + FormatSources(a.Sources), "rows: " + FormatRows(a.Rows), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
