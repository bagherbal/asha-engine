package brokengaugefields

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FiniteBrokenGaugeFieldVariablesCurvatureSearchTheorem() theorem.Theorem {
	const id = "BRIDGE-FINITE-BROKEN-GAUGE-FIELD-CURVATURE-SEARCH"
	const name = "finite broken gauge-field variables / curvature term search"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 96 broken gauge-field audit", Passed: false, Detail: err.Error()}}}
		}
		vars := make([]string, 0, len(a.BrokenVariables))
		for _, v := range a.BrokenVariables {
			vars = append(vars, fmt.Sprintf("%s:%s", v.Name, v.Sector))
		}
		fullVars := make([]string, 0, len(a.FullVariables))
		for _, v := range a.FullVariables {
			fullVars = append(fullVars, fmt.Sprintf("%s:%s", v.Name, v.Sector))
		}
		br := make([]string, 0, len(a.Brackets))
		for _, b := range a.Brackets {
			keys := make([]string, 0, len(b.Components))
			for k := range b.Components {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			comps := make([]string, 0, len(keys))
			for _, k := range keys {
				comps = append(comps, fmt.Sprintf("%.3g%s", b.Components[k], k))
			}
			br = append(br, fmt.Sprintf("[%s,%s]=%s closedBroken=%t", b.Left, b.Right, strings.Join(comps, "+"), b.ClosedInBroken))
		}
		slotDetails := make([]string, 0, len(a.CurvatureSlots))
		for _, s := range a.CurvatureSlots {
			mark := "open"
			if s.Present {
				mark = "present"
			}
			slotDetails = append(slotDetails, fmt.Sprintf("%s=%s", s.Name, mark))
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 95 broken-sector input", Passed: len(a.CandidateHessian) == 3, Detail: fmt.Sprintf("K_candidate=[%.10f, %.10f, %.10f], det=%.10f", a.CandidateHessian[0], a.CandidateHessian[1], a.CandidateHessian[2], a.CandidateDet)},
			{Name: "broken gauge variables typed", Passed: a.BrokenFieldsTyped, Detail: fmt.Sprintf("dim=%d; %s", a.BrokenDimension, strings.Join(vars, "; "))},
			{Name: "full electroweak closure variables", Passed: a.FullEWFieldsRequired && a.FullEWClosed, Detail: fmt.Sprintf("dim=%d; %s", a.FullDimension, strings.Join(fullVars, "; "))},
			{Name: "broken-only closure audit", Passed: a.BrokenOnlyClosed, Detail: strings.Join(br, "; ")},
			{Name: "electromagnetic direction required", Passed: a.RequiresPhoton, Detail: "[T1,T2]=T3=(Z+Q)/2 in the broken basis; Q cannot be discarded from curvature"},
			{Name: "curvature/action slots audited", Passed: true, Detail: strings.Join(slotDetails, "; ")},
			{Name: "finite curvature term derived", Passed: a.CurvatureTermDerived, Detail: "not derived; field-strength/action still open"},
			{Name: "second variation possible", Passed: a.SecondVariationPossible, Detail: "not possible until finite field-strength action is derived"},
			{Name: "diag(1,1,4) selected by action", Passed: a.Diag114ActionSelected, Detail: "still only metric-whitening candidate"},
			{Name: "physical couplings or masses", Passed: a.PhysicalCouplings, Detail: "not derived; no claim of g2/gY/thetaW/alpha or W/Z masses is allowed"},
		}, Notes: []string{a.TruthStatement, "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
