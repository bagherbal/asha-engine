package generation2radialhiggsselfcouplingandnormalizationaudit

import (
	"strings"
	"testing"
)

func TestGate773InheritanceAndRadialExpansion(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate772.Inherited || a.Gate772.CompletedSquareForm != "V_local(phi)=lambda_runtime_eff(phi^dagger phi-v^2/2)^2" || a.Gate772.RealFourCoordinateForm != "V_local(x)=(lambda_runtime_eff/4)(||x||^2-v^2)^2" || a.Gate772.QuarticAirlock != "HiggsQuarticRuntimeCoefficientSeal" || a.Gate772.VEVSeal != "VEVConventionSeal" || a.Gate772.NativeHiggsTheorem {
		t.Fatalf("bad Gate772 inheritance: %+v", a.Gate772)
	}
	if a.Radial.VacuumRepresentative != "x_0=v u_rad" || a.Radial.UnitRadialCondition != "||u_rad||=1" || a.Radial.GaugeChoice != "x=(v+h)u_rad" || a.Radial.NormExpression != "||x||^2=(v+h)^2" || !strings.Contains(a.Radial.ExpansionExpression, "((v+h)^2-v^2)") || a.Radial.RadialGaugeNative || a.Radial.NativeEWSBTheorem {
		t.Fatalf("bad radial expansion: %+v", a.Radial)
	}
}

func TestGate773ExpandedPotentialAndMass(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Expansion.StartingPotential != "V_local(h)=(lambda_runtime_eff/4)[(v+h)^2-v^2]^2" || a.Expansion.ExpandedPotential != "V_local(h)=lambda_runtime_eff v^2 h^2+lambda_runtime_eff v h^3+(lambda_runtime_eff/4)h^4" || a.Expansion.A2Formula != "A_2=lambda_runtime_eff v^2" || a.Expansion.A3Formula != "A_3=lambda_runtime_eff v" || a.Expansion.A4Formula != "A_4=lambda_runtime_eff/4" || !closeRel(a.Expansion.A2GeV2, 7860.072200382293, 1e-15) || !closeRel(a.Expansion.A3GeV, 31.923009292084874, 1e-15) || !closeRel(a.Expansion.A4, 0.032413141262651886, 1e-15) || !a.Expansion.AlgebraicExpansion || a.Expansion.NativeHiggsTheorem {
		t.Fatalf("bad expansion: %+v", a.Expansion)
	}
	if a.Mass.CanonicalMassConvention != "V(h) contains (1/2)m_h^2 h^2" || a.Mass.MassSquaredFormula != "m_H_tree_proxy^2=2lambda_runtime_eff v^2" || !closeRel(a.Mass.MassSquaredGeV2, 15720.144400764586, 1e-15) || !closeRel(a.Mass.MassGeV, 125.38000000304908, 1e-15) || a.Mass.PoleMassTheorem {
		t.Fatalf("bad tree radial mass: %+v", a.Mass)
	}
}

func TestGate773SelfCouplingConventionsAndNumerics(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Conventions.PotentialCoefficientConvention != "V(h)=A_2h^2+A_3h^3+A_4h^4" || a.Conventions.FeynmanRuleConvention != "V(h)=(1/2)m_h^2h^2+(1/3!)lambda_3 h^3+(1/4!)lambda_4 h^4" || a.Conventions.Lambda3Formula != "lambda_3=6lambda_runtime_eff v" || a.Conventions.Lambda4Formula != "lambda_4=6lambda_runtime_eff" || a.Conventions.Lambda3Alternative != "lambda_3=3m_h^2/v" || a.Conventions.Lambda4Alternative != "lambda_4=3m_h^2/v^2" || !a.Conventions.ConventionSeparated || a.Conventions.PhysicalMeasuredCouplings {
		t.Fatalf("bad conventions: %+v", a.Conventions)
	}
	if !a.Numerical.LedgerComputed || !closeRel(a.Numerical.A2GeV2, 7860.072200382293, 1e-15) || !closeRel(a.Numerical.A3GeV, 31.923009292084874, 1e-15) || !closeRel(a.Numerical.A4, 0.032413141262651886, 1e-15) || !closeRel(a.Numerical.MassGeV, 125.38000000304908, 1e-15) || !closeRel(a.Numerical.Lambda3GeV, 191.53805575250925, 1e-15) || !closeRel(a.Numerical.Lambda4, 0.7779153903036453, 1e-15) {
		t.Fatalf("bad numerical ledger: %+v", a.Numerical)
	}
}

func TestGate773FirewallsAndTheoremStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewalls.Audited || a.Firewalls.RadialExpansionNativeHiggs || a.Firewalls.TreeSelfCouplingsMeasured || a.Firewalls.TreeProxyPoleMass || a.Firewalls.RadialGaugeNativeEWSB || a.Firewalls.LambdaRuntimeIndependentTheorem || a.Firewalls.VEVNativeTheorem || a.Firewalls.YukawaOperatorOrEigenvalue || a.Firewalls.HistoryLoopUnitTheorem {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	res := Generation2RadialHiggsSelfCouplingAndNormalizationAuditTheorem().Verify()
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
