package generation2scalarquaternionicmomentewcurvatureprojectionaudit

import "fmt"

func FormatInherited(a InheritedAudit) string {
	return fmt.Sprintf("bridge=%t moment=%t split=%t dynamicsFirewalled=%t WspatialBlocked=%t verdict=%q", a.Gate562QuaternionicBridge, a.Gate562MomentMap, a.Gate562StabilizerOrbitSplit, a.Gate562PhysicalDynamicsFirewalled, a.Gate562WSpatialTransferBlocked, a.Verdict)
}

func FormatOneForm(a FiniteOneFormScalarLaneAudit) string {
	return fmt.Sprintf("A=%q carrier=%q oneForms=%q DA=%q Higgs=%t complex=%d real=%d weak=%q color=%q Dphi=%t SU2prov=%t YukawaFree=%t potential=%t heatKernel=%t verdict=%q", a.Algebra, a.ScalarCarrier, a.OneFormsFormula, a.FluctuatedDiracFormula, a.HiggsDoubletRecovered, a.ComplexDoublets, a.RealScalarDimension, a.WeakRepresentation, a.ColorRepresentation, a.StructuralDphiSocketFound, a.ScalarSU2RepresentationClosed, a.NumericalYukawaFree, a.HiggsPotentialDerived, a.HeatKernelProjectionAvailable, a.Verdict)
}

func FormatQuaternionic(a QuaternionicActionAudit) string {
	return fmt.Sprintf("ImH=%t Hphi=%t action=%t pairing=%q available=%t layer=%q couplingNorm=%t verdict=%q", a.ImHSocketAvailable, a.HphiDoubletModuleAvailable, a.StructuralActionAvailable, a.PairingFormula, a.MomentPairingAvailable, a.AvailableLayer, a.CouplingNormalizationDerived, a.Verdict)
}

func FormatKinetic(a CurvatureKineticProjectionAudit) string {
	return fmt.Sprintf("Dphi=%t DphiSq=%t symbolic=%t coeff=%q nativeCoeff=%t nativeMetric=%t EWcurv=%t EWquad=%t nativeEW=%t secondVar=%t Hessian=%t couplings=%t muCurv=%t muKin=%t layer=%q verdict=%q", a.StructuralDphiSocketFound, a.ProductActionContainsDphiSquared, a.SymbolicKineticProjectionReadOff, a.KineticCoefficientSymbol, a.NativeScalarKineticCoefficientDerived, a.NativeCanonicalScalarMetricDerived, a.ElectroweakCurvatureCarrierTyped, a.ElectroweakQuadraticFamilyTyped, a.NativeElectroweakCurvatureAction, a.FullSecondVariationComputed, a.GaugeHessianActionSelected, a.PhysicalGaugeCouplingsDerived, a.MomentMapTermFoundInFiniteCurvature, a.MomentMapTermFoundInKineticProjection, a.MomentMapLayer, a.Verdict)
}

func FormatMoment(a MomentMapAppearanceAudit) string {
	return fmt.Sprintf("rho=%t muSigma=%t pair=%t oneForm=%t curvature=%t kinetic=%t layer=%q verdict=%q", a.PhiPhiDaggerIdentityAvailable, a.MuSigmaExpressionAvailable, a.PairingMuXAvailable, a.AppearsInFiniteOneForm, a.AppearsInCurvature, a.AppearsInScalarKineticProjection, a.ExactTheoremLayer, a.Verdict)
}

func FormatOrbit(a StabilizerOrbitProjectionAudit) string {
	return fmt.Sprintf("splitAvailable=%t split=%q representation=%t curvatureProjection=%t photon=%t WZ=%t verdict=%q", a.NonzeroMuSplitAvailable, a.Split, a.RecognizedAtRepresentationLevel, a.CurvatureProjectionDistinguishesParts, a.StabilizerUsedForPhotonDirection, a.OrbitUsedForWZDirections, a.Verdict)
}

func FormatU1(a U1MixingFirewallAudit) string {
	return fmt.Sprintf("abelian=%t nullDiagnostic=%t kappa=%t Ynorm=%t theta=%t photon=%t mixing=%t verdict=%q", a.AbelianSocketPresent, a.AbelianNullDirectionDiagnostic, a.U1CompletionCoefficientSelected, a.HyperchargeAbsoluteNormalizationFixed, a.WeakMixingAngleDerived, a.PhotonDirectionDerived, a.PhysicalElectroweakMixingNative, a.Verdict)
}

func FormatMass(a KineticMassFirewallAudit) string {
	return fmt.Sprintf("symbolicDphiSq=%t nativeScalarK=%t nativeGaugeHessian=%t vacuum=%t vev=%t WZ=%t couplings=%t potential=%t verdict=%q", a.SymbolicDphiSquaredChannel, a.NativeScalarKineticCoefficient, a.NativeGaugeKineticHessian, a.ScalarVacuumOrientationDerived, a.HiggsVEVDerived, a.PhysicalWZMassMatrixDerived, a.GaugeCouplingsDerived, a.HiggsPotentialDerived, a.Verdict)
}

func FormatFlavor(a FlavorFirewallAudit) string {
	return fmt.Sprintf("Yukawa=%t generation=%t CKM=%t observed=%t q4=%t tauSpec=%t Wroute=%t separate=%t verdict=%q", a.YukawaEigenvaluesDerived, a.GenerationHierarchyDerived, a.CKMPMNSDerived, a.ObservedFlavorImported, a.Q4PromotedToHiggsFlavor, a.TauEtaPromotedToSpectrum, a.WSpatialWeakPlaneRouteReopened, a.PauliRouteSeparateFromWSpatial, a.Verdict)
}

func FormatFinal(a FinalVerdict) string {
	return fmt.Sprintf("oneForm=%t ImHAction=%t muInCurvKin=%t bookkeepingOnly=%t curvSplit=%t U1=%t kineticMass=%t flavor=%t next=%q verdict=%q", a.FiniteOneFormContainsScalarDoublet, a.ImHActsStructurallyOnHphi, a.MomentMapInsideCurvatureOrKineticData, a.MomentMapRepresentationBookkeepingOnly, a.NonzeroMuCurvatureLevelSplit, a.ElectroweakU1MixingDerived, a.KineticNormalizationAndMassDynamics, a.FlavorDataDerived, a.MissingNextTheorem, a.Verdict)
}
