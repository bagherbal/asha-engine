package generation2levelbchiggserrorbudgetandindependentinputsensitivityaudit

import (
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

const (
	theoremID   = "GATE-792-LEVEL-B-C-HIGGS-ERROR-BUDGET-INDEPENDENT-INPUT-SENSITIVITY"
	theoremName = "Gate 792 — Level-B Error Budget and Independent-Input Sensitivity Audit"
)

func Generation2LevelBCHiggsErrorBudgetAndIndependentInputSensitivityAuditTheorem() theorem.Theorem {
	return theorem.Theorem{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := Cached()
		if err != nil {
			return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build Gate792 analysis", Passed: false, Detail: err.Error()}}, Notes: []string{StatusFirewallPreservedGate792}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate791 Level-B test interface", Passed: a.Gate791.Inherited && a.Gate791.CleanTestObject && a.Gate791.LevelBOnly, Detail: a.Gate791.Verdict},
			{Name: "compute analytic sensitivity formulas", Passed: a.Sensitivity.Computed && closeAbs(a.Sensitivity.DCByNEff, -0.34547216221380384, 1e-15) && closeAbs(a.Sensitivity.DCByLHopf, 0.9549361341977547, 1e-15) && closeAbs(a.Sensitivity.DCByKappaOrient, 0.039757885839527426, 1e-15) && a.Sensitivity.FormulaNEff != "" && a.Sensitivity.FormulaLHopf != "", Detail: FormatSensitivity(a.Sensitivity)},
			{Name: "complete relative elasticity audit", Passed: a.Elasticity.Completed && a.Elasticity.NEffUnitLeverage && a.Elasticity.LHopfSecondChannel && closeAbs(a.Elasticity.ENEff, -1, 1e-15) && closeAbs(a.Elasticity.ELHopf, 0.03663223082862708, 1e-15) && closeAbs(a.Elasticity.EKappaOrient, 0.00021106412551705377, 1e-18), Detail: FormatElasticity(a.Elasticity)},
			{Name: "compute absolute perturbation ledger", Passed: a.Perturbations.Computed && closeAbs(a.Perturbations.DeltaFromNEff, -3.454721622138038e-7, 1e-21) && closeAbs(a.Perturbations.DeltaFromLHopf, 9.549361341977546e-7, 1e-21) && closeAbs(a.Perturbations.DeltaFromS, -4.036181730287719e-9, 1e-23) && containsAll(a.Perturbations.Strongest, []string{"L_Hopf", "N_eff"}), Detail: FormatPerturbations(a.Perturbations)},
			{Name: "compute component-removal diagnostics", Passed: a.Removal.Computed && a.Removal.YukawaImportant && a.Removal.BoundaryMicroLowNumericalLeverage && closeAbs(a.Removal.TopColorShift, 0.0008046575187645733, 1e-16) && closeAbs(a.Removal.BoundaryCorrectionShift, 1.1036177793855018e-7, 1e-18) && closeAbs(a.Removal.CubicRemovalShift, -1.6224799281872038e-12, 1e-24), Detail: FormatRemoval(a.Removal)},
			{Name: "separate numerical sensitivity from theorem pressure", Passed: a.Pressure.Separated && a.Pressure.NEffTopNumerical && a.Pressure.KappaOrientTopFlavor && a.Pressure.LHopfTopHistory && a.Pressure.FWallLowLeverageButTyped && containsAll(a.Pressure.NumericalRanking, []string{"N_eff", "L_Hopf", "xi_boundary"}) && containsAll(a.Pressure.TheoremPressureRanking, []string{"GenerationMixingOperatorSeal", "Yukawa", "RadialHessianHopfTransportSeal", "BoundaryExteriorResponsePackageSeal"}), Detail: FormatPressure(a.Pressure)},
			{Name: "define error-budget categories", Passed: a.ErrorBudget.Defined && containsAll(a.ErrorBudget.Categories, []string{"Type I", "Type II", "Type III", "Type IV"}), Detail: strings.Join(a.ErrorBudget.Categories, "; ")},
			{Name: "classify major inputs by error type", Passed: a.ErrorBudget.NEffClassified && a.ErrorBudget.KappaOrientClassified && a.ErrorBudget.LHopfClassified && a.ErrorBudget.FWallClassified && a.ErrorBudget.ComparisonClassified && containsAll(a.ErrorBudget.Classifications, []string{"N_eff", "kappa_orient", "L_Hopf", "tree-to-pole"}), Detail: strings.Join(a.ErrorBudget.Classifications, "; ")},
			{Name: "complete scientific testability audit", Passed: a.Testability.Completed && a.Testability.NEffBestNumerical && a.Testability.GenerationHistoryNative && a.Testability.TreeToPoleRequired && a.Testability.NumericalSharpnessTarget == "N_eff" && containsAll(a.Testability.NativeClosureTargets, []string{"GenerationMixingOperatorSeal", "RadialHessianHopfTransportSeal", "Yukawa"}), Detail: a.Testability.NumericalSharpnessTarget},
			{Name: "record next branch recommendation", Passed: a.Next.Recorded && a.Next.Recommended == "Gate 793 — N_eff Yukawa Trace Participation Source and Scale-Stability Audit" && containsAll(a.Next.Alternatives, []string{"GenerationMixingOperator", "Tree-to-Pole", "BoundaryExterior"}) && a.Next.Reason != "", Detail: a.Next.Recommended},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Enforced && !a.Firewalls.ErrorBudgetNativeTheorem && !a.Firewalls.LargestSensitivityDeepestTheorem && !a.Firewalls.KappaOrientSmallMeansUnimportant && !a.Firewalls.FWallLowLeverageDisposable && !a.Firewalls.CHiggsPoleMassPrediction && !a.Firewalls.TreeProxyPoleMass && !a.Firewalls.LevelBLevelCPrediction && !a.Firewalls.NEffNativeYukawa && !a.Firewalls.LHopfNativeHistoryLoop && !a.Firewalls.KappaOrientNativeFlavor && !a.Firewalls.ObservedHiggsSourcesBudget && a.Firewalls.Verdict == StatusFirewallPreservedGate792, Detail: a.Firewalls.Verdict},
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
		notes := append([]string{a.Truth, FormatSensitivity(a.Sensitivity), FormatElasticity(a.Elasticity), FormatPerturbations(a.Perturbations), FormatRemoval(a.Removal), FormatPressure(a.Pressure), a.FinalStatement}, Statuses()...)
		return theorem.Result{ID: theoremID, Name: theoremName, Layer: theorem.LayerBridge, Status: status, Checks: checks, Notes: notes}
	}}
}
