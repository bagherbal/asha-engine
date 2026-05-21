package generation2h72defectscalartob2boundaryresponsedescentmapaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate943BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		t.Fatalf("bad identity: %#v", a)
	}
	if !a.Projection.Canonical || a.Projection.Boundary != "B2" || Lambda4V8Rank+B2Rank != H72Rank {
		t.Fatalf("bad projection: %#v", a.Projection)
	}
	if !a.Restriction.ForcesSEquals || a.Restriction.NativeSourceKnown {
		t.Fatalf("bad scalar restriction: %#v", a.Restriction)
	}
	if !a.Insertion.Uniform || a.Insertion.SecondTransportReq {
		t.Fatalf("bad insertion: %#v", a.Insertion)
	}
	if !a.Certificate.TransportComponentSupported || a.Certificate.NativeSourceCertified || a.Certificate.CertificatePassed {
		t.Fatalf("bad certificate II status: %#v", a.Certificate)
	}
	if math.Abs(a.AlphaQuadratic-AlphaQuad) > 1e-18 || math.Abs(a.AlphaTotal-AlphaB) > 1e-18 {
		t.Fatalf("alpha values changed: quad %.19g total %.19g", a.AlphaQuadratic, a.AlphaTotal)
	}
	if !containsAll(a.Supports, Supports()) || !containsAll(a.Failures, Failures()) {
		t.Fatalf("missing supports or failures")
	}
}

func TestGate943Theorem(t *testing.T) {
	res := Generation2H72DefectScalarToB2BoundaryResponseDescentMapAuditTheorem().Verify()
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
		"pi_B : H72 -> B2",
		"S_split * I_B2",
		"CONDITIONAL_SUPPORT_CENTRAL_SCALAR_DESCENT_FORCES_S_EQUALS_S_SPLIT",
		"CONDITIONAL_SUPPORT_S_SPLIT_TRANSPORT_COMPONENT_OF_CERTIFICATE_II_STRONGLY_SUPPORTED",
		"FAILED_ROUTE_NATIVE_STATUS_OF_S_SPLIT_SOURCE_NOT_CERTIFIED",
		"FAILED_ROUTE_CERTIFICATE_II_NOT_FULLY_PASSED",
		NextGate,
	} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker/note %s", want)
		}
	}
}
