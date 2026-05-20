package generation2koidechamberwalloffsetaudit

import (
	"math"
	"testing"
)

func TestGate583KoideChamberWallOffsetAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.MZ.EpsilonDeg-2.26718003289167) > 1e-10 {
		t.Fatalf("unexpected M_Z wall offset: %s", FormatWallPoint(a.MZ))
	}
	if math.Abs(a.Lambda12.EpsilonDeg-2.26738253154505) > 1e-10 {
		t.Fatalf("unexpected Lambda_12 wall offset: %s", FormatWallPoint(a.Lambda12))
	}
	if math.Abs(a.MZ.EpsilonRad-0.039569756309433) > 1e-12 {
		t.Fatalf("unexpected epsilon radians: %s", FormatWallPoint(a.MZ))
	}
	if !a.MZ.InsideCanonicalChamber || !a.Lambda12.InsideCanonicalChamber {
		t.Fatalf("canonical charged-lepton point should remain in chamber: %s / %s", FormatWallPoint(a.MZ), FormatWallPoint(a.Lambda12))
	}
	if math.Abs(a.MZ.ElectronRootOverA-0.0403510719726994) > 1e-12 {
		t.Fatalf("unexpected electron wall coordinate: %s", FormatWallPoint(a.MZ))
	}
	if math.Abs(a.MZ.QuadraticResidual) > 2e-6 {
		t.Fatalf("near-wall quadratic approximation should be sharp: %s", FormatWallPoint(a.MZ))
	}
	if !(a.Transport.EpsilonStable && a.Transport.ChamberPreserved && a.Transport.AmplitudeMovesToward1) {
		t.Fatalf("wall offset should be stable and chamber preserved: %s", FormatTransport(a.Transport))
	}
	if a.Quarks.Up.WallSealValid || a.Quarks.Down.WallSealValid {
		t.Fatalf("quark wall seals should not certify: %s", FormatQuarks(a.Quarks))
	}
	if a.Firewalls.DerivesEpsilon || a.Firewalls.AddsNewCarrier || a.Firewalls.PromotesObservedAsNative || !a.Firewalls.PreservesGate352 {
		t.Fatalf("firewall broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate583Theorem(t *testing.T) {
	res := Generation2KoideChamberWallOffsetAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
