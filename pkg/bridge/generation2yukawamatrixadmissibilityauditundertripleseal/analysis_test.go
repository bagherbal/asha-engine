package generation2yukawamatrixadmissibilityauditundertripleseal

import (
	"strings"
	"testing"
)

func TestGate971BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus || a.NextGate != NextGate {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.Decision.InheritedSealedRail || !a.Decision.R3DualSealPreserved || !a.Decision.ScalarSourceSealPreserved || !a.Decision.PostOrientationSealPreserved || !a.Decision.ExternalGenerationCarrierSealPreserved || !a.Decision.ExternalFlavorOrientationSealPreserved {
		t.Fatalf("seals not preserved: %#v", a.Decision)
	}
	if a.Decision.DerivesNativeFlavor || a.Decision.DerivesNativeYukawaMatrix || a.Decision.DerivesIndividualYukawas || a.Decision.DerivesCKMPMNS || a.Decision.AssignsPhysicalParticles || a.Decision.UpdatesOfficialLedger {
		t.Fatalf("overclaimed native/physical result: %#v", a.Decision)
	}
	joined := strings.Join(append(append(append(a.Supports, a.Failures...), a.Allowed...), a.MatrixNormalForm...), "\n")
	for _, want := range append(RequiredSupports(), RequiredFailures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}

func TestGate971Theorem(t *testing.T) {
	res := Generation2YukawaMatrixAdmissibilityAuditUnderTripleSealTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{Verdict, Classification, ShortStatus, NextGate} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
