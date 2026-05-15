package generation2branchresiduals

import (
	"strings"
	"testing"
)

func TestBuildDefaultGate460(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Sieve.RedactedPreserved || !a.Sieve.SyntheticInteriorAccepted {
		t.Fatalf("expected redacted and synthetic acceptances: %+v", a.Sieve)
	}
	if !(a.Sieve.IncompleteTagRejected && a.Sieve.CausticRejected && a.Sieve.ObservedDataRejected && a.Sieve.NativePromotionRejected && a.Sieve.ProjectiveDomainRejected && a.Sieve.PhaseCosDomainRejected) {
		t.Fatalf("expected all fail-closed routes: %+v", a.Sieve)
	}
	if !a.Firewall.NoObservedMuonMassImported || !a.Firewall.NoCoefficientRayPromotion || !a.Firewall.NoPhaseBranchPromotion {
		t.Fatalf("firewall failed: %+v", a.Firewall)
	}
}

func TestEvaluateResidualSyntheticClosesResiduals(t *testing.T) {
	e := EvaluateResidual(ResidualRequest{Name: "unit", IK: 0.25, ISpec: 0.08, HasNumericPair: true, BridgeOnly: true, CPOddSign: +1, HasCPOddSign: true, C3Sheet: 1, HasC3Sheet: true})
	if !e.Accepted || !e.Evaluated || e.Verdict != StatusSyntheticBranchResidualEvaluated {
		t.Fatalf("unexpected evaluation: %+v", e)
	}
	if !nearlyZero(e.M22Residual) || !nearlyZero(e.IKResidual) || !nearlyZero(e.ISpecResidual) || e.CPSignResidual != 0 || e.C3SheetResidual != 0 {
		t.Fatalf("residuals did not close: %+v", e)
	}
}

func TestEvaluateResidualRejectsUnsafeRoutes(t *testing.T) {
	cases := []struct {
		name string
		req  ResidualRequest
		want string
	}{
		{"missing tag", ResidualRequest{Name: "x", IK: 0.25, ISpec: 0.08, HasNumericPair: true, BridgeOnly: true}, StatusFailedIncompleteBranchTag},
		{"observed", ResidualRequest{Name: "x", IK: 0.25, ISpec: 0.08, HasNumericPair: true, ExplicitObservedData: true, BridgeOnly: true, CPOddSign: 1, HasCPOddSign: true, C3Sheet: 0, HasC3Sheet: true}, StatusFailedObservedDataRejected},
		{"native", ResidualRequest{Name: "x", IK: 0.25, ISpec: 0.08, HasNumericPair: true, NativePromotionClaim: true, BridgeOnly: false, CPOddSign: 1, HasCPOddSign: true, C3Sheet: 0, HasC3Sheet: true}, StatusFailedNativePromotionRejected},
		{"projective", ResidualRequest{Name: "x", IK: 1, ISpec: 0, HasNumericPair: true, BridgeOnly: true, CPOddSign: 1, HasCPOddSign: true, C3Sheet: 0, HasC3Sheet: true}, StatusFailedProjectiveDomainRejected},
		{"phase", ResidualRequest{Name: "x", IK: 0.25, ISpec: 99, HasNumericPair: true, BridgeOnly: true, CPOddSign: 1, HasCPOddSign: true, C3Sheet: 0, HasC3Sheet: true}, StatusFailedPhaseCosDomainRejected},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			e := EvaluateResidual(tc.req)
			if e.Accepted || e.Verdict != tc.want {
				t.Fatalf("expected %s, got %+v", tc.want, e)
			}
		})
	}
}

func TestRenderAuditContainsGate460Language(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	md := RenderAudit(a)
	for _, want := range []string{"# Gate 460 Registry Audit", StatusBridgeOnlyResidualExport, "R_22 = M_22 = 0", "Three-Sector Comparator Multiplex"} {
		if !strings.Contains(md, want) {
			t.Fatalf("audit missing %q", want)
		}
	}
}
