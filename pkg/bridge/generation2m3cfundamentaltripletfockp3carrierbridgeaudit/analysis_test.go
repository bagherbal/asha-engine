package generation2m3cfundamentaltripletfockp3carrierbridgeaudit

import (
	"strings"
	"testing"
)

func TestGate833M3AndFockP3CarrierShapeSupportWithoutIdentification(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.M3Carrier.CanonicalCarrier || !a.M3Carrier.MatrixUnitsAudited || a.M3Carrier.Dim != 3 || a.M3Carrier.MatrixUnitCount != 9 || !a.M3Carrier.SuppliesCarrierProjectors {
		t.Fatalf("M3 carrier not audited: %s", FormatM3Carrier(a.M3Carrier))
	}
	if !a.FockP3.RankThreeCarrier || !a.FockP3.BMinusLSelectorTyped || a.FockP3.M3ActionCertified || a.FockP3.RepresentsM3 {
		t.Fatalf("Fock P3 over/under-certified: %s", FormatFockP3(a.FockP3))
	}
	if !a.ShapeBridge.DimensionMatches || !a.ShapeBridge.FormalIsomorphismsExist || a.ShapeBridge.CanonicalIntertwinerCertified || a.ShapeBridge.TypedBridgeCertified || a.ShapeBridge.TripletsIdentified {
		t.Fatalf("shape bridge over-promoted: %s", FormatShapeBridge(a.ShapeBridge))
	}
	if !containsAll(a.ShapeBridge.Failures, []string{FailureNoCanonicalM3ToFockP3Intertwiner, FailureM3ColorTripletNotIdentifiedWithFockP3, FailureShapeMatchOnlyNotTypedBridge}) {
		t.Fatalf("missing shape-bridge failures: %s", strings.Join(a.ShapeBridge.Failures, ","))
	}
}

func TestGate833IntertwinerRoutesAndTopI3CompatibilityObstructed(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Routes.FiniteAlgebraRouteAudited || !a.Routes.MoritaRouteAudited || !a.Routes.TraceRepresentationRouteAudited {
		t.Fatalf("routes not audited: %s", FormatRoutes(a.Routes))
	}
	if a.Routes.ActionLawCertified || a.Routes.CommutationLawCertified || a.Routes.RepresentationLawCertified || a.Routes.CanonicalIntertwinerCertified || a.Routes.CarrierBridgeCertified {
		t.Fatalf("intertwiner route over-certified: %s", FormatRoutes(a.Routes))
	}
	if !a.TopI3.IdentityShapeMatches || a.TopI3.SameAsM3FundamentalCertified || a.TopI3.SameAsFockP3Certified || a.TopI3.TopToM3ToP3ChainCertified {
		t.Fatalf("top I3 route over-certified: %s", FormatTopI3(a.TopI3))
	}
	if !containsAll(a.TopI3.Failures, []string{FailureTopI3NotCarrierCompatibleWithM3Fundamental, FailureTopI3NotIdentifiedWithFockP3}) {
		t.Fatalf("missing top-I3 failures: %s", strings.Join(a.TopI3.Failures, ","))
	}
}

func TestGate833ImpactFirewallsAndTheorem(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Impact.CarrierShapeSupport || a.Impact.CarrierBridgeCertified || a.Impact.SectorProjectorMapCertified || a.Impact.SectorTraceLedgerCertified || a.Impact.TraceMagnitudeReadoutCertified {
		t.Fatalf("impact over-promoted: %s", FormatImpact(a.Impact))
	}
	if a.Impact.CanPromoteToR3 || a.Impact.CanPromoteToR4 || a.Impact.CanUpdateNEff || a.Impact.CanUpdateCYukawa || a.Impact.CanUpdateCHiggs {
		t.Fatalf("ledger update or R3/R4 promotion allowed: %s", FormatImpact(a.Impact))
	}
	if !a.Firewalls.Enforced || !a.Firewalls.ShapeNotBridge || !a.Firewalls.NoM3P3Intertwiner || !a.Firewalls.NoM3ActionOnP3 || !a.Firewalls.TopI3Separated || !a.Firewalls.CarrierBridgeNotSectorLedger || !a.Firewalls.NoSectorProjectorMap || !a.Firewalls.NoMagnitudeReadout || !a.Firewalls.NotR3 || !a.Firewalls.NotR4 {
		t.Fatalf("firewall failed: %s", a.Firewalls.Verdict)
	}
	res := Generation2M3CFundamentalTripletFockP3CarrierBridgeAuditTheorem().Verify()
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
