package generation2flavorledgerdiagnosticpretestundertripleseal

import (
	"strings"
	"testing"
)

func TestGate965BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus || a.NextGate != NextGate {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.Decision.InheritedTripleSealActive || !a.Decision.R3DualSealPreserved || !a.Decision.ScalarSourceSealPreserved || !a.Decision.PostOrientationSealPreserved || !a.Decision.ExternalGenerationCarrierSealPreserved || !a.Decision.ExternalFlavorOrientationSealPreserved {
		t.Fatalf("triple seal not preserved: %#v", a.Decision)
	}
	if !a.Decision.EpsilonLedgerDiagnosticAllowed || !a.Decision.KappaLedgerDiagnosticAllowed || !a.Decision.KoideShadowDiagnosticAllowed || !a.Decision.CKMPMNSLedgerDiagnosticAllowed || !a.Decision.AllDiagnosticsSealedOnly {
		t.Fatalf("diagnostics not allowed as sealed tests: %#v", a.Decision)
	}
	if a.Decision.NativeFlavorTheoremDerived || a.Decision.YukawaSpectrumDerived || a.Decision.CKMPNMSTheoremDerived || a.Decision.PMNSTheoremDerived || a.Decision.PhysicalParticlesAssigned || a.Decision.OfficialLedgerUpdateAllowed {
		t.Fatalf("overclaimed flavor/native result: %#v", a.Decision)
	}
	for _, d := range a.Diagnostics {
		if d.UsedAsNativeTheorem || d.UsedAsGenerationSource || d.UsedAsFlavorOrientationSource || d.DerivesIndividualYukawas || d.DerivesCKMPMNS || d.DerivesPMNS || d.AssignsPhysicalParticles || d.UpdatesOfficialLedger {
			t.Fatalf("diagnostic overclaimed native/source role: %#v", d)
		}
	}
	joined := strings.Join(appendAll(a.Supports, a.Failures, DiagnosticSupports(a.Diagnostics), DiagnosticFailures(a.Diagnostics)), "\n")
	for _, want := range append(RequiredSupports(), RequiredFailures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}

func TestGate965Theorem(t *testing.T) {
	res := Generation2FlavorLedgerDiagnosticPretestUnderTripleSealTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{Verdict, Classification, ShortStatus, NextGate, "CONDITIONAL_SUPPORT_EPSILON_E_LEDGER_TEST_ALLOWED_UNDER_TRIPLE_SEAL", "CONDITIONAL_SUPPORT_KAPPA_LEDGER_TEST_ALLOWED_UNDER_TRIPLE_SEAL", "CONDITIONAL_SUPPORT_KOIDE_SHADOW_COMPATIBILITY_TEST_ALLOWED_UNDER_TRIPLE_SEAL", "CONDITIONAL_SUPPORT_CKM_PMNS_LEDGER_COMPATIBILITY_TEST_ALLOWED_UNDER_TRIPLE_SEAL", "FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED", "FAILED_ROUTE_DOWNSTREAM_LEDGER_TESTS_ARE_NOT_NATIVE_DERIVATIONS"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
