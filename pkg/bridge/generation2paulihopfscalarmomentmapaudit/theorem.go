package generation2paulihopfscalarmomentmapaudit

import "github.com/bagherbal/asha-engine/pkg/theorem"

func Generation2PauliHopfScalarMomentMapAuditTheorem() theorem.Theorem {
	const id = AuditID
	const name = "Generation 2 Pauli-Hopf scalar moment map audit"
	return theorem.Theorem{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Verify: func() theorem.Result {
		a, err := BuildDefault()
		if err != nil {
			return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.FailedRoute, Checks: []theorem.Check{{Name: "construct Gate560 Pauli-Hopf scalar moment audit", Passed: false, Detail: err.Error()}}}
		}
		checks := []theorem.Check{
			{Name: "inherit sealed Gate 558 eta=Sigma3 lane and Gate 559 transfer obstruction", Passed: a.Inherited.Gate558EtaIsSigma3Candidate && a.Inherited.Gate558HPhiSplitTwoPlusTwo && a.Inherited.Gate559NoLinearTransfer && a.Inherited.Gate559NoTraceRankTransfer, Detail: FormatInherited(a.Inherited)},
			{Name: "certify sealed scalar H_phi=R4~=C2 complex structure", Passed: a.Scalar.IdentityCertified && a.Scalar.RealDimension == 4 && a.Scalar.ComplexDimension == 2 && a.Scalar.SealedCarrierOnly && !a.Scalar.NativeUnsealed, Detail: FormatScalar(a.Scalar)},
			{Name: "construct real symmetric Pauli matrices on sealed H_phi", Passed: len(a.Pauli) == 3 && allPauliGood(a.Pauli), Detail: FormatPauliList(a.Pauli)},
			{Name: "verify Pauli/Cl(3,0) anticommutation relations", Passed: a.Relations.MatricesAvailable && a.Relations.SquaresIdentity && a.Relations.MaxSquareResidual <= 1e-9 && a.Relations.MaxAnticommutatorResidual <= 1e-9 && a.Relations.ConstructedUnderScalarSeal && !a.Relations.NativeUnsealed, Detail: FormatRelations(a.Relations)},
			{Name: "verify scalar moment coordinate formulas", Passed: a.Moment.CoordinatesMatch && len(a.Moment.SamplePoints) >= 4, Detail: FormatMoment(a.Moment)},
			{Name: "verify Hopf identity |mu|^2=(r^2)^2", Passed: a.Hopf.IdentityVerified && a.Hopf.SampleResidualMax <= 1e-9 && a.Hopf.ReliesOnSealedScalarC2Carrier, Detail: FormatHopf(a.Hopf)},
			{Name: "record scalar-sector 4=1+3 radius plus Pauli moment triplet", Passed: a.ScalarSplit.ScalarSectorFourToOnePlus3 && !a.ScalarSplit.IdentifiesGaugeBosons && !a.ScalarSplit.IdentifiesWSpatial && !a.ScalarSplit.IdentifiesWeakIsospin && !a.ScalarSplit.IdentifiesFlavor, Detail: FormatScalarSplit(a.ScalarSplit)},
			{Name: "record scalar-sector nonzero-moment 3=1+2 orbit split", Passed: a.Orbit.NonzeroMomentCondition && a.Orbit.RadialLineCanonical && a.Orbit.OrthogonalPlaneCanonicalGivenMu && a.Orbit.ScalarSectorOnly && !a.Orbit.SelectsWSpatialWeakPlane && !a.Orbit.SelectsGenerationPlane, Detail: FormatOrbit(a.Orbit)},
			{Name: "identify eta records as Sigma3-axis shadow", Passed: a.EtaRelation.EtaEqualsSigma3 && a.EtaRelation.Sigma3AxisShadowOnly && a.EtaRelation.LargerPauliTripletAvailable && !a.EtaRelation.TauEtaPromotedToSpectrum, Detail: FormatEtaRelation(a.EtaRelation)},
			{Name: "preserve transfer firewall to W_spatial, weak planes, and generations", Passed: a.Transfer.PauliMomentTripletAvailable && !a.Transfer.TransferAllowed && !a.Transfer.FunctorToWSpatial && !a.Transfer.FunctorToWeakPlaneCandidates && !a.Transfer.FunctorToGeneration && !a.Transfer.WeakPlaneSelected && !a.Transfer.GenerationHierarchyDerived && !a.Transfer.YukawaTextureDerived && !a.Transfer.CKMPMNSDerived && !a.Transfer.ObservedFlavorImported && !a.Transfer.GaugeBosonIdentification, Detail: FormatTransfer(a.Transfer)},
		}
		return theorem.Result{ID: id, Name: name, Layer: theorem.LayerBridge, Status: theorem.BridgeRequired, Checks: checks, Notes: append(Statuses(), a.Truth, a.Final.MissingNextTheorem)}
	}}
}

func allPauliGood(xs []PauliMatrixAudit) bool {
	if len(xs) != 3 {
		return false
	}
	for _, x := range xs {
		if !x.Symmetric || x.SquareResidual > 1e-9 || x.Rank != 4 || !x.ConstructibleSealed || x.NativeUnsealed {
			return false
		}
	}
	return true
}
