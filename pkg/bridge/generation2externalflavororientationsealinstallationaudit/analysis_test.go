package generation2externalflavororientationsealinstallationaudit

import (
	"strings"
	"testing"
)

func TestGate964BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus || a.NextGate != NextGate {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.Decision.ExternalFlavorOrientationSealInstalled || a.Decision.ExternalFlavorOrientationSealNative {
		t.Fatalf("external seal should be installed but non-native: %#v", a.Decision)
	}
	if !a.Decision.RepresentativeChosenForDiagnostics || !a.Decision.U3OrbitAcknowledged || a.Decision.CanonicalFlavorSelectorCertified {
		t.Fatalf("representative/orbit status wrong: %#v", a.Decision)
	}
	if !a.Decision.TripleSealLaneActive || !a.Decision.R3DualSealPreserved || !a.Decision.ExternalGenerationCarrierSealPreserved {
		t.Fatalf("triple seal lane not preserved: %#v", a.Decision)
	}
	if !a.Decision.DownstreamFlavorLedgerTestsAllowed {
		t.Fatalf("downstream diagnostics should be permitted under seal: %#v", a.Decision)
	}
	if a.Decision.NativeFlavorTheoremCertified || a.Decision.YukawaEigenvaluesDerived || a.Decision.CKMPMNSDerived || a.Decision.PMNSDerived || a.Decision.PhysicalParticlesAssigned || a.Decision.OfficialLedgerUpdateAllowed {
		t.Fatalf("overclaimed flavor results: %#v", a.Decision)
	}
	joined := strings.Join(appendAll(a.Supports, a.Failures, ComponentSupports(a.Components), ComponentFailures(a.Components)), "\n")
	for _, want := range append(RequiredSupports(), RequiredFailures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}

func TestGate964Theorem(t *testing.T) {
	res := Generation2ExternalFlavorOrientationSealInstallationAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{Verdict, Classification, ShortStatus, NextGate, "CONDITIONAL_SUPPORT_EXTERNAL_FLAVOR_ORIENTATION_SEAL_INSTALLED", "CONDITIONAL_SUPPORT_DOWNSTREAM_FLAVOR_LEDGER_TESTS_ALLOWED_UNDER_TRIPLE_SEAL", "FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_NOT_NATIVE_FLAVOR_THEOREM", "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES", "FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
