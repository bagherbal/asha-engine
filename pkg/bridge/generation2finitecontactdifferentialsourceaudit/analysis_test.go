package generation2finitecontactdifferentialsourceaudit

import "testing"

func TestGate568FiniteContactDifferentialSourceSearchAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.K7.Dimension != 7 || a.K7.FrameIsometryResidual > 1e-8 || a.K7.BooleanContainmentResidual > 1e-8 || a.K7.G2ContainmentResidual > 1e-8 {
		t.Fatalf("bad K7: %s", FormatK7(a.K7))
	}
	if a.Boolean.LowerGrade != 3 || a.Boolean.UpperGrade != 4 || a.Boolean.RankFromGram != 56 || !a.Boolean.UnsignedIncidence || a.Boolean.MapsFromK7ToK7 || a.Boolean.HasD2ZeroCertificate || a.Boolean.DefinesContactDifferential {
		t.Fatalf("bad Boolean obstruction: %s", FormatBoolean(a.Boolean))
	}
	if !a.G2.CalibrationSupportAvailable || !a.G2.ProjectorAvailable || a.G2.ProvidesDifferential || a.G2.DefinesDOnK7 {
		t.Fatalf("bad G2 obstruction: %s", FormatG2(a.G2))
	}
	if !a.Projector.PBRestrictionToK7Identity || !a.Projector.PGRestrictionToK7Identity || !a.Projector.PKIdempotent || !a.Projector.ProjectorCommutatorTrivial || a.Projector.AdjacencyOrBoundaryAvailable || a.Projector.RelativePositionDefinesDOnK7 {
		t.Fatalf("bad projector obstruction: %s", FormatProjector(a.Projector))
	}
	if !a.Spectral.Q4ContactSpectralData || a.Spectral.CertifiedDifferential || a.Spectral.DefinesDOnK7 {
		t.Fatalf("bad spectral obstruction: %s", FormatSpectral(a.Spectral))
	}
	if !a.Exterior.FormalExteriorLanguageAvailable || a.Exterior.FiniteExteriorDerivativeOnK7 || a.Exterior.CochainBoundaryOnK7 || a.Exterior.D2ZeroCertificate || a.Exterior.DAlphaComputable {
		t.Fatalf("bad exterior obstruction: %s", FormatExterior(a.Exterior))
	}
	if a.Contact.AlphaAvailable || a.Contact.DOperatorAvailable || a.Contact.DAlphaComputable || a.Contact.ContactVolumeKnown || a.Contact.ContactFormCertified || a.Contact.ReebVectorCertified || a.Contact.K7Splits1Plus6 {
		t.Fatalf("bad contact consequence: %s", FormatContact(a.Contact))
	}
	if a.Time.ContactDToDM || a.Time.ContactDToLorentzianTime || a.Time.ContactDToOSPositivity || a.Time.ContactDToWickRotation || a.Time.ContactDToHilbertReconstruction || a.Time.ContactDToHamiltonianSpectrum || a.Time.ContactDToRGScale || a.Time.ContactDToArrowOfTime || !a.Time.ElectroweakBridgeStillSealed {
		t.Fatalf("bad time firewall: %s", FormatTime(a.Time))
	}
}

func TestGate568Theorem(t *testing.T) {
	res := Generation2FiniteContactDifferentialSourceSearchAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
