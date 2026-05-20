package generation2finitecontactdifferentialsourceaudit

import "fmt"

func FormatK7(a K7CarrierAudit) string {
	return fmt.Sprintf("dim=%d expected=%d ambient=%d index=%.10f frame=%.3e idem=%.3e sym=%.3e boolean=%.3e g2=%.3e verdict=%q", a.Dimension, a.ExpectedDimension, a.AmbientDimension, a.ContactIndex, a.FrameIsometryResidual, a.ProjectorIdempotenceResidual, a.ProjectorSymmetryResidual, a.BooleanContainmentResidual, a.G2ContainmentResidual, a.Verdict)
}
func FormatBoolean(a BooleanDifferentialSourceAudit) string {
	return fmt.Sprintf("grades=%d->%d dims=%d->%d incidence=%dx%d normalized=%dx%d rank=%d iso=%.3e unsigned=%t ambientMiddle=%t K7toK7=%t K7covTo2=%t d2=%t Leibniz=%t definesD=%t verdict=%q", a.LowerGrade, a.UpperGrade, a.LowerDimension, a.UpperDimension, a.IncidenceRows, a.IncidenceCols, a.NormalizedIncidenceRows, a.NormalizedIncidenceCols, a.RankFromGram, a.IsometryResidual, a.UnsignedIncidence, a.MapsIntoAmbientMiddleChamber, a.MapsFromK7ToK7, a.MapsK7CovectorsToK7TwoForms, a.HasD2ZeroCertificate, a.HasLeibnizCertificate, a.DefinesContactDifferential, a.Verdict)
}
func FormatG2(a G2DifferentialSourceAudit) string {
	return fmt.Sprintf("available=%t sector=%d projector=%t calibration=%t differential=%t definesD=%t verdict=%q", a.CalibrationSupportAvailable, a.SectorDimension, a.ProjectorAvailable, a.ProvidesCalibrationForm, a.ProvidesDifferential, a.DefinesDOnK7, a.Verdict)
}
func FormatProjector(a ProjectorDifferentialSourceAudit) string {
	return fmt.Sprintf("PB|K=I:%t PG|K=I:%t PKidempotent=%t commTrivial=%t adjacency=%t relativeDefinesD=%t verdict=%q", a.PBRestrictionToK7Identity, a.PGRestrictionToK7Identity, a.PKIdempotent, a.ProjectorCommutatorTrivial, a.AdjacencyOrBoundaryAvailable, a.RelativePositionDefinesDOnK7, a.Verdict)
}
func FormatSpectral(a SpectralDifferentialSourceAudit) string {
	return fmt.Sprintf("q4=%t endomorphism=%t return=%t differential=%t definesD=%t verdict=%q", a.Q4ContactSpectralData, a.CertifiedContactEndomorphism, a.CertifiedReturnMap, a.CertifiedDifferential, a.DefinesDOnK7, a.Verdict)
}
func FormatExterior(a ExteriorDifferentialAudit) string {
	return fmt.Sprintf("formalExterior=%t wedgeK7=%t dK7=%t cochain=%t d2=%t Leibniz=%t dAlpha=%t verdict=%q", a.FormalExteriorLanguageAvailable, a.WedgeProductOnK7Certified, a.FiniteExteriorDerivativeOnK7, a.CochainBoundaryOnK7, a.D2ZeroCertificate, a.LeibnizRuleCertificate, a.DAlphaComputable, a.Verdict)
}
func FormatContact(a ContactPackageConsequenceAudit) string {
	return fmt.Sprintf("alpha=%t d=%t dAlpha=%t volume=%t contact=%t reeb=%t split=%t verdict=%q", a.AlphaAvailable, a.DOperatorAvailable, a.DAlphaComputable, a.ContactVolumeKnown, a.ContactFormCertified, a.ReebVectorCertified, a.K7Splits1Plus6, a.Verdict)
}
func FormatTime(a ProductTimeFirewallAudit) string {
	return fmt.Sprintf("DM=%t lorentz=%t OS=%t Wick=%t Hilbert=%t Hamiltonian=%t RG=%t arrow=%t EWBridge=%t verdict=%q", a.ContactDToDM, a.ContactDToLorentzianTime, a.ContactDToOSPositivity, a.ContactDToWickRotation, a.ContactDToHilbertReconstruction, a.ContactDToHamiltonianSpectrum, a.ContactDToRGScale, a.ContactDToArrowOfTime, a.ElectroweakBridgeStillSealed, a.Verdict)
}
func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("finiteD=%t booleanPromoted=%t g2Promoted=%t projectorPromoted=%t q4Promoted=%t dAlpha=%t volume=%t reeb=%t physicalTime=%t next=%q verdict=%q", a.FiniteDOperatorFound, a.BooleanIncidencePromoted, a.G2CalibrationPromoted, a.ProjectorRelativePromoted, a.Q4Promoted, a.DAlphaCertified, a.ContactVolumeCertified, a.ReebCertified, a.PhysicalTimeRGOSHilbertOpen, a.MissingNextTheorem, a.Verdict)
}
