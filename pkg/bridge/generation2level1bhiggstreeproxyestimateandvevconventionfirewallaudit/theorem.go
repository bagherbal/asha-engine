package generation2level1bhiggstreeproxyestimateandvevconventionfirewallaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2Level1BHiggsTreeProxyEstimateAndVEVConventionFirewallAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 741 — Level-1B Higgs Tree Proxy Estimate and VEV-Convention Firewall Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate741 Level-1B tree proxy audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate740 Higgs translation firewall", Passed: a.Gate740.Inherited && a.Gate740.Level1BAllowed && a.Gate740.RuntimeLambdaNotPoleMass && a.Gate740.TreeProxyNotPoleMassTheorem && strings.Contains(a.Gate740.Verdict, StatusGate740HiggsTranslationFirewallInherited), Detail: FormatGate740(a.Gate740)},
			{Name: "inherit runtime quartic bridge value", Passed: a.Runtime.ClassifiedAsRuntimeQuartic && a.Runtime.NotIndependentlyDerived && a.Runtime.NotPoleMass && math.Abs(a.Runtime.LambdaRuntimeBridge-0.12965256505047373) < 1e-15 && strings.Contains(a.Runtime.Verdict, StatusRuntimeQuarticBridgeValueInherited), Detail: FormatRuntime(a.Runtime)},
			{Name: "define VEV convention seal", Passed: a.VEV.SuppliedInput && !a.VEV.NativeDerivation && a.VEV.Convention && math.Abs(a.VEV.VGeV-DefaultVEVGeV) < 1e-12 && strings.Contains(a.VEV.Verdict, StatusVEVConventionSealDefined), Detail: FormatVEV(a.VEV)},
			{Name: "apply tree proxy relation", Passed: !a.Proxy.PoleMassPrediction && math.Abs(a.Proxy.TreeProxyGeV-125.38000000298437) < 1e-9 && math.Abs(a.Proxy.SqrtTwoLambdaFactor-math.Sqrt(2*a.Runtime.LambdaRuntimeBridge)) < 1e-16 && strings.Contains(a.Proxy.Verdict, StatusLevel1BTreeProxyEstimateComputed), Detail: FormatProxy(a.Proxy)},
			{Name: "record sensitivity", Passed: a.Sensitivity.LinearInV && a.Sensitivity.HalfPowerInLambda && a.Sensitivity.DeltaMOverMFromDeltaVOverV == 1 && a.Sensitivity.DeltaMOverMFromDeltaLambdaOverLambda == 0.5 && strings.Contains(a.Sensitivity.Verdict, StatusSensitivityToVEVAndLambdaRecorded), Detail: FormatSensitivity(a.Sensitivity)},
			{Name: "carry seals forward", Passed: a.Seals.Explicit && a.Seals.IncludesVEVSeal && a.Seals.IncludesConvention && a.Seals.ProxyRemainsSealed && len(a.Seals.Labels) == 12 && strings.Contains(a.Seals.Verdict, StatusSealDependenceCarriedForward), Detail: FormatSeals(a.Seals)},
			{Name: "enforce pole-mass firewall", Passed: !a.Firewall.TreeProxyEqualsPoleMass && !a.Firewall.RuntimeLambdaIndependentlyDerived && !a.Firewall.HasPoleCorrectionTheorem && !a.Firewall.HasHiggsMassTheorem && !a.Firewall.Level2PredictionAllowed && strings.Contains(a.Firewall.Verdict, StatusPoleMassFirewallEnforced), Detail: FormatFirewall(a.Firewall)},
			{Name: "classify forecast level", Passed: a.Level.Level1BAllowed && !a.Level.Level2Allowed && a.Level.ExplicitSeals && strings.Contains(a.Level.Verdict, StatusLevel1BTreeProxyEstimateAllowedWithExplicitSeals), Detail: FormatLevel(a.Level)},
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
