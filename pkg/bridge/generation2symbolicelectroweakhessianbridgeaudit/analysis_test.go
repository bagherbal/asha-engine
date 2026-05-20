package generation2symbolicelectroweakhessianbridgeaudit

import "testing"

func TestGate564SymbolicElectroweakHessianBridgeAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.Gate563ScalarDoubletLane || !a.Inherited.Gate563ImHActionOnHphi || !a.Inherited.Gate563MomentNotNativeCurvature || !a.Inherited.Gate563NoNativeU1PhotonDirection || !a.Inherited.Gate563NoNativeKineticNormalization || !a.Inherited.Gate563NoFlavorData {
		t.Fatalf("inheritance failed: %s", FormatInherited(a.Inherited))
	}
	if !a.Vacuum.VacuumBridgeSealed || a.Vacuum.VacuumDerivedNatively || !a.Vacuum.StabilizerSolvedSymbolically {
		t.Fatalf("vacuum bridge failed: %s", FormatVacuum(a.Vacuum))
	}
	if a.Charged.PerRealGeneratorCoefficient == "" || a.Charged.ObservedMassImported || a.Charged.NumericalCouplingImported {
		t.Fatalf("charged Hessian failed: %s", FormatCharged(a.Charged))
	}
	if a.Neutral.Matrix[0][0] != "g^2" || a.Neutral.Matrix[0][1] != "-g g'" || a.Neutral.Matrix[1][0] != "-g g'" || a.Neutral.Matrix[1][1] != "g'^2" || a.Neutral.Determinant != "0" || a.Neutral.Rank != 1 {
		t.Fatalf("neutral Hessian failed: %s", FormatNeutral(a.Neutral))
	}
	if !a.Null.DeterminantZero || !a.Null.PhotonSocketOnly || a.Null.PhysicalPhotonDerived || !a.Null.RequiresOSWickHilbertGaugeDynamics {
		t.Fatalf("null socket failed: %s", FormatNull(a.Null))
	}
	if !a.MassRatio.DependsOnKphi || !a.MassRatio.DependsOnV || !a.MassRatio.DependsOnGaugeCouplings || !a.MassRatio.ConventionFactorsSealed || a.MassRatio.ObservedMassImported {
		t.Fatalf("mass ratio failed: %s", FormatMassRatio(a.MassRatio))
	}
	if a.Normalization.NativeNumericalMassDerived || a.Normalization.NativeCouplingDerived || a.Normalization.NativeKphiDerived || a.Normalization.NativeVDerived || a.Normalization.NativeF0Derived || a.Normalization.NativeYukawaTraceADerived || a.Normalization.NativeScalarMetricDerived || a.Normalization.NativeVacuumOrientationDerived {
		t.Fatalf("normalization firewall failed: %s", FormatNormalization(a.Normalization))
	}
	if !a.Relations.Q4ContactOnly || !a.Relations.TauEtaSigma3TraceShadow || !a.Relations.WSpatialWeakPlaneBlocked || !a.Relations.PauliQuaternionicSeparateRoute || a.Relations.FlavorDerived || a.Relations.ObservedDataImported {
		t.Fatalf("relation firewall failed: %s", FormatRelations(a.Relations))
	}
	if !a.Final.SymbolicScalarKineticBridgeProducesHessian || !a.Final.NeutralHessianHasNullDirection || a.Final.PhysicalWZPhotonDynamicsDerived || a.Final.FlavorOrObservedMassDataProduced {
		t.Fatalf("final verdict failed: %s", FormatFinal(a.Final))
	}
}

func TestGate564Theorem(t *testing.T) {
	res := Generation2SymbolicElectroweakHessianBridgeAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
