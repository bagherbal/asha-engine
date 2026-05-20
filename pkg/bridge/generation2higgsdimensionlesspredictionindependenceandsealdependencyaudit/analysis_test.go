package generation2higgsdimensionlesspredictionindependenceandsealdependencyaudit

import (
	"strings"
	"testing"
)

func TestGate780InheritanceAndGraph(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate779.Inherited || a.Gate779.FermiNormalizedIdentity != "4sqrt(2)G_F m_H_tree_proxy^2=C_Higgs" || !closeRel(a.Gate779.CHiggs, 1.0372205204048603, 1e-15) || !closeRel(a.Gate779.GFermiGeVMinus2, 1.1663786999444556e-05, 1e-15) || !closeRel(a.Gate779.TreeMassProxyGeV, 125.38000000304908, 1e-15) || a.Gate779.DerivesGF || a.Gate779.DerivesPoleMass {
		t.Fatalf("bad Gate779 inheritance: %+v", a.Gate779)
	}
	if !a.Graph.Expanded || a.Graph.CHiggsFormula != "C_Higgs=(3/N_eff)[1+L_Hopf(1-kappa_lambda_red)]" || a.Graph.NEffFormula != "N_eff=a^2/b" || !strings.Contains(a.Graph.LHopfFormula, "supp(H_V(x0))") || !strings.Contains(a.Graph.KappaERedFormula, "J_CKM") || !closeRel(a.Graph.CYukawa*a.Graph.CHistory, a.Graph.CHiggs, 1e-15) || !a.Graph.ProductMatchesCHiggs || !a.Graph.HistoryMatchesFormula {
		t.Fatalf("bad dependency graph: %+v", a.Graph)
	}
}

func TestGate780ClassificationAndCircularity(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Classification.Audited || !strings.Contains(a.Classification.LHopfClassification, "not a native HistoryLoop theorem") || !strings.Contains(a.Classification.NEffClassification, "sealed Yukawa") || !strings.Contains(a.Classification.GFClassification, "external electroweak scale") || len(a.Classification.BridgeResponseInputs) < 4 || len(a.Classification.EmpiricalOrSealedInputs) < 6 || len(a.Classification.ExternalScaleInputs) != 2 {
		t.Fatalf("bad classification: %+v", a.Classification)
	}
	if !a.Circularity.Defined || !a.Circularity.KappaLambdaHistoricallyRuntimeTied || !a.Circularity.LambdaRuntimeEffBridgeClosureQuantity || !a.Circularity.LambdaProxyFromYukawaLedgerNotHiggsMass || !a.Circularity.YukawaLedgerStillSealed || !a.Circularity.FWallCompressedAgainstDeficitRelations || !a.Circularity.AnyComponentUsesHiggsOrRuntimeTargetData || !containsAll(a.Circularity.CriticalTargets, []string{"kappa_lambda_red", "lambda_runtime_eff", "lambda_proxy", "F_wall_3_red", "N_eff"}) {
		t.Fatalf("bad circularity audit: %+v", a.Circularity)
	}
}

func TestGate780LevelsRemovalsAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Levels.Defined || !strings.Contains(a.Levels.CurrentLevel, "Level A/B") || !a.Levels.NotLevelC || !a.Levels.NotLevelD || !strings.Contains(a.Levels.Reason, "runtime-linked") {
		t.Fatalf("bad levels: %+v", a.Levels)
	}
	if !a.Removals.Recorded || len(a.Removals.Items) != 7 || !a.Removals.NeedsNativeYukawaOperator || !a.Removals.NeedsScalarMatchingSource || !a.Removals.NeedsFlavorSource || !a.Removals.NeedsHistoryLoopTheorem || !a.Removals.NeedsBoundaryResponseTheorem || !a.Removals.NeedsFermiOrVEVScale || !a.Removals.NeedsTreeToPolePackage {
		t.Fatalf("bad removals: %+v", a.Removals)
	}
	if !a.Firewalls.Enforced || a.Firewalls.CHiggsIndependentIfTargetDataUsed || a.Firewalls.FermiRatioPoleMassTheorem || a.Firewalls.YukawaLedgerNativeYukawaTheorem || a.Firewalls.KappaLambdaRedNativeScalarTheorem || a.Firewalls.LHopfNativeHistoryLoopTheorem || a.Firewalls.GFAShaDerivedScale || a.Firewalls.TreeProxyPoleMass {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate780TheoremStatuses(t *testing.T) {
	res := Generation2HiggsDimensionlessPredictionIndependenceAndSealDependencyAuditTheorem().Verify()
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
