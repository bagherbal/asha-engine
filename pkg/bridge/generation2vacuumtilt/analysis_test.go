package generation2vacuumtilt

import (
	"math"
	"strings"
	"testing"
)

func TestBuildDefaultGate484VacuumTiltAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.C3BasisAudit.AllSectorsExactlyRepresented || !a.C3BasisAudit.RepresentationModuliNeutral {
		t.Fatalf("bad C3 basis audit: %+v", a.C3BasisAudit)
	}
	if !a.KoideAudit.ChargedLeptonPasses || a.KoideAudit.UpQuarkPasses || a.KoideAudit.DownQuarkPasses {
		t.Fatalf("bad Koide audit: %+v", a.KoideAudit)
	}
	if math.Abs(a.KoideAudit.ChargedLeptonResidual) > 1e-4 {
		t.Fatalf("charged-lepton Koide residual too large: %.12g", a.KoideAudit.ChargedLeptonResidual)
	}
	if a.UniversalTilt.OneUniversalTiltVectorSupported || a.UniversalTilt.ReducesModuli {
		t.Fatalf("universal tilt should fail closed: %+v", a.UniversalTilt)
	}
	if a.Compression.FlavorModuliReducedByCurrentGate || a.Firewall.NativeRegistryWritten || a.Firewall.PhysicalDUDComputed || a.Firewall.PhysicalDENuComputed {
		t.Fatalf("firewall leak: %+v %+v", a.Compression, a.Firewall)
	}
}

func TestRenderAuditGate484(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusFailedTiltSliceReparam, StatusChargedLeptonKoideShadowFound, "R/S = sqrt(2)", "physical d_ud = undefined", "charged leptons"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
