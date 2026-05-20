package generation2higgsquarticcoefficientairlockandlambdasymbolfirewallaudit

import (
	"strings"
	"testing"
)

func TestGate770LambdaSymbolFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate769.Inherited || !strings.Contains(a.Gate769.PotentialForm, "lambda_H") || a.Gate769.QuarticCoefficientDerived || a.Gate769.MuSquaredDerived || a.Gate769.NativeScalarPotentialTheorem {
		t.Fatalf("bad Gate769 inheritance: %+v", a.Gate769)
	}
	if a.Symbols.SeparatedObjectCount != 4 || len(a.Symbols.Objects) != 4 || a.Symbols.NotationIdentityAllowed || a.Symbols.NativeIdentities {
		t.Fatalf("bad lambda firewall: %+v", a.Symbols)
	}
	for _, symbol := range []string{"lambda_wall", "lambda_proxy", "lambda_runtime_eff", "lambda_H"} {
		if !hasSymbol(a.Symbols.Objects, symbol) {
			t.Fatalf("missing lambda object %s in %+v", symbol, a.Symbols.Objects)
		}
	}
}

func TestGate770CoefficientAirlockAndScaleConvention(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Potential.CoefficientSymbol != "lambda_H" || !a.Potential.ControlsStabilization || !a.Potential.ControlsRadialHessian || a.Potential.TreeProxyRelation != "m_H_tree^2=2 lambda_H v^2" || a.Potential.DerivedByGate769 || a.Potential.NativeQuarticTheorem {
		t.Fatalf("bad potential quartic typing: %+v", a.Potential)
	}
	if a.Runtime.Symbol != "lambda_runtime_eff" || !strings.Contains(a.Runtime.Formula, "3/N_eff") || !strings.Contains(a.Runtime.Formula, "kappa_lambda_red") || a.Runtime.IndependentScalarRuntimeTheorem || a.Runtime.NativeQuarticTheorem {
		t.Fatalf("bad runtime coefficient typing: %+v", a.Runtime)
	}
	if a.Airlock.SealName != "HiggsQuarticRuntimeCoefficientSeal" || a.Airlock.Identification != "lambda_H := lambda_runtime_eff" || !strings.Contains(a.Airlock.ScaleQualifiedIdentification, "M_Z") || !a.Airlock.Required || !a.Airlock.WithoutSealDistinctObjects || a.Airlock.NativeScalarPotentialTheorem || a.Airlock.NativeQuarticTheorem {
		t.Fatalf("bad airlock: %+v", a.Airlock)
	}
	if !a.Scale.ScalarPotentialNormalizationRequired || !a.Scale.RuntimeScaleRequired || !a.Scale.RenormalizationConventionRequired || !a.Scale.TreeRunningOrBridgeRuntimeRequired || !a.Scale.LawfulOnlyAfterAllSpecified {
		t.Fatalf("bad scale/convention firewall: %+v", a.Scale)
	}
}

func TestGate770MuSquaredConsequenceAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.MuSquared.RequiresQuarticAirlock || !a.MuSquared.RequiresVEVSeal || a.MuSquared.Formula != "mu^2_bridge=-lambda_runtime_eff v^2" || !closeRel(a.MuSquared.LambdaRuntime, lambdaRuntimeEff, 1e-15) || !closeRel(a.MuSquared.VEVGeV, vevConventionGeV, 1e-15) || !closeRel(a.MuSquared.MuSquaredBridgeGeV2, -7860.072200382293, 1e-15) || a.MuSquared.NativeMuSquaredTheorem || a.MuSquared.NativeEWSBTheorem {
		t.Fatalf("bad mu-squared consequence: %+v", a.MuSquared)
	}
	if !a.Firewalls.Audited || a.Firewalls.LambdaWallEqualsLambdaH || a.Firewalls.LambdaProxyEqualsLambdaH || a.Firewalls.LambdaRuntimeEffNativeLambdaH || a.Firewalls.AirlockNativeScalarPotentialTheorem || a.Firewalls.MuSquaredBridgeNativeEWSBTheorem || a.Firewalls.TreeProxyPoleMassTheorem || a.Firewalls.RuntimeQuarticIndependentMass || a.Firewalls.NativeQuarticCoefficientTheorem || a.Firewalls.NativeMuSquaredTheorem || a.Firewalls.NativeVEVTheorem || a.Firewalls.HiggsMassOrPoleMassTheorem || a.Firewalls.YukawaOperatorOrEigenvalueTheorem {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate770TheoremStatuses(t *testing.T) {
	res := Generation2HiggsQuarticCoefficientAirlockAndLambdaSymbolFirewallAuditTheorem().Verify()
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
