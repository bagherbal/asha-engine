package generation2finitecontactcochaincomplexaudit

import "testing"

func TestGate569FiniteContactCochainComplexAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.K7.Dimension != 7 || a.K7.FrameIsometryResidual > 1e-8 || a.K7.BooleanContainmentResidual > 1e-8 || a.K7.G2ContainmentResidual > 1e-8 {
		t.Fatalf("bad K7: %s", FormatK7(a.K7))
	}
	want := []int{1, 7, 21, 35, 35, 21, 7, 1}
	if a.Exterior.VectorDimension != 7 || a.Exterior.TotalDimension != 128 || len(a.Exterior.GradeDimensions) != len(want) || !a.Exterior.HasAbstractExteriorDimensions || a.Exterior.HasCertifiedK7CochainBasis || a.Exterior.HasCertifiedWedgeProductOnK7 || a.Exterior.HasFiniteDOperator {
		t.Fatalf("bad exterior audit: %s", FormatExterior(a.Exterior))
	}
	for i := range want {
		if a.Exterior.GradeDimensions[i] != want[i] {
			t.Fatalf("bad R7 exterior dimensions: %s", FormatExterior(a.Exterior))
		}
	}
	if !a.Boolean.UnsignedIncidence || a.Boolean.Composition34After23Frobenius <= 0 || a.Boolean.D2ZeroForUnsignedIncidence || a.Boolean.SignedOrientationAvailable || a.Boolean.DefinesK7Differential {
		t.Fatalf("bad Boolean d2 obstruction: %s", FormatBoolean(a.Boolean))
	}
	if a.Restriction.K7CochainComplexDefined || a.Restriction.ProjectionFromAmbientFormsToK7Coforms || a.Restriction.PullbackDifferentialDefined || a.Restriction.D2ZeroOnRestrictedComplex {
		t.Fatalf("bad restriction: %s", FormatRestriction(a.Restriction))
	}
	if a.Sources.G2CalibrationSuppliesComplex || a.Sources.ProjectorRelativePositionSuppliesBoundary || a.Sources.Q4SuppliesComplex {
		t.Fatalf("bad source audit: %s", FormatSources(a.Sources))
	}
	if a.Law.HasFullComplex || a.Law.HasD2ZeroCertificate || a.Law.HasLeibnizCertificate || a.Law.HasAlphaCompatibleD || a.Law.DAlphaComputable || a.Law.ContactVolumeComputable || a.Law.ReebComputable {
		t.Fatalf("bad law audit: %s", FormatLaw(a.Law))
	}
	if a.Time.CochainToDM || a.Time.CochainToLorentzianTime || a.Time.CochainToOSPositivity || a.Time.CochainToWickRotation || a.Time.CochainToHilbert || a.Time.CochainToHamiltonian || a.Time.CochainToRGScale || a.Time.CochainToArrowOfTime || !a.Time.ElectroweakBridgeStillSealed {
		t.Fatalf("bad time firewall: %s", FormatTime(a.Time))
	}
}

func TestGate569Theorem(t *testing.T) {
	res := Generation2FiniteContactCochainComplexD2ZeroCertificateAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
