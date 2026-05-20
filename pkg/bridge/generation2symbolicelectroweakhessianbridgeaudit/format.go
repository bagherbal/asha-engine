package generation2symbolicelectroweakhessianbridgeaudit

import (
	"fmt"
	"strings"
)

func FormatInherited(a InheritedAudit) string {
	return fmt.Sprintf("scalarDoublet=%t ImHAction=%t momentNotCurvature=%t noU1Photon=%t noKinetic=%t noFlavor=%t verdict=%q", a.Gate563ScalarDoubletLane, a.Gate563ImHActionOnHphi, a.Gate563MomentNotNativeCurvature, a.Gate563NoNativeU1PhotonDirection, a.Gate563NoNativeKineticNormalization, a.Gate563NoFlavorData, a.Verdict)
}

func FormatVacuum(a ScalarVacuumAudit) string {
	return fmt.Sprintf("carrier=%q phi0=%q norm=%q derived=%t sealed=%t generators=%q hypercharge=%q stabilizer=%q solution=%q solved=%t verdict=%q", a.Carrier, a.VacuumSymbol, a.NormSymbol, a.VacuumDerivedNatively, a.VacuumBridgeSealed, a.GeneratorsConvention, a.HyperchargeConvention, a.StabilizerEquation, a.StabilizerSolution, a.StabilizerSolvedSymbolically, a.Verdict)
}

func FormatCharged(a ChargedSectorAudit) string {
	return fmt.Sprintf("expr=%q realBasis=%q chargedBasis=%q coeffReal=%q coeffCharged=%q observed=%t numericalCoupling=%t verdict=%q", a.KineticExpression, a.RealBasis, a.ChargedBasis, a.PerRealGeneratorCoefficient, a.ChargedPairCoefficient, a.ObservedMassImported, a.NumericalCouplingImported, a.Verdict)
}

func FormatNeutral(a NeutralSectorAudit) string {
	return fmt.Sprintf("basis=%q factor=%q matrix=[[ %s, %s ],[ %s, %s ]] trace=%q det=%s rank=%d eig=(%q,%q) convention=%q verdict=%q", a.Basis, a.OverallFactor, a.Matrix[0][0], a.Matrix[0][1], a.Matrix[1][0], a.Matrix[1][1], a.Trace, a.Determinant, a.Rank, a.Eigenvalues[0], a.Eigenvalues[1], a.Convention, a.Verdict)
}

func FormatNull(a NullDirectionAudit) string {
	return fmt.Sprintf("det0=%t null=%q massive=%q socketOnly=%t physicalPhoton=%t requiresDynamics=%t verdict=%q", a.DeterminantZero, a.NullDirection, a.MassiveDirection, a.PhotonSocketOnly, a.PhysicalPhotonDerived, a.RequiresOSWickHilbertGaugeDynamics, a.Verdict)
}

func FormatMassRatio(a MassRatioAudit) string {
	return fmt.Sprintf("mW=%q mZ=%q ratio=%q Kphi=%t v=%t couplings=%t convention=%t observed=%t verdict=%q", a.MW2Shape, a.MZ2Shape, a.RatioShape, a.DependsOnKphi, a.DependsOnV, a.DependsOnGaugeCouplings, a.ConventionFactorsSealed, a.ObservedMassImported, a.Verdict)
}

func FormatNormalization(a NormalizationFirewallAudit) string {
	return fmt.Sprintf("bridge=[%s] environmental=[%s] nativeMass=%t nativeCoupling=%t nativeKphi=%t nativeV=%t nativeF0=%t nativeA=%t nativeMetric=%t nativeVacuum=%t verdict=%q", strings.Join(a.BridgeVariables, ", "), strings.Join(a.EnvironmentalVariables, ", "), a.NativeNumericalMassDerived, a.NativeCouplingDerived, a.NativeKphiDerived, a.NativeVDerived, a.NativeF0Derived, a.NativeYukawaTraceADerived, a.NativeScalarMetricDerived, a.NativeVacuumOrientationDerived, a.Verdict)
}

func FormatRelations(a RelationFirewallAudit) string {
	return fmt.Sprintf("q4=%t tauEta=%t Wspatial=%t pauliSeparate=%t flavor=%t observed=%t verdict=%q", a.Q4ContactOnly, a.TauEtaSigma3TraceShadow, a.WSpatialWeakPlaneBlocked, a.PauliQuaternionicSeparateRoute, a.FlavorDerived, a.ObservedDataImported, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("hessian=%t neutralNull=%t physicalDynamics=%t variables=[%s] flavorObserved=%t next=%q verdict=%q", a.SymbolicScalarKineticBridgeProducesHessian, a.NeutralHessianHasNullDirection, a.PhysicalWZPhotonDynamicsDerived, strings.Join(a.VariablesBridgeEnvironmental, ", "), a.FlavorOrObservedMassDataProduced, a.MissingNextTheorem, a.Verdict)
}
