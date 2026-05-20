package generation2vevradiusairlockmusquaredconsequenceandvacuumenergyoffsetfirewallaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2VEVRadiusAirlockMuSquaredConsequenceAndVacuumEnergyOffsetFirewallAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 771 — VEV Radius Airlock, Mu-Squared Consequence, and Vacuum-Energy Offset Firewall Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusGate771VEVMuSquaredOffsetBound}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 770 quartic coefficient airlock", Passed: a.Gate770.Inherited && a.Gate770.QuarticAirlockSeal == "HiggsQuarticRuntimeCoefficientSeal" && a.Gate770.Identification == "lambda_H := lambda_runtime_eff" && strings.Contains(a.Gate770.LambdaRuntimeFormula, "3/N_eff") && closeRel(a.Gate770.LambdaRuntimeEff, lambdaRuntimeEff, 1e-15) && !a.Gate770.NativeQuarticCoefficient && !a.Gate770.IndependentScalarRuntime, Detail: FormatGate770(a.Gate770)},
			{Name: "define VEV radius convention seal", Passed: a.VEV.SealName == "VEVConventionSeal" && closeRel(a.VEV.VEVGeV, vevConventionGeV, 1e-15) && a.VEV.PotentialCoordinate == "u=phi^dagger phi" && a.VEV.VacuumCoordinate == "u_0=v^2/2" && closeRel(a.VEV.PhiDaggerPhiAtVacuum, vevConventionGeV*vevConventionGeV/2, 1e-15) && !a.VEV.NativeVEVTheorem, Detail: FormatVEV(a.VEV)},
			{Name: "compute vacuum stationarity condition", Passed: a.Stationarity.PotentialForm == "V(u)=c_0+mu^2 u+lambda_H u^2" && a.Stationarity.Derivative == "dV/du=mu^2+2 lambda_H u" && a.Stationarity.StationarityAt == "u_0=v^2/2" && a.Stationarity.Consequence == "mu^2=-lambda_H v^2" && a.Stationarity.RequiresNonzeroVacuum && a.Stationarity.RequiresQuarticSeal && a.Stationarity.RequiresVEVSeal && !a.Stationarity.NativeEWSBTheorem, Detail: FormatStationarity(a.Stationarity)},
			{Name: "compute mu-squared bridge consequence", Passed: a.MuSquared.Formula == "mu^2_bridge=-lambda_runtime_eff v^2" && closeRel(a.MuSquared.LambdaRuntimeEff, lambdaRuntimeEff, 1e-15) && closeRel(a.MuSquared.VEVGeV, vevConventionGeV, 1e-15) && closeRel(a.MuSquared.MuSquaredBridgeGeV2, -7860.072200382293, 1e-15) && a.MuSquared.RequiresQuarticAirlock && a.MuSquared.RequiresVEVSeal && !a.MuSquared.NativeMuSquaredTheorem && !a.MuSquared.NativeEWSBTheorem, Detail: FormatMuSquared(a.MuSquared)},
			{Name: "reconfirm tree Hessian relation", Passed: a.TreeHessian.Formula == "m_H_tree_proxy^2=-2 mu^2_bridge" && a.TreeHessian.EquivalentFormula == "m_H_tree_proxy^2=2 lambda_runtime_eff v^2" && closeRel(a.TreeHessian.TreeProxySquaredGeV2, 15720.144400764586, 1e-15) && closeRel(a.TreeHessian.TreeProxyGeV, 125.38000000304908, 1e-15) && strings.Contains(a.TreeHessian.RadialHessianEigenvalue, "2 lambda_H v^2") && !a.TreeHessian.TreeProxyPoleMass && !a.TreeHessian.HiggsMassTheorem, Detail: FormatTreeHessian(a.TreeHessian)},
			{Name: "compute vacuum-energy offset form", Passed: strings.Contains(a.Offset.PotentialAtVacuum, "V_min") && a.Offset.StationarySubstitution == "mu^2=-lambda_H v^2" && a.Offset.VMinFormula == "V_min=c_0-(1/4)lambda_H v^4" && strings.Contains(a.Offset.LocalZeroCondition, "c_0=(1/4)lambda_H v^4") && closeRel(a.Offset.C0LocalBridgeGeV4, 119127483.0758411, 1e-14) && closeRel(a.Offset.VMinWithoutC0GeV4, -119127483.0758411, 1e-14) && closeAbs(a.Offset.VMinWithLocalOffsetGeV4, 0, 1e-15) && a.Offset.LocalOffsetConvention && !a.Offset.CosmologicalConstantTheorem && !a.Offset.VacuumEnergyDerivation, Detail: FormatOffset(a.Offset)},
			{Name: "record source-type interpretation", Passed: strings.Contains(a.SourceTypes.LambdaH, "Gate770") && strings.Contains(a.SourceTypes.V, "VEV") && strings.Contains(a.SourceTypes.MuSquaredBridge, "stationarity") && strings.Contains(a.SourceTypes.C0, "local") && strings.Contains(a.SourceTypes.TreeProxy, "not pole mass"), Detail: FormatSourceTypes(a.SourceTypes)},
			{Name: "enforce physical and cosmological firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.VEVNativeTheorem && !a.Firewalls.MuSquaredNativeTheorem && !a.Firewalls.C0CosmologicalConstantTheorem && !a.Firewalls.VMinVacuumEnergyDerivation && !a.Firewalls.TreeProxyPoleMass && !a.Firewalls.LambdaRuntimeIndependentTheorem && !a.Firewalls.QuarticAirlockNativeHiggs && !a.Firewalls.NativeEWSBTheorem && !a.Firewalls.HiggsMassOrPoleMassTheorem && !a.Firewalls.YukawaOperatorOrEigenvalue && a.Firewalls.Verdict == StatusGate771VEVMuSquaredOffsetBound, Detail: FormatFirewalls(a.Firewalls)},
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
