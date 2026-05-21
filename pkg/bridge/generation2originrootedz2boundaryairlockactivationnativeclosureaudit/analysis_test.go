package generation2originrootedz2boundaryairlockactivationnativeclosureaudit

import (
	"strings"
	"testing"
)

func TestGate939BuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited != InheritedStatus || a.Classification != Classification || a.ShortStatus != ShortStatus || a.Truth != FinalTruth {
		t.Fatalf("bad identity: %#v", a)
	}
	if a.FullPassEligible {
		t.Fatalf("Gate 939 should not grant full native R3 without all four certificates")
	}
	if len(a.Clauses) != 4 || !allCollapsed(a.Clauses) || !allBlockFullNative(a.Clauses) || allNativeCertified(a.Clauses) {
		t.Fatalf("bad clause ledger: %#v", a.Clauses)
	}
	if !containsAll(a.Supports, Supports()) {
		t.Fatalf("supports missing")
	}
	if !containsAll(a.Failures, Failures()) {
		t.Fatalf("failures missing")
	}
	if !containsAll(append(a.Failures, clauseFailures(a.Clauses)...), clauseFailures(a.Clauses)) {
		t.Fatalf("clause failures missing")
	}
	if !containsAll(a.Failures, r4Failures(a.R4Boundary)) {
		t.Fatalf("R4 failures missing")
	}
}

func TestGate939Theorem(t *testing.T) {
	res := Generation2OriginRootedZ2BoundaryAirlockActivationNativeClosureAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range append(append(append(Statuses(), Supports()...), Failures()...), []string{
		FinalTruth,
		Classification,
		ShortStatus,
		InheritedStatus,
		MasterFunctorName,
		MasterFunctorID,
		"NativeAdmissibleAirlockSupportLatticeTheorem",
		"NativeSsplitBoundaryResponseParameterTheorem",
		"NativeBoundaryActivationMeasureTheorem",
		"NativeFullAFDescentTheorem or LawfulSpontaneousOrientationTheorem",
		"FAILED_ROUTE_NATIVE_R3_NOT_GRANTED_WITHOUT_ALL_FOUR_CERTIFICATES",
		"FAILED_ROUTE_NO_OFFICIAL_LEDGER_UPDATE_ALLOWED_WITHOUT_NATIVE_R3_CERTIFICATION",
		"RETIRED_FALSE_ROUTE: lambda versus barlambda representative",
		"R4_BOUNDARY: individual Yukawa values",
		NextGate,
	}...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker/note %s", want)
		}
	}
}
