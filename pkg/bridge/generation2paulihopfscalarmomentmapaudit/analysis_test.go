package generation2paulihopfscalarmomentmapaudit

import "testing"

func TestGate560PauliHopfScalarMomentMapAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.Gate558EtaIsSigma3Candidate || !a.Inherited.Gate559NoLinearTransfer || !a.Inherited.Gate559NoTraceRankTransfer {
		t.Fatalf("inheritance failed: %s", FormatInherited(a.Inherited))
	}
	if !a.Scalar.IdentityCertified || a.Scalar.RealDimension != 4 || a.Scalar.ComplexDimension != 2 || !a.Scalar.SealedCarrierOnly || a.Scalar.NativeUnsealed {
		t.Fatalf("scalar structure failed: %s", FormatScalar(a.Scalar))
	}
	if !allPauliGood(a.Pauli) {
		t.Fatalf("pauli matrices failed: %s", FormatPauliList(a.Pauli))
	}
	if !a.Relations.SquaresIdentity || a.Relations.MaxAnticommutatorResidual > 1e-9 || a.Relations.NativeUnsealed {
		t.Fatalf("Cl(3,0) relations failed: %s", FormatRelations(a.Relations))
	}
	if !a.Moment.CoordinatesMatch || len(a.Moment.SamplePoints) < 4 {
		t.Fatalf("moment coordinates failed: %s", FormatMoment(a.Moment))
	}
	if !a.Hopf.IdentityVerified || a.Hopf.SampleResidualMax > 1e-9 {
		t.Fatalf("Hopf identity failed: %s", FormatHopf(a.Hopf))
	}
	if !a.ScalarSplit.ScalarSectorFourToOnePlus3 || a.ScalarSplit.IdentifiesGaugeBosons || a.ScalarSplit.IdentifiesWSpatial || a.ScalarSplit.IdentifiesWeakIsospin || a.ScalarSplit.IdentifiesFlavor {
		t.Fatalf("scalar 4=1+3 firewall failed: %s", FormatScalarSplit(a.ScalarSplit))
	}
	if !a.Orbit.NonzeroMomentCondition || !a.Orbit.ScalarSectorOnly || a.Orbit.SelectsWSpatialWeakPlane || a.Orbit.SelectsGenerationPlane {
		t.Fatalf("moment 3=1+2 firewall failed: %s", FormatOrbit(a.Orbit))
	}
	if !a.EtaRelation.EtaEqualsSigma3 || !a.EtaRelation.Sigma3AxisShadowOnly || !a.EtaRelation.LargerPauliTripletAvailable || a.EtaRelation.TauEtaPromotedToSpectrum || a.EtaRelation.O1Residual > 1e-9 || a.EtaRelation.O2Residual > 1e-9 || a.EtaRelation.O3Residual > 1e-9 {
		t.Fatalf("eta relation failed: %s", FormatEtaRelation(a.EtaRelation))
	}
	if a.Transfer.TransferAllowed || a.Transfer.FunctorToWSpatial || a.Transfer.FunctorToWeakPlaneCandidates || a.Transfer.FunctorToGeneration || a.Transfer.WeakPlaneSelected || a.Transfer.GenerationHierarchyDerived || a.Transfer.YukawaTextureDerived || a.Transfer.CKMPMNSDerived || a.Transfer.ObservedFlavorImported || a.Transfer.GaugeBosonIdentification {
		t.Fatalf("transfer firewall failed: %s", FormatTransfer(a.Transfer))
	}
	if !a.Final.SealedPauliTripletExists || !a.Final.HopfMomentIdentityHolds || !a.Final.ScalarFourToOnePlusThree || !a.Final.NonzeroMomentThreeToOnePlusTwo || !a.Final.EtaIsSigma3Axis || a.Final.LawfulTransferToWOrGeneration {
		t.Fatalf("final verdict failed: %s", FormatFinal(a.Final))
	}
}

func TestGate560Theorem(t *testing.T) {
	res := Generation2PauliHopfScalarMomentMapAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
