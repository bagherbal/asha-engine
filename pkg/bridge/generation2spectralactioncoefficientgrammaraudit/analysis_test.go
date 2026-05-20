package generation2spectralactioncoefficientgrammaraudit

import (
	"math"
	"testing"
)

func TestGate615CoefficientGrammar(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Inherited.XiBoundary-0.0503471644870914) > 1e-12 {
		t.Fatalf("unexpected xi %.15g", a.Inherited.XiBoundary)
	}
	if len(a.Dependencies) < 7 || !hasCoefficient(a.Dependencies, "C_i") || !hasCoefficient(a.Dependencies, "K_phi") || !hasCoefficient(a.Dependencies, "lambda") || !hasCoefficient(a.Dependencies, "f0") {
		t.Fatalf("missing coefficient dependencies: %+v", a.Dependencies)
	}
	if !a.JointDeformation.BridgeExpressible || a.JointDeformation.KnownNativeRelation || a.JointDeformation.ForcesStressEquation {
		t.Fatalf("bad joint deformation audit: %+v", a.JointDeformation)
	}
	if a.JointDeformation.ResidualOverXi > 0.03 {
		t.Fatalf("stress residual too large: %+v", a.JointDeformation)
	}
}

func TestGate615FirewallsAndObstructions(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.ColorDeformation.BridgeExpressible || a.ColorDeformation.NativeRepresentationTrace || !a.ColorDeformation.RequiresSectorSplitF0 || !a.ColorDeformation.RequiresThresholdMatching {
		t.Fatalf("bad color deformation audit: %+v", a.ColorDeformation)
	}
	if !a.ScalarCorrection.BridgeExpressible || !a.ScalarCorrection.ViaBA2 || !a.ScalarCorrection.ViaF0 || !a.ScalarCorrection.ViaMatching || a.ScalarCorrection.Native {
		t.Fatalf("bad scalar correction audit: %+v", a.ScalarCorrection)
	}
	if a.TypeConsistency.RawComparisonSafe || !a.TypeConsistency.NormalizedSafe {
		t.Fatalf("type consistency violated: %+v", a.TypeConsistency)
	}
	if a.NativeObstructions.NativeXi || a.NativeObstructions.NativeSU3Only || a.NativeObstructions.NativeC3LambdaLaw || a.NativeObstructions.NativeF0Split || a.NativeObstructions.NativeLambdaBC || a.NativeObstructions.NativeThresholds {
		t.Fatalf("native obstruction violated: %+v", a.NativeObstructions)
	}
	if a.Firewalls.ClaimsGaugeUnification || a.Firewalls.ClaimsThresholdExistence || a.Firewalls.ClaimsNativeCorrection || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsHiggsStability || a.Firewalls.ClaimsLambdaZero {
		t.Fatalf("firewall violation: %+v", a.Firewalls)
	}
}
