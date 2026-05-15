package generation2leptonpreflight

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate475DefinesLeptonPreflight(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate474NoNativeIK || !a.Inheritance.PMNSBridgeFrontier {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if !a.Schema.RequiresISpecIK || !a.Schema.RequiresBranchTags || !a.Schema.RequiresNeutrinoOrderingPolicy || a.Schema.AllowsPMNSAsRayInput || a.Schema.ComputesNow {
		t.Fatalf("bad schema: %+v", a.Schema)
	}
	if a.Sieve.AcceptedBridgeRows != 1 || a.Sieve.ComputesPMNSResidual || a.Sieve.ComputesIK {
		t.Fatalf("bad sieve: %+v", a.Sieve)
	}
	if !a.Sieve.Probes[3].Accepted || !a.Sieve.Probes[3].BridgeOnly {
		t.Fatalf("complete bridge preflight should pass: %+v", a.Sieve.Probes[3])
	}
	if a.Firewall.LeptonDataImported || a.Firewall.PMNSMatrixComputed || a.Firewall.PMNSNativePrediction || a.Firewall.IKNativeSelectorFound || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall violated: %+v", a.Firewall)
	}
}

func TestRenderAuditContainsGate475Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 475 Registry Audit", StatusPreflightDefined, StatusFailedPMNSAsCoordinate, "e, nu", "PMNS may be a residual target only"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
