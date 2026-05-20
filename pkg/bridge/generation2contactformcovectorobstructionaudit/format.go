package generation2contactformcovectorobstructionaudit

import "fmt"

func FormatK7(a K7BasisMetricAudit) string {
	return fmt.Sprintf("dim=%d expected=%d ambient=%d frame=%dx%d index=%.10f frameResidual=%.3e projectorIdem=%.3e projectorSym=%.3e boolean=%.3e g2=%.3e metricI=%t verdict=%q", a.Dimension, a.ExpectedDimension, a.AmbientDimension, a.FrameRows, a.FrameColumns, a.ContactIndex, a.FrameIsometryResidual, a.ProjectorIdempotenceResidual, a.ProjectorSymmetryResidual, a.BooleanContainmentResidual, a.G2ContainmentResidual, a.InducedMetricIsIdentity, a.Verdict)
}
func FormatSearch(a DistinguishedSearchAudit) string {
	return fmt.Sprintf("PB=%t PG=%t comm=%t relpos=%t booleanTensor=%t g2Calib=%t q4=%t traceRank=%t e0=%t PB|K=I:%t PG|K=I:%t commK=%t found=%t verdict=%q", a.FromPB, a.FromPG, a.FromCommutator, a.FromRelativePosition, a.FromBooleanIncidenceTensor, a.FromG2Calibration, a.FromQ4ContactSpectralBlock, a.FromTraceRankAsymmetry, a.FromCliffordE0Projection, a.PBRestrictionToK7Identity, a.PGRestrictionToK7Identity, a.ProjectorCommutatorOnK7Trivial, a.NativeDistinguishedObjectFound, a.Verdict)
}
func FormatG2(a G2ObstructionAudit) string {
	return fmt.Sprintf("g2=%t transitive=%t extraDatum=%t selects=%t verdict=%q", a.G2StructureAvailable, a.ActsTransitivelyOnUnitDirections, a.ExtraSymmetryBreakingDatumPresent, a.CanSelectReebDirection, a.Verdict)
}
func FormatAlpha(a CandidateAlphaAudit) string {
	return fmt.Sprintf("covector=%t vector=%t nativeBasisIndependent=%t alpha=%t signScale=%t canonical=%t verdict=%q", a.CandidateCovectorFound, a.CandidateVectorFound, a.NativeBasisIndependent, a.AlphaConstructed, a.UniqueUpToSignOrScale, a.FullyCanonical, a.Verdict)
}
func FormatDAlpha(a FiniteDifferentialAudit) string {
	return fmt.Sprintf("exterior=%t finiteD=%t cochain=%t incidence=%t dAlpha=%t verdict=%q", a.ExteriorAlgebraAvailable, a.FiniteDOperatorOnK7Available, a.CochainBoundaryAvailable, a.IncidenceDifferentialOnK7, a.DAlphaComputable, a.Verdict)
}
func FormatContact(a ContactConditionAudit) string {
	return fmt.Sprintf("alpha=%t dAlpha=%t wedgeKnown=%t wedgeNonzero=%t certified=%t verdict=%q", a.AlphaAvailable, a.DAlphaAvailable, a.AlphaWedgeDAlphaCubedKnown, a.AlphaWedgeDAlphaCubedNonzero, a.ContactFormCertified, a.Verdict)
}
func FormatReeb(a ReebAudit) string {
	return fmt.Sprintf("alpha=%t dAlpha=%t alphaR=%t contraction=%t unique=%t split=%t verdict=%q", a.AlphaAvailable, a.DAlphaAvailable, a.SolvedAlphaOfR, a.SolvedContraction, a.UniqueReeb, a.SplitK7As1Plus6, a.Verdict)
}
func FormatQ4(a Q4RelationAudit) string {
	return fmt.Sprintf("q4=%q contact=%t endomorphism=%t reebReturn=%t linearized=%t firewall=%t verdict=%q", a.Polynomial, a.ContactSpectralData, a.CertifiedContactEndomorphism, a.CertifiedReebReturnMap, a.CertifiedLinearizedReebFlow, a.HiggsFlavorYukawaPromotionBlocked, a.Verdict)
}
func FormatE0(a E0RelationAudit) string {
	return fmt.Sprintf("e0Signature=%t projection=%t functor=%t reeb=%t separated=%t verdict=%q", a.CliffordE0AvailableAsSignatureDatum, a.E0ProjectionIntoK7Available, a.E0ToReebFunctorAvailable, a.ReebAvailable, a.SeparationPreserved, a.Verdict)
}
func FormatTime(a ProductTimeFirewallAudit) string {
	return fmt.Sprintf("DM=%t lorentz=%t OS=%t Wick=%t Hilbert=%t Hamiltonian=%t RG=%t arrow=%t EWBridge=%t verdict=%q", a.ContactToDM, a.ContactToLorentzianTime, a.ContactToOSPositivity, a.ContactToWickRotation, a.ContactToHilbertReconstruction, a.ContactToHamiltonianSpectrum, a.ContactToRGScale, a.ContactToArrowOfTime, a.ElectroweakBridgeStillSealed, a.Verdict)
}
func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("distinguished=%t finiteD=%t alpha=%t dAlpha=%t volume=%t reeb=%t split=%t q4Reeb=%t e0Reeb=%t physicalTime=%t next=%q verdict=%q", a.NativeDistinguishedVectorOrCovector, a.FiniteDOperatorOnK7, a.AlphaCertified, a.DAlphaCertified, a.ContactVolumeNonzero, a.ReebCertified, a.K7Splits1Plus6, a.Q4RelatedToReebDynamics, a.E0RelatedToReeb, a.PhysicalTimeRGOSHilbertOpened, a.MissingNextTheorem, a.Verdict)
}
