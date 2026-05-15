package generation2sectordifference

import (
	"math"
	"strings"
	"testing"
)

func TestBuildDefaultGate462(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sieve.ValidUDDifferenceAccepted || a.Sieve.AcceptedCaseCount != 1 || a.Sieve.RejectedCaseCount != 8 {
		t.Fatalf("unexpected sieve closure: %+v", a.Sieve)
	}
	if !(a.Sieve.ObservedCKMPMNSRejected && a.Sieve.NativePredictionRejected && a.Sieve.NativeRelativePromotionRejected && a.Sieve.LeptonPMNSMisrouteRejected && a.Sieve.UniversalityNativeRejected) {
		t.Fatalf("unsafe routes were not rejected: %+v", a.Sieve)
	}
	if a.Firewall.CKMMatrixEntryComputed || a.Firewall.CKMMatrixEntryNative || !a.Firewall.NoObservedCKMImported || !a.Firewall.NoObservedPMNSImported {
		t.Fatalf("CKM/PMNS firewall failed: %+v", a.Firewall)
	}
}

func TestEvaluateRelativeRayComputesBridgeDiagnostic(t *testing.T) {
	rel, accepted, verdict, _ := EvaluateRelativeRay(validUDRays())
	if !accepted || verdict != StatusUDDifferenceBridgeOnlyComputed || !rel.BridgeOnly || rel.ExportsCKMEntry || rel.ExportsNativeObservable {
		t.Fatalf("unexpected relative ray: accepted=%t verdict=%s rel=%+v", accepted, verdict, rel)
	}
	if math.Abs(rel.DeltaAlpha-(-0.30)) > 1e-12 {
		t.Fatalf("unexpected delta alpha: %.16g", rel.DeltaAlpha)
	}
	if rel.ProjectiveDistance <= 0 || rel.PhaseChord <= 0 {
		t.Fatalf("expected nontrivial relative diagnostic: %+v", rel)
	}
}

func TestEvaluateRelativeRayRejectsUnsafeRoutes(t *testing.T) {
	cases := []struct {
		name string
		rays []SectorRay
		want string
	}{
		{"missing sector", []SectorRay{validRay("u", 0.20, 0.30, +1, 0)}, StatusFailedRequiresTwoProvenancedRays},
		{"missing provenance", missingProvenanceRays(), StatusFailedRequiresTwoProvenancedRays},
		{"missing eigenbasis", missingEigenbasisRays(), StatusFailedMissingEigenbasisConvention},
		{"observed", observedMixingRays(), StatusFailedObservedCKMPMNSImportRejected},
		{"native prediction", nativePredictionRays(), StatusFailedCKMPMNSPredictionRejected},
		{"native relative", nativeRelativePromotionRays(), StatusFailedNativeRelativeRayPromotion},
		{"lepton", leptonPMNSMisrouteRays(), StatusFailedLeptonSectorMisrouted},
		{"universality", nativeUniversalityRays(), StatusFailedUniversalityNotNative},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, accepted, verdict, _ := EvaluateRelativeRay(tc.rays)
			if accepted || verdict != tc.want {
				t.Fatalf("expected %s, got accepted=%t verdict=%s", tc.want, accepted, verdict)
			}
		})
	}
}

func TestRenderAuditContainsGate462Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 462 Registry Audit", StatusCKMInterfaceFirewallValidated, "Delta_alpha_ud", "Eigenbasis Convention Ledger"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
