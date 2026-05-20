package generation2witthopfs7contactreebaudit

import "testing"

func TestGate570WittFockHopfS7ContactAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Carrier.HasComplexStructureJ || !a.Carrier.J2EqualsMinusI || !a.Carrier.HasPositiveHermitianMetric || a.Carrier.ComplexDimension != 4 || a.Carrier.RealDimension != 8 {
		t.Fatalf("bad carrier: %s", FormatCarrier(a.Carrier))
	}
	if a.Sphere.SphereRealDimension != 7 || a.Sphere.IdentifiedWithK7 {
		t.Fatalf("bad sphere: %s", FormatSphere(a.Sphere))
	}
	if !a.Contact.ContactVolumeNonzero || a.Contact.ContactVolumeAtBasepoint == 0 || a.Contact.TangentDimension != 7 || a.Contact.HorizontalDimension != 6 {
		t.Fatalf("bad contact: %s", FormatContact(a.Contact))
	}
	if a.Reeb.AlphaOfReeb != 1 || a.Reeb.IReebDAlphaMaxOnTangent > 1e-12 || !a.Reeb.UniqueByContactEquation {
		t.Fatalf("bad Reeb: %s", FormatReeb(a.Reeb))
	}
	if a.Split.TangentDimension != 7 || a.Split.ReebLineDimension != 1 || a.Split.ContactDistributionDim != 6 || a.Split.SumDimension != 7 {
		t.Fatalf("bad split: %s", FormatSplit(a.Split))
	}
	if !a.Quotient.ProjectiveLawSpace || a.Quotient.SpacetimeIdentified || a.Quotient.PhysicalPhaseSpace || a.Quotient.Base != "CP^3" {
		t.Fatalf("bad quotient: %s", FormatQuotient(a.Quotient))
	}
	if !a.Phase.CentralU1Action || !a.Phase.GeneratedByTotalNumber || a.Phase.PhysicalHamiltonianTime {
		t.Fatalf("bad phase: %s", FormatPhase(a.Phase))
	}
	if !a.BL.CommutesWithTotalPhase || !a.BL.DescendsToCP3 || a.BL.SelectsWeakPlane || a.BL.SelectsGeneration {
		t.Fatalf("bad B-L relation: %s", FormatBL(a.BL))
	}
	if a.K7.HopfS7ToK7FunctorFound || a.K7.TangentS7ToK7FunctorFound || a.K7.DimensionMatchPromoted {
		t.Fatalf("bad K7 relation: %s", FormatK7(a.K7))
	}
	if a.Time.ReebToLorentzianTime || a.Time.ReebToOSPositivity || a.Time.ReebToHilbertDynamics || a.Time.ReebToRGScale || a.Time.ReebToObservedHistory || !a.Time.EWBridgeStillBridgeLevel {
		t.Fatalf("bad time firewall: %s", FormatTime(a.Time))
	}
}

func TestGate570Theorem(t *testing.T) {
	res := Generation2WittFockHopfS7ContactFormReebPhaseAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
