package generation2completedsquarehiggspotentialandvacuumoffsetaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2CompletedSquareHiggsPotentialAndVacuumOffsetAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 772 — Completed-Square Higgs Potential and Vacuum-Offset Firewall Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusGate772CompletedSquareBoundary}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 771 VEV, mu-squared, and offset ledger", Passed: a.Gate771.Inherited && a.Gate771.QuarticAirlock == "HiggsQuarticRuntimeCoefficientSeal" && a.Gate771.VEVSeal == "VEVConventionSeal" && a.Gate771.LambdaHIdentification == "lambda_H := lambda_runtime_eff" && a.Gate771.MuSquaredFormula == "mu^2_bridge=-lambda_runtime_eff v^2" && closeRel(a.Gate771.LambdaRuntimeEff, lambdaRuntimeEff, 1e-15) && closeRel(a.Gate771.VEVGeV, vevConventionGeV, 1e-15) && closeRel(a.Gate771.MuSquaredBridgeGeV2, -7860.072200382293, 1e-15) && closeRel(a.Gate771.C0LocalBridgeGeV4, 119127483.0758411, 1e-14) && closeRel(a.Gate771.TreeProxyGeV, 125.38000000304908, 1e-15) && !a.Gate771.NativeVEVTheorem && !a.Gate771.NativeMuSquaredTheorem && !a.Gate771.C0CosmologyTheorem, Detail: FormatGate771(a.Gate771)},
			{Name: "derive completed-square form", Passed: a.Square.StartingPotential == "V(u)=c_0+mu^2 u+lambda_H u^2" && a.Square.Substitution == "mu^2=-lambda_H v^2" && a.Square.Coordinate == "u=phi^dagger phi" && a.Square.ExpandedAfterSeal == "V(u)=c_0-lambda_H v^2 u+lambda_H u^2" && a.Square.CompletedSquare == "V(u)=lambda_H(u-v^2/2)^2+c_0-(1/4)lambda_H v^4" && a.Square.VMinFormula == "V_min=c_0-(1/4)lambda_H v^4" && a.Square.AlgebraicIdentity && !a.Square.NativeHiggsTheorem, Detail: FormatSquare(a.Square)},
			{Name: "record local zero-vacuum offset convention", Passed: a.Offset.ConventionName == "LocalZeroVacuumOffsetConvention" && a.Offset.Condition == "V_min=0" && strings.Contains(a.Offset.C0Formula, "lambda_runtime_eff") && closeRel(a.Offset.C0LocalBridgeGeV4, 119127483.0758411, 1e-14) && closeAbs(a.Offset.VMinWithLocalOffsetGeV4, 0, 1e-15) && a.Offset.LocalOffsetConvention && !a.Offset.CosmologicalConstantTheorem && !a.Offset.VacuumEnergyDerivation, Detail: FormatOffset(a.Offset)},
			{Name: "write local sealed potential in complex and real four-coordinate forms", Passed: a.Local.ComplexCoordinateForm == "V_local(phi)=lambda_runtime_eff(phi^dagger phi-v^2/2)^2" && a.Local.RealFourCoordinateRule == "phi^dagger phi=(1/2)||x||^2" && a.Local.RealFourCoordinateForm == "V_local(x)=(lambda_runtime_eff/4)(||x||^2-v^2)^2" && closeRel(a.Local.LambdaRuntimeEff, lambdaRuntimeEff, 1e-15) && closeRel(a.Local.VEVGeV, vevConventionGeV, 1e-15) && a.Local.NormalizedAfterOffset && !a.Local.NativePotentialTheorem, Detail: FormatLocal(a.Local)},
			{Name: "reconfirm Hessian compatibility and tree proxy", Passed: a.Hessian.VacuumCondition == "||x_0||^2=v^2" && a.Hessian.HessianFormula == "H_V(x_0)=2 lambda_runtime_eff v^2 P_rad" && a.Hessian.SupportProjector == "P_rad=supp(H_V(x_0))" && a.Hessian.TreeProxySquaredFormula == "m_H_tree_proxy^2=2 lambda_runtime_eff v^2" && closeRel(a.Hessian.TreeProxySquaredGeV2, 15720.144400764586, 1e-15) && closeRel(a.Hessian.TreeProxyGeV, 125.38000000304908, 1e-15) && !a.Hessian.TreeProxyPoleMass && !a.Hessian.NativeHistoryLoopTheorem, Detail: FormatHessian(a.Hessian)},
			{Name: "record vacuum orbit and angular flatness", Passed: a.Orbit.ComplexMinimaCondition == "phi^dagger phi=v^2/2" && a.Orbit.RealMinimaCondition == "||x||^2=v^2" && strings.Contains(a.Orbit.OrbitBeforeQuotient, "S^3") && strings.Contains(a.Orbit.AngularFlatness, "flat") && strings.Contains(a.Orbit.RadialNonFlatness, "radial") && !a.Orbit.SelectsCP1Point && !a.Orbit.NativeEWSBTheorem, Detail: FormatOrbit(a.Orbit)},
			{Name: "record source-type interpretation", Passed: strings.Contains(a.SourceTypes.LambdaRuntimeEff, "Gate770") && strings.Contains(a.SourceTypes.V, "VEV") && strings.Contains(a.SourceTypes.MuSquaredBridge, "Gate771") && strings.Contains(a.SourceTypes.C0, "local") && strings.Contains(a.SourceTypes.CompletedSquare, "normalized") && strings.Contains(a.SourceTypes.Interpretation, "not a native"), Detail: FormatSourceTypes(a.SourceTypes)},
			{Name: "enforce cosmological and physical firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.CompletedSquareNativeHiggs && !a.Firewalls.C0CosmologicalConstantTheorem && !a.Firewalls.S3OrbitNativeEWSB && !a.Firewalls.TreeHessianPoleMass && !a.Firewalls.LambdaRuntimeIndependentTheorem && !a.Firewalls.VEVNativeTheorem && !a.Firewalls.HiggsMassOrPoleMassTheorem && !a.Firewalls.YukawaOperatorOrEigenvalue && !a.Firewalls.HistoryLoopUnitTheorem && a.Firewalls.Verdict == StatusGate772CompletedSquareBoundary, Detail: FormatFirewalls(a.Firewalls)},
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
		notes := append([]string{a.Truth}, Statuses()...)
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func closeRel(got, want, tol float64) bool {
	if math.IsNaN(got) || math.IsNaN(want) || math.IsInf(got, 0) || math.IsInf(want, 0) {
		return false
	}
	d := math.Abs(got - want)
	if want == 0 {
		return d <= tol
	}
	return d/math.Abs(want) <= tol
}

func closeAbs(got, want, tol float64) bool {
	if math.IsNaN(got) || math.IsNaN(want) || math.IsInf(got, 0) || math.IsInf(want, 0) {
		return false
	}
	return math.Abs(got-want) <= tol
}
