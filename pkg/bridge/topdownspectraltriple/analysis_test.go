package topdownspectraltriple

import (
	"math"
	"testing"
)

func TestBuildDefaultTopDownSpectralTriple(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.Hilbert.Dimension != 16 || a.Hilbert.LeftDimension != 8 || a.Hilbert.RightDimension != 8 {
		t.Fatalf("unexpected Hilbert audit: %+v", a.Hilbert)
	}
	if !a.Triple.DiracSymmetric || !a.Triple.DiracOffDiagonal || a.Triple.YukawaChannelCount != 8 {
		t.Fatalf("bad Dirac support audit: %+v", a.Triple)
	}
	if !a.Triple.RealStructureInvolutive || !a.Triple.RealStructureCommutesWithD || !a.Triple.RealStructureAnticommutesGamma {
		t.Fatalf("bad J audit: %+v", a.Triple)
	}
	if !a.Triple.GammaInvolutive || !a.Triple.GammaTraceZero || !a.Triple.GammaAnticommutesWithD {
		t.Fatalf("bad gamma audit: %+v", a.Triple)
	}
}

func TestUnitIncidenceGaugeTrace(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if math.Abs(a.Gauge.KSU2T1-2) > 1e-10 || math.Abs(a.Gauge.KSU2T2-2) > 1e-10 || math.Abs(a.Gauge.KSU2T3-2) > 1e-10 {
		t.Fatalf("SU2 traces = %.12g %.12g %.12g, want 2", a.Gauge.KSU2T1, a.Gauge.KSU2T2, a.Gauge.KSU2T3)
	}
	if math.Abs(a.Gauge.KU1Y-(10.0/3.0)) > 1e-10 {
		t.Fatalf("Y trace = %.12g, want 10/3", a.Gauge.KU1Y)
	}
	if math.Abs(a.Gauge.NormalizedY-(5.0/3.0)) > 1e-10 {
		t.Fatalf("normalized Y = %.12g, want 5/3", a.Gauge.NormalizedY)
	}
	if math.Abs(a.Gauge.WeakAngleSeed-(3.0/8.0)) > 1e-10 {
		t.Fatalf("sin2 = %.12g, want 3/8", a.Gauge.WeakAngleSeed)
	}
	if !a.Gauge.BoundaryDiagMatched || !a.Gauge.WeakAngleSeedMatched {
		t.Fatalf("boundary trace should match: %+v", a.Gauge)
	}
}

func TestAmplitudeSensitivityFirewall(t *testing.T) {
	a, err := BuildDefault()
	if err != nil {
		t.Fatal(err)
	}
	if a.AmplitudeSensitivity.UnitAmplitudesDerivedByGate25 {
		t.Fatalf("Gate 25 must not derive unit amplitudes")
	}
	if a.AmplitudeSensitivity.BoundaryRatioStable || a.AmplitudeSensitivity.WeakAngleStable {
		t.Fatalf("deformation should break the ratio: %+v", a.AmplitudeSensitivity)
	}
	if math.Abs(a.AmplitudeSensitivity.DeformedRatioYOverSU2-(295.0/159.0)) > 1e-10 {
		t.Fatalf("deformed ratio = %.12g, want 295/159", a.AmplitudeSensitivity.DeformedRatioYOverSU2)
	}
	if math.Abs(a.AmplitudeSensitivity.DeformedWeakAngle-(159.0/454.0)) > 1e-10 {
		t.Fatalf("deformed sin2 = %.12g, want 159/454", a.AmplitudeSensitivity.DeformedWeakAngle)
	}
	if a.Firewall.ContactModeClassificationSolved || a.Firewall.ThresholdCorrectionsDerived || a.Firewall.RGRunningDerived || a.Firewall.PhysicalCouplingsDerived || a.Firewall.MassSpectrumDerived {
		t.Fatalf("firewall opened incorrectly: %+v", a.Firewall)
	}
}
