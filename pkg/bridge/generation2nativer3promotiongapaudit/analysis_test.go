package generation2nativer3promotiongapaudit

import (
	"strings"
	"testing"
)

func TestGate938ABuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Inherited != InheritedStatus || a.Classification != Classification || a.ShortStatus != ShortStatus || a.Truth != FinalTruth {
		t.Fatalf("bad identity: %#v", a)
	}
	if len(a.Blockers) != 4 || !allPrimary(a.Blockers) {
		t.Fatalf("bad blocker ledger: %#v", a.Blockers)
	}
	if !containsAll(a.Supports, Supports()) {
		t.Fatalf("supports missing")
	}
	if !containsAll(a.Failures, Failures()) {
		t.Fatalf("failures missing")
	}
	if !containsAll(a.Failures, blockerFailures(a.Blockers)) {
		t.Fatalf("blocker failures missing")
	}
	if !containsAll(a.Failures, r4Failures(a.R4Boundary)) {
		t.Fatalf("R4 failures missing")
	}
}

func TestGate938ATheorem(t *testing.T) {
	res := Generation2NativeR3PromotionGapAuditTheorem().Verify()
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
		"Native BoundaryActivationMeasure theorem",
		"Native S_split response-parameter theorem",
		"Native admissible airlock support lattice theorem",
		"Full A_F descent or lawful spontaneous-orientation theorem",
		"RETIRED_PRIMARY_BLOCKER: lambda versus barlambda representative",
		"R4_BOUNDARY: generation carrier",
		NextGate,
	}...) {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing marker/note %s", want)
		}
	}
}
