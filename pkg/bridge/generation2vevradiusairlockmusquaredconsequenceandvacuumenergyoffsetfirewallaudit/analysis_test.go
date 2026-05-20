package generation2vevradiusairlockmusquaredconsequenceandvacuumenergyoffsetfirewallaudit

import (
	"strings"
	"testing"
)

func TestGate771InheritanceAndVEVSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate770.Inherited || a.Gate770.QuarticAirlockSeal != "HiggsQuarticRuntimeCoefficientSeal" || a.Gate770.Identification != "lambda_H := lambda_runtime_eff" || a.Gate770.NativeQuarticCoefficient || a.Gate770.IndependentScalarRuntime {
		t.Fatalf("bad Gate770 inheritance: %+v", a.Gate770)
	}
	if a.VEV.SealName != "VEVConventionSeal" || !closeRel(a.VEV.VEVGeV, vevConventionGeV, 1e-15) || a.VEV.VacuumCoordinate != "u_0=v^2/2" || !closeRel(a.VEV.PhiDaggerPhiAtVacuum, vevConventionGeV*vevConventionGeV/2, 1e-15) || a.VEV.NativeVEVTheorem {
		t.Fatalf("bad VEV convention: %+v", a.VEV)
	}
}

func TestGate771StationarityAndMuSquaredConsequence(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Stationarity.Derivative != "dV/du=mu^2+2 lambda_H u" || a.Stationarity.Consequence != "mu^2=-lambda_H v^2" || !a.Stationarity.RequiresNonzeroVacuum || !a.Stationarity.RequiresQuarticSeal || !a.Stationarity.RequiresVEVSeal || a.Stationarity.NativeEWSBTheorem {
		t.Fatalf("bad stationarity: %+v", a.Stationarity)
	}
	if a.MuSquared.Formula != "mu^2_bridge=-lambda_runtime_eff v^2" || !closeRel(a.MuSquared.LambdaRuntimeEff, lambdaRuntimeEff, 1e-15) || !closeRel(a.MuSquared.VEVGeV, vevConventionGeV, 1e-15) || !closeRel(a.MuSquared.MuSquaredBridgeGeV2, -7860.072200382293, 1e-15) || a.MuSquared.NativeMuSquaredTheorem || a.MuSquared.NativeEWSBTheorem {
		t.Fatalf("bad mu-squared consequence: %+v", a.MuSquared)
	}
}

func TestGate771TreeProxyAndVacuumOffset(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.TreeHessian.Formula != "m_H_tree_proxy^2=-2 mu^2_bridge" || a.TreeHessian.EquivalentFormula != "m_H_tree_proxy^2=2 lambda_runtime_eff v^2" || !closeRel(a.TreeHessian.TreeProxySquaredGeV2, 15720.144400764586, 1e-15) || !closeRel(a.TreeHessian.TreeProxyGeV, 125.38000000304908, 1e-15) || a.TreeHessian.TreeProxyPoleMass || a.TreeHessian.HiggsMassTheorem {
		t.Fatalf("bad tree Hessian relation: %+v", a.TreeHessian)
	}
	if a.Offset.VMinFormula != "V_min=c_0-(1/4)lambda_H v^4" || !strings.Contains(a.Offset.LocalZeroCondition, "c_0=(1/4)lambda_H v^4") || !closeRel(a.Offset.C0LocalBridgeGeV4, 119127483.0758411, 1e-14) || a.Offset.CosmologicalConstantTheorem || a.Offset.VacuumEnergyDerivation {
		t.Fatalf("bad vacuum offset: %+v", a.Offset)
	}
}

func TestGate771FirewallsAndTheoremStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewalls.Audited || a.Firewalls.VEVNativeTheorem || a.Firewalls.MuSquaredNativeTheorem || a.Firewalls.C0CosmologicalConstantTheorem || a.Firewalls.VMinVacuumEnergyDerivation || a.Firewalls.TreeProxyPoleMass || a.Firewalls.LambdaRuntimeIndependentTheorem || a.Firewalls.QuarticAirlockNativeHiggs || a.Firewalls.NativeEWSBTheorem || a.Firewalls.HiggsMassOrPoleMassTheorem || a.Firewalls.YukawaOperatorOrEigenvalue {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	res := Generation2VEVRadiusAirlockMuSquaredConsequenceAndVacuumEnergyOffsetFirewallAuditTheorem().Verify()
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
