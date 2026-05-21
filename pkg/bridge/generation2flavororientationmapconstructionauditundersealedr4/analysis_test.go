package generation2flavororientationmapconstructionauditundersealedr4

import (
	"strings"
	"testing"
)

func TestGate962BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus || a.NextGate != NextGate {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.Decision.ExternalC3DomainAvailable || a.Decision.ExternalC3NativeGeneration {
		t.Fatalf("bad external C3 status: %#v", a.Decision)
	}
	if !a.Decision.U3FamilyGaugeFreedomDetected || !a.Decision.U3OrbitClassAvailable || a.Decision.CanonicalFlavorBasisCertified || a.Decision.CanonicalRepresentativeSelected || a.Decision.FlavorOrientationMapCertified {
		t.Fatalf("bad U3/orientation status: %#v", a.Decision)
	}
	if !a.Decision.AFOrientIsValidInterfaceTarget || a.Decision.AFOrientSuppliesFamilySelector {
		t.Fatalf("bad A_F orient status: %#v", a.Decision)
	}
	if !a.Decision.R3TracebodyAggregateTargetValid || a.Decision.R3RowsUsedAsGenerationLabels {
		t.Fatalf("bad R3 tracebody status: %#v", a.Decision)
	}
	if a.Decision.FlavorFormulaBacksolveAllowed || a.Decision.ObservedFlavorInputAllowed || a.Decision.PhysicalAssignmentAllowed || a.Decision.IndividualYukawaAllowed || a.Decision.CKMPMNSAllowed || a.Decision.OfficialLedgerUpdateAllowed {
		t.Fatalf("overclaimed flavor construction: %#v", a.Decision)
	}
	joined := strings.Join(appendAll(a.Supports, a.Failures, CandidateSupports(a.Candidates), CandidateFailures(a.Candidates)), "\n")
	for _, want := range append(RequiredSupports(), RequiredFailures()...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker %s", want)
		}
	}
}

func TestGate962Theorem(t *testing.T) {
	res := Generation2FlavorOrientationMapConstructionAuditUnderSealedR4Theorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{Verdict, Classification, ShortStatus, NextGate, "CONDITIONAL_SUPPORT_EXTERNAL_C3_DEFINES_FAMILY_ORIENTATION_ORBIT_UP_TO_U3", "FAILED_ROUTE_EXTERNAL_C3_HAS_NO_NATIVE_BASIS_OR_ORDERING", "FAILED_ROUTE_U3_ORBIT_CLASS_IS_NOT_FLAVOR_ORIENTATION_MAP", "FAILED_ROUTE_NO_CANONICAL_REPRESENTATIVE_SELECTED", "FAILED_ROUTE_FLAVOR_FORMULA_BACKSOLVE_IS_CIRCULAR"} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing note %s", want)
		}
	}
}
