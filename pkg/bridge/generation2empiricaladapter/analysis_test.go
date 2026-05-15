package generation2empiricaladapter

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.Sieve.AllowedCount != 4 || a.Sieve.RejectedCount != 5 {
		t.Fatalf("unexpected request sieve: %s", FormatSieve(a.Sieve))
	}
	if a.Schema.AllowsObservedValuesByDefault || a.Schema.AllowsNativeCoefficientExport {
		t.Fatalf("schema leaks forbidden path: %s", FormatSchema(a.Schema))
	}
	if a.Export.ActualObservedValueCount != 0 || a.Export.NativeExportCount != 0 {
		t.Fatalf("dry run exported forbidden values: %s", FormatExport(a.Export))
	}
}

func TestClassifyRejectsForbiddenRoutes(t *testing.T) {
	schema := buildSchema()
	cases := []AdapterRequest{
		{
			Name:                     "spectrum-only native coefficient claim",
			Operation:                "native-coefficient-from-spectrum",
			ValueMode:                ValueModeDummy,
			ComparatorCount:          1,
			HasSectorTag:             true,
			HasRenormalizationScale:  true,
			HasRenormalizationScheme: true,
			HasBridgeLabel:           true,
			ClaimsNativeCoefficient:  true,
		},
		{
			Name:                     "GST promotion",
			Operation:                "gst-native-law",
			ValueMode:                ValueModeDummy,
			ComparatorCount:          2,
			HasSectorTag:             true,
			HasRenormalizationScale:  true,
			HasRenormalizationScheme: true,
			HasBridgeLabel:           true,
			ClaimsGSTNative:          true,
		},
		{
			Name:                     "observed dry-run import",
			Operation:                "observed-local-ray-import",
			ValueMode:                ValueModeObserved,
			ComparatorCount:          2,
			HasSectorTag:             true,
			HasRenormalizationScale:  true,
			HasRenormalizationScheme: true,
			HasBridgeLabel:           true,
			ClaimsLocalRay:           true,
		},
	}
	for _, tc := range cases {
		got := classify(schema, tc)
		if got.Allowed {
			t.Fatalf("forbidden route accepted: %s", FormatRequest(got))
		}
	}
}

func TestClassifyAcceptsOnlyLabelledSymbolicBridge(t *testing.T) {
	schema := buildSchema()
	good := classify(schema, AdapterRequest{
		Name:                     "local dry run",
		Operation:                "local-ray-comparator",
		ValueMode:                ValueModeDummy,
		ComparatorCount:          2,
		HasSectorTag:             true,
		HasRenormalizationScale:  true,
		HasRenormalizationScheme: true,
		HasBridgeLabel:           true,
		UsesSpectrum:             true,
		ClaimsLocalRay:           true,
	})
	if !good.Allowed || good.Classification != "allowed-local-ray-dry-run" {
		t.Fatalf("labelled local bridge was rejected: %s", FormatRequest(good))
	}
	missing := classify(schema, AdapterRequest{
		Name:            "missing metadata",
		Operation:       "local-ray-comparator",
		ValueMode:       ValueModeDummy,
		ComparatorCount: 2,
		HasSectorTag:    true,
		HasBridgeLabel:  true,
		ClaimsLocalRay:  true,
	})
	if missing.Allowed || missing.Classification != "rejected-missing-metadata" {
		t.Fatalf("metadata failure did not fail closed: %s", FormatRequest(missing))
	}
}

func TestTheoremPasses(t *testing.T) {
	res := Generation2EmpiricalTextureAdapterDryRunFirewallTheorem().Run()
	if !res.Passed() {
		t.Fatalf("theorem audit checks did not pass:\n%s", res.Details())
	}
}

func TestRenderAuditContainsKeyStatuses(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{StatusDryRunFirewallValidated, StatusFailedSpectrumOnlyNativePromotionRejected, StatusFailedObservedValuesRejectedByDefault, "local dry run", "native coefficient export: forbidden"} {
		if !stringsContains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}

func stringsContains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
