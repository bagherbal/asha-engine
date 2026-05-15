package generation2leptonobservedadapter

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate478FailsClosedOnCheckedInLeptonLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Import.Loaded || a.Import.AcceptedRows == 0 || !a.Import.AllAcceptedBridgeOnly || !a.Import.LeptonPoliciesComplete {
		t.Fatalf("expected explicit lepton bridge ledger to load through airlock: %+v", a.Import)
	}
	if a.Adapter.DENuComputed || a.Adapter.PMNSResidualComputed || a.Adapter.AlignmentAchieved {
		t.Fatalf("default lepton observed ledger must not compute d_eν/residual: %+v", a.Adapter)
	}
	if !a.Adapter.MissingISpecIKValues || !a.Adapter.MissingBranchTags || !a.Adapter.LeptonDataNoIK {
		t.Fatalf("expected missing comparator/branch diagnostics: %+v", a.Adapter)
	}
	if a.Firewall.NativeRegistryWritten || a.Firewall.PMNSNativePrediction || a.Firewall.PMNSMatrixConstructed || a.Firewall.PMNSEntryComputed {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
}

func TestInvertRejectsLeptonProjectiveBoundary(t *testing.T) {
	one := 1.0
	zero := 0.0
	sig := 1
	sheet := 0
	r, err := invert(SectorInput{Sector: "e", IK: &one, ISpec: &zero, SigmaCP: &sig, NC3: &sheet})
	if err != StatusFailedProjectiveDomain || r.Defined {
		t.Fatalf("expected projective boundary rejection: ray=%+v err=%s", r, err)
	}
}

func TestRenderAuditContainsGate478Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 478 Registry Audit", StatusFailedDENuNotComputableFromFile, "lepton_observed_ledger.json", "Gate478 d_eν = undefined", StatusFailedLeptonDataNoIK, "structurally identical"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
