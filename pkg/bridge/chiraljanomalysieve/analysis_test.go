package chiraljanomalysieve

import (
	"math"
	"testing"
)

func TestGammaTracesAreBranchBlind(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Chiral.AllGammaD2Zero || !a.Chiral.AllGammaD4Zero || a.Chiral.BranchSensitiveViaGamma {
		t.Fatalf("expected branch-blind zero gamma traces: %+v", a.Chiral)
	}
	for _, tr := range a.Chiral.Traces {
		if math.Abs(tr.TrGammaD2) > 1e-14 || math.Abs(tr.TrGammaD4) > 1e-14 {
			t.Fatalf("nonzero gamma trace: %+v", tr)
		}
	}
}

func TestSectorProjectedTracesDistinguishBranchesButDoNotSelect(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sector.BranchSensitive {
		t.Fatal("expected sector-projected traces to distinguish r branches")
	}
	if a.Sector.NativeSelectionFunctional || a.Sector.SelectedBranch != "" {
		t.Fatalf("sector projection must not become a selection rule: %+v", a.Sector)
	}
	if len(a.Sector.Traces) != 2 {
		t.Fatalf("expected two branch traces")
	}
	if math.Abs(a.Sector.Traces[0].TrPCD2-a.Sector.Traces[1].TrPCD2) < 1e-6 {
		t.Fatalf("expected lepton-sector D2 to differ between branches")
	}
}

func TestJAndAnomalyRemainMissing(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.J.PhysicalJDerived || a.J.OppositeActionConstructed {
		t.Fatalf("physical J should not be derived in Gate 289: %+v", a.J)
	}
	if a.Anomaly.CanEliminateBranch || a.Anomaly.AnomalyEquationsDependOnR {
		t.Fatalf("anomaly audit should not eliminate an r branch: %+v", a.Anomaly)
	}
}

func TestBothBranchesSurviveAndHiggsIsNotClaimed(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Sieve.UniqueBranchSelected || len(a.Sieve.SurvivingBranches) != 2 {
		t.Fatalf("expected both branches to survive: %+v", a.Sieve)
	}
	if a.Higgs.HiggsPredictionClaimed || a.Summary.HiggsPredictionDerived {
		t.Fatal("Gate 289 must not claim a Higgs prediction")
	}
	if a.Firewalls.FiniteCorePolluted || !a.Firewalls.DoesNotDiscardSurvivingBranch {
		t.Fatalf("firewall failure: %+v", a.Firewalls)
	}
}
