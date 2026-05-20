package generation2higgstotalcorrectiondecompositionanddilationaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HiggsTotalCorrectionDecompositionAndDilationAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 776 — Higgs Total Correction Decomposition and Dilation Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusGate776TotalCorrectionBoundary}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 775 unified coupling tower", Passed: a.Gate775.Inherited && a.Gate775.UnifiedTower && strings.Contains(a.Gate775.CHiggsFormula, "C_Higgs") && a.Gate775.CYukawaFormula == "C_Yukawa=3/N_eff" && a.Gate775.CHistoryFormula == "C_History=1+L_Hopf(1-kappa_lambda_red)" && !a.Gate775.NativeHiggsTheorem && !a.Gate775.NativeYukawaTheorem && !a.Gate775.NativeHistoryTheorem && !a.Gate775.PhysicalSelfCouplings, Detail: FormatGate775(a.Gate775)},
			{Name: "define History uplift and Yukawa dilution", Passed: a.Definitions.Defined && a.Definitions.DeltaHistoryFormula == "delta_History=L_Hopf(1-kappa_lambda_red)" && a.Definitions.EpsilonYukawaFormula == "epsilon_Yukawa=1-3/N_eff" && a.Definitions.CHistoryFormula == "C_History=1+delta_History" && a.Definitions.CYukawaFormula == "C_Yukawa=1-epsilon_Yukawa" && !a.Definitions.NativeHistoryTheorem && !a.Definitions.NativeYukawaTheorem, Detail: FormatDefinitions(a.Definitions)},
			{Name: "compute total correction expansion", Passed: a.Expansion.Computed && a.Expansion.CHiggsFormula == "C_Higgs=(1-epsilon_Yukawa)(1+delta_History)" && a.Expansion.ExpansionFormula == "C_Higgs=1+delta_History-epsilon_Yukawa-epsilon_Yukawa delta_History" && a.Expansion.DeltaHiggsFormula == "Delta_Higgs=C_Higgs-1=delta_History-epsilon_Yukawa(1+delta_History)" && a.Expansion.YukawaDragFormula == "epsilon_Yukawa(1+delta_History)" && !a.Expansion.NativeHiggsTheorem && !a.Expansion.IndependentRuntimeTheorem, Detail: FormatExpansion(a.Expansion)},
			{Name: "record numerical decomposition ledger", Passed: a.Numerical.Recorded && a.Numerical.Finite && closeRel(a.Numerical.DeltaHistory, 0.03802517792362492, 1e-14) && closeRel(a.Numerical.EpsilonYukawa, 0.0007751811187991509, 1e-15) && closeRel(a.Numerical.YukawaDrag, 0.0008046575187645232, 1e-15) && closeRel(a.Numerical.DeltaHiggs, 0.03722052040486035, 1e-15) && closeRel(a.Numerical.CHiggs, 1.0372205204048603, 1e-15) && closeRel(a.Numerical.CHiggsFromExpanded, a.Numerical.CHiggs, 1e-15) && closeAbs(a.Numerical.ExpansionResidual, 0, 1e-15) && a.Numerical.HistoryDominates, Detail: FormatNumerical(a.Numerical)},
			{Name: "compute radial dilation factor", Passed: a.Dilation.Computed && a.Dilation.Formula == "D_radial=sqrt(C_Higgs)" && a.Dilation.BaselineMassFormula == "m_baseline=v/2" && closeRel(a.Dilation.DilationFactor, 1.0184402389953279, 1e-15) && closeRel(a.Dilation.BaselineMassGeV, 123.1098254, 1e-15) && closeRel(a.Dilation.MassGeV, 125.38000000304908, 1e-15) && closeRel(a.Dilation.MassFromDilationGeV, a.Dilation.MassGeV, 1e-15) && closeAbs(a.Dilation.MassResidual, 0, 1e-18) && !a.Dilation.PoleMassCorrection, Detail: FormatDilation(a.Dilation)},
			{Name: "rewrite coupling tower with Delta_Higgs", Passed: a.Tower.Rewritten && a.Tower.LambdaHFormula == "lambda_H_bridge=(1/8)(1+Delta_Higgs)" && a.Tower.PotentialFormula == "V_local(x)=[(1+Delta_Higgs)/32](||x||^2-v^2)^2" && a.Tower.MassFormula == "m_H_tree=(v/2)sqrt(1+Delta_Higgs)" && a.Tower.Lambda3Formula == "lambda_3=(3/4)v(1+Delta_Higgs)" && a.Tower.Lambda4Formula == "lambda_4=(3/4)(1+Delta_Higgs)" && a.Tower.A2Formula == "A_2=(1/8)(1+Delta_Higgs)v^2" && a.Tower.A3Formula == "A_3=(1/8)(1+Delta_Higgs)v" && a.Tower.A4Formula == "A_4=(1/32)(1+Delta_Higgs)" && !a.Tower.PhysicalSelfCoupling && !a.Tower.PoleMassTheorem, Detail: FormatTower(a.Tower)},
			{Name: "record source-type interpretation", Passed: a.Sources.Recorded && strings.Contains(a.Sources.DeltaHistoryRole, "History") && strings.Contains(a.Sources.EpsilonYukawaRole, "Yukawa") && strings.Contains(a.Sources.YukawaDragRole, "multiplicative drag") && strings.Contains(a.Sources.DeltaHiggsRole, "net sealed") && strings.Contains(a.Sources.TowerRole, "square-root dilation"), Detail: FormatSources(a.Sources)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.DeltaHiggsNativeHiggsTheorem && !a.Firewalls.HistoryUpliftNativeHistoryLoop && !a.Firewalls.YukawaDilutionNativeYukawa && !a.Firewalls.RadialDilationPoleMassCorrection && !a.Firewalls.TreeTowerMeasuredSelfCouplings && !a.Firewalls.YukawaOperatorOrEigenvalue && a.Firewalls.Verdict == StatusGate776TotalCorrectionBoundary, Detail: FormatFirewalls(a.Firewalls)},
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
