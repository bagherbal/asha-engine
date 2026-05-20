package generation2spectralactionabf0canonicalscalarquarticairlockaudit

import "testing"

func TestGate618ABF0ScalarAirlockCore(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited.Verdict != StatusGate617Inherited || a.Inherited.LambdaRuntime >= 0 || a.Inherited.XiBoundary <= 0 {
		t.Fatalf("bad inherited state: %+v", a.Inherited)
	}
	if len(a.ABTraces) != 2 || !traceHas(a.ABTraces, "a") || !traceHas(a.ABTraces, "b") {
		t.Fatalf("bad a,b traces: %+v", a.ABTraces)
	}
	if !a.ABTraces[0].NativeForm || !a.ABTraces[1].NativeForm || !a.ABTraces[0].ObservedValues || !a.ABTraces[1].ObservedValues {
		t.Fatalf("a,b trace classification wrong: %+v", a.ABTraces)
	}
	if a.KineticAudit.KPhiNative || !a.KineticAudit.ATraceNativeForm || a.KineticAudit.CertifiedFormula {
		t.Fatalf("bad kinetic audit: %+v", a.KineticAudit)
	}
	if a.QuarticAudit.LambdaPhiNative || !a.QuarticAudit.BTraceNativeForm || a.QuarticAudit.CertifiedFormula {
		t.Fatalf("bad quartic audit: %+v", a.QuarticAudit)
	}
}

func TestGate618AirlockFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.RatioAudit.CandidateFormula == "" || a.RatioAudit.CLambdaCertified || !a.RatioAudit.RequiresKPhi || !a.RatioAudit.RequiresConvention {
		t.Fatalf("bad ratio audit: %+v", a.RatioAudit)
	}
	if len(a.Conventions) < 6 || allConventionsCertified(a.Conventions) {
		t.Fatalf("bad conventions: %+v", a.Conventions)
	}
	if !a.RuntimeConnection.LambdaMZCanonical || !a.RuntimeConnection.LambdaLambda12V1 || a.RuntimeConnection.MatchingTheorem || a.RuntimeConnection.EquivalentToLambdaCanon {
		t.Fatalf("bad runtime connection: %+v", a.RuntimeConnection)
	}
	if a.StressImpact.CanLiftToLambdaCanon || a.StressImpact.CanNumericallyFixCLambda || a.StressImpact.LambdaRuntime >= 0 {
		t.Fatalf("bad stress impact: %+v", a.StressImpact)
	}
	if a.NativeStatus.NativeKPhi || a.NativeStatus.NativeLambdaPhi || a.NativeStatus.NativeCLambda || a.NativeStatus.NativeABF0ToLambda || a.NativeStatus.NativeRuntimeMatch || a.NativeStatus.NativeVEV || a.NativeStatus.NativeStress {
		t.Fatalf("native status violated: %+v", a.NativeStatus)
	}
	if a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsHiggsStability || a.Firewalls.ClaimsLambdaZero || a.Firewalls.ClaimsGaugeUnification || a.Firewalls.ClaimsNativeStress {
		t.Fatalf("firewall violated: %+v", a.Firewalls)
	}
}
