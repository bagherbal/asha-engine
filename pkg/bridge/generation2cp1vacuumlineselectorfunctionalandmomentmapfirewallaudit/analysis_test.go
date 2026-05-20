package generation2cp1vacuumlineselectorfunctionalandmomentmapfirewallaudit

import (
	"strings"
	"testing"
)

func TestGate763InheritanceAndFunctionalRequirement(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Gate762.Inherited || a.Gate762.Socket != "K7+_J(n) ~= C^2" || !strings.Contains(a.Gate762.MissingObject, "Pi_vac_C") || a.Gate762.MissingObjectRankR != complexLineRealRank || a.Gate762.MissingObjectRankC != complexLineRank || !a.Gate762.PRadSecondary || !a.Gate762.NoCurrentCP1Selector || !a.Gate762.ComplexVacuumLineSealRemains {
		t.Fatalf("bad Gate762 inheritance: %+v", a.Gate762)
	}
	if !strings.Contains(a.Requirement.Question, "functional") || a.Requirement.Domain != "CP1 = P(C^2)" || !strings.Contains(a.Requirement.Target, "Pi_vac_C") || !strings.Contains(a.Requirement.EquivalentHermitianAxis, "su(2)") || !a.Requirement.RequiresNonconstantFunctional || !a.Requirement.RequiresIsolatedCriticalLine || a.Requirement.U2InvariantFunctionalSelectsLine || !a.Requirement.ConstantFunctionalLeavesCP1Flat || a.Requirement.RadialGaugeRequiredForLineSelector {
		t.Fatalf("bad functional requirement: %+v", a.Requirement)
	}
}

func TestGate763MomentMapScalarPotentialAndBoundaryStress(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.MomentMap.MomentMapExistsOnCP1 || !strings.Contains(a.MomentMap.MomentMapFormula, "zz^dagger") || !strings.Contains(a.MomentMap.HamiltonianFormula, "<h") || !a.MomentMap.RequiresSU2Axis || a.MomentMap.SuppliedAxisCertified || !a.MomentMap.WouldSelectEigenlineIfAxis || a.MomentMap.NativeSelectorCertified {
		t.Fatalf("bad moment-map audit: %+v", a.MomentMap)
	}
	if !strings.Contains(a.ScalarPotential.InvariantPotentialFormula, "|z|^2") || !strings.Contains(a.ScalarPotential.AnisotropicPotentialFormula, "H") || !a.ScalarPotential.U2Invariant || !a.ScalarPotential.FlatOnCP1 || !a.ScalarPotential.AnisotropicTermWouldSelectLine || a.ScalarPotential.AnisotropicTermCertified || a.ScalarPotential.NativeOrientationSelector {
		t.Fatalf("bad scalar-potential audit: %+v", a.ScalarPotential)
	}
	if len(a.BoundaryStress.BoundaryObjects) != 5 || !a.BoundaryStress.CollapsedScalarLayer || a.BoundaryStress.ProvidesVectorInK7Plus || a.BoundaryStress.ProvidesHermitianAxisOnC2 || a.BoundaryStress.CanDefineCP1Functional || a.BoundaryStress.CanSelectCP1Point || !a.BoundaryStress.NeedsTypedVectorCoupling {
		t.Fatalf("bad boundary-stress audit: %+v", a.BoundaryStress)
	}
}

func TestGate763FanoOrientationRankingAndOrder(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Fano.SuppliesTwistorFamily || !a.Fano.SuppliesJHSocket || !a.Fano.SuppliesU2Socket || a.Fano.SelectsComplexStructureN || a.Fano.SelectsCP1Point || a.Fano.InvariantVectorInK7Plus || !a.Fano.SymmetryWouldMakeLineArbitrary {
		t.Fatalf("bad Fano/quaternionic audit: %+v", a.Fano)
	}
	if !strings.Contains(a.Orientation.SealName, "ComplexVacuumLineSeal") || !a.Orientation.CanSelectCP1IfSupplied || a.Orientation.NativeTheoremCertified || !a.Orientation.WouldPrecedeRadialGauge || !a.Orientation.WouldNotDeriveEWSB || !a.Orientation.WouldNotDeriveHiggsMass {
		t.Fatalf("bad orientation seal audit: %+v", a.Orientation)
	}
	if !a.Ranking.RankingRecorded || len(a.Ranking.Candidates) != 5 || a.Ranking.AnyNativeSelectorCertified || !strings.Contains(a.Ranking.BestTypedFutureCandidate, "SU(2)") || !strings.Contains(a.Ranking.HighestNativeResult, "no CP1 point") {
		t.Fatalf("bad ranking: %+v", a.Ranking)
	}
	for _, c := range a.Ranking.Candidates {
		if c.NativeSelectorCertified {
			t.Fatalf("candidate unexpectedly native-certified: %+v", c)
		}
	}
	if !a.SealOrder.Preserved || !a.SealOrder.LineBeforeGauge || !a.SealOrder.LHopfAfterGauge || !strings.Contains(a.SealOrder.Step1, "J_H") || !strings.Contains(a.SealOrder.Step2, "CP1") || !strings.Contains(a.SealOrder.Step3, "P_rad") || !strings.Contains(a.SealOrder.Step4, "L_Hopf") {
		t.Fatalf("bad seal order: %+v", a.SealOrder)
	}
}

func TestGate763FirewallsAndTheoremStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewalls.Audited || a.Firewalls.CP1FunctionalNativeSelector || a.Firewalls.MomentMapAxisNative || a.Firewalls.HermitianAxisNative || a.Firewalls.ScalarPotentialNativeOrientation || a.Firewalls.BoundaryStressCP1Functional || a.Firewalls.PiVacCNativeEWSBTheorem || a.Firewalls.RadialGaugeFixingNativeTheorem || a.Firewalls.LHopfNativeHistoryLoopTheorem || a.Firewalls.ScalarRuntimeIndependentTheorem || a.Firewalls.HiggsMassOrPoleMassTheorem || a.Firewalls.YukawaOperatorOrEigenvalueTheorem {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	res := Generation2CP1VacuumLineSelectorFunctionalAndMomentMapFirewallAuditTheorem().Verify()
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
