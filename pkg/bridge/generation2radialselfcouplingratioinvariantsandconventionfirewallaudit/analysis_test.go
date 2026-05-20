package generation2radialselfcouplingratioinvariantsandconventionfirewallaudit

import (
	"strings"
	"testing"
)

func TestGate774InheritanceAndPotentialInvariants(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate773.Inherited || a.Gate773.PotentialCoefficientConvention != "V(h)=A_2 h^2+A_3 h^3+A_4 h^4" || a.Gate773.FeynmanRuleConvention != "V(h)=(1/2)m_h^2h^2+(1/3!)lambda_3 h^3+(1/4!)lambda_4 h^4" || a.Gate773.TreeLaneNativeHiggsTheorem {
		t.Fatalf("bad Gate773 inheritance: %+v", a.Gate773)
	}
	if a.PotentialInvariants.A3SquaredEquals4A2A4 != "A_3^2=4A_2A_4" || a.PotentialInvariants.A3OverA2 != "A_3/A_2=1/v" || a.PotentialInvariants.A4OverA2 != "A_4/A_2=1/(4v^2)" || !a.PotentialInvariants.IndependentOfLambdaRuntime || !a.PotentialInvariants.CompletedSquareSource || a.PotentialInvariants.NativePrediction {
		t.Fatalf("bad potential invariants: %+v", a.PotentialInvariants)
	}
}

func TestGate774FeynmanInvariantsAndNumerics(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.FeynmanInvariants.Lambda3EqualsVLambda4 != "lambda_3=v lambda_4" || a.FeynmanInvariants.Lambda3SquaredIdentity != "lambda_3^2=3m_h^2lambda_4" || a.FeynmanInvariants.Lambda4MassRelation != "lambda_4=3m_h^2/v^2" || a.FeynmanInvariants.Lambda3MassRelation != "lambda_3=3m_h^2/v" || !a.FeynmanInvariants.TreeConventionIdentity || a.FeynmanInvariants.MeasuredCouplingTheorem {
		t.Fatalf("bad Feynman invariants: %+v", a.FeynmanInvariants)
	}
	if !a.Numerical.AuditComputed || !closeRel(a.Numerical.A2GeV2, 7860.072200382293, 1e-15) || !closeRel(a.Numerical.A3GeV, 31.923009292084874, 1e-15) || !closeRel(a.Numerical.A4, 0.032413141262651886, 1e-15) || !closeRel(a.Numerical.MassGeV, 125.38000000304908, 1e-15) || !closeRel(a.Numerical.Lambda3GeV, 191.53805575250925, 1e-15) || !closeRel(a.Numerical.Lambda4, 0.7779153903036453, 1e-15) {
		t.Fatalf("bad numerical ledger: %+v", a.Numerical)
	}
	if !closeAbs(a.Numerical.A3SquaredResidual, 0, 1e-9) || !closeRel(a.Numerical.A3OverA2, a.Numerical.OneOverV, 1e-15) || !closeRel(a.Numerical.A4OverA2, a.Numerical.OneOver4V2, 1e-15) || !closeRel(a.Numerical.Lambda3OverV, a.Numerical.Lambda4, 1e-15) || !closeAbs(a.Numerical.Lambda3SquaredResidual, 0, 1e-8) {
		t.Fatalf("ratio invariant residuals failed: %+v", a.Numerical)
	}
}

func TestGate774ConventionAndPhysicalFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ConventionFirewall.Audited || !a.ConventionFirewall.PotentialCoefficientConvention || !a.ConventionFirewall.FeynmanRuleConvention || !a.ConventionFirewall.TreeOnly || a.ConventionFirewall.PhysicalMeasuredCouplings || a.ConventionFirewall.ColliderObservableTheorem || a.ConventionFirewall.NativeScalarPotentialTheorem {
		t.Fatalf("bad convention firewall: %+v", a.ConventionFirewall)
	}
	if !a.Firewalls.Audited || a.Firewalls.SelfCouplingRatiosMeasured || a.Firewalls.CompletedSquareNativeHiggs || a.Firewalls.TreeProxyPoleMass || a.Firewalls.NativeVEVTheorem || a.Firewalls.NativeScalarRuntimeTheorem || a.Firewalls.YukawaOperatorOrEigenvalue || a.Firewalls.NativeEWSBTheorem || a.Firewalls.HistoryLoopUnitTheorem {
		t.Fatalf("bad physical firewalls: %+v", a.Firewalls)
	}
}

func TestGate774TheoremStatuses(t *testing.T) {
	res := Generation2RadialSelfCouplingRatioInvariantsAndConventionFirewallAuditTheorem().Verify()
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
