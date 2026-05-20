package generation2levelbchiggserrorbudgetandindependentinputsensitivityaudit

import (
	"strings"
	"testing"
)

func TestGate792SensitivityAndElasticity(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate791.Inherited || !a.Gate791.CleanTestObject || !a.Gate791.LevelBOnly {
		t.Fatalf("bad Gate791 inheritance: %+v", a.Gate791)
	}
	if !a.Sensitivity.Computed || !closeAbs(a.Sensitivity.DCByNEff, -0.34547216221380384, 1e-15) || !closeAbs(a.Sensitivity.DCByLHopf, 0.9549361341977547, 1e-15) || !closeAbs(a.Sensitivity.DCByKappaLambdaRed, -0.03975789229626174, 1e-15) || !closeAbs(a.Sensitivity.DCByKappaOrient, 0.039757885839527426, 1e-15) || !closeAbs(a.Sensitivity.DCByS, -0.004036181730287719, 1e-15) || !closeAbs(a.Sensitivity.DCByXiBoundary, 6.456733266535678e-09, 1e-21) {
		t.Fatalf("bad analytic sensitivity: %s", FormatSensitivity(a.Sensitivity))
	}
	if !a.Elasticity.Completed || !a.Elasticity.NEffUnitLeverage || !a.Elasticity.LHopfSecondChannel || !closeAbs(a.Elasticity.ENEff, -1, 1e-15) || !closeAbs(a.Elasticity.ELHopf, 0.03663223082862708, 1e-15) || !closeAbs(a.Elasticity.EAbsLambda, -0.0019050960362564661, 1e-18) || !closeAbs(a.Elasticity.EKappaOrient, 0.00021106412551705377, 1e-18) || !closeAbs(a.Elasticity.ES, -5.029347243414711e-06, 1e-20) || !closeAbs(a.Elasticity.EXiBoundary, 3.1341282342992825e-10, 1e-24) {
		t.Fatalf("bad elasticity: %s", FormatElasticity(a.Elasticity))
	}
}

func TestGate792PerturbationsAndComponentRemoval(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Perturbations.Computed || !containsAll(a.Perturbations.Strongest, []string{"L_Hopf", "N_eff"}) || !closeAbs(a.Perturbations.DeltaFromNEff, -3.454721622138038e-7, 1e-21) || !closeAbs(a.Perturbations.DeltaFromLHopf, 9.549361341977546e-7, 1e-21) || !closeAbs(a.Perturbations.DeltaFromKappaOrient, 3.9757885839527425e-8, 1e-22) || !closeAbs(a.Perturbations.DeltaFromAbsLambda, -3.975789229626174e-8, 1e-22) || !closeAbs(a.Perturbations.DeltaFromS, -4.036181730287719e-9, 1e-23) || !closeAbs(a.Perturbations.DeltaFromXiBoundary, 6.456733266535678e-15, 1e-29) {
		t.Fatalf("bad perturbation ledger: %s", FormatPerturbations(a.Perturbations))
	}
	if !a.Removal.Computed || !a.Removal.YukawaImportant || !a.Removal.BoundaryMicroLowNumericalLeverage || !closeAbs(a.Removal.CHiggsTopColor, 1.038025177923625, 1e-15) || !closeAbs(a.Removal.TopColorShift, 0.0008046575187645733, 1e-16) || !closeAbs(a.Removal.BoundaryCorrectionShift, 1.1036177793855018e-7, 1e-18) || !closeAbs(a.Removal.CubicRemovalShift, -1.6224799281872038e-12, 1e-24) {
		t.Fatalf("bad removal diagnostics: %s", FormatRemoval(a.Removal))
	}
}

func TestGate792PressureErrorBudgetTestabilityAndFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Pressure.Separated || !a.Pressure.NEffTopNumerical || !a.Pressure.KappaOrientTopFlavor || !a.Pressure.LHopfTopHistory || !a.Pressure.FWallLowLeverageButTyped || !containsAll(a.Pressure.NumericalRanking, []string{"N_eff", "L_Hopf", "xi_boundary"}) || !containsAll(a.Pressure.TheoremPressureRanking, []string{"GenerationMixingOperatorSeal", "Yukawa", "RadialHessianHopfTransportSeal", "BoundaryExteriorResponsePackageSeal"}) {
		t.Fatalf("bad pressure split: %s", FormatPressure(a.Pressure))
	}
	if !a.ErrorBudget.Defined || !a.ErrorBudget.NEffClassified || !a.ErrorBudget.KappaOrientClassified || !a.ErrorBudget.LHopfClassified || !a.ErrorBudget.FWallClassified || !a.ErrorBudget.ComparisonClassified || !containsAll(a.ErrorBudget.Categories, []string{"Type I", "Type II", "Type III", "Type IV"}) || !containsAll(a.ErrorBudget.Classifications, []string{"N_eff", "kappa_orient", "L_Hopf", "tree-to-pole"}) {
		t.Fatalf("bad error budget: %+v", a.ErrorBudget)
	}
	if !a.Testability.Completed || !a.Testability.NEffBestNumerical || !a.Testability.GenerationHistoryNative || !a.Testability.TreeToPoleRequired || a.Testability.NumericalSharpnessTarget != "N_eff" || !strings.Contains(a.Testability.EmpiricalComparisonTarget, "tree-to-pole") {
		t.Fatalf("bad testability: %+v", a.Testability)
	}
	if !a.Next.Recorded || a.Next.Recommended != "Gate 793 — N_eff Yukawa Trace Participation Source and Scale-Stability Audit" || a.Next.Reason == "" || !containsAll(a.Next.Alternatives, []string{"GenerationMixingOperator", "Tree-to-Pole", "BoundaryExterior"}) {
		t.Fatalf("bad next recommendation: %+v", a.Next)
	}
	if !a.Firewalls.Enforced || a.Firewalls.ErrorBudgetNativeTheorem || a.Firewalls.LargestSensitivityDeepestTheorem || a.Firewalls.KappaOrientSmallMeansUnimportant || a.Firewalls.FWallLowLeverageDisposable || a.Firewalls.CHiggsPoleMassPrediction || a.Firewalls.TreeProxyPoleMass || a.Firewalls.LevelBLevelCPrediction || a.Firewalls.NEffNativeYukawa || a.Firewalls.LHopfNativeHistoryLoop || a.Firewalls.KappaOrientNativeFlavor || a.Firewalls.ObservedHiggsSourcesBudget {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
}

func TestGate792TheoremStatusesAndFinalStatement(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.FinalStatement, "N_eff is the highest numerical leverage") || !strings.Contains(a.FinalStatement, "largest theorem risks") || !strings.Contains(a.FinalStatement, "Gate 793") {
		t.Fatalf("bad final statement: %s", a.FinalStatement)
	}
	res := Generation2LevelBCHiggsErrorBudgetAndIndependentInputSensitivityAuditTheorem().Verify()
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
