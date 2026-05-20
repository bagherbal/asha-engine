package generation2u2invarianthiggspotentialradialhessianandrankoneeventaudit

import (
	"math"
	"strings"

	"github.com/bagherbal/asha-engine/pkg/theorem"
)

func Generation2U2InvariantHiggsPotentialRadialHessianAndRankOneEventAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Gate 765 — U(2)-Invariant Higgs Potential Radial Hessian and Rank-One Event Audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "build default audit", Passed: false, Detail: err.Error()}}}
		}
		near := func(x, y, tol float64) bool { return math.Abs(x-y) <= tol }
		checks := []theorem.Check{
			{Name: "inherit Gate764 CP1 gauge-orbit radial-rank result", Passed: a.Gate764.Inherited && a.Gate764.Carrier == "K7+_J(n) ~= C^2" && strings.Contains(a.Gate764.RhoPlus, "I_K7+") && strings.Contains(a.Gate764.PRad, "rank-one") && strings.Contains(a.Gate764.LHopfFormula, "1/(8*pi)") && a.Gate764.CP1LocationGaugeRepresentative && a.Gate764.LHopfDependsOnRankOneRadialEvent && !a.Gate764.ComplexVacuumLineScalarSource && strings.Contains(a.Gate764.RemainingScalarSourceQuestion, "rank-one radial"), Detail: FormatGate764(a.Gate764)},
			{Name: "audit standard U2-invariant Higgs potential form", Passed: strings.Contains(a.Potential.PotentialFormula, "mu^2") && strings.Contains(a.Potential.PotentialFormula, "lambda") && a.Potential.RadiusVariable == "r^2 = phi^dagger phi" && a.Potential.DependsOnlyOnRadius && a.Potential.Carrier == "K7+_J(n) ~= C^2" && a.Potential.U2Invariant && a.Potential.ScalarPotentialSupplied && !a.Potential.NativePotentialTheorem && strings.Contains(a.Potential.Verdict, StatusNoNativeASHAScalarPotentialTheorem), Detail: FormatPotential(a.Potential)},
			{Name: "audit CP1 flatness of U2-invariant potential", Passed: strings.Contains(a.CP1Flatness.FixedRadiusOrbit, "CP1") && a.CP1Flatness.PotentialConstantOnCP1 && !a.CP1Flatness.SelectsPiVacC && a.CP1Flatness.ConfirmsGate764Demotion && a.CP1Flatness.RequiresAnisotropyForLine && strings.Contains(a.CP1Flatness.Verdict, StatusCP1SelectorAbsenceExpectedForU2InvariantPotential), Detail: FormatCP1Flatness(a.CP1Flatness)},
			{Name: "record convention-dependent vacuum-radius relation", Passed: a.VacuumRadius.LambdaPositive && a.VacuumRadius.MuSquaredNegative && strings.Contains(a.VacuumRadius.StationaryCondition, "mu^2+2 lambda") && a.VacuumRadius.PhiDaggerPhiRelation == "phi^dagger phi = -mu^2/(2 lambda)" && a.VacuumRadius.VEVConvention == "phi^dagger phi = v^2/2" && a.VacuumRadius.VSquaredRelation == "v^2 = -mu^2/lambda" && a.VacuumRadius.ConventionDependent && !a.VacuumRadius.NativeVEVTheorem && strings.Contains(a.VacuumRadius.Verdict, StatusNoNativeVEVTheorem), Detail: FormatVacuumRadius(a.VacuumRadius)},
			{Name: "type radial Hessian direction and one-plus-three split", Passed: strings.Contains(a.RadialHessian.SuppliedVacuumRepresentative, "phi_0") && a.RadialHessian.RadialPath == "phi(t)=(1+t)phi_0" && a.RadialHessian.DefinesPRad && a.RadialHessian.RadialRealDimension == realRadialRank && a.RadialHessian.AngularRealDimension == angularRealDim && a.RadialHessian.K7PlusRealDimension == k7PlusRealDim && a.RadialHessian.AngularPreservesRadiusFirstOrder && strings.Contains(a.RadialHessian.SplitFormula, "1 + 3") && a.RadialHessian.PotentialHessianSplit && !a.RadialHessian.PhysicalGoldstoneTheorem, Detail: FormatRadialHessian(a.RadialHessian)},
			{Name: "compute rank-one radial event weight", Passed: strings.Contains(a.RankWeight.RhoPlusFormula, "rank(P_rad)/4") && a.RankWeight.RadialRank == realRadialRank && a.RankWeight.K7PlusRealDimension == k7PlusRealDim && near(a.RankWeight.RadialEventWeight, 0.25, 1e-15) && near(a.RankWeight.PhaseLoopPayoff, 1/(2*math.Pi), 1e-15) && near(a.RankWeight.LHopf, 1/(8*math.Pi), 1e-15) && near(a.RankWeight.LHopf, a.RankWeight.ExpectedLHopf, 1e-15) && a.RankWeight.MatchesHistoryLoopQuarter && !a.RankWeight.NativeHistoryLoopTheorem, Detail: FormatRankWeight(a.RankWeight)},
			{Name: "audit complex line versus radial Hessian event", Passed: a.LineVsRadial.ComplexLineRealRank == complexLineRealRank && a.LineVsRadial.RadialHessianEventRealRank == realRadialRank && near(a.LineVsRadial.ComplexLineWeight, 0.5, 1e-15) && near(a.LineVsRadial.RadialHessianEventWeight, 0.25, 1e-15) && near(a.LineVsRadial.ComplexLineLoopUnit, 1/(4*math.Pi), 1e-15) && near(a.LineVsRadial.RadialHessianLoopUnit, 1/(8*math.Pi), 1e-15) && !a.LineVsRadial.FullComplexLineActive && a.LineVsRadial.RadialHessianEventActive && strings.Contains(a.LineVsRadial.PhysicalAmplitudeFluctuation, "one real radial"), Detail: FormatLineVsRadial(a.LineVsRadial)},
			{Name: "record potential capability boundaries", Passed: a.Capability.SourcesRealRadialEventType && a.Capability.SourcesOnePlusThreeSplit && a.Capability.SourcesRankOneEventWeight && !a.Capability.SourcesCP1Point && !a.Capability.SourcesGlobalGaugeFixing && !a.Capability.SourcesPhysicalSU2LTheorem && !a.Capability.SourcesNativeEWSBTheorem && !a.Capability.SourcesHiggsPoleMass && strings.Contains(a.Capability.Verdict, StatusSMLikePotentialSourcesRealRadialEventType), Detail: FormatCapability(a.Capability)},
			{Name: "record conditional source-type upgrade", Passed: strings.Contains(a.Upgrade.BeforeGate765, "imposed") && strings.Contains(a.Upgrade.AfterGate765, "Hessian/amplitude") && a.Upgrade.DependsOnSuppliedPotential && a.Upgrade.DependsOnSuppliedVacuumOrbit && !a.Upgrade.NativeUpgrade && a.Upgrade.ConditionalUpgrade && strings.Contains(a.Upgrade.Verdict, StatusSourceTypeUpgradeRecorded), Detail: FormatUpgrade(a.Upgrade)},
			{Name: "enforce physical firewalls", Passed: a.Firewalls.Audited && !a.Firewalls.SMLikePotentialNativeTheorem && !a.Firewalls.PotentialMinimumNativeVEVTheorem && !a.Firewalls.RadialEventNativeHistoryLoop && !a.Firewalls.OnePlusThreePhysicalGoldstone && !a.Firewalls.CP1FlatnessCompleteEWTheorem && !a.Firewalls.TreeRelationPoleMassTheorem && !a.Firewalls.LHopfNativeTransportTheorem && !a.Firewalls.ScalarRuntimeIndependentTheorem && !a.Firewalls.HiggsMassOrPoleMassTheorem && !a.Firewalls.YukawaOperatorOrEigenvalueTheorem && strings.Contains(a.Firewalls.Verdict, StatusGate765HiggsPotentialRadialEventBoundary), Detail: FormatFirewalls(a.Firewalls)},
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
