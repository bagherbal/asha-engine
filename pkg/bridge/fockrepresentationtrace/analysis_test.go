package fockrepresentationtrace

import "testing"

func TestBuildDefaultRepresentationTrace(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.TraceAudit.FermionStates != 16 || a.TraceAudit.LeftStates != 8 || a.TraceAudit.RightStates != 8 || a.TraceAudit.SU2Doublets != 4 {
		t.Fatalf("bad representation dimensions: %+v", a.TraceAudit)
	}
	if !a.TraceAudit.KSU2T1.Equal(NewRational(2, 1)) || !a.TraceAudit.KSU2T2.Equal(NewRational(2, 1)) || !a.TraceAudit.KSU2T3.Equal(NewRational(2, 1)) {
		t.Fatalf("bad SU2 traces: %s %s %s", a.TraceAudit.KSU2T1, a.TraceAudit.KSU2T2, a.TraceAudit.KSU2T3)
	}
	if !a.TraceAudit.KU1Y.Equal(NewRational(10, 3)) {
		t.Fatalf("bad hypercharge trace: %s", a.TraceAudit.KU1Y)
	}
}

func TestBoundaryRatioIsRepresentationInvariant(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.TraceAudit.BoundaryDiagMatched || !a.TraceAudit.WeakAngleSeedMatched {
		t.Fatalf("boundary not matched: %+v", a.TraceAudit)
	}
	if !a.TraceAudit.NormalizedY.Equal(NewRational(5, 3)) {
		t.Fatalf("Y/SU2=%s, want 5/3", a.TraceAudit.NormalizedY)
	}
	if !a.TraceAudit.WeakAngleSeed.Equal(NewRational(3, 8)) {
		t.Fatalf("sin2=%s, want 3/8", a.TraceAudit.WeakAngleSeed)
	}
	if a.TraceAudit.UsesDiracFourthPower || !a.TraceAudit.AmplitudeIndependent {
		t.Fatalf("representation trace should not use D_F^4 amplitudes: %+v", a.TraceAudit)
	}
}

func TestAmplitudeSeparation(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if !a.Separation.DWeightedFunctionalAmplitudeDependent {
		t.Fatalf("Gate166 D_F^4 diagnostic should be amplitude dependent: %+v", a.Separation)
	}
	if !a.Separation.RepresentationTraceAmplitudeIndependent || !a.Separation.RepresentationUnitRatio.Equal(a.Separation.RepresentationDeformedRatio) || !a.Separation.RepresentationUnitSin2.Equal(a.Separation.RepresentationDeformedSin2) {
		t.Fatalf("representation trace should be amplitude independent: %+v", a.Separation)
	}
	if !a.Separation.Gate166DeformedDWeightedRatio.Equal(NewRational(295, 159)) {
		t.Fatalf("inherited deformed D4 ratio=%s, want 295/159", a.Separation.Gate166DeformedDWeightedRatio)
	}
	if !a.Separation.Gate166DeformedDWeightedSin2.Equal(NewRational(159, 454)) {
		t.Fatalf("inherited deformed D4 sin2=%s, want 159/454", a.Separation.Gate166DeformedDWeightedSin2)
	}
}

func TestYukawaAmplitudeProblemOpened(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.YukawaAudit.YukawaChannels != 8 || a.YukawaAudit.OneGenerationAmplitudeSlots != 8 || a.YukawaAudit.FermionKindBlocks != 4 {
		t.Fatalf("bad Yukawa audit: %+v", a.YukawaAudit)
	}
	if a.YukawaAudit.NumericalAmplitudesDerived || a.YukawaAudit.PhysicalMassesDerived || a.YukawaAudit.MixingMatricesDerived {
		t.Fatalf("masses/mixing should remain open: %+v", a.YukawaAudit)
	}
	if !a.YukawaAudit.MassEigenvaluesAreSingularValues || !a.YukawaAudit.CKMFromLeftMisalignment || !a.YukawaAudit.PMNSFromLeftMisalignment || !a.YukawaAudit.ConnectsGate28TextureSearch {
		t.Fatalf("finite Dirac texture recognition missing: %+v", a.YukawaAudit)
	}
}
