package generation2higgspotentialhessianeigenvalueandtreeproxynormalizationaudit

import (
	"math"
	"strings"
	"testing"
)

func near(x, y float64) bool { return math.Abs(x-y) <= 1e-12 }

func TestGate766InheritanceConventionAndVacuumRadius(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate765.Inherited || !strings.Contains(a.Gate765.Carrier, "R^4") || !strings.Contains(a.Gate765.PotentialForm, "phi^dagger phi") || !strings.Contains(a.Gate765.RadialEventType, "rank-one") || !near(a.Gate765.RankOneEventWeight, 0.25) || !near(a.Gate765.LHopf, 1/(8*math.Pi)) || !a.Gate765.CP1FlatnessExpected || a.Gate765.NativePotentialTheorem || a.Gate765.NativeHistoryLoopTheorem {
		t.Fatalf("bad Gate765 inheritance: %+v", a.Gate765)
	}
	if a.Convention.CarrierComplex != "K7+_J(n) ~= C^2" || a.Convention.CarrierReal != "K7+ ~= R^4" || a.Convention.PhiDaggerPhiConvention != "phi^dagger phi = (1/2)||x||^2" || !strings.Contains(a.Convention.RealPotentialFormula, "lambda/4") || !a.Convention.RequiredForTreeComparison || !a.Convention.ConventionSupplied || a.Convention.NativeCoordinateTheorem {
		t.Fatalf("bad real-coordinate convention: %+v", a.Convention)
	}
	if !a.VacuumRadius.LambdaPositive || !a.VacuumRadius.MuSquaredNegative || a.VacuumRadius.StationaryCondition != "mu^2 + lambda ||x||^2 = 0" || !strings.Contains(a.VacuumRadius.RadiusRelation, "-mu^2/lambda") || a.VacuumRadius.VEVConvention != "phi_0^dagger phi_0 = v^2/2" || a.VacuumRadius.NativeVEVTheorem {
		t.Fatalf("bad vacuum-radius relation: %+v", a.VacuumRadius)
	}
}

func TestGate766HessianEigenvalueAndTreeProxy(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Hessian.GradientFormula, "mu^2 x") || !strings.Contains(a.Hessian.HessianFormula, "2lambda x x^T") || !strings.Contains(a.Hessian.VacuumHessianFormula, "2 lambda v^2 P_rad") || a.Hessian.RadialUnitFormula != "u_rad=x_0/v" || a.Hessian.PRadFormula != "P_rad=u_rad u_rad^T" || !a.Hessian.SymbolicComputation {
		t.Fatalf("bad Hessian computation: %+v", a.Hessian)
	}
	if a.EigenAudit.RadialEigenvalueFormula != "2 lambda v^2" || len(a.EigenAudit.AngularEigenvalues) != angularDim || a.EigenAudit.HessianRank != radialRank || a.EigenAudit.RadialProjector != "P_rad" || !a.EigenAudit.PRadSupportsNonzeroMode || !a.EigenAudit.StrengthensGate765 || a.EigenAudit.PhysicalGoldstoneTheorem {
		t.Fatalf("bad eigenvalue audit: %+v", a.EigenAudit)
	}
	if a.TreeProxy.InsertedLambda != "lambda = lambda_runtime_eff" || !near(a.TreeProxy.LambdaRuntimeEff, lambdaRuntimeEff) || !near(a.TreeProxy.VEVGeV, vevGate741GeV) || !near(a.TreeProxy.TreeMassSquaredGeV2, 2*lambdaRuntimeEff*vevGate741GeV*vevGate741GeV) || !near(a.TreeProxy.TreeMassGeV, 125.38000000304908) || !a.TreeProxy.HessianNormalization || a.TreeProxy.PoleMassTheorem {
		t.Fatalf("bad tree proxy relation: %+v", a.TreeProxy)
	}
}

func TestGate766ThreeFactorAndRoleSeparation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.ThreeFactor.MasterScalarFormula, "1/8") || !strings.Contains(a.ThreeFactor.TreeProxyFormula, "v/2") || !near(a.ThreeFactor.CYukawa, cYukawa) || !near(a.ThreeFactor.CHistory, cHistory) || !near(a.ThreeFactor.TotalCorrection, cYukawa*cHistory) || !near(a.ThreeFactor.BaselineVOverTwoGeV, vevGate741GeV/2) || !near(a.ThreeFactor.CorrectionSqrt, math.Sqrt(cYukawa*cHistory)) || !near(a.ThreeFactor.TreeMassGeV, a.TreeProxy.TreeMassGeV) || !a.ThreeFactor.FeedsTreeProxyForm || a.ThreeFactor.IndependentRuntimeProof {
		t.Fatalf("bad three-factor tree proxy: %+v", a.ThreeFactor)
	}
	if !strings.Contains(a.RoleSplit.HistoryLoopRole, "Tr(rho_plus P_rad)") || !strings.Contains(a.RoleSplit.PotentialHessianRole, "H_V") || !a.RoleSplit.SameProjectorSymbol || !near(a.RoleSplit.HistoryLoopWeight, 0.25) || !near(a.RoleSplit.RadialHessianEigenvalueGeV2, a.TreeProxy.TreeMassSquaredGeV2) || !a.RoleSplit.BridgeAlignmentOnly || a.RoleSplit.NativeAlignmentTheorem {
		t.Fatalf("bad role separation: %+v", a.RoleSplit)
	}
}

func TestGate766FirewallsAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewalls.Audited || a.Firewalls.U2PotentialNativeTheorem || a.Firewalls.VNativeTheorem || a.Firewalls.HessianEigenvaluePoleMassTheorem || a.Firewalls.TreeProxyPoleMassTheorem || a.Firewalls.LambdaRuntimeIndependentTheorem || a.Firewalls.SharedPRadNativeAlignmentTheorem || a.Firewalls.RadialHessianFullEWSBTheorem || a.Firewalls.HiggsMassOrPoleMassTheorem || a.Firewalls.YukawaOperatorOrEigenvalueTheorem {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	res := Generation2HiggsPotentialHessianEigenvalueAndTreeProxyNormalizationAuditTheorem().Verify()
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
