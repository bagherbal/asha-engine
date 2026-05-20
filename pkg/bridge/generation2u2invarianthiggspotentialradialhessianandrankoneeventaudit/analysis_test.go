package generation2u2invarianthiggspotentialradialhessianandrankoneeventaudit

import (
	"math"
	"strings"
	"testing"
)

func near(x, y float64) bool { return math.Abs(x-y) <= 1e-15 }

func TestGate765InheritancePotentialAndCP1Flatness(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate764.Inherited || a.Gate764.Carrier != "K7+_J(n) ~= C^2" || !strings.Contains(a.Gate764.PRad, "rank-one") || !strings.Contains(a.Gate764.LHopfFormula, "1/(8*pi)") || !a.Gate764.CP1LocationGaugeRepresentative || !a.Gate764.LHopfDependsOnRankOneRadialEvent || a.Gate764.ComplexVacuumLineScalarSource || !strings.Contains(a.Gate764.RemainingScalarSourceQuestion, "rank-one radial") {
		t.Fatalf("bad Gate764 inheritance: %+v", a.Gate764)
	}
	if !strings.Contains(a.Potential.PotentialFormula, "mu^2") || !strings.Contains(a.Potential.PotentialFormula, "lambda") || a.Potential.RadiusVariable != "r^2 = phi^dagger phi" || !a.Potential.DependsOnlyOnRadius || a.Potential.Carrier != "K7+_J(n) ~= C^2" || !a.Potential.U2Invariant || !a.Potential.ScalarPotentialSupplied || a.Potential.NativePotentialTheorem {
		t.Fatalf("bad potential audit: %+v", a.Potential)
	}
	if !strings.Contains(a.CP1Flatness.FixedRadiusOrbit, "CP1") || !a.CP1Flatness.PotentialConstantOnCP1 || a.CP1Flatness.SelectsPiVacC || !a.CP1Flatness.ConfirmsGate764Demotion || !a.CP1Flatness.RequiresAnisotropyForLine {
		t.Fatalf("bad CP1 flatness audit: %+v", a.CP1Flatness)
	}
}

func TestGate765VacuumRadiusAndRadialHessian(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.VacuumRadius.LambdaPositive || !a.VacuumRadius.MuSquaredNegative || !strings.Contains(a.VacuumRadius.StationaryCondition, "mu^2+2 lambda") || a.VacuumRadius.PhiDaggerPhiRelation != "phi^dagger phi = -mu^2/(2 lambda)" || a.VacuumRadius.VEVConvention != "phi^dagger phi = v^2/2" || a.VacuumRadius.VSquaredRelation != "v^2 = -mu^2/lambda" || !a.VacuumRadius.ConventionDependent || a.VacuumRadius.NativeVEVTheorem {
		t.Fatalf("bad vacuum-radius audit: %+v", a.VacuumRadius)
	}
	if !strings.Contains(a.RadialHessian.SuppliedVacuumRepresentative, "phi_0") || a.RadialHessian.RadialPath != "phi(t)=(1+t)phi_0" || !a.RadialHessian.DefinesPRad || a.RadialHessian.RadialRealDimension != realRadialRank || a.RadialHessian.AngularRealDimension != angularRealDim || a.RadialHessian.K7PlusRealDimension != k7PlusRealDim || !a.RadialHessian.AngularPreservesRadiusFirstOrder || !strings.Contains(a.RadialHessian.SplitFormula, "1 + 3") || !a.RadialHessian.PotentialHessianSplit || a.RadialHessian.PhysicalGoldstoneTheorem {
		t.Fatalf("bad radial Hessian audit: %+v", a.RadialHessian)
	}
}

func TestGate765RankWeightAndComplexLineDistinction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.RankWeight.RhoPlusFormula, "rank(P_rad)/4") || a.RankWeight.RadialRank != realRadialRank || a.RankWeight.K7PlusRealDimension != k7PlusRealDim || !near(a.RankWeight.RadialEventWeight, 0.25) || !near(a.RankWeight.PhaseLoopPayoff, 1/(2*math.Pi)) || !near(a.RankWeight.LHopf, 1/(8*math.Pi)) || !near(a.RankWeight.LHopf, a.RankWeight.ExpectedLHopf) || !a.RankWeight.MatchesHistoryLoopQuarter || a.RankWeight.NativeHistoryLoopTheorem {
		t.Fatalf("bad rank-one event weight: %+v", a.RankWeight)
	}
	if a.LineVsRadial.ComplexLineRealRank != complexLineRealRank || a.LineVsRadial.RadialHessianEventRealRank != realRadialRank || !near(a.LineVsRadial.ComplexLineWeight, 0.5) || !near(a.LineVsRadial.RadialHessianEventWeight, 0.25) || !near(a.LineVsRadial.ComplexLineLoopUnit, 1/(4*math.Pi)) || !near(a.LineVsRadial.RadialHessianLoopUnit, 1/(8*math.Pi)) || a.LineVsRadial.FullComplexLineActive || !a.LineVsRadial.RadialHessianEventActive || !strings.Contains(a.LineVsRadial.PhysicalAmplitudeFluctuation, "one real radial") {
		t.Fatalf("bad line-vs-radial audit: %+v", a.LineVsRadial)
	}
}

func TestGate765CapabilityUpgradeFirewallsAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Capability.SourcesRealRadialEventType || !a.Capability.SourcesOnePlusThreeSplit || !a.Capability.SourcesRankOneEventWeight || a.Capability.SourcesCP1Point || a.Capability.SourcesGlobalGaugeFixing || a.Capability.SourcesPhysicalSU2LTheorem || a.Capability.SourcesNativeEWSBTheorem || a.Capability.SourcesHiggsPoleMass {
		t.Fatalf("bad potential capability audit: %+v", a.Capability)
	}
	if !strings.Contains(a.Upgrade.BeforeGate765, "imposed") || !strings.Contains(a.Upgrade.AfterGate765, "Hessian/amplitude") || !a.Upgrade.DependsOnSuppliedPotential || !a.Upgrade.DependsOnSuppliedVacuumOrbit || a.Upgrade.NativeUpgrade || !a.Upgrade.ConditionalUpgrade {
		t.Fatalf("bad source-type upgrade: %+v", a.Upgrade)
	}
	if !a.Firewalls.Audited || a.Firewalls.SMLikePotentialNativeTheorem || a.Firewalls.PotentialMinimumNativeVEVTheorem || a.Firewalls.RadialEventNativeHistoryLoop || a.Firewalls.OnePlusThreePhysicalGoldstone || a.Firewalls.CP1FlatnessCompleteEWTheorem || a.Firewalls.TreeRelationPoleMassTheorem || a.Firewalls.LHopfNativeTransportTheorem || a.Firewalls.ScalarRuntimeIndependentTheorem || a.Firewalls.HiggsMassOrPoleMassTheorem || a.Firewalls.YukawaOperatorOrEigenvalueTheorem {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	res := Generation2U2InvariantHiggsPotentialRadialHessianAndRankOneEventAuditTheorem().Verify()
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
