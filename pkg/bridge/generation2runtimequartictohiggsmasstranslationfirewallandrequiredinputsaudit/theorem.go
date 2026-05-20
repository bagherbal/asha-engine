package generation2runtimequartictohiggsmasstranslationfirewallandrequiredinputsaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2RuntimeQuarticToHiggsMassTranslationFirewallAndRequiredInputsAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 740 — Runtime Quartic to Higgs-Mass Translation Firewall and Required Inputs Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate740 Higgs-mass translation firewall audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate739 Level-1 scalar runtime estimate", Passed: a.Gate739.Inherited && a.Gate739.Level1Allowed && a.Gate739.NotIndependentPrediction && a.Gate739.RuntimeBridgeNotHiggsMass && strings.Contains(a.Gate739.Verdict, StatusGate739Level1ScalarRuntimeEstimateInherited), Detail: FormatGate739(a.Gate739)},
			{Name: "classify runtime quartic", Passed: a.Quartic.ClassifiedAsRuntimeQuartic && a.Quartic.BridgeLayer && !a.Quartic.PhysicalPoleMass && strings.Contains(a.Quartic.Verdict, StatusRuntimeQuarticClassified), Detail: FormatQuartic(a.Quartic)},
			{Name: "audit tree-level proxy relation", Passed: a.Proxy.RequiresV && a.Proxy.RequiresConvention && a.Proxy.ConventionDependent && !a.Proxy.PoleMassTheorem && a.Proxy.SqrtTwoLambdaFactor > 0 && math.Abs(a.Proxy.LambdaRuntime-a.Gate739.LambdaRuntimeBridge) < 1e-16 && strings.Contains(a.Proxy.Verdict, StatusTreeLevelProxyRelationAudited), Detail: FormatProxy(a.Proxy)},
			{Name: "list Higgs-mass required inputs", Passed: a.Required.AllListed && len(a.Required.Inputs) == 7 && a.Required.HasVEV && a.Required.HasConvention && a.Required.HasScaleMatching && a.Required.HasRGTransport && a.Required.HasThresholdLoop && a.Required.HasGaugeYukawaTop && a.Required.HasUncertainty && strings.Contains(a.Required.Verdict, StatusHiggsMassRequiredInputsListed), Detail: FormatRequired(a.Required)},
			{Name: "enforce proxy versus pole firewall", Passed: !a.Firewall.RuntimeLambdaEqualsPoleMass && !a.Firewall.TreeProxyEqualsPoleMass && !a.Firewall.NearAgreementIsIndependentPrediction && a.Firewall.RuntimeLambdaNotPoleMass && a.Firewall.TreeProxyNotPoleMassTheorem && strings.Contains(a.Firewall.Verdict, StatusRuntimeLambdaNotPoleMass), Detail: FormatFirewall(a.Firewall)},
			{Name: "carry seal dependence into mass translation", Passed: a.Seals.Explicit && len(a.Seals.Seals) == 10 && a.Seals.TreeProxyWouldRemainLevel1 && a.Seals.NoSealReduction && strings.Contains(a.Seals.Verdict, StatusSealDependenceCarriedIntoMassTranslation), Detail: FormatSeals(a.Seals)},
			{Name: "refine forecast levels", Passed: a.Forecast.Level1AAllowed && a.Forecast.Level1BAllowed && !a.Forecast.Level2Allowed && strings.Contains(a.Forecast.Verdict, StatusLevel2HiggsPoleMassPredictionNotAllowed), Detail: FormatForecast(a.Forecast)},
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
