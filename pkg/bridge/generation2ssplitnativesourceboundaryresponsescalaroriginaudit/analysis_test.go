package generation2ssplitnativesourceboundaryresponsescalaroriginaudit

import (
	"strings"
	"testing"
)

func TestGate941BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		t.Fatalf("bad identity: %#v", a)
	}
	if a.CertificateIIPass || a.Transport.CertificatePassed || a.Transport.NativeCertified {
		t.Fatalf("Gate 941 must not pass Certificate II without native transport: %#v", a.Transport)
	}
	if a.Origin.Status != SourceBridgeStrongNotNative {
		t.Fatalf("expected bridge-strong origin status, got %s", a.Origin.Status)
	}
	if a.Response.AlphaLinear != AlphaLinear || a.Response.AlphaQuadratic != AlphaQuad || a.Response.AlphaTotal != AlphaB {
		t.Fatalf("alpha components changed: %#v", a.Response)
	}
	if !containsAll(a.Supports, Supports()) || !containsAll(a.Failures, Failures()) {
		t.Fatalf("missing supports or failures")
	}
}

func TestGate941Theorem(t *testing.T) {
	res := Generation2SSplitNativeSourceBoundaryResponseScalarOriginAuditTheorem().Verify()
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
		"D_base=(7/72)S_split",
		"AugmentedChamberDefectSplitToBoundaryPairResponseTransportTheorem",
		"CONDITIONAL_SUPPORT_S_SPLIT_HAS_PRIOR_AUGMENTED_CHAMBER_DEFECT_TRACE_ORIGIN",
		"CONDITIONAL_SUPPORT_BOUNDARY_PAIR_SYMMETRY_FORCES_UNIFORM_SCALAR_INSERTION",
		"FAILED_ROUTE_NO_NATIVE_AUGMENTED_CHAMBER_TO_B2_RESPONSE_TRANSPORT_THEOREM",
		"FAILED_ROUTE_CERTIFICATE_II_NOT_PASSED",
		NextGate,
	} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker/note %s", want)
		}
	}
}
