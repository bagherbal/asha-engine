package generation2radialselfcouplingratioinvariantsandconventionfirewallaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2RadialSelfCouplingRatioInvariantsAndConventionFirewallAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 774 — Radial Self-Coupling Ratio Invariants and Convention-Firewall Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusGate774RadialSelfCouplingRatioBoundary}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 773 radial self-coupling coefficients", Passed: a.Gate773.Inherited && a.Gate773.PotentialCoefficientConvention == "V(h)=A_2 h^2+A_3 h^3+A_4 h^4" && a.Gate773.FeynmanRuleConvention == "V(h)=(1/2)m_h^2h^2+(1/3!)lambda_3 h^3+(1/4!)lambda_4 h^4" && a.Gate773.A2Formula == "A_2=lambda_runtime_eff v^2" && a.Gate773.A3Formula == "A_3=lambda_runtime_eff v" && a.Gate773.A4Formula == "A_4=lambda_runtime_eff/4" && a.Gate773.MassSquaredFormula == "m_h^2=2lambda_runtime_eff v^2" && a.Gate773.Lambda3Formula == "lambda_3=6lambda_runtime_eff v" && a.Gate773.Lambda4Formula == "lambda_4=6lambda_runtime_eff" && closeRel(a.Gate773.LambdaRuntimeEff, lambdaRuntimeEff, 1e-15) && closeRel(a.Gate773.VEVGeV, vevConventionGeV, 1e-15) && !a.Gate773.TreeLaneNativeHiggsTheorem, Detail: FormatGate773(a.Gate773)},
			{Name: "derive potential-coefficient ratio invariants", Passed: a.PotentialInvariants.A3SquaredEquals4A2A4 == "A_3^2=4A_2A_4" && a.PotentialInvariants.A3OverA2 == "A_3/A_2=1/v" && a.PotentialInvariants.A4OverA2 == "A_4/A_2=1/(4v^2)" && a.PotentialInvariants.IndependentOfLambdaRuntime && a.PotentialInvariants.CompletedSquareSource && !a.PotentialInvariants.NativePrediction, Detail: FormatPotentialInvariants(a.PotentialInvariants)},
			{Name: "derive Feynman-convention ratio invariants", Passed: a.FeynmanInvariants.Lambda3EqualsVLambda4 == "lambda_3=v lambda_4" && a.FeynmanInvariants.Lambda3SquaredIdentity == "lambda_3^2=3m_h^2lambda_4" && a.FeynmanInvariants.Lambda4MassRelation == "lambda_4=3m_h^2/v^2" && a.FeynmanInvariants.Lambda3MassRelation == "lambda_3=3m_h^2/v" && a.FeynmanInvariants.TreeConventionIdentity && !a.FeynmanInvariants.MeasuredCouplingTheorem, Detail: FormatFeynmanInvariants(a.FeynmanInvariants)},
			{Name: "compute numerical ratio audit", Passed: a.Numerical.AuditComputed && closeRel(a.Numerical.A2GeV2, 7860.072200382293, 1e-15) && closeRel(a.Numerical.A3GeV, 31.923009292084874, 1e-15) && closeRel(a.Numerical.A4, 0.032413141262651886, 1e-15) && closeRel(a.Numerical.MassGeV, 125.38000000304908, 1e-15) && closeRel(a.Numerical.Lambda3GeV, 191.53805575250925, 1e-15) && closeRel(a.Numerical.Lambda4, 0.7779153903036453, 1e-15) && closeAbs(a.Numerical.A3SquaredResidual, 0, 1e-9) && closeRel(a.Numerical.A3OverA2, a.Numerical.OneOverV, 1e-15) && closeRel(a.Numerical.A4OverA2, a.Numerical.OneOver4V2, 1e-15) && closeRel(a.Numerical.Lambda3OverV, a.Numerical.Lambda4, 1e-15) && closeAbs(a.Numerical.Lambda3SquaredResidual, 0, 1e-8), Detail: FormatNumerical(a.Numerical)},
			{Name: "record source-type interpretation", Passed: strings.Contains(a.SourceTypes.Origin, "completed-square") && strings.Contains(a.SourceTypes.LambdaRuntimeRole, "cancels") && strings.Contains(a.SourceTypes.VEVRole, "VEVConventionSeal") && strings.Contains(a.SourceTypes.RatioRole, "internal consistency") && strings.Contains(a.SourceTypes.PredictionFirewall, "not measured"), Detail: FormatSourceTypes(a.SourceTypes)},
			{Name: "audit convention firewall", Passed: a.ConventionFirewall.Audited && a.ConventionFirewall.PotentialCoefficientConvention && a.ConventionFirewall.FeynmanRuleConvention && a.ConventionFirewall.TreeOnly && !a.ConventionFirewall.PhysicalMeasuredCouplings && !a.ConventionFirewall.ColliderObservableTheorem && !a.ConventionFirewall.NativeScalarPotentialTheorem && a.ConventionFirewall.Verdict == StatusConventionFirewallAudited, Detail: FormatConventionFirewall(a.ConventionFirewall)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.SelfCouplingRatiosMeasured && !a.Firewalls.CompletedSquareNativeHiggs && !a.Firewalls.TreeProxyPoleMass && !a.Firewalls.NativeVEVTheorem && !a.Firewalls.NativeScalarRuntimeTheorem && !a.Firewalls.YukawaOperatorOrEigenvalue && !a.Firewalls.NativeEWSBTheorem && !a.Firewalls.HistoryLoopUnitTheorem && a.Firewalls.Verdict == StatusGate774RadialSelfCouplingRatioBoundary, Detail: FormatFirewalls(a.Firewalls)},
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

func closeAbs(got, want, tol float64) bool {
	if math.IsNaN(got) || math.IsNaN(want) || math.IsInf(got, 0) || math.IsInf(want, 0) {
		return false
	}
	return math.Abs(got-want) <= tol
}
