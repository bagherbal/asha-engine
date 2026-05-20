package generation2contactformcovectorobstructionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2ContactFormCertificateAndDistinguishedCovectorObstructionAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 contact form certificate and distinguished covector obstruction audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate567 contact form/covector audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "recover certified K_7 basis and metric", Passed: a.K7.Dimension == 7 && a.K7.InducedMetricIsIdentity && a.K7.BooleanContainmentResidual < 1e-8 && a.K7.G2ContainmentResidual < 1e-8, Detail: FormatK7(a.K7)},
			{Name: "reject distinguished vector/covector search", Passed: a.Search.PBRestrictionToK7Identity && a.Search.PGRestrictionToK7Identity && a.Search.ProjectorCommutatorOnK7Trivial && !a.Search.NativeDistinguishedObjectFound, Detail: FormatSearch(a.Search)},
			{Name: "block G2-only Reeb direction selection", Passed: a.G2.G2StructureAvailable && a.G2.ActsTransitivelyOnUnitDirections && !a.G2.ExtraSymmetryBreakingDatumPresent && !a.G2.CanSelectReebDirection, Detail: FormatG2(a.G2)},
			{Name: "block contact alpha construction", Passed: !a.Alpha.CandidateCovectorFound && !a.Alpha.CandidateVectorFound && !a.Alpha.AlphaConstructed && !a.Alpha.FullyCanonical, Detail: FormatAlpha(a.Alpha)},
			{Name: "block finite d alpha construction", Passed: a.DAlpha.ExteriorAlgebraAvailable && !a.DAlpha.FiniteDOperatorOnK7Available && !a.DAlpha.DAlphaComputable, Detail: FormatDAlpha(a.DAlpha)},
			{Name: "block contact condition alpha wedge dalpha cubed", Passed: !a.Contact.AlphaAvailable && !a.Contact.DAlphaAvailable && !a.Contact.AlphaWedgeDAlphaCubedKnown && !a.Contact.ContactFormCertified, Detail: FormatContact(a.Contact)},
			{Name: "block Reeb vector and 7=1+6 split", Passed: !a.Reeb.SolvedAlphaOfR && !a.Reeb.SolvedContraction && !a.Reeb.UniqueReeb && !a.Reeb.SplitK7As1Plus6, Detail: FormatReeb(a.Reeb)},
			{Name: "keep q4 independent from Reeb dynamics", Passed: a.Q4.ContactSpectralData && !a.Q4.CertifiedContactEndomorphism && !a.Q4.CertifiedReebReturnMap && !a.Q4.CertifiedLinearizedReebFlow && a.Q4.HiggsFlavorYukawaPromotionBlocked, Detail: FormatQ4(a.Q4)},
			{Name: "separate Clifford e0 from Reeb and physical time", Passed: a.E0.CliffordE0AvailableAsSignatureDatum && !a.E0.E0ProjectionIntoK7Available && !a.E0.E0ToReebFunctorAvailable && !a.E0.ReebAvailable && a.E0.SeparationPreserved, Detail: FormatE0(a.E0)},
			{Name: "preserve product-time/RG/OS/Hilbert firewall", Passed: !a.Time.ContactToDM && !a.Time.ContactToLorentzianTime && !a.Time.ContactToOSPositivity && !a.Time.ContactToWickRotation && !a.Time.ContactToHilbertReconstruction && !a.Time.ContactToHamiltonianSpectrum && !a.Time.ContactToRGScale && !a.Time.ContactToArrowOfTime && a.Time.ElectroweakBridgeStillSealed, Detail: FormatTime(a.Time)},
			{Name: "return final contact-form obstruction verdict", Passed: !a.Final.NativeDistinguishedVectorOrCovector && !a.Final.FiniteDOperatorOnK7 && !a.Final.AlphaCertified && !a.Final.DAlphaCertified && !a.Final.ContactVolumeNonzero && !a.Final.ReebCertified && !a.Final.K7Splits1Plus6 && !a.Final.Q4RelatedToReebDynamics && !a.Final.E0RelatedToReeb && !a.Final.PhysicalTimeRGOSHilbertOpened, Detail: FormatFinal(a.Final)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(Statuses(), a.Truth, a.Final.MissingNextTheorem)}
	}}
}
