package generation2contactformcovectorobstructionaudit

import "testing"

func TestGate567ContactFormCovectorObstructionAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.K7.Dimension != 7 || !a.K7.InducedMetricIsIdentity || a.K7.BooleanContainmentResidual > 1e-8 || a.K7.G2ContainmentResidual > 1e-8 {
		t.Fatalf("bad K7 certificate: %s", FormatK7(a.K7))
	}
	if !a.Search.PBRestrictionToK7Identity || !a.Search.PGRestrictionToK7Identity || !a.Search.ProjectorCommutatorOnK7Trivial || a.Search.NativeDistinguishedObjectFound {
		t.Fatalf("distinguished search failed: %s", FormatSearch(a.Search))
	}
	if !a.G2.G2StructureAvailable || !a.G2.ActsTransitivelyOnUnitDirections || a.G2.ExtraSymmetryBreakingDatumPresent || a.G2.CanSelectReebDirection {
		t.Fatalf("G2 obstruction failed: %s", FormatG2(a.G2))
	}
	if a.Alpha.CandidateCovectorFound || a.Alpha.CandidateVectorFound || a.Alpha.AlphaConstructed {
		t.Fatalf("alpha firewall failed: %s", FormatAlpha(a.Alpha))
	}
	if !a.DAlpha.ExteriorAlgebraAvailable || a.DAlpha.FiniteDOperatorOnK7Available || a.DAlpha.DAlphaComputable {
		t.Fatalf("d alpha firewall failed: %s", FormatDAlpha(a.DAlpha))
	}
	if a.Contact.AlphaAvailable || a.Contact.DAlphaAvailable || a.Contact.AlphaWedgeDAlphaCubedKnown || a.Contact.ContactFormCertified {
		t.Fatalf("contact condition firewall failed: %s", FormatContact(a.Contact))
	}
	if a.Reeb.UniqueReeb || a.Reeb.SplitK7As1Plus6 {
		t.Fatalf("Reeb firewall failed: %s", FormatReeb(a.Reeb))
	}
	if !a.Q4.ContactSpectralData || a.Q4.CertifiedContactEndomorphism || a.Q4.CertifiedReebReturnMap || a.Q4.CertifiedLinearizedReebFlow || !a.Q4.HiggsFlavorYukawaPromotionBlocked {
		t.Fatalf("q4 firewall failed: %s", FormatQ4(a.Q4))
	}
	if !a.E0.CliffordE0AvailableAsSignatureDatum || a.E0.E0ProjectionIntoK7Available || a.E0.E0ToReebFunctorAvailable || a.E0.ReebAvailable || !a.E0.SeparationPreserved {
		t.Fatalf("e0 firewall failed: %s", FormatE0(a.E0))
	}
	if a.Time.ContactToDM || a.Time.ContactToLorentzianTime || a.Time.ContactToOSPositivity || a.Time.ContactToWickRotation || a.Time.ContactToHilbertReconstruction || a.Time.ContactToHamiltonianSpectrum || a.Time.ContactToRGScale || a.Time.ContactToArrowOfTime || !a.Time.ElectroweakBridgeStillSealed {
		t.Fatalf("time firewall failed: %s", FormatTime(a.Time))
	}
}

func TestGate567Theorem(t *testing.T) {
	res := Generation2ContactFormCertificateAndDistinguishedCovectorObstructionAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
