package generation2hopfs7k7producttimeairlockaudit

import "fmt"

func FormatHopf(a InheritedHopfAudit) string {
	return fmt.Sprintf("contact=%t reeb=%t split=%t sphere=%q dim=%d CP3=%t totalPhase=%t physicalTime=%t verdict=%q", a.Gate570ContactCertified, a.Gate570ReebCertified, a.Gate570SplitCertified, a.SphereName, a.SphereDimension, a.CP3ProjectiveLawSpace, a.ReebIsTotalFockPhase, a.PhysicalTimeOpened, a.Verdict)
}
func FormatK7(a InheritedK7Audit) string {
	return fmt.Sprintf("K7=%t dim=%d nature=%q S7toK7Already=%t TS7toK7=%t verdict=%q", a.K7CarrierCertified, a.K7Dimension, a.K7Nature, a.HopfS7ToK7Already, a.TangentS7ToK7, a.Verdict)
}
func FormatTypes(a TypeComparisonAudit) string {
	return fmt.Sprintf("Hopf=%q K7=%q sameDim=%t dimPromoted=%t nonlinearLinear=%t basepoint=%t metricContactMismatch=%t functor=%t verdict=%q", a.HopfObjectType, a.K7ObjectType, a.SameRealDimension, a.DimensionMatchPromoted, a.NonlinearToLinearIssue, a.BasepointRequired, a.MetricContactMismatch, a.BasisIndependentFunctor, a.Verdict)
}
func FormatContact(a ContactIntertwinerAudit) string {
	return fmt.Sprintf("candidate=%q basepoint=%t metric=%t alpha=%t reeb=%t horizontal=%t alphaOK=%t reebOK=%t horizontalOK=%t functor=%t verdict=%q", a.CandidateFunctorName, a.RequiresBasepoint, a.RequiresMetricPreservation, a.RequiresAlphaPullback, a.RequiresReebImage, a.RequiresHorizontalImage, a.AlphaPullbackCertified, a.ReebImageCertified, a.HorizontalPlaneCertified, a.FunctorFound, a.Verdict)
}
func FormatQuotient(a QuotientPhaseAudit) string {
	return fmt.Sprintf("quotient=%q CP3dim=%d K7quotient=%t CP3toK7=%t phase=%q K7U1=%t BLdescends=%t BLcanonicalizesK7=%t weakOrGen=%t verdict=%q", a.HopfQuotient, a.CP3Dimension, a.K7QuotientAvailable, a.CP3ToK7FunctorFound, a.TotalPhaseAction, a.K7CentralU1ActionFound, a.BMinusLDescendsToCP3, a.BMinusLCanonicalizesK7, a.WeakPlaneOrGeneration, a.Verdict)
}
func FormatTime(a ProductTimeAirlockAudit) string {
	return fmt.Sprintf("DM=%t Lorentz=%t OS=%t Wick=%t Hilbert=%t Hamiltonian=%t unitary=%t RG=%t cosmological=%t observed=%t EWBridge=%t verdict=%q", a.FockPhaseToDM, a.FockPhaseToLorentzianTime, a.FockPhaseToOSPositivity, a.FockPhaseToWickRotation, a.FockPhaseToHilbert, a.FockPhaseToHamiltonian, a.FockPhaseToUnitaryFlow, a.FockPhaseToRGScale, a.FockPhaseToCosmological, a.FockPhaseToObserved, a.ElectroweakBridgeOnly, a.Verdict)
}
func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("Hopf=%t K7=%t dimOnly=%t HopfToK7=%t TS7toK7=%t productTime=%t RGOSHilbert=%t physicalDynamics=%t next=%q verdict=%q", a.HopfContactInherited, a.K7CarrierInherited, a.DimensionMatchOnly, a.HopfToK7FunctorFound, a.TangentToK7FunctorFound, a.ProductTimeAirlockOpened, a.RGOSHilbertOpened, a.PhysicalDynamicsOpened, a.MissingNextTheorem, a.Verdict)
}
