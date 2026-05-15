package generation2empiricalimportswitch

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate465(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sieve.QuarantinedQuarkMassImportAccepted || !a.Sieve.QuarantinedCKMImportAccepted || a.Sieve.AcceptedCaseCount != 2 || a.Sieve.RejectedCaseCount != 12 {
		t.Fatalf("unexpected sieve closure: %+v", a.Sieve)
	}
	if !(a.Sieve.NativePromotionRejected && a.Sieve.NativeRegistryWriteRejected && a.Sieve.ClosedSwitchRejected && a.Sieve.MissingMetadataRejected && a.Sieve.ObservedDataAsTheoremRejected) {
		t.Fatalf("unsafe import routes were not rejected: %+v", a.Sieve)
	}
	if a.Firewall.EmpiricalDataInNativeRegistry || a.Firewall.NativePredictionFromEmpirical || a.Firewall.NativeLawFromEmpirical || a.Firewall.CKMMatrixNativePrediction || a.Firewall.QuarkMassNativePrediction {
		t.Fatalf("native firewall failed with airlock open: %+v", a.Firewall)
	}
}

func TestEvaluateImportAcceptsOnlySwitchOpenMetadatedBridgeRows(t *testing.T) {
	res, accepted, verdict, _ := EvaluateImport(validQuarkMassImport())
	if !accepted || verdict != StatusQuarantinedImportAccepted || !res.Imported || !res.Quarantined || !res.ComparatorLedgerWritten || res.NativeRegistryWritten {
		t.Fatalf("expected quarantined import, accepted=%t verdict=%s res=%+v", accepted, verdict, res)
	}
	res, accepted, verdict, _ = EvaluateImport(validCKMImport())
	if !accepted || verdict != StatusQuarantinedImportAccepted || !res.Imported || !res.Quarantined || res.NativePredictionLogged || res.NativeLawLogged {
		t.Fatalf("expected quarantined CKM import, accepted=%t verdict=%s res=%+v", accepted, verdict, res)
	}
}

func TestEvaluateImportRejectsUnsafeRoutes(t *testing.T) {
	cases := []struct {
		name string
		req  ImportRequest
		want string
	}{
		{"switch disabled", switchDisabledImport(), StatusFailedSwitchDisabled},
		{"missing source", missingSourceImport(), StatusFailedMissingMetadata},
		{"missing scale", missingScaleImport(), StatusFailedMissingMetadata},
		{"missing scheme", missingSchemeImport(), StatusFailedMissingMetadata},
		{"missing uncertainty", missingUncertaintyImport(), StatusFailedMissingUncertainty},
		{"missing bridge", missingBridgeOnlyImport(), StatusFailedMissingBridgeOnly},
		{"bad ledger", unsupportedLedgerImport(), StatusFailedUnsupportedLedger},
		{"native prediction", nativePredictionImport(), StatusFailedNativePromotion},
		{"native law", nativeLawImport(), StatusFailedNativePromotion},
		{"native registry", nativeRegistryWriteImport(), StatusFailedNativeRegistryWrite},
		{"CKM native", ckmNativePredictionImport(), StatusFailedCKMPMNSNativePrediction},
		{"observed theorem", observedDataAsTheoremImport(), StatusFailedObservedDataAsTheorem},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			_, accepted, verdict, _ := EvaluateImport(tc.req)
			if accepted || verdict != tc.want {
				t.Fatalf("expected %s, got accepted=%t verdict=%s", tc.want, accepted, verdict)
			}
		})
	}
}

func TestRenderAuditContainsGate465Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 465 Registry Audit", StatusEmpiricalImportSwitchValid, "empirical_import", StatusFailedNativePromotion, "quark-sector-comparator-ledger"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
