package generation2higgstotalcorrectiondecompositionanddilationaudit

import (
	"strings"
	"testing"
)

func TestGate776DefinitionsAndExpansion(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate775.Inherited || !a.Gate775.UnifiedTower || a.Gate775.NativeHiggsTheorem || a.Gate775.NativeYukawaTheorem || a.Gate775.NativeHistoryTheorem || a.Gate775.PhysicalSelfCouplings {
		t.Fatalf("bad Gate775 inheritance: %+v", a.Gate775)
	}
	if !a.Definitions.Defined || a.Definitions.DeltaHistoryFormula != "delta_History=L_Hopf(1-kappa_lambda_red)" || a.Definitions.EpsilonYukawaFormula != "epsilon_Yukawa=1-3/N_eff" || a.Definitions.CHistoryFormula != "C_History=1+delta_History" || a.Definitions.CYukawaFormula != "C_Yukawa=1-epsilon_Yukawa" || a.Definitions.NativeHistoryTheorem || a.Definitions.NativeYukawaTheorem {
		t.Fatalf("bad definitions: %+v", a.Definitions)
	}
	if !a.Expansion.Computed || a.Expansion.CHiggsFormula != "C_Higgs=(1-epsilon_Yukawa)(1+delta_History)" || a.Expansion.DeltaHiggsFormula != "Delta_Higgs=C_Higgs-1=delta_History-epsilon_Yukawa(1+delta_History)" || a.Expansion.NativeHiggsTheorem || a.Expansion.IndependentRuntimeTheorem {
		t.Fatalf("bad expansion: %+v", a.Expansion)
	}
}

func TestGate776NumericalDecompositionAndDilation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Numerical.Recorded || !a.Numerical.Finite || !closeRel(a.Numerical.DeltaHistory, 0.03802517792362492, 1e-14) || !closeRel(a.Numerical.EpsilonYukawa, 0.0007751811187991509, 1e-15) || !closeRel(a.Numerical.YukawaDrag, 0.0008046575187645232, 1e-15) || !closeRel(a.Numerical.DeltaHiggs, 0.03722052040486035, 1e-15) || !closeRel(a.Numerical.CHiggs, 1.0372205204048603, 1e-15) || !closeAbs(a.Numerical.ExpansionResidual, 0, 1e-15) || !a.Numerical.HistoryDominates {
		t.Fatalf("bad numerical ledger: %+v", a.Numerical)
	}
	if !a.Dilation.Computed || a.Dilation.Formula != "D_radial=sqrt(C_Higgs)" || !closeRel(a.Dilation.DilationFactor, 1.0184402389953279, 1e-15) || !closeRel(a.Dilation.BaselineMassGeV, 123.1098254, 1e-15) || !closeRel(a.Dilation.MassGeV, 125.38000000304908, 1e-15) || a.Dilation.PoleMassCorrection {
		t.Fatalf("bad dilation: %+v", a.Dilation)
	}
}

func TestGate776TowerSourcesAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Tower.Rewritten || a.Tower.LambdaHFormula != "lambda_H_bridge=(1/8)(1+Delta_Higgs)" || a.Tower.PotentialFormula != "V_local(x)=[(1+Delta_Higgs)/32](||x||^2-v^2)^2" || a.Tower.MassFormula != "m_H_tree=(v/2)sqrt(1+Delta_Higgs)" || a.Tower.Lambda3Formula != "lambda_3=(3/4)v(1+Delta_Higgs)" || a.Tower.Lambda4Formula != "lambda_4=(3/4)(1+Delta_Higgs)" || a.Tower.PhysicalSelfCoupling || a.Tower.PoleMassTheorem {
		t.Fatalf("bad tower: %+v", a.Tower)
	}
	if !a.Sources.Recorded || !strings.Contains(a.Sources.DeltaHistoryRole, "History") || !strings.Contains(a.Sources.EpsilonYukawaRole, "Yukawa") || !strings.Contains(a.Sources.YukawaDragRole, "multiplicative drag") || !strings.Contains(a.Sources.TowerRole, "square-root dilation") {
		t.Fatalf("bad sources: %+v", a.Sources)
	}
	if !a.Firewalls.Audited || a.Firewalls.DeltaHiggsNativeHiggsTheorem || a.Firewalls.HistoryUpliftNativeHistoryLoop || a.Firewalls.YukawaDilutionNativeYukawa || a.Firewalls.RadialDilationPoleMassCorrection || a.Firewalls.TreeTowerMeasuredSelfCouplings || a.Firewalls.YukawaOperatorOrEigenvalue {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate776TheoremStatuses(t *testing.T) {
	res := Generation2HiggsTotalCorrectionDecompositionAndDilationAuditTheorem().Verify()
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
