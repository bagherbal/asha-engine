package generation2fourcertificatenativeclauseexecutionaudit

import (
	"strings"
	"testing"
)

func TestGate940BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		t.Fatalf("bad identity: %#v", a)
	}
	if a.FullNativeEligible || allNativeCertified(a.Certificates) {
		t.Fatalf("Gate 940 must not grant native R3 without all four certificates")
	}
	if len(a.Certificates) != 4 {
		t.Fatalf("expected four certificates, got %d", len(a.Certificates))
	}
	if countStatus(a.Certificates, CertificatePartialSupport) != 2 || countStatus(a.Certificates, CertificateBlocked) != 2 {
		t.Fatalf("expected two partial and two blocked certificates: %#v", a.Certificates)
	}
	if !containsAll(a.Supports, Supports()) || !containsAll(a.Failures, Failures()) {
		t.Fatalf("missing supports or failures")
	}
	if !containsAll(append(a.Failures, certificateFailures(a.Certificates)...), certificateFailures(a.Certificates)) {
		t.Fatalf("missing certificate failures")
	}
	if a.DiagnosticValues.AlphaB != AlphaB || a.DiagnosticValues.NEffOperator != NEffOperator {
		t.Fatalf("diagnostics changed: %#v", a.DiagnosticValues)
	}
}

func TestGate940Theorem(t *testing.T) {
	res := Generation2FourCertificateNativeClauseExecutionAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range []string{
		Verdict,
		Classification,
		ShortStatus,
		InheritedStatus,
		"Projector-generated admissible support lattice",
		"NativeSsplitBoundaryResponseParameterTheorem",
		"NativeBoundaryActivationMeasureTheorem",
		"NativeFullAFDescentTheorem or LawfulSpontaneousOrientationTheorem",
		"CONDITIONAL_SUPPORT_PROJECTOR_GENERATED_SUPPORT_LATTICE_CAN_BE_NATIVE_IN_FINITE_PROJECTOR_CATEGORY",
		"CONDITIONAL_SUPPORT_BOUNDARY_ACTIVATION_MEASURE_CAN_BE_READ_AS_FINITE_NORMALIZED_TRACE_RESPONSE",
		"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED_WITHOUT_ALL_FOUR_CERTIFICATES",
		"FAILED_ROUTE_NO_NATIVE_S_SPLIT_RESPONSE_PARAMETER_THEOREM_IF_SOURCE_NOT_CERTIFIED",
		"FAILED_ROUTE_NO_NATIVE_POST_ORIENTATION_R3_IF_FINITE_ONE_FORM_ORIENTATION_NOT_CERTIFIED",
		NextGate,
	} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker/note %s", want)
		}
	}
}
