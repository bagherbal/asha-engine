package bimoduletracecapacity

import (
	"math"
	"testing"
)

func TestTotalCapacityBoundDoesNotSelectBranch(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if len(a.Stress.Results) != 2 {
		t.Fatalf("expected two branch results")
	}
	if !a.Stress.BothPassTotalCapacity {
		t.Fatalf("expected both branches to pass weak total capacity: %+v", a.Stress)
	}
	if a.Selection.UniqueBranchSelected {
		t.Fatalf("weak capacity must not select branch: %+v", a.Selection)
	}
}

func TestPerSlotDiagnosticWouldSelectRPlusButIsNotNative(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Stress.ExactlyOnePassPerSlot || a.Stress.PerSlotSelectedBranch != "r_plus" {
		t.Fatalf("expected per-slot diagnostic to isolate r_plus: %+v", a.Stress)
	}
	if a.Veto.PerSlotCapacityIsNativeTheorem || a.Veto.RMinusViolatesDerivedGeometry {
		t.Fatalf("per-slot diagnostic must stay unpromoted: %+v", a.Veto)
	}
	if len(a.Selection.SurvivingBranches) != 2 {
		t.Fatalf("both branches must survive native audit: %+v", a.Selection)
	}
}

func TestBranchNumbersMatchGate289Diagnostics(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	byName := map[string]Branch{}
	for _, b := range a.Branches {
		byName[b.Name] = b
	}
	rp := byName["r_plus"]
	rm := byName["r_minus"]
	if math.Abs(rp.LeptonD4-0.937151432355) > 1e-9 || math.Abs(rp.QuarkD4-7.612217870976) > 1e-9 {
		t.Fatalf("unexpected r_plus D4 sector values: %+v", rp)
	}
	if math.Abs(rm.LeptonD4-3.630368759358) > 1e-9 || math.Abs(rm.QuarkD4-4.919000543973) > 1e-9 {
		t.Fatalf("unexpected r_minus D4 sector values: %+v", rm)
	}
}

func TestFirewallsAndHiggsNoClaim(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Higgs.HiggsPredictionClaimed || a.Summary.HiggsPredictionDerived {
		t.Fatal("Gate 290 must not claim Higgs prediction")
	}
	if a.Firewalls.FiniteCorePolluted || !a.Firewalls.DoesNotVetoRMinusWithoutSelector {
		t.Fatalf("firewall failure: %+v", a.Firewalls)
	}
}
