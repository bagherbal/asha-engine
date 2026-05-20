package generation2sealedspatialcp1fstcompatibilityaudit

import "testing"

func TestGate575SealedSpatialCP1CompatibilityWithFiniteSpectralTripleAudit(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Inherited.SealSufficient || !a.Inherited.SealMinimal || !a.Inherited.CP1CP0Split || !a.Inherited.CommutantU2U1 || a.Inherited.SealNative {
		t.Fatalf("bad inheritance: %s", FormatInherited(a.Inherited))
	}
	if !a.Decomposition.CP1CP0SplitExists || a.Decomposition.ProjectorRank != 1 || a.Decomposition.ComplementRank != 2 || a.Decomposition.OrthogonalityResidual > 1e-12 || a.Decomposition.NativeWithoutSeal {
		t.Fatalf("bad decomposition: %s", FormatSealedDecomposition(a.Decomposition))
	}
	if !a.BMinusL.CommutesWithPU || !a.BMinusL.CommutesWithComplement || !a.BMinusL.CompatibilityVacuous || a.BMinusL.SuppliesFurtherSelector {
		t.Fatalf("bad B-L compatibility: %s", FormatBMinusL(a.BMinusL))
	}
	if a.Commutant.Commutant != "u(2)+u(1)" || a.Commutant.Dimension != 5 || !a.Commutant.MatchesGate555Formula {
		t.Fatalf("bad commutant: %s", FormatCommutant(a.Commutant))
	}
	if !a.Quaternionic.ImHSocketAvailableElsewhere || !a.Quaternionic.HPhiDoubletModuleAvailable || !a.Quaternionic.WSpatialTransferBlocked || a.Quaternionic.ImHToSuUperpIntertwinerSupplied || a.Quaternionic.HToEndUperpModuleSupplied {
		t.Fatalf("bad quaternionic audit: %s", FormatQuaternionic(a.Quaternionic))
	}
	if !a.FiniteTriple.FiniteWeakDoubletCarrierExistsElsewhere || a.FiniteTriple.UperpUsedAsFiniteWeakDoubletCarrier || a.FiniteTriple.DCompatibilityForUperp || a.FiniteTriple.FirstOrderCompatibilityForUperp {
		t.Fatalf("bad finite spectral-triple carrier audit: %s", FormatFiniteTriple(a.FiniteTriple))
	}
	if !a.OneForm.FiniteOneFormContainsScalarDoublet || a.OneForm.SealedCP1AppearsInOneFormLane || !a.OneForm.PauliRouteSeparateFromWSpatial {
		t.Fatalf("bad one-form audit: %s", FormatOneForm(a.OneForm))
	}
	if a.WeakPlane.CanCallPhysicalWeakPlane || a.WeakPlane.FiniteSpectralTripleCompatible || a.WeakPlane.QuaternionicCompatible {
		t.Fatalf("bad weak-plane firewall: %s", FormatWeakPlane(a.WeakPlane))
	}
	if a.FlavorEW.GenerationHierarchyDerived || a.FlavorEW.YukawaTextureDerived || a.FlavorEW.CKMPMNSDerived || a.FlavorEW.WeakIsospinDerived {
		t.Fatalf("bad flavor/electroweak firewall: %s", FormatFlavorEW(a.FlavorEW))
	}
	if !a.Boundaries.PauliQuaternionicSocketNotWSpatial || !a.Boundaries.Gate564565BridgeSymbolic || !a.Boundaries.K7TimeRoutesSealed || !a.Boundaries.OrientationSealProjectiveOnly {
		t.Fatalf("bad previous boundaries: %s", FormatBoundaries(a.Boundaries))
	}
	if !a.Final.SealedCP1SplitExistsAlgebraically || !a.Final.CommutesWithBMinusL || a.Final.CarriesNativeOrSealedImHAction || a.Final.PartOfFiniteWeakDoubletCarrier || a.Final.CanBeCalledPhysicalWeakPlane || a.Final.DerivesFlavorOrEWObservedData {
		t.Fatalf("bad final verdict: %s", FormatFinal(a.Final))
	}
}

func TestGate575MatrixUtilities(t *testing.T) {
	p := [][]float64{{0, 0, 0}, {0, 0, 0}, {0, 0, 1}}
	q := subtract3(identity3(), p)
	if projectorIdempotentResidual(p) > 1e-12 || projectorIdempotentResidual(q) > 1e-12 {
		t.Fatalf("projectors not idempotent: P=%v Q=%v", p, q)
	}
	if rankFromTrace(p) != 1 || rankFromTrace(q) != 2 {
		t.Fatalf("bad ranks: trP=%v trQ=%v", trace3(p), trace3(q))
	}
	if productMaxAbs(p, q) > 1e-12 {
		t.Fatalf("P and Q not orthogonal")
	}
	bl := scale3(identity3(), 1.0/3.0)
	if commutatorMaxAbs(bl, p) > 1e-12 || commutatorMaxAbs(bl, q) > 1e-12 {
		t.Fatalf("B-L scalar restriction should commute")
	}
}

func TestGate575Theorem(t *testing.T) {
	res := Generation2SealedSpatialCP1CompatibilityWithFiniteSpectralTripleAuditTheorem().Run()
	if res.Status == "FAILED_ROUTE" {
		t.Fatalf("unexpected failed route: %+v", res)
	}
	for _, c := range res.Checks {
		if !c.Passed {
			t.Fatalf("check failed: %+v", c)
		}
	}
}
