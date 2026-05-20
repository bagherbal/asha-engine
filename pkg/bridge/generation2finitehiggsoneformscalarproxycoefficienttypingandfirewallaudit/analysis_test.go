package generation2finitehiggsoneformscalarproxycoefficienttypingandfirewallaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate753ScalarProxyCoefficientTyping(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate752.Inherited || !a.Gate752.ReducedNormalFormReady || !a.Gate752.LambdaProxyMultiplicative || !a.Gate752.KappaEReducedInserted {
		t.Fatalf("bad Gate752 inheritance: %+v", a.Gate752)
	}
	if !a.Gate620.Inherited || !a.Gate620.ProxyRowsAvailable || !a.Gate620.LowScaleProxyPositive || !a.Gate620.HighScaleRuntimeSignFail || !a.Gate620.SeparateScalarLanes {
		t.Fatalf("bad Gate620 inheritance: %+v", a.Gate620)
	}
	if !a.TraceForms.ATracePositive || !a.TraceForms.BTraceNonNegative || !a.TraceForms.PolynomialTraceFormsNative || !a.TraceForms.EvaluatedYukawaValuesSealed {
		t.Fatalf("bad trace forms: %+v", a.TraceForms)
	}
	if math.Abs(a.BA2.BOverA2MZ-0.33307493962706697) > 1e-15 || a.BA2.NativeOneThirdTheorem {
		t.Fatalf("bad b/a^2 audit: %+v", a.BA2)
	}
	if math.Abs(a.Coefficient.Coefficient-threeEighths) > 1e-15 || !a.Coefficient.GaugeBoundaryNormalization || a.Coefficient.NativeScalarCoefficientTheorem {
		t.Fatalf("bad 3/8 coefficient audit: %+v", a.Coefficient)
	}
}

func TestGate753OneEighthShadowBaseRoleAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.OneEighth.CloseToOneEighth || math.Abs(a.OneEighth.ActualProxy-lambdaProxyMZ) > 1e-14 || math.Abs(a.OneEighth.ShadowIdentityResidual) > 1e-16 {
		t.Fatalf("bad one-eighth shadow: %+v", a.OneEighth)
	}
	if !a.BaseRole.ProxyOutsideLoopBracket || !a.BaseRole.IndependentOfKappaE || !a.BaseRole.IndependentOfFWall3 || !a.BaseRole.IndependentOfLHopf || a.BaseRole.RuntimeDerived {
		t.Fatalf("bad base role: %+v", a.BaseRole)
	}
	if !a.Layers.AllLayersSeparated || !a.Layers.NoCircularRuntimePromotion {
		t.Fatalf("bad source-layer separation: %+v", a.Layers)
	}
	if a.Firewalls.ClaimsNativeBA2OneThird || a.Firewalls.ClaimsNativeThreeEighthsScalar || a.Firewalls.ClaimsNativeScalarProxy || a.Firewalls.ClaimsRuntimeLambda || a.Firewalls.ClaimsHiggsMass || a.Firewalls.ClaimsPoleMass || a.Firewalls.ClaimsYukawaEigenvalues || a.Firewalls.ClaimsFlavorHierarchy || a.Firewalls.ClaimsHistoryLoopUnitSource {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	res := Generation2FiniteHiggsOneFormScalarProxyCoefficientTypingAndFirewallAuditTheorem().Verify()
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
