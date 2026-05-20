package generation2vevscaleairlockanddimensionfulhiggstowerfirewallaudit

import (
	"strings"
	"testing"
)

func TestGate777InheritanceSplitAndVEVSeal(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate776.Inherited || a.Gate776.CHiggsFormula != "C_Higgs=(3/N_eff)[1+L_Hopf(1-kappa_lambda_red)]" || !closeRel(a.Gate776.CHiggs, 1.0372205204048603, 1e-15) || !closeRel(a.Gate776.DilationFactor, 1.0184402389953279, 1e-15) || a.Gate776.NativeHiggsTheorem || a.Gate776.DimensionfulMassTheorem || a.Gate776.PoleMassTheorem {
		t.Fatalf("bad Gate776 inheritance: %+v", a.Gate776)
	}
	if !a.Split.Audited || !containsAll(a.Split.DimensionlessObjects, []string{"C_Higgs", "C_Yukawa", "C_History", "lambda_H_bridge=C_Higgs/8", "lambda_4=(3/4)C_Higgs"}) || a.Split.DimensionfulScaleSeal != "v" || a.Split.MassPower != 1 || a.Split.A2Power != 2 || a.Split.A3Power != 1 || a.Split.Lambda3Power != 1 || a.Split.MuSquaredPower != 2 || a.Split.C0Power != 4 || a.Split.CHiggsDimensionfulMassTheorem {
		t.Fatalf("bad split: %+v", a.Split)
	}
	if !a.VEV.Recorded || a.VEV.SealName != "VEVConventionSeal" || !closeRel(a.VEV.ValueGeV, 246.2196508, 1e-15) || a.VEV.PhiNormConvention != "phi^dagger phi=v^2/2" || !a.VEV.ExternallyRelatedToFermi || a.VEV.NativeVEVTheorem || a.VEV.NativeFermiScaleTheorem {
		t.Fatalf("bad VEV seal: %+v", a.VEV)
	}
}

func TestGate777SensitivityBaselineAndPressure(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sensitivity.Computed || !a.Sensitivity.Finite || a.Sensitivity.MassFractionalFormula != "delta m_H_tree/m_H_tree=delta v/v+(1/2)delta C_Higgs/C_Higgs" || a.Sensitivity.MuSquaredFormula != "delta mu^2/mu^2=delta lambda_H/lambda_H+2 delta v/v" || a.Sensitivity.C0Formula != "delta c0/c0=delta lambda_H/lambda_H+4 delta v/v" || !closeRel(a.Sensitivity.MassSensitivityToV, 1, 1e-15) || !closeRel(a.Sensitivity.MassSensitivityToC, 0.5, 1e-15) || a.Sensitivity.MuSquaredSensitivityToV != 2 || a.Sensitivity.C0SensitivityToV != 4 {
		t.Fatalf("bad sensitivity: %+v", a.Sensitivity)
	}
	if !a.Baseline.Recorded || a.Baseline.BaselineFormula != "m_baseline=v/2" || !closeRel(a.Baseline.BaselineGeV, 123.1098254, 1e-15) || a.Baseline.DilationFormula != "D_radial=sqrt(C_Higgs)" || !closeRel(a.Baseline.DilationFactor, 1.0184402389953279, 1e-15) || !closeRel(a.Baseline.MassGeV, 125.38000000304908, 1e-15) || a.Baseline.DerivedMassTheorem {
		t.Fatalf("bad baseline: %+v", a.Baseline)
	}
	if !a.Pressure.Recorded || !containsAll(a.Pressure.DimensionlessTargets, []string{"N_eff", "kappa_lambda_red", "L_Hopf", "kappa_e_red", "boundary response polynomial"}) || !containsAll(a.Pressure.DimensionfulTargets, []string{"v"}) || a.Pressure.DimensionlessAloneDerivesGeV || !a.Pressure.ElectroweakScaleTheoremRequired {
		t.Fatalf("bad pressure split: %+v", a.Pressure)
	}
}

func TestGate777DerivedLedgerAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Ledger.Finite || !closeRel(a.Ledger.CHiggs, 1.0372205204048603, 1e-15) || !closeRel(a.Ledger.LambdaHBridge, 0.12965256505060754, 1e-15) || !closeRel(a.Ledger.Lambda4Tree, 0.7779153903036453, 1e-15) || !closeRel(a.Ledger.VHalfGeV, 123.1098254, 1e-15) || !closeRel(a.Ledger.MassGeV, 125.38000000304908, 1e-15) || !closeRel(a.Ledger.A2GeV2, 7860.072200382293, 1e-15) || !closeRel(a.Ledger.A3GeV, 31.923009292084874, 1e-15) || !closeRel(a.Ledger.Lambda3GeV, 191.53805575250925, 1e-15) || !closeRel(a.Ledger.MuSquaredGeV2, -7860.072200382293, 1e-15) {
		t.Fatalf("bad derived ledger: %+v", a.Ledger)
	}
	if !a.Firewalls.Audited || a.Firewalls.CHiggsDimensionfulMassTheorem || a.Firewalls.VEVNativeTheorem || a.Firewalls.HalfScaleDerivedHiggsMass || a.Firewalls.FermiScaleNativeASHATheorem || a.Firewalls.TreeProxyPoleMass || a.Firewalls.FullHiggsPrediction || a.Firewalls.YukawaOperatorOrEigenvalue {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate777TheoremStatuses(t *testing.T) {
	res := Generation2VEVScaleAirlockAndDimensionfulHiggsTowerFirewallAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status note %s", want)
		}
	}
}
