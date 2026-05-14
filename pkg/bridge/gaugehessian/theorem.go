package gaugehessian

import (
	"fmt"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func GaugeKineticHessianSecondVariationTheorem() theorem.Theorem {
	const id = "BRIDGE-GAUGE-KINETIC-HESSIAN-SECOND-VARIATION"
	const name = "gauge kinetic Hessian from finite action second variation"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build gauge kinetic Hessian audit", Passed: false, Detail: err.Error()}}}
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 82 action-selection input", Passed: !a.GaugeAction.GaugeKineticActionSelected, Detail: "previous gate exposed positive kinetic diagnostics but selected no gauge action"},
			{Name: "U(1) Hessian search domain", Passed: a.SymmetricHessianDim == 6 && a.DiagonalNoMixingDim == 3 && a.OffDiagonalDim == 3, Detail: fmt.Sprintf("fields=%v; symmetric=%d, diagonal=%d, offDiagonal=%d", a.Fields, a.SymmetricHessianDim, a.DiagonalNoMixingDim, a.OffDiagonalDim)},
			{Name: "finite action slots audited", Passed: len(a.ActionSlots) >= 7, Detail: formatSlots(a.ActionSlots)},
			{Name: "existing derived slots", Passed: DerivedActionSlotCount(a.ActionSlots) == 3, Detail: fmt.Sprintf("derived=%d/7; representation metrics and anomaly/no-mixing data exist, but gauge-field action slots are open", DerivedActionSlotCount(a.ActionSlots))},
			{Name: "candidate Hessian families exposed", Passed: len(a.Candidates) >= 5, Detail: formatCandidates(a.Candidates)},
			{Name: "finite U(1) gauge-field variables", Passed: a.FiniteActionVariablesTyped, Detail: "not yet typed inside a finite gauge action"},
			{Name: "finite U(1) gauge-field action", Passed: a.U1GaugeFieldActionPresent, Detail: "not derived; no native U(1) curvature/kinetic action is available"},
			{Name: "second variation computed", Passed: a.SecondVariationComputed, Detail: "not computable until the finite action exists"},
			{Name: "gauge kinetic Hessian selected", Passed: a.HessianSelected, Detail: "not selected; diagnostic matrices remain families"},
			{Name: "boundary coupling fixed", Passed: a.BoundaryCouplingFixed, Detail: "not fixed; g_* remains open"},
			{Name: "physical alpha", Passed: a.PhysicalAlphaDerived, Detail: "not derived; alpha requires Hessian, boundary scale, and RG matching"},
			{Name: "hidden observed input", Passed: !a.HiddenObservedInputUsed, Detail: "no observed coupling, alpha, weak angle, or mass scale was inserted"},
		}, Notes: []string{a.TruthStatement, "Next: " + a.RecommendedNextGate}}
	}}
}

func formatSlots(xs []ActionSlot) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		state := "open"
		if x.Derived {
			state = "derived"
		} else if x.Present {
			state = "present"
		}
		parts = append(parts, fmt.Sprintf("%s[%s]: %s", x.Name, state, x.Detail))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}

func formatCandidates(xs []HessianCandidate) string {
	parts := make([]string, 0, len(xs))
	for _, x := range xs {
		state := "diagnostic"
		if x.FromSecondVariation {
			state = "derived"
		}
		parts = append(parts, fmt.Sprintf("%s[%s, free=%d]: %s", x.Name, state, x.FreeParameters, x.Detail))
	}
	return "[" + strings.Join(parts, "; ") + "]"
}
