package generation2scalarcanonicalnormalizationspectralquarticairlockaudit

import "testing"

func TestGate617ScalarAirlockCore(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited.Verdict != StatusGate616Inherited || a.Inherited.XiBoundary <= 0 || a.Inherited.LambdaLambda12 >= 0 {
		t.Fatalf("bad inherited state: %+v", a.Inherited)
	}
	for _, symbol := range []string{"K_phi", "Lambda_phi", "lambda_canon", "lambda_runtime", "a", "b", "f0", "v"} {
		if !hasCoeff(a.ScalarCoefficients, symbol) {
			t.Fatalf("missing scalar coefficient %s", symbol)
		}
	}
	if !a.RuntimeConvention.CanonicalSMConvention || !a.RuntimeConvention.BridgeRuntime || !a.RuntimeConvention.RequiresMatching {
		t.Fatalf("bad runtime convention: %+v", a.RuntimeConvention)
	}
	if a.CanonicalMap.ExactFormulaNative || a.CanonicalMap.CanonicalField == "" || a.CanonicalMap.CanonicalQuartic == "" {
		t.Fatalf("bad canonical map: %+v", a.CanonicalMap)
	}
}

func TestGate617AirlockFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ABF0Audit.AAvailableNative || !a.ABF0Audit.BAvailableNative || a.ABF0Audit.FormulaCertified || a.ABF0Audit.KPhiAvailableNative {
		t.Fatalf("bad a,b,f0 audit: %+v", a.ABF0Audit)
	}
	if !a.RuntimeAirlock.IsCanonicalRuntimeLedger || !a.RuntimeAirlock.IsV1Transported || !a.RuntimeAirlock.TopMassSensitive || a.RuntimeAirlock.EquivalentToPreCanonical {
		t.Fatalf("bad runtime airlock: %+v", a.RuntimeAirlock)
	}
	if a.StressSealImpact.CanReplaceByLambdaCanon || a.StressSealImpact.CanReplaceByLambdaPhi || a.StressSealImpact.RuntimeStressResidual <= 0 {
		t.Fatalf("bad stress impact: %+v", a.StressSealImpact)
	}
	if a.NativeStatus.NativeKPhi || a.NativeStatus.NativeScalarMetric || a.NativeStatus.NativeLambdaPhi || a.NativeStatus.NativeABF0ToLambda || a.NativeStatus.NativeVEV || a.NativeStatus.NativeMatching || a.NativeStatus.NativeLambdaZero || a.NativeStatus.NativeStressTheorem {
		t.Fatalf("native status violated: %+v", a.NativeStatus)
	}
	if a.Firewalls.ClaimsLambdaZero || a.Firewalls.ClaimsHiggsStability || a.Firewalls.ClaimsHiggsPoleMass || a.Firewalls.ClaimsNativeQuartic || a.Firewalls.ClaimsNativeStressSeal || a.Firewalls.ClaimsGaugeUnification {
		t.Fatalf("firewall violated: %+v", a.Firewalls)
	}
}
