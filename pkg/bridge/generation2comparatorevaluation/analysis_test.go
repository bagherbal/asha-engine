package generation2comparatorevaluation

import (
	"math"
	"strings"
	"testing"
)

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.Sieve.AcceptedCount != 2 || a.Sieve.RejectedCount != 5 {
		t.Fatalf("unexpected sieve counts: %s", FormatSieve(a.Sieve))
	}
	if !a.Sieve.SyntheticInteriorAccepted || !a.Sieve.RedactedAccepted || !a.Sieve.AllAcceptedBridgeOnly {
		t.Fatalf("safe records not accepted bridge-only: %s", FormatSieve(a.Sieve))
	}
}

func TestEvaluateInverseAndDomainGuards(t *testing.T) {
	ok := Evaluate(ComparatorInput{Name: "ok", Sector: "up", ObservablePair: "{I_spec,I_K}", ValueKind: "synthetic", IK: 0.5, ISpec: 0.1, HasNumericPair: true, BridgeOnly: true})
	if !ok.Accepted || !ok.Evaluated || ok.PhaseBranches != GenericPhaseBranchCount {
		t.Fatalf("interior record should evaluate: %s", FormatEvaluation(ok))
	}
	if math.Abs(ok.Alpha-1) > 1e-12 {
		t.Fatalf("unexpected alpha: %s", FormatEvaluation(ok))
	}
	badIK := Evaluate(ComparatorInput{Name: "bad IK", ValueKind: "synthetic", IK: 1, HasNumericPair: true, BridgeOnly: true})
	if badIK.Accepted || badIK.Verdict != StatusFailedIKDomainRejected {
		t.Fatalf("IK boundary should be rejected: %s", FormatEvaluation(badIK))
	}
	badCos := Evaluate(ComparatorInput{Name: "bad cos", ValueKind: "synthetic", IK: 0, ISpec: 1, HasNumericPair: true, BridgeOnly: true})
	if badCos.Accepted || badCos.Verdict != StatusFailedPhaseCosDomainRejected {
		t.Fatalf("cosine domain should be rejected: %s", FormatEvaluation(badCos))
	}
	caustic := Evaluate(ComparatorInput{Name: "caustic", ValueKind: "synthetic", IK: 0, ISpec: 2 / (3 * math.Sqrt(3)), HasNumericPair: true, BridgeOnly: true})
	if caustic.Accepted || !caustic.Caustic || caustic.Verdict != StatusFailedCausticNotUnique {
		t.Fatalf("caustic should be flagged not accepted: %s", FormatEvaluation(caustic))
	}
}

func TestObservedAndNativePromotionRejected(t *testing.T) {
	observed := Evaluate(ComparatorInput{Name: "obs", ValueKind: "observed", IK: 0.1, ISpec: 0.02, HasNumericPair: true, ExplicitObservedImport: true, BridgeOnly: true})
	if observed.Accepted || observed.Verdict != StatusFailedObservedValueRejected {
		t.Fatalf("observed numeric record should be rejected in redacted harness: %s", FormatEvaluation(observed))
	}
	promote := Evaluate(ComparatorInput{Name: "promote", ValueKind: "synthetic", IK: 0.1, ISpec: 0.02, HasNumericPair: true, BridgeOnly: false, NativePromotionClaim: true})
	if promote.Accepted || promote.Verdict != StatusFailedNativePromotionAttempt {
		t.Fatalf("native promotion should be rejected: %s", FormatEvaluation(promote))
	}
}

func TestRenderAuditContainsGate458Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	md := RenderAudit(a)
	for _, want := range []string{
		"# Gate 458 Registry Audit",
		StatusBridgeOnlyExportValidated,
		"redacted/synthetic",
		"alpha = sqrt(3) I_K",
		StatusFailedObservedValueRejected,
		"no observed value",
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
