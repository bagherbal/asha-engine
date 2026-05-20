package generation2threefactorscalarhiggsmasternormalformandremainingsealpriorityaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate760InheritanceMasterFormulaAndLedger(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate759.Inherited || !a.Gate759.ThreeFactorFormAvailable || a.Gate759.IndependentScalarRuntimeTheorem {
		t.Fatalf("bad Gate759 inheritance: %+v", a.Gate759)
	}
	if math.Abs(a.Gate759.KappaLambdaRed-kappaLambdaRedMZ) > 1e-12 || math.Abs(a.Gate759.CHistory-1.038025177923625) > 1e-12 || math.Abs(a.Gate759.LambdaRuntimeEff-lambdaRuntimeEffMZ) > 1e-15 {
		t.Fatalf("bad Gate759 numerics: %+v", a.Gate759)
	}
	if !a.Master.Defined || a.Master.IndependentScalarRuntimeTheorem {
		t.Fatalf("bad master formula typing: %+v", a.Master)
	}
	if math.Abs(a.Master.CBaseline-oneEighth) > 1e-18 || math.Abs(a.Master.CYukawa-cYukawaMZ) > 1e-15 || math.Abs(a.Master.CHistory-1.038025177923625) > 1e-12 || math.Abs(a.Master.TotalCorrection-1.0372205204048603) > 1e-12 || math.Abs(a.Master.LambdaRuntimeFromMaster-lambdaRuntimeEffMZ) > 1e-15 || math.Abs(a.Master.MasterResidual) > 1e-15 {
		t.Fatalf("bad master formula numerics: %+v", a.Master)
	}
	if !a.Ledger.Recorded || !a.Ledger.Finite || math.Abs(a.Ledger.NEff-nEffMZ) > 1e-15 || math.Abs(a.Ledger.KappaLambdaComplement-0.9556769569304386) > 1e-12 {
		t.Fatalf("bad ledger: %+v", a.Ledger)
	}
}

func TestGate760SourceTypesAndKappaExpansion(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sources.Audited || a.Sources.BaselineScalarTheorem || a.Sources.NEffNativeYukawaTheorem || a.Sources.LHopfNativeTheorem || a.Sources.KappaNativeTheorem {
		t.Fatalf("bad source-type firewall: %+v", a.Sources)
	}
	for _, want := range []string{"top-color", "N_eff", "P_rad", "F_wall_3_red"} {
		joined := strings.Join([]string{a.Sources.BaselineSourceType, a.Sources.CYukawaSourceType, a.Sources.LHopfSourceType, a.Sources.KappaLambdaRedSourceType}, "\n")
		if !strings.Contains(joined, want) {
			t.Fatalf("missing source marker %q in %s", want, joined)
		}
	}
	if !a.KappaExpansion.ExpansionRecorded || a.KappaExpansion.PrimitiveInCurrentBridge || !a.KappaExpansion.ReconstructedFromWallFlavor || a.KappaExpansion.NativeScalarTheorem || a.KappaExpansion.BoundaryGeneratingFunction {
		t.Fatalf("bad kappa expansion firewall: %+v", a.KappaExpansion)
	}
	if !strings.Contains(a.KappaExpansion.Definition, "F_wall_3_red") || !strings.Contains(a.KappaExpansion.FWall3RedFormula, "p_K7") || !strings.Contains(a.KappaExpansion.KappaERedFormula, "J_CKM") {
		t.Fatalf("missing kappa expansion formula content: %+v", a.KappaExpansion)
	}
}

func TestGate760SealPriorityAndFirewalls(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.SealPriority.Audited || !a.SealPriority.Ordered || len(a.SealPriority.Priorities) != 6 {
		t.Fatalf("bad priority audit: %+v", a.SealPriority)
	}
	wantOrder := []string{"P_rad", "n", "N_eff", "kappa_e_red", "F_wall_3_red", "q"}
	for i, want := range wantOrder {
		if a.SealPriority.Priorities[i].Rank != i+1 || a.SealPriority.Priorities[i].Symbol != want {
			t.Fatalf("bad priority order at %d: want %s got %+v", i, want, a.SealPriority.Priorities[i])
		}
	}
	if a.SealPriority.ScalarReductionTarget != "P_rad / L_Hopf" || a.SealPriority.FlavorYukawaReductionTarget != "N_eff" || a.SealPriority.BoundaryReductionTarget != "F_wall_3_red" || a.SealPriority.NativePradSelector || a.SealPriority.NativeBoundaryGenerator {
		t.Fatalf("bad reduction target audit: %+v", a.SealPriority)
	}
	if !a.Firewalls.Audited || a.Firewalls.ThreeFactorIndependentRuntimeTheorem || a.Firewalls.NEffNativeYukawaTheorem || a.Firewalls.LHopfNativeHistoryLoopTheorem || a.Firewalls.KappaLambdaRedNativeScalarTheorem || a.Firewalls.TreeProxyPoleMass || a.Firewalls.HiggsSocketSealsHiggsMassTheorem || a.Firewalls.YukawaOperatorOrEigenvalueTheorem || a.Firewalls.HiggsMassOrPoleMassTheorem {
		t.Fatalf("bad physical firewalls: %+v", a.Firewalls)
	}
}

func TestGate760TheoremVerdictStatuses(t *testing.T) {
	res := Generation2ThreeFactorScalarHiggsMasterNormalFormAndRemainingSealPriorityAuditTheorem().Verify()
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
