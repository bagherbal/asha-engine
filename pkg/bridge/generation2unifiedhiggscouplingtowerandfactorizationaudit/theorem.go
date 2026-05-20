package generation2unifiedhiggscouplingtowerandfactorizationaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2UnifiedHiggsCouplingTowerAndFactorizationAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 775 — Unified Higgs Coupling Tower and Factorization Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusGate775UnifiedHiggsCouplingTowerBoundary}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 774 self-coupling ratio invariants", Passed: a.Gate774.Inherited && a.Gate774.PotentialRatioIdentity == "A_3^2=4A_2A_4" && a.Gate774.FeynmanRatioIdentity == "lambda_3^2=3m_h^2lambda_4" && a.Gate774.RatioInvariantsAreTreeLane && !a.Gate774.PhysicalMeasuredTheorem && !a.Gate774.TreeProxyPoleMass, Detail: FormatGate774(a.Gate774)},
			{Name: "define total correction factor", Passed: a.Correction.Defined && a.Correction.CYukawaFormula == "C_Yukawa=3/N_eff" && a.Correction.CHistoryFormula == "C_History=1+L_Hopf(1-kappa_lambda_red)" && strings.Contains(a.Correction.CHiggsFormula, "C_Higgs=C_Yukawa C_History") && closeRel(a.Correction.CYukawa, cYukawaMZ, 1e-15) && closeRel(a.Correction.CHistory, cHistoryMZ, 1e-15) && closeRel(a.Correction.CHiggs, 1.0372205204048603, 1e-15) && closeAbs(a.Correction.ProductResidual, 0, 1e-18) && !a.Correction.NativeHiggsTheorem && !a.Correction.NativeYukawaTheorem && !a.Correction.NativeHistoryTheorem, Detail: FormatCorrection(a.Correction)},
			{Name: "rewrite quartic coefficient with C_Higgs", Passed: a.Quartic.Airlock == "lambda_H := lambda_runtime_eff" && a.Quartic.LambdaRuntimeFormula == "lambda_runtime_eff=(1/8)C_Higgs" && a.Quartic.LambdaHBridgeFormula == "lambda_H_bridge=C_Higgs/8" && closeRel(a.Quartic.LambdaHBridge, 0.12965256505060754, 1e-15) && closeRel(a.Quartic.LambdaRuntimeFromCHiggs, a.Quartic.LambdaHBridge, 1e-15) && closeAbs(a.Quartic.QuarticResidual, 0, 1e-18) && !a.Quartic.IndependentScalarRuntime && !a.Quartic.NativeQuarticTheorem, Detail: FormatQuartic(a.Quartic)},
			{Name: "rewrite completed-square potential with C_Higgs", Passed: a.Potential.Rewritten && a.Potential.RealFourCoordinateFormula == "V_local(x)=(lambda_runtime_eff/4)(||x||^2-v^2)^2" && a.Potential.UnifiedFactorFormula == "V_local(x)=(C_Higgs/32)(||x||^2-v^2)^2" && a.Potential.LambdaRuntimeFactor == "lambda_runtime_eff=C_Higgs/8" && !a.Potential.NativeHiggsTheorem, Detail: FormatPotential(a.Potential)},
			{Name: "write unified radial coupling tower", Passed: a.Tower.Written && a.Tower.MassSquaredFormula == "m_H_tree^2=(C_Higgs/4)v^2" && a.Tower.MassFormula == "m_H_tree=(v/2)sqrt(C_Higgs)" && a.Tower.A2Formula == "A_2=(C_Higgs/8)v^2" && a.Tower.A3Formula == "A_3=(C_Higgs/8)v" && a.Tower.A4Formula == "A_4=C_Higgs/32" && a.Tower.Lambda3Formula == "lambda_3=(3/4)v C_Higgs" && a.Tower.Lambda4Formula == "lambda_4=(3/4)C_Higgs" && !a.Tower.PhysicalMeasured && !a.Tower.PoleMassTheorem, Detail: FormatTower(a.Tower)},
			{Name: "compute numerical tower ledger", Passed: a.Numerical.TowerComputed && a.Numerical.Finite && closeRel(a.Numerical.CHiggs, 1.0372205204048603, 1e-15) && closeRel(a.Numerical.LambdaHBridge, 0.12965256505060754, 1e-15) && closeRel(a.Numerical.MassGeV, 125.38000000304908, 1e-15) && closeRel(a.Numerical.A2GeV2, 7860.072200382293, 1e-15) && closeRel(a.Numerical.A3GeV, 31.923009292084874, 1e-15) && closeRel(a.Numerical.A4, 0.032413141262651886, 1e-15) && closeRel(a.Numerical.Lambda3GeV, 191.53805575250925, 1e-15) && closeRel(a.Numerical.Lambda4, 0.7779153903036453, 1e-15) && closeRel(a.Numerical.A2FromMassRelation, a.Numerical.A2GeV2, 1e-15), Detail: FormatNumerical(a.Numerical)},
			{Name: "record source-type interpretation", Passed: a.Sources.Recorded && strings.Contains(a.Sources.CHiggsRole, "total scalar correction") && strings.Contains(a.Sources.CYukawaRole, "Yukawa trace participation") && strings.Contains(a.Sources.CHistoryRole, "History") && strings.Contains(a.Sources.TowerRole, "baseline completed-square") && strings.Contains(a.Sources.BaselineRole, "not a native Higgs theorem"), Detail: FormatSources(a.Sources)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.CHiggsNativeHiggsTheorem && !a.Firewalls.CYukawaNativeYukawaTheorem && !a.Firewalls.CHistoryNativeHistoryLoopTheorem && !a.Firewalls.CouplingTowerMeasured && !a.Firewalls.TreeProxyPoleMass && !a.Firewalls.LambdaHBridgeIndependentRuntime && !a.Firewalls.YukawaOperatorOrEigenvalue && a.Firewalls.Verdict == StatusGate775UnifiedHiggsCouplingTowerBoundary, Detail: FormatFirewalls(a.Firewalls)},
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
