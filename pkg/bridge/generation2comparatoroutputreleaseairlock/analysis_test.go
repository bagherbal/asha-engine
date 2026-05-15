package generation2comparatoroutputreleaseairlock

import (
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Truth, "release airlock") {
		t.Fatalf("unexpected truth: %s", a.Truth)
	}
	if !a.Inheritance.Gate545QuarantineOutput || !a.Guard.AbortTriggeredBySynthetic || a.Firewall.NativeRegistryWritten {
		t.Fatalf("Gate546 leaked: inherit=%+v guard=%+v firewall=%+v", a.Inheritance, a.Guard, a.Firewall)
	}
}

func TestReleaseSchemaRows(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Schema.RequiredRows != 15 || a.Schema.HumanReviewRows == 0 || a.Schema.ReproducibilityRows == 0 || a.Schema.SourceAuthenticityRows == 0 || a.Schema.CitationScopeRows == 0 || a.Schema.NativeWriteLockRows == 0 || a.Schema.RollbackRows == 0 {
		t.Fatalf("schema incomplete: %+v", a.Schema)
	}
}

func TestReleaseBlocked(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Guard.ReleaseManifestImported || a.Guard.HumanReviewCompleted || a.Guard.ReproducibilityCompleted || a.Guard.SourceChainAuthenticated || a.Guard.BridgeEvidenceReleaseAllowed || a.Guard.BridgeEvidenceReleased {
		t.Fatalf("release unexpectedly authorized: %+v", a.Guard)
	}
	if !a.Guard.NativeWriteLocked || a.Guard.NativeWriteAuthorization || a.Guard.NativeRegistryWrite {
		t.Fatalf("native lock failed: %+v", a.Guard)
	}
}

func TestPhysicalClaimsRemainFalse(t *testing.T) {
	a, err := Build()
	if err != nil {
		t.Fatal(err)
	}
	if a.Firewall.ComparatorOutputReleased || a.Firewall.BridgeEvidenceClaimReleased || a.Firewall.RealSchwingerSourceImported || a.Firewall.AuthenticatedRealSource || a.Firewall.PhysicalSchwingerFunctionsLoaded || a.Firewall.OSPositivityCertificateLoaded || a.Firewall.WickMapLoaded || a.Firewall.HilbertSpaceReconstructed || a.Firewall.HamiltonianSpectrumLoaded || a.Firewall.UnitaryDynamicsLoaded || a.Firewall.GlobalCausalityLoaded || a.Firewall.TimeArrowLoaded || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall leaked: %+v", a.Firewall)
	}
}

func TestAsTheorem(t *testing.T) {
	res := Generation2ComparatorOutputReleaseAirlockPreflightTheorem().Verify()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s: %s", check.Name, check.Detail)
		}
	}
}
