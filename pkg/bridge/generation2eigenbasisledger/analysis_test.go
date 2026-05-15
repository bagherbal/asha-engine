package generation2eigenbasisledger

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate463(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sieve.ValidConventionAccepted || a.Sieve.AcceptedCaseCount != 1 || a.Sieve.RejectedCaseCount != 10 {
		t.Fatalf("unexpected sieve closure: %+v", a.Sieve)
	}
	if !(a.Sieve.RawPhaseGaugeRejected && a.Sieve.PermutationNativeRejected && a.Sieve.DegenerateSpectrumRejected && a.Sieve.KGenBasisRotationRejected && a.Sieve.ObservedCKMPMNSRejected && a.Sieve.NativePredictionRejected && a.Sieve.EigenbasisNativePromotionRejected && a.Sieve.MatrixExportRejected) {
		t.Fatalf("unsafe routes were not rejected: %+v", a.Sieve)
	}
	if a.Firewall.CKMMatrixEntryComputed || a.Firewall.CKMMatrixEntryNative || a.Firewall.ObservedCKMImported || a.Firewall.ObservedPMNSImported {
		t.Fatalf("CKM/PMNS firewall failed: %+v", a.Firewall)
	}
}

func TestEvaluateConventionPairAcceptsOnlyReadiness(t *testing.T) {
	res, accepted, verdict, _ := EvaluateConventionPair(validUDConventions())
	if !accepted || verdict != StatusConventionLedgerValidated || !res.BridgeOnly || !res.ConventionReady {
		t.Fatalf("unexpected convention result: accepted=%t verdict=%s res=%+v", accepted, verdict, res)
	}
	if res.CKMMatrixComputed || res.PMNSMatrixComputed || res.ExportsNativeObservable {
		t.Fatalf("convention ledger exported forbidden observable: %+v", res)
	}
}

func TestEvaluateConventionPairRejectsUnsafeRoutes(t *testing.T) {
	cases := []struct {
		name string
		rows []SectorConvention
		want string
	}{
		{"missing sector", []SectorConvention{validConvention("u")}, StatusFailedRequiresUDConventions},
		{"missing ordering", missingOrderingConventions(), StatusFailedMissingConvention},
		{"raw phase gauge", rawDiagonalizerConventions(), StatusFailedRawDiagonalizerPhaseGauge},
		{"native ordering", nativeOrderingConventions(), StatusFailedEigenvaluePermutation},
		{"degenerate", degenerateSpectrumConventions(), StatusFailedDegenerateSpectrum},
		{"K rotation", kgenRotationConventions(), StatusFailedKGenBasisRotation},
		{"observed", observedMixingConventions(), StatusFailedObservedCKMPMNSImport},
		{"native prediction", nativePredictionConventions(), StatusFailedCKMPMNSNativePrediction},
		{"native eigenbasis", nativeEigenbasisPromotionConventions(), StatusFailedEigenbasisNativePromotion},
		{"matrix export", matrixExportConventions(), StatusFailedConventionExportsMatrix},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, accepted, verdict, _ := EvaluateConventionPair(tc.rows)
			if accepted || verdict != tc.want {
				t.Fatalf("expected %s, got accepted=%t verdict=%s", tc.want, accepted, verdict)
			}
		})
	}
}

func TestRenderAuditContainsGate463Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 463 Registry Audit", StatusCKMNullAdapterPreconditionSet, "U(1)^3 x S3", "CKM Null Residual Adapter"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
