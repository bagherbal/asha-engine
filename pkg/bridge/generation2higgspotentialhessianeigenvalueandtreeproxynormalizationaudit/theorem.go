package generation2higgspotentialhessianeigenvalueandtreeproxynormalizationaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2HiggsPotentialHessianEigenvalueAndTreeProxyNormalizationAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 766 — Higgs Potential Hessian Eigenvalue and Tree-Proxy Normalization Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default audit", Passed: false, Detail: err.Error()}}}
		}
		near := func(x, y, tol float64) bool { return math.Abs(x-y) <= tol }
		checks := []theorem.Check{
			{Name: "inherit Gate765 Higgs potential radial event typing", Passed: a.Gate765.Inherited && strings.Contains(a.Gate765.Carrier, "R^4") && strings.Contains(a.Gate765.PotentialForm, "mu^2") && strings.Contains(a.Gate765.RadialEventType, "rank-one") && near(a.Gate765.RankOneEventWeight, 0.25, 1e-15) && near(a.Gate765.LHopf, 1/(8*math.Pi), 1e-15) && a.Gate765.CP1FlatnessExpected && !a.Gate765.NativePotentialTheorem && !a.Gate765.NativeHistoryLoopTheorem, Detail: FormatGate765(a.Gate765)},
			{Name: "define real four-coordinate convention", Passed: a.Convention.CarrierComplex == "K7+_J(n) ~= C^2" && a.Convention.CarrierReal == "K7+ ~= R^4" && a.Convention.PhiDaggerPhiConvention == "phi^dagger phi = (1/2)||x||^2" && strings.Contains(a.Convention.RealPotentialFormula, "lambda/4") && a.Convention.RequiredForTreeComparison && a.Convention.ConventionSupplied && !a.Convention.NativeCoordinateTheorem, Detail: FormatConvention(a.Convention)},
			{Name: "record convention-dependent vacuum radius relation", Passed: a.VacuumRadius.LambdaPositive && a.VacuumRadius.MuSquaredNegative && a.VacuumRadius.StationaryCondition == "mu^2 + lambda ||x||^2 = 0" && strings.Contains(a.VacuumRadius.RadiusRelation, "v^2") && a.VacuumRadius.VEVConvention == "phi_0^dagger phi_0 = v^2/2" && !a.VacuumRadius.NativeVEVTheorem && strings.Contains(a.VacuumRadius.Verdict, StatusNoNativeVEVTheorem), Detail: FormatVacuumRadius(a.VacuumRadius)},
			{Name: "compute gradient and Hessian of supplied potential", Passed: strings.Contains(a.Hessian.PotentialFormula, "lambda/4") && strings.Contains(a.Hessian.GradientFormula, "mu^2 x") && strings.Contains(a.Hessian.HessianFormula, "2lambda x x^T") && strings.Contains(a.Hessian.VacuumCondition, "mu^2=-lambda v^2") && strings.Contains(a.Hessian.VacuumHessianFormula, "2 lambda v^2 P_rad") && a.Hessian.RadialUnitFormula == "u_rad=x_0/v" && a.Hessian.PRadFormula == "P_rad=u_rad u_rad^T" && a.Hessian.SymbolicComputation, Detail: FormatHessian(a.Hessian)},
			{Name: "audit radial Hessian eigenvalue and support", Passed: a.EigenAudit.RadialEigenvalueFormula == "2 lambda v^2" && len(a.EigenAudit.AngularEigenvalues) == angularDim && a.EigenAudit.HessianRank == radialRank && a.EigenAudit.RadialProjector == "P_rad" && a.EigenAudit.PRadSupportsNonzeroMode && a.EigenAudit.StrengthensGate765 && !a.EigenAudit.PhysicalGoldstoneTheorem && strings.Contains(a.EigenAudit.Verdict, StatusPRadSupportOfRadialHessianEigenvalue), Detail: FormatEigenAudit(a.EigenAudit)},
			{Name: "reconstruct tree proxy relation from Hessian normalization", Passed: a.TreeProxy.InsertedLambda == "lambda = lambda_runtime_eff" && near(a.TreeProxy.LambdaRuntimeEff, lambdaRuntimeEff, 1e-15) && near(a.TreeProxy.VEVGeV, vevGate741GeV, 1e-12) && near(a.TreeProxy.TreeMassSquaredGeV2, 2*lambdaRuntimeEff*vevGate741GeV*vevGate741GeV, 1e-9) && near(a.TreeProxy.TreeMassGeV, 125.38000000304908, 1e-9) && strings.Contains(a.TreeProxy.RelationFormula, "2 lambda_runtime_eff v^2") && a.TreeProxy.HessianNormalization && !a.TreeProxy.PoleMassTheorem, Detail: FormatTreeProxy(a.TreeProxy)},
			{Name: "write three-factor tree proxy form", Passed: strings.Contains(a.ThreeFactor.MasterScalarFormula, "3/N_eff") && strings.Contains(a.ThreeFactor.TreeProxyFormula, "v/2") && near(a.ThreeFactor.CYukawa, cYukawa, 1e-15) && near(a.ThreeFactor.CHistory, cHistory, 1e-15) && near(a.ThreeFactor.TotalCorrection, cYukawa*cHistory, 1e-15) && near(a.ThreeFactor.BaselineVOverTwoGeV, vevGate741GeV/2, 1e-12) && near(a.ThreeFactor.CorrectionSqrt, math.Sqrt(cYukawa*cHistory), 1e-15) && near(a.ThreeFactor.TreeMassGeV, a.TreeProxy.TreeMassGeV, 1e-9) && a.ThreeFactor.FeedsTreeProxyForm && !a.ThreeFactor.IndependentRuntimeProof, Detail: FormatThreeFactor(a.ThreeFactor)},
			{Name: "separate HistoryLoop and Hessian uses of P_rad", Passed: strings.Contains(a.RoleSplit.HistoryLoopRole, "1/(8*pi)") && strings.Contains(a.RoleSplit.PotentialHessianRole, "2 lambda v^2 P_rad") && a.RoleSplit.SameProjectorSymbol && near(a.RoleSplit.HistoryLoopWeight, 0.25, 1e-15) && near(a.RoleSplit.RadialHessianEigenvalueGeV2, a.TreeProxy.TreeMassSquaredGeV2, 1e-9) && a.RoleSplit.BridgeAlignmentOnly && !a.RoleSplit.NativeAlignmentTheorem, Detail: FormatRoleSplit(a.RoleSplit)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.U2PotentialNativeTheorem && !a.Firewalls.VNativeTheorem && !a.Firewalls.HessianEigenvaluePoleMassTheorem && !a.Firewalls.TreeProxyPoleMassTheorem && !a.Firewalls.LambdaRuntimeIndependentTheorem && !a.Firewalls.SharedPRadNativeAlignmentTheorem && !a.Firewalls.RadialHessianFullEWSBTheorem && !a.Firewalls.HiggsMassOrPoleMassTheorem && !a.Firewalls.YukawaOperatorOrEigenvalueTheorem && strings.Contains(a.Firewalls.Verdict, StatusGate766HiggsHessianTreeProxyBoundary), Detail: FormatFirewalls(a.Firewalls)},
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
