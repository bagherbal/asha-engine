package generation2rankcompleteexternalledger

import (
	"math"
	"strings"
	"testing"
)

func TestBuildDefaultGate471ComputesBridgeOnlyDUD(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Import.Loaded || a.Import.AcceptedRows == 0 || !a.Import.AllAcceptedBridgeOnly {
		t.Fatalf("expected rank-complete bridge ledger to load through airlock: %+v", a.Import)
	}
	if !a.Adapter.DUDComputed || !a.Adapter.CabibboResidualComputed || !a.Adapter.AlignmentAchieved {
		t.Fatalf("expected rank-complete ledger to compute aligned d_ud/residual: %+v", a.Adapter)
	}
	if math.Abs(a.Adapter.DUD-0.225) > 1e-12 {
		t.Fatalf("unexpected d_ud: %.17g", a.Adapter.DUD)
	}
	if a.Firewall.NativeRegistryWritten || a.Firewall.CKMNativePrediction || a.Firewall.CKMMatrixConstructed || a.Firewall.DUDNativePrediction {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
}

func TestInvertRejectsPhaseDomain(t *testing.T) {
	ik := 0.0
	badISpec := 1.0
	sig := 1
	sheet := 0
	r, err := invert(SectorInput{Sector: "u", IK: &ik, ISpec: &badISpec, SigmaCP: &sig, NC3: &sheet})
	if err != StatusFailedPhaseDomain || r.Defined {
		t.Fatalf("expected phase-domain rejection: ray=%+v err=%s", r, err)
	}
}

func TestRenderAuditContainsGate471Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 471 Registry Audit", StatusCKMGeometricAlignmentAchieved, "pdg_rank_complete_ledger.json", "Gate471 d_ud = 0.225", StatusExternalComparatorsNotPDGNative} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
