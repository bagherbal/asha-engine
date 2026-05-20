package generation2finitehiggsoneformscalarproxycoefficienttypingandfirewallaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2FiniteHiggsOneFormScalarProxyCoefficientTypingAndFirewallAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 753 — Finite Higgs One-Form Scalar Proxy Coefficient Typing and Firewall Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate753 scalar proxy coefficient typing audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate752 reduced scalar-Higgs normal form", Passed: a.Gate752.Inherited && a.Gate752.ReducedNormalFormReady && a.Gate752.LambdaProxyMultiplicative && a.Gate752.KappaEReducedInserted && a.Gate752.RuntimePredictionBlocked && a.Gate752.HiggsMassBlocked, Detail: FormatGate752(a.Gate752)},
			{Name: "inherit Gate620 scalar proxy lane", Passed: a.Gate620.Inherited && a.Gate620.ProxyRowsAvailable && a.Gate620.LowScaleProxyPositive && a.Gate620.HighScaleProxyPositive && a.Gate620.LowScaleCloseToRuntime && a.Gate620.HighScaleRuntimeSignFail && a.Gate620.SeparateScalarLanes, Detail: FormatGate620(a.Gate620)},
			{Name: "type finite spectral-action trace forms", Passed: a.TraceForms.ATracePositive && a.TraceForms.BTraceNonNegative && a.TraceForms.PolynomialTraceFormsNative && a.TraceForms.EvaluatedYukawaValuesSealed && strings.Contains(a.TraceForms.AFormula, "Y_u") && strings.Contains(a.TraceForms.BFormula, "^2"), Detail: FormatTraceForms(a.TraceForms)},
			{Name: "audit b/a^2 ratio", Passed: a.BA2.NonNegativeRatio && a.BA2.DimensionlessRatio && a.BA2.TopDominanceCandidate && !a.BA2.NativeOneThirdTheorem && math.Abs(a.BA2.BOverA2MZ-0.33307493962706697) < 1e-15, Detail: FormatBA2(a.BA2)},
			{Name: "type 3/8 coefficient as scalar-proxy airlock", Passed: math.Abs(a.Coefficient.Coefficient-threeEighths) < 1e-15 && a.Coefficient.GaugeBoundaryNormalization && !a.Coefficient.ScalarPotentialCoefficient && a.Coefficient.ScalarConventionAirlock && !a.Coefficient.NativeScalarCoefficientTheorem, Detail: FormatCoefficient(a.Coefficient)},
			{Name: "compute one-eighth proxy shadow", Passed: a.OneEighth.CloseToOneEighth && math.Abs(a.OneEighth.IdealProxy-oneEighth) < 1e-15 && math.Abs(a.OneEighth.ActualProxy-lambdaProxyMZ) < 1e-14 && math.Abs(a.OneEighth.ProxyMinusOneEighth+9.689763985000488e-05) < 1e-14 && math.Abs(a.OneEighth.ShadowIdentityResidual) < 1e-16, Detail: FormatOneEighth(a.OneEighth)},
			{Name: "audit multiplicative base role", Passed: a.BaseRole.ProxyOutsideLoopBracket && a.BaseRole.IndependentOfKappaE && a.BaseRole.IndependentOfFWall3 && a.BaseRole.IndependentOfLHopf && !a.BaseRole.RuntimeDerived && a.BaseRole.TransportFactor > 1 && math.Abs(a.BaseRole.LambdaProxy-lambdaProxyMZ) < 1e-14, Detail: FormatBaseRole(a.BaseRole)},
			{Name: "separate source layers", Passed: a.Layers.AllLayersSeparated && a.Layers.NoCircularRuntimePromotion && len(a.Layers.NativeTraceShapeLayer) == 2 && len(a.Layers.BridgeCoefficientLayer) == 2 && len(a.Layers.EnvironmentalValueLayer) == 2 && len(a.Layers.RuntimeTransportLayer) == 2, Detail: FormatLayers(a.Layers)},
			{Name: "preserve scalar-proxy physical firewalls", Passed: !a.Firewalls.ClaimsNativeBA2OneThird && !a.Firewalls.ClaimsNativeThreeEighthsScalar && !a.Firewalls.ClaimsNativeScalarProxy && !a.Firewalls.ClaimsRuntimeLambda && !a.Firewalls.ClaimsHiggsMass && !a.Firewalls.ClaimsPoleMass && !a.Firewalls.ClaimsYukawaEigenvalues && !a.Firewalls.ClaimsFlavorHierarchy && !a.Firewalls.ClaimsHistoryLoopUnitSource, Detail: FormatFirewalls(a.Firewalls)},
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
