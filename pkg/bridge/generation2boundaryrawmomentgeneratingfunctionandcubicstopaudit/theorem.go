package generation2boundaryrawmomentgeneratingfunctionandcubicstopaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2BoundaryRawMomentGeneratingFunctionAndCubicStopAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 783 — Boundary Raw-Moment Generating Function and Cubic Stop Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate782 boundary-flavor complement", Passed: a.Gate782.Inherited && strings.Contains(a.Gate782.SelectedSubBottleneck, "F_wall_3_red"), Detail: a.Gate782.FormulaLevelIndependence},
			{Name: "audit projector power degeneracy", Passed: a.Operator.Audited && a.Operator.ProjectorIdempotent && strings.Contains(a.Operator.PowerLaw, "s^n P_7") && a.Operator.AllMomentsSameSupport && !a.Operator.IndependentOperatorGeometry, Detail: FormatOperator(a.Operator)},
			{Name: "reduce raw moments to scalar response coordinates", Passed: a.Operator.ScalarRawMomentCoordinate && containsAll(a.Operator.Moments, []string{"M1=p s", "M2=p s^2", "M3=p s^3", "M_n=p s^n"}), Detail: strings.Join(a.Operator.Moments, "; ")},
			{Name: "define response function representation", Passed: a.Response.Audited && a.Response.WrittenAsExpectation && strings.Contains(a.Response.Function, "x+kappa_e_red") && a.Response.MatchesLedger && !a.Response.NativeGeneratingFunction, Detail: FormatResponse(a.Response)},
			{Name: "rewrite cubic coefficient as 2p", Passed: a.Coefficients.Audited && closeRel(a.Coefficients.TwoP, 7.0/36.0, 1e-15) && strings.Contains(a.Coefficients.CubicCoefficient, "-2p") && strings.Contains(a.Coefficients.BoundaryPairCandidate, "7/36") && !a.Coefficients.NativeBoundaryPairTheorem && !a.Coefficients.TypedOperatorMapSourcesCubic, Detail: FormatCoefficients(a.Coefficients)},
			{Name: "audit cubic stop candidates", Passed: a.CubicStop.Audited && !a.CubicStop.ProjectorAlgebraStopsPowers && a.CubicStop.BoundaryPairExteriorCandidate && strings.Contains(a.CubicStop.RequiredExteriorMap, "Lambda^3 B_boundary=0") && !a.CubicStop.RawMomentToExteriorMapCertified && !a.CubicStop.NativeCubicStopTheorem, Detail: FormatCubicStop(a.CubicStop)},
			{Name: "reject untyped M4 fit", Passed: a.CubicStop.UntypedM4FitRejected && !a.CubicStop.TypedFourthOrderCoefficient, Detail: FormatCubicStop(a.CubicStop)},
			{Name: "preserve raw-moment coordinate", Passed: a.Coordinate.Audited && strings.Contains(a.Coordinate.RawMomentForm, "p s^n") && strings.Contains(a.Coordinate.ActiveCoordinate, "raw moments") && !a.Coordinate.VarianceActive && !a.Coordinate.CentralMomentActive && !a.Coordinate.NativeRawMomentCoordinate, Detail: a.Coordinate.ActiveCoordinate},
			{Name: "audit F_wall formula-level runtime target absence", Passed: a.Runtime.Audited && a.Runtime.FormulaLevelRuntimeTargetAbsence && a.Runtime.EvaluableWithoutDirectHiggsRuntimeVariables && !a.Runtime.UsesLambdaRuntime && !a.Runtime.UsesLambdaRuntimeEff && !a.Runtime.UsesTreeMass && !a.Runtime.UsesPoleMass && !a.Runtime.UsesCHiggs && !a.Runtime.UsesGF && !a.Runtime.UsesVEV, Detail: FormatRuntime(a.Runtime)},
			{Name: "record relation to kappa_lambda_red and C_History", Passed: a.Relation.Recorded && strings.Contains(a.Relation.KappaLambdaRelation, "F_wall_3_red") && strings.Contains(a.Relation.FWallLevel, "Level B") && strings.Contains(a.Relation.KappaLambdaLevel, "Level B") && strings.Contains(a.Relation.CHistoryLevel, "Level B") && strings.Contains(a.Relation.CHiggsLevel, "not Level C") && !a.Relation.CHistoryIndependent, Detail: FormatRelation(a.Relation)},
			{Name: "record prediction-level classification", Passed: a.Prediction.Recorded && strings.Contains(a.Prediction.FWall3Level, "Level B") && strings.Contains(a.Prediction.KappaLambdaLevel, "Level B") && strings.Contains(a.Prediction.CHistoryLevel, "Level B") && strings.Contains(a.Prediction.CHiggsLevel, "not Level C"), Detail: a.Prediction.FWall3Level + "; " + a.Prediction.CHiggsLevel},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.ResponseFunctionNative && !a.Firewalls.BoundaryGeneratingFunctionNative && !a.Firewalls.RawMomentCoordinateNative && !a.Firewalls.BoundaryPairStressPullNative && !a.Firewalls.RawMomentExteriorDegreeMapNative && !a.Firewalls.CubicStopNative && !a.Firewalls.TypedFourthOrderCoefficientSource && !a.Firewalls.FWallNativeBoundaryResponse && !a.Firewalls.CHistoryFullIndependentPrediction && !a.Firewalls.TreeProxyPoleMass && !a.Firewalls.YukawaNativeTheorem && a.Firewalls.Verdict == StatusFirewallPreservedGate783, Detail: a.Firewalls.Verdict},
		}
		ok := true
		for _, c := range checks {
			if !c.Passed {
				ok = false
				break
			}
		}
		status := theorem.BridgeRequired
		if !ok {
			status = theorem.FailedRoute
		}
		notes := append([]string{a.Truth, a.FinalStatement}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
