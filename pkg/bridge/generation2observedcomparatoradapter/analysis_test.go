package generation2observedcomparatoradapter

import (
	"math"
	"strings"
	"testing"
)

func TestBuildDefaultGate466(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Airlock.AcceptedRows != RequiredRows || a.Airlock.QuarkMassRowsImported != 6 || a.Airlock.CKMRowsImported != 1 {
		t.Fatalf("unexpected airlock import: %+v", a.Airlock)
	}
	if !a.Airlock.AllAcceptedQuarantined || !a.Airlock.NativePromotionRejectedProbe || !a.Airlock.NativeRegistryWriteRejected || !a.Airlock.ObservedAsTheoremRejected {
		t.Fatalf("airlock failed to quarantine or reject unsafe probes: %+v", a.Airlock)
	}
	if a.Coordinate.DUDDefined || !math.IsNaN(a.Coordinate.DUD) || a.Coordinate.AlignmentAchieved || a.Coordinate.IKComparatorSupplied || a.Coordinate.BranchTagsSupplied {
		t.Fatalf("d_ud should be undefined without I_K and branch tags: %+v", a.Coordinate)
	}
	if a.Firewall.EmpiricalDataInNativeRegistry || a.Firewall.CKMNativePrediction || a.Firewall.QuarkMassNativePrediction || a.Firewall.DUDPromotedNative || a.Firewall.AlignmentPromotedNative {
		t.Fatalf("native firewall violated: %+v", a.Firewall)
	}
}

func TestRenderAuditContainsGate466Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 466 Registry Audit", StatusFailedAlignmentNotComputable, "d_ud = undefined", "|V_us| = 0.225", StatusFailedMassSpectraDoNotDefineRay} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
