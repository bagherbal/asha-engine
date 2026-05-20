package generation2completedsquarehiggspotentialandvacuumoffsetaudit

import (
	"strings"
	"testing"
)

func TestGate772InheritanceAndCompletedSquare(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate771.Inherited || a.Gate771.QuarticAirlock != "HiggsQuarticRuntimeCoefficientSeal" || a.Gate771.VEVSeal != "VEVConventionSeal" || a.Gate771.LambdaHIdentification != "lambda_H := lambda_runtime_eff" || a.Gate771.MuSquaredFormula != "mu^2_bridge=-lambda_runtime_eff v^2" || a.Gate771.NativeVEVTheorem || a.Gate771.NativeMuSquaredTheorem || a.Gate771.C0CosmologyTheorem {
		t.Fatalf("bad Gate771 inheritance: %+v", a.Gate771)
	}
	if a.Square.StartingPotential != "V(u)=c_0+mu^2 u+lambda_H u^2" || a.Square.Substitution != "mu^2=-lambda_H v^2" || a.Square.CompletedSquare != "V(u)=lambda_H(u-v^2/2)^2+c_0-(1/4)lambda_H v^4" || a.Square.VMinFormula != "V_min=c_0-(1/4)lambda_H v^4" || !a.Square.AlgebraicIdentity || a.Square.NativeHiggsTheorem {
		t.Fatalf("bad completed-square form: %+v", a.Square)
	}
}

func TestGate772LocalOffsetAndLocalPotential(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Offset.Condition != "V_min=0" || !strings.Contains(a.Offset.C0Formula, "lambda_runtime_eff") || !closeRel(a.Offset.C0LocalBridgeGeV4, 119127483.0758411, 1e-14) || !closeAbs(a.Offset.VMinWithLocalOffsetGeV4, 0, 1e-15) || !a.Offset.LocalOffsetConvention || a.Offset.CosmologicalConstantTheorem || a.Offset.VacuumEnergyDerivation {
		t.Fatalf("bad local zero offset: %+v", a.Offset)
	}
	if a.Local.ComplexCoordinateForm != "V_local(phi)=lambda_runtime_eff(phi^dagger phi-v^2/2)^2" || a.Local.RealFourCoordinateRule != "phi^dagger phi=(1/2)||x||^2" || a.Local.RealFourCoordinateForm != "V_local(x)=(lambda_runtime_eff/4)(||x||^2-v^2)^2" || !a.Local.NormalizedAfterOffset || a.Local.NativePotentialTheorem {
		t.Fatalf("bad local potential: %+v", a.Local)
	}
}

func TestGate772HessianAndVacuumOrbit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Hessian.VacuumCondition != "||x_0||^2=v^2" || a.Hessian.HessianFormula != "H_V(x_0)=2 lambda_runtime_eff v^2 P_rad" || a.Hessian.SupportProjector != "P_rad=supp(H_V(x_0))" || a.Hessian.TreeProxySquaredFormula != "m_H_tree_proxy^2=2 lambda_runtime_eff v^2" || !closeRel(a.Hessian.TreeProxySquaredGeV2, 15720.144400764586, 1e-15) || !closeRel(a.Hessian.TreeProxyGeV, 125.38000000304908, 1e-15) || a.Hessian.TreeProxyPoleMass || a.Hessian.NativeHistoryLoopTheorem {
		t.Fatalf("bad Hessian compatibility: %+v", a.Hessian)
	}
	if a.Orbit.ComplexMinimaCondition != "phi^dagger phi=v^2/2" || a.Orbit.RealMinimaCondition != "||x||^2=v^2" || !strings.Contains(a.Orbit.OrbitBeforeQuotient, "S^3") || a.Orbit.SelectsCP1Point || a.Orbit.NativeEWSBTheorem {
		t.Fatalf("bad vacuum orbit: %+v", a.Orbit)
	}
}

func TestGate772FirewallsAndTheoremStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewalls.Audited || a.Firewalls.CompletedSquareNativeHiggs || a.Firewalls.C0CosmologicalConstantTheorem || a.Firewalls.S3OrbitNativeEWSB || a.Firewalls.TreeHessianPoleMass || a.Firewalls.LambdaRuntimeIndependentTheorem || a.Firewalls.VEVNativeTheorem || a.Firewalls.HiggsMassOrPoleMassTheorem || a.Firewalls.YukawaOperatorOrEigenvalue || a.Firewalls.HistoryLoopUnitTheorem {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	res := Generation2CompletedSquareHiggsPotentialAndVacuumOffsetAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status note %s", want)
		}
	}
}
