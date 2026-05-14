package ewcurvature

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func FullElectroweakConnectionCurvatureAuditTheorem() theorem.Theorem {
	const id = "BRIDGE-FULL-ELECTROWEAK-CONNECTION-CURVATURE-AUDIT"
	const name = "full electroweak connection curvature / field-strength audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate 97 full EW curvature audit", Passed: false, Detail: err.Error()}}}
		}
		vars := make([]string, 0, len(a.Variables))
		for _, v := range a.Variables {
			vars = append(vars, fmt.Sprintf("%s:%s", v.Name, v.Basis))
		}
		br := make([]string, 0, len(a.Brackets))
		for _, b := range a.Brackets {
			keys := make([]string, 0, len(b.Components))
			for k := range b.Components {
				keys = append(keys, k)
			}
			sort.Strings(keys)
			parts := make([]string, 0, len(keys))
			for _, k := range keys {
				parts = append(parts, fmt.Sprintf("%.3g%s", b.Components[k], k))
			}
			if len(parts) == 0 {
				parts = append(parts, "0")
			}
			br = append(br, fmt.Sprintf("[%s,%s]=%s", b.Left, b.Right, strings.Join(parts, "+")))
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: []theorem.Check{
			{Name: "Gate 96 full-connection input", Passed: a.Gate96.FullEWFieldsRequired && a.Gate96.RequiresPhoton, Detail: "Gate 96 required Q because [T1,T2]=(Z+Q)/2"},
			{Name: "full electroweak variables typed", Passed: a.CurvatureCarrierTyped && a.Dimension == 4, Detail: fmt.Sprintf("dim=%d; %s", a.Dimension, strings.Join(vars, "; "))},
			{Name: "Lie closure in full basis", Passed: a.Closed && a.ClosureResidual < 1e-12, Detail: strings.Join(br, "; ")},
			{Name: "field-strength carrier typed", Passed: a.FullFieldStrengthTyped, Detail: "formal F=dA+A∧A can be written only on the closed {T1,T2,Z,Q} carrier"},
			{Name: "adjoint/Killing diagnostic", Passed: a.AdjointRank == 3, Detail: fmt.Sprintf("rank=%d trace=%.10f matrix=%s", a.AdjointRank, a.AdjointMetricTrace, FormatMatrix(a.AdjointMetric))},
			{Name: "abelian null direction exposed", Passed: len(a.AbelianNullVector) == 4, Detail: "null vector [0,0,-1,1] = Q-Z = 2Y_phi; adjoint curvature does not normalize U(1)"},
			{Name: "semisimple neutral direction exposed", Passed: len(a.SemisimpleDirection) == 4, Detail: "T3=(Z+Q)/2 is the neutral SU(2) direction seen by adjoint curvature"},
			{Name: "positive gauge Hessian selected", Passed: a.AdjointMetricPositive, Detail: "not selected; adjoint diagnostic is degenerate because the abelian direction is null"},
			{Name: "diag(1,1,4) selected by curvature", Passed: a.Diag114SelectedByCurvature, Detail: "not selected; diag(1,1,4) remains metric-whitening candidate from broken-image geometry"},
			{Name: "second variation computed", Passed: a.SecondVariationComputed, Detail: "not computed; finite quadratic action still missing"},
			{Name: "physical couplings or masses", Passed: a.PhysicalCouplings, Detail: "not derived; no g2/gY/thetaW/alpha or W/Z masses are claimed"},
		}, Notes: []string{a.TruthStatement, "rejected claims: " + Join(a.RejectedClaims), "remaining unknowns: " + Join(a.RemainingUnknowns), "Next: " + a.RecommendedNextGate}}
	}}
}
