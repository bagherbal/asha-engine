package generation2hopfs7k7producttimeairlockaudit

import "testing"

func TestGate571HopfS7K7ProductTimeObstructionAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Hopf.Gate570ContactCertified || !a.Hopf.Gate570ReebCertified || !a.Hopf.Gate570SplitCertified || a.Hopf.PhysicalTimeOpened {
		t.Fatalf("bad inherited Hopf package: %s", FormatHopf(a.Hopf))
	}
	if !a.K7.K7CarrierCertified || a.K7.K7Dimension != 7 || a.K7.HopfS7ToK7Already || a.K7.TangentS7ToK7 {
		t.Fatalf("bad K7 inheritance: %s", FormatK7(a.K7))
	}
	if !a.Types.SameRealDimension || a.Types.DimensionMatchPromoted || !a.Types.NonlinearToLinearIssue || a.Types.BasisIndependentFunctor {
		t.Fatalf("bad type comparison: %s", FormatTypes(a.Types))
	}
	if a.Contact.FunctorFound || a.Contact.AlphaPullbackCertified || a.Contact.ReebImageCertified || a.Contact.HorizontalPlaneCertified {
		t.Fatalf("unexpected contact functor: %s", FormatContact(a.Contact))
	}
	if a.Quotient.CP3ToK7FunctorFound || a.Quotient.K7CentralU1ActionFound || a.Quotient.BMinusLCanonicalizesK7 || a.Quotient.WeakPlaneOrGeneration {
		t.Fatalf("unexpected quotient/phase transfer: %s", FormatQuotient(a.Quotient))
	}
	if a.Time.FockPhaseToLorentzianTime || a.Time.FockPhaseToOSPositivity || a.Time.FockPhaseToHilbert || a.Time.FockPhaseToRGScale || a.Time.FockPhaseToHamiltonian || !a.Time.ElectroweakBridgeOnly {
		t.Fatalf("bad time firewall: %s", FormatTime(a.Time))
	}
	if a.Final.HopfToK7FunctorFound || a.Final.TangentToK7FunctorFound || a.Final.ProductTimeAirlockOpened || a.Final.RGOSHilbertOpened || a.Final.PhysicalDynamicsOpened {
		t.Fatalf("bad final verdict: %s", FormatFinal(a.Final))
	}
}

func TestGate571Theorem(t *testing.T) {
	res := Generation2HopfS7K7ProductTimeAirlockObstructionAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
