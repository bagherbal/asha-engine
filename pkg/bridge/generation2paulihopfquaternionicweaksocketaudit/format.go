package generation2paulihopfquaternionicweaksocketaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedAudit) string {
	return fmt.Sprintf("g560Pauli=%t g560Hopf=%t g560Split=%t g561NoSpatial=%t g561NoPlane=%t verdict=%q", a.Gate560PauliTriplet, a.Gate560HopfIdentity, a.Gate560ScalarMomentSplit, a.Gate561NoSpatialIntertwiner, a.Gate561NoCanonicalWeakPlane, a.Verdict)
}

func FormatQuaternionic(a QuaternionicSocketAudit) string {
	return fmt.Sprintf("A=%q H=%t basis=%v dim=%d metric=%t orient=%t bracket=%t unitary=%q ImHWeak=%t dynamics=%t verdict=%q", a.FiniteAlgebra, a.ContainsQuaternionicSummand, a.ImaginaryQuaternionicBasis, a.Dimension, a.MetricNormAvailable, a.OrientationAvailable, a.LieBracketCrossProductAvailable, a.UnitaryGroup, a.ImHAsWeakLieAlgebraStructural, a.PhysicalGaugeDynamicsDerived, a.Verdict)
}

func FormatScalarDoublet(a ScalarDoubletAudit) string {
	return fmt.Sprintf("carrier=%q real=%d complex=%d weak=%q single=%t Hmodule=%t structural=%t dynamical=%t yukawaFree=%t verdict=%q", a.Carrier, a.RealDimension, a.ComplexDimension, a.WeakRepresentation, a.SingleComplexDoubletRecovered, a.LeftHModuleOrEquivalentSU2Doublet, a.RepresentationNativeStructural, a.RepresentationDynamical, a.NumericalYukawaFree, a.Verdict)
}

func FormatRepresentation(a PauliQuaternionRepresentationAudit) string {
	return fmt.Sprintf("rhoH=%t unit=%t antiHermitian=%t hermitianMoment=%t gate560Pauli=%t moduleBasisIndependent=%t axisCanonical=%t convention=%q verdict=%q", a.RhoHAvailable, a.RhoHUnitPreserving, a.ImaginaryUnitsAntiHermitian, a.PauliMatricesHermitianMomentGenerators, a.CliffordPauliFromGate560, a.BasisIndependentAsModule, a.AxisByAxisIdentificationCanonical, a.ConventionFreedom, a.Verdict)
}

func FormatIntertwiner(a IntertwinerAudit) string {
	return fmt.Sprintf("source=%q target=%q exists=%t metric=%t bracket=%t unframedBasisIndependent=%t frameConventional=%t manualSigma3=%t verdict=%q", a.Source, a.Target, a.ModuleIntertwinerExists, a.MetricCompatible, a.LieBracketCompatible, a.BasisIndependentAsUnframedSpaces, a.SpecificSigmaToIJKFrameConventional, a.ManualSigma3ToK, a.Verdict)
}

func FormatMoment(a MomentMapAudit) string {
	return fmt.Sprintf("mu=%q momentMap=%t hopf=%t codomain=%q decomposition=%q normalization=%q gaugeBosons=%t verdict=%q", a.MuFormula, a.MomentMapForSU2Action, a.HopfIdentityInherited, a.Codomain, a.Decomposition, a.NormalizationConvention, a.IdentifiesPhysicalGaugeBosons, a.Verdict)
}

func FormatOrbit(a StabilizerOrbitAudit) string {
	return fmt.Sprintf("nonzero=%t split=%q radial=%t plane=%t scalarQuaternionic=%t WZphoton=%t verdict=%q", a.NonzeroMuCondition, a.Split, a.RadialLineCanonicalGivenMu, a.OrthogonalPlaneCanonicalGivenMetric, a.ScalarQuaternionicOnly, a.IdentifiesWZPhoton, a.Verdict)
}

func FormatEta(a EtaRelationAudit) string {
	return fmt.Sprintf("etaSigma3=%t chosenQuaternionAxis=%t scalarFrame=%t physicalAxis=%t tauShadow=%t verdict=%q", a.EtaEqualsSigma3, a.Sigma3CorrespondsToChosenQuaternionicAxis, a.AxisChosenByScalarFrame, a.AxisPhysicallyCanonical, a.TauEtaSigma3Shadow, a.Verdict)
}

func FormatSpectral(a SpectralTripleCompatibilityAudit) string {
	return fmt.Sprintf("AF=%t gamma=%t J=%t D=%t firstOrder=%t oneForm=%t heat=%t potential=%t dynamics=%t missing=[%s] verdict=%q", a.AFRepresentationStructural, a.GradingCompatibilityInherited, a.JCompatibilityInherited, a.DCompatibilityInherited, a.FirstOrderConditionInherited, a.FiniteOneFormScalarLaneStructural, a.HeatKernelProjectionAvailable, a.HiggsPotentialDerived, a.MassOrDynamicsDerived, strings.Join(a.MissingOrFirewalled, "; "), a.Verdict)
}

func FormatFirewall(a FirewallAudit) string {
	return fmt.Sprintf("preserved=%t weakBosons=%t photon=%t higgsMass=%t generation=%t yukawa=%t ckm=%t observed=%t Wplane=%t verdict=%q", a.Preserved, a.PhysicalWeakBosonsIdentified, a.PhotonIdentified, a.HiggsMassTheorem, a.GenerationHierarchyIdentified, a.YukawaTextureDerived, a.CKMPMNSDerived, a.ObservedFlavorImported, a.WSpatialWeakPlaneSelected, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("ImH3=%t Hmodule=%t PauliImH=%t moment=%t split=%t oneForm=%t dynamics=%t WgenTransfer=%t next=%q verdict=%q", a.ImHNativeOrientedMetricThreeSpace, a.HPhiWeakDoubletModule, a.PauliTripletEquivalentToImH, a.HopfMomentQuaternionicMomentMap, a.NonzeroMuQuaternionicThreeSplit, a.LinkedToFiniteOneFormStructurally, a.PhysicalElectroweakDynamicsDerived, a.LawfulTransferToWSpatialOrGeneration, a.MissingNextTheorem, a.Verdict)
}
