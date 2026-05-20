package generation2finitecontactcochaincomplexaudit

import "fmt"

func FormatK7(a K7CarrierAudit) string {
	return fmt.Sprintf("dim=%d expected=%d ambient=%d index=%.10f frame=%.3e boolean=%.3e g2=%.3e verdict=%q", a.Dimension, a.ExpectedDimension, a.AmbientDimension, a.ContactIndex, a.FrameIsometryResidual, a.BooleanContainmentResidual, a.G2ContainmentResidual, a.Verdict)
}
func FormatExterior(a FormalR7ExteriorAudit) string {
	return fmt.Sprintf("R%d grades=%v total=%d abstract=%t K7basis=%t wedge=%t d=%t verdict=%q", a.VectorDimension, a.GradeDimensions, a.TotalDimension, a.HasAbstractExteriorDimensions, a.HasCertifiedK7CochainBasis, a.HasCertifiedWedgeProductOnK7, a.HasFiniteDOperator, a.Verdict)
}
func FormatBoolean(a BooleanConsecutiveIncidenceAudit) string {
	return fmt.Sprintf("M23=%dx%d M34=%dx%d M34*M23=%dx%d norm=%.3e max=%.3e unsigned=%t d2zero=%t signed=%t definesK7d=%t verdict=%q", a.M23Rows, a.M23Cols, a.M34Rows, a.M34Cols, a.Composition34After23Rows, a.Composition34After23Cols, a.Composition34After23Frobenius, a.Composition34After23MaxAbs, a.UnsignedIncidence, a.D2ZeroForUnsignedIncidence, a.SignedOrientationAvailable, a.DefinesK7Differential, a.Verdict)
}

func FormatRestriction(a K7RestrictionAudit) string {
	return fmt.Sprintf("domain=%q codomain=%q K7=%q complex=%t projection=%t pullbackD=%t d2=%t verdict=%q", a.BooleanIncidenceDomain, a.BooleanIncidenceCodomain, a.K7LivesIn, a.K7CochainComplexDefined, a.ProjectionFromAmbientFormsToK7Coforms, a.PullbackDifferentialDefined, a.D2ZeroOnRestrictedComplex, a.Verdict)
}
func FormatSources(a SourceAudit) string {
	return fmt.Sprintf("G2complex=%t projectorBoundary=%t q4complex=%t verdict=%q", a.G2CalibrationSuppliesComplex, a.ProjectorRelativePositionSuppliesBoundary, a.Q4SuppliesComplex, a.Verdict)
}
func FormatLaw(a DifferentialLawAudit) string {
	return fmt.Sprintf("d0=%t d1=%t d2=%t full=%t d2zero=%t Leibniz=%t alphaD=%t dAlpha=%t volume=%t reeb=%t verdict=%q", a.HasD0ToD1, a.HasD1ToD2, a.HasD2ToD3, a.HasFullComplex, a.HasD2ZeroCertificate, a.HasLeibnizCertificate, a.HasAlphaCompatibleD, a.DAlphaComputable, a.ContactVolumeComputable, a.ReebComputable, a.Verdict)
}
func FormatTime(a ProductTimeFirewallAudit) string {
	return fmt.Sprintf("DM=%t Lorentz=%t OS=%t Wick=%t Hilbert=%t Hamiltonian=%t RG=%t arrow=%t EWBridge=%t verdict=%q", a.CochainToDM, a.CochainToLorentzianTime, a.CochainToOSPositivity, a.CochainToWickRotation, a.CochainToHilbert, a.CochainToHamiltonian, a.CochainToRGScale, a.CochainToArrowOfTime, a.ElectroweakBridgeStillSealed, a.Verdict)
}
func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("K7=%t exteriorOnly=%t unsignedD2Fails=%t signedD=%t K7complex=%t d2=%t Leibniz=%t dAlpha=%t volume=%t reeb=%t physicalTime=%t next=%q verdict=%q", a.K7Certified, a.FormalExteriorDimensionsOnly, a.UnsignedBooleanIncidenceFailsD2, a.SignedFiniteDifferentialFound, a.K7CochainComplexFound, a.D2ZeroCertified, a.LeibnizCertified, a.DAlphaCertified, a.ContactVolumeCertified, a.ReebCertified, a.PhysicalTimeOrRGOpened, a.MissingNextTheorem, a.Verdict)
}
