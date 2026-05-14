package contactlqbetapermission

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func LeptoquarkHyperchargeLocalFieldBetaPermissionTheorem() theorem.Theorem {
	const id = "BRIDGE-LEPTOQUARK-HYPERCHARGE-LOCAL-FIELD-BETA-PERMISSION"
	const name = "leptoquark hypercharge-row and local-field obstruction / beta-permission theorem"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build leptoquark beta-permission audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "Gate 133 SU(2)L obstruction inherited", Passed: a.Gate133SU2ObstructionInherited && a.LeptoquarkRows == 6 && a.CurrentLQSlots == 6 && a.ColorDirections == 3 && a.RealOrientations == 2, Detail: FormatSummary(a.Summary)},
			{Name: "hypercharge and local-field rows are absent", Passed: !a.HyperchargeRowDerived && !a.LocalFieldMapDerived && a.Summary.HyperchargeRowsDerived == 0 && a.Summary.LocalFieldRowsDerived == 0, Detail: FormatRequirements(a.Requirements)},
			{Name: "Lorentz kinetic, mass activation, and decoupling remain underived", Passed: !a.LorentzKineticRowDerived && !a.PoleResidueTheoremDerived && !a.MassActivationDerived && !a.DecouplingRuleDerived && a.Summary.LorentzKineticRowsDerived == 0 && a.Summary.MassActivationRowsDerived == 0 && a.Summary.DecouplingRowsDerived == 0, Detail: FormatRows(a.CandidateRows)},
			{Name: "leptoquark contact beta permission remains closed", Passed: a.BetaPermissionFirewallClosed && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullBetaMatchingTensorDerived, Detail: fmt.Sprintf("rep=%d open=%d beta=%d zero=%d", a.RepresentationCompleteRows, a.RepresentationOpenRows, a.ContactBetaRowsAllowed, a.ContactZeroRowsProved)},
			{Name: "S6 ambiguity and physical-flow nullity are preserved", Passed: a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3, Detail: fmt.Sprintf("s6=%d nullity=%d->%d", a.ResidualS6Choices, a.ResidualNullityBefore, a.ResidualNullityAfter)},
			{Name: "no observed constants or masses leak into the selector", Passed: !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: "no alpha/thetaW/masses/M*/g* input"},
		}, Notes: []string{a.TruthStatement, "requirements: " + FormatRequirements(a.Requirements), "candidate rows: " + FormatRows(a.CandidateRows), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
