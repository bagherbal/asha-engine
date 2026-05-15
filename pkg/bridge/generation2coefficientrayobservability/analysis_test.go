package generation2coefficientrayobservability

import "testing"

func TestBuildDefault(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatalf("BuildDefault failed: %v", err)
	}
	if a.Ray.ProjectiveDimension != 2 || a.Ray.NativeSelectorCount != 0 {
		t.Fatalf("unexpected ray model: %s", FormatRayModel(a.Ray))
	}
	if a.Rank.SpectrumOnlyRank != 1 || a.Rank.SpectrumOnlyResidualDOF != 1 {
		t.Fatalf("spectrum-only rank should be one with one residual coordinate: %s", FormatRankAudit(a.Rank))
	}
	if !a.Rank.TwoScalarLocalWorks || !a.Rank.CPBranchTagRequired {
		t.Fatalf("rank audit did not define local/CP branch protocol: %s", FormatRankAudit(a.Rank))
	}
	if a.Protocol.AllowsNativeCoefficientClaim || a.Protocol.AllowsSpectrumOnlyRayClaim {
		t.Fatalf("protocol leaked a native coefficient claim: %s", FormatProtocol(a.Protocol))
	}
}

func TestJacobianNonzeroGenerically(t *testing.T) {
	j := genericJacobian(0.7, 0.31)
	if j == 0 {
		t.Fatal("generic Jacobian sample unexpectedly zero")
	}
	if genericJacobian(0.7, 0) != 0 {
		t.Fatal("phase caustic phi=0 should have zero Jacobian determinant")
	}
}

func TestObservableMapsRespectFirewall(t *testing.T) {
	r := buildRankAudit()
	sawSpectrumOnly := false
	sawLocal := false
	sawForbidden := false
	for _, m := range r.Maps {
		if m.Name == "normalized spectrum only" {
			sawSpectrumOnly = true
			if m.Rank != 1 || m.LocallyIdentifiesRay {
				t.Fatalf("spectrum-only map incorrectly identifies ray: %s", FormatObservableMap(m))
			}
		}
		if m.Name == "spectrum plus K-addressed overlap" {
			sawLocal = true
			if m.Rank != 2 || !m.LocallyIdentifiesRay || m.NativePromotion {
				t.Fatalf("two-scalar map invalid: %s", FormatObservableMap(m))
			}
		}
		if m.NativePromotion && !m.AllowedByGate453 {
			sawForbidden = true
		}
		if m.NativePromotion && m.AllowedByGate453 {
			t.Fatalf("forbidden promotion allowed: %s", FormatObservableMap(m))
		}
	}
	if !sawSpectrumOnly || !sawLocal || !sawForbidden {
		t.Fatalf("missing expected observable-map cases: spectrum=%t local=%t forbidden=%t", sawSpectrumOnly, sawLocal, sawForbidden)
	}
}

func TestTheoremPasses(t *testing.T) {
	res := Generation2CoefficientRayObservabilityRankAuditTheorem().Run()
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
	for _, want := range []string{StatusComparatorProtocolDefined, StatusFailedSpectrumOnlyUnderdetermined, "I_spec", "I_K", "CP branch tag"} {
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
