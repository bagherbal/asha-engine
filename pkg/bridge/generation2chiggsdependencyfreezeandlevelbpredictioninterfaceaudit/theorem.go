package generation2chiggsdependencyfreezeandlevelbpredictioninterfaceaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

const (
	theoremID   = "GATE-790-C-HIGGS-DEPENDENCY-FREEZE-LEVEL-B-PREDICTION-INTERFACE"
	theoremName = "Gate 790 — C_Higgs Dependency Freeze and Level-B Prediction Interface Audit"
)

func Generation2CHiggsDependencyFreezeAndLevelBPredictionInterfaceAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := Cached()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build analysis", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate789 failure branch", Passed: a.Gate789.Inherited && a.Gate789.FlavorOrientationSealAccepted && !a.Gate789.GenerationMixingOperatorNative, Detail: a.Gate789.InheritedSeal},
			{Name: "write frozen C_Higgs Level-B interface", Passed: a.Interface.Written && a.Interface.CHiggsFormula != "" && a.Interface.ExpandedFWall3 != "" && a.Interface.KappaERed == "kappa_e_red = kappa_orient + kappa_boundary", Detail: a.Interface.CHiggsFormula},
			{Name: "expand kappa_lambda_red into boundary-flavor components", Passed: a.Interface.KappaLambdaRed != "" && a.Interface.FWall3Red != "" && a.Interface.KappaOrient != "" && a.Interface.KappaBoundary != "", Detail: a.Interface.ExpandedFWall3},
			{Name: "recompute frozen numerical ledger", Passed: a.Ledger.Recomputed && closeAbs(a.Ledger.CHiggs, cHiggsSnapshot, 1e-15) && closeAbs(a.Ledger.CHistory, cHistorySnapshot, 1e-15) && closeAbs(a.Ledger.LambdaRuntimeEff, lambdaRuntimeSnapshot, 1e-15), Detail: FormatLedger(a.Ledger)},
			{Name: "classify active dependencies", Passed: a.Classification.Audited && a.Classification.BoundaryGaugeStrong && a.Classification.ExplicitSeals && a.Classification.CHiggs != "" && a.Classification.NEff != "" && a.Classification.LHopf != "", Detail: a.Classification.CHKappaSummary()},
			{Name: "audit runtime target absence", Passed: a.RuntimeAbsence.Audited && containsAll(a.RuntimeAbsence.ForbiddenVariables, []string{"lambda_runtime", "lambda_runtime_eff", "m_H_tree", "m_H_pole", "G_F", "v"}) && !a.RuntimeAbsence.ContainsForbidden && a.RuntimeAbsence.FormulaIndependent && !a.RuntimeAbsence.TheoremIndependent, Detail: a.RuntimeAbsence.Verdict},
			{Name: "record Level-B classification", Passed: a.LevelB.Recorded && a.LevelB.CleanDimensionless && !a.LevelB.LevelCNativePrediction && a.LevelB.PoleMass == "m_H_pole not predicted", Detail: a.LevelB.CiggsSummary()},
			{Name: "define Level-B test protocol", Passed: a.Protocol.Defined && containsAll(a.Protocol.AllowedInputs, []string{"N_eff", "kappa_orient", "boundary coordinates", "p", "L_Hopf"}) && a.Protocol.ForbiddenCircularInput != "" && a.Protocol.RequiresCorrection, Detail: a.Protocol.CompareThrough},
			{Name: "record sensitivity/source pressure order", Passed: a.Sensitivity.Recorded && len(a.Sensitivity.OrderedInputs) == 5 && containsAll(a.Sensitivity.StructuralBottlenecks, []string{"GenerationMixingOperatorSeal", "Yukawa operator", "HistoryLoop transport", "BoundaryExteriorResponsePackageSeal"}), Detail: a.Sensitivity.OrderedInputs[0]},
			{Name: "record seal freeze decision", Passed: a.Freeze.Recorded && a.Freeze.Freezes["kappa_orient"] == "FlavorOrientationReadoutSeal" && a.Freeze.Freezes["F_wall_3_red"] == "BoundaryExteriorResponsePackageSeal" && a.Freeze.Freezes["N_eff"] == "YukawaTraceParticipationSeal" && a.Freeze.Freezes["L_Hopf"] == "RadialHessianHopfTransportSeal", Detail: FormatFreeze(a.Freeze)},
			{Name: "record next branch options", Passed: a.Branches.Recorded && containsAll(a.Branches.Branches, []string{"Branch A", "Branch B", "Branch C", "Branch D"}) && a.Branches.Recommended != "" && a.Branches.Verdict == StatusNextBranchOptionsRecorded, Detail: FormatBranches(a.Branches)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.CHiggsNativeHiggsTheorem && !a.Firewalls.CHiggsPoleMassPrediction && !a.Firewalls.LevelBIsLevelCTheorem && !a.Firewalls.FlavorOrientationNative && !a.Firewalls.NEffNativeYukawa && !a.Firewalls.LHopfNativeHistoryLoop && !a.Firewalls.FWallNativeBoundary && !a.Firewalls.VEVFermiASHAElectroweakScale && !a.Firewalls.TreeProxyPoleMass && a.Firewalls.Verdict == StatusFirewallPreservedGate790, Detail: a.Firewalls.Verdict},
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
		notes := append([]string{a.Truth, FormatLedger(a.Ledger), FormatFreeze(a.Freeze), FormatBranches(a.Branches), a.FinalStatement}, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}

func (d DependencyClassification) CHKappaSummary() string {
	return d.KappaOrient + "; " + d.FWall3Red + "; " + d.NEff + "; " + d.LHopf
}

func (l LevelBClassification) CiggsSummary() string {
	return l.CHiggs + "; " + l.LambdaRuntimeEff + "; " + l.TreeProxy + "; " + l.PoleMass
}
