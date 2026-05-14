package brokenaction

import (
	"fmt"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func BrokenSectorActionSecondVariationSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-BROKEN-SECTOR-ACTION-SECOND-VARIATION"
	const name = "broken-sector action second variation / kinetic Hessian search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 95 broken-sector action audit", Passed: false, Detail: err.Error()}}}
		}
		slotDetails := make([]string, 0, len(a.ActionSlots))
		for _, s := range a.ActionSlots {
			mark := "open"
			if s.Present {
				mark = "present"
			}
			slotDetails = append(slotDetails, fmt.Sprintf("%s=%s", s.Name, mark))
		}
		candDetails := make([]string, 0, len(a.Candidates))
		for _, c := range a.Candidates {
			diag := "free"
			if len(c.Diagonal) == 3 {
				diag = fmt.Sprintf("[%.10f, %.10f, %.10f]", c.Diagonal[0], c.Diagonal[1], c.Diagonal[2])
			}
			candDetails = append(candDetails, fmt.Sprintf("%s diag=%s selected=%t", c.Name, diag, c.Selected))
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 94 diag(1,1,4) input", Passed: a.CandidatePositive && a.WhiteningExact, Detail: fmt.Sprintf("K=[%.10f, %.10f, %.10f], det=%.10f, whiteningExact=%t", a.CandidateK[0], a.CandidateK[1], a.CandidateK[2], a.CandidateDet, a.WhiteningExact)},
			{Name: "finite action slots audited", Passed: true, Detail: strings.Join(slotDetails, "; ")},
			{Name: "candidate second-variation families", Passed: true, Detail: strings.Join(candDetails, "; ")},
			{Name: "covariant-derivative diagnostic exists", Passed: a.CovariantDerivativeDerived, Detail: "DΦ template and broken-generator image diagnostics exist from Gates 84-85"},
			{Name: "finite scalar kinetic action", Passed: a.ScalarKineticActionDerived, Detail: "not derived; active-frame metric remains diagnostic"},
			{Name: "finite gauge field variables", Passed: a.GaugeFieldVariablesDerived, Detail: "not derived as action variables for broken directions"},
			{Name: "finite curvature/field-strength term", Passed: a.CurvatureTermDerived, Detail: "not derived; no action term can yet be varied to produce K_broken"},
			{Name: "second variation computed", Passed: a.SecondVariationComputed, Detail: "open; δ²S/δA_iδA_j has not been computed from a finite action"},
			{Name: "diag(1,1,4) selected by action", Passed: a.Diag114SelectedByAction, Detail: "not selected; currently selected only by metric-whitening diagnostic"},
			{Name: "physical couplings or masses", Passed: a.CouplingRatioDerived || a.PhysicalMassesDerived, Detail: "not derived; no claim of g2/gY/thetaW/alpha or physical W/Z masses is allowed"},
		}, Notes: []string{a.TruthStatement, "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
