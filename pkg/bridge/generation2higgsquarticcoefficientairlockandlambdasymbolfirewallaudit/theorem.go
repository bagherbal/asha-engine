package generation2higgsquarticcoefficientairlockandlambdasymbolfirewallaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HiggsQuarticCoefficientAirlockAndLambdaSymbolFirewallAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 770 — Higgs Quartic Coefficient Airlock and Lambda-Symbol Firewall Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusGate770QuarticCoefficientAirlockBound}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 769 U(2)-invariant potential form with coefficient seals", Passed: a.Gate769.Inherited && strings.Contains(a.Gate769.PotentialForm, "lambda_H") && a.Gate769.QuarticSymbol == "lambda_H" && !a.Gate769.QuarticCoefficientDerived && !a.Gate769.MuSquaredDerived && !a.Gate769.NativeScalarPotentialTheorem, Detail: FormatGate769(a.Gate769)},
			{Name: "separate lambda-like symbols by layer", Passed: a.Symbols.SeparatedObjectCount == 4 && len(a.Symbols.Objects) == 4 && hasSymbol(a.Symbols.Objects, "lambda_wall") && hasSymbol(a.Symbols.Objects, "lambda_proxy") && hasSymbol(a.Symbols.Objects, "lambda_runtime_eff") && hasSymbol(a.Symbols.Objects, "lambda_H") && !a.Symbols.NotationIdentityAllowed && !a.Symbols.NativeIdentities, Detail: FormatLambdaFirewall(a.Symbols)},
			{Name: "type potential quartic coefficient", Passed: strings.Contains(a.Potential.PotentialForm, "lambda_H") && a.Potential.CoefficientSymbol == "lambda_H" && a.Potential.ControlsStabilization && a.Potential.ControlsRadialHessian && a.Potential.TreeProxyRelation == "m_H_tree^2=2 lambda_H v^2" && !a.Potential.DerivedByGate769 && !a.Potential.NativeQuarticTheorem, Detail: FormatPotential(a.Potential)},
			{Name: "type runtime bridge coefficient separately", Passed: a.Runtime.Symbol == "lambda_runtime_eff" && strings.Contains(a.Runtime.Formula, "3/N_eff") && strings.Contains(a.Runtime.Formula, "L_Hopf") && strings.Contains(a.Runtime.TopColorBaseline, "1/8") && strings.Contains(a.Runtime.YukawaParticipationCorrection, "3/N_eff") && strings.Contains(a.Runtime.HistoryLoopTransportUnit, "L_Hopf") && strings.Contains(a.Runtime.ReducedScalarMatchingDeficit, "kappa_lambda_red") && !a.Runtime.IndependentScalarRuntimeTheorem && !a.Runtime.NativeQuarticTheorem, Detail: FormatRuntime(a.Runtime)},
			{Name: "define explicit Higgs quartic runtime coefficient airlock", Passed: a.Airlock.SealName == "HiggsQuarticRuntimeCoefficientSeal" && a.Airlock.Identification == "lambda_H := lambda_runtime_eff" && strings.Contains(a.Airlock.ScaleQualifiedIdentification, "M_Z") && a.Airlock.Required && a.Airlock.WithoutSealDistinctObjects && a.Airlock.TreeProxyAfterSeal == "m_H_tree_proxy^2=2 lambda_runtime_eff v^2" && !a.Airlock.NativeScalarPotentialTheorem && !a.Airlock.NativeQuarticTheorem, Detail: FormatAirlock(a.Airlock)},
			{Name: "audit scale and convention requirements", Passed: a.Scale.ScalarPotentialNormalizationRequired && a.Scale.RuntimeScaleRequired && a.Scale.RenormalizationConventionRequired && a.Scale.TreeRunningOrBridgeRuntimeRequired && strings.Contains(a.Scale.RuntimeScale, "M_Z") && strings.Contains(a.Scale.PotentialConvention, "phi^dagger phi") && a.Scale.LawfulOnlyAfterAllSpecified, Detail: FormatScale(a.Scale)},
			{Name: "record conditional mu-squared consequence", Passed: a.MuSquared.RequiresQuarticAirlock && a.MuSquared.RequiresVEVSeal && a.MuSquared.Formula == "mu^2_bridge=-lambda_runtime_eff v^2" && closeRel(a.MuSquared.LambdaRuntime, lambdaRuntimeEff, 1e-15) && closeRel(a.MuSquared.VEVGeV, vevConventionGeV, 1e-15) && closeRel(a.MuSquared.MuSquaredBridgeGeV2, -7860.072200382293, 1e-15) && !a.MuSquared.NativeMuSquaredTheorem && !a.MuSquared.NativeEWSBTheorem, Detail: FormatMuSquared(a.MuSquared)},
			{Name: "enforce physical and symbol firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.LambdaWallEqualsLambdaH && !a.Firewalls.LambdaProxyEqualsLambdaH && !a.Firewalls.LambdaRuntimeEffNativeLambdaH && !a.Firewalls.AirlockNativeScalarPotentialTheorem && !a.Firewalls.MuSquaredBridgeNativeEWSBTheorem && !a.Firewalls.TreeProxyPoleMassTheorem && !a.Firewalls.RuntimeQuarticIndependentMass && !a.Firewalls.NativeQuarticCoefficientTheorem && !a.Firewalls.NativeMuSquaredTheorem && !a.Firewalls.NativeVEVTheorem && !a.Firewalls.HiggsMassOrPoleMassTheorem && !a.Firewalls.YukawaOperatorOrEigenvalueTheorem && strings.Contains(a.Firewalls.Verdict, StatusGate770QuarticCoefficientAirlockBound), Detail: FormatFirewalls(a.Firewalls)},
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

func hasSymbol(objects []LambdaObject, symbol string) bool {
	for _, obj := range objects {
		if obj.Symbol == symbol && !obj.NativeIdentity {
			return true
		}
	}
	return false
}

func closeRel(got, want, tol float64) bool {
	if math.IsNaN(got) || math.IsNaN(want) || math.IsInf(got, 0) || math.IsInf(want, 0) {
		return false
	}
	d := math.Abs(got - want)
	if want == 0 {
		return d <= tol
	}
	return d/math.Abs(want) <= tol
}
