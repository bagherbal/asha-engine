package generation2chiralitymassbridgefirewallandboundaryrestpressurerelevanceaudit

import (
	"math"
	"strings"
	"testing"
)

func TestGate812ChiralityProjectorFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inheritance.Gate811Inherited || !a.Inheritance.BoundarySecondMomentSelected {
		t.Fatalf("bad inheritance: %+v", a.Inheritance)
	}
	if math.Abs(a.Inheritance.M2-1.624013231638281e-7) > 1e-20 || math.Abs(a.Inheritance.C2Obs-5.8299915722461693) > 5e-13 {
		t.Fatalf("bad Gate 811 numerical inheritance: %+v", a.Inheritance)
	}
	p := a.Pseudoscalar
	if p.NaiveProjectorsIdempotent || p.OmegaSquared != -1 || p.ComplexGammaSquared != 1 {
		t.Fatalf("bad pseudoscalar audit: %s", FormatPseudoscalar(p))
	}
	if math.Abs(p.PPlusSquaredOmega-0.5) > 1e-15 || math.Abs(p.PPlusSquaredScalar) > 1e-15 {
		t.Fatalf("bad P+ square: %s", FormatPseudoscalar(p))
	}
	if !containsAll(p.Failures, []string{StatusNaiveRealProjectorsInvalid, StatusNoNativeRealChirality, StatusComplexAirlockNotNative}) {
		t.Fatalf("missing projector failures: %s", FormatPseudoscalar(p))
	}
}

func TestGate812MassBridgeAndRestPressureFirewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !containsAll(a.HiggsMassBridge.Supports, []string{StatusMassBridgeTyped, StatusMassBridgeEdgeOnly}) {
		t.Fatalf("bad mass bridge supports: %s", FormatMassBridge(a.HiggsMassBridge))
	}
	if !containsAll(a.HiggsMassBridge.Failures, []string{StatusHiggsScalarNoYf, StatusMassBridgeNoEigenvalues, StatusEdgeNoDeltaN, StatusMassBridgeNoTopRest}) {
		t.Fatalf("bad mass bridge failures: %s", FormatMassBridge(a.HiggsMassBridge))
	}
	if !containsAll(a.TraceMagnitude.Failures, []string{StatusNoHermitianTraceOps, StatusNoPositiveTraceAtoms, StatusNoTopColorBlock, StatusNoRestPressureOperator}) {
		t.Fatalf("bad trace magnitude firewall: %s", FormatRequirement(a.TraceMagnitude))
	}
	if !containsAll(a.BoundaryFN.Failures, []string{StatusChiralityNoNineFive, StatusHiggsNoSixPS2, StatusNoBoundaryTraceMap, StatusNoPositiveRestSpectrum}) {
		t.Fatalf("bad boundary-FN firewall: %s", FormatBoundaryFN(a.BoundaryFN))
	}
}

func TestGate812TheoremAndStatusLedger(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Firewalls.Enforced || a.Firewalls.Verdict != StatusFirewallGate812 {
		t.Fatalf("bad firewalls: %+v", a.Firewalls)
	}
	if !strings.Contains(a.Branch.Next, "Gate 813") || !strings.Contains(a.Branch.Next, "Boundary Second-Moment") {
		t.Fatalf("bad branch: %+v", a.Branch)
	}
	res := Generation2ChiralityMassBridgeFirewallAndBoundaryRestPressureRelevanceAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %s %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
