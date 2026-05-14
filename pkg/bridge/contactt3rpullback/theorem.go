package contactt3rpullback

import (
	"fmt"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func ContactT3RPullbackFockToContactIntertwinerSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-CONTACT-T3R-PULLBACK-FOCK-CONTACT-INTERTWINER-SEARCH"
	const name = "contact T3R pullback obstruction / Fock-to-contact intertwiner search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Fock-to-contact intertwiner search", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.Variational, Checks: []theorem.Check{
			{Name: "matter-side T3R/chirality diagnostics are inherited", Passed: a.MatterT3ROperatorFound && a.MatterChiralRestricted && a.MatterMirrorAmbiguity && a.MatterDimension == 16 && a.ScalarDimension == 4 && a.TensorDimension == 64, Detail: FormatSummary(a.Summary)},
			{Name: "generic Fock/tensor maps exist but are not canonical intertwiners", Passed: a.GenericFockToContactMapsExist && a.GenericTensorToContactMapsExist && a.Summary.FockToContactGenericKernelDim == 9 && a.Summary.TensorToContactGenericKernelDim == 57 && a.CanonicalFockToContactMaps == 0 && a.FockToContactIntertwinersDerived == 0, Detail: FormatCandidates(a.Candidates)},
			{Name: "no T3R, chirality, B-L, SU2, or hypercharge pullback rows", Passed: a.T3RPullbackRowsDerived == 0 && a.ChiralityPullbackRowsDerived == 0 && a.BMinusLPullbackRowsDerived == 0 && a.SU2LPullbackRowsDerived == 0 && a.HyperchargeRowsDerived == 0 && a.ElectricChargeRowsDerived == 0, Detail: FormatRows(a.Rows)},
			{Name: "contact representation and beta permission remain closed", Passed: a.BetaPermissionFirewallClosed && a.RepresentationCompleteRows == 0 && a.RepresentationOpenRows == 7 && a.ContactBetaRowsAllowed == 0 && a.ContactZeroRowsProved == 0 && !a.ThresholdCorrectedBetaDerived && !a.FullBetaMatchingTensorDerived, Detail: fmt.Sprintf("rep=%d open=%d beta=%d zero=%d", a.RepresentationCompleteRows, a.RepresentationOpenRows, a.ContactBetaRowsAllowed, a.ContactZeroRowsProved)},
			{Name: "S6 ambiguity and physical-flow nullity are preserved", Passed: a.ResidualS6Choices == 720 && a.ResidualNullityBefore == 3 && a.ResidualNullityAfter == 3, Detail: fmt.Sprintf("s6=%d nullity=%d->%d", a.ResidualS6Choices, a.ResidualNullityBefore, a.ResidualNullityAfter)},
			{Name: "no observed constants, scales, or masses leak in", Passed: !a.HiddenObservedInputUsed && !a.PhysicalWeakAngleDerived && !a.FineStructureDerived && !a.PhysicalMassesDerived && !a.PhysicalScaleDerived, Detail: "no alpha/thetaW/Qobs/masses/M*/g* input"},
		}, Notes: []string{a.TruthStatement, "candidates: " + FormatCandidates(a.Candidates), "rows: " + FormatRows(a.Rows), "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
