package generation2vevscaleairlockanddimensionfulhiggstowerfirewallaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2VEVScaleAirlockAndDimensionfulHiggsTowerFirewallAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 777 — VEV Scale Airlock and Dimensionful Higgs Tower Firewall Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusGate777VEVScaleBoundary}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 776 total correction decomposition", Passed: a.Gate776.Inherited && a.Gate776.CHiggsFormula == "C_Higgs=(3/N_eff)[1+L_Hopf(1-kappa_lambda_red)]" && closeRel(a.Gate776.CHiggs, 1.0372205204048603, 1e-15) && closeRel(a.Gate776.DilationFactor, 1.0184402389953279, 1e-15) && a.Gate776.TreeMassFormula == "m_H_tree=(v/2)sqrt(C_Higgs)" && closeRel(a.Gate776.TreeMassGeV, 125.38000000304908, 1e-15) && !a.Gate776.NativeHiggsTheorem && !a.Gate776.DimensionfulMassTheorem && !a.Gate776.PoleMassTheorem, Detail: FormatGate776(a.Gate776)},
			{Name: "audit dimensionless and dimensionful split", Passed: a.Split.Audited && containsAll(a.Split.DimensionlessObjects, []string{"C_Higgs", "C_Yukawa", "C_History", "lambda_H_bridge=C_Higgs/8", "lambda_4=(3/4)C_Higgs"}) && a.Split.DimensionfulScaleSeal == "v" && a.Split.MassPower == 1 && a.Split.A2Power == 2 && a.Split.A3Power == 1 && a.Split.Lambda3Power == 1 && a.Split.MuSquaredPower == 2 && a.Split.C0Power == 4 && !a.Split.CHiggsDimensionfulMassTheorem, Detail: FormatSplit(a.Split)},
			{Name: "record VEV convention seal", Passed: a.VEV.Recorded && a.VEV.SealName == "VEVConventionSeal" && closeRel(a.VEV.ValueGeV, 246.2196508, 1e-15) && a.VEV.PhiNormConvention == "phi^dagger phi=v^2/2" && a.VEV.ExternallyRelatedToFermi && !a.VEV.NativeVEVTheorem && !a.VEV.NativeFermiScaleTheorem, Detail: FormatVEV(a.VEV)},
			{Name: "compute scale sensitivity", Passed: a.Sensitivity.Computed && a.Sensitivity.Finite && a.Sensitivity.MassFractionalFormula == "delta m_H_tree/m_H_tree=delta v/v+(1/2)delta C_Higgs/C_Higgs" && a.Sensitivity.MuSquaredFormula == "delta mu^2/mu^2=delta lambda_H/lambda_H+2 delta v/v" && a.Sensitivity.C0Formula == "delta c0/c0=delta lambda_H/lambda_H+4 delta v/v" && closeRel(a.Sensitivity.MassSensitivityToV, 1, 1e-15) && closeRel(a.Sensitivity.MassSensitivityToC, 0.5, 1e-15) && a.Sensitivity.MuSquaredSensitivityToV == 2 && a.Sensitivity.C0SensitivityToV == 4, Detail: FormatSensitivity(a.Sensitivity)},
			{Name: "record baseline scale interpretation", Passed: a.Baseline.Recorded && a.Baseline.BaselineFormula == "m_baseline=v/2" && closeRel(a.Baseline.BaselineGeV, 123.1098254, 1e-15) && a.Baseline.DilationFormula == "D_radial=sqrt(C_Higgs)" && closeRel(a.Baseline.DilationFactor, 1.0184402389953279, 1e-15) && a.Baseline.MassFormula == "m_H_tree_proxy=(v/2)D_radial" && closeRel(a.Baseline.MassGeV, 125.38000000304908, 1e-15) && closeRel(a.Baseline.MassFromBaselineGeV, a.Baseline.MassGeV, 1e-15) && closeAbs(a.Baseline.Residual, 0, 1e-18) && !a.Baseline.DerivedMassTheorem, Detail: FormatBaseline(a.Baseline)},
			{Name: "record remaining source pressure split", Passed: a.Pressure.Recorded && containsAll(a.Pressure.DimensionlessTargets, []string{"N_eff", "kappa_lambda_red", "L_Hopf", "kappa_e_red", "boundary response polynomial"}) && containsAll(a.Pressure.DimensionfulTargets, []string{"v"}) && !a.Pressure.DimensionlessAloneDerivesGeV && a.Pressure.ElectroweakScaleTheoremRequired, Detail: FormatPressure(a.Pressure)},
			{Name: "record derived dimensional ledger", Passed: a.Ledger.Finite && closeRel(a.Ledger.CHiggs, 1.0372205204048603, 1e-15) && closeRel(a.Ledger.LambdaHBridge, 0.12965256505060754, 1e-15) && closeRel(a.Ledger.Lambda4Tree, 0.7779153903036453, 1e-15) && closeRel(a.Ledger.VHalfGeV, 123.1098254, 1e-15) && closeRel(a.Ledger.MassGeV, 125.38000000304908, 1e-15) && closeRel(a.Ledger.A2GeV2, 7860.072200382293, 1e-15) && closeRel(a.Ledger.A3GeV, 31.923009292084874, 1e-15) && closeRel(a.Ledger.Lambda3GeV, 191.53805575250925, 1e-15) && closeRel(a.Ledger.MuSquaredGeV2, -7860.072200382293, 1e-15), Detail: FormatLedger(a.Ledger)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.CHiggsDimensionfulMassTheorem && !a.Firewalls.VEVNativeTheorem && !a.Firewalls.HalfScaleDerivedHiggsMass && !a.Firewalls.FermiScaleNativeASHATheorem && !a.Firewalls.TreeProxyPoleMass && !a.Firewalls.FullHiggsPrediction && !a.Firewalls.YukawaOperatorOrEigenvalue && a.Firewalls.Verdict == StatusGate777VEVScaleBoundary, Detail: FormatFirewalls(a.Firewalls)},
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

func containsAll(haystack, needles []string) bool {
	joined := "\x00" + strings.Join(haystack, "\x00") + "\x00"
	for _, n := range needles {
		if !strings.Contains(joined, "\x00"+n+"\x00") {
			return false
		}
	}
	return true
}
