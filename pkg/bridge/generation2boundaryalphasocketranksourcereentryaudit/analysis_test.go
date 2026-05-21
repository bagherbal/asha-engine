package generation2boundaryalphasocketranksourcereentryaudit

import (
	"strings"
	"testing"
)

func TestGate866LinearSocketRankLane(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Alpha.LinearLane.Numerator != PiTopRank || a.Alpha.LinearLane.Denominator != LinearDenominator || !near(a.Alpha.LinearLane.Coefficient, 0.3) || a.Alpha.LinearLane.NativeTransport {
		t.Fatalf("bad linear lane: %s", FormatRankSource(a.Alpha.LinearLane))
	}
}

func TestGate866QuadraticSocketRankLane(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Alpha.QuadraticLane.Numerator != HRminRank || a.Alpha.QuadraticLane.Denominator != QuadraticDenominator || !near(a.Alpha.QuadraticLane.Coefficient, 7.0/72.0) || a.Alpha.QuadraticLane.NativeTransport {
		t.Fatalf("bad quadratic lane: %s", FormatRankSource(a.Alpha.QuadraticLane))
	}
}

func TestGate866AlphaReconstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Alpha.ReconstructsAlpha || !near(a.Alpha.ReconstructedAlpha, AlphaB) || a.Alpha.Native || a.Alpha.TransportMapCertified {
		t.Fatalf("bad alpha reconstruction: %s", FormatAlpha(a.Alpha))
	}
}

func TestGate866DualSevenFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.DualSeven.SameInteger || a.DualSeven.Identified || a.DualSeven.TypedMapCertified || !containsAll(a.DualSeven.Failures, []string{FailureNoHRMinToK7Map, FailureDualSevenNotIdentified}) {
		t.Fatalf("bad dual seven firewall: %s", FormatDualSeven(a.DualSeven))
	}
}

func TestGate866TransportObstructionAndLedgerFreeze(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Obstruction.ActivationMapCertified || a.Obstruction.AlphaNative || a.R3.EligibleForR3 || a.Impact.CanUpdateNEff || a.Impact.CanPromoteToR3 {
		t.Fatalf("overpromoted: %s | %s | %s", FormatObstruction(a.Obstruction), FormatR3(a.R3), FormatImpact(a.Impact))
	}
}

func TestGate866Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) {
		t.Fatalf("firewalls broken: %s", FormatFirewalls(a.Firewalls))
	}
}

func TestGate866Theorem(t *testing.T) {
	res := Generation2BoundaryAlphaSocketRankSourceReEntryAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem failure: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("failed check %s: %s", c.Name, c.Detail)
		}
	}
	joined := strings.Join(res.Notes, "\n")
	for _, want := range Statuses() {
		if !strings.Contains(joined, want) {
			t.Fatalf("missing status %s", want)
		}
	}
}
