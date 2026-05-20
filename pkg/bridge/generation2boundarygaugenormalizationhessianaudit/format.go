package generation2boundarygaugenormalizationhessianaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedHessianAudit) string {
	return fmt.Sprintf("hessian=%t neutralNull=%t noPhysical=%t noFlavor=%t ratio=%q verdict=%q", a.Gate564HessianShape, a.Gate564NeutralNull, a.Gate564NoPhysicalDynamics, a.Gate564NoFlavorData, a.Gate564RatioShape, a.Verdict)
}

func FormatGaugeNorm(a GaugeKineticNormalizationAudit) string {
	return fmt.Sprintf("layer=%q theorem=%q kY=%s formula=%q recovered=%t sin2=%s sinRecovered=%t lowEnergy=%t observed=%t verdict=%q", a.SourceLayer, a.SourceTheorem, a.KY.String(), a.KYFormula, a.KYRecovered, a.BoundarySin2FromSource.String(), a.BoundarySin2Recovered, a.LowEnergyObservedClaim, a.ObservedInputUsed, a.Verdict)
}

func FormatCouplings(a CouplingConventionAudit) string {
	return fmt.Sprintf("g1=%q gprime=%q relation=%q ratio=%s convention=%t nativePhysical=%t verdict=%q", a.CanonicalHyperchargeCoupling, a.Gate564AbelianCoupling, a.Relation, a.RatioUnderBoundaryEquality.String(), a.ConventionVerified, a.NativePhysicalCouplingValue, a.Verdict)
}

func FormatBoundaryEquality(a BoundaryEqualityAudit) string {
	return fmt.Sprintf("equality=%q native=%t bridge=%t absoluteUnit=%t running=%t verdict=%q", a.Equality, a.EqualityNativeTheorem, a.EqualityBridgeBoundary, a.AbsoluteCouplingUnitDerived, a.LowEnergyRunningDerived, a.Verdict)
}

func FormatWeakAngle(a WeakAngleBoundaryAudit) string {
	return fmt.Sprintf("kY=%s gp2g2=%s sin2=%s derivation=%q matches=%t observed=%t verdict=%q", a.KY.String(), a.GPrimeSquaredOverGSquared.String(), a.Sin2ThetaStar.String(), a.Derivation, a.MatchesPreviousASHA, a.ObservedWeakAngleImported, a.Verdict)
}

func FormatHessianRatio(a HessianRatioAlignmentAudit) string {
	return fmt.Sprintf("shape=%q inserted=%s mWmZ=%s derivation=%q physical=%t observed=%t verdict=%q", a.Gate564RatioShape, a.InsertedBoundaryRatio.String(), a.BoundaryMW2OverMZ2.String(), a.Derivation, a.PhysicalLowEnergyMassRatio, a.ObservedMassImported, a.Verdict)
}

func FormatRemaining(a RemainingVariablesFirewall) string {
	return fmt.Sprintf("vars=[%s] Kphi=%t v=%t g=%t gp=%t f0=%t a=%t metric=%t thresholds=%t verdict=%q", strings.Join(a.BridgeEnvironmental, ", "), a.NativeAbsoluteKphi, a.NativeV, a.NativeAbsoluteG, a.NativeAbsoluteGPrime, a.NativeF0, a.NativeYukawaTraceA, a.NativeScalarMetric, a.NativeRGThresholds, a.Verdict)
}

func FormatPhotonFlavor(a PhotonAndFlavorFirewall) string {
	return fmt.Sprintf("socketOnly=%t photon=%t osWickHilbert=%t yukawa=%t ckm=%t generation=%t observedFlavor=%t verdict=%q", a.ASocketSymbolicOnly, a.PhysicalPhotonDerived, a.OSWickHilbertDerived, a.YukawaEigenvalues, a.CKMPMNS, a.GenerationHierarchy, a.ObservedFlavorData, a.Verdict)
}

func FormatRelations(a RelationAudit) string {
	return fmt.Sprintf("q4=%t tauEta=%t Wspatial=%t pauliRoute=%t gate564=%t boundaryOnly=%t verdict=%q", a.Q4ContactOnly, a.TauEtaSigma3TraceShadow, a.WSpatialWeakPlaneBlocked, a.PauliQuaternionicScalarRoute, a.Gate564HessianShape, a.Gate565BoundaryAlignmentOnly, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("kYLayer=%t convention=%t equality=%q sin238=%t ratio58=%t vars=[%s] physical=%t flavor=%t next=%q verdict=%q", a.KYRecoveredCorrectLayer, a.CouplingConventionVerified, a.BoundaryEqualityLayer, a.Sin238Passes, a.HessianRatio58Passes, strings.Join(a.BridgeEnvironmentalVariables, ", "), a.PhysicalLowEnergyPrediction, a.FlavorOrObservedDataProduced, a.MissingNextTheorem, a.Verdict)
}
