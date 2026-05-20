package generation2higgsdimensionlesspredictionindependenceandsealdependencyaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HiggsDimensionlessPredictionIndependenceAndSealDependencyAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 780 — Higgs Dimensionless Prediction Independence and Seal-Dependency Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusGate780HiggsPredictionIndependenceBoundary}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 779 Fermi-normalized ratio", Passed: a.Gate779.Inherited && a.Gate779.FermiNormalizedIdentity == "4sqrt(2)G_F m_H_tree_proxy^2=C_Higgs" && closeRel(a.Gate779.CHiggs, 1.0372205204048603, 1e-15) && closeRel(a.Gate779.VEVGeV, 246.2196508, 1e-15) && closeRel(a.Gate779.GFermiGeVMinus2, 1.1663786999444556e-05, 1e-15) && closeRel(a.Gate779.TreeMassProxyGeV, 125.38000000304908, 1e-15) && closeRel(a.Gate779.FourSqrt2GFMassSquared, 1.0372205204048603, 1e-15) && a.Gate779.DimensionlessTask == "derive or reduce C_Higgs natively" && a.Gate779.ScaleTask == "derive or seal G_F / v" && !a.Gate779.DerivesGF && !a.Gate779.DerivesPoleMass, Detail: FormatGate779(a.Gate779)},
			{Name: "expand C_Higgs dependency graph", Passed: a.Graph.Expanded && a.Graph.CHiggsFormula == "C_Higgs=(3/N_eff)[1+L_Hopf(1-kappa_lambda_red)]" && a.Graph.CYukawaFormula == "C_Yukawa=3/N_eff" && a.Graph.NEffFormula == "N_eff=a^2/b" && a.Graph.CHistoryFormula == "C_History=1+L_Hopf(1-kappa_lambda_red)" && strings.Contains(a.Graph.LHopfFormula, "supp(H_V(x0))") && strings.Contains(a.Graph.KappaLambdaRedFormula, "F_wall_3_red") && strings.Contains(a.Graph.FWall3Formula, "p_K7") && strings.Contains(a.Graph.KappaERedFormula, "theta13") && closeRel(a.Graph.CHiggs, 1.0372205204048603, 1e-15) && closeRel(a.Graph.CYukawa, 0.9992248188812008, 1e-15) && closeRel(a.Graph.CHistory, 1.038025177923625, 1e-15) && closeRel(a.Graph.NEff, 3.0023273474722147, 1e-15) && closeRel(a.Graph.LHopf, 0.039788735772973836, 1e-15) && closeRel(a.Graph.KappaLambdaRed, 0.04432304306956136, 1e-15) && a.Graph.ProductMatchesCHiggs && a.Graph.HistoryMatchesFormula, Detail: FormatGraph(a.Graph)},
			{Name: "audit input independence classification", Passed: a.Classification.Audited && strings.Contains(a.Classification.PK7Classification, "observer-event bridge") && strings.Contains(a.Classification.LHopfClassification, "not a native HistoryLoop theorem") && strings.Contains(a.Classification.NEffClassification, "sealed Yukawa") && strings.Contains(a.Classification.Theta13JCKMClassification, "flavor/empirical") && strings.Contains(a.Classification.BoundaryScalarClassification, "bridge coordinates") && strings.Contains(a.Classification.GFClassification, "external electroweak scale") && strings.Contains(a.Classification.TreeProxyClassification, "tree Hessian proxy") && len(a.Classification.NativeInputs) >= 2 && len(a.Classification.BridgeResponseInputs) >= 4 && len(a.Classification.EmpiricalOrSealedInputs) >= 6 && len(a.Classification.RuntimeDefinedInputs) >= 2 && len(a.Classification.ExternalScaleInputs) >= 2, Detail: FormatClassification(a.Classification)},
			{Name: "define circularity audit", Passed: a.Circularity.Defined && containsAll(a.Circularity.CriticalTargets, []string{"kappa_lambda_red", "lambda_runtime_eff", "lambda_proxy", "F_wall_3_red", "N_eff"}) && a.Circularity.KappaLambdaHistoricallyRuntimeTied && a.Circularity.LambdaRuntimeEffBridgeClosureQuantity && a.Circularity.LambdaProxyFromYukawaLedgerNotHiggsMass && a.Circularity.YukawaLedgerStillSealed && a.Circularity.FWallCompressedAgainstDeficitRelations && a.Circularity.AnyComponentUsesHiggsOrRuntimeTargetData && len(a.Circularity.IndependentPieces) >= 3 && len(a.Circularity.DependentOrSealedPieces) >= 5, Detail: FormatCircularity(a.Circularity)},
			{Name: "define prediction status levels", Passed: a.Levels.Defined && strings.Contains(a.Levels.LevelA, "consistency closure") && strings.Contains(a.Levels.LevelB, "semi-independent") && strings.Contains(a.Levels.LevelC, "independent tree-level prediction") && strings.Contains(a.Levels.LevelD, "physical pole-mass") && strings.Contains(a.Levels.CurrentLevel, "Level A/B") && a.Levels.NotLevelC && a.Levels.NotLevelD && strings.Contains(a.Levels.Reason, "runtime-linked"), Detail: FormatLevels(a.Levels)},
			{Name: "record required removals for prediction", Passed: a.Removals.Recorded && len(a.Removals.Items) == 7 && a.Removals.NeedsNativeYukawaOperator && a.Removals.NeedsScalarMatchingSource && a.Removals.NeedsFlavorSource && a.Removals.NeedsHistoryLoopTheorem && a.Removals.NeedsBoundaryResponseTheorem && a.Removals.NeedsFermiOrVEVScale && a.Removals.NeedsTreeToPolePackage, Detail: FormatRemovals(a.Removals)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.CHiggsIndependentIfTargetDataUsed && !a.Firewalls.FermiRatioPoleMassTheorem && !a.Firewalls.YukawaLedgerNativeYukawaTheorem && !a.Firewalls.KappaLambdaRedNativeScalarTheorem && !a.Firewalls.LHopfNativeHistoryLoopTheorem && !a.Firewalls.GFAShaDerivedScale && !a.Firewalls.TreeProxyPoleMass && a.Firewalls.Verdict == StatusGate780HiggsPredictionIndependenceBoundary, Detail: FormatFirewalls(a.Firewalls)},
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

func containsStatuses(notes []string, wants []string) bool {
	joined := "\x00" + strings.Join(notes, "\x00") + "\x00"
	for _, w := range wants {
		if !strings.Contains(joined, "\x00"+w+"\x00") {
			return false
		}
	}
	return true
}
