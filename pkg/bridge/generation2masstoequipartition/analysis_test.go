package generation2masstoequipartition

import (
	"math"
	"strings"
	"testing"
)

func TestBuildDefaultGate473RejectsMassToEquipartition(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Import.Loaded || a.Import.AcceptedRows != 6 || !a.Import.IKImportRejected || !a.Import.CKMImportRejected {
		t.Fatalf("bad import: %+v", a.Import)
	}
	if !a.Up.ExtremeHierarchy || !a.Down.ExtremeHierarchy {
		t.Fatalf("expected extreme hierarchy: %+v %+v", a.Up, a.Down)
	}
	if a.Up.AlphaOneCompatible || a.Down.AlphaOneCompatible {
		t.Fatalf("alpha=1 should be incompatible with extreme trace-zero spectra: %+v %+v", a.Up, a.Down)
	}
	if a.Up.IKHalfDerived || a.Down.IKHalfDerived || a.Loop.IKDerived || a.Loop.DUDComputed || a.Loop.AlignmentAchieved {
		t.Fatalf("loop must not close: %+v", a.Loop)
	}
	if a.Firewall.NativeRegistryWritten || a.Firewall.CKMNativePrediction || a.Firewall.DUDNativePrediction {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
}

func TestAlphaMaxFromISpec(t *testing.T) {
	if !math.IsInf(alphaMaxFromISpec(0), 1) {
		t.Fatal("zero spectrum invariant should leave alpha unbounded")
	}
	max := 2 / (3 * math.Sqrt(3))
	if math.Abs(alphaMaxFromISpec(max)) > 1e-9 {
		t.Fatalf("maximal extreme invariant should force alpha max 0, got %.12g", alphaMaxFromISpec(max))
	}
	if alphaMaxFromISpec(0.25) <= 0.99 {
		t.Fatalf("I_spec=0.25 should allow alpha around 1")
	}
}

func TestRenderAuditContainsGate473Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 473 Registry Audit", StatusFailedProjectNotAchieved, StatusFailedMassHierarchyNoEquipartition, "d_ud = undefined", "I_K=0.5"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
