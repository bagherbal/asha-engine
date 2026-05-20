package generation2scalarquaternionicmomentewcurvatureprojectionaudit

import "testing"

func TestGate563ScalarQuaternionicMomentElectroweakCurvatureProjectionAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.Gate562QuaternionicBridge || !a.Inherited.Gate562MomentMap || !a.Inherited.Gate562StabilizerOrbitSplit || !a.Inherited.Gate562PhysicalDynamicsFirewalled || !a.Inherited.Gate562WSpatialTransferBlocked {
		t.Fatalf("inheritance failed: %s", FormatInherited(a.Inherited))
	}
	if !a.OneForm.HiggsDoubletRecovered || a.OneForm.ComplexDoublets != 1 || a.OneForm.RealScalarDimension != 4 || !a.OneForm.StructuralDphiSocketFound || !a.OneForm.ScalarSU2RepresentationClosed || !a.OneForm.NumericalYukawaFree || a.OneForm.HiggsPotentialDerived || a.OneForm.HeatKernelProjectionAvailable {
		t.Fatalf("finite one-form lane failed: %s", FormatOneForm(a.OneForm))
	}
	if !a.Quaternionic.ImHSocketAvailable || !a.Quaternionic.HphiDoubletModuleAvailable || !a.Quaternionic.StructuralActionAvailable || !a.Quaternionic.MomentPairingAvailable || a.Quaternionic.CouplingNormalizationDerived {
		t.Fatalf("quaternionic action failed: %s", FormatQuaternionic(a.Quaternionic))
	}
	if !a.Kinetic.StructuralDphiSocketFound || !a.Kinetic.ProductActionContainsDphiSquared || !a.Kinetic.SymbolicKineticProjectionReadOff || a.Kinetic.NativeScalarKineticCoefficientDerived || a.Kinetic.NativeCanonicalScalarMetricDerived || !a.Kinetic.ElectroweakCurvatureCarrierTyped || !a.Kinetic.ElectroweakQuadraticFamilyTyped || a.Kinetic.NativeElectroweakCurvatureAction || a.Kinetic.FullSecondVariationComputed || a.Kinetic.GaugeHessianActionSelected || a.Kinetic.PhysicalGaugeCouplingsDerived || a.Kinetic.MomentMapTermFoundInFiniteCurvature || a.Kinetic.MomentMapTermFoundInKineticProjection {
		t.Fatalf("curvature/kinetic audit failed: %s", FormatKinetic(a.Kinetic))
	}
	if !a.Moment.PhiPhiDaggerIdentityAvailable || !a.Moment.MuSigmaExpressionAvailable || !a.Moment.PairingMuXAvailable || a.Moment.AppearsInFiniteOneForm || a.Moment.AppearsInCurvature || a.Moment.AppearsInScalarKineticProjection {
		t.Fatalf("moment appearance failed: %s", FormatMoment(a.Moment))
	}
	if !a.Orbit.NonzeroMuSplitAvailable || !a.Orbit.RecognizedAtRepresentationLevel || a.Orbit.CurvatureProjectionDistinguishesParts || a.Orbit.StabilizerUsedForPhotonDirection || a.Orbit.OrbitUsedForWZDirections {
		t.Fatalf("orbit projection failed: %s", FormatOrbit(a.Orbit))
	}
	if !a.U1.AbelianSocketPresent || !a.U1.AbelianNullDirectionDiagnostic || a.U1.U1CompletionCoefficientSelected || a.U1.WeakMixingAngleDerived || a.U1.PhotonDirectionDerived || a.U1.PhysicalElectroweakMixingNative {
		t.Fatalf("U1 firewall failed: %s", FormatU1(a.U1))
	}
	if !a.Mass.SymbolicDphiSquaredChannel || a.Mass.NativeScalarKineticCoefficient || a.Mass.NativeGaugeKineticHessian || a.Mass.ScalarVacuumOrientationDerived || a.Mass.HiggsVEVDerived || a.Mass.PhysicalWZMassMatrixDerived || a.Mass.GaugeCouplingsDerived || a.Mass.HiggsPotentialDerived {
		t.Fatalf("kinetic/mass firewall failed: %s", FormatMass(a.Mass))
	}
	if a.Flavor.YukawaEigenvaluesDerived || a.Flavor.GenerationHierarchyDerived || a.Flavor.CKMPMNSDerived || a.Flavor.ObservedFlavorImported || a.Flavor.Q4PromotedToHiggsFlavor || a.Flavor.TauEtaPromotedToSpectrum || a.Flavor.WSpatialWeakPlaneRouteReopened || !a.Flavor.PauliRouteSeparateFromWSpatial {
		t.Fatalf("flavor firewall failed: %s", FormatFlavor(a.Flavor))
	}
	if !a.Final.FiniteOneFormContainsScalarDoublet || !a.Final.ImHActsStructurallyOnHphi || a.Final.MomentMapInsideCurvatureOrKineticData || !a.Final.MomentMapRepresentationBookkeepingOnly || a.Final.NonzeroMuCurvatureLevelSplit || a.Final.ElectroweakU1MixingDerived || a.Final.KineticNormalizationAndMassDynamics || a.Final.FlavorDataDerived {
		t.Fatalf("final verdict failed: %s", FormatFinal(a.Final))
	}
}

func TestGate563Theorem(t *testing.T) {
	res := Generation2ScalarQuaternionicMomentElectroweakCurvatureProjectionAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
