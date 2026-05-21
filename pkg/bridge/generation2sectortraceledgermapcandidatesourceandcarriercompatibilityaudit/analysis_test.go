package generation2sectortraceledgermapcandidatesourceandcarriercompatibilityaudit

import (
	"strings"
	"testing"
)

func TestGate832FiniteSectorProjectorsAreNotMagnitudes(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.FiniteSector.StrongestLawfulSource || !a.FiniteSector.SectorProjectorCandidatesAudited || !a.FiniteSector.SuppliesGaugeInternalCarriers {
		t.Fatalf("finite sector source not audited: %s", FormatFiniteSector(a.FiniteSector))
	}
	if a.FiniteSector.SuppliesYukawaMagnitudes || a.FiniteSector.SuppliesTraceMagnitudeReadout || a.FiniteSector.UsesObservedYukawaData {
		t.Fatalf("finite sector source over-promoted: %s", FormatFiniteSector(a.FiniteSector))
	}
	if !containsAll(a.FiniteSector.Failures, []string{FailureSectorProjectorsDoNotSupplyMagnitudes, FailureSectorProjectorNotYukawaValue, FailureAFSectorIdempotentNotYukawaMagnitude}) {
		t.Fatalf("missing finite-sector failures: %s", strings.Join(a.FiniteSector.Failures, ","))
	}
}

func TestGate832CarrierCompatibilityK7AndDualTripletObstructions(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Compatibility.CarrierMapCertified || a.Compatibility.CarrierCompatibilityCertified || a.Compatibility.FiniteAlgebraCommutationProven || a.Compatibility.Compatible {
		t.Fatalf("aggregate carrier compatibility over-certified: %s", FormatCompatibility(a.Compatibility))
	}
	if !a.Compatibility.NonCircular || !containsAll(a.Compatibility.Failures, []string{FailureAggregateNotCompatibleWithSectorProjectors, FailureNoTypedAggregateToSectorMap, FailureNoCarrierCompatibilityTheorem}) {
		t.Fatalf("missing compatibility obstruction: %s", FormatCompatibility(a.Compatibility))
	}
	if !a.K7Route.CountMatchesK7 || a.K7Route.TypedMapCertified || a.K7Route.ProjectorIdentityCertified || a.K7Route.PromotedToK7 {
		t.Fatalf("K7 route over-promoted: %s", FormatK7Route(a.K7Route))
	}
	if !a.DualTriplet.FiniteAlgebraRouteAudited || !a.DualTriplet.MoritaRouteAudited || !a.DualTriplet.TraceRepresentationRouteAudited || a.DualTriplet.TypedBridgeCertified || a.DualTriplet.TripletsIdentified {
		t.Fatalf("dual-triplet route over-promoted: %s", FormatDualTriplet(a.DualTriplet))
	}
}

func TestGate832ImpactFirewallsAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("impact over-promoted: %s", FormatImpact(a.Impact))
	}
	if !strings.Contains(a.Impact.NextMissingObject, "SectorProjectorMap") || !strings.Contains(a.Impact.NextMissingObject, "SectorTraceMagnitudeReadoutMap") {
		t.Fatalf("next missing object not sharpened: %s", FormatImpact(a.Impact))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.AFProjectorsNotMagnitudes || !a.Firewalls.AggregateNotSectorLedger || !a.Firewalls.NoCarrierMap || !a.Firewalls.NoSectorLedgerMap || !a.Firewalls.NoMagnitudeReadout || !a.Firewalls.DualTripletSeparated || !a.Firewalls.SevenNotK7 || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 {
		t.Fatalf("firewall failed: %s", a.Firewalls.Verdict)
	}
	res := Generation2SectorTraceLedgerMapCandidateSourceAndCarrierCompatibilityAuditTheorem().Verify()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected theorem construction failure: %+v", res)
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
