package generation2boundaryalphapowerresponsetransportmapaudit

import (
	"strings"
	"testing"
)

func TestGate867LinearLane(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Alpha.Linear.Power != 1 || a.Alpha.Linear.Numerator != PiTopRank || a.Alpha.Linear.Denominator != LinearDenominator || a.Alpha.Linear.TransportCertified || a.Alpha.Linear.ResponseOrderDerived {
		t.Fatalf("bad linear lane: %s", FormatLane(a.Alpha.Linear))
	}
}

func TestGate867QuadraticLane(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Alpha.Quadratic.Power != 2 || a.Alpha.Quadratic.Numerator != HRminRank || a.Alpha.Quadratic.Denominator != QuadraticDenominator || a.Alpha.Quadratic.TransportCertified || a.Alpha.Quadratic.ResponseOrderDerived {
		t.Fatalf("bad quadratic lane: %s", FormatLane(a.Alpha.Quadratic))
	}
}

func TestGate867AlphaShapeButNoTransport(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Alpha.ShapeCoherent || !near(a.Alpha.ReconstructedAlpha, AlphaB) || a.Alpha.Native || a.Alpha.TransportMapCertified {
		t.Fatalf("alpha overpromoted or malformed: %s", FormatAlpha(a.Alpha))
	}
}

func TestGate867SharedSAndPowerOrderObstruction(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.SharedS.SameCoordinateUsed || !a.SharedS.BothCodomainsTyped || a.SharedS.TransportIntoBothCertified || a.SharedS.PowerOrderDerived || !containsAll(a.SharedS.Failures, []string{FailureSameSNotTransportedToBoth, FailureLinearVsQuadraticNotDerived}) {
		t.Fatalf("bad shared S audit: %s", FormatSharedS(a.SharedS))
	}
}

func TestGate867DenominatorAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Denominators.LinearDenominator != 10 || a.Denominators.QuadraticDenominator != 72 || !a.Denominators.TypedBoundaryAugmentedDomains || a.Denominators.ActivationTheoremCertified {
		t.Fatalf("bad denominators: %s", FormatDenominators(a.Denominators))
	}
}

func TestGate867Firewalls(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !firewallsOK(a.Firewalls) || a.Impact.CanUpdateNEff || a.Impact.CanPromoteToR3 {
		t.Fatalf("firewalls broken: %s | %s", FormatFirewalls(a.Firewalls), FormatImpact(a.Impact))
	}
}

func TestGate867Theorem(t *testing.T) {
	res := Generation2BoundaryAlphaPowerResponseTransportMapAuditTheorem().Verify()
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
