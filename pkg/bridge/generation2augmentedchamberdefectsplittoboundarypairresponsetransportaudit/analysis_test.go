package generation2augmentedchamberdefectsplittoboundarypairresponsetransportaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate942BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited != InheritedStatus || a.Verdict != Verdict || a.Classification != Classification || a.ShortStatus != ShortStatus {
		t.Fatalf("bad identity: %#v", a)
	}
	if a.Carrier.Lambda4Rank+a.Carrier.BoundaryRank != H72Rank || a.Carrier.TotalRank != H72Rank {
		t.Fatalf("bad carrier ranks: %#v", a.Carrier)
	}
	if a.Interface.Status != TransportStronglySourceTyped {
		t.Fatalf("expected strongly source-typed interface, got %s", a.Interface.Status)
	}
	if a.CertificateIIPassed || a.Identification.CertificateIIPassed || a.Identification.NativeCertified {
		t.Fatalf("Gate 942 must not pass Certificate II without native descent: %#v", a.Identification)
	}
	if math.Abs(a.AlphaQuadratic-AlphaQuad) > 1e-18 || math.Abs(a.AlphaTotal-AlphaB) > 1e-18 {
		t.Fatalf("alpha values changed: quad %.19g total %.19g", a.AlphaQuadratic, a.AlphaTotal)
	}
	if !containsAll(a.Supports, Supports()) || !containsAll(a.Failures, Failures()) {
		t.Fatalf("missing supports or failures")
	}
}

func TestGate942Theorem(t *testing.T) {
	res := Generation2AugmentedChamberDefectSplitToBoundaryPairResponseTransportAuditTheorem().Verify()
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
		"H72 = Lambda^4 V8 plus B2",
		"D_base=(7/72)S_split",
		"alpha_quad=(7/72)S_split^2",
		"CONDITIONAL_SUPPORT_B2_IS_SHARED_BOUNDARY_INTERFACE_OF_H72",
		"CONDITIONAL_SUPPORT_S_SPLIT_TRANSPORTS_ONCE_INTO_EACH_BOUNDARY_FACTOR",
		"CONDITIONAL_SUPPORT_CERTIFICATE_II_STRENGTHENED_BUT_NOT_PASSED",
		"FAILED_ROUTE_NO_NATIVE_H72_DEFECT_TO_B2_RESPONSE_DESCENT_MAP",
		"FAILED_ROUTE_CERTIFICATE_II_NOT_PASSED",
		NextGate,
	} {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker/note %s", want)
		}
	}
}
