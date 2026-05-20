package generation2unifiedhiggscouplingtowerandfactorizationaudit

import (
	"strings"
	"testing"
)

func TestGate775TotalCorrectionAndQuartic(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate774.Inherited || !a.Gate774.RatioInvariantsAreTreeLane || a.Gate774.PhysicalMeasuredTheorem || a.Gate774.TreeProxyPoleMass {
		t.Fatalf("bad Gate774 inheritance: %+v", a.Gate774)
	}
	if !a.Correction.Defined || a.Correction.CYukawaFormula != "C_Yukawa=3/N_eff" || a.Correction.CHistoryFormula != "C_History=1+L_Hopf(1-kappa_lambda_red)" || !closeRel(a.Correction.CHiggs, 1.0372205204048603, 1e-15) || !closeAbs(a.Correction.ProductResidual, 0, 1e-18) || a.Correction.NativeHiggsTheorem || a.Correction.NativeYukawaTheorem || a.Correction.NativeHistoryTheorem {
		t.Fatalf("bad correction factor: %+v", a.Correction)
	}
	if a.Quartic.LambdaHBridgeFormula != "lambda_H_bridge=C_Higgs/8" || !closeRel(a.Quartic.LambdaHBridge, 0.12965256505060754, 1e-15) || a.Quartic.IndependentScalarRuntime || a.Quartic.NativeQuarticTheorem {
		t.Fatalf("bad quartic rewrite: %+v", a.Quartic)
	}
}

func TestGate775PotentialTowerAndNumerics(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Potential.Rewritten || a.Potential.UnifiedFactorFormula != "V_local(x)=(C_Higgs/32)(||x||^2-v^2)^2" || a.Potential.NativeHiggsTheorem {
		t.Fatalf("bad potential rewrite: %+v", a.Potential)
	}
	if !a.Tower.Written || a.Tower.MassSquaredFormula != "m_H_tree^2=(C_Higgs/4)v^2" || a.Tower.MassFormula != "m_H_tree=(v/2)sqrt(C_Higgs)" || a.Tower.A2Formula != "A_2=(C_Higgs/8)v^2" || a.Tower.A3Formula != "A_3=(C_Higgs/8)v" || a.Tower.A4Formula != "A_4=C_Higgs/32" || a.Tower.Lambda3Formula != "lambda_3=(3/4)v C_Higgs" || a.Tower.Lambda4Formula != "lambda_4=(3/4)C_Higgs" || a.Tower.PhysicalMeasured || a.Tower.PoleMassTheorem {
		t.Fatalf("bad tower: %+v", a.Tower)
	}
	if !a.Numerical.TowerComputed || !a.Numerical.Finite || !closeRel(a.Numerical.LambdaHBridge, 0.12965256505060754, 1e-15) || !closeRel(a.Numerical.MassGeV, 125.38000000304908, 1e-15) || !closeRel(a.Numerical.A2GeV2, 7860.072200382293, 1e-15) || !closeRel(a.Numerical.A3GeV, 31.923009292084874, 1e-15) || !closeRel(a.Numerical.A4, 0.032413141262651886, 1e-15) || !closeRel(a.Numerical.Lambda3GeV, 191.53805575250925, 1e-15) || !closeRel(a.Numerical.Lambda4, 0.7779153903036453, 1e-15) {
		t.Fatalf("bad numerical ledger: %+v", a.Numerical)
	}
}

func TestGate775SourceTypesAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sources.Recorded || !strings.Contains(a.Sources.CYukawaRole, "Yukawa trace participation") || !strings.Contains(a.Sources.CHistoryRole, "History") || !strings.Contains(a.Sources.BaselineRole, "not a native Higgs theorem") {
		t.Fatalf("bad sources: %+v", a.Sources)
	}
	if !a.Firewalls.Audited || a.Firewalls.CHiggsNativeHiggsTheorem || a.Firewalls.CYukawaNativeYukawaTheorem || a.Firewalls.CHistoryNativeHistoryLoopTheorem || a.Firewalls.CouplingTowerMeasured || a.Firewalls.TreeProxyPoleMass || a.Firewalls.LambdaHBridgeIndependentRuntime || a.Firewalls.YukawaOperatorOrEigenvalue {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate775TheoremStatuses(t *testing.T) {
	res := Generation2UnifiedHiggsCouplingTowerAndFactorizationAuditTheorem().Verify()
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
