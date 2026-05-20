package generation2chiggsdependencyfreezeandlevelbpredictioninterfaceaudit

import (
	"strings"
	"testing"
)

func TestGate790FrozenInterfaceAndLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate789.Inherited || !a.Gate789.FlavorOrientationSealAccepted || a.Gate789.GenerationMixingOperatorNative {
		t.Fatalf("bad Gate789 inheritance: %+v", a.Gate789)
	}
	if !a.Interface.Written || !strings.Contains(a.Interface.CHiggsFormula, "3/N_eff") || !strings.Contains(a.Interface.ExpandedFWall3, "kappa_orient") || !strings.Contains(a.Interface.ExpandedFWall3, "xi_boundary") {
		t.Fatalf("bad frozen interface: %+v", a.Interface)
	}
	if !closeAbs(a.Ledger.KappaBoundary, kappaBoundarySnapshot, 1e-18) || !closeAbs(a.Ledger.KappaERed, kappaERedSnapshot, 1e-18) || !closeAbs(a.Ledger.FWall3Red, fWall3Snapshot, 1e-18) || !closeAbs(a.Ledger.KappaLambdaRed, kappaLambdaRedSnapshot, 1e-15) || !closeAbs(a.Ledger.CHistory, cHistorySnapshot, 1e-15) || !closeAbs(a.Ledger.CYukawa, cYukawaSnapshot, 1e-15) || !closeAbs(a.Ledger.CHiggs, cHiggsSnapshot, 1e-15) || !closeAbs(a.Ledger.LambdaRuntimeEff, lambdaRuntimeSnapshot, 1e-15) {
		t.Fatalf("bad frozen ledger: %s", FormatLedger(a.Ledger))
	}
}

func TestGate790DependencyClassificationAndRuntimeAbsence(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Classification.Audited || !a.Classification.BoundaryGaugeStrong || !a.Classification.ExplicitSeals {
		t.Fatalf("bad dependency classification: %+v", a.Classification)
	}
	for _, field := range []string{a.Classification.P, a.Classification.KappaOrient, a.Classification.FWall3Red, a.Classification.LHopf, a.Classification.NEff, a.Classification.CHiggs} {
		if field == "" {
			t.Fatalf("empty dependency field in %+v", a.Classification)
		}
	}
	if !a.RuntimeAbsence.Audited || a.RuntimeAbsence.ContainsForbidden || !a.RuntimeAbsence.FormulaIndependent || a.RuntimeAbsence.TheoremIndependent {
		t.Fatalf("bad runtime absence: %+v", a.RuntimeAbsence)
	}
	if !containsAll(a.RuntimeAbsence.ForbiddenVariables, []string{"lambda_runtime", "lambda_runtime_eff", "m_H_tree", "m_H_pole", "G_F", "v"}) {
		t.Fatalf("missing forbidden variables: %+v", a.RuntimeAbsence.ForbiddenVariables)
	}
}

func TestGate790LevelBProtocolSensitivityFreezeAndBranches(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.LevelB.Recorded || !a.LevelB.CleanDimensionless || a.LevelB.LevelCNativePrediction || !strings.Contains(a.LevelB.CHiggs, "Level B") || !strings.Contains(a.LevelB.TreeProxy, "VEV/Fermi") || a.LevelB.PoleMass != "m_H_pole not predicted" {
		t.Fatalf("bad Level-B classification: %+v", a.LevelB)
	}
	if !a.Protocol.Defined || !strings.Contains(a.Protocol.CompareThrough, "4 sqrt(2) G_F") || a.Protocol.ForbiddenCircularInput == "" || !a.Protocol.RequiresCorrection {
		t.Fatalf("bad protocol: %+v", a.Protocol)
	}
	if !a.Sensitivity.Recorded || len(a.Sensitivity.OrderedInputs) != 5 || !containsAll(a.Sensitivity.StructuralBottlenecks, []string{"GenerationMixingOperatorSeal", "Yukawa operator", "HistoryLoop transport", "BoundaryExteriorResponsePackageSeal"}) {
		t.Fatalf("bad sensitivity: %+v", a.Sensitivity)
	}
	if !a.Freeze.Recorded || a.Freeze.Freezes["kappa_orient"] != "FlavorOrientationReadoutSeal" || a.Freeze.Freezes["F_wall_3_red"] != "BoundaryExteriorResponsePackageSeal" || a.Freeze.Freezes["N_eff"] != "YukawaTraceParticipationSeal" || a.Freeze.Freezes["L_Hopf"] != "RadialHessianHopfTransportSeal" {
		t.Fatalf("bad freeze: %+v", a.Freeze)
	}
	if !a.Branches.Recorded || !strings.Contains(a.Branches.Recommended, "Level-B C_Higgs Numerical Interface") || !containsAll(a.Branches.Branches, []string{"Branch A", "Branch B", "Branch C", "Branch D"}) {
		t.Fatalf("bad branches: %+v", a.Branches)
	}
}

func TestGate790FirewallsFinalStatementAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewalls.Enforced || a.Firewalls.CHiggsNativeHiggsTheorem || a.Firewalls.CHiggsPoleMassPrediction || a.Firewalls.LevelBIsLevelCTheorem || a.Firewalls.FlavorOrientationNative || a.Firewalls.NEffNativeYukawa || a.Firewalls.LHopfNativeHistoryLoop || a.Firewalls.FWallNativeBoundary || a.Firewalls.VEVFermiASHAElectroweakScale || a.Firewalls.TreeProxyPoleMass {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	if !strings.Contains(a.FinalStatement, "does not make C_Higgs native") || !strings.Contains(a.FinalStatement, "Level-B dimensionless prediction interface") || !strings.Contains(a.FinalStatement, "Gate 791") {
		t.Fatalf("bad final statement: %s", a.FinalStatement)
	}
	res := Generation2CHiggsDependencyFreezeAndLevelBPredictionInterfaceAuditTheorem().Verify()
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
