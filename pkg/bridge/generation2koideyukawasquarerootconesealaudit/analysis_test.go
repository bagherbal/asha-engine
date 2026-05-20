package generation2koideyukawasquarerootconesealaudit

import "testing"

func TestGate577KoideSquareRootConeSealAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Runtime.Mu0GeV <= 0 || a.Runtime.Lambda12GeV <= 0 || a.Runtime.KoideQe <= 0 {
		t.Fatalf("bad runtime inheritance: %s", FormatRuntime(a.Runtime))
	}
	if !a.Geometry.PositiveConeOnly || a.Geometry.TargetQ != KoideTarget || a.Geometry.TargetAngleDeg != KoideAngleDeg {
		t.Fatalf("bad geometry definition: %s", FormatGeometry(a.Geometry))
	}
	mz := findPoint(a.Comparison.Points, "M_Z", "charged_leptons")
	if !mz.OnKoideCone1e5 || abs(mz.DeltaFromTwoThirds) > 1e-5 || abs(mz.AngleDeltaDeg) > 3e-4 {
		t.Fatalf("charged lepton Koide alignment not sharp: %s", FormatConePoint(mz))
	}
	up := findPoint(a.Comparison.Points, "M_Z", "up_quarks")
	down := findPoint(a.Comparison.Points, "M_Z", "down_quarks")
	if up.OnKoideCone1e4 || down.OnKoideCone1e4 || a.Comparison.KoideUniversalAcrossSectors {
		t.Fatalf("Koide incorrectly universalized: %s", FormatComparison(a.Comparison))
	}
	if a.Seal.NativeDerivation || !a.Seal.BridgeOnly || a.Seal.RemainingContinuousCoordinates != 2 {
		t.Fatalf("bad minimal environmental seal: %s", FormatSeal(a.Seal))
	}
	if !a.Gate352.EmpiricalAlignment || a.Gate352.NativePromotion || a.Gate352.RootTraceNative || a.Gate352.PfaffianCanGenerate {
		t.Fatalf("Gate352 obstruction not preserved: %s", FormatGate352(a.Gate352))
	}
	if a.Firewalls.DerivesChargedLeptonMasses || a.Firewalls.DerivesYukawaEigenvalues || a.Firewalls.DerivesCKM || a.Firewalls.ImportsObservedAsNative || a.Firewalls.AddsNewCarrier || !a.Firewalls.PreservesGate352 {
		t.Fatalf("firewall broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate577ConeIdentity(t *testing.T) {
	p := conePoint("test", "charged_leptons", []string{"e", "mu", "tau"}, []float64{0.0000029350283095504176, 0.0006068707640859305, 0.010205763440624986})
	if abs(p.Q-0.6666605114773856) > 1e-15 {
		t.Fatalf("unexpected Q: %s", FormatConePoint(p))
	}
	if abs(p.PerpOverParallel-1) > 1e-5 {
		t.Fatalf("Koide cone should have perp/parallel near 1: %s", FormatConePoint(p))
	}
}

func TestGate577Theorem(t *testing.T) {
	res := Generation2KoideYukawaSquareRootConeEnvironmentalSealAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
