package generation2scalarquaternionicmomentewcurvatureprojectionaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2ScalarQuaternionicMomentElectroweakCurvatureProjectionAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 scalar/quaternionic moment to electroweak curvature projection audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate563 scalar/quaternionic curvature projection audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit Gate 562 scalar/quaternionic moment bridge and firewalls", Passed: a.Inherited.Gate562QuaternionicBridge && a.Inherited.Gate562MomentMap && a.Inherited.Gate562StabilizerOrbitSplit && a.Inherited.Gate562PhysicalDynamicsFirewalled && a.Inherited.Gate562WSpatialTransferBlocked, Detail: FormatInherited(a.Inherited)},
			{Name: "recover finite one-form scalar SU(2)/H doublet lane", Passed: a.OneForm.HiggsDoubletRecovered && a.OneForm.ComplexDoublets == 1 && a.OneForm.RealScalarDimension == 4 && a.OneForm.StructuralDphiSocketFound && a.OneForm.ScalarSU2RepresentationClosed && a.OneForm.NumericalYukawaFree && !a.OneForm.HiggsPotentialDerived && !a.OneForm.HeatKernelProjectionAvailable, Detail: FormatOneForm(a.OneForm)},
			{Name: "verify structural Im(H) action and moment pairing on H_phi", Passed: a.Quaternionic.ImHSocketAvailable && a.Quaternionic.HphiDoubletModuleAvailable && a.Quaternionic.StructuralActionAvailable && a.Quaternionic.MomentPairingAvailable && !a.Quaternionic.CouplingNormalizationDerived, Detail: FormatQuaternionic(a.Quaternionic)},
			{Name: "locate symbolic Dphi squared and EW curvature sockets while blocking native moment projection", Passed: a.Kinetic.StructuralDphiSocketFound && a.Kinetic.ProductActionContainsDphiSquared && a.Kinetic.SymbolicKineticProjectionReadOff && !a.Kinetic.NativeScalarKineticCoefficientDerived && !a.Kinetic.NativeCanonicalScalarMetricDerived && a.Kinetic.ElectroweakCurvatureCarrierTyped && a.Kinetic.ElectroweakQuadraticFamilyTyped && !a.Kinetic.NativeElectroweakCurvatureAction && !a.Kinetic.FullSecondVariationComputed && !a.Kinetic.GaugeHessianActionSelected && !a.Kinetic.PhysicalGaugeCouplingsDerived && !a.Kinetic.MomentMapTermFoundInFiniteCurvature && !a.Kinetic.MomentMapTermFoundInKineticProjection, Detail: FormatKinetic(a.Kinetic)},
			{Name: "classify moment-map appearance as representation bookkeeping only", Passed: a.Moment.PhiPhiDaggerIdentityAvailable && a.Moment.MuSigmaExpressionAvailable && a.Moment.PairingMuXAvailable && !a.Moment.AppearsInFiniteOneForm && !a.Moment.AppearsInCurvature && !a.Moment.AppearsInScalarKineticProjection, Detail: FormatMoment(a.Moment)},
			{Name: "preserve stabilizer/orbit split only at scalar/quaternionic representation level", Passed: a.Orbit.NonzeroMuSplitAvailable && a.Orbit.RecognizedAtRepresentationLevel && !a.Orbit.CurvatureProjectionDistinguishesParts && !a.Orbit.StabilizerUsedForPhotonDirection && !a.Orbit.OrbitUsedForWZDirections, Detail: FormatOrbit(a.Orbit)},
			{Name: "block native U(1) mixing and photon direction", Passed: a.U1.AbelianSocketPresent && a.U1.AbelianNullDirectionDiagnostic && !a.U1.U1CompletionCoefficientSelected && !a.U1.WeakMixingAngleDerived && !a.U1.PhotonDirectionDerived && !a.U1.PhysicalElectroweakMixingNative, Detail: FormatU1(a.U1)},
			{Name: "block kinetic normalization and W/Z mass dynamics", Passed: a.Mass.SymbolicDphiSquaredChannel && !a.Mass.NativeScalarKineticCoefficient && !a.Mass.NativeGaugeKineticHessian && !a.Mass.ScalarVacuumOrientationDerived && !a.Mass.HiggsVEVDerived && !a.Mass.PhysicalWZMassMatrixDerived && !a.Mass.GaugeCouplingsDerived && !a.Mass.HiggsPotentialDerived, Detail: FormatMass(a.Mass)},
			{Name: "preserve flavor, q4, tau-eta, and W_spatial firewalls", Passed: !a.Flavor.YukawaEigenvaluesDerived && !a.Flavor.GenerationHierarchyDerived && !a.Flavor.CKMPMNSDerived && !a.Flavor.ObservedFlavorImported && !a.Flavor.Q4PromotedToHiggsFlavor && !a.Flavor.TauEtaPromotedToSpectrum && !a.Flavor.WSpatialWeakPlaneRouteReopened && a.Flavor.PauliRouteSeparateFromWSpatial, Detail: FormatFlavor(a.Flavor)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(Statuses(), a.Truth, a.Final.MissingNextTheorem)}
	}}
}
