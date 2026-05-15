package generation2sectormultiplex

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate461(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sieve.IndependentThreeSectorAccepted || !a.Sieve.LabelledBridgeUniversalityAccepted {
		t.Fatalf("expected independent and labelled bridge cases: %+v", a.Sieve)
	}
	if !(a.Sieve.MissingSectorRejected && a.Sieve.NativeUniversalityRejected && a.Sieve.UnlabelledUniversalityRejected && a.Sieve.ObservedDataRejected && a.Sieve.NativePromotionRejected && a.Sieve.SectorContaminationRejected) {
		t.Fatalf("expected all fail-closed routes: %+v", a.Sieve)
	}
	if !a.Firewall.NoCrossSectorUniversalityLaw || !a.Firewall.NoCoefficientRayPromotion || a.Dimensions.SectorRayUniversalityNative {
		t.Fatalf("firewall failed: %+v %+v", a.Firewall, a.Dimensions)
	}
}

func TestEvaluateSectorRecordAcceptsBridgeOnlySynthetic(t *testing.T) {
	e := EvaluateSectorRecord(validRecord("u", 0.20, 0.070, +1, 0))
	if !e.Accepted || !e.Evaluated || e.Verdict != StatusIndependentSectorRaysValid || !e.BridgeOnlyExport {
		t.Fatalf("unexpected evaluation: %+v", e)
	}
	if e.Alpha == 0 || e.Phi == 0 {
		t.Fatalf("expected nontrivial synthetic ray: %+v", e)
	}
}

func TestClassifyRejectsUnsafeRoutes(t *testing.T) {
	cases := []struct {
		name    string
		records []SectorRecord
		want    string
	}{
		{"missing", missingSectorRecords(), StatusFailedMissingSector},
		{"native universality", nativeUniversalityRecords(), StatusFailedNativeUniversality},
		{"unlabelled universality", unlabelledSharedRecords(), StatusFailedUnlabelledUniversality},
		{"observed", observedRecords(), StatusFailedObservedDataRejected},
		{"native promotion", nativePromotionRecords(), StatusFailedNativePromotionRejected},
		{"contamination", contaminatedRecords(), StatusFailedSectorContamination},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			evals := evaluateLedger(tc.records)
			accepted, verdict, _ := classifyCase(tc.name, tc.records, evals)
			if accepted || verdict != tc.want {
				t.Fatalf("expected %s, got accepted=%t verdict=%s evals=%+v", tc.want, accepted, verdict, evals)
			}
		})
	}
}

func TestRenderAuditContainsGate461Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 461 Registry Audit", StatusMultiplexBridgeOnlyValidated, "{u,d,e}", "Sector-Difference Invariant"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
