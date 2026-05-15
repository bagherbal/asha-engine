package generation2ckmnullresidual

import (
	"math"
	"strings"
	"testing"
)

func TestBuildDefaultGate464(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sieve.ValidSyntheticResidualAccepted || a.Sieve.AcceptedCaseCount != 1 || a.Sieve.RejectedCaseCount != 11 {
		t.Fatalf("unexpected sieve closure: %+v", a.Sieve)
	}
	if !(a.Sieve.ObservedCKMPMNSRejected && a.Sieve.NativePredictionRejected && a.Sieve.MatrixExportRejected && a.Sieve.RawDiagonalizerRejected && a.Sieve.GSTSelectorRejected) {
		t.Fatalf("unsafe routes were not rejected: %+v", a.Sieve)
	}
	if a.Firewall.CKMMatrixConstructed || a.Firewall.CKMMatrixEntryComputed || a.Firewall.CKMMatrixEntryNative || a.Firewall.ObservedCKMImported || a.Firewall.ObservedPMNSImported {
		t.Fatalf("CKM firewall failed: %+v", a.Firewall)
	}
}

func TestEvaluateCKMNullResidualComputesOnlyDistance(t *testing.T) {
	in := validInput()
	res, accepted, verdict, _ := EvaluateCKMNullResidual(in)
	if !accepted || verdict != StatusCKMNullResidualComputed || !res.BridgeOnly || !res.SyntheticOnly || !res.ExportsRelativeDiagnostic {
		t.Fatalf("unexpected residual result: accepted=%t verdict=%s res=%+v", accepted, verdict, res)
	}
	wantDAlpha := in.AlphaD - in.AlphaU
	wantDPhi := wrapPi(in.PhiD - in.PhiU)
	wantD := math.Sqrt(wantDAlpha*wantDAlpha + 4*math.Sin(wantDPhi/2)*math.Sin(wantDPhi/2))
	if math.Abs(res.DeltaAlpha-wantDAlpha) > 1e-12 || math.Abs(res.DeltaPhi-wantDPhi) > 1e-12 || math.Abs(res.ProjectiveRayDistance-wantD) > 1e-12 {
		t.Fatalf("bad residual distance: got %+v want Δα=%g Δφ=%g d=%g", res, wantDAlpha, wantDPhi, wantD)
	}
	if res.CKMMatrixConstructed || res.CKMEntryComputed || res.PMNSEntryComputed || res.ExportsNativeObservable || res.ObservedDataImported {
		t.Fatalf("residual exported forbidden observable: %+v", res)
	}
}

func TestEvaluateCKMNullResidualRejectsUnsafeRoutes(t *testing.T) {
	cases := []struct {
		name string
		in   CKMNullInput
		want string
	}{
		{"missing ray", missingRelativeRayInput(), StatusFailedRequiresRelativeRay},
		{"missing eigenbasis", missingEigenbasisInput(), StatusFailedRequiresEigenbasisConvention},
		{"missing branch provenance", missingBranchProvenanceInput(), StatusFailedRequiresBranchAndProvenance},
		{"observed", observedMixingInput(), StatusFailedObservedCKMPMNSImport},
		{"native prediction", nativePredictionInput(), StatusFailedCKMPMNSNativePrediction},
		{"matrix export", matrixExportInput(), StatusFailedCKMMatrixExport},
		{"raw", rawDiagonalizerInput(), StatusFailedRawDiagonalizer},
		{"degenerate", degenerateSpectrumInput(), StatusFailedDegenerateSpectrum},
		{"K rotation", kgenRotationInput(), StatusFailedKGenBasisRotation},
		{"native residual", nativeResidualPromotionInput(), StatusFailedNativeResidualPromotion},
		{"GST selector", gstSelectorInput(), StatusFailedGSTAsCKMSelector},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, accepted, verdict, _ := EvaluateCKMNullResidual(tc.in)
			if accepted || verdict != tc.want {
				t.Fatalf("expected %s, got accepted=%t verdict=%s", tc.want, accepted, verdict)
			}
		})
	}
}

func TestRenderAuditContainsGate464Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 464 Registry Audit", StatusCKMNullResidualFirewallValidated, "d_ud", "not V_CKM"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
