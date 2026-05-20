package generation2boundarygaugenormalizationhessianaudit

import "testing"

func TestGate565BoundaryGaugeNormalizationHessianAlignmentAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.Gate564HessianShape || !a.Inherited.Gate564NeutralNull || !a.Inherited.Gate564NoPhysicalDynamics || !a.Inherited.Gate564NoFlavorData {
		t.Fatalf("inheritance failed: %s", FormatInherited(a.Inherited))
	}
	if !a.GaugeNorm.KYRecovered || a.GaugeNorm.KY != (Rational{5, 3}) || !a.GaugeNorm.BoundarySin2Recovered || a.GaugeNorm.ObservedInputUsed || a.GaugeNorm.LowEnergyObservedClaim {
		t.Fatalf("gauge norm failed: %s", FormatGaugeNorm(a.GaugeNorm))
	}
	if !a.Couplings.ConventionVerified || a.Couplings.RatioUnderBoundaryEquality != (Rational{3, 5}) || a.Couplings.NativePhysicalCouplingValue {
		t.Fatalf("coupling convention failed: %s", FormatCouplings(a.Couplings))
	}
	if a.BoundaryEquality.EqualityNativeTheorem || !a.BoundaryEquality.EqualityBridgeBoundary || a.BoundaryEquality.AbsoluteCouplingUnitDerived || a.BoundaryEquality.LowEnergyRunningDerived {
		t.Fatalf("boundary equality failed: %s", FormatBoundaryEquality(a.BoundaryEquality))
	}
	if !a.WeakAngle.MatchesPreviousASHA || a.WeakAngle.Sin2ThetaStar != (Rational{3, 8}) || a.WeakAngle.ObservedWeakAngleImported {
		t.Fatalf("weak angle failed: %s", FormatWeakAngle(a.WeakAngle))
	}
	if a.HessianRatio.BoundaryMW2OverMZ2 != (Rational{5, 8}) || a.HessianRatio.PhysicalLowEnergyMassRatio || a.HessianRatio.ObservedMassImported {
		t.Fatalf("Hessian ratio failed: %s", FormatHessianRatio(a.HessianRatio))
	}
	if a.Remaining.NativeAbsoluteKphi || a.Remaining.NativeV || a.Remaining.NativeAbsoluteG || a.Remaining.NativeAbsoluteGPrime || a.Remaining.NativeF0 || a.Remaining.NativeYukawaTraceA || a.Remaining.NativeScalarMetric || a.Remaining.NativeRGThresholds {
		t.Fatalf("remaining firewall failed: %s", FormatRemaining(a.Remaining))
	}
	if !a.PhotonFlavor.ASocketSymbolicOnly || a.PhotonFlavor.PhysicalPhotonDerived || a.PhotonFlavor.OSWickHilbertDerived || a.PhotonFlavor.YukawaEigenvalues || a.PhotonFlavor.CKMPMNS || a.PhotonFlavor.GenerationHierarchy || a.PhotonFlavor.ObservedFlavorData {
		t.Fatalf("photon/flavor firewall failed: %s", FormatPhotonFlavor(a.PhotonFlavor))
	}
	if !a.Final.KYRecoveredCorrectLayer || !a.Final.CouplingConventionVerified || !a.Final.Sin238Passes || !a.Final.HessianRatio58Passes || a.Final.PhysicalLowEnergyPrediction || a.Final.FlavorOrObservedDataProduced {
		t.Fatalf("final failed: %s", FormatFinal(a.Final))
	}
}

func TestGate565Theorem(t *testing.T) {
	res := Generation2BoundaryGaugeNormalizationToElectroweakHessianAlignmentAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
