package generation2nullconeikselector

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate480DerivesNullBaseline(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Map.NativeNullConeExists || !a.Map.KTimelikeAssigned || !a.Map.XYSpacelikeAssigned || !a.Map.NullBoundaryDeclaredForGate480 {
		t.Fatalf("bad map: %+v", a.Map)
	}
	if !a.Sieve.NullForcesAEqualsR || a.Sieve.AlphaVac != 1 || a.Sieve.IKVac != 0.5 || a.Sieve.AcceptedNullCases != 2 {
		t.Fatalf("bad null sieve: %+v", a.Sieve)
	}
	if a.Gap.PhysicalSectorIKDefined || a.Gap.DUDComputed || a.Gap.DENuComputed || a.Gap.CKMPrediction || a.Gap.PMNSPrediction {
		t.Fatalf("coordinate gap leaked: %+v", a.Gap)
	}
	if a.Firewall.ObservedMassImported || a.Firewall.CKMImported || a.Firewall.PMNSImported || a.Firewall.VacuumIKPhysicalSectorCoordinate || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
}

func TestIKFormula(t *testing.T) {
	if IK(1) != 0.5 {
		t.Fatalf("expected IK(1)=0.5, got %.15f", IK(1))
	}
	if !evaluateCase("null", 3, 0, 3).AcceptedVacuumBaseline {
		t.Fatal("scaled positive null case should be alpha=1")
	}
	if evaluateCase("timelike", 2, 1, 0).AcceptedVacuumBaseline {
		t.Fatal("timelike case should not be accepted")
	}
}

func TestRenderAuditContainsGate480Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 480 Registry Audit", StatusIKDerived, "alpha_vac", "I_K", "vacuum baseline", StatusFailedPhysicalIKPromotion} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
