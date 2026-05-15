package generation2observednumericaladapter

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate470FailsClosedOnCheckedInLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Import.Loaded || a.Import.AcceptedRows == 0 || !a.Import.AllAcceptedBridgeOnly {
		t.Fatalf("expected explicit bridge ledger to load through airlock: %+v", a.Import)
	}
	if a.Adapter.DUDComputed || a.Adapter.CabibboResidualComputed || a.Adapter.AlignmentAchieved {
		t.Fatalf("default PDG-style ledger must not compute d_ud/residual: %+v", a.Adapter)
	}
	if !a.Adapter.MissingISpecIKValues || !a.Adapter.MissingBranchTags || !a.Adapter.PDGNoIK {
		t.Fatalf("expected missing comparator/branch diagnostics: %+v", a.Adapter)
	}
	if a.Firewall.NativeRegistryWritten || a.Firewall.CKMNativePrediction || a.Firewall.CKMMatrixConstructed {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
}

func TestInvertRejectsProjectiveBoundary(t *testing.T) {
	one := 1.0
	zero := 0.0
	sig := 1
	sheet := 0
	r, err := invert(SectorInput{Sector: "u", IK: &one, ISpec: &zero, SigmaCP: &sig, NC3: &sheet})
	if err != StatusFailedProjectiveDomain || r.Defined {
		t.Fatalf("expected projective boundary rejection: ray=%+v err=%s", r, err)
	}
}

func TestRenderAuditContainsGate470Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 470 Registry Audit", StatusFailedDUDNotComputableFromFile, "pdg_observed_ledger.json", "Gate470 d_ud = undefined", StatusFailedPDGNoIK} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
