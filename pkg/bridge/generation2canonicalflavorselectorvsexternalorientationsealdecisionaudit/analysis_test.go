package generation2canonicalflavorselectorvsexternalorientationsealdecisionaudit

import (
	"strings"
	"testing"
)

func TestGate963BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus || a.NextGate != NextGate {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.Decision.InheritedU3OrbitOnly || !a.Decision.U3FamilyGaugeFreedomRetained || a.Decision.CurrentASHADataBreaksU3Gauge {
		t.Fatalf("bad U3 inherited obstruction status: %#v", a.Decision)
	}
	if a.Decision.CanonicalFlavorSelectorFound || a.Decision.CanonicalRepresentativeSelected {
		t.Fatalf("canonical selector/representative should not be certified: %#v", a.Decision)
	}
	if !a.Decision.ExternalFlavorOrientationSealRequired || !a.Decision.ExternalFlavorOrientationSealCanSelectRepresentative || a.Decision.ExternalFlavorOrientationSealNative {
		t.Fatalf("external flavor orientation seal status wrong: %#v", a.Decision)
	}
	if !a.Decision.DownstreamFlavorLedgerTestsAllowedUnderSeal {
		t.Fatalf("downstream tests must be allowed only under seal: %#v", a.Decision)
	}
	if a.Decision.NativeFlavorTheoremCertified || a.Decision.YukawaEigenvaluesDerived || a.Decision.PhysicalParticlesAssigned || a.Decision.CKMPMNSDerived || a.Decision.OfficialLedgerUpdateAllowed {
		t.Fatalf("overclaimed downstream flavor results: %#v", a.Decision)
	}
	if !a.Decision.R3DualSealPreserved || !a.Decision.ExternalGenerationCarrierSealPreserved {
		t.Fatalf("inherited seals not preserved: %#v", a.Decision)
	}
	joined := strings.Join(appendAll(a.Supports, a.Failures, CandidateSupports(a.Candidates), CandidateFailures(a.Candidates)), "\n")
	for _, want := range append(RequiredSupports(), RequiredFailures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}

func TestGate963Theorem(t *testing.T) {
	res := Generation2CanonicalFlavorSelectorVsExternalOrientationSealDecisionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{Verdict, Classification, ShortStatus, NextGate, "CONDITIONAL_SUPPORT_EXTERNAL_C3_RETAINS_U3_FAMILY_GAUGE_FREEDOM", "CONDITIONAL_SUPPORT_EXTERNAL_FLAVOR_ORIENTATION_SEAL_CAN_SELECT_REPRESENTATIVE", "FAILED_ROUTE_NO_CANONICAL_FLAVOR_SELECTOR_IN_CURRENT_CERTIFICATE", "FAILED_ROUTE_U3_FAMILY_GAUGE_NOT_BROKEN_BY_CURRENT_ASHA_DATA", "FAILED_ROUTE_EXTERNAL_FLAVOR_ORIENTATION_SEAL_NOT_NATIVE_FLAVOR_THEOREM", "FAILED_ROUTE_NO_INDIVIDUAL_YUKAWA_VALUES"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
