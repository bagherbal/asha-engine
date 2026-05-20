package generation2radialhiggsselfcouplingandnormalizationaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2RadialHiggsSelfCouplingAndNormalizationAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 773 — Radial Higgs Self-Coupling Boundary Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusGate773RadialHiggsSelfCouplingBoundary}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 772 completed-square sealed potential", Passed: a.Gate772.Inherited && a.Gate772.CompletedSquareForm == "V_local(phi)=lambda_runtime_eff(phi^dagger phi-v^2/2)^2" && a.Gate772.RealFourCoordinateForm == "V_local(x)=(lambda_runtime_eff/4)(||x||^2-v^2)^2" && a.Gate772.QuarticAirlock == "HiggsQuarticRuntimeCoefficientSeal" && a.Gate772.VEVSeal == "VEVConventionSeal" && closeRel(a.Gate772.LambdaRuntimeEff, lambdaRuntimeEff, 1e-15) && closeRel(a.Gate772.VEVGeV, vevConventionGeV, 1e-15) && !a.Gate772.NativeHiggsTheorem, Detail: FormatGate772(a.Gate772)},
			{Name: "define radial field expansion", Passed: a.Radial.VacuumRepresentative == "x_0=v u_rad" && a.Radial.UnitRadialCondition == "||u_rad||=1" && a.Radial.GaugeChoice == "x=(v+h)u_rad" && a.Radial.NormExpression == "||x||^2=(v+h)^2" && strings.Contains(a.Radial.ExpansionExpression, "((v+h)^2-v^2)") && !a.Radial.RadialGaugeNative && !a.Radial.NativeEWSBTheorem, Detail: FormatRadial(a.Radial)},
			{Name: "expand local radial potential", Passed: a.Expansion.StartingPotential == "V_local(h)=(lambda_runtime_eff/4)[(v+h)^2-v^2]^2" && a.Expansion.ExpandedPotential == "V_local(h)=lambda_runtime_eff v^2 h^2+lambda_runtime_eff v h^3+(lambda_runtime_eff/4)h^4" && a.Expansion.A2Formula == "A_2=lambda_runtime_eff v^2" && a.Expansion.A3Formula == "A_3=lambda_runtime_eff v" && a.Expansion.A4Formula == "A_4=lambda_runtime_eff/4" && closeRel(a.Expansion.A2GeV2, 7860.072200382293, 1e-15) && closeRel(a.Expansion.A3GeV, 31.923009292084874, 1e-15) && closeRel(a.Expansion.A4, 0.032413141262651886, 1e-15) && a.Expansion.AlgebraicExpansion && !a.Expansion.NativeHiggsTheorem, Detail: FormatExpansion(a.Expansion)},
			{Name: "reconfirm tree radial mass", Passed: a.Mass.CanonicalMassConvention == "V(h) contains (1/2)m_h^2 h^2" && a.Mass.MassSquaredFormula == "m_H_tree_proxy^2=2lambda_runtime_eff v^2" && closeRel(a.Mass.LambdaRuntimeEff, lambdaRuntimeEff, 1e-15) && closeRel(a.Mass.VEVGeV, vevConventionGeV, 1e-15) && closeRel(a.Mass.MassSquaredGeV2, 15720.144400764586, 1e-15) && closeRel(a.Mass.MassGeV, 125.38000000304908, 1e-15) && !a.Mass.PoleMassTheorem, Detail: FormatMass(a.Mass)},
			{Name: "separate potential-coefficient and Feynman-rule self-coupling conventions", Passed: a.Conventions.PotentialCoefficientConvention == "V(h)=A_2h^2+A_3h^3+A_4h^4" && a.Conventions.FeynmanRuleConvention == "V(h)=(1/2)m_h^2h^2+(1/3!)lambda_3 h^3+(1/4!)lambda_4 h^4" && a.Conventions.Lambda3Formula == "lambda_3=6lambda_runtime_eff v" && a.Conventions.Lambda4Formula == "lambda_4=6lambda_runtime_eff" && a.Conventions.Lambda3Alternative == "lambda_3=3m_h^2/v" && a.Conventions.Lambda4Alternative == "lambda_4=3m_h^2/v^2" && a.Conventions.ConventionSeparated && !a.Conventions.PhysicalMeasuredCouplings, Detail: FormatConventions(a.Conventions)},
			{Name: "compute numerical self-coupling ledger", Passed: closeRel(a.Numerical.LambdaRuntimeEff, lambdaRuntimeEff, 1e-15) && closeRel(a.Numerical.VEVGeV, vevConventionGeV, 1e-15) && closeRel(a.Numerical.A2GeV2, 7860.072200382293, 1e-15) && closeRel(a.Numerical.A3GeV, 31.923009292084874, 1e-15) && closeRel(a.Numerical.A4, 0.032413141262651886, 1e-15) && closeRel(a.Numerical.MassGeV, 125.38000000304908, 1e-15) && closeRel(a.Numerical.Lambda3GeV, 191.53805575250925, 1e-15) && closeRel(a.Numerical.Lambda4, 0.7779153903036453, 1e-15) && a.Numerical.LedgerComputed, Detail: FormatNumerical(a.Numerical)},
			{Name: "record source-type interpretation", Passed: strings.Contains(a.SourceTypes.LambdaRuntimeEff, "Gate770") && strings.Contains(a.SourceTypes.V, "VEV") && strings.Contains(a.SourceTypes.H, "radial") && strings.Contains(a.SourceTypes.MassProxy, "not pole") && strings.Contains(a.SourceTypes.Lambda3Lambda4, "tree-level") && strings.Contains(a.SourceTypes.Interpretation, "does not promote"), Detail: FormatSourceTypes(a.SourceTypes)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.RadialExpansionNativeHiggs && !a.Firewalls.TreeSelfCouplingsMeasured && !a.Firewalls.TreeProxyPoleMass && !a.Firewalls.RadialGaugeNativeEWSB && !a.Firewalls.LambdaRuntimeIndependentTheorem && !a.Firewalls.VEVNativeTheorem && !a.Firewalls.YukawaOperatorOrEigenvalue && !a.Firewalls.HistoryLoopUnitTheorem && a.Firewalls.Verdict == StatusGate773RadialHiggsSelfCouplingBoundary, Detail: FormatFirewalls(a.Firewalls)},
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
