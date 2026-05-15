package generation2syntheticreleasereviewmanifestadapter

import (
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(a.Truth, "synthetic release-review") {
		t.Fatalf("unexpected truth: %s", a.Truth)
	}
	if !a.Import.ChecksumVerified || !a.Review.BlockedBecauseSynthetic || a.Firewall.NativeRegistryWritten {
		t.Fatalf("Gate547 leaked: import=%+v review=%+v firewall=%+v", a.Import, a.Review, a.Firewall)
	}
}

func TestReleaseManifestRows(t *testing.T) {
	a, err := Build(DefaultLedger)
	if err != nil {
		t.Fatal(err)
	}
	if a.Import.AcceptedRows != 15 || len(a.Import.MissingRows) > 0 || a.Import.RejectedRows != 0 || len(a.Import.DuplicateRows) > 0 {
		t.Fatalf("manifest rows incomplete: %+v", a.Import)
	}
	if !a.Import.AllBridgeOnly || !a.Import.AllReleaseOnly || !a.Import.AllQuarantineOnly || !a.Import.AllDryRunOnly || !a.Import.AllSynthetic || !a.Import.AllNoTheorem {
		t.Fatalf("metadata sieve failed: %+v", a.Import)
	}
}

func TestReleaseReviewBlocked(t *testing.T) {
	a, err := Build(DefaultLedger)
	if err != nil {
		t.Fatal(err)
	}
	if !a.Review.HumanReviewMetadataParsed || !a.Review.ReproducibilityMetadataParsed || !a.Review.SourceChainMetadataParsed || !a.Review.NativeWriteDeltaZero {
		t.Fatalf("review metadata incomplete: %+v", a.Review)
	}
	if a.Review.AuthenticatedSourceChain || a.Review.ReleaseAllowed || a.Review.BridgeEvidenceReleased || a.Review.NativeWriteAuthorization || a.Review.NativeRegistryWrite {
		t.Fatalf("synthetic release unexpectedly authorized: %+v", a.Review)
	}
}

func TestPhysicalClaimsRemainFalse(t *testing.T) {
	a, err := Build(DefaultLedger)
	if err != nil {
		t.Fatal(err)
	}
	if a.Firewall.ComparatorOutputReleased || a.Firewall.BridgeEvidenceClaimReleased || a.Firewall.RealSchwingerSourceImported || a.Firewall.AuthenticatedRealSource || a.Firewall.PhysicalSchwingerFunctionsLoaded || a.Firewall.OSPositivityCertificateLoaded || a.Firewall.WickMapLoaded || a.Firewall.HilbertSpaceReconstructed || a.Firewall.HamiltonianSpectrumLoaded || a.Firewall.UnitaryDynamicsLoaded || a.Firewall.GlobalCausalityLoaded || a.Firewall.TimeArrowLoaded || a.Firewall.NativeRegistryWritten {
		t.Fatalf("firewall leaked: %+v", a.Firewall)
	}
}

func TestAsTheorem(t *testing.T) {
	res := Generation2SyntheticReleaseReviewManifestAdapterDryRunTheorem().Verify()
	for _, check := range res.Checks {
		if !check.Passed {
			t.Fatalf("check failed: %s: %s", check.Name, check.Detail)
		}
	}
}
